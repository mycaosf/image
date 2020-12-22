package main

import (
	"image"
	"image/color"
)

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xffff

	return
}

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}

	r, g, b, _ := c.RGBA()
	return RGB{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8)}
}

type imageRGB struct {
	pix    []uint8
	rect   image.Rectangle
	stride int
}

func (p *imageRGB) ColorModel() color.Model {
	return color.ModelFunc(rgbModel)
}

func (p *imageRGB) Bounds() image.Rectangle {
	return p.rect
}

func (p *imageRGB) At(x, y int) color.Color {
	pos := (y-p.rect.Min.Y)*p.stride + (x-p.rect.Min.X)*3

	return RGB{R: p.pix[pos], G: p.pix[pos+1], B: p.pix[pos+2]}
}
