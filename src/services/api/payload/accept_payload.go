package payload

import (
	"encoding/json"
	"errors"
	"fmt"
	dl "fs/src/services/api/downloader"
	"fs/src/services/api/response"
	tr "fs/src/services/api/transformer"
	"io/ioutil"
	"net/http"
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
func (ap *AcceptPayload) Validate(localPath string) (acceptResponse *response.AcceptResponse, transformers []tr.Transformer) {
	if len(ap.ImageURL) < 1 {
		acceptResponse = response.NewAcceptResponse(400, "ImageUrl cannot be empty.")
		return
	}

	for _, command := range ap.Transformations {
		transformer, err := tr.Create(command)
		if err != nil {
			acceptResponse = response.NewAcceptResponse(400, err.Error())
			return
		}
		transformers = append(transformers, transformer)
	}
	downloader := dl.NewDownloader()
	err := downloader.Fetch(ap.ImageURL, localPath)
	if err != nil {
		acceptResponse = response.NewAcceptResponse(400, err.Error())
		return
	}

	return
}
