package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalMouseButtonReleaseActionInstance interfaces.ICoreAction

func GetGlobalMouseButtonReleaseAction() interfaces.ICoreAction {
	if globalMouseButtonReleaseActionInstance == nil {
		globalMouseButtonReleaseActionInstance = coreAction.NewAction(
			GlobalMouseButtonRelease,
			nil,
		)
	}

	return globalMouseButtonReleaseActionInstance
}
