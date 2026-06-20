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
	"sync"
	"time"

	structure_go "github.com/fullstack-lang/gong/dsm/structure/go"
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

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
	ProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

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
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	DiagramStructures                map[*DiagramStructure]struct{}
	DiagramStructures_instance       map[*DiagramStructure]*DiagramStructure
	DiagramStructures_mapString      map[string]*DiagramStructure
	DiagramStructureOrder            uint
	DiagramStructure_stagedOrder     map[*DiagramStructure]uint
	DiagramStructure_orderStaged     map[uint]*DiagramStructure
	DiagramStructures_reference      map[*DiagramStructure]*DiagramStructure
	DiagramStructures_referenceOrder map[*DiagramStructure]uint

	// insertion point for slice of pointers maps
	DiagramStructure_Part_Shapes_reverseMap map[*PartShape]*DiagramStructure

	DiagramStructure_PartsWhoseNodeIsExpanded_reverseMap map[*Part]*DiagramStructure

	DiagramStructure_Link_Shapes_reverseMap map[*LinkAssociationShape]*DiagramStructure

	DiagramStructure_LinksWhoseNodeIsExpanded_reverseMap map[*Link]*DiagramStructure

	OnAfterDiagramStructureCreateCallback OnAfterCreateInterface[DiagramStructure]
	OnAfterDiagramStructureUpdateCallback OnAfterUpdateInterface[DiagramStructure]
	OnAfterDiagramStructureDeleteCallback OnAfterDeleteInterface[DiagramStructure]
	OnAfterDiagramStructureReadCallback   OnAfterReadInterface[DiagramStructure]

	Librarys                map[*Library]struct{}
	Librarys_instance       map[*Library]*Library
	Librarys_mapString      map[string]*Library
	LibraryOrder            uint
	Library_stagedOrder     map[*Library]uint
	Library_orderStaged     map[uint]*Library
	Librarys_reference      map[*Library]*Library
	Librarys_referenceOrder map[*Library]uint

	// insertion point for slice of pointers maps
	Library_SubLibraries_reverseMap map[*Library]*Library

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_RootSystems_reverseMap map[*System]*Library

	Library_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Links                map[*Link]struct{}
	Links_instance       map[*Link]*Link
	Links_mapString      map[string]*Link
	LinkOrder            uint
	Link_stagedOrder     map[*Link]uint
	Link_orderStaged     map[uint]*Link
	Links_reference      map[*Link]*Link
	Links_referenceOrder map[*Link]uint

	// insertion point for slice of pointers maps
	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	LinkAssociationShapes                map[*LinkAssociationShape]struct{}
	LinkAssociationShapes_instance       map[*LinkAssociationShape]*LinkAssociationShape
	LinkAssociationShapes_mapString      map[string]*LinkAssociationShape
	LinkAssociationShapeOrder            uint
	LinkAssociationShape_stagedOrder     map[*LinkAssociationShape]uint
	LinkAssociationShape_orderStaged     map[uint]*LinkAssociationShape
	LinkAssociationShapes_reference      map[*LinkAssociationShape]*LinkAssociationShape
	LinkAssociationShapes_referenceOrder map[*LinkAssociationShape]uint

	// insertion point for slice of pointers maps
	OnAfterLinkAssociationShapeCreateCallback OnAfterCreateInterface[LinkAssociationShape]
	OnAfterLinkAssociationShapeUpdateCallback OnAfterUpdateInterface[LinkAssociationShape]
	OnAfterLinkAssociationShapeDeleteCallback OnAfterDeleteInterface[LinkAssociationShape]
	OnAfterLinkAssociationShapeReadCallback   OnAfterReadInterface[LinkAssociationShape]

	Parts                map[*Part]struct{}
	Parts_instance       map[*Part]*Part
	Parts_mapString      map[string]*Part
	PartOrder            uint
	Part_stagedOrder     map[*Part]uint
	Part_orderStaged     map[uint]*Part
	Parts_reference      map[*Part]*Part
	Parts_referenceOrder map[*Part]uint

	// insertion point for slice of pointers maps
	OnAfterPartCreateCallback OnAfterCreateInterface[Part]
	OnAfterPartUpdateCallback OnAfterUpdateInterface[Part]
	OnAfterPartDeleteCallback OnAfterDeleteInterface[Part]
	OnAfterPartReadCallback   OnAfterReadInterface[Part]

	PartShapes                map[*PartShape]struct{}
	PartShapes_instance       map[*PartShape]*PartShape
	PartShapes_mapString      map[string]*PartShape
	PartShapeOrder            uint
	PartShape_stagedOrder     map[*PartShape]uint
	PartShape_orderStaged     map[uint]*PartShape
	PartShapes_reference      map[*PartShape]*PartShape
	PartShapes_referenceOrder map[*PartShape]uint

	// insertion point for slice of pointers maps
	OnAfterPartShapeCreateCallback OnAfterCreateInterface[PartShape]
	OnAfterPartShapeUpdateCallback OnAfterUpdateInterface[PartShape]
	OnAfterPartShapeDeleteCallback OnAfterDeleteInterface[PartShape]
	OnAfterPartShapeReadCallback   OnAfterReadInterface[PartShape]

	Systems                map[*System]struct{}
	Systems_instance       map[*System]*System
	Systems_mapString      map[string]*System
	SystemOrder            uint
	System_stagedOrder     map[*System]uint
	System_orderStaged     map[uint]*System
	Systems_reference      map[*System]*System
	Systems_referenceOrder map[*System]uint

	// insertion point for slice of pointers maps
	System_Parts_reverseMap map[*Part]*System

	System_PartsWhoseNodeIsExpanded_reverseMap map[*Part]*System

	System_SubSystems_reverseMap map[*System]*System

	System_SubSystemsWhoseNodeIsExpanded_reverseMap map[*System]*System

	System_Links_reverseMap map[*Link]*System

	System_LinksWhoseNodeIsExpanded_reverseMap map[*Link]*System

	System_DiagramStructures_reverseMap map[*DiagramStructure]*System

	System_DiagramStructuresWhoseNodeIsExpanded_reverseMap map[*DiagramStructure]*System

	OnAfterSystemCreateCallback OnAfterCreateInterface[System]
	OnAfterSystemUpdateCallback OnAfterUpdateInterface[System]
	OnAfterSystemDeleteCallback OnAfterDeleteInterface[System]
	OnAfterSystemReadCallback   OnAfterReadInterface[System]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

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
	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	stage.isApplyingBackwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingBackwardCommit = false
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	stage.isApplyingForwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingForwardCommit = false
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	stage.DiagramStructures_reference = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_instance = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_instance = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint)

	stage.LinkAssociationShapes_reference = make(map[*LinkAssociationShape]*LinkAssociationShape)
	stage.LinkAssociationShapes_instance = make(map[*LinkAssociationShape]*LinkAssociationShape)
	stage.LinkAssociationShapes_referenceOrder = make(map[*LinkAssociationShape]uint)

	stage.Parts_reference = make(map[*Part]*Part)
	stage.Parts_instance = make(map[*Part]*Part)
	stage.Parts_referenceOrder = make(map[*Part]uint)

	stage.PartShapes_reference = make(map[*PartShape]*PartShape)
	stage.PartShapes_instance = make(map[*PartShape]*PartShape)
	stage.PartShapes_referenceOrder = make(map[*PartShape]uint)

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_instance = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxDiagramStructureOrder uint
	var foundDiagramStructure bool
	for _, order := range stage.DiagramStructure_stagedOrder {
		if !foundDiagramStructure || order > maxDiagramStructureOrder {
			maxDiagramStructureOrder = order
			foundDiagramStructure = true
		}
	}
	if foundDiagramStructure {
		stage.DiagramStructureOrder = maxDiagramStructureOrder + 1
	} else {
		stage.DiagramStructureOrder = 0
	}

	var maxLibraryOrder uint
	var foundLibrary bool
	for _, order := range stage.Library_stagedOrder {
		if !foundLibrary || order > maxLibraryOrder {
			maxLibraryOrder = order
			foundLibrary = true
		}
	}
	if foundLibrary {
		stage.LibraryOrder = maxLibraryOrder + 1
	} else {
		stage.LibraryOrder = 0
	}

	var maxLinkOrder uint
	var foundLink bool
	for _, order := range stage.Link_stagedOrder {
		if !foundLink || order > maxLinkOrder {
			maxLinkOrder = order
			foundLink = true
		}
	}
	if foundLink {
		stage.LinkOrder = maxLinkOrder + 1
	} else {
		stage.LinkOrder = 0
	}

	var maxLinkAssociationShapeOrder uint
	var foundLinkAssociationShape bool
	for _, order := range stage.LinkAssociationShape_stagedOrder {
		if !foundLinkAssociationShape || order > maxLinkAssociationShapeOrder {
			maxLinkAssociationShapeOrder = order
			foundLinkAssociationShape = true
		}
	}
	if foundLinkAssociationShape {
		stage.LinkAssociationShapeOrder = maxLinkAssociationShapeOrder + 1
	} else {
		stage.LinkAssociationShapeOrder = 0
	}

	var maxPartOrder uint
	var foundPart bool
	for _, order := range stage.Part_stagedOrder {
		if !foundPart || order > maxPartOrder {
			maxPartOrder = order
			foundPart = true
		}
	}
	if foundPart {
		stage.PartOrder = maxPartOrder + 1
	} else {
		stage.PartOrder = 0
	}

	var maxPartShapeOrder uint
	var foundPartShape bool
	for _, order := range stage.PartShape_stagedOrder {
		if !foundPartShape || order > maxPartShapeOrder {
			maxPartShapeOrder = order
			foundPartShape = true
		}
	}
	if foundPartShape {
		stage.PartShapeOrder = maxPartShapeOrder + 1
	} else {
		stage.PartShapeOrder = 0
	}

	var maxSystemOrder uint
	var foundSystem bool
	for _, order := range stage.System_stagedOrder {
		if !foundSystem || order > maxSystemOrder {
			maxSystemOrder = order
			foundSystem = true
		}
	}
	if foundSystem {
		stage.SystemOrder = maxSystemOrder + 1
	} else {
		stage.SystemOrder = 0
	}

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {
	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
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

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *DiagramStructure:
		tmp := GetStructInstancesByOrder(stage.DiagramStructures, stage.DiagramStructure_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramStructure implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Library:
		tmp := GetStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Library implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Link:
		tmp := GetStructInstancesByOrder(stage.Links, stage.Link_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Link implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LinkAssociationShape:
		tmp := GetStructInstancesByOrder(stage.LinkAssociationShapes, stage.LinkAssociationShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LinkAssociationShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Part:
		tmp := GetStructInstancesByOrder(stage.Parts, stage.Part_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Part implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PartShape:
		tmp := GetStructInstancesByOrder(stage.PartShapes, stage.PartShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PartShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *System:
		tmp := GetStructInstancesByOrder(stage.Systems, stage.System_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *System implements.
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
	case "DiagramStructure":
		res = GetNamedStructInstances(stage.DiagramStructures, stage.DiagramStructure_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Link":
		res = GetNamedStructInstances(stage.Links, stage.Link_stagedOrder)
	case "LinkAssociationShape":
		res = GetNamedStructInstances(stage.LinkAssociationShapes, stage.LinkAssociationShape_stagedOrder)
	case "Part":
		res = GetNamedStructInstances(stage.Parts, stage.Part_stagedOrder)
	case "PartShape":
		res = GetNamedStructInstances(stage.PartShapes, stage.PartShape_stagedOrder)
	case "System":
		res = GetNamedStructInstances(stage.Systems, stage.System_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/structure/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return structure_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return structure_go.GoDiagramsDir
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
	CommitDiagramStructure(diagramstructure *DiagramStructure)
	CheckoutDiagramStructure(diagramstructure *DiagramStructure)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	CommitLinkAssociationShape(linkassociationshape *LinkAssociationShape)
	CheckoutLinkAssociationShape(linkassociationshape *LinkAssociationShape)
	CommitPart(part *Part)
	CheckoutPart(part *Part)
	CommitPartShape(partshape *PartShape)
	CheckoutPartShape(partshape *PartShape)
	CommitSystem(system *System)
	CheckoutSystem(system *System)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		DiagramStructures:           make(map[*DiagramStructure]struct{}),
		DiagramStructures_mapString: make(map[string]*DiagramStructure),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Links:           make(map[*Link]struct{}),
		Links_mapString: make(map[string]*Link),

		LinkAssociationShapes:           make(map[*LinkAssociationShape]struct{}),
		LinkAssociationShapes_mapString: make(map[string]*LinkAssociationShape),

		Parts:           make(map[*Part]struct{}),
		Parts_mapString: make(map[string]*Part),

		PartShapes:           make(map[*PartShape]struct{}),
		PartShapes_mapString: make(map[string]*PartShape),

		Systems:           make(map[*System]struct{}),
		Systems_mapString: make(map[string]*System),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DiagramStructure_stagedOrder: make(map[*DiagramStructure]uint),
		DiagramStructure_orderStaged: make(map[uint]*DiagramStructure),
		DiagramStructures_reference:  make(map[*DiagramStructure]*DiagramStructure),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Link_stagedOrder: make(map[*Link]uint),
		Link_orderStaged: make(map[uint]*Link),
		Links_reference:  make(map[*Link]*Link),

		LinkAssociationShape_stagedOrder: make(map[*LinkAssociationShape]uint),
		LinkAssociationShape_orderStaged: make(map[uint]*LinkAssociationShape),
		LinkAssociationShapes_reference:  make(map[*LinkAssociationShape]*LinkAssociationShape),

		Part_stagedOrder: make(map[*Part]uint),
		Part_orderStaged: make(map[uint]*Part),
		Parts_reference:  make(map[*Part]*Part),

		PartShape_stagedOrder: make(map[*PartShape]uint),
		PartShape_orderStaged: make(map[uint]*PartShape),
		PartShapes_reference:  make(map[*PartShape]*PartShape),

		System_stagedOrder: make(map[*System]uint),
		System_orderStaged: make(map[uint]*System),
		Systems_reference:  make(map[*System]*System),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"DiagramStructure": &DiagramStructureUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Link": &LinkUnmarshaller{},

			"LinkAssociationShape": &LinkAssociationShapeUnmarshaller{},

			"Part": &PartUnmarshaller{},

			"PartShape": &PartShapeUnmarshaller{},

			"System": &SystemUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "DiagramStructure"},
			{name: "Library"},
			{name: "Link"},
			{name: "LinkAssociationShape"},
			{name: "Part"},
			{name: "PartShape"},
			{name: "System"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramStructure:
		return stage.DiagramStructure_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Link:
		return stage.Link_stagedOrder[instance]
	case *LinkAssociationShape:
		return stage.LinkAssociationShape_stagedOrder[instance]
	case *Part:
		return stage.Part_stagedOrder[instance]
	case *PartShape:
		return stage.PartShape_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *DiagramStructure:
		return any(stage.DiagramStructure_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Link:
		return any(stage.Link_orderStaged[order]).(Type)
	case *LinkAssociationShape:
		return any(stage.LinkAssociationShape_orderStaged[order]).(Type)
	case *Part:
		return any(stage.Part_orderStaged[order]).(Type)
	case *PartShape:
		return any(stage.PartShape_orderStaged[order]).(Type)
	case *System:
		return any(stage.System_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DiagramStructure:
		return stage.DiagramStructure_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Link:
		return stage.Link_stagedOrder[instance]
	case *LinkAssociationShape:
		return stage.LinkAssociationShape_stagedOrder[instance]
	case *Part:
		return stage.Part_stagedOrder[instance]
	case *PartShape:
		return stage.PartShape_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
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
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["DiagramStructure"] = len(stage.DiagramStructures)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["LinkAssociationShape"] = len(stage.LinkAssociationShapes)
	stage.Map_GongStructName_InstancesNb["Part"] = len(stage.Parts)
	stage.Map_GongStructName_InstancesNb["PartShape"] = len(stage.PartShapes)
	stage.Map_GongStructName_InstancesNb["System"] = len(stage.Systems)
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
// Stage puts diagramstructure to the model stage
func (diagramstructure *DiagramStructure) Stage(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; !ok {
		stage.DiagramStructures[diagramstructure] = struct{}{}
		stage.DiagramStructure_stagedOrder[diagramstructure] = stage.DiagramStructureOrder
		stage.DiagramStructure_orderStaged[stage.DiagramStructureOrder] = diagramstructure
		stage.DiagramStructureOrder++
	}
	stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure

	return diagramstructure
}

// StagePreserveOrder puts diagramstructure to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramStructureOrder
// - update stage.DiagramStructureOrder accordingly
func (diagramstructure *DiagramStructure) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramStructures[diagramstructure]; !ok {
		stage.DiagramStructures[diagramstructure] = struct{}{}

		if order > stage.DiagramStructureOrder {
			stage.DiagramStructureOrder = order
		}
		stage.DiagramStructure_stagedOrder[diagramstructure] = order
		stage.DiagramStructure_orderStaged[order] = diagramstructure
		stage.DiagramStructureOrder++
	}
	stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure
}

// Unstage removes diagramstructure off the model stage
func (diagramstructure *DiagramStructure) Unstage(stage *Stage) *DiagramStructure {
	delete(stage.DiagramStructures, diagramstructure)
	// issue1150
	// delete(stage.DiagramStructure_stagedOrder, diagramstructure)
	delete(stage.DiagramStructures_mapString, diagramstructure.Name)

	return diagramstructure
}

// UnstageVoid removes diagramstructure off the model stage
func (diagramstructure *DiagramStructure) UnstageVoid(stage *Stage) {
	delete(stage.DiagramStructures, diagramstructure)
	// issue1150
	// delete(stage.DiagramStructure_stagedOrder, diagramstructure)
	delete(stage.DiagramStructures_mapString, diagramstructure.Name)
}

// commit diagramstructure to the back repo (if it is already staged)
func (diagramstructure *DiagramStructure) Commit(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramStructure(diagramstructure)
		}
	}
	return diagramstructure
}

func (diagramstructure *DiagramStructure) CommitVoid(stage *Stage) {
	diagramstructure.Commit(stage)
}

func (diagramstructure *DiagramStructure) StageVoid(stage *Stage) {
	diagramstructure.Stage(stage)
}

// Checkout diagramstructure to the back repo (if it is already staged)
func (diagramstructure *DiagramStructure) Checkout(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramStructure(diagramstructure)
		}
	}
	return diagramstructure
}

// for satisfaction of GongStruct interface
func (diagramstructure *DiagramStructure) GetName() (res string) {
	return diagramstructure.Name
}

// for satisfaction of GongStruct interface
func (diagramstructure *DiagramStructure) SetName(name string) {
	diagramstructure.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}
		stage.Library_stagedOrder[library] = stage.LibraryOrder
		stage.Library_orderStaged[stage.LibraryOrder] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library

	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}

		if order > stage.LibraryOrder {
			stage.LibraryOrder = order
		}
		stage.Library_stagedOrder[library] = order
		stage.Library_orderStaged[order] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)

	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)
}

// commit library to the back repo (if it is already staged)
func (library *Library) Commit(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLibrary(library)
		}
	}
	return library
}

