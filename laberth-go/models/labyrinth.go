package models

// Labyrinth struct
type Labyrinth struct {
	FieldDimentionX, FieldDimentionY, SizeField int
	MovementDistance                            float32
	ArrayToCheck, ArrayToMap                    [][]bool
}
