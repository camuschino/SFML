package utils

import (
	"image/color"

	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// RenderMapAndObjects func
func RenderMapAndObjects(laberth *models.Labyrinth, player, target models.MapPoint, targets []models.Target, imd *imdraw.IMDraw, win *pixelgl.Window) {
	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	imd.Color = colornames.Antiquewhite
	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {
			if laberth.ArrayToMap[i][j] {
				px := getWall(i, j, laberth.SizeField)
				imd.Push(px.Min, px.Max)
				imd.Rectangle(0)
			}
		}
	}

	getObjectsToRender(imd, player, colornames.Aqua, laberth)
	getObjectsToRender(imd, target, colornames.Red, laberth)

	for len(targets) > 0 {
		first := targets[0]
		targets = targets[1:]
		switch first.(type) {
			case models.Coin:
				getObjectsToRender(imd, first.GetMapPoint(), colornames.Blue, laberth)
			case models.Enemy:
				getObjectsToRender(imd, first.GetMapPoint(), colornames.Greenyellow, laberth)
		}
	}

	imd.Draw(win)
	win.Update()
}

func getWall(x, y, sizeField int) (px pixel.Rect) {
	posX := float64(x * sizeField)
	posY := float64(y * sizeField)
	px = pixel.R(posX, posY, posX+float64(sizeField), posY+float64(sizeField))
	return
}

func getObjectsToRender(imd *imdraw.IMDraw, object models.MapPoint, color color.Color, laberth *models.Labyrinth) {

	objectToRender := pixel.Vec{
		X: float64((object.XPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
		Y: float64((object.YPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
	}

	imd.Color = color
	imd.Push(objectToRender)
	imd.Circle(float64(laberth.MovementDistance/2), 0)
}

func renderingStep(point models.MapPoint, labertSizeField int, color color.Color, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Color = color
	px := getWall(point.XPoint, point.YPoint, labertSizeField)
	imd.Push(px.Min, px.Max)
	imd.Rectangle(0)
	imd.Draw(win)
	win.Update()
}
