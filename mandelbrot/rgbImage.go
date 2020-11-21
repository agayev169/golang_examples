package mandelbrot

import (
	"image"
	"image/color"
)

// GrayImage - 8-bit gray image
type GrayImage struct {
	bounds image.Rectangle
	pixels [][]color.Color
}

// ColorModel - Gray
func (img *GrayImage) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds - returns boundaries of an GrayImage
func (img *GrayImage) Bounds() image.Rectangle {
	return img.bounds
}

// At - pixel at (x, y) position
func (img *GrayImage) At(x, y int) color.Color {
	return img.pixels[y][x]
}

// Set - sets pixel at (x, y) position
func (img *GrayImage) Set(x, y int, val uint8) {
	var pixel = color.Gray{
		Y: val,
	}
	img.pixels[y][x] = pixel
}

// CreateImage - creates a GrayImage of (width, height) pixels
func CreateImage(width, height int) GrayImage {
	var imgBounds = image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: width,
			Y: height,
		},
	}

	var pixels = make([][]color.Color, height)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = make([]color.Color, width)
	}

	var img = GrayImage{
		bounds: imgBounds,
		pixels: pixels,
	}

	return img
}
