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
	Buttons                map[*Button]struct{}
	Buttons_instance       map[*Button]*Button
	Buttons_mapString      map[string]*Button
	ButtonOrder            uint
	Button_stagedOrder     map[*Button]uint
	Button_orderStaged     map[uint]*Button
	Buttons_reference      map[*Button]*Button
	Buttons_referenceOrder map[*Button]uint

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback GongOnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback GongOnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback GongOnAfterDeleteInterface[Button]

	ButtonToggles                map[*ButtonToggle]struct{}
	ButtonToggles_instance       map[*ButtonToggle]*ButtonToggle
	ButtonToggles_mapString      map[string]*ButtonToggle
	ButtonToggleOrder            uint
	ButtonToggle_stagedOrder     map[*ButtonToggle]uint
	ButtonToggle_orderStaged     map[uint]*ButtonToggle
	ButtonToggles_reference      map[*ButtonToggle]*ButtonToggle
	ButtonToggles_referenceOrder map[*ButtonToggle]uint

	// insertion point for slice of pointers maps
	OnAfterButtonToggleCreateCallback GongOnAfterCreateInterface[ButtonToggle]
	OnAfterButtonToggleUpdateCallback GongOnAfterUpdateInterface[ButtonToggle]
	OnAfterButtonToggleDeleteCallback GongOnAfterDeleteInterface[ButtonToggle]

	Groups                map[*Group]struct{}
	Groups_instance       map[*Group]*Group
	Groups_mapString      map[string]*Group
	GroupOrder            uint
	Group_stagedOrder     map[*Group]uint
	Group_orderStaged     map[uint]*Group
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint

	// insertion point for slice of pointers maps
	Group_Buttons_reverseMap map[*Button]*Group

	OnAfterGroupCreateCallback GongOnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback GongOnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback GongOnAfterDeleteInterface[Group]

	GroupToogles                map[*GroupToogle]struct{}
	GroupToogles_instance       map[*GroupToogle]*GroupToogle
	GroupToogles_mapString      map[string]*GroupToogle
	GroupToogleOrder            uint
	GroupToogle_stagedOrder     map[*GroupToogle]uint
	GroupToogle_orderStaged     map[uint]*GroupToogle
	GroupToogles_reference      map[*GroupToogle]*GroupToogle
	GroupToogles_referenceOrder map[*GroupToogle]uint

	// insertion point for slice of pointers maps
	GroupToogle_ButtonToggles_reverseMap map[*ButtonToggle]*GroupToogle

	OnAfterGroupToogleCreateCallback GongOnAfterCreateInterface[GroupToogle]
	OnAfterGroupToogleUpdateCallback GongOnAfterUpdateInterface[GroupToogle]
	OnAfterGroupToogleDeleteCallback GongOnAfterDeleteInterface[GroupToogle]

	Layouts                map[*Layout]struct{}
	Layouts_instance       map[*Layout]*Layout
	Layouts_mapString      map[string]*Layout
	LayoutOrder            uint
	Layout_stagedOrder     map[*Layout]uint
	Layout_orderStaged     map[uint]*Layout
	Layouts_reference      map[*Layout]*Layout
	Layouts_referenceOrder map[*Layout]uint

	// insertion point for slice of pointers maps
	Layout_Groups_reverseMap map[*Group]*Layout

	Layout_GroupToogles_reverseMap map[*GroupToogle]*Layout

	OnAfterLayoutCreateCallback GongOnAfterCreateInterface[Layout]
	OnAfterLayoutUpdateCallback GongOnAfterUpdateInterface[Layout]
	OnAfterLayoutDeleteCallback GongOnAfterDeleteInterface[Layout]

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
	__gong__clearReferences(&stage.Buttons_reference, &stage.Buttons_instance, &stage.Buttons_referenceOrder)

	__gong__clearReferences(&stage.ButtonToggles_reference, &stage.ButtonToggles_instance, &stage.ButtonToggles_referenceOrder)

	__gong__clearReferences(&stage.Groups_reference, &stage.Groups_instance, &stage.Groups_referenceOrder)

	__gong__clearReferences(&stage.GroupToogles_reference, &stage.GroupToogles_instance, &stage.GroupToogles_referenceOrder)

	__gong__clearReferences(&stage.Layouts_reference, &stage.Layouts_instance, &stage.Layouts_referenceOrder)

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
	stage.ButtonOrder = __gong__recomputeOrder(stage.Button_stagedOrder)

	stage.ButtonToggleOrder = __gong__recomputeOrder(stage.ButtonToggle_stagedOrder)

	stage.GroupOrder = __gong__recomputeOrder(stage.Group_stagedOrder)

	stage.GroupToogleOrder = __gong__recomputeOrder(stage.GroupToogle_stagedOrder)

	stage.LayoutOrder = __gong__recomputeOrder(stage.Layout_stagedOrder)

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
	case *Button:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Buttons, stage.Button_stagedOrder))
	case *ButtonToggle:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ButtonToggles, stage.ButtonToggle_stagedOrder))
	case *Group:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder))
	case *GroupToogle:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.GroupToogles, stage.GroupToogle_stagedOrder))
	case *Layout:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Layouts, stage.Layout_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/button/go/models"
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
		Buttons:           make(map[*Button]struct{}),
		Buttons_mapString: make(map[string]*Button),

		ButtonToggles:           make(map[*ButtonToggle]struct{}),
		ButtonToggles_mapString: make(map[string]*ButtonToggle),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		GroupToogles:           make(map[*GroupToogle]struct{}),
		GroupToogles_mapString: make(map[string]*GroupToogle),

		Layouts:           make(map[*Layout]struct{}),
		Layouts_mapString: make(map[string]*Layout),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Button_stagedOrder: make(map[*Button]uint),
		Button_orderStaged: make(map[uint]*Button),
		Buttons_reference:  make(map[*Button]*Button),

		ButtonToggle_stagedOrder: make(map[*ButtonToggle]uint),
		ButtonToggle_orderStaged: make(map[uint]*ButtonToggle),
		ButtonToggles_reference:  make(map[*ButtonToggle]*ButtonToggle),

		Group_stagedOrder: make(map[*Group]uint),
		Group_orderStaged: make(map[uint]*Group),
		Groups_reference:  make(map[*Group]*Group),

		GroupToogle_stagedOrder: make(map[*GroupToogle]uint),
		GroupToogle_orderStaged: make(map[uint]*GroupToogle),
		GroupToogles_reference:  make(map[*GroupToogle]*GroupToogle),

		Layout_stagedOrder: make(map[*Layout]uint),
		Layout_orderStaged: make(map[uint]*Layout),
		Layouts_reference:  make(map[*Layout]*Layout),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Button": &ButtonUnmarshaller{},

			"ButtonToggle": &ButtonToggleUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"GroupToogle": &GroupToogleUnmarshaller{},

			"Layout": &LayoutUnmarshaller{},

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
	case *Button:
		return any(stage.Button_orderStaged[order]).(Type)
	case *ButtonToggle:
		return any(stage.ButtonToggle_orderStaged[order]).(Type)
	case *Group:
		return any(stage.Group_orderStaged[order]).(Type)
	case *GroupToogle:
		return any(stage.GroupToogle_orderStaged[order]).(Type)
	case *Layout:
		return any(stage.Layout_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["ButtonToggle"] = len(stage.ButtonToggles)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["GroupToogle"] = len(stage.GroupToogles)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)
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
// Stage puts button to the model stage
func (button *Button) Stage(stage *Stage) *Button {
	__gong__stage(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, button.Name)
	return button
}

