package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowDimentionX, windowDimentionY int     = 500, 500 // window dimention
	sizeBlock                          int     = 20       // large: 100. medium: 50, little: 20, nano: 10
	fieldDimentionX, fieldDimentionY   int     = ((windowDimentionX / sizeBlock) * 2) + 1, ((windowDimentionY / sizeBlock) * 2) + 2
	sizeField                          int     = sizeBlock / 2
	movementDistance                   float32 = float32(sizeField)
)

var (
	arrayToCheck                                                                            [fieldDimentionX][fieldDimentionY]bool
	arrayToMap                                                                              [fieldDimentionX][fieldDimentionY]bool
	playerPositionX, playerPositionY, objectivePositionX, objectivePositionY                int
	mIsMovingUp, mIsMovingRight, mIsMovingLeft, mIsMovingDown, isMapChecked, isMapGenerated bool
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Laberth",
		Bounds: pixel.R(0, 0, float64(windowDimentionY), float64(windowDimentionX)),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	createNewMap()
	setObjectPositions()

	renderMapAndObjects(imd, win)

	for true {
		a := !checkMapByDFS(playerPositionY, playerPositionX, imd, win)
		println(a)
		createNewMap()
		setObjectPositions()
		win.Clear(colornames.Black)
		imd.Clear()
		renderMapAndObjects(imd, win)
	}

	println("FIN")

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}

func createNewMap() {

	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < fieldDimentionX-1; i++ {
		for j := 0; j < fieldDimentionY-2; j++ {
			arrayToCheck[i][j] = false
			r1 := rand.New(s1)

			if i%2 == 0 && j%2 == 0 {
				arrayToMap[i][j] = false
			} else {
				if r1.Intn(2) == 0 {
					arrayToMap[i][j] = false
				} else {
					arrayToMap[i][j] = true
				}
			}
		}
	}

	isMapChecked = false
}

func renderMapAndObjects(imd *imdraw.IMDraw, win *pixelgl.Window) {

	imd.Color = colornames.Antiquewhite
	for i := 0; i < fieldDimentionX-1; i++ {
		for j := 0; j < fieldDimentionY-2; j++ {
			if arrayToMap[i][j] {
				px := getWall(i, j)
				imd.Push(px.Min, px.Max)
				imd.Rectangle(0)
			}
		}
	}

	pyVec, objVec := getObjects()

	imd.Color = colornames.Aqua
	imd.Push(pyVec)
	imd.Circle(float64(movementDistance/2), 0)

	imd.Color = colornames.Greenyellow
	imd.Push(objVec)
	imd.Circle(float64(movementDistance/2), 0)
	imd.Draw(win)
	win.Update()
}

func getWall(x, y int) (px pixel.Rect) {
	posX := float64(x * sizeField)
	posY := float64(y * sizeField)
	px = pixel.R(posY, posX, posY+float64(sizeField), posX+float64(sizeField))
	return
}

func getObjects() (pyVec, objVec pixel.Vec) {

	pyX := float64((playerPositionX * sizeField) + int(movementDistance/2))
	pyVec.X = pyX

	pyY := float64((playerPositionY * sizeField) + int(movementDistance/2))
	pyVec.Y = pyY

	objX := float64((objectivePositionX * sizeField) + int(movementDistance/2))
	objVec.X = objX

	objY := float64((objectivePositionY * sizeField) + int(movementDistance/2))
	objVec.Y = objY
	return
}

func setObjectPositions() {

	xPlayer := 0
	yPlayer := 0
	xObj := 0
	yObj := 0
	validPosition := false

	randX := fieldDimentionX - 1
	randY := fieldDimentionY - 2

	s1 := rand.NewSource(time.Now().UnixNano())

	for !validPosition {
		r1 := rand.New(s1)
		xPlayer = r1.Intn(randX)
		yPlayer = r1.Intn(randY)

		if !arrayToMap[yPlayer][xPlayer] {
			validPosition = true
			playerPositionX = xPlayer
			playerPositionY = yPlayer
		}
	}

	s1 = rand.NewSource(time.Now().UnixNano())
	validPosition = false
	for !validPosition {
		r1 := rand.New(s1)

		xObj = r1.Intn(randX)
		yObj = r1.Intn(randY)

		if !arrayToMap[yObj][xObj] {
			validPosition = true
			objectivePositionX = xObj
			objectivePositionY = yObj
		}
	}
}

func checkMapByDFS(yActual, xActual int, imd *imdraw.IMDraw, win *pixelgl.Window) bool {

	// Check vertical limit in the map.
	if yActual >= (fieldDimentionX-1) || yActual < 0 {
		return false
	}

	// Check horizontal limit in the map.
	if xActual >= (fieldDimentionY-2) || xActual < 0 {
		return false
	}

	// This check if this point is playable.
	if arrayToMap[yActual][xActual] {
		return false
	}

	// Check if this position is already previously checked.
	if arrayToCheck[yActual][xActual] {
		return false
	}

	arrayToCheck[yActual][xActual] = true
	println(xActual, yActual)

	// Check if this point is the point where is the objective.
	if !(playerPositionX == xActual && playerPositionY == yActual) {
		imd.Color = colornames.Yellow
		px := getWall(yActual, xActual)
		imd.Push(px.Min, px.Max)
		imd.Rectangle(0)
		imd.Draw(win)
		win.Update()
		time.Sleep(10 * time.Millisecond)
	}

	// Check if this point is the point where is the objective.
	if objectivePositionX == xActual && objectivePositionY == yActual {
		println("EXITO")
		return true
	}

	if checkMapByDFS(yActual-1, xActual, imd, win) {
		return true
	}

	if checkMapByDFS(yActual, xActual-1, imd, win) {
		return true
	}

	if checkMapByDFS(yActual+1, xActual, imd, win) {
		return true
	}

	if checkMapByDFS(yActual, xActual+1, imd, win) {
		return true
	}

	return false
}

func main() {
	pixelgl.Run(run)
}
