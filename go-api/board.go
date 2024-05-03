package main

type EfficientBoard [RowCount]uint64

func MakeEfficientBoard() *EfficientBoard {
	var a EfficientBoard
	return &a
}

func (b *EfficientBoard) AddPoint(i, j int) {
	b[i] = b[i] | 1 << (ColCount - 1 - j)
}
func (b *EfficientBoard) HasPoint(i, j int) bool {
	val := b[i] & (1 << (ColCount - 1 - j))
	return (val > 0)
}