package main

import (
	"endgame/src/actions"
	"endgame/src/events"
	"endgame/src/handlers"
	"endgame/src/input"
)

func initApp() {
	initGlobalActionsAndEvents()
	initGlobalHandlers()

	input.Init()
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
