package core

type cursor struct {
	row    int
	col    int
	active bool
}

func newCursor() *cursor {
	return &cursor{
		row:    0,
		col:    0,
		active: true,
	}
}

func (c *cursor) scalar() int {
	if !c.active {
		return -1
	}
	return c.row*8 + c.col
}

func (c *cursor) MvUp() {
	if c.row > 0 {
		c.row -= 1
	}
}

func (c *cursor) MvDown() {
	if c.row < 7 {
		c.row += 1
	}
}

func (c *cursor) MvLeft() {
	if c.col > 0 {
		c.col -= 1
	}
}

func (c *cursor) MvRight() {
	if c.col < 7 {
		c.col += 1
	}
}
