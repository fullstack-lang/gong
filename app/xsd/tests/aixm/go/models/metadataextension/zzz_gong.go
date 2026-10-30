// generated code - do not edit
package metadataextension

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
	MD_DatatypeCode_PropertyTypes                map[*MD_DatatypeCode_PropertyType]struct{}
	MD_DatatypeCode_PropertyTypes_instance       map[*MD_DatatypeCode_PropertyType]*MD_DatatypeCode_PropertyType
	MD_DatatypeCode_PropertyTypes_mapString      map[string]*MD_DatatypeCode_PropertyType
	MD_DatatypeCode_PropertyTypeOrder            uint
	MD_DatatypeCode_PropertyType_stagedOrder     map[*MD_DatatypeCode_PropertyType]uint
	MD_DatatypeCode_PropertyType_orderStaged     map[uint]*MD_DatatypeCode_PropertyType
	MD_DatatypeCode_PropertyTypes_reference      map[*MD_DatatypeCode_PropertyType]*MD_DatatypeCode_PropertyType
	MD_DatatypeCode_PropertyTypes_referenceOrder map[*MD_DatatypeCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_DatatypeCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_DatatypeCode_PropertyType]
	OnAfterMD_DatatypeCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_DatatypeCode_PropertyType]
	OnAfterMD_DatatypeCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_DatatypeCode_PropertyType]

	MD_ExtendedElementInformation_PropertyTypes                map[*MD_ExtendedElementInformation_PropertyType]struct{}
	MD_ExtendedElementInformation_PropertyTypes_instance       map[*MD_ExtendedElementInformation_PropertyType]*MD_ExtendedElementInformation_PropertyType
	MD_ExtendedElementInformation_PropertyTypes_mapString      map[string]*MD_ExtendedElementInformation_PropertyType
	MD_ExtendedElementInformation_PropertyTypeOrder            uint
	MD_ExtendedElementInformation_PropertyType_stagedOrder     map[*MD_ExtendedElementInformation_PropertyType]uint
	MD_ExtendedElementInformation_PropertyType_orderStaged     map[uint]*MD_ExtendedElementInformation_PropertyType
	MD_ExtendedElementInformation_PropertyTypes_reference      map[*MD_ExtendedElementInformation_PropertyType]*MD_ExtendedElementInformation_PropertyType
	MD_ExtendedElementInformation_PropertyTypes_referenceOrder map[*MD_ExtendedElementInformation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ExtendedElementInformation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ExtendedElementInformation_PropertyType]
	OnAfterMD_ExtendedElementInformation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ExtendedElementInformation_PropertyType]
	OnAfterMD_ExtendedElementInformation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ExtendedElementInformation_PropertyType]

	MD_ExtendedElementInformation_Types                map[*MD_ExtendedElementInformation_Type]struct{}
	MD_ExtendedElementInformation_Types_instance       map[*MD_ExtendedElementInformation_Type]*MD_ExtendedElementInformation_Type
	MD_ExtendedElementInformation_Types_mapString      map[string]*MD_ExtendedElementInformation_Type
	MD_ExtendedElementInformation_TypeOrder            uint
	MD_ExtendedElementInformation_Type_stagedOrder     map[*MD_ExtendedElementInformation_Type]uint
	MD_ExtendedElementInformation_Type_orderStaged     map[uint]*MD_ExtendedElementInformation_Type
	MD_ExtendedElementInformation_Types_reference      map[*MD_ExtendedElementInformation_Type]*MD_ExtendedElementInformation_Type
	MD_ExtendedElementInformation_Types_referenceOrder map[*MD_ExtendedElementInformation_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ExtendedElementInformation_TypeCreateCallback GongOnAfterCreateInterface[MD_ExtendedElementInformation_Type]
	OnAfterMD_ExtendedElementInformation_TypeUpdateCallback GongOnAfterUpdateInterface[MD_ExtendedElementInformation_Type]
	OnAfterMD_ExtendedElementInformation_TypeDeleteCallback GongOnAfterDeleteInterface[MD_ExtendedElementInformation_Type]

	MD_MetadataExtensionInformation_PropertyTypes                map[*MD_MetadataExtensionInformation_PropertyType]struct{}
	MD_MetadataExtensionInformation_PropertyTypes_instance       map[*MD_MetadataExtensionInformation_PropertyType]*MD_MetadataExtensionInformation_PropertyType
	MD_MetadataExtensionInformation_PropertyTypes_mapString      map[string]*MD_MetadataExtensionInformation_PropertyType
	MD_MetadataExtensionInformation_PropertyTypeOrder            uint
	MD_MetadataExtensionInformation_PropertyType_stagedOrder     map[*MD_MetadataExtensionInformation_PropertyType]uint
	MD_MetadataExtensionInformation_PropertyType_orderStaged     map[uint]*MD_MetadataExtensionInformation_PropertyType
	MD_MetadataExtensionInformation_PropertyTypes_reference      map[*MD_MetadataExtensionInformation_PropertyType]*MD_MetadataExtensionInformation_PropertyType
	MD_MetadataExtensionInformation_PropertyTypes_referenceOrder map[*MD_MetadataExtensionInformation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_MetadataExtensionInformation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_MetadataExtensionInformation_PropertyType]
	OnAfterMD_MetadataExtensionInformation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_MetadataExtensionInformation_PropertyType]
	OnAfterMD_MetadataExtensionInformation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_MetadataExtensionInformation_PropertyType]

	MD_MetadataExtensionInformation_Types                map[*MD_MetadataExtensionInformation_Type]struct{}
	MD_MetadataExtensionInformation_Types_instance       map[*MD_MetadataExtensionInformation_Type]*MD_MetadataExtensionInformation_Type
	MD_MetadataExtensionInformation_Types_mapString      map[string]*MD_MetadataExtensionInformation_Type
	MD_MetadataExtensionInformation_TypeOrder            uint
	MD_MetadataExtensionInformation_Type_stagedOrder     map[*MD_MetadataExtensionInformation_Type]uint
	MD_MetadataExtensionInformation_Type_orderStaged     map[uint]*MD_MetadataExtensionInformation_Type
	MD_MetadataExtensionInformation_Types_reference      map[*MD_MetadataExtensionInformation_Type]*MD_MetadataExtensionInformation_Type
	MD_MetadataExtensionInformation_Types_referenceOrder map[*MD_MetadataExtensionInformation_Type]uint

	// insertion point for slice of pointers maps
	MD_MetadataExtensionInformation_Type_ExtendedElementInformation_reverseMap map[*MD_ExtendedElementInformation_PropertyType]*MD_MetadataExtensionInformation_Type

	OnAfterMD_MetadataExtensionInformation_TypeCreateCallback GongOnAfterCreateInterface[MD_MetadataExtensionInformation_Type]
	OnAfterMD_MetadataExtensionInformation_TypeUpdateCallback GongOnAfterUpdateInterface[MD_MetadataExtensionInformation_Type]
	OnAfterMD_MetadataExtensionInformation_TypeDeleteCallback GongOnAfterDeleteInterface[MD_MetadataExtensionInformation_Type]

	MD_ObligationCode_PropertyTypes                map[*MD_ObligationCode_PropertyType]struct{}
	MD_ObligationCode_PropertyTypes_instance       map[*MD_ObligationCode_PropertyType]*MD_ObligationCode_PropertyType
	MD_ObligationCode_PropertyTypes_mapString      map[string]*MD_ObligationCode_PropertyType
	MD_ObligationCode_PropertyTypeOrder            uint
	MD_ObligationCode_PropertyType_stagedOrder     map[*MD_ObligationCode_PropertyType]uint
	MD_ObligationCode_PropertyType_orderStaged     map[uint]*MD_ObligationCode_PropertyType
	MD_ObligationCode_PropertyTypes_reference      map[*MD_ObligationCode_PropertyType]*MD_ObligationCode_PropertyType
	MD_ObligationCode_PropertyTypes_referenceOrder map[*MD_ObligationCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ObligationCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ObligationCode_PropertyType]
	OnAfterMD_ObligationCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ObligationCode_PropertyType]
	OnAfterMD_ObligationCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ObligationCode_PropertyType]

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
	__gong__clearReferences(&stage.MD_DatatypeCode_PropertyTypes_reference, &stage.MD_DatatypeCode_PropertyTypes_instance, &stage.MD_DatatypeCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_ExtendedElementInformation_PropertyTypes_reference, &stage.MD_ExtendedElementInformation_PropertyTypes_instance, &stage.MD_ExtendedElementInformation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_ExtendedElementInformation_Types_reference, &stage.MD_ExtendedElementInformation_Types_instance, &stage.MD_ExtendedElementInformation_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_MetadataExtensionInformation_PropertyTypes_reference, &stage.MD_MetadataExtensionInformation_PropertyTypes_instance, &stage.MD_MetadataExtensionInformation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_MetadataExtensionInformation_Types_reference, &stage.MD_MetadataExtensionInformation_Types_instance, &stage.MD_MetadataExtensionInformation_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_ObligationCode_PropertyTypes_reference, &stage.MD_ObligationCode_PropertyTypes_instance, &stage.MD_ObligationCode_PropertyTypes_referenceOrder)

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
	stage.MD_DatatypeCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_DatatypeCode_PropertyType_stagedOrder)

	stage.MD_ExtendedElementInformation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ExtendedElementInformation_PropertyType_stagedOrder)

	stage.MD_ExtendedElementInformation_TypeOrder = __gong__recomputeOrder(stage.MD_ExtendedElementInformation_Type_stagedOrder)

	stage.MD_MetadataExtensionInformation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder)

	stage.MD_MetadataExtensionInformation_TypeOrder = __gong__recomputeOrder(stage.MD_MetadataExtensionInformation_Type_stagedOrder)

	stage.MD_ObligationCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ObligationCode_PropertyType_stagedOrder)

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
func (*MD_DatatypeCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyType_stagedOrder)
}

