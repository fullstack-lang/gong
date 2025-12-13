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

	doc_go "github.com/fullstack-lang/gong/lib/doc/go"
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
	AttributeShapes           map[*AttributeShape]struct{}
	AttributeShapes_mapString map[string]*AttributeShape

	// insertion point for slice of pointers maps
	OnAfterAttributeShapeCreateCallback OnAfterCreateInterface[AttributeShape]
	OnAfterAttributeShapeUpdateCallback OnAfterUpdateInterface[AttributeShape]
	OnAfterAttributeShapeDeleteCallback OnAfterDeleteInterface[AttributeShape]
	OnAfterAttributeShapeReadCallback   OnAfterReadInterface[AttributeShape]

	Classdiagrams           map[*Classdiagram]struct{}
	Classdiagrams_mapString map[string]*Classdiagram

	// insertion point for slice of pointers maps
	Classdiagram_GongStructShapes_reverseMap map[*GongStructShape]*Classdiagram

	Classdiagram_GongEnumShapes_reverseMap map[*GongEnumShape]*Classdiagram

	Classdiagram_GongNoteShapes_reverseMap map[*GongNoteShape]*Classdiagram

	OnAfterClassdiagramCreateCallback OnAfterCreateInterface[Classdiagram]
	OnAfterClassdiagramUpdateCallback OnAfterUpdateInterface[Classdiagram]
	OnAfterClassdiagramDeleteCallback OnAfterDeleteInterface[Classdiagram]
	OnAfterClassdiagramReadCallback   OnAfterReadInterface[Classdiagram]

	DiagramPackages           map[*DiagramPackage]struct{}
	DiagramPackages_mapString map[string]*DiagramPackage

	// insertion point for slice of pointers maps
	DiagramPackage_Classdiagrams_reverseMap map[*Classdiagram]*DiagramPackage

	OnAfterDiagramPackageCreateCallback OnAfterCreateInterface[DiagramPackage]
	OnAfterDiagramPackageUpdateCallback OnAfterUpdateInterface[DiagramPackage]
	OnAfterDiagramPackageDeleteCallback OnAfterDeleteInterface[DiagramPackage]
	OnAfterDiagramPackageReadCallback   OnAfterReadInterface[DiagramPackage]

	GongEnumShapes           map[*GongEnumShape]struct{}
	GongEnumShapes_mapString map[string]*GongEnumShape

	// insertion point for slice of pointers maps
	GongEnumShape_GongEnumValueShapes_reverseMap map[*GongEnumValueShape]*GongEnumShape

	OnAfterGongEnumShapeCreateCallback OnAfterCreateInterface[GongEnumShape]
	OnAfterGongEnumShapeUpdateCallback OnAfterUpdateInterface[GongEnumShape]
	OnAfterGongEnumShapeDeleteCallback OnAfterDeleteInterface[GongEnumShape]
	OnAfterGongEnumShapeReadCallback   OnAfterReadInterface[GongEnumShape]

	GongEnumValueShapes           map[*GongEnumValueShape]struct{}
	GongEnumValueShapes_mapString map[string]*GongEnumValueShape

	// insertion point for slice of pointers maps
	OnAfterGongEnumValueShapeCreateCallback OnAfterCreateInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeUpdateCallback OnAfterUpdateInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeDeleteCallback OnAfterDeleteInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeReadCallback   OnAfterReadInterface[GongEnumValueShape]

	GongNoteLinkShapes           map[*GongNoteLinkShape]struct{}
	GongNoteLinkShapes_mapString map[string]*GongNoteLinkShape

	// insertion point for slice of pointers maps
	OnAfterGongNoteLinkShapeCreateCallback OnAfterCreateInterface[GongNoteLinkShape]
	OnAfterGongNoteLinkShapeUpdateCallback OnAfterUpdateInterface[GongNoteLinkShape]
	OnAfterGongNoteLinkShapeDeleteCallback OnAfterDeleteInterface[GongNoteLinkShape]
	OnAfterGongNoteLinkShapeReadCallback   OnAfterReadInterface[GongNoteLinkShape]

	GongNoteShapes           map[*GongNoteShape]struct{}
	GongNoteShapes_mapString map[string]*GongNoteShape

	// insertion point for slice of pointers maps
	GongNoteShape_GongNoteLinkShapes_reverseMap map[*GongNoteLinkShape]*GongNoteShape

	OnAfterGongNoteShapeCreateCallback OnAfterCreateInterface[GongNoteShape]
	OnAfterGongNoteShapeUpdateCallback OnAfterUpdateInterface[GongNoteShape]
	OnAfterGongNoteShapeDeleteCallback OnAfterDeleteInterface[GongNoteShape]
	OnAfterGongNoteShapeReadCallback   OnAfterReadInterface[GongNoteShape]

	GongStructShapes           map[*GongStructShape]struct{}
	GongStructShapes_mapString map[string]*GongStructShape

	// insertion point for slice of pointers maps
	GongStructShape_AttributeShapes_reverseMap map[*AttributeShape]*GongStructShape

	GongStructShape_LinkShapes_reverseMap map[*LinkShape]*GongStructShape

	OnAfterGongStructShapeCreateCallback OnAfterCreateInterface[GongStructShape]
	OnAfterGongStructShapeUpdateCallback OnAfterUpdateInterface[GongStructShape]
	OnAfterGongStructShapeDeleteCallback OnAfterDeleteInterface[GongStructShape]
	OnAfterGongStructShapeReadCallback   OnAfterReadInterface[GongStructShape]

	LinkShapes           map[*LinkShape]struct{}
	LinkShapes_mapString map[string]*LinkShape

	// insertion point for slice of pointers maps
	OnAfterLinkShapeCreateCallback OnAfterCreateInterface[LinkShape]
	OnAfterLinkShapeUpdateCallback OnAfterUpdateInterface[LinkShape]
	OnAfterLinkShapeDeleteCallback OnAfterDeleteInterface[LinkShape]
	OnAfterLinkShapeReadCallback   OnAfterReadInterface[LinkShape]

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
	AttributeShapeOrder            uint
	AttributeShapeMap_Staged_Order map[*AttributeShape]uint

	ClassdiagramOrder            uint
	ClassdiagramMap_Staged_Order map[*Classdiagram]uint

	DiagramPackageOrder            uint
	DiagramPackageMap_Staged_Order map[*DiagramPackage]uint

	GongEnumShapeOrder            uint
	GongEnumShapeMap_Staged_Order map[*GongEnumShape]uint

	GongEnumValueShapeOrder            uint
	GongEnumValueShapeMap_Staged_Order map[*GongEnumValueShape]uint

	GongNoteLinkShapeOrder            uint
	GongNoteLinkShapeMap_Staged_Order map[*GongNoteLinkShape]uint

	GongNoteShapeOrder            uint
	GongNoteShapeMap_Staged_Order map[*GongNoteShape]uint

	GongStructShapeOrder            uint
	GongStructShapeMap_Staged_Order map[*GongStructShape]uint

	LinkShapeOrder            uint
	LinkShapeMap_Staged_Order map[*LinkShape]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
	modified  map[GongstructIF]struct{}
	new       map[GongstructIF]struct{}
	deleted   map[GongstructIF]struct{}
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

func (stage *Stage) GetModified() map[GongstructIF]struct{} {
	return stage.modified
}

func (stage *Stage) GetNew() map[GongstructIF]struct{} {
	return stage.new
}

