package _map

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"endgame/src/utils/map_index"
	"fmt"
)

type coreMap struct {
	reference interfaces.ICoreReference
	name      string
	layers    map[string]interfaces.ICoreMapLayer
}

func (mapObject *coreMap) GetReference() string {
	return mapObject.reference.Get()
}

func (mapObject *coreMap) GetName() string {
	return mapObject.name
}

func (mapObject *coreMap) GetLayers() map[string]interfaces.ICoreMapLayer {
	return mapObject.layers
}

func (mapObject *coreMap) GetLayer(layerName string) interfaces.ICoreMapLayer {
	if !mapObject.hasLayer(layerName) {
		return nil
	}

	return mapObject.layers[layerName]
}

func (mapObject *coreMap) SetLayer(layerName string) {
	if mapObject.hasLayer(layerName) {
		return
	}

	mapObject.layers[layerName] = newMapLayer(layerName)
}

func (mapObject *coreMap) DeleteLayer(layerName string) {
	if !mapObject.hasLayer(layerName) {
		return
	}

	delete(mapObject.layers, layerName)
}

func (mapObject *coreMap) GetObject(
	x uint16,
	y uint16,
	layerName string,
) interfaces.ICoreObject {
	index := map_index.GetIndex(x, y)

	if !mapObject.hasLayer(layerName) {
		return nil
	}

	return mapObject.layers[layerName].GetObject(index)
}

func (mapObject *coreMap) SetObject(
	object interfaces.ICoreObject,
	x uint16,
	y uint16,
	layerName string,
) {
	if !mapObject.hasLayer(layerName) {
		mapObject.SetLayer(layerName)
	}

	index := map_index.GetIndex(x, y)

	mapObject.layers[layerName].SetObject(object, index)
}

func (mapObject *coreMap) DeleteObject(x uint16, y uint16, layerName string) {
	if !mapObject.hasLayer(layerName) {
		return
	}

	index := map_index.GetIndex(x, y)

	mapObject.layers[layerName].DeleteObject(index)
}

func (mapObject *coreMap) GetObjectsForArea(
	areaX uint16,
	areaY uint16,
	areaWidth uint16,
	areaHeight uint16,
) map[string]map[uint32]interfaces.ICoreObject {
	objects := make(map[string]map[uint32]interfaces.ICoreObject)

	for layerName, layer := range mapObject.layers {
		objects[layerName] = layer.GetObjectsForArea(
			areaX,
			areaY,
			areaWidth,
			areaHeight,
		)
	}

	return objects
}

func (mapObject *coreMap) hasLayer(layerName string) bool {
	_, isExist := mapObject.layers[layerName]

	return isExist
}

func NewMap(name string) interfaces.ICoreMap {
	debugMessage := fmt.Sprintf("Created Core Map %s", name)
	logger.LogDebug(debugMessage)

	return &coreMap{
		reference: reference.NewReference(),
		name:      name,
		layers:    make(map[string]interfaces.ICoreMapLayer),
	}
}

func doError(errorEntity error) {
	logger.LogError("map", errorEntity, true)
}
