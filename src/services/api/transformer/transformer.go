package transformer

import (
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/oliamb/cutter"
)

// Transformer provides functionalities for image manipulation
type Transformer interface {
	RemoveExif() (err error)
	Crop(x int, y int, width int, height int) (err error)
}

// ImageTransformer implements Transformer interface
type ImageTransformer struct {
	SourceImg  image.Image
	TargetFile string
}

// RemoveExif removes metadata from the source image
func (it *ImageTransformer) RemoveExif() {

}

// Crop performs cropping
func (it *ImageTransformer) Crop(x int, y int, width int, height int) (err error) {
	croppedImg, err := cutter.Crop(it.SourceImg, cutter.Config{
		Width:  width,
		Height: height,
		Anchor: image.Point{x, y},
	})
	toimg, err := os.Create(it.TargetFile)
	if err != nil {
		log.Fatal(err)
	}
	defer toimg.Close()
	jpeg.Encode(toimg, croppedImg, &jpeg.Options{jpeg.DefaultQuality})

	return
}
