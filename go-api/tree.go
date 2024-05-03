package main

type TreeNode struct {
	parent *TreeNode
	point  *Point
}

func (t *TreeNode) Walk(f func(t *TreeNode)) {
	f(t)
	if t.parent != nil {
		t.parent.Walk(f)
	}
}

func (t *TreeNode) IterateAllFields(f func(p *Point) bool) {
	if t == nil || t.point == nil {
		return
	}
	b := f(t.point)
	if b {
		return
	}
	if t.parent != nil {
		t.parent.IterateAllFields(f)
	}
}

func (t *TreeNode) IterateOnePlayer(f func(p *Point) bool) {
	if t == nil || t.point == nil {
		return
	}
	b := f(t.point)
	if b {
		return
	}
	if t.parent != nil && t.parent.parent != nil && t.parent.parent.point != nil {
		t.parent.parent.IterateOnePlayer(f)
	}
}

func (t *TreeNode) Expand() []Point {
	board := MakeEfficientBoard()
	t.IterateAllFields(func(p *Point) bool {
		board.AddPoint(p[0], p[1])
		return false
	})

	points := make([]Point, 0)
	for i := 0; i < RowCount; i++ {
		for j := 0; j < ColCount; j++ {
			if !board.HasPoint(i, j) {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}
