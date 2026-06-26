package main

import (
	board "github.com/EnotInc/goch/internal/board"
	"github.com/EnotInc/goch/internal/tui"
)

func main() {
	tui := tui.NewTui()

	b := board.NewBoard()
	fen := b.ToFen()

	tui.Draw(fen)
}
