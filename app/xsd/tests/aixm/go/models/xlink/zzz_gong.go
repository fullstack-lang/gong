// generated code - do not edit
package xlink

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
	ArcTypes                map[*ArcType]struct{}
	ArcTypes_instance       map[*ArcType]*ArcType
	ArcTypes_mapString      map[string]*ArcType
	ArcTypeOrder            uint
	ArcType_stagedOrder     map[*ArcType]uint
	ArcType_orderStaged     map[uint]*ArcType
	ArcTypes_reference      map[*ArcType]*ArcType
	ArcTypes_referenceOrder map[*ArcType]uint

	// insertion point for slice of pointers maps
	ArcType_Title_1_reverseMap map[*TitleEltType]*ArcType

	OnAfterArcTypeCreateCallback GongOnAfterCreateInterface[ArcType]
	OnAfterArcTypeUpdateCallback GongOnAfterUpdateInterface[ArcType]
	OnAfterArcTypeDeleteCallback GongOnAfterDeleteInterface[ArcType]

	Extendeds                map[*Extended]struct{}
	Extendeds_instance       map[*Extended]*Extended
	Extendeds_mapString      map[string]*Extended
	ExtendedOrder            uint
	Extended_stagedOrder     map[*Extended]uint
	Extended_orderStaged     map[uint]*Extended
	Extendeds_reference      map[*Extended]*Extended
	Extendeds_referenceOrder map[*Extended]uint

	// insertion point for slice of pointers maps
	OnAfterExtendedCreateCallback GongOnAfterCreateInterface[Extended]
	OnAfterExtendedUpdateCallback GongOnAfterUpdateInterface[Extended]
	OnAfterExtendedDeleteCallback GongOnAfterDeleteInterface[Extended]

	LocatorTypes                map[*LocatorType]struct{}
	LocatorTypes_instance       map[*LocatorType]*LocatorType
	LocatorTypes_mapString      map[string]*LocatorType
	LocatorTypeOrder            uint
	LocatorType_stagedOrder     map[*LocatorType]uint
	LocatorType_orderStaged     map[uint]*LocatorType
	LocatorTypes_reference      map[*LocatorType]*LocatorType
	LocatorTypes_referenceOrder map[*LocatorType]uint

	// insertion point for slice of pointers maps
	LocatorType_Title_1_reverseMap map[*TitleEltType]*LocatorType

	OnAfterLocatorTypeCreateCallback GongOnAfterCreateInterface[LocatorType]
	OnAfterLocatorTypeUpdateCallback GongOnAfterUpdateInterface[LocatorType]
	OnAfterLocatorTypeDeleteCallback GongOnAfterDeleteInterface[LocatorType]

	ResourceTypes                map[*ResourceType]struct{}
	ResourceTypes_instance       map[*ResourceType]*ResourceType
	ResourceTypes_mapString      map[string]*ResourceType
	ResourceTypeOrder            uint
	ResourceType_stagedOrder     map[*ResourceType]uint
	ResourceType_orderStaged     map[uint]*ResourceType
	ResourceTypes_reference      map[*ResourceType]*ResourceType
	ResourceTypes_referenceOrder map[*ResourceType]uint

	// insertion point for slice of pointers maps
	OnAfterResourceTypeCreateCallback GongOnAfterCreateInterface[ResourceType]
	OnAfterResourceTypeUpdateCallback GongOnAfterUpdateInterface[ResourceType]
	OnAfterResourceTypeDeleteCallback GongOnAfterDeleteInterface[ResourceType]

	Simples                map[*Simple]struct{}
	Simples_instance       map[*Simple]*Simple
	Simples_mapString      map[string]*Simple
	SimpleOrder            uint
	Simple_stagedOrder     map[*Simple]uint
	Simple_orderStaged     map[uint]*Simple
	Simples_reference      map[*Simple]*Simple
	Simples_referenceOrder map[*Simple]uint

	// insertion point for slice of pointers maps
	OnAfterSimpleCreateCallback GongOnAfterCreateInterface[Simple]
	OnAfterSimpleUpdateCallback GongOnAfterUpdateInterface[Simple]
	OnAfterSimpleDeleteCallback GongOnAfterDeleteInterface[Simple]

	TitleEltTypes                map[*TitleEltType]struct{}
	TitleEltTypes_instance       map[*TitleEltType]*TitleEltType
	TitleEltTypes_mapString      map[string]*TitleEltType
	TitleEltTypeOrder            uint
	TitleEltType_stagedOrder     map[*TitleEltType]uint
	TitleEltType_orderStaged     map[uint]*TitleEltType
	TitleEltTypes_reference      map[*TitleEltType]*TitleEltType
	TitleEltTypes_referenceOrder map[*TitleEltType]uint

	// insertion point for slice of pointers maps
	OnAfterTitleEltTypeCreateCallback GongOnAfterCreateInterface[TitleEltType]
	OnAfterTitleEltTypeUpdateCallback GongOnAfterUpdateInterface[TitleEltType]
	OnAfterTitleEltTypeDeleteCallback GongOnAfterDeleteInterface[TitleEltType]

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
	__gong__clearReferences(&stage.ArcTypes_reference, &stage.ArcTypes_instance, &stage.ArcTypes_referenceOrder)

	__gong__clearReferences(&stage.Extendeds_reference, &stage.Extendeds_instance, &stage.Extendeds_referenceOrder)

	__gong__clearReferences(&stage.LocatorTypes_reference, &stage.LocatorTypes_instance, &stage.LocatorTypes_referenceOrder)

	__gong__clearReferences(&stage.ResourceTypes_reference, &stage.ResourceTypes_instance, &stage.ResourceTypes_referenceOrder)

	__gong__clearReferences(&stage.Simples_reference, &stage.Simples_instance, &stage.Simples_referenceOrder)

	__gong__clearReferences(&stage.TitleEltTypes_reference, &stage.TitleEltTypes_instance, &stage.TitleEltTypes_referenceOrder)

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
	stage.ArcTypeOrder = __gong__recomputeOrder(stage.ArcType_stagedOrder)

	stage.ExtendedOrder = __gong__recomputeOrder(stage.Extended_stagedOrder)

	stage.LocatorTypeOrder = __gong__recomputeOrder(stage.LocatorType_stagedOrder)

	stage.ResourceTypeOrder = __gong__recomputeOrder(stage.ResourceType_stagedOrder)

	stage.SimpleOrder = __gong__recomputeOrder(stage.Simple_stagedOrder)

	stage.TitleEltTypeOrder = __gong__recomputeOrder(stage.TitleEltType_stagedOrder)

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
func (*ArcType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ArcTypes, stage.ArcType_stagedOrder)
}

