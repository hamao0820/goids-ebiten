package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/goids-ebiten/goids"
)

func main() {
	g := goids.NewGame()
	ebiten.SetWindowSize(goids.Width, goids.Height)
	ebiten.SetWindowTitle("Goids")
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowMousePassthrough(true)
	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	if err := ebiten.RunGameWithOptions(g, op); err != nil {
		panic(err)
	}
}
