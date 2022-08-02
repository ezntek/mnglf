package arrow

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Arrow struct {
	Position rl.Vector2
	Angle    float32
	Length   int32
	TipWidth int32
}

func (arrow *Arrow) Draw() {
	rl.DrawLineV(
		arrow.Position,
		rl.Vector2{},
		color.RGBA{16, 16, 16, 255},
	)
}

func (arrow *Arrow) Update() {
}

func New(pos rl.Vector2, len int32, tipwidth int32, angle float32) *Arrow {
	return &Arrow{
		Position: pos,
		Angle:    angle,
		Length:   len,
		TipWidth: tipwidth,
	}
}
