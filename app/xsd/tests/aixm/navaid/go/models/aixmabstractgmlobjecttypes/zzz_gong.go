// generated code - do not edit
package aixmabstractgmlobjecttypes

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
	AbstractAIXMFeatureBaseTypes                map[*AbstractAIXMFeatureBaseType]struct{}
	AbstractAIXMFeatureBaseTypes_instance       map[*AbstractAIXMFeatureBaseType]*AbstractAIXMFeatureBaseType
	AbstractAIXMFeatureBaseTypes_mapString      map[string]*AbstractAIXMFeatureBaseType
	AbstractAIXMFeatureBaseTypeOrder            uint
	AbstractAIXMFeatureBaseType_stagedOrder     map[*AbstractAIXMFeatureBaseType]uint
	AbstractAIXMFeatureBaseType_orderStaged     map[uint]*AbstractAIXMFeatureBaseType
	AbstractAIXMFeatureBaseTypes_reference      map[*AbstractAIXMFeatureBaseType]*AbstractAIXMFeatureBaseType
	AbstractAIXMFeatureBaseTypes_referenceOrder map[*AbstractAIXMFeatureBaseType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMFeatureBaseTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMFeatureBaseType]
	OnAfterAbstractAIXMFeatureBaseTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMFeatureBaseType]
	OnAfterAbstractAIXMFeatureBaseTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMFeatureBaseType]

	AbstractAIXMFeatureTypes                map[*AbstractAIXMFeatureType]struct{}
	AbstractAIXMFeatureTypes_instance       map[*AbstractAIXMFeatureType]*AbstractAIXMFeatureType
	AbstractAIXMFeatureTypes_mapString      map[string]*AbstractAIXMFeatureType
	AbstractAIXMFeatureTypeOrder            uint
	AbstractAIXMFeatureType_stagedOrder     map[*AbstractAIXMFeatureType]uint
	AbstractAIXMFeatureType_orderStaged     map[uint]*AbstractAIXMFeatureType
	AbstractAIXMFeatureTypes_reference      map[*AbstractAIXMFeatureType]*AbstractAIXMFeatureType
	AbstractAIXMFeatureTypes_referenceOrder map[*AbstractAIXMFeatureType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMFeatureTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMFeatureType]
	OnAfterAbstractAIXMFeatureTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMFeatureType]
	OnAfterAbstractAIXMFeatureTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMFeatureType]

	AbstractAIXMMessageBaseTypes                map[*AbstractAIXMMessageBaseType]struct{}
	AbstractAIXMMessageBaseTypes_instance       map[*AbstractAIXMMessageBaseType]*AbstractAIXMMessageBaseType
	AbstractAIXMMessageBaseTypes_mapString      map[string]*AbstractAIXMMessageBaseType
	AbstractAIXMMessageBaseTypeOrder            uint
	AbstractAIXMMessageBaseType_stagedOrder     map[*AbstractAIXMMessageBaseType]uint
	AbstractAIXMMessageBaseType_orderStaged     map[uint]*AbstractAIXMMessageBaseType
	AbstractAIXMMessageBaseTypes_reference      map[*AbstractAIXMMessageBaseType]*AbstractAIXMMessageBaseType
	AbstractAIXMMessageBaseTypes_referenceOrder map[*AbstractAIXMMessageBaseType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMMessageBaseTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMMessageBaseType]
	OnAfterAbstractAIXMMessageBaseTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMMessageBaseType]
	OnAfterAbstractAIXMMessageBaseTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMMessageBaseType]

	AbstractAIXMMessageTypes                map[*AbstractAIXMMessageType]struct{}
	AbstractAIXMMessageTypes_instance       map[*AbstractAIXMMessageType]*AbstractAIXMMessageType
	AbstractAIXMMessageTypes_mapString      map[string]*AbstractAIXMMessageType
	AbstractAIXMMessageTypeOrder            uint
	AbstractAIXMMessageType_stagedOrder     map[*AbstractAIXMMessageType]uint
	AbstractAIXMMessageType_orderStaged     map[uint]*AbstractAIXMMessageType
	AbstractAIXMMessageTypes_reference      map[*AbstractAIXMMessageType]*AbstractAIXMMessageType
	AbstractAIXMMessageTypes_referenceOrder map[*AbstractAIXMMessageType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMMessageTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMMessageType]
	OnAfterAbstractAIXMMessageTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMMessageType]
	OnAfterAbstractAIXMMessageTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMMessageType]

	AbstractAIXMObjectTypes                map[*AbstractAIXMObjectType]struct{}
	AbstractAIXMObjectTypes_instance       map[*AbstractAIXMObjectType]*AbstractAIXMObjectType
	AbstractAIXMObjectTypes_mapString      map[string]*AbstractAIXMObjectType
	AbstractAIXMObjectTypeOrder            uint
	AbstractAIXMObjectType_stagedOrder     map[*AbstractAIXMObjectType]uint
	AbstractAIXMObjectType_orderStaged     map[uint]*AbstractAIXMObjectType
	AbstractAIXMObjectTypes_reference      map[*AbstractAIXMObjectType]*AbstractAIXMObjectType
	AbstractAIXMObjectTypes_referenceOrder map[*AbstractAIXMObjectType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMObjectTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMObjectType]
	OnAfterAbstractAIXMObjectTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMObjectType]
	OnAfterAbstractAIXMObjectTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMObjectType]

	AbstractAIXMPropertyTypes                map[*AbstractAIXMPropertyType]struct{}
	AbstractAIXMPropertyTypes_instance       map[*AbstractAIXMPropertyType]*AbstractAIXMPropertyType
	AbstractAIXMPropertyTypes_mapString      map[string]*AbstractAIXMPropertyType
	AbstractAIXMPropertyTypeOrder            uint
	AbstractAIXMPropertyType_stagedOrder     map[*AbstractAIXMPropertyType]uint
	AbstractAIXMPropertyType_orderStaged     map[uint]*AbstractAIXMPropertyType
	AbstractAIXMPropertyTypes_reference      map[*AbstractAIXMPropertyType]*AbstractAIXMPropertyType
	AbstractAIXMPropertyTypes_referenceOrder map[*AbstractAIXMPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMPropertyTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMPropertyType]
	OnAfterAbstractAIXMPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMPropertyType]
	OnAfterAbstractAIXMPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMPropertyType]

	AbstractAIXMTimeSliceBaseTypes                map[*AbstractAIXMTimeSliceBaseType]struct{}
	AbstractAIXMTimeSliceBaseTypes_instance       map[*AbstractAIXMTimeSliceBaseType]*AbstractAIXMTimeSliceBaseType
	AbstractAIXMTimeSliceBaseTypes_mapString      map[string]*AbstractAIXMTimeSliceBaseType
	AbstractAIXMTimeSliceBaseTypeOrder            uint
	AbstractAIXMTimeSliceBaseType_stagedOrder     map[*AbstractAIXMTimeSliceBaseType]uint
	AbstractAIXMTimeSliceBaseType_orderStaged     map[uint]*AbstractAIXMTimeSliceBaseType
	AbstractAIXMTimeSliceBaseTypes_reference      map[*AbstractAIXMTimeSliceBaseType]*AbstractAIXMTimeSliceBaseType
	AbstractAIXMTimeSliceBaseTypes_referenceOrder map[*AbstractAIXMTimeSliceBaseType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMTimeSliceBaseTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMTimeSliceBaseType]
	OnAfterAbstractAIXMTimeSliceBaseTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMTimeSliceBaseType]
	OnAfterAbstractAIXMTimeSliceBaseTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMTimeSliceBaseType]

	AbstractAIXMTimeSliceTypes                map[*AbstractAIXMTimeSliceType]struct{}
	AbstractAIXMTimeSliceTypes_instance       map[*AbstractAIXMTimeSliceType]*AbstractAIXMTimeSliceType
	AbstractAIXMTimeSliceTypes_mapString      map[string]*AbstractAIXMTimeSliceType
	AbstractAIXMTimeSliceTypeOrder            uint
	AbstractAIXMTimeSliceType_stagedOrder     map[*AbstractAIXMTimeSliceType]uint
	AbstractAIXMTimeSliceType_orderStaged     map[uint]*AbstractAIXMTimeSliceType
	AbstractAIXMTimeSliceTypes_reference      map[*AbstractAIXMTimeSliceType]*AbstractAIXMTimeSliceType
	AbstractAIXMTimeSliceTypes_referenceOrder map[*AbstractAIXMTimeSliceType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractAIXMTimeSliceTypeCreateCallback GongOnAfterCreateInterface[AbstractAIXMTimeSliceType]
	OnAfterAbstractAIXMTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[AbstractAIXMTimeSliceType]
	OnAfterAbstractAIXMTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[AbstractAIXMTimeSliceType]

	AbstractExtensionTypes                map[*AbstractExtensionType]struct{}
	AbstractExtensionTypes_instance       map[*AbstractExtensionType]*AbstractExtensionType
	AbstractExtensionTypes_mapString      map[string]*AbstractExtensionType
	AbstractExtensionTypeOrder            uint
	AbstractExtensionType_stagedOrder     map[*AbstractExtensionType]uint
	AbstractExtensionType_orderStaged     map[uint]*AbstractExtensionType
	AbstractExtensionTypes_reference      map[*AbstractExtensionType]*AbstractExtensionType
	AbstractExtensionTypes_referenceOrder map[*AbstractExtensionType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractExtensionTypeCreateCallback GongOnAfterCreateInterface[AbstractExtensionType]
	OnAfterAbstractExtensionTypeUpdateCallback GongOnAfterUpdateInterface[AbstractExtensionType]
	OnAfterAbstractExtensionTypeDeleteCallback GongOnAfterDeleteInterface[AbstractExtensionType]

	FeatureMetadataPropertyTypes                map[*FeatureMetadataPropertyType]struct{}
	FeatureMetadataPropertyTypes_instance       map[*FeatureMetadataPropertyType]*FeatureMetadataPropertyType
	FeatureMetadataPropertyTypes_mapString      map[string]*FeatureMetadataPropertyType
	FeatureMetadataPropertyTypeOrder            uint
	FeatureMetadataPropertyType_stagedOrder     map[*FeatureMetadataPropertyType]uint
	FeatureMetadataPropertyType_orderStaged     map[uint]*FeatureMetadataPropertyType
	FeatureMetadataPropertyTypes_reference      map[*FeatureMetadataPropertyType]*FeatureMetadataPropertyType
	FeatureMetadataPropertyTypes_referenceOrder map[*FeatureMetadataPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterFeatureMetadataPropertyTypeCreateCallback GongOnAfterCreateInterface[FeatureMetadataPropertyType]
	OnAfterFeatureMetadataPropertyTypeUpdateCallback GongOnAfterUpdateInterface[FeatureMetadataPropertyType]
	OnAfterFeatureMetadataPropertyTypeDeleteCallback GongOnAfterDeleteInterface[FeatureMetadataPropertyType]

	FeatureTimeSliceMetadataPropertyTypes                map[*FeatureTimeSliceMetadataPropertyType]struct{}
	FeatureTimeSliceMetadataPropertyTypes_instance       map[*FeatureTimeSliceMetadataPropertyType]*FeatureTimeSliceMetadataPropertyType
	FeatureTimeSliceMetadataPropertyTypes_mapString      map[string]*FeatureTimeSliceMetadataPropertyType
	FeatureTimeSliceMetadataPropertyTypeOrder            uint
	FeatureTimeSliceMetadataPropertyType_stagedOrder     map[*FeatureTimeSliceMetadataPropertyType]uint
	FeatureTimeSliceMetadataPropertyType_orderStaged     map[uint]*FeatureTimeSliceMetadataPropertyType
	FeatureTimeSliceMetadataPropertyTypes_reference      map[*FeatureTimeSliceMetadataPropertyType]*FeatureTimeSliceMetadataPropertyType
	FeatureTimeSliceMetadataPropertyTypes_referenceOrder map[*FeatureTimeSliceMetadataPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterFeatureTimeSliceMetadataPropertyTypeCreateCallback GongOnAfterCreateInterface[FeatureTimeSliceMetadataPropertyType]
	OnAfterFeatureTimeSliceMetadataPropertyTypeUpdateCallback GongOnAfterUpdateInterface[FeatureTimeSliceMetadataPropertyType]
	OnAfterFeatureTimeSliceMetadataPropertyTypeDeleteCallback GongOnAfterDeleteInterface[FeatureTimeSliceMetadataPropertyType]

	MessageMetadataPropertyTypes                map[*MessageMetadataPropertyType]struct{}
	MessageMetadataPropertyTypes_instance       map[*MessageMetadataPropertyType]*MessageMetadataPropertyType
	MessageMetadataPropertyTypes_mapString      map[string]*MessageMetadataPropertyType
	MessageMetadataPropertyTypeOrder            uint
	MessageMetadataPropertyType_stagedOrder     map[*MessageMetadataPropertyType]uint
	MessageMetadataPropertyType_orderStaged     map[uint]*MessageMetadataPropertyType
	MessageMetadataPropertyTypes_reference      map[*MessageMetadataPropertyType]*MessageMetadataPropertyType
	MessageMetadataPropertyTypes_referenceOrder map[*MessageMetadataPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMessageMetadataPropertyTypeCreateCallback GongOnAfterCreateInterface[MessageMetadataPropertyType]
	OnAfterMessageMetadataPropertyTypeUpdateCallback GongOnAfterUpdateInterface[MessageMetadataPropertyType]
	OnAfterMessageMetadataPropertyTypeDeleteCallback GongOnAfterDeleteInterface[MessageMetadataPropertyType]

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
	__gong__clearReferences(&stage.AbstractAIXMFeatureBaseTypes_reference, &stage.AbstractAIXMFeatureBaseTypes_instance, &stage.AbstractAIXMFeatureBaseTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMFeatureTypes_reference, &stage.AbstractAIXMFeatureTypes_instance, &stage.AbstractAIXMFeatureTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMMessageBaseTypes_reference, &stage.AbstractAIXMMessageBaseTypes_instance, &stage.AbstractAIXMMessageBaseTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMMessageTypes_reference, &stage.AbstractAIXMMessageTypes_instance, &stage.AbstractAIXMMessageTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMObjectTypes_reference, &stage.AbstractAIXMObjectTypes_instance, &stage.AbstractAIXMObjectTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMPropertyTypes_reference, &stage.AbstractAIXMPropertyTypes_instance, &stage.AbstractAIXMPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMTimeSliceBaseTypes_reference, &stage.AbstractAIXMTimeSliceBaseTypes_instance, &stage.AbstractAIXMTimeSliceBaseTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractAIXMTimeSliceTypes_reference, &stage.AbstractAIXMTimeSliceTypes_instance, &stage.AbstractAIXMTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractExtensionTypes_reference, &stage.AbstractExtensionTypes_instance, &stage.AbstractExtensionTypes_referenceOrder)

	__gong__clearReferences(&stage.FeatureMetadataPropertyTypes_reference, &stage.FeatureMetadataPropertyTypes_instance, &stage.FeatureMetadataPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.FeatureTimeSliceMetadataPropertyTypes_reference, &stage.FeatureTimeSliceMetadataPropertyTypes_instance, &stage.FeatureTimeSliceMetadataPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MessageMetadataPropertyTypes_reference, &stage.MessageMetadataPropertyTypes_instance, &stage.MessageMetadataPropertyTypes_referenceOrder)

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
	stage.AbstractAIXMFeatureBaseTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMFeatureBaseType_stagedOrder)

	stage.AbstractAIXMFeatureTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMFeatureType_stagedOrder)

	stage.AbstractAIXMMessageBaseTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMMessageBaseType_stagedOrder)

	stage.AbstractAIXMMessageTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMMessageType_stagedOrder)

	stage.AbstractAIXMObjectTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMObjectType_stagedOrder)

	stage.AbstractAIXMPropertyTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMPropertyType_stagedOrder)

	stage.AbstractAIXMTimeSliceBaseTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMTimeSliceBaseType_stagedOrder)

	stage.AbstractAIXMTimeSliceTypeOrder = __gong__recomputeOrder(stage.AbstractAIXMTimeSliceType_stagedOrder)

	stage.AbstractExtensionTypeOrder = __gong__recomputeOrder(stage.AbstractExtensionType_stagedOrder)

	stage.FeatureMetadataPropertyTypeOrder = __gong__recomputeOrder(stage.FeatureMetadataPropertyType_stagedOrder)

	stage.FeatureTimeSliceMetadataPropertyTypeOrder = __gong__recomputeOrder(stage.FeatureTimeSliceMetadataPropertyType_stagedOrder)

	stage.MessageMetadataPropertyTypeOrder = __gong__recomputeOrder(stage.MessageMetadataPropertyType_stagedOrder)

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
func (*AbstractAIXMFeatureBaseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseType_stagedOrder)
}

func (*AbstractAIXMFeatureBaseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMFeatureBaseType_orderStaged[order]
}

func (*AbstractAIXMFeatureBaseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMFeatureBaseTypes_mapString
}

func (*AbstractAIXMFeatureBaseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMFeatureBaseTypes
}

func (*AbstractAIXMFeatureBaseType) GongNewInstance() any {
	return new(AbstractAIXMFeatureBaseType)
}

func (*AbstractAIXMFeatureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureType_stagedOrder)
}

func (*AbstractAIXMFeatureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMFeatureType_orderStaged[order]
}

func (*AbstractAIXMFeatureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMFeatureTypes_mapString
}

func (*AbstractAIXMFeatureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMFeatureTypes
}

func (*AbstractAIXMFeatureType) GongNewInstance() any {
	return new(AbstractAIXMFeatureType)
}

func (*AbstractAIXMMessageBaseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseType_stagedOrder)
}

func (*AbstractAIXMMessageBaseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMMessageBaseType_orderStaged[order]
}

func (*AbstractAIXMMessageBaseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMMessageBaseTypes_mapString
}

func (*AbstractAIXMMessageBaseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMMessageBaseTypes
}

func (*AbstractAIXMMessageBaseType) GongNewInstance() any {
	return new(AbstractAIXMMessageBaseType)
}

