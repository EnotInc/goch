package core

import (
	"fmt"
	"os"

	"github.com/EnotInc/goch/internal/ascii"
)

const clear = "\033[0J"

func (core *core) Exit() {
	fmt.Print(ascii.Reset)
	fmt.Print(clear)
	fmt.Print(show_cursor)
	os.Exit(0)
}
