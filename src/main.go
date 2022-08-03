package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tek967/mnglf/src/ball"
	pl "github.com/tek967/rgbapalette"
)

var (
	width  int32 = 1024
	height int32 = 768
)

func drawCheckerboardBG(boxWidth int32, color1, color2 color.RGBA) {
	var isAlternateColor bool
	for y := 0; y < int(height); y += int(boxWidth) {
		for x := 0; x < int(width); x += int(boxWidth) {
			if isAlternateColor {
				rl.DrawRectangle(int32(x), int32(y), int32(boxWidth), int32(boxWidth), color2)
			} else {
				rl.DrawRectangle(int32(x), int32(y), int32(boxWidth), int32(boxWidth), color1)
			}
			isAlternateColor = !isAlternateColor
		}
		isAlternateColor = !isAlternateColor
	}
}

func main() {
	rl.InitWindow(width, height, "mnglf")
	rl.SetTargetFPS(60)

	JetBrainsMonoBold := rl.LoadFont("res/JetBrainsMono-Bold.ttf")

	golfBall := ball.New(17, float32(width)/2-12, float32(height)/2-12, pl.Palette["darkgreen"])

	for !rl.WindowShouldClose() {
		golfBall.Update(&width, &height)

		// drawing
		rl.BeginDrawing()
		drawCheckerboardBG(16, pl.Palette["lightgreen"], color.RGBA{139, 229, 139, 255})
		rl.ClearBackground(pl.Palette["verylightgray"])
		rl.DrawTextEx(JetBrainsMonoBold, fmt.Sprintf("(VEL: %0.2f, %0.2f; POS: %0.2f, %0.2f)", golfBall.Velocity.X, golfBall.Velocity.Y, golfBall.Velocity.X, golfBall.Position.Y), rl.NewVector2(20, 20), 30, float32(JetBrainsMonoBold.BaseSize/6), rl.Black)

		// draw functions
		golfBall.Draw()
		// ---
		rl.EndDrawing()
		// ---
	}
}