func (library *Library) CommitVoid(stage *Stage) {
	library.Commit(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// Checkout library to the back repo (if it is already staged)
func (library *Library) Checkout(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLibrary(library)
		}
	}
	return library
}

// for satisfaction of GongStruct interface
func (library *Library) GetName() (res string) {
	return library.Name
}

// for satisfaction of GongStruct interface
func (library *Library) SetName(name string) {
	library.Name = name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {
	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = struct{}{}
		stage.Link_stagedOrder[link] = stage.LinkOrder
		stage.Link_orderStaged[stage.LinkOrder] = link
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link

	return link
}

// StagePreserveOrder puts link to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkOrder
// - update stage.LinkOrder accordingly
func (link *Link) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = struct{}{}

		if order > stage.LinkOrder {
			stage.LinkOrder = order
		}
		stage.Link_stagedOrder[link] = order
		stage.Link_orderStaged[order] = link
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	delete(stage.Links, link)
	// issue1150
	// delete(stage.Link_stagedOrder, link)
	delete(stage.Links_mapString, link.Name)

	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	delete(stage.Links, link)
	// issue1150
	// delete(stage.Link_stagedOrder, link)
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

func (link *Link) StageVoid(stage *Stage) {
	link.Stage(stage)
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

// for satisfaction of GongStruct interface
func (link *Link) SetName(name string) {
	link.Name = name
}

// Stage puts linkassociationshape to the model stage
func (linkassociationshape *LinkAssociationShape) Stage(stage *Stage) *LinkAssociationShape {
	if _, ok := stage.LinkAssociationShapes[linkassociationshape]; !ok {
		stage.LinkAssociationShapes[linkassociationshape] = struct{}{}
		stage.LinkAssociationShape_stagedOrder[linkassociationshape] = stage.LinkAssociationShapeOrder
		stage.LinkAssociationShape_orderStaged[stage.LinkAssociationShapeOrder] = linkassociationshape
		stage.LinkAssociationShapeOrder++
	}
	stage.LinkAssociationShapes_mapString[linkassociationshape.Name] = linkassociationshape

	return linkassociationshape
}

// StagePreserveOrder puts linkassociationshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkAssociationShapeOrder
// - update stage.LinkAssociationShapeOrder accordingly
func (linkassociationshape *LinkAssociationShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.LinkAssociationShapes[linkassociationshape]; !ok {
		stage.LinkAssociationShapes[linkassociationshape] = struct{}{}

		if order > stage.LinkAssociationShapeOrder {
			stage.LinkAssociationShapeOrder = order
		}
		stage.LinkAssociationShape_stagedOrder[linkassociationshape] = order
		stage.LinkAssociationShape_orderStaged[order] = linkassociationshape
		stage.LinkAssociationShapeOrder++
	}
	stage.LinkAssociationShapes_mapString[linkassociationshape.Name] = linkassociationshape
}

// Unstage removes linkassociationshape off the model stage
func (linkassociationshape *LinkAssociationShape) Unstage(stage *Stage) *LinkAssociationShape {
	delete(stage.LinkAssociationShapes, linkassociationshape)
	// issue1150
	// delete(stage.LinkAssociationShape_stagedOrder, linkassociationshape)
	delete(stage.LinkAssociationShapes_mapString, linkassociationshape.Name)

	return linkassociationshape
}

// UnstageVoid removes linkassociationshape off the model stage
func (linkassociationshape *LinkAssociationShape) UnstageVoid(stage *Stage) {
	delete(stage.LinkAssociationShapes, linkassociationshape)
	// issue1150
	// delete(stage.LinkAssociationShape_stagedOrder, linkassociationshape)
	delete(stage.LinkAssociationShapes_mapString, linkassociationshape.Name)
}

// commit linkassociationshape to the back repo (if it is already staged)
func (linkassociationshape *LinkAssociationShape) Commit(stage *Stage) *LinkAssociationShape {
	if _, ok := stage.LinkAssociationShapes[linkassociationshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLinkAssociationShape(linkassociationshape)
		}
	}
	return linkassociationshape
}

func (linkassociationshape *LinkAssociationShape) CommitVoid(stage *Stage) {
	linkassociationshape.Commit(stage)
}

func (linkassociationshape *LinkAssociationShape) StageVoid(stage *Stage) {
	linkassociationshape.Stage(stage)
}

// Checkout linkassociationshape to the back repo (if it is already staged)
func (linkassociationshape *LinkAssociationShape) Checkout(stage *Stage) *LinkAssociationShape {
	if _, ok := stage.LinkAssociationShapes[linkassociationshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLinkAssociationShape(linkassociationshape)
		}
	}
	return linkassociationshape
}

