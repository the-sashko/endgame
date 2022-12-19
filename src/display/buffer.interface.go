package display

type IBuffer interface {
	GetPixel(x uint16, y uint16) IPixel

	SetPixel(x uint16, y uint16, color IColor)

	GetChangedPixels() map[uint32]IPixel

	RemovePixelFromChanged(x uint16, y uint16)

	Clear()
}
