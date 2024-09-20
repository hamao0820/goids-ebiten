package goids

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/goids-ebiten/gopher"
	"github.com/hamao0820/goids-ebiten/vector"
)

const (
	goidsNum = 20
)

var (
	Width, Height int
)

func init() {
	Width, Height = ebiten.ScreenSizeInFullscreen()
}

type Game struct {
	goids []Goid
}

func NewGame() *Game {
	goids := make([]Goid, 0, goidsNum)
	for i := 0; i < goidsNum; i++ {
		goids = append(goids, NewGoid(vector.New(float64(rand.Intn(Width)), float64(rand.Intn(Height))), 2, 0.1, 100))
	}
	return &Game{goids: goids}
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	mouse := vector.New(float64(x), float64(y))
	for i := 0; i < len(g.goids); i++ {
		goid := &g.goids[i]
		goid.Flock(g.goids, mouse)
		goid.Update(float64(Width), float64(Height))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, goid := range g.goids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(GopherSize)/2, -float64(GopherSize)/2)
		op.GeoM.Translate(goid.position.X, goid.position.Y)
		switch goid.imageType {
		case Front:
			screen.DrawImage(gopher.Front, op)
		case Side:
			screen.DrawImage(gopher.Side, op)
		case Pink:
			screen.DrawImage(gopher.Pink, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}
