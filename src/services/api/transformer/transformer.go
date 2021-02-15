package transformer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
)

// Transformer provides functionality of image manipulation
type Transformer interface {
	Handle(localPath string, number int) (destPath string, err error)
}

// DecodeImage opens file and tries to decode png/jpeg
func DecodeImage(localPath string) (imgOut image.Image, imgType string, err error) {
	infile, _ := os.Open(localPath)
	image.RegisterFormat("png", "\x89PNG\r\n\x1a\n", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	return image.Decode(infile)
}

// WriteImage is responsible for saving an image to the file
func WriteImage(img image.Image, imgType string, localPath string, number int) (destPath string, err error) {
	destPath = localPath + "-" + strconv.Itoa(number) + "." + imgType
	targetImg, err := os.Create(destPath)
	if err != nil {
		err = fmt.Errorf("failed to create a new file")
		return
	}
	defer targetImg.Close()

	switch imgType {
	case "png":
		png.Encode(targetImg, img)
	case "jpeg":
		jpeg.Encode(targetImg, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	default:
		err = fmt.Errorf("no support for type %s", imgType)
		return
	}

	return
}
