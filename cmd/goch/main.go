package main

import "github.com/EnotInc/goch/internal/tui"

const simplefen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

func main() {
	tui := tui.NewTui()
	tui.Draw(simplefen)
}
