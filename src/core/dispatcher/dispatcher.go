package dispatcher

import (
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"fmt"
)

var dispatcherInstance interfaces.ICoreDispatcher

type coreDispatcher struct {
	subscriptions map[string]map[string]interfaces.ICoreHandler
}

func (dispatcher *coreDispatcher) GetSubscriptions(handlerReference string) []string {
	var subscriptions []string

	for eventType, handlers := range dispatcher.subscriptions {
		for reference := range handlers {
			if reference != handlerReference {
				continue
			}

			subscriptions = append(subscriptions, eventType)
			break
		}
	}

	return subscriptions
}

func (dispatcher *coreDispatcher) Subscribe(eventType string, handler interfaces.ICoreHandler) {
	debugMessage := fmt.Sprintf("Handler %s Subscribed To %s Event Type", handler.GetName(), eventType)
	logger.LogDebug(debugMessage)

	_, isExist := dispatcher.subscriptions[eventType]

	if !isExist {
		dispatcher.subscriptions[eventType] = make(map[string]interfaces.
			ICoreHandler)
	}

	dispatcher.subscriptions[eventType][handler.GetReference()] = handler
}

func (dispatcher *coreDispatcher) Unsubscribe(eventType string, handler interfaces.ICoreHandler) {
	debugMessage := fmt.Sprintf("Handler %s Unsubscribed From %s Event Type", handler.GetName(), eventType)
	logger.LogDebug(debugMessage)

	_, isExist := dispatcher.subscriptions[eventType]

	if !isExist {
		return
	}

	_, isExist = dispatcher.subscriptions[eventType][handler.GetReference()]

	if !isExist {
		return
	}

	delete(dispatcher.subscriptions[eventType], handler.GetReference())
}

func (dispatcher *coreDispatcher) Dispatch(eventType string, values map[string]interface{}) {
	debugMessage := fmt.Sprintf("Dispatched %s Event Type", eventType)
	logger.LogDebug(debugMessage)

	_, isExist := dispatcher.subscriptions[eventType]

	if !isExist {
		return
	}

	for _, handler := range dispatcher.subscriptions[eventType] {
		handler.DoHandle(values)
	}
}

func newDispatcher() interfaces.ICoreDispatcher {
	logger.LogDebug("Created Dispatcher")

	return &coreDispatcher{
		subscriptions: make(map[string]map[string]interfaces.ICoreHandler),
	}
}

func GetDispatcher() interfaces.ICoreDispatcher {
	if dispatcherInstance == nil {
		dispatcherInstance = newDispatcher()
	}

	return dispatcherInstance
}
