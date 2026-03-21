package models

import (
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Library struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	RootProducts  []*Product
	RootTasks     []*Task
	RootResources []*Resource

	Notes []*Note

	Diagrams []*Diagram

	objects []AbstractType

	SubLibraries []*Library

	NbPixPerCharacter float64 // stored at the root Library only
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

type Diagram struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	ShowPrefix bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	Width  float64
	Height float64

	Product_Shapes              []*ProductShape
	map_Product_ProductShape    map[*Product]*ProductShape
	ProductsWhoseNodeIsExpanded []*Product // to be made private once in production (no need to persist)
	IsPBSNodeExpanded           bool

	ProductComposition_Shapes           []*ProductCompositionShape
	map_Product_ProductCompositionShape map[*Product]*ProductCompositionShape

	IsWBSNodeExpanded bool

	Task_Shapes                    []*TaskShape
	map_Task_TaskShape             map[*Task]*TaskShape
	TasksWhoseNodeIsExpanded       []*Task // to be made private once in production (no need to persist)ExpandableNodeObject
	TasksWhoseInputNodeIsExpanded  []*Task
	TasksWhoseOutputNodeIsExpanded []*Task

	TaskComposition_Shapes        []*TaskCompositionShape
	map_Task_TaskCompositionShape map[*Task]*TaskCompositionShape

	TaskInputShapes         []*TaskInputShape
	map_Task_TaskInputShape map[taskProductKey]*TaskInputShape

	TaskOutputShapes         []*TaskOutputShape
	map_Task_TaskOutputShape map[taskProductKey]*TaskOutputShape

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NoteProductShapes         []*NoteProductShape
	map_Note_NoteProductShape map[noteProductKey]*NoteProductShape

	NoteTaskShapes         []*NoteTaskShape
	map_Note_NoteTaskShape map[noteTaskKey]*NoteTaskShape

	NoteResourceShapes         []*NoteResourceShape
	map_Note_NoteResourceShape map[noteResourceKey]*NoteResourceShape

	Resource_Shapes              []*ResourceShape
	map_Resource_ResourceShape   map[*Resource]*ResourceShape
	ResourcesWhoseNodeIsExpanded []*Resource
	IsResourcesNodeExpanded      bool

	ResourceComposition_Shapes            []*ResourceCompositionShape
	map_Resource_ResourceCompositionShape map[*Resource]*ResourceCompositionShape

	ResourceTaskShapes             []*ResourceTaskShape
	map_Resource_ResourceTaskShape map[resourceTaskKey]*ResourceTaskShape

	map_Product_Rect  map[*Product]*svg.Rect
	map_Task_Rect     map[*Task]*svg.Rect
	map_Note_Rect     map[*Note]*svg.Rect
	map_Resource_Rect map[*Resource]*svg.Rect

	map_SvgRect_ProductShape  map[*svg.Rect]*ProductShape
	map_SvgRect_TaskShape     map[*svg.Rect]*TaskShape
	map_SvgRect_NoteShape     map[*svg.Rect]*NoteShape
	map_SvgRect_ResourceShape map[*svg.Rect]*ResourceShape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

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
)
