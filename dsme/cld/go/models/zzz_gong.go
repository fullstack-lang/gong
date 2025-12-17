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

	cld_go "github.com/fullstack-lang/gong/dsme/cld/go"
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
	Category1s           map[*Category1]struct{}
	Category1s_mapString map[string]*Category1

	// insertion point for slice of pointers maps
	OnAfterCategory1CreateCallback OnAfterCreateInterface[Category1]
	OnAfterCategory1UpdateCallback OnAfterUpdateInterface[Category1]
	OnAfterCategory1DeleteCallback OnAfterDeleteInterface[Category1]
	OnAfterCategory1ReadCallback   OnAfterReadInterface[Category1]

	Category1Shapes           map[*Category1Shape]struct{}
	Category1Shapes_mapString map[string]*Category1Shape

	// insertion point for slice of pointers maps
	OnAfterCategory1ShapeCreateCallback OnAfterCreateInterface[Category1Shape]
	OnAfterCategory1ShapeUpdateCallback OnAfterUpdateInterface[Category1Shape]
	OnAfterCategory1ShapeDeleteCallback OnAfterDeleteInterface[Category1Shape]
	OnAfterCategory1ShapeReadCallback   OnAfterReadInterface[Category1Shape]

	Category2s           map[*Category2]struct{}
	Category2s_mapString map[string]*Category2

	// insertion point for slice of pointers maps
	OnAfterCategory2CreateCallback OnAfterCreateInterface[Category2]
	OnAfterCategory2UpdateCallback OnAfterUpdateInterface[Category2]
	OnAfterCategory2DeleteCallback OnAfterDeleteInterface[Category2]
	OnAfterCategory2ReadCallback   OnAfterReadInterface[Category2]

	Category2Shapes           map[*Category2Shape]struct{}
	Category2Shapes_mapString map[string]*Category2Shape

	// insertion point for slice of pointers maps
	OnAfterCategory2ShapeCreateCallback OnAfterCreateInterface[Category2Shape]
	OnAfterCategory2ShapeUpdateCallback OnAfterUpdateInterface[Category2Shape]
	OnAfterCategory2ShapeDeleteCallback OnAfterDeleteInterface[Category2Shape]
	OnAfterCategory2ShapeReadCallback   OnAfterReadInterface[Category2Shape]

	Category3s           map[*Category3]struct{}
	Category3s_mapString map[string]*Category3

	// insertion point for slice of pointers maps
	OnAfterCategory3CreateCallback OnAfterCreateInterface[Category3]
	OnAfterCategory3UpdateCallback OnAfterUpdateInterface[Category3]
	OnAfterCategory3DeleteCallback OnAfterDeleteInterface[Category3]
	OnAfterCategory3ReadCallback   OnAfterReadInterface[Category3]

	Category3Shapes           map[*Category3Shape]struct{}
	Category3Shapes_mapString map[string]*Category3Shape

	// insertion point for slice of pointers maps
	OnAfterCategory3ShapeCreateCallback OnAfterCreateInterface[Category3Shape]
	OnAfterCategory3ShapeUpdateCallback OnAfterUpdateInterface[Category3Shape]
	OnAfterCategory3ShapeDeleteCallback OnAfterDeleteInterface[Category3Shape]
	OnAfterCategory3ShapeReadCallback   OnAfterReadInterface[Category3Shape]

	ControlPointShapes           map[*ControlPointShape]struct{}
	ControlPointShapes_mapString map[string]*ControlPointShape

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback OnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback OnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback OnAfterDeleteInterface[ControlPointShape]
	OnAfterControlPointShapeReadCallback   OnAfterReadInterface[ControlPointShape]

	Desks           map[*Desk]struct{}
	Desks_mapString map[string]*Desk

	// insertion point for slice of pointers maps
	OnAfterDeskCreateCallback OnAfterCreateInterface[Desk]
	OnAfterDeskUpdateCallback OnAfterUpdateInterface[Desk]
	OnAfterDeskDeleteCallback OnAfterDeleteInterface[Desk]
	OnAfterDeskReadCallback   OnAfterReadInterface[Desk]

	Diagrams           map[*Diagram]struct{}
	Diagrams_mapString map[string]*Diagram

	// insertion point for slice of pointers maps
	Diagram_Category1Shapes_reverseMap map[*Category1Shape]*Diagram

	Diagram_Category2Shapes_reverseMap map[*Category2Shape]*Diagram

	Diagram_Category3Shapes_reverseMap map[*Category3Shape]*Diagram

	Diagram_InfluenceShapes_reverseMap map[*InfluenceShape]*Diagram

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Influences           map[*Influence]struct{}
	Influences_mapString map[string]*Influence

	// insertion point for slice of pointers maps
	OnAfterInfluenceCreateCallback OnAfterCreateInterface[Influence]
	OnAfterInfluenceUpdateCallback OnAfterUpdateInterface[Influence]
	OnAfterInfluenceDeleteCallback OnAfterDeleteInterface[Influence]
	OnAfterInfluenceReadCallback   OnAfterReadInterface[Influence]

	InfluenceShapes           map[*InfluenceShape]struct{}
	InfluenceShapes_mapString map[string]*InfluenceShape

	// insertion point for slice of pointers maps
	InfluenceShape_ControlPointShapes_reverseMap map[*ControlPointShape]*InfluenceShape

	OnAfterInfluenceShapeCreateCallback OnAfterCreateInterface[InfluenceShape]
	OnAfterInfluenceShapeUpdateCallback OnAfterUpdateInterface[InfluenceShape]
	OnAfterInfluenceShapeDeleteCallback OnAfterDeleteInterface[InfluenceShape]
	OnAfterInfluenceShapeReadCallback   OnAfterReadInterface[InfluenceShape]

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
	Category1Order            uint
	Category1Map_Staged_Order map[*Category1]uint

	Category1ShapeOrder            uint
	Category1ShapeMap_Staged_Order map[*Category1Shape]uint

	Category2Order            uint
	Category2Map_Staged_Order map[*Category2]uint

	Category2ShapeOrder            uint
	Category2ShapeMap_Staged_Order map[*Category2Shape]uint

	Category3Order            uint
	Category3Map_Staged_Order map[*Category3]uint

	Category3ShapeOrder            uint
	Category3ShapeMap_Staged_Order map[*Category3Shape]uint

	ControlPointShapeOrder            uint
	ControlPointShapeMap_Staged_Order map[*ControlPointShape]uint

	DeskOrder            uint
	DeskMap_Staged_Order map[*Desk]uint

	DiagramOrder            uint
	DiagramMap_Staged_Order map[*Diagram]uint

	InfluenceOrder            uint
	InfluenceMap_Staged_Order map[*Influence]uint

	InfluenceShapeOrder            uint
	InfluenceShapeMap_Staged_Order map[*InfluenceShape]uint

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
	case *Category1:
		tmp := GetStructInstancesByOrder(stage.Category1s, stage.Category1Map_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category1 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Category1Shape:
		tmp := GetStructInstancesByOrder(stage.Category1Shapes, stage.Category1ShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category1Shape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Category2:
		tmp := GetStructInstancesByOrder(stage.Category2s, stage.Category2Map_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category2 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Category2Shape:
		tmp := GetStructInstancesByOrder(stage.Category2Shapes, stage.Category2ShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category2Shape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Category3:
		tmp := GetStructInstancesByOrder(stage.Category3s, stage.Category3Map_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category3 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Category3Shape:
		tmp := GetStructInstancesByOrder(stage.Category3Shapes, stage.Category3ShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Category3Shape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlPointShape:
		tmp := GetStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlPointShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Desk:
		tmp := GetStructInstancesByOrder(stage.Desks, stage.DeskMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Desk implements.
			res = append(res, any(v).(T))
		}
		return res
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
	case *Influence:
		tmp := GetStructInstancesByOrder(stage.Influences, stage.InfluenceMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Influence implements.
			res = append(res, any(v).(T))
		}
		return res
	case *InfluenceShape:
		tmp := GetStructInstancesByOrder(stage.InfluenceShapes, stage.InfluenceShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *InfluenceShape implements.
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
	case "Category1":
		res = GetNamedStructInstances(stage.Category1s, stage.Category1Map_Staged_Order)
	case "Category1Shape":
		res = GetNamedStructInstances(stage.Category1Shapes, stage.Category1ShapeMap_Staged_Order)
	case "Category2":
		res = GetNamedStructInstances(stage.Category2s, stage.Category2Map_Staged_Order)
	case "Category2Shape":
		res = GetNamedStructInstances(stage.Category2Shapes, stage.Category2ShapeMap_Staged_Order)
	case "Category3":
		res = GetNamedStructInstances(stage.Category3s, stage.Category3Map_Staged_Order)
	case "Category3Shape":
		res = GetNamedStructInstances(stage.Category3Shapes, stage.Category3ShapeMap_Staged_Order)
	case "ControlPointShape":
		res = GetNamedStructInstances(stage.ControlPointShapes, stage.ControlPointShapeMap_Staged_Order)
	case "Desk":
		res = GetNamedStructInstances(stage.Desks, stage.DeskMap_Staged_Order)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.DiagramMap_Staged_Order)
	case "Influence":
		res = GetNamedStructInstances(stage.Influences, stage.InfluenceMap_Staged_Order)
	case "InfluenceShape":
		res = GetNamedStructInstances(stage.InfluenceShapes, stage.InfluenceShapeMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/dsme/cld/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return cld_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return cld_go.GoDiagramsDir
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
	CommitCategory1(category1 *Category1)
	CheckoutCategory1(category1 *Category1)
	CommitCategory1Shape(category1shape *Category1Shape)
	CheckoutCategory1Shape(category1shape *Category1Shape)
	CommitCategory2(category2 *Category2)
	CheckoutCategory2(category2 *Category2)
	CommitCategory2Shape(category2shape *Category2Shape)
	CheckoutCategory2Shape(category2shape *Category2Shape)
	CommitCategory3(category3 *Category3)
	CheckoutCategory3(category3 *Category3)
	CommitCategory3Shape(category3shape *Category3Shape)
	CheckoutCategory3Shape(category3shape *Category3Shape)
	CommitControlPointShape(controlpointshape *ControlPointShape)
	CheckoutControlPointShape(controlpointshape *ControlPointShape)
	CommitDesk(desk *Desk)
	CheckoutDesk(desk *Desk)
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitInfluence(influence *Influence)
	CheckoutInfluence(influence *Influence)
	CommitInfluenceShape(influenceshape *InfluenceShape)
	CheckoutInfluenceShape(influenceshape *InfluenceShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Category1s:           make(map[*Category1]struct{}),
		Category1s_mapString: make(map[string]*Category1),

		Category1Shapes:           make(map[*Category1Shape]struct{}),
		Category1Shapes_mapString: make(map[string]*Category1Shape),

		Category2s:           make(map[*Category2]struct{}),
		Category2s_mapString: make(map[string]*Category2),

		Category2Shapes:           make(map[*Category2Shape]struct{}),
		Category2Shapes_mapString: make(map[string]*Category2Shape),

		Category3s:           make(map[*Category3]struct{}),
		Category3s_mapString: make(map[string]*Category3),

		Category3Shapes:           make(map[*Category3Shape]struct{}),
		Category3Shapes_mapString: make(map[string]*Category3Shape),

		ControlPointShapes:           make(map[*ControlPointShape]struct{}),
		ControlPointShapes_mapString: make(map[string]*ControlPointShape),

		Desks:           make(map[*Desk]struct{}),
		Desks_mapString: make(map[string]*Desk),

		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Influences:           make(map[*Influence]struct{}),
		Influences_mapString: make(map[string]*Influence),

		InfluenceShapes:           make(map[*InfluenceShape]struct{}),
		InfluenceShapes_mapString: make(map[string]*InfluenceShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Category1Map_Staged_Order: make(map[*Category1]uint),

		Category1ShapeMap_Staged_Order: make(map[*Category1Shape]uint),

		Category2Map_Staged_Order: make(map[*Category2]uint),

		Category2ShapeMap_Staged_Order: make(map[*Category2Shape]uint),

		Category3Map_Staged_Order: make(map[*Category3]uint),

		Category3ShapeMap_Staged_Order: make(map[*Category3Shape]uint),

		ControlPointShapeMap_Staged_Order: make(map[*ControlPointShape]uint),

		DeskMap_Staged_Order: make(map[*Desk]uint),

		DiagramMap_Staged_Order: make(map[*Diagram]uint),

		InfluenceMap_Staged_Order: make(map[*Influence]uint),

		InfluenceShapeMap_Staged_Order: make(map[*InfluenceShape]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Category1"},
			{name: "Category1Shape"},
			{name: "Category2"},
			{name: "Category2Shape"},
			{name: "Category3"},
			{name: "Category3Shape"},
			{name: "ControlPointShape"},
			{name: "Desk"},
			{name: "Diagram"},
			{name: "Influence"},
			{name: "InfluenceShape"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Category1:
		return stage.Category1Map_Staged_Order[instance]
	case *Category1Shape:
		return stage.Category1ShapeMap_Staged_Order[instance]
	case *Category2:
		return stage.Category2Map_Staged_Order[instance]
	case *Category2Shape:
		return stage.Category2ShapeMap_Staged_Order[instance]
	case *Category3:
		return stage.Category3Map_Staged_Order[instance]
	case *Category3Shape:
		return stage.Category3ShapeMap_Staged_Order[instance]
	case *ControlPointShape:
		return stage.ControlPointShapeMap_Staged_Order[instance]
	case *Desk:
		return stage.DeskMap_Staged_Order[instance]
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Influence:
		return stage.InfluenceMap_Staged_Order[instance]
	case *InfluenceShape:
		return stage.InfluenceShapeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Category1:
		return stage.Category1Map_Staged_Order[instance]
	case *Category1Shape:
		return stage.Category1ShapeMap_Staged_Order[instance]
	case *Category2:
		return stage.Category2Map_Staged_Order[instance]
	case *Category2Shape:
		return stage.Category2ShapeMap_Staged_Order[instance]
	case *Category3:
		return stage.Category3Map_Staged_Order[instance]
	case *Category3Shape:
		return stage.Category3ShapeMap_Staged_Order[instance]
	case *ControlPointShape:
		return stage.ControlPointShapeMap_Staged_Order[instance]
	case *Desk:
		return stage.DeskMap_Staged_Order[instance]
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Influence:
		return stage.InfluenceMap_Staged_Order[instance]
	case *InfluenceShape:
		return stage.InfluenceShapeMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Category1"] = len(stage.Category1s)
	stage.Map_GongStructName_InstancesNb["Category1Shape"] = len(stage.Category1Shapes)
	stage.Map_GongStructName_InstancesNb["Category2"] = len(stage.Category2s)
	stage.Map_GongStructName_InstancesNb["Category2Shape"] = len(stage.Category2Shapes)
	stage.Map_GongStructName_InstancesNb["Category3"] = len(stage.Category3s)
	stage.Map_GongStructName_InstancesNb["Category3Shape"] = len(stage.Category3Shapes)
	stage.Map_GongStructName_InstancesNb["ControlPointShape"] = len(stage.ControlPointShapes)
	stage.Map_GongStructName_InstancesNb["Desk"] = len(stage.Desks)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Influence"] = len(stage.Influences)
	stage.Map_GongStructName_InstancesNb["InfluenceShape"] = len(stage.InfluenceShapes)
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
// Stage puts category1 to the model stage
func (category1 *Category1) Stage(stage *Stage) *Category1 {

	if _, ok := stage.Category1s[category1]; !ok {
		stage.Category1s[category1] = struct{}{}
		stage.Category1Map_Staged_Order[category1] = stage.Category1Order
		stage.Category1Order++
	}
	stage.Category1s_mapString[category1.Name] = category1

	return category1
}

// StagePreserveOrder puts category1 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category1Order
// - update stage.Category1Order accordingly
func (category1 *Category1) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category1s[category1]; !ok {
		stage.Category1s[category1] = struct{}{}

		if order > stage.Category1Order {
			stage.Category1Order = order
		}
		stage.Category1Map_Staged_Order[category1] = stage.Category1Order
		stage.Category1Order++
	}
	stage.Category1s_mapString[category1.Name] = category1
}

// Unstage removes category1 off the model stage
func (category1 *Category1) Unstage(stage *Stage) *Category1 {
	delete(stage.Category1s, category1)
	delete(stage.Category1Map_Staged_Order, category1)
	delete(stage.Category1s_mapString, category1.Name)

	return category1
}

// UnstageVoid removes category1 off the model stage
func (category1 *Category1) UnstageVoid(stage *Stage) {
	delete(stage.Category1s, category1)
	delete(stage.Category1Map_Staged_Order, category1)
	delete(stage.Category1s_mapString, category1.Name)
}

// commit category1 to the back repo (if it is already staged)
func (category1 *Category1) Commit(stage *Stage) *Category1 {
	if _, ok := stage.Category1s[category1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory1(category1)
		}
	}
	return category1
}

func (category1 *Category1) CommitVoid(stage *Stage) {
	category1.Commit(stage)
}

func (category1 *Category1) StageVoid(stage *Stage) {
	category1.Stage(stage)
}

// Checkout category1 to the back repo (if it is already staged)
func (category1 *Category1) Checkout(stage *Stage) *Category1 {
	if _, ok := stage.Category1s[category1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory1(category1)
		}
	}
	return category1
}

// for satisfaction of GongStruct interface
func (category1 *Category1) GetName() (res string) {
	return category1.Name
}

// for satisfaction of GongStruct interface
func (category1 *Category1) SetName(name string) (){
	category1.Name = name
}

// Stage puts category1shape to the model stage
func (category1shape *Category1Shape) Stage(stage *Stage) *Category1Shape {

	if _, ok := stage.Category1Shapes[category1shape]; !ok {
		stage.Category1Shapes[category1shape] = struct{}{}
		stage.Category1ShapeMap_Staged_Order[category1shape] = stage.Category1ShapeOrder
		stage.Category1ShapeOrder++
	}
	stage.Category1Shapes_mapString[category1shape.Name] = category1shape

	return category1shape
}

// StagePreserveOrder puts category1shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category1ShapeOrder
// - update stage.Category1ShapeOrder accordingly
func (category1shape *Category1Shape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category1Shapes[category1shape]; !ok {
		stage.Category1Shapes[category1shape] = struct{}{}

		if order > stage.Category1ShapeOrder {
			stage.Category1ShapeOrder = order
		}
		stage.Category1ShapeMap_Staged_Order[category1shape] = stage.Category1ShapeOrder
		stage.Category1ShapeOrder++
	}
	stage.Category1Shapes_mapString[category1shape.Name] = category1shape
}

// Unstage removes category1shape off the model stage
func (category1shape *Category1Shape) Unstage(stage *Stage) *Category1Shape {
	delete(stage.Category1Shapes, category1shape)
	delete(stage.Category1ShapeMap_Staged_Order, category1shape)
	delete(stage.Category1Shapes_mapString, category1shape.Name)

	return category1shape
}

// UnstageVoid removes category1shape off the model stage
func (category1shape *Category1Shape) UnstageVoid(stage *Stage) {
	delete(stage.Category1Shapes, category1shape)
	delete(stage.Category1ShapeMap_Staged_Order, category1shape)
	delete(stage.Category1Shapes_mapString, category1shape.Name)
}

// commit category1shape to the back repo (if it is already staged)
func (category1shape *Category1Shape) Commit(stage *Stage) *Category1Shape {
	if _, ok := stage.Category1Shapes[category1shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory1Shape(category1shape)
		}
	}
	return category1shape
}

func (category1shape *Category1Shape) CommitVoid(stage *Stage) {
	category1shape.Commit(stage)
}

func (category1shape *Category1Shape) StageVoid(stage *Stage) {
	category1shape.Stage(stage)
}

// Checkout category1shape to the back repo (if it is already staged)
func (category1shape *Category1Shape) Checkout(stage *Stage) *Category1Shape {
	if _, ok := stage.Category1Shapes[category1shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory1Shape(category1shape)
		}
	}
	return category1shape
}

// for satisfaction of GongStruct interface
func (category1shape *Category1Shape) GetName() (res string) {
	return category1shape.Name
}

// for satisfaction of GongStruct interface
func (category1shape *Category1Shape) SetName(name string) (){
	category1shape.Name = name
}

// Stage puts category2 to the model stage
func (category2 *Category2) Stage(stage *Stage) *Category2 {

	if _, ok := stage.Category2s[category2]; !ok {
		stage.Category2s[category2] = struct{}{}
		stage.Category2Map_Staged_Order[category2] = stage.Category2Order
		stage.Category2Order++
	}
	stage.Category2s_mapString[category2.Name] = category2

	return category2
}

// StagePreserveOrder puts category2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category2Order
// - update stage.Category2Order accordingly
func (category2 *Category2) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category2s[category2]; !ok {
		stage.Category2s[category2] = struct{}{}

		if order > stage.Category2Order {
			stage.Category2Order = order
		}
		stage.Category2Map_Staged_Order[category2] = stage.Category2Order
		stage.Category2Order++
	}
	stage.Category2s_mapString[category2.Name] = category2
}

// Unstage removes category2 off the model stage
func (category2 *Category2) Unstage(stage *Stage) *Category2 {
	delete(stage.Category2s, category2)
	delete(stage.Category2Map_Staged_Order, category2)
	delete(stage.Category2s_mapString, category2.Name)

	return category2
}

// UnstageVoid removes category2 off the model stage
func (category2 *Category2) UnstageVoid(stage *Stage) {
	delete(stage.Category2s, category2)
	delete(stage.Category2Map_Staged_Order, category2)
	delete(stage.Category2s_mapString, category2.Name)
}

// commit category2 to the back repo (if it is already staged)
func (category2 *Category2) Commit(stage *Stage) *Category2 {
	if _, ok := stage.Category2s[category2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory2(category2)
		}
	}
	return category2
}

func (category2 *Category2) CommitVoid(stage *Stage) {
	category2.Commit(stage)
}

func (category2 *Category2) StageVoid(stage *Stage) {
	category2.Stage(stage)
}

// Checkout category2 to the back repo (if it is already staged)
func (category2 *Category2) Checkout(stage *Stage) *Category2 {
	if _, ok := stage.Category2s[category2]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory2(category2)
		}
	}
	return category2
}

// for satisfaction of GongStruct interface
func (category2 *Category2) GetName() (res string) {
	return category2.Name
}

// for satisfaction of GongStruct interface
func (category2 *Category2) SetName(name string) (){
	category2.Name = name
}

// Stage puts category2shape to the model stage
func (category2shape *Category2Shape) Stage(stage *Stage) *Category2Shape {

	if _, ok := stage.Category2Shapes[category2shape]; !ok {
		stage.Category2Shapes[category2shape] = struct{}{}
		stage.Category2ShapeMap_Staged_Order[category2shape] = stage.Category2ShapeOrder
		stage.Category2ShapeOrder++
	}
	stage.Category2Shapes_mapString[category2shape.Name] = category2shape

	return category2shape
}

// StagePreserveOrder puts category2shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category2ShapeOrder
// - update stage.Category2ShapeOrder accordingly
func (category2shape *Category2Shape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category2Shapes[category2shape]; !ok {
		stage.Category2Shapes[category2shape] = struct{}{}

		if order > stage.Category2ShapeOrder {
			stage.Category2ShapeOrder = order
		}
		stage.Category2ShapeMap_Staged_Order[category2shape] = stage.Category2ShapeOrder
		stage.Category2ShapeOrder++
	}
	stage.Category2Shapes_mapString[category2shape.Name] = category2shape
}

// Unstage removes category2shape off the model stage
func (category2shape *Category2Shape) Unstage(stage *Stage) *Category2Shape {
	delete(stage.Category2Shapes, category2shape)
	delete(stage.Category2ShapeMap_Staged_Order, category2shape)
	delete(stage.Category2Shapes_mapString, category2shape.Name)

	return category2shape
}

// UnstageVoid removes category2shape off the model stage
func (category2shape *Category2Shape) UnstageVoid(stage *Stage) {
	delete(stage.Category2Shapes, category2shape)
	delete(stage.Category2ShapeMap_Staged_Order, category2shape)
	delete(stage.Category2Shapes_mapString, category2shape.Name)
}

// commit category2shape to the back repo (if it is already staged)
func (category2shape *Category2Shape) Commit(stage *Stage) *Category2Shape {
	if _, ok := stage.Category2Shapes[category2shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory2Shape(category2shape)
		}
	}
	return category2shape
}

func (category2shape *Category2Shape) CommitVoid(stage *Stage) {
	category2shape.Commit(stage)
}

func (category2shape *Category2Shape) StageVoid(stage *Stage) {
	category2shape.Stage(stage)
}

// Checkout category2shape to the back repo (if it is already staged)
func (category2shape *Category2Shape) Checkout(stage *Stage) *Category2Shape {
	if _, ok := stage.Category2Shapes[category2shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory2Shape(category2shape)
		}
	}
	return category2shape
}

