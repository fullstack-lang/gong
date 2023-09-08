package models

type FormFieldSelect struct {
	Name  string
	Value *Option

	Options []*Option

	CanBeEmpty bool
}
