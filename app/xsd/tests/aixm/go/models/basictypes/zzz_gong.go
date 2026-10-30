// generated code - do not edit
package basictypes

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
	CodeListTypes                map[*CodeListType]struct{}
	CodeListTypes_instance       map[*CodeListType]*CodeListType
	CodeListTypes_mapString      map[string]*CodeListType
	CodeListTypeOrder            uint
	CodeListType_stagedOrder     map[*CodeListType]uint
	CodeListType_orderStaged     map[uint]*CodeListType
	CodeListTypes_reference      map[*CodeListType]*CodeListType
	CodeListTypes_referenceOrder map[*CodeListType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeListTypeCreateCallback GongOnAfterCreateInterface[CodeListType]
	OnAfterCodeListTypeUpdateCallback GongOnAfterUpdateInterface[CodeListType]
	OnAfterCodeListTypeDeleteCallback GongOnAfterDeleteInterface[CodeListType]

	CodeOrNilReasonListTypes                map[*CodeOrNilReasonListType]struct{}
	CodeOrNilReasonListTypes_instance       map[*CodeOrNilReasonListType]*CodeOrNilReasonListType
	CodeOrNilReasonListTypes_mapString      map[string]*CodeOrNilReasonListType
	CodeOrNilReasonListTypeOrder            uint
	CodeOrNilReasonListType_stagedOrder     map[*CodeOrNilReasonListType]uint
	CodeOrNilReasonListType_orderStaged     map[uint]*CodeOrNilReasonListType
	CodeOrNilReasonListTypes_reference      map[*CodeOrNilReasonListType]*CodeOrNilReasonListType
	CodeOrNilReasonListTypes_referenceOrder map[*CodeOrNilReasonListType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOrNilReasonListTypeCreateCallback GongOnAfterCreateInterface[CodeOrNilReasonListType]
	OnAfterCodeOrNilReasonListTypeUpdateCallback GongOnAfterUpdateInterface[CodeOrNilReasonListType]
	OnAfterCodeOrNilReasonListTypeDeleteCallback GongOnAfterDeleteInterface[CodeOrNilReasonListType]

	CodeTypes                map[*CodeType]struct{}
	CodeTypes_instance       map[*CodeType]*CodeType
	CodeTypes_mapString      map[string]*CodeType
	CodeTypeOrder            uint
	CodeType_stagedOrder     map[*CodeType]uint
	CodeType_orderStaged     map[uint]*CodeType
	CodeTypes_reference      map[*CodeType]*CodeType
	CodeTypes_referenceOrder map[*CodeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTypeCreateCallback GongOnAfterCreateInterface[CodeType]
	OnAfterCodeTypeUpdateCallback GongOnAfterUpdateInterface[CodeType]
	OnAfterCodeTypeDeleteCallback GongOnAfterDeleteInterface[CodeType]

	CodeWithAuthorityTypes                map[*CodeWithAuthorityType]struct{}
	CodeWithAuthorityTypes_instance       map[*CodeWithAuthorityType]*CodeWithAuthorityType
	CodeWithAuthorityTypes_mapString      map[string]*CodeWithAuthorityType
	CodeWithAuthorityTypeOrder            uint
	CodeWithAuthorityType_stagedOrder     map[*CodeWithAuthorityType]uint
	CodeWithAuthorityType_orderStaged     map[uint]*CodeWithAuthorityType
	CodeWithAuthorityTypes_reference      map[*CodeWithAuthorityType]*CodeWithAuthorityType
	CodeWithAuthorityTypes_referenceOrder map[*CodeWithAuthorityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeWithAuthorityTypeCreateCallback GongOnAfterCreateInterface[CodeWithAuthorityType]
	OnAfterCodeWithAuthorityTypeUpdateCallback GongOnAfterUpdateInterface[CodeWithAuthorityType]
	OnAfterCodeWithAuthorityTypeDeleteCallback GongOnAfterDeleteInterface[CodeWithAuthorityType]

	CoordinatesTypes                map[*CoordinatesType]struct{}
	CoordinatesTypes_instance       map[*CoordinatesType]*CoordinatesType
	CoordinatesTypes_mapString      map[string]*CoordinatesType
	CoordinatesTypeOrder            uint
	CoordinatesType_stagedOrder     map[*CoordinatesType]uint
	CoordinatesType_orderStaged     map[uint]*CoordinatesType
	CoordinatesTypes_reference      map[*CoordinatesType]*CoordinatesType
	CoordinatesTypes_referenceOrder map[*CoordinatesType]uint

	// insertion point for slice of pointers maps
	OnAfterCoordinatesTypeCreateCallback GongOnAfterCreateInterface[CoordinatesType]
	OnAfterCoordinatesTypeUpdateCallback GongOnAfterUpdateInterface[CoordinatesType]
	OnAfterCoordinatesTypeDeleteCallback GongOnAfterDeleteInterface[CoordinatesType]

	MeasureListTypes                map[*MeasureListType]struct{}
	MeasureListTypes_instance       map[*MeasureListType]*MeasureListType
	MeasureListTypes_mapString      map[string]*MeasureListType
	MeasureListTypeOrder            uint
	MeasureListType_stagedOrder     map[*MeasureListType]uint
	MeasureListType_orderStaged     map[uint]*MeasureListType
	MeasureListTypes_reference      map[*MeasureListType]*MeasureListType
	MeasureListTypes_referenceOrder map[*MeasureListType]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureListTypeCreateCallback GongOnAfterCreateInterface[MeasureListType]
	OnAfterMeasureListTypeUpdateCallback GongOnAfterUpdateInterface[MeasureListType]
	OnAfterMeasureListTypeDeleteCallback GongOnAfterDeleteInterface[MeasureListType]

	MeasureOrNilReasonListTypes                map[*MeasureOrNilReasonListType]struct{}
	MeasureOrNilReasonListTypes_instance       map[*MeasureOrNilReasonListType]*MeasureOrNilReasonListType
	MeasureOrNilReasonListTypes_mapString      map[string]*MeasureOrNilReasonListType
	MeasureOrNilReasonListTypeOrder            uint
	MeasureOrNilReasonListType_stagedOrder     map[*MeasureOrNilReasonListType]uint
	MeasureOrNilReasonListType_orderStaged     map[uint]*MeasureOrNilReasonListType
	MeasureOrNilReasonListTypes_reference      map[*MeasureOrNilReasonListType]*MeasureOrNilReasonListType
	MeasureOrNilReasonListTypes_referenceOrder map[*MeasureOrNilReasonListType]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureOrNilReasonListTypeCreateCallback GongOnAfterCreateInterface[MeasureOrNilReasonListType]
	OnAfterMeasureOrNilReasonListTypeUpdateCallback GongOnAfterUpdateInterface[MeasureOrNilReasonListType]
	OnAfterMeasureOrNilReasonListTypeDeleteCallback GongOnAfterDeleteInterface[MeasureOrNilReasonListType]

	MeasureTypes                map[*MeasureType]struct{}
	MeasureTypes_instance       map[*MeasureType]*MeasureType
	MeasureTypes_mapString      map[string]*MeasureType
	MeasureTypeOrder            uint
	MeasureType_stagedOrder     map[*MeasureType]uint
	MeasureType_orderStaged     map[uint]*MeasureType
	MeasureTypes_reference      map[*MeasureType]*MeasureType
	MeasureTypes_referenceOrder map[*MeasureType]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureTypeCreateCallback GongOnAfterCreateInterface[MeasureType]
	OnAfterMeasureTypeUpdateCallback GongOnAfterUpdateInterface[MeasureType]
	OnAfterMeasureTypeDeleteCallback GongOnAfterDeleteInterface[MeasureType]

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
	__gong__clearReferences(&stage.CodeListTypes_reference, &stage.CodeListTypes_instance, &stage.CodeListTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOrNilReasonListTypes_reference, &stage.CodeOrNilReasonListTypes_instance, &stage.CodeOrNilReasonListTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTypes_reference, &stage.CodeTypes_instance, &stage.CodeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWithAuthorityTypes_reference, &stage.CodeWithAuthorityTypes_instance, &stage.CodeWithAuthorityTypes_referenceOrder)

	__gong__clearReferences(&stage.CoordinatesTypes_reference, &stage.CoordinatesTypes_instance, &stage.CoordinatesTypes_referenceOrder)

	__gong__clearReferences(&stage.MeasureListTypes_reference, &stage.MeasureListTypes_instance, &stage.MeasureListTypes_referenceOrder)

	__gong__clearReferences(&stage.MeasureOrNilReasonListTypes_reference, &stage.MeasureOrNilReasonListTypes_instance, &stage.MeasureOrNilReasonListTypes_referenceOrder)

	__gong__clearReferences(&stage.MeasureTypes_reference, &stage.MeasureTypes_instance, &stage.MeasureTypes_referenceOrder)

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
	stage.CodeListTypeOrder = __gong__recomputeOrder(stage.CodeListType_stagedOrder)

	stage.CodeOrNilReasonListTypeOrder = __gong__recomputeOrder(stage.CodeOrNilReasonListType_stagedOrder)

	stage.CodeTypeOrder = __gong__recomputeOrder(stage.CodeType_stagedOrder)

	stage.CodeWithAuthorityTypeOrder = __gong__recomputeOrder(stage.CodeWithAuthorityType_stagedOrder)

	stage.CoordinatesTypeOrder = __gong__recomputeOrder(stage.CoordinatesType_stagedOrder)

	stage.MeasureListTypeOrder = __gong__recomputeOrder(stage.MeasureListType_stagedOrder)

	stage.MeasureOrNilReasonListTypeOrder = __gong__recomputeOrder(stage.MeasureOrNilReasonListType_stagedOrder)

	stage.MeasureTypeOrder = __gong__recomputeOrder(stage.MeasureType_stagedOrder)

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
func (*CodeListType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeListTypes, stage.CodeListType_stagedOrder)
}

func (*CodeListType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeListType_orderStaged[order]
}

func (*CodeListType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeListTypes_mapString
}

func (*CodeListType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeListTypes
}

func (*CodeListType) GongNewInstance() any {
	return new(CodeListType)
}

func (*CodeOrNilReasonListType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListType_stagedOrder)
}

func (*CodeOrNilReasonListType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOrNilReasonListType_orderStaged[order]
}

func (*CodeOrNilReasonListType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOrNilReasonListTypes_mapString
}

func (*CodeOrNilReasonListType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOrNilReasonListTypes
}

func (*CodeOrNilReasonListType) GongNewInstance() any {
	return new(CodeOrNilReasonListType)
}

func (*CodeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTypes, stage.CodeType_stagedOrder)
}

func (*CodeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeType_orderStaged[order]
}

func (*CodeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTypes_mapString
}

func (*CodeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTypes
}

func (*CodeType) GongNewInstance() any {
	return new(CodeType)
}

func (*CodeWithAuthorityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityType_stagedOrder)
}