func (*ArcType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ArcType_orderStaged[order]
}

func (*ArcType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ArcTypes_mapString
}

func (*ArcType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ArcTypes
}

func (*ArcType) GongNewInstance() any {
	return new(ArcType)
}

func (*Extended) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Extendeds, stage.Extended_stagedOrder)
}

func (*Extended) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Extended_orderStaged[order]
}

func (*Extended) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Extendeds_mapString
}

func (*Extended) GongGetInstancesSet(stage *Stage) any {
	return &stage.Extendeds
}

func (*Extended) GongNewInstance() any {
	return new(Extended)
}

func (*LocatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LocatorTypes, stage.LocatorType_stagedOrder)
}

func (*LocatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LocatorType_orderStaged[order]
}

func (*LocatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LocatorTypes_mapString
}

func (*LocatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LocatorTypes
}

func (*LocatorType) GongNewInstance() any {
	return new(LocatorType)
}

func (*ResourceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ResourceTypes, stage.ResourceType_stagedOrder)
}

func (*ResourceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ResourceType_orderStaged[order]
}

func (*ResourceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ResourceTypes_mapString
}

func (*ResourceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ResourceTypes
}

func (*ResourceType) GongNewInstance() any {
	return new(ResourceType)
}

func (*Simple) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Simples, stage.Simple_stagedOrder)
}

func (*Simple) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Simple_orderStaged[order]
}

func (*Simple) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Simples_mapString
}

func (*Simple) GongGetInstancesSet(stage *Stage) any {
	return &stage.Simples
}

func (*Simple) GongNewInstance() any {
	return new(Simple)
}

func (*TitleEltType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TitleEltTypes, stage.TitleEltType_stagedOrder)
}

func (*TitleEltType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TitleEltType_orderStaged[order]
}

func (*TitleEltType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TitleEltTypes_mapString
}

