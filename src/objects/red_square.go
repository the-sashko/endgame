package objects

import (
	"endgame/src/core/interfaces"
	"endgame/src/display"
	"endgame/src/utils/logger"
	"endgame/src/utils/map_index"
)

const redSquareDefaultTexture = "../res/red_square.bmp"

type redSquareObjectAdapter struct{}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) Destroy() {
	logger.LogDebug("Destroyed Red Square Object Adapter")

	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleCollision(
	collidedObject interfaces.ICoreObject,
) {
	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleMouseHover() {
	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleMouseButtonPress(
	mouseButton interfaces.ICoreButton,
) {
	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleMouseButtonRelease(
	mouseButton interfaces.ICoreButton,
) {
	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleMouseButtonClick(
	mouseButton interfaces.ICoreButton,
) {
	//todo
}

func (redSquareObjectAdapterInstance *redSquareObjectAdapter) DoHandleMove(
	deltaX int32,
	deltaY int32,
) {
	//todo
}

func NewRedSquareObject(
	name string,
) display.IDisplayObject {
	logger.LogDebug("Created Red Square Object Adapter")

	shape := make(map[uint32]bool)

	for x := uint16(0); x < 164; x++ {
		for y := uint16(0); y < 64; y++ {
			shape[map_index.GetIndex(x, y)] = true
		}
	}

	texturePaths := make(map[string]string)
	texturePaths["default"] = redSquareDefaultTexture

	redSquareObjectAdapterInstance := &redSquareObjectAdapter{}

	return display.NewDisplayObject(
		name,
		shape,
		texturePaths,
		redSquareObjectAdapterInstance,
	)
}
