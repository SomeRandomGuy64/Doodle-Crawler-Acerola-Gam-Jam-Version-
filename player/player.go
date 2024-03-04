package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	XPosition  int32
	YPosition  int32
	MoveAmount int32
	Radius     float32
}

func (player Player) Draw() {
	rl.DrawCircle(player.XPosition, player.YPosition, player.Radius, rl.Yellow)
}

func (player *Player) Move() {
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		player.XPosition += player.MoveAmount
	}

	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		player.XPosition -= player.MoveAmount
	}

	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown) {
		player.YPosition += player.MoveAmount
	}

	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp) {
		player.YPosition -= player.MoveAmount
	}
}

func New() Player {
	return Player{
		XPosition:  81,
		YPosition:  81,
		MoveAmount: 54,
		Radius:     4,
	}
}
