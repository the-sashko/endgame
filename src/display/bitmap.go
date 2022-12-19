package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type bitmapImage struct {
	name   string
	pixels map[uint32]IColor
}

func (bitmapImageObject *bitmapImage) loadFromFile(filePath string) {
	bmpSurface, err := sdl.LoadBMP(filePath)

	if err != nil {
		doError(err)
	}

	for x := uint16(0); x < uint16(bmpSurface.W); x++ {
		for y := uint16(0); y < uint16(bmpSurface.H); y++ {
			red, green, blue, alpha := bmpSurface.At(int(x), int(y)).RGBA()

			index := uint32(x)
			index = (index << 16) | uint32(y)

			color := NewColor(red, green, blue, alpha)

			bitmapImageObject.pixels[index] = color
		}
	}
}

func (bitmapImageObject *bitmapImage) GetName() string {
	return bitmapImageObject.name
}

func (bitmapImageObject *bitmapImage) GetPixels() map[uint32]IColor {
	return bitmapImageObject.pixels
}

func newBitmapImage(name string, filePath string) IBitmapImage {
	bitmapImageInstance := &bitmapImage{
		name:   name,
		pixels: make(map[uint32]IColor),
	}

	bitmapImageInstance.loadFromFile(filePath)

	return bitmapImageInstance
}
