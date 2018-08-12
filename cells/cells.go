package cells

import (
	"math/rand"
)

const (
	N CellState = 1
	S CellState = 2
	E CellState = 4
	W CellState = 8
)

func (s CellState) Opposite() CellState {
	switch s {
	case N:
		return S
	case S:
		return N
	case E:
		return W
	case W:
		return E
	}
	return 0
}

func shuffle(r *rand.Rand, vals []CellState) {
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

type CellState int

type CellType interface {
	Directions(r *rand.Rand) []CellState
	CellForDirection(cx, cy int, d CellState) (int, int)
}

type Square struct {
	directions []CellState
}

func NewSquare() Square {
	return Square{[]CellState{N, S, E, W}}
}

func (s Square) Directions(r *rand.Rand) []CellState {
	shuffle(r, s.directions)
	return s.directions
}

func (s Square) CellForDirection(cx, cy int, d CellState) (int, int) {
	var dx, dy int
	switch d {
	case N:
		dx, dy = 0, -1
	case S:
		dx, dy = 0, 1
	case E:
		dx, dy = 1, 0
	case W:
		dx, dy = -1, 0
	}
	return cx + dx, cy + dy
}
