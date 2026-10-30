// generated code - do not edit
package measures

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
	AngleTypes                map[*AngleType]struct{}
	AngleTypes_instance       map[*AngleType]*AngleType
	AngleTypes_mapString      map[string]*AngleType
	AngleTypeOrder            uint
	AngleType_stagedOrder     map[*AngleType]uint
	AngleType_orderStaged     map[uint]*AngleType
	AngleTypes_reference      map[*AngleType]*AngleType
	AngleTypes_referenceOrder map[*AngleType]uint

	// insertion point for slice of pointers maps
	OnAfterAngleTypeCreateCallback GongOnAfterCreateInterface[AngleType]
	OnAfterAngleTypeUpdateCallback GongOnAfterUpdateInterface[AngleType]
	OnAfterAngleTypeDeleteCallback GongOnAfterDeleteInterface[AngleType]

	AreaTypes                map[*AreaType]struct{}
	AreaTypes_instance       map[*AreaType]*AreaType
	AreaTypes_mapString      map[string]*AreaType
	AreaTypeOrder            uint
	AreaType_stagedOrder     map[*AreaType]uint
	AreaType_orderStaged     map[uint]*AreaType
	AreaTypes_reference      map[*AreaType]*AreaType
	AreaTypes_referenceOrder map[*AreaType]uint

	// insertion point for slice of pointers maps
	OnAfterAreaTypeCreateCallback GongOnAfterCreateInterface[AreaType]
	OnAfterAreaTypeUpdateCallback GongOnAfterUpdateInterface[AreaType]
	OnAfterAreaTypeDeleteCallback GongOnAfterDeleteInterface[AreaType]

	GridLengthTypes                map[*GridLengthType]struct{}
	GridLengthTypes_instance       map[*GridLengthType]*GridLengthType
	GridLengthTypes_mapString      map[string]*GridLengthType
	GridLengthTypeOrder            uint
	GridLengthType_stagedOrder     map[*GridLengthType]uint
	GridLengthType_orderStaged     map[uint]*GridLengthType
	GridLengthTypes_reference      map[*GridLengthType]*GridLengthType
	GridLengthTypes_referenceOrder map[*GridLengthType]uint

	// insertion point for slice of pointers maps
	OnAfterGridLengthTypeCreateCallback GongOnAfterCreateInterface[GridLengthType]
	OnAfterGridLengthTypeUpdateCallback GongOnAfterUpdateInterface[GridLengthType]
	OnAfterGridLengthTypeDeleteCallback GongOnAfterDeleteInterface[GridLengthType]

	LengthTypes                map[*LengthType]struct{}
	LengthTypes_instance       map[*LengthType]*LengthType
	LengthTypes_mapString      map[string]*LengthType
	LengthTypeOrder            uint
	LengthType_stagedOrder     map[*LengthType]uint
	LengthType_orderStaged     map[uint]*LengthType
	LengthTypes_reference      map[*LengthType]*LengthType
	LengthTypes_referenceOrder map[*LengthType]uint

	// insertion point for slice of pointers maps
	OnAfterLengthTypeCreateCallback GongOnAfterCreateInterface[LengthType]
	OnAfterLengthTypeUpdateCallback GongOnAfterUpdateInterface[LengthType]
	OnAfterLengthTypeDeleteCallback GongOnAfterDeleteInterface[LengthType]

	ScaleTypes                map[*ScaleType]struct{}
	ScaleTypes_instance       map[*ScaleType]*ScaleType
	ScaleTypes_mapString      map[string]*ScaleType
	ScaleTypeOrder            uint
	ScaleType_stagedOrder     map[*ScaleType]uint
	ScaleType_orderStaged     map[uint]*ScaleType
	ScaleTypes_reference      map[*ScaleType]*ScaleType
	ScaleTypes_referenceOrder map[*ScaleType]uint

	// insertion point for slice of pointers maps
	OnAfterScaleTypeCreateCallback GongOnAfterCreateInterface[ScaleType]
	OnAfterScaleTypeUpdateCallback GongOnAfterUpdateInterface[ScaleType]
	OnAfterScaleTypeDeleteCallback GongOnAfterDeleteInterface[ScaleType]

	SpeedTypes                map[*SpeedType]struct{}
	SpeedTypes_instance       map[*SpeedType]*SpeedType
	SpeedTypes_mapString      map[string]*SpeedType
	SpeedTypeOrder            uint
	SpeedType_stagedOrder     map[*SpeedType]uint
	SpeedType_orderStaged     map[uint]*SpeedType
	SpeedTypes_reference      map[*SpeedType]*SpeedType
	SpeedTypes_referenceOrder map[*SpeedType]uint

	// insertion point for slice of pointers maps
	OnAfterSpeedTypeCreateCallback GongOnAfterCreateInterface[SpeedType]
	OnAfterSpeedTypeUpdateCallback GongOnAfterUpdateInterface[SpeedType]
	OnAfterSpeedTypeDeleteCallback GongOnAfterDeleteInterface[SpeedType]

	TimeTypes                map[*TimeType]struct{}
	TimeTypes_instance       map[*TimeType]*TimeType
	TimeTypes_mapString      map[string]*TimeType
	TimeTypeOrder            uint
	TimeType_stagedOrder     map[*TimeType]uint
	TimeType_orderStaged     map[uint]*TimeType
	TimeTypes_reference      map[*TimeType]*TimeType
	TimeTypes_referenceOrder map[*TimeType]uint

	// insertion point for slice of pointers maps
	OnAfterTimeTypeCreateCallback GongOnAfterCreateInterface[TimeType]
	OnAfterTimeTypeUpdateCallback GongOnAfterUpdateInterface[TimeType]
	OnAfterTimeTypeDeleteCallback GongOnAfterDeleteInterface[TimeType]

	VolumeTypes                map[*VolumeType]struct{}
	VolumeTypes_instance       map[*VolumeType]*VolumeType
	VolumeTypes_mapString      map[string]*VolumeType
	VolumeTypeOrder            uint
	VolumeType_stagedOrder     map[*VolumeType]uint
	VolumeType_orderStaged     map[uint]*VolumeType
	VolumeTypes_reference      map[*VolumeType]*VolumeType
	VolumeTypes_referenceOrder map[*VolumeType]uint

	// insertion point for slice of pointers maps
	OnAfterVolumeTypeCreateCallback GongOnAfterCreateInterface[VolumeType]
	OnAfterVolumeTypeUpdateCallback GongOnAfterUpdateInterface[VolumeType]
	OnAfterVolumeTypeDeleteCallback GongOnAfterDeleteInterface[VolumeType]

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
	__gong__clearReferences(&stage.AngleTypes_reference, &stage.AngleTypes_instance, &stage.AngleTypes_referenceOrder)

	__gong__clearReferences(&stage.AreaTypes_reference, &stage.AreaTypes_instance, &stage.AreaTypes_referenceOrder)

	__gong__clearReferences(&stage.GridLengthTypes_reference, &stage.GridLengthTypes_instance, &stage.GridLengthTypes_referenceOrder)

	__gong__clearReferences(&stage.LengthTypes_reference, &stage.LengthTypes_instance, &stage.LengthTypes_referenceOrder)

	__gong__clearReferences(&stage.ScaleTypes_reference, &stage.ScaleTypes_instance, &stage.ScaleTypes_referenceOrder)

	__gong__clearReferences(&stage.SpeedTypes_reference, &stage.SpeedTypes_instance, &stage.SpeedTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.VolumeTypes_reference, &stage.VolumeTypes_instance, &stage.VolumeTypes_referenceOrder)

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
	stage.AngleTypeOrder = __gong__recomputeOrder(stage.AngleType_stagedOrder)

	stage.AreaTypeOrder = __gong__recomputeOrder(stage.AreaType_stagedOrder)

	stage.GridLengthTypeOrder = __gong__recomputeOrder(stage.GridLengthType_stagedOrder)

	stage.LengthTypeOrder = __gong__recomputeOrder(stage.LengthType_stagedOrder)

	stage.ScaleTypeOrder = __gong__recomputeOrder(stage.ScaleType_stagedOrder)

	stage.SpeedTypeOrder = __gong__recomputeOrder(stage.SpeedType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.VolumeTypeOrder = __gong__recomputeOrder(stage.VolumeType_stagedOrder)

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
func (*AngleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AngleTypes, stage.AngleType_stagedOrder)
}

