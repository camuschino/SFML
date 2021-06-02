package utils

import (
	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	coorClone models.Coords
)

func CheckTargetPosition(win *pixelgl.Window, imd *imdraw.IMDraw, laberth *models.Labyrinth, target *models.Coords) {

	if win.JustPressed(pixelgl.KeyUp) {
		coorClone = *target
		coorClone.YPoint++
		if CheckLimit(coorClone.YPoint, len(laberth.ArrayToMap[0])) && CheckMapPoint(coorClone, laberth) {
			drawObjectToRender(imd, *target, colornames.Black, laberth)
			target.YPoint++
			drawObjectToRender(imd, *target, colornames.Red, laberth)
		}
	}

	if win.JustPressed(pixelgl.KeyDown) {
		coorClone = *target
		coorClone.YPoint--
		if CheckLimit(coorClone.YPoint, len(laberth.ArrayToMap[0])) && CheckMapPoint(coorClone, laberth) {
			drawObjectToRender(imd, *target, colornames.Black, laberth)
			target.YPoint--
			drawObjectToRender(imd, *target, colornames.Red, laberth)
		}
	}

	if win.JustPressed(pixelgl.KeyRight) {
		coorClone = *target
		coorClone.XPoint++
		if CheckLimit(coorClone.XPoint, len(laberth.ArrayToMap)) && CheckMapPoint(coorClone, laberth) {
			drawObjectToRender(imd, *target, colornames.Black, laberth)
			target.XPoint++
			drawObjectToRender(imd, *target, colornames.Red, laberth)
		}
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		coorClone = *target
		coorClone.XPoint--
		if CheckLimit(coorClone.XPoint, len(laberth.ArrayToMap)) && CheckMapPoint(coorClone, laberth) {
			drawObjectToRender(imd, *target, colornames.Black, laberth)
			target.XPoint--
			drawObjectToRender(imd, *target, colornames.Red, laberth)
		}
	}
}
