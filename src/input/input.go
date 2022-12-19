package input

import (
	"endgame/src/actions"
	"endgame/src/core/action"
	"endgame/src/core/event"
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
		actions.GetQuitAppAction().Fire(nil)
	}()
}

func doHandleInputsSdl() {
	for sdlEvent := sdl.PollEvent(); sdlEvent != nil; sdlEvent = sdl.
		PollEvent() {
		switch sdlEvent.(type) {
		case *sdl.QuitEvent:
			actions.GetQuitAppAction().Fire(nil)
			break
		case *sdl.MouseButtonEvent:

			values := make(map[string]interface{})
			values["x"] = uint16(sdlEvent.(*sdl.MouseButtonEvent).X)
			values["y"] = uint16(sdlEvent.(*sdl.MouseButtonEvent).Y)

			eventClick := event.NewEvent("e_click")

			actionClick := action.NewAction("a_click", nil)
			actionClick.AddEvent(eventClick)

			actionClick.Fire(values)

			break
		}
	}
}
