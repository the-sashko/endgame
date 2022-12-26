package display

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/object"
	"endgame/src/utils/logger"
	"fmt"
)

type displayObject struct {
	currentTextureName string
	textures           map[string]IBitmapImage
	coreObject         interfaces.ICoreObject
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

func (object *displayObject) SetTexture(textureName string) {
	object.currentTextureName = textureName
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

func (object *displayObject) DoHandleMove(deltaX uint16, deltaY uint16) {
	object.coreObject.DoHandleMove(deltaX, deltaY)
}

func NewDisplayObject(
	name string,
	shape map[uint32]bool,
	textureNames map[string]string,
	adapter interfaces.ICoreObjectAdapter,
) IDisplayObject {
	debugMessage := fmt.Sprintf("Created %s Display Object", name)
	logger.LogDebug(debugMessage)

	currentTextureName := "default"

	textures := make(map[string]IBitmapImage)

	for textureName, texturePath := range textureNames {
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