func (*AngleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AngleType_orderStaged[order]
}

func (*AngleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AngleTypes_mapString
}

func (*AngleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AngleTypes
}

func (*AngleType) GongNewInstance() any {
	return new(AngleType)
}

func (*AreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AreaTypes, stage.AreaType_stagedOrder)
}

func (*AreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AreaType_orderStaged[order]
}

func (*AreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AreaTypes_mapString
}

func (*AreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AreaTypes
}

func (*AreaType) GongNewInstance() any {
	return new(AreaType)
}

func (*GridLengthType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GridLengthTypes, stage.GridLengthType_stagedOrder)
}

func (*GridLengthType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GridLengthType_orderStaged[order]
}

func (*GridLengthType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GridLengthTypes_mapString
}

func (*GridLengthType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GridLengthTypes
}

func (*GridLengthType) GongNewInstance() any {
	return new(GridLengthType)
}

func (*LengthType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LengthTypes, stage.LengthType_stagedOrder)
}

func (*LengthType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LengthType_orderStaged[order]
}

func (*LengthType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LengthTypes_mapString
}

func (*LengthType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LengthTypes
}

func (*LengthType) GongNewInstance() any {
	return new(LengthType)
}

func (*ScaleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ScaleTypes, stage.ScaleType_stagedOrder)
}