func (stage *Stage) GetDeleted() map[GongstructIF]struct{} {
	return stage.deleted
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
	case *AttributeShape:
		tmp := GetStructInstancesByOrder(stage.AttributeShapes, stage.AttributeShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AttributeShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Classdiagram:
		tmp := GetStructInstancesByOrder(stage.Classdiagrams, stage.ClassdiagramMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Classdiagram implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DiagramPackage:
		tmp := GetStructInstancesByOrder(stage.DiagramPackages, stage.DiagramPackageMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramPackage implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongEnumShape:
		tmp := GetStructInstancesByOrder(stage.GongEnumShapes, stage.GongEnumShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongEnumShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongEnumValueShape:
		tmp := GetStructInstancesByOrder(stage.GongEnumValueShapes, stage.GongEnumValueShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongEnumValueShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongNoteLinkShape:
		tmp := GetStructInstancesByOrder(stage.GongNoteLinkShapes, stage.GongNoteLinkShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongNoteLinkShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongNoteShape:
		tmp := GetStructInstancesByOrder(stage.GongNoteShapes, stage.GongNoteShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongNoteShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongStructShape:
		tmp := GetStructInstancesByOrder(stage.GongStructShapes, stage.GongStructShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongStructShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LinkShape:
		tmp := GetStructInstancesByOrder(stage.LinkShapes, stage.LinkShapeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LinkShape implements.
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
	case "AttributeShape":
		res = GetNamedStructInstances(stage.AttributeShapes, stage.AttributeShapeMap_Staged_Order)
	case "Classdiagram":
		res = GetNamedStructInstances(stage.Classdiagrams, stage.ClassdiagramMap_Staged_Order)
	case "DiagramPackage":
		res = GetNamedStructInstances(stage.DiagramPackages, stage.DiagramPackageMap_Staged_Order)
	case "GongEnumShape":
		res = GetNamedStructInstances(stage.GongEnumShapes, stage.GongEnumShapeMap_Staged_Order)
	case "GongEnumValueShape":
		res = GetNamedStructInstances(stage.GongEnumValueShapes, stage.GongEnumValueShapeMap_Staged_Order)
	case "GongNoteLinkShape":
		res = GetNamedStructInstances(stage.GongNoteLinkShapes, stage.GongNoteLinkShapeMap_Staged_Order)
	case "GongNoteShape":
		res = GetNamedStructInstances(stage.GongNoteShapes, stage.GongNoteShapeMap_Staged_Order)
	case "GongStructShape":
		res = GetNamedStructInstances(stage.GongStructShapes, stage.GongStructShapeMap_Staged_Order)
	case "LinkShape":
		res = GetNamedStructInstances(stage.LinkShapes, stage.LinkShapeMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/doc/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return doc_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return doc_go.GoDiagramsDir
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
	CommitAttributeShape(attributeshape *AttributeShape)
	CheckoutAttributeShape(attributeshape *AttributeShape)
	CommitClassdiagram(classdiagram *Classdiagram)
	CheckoutClassdiagram(classdiagram *Classdiagram)
	CommitDiagramPackage(diagrampackage *DiagramPackage)
	CheckoutDiagramPackage(diagrampackage *DiagramPackage)
	CommitGongEnumShape(gongenumshape *GongEnumShape)
	CheckoutGongEnumShape(gongenumshape *GongEnumShape)
	CommitGongEnumValueShape(gongenumvalueshape *GongEnumValueShape)
	CheckoutGongEnumValueShape(gongenumvalueshape *GongEnumValueShape)
	CommitGongNoteLinkShape(gongnotelinkshape *GongNoteLinkShape)
	CheckoutGongNoteLinkShape(gongnotelinkshape *GongNoteLinkShape)
	CommitGongNoteShape(gongnoteshape *GongNoteShape)
	CheckoutGongNoteShape(gongnoteshape *GongNoteShape)
	CommitGongStructShape(gongstructshape *GongStructShape)
	CheckoutGongStructShape(gongstructshape *GongStructShape)
	CommitLinkShape(linkshape *LinkShape)
	CheckoutLinkShape(linkshape *LinkShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		AttributeShapes:           make(map[*AttributeShape]struct{}),
		AttributeShapes_mapString: make(map[string]*AttributeShape),

		Classdiagrams:           make(map[*Classdiagram]struct{}),
		Classdiagrams_mapString: make(map[string]*Classdiagram),

		DiagramPackages:           make(map[*DiagramPackage]struct{}),
		DiagramPackages_mapString: make(map[string]*DiagramPackage),

		GongEnumShapes:           make(map[*GongEnumShape]struct{}),
		GongEnumShapes_mapString: make(map[string]*GongEnumShape),

		GongEnumValueShapes:           make(map[*GongEnumValueShape]struct{}),
		GongEnumValueShapes_mapString: make(map[string]*GongEnumValueShape),

		GongNoteLinkShapes:           make(map[*GongNoteLinkShape]struct{}),
		GongNoteLinkShapes_mapString: make(map[string]*GongNoteLinkShape),

		GongNoteShapes:           make(map[*GongNoteShape]struct{}),
		GongNoteShapes_mapString: make(map[string]*GongNoteShape),

		GongStructShapes:           make(map[*GongStructShape]struct{}),
		GongStructShapes_mapString: make(map[string]*GongStructShape),

		LinkShapes:           make(map[*LinkShape]struct{}),
		LinkShapes_mapString: make(map[string]*LinkShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AttributeShapeMap_Staged_Order: make(map[*AttributeShape]uint),

		ClassdiagramMap_Staged_Order: make(map[*Classdiagram]uint),

		DiagramPackageMap_Staged_Order: make(map[*DiagramPackage]uint),

		GongEnumShapeMap_Staged_Order: make(map[*GongEnumShape]uint),

		GongEnumValueShapeMap_Staged_Order: make(map[*GongEnumValueShape]uint),

		GongNoteLinkShapeMap_Staged_Order: make(map[*GongNoteLinkShape]uint),

		GongNoteShapeMap_Staged_Order: make(map[*GongNoteShape]uint),

		GongStructShapeMap_Staged_Order: make(map[*GongStructShape]uint),

		LinkShapeMap_Staged_Order: make(map[*LinkShape]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AttributeShape"},
			{name: "Classdiagram"},
			{name: "DiagramPackage"},
			{name: "GongEnumShape"},
			{name: "GongEnumValueShape"},
			{name: "GongNoteLinkShape"},
			{name: "GongNoteShape"},
			{name: "GongStructShape"},
			{name: "LinkShape"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
		new:       make(map[GongstructIF]struct{}),
		modified:  make(map[GongstructIF]struct{}),
		deleted:   make(map[GongstructIF]struct{}),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AttributeShape:
		return stage.AttributeShapeMap_Staged_Order[instance]
	case *Classdiagram:
		return stage.ClassdiagramMap_Staged_Order[instance]
	case *DiagramPackage:
		return stage.DiagramPackageMap_Staged_Order[instance]
	case *GongEnumShape:
		return stage.GongEnumShapeMap_Staged_Order[instance]
	case *GongEnumValueShape:
		return stage.GongEnumValueShapeMap_Staged_Order[instance]
	case *GongNoteLinkShape:
		return stage.GongNoteLinkShapeMap_Staged_Order[instance]
	case *GongNoteShape:
		return stage.GongNoteShapeMap_Staged_Order[instance]
	case *GongStructShape:
		return stage.GongStructShapeMap_Staged_Order[instance]
	case *LinkShape:
		return stage.LinkShapeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AttributeShape:
		return stage.AttributeShapeMap_Staged_Order[instance]
	case *Classdiagram:
		return stage.ClassdiagramMap_Staged_Order[instance]
	case *DiagramPackage:
		return stage.DiagramPackageMap_Staged_Order[instance]
	case *GongEnumShape:
		return stage.GongEnumShapeMap_Staged_Order[instance]
	case *GongEnumValueShape:
		return stage.GongEnumValueShapeMap_Staged_Order[instance]
	case *GongNoteLinkShape:
		return stage.GongNoteLinkShapeMap_Staged_Order[instance]
	case *GongNoteShape:
		return stage.GongNoteShapeMap_Staged_Order[instance]
	case *GongStructShape:
		return stage.GongStructShapeMap_Staged_Order[instance]
	case *LinkShape:
		return stage.LinkShapeMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["AttributeShape"] = len(stage.AttributeShapes)
	stage.Map_GongStructName_InstancesNb["Classdiagram"] = len(stage.Classdiagrams)
	stage.Map_GongStructName_InstancesNb["DiagramPackage"] = len(stage.DiagramPackages)
	stage.Map_GongStructName_InstancesNb["GongEnumShape"] = len(stage.GongEnumShapes)
	stage.Map_GongStructName_InstancesNb["GongEnumValueShape"] = len(stage.GongEnumValueShapes)
	stage.Map_GongStructName_InstancesNb["GongNoteLinkShape"] = len(stage.GongNoteLinkShapes)
	stage.Map_GongStructName_InstancesNb["GongNoteShape"] = len(stage.GongNoteShapes)
	stage.Map_GongStructName_InstancesNb["GongStructShape"] = len(stage.GongStructShapes)
	stage.Map_GongStructName_InstancesNb["LinkShape"] = len(stage.LinkShapes)
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
// Stage puts attributeshape to the model stage
func (attributeshape *AttributeShape) Stage(stage *Stage) *AttributeShape {

	if _, ok := stage.AttributeShapes[attributeshape]; !ok {
		stage.AttributeShapes[attributeshape] = struct{}{}
		stage.AttributeShapeMap_Staged_Order[attributeshape] = stage.AttributeShapeOrder
		stage.AttributeShapeOrder++
		stage.new[attributeshape] = struct{}{}
		delete(stage.deleted, attributeshape)
	} else {
		if _, ok := stage.new[attributeshape]; !ok {
			stage.modified[attributeshape] = struct{}{}
		}
	}
	stage.AttributeShapes_mapString[attributeshape.Name] = attributeshape

	return attributeshape
}

// Unstage removes attributeshape off the model stage
func (attributeshape *AttributeShape) Unstage(stage *Stage) *AttributeShape {
	delete(stage.AttributeShapes, attributeshape)
	delete(stage.AttributeShapes_mapString, attributeshape.Name)

	if _, ok := stage.reference[attributeshape]; ok {
		stage.deleted[attributeshape] = struct{}{}
	} else {
		delete(stage.new, attributeshape)
	}
	return attributeshape
}

// UnstageVoid removes attributeshape off the model stage
func (attributeshape *AttributeShape) UnstageVoid(stage *Stage) {
	delete(stage.AttributeShapes, attributeshape)
	delete(stage.AttributeShapes_mapString, attributeshape.Name)
}

// commit attributeshape to the back repo (if it is already staged)
func (attributeshape *AttributeShape) Commit(stage *Stage) *AttributeShape {
	if _, ok := stage.AttributeShapes[attributeshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAttributeShape(attributeshape)
		}
	}
	return attributeshape
}

func (attributeshape *AttributeShape) CommitVoid(stage *Stage) {
	attributeshape.Commit(stage)
}

func (attributeshape *AttributeShape) StageVoid(stage *Stage) {
	attributeshape.Stage(stage)
}

// Checkout attributeshape to the back repo (if it is already staged)
func (attributeshape *AttributeShape) Checkout(stage *Stage) *AttributeShape {
	if _, ok := stage.AttributeShapes[attributeshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAttributeShape(attributeshape)
		}
	}
	return attributeshape
}

// for satisfaction of GongStruct interface
func (attributeshape *AttributeShape) GetName() (res string) {
	return attributeshape.Name
}

// for satisfaction of GongStruct interface
func (attributeshape *AttributeShape) SetName(name string) (){
	attributeshape.Name = name
}

// Stage puts classdiagram to the model stage
func (classdiagram *Classdiagram) Stage(stage *Stage) *Classdiagram {

	if _, ok := stage.Classdiagrams[classdiagram]; !ok {
		stage.Classdiagrams[classdiagram] = struct{}{}
		stage.ClassdiagramMap_Staged_Order[classdiagram] = stage.ClassdiagramOrder
		stage.ClassdiagramOrder++
		stage.new[classdiagram] = struct{}{}
		delete(stage.deleted, classdiagram)
	} else {
		if _, ok := stage.new[classdiagram]; !ok {
			stage.modified[classdiagram] = struct{}{}
		}
	}
	stage.Classdiagrams_mapString[classdiagram.Name] = classdiagram

	return classdiagram
}

// Unstage removes classdiagram off the model stage
func (classdiagram *Classdiagram) Unstage(stage *Stage) *Classdiagram {
	delete(stage.Classdiagrams, classdiagram)
	delete(stage.Classdiagrams_mapString, classdiagram.Name)

	if _, ok := stage.reference[classdiagram]; ok {
		stage.deleted[classdiagram] = struct{}{}
	} else {
		delete(stage.new, classdiagram)
	}
	return classdiagram
}

// UnstageVoid removes classdiagram off the model stage
func (classdiagram *Classdiagram) UnstageVoid(stage *Stage) {
	delete(stage.Classdiagrams, classdiagram)
	delete(stage.Classdiagrams_mapString, classdiagram.Name)
}

// commit classdiagram to the back repo (if it is already staged)
func (classdiagram *Classdiagram) Commit(stage *Stage) *Classdiagram {
	if _, ok := stage.Classdiagrams[classdiagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitClassdiagram(classdiagram)
		}
	}
	return classdiagram
}

func (classdiagram *Classdiagram) CommitVoid(stage *Stage) {
	classdiagram.Commit(stage)
}

func (classdiagram *Classdiagram) StageVoid(stage *Stage) {
	classdiagram.Stage(stage)
}

// Checkout classdiagram to the back repo (if it is already staged)
func (classdiagram *Classdiagram) Checkout(stage *Stage) *Classdiagram {
	if _, ok := stage.Classdiagrams[classdiagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutClassdiagram(classdiagram)
		}
	}
	return classdiagram
}

// for satisfaction of GongStruct interface
func (classdiagram *Classdiagram) GetName() (res string) {
	return classdiagram.Name
}

// for satisfaction of GongStruct interface
func (classdiagram *Classdiagram) SetName(name string) (){
	classdiagram.Name = name
}

// Stage puts diagrampackage to the model stage
func (diagrampackage *DiagramPackage) Stage(stage *Stage) *DiagramPackage {

	if _, ok := stage.DiagramPackages[diagrampackage]; !ok {
		stage.DiagramPackages[diagrampackage] = struct{}{}
		stage.DiagramPackageMap_Staged_Order[diagrampackage] = stage.DiagramPackageOrder
		stage.DiagramPackageOrder++
		stage.new[diagrampackage] = struct{}{}
		delete(stage.deleted, diagrampackage)
	} else {
		if _, ok := stage.new[diagrampackage]; !ok {
			stage.modified[diagrampackage] = struct{}{}
		}
	}
	stage.DiagramPackages_mapString[diagrampackage.Name] = diagrampackage

	return diagrampackage
}

// Unstage removes diagrampackage off the model stage
func (diagrampackage *DiagramPackage) Unstage(stage *Stage) *DiagramPackage {
	delete(stage.DiagramPackages, diagrampackage)
	delete(stage.DiagramPackages_mapString, diagrampackage.Name)

	if _, ok := stage.reference[diagrampackage]; ok {
		stage.deleted[diagrampackage] = struct{}{}
	} else {
		delete(stage.new, diagrampackage)
	}
	return diagrampackage
}

// UnstageVoid removes diagrampackage off the model stage
func (diagrampackage *DiagramPackage) UnstageVoid(stage *Stage) {
	delete(stage.DiagramPackages, diagrampackage)
	delete(stage.DiagramPackages_mapString, diagrampackage.Name)
}

// commit diagrampackage to the back repo (if it is already staged)
func (diagrampackage *DiagramPackage) Commit(stage *Stage) *DiagramPackage {
	if _, ok := stage.DiagramPackages[diagrampackage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramPackage(diagrampackage)
		}
	}
	return diagrampackage
}

func (diagrampackage *DiagramPackage) CommitVoid(stage *Stage) {
	diagrampackage.Commit(stage)
}

func (diagrampackage *DiagramPackage) StageVoid(stage *Stage) {
	diagrampackage.Stage(stage)
}

// Checkout diagrampackage to the back repo (if it is already staged)
func (diagrampackage *DiagramPackage) Checkout(stage *Stage) *DiagramPackage {
	if _, ok := stage.DiagramPackages[diagrampackage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramPackage(diagrampackage)
		}
	}
	return diagrampackage
}

// for satisfaction of GongStruct interface
func (diagrampackage *DiagramPackage) GetName() (res string) {
	return diagrampackage.Name
}

// for satisfaction of GongStruct interface
func (diagrampackage *DiagramPackage) SetName(name string) (){
	diagrampackage.Name = name
}

// Stage puts gongenumshape to the model stage
func (gongenumshape *GongEnumShape) Stage(stage *Stage) *GongEnumShape {

	if _, ok := stage.GongEnumShapes[gongenumshape]; !ok {
		stage.GongEnumShapes[gongenumshape] = struct{}{}
		stage.GongEnumShapeMap_Staged_Order[gongenumshape] = stage.GongEnumShapeOrder
		stage.GongEnumShapeOrder++
		stage.new[gongenumshape] = struct{}{}
		delete(stage.deleted, gongenumshape)
	} else {
		if _, ok := stage.new[gongenumshape]; !ok {
			stage.modified[gongenumshape] = struct{}{}
		}
	}
	stage.GongEnumShapes_mapString[gongenumshape.Name] = gongenumshape

	return gongenumshape
}

// Unstage removes gongenumshape off the model stage
func (gongenumshape *GongEnumShape) Unstage(stage *Stage) *GongEnumShape {
	delete(stage.GongEnumShapes, gongenumshape)
	delete(stage.GongEnumShapes_mapString, gongenumshape.Name)

	if _, ok := stage.reference[gongenumshape]; ok {
		stage.deleted[gongenumshape] = struct{}{}
	} else {
		delete(stage.new, gongenumshape)
	}
	return gongenumshape
}

// UnstageVoid removes gongenumshape off the model stage
func (gongenumshape *GongEnumShape) UnstageVoid(stage *Stage) {
	delete(stage.GongEnumShapes, gongenumshape)
	delete(stage.GongEnumShapes_mapString, gongenumshape.Name)
}

// commit gongenumshape to the back repo (if it is already staged)
func (gongenumshape *GongEnumShape) Commit(stage *Stage) *GongEnumShape {
	if _, ok := stage.GongEnumShapes[gongenumshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnumShape(gongenumshape)
		}
	}
	return gongenumshape
}

func (gongenumshape *GongEnumShape) CommitVoid(stage *Stage) {
	gongenumshape.Commit(stage)
}

func (gongenumshape *GongEnumShape) StageVoid(stage *Stage) {
	gongenumshape.Stage(stage)
}

// Checkout gongenumshape to the back repo (if it is already staged)
func (gongenumshape *GongEnumShape) Checkout(stage *Stage) *GongEnumShape {
	if _, ok := stage.GongEnumShapes[gongenumshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnumShape(gongenumshape)
		}
	}
	return gongenumshape
}

// for satisfaction of GongStruct interface
func (gongenumshape *GongEnumShape) GetName() (res string) {
	return gongenumshape.Name
}

// for satisfaction of GongStruct interface
func (gongenumshape *GongEnumShape) SetName(name string) (){
	gongenumshape.Name = name
}

// Stage puts gongenumvalueshape to the model stage
func (gongenumvalueshape *GongEnumValueShape) Stage(stage *Stage) *GongEnumValueShape {

	if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; !ok {
		stage.GongEnumValueShapes[gongenumvalueshape] = struct{}{}
		stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshape] = stage.GongEnumValueShapeOrder
		stage.GongEnumValueShapeOrder++
		stage.new[gongenumvalueshape] = struct{}{}
		delete(stage.deleted, gongenumvalueshape)
	} else {
		if _, ok := stage.new[gongenumvalueshape]; !ok {
			stage.modified[gongenumvalueshape] = struct{}{}
		}
	}
	stage.GongEnumValueShapes_mapString[gongenumvalueshape.Name] = gongenumvalueshape

	return gongenumvalueshape
}

// Unstage removes gongenumvalueshape off the model stage
func (gongenumvalueshape *GongEnumValueShape) Unstage(stage *Stage) *GongEnumValueShape {
	delete(stage.GongEnumValueShapes, gongenumvalueshape)
	delete(stage.GongEnumValueShapes_mapString, gongenumvalueshape.Name)

	if _, ok := stage.reference[gongenumvalueshape]; ok {
		stage.deleted[gongenumvalueshape] = struct{}{}
	} else {
		delete(stage.new, gongenumvalueshape)
	}
	return gongenumvalueshape
}

// UnstageVoid removes gongenumvalueshape off the model stage
func (gongenumvalueshape *GongEnumValueShape) UnstageVoid(stage *Stage) {
	delete(stage.GongEnumValueShapes, gongenumvalueshape)
	delete(stage.GongEnumValueShapes_mapString, gongenumvalueshape.Name)
}

// commit gongenumvalueshape to the back repo (if it is already staged)
func (gongenumvalueshape *GongEnumValueShape) Commit(stage *Stage) *GongEnumValueShape {
	if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnumValueShape(gongenumvalueshape)
		}
	}
	return gongenumvalueshape
}

func (gongenumvalueshape *GongEnumValueShape) CommitVoid(stage *Stage) {
	gongenumvalueshape.Commit(stage)
}

func (gongenumvalueshape *GongEnumValueShape) StageVoid(stage *Stage) {
	gongenumvalueshape.Stage(stage)
}

// Checkout gongenumvalueshape to the back repo (if it is already staged)
func (gongenumvalueshape *GongEnumValueShape) Checkout(stage *Stage) *GongEnumValueShape {
	if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnumValueShape(gongenumvalueshape)
		}
	}
	return gongenumvalueshape
}

// for satisfaction of GongStruct interface
func (gongenumvalueshape *GongEnumValueShape) GetName() (res string) {
	return gongenumvalueshape.Name
}

// for satisfaction of GongStruct interface
func (gongenumvalueshape *GongEnumValueShape) SetName(name string) (){
	gongenumvalueshape.Name = name
}

// Stage puts gongnotelinkshape to the model stage
func (gongnotelinkshape *GongNoteLinkShape) Stage(stage *Stage) *GongNoteLinkShape {

	if _, ok := stage.GongNoteLinkShapes[gongnotelinkshape]; !ok {
		stage.GongNoteLinkShapes[gongnotelinkshape] = struct{}{}
		stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshape] = stage.GongNoteLinkShapeOrder
		stage.GongNoteLinkShapeOrder++
		stage.new[gongnotelinkshape] = struct{}{}
		delete(stage.deleted, gongnotelinkshape)
	} else {
		if _, ok := stage.new[gongnotelinkshape]; !ok {
			stage.modified[gongnotelinkshape] = struct{}{}
		}
	}
	stage.GongNoteLinkShapes_mapString[gongnotelinkshape.Name] = gongnotelinkshape

	return gongnotelinkshape
}

// Unstage removes gongnotelinkshape off the model stage
func (gongnotelinkshape *GongNoteLinkShape) Unstage(stage *Stage) *GongNoteLinkShape {
	delete(stage.GongNoteLinkShapes, gongnotelinkshape)
	delete(stage.GongNoteLinkShapes_mapString, gongnotelinkshape.Name)

	if _, ok := stage.reference[gongnotelinkshape]; ok {
		stage.deleted[gongnotelinkshape] = struct{}{}
	} else {
		delete(stage.new, gongnotelinkshape)
	}
	return gongnotelinkshape
}

// UnstageVoid removes gongnotelinkshape off the model stage
func (gongnotelinkshape *GongNoteLinkShape) UnstageVoid(stage *Stage) {
	delete(stage.GongNoteLinkShapes, gongnotelinkshape)
	delete(stage.GongNoteLinkShapes_mapString, gongnotelinkshape.Name)
}

// commit gongnotelinkshape to the back repo (if it is already staged)
func (gongnotelinkshape *GongNoteLinkShape) Commit(stage *Stage) *GongNoteLinkShape {
	if _, ok := stage.GongNoteLinkShapes[gongnotelinkshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongNoteLinkShape(gongnotelinkshape)
		}
	}
	return gongnotelinkshape
}

