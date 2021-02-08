package payload

import (
	"encoding/json"
	"errors"
	dl "fs/src/services/api/downloader"
	"fs/src/services/api/response"
	"io/ioutil"
	"net/http"
)

// AcceptPayload defines the structure of data provided by the user
type AcceptPayload struct {
	ImageURL string `json:"imageUrl"`
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

	return payload, nil
}

// Validate validates given payload
func (ap *AcceptPayload) Validate() response.AcceptResponse {
	if len(ap.ImageURL) < 1 {
		return response.NewAcceptResponse(400, "ImageUrl cannot be empty.")
	}

	downloader := dl.NewDownloader()
	err := downloader.Fetch(ap.ImageURL)
	if err != nil {
		return response.NewAcceptResponse(400, err.Error())
	}

	return response.NewAcceptResponse(200, "OK")
}
