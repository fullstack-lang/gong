package models

type FieldInterface interface {
	GetName() string
	GetIndex() int // get the field index with its belonging gongstruct
}