func (*CodeWithAuthorityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeWithAuthorityType_orderStaged[order]
}

func (*CodeWithAuthorityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeWithAuthorityTypes_mapString
}

func (*CodeWithAuthorityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeWithAuthorityTypes
}

func (*CodeWithAuthorityType) GongNewInstance() any {
	return new(CodeWithAuthorityType)
}

func (*CoordinatesType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CoordinatesTypes, stage.CoordinatesType_stagedOrder)
}

func (*CoordinatesType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CoordinatesType_orderStaged[order]
}

func (*CoordinatesType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CoordinatesTypes_mapString
}

func (*CoordinatesType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CoordinatesTypes
}

func (*CoordinatesType) GongNewInstance() any {
	return new(CoordinatesType)
}

func (*MeasureListType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureListTypes, stage.MeasureListType_stagedOrder)
}

func (*MeasureListType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureListType_orderStaged[order]
}

func (*MeasureListType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureListTypes_mapString
}

func (*MeasureListType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureListTypes
}

func (*MeasureListType) GongNewInstance() any {
	return new(MeasureListType)
}

func (*MeasureOrNilReasonListType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListType_stagedOrder)
}

func (*MeasureOrNilReasonListType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureOrNilReasonListType_orderStaged[order]
}

func (*MeasureOrNilReasonListType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureOrNilReasonListTypes_mapString
}

