package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalKeyboardKeyPressEventInstance interfaces.ICoreEvent

func GetGlobalKeyboardKeyPressEvent() interfaces.ICoreEvent {
	if globalKeyboardKeyPressEventInstance == nil {
		globalKeyboardKeyPressEventInstance = coreEvents.NewEvent(
			GlobalKeyboardKeyPressType,
		)
	}

	return globalKeyboardKeyPressEventInstance
}
