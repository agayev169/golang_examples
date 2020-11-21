# Mandelbrot Set

Implementation of the Mandelbrot set in golang.

The result of the algorithm with `setBeginX = -2.5`, `setBeginY = -1.0` and `stepSize = 0.002` for an 1920x1080 image:
![Result of the algorithm](imgs/img.png)

## Installation
`go get github.com/agayev169/golang_examples/mandelbrot`

## Usage
```
package main

import (
	"github.com/agayev169/golang_examples/mandelbrot"
	"image/png"
	"os"
)

func main() {
	img := mandelbrot.Mandelbrot(1920, 1080, -2.5, -1.0, 0.002)

	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
    
	png.Encode(f, img)
}

```