package tui

import (
	"fmt"
	"strings"
)

const (
	reset    = "\033[0m"
	black    = "\033[42m"
	white    = "\033[43m"
	cursored = "\033[46m"

	selected = "\033[34m"
)

func (t *tui) Draw(fen string, cursor int, from int, moves []int) {
	var board strings.Builder
	pieces := t.decode(fen, moves)
	for row := range 8 {
		fmt.Fprintf(&board, "%d ", row+1)
		for col := range 8 {
			var color string
			pos := row*8 + col

			if cursor == pos {
				color = cursored
			} else {
				switch (row+col)%2 == 0 {
				case false:
					color = black
				case true:
					color = white
				}
			}

			board.WriteString(color)
			piece := pieces[row*8+col]
			if from == pos {
				board.WriteString(selected)
			}

			board.WriteString(string(piece))
			board.WriteString(reset)
		}
		board.WriteByte('\n')
	}
	board.WriteString("  a b c d e f g h")
	fmt.Print(board.String())
}
