package scene

import (
	"image/color"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	buttons []*button
	cursor  int
}

func (m *MainMenu) EnterState() error {
	return nil
}
func (m *MainMenu) Update(gameC GameController) (Scene, error) {
	if scn, err := m.processMouseClick(); err != nil {
		return nil, err
	} else if scn != nil {
		return scn, nil
	}

	if scn, err := m.processKeyboardInput(); err != nil {
		return nil, err
	} else if scn != nil {
		return scn, nil
	}
	return nil, nil

}
func (m *MainMenu) processMouseClick() (Scene, error) {

	for i, button := range m.buttons {
		if button.CheckIfWasHovered() {
			m.cursor = i
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for _, button := range m.buttons {
			if !button.CheckIfWasClicked() {
				continue
			}
			newScene, err := button.OnClick()
			if err != nil {
				return nil, err
			}
			if newScene != nil {
				return newScene, nil
			}
		}
	}
	return nil, nil

}
func (m *MainMenu) processKeyboardInput() (Scene, error) {
	if rl.IsKeyPressed(rl.KeyUp) {
		if m.cursor > 0 {
			m.cursor--
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if m.cursor < len(m.buttons)-1 {
			m.cursor++
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		scn, err := m.buttons[m.cursor].OnClick()
		return scn, err
	}
	return nil, nil
}
func (m *MainMenu) Draw() {

	rl.DrawText("NET BOMBER", 10, 10, 32, rl.Black)
	for i, button := range m.buttons {
		if i == m.cursor {
			// Highlight the button if it's selected
			button.Draw(color.RGBA{255, 0, 0, 255})
		} else {
			button.Draw(rl.LightGray)
		}

	}
}

//	func repeatingKeyPressed(key ebiten.Key) bool {
//		const (
//			delay    = 30
//			interval = 3
//		)
//		d := inpututil.KeyPressDuration(key)
//		if d == 1 {
//			return true
//		}
//		if d >= delay && (d-delay)%interval == 0 {
//			return true
//		}
//		return false
//	}
func NewMainMenu() *MainMenu {
	buttons := []*button{
		NewButton(100, 100, 200, 50, "New Game", func() (Scene, error) {
			rl.TraceLog(rl.LogInfo, "New Game clicked!")
			// Get window size
			windowWidth, windowHeight := rl.GetScreenWidth(), rl.GetScreenHeight()
			gameLobby := NewGameLobby(int32(windowWidth), int32(windowHeight))
			return gameLobby, nil

		}),
		NewButton(100, 200, 200, 50, "Options", func() (Scene, error) {
			rl.TraceLog(rl.LogInfo, "Button Option clicked!")
			return nil, nil
		}),
		NewButton(100, 300, 200, 50, "Exit", func() (Scene, error) {
			os.Exit(0)
			return nil, nil
		}),
	}
	// buttons := []*component.Button{}
	return &MainMenu{
		buttons: buttons,
		cursor:  0,
	}
}

type button struct {
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Text    string
	OnClick func() (Scene, error)
}

func NewButton(x, y, width, height int32, text string, onClick func() (Scene, error)) *button {
	return &button{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Text:    text,
		OnClick: onClick,
	}
}
func (b *button) Draw(clr color.RGBA) {
	var choosedColor color.RGBA
	if clr == (color.RGBA{}) {
		choosedColor = rl.LightGray
	} else {
		choosedColor = clr
	}
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, choosedColor)

	// Draw the button text
	textWidth := rl.MeasureText(b.Text, 20)
	textX := b.X + (b.Width-textWidth)/2
	textY := b.Y + (b.Height-20)/2
	rl.DrawText(b.Text, textX, textY, 20, rl.Black)

	// Check for mouse click
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
	}
}
func (b *button) CheckIfWasClicked() bool {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if mouseX >= b.X && mouseX <= b.X+b.Width && mouseY >= b.Y && mouseY <= b.Y+b.Height {
		return true
	}
	return false
}
func (b *button) CheckIfWasHovered() bool {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()

	if mouseX >= b.X && mouseX <= b.X+b.Width && mouseY >= b.Y && mouseY <= b.Y+b.Height {
		return true
	}
	return false
}

var _ Scene = (*MainMenu)(nil)
