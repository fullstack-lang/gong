package models

// singloton
type Root struct {
	Name string

	Projects []*Project

	NbPixPerCharacter float64
}

type Project struct {
	Name string

	RootProducts      []*Product
	IsPBSNodeExpanded bool

	RootTasks         []*Task
	IsWBSNodeExpanded bool

	Diagrams               []*Diagram
	IsDiagramsNodeExpanded bool

	Notes               []*Note
	IsNotesNodeExpanded bool

	ExpandableNodeObject
}

type ExpandableNodeObject struct {
	IsExpanded bool // to be made private once in production (no need to persist)

	// ComputedPrefix is automaticaly computed by the semantic enforcing mechanism
	ComputedPrefix string
	computedPrefix []int

	// When the full PBS is displayed, the computedWidth is the number of node
	// aligned below. A leaf node has a computedWidth of 1
	computedWidth int
}

// Note brings information to a diagram
type Note struct {
	//gong:text width:300 height:300
	Name string

	Products []*Product

	Tasks []*Task

	IsExpanded bool
}

// GetComputedPrefix implements [AbstractType].
func (note *Note) GetComputedPrefix() string {
	return ""
}

// GetComputedWidth implements [AbstractType].
func (note *Note) GetComputedWidth() int {
	return 0
}

// GetIsExpanded implements [AbstractType].
func (note *Note) GetIsExpanded() bool {
	return note.IsExpanded
}

// SetComputedPrefix implements [AbstractType].
func (note *Note) SetComputedPrefix(string) {
}

// SetIsExpanded implements [AbstractType].
func (note *Note) SetIsExpanded(val bool) {
	note.IsExpanded = val
}

// SetComputedWidth implements [AbstractType].
func (note *Note) SetComputedWidth(int) {
}

func (r *ExpandableNodeObject) GetComputedWidth() int {
	return r.computedWidth
}

type Task struct {
	Name string

	//gong:text width:300 height:300
	Description string

	SubTasks []*Task

	ExpandableNodeObject

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
// - They are Directed Acyclic Graph (DAG)
// - A [models.Product]/[models.Task] belongs to at most one PBS/WBS.
// Those invariants allow prefix and parent to be computed at each UX loop
const NoteSemantic = ""

type Product struct {
	Name string

	//gong:text width:300 height:300
	Description string

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

	// parentProduct is a computed field
	// since a Product belongs to at most one WBS,
	// a parentProduct is computed at each UX look. It can be null if the
	// product is a root product or an orphaned product
	parentProduct *Product
}

type AbstractType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
	GetComputedPrefix() string
	SetComputedPrefix(string)
	GetComputedWidth() int
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

var (
	_ AbstractType = (*Product)(nil)
	_ AbstractType = (*Project)(nil)
	_ AbstractType = (*Task)(nil)
	_ AbstractType = (*Note)(nil)
)
