package models

type Port struct {
	Name string

	//gong:text width:300 height:300
	Description string

	outControlFlows []*ControlFlow
	inControlFlows  []*ControlFlow

	outDataFlows []*DataFlow
	inDataFlows  []*DataFlow

	owningPart *Part

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Port)(nil)

func (s *Port) GetOwningPart() *Part {
	return s.owningPart
}
