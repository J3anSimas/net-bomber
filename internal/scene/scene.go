package scene

type Scene interface {
	Update(gameC GameController) (Scene, error)
	EnterState() error
	Draw()
}
type GameController interface {
	Update() error
	EnterState() error
	Draw()
	SetScene(scene Scene)
	GetWindowSize() (int32, int32)
}
