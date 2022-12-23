package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalAppStartActionInstance interfaces.ICoreAction

func GetGlobalAppStartAction() interfaces.ICoreAction {
	if globalAppStartActionInstance == nil {
		globalAppStartActionInstance = coreAction.NewAction(GlobalAppStart, nil)
	}

	return globalAppStartActionInstance
}
