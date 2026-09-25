// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

const (
	GongProbeTreeSidebarSuffix           = ":sidebar of the probe"
	GongProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	GongProbeTableSuffix                 = ":table of the probe"
	GongProbeNotificationTableSuffix     = ":notification table of the probe"
	GongProbeFormSuffix                  = ":form of the probe"
	GongProbeSplitSuffix                 = ":probe of the probe"
	GongProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
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
	ArtefactTypes                map[*ArtefactType]struct{}
	ArtefactTypes_instance       map[*ArtefactType]*ArtefactType
	ArtefactTypes_mapString      map[string]*ArtefactType
	ArtefactTypeOrder            uint
	ArtefactType_stagedOrder     map[*ArtefactType]uint
	ArtefactType_orderStaged     map[uint]*ArtefactType
	ArtefactTypes_reference      map[*ArtefactType]*ArtefactType
	ArtefactTypes_referenceOrder map[*ArtefactType]uint

	// insertion point for slice of pointers maps
	OnAfterArtefactTypeCreateCallback GongOnAfterCreateInterface[ArtefactType]
	OnAfterArtefactTypeUpdateCallback GongOnAfterUpdateInterface[ArtefactType]
	OnAfterArtefactTypeDeleteCallback GongOnAfterDeleteInterface[ArtefactType]

	ArtefactTypeShapes                map[*ArtefactTypeShape]struct{}
	ArtefactTypeShapes_instance       map[*ArtefactTypeShape]*ArtefactTypeShape
	ArtefactTypeShapes_mapString      map[string]*ArtefactTypeShape
	ArtefactTypeShapeOrder            uint
	ArtefactTypeShape_stagedOrder     map[*ArtefactTypeShape]uint
	ArtefactTypeShape_orderStaged     map[uint]*ArtefactTypeShape
	ArtefactTypeShapes_reference      map[*ArtefactTypeShape]*ArtefactTypeShape
	ArtefactTypeShapes_referenceOrder map[*ArtefactTypeShape]uint

	// insertion point for slice of pointers maps
	OnAfterArtefactTypeShapeCreateCallback GongOnAfterCreateInterface[ArtefactTypeShape]
	OnAfterArtefactTypeShapeUpdateCallback GongOnAfterUpdateInterface[ArtefactTypeShape]
	OnAfterArtefactTypeShapeDeleteCallback GongOnAfterDeleteInterface[ArtefactTypeShape]

	Artists                map[*Artist]struct{}
	Artists_instance       map[*Artist]*Artist
	Artists_mapString      map[string]*Artist
	ArtistOrder            uint
	Artist_stagedOrder     map[*Artist]uint
	Artist_orderStaged     map[uint]*Artist
	Artists_reference      map[*Artist]*Artist
	Artists_referenceOrder map[*Artist]uint

	// insertion point for slice of pointers maps
	OnAfterArtistCreateCallback GongOnAfterCreateInterface[Artist]
	OnAfterArtistUpdateCallback GongOnAfterUpdateInterface[Artist]
	OnAfterArtistDeleteCallback GongOnAfterDeleteInterface[Artist]

	ArtistShapes                map[*ArtistShape]struct{}
	ArtistShapes_instance       map[*ArtistShape]*ArtistShape
	ArtistShapes_mapString      map[string]*ArtistShape
	ArtistShapeOrder            uint
	ArtistShape_stagedOrder     map[*ArtistShape]uint
	ArtistShape_orderStaged     map[uint]*ArtistShape
	ArtistShapes_reference      map[*ArtistShape]*ArtistShape
	ArtistShapes_referenceOrder map[*ArtistShape]uint

	// insertion point for slice of pointers maps
	OnAfterArtistShapeCreateCallback GongOnAfterCreateInterface[ArtistShape]
	OnAfterArtistShapeUpdateCallback GongOnAfterUpdateInterface[ArtistShape]
	OnAfterArtistShapeDeleteCallback GongOnAfterDeleteInterface[ArtistShape]

	ControlPointShapes                map[*ControlPointShape]struct{}
	ControlPointShapes_instance       map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_mapString      map[string]*ControlPointShape
	ControlPointShapeOrder            uint
	ControlPointShape_stagedOrder     map[*ControlPointShape]uint
	ControlPointShape_orderStaged     map[uint]*ControlPointShape
	ControlPointShapes_reference      map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_referenceOrder map[*ControlPointShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback GongOnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback GongOnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback GongOnAfterDeleteInterface[ControlPointShape]

	Desks                map[*Desk]struct{}
	Desks_instance       map[*Desk]*Desk
	Desks_mapString      map[string]*Desk
	DeskOrder            uint
	Desk_stagedOrder     map[*Desk]uint
	Desk_orderStaged     map[uint]*Desk
	Desks_reference      map[*Desk]*Desk
	Desks_referenceOrder map[*Desk]uint

	// insertion point for slice of pointers maps
	OnAfterDeskCreateCallback GongOnAfterCreateInterface[Desk]
	OnAfterDeskUpdateCallback GongOnAfterUpdateInterface[Desk]
	OnAfterDeskDeleteCallback GongOnAfterDeleteInterface[Desk]

	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_MovementShapes_reverseMap map[*MovementShape]*Diagram

	Diagram_ArtefactTypeShapes_reverseMap map[*ArtefactTypeShape]*Diagram

	Diagram_ArtistShapes_reverseMap map[*ArtistShape]*Diagram

	Diagram_InfluenceShapes_reverseMap map[*InfluenceShape]*Diagram

	OnAfterDiagramCreateCallback GongOnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback GongOnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback GongOnAfterDeleteInterface[Diagram]

	Influences                map[*Influence]struct{}
	Influences_instance       map[*Influence]*Influence
	Influences_mapString      map[string]*Influence
	InfluenceOrder            uint
	Influence_stagedOrder     map[*Influence]uint
	Influence_orderStaged     map[uint]*Influence
	Influences_reference      map[*Influence]*Influence
	Influences_referenceOrder map[*Influence]uint

	// insertion point for slice of pointers maps
	OnAfterInfluenceCreateCallback GongOnAfterCreateInterface[Influence]
	OnAfterInfluenceUpdateCallback GongOnAfterUpdateInterface[Influence]
	OnAfterInfluenceDeleteCallback GongOnAfterDeleteInterface[Influence]

	InfluenceShapes                map[*InfluenceShape]struct{}
	InfluenceShapes_instance       map[*InfluenceShape]*InfluenceShape
	InfluenceShapes_mapString      map[string]*InfluenceShape
	InfluenceShapeOrder            uint
	InfluenceShape_stagedOrder     map[*InfluenceShape]uint
	InfluenceShape_orderStaged     map[uint]*InfluenceShape
	InfluenceShapes_reference      map[*InfluenceShape]*InfluenceShape
	InfluenceShapes_referenceOrder map[*InfluenceShape]uint

	// insertion point for slice of pointers maps
	InfluenceShape_ControlPointShapes_reverseMap map[*ControlPointShape]*InfluenceShape

	OnAfterInfluenceShapeCreateCallback GongOnAfterCreateInterface[InfluenceShape]
	OnAfterInfluenceShapeUpdateCallback GongOnAfterUpdateInterface[InfluenceShape]
	OnAfterInfluenceShapeDeleteCallback GongOnAfterDeleteInterface[InfluenceShape]

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

	OnAfterLibraryCreateCallback GongOnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback GongOnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback GongOnAfterDeleteInterface[Library]

	Movements                map[*Movement]struct{}
	Movements_instance       map[*Movement]*Movement
	Movements_mapString      map[string]*Movement
	MovementOrder            uint
	Movement_stagedOrder     map[*Movement]uint
	Movement_orderStaged     map[uint]*Movement
	Movements_reference      map[*Movement]*Movement
	Movements_referenceOrder map[*Movement]uint

	// insertion point for slice of pointers maps
	Movement_Places_reverseMap map[*Place]*Movement

	OnAfterMovementCreateCallback GongOnAfterCreateInterface[Movement]
	OnAfterMovementUpdateCallback GongOnAfterUpdateInterface[Movement]
	OnAfterMovementDeleteCallback GongOnAfterDeleteInterface[Movement]

	MovementShapes                map[*MovementShape]struct{}
	MovementShapes_instance       map[*MovementShape]*MovementShape
	MovementShapes_mapString      map[string]*MovementShape
	MovementShapeOrder            uint
	MovementShape_stagedOrder     map[*MovementShape]uint
	MovementShape_orderStaged     map[uint]*MovementShape
	MovementShapes_reference      map[*MovementShape]*MovementShape
	MovementShapes_referenceOrder map[*MovementShape]uint

	// insertion point for slice of pointers maps
	OnAfterMovementShapeCreateCallback GongOnAfterCreateInterface[MovementShape]
	OnAfterMovementShapeUpdateCallback GongOnAfterUpdateInterface[MovementShape]
	OnAfterMovementShapeDeleteCallback GongOnAfterDeleteInterface[MovementShape]

	Places                map[*Place]struct{}
	Places_instance       map[*Place]*Place
	Places_mapString      map[string]*Place
	PlaceOrder            uint
	Place_stagedOrder     map[*Place]uint
	Place_orderStaged     map[uint]*Place
	Places_reference      map[*Place]*Place
	Places_referenceOrder map[*Place]uint

	// insertion point for slice of pointers maps
	OnAfterPlaceCreateCallback GongOnAfterCreateInterface[Place]
	OnAfterPlaceUpdateCallback GongOnAfterUpdateInterface[Place]
	OnAfterPlaceDeleteCallback GongOnAfterDeleteInterface[Place]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF

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

type GongStage = Stage

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
	err := stage.ParseAstString(commitToApply, true)
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
	err := stage.ParseAstString(commitToApply, true)
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
	__gong__clearReferences(&stage.ArtefactTypes_reference, &stage.ArtefactTypes_instance, &stage.ArtefactTypes_referenceOrder)

	__gong__clearReferences(&stage.ArtefactTypeShapes_reference, &stage.ArtefactTypeShapes_instance, &stage.ArtefactTypeShapes_referenceOrder)

	__gong__clearReferences(&stage.Artists_reference, &stage.Artists_instance, &stage.Artists_referenceOrder)

	__gong__clearReferences(&stage.ArtistShapes_reference, &stage.ArtistShapes_instance, &stage.ArtistShapes_referenceOrder)

	__gong__clearReferences(&stage.ControlPointShapes_reference, &stage.ControlPointShapes_instance, &stage.ControlPointShapes_referenceOrder)

	__gong__clearReferences(&stage.Desks_reference, &stage.Desks_instance, &stage.Desks_referenceOrder)

	__gong__clearReferences(&stage.Diagrams_reference, &stage.Diagrams_instance, &stage.Diagrams_referenceOrder)

	__gong__clearReferences(&stage.Influences_reference, &stage.Influences_instance, &stage.Influences_referenceOrder)

	__gong__clearReferences(&stage.InfluenceShapes_reference, &stage.InfluenceShapes_instance, &stage.InfluenceShapes_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.Movements_reference, &stage.Movements_instance, &stage.Movements_referenceOrder)

	__gong__clearReferences(&stage.MovementShapes_reference, &stage.MovementShapes_instance, &stage.MovementShapes_referenceOrder)

	__gong__clearReferences(&stage.Places_reference, &stage.Places_instance, &stage.Places_referenceOrder)

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
	stage.ArtefactTypeOrder = __gong__recomputeOrder(stage.ArtefactType_stagedOrder)

	stage.ArtefactTypeShapeOrder = __gong__recomputeOrder(stage.ArtefactTypeShape_stagedOrder)

	stage.ArtistOrder = __gong__recomputeOrder(stage.Artist_stagedOrder)

	stage.ArtistShapeOrder = __gong__recomputeOrder(stage.ArtistShape_stagedOrder)

	stage.ControlPointShapeOrder = __gong__recomputeOrder(stage.ControlPointShape_stagedOrder)

	stage.DeskOrder = __gong__recomputeOrder(stage.Desk_stagedOrder)

	stage.DiagramOrder = __gong__recomputeOrder(stage.Diagram_stagedOrder)

	stage.InfluenceOrder = __gong__recomputeOrder(stage.Influence_stagedOrder)

	stage.InfluenceShapeOrder = __gong__recomputeOrder(stage.InfluenceShape_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.MovementOrder = __gong__recomputeOrder(stage.Movement_stagedOrder)

	stage.MovementShapeOrder = __gong__recomputeOrder(stage.MovementShape_stagedOrder)

	stage.PlaceOrder = __gong__recomputeOrder(stage.Place_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *ArtefactType:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ArtefactTypes, stage.ArtefactType_stagedOrder))
	case *ArtefactTypeShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder))
	case *Artist:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Artists, stage.Artist_stagedOrder))
	case *ArtistShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ArtistShapes, stage.ArtistShape_stagedOrder))
	case *ControlPointShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder))
	case *Desk:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Desks, stage.Desk_stagedOrder))
	case *Diagram:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder))
	case *Influence:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Influences, stage.Influence_stagedOrder))
	case *InfluenceShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *Movement:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Movements, stage.Movement_stagedOrder))
	case *MovementShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MovementShapes, stage.MovementShape_stagedOrder))
	case *Place:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Places, stage.Place_stagedOrder))

	}
	return
}

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
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
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		ArtefactTypes:           make(map[*ArtefactType]struct{}),
		ArtefactTypes_mapString: make(map[string]*ArtefactType),

		ArtefactTypeShapes:           make(map[*ArtefactTypeShape]struct{}),
		ArtefactTypeShapes_mapString: make(map[string]*ArtefactTypeShape),

		Artists:           make(map[*Artist]struct{}),
		Artists_mapString: make(map[string]*Artist),

		ArtistShapes:           make(map[*ArtistShape]struct{}),
		ArtistShapes_mapString: make(map[string]*ArtistShape),

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

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Movements:           make(map[*Movement]struct{}),
		Movements_mapString: make(map[string]*Movement),

		MovementShapes:           make(map[*MovementShape]struct{}),
		MovementShapes_mapString: make(map[string]*MovementShape),

		Places:           make(map[*Place]struct{}),
		Places_mapString: make(map[string]*Place),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ArtefactType_stagedOrder: make(map[*ArtefactType]uint),
		ArtefactType_orderStaged: make(map[uint]*ArtefactType),
		ArtefactTypes_reference:  make(map[*ArtefactType]*ArtefactType),

		ArtefactTypeShape_stagedOrder: make(map[*ArtefactTypeShape]uint),
		ArtefactTypeShape_orderStaged: make(map[uint]*ArtefactTypeShape),
		ArtefactTypeShapes_reference:  make(map[*ArtefactTypeShape]*ArtefactTypeShape),

		Artist_stagedOrder: make(map[*Artist]uint),
		Artist_orderStaged: make(map[uint]*Artist),
		Artists_reference:  make(map[*Artist]*Artist),

		ArtistShape_stagedOrder: make(map[*ArtistShape]uint),
		ArtistShape_orderStaged: make(map[uint]*ArtistShape),
		ArtistShapes_reference:  make(map[*ArtistShape]*ArtistShape),

		ControlPointShape_stagedOrder: make(map[*ControlPointShape]uint),
		ControlPointShape_orderStaged: make(map[uint]*ControlPointShape),
		ControlPointShapes_reference:  make(map[*ControlPointShape]*ControlPointShape),

		Desk_stagedOrder: make(map[*Desk]uint),
		Desk_orderStaged: make(map[uint]*Desk),
		Desks_reference:  make(map[*Desk]*Desk),

		Diagram_stagedOrder: make(map[*Diagram]uint),
		Diagram_orderStaged: make(map[uint]*Diagram),
		Diagrams_reference:  make(map[*Diagram]*Diagram),

		Influence_stagedOrder: make(map[*Influence]uint),
		Influence_orderStaged: make(map[uint]*Influence),
		Influences_reference:  make(map[*Influence]*Influence),

		InfluenceShape_stagedOrder: make(map[*InfluenceShape]uint),
		InfluenceShape_orderStaged: make(map[uint]*InfluenceShape),
		InfluenceShapes_reference:  make(map[*InfluenceShape]*InfluenceShape),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Movement_stagedOrder: make(map[*Movement]uint),
		Movement_orderStaged: make(map[uint]*Movement),
		Movements_reference:  make(map[*Movement]*Movement),

		MovementShape_stagedOrder: make(map[*MovementShape]uint),
		MovementShape_orderStaged: make(map[uint]*MovementShape),
		MovementShapes_reference:  make(map[*MovementShape]*MovementShape),

		Place_stagedOrder: make(map[*Place]uint),
		Place_orderStaged: make(map[uint]*Place),
		Places_reference:  make(map[*Place]*Place),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"ArtefactType": &ArtefactTypeUnmarshaller{},

			"ArtefactTypeShape": &ArtefactTypeShapeUnmarshaller{},

			"Artist": &ArtistUnmarshaller{},

			"ArtistShape": &ArtistShapeUnmarshaller{},

			"ControlPointShape": &ControlPointShapeUnmarshaller{},

			"Desk": &DeskUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"Influence": &InfluenceUnmarshaller{},

			"InfluenceShape": &InfluenceShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Movement": &MovementUnmarshaller{},

			"MovementShape": &MovementShapeUnmarshaller{},

			"Place": &PlaceUnmarshaller{},

			// end of insertion point
		},

		navigationMode: GongNavigationModeNormal,
	}

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *ArtefactType:
		return any(stage.ArtefactType_orderStaged[order]).(Type)
	case *ArtefactTypeShape:
		return any(stage.ArtefactTypeShape_orderStaged[order]).(Type)
	case *Artist:
		return any(stage.Artist_orderStaged[order]).(Type)
	case *ArtistShape:
		return any(stage.ArtistShape_orderStaged[order]).(Type)
	case *ControlPointShape:
		return any(stage.ControlPointShape_orderStaged[order]).(Type)
	case *Desk:
		return any(stage.Desk_orderStaged[order]).(Type)
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *Influence:
		return any(stage.Influence_orderStaged[order]).(Type)
	case *InfluenceShape:
		return any(stage.InfluenceShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Movement:
		return any(stage.Movement_orderStaged[order]).(Type)
	case *MovementShape:
		return any(stage.MovementShape_orderStaged[order]).(Type)
	case *Place:
		return any(stage.Place_orderStaged[order]).(Type)
	default:
		return // should not happen
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
	stage.Map_GongStructName_InstancesNb["ArtefactType"] = len(stage.ArtefactTypes)
	stage.Map_GongStructName_InstancesNb["ArtefactTypeShape"] = len(stage.ArtefactTypeShapes)
	stage.Map_GongStructName_InstancesNb["Artist"] = len(stage.Artists)
	stage.Map_GongStructName_InstancesNb["ArtistShape"] = len(stage.ArtistShapes)
	stage.Map_GongStructName_InstancesNb["ControlPointShape"] = len(stage.ControlPointShapes)
	stage.Map_GongStructName_InstancesNb["Desk"] = len(stage.Desks)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Influence"] = len(stage.Influences)
	stage.Map_GongStructName_InstancesNb["InfluenceShape"] = len(stage.InfluenceShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Movement"] = len(stage.Movements)
	stage.Map_GongStructName_InstancesNb["MovementShape"] = len(stage.MovementShapes)
	stage.Map_GongStructName_InstancesNb["Place"] = len(stage.Places)
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
// Stage puts artefacttype to the model stage
func (artefacttype *ArtefactType) Stage(stage *Stage) *ArtefactType {
	__gong__stage(stage.ArtefactTypes, stage.ArtefactType_stagedOrder, stage.ArtefactType_orderStaged, &stage.ArtefactTypeOrder, stage.ArtefactTypes_mapString, artefacttype, artefacttype.Name)
	return artefacttype
}

// StagePreserveOrder puts artefacttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtefactTypeOrder
// - update stage.ArtefactTypeOrder accordingly
func (artefacttype *ArtefactType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArtefactTypes, stage.ArtefactType_stagedOrder, stage.ArtefactType_orderStaged, &stage.ArtefactTypeOrder, stage.ArtefactTypes_mapString, artefacttype, order, artefacttype.Name)
}

// Unstage removes artefacttype off the model stage
func (artefacttype *ArtefactType) Unstage(stage *Stage) *ArtefactType {
	__gong__unstage(stage.ArtefactTypes, stage.ArtefactTypes_mapString, artefacttype, artefacttype.Name)
	return artefacttype
}

// UnstageVoid removes artefacttype off the model stage
func (artefacttype *ArtefactType) UnstageVoid(stage *Stage) {
	artefacttype.Unstage(stage)
}

func (artefacttype *ArtefactType) StageVoid(stage *Stage) {
	artefacttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (artefacttype *ArtefactType) GetName() (res string) {
	return artefacttype.Name
}

// for satisfaction of GongStruct interface
func (artefacttype *ArtefactType) SetName(name string) {
	artefacttype.Name = name
}

// Stage puts artefacttypeshape to the model stage
func (artefacttypeshape *ArtefactTypeShape) Stage(stage *Stage) *ArtefactTypeShape {
	__gong__stage(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder, stage.ArtefactTypeShape_orderStaged, &stage.ArtefactTypeShapeOrder, stage.ArtefactTypeShapes_mapString, artefacttypeshape, artefacttypeshape.Name)
	return artefacttypeshape
}

// StagePreserveOrder puts artefacttypeshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtefactTypeShapeOrder
// - update stage.ArtefactTypeShapeOrder accordingly
func (artefacttypeshape *ArtefactTypeShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder, stage.ArtefactTypeShape_orderStaged, &stage.ArtefactTypeShapeOrder, stage.ArtefactTypeShapes_mapString, artefacttypeshape, order, artefacttypeshape.Name)
}

// Unstage removes artefacttypeshape off the model stage
func (artefacttypeshape *ArtefactTypeShape) Unstage(stage *Stage) *ArtefactTypeShape {
	__gong__unstage(stage.ArtefactTypeShapes, stage.ArtefactTypeShapes_mapString, artefacttypeshape, artefacttypeshape.Name)
	return artefacttypeshape
}

// UnstageVoid removes artefacttypeshape off the model stage
func (artefacttypeshape *ArtefactTypeShape) UnstageVoid(stage *Stage) {
	artefacttypeshape.Unstage(stage)
}

func (artefacttypeshape *ArtefactTypeShape) StageVoid(stage *Stage) {
	artefacttypeshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (artefacttypeshape *ArtefactTypeShape) GetName() (res string) {
	return artefacttypeshape.Name
}

// for satisfaction of GongStruct interface
func (artefacttypeshape *ArtefactTypeShape) SetName(name string) {
	artefacttypeshape.Name = name
}

// Stage puts artist to the model stage
func (artist *Artist) Stage(stage *Stage) *Artist {
	__gong__stage(stage.Artists, stage.Artist_stagedOrder, stage.Artist_orderStaged, &stage.ArtistOrder, stage.Artists_mapString, artist, artist.Name)
	return artist
}

// StagePreserveOrder puts artist to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtistOrder
// - update stage.ArtistOrder accordingly
func (artist *Artist) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Artists, stage.Artist_stagedOrder, stage.Artist_orderStaged, &stage.ArtistOrder, stage.Artists_mapString, artist, order, artist.Name)
}

