package models

// Labyrinth struct
type Labyrinth struct {
	SizeField        int
	MovementDistance float32
	ArrayToCheck     [][]bool
	ArrayToMap       [][]MapPointable
}

// SetArrays func
func (labert *Labyrinth) SetArrays(fieldDimentionX, fieldDimentionY int) {
	labert.ArrayToCheck = make([][]bool, fieldDimentionX)
	labert.ArrayToMap = make([][]MapPointable, fieldDimentionX)

	for i := range labert.ArrayToCheck {
		labert.ArrayToCheck[i] = make([]bool, fieldDimentionY)
		labert.ArrayToMap[i] = make([]MapPointable, fieldDimentionY)
	}
}
