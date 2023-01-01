package global_state

import (
	"endgame/src/display"
	"endgame/src/settings"
)

type IGlobalState interface {
	Get(name string) interface{}
	Set(name string, value interface{})
	Has(name string) bool
	Delete(name string)
	GetMap(mapName string) display.IDisplayMap
	SetMap(mapName string, displayMap display.IDisplayMap)
	DeleteMap(mapName string)
	GetSettings() settings.ISettings
}
