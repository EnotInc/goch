package board

const (
	black_pawn_home = 0b0000000011111111000000000000000000000000000000000000000000000000
	white_pawn_home = 0b0000000000000000000000000000000000000000000000001111111100000000

	upper_border  uint64 = 0b1111111100000000000000000000000000000000000000000000000000000000
	bottom_border uint64 = 0b0000000000000000000000000000000000000000000000000000000011111111
	left_border   uint64 = 0b1000000010000000100000001000000010000000100000001000000010000000
	right_border  uint64 = 0b0000000100000001000000010000000100000001000000010000000100000001
)

func (b *Board) generalBorderMask() uint64 {
	var mask uint64 = 0b1111111111111111111111111111111111111111111111111111111111111111

	if b.from&left_border != 0 {
		mask ^= right_border
	}
	if b.from&right_border != 0 {
		mask ^= left_border
	}

	return mask
}

func (b *Board) knightBorderMask() uint64 {
	var mask uint64 = 0b1111111111111111111111111111111111111111111111111111111111111111

	if b.from<<1&left_border != 0 {
		mask ^= right_border
	}
	if b.from>>1&right_border != 0 {
		mask ^= left_border
	}
	if b.from&left_border != 0 {
		mask ^= right_border
		mask ^= b.from >> 6
		mask ^= b.from << 10
	}
	if b.from&right_border != 0 {
		mask ^= left_border
		mask ^= b.from << 6
		mask ^= b.from >> 10
	}

	return mask
}

func (b *Board) pawnMove() {
	// TODO: implement en passant
	// TODO: implement promotion
	p := b.moved_piese
	switch p.side {
	case white:
		b.moves |= b.from << 8 & b.empty
		if (b.from&white_pawn_home != 0) && (b.from<<8&b.empty != 0) {
			b.moves |= b.from << 16 & b.empty
		}
		b.moves |= b.from << 7 & b.bpieces & b.generalBorderMask()
		b.moves |= b.from << 9 & b.bpieces & b.generalBorderMask()
	case black:
		b.moves |= b.from >> 8 & b.empty
		if (b.from&black_pawn_home != 0) && (b.from>>8&b.empty != 0) {
			b.moves |= b.from >> 16 & b.empty
		}
		b.moves |= b.from >> 7 & b.wpieces & b.generalBorderMask()
		b.moves |= b.from >> 9 & b.wpieces & b.generalBorderMask()
	}

}

func (b *Board) kingMove() {
	// TODO: implement castling
	b.moves |= (b.from << 1) & b.generalBorderMask()
	b.moves |= (b.from >> 1) & b.generalBorderMask()
	b.moves |= (b.from << 8) & b.generalBorderMask()
	b.moves |= (b.from >> 8) & b.generalBorderMask()
	b.moves |= (b.from << 7) & b.generalBorderMask()
	b.moves |= (b.from >> 7) & b.generalBorderMask()
	b.moves |= (b.from << 9) & b.generalBorderMask()
	b.moves |= (b.from >> 9) & b.generalBorderMask()
}

func (b *Board) knightMove() {
	b.moves |= (b.from << 17) & b.knightBorderMask()
	b.moves |= (b.from >> 17) & b.knightBorderMask()
	b.moves |= (b.from << 15) & b.knightBorderMask()
	b.moves |= (b.from >> 15) & b.knightBorderMask()
	b.moves |= (b.from << 10) & b.knightBorderMask()
	b.moves |= (b.from >> 10) & b.knightBorderMask()
	b.moves |= (b.from << 6) & b.knightBorderMask()
	b.moves |= (b.from >> 6) & b.knightBorderMask()
}

func (b *Board) rookMove(enemy uint64) {
	b.scanPlus(enemy)
}

func (b *Board) bishopMove(enemy uint64) {
	b.scanCross(enemy)
}

func (b *Board) queenMove(enemy uint64) {
	b.scanCross(enemy)
	b.scanPlus(enemy)
}

// x
func (b *Board) scanCross(enemy uint64) {
	const tlbr = 9 // topleft - topright
	const trbl = 7 // botleft - botright

	tl := b.from
	br := b.from
	tr := b.from
	bl := b.from

	moved := true
	for moved {
		moved = false

		// -==[ moving white there is a free celss ]==-
		if tl&upper_border == 0 && tl&left_border == 0 && tl<<tlbr&b.empty != 0 {
			tl = tl << tlbr
			b.moves |= tl
			moved = true
		}
		if br&bottom_border == 0 && br&right_border == 0 && br>>tlbr&b.empty != 0 {
			br = br >> tlbr
			b.moves |= br
			moved = true
		}
		if tr&upper_border == 0 && tr&right_border == 0 && tr<<trbl&b.empty != 0 {
			tr = tr << trbl
			b.moves |= tr
			moved = true
		}
		if bl&bottom_border == 0 && bl&left_border == 0 && bl>>trbl&b.empty != 0 {
			bl = bl >> trbl
			b.moves |= bl
			moved = true
		}

		// -==[ capturing enemies pieces ]==-
		if tl<<tlbr&enemy != 0 && tl<<tlbr&right_border == 0 {
			b.moves |= tl << tlbr
		}
		if br>>tlbr&enemy != 0 && br>>tlbr&left_border == 0 {
			b.moves |= br >> tlbr
		}
		if tr<<trbl&enemy != 0 && tr<<trbl&left_border == 0 {
			b.moves |= tr << trbl
		}
		if bl>>trbl&enemy != 0 && bl>>trbl&right_border == 0 {
			b.moves |= bl >> trbl
		}
	}

}

// +
func (b *Board) scanPlus(enemy uint64) {
	const lr = 1 // left - right
	const ud = 8 // up - down

	l := b.from
	r := b.from
	u := b.from
	d := b.from

	moved := true
	for moved {
		moved = false
		// -==[ moving white there is a free celss ]==-

		if l&left_border == 0 && l<<lr&b.empty != 0 {
			l = l << lr
			b.moves |= l
			moved = true
		}
		if r&right_border == 0 && r>>lr&b.empty != 0 {
			r = r >> lr
			b.moves |= r
			moved = true
		}
		if u&upper_border == 0 && u<<ud&b.empty != 0 {
			u = u << ud
			b.moves |= u
			moved = true
		}
		if d&bottom_border == 0 && d>>ud&b.empty != 0 {
			d = d >> ud
			b.moves |= d
			moved = true
		}

		// -==[ capturing enemies pieces ]==-
		if l<<lr&enemy != 0 {
			b.moves |= l << lr
		}
		if r>>lr&enemy != 0 {
			b.moves |= r >> lr
		}
		if u<<ud&enemy != 0 {
			b.moves |= u << ud
		}
		if d>>ud&enemy != 0 {
			b.moves |= d >> ud
		}
	}
}
