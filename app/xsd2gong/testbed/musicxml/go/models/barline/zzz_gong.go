// generated code - do not edit
package barline

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
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

type GongMarshallingMode = gong_runtime.GongMarshallingMode

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal = gong_runtime.GongMarshallingNormal

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit = gong_runtime.GongMarshallingAppendCommit
)

type gongStageNavigationMode = gong_runtime.StageNavigationMode

const (
	GongNavigationModeNormal     = gong_runtime.GongNavigationModeNormal
	GongNavigationModeNavigating = gong_runtime.GongNavigationModeNavigating
)

const (
	GongProbeTreeSidebarSuffix           = gong_runtime.GongProbeTreeSidebarSuffix
	GongProbeNavigationTreeSidebarSuffix = gong_runtime.GongProbeNavigationTreeSidebarSuffix
	GongProbeTableSuffix                 = gong_runtime.GongProbeTableSuffix
	GongProbeNotificationTableSuffix     = gong_runtime.GongProbeNotificationTableSuffix
	GongProbeFormSuffix                  = gong_runtime.GongProbeFormSuffix
	GongProbeSplitSuffix                 = gong_runtime.GongProbeSplitSuffix
	GongProbeLoadSuffix                  = gong_runtime.GongProbeLoadSuffix
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
	gong_runtime.StageCore

	// insertion point for definition of arrays registering instances
	BarStyleColors                map[*BarStyleColor]struct{}
	BarStyleColors_instance       map[*BarStyleColor]*BarStyleColor
	BarStyleColors_mapString      map[string]*BarStyleColor
	BarStyleColorOrder            uint
	BarStyleColor_stagedOrder     map[*BarStyleColor]uint
	BarStyleColor_orderStaged     map[uint]*BarStyleColor
	BarStyleColors_reference      map[*BarStyleColor]*BarStyleColor
	BarStyleColors_referenceOrder map[*BarStyleColor]uint

	// insertion point for slice of pointers maps
	OnAfterBarStyleColorCreateCallback GongOnAfterCreateInterface[BarStyleColor]
	OnAfterBarStyleColorUpdateCallback GongOnAfterUpdateInterface[BarStyleColor]
	OnAfterBarStyleColorDeleteCallback GongOnAfterDeleteInterface[BarStyleColor]

	Barlines                map[*Barline]struct{}
	Barlines_instance       map[*Barline]*Barline
	Barlines_mapString      map[string]*Barline
	BarlineOrder            uint
	Barline_stagedOrder     map[*Barline]uint
	Barline_orderStaged     map[uint]*Barline
	Barlines_reference      map[*Barline]*Barline
	Barlines_referenceOrder map[*Barline]uint

	// insertion point for slice of pointers maps
	OnAfterBarlineCreateCallback GongOnAfterCreateInterface[Barline]
	OnAfterBarlineUpdateCallback GongOnAfterUpdateInterface[Barline]
	OnAfterBarlineDeleteCallback GongOnAfterDeleteInterface[Barline]

	Endings                map[*Ending]struct{}
	Endings_instance       map[*Ending]*Ending
	Endings_mapString      map[string]*Ending
	EndingOrder            uint
	Ending_stagedOrder     map[*Ending]uint
	Ending_orderStaged     map[uint]*Ending
	Endings_reference      map[*Ending]*Ending
	Endings_referenceOrder map[*Ending]uint

	// insertion point for slice of pointers maps
	OnAfterEndingCreateCallback GongOnAfterCreateInterface[Ending]
	OnAfterEndingUpdateCallback GongOnAfterUpdateInterface[Ending]
	OnAfterEndingDeleteCallback GongOnAfterDeleteInterface[Ending]

	Repeats                map[*Repeat]struct{}
	Repeats_instance       map[*Repeat]*Repeat
	Repeats_mapString      map[string]*Repeat
	RepeatOrder            uint
	Repeat_stagedOrder     map[*Repeat]uint
	Repeat_orderStaged     map[uint]*Repeat
	Repeats_reference      map[*Repeat]*Repeat
	Repeats_referenceOrder map[*Repeat]uint

	// insertion point for slice of pointers maps
	OnAfterRepeatCreateCallback GongOnAfterCreateInterface[Repeat]
	OnAfterRepeatUpdateCallback GongOnAfterUpdateInterface[Repeat]
	OnAfterRepeatDeleteCallback GongOnAfterDeleteInterface[Repeat]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

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
}