func (*MeasureOrNilReasonListType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureOrNilReasonListTypes
}

func (*MeasureOrNilReasonListType) GongNewInstance() any {
	return new(MeasureOrNilReasonListType)
}

func (*MeasureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureTypes, stage.MeasureType_stagedOrder)
}

func (*MeasureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureType_orderStaged[order]
}

func (*MeasureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureTypes_mapString
}

func (*MeasureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureTypes
}

func (*MeasureType) GongNewInstance() any {
	return new(MeasureType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
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
		CodeListTypes:           make(map[*CodeListType]struct{}),
		CodeListTypes_mapString: make(map[string]*CodeListType),

		CodeOrNilReasonListTypes:           make(map[*CodeOrNilReasonListType]struct{}),
		CodeOrNilReasonListTypes_mapString: make(map[string]*CodeOrNilReasonListType),

		CodeTypes:           make(map[*CodeType]struct{}),
		CodeTypes_mapString: make(map[string]*CodeType),

		CodeWithAuthorityTypes:           make(map[*CodeWithAuthorityType]struct{}),
		CodeWithAuthorityTypes_mapString: make(map[string]*CodeWithAuthorityType),

		CoordinatesTypes:           make(map[*CoordinatesType]struct{}),
		CoordinatesTypes_mapString: make(map[string]*CoordinatesType),

		MeasureListTypes:           make(map[*MeasureListType]struct{}),
		MeasureListTypes_mapString: make(map[string]*MeasureListType),

		MeasureOrNilReasonListTypes:           make(map[*MeasureOrNilReasonListType]struct{}),
		MeasureOrNilReasonListTypes_mapString: make(map[string]*MeasureOrNilReasonListType),

		MeasureTypes:           make(map[*MeasureType]struct{}),
		MeasureTypes_mapString: make(map[string]*MeasureType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CodeListType_stagedOrder: make(map[*CodeListType]uint),
		CodeListType_orderStaged: make(map[uint]*CodeListType),
		CodeListTypes_reference:  make(map[*CodeListType]*CodeListType),

		CodeOrNilReasonListType_stagedOrder: make(map[*CodeOrNilReasonListType]uint),
		CodeOrNilReasonListType_orderStaged: make(map[uint]*CodeOrNilReasonListType),
		CodeOrNilReasonListTypes_reference:  make(map[*CodeOrNilReasonListType]*CodeOrNilReasonListType),

		CodeType_stagedOrder: make(map[*CodeType]uint),
		CodeType_orderStaged: make(map[uint]*CodeType),
		CodeTypes_reference:  make(map[*CodeType]*CodeType),

		CodeWithAuthorityType_stagedOrder: make(map[*CodeWithAuthorityType]uint),
		CodeWithAuthorityType_orderStaged: make(map[uint]*CodeWithAuthorityType),
		CodeWithAuthorityTypes_reference:  make(map[*CodeWithAuthorityType]*CodeWithAuthorityType),

		CoordinatesType_stagedOrder: make(map[*CoordinatesType]uint),
		CoordinatesType_orderStaged: make(map[uint]*CoordinatesType),
		CoordinatesTypes_reference:  make(map[*CoordinatesType]*CoordinatesType),

		MeasureListType_stagedOrder: make(map[*MeasureListType]uint),
		MeasureListType_orderStaged: make(map[uint]*MeasureListType),
		MeasureListTypes_reference:  make(map[*MeasureListType]*MeasureListType),

		MeasureOrNilReasonListType_stagedOrder: make(map[*MeasureOrNilReasonListType]uint),
		MeasureOrNilReasonListType_orderStaged: make(map[uint]*MeasureOrNilReasonListType),
		MeasureOrNilReasonListTypes_reference:  make(map[*MeasureOrNilReasonListType]*MeasureOrNilReasonListType),

		MeasureType_stagedOrder: make(map[*MeasureType]uint),
		MeasureType_orderStaged: make(map[uint]*MeasureType),
		MeasureTypes_reference:  make(map[*MeasureType]*MeasureType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeListType": &CodeListTypeUnmarshaller{},

			"CodeOrNilReasonListType": &CodeOrNilReasonListTypeUnmarshaller{},

			"CodeType": &CodeTypeUnmarshaller{},

			"CodeWithAuthorityType": &CodeWithAuthorityTypeUnmarshaller{},

			"CoordinatesType": &CoordinatesTypeUnmarshaller{},

			"MeasureListType": &MeasureListTypeUnmarshaller{},

			"MeasureOrNilReasonListType": &MeasureOrNilReasonListTypeUnmarshaller{},

			"MeasureType": &MeasureTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeListType"] = len(stage.CodeListTypes)
	stage.Map_GongStructName_InstancesNb["CodeOrNilReasonListType"] = len(stage.CodeOrNilReasonListTypes)
	stage.Map_GongStructName_InstancesNb["CodeType"] = len(stage.CodeTypes)
	stage.Map_GongStructName_InstancesNb["CodeWithAuthorityType"] = len(stage.CodeWithAuthorityTypes)
	stage.Map_GongStructName_InstancesNb["CoordinatesType"] = len(stage.CoordinatesTypes)
	stage.Map_GongStructName_InstancesNb["MeasureListType"] = len(stage.MeasureListTypes)
	stage.Map_GongStructName_InstancesNb["MeasureOrNilReasonListType"] = len(stage.MeasureOrNilReasonListTypes)
	stage.Map_GongStructName_InstancesNb["MeasureType"] = len(stage.MeasureTypes)
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
// Stage puts codelisttype to the model stage
func (codelisttype *CodeListType) Stage(stage *Stage) *CodeListType {
	__gong__stage(stage.CodeListTypes, stage.CodeListType_stagedOrder, stage.CodeListType_orderStaged, &stage.CodeListTypeOrder, stage.CodeListTypes_mapString, codelisttype, codelisttype.Name)
	return codelisttype
}

// StagePreserveOrder puts codelisttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeListTypeOrder
// - update stage.CodeListTypeOrder accordingly
func (codelisttype *CodeListType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeListTypes, stage.CodeListType_stagedOrder, stage.CodeListType_orderStaged, &stage.CodeListTypeOrder, stage.CodeListTypes_mapString, codelisttype, order, codelisttype.Name)
}

// Unstage removes codelisttype off the model stage
func (codelisttype *CodeListType) Unstage(stage *Stage) *CodeListType {
	__gong__unstage(stage.CodeListTypes, stage.CodeListTypes_mapString, codelisttype, codelisttype.Name)
	return codelisttype
}

// UnstageVoid removes codelisttype off the model stage
func (codelisttype *CodeListType) UnstageVoid(stage *Stage) {
	codelisttype.Unstage(stage)
}

func (codelisttype *CodeListType) StageVoid(stage *Stage) {
	codelisttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelisttype *CodeListType) GetName() (res string) {
	return codelisttype.Name
}

// for satisfaction of GongStruct interface
func (codelisttype *CodeListType) SetName(name string) {
	codelisttype.Name = name
}

// Stage puts codeornilreasonlisttype to the model stage
func (codeornilreasonlisttype *CodeOrNilReasonListType) Stage(stage *Stage) *CodeOrNilReasonListType {
	__gong__stage(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListType_stagedOrder, stage.CodeOrNilReasonListType_orderStaged, &stage.CodeOrNilReasonListTypeOrder, stage.CodeOrNilReasonListTypes_mapString, codeornilreasonlisttype, codeornilreasonlisttype.Name)
	return codeornilreasonlisttype
}

// StagePreserveOrder puts codeornilreasonlisttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOrNilReasonListTypeOrder
// - update stage.CodeOrNilReasonListTypeOrder accordingly
func (codeornilreasonlisttype *CodeOrNilReasonListType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListType_stagedOrder, stage.CodeOrNilReasonListType_orderStaged, &stage.CodeOrNilReasonListTypeOrder, stage.CodeOrNilReasonListTypes_mapString, codeornilreasonlisttype, order, codeornilreasonlisttype.Name)
}

// Unstage removes codeornilreasonlisttype off the model stage
func (codeornilreasonlisttype *CodeOrNilReasonListType) Unstage(stage *Stage) *CodeOrNilReasonListType {
	__gong__unstage(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListTypes_mapString, codeornilreasonlisttype, codeornilreasonlisttype.Name)
	return codeornilreasonlisttype
}

// UnstageVoid removes codeornilreasonlisttype off the model stage
func (codeornilreasonlisttype *CodeOrNilReasonListType) UnstageVoid(stage *Stage) {
	codeornilreasonlisttype.Unstage(stage)
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) StageVoid(stage *Stage) {
	codeornilreasonlisttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeornilreasonlisttype *CodeOrNilReasonListType) GetName() (res string) {
	return codeornilreasonlisttype.Name
}

// for satisfaction of GongStruct interface
func (codeornilreasonlisttype *CodeOrNilReasonListType) SetName(name string) {
	codeornilreasonlisttype.Name = name
}

// Stage puts codetype to the model stage
func (codetype *CodeType) Stage(stage *Stage) *CodeType {
	__gong__stage(stage.CodeTypes, stage.CodeType_stagedOrder, stage.CodeType_orderStaged, &stage.CodeTypeOrder, stage.CodeTypes_mapString, codetype, codetype.Name)
	return codetype
}

// StagePreserveOrder puts codetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTypeOrder
// - update stage.CodeTypeOrder accordingly
func (codetype *CodeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTypes, stage.CodeType_stagedOrder, stage.CodeType_orderStaged, &stage.CodeTypeOrder, stage.CodeTypes_mapString, codetype, order, codetype.Name)
}

// Unstage removes codetype off the model stage
func (codetype *CodeType) Unstage(stage *Stage) *CodeType {
	__gong__unstage(stage.CodeTypes, stage.CodeTypes_mapString, codetype, codetype.Name)
	return codetype
}

// UnstageVoid removes codetype off the model stage
func (codetype *CodeType) UnstageVoid(stage *Stage) {
	codetype.Unstage(stage)
}

func (codetype *CodeType) StageVoid(stage *Stage) {
	codetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetype *CodeType) GetName() (res string) {
	return codetype.Name
}

// for satisfaction of GongStruct interface
func (codetype *CodeType) SetName(name string) {
	codetype.Name = name
}

// Stage puts codewithauthoritytype to the model stage
func (codewithauthoritytype *CodeWithAuthorityType) Stage(stage *Stage) *CodeWithAuthorityType {
	__gong__stage(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityType_stagedOrder, stage.CodeWithAuthorityType_orderStaged, &stage.CodeWithAuthorityTypeOrder, stage.CodeWithAuthorityTypes_mapString, codewithauthoritytype, codewithauthoritytype.Name)
	return codewithauthoritytype
}

// StagePreserveOrder puts codewithauthoritytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeWithAuthorityTypeOrder
// - update stage.CodeWithAuthorityTypeOrder accordingly
func (codewithauthoritytype *CodeWithAuthorityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityType_stagedOrder, stage.CodeWithAuthorityType_orderStaged, &stage.CodeWithAuthorityTypeOrder, stage.CodeWithAuthorityTypes_mapString, codewithauthoritytype, order, codewithauthoritytype.Name)
}

