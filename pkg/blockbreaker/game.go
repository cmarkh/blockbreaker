package blockbreaker

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Style  tcell.Style
	Size   Size
	Round  int
	Speed  time.Duration //sets the speed of the game loop. each loop contains one ball move so faster loop = faster ball
	Ball   Ball
	Paddle Paddle
}

type Size struct {
	Width, Height  int
	StartX, StartY int
}

func Start() (err error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return
	}
	err = screen.Init()
	if err != nil {
		return
	}

	game := NewGame(screen)
	go game.Loop()

	for x := -1; ; x++ {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			tcell.NewEventResize(10, 20)
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func (g *Game) Loop() {

	ticker := time.NewTicker(g.Speed)

	for x := 0; ; x++ {
		g.Screen.Clear()

		g.Borders()
		g.Screen.SetContent(g.Size.StartX, g.Size.Height+1,
			[]rune(fmt.Sprintf("Round %d", g.Round))[0], []rune(fmt.Sprintf("Round %d", g.Round))[1:],
			g.Style)
		g.Screen.SetContent(g.Size.Width+1-len(fmt.Sprint(x+1)), g.Size.Height+1,
			[]rune(fmt.Sprint(x + 1))[0], []rune(fmt.Sprint(x + 1))[1:], g.Style)

		g.Ball.CheckEdges(g.Size) //g.Screen.Size()
		g.Ball.Move()
		g.Screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Form, nil, g.Style)

		g.Screen.SetContent(g.Paddle.X, g.Paddle.Y, g.Paddle.Form[0], g.Paddle.Form[1:], g.Style)

		g.Screen.Show()

		time.Sleep(time.Second / 2)

		<-ticker.C //maintain max speed of game loop
	}
}

func NewGame(s tcell.Screen) (game *Game) {
	game = &Game{
		Screen: s,
		Style:  tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault),
		Size:   Size{80, 25, 0, 0},
		Round:  1,
		Speed:  time.Millisecond * 40,
		Ball: Ball{
			XSpeed: 1,
			YSpeed: 1,
			Form:   '\u25CF',
		},
		Paddle: Paddle{
			Form: [3]Rune{'\u0258', '\u0258', '\u0258'},
		},
	}
	game.Ball.X = game.Size.Width / 2
	game.Ball.Y = game.Size.Height - 1

	game.Paddle.X = game.Size.Width / 2
	game.Paddle.Y = game.Size.Height - 1

	game.Screen.SetStyle(game.Style)

	return
}

func (g *Game) NewRound() {
	g.Round++

	g.Speed -= time.Millisecond * time.Duration(g.Round)

	g.Ball.X = (g.Size.Width)/2 - 1
	g.Ball.Y = (g.Size.Height)
}

func (g *Game) Borders() {
	// Draw borders
	for col := g.Size.StartX; col <= g.Size.Width; col++ {
		g.Screen.SetContent(col, g.Size.StartY, tcell.RuneHLine, nil, g.Style)
		g.Screen.SetContent(col, g.Size.Height, tcell.RuneHLine, nil, g.Style)
	}
	for row := g.Size.StartY; row < g.Size.Height; row++ {
		g.Screen.SetContent(g.Size.StartY, row, tcell.RuneVLine, nil, g.Style)
		g.Screen.SetContent(g.Size.Width, row, tcell.RuneVLine, nil, g.Style)
	}

	g.Screen.SetContent(g.Size.StartX, g.Size.StartY, tcell.RuneULCorner, nil, g.Style)
	g.Screen.SetContent(g.Size.Width, g.Size.StartY, tcell.RuneURCorner, nil, g.Style)
	g.Screen.SetContent(g.Size.StartX, g.Size.Height, tcell.RuneLLCorner, nil, g.Style)
	g.Screen.SetContent(g.Size.Width, g.Size.Height, tcell.RuneLRCorner, nil, g.Style)
}
