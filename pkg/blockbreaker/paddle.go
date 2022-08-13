package blockbreaker

import "log"

type Paddle struct {
	X, Y    int
	XOffset int //NOTE: X represents the middle of the paddle
	Form    []rune
}

func NewPaddle(width int) (p Paddle) {
	if width%2 == 0 {
		log.Fatal("Paddle must have an uneven size.")
	}
	p.XOffset = width / 2
	for i := 0; i < width; i++ {
		p.Form = append(p.Form, '\u2583')
	}
	return
}
