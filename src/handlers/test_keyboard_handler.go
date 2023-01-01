package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/display"
	"endgame/src/global_state"
)

var testKeyboardHandlerInstance interfaces.ICoreHandler

type testKeyboardHandler struct{}

func (handlerObject *testKeyboardHandler) DoHandle(
	values interfaces.ICoreValuesObject,
) {
	deltaX := int32(0)
	deltaY := int32(0)

	buttonName := values.Get("button_name").(string)

	switch {
	case buttonName == "w" || buttonName == "up":
		deltaY = deltaY - 5
		break
	case buttonName == "a" || buttonName == "left":
		deltaX = deltaX - 5
		break
	case buttonName == "s" || buttonName == "down":
		deltaY = deltaY + 5
		break
	case buttonName == "d" || buttonName == "right":
		deltaX = deltaX + 5
		break
	default:
		return
	}

	global_state.
		GetGlobalState().
		GetObject("test").(display.IDisplayObject).
		DoMove(deltaX, deltaY)

	display.GetInstance().ClearAll()
}

func GetTestKeyboardHandler() interfaces.ICoreHandler {
	if testKeyboardHandlerInstance == nil {
		testKeyboardHandlerInstance = coreHandler.NewHandler(
			"test_keyboard_handler",
			new(testKeyboardHandler),
		)
	}

	return testKeyboardHandlerInstance
}
