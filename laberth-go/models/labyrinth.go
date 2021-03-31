package models

// Labyrinth struct
type Labyrinth struct {
	SizeField                int
	MovementDistance         float32
	ArrayToCheck, ArrayToMap [][]bool
}

// SetArrays func
func (labert *Labyrinth) SetArrays(fieldDimentionX, fieldDimentionY int) {
	labert.ArrayToCheck = make([][]bool, fieldDimentionX, fieldDimentionX)
	labert.ArrayToMap = make([][]bool, fieldDimentionX, fieldDimentionX)

	for i := range labert.ArrayToCheck {
		labert.ArrayToMap[i] = make([]bool, fieldDimentionY, fieldDimentionY)
		labert.ArrayToCheck[i] = make([]bool, fieldDimentionY, fieldDimentionY)
	}
}
