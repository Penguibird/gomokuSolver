package main;

type Field int;

const (
	Empty Field = iota;
	Cross
	Circle
)
func (s Field) String() string {
	switch s {
	case Empty:
		return " "
	case Cross:
		return "❌"
	case Circle:
		return "⭕"
	}
	return "unknown"
}