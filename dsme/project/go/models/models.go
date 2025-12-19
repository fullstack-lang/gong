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
