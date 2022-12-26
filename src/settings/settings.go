package settings

import (
	"endgame/src/core/interfaces"
	"endgame/src/core/values_object"
)

const (
	defaultDebug   = false
	defaultAppName = "Endgame"
)

var settingsInstance ISettings

type settingsType struct {
	values interfaces.ICoreValuesObject
}

func (settings *settingsType) IsDebug() bool {
	return settings.values.Get("is_debug").(bool)
}

func (settings *settingsType) SetDebug(isDebug bool) {
	settings.values.Set("is_debug", isDebug)
}

func (settings *settingsType) GetAppName() string {
	return settings.values.Get("app_name").(string)
}

func (settings *settingsType) SetAppName(appName string) {
	settings.values.Set("app_name", appName)
}

func (settings *settingsType) GetScreenResolution() IScreenResolution {
	return settings.values.Get("screen_resolution").(IScreenResolution)
}

func (settings *settingsType) SetScreenResolution(
	ScreenResolutionObject IScreenResolution,
) {
	settings.values.Set("screen_resolution", ScreenResolutionObject)
}

func newSettings() ISettings {
	settingsInstance = &settingsType{
		values: values_object.NewValuesObject(make(map[string]interface{})),
	}

	settingsInstance.SetDebug(defaultDebug)
	settingsInstance.SetAppName(defaultAppName)
	settingsInstance.SetScreenResolution(newScreenResolution())

	return settingsInstance
}

func GetSettings() ISettings {
	if settingsInstance == nil {
		settingsInstance = newSettings()
	}

	return settingsInstance
}
