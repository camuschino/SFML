package utils

import (
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
		result = checkMapByDFS(player, target, laberth, imd, win)
	}
	return
}

func checkMapByBFS(player, target models.MapPoint, laberth models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) (validMap bool) {
	slice := []models.MapPoint{}

	validMap = false

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	upPoint, downPoint, leftPoint, rightPoint := player, player, player, player

	downPoint.YPoint--
	if checkLimit(downPoint.YPoint, fieldDimentionY) && !checkPointIsWall(downPoint, laberth) {
		slice = append(slice, downPoint)
	}

	leftPoint.XPoint--
	if checkLimit(leftPoint.XPoint, fieldDimentionX) && !checkPointIsWall(leftPoint, laberth) {
		slice = append(slice, leftPoint)
	}

	upPoint.YPoint++
	if checkLimit(upPoint.YPoint, fieldDimentionY) && !checkPointIsWall(upPoint, laberth) {
		slice = append(slice, upPoint)
	}

	rightPoint.XPoint++
	if checkLimit(rightPoint.XPoint, fieldDimentionX) && !checkPointIsWall(rightPoint, laberth) {
		slice = append(slice, rightPoint)
	}

	for len(slice) > 0 {
		first := slice[0]

		slice = slice[1:]

		// Check vertical limit in the map.
		if !checkLimit(first.YPoint, fieldDimentionY) {
			continue
		}

		// Check horizontal limit in the map.
		if !checkLimit(first.XPoint, fieldDimentionX) {
			continue
		}

		// This check if this point is playable. (true means false, because there's a wall)
		if !checkMapPoint(first, laberth) {
			continue
		}

		laberth.ArrayToCheck[first.XPoint][first.YPoint] = true

		if target == first {
			renderingStep(first, laberth.SizeField, colornames.Blue, imd, win)
			time.Sleep(1000 * time.Millisecond)
			validMap = true
			break
		}

		renderingStep(first, laberth.SizeField, colornames.Greenyellow, imd, win)

		time.Sleep(10 * time.Millisecond)

		leftPoint = first
		leftPoint.YPoint--

		if checkLimit(leftPoint.YPoint, fieldDimentionY) && checkMapPoint(leftPoint, laberth) {
			slice = append(slice, leftPoint)
		}

		downPoint = first
		downPoint.XPoint--

		if checkLimit(downPoint.XPoint, fieldDimentionX) && checkMapPoint(downPoint, laberth) {
			slice = append(slice, downPoint)
		}

		upPoint = first
		upPoint.YPoint++

		if checkLimit(upPoint.YPoint, fieldDimentionY) && checkMapPoint(upPoint, laberth) {
			slice = append(slice, upPoint)
		}

		rightPoint = first
		rightPoint.XPoint++

		if checkLimit(rightPoint.XPoint, fieldDimentionX) && checkMapPoint(rightPoint, laberth) {
			slice = append(slice, rightPoint)
		}
	}

	return
}

func checkLimit(currentValue, limit int) bool {
	return currentValue >= 0 && currentValue < limit-1
}

func checkMapPoint(point models.MapPoint, laberth models.Labyrinth) bool {
	return !checkPointIsWall(point, laberth) && !checkPointIsAlreadyTested(point, laberth)
}

func checkPointIsWall(point models.MapPoint, laberth models.Labyrinth) bool {
	return laberth.ArrayToMap[point.XPoint][point.YPoint]
}

func checkPointIsAlreadyTested(point models.MapPoint, laberth models.Labyrinth) bool {
	return laberth.ArrayToCheck[point.XPoint][point.YPoint]
}

func checkMapByDFS(actualPoint, target models.MapPoint, laberth models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) bool {

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	yActual := actualPoint.YPoint
	xActual := actualPoint.XPoint

	// Check vertical limit in the map.
	if !checkLimit(xActual, fieldDimentionX) {
		return false
	}

	// Check horizontal limit in the map.
	if !checkLimit(yActual, fieldDimentionY) {
		return false
	}

	// This check if this point is playable.
	if laberth.ArrayToMap[xActual][yActual] {
		return false
	}

	// Check if this position is already previously checked.
	if laberth.ArrayToCheck[xActual][yActual] {
		return false
	}

	laberth.ArrayToCheck[xActual][yActual] = true

	if target == actualPoint {
		imd.Color = colornames.Red
		px := getWall(xActual, yActual, laberth.SizeField)
		imd.Push(px.Min, px.Max)
		imd.Rectangle(0)
		imd.Draw(win)
		win.Update()
		time.Sleep(1000 * time.Millisecond)
		return true
	}

	imd.Color = colornames.Yellow
	px := getWall(xActual, yActual, laberth.SizeField)
	imd.Push(px.Min, px.Max)
	imd.Rectangle(0)
	imd.Draw(win)
	win.Update()

	time.Sleep(10 * time.Millisecond)

	if checkMapByDFS(models.MapPoint{XPoint: xActual - 1, YPoint: yActual}, target, laberth, imd, win) {
		return true
	}

	if checkMapByDFS(models.MapPoint{XPoint: xActual, YPoint: yActual - 1}, target, laberth, imd, win) {
		return true
	}

	if checkMapByDFS(models.MapPoint{XPoint: xActual + 1, YPoint: yActual}, target, laberth, imd, win) {
		return true
	}

	if checkMapByDFS(models.MapPoint{XPoint: xActual, YPoint: yActual + 1}, target, laberth, imd, win) {
		return true
	}

	return false
}
