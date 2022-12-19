package interfaces

type ICoreHandlerAdapter interface {
	DoHandle(value map[string]interface{})
}
