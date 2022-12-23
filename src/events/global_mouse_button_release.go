package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalMouseButtonReleaseEventInstance interfaces.ICoreEvent

func GetGlobalMouseButtonReleaseEvent() interfaces.ICoreEvent {
	if globalMouseButtonReleaseEventInstance == nil {
		globalMouseButtonReleaseEventInstance = coreEvents.NewEvent(
			GlobalMouseButtonReleaseType,
		)
	}

	return globalMouseButtonReleaseEventInstance
}
