package player

import (
	"doodle-crawler/directions"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	XPosition  int32
	YPosition  int32
	MoveAmount int32
	Radius     float32
	Facing     directions.Direction
}

func (player Player) Draw() {
	rl.DrawCircle(player.XPosition, player.YPosition, player.Radius, rl.Yellow)

	switch player.Facing {
	case directions.North:
		rl.DrawLine(player.XPosition, player.YPosition, player.XPosition, player.YPosition-10, rl.Yellow)
	case directions.East:
		rl.DrawLine(player.XPosition, player.YPosition, player.XPosition+10, player.YPosition, rl.Yellow)
	case directions.South:
		rl.DrawLine(player.XPosition, player.YPosition, player.XPosition, player.YPosition+10, rl.Yellow)
	case directions.West:
		rl.DrawLine(player.XPosition, player.YPosition, player.XPosition-10, player.YPosition, rl.Yellow)
	}
}

func (player *Player) Move() {

	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp) {
		switch player.Facing {
		case directions.North:
			player.YPosition -= player.MoveAmount
		case directions.East:
			player.XPosition += player.MoveAmount
		case directions.South:
			player.YPosition += player.MoveAmount
		case directions.West:
			player.XPosition -= player.MoveAmount
		}
	}

	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		player.Facing = (player.Facing + 1) % 4
	}

	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		player.Facing = (player.Facing - 1 + 4) % 4
	}

	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown) {
		player.Facing = (player.Facing + 2) % 4
	}
}

func New() Player {
	return Player{
		XPosition:  81,
		YPosition:  81,
		MoveAmount: 54,
		Radius:     4,
		Facing:     directions.East,
	}
}
