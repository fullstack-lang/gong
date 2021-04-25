package models

type GongTimeField struct {
	Name string
}

func (gongTimeField *GongTimeField) GetName() string {
	return gongTimeField.Name
}
