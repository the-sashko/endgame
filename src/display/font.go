package display

import "C"
import (
	"fmt"
	"github.com/veandco/go-sdl2/ttf"
)

const fontsDirPath = "../res/fonts/"

var fontInstances map[string]IFont

type font struct {
	name          string
	ttfFontObject *ttf.Font
}

func (fontObject *font) GetName() string {
	return fontObject.name
}

func (fontObject *font) GetTtfFont() *ttf.Font {
	return fontObject.ttfFontObject
}

func (fontObject *font) Close() {
	fontObject.ttfFontObject.Close()
}

func newFont(name string, filePath string, size byte) IFont {
	ttfFontObject, err := ttf.OpenFont(filePath, int(size))

	if err != nil {
		doError(err)
	}

	fontObject := &font{
		name:          name,
		ttfFontObject: ttfFontObject,
	}

	return fontObject
}

func getFontInstance(name string, size byte) IFont {
	if fontInstances == nil {
		fontInstances = make(map[string]IFont)
	}

	fontName := fmt.Sprintf("%s_%d", name, size)
	fontFilePath := fmt.Sprintf("%s%s.ttf", fontsDirPath, name)

	_, isExist := fontInstances[fontName]

	if !isExist {
		fontInstances[name] = newFont(fontName, fontFilePath, size)
	}

	return fontInstances[name]
}
