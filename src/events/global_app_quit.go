package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalAppQuitEventInstance interfaces.ICoreEvent

func GetGlobalAppQuitEvent() interfaces.ICoreEvent {
	if globalAppQuitEventInstance == nil {
		globalAppQuitEventInstance = coreEvents.NewEvent(GlobalAppQuitType)
	}

	return globalAppQuitEventInstance
}
