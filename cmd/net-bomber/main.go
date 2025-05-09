package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"snake-go/internal/scene"
)

var (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 720
)

type Game struct {
	currentScene  scene.Scene
	mainMenuScene *scene.MainMenu
}

func NewGame(winWidth, winHeight int) ebiten.Game {
	mainMenu := scene.NewMainMenu()
	err := mainMenu.EnterState()
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		currentScene:  mainMenu,
		mainMenuScene: mainMenu,
	}
}

func (g *Game) Update() error {
	err := g.currentScene.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Net Bomber")
	ebiten.SetVsyncEnabled(true)

	game := NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
