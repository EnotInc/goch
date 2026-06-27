package board

import (
	"fmt"
	"strings"
)

func (b *Board) ToFen() string {
	var fen strings.Builder
	var empty int = 0

	writeEmpty := func() {
		if empty != 0 {
			fmt.Fprintf(&fen, "%d", empty)
			empty = 0
		}
	}

	for i := 63; i >= 0; i-- {
		var mask uint64 = 1 << i

		// -==[ black pieces ]==-
		if mask&b.bking != 0 {
			writeEmpty()
			fen.WriteByte('k')
		} else if mask&b.bqueen != 0 {
			writeEmpty()
			fen.WriteByte('q')
		} else if mask&b.bbishop != 0 {
			writeEmpty()
			fen.WriteByte('b')
		} else if mask&b.bknight != 0 {
			writeEmpty()
			fen.WriteByte('n')
		} else if mask&b.brook != 0 {
			writeEmpty()
			fen.WriteByte('r')
		} else if mask&b.bpawn != 0 {
			writeEmpty()
			fen.WriteByte('p')
		} else

		// -==[ white pieces ]==-
		if mask&b.wking != 0 {
			writeEmpty()
			fen.WriteByte('K')
		} else if mask&b.wqueen != 0 {
			writeEmpty()
			fen.WriteByte('Q')
		} else if mask&b.wbishop != 0 {
			writeEmpty()
			fen.WriteByte('B')
		} else if mask&b.wknight != 0 {
			writeEmpty()
			fen.WriteByte('N')
		} else if mask&b.wrook != 0 {
			writeEmpty()
			fen.WriteByte('R')
		} else if mask&b.wpawn != 0 {
			writeEmpty()
			fen.WriteByte('P')

		} else {
			empty += 1
		}

		// -==[ new line ]==-
		if empty == 8 {
			fmt.Fprintf(&fen, "%d", empty)
			fen.WriteByte('/')
			empty = 0
		} else if i%8 == 0 {
			writeEmpty()
			if i > 0 { // not including last '/' symbol
				fen.WriteByte('/')
			}
		}
	}

	return fen.String()
}
