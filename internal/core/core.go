package core

import (
	"bufio"
	"fmt"
	"os"

	board "github.com/EnotInc/goch/internal/board"
	"github.com/EnotInc/goch/internal/tui"
	"golang.org/x/term"
)

const (
	save  = "\033[s"
	reset = "\033[u"

	hide_cursor = "\033[?25l"
	show_cursor = "\033[?25h"

	clearline = "\033[0K"
)

type core struct {
	fdin  int
	fdout int
	old   *term.State
}

func Init() *core {
	_fdin := int(os.Stdin.Fd())
	_fdout := int(os.Stdout.Fd())

	old, err := term.MakeRaw(_fdin)
	if err != nil {
		panic(err)
	}

	return &core{
		fdin:  _fdin,
		fdout: _fdout,
		old:   old,
	}
}

const (
	enter = 13
	space = 32
	esc   = 27
)

func (core *core) Run() {
	tui := tui.NewTui()
	b := board.NewBoard()
	c := newCursor()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(save, hide_cursor)

	// first draw
	fen := b.ToFen()
	tui.Draw(fen, c.scalar(), -1, nil)

	var i = 0

	var quit = false
	for !quit {
		i += 1
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		switch key { // TODO: add arrows support
		case 'h':
			c.MvLeft()
		case 'j':
			c.MvDown()
		case 'k':
			c.MvUp()
		case 'l':
			c.MvRight()
		case 'q':
			quit = true
		case enter, space:
			b.AddMove(c.scalar())
		case esc:
			b.Cancel_selection()
		}

		fen := b.ToFen()
		fmt.Print(reset)
		tui.Draw(fen, c.scalar(), b.From(), b.Moves())
	}

	fmt.Print(show_cursor)
}