func (*MD_DatatypeCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_DatatypeCode_PropertyType_orderStaged[order]
}

func (*MD_DatatypeCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_DatatypeCode_PropertyTypes_mapString
}

func (*MD_DatatypeCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_DatatypeCode_PropertyTypes
}

func (*MD_DatatypeCode_PropertyType) GongNewInstance() any {
	return new(MD_DatatypeCode_PropertyType)
}

func (*MD_ExtendedElementInformation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyType_stagedOrder)
}

func (*MD_ExtendedElementInformation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ExtendedElementInformation_PropertyType_orderStaged[order]
}

func (*MD_ExtendedElementInformation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ExtendedElementInformation_PropertyTypes_mapString
}

func (*MD_ExtendedElementInformation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ExtendedElementInformation_PropertyTypes
}

func (*MD_ExtendedElementInformation_PropertyType) GongNewInstance() any {
	return new(MD_ExtendedElementInformation_PropertyType)
}

func (*MD_ExtendedElementInformation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Type_stagedOrder)
}

func (*MD_ExtendedElementInformation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ExtendedElementInformation_Type_orderStaged[order]
}

func (*MD_ExtendedElementInformation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ExtendedElementInformation_Types_mapString
}

func (*MD_ExtendedElementInformation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ExtendedElementInformation_Types
}

