package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalMouseScrollActionInstance interfaces.ICoreAction

func GetGlobalMouseScrollAction() interfaces.ICoreAction {
	if globalMouseScrollActionInstance == nil {
		globalMouseScrollActionInstance = coreAction.NewAction(
			GlobalMouseScroll,
			nil,
		)
	}

	return globalMouseScrollActionInstance
}
