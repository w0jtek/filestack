package transformer

import (
	"fmt"
	"image/color"
)

// Rotate rotates image by multiple of 90 degrees
type Rotate struct {
	degrees int
}

// Rotate constructor
func NewRotate(degrees int) *Rotate {
	return &Rotate{
		degrees: degrees,
	}
}

// Handle applies image rotation
func (t *Rotate) Handle(localPath string) (err error) {
	if t.degrees%90 != 0 {
		return fmt.Errorf("Incorrect value for rotate - value should be a multiple of 90.")
	}
	img, imgType, err := DecodeImage(localPath)
	// rewriting image does the job of removing exif metadata
	return RewriteImage(img, imgType, localPath)
}

func RotateMatrixBy90(matrix [][]color.Color) [][]color.Color {
	srcHeight := len(matrix)
	srcWidth := len(matrix[0])

	targetHeight := srcWidth
	targetWidth := srcHeight

	var out [][]color.Color
	out = make([][]color.Color, targetHeight)
	for i := range out {
		out[i] = make([]color.Color, targetWidth)
	}

	var stream []color.Color

	for y := 0; y < srcHeight; y++ {
		for x := 0; x < srcWidth; x++ {
			stream = append(stream, matrix[y][x])
		}
	}

	j := 0
	for x := targetWidth - 1; x >= 0; x-- {
		for y := 0; y < targetHeight; y++ {
			out[y][x] = stream[j]
			j = j + 1
		}
	}

	return out
}

func RotateMatrixBy180(matrix [][]color.Color) [][]color.Color {
	height := len(matrix)
	width := len(matrix[0])

	var out [][]color.Color
	out = make([][]color.Color, height)
	for i := range out {
		out[i] = make([]color.Color, width)
	}

	var stream []color.Color

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			stream = append(stream, matrix[y][x])
		}
	}

	j := 0
	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			out[y][x] = stream[j]
			j = j + 1
		}
	}

	return out
}

func RotateMatrixBy270(matrix [][]color.Color) [][]color.Color {
	srcHeight := len(matrix)
	srcWidth := len(matrix[0])

	targetHeight := srcWidth
	targetWidth := srcHeight

	var out [][]color.Color
	out = make([][]color.Color, targetHeight)
	for i := range out {
		out[i] = make([]color.Color, targetWidth)
	}

	var stream []color.Color

	for y := 0; y < srcHeight; y++ {
		for x := 0; x < srcWidth; x++ {
			stream = append(stream, matrix[y][x])
		}
	}

	j := 0
	for x := 0; x <= targetWidth-1; x++ {
		for y := targetHeight - 1; y >= 0; y-- {
			out[y][x] = stream[j]
			j = j + 1
		}
	}

	return out
}
