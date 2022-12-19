package display

type IBitmapImage interface {
	GetName() string
	GetPixels() map[uint32]IColor
}
