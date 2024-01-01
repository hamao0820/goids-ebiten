package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/goids-ebiten/goids"
)

func main() {
	g := goids.NewGame()
	ebiten.SetWindowSize(goids.Width, goids.Height)
	ebiten.SetWindowTitle("Goids")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
