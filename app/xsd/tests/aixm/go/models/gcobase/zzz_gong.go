// generated code - do not edit
package gcobase

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
	AbstractObject_Types                map[*AbstractObject_Type]struct{}
	AbstractObject_Types_instance       map[*AbstractObject_Type]*AbstractObject_Type
	AbstractObject_Types_mapString      map[string]*AbstractObject_Type
	AbstractObject_TypeOrder            uint
	AbstractObject_Type_stagedOrder     map[*AbstractObject_Type]uint
	AbstractObject_Type_orderStaged     map[uint]*AbstractObject_Type
	AbstractObject_Types_reference      map[*AbstractObject_Type]*AbstractObject_Type
	AbstractObject_Types_referenceOrder map[*AbstractObject_Type]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractObject_TypeCreateCallback GongOnAfterCreateInterface[AbstractObject_Type]
	OnAfterAbstractObject_TypeUpdateCallback GongOnAfterUpdateInterface[AbstractObject_Type]
	OnAfterAbstractObject_TypeDeleteCallback GongOnAfterDeleteInterface[AbstractObject_Type]

	CodeListValue_Types                map[*CodeListValue_Type]struct{}
	CodeListValue_Types_instance       map[*CodeListValue_Type]*CodeListValue_Type
	CodeListValue_Types_mapString      map[string]*CodeListValue_Type
	CodeListValue_TypeOrder            uint
	CodeListValue_Type_stagedOrder     map[*CodeListValue_Type]uint
	CodeListValue_Type_orderStaged     map[uint]*CodeListValue_Type
	CodeListValue_Types_reference      map[*CodeListValue_Type]*CodeListValue_Type
	CodeListValue_Types_referenceOrder map[*CodeListValue_Type]uint

	// insertion point for slice of pointers maps
	OnAfterCodeListValue_TypeCreateCallback GongOnAfterCreateInterface[CodeListValue_Type]
	OnAfterCodeListValue_TypeUpdateCallback GongOnAfterUpdateInterface[CodeListValue_Type]
	OnAfterCodeListValue_TypeDeleteCallback GongOnAfterDeleteInterface[CodeListValue_Type]

	ObjectReference_PropertyTypes                map[*ObjectReference_PropertyType]struct{}
	ObjectReference_PropertyTypes_instance       map[*ObjectReference_PropertyType]*ObjectReference_PropertyType
	ObjectReference_PropertyTypes_mapString      map[string]*ObjectReference_PropertyType
	ObjectReference_PropertyTypeOrder            uint
	ObjectReference_PropertyType_stagedOrder     map[*ObjectReference_PropertyType]uint
	ObjectReference_PropertyType_orderStaged     map[uint]*ObjectReference_PropertyType
	ObjectReference_PropertyTypes_reference      map[*ObjectReference_PropertyType]*ObjectReference_PropertyType
	ObjectReference_PropertyTypes_referenceOrder map[*ObjectReference_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterObjectReference_PropertyTypeCreateCallback GongOnAfterCreateInterface[ObjectReference_PropertyType]
	OnAfterObjectReference_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[ObjectReference_PropertyType]
	OnAfterObjectReference_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[ObjectReference_PropertyType]

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
	__gong__clearReferences(&stage.AbstractObject_Types_reference, &stage.AbstractObject_Types_instance, &stage.AbstractObject_Types_referenceOrder)

	__gong__clearReferences(&stage.CodeListValue_Types_reference, &stage.CodeListValue_Types_instance, &stage.CodeListValue_Types_referenceOrder)

	__gong__clearReferences(&stage.ObjectReference_PropertyTypes_reference, &stage.ObjectReference_PropertyTypes_instance, &stage.ObjectReference_PropertyTypes_referenceOrder)

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
	stage.AbstractObject_TypeOrder = __gong__recomputeOrder(stage.AbstractObject_Type_stagedOrder)

	stage.CodeListValue_TypeOrder = __gong__recomputeOrder(stage.CodeListValue_Type_stagedOrder)

	stage.ObjectReference_PropertyTypeOrder = __gong__recomputeOrder(stage.ObjectReference_PropertyType_stagedOrder)

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
func (*AbstractObject_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractObject_Types, stage.AbstractObject_Type_stagedOrder)
}

