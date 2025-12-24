// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	project_go "github.com/fullstack-lang/gong/dsme/project/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs
var _ = strings.Clone("")

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")
var _ = errUnkownEnum

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void
var _ = __member

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Diagrams           map[*Diagram]struct{}
	Diagrams_mapString map[string]*Diagram

	// insertion point for slice of pointers maps
	Diagram_Product_Shapes_reverseMap map[*ProductShape]*Diagram

	Diagram_ProductsWhoseNodeIsExpanded_reverseMap map[*Product]*Diagram

	Diagram_ProductComposition_Shapes_reverseMap map[*ProductCompositionShape]*Diagram

	Diagram_Task_Shapes_reverseMap map[*TaskShape]*Diagram

	Diagram_TasksWhoseNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhoseInputNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TaskComposition_Shapes_reverseMap map[*TaskCompositionShape]*Diagram

	Diagram_TaskInputShapes_reverseMap map[*TaskInputShape]*Diagram

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Products           map[*Product]struct{}
	Products_mapString map[string]*Product

	// insertion point for slice of pointers maps
	Product_SubProducts_reverseMap map[*Product]*Product

	OnAfterProductCreateCallback OnAfterCreateInterface[Product]
	OnAfterProductUpdateCallback OnAfterUpdateInterface[Product]
	OnAfterProductDeleteCallback OnAfterDeleteInterface[Product]
	OnAfterProductReadCallback   OnAfterReadInterface[Product]

	ProductCompositionShapes           map[*ProductCompositionShape]struct{}
	ProductCompositionShapes_mapString map[string]*ProductCompositionShape

	// insertion point for slice of pointers maps
	OnAfterProductCompositionShapeCreateCallback OnAfterCreateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeUpdateCallback OnAfterUpdateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeDeleteCallback OnAfterDeleteInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeReadCallback   OnAfterReadInterface[ProductCompositionShape]

	ProductShapes           map[*ProductShape]struct{}
	ProductShapes_mapString map[string]*ProductShape

	// insertion point for slice of pointers maps
	OnAfterProductShapeCreateCallback OnAfterCreateInterface[ProductShape]
	OnAfterProductShapeUpdateCallback OnAfterUpdateInterface[ProductShape]
	OnAfterProductShapeDeleteCallback OnAfterDeleteInterface[ProductShape]
	OnAfterProductShapeReadCallback   OnAfterReadInterface[ProductShape]

	Projects           map[*Project]struct{}
	Projects_mapString map[string]*Project

	// insertion point for slice of pointers maps
	Project_RootProducts_reverseMap map[*Product]*Project

	Project_RootTasks_reverseMap map[*Task]*Project

	Project_Diagrams_reverseMap map[*Diagram]*Project

	OnAfterProjectCreateCallback OnAfterCreateInterface[Project]
	OnAfterProjectUpdateCallback OnAfterUpdateInterface[Project]
	OnAfterProjectDeleteCallback OnAfterDeleteInterface[Project]
	OnAfterProjectReadCallback   OnAfterReadInterface[Project]

	Roots           map[*Root]struct{}
	Roots_mapString map[string]*Root

	// insertion point for slice of pointers maps
	Root_Projects_reverseMap map[*Project]*Root

	Root_OrphanedProducts_reverseMap map[*Product]*Root

	Root_OrphanedTasks_reverseMap map[*Task]*Root

	OnAfterRootCreateCallback OnAfterCreateInterface[Root]
	OnAfterRootUpdateCallback OnAfterUpdateInterface[Root]
	OnAfterRootDeleteCallback OnAfterDeleteInterface[Root]
	OnAfterRootReadCallback   OnAfterReadInterface[Root]

	Tasks           map[*Task]struct{}
	Tasks_mapString map[string]*Task

	// insertion point for slice of pointers maps
	Task_SubTasks_reverseMap map[*Task]*Task

	Task_Inputs_reverseMap map[*Product]*Task

	Task_Outputs_reverseMap map[*Product]*Task

	OnAfterTaskCreateCallback OnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback OnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback OnAfterDeleteInterface[Task]
	OnAfterTaskReadCallback   OnAfterReadInterface[Task]

	TaskCompositionShapes           map[*TaskCompositionShape]struct{}
	TaskCompositionShapes_mapString map[string]*TaskCompositionShape

	// insertion point for slice of pointers maps
	OnAfterTaskCompositionShapeCreateCallback OnAfterCreateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeUpdateCallback OnAfterUpdateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeDeleteCallback OnAfterDeleteInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeReadCallback   OnAfterReadInterface[TaskCompositionShape]

	TaskInputShapes           map[*TaskInputShape]struct{}
	TaskInputShapes_mapString map[string]*TaskInputShape

	// insertion point for slice of pointers maps
	OnAfterTaskInputShapeCreateCallback OnAfterCreateInterface[TaskInputShape]
	OnAfterTaskInputShapeUpdateCallback OnAfterUpdateInterface[TaskInputShape]
	OnAfterTaskInputShapeDeleteCallback OnAfterDeleteInterface[TaskInputShape]
	OnAfterTaskInputShapeReadCallback   OnAfterReadInterface[TaskInputShape]

	TaskShapes           map[*TaskShape]struct{}
	TaskShapes_mapString map[string]*TaskShape

	// insertion point for slice of pointers maps
	OnAfterTaskShapeCreateCallback OnAfterCreateInterface[TaskShape]
	OnAfterTaskShapeUpdateCallback OnAfterUpdateInterface[TaskShape]
	OnAfterTaskShapeDeleteCallback OnAfterDeleteInterface[TaskShape]
	OnAfterTaskShapeReadCallback   OnAfterReadInterface[TaskShape]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	DiagramOrder            uint
	DiagramMap_Staged_Order map[*Diagram]uint

	ProductOrder            uint
	ProductMap_Staged_Order map[*Product]uint

	ProductCompositionShapeOrder            uint
	ProductCompositionShapeMap_Staged_Order map[*ProductCompositionShape]uint

	ProductShapeOrder            uint
	ProductShapeMap_Staged_Order map[*ProductShape]uint

	ProjectOrder            uint
	ProjectMap_Staged_Order map[*Project]uint

	RootOrder            uint
	RootMap_Staged_Order map[*Root]uint

	TaskOrder            uint
	TaskMap_Staged_Order map[*Task]uint

	TaskCompositionShapeOrder            uint
	TaskCompositionShapeMap_Staged_Order map[*TaskCompositionShape]uint

	TaskInputShapeOrder            uint
	TaskInputShapeMap_Staged_Order map[*TaskInputShape]uint

	TaskShapeOrder            uint
	TaskShapeMap_Staged_Order map[*TaskShape]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func (stage *Stage) GetReference() map[GongstructIF]GongstructIF {
	return stage.reference
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Diagram:
		tmp := GetStructInstancesByOrder(stage.Diagrams, stage.DiagramMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Diagram implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Product:
		tmp := GetStructInstancesByOrder(stage.Products, stage.ProductMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Product implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProductCompositionShape:
		tmp := GetStructInstancesByOrder(stage.ProductCompositionShapes, stage.ProductCompositionShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProductCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProductShape:
		tmp := GetStructInstancesByOrder(stage.ProductShapes, stage.ProductShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProductShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Project:
		tmp := GetStructInstancesByOrder(stage.Projects, stage.ProjectMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Project implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Root:
		tmp := GetStructInstancesByOrder(stage.Roots, stage.RootMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Root implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Task:
		tmp := GetStructInstancesByOrder(stage.Tasks, stage.TaskMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Task implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskCompositionShape:
		tmp := GetStructInstancesByOrder(stage.TaskCompositionShapes, stage.TaskCompositionShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskInputShape:
		tmp := GetStructInstancesByOrder(stage.TaskInputShapes, stage.TaskInputShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskInputShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskShape:
		tmp := GetStructInstancesByOrder(stage.TaskShapes, stage.TaskShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskShape implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.DiagramMap_Staged_Order)
	case "Product":
		res = GetNamedStructInstances(stage.Products, stage.ProductMap_Staged_Order)
	case "ProductCompositionShape":
		res = GetNamedStructInstances(stage.ProductCompositionShapes, stage.ProductCompositionShapeMap_Staged_Order)
	case "ProductShape":
		res = GetNamedStructInstances(stage.ProductShapes, stage.ProductShapeMap_Staged_Order)
	case "Project":
		res = GetNamedStructInstances(stage.Projects, stage.ProjectMap_Staged_Order)
	case "Root":
		res = GetNamedStructInstances(stage.Roots, stage.RootMap_Staged_Order)
	case "Task":
		res = GetNamedStructInstances(stage.Tasks, stage.TaskMap_Staged_Order)
	case "TaskCompositionShape":
		res = GetNamedStructInstances(stage.TaskCompositionShapes, stage.TaskCompositionShapeMap_Staged_Order)
	case "TaskInputShape":
		res = GetNamedStructInstances(stage.TaskInputShapes, stage.TaskInputShapeMap_Staged_Order)
	case "TaskShape":
		res = GetNamedStructInstances(stage.TaskShapes, stage.TaskShapeMap_Staged_Order)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsme/project/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return project_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return project_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitProduct(product *Product)
	CheckoutProduct(product *Product)
	CommitProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CheckoutProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CommitProductShape(productshape *ProductShape)
	CheckoutProductShape(productshape *ProductShape)
	CommitProject(project *Project)
	CheckoutProject(project *Project)
	CommitRoot(root *Root)
	CheckoutRoot(root *Root)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	CommitTaskCompositionShape(taskcompositionshape *TaskCompositionShape)
	CheckoutTaskCompositionShape(taskcompositionshape *TaskCompositionShape)
	CommitTaskInputShape(taskinputshape *TaskInputShape)
	CheckoutTaskInputShape(taskinputshape *TaskInputShape)
	CommitTaskShape(taskshape *TaskShape)
	CheckoutTaskShape(taskshape *TaskShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Products:           make(map[*Product]struct{}),
		Products_mapString: make(map[string]*Product),

		ProductCompositionShapes:           make(map[*ProductCompositionShape]struct{}),
		ProductCompositionShapes_mapString: make(map[string]*ProductCompositionShape),

		ProductShapes:           make(map[*ProductShape]struct{}),
		ProductShapes_mapString: make(map[string]*ProductShape),

		Projects:           make(map[*Project]struct{}),
		Projects_mapString: make(map[string]*Project),

		Roots:           make(map[*Root]struct{}),
		Roots_mapString: make(map[string]*Root),

		Tasks:           make(map[*Task]struct{}),
		Tasks_mapString: make(map[string]*Task),

		TaskCompositionShapes:           make(map[*TaskCompositionShape]struct{}),
		TaskCompositionShapes_mapString: make(map[string]*TaskCompositionShape),

		TaskInputShapes:           make(map[*TaskInputShape]struct{}),
		TaskInputShapes_mapString: make(map[string]*TaskInputShape),

		TaskShapes:           make(map[*TaskShape]struct{}),
		TaskShapes_mapString: make(map[string]*TaskShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DiagramMap_Staged_Order: make(map[*Diagram]uint),

		ProductMap_Staged_Order: make(map[*Product]uint),

		ProductCompositionShapeMap_Staged_Order: make(map[*ProductCompositionShape]uint),

		ProductShapeMap_Staged_Order: make(map[*ProductShape]uint),

		ProjectMap_Staged_Order: make(map[*Project]uint),

		RootMap_Staged_Order: make(map[*Root]uint),

		TaskMap_Staged_Order: make(map[*Task]uint),

		TaskCompositionShapeMap_Staged_Order: make(map[*TaskCompositionShape]uint),

		TaskInputShapeMap_Staged_Order: make(map[*TaskInputShape]uint),

		TaskShapeMap_Staged_Order: make(map[*TaskShape]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Diagram"},
			{name: "Product"},
			{name: "ProductCompositionShape"},
			{name: "ProductShape"},
			{name: "Project"},
			{name: "Root"},
			{name: "Task"},
			{name: "TaskCompositionShape"},
			{name: "TaskInputShape"},
			{name: "TaskShape"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Product:
		return stage.ProductMap_Staged_Order[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShapeMap_Staged_Order[instance]
	case *ProductShape:
		return stage.ProductShapeMap_Staged_Order[instance]
	case *Project:
		return stage.ProjectMap_Staged_Order[instance]
	case *Root:
		return stage.RootMap_Staged_Order[instance]
	case *Task:
		return stage.TaskMap_Staged_Order[instance]
	case *TaskCompositionShape:
		return stage.TaskCompositionShapeMap_Staged_Order[instance]
	case *TaskInputShape:
		return stage.TaskInputShapeMap_Staged_Order[instance]
	case *TaskShape:
		return stage.TaskShapeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Product:
		return stage.ProductMap_Staged_Order[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShapeMap_Staged_Order[instance]
	case *ProductShape:
		return stage.ProductShapeMap_Staged_Order[instance]
	case *Project:
		return stage.ProjectMap_Staged_Order[instance]
	case *Root:
		return stage.RootMap_Staged_Order[instance]
	case *Task:
		return stage.TaskMap_Staged_Order[instance]
	case *TaskCompositionShape:
		return stage.TaskCompositionShapeMap_Staged_Order[instance]
	case *TaskInputShape:
		return stage.TaskInputShapeMap_Staged_Order[instance]
	case *TaskShape:
		return stage.TaskShapeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Product"] = len(stage.Products)
	stage.Map_GongStructName_InstancesNb["ProductCompositionShape"] = len(stage.ProductCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ProductShape"] = len(stage.ProductShapes)
	stage.Map_GongStructName_InstancesNb["Project"] = len(stage.Projects)
	stage.Map_GongStructName_InstancesNb["Root"] = len(stage.Roots)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)
	stage.Map_GongStructName_InstancesNb["TaskCompositionShape"] = len(stage.TaskCompositionShapes)
	stage.Map_GongStructName_InstancesNb["TaskInputShape"] = len(stage.TaskInputShapes)
	stage.Map_GongStructName_InstancesNb["TaskShape"] = len(stage.TaskShapes)
}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {

	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}
		stage.DiagramMap_Staged_Order[diagram] = stage.DiagramOrder
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram

	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}

		if order > stage.DiagramOrder {
			stage.DiagramOrder = order
		}
		stage.DiagramMap_Staged_Order[diagram] = stage.DiagramOrder
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	delete(stage.Diagrams, diagram)
	delete(stage.DiagramMap_Staged_Order, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)

	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	delete(stage.Diagrams, diagram)
	delete(stage.DiagramMap_Staged_Order, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)
}

// commit diagram to the back repo (if it is already staged)
func (diagram *Diagram) Commit(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagram(diagram)
		}
	}
	return diagram
}

func (diagram *Diagram) CommitVoid(stage *Stage) {
	diagram.Commit(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// Checkout diagram to the back repo (if it is already staged)
func (diagram *Diagram) Checkout(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagram(diagram)
		}
	}
	return diagram
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) GetName() (res string) {
	return diagram.Name
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) SetName(name string) (){
	diagram.Name = name
}

// Stage puts product to the model stage
func (product *Product) Stage(stage *Stage) *Product {

	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}
		stage.ProductMap_Staged_Order[product] = stage.ProductOrder
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product

	return product
}

// StagePreserveOrder puts product to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductOrder
// - update stage.ProductOrder accordingly
func (product *Product) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}

		if order > stage.ProductOrder {
			stage.ProductOrder = order
		}
		stage.ProductMap_Staged_Order[product] = stage.ProductOrder
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product
}

// Unstage removes product off the model stage
func (product *Product) Unstage(stage *Stage) *Product {
	delete(stage.Products, product)
	delete(stage.ProductMap_Staged_Order, product)
	delete(stage.Products_mapString, product.Name)

	return product
}

// UnstageVoid removes product off the model stage
func (product *Product) UnstageVoid(stage *Stage) {
	delete(stage.Products, product)
	delete(stage.ProductMap_Staged_Order, product)
	delete(stage.Products_mapString, product.Name)
}

// commit product to the back repo (if it is already staged)
func (product *Product) Commit(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProduct(product)
		}
	}
	return product
}

func (product *Product) CommitVoid(stage *Stage) {
	product.Commit(stage)
}

func (product *Product) StageVoid(stage *Stage) {
	product.Stage(stage)
}

// Checkout product to the back repo (if it is already staged)
func (product *Product) Checkout(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProduct(product)
		}
	}
	return product
}

// for satisfaction of GongStruct interface
func (product *Product) GetName() (res string) {
	return product.Name
}

// for satisfaction of GongStruct interface
func (product *Product) SetName(name string) (){
	product.Name = name
}

// Stage puts productcompositionshape to the model stage
func (productcompositionshape *ProductCompositionShape) Stage(stage *Stage) *ProductCompositionShape {

	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; !ok {
		stage.ProductCompositionShapes[productcompositionshape] = struct{}{}
		stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape] = stage.ProductCompositionShapeOrder
		stage.ProductCompositionShapeOrder++
	}
	stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape

	return productcompositionshape
}

// StagePreserveOrder puts productcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductCompositionShapeOrder
// - update stage.ProductCompositionShapeOrder accordingly
func (productcompositionshape *ProductCompositionShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; !ok {
		stage.ProductCompositionShapes[productcompositionshape] = struct{}{}

		if order > stage.ProductCompositionShapeOrder {
			stage.ProductCompositionShapeOrder = order
		}
		stage.ProductCompositionShapeMap_Staged_Order[productcompositionshape] = stage.ProductCompositionShapeOrder
		stage.ProductCompositionShapeOrder++
	}
	stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape
}

// Unstage removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) Unstage(stage *Stage) *ProductCompositionShape {
	delete(stage.ProductCompositionShapes, productcompositionshape)
	delete(stage.ProductCompositionShapeMap_Staged_Order, productcompositionshape)
	delete(stage.ProductCompositionShapes_mapString, productcompositionshape.Name)

	return productcompositionshape
}

// UnstageVoid removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.ProductCompositionShapes, productcompositionshape)
	delete(stage.ProductCompositionShapeMap_Staged_Order, productcompositionshape)
	delete(stage.ProductCompositionShapes_mapString, productcompositionshape.Name)
}

// commit productcompositionshape to the back repo (if it is already staged)
func (productcompositionshape *ProductCompositionShape) Commit(stage *Stage) *ProductCompositionShape {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProductCompositionShape(productcompositionshape)
		}
	}
	return productcompositionshape
}

