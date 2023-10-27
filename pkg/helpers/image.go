package helpers

import (
	"bytes"
	"encoding/base64"
	"os"
	"strings"

	"github.com/google/uuid"
)

func ConvertBase64ImageToFile(base64Image string) (*os.File, error) {
	base64Image = strings.SplitN(base64Image, ",", 2)[1]

	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(imageData)

	tempFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	_, err = buffer.WriteTo(tempFile)
	if err != nil {
		return nil, err
	}

	newFileName := uuid.NewString() + ".jpg"
	os.Rename(tempFile.Name(), newFileName)

	file, err := os.Open(newFileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
