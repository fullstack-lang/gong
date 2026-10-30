// generated code - do not edit
package applicationschema

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
	MD_ApplicationSchemaInformation_PropertyTypes                map[*MD_ApplicationSchemaInformation_PropertyType]struct{}
	MD_ApplicationSchemaInformation_PropertyTypes_instance       map[*MD_ApplicationSchemaInformation_PropertyType]*MD_ApplicationSchemaInformation_PropertyType
	MD_ApplicationSchemaInformation_PropertyTypes_mapString      map[string]*MD_ApplicationSchemaInformation_PropertyType
	MD_ApplicationSchemaInformation_PropertyTypeOrder            uint
	MD_ApplicationSchemaInformation_PropertyType_stagedOrder     map[*MD_ApplicationSchemaInformation_PropertyType]uint
	MD_ApplicationSchemaInformation_PropertyType_orderStaged     map[uint]*MD_ApplicationSchemaInformation_PropertyType
	MD_ApplicationSchemaInformation_PropertyTypes_reference      map[*MD_ApplicationSchemaInformation_PropertyType]*MD_ApplicationSchemaInformation_PropertyType
	MD_ApplicationSchemaInformation_PropertyTypes_referenceOrder map[*MD_ApplicationSchemaInformation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ApplicationSchemaInformation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ApplicationSchemaInformation_PropertyType]
	OnAfterMD_ApplicationSchemaInformation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ApplicationSchemaInformation_PropertyType]
	OnAfterMD_ApplicationSchemaInformation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ApplicationSchemaInformation_PropertyType]

	MD_ApplicationSchemaInformation_Types                map[*MD_ApplicationSchemaInformation_Type]struct{}
	MD_ApplicationSchemaInformation_Types_instance       map[*MD_ApplicationSchemaInformation_Type]*MD_ApplicationSchemaInformation_Type
	MD_ApplicationSchemaInformation_Types_mapString      map[string]*MD_ApplicationSchemaInformation_Type
	MD_ApplicationSchemaInformation_TypeOrder            uint
	MD_ApplicationSchemaInformation_Type_stagedOrder     map[*MD_ApplicationSchemaInformation_Type]uint
	MD_ApplicationSchemaInformation_Type_orderStaged     map[uint]*MD_ApplicationSchemaInformation_Type
	MD_ApplicationSchemaInformation_Types_reference      map[*MD_ApplicationSchemaInformation_Type]*MD_ApplicationSchemaInformation_Type
	MD_ApplicationSchemaInformation_Types_referenceOrder map[*MD_ApplicationSchemaInformation_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ApplicationSchemaInformation_TypeCreateCallback GongOnAfterCreateInterface[MD_ApplicationSchemaInformation_Type]
	OnAfterMD_ApplicationSchemaInformation_TypeUpdateCallback GongOnAfterUpdateInterface[MD_ApplicationSchemaInformation_Type]
	OnAfterMD_ApplicationSchemaInformation_TypeDeleteCallback GongOnAfterDeleteInterface[MD_ApplicationSchemaInformation_Type]

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
	__gong__clearReferences(&stage.MD_ApplicationSchemaInformation_PropertyTypes_reference, &stage.MD_ApplicationSchemaInformation_PropertyTypes_instance, &stage.MD_ApplicationSchemaInformation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_ApplicationSchemaInformation_Types_reference, &stage.MD_ApplicationSchemaInformation_Types_instance, &stage.MD_ApplicationSchemaInformation_Types_referenceOrder)

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
	stage.MD_ApplicationSchemaInformation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ApplicationSchemaInformation_PropertyType_stagedOrder)

	stage.MD_ApplicationSchemaInformation_TypeOrder = __gong__recomputeOrder(stage.MD_ApplicationSchemaInformation_Type_stagedOrder)

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
func (*MD_ApplicationSchemaInformation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ApplicationSchemaInformation_PropertyTypes, stage.MD_ApplicationSchemaInformation_PropertyType_stagedOrder)
}

