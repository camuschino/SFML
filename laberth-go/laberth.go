package main

import (
	"github.com/camuschino/laberth-go/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowDimentionX, windowDimentionY, sizeBlock int     = 500, 500, 50 // window dimention AND large: 100. medium: 50, little: 20, nano: 10
	fieldDimentionX, fieldDimentionY              int     = ((windowDimentionX / sizeBlock) * 2) + 1, ((windowDimentionY / sizeBlock) * 2) + 2
	sizeField                                     int     = sizeBlock / 2
	movementDistance                              float32 = float32(sizeField)
)

var (
	arrayToCheck, arrayToMap                                                                [fieldDimentionX][fieldDimentionY]bool
	playerPositionX, playerPositionY, objectivePositionX, objectivePositionY                int
	mIsMovingUp, mIsMovingRight, mIsMovingLeft, mIsMovingDown, isMapChecked, isMapGenerated bool
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

	player, target := utils.SetObjectPositions(laberth)
	utils.RenderMapAndObjects(laberth, player, target, imd, win)

	valid := utils.ValidateMap("BFS", player, target, laberth, imd, win)
	println(valid)

	for true {
		laberth := utils.CreateNewMap(fieldDimentionX, fieldDimentionY, sizeField, movementDistance)

		player, target := utils.SetObjectPositions(laberth)
		win.Clear(colornames.Black)
		imd.Clear()
		utils.RenderMapAndObjects(laberth, player, target, imd, win)

		valid := utils.ValidateMap("BFS", player, target, laberth, imd, win)
		println(valid)
	}

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}

func checkLimit(currentValue, limit int) bool {
	return currentValue >= 0 && currentValue < limit
}

// func checkMapByDFS(yActual, xActual int, imd *imdraw.IMDraw, win *pixelgl.Window) bool {

// 	// Check vertical limit in the map.
// 	if yActual >= (fieldDimentionX-1) || yActual < 0 {
// 		return false
// 	}

// 	// Check horizontal limit in the map.
// 	if xActual >= (fieldDimentionY-2) || xActual < 0 {
// 		return false
// 	}

// 	// This check if this point is playable.
// 	if arrayToMap[yActual][xActual] {
// 		return false
// 	}

// 	// Check if this position is already previously checked.
// 	if arrayToCheck[yActual][xActual] {
// 		return false
// 	}

// 	arrayToCheck[yActual][xActual] = true
// 	// println(xActual, yActual)

// 	// Check if this point is the point where is the objective.
// 	if !(playerPositionX == yActual && playerPositionY == xActual) {
// 		imd.Color = colornames.Yellow
// 		px := getWall(yActual, xActual)
// 		imd.Push(px.Min, px.Max)
// 		imd.Rectangle(0)
// 		imd.Draw(win)
// 		win.Update()
// 		time.Sleep(20 * time.Millisecond)
// 	}

// 	// Check if this point is the point where is the objective.
// 	if objectivePositionY == xActual && objectivePositionX == yActual {
// 		// println("EXITO")
// 		return true
// 	}

// 	if checkMapByDFS(yActual-1, xActual, imd, win) {
// 		return true
// 	}

// 	if checkMapByDFS(yActual, xActual-1, imd, win) {
// 		return true
// 	}

// 	if checkMapByDFS(yActual+1, xActual, imd, win) {
// 		return true
// 	}

// 	if checkMapByDFS(yActual, xActual+1, imd, win) {
// 		return true
// 	}

// 	return false
// }

func main() {
	pixelgl.Run(run)
}
