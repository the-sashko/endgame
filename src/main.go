package main

import (
	"endgame/src/actions"
	"endgame/src/display"
	"endgame/src/handlers"
	"endgame/src/input"
	"endgame/src/utils/fps"
	"endgame/src/utils/settings"
	"fmt"
)

func main() {
	initApp()

	appSettings := settings.GetSettings()
	appSettings.SetDebug(true)

	actions.GetGlobalAppStartAction().Fire(nil)

	//handlerTest := handlers.NewTestHandler()
	//handlerTest.Subscribe("e_click")
	handlers.NewClickHandler().Subscribe("global_event_click")

	display.Init(
		"EndGame Demo",
		640,
		480,
		false,
	)

	displayObject := display.GetInstance()

	defer displayObject.Quit()

	testImage := display.GetBitmapImage("test", "../res/test.bmp")

	displayObject.DrawBitmapImageToDefaultBuffer(testImage, 20, 20)

	_ = displayObject.AddTextWithDefaultStyle("Hello word!", 20, 20)

	for {
		if settings.GetSettings().IsDebug() {
			fps.GetInstance().StartFrame()
			display.GetInstance().SetWindowTitle(
				fmt.Sprintf("EndGame Demo - FPS: %s", fps.GetInstance().GetFramesPerSecondString()),
			)
		}

		loop()

		if settings.GetSettings().IsDebug() {
			fps.GetInstance().EndFrame()
		}
	}
}

func loop() {
	input.DoHandleInputs()
	display.GetInstance().Render()
}