func (gongnotelinkshape *GongNoteLinkShape) CommitVoid(stage *Stage) {
	gongnotelinkshape.Commit(stage)
}

func (gongnotelinkshape *GongNoteLinkShape) StageVoid(stage *Stage) {
	gongnotelinkshape.Stage(stage)
}

// Checkout gongnotelinkshape to the back repo (if it is already staged)
func (gongnotelinkshape *GongNoteLinkShape) Checkout(stage *Stage) *GongNoteLinkShape {
	if _, ok := stage.GongNoteLinkShapes[gongnotelinkshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongNoteLinkShape(gongnotelinkshape)
		}
	}
	return gongnotelinkshape
}

// for satisfaction of GongStruct interface
func (gongnotelinkshape *GongNoteLinkShape) GetName() (res string) {
	return gongnotelinkshape.Name
}

// for satisfaction of GongStruct interface
func (gongnotelinkshape *GongNoteLinkShape) SetName(name string) (){
	gongnotelinkshape.Name = name
}

// Stage puts gongnoteshape to the model stage
func (gongnoteshape *GongNoteShape) Stage(stage *Stage) *GongNoteShape {

	if _, ok := stage.GongNoteShapes[gongnoteshape]; !ok {
		stage.GongNoteShapes[gongnoteshape] = struct{}{}
		stage.GongNoteShapeMap_Staged_Order[gongnoteshape] = stage.GongNoteShapeOrder
		stage.GongNoteShapeOrder++
		stage.new[gongnoteshape] = struct{}{}
		delete(stage.deleted, gongnoteshape)
	} else {
		if _, ok := stage.new[gongnoteshape]; !ok {
			stage.modified[gongnoteshape] = struct{}{}
		}
	}
	stage.GongNoteShapes_mapString[gongnoteshape.Name] = gongnoteshape

	return gongnoteshape
}

