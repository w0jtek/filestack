package payload

import (
	"encoding/json"
	"errors"
	"fs/src/services/api/response"
	"io/ioutil"
	"net/http"
)

type AcceptPayload struct {
	ImageURL string `json:"imageUrl"`
}

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

func (ap *AcceptPayload) Validate() response.AcceptResponse {
	if len(ap.ImageURL) < 1 {
		return response.NewAcceptResponse(400, "ImageUrl cannot be empty.")
	}

	return response.NewAcceptResponse(200, "OK")
}
