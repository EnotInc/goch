package tui

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	black = "\033[42m"
	white = "\033[43m"

	save    = "\033[s"
	restore = "\033[u"
	back    = "\033[1A"
	clear   = "\033[2K"
)

func (t *tui) Draw(fen string) {
	var board strings.Builder
	pieces := t.decode(fen)
	for row := range 8 {
		fmt.Fprintf(&board, "%d ", row+1)
		for col := range 8 {
			color := black
			if (row+col)%2 == 0 {
				color = white
			}

			piece := pieces[row*8+col]

			board.WriteString(color)
			board.WriteString(string(piece))
			board.WriteString(reset)
		}
		board.WriteByte('\n')
	}
	board.WriteString("  a b c d e f g h")
	fmt.Print(board.String())
}
