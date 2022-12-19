package display

import "github.com/veandco/go-sdl2/sdl"

type color struct {
	red   uint32
	green uint32
	blue  uint32
	alpha uint32
}

func (colorObject *color) RGBA() (
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
) {
	return colorObject.GetRed(),
		colorObject.GetGreen(),
		colorObject.GetBlue(),
		colorObject.GetAlpha()
}

func (colorObject *color) GetSdlColor() sdl.Color {
	return sdl.Color{
		R: byte(colorObject.GetRed()),
		G: byte(colorObject.GetGreen()),
		B: byte(colorObject.GetBlue()),
		A: byte(colorObject.GetAlpha()),
	}
}

func (colorObject *color) GetRed() uint32 {
	return colorObject.red
}

func (colorObject *color) SetRed(red uint32) {
	colorObject.red = red
}

func (colorObject *color) GetGreen() uint32 {
	return colorObject.green
}

func (colorObject *color) SetGreen(green uint32) {
	colorObject.green = green
}

func (colorObject *color) GetBlue() uint32 {
	return colorObject.blue
}

func (colorObject *color) SetBlue(blue uint32) {
	colorObject.blue = blue
}

func (colorObject *color) GetAlpha() uint32 {
	return colorObject.alpha
}

func (colorObject *color) SetAlpha(alpha uint32) {
	colorObject.alpha = alpha
}

func NewColor(
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
) IColor {
	return &color{
		red:   red,
		blue:  blue,
		green: green,
		alpha: alpha,
	}
}
