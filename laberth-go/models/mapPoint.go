package models

type MapPoint struct {
	TargetInPoint Target
	WallInPoint   bool
}

type Target interface {
	Collision(playerScore int) int
	GetMapPoint() XYMapPoint
}

type XYMapPoint struct {
	XPoint, YPoint int
}

type Coin struct {
	Score        int
	MapPointCoin XYMapPoint
}

type Enemy struct {
	Score         int
	MapPointEnemy XYMapPoint
}

func (c Coin) GetMapPoint() XYMapPoint {
	return c.MapPointCoin
}

func (e Enemy) GetMapPoint() XYMapPoint {
	return e.MapPointEnemy
}

func (c Coin) Collision(playerScore int) int {
	return playerScore + c.Score
}

func (e Enemy) Collision(playerScore int) int {
	return playerScore - e.Score
}
