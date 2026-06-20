package models

type Part struct {
	Name string

	// When the part has one or more system/resource
	// it might be conveniant to consider it, semanticaly/visualy, as
	// the system/resource itself.
	// IsNameOverridenWithFirstAllocatedSystemResource is true in this case
	IsSystemResource bool

	//gong:text width:300 height:300
	Description string

	// Resources is the list of resources used by the part
	Resources               []*Resource
	IsResourcesNodeExpanded bool

	// Systemess is the list of systemes performed by the part
	Systemes               []*System
	IsSystemesNodeExpanded bool

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