func (*MD_ApplicationSchemaInformation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ApplicationSchemaInformation_PropertyType_orderStaged[order]
}

func (*MD_ApplicationSchemaInformation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString
}

func (*MD_ApplicationSchemaInformation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ApplicationSchemaInformation_PropertyTypes
}

func (*MD_ApplicationSchemaInformation_PropertyType) GongNewInstance() any {
	return new(MD_ApplicationSchemaInformation_PropertyType)
}

func (*MD_ApplicationSchemaInformation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ApplicationSchemaInformation_Types, stage.MD_ApplicationSchemaInformation_Type_stagedOrder)
}

func (*MD_ApplicationSchemaInformation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ApplicationSchemaInformation_Type_orderStaged[order]
}

func (*MD_ApplicationSchemaInformation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ApplicationSchemaInformation_Types_mapString
}

func (*MD_ApplicationSchemaInformation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ApplicationSchemaInformation_Types
}

func (*MD_ApplicationSchemaInformation_Type) GongNewInstance() any {
	return new(MD_ApplicationSchemaInformation_Type)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/applicationschema"
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
		MD_ApplicationSchemaInformation_PropertyTypes:           make(map[*MD_ApplicationSchemaInformation_PropertyType]struct{}),
		MD_ApplicationSchemaInformation_PropertyTypes_mapString: make(map[string]*MD_ApplicationSchemaInformation_PropertyType),

		MD_ApplicationSchemaInformation_Types:           make(map[*MD_ApplicationSchemaInformation_Type]struct{}),
		MD_ApplicationSchemaInformation_Types_mapString: make(map[string]*MD_ApplicationSchemaInformation_Type),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		MD_ApplicationSchemaInformation_PropertyType_stagedOrder: make(map[*MD_ApplicationSchemaInformation_PropertyType]uint),
		MD_ApplicationSchemaInformation_PropertyType_orderStaged: make(map[uint]*MD_ApplicationSchemaInformation_PropertyType),
		MD_ApplicationSchemaInformation_PropertyTypes_reference:  make(map[*MD_ApplicationSchemaInformation_PropertyType]*MD_ApplicationSchemaInformation_PropertyType),

		MD_ApplicationSchemaInformation_Type_stagedOrder: make(map[*MD_ApplicationSchemaInformation_Type]uint),
		MD_ApplicationSchemaInformation_Type_orderStaged: make(map[uint]*MD_ApplicationSchemaInformation_Type),
		MD_ApplicationSchemaInformation_Types_reference:  make(map[*MD_ApplicationSchemaInformation_Type]*MD_ApplicationSchemaInformation_Type),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"MD_ApplicationSchemaInformation_PropertyType": &MD_ApplicationSchemaInformation_PropertyTypeUnmarshaller{},

			"MD_ApplicationSchemaInformation_Type": &MD_ApplicationSchemaInformation_TypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["MD_ApplicationSchemaInformation_PropertyType"] = len(stage.MD_ApplicationSchemaInformation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_ApplicationSchemaInformation_Type"] = len(stage.MD_ApplicationSchemaInformation_Types)
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
// Stage puts md_applicationschemainformation_propertytype to the model stage
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) Stage(stage *Stage) *MD_ApplicationSchemaInformation_PropertyType {
	__gong__stage(stage.MD_ApplicationSchemaInformation_PropertyTypes, stage.MD_ApplicationSchemaInformation_PropertyType_stagedOrder, stage.MD_ApplicationSchemaInformation_PropertyType_orderStaged, &stage.MD_ApplicationSchemaInformation_PropertyTypeOrder, stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString, md_applicationschemainformation_propertytype, md_applicationschemainformation_propertytype.Name)
	return md_applicationschemainformation_propertytype
}

// StagePreserveOrder puts md_applicationschemainformation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ApplicationSchemaInformation_PropertyTypeOrder
// - update stage.MD_ApplicationSchemaInformation_PropertyTypeOrder accordingly
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ApplicationSchemaInformation_PropertyTypes, stage.MD_ApplicationSchemaInformation_PropertyType_stagedOrder, stage.MD_ApplicationSchemaInformation_PropertyType_orderStaged, &stage.MD_ApplicationSchemaInformation_PropertyTypeOrder, stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString, md_applicationschemainformation_propertytype, order, md_applicationschemainformation_propertytype.Name)
}