// StagePreserveOrder puts button to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonOrder
// - update stage.ButtonOrder accordingly
func (button *Button) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, order, button.Name)
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	__gong__unstage(stage.Buttons, stage.Buttons_mapString, button, button.Name)
	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	button.Unstage(stage)
}

func (button *Button) StageVoid(stage *Stage) {
	button.Stage(stage)
}

// for satisfaction of GongStruct interface
func (button *Button) GetName() (res string) {
	return button.Name
}

// for satisfaction of GongStruct interface
func (button *Button) SetName(name string) {
	button.Name = name
}

// Stage puts buttontoggle to the model stage
func (buttontoggle *ButtonToggle) Stage(stage *Stage) *ButtonToggle {
	__gong__stage(stage.ButtonToggles, stage.ButtonToggle_stagedOrder, stage.ButtonToggle_orderStaged, &stage.ButtonToggleOrder, stage.ButtonToggles_mapString, buttontoggle, buttontoggle.Name)
	return buttontoggle
}

// StagePreserveOrder puts buttontoggle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonToggleOrder
// - update stage.ButtonToggleOrder accordingly
func (buttontoggle *ButtonToggle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ButtonToggles, stage.ButtonToggle_stagedOrder, stage.ButtonToggle_orderStaged, &stage.ButtonToggleOrder, stage.ButtonToggles_mapString, buttontoggle, order, buttontoggle.Name)
}

