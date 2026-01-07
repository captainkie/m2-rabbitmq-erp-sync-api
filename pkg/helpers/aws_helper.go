package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Connect to S3
func SetupS3Client() *s3.S3 {
	region := os.Getenv("AWS_S3_REGION")
	accessKey := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")

	awsSession, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	svc := s3.New(awsSession)
	return svc
}

// List all objects in a bucket return []string
func ListObjects(s3Client *s3.S3, bucket, prefix string) ([]string, error) {
	var objects []string

	resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})

	if err != nil {
		return objects, err
	}

	for _, obj := range resp.Contents {
		fileNameWithExtension := filepath.Base(*obj.Key)
		if strings.HasSuffix(fileNameWithExtension, "/") {
			// Handle folders, if needed
		} else {
			// Handle files
			objects = append(objects, fileNameWithExtension)
		}
	}

	return objects, nil
}

// get file from S3
func GetFile(s3Client *s3.S3, bucket, imageKey, localPath string) (*os.File, error) {
	file, err := os.Create(localPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(imageKey),
	})

	if err != nil {
		return nil, err
	}

	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Downloaded file: %s\n", localPath)

	return file, nil
}

// get file from S3 as base64
func GetFileAsBase64(s3Client *s3.S3, bucket, imageKey string) ([]byte, string, error) {
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(imageKey),
	})

	if err != nil {
		return nil, "", err
	}

	defer resp.Body.Close()

	// Read the object content
	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	// Get the MIME type
	mimeType := http.DetectContentType(fileContent)

	return fileContent, mimeType, nil
}

// Upload a file to S3
func UploadFile(s3Client *s3.S3, bucket, prefix, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(prefix + "/" + filepath.Base(filePath)),
		Body:   file,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded file: %s\n", filePath)

	return nil
}

// Delete file from S3
func DeleteFile(s3Client *s3.S3, bucket, imageKey string) error {
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(imageKey),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Deleted file: %s\n", imageKey)

	return nil
}

// Delete all objects and the folder itself
func DeleteObjectsAndFolder(s3Client *s3.S3, bucket, prefix string) error {
	// List all objects with the specified prefix
	resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})

	if err != nil {
		return err
	}

	// Delete each object
	for _, obj := range resp.Contents {
		_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(*obj.Key),
		})

		if err != nil {
			return err
		}

		fmt.Printf("Deleted object: %s\n", *obj.Key)
	}

	// Delete the folder itself
	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(prefix),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Deleted folder: %s\n", prefix)

	return nil
}
