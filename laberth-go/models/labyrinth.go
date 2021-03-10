package models

// Labyrinth struct
type Labyrinth struct {
	FieldDimentionX, FieldDimentionY, SizeField int
	MovementDistance                            float32
	ArrayToCheck, ArrayToMap                    [][]bool
}

// SetArrays func
func (labert *Labyrinth) SetArrays() {
	labert.ArrayToCheck = make([][]bool, labert.FieldDimentionY)
	labert.ArrayToMap = make([][]bool, labert.FieldDimentionY)

	for i := range labert.ArrayToCheck {
		labert.ArrayToMap[i] = make([]bool, labert.FieldDimentionX)
		labert.ArrayToCheck[i] = make([]bool, labert.FieldDimentionX)
	}
}
