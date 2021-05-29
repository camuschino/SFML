package utils

import (
	"time"

	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	target    models.Coords
	sizeField int
)

type LabSolver interface {
	checkMapByBFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) int
	checkMapByDFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, score int) models.MapBool
}

type AlgorithmsSearching struct {
	numberOfRounds int
}

// ValidateMap function which works fine
func ValidateMap(algorithm string, player, targetOriginal models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, sizeFieldOriginal int) (score int) {
	var seeker LabSolver
	algh := AlgorithmsSearching{}
	seeker = algh

	target = targetOriginal
	sizeField = sizeFieldOriginal

	switch algorithm {
	case "BFS":
		score = seeker.checkMapByBFS(player, target, laberth, imd, win)
		algh.numberOfRounds = algh.numberOfRounds + score
	case "DFS":
		seeker.checkMapByDFS(player, target, laberth, imd, win, 0)
		// algh.numberOfRounds = algh.numberOfRounds + score
		return 0
	}
	return
}

func (algh AlgorithmsSearching) checkMapByBFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) (score int) {
	slice := []models.Coords{}

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

		// This check if this point is playable. (true means false, because there's a wall)
		if !checkMapPoint(first, laberth) {
			continue
		}

		laberth.ArrayToCheck[first.XPoint][first.YPoint] = true

		if target == first {
			renderingStep(first, laberth.SizeField, colornames.Blue, imd, win)
			time.Sleep(1000 * time.Millisecond)
			return score
		}

		switch mapPointable := laberth.ArrayToMap[first.XPoint][first.YPoint].(type) {
		case models.MapPoint:
			score = mapPointable.TargetInPoint.Collision(score)
			println(score)
		}

		if win.JustPressed(pixelgl.KeyUp) && !bool(laberth.ArrayToMap[target.XPoint][target.YPoint+1].(models.MapBool)) {
			getObjectsToRender(imd, target, colornames.Black, laberth)
			target.YPoint += sizeField
			getObjectsToRender(imd, target, colornames.Red, laberth)
		}

		if win.JustPressed(pixelgl.KeyDown) && !bool(laberth.ArrayToMap[target.XPoint][target.YPoint-1].(models.MapBool)) {
			getObjectsToRender(imd, target, colornames.Black, laberth)
			target.YPoint -= sizeField
			getObjectsToRender(imd, target, colornames.Red, laberth)
		}

		if win.JustPressed(pixelgl.KeyRight) && !bool(laberth.ArrayToMap[target.XPoint-1][target.YPoint].(models.MapBool)) {
			getObjectsToRender(imd, target, colornames.Black, laberth)
			target.XPoint += sizeField
			getObjectsToRender(imd, target, colornames.Red, laberth)
		}

		if win.JustPressed(pixelgl.KeyLeft) && !bool(laberth.ArrayToMap[target.XPoint+1][target.YPoint].(models.MapBool)) {
			getObjectsToRender(imd, target, colornames.Black, laberth)
			target.XPoint -= sizeField
			getObjectsToRender(imd, target, colornames.Red, laberth)
		}

		renderingStep(first, laberth.SizeField, colornames.Greenyellow, imd, win)

		time.Sleep(1 * time.Millisecond)

		upPoint, downPoint, leftPoint, rightPoint := first, first, first, first

		leftPoint.YPoint--
		if checkLimit(leftPoint.YPoint, fieldDimentionY) && checkMapPoint(leftPoint, laberth) {
			slice = append(slice, leftPoint)
		}

		downPoint.XPoint--
		if checkLimit(downPoint.XPoint, fieldDimentionX) && checkMapPoint(downPoint, laberth) {
			slice = append(slice, downPoint)
		}

		upPoint.YPoint++
		if checkLimit(upPoint.YPoint, fieldDimentionY) && checkMapPoint(upPoint, laberth) {
			slice = append(slice, upPoint)
		}

		rightPoint.XPoint++
		if checkLimit(rightPoint.XPoint, fieldDimentionX) && checkMapPoint(rightPoint, laberth) {
			slice = append(slice, rightPoint)
		}
	}

	return 0
}

func (algh AlgorithmsSearching) checkMapByDFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, score int) models.MapBool {

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	// Check vertical limit in the map.
	if !checkLimit(player.XPoint, fieldDimentionX) {
		return false
	}

	// Check horizontal limit in the map.
	if !checkLimit(player.YPoint, fieldDimentionY) {
		return false
	}

	// This check if this point is playable, AND
	// Check if this position is already previously checked.
	if !checkMapPoint(player, laberth) {
		return false
	}

	laberth.ArrayToCheck[player.XPoint][player.YPoint] = true

	switch mapPointable := laberth.ArrayToMap[player.XPoint][player.YPoint].(type) {
	case models.MapPoint:
		score = mapPointable.TargetInPoint.Collision(score)
		println(score)
	}

	if target == player {
		renderingStep(player, laberth.SizeField, colornames.Blue, imd, win)
		time.Sleep(1000 * time.Millisecond)
		return true
	}

	renderingStep(player, laberth.SizeField, colornames.Greenyellow, imd, win)

	time.Sleep(10 * time.Millisecond)

	leftPoint := player
	leftPoint.XPoint--
	if algh.checkMapByDFS(leftPoint, target, laberth, imd, win, score) {
		return true
	}

	downPoint := player
	downPoint.YPoint--
	if algh.checkMapByDFS(downPoint, target, laberth, imd, win, score) {
		return true
	}

	rightPoint := player
	rightPoint.XPoint++
	if algh.checkMapByDFS(rightPoint, target, laberth, imd, win, score) {
		return true
	}

	upPoint := player
	upPoint.YPoint++

	return algh.checkMapByDFS(upPoint, target, laberth, imd, win, score)
}

func checkLimit(currentValue, limit int) bool {
	return currentValue >= 0 && currentValue < limit-1
}

func checkMapPoint(point models.Coords, laberth *models.Labyrinth) bool {
	return !checkPointIsWall(point, laberth) && !checkPointIsAlreadyTested(point, laberth)
}

func checkPointIsWall(point models.Coords, laberth *models.Labyrinth) bool {

	switch mapPointale := laberth.ArrayToMap[point.XPoint][point.YPoint].(type) {
	case models.MapBool:
		return bool(mapPointale)
	default:
		return false
	}
}

func checkPointIsAlreadyTested(point models.Coords, laberth *models.Labyrinth) bool {
	return laberth.ArrayToCheck[point.XPoint][point.YPoint]
}
