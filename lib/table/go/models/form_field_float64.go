package models

type FormFieldFloat64 struct {
	Name  string
	Value float64

	HasMinValidator bool
	MinValue        float64

	HasMaxValidator bool
	MaxValue        float64
}
