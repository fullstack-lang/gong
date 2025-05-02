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
	"time"

	doc2_go "github.com/fullstack-lang/gong/lib/doc2/go"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

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

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Classdiagrams           map[*Classdiagram]any
	Classdiagrams_mapString map[string]*Classdiagram

	// insertion point for slice of pointers maps
	Classdiagram_GongStructShapes_reverseMap map[*GongStructShape]*Classdiagram

	Classdiagram_GongEnumShapes_reverseMap map[*GongEnumShape]*Classdiagram

	Classdiagram_NoteShapes_reverseMap map[*NoteShape]*Classdiagram

	OnAfterClassdiagramCreateCallback OnAfterCreateInterface[Classdiagram]
	OnAfterClassdiagramUpdateCallback OnAfterUpdateInterface[Classdiagram]
	OnAfterClassdiagramDeleteCallback OnAfterDeleteInterface[Classdiagram]
	OnAfterClassdiagramReadCallback   OnAfterReadInterface[Classdiagram]

	DiagramPackages           map[*DiagramPackage]any
	DiagramPackages_mapString map[string]*DiagramPackage

	// insertion point for slice of pointers maps
	DiagramPackage_Classdiagrams_reverseMap map[*Classdiagram]*DiagramPackage

	DiagramPackage_Umlscs_reverseMap map[*Umlsc]*DiagramPackage

	OnAfterDiagramPackageCreateCallback OnAfterCreateInterface[DiagramPackage]
	OnAfterDiagramPackageUpdateCallback OnAfterUpdateInterface[DiagramPackage]
	OnAfterDiagramPackageDeleteCallback OnAfterDeleteInterface[DiagramPackage]
	OnAfterDiagramPackageReadCallback   OnAfterReadInterface[DiagramPackage]

	Fields           map[*FieldShape]any
	Fields_mapString map[string]*FieldShape

	// insertion point for slice of pointers maps
	OnAfterFieldCreateCallback OnAfterCreateInterface[FieldShape]
	OnAfterFieldUpdateCallback OnAfterUpdateInterface[FieldShape]
	OnAfterFieldDeleteCallback OnAfterDeleteInterface[FieldShape]
	OnAfterFieldReadCallback   OnAfterReadInterface[FieldShape]

	GongEnumShapes           map[*GongEnumShape]any
	GongEnumShapes_mapString map[string]*GongEnumShape

	// insertion point for slice of pointers maps
	GongEnumShape_GongEnumValueEntrys_reverseMap map[*GongEnumValueEntry]*GongEnumShape

	OnAfterGongEnumShapeCreateCallback OnAfterCreateInterface[GongEnumShape]
	OnAfterGongEnumShapeUpdateCallback OnAfterUpdateInterface[GongEnumShape]
	OnAfterGongEnumShapeDeleteCallback OnAfterDeleteInterface[GongEnumShape]
	OnAfterGongEnumShapeReadCallback   OnAfterReadInterface[GongEnumShape]

	GongEnumValueEntrys           map[*GongEnumValueEntry]any
	GongEnumValueEntrys_mapString map[string]*GongEnumValueEntry

	// insertion point for slice of pointers maps
	OnAfterGongEnumValueEntryCreateCallback OnAfterCreateInterface[GongEnumValueEntry]
	OnAfterGongEnumValueEntryUpdateCallback OnAfterUpdateInterface[GongEnumValueEntry]
	OnAfterGongEnumValueEntryDeleteCallback OnAfterDeleteInterface[GongEnumValueEntry]
	OnAfterGongEnumValueEntryReadCallback   OnAfterReadInterface[GongEnumValueEntry]

	GongStructShapes           map[*GongStructShape]any
	GongStructShapes_mapString map[string]*GongStructShape

	// insertion point for slice of pointers maps
	GongStructShape_Fields_reverseMap map[*FieldShape]*GongStructShape

	GongStructShape_Links_reverseMap map[*Link]*GongStructShape

	OnAfterGongStructShapeCreateCallback OnAfterCreateInterface[GongStructShape]
	OnAfterGongStructShapeUpdateCallback OnAfterUpdateInterface[GongStructShape]
	OnAfterGongStructShapeDeleteCallback OnAfterDeleteInterface[GongStructShape]
	OnAfterGongStructShapeReadCallback   OnAfterReadInterface[GongStructShape]

	Links           map[*Link]any
	Links_mapString map[string]*Link

	// insertion point for slice of pointers maps
	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	NoteShapes           map[*NoteShape]any
	NoteShapes_mapString map[string]*NoteShape

	// insertion point for slice of pointers maps
	NoteShape_NoteShapeLinks_reverseMap map[*NoteShapeLink]*NoteShape

	OnAfterNoteShapeCreateCallback OnAfterCreateInterface[NoteShape]
	OnAfterNoteShapeUpdateCallback OnAfterUpdateInterface[NoteShape]
	OnAfterNoteShapeDeleteCallback OnAfterDeleteInterface[NoteShape]
	OnAfterNoteShapeReadCallback   OnAfterReadInterface[NoteShape]

	NoteShapeLinks           map[*NoteShapeLink]any
	NoteShapeLinks_mapString map[string]*NoteShapeLink

	// insertion point for slice of pointers maps
	OnAfterNoteShapeLinkCreateCallback OnAfterCreateInterface[NoteShapeLink]
	OnAfterNoteShapeLinkUpdateCallback OnAfterUpdateInterface[NoteShapeLink]
	OnAfterNoteShapeLinkDeleteCallback OnAfterDeleteInterface[NoteShapeLink]
	OnAfterNoteShapeLinkReadCallback   OnAfterReadInterface[NoteShapeLink]

	Positions           map[*Position]any
	Positions_mapString map[string]*Position

	// insertion point for slice of pointers maps
	OnAfterPositionCreateCallback OnAfterCreateInterface[Position]
	OnAfterPositionUpdateCallback OnAfterUpdateInterface[Position]
	OnAfterPositionDeleteCallback OnAfterDeleteInterface[Position]
	OnAfterPositionReadCallback   OnAfterReadInterface[Position]

	UmlStates           map[*UmlState]any
	UmlStates_mapString map[string]*UmlState

	// insertion point for slice of pointers maps
	OnAfterUmlStateCreateCallback OnAfterCreateInterface[UmlState]
	OnAfterUmlStateUpdateCallback OnAfterUpdateInterface[UmlState]
	OnAfterUmlStateDeleteCallback OnAfterDeleteInterface[UmlState]
	OnAfterUmlStateReadCallback   OnAfterReadInterface[UmlState]

	Umlscs           map[*Umlsc]any
	Umlscs_mapString map[string]*Umlsc

	// insertion point for slice of pointers maps
	Umlsc_States_reverseMap map[*UmlState]*Umlsc

	OnAfterUmlscCreateCallback OnAfterCreateInterface[Umlsc]
	OnAfterUmlscUpdateCallback OnAfterUpdateInterface[Umlsc]
	OnAfterUmlscDeleteCallback OnAfterDeleteInterface[Umlsc]
	OnAfterUmlscReadCallback   OnAfterReadInterface[Umlsc]

	Vertices           map[*Vertice]any
	Vertices_mapString map[string]*Vertice

	// insertion point for slice of pointers maps
	OnAfterVerticeCreateCallback OnAfterCreateInterface[Vertice]
	OnAfterVerticeUpdateCallback OnAfterUpdateInterface[Vertice]
	OnAfterVerticeDeleteCallback OnAfterDeleteInterface[Vertice]
	OnAfterVerticeReadCallback   OnAfterReadInterface[Vertice]

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
	ClassdiagramOrder            uint
	ClassdiagramMap_Staged_Order map[*Classdiagram]uint

	DiagramPackageOrder            uint
	DiagramPackageMap_Staged_Order map[*DiagramPackage]uint

	FieldOrder            uint
	FieldMap_Staged_Order map[*FieldShape]uint

	GongEnumShapeOrder            uint
	GongEnumShapeMap_Staged_Order map[*GongEnumShape]uint

	GongEnumValueEntryOrder            uint
	GongEnumValueEntryMap_Staged_Order map[*GongEnumValueEntry]uint

	GongStructShapeOrder            uint
	GongStructShapeMap_Staged_Order map[*GongStructShape]uint

	LinkOrder            uint
	LinkMap_Staged_Order map[*Link]uint

	NoteShapeOrder            uint
	NoteShapeMap_Staged_Order map[*NoteShape]uint

	NoteShapeLinkOrder            uint
	NoteShapeLinkMap_Staged_Order map[*NoteShapeLink]uint

	PositionOrder            uint
	PositionMap_Staged_Order map[*Position]uint

	UmlStateOrder            uint
	UmlStateMap_Staged_Order map[*UmlState]uint

	UmlscOrder            uint
	UmlscMap_Staged_Order map[*Umlsc]uint

	VerticeOrder            uint
	VerticeMap_Staged_Order map[*Vertice]uint

	// end of insertion point

	NamedStructs []*NamedStruct
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

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

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "Classdiagram":
		res = GetNamedStructInstances(stage.Classdiagrams, stage.ClassdiagramMap_Staged_Order)
	case "DiagramPackage":
		res = GetNamedStructInstances(stage.DiagramPackages, stage.DiagramPackageMap_Staged_Order)
	case "Field":
		res = GetNamedStructInstances(stage.Fields, stage.FieldMap_Staged_Order)
	case "GongEnumShape":
		res = GetNamedStructInstances(stage.GongEnumShapes, stage.GongEnumShapeMap_Staged_Order)
	case "GongEnumValueEntry":
		res = GetNamedStructInstances(stage.GongEnumValueEntrys, stage.GongEnumValueEntryMap_Staged_Order)
	case "GongStructShape":
		res = GetNamedStructInstances(stage.GongStructShapes, stage.GongStructShapeMap_Staged_Order)
	case "Link":
		res = GetNamedStructInstances(stage.Links, stage.LinkMap_Staged_Order)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShapeMap_Staged_Order)
	case "NoteShapeLink":
		res = GetNamedStructInstances(stage.NoteShapeLinks, stage.NoteShapeLinkMap_Staged_Order)
	case "Position":
		res = GetNamedStructInstances(stage.Positions, stage.PositionMap_Staged_Order)
	case "UmlState":
		res = GetNamedStructInstances(stage.UmlStates, stage.UmlStateMap_Staged_Order)
	case "Umlsc":
		res = GetNamedStructInstances(stage.Umlscs, stage.UmlscMap_Staged_Order)
	case "Vertice":
		res = GetNamedStructInstances(stage.Vertices, stage.VerticeMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/doc2/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return doc2_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return doc2_go.GoDiagramsDir
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
	CommitClassdiagram(classdiagram *Classdiagram)
	CheckoutClassdiagram(classdiagram *Classdiagram)
	CommitDiagramPackage(diagrampackage *DiagramPackage)
	CheckoutDiagramPackage(diagrampackage *DiagramPackage)
	CommitField(field *FieldShape)
	CheckoutField(field *FieldShape)
	CommitGongEnumShape(gongenumshape *GongEnumShape)
	CheckoutGongEnumShape(gongenumshape *GongEnumShape)
	CommitGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry)
	CheckoutGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry)
	CommitGongStructShape(gongstructshape *GongStructShape)
	CheckoutGongStructShape(gongstructshape *GongStructShape)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteShapeLink(noteshapelink *NoteShapeLink)
	CheckoutNoteShapeLink(noteshapelink *NoteShapeLink)
	CommitPosition(position *Position)
	CheckoutPosition(position *Position)
	CommitUmlState(umlstate *UmlState)
	CheckoutUmlState(umlstate *UmlState)
	CommitUmlsc(umlsc *Umlsc)
	CheckoutUmlsc(umlsc *Umlsc)
	CommitVertice(vertice *Vertice)
	CheckoutVertice(vertice *Vertice)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Classdiagrams:           make(map[*Classdiagram]any),
		Classdiagrams_mapString: make(map[string]*Classdiagram),

		DiagramPackages:           make(map[*DiagramPackage]any),
		DiagramPackages_mapString: make(map[string]*DiagramPackage),

		Fields:           make(map[*FieldShape]any),
		Fields_mapString: make(map[string]*FieldShape),

		GongEnumShapes:           make(map[*GongEnumShape]any),
		GongEnumShapes_mapString: make(map[string]*GongEnumShape),

		GongEnumValueEntrys:           make(map[*GongEnumValueEntry]any),
		GongEnumValueEntrys_mapString: make(map[string]*GongEnumValueEntry),

		GongStructShapes:           make(map[*GongStructShape]any),
		GongStructShapes_mapString: make(map[string]*GongStructShape),

		Links:           make(map[*Link]any),
		Links_mapString: make(map[string]*Link),

		NoteShapes:           make(map[*NoteShape]any),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteShapeLinks:           make(map[*NoteShapeLink]any),
		NoteShapeLinks_mapString: make(map[string]*NoteShapeLink),

		Positions:           make(map[*Position]any),
		Positions_mapString: make(map[string]*Position),

		UmlStates:           make(map[*UmlState]any),
		UmlStates_mapString: make(map[string]*UmlState),

		Umlscs:           make(map[*Umlsc]any),
		Umlscs_mapString: make(map[string]*Umlsc),

		Vertices:           make(map[*Vertice]any),
		Vertices_mapString: make(map[string]*Vertice),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ClassdiagramMap_Staged_Order: make(map[*Classdiagram]uint),

		DiagramPackageMap_Staged_Order: make(map[*DiagramPackage]uint),

		FieldMap_Staged_Order: make(map[*FieldShape]uint),

		GongEnumShapeMap_Staged_Order: make(map[*GongEnumShape]uint),

		GongEnumValueEntryMap_Staged_Order: make(map[*GongEnumValueEntry]uint),

		GongStructShapeMap_Staged_Order: make(map[*GongStructShape]uint),

		LinkMap_Staged_Order: make(map[*Link]uint),

		NoteShapeMap_Staged_Order: make(map[*NoteShape]uint),

		NoteShapeLinkMap_Staged_Order: make(map[*NoteShapeLink]uint),

		PositionMap_Staged_Order: make(map[*Position]uint),

		UmlStateMap_Staged_Order: make(map[*UmlState]uint),

		UmlscMap_Staged_Order: make(map[*Umlsc]uint),

		VerticeMap_Staged_Order: make(map[*Vertice]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			&NamedStruct{name: "Classdiagram"},
			&NamedStruct{name: "DiagramPackage"},
			&NamedStruct{name: "Field"},
			&NamedStruct{name: "GongEnumShape"},
			&NamedStruct{name: "GongEnumValueEntry"},
			&NamedStruct{name: "GongStructShape"},
			&NamedStruct{name: "Link"},
			&NamedStruct{name: "NoteShape"},
			&NamedStruct{name: "NoteShapeLink"},
			&NamedStruct{name: "Position"},
			&NamedStruct{name: "UmlState"},
			&NamedStruct{name: "Umlsc"},
			&NamedStruct{name: "Vertice"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Classdiagram:
		return stage.ClassdiagramMap_Staged_Order[instance]
	case *DiagramPackage:
		return stage.DiagramPackageMap_Staged_Order[instance]
	case *FieldShape:
		return stage.FieldMap_Staged_Order[instance]
	case *GongEnumShape:
		return stage.GongEnumShapeMap_Staged_Order[instance]
	case *GongEnumValueEntry:
		return stage.GongEnumValueEntryMap_Staged_Order[instance]
	case *GongStructShape:
		return stage.GongStructShapeMap_Staged_Order[instance]
	case *Link:
		return stage.LinkMap_Staged_Order[instance]
	case *NoteShape:
		return stage.NoteShapeMap_Staged_Order[instance]
	case *NoteShapeLink:
		return stage.NoteShapeLinkMap_Staged_Order[instance]
	case *Position:
		return stage.PositionMap_Staged_Order[instance]
	case *UmlState:
		return stage.UmlStateMap_Staged_Order[instance]
	case *Umlsc:
		return stage.UmlscMap_Staged_Order[instance]
	case *Vertice:
		return stage.VerticeMap_Staged_Order[instance]
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

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Classdiagram"] = len(stage.Classdiagrams)
	stage.Map_GongStructName_InstancesNb["DiagramPackage"] = len(stage.DiagramPackages)
	stage.Map_GongStructName_InstancesNb["Field"] = len(stage.Fields)
	stage.Map_GongStructName_InstancesNb["GongEnumShape"] = len(stage.GongEnumShapes)
	stage.Map_GongStructName_InstancesNb["GongEnumValueEntry"] = len(stage.GongEnumValueEntrys)
	stage.Map_GongStructName_InstancesNb["GongStructShape"] = len(stage.GongStructShapes)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteShapeLink"] = len(stage.NoteShapeLinks)
	stage.Map_GongStructName_InstancesNb["Position"] = len(stage.Positions)
	stage.Map_GongStructName_InstancesNb["UmlState"] = len(stage.UmlStates)
	stage.Map_GongStructName_InstancesNb["Umlsc"] = len(stage.Umlscs)
	stage.Map_GongStructName_InstancesNb["Vertice"] = len(stage.Vertices)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Classdiagram"] = len(stage.Classdiagrams)
	stage.Map_GongStructName_InstancesNb["DiagramPackage"] = len(stage.DiagramPackages)
	stage.Map_GongStructName_InstancesNb["Field"] = len(stage.Fields)
	stage.Map_GongStructName_InstancesNb["GongEnumShape"] = len(stage.GongEnumShapes)
	stage.Map_GongStructName_InstancesNb["GongEnumValueEntry"] = len(stage.GongEnumValueEntrys)
	stage.Map_GongStructName_InstancesNb["GongStructShape"] = len(stage.GongStructShapes)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteShapeLink"] = len(stage.NoteShapeLinks)
	stage.Map_GongStructName_InstancesNb["Position"] = len(stage.Positions)
	stage.Map_GongStructName_InstancesNb["UmlState"] = len(stage.UmlStates)
	stage.Map_GongStructName_InstancesNb["Umlsc"] = len(stage.Umlscs)
	stage.Map_GongStructName_InstancesNb["Vertice"] = len(stage.Vertices)

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
// Stage puts classdiagram to the model stage
func (classdiagram *Classdiagram) Stage(stage *Stage) *Classdiagram {

	if _, ok := stage.Classdiagrams[classdiagram]; !ok {
		stage.Classdiagrams[classdiagram] = __member
		stage.ClassdiagramMap_Staged_Order[classdiagram] = stage.ClassdiagramOrder
		stage.ClassdiagramOrder++
	}
	stage.Classdiagrams_mapString[classdiagram.Name] = classdiagram

	return classdiagram
}

// Unstage removes classdiagram off the model stage
func (classdiagram *Classdiagram) Unstage(stage *Stage) *Classdiagram {
	delete(stage.Classdiagrams, classdiagram)
	delete(stage.Classdiagrams_mapString, classdiagram.Name)
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

// Stage puts diagrampackage to the model stage
func (diagrampackage *DiagramPackage) Stage(stage *Stage) *DiagramPackage {

	if _, ok := stage.DiagramPackages[diagrampackage]; !ok {
		stage.DiagramPackages[diagrampackage] = __member
		stage.DiagramPackageMap_Staged_Order[diagrampackage] = stage.DiagramPackageOrder
		stage.DiagramPackageOrder++
	}
	stage.DiagramPackages_mapString[diagrampackage.Name] = diagrampackage

	return diagrampackage
}

// Unstage removes diagrampackage off the model stage
func (diagrampackage *DiagramPackage) Unstage(stage *Stage) *DiagramPackage {
	delete(stage.DiagramPackages, diagrampackage)
	delete(stage.DiagramPackages_mapString, diagrampackage.Name)
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

// Stage puts field to the model stage
func (field *FieldShape) Stage(stage *Stage) *FieldShape {

	if _, ok := stage.Fields[field]; !ok {
		stage.Fields[field] = __member
		stage.FieldMap_Staged_Order[field] = stage.FieldOrder
		stage.FieldOrder++
	}
	stage.Fields_mapString[field.Name] = field

	return field
}

// Unstage removes field off the model stage
func (field *FieldShape) Unstage(stage *Stage) *FieldShape {
	delete(stage.Fields, field)
	delete(stage.Fields_mapString, field.Name)
	return field
}

// UnstageVoid removes field off the model stage
func (field *FieldShape) UnstageVoid(stage *Stage) {
	delete(stage.Fields, field)
	delete(stage.Fields_mapString, field.Name)
}

// commit field to the back repo (if it is already staged)
func (field *FieldShape) Commit(stage *Stage) *FieldShape {
	if _, ok := stage.Fields[field]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitField(field)
		}
	}
	return field
}

func (field *FieldShape) CommitVoid(stage *Stage) {
	field.Commit(stage)
}

// Checkout field to the back repo (if it is already staged)
func (field *FieldShape) Checkout(stage *Stage) *FieldShape {
	if _, ok := stage.Fields[field]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutField(field)
		}
	}
	return field
}

