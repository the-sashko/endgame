package display

type IText interface {
	GetReference() string
	GetValue() string
	GetFont() IFont
	GetColor() IColor
	GetX() uint16
	GetY() uint16
}
