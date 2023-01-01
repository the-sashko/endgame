package main

import (
	"endgame/src/actions"
	"endgame/src/display"
	"endgame/src/events"
	"endgame/src/global_state"
	"endgame/src/handlers"
	"endgame/src/input"
	"endgame/src/objects"
	"endgame/src/settings"
	"endgame/src/utils/fps"
	"endgame/src/utils/map_index"
	"fmt"
)

func main() {
	initApp()

	actions.GetGlobalAppStartAction().Fire(nil)

	handlers.GetTestClickHandler().Subscribe(events.GlobalMouseButtonClickType)
	handlers.GetTestMoveHandler().Subscribe(events.GlobalMouseMoveType)
	handlers.GetTestKeyboardHandler().Subscribe(events.GlobalKeyboardKeyReleaseType)

	displayObject := display.GetInstance()

	defer displayObject.Quit()

	//testImage := display.GetBitmapImage("test", "../res/test.bmp")

	//displayObject.DrawBitmapImageToDefaultBuffer(testImage, 20, 20)

	//_ = displayObject.AddTextWithDefaultStyle("Hello word!", 20, 20)

	x := uint16(5)
	y := uint16(5)

	global_state.GetGlobalState().GetMap("default").SetObjectToDefaultLayer(
		objects.NewRedSquareObject("test"),
		x,
		y,
	)

	global_state.GetGlobalState().Set("test_obj_x", x)
	global_state.GetGlobalState().Set("test_obj_y", y)

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

	dso := display.GetDisplayScene()

	mp := global_state.GetGlobalState().GetMap("default")

	objs := mp.GetObjectsForArea(dso.GetX(), dso.GetY(), dso.GetWidth(), dso.GetHeight())

	dso.SetObjects(objs)

	dso.DoFillPixels()

	for bufferLayer, pixels := range dso.GetPixels() {
		for index, pixel := range pixels {
			x, y := map_index.RetrieveXY(index)

			red, green, blue, alpha := pixel.RGBA()

			display.GetInstance().DrawPixel(
				x,
				y,
				red,
				green,
				blue,
				alpha,
				bufferLayer,
			)
		}
	}

	display.GetInstance().Render()
}
