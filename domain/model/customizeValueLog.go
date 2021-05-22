package model

type CustomizeValueLog struct {
	Idx         uint
	UserIdx     uint
	QuestionIdx uint
	Answer      bool

	deleted uint
	created string
	updated string
}
