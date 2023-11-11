package helpers

import (
	"bytes"
	"cybersafe-backend-api/pkg/errutil"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"image/png"

	"io"
	"strings"

	"github.com/nfnt/resize"
)

func ConvertBase64ImageToFile(base64Image string) (*bytes.Buffer, error) {
	imageData := strings.Split(base64Image, ",")
	if len(imageData) != 2 {
		return nil, errutil.ErrInvalidBase64Image
	}

	decodedImage := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageData[1]))

	buffer := new(bytes.Buffer)
	_, err := io.Copy(buffer, decodedImage)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func ResizeImage(file *bytes.Buffer, weight, height uint) (*bytes.Buffer, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	resizedImage := resize.Resize(weight, height, img, resize.Lanczos3)

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, resizedImage)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func HandleBase64Image(base64Image string, weight, height uint) (io.Reader, error) {
	buffer, err := ConvertBase64ImageToFile(base64Image)
	if err != nil {
		return nil, err
	}

	resizedImage, err := ResizeImage(buffer, weight, height)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(resizedImage.Bytes())

	return reader, nil
}