// Unstage removes gongnoteshape off the model stage
func (gongnoteshape *GongNoteShape) Unstage(stage *Stage) *GongNoteShape {
	delete(stage.GongNoteShapes, gongnoteshape)
	delete(stage.GongNoteShapes_mapString, gongnoteshape.Name)

	if _, ok := stage.reference[gongnoteshape]; ok {
		stage.deleted[gongnoteshape] = struct{}{}
	} else {
		delete(stage.new, gongnoteshape)
	}
	return gongnoteshape
}

// UnstageVoid removes gongnoteshape off the model stage
func (gongnoteshape *GongNoteShape) UnstageVoid(stage *Stage) {
	delete(stage.GongNoteShapes, gongnoteshape)
	delete(stage.GongNoteShapes_mapString, gongnoteshape.Name)
}

// commit gongnoteshape to the back repo (if it is already staged)
func (gongnoteshape *GongNoteShape) Commit(stage *Stage) *GongNoteShape {
	if _, ok := stage.GongNoteShapes[gongnoteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongNoteShape(gongnoteshape)
		}
	}
	return gongnoteshape
}

func (gongnoteshape *GongNoteShape) CommitVoid(stage *Stage) {
	gongnoteshape.Commit(stage)
}

func (gongnoteshape *GongNoteShape) StageVoid(stage *Stage) {
	gongnoteshape.Stage(stage)
}

// Checkout gongnoteshape to the back repo (if it is already staged)
func (gongnoteshape *GongNoteShape) Checkout(stage *Stage) *GongNoteShape {
	if _, ok := stage.GongNoteShapes[gongnoteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongNoteShape(gongnoteshape)
		}
	}
	return gongnoteshape
}

// for satisfaction of GongStruct interface
func (gongnoteshape *GongNoteShape) GetName() (res string) {
	return gongnoteshape.Name
}

// for satisfaction of GongStruct interface
func (gongnoteshape *GongNoteShape) SetName(name string) (){
	gongnoteshape.Name = name
}

// Stage puts gongstructshape to the model stage
func (gongstructshape *GongStructShape) Stage(stage *Stage) *GongStructShape {

	if _, ok := stage.GongStructShapes[gongstructshape]; !ok {
		stage.GongStructShapes[gongstructshape] = struct{}{}
		stage.GongStructShapeMap_Staged_Order[gongstructshape] = stage.GongStructShapeOrder
		stage.GongStructShapeOrder++
		stage.new[gongstructshape] = struct{}{}
		delete(stage.deleted, gongstructshape)
	} else {
		if _, ok := stage.new[gongstructshape]; !ok {
			stage.modified[gongstructshape] = struct{}{}
		}
	}
	stage.GongStructShapes_mapString[gongstructshape.Name] = gongstructshape

	return gongstructshape
}

// Unstage removes gongstructshape off the model stage
func (gongstructshape *GongStructShape) Unstage(stage *Stage) *GongStructShape {
	delete(stage.GongStructShapes, gongstructshape)
	delete(stage.GongStructShapes_mapString, gongstructshape.Name)

	if _, ok := stage.reference[gongstructshape]; ok {
		stage.deleted[gongstructshape] = struct{}{}
	} else {
		delete(stage.new, gongstructshape)
	}
	return gongstructshape
}

// UnstageVoid removes gongstructshape off the model stage
func (gongstructshape *GongStructShape) UnstageVoid(stage *Stage) {
	delete(stage.GongStructShapes, gongstructshape)
	delete(stage.GongStructShapes_mapString, gongstructshape.Name)
}

// commit gongstructshape to the back repo (if it is already staged)
func (gongstructshape *GongStructShape) Commit(stage *Stage) *GongStructShape {
	if _, ok := stage.GongStructShapes[gongstructshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongStructShape(gongstructshape)
		}
	}
	return gongstructshape
}

func (gongstructshape *GongStructShape) CommitVoid(stage *Stage) {
	gongstructshape.Commit(stage)
}

func (gongstructshape *GongStructShape) StageVoid(stage *Stage) {
	gongstructshape.Stage(stage)
}

// Checkout gongstructshape to the back repo (if it is already staged)
func (gongstructshape *GongStructShape) Checkout(stage *Stage) *GongStructShape {
	if _, ok := stage.GongStructShapes[gongstructshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongStructShape(gongstructshape)
		}
	}
	return gongstructshape
}

// for satisfaction of GongStruct interface
func (gongstructshape *GongStructShape) GetName() (res string) {
	return gongstructshape.Name
}

// for satisfaction of GongStruct interface
func (gongstructshape *GongStructShape) SetName(name string) (){
	gongstructshape.Name = name
}

// Stage puts linkshape to the model stage
func (linkshape *LinkShape) Stage(stage *Stage) *LinkShape {

	if _, ok := stage.LinkShapes[linkshape]; !ok {
		stage.LinkShapes[linkshape] = struct{}{}
		stage.LinkShapeMap_Staged_Order[linkshape] = stage.LinkShapeOrder
		stage.LinkShapeOrder++
		stage.new[linkshape] = struct{}{}
		delete(stage.deleted, linkshape)
	} else {
		if _, ok := stage.new[linkshape]; !ok {
			stage.modified[linkshape] = struct{}{}
		}
	}
	stage.LinkShapes_mapString[linkshape.Name] = linkshape

	return linkshape
}

// Unstage removes linkshape off the model stage
func (linkshape *LinkShape) Unstage(stage *Stage) *LinkShape {
	delete(stage.LinkShapes, linkshape)
	delete(stage.LinkShapes_mapString, linkshape.Name)

	if _, ok := stage.reference[linkshape]; ok {
		stage.deleted[linkshape] = struct{}{}
	} else {
		delete(stage.new, linkshape)
	}
	return linkshape
}

