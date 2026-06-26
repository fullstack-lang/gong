package models

// common fields among fields (ah ah)
type AbstractField struct {

	// some field mark the start of an accordion //gong:accordion-start //gong:accordion-end
	IsAccordionStart bool
	IsAccordionEnd   bool
}

type FieldInterface interface {
	GetName() string
	GetIndex() int // get the field index with its belonging gongstruct

	// GetCompositeStructName provides, if it is from a composite struct, the name of the composite struct
	//  where the field is present.
	// It is empty otherwise
	GetCompositeStructName() string
}