// for satisfaction of GongStruct interface
func (linkassociationshape *LinkAssociationShape) GetName() (res string) {
	return linkassociationshape.Name
}

// for satisfaction of GongStruct interface
func (linkassociationshape *LinkAssociationShape) SetName(name string) {
	linkassociationshape.Name = name
}

// Stage puts part to the model stage
func (part *Part) Stage(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; !ok {
		stage.Parts[part] = struct{}{}
		stage.Part_stagedOrder[part] = stage.PartOrder
		stage.Part_orderStaged[stage.PartOrder] = part
		stage.PartOrder++
	}
	stage.Parts_mapString[part.Name] = part

	return part
}

// StagePreserveOrder puts part to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartOrder
// - update stage.PartOrder accordingly
func (part *Part) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Parts[part]; !ok {
		stage.Parts[part] = struct{}{}

		if order > stage.PartOrder {
			stage.PartOrder = order
		}
		stage.Part_stagedOrder[part] = order
		stage.Part_orderStaged[order] = part
		stage.PartOrder++
	}
	stage.Parts_mapString[part.Name] = part
}

// Unstage removes part off the model stage
func (part *Part) Unstage(stage *Stage) *Part {
	delete(stage.Parts, part)
	// issue1150
	// delete(stage.Part_stagedOrder, part)
	delete(stage.Parts_mapString, part.Name)

	return part
}

