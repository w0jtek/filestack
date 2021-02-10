package transformer

import (
	"image"

	"github.com/oliamb/cutter"
)

// Crop applies image cropping
type Crop struct {
	x      int
	y      int
	width  int
	height int
}

// NewCrop constructor
func NewCrop(x int, y int, width int, height int) *Crop {
	return &Crop{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// Handle applies image cropping
func (t *Crop) Handle(localPath string) (err error) {

	img, imgType, err := DecodeImage(localPath)
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  t.width,
		Height: t.height,
		Anchor: image.Point{t.x, t.y},
	})

	return RewriteImage(croppedImg, imgType, localPath)
}
