package models

type Enemy struct {
	Score         int
	MapPointEnemy Coords
}

func (e Enemy) GetMapPoint() Coords {
	return e.MapPointEnemy
}

func (e Enemy) SetMapPoint(mapPoint Coords) Target {
	e.MapPointEnemy = mapPoint
	var t Target = e
	return t
}

func (e Enemy) GetScore() int {
	return e.Score
}

func (e Enemy) SetScore(score int) Target {
	e.Score = score
	var t Target = e
	return t
}

func (e Enemy) Collision(playerScore int) int {
	return playerScore - e.Score
}
