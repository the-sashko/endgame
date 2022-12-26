package display

import (
	"endgame/src/core/interfaces"
	_map "endgame/src/core/map"
	"endgame/src/utils/logger"
	"endgame/src/utils/map_index"
	"errors"
)

var displaySceneInstance IDisplayScene

type displayScene struct {
	displayBufferLayers map[string]byte
	pixels              map[byte]map[uint32]IColor
	coreScene           interfaces.ICoreScene
}

func (scene *displayScene) GetX() uint16 {
	return scene.coreScene.GetX()
}

func (scene *displayScene) GetY() uint16 {
	return scene.coreScene.GetY()
}

func (scene *displayScene) GetWidth() uint16 {
	return scene.coreScene.GetWidth()
}

func (scene *displayScene) GetHeight() uint16 {
	return scene.coreScene.GetHeight()
}

func (scene *displayScene) SetX(x uint16) {
	scene.coreScene.SetX(x)
}

func (scene *displayScene) SetY(y uint16) {
	scene.coreScene.SetY(y)
}

func (scene *displayScene) SetDisplayBufferLayer(
	mapLayerName string,
	bufferLayer byte,
) {
	scene.displayBufferLayers[mapLayerName] = bufferLayer
}

func (scene *displayScene) GetPixels() map[byte]map[uint32]IColor {
	return scene.pixels
}

func (scene *displayScene) SetObjects(
	objects map[string]map[uint32]interfaces.ICoreObject,
) {
	scene.coreScene.SetObjects(objects)
}

func (scene *displayScene) DeleteDisplayBufferLayer(mapLayerName string) {
	if mapLayerName == defaultMapLayer {
		err := errors.New("cannot delete default map layer")
		doError(err)
	}

	delete(scene.displayBufferLayers, mapLayerName)
}

func (scene *displayScene) DoFillPixels() {
	coreObjects := scene.getObjects()

	for mapLayerName, displayBufferLayer := range scene.displayBufferLayers {
		for objectIndex, coreObject := range coreObjects[mapLayerName] {
			if coreObject == nil {
				continue
			}

			pixels := coreObject.(IDisplayObject).GetCurrentTexture()

			for pixelIndex, pixel := range pixels {
				objectX, objectY := map_index.RetrieveXY(objectIndex)
				pixelX, pixelY := map_index.RetrieveXY(pixelIndex)

				pixelX += objectX
				pixelY += objectY

				if pixelX >= scene.GetWidth() || pixelY >= scene.GetHeight() {
					continue
				}

				scenePixelIndex := map_index.GetIndex(pixelX, pixelY)

				scene.pixels[displayBufferLayer][scenePixelIndex] = pixel
			}
		}
	}
}

func (scene *displayScene) getObjects() map[string]map[uint32]interfaces.ICoreObject {
	displayObjects := scene.coreScene.GetObjects()

	return displayObjects
}

func initDisplayScene(width uint16, height uint16) {
	if displaySceneInstance != nil {
		err := errors.New("display scene instance is already exists")
		doError(err)
	}

	_map.InitCoreScene(width, height)

	displaySceneInstance = newDisplayScene()

	displaySceneInstance.SetDisplayBufferLayer(
		defaultMapLayer,
		defaultBufferLayer,
	)
}

func newDisplayScene() IDisplayScene {
	logger.LogDebug("Created Display Scene")

	return &displayScene{
		displayBufferLayers: make(map[string]byte),
		pixels:              make(map[byte]map[uint32]IColor),
		coreScene:           _map.GetCoreScene(),
	}
}

func GetDisplayScene() IDisplayScene {
	if displaySceneInstance == nil {
		err := errors.New("display scene not initialized")
		doError(err)
	}

	return displaySceneInstance
}
