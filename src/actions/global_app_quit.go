package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var globalAppQuitActionInstance interfaces.ICoreAction

func GetGlobalAppQuitAction() interfaces.ICoreAction {
	if globalAppQuitActionInstance == nil {
		globalAppQuitActionInstance = coreAction.NewAction(GlobalAppQuit, nil)
	}

	return globalAppQuitActionInstance
}
