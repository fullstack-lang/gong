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
	Arrows                map[*Arrow]struct{}
	Arrows_instance       map[*Arrow]*Arrow
	Arrows_mapString      map[string]*Arrow
	ArrowOrder            uint
	Arrow_stagedOrder     map[*Arrow]uint
	Arrow_orderStaged     map[uint]*Arrow
	Arrows_reference      map[*Arrow]*Arrow
	Arrows_referenceOrder map[*Arrow]uint

	// insertion point for slice of pointers maps
	OnAfterArrowCreateCallback GongOnAfterCreateInterface[Arrow]
	OnAfterArrowUpdateCallback GongOnAfterUpdateInterface[Arrow]
	OnAfterArrowDeleteCallback GongOnAfterDeleteInterface[Arrow]

	Bars                map[*Bar]struct{}
	Bars_instance       map[*Bar]*Bar
	Bars_mapString      map[string]*Bar
	BarOrder            uint
	Bar_stagedOrder     map[*Bar]uint
	Bar_orderStaged     map[uint]*Bar
	Bars_reference      map[*Bar]*Bar
	Bars_referenceOrder map[*Bar]uint

	// insertion point for slice of pointers maps
	OnAfterBarCreateCallback GongOnAfterCreateInterface[Bar]
	OnAfterBarUpdateCallback GongOnAfterUpdateInterface[Bar]
	OnAfterBarDeleteCallback GongOnAfterDeleteInterface[Bar]

	Gantts                map[*Gantt]struct{}
	Gantts_instance       map[*Gantt]*Gantt
	Gantts_mapString      map[string]*Gantt
	GanttOrder            uint
	Gantt_stagedOrder     map[*Gantt]uint
	Gantt_orderStaged     map[uint]*Gantt
	Gantts_reference      map[*Gantt]*Gantt
	Gantts_referenceOrder map[*Gantt]uint

	// insertion point for slice of pointers maps
	Gantt_Lanes_reverseMap map[*Lane]*Gantt

	Gantt_Milestones_reverseMap map[*Milestone]*Gantt

	Gantt_Groups_reverseMap map[*Group]*Gantt

	Gantt_Arrows_reverseMap map[*Arrow]*Gantt

	OnAfterGanttCreateCallback GongOnAfterCreateInterface[Gantt]
	OnAfterGanttUpdateCallback GongOnAfterUpdateInterface[Gantt]
	OnAfterGanttDeleteCallback GongOnAfterDeleteInterface[Gantt]

	Groups                map[*Group]struct{}
	Groups_instance       map[*Group]*Group
	Groups_mapString      map[string]*Group
	GroupOrder            uint
	Group_stagedOrder     map[*Group]uint
	Group_orderStaged     map[uint]*Group
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint

	// insertion point for slice of pointers maps
	Group_GroupLanes_reverseMap map[*Lane]*Group

	OnAfterGroupCreateCallback GongOnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback GongOnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback GongOnAfterDeleteInterface[Group]

	Lanes                map[*Lane]struct{}
	Lanes_instance       map[*Lane]*Lane
	Lanes_mapString      map[string]*Lane
	LaneOrder            uint
	Lane_stagedOrder     map[*Lane]uint
	Lane_orderStaged     map[uint]*Lane
	Lanes_reference      map[*Lane]*Lane
	Lanes_referenceOrder map[*Lane]uint

	// insertion point for slice of pointers maps
	Lane_Bars_reverseMap map[*Bar]*Lane

	OnAfterLaneCreateCallback GongOnAfterCreateInterface[Lane]
	OnAfterLaneUpdateCallback GongOnAfterUpdateInterface[Lane]
	OnAfterLaneDeleteCallback GongOnAfterDeleteInterface[Lane]

	LaneUses                map[*LaneUse]struct{}
	LaneUses_instance       map[*LaneUse]*LaneUse
	LaneUses_mapString      map[string]*LaneUse
	LaneUseOrder            uint
	LaneUse_stagedOrder     map[*LaneUse]uint
	LaneUse_orderStaged     map[uint]*LaneUse
	LaneUses_reference      map[*LaneUse]*LaneUse
	LaneUses_referenceOrder map[*LaneUse]uint

	// insertion point for slice of pointers maps
	OnAfterLaneUseCreateCallback GongOnAfterCreateInterface[LaneUse]
	OnAfterLaneUseUpdateCallback GongOnAfterUpdateInterface[LaneUse]
	OnAfterLaneUseDeleteCallback GongOnAfterDeleteInterface[LaneUse]

	Milestones                map[*Milestone]struct{}
	Milestones_instance       map[*Milestone]*Milestone
	Milestones_mapString      map[string]*Milestone
	MilestoneOrder            uint
	Milestone_stagedOrder     map[*Milestone]uint
	Milestone_orderStaged     map[uint]*Milestone
	Milestones_reference      map[*Milestone]*Milestone
	Milestones_referenceOrder map[*Milestone]uint

	// insertion point for slice of pointers maps
	Milestone_LanesToDisplay_reverseMap map[*Lane]*Milestone

	OnAfterMilestoneCreateCallback GongOnAfterCreateInterface[Milestone]
	OnAfterMilestoneUpdateCallback GongOnAfterUpdateInterface[Milestone]
	OnAfterMilestoneDeleteCallback GongOnAfterDeleteInterface[Milestone]

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
	__gong__clearReferences(&stage.Arrows_reference, &stage.Arrows_instance, &stage.Arrows_referenceOrder)

	__gong__clearReferences(&stage.Bars_reference, &stage.Bars_instance, &stage.Bars_referenceOrder)

	__gong__clearReferences(&stage.Gantts_reference, &stage.Gantts_instance, &stage.Gantts_referenceOrder)

	__gong__clearReferences(&stage.Groups_reference, &stage.Groups_instance, &stage.Groups_referenceOrder)

	__gong__clearReferences(&stage.Lanes_reference, &stage.Lanes_instance, &stage.Lanes_referenceOrder)

	__gong__clearReferences(&stage.LaneUses_reference, &stage.LaneUses_instance, &stage.LaneUses_referenceOrder)

	__gong__clearReferences(&stage.Milestones_reference, &stage.Milestones_instance, &stage.Milestones_referenceOrder)

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
	stage.ArrowOrder = __gong__recomputeOrder(stage.Arrow_stagedOrder)

	stage.BarOrder = __gong__recomputeOrder(stage.Bar_stagedOrder)

	stage.GanttOrder = __gong__recomputeOrder(stage.Gantt_stagedOrder)

	stage.GroupOrder = __gong__recomputeOrder(stage.Group_stagedOrder)

	stage.LaneOrder = __gong__recomputeOrder(stage.Lane_stagedOrder)

	stage.LaneUseOrder = __gong__recomputeOrder(stage.LaneUse_stagedOrder)

	stage.MilestoneOrder = __gong__recomputeOrder(stage.Milestone_stagedOrder)

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
	case *Arrow:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Arrows, stage.Arrow_stagedOrder))
	case *Bar:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Bars, stage.Bar_stagedOrder))
	case *Gantt:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Gantts, stage.Gantt_stagedOrder))
	case *Group:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder))
	case *Lane:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Lanes, stage.Lane_stagedOrder))
	case *LaneUse:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.LaneUses, stage.LaneUse_stagedOrder))
	case *Milestone:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Milestones, stage.Milestone_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/gantt/go/models"
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
		Arrows:           make(map[*Arrow]struct{}),
		Arrows_mapString: make(map[string]*Arrow),

		Bars:           make(map[*Bar]struct{}),
		Bars_mapString: make(map[string]*Bar),

		Gantts:           make(map[*Gantt]struct{}),
		Gantts_mapString: make(map[string]*Gantt),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		Lanes:           make(map[*Lane]struct{}),
		Lanes_mapString: make(map[string]*Lane),

		LaneUses:           make(map[*LaneUse]struct{}),
		LaneUses_mapString: make(map[string]*LaneUse),

		Milestones:           make(map[*Milestone]struct{}),
		Milestones_mapString: make(map[string]*Milestone),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Arrow_stagedOrder: make(map[*Arrow]uint),
		Arrow_orderStaged: make(map[uint]*Arrow),
		Arrows_reference:  make(map[*Arrow]*Arrow),

		Bar_stagedOrder: make(map[*Bar]uint),
		Bar_orderStaged: make(map[uint]*Bar),
		Bars_reference:  make(map[*Bar]*Bar),

		Gantt_stagedOrder: make(map[*Gantt]uint),
		Gantt_orderStaged: make(map[uint]*Gantt),
		Gantts_reference:  make(map[*Gantt]*Gantt),

		Group_stagedOrder: make(map[*Group]uint),
		Group_orderStaged: make(map[uint]*Group),
		Groups_reference:  make(map[*Group]*Group),

		Lane_stagedOrder: make(map[*Lane]uint),
		Lane_orderStaged: make(map[uint]*Lane),
		Lanes_reference:  make(map[*Lane]*Lane),

		LaneUse_stagedOrder: make(map[*LaneUse]uint),
		LaneUse_orderStaged: make(map[uint]*LaneUse),
		LaneUses_reference:  make(map[*LaneUse]*LaneUse),

		Milestone_stagedOrder: make(map[*Milestone]uint),
		Milestone_orderStaged: make(map[uint]*Milestone),
		Milestones_reference:  make(map[*Milestone]*Milestone),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Arrow": &ArrowUnmarshaller{},

			"Bar": &BarUnmarshaller{},

			"Gantt": &GanttUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"Lane": &LaneUnmarshaller{},

			"LaneUse": &LaneUseUnmarshaller{},

			"Milestone": &MilestoneUnmarshaller{},

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
	case *Arrow:
		return any(stage.Arrow_orderStaged[order]).(Type)
	case *Bar:
		return any(stage.Bar_orderStaged[order]).(Type)
	case *Gantt:
		return any(stage.Gantt_orderStaged[order]).(Type)
	case *Group:
		return any(stage.Group_orderStaged[order]).(Type)
	case *Lane:
		return any(stage.Lane_orderStaged[order]).(Type)
	case *LaneUse:
		return any(stage.LaneUse_orderStaged[order]).(Type)
	case *Milestone:
		return any(stage.Milestone_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Arrow"] = len(stage.Arrows)
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Gantt"] = len(stage.Gantts)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Lane"] = len(stage.Lanes)
	stage.Map_GongStructName_InstancesNb["LaneUse"] = len(stage.LaneUses)
	stage.Map_GongStructName_InstancesNb["Milestone"] = len(stage.Milestones)
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
// Stage puts arrow to the model stage
func (arrow *Arrow) Stage(stage *Stage) *Arrow {
	__gong__stage(stage.Arrows, stage.Arrow_stagedOrder, stage.Arrow_orderStaged, &stage.ArrowOrder, stage.Arrows_mapString, arrow, arrow.Name)
	return arrow
}