type GongStage = Stage

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterBeforeCommitHook(func() { hook(s) })
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterAfterCommitHook(func() { hook(s) })
}

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	return stage.StageCore.ApplyBackwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ApplyForwardCommit applies the commit after the current one
func (stage *Stage) ApplyForwardCommit() error {
	return stage.StageCore.ApplyForwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	stage.StageCore.ResetHardCore()

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.StageCore.SquashCore()

	// insertion point for clear references
	__gong__clearReferences(&stage.BarStyleColors_reference, &stage.BarStyleColors_instance, &stage.BarStyleColors_referenceOrder)

	__gong__clearReferences(&stage.Barlines_reference, &stage.Barlines_instance, &stage.Barlines_referenceOrder)

	__gong__clearReferences(&stage.Endings_reference, &stage.Endings_instance, &stage.Endings_referenceOrder)

	__gong__clearReferences(&stage.Repeats_reference, &stage.Repeats_instance, &stage.Repeats_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	stage.StageCore.EndSquash()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.BarStyleColorOrder = __gong__recomputeOrder(stage.BarStyleColor_stagedOrder)

	stage.BarlineOrder = __gong__recomputeOrder(stage.Barline_stagedOrder)

	stage.EndingOrder = __gong__recomputeOrder(stage.Ending_stagedOrder)

	stage.RepeatOrder = __gong__recomputeOrder(stage.Repeat_stagedOrder)

	// end of insertion point for max order recomputation
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

// insertion point for stage ops
func (*BarStyleColor) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BarStyleColors, stage.BarStyleColor_stagedOrder)
}

func (*BarStyleColor) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BarStyleColor_orderStaged[order]
}

func (*BarStyleColor) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BarStyleColors_mapString
}

func (*BarStyleColor) GongGetInstancesSet(stage *Stage) any {
	return &stage.BarStyleColors
}

func (*BarStyleColor) GongNewInstance() any {
	return new(BarStyleColor)
}

func (*Barline) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Barlines, stage.Barline_stagedOrder)
}

func (*Barline) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Barline_orderStaged[order]
}

func (*Barline) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Barlines_mapString
}

func (*Barline) GongGetInstancesSet(stage *Stage) any {
	return &stage.Barlines
}

func (*Barline) GongNewInstance() any {
	return new(Barline)
}

func (*Ending) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Endings, stage.Ending_stagedOrder)
}

func (*Ending) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Ending_orderStaged[order]
}

func (*Ending) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Endings_mapString
}

func (*Ending) GongGetInstancesSet(stage *Stage) any {
	return &stage.Endings
}

func (*Ending) GongNewInstance() any {
	return new(Ending)
}

func (*Repeat) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Repeats, stage.Repeat_stagedOrder)
}

func (*Repeat) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Repeat_orderStaged[order]
}

func (*Repeat) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Repeats_mapString
}

func (*Repeat) GongGetInstancesSet(stage *Stage) any {
	return &stage.Repeats
}

func (*Repeat) GongNewInstance() any {
	return new(Repeat)
}


// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	if stage == nil {
		return nil
	}
	var t T
	return t.GongGetInstancesByOrder(stage).([]T)
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
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
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
		BarStyleColors:           make(map[*BarStyleColor]struct{}),
		BarStyleColors_mapString: make(map[string]*BarStyleColor),

		Barlines:           make(map[*Barline]struct{}),
		Barlines_mapString: make(map[string]*Barline),

		Endings:           make(map[*Ending]struct{}),
		Endings_mapString: make(map[string]*Ending),

		Repeats:           make(map[*Repeat]struct{}),
		Repeats_mapString: make(map[string]*Repeat),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		BarStyleColor_stagedOrder: make(map[*BarStyleColor]uint),
		BarStyleColor_orderStaged: make(map[uint]*BarStyleColor),
		BarStyleColors_reference:  make(map[*BarStyleColor]*BarStyleColor),

		Barline_stagedOrder: make(map[*Barline]uint),
		Barline_orderStaged: make(map[uint]*Barline),
		Barlines_reference:  make(map[*Barline]*Barline),

		Ending_stagedOrder: make(map[*Ending]uint),
		Ending_orderStaged: make(map[uint]*Ending),
		Endings_reference:  make(map[*Ending]*Ending),

		Repeat_stagedOrder: make(map[*Repeat]uint),
		Repeat_orderStaged: make(map[uint]*Repeat),
		Repeats_reference:  make(map[*Repeat]*Repeat),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"BarStyleColor": &BarStyleColorUnmarshaller{},

			"Barline": &BarlineUnmarshaller{},

			"Ending": &EndingUnmarshaller{},

			"Repeat": &RepeatUnmarshaller{},

			// end of insertion point
		},
	}
	stage.StageCore.SetName(name)
	stage.StageCore.SetNavigationMode(GongNavigationModeNormal)

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
	if stage == nil {
		return
	}
	var t Type
	val := t.GongGetInstanceFromOrder(stage, order)
	if val != nil {
		res = val.(Type)
	}
	return
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.StageCore.GetBeforeCommitHooks()
	stage.StageCore.ClearBeforeCommitHooks()
	tmp3 := stage.StageCore.GetAfterCommitHooks()
	stage.StageCore.ClearAfterCommitHooks()
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.StageCore.SetBeforeCommitHooks(tmp2)
	stage.StageCore.SetAfterCommitHooks(tmp3)
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
	stage.StageCore.RunBeforeCommitHooks()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.IsInDeltaMode() && stage.GetNavigationMode() == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
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
	stage.StageCore.RunAfterCommitHooks()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BarStyleColor"] = len(stage.BarStyleColors)
	stage.Map_GongStructName_InstancesNb["Barline"] = len(stage.Barlines)
	stage.Map_GongStructName_InstancesNb["Ending"] = len(stage.Endings)
	stage.Map_GongStructName_InstancesNb["Repeat"] = len(stage.Repeats)
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
// Stage puts barstylecolor to the model stage
func (barstylecolor *BarStyleColor) Stage(stage *Stage) *BarStyleColor {
	__gong__stage(stage.BarStyleColors, stage.BarStyleColor_stagedOrder, stage.BarStyleColor_orderStaged, &stage.BarStyleColorOrder, stage.BarStyleColors_mapString, barstylecolor, barstylecolor.Name)
	return barstylecolor
}

// StagePreserveOrder puts barstylecolor to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BarStyleColorOrder
// - update stage.BarStyleColorOrder accordingly
func (barstylecolor *BarStyleColor) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BarStyleColors, stage.BarStyleColor_stagedOrder, stage.BarStyleColor_orderStaged, &stage.BarStyleColorOrder, stage.BarStyleColors_mapString, barstylecolor, order, barstylecolor.Name)
}

// Unstage removes barstylecolor off the model stage
func (barstylecolor *BarStyleColor) Unstage(stage *Stage) *BarStyleColor {
	__gong__unstage(stage.BarStyleColors, stage.BarStyleColors_mapString, barstylecolor, barstylecolor.Name)
	return barstylecolor
}

// UnstageVoid removes barstylecolor off the model stage
func (barstylecolor *BarStyleColor) UnstageVoid(stage *Stage) {
	barstylecolor.Unstage(stage)
}

func (barstylecolor *BarStyleColor) StageVoid(stage *Stage) {
	barstylecolor.Stage(stage)
}

