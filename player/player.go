package player

import (
	"doodle-crawler/directions"
	"doodle-crawler/worldMaps"
	"fmt"
	"math"

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

func (player Player) DrawRays(mapDetails []int32, worldMap worldMaps.WorldMap) {
	var ray, mapX, mapY, mapPosition, depthOfField int32
	var rayX, rayY, rayAngle, xOffset, yOffset float32

	rayAngle = (float32(player.Facing) - 1) * (0.5 * math.Pi)

	for ray = 0; ray < 2; ray++ {

		if rayAngle < 0 {
			rayAngle += (2 * math.Pi)
		}

		//Check horizontal lines
		depthOfField = 0
		aTan := float32(-1 / math.Tan(float64(rayAngle)))

		//Looking northwards
		if rayAngle > math.Pi {
			rayY = (float32((player.YPosition>>5)<<5) - 0.0001)
			rayX = ((float32(player.YPosition)-rayY)*aTan + float32(player.XPosition))
			yOffset = -float32(player.MoveAmount)
			xOffset = -yOffset * aTan
		}

		//looking southwards
		if rayAngle < math.Pi {
			rayY = (float32((player.YPosition>>5)<<5) + 64)
			rayX = (float32(player.YPosition)-rayY)*aTan + float32(player.XPosition)
			yOffset = float32(player.MoveAmount)
			xOffset = -yOffset * aTan
		}

		//looking straight left or right
		if rayAngle == 0 || rayAngle == math.Pi {
			rayX = float32(player.XPosition)
			rayY = float32(player.YPosition)
			depthOfField = 16
		}

		for depthOfField < 16 {
			mapX = int32(math.Floor(float64(rayX / float32(player.MoveAmount))))
			mapY = int32(math.Floor(float64(rayY / float32(player.MoveAmount))))
			mapPosition = mapY*worldMap.XSize + mapX

			if mapPosition < worldMap.XSize*worldMap.YSize && mapDetails[mapPosition] == 1 {
				depthOfField = 16
			} else {
				rayX += xOffset
				rayY += yOffset
				depthOfField += 1
			}
		}
		rl.DrawLine(player.XPosition, player.YPosition, int32(rayX), int32(rayY), rl.Green)

		//Check vertical lines
		depthOfField = 0
		nTan := float32(-math.Tan(float64(rayAngle)))

		//Looking westwards
		if rayAngle > math.Pi/2 && rayAngle < (3*math.Pi)/2 {
			rayX = (float32((player.XPosition>>5)<<5) - 0.0001)
			rayY = ((float32(player.XPosition)-rayX)*nTan + float32(player.YPosition))
			xOffset = -float32(player.MoveAmount)
			yOffset = -xOffset * nTan
		}

		//looking eastwards
		if rayAngle < math.Pi/2 || rayAngle > (3*math.Pi)/2 {
			rayX = (float32((player.XPosition>>5)<<5) + 64)
			rayY = (float32(player.XPosition)-rayX)*nTan + float32(player.YPosition)
			xOffset = float32(player.MoveAmount)
			yOffset = -xOffset * nTan
		}

		//looking straight up or down
		if rayAngle == math.Pi/2 || rayAngle == (3*math.Pi)/2 {
			rayX = float32(player.XPosition)
			rayY = float32(player.YPosition)
			depthOfField = 16
		}

		for depthOfField < 16 {
			mapX = int32(math.Floor(float64(rayX / float32(player.MoveAmount))))
			mapY = int32(math.Floor(float64(rayY / float32(player.MoveAmount))))
			mapPosition = mapY*worldMap.XSize + mapX

			if mapPosition < worldMap.XSize*worldMap.YSize && mapDetails[mapPosition] == 1 {
				depthOfField = 16
			} else {
				rayX += xOffset
				rayY += yOffset
				depthOfField += 1
			}
		}
		rl.DrawLine(player.XPosition, player.YPosition, int32(rayX), int32(rayY), rl.Red)
		fmt.Println(rayAngle, rayX, rayY, player.XPosition, player.YPosition)

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
		XPosition:  48,
		YPosition:  48,
		MoveAmount: 32,
		Radius:     4,
		Facing:     directions.East,
	}
}