func (*MD_ExtendedElementInformation_Type) GongNewInstance() any {
	return new(MD_ExtendedElementInformation_Type)
}

func (*MD_MetadataExtensionInformation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder)
}

func (*MD_MetadataExtensionInformation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_MetadataExtensionInformation_PropertyType_orderStaged[order]
}

func (*MD_MetadataExtensionInformation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_MetadataExtensionInformation_PropertyTypes_mapString
}

func (*MD_MetadataExtensionInformation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_MetadataExtensionInformation_PropertyTypes
}

func (*MD_MetadataExtensionInformation_PropertyType) GongNewInstance() any {
	return new(MD_MetadataExtensionInformation_PropertyType)
}

func (*MD_MetadataExtensionInformation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Type_stagedOrder)
}

func (*MD_MetadataExtensionInformation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_MetadataExtensionInformation_Type_orderStaged[order]
}

func (*MD_MetadataExtensionInformation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_MetadataExtensionInformation_Types_mapString
}

func (*MD_MetadataExtensionInformation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_MetadataExtensionInformation_Types
}

func (*MD_MetadataExtensionInformation_Type) GongNewInstance() any {
	return new(MD_MetadataExtensionInformation_Type)
}

func (*MD_ObligationCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyType_stagedOrder)
}

func (*MD_ObligationCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ObligationCode_PropertyType_orderStaged[order]
}

func (*MD_ObligationCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ObligationCode_PropertyTypes_mapString
}

func (*MD_ObligationCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ObligationCode_PropertyTypes
}

