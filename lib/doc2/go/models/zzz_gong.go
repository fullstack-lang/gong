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

// can be used for
//     days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs

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
	AttributeShapes           map[*AttributeShape]any
	AttributeShapes_mapString map[string]*AttributeShape

	// insertion point for slice of pointers maps
	OnAfterAttributeShapeCreateCallback OnAfterCreateInterface[AttributeShape]
	OnAfterAttributeShapeUpdateCallback OnAfterUpdateInterface[AttributeShape]
	OnAfterAttributeShapeDeleteCallback OnAfterDeleteInterface[AttributeShape]
	OnAfterAttributeShapeReadCallback   OnAfterReadInterface[AttributeShape]

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

	OnAfterDiagramPackageCreateCallback OnAfterCreateInterface[DiagramPackage]
	OnAfterDiagramPackageUpdateCallback OnAfterUpdateInterface[DiagramPackage]
	OnAfterDiagramPackageDeleteCallback OnAfterDeleteInterface[DiagramPackage]
	OnAfterDiagramPackageReadCallback   OnAfterReadInterface[DiagramPackage]

	GongEnumShapes           map[*GongEnumShape]any
	GongEnumShapes_mapString map[string]*GongEnumShape

	// insertion point for slice of pointers maps
	GongEnumShape_GongEnumValueShapes_reverseMap map[*GongEnumValueShape]*GongEnumShape

	OnAfterGongEnumShapeCreateCallback OnAfterCreateInterface[GongEnumShape]
	OnAfterGongEnumShapeUpdateCallback OnAfterUpdateInterface[GongEnumShape]
	OnAfterGongEnumShapeDeleteCallback OnAfterDeleteInterface[GongEnumShape]
	OnAfterGongEnumShapeReadCallback   OnAfterReadInterface[GongEnumShape]

	GongEnumValueShapes           map[*GongEnumValueShape]any
	GongEnumValueShapes_mapString map[string]*GongEnumValueShape

	// insertion point for slice of pointers maps
	OnAfterGongEnumValueShapeCreateCallback OnAfterCreateInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeUpdateCallback OnAfterUpdateInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeDeleteCallback OnAfterDeleteInterface[GongEnumValueShape]
	OnAfterGongEnumValueShapeReadCallback   OnAfterReadInterface[GongEnumValueShape]

	GongStructShapes           map[*GongStructShape]any
	GongStructShapes_mapString map[string]*GongStructShape

	// insertion point for slice of pointers maps
	GongStructShape_AttributeShapes_reverseMap map[*AttributeShape]*GongStructShape

	GongStructShape_LinkShapes_reverseMap map[*LinkShape]*GongStructShape

	OnAfterGongStructShapeCreateCallback OnAfterCreateInterface[GongStructShape]
	OnAfterGongStructShapeUpdateCallback OnAfterUpdateInterface[GongStructShape]
	OnAfterGongStructShapeDeleteCallback OnAfterDeleteInterface[GongStructShape]
	OnAfterGongStructShapeReadCallback   OnAfterReadInterface[GongStructShape]

	LinkShapes           map[*LinkShape]any
	LinkShapes_mapString map[string]*LinkShape

	// insertion point for slice of pointers maps
	OnAfterLinkShapeCreateCallback OnAfterCreateInterface[LinkShape]
	OnAfterLinkShapeUpdateCallback OnAfterUpdateInterface[LinkShape]
	OnAfterLinkShapeDeleteCallback OnAfterDeleteInterface[LinkShape]
	OnAfterLinkShapeReadCallback   OnAfterReadInterface[LinkShape]

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

	GongStructShapeOrder            uint
	GongStructShapeMap_Staged_Order map[*GongStructShape]uint

	LinkShapeOrder            uint
	LinkShapeMap_Staged_Order map[*LinkShape]uint

	NoteShapeOrder            uint
	NoteShapeMap_Staged_Order map[*NoteShape]uint

	NoteShapeLinkOrder            uint
	NoteShapeLinkMap_Staged_Order map[*NoteShapeLink]uint

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
		case "GongStructShape":
			res = GetNamedStructInstances(stage.GongStructShapes, stage.GongStructShapeMap_Staged_Order)
		case "LinkShape":
			res = GetNamedStructInstances(stage.LinkShapes, stage.LinkShapeMap_Staged_Order)
		case "NoteShape":
			res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShapeMap_Staged_Order)
		case "NoteShapeLink":
			res = GetNamedStructInstances(stage.NoteShapeLinks, stage.NoteShapeLinkMap_Staged_Order)
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
	CommitGongStructShape(gongstructshape *GongStructShape)
	CheckoutGongStructShape(gongstructshape *GongStructShape)
	CommitLinkShape(linkshape *LinkShape)
	CheckoutLinkShape(linkshape *LinkShape)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteShapeLink(noteshapelink *NoteShapeLink)
	CheckoutNoteShapeLink(noteshapelink *NoteShapeLink)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		AttributeShapes:           make(map[*AttributeShape]any),
		AttributeShapes_mapString: make(map[string]*AttributeShape),

		Classdiagrams:           make(map[*Classdiagram]any),
		Classdiagrams_mapString: make(map[string]*Classdiagram),

		DiagramPackages:           make(map[*DiagramPackage]any),
		DiagramPackages_mapString: make(map[string]*DiagramPackage),

		GongEnumShapes:           make(map[*GongEnumShape]any),
		GongEnumShapes_mapString: make(map[string]*GongEnumShape),

		GongEnumValueShapes:           make(map[*GongEnumValueShape]any),
		GongEnumValueShapes_mapString: make(map[string]*GongEnumValueShape),

		GongStructShapes:           make(map[*GongStructShape]any),
		GongStructShapes_mapString: make(map[string]*GongStructShape),

		LinkShapes:           make(map[*LinkShape]any),
		LinkShapes_mapString: make(map[string]*LinkShape),

		NoteShapes:           make(map[*NoteShape]any),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteShapeLinks:           make(map[*NoteShapeLink]any),
		NoteShapeLinks_mapString: make(map[string]*NoteShapeLink),

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

		GongStructShapeMap_Staged_Order: make(map[*GongStructShape]uint),

		LinkShapeMap_Staged_Order: make(map[*LinkShape]uint),

		NoteShapeMap_Staged_Order: make(map[*NoteShape]uint),

		NoteShapeLinkMap_Staged_Order: make(map[*NoteShapeLink]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AttributeShape"},
			{name: "Classdiagram"},
			{name: "DiagramPackage"},
			{name: "GongEnumShape"},
			{name: "GongEnumValueShape"},
			{name: "GongStructShape"},
			{name: "LinkShape"},
			{name: "NoteShape"},
			{name: "NoteShapeLink"},
		}, // end of insertion point
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
	case *GongStructShape:
		return stage.GongStructShapeMap_Staged_Order[instance]
	case *LinkShape:
		return stage.LinkShapeMap_Staged_Order[instance]
	case *NoteShape:
		return stage.NoteShapeMap_Staged_Order[instance]
	case *NoteShapeLink:
		return stage.NoteShapeLinkMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["AttributeShape"] = len(stage.AttributeShapes)
	stage.Map_GongStructName_InstancesNb["Classdiagram"] = len(stage.Classdiagrams)
	stage.Map_GongStructName_InstancesNb["DiagramPackage"] = len(stage.DiagramPackages)
	stage.Map_GongStructName_InstancesNb["GongEnumShape"] = len(stage.GongEnumShapes)
	stage.Map_GongStructName_InstancesNb["GongEnumValueShape"] = len(stage.GongEnumValueShapes)
	stage.Map_GongStructName_InstancesNb["GongStructShape"] = len(stage.GongStructShapes)
	stage.Map_GongStructName_InstancesNb["LinkShape"] = len(stage.LinkShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteShapeLink"] = len(stage.NoteShapeLinks)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AttributeShape"] = len(stage.AttributeShapes)
	stage.Map_GongStructName_InstancesNb["Classdiagram"] = len(stage.Classdiagrams)
	stage.Map_GongStructName_InstancesNb["DiagramPackage"] = len(stage.DiagramPackages)
	stage.Map_GongStructName_InstancesNb["GongEnumShape"] = len(stage.GongEnumShapes)
	stage.Map_GongStructName_InstancesNb["GongEnumValueShape"] = len(stage.GongEnumValueShapes)
	stage.Map_GongStructName_InstancesNb["GongStructShape"] = len(stage.GongStructShapes)
	stage.Map_GongStructName_InstancesNb["LinkShape"] = len(stage.LinkShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteShapeLink"] = len(stage.NoteShapeLinks)

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
		stage.AttributeShapes[attributeshape] = __member
		stage.AttributeShapeMap_Staged_Order[attributeshape] = stage.AttributeShapeOrder
		stage.AttributeShapeOrder++
	}
	stage.AttributeShapes_mapString[attributeshape.Name] = attributeshape

	return attributeshape
}

