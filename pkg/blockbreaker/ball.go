package blockbreaker

import "fmt"

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

func (g *Game) CheckEdges() {
	if g.Ball.X <= g.Size.StartX || g.Ball.X >= g.Size.Width {
		g.Ball.XSpeed *= -1
	}
	if g.Ball.Y <= g.Size.StartY || g.Ball.Y >= g.Size.Height {
		g.Ball.YSpeed *= -1
	}
	if g.Ball.Y == g.Paddle.Y {
		if g.Ball.X <= g.Paddle.X+g.Paddle.XOffset &&
			g.Ball.X >= g.Paddle.X-g.Paddle.XOffset {
			g.Ball.YSpeed *= -1
		}
	}
}

func (g *Game) GameOver() bool {
	if g.Ball.Y == g.Size.Height {
		str := "You Lose!"
		g.Screen.SetContent((g.Size.Width-len(str))/2, g.Size.Height/2-1,
			[]rune(str)[0], []rune(str)[1:], g.Style)
		str = "Type \"new\" to start a new game"
		g.Screen.SetContent((g.Size.Width-len(str))/2, g.Size.Height/2,
			[]rune(str)[0], []rune(str)[1:], g.Style)

		g.Screen.Show()
		fmt.Scanln()
		*g = *NewGame(g.Screen)
		return true
	}
	return false
}
