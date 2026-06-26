package tui

import (
	"fmt"
	"strconv"
)

type tui struct {
	board   string
	message string
}

func NewTui() *tui {
	return &tui{}
}

func (t *tui) decode(fen string, moves []int) []piece {
	var pieces []piece
	rows := 7
	col := 8

	for _, ch := range fen {
		amount := 0
		switch ch {
		// -==[ rows ]==-
		case '/':
			rows -= 1
			col = 8
			continue

		// -==[ black pieces ]==-
		case 'k':
			pieces = append(pieces, bking)
		case 'q':
			pieces = append(pieces, bqueen)
		case 'r':
			pieces = append(pieces, brook)
		case 'n':
			pieces = append(pieces, bknight)
		case 'b':
			pieces = append(pieces, bbishop)
		case 'p':
			pieces = append(pieces, bpawn)

		// -==[ white pieces ]==-
		case 'K':
			pieces = append(pieces, wking)
		case 'Q':
			pieces = append(pieces, wqueen)
		case 'R':
			pieces = append(pieces, wrook)
		case 'N':
			pieces = append(pieces, wknight)
		case 'B':
			pieces = append(pieces, wbishop)
		case 'P':
			pieces = append(pieces, wpawn)

		// -==[ numbers ]==-
		default:
			a, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(fmt.Sprintf("Unable to decode fen string. Unknown char: %c\nfen: %s\nerr: %s", ch, fen, err))
			}
			for range a {
				pieces = append(pieces, none)
			}
			amount = a
		}
		col -= amount

		if col < 0 {
			panic(fmt.Sprintf("Unable to decode fen string: %s", fen))
		}
	}

	if rows < 0 {
		panic(fmt.Sprintf("Unable to decode fen string.\nWrong rows amount.\nfen: %s", fen))
	}

	for _, m := range moves {
		pieces[m] = move
	}

	return pieces
}
