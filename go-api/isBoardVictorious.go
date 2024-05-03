package main

import (
	"sync"
)

type Board struct {
	board *[][]Field
}

var directions = &[8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	// {0, 1},
	// {1, -1},
	// {1, 0},
	// {1, 1},
}

type IsBoardVictoriousReturn struct {
	isGameOver bool
	PlayerType
}

var victoriousBoardsHashMap = &sync.Map{} // make(map[BoardHash]*IsBoardVictoriousReturn)

func IsBoardVictorious(tree *TreeNode) (isGameOver bool, winner bool) {
	// h := HashBoard(&Board{board: b}, Null)
	// v, ok := victoriousBoardsHashMap.Load(h)
	// if ok {
	// 	r := v.(IsBoardVictoriousReturn)
	// 	// fmt.Println("Cache hit")
	// 	return r.isGameOver, r.PlayerType
	// } else {
	// 	defer func() {
	// 		victoriousBoardsHashMap.Store(h, IsBoardVictoriousReturn{isGameOver: isGameOver, PlayerType: winner})
	// 	}()
	// }

	i := 0
	tree.IterateAllFields(func(p *Point) bool {
		i++
		return false
	})
	if i >= (RowCount * ColCount) {
		return true, false
	}

	tree.IterateOnePlayer(func(p *Point) bool {
		isGameOver = check(tree, p)
		if isGameOver {
			winner = true
			return true
		}
		return false
	})

	return
}

// f field to compare to - to check circles we'd pass Circle
func check(t *TreeNode, p *Point) bool {
	ch := make(chan bool, 8)
	for _, dir := range directions {
		go checkDirection(t, p, dir, ch)
	}
	for i := 0; i < len(directions); i++ {
		v, ok := <-ch
		if v == true || !ok {
			close(ch)
			return true
		}
	}
	return false
}

func checkDirection(t *TreeNode, p *Point, dir [2]int, ch chan bool) {
	for counter := 1; counter < 3; counter++ {
		i := p[0] + (dir[0] * counter)
		j := p[1] + (dir[1] * counter)
		if i < 0 || j < 0 || i >= RowCount || j >= ColCount {
			ch <- false
			return
		}
		matchFound := false
		pointToCheck := [2]int{i, j}
		t.IterateOnePlayer(func(p *Point) bool {
			if *p == pointToCheck {
				matchFound = true
				//break loop
				return true
			}
			//continue iterating
			return false
		})
		if !matchFound {
			ch <- false
		} else {
			ch <- true
		}
	}
}
