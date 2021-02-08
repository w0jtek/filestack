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
	}

	acceptPayload.Validate().Render(w)
}
