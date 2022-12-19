package display

type IDisplay interface {
	Quit()

	GetWidth() uint16

	GetHeight() uint16

	GetWindowTitle() string

	SetWindowTitle(title string)

	DrawBitmapImage(
		bitmapImageObject IBitmapImage,
		x uint16,
		y uint16,
		bufferLayer byte,
	)

	DrawBitmapImageToDefaultBuffer(
		bitmapImageObject IBitmapImage,
		x uint16,
		y uint16,
	)

	DrawPixel(
		x uint16,
		y uint16,
		red uint32,
		green uint32,
		blue uint32,
		alpha uint32,
		bufferLayer byte,
	)

	DrawPixelToDefaultBuffer(
		x uint16,
		y uint16,
		red uint32,
		green uint32,
		blue uint32,
		alpha uint32,
	)

	AddText(
		value string,
		x uint16,
		y uint16,
		fontName string,
		size byte,
		red uint32,
		green uint32,
		blue uint32,
		alpha uint32,
	) string

	AddTextWithDefaultStyle(
		value string,
		x uint16,
		y uint16,
	) string

	GetText(reference string) IText

	GetTexts() map[string]IText

	RemoveText(reference string)

	RemoveAllTexts()

	ClearBuffer(bufferLayer byte)

	ClearDefaultBuffer()

	ClearAll()

	ClearScreen()

	Render()
}
