package models

type ControlFlow struct {
	Name string

	Start *Task

	End *Task
}
