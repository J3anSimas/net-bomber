package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() error
	EnterState() error
	Draw(*ebiten.Image)
}