// Unstage removes md_applicationschemainformation_propertytype off the model stage
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) Unstage(stage *Stage) *MD_ApplicationSchemaInformation_PropertyType {
	__gong__unstage(stage.MD_ApplicationSchemaInformation_PropertyTypes, stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString, md_applicationschemainformation_propertytype, md_applicationschemainformation_propertytype.Name)
	return md_applicationschemainformation_propertytype
}

// UnstageVoid removes md_applicationschemainformation_propertytype off the model stage
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) UnstageVoid(stage *Stage) {
	md_applicationschemainformation_propertytype.Unstage(stage)
}

func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) StageVoid(stage *Stage) {
	md_applicationschemainformation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GetName() (res string) {
	return md_applicationschemainformation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) SetName(name string) {
	md_applicationschemainformation_propertytype.Name = name
}

// Stage puts md_applicationschemainformation_type to the model stage
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) Stage(stage *Stage) *MD_ApplicationSchemaInformation_Type {
	__gong__stage(stage.MD_ApplicationSchemaInformation_Types, stage.MD_ApplicationSchemaInformation_Type_stagedOrder, stage.MD_ApplicationSchemaInformation_Type_orderStaged, &stage.MD_ApplicationSchemaInformation_TypeOrder, stage.MD_ApplicationSchemaInformation_Types_mapString, md_applicationschemainformation_type, md_applicationschemainformation_type.Name)
	return md_applicationschemainformation_type
}

// StagePreserveOrder puts md_applicationschemainformation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ApplicationSchemaInformation_TypeOrder
// - update stage.MD_ApplicationSchemaInformation_TypeOrder accordingly
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ApplicationSchemaInformation_Types, stage.MD_ApplicationSchemaInformation_Type_stagedOrder, stage.MD_ApplicationSchemaInformation_Type_orderStaged, &stage.MD_ApplicationSchemaInformation_TypeOrder, stage.MD_ApplicationSchemaInformation_Types_mapString, md_applicationschemainformation_type, order, md_applicationschemainformation_type.Name)
}

// Unstage removes md_applicationschemainformation_type off the model stage
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) Unstage(stage *Stage) *MD_ApplicationSchemaInformation_Type {
	__gong__unstage(stage.MD_ApplicationSchemaInformation_Types, stage.MD_ApplicationSchemaInformation_Types_mapString, md_applicationschemainformation_type, md_applicationschemainformation_type.Name)
	return md_applicationschemainformation_type
}

// UnstageVoid removes md_applicationschemainformation_type off the model stage
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) UnstageVoid(stage *Stage) {
	md_applicationschemainformation_type.Unstage(stage)
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) StageVoid(stage *Stage) {
	md_applicationschemainformation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GetName() (res string) {
	return md_applicationschemainformation_type.Name
}

// for satisfaction of GongStruct interface
func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) SetName(name string) {
	md_applicationschemainformation_type.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.MD_ApplicationSchemaInformation_PropertyTypes, &stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString, &stage.MD_ApplicationSchemaInformation_PropertyType_stagedOrder, &stage.MD_ApplicationSchemaInformation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_ApplicationSchemaInformation_Types, &stage.MD_ApplicationSchemaInformation_Types_mapString, &stage.MD_ApplicationSchemaInformation_Type_stagedOrder, &stage.MD_ApplicationSchemaInformation_TypeOrder)

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
func (MD_ApplicationSchemaInformation_PropertyType) GongGetAssociationName() any {
	return &MD_ApplicationSchemaInformation_PropertyType{
			MD_ApplicationSchemaInformation: &MD_ApplicationSchemaInformation_Type{Name: "MD_ApplicationSchemaInformation"},
	}
}

