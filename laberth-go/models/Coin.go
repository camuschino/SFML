package models

type Coin struct {
	Score        int
	MapPointCoin Coords
}

func (c Coin) GetMapPoint() Coords {
	return c.MapPointCoin
}

func (c Coin) SetMapPoint(mapPoint Coords) Target {
	c.MapPointCoin = mapPoint
	var t Target = c
	return t
}

func (c Coin) GetScore() int {
	return c.Score
}

func (c Coin) SetScore(score int) Target {
	c.Score = score
	var t Target = c
	return t
}

func (c Coin) Collision(playerScore int) int {
	return playerScore + c.Score
}
