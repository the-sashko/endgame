package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"os"
)

var globalAppQuitHandlerInstance interfaces.ICoreHandler

type globalAppQuitHandler struct{}

func (handlerObject *globalAppQuitHandler) DoHandle(
	values interfaces.ICoreValuesObject,
) {
	logger.LogMessage("Quit", logger.TypeDefault)
	os.Exit(0)
}

func GetGlobalAppQuitHandler() interfaces.ICoreHandler {
	if globalAppQuitHandlerInstance == nil {
		globalAppQuitHandlerInstance = coreHandler.NewHandler(
			GlobalAppQuit,
			new(globalAppQuitHandler),
		)
	}

	return globalAppQuitHandlerInstance
}