// for satisfaction of GongStruct interface
func (field *FieldShape) GetName() (res string) {
	return field.Name
}

// Stage puts gongenumshape to the model stage
func (gongenumshape *GongEnumShape) Stage(stage *Stage) *GongEnumShape {

	if _, ok := stage.GongEnumShapes[gongenumshape]; !ok {
		stage.GongEnumShapes[gongenumshape] = __member
		stage.GongEnumShapeMap_Staged_Order[gongenumshape] = stage.GongEnumShapeOrder
		stage.GongEnumShapeOrder++
	}
	stage.GongEnumShapes_mapString[gongenumshape.Name] = gongenumshape

	return gongenumshape
}

// Unstage removes gongenumshape off the model stage
func (gongenumshape *GongEnumShape) Unstage(stage *Stage) *GongEnumShape {
	delete(stage.GongEnumShapes, gongenumshape)
	delete(stage.GongEnumShapes_mapString, gongenumshape.Name)
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

// Stage puts gongenumvalueentry to the model stage
func (gongenumvalueentry *GongEnumValueEntry) Stage(stage *Stage) *GongEnumValueEntry {

	if _, ok := stage.GongEnumValueEntrys[gongenumvalueentry]; !ok {
		stage.GongEnumValueEntrys[gongenumvalueentry] = __member
		stage.GongEnumValueEntryMap_Staged_Order[gongenumvalueentry] = stage.GongEnumValueEntryOrder
		stage.GongEnumValueEntryOrder++
	}
	stage.GongEnumValueEntrys_mapString[gongenumvalueentry.Name] = gongenumvalueentry

	return gongenumvalueentry
}

// Unstage removes gongenumvalueentry off the model stage
func (gongenumvalueentry *GongEnumValueEntry) Unstage(stage *Stage) *GongEnumValueEntry {
	delete(stage.GongEnumValueEntrys, gongenumvalueentry)
	delete(stage.GongEnumValueEntrys_mapString, gongenumvalueentry.Name)
	return gongenumvalueentry
}

// UnstageVoid removes gongenumvalueentry off the model stage
func (gongenumvalueentry *GongEnumValueEntry) UnstageVoid(stage *Stage) {
	delete(stage.GongEnumValueEntrys, gongenumvalueentry)
	delete(stage.GongEnumValueEntrys_mapString, gongenumvalueentry.Name)
}

// commit gongenumvalueentry to the back repo (if it is already staged)
func (gongenumvalueentry *GongEnumValueEntry) Commit(stage *Stage) *GongEnumValueEntry {
	if _, ok := stage.GongEnumValueEntrys[gongenumvalueentry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnumValueEntry(gongenumvalueentry)
		}
	}
	return gongenumvalueentry
}

