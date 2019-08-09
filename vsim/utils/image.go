package utils

import (
	"image"
	"image/color"
	"math"
)

func ImageBoundsMatch(a, b image.Rectangle) bool {
	return a.Min.X == b.Min.X && a.Min.Y == b.Min.Y && a.Max.X == b.Max.X && a.Max.Y == b.Max.Y
}

func GetImageDiff(imageA, imageB image.Image) float64 {
	bounds := imageA.Bounds()

	var diff float64

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			diff += getColorDiff(imageA.At(x, y), imageB.At(x, y))
		}
	}

	return diff
}

func getColorDiff(colorA, colorB color.Color) (diff float64) {
	r1, g1, b1, a1 := colorA.RGBA()
	r2, g2, b2, a2 := colorB.RGBA()

	diff += math.Abs(float64(r1 - r2))
	diff += math.Abs(float64(g1 - g2))
	diff += math.Abs(float64(b1 - b2))
	diff += math.Abs(float64(a1 - a2))

	return diff / 255
}
