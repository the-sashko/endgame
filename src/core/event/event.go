package event

import (
	"endgame/src/core/dispatcher"
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"fmt"
)

type coreEvent struct {
	reference interfaces.ICoreReference
	typeName  string
}

func (event *coreEvent) GetReference() string {
	return event.reference.Get()
}

func (event *coreEvent) GetTypeName() string {
	return event.typeName
}

func (event *coreEvent) Fire(values interfaces.ICoreValuesObject) {
	debugMessage := fmt.Sprintf("Fired Event With %s Event Type", event.GetTypeName())
	logger.LogDebug(debugMessage)

	dispatcher.GetDispatcher().Dispatch(event.GetTypeName(), values)
}

func NewEvent(typeName string) interfaces.ICoreEvent {
	debugMessage := fmt.Sprintf("Created Event With %s Event Type", typeName)
	logger.LogDebug(debugMessage)

	return &coreEvent{
		reference: reference.NewReference(),
		typeName:  typeName,
	}
}