func (*AbstractAIXMMessageType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageType_stagedOrder)
}

func (*AbstractAIXMMessageType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMMessageType_orderStaged[order]
}

func (*AbstractAIXMMessageType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMMessageTypes_mapString
}

func (*AbstractAIXMMessageType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMMessageTypes
}

func (*AbstractAIXMMessageType) GongNewInstance() any {
	return new(AbstractAIXMMessageType)
}

func (*AbstractAIXMObjectType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectType_stagedOrder)
}

func (*AbstractAIXMObjectType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMObjectType_orderStaged[order]
}

func (*AbstractAIXMObjectType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMObjectTypes_mapString
}

func (*AbstractAIXMObjectType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMObjectTypes
}

func (*AbstractAIXMObjectType) GongNewInstance() any {
	return new(AbstractAIXMObjectType)
}

func (*AbstractAIXMPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyType_stagedOrder)
}

func (*AbstractAIXMPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMPropertyType_orderStaged[order]
}

func (*AbstractAIXMPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMPropertyTypes_mapString
}

func (*AbstractAIXMPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMPropertyTypes
}

func (*AbstractAIXMPropertyType) GongNewInstance() any {
	return new(AbstractAIXMPropertyType)
}

func (*AbstractAIXMTimeSliceBaseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseType_stagedOrder)
}

func (*AbstractAIXMTimeSliceBaseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMTimeSliceBaseType_orderStaged[order]
}

func (*AbstractAIXMTimeSliceBaseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMTimeSliceBaseTypes_mapString
}

func (*AbstractAIXMTimeSliceBaseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMTimeSliceBaseTypes
}

func (*AbstractAIXMTimeSliceBaseType) GongNewInstance() any {
	return new(AbstractAIXMTimeSliceBaseType)
}

func (*AbstractAIXMTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceType_stagedOrder)
}

func (*AbstractAIXMTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractAIXMTimeSliceType_orderStaged[order]
}

func (*AbstractAIXMTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractAIXMTimeSliceTypes_mapString
}

func (*AbstractAIXMTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractAIXMTimeSliceTypes
}

func (*AbstractAIXMTimeSliceType) GongNewInstance() any {
	return new(AbstractAIXMTimeSliceType)
}

func (*AbstractExtensionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractExtensionTypes, stage.AbstractExtensionType_stagedOrder)
}

func (*AbstractExtensionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractExtensionType_orderStaged[order]
}

func (*AbstractExtensionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractExtensionTypes_mapString
}

func (*AbstractExtensionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractExtensionTypes
}

func (*AbstractExtensionType) GongNewInstance() any {
	return new(AbstractExtensionType)
}

func (*FeatureMetadataPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyType_stagedOrder)
}

func (*FeatureMetadataPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FeatureMetadataPropertyType_orderStaged[order]
}

func (*FeatureMetadataPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FeatureMetadataPropertyTypes_mapString
}

func (*FeatureMetadataPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.FeatureMetadataPropertyTypes
}

func (*FeatureMetadataPropertyType) GongNewInstance() any {
	return new(FeatureMetadataPropertyType)
}

func (*FeatureTimeSliceMetadataPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyType_stagedOrder)
}

func (*FeatureTimeSliceMetadataPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FeatureTimeSliceMetadataPropertyType_orderStaged[order]
}

func (*FeatureTimeSliceMetadataPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FeatureTimeSliceMetadataPropertyTypes_mapString
}

func (*FeatureTimeSliceMetadataPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.FeatureTimeSliceMetadataPropertyTypes
}

func (*FeatureTimeSliceMetadataPropertyType) GongNewInstance() any {
	return new(FeatureTimeSliceMetadataPropertyType)
}

func (*MessageMetadataPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyType_stagedOrder)
}

func (*MessageMetadataPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MessageMetadataPropertyType_orderStaged[order]
}

func (*MessageMetadataPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MessageMetadataPropertyTypes_mapString
}

func (*MessageMetadataPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MessageMetadataPropertyTypes
}

