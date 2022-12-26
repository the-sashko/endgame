package display

import "endgame/src/utils/map_index"

type buffer struct {
	pixels        map[uint16]map[uint16]IPixel
	changedPixels map[uint32]IPixel
}

func (bufferObject *buffer) GetPixel(x uint16, y uint16) IPixel {
	if !bufferObject.hasPixel(x, y) {
		return nil
	}

	return bufferObject.pixels[x][y]
}

func (bufferObject *buffer) SetPixel(x uint16, y uint16, color IColor) {
	if bufferObject.hasPixel(x, y) {
		bufferObject.GetPixel(x, y).SetColor(color)
	}

	_, isExist := bufferObject.pixels[x]

	if !isExist || bufferObject.pixels[x] == nil {
		bufferObject.pixels[x] = make(map[uint16]IPixel)
	}

	bufferObject.pixels[x][y] = newPixel(color)

	pixelObject := bufferObject.GetPixel(x, y)

	if pixelObject.IsChanged() {
		bufferObject.setPixelToChanged(x, y, pixelObject)
	}
}

func (bufferObject *buffer) GetChangedPixels() map[uint32]IPixel {
	return bufferObject.changedPixels
}

func (bufferObject *buffer) SetPixelChanged(x uint16, y uint16) {
	index := map_index.GetIndex(x, y)

	pixelObject := bufferObject.GetPixel(x, y)

	if pixelObject == nil {
		return
	}

	_, isExist := bufferObject.changedPixels[index]

	if isExist {
		return
	}

	pixelObject.SetChanged(true)
	bufferObject.changedPixels[index] = pixelObject
}

func (bufferObject *buffer) RemovePixelFromChanged(x uint16, y uint16) {
	index := map_index.GetIndex(x, y)

	_, isExist := bufferObject.changedPixels[index]

	if !isExist {
		return
	}

	delete(bufferObject.changedPixels, index)
}

func (bufferObject *buffer) Clear() {
	bufferObject.pixels = map[uint16]map[uint16]IPixel{}
	bufferObject.changedPixels = map[uint32]IPixel{}
}

func (bufferObject *buffer) hasPixel(x uint16, y uint16) bool {
	_, isExist := bufferObject.pixels[x]

	if !isExist || bufferObject.pixels[x] == nil {
		return false
	}

	_, isExist = bufferObject.pixels[x][y]

	if !isExist {
		return false
	}

	return bufferObject.pixels[x][y] != nil
}

func (bufferObject *buffer) setPixelToChanged(
	x uint16,
	y uint16,
	pixel IPixel,
) {
	index := map_index.GetIndex(x, y)

	bufferObject.changedPixels[index] = pixel
}

func newBuffer() IBuffer {
	return &buffer{
		pixels:        map[uint16]map[uint16]IPixel{},
		changedPixels: map[uint32]IPixel{},
	}
}