// for satisfaction of GongStruct interface
func (category2shape *Category2Shape) GetName() (res string) {
	return category2shape.Name
}

// for satisfaction of GongStruct interface
func (category2shape *Category2Shape) SetName(name string) (){
	category2shape.Name = name
}

// Stage puts category3 to the model stage
func (category3 *Category3) Stage(stage *Stage) *Category3 {

	if _, ok := stage.Category3s[category3]; !ok {
		stage.Category3s[category3] = struct{}{}
		stage.Category3Map_Staged_Order[category3] = stage.Category3Order
		stage.Category3Order++
	}
	stage.Category3s_mapString[category3.Name] = category3

	return category3
}

// StagePreserveOrder puts category3 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category3Order
// - update stage.Category3Order accordingly
func (category3 *Category3) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category3s[category3]; !ok {
		stage.Category3s[category3] = struct{}{}

		if order > stage.Category3Order {
			stage.Category3Order = order
		}
		stage.Category3Map_Staged_Order[category3] = stage.Category3Order
		stage.Category3Order++
	}
	stage.Category3s_mapString[category3.Name] = category3
}

// Unstage removes category3 off the model stage
func (category3 *Category3) Unstage(stage *Stage) *Category3 {
	delete(stage.Category3s, category3)
	delete(stage.Category3Map_Staged_Order, category3)
	delete(stage.Category3s_mapString, category3.Name)

	return category3
}

