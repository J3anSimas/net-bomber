package scene

import rl "github.com/gen2brain/raylib-go/raylib"

type PlayerInfo struct {
	name string
	id   int
}

type GameLobby struct {
	players []PlayerInfo
}

// Draw implements Scene.
func (g *GameLobby) Draw() {
	rl.DrawText("Game Lobby", 10, 10, 20, rl.Black)
}

// EnterState implements Scene.
func (g *GameLobby) EnterState() error {
	// Initialize game lobby resources here if needed
	return nil
}

// Update implements Scene.
func (g *GameLobby) Update(gameC GameController) (Scene, error) {
	return nil, nil
}

func NewGameLobby(windowWidth, windowHeight int32) *GameLobby {
	return &GameLobby{}
}

var _ Scene = (*GameLobby)(nil)
