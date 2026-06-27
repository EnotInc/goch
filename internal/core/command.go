package core

import (
	"slices"
	"strconv"
)

func (core *core) tryParce() {
	if len(core.input) > 1 {
		core.input = core.input[1:]
	}

	if len(core.input) == 2 && slices.Contains([]rune("abcdefghABCDEFGH"), core.input[0]) {
		col := -1
		switch core.input[0] {
		case 'a', 'A':
			col = 0
		case 'b', 'B':
			col = 1
		case 'c', 'C':
			col = 2
		case 'd', 'D':
			col = 3
		case 'e', 'E':
			col = 4
		case 'f', 'F':
			col = 5
		case 'g', 'G':
			col = 6
		case 'h', 'H':
			col = 7
		}

		row, err := strconv.Atoi(string(core.input[1]))
		if err != nil || row == -1 || col > 7 || col < 0 {
			core.input = []rune{}
			return
		}

		row -= 1
		core.cursor.SetPos(row, col)
	}

	core.input = []rune{}
}