// UnstageVoid removes part off the model stage
func (part *Part) UnstageVoid(stage *Stage) {
	delete(stage.Parts, part)
	// issue1150
	// delete(stage.Part_stagedOrder, part)
	delete(stage.Parts_mapString, part.Name)
}

// commit part to the back repo (if it is already staged)
func (part *Part) Commit(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPart(part)
		}
	}
	return part
}

func (part *Part) CommitVoid(stage *Stage) {
	part.Commit(stage)
}

func (part *Part) StageVoid(stage *Stage) {
	part.Stage(stage)
}

// Checkout part to the back repo (if it is already staged)
func (part *Part) Checkout(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPart(part)
		}
	}
	return part
}

// for satisfaction of GongStruct interface
func (part *Part) GetName() (res string) {
	return part.Name
}

// for satisfaction of GongStruct interface
func (part *Part) SetName(name string) {
	part.Name = name
}

// Stage puts partshape to the model stage
func (partshape *PartShape) Stage(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; !ok {
		stage.PartShapes[partshape] = struct{}{}
		stage.PartShape_stagedOrder[partshape] = stage.PartShapeOrder
		stage.PartShape_orderStaged[stage.PartShapeOrder] = partshape
		stage.PartShapeOrder++
	}
	stage.PartShapes_mapString[partshape.Name] = partshape

	return partshape
}

// StagePreserveOrder puts partshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartShapeOrder
// - update stage.PartShapeOrder accordingly
func (partshape *PartShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PartShapes[partshape]; !ok {
		stage.PartShapes[partshape] = struct{}{}

		if order > stage.PartShapeOrder {
			stage.PartShapeOrder = order
		}
		stage.PartShape_stagedOrder[partshape] = order
		stage.PartShape_orderStaged[order] = partshape
		stage.PartShapeOrder++
	}
	stage.PartShapes_mapString[partshape.Name] = partshape
}

// Unstage removes partshape off the model stage
func (partshape *PartShape) Unstage(stage *Stage) *PartShape {
	delete(stage.PartShapes, partshape)
	// issue1150
	// delete(stage.PartShape_stagedOrder, partshape)
	delete(stage.PartShapes_mapString, partshape.Name)

	return partshape
}

// UnstageVoid removes partshape off the model stage
func (partshape *PartShape) UnstageVoid(stage *Stage) {
	delete(stage.PartShapes, partshape)
	// issue1150
	// delete(stage.PartShape_stagedOrder, partshape)
	delete(stage.PartShapes_mapString, partshape.Name)
}

// commit partshape to the back repo (if it is already staged)
func (partshape *PartShape) Commit(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPartShape(partshape)
		}
	}
	return partshape
}

