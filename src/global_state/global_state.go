package global_state

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/values_object"
	"endgame/src/display"
	settingsPackage "endgame/src/settings"
)

const defaultMapName = "default"

var globalStateInstance IGlobalState

type globalState struct {
	valuesObject interfaces.ICoreValuesObject
}

func (globalStateObject *globalState) Get(name string) interface{} {
	return globalStateObject.valuesObject.Get(name)
}

func (globalStateObject *globalState) Set(name string, value interface{}) {
	globalStateObject.valuesObject.Set(name, value)
}

func (globalStateObject *globalState) Has(name string) bool {
	return globalStateObject.valuesObject.Has(name)
}

func (globalStateObject *globalState) Delete(name string) {
	globalStateObject.Has(name)
}

func (globalStateObject *globalState) GetCurrentMapName() string {
	return globalStateObject.Get("current_map_name").(string)
}

func (globalStateObject *globalState) GetMaps() interfaces.ICoreValuesObject {
	return globalStateObject.Get("maps").(interfaces.ICoreValuesObject)
}

func (globalStateObject *globalState) GetMap(
	mapName string,
) display.IDisplayMap {
	return globalStateObject.GetMaps().Get(mapName).(display.IDisplayMap)
}

func (globalStateObject *globalState) GetCurrentMap() display.IDisplayMap {
	currentMapName := globalStateObject.GetCurrentMapName()

	return globalStateObject.GetMap(currentMapName).(display.IDisplayMap)
}

func (globalStateObject *globalState) GetObjects() interfaces.ICoreValuesObject {
	return globalStateObject.Get("objects").(interfaces.ICoreValuesObject)
}

func (globalStateObject *globalState) GetObject(
	objectName string,
) display.IDisplayObject {
	return globalStateObject.GetObjects().Get(objectName).(display.IDisplayObject)
}

func (globalStateObject *globalState) SetCurrentMapName(mapName string) {
	globalStateObject.Set("current_map_name", mapName)
}

func (globalStateObject *globalState) SetMap(displayMap display.IDisplayMap) {
	globalStateObject.GetMaps().Set(displayMap.GetName(), displayMap)
}

func (globalStateObject *globalState) SetObject(
	displayObject display.IDisplayObject,
) {
	globalStateObject.GetObjects().Set(displayObject.GetName(), displayObject)
}

func (globalStateObject *globalState) DeleteMap(mapName string) {
	globalStateObject.GetMaps().Delete(mapName)
}

func (globalStateObject *globalState) DeleteObject(objectName string) {
	if !globalStateObject.GetObjects().Has(objectName) {
		return
	}

	globalStateObject.GetObject(objectName).Destroy()

	globalStateObject.GetObjects().Delete(objectName)
}

func (globalStateObject *globalState) GetSettings() settingsPackage.ISettings {
	return globalStateObject.Get("settings").(settingsPackage.ISettings)
}

func newGlobalState() IGlobalState {
	globalStateObject := &globalState{
		valuesObject: values_object.NewValuesObject(
			make(map[string]interface{}),
		),
	}

	globalStateObject.Set(
		"maps",
		values_object.NewValuesObject(
			make(map[string]interface{}),
		),
	)

	globalStateObject.Set(
		"objects",
		values_object.NewValuesObject(
			make(map[string]interface{}),
		),
	)

	globalStateObject.SetCurrentMapName(defaultMapName)
	globalStateObject.Set("settings", settingsPackage.GetSettings())

	return globalStateObject
}

func GetGlobalState() IGlobalState {
	if globalStateInstance == nil {
		globalStateInstance = newGlobalState()
	}

	return globalStateInstance
}