func (*ScaleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ScaleType_orderStaged[order]
}

func (*ScaleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ScaleTypes_mapString
}

func (*ScaleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ScaleTypes
}

func (*ScaleType) GongNewInstance() any {
	return new(ScaleType)
}

func (*SpeedType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SpeedTypes, stage.SpeedType_stagedOrder)
}

func (*SpeedType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SpeedType_orderStaged[order]
}

func (*SpeedType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SpeedTypes_mapString
}

func (*SpeedType) GongGetInstancesSet(stage *Stage) any {
	return &stage.SpeedTypes
}

func (*SpeedType) GongNewInstance() any {
	return new(SpeedType)
}

func (*TimeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimeTypes, stage.TimeType_stagedOrder)
}

func (*TimeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimeType_orderStaged[order]
}

func (*TimeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimeTypes_mapString
}

func (*TimeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimeTypes
}

func (*TimeType) GongNewInstance() any {
	return new(TimeType)
}

func (*VolumeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VolumeTypes, stage.VolumeType_stagedOrder)
}

func (*VolumeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VolumeType_orderStaged[order]
}

func (*VolumeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VolumeTypes_mapString
}

func (*VolumeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VolumeTypes
}

func (*VolumeType) GongNewInstance() any {
	return new(VolumeType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/measures"
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
		AngleTypes:           make(map[*AngleType]struct{}),
		AngleTypes_mapString: make(map[string]*AngleType),

		AreaTypes:           make(map[*AreaType]struct{}),
		AreaTypes_mapString: make(map[string]*AreaType),

		GridLengthTypes:           make(map[*GridLengthType]struct{}),
		GridLengthTypes_mapString: make(map[string]*GridLengthType),

		LengthTypes:           make(map[*LengthType]struct{}),
		LengthTypes_mapString: make(map[string]*LengthType),

		ScaleTypes:           make(map[*ScaleType]struct{}),
		ScaleTypes_mapString: make(map[string]*ScaleType),

		SpeedTypes:           make(map[*SpeedType]struct{}),
		SpeedTypes_mapString: make(map[string]*SpeedType),

		TimeTypes:           make(map[*TimeType]struct{}),
		TimeTypes_mapString: make(map[string]*TimeType),

		VolumeTypes:           make(map[*VolumeType]struct{}),
		VolumeTypes_mapString: make(map[string]*VolumeType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AngleType_stagedOrder: make(map[*AngleType]uint),
		AngleType_orderStaged: make(map[uint]*AngleType),
		AngleTypes_reference:  make(map[*AngleType]*AngleType),

		AreaType_stagedOrder: make(map[*AreaType]uint),
		AreaType_orderStaged: make(map[uint]*AreaType),
		AreaTypes_reference:  make(map[*AreaType]*AreaType),

		GridLengthType_stagedOrder: make(map[*GridLengthType]uint),
		GridLengthType_orderStaged: make(map[uint]*GridLengthType),
		GridLengthTypes_reference:  make(map[*GridLengthType]*GridLengthType),

		LengthType_stagedOrder: make(map[*LengthType]uint),
		LengthType_orderStaged: make(map[uint]*LengthType),
		LengthTypes_reference:  make(map[*LengthType]*LengthType),

		ScaleType_stagedOrder: make(map[*ScaleType]uint),
		ScaleType_orderStaged: make(map[uint]*ScaleType),
		ScaleTypes_reference:  make(map[*ScaleType]*ScaleType),

		SpeedType_stagedOrder: make(map[*SpeedType]uint),
		SpeedType_orderStaged: make(map[uint]*SpeedType),
		SpeedTypes_reference:  make(map[*SpeedType]*SpeedType),

		TimeType_stagedOrder: make(map[*TimeType]uint),
		TimeType_orderStaged: make(map[uint]*TimeType),
		TimeTypes_reference:  make(map[*TimeType]*TimeType),

		VolumeType_stagedOrder: make(map[*VolumeType]uint),
		VolumeType_orderStaged: make(map[uint]*VolumeType),
		VolumeTypes_reference:  make(map[*VolumeType]*VolumeType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AngleType": &AngleTypeUnmarshaller{},

			"AreaType": &AreaTypeUnmarshaller{},

			"GridLengthType": &GridLengthTypeUnmarshaller{},

			"LengthType": &LengthTypeUnmarshaller{},

			"ScaleType": &ScaleTypeUnmarshaller{},

			"SpeedType": &SpeedTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"VolumeType": &VolumeTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AngleType"] = len(stage.AngleTypes)
	stage.Map_GongStructName_InstancesNb["AreaType"] = len(stage.AreaTypes)
	stage.Map_GongStructName_InstancesNb["GridLengthType"] = len(stage.GridLengthTypes)
	stage.Map_GongStructName_InstancesNb["LengthType"] = len(stage.LengthTypes)
	stage.Map_GongStructName_InstancesNb["ScaleType"] = len(stage.ScaleTypes)
	stage.Map_GongStructName_InstancesNb["SpeedType"] = len(stage.SpeedTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["VolumeType"] = len(stage.VolumeTypes)
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
// Stage puts angletype to the model stage
func (angletype *AngleType) Stage(stage *Stage) *AngleType {
	__gong__stage(stage.AngleTypes, stage.AngleType_stagedOrder, stage.AngleType_orderStaged, &stage.AngleTypeOrder, stage.AngleTypes_mapString, angletype, angletype.Name)
	return angletype
}

// StagePreserveOrder puts angletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AngleTypeOrder
// - update stage.AngleTypeOrder accordingly
func (angletype *AngleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AngleTypes, stage.AngleType_stagedOrder, stage.AngleType_orderStaged, &stage.AngleTypeOrder, stage.AngleTypes_mapString, angletype, order, angletype.Name)
}

// Unstage removes angletype off the model stage
func (angletype *AngleType) Unstage(stage *Stage) *AngleType {
	__gong__unstage(stage.AngleTypes, stage.AngleTypes_mapString, angletype, angletype.Name)
	return angletype
}

// UnstageVoid removes angletype off the model stage
func (angletype *AngleType) UnstageVoid(stage *Stage) {
	angletype.Unstage(stage)
}

func (angletype *AngleType) StageVoid(stage *Stage) {
	angletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (angletype *AngleType) GetName() (res string) {
	return angletype.Name
}

// for satisfaction of GongStruct interface
func (angletype *AngleType) SetName(name string) {
	angletype.Name = name
}

// Stage puts areatype to the model stage
func (areatype *AreaType) Stage(stage *Stage) *AreaType {
	__gong__stage(stage.AreaTypes, stage.AreaType_stagedOrder, stage.AreaType_orderStaged, &stage.AreaTypeOrder, stage.AreaTypes_mapString, areatype, areatype.Name)
	return areatype
}

// StagePreserveOrder puts areatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AreaTypeOrder
// - update stage.AreaTypeOrder accordingly
func (areatype *AreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AreaTypes, stage.AreaType_stagedOrder, stage.AreaType_orderStaged, &stage.AreaTypeOrder, stage.AreaTypes_mapString, areatype, order, areatype.Name)
}

// Unstage removes areatype off the model stage
func (areatype *AreaType) Unstage(stage *Stage) *AreaType {
	__gong__unstage(stage.AreaTypes, stage.AreaTypes_mapString, areatype, areatype.Name)
	return areatype
}

// UnstageVoid removes areatype off the model stage
func (areatype *AreaType) UnstageVoid(stage *Stage) {
	areatype.Unstage(stage)
}

func (areatype *AreaType) StageVoid(stage *Stage) {
	areatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (areatype *AreaType) GetName() (res string) {
	return areatype.Name
}

// for satisfaction of GongStruct interface
func (areatype *AreaType) SetName(name string) {
	areatype.Name = name
}

// Stage puts gridlengthtype to the model stage
func (gridlengthtype *GridLengthType) Stage(stage *Stage) *GridLengthType {
	__gong__stage(stage.GridLengthTypes, stage.GridLengthType_stagedOrder, stage.GridLengthType_orderStaged, &stage.GridLengthTypeOrder, stage.GridLengthTypes_mapString, gridlengthtype, gridlengthtype.Name)
	return gridlengthtype
}

// StagePreserveOrder puts gridlengthtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GridLengthTypeOrder
// - update stage.GridLengthTypeOrder accordingly
func (gridlengthtype *GridLengthType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GridLengthTypes, stage.GridLengthType_stagedOrder, stage.GridLengthType_orderStaged, &stage.GridLengthTypeOrder, stage.GridLengthTypes_mapString, gridlengthtype, order, gridlengthtype.Name)
}

// Unstage removes gridlengthtype off the model stage
func (gridlengthtype *GridLengthType) Unstage(stage *Stage) *GridLengthType {
	__gong__unstage(stage.GridLengthTypes, stage.GridLengthTypes_mapString, gridlengthtype, gridlengthtype.Name)
	return gridlengthtype
}

// UnstageVoid removes gridlengthtype off the model stage
func (gridlengthtype *GridLengthType) UnstageVoid(stage *Stage) {
	gridlengthtype.Unstage(stage)
}

func (gridlengthtype *GridLengthType) StageVoid(stage *Stage) {
	gridlengthtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (gridlengthtype *GridLengthType) GetName() (res string) {
	return gridlengthtype.Name
}

// for satisfaction of GongStruct interface
func (gridlengthtype *GridLengthType) SetName(name string) {
	gridlengthtype.Name = name
}

// Stage puts lengthtype to the model stage
func (lengthtype *LengthType) Stage(stage *Stage) *LengthType {
	__gong__stage(stage.LengthTypes, stage.LengthType_stagedOrder, stage.LengthType_orderStaged, &stage.LengthTypeOrder, stage.LengthTypes_mapString, lengthtype, lengthtype.Name)
	return lengthtype
}

// StagePreserveOrder puts lengthtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LengthTypeOrder
// - update stage.LengthTypeOrder accordingly
func (lengthtype *LengthType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LengthTypes, stage.LengthType_stagedOrder, stage.LengthType_orderStaged, &stage.LengthTypeOrder, stage.LengthTypes_mapString, lengthtype, order, lengthtype.Name)
}

// Unstage removes lengthtype off the model stage
func (lengthtype *LengthType) Unstage(stage *Stage) *LengthType {
	__gong__unstage(stage.LengthTypes, stage.LengthTypes_mapString, lengthtype, lengthtype.Name)
	return lengthtype
}

// UnstageVoid removes lengthtype off the model stage
func (lengthtype *LengthType) UnstageVoid(stage *Stage) {
	lengthtype.Unstage(stage)
}

func (lengthtype *LengthType) StageVoid(stage *Stage) {
	lengthtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lengthtype *LengthType) GetName() (res string) {
	return lengthtype.Name
}

// for satisfaction of GongStruct interface
func (lengthtype *LengthType) SetName(name string) {
	lengthtype.Name = name
}

// Stage puts scaletype to the model stage
func (scaletype *ScaleType) Stage(stage *Stage) *ScaleType {
	__gong__stage(stage.ScaleTypes, stage.ScaleType_stagedOrder, stage.ScaleType_orderStaged, &stage.ScaleTypeOrder, stage.ScaleTypes_mapString, scaletype, scaletype.Name)
	return scaletype
}

// StagePreserveOrder puts scaletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScaleTypeOrder
// - update stage.ScaleTypeOrder accordingly
func (scaletype *ScaleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ScaleTypes, stage.ScaleType_stagedOrder, stage.ScaleType_orderStaged, &stage.ScaleTypeOrder, stage.ScaleTypes_mapString, scaletype, order, scaletype.Name)
}

// Unstage removes scaletype off the model stage
func (scaletype *ScaleType) Unstage(stage *Stage) *ScaleType {
	__gong__unstage(stage.ScaleTypes, stage.ScaleTypes_mapString, scaletype, scaletype.Name)
	return scaletype
}

// UnstageVoid removes scaletype off the model stage
func (scaletype *ScaleType) UnstageVoid(stage *Stage) {
	scaletype.Unstage(stage)
}

func (scaletype *ScaleType) StageVoid(stage *Stage) {
	scaletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scaletype *ScaleType) GetName() (res string) {
	return scaletype.Name
}

// for satisfaction of GongStruct interface
func (scaletype *ScaleType) SetName(name string) {
	scaletype.Name = name
}

// Stage puts speedtype to the model stage
func (speedtype *SpeedType) Stage(stage *Stage) *SpeedType {
	__gong__stage(stage.SpeedTypes, stage.SpeedType_stagedOrder, stage.SpeedType_orderStaged, &stage.SpeedTypeOrder, stage.SpeedTypes_mapString, speedtype, speedtype.Name)
	return speedtype
}

// StagePreserveOrder puts speedtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpeedTypeOrder
// - update stage.SpeedTypeOrder accordingly
func (speedtype *SpeedType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SpeedTypes, stage.SpeedType_stagedOrder, stage.SpeedType_orderStaged, &stage.SpeedTypeOrder, stage.SpeedTypes_mapString, speedtype, order, speedtype.Name)
}

// Unstage removes speedtype off the model stage
func (speedtype *SpeedType) Unstage(stage *Stage) *SpeedType {
	__gong__unstage(stage.SpeedTypes, stage.SpeedTypes_mapString, speedtype, speedtype.Name)
	return speedtype
}

// UnstageVoid removes speedtype off the model stage
func (speedtype *SpeedType) UnstageVoid(stage *Stage) {
	speedtype.Unstage(stage)
}

func (speedtype *SpeedType) StageVoid(stage *Stage) {
	speedtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (speedtype *SpeedType) GetName() (res string) {
	return speedtype.Name
}

// for satisfaction of GongStruct interface
func (speedtype *SpeedType) SetName(name string) {
	speedtype.Name = name
}

// Stage puts timetype to the model stage
func (timetype *TimeType) Stage(stage *Stage) *TimeType {
	__gong__stage(stage.TimeTypes, stage.TimeType_stagedOrder, stage.TimeType_orderStaged, &stage.TimeTypeOrder, stage.TimeTypes_mapString, timetype, timetype.Name)
	return timetype
}

// StagePreserveOrder puts timetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimeTypeOrder
// - update stage.TimeTypeOrder accordingly
func (timetype *TimeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimeTypes, stage.TimeType_stagedOrder, stage.TimeType_orderStaged, &stage.TimeTypeOrder, stage.TimeTypes_mapString, timetype, order, timetype.Name)
}

// Unstage removes timetype off the model stage
func (timetype *TimeType) Unstage(stage *Stage) *TimeType {
	__gong__unstage(stage.TimeTypes, stage.TimeTypes_mapString, timetype, timetype.Name)
	return timetype
}

// UnstageVoid removes timetype off the model stage
func (timetype *TimeType) UnstageVoid(stage *Stage) {
	timetype.Unstage(stage)
}

func (timetype *TimeType) StageVoid(stage *Stage) {
	timetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timetype *TimeType) GetName() (res string) {
	return timetype.Name
}

// for satisfaction of GongStruct interface
func (timetype *TimeType) SetName(name string) {
	timetype.Name = name
}

// Stage puts volumetype to the model stage
func (volumetype *VolumeType) Stage(stage *Stage) *VolumeType {
	__gong__stage(stage.VolumeTypes, stage.VolumeType_stagedOrder, stage.VolumeType_orderStaged, &stage.VolumeTypeOrder, stage.VolumeTypes_mapString, volumetype, volumetype.Name)
	return volumetype
}

// StagePreserveOrder puts volumetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VolumeTypeOrder
// - update stage.VolumeTypeOrder accordingly
func (volumetype *VolumeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VolumeTypes, stage.VolumeType_stagedOrder, stage.VolumeType_orderStaged, &stage.VolumeTypeOrder, stage.VolumeTypes_mapString, volumetype, order, volumetype.Name)
}

// Unstage removes volumetype off the model stage
func (volumetype *VolumeType) Unstage(stage *Stage) *VolumeType {
	__gong__unstage(stage.VolumeTypes, stage.VolumeTypes_mapString, volumetype, volumetype.Name)
	return volumetype
}

// UnstageVoid removes volumetype off the model stage
func (volumetype *VolumeType) UnstageVoid(stage *Stage) {
	volumetype.Unstage(stage)
}

func (volumetype *VolumeType) StageVoid(stage *Stage) {
	volumetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (volumetype *VolumeType) GetName() (res string) {
	return volumetype.Name
}

// for satisfaction of GongStruct interface
func (volumetype *VolumeType) SetName(name string) {
	volumetype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AngleTypes, &stage.AngleTypes_mapString, &stage.AngleType_stagedOrder, &stage.AngleTypeOrder)

	__gong__resetStageType(&stage.AreaTypes, &stage.AreaTypes_mapString, &stage.AreaType_stagedOrder, &stage.AreaTypeOrder)

	__gong__resetStageType(&stage.GridLengthTypes, &stage.GridLengthTypes_mapString, &stage.GridLengthType_stagedOrder, &stage.GridLengthTypeOrder)

	__gong__resetStageType(&stage.LengthTypes, &stage.LengthTypes_mapString, &stage.LengthType_stagedOrder, &stage.LengthTypeOrder)

	__gong__resetStageType(&stage.ScaleTypes, &stage.ScaleTypes_mapString, &stage.ScaleType_stagedOrder, &stage.ScaleTypeOrder)

	__gong__resetStageType(&stage.SpeedTypes, &stage.SpeedTypes_mapString, &stage.SpeedType_stagedOrder, &stage.SpeedTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.VolumeTypes, &stage.VolumeTypes_mapString, &stage.VolumeType_stagedOrder, &stage.VolumeTypeOrder)

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
func (AngleType) GongGetAssociationName() any {
	return &AngleType{
	}
}

func (AreaType) GongGetAssociationName() any {
	return &AreaType{
	}
}

func (GridLengthType) GongGetAssociationName() any {
	return &GridLengthType{
	}
}

func (LengthType) GongGetAssociationName() any {
	return &LengthType{
	}
}

func (ScaleType) GongGetAssociationName() any {
	return &ScaleType{
	}
}

func (SpeedType) GongGetAssociationName() any {
	return &SpeedType{
	}
}

func (TimeType) GongGetAssociationName() any {
	return &TimeType{
	}
}

func (VolumeType) GongGetAssociationName() any {
	return &VolumeType{
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
	// reverse maps of direct associations of AngleType
	case AngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AreaType
	case AreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GridLengthType
	case GridLengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LengthType
	case LengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScaleType
	case ScaleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpeedType
	case SpeedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VolumeType
	case VolumeType:
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
	// reverse maps of direct associations of AngleType
	case AngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AreaType
	case AreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GridLengthType
	case GridLengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LengthType
	case LengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ScaleType
	case ScaleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpeedType
	case SpeedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VolumeType
	case VolumeType:
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
func (*AngleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AreaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GridLengthType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LengthType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ScaleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SpeedType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*VolumeType) GongGetReverseFields() []GongReverseField {
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
func (angletype *AngleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (areatype *AreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (gridlengthtype *GridLengthType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (lengthtype *LengthType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (scaletype *ScaleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (speedtype *SpeedType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (timetype *TimeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (volumetype *VolumeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
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
func (angletype *AngleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = angletype.Name
	case "EnclosedText":
		res.valueString = angletype.EnclosedText
	}
	return
}

func (areatype *AreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = areatype.Name
	case "EnclosedText":
		res.valueString = areatype.EnclosedText
	}
	return
}

func (gridlengthtype *GridLengthType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gridlengthtype.Name
	case "EnclosedText":
		res.valueString = gridlengthtype.EnclosedText
	}
	return
}

func (lengthtype *LengthType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lengthtype.Name
	case "EnclosedText":
		res.valueString = lengthtype.EnclosedText
	}
	return
}

func (scaletype *ScaleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scaletype.Name
	case "EnclosedText":
		res.valueString = scaletype.EnclosedText
	}
	return
}

func (speedtype *SpeedType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = speedtype.Name
	case "EnclosedText":
		res.valueString = speedtype.EnclosedText
	}
	return
}

func (timetype *TimeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timetype.Name
	case "EnclosedText":
		res.valueString = timetype.EnclosedText
	}
	return
}

func (volumetype *VolumeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = volumetype.Name
	case "EnclosedText":
		res.valueString = volumetype.EnclosedText
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
func (angletype *AngleType) GongGetGongstructName() string {
	return "AngleType"
}

func (areatype *AreaType) GongGetGongstructName() string {
	return "AreaType"
}

func (gridlengthtype *GridLengthType) GongGetGongstructName() string {
	return "GridLengthType"
}

func (lengthtype *LengthType) GongGetGongstructName() string {
	return "LengthType"
}

func (scaletype *ScaleType) GongGetGongstructName() string {
	return "ScaleType"
}

func (speedtype *SpeedType) GongGetGongstructName() string {
	return "SpeedType"
}

func (timetype *TimeType) GongGetGongstructName() string {
	return "TimeType"
}

func (volumetype *VolumeType) GongGetGongstructName() string {
	return "VolumeType"
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
	__gong__rebuildMapString(stage.AngleTypes, &stage.AngleTypes_mapString)

	__gong__rebuildMapString(stage.AreaTypes, &stage.AreaTypes_mapString)

	__gong__rebuildMapString(stage.GridLengthTypes, &stage.GridLengthTypes_mapString)

	__gong__rebuildMapString(stage.LengthTypes, &stage.LengthTypes_mapString)

	__gong__rebuildMapString(stage.ScaleTypes, &stage.ScaleTypes_mapString)

	__gong__rebuildMapString(stage.SpeedTypes, &stage.SpeedTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.VolumeTypes, &stage.VolumeTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
