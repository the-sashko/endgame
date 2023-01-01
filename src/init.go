package main

import (
	"endgame/src/actions"
	"endgame/src/display"
	"endgame/src/events"
	"endgame/src/global_state"
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
	global_state.GetGlobalState().GetSettings().SetDebug(true)
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
		global_state.GetGlobalState().GetSettings().GetAppName(),
		global_state.GetGlobalState().GetSettings().GetScreenResolution().GetWidth(),
		global_state.GetGlobalState().GetSettings().GetScreenResolution().GetHeight(),
		global_state.GetGlobalState().GetSettings().GetScreenResolution().IsFullScreen(),
	)
}

func initMaps() {
	defaultMap := display.NewDisplayMap("default")
	global_state.GetGlobalState().SetMap("default", defaultMap)
}

func initScene() {
	scene := display.GetDisplayScene()
	scene.SetX(0)
	scene.SetY(0)
}
