package worldMaps

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WorldMap struct {
	XSize      int32
	YSize      int32
	BlockSize  int32
	MapDetails []int32
}

func (worldMap WorldMap) DrawMap() {
	var xPosition, yPosition, xOffset, yOffset int32
	var usedColor color.RGBA

	for yPosition = 0; yPosition < worldMap.YSize; yPosition++ {
		for xPosition = 0; xPosition < worldMap.XSize; xPosition++ {

			if worldMap.MapDetails[yPosition*worldMap.XSize+xPosition] == 1 {
				usedColor = rl.White
			} else {
				usedColor = rl.Black
			}
			xOffset = xPosition * worldMap.BlockSize
			yOffset = yPosition * worldMap.BlockSize
			rl.DrawRectangle(xOffset+1, yOffset+1, worldMap.BlockSize-1, worldMap.BlockSize-1, usedColor)
		}
	}
}

func New(xSize, ySize int32, mapDetails []int32) WorldMap {
	return WorldMap{
		XSize:      xSize,
		YSize:      ySize,
		BlockSize:  32,
		MapDetails: mapDetails,
	}
}
