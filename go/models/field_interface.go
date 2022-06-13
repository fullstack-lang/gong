package models

type FieldInterface interface {
	GetName() string
	GetIndex() int // get the field index with its belonging gongstruct

	// GetCompositeStructName provides, if it is from a composite struct, the name of the composite struct
	//  where the field is present.
	// It is empty otherwise
	GetCompositeStructName() string
}
