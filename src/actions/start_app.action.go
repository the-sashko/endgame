package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var startAppActionInstance interfaces.ICoreAction

const startAppActionName = "g_action_start_app"

func GetStartAppAction() interfaces.ICoreAction {
	if startAppActionInstance == nil {
		startAppActionInstance = coreAction.NewAction(startAppActionName, nil)
	}

	return startAppActionInstance
}
