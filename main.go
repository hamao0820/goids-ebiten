package main

import "github.com/hamao0820/goids-ebiten/game"

func main() {
	g := game.New()
	if err := g.Run(); err != nil {
		panic(err)
	}
}
