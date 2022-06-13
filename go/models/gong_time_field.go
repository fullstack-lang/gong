package models

type GongTimeField struct {
	Name string

	Index int

	CompositeStructName string
}

func (gongTimeField *GongTimeField) GetIndex() int {
	return gongTimeField.Index
}

func (gongTimeField *GongTimeField) GetCompositeStructName() string {
	return gongTimeField.CompositeStructName
}
