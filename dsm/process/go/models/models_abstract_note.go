package models

type Note struct {
	//gong:text width:300 height:300
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsTasksNodeExpanded bool
	Tasks               []*Task

	IsDataFlowsNodeExpanded bool
	DataFlows               []*DataFlow

	IsControlFlowsNodeExpanded bool
	ControlFlows               []*ControlFlow
}

var _ AbstractType = (*Note)(nil)