func (gongenumvalueentry *GongEnumValueEntry) CommitVoid(stage *Stage) {
	gongenumvalueentry.Commit(stage)
}

// Checkout gongenumvalueentry to the back repo (if it is already staged)
func (gongenumvalueentry *GongEnumValueEntry) Checkout(stage *Stage) *GongEnumValueEntry {
	if _, ok := stage.GongEnumValueEntrys[gongenumvalueentry]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnumValueEntry(gongenumvalueentry)
		}
	}
	return gongenumvalueentry
}

// for satisfaction of GongStruct interface
func (gongenumvalueentry *GongEnumValueEntry) GetName() (res string) {
	return gongenumvalueentry.Name
}

// Stage puts gongstructshape to the model stage
func (gongstructshape *GongStructShape) Stage(stage *Stage) *GongStructShape {

	if _, ok := stage.GongStructShapes[gongstructshape]; !ok {
		stage.GongStructShapes[gongstructshape] = __member
		stage.GongStructShapeMap_Staged_Order[gongstructshape] = stage.GongStructShapeOrder
		stage.GongStructShapeOrder++
	}
	stage.GongStructShapes_mapString[gongstructshape.Name] = gongstructshape

	return gongstructshape
}

// Unstage removes gongstructshape off the model stage
func (gongstructshape *GongStructShape) Unstage(stage *Stage) *GongStructShape {
	delete(stage.GongStructShapes, gongstructshape)
	delete(stage.GongStructShapes_mapString, gongstructshape.Name)
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

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {

	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = __member
		stage.LinkMap_Staged_Order[link] = stage.LinkOrder
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *Stage) {
	link.Commit(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLink(link)
		}
	}
	return link
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// Stage puts noteshape to the model stage
func (noteshape *NoteShape) Stage(stage *Stage) *NoteShape {

	if _, ok := stage.NoteShapes[noteshape]; !ok {
		stage.NoteShapes[noteshape] = __member
		stage.NoteShapeMap_Staged_Order[noteshape] = stage.NoteShapeOrder
		stage.NoteShapeOrder++
	}
	stage.NoteShapes_mapString[noteshape.Name] = noteshape

	return noteshape
}

// Unstage removes noteshape off the model stage
func (noteshape *NoteShape) Unstage(stage *Stage) *NoteShape {
	delete(stage.NoteShapes, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)
	return noteshape
}

// UnstageVoid removes noteshape off the model stage
func (noteshape *NoteShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteShapes, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)
}

