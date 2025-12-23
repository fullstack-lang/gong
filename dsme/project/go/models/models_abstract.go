package models

// singloton
type Root struct {
	Name string

	Projects []*Project

	// product that do not belong to projects
	OrphanedProducts []*Product

	// task that do not belong to projects
	OrphanedTasks []*Task

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

	ExpandableNodeObject
}

type ExpandableNodeObject struct {
	IsExpanded bool // to be made private once in production (no need to persist)

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

	// parentTask is a computed field
	// since a Task belongs to at most one WBS,
	// a parentTask is computed at each UX look. It can be null if the
	// task is a root task or an orphaned task
	parentTask *Task
}

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

type ProjectElementType interface {
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

var (
	_ ProjectElementType = (*Product)(nil)
	_ ProjectElementType = (*Project)(nil)
	_ ProjectElementType = (*Task)(nil)
)
