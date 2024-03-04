package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var playerPositionX, playerPositionY, playerSpeed int32
	var playerRadius float32
	var mapX, mapY, mapS int32 //mapX and mapY determine the width and height whilst mapS determines the size of the squares

	mapX = 16
	mapY = 16
	mapS = 55
	mapOne := []int32{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	playerPositionX = 300
	playerPositionY = 300
	playerRadius = 4
	playerSpeed = 5

	rl.InitWindow(1024, 896, "Doodle Crawler")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Gray)

		drawMap(mapX, mapY, mapS, mapOne)
		drawPlayer(&playerPositionX, &playerPositionY, &playerRadius)
		movePlayer(&playerPositionX, &playerPositionY, playerSpeed)

		rl.EndDrawing()
	}
}

func drawMap(mapX, mapY, mapS int32, worldMap []int32) {
	var xPosition, yPosition, xOffset, yOffset int32
	var usedColor color.RGBA

	for yPosition = 0; yPosition < mapY; yPosition++ {
		for xPosition = 0; xPosition < mapX; xPosition++ {

			if worldMap[yPosition*mapX+xPosition] == 1 {
				usedColor = rl.White
			} else {
				usedColor = rl.Black
			}
			xOffset = xPosition * mapS
			yOffset = yPosition * mapS
			rl.DrawRectangle(xOffset+1, yOffset+1, mapS-1, mapS-1, usedColor)
		}
	}
}

func drawPlayer(playerPositionX, playerPositionY *int32, playerRadius *float32) {
	rl.DrawCircle(*playerPositionX, *playerPositionY, *playerRadius, rl.Yellow)
}

func movePlayer(playerPositionX, playerPositionY *int32, playerSpeed int32) {
	if rl.IsKeyDown(rl.KeyRight) {
		*playerPositionX += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		*playerPositionX -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyDown) {
		*playerPositionY += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyUp) {
		*playerPositionY -= playerSpeed
	}
}