// commit noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Commit(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteShape(noteshape)
		}
	}
	return noteshape
}

func (noteshape *NoteShape) CommitVoid(stage *Stage) {
	noteshape.Commit(stage)
}

// Checkout noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Checkout(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteShape(noteshape)
		}
	}
	return noteshape
}

// for satisfaction of GongStruct interface
func (noteshape *NoteShape) GetName() (res string) {
	return noteshape.Name
}

// Stage puts noteshapelink to the model stage
func (noteshapelink *NoteShapeLink) Stage(stage *Stage) *NoteShapeLink {

	if _, ok := stage.NoteShapeLinks[noteshapelink]; !ok {
		stage.NoteShapeLinks[noteshapelink] = __member
		stage.NoteShapeLinkMap_Staged_Order[noteshapelink] = stage.NoteShapeLinkOrder
		stage.NoteShapeLinkOrder++
	}
	stage.NoteShapeLinks_mapString[noteshapelink.Name] = noteshapelink

	return noteshapelink
}

// Unstage removes noteshapelink off the model stage
func (noteshapelink *NoteShapeLink) Unstage(stage *Stage) *NoteShapeLink {
	delete(stage.NoteShapeLinks, noteshapelink)
	delete(stage.NoteShapeLinks_mapString, noteshapelink.Name)
	return noteshapelink
}

// UnstageVoid removes noteshapelink off the model stage
func (noteshapelink *NoteShapeLink) UnstageVoid(stage *Stage) {
	delete(stage.NoteShapeLinks, noteshapelink)
	delete(stage.NoteShapeLinks_mapString, noteshapelink.Name)
}

// commit noteshapelink to the back repo (if it is already staged)
func (noteshapelink *NoteShapeLink) Commit(stage *Stage) *NoteShapeLink {
	if _, ok := stage.NoteShapeLinks[noteshapelink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteShapeLink(noteshapelink)
		}
	}
	return noteshapelink
}

func (noteshapelink *NoteShapeLink) CommitVoid(stage *Stage) {
	noteshapelink.Commit(stage)
}

// Checkout noteshapelink to the back repo (if it is already staged)
func (noteshapelink *NoteShapeLink) Checkout(stage *Stage) *NoteShapeLink {
	if _, ok := stage.NoteShapeLinks[noteshapelink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteShapeLink(noteshapelink)
		}
	}
	return noteshapelink
}

// for satisfaction of GongStruct interface
func (noteshapelink *NoteShapeLink) GetName() (res string) {
	return noteshapelink.Name
}

// Stage puts position to the model stage
func (position *Position) Stage(stage *Stage) *Position {

	if _, ok := stage.Positions[position]; !ok {
		stage.Positions[position] = __member
		stage.PositionMap_Staged_Order[position] = stage.PositionOrder
		stage.PositionOrder++
	}
	stage.Positions_mapString[position.Name] = position

	return position
}

// Unstage removes position off the model stage
func (position *Position) Unstage(stage *Stage) *Position {
	delete(stage.Positions, position)
	delete(stage.Positions_mapString, position.Name)
	return position
}

// UnstageVoid removes position off the model stage
func (position *Position) UnstageVoid(stage *Stage) {
	delete(stage.Positions, position)
	delete(stage.Positions_mapString, position.Name)
}

// commit position to the back repo (if it is already staged)
func (position *Position) Commit(stage *Stage) *Position {
	if _, ok := stage.Positions[position]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPosition(position)
		}
	}
	return position
}

func (position *Position) CommitVoid(stage *Stage) {
	position.Commit(stage)
}

// Checkout position to the back repo (if it is already staged)
func (position *Position) Checkout(stage *Stage) *Position {
	if _, ok := stage.Positions[position]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPosition(position)
		}
	}
	return position
}

// for satisfaction of GongStruct interface
func (position *Position) GetName() (res string) {
	return position.Name
}

// Stage puts umlstate to the model stage
func (umlstate *UmlState) Stage(stage *Stage) *UmlState {

	if _, ok := stage.UmlStates[umlstate]; !ok {
		stage.UmlStates[umlstate] = __member
		stage.UmlStateMap_Staged_Order[umlstate] = stage.UmlStateOrder
		stage.UmlStateOrder++
	}
	stage.UmlStates_mapString[umlstate.Name] = umlstate

	return umlstate
}

// Unstage removes umlstate off the model stage
func (umlstate *UmlState) Unstage(stage *Stage) *UmlState {
	delete(stage.UmlStates, umlstate)
	delete(stage.UmlStates_mapString, umlstate.Name)
	return umlstate
}

// UnstageVoid removes umlstate off the model stage
func (umlstate *UmlState) UnstageVoid(stage *Stage) {
	delete(stage.UmlStates, umlstate)
	delete(stage.UmlStates_mapString, umlstate.Name)
}

// commit umlstate to the back repo (if it is already staged)
func (umlstate *UmlState) Commit(stage *Stage) *UmlState {
	if _, ok := stage.UmlStates[umlstate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUmlState(umlstate)
		}
	}
	return umlstate
}

func (umlstate *UmlState) CommitVoid(stage *Stage) {
	umlstate.Commit(stage)
}

// Checkout umlstate to the back repo (if it is already staged)
func (umlstate *UmlState) Checkout(stage *Stage) *UmlState {
	if _, ok := stage.UmlStates[umlstate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUmlState(umlstate)
		}
	}
	return umlstate
}

// for satisfaction of GongStruct interface
func (umlstate *UmlState) GetName() (res string) {
	return umlstate.Name
}

// Stage puts umlsc to the model stage
func (umlsc *Umlsc) Stage(stage *Stage) *Umlsc {

	if _, ok := stage.Umlscs[umlsc]; !ok {
		stage.Umlscs[umlsc] = __member
		stage.UmlscMap_Staged_Order[umlsc] = stage.UmlscOrder
		stage.UmlscOrder++
	}
	stage.Umlscs_mapString[umlsc.Name] = umlsc

	return umlsc
}

// Unstage removes umlsc off the model stage
func (umlsc *Umlsc) Unstage(stage *Stage) *Umlsc {
	delete(stage.Umlscs, umlsc)
	delete(stage.Umlscs_mapString, umlsc.Name)
	return umlsc
}

// UnstageVoid removes umlsc off the model stage
func (umlsc *Umlsc) UnstageVoid(stage *Stage) {
	delete(stage.Umlscs, umlsc)
	delete(stage.Umlscs_mapString, umlsc.Name)
}

// commit umlsc to the back repo (if it is already staged)
func (umlsc *Umlsc) Commit(stage *Stage) *Umlsc {
	if _, ok := stage.Umlscs[umlsc]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUmlsc(umlsc)
		}
	}
	return umlsc
}

func (umlsc *Umlsc) CommitVoid(stage *Stage) {
	umlsc.Commit(stage)
}

// Checkout umlsc to the back repo (if it is already staged)
func (umlsc *Umlsc) Checkout(stage *Stage) *Umlsc {
	if _, ok := stage.Umlscs[umlsc]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUmlsc(umlsc)
		}
	}
	return umlsc
}

// for satisfaction of GongStruct interface
func (umlsc *Umlsc) GetName() (res string) {
	return umlsc.Name
}

// Stage puts vertice to the model stage
func (vertice *Vertice) Stage(stage *Stage) *Vertice {

	if _, ok := stage.Vertices[vertice]; !ok {
		stage.Vertices[vertice] = __member
		stage.VerticeMap_Staged_Order[vertice] = stage.VerticeOrder
		stage.VerticeOrder++
	}
	stage.Vertices_mapString[vertice.Name] = vertice

	return vertice
}

// Unstage removes vertice off the model stage
func (vertice *Vertice) Unstage(stage *Stage) *Vertice {
	delete(stage.Vertices, vertice)
	delete(stage.Vertices_mapString, vertice.Name)
	return vertice
}

// UnstageVoid removes vertice off the model stage
func (vertice *Vertice) UnstageVoid(stage *Stage) {
	delete(stage.Vertices, vertice)
	delete(stage.Vertices_mapString, vertice.Name)
}

// commit vertice to the back repo (if it is already staged)
func (vertice *Vertice) Commit(stage *Stage) *Vertice {
	if _, ok := stage.Vertices[vertice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitVertice(vertice)
		}
	}
	return vertice
}

func (vertice *Vertice) CommitVoid(stage *Stage) {
	vertice.Commit(stage)
}

// Checkout vertice to the back repo (if it is already staged)
func (vertice *Vertice) Checkout(stage *Stage) *Vertice {
	if _, ok := stage.Vertices[vertice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutVertice(vertice)
		}
	}
	return vertice
}

