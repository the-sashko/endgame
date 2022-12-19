package fps

type IFps interface {
	StartFrame()
	EndFrame()
	GetFramesPerSecond() uint16
	GetFramesPerSecondString() string
}
