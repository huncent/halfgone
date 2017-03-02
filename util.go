package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"math"
	"math/rand"
	"os"
)

// LoadImage reads and loads an image from a file path.
func LoadImage(path string) (image.Image, error) {
	infile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImagePNG save an image to a PNG file.
func SaveImagePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, img)
	return nil
}

// min returns the smallest of two ints.
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Convert a uint16 to a int8 taking into account bounds.
func i16ToUI8(x int16) uint8 {
	switch {
	case x < 1:
		return uint8(0)
	case x > 254:
		return uint8(255)
	}
	return uint8(x)
}

// randInt generates a random int in range [min, max).
func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

// randPoint generates a point with random coordinates from some given bounds.
func randPoint(bounds image.Rectangle, rng *rand.Rand) image.Point {
	return image.Point{
		X: randInt(bounds.Min.X, bounds.Max.X, rng),
		Y: randInt(bounds.Min.Y, bounds.Max.Y, rng),
	}
}

// randColor generates a random non-transparent color.
func randColor() color.Color {
	return color.NRGBA{
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		uint8(rand.Intn(256)),
		255,
	}
}

// square returns the square of an integer.
func square(x int) int {
	return x * x
}

// euclideanDistance calculates the L2 distance between two points.
func euclideanDistance(a, b image.Point) float64 {
	return math.Pow(float64(square(b.X-a.X)+square(b.Y-a.Y)), 0.5)
}

// binarySearchInt searches for the index of the first value in a sorted slice which is above a
// given value.
func binarySearchInt(value int, ints []int) int {
	var (
		index = -1
		a     = 0
		b     = len(ints) / 2
		c     = len(ints) - 1
	)

	for a != b && b != c {
		if value <= ints[b] {
			c = b
			b = (a + c) / 2
		} else {
			a = b
			b = (a + c + 1) / 2
		}
		if value <= ints[b] {
			index = b
		}
	}

	return index
}