func (*TitleEltType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TitleEltTypes
}

func (*TitleEltType) GongNewInstance() any {
	return new(TitleEltType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"
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
		ArcTypes:           make(map[*ArcType]struct{}),
		ArcTypes_mapString: make(map[string]*ArcType),

		Extendeds:           make(map[*Extended]struct{}),
		Extendeds_mapString: make(map[string]*Extended),

		LocatorTypes:           make(map[*LocatorType]struct{}),
		LocatorTypes_mapString: make(map[string]*LocatorType),

		ResourceTypes:           make(map[*ResourceType]struct{}),
		ResourceTypes_mapString: make(map[string]*ResourceType),

		Simples:           make(map[*Simple]struct{}),
		Simples_mapString: make(map[string]*Simple),

		TitleEltTypes:           make(map[*TitleEltType]struct{}),
		TitleEltTypes_mapString: make(map[string]*TitleEltType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ArcType_stagedOrder: make(map[*ArcType]uint),
		ArcType_orderStaged: make(map[uint]*ArcType),
		ArcTypes_reference:  make(map[*ArcType]*ArcType),

		Extended_stagedOrder: make(map[*Extended]uint),
		Extended_orderStaged: make(map[uint]*Extended),
		Extendeds_reference:  make(map[*Extended]*Extended),

		LocatorType_stagedOrder: make(map[*LocatorType]uint),
		LocatorType_orderStaged: make(map[uint]*LocatorType),
		LocatorTypes_reference:  make(map[*LocatorType]*LocatorType),

		ResourceType_stagedOrder: make(map[*ResourceType]uint),
		ResourceType_orderStaged: make(map[uint]*ResourceType),
		ResourceTypes_reference:  make(map[*ResourceType]*ResourceType),

		Simple_stagedOrder: make(map[*Simple]uint),
		Simple_orderStaged: make(map[uint]*Simple),
		Simples_reference:  make(map[*Simple]*Simple),

		TitleEltType_stagedOrder: make(map[*TitleEltType]uint),
		TitleEltType_orderStaged: make(map[uint]*TitleEltType),
		TitleEltTypes_reference:  make(map[*TitleEltType]*TitleEltType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"ArcType": &ArcTypeUnmarshaller{},

			"Extended": &ExtendedUnmarshaller{},

			"LocatorType": &LocatorTypeUnmarshaller{},

			"ResourceType": &ResourceTypeUnmarshaller{},

			"Simple": &SimpleUnmarshaller{},

			"TitleEltType": &TitleEltTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["ArcType"] = len(stage.ArcTypes)
	stage.Map_GongStructName_InstancesNb["Extended"] = len(stage.Extendeds)
	stage.Map_GongStructName_InstancesNb["LocatorType"] = len(stage.LocatorTypes)
	stage.Map_GongStructName_InstancesNb["ResourceType"] = len(stage.ResourceTypes)
	stage.Map_GongStructName_InstancesNb["Simple"] = len(stage.Simples)
	stage.Map_GongStructName_InstancesNb["TitleEltType"] = len(stage.TitleEltTypes)
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
// Stage puts arctype to the model stage
func (arctype *ArcType) Stage(stage *Stage) *ArcType {
	__gong__stage(stage.ArcTypes, stage.ArcType_stagedOrder, stage.ArcType_orderStaged, &stage.ArcTypeOrder, stage.ArcTypes_mapString, arctype, arctype.Name)
	return arctype
}

// StagePreserveOrder puts arctype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArcTypeOrder
// - update stage.ArcTypeOrder accordingly
func (arctype *ArcType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArcTypes, stage.ArcType_stagedOrder, stage.ArcType_orderStaged, &stage.ArcTypeOrder, stage.ArcTypes_mapString, arctype, order, arctype.Name)
}

// Unstage removes arctype off the model stage
func (arctype *ArcType) Unstage(stage *Stage) *ArcType {
	__gong__unstage(stage.ArcTypes, stage.ArcTypes_mapString, arctype, arctype.Name)
	return arctype
}

// UnstageVoid removes arctype off the model stage
func (arctype *ArcType) UnstageVoid(stage *Stage) {
	arctype.Unstage(stage)
}

func (arctype *ArcType) StageVoid(stage *Stage) {
	arctype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arctype *ArcType) GetName() (res string) {
	return arctype.Name
}

// for satisfaction of GongStruct interface
func (arctype *ArcType) SetName(name string) {
	arctype.Name = name
}

// Stage puts extended to the model stage
func (extended *Extended) Stage(stage *Stage) *Extended {
	__gong__stage(stage.Extendeds, stage.Extended_stagedOrder, stage.Extended_orderStaged, &stage.ExtendedOrder, stage.Extendeds_mapString, extended, extended.Name)
	return extended
}

// StagePreserveOrder puts extended to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtendedOrder
// - update stage.ExtendedOrder accordingly
func (extended *Extended) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Extendeds, stage.Extended_stagedOrder, stage.Extended_orderStaged, &stage.ExtendedOrder, stage.Extendeds_mapString, extended, order, extended.Name)
}

// Unstage removes extended off the model stage
func (extended *Extended) Unstage(stage *Stage) *Extended {
	__gong__unstage(stage.Extendeds, stage.Extendeds_mapString, extended, extended.Name)
	return extended
}

// UnstageVoid removes extended off the model stage
func (extended *Extended) UnstageVoid(stage *Stage) {
	extended.Unstage(stage)
}

func (extended *Extended) StageVoid(stage *Stage) {
	extended.Stage(stage)
}

// for satisfaction of GongStruct interface
func (extended *Extended) GetName() (res string) {
	return extended.Name
}

// for satisfaction of GongStruct interface
func (extended *Extended) SetName(name string) {
	extended.Name = name
}

// Stage puts locatortype to the model stage
func (locatortype *LocatorType) Stage(stage *Stage) *LocatorType {
	__gong__stage(stage.LocatorTypes, stage.LocatorType_stagedOrder, stage.LocatorType_orderStaged, &stage.LocatorTypeOrder, stage.LocatorTypes_mapString, locatortype, locatortype.Name)
	return locatortype
}

// StagePreserveOrder puts locatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LocatorTypeOrder
// - update stage.LocatorTypeOrder accordingly
func (locatortype *LocatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LocatorTypes, stage.LocatorType_stagedOrder, stage.LocatorType_orderStaged, &stage.LocatorTypeOrder, stage.LocatorTypes_mapString, locatortype, order, locatortype.Name)
}

