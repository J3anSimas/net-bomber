package main

import (
	"log"
	"snake-go/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 720
)

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Net Bomber")
	ebiten.SetVsyncEnabled(true)

	game := game.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
