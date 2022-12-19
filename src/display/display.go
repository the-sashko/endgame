package display

import (
	"endgame/src/utils/logger"
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const defaultBufferLayer = 0
const defaultTextFont = "default"
const defaultTextSize = 14
const defaultTextColorRed = 0xff
const defaultTextColorBlue = 0xff
const defaultTextColorGreen = 0xff
const defaultTextColorAlpha = 0xff

var displayInstance IDisplay

type display struct {
	window     *sdl.Window
	buffers    map[byte]IBuffer
	textBuffer ITextBuffer
}

func (displayObject *display) Quit() {
	defer func(window *sdl.Window) {
		err := window.Destroy()

		if err != nil {
			doError(err)
		}
	}(displayObject.window)

	defer sdl.Quit()
}

func (displayObject *display) GetWidth() uint16 {
	return uint16(displayObject.getSurface().W)
}

func (displayObject *display) GetHeight() uint16 {
	return uint16(displayObject.getSurface().H)
}

func (displayObject *display) GetWindowTitle() string {
	return displayObject.window.GetTitle()
}

func (displayObject *display) SetWindowTitle(title string) {
	displayObject.window.SetTitle(title)
}

func (displayObject *display) DrawBitmapImage(
	bitmapImageObject IBitmapImage,
	x uint16,
	y uint16,
	bufferLayer byte,
) {
	for index, pixelColor := range bitmapImageObject.GetPixels() {
		pixelX := index >> 16
		pixelY := index - (pixelX << 16)

		displayObject.DrawPixel(
			x+uint16(pixelX),
			y+uint16(pixelY),
			pixelColor.GetRed(),
			pixelColor.GetGreen(),
			pixelColor.GetBlue(),
			pixelColor.GetAlpha(),
			bufferLayer,
		)
	}
}

func (displayObject *display) DrawBitmapImageToDefaultBuffer(
	bitmapImageObject IBitmapImage,
	x uint16,
	y uint16,
) {
	displayObject.DrawBitmapImage(bitmapImageObject, x, y, defaultBufferLayer)
}

func (displayObject *display) DrawPixel(
	x uint16,
	y uint16,
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
	bufferLayer byte,
) {
	if x >= displayObject.GetWidth() || y >= displayObject.GetHeight() {
		return
	}

	colorObject := NewColor(red, green, blue, alpha)

	_, isExist := displayObject.buffers[bufferLayer]

	if !isExist {
		displayObject.buffers[bufferLayer] = newBuffer()
	}

	displayObject.buffers[bufferLayer].SetPixel(x, y, colorObject)
}

func (displayObject *display) DrawPixelToDefaultBuffer(
	x uint16,
	y uint16,
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
) {
	displayObject.DrawPixel(x, y, red, green, blue, alpha, defaultBufferLayer)
}

func (displayObject *display) AddText(
	value string,
	x uint16,
	y uint16,
	fontName string,
	size byte,
	red uint32,
	green uint32,
	blue uint32,
	alpha uint32,
) string {
	textObject := newText(value, x, y, fontName, size, red, green, blue, alpha)
	displayObject.textBuffer.AddText(textObject)

	return textObject.GetReference()
}

func (displayObject *display) AddTextWithDefaultStyle(
	value string,
	x uint16,
	y uint16,
) string {
	textObject := newText(
		value,
		x,
		y,
		defaultTextFont,
		defaultTextSize,
		defaultTextColorRed,
		defaultTextColorGreen,
		defaultTextColorBlue,
		defaultTextColorAlpha,
	)

	displayObject.textBuffer.AddText(textObject)

	return textObject.GetReference()
}

func (displayObject *display) GetText(reference string) IText {
	return displayObject.textBuffer.GetText(reference)
}

func (displayObject *display) GetTexts() map[string]IText {
	return displayObject.textBuffer.GetTexts()
}

func (displayObject *display) RemoveText(reference string) {
	displayObject.textBuffer.RemoveText(reference)
}

func (displayObject *display) RemoveAllTexts() {
	displayObject.textBuffer.Clear()
}

func (displayObject *display) ClearBuffer(bufferLayer byte) {
	_, isExist := displayObject.buffers[bufferLayer]

	if !isExist {
		return
	}

	displayObject.buffers[bufferLayer].Clear()
}

func (displayObject *display) ClearDefaultBuffer() {
	displayObject.ClearBuffer(defaultBufferLayer)
}

func (displayObject *display) ClearAll() {
	for bufferLayer := range displayObject.buffers {
		displayObject.ClearBuffer(bufferLayer)
	}

	displayObject.RemoveAllTexts()

	surface := displayObject.getSurface()

	err := surface.FillRect(nil, 0)

	if err != nil {
		doError(err)
	}
}

func (displayObject *display) ClearScreen() {
	displayObject.ClearAll()
	displayObject.initBuffers()
	displayObject.Render()
}

func (displayObject *display) Render() {
	displayObject.renderFromBuffers()
	displayObject.renderFromTextBuffer()

	err := displayObject.window.UpdateSurface()

	if err != nil {
		doError(err)
	}
}

func (displayObject *display) init(
	title string,
	width uint16,
	height uint16,
	isFullScreen bool,
) {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		doError(err)
	}

	displayObject.initWindow(title, width, height, isFullScreen)
	displayObject.initBuffers()

	surface := displayObject.getSurface()

	err = surface.FillRect(nil, 0)

	if err != nil {
		doError(err)
	}
}

