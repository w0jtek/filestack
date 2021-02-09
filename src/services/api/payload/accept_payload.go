package payload

import (
	"encoding/json"
	"errors"
	"fmt"
	dl "fs/src/services/api/downloader"
	"fs/src/services/api/response"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
)

// AcceptPayload defines the structure of data provided by the user
type AcceptPayload struct {
	ImageURL        string   `json:"imageUrl"`
	Transformations []string `json:"transformations"`
}

// NewAcceptPayload constructor
func NewAcceptPayload(r *http.Request) (payload *AcceptPayload, err error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("failed to read body")
	}

	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		return nil, errors.New("failed to unmarshal body")
	}

	fmt.Printf("%v", payload)

	return payload, nil
}

// Validate validates given payload
func (ap *AcceptPayload) Validate(localPath string) (imageOut image.Image, acceptResponse *response.AcceptResponse, imageType string) {
	imageType = ""
	if len(ap.ImageURL) < 1 {
		acceptResponse = response.NewAcceptResponse(400, "ImageUrl cannot be empty.")
		return
	}

	downloader := dl.NewDownloader()
	err := downloader.Fetch(ap.ImageURL, localPath)
	if err != nil {
		acceptResponse = response.NewAcceptResponse(400, err.Error())
		return
	}

	infile, _ := os.Open(localPath)
	image.RegisterFormat("png", "\x89PNG\r\n\x1a\n", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	imageOut, imageType, err = image.Decode(infile)
	if err != nil {
		acceptResponse = response.NewAcceptResponse(400, "it is not a valid png nor jpeg file")
		return
	}

	return
}
