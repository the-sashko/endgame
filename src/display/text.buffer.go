package display

type textBuffer struct {
	//textsOnScreen map[string]uint32
	texts map[string]IText
}

func (textBufferObject *textBuffer) AddText(textObject IText) {
	textBufferObject.texts[textObject.GetReference()] = textObject
}

func (textBufferObject *textBuffer) GetText(reference string) IText {
	_, isExist := textBufferObject.texts[reference]

	if !isExist {
		return nil
	}

	return textBufferObject.texts[reference]
}

func (textBufferObject *textBuffer) GetTexts() map[string]IText {
	return textBufferObject.texts
}

func (textBufferObject *textBuffer) RemoveText(reference string) {
	_, isExist := textBufferObject.texts[reference]

	if !isExist {
		return
	}

	delete(textBufferObject.texts, reference)
}

func (textBufferObject *textBuffer) Clear() {
	for _, fontInstance := range fontInstances {
		fontInstance.Close()
	}

	textBufferObject.texts = make(map[string]IText)
}

func newTextBuffer() ITextBuffer {
	return &textBuffer{
		texts: make(map[string]IText),
	}
}