// Unstage removes codewithauthoritytype off the model stage
func (codewithauthoritytype *CodeWithAuthorityType) Unstage(stage *Stage) *CodeWithAuthorityType {
	__gong__unstage(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityTypes_mapString, codewithauthoritytype, codewithauthoritytype.Name)
	return codewithauthoritytype
}

// UnstageVoid removes codewithauthoritytype off the model stage
func (codewithauthoritytype *CodeWithAuthorityType) UnstageVoid(stage *Stage) {
	codewithauthoritytype.Unstage(stage)
}

func (codewithauthoritytype *CodeWithAuthorityType) StageVoid(stage *Stage) {
	codewithauthoritytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codewithauthoritytype *CodeWithAuthorityType) GetName() (res string) {
	return codewithauthoritytype.Name
}

// for satisfaction of GongStruct interface
func (codewithauthoritytype *CodeWithAuthorityType) SetName(name string) {
	codewithauthoritytype.Name = name
}

// Stage puts coordinatestype to the model stage
func (coordinatestype *CoordinatesType) Stage(stage *Stage) *CoordinatesType {
	__gong__stage(stage.CoordinatesTypes, stage.CoordinatesType_stagedOrder, stage.CoordinatesType_orderStaged, &stage.CoordinatesTypeOrder, stage.CoordinatesTypes_mapString, coordinatestype, coordinatestype.Name)
	return coordinatestype
}