// Unstage removes artist off the model stage
func (artist *Artist) Unstage(stage *Stage) *Artist {
	__gong__unstage(stage.Artists, stage.Artists_mapString, artist, artist.Name)
	return artist
}

// UnstageVoid removes artist off the model stage
func (artist *Artist) UnstageVoid(stage *Stage) {
	artist.Unstage(stage)
}

func (artist *Artist) StageVoid(stage *Stage) {
	artist.Stage(stage)
}

// for satisfaction of GongStruct interface
func (artist *Artist) GetName() (res string) {
	return artist.Name
}

// for satisfaction of GongStruct interface
func (artist *Artist) SetName(name string) {
	artist.Name = name
}

// Stage puts artistshape to the model stage
func (artistshape *ArtistShape) Stage(stage *Stage) *ArtistShape {
	__gong__stage(stage.ArtistShapes, stage.ArtistShape_stagedOrder, stage.ArtistShape_orderStaged, &stage.ArtistShapeOrder, stage.ArtistShapes_mapString, artistshape, artistshape.Name)
	return artistshape
}

// StagePreserveOrder puts artistshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtistShapeOrder
// - update stage.ArtistShapeOrder accordingly
func (artistshape *ArtistShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArtistShapes, stage.ArtistShape_stagedOrder, stage.ArtistShape_orderStaged, &stage.ArtistShapeOrder, stage.ArtistShapes_mapString, artistshape, order, artistshape.Name)
}

