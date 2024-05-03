package main

import (
	"fmt"
	"sync"
)

var probabilityHashMap = &sync.Map{}

type DoubleBoard struct {
	circles EfficientBoard
	crosses EfficientBoard
	turn    PlayerType
}

func helper(s *TreeNode) EfficientBoard {
	l := MakeEfficientBoard()
	s.IterateOnePlayer(func(p *Point) bool {
		l.AddPoint(p[0], p[1])
		return false
	})
	return *l
}

func toDoubleBoard(move *TreeNode, p PlayerType) DoubleBoard {
	var d DoubleBoard
	if p == Circles {
		d.circles = helper(move)
		d.crosses = helper(move.parent)
	} else {
		d.crosses = helper(move)
		d.circles = helper(move.parent)
	}
	d.turn = p
	return d
}

// make(map[BoardHash]*[2]float64)
// var i = 0

func ProbabilityOfWin(move *TreeNode, p PlayerType, sendChannel chan ProbabilityOfWinReturn) {
	var circles, crosses float64
	h := toDoubleBoard(move, p)
	x, ok := (probabilityHashMap.Load(h))
	if ok {
		r := x.([2]float64)
		fmt.Println("Cache hit")
		sendChannel <- ProbabilityOfWinReturn{
			circles: r[0],
			crosses: r[1],
			move:    move,
		}
		return
	} else {
		defer func() {
			probabilityHashMap.Store(h, [2]float64{circles, crosses})
		}()
	}
	isGameOver, isWinner := IsBoardVictorious(move)
	if isGameOver {
		if isWinner {
			if p == Circles {
				circles, crosses = 1.0, 0.0
			} else if p == Crosses {
				circles, crosses = 0.0, 1.0
			}
		} else {
			circles, crosses = 0.0, 0.0
		}

		sendChannel <- ProbabilityOfWinReturn{
			circles: circles,
			crosses: crosses,
			move:    move,
		}
		return
	}

	possibleMoves := move.Expand()
	p = p.Toggle()
	// receiveChannel := make(chan ProbabilityOfWinReturn, len)
	receiveChannel := make(chan ProbabilityOfWinReturn, len(possibleMoves))
	for _, m := range possibleMoves {
		ProbabilityOfWin(&TreeNode{
			parent: move,
			point:  &m,
		}, p, receiveChannel)
	}
	for i := 0; i < len(possibleMoves); i++ {
		v := <-receiveChannel
		circles += v.circles
		crosses += v.crosses
	}
	circles = circles / float64(len(possibleMoves))
	crosses = crosses / float64(len(possibleMoves))
	// fmt.Println("Finishing execution", move.point)
	sendChannel <- ProbabilityOfWinReturn{
		circles: circles,
		crosses: crosses,
		move:    move,
	}
}