func (*MD_ObligationCode_PropertyType) GongNewInstance() any {
	return new(MD_ObligationCode_PropertyType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension"
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
		MD_DatatypeCode_PropertyTypes:           make(map[*MD_DatatypeCode_PropertyType]struct{}),
		MD_DatatypeCode_PropertyTypes_mapString: make(map[string]*MD_DatatypeCode_PropertyType),

		MD_ExtendedElementInformation_PropertyTypes:           make(map[*MD_ExtendedElementInformation_PropertyType]struct{}),
		MD_ExtendedElementInformation_PropertyTypes_mapString: make(map[string]*MD_ExtendedElementInformation_PropertyType),

		MD_ExtendedElementInformation_Types:           make(map[*MD_ExtendedElementInformation_Type]struct{}),
		MD_ExtendedElementInformation_Types_mapString: make(map[string]*MD_ExtendedElementInformation_Type),

		MD_MetadataExtensionInformation_PropertyTypes:           make(map[*MD_MetadataExtensionInformation_PropertyType]struct{}),
		MD_MetadataExtensionInformation_PropertyTypes_mapString: make(map[string]*MD_MetadataExtensionInformation_PropertyType),

		MD_MetadataExtensionInformation_Types:           make(map[*MD_MetadataExtensionInformation_Type]struct{}),
		MD_MetadataExtensionInformation_Types_mapString: make(map[string]*MD_MetadataExtensionInformation_Type),

		MD_ObligationCode_PropertyTypes:           make(map[*MD_ObligationCode_PropertyType]struct{}),
		MD_ObligationCode_PropertyTypes_mapString: make(map[string]*MD_ObligationCode_PropertyType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		MD_DatatypeCode_PropertyType_stagedOrder: make(map[*MD_DatatypeCode_PropertyType]uint),
		MD_DatatypeCode_PropertyType_orderStaged: make(map[uint]*MD_DatatypeCode_PropertyType),
		MD_DatatypeCode_PropertyTypes_reference:  make(map[*MD_DatatypeCode_PropertyType]*MD_DatatypeCode_PropertyType),

		MD_ExtendedElementInformation_PropertyType_stagedOrder: make(map[*MD_ExtendedElementInformation_PropertyType]uint),
		MD_ExtendedElementInformation_PropertyType_orderStaged: make(map[uint]*MD_ExtendedElementInformation_PropertyType),
		MD_ExtendedElementInformation_PropertyTypes_reference:  make(map[*MD_ExtendedElementInformation_PropertyType]*MD_ExtendedElementInformation_PropertyType),

		MD_ExtendedElementInformation_Type_stagedOrder: make(map[*MD_ExtendedElementInformation_Type]uint),
		MD_ExtendedElementInformation_Type_orderStaged: make(map[uint]*MD_ExtendedElementInformation_Type),
		MD_ExtendedElementInformation_Types_reference:  make(map[*MD_ExtendedElementInformation_Type]*MD_ExtendedElementInformation_Type),

		MD_MetadataExtensionInformation_PropertyType_stagedOrder: make(map[*MD_MetadataExtensionInformation_PropertyType]uint),
		MD_MetadataExtensionInformation_PropertyType_orderStaged: make(map[uint]*MD_MetadataExtensionInformation_PropertyType),
		MD_MetadataExtensionInformation_PropertyTypes_reference:  make(map[*MD_MetadataExtensionInformation_PropertyType]*MD_MetadataExtensionInformation_PropertyType),

		MD_MetadataExtensionInformation_Type_stagedOrder: make(map[*MD_MetadataExtensionInformation_Type]uint),
		MD_MetadataExtensionInformation_Type_orderStaged: make(map[uint]*MD_MetadataExtensionInformation_Type),
		MD_MetadataExtensionInformation_Types_reference:  make(map[*MD_MetadataExtensionInformation_Type]*MD_MetadataExtensionInformation_Type),

		MD_ObligationCode_PropertyType_stagedOrder: make(map[*MD_ObligationCode_PropertyType]uint),
		MD_ObligationCode_PropertyType_orderStaged: make(map[uint]*MD_ObligationCode_PropertyType),
		MD_ObligationCode_PropertyTypes_reference:  make(map[*MD_ObligationCode_PropertyType]*MD_ObligationCode_PropertyType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"MD_DatatypeCode_PropertyType": &MD_DatatypeCode_PropertyTypeUnmarshaller{},

			"MD_ExtendedElementInformation_PropertyType": &MD_ExtendedElementInformation_PropertyTypeUnmarshaller{},

			"MD_ExtendedElementInformation_Type": &MD_ExtendedElementInformation_TypeUnmarshaller{},

			"MD_MetadataExtensionInformation_PropertyType": &MD_MetadataExtensionInformation_PropertyTypeUnmarshaller{},

			"MD_MetadataExtensionInformation_Type": &MD_MetadataExtensionInformation_TypeUnmarshaller{},

			"MD_ObligationCode_PropertyType": &MD_ObligationCode_PropertyTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["MD_DatatypeCode_PropertyType"] = len(stage.MD_DatatypeCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_ExtendedElementInformation_PropertyType"] = len(stage.MD_ExtendedElementInformation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_ExtendedElementInformation_Type"] = len(stage.MD_ExtendedElementInformation_Types)
	stage.Map_GongStructName_InstancesNb["MD_MetadataExtensionInformation_PropertyType"] = len(stage.MD_MetadataExtensionInformation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_MetadataExtensionInformation_Type"] = len(stage.MD_MetadataExtensionInformation_Types)
	stage.Map_GongStructName_InstancesNb["MD_ObligationCode_PropertyType"] = len(stage.MD_ObligationCode_PropertyTypes)
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
// Stage puts md_datatypecode_propertytype to the model stage
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) Stage(stage *Stage) *MD_DatatypeCode_PropertyType {
	__gong__stage(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyType_stagedOrder, stage.MD_DatatypeCode_PropertyType_orderStaged, &stage.MD_DatatypeCode_PropertyTypeOrder, stage.MD_DatatypeCode_PropertyTypes_mapString, md_datatypecode_propertytype, md_datatypecode_propertytype.Name)
	return md_datatypecode_propertytype
}

// StagePreserveOrder puts md_datatypecode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_DatatypeCode_PropertyTypeOrder
// - update stage.MD_DatatypeCode_PropertyTypeOrder accordingly
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyType_stagedOrder, stage.MD_DatatypeCode_PropertyType_orderStaged, &stage.MD_DatatypeCode_PropertyTypeOrder, stage.MD_DatatypeCode_PropertyTypes_mapString, md_datatypecode_propertytype, order, md_datatypecode_propertytype.Name)
}

// Unstage removes md_datatypecode_propertytype off the model stage
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) Unstage(stage *Stage) *MD_DatatypeCode_PropertyType {
	__gong__unstage(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyTypes_mapString, md_datatypecode_propertytype, md_datatypecode_propertytype.Name)
	return md_datatypecode_propertytype
}

// UnstageVoid removes md_datatypecode_propertytype off the model stage
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) UnstageVoid(stage *Stage) {
	md_datatypecode_propertytype.Unstage(stage)
}

func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) StageVoid(stage *Stage) {
	md_datatypecode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GetName() (res string) {
	return md_datatypecode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) SetName(name string) {
	md_datatypecode_propertytype.Name = name
}

// Stage puts md_extendedelementinformation_propertytype to the model stage
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) Stage(stage *Stage) *MD_ExtendedElementInformation_PropertyType {
	__gong__stage(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyType_stagedOrder, stage.MD_ExtendedElementInformation_PropertyType_orderStaged, &stage.MD_ExtendedElementInformation_PropertyTypeOrder, stage.MD_ExtendedElementInformation_PropertyTypes_mapString, md_extendedelementinformation_propertytype, md_extendedelementinformation_propertytype.Name)
	return md_extendedelementinformation_propertytype
}

// StagePreserveOrder puts md_extendedelementinformation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ExtendedElementInformation_PropertyTypeOrder
// - update stage.MD_ExtendedElementInformation_PropertyTypeOrder accordingly
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyType_stagedOrder, stage.MD_ExtendedElementInformation_PropertyType_orderStaged, &stage.MD_ExtendedElementInformation_PropertyTypeOrder, stage.MD_ExtendedElementInformation_PropertyTypes_mapString, md_extendedelementinformation_propertytype, order, md_extendedelementinformation_propertytype.Name)
}

