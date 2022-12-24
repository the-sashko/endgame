package keyboard

import (
	"endgame/src/core/input/button"
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"fmt"
)

var coreKeyboardInstance interfaces.ICoreKeyboard

type coreKeyboard struct {
	buttons map[string]interfaces.ICoreButton
}

func (keyboard *coreKeyboard) GetButton(
	buttonName string,
) interfaces.ICoreButton {
	buttonName = fmt.Sprintf("keyboard_%s_button", buttonName)

	if !keyboard.hasButton(buttonName) {
		keyboard.createButton(buttonName)
	}

	return keyboard.buttons[buttonName]
}

func (keyboard *coreKeyboard) hasButton(buttonName string) bool {
	_, isExist := keyboard.buttons[buttonName]

	return isExist
}

func (keyboard *coreKeyboard) createButton(buttonName string) {
	keyboard.buttons[buttonName] = button.NewButton(buttonName)
}

func newKeyboard() interfaces.ICoreKeyboard {
	logger.LogDebug("Created Core Keyboard Object")

	return &coreKeyboard{
		buttons: make(map[string]interfaces.ICoreButton),
	}
}

func GetKeyboard() interfaces.ICoreKeyboard {
	if coreKeyboardInstance == nil {
		coreKeyboardInstance = newKeyboard()
	}

	return coreKeyboardInstance
}