// StagePreserveOrder puts arrow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArrowOrder
// - update stage.ArrowOrder accordingly
func (arrow *Arrow) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Arrows, stage.Arrow_stagedOrder, stage.Arrow_orderStaged, &stage.ArrowOrder, stage.Arrows_mapString, arrow, order, arrow.Name)
}

// Unstage removes arrow off the model stage
func (arrow *Arrow) Unstage(stage *Stage) *Arrow {
	__gong__unstage(stage.Arrows, stage.Arrows_mapString, arrow, arrow.Name)
	return arrow
}

// UnstageVoid removes arrow off the model stage
func (arrow *Arrow) UnstageVoid(stage *Stage) {
	arrow.Unstage(stage)
}

func (arrow *Arrow) StageVoid(stage *Stage) {
	arrow.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) GetName() (res string) {
	return arrow.Name
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) SetName(name string) {
	arrow.Name = name
}

// Stage puts bar to the model stage
func (bar *Bar) Stage(stage *Stage) *Bar {
	__gong__stage(stage.Bars, stage.Bar_stagedOrder, stage.Bar_orderStaged, &stage.BarOrder, stage.Bars_mapString, bar, bar.Name)
	return bar
}

// StagePreserveOrder puts bar to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BarOrder
// - update stage.BarOrder accordingly
func (bar *Bar) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Bars, stage.Bar_stagedOrder, stage.Bar_orderStaged, &stage.BarOrder, stage.Bars_mapString, bar, order, bar.Name)
}

