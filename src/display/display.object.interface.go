package display

import "endgame/src/core/interfaces"

type IDisplayObject interface {
	GetReference() string
	GetName() string
	GetShape() map[uint32]bool
	GetCurrentTexture() map[uint32]IColor
	GetX() uint16
	GetY() uint16
	GetMap() IDisplayMap
	GetMapLayerName() string
	SetTexture(textureName string)
	SetX(x uint16)
	SetY(y uint16)
	SetMap(mapObject IDisplayMap)
	SetMapLayerName(mapLayerName string)
	Destroy()
	DoMove(deltaX int32, deltaY int32)
	DoHandleCollision(collidedObject interfaces.ICoreObject)
	DoHandleMouseHover()
	DoHandleMouseButtonPress(mouseButton interfaces.ICoreButton)
	DoHandleMouseButtonRelease(mouseButton interfaces.ICoreButton)
	DoHandleMouseButtonClick(mouseButton interfaces.ICoreButton)
	DoHandleMove(deltaX int32, deltaY int32)
}
