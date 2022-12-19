package interfaces

type ICoreAction interface {
	GetReference() string
	GetName() string
	Fire(values map[string]interface{})
	GetEvents() map[string]ICoreEvent
	HasEvent(eventType string) bool
	GetEvent(eventType string) ICoreEvent
	AddEvent(event ICoreEvent)
	RemoveEvent(eventType string)
}
