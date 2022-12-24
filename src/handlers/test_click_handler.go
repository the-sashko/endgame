package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/input/mouse"
	"endgame/src/core/interfaces"
	"endgame/src/display"
)

var testClickHandlerInstance interfaces.ICoreHandler

type testClickHandler struct{}

func (handlerObject *testClickHandler) DoHandle(values map[string]interface{}) {
	x := mouse.GetMouse().GetX()
	y := mouse.GetMouse().GetY()

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

func GetTestClickHandler() interfaces.ICoreHandler {
	if testClickHandlerInstance == nil {
		testClickHandlerInstance = coreHandler.NewHandler(
			"test_click_handler",
			new(testClickHandler),
		)
	}

	return testClickHandlerInstance
}