func (productcompositionshape *ProductCompositionShape) CommitVoid(stage *Stage) {
	productcompositionshape.Commit(stage)
}

func (productcompositionshape *ProductCompositionShape) StageVoid(stage *Stage) {
	productcompositionshape.Stage(stage)
}

// Checkout productcompositionshape to the back repo (if it is already staged)
func (productcompositionshape *ProductCompositionShape) Checkout(stage *Stage) *ProductCompositionShape {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProductCompositionShape(productcompositionshape)
		}
	}
	return productcompositionshape
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) GetName() (res string) {
	return productcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) SetName(name string) (){
	productcompositionshape.Name = name
}

// Stage puts productshape to the model stage
func (productshape *ProductShape) Stage(stage *Stage) *ProductShape {

	if _, ok := stage.ProductShapes[productshape]; !ok {
		stage.ProductShapes[productshape] = struct{}{}
		stage.ProductShapeMap_Staged_Order[productshape] = stage.ProductShapeOrder
		stage.ProductShapeOrder++
	}
	stage.ProductShapes_mapString[productshape.Name] = productshape

	return productshape
}

// StagePreserveOrder puts productshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductShapeOrder
// - update stage.ProductShapeOrder accordingly
func (productshape *ProductShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ProductShapes[productshape]; !ok {
		stage.ProductShapes[productshape] = struct{}{}

		if order > stage.ProductShapeOrder {
			stage.ProductShapeOrder = order
		}
		stage.ProductShapeMap_Staged_Order[productshape] = stage.ProductShapeOrder
		stage.ProductShapeOrder++
	}
	stage.ProductShapes_mapString[productshape.Name] = productshape
}