// StagePreserveOrder puts coordinatestype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CoordinatesTypeOrder
// - update stage.CoordinatesTypeOrder accordingly
func (coordinatestype *CoordinatesType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CoordinatesTypes, stage.CoordinatesType_stagedOrder, stage.CoordinatesType_orderStaged, &stage.CoordinatesTypeOrder, stage.CoordinatesTypes_mapString, coordinatestype, order, coordinatestype.Name)
}

// Unstage removes coordinatestype off the model stage
func (coordinatestype *CoordinatesType) Unstage(stage *Stage) *CoordinatesType {
	__gong__unstage(stage.CoordinatesTypes, stage.CoordinatesTypes_mapString, coordinatestype, coordinatestype.Name)
	return coordinatestype
}

// UnstageVoid removes coordinatestype off the model stage
func (coordinatestype *CoordinatesType) UnstageVoid(stage *Stage) {
	coordinatestype.Unstage(stage)
}

func (coordinatestype *CoordinatesType) StageVoid(stage *Stage) {
	coordinatestype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coordinatestype *CoordinatesType) GetName() (res string) {
	return coordinatestype.Name
}

// for satisfaction of GongStruct interface
func (coordinatestype *CoordinatesType) SetName(name string) {
	coordinatestype.Name = name
}

// Stage puts measurelisttype to the model stage
func (measurelisttype *MeasureListType) Stage(stage *Stage) *MeasureListType {
	__gong__stage(stage.MeasureListTypes, stage.MeasureListType_stagedOrder, stage.MeasureListType_orderStaged, &stage.MeasureListTypeOrder, stage.MeasureListTypes_mapString, measurelisttype, measurelisttype.Name)
	return measurelisttype
}

