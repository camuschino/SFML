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

func generateRandMapPoint(randX, randY int) (mapPoint models.XYMapPoint) {

	mapPoint.XPoint = generateRandInt(randX)
	mapPoint.YPoint = generateRandInt(randY)

	return
}

func generateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.XYMapPoint) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	mapPoint = generateRandMapPoint(randX, randY)

	mapPointable := laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint]

	for mapPointable.WallInPoint {
		mapPoint = generateRandMapPoint(randX, randY)
		mapPointable.WallInPoint = laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].WallInPoint
	}

	return
}

// SetObjectPositions func
func SetObjectPositions(laberth *models.Labyrinth) (player, target models.XYMapPoint) {

	player = generateValidMapPoint(laberth)
	target = generateValidMapPoint(laberth)

	for i := 0; i < 10; i++ {
		switch generateRandInt(2) {
		case 0:
			mapPoint := generateValidMapPoint(laberth)
			newEnemy := models.Enemy{
				Score:         20,
				MapPointEnemy: mapPoint,
			}
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].TargetInPoint = newEnemy
		case 1:
			mapPoint := generateValidMapPoint(laberth)
			newCoin := models.Coin{
				Score:        10,
				MapPointCoin: mapPoint,
			}
			laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint].TargetInPoint = newCoin
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

	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {
			r1 := rand.New(s1)

			newMapPoint := models.MapPoint{}

			if i%2 == 0 && j%2 == 0 {
				newMapPoint.WallInPoint = false
				newMapPoint.TargetInPoint = nil

				laberth.ArrayToMap[i][j] = newMapPoint
			} else {
				if r1.Intn(2) == 0 {
					newMapPoint.WallInPoint = false
					newMapPoint.TargetInPoint = nil

					laberth.ArrayToMap[i][j] = newMapPoint
				} else {
					newMapPoint.WallInPoint = true
					newMapPoint.TargetInPoint = nil

					laberth.ArrayToMap[i][j] = newMapPoint
				}
			}
		}
	}

	return
}
