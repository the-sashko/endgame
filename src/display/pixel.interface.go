package display

type IPixel interface {
	GetColor() IColor

	SetColor(color IColor)

	IsChanged() bool

	SetChanged(isChanged bool)
}
