package button

import (
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
	"fmt"
)

type coreButton struct {
	name  string
	state uint8
}

func (button *coreButton) GetName() string {
	return button.name
}

func (button *coreButton) GetState() uint8 {
	return button.state
}

func (button *coreButton) SetState(state uint8) {
	button.state = state
}

func NewButton(name string) interfaces.ICoreButton {
	debugMessage := fmt.Sprintf("Created %s Button", name)
	logger.LogDebug(debugMessage)

	return &coreButton{
		name:  name,
		state: StateReleased,
	}
}