// Unstage removes attributeshape off the model stage
func (attributeshape *AttributeShape) Unstage(stage *Stage) *AttributeShape {
	delete(stage.AttributeShapes, attributeshape)
	delete(stage.AttributeShapes_mapString, attributeshape.Name)
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

// Stage puts gongenumvalueshape to the model stage
func (gongenumvalueshape *GongEnumValueShape) Stage(stage *Stage) *GongEnumValueShape {

	if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; !ok {
		stage.GongEnumValueShapes[gongenumvalueshape] = __member
		stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshape] = stage.GongEnumValueShapeOrder
		stage.GongEnumValueShapeOrder++
	}
	stage.GongEnumValueShapes_mapString[gongenumvalueshape.Name] = gongenumvalueshape

	return gongenumvalueshape
}

// Unstage removes gongenumvalueshape off the model stage
func (gongenumvalueshape *GongEnumValueShape) Unstage(stage *Stage) *GongEnumValueShape {
	delete(stage.GongEnumValueShapes, gongenumvalueshape)
	delete(stage.GongEnumValueShapes_mapString, gongenumvalueshape.Name)
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

// Stage puts linkshape to the model stage
func (linkshape *LinkShape) Stage(stage *Stage) *LinkShape {

	if _, ok := stage.LinkShapes[linkshape]; !ok {
		stage.LinkShapes[linkshape] = __member
		stage.LinkShapeMap_Staged_Order[linkshape] = stage.LinkShapeOrder
		stage.LinkShapeOrder++
	}
	stage.LinkShapes_mapString[linkshape.Name] = linkshape

	return linkshape
}

// Unstage removes linkshape off the model stage
func (linkshape *LinkShape) Unstage(stage *Stage) *LinkShape {
	delete(stage.LinkShapes, linkshape)
	delete(stage.LinkShapes_mapString, linkshape.Name)
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAttributeShape(AttributeShape *AttributeShape)
	CreateORMClassdiagram(Classdiagram *Classdiagram)
	CreateORMDiagramPackage(DiagramPackage *DiagramPackage)
	CreateORMGongEnumShape(GongEnumShape *GongEnumShape)
	CreateORMGongEnumValueShape(GongEnumValueShape *GongEnumValueShape)
	CreateORMGongStructShape(GongStructShape *GongStructShape)
	CreateORMLinkShape(LinkShape *LinkShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteShapeLink(NoteShapeLink *NoteShapeLink)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAttributeShape(AttributeShape *AttributeShape)
	DeleteORMClassdiagram(Classdiagram *Classdiagram)
	DeleteORMDiagramPackage(DiagramPackage *DiagramPackage)
	DeleteORMGongEnumShape(GongEnumShape *GongEnumShape)
	DeleteORMGongEnumValueShape(GongEnumValueShape *GongEnumValueShape)
	DeleteORMGongStructShape(GongStructShape *GongStructShape)
	DeleteORMLinkShape(LinkShape *LinkShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteShapeLink(NoteShapeLink *NoteShapeLink)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AttributeShapes = make(map[*AttributeShape]any)
	stage.AttributeShapes_mapString = make(map[string]*AttributeShape)
	stage.AttributeShapeMap_Staged_Order = make(map[*AttributeShape]uint)
	stage.AttributeShapeOrder = 0

	stage.Classdiagrams = make(map[*Classdiagram]any)
	stage.Classdiagrams_mapString = make(map[string]*Classdiagram)
	stage.ClassdiagramMap_Staged_Order = make(map[*Classdiagram]uint)
	stage.ClassdiagramOrder = 0

	stage.DiagramPackages = make(map[*DiagramPackage]any)
	stage.DiagramPackages_mapString = make(map[string]*DiagramPackage)
	stage.DiagramPackageMap_Staged_Order = make(map[*DiagramPackage]uint)
	stage.DiagramPackageOrder = 0

	stage.GongEnumShapes = make(map[*GongEnumShape]any)
	stage.GongEnumShapes_mapString = make(map[string]*GongEnumShape)
	stage.GongEnumShapeMap_Staged_Order = make(map[*GongEnumShape]uint)
	stage.GongEnumShapeOrder = 0

	stage.GongEnumValueShapes = make(map[*GongEnumValueShape]any)
	stage.GongEnumValueShapes_mapString = make(map[string]*GongEnumValueShape)
	stage.GongEnumValueShapeMap_Staged_Order = make(map[*GongEnumValueShape]uint)
	stage.GongEnumValueShapeOrder = 0

	stage.GongStructShapes = make(map[*GongStructShape]any)
	stage.GongStructShapes_mapString = make(map[string]*GongStructShape)
	stage.GongStructShapeMap_Staged_Order = make(map[*GongStructShape]uint)
	stage.GongStructShapeOrder = 0

	stage.LinkShapes = make(map[*LinkShape]any)
	stage.LinkShapes_mapString = make(map[string]*LinkShape)
	stage.LinkShapeMap_Staged_Order = make(map[*LinkShape]uint)
	stage.LinkShapeOrder = 0

	stage.NoteShapes = make(map[*NoteShape]any)
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShapeMap_Staged_Order = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteShapeLinks = make(map[*NoteShapeLink]any)
	stage.NoteShapeLinks_mapString = make(map[string]*NoteShapeLink)
	stage.NoteShapeLinkMap_Staged_Order = make(map[*NoteShapeLink]uint)
	stage.NoteShapeLinkOrder = 0

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

	stage.GongStructShapes = nil
	stage.GongStructShapes_mapString = nil

	stage.LinkShapes = nil
	stage.LinkShapes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteShapeLinks = nil
	stage.NoteShapeLinks_mapString = nil

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

	for gongstructshape := range stage.GongStructShapes {
		gongstructshape.Unstage(stage)
	}

	for linkshape := range stage.LinkShapes {
		linkshape.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for noteshapelink := range stage.NoteShapeLinks {
		noteshapelink.Unstage(stage)
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
	case map[*GongStructShape]any:
		return any(&stage.GongStructShapes).(*Type)
	case map[*LinkShape]any:
		return any(&stage.LinkShapes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteShapeLink]any:
		return any(&stage.NoteShapeLinks).(*Type)
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
	case map[string]*AttributeShape:
		return any(&stage.AttributeShapes_mapString).(*Type)
	case map[string]*Classdiagram:
		return any(&stage.Classdiagrams_mapString).(*Type)
	case map[string]*DiagramPackage:
		return any(&stage.DiagramPackages_mapString).(*Type)
	case map[string]*GongEnumShape:
		return any(&stage.GongEnumShapes_mapString).(*Type)
	case map[string]*GongEnumValueShape:
		return any(&stage.GongEnumValueShapes_mapString).(*Type)
	case map[string]*GongStructShape:
		return any(&stage.GongStructShapes_mapString).(*Type)
	case map[string]*LinkShape:
		return any(&stage.LinkShapes_mapString).(*Type)
	case map[string]*NoteShape:
		return any(&stage.NoteShapes_mapString).(*Type)
	case map[string]*NoteShapeLink:
		return any(&stage.NoteShapeLinks_mapString).(*Type)
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
	case AttributeShape:
		return any(&stage.AttributeShapes).(*map[*Type]any)
	case Classdiagram:
		return any(&stage.Classdiagrams).(*map[*Type]any)
	case DiagramPackage:
		return any(&stage.DiagramPackages).(*map[*Type]any)
	case GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[*Type]any)
	case GongEnumValueShape:
		return any(&stage.GongEnumValueShapes).(*map[*Type]any)
	case GongStructShape:
		return any(&stage.GongStructShapes).(*map[*Type]any)
	case LinkShape:
		return any(&stage.LinkShapes).(*map[*Type]any)
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]any)
	case NoteShapeLink:
		return any(&stage.NoteShapeLinks).(*map[*Type]any)
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
	case *AttributeShape:
		return any(&stage.AttributeShapes).(*map[Type]any)
	case *Classdiagram:
		return any(&stage.Classdiagrams).(*map[Type]any)
	case *DiagramPackage:
		return any(&stage.DiagramPackages).(*map[Type]any)
	case *GongEnumShape:
		return any(&stage.GongEnumShapes).(*map[Type]any)
	case *GongEnumValueShape:
		return any(&stage.GongEnumValueShapes).(*map[Type]any)
	case *GongStructShape:
		return any(&stage.GongStructShapes).(*map[Type]any)
	case *LinkShape:
		return any(&stage.LinkShapes).(*map[Type]any)
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]any)
	case *NoteShapeLink:
		return any(&stage.NoteShapeLinks).(*map[Type]any)
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
	case GongStructShape:
		return any(&stage.GongStructShapes_mapString).(*map[string]*Type)
	case LinkShape:
		return any(&stage.LinkShapes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteShapeLink:
		return any(&stage.NoteShapeLinks_mapString).(*map[string]*Type)
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
		}
	// reverse maps of direct associations of GongEnumShape
	case GongEnumShape:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValueShapes":
			res := make(map[*GongEnumValueShape]*GongEnumShape)
			for gongenumshape := range stage.GongEnumShapes {
				for _, gongenumvalueshape_ := range gongenumshape.GongEnumValueShapes {
					res[gongenumvalueshape_] = gongenumshape
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GongEnumValueShape
	case GongEnumValueShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStructShape
	case GongStructShape:
		switch fieldname {
		// insertion point for per direct association field
		case "AttributeShapes":
			res := make(map[*AttributeShape]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, attributeshape_ := range gongstructshape.AttributeShapes {
					res[attributeshape_] = gongstructshape
				}
			}
			return any(res).(map[*End]*Start)
		case "LinkShapes":
			res := make(map[*LinkShape]*GongStructShape)
			for gongstructshape := range stage.GongStructShapes {
				for _, linkshape_ := range gongstructshape.LinkShapes {
					res[linkshape_] = gongstructshape
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of LinkShape
	case LinkShape:
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
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AttributeShape:
		res = "AttributeShape"
	case Classdiagram:
		res = "Classdiagram"
	case DiagramPackage:
		res = "DiagramPackage"
	case GongEnumShape:
		res = "GongEnumShape"
	case GongEnumValueShape:
		res = "GongEnumValueShape"
	case GongStructShape:
		res = "GongStructShape"
	case LinkShape:
		res = "LinkShape"
	case NoteShape:
		res = "NoteShape"
	case NoteShapeLink:
		res = "NoteShapeLink"
	}
	return res
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
	case *GongStructShape:
		res = "GongStructShape"
	case *LinkShape:
		res = "LinkShape"
	case *NoteShape:
		res = "NoteShape"
	case *NoteShapeLink:
		res = "NoteShapeLink"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AttributeShape:
		res = []string{"Name", "Identifier", "FieldTypeAsString", "Structname", "Fieldtypename"}
	case Classdiagram:
		res = []string{"Name", "GongStructShapes", "GongEnumShapes", "NoteShapes", "IsInRenameMode", "IsExpanded", "NodeGongStructsIsExpanded", "NodeGongStructNodeExpansionBinaryEncoding", "NodeGongEnumsIsExpanded", "NodeGongEnumNodeExpansionBinaryEncoding", "NodeGongNotesIsExpanded", "NodeGongNoteNodeExpansionBinaryEncoding"}
	case DiagramPackage:
		res = []string{"Name", "Path", "GongModelPath", "Classdiagrams", "SelectedClassdiagram", "AbsolutePathToDiagramPackage"}
	case GongEnumShape:
		res = []string{"Name", "X", "Y", "Identifier", "GongEnumValueShapes", "Width", "Height", "IsExpanded"}
	case GongEnumValueShape:
		res = []string{"Name", "Identifier"}
	case GongStructShape:
		res = []string{"Name", "X", "Y", "Identifier", "ShowNbInstances", "NbInstances", "AttributeShapes", "LinkShapes", "Width", "Height", "IsSelected"}
	case LinkShape:
		res = []string{"Name", "Identifier", "Fieldtypename", "FieldOffsetX", "FieldOffsetY", "TargetMultiplicity", "TargetMultiplicityOffsetX", "TargetMultiplicityOffsetY", "SourceMultiplicity", "SourceMultiplicityOffsetX", "SourceMultiplicityOffsetY", "X", "Y", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case NoteShape:
		res = []string{"Name", "Identifier", "Body", "BodyHTML", "X", "Y", "Width", "Height", "Matched", "NoteShapeLinks", "IsExpanded"}
	case NoteShapeLink:
		res = []string{"Name", "Identifier", "Type"}
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
	case AttributeShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "AttributeShapes"
		res = append(res, rf)
	case Classdiagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramPackage"
		rf.Fieldname = "Classdiagrams"
		res = append(res, rf)
	case DiagramPackage:
		var rf ReverseField
		_ = rf
	case GongEnumShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongEnumShapes"
		res = append(res, rf)
	case GongEnumValueShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongEnumShape"
		rf.Fieldname = "GongEnumValueShapes"
		res = append(res, rf)
	case GongStructShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Classdiagram"
		rf.Fieldname = "GongStructShapes"
		res = append(res, rf)
	case LinkShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStructShape"
		rf.Fieldname = "LinkShapes"
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
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AttributeShape:
		res = []string{"Name", "Identifier", "FieldTypeAsString", "Structname", "Fieldtypename"}
	case *Classdiagram:
		res = []string{"Name", "GongStructShapes", "GongEnumShapes", "NoteShapes", "IsInRenameMode", "IsExpanded", "NodeGongStructsIsExpanded", "NodeGongStructNodeExpansionBinaryEncoding", "NodeGongEnumsIsExpanded", "NodeGongEnumNodeExpansionBinaryEncoding", "NodeGongNotesIsExpanded", "NodeGongNoteNodeExpansionBinaryEncoding"}
	case *DiagramPackage:
		res = []string{"Name", "Path", "GongModelPath", "Classdiagrams", "SelectedClassdiagram", "AbsolutePathToDiagramPackage"}
	case *GongEnumShape:
		res = []string{"Name", "X", "Y", "Identifier", "GongEnumValueShapes", "Width", "Height", "IsExpanded"}
	case *GongEnumValueShape:
		res = []string{"Name", "Identifier"}
	case *GongStructShape:
		res = []string{"Name", "X", "Y", "Identifier", "ShowNbInstances", "NbInstances", "AttributeShapes", "LinkShapes", "Width", "Height", "IsSelected"}
	case *LinkShape:
		res = []string{"Name", "Identifier", "Fieldtypename", "FieldOffsetX", "FieldOffsetY", "TargetMultiplicity", "TargetMultiplicityOffsetX", "TargetMultiplicityOffsetY", "SourceMultiplicity", "SourceMultiplicityOffsetX", "SourceMultiplicityOffsetY", "X", "Y", "StartOrientation", "StartRatio", "EndOrientation", "EndRatio", "CornerOffsetRatio"}
	case *NoteShape:
		res = []string{"Name", "Identifier", "Body", "BodyHTML", "X", "Y", "Width", "Height", "Matched", "NoteShapeLinks", "IsExpanded"}
	case *NoteShapeLink:
		res = []string{"Name", "Identifier", "Type"}
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
	case *AttributeShape:
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
		case "IsInRenameMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInRenameMode)
			res.valueBool = inferedInstance.IsInRenameMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongStructsIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongStructsIsExpanded)
			res.valueBool = inferedInstance.NodeGongStructsIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongStructNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongStructNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongStructNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NodeGongEnumsIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongEnumsIsExpanded)
			res.valueBool = inferedInstance.NodeGongEnumsIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongEnumNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongEnumNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongEnumNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NodeGongNotesIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongNotesIsExpanded)
			res.valueBool = inferedInstance.NodeGongNotesIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongNoteNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongNoteNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongNoteNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
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
		case "AbsolutePathToDiagramPackage":
			res.valueString = inferedInstance.AbsolutePathToDiagramPackage
		}
	case *GongEnumShape:
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
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "GongEnumValueShapes":
			for idx, __instance__ := range inferedInstance.GongEnumValueShapes {
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
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case *GongEnumValueShape:
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
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
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
		case "AttributeShapes":
			for idx, __instance__ := range inferedInstance.AttributeShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "LinkShapes":
			for idx, __instance__ := range inferedInstance.LinkShapes {
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
		}
	case *LinkShape:
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
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
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
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
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
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case AttributeShape:
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
		case "IsInRenameMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInRenameMode)
			res.valueBool = inferedInstance.IsInRenameMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongStructsIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongStructsIsExpanded)
			res.valueBool = inferedInstance.NodeGongStructsIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongStructNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongStructNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongStructNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NodeGongEnumsIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongEnumsIsExpanded)
			res.valueBool = inferedInstance.NodeGongEnumsIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongEnumNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongEnumNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongEnumNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NodeGongNotesIsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.NodeGongNotesIsExpanded)
			res.valueBool = inferedInstance.NodeGongNotesIsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "NodeGongNoteNodeExpansionBinaryEncoding":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NodeGongNoteNodeExpansionBinaryEncoding)
			res.valueInt = inferedInstance.NodeGongNoteNodeExpansionBinaryEncoding
			res.GongFieldValueType = GongFieldValueTypeInt
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
		case "AbsolutePathToDiagramPackage":
			res.valueString = inferedInstance.AbsolutePathToDiagramPackage
		}
	case GongEnumShape:
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
		case "Identifier":
			res.valueString = inferedInstance.Identifier
		case "GongEnumValueShapes":
			for idx, __instance__ := range inferedInstance.GongEnumValueShapes {
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
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		}
	case GongEnumValueShape:
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
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
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
		case "AttributeShapes":
			for idx, __instance__ := range inferedInstance.AttributeShapes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "LinkShapes":
			for idx, __instance__ := range inferedInstance.LinkShapes {
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
		}
	case LinkShape:
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
		case "X":
			res.valueString = fmt.Sprintf("%f", inferedInstance.X)
			res.valueFloat = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Y)
			res.valueFloat = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeFloat
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
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
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
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
