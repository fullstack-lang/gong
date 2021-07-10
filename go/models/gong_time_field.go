package models

type GongTimeField struct {
	Name string

	Index int
}

func (gongTimeField *GongTimeField) GetName() string {
	return gongTimeField.Name
}

func (gongTimeField *GongTimeField) GetIndex() int {
	return gongTimeField.Index
}
