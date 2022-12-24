package input

import (
	"endgame/src/actions"
	"github.com/veandco/go-sdl2/sdl"
)

func doHandleInputsFromSdl() {
	for sdlEvent := sdl.PollEvent(); sdlEvent != nil; sdlEvent = sdl.PollEvent() {
		doHandleSdlEvent(sdlEvent)
	}
}

func doHandleSdlEvent(sdlEvent sdl.Event) {
	switch sdlEvent.(type) {
	case *sdl.QuitEvent:
		doHandleQuitSdlEvent()
		break
	case *sdl.MouseButtonEvent:
		doHandleMouseButtonSdlEvent(sdlEvent)
		break

	case *sdl.MouseMotionEvent:
		doHandleMouseMotionSdlEvent(sdlEvent)
		break

	case *sdl.MouseWheelEvent:
		doHandleMouseWheelSdlEvent(sdlEvent)
		break

	case *sdl.KeyboardEvent:
		doHandleKeyboardSdlEvent(sdlEvent)
		break
	}
}

func doHandleQuitSdlEvent() {
	actions.GetGlobalAppQuitAction().Fire(nil)
}
