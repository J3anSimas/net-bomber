package game

import (
	"log"
	"snake-go/internal/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentScene   scene.Scene
	mainMenuScene  *scene.MainMenu
	gameLobbyScene *scene.GameLobby
}

func NewGame(winWidth, winHeight int) ebiten.Game {
	mainMenu := scene.NewMainMenu()
	gameLobby := scene.NewGameLobby(winWidth, winHeight)
	err := mainMenu.EnterState()
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		currentScene:   gameLobby,
		mainMenuScene:  mainMenu,
		gameLobbyScene: gameLobby,
	}
}
func (g *Game) EnterState() error {
	err := g.currentScene.EnterState()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Update() error {
	err := g.currentScene.Update(g)
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
func (g *Game) SetScene(scene scene.Scene) {
	g.currentScene = scene
	err := g.currentScene.EnterState()
	if err != nil {
		log.Fatal(err)
	}
}
func (g *Game) GetWindowSize() (int, int) {
	return ebiten.WindowSize()
}

// var _ scene.GameController = (*Game)(nil)