// Unstage removes bar off the model stage
func (bar *Bar) Unstage(stage *Stage) *Bar {
	__gong__unstage(stage.Bars, stage.Bars_mapString, bar, bar.Name)
	return bar
}

// UnstageVoid removes bar off the model stage
func (bar *Bar) UnstageVoid(stage *Stage) {
	bar.Unstage(stage)
}

func (bar *Bar) StageVoid(stage *Stage) {
	bar.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bar *Bar) GetName() (res string) {
	return bar.Name
}

// for satisfaction of GongStruct interface
func (bar *Bar) SetName(name string) {
	bar.Name = name
}

// Stage puts gantt to the model stage
func (gantt *Gantt) Stage(stage *Stage) *Gantt {
	__gong__stage(stage.Gantts, stage.Gantt_stagedOrder, stage.Gantt_orderStaged, &stage.GanttOrder, stage.Gantts_mapString, gantt, gantt.Name)
	return gantt
}

// StagePreserveOrder puts gantt to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GanttOrder
// - update stage.GanttOrder accordingly
func (gantt *Gantt) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Gantts, stage.Gantt_stagedOrder, stage.Gantt_orderStaged, &stage.GanttOrder, stage.Gantts_mapString, gantt, order, gantt.Name)
}

// Unstage removes gantt off the model stage
func (gantt *Gantt) Unstage(stage *Stage) *Gantt {
	__gong__unstage(stage.Gantts, stage.Gantts_mapString, gantt, gantt.Name)
	return gantt
}

