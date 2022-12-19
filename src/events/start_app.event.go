package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var startAppEventInstance interfaces.ICoreEvent

func GetStartAppEvent() interfaces.ICoreEvent {
	if startAppEventInstance == nil {
		startAppEventInstance = coreEvents.NewEvent(StartAppEventType)
	}

	return startAppEventInstance
}
