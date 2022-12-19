package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/events"
	"endgame/src/utils/logger"
	"os"
)

var quitAppHandlerInstance interfaces.ICoreHandler

type quitAppHandler struct{}

func (handlerObject *quitAppHandler) DoHandle(values map[string]interface{}) {
	logger.LogMessage("Quit", logger.TypeDefault)
	os.Exit(0)
}

func createQuitAppHandler() {
	quitAppHandlerInstance := coreHandler.NewHandler("quit_app", new(quitAppHandler))
	quitAppHandlerInstance.Subscribe(events.QuitAppEventType)
}

func GetQuitAppHandler() interfaces.ICoreHandler {
	if quitAppHandlerInstance == nil {
		createQuitAppHandler()
	}

	return quitAppHandlerInstance
}