// Unstage removes locatortype off the model stage
func (locatortype *LocatorType) Unstage(stage *Stage) *LocatorType {
	__gong__unstage(stage.LocatorTypes, stage.LocatorTypes_mapString, locatortype, locatortype.Name)
	return locatortype
}

// UnstageVoid removes locatortype off the model stage
func (locatortype *LocatorType) UnstageVoid(stage *Stage) {
	locatortype.Unstage(stage)
}

func (locatortype *LocatorType) StageVoid(stage *Stage) {
	locatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (locatortype *LocatorType) GetName() (res string) {
	return locatortype.Name
}

// for satisfaction of GongStruct interface
func (locatortype *LocatorType) SetName(name string) {
	locatortype.Name = name
}

// Stage puts resourcetype to the model stage
func (resourcetype *ResourceType) Stage(stage *Stage) *ResourceType {
	__gong__stage(stage.ResourceTypes, stage.ResourceType_stagedOrder, stage.ResourceType_orderStaged, &stage.ResourceTypeOrder, stage.ResourceTypes_mapString, resourcetype, resourcetype.Name)
	return resourcetype
}

// StagePreserveOrder puts resourcetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceTypeOrder
// - update stage.ResourceTypeOrder accordingly
func (resourcetype *ResourceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ResourceTypes, stage.ResourceType_stagedOrder, stage.ResourceType_orderStaged, &stage.ResourceTypeOrder, stage.ResourceTypes_mapString, resourcetype, order, resourcetype.Name)
}

// Unstage removes resourcetype off the model stage
func (resourcetype *ResourceType) Unstage(stage *Stage) *ResourceType {
	__gong__unstage(stage.ResourceTypes, stage.ResourceTypes_mapString, resourcetype, resourcetype.Name)
	return resourcetype
}

// UnstageVoid removes resourcetype off the model stage
func (resourcetype *ResourceType) UnstageVoid(stage *Stage) {
	resourcetype.Unstage(stage)
}

