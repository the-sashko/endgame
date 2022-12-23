package events

import (
	coreEvents "endgame/src/core/event"
	"endgame/src/core/interfaces"
)

var globalMouseButtonClickEventInstance interfaces.ICoreEvent

func GetGlobalMouseButtonClickEvent() interfaces.ICoreEvent {
	if globalMouseButtonClickEventInstance == nil {
		globalMouseButtonClickEventInstance = coreEvents.NewEvent(
			GlobalMouseButtonClickType,
		)
	}

	return globalMouseButtonClickEventInstance
}
