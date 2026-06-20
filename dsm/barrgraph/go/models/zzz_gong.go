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

	barrgraph_go "github.com/fullstack-lang/gong/dsm/barrgraph/go"
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
	ArtefactTypes                map[*ArtefactType]struct{}
	ArtefactTypes_instance       map[*ArtefactType]*ArtefactType
	ArtefactTypes_mapString      map[string]*ArtefactType
	ArtefactTypeOrder            uint
	ArtefactType_stagedOrder     map[*ArtefactType]uint
	ArtefactType_orderStaged     map[uint]*ArtefactType
	ArtefactTypes_reference      map[*ArtefactType]*ArtefactType
	ArtefactTypes_referenceOrder map[*ArtefactType]uint

	// insertion point for slice of pointers maps
	OnAfterArtefactTypeCreateCallback OnAfterCreateInterface[ArtefactType]
	OnAfterArtefactTypeUpdateCallback OnAfterUpdateInterface[ArtefactType]
	OnAfterArtefactTypeDeleteCallback OnAfterDeleteInterface[ArtefactType]
	OnAfterArtefactTypeReadCallback   OnAfterReadInterface[ArtefactType]

	ArtefactTypeShapes                map[*ArtefactTypeShape]struct{}
	ArtefactTypeShapes_instance       map[*ArtefactTypeShape]*ArtefactTypeShape
	ArtefactTypeShapes_mapString      map[string]*ArtefactTypeShape
	ArtefactTypeShapeOrder            uint
	ArtefactTypeShape_stagedOrder     map[*ArtefactTypeShape]uint
	ArtefactTypeShape_orderStaged     map[uint]*ArtefactTypeShape
	ArtefactTypeShapes_reference      map[*ArtefactTypeShape]*ArtefactTypeShape
	ArtefactTypeShapes_referenceOrder map[*ArtefactTypeShape]uint

	// insertion point for slice of pointers maps
	OnAfterArtefactTypeShapeCreateCallback OnAfterCreateInterface[ArtefactTypeShape]
	OnAfterArtefactTypeShapeUpdateCallback OnAfterUpdateInterface[ArtefactTypeShape]
	OnAfterArtefactTypeShapeDeleteCallback OnAfterDeleteInterface[ArtefactTypeShape]
	OnAfterArtefactTypeShapeReadCallback   OnAfterReadInterface[ArtefactTypeShape]

	Artists                map[*Artist]struct{}
	Artists_instance       map[*Artist]*Artist
	Artists_mapString      map[string]*Artist
	ArtistOrder            uint
	Artist_stagedOrder     map[*Artist]uint
	Artist_orderStaged     map[uint]*Artist
	Artists_reference      map[*Artist]*Artist
	Artists_referenceOrder map[*Artist]uint

	// insertion point for slice of pointers maps
	OnAfterArtistCreateCallback OnAfterCreateInterface[Artist]
	OnAfterArtistUpdateCallback OnAfterUpdateInterface[Artist]
	OnAfterArtistDeleteCallback OnAfterDeleteInterface[Artist]
	OnAfterArtistReadCallback   OnAfterReadInterface[Artist]

	ArtistShapes                map[*ArtistShape]struct{}
	ArtistShapes_instance       map[*ArtistShape]*ArtistShape
	ArtistShapes_mapString      map[string]*ArtistShape
	ArtistShapeOrder            uint
	ArtistShape_stagedOrder     map[*ArtistShape]uint
	ArtistShape_orderStaged     map[uint]*ArtistShape
	ArtistShapes_reference      map[*ArtistShape]*ArtistShape
	ArtistShapes_referenceOrder map[*ArtistShape]uint

	// insertion point for slice of pointers maps
	OnAfterArtistShapeCreateCallback OnAfterCreateInterface[ArtistShape]
	OnAfterArtistShapeUpdateCallback OnAfterUpdateInterface[ArtistShape]
	OnAfterArtistShapeDeleteCallback OnAfterDeleteInterface[ArtistShape]
	OnAfterArtistShapeReadCallback   OnAfterReadInterface[ArtistShape]

	ControlPointShapes                map[*ControlPointShape]struct{}
	ControlPointShapes_instance       map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_mapString      map[string]*ControlPointShape
	ControlPointShapeOrder            uint
	ControlPointShape_stagedOrder     map[*ControlPointShape]uint
	ControlPointShape_orderStaged     map[uint]*ControlPointShape
	ControlPointShapes_reference      map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_referenceOrder map[*ControlPointShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback OnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback OnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback OnAfterDeleteInterface[ControlPointShape]
	OnAfterControlPointShapeReadCallback   OnAfterReadInterface[ControlPointShape]

	Desks                map[*Desk]struct{}
	Desks_instance       map[*Desk]*Desk
	Desks_mapString      map[string]*Desk
	DeskOrder            uint
	Desk_stagedOrder     map[*Desk]uint
	Desk_orderStaged     map[uint]*Desk
	Desks_reference      map[*Desk]*Desk
	Desks_referenceOrder map[*Desk]uint

	// insertion point for slice of pointers maps
	OnAfterDeskCreateCallback OnAfterCreateInterface[Desk]
	OnAfterDeskUpdateCallback OnAfterUpdateInterface[Desk]
	OnAfterDeskDeleteCallback OnAfterDeleteInterface[Desk]
	OnAfterDeskReadCallback   OnAfterReadInterface[Desk]

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

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Influences                map[*Influence]struct{}
	Influences_instance       map[*Influence]*Influence
	Influences_mapString      map[string]*Influence
	InfluenceOrder            uint
	Influence_stagedOrder     map[*Influence]uint
	Influence_orderStaged     map[uint]*Influence
	Influences_reference      map[*Influence]*Influence
	Influences_referenceOrder map[*Influence]uint

	// insertion point for slice of pointers maps
	OnAfterInfluenceCreateCallback OnAfterCreateInterface[Influence]
	OnAfterInfluenceUpdateCallback OnAfterUpdateInterface[Influence]
	OnAfterInfluenceDeleteCallback OnAfterDeleteInterface[Influence]
	OnAfterInfluenceReadCallback   OnAfterReadInterface[Influence]

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

	OnAfterInfluenceShapeCreateCallback OnAfterCreateInterface[InfluenceShape]
	OnAfterInfluenceShapeUpdateCallback OnAfterUpdateInterface[InfluenceShape]
	OnAfterInfluenceShapeDeleteCallback OnAfterDeleteInterface[InfluenceShape]
	OnAfterInfluenceShapeReadCallback   OnAfterReadInterface[InfluenceShape]

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

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

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

	OnAfterMovementCreateCallback OnAfterCreateInterface[Movement]
	OnAfterMovementUpdateCallback OnAfterUpdateInterface[Movement]
	OnAfterMovementDeleteCallback OnAfterDeleteInterface[Movement]
	OnAfterMovementReadCallback   OnAfterReadInterface[Movement]

	MovementShapes                map[*MovementShape]struct{}
	MovementShapes_instance       map[*MovementShape]*MovementShape
	MovementShapes_mapString      map[string]*MovementShape
	MovementShapeOrder            uint
	MovementShape_stagedOrder     map[*MovementShape]uint
	MovementShape_orderStaged     map[uint]*MovementShape
	MovementShapes_reference      map[*MovementShape]*MovementShape
	MovementShapes_referenceOrder map[*MovementShape]uint

	// insertion point for slice of pointers maps
	OnAfterMovementShapeCreateCallback OnAfterCreateInterface[MovementShape]
	OnAfterMovementShapeUpdateCallback OnAfterUpdateInterface[MovementShape]
	OnAfterMovementShapeDeleteCallback OnAfterDeleteInterface[MovementShape]
	OnAfterMovementShapeReadCallback   OnAfterReadInterface[MovementShape]

	Places                map[*Place]struct{}
	Places_instance       map[*Place]*Place
	Places_mapString      map[string]*Place
	PlaceOrder            uint
	Place_stagedOrder     map[*Place]uint
	Place_orderStaged     map[uint]*Place
	Places_reference      map[*Place]*Place
	Places_referenceOrder map[*Place]uint

	// insertion point for slice of pointers maps
	OnAfterPlaceCreateCallback OnAfterCreateInterface[Place]
	OnAfterPlaceUpdateCallback OnAfterUpdateInterface[Place]
	OnAfterPlaceDeleteCallback OnAfterDeleteInterface[Place]
	OnAfterPlaceReadCallback   OnAfterReadInterface[Place]

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
	stage.ArtefactTypes_reference = make(map[*ArtefactType]*ArtefactType)
	stage.ArtefactTypes_instance = make(map[*ArtefactType]*ArtefactType)
	stage.ArtefactTypes_referenceOrder = make(map[*ArtefactType]uint)

	stage.ArtefactTypeShapes_reference = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	stage.ArtefactTypeShapes_instance = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	stage.ArtefactTypeShapes_referenceOrder = make(map[*ArtefactTypeShape]uint)

	stage.Artists_reference = make(map[*Artist]*Artist)
	stage.Artists_instance = make(map[*Artist]*Artist)
	stage.Artists_referenceOrder = make(map[*Artist]uint)

	stage.ArtistShapes_reference = make(map[*ArtistShape]*ArtistShape)
	stage.ArtistShapes_instance = make(map[*ArtistShape]*ArtistShape)
	stage.ArtistShapes_referenceOrder = make(map[*ArtistShape]uint)

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint)

	stage.Desks_reference = make(map[*Desk]*Desk)
	stage.Desks_instance = make(map[*Desk]*Desk)
	stage.Desks_referenceOrder = make(map[*Desk]uint)

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint)

	stage.Influences_reference = make(map[*Influence]*Influence)
	stage.Influences_instance = make(map[*Influence]*Influence)
	stage.Influences_referenceOrder = make(map[*Influence]uint)

	stage.InfluenceShapes_reference = make(map[*InfluenceShape]*InfluenceShape)
	stage.InfluenceShapes_instance = make(map[*InfluenceShape]*InfluenceShape)
	stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Movements_reference = make(map[*Movement]*Movement)
	stage.Movements_instance = make(map[*Movement]*Movement)
	stage.Movements_referenceOrder = make(map[*Movement]uint)

	stage.MovementShapes_reference = make(map[*MovementShape]*MovementShape)
	stage.MovementShapes_instance = make(map[*MovementShape]*MovementShape)
	stage.MovementShapes_referenceOrder = make(map[*MovementShape]uint)

	stage.Places_reference = make(map[*Place]*Place)
	stage.Places_instance = make(map[*Place]*Place)
	stage.Places_referenceOrder = make(map[*Place]uint)

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
	var maxArtefactTypeOrder uint
	var foundArtefactType bool
	for _, order := range stage.ArtefactType_stagedOrder {
		if !foundArtefactType || order > maxArtefactTypeOrder {
			maxArtefactTypeOrder = order
			foundArtefactType = true
		}
	}
	if foundArtefactType {
		stage.ArtefactTypeOrder = maxArtefactTypeOrder + 1
	} else {
		stage.ArtefactTypeOrder = 0
	}

	var maxArtefactTypeShapeOrder uint
	var foundArtefactTypeShape bool
	for _, order := range stage.ArtefactTypeShape_stagedOrder {
		if !foundArtefactTypeShape || order > maxArtefactTypeShapeOrder {
			maxArtefactTypeShapeOrder = order
			foundArtefactTypeShape = true
		}
	}
	if foundArtefactTypeShape {
		stage.ArtefactTypeShapeOrder = maxArtefactTypeShapeOrder + 1
	} else {
		stage.ArtefactTypeShapeOrder = 0
	}

	var maxArtistOrder uint
	var foundArtist bool
	for _, order := range stage.Artist_stagedOrder {
		if !foundArtist || order > maxArtistOrder {
			maxArtistOrder = order
			foundArtist = true
		}
	}
	if foundArtist {
		stage.ArtistOrder = maxArtistOrder + 1
	} else {
		stage.ArtistOrder = 0
	}

	var maxArtistShapeOrder uint
	var foundArtistShape bool
	for _, order := range stage.ArtistShape_stagedOrder {
		if !foundArtistShape || order > maxArtistShapeOrder {
			maxArtistShapeOrder = order
			foundArtistShape = true
		}
	}
	if foundArtistShape {
		stage.ArtistShapeOrder = maxArtistShapeOrder + 1
	} else {
		stage.ArtistShapeOrder = 0
	}

	var maxControlPointShapeOrder uint
	var foundControlPointShape bool
	for _, order := range stage.ControlPointShape_stagedOrder {
		if !foundControlPointShape || order > maxControlPointShapeOrder {
			maxControlPointShapeOrder = order
			foundControlPointShape = true
		}
	}
	if foundControlPointShape {
		stage.ControlPointShapeOrder = maxControlPointShapeOrder + 1
	} else {
		stage.ControlPointShapeOrder = 0
	}

	var maxDeskOrder uint
	var foundDesk bool
	for _, order := range stage.Desk_stagedOrder {
		if !foundDesk || order > maxDeskOrder {
			maxDeskOrder = order
			foundDesk = true
		}
	}
	if foundDesk {
		stage.DeskOrder = maxDeskOrder + 1
	} else {
		stage.DeskOrder = 0
	}

	var maxDiagramOrder uint
	var foundDiagram bool
	for _, order := range stage.Diagram_stagedOrder {
		if !foundDiagram || order > maxDiagramOrder {
			maxDiagramOrder = order
			foundDiagram = true
		}
	}
	if foundDiagram {
		stage.DiagramOrder = maxDiagramOrder + 1
	} else {
		stage.DiagramOrder = 0
	}

	var maxInfluenceOrder uint
	var foundInfluence bool
	for _, order := range stage.Influence_stagedOrder {
		if !foundInfluence || order > maxInfluenceOrder {
			maxInfluenceOrder = order
			foundInfluence = true
		}
	}
	if foundInfluence {
		stage.InfluenceOrder = maxInfluenceOrder + 1
	} else {
		stage.InfluenceOrder = 0
	}

	var maxInfluenceShapeOrder uint
	var foundInfluenceShape bool
	for _, order := range stage.InfluenceShape_stagedOrder {
		if !foundInfluenceShape || order > maxInfluenceShapeOrder {
			maxInfluenceShapeOrder = order
			foundInfluenceShape = true
		}
	}
	if foundInfluenceShape {
		stage.InfluenceShapeOrder = maxInfluenceShapeOrder + 1
	} else {
		stage.InfluenceShapeOrder = 0
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

	var maxMovementOrder uint
	var foundMovement bool
	for _, order := range stage.Movement_stagedOrder {
		if !foundMovement || order > maxMovementOrder {
			maxMovementOrder = order
			foundMovement = true
		}
	}
	if foundMovement {
		stage.MovementOrder = maxMovementOrder + 1
	} else {
		stage.MovementOrder = 0
	}

	var maxMovementShapeOrder uint
	var foundMovementShape bool
	for _, order := range stage.MovementShape_stagedOrder {
		if !foundMovementShape || order > maxMovementShapeOrder {
			maxMovementShapeOrder = order
			foundMovementShape = true
		}
	}
	if foundMovementShape {
		stage.MovementShapeOrder = maxMovementShapeOrder + 1
	} else {
		stage.MovementShapeOrder = 0
	}

	var maxPlaceOrder uint
	var foundPlace bool
	for _, order := range stage.Place_stagedOrder {
		if !foundPlace || order > maxPlaceOrder {
			maxPlaceOrder = order
			foundPlace = true
		}
	}
	if foundPlace {
		stage.PlaceOrder = maxPlaceOrder + 1
	} else {
		stage.PlaceOrder = 0
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
	case *ArtefactType:
		tmp := GetStructInstancesByOrder(stage.ArtefactTypes, stage.ArtefactType_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ArtefactType implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ArtefactTypeShape:
		tmp := GetStructInstancesByOrder(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ArtefactTypeShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Artist:
		tmp := GetStructInstancesByOrder(stage.Artists, stage.Artist_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Artist implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ArtistShape:
		tmp := GetStructInstancesByOrder(stage.ArtistShapes, stage.ArtistShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ArtistShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlPointShape:
		tmp := GetStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder)

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
		tmp := GetStructInstancesByOrder(stage.Desks, stage.Desk_stagedOrder)

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
		tmp := GetStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder)

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
		tmp := GetStructInstancesByOrder(stage.Influences, stage.Influence_stagedOrder)

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
		tmp := GetStructInstancesByOrder(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder)

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
	case *Movement:
		tmp := GetStructInstancesByOrder(stage.Movements, stage.Movement_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Movement implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MovementShape:
		tmp := GetStructInstancesByOrder(stage.MovementShapes, stage.MovementShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MovementShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Place:
		tmp := GetStructInstancesByOrder(stage.Places, stage.Place_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Place implements.
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
	case "ArtefactType":
		res = GetNamedStructInstances(stage.ArtefactTypes, stage.ArtefactType_stagedOrder)
	case "ArtefactTypeShape":
		res = GetNamedStructInstances(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder)
	case "Artist":
		res = GetNamedStructInstances(stage.Artists, stage.Artist_stagedOrder)
	case "ArtistShape":
		res = GetNamedStructInstances(stage.ArtistShapes, stage.ArtistShape_stagedOrder)
	case "ControlPointShape":
		res = GetNamedStructInstances(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder)
	case "Desk":
		res = GetNamedStructInstances(stage.Desks, stage.Desk_stagedOrder)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.Diagram_stagedOrder)
	case "Influence":
		res = GetNamedStructInstances(stage.Influences, stage.Influence_stagedOrder)
	case "InfluenceShape":
		res = GetNamedStructInstances(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Movement":
		res = GetNamedStructInstances(stage.Movements, stage.Movement_stagedOrder)
	case "MovementShape":
		res = GetNamedStructInstances(stage.MovementShapes, stage.MovementShape_stagedOrder)
	case "Place":
		res = GetNamedStructInstances(stage.Places, stage.Place_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return barrgraph_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return barrgraph_go.GoDiagramsDir
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
	CommitArtefactType(artefacttype *ArtefactType)
	CheckoutArtefactType(artefacttype *ArtefactType)
	CommitArtefactTypeShape(artefacttypeshape *ArtefactTypeShape)
	CheckoutArtefactTypeShape(artefacttypeshape *ArtefactTypeShape)
	CommitArtist(artist *Artist)
	CheckoutArtist(artist *Artist)
	CommitArtistShape(artistshape *ArtistShape)
	CheckoutArtistShape(artistshape *ArtistShape)
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
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitMovement(movement *Movement)
	CheckoutMovement(movement *Movement)
	CommitMovementShape(movementshape *MovementShape)
	CheckoutMovementShape(movementshape *MovementShape)
	CommitPlace(place *Place)
	CheckoutPlace(place *Place)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ArtefactType"},
			{name: "ArtefactTypeShape"},
			{name: "Artist"},
			{name: "ArtistShape"},
			{name: "ControlPointShape"},
			{name: "Desk"},
			{name: "Diagram"},
			{name: "Influence"},
			{name: "InfluenceShape"},
			{name: "Library"},
			{name: "Movement"},
			{name: "MovementShape"},
			{name: "Place"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ArtefactType:
		return stage.ArtefactType_stagedOrder[instance]
	case *ArtefactTypeShape:
		return stage.ArtefactTypeShape_stagedOrder[instance]
	case *Artist:
		return stage.Artist_stagedOrder[instance]
	case *ArtistShape:
		return stage.ArtistShape_stagedOrder[instance]
	case *ControlPointShape:
		return stage.ControlPointShape_stagedOrder[instance]
	case *Desk:
		return stage.Desk_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Influence:
		return stage.Influence_stagedOrder[instance]
	case *InfluenceShape:
		return stage.InfluenceShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Movement:
		return stage.Movement_stagedOrder[instance]
	case *MovementShape:
		return stage.MovementShape_stagedOrder[instance]
	case *Place:
		return stage.Place_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ArtefactType:
		return stage.ArtefactType_stagedOrder[instance]
	case *ArtefactTypeShape:
		return stage.ArtefactTypeShape_stagedOrder[instance]
	case *Artist:
		return stage.Artist_stagedOrder[instance]
	case *ArtistShape:
		return stage.ArtistShape_stagedOrder[instance]
	case *ControlPointShape:
		return stage.ControlPointShape_stagedOrder[instance]
	case *Desk:
		return stage.Desk_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Influence:
		return stage.Influence_stagedOrder[instance]
	case *InfluenceShape:
		return stage.InfluenceShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Movement:
		return stage.Movement_stagedOrder[instance]
	case *MovementShape:
		return stage.MovementShape_stagedOrder[instance]
	case *Place:
		return stage.Place_stagedOrder[instance]
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
	if _, ok := stage.ArtefactTypes[artefacttype]; !ok {
		stage.ArtefactTypes[artefacttype] = struct{}{}
		stage.ArtefactType_stagedOrder[artefacttype] = stage.ArtefactTypeOrder
		stage.ArtefactType_orderStaged[stage.ArtefactTypeOrder] = artefacttype
		stage.ArtefactTypeOrder++
	}
	stage.ArtefactTypes_mapString[artefacttype.Name] = artefacttype

	return artefacttype
}

// StagePreserveOrder puts artefacttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtefactTypeOrder
// - update stage.ArtefactTypeOrder accordingly
func (artefacttype *ArtefactType) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ArtefactTypes[artefacttype]; !ok {
		stage.ArtefactTypes[artefacttype] = struct{}{}

		if order > stage.ArtefactTypeOrder {
			stage.ArtefactTypeOrder = order
		}
		stage.ArtefactType_stagedOrder[artefacttype] = order
		stage.ArtefactType_orderStaged[order] = artefacttype
		stage.ArtefactTypeOrder++
	}
	stage.ArtefactTypes_mapString[artefacttype.Name] = artefacttype
}

// Unstage removes artefacttype off the model stage
func (artefacttype *ArtefactType) Unstage(stage *Stage) *ArtefactType {
	delete(stage.ArtefactTypes, artefacttype)
	// issue1150
	// delete(stage.ArtefactType_stagedOrder, artefacttype)
	delete(stage.ArtefactTypes_mapString, artefacttype.Name)

	return artefacttype
}

// UnstageVoid removes artefacttype off the model stage
func (artefacttype *ArtefactType) UnstageVoid(stage *Stage) {
	delete(stage.ArtefactTypes, artefacttype)
	// issue1150
	// delete(stage.ArtefactType_stagedOrder, artefacttype)
	delete(stage.ArtefactTypes_mapString, artefacttype.Name)
}

// commit artefacttype to the back repo (if it is already staged)
func (artefacttype *ArtefactType) Commit(stage *Stage) *ArtefactType {
	if _, ok := stage.ArtefactTypes[artefacttype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArtefactType(artefacttype)
		}
	}
	return artefacttype
}

func (artefacttype *ArtefactType) CommitVoid(stage *Stage) {
	artefacttype.Commit(stage)
}

func (artefacttype *ArtefactType) StageVoid(stage *Stage) {
	artefacttype.Stage(stage)
}

// Checkout artefacttype to the back repo (if it is already staged)
func (artefacttype *ArtefactType) Checkout(stage *Stage) *ArtefactType {
	if _, ok := stage.ArtefactTypes[artefacttype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArtefactType(artefacttype)
		}
	}
	return artefacttype
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
	if _, ok := stage.ArtefactTypeShapes[artefacttypeshape]; !ok {
		stage.ArtefactTypeShapes[artefacttypeshape] = struct{}{}
		stage.ArtefactTypeShape_stagedOrder[artefacttypeshape] = stage.ArtefactTypeShapeOrder
		stage.ArtefactTypeShape_orderStaged[stage.ArtefactTypeShapeOrder] = artefacttypeshape
		stage.ArtefactTypeShapeOrder++
	}
	stage.ArtefactTypeShapes_mapString[artefacttypeshape.Name] = artefacttypeshape

	return artefacttypeshape
}

// StagePreserveOrder puts artefacttypeshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtefactTypeShapeOrder
// - update stage.ArtefactTypeShapeOrder accordingly
func (artefacttypeshape *ArtefactTypeShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ArtefactTypeShapes[artefacttypeshape]; !ok {
		stage.ArtefactTypeShapes[artefacttypeshape] = struct{}{}

		if order > stage.ArtefactTypeShapeOrder {
			stage.ArtefactTypeShapeOrder = order
		}
		stage.ArtefactTypeShape_stagedOrder[artefacttypeshape] = order
		stage.ArtefactTypeShape_orderStaged[order] = artefacttypeshape
		stage.ArtefactTypeShapeOrder++
	}
	stage.ArtefactTypeShapes_mapString[artefacttypeshape.Name] = artefacttypeshape
}

// Unstage removes artefacttypeshape off the model stage
func (artefacttypeshape *ArtefactTypeShape) Unstage(stage *Stage) *ArtefactTypeShape {
	delete(stage.ArtefactTypeShapes, artefacttypeshape)
	// issue1150
	// delete(stage.ArtefactTypeShape_stagedOrder, artefacttypeshape)
	delete(stage.ArtefactTypeShapes_mapString, artefacttypeshape.Name)

	return artefacttypeshape
}

// UnstageVoid removes artefacttypeshape off the model stage
func (artefacttypeshape *ArtefactTypeShape) UnstageVoid(stage *Stage) {
	delete(stage.ArtefactTypeShapes, artefacttypeshape)
	// issue1150
	// delete(stage.ArtefactTypeShape_stagedOrder, artefacttypeshape)
	delete(stage.ArtefactTypeShapes_mapString, artefacttypeshape.Name)
}

// commit artefacttypeshape to the back repo (if it is already staged)
func (artefacttypeshape *ArtefactTypeShape) Commit(stage *Stage) *ArtefactTypeShape {
	if _, ok := stage.ArtefactTypeShapes[artefacttypeshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArtefactTypeShape(artefacttypeshape)
		}
	}
	return artefacttypeshape
}

func (artefacttypeshape *ArtefactTypeShape) CommitVoid(stage *Stage) {
	artefacttypeshape.Commit(stage)
}

func (artefacttypeshape *ArtefactTypeShape) StageVoid(stage *Stage) {
	artefacttypeshape.Stage(stage)
}

// Checkout artefacttypeshape to the back repo (if it is already staged)
func (artefacttypeshape *ArtefactTypeShape) Checkout(stage *Stage) *ArtefactTypeShape {
	if _, ok := stage.ArtefactTypeShapes[artefacttypeshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArtefactTypeShape(artefacttypeshape)
		}
	}
	return artefacttypeshape
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
	if _, ok := stage.Artists[artist]; !ok {
		stage.Artists[artist] = struct{}{}
		stage.Artist_stagedOrder[artist] = stage.ArtistOrder
		stage.Artist_orderStaged[stage.ArtistOrder] = artist
		stage.ArtistOrder++
	}
	stage.Artists_mapString[artist.Name] = artist

	return artist
}

// StagePreserveOrder puts artist to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtistOrder
// - update stage.ArtistOrder accordingly
func (artist *Artist) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Artists[artist]; !ok {
		stage.Artists[artist] = struct{}{}

		if order > stage.ArtistOrder {
			stage.ArtistOrder = order
		}
		stage.Artist_stagedOrder[artist] = order
		stage.Artist_orderStaged[order] = artist
		stage.ArtistOrder++
	}
	stage.Artists_mapString[artist.Name] = artist
}

// Unstage removes artist off the model stage
func (artist *Artist) Unstage(stage *Stage) *Artist {
	delete(stage.Artists, artist)
	// issue1150
	// delete(stage.Artist_stagedOrder, artist)
	delete(stage.Artists_mapString, artist.Name)

	return artist
}

// UnstageVoid removes artist off the model stage
func (artist *Artist) UnstageVoid(stage *Stage) {
	delete(stage.Artists, artist)
	// issue1150
	// delete(stage.Artist_stagedOrder, artist)
	delete(stage.Artists_mapString, artist.Name)
}

// commit artist to the back repo (if it is already staged)
func (artist *Artist) Commit(stage *Stage) *Artist {
	if _, ok := stage.Artists[artist]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArtist(artist)
		}
	}
	return artist
}

