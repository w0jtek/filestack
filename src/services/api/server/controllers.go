package server

import (
	"fs/src/services/api/payload"
	"fs/src/services/api/response"
	"fs/src/services/api/transformer"
	"log"
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
	img, acceptResponse, imageType := acceptPayload.Validate(localPath)
	if acceptResponse != nil {
		acceptResponse.Render(w)
		return
	}

	transformer := transformer.ImageTransformer{
		SourceImg:  img,
		TargetFile: "/tmp/image-fetched2.tmp",
	}

	err = transformer.Crop(0, 0, 200, 100)
	if err != nil {
		log.Fatal(err)
	}

	contentType := "image/" + imageType
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, "/tmp/image-fetched2.tmp")
}
