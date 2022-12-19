package interfaces

type ICoreHandler interface {
	GetReference() string
	GetName() string
	DoHandle(values map[string]interface{})
	Subscribe(eventType string)
	Unsubscribe(eventType string)
	GetSubscriptions() []string
}
