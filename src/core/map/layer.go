package _map

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"endgame/src/utils/map_index"
	"fmt"
)

type coreMapLayer struct {
	reference interfaces.ICoreReference
	name      string
	objects   map[uint32]interfaces.ICoreObject
}

func (mapLayer *coreMapLayer) GetReference() string {
	return mapLayer.reference.Get()
}

func (mapLayer *coreMapLayer) GetName() string {
	return mapLayer.name
}

func (mapLayer *coreMapLayer) GetObject(index uint32) interfaces.ICoreObject {
	if !mapLayer.HasObject(index) {
		return nil
	}

	return mapLayer.objects[index]
}

func (mapLayer *coreMapLayer) SetObject(
	object interfaces.ICoreObject,
	index uint32,
) {
	mapLayer.objects[index] = object
}

func (mapLayer *coreMapLayer) DeleteObject(index uint32) {
	if !mapLayer.HasObject(index) {
		return
	}

	delete(mapLayer.objects, index)
}

func (mapLayer *coreMapLayer) GetObjectsForArea(
	areaX uint16,
	areaY uint16,
	areaWidth uint16,
	areaHeight uint16,
) map[uint32]interfaces.ICoreObject {
	objects := make(map[uint32]interfaces.ICoreObject)

	for x := areaX; x < areaX+areaWidth; x++ {
		for y := areaY; y < areaY+areaHeight; y++ {
			index := map_index.GetIndex(x, y)

			if !mapLayer.HasObject(index) {
				continue
			}

			objects[index] = mapLayer.objects[index]
		}
	}

	return objects
}

func (mapLayer *coreMapLayer) HasObject(index uint32) bool {
	_, isExist := mapLayer.objects[index]

	return isExist
}

func newMapLayer(name string) interfaces.ICoreMapLayer {
	debugMessage := fmt.Sprintf("Created Core Map Layer %s", name)
	logger.LogDebug(debugMessage)

	return &coreMapLayer{
		reference: reference.NewReference(),
		name:      name,
		objects:   make(map[uint32]interfaces.ICoreObject),
	}
}
