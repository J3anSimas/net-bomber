package game

import (
	"log"

	"github.com/j3ansimas/net-bomber/internal/scene"
)

type Game struct {
	currentScene   scene.Scene
	mainMenuScene  *scene.MainMenu
	gameLobbyScene *scene.GameLobby
	WindowWidth    int32
	WindowHeight   int32
}

func NewGame(winWidth, winHeight int32) *Game {
	mainMenu := scene.NewMainMenu()
	gameLobby := scene.NewGameLobby(winWidth, winHeight)
	err := mainMenu.EnterState()
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		currentScene:   mainMenu,
		mainMenuScene:  mainMenu,
		gameLobbyScene: gameLobby,
		WindowWidth:    winWidth,
		WindowHeight:   winHeight,
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
	scn, err := g.currentScene.Update(g)
	if err != nil {
		return err
	}
	if scn != nil {
		g.SetScene(scn)
	}
	return nil
}

func (g *Game) Draw() {
	g.currentScene.Draw()
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
func (g *Game) GetWindowSize() (int32, int32) {
	return g.WindowWidth, g.WindowHeight
}

// var _ scene.GameController = (*Game)(nil)
