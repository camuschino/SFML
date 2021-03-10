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

// func createNewMap() {

// 	s1 := rand.NewSource(time.Now().UnixNano())

// 	for i := 0; i < fieldDimentionX-1; i++ {
// 		for j := 0; j < fieldDimentionY-2; j++ {
// 			arrayToCheck[i][j] = false
// 			r1 := rand.New(s1)

// 			if i%2 == 0 && j%2 == 0 {
// 				arrayToMap[i][j] = false
// 			} else {
// 				if r1.Intn(2) == 0 {
// 					arrayToMap[i][j] = false
// 				} else {
// 					arrayToMap[i][j] = true
// 				}
// 			}
// 		}
// 	}
// }

// func checkMapByBFS(xActual, yActual int, imd *imdraw.IMDraw, win *pixelgl.Window) (validMap bool) {
// 	queue := list.New()
// 	objectivePoint := models.MapPoint{XPoint: objectivePositionX, YPoint: objectivePositionY}
// 	arrayToCheck[xActual][yActual] = true

// 	validMap = false

// 	if yActual > 0 && !arrayToMap[xActual][yActual-1] {
// 		queue.PushBack(models.MapPoint{XPoint: xActual, YPoint: yActual - 1})
// 	}
// 	if xActual > 0 && !arrayToMap[xActual-1][yActual] {
// 		queue.PushBack(models.MapPoint{XPoint: xActual - 1, YPoint: yActual})
// 	}
// 	if yActual < (fieldDimentionY-2) && !arrayToMap[xActual][yActual+1] {
// 		queue.PushBack(models.MapPoint{XPoint: xActual, YPoint: yActual + 1})
// 	}
// 	if xActual < (fieldDimentionX-1) && !arrayToMap[xActual+1][yActual] {
// 		queue.PushBack(models.MapPoint{XPoint: xActual + 1, YPoint: yActual})
// 	}

// 	xCurrent := xActual
// 	yCurrent := yActual

// 	for queue.Len() > 0 {
// 		// println("size: ", queue.Len())
// 		currentPoint := queue.Front()

// 		xCurrent = currentPoint.Value.(models.MapPoint).XPoint
// 		yCurrent = currentPoint.Value.(models.MapPoint).YPoint

// 		println(xCurrent, yCurrent)

// 		queue.Remove(currentPoint)

// 		// Check horizontal limit in the map.
// 		if !checkLimit(yCurrent, (fieldDimentionY - 2)) {
// 			continue
// 		}

// 		// Check vertical limit in the map.
// 		if !checkLimit(xCurrent, (fieldDimentionX - 1)) {
// 			continue
// 		}

// 		// This check if this point is playable. (true means false, because there's a wall)
// 		if arrayToMap[xCurrent][yCurrent] || arrayToCheck[xCurrent][yCurrent] {
// 			continue
// 		}

// 		arrayToCheck[xCurrent][yCurrent] = true

// 		imd.Color = colornames.Yellow
// 		px := getWall(xCurrent, yCurrent)
// 		imd.Push(px.Min, px.Max)
// 		imd.Rectangle(0)
// 		imd.Draw(win)
// 		win.Update()
// 		time.Sleep(20 * time.Millisecond)

// 		if objectivePoint == currentPoint.Value {
// 			validMap = true
// 			break
// 		}

// 		if checkLimit(yCurrent-1, (fieldDimentionY-2)) && !arrayToMap[xCurrent][yCurrent-1] && !arrayToCheck[xCurrent][yCurrent-1] {
// 			queue.PushBack(models.MapPoint{XPoint: xCurrent, YPoint: yCurrent - 1})
// 		}

// 		if checkLimit(xCurrent-1, (fieldDimentionX-1)) && !arrayToMap[xCurrent-1][yCurrent] && !arrayToCheck[xCurrent-1][yCurrent] {
// 			queue.PushBack(models.MapPoint{XPoint: xCurrent - 1, YPoint: yCurrent})
// 		}

// 		if checkLimit(yCurrent+1, (fieldDimentionY-2)) && !arrayToMap[xCurrent][yCurrent+1] && !arrayToCheck[xCurrent][yCurrent+1] {
// 			queue.PushBack(models.MapPoint{XPoint: xCurrent, YPoint: yCurrent + 1})
// 		}

// 		if checkLimit(xCurrent+1, (fieldDimentionX-1)) && !arrayToMap[xCurrent+1][yCurrent] && !arrayToCheck[xCurrent+1][yCurrent] {
// 			queue.PushBack(models.MapPoint{XPoint: xCurrent + 1, YPoint: yCurrent})
// 		}
// 	}

// 	return
// }

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
