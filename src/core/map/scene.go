package _map

import (
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"errors"
)

var coreSceneInstance interfaces.ICoreScene

type coreScene struct {
	x       uint16
	y       uint16
	width   uint16
	height  uint16
	objects map[string]map[uint32]interfaces.ICoreObject
}

func (scene *coreScene) GetX() uint16 {
	return scene.x
}

func (scene *coreScene) GetY() uint16 {
	return scene.y
}

func (scene *coreScene) GetWidth() uint16 {
	return scene.width
}

func (scene *coreScene) GetHeight() uint16 {
	return scene.height
}

func (scene *coreScene) GetObjects() map[string]map[uint32]interfaces.ICoreObject {
	return scene.objects
}

func (scene *coreScene) SetX(x uint16) {
	scene.x = x
}

func (scene *coreScene) SetY(y uint16) {
	scene.y = y
}

func (scene *coreScene) SetObjects(
	objects map[string]map[uint32]interfaces.ICoreObject,
) {
	scene.objects = objects
}

func (scene *coreScene) hasObject(layerName string, index uint32) bool {
	_, isExist := scene.objects[layerName]

	if !isExist {
		return false
	}

	_, isExist = scene.objects[layerName][index]

	if !isExist {
		return false
	}

	return true
}

func newCoreScene(width uint16, height uint16) interfaces.ICoreScene {
	logger.LogDebug("Created Core Map Scene")

	return &coreScene{
		x:       uint16(0),
		y:       uint16(0),
		width:   width,
		height:  height,
		objects: make(map[string]map[uint32]interfaces.ICoreObject),
	}
}

func InitCoreScene(width uint16, height uint16) {
	if coreSceneInstance != nil {
		err := errors.New("core scene instance is already exists")
		doError(err)
	}

	coreSceneInstance = newCoreScene(width, height)
}

func GetCoreScene() interfaces.ICoreScene {
	if coreSceneInstance == nil {
		err := errors.New("core scene not initialized")
		doError(err)
	}

	return coreSceneInstance
}