// Unstage removes buttontoggle off the model stage
func (buttontoggle *ButtonToggle) Unstage(stage *Stage) *ButtonToggle {
	__gong__unstage(stage.ButtonToggles, stage.ButtonToggles_mapString, buttontoggle, buttontoggle.Name)
	return buttontoggle
}

// UnstageVoid removes buttontoggle off the model stage
func (buttontoggle *ButtonToggle) UnstageVoid(stage *Stage) {
	buttontoggle.Unstage(stage)
}

func (buttontoggle *ButtonToggle) StageVoid(stage *Stage) {
	buttontoggle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (buttontoggle *ButtonToggle) GetName() (res string) {
	return buttontoggle.Name
}

// for satisfaction of GongStruct interface
func (buttontoggle *ButtonToggle) SetName(name string) {
	buttontoggle.Name = name
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

// Stage puts grouptoogle to the model stage
func (grouptoogle *GroupToogle) Stage(stage *Stage) *GroupToogle {
	__gong__stage(stage.GroupToogles, stage.GroupToogle_stagedOrder, stage.GroupToogle_orderStaged, &stage.GroupToogleOrder, stage.GroupToogles_mapString, grouptoogle, grouptoogle.Name)
	return grouptoogle
}

// StagePreserveOrder puts grouptoogle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupToogleOrder
// - update stage.GroupToogleOrder accordingly
func (grouptoogle *GroupToogle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroupToogles, stage.GroupToogle_stagedOrder, stage.GroupToogle_orderStaged, &stage.GroupToogleOrder, stage.GroupToogles_mapString, grouptoogle, order, grouptoogle.Name)
}

// Unstage removes grouptoogle off the model stage
func (grouptoogle *GroupToogle) Unstage(stage *Stage) *GroupToogle {
	__gong__unstage(stage.GroupToogles, stage.GroupToogles_mapString, grouptoogle, grouptoogle.Name)
	return grouptoogle
}

// UnstageVoid removes grouptoogle off the model stage
func (grouptoogle *GroupToogle) UnstageVoid(stage *Stage) {
	grouptoogle.Unstage(stage)
}

func (grouptoogle *GroupToogle) StageVoid(stage *Stage) {
	grouptoogle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (grouptoogle *GroupToogle) GetName() (res string) {
	return grouptoogle.Name
}

// for satisfaction of GongStruct interface
func (grouptoogle *GroupToogle) SetName(name string) {
	grouptoogle.Name = name
}

// Stage puts layout to the model stage
func (layout *Layout) Stage(stage *Stage) *Layout {
	__gong__stage(stage.Layouts, stage.Layout_stagedOrder, stage.Layout_orderStaged, &stage.LayoutOrder, stage.Layouts_mapString, layout, layout.Name)
	return layout
}

// StagePreserveOrder puts layout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LayoutOrder
// - update stage.LayoutOrder accordingly
func (layout *Layout) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Layouts, stage.Layout_stagedOrder, stage.Layout_orderStaged, &stage.LayoutOrder, stage.Layouts_mapString, layout, order, layout.Name)
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *Stage) *Layout {
	__gong__unstage(stage.Layouts, stage.Layouts_mapString, layout, layout.Name)
	return layout
}

