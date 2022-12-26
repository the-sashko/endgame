package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
)

var globalAppStartHandlerInstance interfaces.ICoreHandler

type globalAppStartHandler struct{}

func (handlerObject *globalAppStartHandler) DoHandle(
	values interfaces.ICoreValuesObject,
) {
	logger.LogMessage("Started", logger.TypeDefault)
}

func GetGlobalAppStartHandler() interfaces.ICoreHandler {
	if globalAppStartHandlerInstance == nil {
		globalAppStartHandlerInstance = coreHandler.NewHandler(
			GlobalAppStart,
			new(globalAppStartHandler),
		)
	}

	return globalAppStartHandlerInstance
}
