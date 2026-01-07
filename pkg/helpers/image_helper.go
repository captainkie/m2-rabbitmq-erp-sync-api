package helpers

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func GetMimeTypeByFileName(fileName string) (string, error) {
	mime, err := mimetype.DetectFile(fileName)
	if err != nil {
		return "", err
	}
	return mime.String(), nil
}

func GetMimeTypeByImageData(data []byte) (string, error) {
	mime := mimetype.Detect(data)
	return mime.String(), nil
}

func ConvertImageToBase64(imagePath string) (string, error) {
	// Read the image file
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	// Read the image data
	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	// Encode the image data to base64
	base64String := base64.StdEncoding.EncodeToString(imageData)

	return base64String, nil
}
