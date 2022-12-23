package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalKeyboardKeyReleaseEventInstance interfaces.ICoreEvent

func GetGlobalKeyboardKeyReleaseEvent() interfaces.ICoreEvent {
	if globalKeyboardKeyReleaseEventInstance == nil {
		globalKeyboardKeyReleaseEventInstance = coreEvents.NewEvent(
			GlobalKeyboardKeyReleaseType,
		)
	}

	return globalKeyboardKeyReleaseEventInstance
}