func (*MessageMetadataPropertyType) GongNewInstance() any {
	return new(MessageMetadataPropertyType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
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
		AbstractAIXMFeatureBaseTypes:           make(map[*AbstractAIXMFeatureBaseType]struct{}),
		AbstractAIXMFeatureBaseTypes_mapString: make(map[string]*AbstractAIXMFeatureBaseType),

		AbstractAIXMFeatureTypes:           make(map[*AbstractAIXMFeatureType]struct{}),
		AbstractAIXMFeatureTypes_mapString: make(map[string]*AbstractAIXMFeatureType),

		AbstractAIXMMessageBaseTypes:           make(map[*AbstractAIXMMessageBaseType]struct{}),
		AbstractAIXMMessageBaseTypes_mapString: make(map[string]*AbstractAIXMMessageBaseType),

		AbstractAIXMMessageTypes:           make(map[*AbstractAIXMMessageType]struct{}),
		AbstractAIXMMessageTypes_mapString: make(map[string]*AbstractAIXMMessageType),

		AbstractAIXMObjectTypes:           make(map[*AbstractAIXMObjectType]struct{}),
		AbstractAIXMObjectTypes_mapString: make(map[string]*AbstractAIXMObjectType),

		AbstractAIXMPropertyTypes:           make(map[*AbstractAIXMPropertyType]struct{}),
		AbstractAIXMPropertyTypes_mapString: make(map[string]*AbstractAIXMPropertyType),

		AbstractAIXMTimeSliceBaseTypes:           make(map[*AbstractAIXMTimeSliceBaseType]struct{}),
		AbstractAIXMTimeSliceBaseTypes_mapString: make(map[string]*AbstractAIXMTimeSliceBaseType),

		AbstractAIXMTimeSliceTypes:           make(map[*AbstractAIXMTimeSliceType]struct{}),
		AbstractAIXMTimeSliceTypes_mapString: make(map[string]*AbstractAIXMTimeSliceType),

		AbstractExtensionTypes:           make(map[*AbstractExtensionType]struct{}),
		AbstractExtensionTypes_mapString: make(map[string]*AbstractExtensionType),

		FeatureMetadataPropertyTypes:           make(map[*FeatureMetadataPropertyType]struct{}),
		FeatureMetadataPropertyTypes_mapString: make(map[string]*FeatureMetadataPropertyType),

		FeatureTimeSliceMetadataPropertyTypes:           make(map[*FeatureTimeSliceMetadataPropertyType]struct{}),
		FeatureTimeSliceMetadataPropertyTypes_mapString: make(map[string]*FeatureTimeSliceMetadataPropertyType),

		MessageMetadataPropertyTypes:           make(map[*MessageMetadataPropertyType]struct{}),
		MessageMetadataPropertyTypes_mapString: make(map[string]*MessageMetadataPropertyType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractAIXMFeatureBaseType_stagedOrder: make(map[*AbstractAIXMFeatureBaseType]uint),
		AbstractAIXMFeatureBaseType_orderStaged: make(map[uint]*AbstractAIXMFeatureBaseType),
		AbstractAIXMFeatureBaseTypes_reference:  make(map[*AbstractAIXMFeatureBaseType]*AbstractAIXMFeatureBaseType),

		AbstractAIXMFeatureType_stagedOrder: make(map[*AbstractAIXMFeatureType]uint),
		AbstractAIXMFeatureType_orderStaged: make(map[uint]*AbstractAIXMFeatureType),
		AbstractAIXMFeatureTypes_reference:  make(map[*AbstractAIXMFeatureType]*AbstractAIXMFeatureType),

		AbstractAIXMMessageBaseType_stagedOrder: make(map[*AbstractAIXMMessageBaseType]uint),
		AbstractAIXMMessageBaseType_orderStaged: make(map[uint]*AbstractAIXMMessageBaseType),
		AbstractAIXMMessageBaseTypes_reference:  make(map[*AbstractAIXMMessageBaseType]*AbstractAIXMMessageBaseType),

		AbstractAIXMMessageType_stagedOrder: make(map[*AbstractAIXMMessageType]uint),
		AbstractAIXMMessageType_orderStaged: make(map[uint]*AbstractAIXMMessageType),
		AbstractAIXMMessageTypes_reference:  make(map[*AbstractAIXMMessageType]*AbstractAIXMMessageType),

		AbstractAIXMObjectType_stagedOrder: make(map[*AbstractAIXMObjectType]uint),
		AbstractAIXMObjectType_orderStaged: make(map[uint]*AbstractAIXMObjectType),
		AbstractAIXMObjectTypes_reference:  make(map[*AbstractAIXMObjectType]*AbstractAIXMObjectType),

		AbstractAIXMPropertyType_stagedOrder: make(map[*AbstractAIXMPropertyType]uint),
		AbstractAIXMPropertyType_orderStaged: make(map[uint]*AbstractAIXMPropertyType),
		AbstractAIXMPropertyTypes_reference:  make(map[*AbstractAIXMPropertyType]*AbstractAIXMPropertyType),

		AbstractAIXMTimeSliceBaseType_stagedOrder: make(map[*AbstractAIXMTimeSliceBaseType]uint),
		AbstractAIXMTimeSliceBaseType_orderStaged: make(map[uint]*AbstractAIXMTimeSliceBaseType),
		AbstractAIXMTimeSliceBaseTypes_reference:  make(map[*AbstractAIXMTimeSliceBaseType]*AbstractAIXMTimeSliceBaseType),

		AbstractAIXMTimeSliceType_stagedOrder: make(map[*AbstractAIXMTimeSliceType]uint),
		AbstractAIXMTimeSliceType_orderStaged: make(map[uint]*AbstractAIXMTimeSliceType),
		AbstractAIXMTimeSliceTypes_reference:  make(map[*AbstractAIXMTimeSliceType]*AbstractAIXMTimeSliceType),

		AbstractExtensionType_stagedOrder: make(map[*AbstractExtensionType]uint),
		AbstractExtensionType_orderStaged: make(map[uint]*AbstractExtensionType),
		AbstractExtensionTypes_reference:  make(map[*AbstractExtensionType]*AbstractExtensionType),

		FeatureMetadataPropertyType_stagedOrder: make(map[*FeatureMetadataPropertyType]uint),
		FeatureMetadataPropertyType_orderStaged: make(map[uint]*FeatureMetadataPropertyType),
		FeatureMetadataPropertyTypes_reference:  make(map[*FeatureMetadataPropertyType]*FeatureMetadataPropertyType),

		FeatureTimeSliceMetadataPropertyType_stagedOrder: make(map[*FeatureTimeSliceMetadataPropertyType]uint),
		FeatureTimeSliceMetadataPropertyType_orderStaged: make(map[uint]*FeatureTimeSliceMetadataPropertyType),
		FeatureTimeSliceMetadataPropertyTypes_reference:  make(map[*FeatureTimeSliceMetadataPropertyType]*FeatureTimeSliceMetadataPropertyType),

		MessageMetadataPropertyType_stagedOrder: make(map[*MessageMetadataPropertyType]uint),
		MessageMetadataPropertyType_orderStaged: make(map[uint]*MessageMetadataPropertyType),
		MessageMetadataPropertyTypes_reference:  make(map[*MessageMetadataPropertyType]*MessageMetadataPropertyType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractAIXMFeatureBaseType": &AbstractAIXMFeatureBaseTypeUnmarshaller{},

			"AbstractAIXMFeatureType": &AbstractAIXMFeatureTypeUnmarshaller{},

			"AbstractAIXMMessageBaseType": &AbstractAIXMMessageBaseTypeUnmarshaller{},

			"AbstractAIXMMessageType": &AbstractAIXMMessageTypeUnmarshaller{},

			"AbstractAIXMObjectType": &AbstractAIXMObjectTypeUnmarshaller{},

			"AbstractAIXMPropertyType": &AbstractAIXMPropertyTypeUnmarshaller{},

			"AbstractAIXMTimeSliceBaseType": &AbstractAIXMTimeSliceBaseTypeUnmarshaller{},

			"AbstractAIXMTimeSliceType": &AbstractAIXMTimeSliceTypeUnmarshaller{},

			"AbstractExtensionType": &AbstractExtensionTypeUnmarshaller{},

			"FeatureMetadataPropertyType": &FeatureMetadataPropertyTypeUnmarshaller{},

			"FeatureTimeSliceMetadataPropertyType": &FeatureTimeSliceMetadataPropertyTypeUnmarshaller{},

			"MessageMetadataPropertyType": &MessageMetadataPropertyTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractAIXMFeatureBaseType"] = len(stage.AbstractAIXMFeatureBaseTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMFeatureType"] = len(stage.AbstractAIXMFeatureTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMMessageBaseType"] = len(stage.AbstractAIXMMessageBaseTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMMessageType"] = len(stage.AbstractAIXMMessageTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMObjectType"] = len(stage.AbstractAIXMObjectTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMPropertyType"] = len(stage.AbstractAIXMPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMTimeSliceBaseType"] = len(stage.AbstractAIXMTimeSliceBaseTypes)
	stage.Map_GongStructName_InstancesNb["AbstractAIXMTimeSliceType"] = len(stage.AbstractAIXMTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["AbstractExtensionType"] = len(stage.AbstractExtensionTypes)
	stage.Map_GongStructName_InstancesNb["FeatureMetadataPropertyType"] = len(stage.FeatureMetadataPropertyTypes)
	stage.Map_GongStructName_InstancesNb["FeatureTimeSliceMetadataPropertyType"] = len(stage.FeatureTimeSliceMetadataPropertyTypes)
	stage.Map_GongStructName_InstancesNb["MessageMetadataPropertyType"] = len(stage.MessageMetadataPropertyTypes)
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
// Stage puts abstractaixmfeaturebasetype to the model stage
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) Stage(stage *Stage) *AbstractAIXMFeatureBaseType {
	__gong__stage(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseType_stagedOrder, stage.AbstractAIXMFeatureBaseType_orderStaged, &stage.AbstractAIXMFeatureBaseTypeOrder, stage.AbstractAIXMFeatureBaseTypes_mapString, abstractaixmfeaturebasetype, abstractaixmfeaturebasetype.Name)
	return abstractaixmfeaturebasetype
}

// StagePreserveOrder puts abstractaixmfeaturebasetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMFeatureBaseTypeOrder
// - update stage.AbstractAIXMFeatureBaseTypeOrder accordingly
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseType_stagedOrder, stage.AbstractAIXMFeatureBaseType_orderStaged, &stage.AbstractAIXMFeatureBaseTypeOrder, stage.AbstractAIXMFeatureBaseTypes_mapString, abstractaixmfeaturebasetype, order, abstractaixmfeaturebasetype.Name)
}

// Unstage removes abstractaixmfeaturebasetype off the model stage
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) Unstage(stage *Stage) *AbstractAIXMFeatureBaseType {
	__gong__unstage(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseTypes_mapString, abstractaixmfeaturebasetype, abstractaixmfeaturebasetype.Name)
	return abstractaixmfeaturebasetype
}

// UnstageVoid removes abstractaixmfeaturebasetype off the model stage
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) UnstageVoid(stage *Stage) {
	abstractaixmfeaturebasetype.Unstage(stage)
}

func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) StageVoid(stage *Stage) {
	abstractaixmfeaturebasetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GetName() (res string) {
	return abstractaixmfeaturebasetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) SetName(name string) {
	abstractaixmfeaturebasetype.Name = name
}

// Stage puts abstractaixmfeaturetype to the model stage
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) Stage(stage *Stage) *AbstractAIXMFeatureType {
	__gong__stage(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureType_stagedOrder, stage.AbstractAIXMFeatureType_orderStaged, &stage.AbstractAIXMFeatureTypeOrder, stage.AbstractAIXMFeatureTypes_mapString, abstractaixmfeaturetype, abstractaixmfeaturetype.Name)
	return abstractaixmfeaturetype
}

// StagePreserveOrder puts abstractaixmfeaturetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMFeatureTypeOrder
// - update stage.AbstractAIXMFeatureTypeOrder accordingly
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureType_stagedOrder, stage.AbstractAIXMFeatureType_orderStaged, &stage.AbstractAIXMFeatureTypeOrder, stage.AbstractAIXMFeatureTypes_mapString, abstractaixmfeaturetype, order, abstractaixmfeaturetype.Name)
}

// Unstage removes abstractaixmfeaturetype off the model stage
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) Unstage(stage *Stage) *AbstractAIXMFeatureType {
	__gong__unstage(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureTypes_mapString, abstractaixmfeaturetype, abstractaixmfeaturetype.Name)
	return abstractaixmfeaturetype
}

// UnstageVoid removes abstractaixmfeaturetype off the model stage
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) UnstageVoid(stage *Stage) {
	abstractaixmfeaturetype.Unstage(stage)
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) StageVoid(stage *Stage) {
	abstractaixmfeaturetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GetName() (res string) {
	return abstractaixmfeaturetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) SetName(name string) {
	abstractaixmfeaturetype.Name = name
}

// Stage puts abstractaixmmessagebasetype to the model stage
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) Stage(stage *Stage) *AbstractAIXMMessageBaseType {
	__gong__stage(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseType_stagedOrder, stage.AbstractAIXMMessageBaseType_orderStaged, &stage.AbstractAIXMMessageBaseTypeOrder, stage.AbstractAIXMMessageBaseTypes_mapString, abstractaixmmessagebasetype, abstractaixmmessagebasetype.Name)
	return abstractaixmmessagebasetype
}

