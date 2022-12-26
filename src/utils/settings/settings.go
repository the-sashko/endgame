package settings

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/values_object"
)

var settingsInstance *settingsType

type settingsType struct {
	values interfaces.ICoreValuesObject
}

func (settings *settingsType) IsDebug() bool {
	return settings.values.Get("is_debug").(bool)
}

func (settings *settingsType) SetDebug(isDebug bool) {
	settings.values.Set("is_debug", isDebug)
}

func newSettings() *settingsType {
	return &settingsType{
		values: values_object.NewValuesObject(make(map[string]interface{})),
	}
}

func GetSettings() ISettings {
	if settingsInstance == nil {
		settingsInstance = newSettings()
	}

	return settingsInstance
}
