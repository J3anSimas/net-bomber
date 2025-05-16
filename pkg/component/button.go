package component

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	X       int32
	Y       int32
	Width   int32
	Height  int32
	Text    string
	OnClick func()
}

func NewButton(x, y, width, height int32, text string, onClick func()) *Button {
	return &Button{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Text:    text,
		OnClick: onClick,
	}
}
func (b *Button) Draw(clr color.RGBA) {
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
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		if mouseX >= b.X && mouseX <= b.X+b.Width && mouseY >= b.Y && mouseY <= b.Y+b.Height {
			b.OnClick()
		}
	}
}
