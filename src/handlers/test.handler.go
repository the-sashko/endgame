package handlers

import (
	coreHandler "endgame/src/core/handler"
	"endgame/src/core/interfaces"
	"endgame/src/display"
	"math/rand"
)

type testHandler struct{}

func (handlerObject *testHandler) DoHandle(values map[string]interface{}) {
	return
	displayObject := display.GetInstance()

	/*displayObject.ClearScreen()
	  ref := displayObject.AddTextWithDefaultStyle("TEST", 0, 0)
	  displayObject.Render()
	  displayObject.RemoveText(ref)
	  time.Sleep(3 * time.Second)
	  displayObject.Render()
	  time.Sleep(3 * time.Second)*/

	for x := uint16(0); x < 640; x++ {
		for y := uint16(0); y < 480; y++ {
			displayObject.DrawPixelToDefaultBuffer(
				x,
				y,
				rand.Uint32(),
				rand.Uint32(),
				rand.Uint32(),
				0xffffff,
			)
		}
	}
}

func NewTestHandler() interfaces.ICoreHandler {
	return coreHandler.NewHandler("test", new(testHandler))
}
