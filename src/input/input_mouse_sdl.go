package input

import (
	"endgame/src/actions"
	"endgame/src/core/input/button"
	"endgame/src/core/input/mouse"
	"endgame/src/core/values_object"
	"github.com/veandco/go-sdl2/sdl"
)

func doHandleMouseButtonSdlEvent(sdlEvent sdl.Event) {
	mouseObject := mouse.GetMouse()

	values := values_object.NewValuesObject(make(map[string]interface{}))

	mouseObject.SetX(uint16(sdlEvent.(*sdl.MouseButtonEvent).X))
	mouseObject.SetY(uint16(sdlEvent.(*sdl.MouseButtonEvent).Y))

	mouseButtonName := mouse.UnknownButtonName
	mouseButtonState := retrieveButtonStateFromSdlEvent(sdlEvent)

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

	values.Set("mouse_button_name", mouseButtonName)
	values.Set("mouse_button_state", mouseButtonState)

	if sdlEvent.(*sdl.MouseButtonEvent).Clicks > 0 {
		values.Set("clicks", sdlEvent.(*sdl.MouseButtonEvent).Clicks)
		actions.GetGlobalMouseButtonClickAction().Fire(values)

		return
	}

	if mouseButtonState == button.StatePressed {
		actions.GetGlobalMouseButtonPressAction().Fire(values)

		return
	}

	if mouseButtonState == button.StateReleased {
		actions.GetGlobalMouseButtonReleaseAction().Fire(values)

		return
	}
}

func retrieveButtonStateFromSdlEvent(sdlEvent sdl.Event) uint8 {
	var mouseButtonState uint8

	mouseButtonState = button.StateUnknown

	switch sdlEvent.(*sdl.MouseButtonEvent).State {
	case sdl.PRESSED:
		mouseButtonState = button.StatePressed
		break
	case sdl.RELEASED:
		mouseButtonState = button.StateReleased
		break
	}

	return mouseButtonState
}

func doHandleMouseMotionSdlEvent(sdlEvent sdl.Event) {
	mouseObject := mouse.GetMouse()

	mouseObject.SetX(uint16(sdlEvent.(*sdl.MouseMotionEvent).X))
	mouseObject.SetY(uint16(sdlEvent.(*sdl.MouseMotionEvent).Y))

	actions.GetGlobalMouseMoveAction().Fire(nil)
}

func doHandleMouseWheelSdlEvent(sdlEvent sdl.Event) {
	values := values_object.NewValuesObject(make(map[string]interface{}))

	x := sdlEvent.(*sdl.MouseWheelEvent).X
	y := sdlEvent.(*sdl.MouseWheelEvent).Y

	preciseX := sdlEvent.(*sdl.MouseWheelEvent).PreciseX
	preciseY := sdlEvent.(*sdl.MouseWheelEvent).PreciseY

	if sdlEvent.(*sdl.MouseWheelEvent).Direction == sdl.MOUSEWHEEL_FLIPPED {
		x = -x
		y = -y
		preciseX = -preciseX
		preciseY = -preciseY
	}

	values.Set("x", x)
	values.Set("y", y)

	precise := make(map[string]float32)

	precise["x"] = preciseX
	precise["y"] = preciseY

	values.Set("precise", precise)

	actions.GetGlobalMouseScrollAction().Fire(values)
}