// Unstage removes productshape off the model stage
func (productshape *ProductShape) Unstage(stage *Stage) *ProductShape {
	delete(stage.ProductShapes, productshape)
	delete(stage.ProductShapeMap_Staged_Order, productshape)
	delete(stage.ProductShapes_mapString, productshape.Name)

	return productshape
}

// UnstageVoid removes productshape off the model stage
func (productshape *ProductShape) UnstageVoid(stage *Stage) {
	delete(stage.ProductShapes, productshape)
	delete(stage.ProductShapeMap_Staged_Order, productshape)
	delete(stage.ProductShapes_mapString, productshape.Name)
}

// commit productshape to the back repo (if it is already staged)
func (productshape *ProductShape) Commit(stage *Stage) *ProductShape {
	if _, ok := stage.ProductShapes[productshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProductShape(productshape)
		}
	}
	return productshape
}

func (productshape *ProductShape) CommitVoid(stage *Stage) {
	productshape.Commit(stage)
}

func (productshape *ProductShape) StageVoid(stage *Stage) {
	productshape.Stage(stage)
}

// Checkout productshape to the back repo (if it is already staged)
func (productshape *ProductShape) Checkout(stage *Stage) *ProductShape {
	if _, ok := stage.ProductShapes[productshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProductShape(productshape)
		}
	}
	return productshape
}

// for satisfaction of GongStruct interface
func (productshape *ProductShape) GetName() (res string) {
	return productshape.Name
}

// for satisfaction of GongStruct interface
func (productshape *ProductShape) SetName(name string) (){
	productshape.Name = name
}

// Stage puts project to the model stage
func (project *Project) Stage(stage *Stage) *Project {

	if _, ok := stage.Projects[project]; !ok {
		stage.Projects[project] = struct{}{}
		stage.ProjectMap_Staged_Order[project] = stage.ProjectOrder
		stage.ProjectOrder++
	}
	stage.Projects_mapString[project.Name] = project

	return project
}

// StagePreserveOrder puts project to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProjectOrder
// - update stage.ProjectOrder accordingly
func (project *Project) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Projects[project]; !ok {
		stage.Projects[project] = struct{}{}

		if order > stage.ProjectOrder {
			stage.ProjectOrder = order
		}
		stage.ProjectMap_Staged_Order[project] = stage.ProjectOrder
		stage.ProjectOrder++
	}
	stage.Projects_mapString[project.Name] = project
}

// Unstage removes project off the model stage
func (project *Project) Unstage(stage *Stage) *Project {
	delete(stage.Projects, project)
	delete(stage.ProjectMap_Staged_Order, project)
	delete(stage.Projects_mapString, project.Name)

	return project
}

// UnstageVoid removes project off the model stage
func (project *Project) UnstageVoid(stage *Stage) {
	delete(stage.Projects, project)
	delete(stage.ProjectMap_Staged_Order, project)
	delete(stage.Projects_mapString, project.Name)
}

// commit project to the back repo (if it is already staged)
func (project *Project) Commit(stage *Stage) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProject(project)
		}
	}
	return project
}

func (project *Project) CommitVoid(stage *Stage) {
	project.Commit(stage)
}

func (project *Project) StageVoid(stage *Stage) {
	project.Stage(stage)
}

// Checkout project to the back repo (if it is already staged)
func (project *Project) Checkout(stage *Stage) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProject(project)
		}
	}
	return project
}

// for satisfaction of GongStruct interface
func (project *Project) GetName() (res string) {
	return project.Name
}

// for satisfaction of GongStruct interface
func (project *Project) SetName(name string) (){
	project.Name = name
}

// Stage puts root to the model stage
func (root *Root) Stage(stage *Stage) *Root {

	if _, ok := stage.Roots[root]; !ok {
		stage.Roots[root] = struct{}{}
		stage.RootMap_Staged_Order[root] = stage.RootOrder
		stage.RootOrder++
	}
	stage.Roots_mapString[root.Name] = root

	return root
}

// StagePreserveOrder puts root to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RootOrder
// - update stage.RootOrder accordingly
func (root *Root) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Roots[root]; !ok {
		stage.Roots[root] = struct{}{}

		if order > stage.RootOrder {
			stage.RootOrder = order
		}
		stage.RootMap_Staged_Order[root] = stage.RootOrder
		stage.RootOrder++
	}
	stage.Roots_mapString[root.Name] = root
}

// Unstage removes root off the model stage
func (root *Root) Unstage(stage *Stage) *Root {
	delete(stage.Roots, root)
	delete(stage.RootMap_Staged_Order, root)
	delete(stage.Roots_mapString, root.Name)

	return root
}

// UnstageVoid removes root off the model stage
func (root *Root) UnstageVoid(stage *Stage) {
	delete(stage.Roots, root)
	delete(stage.RootMap_Staged_Order, root)
	delete(stage.Roots_mapString, root.Name)
}

// commit root to the back repo (if it is already staged)
func (root *Root) Commit(stage *Stage) *Root {
	if _, ok := stage.Roots[root]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRoot(root)
		}
	}
	return root
}

func (root *Root) CommitVoid(stage *Stage) {
	root.Commit(stage)
}

func (root *Root) StageVoid(stage *Stage) {
	root.Stage(stage)
}

