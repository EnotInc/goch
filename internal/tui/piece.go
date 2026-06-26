package tui

type piece string

const (
	wking   piece = "\033[1m\u2654 "
	wqueen  piece = "\033[1m\u2655 "
	wrook   piece = "\033[1m\u2656 "
	wbishop piece = "\033[1m\u2657 "
	wknight piece = "\033[1m\u2658 "
	wpawn   piece = "\033[1m\u2659 "

	bking   piece = "\033[1m\u265a "
	bqueen  piece = "\033[1m\u265b "
	brook   piece = "\033[1m\u265c "
	bbishop piece = "\033[1m\u265d "
	bknight piece = "\033[1m\u265e "
	bpawn   piece = "\033[1m\u265f "

	none piece = "  "
)
