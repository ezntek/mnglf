package ball

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Radius   float32
	Position rl.Vector2
	Velocity rl.Vector2
	Color    color.RGBA
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, b.Color)
}

func New(radius float32, x, y float32, color color.RGBA) *Ball {
	return &Ball{
		Radius:   radius,
		Position: rl.NewVector2(x, y),
		Velocity: rl.NewVector2(0, 0),
		Color:    color,
	}
}
