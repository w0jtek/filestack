package transformer

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestRotateMap(t *testing.T) {

	sourceMatrix := [][]color.Color{
		{color.Black, color.Black, color.Black, color.White},
		{color.Black, color.White, color.White, color.White},
		{color.White, color.Black, color.White, color.Black},
	}

	t.Run("rotate 90 degrees", func(t *testing.T) {
		result := RotateMatrixBy90(sourceMatrix)

		expectedResult := [][]color.Color{
			{color.White, color.Black, color.Black},
			{color.Black, color.White, color.Black},
			{color.White, color.White, color.Black},
			{color.Black, color.White, color.White},
		}

		assert.Equal(t, len(expectedResult), len(result))

		if len(result) == 0 {
			t.Fatal("Expected not empty data set.")
		}

		assert.Equal(t, len(expectedResult[0]), len(result[0]))

		for x, row := range result {
			for y, _ := range row {
				assert.Equal(t, expectedResult[x][y], result[x][y])
			}
		}
	})

	t.Run("rotate 180 degrees", func(t *testing.T) {
		result := RotateMatrixBy180(sourceMatrix)

		expectedResult := [][]color.Color{
			{color.Black, color.White, color.Black, color.White},
			{color.White, color.White, color.White, color.Black},
			{color.White, color.Black, color.Black, color.Black},
		}

		assert.Equal(t, len(expectedResult), len(result))

		if len(result) == 0 {
			t.Fatal("Expected not empty data set.")
		}

		assert.Equal(t, len(expectedResult[0]), len(result[0]))

		for x, row := range result {
			for y, _ := range row {
				assert.Equal(t, expectedResult[x][y], result[x][y])
			}
		}
	})

	t.Run("rotate 270 degrees", func(t *testing.T) {
		result := RotateMatrixBy270(sourceMatrix)

		expectedResult := [][]color.Color{
			{color.White, color.White, color.Black},
			{color.Black, color.White, color.White},
			{color.Black, color.White, color.Black},
			{color.Black, color.Black, color.White},
		}

		assert.Equal(t, len(expectedResult), len(result))

		if len(result) == 0 {
			t.Fatal("Expected not empty data set.")
		}

		assert.Equal(t, len(expectedResult[0]), len(result[0]))

		for x, row := range result {
			for y, _ := range row {
				assert.Equal(t, expectedResult[x][y], result[x][y])
			}
		}
	})

}
