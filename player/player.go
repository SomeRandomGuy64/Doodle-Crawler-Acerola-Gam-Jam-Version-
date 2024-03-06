package player

import (
	"doodle-crawler/directions"
	"doodle-crawler/worldMaps"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DR = 0.0174533 //one degree in radians
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
	var rayX, rayY, rayAngle, xOffset, yOffset, disH, disV float32

	rayAngle = (float32(player.Facing)-1)*(0.5*math.Pi) - 30*DR

	for ray = 0; ray < 60; ray++ {

		if rayAngle < 0 {
			rayAngle += (2 * math.Pi)
		}

		//Check horizontal lines
		depthOfField = 0
		disH = 10000000
		horizontalX := player.XPosition
		horizontalY := player.YPosition
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
			rayY = (float32((player.YPosition>>5)<<5) + float32(player.MoveAmount))
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

			if mapPosition > 0 && mapPosition < worldMap.XSize*worldMap.YSize && mapDetails[mapPosition] == 1 {
				horizontalX = int32(rayX)
				horizontalY = int32(rayY)
				disH = distance(float32(player.XPosition), float32(player.YPosition), float32(horizontalX), float32(horizontalY))
				depthOfField = 16
			} else {
				rayX += xOffset
				rayY += yOffset
				depthOfField += 1
			}
		}

		//Check vertical lines
		depthOfField = 0
		disV = 10000000
		verticalX := player.XPosition
		verticalY := player.YPosition
		nTan := float32(-math.Tan(float64(rayAngle)))

		//Looking westwards
		if rayAngle > math.Pi/2 && rayAngle < (3*math.Pi)/2 {
			rayX = (float32((player.XPosition>>5)<<5) - 0.00001)
			rayY = ((float32(player.XPosition)-rayX)*nTan + float32(player.YPosition))
			xOffset = -float32(player.MoveAmount)
			yOffset = -xOffset * nTan
		}

		//looking eastwards
		if rayAngle < math.Pi/2 || rayAngle > (3*math.Pi)/2 {
			rayX = (float32((player.XPosition>>5)<<5) + float32(player.MoveAmount))
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

			if mapPosition > 0 && mapPosition < worldMap.XSize*worldMap.YSize && mapDetails[mapPosition] == 1 {
				verticalX = int32(rayX)
				verticalY = int32(rayY)
				disV = distance(float32(player.XPosition), float32(player.YPosition), float32(verticalX), float32(verticalY))
				depthOfField = 16
			} else {
				rayX += xOffset
				rayY += yOffset
				depthOfField += 1
			}
		}

		if disV < disH {
			rayX = float32(verticalX)
			rayY = float32(verticalY)
		}
		if disH < disV {
			rayX = float32(horizontalX)
			rayY = float32(horizontalY)
		}

		rl.DrawLine(player.XPosition, player.YPosition, int32(rayX), int32(rayY), rl.Red)
		rayAngle += DR
	}
}

func distance(aX, aY, bX, bY float32) float32 {
	return float32(math.Sqrt((float64(bX-aX)*float64(bX-aX) + float64(bY-aY)*float64(bY-aY))))
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
