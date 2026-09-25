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

	Menus                map[*Menu]struct{}
	Menus_instance       map[*Menu]*Menu
	Menus_mapString      map[string]*Menu
	MenuOrder            uint
	Menu_stagedOrder     map[*Menu]uint
	Menu_orderStaged     map[uint]*Menu
	Menus_reference      map[*Menu]*Menu
	Menus_referenceOrder map[*Menu]uint

	// insertion point for slice of pointers maps
	Menu_Buttons_reverseMap map[*Button]*Menu

	OnAfterMenuCreateCallback GongOnAfterCreateInterface[Menu]
	OnAfterMenuUpdateCallback GongOnAfterUpdateInterface[Menu]
	OnAfterMenuDeleteCallback GongOnAfterDeleteInterface[Menu]

	Nodes                map[*Node]struct{}
	Nodes_instance       map[*Node]*Node
	Nodes_mapString      map[string]*Node
	NodeOrder            uint
	Node_stagedOrder     map[*Node]uint
	Node_orderStaged     map[uint]*Node
	Nodes_reference      map[*Node]*Node
	Nodes_referenceOrder map[*Node]uint

	// insertion point for slice of pointers maps
	Node_Children_reverseMap map[*Node]*Node

	Node_Buttons_reverseMap map[*Button]*Node

	OnAfterNodeCreateCallback GongOnAfterCreateInterface[Node]
	OnAfterNodeUpdateCallback GongOnAfterUpdateInterface[Node]
	OnAfterNodeDeleteCallback GongOnAfterDeleteInterface[Node]

	SVGIcons                map[*SVGIcon]struct{}
	SVGIcons_instance       map[*SVGIcon]*SVGIcon
	SVGIcons_mapString      map[string]*SVGIcon
	SVGIconOrder            uint
	SVGIcon_stagedOrder     map[*SVGIcon]uint
	SVGIcon_orderStaged     map[uint]*SVGIcon
	SVGIcons_reference      map[*SVGIcon]*SVGIcon
	SVGIcons_referenceOrder map[*SVGIcon]uint

	// insertion point for slice of pointers maps
	OnAfterSVGIconCreateCallback GongOnAfterCreateInterface[SVGIcon]
	OnAfterSVGIconUpdateCallback GongOnAfterUpdateInterface[SVGIcon]
	OnAfterSVGIconDeleteCallback GongOnAfterDeleteInterface[SVGIcon]

	Trees                map[*Tree]struct{}
	Trees_instance       map[*Tree]*Tree
	Trees_mapString      map[string]*Tree
	TreeOrder            uint
	Tree_stagedOrder     map[*Tree]uint
	Tree_orderStaged     map[uint]*Tree
	Trees_reference      map[*Tree]*Tree
	Trees_referenceOrder map[*Tree]uint

	// insertion point for slice of pointers maps
	Tree_RootNodes_reverseMap map[*Node]*Tree

	OnAfterTreeCreateCallback GongOnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback GongOnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback GongOnAfterDeleteInterface[Tree]

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

	__gong__clearReferences(&stage.Menus_reference, &stage.Menus_instance, &stage.Menus_referenceOrder)

	__gong__clearReferences(&stage.Nodes_reference, &stage.Nodes_instance, &stage.Nodes_referenceOrder)

	__gong__clearReferences(&stage.SVGIcons_reference, &stage.SVGIcons_instance, &stage.SVGIcons_referenceOrder)

	__gong__clearReferences(&stage.Trees_reference, &stage.Trees_instance, &stage.Trees_referenceOrder)

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

	stage.MenuOrder = __gong__recomputeOrder(stage.Menu_stagedOrder)

	stage.NodeOrder = __gong__recomputeOrder(stage.Node_stagedOrder)

	stage.SVGIconOrder = __gong__recomputeOrder(stage.SVGIcon_stagedOrder)

	stage.TreeOrder = __gong__recomputeOrder(stage.Tree_stagedOrder)

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
	case *Menu:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Menus, stage.Menu_stagedOrder))
	case *Node:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Nodes, stage.Node_stagedOrder))
	case *SVGIcon:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SVGIcons, stage.SVGIcon_stagedOrder))
	case *Tree:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Trees, stage.Tree_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/tree/go/models"
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

		Menus:           make(map[*Menu]struct{}),
		Menus_mapString: make(map[string]*Menu),

		Nodes:           make(map[*Node]struct{}),
		Nodes_mapString: make(map[string]*Node),

		SVGIcons:           make(map[*SVGIcon]struct{}),
		SVGIcons_mapString: make(map[string]*SVGIcon),

		Trees:           make(map[*Tree]struct{}),
		Trees_mapString: make(map[string]*Tree),

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

		Menu_stagedOrder: make(map[*Menu]uint),
		Menu_orderStaged: make(map[uint]*Menu),
		Menus_reference:  make(map[*Menu]*Menu),

		Node_stagedOrder: make(map[*Node]uint),
		Node_orderStaged: make(map[uint]*Node),
		Nodes_reference:  make(map[*Node]*Node),

		SVGIcon_stagedOrder: make(map[*SVGIcon]uint),
		SVGIcon_orderStaged: make(map[uint]*SVGIcon),
		SVGIcons_reference:  make(map[*SVGIcon]*SVGIcon),

		Tree_stagedOrder: make(map[*Tree]uint),
		Tree_orderStaged: make(map[uint]*Tree),
		Trees_reference:  make(map[*Tree]*Tree),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Button": &ButtonUnmarshaller{},

			"Menu": &MenuUnmarshaller{},

			"Node": &NodeUnmarshaller{},

			"SVGIcon": &SVGIconUnmarshaller{},

			"Tree": &TreeUnmarshaller{},

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
	case *Menu:
		return any(stage.Menu_orderStaged[order]).(Type)
	case *Node:
		return any(stage.Node_orderStaged[order]).(Type)
	case *SVGIcon:
		return any(stage.SVGIcon_orderStaged[order]).(Type)
	case *Tree:
		return any(stage.Tree_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Menu"] = len(stage.Menus)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["SVGIcon"] = len(stage.SVGIcons)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)
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

// Stage puts menu to the model stage
func (menu *Menu) Stage(stage *Stage) *Menu {
	__gong__stage(stage.Menus, stage.Menu_stagedOrder, stage.Menu_orderStaged, &stage.MenuOrder, stage.Menus_mapString, menu, menu.Name)
	return menu
}

// StagePreserveOrder puts menu to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MenuOrder
// - update stage.MenuOrder accordingly
func (menu *Menu) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Menus, stage.Menu_stagedOrder, stage.Menu_orderStaged, &stage.MenuOrder, stage.Menus_mapString, menu, order, menu.Name)
}

