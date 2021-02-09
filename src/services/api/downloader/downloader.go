package downloader

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

// Downloader is responsible for fetching image files
type Downloader struct {
}

// NewDownloader constructor
func NewDownloader() *Downloader {
	return &Downloader{}
}

// Fetch downloads files
func (d *Downloader) Fetch(imageURL string) error {
	defaultError := fmt.Errorf("fetching the image has failed")
	httpResponse, err := http.Get(imageURL)
	if httpResponse.StatusCode != 200 || err != nil {
		return defaultError
	}
	defer httpResponse.Body.Close()

	// For simplicity, file will be saved to a hard-coded location
	localImagePath := "/tmp/image-fetched"
	file, err := os.Create(localImagePath)
	if err != nil {
		return defaultError
	}
	defer file.Close()

	length, err := io.Copy(file, httpResponse.Body)
	if length == 0 || err != nil {
		return defaultError
	}

	infile, _ := os.Open(localImagePath)
	image.RegisterFormat("png", "\x89PNG\r\n\x1a\n", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	_, _, err = image.Decode(infile)
	if err != nil {
		return fmt.Errorf("it is not a valid png nor jpeg file")
	}

	return nil
}