// Unstage removes artistshape off the model stage
func (artistshape *ArtistShape) Unstage(stage *Stage) *ArtistShape {
	__gong__unstage(stage.ArtistShapes, stage.ArtistShapes_mapString, artistshape, artistshape.Name)
	return artistshape
}

// UnstageVoid removes artistshape off the model stage
func (artistshape *ArtistShape) UnstageVoid(stage *Stage) {
	artistshape.Unstage(stage)
}

func (artistshape *ArtistShape) StageVoid(stage *Stage) {
	artistshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (artistshape *ArtistShape) GetName() (res string) {
	return artistshape.Name
}

// for satisfaction of GongStruct interface
func (artistshape *ArtistShape) SetName(name string) {
	artistshape.Name = name
}

// Stage puts controlpointshape to the model stage
func (controlpointshape *ControlPointShape) Stage(stage *Stage) *ControlPointShape {
	__gong__stage(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// StagePreserveOrder puts controlpointshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointShapeOrder
// - update stage.ControlPointShapeOrder accordingly
func (controlpointshape *ControlPointShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, order, controlpointshape.Name)
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	__gong__unstage(stage.ControlPointShapes, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	controlpointshape.Unstage(stage)
}

func (controlpointshape *ControlPointShape) StageVoid(stage *Stage) {
	controlpointshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) GetName() (res string) {
	return controlpointshape.Name
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) SetName(name string) {
	controlpointshape.Name = name
}

// Stage puts desk to the model stage
func (desk *Desk) Stage(stage *Stage) *Desk {
	__gong__stage(stage.Desks, stage.Desk_stagedOrder, stage.Desk_orderStaged, &stage.DeskOrder, stage.Desks_mapString, desk, desk.Name)
	return desk
}

// StagePreserveOrder puts desk to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeskOrder
// - update stage.DeskOrder accordingly
func (desk *Desk) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Desks, stage.Desk_stagedOrder, stage.Desk_orderStaged, &stage.DeskOrder, stage.Desks_mapString, desk, order, desk.Name)
}

// Unstage removes desk off the model stage
func (desk *Desk) Unstage(stage *Stage) *Desk {
	__gong__unstage(stage.Desks, stage.Desks_mapString, desk, desk.Name)
	return desk
}

// UnstageVoid removes desk off the model stage
func (desk *Desk) UnstageVoid(stage *Stage) {
	desk.Unstage(stage)
}

func (desk *Desk) StageVoid(stage *Stage) {
	desk.Stage(stage)
}

// for satisfaction of GongStruct interface
func (desk *Desk) GetName() (res string) {
	return desk.Name
}

// for satisfaction of GongStruct interface
func (desk *Desk) SetName(name string) {
	desk.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {
	__gong__stage(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, order, diagram.Name)
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	__gong__unstage(stage.Diagrams, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	diagram.Unstage(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) GetName() (res string) {
	return diagram.Name
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) SetName(name string) {
	diagram.Name = name
}

// Stage puts influence to the model stage
func (influence *Influence) Stage(stage *Stage) *Influence {
	__gong__stage(stage.Influences, stage.Influence_stagedOrder, stage.Influence_orderStaged, &stage.InfluenceOrder, stage.Influences_mapString, influence, influence.Name)
	return influence
}

// StagePreserveOrder puts influence to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InfluenceOrder
// - update stage.InfluenceOrder accordingly
func (influence *Influence) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Influences, stage.Influence_stagedOrder, stage.Influence_orderStaged, &stage.InfluenceOrder, stage.Influences_mapString, influence, order, influence.Name)
}

// Unstage removes influence off the model stage
func (influence *Influence) Unstage(stage *Stage) *Influence {
	__gong__unstage(stage.Influences, stage.Influences_mapString, influence, influence.Name)
	return influence
}

// UnstageVoid removes influence off the model stage
func (influence *Influence) UnstageVoid(stage *Stage) {
	influence.Unstage(stage)
}

func (influence *Influence) StageVoid(stage *Stage) {
	influence.Stage(stage)
}

// for satisfaction of GongStruct interface
func (influence *Influence) GetName() (res string) {
	return influence.Name
}

// for satisfaction of GongStruct interface
func (influence *Influence) SetName(name string) {
	influence.Name = name
}

// Stage puts influenceshape to the model stage
func (influenceshape *InfluenceShape) Stage(stage *Stage) *InfluenceShape {
	__gong__stage(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder, stage.InfluenceShape_orderStaged, &stage.InfluenceShapeOrder, stage.InfluenceShapes_mapString, influenceshape, influenceshape.Name)
	return influenceshape
}

// StagePreserveOrder puts influenceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InfluenceShapeOrder
// - update stage.InfluenceShapeOrder accordingly
func (influenceshape *InfluenceShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder, stage.InfluenceShape_orderStaged, &stage.InfluenceShapeOrder, stage.InfluenceShapes_mapString, influenceshape, order, influenceshape.Name)
}

// Unstage removes influenceshape off the model stage
func (influenceshape *InfluenceShape) Unstage(stage *Stage) *InfluenceShape {
	__gong__unstage(stage.InfluenceShapes, stage.InfluenceShapes_mapString, influenceshape, influenceshape.Name)
	return influenceshape
}

// UnstageVoid removes influenceshape off the model stage
func (influenceshape *InfluenceShape) UnstageVoid(stage *Stage) {
	influenceshape.Unstage(stage)
}

func (influenceshape *InfluenceShape) StageVoid(stage *Stage) {
	influenceshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (influenceshape *InfluenceShape) GetName() (res string) {
	return influenceshape.Name
}

// for satisfaction of GongStruct interface
func (influenceshape *InfluenceShape) SetName(name string) {
	influenceshape.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	__gong__stage(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, library.Name)
	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, order, library.Name)
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	__gong__unstage(stage.Librarys, stage.Librarys_mapString, library, library.Name)
	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	library.Unstage(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// for satisfaction of GongStruct interface
func (library *Library) GetName() (res string) {
	return library.Name
}

// for satisfaction of GongStruct interface
func (library *Library) SetName(name string) {
	library.Name = name
}

// Stage puts movement to the model stage
func (movement *Movement) Stage(stage *Stage) *Movement {
	__gong__stage(stage.Movements, stage.Movement_stagedOrder, stage.Movement_orderStaged, &stage.MovementOrder, stage.Movements_mapString, movement, movement.Name)
	return movement
}

// StagePreserveOrder puts movement to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MovementOrder
// - update stage.MovementOrder accordingly
func (movement *Movement) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Movements, stage.Movement_stagedOrder, stage.Movement_orderStaged, &stage.MovementOrder, stage.Movements_mapString, movement, order, movement.Name)
}

// Unstage removes movement off the model stage
func (movement *Movement) Unstage(stage *Stage) *Movement {
	__gong__unstage(stage.Movements, stage.Movements_mapString, movement, movement.Name)
	return movement
}

// UnstageVoid removes movement off the model stage
func (movement *Movement) UnstageVoid(stage *Stage) {
	movement.Unstage(stage)
}

func (movement *Movement) StageVoid(stage *Stage) {
	movement.Stage(stage)
}

// for satisfaction of GongStruct interface
func (movement *Movement) GetName() (res string) {
	return movement.Name
}

// for satisfaction of GongStruct interface
func (movement *Movement) SetName(name string) {
	movement.Name = name
}

// Stage puts movementshape to the model stage
func (movementshape *MovementShape) Stage(stage *Stage) *MovementShape {
	__gong__stage(stage.MovementShapes, stage.MovementShape_stagedOrder, stage.MovementShape_orderStaged, &stage.MovementShapeOrder, stage.MovementShapes_mapString, movementshape, movementshape.Name)
	return movementshape
}

// StagePreserveOrder puts movementshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MovementShapeOrder
// - update stage.MovementShapeOrder accordingly
func (movementshape *MovementShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MovementShapes, stage.MovementShape_stagedOrder, stage.MovementShape_orderStaged, &stage.MovementShapeOrder, stage.MovementShapes_mapString, movementshape, order, movementshape.Name)
}