func (MD_ApplicationSchemaInformation_Type) GongGetAssociationName() any {
	return &MD_ApplicationSchemaInformation_Type{
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
	// reverse maps of direct associations of MD_ApplicationSchemaInformation_PropertyType
	case MD_ApplicationSchemaInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_ApplicationSchemaInformation":
			res := make(map[*MD_ApplicationSchemaInformation_Type][]*MD_ApplicationSchemaInformation_PropertyType)
			for md_applicationschemainformation_propertytype := range stage.MD_ApplicationSchemaInformation_PropertyTypes {
				if md_applicationschemainformation_propertytype.MD_ApplicationSchemaInformation != nil {
					md_applicationschemainformation_type_ := md_applicationschemainformation_propertytype.MD_ApplicationSchemaInformation
					var md_applicationschemainformation_propertytypes []*MD_ApplicationSchemaInformation_PropertyType
					_, ok := res[md_applicationschemainformation_type_]
					if ok {
						md_applicationschemainformation_propertytypes = res[md_applicationschemainformation_type_]
					} else {
						md_applicationschemainformation_propertytypes = make([]*MD_ApplicationSchemaInformation_PropertyType, 0)
					}
					md_applicationschemainformation_propertytypes = append(md_applicationschemainformation_propertytypes, md_applicationschemainformation_propertytype)
					res[md_applicationschemainformation_type_] = md_applicationschemainformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ApplicationSchemaInformation_Type
	case MD_ApplicationSchemaInformation_Type:
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
	// reverse maps of direct associations of MD_ApplicationSchemaInformation_PropertyType
	case MD_ApplicationSchemaInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ApplicationSchemaInformation_Type
	case MD_ApplicationSchemaInformation_Type:
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
func (*MD_ApplicationSchemaInformation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ApplicationSchemaInformation_Type) GongGetReverseFields() []GongReverseField {
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
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "MD_ApplicationSchemaInformation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ApplicationSchemaInformation_Type",
		},
	}
	return
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_applicationschemainformation_propertytype.Name
	case "NilReason":
		res.valueString = md_applicationschemainformation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_applicationschemainformation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_applicationschemainformation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_applicationschemainformation_propertytype.Href
	case "Role":
		res.valueString = md_applicationschemainformation_propertytype.Role
	case "Arcrole":
		res.valueString = md_applicationschemainformation_propertytype.Arcrole
	case "Title":
		res.valueString = md_applicationschemainformation_propertytype.Title
	case "Show":
		res.valueString = md_applicationschemainformation_propertytype.Show
	case "Actuate":
		res.valueString = md_applicationschemainformation_propertytype.Actuate
	case "MD_ApplicationSchemaInformation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_applicationschemainformation_propertytype.MD_ApplicationSchemaInformation != nil {
			res.valueString = md_applicationschemainformation_propertytype.MD_ApplicationSchemaInformation.Name
			res.ids = md_applicationschemainformation_propertytype.MD_ApplicationSchemaInformation.GongGetUUID(stage)
		}
	}
	return
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_applicationschemainformation_type.Name
	case "Id":
		res.valueString = md_applicationschemainformation_type.Id
	case "Uuid":
		res.valueString = md_applicationschemainformation_type.Uuid
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
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongGetGongstructName() string {
	return "MD_ApplicationSchemaInformation_PropertyType"
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongGetGongstructName() string {
	return "MD_ApplicationSchemaInformation_Type"
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
	__gong__rebuildMapString(stage.MD_ApplicationSchemaInformation_PropertyTypes, &stage.MD_ApplicationSchemaInformation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_ApplicationSchemaInformation_Types, &stage.MD_ApplicationSchemaInformation_Types_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
