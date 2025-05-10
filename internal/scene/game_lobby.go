package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerInfo struct {
	name string
	id   int
}

type GameLobby struct {
	players []PlayerInfo
}

// Draw implements Scene.
func (g *GameLobby) Draw(screen *ebiten.Image) {
	offset := 20
	windowWidth, windowHeight := ebiten.WindowSize()
	vector.StrokeRect(screen, float32(0+offset), float32(0+offset), float32(windowWidth-offset*2), float32(windowHeight-offset*2), 5, color.White, true)
	for i := range 4 {
		x := 30*i + 10
		y := 30
		vector.StrokeRect(screen, float32(x), float32(y), float32(100), float32(300), 5, color.White, true)
		if i < len(g.players) {
			player := g.players[i]
			vector.Draw(screen, player.name, float32(x+10), float32(y+10), color.White, nil)
		}
	}
}

// EnterState implements Scene.
func (g *GameLobby) EnterState() error {
	// Initialize game lobby resources here if needed
	return nil
}

// Update implements Scene.
func (g *GameLobby) Update(gameC GameController) error {
	return nil
}

func NewGameLobby(windowWidth, windowHeight int) *GameLobby {
	return &GameLobby{}
}

var _ Scene = (*GameLobby)(nil)
