package models

import "time"

// singloton
type Root struct {
	Name string

	Libraries []*Library

	NbPixPerCharacter float64
}

type Library struct {
	Name string

	RootProducts  []*Product
	RootTasks     []*Task
	RootResources []*Resource

	Notes []*Note

	Diagrams []*Diagram

	AbstractTypeFields
	LibraryAbstractFields

	objects []AbstractType

	SubLibraries []*Library
}

type LibraryAbstractFields struct {
	OwningLibrary *Library
}

type LibraryOwnedType interface {
	GetOwnlingLibrary() *Library
	SetOwningLibrary(library *Library)
}

func (r *LibraryAbstractFields) GetOwnlingLibrary() *Library {
	return r.OwningLibrary
}

func (r *LibraryAbstractFields) SetOwningLibrary(library *Library) {
	r.OwningLibrary = library
}

// Note brings information to a diagram
type Note struct {
	//gong:text width:300 height:300
	Name string

	Products  []*Product
	Tasks     []*Task
	Resources []*Resource

	AbstractTypeFields
	LibraryAbstractFields
}

type Task struct {
	Name string

	Start time.Time
	End   time.Time

	//gong:text width:300 height:300
	Description string

	SubTasks []*Task

	AbstractTypeFields
	LibraryAbstractFields

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

	//gong:text width:300 height:300
	Description string

	SubProducts []*Product

	AbstractTypeFields
	LibraryAbstractFields

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

	//gong:text width:300 height:300
	Description string

	Tasks []*Task

	SubResources []*Resource

	// parentResource is a computed field
	// since a Resource belongs to at most one RBS,
	// a parentResource is computed at each UX look. It can be null if the
	// resource is a root resource or an orphaned resource
	parentResource *Resource

	AbstractTypeFields
	LibraryAbstractFields
}

var (
	_ AbstractType = (*Product)(nil)
	_ AbstractType = (*Library)(nil)
	_ AbstractType = (*Task)(nil)
	_ AbstractType = (*Note)(nil)
	_ AbstractType = (*Resource)(nil)
)
