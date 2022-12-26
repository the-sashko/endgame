package settings

type ISettings interface {
	IsDebug() bool
	SetDebug(isDebug bool)
}
