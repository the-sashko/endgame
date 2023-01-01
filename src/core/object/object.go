package object

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"fmt"
)

type coreObject struct {
	reference interfaces.ICoreReference
	name      string
	shape     map[uint32]bool
	adapter   interfaces.ICoreObjectAdapter
}

func (object *coreObject) GetReference() string {
	return object.reference.Get()
}

func (object *coreObject) GetName() string {
	return object.name
}

func (object *coreObject) GetShape() map[uint32]bool {
	return object.shape
}

func (object *coreObject) Destroy() {
	debugMessage := fmt.Sprintf("Destroyed %s Object", object.GetName())
	logger.LogDebug(debugMessage)

	object.adapter.Destroy()
	object.adapter = nil
}

func (object *coreObject) DoHandleCollision(
	collidedObject interfaces.ICoreObject,
) {
	debugMessage := fmt.Sprintf(
		"Colided %s Object With %s Object",
		object.GetName(),
		collidedObject.GetName(),
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleCollision(collidedObject)
}

func (object *coreObject) DoHandleMouseHover() {
	debugMessage := fmt.Sprintf(
		"Mouse Hover On %s Object",
		object.GetName(),
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleMouseHover()
}

func (object *coreObject) DoHandleMouseButtonPress(
	mouseButton interfaces.ICoreButton,
) {
	debugMessage := fmt.Sprintf(
		"Mouse Button Press On %s Object",
		object.GetName(),
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleMouseButtonPress(mouseButton)
}

func (object *coreObject) DoHandleMouseButtonRelease(
	mouseButton interfaces.ICoreButton,
) {
	debugMessage := fmt.Sprintf(
		"Mouse Button Release On %s Object",
		object.GetName(),
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleMouseButtonRelease(mouseButton)
}

func (object *coreObject) DoHandleMouseButtonClick(
	mouseButton interfaces.ICoreButton,
) {
	debugMessage := fmt.Sprintf(
		"Mouse Button Click On %s Object",
		object.GetName(),
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleMouseButtonClick(mouseButton)
}

func (object *coreObject) DoHandleMove(deltaX int32, deltaY int32) {
	debugMessage := fmt.Sprintf(
		"Move %s Object: %d X %d Y",
		object.GetName(),
		deltaX,
		deltaY,
	)

	logger.LogDebug(debugMessage)

	object.adapter.DoHandleMove(deltaX, deltaY)
}

func NewCoreObject(
	name string,
	shape map[uint32]bool,
	adapter interfaces.ICoreObjectAdapter,
) interfaces.ICoreObject {
	debugMessage := fmt.Sprintf("Created %s Object", name)
	logger.LogDebug(debugMessage)

	return &coreObject{
		reference: reference.NewReference(),
		name:      name,
		shape:     shape,
		adapter:   adapter,
	}
}
