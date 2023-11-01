package aws

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
}

func GetS3Client(sdkConfig aws.Config) S3Client {
	return S3Client{Client: s3.NewFromConfig(sdkConfig)}
}

func (c *S3Client) UploadFile(bucketName string, objectKey string, file io.Reader) error {
	_, err := c.Client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
	if err != nil {
		log.Println("Error uploading file to S3: ", err)

		return err
	}

	return err
}
