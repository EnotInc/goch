package tui

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	black = "\033[42m"
	white = "\033[43m"

	cur_black = "\033[46m"
	cur_white = "\033[46m"
)

func (t *tui) Draw(fen string, cursor int) {
	var board strings.Builder
	pieces := t.decode(fen)
	for row := range 8 {
		fmt.Fprintf(&board, "%d ", row+1)
		for col := range 8 {
			var color string
			pos := row*8 + col

			switch (row+col)%2 == 0 {
			case false:
				if cursor == pos {
					color = cur_black
				} else {
					color = black
				}

			case true:
				if cursor == pos {
					color = cur_white
				} else {
					color = white
				}
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
