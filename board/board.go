package board

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/chrisvaughn/amazin/cells"
)

type Board struct {
	grid [][]cells.CellState
	ct   cells.CellType
	rand *rand.Rand
}

func NewBoard(width, height int, opts ...func(*Board)) *Board {
	g := make([][]cells.CellState, height)
	for i := range g {
		g[i] = make([]cells.CellState, width)
	}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := &Board{g, cells.NewSquare(), r1}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func WithCellType(c cells.CellType) func(*Board) {
	return func(b *Board) {
		b.ct = c
	}
}

func WithSeed(seed int64) func(*Board) {
	return func(b *Board) {
		b.rand = rand.New(rand.NewSource(seed))
	}
}

func (b *Board) Generate() {
	b.carve(0, 0)
}

// carve will carve a passage from the current cell at x, y
func (b *Board) carve(cx, cy int) {
	for _, d := range b.ct.Directions(b.rand) {
		nx, ny := b.ct.CellForDirection(cx, cy, d)
		if b.inBounds(nx, ny) && b.grid[ny][nx] == 0 {
			b.grid[cy][cx] |= d
			b.grid[ny][nx] |= d.Opposite()
			b.carve(nx, ny)
		}
	}

}

// carve will carve a passage from the current cell at x, y
func (b *Board) inBounds(cx, cy int) bool {
	return cx >= 0 && cx < len(b.grid[0]) && cy >= 0 && cy < len(b.grid)
}

func (b *Board) Display() {
	for y, h := range b.grid {
		fmt.Print("|")
		for x, cell := range h {
			if cell&2 != 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}
			if cell&4 != 0 {
				if (cell|b.grid[y][x+1])&2 != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}
			} else {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}
