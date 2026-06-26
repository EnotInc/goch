package internal

import (
	"fmt"
	"strings"
)

type board struct {
	wking   uint64
	wqueen  uint64
	wrook   uint64
	wbishop uint64
	wknight uint64
	wpawn   uint64

	bking   uint64
	bqueen  uint64
	brook   uint64
	bbishop uint64
	bknight uint64
	bpawn   uint64

	empty   uint64
	taken   uint64
	wpieces uint64
	bpieces uint64
}

func NewBoard() *board {
	b := &board{
		bking:   0b0001000000000000000000000000000000000000000000000000000000000000,
		bqueen:  0b0000100000000000000000000000000000000000000000000000000000000000,
		bbishop: 0b0010010000000000000000000000000000000000000000000000000000000000,
		bknight: 0b0100001000000000000000000000000000000000000000000000000000000000,
		brook:   0b1000000100000000000000000000000000000000000000000000000000000000,
		bpawn:   0b0000000011111111000000000000000000000000000000000000000000000000,

		wking:   0b0000000000000000000000000000000000000000000000000000000000010000,
		wqueen:  0b0000000000000000000000000000000000000000000000000000000000001000,
		wbishop: 0b0000000000000000000000000000000000000000000000000000000000100100,
		wknight: 0b0000000000000000000000000000000000000000000000000000000001000010,
		wrook:   0b0000000000000000000000000000000000000000000000000000000010000001,
		wpawn:   0b0000000000000000000000000000000000000000000000001111111100000000,
	}
	b.updatePieces()

	return b
}

func (b *board) updatePieces() {
	b.bpieces = b.bbishop | b.bking | b.bknight | b.bpawn | b.bqueen | b.brook
	b.wpieces = b.wbishop | b.wking | b.wknight | b.wpawn | b.wqueen | b.brook

	b.taken = b.bpieces | b.wpieces
	b.empty = ^b.taken
}

func (b *board) ToFen() string {
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
		} else if i%8 == 0 && i > 0 {
			writeEmpty()
			fen.WriteByte('/')
		}
	}

	return fen.String()
}
