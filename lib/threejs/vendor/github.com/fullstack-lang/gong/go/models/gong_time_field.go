package models

type GongTimeField struct {
	Name string

	Index int

	CompositeStructName string
	AbstractField

	//  is there is a magic gong code with bespoketimeserializeformat "2006-01-02"
	// Date2 time.Time
	BespokeTimeFormat string

	//  is there is a magic gong code with time-form-only
	TimeFormOnly bool
}

func (gongTimeField *GongTimeField) GetIndex() int {
	return gongTimeField.Index
}

func (gongTimeField *GongTimeField) GetCompositeStructName() string {
	return gongTimeField.CompositeStructName
}
