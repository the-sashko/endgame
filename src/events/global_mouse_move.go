package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalMouseMoveEventInstance interfaces.ICoreEvent

func GetGlobalMouseMoveEvent() interfaces.ICoreEvent {
	if globalMouseMoveEventInstance == nil {
		globalMouseMoveEventInstance = coreEvents.NewEvent(GlobalMouseMoveType)
	}

	return globalMouseMoveEventInstance
}
