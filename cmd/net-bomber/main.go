package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j3ansimas/net-bomber/internal/game"
	"github.com/joho/godotenv"
)

var (
	WINDOW_WIDTH  int32 = 1280
	WINDOW_HEIGHT int32 = 720
)

func main() {
	godotenv.Load()
	env := os.Getenv("ENV")
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	g := game.NewGame(WINDOW_WIDTH, WINDOW_HEIGHT)
	_ = g
	fmt.Printf("Environment: %s", env)
	if env == "dev" {
		fmt.Println("Dev mode")
		monitorCount := rl.GetMonitorCount()
		if monitorCount > 1 {
			// Move a janela para o segundo monitor (índice 1)
			rl.SetWindowMonitor(1)
		}
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		g.Update()
		g.Draw()
		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
}
