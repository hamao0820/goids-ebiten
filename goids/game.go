package goids

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Width  = 640
	Height = 480
)

var (
	gopher *ebiten.Image
)

func init() {
	var err error
	gopher, _, err = ebitenutil.NewImageFromFile("assets/images/gopher.png")
	if err != nil {
		panic(err)
	}
}

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(gopher, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}
