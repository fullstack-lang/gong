package models

type Part struct {
	Name string

	//gong:text width:300 height:300
	Description string

	Ports []*Port

	// A part is an instance of a system within another system.
	// This is an instance of the whole/part meta model pattern.
	// by default, the part name is the system type name, but it can be overridden by the user.
	TypeOfPart              *System
	IsPartNameNotSystemName bool

	IsControlFlowsNodeExpanded bool
	ControlFlows               []*ControlFlow

	PortWhoseOutControlFlowsNodeIsExpanded []*Port
	PortWhoseInControlFlowsNodeIsExpanded  []*Port

	IsDataFlowsNodeExpanded bool
	inDataFlows             []*DataFlow
	outDataFlows            []*DataFlow

	PortWhoseOutDataFlowsNodeIsExpanded []*Port
	PortWhoseInDataFlowsNodeIsExpanded  []*Port

	owningSystem *System

	PartAnchoredPath []*PartAnchoredPath

	LibraryAbstractFields
	AbstractTypeFields

	IsPortsNodeExpanded bool
}

var _ AbstractType = (*Part)(nil)

func (s *Part) GetOwningSystem() *System {
	return s.owningSystem
}