// Checkout root to the back repo (if it is already staged)
func (root *Root) Checkout(stage *Stage) *Root {
	if _, ok := stage.Roots[root]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRoot(root)
		}
	}
	return root
}

// for satisfaction of GongStruct interface
func (root *Root) GetName() (res string) {
	return root.Name
}

// for satisfaction of GongStruct interface
func (root *Root) SetName(name string) (){
	root.Name = name
}

// Stage puts task to the model stage
func (task *Task) Stage(stage *Stage) *Task {

	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}
		stage.TaskMap_Staged_Order[task] = stage.TaskOrder
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task

	return task
}

// StagePreserveOrder puts task to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOrder
// - update stage.TaskOrder accordingly
func (task *Task) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}

		if order > stage.TaskOrder {
			stage.TaskOrder = order
		}
		stage.TaskMap_Staged_Order[task] = stage.TaskOrder
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task
}

// Unstage removes task off the model stage
func (task *Task) Unstage(stage *Stage) *Task {
	delete(stage.Tasks, task)
	delete(stage.TaskMap_Staged_Order, task)
	delete(stage.Tasks_mapString, task.Name)

	return task
}

// UnstageVoid removes task off the model stage
func (task *Task) UnstageVoid(stage *Stage) {
	delete(stage.Tasks, task)
	delete(stage.TaskMap_Staged_Order, task)
	delete(stage.Tasks_mapString, task.Name)
}

// commit task to the back repo (if it is already staged)
func (task *Task) Commit(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTask(task)
		}
	}
	return task
}

func (task *Task) CommitVoid(stage *Stage) {
	task.Commit(stage)
}

func (task *Task) StageVoid(stage *Stage) {
	task.Stage(stage)
}

// Checkout task to the back repo (if it is already staged)
func (task *Task) Checkout(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTask(task)
		}
	}
	return task
}

// for satisfaction of GongStruct interface
func (task *Task) GetName() (res string) {
	return task.Name
}

// for satisfaction of GongStruct interface
func (task *Task) SetName(name string) (){
	task.Name = name
}

// Stage puts taskcompositionshape to the model stage
func (taskcompositionshape *TaskCompositionShape) Stage(stage *Stage) *TaskCompositionShape {

	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; !ok {
		stage.TaskCompositionShapes[taskcompositionshape] = struct{}{}
		stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape] = stage.TaskCompositionShapeOrder
		stage.TaskCompositionShapeOrder++
	}
	stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape

	return taskcompositionshape
}

// StagePreserveOrder puts taskcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskCompositionShapeOrder
// - update stage.TaskCompositionShapeOrder accordingly
func (taskcompositionshape *TaskCompositionShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; !ok {
		stage.TaskCompositionShapes[taskcompositionshape] = struct{}{}

		if order > stage.TaskCompositionShapeOrder {
			stage.TaskCompositionShapeOrder = order
		}
		stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshape] = stage.TaskCompositionShapeOrder
		stage.TaskCompositionShapeOrder++
	}
	stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape
}

// Unstage removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) Unstage(stage *Stage) *TaskCompositionShape {
	delete(stage.TaskCompositionShapes, taskcompositionshape)
	delete(stage.TaskCompositionShapeMap_Staged_Order, taskcompositionshape)
	delete(stage.TaskCompositionShapes_mapString, taskcompositionshape.Name)

	return taskcompositionshape
}

// UnstageVoid removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskCompositionShapes, taskcompositionshape)
	delete(stage.TaskCompositionShapeMap_Staged_Order, taskcompositionshape)
	delete(stage.TaskCompositionShapes_mapString, taskcompositionshape.Name)
}

// commit taskcompositionshape to the back repo (if it is already staged)
func (taskcompositionshape *TaskCompositionShape) Commit(stage *Stage) *TaskCompositionShape {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskCompositionShape(taskcompositionshape)
		}
	}
	return taskcompositionshape
}

func (taskcompositionshape *TaskCompositionShape) CommitVoid(stage *Stage) {
	taskcompositionshape.Commit(stage)
}

func (taskcompositionshape *TaskCompositionShape) StageVoid(stage *Stage) {
	taskcompositionshape.Stage(stage)
}

// Checkout taskcompositionshape to the back repo (if it is already staged)
func (taskcompositionshape *TaskCompositionShape) Checkout(stage *Stage) *TaskCompositionShape {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskCompositionShape(taskcompositionshape)
		}
	}
	return taskcompositionshape
}

// for satisfaction of GongStruct interface
func (taskcompositionshape *TaskCompositionShape) GetName() (res string) {
	return taskcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (taskcompositionshape *TaskCompositionShape) SetName(name string) (){
	taskcompositionshape.Name = name
}

// Stage puts taskinputshape to the model stage
func (taskinputshape *TaskInputShape) Stage(stage *Stage) *TaskInputShape {

	if _, ok := stage.TaskInputShapes[taskinputshape]; !ok {
		stage.TaskInputShapes[taskinputshape] = struct{}{}
		stage.TaskInputShapeMap_Staged_Order[taskinputshape] = stage.TaskInputShapeOrder
		stage.TaskInputShapeOrder++
	}
	stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape

	return taskinputshape
}

// StagePreserveOrder puts taskinputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskInputShapeOrder
// - update stage.TaskInputShapeOrder accordingly
func (taskinputshape *TaskInputShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TaskInputShapes[taskinputshape]; !ok {
		stage.TaskInputShapes[taskinputshape] = struct{}{}

		if order > stage.TaskInputShapeOrder {
			stage.TaskInputShapeOrder = order
		}
		stage.TaskInputShapeMap_Staged_Order[taskinputshape] = stage.TaskInputShapeOrder
		stage.TaskInputShapeOrder++
	}
	stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape
}

// Unstage removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) Unstage(stage *Stage) *TaskInputShape {
	delete(stage.TaskInputShapes, taskinputshape)
	delete(stage.TaskInputShapeMap_Staged_Order, taskinputshape)
	delete(stage.TaskInputShapes_mapString, taskinputshape.Name)

	return taskinputshape
}

// UnstageVoid removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskInputShapes, taskinputshape)
	delete(stage.TaskInputShapeMap_Staged_Order, taskinputshape)
	delete(stage.TaskInputShapes_mapString, taskinputshape.Name)
}

// commit taskinputshape to the back repo (if it is already staged)
func (taskinputshape *TaskInputShape) Commit(stage *Stage) *TaskInputShape {
	if _, ok := stage.TaskInputShapes[taskinputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskInputShape(taskinputshape)
		}
	}
	return taskinputshape
}

func (taskinputshape *TaskInputShape) CommitVoid(stage *Stage) {
	taskinputshape.Commit(stage)
}

func (taskinputshape *TaskInputShape) StageVoid(stage *Stage) {
	taskinputshape.Stage(stage)
}

// Checkout taskinputshape to the back repo (if it is already staged)
func (taskinputshape *TaskInputShape) Checkout(stage *Stage) *TaskInputShape {
	if _, ok := stage.TaskInputShapes[taskinputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskInputShape(taskinputshape)
		}
	}
	return taskinputshape
}

// for satisfaction of GongStruct interface
func (taskinputshape *TaskInputShape) GetName() (res string) {
	return taskinputshape.Name
}

// for satisfaction of GongStruct interface
func (taskinputshape *TaskInputShape) SetName(name string) (){
	taskinputshape.Name = name
}

// Stage puts taskshape to the model stage
func (taskshape *TaskShape) Stage(stage *Stage) *TaskShape {

	if _, ok := stage.TaskShapes[taskshape]; !ok {
		stage.TaskShapes[taskshape] = struct{}{}
		stage.TaskShapeMap_Staged_Order[taskshape] = stage.TaskShapeOrder
		stage.TaskShapeOrder++
	}
	stage.TaskShapes_mapString[taskshape.Name] = taskshape

	return taskshape
}

// StagePreserveOrder puts taskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskShapeOrder
// - update stage.TaskShapeOrder accordingly
func (taskshape *TaskShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TaskShapes[taskshape]; !ok {
		stage.TaskShapes[taskshape] = struct{}{}

		if order > stage.TaskShapeOrder {
			stage.TaskShapeOrder = order
		}
		stage.TaskShapeMap_Staged_Order[taskshape] = stage.TaskShapeOrder
		stage.TaskShapeOrder++
	}
	stage.TaskShapes_mapString[taskshape.Name] = taskshape
}

