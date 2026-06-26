package board

func (b *board) AddMove(cursor int) {
	if cursor == -1 {
		return
	}

	if b.from == 0 {
		b.from = 1 << (63 - cursor)
		b.setMoves()
	} else {
		b.to = 1 << (63 - cursor)
		b.performMove()
	}
}

// build moves board map
// shows every possible position where piece can go to
func (b *board) setMoves() {
	b.updatePieces()
	b.select_piece()
	p := b.moved_piese

	var enemy_pieces uint64
	var your_pieces uint64

	switch p.side {
	case white:
		your_pieces = b.wpieces
		enemy_pieces = b.bpieces
	case black:
		your_pieces = b.bpieces
		enemy_pieces = b.wpieces
	case none:
		return
	}

	switch p.piece_t {
	case king:
		b.kingMove()
	case queen:
		b.queenMove(enemy_pieces)
	case rook:
		b.rookMove(enemy_pieces)
	case bishop:
		b.bishopMove(enemy_pieces)
	case knight:
		b.knightMove()
	case pawn:
		b.pawnMove()
	case nat:
		return // TODO: set error message
	}

	b.moves &= ^your_pieces
}

func (b *board) performMove() {
	if (b.moves&b.to == 0) || (b.from&b.to != 0) {
		b.Cancel_selection()
		return
	}

	b.capture()
	p := b.moved_piese
	switch p.side {
	case white:
		switch p.piece_t {
		case king:
			b.wking ^= b.from
			b.wking |= b.to
		case queen:
			b.wqueen ^= b.from
			b.wqueen |= b.to
		case rook:
			b.wrook ^= b.from
			b.wrook |= b.to
		case bishop:
			b.wbishop ^= b.from
			b.wbishop |= b.to
		case knight:
			b.wknight ^= b.from
			b.wknight |= b.to
		case pawn:
			b.wpawn ^= b.from
			b.wpawn |= b.to
		}

	case black:
		switch p.piece_t {
		case king:
			b.bking ^= b.from
			b.bking |= b.to
		case queen:
			b.bqueen ^= b.from
			b.bqueen |= b.to
		case rook:
			b.brook ^= b.from
			b.brook |= b.to
		case bishop:
			b.bbishop ^= b.from
			b.bbishop |= b.to
		case knight:
			b.bknight ^= b.from
			b.bknight |= b.to
		case pawn:
			b.bpawn ^= b.from
			b.bpawn |= b.to
		}

	}

	b.Cancel_selection()
	b.updatePieces()
}

func (b *board) Cancel_selection() {
	b.from = 0
	b.to = 0
	b.moves = 0
	b.moved_piese = piece{
		piece_t: nat,
		side:    none,
	}

}

func (b *board) capture() {
	// TODO: keep track of captured pieces

	// -==[ black pieces ]==-
	if b.to&b.bking != 0 {
		b.bking ^= b.to
	} else if b.to&b.bqueen != 0 {
		b.bqueen ^= b.to
	} else if b.to&b.bbishop != 0 {
		b.bbishop ^= b.to
	} else if b.to&b.bknight != 0 {
		b.bknight ^= b.to
	} else if b.to&b.brook != 0 {
		b.brook ^= b.to
	} else if b.to&b.bpawn != 0 {
		b.bpawn ^= b.to
	} else

	// -==[ white pieces ]==-
	if b.to&b.wking != 0 {
		b.wking ^= b.to
	} else if b.to&b.wqueen != 0 {
		b.wqueen ^= b.to
	} else if b.to&b.wbishop != 0 {
		b.wbishop ^= b.to
	} else if b.to&b.wknight != 0 {
		b.wknight ^= b.to
	} else if b.to&b.wrook != 0 {
		b.wrook ^= b.to
	} else if b.to&b.wpawn != 0 {
		b.wpawn ^= b.to
	}
}

func (b *board) From() int {
	for i := range 64 {
		var mask uint64 = 1 << (63 - i)
		if mask&b.from != 0 {
			return i
		}

	}
	return -1
}

func (b *board) Moves() []int {
	var moves []int = nil
	for i := range 64 {
		var mask uint64 = 1 << (63 - i)
		if mask&b.moves != 0 {
			moves = append(moves, i)
		}

	}
	return moves
}
