package main

import (
	"github.com/camuschino/laberth-go/models"
	"github.com/camuschino/laberth-go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowDimentionX, windowDimentionY, sizeBlock int     = 600, 600, 50 // window dimention AND large: 100. medium: 50, little: 20, nano: 10
	fieldDimentionX, fieldDimentionY              int     = ((windowDimentionX / sizeBlock) * 2) + 1, ((windowDimentionY / sizeBlock) * 2) + 1
	sizeField                                     int     = sizeBlock / 2
	movementDistance                              float32 = float32(sizeField)
)

var (
	laberth        models.Labyrinth
	player, target models.Coords
)

func getWindowConfigs() pixelgl.WindowConfig {
	return pixelgl.WindowConfig{
		Title:  "Laberth",
		Bounds: pixel.R(0, 0, float64(windowDimentionY), float64(windowDimentionX)),
		VSync:  true,
	}
}

func getWindowAndImd() (*pixelgl.Window, *imdraw.IMDraw) {
	win, err := pixelgl.NewWindow(getWindowConfigs())
	if err != nil {
		panic(err)
	}
	return win, imdraw.New(nil)
}

func getNewEmptyMap() models.Labyrinth {
	laberth := models.Labyrinth{
		SizeField:        sizeField,
		MovementDistance: movementDistance,
	}
	laberth.CreateNewEmptyMap(fieldDimentionX, fieldDimentionY)
	return laberth
}

func getNewMapAndObjects(newEmptyMap *models.Labyrinth) (models.Labyrinth, models.Coords, models.Coords) {

	utils.CreateNewLabyrinth(newEmptyMap)
	player, target := utils.SetObjectPositions(newEmptyMap)

	return *newEmptyMap, player, target
}

func run() {

	win, imd := getWindowAndImd()
	newEmptyMap := getNewEmptyMap()

	for {
		laberth, player, target = getNewMapAndObjects(&newEmptyMap)
		win.Clear(colornames.Black)
		imd.Clear()
		utils.RenderMapAndObjects(&laberth, player, target, imd, win)
		utils.ValidateMap("", player, target, &laberth, imd, win, sizeField)
	}
}

func main() {
	pixelgl.Run(run)
}
