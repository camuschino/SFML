package utils

import (
	"sync"
	"time"

	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	fieldDimentionX, fieldDimentionY int
)

// ValidateMap function which works fine
func ValidateMap(algorithm string, player models.Coords, target *models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, sizeFieldOriginal int) (score int) {

	fieldDimentionX = len(laberth.ArrayToMap)
	fieldDimentionY = len(laberth.ArrayToMap[0])

	switch algorithm {
	case "DFS":
		checkMapByBFS(player, target, laberth, imd, win)
		// algh.numberOfRounds = algh.numberOfRounds + score
		return 0
	case "BFS":
	default:
		checkMapByBFS(player, target, laberth, imd, win)
	}
	return
}

func forBFS(slice []models.Coords, target *models.Coords, mutex *sync.Mutex, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) {

	var first models.Coords

	var score int

	first = slice[0]

	for ; len(slice) > 0; first, slice = slice[0], slice[1:] {

		// This check if this point is playable. (true means false, because there's a wall)
		if !CheckMapPoint(first, laberth) {
			continue
		}

		laberth.ArrayToCheck[first.XPoint][first.YPoint] = true

		if *target == first {
			RenderingStep(first, laberth, colornames.Blue, mutex, imd, win)
			time.Sleep(1000 * time.Millisecond)
			laberth.Over <- true
		}

		switch mapPointable := laberth.ArrayToMap[first.XPoint][first.YPoint].(type) {
		case models.MapPoint:
			score = mapPointable.TargetInPoint.Collision(score)
			println(score)
		}

		RenderingStep(first, laberth, colornames.Greenyellow, mutex, imd, win)

		getCoordsSlice(first, &slice, laberth)
	}

	laberth.Over <- false
}

func getCoordsSlice(player models.Coords, slice *[]models.Coords, laberth *models.Labyrinth) {

	upPoint, downPoint, leftPoint, rightPoint := player, player, player, player

	upPoint.YPoint++
	if CheckLimit(upPoint.YPoint, fieldDimentionY) && !CheckPointIsWall(upPoint, laberth) {
		*slice = append(*slice, upPoint)
	}

	rightPoint.XPoint++
	if CheckLimit(rightPoint.XPoint, fieldDimentionX) && !CheckPointIsWall(rightPoint, laberth) {
		*slice = append(*slice, rightPoint)
	}

	downPoint.YPoint--
	if CheckLimit(downPoint.YPoint, fieldDimentionY) && !CheckPointIsWall(downPoint, laberth) {
		*slice = append(*slice, downPoint)
	}

	leftPoint.XPoint--
	if CheckLimit(leftPoint.XPoint, fieldDimentionX) && !CheckPointIsWall(leftPoint, laberth) {
		*slice = append(*slice, leftPoint)
	}
}

func checkMapByBFS(player models.Coords, target *models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window) {

	mutex := &sync.Mutex{}
	var slice []models.Coords
	var laberthOver bool = false
	time.Sleep(2500 * time.Millisecond)

	for {
		getCoordsSlice(player, &slice, laberth)
		if len(slice) == 0 {
			continue
		}
		go func() {
			laberthOver = <-laberth.Over
		}()
		if !laberthOver {
			go forBFS(slice, target, mutex, laberth, imd, win)
		} else {
			println("FINISHED!!!")
			break
		}

		time.Sleep(500 * time.Millisecond)
		player = GenerateValidMapPoint(laberth)
	}
}

// func checkMapByDFS(player, target models.Coords, laberth *models.Labyrinth, imd *imdraw.IMDraw, win *pixelgl.Window, score int) models.MapBool {

// 	fieldDimentionX := len(laberth.ArrayToMap)
// 	fieldDimentionY := len(laberth.ArrayToMap[0])

// 	// Check vertical limit in the map.
// 	if !CheckLimit(player.XPoint, fieldDimentionX) {
// 		return false
// 	}

// 	// Check horizontal limit in the map.
// 	if !CheckLimit(player.YPoint, fieldDimentionY) {
// 		return false
// 	}

// 	// This check if this point is playable, AND
// 	// Check if this position is already previously checked.
// 	if !CheckMapPoint(player, laberth) {
// 		return false
// 	}

// 	laberth.ArrayToCheck[player.XPoint][player.YPoint] = true

// 	switch mapPointable := laberth.ArrayToMap[player.XPoint][player.YPoint].(type) {
// 	case models.MapPoint:
// 		score = mapPointable.TargetInPoint.Collision(score)
// 		println(score)
// 	}

// 	if target == player {
// 		RenderingStep(player, laberth.SizeField, colornames.Blue, imd, win)
// 		time.Sleep(1000 * time.Millisecond)
// 		return true
// 	}

// 	RenderingStep(player, laberth.SizeField, colornames.Greenyellow, imd, win)

// 	time.Sleep(10 * time.Millisecond)

// 	leftPoint := player
// 	leftPoint.XPoint--
// 	if checkMapByDFS(leftPoint, target, laberth, imd, win, score) {
// 		return true
// 	}

// 	downPoint := player
// 	downPoint.YPoint--
// 	if checkMapByDFS(downPoint, target, laberth, imd, win, score) {
// 		return true
// 	}

// 	rightPoint := player
// 	rightPoint.XPoint++
// 	if checkMapByDFS(rightPoint, target, laberth, imd, win, score) {
// 		return true
// 	}

// 	upPoint := player
// 	upPoint.YPoint++

// 	return checkMapByDFS(upPoint, target, laberth, imd, win, score)
// }

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