// for satisfaction of GongStruct interface
func (vertice *Vertice) GetName() (res string) {
	return vertice.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMClassdiagram(Classdiagram *Classdiagram)
	CreateORMDiagramPackage(DiagramPackage *DiagramPackage)
	CreateORMField(Field *FieldShape)
	CreateORMGongEnumShape(GongEnumShape *GongEnumShape)
	CreateORMGongEnumValueEntry(GongEnumValueEntry *GongEnumValueEntry)
	CreateORMGongStructShape(GongStructShape *GongStructShape)
	CreateORMLink(Link *Link)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteShapeLink(NoteShapeLink *NoteShapeLink)
	CreateORMPosition(Position *Position)
	CreateORMUmlState(UmlState *UmlState)
	CreateORMUmlsc(Umlsc *Umlsc)
	CreateORMVertice(Vertice *Vertice)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMClassdiagram(Classdiagram *Classdiagram)
	DeleteORMDiagramPackage(DiagramPackage *DiagramPackage)
	DeleteORMField(Field *FieldShape)
	DeleteORMGongEnumShape(GongEnumShape *GongEnumShape)
	DeleteORMGongEnumValueEntry(GongEnumValueEntry *GongEnumValueEntry)
	DeleteORMGongStructShape(GongStructShape *GongStructShape)
	DeleteORMLink(Link *Link)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteShapeLink(NoteShapeLink *NoteShapeLink)
	DeleteORMPosition(Position *Position)
	DeleteORMUmlState(UmlState *UmlState)
	DeleteORMUmlsc(Umlsc *Umlsc)
	DeleteORMVertice(Vertice *Vertice)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Classdiagrams = make(map[*Classdiagram]any)
	stage.Classdiagrams_mapString = make(map[string]*Classdiagram)
	stage.ClassdiagramMap_Staged_Order = make(map[*Classdiagram]uint)
	stage.ClassdiagramOrder = 0

	stage.DiagramPackages = make(map[*DiagramPackage]any)
	stage.DiagramPackages_mapString = make(map[string]*DiagramPackage)
	stage.DiagramPackageMap_Staged_Order = make(map[*DiagramPackage]uint)
	stage.DiagramPackageOrder = 0

	stage.Fields = make(map[*FieldShape]any)
	stage.Fields_mapString = make(map[string]*FieldShape)
	stage.FieldMap_Staged_Order = make(map[*FieldShape]uint)
	stage.FieldOrder = 0

	stage.GongEnumShapes = make(map[*GongEnumShape]any)
	stage.GongEnumShapes_mapString = make(map[string]*GongEnumShape)
	stage.GongEnumShapeMap_Staged_Order = make(map[*GongEnumShape]uint)
	stage.GongEnumShapeOrder = 0

	stage.GongEnumValueEntrys = make(map[*GongEnumValueEntry]any)
	stage.GongEnumValueEntrys_mapString = make(map[string]*GongEnumValueEntry)
	stage.GongEnumValueEntryMap_Staged_Order = make(map[*GongEnumValueEntry]uint)
	stage.GongEnumValueEntryOrder = 0

	stage.GongStructShapes = make(map[*GongStructShape]any)
	stage.GongStructShapes_mapString = make(map[string]*GongStructShape)
	stage.GongStructShapeMap_Staged_Order = make(map[*GongStructShape]uint)
	stage.GongStructShapeOrder = 0

	stage.Links = make(map[*Link]any)
	stage.Links_mapString = make(map[string]*Link)
	stage.LinkMap_Staged_Order = make(map[*Link]uint)
	stage.LinkOrder = 0

	stage.NoteShapes = make(map[*NoteShape]any)
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShapeMap_Staged_Order = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteShapeLinks = make(map[*NoteShapeLink]any)
	stage.NoteShapeLinks_mapString = make(map[string]*NoteShapeLink)
	stage.NoteShapeLinkMap_Staged_Order = make(map[*NoteShapeLink]uint)
	stage.NoteShapeLinkOrder = 0

	stage.Positions = make(map[*Position]any)
	stage.Positions_mapString = make(map[string]*Position)
	stage.PositionMap_Staged_Order = make(map[*Position]uint)
	stage.PositionOrder = 0

	stage.UmlStates = make(map[*UmlState]any)
	stage.UmlStates_mapString = make(map[string]*UmlState)
	stage.UmlStateMap_Staged_Order = make(map[*UmlState]uint)
	stage.UmlStateOrder = 0

	stage.Umlscs = make(map[*Umlsc]any)
	stage.Umlscs_mapString = make(map[string]*Umlsc)
	stage.UmlscMap_Staged_Order = make(map[*Umlsc]uint)
	stage.UmlscOrder = 0

	stage.Vertices = make(map[*Vertice]any)
	stage.Vertices_mapString = make(map[string]*Vertice)
	stage.VerticeMap_Staged_Order = make(map[*Vertice]uint)
	stage.VerticeOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Classdiagrams = nil
	stage.Classdiagrams_mapString = nil

	stage.DiagramPackages = nil
	stage.DiagramPackages_mapString = nil

	stage.Fields = nil
	stage.Fields_mapString = nil

	stage.GongEnumShapes = nil
	stage.GongEnumShapes_mapString = nil

	stage.GongEnumValueEntrys = nil
	stage.GongEnumValueEntrys_mapString = nil

	stage.GongStructShapes = nil
	stage.GongStructShapes_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteShapeLinks = nil
	stage.NoteShapeLinks_mapString = nil

	stage.Positions = nil
	stage.Positions_mapString = nil

	stage.UmlStates = nil
	stage.UmlStates_mapString = nil

	stage.Umlscs = nil
	stage.Umlscs_mapString = nil

	stage.Vertices = nil
	stage.Vertices_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for classdiagram := range stage.Classdiagrams {
		classdiagram.Unstage(stage)
	}

	for diagrampackage := range stage.DiagramPackages {
		diagrampackage.Unstage(stage)
	}

	for field := range stage.Fields {
		field.Unstage(stage)
	}

	for gongenumshape := range stage.GongEnumShapes {
		gongenumshape.Unstage(stage)
	}

	for gongenumvalueentry := range stage.GongEnumValueEntrys {
		gongenumvalueentry.Unstage(stage)
	}

	for gongstructshape := range stage.GongStructShapes {
		gongstructshape.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for noteshapelink := range stage.NoteShapeLinks {
		noteshapelink.Unstage(stage)
	}

	for position := range stage.Positions {
		position.Unstage(stage)
	}

	for umlstate := range stage.UmlStates {
		umlstate.Unstage(stage)
	}

	for umlsc := range stage.Umlscs {
		umlsc.Unstage(stage)
	}

	for vertice := range stage.Vertices {
		vertice.Unstage(stage)
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
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

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
	case map[*Classdiagram]any:
		return any(&stage.Classdiagrams).(*Type)
	case map[*DiagramPackage]any:
		return any(&stage.DiagramPackages).(*Type)
	case map[*FieldShape]any:
		return any(&stage.Fields).(*Type)
	case map[*GongEnumShape]any:
		return any(&stage.GongEnumShapes).(*Type)
	case map[*GongEnumValueEntry]any:
		return any(&stage.GongEnumValueEntrys).(*Type)
	case map[*GongStructShape]any:
		return any(&stage.GongStructShapes).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteShapeLink]any:
		return any(&stage.NoteShapeLinks).(*Type)
	case map[*Position]any:
		return any(&stage.Positions).(*Type)
	case map[*UmlState]any:
		return any(&stage.UmlStates).(*Type)
	case map[*Umlsc]any:
		return any(&stage.Umlscs).(*Type)
	case map[*Vertice]any:
		return any(&stage.Vertices).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Classdiagram:
		return any(&stage.Classdiagrams_mapString).(*Type)
	case map[string]*DiagramPackage:
		return any(&stage.DiagramPackages_mapString).(*Type)
	case map[string]*FieldShape:
		return any(&stage.Fields_mapString).(*Type)
	case map[string]*GongEnumShape:
		return any(&stage.GongEnumShapes_mapString).(*Type)
	case map[string]*GongEnumValueEntry:
		return any(&stage.GongEnumValueEntrys_mapString).(*Type)
	case map[string]*GongStructShape:
		return any(&stage.GongStructShapes_mapString).(*Type)
	case map[string]*Link:
		return any(&stage.Links_mapString).(*Type)
	case map[string]*NoteShape:
		return any(&stage.NoteShapes_mapString).(*Type)
	case map[string]*NoteShapeLink:
		return any(&stage.NoteShapeLinks_mapString).(*Type)
	case map[string]*Position:
		return any(&stage.Positions_mapString).(*Type)
	case map[string]*UmlState:
		return any(&stage.UmlStates_mapString).(*Type)
	case map[string]*Umlsc:
		return any(&stage.Umlscs_mapString).(*Type)
	case map[string]*Vertice:
		return any(&stage.Vertices_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Classdiagram:
		return any(&stage.Classdiagrams).(*map[*Type]any)
	case DiagramPackage:
		return any(&stage.DiagramPackages).(*map[*Type]any)
	case FieldShape:
		return any(&stage.Fields).(*map[*Type]any)
	case GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[*Type]any)
	case GongEnumValueEntry:
		return any(&stage.GongEnumValueEntrys).(*map[*Type]any)
	case GongStructShape:
		return any(&stage.GongStructShapes).(*map[*Type]any)
	case Link:
		return any(&stage.Links).(*map[*Type]any)
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]any)
	case NoteShapeLink:
		return any(&stage.NoteShapeLinks).(*map[*Type]any)
	case Position:
		return any(&stage.Positions).(*map[*Type]any)
	case UmlState:
		return any(&stage.UmlStates).(*map[*Type]any)
	case Umlsc:
		return any(&stage.Umlscs).(*map[*Type]any)
	case Vertice:
		return any(&stage.Vertices).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Classdiagram:
		return any(&stage.Classdiagrams).(*map[Type]any)
	case *DiagramPackage:
		return any(&stage.DiagramPackages).(*map[Type]any)
	case *FieldShape:
		return any(&stage.Fields).(*map[Type]any)
	case *GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[Type]any)
	case *GongEnumValueEntry:
		return any(&stage.GongEnumValueEntrys).(*map[Type]any)
	case *GongStructShape:
		return any(&stage.GongStructShapes).(*map[Type]any)
	case *Link:
		return any(&stage.Links).(*map[Type]any)
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]any)
	case *NoteShapeLink:
		return any(&stage.NoteShapeLinks).(*map[Type]any)
	case *Position:
		return any(&stage.Positions).(*map[Type]any)
	case *UmlState:
		return any(&stage.UmlStates).(*map[Type]any)
	case *Umlsc:
		return any(&stage.Umlscs).(*map[Type]any)
	case *Vertice:
		return any(&stage.Vertices).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Classdiagram:
		return any(&stage.Classdiagrams_mapString).(*map[string]*Type)
	case DiagramPackage:
		return any(&stage.DiagramPackages_mapString).(*map[string]*Type)
	case FieldShape:
		return any(&stage.Fields_mapString).(*map[string]*Type)
	case GongEnumShape:
		return any(&stage.GongEnumShapes_mapString).(*map[string]*Type)
	case GongEnumValueEntry:
		return any(&stage.GongEnumValueEntrys_mapString).(*map[string]*Type)
	case GongStructShape:
		return any(&stage.GongStructShapes_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteShapeLink:
		return any(&stage.NoteShapeLinks_mapString).(*map[string]*Type)
	case Position:
		return any(&stage.Positions_mapString).(*map[string]*Type)
	case UmlState:
		return any(&stage.UmlStates_mapString).(*map[string]*Type)
	case Umlsc:
		return any(&stage.Umlscs_mapString).(*map[string]*Type)
	case Vertice:
		return any(&stage.Vertices_mapString).(*map[string]*Type)
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
	case Classdiagram:
		return any(&Classdiagram{
			// Initialisation of associations
			// field is initialized with an instance of GongStructShape with the name of the field
			GongStructShapes: []*GongStructShape{{Name: "GongStructShapes"}},
			// field is initialized with an instance of GongEnumShape with the name of the field
			GongEnumShapes: []*GongEnumShape{{Name: "GongEnumShapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			NoteShapes: []*NoteShape{{Name: "NoteShapes"}},
		}).(*Type)
	case DiagramPackage:
		return any(&DiagramPackage{
			// Initialisation of associations
			// field is initialized with an instance of Classdiagram with the name of the field
			Classdiagrams: []*Classdiagram{{Name: "Classdiagrams"}},
			// field is initialized with an instance of Classdiagram with the name of the field
			SelectedClassdiagram: &Classdiagram{Name: "SelectedClassdiagram"},
			// field is initialized with an instance of Umlsc with the name of the field
			Umlscs: []*Umlsc{{Name: "Umlscs"}},
		}).(*Type)
	case FieldShape:
		return any(&FieldShape{
			// Initialisation of associations
		}).(*Type)
	case GongEnumShape:
		return any(&GongEnumShape{
			// Initialisation of associations
			// field is initialized with an instance of Position with the name of the field
			Position: &Position{Name: "Position"},
			// field is initialized with an instance of GongEnumValueEntry with the name of the field
			GongEnumValueEntrys: []*GongEnumValueEntry{{Name: "GongEnumValueEntrys"}},
		}).(*Type)
	case GongEnumValueEntry:
		return any(&GongEnumValueEntry{
			// Initialisation of associations
		}).(*Type)
	case GongStructShape:
		return any(&GongStructShape{
			// Initialisation of associations
			// field is initialized with an instance of Position with the name of the field
			Position: &Position{Name: "Position"},
			// field is initialized with an instance of Field with the name of the field
			FieldShapes: []*FieldShape{{Name: "Fields"}},
			// field is initialized with an instance of Link with the name of the field
			Links: []*Link{{Name: "Links"}},
		}).(*Type)
	case Link:
		return any(&Link{
			// Initialisation of associations
			// field is initialized with an instance of Vertice with the name of the field
			Middlevertice: &Vertice{Name: "Middlevertice"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of NoteShapeLink with the name of the field
			NoteShapeLinks: []*NoteShapeLink{{Name: "NoteShapeLinks"}},
		}).(*Type)
	case NoteShapeLink:
		return any(&NoteShapeLink{
			// Initialisation of associations
		}).(*Type)
	case Position:
		return any(&Position{
			// Initialisation of associations
		}).(*Type)
	case UmlState:
		return any(&UmlState{
			// Initialisation of associations
		}).(*Type)
	case Umlsc:
		return any(&Umlsc{
			// Initialisation of associations
			// field is initialized with an instance of UmlState with the name of the field
			States: []*UmlState{{Name: "States"}},
		}).(*Type)
	case Vertice:
		return any(&Vertice{
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
	// reverse maps of direct associations of Field
	case FieldShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumShape
	case GongEnumShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Position":
			res := make(map[*Position][]*GongEnumShape)
			for gongenumshape := range stage.GongEnumShapes {
				if gongenumshape.Position != nil {
					position_ := gongenumshape.Position
					var gongenumshapes []*GongEnumShape
					_, ok := res[position_]
					if ok {
						gongenumshapes = res[position_]
					} else {
						gongenumshapes = make([]*GongEnumShape, 0)
					}
					gongenumshapes = append(gongenumshapes, gongenumshape)
					res[position_] = gongenumshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumValueEntry
	case GongEnumValueEntry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStructShape
	case GongStructShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Position":
			res := make(map[*Position][]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				if gongstructshape.Position != nil {
					position_ := gongstructshape.Position
					var gongstructshapes []*GongStructShape
					_, ok := res[position_]
					if ok {
						gongstructshapes = res[position_]
					} else {
						gongstructshapes = make([]*GongStructShape, 0)
					}
					gongstructshapes = append(gongstructshapes, gongstructshape)
					res[position_] = gongstructshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		case "Middlevertice":
			res := make(map[*Vertice][]*Link)
			for link := range stage.Links {
				if link.Middlevertice != nil {
					vertice_ := link.Middlevertice
					var links []*Link
					_, ok := res[vertice_]
					if ok {
						links = res[vertice_]
					} else {
						links = make([]*Link, 0)
					}
					links = append(links, link)
					res[vertice_] = links
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShapeLink
	case NoteShapeLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Position
	case Position:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UmlState
	case UmlState:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Umlsc
	case Umlsc:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Vertice
	case Vertice:
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Classdiagram
	case Classdiagram:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStructShapes":
			res := make(map[*GongStructShape]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, gongstructshape_ := range classdiagram.GongStructShapes {
					res[gongstructshape_] = classdiagram
				}
			}
			return any(res).(map[*End]*Start)
		case "GongEnumShapes":
			res := make(map[*GongEnumShape]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, gongenumshape_ := range classdiagram.GongEnumShapes {
					res[gongenumshape_] = classdiagram
				}
			}
			return any(res).(map[*End]*Start)
		case "NoteShapes":
			res := make(map[*NoteShape]*Classdiagram)
			for classdiagram := range stage.Classdiagrams {
				for _, noteshape_ := range classdiagram.NoteShapes {
					res[noteshape_] = classdiagram
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DiagramPackage
	case DiagramPackage:
		switch fieldname {
		// insertion point for per direct association field
		case "Classdiagrams":
			res := make(map[*Classdiagram]*DiagramPackage)
			for diagrampackage := range stage.DiagramPackages {
				for _, classdiagram_ := range diagrampackage.Classdiagrams {
					res[classdiagram_] = diagrampackage
				}
			}
			return any(res).(map[*End]*Start)
		case "Umlscs":
			res := make(map[*Umlsc]*DiagramPackage)
			for diagrampackage := range stage.DiagramPackages {
				for _, umlsc_ := range diagrampackage.Umlscs {
					res[umlsc_] = diagrampackage
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Field
	case FieldShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumShape
	case GongEnumShape:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValueEntrys":
			res := make(map[*GongEnumValueEntry]*GongEnumShape)
			for gongenumshape := range stage.GongEnumShapes {
				for _, gongenumvalueentry_ := range gongenumshape.GongEnumValueEntrys {
					res[gongenumvalueentry_] = gongenumshape
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GongEnumValueEntry
	case GongEnumValueEntry:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStructShape
	case GongStructShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Fields":
			res := make(map[*FieldShape]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, field_ := range gongstructshape.FieldShapes {
					res[field_] = gongstructshape
				}
			}
			return any(res).(map[*End]*Start)
		case "Links":
			res := make(map[*Link]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, link_ := range gongstructshape.Links {
					res[link_] = gongstructshape
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		case "NoteShapeLinks":
			res := make(map[*NoteShapeLink]*NoteShape)
			for noteshape := range stage.NoteShapes {
				for _, noteshapelink_ := range noteshape.NoteShapeLinks {
					res[noteshapelink_] = noteshape
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of NoteShapeLink
	case NoteShapeLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Position
	case Position:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UmlState
	case UmlState:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Umlsc
	case Umlsc:
		switch fieldname {
		// insertion point for per direct association field
		case "States":
			res := make(map[*UmlState]*Umlsc)
			for umlsc := range stage.Umlscs {
				for _, umlstate_ := range umlsc.States {
					res[umlstate_] = umlsc
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Vertice
	case Vertice:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Classdiagram:
		res = "Classdiagram"
	case DiagramPackage:
		res = "DiagramPackage"
	case FieldShape:
		res = "Field"
	case GongEnumShape:
		res = "GongEnumShape"
	case GongEnumValueEntry:
		res = "GongEnumValueEntry"
	case GongStructShape:
		res = "GongStructShape"
	case Link:
		res = "Link"
	case NoteShape:
		res = "NoteShape"
	case NoteShapeLink:
		res = "NoteShapeLink"
	case Position:
		res = "Position"
	case UmlState:
		res = "UmlState"
	case Umlsc:
		res = "Umlsc"
	case Vertice:
		res = "Vertice"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Classdiagram:
		res = "Classdiagram"
	case *DiagramPackage:
		res = "DiagramPackage"
	case *FieldShape:
		res = "Field"
	case *GongEnumShape:
		res = "GongEnumShape"
	case *GongEnumValueEntry:
		res = "GongEnumValueEntry"
	case *GongStructShape:
		res = "GongStructShape"
	case *Link:
		res = "Link"
	case *NoteShape:
		res = "NoteShape"
	case *NoteShapeLink:
		res = "NoteShapeLink"
	case *Position:
		res = "Position"
	case *UmlState:
		res = "UmlState"
	case *Umlsc:
		res = "Umlsc"
	case *Vertice:
		res = "Vertice"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Classdiagram:
		res = []string{"Name", "GongStructShapes", "GongEnumShapes", "NoteShapes", "IsInDrawMode", "IsExpanded"}
	case DiagramPackage:
		res = []string{"Name", "Path", "GongModelPath", "Classdiagrams", "SelectedClassdiagram", "Umlscs", "IsEditable", "IsReloaded", "AbsolutePathToDiagramPackage"}
	case FieldShape:
		res = []string{"Name", "Identifier", "FieldTypeAsString", "Structname", "Fieldtypename"}
	case GongEnumShape:
		res = []string{"Name", "Position", "Identifier", "GongEnumValueEntrys", "Width", "Height"}
	case GongEnumValueEntry:
		res = []string{"Name", "Identifier"}
	case GongStructShape:
		res = []string{"Name", "Position", "Identifier", "ShowNbInstances", "NbInstances", "Fields", "Links", "Width", "Height", "IsSelected", "IsExpanded"}
	case Link:
		res = []string{"Name", "Identifier", "Fieldtypename", "FieldOffsetX", "FieldOffsetY", "TargetMultiplicity", "TargetMultiplicityOffsetX", "TargetMultiplicityOffsetY", "SourceMultiplicity", "SourceMultiplicityOffsetX", "SourceMultiplicityOffsetY", "Middlevertice", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case NoteShape:
		res = []string{"Name", "Identifier", "Body", "BodyHTML", "X", "Y", "Width", "Height", "Matched", "NoteShapeLinks"}
	case NoteShapeLink:
		res = []string{"Name", "Identifier", "Type"}
	case Position:
		res = []string{"X", "Y", "Name"}
	case UmlState:
		res = []string{"Name", "X", "Y"}
	case Umlsc:
		res = []string{"Name", "States", "Activestate", "IsInDrawMode"}
	case Vertice:
		res = []string{"X", "Y", "Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Classdiagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramPackage"
		rf.Fieldname = "Classdiagrams"
		res = append(res, rf)
	case DiagramPackage:
		var rf ReverseField
		_ = rf
	case FieldShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "Fields"
		res = append(res, rf)
	case GongEnumShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongEnumShapes"
		res = append(res, rf)
	case GongEnumValueEntry:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongEnumShape"
		rf.Fieldname = "GongEnumValueEntrys"
		res = append(res, rf)
	case GongStructShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongStructShapes"
		res = append(res, rf)
	case Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "Links"
		res = append(res, rf)
	case NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "NoteShapes"
		res = append(res, rf)
	case NoteShapeLink:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "NoteShape"
		rf.Fieldname = "NoteShapeLinks"
		res = append(res, rf)
	case Position:
		var rf ReverseField
		_ = rf
	case UmlState:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Umlsc"
		rf.Fieldname = "States"
		res = append(res, rf)
	case Umlsc:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramPackage"
		rf.Fieldname = "Umlscs"
		res = append(res, rf)
	case Vertice:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Classdiagram:
		res = []string{"Name", "GongStructShapes", "GongEnumShapes", "NoteShapes", "IsInDrawMode", "IsExpanded"}
	case *DiagramPackage:
		res = []string{"Name", "Path", "GongModelPath", "Classdiagrams", "SelectedClassdiagram", "Umlscs", "IsEditable", "IsReloaded", "AbsolutePathToDiagramPackage"}
	case *FieldShape:
		res = []string{"Name", "Identifier", "FieldTypeAsString", "Structname", "Fieldtypename"}
	case *GongEnumShape:
		res = []string{"Name", "Position", "Identifier", "GongEnumValueEntrys", "Width", "Height"}
	case *GongEnumValueEntry:
		res = []string{"Name", "Identifier"}
	case *GongStructShape:
		res = []string{"Name", "Position", "Identifier", "ShowNbInstances", "NbInstances", "Fields", "Links", "Width", "Height", "IsSelected", "IsExpanded"}
	case *Link:
		res = []string{"Name", "Identifier", "Fieldtypename", "FieldOffsetX", "FieldOffsetY", "TargetMultiplicity", "TargetMultiplicityOffsetX", "TargetMultiplicityOffsetY", "SourceMultiplicity", "SourceMultiplicityOffsetX", "SourceMultiplicityOffsetY", "Middlevertice", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case *NoteShape:
		res = []string{"Name", "Identifier", "Body", "BodyHTML", "X", "Y", "Width", "Height", "Matched", "NoteShapeLinks"}
	case *NoteShapeLink:
		res = []string{"Name", "Identifier", "Type"}
	case *Position:
		res = []string{"X", "Y", "Name"}
	case *UmlState:
		res = []string{"Name", "X", "Y"}
	case *Umlsc:
		res = []string{"Name", "States", "Activestate", "IsInDrawMode"}
	case *Vertice:
		res = []string{"X", "Y", "Name"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
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

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Classdiagram:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "GongStructShapes":
			for idx, __instance__ := range inferedInstance.GongStructShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "GongEnumShapes":
			for idx, __instance__ := range inferedInstance.GongEnumShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "NoteShapes":
			for idx, __instance__ := range inferedInstance.NoteShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsInDrawMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInDrawMode)
			res.valueBool = inferedInstance.IsInDrawMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *DiagramPackage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Path":
			res.valueString = inferedInstance.Path
		case "GongModelPath":
			res.valueString = inferedInstance.GongModelPath
		case "Classdiagrams":
			for idx, __instance__ := range inferedInstance.Classdiagrams {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SelectedClassdiagram":
			if inferedInstance.SelectedClassdiagram != nil {
				res.valueString = inferedInstance.SelectedClassdiagram.Name
			}
		case "Umlscs":
			for idx, __instance__ := range inferedInstance.Umlscs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsEditable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsEditable)
			res.valueBool = inferedInstance.IsEditable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsReloaded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsReloaded)
			res.valueBool = inferedInstance.IsReloaded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AbsolutePathToDiagramPackage":
			res.valueString = inferedInstance.AbsolutePathToDiagramPackage
		}
	case *FieldShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "FieldTypeAsString":
			res.valueString = inferedInstance.FieldTypeAsString
		case "Structname":
			res.valueString = inferedInstance.Structname
		case "Fieldtypename":
			res.valueString = inferedInstance.Fieldtypename
		}
	case *GongEnumShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Position":
			if inferedInstance.Position != nil {
				res.valueString = inferedInstance.Position.Name
			}
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "GongEnumValueEntrys":
			for idx, __instance__ := range inferedInstance.GongEnumValueEntrys {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *GongEnumValueEntry:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		}
	case *GongStructShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Position":
			if inferedInstance.Position != nil {
				res.valueString = inferedInstance.Position.Name
			}
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "ShowNbInstances":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowNbInstances)
			res.valueBool = inferedInstance.ShowNbInstances
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NbInstances":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbInstances)
			res.valueInt = inferedInstance.NbInstances
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Fields":
			for idx, __instance__ := range inferedInstance.FieldShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "IsSelected":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelected)
			res.valueBool = inferedInstance.IsSelected
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Fieldtypename":
			res.valueString = inferedInstance.Fieldtypename
		case "FieldOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FieldOffsetX)
			res.valueFloat = inferedInstance.FieldOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FieldOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FieldOffsetY)
			res.valueFloat = inferedInstance.FieldOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TargetMultiplicity":
			enum := inferedInstance.TargetMultiplicity
			res.valueString = enum.ToCodeString()
		case "TargetMultiplicityOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetMultiplicityOffsetX)
			res.valueFloat = inferedInstance.TargetMultiplicityOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TargetMultiplicityOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetMultiplicityOffsetY)
			res.valueFloat = inferedInstance.TargetMultiplicityOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SourceMultiplicity":
			enum := inferedInstance.SourceMultiplicity
			res.valueString = enum.ToCodeString()
		case "SourceMultiplicityOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SourceMultiplicityOffsetX)
			res.valueFloat = inferedInstance.SourceMultiplicityOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SourceMultiplicityOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SourceMultiplicityOffsetY)
			res.valueFloat = inferedInstance.SourceMultiplicityOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Middlevertice":
			if inferedInstance.Middlevertice != nil {
				res.valueString = inferedInstance.Middlevertice.Name
			}
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res.valueString = enum.ToCodeString()
		case "StartRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartRatio)
			res.valueFloat = inferedInstance.StartRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res.valueString = enum.ToCodeString()
		case "EndRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndRatio)
			res.valueFloat = inferedInstance.EndRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerOffsetRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
			res.valueFloat = inferedInstance.CornerOffsetRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *NoteShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Body":
			res.valueString = inferedInstance.Body
		case "BodyHTML":
			res.valueString = inferedInstance.BodyHTML
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Matched":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Matched)
			res.valueBool = inferedInstance.Matched
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NoteShapeLinks":
			for idx, __instance__ := range inferedInstance.NoteShapeLinks {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *NoteShapeLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Type":
			enum := inferedInstance.Type
			res.valueString = enum.ToCodeString()
		}
	case *Position:
		switch fieldName {
		// string value of fields
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *UmlState:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *Umlsc:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "States":
			for idx, __instance__ := range inferedInstance.States {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Activestate":
			res.valueString = inferedInstance.Activestate
		case "IsInDrawMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInDrawMode)
			res.valueBool = inferedInstance.IsInDrawMode
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *Vertice:
		switch fieldName {
		// string value of fields
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Classdiagram:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "GongStructShapes":
			for idx, __instance__ := range inferedInstance.GongStructShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "GongEnumShapes":
			for idx, __instance__ := range inferedInstance.GongEnumShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "NoteShapes":
			for idx, __instance__ := range inferedInstance.NoteShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsInDrawMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInDrawMode)
			res.valueBool = inferedInstance.IsInDrawMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case DiagramPackage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Path":
			res.valueString = inferedInstance.Path
		case "GongModelPath":
			res.valueString = inferedInstance.GongModelPath
		case "Classdiagrams":
			for idx, __instance__ := range inferedInstance.Classdiagrams {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SelectedClassdiagram":
			if inferedInstance.SelectedClassdiagram != nil {
				res.valueString = inferedInstance.SelectedClassdiagram.Name
			}
		case "Umlscs":
			for idx, __instance__ := range inferedInstance.Umlscs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "IsEditable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsEditable)
			res.valueBool = inferedInstance.IsEditable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsReloaded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsReloaded)
			res.valueBool = inferedInstance.IsReloaded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AbsolutePathToDiagramPackage":
			res.valueString = inferedInstance.AbsolutePathToDiagramPackage
		}
	case FieldShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "FieldTypeAsString":
			res.valueString = inferedInstance.FieldTypeAsString
		case "Structname":
			res.valueString = inferedInstance.Structname
		case "Fieldtypename":
			res.valueString = inferedInstance.Fieldtypename
		}
	case GongEnumShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Position":
			if inferedInstance.Position != nil {
				res.valueString = inferedInstance.Position.Name
			}
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "GongEnumValueEntrys":
			for idx, __instance__ := range inferedInstance.GongEnumValueEntrys {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case GongEnumValueEntry:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		}
	case GongStructShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Position":
			if inferedInstance.Position != nil {
				res.valueString = inferedInstance.Position.Name
			}
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "ShowNbInstances":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowNbInstances)
			res.valueBool = inferedInstance.ShowNbInstances
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NbInstances":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbInstances)
			res.valueInt = inferedInstance.NbInstances
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Fields":
			for idx, __instance__ := range inferedInstance.FieldShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "IsSelected":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSelected)
			res.valueBool = inferedInstance.IsSelected
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Fieldtypename":
			res.valueString = inferedInstance.Fieldtypename
		case "FieldOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FieldOffsetX)
			res.valueFloat = inferedInstance.FieldOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "FieldOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FieldOffsetY)
			res.valueFloat = inferedInstance.FieldOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TargetMultiplicity":
			enum := inferedInstance.TargetMultiplicity
			res.valueString = enum.ToCodeString()
		case "TargetMultiplicityOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetMultiplicityOffsetX)
			res.valueFloat = inferedInstance.TargetMultiplicityOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TargetMultiplicityOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TargetMultiplicityOffsetY)
			res.valueFloat = inferedInstance.TargetMultiplicityOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SourceMultiplicity":
			enum := inferedInstance.SourceMultiplicity
			res.valueString = enum.ToCodeString()
		case "SourceMultiplicityOffsetX":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SourceMultiplicityOffsetX)
			res.valueFloat = inferedInstance.SourceMultiplicityOffsetX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "SourceMultiplicityOffsetY":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SourceMultiplicityOffsetY)
			res.valueFloat = inferedInstance.SourceMultiplicityOffsetY
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Middlevertice":
			if inferedInstance.Middlevertice != nil {
				res.valueString = inferedInstance.Middlevertice.Name
			}
		case "StartOrientation":
			enum := inferedInstance.StartOrientation
			res.valueString = enum.ToCodeString()
		case "StartRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StartRatio)
			res.valueFloat = inferedInstance.StartRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndOrientation":
			enum := inferedInstance.EndOrientation
			res.valueString = enum.ToCodeString()
		case "EndRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.EndRatio)
			res.valueFloat = inferedInstance.EndRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "CornerOffsetRatio":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CornerOffsetRatio)
			res.valueFloat = inferedInstance.CornerOffsetRatio
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case NoteShape:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Body":
			res.valueString = inferedInstance.Body
		case "BodyHTML":
			res.valueString = inferedInstance.BodyHTML
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Width":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Width)
			res.valueFloat = inferedInstance.Width
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Height":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Height)
			res.valueFloat = inferedInstance.Height
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Matched":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Matched)
			res.valueBool = inferedInstance.Matched
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NoteShapeLinks":
			for idx, __instance__ := range inferedInstance.NoteShapeLinks {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case NoteShapeLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "Type":
			enum := inferedInstance.Type
			res.valueString = enum.ToCodeString()
		}
	case Position:
		switch fieldName {
		// string value of fields
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case UmlState:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case Umlsc:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "States":
			for idx, __instance__ := range inferedInstance.States {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Activestate":
			res.valueString = inferedInstance.Activestate
		case "IsInDrawMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInDrawMode)
			res.valueBool = inferedInstance.IsInDrawMode
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case Vertice:
		switch fieldName {
		// string value of fields
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