// Unstage removes menu off the model stage
func (menu *Menu) Unstage(stage *Stage) *Menu {
	__gong__unstage(stage.Menus, stage.Menus_mapString, menu, menu.Name)
	return menu
}

// UnstageVoid removes menu off the model stage
func (menu *Menu) UnstageVoid(stage *Stage) {
	menu.Unstage(stage)
}

func (menu *Menu) StageVoid(stage *Stage) {
	menu.Stage(stage)
}

// for satisfaction of GongStruct interface
func (menu *Menu) GetName() (res string) {
	return menu.Name
}

// for satisfaction of GongStruct interface
func (menu *Menu) SetName(name string) {
	menu.Name = name
}

// Stage puts node to the model stage
func (node *Node) Stage(stage *Stage) *Node {
	__gong__stage(stage.Nodes, stage.Node_stagedOrder, stage.Node_orderStaged, &stage.NodeOrder, stage.Nodes_mapString, node, node.Name)
	return node
}

// StagePreserveOrder puts node to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NodeOrder
// - update stage.NodeOrder accordingly
func (node *Node) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Nodes, stage.Node_stagedOrder, stage.Node_orderStaged, &stage.NodeOrder, stage.Nodes_mapString, node, order, node.Name)
}

// Unstage removes node off the model stage
func (node *Node) Unstage(stage *Stage) *Node {
	__gong__unstage(stage.Nodes, stage.Nodes_mapString, node, node.Name)
	return node
}