// StagePreserveOrder puts abstractaixmmessagebasetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMMessageBaseTypeOrder
// - update stage.AbstractAIXMMessageBaseTypeOrder accordingly
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseType_stagedOrder, stage.AbstractAIXMMessageBaseType_orderStaged, &stage.AbstractAIXMMessageBaseTypeOrder, stage.AbstractAIXMMessageBaseTypes_mapString, abstractaixmmessagebasetype, order, abstractaixmmessagebasetype.Name)
}

// Unstage removes abstractaixmmessagebasetype off the model stage
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) Unstage(stage *Stage) *AbstractAIXMMessageBaseType {
	__gong__unstage(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseTypes_mapString, abstractaixmmessagebasetype, abstractaixmmessagebasetype.Name)
	return abstractaixmmessagebasetype
}

// UnstageVoid removes abstractaixmmessagebasetype off the model stage
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) UnstageVoid(stage *Stage) {
	abstractaixmmessagebasetype.Unstage(stage)
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) StageVoid(stage *Stage) {
	abstractaixmmessagebasetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GetName() (res string) {
	return abstractaixmmessagebasetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) SetName(name string) {
	abstractaixmmessagebasetype.Name = name
}

// Stage puts abstractaixmmessagetype to the model stage
func (abstractaixmmessagetype *AbstractAIXMMessageType) Stage(stage *Stage) *AbstractAIXMMessageType {
	__gong__stage(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageType_stagedOrder, stage.AbstractAIXMMessageType_orderStaged, &stage.AbstractAIXMMessageTypeOrder, stage.AbstractAIXMMessageTypes_mapString, abstractaixmmessagetype, abstractaixmmessagetype.Name)
	return abstractaixmmessagetype
}

// StagePreserveOrder puts abstractaixmmessagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMMessageTypeOrder
// - update stage.AbstractAIXMMessageTypeOrder accordingly
func (abstractaixmmessagetype *AbstractAIXMMessageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageType_stagedOrder, stage.AbstractAIXMMessageType_orderStaged, &stage.AbstractAIXMMessageTypeOrder, stage.AbstractAIXMMessageTypes_mapString, abstractaixmmessagetype, order, abstractaixmmessagetype.Name)
}

// Unstage removes abstractaixmmessagetype off the model stage
func (abstractaixmmessagetype *AbstractAIXMMessageType) Unstage(stage *Stage) *AbstractAIXMMessageType {
	__gong__unstage(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageTypes_mapString, abstractaixmmessagetype, abstractaixmmessagetype.Name)
	return abstractaixmmessagetype
}

// UnstageVoid removes abstractaixmmessagetype off the model stage
func (abstractaixmmessagetype *AbstractAIXMMessageType) UnstageVoid(stage *Stage) {
	abstractaixmmessagetype.Unstage(stage)
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) StageVoid(stage *Stage) {
	abstractaixmmessagetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmmessagetype *AbstractAIXMMessageType) GetName() (res string) {
	return abstractaixmmessagetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmmessagetype *AbstractAIXMMessageType) SetName(name string) {
	abstractaixmmessagetype.Name = name
}

// Stage puts abstractaixmobjecttype to the model stage
func (abstractaixmobjecttype *AbstractAIXMObjectType) Stage(stage *Stage) *AbstractAIXMObjectType {
	__gong__stage(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectType_stagedOrder, stage.AbstractAIXMObjectType_orderStaged, &stage.AbstractAIXMObjectTypeOrder, stage.AbstractAIXMObjectTypes_mapString, abstractaixmobjecttype, abstractaixmobjecttype.Name)
	return abstractaixmobjecttype
}

// StagePreserveOrder puts abstractaixmobjecttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMObjectTypeOrder
// - update stage.AbstractAIXMObjectTypeOrder accordingly
func (abstractaixmobjecttype *AbstractAIXMObjectType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectType_stagedOrder, stage.AbstractAIXMObjectType_orderStaged, &stage.AbstractAIXMObjectTypeOrder, stage.AbstractAIXMObjectTypes_mapString, abstractaixmobjecttype, order, abstractaixmobjecttype.Name)
}

// Unstage removes abstractaixmobjecttype off the model stage
func (abstractaixmobjecttype *AbstractAIXMObjectType) Unstage(stage *Stage) *AbstractAIXMObjectType {
	__gong__unstage(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectTypes_mapString, abstractaixmobjecttype, abstractaixmobjecttype.Name)
	return abstractaixmobjecttype
}

// UnstageVoid removes abstractaixmobjecttype off the model stage
func (abstractaixmobjecttype *AbstractAIXMObjectType) UnstageVoid(stage *Stage) {
	abstractaixmobjecttype.Unstage(stage)
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) StageVoid(stage *Stage) {
	abstractaixmobjecttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmobjecttype *AbstractAIXMObjectType) GetName() (res string) {
	return abstractaixmobjecttype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmobjecttype *AbstractAIXMObjectType) SetName(name string) {
	abstractaixmobjecttype.Name = name
}

// Stage puts abstractaixmpropertytype to the model stage
func (abstractaixmpropertytype *AbstractAIXMPropertyType) Stage(stage *Stage) *AbstractAIXMPropertyType {
	__gong__stage(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyType_stagedOrder, stage.AbstractAIXMPropertyType_orderStaged, &stage.AbstractAIXMPropertyTypeOrder, stage.AbstractAIXMPropertyTypes_mapString, abstractaixmpropertytype, abstractaixmpropertytype.Name)
	return abstractaixmpropertytype
}

// StagePreserveOrder puts abstractaixmpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMPropertyTypeOrder
// - update stage.AbstractAIXMPropertyTypeOrder accordingly
func (abstractaixmpropertytype *AbstractAIXMPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyType_stagedOrder, stage.AbstractAIXMPropertyType_orderStaged, &stage.AbstractAIXMPropertyTypeOrder, stage.AbstractAIXMPropertyTypes_mapString, abstractaixmpropertytype, order, abstractaixmpropertytype.Name)
}

// Unstage removes abstractaixmpropertytype off the model stage
func (abstractaixmpropertytype *AbstractAIXMPropertyType) Unstage(stage *Stage) *AbstractAIXMPropertyType {
	__gong__unstage(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyTypes_mapString, abstractaixmpropertytype, abstractaixmpropertytype.Name)
	return abstractaixmpropertytype
}

// UnstageVoid removes abstractaixmpropertytype off the model stage
func (abstractaixmpropertytype *AbstractAIXMPropertyType) UnstageVoid(stage *Stage) {
	abstractaixmpropertytype.Unstage(stage)
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) StageVoid(stage *Stage) {
	abstractaixmpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmpropertytype *AbstractAIXMPropertyType) GetName() (res string) {
	return abstractaixmpropertytype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmpropertytype *AbstractAIXMPropertyType) SetName(name string) {
	abstractaixmpropertytype.Name = name
}

// Stage puts abstractaixmtimeslicebasetype to the model stage
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) Stage(stage *Stage) *AbstractAIXMTimeSliceBaseType {
	__gong__stage(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseType_stagedOrder, stage.AbstractAIXMTimeSliceBaseType_orderStaged, &stage.AbstractAIXMTimeSliceBaseTypeOrder, stage.AbstractAIXMTimeSliceBaseTypes_mapString, abstractaixmtimeslicebasetype, abstractaixmtimeslicebasetype.Name)
	return abstractaixmtimeslicebasetype
}

// StagePreserveOrder puts abstractaixmtimeslicebasetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMTimeSliceBaseTypeOrder
// - update stage.AbstractAIXMTimeSliceBaseTypeOrder accordingly
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseType_stagedOrder, stage.AbstractAIXMTimeSliceBaseType_orderStaged, &stage.AbstractAIXMTimeSliceBaseTypeOrder, stage.AbstractAIXMTimeSliceBaseTypes_mapString, abstractaixmtimeslicebasetype, order, abstractaixmtimeslicebasetype.Name)
}

// Unstage removes abstractaixmtimeslicebasetype off the model stage
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) Unstage(stage *Stage) *AbstractAIXMTimeSliceBaseType {
	__gong__unstage(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseTypes_mapString, abstractaixmtimeslicebasetype, abstractaixmtimeslicebasetype.Name)
	return abstractaixmtimeslicebasetype
}

