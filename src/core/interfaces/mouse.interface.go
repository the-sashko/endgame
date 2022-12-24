package interfaces

type ICoreMouse interface {
	GetX() uint16
	GetY() uint16
	GetLeftButton() ICoreButton
	GetRightButton() ICoreButton
	GetMiddleButton() ICoreButton
	GetXOneButton() ICoreButton
	GetXTwoButton() ICoreButton
	SetX(x uint16)
	SetY(y uint16)
}
