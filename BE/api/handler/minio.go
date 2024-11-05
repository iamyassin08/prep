package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iamyassin08/prep/storage"
	"github.com/minio/minio-go/v7"
)

// func removeFileHelper(folder string, id int32, c *fiber.Ctx) (string, error) {
// 	bucketName := os.Getenv("MINIO_BUCKET")
// 	file, err := c.FormFile("file")
// 	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
// 	fmt.Println(id)
// 	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
// 	if err != nil {
// 		return "", err
// 	}
// 	return "", nil
// }

func uploadFileHelper(folder string, id int32, c *fiber.Ctx) (string, error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	file, err := c.FormFile("file")
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println(id)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	if err != nil {
		return "", err
	}

	// Get Buffer from file
	buffer, err := file.Open()

	if err != nil {
		return "", err
	}
	defer buffer.Close()

	objectName := fmt.Sprintf("%s/%d/images/%s", folder, id, file.Filename)

	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size
	// fmt.Println(fileBuffer)
	fmt.Println(objectName)
	fmt.Println(contentType)
	// Upload the zip file with PutObject
	info, err := storage.MINIO_CLIENT.PutObject(c.Context(), bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return objectName, err
}