// UnstageVoid removes gantt off the model stage
func (gantt *Gantt) UnstageVoid(stage *Stage) {
	gantt.Unstage(stage)
}

func (gantt *Gantt) StageVoid(stage *Stage) {
	gantt.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gantt *Gantt) GetName() (res string) {
	return gantt.Name
}

// for satisfaction of GongStruct interface
func (gantt *Gantt) SetName(name string) {
	gantt.Name = name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {
	__gong__stage(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, group.Name)
	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, order, group.Name)
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	__gong__unstage(stage.Groups, stage.Groups_mapString, group, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	group.Unstage(stage)
}

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
}

// for satisfaction of GongStruct interface
func (group *Group) SetName(name string) {
	group.Name = name
}

// Stage puts lane to the model stage
func (lane *Lane) Stage(stage *Stage) *Lane {
	__gong__stage(stage.Lanes, stage.Lane_stagedOrder, stage.Lane_orderStaged, &stage.LaneOrder, stage.Lanes_mapString, lane, lane.Name)
	return lane
}

// StagePreserveOrder puts lane to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LaneOrder
// - update stage.LaneOrder accordingly
func (lane *Lane) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Lanes, stage.Lane_stagedOrder, stage.Lane_orderStaged, &stage.LaneOrder, stage.Lanes_mapString, lane, order, lane.Name)
}

// Unstage removes lane off the model stage
func (lane *Lane) Unstage(stage *Stage) *Lane {
	__gong__unstage(stage.Lanes, stage.Lanes_mapString, lane, lane.Name)
	return lane
}

// UnstageVoid removes lane off the model stage
func (lane *Lane) UnstageVoid(stage *Stage) {
	lane.Unstage(stage)
}

func (lane *Lane) StageVoid(stage *Stage) {
	lane.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lane *Lane) GetName() (res string) {
	return lane.Name
}

// for satisfaction of GongStruct interface
func (lane *Lane) SetName(name string) {
	lane.Name = name
}

// Stage puts laneuse to the model stage
func (laneuse *LaneUse) Stage(stage *Stage) *LaneUse {
	__gong__stage(stage.LaneUses, stage.LaneUse_stagedOrder, stage.LaneUse_orderStaged, &stage.LaneUseOrder, stage.LaneUses_mapString, laneuse, laneuse.Name)
	return laneuse
}

// StagePreserveOrder puts laneuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LaneUseOrder
// - update stage.LaneUseOrder accordingly
func (laneuse *LaneUse) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LaneUses, stage.LaneUse_stagedOrder, stage.LaneUse_orderStaged, &stage.LaneUseOrder, stage.LaneUses_mapString, laneuse, order, laneuse.Name)
}

// Unstage removes laneuse off the model stage
func (laneuse *LaneUse) Unstage(stage *Stage) *LaneUse {
	__gong__unstage(stage.LaneUses, stage.LaneUses_mapString, laneuse, laneuse.Name)
	return laneuse
}

// UnstageVoid removes laneuse off the model stage
func (laneuse *LaneUse) UnstageVoid(stage *Stage) {
	laneuse.Unstage(stage)
}

func (laneuse *LaneUse) StageVoid(stage *Stage) {
	laneuse.Stage(stage)
}

// for satisfaction of GongStruct interface
func (laneuse *LaneUse) GetName() (res string) {
	return laneuse.Name
}

