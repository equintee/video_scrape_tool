package content_parser

import (
	"context"
	"log"
	"os"
	"video_scrape_tool/cmd/api/models"
	"video_scrape_tool/cmd/parsers/twitter"
	pb "video_scrape_tool/cmd/pb"
	"video_scrape_tool/cmd/util"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ContentDto struct {
	File    *os.File
	Song    models.Song
	Request ScrapeRequest
}
type ScrapeRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
}

type ContentService interface {
	Scrape(c echo.Context) error
}

type Service struct {
	twitterService          twitter.Service
	musicRecognitionService pb.MusicRecoginitonServiceClient
	minio                   util.ContentStorageInterface
	collection              *mongo.Collection
	bucketName              string
}

func NewService() *Service {
	instance := &Service{
		twitterService:          twitter.NewService(),
		musicRecognitionService: pb.GetMusicRecognitionService(),
		minio:                   util.GetInstance(),
		collection:              util.GetDatabase("contents"),
		bucketName:              "contents",
	}
	instance.minio.CreateBucket(instance.bucketName)
	return instance
}

func (s *Service) Scrape(c echo.Context) error {
	var request ScrapeRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, err)
		return nil
	}

	var content ContentDto

	switch request.Type {
	case "twitter":
		content.File, _ = s.twitterService.Scrape(request.Source)
	default:
		return nil
	}

	var entity models.Content
	err := copier.Copy(&entity, request)

	if err != nil {
		c.JSON(500, "Error copying entity.")
	}

	contentUrl := s.minio.SaveObject(s.bucketName, "", content.File.Name())
	entity.ContentUrl = contentUrl

	song := s.findSong(content.File)
	entity.Song = song

	//TODO: check if better way
	entity.Id = bson.NewObjectID()

	_, err = s.collection.InsertOne(context.Background(), entity)

	if err != nil {
		return err
	}

	c.JSON(200, entity)
	return nil
}

func (s *Service) findSong(file *os.File) models.Song {
	audioClip, err := util.ExtractAudioFromVideo(file)

	if err != nil {
		log.Println(err)
	}

	song, err := s.musicRecognitionService.RecognizeSong(context.Background(), &pb.RecognizeSongRequest{
		AudioClip: audioClip,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return models.Song{
		Name:   song.SongName,
		Artist: song.Artist,
	}
}
