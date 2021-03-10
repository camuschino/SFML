package models

// Labyrinth struct
type Labyrinth struct {
	SizeField                int
	MovementDistance         float32
	ArrayToCheck, ArrayToMap [][]bool
}

// SetArrays func
func (labert *Labyrinth) SetArrays(fieldDimentionX, fieldDimentionY int) {
	labert.ArrayToCheck = make([][]bool, fieldDimentionX)
	labert.ArrayToMap = make([][]bool, fieldDimentionX)

	for i := range labert.ArrayToCheck {
		labert.ArrayToMap[i] = make([]bool, fieldDimentionY)
		labert.ArrayToCheck[i] = make([]bool, fieldDimentionY)
	}
}