// for satisfaction of GongStruct interface
func (laneuse *LaneUse) SetName(name string) {
	laneuse.Name = name
}

// Stage puts milestone to the model stage
func (milestone *Milestone) Stage(stage *Stage) *Milestone {
	__gong__stage(stage.Milestones, stage.Milestone_stagedOrder, stage.Milestone_orderStaged, &stage.MilestoneOrder, stage.Milestones_mapString, milestone, milestone.Name)
	return milestone
}

// StagePreserveOrder puts milestone to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MilestoneOrder
// - update stage.MilestoneOrder accordingly
func (milestone *Milestone) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Milestones, stage.Milestone_stagedOrder, stage.Milestone_orderStaged, &stage.MilestoneOrder, stage.Milestones_mapString, milestone, order, milestone.Name)
}

// Unstage removes milestone off the model stage
func (milestone *Milestone) Unstage(stage *Stage) *Milestone {
	__gong__unstage(stage.Milestones, stage.Milestones_mapString, milestone, milestone.Name)
	return milestone
}

// UnstageVoid removes milestone off the model stage
func (milestone *Milestone) UnstageVoid(stage *Stage) {
	milestone.Unstage(stage)
}

func (milestone *Milestone) StageVoid(stage *Stage) {
	milestone.Stage(stage)
}

// for satisfaction of GongStruct interface
func (milestone *Milestone) GetName() (res string) {
	return milestone.Name
}

