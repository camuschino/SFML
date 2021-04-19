package utils

import (
	"math/rand"
	"time"

	"github.com/camuschino/laberth-go/models"
)

func generateRandInt(axisValue int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(axisValue)
}

func generateRandMapPoint(randX, randY int) (mapPoint models.MapPoint) {

	mapPoint.XPoint = generateRandInt(randX)
	mapPoint.YPoint = generateRandInt(randY)

	return
}

func generateValidMapPoint(laberth *models.Labyrinth) (mapPoint models.MapPoint) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	mapPoint = generateRandMapPoint(randX, randY)

	for laberth.ArrayToMap[mapPoint.XPoint][mapPoint.YPoint] {
		mapPoint = generateRandMapPoint(randX, randY)
	}

	return
}

// SetObjectPositions func
func SetObjectPositions(laberth *models.Labyrinth) (player, target models.MapPoint, targets []models.Target) {

	player = generateValidMapPoint(laberth)
	target = generateValidMapPoint(laberth)

	for i := 0; i < 5;i++ {
		var newTarget models.Target

		switch generateRandInt(2) {
		case 0:
			newEnemy := models.Enemy{
				MapPointEnemy: generateValidMapPoint(laberth),
				Score: 20,
			}
			newTarget = newEnemy
		case 1:
			newCoin := models.Coin{
				MapPointCoin: generateValidMapPoint(laberth),
				Score: 10,
			}
			newTarget = newCoin
		}

		targets = append(targets, newTarget)
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

			if i%2 == 0 && j%2 == 0 {
				laberth.ArrayToMap[i][j] = false
			} else {
				if r1.Intn(2) == 0 {
					laberth.ArrayToMap[i][j] = false
				} else {
					laberth.ArrayToMap[i][j] = true
				}
			}
		}
	}

	return
}
