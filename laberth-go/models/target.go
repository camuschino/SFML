package models

type Target interface {
	Collision(playerScore int) int
	GetMapPoint() MapPoint
}

type Coin struct {
	Score int
	MapPointCoin MapPoint
}

type Enemy struct {
	Score int
	MapPointEnemy MapPoint
}

func (c Coin) GetMapPoint() MapPoint {
	return c.MapPointCoin
}

func (c Enemy) GetMapPoint() MapPoint {
	return c.MapPointEnemy
}

func (c Coin) Collision(playerScore int) int {
	return playerScore - c.Score
}

func (e Enemy) Collision(playerScore int) int {
	return playerScore - e.Score
}