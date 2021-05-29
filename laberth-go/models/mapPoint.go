package models

type Target interface {
	Collision(playerScore int) int
	GetMapPoint() Coords
	SetMapPoint(mapPoint Coords) Target
	GetScore() int
	SetScore(score int) Target
}

type MapPointable interface{}

type MapPoint struct {
	TargetInPoint Target
}

type MapBool bool

type Coords struct {
	XPoint, YPoint int
}
