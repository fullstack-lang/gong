package models

// FormNames
type FormNames string

const (
	DefaultForm FormNames = "default form"

	// Name of the form for editing models objects
	ModelForm FormNames = "Model Form"

	// Name of the form for editing shape objects
	ShapeForm FormNames = "Shape Form"
)
