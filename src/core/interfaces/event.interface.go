package interfaces

type ICoreEvent interface {
	GetReference() string
	GetTypeName() string
	Fire(values ICoreValuesObject)
}
