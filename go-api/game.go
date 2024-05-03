package main

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
)

type Game struct {
	currentState   *TreeNode
	startingPlayer PlayerType
}

type Point [2]int

func (g *Game) MakeOpponentMove(i, j int) {
	newNode := &TreeNode{
		parent: g.currentState,
		point:  &Point{i, j},
	}
	g.currentState = newNode
}

func sliceFromIterator(iterator func(f func(p *Point) bool)) []Point {
	s := make([]Point, 0)
	iterator(func(p *Point) bool {
		s = append(s, *p)
		return false
	})
	return s
}

func (g *Game) ServeFieldList(w http.ResponseWriter) {
	circles := sliceFromIterator(g.currentState.IterateOnePlayer)
	var crosses []Point
	if g.currentState.parent != nil {
		crosses = sliceFromIterator(g.currentState.parent.IterateOnePlayer)
	} else {
		crosses = make([]Point, 0)
	}
	res, err := json.Marshal(struct {
		Circles []Point
		Crosses []Point
	}{
		Circles: circles,
		Crosses: crosses,
	})
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		// fmt.Println(err)
	}
}

func choose(circles, crosses float64, p PlayerType) float64 {
	if p == Circles {
		return circles
	}
	if p == Crosses {
		return crosses
	}
	return 0
}

type ProbabilityOfWinReturn struct {
	circles float64
	crosses float64
	move    *TreeNode
}

func (g *Game) MakeNextMove() (newMove *Point, maxProbability float64) {
	possibleMoves := g.currentState.Expand()
	// state := g.GetBoardState()
	maxProbability = -1
	parent := g.currentState
	ch := make(chan ProbabilityOfWinReturn, len(possibleMoves))
	for _, v := range possibleMoves {
		go ProbabilityOfWin(
			&TreeNode{
				parent: parent,
				point:  &v,
			},
			g.startingPlayer,
			ch,
		)
	}
	for i := 0; i < len(possibleMoves); i++ {
		v := <-ch
		p := choose(v.circles, v.crosses, g.startingPlayer)
		fmt.Println(v.move.point, p)
		if p >= maxProbability {
			maxProbability = p
			newMove = v.move.point
		}
	}
	fmt.Println("Choosing point ", *newMove, " with probability ", maxProbability)
	newNode := &TreeNode{
		parent: g.currentState,
		point:  (*Point)(newMove),
	}
	g.currentState = newNode
	return
}

