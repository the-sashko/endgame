package display

import (
	"endgame/src/core/interfaces"
	_map "endgame/src/core/map"
	"endgame/src/utils/logger"
	"errors"
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

func (displayMapObject *displayMap) GetDefaultLayer() interfaces.ICoreMapLayer {
	return displayMapObject.coreMap.GetLayer(defaultMapLayer)
}

func (displayMapObject *displayMap) SetLayer(layerName string) {
	displayMapObject.coreMap.SetLayer(layerName)
}

func (displayMapObject *displayMap) DeleteLayer(layerName string) {
	if layerName == defaultMapLayer {
		err := errors.New("cannot delete default map layer")
		doError(err)
	}

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

func (displayMapObject *displayMap) GetObjectFromDefaultLayer(
	x uint16,
	y uint16,
) IDisplayObject {
	return displayMapObject.GetObject(x, y, defaultMapLayer)
}

func (displayMapObject *displayMap) SetObject(
	object IDisplayObject,
	x uint16,
	y uint16,
	layerName string,
) {
	object.SetMap(displayMapObject)
	object.SetMapLayerName(layerName)
	object.SetX(x)
	object.SetY(y)

	displayMapObject.coreMap.SetObject(object, x, y, layerName)
}

func (displayMapObject *displayMap) SetObjectToDefaultLayer(
	object IDisplayObject,
	x uint16,
	y uint16,
) {
	displayMapObject.SetObject(object, x, y, defaultMapLayer)
}

func (displayMapObject *displayMap) HasObject(
	x uint16,
	y uint16,
	layerName string,
) bool {
	return displayMapObject.coreMap.HasObject(x, y, layerName)
}

func (displayMapObject *displayMap) HasObjectOnDefaultLayer(
	x uint16,
	y uint16,
) bool {
	return displayMapObject.HasObject(x, y, defaultMapLayer)
}

func (displayMapObject *displayMap) DeleteObject(
	x uint16,
	y uint16,
	layerName string,
) {
	displayMapObject.coreMap.DeleteObject(x, y, layerName)
}

func (displayMapObject *displayMap) DeleteObjectFromDefaultLayer(
	x uint16,
	y uint16,
) {
	displayMapObject.DeleteObject(x, y, defaultMapLayer)
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

func NewDisplayMap(name string) IDisplayMap {
	debugMessage := fmt.Sprintf("Created Display Map %s", name)
	logger.LogDebug(debugMessage)

	displayMap := &displayMap{
		coreMap: _map.NewMap(name),
	}

	displayMap.SetLayer(defaultMapLayer)

	return displayMap
}
