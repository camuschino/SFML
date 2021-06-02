package utils

import (
	"time"

	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	target                                      models.Coords
	sizeField, fieldDimentionX, fieldDimentionY int
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

	fieldDimentionX = len(laberth.ArrayToMap)
	fieldDimentionY = len(laberth.ArrayToMap[0])

	switch algorithm {
	case "DFS":
		seeker.checkMapByBFS(player, target, laberth, imd, win)
		// algh.numberOfRounds = algh.numberOfRounds + score
		return 0
	case "BFS":
	default:
		score = seeker.checkMapByBFS(player, target, laberth, imd, win)
		algh.numberOfRounds = algh.numberOfRounds + score
	}
	return
}

func (algh AlgorithmsSearching) checkMapByBFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) (score int) {
	slice := []models.Coords{}

	upPoint, downPoint, leftPoint, rightPoint := player, player, player, player

	upPoint.YPoint++
	if CheckLimit(upPoint.YPoint, fieldDimentionY) && !CheckPointIsWall(upPoint, laberth) {
		slice = append(slice, upPoint)
	}

	rightPoint.XPoint++
	if CheckLimit(rightPoint.XPoint, fieldDimentionX) && !CheckPointIsWall(rightPoint, laberth) {
		slice = append(slice, rightPoint)
	}

	downPoint.YPoint--
	if CheckLimit(downPoint.YPoint, fieldDimentionY) && !CheckPointIsWall(downPoint, laberth) {
		slice = append(slice, downPoint)
	}

	leftPoint.XPoint--
	if CheckLimit(leftPoint.XPoint, fieldDimentionX) && !CheckPointIsWall(leftPoint, laberth) {
		slice = append(slice, leftPoint)
	}

	var first models.Coords

	for ; len(slice) > 0; slice = slice[1:] {
		first = slice[0]

		// This check if this point is playable. (true means false, because there's a wall)
		if !CheckMapPoint(first, laberth) {
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

		renderingStep(first, laberth.SizeField, colornames.Greenyellow, imd, win)
		CheckTargetPosition(win, imd, laberth, &target)

		time.Sleep(15 * time.Millisecond)

		upPoint, downPoint, leftPoint, rightPoint := first, first, first, first

		upPoint.YPoint++
		if CheckLimit(upPoint.YPoint, fieldDimentionY) && CheckMapPoint(upPoint, laberth) {
			slice = append(slice, upPoint)
		}

		rightPoint.XPoint++
		if CheckLimit(rightPoint.XPoint, fieldDimentionX) && CheckMapPoint(rightPoint, laberth) {
			slice = append(slice, rightPoint)
		}

		leftPoint.YPoint--
		if CheckLimit(leftPoint.YPoint, fieldDimentionY) && CheckMapPoint(leftPoint, laberth) {
			slice = append(slice, leftPoint)
		}

		downPoint.XPoint--
		if CheckLimit(downPoint.XPoint, fieldDimentionX) && CheckMapPoint(downPoint, laberth) {
			slice = append(slice, downPoint)
		}
	}

	return 0
}

func (algh AlgorithmsSearching) checkMapByDFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, score int) models.MapBool {

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	// Check vertical limit in the map.
	if !CheckLimit(player.XPoint, fieldDimentionX) {
		return false
	}

	// Check horizontal limit in the map.
	if !CheckLimit(player.YPoint, fieldDimentionY) {
		return false
	}

	// This check if this point is playable, AND
	// Check if this position is already previously checked.
	if !CheckMapPoint(player, laberth) {
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

func CheckLimit(currentValue, limit int) bool {
	return currentValue >= 0 && currentValue < limit-1
}

func CheckMapPoint(point models.Coords, laberth *models.Labyrinth) bool {
	return !CheckPointIsWall(point, laberth) && !CheckPointIsAlreadyTested(point, laberth)
}

func CheckPointIsWall(point models.Coords, laberth *models.Labyrinth) bool {

	switch mapPointale := laberth.ArrayToMap[point.XPoint][point.YPoint].(type) {
	case models.MapBool:
		return bool(mapPointale)
	default:
		return false
	}
}

func CheckPointIsAlreadyTested(point models.Coords, laberth *models.Labyrinth) bool {
	return laberth.ArrayToCheck[point.XPoint][point.YPoint]
}
