package action

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"fmt"
)

type coreAction struct {
	reference interfaces.ICoreReference
	name      string
	events    map[string]interfaces.ICoreEvent
}

func (action *coreAction) GetReference() string {
	return action.reference.Get()
}

func (action *coreAction) GetName() string {
	return action.name
}

func (action *coreAction) Fire(values interfaces.ICoreValuesObject) {
	debugMessage := fmt.Sprintf("Fired %s Action", action.GetName())
	logger.LogDebug(debugMessage)

	for _, event := range action.events {
		event.Fire(values)
	}
}

func (action *coreAction) GetEvents() map[string]interfaces.ICoreEvent {
	return action.events
}

func (action *coreAction) HasEvent(eventType string) bool {
	_, isExist := action.events[eventType]

	return isExist
}

func (action *coreAction) GetEvent(eventType string) interfaces.ICoreEvent {
	if !action.HasEvent(eventType) {
		return nil
	}

	return action.events[eventType]
}

func (action *coreAction) AddEvent(event interfaces.ICoreEvent) {
	debugMessage := fmt.Sprintf("Added %s Event Type To %s Action", event.GetTypeName(), action.GetName())
	logger.LogDebug(debugMessage)

	action.events[event.GetTypeName()] = event
}

func (action *coreAction) RemoveEvent(eventType string) {
	debugMessage := fmt.Sprintf("Removed %s Event Type From %s Action", eventType, action.GetName())
	logger.LogDebug(debugMessage)

	delete(action.events, eventType)
}

func NewAction(name string, events []interfaces.ICoreEvent) interfaces.ICoreAction {
	debugMessage := fmt.Sprintf("Created %s Action", name)
	logger.LogDebug(debugMessage)

	eventsMap := make(map[string]interfaces.ICoreEvent)

	for _, event := range events {
		eventsMap[event.GetTypeName()] = event
	}

	return &coreAction{
		reference: reference.NewReference(),
		name:      name,
		events:    eventsMap,
	}
}
