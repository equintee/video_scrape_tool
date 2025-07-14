package util

import (
	"context"
	"errors"
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