// Unstage removes md_extendedelementinformation_propertytype off the model stage
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) Unstage(stage *Stage) *MD_ExtendedElementInformation_PropertyType {
	__gong__unstage(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyTypes_mapString, md_extendedelementinformation_propertytype, md_extendedelementinformation_propertytype.Name)
	return md_extendedelementinformation_propertytype
}

// UnstageVoid removes md_extendedelementinformation_propertytype off the model stage
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) UnstageVoid(stage *Stage) {
	md_extendedelementinformation_propertytype.Unstage(stage)
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) StageVoid(stage *Stage) {
	md_extendedelementinformation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GetName() (res string) {
	return md_extendedelementinformation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) SetName(name string) {
	md_extendedelementinformation_propertytype.Name = name
}

// Stage puts md_extendedelementinformation_type to the model stage
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) Stage(stage *Stage) *MD_ExtendedElementInformation_Type {
	__gong__stage(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Type_stagedOrder, stage.MD_ExtendedElementInformation_Type_orderStaged, &stage.MD_ExtendedElementInformation_TypeOrder, stage.MD_ExtendedElementInformation_Types_mapString, md_extendedelementinformation_type, md_extendedelementinformation_type.Name)
	return md_extendedelementinformation_type
}

// StagePreserveOrder puts md_extendedelementinformation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ExtendedElementInformation_TypeOrder
// - update stage.MD_ExtendedElementInformation_TypeOrder accordingly
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Type_stagedOrder, stage.MD_ExtendedElementInformation_Type_orderStaged, &stage.MD_ExtendedElementInformation_TypeOrder, stage.MD_ExtendedElementInformation_Types_mapString, md_extendedelementinformation_type, order, md_extendedelementinformation_type.Name)
}

// Unstage removes md_extendedelementinformation_type off the model stage
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) Unstage(stage *Stage) *MD_ExtendedElementInformation_Type {
	__gong__unstage(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Types_mapString, md_extendedelementinformation_type, md_extendedelementinformation_type.Name)
	return md_extendedelementinformation_type
}

// UnstageVoid removes md_extendedelementinformation_type off the model stage
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) UnstageVoid(stage *Stage) {
	md_extendedelementinformation_type.Unstage(stage)
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) StageVoid(stage *Stage) {
	md_extendedelementinformation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GetName() (res string) {
	return md_extendedelementinformation_type.Name
}

// for satisfaction of GongStruct interface
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) SetName(name string) {
	md_extendedelementinformation_type.Name = name
}

// Stage puts md_metadataextensioninformation_propertytype to the model stage
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) Stage(stage *Stage) *MD_MetadataExtensionInformation_PropertyType {
	__gong__stage(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder, stage.MD_MetadataExtensionInformation_PropertyType_orderStaged, &stage.MD_MetadataExtensionInformation_PropertyTypeOrder, stage.MD_MetadataExtensionInformation_PropertyTypes_mapString, md_metadataextensioninformation_propertytype, md_metadataextensioninformation_propertytype.Name)
	return md_metadataextensioninformation_propertytype
}

// StagePreserveOrder puts md_metadataextensioninformation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_MetadataExtensionInformation_PropertyTypeOrder
// - update stage.MD_MetadataExtensionInformation_PropertyTypeOrder accordingly
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder, stage.MD_MetadataExtensionInformation_PropertyType_orderStaged, &stage.MD_MetadataExtensionInformation_PropertyTypeOrder, stage.MD_MetadataExtensionInformation_PropertyTypes_mapString, md_metadataextensioninformation_propertytype, order, md_metadataextensioninformation_propertytype.Name)
}

// Unstage removes md_metadataextensioninformation_propertytype off the model stage
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) Unstage(stage *Stage) *MD_MetadataExtensionInformation_PropertyType {
	__gong__unstage(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyTypes_mapString, md_metadataextensioninformation_propertytype, md_metadataextensioninformation_propertytype.Name)
	return md_metadataextensioninformation_propertytype
}

// UnstageVoid removes md_metadataextensioninformation_propertytype off the model stage
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) UnstageVoid(stage *Stage) {
	md_metadataextensioninformation_propertytype.Unstage(stage)
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) StageVoid(stage *Stage) {
	md_metadataextensioninformation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GetName() (res string) {
	return md_metadataextensioninformation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) SetName(name string) {
	md_metadataextensioninformation_propertytype.Name = name
}

// Stage puts md_metadataextensioninformation_type to the model stage
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) Stage(stage *Stage) *MD_MetadataExtensionInformation_Type {
	__gong__stage(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Type_stagedOrder, stage.MD_MetadataExtensionInformation_Type_orderStaged, &stage.MD_MetadataExtensionInformation_TypeOrder, stage.MD_MetadataExtensionInformation_Types_mapString, md_metadataextensioninformation_type, md_metadataextensioninformation_type.Name)
	return md_metadataextensioninformation_type
}

// StagePreserveOrder puts md_metadataextensioninformation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_MetadataExtensionInformation_TypeOrder
// - update stage.MD_MetadataExtensionInformation_TypeOrder accordingly
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Type_stagedOrder, stage.MD_MetadataExtensionInformation_Type_orderStaged, &stage.MD_MetadataExtensionInformation_TypeOrder, stage.MD_MetadataExtensionInformation_Types_mapString, md_metadataextensioninformation_type, order, md_metadataextensioninformation_type.Name)
}

