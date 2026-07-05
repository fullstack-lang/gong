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
	TreeAbstractTypeFields

	Products  []*Product
	Tasks     []*Task
	Resources []*Resource
}

type Task struct {
	Name string

	//gong:text width:300 height:300
	Description string

	//gong:time-form-only
	Start time.Time

	//gong:time-form-only
	End time.Time

	//gong:accordion-start "Duration"
	DurationYears  float64
	DurationMonths float64
	DurationWeeks  float64
	DurationDays   float64
	DurationHours  float64
	//gong:accordion-end
	IsEndDateComputedFromDuration bool

	//gong:accordion-start "Predecessors"
	Predecessors []*Task
	//gong:accordion-end
	IsStartDateComputedFromPredecessors bool

	IsMilestone bool

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
	//gong:accordion-start "Completion Display"
	IsWithCompletion bool
	//gong:accordion-end
	Completion CompletionEnum

	// DisplayVerticalBar indicates wether the task
	// has a vertical vertical on the whole gantt when it is a milestone
	DisplayVerticalBar bool

	// a red diamond a text anchor will be displayed
	TaskGroupsToDisplay []*TaskGroup

	//gong:accordion-start "Custom positions"
	TextPosition TextPositionEnum
	XOffset      float64
	//gong:accordion-end
	YOffset float64

	IsImport       bool
	ReferencedTask *Task

	SubTasks []*Task

	LibraryAbstractFields
	AbstractTypeFields
	TreeAbstractTypeFields
}

// TextPositionEnum
type TextPositionEnum string

const (
	TEXT_POSITION_TOP    TextPositionEnum = "POSITION_TOP"
	TEXT_POSITION_BOTTOM TextPositionEnum = "POSITION_BOTTOM"
	TEXT_POSITION_LEFT   TextPositionEnum = "POSITION_LEFT"
	TEXT_POSITION_RIGHT  TextPositionEnum = "POSITION_RIGHT"
	TEXT_POSITION_CENTER TextPositionEnum = "POSITION_CENTER"
)

// CompletionEnum
type CompletionEnum string

const (
	PERCENT_100 CompletionEnum = "100 %"
	PERCENT_075 CompletionEnum = "75 %"
	PERCENT_050 CompletionEnum = "50 %"
	PERCENT_025 CompletionEnum = "25 %"
	PERCENT_000 CompletionEnum = "0 %"
)

// TimeStepScaleEnum
type TimeStepScaleEnum string

const (
	YEARS  TimeStepScaleEnum = "YEARS"
	MONTHS TimeStepScaleEnum = "MONTHS"
	WEEKS  TimeStepScaleEnum = "WEEKS"
	DAYS   TimeStepScaleEnum = "DAYS"
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

	IsImport          bool
	ReferencedProduct *Product

	LibraryAbstractFields
	AbstractTypeFields
	TreeAbstractTypeFields
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

	LibraryAbstractFields
	AbstractTypeFields
	TreeAbstractTypeFields

	IsImport           bool
	ReferencedResource *Resource
}

var (
	_ AbstractType = (*Product)(nil)
	_ AbstractType = (*Library)(nil)
	_ AbstractType = (*Task)(nil)
	_ AbstractType = (*Note)(nil)
	_ AbstractType = (*Resource)(nil)
	_ AbstractType = (*TaskGroup)(nil)
)