func (artist *Artist) CommitVoid(stage *Stage) {
	artist.Commit(stage)
}

func (artist *Artist) StageVoid(stage *Stage) {
	artist.Stage(stage)
}

// Checkout artist to the back repo (if it is already staged)
func (artist *Artist) Checkout(stage *Stage) *Artist {
	if _, ok := stage.Artists[artist]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArtist(artist)
		}
	}
	return artist
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
	if _, ok := stage.ArtistShapes[artistshape]; !ok {
		stage.ArtistShapes[artistshape] = struct{}{}
		stage.ArtistShape_stagedOrder[artistshape] = stage.ArtistShapeOrder
		stage.ArtistShape_orderStaged[stage.ArtistShapeOrder] = artistshape
		stage.ArtistShapeOrder++
	}
	stage.ArtistShapes_mapString[artistshape.Name] = artistshape

	return artistshape
}

// StagePreserveOrder puts artistshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArtistShapeOrder
// - update stage.ArtistShapeOrder accordingly
func (artistshape *ArtistShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ArtistShapes[artistshape]; !ok {
		stage.ArtistShapes[artistshape] = struct{}{}

		if order > stage.ArtistShapeOrder {
			stage.ArtistShapeOrder = order
		}
		stage.ArtistShape_stagedOrder[artistshape] = order
		stage.ArtistShape_orderStaged[order] = artistshape
		stage.ArtistShapeOrder++
	}
	stage.ArtistShapes_mapString[artistshape.Name] = artistshape
}

