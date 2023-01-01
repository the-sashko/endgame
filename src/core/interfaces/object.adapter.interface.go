package interfaces

type ICoreObjectAdapter interface {
	Destroy()
	DoHandleCollision(collidedObject ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton ICoreButton)
	DoHandleMouseButtonRelease(mouseButton ICoreButton)
	DoHandleMouseButtonClick(mouseButton ICoreButton)
	DoHandleMove(deltaX int32, deltaY int32)
}
