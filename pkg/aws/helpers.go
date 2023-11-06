package aws

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/pkg/errutil"
	"cybersafe-backend-api/pkg/helpers"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Imageler interface {
	SetImageURL(string)
}

func HandleImageAndUploadToS3(base64Image, bucketName, bucketFolder, bucketURL string, c *components.HTTPComponents, model Imageler, height, weight uint) error {
	if base64Image != "" {
		imgReader, err := helpers.HandleBase64Image(base64Image, height, weight)
		if err != nil {
			return errutil.ErrInvalidBase64Image
		}

		imgURL := fmt.Sprintf("%s/%s.png", bucketFolder, uuid.NewString())
		s3Client := GetS3Client(GetAWSConfig(c.Components))
		err = s3Client.UploadFile(bucketName, imgURL, imgReader)
		if err != nil {
			log.Println("Error uploading file to S3", err)

			return errutil.ErrUnexpectedError
		}

		model.SetImageURL(bucketURL + imgURL)
	}

	return nil
}
