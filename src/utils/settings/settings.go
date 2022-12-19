package settings

var settingsInstance *settingsType

type settingsType struct {
	isDebug bool
}

func (settings *settingsType) IsDebug() bool {
	return settings.isDebug
}

func (settings *settingsType) SetDebug(isDebug bool) {
	settings.isDebug = isDebug
}

func newSettings() *settingsType {
	return &settingsType{
		isDebug: false,
	}
}

func GetSettings() *settingsType {
	if settingsInstance == nil {
		settingsInstance = newSettings()
	}

	return settingsInstance
}
