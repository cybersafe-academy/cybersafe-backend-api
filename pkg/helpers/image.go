package helpers

import (
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func ConvertBase64ImageToFile(base64Image string) (*os.File, error) {
	// Remove the prefix
	imgData := strings.Split(base64Image, ",")[1]

	// Decode the base64 string
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgData))

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	// Create a new file
	file, err := os.Create("tmp.png")
	if err != nil {
		return nil, err
	}

	// Encode the image to file
	err = png.Encode(file, img)
	if err != nil {
		return nil, err
	}

	// Return the file pointer to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ResizeImage(inputFile *os.File, weight, height uint) (*os.File, error) {
	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, err
	}

	// Resize the image to width 1280px and height 720px
	img = resize.Resize(weight, height, img, resize.Lanczos3)

	outputFile, err := os.Create(uuid.NewString() + ".png")
	if err != nil {
		return nil, err
	}

	// Crie um novo arquivo de imagem de saída
	err = png.Encode(outputFile, img)
	if err != nil {
		return nil, err
	}

	// Return the file pointer to the beginning
	_, err = outputFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return outputFile, nil
}
