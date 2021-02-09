package downloader

import (
	"fmt"
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
func (d *Downloader) Fetch(imageURL string, localPath string) error {
	defaultError := fmt.Errorf("fetching the image has failed")
	httpResponse, err := http.Get(imageURL)
	if httpResponse.StatusCode != 200 || err != nil {
		return defaultError
	}
	defer httpResponse.Body.Close()

	// For simplicity, file will be saved to a hard-coded location
	file, err := os.Create(localPath)
	if err != nil {
		return defaultError
	}
	defer file.Close()

	length, err := io.Copy(file, httpResponse.Body)
	if length == 0 || err != nil {
		return defaultError
	}

	return nil
}