// Unstage removes artistshape off the model stage
func (artistshape *ArtistShape) Unstage(stage *Stage) *ArtistShape {
	delete(stage.ArtistShapes, artistshape)
	// issue1150
	// delete(stage.ArtistShape_stagedOrder, artistshape)
	delete(stage.ArtistShapes_mapString, artistshape.Name)

	return artistshape
}

// UnstageVoid removes artistshape off the model stage
func (artistshape *ArtistShape) UnstageVoid(stage *Stage) {
	delete(stage.ArtistShapes, artistshape)
	// issue1150
	// delete(stage.ArtistShape_stagedOrder, artistshape)
	delete(stage.ArtistShapes_mapString, artistshape.Name)
}

// commit artistshape to the back repo (if it is already staged)
func (artistshape *ArtistShape) Commit(stage *Stage) *ArtistShape {
	if _, ok := stage.ArtistShapes[artistshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArtistShape(artistshape)
		}
	}
	return artistshape
}

func (artistshape *ArtistShape) CommitVoid(stage *Stage) {
	artistshape.Commit(stage)
}

func (artistshape *ArtistShape) StageVoid(stage *Stage) {
	artistshape.Stage(stage)
}

// Checkout artistshape to the back repo (if it is already staged)
func (artistshape *ArtistShape) Checkout(stage *Stage) *ArtistShape {
	if _, ok := stage.ArtistShapes[artistshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArtistShape(artistshape)
		}
	}
	return artistshape
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
	if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
		stage.ControlPointShapes[controlpointshape] = struct{}{}
		stage.ControlPointShape_stagedOrder[controlpointshape] = stage.ControlPointShapeOrder
		stage.ControlPointShape_orderStaged[stage.ControlPointShapeOrder] = controlpointshape
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
		stage.ControlPointShape_stagedOrder[controlpointshape] = order
		stage.ControlPointShape_orderStaged[order] = controlpointshape
		stage.ControlPointShapeOrder++
	}
	stage.ControlPointShapes_mapString[controlpointshape.Name] = controlpointshape
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	delete(stage.ControlPointShapes, controlpointshape)
	// issue1150
	// delete(stage.ControlPointShape_stagedOrder, controlpointshape)
	delete(stage.ControlPointShapes_mapString, controlpointshape.Name)

	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	delete(stage.ControlPointShapes, controlpointshape)
	// issue1150
	// delete(stage.ControlPointShape_stagedOrder, controlpointshape)
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
func (controlpointshape *ControlPointShape) SetName(name string) {
	controlpointshape.Name = name
}

