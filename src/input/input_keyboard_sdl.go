package input

import (
	"endgame/src/actions"
	"endgame/src/core/input/button"
	"endgame/src/core/input/keyboard"
	"endgame/src/core/values_object"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
)

func doHandleKeyboardSdlEvent(sdlEvent sdl.Event) {
	values := values_object.NewValuesObject(make(map[string]interface{}))

	buttonName := retrieveButtonNameFromSdlEvent(sdlEvent)
	buttonState := retrieveKeyboardButtonStateFromSdlEvent(sdlEvent)

	keyboard.GetKeyboard().GetButton(buttonName).SetState(buttonState)

	values.Set("button_name", buttonName)

	if buttonState == button.StatePressed {
		actions.GetGlobalKeyboardKeyPressAction().Fire(values)

		return
	}

	if buttonState == button.StateReleased {
		actions.GetGlobalKeyboardKeyReleaseAction().Fire(values)
	}
}

func retrieveButtonNameFromSdlEvent(sdlEvent sdl.Event) string {
	buttonName := sdl.GetScancodeName(
		sdlEvent.(*sdl.KeyboardEvent).Keysym.Scancode,
	)

	if buttonName == "" {
		buttonName = "unknown"
	}

	buttonName = strings.ToLower(buttonName)

	return strings.Replace(buttonName, " ", "_", -1)
}

func retrieveKeyboardButtonStateFromSdlEvent(sdlEvent sdl.Event) uint8 {
	var buttonState uint8

	buttonState = button.StateUnknown

	if sdlEvent.(*sdl.MouseButtonEvent).Type == sdl.MOUSEBUTTONDOWN {
		buttonState = button.StatePressed
	}

	if sdlEvent.(*sdl.MouseButtonEvent).Type == sdl.MOUSEBUTTONUP {
		buttonState = button.StateReleased
	}

	return buttonState
}
