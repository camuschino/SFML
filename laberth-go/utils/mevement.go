package utils

import (
	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel/pixelgl"
)

func Running(win *pixelgl.Window, labeth models.Labyrinth, target models.Coords) {
	// if win.JustPressed(pixelgl.KeyUp) {
	// 	target.YPoint++
	// 	println(2)
	// 	if CheckLimit(target.YPoint, fieldDimentionY) && !CheckPointIsWall(target, laberth) {
	// 		getObjectsToRender(imd, target, colornames.Black, laberth)
	// 		target.YPoint += sizeField
	// 		getObjectsToRender(imd, target, colornames.Red, laberth)
	// 		win.Update()
	// 	}
	// }

	// if win.JustPressed(pixelgl.KeyDown) && !bool(laberth.ArrayToMap[target.XPoint][target.YPoint-1].(models.MapBool)) {
	// 	getObjectsToRender(imd, target, colornames.Black, laberth)
	// 	target.YPoint -= sizeField
	// 	getObjectsToRender(imd, target, colornames.Red, laberth)
	// }

	// if win.JustPressed(pixelgl.KeyRight) && !bool(laberth.ArrayToMap[target.XPoint-1][target.YPoint].(models.MapBool)) {
	// 	getObjectsToRender(imd, target, colornames.Black, laberth)
	// 	target.XPoint += sizeField
	// 	getObjectsToRender(imd, target, colornames.Red, laberth)
	// }

	// if win.JustPressed(pixelgl.KeyLeft) && !bool(laberth.ArrayToMap[target.XPoint+1][target.YPoint].(models.MapBool)) {
	// 	getObjectsToRender(imd, target, colornames.Black, laberth)
	// 	target.XPoint -= sizeField
	// 	getObjectsToRender(imd, target, colornames.Red, laberth)
	// }
}
