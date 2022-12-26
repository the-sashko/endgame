package map_index

func GetIndex(x uint16, y uint16) uint32 {
	index := uint32(x)
	index = (index << 16) | uint32(y)

	return index
}

func RetrieveXY(index uint32) (uint16, uint16) {
	x := index >> 16
	y := index - (x << 16)

	return uint16(x), uint16(y)
}
