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
	// if i < len(g.players) {
	// 	player := g.players[i]
	//
	// }
	for i := range 4 {
		vector.StrokeRect(screen, float32(offset+16+(i*offset+16)), 16, 304, 500, 2, color.RGBA{255, 0, 0, 255}, true)
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
