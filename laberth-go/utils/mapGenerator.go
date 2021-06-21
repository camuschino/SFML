package utils

import (
	"math/rand"
	"time"

	"github.com/camuschino/laberth-go/models"
)

func generateRandInt(randLimit int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(randLimit)
}

func generateRandMapPoint(randX, randY int) models.Coords {
	return models.Coords{
		XPoint: generateRandInt(randX),
		YPoint: generateRandInt(randY),
	}
}

func GenerateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.Coords) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	for {
		mapPoint = generateRandMapPoint(randX, randY)
		switch mapPointable := laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].(type) {
		case models.MapBool:
			if mapPointable {
				continue
			}
		case models.MapPointable:
		default:
			continue
		}
		break
	}

	return
}

func getNewTarget(mapPoint models.Coords, newTarget models.Target) models.MapPointable {
	var score int

	switch newTarget.(type) {
	case models.Enemy:
		score = 20
	case models.Coin:
	default:
		score = 10
	}

	return models.MapPoint{
		TargetInPoint: newTarget.SetScore(score).SetMapPoint(mapPoint),
	}
}

// SetObjectPositions func
func SetObjectPositions(laberth *models.Labyrinth) (player, target models.Coords) {

	player = GenerateValidMapPoint(laberth)
	target = GenerateValidMapPoint(laberth)
	var mapPoint models.Coords

	for i := 0; i < 10; i++ {
		mapPoint = GenerateValidMapPoint(laberth)
		switch generateRandInt(2) {
		case 0:
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] = getNewTarget(mapPoint, models.Enemy{})
		case 1:
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] = getNewTarget(mapPoint, models.Coin{})
		}
	}

	return
}

// CreateNewMap func
func CreateNewLabyrinth(laberth *models.Labyrinth) {
	var newMapBool models.MapBool
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {
			if (i%2 == 0 && j%2 == 0) || r1.Intn(2) == 0 {
				newMapBool = false
			} else {
				newMapBool = true
			}
			laberth.ArrayToMap[i][j] = newMapBool
		}
	}
	laberth.ArrayToCheck = make([][]bool, fieldDimentionX)
	for i := range laberth.ArrayToCheck {
		laberth.ArrayToCheck[i] = make([]bool, fieldDimentionY)
	}
}