// UnstageVoid removes linkshape off the model stage
func (linkshape *LinkShape) UnstageVoid(stage *Stage) {
	delete(stage.LinkShapes, linkshape)
	delete(stage.LinkShapes_mapString, linkshape.Name)
}

// commit linkshape to the back repo (if it is already staged)
func (linkshape *LinkShape) Commit(stage *Stage) *LinkShape {
	if _, ok := stage.LinkShapes[linkshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLinkShape(linkshape)
		}
	}
	return linkshape
}

func (linkshape *LinkShape) CommitVoid(stage *Stage) {
	linkshape.Commit(stage)
}

func (linkshape *LinkShape) StageVoid(stage *Stage) {
	linkshape.Stage(stage)
}

// Checkout linkshape to the back repo (if it is already staged)
func (linkshape *LinkShape) Checkout(stage *Stage) *LinkShape {
	if _, ok := stage.LinkShapes[linkshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLinkShape(linkshape)
		}
	}
	return linkshape
}

// for satisfaction of GongStruct interface
func (linkshape *LinkShape) GetName() (res string) {
	return linkshape.Name
}

// for satisfaction of GongStruct interface
func (linkshape *LinkShape) SetName(name string) (){
	linkshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAttributeShape(AttributeShape *AttributeShape)
	CreateORMClassdiagram(Classdiagram *Classdiagram)
	CreateORMDiagramPackage(DiagramPackage *DiagramPackage)
	CreateORMGongEnumShape(GongEnumShape *GongEnumShape)
	CreateORMGongEnumValueShape(GongEnumValueShape *GongEnumValueShape)
	CreateORMGongNoteLinkShape(GongNoteLinkShape *GongNoteLinkShape)
	CreateORMGongNoteShape(GongNoteShape *GongNoteShape)
	CreateORMGongStructShape(GongStructShape *GongStructShape)
	CreateORMLinkShape(LinkShape *LinkShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAttributeShape(AttributeShape *AttributeShape)
	DeleteORMClassdiagram(Classdiagram *Classdiagram)
	DeleteORMDiagramPackage(DiagramPackage *DiagramPackage)
	DeleteORMGongEnumShape(GongEnumShape *GongEnumShape)
	DeleteORMGongEnumValueShape(GongEnumValueShape *GongEnumValueShape)
	DeleteORMGongNoteLinkShape(GongNoteLinkShape *GongNoteLinkShape)
	DeleteORMGongNoteShape(GongNoteShape *GongNoteShape)
	DeleteORMGongStructShape(GongStructShape *GongStructShape)
	DeleteORMLinkShape(LinkShape *LinkShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AttributeShapes = make(map[*AttributeShape]struct{})
	stage.AttributeShapes_mapString = make(map[string]*AttributeShape)
	stage.AttributeShapeMap_Staged_Order = make(map[*AttributeShape]uint)
	stage.AttributeShapeOrder = 0

	stage.Classdiagrams = make(map[*Classdiagram]struct{})
	stage.Classdiagrams_mapString = make(map[string]*Classdiagram)
	stage.ClassdiagramMap_Staged_Order = make(map[*Classdiagram]uint)
	stage.ClassdiagramOrder = 0

	stage.DiagramPackages = make(map[*DiagramPackage]struct{})
	stage.DiagramPackages_mapString = make(map[string]*DiagramPackage)
	stage.DiagramPackageMap_Staged_Order = make(map[*DiagramPackage]uint)
	stage.DiagramPackageOrder = 0

	stage.GongEnumShapes = make(map[*GongEnumShape]struct{})
	stage.GongEnumShapes_mapString = make(map[string]*GongEnumShape)
	stage.GongEnumShapeMap_Staged_Order = make(map[*GongEnumShape]uint)
	stage.GongEnumShapeOrder = 0

	stage.GongEnumValueShapes = make(map[*GongEnumValueShape]struct{})
	stage.GongEnumValueShapes_mapString = make(map[string]*GongEnumValueShape)
	stage.GongEnumValueShapeMap_Staged_Order = make(map[*GongEnumValueShape]uint)
	stage.GongEnumValueShapeOrder = 0

	stage.GongNoteLinkShapes = make(map[*GongNoteLinkShape]struct{})
	stage.GongNoteLinkShapes_mapString = make(map[string]*GongNoteLinkShape)
	stage.GongNoteLinkShapeMap_Staged_Order = make(map[*GongNoteLinkShape]uint)
	stage.GongNoteLinkShapeOrder = 0

	stage.GongNoteShapes = make(map[*GongNoteShape]struct{})
	stage.GongNoteShapes_mapString = make(map[string]*GongNoteShape)
	stage.GongNoteShapeMap_Staged_Order = make(map[*GongNoteShape]uint)
	stage.GongNoteShapeOrder = 0

	stage.GongStructShapes = make(map[*GongStructShape]struct{})
	stage.GongStructShapes_mapString = make(map[string]*GongStructShape)
	stage.GongStructShapeMap_Staged_Order = make(map[*GongStructShape]uint)
	stage.GongStructShapeOrder = 0

	stage.LinkShapes = make(map[*LinkShape]struct{})
	stage.LinkShapes_mapString = make(map[string]*LinkShape)
	stage.LinkShapeMap_Staged_Order = make(map[*LinkShape]uint)
	stage.LinkShapeOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AttributeShapes = nil
	stage.AttributeShapes_mapString = nil

	stage.Classdiagrams = nil
	stage.Classdiagrams_mapString = nil

	stage.DiagramPackages = nil
	stage.DiagramPackages_mapString = nil

	stage.GongEnumShapes = nil
	stage.GongEnumShapes_mapString = nil

	stage.GongEnumValueShapes = nil
	stage.GongEnumValueShapes_mapString = nil

	stage.GongNoteLinkShapes = nil
	stage.GongNoteLinkShapes_mapString = nil

	stage.GongNoteShapes = nil
	stage.GongNoteShapes_mapString = nil

	stage.GongStructShapes = nil
	stage.GongStructShapes_mapString = nil

	stage.LinkShapes = nil
	stage.LinkShapes_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for attributeshape := range stage.AttributeShapes {
		attributeshape.Unstage(stage)
	}

	for classdiagram := range stage.Classdiagrams {
		classdiagram.Unstage(stage)
	}

	for diagrampackage := range stage.DiagramPackages {
		diagrampackage.Unstage(stage)
	}

	for gongenumshape := range stage.GongEnumShapes {
		gongenumshape.Unstage(stage)
	}

	for gongenumvalueshape := range stage.GongEnumValueShapes {
		gongenumvalueshape.Unstage(stage)
	}

	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		gongnotelinkshape.Unstage(stage)
	}

	for gongnoteshape := range stage.GongNoteShapes {
		gongnoteshape.Unstage(stage)
	}

	for gongstructshape := range stage.GongStructShapes {
		gongstructshape.Unstage(stage)
	}

	for linkshape := range stage.LinkShapes {
		linkshape.Unstage(stage)
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
	case map[*AttributeShape]any:
		return any(&stage.AttributeShapes).(*Type)
	case map[*Classdiagram]any:
		return any(&stage.Classdiagrams).(*Type)
	case map[*DiagramPackage]any:
		return any(&stage.DiagramPackages).(*Type)
	case map[*GongEnumShape]any:
		return any(&stage.GongEnumShapes).(*Type)
	case map[*GongEnumValueShape]any:
		return any(&stage.GongEnumValueShapes).(*Type)
	case map[*GongNoteLinkShape]any:
		return any(&stage.GongNoteLinkShapes).(*Type)
	case map[*GongNoteShape]any:
		return any(&stage.GongNoteShapes).(*Type)
	case map[*GongStructShape]any:
		return any(&stage.GongStructShapes).(*Type)
	case map[*LinkShape]any:
		return any(&stage.LinkShapes).(*Type)
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
	case *AttributeShape:
		return any(stage.AttributeShapes_mapString).(map[string]Type)
	case *Classdiagram:
		return any(stage.Classdiagrams_mapString).(map[string]Type)
	case *DiagramPackage:
		return any(stage.DiagramPackages_mapString).(map[string]Type)
	case *GongEnumShape:
		return any(stage.GongEnumShapes_mapString).(map[string]Type)
	case *GongEnumValueShape:
		return any(stage.GongEnumValueShapes_mapString).(map[string]Type)
	case *GongNoteLinkShape:
		return any(stage.GongNoteLinkShapes_mapString).(map[string]Type)
	case *GongNoteShape:
		return any(stage.GongNoteShapes_mapString).(map[string]Type)
	case *GongStructShape:
		return any(stage.GongStructShapes_mapString).(map[string]Type)
	case *LinkShape:
		return any(stage.LinkShapes_mapString).(map[string]Type)
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
	case AttributeShape:
		return any(&stage.AttributeShapes).(*map[*Type]struct{})
	case Classdiagram:
		return any(&stage.Classdiagrams).(*map[*Type]struct{})
	case DiagramPackage:
		return any(&stage.DiagramPackages).(*map[*Type]struct{})
	case GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[*Type]struct{})
	case GongEnumValueShape:
		return any(&stage.GongEnumValueShapes).(*map[*Type]struct{})
	case GongNoteLinkShape:
		return any(&stage.GongNoteLinkShapes).(*map[*Type]struct{})
	case GongNoteShape:
		return any(&stage.GongNoteShapes).(*map[*Type]struct{})
	case GongStructShape:
		return any(&stage.GongStructShapes).(*map[*Type]struct{})
	case LinkShape:
		return any(&stage.LinkShapes).(*map[*Type]struct{})
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
	case *AttributeShape:
		return any(&stage.AttributeShapes).(*map[Type]struct{})
	case *Classdiagram:
		return any(&stage.Classdiagrams).(*map[Type]struct{})
	case *DiagramPackage:
		return any(&stage.DiagramPackages).(*map[Type]struct{})
	case *GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[Type]struct{})
	case *GongEnumValueShape:
		return any(&stage.GongEnumValueShapes).(*map[Type]struct{})
	case *GongNoteLinkShape:
		return any(&stage.GongNoteLinkShapes).(*map[Type]struct{})
	case *GongNoteShape:
		return any(&stage.GongNoteShapes).(*map[Type]struct{})
	case *GongStructShape:
		return any(&stage.GongStructShapes).(*map[Type]struct{})
	case *LinkShape:
		return any(&stage.LinkShapes).(*map[Type]struct{})
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
	case AttributeShape:
		return any(&stage.AttributeShapes_mapString).(*map[string]*Type)
	case Classdiagram:
		return any(&stage.Classdiagrams_mapString).(*map[string]*Type)
	case DiagramPackage:
		return any(&stage.DiagramPackages_mapString).(*map[string]*Type)
	case GongEnumShape:
		return any(&stage.GongEnumShapes_mapString).(*map[string]*Type)
	case GongEnumValueShape:
		return any(&stage.GongEnumValueShapes_mapString).(*map[string]*Type)
	case GongNoteLinkShape:
		return any(&stage.GongNoteLinkShapes_mapString).(*map[string]*Type)
	case GongNoteShape:
		return any(&stage.GongNoteShapes_mapString).(*map[string]*Type)
	case GongStructShape:
		return any(&stage.GongStructShapes_mapString).(*map[string]*Type)
	case LinkShape:
		return any(&stage.LinkShapes_mapString).(*map[string]*Type)
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
	case AttributeShape:
		return any(&AttributeShape{
			// Initialisation of associations
		}).(*Type)
	case Classdiagram:
		return any(&Classdiagram{
			// Initialisation of associations
			// field is initialized with an instance of GongStructShape with the name of the field
			GongStructShapes: []*GongStructShape{{Name: "GongStructShapes"}},
			// field is initialized with an instance of GongEnumShape with the name of the field
			GongEnumShapes: []*GongEnumShape{{Name: "GongEnumShapes"}},
			// field is initialized with an instance of GongNoteShape with the name of the field
			GongNoteShapes: []*GongNoteShape{{Name: "GongNoteShapes"}},
		}).(*Type)
	case DiagramPackage:
		return any(&DiagramPackage{
			// Initialisation of associations
			// field is initialized with an instance of Classdiagram with the name of the field
			Classdiagrams: []*Classdiagram{{Name: "Classdiagrams"}},
			// field is initialized with an instance of Classdiagram with the name of the field
			SelectedClassdiagram: &Classdiagram{Name: "SelectedClassdiagram"},
		}).(*Type)
	case GongEnumShape:
		return any(&GongEnumShape{
			// Initialisation of associations
			// field is initialized with an instance of GongEnumValueShape with the name of the field
			GongEnumValueShapes: []*GongEnumValueShape{{Name: "GongEnumValueShapes"}},
		}).(*Type)
	case GongEnumValueShape:
		return any(&GongEnumValueShape{
			// Initialisation of associations
		}).(*Type)
	case GongNoteLinkShape:
		return any(&GongNoteLinkShape{
			// Initialisation of associations
		}).(*Type)
	case GongNoteShape:
		return any(&GongNoteShape{
			// Initialisation of associations
			// field is initialized with an instance of GongNoteLinkShape with the name of the field
			GongNoteLinkShapes: []*GongNoteLinkShape{{Name: "GongNoteLinkShapes"}},
		}).(*Type)
	case GongStructShape:
		return any(&GongStructShape{
			// Initialisation of associations
			// field is initialized with an instance of AttributeShape with the name of the field
			AttributeShapes: []*AttributeShape{{Name: "AttributeShapes"}},
			// field is initialized with an instance of LinkShape with the name of the field
			LinkShapes: []*LinkShape{{Name: "LinkShapes"}},
		}).(*Type)
	case LinkShape:
		return any(&LinkShape{
			// Initialisation of associations
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
	// reverse maps of direct associations of AttributeShape
	case AttributeShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Classdiagram
	case Classdiagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramPackage
	case DiagramPackage:
		switch fieldname {
		// insertion point for per direct association field
		case "SelectedClassdiagram":
			res := make(map[*Classdiagram][]*DiagramPackage)
			for diagrampackage := range stage.DiagramPackages {
				if diagrampackage.SelectedClassdiagram != nil {
					classdiagram_ := diagrampackage.SelectedClassdiagram
					var diagrampackages []*DiagramPackage
					_, ok := res[classdiagram_]
					if ok {
						diagrampackages = res[classdiagram_]
					} else {
						diagrampackages = make([]*DiagramPackage, 0)
					}
					diagrampackages = append(diagrampackages, diagrampackage)
					res[classdiagram_] = diagrampackages
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumShape
	case GongEnumShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumValueShape
	case GongEnumValueShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNoteLinkShape
	case GongNoteLinkShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNoteShape
	case GongNoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStructShape
	case GongStructShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LinkShape
	case LinkShape:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of AttributeShape
	case AttributeShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Classdiagram
	case Classdiagram:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStructShapes":
			res := make(map[*GongStructShape][]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, gongstructshape_ := range classdiagram.GongStructShapes {
					res[gongstructshape_] = append(res[gongstructshape_], classdiagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GongEnumShapes":
			res := make(map[*GongEnumShape][]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, gongenumshape_ := range classdiagram.GongEnumShapes {
					res[gongenumshape_] = append(res[gongenumshape_], classdiagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GongNoteShapes":
			res := make(map[*GongNoteShape][]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, gongnoteshape_ := range classdiagram.GongNoteShapes {
					res[gongnoteshape_] = append(res[gongnoteshape_], classdiagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramPackage
	case DiagramPackage:
		switch fieldname {
		// insertion point for per direct association field
		case "Classdiagrams":
			res := make(map[*Classdiagram][]*DiagramPackage)
			for diagrampackage := range stage.DiagramPackages {
				for _, classdiagram_ := range diagrampackage.Classdiagrams {
					res[classdiagram_] = append(res[classdiagram_], diagrampackage)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumShape
	case GongEnumShape:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValueShapes":
			res := make(map[*GongEnumValueShape][]*GongEnumShape)
			for gongenumshape := range stage.GongEnumShapes {
				for _, gongenumvalueshape_ := range gongenumshape.GongEnumValueShapes {
					res[gongenumvalueshape_] = append(res[gongenumvalueshape_], gongenumshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumValueShape
	case GongEnumValueShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNoteLinkShape
	case GongNoteLinkShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNoteShape
	case GongNoteShape:
		switch fieldname {
		// insertion point for per direct association field
		case "GongNoteLinkShapes":
			res := make(map[*GongNoteLinkShape][]*GongNoteShape)
			for gongnoteshape := range stage.GongNoteShapes {
				for _, gongnotelinkshape_ := range gongnoteshape.GongNoteLinkShapes {
					res[gongnotelinkshape_] = append(res[gongnotelinkshape_], gongnoteshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongStructShape
	case GongStructShape:
		switch fieldname {
		// insertion point for per direct association field
		case "AttributeShapes":
			res := make(map[*AttributeShape][]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, attributeshape_ := range gongstructshape.AttributeShapes {
					res[attributeshape_] = append(res[attributeshape_], gongstructshape)
				}
			}
			return any(res).(map[*End][]*Start)
		case "LinkShapes":
			res := make(map[*LinkShape][]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, linkshape_ := range gongstructshape.LinkShapes {
					res[linkshape_] = append(res[linkshape_], gongstructshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinkShape
	case LinkShape:
		switch fieldname {
		// insertion point for per direct association field
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
	case *AttributeShape:
		res = "AttributeShape"
	case *Classdiagram:
		res = "Classdiagram"
	case *DiagramPackage:
		res = "DiagramPackage"
	case *GongEnumShape:
		res = "GongEnumShape"
	case *GongEnumValueShape:
		res = "GongEnumValueShape"
	case *GongNoteLinkShape:
		res = "GongNoteLinkShape"
	case *GongNoteShape:
		res = "GongNoteShape"
	case *GongStructShape:
		res = "GongStructShape"
	case *LinkShape:
		res = "LinkShape"
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
	case *AttributeShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "AttributeShapes"
		res = append(res, rf)
	case *Classdiagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramPackage"
		rf.Fieldname = "Classdiagrams"
		res = append(res, rf)
	case *DiagramPackage:
		var rf ReverseField
		_ = rf
	case *GongEnumShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongEnumShapes"
		res = append(res, rf)
	case *GongEnumValueShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongEnumShape"
		rf.Fieldname = "GongEnumValueShapes"
		res = append(res, rf)
	case *GongNoteLinkShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongNoteShape"
		rf.Fieldname = "GongNoteLinkShapes"
		res = append(res, rf)
	case *GongNoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongNoteShapes"
		res = append(res, rf)
	case *GongStructShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongStructShapes"
		res = append(res, rf)
	case *LinkShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "LinkShapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (attributeshape *AttributeShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FieldTypeAsString",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Structname",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Fieldtypename",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (classdiagram *Classdiagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsIncludedInStaticWebSite",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongStructShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongStructShape",
		},
		{
			Name:                 "GongEnumShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongEnumShape",
		},
		{
			Name:                 "GongNoteShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongNoteShape",
		},
		{
			Name:               "ShowNbInstances",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ShowMultiplicity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ShowLinkNames",
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
			Name:               "NodeGongStructsIsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NodeGongStructNodeExpansion",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NodeGongEnumsIsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NodeGongEnumNodeExpansion",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NodeGongNotesIsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NodeGongNoteNodeExpansion",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (diagrampackage *DiagramPackage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Path",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "GongModelPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Classdiagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Classdiagram",
		},
		{
			Name:                 "SelectedClassdiagram",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Classdiagram",
		},
		{
			Name:               "AbsolutePathToDiagramPackage",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongenumshape *GongEnumShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
			Name:               "IdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongEnumValueShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongEnumValueShape",
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Identifier",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongnoteshape *GongNoteShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Identifier",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Body",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BodyHTML",
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
		{
			Name:               "Matched",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongNoteLinkShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongNoteLinkShape",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongstructshape *GongStructShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
			Name:               "IdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "AttributeShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributeShape",
		},
		{
			Name:                 "LinkShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkShape",
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (linkshape *LinkShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FieldTypeIdentifierMeta",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FieldOffsetX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FieldOffsetY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TargetMultiplicity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TargetMultiplicityOffsetX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TargetMultiplicityOffsetY",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "SourceMultiplicity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "SourceMultiplicityOffsetX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "SourceMultiplicityOffsetY",
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
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
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
func (attributeshape *AttributeShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = attributeshape.Name
	case "FieldTypeAsString":
		res.valueString = attributeshape.FieldTypeAsString
	case "Structname":
		res.valueString = attributeshape.Structname
	case "Fieldtypename":
		res.valueString = attributeshape.Fieldtypename
	}
	return
}
func (classdiagram *Classdiagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = classdiagram.Name
	case "Description":
		res.valueString = classdiagram.Description
	case "IsIncludedInStaticWebSite":
		res.valueString = fmt.Sprintf("%t", classdiagram.IsIncludedInStaticWebSite)
		res.valueBool = classdiagram.IsIncludedInStaticWebSite
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GongStructShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range classdiagram.GongStructShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "GongEnumShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range classdiagram.GongEnumShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "GongNoteShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range classdiagram.GongNoteShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "ShowNbInstances":
		res.valueString = fmt.Sprintf("%t", classdiagram.ShowNbInstances)
		res.valueBool = classdiagram.ShowNbInstances
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowMultiplicity":
		res.valueString = fmt.Sprintf("%t", classdiagram.ShowMultiplicity)
		res.valueBool = classdiagram.ShowMultiplicity
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowLinkNames":
		res.valueString = fmt.Sprintf("%t", classdiagram.ShowLinkNames)
		res.valueBool = classdiagram.ShowLinkNames
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInRenameMode":
		res.valueString = fmt.Sprintf("%t", classdiagram.IsInRenameMode)
		res.valueBool = classdiagram.IsInRenameMode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", classdiagram.IsExpanded)
		res.valueBool = classdiagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NodeGongStructsIsExpanded":
		res.valueString = fmt.Sprintf("%t", classdiagram.NodeGongStructsIsExpanded)
		res.valueBool = classdiagram.NodeGongStructsIsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NodeGongStructNodeExpansion":
		res.valueString = classdiagram.NodeGongStructNodeExpansion
	case "NodeGongEnumsIsExpanded":
		res.valueString = fmt.Sprintf("%t", classdiagram.NodeGongEnumsIsExpanded)
		res.valueBool = classdiagram.NodeGongEnumsIsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NodeGongEnumNodeExpansion":
		res.valueString = classdiagram.NodeGongEnumNodeExpansion
	case "NodeGongNotesIsExpanded":
		res.valueString = fmt.Sprintf("%t", classdiagram.NodeGongNotesIsExpanded)
		res.valueBool = classdiagram.NodeGongNotesIsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NodeGongNoteNodeExpansion":
		res.valueString = classdiagram.NodeGongNoteNodeExpansion
	}
	return
}
func (diagrampackage *DiagramPackage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagrampackage.Name
	case "Path":
		res.valueString = diagrampackage.Path
	case "GongModelPath":
		res.valueString = diagrampackage.GongModelPath
	case "Classdiagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagrampackage.Classdiagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "SelectedClassdiagram":
		res.GongFieldValueType = GongFieldValueTypePointer
		if diagrampackage.SelectedClassdiagram != nil {
			res.valueString = diagrampackage.SelectedClassdiagram.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, diagrampackage.SelectedClassdiagram))
		}
	case "AbsolutePathToDiagramPackage":
		res.valueString = diagrampackage.AbsolutePathToDiagramPackage
	}
	return
}
func (gongenumshape *GongEnumShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenumshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", gongenumshape.X)
		res.valueFloat = gongenumshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", gongenumshape.Y)
		res.valueFloat = gongenumshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "GongEnumValueShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongenumshape.GongEnumValueShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Width":
		res.valueString = fmt.Sprintf("%f", gongenumshape.Width)
		res.valueFloat = gongenumshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", gongenumshape.Height)
		res.valueFloat = gongenumshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", gongenumshape.IsExpanded)
		res.valueBool = gongenumshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (gongenumvalueshape *GongEnumValueShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenumvalueshape.Name
	}
	return
}
func (gongnotelinkshape *GongNoteLinkShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongnotelinkshape.Name
	case "Identifier":
		res.valueString = gongnotelinkshape.Identifier
	case "Type":
		enum := gongnotelinkshape.Type
		res.valueString = enum.ToCodeString()
	}
	return
}
func (gongnoteshape *GongNoteShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongnoteshape.Name
	case "Identifier":
		res.valueString = gongnoteshape.Identifier
	case "Body":
		res.valueString = gongnoteshape.Body
	case "BodyHTML":
		res.valueString = gongnoteshape.BodyHTML
	case "X":
		res.valueString = fmt.Sprintf("%f", gongnoteshape.X)
		res.valueFloat = gongnoteshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", gongnoteshape.Y)
		res.valueFloat = gongnoteshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", gongnoteshape.Width)
		res.valueFloat = gongnoteshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", gongnoteshape.Height)
		res.valueFloat = gongnoteshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Matched":
		res.valueString = fmt.Sprintf("%t", gongnoteshape.Matched)
		res.valueBool = gongnoteshape.Matched
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GongNoteLinkShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongnoteshape.GongNoteLinkShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", gongnoteshape.IsExpanded)
		res.valueBool = gongnoteshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (gongstructshape *GongStructShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongstructshape.Name
	case "X":
		res.valueString = fmt.Sprintf("%f", gongstructshape.X)
		res.valueFloat = gongstructshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", gongstructshape.Y)
		res.valueFloat = gongstructshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AttributeShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstructshape.AttributeShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "LinkShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstructshape.LinkShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Width":
		res.valueString = fmt.Sprintf("%f", gongstructshape.Width)
		res.valueFloat = gongstructshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", gongstructshape.Height)
		res.valueFloat = gongstructshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsSelected":
		res.valueString = fmt.Sprintf("%t", gongstructshape.IsSelected)
		res.valueBool = gongstructshape.IsSelected
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (linkshape *LinkShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linkshape.Name
	case "FieldOffsetX":
		res.valueString = fmt.Sprintf("%f", linkshape.FieldOffsetX)
		res.valueFloat = linkshape.FieldOffsetX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FieldOffsetY":
		res.valueString = fmt.Sprintf("%f", linkshape.FieldOffsetY)
		res.valueFloat = linkshape.FieldOffsetY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetMultiplicity":
		enum := linkshape.TargetMultiplicity
		res.valueString = enum.ToCodeString()
	case "TargetMultiplicityOffsetX":
		res.valueString = fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetX)
		res.valueFloat = linkshape.TargetMultiplicityOffsetX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TargetMultiplicityOffsetY":
		res.valueString = fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetY)
		res.valueFloat = linkshape.TargetMultiplicityOffsetY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SourceMultiplicity":
		enum := linkshape.SourceMultiplicity
		res.valueString = enum.ToCodeString()
	case "SourceMultiplicityOffsetX":
		res.valueString = fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetX)
		res.valueFloat = linkshape.SourceMultiplicityOffsetX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "SourceMultiplicityOffsetY":
		res.valueString = fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetY)
		res.valueFloat = linkshape.SourceMultiplicityOffsetY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "X":
		res.valueString = fmt.Sprintf("%f", linkshape.X)
		res.valueFloat = linkshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", linkshape.Y)
		res.valueFloat = linkshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := linkshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", linkshape.StartRatio)
		res.valueFloat = linkshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndOrientation":
		enum := linkshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", linkshape.EndRatio)
		res.valueFloat = linkshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", linkshape.CornerOffsetRatio)
		res.valueFloat = linkshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (attributeshape *AttributeShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		attributeshape.Name = value.GetValueString()
	case "FieldTypeAsString":
		attributeshape.FieldTypeAsString = value.GetValueString()
	case "Structname":
		attributeshape.Structname = value.GetValueString()
	case "Fieldtypename":
		attributeshape.Fieldtypename = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (classdiagram *Classdiagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		classdiagram.Name = value.GetValueString()
	case "Description":
		classdiagram.Description = value.GetValueString()
	case "IsIncludedInStaticWebSite":
		classdiagram.IsIncludedInStaticWebSite = value.GetValueBool()
	case "GongStructShapes":
		classdiagram.GongStructShapes = make([]*GongStructShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongStructShapes {
					if stage.GongStructShapeMap_Staged_Order[__instance__] == uint(id) {
						classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, __instance__)
						break
					}
				}
			}
		}
	case "GongEnumShapes":
		classdiagram.GongEnumShapes = make([]*GongEnumShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongEnumShapes {
					if stage.GongEnumShapeMap_Staged_Order[__instance__] == uint(id) {
						classdiagram.GongEnumShapes = append(classdiagram.GongEnumShapes, __instance__)
						break
					}
				}
			}
		}
	case "GongNoteShapes":
		classdiagram.GongNoteShapes = make([]*GongNoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongNoteShapes {
					if stage.GongNoteShapeMap_Staged_Order[__instance__] == uint(id) {
						classdiagram.GongNoteShapes = append(classdiagram.GongNoteShapes, __instance__)
						break
					}
				}
			}
		}
	case "ShowNbInstances":
		classdiagram.ShowNbInstances = value.GetValueBool()
	case "ShowMultiplicity":
		classdiagram.ShowMultiplicity = value.GetValueBool()
	case "ShowLinkNames":
		classdiagram.ShowLinkNames = value.GetValueBool()
	case "IsInRenameMode":
		classdiagram.IsInRenameMode = value.GetValueBool()
	case "IsExpanded":
		classdiagram.IsExpanded = value.GetValueBool()
	case "NodeGongStructsIsExpanded":
		classdiagram.NodeGongStructsIsExpanded = value.GetValueBool()
	case "NodeGongStructNodeExpansion":
		classdiagram.NodeGongStructNodeExpansion = value.GetValueString()
	case "NodeGongEnumsIsExpanded":
		classdiagram.NodeGongEnumsIsExpanded = value.GetValueBool()
	case "NodeGongEnumNodeExpansion":
		classdiagram.NodeGongEnumNodeExpansion = value.GetValueString()
	case "NodeGongNotesIsExpanded":
		classdiagram.NodeGongNotesIsExpanded = value.GetValueBool()
	case "NodeGongNoteNodeExpansion":
		classdiagram.NodeGongNoteNodeExpansion = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagrampackage *DiagramPackage) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagrampackage.Name = value.GetValueString()
	case "Path":
		diagrampackage.Path = value.GetValueString()
	case "GongModelPath":
		diagrampackage.GongModelPath = value.GetValueString()
	case "Classdiagrams":
		diagrampackage.Classdiagrams = make([]*Classdiagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Classdiagrams {
					if stage.ClassdiagramMap_Staged_Order[__instance__] == uint(id) {
						diagrampackage.Classdiagrams = append(diagrampackage.Classdiagrams, __instance__)
						break
					}
				}
			}
		}
	case "SelectedClassdiagram":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			diagrampackage.SelectedClassdiagram = nil
			for __instance__ := range stage.Classdiagrams {
				if stage.ClassdiagramMap_Staged_Order[__instance__] == uint(id) {
					diagrampackage.SelectedClassdiagram = __instance__
					break
				}
			}
		}
	case "AbsolutePathToDiagramPackage":
		diagrampackage.AbsolutePathToDiagramPackage = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongenumshape *GongEnumShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongenumshape.Name = value.GetValueString()
	case "X":
		gongenumshape.X = value.GetValueFloat()
	case "Y":
		gongenumshape.Y = value.GetValueFloat()
	case "GongEnumValueShapes":
		gongenumshape.GongEnumValueShapes = make([]*GongEnumValueShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongEnumValueShapes {
					if stage.GongEnumValueShapeMap_Staged_Order[__instance__] == uint(id) {
						gongenumshape.GongEnumValueShapes = append(gongenumshape.GongEnumValueShapes, __instance__)
						break
					}
				}
			}
		}
	case "Width":
		gongenumshape.Width = value.GetValueFloat()
	case "Height":
		gongenumshape.Height = value.GetValueFloat()
	case "IsExpanded":
		gongenumshape.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongenumvalueshape *GongEnumValueShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongenumvalueshape.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongnotelinkshape *GongNoteLinkShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongnotelinkshape.Name = value.GetValueString()
	case "Identifier":
		gongnotelinkshape.Identifier = value.GetValueString()
	case "Type":
		gongnotelinkshape.Type.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongnoteshape *GongNoteShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongnoteshape.Name = value.GetValueString()
	case "Identifier":
		gongnoteshape.Identifier = value.GetValueString()
	case "Body":
		gongnoteshape.Body = value.GetValueString()
	case "BodyHTML":
		gongnoteshape.BodyHTML = value.GetValueString()
	case "X":
		gongnoteshape.X = value.GetValueFloat()
	case "Y":
		gongnoteshape.Y = value.GetValueFloat()
	case "Width":
		gongnoteshape.Width = value.GetValueFloat()
	case "Height":
		gongnoteshape.Height = value.GetValueFloat()
	case "Matched":
		gongnoteshape.Matched = value.GetValueBool()
	case "GongNoteLinkShapes":
		gongnoteshape.GongNoteLinkShapes = make([]*GongNoteLinkShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongNoteLinkShapes {
					if stage.GongNoteLinkShapeMap_Staged_Order[__instance__] == uint(id) {
						gongnoteshape.GongNoteLinkShapes = append(gongnoteshape.GongNoteLinkShapes, __instance__)
						break
					}
				}
			}
		}
	case "IsExpanded":
		gongnoteshape.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongstructshape *GongStructShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongstructshape.Name = value.GetValueString()
	case "X":
		gongstructshape.X = value.GetValueFloat()
	case "Y":
		gongstructshape.Y = value.GetValueFloat()
	case "AttributeShapes":
		gongstructshape.AttributeShapes = make([]*AttributeShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AttributeShapes {
					if stage.AttributeShapeMap_Staged_Order[__instance__] == uint(id) {
						gongstructshape.AttributeShapes = append(gongstructshape.AttributeShapes, __instance__)
						break
					}
				}
			}
		}
	case "LinkShapes":
		gongstructshape.LinkShapes = make([]*LinkShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.LinkShapes {
					if stage.LinkShapeMap_Staged_Order[__instance__] == uint(id) {
						gongstructshape.LinkShapes = append(gongstructshape.LinkShapes, __instance__)
						break
					}
				}
			}
		}
	case "Width":
		gongstructshape.Width = value.GetValueFloat()
	case "Height":
		gongstructshape.Height = value.GetValueFloat()
	case "IsSelected":
		gongstructshape.IsSelected = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (linkshape *LinkShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		linkshape.Name = value.GetValueString()
	case "FieldOffsetX":
		linkshape.FieldOffsetX = value.GetValueFloat()
	case "FieldOffsetY":
		linkshape.FieldOffsetY = value.GetValueFloat()
	case "TargetMultiplicity":
		linkshape.TargetMultiplicity.FromCodeString(value.GetValueString())
	case "TargetMultiplicityOffsetX":
		linkshape.TargetMultiplicityOffsetX = value.GetValueFloat()
	case "TargetMultiplicityOffsetY":
		linkshape.TargetMultiplicityOffsetY = value.GetValueFloat()
	case "SourceMultiplicity":
		linkshape.SourceMultiplicity.FromCodeString(value.GetValueString())
	case "SourceMultiplicityOffsetX":
		linkshape.SourceMultiplicityOffsetX = value.GetValueFloat()
	case "SourceMultiplicityOffsetY":
		linkshape.SourceMultiplicityOffsetY = value.GetValueFloat()
	case "X":
		linkshape.X = value.GetValueFloat()
	case "Y":
		linkshape.Y = value.GetValueFloat()
	case "StartOrientation":
		linkshape.StartOrientation.FromCodeString(value.GetValueString())
	case "StartRatio":
		linkshape.StartRatio = value.GetValueFloat()
	case "EndOrientation":
		linkshape.EndOrientation.FromCodeString(value.GetValueString())
	case "EndRatio":
		linkshape.EndRatio = value.GetValueFloat()
	case "CornerOffsetRatio":
		linkshape.CornerOffsetRatio = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (attributeshape *AttributeShape) GongGetGongstructName() string {
	return "AttributeShape"
}

func (classdiagram *Classdiagram) GongGetGongstructName() string {
	return "Classdiagram"
}

func (diagrampackage *DiagramPackage) GongGetGongstructName() string {
	return "DiagramPackage"
}

func (gongenumshape *GongEnumShape) GongGetGongstructName() string {
	return "GongEnumShape"
}

func (gongenumvalueshape *GongEnumValueShape) GongGetGongstructName() string {
	return "GongEnumValueShape"
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetGongstructName() string {
	return "GongNoteLinkShape"
}

func (gongnoteshape *GongNoteShape) GongGetGongstructName() string {
	return "GongNoteShape"
}

func (gongstructshape *GongStructShape) GongGetGongstructName() string {
	return "GongStructShape"
}

func (linkshape *LinkShape) GongGetGongstructName() string {
	return "LinkShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.AttributeShapes_mapString = make(map[string]*AttributeShape)
	for attributeshape := range stage.AttributeShapes {
		stage.AttributeShapes_mapString[attributeshape.Name] = attributeshape
	}

	stage.Classdiagrams_mapString = make(map[string]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		stage.Classdiagrams_mapString[classdiagram.Name] = classdiagram
	}

	stage.DiagramPackages_mapString = make(map[string]*DiagramPackage)
	for diagrampackage := range stage.DiagramPackages {
		stage.DiagramPackages_mapString[diagrampackage.Name] = diagrampackage
	}

	stage.GongEnumShapes_mapString = make(map[string]*GongEnumShape)
	for gongenumshape := range stage.GongEnumShapes {
		stage.GongEnumShapes_mapString[gongenumshape.Name] = gongenumshape
	}

	stage.GongEnumValueShapes_mapString = make(map[string]*GongEnumValueShape)
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		stage.GongEnumValueShapes_mapString[gongenumvalueshape.Name] = gongenumvalueshape
	}

	stage.GongNoteLinkShapes_mapString = make(map[string]*GongNoteLinkShape)
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		stage.GongNoteLinkShapes_mapString[gongnotelinkshape.Name] = gongnotelinkshape
	}

	stage.GongNoteShapes_mapString = make(map[string]*GongNoteShape)
	for gongnoteshape := range stage.GongNoteShapes {
		stage.GongNoteShapes_mapString[gongnoteshape.Name] = gongnoteshape
	}

	stage.GongStructShapes_mapString = make(map[string]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		stage.GongStructShapes_mapString[gongstructshape.Name] = gongstructshape
	}

	stage.LinkShapes_mapString = make(map[string]*LinkShape)
	for linkshape := range stage.LinkShapes {
		stage.LinkShapes_mapString[linkshape.Name] = linkshape
	}

}
// Last line of the template
