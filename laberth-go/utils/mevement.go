package utils

import (
	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	nextPosition models.Coords
)

func moveTarget(nextPosition models.Coords, target *models.Coords, positionToCheck, limit int, imd *imdraw.IMDraw, laberth *models.Labyrinth) {

	if CheckLimit(positionToCheck, limit) && CheckMapPoint(nextPosition, laberth) {
		DrawObjectToRender(imd, *target, colornames.Black, laberth)
		*target = nextPosition
	}
}

func CheckTargetPosition(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {
	imd.Clear()

	if win.JustPressed(pixelgl.KeyUp) {
		nextPosition = *target
		nextPosition.YPoint++
		moveTarget(nextPosition, target, nextPosition.YPoint, len(laberth.ArrayToMap[0]), imd, laberth)
	}

	if win.JustPressed(pixelgl.KeyDown) {
		nextPosition = *target
		nextPosition.YPoint--
		moveTarget(nextPosition, target, nextPosition.YPoint, len(laberth.ArrayToMap[0]), imd, laberth)
	}

	if win.JustPressed(pixelgl.KeyRight) {
		nextPosition = *target
		nextPosition.XPoint++
		moveTarget(nextPosition, target, nextPosition.XPoint, len(laberth.ArrayToMap[0]), imd, laberth)
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		nextPosition = *target
		nextPosition.XPoint--
		moveTarget(nextPosition, target, nextPosition.XPoint, len(laberth.ArrayToMap[0]), imd, laberth)
	}

	DrawObjectToRender(imd, *target, colornames.Red, laberth)
	imd.Draw(win)
}
