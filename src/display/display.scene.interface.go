package display

import (
	"endgame/src/core/interfaces"
)

type IDisplayScene interface {
	GetX() uint16
	GetY() uint16
	GetWidth() uint16
	GetHeight() uint16
	SetX(x uint16)
	SetY(y uint16)
	SetDisplayBufferLayer(mapLayerName string, displayBufferLayer byte)
	GetPixels() map[byte]map[uint32]IColor
	SetObjects(objects map[string]map[uint32]interfaces.ICoreObject)
	DeleteDisplayBufferLayer(mapLayerName string)
	DoFillPixels()
}
