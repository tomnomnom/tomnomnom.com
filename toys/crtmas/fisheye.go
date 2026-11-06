// Stolen and lightly modified from https://github.com/unixpickle/feDisplacementMaps
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const Width = 800
const Height = 600

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: fisheye <output.png>")
		os.Exit(1)
	}

	image := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			distanceVecX := float64(x - Width/2)
			distanceVecY := float64(y - Height/2)
			mag := math.Sqrt(math.Pow(distanceVecX, 2) + math.Pow(distanceVecY, 2))
			maxMag := math.Sqrt2 * Width / 2
			displacement := (mag - maxMag/2) / (maxMag / 2)
			displacement /= maxMag / 128
			image.Set(x, y, color.RGBA{
				R: uint8(displacement*distanceVecX + 128),
				G: uint8(displacement*distanceVecY + 128),
				B: 128,
				A: 255,
			})
		}
	}

	out, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer out.Close()
	png.Encode(out, image)
}
