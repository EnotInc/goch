package board

func (b *board) AddMove(cursor int) {
	if b.from == 0 {
		b.from = 1 << cursor
	} else {
		b.to = 1 << cursor
		b.performMove()
	}
}

// build moves board map
// shows every possible position where piece can go to
func (b *board) setMoves() {
}

func (b *board) performMove() {
	b.setMoves()
	b.from = 0
	b.to = 0
}

func (b *board) From() int {
	for i := range 64 {
		var mask uint64 = 1 << i
		if mask&b.from != 0 {
			return i
		}

	}
	return -1
}

func (b *board) Moves() []int {
	var moves []int = nil
	for i := range 64 {
		var mask uint64 = 1 << i
		if mask&b.moves != 0 {
			moves = append(moves, i)
		}

	}
	return moves
}
