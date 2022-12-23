package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalMouseButtonClickActionInstance interfaces.ICoreAction

func GetGlobalMouseButtonClickAction() interfaces.ICoreAction {
	if globalMouseButtonClickActionInstance == nil {
		globalMouseButtonClickActionInstance = coreAction.NewAction(
			GlobalMouseButtonClick,
			nil,
		)
	}

	return globalMouseButtonClickActionInstance
}
