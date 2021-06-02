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
func RenderMapAndObjects(laberth *models.Labyrinth, player, target models.Coords, imd *imdraw.IMDraw, win *pixelgl.Window) {
	fieldDimentionX := len(laberth.ArrayToMap)
	fieldDimentionY := len(laberth.ArrayToMap[0])

	imd.Color = colornames.Antiquewhite
	for i := 0; i < fieldDimentionX; i++ {
		for j := 0; j < fieldDimentionY; j++ {

			switch mapPointable := laberth.ArrayToMap[i][j].(type) {
			case models.MapBool:
				if mapPointable {
					px := getWall(i, j, laberth.SizeField)
					imd.Color = colornames.White
					imd.Push(px.Min, px.Max)
					imd.Rectangle(0)
				}
			case models.MapPoint:
				target := mapPointable.TargetInPoint
				switch target.(type) {
				case models.Coin:
					drawObjectToRender(imd, target.GetMapPoint(), colornames.Blue, laberth)
				case models.Enemy:
					drawObjectToRender(imd, target.GetMapPoint(), colornames.Greenyellow, laberth)
				}
			}
		}
	}

	drawObjectToRender(imd, player, colornames.Aqua, laberth)
	drawObjectToRender(imd, target, colornames.Red, laberth)

	imd.Draw(win)
	win.Update()
}

func getWall(x, y, sizeField int) (px pixel.Rect) {
	posX := float64(x * sizeField)
	posY := float64(y * sizeField)
	px = pixel.R(posX, posY, posX+float64(sizeField), posY+float64(sizeField))
	return
}

func drawObjectToRender(imd *imdraw.IMDraw, object models.Coords, color color.Color, laberth *models.Labyrinth) {

	objectToRender := pixel.Vec{
		X: float64((object.XPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
		Y: float64((object.YPoint * laberth.SizeField) + int(laberth.MovementDistance/2)),
	}

	imd.Color = color
	imd.Push(objectToRender)
	imd.Circle(float64(laberth.MovementDistance/2), 0)
}

func renderingStep(point models.Coords, labertSizeField int, color color.Color, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Color = color
	px := getWall(point.XPoint, point.YPoint, labertSizeField)
	imd.Push(px.Min, px.Max)
	imd.Rectangle(0)
	imd.Draw(win)
	win.Update()
}
