package models

type GongTimeField struct {
	Name string

	Index int
}

func (gongTimeField *GongTimeField) GetIndex() int {
	return gongTimeField.Index
}
