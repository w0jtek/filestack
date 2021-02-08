package downloader

import (
	"fmt"
	"image/jpeg"
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

	// // For simplicity, file will be saved to a hard-coded location
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
	_, err = jpeg.DecodeConfig(infile)
	if err != nil {
		return fmt.Errorf("it is not a valid jpeg file")
	}

	return nil
}
