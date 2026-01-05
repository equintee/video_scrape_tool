package content_parser

import (
	"bytes"
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"video_scrape_tool/cmd/api/models"
	"video_scrape_tool/cmd/parsers/twitter"
	pb "video_scrape_tool/cmd/pb"
	"video_scrape_tool/cmd/util"

	"github.com/gorilla/schema"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ContentDto struct {
	File    *os.File
	Song    models.Song
	Request ScrapeRequest
}

type ContentService interface {
	Scrape(c echo.Context) error
	UpdateContent(c echo.Context) error
	GetContentMetaData(c echo.Context) error
	GetContentChunk(c echo.Context) ([]byte, error)
	GetFilters(c echo.Context) (FiltersResponse, error)
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

type ScrapeRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
}

func (s *Service) Scrape(c echo.Context) error {
	var request ScrapeRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, err)
		return nil
	}

	if request.Tags == nil {
		request.Tags = make([]string, 0)
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

	var response ContentResponse
	copier.Copy(&response, &entity)
	if err = c.JSON(200, response); err != nil {
		return err
	}

	return nil
}

type GetContentRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Songs       []string `json:"song"`
	PageSize    int      `json:"page_size" minimum:"1"`
	PageNum     int      `json:"page_num" minimum:"1"`
}

type ContentResponse struct {
	Id          bson.ObjectID `json:"id" bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Tags        []string      `json:"tags"`
	ContentUrl  string        `json:"content_url" bson:"content_url"`
	Type        string        `json:"type"`
	Song        SongDto       `json:"song"`
}

type SongDto struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

func (s *Service) GetContentMetaData(c echo.Context) error {
	var request GetContentRequest
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&request, c.QueryParams()); err != nil {
		c.JSON(400, err)
		return err
	}
	var filters []bson.M

	if request.Name != "" {
		filters = append(filters, bson.M{"name": bson.M{"$regex": request.Name, "$options": "i"}})
	}
	if request.Description != "" {
		filters = append(filters, bson.M{"description": bson.M{"$regex": request.Description, "$options": "i"}})
	}

	if len(request.Tags) > 0 {
		filters = append(filters, bson.M{"tags": bson.M{"$in": request.Tags}})
	}

	if request.Source != "" {
		filters = append(filters, bson.M{"source": request.Source})
	}

	if request.Type != "" {
		filters = append(filters, bson.M{"type": request.Type})
	}

	if len(request.Songs) > 0 {
		filters = append(filters, bson.M{"song.name": bson.M{"$in": request.Songs}})
	}

	var filter bson.M
	if len(filters) > 0 {
		filter = bson.M{"$and": filters}
	} else {
		filter = bson.M{}
	}

	skip := int64(request.PageSize * (request.PageNum - 1))
	limit := int64(request.PageNum)
	findOptions := options.Find().SetSkip(skip).SetLimit(limit)
	result, err := s.collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return err
	}

	var contents []ContentResponse
	err = result.All(context.Background(), &contents)
	if err != nil {
		return err
	}

	if contents == nil {
		contents = []ContentResponse{}
	}

	if err = c.JSON(200, contents); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetContentChunk(c echo.Context) ([]byte, error) {
	contentId := c.QueryParam("contentId")
	rangeHeader := c.Request().Header.Get("Range")
	rangeHeader = strings.ReplaceAll(rangeHeader, "bytes=", "")
	split := strings.Split(rangeHeader, "-")

	start, err := strconv.ParseInt(split[0], 10, 64)
	var end int64
	end = start + 100000
	if split[1] != "" {
		end, err = strconv.ParseInt(split[1], 10, 64)
	}

	chunk, err := s.minio.GetChunk(s.bucketName, contentId, start, end)
	if err != nil {
		c.Response().Status = 204
		return nil, err
	}

	response := c.Response()
	response.Status = 206
	contentRange := "bytes " + strconv.FormatInt(chunk.Start, 10) + "-" + strconv.FormatInt(chunk.End, 10) + "/" + strconv.FormatInt(chunk.Size, 10)
	response.Header().Set("Content-Range", contentRange)

	if err = c.Stream(206, "video/mp4", bytes.NewReader(chunk.Data)); err != nil {
		return nil, err
	}
	return chunk.Data, err
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

type FiltersResponse struct {
	Tags  []string `json:"tags"`
	Songs []string `json:"songs"`
}

func (s *Service) GetFilters(c echo.Context) (FiltersResponse, error) {
	var tags []string
	s.collection.Distinct(context.Background(), "tags", bson.D{}).Decode(&tags)

	var songs []string
	s.collection.Distinct(context.Background(), "song.name", bson.D{}).Decode(&songs)

	response := &FiltersResponse{
		Tags:  tags,
		Songs: songs,
	}
	c.JSON(200, response)

	return *response, nil
}

type UpdateRequest struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func (s *Service) UpdateContent(c echo.Context) error {
	var request UpdateRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, err)
		return err
	}

	hex, _ := bson.ObjectIDFromHex(request.Id)
	cursor := s.collection.FindOne(context.Background(), bson.M{"_id": hex})
	if err := cursor.Err(); err != nil {
		c.JSON(422, "Content not found.")
	}
	var entity models.Content
	cursor.Decode(&entity)
	if err := copier.Copy(&entity, &request); err != nil {
		c.JSON(400, err)
	}

	copier.Copy(&entity, &request)
	update, _ := s.collection.ReplaceOne(context.Background(), bson.M{"_id": hex}, entity)
	c.JSON(200, update)
	return nil
}
