package display

import "endgame/src/core/interfaces"

type IDisplayObject interface {
	GetReference() string
	GetName() string
	GetShape() map[uint32]bool
	GetCurrentTexture() map[uint32]IColor
	SetTexture(textureName string)
	DoHandleCollision(collidedObject interfaces.ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton interfaces.ICoreButton)
	DoHandleMouseButtonRelease(mouseButton interfaces.ICoreButton)
	DoHandleMouseButtonClick(mouseButton interfaces.ICoreButton)
	DoHandleMove(deltaX uint16, deltaY uint16)
}
