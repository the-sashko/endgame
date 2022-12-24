package interfaces

type ICoreKeyboard interface {
	GetButton(buttonName string) ICoreButton
}
