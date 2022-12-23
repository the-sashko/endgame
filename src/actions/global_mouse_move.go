package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalMouseMoveActionInstance interfaces.ICoreAction

func GetGlobalMouseMoveAction() interfaces.ICoreAction {
	if globalMouseMoveActionInstance == nil {
		globalMouseMoveActionInstance = coreAction.NewAction(
			GlobalMouseMove,
			nil,
		)
	}

	return globalMouseMoveActionInstance
}
