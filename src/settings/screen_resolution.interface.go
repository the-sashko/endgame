package settings

type IScreenResolution interface {
	GetWidth() uint16
	GetHeight() uint16
	IsFullScreen() bool
	SetWidth(width uint16)
	SetHeight(height uint16)
	SetFullScreen(isFullScreen bool)
}