// Unstage removes movementshape off the model stage
func (movementshape *MovementShape) Unstage(stage *Stage) *MovementShape {
	__gong__unstage(stage.MovementShapes, stage.MovementShapes_mapString, movementshape, movementshape.Name)
	return movementshape
}

// UnstageVoid removes movementshape off the model stage
func (movementshape *MovementShape) UnstageVoid(stage *Stage) {
	movementshape.Unstage(stage)
}

func (movementshape *MovementShape) StageVoid(stage *Stage) {
	movementshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (movementshape *MovementShape) GetName() (res string) {
	return movementshape.Name
}

// for satisfaction of GongStruct interface
func (movementshape *MovementShape) SetName(name string) {
	movementshape.Name = name
}

// Stage puts place to the model stage
func (place *Place) Stage(stage *Stage) *Place {
	__gong__stage(stage.Places, stage.Place_stagedOrder, stage.Place_orderStaged, &stage.PlaceOrder, stage.Places_mapString, place, place.Name)
	return place
}

// StagePreserveOrder puts place to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlaceOrder
// - update stage.PlaceOrder accordingly
func (place *Place) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Places, stage.Place_stagedOrder, stage.Place_orderStaged, &stage.PlaceOrder, stage.Places_mapString, place, order, place.Name)
}

