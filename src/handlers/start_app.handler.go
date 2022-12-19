package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/events"
	"endgame/src/utils/logger"
)

var startAppHandlerInstance interfaces.ICoreHandler

type startAppHandler struct{}

func (handlerObject *startAppHandler) DoHandle(values map[string]interface{}) {
	logger.LogMessage("Started", logger.TypeDefault)
}

func createStartAppHandler() {
	startAppHandlerInstance := coreHandler.NewHandler("start_app", new(startAppHandler))
	startAppHandlerInstance.Subscribe(events.StartAppEventType)
}

func GetStartAppHandler() interfaces.ICoreHandler {
	if startAppHandlerInstance == nil {
		createStartAppHandler()
	}

	return startAppHandlerInstance
}
