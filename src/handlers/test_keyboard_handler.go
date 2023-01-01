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
	buttonName := values.Get("button_name").(string)

	x := global_state.GetGlobalState().Get("test_obj_x").(uint16)
	y := global_state.GetGlobalState().Get("test_obj_y").(uint16)

	m := global_state.GetGlobalState().GetMap("default")

	obj := m.GetObjectFromDefaultLayer(x, y)

	switch {
	case buttonName == "w" || buttonName == "up":
		m.DeleteObjectFromDefaultLayer(x, y)
		y = y - 5
		break
	case buttonName == "a" || buttonName == "left":
		m.DeleteObjectFromDefaultLayer(x, y)
		x = x - 5
		break
	case buttonName == "s" || buttonName == "down":
		m.DeleteObjectFromDefaultLayer(x, y)
		y = y + 5
		break
	case buttonName == "d" || buttonName == "right":
		m.DeleteObjectFromDefaultLayer(x, y)
		x = x + 5
		break
	default:
		return
	}

	global_state.GetGlobalState().Set("test_obj_x", x)
	global_state.GetGlobalState().Set("test_obj_y", y)
	m.SetObjectToDefaultLayer(obj, x, y)

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