func (resourcetype *ResourceType) StageVoid(stage *Stage) {
	resourcetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (resourcetype *ResourceType) GetName() (res string) {
	return resourcetype.Name
}

// for satisfaction of GongStruct interface
func (resourcetype *ResourceType) SetName(name string) {
	resourcetype.Name = name
}

// Stage puts simple to the model stage
func (simple *Simple) Stage(stage *Stage) *Simple {
	__gong__stage(stage.Simples, stage.Simple_stagedOrder, stage.Simple_orderStaged, &stage.SimpleOrder, stage.Simples_mapString, simple, simple.Name)
	return simple
}

// StagePreserveOrder puts simple to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SimpleOrder
// - update stage.SimpleOrder accordingly
func (simple *Simple) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Simples, stage.Simple_stagedOrder, stage.Simple_orderStaged, &stage.SimpleOrder, stage.Simples_mapString, simple, order, simple.Name)
}

// Unstage removes simple off the model stage
func (simple *Simple) Unstage(stage *Stage) *Simple {
	__gong__unstage(stage.Simples, stage.Simples_mapString, simple, simple.Name)
	return simple
}

// UnstageVoid removes simple off the model stage
func (simple *Simple) UnstageVoid(stage *Stage) {
	simple.Unstage(stage)
}

func (simple *Simple) StageVoid(stage *Stage) {
	simple.Stage(stage)
}

// for satisfaction of GongStruct interface
func (simple *Simple) GetName() (res string) {
	return simple.Name
}

// for satisfaction of GongStruct interface
func (simple *Simple) SetName(name string) {
	simple.Name = name
}

// Stage puts titleelttype to the model stage
func (titleelttype *TitleEltType) Stage(stage *Stage) *TitleEltType {
	__gong__stage(stage.TitleEltTypes, stage.TitleEltType_stagedOrder, stage.TitleEltType_orderStaged, &stage.TitleEltTypeOrder, stage.TitleEltTypes_mapString, titleelttype, titleelttype.Name)
	return titleelttype
}

// StagePreserveOrder puts titleelttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TitleEltTypeOrder
// - update stage.TitleEltTypeOrder accordingly
func (titleelttype *TitleEltType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TitleEltTypes, stage.TitleEltType_stagedOrder, stage.TitleEltType_orderStaged, &stage.TitleEltTypeOrder, stage.TitleEltTypes_mapString, titleelttype, order, titleelttype.Name)
}

// Unstage removes titleelttype off the model stage
func (titleelttype *TitleEltType) Unstage(stage *Stage) *TitleEltType {
	__gong__unstage(stage.TitleEltTypes, stage.TitleEltTypes_mapString, titleelttype, titleelttype.Name)
	return titleelttype
}

// UnstageVoid removes titleelttype off the model stage
func (titleelttype *TitleEltType) UnstageVoid(stage *Stage) {
	titleelttype.Unstage(stage)
}

