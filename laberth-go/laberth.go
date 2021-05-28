package main

import (
	"github.com/camuschino/laberth-go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowDimentionX, windowDimentionY, sizeBlock int     = 100, 100, 10 // window dimention AND large: 100. medium: 50, little: 20, nano: 10
	fieldDimentionX, fieldDimentionY              int     = ((windowDimentionX / sizeBlock) * 2) + 1, ((windowDimentionY / sizeBlock) * 2) + 1
	sizeField                                     int     = sizeBlock / 2
	movementDistance                              float32 = float32(sizeField)
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Laberth",
		Bounds: pixel.R(0, 0, float64(windowDimentionY), float64(windowDimentionX)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	laberth := utils.CreateNewMap(fieldDimentionX, fieldDimentionY, sizeField, movementDistance)

	player, target := utils.SetObjectPositions(&laberth)
	utils.RenderMapAndObjects(&laberth, player, target, imd, win)

	algorithm := "BFS"

	valid := utils.ValidateMap(algorithm, player, target, &laberth, imd, win)
	println(valid)

	for {
		laberth := utils.CreateNewMap(fieldDimentionX, fieldDimentionY, sizeField, movementDistance)

		player, target := utils.SetObjectPositions(&laberth)
		win.Clear(colornames.Black)
		imd.Clear()
		utils.RenderMapAndObjects(&laberth, player, target, imd, win)

		if algorithm == "BFS" {
			algorithm = "DFS"
		} else {
			algorithm = "BFS"
		}

		valid := utils.ValidateMap(algorithm, player, target, &laberth, imd, win)
		println(valid)
	}
}

func main() {
	pixelgl.Run(run)
}
