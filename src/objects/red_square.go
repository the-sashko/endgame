package objects

import (
	"endgame/src/core/interfaces"
	"endgame/src/display"
	"endgame/src/utils/logger"
	"endgame/src/utils/map_index"
)

const redSquareDefaultTexture = "../res/red_square.bmp"

type redSquareObjectAdapter struct{}

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
	deltaX uint16,
	deltaY uint16,
) {
	//todo
}

func NewRedSquareObject(
	name string,
) display.IDisplayObject {
	logger.LogDebug("Created Red Square Object Adapter")

	shape := make(map[uint32]bool)

	for x := uint16(0); x < 16; x++ {
		for y := uint16(0); y < 16; y++ {
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
