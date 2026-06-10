package models

import (
	"time"
)

// Note brings information to a diagram
type Note struct {
	//gong:text width:300 height:300
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Products  []*Product
	Tasks     []*Task
	Resources []*Resource
}

type Task struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsImport       bool
	ReferencedTask *Task

	Start time.Time
	End   time.Time

	//gong:text width:300 height:300
	Description string

	SubTasks []*Task

	Inputs               []*Product
	IsInputsNodeExpanded bool

	Outputs               []*Product
	IsOutputsNodeExpanded bool

	// parentTask is a computed field
	// since a Task belongs to at most one WBS,
	// a parentTask is computed at each UX look. It can be null if the
	// task is a root task or an orphaned task
	parentTask *Task

	// Completion Management
	IsWithCompletion bool
	Completion       CompletionEnum
}

// CompletionEnum
type CompletionEnum string

const (
	PERCENT_100 CompletionEnum = "100 %"
	PERCENT_075 CompletionEnum = "75 %"
	PERCENT_050 CompletionEnum = "50 %"
	PERCENT_025 CompletionEnum = "25 %"
	PERCENT_000 CompletionEnum = "0 %"
)

// A TaskGroup is a set of task.
// A task can belong to multiple task groups.
// A TaskGroup is used to be displayed in a single lane when the diagram is a time diagram.
type TaskGroup struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Tasks []*Task
}

// GONGDOC(NoteSemantic): Note on the models semantic
//
// [models.Product] and [models.Task] are in Product Breakdown Structure (PBS)
// and Work Breakdown Structure (WBS)
// PBS/WBS have 2 invariants that are enforced at each UX loop:
// - They are Trees
// - A [models.Product]/[models.Task] belongs to at most one PBS/WBS.
// Those invariants allow prefix and parent to be computed at each UX loop
const NoteSemantic = ""

type Product struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsImport          bool
	ReferencedProduct *Product

	//gong:text width:300 height:300
	Description string

	SubProducts []*Product

	// producers are computed from [models.Task.Outputs]
	// this is a computed field, therefore, not exported
	producers               []*Task
	IsProducersNodeExpanded bool

	// consumers are computed from [models.Task.Inputs]
	// this is a computed field, therefore, not exported
	consumers               []*Task
	IsConsumersNodeExpanded bool

	// parentProduct is a computed field
	// since a Product belongs to at most one WBS,
	// a parentProduct is computed at each UX look. It can be null if the
	// product is a root product or an orphaned product
	parentProduct *Product
}

type Resource struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsImport           bool
	ReferencedResource *Resource

	//gong:text width:300 height:300
	Description string

	Tasks []*Task

	SubResources []*Resource

	// parentResource is a computed field
	// since a Resource belongs to at most one RBS,
	// a parentResource is computed at each UX look. It can be null if the
	// resource is a root resource or an orphaned resource
	parentResource *Resource
}

var (
	_ AbstractType = (*Product)(nil)
	_ AbstractType = (*Library)(nil)
	_ AbstractType = (*Task)(nil)
	_ AbstractType = (*Note)(nil)
	_ AbstractType = (*Resource)(nil)
	_ AbstractType = (*TaskGroup)(nil)
)
