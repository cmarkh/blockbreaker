package blockbreaker

type Ball struct {
	X      int
	Y      int
	XSpeed int
	YSpeed int
	Form   rune
}

func (b *Ball) Move() {
	b.X += b.XSpeed
	b.Y += b.YSpeed
}

func (b *Ball) CheckEdges(size Size) {
	if b.X <= size.StartX || b.X >= size.Width {
		b.XSpeed *= -1
	}
	if b.Y <= size.StartY || b.Y >= size.Height {
		b.YSpeed *= -1
	}
}
