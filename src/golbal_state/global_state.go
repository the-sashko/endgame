package golbal_state

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/values_object"
	"endgame/src/display"
	"endgame/src/utils/settings"
	"fmt"
)

var globalStateInstance IGlobalState

type globalState struct {
	valuesObject   interfaces.ICoreValuesObject
	settingsObject settings.ISettings
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

func (globalStateObject *globalState) GetMap(
	mapName string,
) display.IDisplayMap {
	valueName := globalStateObject.retrieveMapValueName(mapName)

	if !globalStateObject.Has(valueName) {
		return nil
	}

	return globalStateObject.Get(valueName).(display.IDisplayMap)
}

func (globalStateObject *globalState) SetMap(
	mapName string,
	displayMap display.IDisplayMap,
) {
	valueName := globalStateObject.retrieveMapValueName(mapName)

	globalStateObject.Set(valueName, displayMap)
}

func (globalStateObject *globalState) DeleteMap(mapName string) {
	valueName := globalStateObject.retrieveMapValueName(mapName)

	globalStateObject.Delete(valueName)
}

func (globalStateObject *globalState) retrieveMapValueName(
	mapName string,
) string {
	return fmt.Sprintf("%s_display_map", mapName)
}

func (globalStateObject *globalState) GetSettings() settings.ISettings {
	return globalStateObject.settingsObject
}

func newGlobalState() IGlobalState {
	return &globalState{
		valuesObject: values_object.NewValuesObject(
			make(map[string]interface{}),
		),
		settingsObject: settings.GetSettings(),
	}
}

func GetGlobalState() IGlobalState {
	if globalStateInstance == nil {
		globalStateInstance = newGlobalState()
	}

	return globalStateInstance
}
