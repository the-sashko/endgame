package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalMouseButtonPressEventInstance interfaces.ICoreEvent

func GetGlobalMouseButtonPressEvent() interfaces.ICoreEvent {
	if globalMouseButtonPressEventInstance == nil {
		globalMouseButtonPressEventInstance = coreEvents.NewEvent(
			GlobalMouseButtonPressType,
		)
	}

	return globalMouseButtonPressEventInstance
}
