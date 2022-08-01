package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	pl "github.com/tek967/rgbapalette"
)

const (
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

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		drawCheckerboardBG(16, pl.Palette["lightgreen"], color.RGBA{139, 229, 139, 255})
		rl.ClearBackground(pl.Palette["verylightgray"])
		rl.EndDrawing()
	}
}