// Unstage removes md_metadataextensioninformation_type off the model stage
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) Unstage(stage *Stage) *MD_MetadataExtensionInformation_Type {
	__gong__unstage(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Types_mapString, md_metadataextensioninformation_type, md_metadataextensioninformation_type.Name)
	return md_metadataextensioninformation_type
}

// UnstageVoid removes md_metadataextensioninformation_type off the model stage
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) UnstageVoid(stage *Stage) {
	md_metadataextensioninformation_type.Unstage(stage)
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) StageVoid(stage *Stage) {
	md_metadataextensioninformation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GetName() (res string) {
	return md_metadataextensioninformation_type.Name
}

// for satisfaction of GongStruct interface
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) SetName(name string) {
	md_metadataextensioninformation_type.Name = name
}

// Stage puts md_obligationcode_propertytype to the model stage
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) Stage(stage *Stage) *MD_ObligationCode_PropertyType {
	__gong__stage(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyType_stagedOrder, stage.MD_ObligationCode_PropertyType_orderStaged, &stage.MD_ObligationCode_PropertyTypeOrder, stage.MD_ObligationCode_PropertyTypes_mapString, md_obligationcode_propertytype, md_obligationcode_propertytype.Name)
	return md_obligationcode_propertytype
}

// StagePreserveOrder puts md_obligationcode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ObligationCode_PropertyTypeOrder
// - update stage.MD_ObligationCode_PropertyTypeOrder accordingly
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyType_stagedOrder, stage.MD_ObligationCode_PropertyType_orderStaged, &stage.MD_ObligationCode_PropertyTypeOrder, stage.MD_ObligationCode_PropertyTypes_mapString, md_obligationcode_propertytype, order, md_obligationcode_propertytype.Name)
}

// Unstage removes md_obligationcode_propertytype off the model stage
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) Unstage(stage *Stage) *MD_ObligationCode_PropertyType {
	__gong__unstage(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyTypes_mapString, md_obligationcode_propertytype, md_obligationcode_propertytype.Name)
	return md_obligationcode_propertytype
}

