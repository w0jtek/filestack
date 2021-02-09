package server

import (
	"fs/src/services/api/payload"
	"fs/src/services/api/response"
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
	acceptResponse, imageType := acceptPayload.Validate(localPath)
	if acceptResponse != nil {
		acceptResponse.Render(w)
		return
	}

	contentType := "image/" + imageType
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, localPath)
}
