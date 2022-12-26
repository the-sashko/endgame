package display

import (
	"endgame/src/core/interfaces"
)

type IDisplayMap interface {
	GetReference() string
	GetName() string
	GetLayers() map[string]interfaces.ICoreMapLayer
	GetLayer(layerName string) interfaces.ICoreMapLayer
	SetLayer(layerName string)
	DeleteLayer(layerName string)
	GetObject(x uint16, y uint16, layerName string) IDisplayObject
	SetObject(object IDisplayObject, x uint16, y uint16, layerName string)
	DeleteObject(x uint16, y uint16, layerName string)
	GetObjectsForArea(
		areaX uint16,
		areaY uint16,
		areaWidth uint16,
		areaHeight uint16,
	) map[string]map[uint32]interfaces.ICoreObject
}
