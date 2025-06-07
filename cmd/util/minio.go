package util

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func init() {
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_USERNAME")
	secretAccessKey := os.Getenv("MINIO_PASSWORD")

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	client = minioClient
	log.Printf("Minio connection sucesfull") // minioClient is now set up
}

func CreateBucket(bucketName string) {
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	}
}

func SaveObject(bucketName string, objectName string, filePath string) string {
	exists, err := client.BucketExists(context.Background(), bucketName)

	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		log.Fatalln(errors.New("Bucket does not exist."))
	}
	object, err := client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})

	if err != nil {
		log.Fatalln("Error uploading file.")
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
	return objectName
}
