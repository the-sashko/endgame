package fps

import (
	"strconv"
	"time"
)

var fpsInstance *fps

const framesCount = 1

type fps struct {
	timestamp             int64
	microsecondsPerFrames []int64
	framesPerSecond       uint16
}

func (fpsObject *fps) StartFrame() {
	fpsObject.timestamp = time.Now().UnixMicro()
}

func (fpsObject *fps) EndFrame() {
	timestampDiff := time.Now().UnixMicro() - fpsObject.timestamp

	fpsObject.microsecondsPerFrames = append(
		fpsObject.microsecondsPerFrames,
		timestampDiff,
	)

	var microsecondsPerFrameSum int64 = 0

	for _, microsecondsPerFrame := range fpsObject.microsecondsPerFrames {
		microsecondsPerFrameSum += microsecondsPerFrame
	}

	microsecondsPerFrameAvg := microsecondsPerFrameSum / framesCount

	framesPerSecond := 1000000000 / microsecondsPerFrameAvg

	if framesPerSecond > 65535 {
		framesPerSecond = 65535
	}

	fpsObject.framesPerSecond = uint16(framesPerSecond)

	if len(fpsObject.microsecondsPerFrames) == framesCount {
		fpsObject.microsecondsPerFrames = make([]int64, 0)
	}
}

func (fpsObject *fps) GetFramesPerSecond() uint16 {
	return fpsObject.framesPerSecond
}

func (fpsObject *fps) GetFramesPerSecondString() string {
	if fpsObject.framesPerSecond == 0 {
		return "---"
	}

	if fpsObject.framesPerSecond > 999 {
		//return "999+"
	}

	return strconv.Itoa(int(fpsObject.framesPerSecond))
}

func newFps() *fps {
	return &fps{
		microsecondsPerFrames: make([]int64, 0),
		framesPerSecond:       0,
	}
}

func GetInstance() IFps {
	if fpsInstance == nil {
		fpsInstance = newFps()
	}

	return fpsInstance
}
