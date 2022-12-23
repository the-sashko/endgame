package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalAppStartEventInstance interfaces.ICoreEvent

func GetGlobalAppStartEvent() interfaces.ICoreEvent {
	if globalAppStartEventInstance == nil {
		globalAppStartEventInstance = coreEvents.NewEvent(GlobalAppStartType)
	}

	return globalAppStartEventInstance
}
