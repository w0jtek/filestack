package server

import (
	"fs/src/services/api/payload"
	"fs/src/services/api/response"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/oliamb/cutter"
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

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  100,
		Height: 100,
		Anchor: image.Point{100, 100},
	})
	toimg, err := os.Create("/tmp/image-fetched2.tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer toimg.Close()
	jpeg.Encode(toimg, croppedImg, &jpeg.Options{jpeg.DefaultQuality})

	contentType := "image/" + imageType
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, "/tmp/image-fetched2.tmp")

}
