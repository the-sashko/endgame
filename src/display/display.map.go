package display

import (
	"endgame/src/core/interfaces"
	_map "endgame/src/core/map"
	"endgame/src/utils/logger"
	"fmt"
)

type displayMap struct {
	coreMap interfaces.ICoreMap
}

func (displayMapObject *displayMap) GetReference() string {
	return displayMapObject.coreMap.GetReference()
}

func (displayMapObject *displayMap) GetName() string {
	return displayMapObject.coreMap.GetName()
}

func (displayMapObject *displayMap) GetLayers() map[string]interfaces.ICoreMapLayer {
	return displayMapObject.coreMap.GetLayers()
}

func (displayMapObject *displayMap) GetLayer(
	layerName string,
) interfaces.ICoreMapLayer {
	return displayMapObject.coreMap.GetLayer(layerName)
}

func (displayMapObject *displayMap) SetLayer(layerName string) {
	displayMapObject.coreMap.SetLayer(layerName)
}

func (displayMapObject *displayMap) DeleteLayer(layerName string) {
	displayMapObject.coreMap.DeleteLayer(layerName)
}

func (displayMapObject *displayMap) GetObject(
	x uint16,
	y uint16,
	layerName string,
) IDisplayObject {
	object := displayMapObject.coreMap.GetObject(x, y, layerName)

	if object == nil {
		return nil
	}

	return object.(IDisplayObject)
}

func (displayMapObject *displayMap) SetObject(
	object IDisplayObject,
	x uint16,
	y uint16,
	layerName string,
) {
	displayMapObject.coreMap.SetObject(object, x, y, layerName)
}

func (displayMapObject *displayMap) DeleteObject(
	x uint16,
	y uint16,
	layerName string,
) {
	displayMapObject.coreMap.DeleteObject(x, y, layerName)
}

func (displayMapObject *displayMap) GetObjectsForArea(
	areaX uint16,
	areaY uint16,
	areaWidth uint16,
	areaHeight uint16,
) map[string]map[uint32]interfaces.ICoreObject {
	return displayMapObject.coreMap.GetObjectsForArea(
		areaX,
		areaY,
		areaWidth,
		areaHeight,
	)
}

func newDisplayMap(name string) IDisplayMap {
	debugMessage := fmt.Sprintf("Created Display Map %s", name)
	logger.LogDebug(debugMessage)

	return &displayMap{
		coreMap: _map.NewMap(name),
	}
}