// UnstageVoid removes category3 off the model stage
func (category3 *Category3) UnstageVoid(stage *Stage) {
	delete(stage.Category3s, category3)
	delete(stage.Category3Map_Staged_Order, category3)
	delete(stage.Category3s_mapString, category3.Name)
}

// commit category3 to the back repo (if it is already staged)
func (category3 *Category3) Commit(stage *Stage) *Category3 {
	if _, ok := stage.Category3s[category3]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory3(category3)
		}
	}
	return category3
}

func (category3 *Category3) CommitVoid(stage *Stage) {
	category3.Commit(stage)
}

func (category3 *Category3) StageVoid(stage *Stage) {
	category3.Stage(stage)
}

// Checkout category3 to the back repo (if it is already staged)
func (category3 *Category3) Checkout(stage *Stage) *Category3 {
	if _, ok := stage.Category3s[category3]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory3(category3)
		}
	}
	return category3
}

// for satisfaction of GongStruct interface
func (category3 *Category3) GetName() (res string) {
	return category3.Name
}

// for satisfaction of GongStruct interface
func (category3 *Category3) SetName(name string) (){
	category3.Name = name
}

// Stage puts category3shape to the model stage
func (category3shape *Category3Shape) Stage(stage *Stage) *Category3Shape {

	if _, ok := stage.Category3Shapes[category3shape]; !ok {
		stage.Category3Shapes[category3shape] = struct{}{}
		stage.Category3ShapeMap_Staged_Order[category3shape] = stage.Category3ShapeOrder
		stage.Category3ShapeOrder++
	}
	stage.Category3Shapes_mapString[category3shape.Name] = category3shape

	return category3shape
}

// StagePreserveOrder puts category3shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Category3ShapeOrder
// - update stage.Category3ShapeOrder accordingly
func (category3shape *Category3Shape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Category3Shapes[category3shape]; !ok {
		stage.Category3Shapes[category3shape] = struct{}{}

		if order > stage.Category3ShapeOrder {
			stage.Category3ShapeOrder = order
		}
		stage.Category3ShapeMap_Staged_Order[category3shape] = stage.Category3ShapeOrder
		stage.Category3ShapeOrder++
	}
	stage.Category3Shapes_mapString[category3shape.Name] = category3shape
}

// Unstage removes category3shape off the model stage
func (category3shape *Category3Shape) Unstage(stage *Stage) *Category3Shape {
	delete(stage.Category3Shapes, category3shape)
	delete(stage.Category3ShapeMap_Staged_Order, category3shape)
	delete(stage.Category3Shapes_mapString, category3shape.Name)

	return category3shape
}

// UnstageVoid removes category3shape off the model stage
func (category3shape *Category3Shape) UnstageVoid(stage *Stage) {
	delete(stage.Category3Shapes, category3shape)
	delete(stage.Category3ShapeMap_Staged_Order, category3shape)
	delete(stage.Category3Shapes_mapString, category3shape.Name)
}

// commit category3shape to the back repo (if it is already staged)
func (category3shape *Category3Shape) Commit(stage *Stage) *Category3Shape {
	if _, ok := stage.Category3Shapes[category3shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCategory3Shape(category3shape)
		}
	}
	return category3shape
}

func (category3shape *Category3Shape) CommitVoid(stage *Stage) {
	category3shape.Commit(stage)
}

func (category3shape *Category3Shape) StageVoid(stage *Stage) {
	category3shape.Stage(stage)
}

// Checkout category3shape to the back repo (if it is already staged)
func (category3shape *Category3Shape) Checkout(stage *Stage) *Category3Shape {
	if _, ok := stage.Category3Shapes[category3shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCategory3Shape(category3shape)
		}
	}
	return category3shape
}

// for satisfaction of GongStruct interface
func (category3shape *Category3Shape) GetName() (res string) {
	return category3shape.Name
}

// for satisfaction of GongStruct interface
func (category3shape *Category3Shape) SetName(name string) (){
	category3shape.Name = name
}

// Stage puts controlpointshape to the model stage
func (controlpointshape *ControlPointShape) Stage(stage *Stage) *ControlPointShape {

	if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
		stage.ControlPointShapes[controlpointshape] = struct{}{}
		stage.ControlPointShapeMap_Staged_Order[controlpointshape] = stage.ControlPointShapeOrder
		stage.ControlPointShapeOrder++
	}
	stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape

	return controlpointshape
}

// StagePreserveOrder puts controlpointshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointShapeOrder
// - update stage.ControlPointShapeOrder accordingly
func (controlpointshape *ControlPointShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
		stage.ControlPointShapes[controlpointshape] = struct{}{}

		if order > stage.ControlPointShapeOrder {
			stage.ControlPointShapeOrder = order
		}
		stage.ControlPointShapeMap_Staged_Order[controlpointshape] = stage.ControlPointShapeOrder
		stage.ControlPointShapeOrder++
	}
	stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	delete(stage.ControlPointShapes, controlpointshape)
	delete(stage.ControlPointShapeMap_Staged_Order, controlpointshape)
	delete(stage.ControlPointShapes_mapString, controlpointshape.Name)

	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	delete(stage.ControlPointShapes, controlpointshape)
	delete(stage.ControlPointShapeMap_Staged_Order, controlpointshape)
	delete(stage.ControlPointShapes_mapString, controlpointshape.Name)
}

