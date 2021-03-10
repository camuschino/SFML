package utils

import (
	"github.com/camuschino/laberth-go/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// RenderMapAndObjects func
func RenderMapAndObjects(laberth models.Labyrinth, player, target models.MapPoint, imd *imdraw.IMDraw, win *pixelgl.Window) {

	imd.Color = colornames.Antiquewhite
	for i := 0; i < laberth.FieldDimentionX-1; i++ {
		for j := 0; j < laberth.FieldDimentionY-2; j++ {
			if laberth.ArrayToMap[i][j] {
				px := getWall(i, j, laberth.SizeField)
				imd.Push(px.Min, px.Max)
				imd.Rectangle(0)
			}
		}
	}

	pyVec := getObjectsToRender(player, laberth)
	tarVec := getObjectsToRender(target, laberth)

	imd.Color = colornames.Aqua
	imd.Push(pyVec)
	imd.Circle(float64(laberth.MovementDistance/2), 0)

	imd.Color = colornames.Greenyellow
	imd.Push(tarVec)
	imd.Circle(float64(laberth.MovementDistance/2), 0)
	imd.Draw(win)
	win.Update()
}

func getWall(x, y, sizeField int) (px pixel.Rect) {
	posX := float64(x * sizeField)
	posY := float64(y * sizeField)
	px = pixel.R(posX, posY, posX+float64(sizeField), posY+float64(sizeField))
	return
}

func getObjectsToRender(object models.MapPoint, laberth models.Labyrinth) (objectToRender pixel.Vec) {

	objectToRender.X = float64((object.XPoint * laberth.SizeField) + int(laberth.MovementDistance/2))
	objectToRender.Y = float64((object.YPoint * laberth.SizeField) + int(laberth.MovementDistance/2))

	return
}
