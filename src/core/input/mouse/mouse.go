package mouse

import (
	"endgame/src/core/input/button"
	"endgame/src/core/interfaces"
	"endgame/src/utils/logger"
)

var coreMouseInstance interfaces.ICoreMouse

type coreMouse struct {
	x            uint16
	y            uint16
	leftButton   interfaces.ICoreButton
	rightButton  interfaces.ICoreButton
	middleButton interfaces.ICoreButton
	xOneButton   interfaces.ICoreButton
	xTwoButton   interfaces.ICoreButton
}

func (mouse *coreMouse) GetX() uint16 {
	return mouse.x
}

func (mouse *coreMouse) GetY() uint16 {
	return mouse.y
}

func (mouse *coreMouse) GetLeftButton() interfaces.ICoreButton {
	return mouse.leftButton
}

func (mouse *coreMouse) GetRightButton() interfaces.ICoreButton {
	return mouse.rightButton
}

func (mouse *coreMouse) GetMiddleButton() interfaces.ICoreButton {
	return mouse.middleButton
}

func (mouse *coreMouse) GetXOneButton() interfaces.ICoreButton {
	return mouse.xOneButton
}

func (mouse *coreMouse) GetXTwoButton() interfaces.ICoreButton {
	return mouse.xTwoButton
}

func (mouse *coreMouse) SetX(x uint16) {
	mouse.x = x
}

func (mouse *coreMouse) SetY(y uint16) {
	mouse.y = y
}

func newMouse() interfaces.ICoreMouse {
	logger.LogDebug("Created Core Mouse Object")

	return &coreMouse{
		x:            0,
		y:            0,
		leftButton:   button.NewButton(LeftButtonName),
		rightButton:  button.NewButton(RightButtonName),
		middleButton: button.NewButton(MiddleButtonName),
		xOneButton:   button.NewButton(XOneButtonName),
		xTwoButton:   button.NewButton(XTwoButtonName),
	}
}

func GetMouse() interfaces.ICoreMouse {
	if coreMouseInstance == nil {
		coreMouseInstance = newMouse()
	}

	return coreMouseInstance
}
