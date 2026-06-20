package models

type Part struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsPortsNodeExpanded bool
	Ports               []*Port

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
}

var _ AbstractType = (*Part)(nil)

func (s *Part) GetOwningSystem() *System {
	return s.owningSystem
}