// Unstage removes taskshape off the model stage
func (taskshape *TaskShape) Unstage(stage *Stage) *TaskShape {
	delete(stage.TaskShapes, taskshape)
	delete(stage.TaskShapeMap_Staged_Order, taskshape)
	delete(stage.TaskShapes_mapString, taskshape.Name)

	return taskshape
}

// UnstageVoid removes taskshape off the model stage
func (taskshape *TaskShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskShapes, taskshape)
	delete(stage.TaskShapeMap_Staged_Order, taskshape)
	delete(stage.TaskShapes_mapString, taskshape.Name)
}

// commit taskshape to the back repo (if it is already staged)
func (taskshape *TaskShape) Commit(stage *Stage) *TaskShape {
	if _, ok := stage.TaskShapes[taskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskShape(taskshape)
		}
	}
	return taskshape
}

func (taskshape *TaskShape) CommitVoid(stage *Stage) {
	taskshape.Commit(stage)
}

func (taskshape *TaskShape) StageVoid(stage *Stage) {
	taskshape.Stage(stage)
}

// Checkout taskshape to the back repo (if it is already staged)
func (taskshape *TaskShape) Checkout(stage *Stage) *TaskShape {
	if _, ok := stage.TaskShapes[taskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskShape(taskshape)
		}
	}
	return taskshape
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) GetName() (res string) {
	return taskshape.Name
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) SetName(name string) (){
	taskshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDiagram(Diagram *Diagram)
	CreateORMProduct(Product *Product)
	CreateORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	CreateORMProductShape(ProductShape *ProductShape)
	CreateORMProject(Project *Project)
	CreateORMRoot(Root *Root)
	CreateORMTask(Task *Task)
	CreateORMTaskCompositionShape(TaskCompositionShape *TaskCompositionShape)
	CreateORMTaskInputShape(TaskInputShape *TaskInputShape)
	CreateORMTaskShape(TaskShape *TaskShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMProduct(Product *Product)
	DeleteORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	DeleteORMProductShape(ProductShape *ProductShape)
	DeleteORMProject(Project *Project)
	DeleteORMRoot(Root *Root)
	DeleteORMTask(Task *Task)
	DeleteORMTaskCompositionShape(TaskCompositionShape *TaskCompositionShape)
	DeleteORMTaskInputShape(TaskInputShape *TaskInputShape)
	DeleteORMTaskShape(TaskShape *TaskShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.DiagramMap_Staged_Order = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Products = make(map[*Product]struct{})
	stage.Products_mapString = make(map[string]*Product)
	stage.ProductMap_Staged_Order = make(map[*Product]uint)
	stage.ProductOrder = 0

	stage.ProductCompositionShapes = make(map[*ProductCompositionShape]struct{})
	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	stage.ProductCompositionShapeMap_Staged_Order = make(map[*ProductCompositionShape]uint)
	stage.ProductCompositionShapeOrder = 0

	stage.ProductShapes = make(map[*ProductShape]struct{})
	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	stage.ProductShapeMap_Staged_Order = make(map[*ProductShape]uint)
	stage.ProductShapeOrder = 0

	stage.Projects = make(map[*Project]struct{})
	stage.Projects_mapString = make(map[string]*Project)
	stage.ProjectMap_Staged_Order = make(map[*Project]uint)
	stage.ProjectOrder = 0

	stage.Roots = make(map[*Root]struct{})
	stage.Roots_mapString = make(map[string]*Root)
	stage.RootMap_Staged_Order = make(map[*Root]uint)
	stage.RootOrder = 0

	stage.Tasks = make(map[*Task]struct{})
	stage.Tasks_mapString = make(map[string]*Task)
	stage.TaskMap_Staged_Order = make(map[*Task]uint)
	stage.TaskOrder = 0

	stage.TaskCompositionShapes = make(map[*TaskCompositionShape]struct{})
	stage.TaskCompositionShapes_mapString = make(map[string]*TaskCompositionShape)
	stage.TaskCompositionShapeMap_Staged_Order = make(map[*TaskCompositionShape]uint)
	stage.TaskCompositionShapeOrder = 0

	stage.TaskInputShapes = make(map[*TaskInputShape]struct{})
	stage.TaskInputShapes_mapString = make(map[string]*TaskInputShape)
	stage.TaskInputShapeMap_Staged_Order = make(map[*TaskInputShape]uint)
	stage.TaskInputShapeOrder = 0

	stage.TaskShapes = make(map[*TaskShape]struct{})
	stage.TaskShapes_mapString = make(map[string]*TaskShape)
	stage.TaskShapeMap_Staged_Order = make(map[*TaskShape]uint)
	stage.TaskShapeOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Products = nil
	stage.Products_mapString = nil

	stage.ProductCompositionShapes = nil
	stage.ProductCompositionShapes_mapString = nil

	stage.ProductShapes = nil
	stage.ProductShapes_mapString = nil

	stage.Projects = nil
	stage.Projects_mapString = nil

	stage.Roots = nil
	stage.Roots_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

	stage.TaskCompositionShapes = nil
	stage.TaskCompositionShapes_mapString = nil

	stage.TaskInputShapes = nil
	stage.TaskInputShapes_mapString = nil

	stage.TaskShapes = nil
	stage.TaskShapes_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for diagram := range stage.Diagrams {
		diagram.Unstage(stage)
	}

	for product := range stage.Products {
		product.Unstage(stage)
	}

	for productcompositionshape := range stage.ProductCompositionShapes {
		productcompositionshape.Unstage(stage)
	}

	for productshape := range stage.ProductShapes {
		productshape.Unstage(stage)
	}

	for project := range stage.Projects {
		project.Unstage(stage)
	}

	for root := range stage.Roots {
		root.Unstage(stage)
	}

	for task := range stage.Tasks {
		task.Unstage(stage)
	}

	for taskcompositionshape := range stage.TaskCompositionShapes {
		taskcompositionshape.Unstage(stage)
	}

	for taskinputshape := range stage.TaskInputShapes {
		taskinputshape.Unstage(stage)
	}

	for taskshape := range stage.TaskShapes {
		taskshape.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Product]any:
		return any(&stage.Products).(*Type)
	case map[*ProductCompositionShape]any:
		return any(&stage.ProductCompositionShapes).(*Type)
	case map[*ProductShape]any:
		return any(&stage.ProductShapes).(*Type)
	case map[*Project]any:
		return any(&stage.Projects).(*Type)
	case map[*Root]any:
		return any(&stage.Roots).(*Type)
	case map[*Task]any:
		return any(&stage.Tasks).(*Type)
	case map[*TaskCompositionShape]any:
		return any(&stage.TaskCompositionShapes).(*Type)
	case map[*TaskInputShape]any:
		return any(&stage.TaskInputShapes).(*Type)
	case map[*TaskShape]any:
		return any(&stage.TaskShapes).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Product:
		return any(stage.Products_mapString).(map[string]Type)
	case *ProductCompositionShape:
		return any(stage.ProductCompositionShapes_mapString).(map[string]Type)
	case *ProductShape:
		return any(stage.ProductShapes_mapString).(map[string]Type)
	case *Project:
		return any(stage.Projects_mapString).(map[string]Type)
	case *Root:
		return any(stage.Roots_mapString).(map[string]Type)
	case *Task:
		return any(stage.Tasks_mapString).(map[string]Type)
	case *TaskCompositionShape:
		return any(stage.TaskCompositionShapes_mapString).(map[string]Type)
	case *TaskInputShape:
		return any(stage.TaskInputShapes_mapString).(map[string]Type)
	case *TaskShape:
		return any(stage.TaskShapes_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Product:
		return any(&stage.Products).(*map[*Type]struct{})
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[*Type]struct{})
	case ProductShape:
		return any(&stage.ProductShapes).(*map[*Type]struct{})
	case Project:
		return any(&stage.Projects).(*map[*Type]struct{})
	case Root:
		return any(&stage.Roots).(*map[*Type]struct{})
	case Task:
		return any(&stage.Tasks).(*map[*Type]struct{})
	case TaskCompositionShape:
		return any(&stage.TaskCompositionShapes).(*map[*Type]struct{})
	case TaskInputShape:
		return any(&stage.TaskInputShapes).(*map[*Type]struct{})
	case TaskShape:
		return any(&stage.TaskShapes).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Product:
		return any(&stage.Products).(*map[Type]struct{})
	case *ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[Type]struct{})
	case *ProductShape:
		return any(&stage.ProductShapes).(*map[Type]struct{})
	case *Project:
		return any(&stage.Projects).(*map[Type]struct{})
	case *Root:
		return any(&stage.Roots).(*map[Type]struct{})
	case *Task:
		return any(&stage.Tasks).(*map[Type]struct{})
	case *TaskCompositionShape:
		return any(&stage.TaskCompositionShapes).(*map[Type]struct{})
	case *TaskInputShape:
		return any(&stage.TaskInputShapes).(*map[Type]struct{})
	case *TaskShape:
		return any(&stage.TaskShapes).(*map[Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Product:
		return any(&stage.Products_mapString).(*map[string]*Type)
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes_mapString).(*map[string]*Type)
	case ProductShape:
		return any(&stage.ProductShapes_mapString).(*map[string]*Type)
	case Project:
		return any(&stage.Projects_mapString).(*map[string]*Type)
	case Root:
		return any(&stage.Roots_mapString).(*map[string]*Type)
	case Task:
		return any(&stage.Tasks_mapString).(*map[string]*Type)
	case TaskCompositionShape:
		return any(&stage.TaskCompositionShapes_mapString).(*map[string]*Type)
	case TaskInputShape:
		return any(&stage.TaskInputShapes_mapString).(*map[string]*Type)
	case TaskShape:
		return any(&stage.TaskShapes_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of ProductShape with the name of the field
			Product_Shapes: []*ProductShape{{Name: "Product_Shapes"}},
			// field is initialized with an instance of Product with the name of the field
			ProductsWhoseNodeIsExpanded: []*Product{{Name: "ProductsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ProductCompositionShape with the name of the field
			ProductComposition_Shapes: []*ProductCompositionShape{{Name: "ProductComposition_Shapes"}},
			// field is initialized with an instance of TaskShape with the name of the field
			Task_Shapes: []*TaskShape{{Name: "Task_Shapes"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseInputNodeIsExpanded: []*Task{{Name: "TasksWhoseInputNodeIsExpanded"}},
			// field is initialized with an instance of TaskCompositionShape with the name of the field
			TaskComposition_Shapes: []*TaskCompositionShape{{Name: "TaskComposition_Shapes"}},
			// field is initialized with an instance of TaskInputShape with the name of the field
			TaskInputShapes: []*TaskInputShape{{Name: "TaskInputShapes"}},
		}).(*Type)
	case Product:
		return any(&Product{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			SubProducts: []*Product{{Name: "SubProducts"}},
		}).(*Type)
	case ProductCompositionShape:
		return any(&ProductCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case ProductShape:
		return any(&ProductShape{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case Project:
		return any(&Project{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			RootProducts: []*Product{{Name: "RootProducts"}},
			// field is initialized with an instance of Task with the name of the field
			RootTasks: []*Task{{Name: "RootTasks"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case Root:
		return any(&Root{
			// Initialisation of associations
			// field is initialized with an instance of Project with the name of the field
			Projects: []*Project{{Name: "Projects"}},
			// field is initialized with an instance of Product with the name of the field
			OrphanedProducts: []*Product{{Name: "OrphanedProducts"}},
			// field is initialized with an instance of Task with the name of the field
			OrphanedTasks: []*Task{{Name: "OrphanedTasks"}},
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			SubTasks: []*Task{{Name: "SubTasks"}},
			// field is initialized with an instance of Product with the name of the field
			Inputs: []*Product{{Name: "Inputs"}},
			// field is initialized with an instance of Product with the name of the field
			Outputs: []*Product{{Name: "Outputs"}},
		}).(*Type)
	case TaskCompositionShape:
		return any(&TaskCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskInputShape:
		return any(&TaskInputShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case TaskShape:
		return any(&TaskShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProductCompositionShape
	case ProductCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*ProductCompositionShape)
			for productcompositionshape := range stage.ProductCompositionShapes {
				if productcompositionshape.Product != nil {
					product_ := productcompositionshape.Product
					var productcompositionshapes []*ProductCompositionShape
					_, ok := res[product_]
					if ok {
						productcompositionshapes = res[product_]
					} else {
						productcompositionshapes = make([]*ProductCompositionShape, 0)
					}
					productcompositionshapes = append(productcompositionshapes, productcompositionshape)
					res[product_] = productcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductShape
	case ProductShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*ProductShape)
			for productshape := range stage.ProductShapes {
				if productshape.Product != nil {
					product_ := productshape.Product
					var productshapes []*ProductShape
					_, ok := res[product_]
					if ok {
						productshapes = res[product_]
					} else {
						productshapes = make([]*ProductShape, 0)
					}
					productshapes = append(productshapes, productshape)
					res[product_] = productshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskCompositionShape
	case TaskCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskCompositionShape)
			for taskcompositionshape := range stage.TaskCompositionShapes {
				if taskcompositionshape.Task != nil {
					task_ := taskcompositionshape.Task
					var taskcompositionshapes []*TaskCompositionShape
					_, ok := res[task_]
					if ok {
						taskcompositionshapes = res[task_]
					} else {
						taskcompositionshapes = make([]*TaskCompositionShape, 0)
					}
					taskcompositionshapes = append(taskcompositionshapes, taskcompositionshape)
					res[task_] = taskcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskInputShape
	case TaskInputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskInputShape)
			for taskinputshape := range stage.TaskInputShapes {
				if taskinputshape.Task != nil {
					task_ := taskinputshape.Task
					var taskinputshapes []*TaskInputShape
					_, ok := res[task_]
					if ok {
						taskinputshapes = res[task_]
					} else {
						taskinputshapes = make([]*TaskInputShape, 0)
					}
					taskinputshapes = append(taskinputshapes, taskinputshape)
					res[task_] = taskinputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Product":
			res := make(map[*Product][]*TaskInputShape)
			for taskinputshape := range stage.TaskInputShapes {
				if taskinputshape.Product != nil {
					product_ := taskinputshape.Product
					var taskinputshapes []*TaskInputShape
					_, ok := res[product_]
					if ok {
						taskinputshapes = res[product_]
					} else {
						taskinputshapes = make([]*TaskInputShape, 0)
					}
					taskinputshapes = append(taskinputshapes, taskinputshape)
					res[product_] = taskinputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskShape
	case TaskShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskShape)
			for taskshape := range stage.TaskShapes {
				if taskshape.Task != nil {
					task_ := taskshape.Task
					var taskshapes []*TaskShape
					_, ok := res[task_]
					if ok {
						taskshapes = res[task_]
					} else {
						taskshapes = make([]*TaskShape, 0)
					}
					taskshapes = append(taskshapes, taskshape)
					res[task_] = taskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "Product_Shapes":
			res := make(map[*ProductShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productshape_ := range diagram.Product_Shapes {
					res[productshape_] = append(res[productshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProductsWhoseNodeIsExpanded":
			res := make(map[*Product][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, product_ := range diagram.ProductsWhoseNodeIsExpanded {
					res[product_] = append(res[product_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProductComposition_Shapes":
			res := make(map[*ProductCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productcompositionshape_ := range diagram.ProductComposition_Shapes {
					res[productcompositionshape_] = append(res[productcompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Task_Shapes":
			res := make(map[*TaskShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskshape_ := range diagram.Task_Shapes {
					res[taskshape_] = append(res[taskshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhoseNodeIsExpanded {
					res[task_] = append(res[task_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseInputNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhoseInputNodeIsExpanded {
					res[task_] = append(res[task_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskComposition_Shapes":
			res := make(map[*TaskCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskcompositionshape_ := range diagram.TaskComposition_Shapes {
					res[taskcompositionshape_] = append(res[taskcompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskInputShapes":
			res := make(map[*TaskInputShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskinputshape_ := range diagram.TaskInputShapes {
					res[taskinputshape_] = append(res[taskinputshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		case "SubProducts":
			res := make(map[*Product][]*Product)
			for product := range stage.Products {
				for _, product_ := range product.SubProducts {
					res[product_] = append(res[product_], product)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductCompositionShape
	case ProductCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProductShape
	case ProductShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		case "RootProducts":
			res := make(map[*Product][]*Project)
			for project := range stage.Projects {
				for _, product_ := range project.RootProducts {
					res[product_] = append(res[product_], project)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootTasks":
			res := make(map[*Task][]*Project)
			for project := range stage.Projects {
				for _, task_ := range project.RootTasks {
					res[task_] = append(res[task_], project)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagrams":
			res := make(map[*Diagram][]*Project)
			for project := range stage.Projects {
				for _, diagram_ := range project.Diagrams {
					res[diagram_] = append(res[diagram_], project)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		case "Projects":
			res := make(map[*Project][]*Root)
			for root := range stage.Roots {
				for _, project_ := range root.Projects {
					res[project_] = append(res[project_], root)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OrphanedProducts":
			res := make(map[*Product][]*Root)
			for root := range stage.Roots {
				for _, product_ := range root.OrphanedProducts {
					res[product_] = append(res[product_], root)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OrphanedTasks":
			res := make(map[*Task][]*Root)
			for root := range stage.Roots {
				for _, task_ := range root.OrphanedTasks {
					res[task_] = append(res[task_], root)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "SubTasks":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				for _, task_ := range task.SubTasks {
					res[task_] = append(res[task_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Inputs":
			res := make(map[*Product][]*Task)
			for task := range stage.Tasks {
				for _, product_ := range task.Inputs {
					res[product_] = append(res[product_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Outputs":
			res := make(map[*Product][]*Task)
			for task := range stage.Tasks {
				for _, product_ := range task.Outputs {
					res[product_] = append(res[product_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskCompositionShape
	case TaskCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskInputShape
	case TaskInputShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskShape
	case TaskShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Diagram:
		res = "Diagram"
	case *Product:
		res = "Product"
	case *ProductCompositionShape:
		res = "ProductCompositionShape"
	case *ProductShape:
		res = "ProductShape"
	case *Project:
		res = "Project"
	case *Root:
		res = "Root"
	case *Task:
		res = "Task"
	case *TaskCompositionShape:
		res = "TaskCompositionShape"
	case *TaskInputShape:
		res = "TaskInputShape"
	case *TaskShape:
		res = "TaskShape"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Diagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Project"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case *Product:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Product"
		rf.Fieldname = "SubProducts"
		res = append(res, rf)
		rf.GongstructName = "Project"
		rf.Fieldname = "RootProducts"
		res = append(res, rf)
		rf.GongstructName = "Root"
		rf.Fieldname = "OrphanedProducts"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "Inputs"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "Outputs"
		res = append(res, rf)
	case *ProductCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductComposition_Shapes"
		res = append(res, rf)
	case *ProductShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Product_Shapes"
		res = append(res, rf)
	case *Project:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Root"
		rf.Fieldname = "Projects"
		res = append(res, rf)
	case *Root:
		var rf ReverseField
		_ = rf
	case *Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhoseInputNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Project"
		rf.Fieldname = "RootTasks"
		res = append(res, rf)
		rf.GongstructName = "Root"
		rf.Fieldname = "OrphanedTasks"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "SubTasks"
		res = append(res, rf)
	case *TaskCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskComposition_Shapes"
		res = append(res, rf)
	case *TaskInputShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskInputShapes"
		res = append(res, rf)
	case *TaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Task_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (diagram *Diagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsInRenameMode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Product_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductShape",
		},
		{
			Name:                 "ProductsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsPBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "ProductComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductCompositionShape",
		},
		{
			Name:               "IsWBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Task_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskShape",
		},
		{
			Name:                 "TasksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TasksWhoseInputNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TaskComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskCompositionShape",
		},
		{
			Name:                 "TaskInputShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskInputShape",
		},
	}
	return
}

func (product *Product) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "SubProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsProducersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsConsumersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (productshape *ProductShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (project *Project) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "RootProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsPBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "RootTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsWBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:               "IsDiagramsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (root *Root) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Projects",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Project",
		},
		{
			Name:                 "OrphanedProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:                 "OrphanedTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (task *Task) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "SubTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Inputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsInputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Outputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsOutputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (taskinputshape *TaskInputShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (taskshape *TaskShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {

	var ret Type
	return ret.GongGetFieldHeaders()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

// insertion point for generic get gongstruct field value
func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagram.IsChecked)
		res.valueBool = diagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagram.IsEditable_)
		res.valueBool = diagram.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInRenameMode":
		res.valueString = fmt.Sprintf("%t", diagram.IsInRenameMode)
		res.valueBool = diagram.IsInRenameMode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsExpanded)
		res.valueBool = diagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = diagram.ComputedPrefix
	case "Product_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Product_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "ProductsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ProductsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsPBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsPBSNodeExpanded)
		res.valueBool = diagram.IsPBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ProductComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ProductComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsWBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsWBSNodeExpanded)
		res.valueBool = diagram.IsWBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Task_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Task_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "TasksWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "TasksWhoseInputNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhoseInputNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "TaskComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "TaskInputShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskInputShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (product *Product) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = product.Name
	case "SubProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range product.SubProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsExpanded)
		res.valueBool = product.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = product.ComputedPrefix
	case "IsProducersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsProducersNodeExpanded)
		res.valueBool = product.IsProducersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsConsumersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsConsumersNodeExpanded)
		res.valueBool = product.IsConsumersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (productcompositionshape *ProductCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = productcompositionshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productcompositionshape.Product != nil {
			res.valueString = productcompositionshape.Product.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, productcompositionshape.Product))
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.StartRatio)
		res.valueFloat = productcompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.EndRatio)
		res.valueFloat = productcompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := productcompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := productcompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.CornerOffsetRatio)
		res.valueFloat = productcompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (productshape *ProductShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = productshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productshape.Product != nil {
			res.valueString = productshape.Product.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, productshape.Product))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", productshape.IsExpanded)
		res.valueBool = productshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", productshape.X)
		res.valueFloat = productshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", productshape.Y)
		res.valueFloat = productshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", productshape.Width)
		res.valueFloat = productshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", productshape.Height)
		res.valueFloat = productshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (project *Project) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = project.Name
	case "RootProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range project.RootProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsPBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", project.IsPBSNodeExpanded)
		res.valueBool = project.IsPBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RootTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range project.RootTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsWBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", project.IsWBSNodeExpanded)
		res.valueBool = project.IsWBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range project.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsDiagramsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", project.IsDiagramsNodeExpanded)
		res.valueBool = project.IsDiagramsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", project.IsExpanded)
		res.valueBool = project.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = project.ComputedPrefix
	}
	return
}
func (root *Root) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = root.Name
	case "Projects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range root.Projects {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "OrphanedProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range root.OrphanedProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "OrphanedTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range root.OrphanedTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", root.NbPixPerCharacter)
		res.valueFloat = root.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (task *Task) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = task.Name
	case "SubTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.SubTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsExpanded)
		res.valueBool = task.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = task.ComputedPrefix
	case "Inputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.Inputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsInputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsInputsNodeExpanded)
		res.valueBool = task.IsInputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Outputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.Outputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsOutputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsOutputsNodeExpanded)
		res.valueBool = task.IsOutputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (taskcompositionshape *TaskCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskcompositionshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskcompositionshape.Task != nil {
			res.valueString = taskcompositionshape.Task.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, taskcompositionshape.Task))
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.StartRatio)
		res.valueFloat = taskcompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.EndRatio)
		res.valueFloat = taskcompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskcompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskcompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.CornerOffsetRatio)
		res.valueFloat = taskcompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (taskinputshape *TaskInputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskinputshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskinputshape.Task != nil {
			res.valueString = taskinputshape.Task.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, taskinputshape.Task))
		}
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskinputshape.Product != nil {
			res.valueString = taskinputshape.Product.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, taskinputshape.Product))
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.StartRatio)
		res.valueFloat = taskinputshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.EndRatio)
		res.valueFloat = taskinputshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskinputshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskinputshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.CornerOffsetRatio)
		res.valueFloat = taskinputshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (taskshape *TaskShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskshape.Task != nil {
			res.valueString = taskshape.Task.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, taskshape.Task))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", taskshape.IsExpanded)
		res.valueBool = taskshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", taskshape.X)
		res.valueFloat = taskshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", taskshape.Y)
		res.valueFloat = taskshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", taskshape.Width)
		res.valueFloat = taskshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", taskshape.Height)
		res.valueFloat = taskshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (diagram *Diagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagram.Name = value.GetValueString()
	case "IsChecked":
		diagram.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagram.IsEditable_ = value.GetValueBool()
	case "IsInRenameMode":
		diagram.IsInRenameMode = value.GetValueBool()
	case "IsExpanded":
		diagram.IsExpanded = value.GetValueBool()
	case "ComputedPrefix":
		diagram.ComputedPrefix = value.GetValueString()
	case "Product_Shapes":
		diagram.Product_Shapes = make([]*ProductShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProductShapes {
					if stage.ProductShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Product_Shapes = append(diagram.Product_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ProductsWhoseNodeIsExpanded":
		diagram.ProductsWhoseNodeIsExpanded = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						diagram.ProductsWhoseNodeIsExpanded = append(diagram.ProductsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsPBSNodeExpanded":
		diagram.IsPBSNodeExpanded = value.GetValueBool()
	case "ProductComposition_Shapes":
		diagram.ProductComposition_Shapes = make([]*ProductCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProductCompositionShapes {
					if stage.ProductCompositionShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.ProductComposition_Shapes = append(diagram.ProductComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsWBSNodeExpanded":
		diagram.IsWBSNodeExpanded = value.GetValueBool()
	case "Task_Shapes":
		diagram.Task_Shapes = make([]*TaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskShapes {
					if stage.TaskShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Task_Shapes = append(diagram.Task_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseNodeIsExpanded":
		diagram.TasksWhoseNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						diagram.TasksWhoseNodeIsExpanded = append(diagram.TasksWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseInputNodeIsExpanded":
		diagram.TasksWhoseInputNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						diagram.TasksWhoseInputNodeIsExpanded = append(diagram.TasksWhoseInputNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TaskComposition_Shapes":
		diagram.TaskComposition_Shapes = make([]*TaskCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskCompositionShapes {
					if stage.TaskCompositionShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.TaskComposition_Shapes = append(diagram.TaskComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "TaskInputShapes":
		diagram.TaskInputShapes = make([]*TaskInputShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskInputShapes {
					if stage.TaskInputShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.TaskInputShapes = append(diagram.TaskInputShapes, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (product *Product) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		product.Name = value.GetValueString()
	case "SubProducts":
		product.SubProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						product.SubProducts = append(product.SubProducts, __instance__)
						break
					}
				}
			}
		}
	case "IsExpanded":
		product.IsExpanded = value.GetValueBool()
	case "ComputedPrefix":
		product.ComputedPrefix = value.GetValueString()
	case "IsProducersNodeExpanded":
		product.IsProducersNodeExpanded = value.GetValueBool()
	case "IsConsumersNodeExpanded":
		product.IsConsumersNodeExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (productcompositionshape *ProductCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		productcompositionshape.Name = value.GetValueString()
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			productcompositionshape.Product = nil
			for __instance__ := range stage.Products {
				if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
					productcompositionshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		productcompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		productcompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		productcompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		productcompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		productcompositionshape.CornerOffsetRatio = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (productshape *ProductShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		productshape.Name = value.GetValueString()
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			productshape.Product = nil
			for __instance__ := range stage.Products {
				if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
					productshape.Product = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		productshape.IsExpanded = value.GetValueBool()
	case "X":
		productshape.X = value.GetValueFloat()
	case "Y":
		productshape.Y = value.GetValueFloat()
	case "Width":
		productshape.Width = value.GetValueFloat()
	case "Height":
		productshape.Height = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (project *Project) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		project.Name = value.GetValueString()
	case "RootProducts":
		project.RootProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						project.RootProducts = append(project.RootProducts, __instance__)
						break
					}
				}
			}
		}
	case "IsPBSNodeExpanded":
		project.IsPBSNodeExpanded = value.GetValueBool()
	case "RootTasks":
		project.RootTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						project.RootTasks = append(project.RootTasks, __instance__)
						break
					}
				}
			}
		}
	case "IsWBSNodeExpanded":
		project.IsWBSNodeExpanded = value.GetValueBool()
	case "Diagrams":
		project.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.DiagramMap_Staged_Order[__instance__] == uint(id) {
						project.Diagrams = append(project.Diagrams, __instance__)
						break
					}
				}
			}
		}
	case "IsDiagramsNodeExpanded":
		project.IsDiagramsNodeExpanded = value.GetValueBool()
	case "IsExpanded":
		project.IsExpanded = value.GetValueBool()
	case "ComputedPrefix":
		project.ComputedPrefix = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (root *Root) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		root.Name = value.GetValueString()
	case "Projects":
		root.Projects = make([]*Project, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Projects {
					if stage.ProjectMap_Staged_Order[__instance__] == uint(id) {
						root.Projects = append(root.Projects, __instance__)
						break
					}
				}
			}
		}
	case "OrphanedProducts":
		root.OrphanedProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						root.OrphanedProducts = append(root.OrphanedProducts, __instance__)
						break
					}
				}
			}
		}
	case "OrphanedTasks":
		root.OrphanedTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						root.OrphanedTasks = append(root.OrphanedTasks, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		root.NbPixPerCharacter = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (task *Task) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		task.Name = value.GetValueString()
	case "SubTasks":
		task.SubTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						task.SubTasks = append(task.SubTasks, __instance__)
						break
					}
				}
			}
		}
	case "IsExpanded":
		task.IsExpanded = value.GetValueBool()
	case "ComputedPrefix":
		task.ComputedPrefix = value.GetValueString()
	case "Inputs":
		task.Inputs = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						task.Inputs = append(task.Inputs, __instance__)
						break
					}
				}
			}
		}
	case "IsInputsNodeExpanded":
		task.IsInputsNodeExpanded = value.GetValueBool()
	case "Outputs":
		task.Outputs = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						task.Outputs = append(task.Outputs, __instance__)
						break
					}
				}
			}
		}
	case "IsOutputsNodeExpanded":
		task.IsOutputsNodeExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskcompositionshape *TaskCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskcompositionshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskcompositionshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
					taskcompositionshape.Task = __instance__
					break
				}
			}
		}
	case "StartRatio":
		taskcompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		taskcompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		taskcompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		taskcompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		taskcompositionshape.CornerOffsetRatio = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskinputshape *TaskInputShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskinputshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskinputshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
					taskinputshape.Task = __instance__
					break
				}
			}
		}
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskinputshape.Product = nil
			for __instance__ := range stage.Products {
				if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
					taskinputshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		taskinputshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		taskinputshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		taskinputshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		taskinputshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		taskinputshape.CornerOffsetRatio = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskshape *TaskShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
					taskshape.Task = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		taskshape.IsExpanded = value.GetValueBool()
	case "X":
		taskshape.X = value.GetValueFloat()
	case "Y":
		taskshape.Y = value.GetValueFloat()
	case "Width":
		taskshape.Width = value.GetValueFloat()
	case "Height":
		taskshape.Height = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (product *Product) GongGetGongstructName() string {
	return "Product"
}

func (productcompositionshape *ProductCompositionShape) GongGetGongstructName() string {
	return "ProductCompositionShape"
}

func (productshape *ProductShape) GongGetGongstructName() string {
	return "ProductShape"
}

func (project *Project) GongGetGongstructName() string {
	return "Project"
}

func (root *Root) GongGetGongstructName() string {
	return "Root"
}

func (task *Task) GongGetGongstructName() string {
	return "Task"
}

func (taskcompositionshape *TaskCompositionShape) GongGetGongstructName() string {
	return "TaskCompositionShape"
}

func (taskinputshape *TaskInputShape) GongGetGongstructName() string {
	return "TaskInputShape"
}

func (taskshape *TaskShape) GongGetGongstructName() string {
	return "TaskShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Diagrams_mapString = make(map[string]*Diagram)
	for diagram := range stage.Diagrams {
		stage.Diagrams_mapString[diagram.Name] = diagram
	}

	stage.Products_mapString = make(map[string]*Product)
	for product := range stage.Products {
		stage.Products_mapString[product.Name] = product
	}

	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	for productcompositionshape := range stage.ProductCompositionShapes {
		stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape
	}

	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	for productshape := range stage.ProductShapes {
		stage.ProductShapes_mapString[productshape.Name] = productshape
	}

	stage.Projects_mapString = make(map[string]*Project)
	for project := range stage.Projects {
		stage.Projects_mapString[project.Name] = project
	}

	stage.Roots_mapString = make(map[string]*Root)
	for root := range stage.Roots {
		stage.Roots_mapString[root.Name] = root
	}

	stage.Tasks_mapString = make(map[string]*Task)
	for task := range stage.Tasks {
		stage.Tasks_mapString[task.Name] = task
	}

	stage.TaskCompositionShapes_mapString = make(map[string]*TaskCompositionShape)
	for taskcompositionshape := range stage.TaskCompositionShapes {
		stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape
	}

	stage.TaskInputShapes_mapString = make(map[string]*TaskInputShape)
	for taskinputshape := range stage.TaskInputShapes {
		stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape
	}

	stage.TaskShapes_mapString = make(map[string]*TaskShape)
	for taskshape := range stage.TaskShapes {
		stage.TaskShapes_mapString[taskshape.Name] = taskshape
	}

}
// Last line of the template