// Stage puts desk to the model stage
func (desk *Desk) Stage(stage *Stage) *Desk {
	if _, ok := stage.Desks[desk]; !ok {
		stage.Desks[desk] = struct{}{}
		stage.Desk_stagedOrder[desk] = stage.DeskOrder
		stage.Desk_orderStaged[stage.DeskOrder] = desk
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
		stage.Desk_stagedOrder[desk] = order
		stage.Desk_orderStaged[order] = desk
		stage.DeskOrder++
	}
	stage.Desks_mapString[desk.Name] = desk
}

// Unstage removes desk off the model stage
func (desk *Desk) Unstage(stage *Stage) *Desk {
	delete(stage.Desks, desk)
	// issue1150
	// delete(stage.Desk_stagedOrder, desk)
	delete(stage.Desks_mapString, desk.Name)

	return desk
}

// UnstageVoid removes desk off the model stage
func (desk *Desk) UnstageVoid(stage *Stage) {
	delete(stage.Desks, desk)
	// issue1150
	// delete(stage.Desk_stagedOrder, desk)
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
func (desk *Desk) SetName(name string) {
	desk.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}
		stage.Diagram_stagedOrder[diagram] = stage.DiagramOrder
		stage.Diagram_orderStaged[stage.DiagramOrder] = diagram
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
		stage.Diagram_stagedOrder[diagram] = order
		stage.Diagram_orderStaged[order] = diagram
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)

	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
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
func (diagram *Diagram) SetName(name string) {
	diagram.Name = name
}

// Stage puts influence to the model stage
func (influence *Influence) Stage(stage *Stage) *Influence {
	if _, ok := stage.Influences[influence]; !ok {
		stage.Influences[influence] = struct{}{}
		stage.Influence_stagedOrder[influence] = stage.InfluenceOrder
		stage.Influence_orderStaged[stage.InfluenceOrder] = influence
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
		stage.Influence_stagedOrder[influence] = order
		stage.Influence_orderStaged[order] = influence
		stage.InfluenceOrder++
	}
	stage.Influences_mapString[influence.Name] = influence
}

// Unstage removes influence off the model stage
func (influence *Influence) Unstage(stage *Stage) *Influence {
	delete(stage.Influences, influence)
	// issue1150
	// delete(stage.Influence_stagedOrder, influence)
	delete(stage.Influences_mapString, influence.Name)

	return influence
}

