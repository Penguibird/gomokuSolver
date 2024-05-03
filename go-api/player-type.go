package main

type PlayerType int

const (
	Null PlayerType = iota
	Crosses
	Circles
)

func (s PlayerType) String() string {
	switch s {
	case Circles:
		return "Circles"
	case Null:
		return "Null"
	case Crosses:
		return "Crosses"
	}
	return "unknown"
}

func (t PlayerType) Toggle() PlayerType {
	switch t {
	case Circles:
		return Crosses
	case Crosses:
		return Circles
	case Null:
		return Null
	}
	return Null
}
