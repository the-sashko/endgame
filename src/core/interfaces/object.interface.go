package interfaces

type ICoreObject interface {
	GetReference() string
	GetName() string
	GetShape() map[uint32]bool
	DoHandleCollision(collidedObject ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton ICoreButton)
	DoHandleMouseButtonRelease(mouseButton ICoreButton)
	DoHandleMouseButtonClick(mouseButton ICoreButton)
	DoHandleMove(deltaX uint16, deltaY uint16)
}
