package transformer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// Transformer provides functionality of image manipulation
type Transformer interface {
	Handle(localPath string) (err error)
}

// DecodeImage opens file and tries to decode png/jpeg
func DecodeImage(localPath string) (imgOut image.Image, imgType string, err error) {
	infile, _ := os.Open(localPath)
	image.RegisterFormat("png", "\x89PNG\r\n\x1a\n", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	return image.Decode(infile)
}

// RewriteImage is responsible for saving an image to the file
func RewriteImage(img image.Image, imgType string, localPath string) error {
	localTmpPath := localPath + ".tmp"
	targetImg, err := os.Create(localTmpPath)
	if err != nil {
		return fmt.Errorf("Failed to create a new file.")
	}
	defer targetImg.Close()

	switch imgType {
	case "png":
		png.Encode(targetImg, img)
	case "jpeg":
		jpeg.Encode(targetImg, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	default:
		return fmt.Errorf("No support for type %s", imgType)
	}

	os.Rename(localTmpPath, localPath)

	return nil
}
