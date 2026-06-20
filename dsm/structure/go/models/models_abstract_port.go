package models

type Port struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsStartPort bool
	IsEndPort   bool

	outControlFlows []*ControlFlow
	inControlFlows  []*ControlFlow

	outDataFlows []*DataFlow
	inDataFlows  []*DataFlow

	// A port is an instance of a system within another system.
	// This is an instance of the whole/part meta model pattern.
	// by default, the port name is the system type name, but it can be overridden by the user.
	Type                     *System
	IsPortNameNotSystemName bool

	owningPart *Part
}

var _ AbstractType = (*Port)(nil)

func (s *Port) GetOwningPart() *Part {
	return s.owningPart
}
