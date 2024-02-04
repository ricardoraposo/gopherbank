package utils

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Region           string = os.Getenv("AWS_REGION")
	awsAccessKeyID     string = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey string = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsConfig                 = aws.Config{
		Region:      s3Region,
		Credentials: credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSecretAccessKey, ""),
	}
)

func newS3Uploader() (*manager.Uploader, error) {
	client := s3.NewFromConfig(awsConfig)
	uploader := manager.NewUploader(client)
	return uploader, nil
}

func UploadToS3(file io.Reader, key string) (*manager.UploadOutput, error) {
	uploader, err := newS3Uploader()
	if err != nil {
		return nil, err
	}

	return uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("gopherbankblobs"),
		Key:    aws.String(key),
		Body:   file,
        ContentType: aws.String("image/png"),
		ACL:    "public-read",
	})
}
