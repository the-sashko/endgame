package display

type pixel struct {
	color     IColor
	isChanged bool
}

func (pixelObject *pixel) GetColor() IColor {
	return pixelObject.color
}

func (pixelObject *pixel) SetColor(color IColor) {
	if pixelObject.color == color {
		return
	}

	pixelObject.color = color
	pixelObject.SetChanged(true)
}

func (pixelObject *pixel) IsChanged() bool {
	return pixelObject.isChanged
}

func (pixelObject *pixel) SetChanged(isChanged bool) {
	pixelObject.isChanged = isChanged
}

func newPixel(color IColor) IPixel {
	return &pixel{
		color:     color,
		isChanged: true,
	}
}
