package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func PutFile(sess *session.Session, bucket string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file " + filename)
		return err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(
		&s3manager.UploadInput{
			Bucket: &bucket,
			Key:    aws.String("profile-pictures/" + filename),
			Body:   file,
		})
	if err != nil {
		return err
	}

	return nil
}