// UnstageVoid removes layout off the model stage
func (layout *Layout) UnstageVoid(stage *Stage) {
	layout.Unstage(stage)
}

func (layout *Layout) StageVoid(stage *Stage) {
	layout.Stage(stage)
}

// for satisfaction of GongStruct interface
func (layout *Layout) GetName() (res string) {
	return layout.Name
}

// for satisfaction of GongStruct interface
func (layout *Layout) SetName(name string) {
	layout.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Buttons, &stage.Buttons_mapString, &stage.Button_stagedOrder, &stage.ButtonOrder)

	__gong__resetStageType(&stage.ButtonToggles, &stage.ButtonToggles_mapString, &stage.ButtonToggle_stagedOrder, &stage.ButtonToggleOrder)

	__gong__resetStageType(&stage.Groups, &stage.Groups_mapString, &stage.Group_stagedOrder, &stage.GroupOrder)

	__gong__resetStageType(&stage.GroupToogles, &stage.GroupToogles_mapString, &stage.GroupToogle_stagedOrder, &stage.GroupToogleOrder)

	__gong__resetStageType(&stage.Layouts, &stage.Layouts_mapString, &stage.Layout_stagedOrder, &stage.LayoutOrder)

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
	case *Button:
		return any(stage.Buttons_mapString).(map[string]Type)
	case *ButtonToggle:
		return any(stage.ButtonToggles_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *GroupToogle:
		return any(stage.GroupToogles_mapString).(map[string]Type)
	case *Layout:
		return any(stage.Layouts_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Button:
		return any(&stage.Buttons).(*map[Type]struct{})
	case *ButtonToggle:
		return any(&stage.ButtonToggles).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *GroupToogle:
		return any(&stage.GroupToogles).(*map[Type]struct{})
	case *Layout:
		return any(&stage.Layouts).(*map[Type]struct{})
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
	case Group:
		return any(&Group{
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case GroupToogle:
		return any(&GroupToogle{
			ButtonToggles: []*ButtonToggle{{Name: "ButtonToggles"}},
		}).(*Type)
	case Layout:
		return any(&Layout{
			Groups: []*Group{{Name: "Groups"}},
			GroupToogles: []*GroupToogle{{Name: "GroupToogles"}},
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
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ButtonToggle
	case ButtonToggle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupToogle
	case GroupToogle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layout
	case Layout:
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
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ButtonToggle
	case ButtonToggle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Buttons":
			res := make(map[*Button][]*Group)
			for group := range stage.Groups {
				for _, button_ := range group.Buttons {
					res[button_] = append(res[button_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GroupToogle
	case GroupToogle:
		switch fieldname {
		// insertion point for per direct association field
		case "ButtonToggles":
			res := make(map[*ButtonToggle][]*GroupToogle)
			for grouptoogle := range stage.GroupToogles {
				for _, buttontoggle_ := range grouptoogle.ButtonToggles {
					res[buttontoggle_] = append(res[buttontoggle_], grouptoogle)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Layout
	case Layout:
		switch fieldname {
		// insertion point for per direct association field
		case "Groups":
			res := make(map[*Group][]*Layout)
			for layout := range stage.Layouts {
				for _, group_ := range layout.Groups {
					res[group_] = append(res[group_], layout)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GroupToogles":
			res := make(map[*GroupToogle][]*Layout)
			for layout := range stage.Layouts {
				for _, grouptoogle_ := range layout.GroupToogles {
					res[grouptoogle_] = append(res[grouptoogle_], layout)
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
	case *Button:
		res = any(new(Button)).(Type)
	case *ButtonToggle:
		res = any(new(ButtonToggle)).(Type)
	case *Group:
		res = any(new(Group)).(Type)
	case *GroupToogle:
		res = any(new(GroupToogle)).(Type)
	case *Layout:
		res = any(new(Layout)).(Type)
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
	case *Button:
		res = "Button"
	case *ButtonToggle:
		res = "ButtonToggle"
	case *Group:
		res = "Group"
	case *GroupToogle:
		res = "GroupToogle"
	case *Layout:
		res = "Layout"
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
	case *Button:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
	case *ButtonToggle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GroupToogle"
		rf.Fieldname = "ButtonToggles"
		res = append(res, rf)
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *GroupToogle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "GroupToogles"
		res = append(res, rf)
	case *Layout:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (button *Button) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Color",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MatButtonPaletteType",
		},
		{
			Name:                 "MatButtonType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MatButtonType",
		},
		{
			Name:                 "MatButtonAppearance",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "MatButtonAppearance",
		},
		{
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ToolTipPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ToolTipPositionEnum",
		},
	}
	return
}

func (buttontoggle *ButtonToggle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "Percentage",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "Buttons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Button",
		},
		{
			Name:               "NbColumns",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (grouptoogle *GroupToogle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Percentage",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "ButtonToggles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ButtonToggle",
		},
		{
			Name:               "IsSingleSelector",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (layout *Layout) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "GroupToogles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroupToogle",
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
func (button *Button) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = button.Name
	case "Label":
		res.valueString = button.Label
	case "Icon":
		res.valueString = button.Icon
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", button.IsDisabled)
		res.valueBool = button.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Color":
		enum := button.Color
		res.valueString = enum.ToCodeString()
	case "MatButtonType":
		enum := button.MatButtonType
		res.valueString = enum.ToCodeString()
	case "MatButtonAppearance":
		enum := button.MatButtonAppearance
		res.valueString = enum.ToCodeString()
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", button.HasToolTip)
		res.valueBool = button.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = button.ToolTipText
	case "ToolTipPosition":
		enum := button.ToolTipPosition
		res.valueString = enum.ToCodeString()
	}
	return
}

func (buttontoggle *ButtonToggle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = buttontoggle.Name
	case "Label":
		res.valueString = buttontoggle.Label
	case "Icon":
		res.valueString = buttontoggle.Icon
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", buttontoggle.IsDisabled)
		res.valueBool = buttontoggle.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", buttontoggle.IsChecked)
		res.valueBool = buttontoggle.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "Percentage":
		res.valueString = fmt.Sprintf("%f", group.Percentage)
		res.valueFloat = group.Percentage
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Buttons":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Buttons {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbColumns":
		res.valueString = fmt.Sprintf("%d", group.NbColumns)
		res.valueInt = group.NbColumns
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (grouptoogle *GroupToogle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = grouptoogle.Name
	case "Percentage":
		res.valueString = fmt.Sprintf("%f", grouptoogle.Percentage)
		res.valueFloat = grouptoogle.Percentage
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ButtonToggles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range grouptoogle.ButtonToggles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSingleSelector":
		res.valueString = fmt.Sprintf("%t", grouptoogle.IsSingleSelector)
		res.valueBool = grouptoogle.IsSingleSelector
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (layout *Layout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = layout.Name
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layout.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GroupToogles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layout.GroupToogles {
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
func (button *Button) GongGetGongstructName() string {
	return "Button"
}

func (buttontoggle *ButtonToggle) GongGetGongstructName() string {
	return "ButtonToggle"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (grouptoogle *GroupToogle) GongGetGongstructName() string {
	return "GroupToogle"
}

func (layout *Layout) GongGetGongstructName() string {
	return "Layout"
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
	__gong__rebuildMapString(stage.Buttons, &stage.Buttons_mapString)

	__gong__rebuildMapString(stage.ButtonToggles, &stage.ButtonToggles_mapString)

	__gong__rebuildMapString(stage.Groups, &stage.Groups_mapString)

	__gong__rebuildMapString(stage.GroupToogles, &stage.GroupToogles_mapString)

	__gong__rebuildMapString(stage.Layouts, &stage.Layouts_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
