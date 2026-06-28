package core

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/EnotInc/goch/internal/ascii"
	board "github.com/EnotInc/goch/internal/board"
	"github.com/EnotInc/goch/internal/tui"
	"golang.org/x/term"
)

const (
	hide_cursor = "\033[?25l"
	show_cursor = "\033[?25h"
)

type core struct {
	fdin   int
	fdout  int
	old    *term.State
	board  *board.Board
	cursor *cursor

	input []rune
}

func Init() *core {
	_fdin := int(os.Stdin.Fd())
	_fdout := int(os.Stdout.Fd())

	old, err := term.MakeRaw(_fdin)
	if err != nil {
		panic(err)
	}

	c := newCursor()
	b := board.NewBoard()
	return &core{
		fdin:   _fdin,
		fdout:  _fdout,
		old:    old,
		board:  b,
		cursor: c,
	}
}

const (
	enter     = 13
	space     = 32
	esc       = 27
	backspace = 127
)

func (core *core) Run() {
	tui := tui.NewTui()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(ascii.Save, hide_cursor)

	// first draw
	fen := core.board.ToFen()
	tui.Draw(fen, core.cursor.scalar(), -1, nil)

	var i = 0

	for {
		i += 1
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if len(core.input) == 0 {
			core.normal(key)
		} else {
			core.command(key)
		}

		fen := core.board.ToFen()
		fmt.Print(ascii.Reset)
		tui.SetCommand(string(core.input))
		tui.Draw(fen, core.cursor.scalar(), core.board.From(), core.board.Moves())
	}
}

func (core *core) normal(key rune) {
	switch key { // TODO: add arrows support
	case 'h':
		core.cursor.MvLeft()
	case 'j':
		core.cursor.MvDown()
	case 'k':
		core.cursor.MvUp()
	case 'l':
		core.cursor.MvRight()
	case 'q':
		core.Exit()
	case enter, space:
		core.board.AddMove(core.cursor.scalar())
	case esc:
		core.board.Cancel_selection()
	case ':':
		core.input = append(core.input, key)
		core.cursor.active = false
	}
}

func (core *core) command(key rune) {
	switch key {
	case esc:
		core.input = []rune{}
		core.cursor.active = true
	case enter:
		core.tryParce()
		core.cursor.active = true
	case backspace:
		if len(core.input) > 0 {
			core.input = core.input[:len(core.input)-1]
		}
	default:
		if len(core.input) > 8 {
			return
		}
		if slices.Contains([]rune("q1234567890abdcefghABCDEFGH"), key) {
			core.input = append(core.input, key)
		}
	}
}
