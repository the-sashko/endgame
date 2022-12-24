package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalMouseScrollEventInstance interfaces.ICoreEvent

func GetGlobalMouseScrollEvent() interfaces.ICoreEvent {
	if globalMouseScrollEventInstance == nil {
		globalMouseScrollEventInstance = coreEvents.NewEvent(
			GlobalMouseScrollType,
		)
	}

	return globalMouseScrollEventInstance
}
