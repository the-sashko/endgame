package interfaces

type ICoreObjectAdapter interface {
	DoHandleCollision(collidedObject ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton ICoreButton)
	DoHandleMouseButtonRelease(mouseButton ICoreButton)
	DoHandleMouseButtonClick(mouseButton ICoreButton)
	DoHandleMove(deltaX uint16, deltaY uint16)
}
