package utils

import (
	"container/list"
	"time"

	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// ValidateMap function which works fine
func ValidateMap(algorithm string, player, target models.MapPoint, laberth models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) (result bool) {

	switch algorithm {
	case "BFS":
		result = checkMapByBFS(player, target, laberth, imd, win)
	case "DFS":
		// result = checkMapByDFS()
	}
	return
}

func checkMapByBFS(player, target models.MapPoint, laberth models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) (validMap bool) {
	queue := list.New()
	laberth.ArrayToCheck[player.XPoint][player.YPoint] = true

	validMap = false

	if player.YPoint > 0 && !laberth.ArrayToMap[player.XPoint][player.YPoint-1] {
		queue.PushBack(models.MapPoint{XPoint: player.XPoint, YPoint: player.YPoint - 1})
	}
	if player.XPoint > 0 && !laberth.ArrayToMap[player.XPoint-1][player.YPoint] {
		queue.PushBack(models.MapPoint{XPoint: player.XPoint - 1, YPoint: player.YPoint})
	}
	if player.YPoint < (laberth.FieldDimentionY-2) && !laberth.ArrayToMap[player.XPoint][player.YPoint+1] {
		queue.PushBack(models.MapPoint{XPoint: player.XPoint, YPoint: player.YPoint + 1})
	}
	if player.XPoint < (laberth.FieldDimentionX-1) && !laberth.ArrayToMap[player.XPoint+1][player.YPoint] {
		queue.PushBack(models.MapPoint{XPoint: player.XPoint + 1, YPoint: player.YPoint})
	}

	xCurrent := 0
	yCurrent := 0

	time.Sleep(1000 * time.Millisecond)

	for queue.Len() > 0 {
		currentPoint := queue.Front()

		xCurrent = currentPoint.Value.(models.MapPoint).XPoint
		yCurrent = currentPoint.Value.(models.MapPoint).YPoint

		println(xCurrent, yCurrent)

		queue.Remove(currentPoint)

		// Check horizontal limit in the map.
		if !checkLimit(yCurrent, (laberth.FieldDimentionY - 2)) {
			continue
		}

		// Check vertical limit in the map.
		if !checkLimit(xCurrent, (laberth.FieldDimentionX - 1)) {
			continue
		}

		// This check if this point is playable. (true means false, because there's a wall)
		if laberth.ArrayToMap[xCurrent][yCurrent] || laberth.ArrayToCheck[xCurrent][yCurrent] {
			continue
		}

		laberth.ArrayToCheck[xCurrent][yCurrent] = true

		if target == currentPoint.Value {
			imd.Color = colornames.Red
			px := getWall(xCurrent, yCurrent, laberth.SizeField)
			imd.Push(px.Min, px.Max)
			imd.Rectangle(0)
			imd.Draw(win)
			win.Update()
			time.Sleep(1000 * time.Millisecond)
			validMap = true
			break
		} else {
			imd.Color = colornames.Yellow
			px := getWall(xCurrent, yCurrent, laberth.SizeField)
			imd.Push(px.Min, px.Max)
			imd.Rectangle(0)
			imd.Draw(win)
			win.Update()
		}
		time.Sleep(10 * time.Millisecond)

		if checkLimit(yCurrent-1, (laberth.FieldDimentionY-2)) && !laberth.ArrayToMap[xCurrent][yCurrent-1] && !laberth.ArrayToCheck[xCurrent][yCurrent-1] {
			queue.PushBack(models.MapPoint{XPoint: xCurrent, YPoint: yCurrent - 1})
		}

		if checkLimit(xCurrent-1, (laberth.FieldDimentionX-1)) && !laberth.ArrayToMap[xCurrent-1][yCurrent] && !laberth.ArrayToCheck[xCurrent-1][yCurrent] {
			queue.PushBack(models.MapPoint{XPoint: xCurrent - 1, YPoint: yCurrent})
		}

		if checkLimit(yCurrent+1, (laberth.FieldDimentionY-2)) && !laberth.ArrayToMap[xCurrent][yCurrent+1] && !laberth.ArrayToCheck[xCurrent][yCurrent+1] {
			queue.PushBack(models.MapPoint{XPoint: xCurrent, YPoint: yCurrent + 1})
		}

		if checkLimit(xCurrent+1, (laberth.FieldDimentionX-1)) && !laberth.ArrayToMap[xCurrent+1][yCurrent] && !laberth.ArrayToCheck[xCurrent+1][yCurrent] {
			queue.PushBack(models.MapPoint{XPoint: xCurrent + 1, YPoint: yCurrent})
		}
	}

	return
}

func checkLimit(currentValue, limit int) bool {
	return currentValue >= 0 && currentValue < limit
}

// func checkMapByDFS(yActual, xActual int, imd *imdraw.IMDraw, win *pixelgl.Window) bool {

// 	// Check vertical limit in the map.
// 	if yActual >= (laberth.FieldDimentionX-1) || yActual < 0 {
// 		return false
// 	}

// 	// Check horizontal limit in the map.
// 	if xActual >= (laberth.FieldDimentionY-2) || xActual < 0 {
// 		return false
// 	}

// 	// This check if this point is playable.
// 	if laberth.ArrayToMap[yActual][xActual] {
// 		return false
// 	}

// 	// Check if this position is already previously checked.
// 	if laberth.ArrayToCheck[yActual][xActual] {
// 		return false
// 	}

// 	laberth.ArrayToCheck[yActual][xActual] = true
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
