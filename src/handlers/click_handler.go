package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/display"
)

type clickHandler struct{}

func (handlerObject *clickHandler) DoHandle(values map[string]interface{}) {
	x := values["x"].(uint16)
	y := values["y"].(uint16)
	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y-1, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y+1, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y-1, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y+1, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y-1, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y, 0xffffff, 0, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y+1, 0xffffff, 0, 0,
		0xffffff)
}

func NewClickHandler() interfaces.ICoreHandler {
	return coreHandler.NewHandler("click", new(clickHandler))
}