// Unstage removes place off the model stage
func (place *Place) Unstage(stage *Stage) *Place {
	__gong__unstage(stage.Places, stage.Places_mapString, place, place.Name)
	return place
}

// UnstageVoid removes place off the model stage
func (place *Place) UnstageVoid(stage *Stage) {
	place.Unstage(stage)
}

func (place *Place) StageVoid(stage *Stage) {
	place.Stage(stage)
}

// for satisfaction of GongStruct interface
func (place *Place) GetName() (res string) {
	return place.Name
}

// for satisfaction of GongStruct interface
func (place *Place) SetName(name string) {
	place.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.ArtefactTypes, &stage.ArtefactTypes_mapString, &stage.ArtefactType_stagedOrder, &stage.ArtefactTypeOrder)

	__gong__resetStageType(&stage.ArtefactTypeShapes, &stage.ArtefactTypeShapes_mapString, &stage.ArtefactTypeShape_stagedOrder, &stage.ArtefactTypeShapeOrder)

	__gong__resetStageType(&stage.Artists, &stage.Artists_mapString, &stage.Artist_stagedOrder, &stage.ArtistOrder)

	__gong__resetStageType(&stage.ArtistShapes, &stage.ArtistShapes_mapString, &stage.ArtistShape_stagedOrder, &stage.ArtistShapeOrder)

	__gong__resetStageType(&stage.ControlPointShapes, &stage.ControlPointShapes_mapString, &stage.ControlPointShape_stagedOrder, &stage.ControlPointShapeOrder)

	__gong__resetStageType(&stage.Desks, &stage.Desks_mapString, &stage.Desk_stagedOrder, &stage.DeskOrder)

	__gong__resetStageType(&stage.Diagrams, &stage.Diagrams_mapString, &stage.Diagram_stagedOrder, &stage.DiagramOrder)

	__gong__resetStageType(&stage.Influences, &stage.Influences_mapString, &stage.Influence_stagedOrder, &stage.InfluenceOrder)

	__gong__resetStageType(&stage.InfluenceShapes, &stage.InfluenceShapes_mapString, &stage.InfluenceShape_stagedOrder, &stage.InfluenceShapeOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.Movements, &stage.Movements_mapString, &stage.Movement_stagedOrder, &stage.MovementOrder)

	__gong__resetStageType(&stage.MovementShapes, &stage.MovementShapes_mapString, &stage.MovementShape_stagedOrder, &stage.MovementShapeOrder)

	__gong__resetStageType(&stage.Places, &stage.Places_mapString, &stage.Place_stagedOrder, &stage.PlaceOrder)

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct any

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ArtefactType:
		return any(stage.ArtefactTypes_mapString).(map[string]Type)
	case *ArtefactTypeShape:
		return any(stage.ArtefactTypeShapes_mapString).(map[string]Type)
	case *Artist:
		return any(stage.Artists_mapString).(map[string]Type)
	case *ArtistShape:
		return any(stage.ArtistShapes_mapString).(map[string]Type)
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
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Movement:
		return any(stage.Movements_mapString).(map[string]Type)
	case *MovementShape:
		return any(stage.MovementShapes_mapString).(map[string]Type)
	case *Place:
		return any(stage.Places_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *ArtefactType:
		return any(&stage.ArtefactTypes).(*map[Type]struct{})
	case *ArtefactTypeShape:
		return any(&stage.ArtefactTypeShapes).(*map[Type]struct{})
	case *Artist:
		return any(&stage.Artists).(*map[Type]struct{})
	case *ArtistShape:
		return any(&stage.ArtistShapes).(*map[Type]struct{})
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
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Movement:
		return any(&stage.Movements).(*map[Type]struct{})
	case *MovementShape:
		return any(&stage.MovementShapes).(*map[Type]struct{})
	case *Place:
		return any(&stage.Places).(*map[Type]struct{})
	default:
		return nil
	}
}

// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case ArtefactTypeShape:
		return any(&ArtefactTypeShape{
			ArtefactType: &ArtefactType{Name: "ArtefactType"},
		}).(*Type)
	case Artist:
		return any(&Artist{
			Place: &Place{Name: "Place"},
		}).(*Type)
	case ArtistShape:
		return any(&ArtistShape{
			Artist: &Artist{Name: "Artist"},
		}).(*Type)
	case Desk:
		return any(&Desk{
			SelectedDiagram: &Diagram{Name: "SelectedDiagram"},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			MovementShapes: []*MovementShape{{Name: "MovementShapes"}},
			ArtefactTypeShapes: []*ArtefactTypeShape{{Name: "ArtefactTypeShapes"}},
			ArtistShapes: []*ArtistShape{{Name: "ArtistShapes"}},
			InfluenceShapes: []*InfluenceShape{{Name: "InfluenceShapes"}},
		}).(*Type)
	case Influence:
		return any(&Influence{
			SourceMovement: &Movement{Name: "SourceMovement"},
			SourceArtefactType: &ArtefactType{Name: "SourceArtefactType"},
			SourceArtist: &Artist{Name: "SourceArtist"},
			TargetMovement: &Movement{Name: "TargetMovement"},
			TargetArtefactType: &ArtefactType{Name: "TargetArtefactType"},
			TargetArtist: &Artist{Name: "TargetArtist"},
		}).(*Type)
	case InfluenceShape:
		return any(&InfluenceShape{
			Influence: &Influence{Name: "Influence"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Movement:
		return any(&Movement{
			Places: []*Place{{Name: "Places"}},
		}).(*Type)
	case MovementShape:
		return any(&MovementShape{
			Movement: &Movement{Name: "Movement"},
		}).(*Type)
	default:
		return &ret
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ArtefactType
	case ArtefactType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArtefactTypeShape
	case ArtefactTypeShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ArtefactType":
			res := make(map[*ArtefactType][]*ArtefactTypeShape)
			for artefacttypeshape := range stage.ArtefactTypeShapes {
				if artefacttypeshape.ArtefactType != nil {
					artefacttype_ := artefacttypeshape.ArtefactType
					var artefacttypeshapes []*ArtefactTypeShape
					_, ok := res[artefacttype_]
					if ok {
						artefacttypeshapes = res[artefacttype_]
					} else {
						artefacttypeshapes = make([]*ArtefactTypeShape, 0)
					}
					artefacttypeshapes = append(artefacttypeshapes, artefacttypeshape)
					res[artefacttype_] = artefacttypeshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Artist
	case Artist:
		switch fieldname {
		// insertion point for per direct association field
		case "Place":
			res := make(map[*Place][]*Artist)
			for artist := range stage.Artists {
				if artist.Place != nil {
					place_ := artist.Place
					var artists []*Artist
					_, ok := res[place_]
					if ok {
						artists = res[place_]
					} else {
						artists = make([]*Artist, 0)
					}
					artists = append(artists, artist)
					res[place_] = artists
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ArtistShape
	case ArtistShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Artist":
			res := make(map[*Artist][]*ArtistShape)
			for artistshape := range stage.ArtistShapes {
				if artistshape.Artist != nil {
					artist_ := artistshape.Artist
					var artistshapes []*ArtistShape
					_, ok := res[artist_]
					if ok {
						artistshapes = res[artist_]
					} else {
						artistshapes = make([]*ArtistShape, 0)
					}
					artistshapes = append(artistshapes, artistshape)
					res[artist_] = artistshapes
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
		case "SourceMovement":
			res := make(map[*Movement][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceMovement != nil {
					movement_ := influence.SourceMovement
					var influences []*Influence
					_, ok := res[movement_]
					if ok {
						influences = res[movement_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[movement_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "SourceArtefactType":
			res := make(map[*ArtefactType][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceArtefactType != nil {
					artefacttype_ := influence.SourceArtefactType
					var influences []*Influence
					_, ok := res[artefacttype_]
					if ok {
						influences = res[artefacttype_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[artefacttype_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "SourceArtist":
			res := make(map[*Artist][]*Influence)
			for influence := range stage.Influences {
				if influence.SourceArtist != nil {
					artist_ := influence.SourceArtist
					var influences []*Influence
					_, ok := res[artist_]
					if ok {
						influences = res[artist_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[artist_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetMovement":
			res := make(map[*Movement][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetMovement != nil {
					movement_ := influence.TargetMovement
					var influences []*Influence
					_, ok := res[movement_]
					if ok {
						influences = res[movement_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[movement_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetArtefactType":
			res := make(map[*ArtefactType][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetArtefactType != nil {
					artefacttype_ := influence.TargetArtefactType
					var influences []*Influence
					_, ok := res[artefacttype_]
					if ok {
						influences = res[artefacttype_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[artefacttype_] = influences
				}
			}
			return any(res).(map[*End][]*Start)
		case "TargetArtist":
			res := make(map[*Artist][]*Influence)
			for influence := range stage.Influences {
				if influence.TargetArtist != nil {
					artist_ := influence.TargetArtist
					var influences []*Influence
					_, ok := res[artist_]
					if ok {
						influences = res[artist_]
					} else {
						influences = make([]*Influence, 0)
					}
					influences = append(influences, influence)
					res[artist_] = influences
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
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Movement
	case Movement:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MovementShape
	case MovementShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Movement":
			res := make(map[*Movement][]*MovementShape)
			for movementshape := range stage.MovementShapes {
				if movementshape.Movement != nil {
					movement_ := movementshape.Movement
					var movementshapes []*MovementShape
					_, ok := res[movement_]
					if ok {
						movementshapes = res[movement_]
					} else {
						movementshapes = make([]*MovementShape, 0)
					}
					movementshapes = append(movementshapes, movementshape)
					res[movement_] = movementshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Place
	case Place:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of ArtefactType
	case ArtefactType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArtefactTypeShape
	case ArtefactTypeShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Artist
	case Artist:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArtistShape
	case ArtistShape:
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
		case "MovementShapes":
			res := make(map[*MovementShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, movementshape_ := range diagram.MovementShapes {
					res[movementshape_] = append(res[movementshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ArtefactTypeShapes":
			res := make(map[*ArtefactTypeShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, artefacttypeshape_ := range diagram.ArtefactTypeShapes {
					res[artefacttypeshape_] = append(res[artefacttypeshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ArtistShapes":
			res := make(map[*ArtistShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, artistshape_ := range diagram.ArtistShapes {
					res[artistshape_] = append(res[artistshape_], diagram)
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
		}
	// reverse maps of direct associations of Movement
	case Movement:
		switch fieldname {
		// insertion point for per direct association field
		case "Places":
			res := make(map[*Place][]*Movement)
			for movement := range stage.Movements {
				for _, place_ := range movement.Places {
					res[place_] = append(res[place_], movement)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MovementShape
	case MovementShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Place
	case Place:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *ArtefactType:
		res = any(new(ArtefactType)).(Type)
	case *ArtefactTypeShape:
		res = any(new(ArtefactTypeShape)).(Type)
	case *Artist:
		res = any(new(Artist)).(Type)
	case *ArtistShape:
		res = any(new(ArtistShape)).(Type)
	case *ControlPointShape:
		res = any(new(ControlPointShape)).(Type)
	case *Desk:
		res = any(new(Desk)).(Type)
	case *Diagram:
		res = any(new(Diagram)).(Type)
	case *Influence:
		res = any(new(Influence)).(Type)
	case *InfluenceShape:
		res = any(new(InfluenceShape)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *Movement:
		res = any(new(Movement)).(Type)
	case *MovementShape:
		res = any(new(MovementShape)).(Type)
	case *Place:
		res = any(new(Place)).(Type)
	}
	return res
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ArtefactType:
		res = "ArtefactType"
	case *ArtefactTypeShape:
		res = "ArtefactTypeShape"
	case *Artist:
		res = "Artist"
	case *ArtistShape:
		res = "ArtistShape"
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
	case *Library:
		res = "Library"
	case *Movement:
		res = "Movement"
	case *MovementShape:
		res = "MovementShape"
	case *Place:
		res = "Place"
	}
	return res
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	res = make([]GongReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *ArtefactType:
		var rf ReverseField
		_ = rf
	case *ArtefactTypeShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ArtefactTypeShapes"
		res = append(res, rf)
	case *Artist:
		var rf ReverseField
		_ = rf
	case *ArtistShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ArtistShapes"
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
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Movement:
		var rf ReverseField
		_ = rf
	case *MovementShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "MovementShapes"
		res = append(res, rf)
	case *Place:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Movement"
		rf.Fieldname = "Places"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (artefacttype *ArtefactType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ArtefactType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ArtefactType",
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
	}
	return
}

func (artist *Artist) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsDead",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DateOfDeath",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:                 "Place",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Place",
		},
	}
	return
}

func (artistshape *ArtistShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Artist",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Artist",
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
			Name:               "ImagePng_X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ImagePng_Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ImagePng_Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ImagePng_Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ImagePng_X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ImagePng_Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "ImagePng_RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:               "ImagePngBase64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (controlpointshape *ControlPointShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsStartShapeTheClosestShape",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (desk *Desk) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "MovementShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MovementShape",
		},
		{
			Name:                 "ArtefactTypeShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArtefactTypeShape",
		},
		{
			Name:                 "ArtistShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArtistShape",
		},
		{
			Name:                 "InfluenceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "InfluenceShape",
		},
		{
			Name:               "IsEditable",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMovementCategoryNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsArtefactTypeCategoryNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsArtistCategoryNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsInfluenceCategoryNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMovementCategoryHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsArtefactTypeCategoryHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsArtistCategoryHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsInfluenceCategoryHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "StartDate",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "EndDate",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "NbYearsForIntervals",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "XMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "YMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "NextVerticalDateXMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RedColorCode",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BackgroundGreyColorCode",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "GrayColorCode",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomBoxYOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "BottomBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "BottomBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "BottomBoxFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomBoxFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomBoxFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomBoxLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomBoxLetterColorCode",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MovementRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "MovementTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "MovementDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "MovementFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MajorMovementFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MinorMovementFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "AbstractMovementFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "AbstractMovementRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "AbstractMovementTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "AbstractDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:                 "MovementDateRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "MovementDateTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "MovementDateTextDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "MovementDateAndPlacesFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementDateAndPlacesFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementDateAndPlacesFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementDateAndPlacesLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MovementBelowArcY_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "MovementBelowArcY_OffsetPerPlace",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "MovementPlacesRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "MovementPlacesTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "MovementPlacesDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "ArtefactTypeFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtefactTypeFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtefactTypeFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtefactTypeLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ArtefactTypeRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "ArtefactDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "ArtefactTypeStrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "ArtistRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "ArtistTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "ArtistDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "ArtistFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MajorArtistFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MinorArtistFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ArtistDateRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "ArtistDateTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "ArtistDateDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "ArtistDateAndPlacesFontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistDateAndPlacesFontWeigth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistDateAndPlacesFontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ArtistDateAndPlacesLetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ArtistPlacesRectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:                 "ArtistPlacesTextAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextAnchorType",
		},
		{
			Name:                 "ArtistPlacesDominantBaselineType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DominantBaselineType",
		},
		{
			Name:               "InfluenceArrowSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "InfluenceArrowStartOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "InfluenceArrowEndOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "InfluenceCornerRadius",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "InfluenceDashedLinePattern",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (influence *Influence) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "SourceMovement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Movement",
		},
		{
			Name:                 "SourceArtefactType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ArtefactType",
		},
		{
			Name:                 "SourceArtist",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Artist",
		},
		{
			Name:                 "TargetMovement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Movement",
		},
		{
			Name:                 "TargetArtefactType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ArtefactType",
		},
		{
			Name:                 "TargetArtist",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Artist",
		},
		{
			Name:               "IsHypothtical",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (influenceshape *InfluenceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Influence",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Influence",
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
			Name:               "Description",
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
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (movement *Movement) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Date",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "HideDate",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Places",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Place",
		},
		{
			Name:               "HasTaxonomicFilter",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "TaxonomicFilter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsFeatured",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "FeaturePrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsMajor",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMinor",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AdditionnalName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (movementshape *MovementShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Movement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Movement",
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
	}
	return
}

func (place *Place) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
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
func (artefacttype *ArtefactType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = artefacttype.Name
	case "ComputedPrefix":
		res.valueString = artefacttype.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", artefacttype.IsExpanded)
		res.valueBool = artefacttype.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = artefacttypeshape.Name
	case "ArtefactType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if artefacttypeshape.ArtefactType != nil {
			res.valueString = artefacttypeshape.ArtefactType.Name
			res.ids = artefacttypeshape.ArtefactType.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", artefacttypeshape.X)
		res.valueFloat = artefacttypeshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", artefacttypeshape.Y)
		res.valueFloat = artefacttypeshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", artefacttypeshape.Width)
		res.valueFloat = artefacttypeshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", artefacttypeshape.Height)
		res.valueFloat = artefacttypeshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", artefacttypeshape.IsHidden)
		res.valueBool = artefacttypeshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (artist *Artist) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = artist.Name
	case "ComputedPrefix":
		res.valueString = artist.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", artist.IsExpanded)
		res.valueBool = artist.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsDead":
		res.valueString = fmt.Sprintf("%t", artist.IsDead)
		res.valueBool = artist.IsDead
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DateOfDeath":
		res.valueString = artist.DateOfDeath.String()
	case "Place":
		res.GongFieldValueType = GongFieldValueTypePointer
		if artist.Place != nil {
			res.valueString = artist.Place.Name
			res.ids = artist.Place.GongGetUUID(stage)
		}
	}
	return
}

func (artistshape *ArtistShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = artistshape.Name
	case "Artist":
		res.GongFieldValueType = GongFieldValueTypePointer
		if artistshape.Artist != nil {
			res.valueString = artistshape.Artist.Name
			res.ids = artistshape.Artist.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", artistshape.X)
		res.valueFloat = artistshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", artistshape.Y)
		res.valueFloat = artistshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", artistshape.Width)
		res.valueFloat = artistshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", artistshape.Height)
		res.valueFloat = artistshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", artistshape.IsHidden)
		res.valueBool = artistshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ImagePng_X":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_X)
		res.valueFloat = artistshape.ImagePng_X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_Y":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_Y)
		res.valueFloat = artistshape.ImagePng_Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_Width":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_Width)
		res.valueFloat = artistshape.ImagePng_Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_Height":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_Height)
		res.valueFloat = artistshape.ImagePng_Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_X_Offset":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_X_Offset)
		res.valueFloat = artistshape.ImagePng_X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_Y_Offset":
		res.valueString = fmt.Sprintf("%f", artistshape.ImagePng_Y_Offset)
		res.valueFloat = artistshape.ImagePng_Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ImagePng_RectAnchorType":
		enum := artistshape.ImagePng_RectAnchorType
		res.valueString = enum.ToCodeString()
	case "ImagePngBase64Content":
		res.valueString = artistshape.ImagePngBase64Content
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
			res.ids = desk.SelectedDiagram.GongGetUUID(stage)
		}
	}
	return
}

func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "ComputedPrefix":
		res.valueString = diagram.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsExpanded)
		res.valueBool = diagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagram.IsChecked)
		res.valueBool = diagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MovementShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.MovementShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ArtefactTypeShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ArtefactTypeShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ArtistShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ArtistShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "InfluenceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.InfluenceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEditable":
		res.valueString = fmt.Sprintf("%t", diagram.IsEditable)
		res.valueBool = diagram.IsEditable
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsNodeExpanded)
		res.valueBool = diagram.IsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMovementCategoryNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsMovementCategoryNodeExpanded)
		res.valueBool = diagram.IsMovementCategoryNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsArtefactTypeCategoryNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryNodeExpanded)
		res.valueBool = diagram.IsArtefactTypeCategoryNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsArtistCategoryNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsArtistCategoryNodeExpanded)
		res.valueBool = diagram.IsArtistCategoryNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInfluenceCategoryNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded)
		res.valueBool = diagram.IsInfluenceCategoryNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMovementCategoryHidden":
		res.valueString = fmt.Sprintf("%t", diagram.IsMovementCategoryHidden)
		res.valueBool = diagram.IsMovementCategoryHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsArtefactTypeCategoryHidden":
		res.valueString = fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryHidden)
		res.valueBool = diagram.IsArtefactTypeCategoryHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsArtistCategoryHidden":
		res.valueString = fmt.Sprintf("%t", diagram.IsArtistCategoryHidden)
		res.valueBool = diagram.IsArtistCategoryHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInfluenceCategoryHidden":
		res.valueString = fmt.Sprintf("%t", diagram.IsInfluenceCategoryHidden)
		res.valueBool = diagram.IsInfluenceCategoryHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "StartDate":
		res.valueString = diagram.StartDate.String()
	case "EndDate":
		res.valueString = diagram.EndDate.String()
	case "NbYearsForIntervals":
		res.valueString = fmt.Sprintf("%d", diagram.NbYearsForIntervals)
		res.valueInt = diagram.NbYearsForIntervals
		res.GongFieldValueType = GongFieldValueTypeInt
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
	case "NextVerticalDateXMargin":
		res.valueString = fmt.Sprintf("%f", diagram.NextVerticalDateXMargin)
		res.valueFloat = diagram.NextVerticalDateXMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RedColorCode":
		res.valueString = diagram.RedColorCode
	case "BackgroundGreyColorCode":
		res.valueString = diagram.BackgroundGreyColorCode
	case "GrayColorCode":
		res.valueString = diagram.GrayColorCode
	case "BottomBoxYOffset":
		res.valueString = fmt.Sprintf("%f", diagram.BottomBoxYOffset)
		res.valueFloat = diagram.BottomBoxYOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "BottomBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagram.BottomBoxWidth)
		res.valueFloat = diagram.BottomBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "BottomBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagram.BottomBoxHeigth)
		res.valueFloat = diagram.BottomBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "BottomBoxFontSize":
		res.valueString = diagram.BottomBoxFontSize
	case "BottomBoxFontWeigth":
		res.valueString = diagram.BottomBoxFontWeigth
	case "BottomBoxFontFamily":
		res.valueString = diagram.BottomBoxFontFamily
	case "BottomBoxLetterSpacing":
		res.valueString = diagram.BottomBoxLetterSpacing
	case "BottomBoxLetterColorCode":
		res.valueString = diagram.BottomBoxLetterColorCode
	case "MovementRectAnchorType":
		enum := diagram.MovementRectAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementTextAnchorType":
		enum := diagram.MovementTextAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementDominantBaselineType":
		enum := diagram.MovementDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "MovementFontSize":
		res.valueString = diagram.MovementFontSize
	case "MajorMovementFontSize":
		res.valueString = diagram.MajorMovementFontSize
	case "MinorMovementFontSize":
		res.valueString = diagram.MinorMovementFontSize
	case "MovementFontWeigth":
		res.valueString = diagram.MovementFontWeigth
	case "MovementFontFamily":
		res.valueString = diagram.MovementFontFamily
	case "MovementLetterSpacing":
		res.valueString = diagram.MovementLetterSpacing
	case "AbstractMovementFontSize":
		res.valueString = diagram.AbstractMovementFontSize
	case "AbstractMovementRectAnchorType":
		enum := diagram.AbstractMovementRectAnchorType
		res.valueString = enum.ToCodeString()
	case "AbstractMovementTextAnchorType":
		enum := diagram.AbstractMovementTextAnchorType
		res.valueString = enum.ToCodeString()
	case "AbstractDominantBaselineType":
		enum := diagram.AbstractDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "MovementDateRectAnchorType":
		enum := diagram.MovementDateRectAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementDateTextAnchorType":
		enum := diagram.MovementDateTextAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementDateTextDominantBaselineType":
		enum := diagram.MovementDateTextDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "MovementDateAndPlacesFontSize":
		res.valueString = diagram.MovementDateAndPlacesFontSize
	case "MovementDateAndPlacesFontWeigth":
		res.valueString = diagram.MovementDateAndPlacesFontWeigth
	case "MovementDateAndPlacesFontFamily":
		res.valueString = diagram.MovementDateAndPlacesFontFamily
	case "MovementDateAndPlacesLetterSpacing":
		res.valueString = diagram.MovementDateAndPlacesLetterSpacing
	case "MovementBelowArcY_Offset":
		res.valueString = fmt.Sprintf("%f", diagram.MovementBelowArcY_Offset)
		res.valueFloat = diagram.MovementBelowArcY_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MovementBelowArcY_OffsetPerPlace":
		res.valueString = fmt.Sprintf("%f", diagram.MovementBelowArcY_OffsetPerPlace)
		res.valueFloat = diagram.MovementBelowArcY_OffsetPerPlace
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MovementPlacesRectAnchorType":
		enum := diagram.MovementPlacesRectAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementPlacesTextAnchorType":
		enum := diagram.MovementPlacesTextAnchorType
		res.valueString = enum.ToCodeString()
	case "MovementPlacesDominantBaselineType":
		enum := diagram.MovementPlacesDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "ArtefactTypeFontSize":
		res.valueString = diagram.ArtefactTypeFontSize
	case "ArtefactTypeFontWeigth":
		res.valueString = diagram.ArtefactTypeFontWeigth
	case "ArtefactTypeFontFamily":
		res.valueString = diagram.ArtefactTypeFontFamily
	case "ArtefactTypeLetterSpacing":
		res.valueString = diagram.ArtefactTypeLetterSpacing
	case "ArtefactTypeRectAnchorType":
		enum := diagram.ArtefactTypeRectAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtefactDominantBaselineType":
		enum := diagram.ArtefactDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "ArtefactTypeStrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.ArtefactTypeStrokeWidth)
		res.valueFloat = diagram.ArtefactTypeStrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArtistRectAnchorType":
		enum := diagram.ArtistRectAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistTextAnchorType":
		enum := diagram.ArtistTextAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistDominantBaselineType":
		enum := diagram.ArtistDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "ArtistFontSize":
		res.valueString = diagram.ArtistFontSize
	case "MajorArtistFontSize":
		res.valueString = diagram.MajorArtistFontSize
	case "MinorArtistFontSize":
		res.valueString = diagram.MinorArtistFontSize
	case "ArtistFontWeigth":
		res.valueString = diagram.ArtistFontWeigth
	case "ArtistFontFamily":
		res.valueString = diagram.ArtistFontFamily
	case "ArtistLetterSpacing":
		res.valueString = diagram.ArtistLetterSpacing
	case "ArtistDateRectAnchorType":
		enum := diagram.ArtistDateRectAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistDateTextAnchorType":
		enum := diagram.ArtistDateTextAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistDateDominantBaselineType":
		enum := diagram.ArtistDateDominantBaselineType
		res.valueString = enum.ToCodeString()
	case "ArtistDateAndPlacesFontSize":
		res.valueString = diagram.ArtistDateAndPlacesFontSize
	case "ArtistDateAndPlacesFontWeigth":
		res.valueString = diagram.ArtistDateAndPlacesFontWeigth
	case "ArtistDateAndPlacesFontFamily":
		res.valueString = diagram.ArtistDateAndPlacesFontFamily
	case "ArtistDateAndPlacesLetterSpacing":
		res.valueString = diagram.ArtistDateAndPlacesLetterSpacing
	case "ArtistPlacesRectAnchorType":
		enum := diagram.ArtistPlacesRectAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistPlacesTextAnchorType":
		enum := diagram.ArtistPlacesTextAnchorType
		res.valueString = enum.ToCodeString()
	case "ArtistPlacesDominantBaselineType":
		enum := diagram.ArtistPlacesDominantBaselineType
		res.valueString = enum.ToCodeString()
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
	case "ComputedPrefix":
		res.valueString = influence.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", influence.IsExpanded)
		res.valueBool = influence.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SourceMovement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceMovement != nil {
			res.valueString = influence.SourceMovement.Name
			res.ids = influence.SourceMovement.GongGetUUID(stage)
		}
	case "SourceArtefactType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceArtefactType != nil {
			res.valueString = influence.SourceArtefactType.Name
			res.ids = influence.SourceArtefactType.GongGetUUID(stage)
		}
	case "SourceArtist":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.SourceArtist != nil {
			res.valueString = influence.SourceArtist.Name
			res.ids = influence.SourceArtist.GongGetUUID(stage)
		}
	case "TargetMovement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetMovement != nil {
			res.valueString = influence.TargetMovement.Name
			res.ids = influence.TargetMovement.GongGetUUID(stage)
		}
	case "TargetArtefactType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetArtefactType != nil {
			res.valueString = influence.TargetArtefactType.Name
			res.ids = influence.TargetArtefactType.GongGetUUID(stage)
		}
	case "TargetArtist":
		res.GongFieldValueType = GongFieldValueTypePointer
		if influence.TargetArtist != nil {
			res.valueString = influence.TargetArtist.Name
			res.ids = influence.TargetArtist.GongGetUUID(stage)
		}
	case "IsHypothtical":
		res.valueString = fmt.Sprintf("%t", influence.IsHypothtical)
		res.valueBool = influence.IsHypothtical
		res.GongFieldValueType = GongFieldValueTypeBool
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
			res.ids = influenceshape.Influence.GongGetUUID(stage)
		}
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", influenceshape.IsHidden)
		res.valueBool = influenceshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range influenceshape.ControlPointShapes {
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
	case "Description":
		res.valueString = library.Description
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (movement *Movement) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = movement.Name
	case "ComputedPrefix":
		res.valueString = movement.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", movement.IsExpanded)
		res.valueBool = movement.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Date":
		res.valueString = movement.Date.String()
	case "HideDate":
		res.valueString = fmt.Sprintf("%t", movement.HideDate)
		res.valueBool = movement.HideDate
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Places":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range movement.Places {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HasTaxonomicFilter":
		res.valueString = fmt.Sprintf("%t", movement.HasTaxonomicFilter)
		res.valueBool = movement.HasTaxonomicFilter
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TaxonomicFilter":
		res.valueString = movement.TaxonomicFilter
	case "IsFeatured":
		res.valueString = fmt.Sprintf("%t", movement.IsFeatured)
		res.valueBool = movement.IsFeatured
		res.GongFieldValueType = GongFieldValueTypeBool
	case "FeaturePrefix":
		res.valueString = movement.FeaturePrefix
	case "IsMajor":
		res.valueString = fmt.Sprintf("%t", movement.IsMajor)
		res.valueBool = movement.IsMajor
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMinor":
		res.valueString = fmt.Sprintf("%t", movement.IsMinor)
		res.valueBool = movement.IsMinor
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AdditionnalName":
		res.valueString = movement.AdditionnalName
	}
	return
}

func (movementshape *MovementShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = movementshape.Name
	case "Movement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if movementshape.Movement != nil {
			res.valueString = movementshape.Movement.Name
			res.ids = movementshape.Movement.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", movementshape.X)
		res.valueFloat = movementshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", movementshape.Y)
		res.valueFloat = movementshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", movementshape.Width)
		res.valueFloat = movementshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", movementshape.Height)
		res.valueFloat = movementshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", movementshape.IsHidden)
		res.valueBool = movementshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (place *Place) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = place.Name
	}
	return
}

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (artefacttype *ArtefactType) GongGetGongstructName() string {
	return "ArtefactType"
}

func (artefacttypeshape *ArtefactTypeShape) GongGetGongstructName() string {
	return "ArtefactTypeShape"
}

func (artist *Artist) GongGetGongstructName() string {
	return "Artist"
}

func (artistshape *ArtistShape) GongGetGongstructName() string {
	return "ArtistShape"
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

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (movement *Movement) GongGetGongstructName() string {
	return "Movement"
}

func (movementshape *MovementShape) GongGetGongstructName() string {
	return "MovementShape"
}

func (place *Place) GongGetGongstructName() string {
	return "Place"
}

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.ArtefactTypes, &stage.ArtefactTypes_mapString)

	__gong__rebuildMapString(stage.ArtefactTypeShapes, &stage.ArtefactTypeShapes_mapString)

	__gong__rebuildMapString(stage.Artists, &stage.Artists_mapString)

	__gong__rebuildMapString(stage.ArtistShapes, &stage.ArtistShapes_mapString)

	__gong__rebuildMapString(stage.ControlPointShapes, &stage.ControlPointShapes_mapString)

	__gong__rebuildMapString(stage.Desks, &stage.Desks_mapString)

	__gong__rebuildMapString(stage.Diagrams, &stage.Diagrams_mapString)

	__gong__rebuildMapString(stage.Influences, &stage.Influences_mapString)

	__gong__rebuildMapString(stage.InfluenceShapes, &stage.InfluenceShapes_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.Movements, &stage.Movements_mapString)

	__gong__rebuildMapString(stage.MovementShapes, &stage.MovementShapes_mapString)

	__gong__rebuildMapString(stage.Places, &stage.Places_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