// for satisfaction of GongStruct interface
func (barstylecolor *BarStyleColor) GetName() (res string) {
	return barstylecolor.Name
}

// for satisfaction of GongStruct interface
func (barstylecolor *BarStyleColor) SetName(name string) {
	barstylecolor.Name = name
}

// Stage puts barline to the model stage
func (barline *Barline) Stage(stage *Stage) *Barline {
	__gong__stage(stage.Barlines, stage.Barline_stagedOrder, stage.Barline_orderStaged, &stage.BarlineOrder, stage.Barlines_mapString, barline, barline.Name)
	return barline
}

// StagePreserveOrder puts barline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BarlineOrder
// - update stage.BarlineOrder accordingly
func (barline *Barline) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Barlines, stage.Barline_stagedOrder, stage.Barline_orderStaged, &stage.BarlineOrder, stage.Barlines_mapString, barline, order, barline.Name)
}

// Unstage removes barline off the model stage
func (barline *Barline) Unstage(stage *Stage) *Barline {
	__gong__unstage(stage.Barlines, stage.Barlines_mapString, barline, barline.Name)
	return barline
}

// UnstageVoid removes barline off the model stage
func (barline *Barline) UnstageVoid(stage *Stage) {
	barline.Unstage(stage)
}

func (barline *Barline) StageVoid(stage *Stage) {
	barline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (barline *Barline) GetName() (res string) {
	return barline.Name
}

// for satisfaction of GongStruct interface
func (barline *Barline) SetName(name string) {
	barline.Name = name
}

// Stage puts ending to the model stage
func (ending *Ending) Stage(stage *Stage) *Ending {
	__gong__stage(stage.Endings, stage.Ending_stagedOrder, stage.Ending_orderStaged, &stage.EndingOrder, stage.Endings_mapString, ending, ending.Name)
	return ending
}

// StagePreserveOrder puts ending to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EndingOrder
// - update stage.EndingOrder accordingly
func (ending *Ending) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Endings, stage.Ending_stagedOrder, stage.Ending_orderStaged, &stage.EndingOrder, stage.Endings_mapString, ending, order, ending.Name)
}

// Unstage removes ending off the model stage
func (ending *Ending) Unstage(stage *Stage) *Ending {
	__gong__unstage(stage.Endings, stage.Endings_mapString, ending, ending.Name)
	return ending
}

// UnstageVoid removes ending off the model stage
func (ending *Ending) UnstageVoid(stage *Stage) {
	ending.Unstage(stage)
}

func (ending *Ending) StageVoid(stage *Stage) {
	ending.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ending *Ending) GetName() (res string) {
	return ending.Name
}

// for satisfaction of GongStruct interface
func (ending *Ending) SetName(name string) {
	ending.Name = name
}

// Stage puts repeat to the model stage
func (repeat *Repeat) Stage(stage *Stage) *Repeat {
	__gong__stage(stage.Repeats, stage.Repeat_stagedOrder, stage.Repeat_orderStaged, &stage.RepeatOrder, stage.Repeats_mapString, repeat, repeat.Name)
	return repeat
}

// StagePreserveOrder puts repeat to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RepeatOrder
// - update stage.RepeatOrder accordingly
func (repeat *Repeat) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Repeats, stage.Repeat_stagedOrder, stage.Repeat_orderStaged, &stage.RepeatOrder, stage.Repeats_mapString, repeat, order, repeat.Name)
}

// Unstage removes repeat off the model stage
func (repeat *Repeat) Unstage(stage *Stage) *Repeat {
	__gong__unstage(stage.Repeats, stage.Repeats_mapString, repeat, repeat.Name)
	return repeat
}

// UnstageVoid removes repeat off the model stage
func (repeat *Repeat) UnstageVoid(stage *Stage) {
	repeat.Unstage(stage)
}

func (repeat *Repeat) StageVoid(stage *Stage) {
	repeat.Stage(stage)
}