// for satisfaction of GongStruct interface
func (milestone *Milestone) SetName(name string) {
	milestone.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Arrows, &stage.Arrows_mapString, &stage.Arrow_stagedOrder, &stage.ArrowOrder)

	__gong__resetStageType(&stage.Bars, &stage.Bars_mapString, &stage.Bar_stagedOrder, &stage.BarOrder)

	__gong__resetStageType(&stage.Gantts, &stage.Gantts_mapString, &stage.Gantt_stagedOrder, &stage.GanttOrder)

	__gong__resetStageType(&stage.Groups, &stage.Groups_mapString, &stage.Group_stagedOrder, &stage.GroupOrder)

	__gong__resetStageType(&stage.Lanes, &stage.Lanes_mapString, &stage.Lane_stagedOrder, &stage.LaneOrder)

	__gong__resetStageType(&stage.LaneUses, &stage.LaneUses_mapString, &stage.LaneUse_stagedOrder, &stage.LaneUseOrder)

	__gong__resetStageType(&stage.Milestones, &stage.Milestones_mapString, &stage.Milestone_stagedOrder, &stage.MilestoneOrder)

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
	case *Arrow:
		return any(stage.Arrows_mapString).(map[string]Type)
	case *Bar:
		return any(stage.Bars_mapString).(map[string]Type)
	case *Gantt:
		return any(stage.Gantts_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *Lane:
		return any(stage.Lanes_mapString).(map[string]Type)
	case *LaneUse:
		return any(stage.LaneUses_mapString).(map[string]Type)
	case *Milestone:
		return any(stage.Milestones_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Arrow:
		return any(&stage.Arrows).(*map[Type]struct{})
	case *Bar:
		return any(&stage.Bars).(*map[Type]struct{})
	case *Gantt:
		return any(&stage.Gantts).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *Lane:
		return any(&stage.Lanes).(*map[Type]struct{})
	case *LaneUse:
		return any(&stage.LaneUses).(*map[Type]struct{})
	case *Milestone:
		return any(&stage.Milestones).(*map[Type]struct{})
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
	case Arrow:
		return any(&Arrow{
			From: &Bar{Name: "From"},
			To: &Bar{Name: "To"},
		}).(*Type)
	case Gantt:
		return any(&Gantt{
			Lanes: []*Lane{{Name: "Lanes"}},
			Milestones: []*Milestone{{Name: "Milestones"}},
			Groups: []*Group{{Name: "Groups"}},
			Arrows: []*Arrow{{Name: "Arrows"}},
		}).(*Type)
	case Group:
		return any(&Group{
			GroupLanes: []*Lane{{Name: "GroupLanes"}},
		}).(*Type)
	case Lane:
		return any(&Lane{
			Bars: []*Bar{{Name: "Bars"}},
		}).(*Type)
	case LaneUse:
		return any(&LaneUse{
			Lane: &Lane{Name: "Lane"},
		}).(*Type)
	case Milestone:
		return any(&Milestone{
			LanesToDisplay: []*Lane{{Name: "LanesToDisplay"}},
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
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		case "From":
			res := make(map[*Bar][]*Arrow)
			for arrow := range stage.Arrows {
				if arrow.From != nil {
					bar_ := arrow.From
					var arrows []*Arrow
					_, ok := res[bar_]
					if ok {
						arrows = res[bar_]
					} else {
						arrows = make([]*Arrow, 0)
					}
					arrows = append(arrows, arrow)
					res[bar_] = arrows
				}
			}
			return any(res).(map[*End][]*Start)
		case "To":
			res := make(map[*Bar][]*Arrow)
			for arrow := range stage.Arrows {
				if arrow.To != nil {
					bar_ := arrow.To
					var arrows []*Arrow
					_, ok := res[bar_]
					if ok {
						arrows = res[bar_]
					} else {
						arrows = make([]*Arrow, 0)
					}
					arrows = append(arrows, arrow)
					res[bar_] = arrows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gantt
	case Gantt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lane
	case Lane:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LaneUse
	case LaneUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Lane":
			res := make(map[*Lane][]*LaneUse)
			for laneuse := range stage.LaneUses {
				if laneuse.Lane != nil {
					lane_ := laneuse.Lane
					var laneuses []*LaneUse
					_, ok := res[lane_]
					if ok {
						laneuses = res[lane_]
					} else {
						laneuses = make([]*LaneUse, 0)
					}
					laneuses = append(laneuses, laneuse)
					res[lane_] = laneuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Milestone
	case Milestone:
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
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gantt
	case Gantt:
		switch fieldname {
		// insertion point for per direct association field
		case "Lanes":
			res := make(map[*Lane][]*Gantt)
			for gantt := range stage.Gantts {
				for _, lane_ := range gantt.Lanes {
					res[lane_] = append(res[lane_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Milestones":
			res := make(map[*Milestone][]*Gantt)
			for gantt := range stage.Gantts {
				for _, milestone_ := range gantt.Milestones {
					res[milestone_] = append(res[milestone_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Gantt)
			for gantt := range stage.Gantts {
				for _, group_ := range gantt.Groups {
					res[group_] = append(res[group_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arrows":
			res := make(map[*Arrow][]*Gantt)
			for gantt := range stage.Gantts {
				for _, arrow_ := range gantt.Arrows {
					res[arrow_] = append(res[arrow_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupLanes":
			res := make(map[*Lane][]*Group)
			for group := range stage.Groups {
				for _, lane_ := range group.GroupLanes {
					res[lane_] = append(res[lane_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Lane
	case Lane:
		switch fieldname {
		// insertion point for per direct association field
		case "Bars":
			res := make(map[*Bar][]*Lane)
			for lane := range stage.Lanes {
				for _, bar_ := range lane.Bars {
					res[bar_] = append(res[bar_], lane)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LaneUse
	case LaneUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Milestone
	case Milestone:
		switch fieldname {
		// insertion point for per direct association field
		case "LanesToDisplay":
			res := make(map[*Lane][]*Milestone)
			for milestone := range stage.Milestones {
				for _, lane_ := range milestone.LanesToDisplay {
					res[lane_] = append(res[lane_], milestone)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *Arrow:
		res = any(new(Arrow)).(Type)
	case *Bar:
		res = any(new(Bar)).(Type)
	case *Gantt:
		res = any(new(Gantt)).(Type)
	case *Group:
		res = any(new(Group)).(Type)
	case *Lane:
		res = any(new(Lane)).(Type)
	case *LaneUse:
		res = any(new(LaneUse)).(Type)
	case *Milestone:
		res = any(new(Milestone)).(Type)
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
	case *Arrow:
		res = "Arrow"
	case *Bar:
		res = "Bar"
	case *Gantt:
		res = "Gantt"
	case *Group:
		res = "Group"
	case *Lane:
		res = "Lane"
	case *LaneUse:
		res = "LaneUse"
	case *Milestone:
		res = "Milestone"
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
	case *Arrow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Arrows"
		res = append(res, rf)
	case *Bar:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Lane"
		rf.Fieldname = "Bars"
		res = append(res, rf)
	case *Gantt:
		var rf ReverseField
		_ = rf
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *Lane:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Lanes"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "GroupLanes"
		res = append(res, rf)
		rf.GongstructName = "Milestone"
		rf.Fieldname = "LanesToDisplay"
		res = append(res, rf)
	case *LaneUse:
		var rf ReverseField
		_ = rf
	case *Milestone:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Milestones"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (arrow *Arrow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "From",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bar",
		},
		{
			Name:                 "To",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bar",
		},
		{
			Name:               "OptionnalColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OptionnalStroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (bar *Bar) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "End",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ComputedDuration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
		{
			Name:               "OptionnalColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OptionnalStroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (gantt *Gantt) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedStart",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ComputedEnd",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ComputedDuration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
		{
			Name:               "UseManualStartAndEndDates",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ManualStart",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ManualEnd",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "LaneHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RatioBarToLaneHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "YTopMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XLeftText",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TextHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XLeftLanes",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XRightMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ArrowLengthToTheRightOfStartBar",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ArrowTipLenght",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TimeLine_Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeLine_FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TimeLine_Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeLine_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Group_Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Group_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Group_StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DateYOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "AlignOnStartEndOnYearStart",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Lanes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
		},
		{
			Name:                 "Milestones",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Milestone",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Arrows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Arrow",
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "GroupLanes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
		},
	}
	return
}

func (lane *Lane) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Bars",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Bar",
		},
	}
	return
}

func (laneuse *LaneUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Lane",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Lane",
		},
	}
	return
}

func (milestone *Milestone) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Date",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "DisplayVerticalBar",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LanesToDisplay",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
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
func (arrow *Arrow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arrow.Name
	case "From":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arrow.From != nil {
			res.valueString = arrow.From.Name
			res.ids = arrow.From.GongGetUUID(stage)
		}
	case "To":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arrow.To != nil {
			res.valueString = arrow.To.Name
			res.ids = arrow.To.GongGetUUID(stage)
		}
	case "OptionnalColor":
		res.valueString = arrow.OptionnalColor
	case "OptionnalStroke":
		res.valueString = arrow.OptionnalStroke
	}
	return
}

func (bar *Bar) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bar.Name
	case "Start":
		res.valueString = bar.Start.String()
	case "End":
		res.valueString = bar.End.String()
	case "ComputedDuration":
		if math.Abs(bar.ComputedDuration.Hours()) >= 24 {
			days := __gong__abs(int(int(bar.ComputedDuration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(bar.ComputedDuration.Hours()) % 24
			remainingMinutes := int(bar.ComputedDuration.Minutes()) % 60
			remainingSeconds := int(bar.ComputedDuration.Seconds()) % 60

			if bar.ComputedDuration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", bar.ComputedDuration.String())
		}
	case "OptionnalColor":
		res.valueString = bar.OptionnalColor
	case "OptionnalStroke":
		res.valueString = bar.OptionnalStroke
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", bar.FillOpacity)
		res.valueFloat = bar.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", bar.StrokeWidth)
		res.valueFloat = bar.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = bar.StrokeDashArray
	}
	return
}

func (gantt *Gantt) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gantt.Name
	case "ComputedStart":
		res.valueString = gantt.ComputedStart.String()
	case "ComputedEnd":
		res.valueString = gantt.ComputedEnd.String()
	case "ComputedDuration":
		if math.Abs(gantt.ComputedDuration.Hours()) >= 24 {
			days := __gong__abs(int(int(gantt.ComputedDuration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(gantt.ComputedDuration.Hours()) % 24
			remainingMinutes := int(gantt.ComputedDuration.Minutes()) % 60
			remainingSeconds := int(gantt.ComputedDuration.Seconds()) % 60

			if gantt.ComputedDuration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", gantt.ComputedDuration.String())
		}
	case "UseManualStartAndEndDates":
		res.valueString = fmt.Sprintf("%t", gantt.UseManualStartAndEndDates)
		res.valueBool = gantt.UseManualStartAndEndDates
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ManualStart":
		res.valueString = gantt.ManualStart.String()
	case "ManualEnd":
		res.valueString = gantt.ManualEnd.String()
	case "LaneHeight":
		res.valueString = fmt.Sprintf("%f", gantt.LaneHeight)
		res.valueFloat = gantt.LaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RatioBarToLaneHeight":
		res.valueString = fmt.Sprintf("%f", gantt.RatioBarToLaneHeight)
		res.valueFloat = gantt.RatioBarToLaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "YTopMargin":
		res.valueString = fmt.Sprintf("%f", gantt.YTopMargin)
		res.valueFloat = gantt.YTopMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftText":
		res.valueString = fmt.Sprintf("%f", gantt.XLeftText)
		res.valueFloat = gantt.XLeftText
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TextHeight":
		res.valueString = fmt.Sprintf("%f", gantt.TextHeight)
		res.valueFloat = gantt.TextHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftLanes":
		res.valueString = fmt.Sprintf("%f", gantt.XLeftLanes)
		res.valueFloat = gantt.XLeftLanes
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XRightMargin":
		res.valueString = fmt.Sprintf("%f", gantt.XRightMargin)
		res.valueFloat = gantt.XRightMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowLengthToTheRightOfStartBar":
		res.valueString = fmt.Sprintf("%f", gantt.ArrowLengthToTheRightOfStartBar)
		res.valueFloat = gantt.ArrowLengthToTheRightOfStartBar
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowTipLenght":
		res.valueString = fmt.Sprintf("%f", gantt.ArrowTipLenght)
		res.valueFloat = gantt.ArrowTipLenght
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Color":
		res.valueString = gantt.TimeLine_Color
	case "TimeLine_FillOpacity":
		res.valueString = fmt.Sprintf("%f", gantt.TimeLine_FillOpacity)
		res.valueFloat = gantt.TimeLine_FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Stroke":
		res.valueString = gantt.TimeLine_Stroke
	case "TimeLine_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", gantt.TimeLine_StrokeWidth)
		res.valueFloat = gantt.TimeLine_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Group_Stroke":
		res.valueString = gantt.Group_Stroke
	case "Group_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", gantt.Group_StrokeWidth)
		res.valueFloat = gantt.Group_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Group_StrokeDashArray":
		res.valueString = gantt.Group_StrokeDashArray
	case "DateYOffset":
		res.valueString = fmt.Sprintf("%f", gantt.DateYOffset)
		res.valueFloat = gantt.DateYOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AlignOnStartEndOnYearStart":
		res.valueString = fmt.Sprintf("%t", gantt.AlignOnStartEndOnYearStart)
		res.valueBool = gantt.AlignOnStartEndOnYearStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Lanes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Lanes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Milestones":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Milestones {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Arrows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Arrows {
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

func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "GroupLanes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.GroupLanes {
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

func (lane *Lane) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lane.Name
	case "Order":
		res.valueString = fmt.Sprintf("%d", lane.Order)
		res.valueInt = lane.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Bars":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lane.Bars {
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

func (laneuse *LaneUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = laneuse.Name
	case "Lane":
		res.GongFieldValueType = GongFieldValueTypePointer
		if laneuse.Lane != nil {
			res.valueString = laneuse.Lane.Name
			res.ids = laneuse.Lane.GongGetUUID(stage)
		}
	}
	return
}

func (milestone *Milestone) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = milestone.Name
	case "Date":
		res.valueString = milestone.Date.String()
	case "DisplayVerticalBar":
		res.valueString = fmt.Sprintf("%t", milestone.DisplayVerticalBar)
		res.valueBool = milestone.DisplayVerticalBar
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LanesToDisplay":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range milestone.LanesToDisplay {
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (arrow *Arrow) GongGetGongstructName() string {
	return "Arrow"
}

func (bar *Bar) GongGetGongstructName() string {
	return "Bar"
}

func (gantt *Gantt) GongGetGongstructName() string {
	return "Gantt"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (lane *Lane) GongGetGongstructName() string {
	return "Lane"
}

func (laneuse *LaneUse) GongGetGongstructName() string {
	return "LaneUse"
}

func (milestone *Milestone) GongGetGongstructName() string {
	return "Milestone"
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
	__gong__rebuildMapString(stage.Arrows, &stage.Arrows_mapString)

	__gong__rebuildMapString(stage.Bars, &stage.Bars_mapString)

	__gong__rebuildMapString(stage.Gantts, &stage.Gantts_mapString)

	__gong__rebuildMapString(stage.Groups, &stage.Groups_mapString)

	__gong__rebuildMapString(stage.Lanes, &stage.Lanes_mapString)

	__gong__rebuildMapString(stage.LaneUses, &stage.LaneUses_mapString)

	__gong__rebuildMapString(stage.Milestones, &stage.Milestones_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
