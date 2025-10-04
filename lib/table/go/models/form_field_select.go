package models

type FormFieldSelect struct {
	Name  string
	Value *Option

	Options []*Option

	CanBeEmpty bool

	// PreserveInitialOrder overides default alphabetical
	// ordering of Options in the pulldown menu
	PreserveInitialOrder bool
}