// UnstageVoid removes influence off the model stage
func (influence *Influence) UnstageVoid(stage *Stage) {
	delete(stage.Influences, influence)
	// issue1150
	// delete(stage.Influence_stagedOrder, influence)
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
func (influence *Influence) SetName(name string) {
	influence.Name = name
}

// Stage puts influenceshape to the model stage
func (influenceshape *InfluenceShape) Stage(stage *Stage) *InfluenceShape {
	if _, ok := stage.InfluenceShapes[influenceshape]; !ok {
		stage.InfluenceShapes[influenceshape] = struct{}{}
		stage.InfluenceShape_stagedOrder[influenceshape] = stage.InfluenceShapeOrder
		stage.InfluenceShape_orderStaged[stage.InfluenceShapeOrder] = influenceshape
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
		stage.InfluenceShape_stagedOrder[influenceshape] = order
		stage.InfluenceShape_orderStaged[order] = influenceshape
		stage.InfluenceShapeOrder++
	}
	stage.InfluenceShapes_mapString[influenceshape.Name] = influenceshape
}

// Unstage removes influenceshape off the model stage
func (influenceshape *InfluenceShape) Unstage(stage *Stage) *InfluenceShape {
	delete(stage.InfluenceShapes, influenceshape)
	// issue1150
	// delete(stage.InfluenceShape_stagedOrder, influenceshape)
	delete(stage.InfluenceShapes_mapString, influenceshape.Name)

	return influenceshape
}

// UnstageVoid removes influenceshape off the model stage
func (influenceshape *InfluenceShape) UnstageVoid(stage *Stage) {
	delete(stage.InfluenceShapes, influenceshape)
	// issue1150
	// delete(stage.InfluenceShape_stagedOrder, influenceshape)
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
func (influenceshape *InfluenceShape) SetName(name string) {
	influenceshape.Name = name
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

// Stage puts movement to the model stage
func (movement *Movement) Stage(stage *Stage) *Movement {
	if _, ok := stage.Movements[movement]; !ok {
		stage.Movements[movement] = struct{}{}
		stage.Movement_stagedOrder[movement] = stage.MovementOrder
		stage.Movement_orderStaged[stage.MovementOrder] = movement
		stage.MovementOrder++
	}
	stage.Movements_mapString[movement.Name] = movement

	return movement
}

// StagePreserveOrder puts movement to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MovementOrder
// - update stage.MovementOrder accordingly
func (movement *Movement) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Movements[movement]; !ok {
		stage.Movements[movement] = struct{}{}

		if order > stage.MovementOrder {
			stage.MovementOrder = order
		}
		stage.Movement_stagedOrder[movement] = order
		stage.Movement_orderStaged[order] = movement
		stage.MovementOrder++
	}
	stage.Movements_mapString[movement.Name] = movement
}

// Unstage removes movement off the model stage
func (movement *Movement) Unstage(stage *Stage) *Movement {
	delete(stage.Movements, movement)
	// issue1150
	// delete(stage.Movement_stagedOrder, movement)
	delete(stage.Movements_mapString, movement.Name)

	return movement
}

// UnstageVoid removes movement off the model stage
func (movement *Movement) UnstageVoid(stage *Stage) {
	delete(stage.Movements, movement)
	// issue1150
	// delete(stage.Movement_stagedOrder, movement)
	delete(stage.Movements_mapString, movement.Name)
}

// commit movement to the back repo (if it is already staged)
func (movement *Movement) Commit(stage *Stage) *Movement {
	if _, ok := stage.Movements[movement]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMovement(movement)
		}
	}
	return movement
}

func (movement *Movement) CommitVoid(stage *Stage) {
	movement.Commit(stage)
}

func (movement *Movement) StageVoid(stage *Stage) {
	movement.Stage(stage)
}

// Checkout movement to the back repo (if it is already staged)
func (movement *Movement) Checkout(stage *Stage) *Movement {
	if _, ok := stage.Movements[movement]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMovement(movement)
		}
	}
	return movement
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
	if _, ok := stage.MovementShapes[movementshape]; !ok {
		stage.MovementShapes[movementshape] = struct{}{}
		stage.MovementShape_stagedOrder[movementshape] = stage.MovementShapeOrder
		stage.MovementShape_orderStaged[stage.MovementShapeOrder] = movementshape
		stage.MovementShapeOrder++
	}
	stage.MovementShapes_mapString[movementshape.Name] = movementshape

	return movementshape
}

// StagePreserveOrder puts movementshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MovementShapeOrder
// - update stage.MovementShapeOrder accordingly
func (movementshape *MovementShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MovementShapes[movementshape]; !ok {
		stage.MovementShapes[movementshape] = struct{}{}

		if order > stage.MovementShapeOrder {
			stage.MovementShapeOrder = order
		}
		stage.MovementShape_stagedOrder[movementshape] = order
		stage.MovementShape_orderStaged[order] = movementshape
		stage.MovementShapeOrder++
	}
	stage.MovementShapes_mapString[movementshape.Name] = movementshape
}

// Unstage removes movementshape off the model stage
func (movementshape *MovementShape) Unstage(stage *Stage) *MovementShape {
	delete(stage.MovementShapes, movementshape)
	// issue1150
	// delete(stage.MovementShape_stagedOrder, movementshape)
	delete(stage.MovementShapes_mapString, movementshape.Name)

	return movementshape
}

// UnstageVoid removes movementshape off the model stage
func (movementshape *MovementShape) UnstageVoid(stage *Stage) {
	delete(stage.MovementShapes, movementshape)
	// issue1150
	// delete(stage.MovementShape_stagedOrder, movementshape)
	delete(stage.MovementShapes_mapString, movementshape.Name)
}

// commit movementshape to the back repo (if it is already staged)
func (movementshape *MovementShape) Commit(stage *Stage) *MovementShape {
	if _, ok := stage.MovementShapes[movementshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMovementShape(movementshape)
		}
	}
	return movementshape
}

func (movementshape *MovementShape) CommitVoid(stage *Stage) {
	movementshape.Commit(stage)
}

func (movementshape *MovementShape) StageVoid(stage *Stage) {
	movementshape.Stage(stage)
}

// Checkout movementshape to the back repo (if it is already staged)
func (movementshape *MovementShape) Checkout(stage *Stage) *MovementShape {
	if _, ok := stage.MovementShapes[movementshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMovementShape(movementshape)
		}
	}
	return movementshape
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
	if _, ok := stage.Places[place]; !ok {
		stage.Places[place] = struct{}{}
		stage.Place_stagedOrder[place] = stage.PlaceOrder
		stage.Place_orderStaged[stage.PlaceOrder] = place
		stage.PlaceOrder++
	}
	stage.Places_mapString[place.Name] = place

	return place
}

// StagePreserveOrder puts place to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlaceOrder
// - update stage.PlaceOrder accordingly
func (place *Place) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Places[place]; !ok {
		stage.Places[place] = struct{}{}

		if order > stage.PlaceOrder {
			stage.PlaceOrder = order
		}
		stage.Place_stagedOrder[place] = order
		stage.Place_orderStaged[order] = place
		stage.PlaceOrder++
	}
	stage.Places_mapString[place.Name] = place
}

// Unstage removes place off the model stage
func (place *Place) Unstage(stage *Stage) *Place {
	delete(stage.Places, place)
	// issue1150
	// delete(stage.Place_stagedOrder, place)
	delete(stage.Places_mapString, place.Name)

	return place
}

// UnstageVoid removes place off the model stage
func (place *Place) UnstageVoid(stage *Stage) {
	delete(stage.Places, place)
	// issue1150
	// delete(stage.Place_stagedOrder, place)
	delete(stage.Places_mapString, place.Name)
}

// commit place to the back repo (if it is already staged)
func (place *Place) Commit(stage *Stage) *Place {
	if _, ok := stage.Places[place]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlace(place)
		}
	}
	return place
}

func (place *Place) CommitVoid(stage *Stage) {
	place.Commit(stage)
}

func (place *Place) StageVoid(stage *Stage) {
	place.Stage(stage)
}

// Checkout place to the back repo (if it is already staged)
func (place *Place) Checkout(stage *Stage) *Place {
	if _, ok := stage.Places[place]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlace(place)
		}
	}
	return place
}

// for satisfaction of GongStruct interface
func (place *Place) GetName() (res string) {
	return place.Name
}

