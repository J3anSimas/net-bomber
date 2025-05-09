package scene

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"log"
	"os"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	choicesFontSize = float64(64)
	titleFontSize   = float64(128)
)

type MainMenu struct {
	choices []string
	cursor  int
}

func (m *MainMenu) EnterState() error {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
	return nil
}
func (m *MainMenu) Update() error {
	if repeatingKeyPressed(ebiten.KeyUp) || repeatingKeyPressed(ebiten.KeyW) {
		if m.cursor > 0 {
			m.cursor--
		}
	}
	if repeatingKeyPressed(ebiten.KeyDown) || repeatingKeyPressed(ebiten.KeyS) {
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	}
	if repeatingKeyPressed(ebiten.KeyEnter) {
		if m.cursor == len(m.choices)-1 {
			os.Exit(0)
		}
	}
	return nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 27, G: 34, B: 61, A: 255})
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.RGBA{R: 12, G: 59, B: 242, A: 255})
	op.GeoM.Translate(20, 10)
	text.Draw(screen, "NET BOMBER", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   titleFontSize,
	}, op)
	y := 20 + titleFontSize + choicesFontSize
	for i, choice := range m.choices {
		op = &text.DrawOptions{}
		op.GeoM.Translate(20, y)
		if i == m.cursor {
			op.ColorScale.ScaleWithColor(color.RGBA{R: 50, G: 255, B: 50, A: 255})
		} else {
			op.ColorScale.ScaleWithColor(color.White)
		}
		text.Draw(screen, choice, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   choicesFontSize,
		}, op)
		y += choicesFontSize + 16
	}
}
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
func NewMainMenu() *MainMenu {
	return &MainMenu{
		choices: []string{"New Room", "Join Room", "Options", "Exit"},
		cursor:  0,
	}
}