// StagePreserveOrder puts measurelisttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureListTypeOrder
// - update stage.MeasureListTypeOrder accordingly
func (measurelisttype *MeasureListType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureListTypes, stage.MeasureListType_stagedOrder, stage.MeasureListType_orderStaged, &stage.MeasureListTypeOrder, stage.MeasureListTypes_mapString, measurelisttype, order, measurelisttype.Name)
}

// Unstage removes measurelisttype off the model stage
func (measurelisttype *MeasureListType) Unstage(stage *Stage) *MeasureListType {
	__gong__unstage(stage.MeasureListTypes, stage.MeasureListTypes_mapString, measurelisttype, measurelisttype.Name)
	return measurelisttype
}

// UnstageVoid removes measurelisttype off the model stage
func (measurelisttype *MeasureListType) UnstageVoid(stage *Stage) {
	measurelisttype.Unstage(stage)
}

func (measurelisttype *MeasureListType) StageVoid(stage *Stage) {
	measurelisttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measurelisttype *MeasureListType) GetName() (res string) {
	return measurelisttype.Name
}

// for satisfaction of GongStruct interface
func (measurelisttype *MeasureListType) SetName(name string) {
	measurelisttype.Name = name
}

// Stage puts measureornilreasonlisttype to the model stage
func (measureornilreasonlisttype *MeasureOrNilReasonListType) Stage(stage *Stage) *MeasureOrNilReasonListType {
	__gong__stage(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListType_stagedOrder, stage.MeasureOrNilReasonListType_orderStaged, &stage.MeasureOrNilReasonListTypeOrder, stage.MeasureOrNilReasonListTypes_mapString, measureornilreasonlisttype, measureornilreasonlisttype.Name)
	return measureornilreasonlisttype
}

// StagePreserveOrder puts measureornilreasonlisttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureOrNilReasonListTypeOrder
// - update stage.MeasureOrNilReasonListTypeOrder accordingly
func (measureornilreasonlisttype *MeasureOrNilReasonListType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListType_stagedOrder, stage.MeasureOrNilReasonListType_orderStaged, &stage.MeasureOrNilReasonListTypeOrder, stage.MeasureOrNilReasonListTypes_mapString, measureornilreasonlisttype, order, measureornilreasonlisttype.Name)
}

