package models

type Target interface {
	Collision(playerScore int) int
	GetMapPoint() Coords
	SetMapPoint(mapPoint Coords) Target
	GetScore() int
	SetScore(score int) Target
}

type MapPointable interface {
	Step() MapPointable
}

type MapPoint struct {
	TargetInPoint Target
}

type MapBool bool

func (mapBool MapBool) Step() MapPointable {
	return mapBool
}

func (mapPoint MapPoint) Step() MapPointable {
	return mapPoint
}

type Coords struct {
	XPoint, YPoint int
}
