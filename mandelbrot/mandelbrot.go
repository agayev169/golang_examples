package mandelbrot

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math"
)

// Mandelbrot - returns Mandelbrot set as a GrayImage
func Mandelbrot(imgWidth, imgHeight int, setBeginX, setBeginY, stepSize float64) image.Image {
	img := CreateImage(imgWidth, imgHeight)

	for y, setCurrentY := 0, setBeginY; y < imgHeight; y, setCurrentY = y+1, setCurrentY+stepSize {
		for x, setCurrentX := 0, setBeginX; x < imgWidth; x, setCurrentX = x+1, setCurrentX+stepSize {
			value := mandelbrotFunction(complex(setCurrentX, setCurrentY), 255, 16.0)
			img.Set(x, y, value)
		}
	}

	return &img
}

// Zoom - returns a GIF of zooming in the Mandelbrot set
func Zoom(imgWidth, imgHeight int, setCenterX, setCenterY, stepSizeBegin, stepCoef float64,
	frameCount uint) *gif.GIF {
	steps := make([]float64, frameCount, frameCount)
	for i := uint(0); i < frameCount; i++ {
		steps[i] = stepSizeBegin * math.Pow(stepCoef, float64(i))
	}

	res := &gif.GIF{}

	c := make(chan struct {
		Index uint
		Frame *image.Paletted
	}, 100)

	for i := uint(0); i < frameCount; i++ {
		setBeginX := setCenterX - float64(imgWidth/2)*steps[i]
		setBeginY := setCenterY - float64(imgHeight/2)*steps[i]
		go func(i uint, setBeginX, setBeginY float64, c chan<- struct {
			Index uint
			Frame *image.Paletted
		}) {
			frame := Mandelbrot(imgWidth, imgHeight, setBeginX, setBeginY, steps[i])
			paletted := image.NewPaletted(frame.Bounds(), palette.Plan9)
			draw.Draw(paletted, paletted.Rect, frame, paletted.Rect.Min, draw.Over)
			c <- struct {
				Index uint
				Frame *image.Paletted
			}{Index: i, Frame: paletted}
		}(i, setBeginX, setBeginY, c)
		res.Image = append(res.Image, nil)
		res.Delay = append(res.Delay, 0)
	}

	for i := uint(0); i < frameCount; i++ {
		frameInfo := <-c
		res.Image[frameInfo.Index] = frameInfo.Frame
	}

	return res
}

// mandelbrotFunction - returns the pixel value for (x, y) position
func mandelbrotFunction(c complex128, steps int, threshold float64) uint8 {
	res := uint8(steps)
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
