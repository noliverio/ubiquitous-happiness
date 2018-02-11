package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

// creates a really interesting trippy holographic card look
func (m Image) At(x int, y int) color.Color {
	c := uint8(x ^ y*2)
	d := uint8(y - x)
	e := uint8(x * y)
	return color.RGBA{c, e, 255, d}
}

func (m Image) Bounds() image.Rectangle {
	rect := image.Rect(0, 0, m.Height, m.Width)
	return rect
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
