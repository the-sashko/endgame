package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/input/mouse"
	"endgame/src/core/interfaces"
	"endgame/src/display"
)

var testMoveHandlerInstance interfaces.ICoreHandler

type testMoveHandler struct{}

func (handlerObject *testMoveHandler) DoHandle(values map[string]interface{}) {
	x := mouse.GetMouse().GetX()
	y := mouse.GetMouse().GetY()

	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y-1, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x-1, y+1, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y-1, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x, y+1, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y-1, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y, 0, 0xffffff, 0,
		0xffffff)
	display.GetInstance().DrawPixelToDefaultBuffer(x+1, y+1, 0, 0xffffff, 0,
		0xffffff)
}

func GetTestMoveHandler() interfaces.ICoreHandler {
	if testMoveHandlerInstance == nil {
		testMoveHandlerInstance = coreHandler.NewHandler(
			"test_move_handler",
			new(testMoveHandler),
		)
	}

	return testMoveHandlerInstance
}
