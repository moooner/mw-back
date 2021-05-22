package model

type Question struct {
	Idx          uint
	QuestionText string
	Delta        float64

	deleted uint
	created string
	updated string
}