// UnstageVoid removes abstractaixmtimeslicebasetype off the model stage
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) UnstageVoid(stage *Stage) {
	abstractaixmtimeslicebasetype.Unstage(stage)
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) StageVoid(stage *Stage) {
	abstractaixmtimeslicebasetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GetName() (res string) {
	return abstractaixmtimeslicebasetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) SetName(name string) {
	abstractaixmtimeslicebasetype.Name = name
}

// Stage puts abstractaixmtimeslicetype to the model stage
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) Stage(stage *Stage) *AbstractAIXMTimeSliceType {
	__gong__stage(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceType_stagedOrder, stage.AbstractAIXMTimeSliceType_orderStaged, &stage.AbstractAIXMTimeSliceTypeOrder, stage.AbstractAIXMTimeSliceTypes_mapString, abstractaixmtimeslicetype, abstractaixmtimeslicetype.Name)
	return abstractaixmtimeslicetype
}

// StagePreserveOrder puts abstractaixmtimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractAIXMTimeSliceTypeOrder
// - update stage.AbstractAIXMTimeSliceTypeOrder accordingly
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceType_stagedOrder, stage.AbstractAIXMTimeSliceType_orderStaged, &stage.AbstractAIXMTimeSliceTypeOrder, stage.AbstractAIXMTimeSliceTypes_mapString, abstractaixmtimeslicetype, order, abstractaixmtimeslicetype.Name)
}

// Unstage removes abstractaixmtimeslicetype off the model stage
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) Unstage(stage *Stage) *AbstractAIXMTimeSliceType {
	__gong__unstage(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceTypes_mapString, abstractaixmtimeslicetype, abstractaixmtimeslicetype.Name)
	return abstractaixmtimeslicetype
}

// UnstageVoid removes abstractaixmtimeslicetype off the model stage
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) UnstageVoid(stage *Stage) {
	abstractaixmtimeslicetype.Unstage(stage)
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) StageVoid(stage *Stage) {
	abstractaixmtimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GetName() (res string) {
	return abstractaixmtimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) SetName(name string) {
	abstractaixmtimeslicetype.Name = name
}

// Stage puts abstractextensiontype to the model stage
func (abstractextensiontype *AbstractExtensionType) Stage(stage *Stage) *AbstractExtensionType {
	__gong__stage(stage.AbstractExtensionTypes, stage.AbstractExtensionType_stagedOrder, stage.AbstractExtensionType_orderStaged, &stage.AbstractExtensionTypeOrder, stage.AbstractExtensionTypes_mapString, abstractextensiontype, abstractextensiontype.Name)
	return abstractextensiontype
}

// StagePreserveOrder puts abstractextensiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractExtensionTypeOrder
// - update stage.AbstractExtensionTypeOrder accordingly
func (abstractextensiontype *AbstractExtensionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractExtensionTypes, stage.AbstractExtensionType_stagedOrder, stage.AbstractExtensionType_orderStaged, &stage.AbstractExtensionTypeOrder, stage.AbstractExtensionTypes_mapString, abstractextensiontype, order, abstractextensiontype.Name)
}

// Unstage removes abstractextensiontype off the model stage
func (abstractextensiontype *AbstractExtensionType) Unstage(stage *Stage) *AbstractExtensionType {
	__gong__unstage(stage.AbstractExtensionTypes, stage.AbstractExtensionTypes_mapString, abstractextensiontype, abstractextensiontype.Name)
	return abstractextensiontype
}

// UnstageVoid removes abstractextensiontype off the model stage
func (abstractextensiontype *AbstractExtensionType) UnstageVoid(stage *Stage) {
	abstractextensiontype.Unstage(stage)
}

func (abstractextensiontype *AbstractExtensionType) StageVoid(stage *Stage) {
	abstractextensiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractextensiontype *AbstractExtensionType) GetName() (res string) {
	return abstractextensiontype.Name
}

// for satisfaction of GongStruct interface
func (abstractextensiontype *AbstractExtensionType) SetName(name string) {
	abstractextensiontype.Name = name
}

// Stage puts featuremetadatapropertytype to the model stage
func (featuremetadatapropertytype *FeatureMetadataPropertyType) Stage(stage *Stage) *FeatureMetadataPropertyType {
	__gong__stage(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyType_stagedOrder, stage.FeatureMetadataPropertyType_orderStaged, &stage.FeatureMetadataPropertyTypeOrder, stage.FeatureMetadataPropertyTypes_mapString, featuremetadatapropertytype, featuremetadatapropertytype.Name)
	return featuremetadatapropertytype
}

// StagePreserveOrder puts featuremetadatapropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FeatureMetadataPropertyTypeOrder
// - update stage.FeatureMetadataPropertyTypeOrder accordingly
func (featuremetadatapropertytype *FeatureMetadataPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyType_stagedOrder, stage.FeatureMetadataPropertyType_orderStaged, &stage.FeatureMetadataPropertyTypeOrder, stage.FeatureMetadataPropertyTypes_mapString, featuremetadatapropertytype, order, featuremetadatapropertytype.Name)
}

// Unstage removes featuremetadatapropertytype off the model stage
func (featuremetadatapropertytype *FeatureMetadataPropertyType) Unstage(stage *Stage) *FeatureMetadataPropertyType {
	__gong__unstage(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyTypes_mapString, featuremetadatapropertytype, featuremetadatapropertytype.Name)
	return featuremetadatapropertytype
}

// UnstageVoid removes featuremetadatapropertytype off the model stage
func (featuremetadatapropertytype *FeatureMetadataPropertyType) UnstageVoid(stage *Stage) {
	featuremetadatapropertytype.Unstage(stage)
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) StageVoid(stage *Stage) {
	featuremetadatapropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (featuremetadatapropertytype *FeatureMetadataPropertyType) GetName() (res string) {
	return featuremetadatapropertytype.Name
}

// for satisfaction of GongStruct interface
func (featuremetadatapropertytype *FeatureMetadataPropertyType) SetName(name string) {
	featuremetadatapropertytype.Name = name
}

// Stage puts featuretimeslicemetadatapropertytype to the model stage
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) Stage(stage *Stage) *FeatureTimeSliceMetadataPropertyType {
	__gong__stage(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyType_stagedOrder, stage.FeatureTimeSliceMetadataPropertyType_orderStaged, &stage.FeatureTimeSliceMetadataPropertyTypeOrder, stage.FeatureTimeSliceMetadataPropertyTypes_mapString, featuretimeslicemetadatapropertytype, featuretimeslicemetadatapropertytype.Name)
	return featuretimeslicemetadatapropertytype
}

// StagePreserveOrder puts featuretimeslicemetadatapropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FeatureTimeSliceMetadataPropertyTypeOrder
// - update stage.FeatureTimeSliceMetadataPropertyTypeOrder accordingly
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyType_stagedOrder, stage.FeatureTimeSliceMetadataPropertyType_orderStaged, &stage.FeatureTimeSliceMetadataPropertyTypeOrder, stage.FeatureTimeSliceMetadataPropertyTypes_mapString, featuretimeslicemetadatapropertytype, order, featuretimeslicemetadatapropertytype.Name)
}

// Unstage removes featuretimeslicemetadatapropertytype off the model stage
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) Unstage(stage *Stage) *FeatureTimeSliceMetadataPropertyType {
	__gong__unstage(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyTypes_mapString, featuretimeslicemetadatapropertytype, featuretimeslicemetadatapropertytype.Name)
	return featuretimeslicemetadatapropertytype
}

// UnstageVoid removes featuretimeslicemetadatapropertytype off the model stage
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) UnstageVoid(stage *Stage) {
	featuretimeslicemetadatapropertytype.Unstage(stage)
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) StageVoid(stage *Stage) {
	featuretimeslicemetadatapropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GetName() (res string) {
	return featuretimeslicemetadatapropertytype.Name
}

// for satisfaction of GongStruct interface
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) SetName(name string) {
	featuretimeslicemetadatapropertytype.Name = name
}

// Stage puts messagemetadatapropertytype to the model stage
func (messagemetadatapropertytype *MessageMetadataPropertyType) Stage(stage *Stage) *MessageMetadataPropertyType {
	__gong__stage(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyType_stagedOrder, stage.MessageMetadataPropertyType_orderStaged, &stage.MessageMetadataPropertyTypeOrder, stage.MessageMetadataPropertyTypes_mapString, messagemetadatapropertytype, messagemetadatapropertytype.Name)
	return messagemetadatapropertytype
}

// StagePreserveOrder puts messagemetadatapropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageMetadataPropertyTypeOrder
// - update stage.MessageMetadataPropertyTypeOrder accordingly
func (messagemetadatapropertytype *MessageMetadataPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyType_stagedOrder, stage.MessageMetadataPropertyType_orderStaged, &stage.MessageMetadataPropertyTypeOrder, stage.MessageMetadataPropertyTypes_mapString, messagemetadatapropertytype, order, messagemetadatapropertytype.Name)
}

