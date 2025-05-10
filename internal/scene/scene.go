package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(gameC GameController) error
	EnterState() error
	Draw(screen *ebiten.Image)
}
type GameController interface {
	Update() error
	EnterState() error
	Draw(screen *ebiten.Image)
	SetScene(scene Scene)
	GetWindowSize() (int, int)
}