// for satisfaction of GongStruct interface
func (place *Place) SetName(name string) {
	place.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMArtefactType(ArtefactType *ArtefactType)
	CreateORMArtefactTypeShape(ArtefactTypeShape *ArtefactTypeShape)
	CreateORMArtist(Artist *Artist)
	CreateORMArtistShape(ArtistShape *ArtistShape)
	CreateORMControlPointShape(ControlPointShape *ControlPointShape)
	CreateORMDesk(Desk *Desk)
	CreateORMDiagram(Diagram *Diagram)
	CreateORMInfluence(Influence *Influence)
	CreateORMInfluenceShape(InfluenceShape *InfluenceShape)
	CreateORMLibrary(Library *Library)
	CreateORMMovement(Movement *Movement)
	CreateORMMovementShape(MovementShape *MovementShape)
	CreateORMPlace(Place *Place)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMArtefactType(ArtefactType *ArtefactType)
	DeleteORMArtefactTypeShape(ArtefactTypeShape *ArtefactTypeShape)
	DeleteORMArtist(Artist *Artist)
	DeleteORMArtistShape(ArtistShape *ArtistShape)
	DeleteORMControlPointShape(ControlPointShape *ControlPointShape)
	DeleteORMDesk(Desk *Desk)
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMInfluence(Influence *Influence)
	DeleteORMInfluenceShape(InfluenceShape *InfluenceShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMMovement(Movement *Movement)
	DeleteORMMovementShape(MovementShape *MovementShape)
	DeleteORMPlace(Place *Place)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.ArtefactTypes = make(map[*ArtefactType]struct{})
	stage.ArtefactTypes_mapString = make(map[string]*ArtefactType)
	stage.ArtefactType_stagedOrder = make(map[*ArtefactType]uint)
	stage.ArtefactTypeOrder = 0

	stage.ArtefactTypeShapes = make(map[*ArtefactTypeShape]struct{})
	stage.ArtefactTypeShapes_mapString = make(map[string]*ArtefactTypeShape)
	stage.ArtefactTypeShape_stagedOrder = make(map[*ArtefactTypeShape]uint)
	stage.ArtefactTypeShapeOrder = 0

	stage.Artists = make(map[*Artist]struct{})
	stage.Artists_mapString = make(map[string]*Artist)
	stage.Artist_stagedOrder = make(map[*Artist]uint)
	stage.ArtistOrder = 0

	stage.ArtistShapes = make(map[*ArtistShape]struct{})
	stage.ArtistShapes_mapString = make(map[string]*ArtistShape)
	stage.ArtistShape_stagedOrder = make(map[*ArtistShape]uint)
	stage.ArtistShapeOrder = 0

	stage.ControlPointShapes = make(map[*ControlPointShape]struct{})
	stage.ControlPointShapes_mapString = make(map[string]*ControlPointShape)
	stage.ControlPointShape_stagedOrder = make(map[*ControlPointShape]uint)
	stage.ControlPointShapeOrder = 0

	stage.Desks = make(map[*Desk]struct{})
	stage.Desks_mapString = make(map[string]*Desk)
	stage.Desk_stagedOrder = make(map[*Desk]uint)
	stage.DeskOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.Diagram_stagedOrder = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Influences = make(map[*Influence]struct{})
	stage.Influences_mapString = make(map[string]*Influence)
	stage.Influence_stagedOrder = make(map[*Influence]uint)
	stage.InfluenceOrder = 0

	stage.InfluenceShapes = make(map[*InfluenceShape]struct{})
	stage.InfluenceShapes_mapString = make(map[string]*InfluenceShape)
	stage.InfluenceShape_stagedOrder = make(map[*InfluenceShape]uint)
	stage.InfluenceShapeOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Movements = make(map[*Movement]struct{})
	stage.Movements_mapString = make(map[string]*Movement)
	stage.Movement_stagedOrder = make(map[*Movement]uint)
	stage.MovementOrder = 0

	stage.MovementShapes = make(map[*MovementShape]struct{})
	stage.MovementShapes_mapString = make(map[string]*MovementShape)
	stage.MovementShape_stagedOrder = make(map[*MovementShape]uint)
	stage.MovementShapeOrder = 0

	stage.Places = make(map[*Place]struct{})
	stage.Places_mapString = make(map[string]*Place)
	stage.Place_stagedOrder = make(map[*Place]uint)
	stage.PlaceOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.ArtefactTypes = nil
	stage.ArtefactTypes_mapString = nil

	stage.ArtefactTypeShapes = nil
	stage.ArtefactTypeShapes_mapString = nil

	stage.Artists = nil
	stage.Artists_mapString = nil

	stage.ArtistShapes = nil
	stage.ArtistShapes_mapString = nil

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

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Movements = nil
	stage.Movements_mapString = nil

	stage.MovementShapes = nil
	stage.MovementShapes_mapString = nil

	stage.Places = nil
	stage.Places_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for artefacttype := range stage.ArtefactTypes {
		artefacttype.Unstage(stage)
	}

	for artefacttypeshape := range stage.ArtefactTypeShapes {
		artefacttypeshape.Unstage(stage)
	}

	for artist := range stage.Artists {
		artist.Unstage(stage)
	}

	for artistshape := range stage.ArtistShapes {
		artistshape.Unstage(stage)
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

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for movement := range stage.Movements {
		movement.Unstage(stage)
	}

	for movementshape := range stage.MovementShapes {
		movementshape.Unstage(stage)
	}

	for place := range stage.Places {
		place.Unstage(stage)
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
	case map[*ArtefactType]any:
		return any(&stage.ArtefactTypes).(*Type)
	case map[*ArtefactTypeShape]any:
		return any(&stage.ArtefactTypeShapes).(*Type)
	case map[*Artist]any:
		return any(&stage.Artists).(*Type)
	case map[*ArtistShape]any:
		return any(&stage.ArtistShapes).(*Type)
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
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Movement]any:
		return any(&stage.Movements).(*Type)
	case map[*MovementShape]any:
		return any(&stage.MovementShapes).(*Type)
	case map[*Place]any:
		return any(&stage.Places).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ArtefactType:
		return any(&stage.ArtefactTypes).(*map[*Type]struct{})
	case ArtefactTypeShape:
		return any(&stage.ArtefactTypeShapes).(*map[*Type]struct{})
	case Artist:
		return any(&stage.Artists).(*map[*Type]struct{})
	case ArtistShape:
		return any(&stage.ArtistShapes).(*map[*Type]struct{})
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
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Movement:
		return any(&stage.Movements).(*map[*Type]struct{})
	case MovementShape:
		return any(&stage.MovementShapes).(*map[*Type]struct{})
	case Place:
		return any(&stage.Places).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case ArtefactType:
		return any(&stage.ArtefactTypes_mapString).(*map[string]*Type)
	case ArtefactTypeShape:
		return any(&stage.ArtefactTypeShapes_mapString).(*map[string]*Type)
	case Artist:
		return any(&stage.Artists_mapString).(*map[string]*Type)
	case ArtistShape:
		return any(&stage.ArtistShapes_mapString).(*map[string]*Type)
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
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Movement:
		return any(&stage.Movements_mapString).(*map[string]*Type)
	case MovementShape:
		return any(&stage.MovementShapes_mapString).(*map[string]*Type)
	case Place:
		return any(&stage.Places_mapString).(*map[string]*Type)
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
	case ArtefactType:
		return any(&ArtefactType{
			// Initialisation of associations
		}).(*Type)
	case ArtefactTypeShape:
		return any(&ArtefactTypeShape{
			// Initialisation of associations
			// field is initialized with an instance of ArtefactType with the name of the field
			ArtefactType: &ArtefactType{Name: "ArtefactType"},
		}).(*Type)
	case Artist:
		return any(&Artist{
			// Initialisation of associations
			// field is initialized with an instance of Place with the name of the field
			Place: &Place{Name: "Place"},
		}).(*Type)
	case ArtistShape:
		return any(&ArtistShape{
			// Initialisation of associations
			// field is initialized with an instance of Artist with the name of the field
			Artist: &Artist{Name: "Artist"},
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
			// field is initialized with an instance of MovementShape with the name of the field
			MovementShapes: []*MovementShape{{Name: "MovementShapes"}},
			// field is initialized with an instance of ArtefactTypeShape with the name of the field
			ArtefactTypeShapes: []*ArtefactTypeShape{{Name: "ArtefactTypeShapes"}},
			// field is initialized with an instance of ArtistShape with the name of the field
			ArtistShapes: []*ArtistShape{{Name: "ArtistShapes"}},
			// field is initialized with an instance of InfluenceShape with the name of the field
			InfluenceShapes: []*InfluenceShape{{Name: "InfluenceShapes"}},
		}).(*Type)
	case Influence:
		return any(&Influence{
			// Initialisation of associations
			// field is initialized with an instance of Movement with the name of the field
			SourceMovement: &Movement{Name: "SourceMovement"},
			// field is initialized with an instance of ArtefactType with the name of the field
			SourceArtefactType: &ArtefactType{Name: "SourceArtefactType"},
			// field is initialized with an instance of Artist with the name of the field
			SourceArtist: &Artist{Name: "SourceArtist"},
			// field is initialized with an instance of Movement with the name of the field
			TargetMovement: &Movement{Name: "TargetMovement"},
			// field is initialized with an instance of ArtefactType with the name of the field
			TargetArtefactType: &ArtefactType{Name: "TargetArtefactType"},
			// field is initialized with an instance of Artist with the name of the field
			TargetArtist: &Artist{Name: "TargetArtist"},
		}).(*Type)
	case InfluenceShape:
		return any(&InfluenceShape{
			// Initialisation of associations
			// field is initialized with an instance of Influence with the name of the field
			Influence: &Influence{Name: "Influence"},
			// field is initialized with an instance of ControlPointShape with the name of the field
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Movement:
		return any(&Movement{
			// Initialisation of associations
			// field is initialized with an instance of Place with the name of the field
			Places: []*Place{{Name: "Places"}},
		}).(*Type)
	case MovementShape:
		return any(&MovementShape{
			// Initialisation of associations
			// field is initialized with an instance of Movement with the name of the field
			Movement: &Movement{Name: "Movement"},
		}).(*Type)
	case Place:
		return any(&Place{
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
	case "LayoutDirection":
		enum := artefacttype.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := artist.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := diagram.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := influence.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
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
	case "LayoutDirection":
		enum := movement.LayoutDirection
		res.valueString = enum.ToCodeString()
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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (artefacttype *ArtefactType) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		artefacttype.Name = value.GetValueString()
	case "ComputedPrefix":
		artefacttype.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		artefacttype.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		artefacttype.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (artefacttypeshape *ArtefactTypeShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		artefacttypeshape.Name = value.GetValueString()
	case "ArtefactType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			artefacttypeshape.ArtefactType = nil
			for __instance__ := range stage.ArtefactTypes {
				if stage.ArtefactType_stagedOrder[__instance__] == uint(id) {
					artefacttypeshape.ArtefactType = __instance__
					break
				}
			}
		}
	case "X":
		artefacttypeshape.X = value.GetValueFloat()
	case "Y":
		artefacttypeshape.Y = value.GetValueFloat()
	case "Width":
		artefacttypeshape.Width = value.GetValueFloat()
	case "Height":
		artefacttypeshape.Height = value.GetValueFloat()
	case "IsHidden":
		artefacttypeshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (artist *Artist) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		artist.Name = value.GetValueString()
	case "ComputedPrefix":
		artist.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		artist.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		artist.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsDead":
		artist.IsDead = value.GetValueBool()
	case "Place":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			artist.Place = nil
			for __instance__ := range stage.Places {
				if stage.Place_stagedOrder[__instance__] == uint(id) {
					artist.Place = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (artistshape *ArtistShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		artistshape.Name = value.GetValueString()
	case "Artist":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			artistshape.Artist = nil
			for __instance__ := range stage.Artists {
				if stage.Artist_stagedOrder[__instance__] == uint(id) {
					artistshape.Artist = __instance__
					break
				}
			}
		}
	case "X":
		artistshape.X = value.GetValueFloat()
	case "Y":
		artistshape.Y = value.GetValueFloat()
	case "Width":
		artistshape.Width = value.GetValueFloat()
	case "Height":
		artistshape.Height = value.GetValueFloat()
	case "IsHidden":
		artistshape.IsHidden = value.GetValueBool()
	case "ImagePng_X":
		artistshape.ImagePng_X = value.GetValueFloat()
	case "ImagePng_Y":
		artistshape.ImagePng_Y = value.GetValueFloat()
	case "ImagePng_Width":
		artistshape.ImagePng_Width = value.GetValueFloat()
	case "ImagePng_Height":
		artistshape.ImagePng_Height = value.GetValueFloat()
	case "ImagePng_X_Offset":
		artistshape.ImagePng_X_Offset = value.GetValueFloat()
	case "ImagePng_Y_Offset":
		artistshape.ImagePng_Y_Offset = value.GetValueFloat()
	case "ImagePng_RectAnchorType":
		artistshape.ImagePng_RectAnchorType.FromCodeString(value.GetValueString())
	case "ImagePngBase64Content":
		artistshape.ImagePngBase64Content = value.GetValueString()
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
				if stage.Diagram_stagedOrder[__instance__] == uint(id) {
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
	case "ComputedPrefix":
		diagram.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagram.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		diagram.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsChecked":
		diagram.IsChecked = value.GetValueBool()
	case "MovementShapes":
		diagram.MovementShapes = make([]*MovementShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.MovementShapes {
					if stage.MovementShape_stagedOrder[__instance__] == uint(id) {
						diagram.MovementShapes = append(diagram.MovementShapes, __instance__)
						break
					}
				}
			}
		}
	case "ArtefactTypeShapes":
		diagram.ArtefactTypeShapes = make([]*ArtefactTypeShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ArtefactTypeShapes {
					if stage.ArtefactTypeShape_stagedOrder[__instance__] == uint(id) {
						diagram.ArtefactTypeShapes = append(diagram.ArtefactTypeShapes, __instance__)
						break
					}
				}
			}
		}
	case "ArtistShapes":
		diagram.ArtistShapes = make([]*ArtistShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ArtistShapes {
					if stage.ArtistShape_stagedOrder[__instance__] == uint(id) {
						diagram.ArtistShapes = append(diagram.ArtistShapes, __instance__)
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
					if stage.InfluenceShape_stagedOrder[__instance__] == uint(id) {
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
	case "IsMovementCategoryNodeExpanded":
		diagram.IsMovementCategoryNodeExpanded = value.GetValueBool()
	case "IsArtefactTypeCategoryNodeExpanded":
		diagram.IsArtefactTypeCategoryNodeExpanded = value.GetValueBool()
	case "IsArtistCategoryNodeExpanded":
		diagram.IsArtistCategoryNodeExpanded = value.GetValueBool()
	case "IsInfluenceCategoryNodeExpanded":
		diagram.IsInfluenceCategoryNodeExpanded = value.GetValueBool()
	case "IsMovementCategoryHidden":
		diagram.IsMovementCategoryHidden = value.GetValueBool()
	case "IsArtefactTypeCategoryHidden":
		diagram.IsArtefactTypeCategoryHidden = value.GetValueBool()
	case "IsArtistCategoryHidden":
		diagram.IsArtistCategoryHidden = value.GetValueBool()
	case "IsInfluenceCategoryHidden":
		diagram.IsInfluenceCategoryHidden = value.GetValueBool()
	case "NbYearsForIntervals":
		diagram.NbYearsForIntervals = int(value.GetValueInt())
	case "XMargin":
		diagram.XMargin = value.GetValueFloat()
	case "YMargin":
		diagram.YMargin = value.GetValueFloat()
	case "Height":
		diagram.Height = value.GetValueFloat()
	case "NextVerticalDateXMargin":
		diagram.NextVerticalDateXMargin = value.GetValueFloat()
	case "RedColorCode":
		diagram.RedColorCode = value.GetValueString()
	case "BackgroundGreyColorCode":
		diagram.BackgroundGreyColorCode = value.GetValueString()
	case "GrayColorCode":
		diagram.GrayColorCode = value.GetValueString()
	case "BottomBoxYOffset":
		diagram.BottomBoxYOffset = value.GetValueFloat()
	case "BottomBoxWidth":
		diagram.BottomBoxWidth = value.GetValueFloat()
	case "BottomBoxHeigth":
		diagram.BottomBoxHeigth = value.GetValueFloat()
	case "BottomBoxFontSize":
		diagram.BottomBoxFontSize = value.GetValueString()
	case "BottomBoxFontWeigth":
		diagram.BottomBoxFontWeigth = value.GetValueString()
	case "BottomBoxFontFamily":
		diagram.BottomBoxFontFamily = value.GetValueString()
	case "BottomBoxLetterSpacing":
		diagram.BottomBoxLetterSpacing = value.GetValueString()
	case "BottomBoxLetterColorCode":
		diagram.BottomBoxLetterColorCode = value.GetValueString()
	case "MovementRectAnchorType":
		diagram.MovementRectAnchorType.FromCodeString(value.GetValueString())
	case "MovementTextAnchorType":
		diagram.MovementTextAnchorType.FromCodeString(value.GetValueString())
	case "MovementDominantBaselineType":
		diagram.MovementDominantBaselineType.FromCodeString(value.GetValueString())
	case "MovementFontSize":
		diagram.MovementFontSize = value.GetValueString()
	case "MajorMovementFontSize":
		diagram.MajorMovementFontSize = value.GetValueString()
	case "MinorMovementFontSize":
		diagram.MinorMovementFontSize = value.GetValueString()
	case "MovementFontWeigth":
		diagram.MovementFontWeigth = value.GetValueString()
	case "MovementFontFamily":
		diagram.MovementFontFamily = value.GetValueString()
	case "MovementLetterSpacing":
		diagram.MovementLetterSpacing = value.GetValueString()
	case "AbstractMovementFontSize":
		diagram.AbstractMovementFontSize = value.GetValueString()
	case "AbstractMovementRectAnchorType":
		diagram.AbstractMovementRectAnchorType.FromCodeString(value.GetValueString())
	case "AbstractMovementTextAnchorType":
		diagram.AbstractMovementTextAnchorType.FromCodeString(value.GetValueString())
	case "AbstractDominantBaselineType":
		diagram.AbstractDominantBaselineType.FromCodeString(value.GetValueString())
	case "MovementDateRectAnchorType":
		diagram.MovementDateRectAnchorType.FromCodeString(value.GetValueString())
	case "MovementDateTextAnchorType":
		diagram.MovementDateTextAnchorType.FromCodeString(value.GetValueString())
	case "MovementDateTextDominantBaselineType":
		diagram.MovementDateTextDominantBaselineType.FromCodeString(value.GetValueString())
	case "MovementDateAndPlacesFontSize":
		diagram.MovementDateAndPlacesFontSize = value.GetValueString()
	case "MovementDateAndPlacesFontWeigth":
		diagram.MovementDateAndPlacesFontWeigth = value.GetValueString()
	case "MovementDateAndPlacesFontFamily":
		diagram.MovementDateAndPlacesFontFamily = value.GetValueString()
	case "MovementDateAndPlacesLetterSpacing":
		diagram.MovementDateAndPlacesLetterSpacing = value.GetValueString()
	case "MovementBelowArcY_Offset":
		diagram.MovementBelowArcY_Offset = value.GetValueFloat()
	case "MovementBelowArcY_OffsetPerPlace":
		diagram.MovementBelowArcY_OffsetPerPlace = value.GetValueFloat()
	case "MovementPlacesRectAnchorType":
		diagram.MovementPlacesRectAnchorType.FromCodeString(value.GetValueString())
	case "MovementPlacesTextAnchorType":
		diagram.MovementPlacesTextAnchorType.FromCodeString(value.GetValueString())
	case "MovementPlacesDominantBaselineType":
		diagram.MovementPlacesDominantBaselineType.FromCodeString(value.GetValueString())
	case "ArtefactTypeFontSize":
		diagram.ArtefactTypeFontSize = value.GetValueString()
	case "ArtefactTypeFontWeigth":
		diagram.ArtefactTypeFontWeigth = value.GetValueString()
	case "ArtefactTypeFontFamily":
		diagram.ArtefactTypeFontFamily = value.GetValueString()
	case "ArtefactTypeLetterSpacing":
		diagram.ArtefactTypeLetterSpacing = value.GetValueString()
	case "ArtefactTypeRectAnchorType":
		diagram.ArtefactTypeRectAnchorType.FromCodeString(value.GetValueString())
	case "ArtefactDominantBaselineType":
		diagram.ArtefactDominantBaselineType.FromCodeString(value.GetValueString())
	case "ArtefactTypeStrokeWidth":
		diagram.ArtefactTypeStrokeWidth = value.GetValueFloat()
	case "ArtistRectAnchorType":
		diagram.ArtistRectAnchorType.FromCodeString(value.GetValueString())
	case "ArtistTextAnchorType":
		diagram.ArtistTextAnchorType.FromCodeString(value.GetValueString())
	case "ArtistDominantBaselineType":
		diagram.ArtistDominantBaselineType.FromCodeString(value.GetValueString())
	case "ArtistFontSize":
		diagram.ArtistFontSize = value.GetValueString()
	case "MajorArtistFontSize":
		diagram.MajorArtistFontSize = value.GetValueString()
	case "MinorArtistFontSize":
		diagram.MinorArtistFontSize = value.GetValueString()
	case "ArtistFontWeigth":
		diagram.ArtistFontWeigth = value.GetValueString()
	case "ArtistFontFamily":
		diagram.ArtistFontFamily = value.GetValueString()
	case "ArtistLetterSpacing":
		diagram.ArtistLetterSpacing = value.GetValueString()
	case "ArtistDateRectAnchorType":
		diagram.ArtistDateRectAnchorType.FromCodeString(value.GetValueString())
	case "ArtistDateTextAnchorType":
		diagram.ArtistDateTextAnchorType.FromCodeString(value.GetValueString())
	case "ArtistDateDominantBaselineType":
		diagram.ArtistDateDominantBaselineType.FromCodeString(value.GetValueString())
	case "ArtistDateAndPlacesFontSize":
		diagram.ArtistDateAndPlacesFontSize = value.GetValueString()
	case "ArtistDateAndPlacesFontWeigth":
		diagram.ArtistDateAndPlacesFontWeigth = value.GetValueString()
	case "ArtistDateAndPlacesFontFamily":
		diagram.ArtistDateAndPlacesFontFamily = value.GetValueString()
	case "ArtistDateAndPlacesLetterSpacing":
		diagram.ArtistDateAndPlacesLetterSpacing = value.GetValueString()
	case "ArtistPlacesRectAnchorType":
		diagram.ArtistPlacesRectAnchorType.FromCodeString(value.GetValueString())
	case "ArtistPlacesTextAnchorType":
		diagram.ArtistPlacesTextAnchorType.FromCodeString(value.GetValueString())
	case "ArtistPlacesDominantBaselineType":
		diagram.ArtistPlacesDominantBaselineType.FromCodeString(value.GetValueString())
	case "InfluenceArrowSize":
		diagram.InfluenceArrowSize = value.GetValueFloat()
	case "InfluenceArrowStartOffset":
		diagram.InfluenceArrowStartOffset = value.GetValueFloat()
	case "InfluenceArrowEndOffset":
		diagram.InfluenceArrowEndOffset = value.GetValueFloat()
	case "InfluenceCornerRadius":
		diagram.InfluenceCornerRadius = value.GetValueFloat()
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
	case "ComputedPrefix":
		influence.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		influence.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		influence.LayoutDirection.FromCodeString(value.GetValueString())
	case "SourceMovement":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceMovement = nil
			for __instance__ := range stage.Movements {
				if stage.Movement_stagedOrder[__instance__] == uint(id) {
					influence.SourceMovement = __instance__
					break
				}
			}
		}
	case "SourceArtefactType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceArtefactType = nil
			for __instance__ := range stage.ArtefactTypes {
				if stage.ArtefactType_stagedOrder[__instance__] == uint(id) {
					influence.SourceArtefactType = __instance__
					break
				}
			}
		}
	case "SourceArtist":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.SourceArtist = nil
			for __instance__ := range stage.Artists {
				if stage.Artist_stagedOrder[__instance__] == uint(id) {
					influence.SourceArtist = __instance__
					break
				}
			}
		}
	case "TargetMovement":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetMovement = nil
			for __instance__ := range stage.Movements {
				if stage.Movement_stagedOrder[__instance__] == uint(id) {
					influence.TargetMovement = __instance__
					break
				}
			}
		}
	case "TargetArtefactType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetArtefactType = nil
			for __instance__ := range stage.ArtefactTypes {
				if stage.ArtefactType_stagedOrder[__instance__] == uint(id) {
					influence.TargetArtefactType = __instance__
					break
				}
			}
		}
	case "TargetArtist":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			influence.TargetArtist = nil
			for __instance__ := range stage.Artists {
				if stage.Artist_stagedOrder[__instance__] == uint(id) {
					influence.TargetArtist = __instance__
					break
				}
			}
		}
	case "IsHypothtical":
		influence.IsHypothtical = value.GetValueBool()
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
				if stage.Influence_stagedOrder[__instance__] == uint(id) {
					influenceshape.Influence = __instance__
					break
				}
			}
		}
	case "IsHidden":
		influenceshape.IsHidden = value.GetValueBool()
	case "ControlPointShapes":
		influenceshape.ControlPointShapes = make([]*ControlPointShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlPointShapes {
					if stage.ControlPointShape_stagedOrder[__instance__] == uint(id) {
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

func (library *Library) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		library.Name = value.GetValueString()
	case "Description":
		library.Description = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
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
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (movement *Movement) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		movement.Name = value.GetValueString()
	case "ComputedPrefix":
		movement.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		movement.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		movement.LayoutDirection.FromCodeString(value.GetValueString())
	case "HideDate":
		movement.HideDate = value.GetValueBool()
	case "Places":
		movement.Places = make([]*Place, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Places {
					if stage.Place_stagedOrder[__instance__] == uint(id) {
						movement.Places = append(movement.Places, __instance__)
						break
					}
				}
			}
		}
	case "HasTaxonomicFilter":
		movement.HasTaxonomicFilter = value.GetValueBool()
	case "TaxonomicFilter":
		movement.TaxonomicFilter = value.GetValueString()
	case "IsFeatured":
		movement.IsFeatured = value.GetValueBool()
	case "FeaturePrefix":
		movement.FeaturePrefix = value.GetValueString()
	case "IsMajor":
		movement.IsMajor = value.GetValueBool()
	case "IsMinor":
		movement.IsMinor = value.GetValueBool()
	case "AdditionnalName":
		movement.AdditionnalName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (movementshape *MovementShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		movementshape.Name = value.GetValueString()
	case "Movement":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			movementshape.Movement = nil
			for __instance__ := range stage.Movements {
				if stage.Movement_stagedOrder[__instance__] == uint(id) {
					movementshape.Movement = __instance__
					break
				}
			}
		}
	case "X":
		movementshape.X = value.GetValueFloat()
	case "Y":
		movementshape.Y = value.GetValueFloat()
	case "Width":
		movementshape.Width = value.GetValueFloat()
	case "Height":
		movementshape.Height = value.GetValueFloat()
	case "IsHidden":
		movementshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (place *Place) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		place.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.ArtefactTypes_mapString = make(map[string]*ArtefactType)
	for artefacttype := range stage.ArtefactTypes {
		stage.ArtefactTypes_mapString[artefacttype.Name] = artefacttype
	}

	stage.ArtefactTypeShapes_mapString = make(map[string]*ArtefactTypeShape)
	for artefacttypeshape := range stage.ArtefactTypeShapes {
		stage.ArtefactTypeShapes_mapString[artefacttypeshape.Name] = artefacttypeshape
	}

	stage.Artists_mapString = make(map[string]*Artist)
	for artist := range stage.Artists {
		stage.Artists_mapString[artist.Name] = artist
	}

	stage.ArtistShapes_mapString = make(map[string]*ArtistShape)
	for artistshape := range stage.ArtistShapes {
		stage.ArtistShapes_mapString[artistshape.Name] = artistshape
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

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Movements_mapString = make(map[string]*Movement)
	for movement := range stage.Movements {
		stage.Movements_mapString[movement.Name] = movement
	}

	stage.MovementShapes_mapString = make(map[string]*MovementShape)
	for movementshape := range stage.MovementShapes {
		stage.MovementShapes_mapString[movementshape.Name] = movementshape
	}

	stage.Places_mapString = make(map[string]*Place)
	for place := range stage.Places {
		stage.Places_mapString[place.Name] = place
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
