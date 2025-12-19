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
}

type Task struct {
	Name string

	SubTasks []*Task

	ExpandableNodeObject
}

type Product struct {
	Name string

	// ComputedPrefix is automaticaly computed by the semantic enforcing mechanism
	ComputedPrefix string

	SubProducts []*Product

	ExpandableNodeObject
}

type NodeType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
}

func (r *ExpandableNodeObject) GetIsExpanded() bool {
	return r.IsExpanded
}

func (r *ExpandableNodeObject) SetIsExpanded(isExpanded bool) {
	r.IsExpanded = isExpanded
}

var _ NodeType = (*Product)(nil)
var _ NodeType = (*Project)(nil)
