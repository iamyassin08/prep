package storage

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MINIO_CLIENT *minio.Client

func ConnectToMinioClient() (*minio.Client, error) {
	endpoint, _ := os.LookupEnv("MINIO_ENDPOINT")
	accessKeyID, _ := os.LookupEnv("MINIO_ACCESS_KEY")
	secretAccessKey, _ := os.LookupEnv("MINIO_SECRET_ACCESS_KEY")
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return &minio.Client{}, err
	}
	log.Printf("%#v\n", minioClient)
	return minioClient, nil
}
