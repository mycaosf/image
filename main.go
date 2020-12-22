package main

import (
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"unsafe"
)

//#include "image.h"
import "C"

//export ImageEncode
func ImageEncode(p *C.imageEncodeParam) C.int {
	var img image.Image
	w := int(p.width)
	h := int(p.height)
	data := (*[0x7fffffff]uint8)(unsafe.Pointer(p.data))
	rc := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: w, Y: h}}

	switch p.c {
	case C.imageColorGray8:
		img = &image.Gray{Pix: data[:], Stride: w, Rect: rc}
	case C.imageColorGray16:
		img = &image.Gray16{Pix: data[:], Stride: w * 2, Rect: rc}
	case C.imageColorRGB:
		img = &imageRGB{pix: data[:], stride: w * 3, rect: rc}
	case C.imageColorRGBA:
		img = &image.NRGBA{Pix: data[:], Stride: w * 4, Rect: rc}
	case C.imageColorRGBA64:
		img = &image.NRGBA64{Pix: data[:], Stride: w * 8, Rect: rc}
	case C.imageColorCMYK:
		img = &image.CMYK{Pix: data[:], Stride: w * 4, Rect: rc}
	default:
		return -1
	}

	fileName := C.GoString(p.fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return -2
	}

	defer file.Close()

	switch p.t {
	case C.imageTypePNG:
		err = png.Encode(file, img)
	case C.imageTypeJPEG:
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	case C.imageTypeBMP:
		err = bmp.Encode(file, img)
	case C.imageTypeTIFF:
		err = tiff.Encode(file, img, &tiff.Options{Compression: tiff.LZW, Predictor: true})
	default:
		return -3
	}

	if err != nil {
		return -4
	}

	return 1
}

func main() {
}
