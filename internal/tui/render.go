package tui

import (
	"fmt"
	"strings"

	"github.com/EnotInc/goch/internal/ascii"
)

const (
	reset    = "\033[0m"
	black    = "\033[42m"
	white    = "\033[43m"
	cursored = "\033[46m"

	selected = "\033[35m"
)

func (t *tui) Draw(fen string, cursor int, from int, moves []int) {
	var board strings.Builder
	board.WriteString(ascii.Reset)
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
		board.WriteString("\n\r")
	}
	board.WriteString("  a b c d e f g h\n\r")
	board.WriteString(ascii.Clearline)

	if len(t.command) > 0 {
		underline := "\033[4m"
		overline := "\033[53m"
		cmd := strings.Split(t.command, ":")
		fill := strings.Repeat(" ", 16-len(t.command)) // 16 - board width
		fmt.Fprintf(&board, " %s%s: %s%s", underline, overline, cmd[1], fill)
	}

	fmt.Print(board.String())
}
