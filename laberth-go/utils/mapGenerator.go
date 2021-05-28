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

func generateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.Coords) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	mapPoint = generateRandMapPoint(randX, randY)

	mapPointable := laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint]

	for mapPointable.(models.MapBool) {
		mapPoint = generateRandMapPoint(randX, randY)
		mapPointable = laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].(models.MapBool)
	}

	return
}

func getNewTarget(mapPoint models.Coords, newTarget models.Target) models.MapPointable {
	var t models.MapPoint

	switch newTarget.(type) {
	case models.Enemy:
		t.TargetInPoint = newTarget.SetScore(20).SetMapPoint(mapPoint)
	case models.Coin:
		t.TargetInPoint = newTarget.SetScore(10).SetMapPoint(mapPoint)
	default:
		t.TargetInPoint = newTarget.SetScore(10).SetMapPoint(mapPoint)
	}
	return t
}

// SetObjectPositions func
func SetObjectPositions(laberth *models.Labyrinth) (player, target models.Coords) {

	player = generateValidMapPoint(laberth)
	target = generateValidMapPoint(laberth)

	for i := 0; i < 10; i++ {
		mapPoint := generateValidMapPoint(laberth)
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
func CreateNewMap(xSize, ySize, sizeField int, movementDistance float32) (laberth models.Labyrinth) {

	s1 := rand.NewSource(time.Now().UnixNano())

	laberth = models.Labyrinth{
		SizeField:        sizeField,
		MovementDistance: movementDistance,
	}

	laberth.SetArrays(xSize, ySize)

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])
	var newMapBool models.MapBool
	r1 := rand.New(s1)

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

	return
}
