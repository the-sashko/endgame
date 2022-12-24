package interfaces

type ICoreButton interface {
	GetName() string
	GetState() uint8
	SetState(state uint8)
}