// UnstageVoid removes node off the model stage
func (node *Node) UnstageVoid(stage *Stage) {
	node.Unstage(stage)
}

func (node *Node) StageVoid(stage *Stage) {
	node.Stage(stage)
}

// for satisfaction of GongStruct interface
func (node *Node) GetName() (res string) {
	return node.Name
}

// for satisfaction of GongStruct interface
func (node *Node) SetName(name string) {
	node.Name = name
}

// Stage puts svgicon to the model stage
func (svgicon *SVGIcon) Stage(stage *Stage) *SVGIcon {
	__gong__stage(stage.SVGIcons, stage.SVGIcon_stagedOrder, stage.SVGIcon_orderStaged, &stage.SVGIconOrder, stage.SVGIcons_mapString, svgicon, svgicon.Name)
	return svgicon
}

// StagePreserveOrder puts svgicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SVGIconOrder
// - update stage.SVGIconOrder accordingly
func (svgicon *SVGIcon) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SVGIcons, stage.SVGIcon_stagedOrder, stage.SVGIcon_orderStaged, &stage.SVGIconOrder, stage.SVGIcons_mapString, svgicon, order, svgicon.Name)
}

// Unstage removes svgicon off the model stage
func (svgicon *SVGIcon) Unstage(stage *Stage) *SVGIcon {
	__gong__unstage(stage.SVGIcons, stage.SVGIcons_mapString, svgicon, svgicon.Name)
	return svgicon
}

// UnstageVoid removes svgicon off the model stage
func (svgicon *SVGIcon) UnstageVoid(stage *Stage) {
	svgicon.Unstage(stage)
}

func (svgicon *SVGIcon) StageVoid(stage *Stage) {
	svgicon.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) GetName() (res string) {
	return svgicon.Name
}

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) SetName(name string) {
	svgicon.Name = name
}

// Stage puts tree to the model stage
func (tree *Tree) Stage(stage *Stage) *Tree {
	__gong__stage(stage.Trees, stage.Tree_stagedOrder, stage.Tree_orderStaged, &stage.TreeOrder, stage.Trees_mapString, tree, tree.Name)
	return tree
}

// StagePreserveOrder puts tree to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TreeOrder
// - update stage.TreeOrder accordingly
func (tree *Tree) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Trees, stage.Tree_stagedOrder, stage.Tree_orderStaged, &stage.TreeOrder, stage.Trees_mapString, tree, order, tree.Name)
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *Stage) *Tree {
	__gong__unstage(stage.Trees, stage.Trees_mapString, tree, tree.Name)
	return tree
}

// UnstageVoid removes tree off the model stage
func (tree *Tree) UnstageVoid(stage *Stage) {
	tree.Unstage(stage)
}

func (tree *Tree) StageVoid(stage *Stage) {
	tree.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tree *Tree) GetName() (res string) {
	return tree.Name
}

