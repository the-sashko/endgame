package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalKeyboardKeyReleaseActionInstance interfaces.ICoreAction

func GetGlobalKeyboardKeyReleaseAction() interfaces.ICoreAction {
	if globalKeyboardKeyReleaseActionInstance == nil {
		globalKeyboardKeyReleaseActionInstance = coreAction.NewAction(
			GlobalKeyboardKeyRelease,
			nil,
		)
	}

	return globalKeyboardKeyReleaseActionInstance
}
