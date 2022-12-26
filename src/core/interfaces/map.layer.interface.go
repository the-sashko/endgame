package interfaces

type ICoreMapLayer interface {
	GetReference() string
	GetName() string
	GetObject(index uint32) ICoreObject
	SetObject(object ICoreObject, index uint32)
	DeleteObject(index uint32)
	GetObjectsForArea(
		areaX uint16,
		areaY uint16,
		areaWidth uint16,
		areaHeight uint16,
	) map[uint32]ICoreObject
}
