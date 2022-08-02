package ball

import (
	"fmt"
	"image/color"
	"math"

	// "math"

	rl "github.com/gen2brain/raylib-go/raylib"
	//pl "github.com/tek967/rgbapalette"
)

type Ball struct {
	count                 int
	isBallMoving          bool
	Radius                float32
	Position              rl.Vector2
	OneDVelocity          float32
	Velocity              rl.Vector2
	Color                 color.RGBA
	friction              float32
	initialCursorPosition rl.Vector2
	mouseCurrentPosition  rl.Vector2
}

func GetDistance(a, b rl.Vector2) float32 {
	return float32(math.Sqrt(math.Pow(math.Abs(float64(a.X)-float64(a.Y)), 2) + math.Pow(math.Abs(float64(a.Y)-float64(b.Y)), 2)))
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, b.Color)
}

func (b *Ball) Collision(sw, sh *int32) {
	if int32(b.Position.X) > *sw {
		b.Velocity.X = -b.Velocity.X
		b.Position.X = float32(*sw) - 2
	} else if int32(b.Position.X) < 0 {
		b.Velocity.X = -b.Velocity.X
		b.Position.X = 2
	}
	if int32(b.Position.Y) > *sh {
		b.Velocity.Y = -b.Velocity.Y
		b.Position.Y = float32(*sh) - 2
	} else if int32(b.Position.Y) < 0 {
		b.Velocity.Y = -b.Velocity.Y
		b.Position.Y = 2
	}
}

func (b *Ball) Update(sw, sh *int32) {
	b.count++
	b.Collision(sw, sh)

	if !b.isBallMoving {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.mouseCurrentPosition = rl.GetMousePosition()
			b.initialCursorPosition = rl.GetMousePosition()
		}
	}
	if rl.IsMouseButtonDown(rl.MouseLeftButton) && b.initialCursorPosition.X != 0 && b.initialCursorPosition.Y != 0 {
		b.mouseCurrentPosition = rl.GetMousePosition()
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		b.isBallMoving = true
		b.Velocity.X = -(b.mouseCurrentPosition.X - b.initialCursorPosition.X) / 5.7
		b.Velocity.Y = -(b.mouseCurrentPosition.Y - b.initialCursorPosition.Y) / 5.7
	}
	if b.isBallMoving {
		if !rl.IsMouseButtonDown(rl.MouseLeftButton) && (b.Velocity.X != 0 || b.Velocity.Y != 0) {
			if b.Velocity.X > 0 {
				b.Velocity.X -= b.friction
			} else if b.Velocity.X < 0 {
				b.Velocity.X += b.friction
			}

			if b.Velocity.Y > 0 {
				b.Velocity.Y -= b.friction
			} else if b.Velocity.Y < 0 {
				b.Velocity.Y += b.friction
			}

			if b.Velocity.X > -b.friction && b.Velocity.X < b.friction {
				b.Velocity.X = 0
			}
			if b.Velocity.Y > -b.friction && b.Velocity.Y < b.friction {
				b.Velocity.Y = 0
			}
		}
	}
	if b.Velocity.X == 0 && b.Velocity.Y == 0 {
		b.isBallMoving = false
	}
	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y
	fmt.Printf("(VEL: %0.2f, %0.2f; POS: %0.2f, %0.2f) %d\n", b.Velocity.X, b.Velocity.Y, b.Position.X, b.Position.Y, b.count)
}

func New(radius float32, x, y float32, color color.RGBA) *Ball {
	return &Ball{
		Radius:   radius,
		friction: 1.2,
		Position: rl.NewVector2(x, y),
		Velocity: rl.NewVector2(0, 0),
		Color:    color,
	}
}
