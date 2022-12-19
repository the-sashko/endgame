package display

import "github.com/veandco/go-sdl2/ttf"

type IFont interface {
	GetName() string
	GetTtfFont() *ttf.Font
	Close()
}
