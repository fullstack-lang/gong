package models

type Task struct {
	Name string

	ParentTask *Task
}

type Product struct {
	Name string

	ParentProduct *Product

	IsExpanded bool
}

type Project struct {
	Name string

	RootTasks    []*Task
	RootProducts []*Product
}

// singloton
type Root struct {
	Name string

	Projects []*Project
}
