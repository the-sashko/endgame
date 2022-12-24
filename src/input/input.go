package input

import (
	"endgame/src/actions"
	"endgame/src/core/input/button"
	"endgame/src/core/input/mouse"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	initSyscallHandling()
}

func DoHandleInputs() {
	doHandleInputsSdl()
}

func initSyscallHandling() {
	osSignalChannel := make(chan os.Signal, 1)

	signal.Notify(
		osSignalChannel,
		os.Interrupt,
		os.Kill,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-osSignalChannel
		actions.GetGlobalAppQuitAction().Fire(nil)
	}()
}

func doHandleInputsSdl() {
	for sdlEvent := sdl.PollEvent(); sdlEvent != nil; sdlEvent = sdl.
		PollEvent() {
		switch sdlEvent.(type) {
		case *sdl.QuitEvent:
			actions.GetGlobalAppQuitAction().Fire(nil)
			break
		case *sdl.MouseButtonEvent:
			var mouseButtonName string
			var mouseButtonState uint8

			mouseObject := mouse.GetMouse()

			values := make(map[string]interface{})

			mouseObject.SetX(uint16(sdlEvent.(*sdl.MouseButtonEvent).X))
			mouseObject.SetY(uint16(sdlEvent.(*sdl.MouseButtonEvent).Y))

			mouseButtonName = mouse.UnknownButtonName
			mouseButtonState = button.StateUnknown

			switch sdlEvent.(*sdl.MouseButtonEvent).State {
			case sdl.PRESSED:
				mouseButtonState = button.StatePressed
				break
			case sdl.RELEASED:
				mouseButtonState = button.StateReleased
				break
			}

			switch sdlEvent.(*sdl.MouseButtonEvent).Button {
			case sdl.BUTTON_LEFT:
				mouseObject.GetLeftButton().SetState(mouseButtonState)
				mouseButtonName = mouse.LeftButtonName
				break
			case sdl.BUTTON_RIGHT:
				mouseObject.GetRightButton().SetState(mouseButtonState)
				mouseButtonName = mouse.RightButtonName
				break
			case sdl.BUTTON_MIDDLE:
				mouseObject.GetMiddleButton().SetState(mouseButtonState)
				mouseButtonName = mouse.MiddleButtonName
				break
			case sdl.BUTTON_X1:
				mouseObject.GetXOneButton().SetState(mouseButtonState)
				mouseButtonName = mouse.XOneButtonName
				break
			case sdl.BUTTON_X2:
				mouseObject.GetXTwoButton().SetState(mouseButtonState)
				mouseButtonName = mouse.XTwoButtonName
				break
			}

			values["mouse_button_name"] = mouseButtonName
			values["mouse_button_state"] = mouseButtonState

			if sdlEvent.(*sdl.MouseButtonEvent).Clicks > 0 {
				values["clicks"] = sdlEvent.(*sdl.MouseButtonEvent).Clicks
				actions.GetGlobalMouseButtonClickAction().Fire(values)

				continue
			}

			if mouseButtonState == button.StatePressed {
				actions.GetGlobalMouseButtonPressAction().Fire(values)

				continue
			}

			if mouseButtonState == button.StateReleased {
				actions.GetGlobalMouseButtonReleaseAction().Fire(values)

				continue
			}

			break
		}
	}
}
