package models

// Labyrinth struct
type Labyrinth struct {
	SizeField        int
	MovementDistance float32
	ArrayToCheck     [][]bool
	ArrayToMap       [][]MapPoint
}

// SetArrays func
func (labert *Labyrinth) SetArrays(fieldDimentionX, fieldDimentionY int) {
	labert.ArrayToCheck = make([][]bool, fieldDimentionX)
	labert.ArrayToMap = make([][]MapPoint, fieldDimentionX)

	for i := range labert.ArrayToCheck {
		labert.ArrayToMap[i] = make([]MapPoint, fieldDimentionY)
		labert.ArrayToCheck[i] = make([]bool, fieldDimentionY)
	}
}
