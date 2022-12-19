package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var quitAppEventInstance interfaces.ICoreEvent

func GetQuitAppEvent() interfaces.ICoreEvent {
	if quitAppEventInstance == nil {
		quitAppEventInstance = coreEvents.NewEvent(QuitAppEventType)
	}

	return quitAppEventInstance
}
