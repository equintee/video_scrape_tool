package util

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ContentStorage struct {
	client *minio.Client
}
type ContentStorageInterface interface {
	CreateBucket(bucketName string)
	SaveObject(bucketName string, objectName string, filePath string) string
	GetChunk(bucketName string, objectName string, start int64, end int64) (*Chunk, error)
}

var instance *ContentStorage

func GetInstance() *ContentStorage {
	return instance
}

func init() {
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_USERNAME")
	secretAccessKey := os.Getenv("MINIO_PASSWORD")

	// Initialize minio database object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	instance = &ContentStorage{
		client: minioClient,
	}

	log.Printf("ContentStorage connection sucesfull") // minioClient is now set up
}

func (c *ContentStorage) CreateBucket(bucketName string) {
	exists, err := c.client.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		err = c.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	}
}

func (c *ContentStorage) SaveObject(bucketName string, objectName string, filePath string) string {
	exists, err := c.client.BucketExists(context.Background(), bucketName)

	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		log.Fatalln(errors.New("Bucket does not exist."))
	}

	if objectName == "" {
		extension := filepath.Ext(filePath)
		newUUID, _ := uuid.NewUUID()
		objectName = newUUID.String() + extension
	}
	object, err := c.client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})

	if err != nil {
		log.Fatalln("Error uploading file.")
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
	return objectName
}

type Chunk struct {
	Data  []byte
	Start int64
	End   int64
	Size  int64
}

func (c *ContentStorage) GetChunk(bucketName string, objectName string, start int64, end int64) (*Chunk, error) {
	options := minio.GetObjectOptions{}
	options.SetRange(start, end)
	object, err := c.client.GetObject(context.TODO(), bucketName, objectName, options)
	if err != nil {
		return nil, err
	}
	defer object.Close()

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	stat, err := c.client.GetObject(context.TODO(), bucketName, objectName, options)
	size, err := stat.Stat()
	if err != nil {
		return nil, err
	}

	var chunk Chunk
	chunk.Start = start
	chunk.End = min(end, size.Size-1)
	chunk.Size = size.Size
	chunk.Data = data

	return &chunk, nil
}
