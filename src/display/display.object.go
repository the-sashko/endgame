package display

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/object"
	"endgame/src/utils/logger"
	"errors"
	"fmt"
)

type displayObject struct {
	currentTextureName string
	textures           map[string]IBitmapImage
	coreObject         interfaces.ICoreObject
	x                  uint16
	y                  uint16
	mapObject          IDisplayMap
	mapLayerName       string
}

func (object *displayObject) GetReference() string {
	return object.coreObject.GetReference()
}

func (object *displayObject) GetName() string {
	return object.coreObject.GetName()
}

func (object *displayObject) GetShape() map[uint32]bool {
	return object.coreObject.GetShape()
}

func (object *displayObject) GetCurrentTexture() map[uint32]IColor {
	_, isExist := object.textures[object.currentTextureName]

	if !isExist {
		return nil
	}

	pixels := make(map[uint32]IColor)

	for index, colorObject := range object.textures[object.currentTextureName].GetPixels() {
		_, isExist = object.coreObject.GetShape()[index]

		if !isExist {
			continue
		}

		pixels[index] = colorObject
	}

	return pixels
}

func (object *displayObject) GetX() uint16 {
	return object.x
}

func (object *displayObject) GetY() uint16 {
	return object.y
}

func (object *displayObject) GetMap() IDisplayMap {
	if object.mapObject == nil {
		errorMessage := fmt.Sprintf(
			"Display Object %s Has No Map",
			object.GetName(),
		)

		doError(errors.New(errorMessage))
	}

	if object.mapObject.GetObject(
		object.x,
		object.y,
		object.mapLayerName,
	) == nil {
		errorMessage := fmt.Sprintf(
			"Not Found Object On Map. Map: %s Layer: %s X: %d Y: %d",
			object.mapObject.GetName(),
			object.mapLayerName,
			object.x,
			object.y,
		)

		doError(errors.New(errorMessage))
	}

	if object.mapObject.GetObject(
		object.x,
		object.y,
		object.mapLayerName,
	).GetReference() != object.GetReference() {
		errorMessage := fmt.Sprintf(
			"Map objects mismatch. Map: %s Layer: %s X: %d Y: %d",
			object.mapObject.GetName(),
			object.mapLayerName,
			object.x,
			object.y,
		)

		doError(errors.New(errorMessage))
	}

	return object.mapObject
}

func (object *displayObject) GetMapLayerName() string {
	return object.mapLayerName
}

func (object *displayObject) SetTexture(textureName string) {
	object.currentTextureName = textureName
}

func (object *displayObject) SetX(x uint16) {
	object.x = x
}

func (object *displayObject) SetY(y uint16) {
	object.y = y
}

func (object *displayObject) SetMap(mapObject IDisplayMap) {
	object.mapObject = mapObject
}

func (object *displayObject) SetMapLayerName(mapLayerName string) {
	object.mapLayerName = mapLayerName
}

func (object *displayObject) DoMove(deltaX int32, deltaY int32) {
	logMessage := fmt.Sprintf(
		"Moving %s Display Object X: %d Y: %d",
		object.GetName(),
		deltaX,
		deltaY,
	)

	logger.LogDebug(logMessage)

	mapObject := object.GetMap()

	fromX := object.x
	fromY := object.y

	toX := int32(fromX) + deltaX
	toY := int32(fromY) + deltaY

	if toX < 0 {
		toX = 0
	}

	if toY < 0 {
		toY = 0
	}

	object.SetX(uint16(toX))
	object.SetY(uint16(toY))

	mapObject.DeleteObject(fromX, fromY, object.mapLayerName)
	mapObject.SetObject(object, uint16(toX), uint16(toY), object.mapLayerName)

	object.DoHandleMove(deltaX, deltaY)
}

func (object *displayObject) Destroy() {
	debugMessage := fmt.Sprintf("Destroyed %s Display Object", object.GetName())
	logger.LogDebug(debugMessage)

	object.GetMap().DeleteObject(
		object.GetX(),
		object.GetY(),
		object.GetMapLayerName(),
	)

	object.coreObject.Destroy()
	object.coreObject = nil
}

func (object *displayObject) DoHandleCollision(
	collidedObject interfaces.ICoreObject,
) {
	object.coreObject.DoHandleCollision(collidedObject)
}

func (object *displayObject) DoHandleMouseHover() {
	object.coreObject.DoHandleMouseHover()
}

func (object *displayObject) DoHandleMouseButtonPress(
	mouseButton interfaces.ICoreButton,
) {
	object.coreObject.DoHandleMouseButtonPress(mouseButton)
}

func (object *displayObject) DoHandleMouseButtonRelease(
	mouseButton interfaces.ICoreButton,
) {
	object.coreObject.DoHandleMouseButtonRelease(mouseButton)
}

func (object *displayObject) DoHandleMouseButtonClick(
	mouseButton interfaces.ICoreButton,
) {
	object.coreObject.DoHandleMouseButtonClick(mouseButton)
}

func (object *displayObject) DoHandleMove(deltaX int32, deltaY int32) {
	object.coreObject.DoHandleMove(deltaX, deltaY)
}

func NewDisplayObject(
	name string,
	shape map[uint32]bool,
	texturePaths map[string]string,
	adapter interfaces.ICoreObjectAdapter,
) IDisplayObject {
	debugMessage := fmt.Sprintf("Created %s Display Object", name)
	logger.LogDebug(debugMessage)

	currentTextureName := "default"

	textures := make(map[string]IBitmapImage)

	for textureName, texturePath := range texturePaths {
		textures[textureName] = newBitmapImage(textureName, texturePath)
	}

	return &displayObject{
		textures:           textures,
		currentTextureName: currentTextureName,
		coreObject: object.NewCoreObject(
			name,
			shape,
			adapter,
		),
	}
}
