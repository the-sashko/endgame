package actions

import (
	coreAction "endgame/src/core/action"
	"endgame/src/core/interfaces"
)

var quitAppActionInstance interfaces.ICoreAction

const quitAppActionName = "g_action_quit_app"

func GetQuitAppAction() interfaces.ICoreAction {
	if quitAppActionInstance == nil {
		quitAppActionInstance = coreAction.NewAction(quitAppActionName, nil)
	}

	return quitAppActionInstance
}