// UnstageVoid removes md_obligationcode_propertytype off the model stage
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) UnstageVoid(stage *Stage) {
	md_obligationcode_propertytype.Unstage(stage)
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) StageVoid(stage *Stage) {
	md_obligationcode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GetName() (res string) {
	return md_obligationcode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) SetName(name string) {
	md_obligationcode_propertytype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.MD_DatatypeCode_PropertyTypes, &stage.MD_DatatypeCode_PropertyTypes_mapString, &stage.MD_DatatypeCode_PropertyType_stagedOrder, &stage.MD_DatatypeCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_ExtendedElementInformation_PropertyTypes, &stage.MD_ExtendedElementInformation_PropertyTypes_mapString, &stage.MD_ExtendedElementInformation_PropertyType_stagedOrder, &stage.MD_ExtendedElementInformation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_ExtendedElementInformation_Types, &stage.MD_ExtendedElementInformation_Types_mapString, &stage.MD_ExtendedElementInformation_Type_stagedOrder, &stage.MD_ExtendedElementInformation_TypeOrder)

	__gong__resetStageType(&stage.MD_MetadataExtensionInformation_PropertyTypes, &stage.MD_MetadataExtensionInformation_PropertyTypes_mapString, &stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder, &stage.MD_MetadataExtensionInformation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_MetadataExtensionInformation_Types, &stage.MD_MetadataExtensionInformation_Types_mapString, &stage.MD_MetadataExtensionInformation_Type_stagedOrder, &stage.MD_MetadataExtensionInformation_TypeOrder)

	__gong__resetStageType(&stage.MD_ObligationCode_PropertyTypes, &stage.MD_ObligationCode_PropertyTypes_mapString, &stage.MD_ObligationCode_PropertyType_stagedOrder, &stage.MD_ObligationCode_PropertyTypeOrder)

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
func (MD_DatatypeCode_PropertyType) GongGetAssociationName() any {
	return &MD_DatatypeCode_PropertyType{
	}
}

func (MD_ExtendedElementInformation_PropertyType) GongGetAssociationName() any {
	return &MD_ExtendedElementInformation_PropertyType{
			MD_ExtendedElementInformation: &MD_ExtendedElementInformation_Type{Name: "MD_ExtendedElementInformation"},
	}
}

func (MD_ExtendedElementInformation_Type) GongGetAssociationName() any {
	return &MD_ExtendedElementInformation_Type{
			Obligation: &MD_ObligationCode_PropertyType{Name: "Obligation"},
			DataType: &MD_DatatypeCode_PropertyType{Name: "DataType"},
	}
}

func (MD_MetadataExtensionInformation_PropertyType) GongGetAssociationName() any {
	return &MD_MetadataExtensionInformation_PropertyType{
			MD_MetadataExtensionInformation: &MD_MetadataExtensionInformation_Type{Name: "MD_MetadataExtensionInformation"},
	}
}

func (MD_MetadataExtensionInformation_Type) GongGetAssociationName() any {
	return &MD_MetadataExtensionInformation_Type{
			ExtendedElementInformation: []*MD_ExtendedElementInformation_PropertyType{{Name: "ExtendedElementInformation"}},
	}
}

func (MD_ObligationCode_PropertyType) GongGetAssociationName() any {
	return &MD_ObligationCode_PropertyType{
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
	// reverse maps of direct associations of MD_DatatypeCode_PropertyType
	case MD_DatatypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ExtendedElementInformation_PropertyType
	case MD_ExtendedElementInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_ExtendedElementInformation":
			res := make(map[*MD_ExtendedElementInformation_Type][]*MD_ExtendedElementInformation_PropertyType)
			for md_extendedelementinformation_propertytype := range stage.MD_ExtendedElementInformation_PropertyTypes {
				if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
					md_extendedelementinformation_type_ := md_extendedelementinformation_propertytype.MD_ExtendedElementInformation
					var md_extendedelementinformation_propertytypes []*MD_ExtendedElementInformation_PropertyType
					_, ok := res[md_extendedelementinformation_type_]
					if ok {
						md_extendedelementinformation_propertytypes = res[md_extendedelementinformation_type_]
					} else {
						md_extendedelementinformation_propertytypes = make([]*MD_ExtendedElementInformation_PropertyType, 0)
					}
					md_extendedelementinformation_propertytypes = append(md_extendedelementinformation_propertytypes, md_extendedelementinformation_propertytype)
					res[md_extendedelementinformation_type_] = md_extendedelementinformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ExtendedElementInformation_Type
	case MD_ExtendedElementInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "Obligation":
			res := make(map[*MD_ObligationCode_PropertyType][]*MD_ExtendedElementInformation_Type)
			for md_extendedelementinformation_type := range stage.MD_ExtendedElementInformation_Types {
				if md_extendedelementinformation_type.Obligation != nil {
					md_obligationcode_propertytype_ := md_extendedelementinformation_type.Obligation
					var md_extendedelementinformation_types []*MD_ExtendedElementInformation_Type
					_, ok := res[md_obligationcode_propertytype_]
					if ok {
						md_extendedelementinformation_types = res[md_obligationcode_propertytype_]
					} else {
						md_extendedelementinformation_types = make([]*MD_ExtendedElementInformation_Type, 0)
					}
					md_extendedelementinformation_types = append(md_extendedelementinformation_types, md_extendedelementinformation_type)
					res[md_obligationcode_propertytype_] = md_extendedelementinformation_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataType":
			res := make(map[*MD_DatatypeCode_PropertyType][]*MD_ExtendedElementInformation_Type)
			for md_extendedelementinformation_type := range stage.MD_ExtendedElementInformation_Types {
				if md_extendedelementinformation_type.DataType != nil {
					md_datatypecode_propertytype_ := md_extendedelementinformation_type.DataType
					var md_extendedelementinformation_types []*MD_ExtendedElementInformation_Type
					_, ok := res[md_datatypecode_propertytype_]
					if ok {
						md_extendedelementinformation_types = res[md_datatypecode_propertytype_]
					} else {
						md_extendedelementinformation_types = make([]*MD_ExtendedElementInformation_Type, 0)
					}
					md_extendedelementinformation_types = append(md_extendedelementinformation_types, md_extendedelementinformation_type)
					res[md_datatypecode_propertytype_] = md_extendedelementinformation_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_MetadataExtensionInformation_PropertyType
	case MD_MetadataExtensionInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_MetadataExtensionInformation":
			res := make(map[*MD_MetadataExtensionInformation_Type][]*MD_MetadataExtensionInformation_PropertyType)
			for md_metadataextensioninformation_propertytype := range stage.MD_MetadataExtensionInformation_PropertyTypes {
				if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
					md_metadataextensioninformation_type_ := md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation
					var md_metadataextensioninformation_propertytypes []*MD_MetadataExtensionInformation_PropertyType
					_, ok := res[md_metadataextensioninformation_type_]
					if ok {
						md_metadataextensioninformation_propertytypes = res[md_metadataextensioninformation_type_]
					} else {
						md_metadataextensioninformation_propertytypes = make([]*MD_MetadataExtensionInformation_PropertyType, 0)
					}
					md_metadataextensioninformation_propertytypes = append(md_metadataextensioninformation_propertytypes, md_metadataextensioninformation_propertytype)
					res[md_metadataextensioninformation_type_] = md_metadataextensioninformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_MetadataExtensionInformation_Type
	case MD_MetadataExtensionInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ObligationCode_PropertyType
	case MD_ObligationCode_PropertyType:
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
	// reverse maps of direct associations of MD_DatatypeCode_PropertyType
	case MD_DatatypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ExtendedElementInformation_PropertyType
	case MD_ExtendedElementInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ExtendedElementInformation_Type
	case MD_ExtendedElementInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_MetadataExtensionInformation_PropertyType
	case MD_MetadataExtensionInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_MetadataExtensionInformation_Type
	case MD_MetadataExtensionInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ExtendedElementInformation":
			res := make(map[*MD_ExtendedElementInformation_PropertyType][]*MD_MetadataExtensionInformation_Type)
			for md_metadataextensioninformation_type := range stage.MD_MetadataExtensionInformation_Types {
				for _, md_extendedelementinformation_propertytype_ := range md_metadataextensioninformation_type.ExtendedElementInformation {
					res[md_extendedelementinformation_propertytype_] = append(res[md_extendedelementinformation_propertytype_], md_metadataextensioninformation_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ObligationCode_PropertyType
	case MD_ObligationCode_PropertyType:
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
func (*MD_DatatypeCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ExtendedElementInformation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MD_MetadataExtensionInformation_Type",
			Fieldname: "ExtendedElementInformation",
		},
	}
}

func (*MD_ExtendedElementInformation_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_MetadataExtensionInformation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_MetadataExtensionInformation_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ObligationCode_PropertyType) GongGetReverseFields() []GongReverseField {
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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_ExtendedElementInformation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ExtendedElementInformation_Type",
		},
	}
	return
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "Obligation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ObligationCode_PropertyType",
		},
		{
			Name:                 "DataType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_DatatypeCode_PropertyType",
		},
	}
	return
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_MetadataExtensionInformation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_MetadataExtensionInformation_Type",
		},
	}
	return
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "ExtendedElementInformation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_ExtendedElementInformation_PropertyType",
		},
	}
	return
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_ObligationCode",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_MD_ObligationCode_Type",
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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_datatypecode_propertytype.Name
	case "NilReason":
		res.valueString = md_datatypecode_propertytype.NilReason
	}
	return
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_extendedelementinformation_propertytype.Name
	case "NilReason":
		res.valueString = md_extendedelementinformation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_extendedelementinformation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_extendedelementinformation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_extendedelementinformation_propertytype.Href
	case "Role":
		res.valueString = md_extendedelementinformation_propertytype.Role
	case "Arcrole":
		res.valueString = md_extendedelementinformation_propertytype.Arcrole
	case "Title":
		res.valueString = md_extendedelementinformation_propertytype.Title
	case "Show":
		res.valueString = md_extendedelementinformation_propertytype.Show
	case "Actuate":
		res.valueString = md_extendedelementinformation_propertytype.Actuate
	case "MD_ExtendedElementInformation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
			res.valueString = md_extendedelementinformation_propertytype.MD_ExtendedElementInformation.Name
			res.ids = md_extendedelementinformation_propertytype.MD_ExtendedElementInformation.GongGetUUID(stage)
		}
	}
	return
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_extendedelementinformation_type.Name
	case "Id":
		res.valueString = md_extendedelementinformation_type.Id
	case "Uuid":
		res.valueString = md_extendedelementinformation_type.Uuid
	case "Obligation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_extendedelementinformation_type.Obligation != nil {
			res.valueString = md_extendedelementinformation_type.Obligation.Name
			res.ids = md_extendedelementinformation_type.Obligation.GongGetUUID(stage)
		}
	case "DataType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_extendedelementinformation_type.DataType != nil {
			res.valueString = md_extendedelementinformation_type.DataType.Name
			res.ids = md_extendedelementinformation_type.DataType.GongGetUUID(stage)
		}
	}
	return
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_metadataextensioninformation_propertytype.Name
	case "NilReason":
		res.valueString = md_metadataextensioninformation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_metadataextensioninformation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_metadataextensioninformation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_metadataextensioninformation_propertytype.Href
	case "Role":
		res.valueString = md_metadataextensioninformation_propertytype.Role
	case "Arcrole":
		res.valueString = md_metadataextensioninformation_propertytype.Arcrole
	case "Title":
		res.valueString = md_metadataextensioninformation_propertytype.Title
	case "Show":
		res.valueString = md_metadataextensioninformation_propertytype.Show
	case "Actuate":
		res.valueString = md_metadataextensioninformation_propertytype.Actuate
	case "MD_MetadataExtensionInformation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
			res.valueString = md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation.Name
			res.ids = md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation.GongGetUUID(stage)
		}
	}
	return
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_metadataextensioninformation_type.Name
	case "Id":
		res.valueString = md_metadataextensioninformation_type.Id
	case "Uuid":
		res.valueString = md_metadataextensioninformation_type.Uuid
	case "ExtendedElementInformation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_metadataextensioninformation_type.ExtendedElementInformation {
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

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_obligationcode_propertytype.Name
	case "NilReason":
		res.valueString = md_obligationcode_propertytype.NilReason
	case "MD_ObligationCode":
		enum := md_obligationcode_propertytype.MD_ObligationCode
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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongGetGongstructName() string {
	return "MD_DatatypeCode_PropertyType"
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongGetGongstructName() string {
	return "MD_ExtendedElementInformation_PropertyType"
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongGetGongstructName() string {
	return "MD_ExtendedElementInformation_Type"
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongGetGongstructName() string {
	return "MD_MetadataExtensionInformation_PropertyType"
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongGetGongstructName() string {
	return "MD_MetadataExtensionInformation_Type"
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongGetGongstructName() string {
	return "MD_ObligationCode_PropertyType"
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
	__gong__rebuildMapString(stage.MD_DatatypeCode_PropertyTypes, &stage.MD_DatatypeCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_ExtendedElementInformation_PropertyTypes, &stage.MD_ExtendedElementInformation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_ExtendedElementInformation_Types, &stage.MD_ExtendedElementInformation_Types_mapString)

	__gong__rebuildMapString(stage.MD_MetadataExtensionInformation_PropertyTypes, &stage.MD_MetadataExtensionInformation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_MetadataExtensionInformation_Types, &stage.MD_MetadataExtensionInformation_Types_mapString)

	__gong__rebuildMapString(stage.MD_ObligationCode_PropertyTypes, &stage.MD_ObligationCode_PropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
