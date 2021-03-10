package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/camuschino/laberth-go/models"
)

// SetObjectPositions func
func SetObjectPositions(laberth models.Labyrinth) (player, target models.MapPoint) {

	randX := len(laberth.ArrayToMap) - 2
	randY := len(laberth.ArrayToMap[0]) - 2

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	player.XPoint = r1.Intn(randX)
	player.YPoint = r1.Intn(randY)

	for laberth.ArrayToMap[player.XPoint][player.YPoint] {
		r1 = rand.New(s1)
		player.XPoint = r1.Intn(randX)
		player.YPoint = r1.Intn(randY)
	}

	target.XPoint = r1.Intn(randX)
	target.YPoint = r1.Intn(randY)

	for laberth.ArrayToMap[target.XPoint][target.YPoint] {
		r1 = rand.New(s1)
		target.XPoint = r1.Intn(randX)
		target.YPoint = r1.Intn(randY)
	}
	fmt.Println(player)
	fmt.Println(target)

	return
}

// CreateNewMap func
func CreateNewMap(xSize, ySize, sizeField int, movementDistance float32) (laberth models.Labyrinth) {

	s1 := rand.NewSource(time.Now().UnixNano())

	laberth = models.Labyrinth{
		// FieldDimentionX:  xSize,
		// FieldDimentionY:  ySize,
		SizeField:        sizeField,
		MovementDistance: movementDistance,
	}

	laberth.SetArrays(xSize, ySize)

	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	for i := 0; i < fieldDimentionX-1; i++ {
		for j := 0; j < fieldDimentionY-2; j++ {
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
