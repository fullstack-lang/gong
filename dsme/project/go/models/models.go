package models

type Project struct {
	Name string

	RootTasks    []*Task
	RootProducts []*Product

	IsExpanded bool
}

// singloton
type Root struct {
	Name string

	Projects []*Project

	// product that do not belong to projects
	OrphanedProducts []*Product
}

type Task struct {
	Name string

	ParentTask *Task
}

type Product struct {
	Name string

	SubProducts []*Product

	IsExpanded bool
}

type NodeType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
}

func (product *Product) GetIsExpanded() bool {
	return product.IsExpanded
}

func (product *Product) SetIsExpanded(isExpanded bool) {
	product.IsExpanded = isExpanded
}

var _ NodeType = (*Product)(nil)

func (project *Project) GetIsExpanded() bool {
	return project.IsExpanded
}

func (project *Project) SetIsExpanded(isExpanded bool) {
	project.IsExpanded = isExpanded
}

var _ NodeType = (*Project)(nil)
