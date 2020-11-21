package mandelbrot

import (
	"image"
	"math"
)

// Mandelbrot - returns Mandelbrot set as a GrayImage
func Mandelbrot(imgWidth, imgHeight int, setBeginX, setBeginY, stepSize float64) image.Image {
	img := CreateImage(imgWidth, imgHeight)

	for y, setCurrentY := 0, setBeginY; y < imgHeight; y, setCurrentY = y+1, setCurrentY+stepSize {
		for x, setCurrentX := 0, setBeginX; x < imgWidth; x, setCurrentX = x+1, setCurrentX+stepSize {
			value := mandelbrotFunction(complex(setCurrentX, setCurrentY), 250, 16.0)
			img.Set(x, y, value)
		}
	}

	return &img
}

// mandelbrotFunction - returns the pixel value for (x, y) position
func mandelbrotFunction(c complex128, steps int, threshold float64) uint8 {
	res := uint8(255)
	var z complex128 = 0

	for i := 0; i < steps; i++ {
		z = z*z + c

		magn := mag(z)
		if magn > threshold {
			return res
		}
		res--
	}

	return res
}

// mag - returns the magnitude of a complex number
func mag(c complex128) float64 {
	r := real(c)
	i := imag(c)

	return math.Sqrt(r*r + i*i)
}
