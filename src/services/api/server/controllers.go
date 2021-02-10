package server

import (
	"fs/src/services/api/payload"
	"fs/src/services/api/response"
	"fs/src/services/api/transformer"
	"net/http"
)

func handleAccept(w http.ResponseWriter, r *http.Request) {
	acceptPayload, err := payload.NewAcceptPayload(r)
	if err != nil {
		response.NewAcceptResponse(
			400,
			"Reading request body has failed.",
		).Render(w)
		return
	}

	localPath := "/tmp/image-fetched.tmp"
	acceptResponse, transformers := acceptPayload.Validate(localPath)
	if acceptResponse != nil {
		acceptResponse.Render(w)
		return
	}

	_, imgType, err := transformer.DecodeImage(localPath)
	if err != nil {
		response.NewAcceptResponse(
			400,
			"This file is not a valid png nor jpeg image.",
		).Render(w)
		return
	}

	for _, transformer := range transformers {
		transformer.Handle(localPath)
	}

	w.Header().Set("Content-Type", "image/"+imgType)
	http.ServeFile(w, r, localPath)
}