// Unstage removes measureornilreasonlisttype off the model stage
func (measureornilreasonlisttype *MeasureOrNilReasonListType) Unstage(stage *Stage) *MeasureOrNilReasonListType {
	__gong__unstage(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListTypes_mapString, measureornilreasonlisttype, measureornilreasonlisttype.Name)
	return measureornilreasonlisttype
}

// UnstageVoid removes measureornilreasonlisttype off the model stage
func (measureornilreasonlisttype *MeasureOrNilReasonListType) UnstageVoid(stage *Stage) {
	measureornilreasonlisttype.Unstage(stage)
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) StageVoid(stage *Stage) {
	measureornilreasonlisttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measureornilreasonlisttype *MeasureOrNilReasonListType) GetName() (res string) {
	return measureornilreasonlisttype.Name
}

// for satisfaction of GongStruct interface
func (measureornilreasonlisttype *MeasureOrNilReasonListType) SetName(name string) {
	measureornilreasonlisttype.Name = name
}

// Stage puts measuretype to the model stage
func (measuretype *MeasureType) Stage(stage *Stage) *MeasureType {
	__gong__stage(stage.MeasureTypes, stage.MeasureType_stagedOrder, stage.MeasureType_orderStaged, &stage.MeasureTypeOrder, stage.MeasureTypes_mapString, measuretype, measuretype.Name)
	return measuretype
}

// StagePreserveOrder puts measuretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureTypeOrder
// - update stage.MeasureTypeOrder accordingly
func (measuretype *MeasureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureTypes, stage.MeasureType_stagedOrder, stage.MeasureType_orderStaged, &stage.MeasureTypeOrder, stage.MeasureTypes_mapString, measuretype, order, measuretype.Name)
}

// Unstage removes measuretype off the model stage
func (measuretype *MeasureType) Unstage(stage *Stage) *MeasureType {
	__gong__unstage(stage.MeasureTypes, stage.MeasureTypes_mapString, measuretype, measuretype.Name)
	return measuretype
}

// UnstageVoid removes measuretype off the model stage
func (measuretype *MeasureType) UnstageVoid(stage *Stage) {
	measuretype.Unstage(stage)
}

func (measuretype *MeasureType) StageVoid(stage *Stage) {
	measuretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measuretype *MeasureType) GetName() (res string) {
	return measuretype.Name
}

// for satisfaction of GongStruct interface
func (measuretype *MeasureType) SetName(name string) {
	measuretype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeListTypes, &stage.CodeListTypes_mapString, &stage.CodeListType_stagedOrder, &stage.CodeListTypeOrder)

	__gong__resetStageType(&stage.CodeOrNilReasonListTypes, &stage.CodeOrNilReasonListTypes_mapString, &stage.CodeOrNilReasonListType_stagedOrder, &stage.CodeOrNilReasonListTypeOrder)

	__gong__resetStageType(&stage.CodeTypes, &stage.CodeTypes_mapString, &stage.CodeType_stagedOrder, &stage.CodeTypeOrder)

	__gong__resetStageType(&stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString, &stage.CodeWithAuthorityType_stagedOrder, &stage.CodeWithAuthorityTypeOrder)

	__gong__resetStageType(&stage.CoordinatesTypes, &stage.CoordinatesTypes_mapString, &stage.CoordinatesType_stagedOrder, &stage.CoordinatesTypeOrder)

	__gong__resetStageType(&stage.MeasureListTypes, &stage.MeasureListTypes_mapString, &stage.MeasureListType_stagedOrder, &stage.MeasureListTypeOrder)

	__gong__resetStageType(&stage.MeasureOrNilReasonListTypes, &stage.MeasureOrNilReasonListTypes_mapString, &stage.MeasureOrNilReasonListType_stagedOrder, &stage.MeasureOrNilReasonListTypeOrder)

	__gong__resetStageType(&stage.MeasureTypes, &stage.MeasureTypes_mapString, &stage.MeasureType_stagedOrder, &stage.MeasureTypeOrder)

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
func (CodeListType) GongGetAssociationName() any {
	return &CodeListType{
	}
}

func (CodeOrNilReasonListType) GongGetAssociationName() any {
	return &CodeOrNilReasonListType{
	}
}

func (CodeType) GongGetAssociationName() any {
	return &CodeType{
	}
}

func (CodeWithAuthorityType) GongGetAssociationName() any {
	return &CodeWithAuthorityType{
	}
}

func (CoordinatesType) GongGetAssociationName() any {
	return &CoordinatesType{
	}
}

func (MeasureListType) GongGetAssociationName() any {
	return &MeasureListType{
	}
}

func (MeasureOrNilReasonListType) GongGetAssociationName() any {
	return &MeasureOrNilReasonListType{
	}
}

