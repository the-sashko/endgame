package settings

var screenResolutionInstance IScreenResolution

const (
	defaultScreenResolutionWidth  = 640
	defaultScreenResolutionHeight = 480
	defaultIsFullScreen           = false
)

type screenResolution struct {
	width        uint16
	height       uint16
	isFullScreen bool
}

func (screenResolutionObject *screenResolution) GetWidth() uint16 {
	return screenResolutionObject.width
}

func (screenResolutionObject *screenResolution) GetHeight() uint16 {
	return screenResolutionObject.height
}

func (screenResolutionObject *screenResolution) IsFullScreen() bool {
	return screenResolutionObject.isFullScreen
}

func (screenResolutionObject *screenResolution) SetWidth(width uint16) {
	screenResolutionObject.width = width
}

func (screenResolutionObject *screenResolution) SetHeight(height uint16) {
	screenResolutionObject.height = height
}

func (screenResolutionObject *screenResolution) SetFullScreen(isFullScreen bool) {
	screenResolutionObject.isFullScreen = isFullScreen
}

func newScreenResolution() IScreenResolution {
	return &screenResolution{
		width:        defaultScreenResolutionWidth,
		height:       defaultScreenResolutionHeight,
		isFullScreen: defaultIsFullScreen,
	}
}

func getScreenResolution() IScreenResolution {
	if screenResolutionInstance == nil {
		screenResolutionInstance = newScreenResolution()
	}

	return screenResolutionInstance
}
