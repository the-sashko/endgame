package main

import (
	"endgame/src/actions"
	"endgame/src/display"
	"endgame/src/events"
	"endgame/src/golbal_state"
	"endgame/src/handlers"
	"endgame/src/input"
)

func initApp() {
	initSettings()
	initGlobalActionsAndEvents()
	initGlobalHandlers()
	initDisplay()
	initMaps()
	initScene()

	input.Init()
}

func initSettings() {
	golbal_state.GetGlobalState().GetSettings().SetDebug(true)
}

func initGlobalActionsAndEvents() {
	actions.GetGlobalAppStartAction().AddEvent(
		events.GetGlobalAppStartEvent(),
	)

	actions.GetGlobalAppQuitAction().AddEvent(
		events.GetGlobalAppQuitEvent(),
	)

	actions.GetGlobalKeyboardKeyPressAction().AddEvent(
		events.GetGlobalKeyboardKeyPressEvent(),
	)

	actions.GetGlobalKeyboardKeyReleaseAction().AddEvent(
		events.GetGlobalKeyboardKeyReleaseEvent(),
	)

	actions.GetGlobalMouseMoveAction().AddEvent(
		events.GetGlobalMouseMoveEvent(),
	)

	actions.GetGlobalMouseButtonPressAction().AddEvent(
		events.GetGlobalMouseButtonPressEvent(),
	)

	actions.GetGlobalMouseButtonReleaseAction().AddEvent(
		events.GetGlobalMouseButtonReleaseEvent(),
	)

	actions.GetGlobalMouseButtonClickAction().AddEvent(
		events.GetGlobalMouseButtonClickEvent(),
	)

	actions.GetGlobalMouseMoveAction().AddEvent(
		events.GetGlobalMouseMoveEvent(),
	)

	actions.GetGlobalMouseScrollAction().AddEvent(
		events.GetGlobalMouseScrollEvent(),
	)
}

func initGlobalHandlers() {
	handlers.GetGlobalAppStartHandler().Subscribe(events.GlobalAppStartType)
	handlers.GetGlobalAppQuitHandler().Subscribe(events.GlobalAppQuitType)
}

func initDisplay() {
	display.Init(
		golbal_state.GetGlobalState().GetSettings().GetAppName(),
		golbal_state.GetGlobalState().GetSettings().GetScreenResolution().GetWidth(),
		golbal_state.GetGlobalState().GetSettings().GetScreenResolution().GetHeight(),
		golbal_state.GetGlobalState().GetSettings().GetScreenResolution().IsFullScreen(),
	)
}

func initMaps() {
	defMap := display.NewDisplayMap("default")
	golbal_state.GetGlobalState().SetMap("default", defMap)
}

func initScene() {
	scene := display.GetDisplayScene()
	scene.SetX(0)
	scene.SetY(0)
}