func (MeasureType) GongGetAssociationName() any {
	return &MeasureType{
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
	// reverse maps of direct associations of CodeListType
	case CodeListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrNilReasonListType
	case CodeOrNilReasonListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeType
	case CodeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWithAuthorityType
	case CodeWithAuthorityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CoordinatesType
	case CoordinatesType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureListType
	case MeasureListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureOrNilReasonListType
	case MeasureOrNilReasonListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureType
	case MeasureType:
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
	// reverse maps of direct associations of CodeListType
	case CodeListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrNilReasonListType
	case CodeOrNilReasonListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeType
	case CodeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWithAuthorityType
	case CodeWithAuthorityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CoordinatesType
	case CoordinatesType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureListType
	case MeasureListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureOrNilReasonListType
	case MeasureOrNilReasonListType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureType
	case MeasureType:
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
func (*CodeListType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOrNilReasonListType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeWithAuthorityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CoordinatesType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureListType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureOrNilReasonListType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureType) GongGetReverseFields() []GongReverseField {
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
func (codelisttype *CodeListType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeSpace",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeSpace",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetype *CodeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeSpace",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeSpace",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coordinatestype *CoordinatesType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Decimal",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Cs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ts",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measurelisttype *MeasureListType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Uom",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Uom",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (measuretype *MeasureType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Uom",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeFloat,
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
func (codelisttype *CodeListType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelisttype.Name
	case "CodeSpace":
		res.valueString = codelisttype.CodeSpace
	case "EnclosedText":
		res.valueString = codelisttype.EnclosedText
	}
	return
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeornilreasonlisttype.Name
	case "CodeSpace":
		res.valueString = codeornilreasonlisttype.CodeSpace
	case "EnclosedText":
		res.valueString = codeornilreasonlisttype.EnclosedText
	}
	return
}

func (codetype *CodeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetype.Name
	case "CodeSpace":
		res.valueString = codetype.CodeSpace
	case "EnclosedText":
		res.valueString = codetype.EnclosedText
	}
	return
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codewithauthoritytype.Name
	case "CodeSpace":
		res.valueString = codewithauthoritytype.CodeSpace
	case "EnclosedText":
		res.valueString = codewithauthoritytype.EnclosedText
	}
	return
}

func (coordinatestype *CoordinatesType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coordinatestype.Name
	case "Decimal":
		res.valueString = coordinatestype.Decimal
	case "Cs":
		res.valueString = coordinatestype.Cs
	case "Ts":
		res.valueString = coordinatestype.Ts
	case "EnclosedText":
		res.valueString = coordinatestype.EnclosedText
	}
	return
}

func (measurelisttype *MeasureListType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measurelisttype.Name
	case "Uom":
		res.valueString = measurelisttype.Uom
	case "EnclosedText":
		res.valueString = measurelisttype.EnclosedText
	}
	return
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measureornilreasonlisttype.Name
	case "Uom":
		res.valueString = measureornilreasonlisttype.Uom
	case "EnclosedText":
		res.valueString = measureornilreasonlisttype.EnclosedText
	}
	return
}

func (measuretype *MeasureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measuretype.Name
	case "Uom":
		res.valueString = measuretype.Uom
	case "EnclosedText":
		res.valueString = fmt.Sprintf("%f", measuretype.EnclosedText)
		res.valueFloat = measuretype.EnclosedText
		res.GongFieldValueType = GongFieldValueTypeFloat
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
func (codelisttype *CodeListType) GongGetGongstructName() string {
	return "CodeListType"
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetGongstructName() string {
	return "CodeOrNilReasonListType"
}

func (codetype *CodeType) GongGetGongstructName() string {
	return "CodeType"
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetGongstructName() string {
	return "CodeWithAuthorityType"
}

func (coordinatestype *CoordinatesType) GongGetGongstructName() string {
	return "CoordinatesType"
}

func (measurelisttype *MeasureListType) GongGetGongstructName() string {
	return "MeasureListType"
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetGongstructName() string {
	return "MeasureOrNilReasonListType"
}

func (measuretype *MeasureType) GongGetGongstructName() string {
	return "MeasureType"
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
	__gong__rebuildMapString(stage.CodeListTypes, &stage.CodeListTypes_mapString)

	__gong__rebuildMapString(stage.CodeOrNilReasonListTypes, &stage.CodeOrNilReasonListTypes_mapString)

	__gong__rebuildMapString(stage.CodeTypes, &stage.CodeTypes_mapString)

	__gong__rebuildMapString(stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString)

	__gong__rebuildMapString(stage.CoordinatesTypes, &stage.CoordinatesTypes_mapString)

	__gong__rebuildMapString(stage.MeasureListTypes, &stage.MeasureListTypes_mapString)

	__gong__rebuildMapString(stage.MeasureOrNilReasonListTypes, &stage.MeasureOrNilReasonListTypes_mapString)

	__gong__rebuildMapString(stage.MeasureTypes, &stage.MeasureTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
