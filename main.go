package main

import (
	"github.com/chrisvaughn/amazin/board"
)

func main() {
	board := board.NewBoard(20, 20)
	board.Generate()
	board.Display()
}
