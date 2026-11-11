package models

type Part struct {
	Name string

	//gong:text width:300 height:300
	Description string

	Ports []*Port

	// A part is an instance of a system within another system.
	// This is an instance of the whole/part meta model pattern.
	Type                    *System

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

	IsPortsNodeExpanded         bool
	IsTypePartsNodeExpanded     bool
	IsTypeDataFlowsNodeExpanded bool
}

var _ AbstractType = (*Part)(nil)

func (s *Part) GetOwningSystem() *System {
	return s.owningSystem
}
