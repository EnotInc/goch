package board

type piece_t int

const (
	_ piece_t = iota
	king
	queen
	rook
	bishop
	knight
	pawn

	nat // not a type
)

type side int

const (
	_ side = iota
	black
	white

	none
)

type piece struct {
	piece_t piece_t
	side    side
}

func (b *Board) select_piece() {
	p := piece{}

	if b.from&b.bpieces != 0 {
		p.side = black
		if b.from&b.bking != 0 {
			p.piece_t = king
		} else if b.from&b.bqueen != 0 {
			p.piece_t = queen
		} else if b.from&b.brook != 0 {
			p.piece_t = rook
		} else if b.from&b.bbishop != 0 {
			p.piece_t = bishop
		} else if b.from&b.bknight != 0 {
			p.piece_t = knight
		} else if b.from&b.bpawn != 0 {
			p.piece_t = pawn
		}
	} else if b.from&b.wpieces != 0 {
		p.side = white
		if b.from&b.wking != 0 {
			p.piece_t = king
		} else if b.from&b.wqueen != 0 {
			p.piece_t = queen
		} else if b.from&b.wrook != 0 {
			p.piece_t = rook
		} else if b.from&b.wbishop != 0 {
			p.piece_t = bishop
		} else if b.from&b.wknight != 0 {
			p.piece_t = knight
		} else if b.from&b.wpawn != 0 {
			p.piece_t = pawn
		}
	} else {
		p.side = none
		p.piece_t = nat
	}

	b.moved_piese = p
}
