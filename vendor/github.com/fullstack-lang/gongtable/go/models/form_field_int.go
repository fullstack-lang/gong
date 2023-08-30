package models

type FormFieldInt struct {
	Name  string
	Value int

	HasMinValidator bool
	MinValue        int

	HasMaxValidator bool
	MaxValue        int
}