// for satisfaction of GongStruct interface
func (tree *Tree) SetName(name string) {
	tree.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Buttons, &stage.Buttons_mapString, &stage.Button_stagedOrder, &stage.ButtonOrder)

	__gong__resetStageType(&stage.Menus, &stage.Menus_mapString, &stage.Menu_stagedOrder, &stage.MenuOrder)

	__gong__resetStageType(&stage.Nodes, &stage.Nodes_mapString, &stage.Node_stagedOrder, &stage.NodeOrder)

	__gong__resetStageType(&stage.SVGIcons, &stage.SVGIcons_mapString, &stage.SVGIcon_stagedOrder, &stage.SVGIconOrder)

	__gong__resetStageType(&stage.Trees, &stage.Trees_mapString, &stage.Tree_stagedOrder, &stage.TreeOrder)

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
	case *Menu:
		return any(stage.Menus_mapString).(map[string]Type)
	case *Node:
		return any(stage.Nodes_mapString).(map[string]Type)
	case *SVGIcon:
		return any(stage.SVGIcons_mapString).(map[string]Type)
	case *Tree:
		return any(stage.Trees_mapString).(map[string]Type)
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
	case *Menu:
		return any(&stage.Menus).(*map[Type]struct{})
	case *Node:
		return any(&stage.Nodes).(*map[Type]struct{})
	case *SVGIcon:
		return any(&stage.SVGIcons).(*map[Type]struct{})
	case *Tree:
		return any(&stage.Trees).(*map[Type]struct{})
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
	case Button:
		return any(&Button{
			SVGIcon: &SVGIcon{Name: "SVGIcon"},
		}).(*Type)
	case Menu:
		return any(&Menu{
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case Node:
		return any(&Node{
			PreceedingSVGIcon: &SVGIcon{Name: "PreceedingSVGIcon"},
			Children: []*Node{{Name: "Children"}},
			Buttons: []*Button{{Name: "Buttons"}},
			Menu: &Menu{Name: "Menu"},
		}).(*Type)
	case Tree:
		return any(&Tree{
			RootNodes: []*Node{{Name: "RootNodes"}},
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
		case "SVGIcon":
			res := make(map[*SVGIcon][]*Button)
			for button := range stage.Buttons {
				if button.SVGIcon != nil {
					svgicon_ := button.SVGIcon
					var buttons []*Button
					_, ok := res[svgicon_]
					if ok {
						buttons = res[svgicon_]
					} else {
						buttons = make([]*Button, 0)
					}
					buttons = append(buttons, button)
					res[svgicon_] = buttons
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Menu
	case Menu:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "PreceedingSVGIcon":
			res := make(map[*SVGIcon][]*Node)
			for node := range stage.Nodes {
				if node.PreceedingSVGIcon != nil {
					svgicon_ := node.PreceedingSVGIcon
					var nodes []*Node
					_, ok := res[svgicon_]
					if ok {
						nodes = res[svgicon_]
					} else {
						nodes = make([]*Node, 0)
					}
					nodes = append(nodes, node)
					res[svgicon_] = nodes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Menu":
			res := make(map[*Menu][]*Node)
			for node := range stage.Nodes {
				if node.Menu != nil {
					menu_ := node.Menu
					var nodes []*Node
					_, ok := res[menu_]
					if ok {
						nodes = res[menu_]
					} else {
						nodes = make([]*Node, 0)
					}
					nodes = append(nodes, node)
					res[menu_] = nodes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
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
	// reverse maps of direct associations of Menu
	case Menu:
		switch fieldname {
		// insertion point for per direct association field
		case "Buttons":
			res := make(map[*Button][]*Menu)
			for menu := range stage.Menus {
				for _, button_ := range menu.Buttons {
					res[button_] = append(res[button_], menu)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "Children":
			res := make(map[*Node][]*Node)
			for node := range stage.Nodes {
				for _, node_ := range node.Children {
					res[node_] = append(res[node_], node)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Buttons":
			res := make(map[*Button][]*Node)
			for node := range stage.Nodes {
				for _, button_ := range node.Buttons {
					res[button_] = append(res[button_], node)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		case "RootNodes":
			res := make(map[*Node][]*Tree)
			for tree := range stage.Trees {
				for _, node_ := range tree.RootNodes {
					res[node_] = append(res[node_], tree)
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
	case *Menu:
		res = any(new(Menu)).(Type)
	case *Node:
		res = any(new(Node)).(Type)
	case *SVGIcon:
		res = any(new(SVGIcon)).(Type)
	case *Tree:
		res = any(new(Tree)).(Type)
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
	case *Menu:
		res = "Menu"
	case *Node:
		res = "Node"
	case *SVGIcon:
		res = "SVGIcon"
	case *Tree:
		res = "Tree"
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
		rf.GongstructName = "Menu"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
		rf.GongstructName = "Node"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
	case *Menu:
		var rf ReverseField
		_ = rf
	case *Node:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Node"
		rf.Fieldname = "Children"
		res = append(res, rf)
		rf.GongstructName = "Tree"
		rf.Fieldname = "RootNodes"
		res = append(res, rf)
	case *SVGIcon:
		var rf ReverseField
		_ = rf
	case *Tree:
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
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SVGIcon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SVGIcon",
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
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
		{
			Name:               "ClientOnX",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ClientOnY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (menu *Menu) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Buttons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Button",
		},
	}
	return
}

func (node *Node) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsWithPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Prefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "FontStyleEnum",
		},
		{
			Name:               "BackgroundColor",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasCheckboxButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsCheckboxDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CheckboxHasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "CheckboxToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CheckboxToolTipPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ToolTipPositionEnum",
		},
		{
			Name:               "HasSecondCheckboxButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSecondCheckboxChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSecondCheckboxDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SecondCheckboxHasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SecondCheckboxToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SecondCheckboxToolTipPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ToolTipPositionEnum",
		},
		{
			Name:               "TextAfterSecondCheckbox",
			GongFieldValueType: GongFieldValueTypeString,
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
		{
			Name:               "ClientOnY",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsInEditMode",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsNodeClickable",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsWithPreceedingIcon",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "PreceedingIcon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PreceedingSVGIcon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SVGIcon",
		},
		{
			Name:                 "Children",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "Buttons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Button",
		},
		{
			Name:                 "Menu",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Menu",
		},
	}
	return
}

func (svgicon *SVGIcon) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tree *Tree) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RootNodes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Node",
		},
		{
			Name:               "HaveSearch",
			GongFieldValueType: GongFieldValueTypeBool,
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
	case "Icon":
		res.valueString = button.Icon
	case "SVGIcon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if button.SVGIcon != nil {
			res.valueString = button.SVGIcon.Name
			res.ids = button.SVGIcon.GongGetUUID(stage)
		}
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", button.IsDisabled)
		res.valueBool = button.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", button.HasToolTip)
		res.valueBool = button.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = button.ToolTipText
	case "ToolTipPosition":
		enum := button.ToolTipPosition
		res.valueString = enum.ToCodeString()
	case "ClientOnX":
		res.valueString = fmt.Sprintf("%f", button.ClientOnX)
		res.valueFloat = button.ClientOnX
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ClientOnY":
		res.valueString = fmt.Sprintf("%f", button.ClientOnY)
		res.valueFloat = button.ClientOnY
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (menu *Menu) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = menu.Name
	case "Buttons":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range menu.Buttons {
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

func (node *Node) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = node.Name
	case "IsWithPrefix":
		res.valueString = fmt.Sprintf("%t", node.IsWithPrefix)
		res.valueBool = node.IsWithPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Prefix":
		res.valueString = node.Prefix
	case "FontStyle":
		enum := node.FontStyle
		res.valueString = enum.ToCodeString()
	case "BackgroundColor":
		res.valueString = node.BackgroundColor
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", node.IsExpanded)
		res.valueBool = node.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasCheckboxButton":
		res.valueString = fmt.Sprintf("%t", node.HasCheckboxButton)
		res.valueBool = node.HasCheckboxButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", node.IsChecked)
		res.valueBool = node.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsCheckboxDisabled":
		res.valueString = fmt.Sprintf("%t", node.IsCheckboxDisabled)
		res.valueBool = node.IsCheckboxDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CheckboxHasToolTip":
		res.valueString = fmt.Sprintf("%t", node.CheckboxHasToolTip)
		res.valueBool = node.CheckboxHasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CheckboxToolTipText":
		res.valueString = node.CheckboxToolTipText
	case "CheckboxToolTipPosition":
		enum := node.CheckboxToolTipPosition
		res.valueString = enum.ToCodeString()
	case "HasSecondCheckboxButton":
		res.valueString = fmt.Sprintf("%t", node.HasSecondCheckboxButton)
		res.valueBool = node.HasSecondCheckboxButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSecondCheckboxChecked":
		res.valueString = fmt.Sprintf("%t", node.IsSecondCheckboxChecked)
		res.valueBool = node.IsSecondCheckboxChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSecondCheckboxDisabled":
		res.valueString = fmt.Sprintf("%t", node.IsSecondCheckboxDisabled)
		res.valueBool = node.IsSecondCheckboxDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SecondCheckboxHasToolTip":
		res.valueString = fmt.Sprintf("%t", node.SecondCheckboxHasToolTip)
		res.valueBool = node.SecondCheckboxHasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SecondCheckboxToolTipText":
		res.valueString = node.SecondCheckboxToolTipText
	case "SecondCheckboxToolTipPosition":
		enum := node.SecondCheckboxToolTipPosition
		res.valueString = enum.ToCodeString()
	case "TextAfterSecondCheckbox":
		res.valueString = node.TextAfterSecondCheckbox
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", node.HasToolTip)
		res.valueBool = node.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = node.ToolTipText
	case "ToolTipPosition":
		enum := node.ToolTipPosition
		res.valueString = enum.ToCodeString()
	case "ClientOnY":
		res.valueString = fmt.Sprintf("%f", node.ClientOnY)
		res.valueFloat = node.ClientOnY
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsInEditMode":
		res.valueString = fmt.Sprintf("%t", node.IsInEditMode)
		res.valueBool = node.IsInEditMode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsNodeClickable":
		res.valueString = fmt.Sprintf("%t", node.IsNodeClickable)
		res.valueBool = node.IsNodeClickable
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsWithPreceedingIcon":
		res.valueString = fmt.Sprintf("%t", node.IsWithPreceedingIcon)
		res.valueBool = node.IsWithPreceedingIcon
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PreceedingIcon":
		res.valueString = node.PreceedingIcon
	case "PreceedingSVGIcon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if node.PreceedingSVGIcon != nil {
			res.valueString = node.PreceedingSVGIcon.Name
			res.ids = node.PreceedingSVGIcon.GongGetUUID(stage)
		}
	case "Children":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range node.Children {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Buttons":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range node.Buttons {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Menu":
		res.GongFieldValueType = GongFieldValueTypePointer
		if node.Menu != nil {
			res.valueString = node.Menu.Name
			res.ids = node.Menu.GongGetUUID(stage)
		}
	}
	return
}

func (svgicon *SVGIcon) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svgicon.Name
	case "SVG":
		res.valueString = svgicon.SVG
	}
	return
}

func (tree *Tree) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tree.Name
	case "RootNodes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range tree.RootNodes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HaveSearch":
		res.valueString = fmt.Sprintf("%t", tree.HaveSearch)
		res.valueBool = tree.HaveSearch
		res.GongFieldValueType = GongFieldValueTypeBool
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

func (menu *Menu) GongGetGongstructName() string {
	return "Menu"
}

func (node *Node) GongGetGongstructName() string {
	return "Node"
}

func (svgicon *SVGIcon) GongGetGongstructName() string {
	return "SVGIcon"
}

func (tree *Tree) GongGetGongstructName() string {
	return "Tree"
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

	__gong__rebuildMapString(stage.Menus, &stage.Menus_mapString)

	__gong__rebuildMapString(stage.Nodes, &stage.Nodes_mapString)

	__gong__rebuildMapString(stage.SVGIcons, &stage.SVGIcons_mapString)

	__gong__rebuildMapString(stage.Trees, &stage.Trees_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
