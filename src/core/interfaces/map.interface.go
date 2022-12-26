package interfaces

type ICoreMap interface {
	GetReference() string
	GetName() string
	GetLayers() map[string]ICoreMapLayer
	GetLayer(layerName string) ICoreMapLayer
	SetLayer(layerName string)
	DeleteLayer(layerName string)
	GetObject(x uint16, y uint16, layerName string) ICoreObject
	SetObject(object ICoreObject, x uint16, y uint16, layerName string)
	DeleteObject(x uint16, y uint16, layerName string)
	GetObjectsForArea(
		areaX uint16,
		areaY uint16,
		areaWidth uint16,
		areaHeight uint16,
	) map[string]map[uint32]ICoreObject
}