func (titleelttype *TitleEltType) StageVoid(stage *Stage) {
	titleelttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (titleelttype *TitleEltType) GetName() (res string) {
	return titleelttype.Name
}

// for satisfaction of GongStruct interface
func (titleelttype *TitleEltType) SetName(name string) {
	titleelttype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.ArcTypes, &stage.ArcTypes_mapString, &stage.ArcType_stagedOrder, &stage.ArcTypeOrder)

	__gong__resetStageType(&stage.Extendeds, &stage.Extendeds_mapString, &stage.Extended_stagedOrder, &stage.ExtendedOrder)

	__gong__resetStageType(&stage.LocatorTypes, &stage.LocatorTypes_mapString, &stage.LocatorType_stagedOrder, &stage.LocatorTypeOrder)

	__gong__resetStageType(&stage.ResourceTypes, &stage.ResourceTypes_mapString, &stage.ResourceType_stagedOrder, &stage.ResourceTypeOrder)

	__gong__resetStageType(&stage.Simples, &stage.Simples_mapString, &stage.Simple_stagedOrder, &stage.SimpleOrder)

	__gong__resetStageType(&stage.TitleEltTypes, &stage.TitleEltTypes_mapString, &stage.TitleEltType_stagedOrder, &stage.TitleEltTypeOrder)

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
func (ArcType) GongGetAssociationName() any {
	return &ArcType{
			Title_1: []*TitleEltType{{Name: "Title_1"}},
	}
}

func (Extended) GongGetAssociationName() any {
	return &Extended{
			Title_1: &TitleEltType{Name: "Title_1"},
			Resource: &ResourceType{Name: "Resource"},
			Locator: &LocatorType{Name: "Locator"},
			Arc: &ArcType{Name: "Arc"},
	}
}

func (LocatorType) GongGetAssociationName() any {
	return &LocatorType{
			Title_1: []*TitleEltType{{Name: "Title_1"}},
	}
}

func (ResourceType) GongGetAssociationName() any {
	return &ResourceType{
	}
}

func (Simple) GongGetAssociationName() any {
	return &Simple{
	}
}

func (TitleEltType) GongGetAssociationName() any {
	return &TitleEltType{
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
	// reverse maps of direct associations of ArcType
	case ArcType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Extended
	case Extended:
		switch fieldname {
		// insertion point for per direct association field
		case "Title_1":
			res := make(map[*TitleEltType][]*Extended)
			for extended := range stage.Extendeds {
				if extended.Title_1 != nil {
					titleelttype_ := extended.Title_1
					var extendeds []*Extended
					_, ok := res[titleelttype_]
					if ok {
						extendeds = res[titleelttype_]
					} else {
						extendeds = make([]*Extended, 0)
					}
					extendeds = append(extendeds, extended)
					res[titleelttype_] = extendeds
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resource":
			res := make(map[*ResourceType][]*Extended)
			for extended := range stage.Extendeds {
				if extended.Resource != nil {
					resourcetype_ := extended.Resource
					var extendeds []*Extended
					_, ok := res[resourcetype_]
					if ok {
						extendeds = res[resourcetype_]
					} else {
						extendeds = make([]*Extended, 0)
					}
					extendeds = append(extendeds, extended)
					res[resourcetype_] = extendeds
				}
			}
			return any(res).(map[*End][]*Start)
		case "Locator":
			res := make(map[*LocatorType][]*Extended)
			for extended := range stage.Extendeds {
				if extended.Locator != nil {
					locatortype_ := extended.Locator
					var extendeds []*Extended
					_, ok := res[locatortype_]
					if ok {
						extendeds = res[locatortype_]
					} else {
						extendeds = make([]*Extended, 0)
					}
					extendeds = append(extendeds, extended)
					res[locatortype_] = extendeds
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arc":
			res := make(map[*ArcType][]*Extended)
			for extended := range stage.Extendeds {
				if extended.Arc != nil {
					arctype_ := extended.Arc
					var extendeds []*Extended
					_, ok := res[arctype_]
					if ok {
						extendeds = res[arctype_]
					} else {
						extendeds = make([]*Extended, 0)
					}
					extendeds = append(extendeds, extended)
					res[arctype_] = extendeds
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LocatorType
	case LocatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ResourceType
	case ResourceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Simple
	case Simple:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TitleEltType
	case TitleEltType:
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
	// reverse maps of direct associations of ArcType
	case ArcType:
		switch fieldname {
		// insertion point for per direct association field
		case "Title_1":
			res := make(map[*TitleEltType][]*ArcType)
			for arctype := range stage.ArcTypes {
				for _, titleelttype_ := range arctype.Title_1 {
					res[titleelttype_] = append(res[titleelttype_], arctype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Extended
	case Extended:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LocatorType
	case LocatorType:
		switch fieldname {
		// insertion point for per direct association field
		case "Title_1":
			res := make(map[*TitleEltType][]*LocatorType)
			for locatortype := range stage.LocatorTypes {
				for _, titleelttype_ := range locatortype.Title_1 {
					res[titleelttype_] = append(res[titleelttype_], locatortype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ResourceType
	case ResourceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Simple
	case Simple:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TitleEltType
	case TitleEltType:
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
func (*ArcType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Extended) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LocatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ResourceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Simple) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TitleEltType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ArcType",
			Fieldname: "Title_1",
		},
		{
			GongstructName: "LocatorType",
			Fieldname: "Title_1",
		},
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
func (arctype *ArcType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Arcrole",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Show",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Actuate",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "From",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "To",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Title_1",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TitleEltType",
		},
	}
	return
}

func (extended *Extended) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Title_1",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TitleEltType",
		},
		{
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ResourceType",
		},
		{
			Name:                 "Locator",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LocatorType",
		},
		{
			Name:                 "Arc",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ArcType",
		},
	}
	return
}

func (locatortype *LocatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Title_1",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TitleEltType",
		},
	}
	return
}

func (resourcetype *ResourceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (simple *Simple) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Arcrole",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Show",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Actuate",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (titleelttype *TitleEltType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Lang",
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
func (arctype *ArcType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arctype.Name
	case "Type_XSD":
		res.valueString = arctype.Type_XSD
	case "Arcrole":
		res.valueString = arctype.Arcrole
	case "Title":
		res.valueString = arctype.Title
	case "Show":
		res.valueString = arctype.Show
	case "Actuate":
		res.valueString = arctype.Actuate
	case "From":
		res.valueString = arctype.From
	case "To":
		res.valueString = arctype.To
	case "Title_1":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range arctype.Title_1 {
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

func (extended *Extended) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = extended.Name
	case "Type_XSD":
		res.valueString = extended.Type_XSD
	case "Role":
		res.valueString = extended.Role
	case "Title":
		res.valueString = extended.Title
	case "Title_1":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extended.Title_1 != nil {
			res.valueString = extended.Title_1.Name
			res.ids = extended.Title_1.GongGetUUID(stage)
		}
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extended.Resource != nil {
			res.valueString = extended.Resource.Name
			res.ids = extended.Resource.GongGetUUID(stage)
		}
	case "Locator":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extended.Locator != nil {
			res.valueString = extended.Locator.Name
			res.ids = extended.Locator.GongGetUUID(stage)
		}
	case "Arc":
		res.GongFieldValueType = GongFieldValueTypePointer
		if extended.Arc != nil {
			res.valueString = extended.Arc.Name
			res.ids = extended.Arc.GongGetUUID(stage)
		}
	}
	return
}

func (locatortype *LocatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = locatortype.Name
	case "Type_XSD":
		res.valueString = locatortype.Type_XSD
	case "Href":
		res.valueString = locatortype.Href
	case "Role":
		res.valueString = locatortype.Role
	case "Title":
		res.valueString = locatortype.Title
	case "Label":
		res.valueString = locatortype.Label
	case "Title_1":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range locatortype.Title_1 {
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

func (resourcetype *ResourceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resourcetype.Name
	case "Type_XSD":
		res.valueString = resourcetype.Type_XSD
	case "Role":
		res.valueString = resourcetype.Role
	case "Title":
		res.valueString = resourcetype.Title
	case "Label":
		res.valueString = resourcetype.Label
	}
	return
}

func (simple *Simple) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = simple.Name
	case "Type_XSD":
		res.valueString = simple.Type_XSD
	case "Href":
		res.valueString = simple.Href
	case "Role":
		res.valueString = simple.Role
	case "Arcrole":
		res.valueString = simple.Arcrole
	case "Title":
		res.valueString = simple.Title
	case "Show":
		res.valueString = simple.Show
	case "Actuate":
		res.valueString = simple.Actuate
	}
	return
}

func (titleelttype *TitleEltType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = titleelttype.Name
	case "Type_XSD":
		res.valueString = titleelttype.Type_XSD
	case "Lang":
		res.valueString = titleelttype.Lang
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
func (arctype *ArcType) GongGetGongstructName() string {
	return "ArcType"
}

func (extended *Extended) GongGetGongstructName() string {
	return "Extended"
}

func (locatortype *LocatorType) GongGetGongstructName() string {
	return "LocatorType"
}

func (resourcetype *ResourceType) GongGetGongstructName() string {
	return "ResourceType"
}

func (simple *Simple) GongGetGongstructName() string {
	return "Simple"
}

func (titleelttype *TitleEltType) GongGetGongstructName() string {
	return "TitleEltType"
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
	__gong__rebuildMapString(stage.ArcTypes, &stage.ArcTypes_mapString)

	__gong__rebuildMapString(stage.Extendeds, &stage.Extendeds_mapString)

	__gong__rebuildMapString(stage.LocatorTypes, &stage.LocatorTypes_mapString)

	__gong__rebuildMapString(stage.ResourceTypes, &stage.ResourceTypes_mapString)

	__gong__rebuildMapString(stage.Simples, &stage.Simples_mapString)

	__gong__rebuildMapString(stage.TitleEltTypes, &stage.TitleEltTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