// Unstage removes messagemetadatapropertytype off the model stage
func (messagemetadatapropertytype *MessageMetadataPropertyType) Unstage(stage *Stage) *MessageMetadataPropertyType {
	__gong__unstage(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyTypes_mapString, messagemetadatapropertytype, messagemetadatapropertytype.Name)
	return messagemetadatapropertytype
}

// UnstageVoid removes messagemetadatapropertytype off the model stage
func (messagemetadatapropertytype *MessageMetadataPropertyType) UnstageVoid(stage *Stage) {
	messagemetadatapropertytype.Unstage(stage)
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) StageVoid(stage *Stage) {
	messagemetadatapropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (messagemetadatapropertytype *MessageMetadataPropertyType) GetName() (res string) {
	return messagemetadatapropertytype.Name
}

// for satisfaction of GongStruct interface
func (messagemetadatapropertytype *MessageMetadataPropertyType) SetName(name string) {
	messagemetadatapropertytype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractAIXMFeatureBaseTypes, &stage.AbstractAIXMFeatureBaseTypes_mapString, &stage.AbstractAIXMFeatureBaseType_stagedOrder, &stage.AbstractAIXMFeatureBaseTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMFeatureTypes, &stage.AbstractAIXMFeatureTypes_mapString, &stage.AbstractAIXMFeatureType_stagedOrder, &stage.AbstractAIXMFeatureTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMMessageBaseTypes, &stage.AbstractAIXMMessageBaseTypes_mapString, &stage.AbstractAIXMMessageBaseType_stagedOrder, &stage.AbstractAIXMMessageBaseTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMMessageTypes, &stage.AbstractAIXMMessageTypes_mapString, &stage.AbstractAIXMMessageType_stagedOrder, &stage.AbstractAIXMMessageTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMObjectTypes, &stage.AbstractAIXMObjectTypes_mapString, &stage.AbstractAIXMObjectType_stagedOrder, &stage.AbstractAIXMObjectTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMPropertyTypes, &stage.AbstractAIXMPropertyTypes_mapString, &stage.AbstractAIXMPropertyType_stagedOrder, &stage.AbstractAIXMPropertyTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMTimeSliceBaseTypes, &stage.AbstractAIXMTimeSliceBaseTypes_mapString, &stage.AbstractAIXMTimeSliceBaseType_stagedOrder, &stage.AbstractAIXMTimeSliceBaseTypeOrder)

	__gong__resetStageType(&stage.AbstractAIXMTimeSliceTypes, &stage.AbstractAIXMTimeSliceTypes_mapString, &stage.AbstractAIXMTimeSliceType_stagedOrder, &stage.AbstractAIXMTimeSliceTypeOrder)

	__gong__resetStageType(&stage.AbstractExtensionTypes, &stage.AbstractExtensionTypes_mapString, &stage.AbstractExtensionType_stagedOrder, &stage.AbstractExtensionTypeOrder)

	__gong__resetStageType(&stage.FeatureMetadataPropertyTypes, &stage.FeatureMetadataPropertyTypes_mapString, &stage.FeatureMetadataPropertyType_stagedOrder, &stage.FeatureMetadataPropertyTypeOrder)

	__gong__resetStageType(&stage.FeatureTimeSliceMetadataPropertyTypes, &stage.FeatureTimeSliceMetadataPropertyTypes_mapString, &stage.FeatureTimeSliceMetadataPropertyType_stagedOrder, &stage.FeatureTimeSliceMetadataPropertyTypeOrder)

	__gong__resetStageType(&stage.MessageMetadataPropertyTypes, &stage.MessageMetadataPropertyTypes_mapString, &stage.MessageMetadataPropertyType_stagedOrder, &stage.MessageMetadataPropertyTypeOrder)

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
func (AbstractAIXMFeatureBaseType) GongGetAssociationName() any {
	return &AbstractAIXMFeatureBaseType{
	}
}

func (AbstractAIXMFeatureType) GongGetAssociationName() any {
	return &AbstractAIXMFeatureType{
			FeatureMetadata: &FeatureMetadataPropertyType{Name: "FeatureMetadata"},
	}
}

func (AbstractAIXMMessageBaseType) GongGetAssociationName() any {
	return &AbstractAIXMMessageBaseType{
			MessageMetadata: &MessageMetadataPropertyType{Name: "MessageMetadata"},
	}
}

func (AbstractAIXMMessageType) GongGetAssociationName() any {
	return &AbstractAIXMMessageType{
			MessageMetadata: &MessageMetadataPropertyType{Name: "MessageMetadata"},
	}
}

func (AbstractAIXMObjectType) GongGetAssociationName() any {
	return &AbstractAIXMObjectType{
	}
}

func (AbstractAIXMPropertyType) GongGetAssociationName() any {
	return &AbstractAIXMPropertyType{
	}
}

func (AbstractAIXMTimeSliceBaseType) GongGetAssociationName() any {
	return &AbstractAIXMTimeSliceBaseType{
	}
}

func (AbstractAIXMTimeSliceType) GongGetAssociationName() any {
	return &AbstractAIXMTimeSliceType{
			TimeSliceMetadata: &FeatureTimeSliceMetadataPropertyType{Name: "TimeSliceMetadata"},
	}
}

func (AbstractExtensionType) GongGetAssociationName() any {
	return &AbstractExtensionType{
	}
}

func (FeatureMetadataPropertyType) GongGetAssociationName() any {
	return &FeatureMetadataPropertyType{
	}
}

func (FeatureTimeSliceMetadataPropertyType) GongGetAssociationName() any {
	return &FeatureTimeSliceMetadataPropertyType{
	}
}

func (MessageMetadataPropertyType) GongGetAssociationName() any {
	return &MessageMetadataPropertyType{
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
	// reverse maps of direct associations of AbstractAIXMFeatureBaseType
	case AbstractAIXMFeatureBaseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMFeatureType
	case AbstractAIXMFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		case "FeatureMetadata":
			res := make(map[*FeatureMetadataPropertyType][]*AbstractAIXMFeatureType)
			for abstractaixmfeaturetype := range stage.AbstractAIXMFeatureTypes {
				if abstractaixmfeaturetype.FeatureMetadata != nil {
					featuremetadatapropertytype_ := abstractaixmfeaturetype.FeatureMetadata
					var abstractaixmfeaturetypes []*AbstractAIXMFeatureType
					_, ok := res[featuremetadatapropertytype_]
					if ok {
						abstractaixmfeaturetypes = res[featuremetadatapropertytype_]
					} else {
						abstractaixmfeaturetypes = make([]*AbstractAIXMFeatureType, 0)
					}
					abstractaixmfeaturetypes = append(abstractaixmfeaturetypes, abstractaixmfeaturetype)
					res[featuremetadatapropertytype_] = abstractaixmfeaturetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractAIXMMessageBaseType
	case AbstractAIXMMessageBaseType:
		switch fieldname {
		// insertion point for per direct association field
		case "MessageMetadata":
			res := make(map[*MessageMetadataPropertyType][]*AbstractAIXMMessageBaseType)
			for abstractaixmmessagebasetype := range stage.AbstractAIXMMessageBaseTypes {
				if abstractaixmmessagebasetype.MessageMetadata != nil {
					messagemetadatapropertytype_ := abstractaixmmessagebasetype.MessageMetadata
					var abstractaixmmessagebasetypes []*AbstractAIXMMessageBaseType
					_, ok := res[messagemetadatapropertytype_]
					if ok {
						abstractaixmmessagebasetypes = res[messagemetadatapropertytype_]
					} else {
						abstractaixmmessagebasetypes = make([]*AbstractAIXMMessageBaseType, 0)
					}
					abstractaixmmessagebasetypes = append(abstractaixmmessagebasetypes, abstractaixmmessagebasetype)
					res[messagemetadatapropertytype_] = abstractaixmmessagebasetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractAIXMMessageType
	case AbstractAIXMMessageType:
		switch fieldname {
		// insertion point for per direct association field
		case "MessageMetadata":
			res := make(map[*MessageMetadataPropertyType][]*AbstractAIXMMessageType)
			for abstractaixmmessagetype := range stage.AbstractAIXMMessageTypes {
				if abstractaixmmessagetype.MessageMetadata != nil {
					messagemetadatapropertytype_ := abstractaixmmessagetype.MessageMetadata
					var abstractaixmmessagetypes []*AbstractAIXMMessageType
					_, ok := res[messagemetadatapropertytype_]
					if ok {
						abstractaixmmessagetypes = res[messagemetadatapropertytype_]
					} else {
						abstractaixmmessagetypes = make([]*AbstractAIXMMessageType, 0)
					}
					abstractaixmmessagetypes = append(abstractaixmmessagetypes, abstractaixmmessagetype)
					res[messagemetadatapropertytype_] = abstractaixmmessagetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractAIXMObjectType
	case AbstractAIXMObjectType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMPropertyType
	case AbstractAIXMPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMTimeSliceBaseType
	case AbstractAIXMTimeSliceBaseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMTimeSliceType
	case AbstractAIXMTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSliceMetadata":
			res := make(map[*FeatureTimeSliceMetadataPropertyType][]*AbstractAIXMTimeSliceType)
			for abstractaixmtimeslicetype := range stage.AbstractAIXMTimeSliceTypes {
				if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
					featuretimeslicemetadatapropertytype_ := abstractaixmtimeslicetype.TimeSliceMetadata
					var abstractaixmtimeslicetypes []*AbstractAIXMTimeSliceType
					_, ok := res[featuretimeslicemetadatapropertytype_]
					if ok {
						abstractaixmtimeslicetypes = res[featuretimeslicemetadatapropertytype_]
					} else {
						abstractaixmtimeslicetypes = make([]*AbstractAIXMTimeSliceType, 0)
					}
					abstractaixmtimeslicetypes = append(abstractaixmtimeslicetypes, abstractaixmtimeslicetype)
					res[featuretimeslicemetadatapropertytype_] = abstractaixmtimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractExtensionType
	case AbstractExtensionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FeatureMetadataPropertyType
	case FeatureMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FeatureTimeSliceMetadataPropertyType
	case FeatureTimeSliceMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MessageMetadataPropertyType
	case MessageMetadataPropertyType:
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
	// reverse maps of direct associations of AbstractAIXMFeatureBaseType
	case AbstractAIXMFeatureBaseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMFeatureType
	case AbstractAIXMFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMMessageBaseType
	case AbstractAIXMMessageBaseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMMessageType
	case AbstractAIXMMessageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMObjectType
	case AbstractAIXMObjectType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMPropertyType
	case AbstractAIXMPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMTimeSliceBaseType
	case AbstractAIXMTimeSliceBaseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractAIXMTimeSliceType
	case AbstractAIXMTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractExtensionType
	case AbstractExtensionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FeatureMetadataPropertyType
	case FeatureMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FeatureTimeSliceMetadataPropertyType
	case FeatureTimeSliceMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MessageMetadataPropertyType
	case MessageMetadataPropertyType:
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
func (*AbstractAIXMFeatureBaseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMFeatureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMMessageBaseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMMessageType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMObjectType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMTimeSliceBaseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractAIXMTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractExtensionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FeatureMetadataPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FeatureTimeSliceMetadataPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MessageMetadataPropertyType) GongGetReverseFields() []GongReverseField {
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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FeatureMetadata",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FeatureMetadataPropertyType",
		},
	}
	return
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SequenceNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "MessageMetadata",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MessageMetadataPropertyType",
		},
	}
	return
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SequenceNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "MessageMetadata",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MessageMetadataPropertyType",
		},
		{
			Name:               "AggregationType",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Interpretation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_Interpretation_SimpleType",
		},
		{
			Name:               "SequenceNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "CorrectionNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "TimeSliceMetadata",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FeatureTimeSliceMetadataPropertyType",
		},
	}
	return
}

func (abstractextensiontype *AbstractExtensionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Owns",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Owns",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Owns",
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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmfeaturebasetype.Name
	case "Id":
		res.valueString = abstractaixmfeaturebasetype.Id
	case "Description":
		res.valueString = abstractaixmfeaturebasetype.Description
	}
	return
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmfeaturetype.Name
	case "Id":
		res.valueString = abstractaixmfeaturetype.Id
	case "Description":
		res.valueString = abstractaixmfeaturetype.Description
	case "FeatureMetadata":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractaixmfeaturetype.FeatureMetadata != nil {
			res.valueString = abstractaixmfeaturetype.FeatureMetadata.Name
			res.ids = abstractaixmfeaturetype.FeatureMetadata.GongGetUUID(stage)
		}
	}
	return
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmmessagebasetype.Name
	case "Id":
		res.valueString = abstractaixmmessagebasetype.Id
	case "Description":
		res.valueString = abstractaixmmessagebasetype.Description
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", abstractaixmmessagebasetype.SequenceNumber)
		res.valueInt = abstractaixmmessagebasetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MessageMetadata":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractaixmmessagebasetype.MessageMetadata != nil {
			res.valueString = abstractaixmmessagebasetype.MessageMetadata.Name
			res.ids = abstractaixmmessagebasetype.MessageMetadata.GongGetUUID(stage)
		}
	}
	return
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmmessagetype.Name
	case "Id":
		res.valueString = abstractaixmmessagetype.Id
	case "Description":
		res.valueString = abstractaixmmessagetype.Description
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", abstractaixmmessagetype.SequenceNumber)
		res.valueInt = abstractaixmmessagetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MessageMetadata":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractaixmmessagetype.MessageMetadata != nil {
			res.valueString = abstractaixmmessagetype.MessageMetadata.Name
			res.ids = abstractaixmmessagetype.MessageMetadata.GongGetUUID(stage)
		}
	case "AggregationType":
		res.valueString = abstractaixmmessagetype.AggregationType
	}
	return
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmobjecttype.Name
	case "Id":
		res.valueString = abstractaixmobjecttype.Id
	}
	return
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmpropertytype.Name
	case "NilReason":
		res.valueString = abstractaixmpropertytype.NilReason
	}
	return
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmtimeslicebasetype.Name
	case "Id":
		res.valueString = abstractaixmtimeslicebasetype.Id
	}
	return
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractaixmtimeslicetype.Name
	case "Id":
		res.valueString = abstractaixmtimeslicetype.Id
	case "Description":
		res.valueString = abstractaixmtimeslicetype.Description
	case "Interpretation":
		enum := abstractaixmtimeslicetype.Interpretation
		res.valueString = enum.ToCodeString()
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", abstractaixmtimeslicetype.SequenceNumber)
		res.valueInt = abstractaixmtimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", abstractaixmtimeslicetype.CorrectionNumber)
		res.valueInt = abstractaixmtimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "TimeSliceMetadata":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
			res.valueString = abstractaixmtimeslicetype.TimeSliceMetadata.Name
			res.ids = abstractaixmtimeslicetype.TimeSliceMetadata.GongGetUUID(stage)
		}
	}
	return
}