// commit controlpointshape to the back repo (if it is already staged)
func (controlpointshape *ControlPointShape) Commit(stage *Stage) *ControlPointShape {
	if _, ok := stage.ControlPointShapes[controlpointshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlPointShape(controlpointshape)
		}
	}
	return controlpointshape
}

func (controlpointshape *ControlPointShape) CommitVoid(stage *Stage) {
	controlpointshape.Commit(stage)
}

func (controlpointshape *ControlPointShape) StageVoid(stage *Stage) {
	controlpointshape.Stage(stage)
}

// Checkout controlpointshape to the back repo (if it is already staged)
func (controlpointshape *ControlPointShape) Checkout(stage *Stage) *ControlPointShape {
	if _, ok := stage.ControlPointShapes[controlpointshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlPointShape(controlpointshape)
		}
	}
	return controlpointshape
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) GetName() (res string) {
	return controlpointshape.Name
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) SetName(name string) (){
	controlpointshape.Name = name
}

// Stage puts desk to the model stage
func (desk *Desk) Stage(stage *Stage) *Desk {

	if _, ok := stage.Desks[desk]; !ok {
		stage.Desks[desk] = struct{}{}
		stage.DeskMap_Staged_Order[desk] = stage.DeskOrder
		stage.DeskOrder++
	}
	stage.Desks_mapString[desk.Name] = desk

	return desk
}

// StagePreserveOrder puts desk to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeskOrder
// - update stage.DeskOrder accordingly
func (desk *Desk) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Desks[desk]; !ok {
		stage.Desks[desk] = struct{}{}

		if order > stage.DeskOrder {
			stage.DeskOrder = order
		}
		stage.DeskMap_Staged_Order[desk] = stage.DeskOrder
		stage.DeskOrder++
	}
	stage.Desks_mapString[desk.Name] = desk
}

// Unstage removes desk off the model stage
func (desk *Desk) Unstage(stage *Stage) *Desk {
	delete(stage.Desks, desk)
	delete(stage.DeskMap_Staged_Order, desk)
	delete(stage.Desks_mapString, desk.Name)

	return desk
}

// UnstageVoid removes desk off the model stage
func (desk *Desk) UnstageVoid(stage *Stage) {
	delete(stage.Desks, desk)
	delete(stage.DeskMap_Staged_Order, desk)
	delete(stage.Desks_mapString, desk.Name)
}

// commit desk to the back repo (if it is already staged)
func (desk *Desk) Commit(stage *Stage) *Desk {
	if _, ok := stage.Desks[desk]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDesk(desk)
		}
	}
	return desk
}

func (desk *Desk) CommitVoid(stage *Stage) {
	desk.Commit(stage)
}

func (desk *Desk) StageVoid(stage *Stage) {
	desk.Stage(stage)
}

// Checkout desk to the back repo (if it is already staged)
func (desk *Desk) Checkout(stage *Stage) *Desk {
	if _, ok := stage.Desks[desk]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDesk(desk)
		}
	}
	return desk
}

// for satisfaction of GongStruct interface
func (desk *Desk) GetName() (res string) {
	return desk.Name
}

// for satisfaction of GongStruct interface
func (desk *Desk) SetName(name string) (){
	desk.Name = name
}

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

// Stage puts influence to the model stage
func (influence *Influence) Stage(stage *Stage) *Influence {

	if _, ok := stage.Influences[influence]; !ok {
		stage.Influences[influence] = struct{}{}
		stage.InfluenceMap_Staged_Order[influence] = stage.InfluenceOrder
		stage.InfluenceOrder++
	}
	stage.Influences_mapString[influence.Name] = influence

	return influence
}

// StagePreserveOrder puts influence to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InfluenceOrder
// - update stage.InfluenceOrder accordingly
func (influence *Influence) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Influences[influence]; !ok {
		stage.Influences[influence] = struct{}{}

		if order > stage.InfluenceOrder {
			stage.InfluenceOrder = order
		}
		stage.InfluenceMap_Staged_Order[influence] = stage.InfluenceOrder
		stage.InfluenceOrder++
	}
	stage.Influences_mapString[influence.Name] = influence
}

// Unstage removes influence off the model stage
func (influence *Influence) Unstage(stage *Stage) *Influence {
	delete(stage.Influences, influence)
	delete(stage.InfluenceMap_Staged_Order, influence)
	delete(stage.Influences_mapString, influence.Name)

	return influence
}

// UnstageVoid removes influence off the model stage
func (influence *Influence) UnstageVoid(stage *Stage) {
	delete(stage.Influences, influence)
	delete(stage.InfluenceMap_Staged_Order, influence)
	delete(stage.Influences_mapString, influence.Name)
}

// commit influence to the back repo (if it is already staged)
func (influence *Influence) Commit(stage *Stage) *Influence {
	if _, ok := stage.Influences[influence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitInfluence(influence)
		}
	}
	return influence
}

func (influence *Influence) CommitVoid(stage *Stage) {
	influence.Commit(stage)
}

func (influence *Influence) StageVoid(stage *Stage) {
	influence.Stage(stage)
}

// Checkout influence to the back repo (if it is already staged)
func (influence *Influence) Checkout(stage *Stage) *Influence {
	if _, ok := stage.Influences[influence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutInfluence(influence)
		}
	}
	return influence
}

// for satisfaction of GongStruct interface
func (influence *Influence) GetName() (res string) {
	return influence.Name
}

// for satisfaction of GongStruct interface
func (influence *Influence) SetName(name string) (){
	influence.Name = name
}

// Stage puts influenceshape to the model stage
func (influenceshape *InfluenceShape) Stage(stage *Stage) *InfluenceShape {

	if _, ok := stage.InfluenceShapes[influenceshape]; !ok {
		stage.InfluenceShapes[influenceshape] = struct{}{}
		stage.InfluenceShapeMap_Staged_Order[influenceshape] = stage.InfluenceShapeOrder
		stage.InfluenceShapeOrder++
	}
	stage.InfluenceShapes_mapString[influenceshape.Name] = influenceshape

	return influenceshape
}

// StagePreserveOrder puts influenceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InfluenceShapeOrder
// - update stage.InfluenceShapeOrder accordingly
func (influenceshape *InfluenceShape) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.InfluenceShapes[influenceshape]; !ok {
		stage.InfluenceShapes[influenceshape] = struct{}{}

		if order > stage.InfluenceShapeOrder {
			stage.InfluenceShapeOrder = order
		}
		stage.InfluenceShapeMap_Staged_Order[influenceshape] = stage.InfluenceShapeOrder
		stage.InfluenceShapeOrder++
	}
	stage.InfluenceShapes_mapString[influenceshape.Name] = influenceshape
}

// Unstage removes influenceshape off the model stage
func (influenceshape *InfluenceShape) Unstage(stage *Stage) *InfluenceShape {
	delete(stage.InfluenceShapes, influenceshape)
	delete(stage.InfluenceShapeMap_Staged_Order, influenceshape)
	delete(stage.InfluenceShapes_mapString, influenceshape.Name)

	return influenceshape
}

// UnstageVoid removes influenceshape off the model stage
func (influenceshape *InfluenceShape) UnstageVoid(stage *Stage) {
	delete(stage.InfluenceShapes, influenceshape)
	delete(stage.InfluenceShapeMap_Staged_Order, influenceshape)
	delete(stage.InfluenceShapes_mapString, influenceshape.Name)
}

// commit influenceshape to the back repo (if it is already staged)
func (influenceshape *InfluenceShape) Commit(stage *Stage) *InfluenceShape {
	if _, ok := stage.InfluenceShapes[influenceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitInfluenceShape(influenceshape)
		}
	}
	return influenceshape
}

func (influenceshape *InfluenceShape) CommitVoid(stage *Stage) {
	influenceshape.Commit(stage)
}

func (influenceshape *InfluenceShape) StageVoid(stage *Stage) {
	influenceshape.Stage(stage)
}

// Checkout influenceshape to the back repo (if it is already staged)
func (influenceshape *InfluenceShape) Checkout(stage *Stage) *InfluenceShape {
	if _, ok := stage.InfluenceShapes[influenceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutInfluenceShape(influenceshape)
		}
	}
	return influenceshape
}

// for satisfaction of GongStruct interface
func (influenceshape *InfluenceShape) GetName() (res string) {
	return influenceshape.Name
}

