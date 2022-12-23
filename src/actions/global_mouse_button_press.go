package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalMouseButtonPressActionInstance interfaces.ICoreAction

func GetGlobalMouseButtonPressAction() interfaces.ICoreAction {
	if globalMouseButtonPressActionInstance == nil {
		globalMouseButtonPressActionInstance = coreAction.NewAction(
			GlobalMouseButtonPress,
			nil,
		)
	}

	return globalMouseButtonPressActionInstance
}
