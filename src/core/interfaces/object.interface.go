package interfaces

type ICoreObject interface {
	GetReference() string
	GetName() string
	GetShape() map[uint32]bool
	Destroy()
	DoHandleCollision(collidedObject ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton ICoreButton)
	DoHandleMouseButtonRelease(mouseButton ICoreButton)
	DoHandleMouseButtonClick(mouseButton ICoreButton)
	DoHandleMove(deltaX int32, deltaY int32)
}
