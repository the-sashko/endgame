package settings

type ISettings interface {
	IsDebug() bool
	SetDebug(isDebug bool)
	GetAppName() string
	SetAppName(appName string)
	GetScreenResolution() IScreenResolution
	SetScreenResolution(ScreenResolutionObject IScreenResolution)
}
