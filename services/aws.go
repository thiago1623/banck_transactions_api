package services

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWSService strcut of service s3
type AWSService struct {
	S3Client *s3.Client
}

// NewAWSService access to s3 with IAM credentials
func NewAWSService(accessKeyID, secretAccessKey string) (*AWSService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return &AWSService{
		S3Client: client,
	}, nil
}

// UploadFile upload file to specific s3 bucket
func (awsSvc AWSService) UploadFile(bucketName string, bucketKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error while opening the file", err)
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file: %v\n", cerr)
		}
	}()
	_, err = awsSvc.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketKey),
		Body:   file,
	})

	if err != nil {
		log.Println("Error while uploading the file", err)
	}
	return err
}
