package global_state

import (
	"endgame/src/core/interfaces"
	"endgame/src/display"
	settingsPackage "endgame/src/settings"
)

type IGlobalState interface {
	Get(name string) interface{}
	Set(name string, value interface{})
	Has(name string) bool
	Delete(name string)
	GetCurrentMapName() string
	GetMaps() interfaces.ICoreValuesObject
	GetMap(mapName string) display.IDisplayMap
	GetCurrentMap() display.IDisplayMap
	GetObjects() interfaces.ICoreValuesObject
	GetObject(objectName string) display.IDisplayObject
	SetCurrentMapName(mapName string)
	SetMap(displayMap display.IDisplayMap)
	SetObject(displayObject display.IDisplayObject)
	DeleteMap(mapName string)
	DeleteObject(objectName string)
	GetSettings() settingsPackage.ISettings
}
