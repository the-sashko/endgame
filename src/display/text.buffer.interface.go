package display

type ITextBuffer interface {
	AddText(textObject IText)
	GetText(reference string) IText
	GetTexts() map[string]IText
	RemoveText(reference string)
	Clear()
}
