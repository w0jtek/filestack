package transformer

import (
	"fmt"
	"image"
	"image/color"
	"os"
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
		return fmt.Errorf("Degrees must be a multiple of 90.")
	}

	img, imgType, err := DecodeImage(localPath)

	reader, err := os.Open(localPath)
	defer reader.Close()
	if err != nil {
		return fmt.Errorf("Cannot read file.")
	}

	imgConfig, _, err := image.DecodeConfig(reader)
	if err != nil {
		return fmt.Errorf("Cannot get image config.")
	}

	widthSrc := imgConfig.Width
	heightSrc := imgConfig.Height
	widthDest := widthSrc
	heightDest := heightSrc

	var matrixSrc [][]color.Color
	matrixSrc = make([][]color.Color, heightSrc)
	for i := range matrixSrc {
		matrixSrc[i] = make([]color.Color, widthSrc)
	}
	for y := 0; y < heightSrc; y++ {
		for x := 0; x < widthSrc; x++ {
			matrixSrc[y][x] = img.At(x, y)
		}
	}

	degreesModulo := (t.degrees / 90) % 4
	if degreesModulo == 0 {
		return nil
	}

	if degreesModulo == 1 || degreesModulo == 3 {
		widthDest = heightSrc
		heightDest = widthSrc
	}
	matrixDest := RotateMatrix(matrixSrc, degreesModulo)

	imgDest := image.NewRGBA(image.Rect(0, 0, widthDest-1, heightDest-1))
	for y := 0; y < heightDest; y++ {
		for x := 0; x < widthDest; x++ {
			imgDest.Set(x, y, matrixDest[y][x])
		}
	}

	// rewriting image does the job of removing exif metadata
	return RewriteImage(imgDest, imgType, localPath)
}

func RotateMatrix(matrix [][]color.Color, rotations int) [][]color.Color {
	switch rotations {
	case 1:
		return RotateMatrixBy90(matrix)
	case 2:
		return RotateMatrixBy180(matrix)
	case 3:
		return RotateMatrixBy270(matrix)
	default:
		return matrix
	}
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