func (*AbstractObject_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractObject_Type_orderStaged[order]
}

func (*AbstractObject_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractObject_Types_mapString
}

func (*AbstractObject_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractObject_Types
}

func (*AbstractObject_Type) GongNewInstance() any {
	return new(AbstractObject_Type)
}

func (*CodeListValue_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeListValue_Types, stage.CodeListValue_Type_stagedOrder)
}

func (*CodeListValue_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeListValue_Type_orderStaged[order]
}

func (*CodeListValue_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeListValue_Types_mapString
}

func (*CodeListValue_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeListValue_Types
}

func (*CodeListValue_Type) GongNewInstance() any {
	return new(CodeListValue_Type)
}

func (*ObjectReference_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ObjectReference_PropertyTypes, stage.ObjectReference_PropertyType_stagedOrder)
}

func (*ObjectReference_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ObjectReference_PropertyType_orderStaged[order]
}

func (*ObjectReference_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ObjectReference_PropertyTypes_mapString
}

func (*ObjectReference_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ObjectReference_PropertyTypes
}

func (*ObjectReference_PropertyType) GongNewInstance() any {
	return new(ObjectReference_PropertyType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gcobase"
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
		AbstractObject_Types:           make(map[*AbstractObject_Type]struct{}),
		AbstractObject_Types_mapString: make(map[string]*AbstractObject_Type),

		CodeListValue_Types:           make(map[*CodeListValue_Type]struct{}),
		CodeListValue_Types_mapString: make(map[string]*CodeListValue_Type),

		ObjectReference_PropertyTypes:           make(map[*ObjectReference_PropertyType]struct{}),
		ObjectReference_PropertyTypes_mapString: make(map[string]*ObjectReference_PropertyType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractObject_Type_stagedOrder: make(map[*AbstractObject_Type]uint),
		AbstractObject_Type_orderStaged: make(map[uint]*AbstractObject_Type),
		AbstractObject_Types_reference:  make(map[*AbstractObject_Type]*AbstractObject_Type),

		CodeListValue_Type_stagedOrder: make(map[*CodeListValue_Type]uint),
		CodeListValue_Type_orderStaged: make(map[uint]*CodeListValue_Type),
		CodeListValue_Types_reference:  make(map[*CodeListValue_Type]*CodeListValue_Type),

		ObjectReference_PropertyType_stagedOrder: make(map[*ObjectReference_PropertyType]uint),
		ObjectReference_PropertyType_orderStaged: make(map[uint]*ObjectReference_PropertyType),
		ObjectReference_PropertyTypes_reference:  make(map[*ObjectReference_PropertyType]*ObjectReference_PropertyType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractObject_Type": &AbstractObject_TypeUnmarshaller{},

			"CodeListValue_Type": &CodeListValue_TypeUnmarshaller{},

			"ObjectReference_PropertyType": &ObjectReference_PropertyTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractObject_Type"] = len(stage.AbstractObject_Types)
	stage.Map_GongStructName_InstancesNb["CodeListValue_Type"] = len(stage.CodeListValue_Types)
	stage.Map_GongStructName_InstancesNb["ObjectReference_PropertyType"] = len(stage.ObjectReference_PropertyTypes)
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
// Stage puts abstractobject_type to the model stage
func (abstractobject_type *AbstractObject_Type) Stage(stage *Stage) *AbstractObject_Type {
	__gong__stage(stage.AbstractObject_Types, stage.AbstractObject_Type_stagedOrder, stage.AbstractObject_Type_orderStaged, &stage.AbstractObject_TypeOrder, stage.AbstractObject_Types_mapString, abstractobject_type, abstractobject_type.Name)
	return abstractobject_type
}

// StagePreserveOrder puts abstractobject_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractObject_TypeOrder
// - update stage.AbstractObject_TypeOrder accordingly
func (abstractobject_type *AbstractObject_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractObject_Types, stage.AbstractObject_Type_stagedOrder, stage.AbstractObject_Type_orderStaged, &stage.AbstractObject_TypeOrder, stage.AbstractObject_Types_mapString, abstractobject_type, order, abstractobject_type.Name)
}

// Unstage removes abstractobject_type off the model stage
func (abstractobject_type *AbstractObject_Type) Unstage(stage *Stage) *AbstractObject_Type {
	__gong__unstage(stage.AbstractObject_Types, stage.AbstractObject_Types_mapString, abstractobject_type, abstractobject_type.Name)
	return abstractobject_type
}

// UnstageVoid removes abstractobject_type off the model stage
func (abstractobject_type *AbstractObject_Type) UnstageVoid(stage *Stage) {
	abstractobject_type.Unstage(stage)
}

func (abstractobject_type *AbstractObject_Type) StageVoid(stage *Stage) {
	abstractobject_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractobject_type *AbstractObject_Type) GetName() (res string) {
	return abstractobject_type.Name
}

// for satisfaction of GongStruct interface
func (abstractobject_type *AbstractObject_Type) SetName(name string) {
	abstractobject_type.Name = name
}

// Stage puts codelistvalue_type to the model stage
func (codelistvalue_type *CodeListValue_Type) Stage(stage *Stage) *CodeListValue_Type {
	__gong__stage(stage.CodeListValue_Types, stage.CodeListValue_Type_stagedOrder, stage.CodeListValue_Type_orderStaged, &stage.CodeListValue_TypeOrder, stage.CodeListValue_Types_mapString, codelistvalue_type, codelistvalue_type.Name)
	return codelistvalue_type
}

// StagePreserveOrder puts codelistvalue_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeListValue_TypeOrder
// - update stage.CodeListValue_TypeOrder accordingly
func (codelistvalue_type *CodeListValue_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeListValue_Types, stage.CodeListValue_Type_stagedOrder, stage.CodeListValue_Type_orderStaged, &stage.CodeListValue_TypeOrder, stage.CodeListValue_Types_mapString, codelistvalue_type, order, codelistvalue_type.Name)
}

// Unstage removes codelistvalue_type off the model stage
func (codelistvalue_type *CodeListValue_Type) Unstage(stage *Stage) *CodeListValue_Type {
	__gong__unstage(stage.CodeListValue_Types, stage.CodeListValue_Types_mapString, codelistvalue_type, codelistvalue_type.Name)
	return codelistvalue_type
}

// UnstageVoid removes codelistvalue_type off the model stage
func (codelistvalue_type *CodeListValue_Type) UnstageVoid(stage *Stage) {
	codelistvalue_type.Unstage(stage)
}

func (codelistvalue_type *CodeListValue_Type) StageVoid(stage *Stage) {
	codelistvalue_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelistvalue_type *CodeListValue_Type) GetName() (res string) {
	return codelistvalue_type.Name
}

// for satisfaction of GongStruct interface
func (codelistvalue_type *CodeListValue_Type) SetName(name string) {
	codelistvalue_type.Name = name
}

// Stage puts objectreference_propertytype to the model stage
func (objectreference_propertytype *ObjectReference_PropertyType) Stage(stage *Stage) *ObjectReference_PropertyType {
	__gong__stage(stage.ObjectReference_PropertyTypes, stage.ObjectReference_PropertyType_stagedOrder, stage.ObjectReference_PropertyType_orderStaged, &stage.ObjectReference_PropertyTypeOrder, stage.ObjectReference_PropertyTypes_mapString, objectreference_propertytype, objectreference_propertytype.Name)
	return objectreference_propertytype
}

// StagePreserveOrder puts objectreference_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObjectReference_PropertyTypeOrder
// - update stage.ObjectReference_PropertyTypeOrder accordingly
func (objectreference_propertytype *ObjectReference_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ObjectReference_PropertyTypes, stage.ObjectReference_PropertyType_stagedOrder, stage.ObjectReference_PropertyType_orderStaged, &stage.ObjectReference_PropertyTypeOrder, stage.ObjectReference_PropertyTypes_mapString, objectreference_propertytype, order, objectreference_propertytype.Name)
}

// Unstage removes objectreference_propertytype off the model stage
func (objectreference_propertytype *ObjectReference_PropertyType) Unstage(stage *Stage) *ObjectReference_PropertyType {
	__gong__unstage(stage.ObjectReference_PropertyTypes, stage.ObjectReference_PropertyTypes_mapString, objectreference_propertytype, objectreference_propertytype.Name)
	return objectreference_propertytype
}

// UnstageVoid removes objectreference_propertytype off the model stage
func (objectreference_propertytype *ObjectReference_PropertyType) UnstageVoid(stage *Stage) {
	objectreference_propertytype.Unstage(stage)
}

func (objectreference_propertytype *ObjectReference_PropertyType) StageVoid(stage *Stage) {
	objectreference_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (objectreference_propertytype *ObjectReference_PropertyType) GetName() (res string) {
	return objectreference_propertytype.Name
}

// for satisfaction of GongStruct interface
func (objectreference_propertytype *ObjectReference_PropertyType) SetName(name string) {
	objectreference_propertytype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractObject_Types, &stage.AbstractObject_Types_mapString, &stage.AbstractObject_Type_stagedOrder, &stage.AbstractObject_TypeOrder)

	__gong__resetStageType(&stage.CodeListValue_Types, &stage.CodeListValue_Types_mapString, &stage.CodeListValue_Type_stagedOrder, &stage.CodeListValue_TypeOrder)

	__gong__resetStageType(&stage.ObjectReference_PropertyTypes, &stage.ObjectReference_PropertyTypes_mapString, &stage.ObjectReference_PropertyType_stagedOrder, &stage.ObjectReference_PropertyTypeOrder)

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
func (AbstractObject_Type) GongGetAssociationName() any {
	return &AbstractObject_Type{
	}
}

func (CodeListValue_Type) GongGetAssociationName() any {
	return &CodeListValue_Type{
	}
}

func (ObjectReference_PropertyType) GongGetAssociationName() any {
	return &ObjectReference_PropertyType{
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
	// reverse maps of direct associations of AbstractObject_Type
	case AbstractObject_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeListValue_Type
	case CodeListValue_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ObjectReference_PropertyType
	case ObjectReference_PropertyType:
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
	// reverse maps of direct associations of AbstractObject_Type
	case AbstractObject_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeListValue_Type
	case CodeListValue_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ObjectReference_PropertyType
	case ObjectReference_PropertyType:
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
func (*AbstractObject_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeListValue_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ObjectReference_PropertyType) GongGetReverseFields() []GongReverseField {
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
func (abstractobject_type *AbstractObject_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Uuid",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codelistvalue_type *CodeListValue_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeList",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CodeListValue",
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

func (objectreference_propertytype *ObjectReference_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Uuidref",
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
func (abstractobject_type *AbstractObject_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractobject_type.Name
	case "Id":
		res.valueString = abstractobject_type.Id
	case "Uuid":
		res.valueString = abstractobject_type.Uuid
	}
	return
}

func (codelistvalue_type *CodeListValue_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelistvalue_type.Name
	case "CodeList":
		res.valueString = codelistvalue_type.CodeList
	case "CodeListValue":
		res.valueString = codelistvalue_type.CodeListValue
	case "CodeSpace":
		res.valueString = codelistvalue_type.CodeSpace
	case "EnclosedText":
		res.valueString = codelistvalue_type.EnclosedText
	}
	return
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = objectreference_propertytype.Name
	case "NilReason":
		res.valueString = objectreference_propertytype.NilReason
	case "Uuidref":
		res.valueString = objectreference_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = objectreference_propertytype.Type_XSD
	case "Href":
		res.valueString = objectreference_propertytype.Href
	case "Role":
		res.valueString = objectreference_propertytype.Role
	case "Arcrole":
		res.valueString = objectreference_propertytype.Arcrole
	case "Title":
		res.valueString = objectreference_propertytype.Title
	case "Show":
		res.valueString = objectreference_propertytype.Show
	case "Actuate":
		res.valueString = objectreference_propertytype.Actuate
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
func (abstractobject_type *AbstractObject_Type) GongGetGongstructName() string {
	return "AbstractObject_Type"
}

func (codelistvalue_type *CodeListValue_Type) GongGetGongstructName() string {
	return "CodeListValue_Type"
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongGetGongstructName() string {
	return "ObjectReference_PropertyType"
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
	__gong__rebuildMapString(stage.AbstractObject_Types, &stage.AbstractObject_Types_mapString)

	__gong__rebuildMapString(stage.CodeListValue_Types, &stage.CodeListValue_Types_mapString)

	__gong__rebuildMapString(stage.ObjectReference_PropertyTypes, &stage.ObjectReference_PropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
