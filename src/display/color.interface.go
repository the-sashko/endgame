package display

import "github.com/veandco/go-sdl2/sdl"

type IColor interface {
	RGBA() (red uint32, green uint32, blue uint32, alpha uint32)

	GetSdlColor() sdl.Color

	GetRed() uint32

	SetRed(red uint32)

	GetGreen() uint32

	SetGreen(green uint32)

	GetBlue() uint32

	SetBlue(blue uint32)

	GetAlpha() uint32

	SetAlpha(alpha uint32)
}
