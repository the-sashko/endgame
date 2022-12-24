package main

import (
	"endgame/src/actions"
	"endgame/src/display"
	"endgame/src/events"
	"endgame/src/handlers"
	"endgame/src/input"
	"endgame/src/utils/settings"
)

func initApp() {
	initSettings()
	initGlobalActionsAndEvents()
	initGlobalHandlers()
	iniDisplay()

	input.Init()
}

func initSettings() {
	appSettings := settings.GetSettings()
	appSettings.SetDebug(true)
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
}

func initGlobalHandlers() {
	handlers.GetGlobalAppStartHandler().Subscribe(events.GlobalAppStartType)
	handlers.GetGlobalAppQuitHandler().Subscribe(events.GlobalAppQuitType)
}

func iniDisplay() {
	display.Init(
		"EndGame Demo",
		640,
		480,
		false,
	)
}
