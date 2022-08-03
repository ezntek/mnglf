package powermeter

import rl "github.com/gen2brain/raylib-go/raylib"

type PowerMeter struct {
	borderRect  rl.Rectangle
	innerRect   rl.Rectangle
	fillPercent float32
	IsHidden    bool
	Position    rl.Vector2
}

func (p *PowerMeter) Draw() {
	if !p.IsHidden {
		rl.DrawRectangleRec(p.borderRect, rl.Black)
		rl.DrawRectangleRec(p.innerRect, rl.Green)
	}
}

func New(pos rl.Vector2, initialFillPercent float32) *PowerMeter {
	return &PowerMeter{
		borderRect:  rl.NewRectangle(pos.X-5, pos.Y-5, 60, 20),
		innerRect:   rl.NewRectangle(pos.X, pos.Y, initialFillPercent*50, 10),
		fillPercent: initialFillPercent,
		IsHidden:    false,
		Position:    pos,
	}
}