func (displayObject *display) initWindow(
	title string,
	width uint16,
	height uint16,
	isFullScreen bool,
) {
	var flags uint32

	flags = sdl.WINDOW_SHOWN

	if isFullScreen {
		flags += sdl.WINDOW_FULLSCREEN
	}

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width),
		int32(height),
		flags,
	)

	if err != nil {
		doError(err)
	}

	displayObject.window = window
}

func (displayObject *display) initBuffers() {
	displayObject.buffers = make(map[byte]IBuffer)
	displayObject.buffers[defaultBufferLayer] = newBuffer()

	displayObject.textBuffer = newTextBuffer()
}

func (displayObject *display) renderFromBuffers() {
	for _, bufferObject := range displayObject.buffers {
		displayObject.renderFromBuffer(bufferObject)
	}
}

func (displayObject *display) renderFromBuffer(bufferObject IBuffer) {
	for index, pixelObject := range bufferObject.GetChangedPixels() {
		x := index >> 16
		y := index - (x << 16)

		displayObject.renderPixel(bufferObject, pixelObject, uint16(x), uint16(y))
	}
}

func (displayObject *display) renderFromTextBuffer() {
	for _, textObject := range displayObject.textBuffer.GetTexts() {
		displayObject.renderText(textObject)
	}
}

func (displayObject *display) renderText(textObject IText) {
	ttfFont := textObject.GetFont().GetTtfFont()

	textSurface, err := ttfFont.RenderUTF8Blended(
		textObject.GetValue(),
		textObject.GetColor().GetSdlColor(),
	)

	if err != nil {
		doError(err)
	}

	defer textSurface.Free()

	err = textSurface.Blit(
		nil,
		displayObject.getSurface(),
		&sdl.Rect{
			X: int32(textObject.GetX()),
			Y: int32(textObject.GetY()),
			W: 0,
			H: 0,
		},
	)

	/*var xx uint16
	  var yy uint16

	  for xx = 0; xx < uint16(textSurface.W); xx++ {
	  	for yy = 0; yy < uint16(textSurface.H); yy++ {
	  		displayObject.DrawPixelToDefaultBuffer(xx+textObject.GetX(),
	  			yy+textObject.GetY(), 255,
	  			0, 0, 255)
	  	}
	  }*/

	if err != nil {
		doError(err)
	}
}

func (displayObject *display) renderPixel(
	bufferObject IBuffer,
	pixelObject IPixel,
	x uint16,
	y uint16,
) {
	if !pixelObject.IsChanged() {
		bufferObject.RemovePixelFromChanged(x, y)

		return
	}

	pixelObject.SetChanged(false)

	surface, err := displayObject.window.GetSurface()

	if err != nil {
		doError(err)
	}

	surface.Set(int(x), int(y), pixelObject.GetColor())

	bufferObject.RemovePixelFromChanged(x, y)
}

func (displayObject *display) getSurface() *sdl.Surface {
	surface, err := displayObject.window.GetSurface()

	if err != nil {
		doError(err)
	}

	return surface
}

func Init(
	title string,
	width uint16,
	height uint16,
	isFullScreen bool,
) {
	if displayInstance != nil {
		err := errors.New("display instance is already exists")
		doError(err)
	}

	err := ttf.Init()

	if err != nil {
		return
	}

	displayObject := new(display)
	displayObject.init(title, width, height, isFullScreen)

	displayInstance = displayObject
}

func GetInstance() IDisplay {
	if displayInstance == nil {
		err := errors.New("display not initialized")
		doError(err)
	}

	return displayInstance
}

func doError(errorEntity error) {
	logger.LogError("display", errorEntity, true)
}
