package server

import (
	"fs/src/services/api/compressor"
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

	_, _, err = transformer.DecodeImage(localPath)
	if err != nil {
		response.NewAcceptResponse(
			400,
			"This file is not a valid png nor jpeg image.",
		).Render(w)
		return
	}

	var transformedPaths []string
	for i, transformer := range transformers {
		pathTransformed, err := transformer.Handle(localPath, i)
		if err != nil {
			response.NewAcceptResponse(
				400,
				err.Error(),
			).Render(w)
			return
		}
		transformedPaths = append(transformedPaths, pathTransformed)
	}

	zipPath := "/tmp/archive.zip"
	zipCompressor := compressor.NewCompressor(transformedPaths)
	zipCompressor.Save(zipPath)

	w.Header().Set("Content-Type", "application/zip")
	http.ServeFile(w, r, zipPath)
}
