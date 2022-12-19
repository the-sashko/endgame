package handler

import (
	"endgame/src/core/dispatcher"
	"endgame/src/core/interfaces"
	"endgame/src/core/reference"
	"endgame/src/utils/logger"
	"fmt"
)

type coreHandler struct {
	reference interfaces.ICoreReference
	name      string
	adapter   interfaces.ICoreHandlerAdapter
}

func (handler *coreHandler) GetReference() string {
	return handler.reference.Get()
}

func (handler *coreHandler) GetName() string {
	return handler.name
}

func (handler *coreHandler) DoHandle(values map[string]interface{}) {
	debugMessage := fmt.Sprintf("%s Handler Handled Action Event Data", handler.GetName())
	logger.LogDebug(debugMessage)

	handler.adapter.DoHandle(values)
}

func (handler *coreHandler) Subscribe(eventType string) {
	dispatcher.GetDispatcher().Subscribe(eventType, handler)
}

func (handler *coreHandler) Unsubscribe(eventType string) {
	dispatcher.GetDispatcher().Unsubscribe(eventType, handler)
}

func (handler *coreHandler) GetSubscriptions() []string {
	return dispatcher.GetDispatcher().GetSubscriptions(handler.GetReference())
}

func NewHandler(name string, adapter interfaces.ICoreHandlerAdapter) interfaces.ICoreHandler {
	debugMessage := fmt.Sprintf("Created %s Handler", name)
	logger.LogDebug(debugMessage)

	return &coreHandler{
		reference: reference.NewReference(),
		name:      name,
		adapter:   adapter,
	}
}
