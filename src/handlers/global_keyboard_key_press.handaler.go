package handlers

import (
	"endgame/src/actions"
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
)

var globalKeyboardKeyPressHandlerInstance interfaces.ICoreHandler

type globalKeyboardKeyPressHandler struct{}

func (handlerObject *globalKeyboardKeyPressHandler) DoHandle(
	values interfaces.ICoreValuesObject,
) {
	if values.Get("button_name").(string) != "escape" {
		return
	}

	actions.GetGlobalAppQuitAction().Fire(nil)
}

func GetGlobalKeyboardKeyPressHandler() interfaces.ICoreHandler {
	if globalKeyboardKeyPressHandlerInstance == nil {
		globalKeyboardKeyPressHandlerInstance = coreHandler.NewHandler(
			GlobalKeyboardKeyPress,
			new(globalKeyboardKeyPressHandler),
		)
	}

	return globalKeyboardKeyPressHandlerInstance
}
