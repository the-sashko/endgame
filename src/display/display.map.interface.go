package display

import (
	"endgame/src/core/interfaces"
)

type IDisplayMap interface {
	GetReference() string
	GetName() string
	GetLayers() map[string]interfaces.ICoreMapLayer
	GetLayer(layerName string) interfaces.ICoreMapLayer
	GetDefaultLayer() interfaces.ICoreMapLayer
	SetLayer(layerName string)
	DeleteLayer(layerName string)
	GetObject(x uint16, y uint16, layerName string) IDisplayObject
	GetObjectFromDefaultLayer(x uint16, y uint16) IDisplayObject
	SetObject(object IDisplayObject, x uint16, y uint16, layerName string)
	SetObjectToDefaultLayer(object IDisplayObject, x uint16, y uint16)
	DeleteObject(x uint16, y uint16, layerName string)
	DeleteObjectFromDefaultLayer(x uint16, y uint16)
	GetObjectsForArea(
		areaX uint16,
		areaY uint16,
		areaWidth uint16,
		areaHeight uint16,
	) map[string]map[uint32]interfaces.ICoreObject
}
