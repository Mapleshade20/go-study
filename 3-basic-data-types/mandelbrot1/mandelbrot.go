package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		subpixels              = 2 // 2x2 subpixels per pixel
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// Initialize color components for averaging
			var r, g, b, a uint32

			// Calculate color for each subpixel
			for sy := 0; sy < subpixels; sy++ {
				for sx := 0; sx < subpixels; sx++ {
					x := float64(px*subpixels+sx)/float64(width*subpixels)*(xmax-xmin) + xmin
					y := float64(py*subpixels+sy)/float64(height*subpixels)*(ymax-ymin) + ymin
					z := complex(x, y)
					subColor := mandelbrot(z)
					sr, sg, sb, sa := subColor.RGBA()
					r += sr
					g += sg
					b += sb
					a += sa
				}
			}

			// Average the color components
			totalSubpixels := uint32(subpixels * subpixels)
			r /= totalSubpixels
			g /= totalSubpixels
			b /= totalSubpixels
			a /= totalSubpixels

			// Set the averaged color for the pixel
			img.Set(px, py, color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 250
		contrast   = 50
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 255 - contrast/2*n, 255 - contrast/4*n, 255}
		}
	}
	return color.Black
}
