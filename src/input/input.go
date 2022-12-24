package input

import (
	"endgame/src/actions"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	initSyscallHandling()
}

func DoHandleInputs() {
	doHandleInputsFromSdl()
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