// for satisfaction of GongStruct interface
func (influenceshape *InfluenceShape) SetName(name string) (){
	influenceshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCategory1(Category1 *Category1)
	CreateORMCategory1Shape(Category1Shape *Category1Shape)
	CreateORMCategory2(Category2 *Category2)
	CreateORMCategory2Shape(Category2Shape *Category2Shape)
	CreateORMCategory3(Category3 *Category3)
	CreateORMCategory3Shape(Category3Shape *Category3Shape)
	CreateORMControlPointShape(ControlPointShape *ControlPointShape)
	CreateORMDesk(Desk *Desk)
	CreateORMDiagram(Diagram *Diagram)
	CreateORMInfluence(Influence *Influence)
	CreateORMInfluenceShape(InfluenceShape *InfluenceShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCategory1(Category1 *Category1)
	DeleteORMCategory1Shape(Category1Shape *Category1Shape)
	DeleteORMCategory2(Category2 *Category2)
	DeleteORMCategory2Shape(Category2Shape *Category2Shape)
	DeleteORMCategory3(Category3 *Category3)
	DeleteORMCategory3Shape(Category3Shape *Category3Shape)
	DeleteORMControlPointShape(ControlPointShape *ControlPointShape)
	DeleteORMDesk(Desk *Desk)
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMInfluence(Influence *Influence)
	DeleteORMInfluenceShape(InfluenceShape *InfluenceShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Category1s = make(map[*Category1]struct{})
	stage.Category1s_mapString = make(map[string]*Category1)
	stage.Category1Map_Staged_Order = make(map[*Category1]uint)
	stage.Category1Order = 0

	stage.Category1Shapes = make(map[*Category1Shape]struct{})
	stage.Category1Shapes_mapString = make(map[string]*Category1Shape)
	stage.Category1ShapeMap_Staged_Order = make(map[*Category1Shape]uint)
	stage.Category1ShapeOrder = 0

	stage.Category2s = make(map[*Category2]struct{})
	stage.Category2s_mapString = make(map[string]*Category2)
	stage.Category2Map_Staged_Order = make(map[*Category2]uint)
	stage.Category2Order = 0

	stage.Category2Shapes = make(map[*Category2Shape]struct{})
	stage.Category2Shapes_mapString = make(map[string]*Category2Shape)
	stage.Category2ShapeMap_Staged_Order = make(map[*Category2Shape]uint)
	stage.Category2ShapeOrder = 0

	stage.Category3s = make(map[*Category3]struct{})
	stage.Category3s_mapString = make(map[string]*Category3)
	stage.Category3Map_Staged_Order = make(map[*Category3]uint)
	stage.Category3Order = 0

	stage.Category3Shapes = make(map[*Category3Shape]struct{})
	stage.Category3Shapes_mapString = make(map[string]*Category3Shape)
	stage.Category3ShapeMap_Staged_Order = make(map[*Category3Shape]uint)
	stage.Category3ShapeOrder = 0

	stage.ControlPointShapes = make(map[*ControlPointShape]struct{})
	stage.ControlPointShapes_mapString = make(map[string]*ControlPointShape)
	stage.ControlPointShapeMap_Staged_Order = make(map[*ControlPointShape]uint)
	stage.ControlPointShapeOrder = 0

	stage.Desks = make(map[*Desk]struct{})
	stage.Desks_mapString = make(map[string]*Desk)
	stage.DeskMap_Staged_Order = make(map[*Desk]uint)
	stage.DeskOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.DiagramMap_Staged_Order = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Influences = make(map[*Influence]struct{})
	stage.Influences_mapString = make(map[string]*Influence)
	stage.InfluenceMap_Staged_Order = make(map[*Influence]uint)
	stage.InfluenceOrder = 0

	stage.InfluenceShapes = make(map[*InfluenceShape]struct{})
	stage.InfluenceShapes_mapString = make(map[string]*InfluenceShape)
	stage.InfluenceShapeMap_Staged_Order = make(map[*InfluenceShape]uint)
	stage.InfluenceShapeOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Category1s = nil
	stage.Category1s_mapString = nil

	stage.Category1Shapes = nil
	stage.Category1Shapes_mapString = nil

	stage.Category2s = nil
	stage.Category2s_mapString = nil

	stage.Category2Shapes = nil
	stage.Category2Shapes_mapString = nil

	stage.Category3s = nil
	stage.Category3s_mapString = nil

	stage.Category3Shapes = nil
	stage.Category3Shapes_mapString = nil

	stage.ControlPointShapes = nil
	stage.ControlPointShapes_mapString = nil

	stage.Desks = nil
	stage.Desks_mapString = nil

	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Influences = nil
	stage.Influences_mapString = nil

	stage.InfluenceShapes = nil
	stage.InfluenceShapes_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for category1 := range stage.Category1s {
		category1.Unstage(stage)
	}

	for category1shape := range stage.Category1Shapes {
		category1shape.Unstage(stage)
	}

	for category2 := range stage.Category2s {
		category2.Unstage(stage)
	}

	for category2shape := range stage.Category2Shapes {
		category2shape.Unstage(stage)
	}

	for category3 := range stage.Category3s {
		category3.Unstage(stage)
	}

	for category3shape := range stage.Category3Shapes {
		category3shape.Unstage(stage)
	}

	for controlpointshape := range stage.ControlPointShapes {
		controlpointshape.Unstage(stage)
	}

	for desk := range stage.Desks {
		desk.Unstage(stage)
	}

	for diagram := range stage.Diagrams {
		diagram.Unstage(stage)
	}

	for influence := range stage.Influences {
		influence.Unstage(stage)
	}

	for influenceshape := range stage.InfluenceShapes {
		influenceshape.Unstage(stage)
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
	case map[*Category1]any:
		return any(&stage.Category1s).(*Type)
	case map[*Category1Shape]any:
		return any(&stage.Category1Shapes).(*Type)
	case map[*Category2]any:
		return any(&stage.Category2s).(*Type)
	case map[*Category2Shape]any:
		return any(&stage.Category2Shapes).(*Type)
	case map[*Category3]any:
		return any(&stage.Category3s).(*Type)
	case map[*Category3Shape]any:
		return any(&stage.Category3Shapes).(*Type)
	case map[*ControlPointShape]any:
		return any(&stage.ControlPointShapes).(*Type)
	case map[*Desk]any:
		return any(&stage.Desks).(*Type)
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Influence]any:
		return any(&stage.Influences).(*Type)
	case map[*InfluenceShape]any:
		return any(&stage.InfluenceShapes).(*Type)
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
	case *Category1:
		return any(stage.Category1s_mapString).(map[string]Type)
	case *Category1Shape:
		return any(stage.Category1Shapes_mapString).(map[string]Type)
	case *Category2:
		return any(stage.Category2s_mapString).(map[string]Type)
	case *Category2Shape:
		return any(stage.Category2Shapes_mapString).(map[string]Type)
	case *Category3:
		return any(stage.Category3s_mapString).(map[string]Type)
	case *Category3Shape:
		return any(stage.Category3Shapes_mapString).(map[string]Type)
	case *ControlPointShape:
		return any(stage.ControlPointShapes_mapString).(map[string]Type)
	case *Desk:
		return any(stage.Desks_mapString).(map[string]Type)
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Influence:
		return any(stage.Influences_mapString).(map[string]Type)
	case *InfluenceShape:
		return any(stage.InfluenceShapes_mapString).(map[string]Type)
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
	case Category1:
		return any(&stage.Category1s).(*map[*Type]struct{})
	case Category1Shape:
		return any(&stage.Category1Shapes).(*map[*Type]struct{})
	case Category2:
		return any(&stage.Category2s).(*map[*Type]struct{})
	case Category2Shape:
		return any(&stage.Category2Shapes).(*map[*Type]struct{})
	case Category3:
		return any(&stage.Category3s).(*map[*Type]struct{})
	case Category3Shape:
		return any(&stage.Category3Shapes).(*map[*Type]struct{})
	case ControlPointShape:
		return any(&stage.ControlPointShapes).(*map[*Type]struct{})
	case Desk:
		return any(&stage.Desks).(*map[*Type]struct{})
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Influence:
		return any(&stage.Influences).(*map[*Type]struct{})
	case InfluenceShape:
		return any(&stage.InfluenceShapes).(*map[*Type]struct{})
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
	case *Category1:
		return any(&stage.Category1s).(*map[Type]struct{})
	case *Category1Shape:
		return any(&stage.Category1Shapes).(*map[Type]struct{})
	case *Category2:
		return any(&stage.Category2s).(*map[Type]struct{})
	case *Category2Shape:
		return any(&stage.Category2Shapes).(*map[Type]struct{})
	case *Category3:
		return any(&stage.Category3s).(*map[Type]struct{})
	case *Category3Shape:
		return any(&stage.Category3Shapes).(*map[Type]struct{})
	case *ControlPointShape:
		return any(&stage.ControlPointShapes).(*map[Type]struct{})
	case *Desk:
		return any(&stage.Desks).(*map[Type]struct{})
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Influence:
		return any(&stage.Influences).(*map[Type]struct{})
	case *InfluenceShape:
		return any(&stage.InfluenceShapes).(*map[Type]struct{})
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
	case Category1:
		return any(&stage.Category1s_mapString).(*map[string]*Type)
	case Category1Shape:
		return any(&stage.Category1Shapes_mapString).(*map[string]*Type)
	case Category2:
		return any(&stage.Category2s_mapString).(*map[string]*Type)
	case Category2Shape:
		return any(&stage.Category2Shapes_mapString).(*map[string]*Type)
	case Category3:
		return any(&stage.Category3s_mapString).(*map[string]*Type)
	case Category3Shape:
		return any(&stage.Category3Shapes_mapString).(*map[string]*Type)
	case ControlPointShape:
		return any(&stage.ControlPointShapes_mapString).(*map[string]*Type)
	case Desk:
		return any(&stage.Desks_mapString).(*map[string]*Type)
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Influence:
		return any(&stage.Influences_mapString).(*map[string]*Type)
	case InfluenceShape:
		return any(&stage.InfluenceShapes_mapString).(*map[string]*Type)
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
	case Category1:
		return any(&Category1{
			// Initialisation of associations
		}).(*Type)
	case Category1Shape:
		return any(&Category1Shape{
			// Initialisation of associations
			// field is initialized with an instance of Category1 with the name of the field
			Category1: &Category1{Name: "Category1"},
		}).(*Type)
	case Category2:
		return any(&Category2{
			// Initialisation of associations
		}).(*Type)
	case Category2Shape:
		return any(&Category2Shape{
			// Initialisation of associations
			// field is initialized with an instance of Category2 with the name of the field
			Category2: &Category2{Name: "Category2"},
		}).(*Type)
	case Category3:
		return any(&Category3{
			// Initialisation of associations
		}).(*Type)
	case Category3Shape:
		return any(&Category3Shape{
			// Initialisation of associations
			// field is initialized with an instance of Category3 with the name of the field
			Category3: &Category3{Name: "Category3"},
		}).(*Type)
	case ControlPointShape:
		return any(&ControlPointShape{
			// Initialisation of associations
		}).(*Type)
	case Desk:
		return any(&Desk{
			// Initialisation of associations
			// field is initialized with an instance of Diagram with the name of the field
			SelectedDiagram: &Diagram{Name: "SelectedDiagram"},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of Category1Shape with the name of the field
			Category1Shapes: []*Category1Shape{{Name: "Category1Shapes"}},
			// field is initialized with an instance of Category2Shape with the name of the field
			Category2Shapes: []*Category2Shape{{Name: "Category2Shapes"}},
			// field is initialized with an instance of Category3Shape with the name of the field
			Category3Shapes: []*Category3Shape{{Name: "Category3Shapes"}},
			// field is initialized with an instance of InfluenceShape with the name of the field
			InfluenceShapes: []*InfluenceShape{{Name: "InfluenceShapes"}},
		}).(*Type)
	case Influence:
		return any(&Influence{
			// Initialisation of associations
			// field is initialized with an instance of Category1 with the name of the field
			SourceCategory1: &Category1{Name: "SourceCategory1"},
			// field is initialized with an instance of Category3 with the name of the field
			SourceCategory2: &Category3{Name: "SourceCategory2"},
			// field is initialized with an instance of Category2 with the name of the field
			SourceCategory3: &Category2{Name: "SourceCategory3"},
			// field is initialized with an instance of Category1 with the name of the field
			TargetCategory1: &Category1{Name: "TargetCategory1"},
			// field is initialized with an instance of Category3 with the name of the field
			TargetCategory2: &Category3{Name: "TargetCategory2"},
			// field is initialized with an instance of Category2 with the name of the field
			TargetCategory3: &Category2{Name: "TargetCategory3"},
		}).(*Type)
	case InfluenceShape:
		return any(&InfluenceShape{
			// Initialisation of associations
			// field is initialized with an instance of Influence with the name of the field
			Influence: &Influence{Name: "Influence"},
			// field is initialized with an instance of ControlPointShape with the name of the field
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
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
	// reverse maps of direct associations of Category1
	case Category1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category1Shape
	case Category1Shape:
		switch fieldname {
		// insertion point for per direct association field
		case "Category1":
			res := make(map[*Category1][]*Category1Shape)
			for category1shape := range stage.Category1Shapes {
				if category1shape.Category1 != nil {
					category1_ := category1shape.Category1
					var category1shapes []*Category1Shape
					_, ok := res[category1_]
					if ok {
						category1shapes = res[category1_]
					} else {
						category1shapes = make([]*Category1Shape, 0)
					}
					category1shapes = append(category1shapes, category1shape)
					res[category1_] = category1shapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Category2
	case Category2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category2Shape
	case Category2Shape:
		switch fieldname {
		// insertion point for per direct association field
		case "Category2":
			res := make(map[*Category2][]*Category2Shape)
			for category2shape := range stage.Category2Shapes {
				if category2shape.Category2 != nil {
					category2_ := category2shape.Category2
					var category2shapes []*Category2Shape
					_, ok := res[category2_]
					if ok {
						category2shapes = res[category2_]
					} else {
						category2shapes = make([]*Category2Shape, 0)
					}
					category2shapes = append(category2shapes, category2shape)
					res[category2_] = category2shapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Category3
	case Category3:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category3Shape
	case Category3Shape:
		switch fieldname {
		// insertion point for per direct association field
		case "Category3":
			res := make(map[*Category3][]*Category3Shape)
			for category3shape := range stage.Category3Shapes {
				if category3shape.Category3 != nil {
					category3_ := category3shape.Category3
					var category3shapes []*Category3Shape
					_, ok := res[category3_]
					if ok {
						category3shapes = res[category3_]
					} else {
						category3shapes = make([]*Category3Shape, 0)
					}
					category3shapes = append(category3shapes, category3shape)
					res[category3_] = category3shapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Desk
	case Desk:
		switch fieldname {
		// insertion point for per direct association field
		case "SelectedDiagram":
			res := make(map[*Diagram][]*Desk)
			for desk := range stage.Desks {
				if desk.SelectedDiagram != nil {
					diagram_ := desk.SelectedDiagram
					var desks []*Desk
					_, ok := res[diagram_]
					if ok {
						desks = res[diagram_]
					} else {
						desks = make([]*Desk, 0)
					}
					desks = append(desks, desk)
					res[diagram_] = desks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Influence
	case Influence:
		switch fieldname {
		// insertion point for per direct association field
		case "SourceCategory1":
			res := make(map[*Category1][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceCategory1 != nil {
					category1_ := influence.SourceCategory1
					var influences []*Influence
					_, ok := res[category1_]
					if ok {
						influences = res[category1_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category1_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "SourceCategory2":
			res := make(map[*Category3][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceCategory2 != nil {
					category3_ := influence.SourceCategory2
					var influences []*Influence
					_, ok := res[category3_]
					if ok {
						influences = res[category3_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category3_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "SourceCategory3":
			res := make(map[*Category2][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceCategory3 != nil {
					category2_ := influence.SourceCategory3
					var influences []*Influence
					_, ok := res[category2_]
					if ok {
						influences = res[category2_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category2_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetCategory1":
			res := make(map[*Category1][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetCategory1 != nil {
					category1_ := influence.TargetCategory1
					var influences []*Influence
					_, ok := res[category1_]
					if ok {
						influences = res[category1_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category1_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetCategory2":
			res := make(map[*Category3][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetCategory2 != nil {
					category3_ := influence.TargetCategory2
					var influences []*Influence
					_, ok := res[category3_]
					if ok {
						influences = res[category3_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category3_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetCategory3":
			res := make(map[*Category2][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetCategory3 != nil {
					category2_ := influence.TargetCategory3
					var influences []*Influence
					_, ok := res[category2_]
					if ok {
						influences = res[category2_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[category2_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of InfluenceShape
	case InfluenceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Influence":
			res := make(map[*Influence][]*InfluenceShape)
			for influenceshape := range stage.InfluenceShapes {
				if influenceshape.Influence != nil {
					influence_ := influenceshape.Influence
					var influenceshapes []*InfluenceShape
					_, ok := res[influence_]
					if ok {
						influenceshapes = res[influence_]
					} else {
						influenceshapes = make([]*InfluenceShape, 0)
					}
					influenceshapes = append(influenceshapes, influenceshape)
					res[influence_] = influenceshapes
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
	// reverse maps of direct associations of Category1
	case Category1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category1Shape
	case Category1Shape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category2
	case Category2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category2Shape
	case Category2Shape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category3
	case Category3:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Category3Shape
	case Category3Shape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Desk
	case Desk:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "Category1Shapes":
			res := make(map[*Category1Shape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, category1shape_ := range diagram.Category1Shapes {
					res[category1shape_] = append(res[category1shape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Category2Shapes":
			res := make(map[*Category2Shape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, category2shape_ := range diagram.Category2Shapes {
					res[category2shape_] = append(res[category2shape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Category3Shapes":
			res := make(map[*Category3Shape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, category3shape_ := range diagram.Category3Shapes {
					res[category3shape_] = append(res[category3shape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "InfluenceShapes":
			res := make(map[*InfluenceShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, influenceshape_ := range diagram.InfluenceShapes {
					res[influenceshape_] = append(res[influenceshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Influence
	case Influence:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InfluenceShape
	case InfluenceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*InfluenceShape)
			for influenceshape := range stage.InfluenceShapes {
				for _, controlpointshape_ := range influenceshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], influenceshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Category1:
		res = "Category1"
	case *Category1Shape:
		res = "Category1Shape"
	case *Category2:
		res = "Category2"
	case *Category2Shape:
		res = "Category2Shape"
	case *Category3:
		res = "Category3"
	case *Category3Shape:
		res = "Category3Shape"
	case *ControlPointShape:
		res = "ControlPointShape"
	case *Desk:
		res = "Desk"
	case *Diagram:
		res = "Diagram"
	case *Influence:
		res = "Influence"
	case *InfluenceShape:
		res = "InfluenceShape"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Category1:
		var rf ReverseField
		_ = rf
	case *Category1Shape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Category1Shapes"
		res = append(res, rf)
	case *Category2:
		var rf ReverseField
		_ = rf
	case *Category2Shape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Category2Shapes"
		res = append(res, rf)
	case *Category3:
		var rf ReverseField
		_ = rf
	case *Category3Shape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Category3Shapes"
		res = append(res, rf)
	case *ControlPointShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "InfluenceShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
	case *Desk:
		var rf ReverseField
		_ = rf
	case *Diagram:
		var rf ReverseField
		_ = rf
	case *Influence:
		var rf ReverseField
		_ = rf
	case *InfluenceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "InfluenceShapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (category1 *Category1) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (category1shape *Category1Shape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Category1",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category1",
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

func (category2 *Category2) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (category2shape *Category2Shape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Category2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category2",
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

func (category3 *Category3) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (category3shape *Category3Shape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Category3",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category3",
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

func (controlpointshape *ControlPointShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsStartShapeTheClosestShape",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (desk *Desk) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "SelectedDiagram",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Diagram",
		},
	}
	return
}

func (diagram *Diagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Category1Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Category1Shape",
		},
		{
			Name:                 "Category2Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Category2Shape",
		},
		{
			Name:                 "Category3Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Category3Shape",
		},
		{
			Name:                 "InfluenceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "InfluenceShape",
		},
		{
			Name:               "IsEditable",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory1NodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory2NodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory3NodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsInfluenceCategoryNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory1Shown",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory2Shown",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsCategory3Shown",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsInfluenceCategoryShown",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "XMargin",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "YMargin",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RedColorCode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BackgroundGreyColorCode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "GrayColorCode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1RectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1TextAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1DominantBaselineType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1FontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1FontWeigth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1FontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category1LetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2TypeFontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2TypeFontWeigth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2TypeFontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2TypeLetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2TypeRectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2DominantBaselineType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category2StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3RectAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3TextAnchorType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3DominantBaselineType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3FontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3FontWeigth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3FontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Category3LetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceStrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceArrowSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceArrowStartOffset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceArrowEndOffset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceCornerRadius",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceFontSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceFontWeigth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceFontFamily",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceLetterSpacing",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "InfluenceDashedLinePattern",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (influence *Influence) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "SourceCategory1",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category1",
		},
		{
			Name:                 "SourceCategory2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category3",
		},
		{
			Name:                 "SourceCategory3",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category2",
		},
		{
			Name:                 "TargetCategory1",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category1",
		},
		{
			Name:                 "TargetCategory2",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category3",
		},
		{
			Name:                 "TargetCategory3",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Category2",
		},
		{
			Name:               "IsHypothtical",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TextAtEndOfArrow",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (influenceshape *InfluenceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Influence",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Influence",
		},
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
func (category1 *Category1) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category1.Name
	}
	return
}
func (category1shape *Category1Shape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category1shape.Name
	case "Category1":
		res.GongFieldValueType = GongFieldValueTypePointer
		if category1shape.Category1 != nil {
			res.valueString = category1shape.Category1.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, category1shape.Category1))
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", category1shape.X)
		res.valueFloat = category1shape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", category1shape.Y)
		res.valueFloat = category1shape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", category1shape.Width)
		res.valueFloat = category1shape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", category1shape.Height)
		res.valueFloat = category1shape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (category2 *Category2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category2.Name
	}
	return
}
func (category2shape *Category2Shape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category2shape.Name
	case "Category2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if category2shape.Category2 != nil {
			res.valueString = category2shape.Category2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, category2shape.Category2))
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", category2shape.X)
		res.valueFloat = category2shape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", category2shape.Y)
		res.valueFloat = category2shape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", category2shape.Width)
		res.valueFloat = category2shape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", category2shape.Height)
		res.valueFloat = category2shape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (category3 *Category3) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category3.Name
	}
	return
}
func (category3shape *Category3Shape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = category3shape.Name
	case "Category3":
		res.GongFieldValueType = GongFieldValueTypePointer
		if category3shape.Category3 != nil {
			res.valueString = category3shape.Category3.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, category3shape.Category3))
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", category3shape.X)
		res.valueFloat = category3shape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", category3shape.Y)
		res.valueFloat = category3shape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", category3shape.Width)
		res.valueFloat = category3shape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", category3shape.Height)
		res.valueFloat = category3shape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func (controlpointshape *ControlPointShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlpointshape.Name
	case "X_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.X_Relative)
		res.valueFloat = controlpointshape.X_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.Y_Relative)
		res.valueFloat = controlpointshape.Y_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsStartShapeTheClosestShape":
		res.valueString = fmt.Sprintf("%t", controlpointshape.IsStartShapeTheClosestShape)
		res.valueBool = controlpointshape.IsStartShapeTheClosestShape
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (desk *Desk) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = desk.Name
	case "SelectedDiagram":
		res.GongFieldValueType = GongFieldValueTypePointer
		if desk.SelectedDiagram != nil {
			res.valueString = desk.SelectedDiagram.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, desk.SelectedDiagram))
		}
	}
	return
}
func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "Category1Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Category1Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Category2Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Category2Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Category3Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Category3Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "InfluenceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.InfluenceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsEditable":
		res.valueString = fmt.Sprintf("%t", diagram.IsEditable)
		res.valueBool = diagram.IsEditable
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsNodeExpanded)
		res.valueBool = diagram.IsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory1NodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory1NodeExpanded)
		res.valueBool = diagram.IsCategory1NodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory2NodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory2NodeExpanded)
		res.valueBool = diagram.IsCategory2NodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory3NodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory3NodeExpanded)
		res.valueBool = diagram.IsCategory3NodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInfluenceCategoryNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded)
		res.valueBool = diagram.IsInfluenceCategoryNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory1Shown":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory1Shown)
		res.valueBool = diagram.IsCategory1Shown
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory2Shown":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory2Shown)
		res.valueBool = diagram.IsCategory2Shown
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCategory3Shown":
		res.valueString = fmt.Sprintf("%t", diagram.IsCategory3Shown)
		res.valueBool = diagram.IsCategory3Shown
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInfluenceCategoryShown":
		res.valueString = fmt.Sprintf("%t", diagram.IsInfluenceCategoryShown)
		res.valueBool = diagram.IsInfluenceCategoryShown
		res.GongFieldValueType = GongFieldValueTypeBool
	case "XMargin":
		res.valueString = fmt.Sprintf("%f", diagram.XMargin)
		res.valueFloat = diagram.XMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "YMargin":
		res.valueString = fmt.Sprintf("%f", diagram.YMargin)
		res.valueFloat = diagram.YMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagram.Height)
		res.valueFloat = diagram.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagram.Width)
		res.valueFloat = diagram.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", diagram.NbPixPerCharacter)
		res.valueFloat = diagram.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RedColorCode":
		res.valueString = diagram.RedColorCode
	case "BackgroundGreyColorCode":
		res.valueString = diagram.BackgroundGreyColorCode
	case "GrayColorCode":
		res.valueString = diagram.GrayColorCode
	case "Category1RectAnchorType":
		enum := diagram.Category1RectAnchorType
		res.valueString = enum.ToCodeString()
	case "Category1TextAnchorType":
		enum := diagram.Category1TextAnchorType
		res.valueString = enum.ToCodeString()
	case "Category1DominantBaselineType":
		enum := diagram.Category1DominantBaselineType
		res.valueString = enum.ToCodeString()
	case "Category1FontSize":
		res.valueString = diagram.Category1FontSize
	case "Category1FontWeigth":
		res.valueString = diagram.Category1FontWeigth
	case "Category1FontFamily":
		res.valueString = diagram.Category1FontFamily
	case "Category1LetterSpacing":
		res.valueString = diagram.Category1LetterSpacing
	case "Category2TypeFontSize":
		res.valueString = diagram.Category2TypeFontSize
	case "Category2TypeFontWeigth":
		res.valueString = diagram.Category2TypeFontWeigth
	case "Category2TypeFontFamily":
		res.valueString = diagram.Category2TypeFontFamily
	case "Category2TypeLetterSpacing":
		res.valueString = diagram.Category2TypeLetterSpacing
	case "Category2TypeRectAnchorType":
		enum := diagram.Category2TypeRectAnchorType
		res.valueString = enum.ToCodeString()
	case "Category2DominantBaselineType":
		enum := diagram.Category2DominantBaselineType
		res.valueString = enum.ToCodeString()
	case "Category2StrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.Category2StrokeWidth)
		res.valueFloat = diagram.Category2StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Category3RectAnchorType":
		enum := diagram.Category3RectAnchorType
		res.valueString = enum.ToCodeString()
	case "Category3TextAnchorType":
		enum := diagram.Category3TextAnchorType
		res.valueString = enum.ToCodeString()
	case "Category3DominantBaselineType":
		enum := diagram.Category3DominantBaselineType
		res.valueString = enum.ToCodeString()
	case "Category3FontSize":
		res.valueString = diagram.Category3FontSize
	case "Category3FontWeigth":
		res.valueString = diagram.Category3FontWeigth
	case "Category3FontFamily":
		res.valueString = diagram.Category3FontFamily
	case "Category3LetterSpacing":
		res.valueString = diagram.Category3LetterSpacing
	case "InfluenceStrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.InfluenceStrokeWidth)
		res.valueFloat = diagram.InfluenceStrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InfluenceArrowSize":
		res.valueString = fmt.Sprintf("%f", diagram.InfluenceArrowSize)
		res.valueFloat = diagram.InfluenceArrowSize
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InfluenceArrowStartOffset":
		res.valueString = fmt.Sprintf("%f", diagram.InfluenceArrowStartOffset)
		res.valueFloat = diagram.InfluenceArrowStartOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InfluenceArrowEndOffset":
		res.valueString = fmt.Sprintf("%f", diagram.InfluenceArrowEndOffset)
		res.valueFloat = diagram.InfluenceArrowEndOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InfluenceCornerRadius":
		res.valueString = fmt.Sprintf("%f", diagram.InfluenceCornerRadius)
		res.valueFloat = diagram.InfluenceCornerRadius
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "InfluenceFontSize":
		res.valueString = diagram.InfluenceFontSize
	case "InfluenceFontWeigth":
		res.valueString = diagram.InfluenceFontWeigth
	case "InfluenceFontFamily":
		res.valueString = diagram.InfluenceFontFamily
	case "InfluenceLetterSpacing":
		res.valueString = diagram.InfluenceLetterSpacing
	case "InfluenceDashedLinePattern":
		res.valueString = diagram.InfluenceDashedLinePattern
	}
	return
}
func (influence *Influence) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = influence.Name
	case "SourceCategory1":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceCategory1 != nil {
			res.valueString = influence.SourceCategory1.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.SourceCategory1))
		}
	case "SourceCategory2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceCategory2 != nil {
			res.valueString = influence.SourceCategory2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.SourceCategory2))
		}
	case "SourceCategory3":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceCategory3 != nil {
			res.valueString = influence.SourceCategory3.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.SourceCategory3))
		}
	case "TargetCategory1":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetCategory1 != nil {
			res.valueString = influence.TargetCategory1.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.TargetCategory1))
		}
	case "TargetCategory2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetCategory2 != nil {
			res.valueString = influence.TargetCategory2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.TargetCategory2))
		}
	case "TargetCategory3":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetCategory3 != nil {
			res.valueString = influence.TargetCategory3.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influence.TargetCategory3))
		}
	case "IsHypothtical":
		res.valueString = fmt.Sprintf("%t", influence.IsHypothtical)
		res.valueBool = influence.IsHypothtical
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TextAtEndOfArrow":
		res.valueString = influence.TextAtEndOfArrow
	}
	return
}
func (influenceshape *InfluenceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = influenceshape.Name
	case "Influence":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influenceshape.Influence != nil {
			res.valueString = influenceshape.Influence.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, influenceshape.Influence))
		}
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range influenceshape.ControlPointShapes {
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
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (category1 *Category1) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category1.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (category1shape *Category1Shape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category1shape.Name = value.GetValueString()
	case "Category1":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			category1shape.Category1 = nil
			for __instance__ := range stage.Category1s {
				if stage.Category1Map_Staged_Order[__instance__] == uint(id) {
					category1shape.Category1 = __instance__
					break
				}
			}
		}
	case "X":
		category1shape.X = value.GetValueFloat()
	case "Y":
		category1shape.Y = value.GetValueFloat()
	case "Width":
		category1shape.Width = value.GetValueFloat()
	case "Height":
		category1shape.Height = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (category2 *Category2) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category2.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (category2shape *Category2Shape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category2shape.Name = value.GetValueString()
	case "Category2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			category2shape.Category2 = nil
			for __instance__ := range stage.Category2s {
				if stage.Category2Map_Staged_Order[__instance__] == uint(id) {
					category2shape.Category2 = __instance__
					break
				}
			}
		}
	case "X":
		category2shape.X = value.GetValueFloat()
	case "Y":
		category2shape.Y = value.GetValueFloat()
	case "Width":
		category2shape.Width = value.GetValueFloat()
	case "Height":
		category2shape.Height = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (category3 *Category3) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category3.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (category3shape *Category3Shape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		category3shape.Name = value.GetValueString()
	case "Category3":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			category3shape.Category3 = nil
			for __instance__ := range stage.Category3s {
				if stage.Category3Map_Staged_Order[__instance__] == uint(id) {
					category3shape.Category3 = __instance__
					break
				}
			}
		}
	case "X":
		category3shape.X = value.GetValueFloat()
	case "Y":
		category3shape.Y = value.GetValueFloat()
	case "Width":
		category3shape.Width = value.GetValueFloat()
	case "Height":
		category3shape.Height = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlpointshape *ControlPointShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlpointshape.Name = value.GetValueString()
	case "X_Relative":
		controlpointshape.X_Relative = value.GetValueFloat()
	case "Y_Relative":
		controlpointshape.Y_Relative = value.GetValueFloat()
	case "IsStartShapeTheClosestShape":
		controlpointshape.IsStartShapeTheClosestShape = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (desk *Desk) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		desk.Name = value.GetValueString()
	case "SelectedDiagram":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			desk.SelectedDiagram = nil
			for __instance__ := range stage.Diagrams {
				if stage.DiagramMap_Staged_Order[__instance__] == uint(id) {
					desk.SelectedDiagram = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagram *Diagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagram.Name = value.GetValueString()
	case "Category1Shapes":
		diagram.Category1Shapes = make([]*Category1Shape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Category1Shapes {
					if stage.Category1ShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Category1Shapes = append(diagram.Category1Shapes, __instance__)
						break
					}
				}
			}
		}
	case "Category2Shapes":
		diagram.Category2Shapes = make([]*Category2Shape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Category2Shapes {
					if stage.Category2ShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Category2Shapes = append(diagram.Category2Shapes, __instance__)
						break
					}
				}
			}
		}
	case "Category3Shapes":
		diagram.Category3Shapes = make([]*Category3Shape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Category3Shapes {
					if stage.Category3ShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Category3Shapes = append(diagram.Category3Shapes, __instance__)
						break
					}
				}
			}
		}
	case "InfluenceShapes":
		diagram.InfluenceShapes = make([]*InfluenceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.InfluenceShapes {
					if stage.InfluenceShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.InfluenceShapes = append(diagram.InfluenceShapes, __instance__)
						break
					}
				}
			}
		}
	case "IsEditable":
		diagram.IsEditable = value.GetValueBool()
	case "IsNodeExpanded":
		diagram.IsNodeExpanded = value.GetValueBool()
	case "IsCategory1NodeExpanded":
		diagram.IsCategory1NodeExpanded = value.GetValueBool()
	case "IsCategory2NodeExpanded":
		diagram.IsCategory2NodeExpanded = value.GetValueBool()
	case "IsCategory3NodeExpanded":
		diagram.IsCategory3NodeExpanded = value.GetValueBool()
	case "IsInfluenceCategoryNodeExpanded":
		diagram.IsInfluenceCategoryNodeExpanded = value.GetValueBool()
	case "IsCategory1Shown":
		diagram.IsCategory1Shown = value.GetValueBool()
	case "IsCategory2Shown":
		diagram.IsCategory2Shown = value.GetValueBool()
	case "IsCategory3Shown":
		diagram.IsCategory3Shown = value.GetValueBool()
	case "IsInfluenceCategoryShown":
		diagram.IsInfluenceCategoryShown = value.GetValueBool()
	case "XMargin":
		diagram.XMargin = value.GetValueFloat()
	case "YMargin":
		diagram.YMargin = value.GetValueFloat()
	case "Height":
		diagram.Height = value.GetValueFloat()
	case "Width":
		diagram.Width = value.GetValueFloat()
	case "NbPixPerCharacter":
		diagram.NbPixPerCharacter = value.GetValueFloat()
	case "RedColorCode":
		diagram.RedColorCode = value.GetValueString()
	case "BackgroundGreyColorCode":
		diagram.BackgroundGreyColorCode = value.GetValueString()
	case "GrayColorCode":
		diagram.GrayColorCode = value.GetValueString()
	case "Category1RectAnchorType":
		diagram.Category1RectAnchorType.FromCodeString(value.GetValueString())
	case "Category1TextAnchorType":
		diagram.Category1TextAnchorType.FromCodeString(value.GetValueString())
	case "Category1DominantBaselineType":
		diagram.Category1DominantBaselineType.FromCodeString(value.GetValueString())
	case "Category1FontSize":
		diagram.Category1FontSize = value.GetValueString()
	case "Category1FontWeigth":
		diagram.Category1FontWeigth = value.GetValueString()
	case "Category1FontFamily":
		diagram.Category1FontFamily = value.GetValueString()
	case "Category1LetterSpacing":
		diagram.Category1LetterSpacing = value.GetValueString()
	case "Category2TypeFontSize":
		diagram.Category2TypeFontSize = value.GetValueString()
	case "Category2TypeFontWeigth":
		diagram.Category2TypeFontWeigth = value.GetValueString()
	case "Category2TypeFontFamily":
		diagram.Category2TypeFontFamily = value.GetValueString()
	case "Category2TypeLetterSpacing":
		diagram.Category2TypeLetterSpacing = value.GetValueString()
	case "Category2TypeRectAnchorType":
		diagram.Category2TypeRectAnchorType.FromCodeString(value.GetValueString())
	case "Category2DominantBaselineType":
		diagram.Category2DominantBaselineType.FromCodeString(value.GetValueString())
	case "Category2StrokeWidth":
		diagram.Category2StrokeWidth = value.GetValueFloat()
	case "Category3RectAnchorType":
		diagram.Category3RectAnchorType.FromCodeString(value.GetValueString())
	case "Category3TextAnchorType":
		diagram.Category3TextAnchorType.FromCodeString(value.GetValueString())
	case "Category3DominantBaselineType":
		diagram.Category3DominantBaselineType.FromCodeString(value.GetValueString())
	case "Category3FontSize":
		diagram.Category3FontSize = value.GetValueString()
	case "Category3FontWeigth":
		diagram.Category3FontWeigth = value.GetValueString()
	case "Category3FontFamily":
		diagram.Category3FontFamily = value.GetValueString()
	case "Category3LetterSpacing":
		diagram.Category3LetterSpacing = value.GetValueString()
	case "InfluenceStrokeWidth":
		diagram.InfluenceStrokeWidth = value.GetValueFloat()
	case "InfluenceArrowSize":
		diagram.InfluenceArrowSize = value.GetValueFloat()
	case "InfluenceArrowStartOffset":
		diagram.InfluenceArrowStartOffset = value.GetValueFloat()
	case "InfluenceArrowEndOffset":
		diagram.InfluenceArrowEndOffset = value.GetValueFloat()
	case "InfluenceCornerRadius":
		diagram.InfluenceCornerRadius = value.GetValueFloat()
	case "InfluenceFontSize":
		diagram.InfluenceFontSize = value.GetValueString()
	case "InfluenceFontWeigth":
		diagram.InfluenceFontWeigth = value.GetValueString()
	case "InfluenceFontFamily":
		diagram.InfluenceFontFamily = value.GetValueString()
	case "InfluenceLetterSpacing":
		diagram.InfluenceLetterSpacing = value.GetValueString()
	case "InfluenceDashedLinePattern":
		diagram.InfluenceDashedLinePattern = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (influence *Influence) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		influence.Name = value.GetValueString()
	case "SourceCategory1":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceCategory1 = nil
			for __instance__ := range stage.Category1s {
				if stage.Category1Map_Staged_Order[__instance__] == uint(id) {
					influence.SourceCategory1 = __instance__
					break
				}
			}
		}
	case "SourceCategory2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceCategory2 = nil
			for __instance__ := range stage.Category3s {
				if stage.Category3Map_Staged_Order[__instance__] == uint(id) {
					influence.SourceCategory2 = __instance__
					break
				}
			}
		}
	case "SourceCategory3":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceCategory3 = nil
			for __instance__ := range stage.Category2s {
				if stage.Category2Map_Staged_Order[__instance__] == uint(id) {
					influence.SourceCategory3 = __instance__
					break
				}
			}
		}
	case "TargetCategory1":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetCategory1 = nil
			for __instance__ := range stage.Category1s {
				if stage.Category1Map_Staged_Order[__instance__] == uint(id) {
					influence.TargetCategory1 = __instance__
					break
				}
			}
		}
	case "TargetCategory2":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetCategory2 = nil
			for __instance__ := range stage.Category3s {
				if stage.Category3Map_Staged_Order[__instance__] == uint(id) {
					influence.TargetCategory2 = __instance__
					break
				}
			}
		}
	case "TargetCategory3":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetCategory3 = nil
			for __instance__ := range stage.Category2s {
				if stage.Category2Map_Staged_Order[__instance__] == uint(id) {
					influence.TargetCategory3 = __instance__
					break
				}
			}
		}
	case "IsHypothtical":
		influence.IsHypothtical = value.GetValueBool()
	case "TextAtEndOfArrow":
		influence.TextAtEndOfArrow = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (influenceshape *InfluenceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		influenceshape.Name = value.GetValueString()
	case "Influence":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influenceshape.Influence = nil
			for __instance__ := range stage.Influences {
				if stage.InfluenceMap_Staged_Order[__instance__] == uint(id) {
					influenceshape.Influence = __instance__
					break
				}
			}
		}
	case "ControlPointShapes":
		influenceshape.ControlPointShapes = make([]*ControlPointShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlPointShapes {
					if stage.ControlPointShapeMap_Staged_Order[__instance__] == uint(id) {
						influenceshape.ControlPointShapes = append(influenceshape.ControlPointShapes, __instance__)
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

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (category1 *Category1) GongGetGongstructName() string {
	return "Category1"
}

func (category1shape *Category1Shape) GongGetGongstructName() string {
	return "Category1Shape"
}

func (category2 *Category2) GongGetGongstructName() string {
	return "Category2"
}

func (category2shape *Category2Shape) GongGetGongstructName() string {
	return "Category2Shape"
}

func (category3 *Category3) GongGetGongstructName() string {
	return "Category3"
}

func (category3shape *Category3Shape) GongGetGongstructName() string {
	return "Category3Shape"
}

func (controlpointshape *ControlPointShape) GongGetGongstructName() string {
	return "ControlPointShape"
}

func (desk *Desk) GongGetGongstructName() string {
	return "Desk"
}

func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (influence *Influence) GongGetGongstructName() string {
	return "Influence"
}

func (influenceshape *InfluenceShape) GongGetGongstructName() string {
	return "InfluenceShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Category1s_mapString = make(map[string]*Category1)
	for category1 := range stage.Category1s {
		stage.Category1s_mapString[category1.Name] = category1
	}

	stage.Category1Shapes_mapString = make(map[string]*Category1Shape)
	for category1shape := range stage.Category1Shapes {
		stage.Category1Shapes_mapString[category1shape.Name] = category1shape
	}

	stage.Category2s_mapString = make(map[string]*Category2)
	for category2 := range stage.Category2s {
		stage.Category2s_mapString[category2.Name] = category2
	}

	stage.Category2Shapes_mapString = make(map[string]*Category2Shape)
	for category2shape := range stage.Category2Shapes {
		stage.Category2Shapes_mapString[category2shape.Name] = category2shape
	}

	stage.Category3s_mapString = make(map[string]*Category3)
	for category3 := range stage.Category3s {
		stage.Category3s_mapString[category3.Name] = category3
	}

	stage.Category3Shapes_mapString = make(map[string]*Category3Shape)
	for category3shape := range stage.Category3Shapes {
		stage.Category3Shapes_mapString[category3shape.Name] = category3shape
	}

	stage.ControlPointShapes_mapString = make(map[string]*ControlPointShape)
	for controlpointshape := range stage.ControlPointShapes {
		stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape
	}

	stage.Desks_mapString = make(map[string]*Desk)
	for desk := range stage.Desks {
		stage.Desks_mapString[desk.Name] = desk
	}

	stage.Diagrams_mapString = make(map[string]*Diagram)
	for diagram := range stage.Diagrams {
		stage.Diagrams_mapString[diagram.Name] = diagram
	}

	stage.Influences_mapString = make(map[string]*Influence)
	for influence := range stage.Influences {
		stage.Influences_mapString[influence.Name] = influence
	}

	stage.InfluenceShapes_mapString = make(map[string]*InfluenceShape)
	for influenceshape := range stage.InfluenceShapes {
		stage.InfluenceShapes_mapString[influenceshape.Name] = influenceshape
	}

}
// Last line of the template
