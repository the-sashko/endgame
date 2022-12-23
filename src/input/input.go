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
			values := make(map[string]interface{})

			values["x"] = uint16(sdlEvent.(*sdl.MouseButtonEvent).X)
			values["y"] = uint16(sdlEvent.(*sdl.MouseButtonEvent).Y)
			values["button"] = sdlEvent.(*sdl.MouseButtonEvent).Button
			values["state"] = sdlEvent.(*sdl.MouseButtonEvent).State
			values["clicks"] = sdlEvent.(*sdl.MouseButtonEvent).Clicks
			values["type"] = sdlEvent.(*sdl.MouseButtonEvent).Type

			eventClick := event.NewEvent("global_event_click")

			actionClick := action.NewAction("global_action_click", nil)
			actionClick.AddEvent(eventClick)

			actionClick.Fire(values)

			break
		}
	}
}
