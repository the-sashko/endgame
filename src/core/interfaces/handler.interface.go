package interfaces

type ICoreHandler interface {
	GetReference() string
	GetName() string
	DoHandle(values ICoreValuesObject)
	Subscribe(eventType string)
	Unsubscribe(eventType string)
	GetSubscriptions() []string
}
