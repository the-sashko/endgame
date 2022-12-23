package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalKeyboardKeyPressActionInstance interfaces.ICoreAction

func GetGlobalKeyboardKeyPressAction() interfaces.ICoreAction {
	if globalKeyboardKeyPressActionInstance == nil {
		globalKeyboardKeyPressActionInstance = coreAction.NewAction(
			GlobalKeyboardKeyPress,
			nil,
		)
	}

	return globalKeyboardKeyPressActionInstance
}