// for satisfaction of GongStruct interface
func (repeat *Repeat) GetName() (res string) {
	return repeat.Name
}

// for satisfaction of GongStruct interface
func (repeat *Repeat) SetName(name string) {
	repeat.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.BarStyleColors, &stage.BarStyleColors_mapString, &stage.BarStyleColor_stagedOrder, &stage.BarStyleColorOrder)

	__gong__resetStageType(&stage.Barlines, &stage.Barlines_mapString, &stage.Barline_stagedOrder, &stage.BarlineOrder)

	__gong__resetStageType(&stage.Endings, &stage.Endings_mapString, &stage.Ending_stagedOrder, &stage.EndingOrder)

	__gong__resetStageType(&stage.Repeats, &stage.Repeats_mapString, &stage.Repeat_stagedOrder, &stage.RepeatOrder)

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
type Gongstruct interface {
	GongGetAssociationName() any
}

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

	GongGetInstancesByOrder(stage *Stage) any
	GongGetInstanceFromOrder(stage *Stage, order uint) any
	GongGetInstancesMapByName(stage *Stage) any
	GongGetInstancesSet(stage *Stage) any
	GongNewInstance() any
	GongGetReverseFields() []GongReverseField
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
	if stage == nil {
		return nil
	}
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetInstancesMapByName(stage).(map[string]Type)
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesSet(stage).(*map[Type]struct{})
}

// insertion point for instance with special fields
func (BarStyleColor) GongGetAssociationName() any {
	return &BarStyleColor{
	}
}

func (Barline) GongGetAssociationName() any {
	return &Barline{
			BarStyle: &BarStyleColor{Name: "BarStyle"},
			Ending: &Ending{Name: "Ending"},
			Repeat: &Repeat{Name: "Repeat"},
	}
}

func (Ending) GongGetAssociationName() any {
	return &Ending{
	}
}

func (Repeat) GongGetAssociationName() any {
	return &Repeat{
	}
}


// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var t Type
	return t.GongGetAssociationName().(*Type)
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
	// reverse maps of direct associations of BarStyleColor
	case BarStyleColor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Barline
	case Barline:
		switch fieldname {
		// insertion point for per direct association field
		case "BarStyle":
			res := make(map[*BarStyleColor][]*Barline)
			for barline := range stage.Barlines {
				if barline.BarStyle != nil {
					barstylecolor_ := barline.BarStyle
					var barlines []*Barline
					_, ok := res[barstylecolor_]
					if ok {
						barlines = res[barstylecolor_]
					} else {
						barlines = make([]*Barline, 0)
					}
					barlines = append(barlines, barline)
					res[barstylecolor_] = barlines
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ending":
			res := make(map[*Ending][]*Barline)
			for barline := range stage.Barlines {
				if barline.Ending != nil {
					ending_ := barline.Ending
					var barlines []*Barline
					_, ok := res[ending_]
					if ok {
						barlines = res[ending_]
					} else {
						barlines = make([]*Barline, 0)
					}
					barlines = append(barlines, barline)
					res[ending_] = barlines
				}
			}
			return any(res).(map[*End][]*Start)
		case "Repeat":
			res := make(map[*Repeat][]*Barline)
			for barline := range stage.Barlines {
				if barline.Repeat != nil {
					repeat_ := barline.Repeat
					var barlines []*Barline
					_, ok := res[repeat_]
					if ok {
						barlines = res[repeat_]
					} else {
						barlines = make([]*Barline, 0)
					}
					barlines = append(barlines, barline)
					res[repeat_] = barlines
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Ending
	case Ending:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Repeat
	case Repeat:
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
	// reverse maps of direct associations of BarStyleColor
	case BarStyleColor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Barline
	case Barline:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Ending
	case Ending:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Repeat
	case Repeat:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var t Type
	return t.GongNewInstance().(Type)
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
	var t Type
	if any(t) == nil {
		return ""
	}
	return t.GongGetGongstructName()
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

// insertion point for generic get reverse fields
func (*BarStyleColor) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Barline) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Ending) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Repeat) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}


func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetReverseFields()
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (barstylecolor *BarStyleColor) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBarStyle",
		},
	}
	return
}

