package interfaces

type ICoreDispatcher interface {
	GetSubscriptions(handlerReference string) []string
	Subscribe(eventType string, handler ICoreHandler)
	Unsubscribe(eventType string, handler ICoreHandler)
	Dispatch(eventType string, values map[string]interface{})
}
