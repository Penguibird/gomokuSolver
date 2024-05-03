package main

import (
	"fmt"
	"testing"
)

func TestAddToBoard(t *testing.T) {
	g := Game{
		currentState:   nil,
		startingPlayer: Circles,
	}
	t.Run("Adding to board creates good result", func(t *testing.T) {
		g.MakeOpponentMove(0, 0)
		g.MakeOpponentMove(1, 1)
		g.MakeOpponentMove(0, 2)
		g.MakeOpponentMove(2, 3)
		i := 0
		s := ""
		g.currentState.IterateAllFields(func(p *Point) bool {
			i++
			s += fmt.Sprint(p[1])
			return false
		})
		if i != 4 || s != "3210" {
			t.Log(sliceFromIterator(g.currentState.IterateAllFields))
			t.Error(i, s)
		}
	})
}

func TestEfficientBoard(t *testing.T) {
	t.Run("Adding to board puts the point there", func(t *testing.T) {
		b := MakeEfficientBoard()
		b.AddPoint(0, 1)
		if !b.HasPoint(0, 1) {
			t.Error("Board doesnt have the point added", b)
		}
	})
}

func TestIsBoardVictorious(t *testing.T) {
	t.Run("Game correctly adds bits", TestAddToBoard)

	t.Run("Winning board is correctly recognized", func(t *testing.T) {
		g := Game{
			currentState:   nil,
			startingPlayer: Circles,
		}
		g.MakeOpponentMove(0, 0)
		g.MakeOpponentMove(3, 3)

		g.MakeOpponentMove(0, 1)
		g.MakeOpponentMove(2, 0)

		g.MakeOpponentMove(0, 2)
		g.MakeOpponentMove(1, 3)

		g.MakeOpponentMove(0, 3)

		b, w := IsBoardVictorious(g.currentState)
		if !b || !w {
			t.Error(b, w, g)
		}
	})

	t.Run("Full board is correctly recognized", func(t *testing.T) {
		g := Game{
			currentState:   nil,
			startingPlayer: Circles,
		}
		g.MakeOpponentMove(0, 0)
		g.MakeOpponentMove(0, 1)
		g.MakeOpponentMove(0, 2)
		g.MakeOpponentMove(0, 3)
		g.MakeOpponentMove(1, 0)
		g.MakeOpponentMove(1, 1)
		g.MakeOpponentMove(1, 2)
		g.MakeOpponentMove(1, 3)
		g.MakeOpponentMove(2, 0)
		g.MakeOpponentMove(2, 1)
		g.MakeOpponentMove(2, 2)
		g.MakeOpponentMove(2, 3)
		g.MakeOpponentMove(3, 0)
		g.MakeOpponentMove(3, 1)
		g.MakeOpponentMove(3, 2)
		g.MakeOpponentMove(3, 3)

		b, w := IsBoardVictorious(g.currentState)
		if !(b && !w) {
			t.Error(b, w, g)
		}
	})
}