func (barline *Barline) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumRightLeftMiddle",
		},
		{
			Name:               "Segno",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Coda",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Divisions",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BarStyle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BarStyleColor",
		},
		{
			Name:                 "Ending",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Ending",
		},
		{
			Name:                 "Repeat",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Repeat",
		},
	}
	return
}

func (ending *Ending) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartStopDiscontinue",
		},
		{
			Name:               "EndLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TextX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TextY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "System",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumSystemRelation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (repeat *Repeat) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBackwardForward",
		},
		{
			Name:               "Times",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Winged",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumWinged",
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
func (barstylecolor *BarStyleColor) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = barstylecolor.Name
	case "Color":
		res.valueString = barstylecolor.Color
	case "Value":
		enum := barstylecolor.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (barline *Barline) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = barline.Name
	case "Location":
		enum := barline.Location
		res.valueString = enum.ToCodeString()
	case "Segno":
		res.valueString = barline.Segno
	case "Coda":
		res.valueString = barline.Coda
	case "Divisions":
		res.valueString = barline.Divisions
	case "Id":
		res.valueString = barline.Id
	case "BarStyle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if barline.BarStyle != nil {
			res.valueString = barline.BarStyle.Name
			res.ids = barline.BarStyle.GongGetUUID(stage)
		}
	case "Ending":
		res.GongFieldValueType = GongFieldValueTypePointer
		if barline.Ending != nil {
			res.valueString = barline.Ending.Name
			res.ids = barline.Ending.GongGetUUID(stage)
		}
	case "Repeat":
		res.GongFieldValueType = GongFieldValueTypePointer
		if barline.Repeat != nil {
			res.valueString = barline.Repeat.Name
			res.ids = barline.Repeat.GongGetUUID(stage)
		}
	}
	return
}

func (ending *Ending) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ending.Name
	case "Number":
		res.valueString = ending.Number
	case "Type":
		enum := ending.Type
		res.valueString = enum.ToCodeString()
	case "EndLength":
		res.valueString = ending.EndLength
	case "TextX":
		res.valueString = ending.TextX
	case "TextY":
		res.valueString = ending.TextY
	case "DefaultX":
		res.valueString = ending.DefaultX
	case "DefaultY":
		res.valueString = ending.DefaultY
	case "RelativeX":
		res.valueString = ending.RelativeX
	case "RelativeY":
		res.valueString = ending.RelativeY
	case "FontFamily":
		res.valueString = ending.FontFamily
	case "FontSize":
		res.valueString = ending.FontSize
	case "Color":
		res.valueString = ending.Color
	case "System":
		enum := ending.System
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = ending.Value
	}
	return
}

func (repeat *Repeat) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = repeat.Name
	case "Direction":
		enum := repeat.Direction
		res.valueString = enum.ToCodeString()
	case "Times":
		res.valueString = fmt.Sprintf("%d", repeat.Times)
		res.valueInt = repeat.Times
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Winged":
		enum := repeat.Winged
		res.valueString = enum.ToCodeString()
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
func (barstylecolor *BarStyleColor) GongGetGongstructName() string {
	return "BarStyleColor"
}

func (barline *Barline) GongGetGongstructName() string {
	return "Barline"
}

func (ending *Ending) GongGetGongstructName() string {
	return "Ending"
}

func (repeat *Repeat) GongGetGongstructName() string {
	return "Repeat"
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
	__gong__rebuildMapString(stage.BarStyleColors, &stage.BarStyleColors_mapString)

	__gong__rebuildMapString(stage.Barlines, &stage.Barlines_mapString)

	__gong__rebuildMapString(stage.Endings, &stage.Endings_mapString)

	__gong__rebuildMapString(stage.Repeats, &stage.Repeats_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
