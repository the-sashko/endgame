package display

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
)

type text struct {
	reference interfaces.ICoreReference
	value     string
	font      IFont
	color     IColor
	x         uint16
	y         uint16
}

func (textObject *text) GetReference() string {
	return textObject.reference.Get()
}

func (textObject *text) GetValue() string {
	return textObject.value
}

func (textObject *text) GetFont() IFont {
	return textObject.font
}

func (textObject *text) GetColor() IColor {
	return textObject.color
}

func (textObject *text) GetX() uint16 {
	return textObject.x
}

func (textObject *text) GetY() uint16 {
	return textObject.x
}

func newText(
	value string,
	x uint16,
	y uint16,
	fontName string,
	size byte,
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
) IText {
	return &text{
		reference: reference.NewReference(),
		value:     value,
		font:      getFontInstance(fontName, size),
		color:     NewColor(red, green, blue, alpha),
		x:         x,
		y:         y,
	}
}
