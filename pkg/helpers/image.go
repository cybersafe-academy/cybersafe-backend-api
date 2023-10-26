package helpers

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func ConvertBase64ImageToFile(base64String string, fileName string) (*os.File, error) {
	prefix := "data:image/"
	suffix := ";base64,"
	var imageType string

	if strings.HasPrefix(base64String, prefix) && strings.Contains(base64String, suffix) {
		start := len(prefix)
		end := strings.Index(base64String, ";base64")
		imageType = base64String[start:end]
		base64String = base64String[end+len(suffix):]
	}

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	outputFileName := fmt.Sprintf("%s.%s", fileName, imageType)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return nil, err
	}

	_, err = outputFile.Write(data)
	if err != nil {
		outputFile.Close()
		return nil, err
	}

	return outputFile, nil
}