func (partshape *PartShape) CommitVoid(stage *Stage) {
	partshape.Commit(stage)
}

func (partshape *PartShape) StageVoid(stage *Stage) {
	partshape.Stage(stage)
}

// Checkout partshape to the back repo (if it is already staged)
func (partshape *PartShape) Checkout(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPartShape(partshape)
		}
	}
	return partshape
}

// for satisfaction of GongStruct interface
func (partshape *PartShape) GetName() (res string) {
	return partshape.Name
}

// for satisfaction of GongStruct interface
func (partshape *PartShape) SetName(name string) {
	partshape.Name = name
}

// Stage puts system to the model stage
func (system *System) Stage(stage *Stage) *System {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}
		stage.System_stagedOrder[system] = stage.SystemOrder
		stage.System_orderStaged[stage.SystemOrder] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system

	return system
}

// StagePreserveOrder puts system to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemOrder
// - update stage.SystemOrder accordingly
func (system *System) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}

		if order > stage.SystemOrder {
			stage.SystemOrder = order
		}
		stage.System_stagedOrder[system] = order
		stage.System_orderStaged[order] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system
}

// Unstage removes system off the model stage
func (system *System) Unstage(stage *Stage) *System {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)

	return system
}

// UnstageVoid removes system off the model stage
func (system *System) UnstageVoid(stage *Stage) {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)
}

// commit system to the back repo (if it is already staged)
func (system *System) Commit(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystem(system)
		}
	}
	return system
}

func (system *System) CommitVoid(stage *Stage) {
	system.Commit(stage)
}

func (system *System) StageVoid(stage *Stage) {
	system.Stage(stage)
}

// Checkout system to the back repo (if it is already staged)
func (system *System) Checkout(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystem(system)
		}
	}
	return system
}

// for satisfaction of GongStruct interface
func (system *System) GetName() (res string) {
	return system.Name
}

