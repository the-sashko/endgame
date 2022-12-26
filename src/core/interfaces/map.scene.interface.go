package interfaces

type ICoreScene interface {
	GetX() uint16
	GetY() uint16
	GetWidth() uint16
	GetHeight() uint16
	GetObjects() map[string]map[uint32]ICoreObject
	SetX(x uint16)
	SetY(y uint16)
	SetObjects(objects map[string]map[uint32]ICoreObject)
}
