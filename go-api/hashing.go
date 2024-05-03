package main

type SingleBoardAsBits [RowCount]uint64

type BoardHash struct {
	turn    PlayerType
	circles SingleBoardAsBits
	crosses SingleBoardAsBits
}