func (abstractextensiontype *AbstractExtensionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractextensiontype.Name
	case "Id":
		res.valueString = abstractextensiontype.Id
	case "Description":
		res.valueString = abstractextensiontype.Description
	}
	return
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = featuremetadatapropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", featuremetadatapropertytype.Owns)
		res.valueBool = featuremetadatapropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = featuretimeslicemetadatapropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", featuretimeslicemetadatapropertytype.Owns)
		res.valueBool = featuretimeslicemetadatapropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = messagemetadatapropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", messagemetadatapropertytype.Owns)
		res.valueBool = messagemetadatapropertytype.Owns
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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetGongstructName() string {
	return "AbstractAIXMFeatureBaseType"
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetGongstructName() string {
	return "AbstractAIXMFeatureType"
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetGongstructName() string {
	return "AbstractAIXMMessageBaseType"
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetGongstructName() string {
	return "AbstractAIXMMessageType"
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetGongstructName() string {
	return "AbstractAIXMObjectType"
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetGongstructName() string {
	return "AbstractAIXMPropertyType"
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetGongstructName() string {
	return "AbstractAIXMTimeSliceBaseType"
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetGongstructName() string {
	return "AbstractAIXMTimeSliceType"
}

func (abstractextensiontype *AbstractExtensionType) GongGetGongstructName() string {
	return "AbstractExtensionType"
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetGongstructName() string {
	return "FeatureMetadataPropertyType"
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetGongstructName() string {
	return "FeatureTimeSliceMetadataPropertyType"
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetGongstructName() string {
	return "MessageMetadataPropertyType"
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
	__gong__rebuildMapString(stage.AbstractAIXMFeatureBaseTypes, &stage.AbstractAIXMFeatureBaseTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMFeatureTypes, &stage.AbstractAIXMFeatureTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMMessageBaseTypes, &stage.AbstractAIXMMessageBaseTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMMessageTypes, &stage.AbstractAIXMMessageTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMObjectTypes, &stage.AbstractAIXMObjectTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMPropertyTypes, &stage.AbstractAIXMPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMTimeSliceBaseTypes, &stage.AbstractAIXMTimeSliceBaseTypes_mapString)

	__gong__rebuildMapString(stage.AbstractAIXMTimeSliceTypes, &stage.AbstractAIXMTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.AbstractExtensionTypes, &stage.AbstractExtensionTypes_mapString)

	__gong__rebuildMapString(stage.FeatureMetadataPropertyTypes, &stage.FeatureMetadataPropertyTypes_mapString)

	__gong__rebuildMapString(stage.FeatureTimeSliceMetadataPropertyTypes, &stage.FeatureTimeSliceMetadataPropertyTypes_mapString)

	__gong__rebuildMapString(stage.MessageMetadataPropertyTypes, &stage.MessageMetadataPropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