// for satisfaction of GongStruct interface
func (system *System) SetName(name string) {
	system.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDiagramStructure(DiagramStructure *DiagramStructure)
	CreateORMLibrary(Library *Library)
	CreateORMLink(Link *Link)
	CreateORMLinkAssociationShape(LinkAssociationShape *LinkAssociationShape)
	CreateORMPart(Part *Part)
	CreateORMPartShape(PartShape *PartShape)
	CreateORMSystem(System *System)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDiagramStructure(DiagramStructure *DiagramStructure)
	DeleteORMLibrary(Library *Library)
	DeleteORMLink(Link *Link)
	DeleteORMLinkAssociationShape(LinkAssociationShape *LinkAssociationShape)
	DeleteORMPart(Part *Part)
	DeleteORMPartShape(PartShape *PartShape)
	DeleteORMSystem(System *System)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.DiagramStructures = make(map[*DiagramStructure]struct{})
	stage.DiagramStructures_mapString = make(map[string]*DiagramStructure)
	stage.DiagramStructure_stagedOrder = make(map[*DiagramStructure]uint)
	stage.DiagramStructureOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Links = make(map[*Link]struct{})
	stage.Links_mapString = make(map[string]*Link)
	stage.Link_stagedOrder = make(map[*Link]uint)
	stage.LinkOrder = 0

	stage.LinkAssociationShapes = make(map[*LinkAssociationShape]struct{})
	stage.LinkAssociationShapes_mapString = make(map[string]*LinkAssociationShape)
	stage.LinkAssociationShape_stagedOrder = make(map[*LinkAssociationShape]uint)
	stage.LinkAssociationShapeOrder = 0

	stage.Parts = make(map[*Part]struct{})
	stage.Parts_mapString = make(map[string]*Part)
	stage.Part_stagedOrder = make(map[*Part]uint)
	stage.PartOrder = 0

	stage.PartShapes = make(map[*PartShape]struct{})
	stage.PartShapes_mapString = make(map[string]*PartShape)
	stage.PartShape_stagedOrder = make(map[*PartShape]uint)
	stage.PartShapeOrder = 0

	stage.Systems = make(map[*System]struct{})
	stage.Systems_mapString = make(map[string]*System)
	stage.System_stagedOrder = make(map[*System]uint)
	stage.SystemOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.DiagramStructures = nil
	stage.DiagramStructures_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

	stage.LinkAssociationShapes = nil
	stage.LinkAssociationShapes_mapString = nil

	stage.Parts = nil
	stage.Parts_mapString = nil

	stage.PartShapes = nil
	stage.PartShapes_mapString = nil

	stage.Systems = nil
	stage.Systems_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for diagramstructure := range stage.DiagramStructures {
		diagramstructure.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
	}

	for linkassociationshape := range stage.LinkAssociationShapes {
		linkassociationshape.Unstage(stage)
	}

	for part := range stage.Parts {
		part.Unstage(stage)
	}

	for partshape := range stage.PartShapes {
		partshape.Unstage(stage)
	}

	for system := range stage.Systems {
		system.Unstage(stage)
	}

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

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
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
	GongGetUUID(stage *Stage) string
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
	case map[*DiagramStructure]any:
		return any(&stage.DiagramStructures).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
	case map[*LinkAssociationShape]any:
		return any(&stage.LinkAssociationShapes).(*Type)
	case map[*Part]any:
		return any(&stage.Parts).(*Type)
	case map[*PartShape]any:
		return any(&stage.PartShapes).(*Type)
	case map[*System]any:
		return any(&stage.Systems).(*Type)
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
	case *DiagramStructure:
		return any(stage.DiagramStructures_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Link:
		return any(stage.Links_mapString).(map[string]Type)
	case *LinkAssociationShape:
		return any(stage.LinkAssociationShapes_mapString).(map[string]Type)
	case *Part:
		return any(stage.Parts_mapString).(map[string]Type)
	case *PartShape:
		return any(stage.PartShapes_mapString).(map[string]Type)
	case *System:
		return any(stage.Systems_mapString).(map[string]Type)
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
	case DiagramStructure:
		return any(&stage.DiagramStructures).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Link:
		return any(&stage.Links).(*map[*Type]struct{})
	case LinkAssociationShape:
		return any(&stage.LinkAssociationShapes).(*map[*Type]struct{})
	case Part:
		return any(&stage.Parts).(*map[*Type]struct{})
	case PartShape:
		return any(&stage.PartShapes).(*map[*Type]struct{})
	case System:
		return any(&stage.Systems).(*map[*Type]struct{})
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
	case *DiagramStructure:
		return any(&stage.DiagramStructures).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Link:
		return any(&stage.Links).(*map[Type]struct{})
	case *LinkAssociationShape:
		return any(&stage.LinkAssociationShapes).(*map[Type]struct{})
	case *Part:
		return any(&stage.Parts).(*map[Type]struct{})
	case *PartShape:
		return any(&stage.PartShapes).(*map[Type]struct{})
	case *System:
		return any(&stage.Systems).(*map[Type]struct{})
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
	case DiagramStructure:
		return any(&stage.DiagramStructures_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
	case LinkAssociationShape:
		return any(&stage.LinkAssociationShapes_mapString).(*map[string]*Type)
	case Part:
		return any(&stage.Parts_mapString).(*map[string]*Type)
	case PartShape:
		return any(&stage.PartShapes_mapString).(*map[string]*Type)
	case System:
		return any(&stage.Systems_mapString).(*map[string]*Type)
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
	case DiagramStructure:
		return any(&DiagramStructure{
			// Initialisation of associations
			// field is initialized with an instance of PartShape with the name of the field
			Part_Shapes: []*PartShape{{Name: "Part_Shapes"}},
			// field is initialized with an instance of Part with the name of the field
			PartsWhoseNodeIsExpanded: []*Part{{Name: "PartsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of LinkAssociationShape with the name of the field
			Link_Shapes: []*LinkAssociationShape{{Name: "Link_Shapes"}},
			// field is initialized with an instance of Link with the name of the field
			LinksWhoseNodeIsExpanded: []*Link{{Name: "LinksWhoseNodeIsExpanded"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			RootSystems: []*System{{Name: "RootSystems"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Link:
		return any(&Link{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Source: &Part{Name: "Source"},
			// field is initialized with an instance of Part with the name of the field
			Target: &Part{Name: "Target"},
		}).(*Type)
	case LinkAssociationShape:
		return any(&LinkAssociationShape{
			// Initialisation of associations
			// field is initialized with an instance of Link with the name of the field
			Link: &Link{Name: "Link"},
		}).(*Type)
	case Part:
		return any(&Part{
			// Initialisation of associations
		}).(*Type)
	case PartShape:
		return any(&PartShape{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Part: &Part{Name: "Part"},
		}).(*Type)
	case System:
		return any(&System{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Parts: []*Part{{Name: "Parts"}},
			// field is initialized with an instance of Part with the name of the field
			PartsWhoseNodeIsExpanded: []*Part{{Name: "PartsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			SubSystems: []*System{{Name: "SubSystems"}},
			// field is initialized with an instance of System with the name of the field
			SubSystemsWhoseNodeIsExpanded: []*System{{Name: "SubSystemsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Link with the name of the field
			Links: []*Link{{Name: "Links"}},
			// field is initialized with an instance of Link with the name of the field
			LinksWhoseNodeIsExpanded: []*Link{{Name: "LinksWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DiagramStructure with the name of the field
			DiagramStructures: []*DiagramStructure{{Name: "DiagramStructures"}},
			// field is initialized with an instance of DiagramStructure with the name of the field
			DiagramStructuresWhoseNodeIsExpanded: []*DiagramStructure{{Name: "DiagramStructuresWhoseNodeIsExpanded"}},
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
	// reverse maps of direct associations of DiagramStructure
	case DiagramStructure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		case "Source":
			res := make(map[*Part][]*Link)
			for link := range stage.Links {
				if link.Source != nil {
					part_ := link.Source
					var links []*Link
					_, ok := res[part_]
					if ok {
						links = res[part_]
					} else {
						links = make([]*Link, 0)
					}
					links = append(links, link)
					res[part_] = links
				}
			}
			return any(res).(map[*End][]*Start)
		case "Target":
			res := make(map[*Part][]*Link)
			for link := range stage.Links {
				if link.Target != nil {
					part_ := link.Target
					var links []*Link
					_, ok := res[part_]
					if ok {
						links = res[part_]
					} else {
						links = make([]*Link, 0)
					}
					links = append(links, link)
					res[part_] = links
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinkAssociationShape
	case LinkAssociationShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Link":
			res := make(map[*Link][]*LinkAssociationShape)
			for linkassociationshape := range stage.LinkAssociationShapes {
				if linkassociationshape.Link != nil {
					link_ := linkassociationshape.Link
					var linkassociationshapes []*LinkAssociationShape
					_, ok := res[link_]
					if ok {
						linkassociationshapes = res[link_]
					} else {
						linkassociationshapes = make([]*LinkAssociationShape, 0)
					}
					linkassociationshapes = append(linkassociationshapes, linkassociationshape)
					res[link_] = linkassociationshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartShape
	case PartShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*Part][]*PartShape)
			for partshape := range stage.PartShapes {
				if partshape.Part != nil {
					part_ := partshape.Part
					var partshapes []*PartShape
					_, ok := res[part_]
					if ok {
						partshapes = res[part_]
					} else {
						partshapes = make([]*PartShape, 0)
					}
					partshapes = append(partshapes, partshape)
					res[part_] = partshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of System
	case System:
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
	// reverse maps of direct associations of DiagramStructure
	case DiagramStructure:
		switch fieldname {
		// insertion point for per direct association field
		case "Part_Shapes":
			res := make(map[*PartShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, partshape_ := range diagramstructure.Part_Shapes {
					res[partshape_] = append(res[partshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartsWhoseNodeIsExpanded":
			res := make(map[*Part][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, part_ := range diagramstructure.PartsWhoseNodeIsExpanded {
					res[part_] = append(res[part_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Link_Shapes":
			res := make(map[*LinkAssociationShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, linkassociationshape_ := range diagramstructure.Link_Shapes {
					res[linkassociationshape_] = append(res[linkassociationshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "LinksWhoseNodeIsExpanded":
			res := make(map[*Link][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, link_ := range diagramstructure.LinksWhoseNodeIsExpanded {
					res[link_] = append(res[link_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootSystems":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.RootSystems {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LinkAssociationShape
	case LinkAssociationShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartShape
	case PartShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		case "Parts":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.Parts {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartsWhoseNodeIsExpanded":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.PartsWhoseNodeIsExpanded {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystems":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystems {
					res[system_] = append(res[system_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Links":
			res := make(map[*Link][]*System)
			for system := range stage.Systems {
				for _, link_ := range system.Links {
					res[link_] = append(res[link_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "LinksWhoseNodeIsExpanded":
			res := make(map[*Link][]*System)
			for system := range stage.Systems {
				for _, link_ := range system.LinksWhoseNodeIsExpanded {
					res[link_] = append(res[link_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramStructures":
			res := make(map[*DiagramStructure][]*System)
			for system := range stage.Systems {
				for _, diagramstructure_ := range system.DiagramStructures {
					res[diagramstructure_] = append(res[diagramstructure_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramStructuresWhoseNodeIsExpanded":
			res := make(map[*DiagramStructure][]*System)
			for system := range stage.Systems {
				for _, diagramstructure_ := range system.DiagramStructuresWhoseNodeIsExpanded {
					res[diagramstructure_] = append(res[diagramstructure_], system)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *DiagramStructure:
		res = "DiagramStructure"
	case *Library:
		res = "Library"
	case *Link:
		res = "Link"
	case *LinkAssociationShape:
		res = "LinkAssociationShape"
	case *Part:
		res = "Part"
	case *PartShape:
		res = "PartShape"
	case *System:
		res = "System"
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
	case *DiagramStructure:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramStructures"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramStructuresWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "LinksWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Links"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "LinksWhoseNodeIsExpanded"
		res = append(res, rf)
	case *LinkAssociationShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Link_Shapes"
		res = append(res, rf)
	case *Part:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "PartsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Parts"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "PartsWhoseNodeIsExpanded"
		res = append(res, rf)
	case *PartShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Part_Shapes"
		res = append(res, rf)
	case *System:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "RootSystems"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystems"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystemsWhoseNodeIsExpanded"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (diagramstructure *DiagramStructure) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "Part_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PartShape",
		},
		{
			Name:               "IsPartsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PartsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "Link_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinkAssociationShape",
		},
		{
			Name:               "IsLinksNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LinksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Link",
		},
	}
	return
}

func (library *Library) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "IsSubLibrariesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibrariesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LogoSVGFile",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "RootSystems",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsSystemsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
	}
	return
}

func (link *Link) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "Source",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "Target",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
	}
	return
}

func (linkassociationshape *LinkAssociationShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Link",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Link",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (part *Part) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
	}
	return
}

func (partshape *PartShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "WidthWeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "OverideLayoutDirection",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
	}
	return
}

func (system *System) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "Parts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:               "IsPartsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PartsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "SubSystems",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsSubSystemsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubSystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "Links",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Link",
		},
		{
			Name:               "IsLinksNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LinksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Link",
		},
		{
			Name:                 "DiagramStructures",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramStructure",
		},
		{
			Name:               "IsDiagramStructuresNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DiagramStructuresWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramStructure",
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
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
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
func (diagramstructure *DiagramStructure) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramstructure.Name
	case "ComputedPrefix":
		res.valueString = diagramstructure.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsExpanded)
		res.valueBool = diagramstructure.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := diagramstructure.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsChecked)
		res.valueBool = diagramstructure.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsEditable_)
		res.valueBool = diagramstructure.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsShowPrefix)
		res.valueBool = diagramstructure.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramstructure.Width)
		res.valueFloat = diagramstructure.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramstructure.Height)
		res.valueFloat = diagramstructure.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramstructure.DefaultBoxWidth)
		res.valueFloat = diagramstructure.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramstructure.DefaultBoxHeigth)
		res.valueFloat = diagramstructure.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Part_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Part_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPartsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsPartsNodeExpanded)
		res.valueBool = diagramstructure.IsPartsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PartsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.PartsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Link_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Link_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsLinksNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsLinksNodeExpanded)
		res.valueBool = diagramstructure.IsLinksNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LinksWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.LinksWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
	case "SubLibraries":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibraries {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubLibrariesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSubLibrariesNodeExpanded)
		res.valueBool = library.IsSubLibrariesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibrariesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibrariesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", library.NbPixPerCharacter)
		res.valueFloat = library.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LogoSVGFile":
		res.valueString = library.LogoSVGFile
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RootSystems":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootSystems {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSystemsNodeExpanded)
		res.valueBool = library.IsSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (link *Link) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = link.Name
	case "ComputedPrefix":
		res.valueString = link.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", link.IsExpanded)
		res.valueBool = link.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := link.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Source":
		res.GongFieldValueType = GongFieldValueTypePointer
		if link.Source != nil {
			res.valueString = link.Source.Name
			res.ids = link.Source.GongGetUUID(stage)
		}
	case "Target":
		res.GongFieldValueType = GongFieldValueTypePointer
		if link.Target != nil {
			res.valueString = link.Target.Name
			res.ids = link.Target.GongGetUUID(stage)
		}
	}
	return
}

func (linkassociationshape *LinkAssociationShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linkassociationshape.Name
	case "Link":
		res.GongFieldValueType = GongFieldValueTypePointer
		if linkassociationshape.Link != nil {
			res.valueString = linkassociationshape.Link.Name
			res.ids = linkassociationshape.Link.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", linkassociationshape.StartRatio)
		res.valueFloat = linkassociationshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", linkassociationshape.EndRatio)
		res.valueFloat = linkassociationshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := linkassociationshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := linkassociationshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", linkassociationshape.CornerOffsetRatio)
		res.valueFloat = linkassociationshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", linkassociationshape.IsHidden)
		res.valueBool = linkassociationshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (part *Part) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = part.Name
	case "ComputedPrefix":
		res.valueString = part.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", part.IsExpanded)
		res.valueBool = part.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := part.LayoutDirection
		res.valueString = enum.ToCodeString()
	}
	return
}

func (partshape *PartShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partshape.Name
	case "Part":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partshape.Part != nil {
			res.valueString = partshape.Part.Name
			res.ids = partshape.Part.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", partshape.IsExpanded)
		res.valueBool = partshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", partshape.X)
		res.valueFloat = partshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", partshape.Y)
		res.valueFloat = partshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", partshape.Width)
		res.valueFloat = partshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", partshape.Height)
		res.valueFloat = partshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", partshape.IsHidden)
		res.valueBool = partshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "WidthWeight":
		res.valueString = fmt.Sprintf("%f", partshape.WidthWeight)
		res.valueFloat = partshape.WidthWeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "OverideLayoutDirection":
		res.valueString = fmt.Sprintf("%t", partshape.OverideLayoutDirection)
		res.valueBool = partshape.OverideLayoutDirection
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := partshape.LayoutDirection
		res.valueString = enum.ToCodeString()
	}
	return
}

func (system *System) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = system.Name
	case "ComputedPrefix":
		res.valueString = system.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsExpanded)
		res.valueBool = system.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := system.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Parts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Parts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPartsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsPartsNodeExpanded)
		res.valueBool = system.IsPartsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PartsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.PartsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubSystems":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystems {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsSubSystemsNodeExpanded)
		res.valueBool = system.IsSubSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubSystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Links":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Links {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsLinksNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsLinksNodeExpanded)
		res.valueBool = system.IsLinksNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LinksWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.LinksWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramStructures":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramStructures {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDiagramStructuresNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsDiagramStructuresNodeExpanded)
		res.valueBool = system.IsDiagramStructuresNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DiagramStructuresWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramStructuresWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (diagramstructure *DiagramStructure) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramstructure.Name = value.GetValueString()
	case "ComputedPrefix":
		diagramstructure.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramstructure.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		diagramstructure.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsChecked":
		diagramstructure.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramstructure.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramstructure.IsShowPrefix = value.GetValueBool()
	case "Width":
		diagramstructure.Width = value.GetValueFloat()
	case "Height":
		diagramstructure.Height = value.GetValueFloat()
	case "DefaultBoxWidth":
		diagramstructure.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramstructure.DefaultBoxHeigth = value.GetValueFloat()
	case "Part_Shapes":
		diagramstructure.Part_Shapes = make([]*PartShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PartShapes {
					if stage.PartShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Part_Shapes = append(diagramstructure.Part_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsPartsNodeExpanded":
		diagramstructure.IsPartsNodeExpanded = value.GetValueBool()
	case "PartsWhoseNodeIsExpanded":
		diagramstructure.PartsWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						diagramstructure.PartsWhoseNodeIsExpanded = append(diagramstructure.PartsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Link_Shapes":
		diagramstructure.Link_Shapes = make([]*LinkAssociationShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.LinkAssociationShapes {
					if stage.LinkAssociationShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Link_Shapes = append(diagramstructure.Link_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsLinksNodeExpanded":
		diagramstructure.IsLinksNodeExpanded = value.GetValueBool()
	case "LinksWhoseNodeIsExpanded":
		diagramstructure.LinksWhoseNodeIsExpanded = make([]*Link, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Links {
					if stage.Link_stagedOrder[__instance__] == uint(id) {
						diagramstructure.LinksWhoseNodeIsExpanded = append(diagramstructure.LinksWhoseNodeIsExpanded, __instance__)
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

func (library *Library) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		library.Name = value.GetValueString()
	case "SubLibraries":
		library.SubLibraries = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibraries = append(library.SubLibraries, __instance__)
						break
					}
				}
			}
		}
	case "IsSubLibrariesNodeExpanded":
		library.IsSubLibrariesNodeExpanded = value.GetValueBool()
	case "SubLibrariesWhoseNodeIsExpanded":
		library.SubLibrariesWhoseNodeIsExpanded = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibrariesWhoseNodeIsExpanded = append(library.SubLibrariesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "RootSystems":
		library.RootSystems = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.RootSystems = append(library.RootSystems, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemsNodeExpanded":
		library.IsSystemsNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		library.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.SystemsWhoseNodeIsExpanded = append(library.SystemsWhoseNodeIsExpanded, __instance__)
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

func (link *Link) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		link.Name = value.GetValueString()
	case "ComputedPrefix":
		link.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		link.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		link.LayoutDirection.FromCodeString(value.GetValueString())
	case "Source":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			link.Source = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					link.Source = __instance__
					break
				}
			}
		}
	case "Target":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			link.Target = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					link.Target = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (linkassociationshape *LinkAssociationShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		linkassociationshape.Name = value.GetValueString()
	case "Link":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			linkassociationshape.Link = nil
			for __instance__ := range stage.Links {
				if stage.Link_stagedOrder[__instance__] == uint(id) {
					linkassociationshape.Link = __instance__
					break
				}
			}
		}
	case "StartRatio":
		linkassociationshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		linkassociationshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		linkassociationshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		linkassociationshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		linkassociationshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		linkassociationshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (part *Part) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		part.Name = value.GetValueString()
	case "ComputedPrefix":
		part.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		part.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		part.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (partshape *PartShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		partshape.Name = value.GetValueString()
	case "Part":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			partshape.Part = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					partshape.Part = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		partshape.IsExpanded = value.GetValueBool()
	case "X":
		partshape.X = value.GetValueFloat()
	case "Y":
		partshape.Y = value.GetValueFloat()
	case "Width":
		partshape.Width = value.GetValueFloat()
	case "Height":
		partshape.Height = value.GetValueFloat()
	case "IsHidden":
		partshape.IsHidden = value.GetValueBool()
	case "WidthWeight":
		partshape.WidthWeight = value.GetValueFloat()
	case "OverideLayoutDirection":
		partshape.OverideLayoutDirection = value.GetValueBool()
	case "LayoutDirection":
		partshape.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (system *System) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		system.Name = value.GetValueString()
	case "ComputedPrefix":
		system.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		system.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		system.LayoutDirection.FromCodeString(value.GetValueString())
	case "Parts":
		system.Parts = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.Parts = append(system.Parts, __instance__)
						break
					}
				}
			}
		}
	case "IsPartsNodeExpanded":
		system.IsPartsNodeExpanded = value.GetValueBool()
	case "PartsWhoseNodeIsExpanded":
		system.PartsWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.PartsWhoseNodeIsExpanded = append(system.PartsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "SubSystems":
		system.SubSystems = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystems = append(system.SubSystems, __instance__)
						break
					}
				}
			}
		}
	case "IsSubSystemsNodeExpanded":
		system.IsSubSystemsNodeExpanded = value.GetValueBool()
	case "SubSystemsWhoseNodeIsExpanded":
		system.SubSystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystemsWhoseNodeIsExpanded = append(system.SubSystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Links":
		system.Links = make([]*Link, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Links {
					if stage.Link_stagedOrder[__instance__] == uint(id) {
						system.Links = append(system.Links, __instance__)
						break
					}
				}
			}
		}
	case "IsLinksNodeExpanded":
		system.IsLinksNodeExpanded = value.GetValueBool()
	case "LinksWhoseNodeIsExpanded":
		system.LinksWhoseNodeIsExpanded = make([]*Link, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Links {
					if stage.Link_stagedOrder[__instance__] == uint(id) {
						system.LinksWhoseNodeIsExpanded = append(system.LinksWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "DiagramStructures":
		system.DiagramStructures = make([]*DiagramStructure, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramStructures {
					if stage.DiagramStructure_stagedOrder[__instance__] == uint(id) {
						system.DiagramStructures = append(system.DiagramStructures, __instance__)
						break
					}
				}
			}
		}
	case "IsDiagramStructuresNodeExpanded":
		system.IsDiagramStructuresNodeExpanded = value.GetValueBool()
	case "DiagramStructuresWhoseNodeIsExpanded":
		system.DiagramStructuresWhoseNodeIsExpanded = make([]*DiagramStructure, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramStructures {
					if stage.DiagramStructure_stagedOrder[__instance__] == uint(id) {
						system.DiagramStructuresWhoseNodeIsExpanded = append(system.DiagramStructuresWhoseNodeIsExpanded, __instance__)
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
func (diagramstructure *DiagramStructure) GongGetGongstructName() string {
	return "DiagramStructure"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (link *Link) GongGetGongstructName() string {
	return "Link"
}

func (linkassociationshape *LinkAssociationShape) GongGetGongstructName() string {
	return "LinkAssociationShape"
}

func (part *Part) GongGetGongstructName() string {
	return "Part"
}

func (partshape *PartShape) GongGetGongstructName() string {
	return "PartShape"
}

func (system *System) GongGetGongstructName() string {
	return "System"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.DiagramStructures_mapString = make(map[string]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Links_mapString = make(map[string]*Link)
	for link := range stage.Links {
		stage.Links_mapString[link.Name] = link
	}

	stage.LinkAssociationShapes_mapString = make(map[string]*LinkAssociationShape)
	for linkassociationshape := range stage.LinkAssociationShapes {
		stage.LinkAssociationShapes_mapString[linkassociationshape.Name] = linkassociationshape
	}

	stage.Parts_mapString = make(map[string]*Part)
	for part := range stage.Parts {
		stage.Parts_mapString[part.Name] = part
	}

	stage.PartShapes_mapString = make(map[string]*PartShape)
	for partshape := range stage.PartShapes {
		stage.PartShapes_mapString[partshape.Name] = partshape
	}

	stage.Systems_mapString = make(map[string]*System)
	for system := range stage.Systems {
		stage.Systems_mapString[system.Name] = system
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
