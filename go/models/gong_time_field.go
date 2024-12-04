package models

type GongTimeField struct {
	Name string

	Index int

	CompositeStructName string

	//  is there is a magic gong code with bespoketimeserializeformat "2006-01-02"
	// Date2 time.Time
	BespokeTimeFormat string
}

func (gongTimeField *GongTimeField) GetIndex() int {
	return gongTimeField.Index
}

func (gongTimeField *GongTimeField) GetCompositeStructName() string {
	return gongTimeField.CompositeStructName
}
