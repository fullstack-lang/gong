package models

type Project struct {
	Name string

	RootTasks    []*Task
	RootProducts []*Product

	ExpandableNodeObject
}

// singloton
type Root struct {
	Name string

	Projects []*Project

	// product that do not belong to projects
	OrphanedProducts []*Product

	// task that do not belong to projects
	OrphanedTasks []*Task
}

type ExpandableNodeObject struct {
	IsExpanded bool

	// ComputedPrefix is automaticaly computed by the semantic enforcing mechanism
	ComputedPrefix string
}
type Task struct {
	Name string

	SubTasks []*Task

	ExpandableNodeObject

	Inputs               []*Product
	IsInputsNodeExpanded bool

	Outputs               []*Product
	IsOutputsNodeExpanded bool
}

type Product struct {
	Name string

	SubProducts []*Product

	ExpandableNodeObject

	// producers are computed from [models.Task.Outputs]
	// this is a computed field, therefore, not exported
	producers               []*Task
	IsProducersNodeExpanded bool

	// consumers are computed from [models.Task.Inputs]
	// this is a computed field, therefore, not exported
	consumers               []*Task
	IsConsumersNodeExpanded bool
}

type NodeType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
	GetComputedPrefix() string
	SetComputedPrefix(string)
}

func (r *ExpandableNodeObject) GetIsExpanded() bool {
	return r.IsExpanded
}

func (r *ExpandableNodeObject) SetIsExpanded(isExpanded bool) {
	r.IsExpanded = isExpanded
}

func (r *ExpandableNodeObject) GetComputedPrefix() string {
	return r.ComputedPrefix
}

func (r *ExpandableNodeObject) SetComputedPrefix(ComputedPrefix string) {
	r.ComputedPrefix = ComputedPrefix
}

var _ NodeType = (*Product)(nil)
var _ NodeType = (*Project)(nil)
var _ NodeType = (*Task)(nil)
