// generated code - do not edit
package gml

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
	AbstractFeatureMemberTypes                map[*AbstractFeatureMemberType]struct{}
	AbstractFeatureMemberTypes_instance       map[*AbstractFeatureMemberType]*AbstractFeatureMemberType
	AbstractFeatureMemberTypes_mapString      map[string]*AbstractFeatureMemberType
	AbstractFeatureMemberTypeOrder            uint
	AbstractFeatureMemberType_stagedOrder     map[*AbstractFeatureMemberType]uint
	AbstractFeatureMemberType_orderStaged     map[uint]*AbstractFeatureMemberType
	AbstractFeatureMemberTypes_reference      map[*AbstractFeatureMemberType]*AbstractFeatureMemberType
	AbstractFeatureMemberTypes_referenceOrder map[*AbstractFeatureMemberType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractFeatureMemberTypeCreateCallback GongOnAfterCreateInterface[AbstractFeatureMemberType]
	OnAfterAbstractFeatureMemberTypeUpdateCallback GongOnAfterUpdateInterface[AbstractFeatureMemberType]
	OnAfterAbstractFeatureMemberTypeDeleteCallback GongOnAfterDeleteInterface[AbstractFeatureMemberType]

	AbstractFeatureTypes                map[*AbstractFeatureType]struct{}
	AbstractFeatureTypes_instance       map[*AbstractFeatureType]*AbstractFeatureType
	AbstractFeatureTypes_mapString      map[string]*AbstractFeatureType
	AbstractFeatureTypeOrder            uint
	AbstractFeatureType_stagedOrder     map[*AbstractFeatureType]uint
	AbstractFeatureType_orderStaged     map[uint]*AbstractFeatureType
	AbstractFeatureTypes_reference      map[*AbstractFeatureType]*AbstractFeatureType
	AbstractFeatureTypes_referenceOrder map[*AbstractFeatureType]uint

	// insertion point for slice of pointers maps
	AbstractFeatureType_NameXSD_reverseMap map[*CodeType]*AbstractFeatureType

	OnAfterAbstractFeatureTypeCreateCallback GongOnAfterCreateInterface[AbstractFeatureType]
	OnAfterAbstractFeatureTypeUpdateCallback GongOnAfterUpdateInterface[AbstractFeatureType]
	OnAfterAbstractFeatureTypeDeleteCallback GongOnAfterDeleteInterface[AbstractFeatureType]

	AbstractGMLTypes                map[*AbstractGMLType]struct{}
	AbstractGMLTypes_instance       map[*AbstractGMLType]*AbstractGMLType
	AbstractGMLTypes_mapString      map[string]*AbstractGMLType
	AbstractGMLTypeOrder            uint
	AbstractGMLType_stagedOrder     map[*AbstractGMLType]uint
	AbstractGMLType_orderStaged     map[uint]*AbstractGMLType
	AbstractGMLTypes_reference      map[*AbstractGMLType]*AbstractGMLType
	AbstractGMLTypes_referenceOrder map[*AbstractGMLType]uint

	// insertion point for slice of pointers maps
	AbstractGMLType_NameXSD_reverseMap map[*CodeType]*AbstractGMLType

	OnAfterAbstractGMLTypeCreateCallback GongOnAfterCreateInterface[AbstractGMLType]
	OnAfterAbstractGMLTypeUpdateCallback GongOnAfterUpdateInterface[AbstractGMLType]
	OnAfterAbstractGMLTypeDeleteCallback GongOnAfterDeleteInterface[AbstractGMLType]

	AbstractGeometricPrimitiveTypes                map[*AbstractGeometricPrimitiveType]struct{}
	AbstractGeometricPrimitiveTypes_instance       map[*AbstractGeometricPrimitiveType]*AbstractGeometricPrimitiveType
	AbstractGeometricPrimitiveTypes_mapString      map[string]*AbstractGeometricPrimitiveType
	AbstractGeometricPrimitiveTypeOrder            uint
	AbstractGeometricPrimitiveType_stagedOrder     map[*AbstractGeometricPrimitiveType]uint
	AbstractGeometricPrimitiveType_orderStaged     map[uint]*AbstractGeometricPrimitiveType
	AbstractGeometricPrimitiveTypes_reference      map[*AbstractGeometricPrimitiveType]*AbstractGeometricPrimitiveType
	AbstractGeometricPrimitiveTypes_referenceOrder map[*AbstractGeometricPrimitiveType]uint

	// insertion point for slice of pointers maps
	AbstractGeometricPrimitiveType_NameXSD_reverseMap map[*CodeType]*AbstractGeometricPrimitiveType

	OnAfterAbstractGeometricPrimitiveTypeCreateCallback GongOnAfterCreateInterface[AbstractGeometricPrimitiveType]
	OnAfterAbstractGeometricPrimitiveTypeUpdateCallback GongOnAfterUpdateInterface[AbstractGeometricPrimitiveType]
	OnAfterAbstractGeometricPrimitiveTypeDeleteCallback GongOnAfterDeleteInterface[AbstractGeometricPrimitiveType]

	AbstractGeometryTypes                map[*AbstractGeometryType]struct{}
	AbstractGeometryTypes_instance       map[*AbstractGeometryType]*AbstractGeometryType
	AbstractGeometryTypes_mapString      map[string]*AbstractGeometryType
	AbstractGeometryTypeOrder            uint
	AbstractGeometryType_stagedOrder     map[*AbstractGeometryType]uint
	AbstractGeometryType_orderStaged     map[uint]*AbstractGeometryType
	AbstractGeometryTypes_reference      map[*AbstractGeometryType]*AbstractGeometryType
	AbstractGeometryTypes_referenceOrder map[*AbstractGeometryType]uint

	// insertion point for slice of pointers maps
	AbstractGeometryType_NameXSD_reverseMap map[*CodeType]*AbstractGeometryType

	OnAfterAbstractGeometryTypeCreateCallback GongOnAfterCreateInterface[AbstractGeometryType]
	OnAfterAbstractGeometryTypeUpdateCallback GongOnAfterUpdateInterface[AbstractGeometryType]
	OnAfterAbstractGeometryTypeDeleteCallback GongOnAfterDeleteInterface[AbstractGeometryType]

	AbstractMetadataPropertyTypes                map[*AbstractMetadataPropertyType]struct{}
	AbstractMetadataPropertyTypes_instance       map[*AbstractMetadataPropertyType]*AbstractMetadataPropertyType
	AbstractMetadataPropertyTypes_mapString      map[string]*AbstractMetadataPropertyType
	AbstractMetadataPropertyTypeOrder            uint
	AbstractMetadataPropertyType_stagedOrder     map[*AbstractMetadataPropertyType]uint
	AbstractMetadataPropertyType_orderStaged     map[uint]*AbstractMetadataPropertyType
	AbstractMetadataPropertyTypes_reference      map[*AbstractMetadataPropertyType]*AbstractMetadataPropertyType
	AbstractMetadataPropertyTypes_referenceOrder map[*AbstractMetadataPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractMetadataPropertyTypeCreateCallback GongOnAfterCreateInterface[AbstractMetadataPropertyType]
	OnAfterAbstractMetadataPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AbstractMetadataPropertyType]
	OnAfterAbstractMetadataPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AbstractMetadataPropertyType]

	AbstractTimeGeometricPrimitiveTypes                map[*AbstractTimeGeometricPrimitiveType]struct{}
	AbstractTimeGeometricPrimitiveTypes_instance       map[*AbstractTimeGeometricPrimitiveType]*AbstractTimeGeometricPrimitiveType
	AbstractTimeGeometricPrimitiveTypes_mapString      map[string]*AbstractTimeGeometricPrimitiveType
	AbstractTimeGeometricPrimitiveTypeOrder            uint
	AbstractTimeGeometricPrimitiveType_stagedOrder     map[*AbstractTimeGeometricPrimitiveType]uint
	AbstractTimeGeometricPrimitiveType_orderStaged     map[uint]*AbstractTimeGeometricPrimitiveType
	AbstractTimeGeometricPrimitiveTypes_reference      map[*AbstractTimeGeometricPrimitiveType]*AbstractTimeGeometricPrimitiveType
	AbstractTimeGeometricPrimitiveTypes_referenceOrder map[*AbstractTimeGeometricPrimitiveType]uint

	// insertion point for slice of pointers maps
	AbstractTimeGeometricPrimitiveType_NameXSD_reverseMap map[*CodeType]*AbstractTimeGeometricPrimitiveType

	OnAfterAbstractTimeGeometricPrimitiveTypeCreateCallback GongOnAfterCreateInterface[AbstractTimeGeometricPrimitiveType]
	OnAfterAbstractTimeGeometricPrimitiveTypeUpdateCallback GongOnAfterUpdateInterface[AbstractTimeGeometricPrimitiveType]
	OnAfterAbstractTimeGeometricPrimitiveTypeDeleteCallback GongOnAfterDeleteInterface[AbstractTimeGeometricPrimitiveType]

	AbstractTimeObjectTypes                map[*AbstractTimeObjectType]struct{}
	AbstractTimeObjectTypes_instance       map[*AbstractTimeObjectType]*AbstractTimeObjectType
	AbstractTimeObjectTypes_mapString      map[string]*AbstractTimeObjectType
	AbstractTimeObjectTypeOrder            uint
	AbstractTimeObjectType_stagedOrder     map[*AbstractTimeObjectType]uint
	AbstractTimeObjectType_orderStaged     map[uint]*AbstractTimeObjectType
	AbstractTimeObjectTypes_reference      map[*AbstractTimeObjectType]*AbstractTimeObjectType
	AbstractTimeObjectTypes_referenceOrder map[*AbstractTimeObjectType]uint

	// insertion point for slice of pointers maps
	AbstractTimeObjectType_NameXSD_reverseMap map[*CodeType]*AbstractTimeObjectType

	OnAfterAbstractTimeObjectTypeCreateCallback GongOnAfterCreateInterface[AbstractTimeObjectType]
	OnAfterAbstractTimeObjectTypeUpdateCallback GongOnAfterUpdateInterface[AbstractTimeObjectType]
	OnAfterAbstractTimeObjectTypeDeleteCallback GongOnAfterDeleteInterface[AbstractTimeObjectType]

	AbstractTimePrimitiveTypes                map[*AbstractTimePrimitiveType]struct{}
	AbstractTimePrimitiveTypes_instance       map[*AbstractTimePrimitiveType]*AbstractTimePrimitiveType
	AbstractTimePrimitiveTypes_mapString      map[string]*AbstractTimePrimitiveType
	AbstractTimePrimitiveTypeOrder            uint
	AbstractTimePrimitiveType_stagedOrder     map[*AbstractTimePrimitiveType]uint
	AbstractTimePrimitiveType_orderStaged     map[uint]*AbstractTimePrimitiveType
	AbstractTimePrimitiveTypes_reference      map[*AbstractTimePrimitiveType]*AbstractTimePrimitiveType
	AbstractTimePrimitiveTypes_referenceOrder map[*AbstractTimePrimitiveType]uint

	// insertion point for slice of pointers maps
	AbstractTimePrimitiveType_NameXSD_reverseMap map[*CodeType]*AbstractTimePrimitiveType

	OnAfterAbstractTimePrimitiveTypeCreateCallback GongOnAfterCreateInterface[AbstractTimePrimitiveType]
	OnAfterAbstractTimePrimitiveTypeUpdateCallback GongOnAfterUpdateInterface[AbstractTimePrimitiveType]
	OnAfterAbstractTimePrimitiveTypeDeleteCallback GongOnAfterDeleteInterface[AbstractTimePrimitiveType]

	AbstractTimeSliceTypes                map[*AbstractTimeSliceType]struct{}
	AbstractTimeSliceTypes_instance       map[*AbstractTimeSliceType]*AbstractTimeSliceType
	AbstractTimeSliceTypes_mapString      map[string]*AbstractTimeSliceType
	AbstractTimeSliceTypeOrder            uint
	AbstractTimeSliceType_stagedOrder     map[*AbstractTimeSliceType]uint
	AbstractTimeSliceType_orderStaged     map[uint]*AbstractTimeSliceType
	AbstractTimeSliceTypes_reference      map[*AbstractTimeSliceType]*AbstractTimeSliceType
	AbstractTimeSliceTypes_referenceOrder map[*AbstractTimeSliceType]uint

	// insertion point for slice of pointers maps
	AbstractTimeSliceType_NameXSD_reverseMap map[*CodeType]*AbstractTimeSliceType

	OnAfterAbstractTimeSliceTypeCreateCallback GongOnAfterCreateInterface[AbstractTimeSliceType]
	OnAfterAbstractTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[AbstractTimeSliceType]
	OnAfterAbstractTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[AbstractTimeSliceType]

	BoundingShapeTypes                map[*BoundingShapeType]struct{}
	BoundingShapeTypes_instance       map[*BoundingShapeType]*BoundingShapeType
	BoundingShapeTypes_mapString      map[string]*BoundingShapeType
	BoundingShapeTypeOrder            uint
	BoundingShapeType_stagedOrder     map[*BoundingShapeType]uint
	BoundingShapeType_orderStaged     map[uint]*BoundingShapeType
	BoundingShapeTypes_reference      map[*BoundingShapeType]*BoundingShapeType
	BoundingShapeTypes_referenceOrder map[*BoundingShapeType]uint

	// insertion point for slice of pointers maps
	OnAfterBoundingShapeTypeCreateCallback GongOnAfterCreateInterface[BoundingShapeType]
	OnAfterBoundingShapeTypeUpdateCallback GongOnAfterUpdateInterface[BoundingShapeType]
	OnAfterBoundingShapeTypeDeleteCallback GongOnAfterDeleteInterface[BoundingShapeType]

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

	DirectPositionTypes                map[*DirectPositionType]struct{}
	DirectPositionTypes_instance       map[*DirectPositionType]*DirectPositionType
	DirectPositionTypes_mapString      map[string]*DirectPositionType
	DirectPositionTypeOrder            uint
	DirectPositionType_stagedOrder     map[*DirectPositionType]uint
	DirectPositionType_orderStaged     map[uint]*DirectPositionType
	DirectPositionTypes_reference      map[*DirectPositionType]*DirectPositionType
	DirectPositionTypes_referenceOrder map[*DirectPositionType]uint

	// insertion point for slice of pointers maps
	OnAfterDirectPositionTypeCreateCallback GongOnAfterCreateInterface[DirectPositionType]
	OnAfterDirectPositionTypeUpdateCallback GongOnAfterUpdateInterface[DirectPositionType]
	OnAfterDirectPositionTypeDeleteCallback GongOnAfterDeleteInterface[DirectPositionType]

	DynamicFeatureTypes                map[*DynamicFeatureType]struct{}
	DynamicFeatureTypes_instance       map[*DynamicFeatureType]*DynamicFeatureType
	DynamicFeatureTypes_mapString      map[string]*DynamicFeatureType
	DynamicFeatureTypeOrder            uint
	DynamicFeatureType_stagedOrder     map[*DynamicFeatureType]uint
	DynamicFeatureType_orderStaged     map[uint]*DynamicFeatureType
	DynamicFeatureTypes_reference      map[*DynamicFeatureType]*DynamicFeatureType
	DynamicFeatureTypes_referenceOrder map[*DynamicFeatureType]uint

	// insertion point for slice of pointers maps
	DynamicFeatureType_NameXSD_reverseMap map[*CodeType]*DynamicFeatureType

	OnAfterDynamicFeatureTypeCreateCallback GongOnAfterCreateInterface[DynamicFeatureType]
	OnAfterDynamicFeatureTypeUpdateCallback GongOnAfterUpdateInterface[DynamicFeatureType]
	OnAfterDynamicFeatureTypeDeleteCallback GongOnAfterDeleteInterface[DynamicFeatureType]

	EnvelopeTypes                map[*EnvelopeType]struct{}
	EnvelopeTypes_instance       map[*EnvelopeType]*EnvelopeType
	EnvelopeTypes_mapString      map[string]*EnvelopeType
	EnvelopeTypeOrder            uint
	EnvelopeType_stagedOrder     map[*EnvelopeType]uint
	EnvelopeType_orderStaged     map[uint]*EnvelopeType
	EnvelopeTypes_reference      map[*EnvelopeType]*EnvelopeType
	EnvelopeTypes_referenceOrder map[*EnvelopeType]uint

	// insertion point for slice of pointers maps
	OnAfterEnvelopeTypeCreateCallback GongOnAfterCreateInterface[EnvelopeType]
	OnAfterEnvelopeTypeUpdateCallback GongOnAfterUpdateInterface[EnvelopeType]
	OnAfterEnvelopeTypeDeleteCallback GongOnAfterDeleteInterface[EnvelopeType]

	PointTypes                map[*PointType]struct{}
	PointTypes_instance       map[*PointType]*PointType
	PointTypes_mapString      map[string]*PointType
	PointTypeOrder            uint
	PointType_stagedOrder     map[*PointType]uint
	PointType_orderStaged     map[uint]*PointType
	PointTypes_reference      map[*PointType]*PointType
	PointTypes_referenceOrder map[*PointType]uint

	// insertion point for slice of pointers maps
	PointType_NameXSD_reverseMap map[*CodeType]*PointType

	OnAfterPointTypeCreateCallback GongOnAfterCreateInterface[PointType]
	OnAfterPointTypeUpdateCallback GongOnAfterUpdateInterface[PointType]
	OnAfterPointTypeDeleteCallback GongOnAfterDeleteInterface[PointType]

	TimeInstantTypes                map[*TimeInstantType]struct{}
	TimeInstantTypes_instance       map[*TimeInstantType]*TimeInstantType
	TimeInstantTypes_mapString      map[string]*TimeInstantType
	TimeInstantTypeOrder            uint
	TimeInstantType_stagedOrder     map[*TimeInstantType]uint
	TimeInstantType_orderStaged     map[uint]*TimeInstantType
	TimeInstantTypes_reference      map[*TimeInstantType]*TimeInstantType
	TimeInstantTypes_referenceOrder map[*TimeInstantType]uint

	// insertion point for slice of pointers maps
	TimeInstantType_NameXSD_reverseMap map[*CodeType]*TimeInstantType

	OnAfterTimeInstantTypeCreateCallback GongOnAfterCreateInterface[TimeInstantType]
	OnAfterTimeInstantTypeUpdateCallback GongOnAfterUpdateInterface[TimeInstantType]
	OnAfterTimeInstantTypeDeleteCallback GongOnAfterDeleteInterface[TimeInstantType]

	TimePeriodTypes                map[*TimePeriodType]struct{}
	TimePeriodTypes_instance       map[*TimePeriodType]*TimePeriodType
	TimePeriodTypes_mapString      map[string]*TimePeriodType
	TimePeriodTypeOrder            uint
	TimePeriodType_stagedOrder     map[*TimePeriodType]uint
	TimePeriodType_orderStaged     map[uint]*TimePeriodType
	TimePeriodTypes_reference      map[*TimePeriodType]*TimePeriodType
	TimePeriodTypes_referenceOrder map[*TimePeriodType]uint

	// insertion point for slice of pointers maps
	TimePeriodType_NameXSD_reverseMap map[*CodeType]*TimePeriodType

	OnAfterTimePeriodTypeCreateCallback GongOnAfterCreateInterface[TimePeriodType]
	OnAfterTimePeriodTypeUpdateCallback GongOnAfterUpdateInterface[TimePeriodType]
	OnAfterTimePeriodTypeDeleteCallback GongOnAfterDeleteInterface[TimePeriodType]

	TimePositionTypes                map[*TimePositionType]struct{}
	TimePositionTypes_instance       map[*TimePositionType]*TimePositionType
	TimePositionTypes_mapString      map[string]*TimePositionType
	TimePositionTypeOrder            uint
	TimePositionType_stagedOrder     map[*TimePositionType]uint
	TimePositionType_orderStaged     map[uint]*TimePositionType
	TimePositionTypes_reference      map[*TimePositionType]*TimePositionType
	TimePositionTypes_referenceOrder map[*TimePositionType]uint

	// insertion point for slice of pointers maps
	OnAfterTimePositionTypeCreateCallback GongOnAfterCreateInterface[TimePositionType]
	OnAfterTimePositionTypeUpdateCallback GongOnAfterUpdateInterface[TimePositionType]
	OnAfterTimePositionTypeDeleteCallback GongOnAfterDeleteInterface[TimePositionType]

	TimePrimitivePropertyTypes                map[*TimePrimitivePropertyType]struct{}
	TimePrimitivePropertyTypes_instance       map[*TimePrimitivePropertyType]*TimePrimitivePropertyType
	TimePrimitivePropertyTypes_mapString      map[string]*TimePrimitivePropertyType
	TimePrimitivePropertyTypeOrder            uint
	TimePrimitivePropertyType_stagedOrder     map[*TimePrimitivePropertyType]uint
	TimePrimitivePropertyType_orderStaged     map[uint]*TimePrimitivePropertyType
	TimePrimitivePropertyTypes_reference      map[*TimePrimitivePropertyType]*TimePrimitivePropertyType
	TimePrimitivePropertyTypes_referenceOrder map[*TimePrimitivePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterTimePrimitivePropertyTypeCreateCallback GongOnAfterCreateInterface[TimePrimitivePropertyType]
	OnAfterTimePrimitivePropertyTypeUpdateCallback GongOnAfterUpdateInterface[TimePrimitivePropertyType]
	OnAfterTimePrimitivePropertyTypeDeleteCallback GongOnAfterDeleteInterface[TimePrimitivePropertyType]

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
	__gong__clearReferences(&stage.AbstractFeatureMemberTypes_reference, &stage.AbstractFeatureMemberTypes_instance, &stage.AbstractFeatureMemberTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractFeatureTypes_reference, &stage.AbstractFeatureTypes_instance, &stage.AbstractFeatureTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGMLTypes_reference, &stage.AbstractGMLTypes_instance, &stage.AbstractGMLTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGeometricPrimitiveTypes_reference, &stage.AbstractGeometricPrimitiveTypes_instance, &stage.AbstractGeometricPrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGeometryTypes_reference, &stage.AbstractGeometryTypes_instance, &stage.AbstractGeometryTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractMetadataPropertyTypes_reference, &stage.AbstractMetadataPropertyTypes_instance, &stage.AbstractMetadataPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeGeometricPrimitiveTypes_reference, &stage.AbstractTimeGeometricPrimitiveTypes_instance, &stage.AbstractTimeGeometricPrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeObjectTypes_reference, &stage.AbstractTimeObjectTypes_instance, &stage.AbstractTimeObjectTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimePrimitiveTypes_reference, &stage.AbstractTimePrimitiveTypes_instance, &stage.AbstractTimePrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeSliceTypes_reference, &stage.AbstractTimeSliceTypes_instance, &stage.AbstractTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.BoundingShapeTypes_reference, &stage.BoundingShapeTypes_instance, &stage.BoundingShapeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTypes_reference, &stage.CodeTypes_instance, &stage.CodeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWithAuthorityTypes_reference, &stage.CodeWithAuthorityTypes_instance, &stage.CodeWithAuthorityTypes_referenceOrder)

	__gong__clearReferences(&stage.DirectPositionTypes_reference, &stage.DirectPositionTypes_instance, &stage.DirectPositionTypes_referenceOrder)

	__gong__clearReferences(&stage.DynamicFeatureTypes_reference, &stage.DynamicFeatureTypes_instance, &stage.DynamicFeatureTypes_referenceOrder)

	__gong__clearReferences(&stage.EnvelopeTypes_reference, &stage.EnvelopeTypes_instance, &stage.EnvelopeTypes_referenceOrder)

	__gong__clearReferences(&stage.PointTypes_reference, &stage.PointTypes_instance, &stage.PointTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeInstantTypes_reference, &stage.TimeInstantTypes_instance, &stage.TimeInstantTypes_referenceOrder)

	__gong__clearReferences(&stage.TimePeriodTypes_reference, &stage.TimePeriodTypes_instance, &stage.TimePeriodTypes_referenceOrder)

	__gong__clearReferences(&stage.TimePositionTypes_reference, &stage.TimePositionTypes_instance, &stage.TimePositionTypes_referenceOrder)

	__gong__clearReferences(&stage.TimePrimitivePropertyTypes_reference, &stage.TimePrimitivePropertyTypes_instance, &stage.TimePrimitivePropertyTypes_referenceOrder)

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
	stage.AbstractFeatureMemberTypeOrder = __gong__recomputeOrder(stage.AbstractFeatureMemberType_stagedOrder)

	stage.AbstractFeatureTypeOrder = __gong__recomputeOrder(stage.AbstractFeatureType_stagedOrder)

	stage.AbstractGMLTypeOrder = __gong__recomputeOrder(stage.AbstractGMLType_stagedOrder)

	stage.AbstractGeometricPrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractGeometricPrimitiveType_stagedOrder)

	stage.AbstractGeometryTypeOrder = __gong__recomputeOrder(stage.AbstractGeometryType_stagedOrder)

	stage.AbstractMetadataPropertyTypeOrder = __gong__recomputeOrder(stage.AbstractMetadataPropertyType_stagedOrder)

	stage.AbstractTimeGeometricPrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractTimeGeometricPrimitiveType_stagedOrder)

	stage.AbstractTimeObjectTypeOrder = __gong__recomputeOrder(stage.AbstractTimeObjectType_stagedOrder)

	stage.AbstractTimePrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractTimePrimitiveType_stagedOrder)

	stage.AbstractTimeSliceTypeOrder = __gong__recomputeOrder(stage.AbstractTimeSliceType_stagedOrder)

	stage.BoundingShapeTypeOrder = __gong__recomputeOrder(stage.BoundingShapeType_stagedOrder)

	stage.CodeTypeOrder = __gong__recomputeOrder(stage.CodeType_stagedOrder)

	stage.CodeWithAuthorityTypeOrder = __gong__recomputeOrder(stage.CodeWithAuthorityType_stagedOrder)

	stage.DirectPositionTypeOrder = __gong__recomputeOrder(stage.DirectPositionType_stagedOrder)

	stage.DynamicFeatureTypeOrder = __gong__recomputeOrder(stage.DynamicFeatureType_stagedOrder)

	stage.EnvelopeTypeOrder = __gong__recomputeOrder(stage.EnvelopeType_stagedOrder)

	stage.PointTypeOrder = __gong__recomputeOrder(stage.PointType_stagedOrder)

	stage.TimeInstantTypeOrder = __gong__recomputeOrder(stage.TimeInstantType_stagedOrder)

	stage.TimePeriodTypeOrder = __gong__recomputeOrder(stage.TimePeriodType_stagedOrder)

	stage.TimePositionTypeOrder = __gong__recomputeOrder(stage.TimePositionType_stagedOrder)

	stage.TimePrimitivePropertyTypeOrder = __gong__recomputeOrder(stage.TimePrimitivePropertyType_stagedOrder)

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
func (*AbstractFeatureMemberType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberType_stagedOrder)
}

func (*AbstractFeatureMemberType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractFeatureMemberType_orderStaged[order]
}

func (*AbstractFeatureMemberType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractFeatureMemberTypes_mapString
}

func (*AbstractFeatureMemberType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractFeatureMemberTypes
}

func (*AbstractFeatureMemberType) GongNewInstance() any {
	return new(AbstractFeatureMemberType)
}

func (*AbstractFeatureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractFeatureTypes, stage.AbstractFeatureType_stagedOrder)
}

func (*AbstractFeatureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractFeatureType_orderStaged[order]
}

func (*AbstractFeatureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractFeatureTypes_mapString
}

func (*AbstractFeatureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractFeatureTypes
}

func (*AbstractFeatureType) GongNewInstance() any {
	return new(AbstractFeatureType)
}

func (*AbstractGMLType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractGMLTypes, stage.AbstractGMLType_stagedOrder)
}

func (*AbstractGMLType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractGMLType_orderStaged[order]
}

func (*AbstractGMLType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractGMLTypes_mapString
}

func (*AbstractGMLType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractGMLTypes
}

func (*AbstractGMLType) GongNewInstance() any {
	return new(AbstractGMLType)
}

func (*AbstractGeometricPrimitiveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveType_stagedOrder)
}

func (*AbstractGeometricPrimitiveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractGeometricPrimitiveType_orderStaged[order]
}

func (*AbstractGeometricPrimitiveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractGeometricPrimitiveTypes_mapString
}

func (*AbstractGeometricPrimitiveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractGeometricPrimitiveTypes
}

func (*AbstractGeometricPrimitiveType) GongNewInstance() any {
	return new(AbstractGeometricPrimitiveType)
}

func (*AbstractGeometryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractGeometryTypes, stage.AbstractGeometryType_stagedOrder)
}

func (*AbstractGeometryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractGeometryType_orderStaged[order]
}

func (*AbstractGeometryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractGeometryTypes_mapString
}

func (*AbstractGeometryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractGeometryTypes
}

func (*AbstractGeometryType) GongNewInstance() any {
	return new(AbstractGeometryType)
}

func (*AbstractMetadataPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyType_stagedOrder)
}

func (*AbstractMetadataPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractMetadataPropertyType_orderStaged[order]
}

func (*AbstractMetadataPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractMetadataPropertyTypes_mapString
}

func (*AbstractMetadataPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractMetadataPropertyTypes
}

func (*AbstractMetadataPropertyType) GongNewInstance() any {
	return new(AbstractMetadataPropertyType)
}

func (*AbstractTimeGeometricPrimitiveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveType_stagedOrder)
}

func (*AbstractTimeGeometricPrimitiveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractTimeGeometricPrimitiveType_orderStaged[order]
}

func (*AbstractTimeGeometricPrimitiveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractTimeGeometricPrimitiveTypes_mapString
}

func (*AbstractTimeGeometricPrimitiveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractTimeGeometricPrimitiveTypes
}

func (*AbstractTimeGeometricPrimitiveType) GongNewInstance() any {
	return new(AbstractTimeGeometricPrimitiveType)
}

func (*AbstractTimeObjectType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectType_stagedOrder)
}

func (*AbstractTimeObjectType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractTimeObjectType_orderStaged[order]
}

func (*AbstractTimeObjectType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractTimeObjectTypes_mapString
}

func (*AbstractTimeObjectType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractTimeObjectTypes
}

func (*AbstractTimeObjectType) GongNewInstance() any {
	return new(AbstractTimeObjectType)
}

func (*AbstractTimePrimitiveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveType_stagedOrder)
}

func (*AbstractTimePrimitiveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractTimePrimitiveType_orderStaged[order]
}

func (*AbstractTimePrimitiveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractTimePrimitiveTypes_mapString
}

func (*AbstractTimePrimitiveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractTimePrimitiveTypes
}

func (*AbstractTimePrimitiveType) GongNewInstance() any {
	return new(AbstractTimePrimitiveType)
}

func (*AbstractTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceType_stagedOrder)
}

func (*AbstractTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractTimeSliceType_orderStaged[order]
}

func (*AbstractTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractTimeSliceTypes_mapString
}

func (*AbstractTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractTimeSliceTypes
}

func (*AbstractTimeSliceType) GongNewInstance() any {
	return new(AbstractTimeSliceType)
}

func (*BoundingShapeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BoundingShapeTypes, stage.BoundingShapeType_stagedOrder)
}

func (*BoundingShapeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BoundingShapeType_orderStaged[order]
}

func (*BoundingShapeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BoundingShapeTypes_mapString
}

func (*BoundingShapeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.BoundingShapeTypes
}

func (*BoundingShapeType) GongNewInstance() any {
	return new(BoundingShapeType)
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

func (*DirectPositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DirectPositionTypes, stage.DirectPositionType_stagedOrder)
}

func (*DirectPositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DirectPositionType_orderStaged[order]
}

func (*DirectPositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DirectPositionTypes_mapString
}

func (*DirectPositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DirectPositionTypes
}

func (*DirectPositionType) GongNewInstance() any {
	return new(DirectPositionType)
}

func (*DynamicFeatureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DynamicFeatureTypes, stage.DynamicFeatureType_stagedOrder)
}

func (*DynamicFeatureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DynamicFeatureType_orderStaged[order]
}

func (*DynamicFeatureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DynamicFeatureTypes_mapString
}

func (*DynamicFeatureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DynamicFeatureTypes
}

func (*DynamicFeatureType) GongNewInstance() any {
	return new(DynamicFeatureType)
}

func (*EnvelopeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EnvelopeTypes, stage.EnvelopeType_stagedOrder)
}

func (*EnvelopeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EnvelopeType_orderStaged[order]
}

func (*EnvelopeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EnvelopeTypes_mapString
}

func (*EnvelopeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.EnvelopeTypes
}

func (*EnvelopeType) GongNewInstance() any {
	return new(EnvelopeType)
}

func (*PointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PointTypes, stage.PointType_stagedOrder)
}

func (*PointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PointType_orderStaged[order]
}

func (*PointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PointTypes_mapString
}

func (*PointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.PointTypes
}

func (*PointType) GongNewInstance() any {
	return new(PointType)
}

func (*TimeInstantType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimeInstantTypes, stage.TimeInstantType_stagedOrder)
}

func (*TimeInstantType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimeInstantType_orderStaged[order]
}

func (*TimeInstantType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimeInstantTypes_mapString
}

func (*TimeInstantType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimeInstantTypes
}

func (*TimeInstantType) GongNewInstance() any {
	return new(TimeInstantType)
}

func (*TimePeriodType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimePeriodTypes, stage.TimePeriodType_stagedOrder)
}

func (*TimePeriodType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimePeriodType_orderStaged[order]
}

func (*TimePeriodType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimePeriodTypes_mapString
}

func (*TimePeriodType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimePeriodTypes
}

func (*TimePeriodType) GongNewInstance() any {
	return new(TimePeriodType)
}

func (*TimePositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimePositionTypes, stage.TimePositionType_stagedOrder)
}

func (*TimePositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimePositionType_orderStaged[order]
}

func (*TimePositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimePositionTypes_mapString
}

func (*TimePositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimePositionTypes
}

func (*TimePositionType) GongNewInstance() any {
	return new(TimePositionType)
}

func (*TimePrimitivePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyType_stagedOrder)
}

func (*TimePrimitivePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimePrimitivePropertyType_orderStaged[order]
}

func (*TimePrimitivePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimePrimitivePropertyTypes_mapString
}

func (*TimePrimitivePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimePrimitivePropertyTypes
}

func (*TimePrimitivePropertyType) GongNewInstance() any {
	return new(TimePrimitivePropertyType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
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
		AbstractFeatureMemberTypes:           make(map[*AbstractFeatureMemberType]struct{}),
		AbstractFeatureMemberTypes_mapString: make(map[string]*AbstractFeatureMemberType),

		AbstractFeatureTypes:           make(map[*AbstractFeatureType]struct{}),
		AbstractFeatureTypes_mapString: make(map[string]*AbstractFeatureType),

		AbstractGMLTypes:           make(map[*AbstractGMLType]struct{}),
		AbstractGMLTypes_mapString: make(map[string]*AbstractGMLType),

		AbstractGeometricPrimitiveTypes:           make(map[*AbstractGeometricPrimitiveType]struct{}),
		AbstractGeometricPrimitiveTypes_mapString: make(map[string]*AbstractGeometricPrimitiveType),

		AbstractGeometryTypes:           make(map[*AbstractGeometryType]struct{}),
		AbstractGeometryTypes_mapString: make(map[string]*AbstractGeometryType),

		AbstractMetadataPropertyTypes:           make(map[*AbstractMetadataPropertyType]struct{}),
		AbstractMetadataPropertyTypes_mapString: make(map[string]*AbstractMetadataPropertyType),

		AbstractTimeGeometricPrimitiveTypes:           make(map[*AbstractTimeGeometricPrimitiveType]struct{}),
		AbstractTimeGeometricPrimitiveTypes_mapString: make(map[string]*AbstractTimeGeometricPrimitiveType),

		AbstractTimeObjectTypes:           make(map[*AbstractTimeObjectType]struct{}),
		AbstractTimeObjectTypes_mapString: make(map[string]*AbstractTimeObjectType),

		AbstractTimePrimitiveTypes:           make(map[*AbstractTimePrimitiveType]struct{}),
		AbstractTimePrimitiveTypes_mapString: make(map[string]*AbstractTimePrimitiveType),

		AbstractTimeSliceTypes:           make(map[*AbstractTimeSliceType]struct{}),
		AbstractTimeSliceTypes_mapString: make(map[string]*AbstractTimeSliceType),

		BoundingShapeTypes:           make(map[*BoundingShapeType]struct{}),
		BoundingShapeTypes_mapString: make(map[string]*BoundingShapeType),

		CodeTypes:           make(map[*CodeType]struct{}),
		CodeTypes_mapString: make(map[string]*CodeType),

		CodeWithAuthorityTypes:           make(map[*CodeWithAuthorityType]struct{}),
		CodeWithAuthorityTypes_mapString: make(map[string]*CodeWithAuthorityType),

		DirectPositionTypes:           make(map[*DirectPositionType]struct{}),
		DirectPositionTypes_mapString: make(map[string]*DirectPositionType),

		DynamicFeatureTypes:           make(map[*DynamicFeatureType]struct{}),
		DynamicFeatureTypes_mapString: make(map[string]*DynamicFeatureType),

		EnvelopeTypes:           make(map[*EnvelopeType]struct{}),
		EnvelopeTypes_mapString: make(map[string]*EnvelopeType),

		PointTypes:           make(map[*PointType]struct{}),
		PointTypes_mapString: make(map[string]*PointType),

		TimeInstantTypes:           make(map[*TimeInstantType]struct{}),
		TimeInstantTypes_mapString: make(map[string]*TimeInstantType),

		TimePeriodTypes:           make(map[*TimePeriodType]struct{}),
		TimePeriodTypes_mapString: make(map[string]*TimePeriodType),

		TimePositionTypes:           make(map[*TimePositionType]struct{}),
		TimePositionTypes_mapString: make(map[string]*TimePositionType),

		TimePrimitivePropertyTypes:           make(map[*TimePrimitivePropertyType]struct{}),
		TimePrimitivePropertyTypes_mapString: make(map[string]*TimePrimitivePropertyType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractFeatureMemberType_stagedOrder: make(map[*AbstractFeatureMemberType]uint),
		AbstractFeatureMemberType_orderStaged: make(map[uint]*AbstractFeatureMemberType),
		AbstractFeatureMemberTypes_reference:  make(map[*AbstractFeatureMemberType]*AbstractFeatureMemberType),

		AbstractFeatureType_stagedOrder: make(map[*AbstractFeatureType]uint),
		AbstractFeatureType_orderStaged: make(map[uint]*AbstractFeatureType),
		AbstractFeatureTypes_reference:  make(map[*AbstractFeatureType]*AbstractFeatureType),

		AbstractGMLType_stagedOrder: make(map[*AbstractGMLType]uint),
		AbstractGMLType_orderStaged: make(map[uint]*AbstractGMLType),
		AbstractGMLTypes_reference:  make(map[*AbstractGMLType]*AbstractGMLType),

		AbstractGeometricPrimitiveType_stagedOrder: make(map[*AbstractGeometricPrimitiveType]uint),
		AbstractGeometricPrimitiveType_orderStaged: make(map[uint]*AbstractGeometricPrimitiveType),
		AbstractGeometricPrimitiveTypes_reference:  make(map[*AbstractGeometricPrimitiveType]*AbstractGeometricPrimitiveType),

		AbstractGeometryType_stagedOrder: make(map[*AbstractGeometryType]uint),
		AbstractGeometryType_orderStaged: make(map[uint]*AbstractGeometryType),
		AbstractGeometryTypes_reference:  make(map[*AbstractGeometryType]*AbstractGeometryType),

		AbstractMetadataPropertyType_stagedOrder: make(map[*AbstractMetadataPropertyType]uint),
		AbstractMetadataPropertyType_orderStaged: make(map[uint]*AbstractMetadataPropertyType),
		AbstractMetadataPropertyTypes_reference:  make(map[*AbstractMetadataPropertyType]*AbstractMetadataPropertyType),

		AbstractTimeGeometricPrimitiveType_stagedOrder: make(map[*AbstractTimeGeometricPrimitiveType]uint),
		AbstractTimeGeometricPrimitiveType_orderStaged: make(map[uint]*AbstractTimeGeometricPrimitiveType),
		AbstractTimeGeometricPrimitiveTypes_reference:  make(map[*AbstractTimeGeometricPrimitiveType]*AbstractTimeGeometricPrimitiveType),

		AbstractTimeObjectType_stagedOrder: make(map[*AbstractTimeObjectType]uint),
		AbstractTimeObjectType_orderStaged: make(map[uint]*AbstractTimeObjectType),
		AbstractTimeObjectTypes_reference:  make(map[*AbstractTimeObjectType]*AbstractTimeObjectType),

		AbstractTimePrimitiveType_stagedOrder: make(map[*AbstractTimePrimitiveType]uint),
		AbstractTimePrimitiveType_orderStaged: make(map[uint]*AbstractTimePrimitiveType),
		AbstractTimePrimitiveTypes_reference:  make(map[*AbstractTimePrimitiveType]*AbstractTimePrimitiveType),

		AbstractTimeSliceType_stagedOrder: make(map[*AbstractTimeSliceType]uint),
		AbstractTimeSliceType_orderStaged: make(map[uint]*AbstractTimeSliceType),
		AbstractTimeSliceTypes_reference:  make(map[*AbstractTimeSliceType]*AbstractTimeSliceType),

		BoundingShapeType_stagedOrder: make(map[*BoundingShapeType]uint),
		BoundingShapeType_orderStaged: make(map[uint]*BoundingShapeType),
		BoundingShapeTypes_reference:  make(map[*BoundingShapeType]*BoundingShapeType),

		CodeType_stagedOrder: make(map[*CodeType]uint),
		CodeType_orderStaged: make(map[uint]*CodeType),
		CodeTypes_reference:  make(map[*CodeType]*CodeType),

		CodeWithAuthorityType_stagedOrder: make(map[*CodeWithAuthorityType]uint),
		CodeWithAuthorityType_orderStaged: make(map[uint]*CodeWithAuthorityType),
		CodeWithAuthorityTypes_reference:  make(map[*CodeWithAuthorityType]*CodeWithAuthorityType),

		DirectPositionType_stagedOrder: make(map[*DirectPositionType]uint),
		DirectPositionType_orderStaged: make(map[uint]*DirectPositionType),
		DirectPositionTypes_reference:  make(map[*DirectPositionType]*DirectPositionType),

		DynamicFeatureType_stagedOrder: make(map[*DynamicFeatureType]uint),
		DynamicFeatureType_orderStaged: make(map[uint]*DynamicFeatureType),
		DynamicFeatureTypes_reference:  make(map[*DynamicFeatureType]*DynamicFeatureType),

		EnvelopeType_stagedOrder: make(map[*EnvelopeType]uint),
		EnvelopeType_orderStaged: make(map[uint]*EnvelopeType),
		EnvelopeTypes_reference:  make(map[*EnvelopeType]*EnvelopeType),

		PointType_stagedOrder: make(map[*PointType]uint),
		PointType_orderStaged: make(map[uint]*PointType),
		PointTypes_reference:  make(map[*PointType]*PointType),

		TimeInstantType_stagedOrder: make(map[*TimeInstantType]uint),
		TimeInstantType_orderStaged: make(map[uint]*TimeInstantType),
		TimeInstantTypes_reference:  make(map[*TimeInstantType]*TimeInstantType),

		TimePeriodType_stagedOrder: make(map[*TimePeriodType]uint),
		TimePeriodType_orderStaged: make(map[uint]*TimePeriodType),
		TimePeriodTypes_reference:  make(map[*TimePeriodType]*TimePeriodType),

		TimePositionType_stagedOrder: make(map[*TimePositionType]uint),
		TimePositionType_orderStaged: make(map[uint]*TimePositionType),
		TimePositionTypes_reference:  make(map[*TimePositionType]*TimePositionType),

		TimePrimitivePropertyType_stagedOrder: make(map[*TimePrimitivePropertyType]uint),
		TimePrimitivePropertyType_orderStaged: make(map[uint]*TimePrimitivePropertyType),
		TimePrimitivePropertyTypes_reference:  make(map[*TimePrimitivePropertyType]*TimePrimitivePropertyType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractFeatureMemberType": &AbstractFeatureMemberTypeUnmarshaller{},

			"AbstractFeatureType": &AbstractFeatureTypeUnmarshaller{},

			"AbstractGMLType": &AbstractGMLTypeUnmarshaller{},

			"AbstractGeometricPrimitiveType": &AbstractGeometricPrimitiveTypeUnmarshaller{},

			"AbstractGeometryType": &AbstractGeometryTypeUnmarshaller{},

			"AbstractMetadataPropertyType": &AbstractMetadataPropertyTypeUnmarshaller{},

			"AbstractTimeGeometricPrimitiveType": &AbstractTimeGeometricPrimitiveTypeUnmarshaller{},

			"AbstractTimeObjectType": &AbstractTimeObjectTypeUnmarshaller{},

			"AbstractTimePrimitiveType": &AbstractTimePrimitiveTypeUnmarshaller{},

			"AbstractTimeSliceType": &AbstractTimeSliceTypeUnmarshaller{},

			"BoundingShapeType": &BoundingShapeTypeUnmarshaller{},

			"CodeType": &CodeTypeUnmarshaller{},

			"CodeWithAuthorityType": &CodeWithAuthorityTypeUnmarshaller{},

			"DirectPositionType": &DirectPositionTypeUnmarshaller{},

			"DynamicFeatureType": &DynamicFeatureTypeUnmarshaller{},

			"EnvelopeType": &EnvelopeTypeUnmarshaller{},

			"PointType": &PointTypeUnmarshaller{},

			"TimeInstantType": &TimeInstantTypeUnmarshaller{},

			"TimePeriodType": &TimePeriodTypeUnmarshaller{},

			"TimePositionType": &TimePositionTypeUnmarshaller{},

			"TimePrimitivePropertyType": &TimePrimitivePropertyTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractFeatureMemberType"] = len(stage.AbstractFeatureMemberTypes)
	stage.Map_GongStructName_InstancesNb["AbstractFeatureType"] = len(stage.AbstractFeatureTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGMLType"] = len(stage.AbstractGMLTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGeometricPrimitiveType"] = len(stage.AbstractGeometricPrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGeometryType"] = len(stage.AbstractGeometryTypes)
	stage.Map_GongStructName_InstancesNb["AbstractMetadataPropertyType"] = len(stage.AbstractMetadataPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeGeometricPrimitiveType"] = len(stage.AbstractTimeGeometricPrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeObjectType"] = len(stage.AbstractTimeObjectTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimePrimitiveType"] = len(stage.AbstractTimePrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeSliceType"] = len(stage.AbstractTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["BoundingShapeType"] = len(stage.BoundingShapeTypes)
	stage.Map_GongStructName_InstancesNb["CodeType"] = len(stage.CodeTypes)
	stage.Map_GongStructName_InstancesNb["CodeWithAuthorityType"] = len(stage.CodeWithAuthorityTypes)
	stage.Map_GongStructName_InstancesNb["DirectPositionType"] = len(stage.DirectPositionTypes)
	stage.Map_GongStructName_InstancesNb["DynamicFeatureType"] = len(stage.DynamicFeatureTypes)
	stage.Map_GongStructName_InstancesNb["EnvelopeType"] = len(stage.EnvelopeTypes)
	stage.Map_GongStructName_InstancesNb["PointType"] = len(stage.PointTypes)
	stage.Map_GongStructName_InstancesNb["TimeInstantType"] = len(stage.TimeInstantTypes)
	stage.Map_GongStructName_InstancesNb["TimePeriodType"] = len(stage.TimePeriodTypes)
	stage.Map_GongStructName_InstancesNb["TimePositionType"] = len(stage.TimePositionTypes)
	stage.Map_GongStructName_InstancesNb["TimePrimitivePropertyType"] = len(stage.TimePrimitivePropertyTypes)
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
// Stage puts abstractfeaturemembertype to the model stage
func (abstractfeaturemembertype *AbstractFeatureMemberType) Stage(stage *Stage) *AbstractFeatureMemberType {
	__gong__stage(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberType_stagedOrder, stage.AbstractFeatureMemberType_orderStaged, &stage.AbstractFeatureMemberTypeOrder, stage.AbstractFeatureMemberTypes_mapString, abstractfeaturemembertype, abstractfeaturemembertype.Name)
	return abstractfeaturemembertype
}

// StagePreserveOrder puts abstractfeaturemembertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractFeatureMemberTypeOrder
// - update stage.AbstractFeatureMemberTypeOrder accordingly
func (abstractfeaturemembertype *AbstractFeatureMemberType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberType_stagedOrder, stage.AbstractFeatureMemberType_orderStaged, &stage.AbstractFeatureMemberTypeOrder, stage.AbstractFeatureMemberTypes_mapString, abstractfeaturemembertype, order, abstractfeaturemembertype.Name)
}

// Unstage removes abstractfeaturemembertype off the model stage
func (abstractfeaturemembertype *AbstractFeatureMemberType) Unstage(stage *Stage) *AbstractFeatureMemberType {
	__gong__unstage(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberTypes_mapString, abstractfeaturemembertype, abstractfeaturemembertype.Name)
	return abstractfeaturemembertype
}

// UnstageVoid removes abstractfeaturemembertype off the model stage
func (abstractfeaturemembertype *AbstractFeatureMemberType) UnstageVoid(stage *Stage) {
	abstractfeaturemembertype.Unstage(stage)
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) StageVoid(stage *Stage) {
	abstractfeaturemembertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractfeaturemembertype *AbstractFeatureMemberType) GetName() (res string) {
	return abstractfeaturemembertype.Name
}

// for satisfaction of GongStruct interface
func (abstractfeaturemembertype *AbstractFeatureMemberType) SetName(name string) {
	abstractfeaturemembertype.Name = name
}

// Stage puts abstractfeaturetype to the model stage
func (abstractfeaturetype *AbstractFeatureType) Stage(stage *Stage) *AbstractFeatureType {
	__gong__stage(stage.AbstractFeatureTypes, stage.AbstractFeatureType_stagedOrder, stage.AbstractFeatureType_orderStaged, &stage.AbstractFeatureTypeOrder, stage.AbstractFeatureTypes_mapString, abstractfeaturetype, abstractfeaturetype.Name)
	return abstractfeaturetype
}

// StagePreserveOrder puts abstractfeaturetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractFeatureTypeOrder
// - update stage.AbstractFeatureTypeOrder accordingly
func (abstractfeaturetype *AbstractFeatureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractFeatureTypes, stage.AbstractFeatureType_stagedOrder, stage.AbstractFeatureType_orderStaged, &stage.AbstractFeatureTypeOrder, stage.AbstractFeatureTypes_mapString, abstractfeaturetype, order, abstractfeaturetype.Name)
}

// Unstage removes abstractfeaturetype off the model stage
func (abstractfeaturetype *AbstractFeatureType) Unstage(stage *Stage) *AbstractFeatureType {
	__gong__unstage(stage.AbstractFeatureTypes, stage.AbstractFeatureTypes_mapString, abstractfeaturetype, abstractfeaturetype.Name)
	return abstractfeaturetype
}

// UnstageVoid removes abstractfeaturetype off the model stage
func (abstractfeaturetype *AbstractFeatureType) UnstageVoid(stage *Stage) {
	abstractfeaturetype.Unstage(stage)
}

func (abstractfeaturetype *AbstractFeatureType) StageVoid(stage *Stage) {
	abstractfeaturetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractfeaturetype *AbstractFeatureType) GetName() (res string) {
	return abstractfeaturetype.Name
}

// for satisfaction of GongStruct interface
func (abstractfeaturetype *AbstractFeatureType) SetName(name string) {
	abstractfeaturetype.Name = name
}

// Stage puts abstractgmltype to the model stage
func (abstractgmltype *AbstractGMLType) Stage(stage *Stage) *AbstractGMLType {
	__gong__stage(stage.AbstractGMLTypes, stage.AbstractGMLType_stagedOrder, stage.AbstractGMLType_orderStaged, &stage.AbstractGMLTypeOrder, stage.AbstractGMLTypes_mapString, abstractgmltype, abstractgmltype.Name)
	return abstractgmltype
}

// StagePreserveOrder puts abstractgmltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractGMLTypeOrder
// - update stage.AbstractGMLTypeOrder accordingly
func (abstractgmltype *AbstractGMLType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractGMLTypes, stage.AbstractGMLType_stagedOrder, stage.AbstractGMLType_orderStaged, &stage.AbstractGMLTypeOrder, stage.AbstractGMLTypes_mapString, abstractgmltype, order, abstractgmltype.Name)
}

// Unstage removes abstractgmltype off the model stage
func (abstractgmltype *AbstractGMLType) Unstage(stage *Stage) *AbstractGMLType {
	__gong__unstage(stage.AbstractGMLTypes, stage.AbstractGMLTypes_mapString, abstractgmltype, abstractgmltype.Name)
	return abstractgmltype
}

// UnstageVoid removes abstractgmltype off the model stage
func (abstractgmltype *AbstractGMLType) UnstageVoid(stage *Stage) {
	abstractgmltype.Unstage(stage)
}

func (abstractgmltype *AbstractGMLType) StageVoid(stage *Stage) {
	abstractgmltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractgmltype *AbstractGMLType) GetName() (res string) {
	return abstractgmltype.Name
}

// for satisfaction of GongStruct interface
func (abstractgmltype *AbstractGMLType) SetName(name string) {
	abstractgmltype.Name = name
}

// Stage puts abstractgeometricprimitivetype to the model stage
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) Stage(stage *Stage) *AbstractGeometricPrimitiveType {
	__gong__stage(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveType_stagedOrder, stage.AbstractGeometricPrimitiveType_orderStaged, &stage.AbstractGeometricPrimitiveTypeOrder, stage.AbstractGeometricPrimitiveTypes_mapString, abstractgeometricprimitivetype, abstractgeometricprimitivetype.Name)
	return abstractgeometricprimitivetype
}

// StagePreserveOrder puts abstractgeometricprimitivetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractGeometricPrimitiveTypeOrder
// - update stage.AbstractGeometricPrimitiveTypeOrder accordingly
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveType_stagedOrder, stage.AbstractGeometricPrimitiveType_orderStaged, &stage.AbstractGeometricPrimitiveTypeOrder, stage.AbstractGeometricPrimitiveTypes_mapString, abstractgeometricprimitivetype, order, abstractgeometricprimitivetype.Name)
}

// Unstage removes abstractgeometricprimitivetype off the model stage
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) Unstage(stage *Stage) *AbstractGeometricPrimitiveType {
	__gong__unstage(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveTypes_mapString, abstractgeometricprimitivetype, abstractgeometricprimitivetype.Name)
	return abstractgeometricprimitivetype
}

// UnstageVoid removes abstractgeometricprimitivetype off the model stage
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) UnstageVoid(stage *Stage) {
	abstractgeometricprimitivetype.Unstage(stage)
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) StageVoid(stage *Stage) {
	abstractgeometricprimitivetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GetName() (res string) {
	return abstractgeometricprimitivetype.Name
}

// for satisfaction of GongStruct interface
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) SetName(name string) {
	abstractgeometricprimitivetype.Name = name
}

// Stage puts abstractgeometrytype to the model stage
func (abstractgeometrytype *AbstractGeometryType) Stage(stage *Stage) *AbstractGeometryType {
	__gong__stage(stage.AbstractGeometryTypes, stage.AbstractGeometryType_stagedOrder, stage.AbstractGeometryType_orderStaged, &stage.AbstractGeometryTypeOrder, stage.AbstractGeometryTypes_mapString, abstractgeometrytype, abstractgeometrytype.Name)
	return abstractgeometrytype
}

// StagePreserveOrder puts abstractgeometrytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractGeometryTypeOrder
// - update stage.AbstractGeometryTypeOrder accordingly
func (abstractgeometrytype *AbstractGeometryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractGeometryTypes, stage.AbstractGeometryType_stagedOrder, stage.AbstractGeometryType_orderStaged, &stage.AbstractGeometryTypeOrder, stage.AbstractGeometryTypes_mapString, abstractgeometrytype, order, abstractgeometrytype.Name)
}

// Unstage removes abstractgeometrytype off the model stage
func (abstractgeometrytype *AbstractGeometryType) Unstage(stage *Stage) *AbstractGeometryType {
	__gong__unstage(stage.AbstractGeometryTypes, stage.AbstractGeometryTypes_mapString, abstractgeometrytype, abstractgeometrytype.Name)
	return abstractgeometrytype
}

// UnstageVoid removes abstractgeometrytype off the model stage
func (abstractgeometrytype *AbstractGeometryType) UnstageVoid(stage *Stage) {
	abstractgeometrytype.Unstage(stage)
}

func (abstractgeometrytype *AbstractGeometryType) StageVoid(stage *Stage) {
	abstractgeometrytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractgeometrytype *AbstractGeometryType) GetName() (res string) {
	return abstractgeometrytype.Name
}

// for satisfaction of GongStruct interface
func (abstractgeometrytype *AbstractGeometryType) SetName(name string) {
	abstractgeometrytype.Name = name
}

// Stage puts abstractmetadatapropertytype to the model stage
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) Stage(stage *Stage) *AbstractMetadataPropertyType {
	__gong__stage(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyType_stagedOrder, stage.AbstractMetadataPropertyType_orderStaged, &stage.AbstractMetadataPropertyTypeOrder, stage.AbstractMetadataPropertyTypes_mapString, abstractmetadatapropertytype, abstractmetadatapropertytype.Name)
	return abstractmetadatapropertytype
}

// StagePreserveOrder puts abstractmetadatapropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractMetadataPropertyTypeOrder
// - update stage.AbstractMetadataPropertyTypeOrder accordingly
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyType_stagedOrder, stage.AbstractMetadataPropertyType_orderStaged, &stage.AbstractMetadataPropertyTypeOrder, stage.AbstractMetadataPropertyTypes_mapString, abstractmetadatapropertytype, order, abstractmetadatapropertytype.Name)
}

// Unstage removes abstractmetadatapropertytype off the model stage
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) Unstage(stage *Stage) *AbstractMetadataPropertyType {
	__gong__unstage(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyTypes_mapString, abstractmetadatapropertytype, abstractmetadatapropertytype.Name)
	return abstractmetadatapropertytype
}

// UnstageVoid removes abstractmetadatapropertytype off the model stage
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) UnstageVoid(stage *Stage) {
	abstractmetadatapropertytype.Unstage(stage)
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) StageVoid(stage *Stage) {
	abstractmetadatapropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GetName() (res string) {
	return abstractmetadatapropertytype.Name
}

// for satisfaction of GongStruct interface
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) SetName(name string) {
	abstractmetadatapropertytype.Name = name
}

// Stage puts abstracttimegeometricprimitivetype to the model stage
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) Stage(stage *Stage) *AbstractTimeGeometricPrimitiveType {
	__gong__stage(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveType_stagedOrder, stage.AbstractTimeGeometricPrimitiveType_orderStaged, &stage.AbstractTimeGeometricPrimitiveTypeOrder, stage.AbstractTimeGeometricPrimitiveTypes_mapString, abstracttimegeometricprimitivetype, abstracttimegeometricprimitivetype.Name)
	return abstracttimegeometricprimitivetype
}

// StagePreserveOrder puts abstracttimegeometricprimitivetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractTimeGeometricPrimitiveTypeOrder
// - update stage.AbstractTimeGeometricPrimitiveTypeOrder accordingly
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveType_stagedOrder, stage.AbstractTimeGeometricPrimitiveType_orderStaged, &stage.AbstractTimeGeometricPrimitiveTypeOrder, stage.AbstractTimeGeometricPrimitiveTypes_mapString, abstracttimegeometricprimitivetype, order, abstracttimegeometricprimitivetype.Name)
}

// Unstage removes abstracttimegeometricprimitivetype off the model stage
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) Unstage(stage *Stage) *AbstractTimeGeometricPrimitiveType {
	__gong__unstage(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveTypes_mapString, abstracttimegeometricprimitivetype, abstracttimegeometricprimitivetype.Name)
	return abstracttimegeometricprimitivetype
}

// UnstageVoid removes abstracttimegeometricprimitivetype off the model stage
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) UnstageVoid(stage *Stage) {
	abstracttimegeometricprimitivetype.Unstage(stage)
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) StageVoid(stage *Stage) {
	abstracttimegeometricprimitivetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GetName() (res string) {
	return abstracttimegeometricprimitivetype.Name
}

// for satisfaction of GongStruct interface
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) SetName(name string) {
	abstracttimegeometricprimitivetype.Name = name
}

// Stage puts abstracttimeobjecttype to the model stage
func (abstracttimeobjecttype *AbstractTimeObjectType) Stage(stage *Stage) *AbstractTimeObjectType {
	__gong__stage(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectType_stagedOrder, stage.AbstractTimeObjectType_orderStaged, &stage.AbstractTimeObjectTypeOrder, stage.AbstractTimeObjectTypes_mapString, abstracttimeobjecttype, abstracttimeobjecttype.Name)
	return abstracttimeobjecttype
}

// StagePreserveOrder puts abstracttimeobjecttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractTimeObjectTypeOrder
// - update stage.AbstractTimeObjectTypeOrder accordingly
func (abstracttimeobjecttype *AbstractTimeObjectType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectType_stagedOrder, stage.AbstractTimeObjectType_orderStaged, &stage.AbstractTimeObjectTypeOrder, stage.AbstractTimeObjectTypes_mapString, abstracttimeobjecttype, order, abstracttimeobjecttype.Name)
}

// Unstage removes abstracttimeobjecttype off the model stage
func (abstracttimeobjecttype *AbstractTimeObjectType) Unstage(stage *Stage) *AbstractTimeObjectType {
	__gong__unstage(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectTypes_mapString, abstracttimeobjecttype, abstracttimeobjecttype.Name)
	return abstracttimeobjecttype
}

// UnstageVoid removes abstracttimeobjecttype off the model stage
func (abstracttimeobjecttype *AbstractTimeObjectType) UnstageVoid(stage *Stage) {
	abstracttimeobjecttype.Unstage(stage)
}

func (abstracttimeobjecttype *AbstractTimeObjectType) StageVoid(stage *Stage) {
	abstracttimeobjecttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstracttimeobjecttype *AbstractTimeObjectType) GetName() (res string) {
	return abstracttimeobjecttype.Name
}

// for satisfaction of GongStruct interface
func (abstracttimeobjecttype *AbstractTimeObjectType) SetName(name string) {
	abstracttimeobjecttype.Name = name
}

// Stage puts abstracttimeprimitivetype to the model stage
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) Stage(stage *Stage) *AbstractTimePrimitiveType {
	__gong__stage(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveType_stagedOrder, stage.AbstractTimePrimitiveType_orderStaged, &stage.AbstractTimePrimitiveTypeOrder, stage.AbstractTimePrimitiveTypes_mapString, abstracttimeprimitivetype, abstracttimeprimitivetype.Name)
	return abstracttimeprimitivetype
}

// StagePreserveOrder puts abstracttimeprimitivetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractTimePrimitiveTypeOrder
// - update stage.AbstractTimePrimitiveTypeOrder accordingly
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveType_stagedOrder, stage.AbstractTimePrimitiveType_orderStaged, &stage.AbstractTimePrimitiveTypeOrder, stage.AbstractTimePrimitiveTypes_mapString, abstracttimeprimitivetype, order, abstracttimeprimitivetype.Name)
}

// Unstage removes abstracttimeprimitivetype off the model stage
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) Unstage(stage *Stage) *AbstractTimePrimitiveType {
	__gong__unstage(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveTypes_mapString, abstracttimeprimitivetype, abstracttimeprimitivetype.Name)
	return abstracttimeprimitivetype
}

// UnstageVoid removes abstracttimeprimitivetype off the model stage
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) UnstageVoid(stage *Stage) {
	abstracttimeprimitivetype.Unstage(stage)
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) StageVoid(stage *Stage) {
	abstracttimeprimitivetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GetName() (res string) {
	return abstracttimeprimitivetype.Name
}

// for satisfaction of GongStruct interface
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) SetName(name string) {
	abstracttimeprimitivetype.Name = name
}

// Stage puts abstracttimeslicetype to the model stage
func (abstracttimeslicetype *AbstractTimeSliceType) Stage(stage *Stage) *AbstractTimeSliceType {
	__gong__stage(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceType_stagedOrder, stage.AbstractTimeSliceType_orderStaged, &stage.AbstractTimeSliceTypeOrder, stage.AbstractTimeSliceTypes_mapString, abstracttimeslicetype, abstracttimeslicetype.Name)
	return abstracttimeslicetype
}

// StagePreserveOrder puts abstracttimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractTimeSliceTypeOrder
// - update stage.AbstractTimeSliceTypeOrder accordingly
func (abstracttimeslicetype *AbstractTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceType_stagedOrder, stage.AbstractTimeSliceType_orderStaged, &stage.AbstractTimeSliceTypeOrder, stage.AbstractTimeSliceTypes_mapString, abstracttimeslicetype, order, abstracttimeslicetype.Name)
}

// Unstage removes abstracttimeslicetype off the model stage
func (abstracttimeslicetype *AbstractTimeSliceType) Unstage(stage *Stage) *AbstractTimeSliceType {
	__gong__unstage(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceTypes_mapString, abstracttimeslicetype, abstracttimeslicetype.Name)
	return abstracttimeslicetype
}

// UnstageVoid removes abstracttimeslicetype off the model stage
func (abstracttimeslicetype *AbstractTimeSliceType) UnstageVoid(stage *Stage) {
	abstracttimeslicetype.Unstage(stage)
}

func (abstracttimeslicetype *AbstractTimeSliceType) StageVoid(stage *Stage) {
	abstracttimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstracttimeslicetype *AbstractTimeSliceType) GetName() (res string) {
	return abstracttimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (abstracttimeslicetype *AbstractTimeSliceType) SetName(name string) {
	abstracttimeslicetype.Name = name
}

// Stage puts boundingshapetype to the model stage
func (boundingshapetype *BoundingShapeType) Stage(stage *Stage) *BoundingShapeType {
	__gong__stage(stage.BoundingShapeTypes, stage.BoundingShapeType_stagedOrder, stage.BoundingShapeType_orderStaged, &stage.BoundingShapeTypeOrder, stage.BoundingShapeTypes_mapString, boundingshapetype, boundingshapetype.Name)
	return boundingshapetype
}

// StagePreserveOrder puts boundingshapetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BoundingShapeTypeOrder
// - update stage.BoundingShapeTypeOrder accordingly
func (boundingshapetype *BoundingShapeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BoundingShapeTypes, stage.BoundingShapeType_stagedOrder, stage.BoundingShapeType_orderStaged, &stage.BoundingShapeTypeOrder, stage.BoundingShapeTypes_mapString, boundingshapetype, order, boundingshapetype.Name)
}

// Unstage removes boundingshapetype off the model stage
func (boundingshapetype *BoundingShapeType) Unstage(stage *Stage) *BoundingShapeType {
	__gong__unstage(stage.BoundingShapeTypes, stage.BoundingShapeTypes_mapString, boundingshapetype, boundingshapetype.Name)
	return boundingshapetype
}

// UnstageVoid removes boundingshapetype off the model stage
func (boundingshapetype *BoundingShapeType) UnstageVoid(stage *Stage) {
	boundingshapetype.Unstage(stage)
}

func (boundingshapetype *BoundingShapeType) StageVoid(stage *Stage) {
	boundingshapetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (boundingshapetype *BoundingShapeType) GetName() (res string) {
	return boundingshapetype.Name
}

// for satisfaction of GongStruct interface
func (boundingshapetype *BoundingShapeType) SetName(name string) {
	boundingshapetype.Name = name
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

// Stage puts directpositiontype to the model stage
func (directpositiontype *DirectPositionType) Stage(stage *Stage) *DirectPositionType {
	__gong__stage(stage.DirectPositionTypes, stage.DirectPositionType_stagedOrder, stage.DirectPositionType_orderStaged, &stage.DirectPositionTypeOrder, stage.DirectPositionTypes_mapString, directpositiontype, directpositiontype.Name)
	return directpositiontype
}

// StagePreserveOrder puts directpositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectPositionTypeOrder
// - update stage.DirectPositionTypeOrder accordingly
func (directpositiontype *DirectPositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DirectPositionTypes, stage.DirectPositionType_stagedOrder, stage.DirectPositionType_orderStaged, &stage.DirectPositionTypeOrder, stage.DirectPositionTypes_mapString, directpositiontype, order, directpositiontype.Name)
}

// Unstage removes directpositiontype off the model stage
func (directpositiontype *DirectPositionType) Unstage(stage *Stage) *DirectPositionType {
	__gong__unstage(stage.DirectPositionTypes, stage.DirectPositionTypes_mapString, directpositiontype, directpositiontype.Name)
	return directpositiontype
}

// UnstageVoid removes directpositiontype off the model stage
func (directpositiontype *DirectPositionType) UnstageVoid(stage *Stage) {
	directpositiontype.Unstage(stage)
}

func (directpositiontype *DirectPositionType) StageVoid(stage *Stage) {
	directpositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (directpositiontype *DirectPositionType) GetName() (res string) {
	return directpositiontype.Name
}

// for satisfaction of GongStruct interface
func (directpositiontype *DirectPositionType) SetName(name string) {
	directpositiontype.Name = name
}

// Stage puts dynamicfeaturetype to the model stage
func (dynamicfeaturetype *DynamicFeatureType) Stage(stage *Stage) *DynamicFeatureType {
	__gong__stage(stage.DynamicFeatureTypes, stage.DynamicFeatureType_stagedOrder, stage.DynamicFeatureType_orderStaged, &stage.DynamicFeatureTypeOrder, stage.DynamicFeatureTypes_mapString, dynamicfeaturetype, dynamicfeaturetype.Name)
	return dynamicfeaturetype
}

// StagePreserveOrder puts dynamicfeaturetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DynamicFeatureTypeOrder
// - update stage.DynamicFeatureTypeOrder accordingly
func (dynamicfeaturetype *DynamicFeatureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DynamicFeatureTypes, stage.DynamicFeatureType_stagedOrder, stage.DynamicFeatureType_orderStaged, &stage.DynamicFeatureTypeOrder, stage.DynamicFeatureTypes_mapString, dynamicfeaturetype, order, dynamicfeaturetype.Name)
}

// Unstage removes dynamicfeaturetype off the model stage
func (dynamicfeaturetype *DynamicFeatureType) Unstage(stage *Stage) *DynamicFeatureType {
	__gong__unstage(stage.DynamicFeatureTypes, stage.DynamicFeatureTypes_mapString, dynamicfeaturetype, dynamicfeaturetype.Name)
	return dynamicfeaturetype
}

// UnstageVoid removes dynamicfeaturetype off the model stage
func (dynamicfeaturetype *DynamicFeatureType) UnstageVoid(stage *Stage) {
	dynamicfeaturetype.Unstage(stage)
}

func (dynamicfeaturetype *DynamicFeatureType) StageVoid(stage *Stage) {
	dynamicfeaturetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dynamicfeaturetype *DynamicFeatureType) GetName() (res string) {
	return dynamicfeaturetype.Name
}

// for satisfaction of GongStruct interface
func (dynamicfeaturetype *DynamicFeatureType) SetName(name string) {
	dynamicfeaturetype.Name = name
}

// Stage puts envelopetype to the model stage
func (envelopetype *EnvelopeType) Stage(stage *Stage) *EnvelopeType {
	__gong__stage(stage.EnvelopeTypes, stage.EnvelopeType_stagedOrder, stage.EnvelopeType_orderStaged, &stage.EnvelopeTypeOrder, stage.EnvelopeTypes_mapString, envelopetype, envelopetype.Name)
	return envelopetype
}

// StagePreserveOrder puts envelopetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EnvelopeTypeOrder
// - update stage.EnvelopeTypeOrder accordingly
func (envelopetype *EnvelopeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EnvelopeTypes, stage.EnvelopeType_stagedOrder, stage.EnvelopeType_orderStaged, &stage.EnvelopeTypeOrder, stage.EnvelopeTypes_mapString, envelopetype, order, envelopetype.Name)
}

// Unstage removes envelopetype off the model stage
func (envelopetype *EnvelopeType) Unstage(stage *Stage) *EnvelopeType {
	__gong__unstage(stage.EnvelopeTypes, stage.EnvelopeTypes_mapString, envelopetype, envelopetype.Name)
	return envelopetype
}

// UnstageVoid removes envelopetype off the model stage
func (envelopetype *EnvelopeType) UnstageVoid(stage *Stage) {
	envelopetype.Unstage(stage)
}

func (envelopetype *EnvelopeType) StageVoid(stage *Stage) {
	envelopetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (envelopetype *EnvelopeType) GetName() (res string) {
	return envelopetype.Name
}

// for satisfaction of GongStruct interface
func (envelopetype *EnvelopeType) SetName(name string) {
	envelopetype.Name = name
}

// Stage puts pointtype to the model stage
func (pointtype *PointType) Stage(stage *Stage) *PointType {
	__gong__stage(stage.PointTypes, stage.PointType_stagedOrder, stage.PointType_orderStaged, &stage.PointTypeOrder, stage.PointTypes_mapString, pointtype, pointtype.Name)
	return pointtype
}

// StagePreserveOrder puts pointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PointTypeOrder
// - update stage.PointTypeOrder accordingly
func (pointtype *PointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PointTypes, stage.PointType_stagedOrder, stage.PointType_orderStaged, &stage.PointTypeOrder, stage.PointTypes_mapString, pointtype, order, pointtype.Name)
}

// Unstage removes pointtype off the model stage
func (pointtype *PointType) Unstage(stage *Stage) *PointType {
	__gong__unstage(stage.PointTypes, stage.PointTypes_mapString, pointtype, pointtype.Name)
	return pointtype
}

// UnstageVoid removes pointtype off the model stage
func (pointtype *PointType) UnstageVoid(stage *Stage) {
	pointtype.Unstage(stage)
}

func (pointtype *PointType) StageVoid(stage *Stage) {
	pointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pointtype *PointType) GetName() (res string) {
	return pointtype.Name
}

// for satisfaction of GongStruct interface
func (pointtype *PointType) SetName(name string) {
	pointtype.Name = name
}

// Stage puts timeinstanttype to the model stage
func (timeinstanttype *TimeInstantType) Stage(stage *Stage) *TimeInstantType {
	__gong__stage(stage.TimeInstantTypes, stage.TimeInstantType_stagedOrder, stage.TimeInstantType_orderStaged, &stage.TimeInstantTypeOrder, stage.TimeInstantTypes_mapString, timeinstanttype, timeinstanttype.Name)
	return timeinstanttype
}

// StagePreserveOrder puts timeinstanttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimeInstantTypeOrder
// - update stage.TimeInstantTypeOrder accordingly
func (timeinstanttype *TimeInstantType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimeInstantTypes, stage.TimeInstantType_stagedOrder, stage.TimeInstantType_orderStaged, &stage.TimeInstantTypeOrder, stage.TimeInstantTypes_mapString, timeinstanttype, order, timeinstanttype.Name)
}

// Unstage removes timeinstanttype off the model stage
func (timeinstanttype *TimeInstantType) Unstage(stage *Stage) *TimeInstantType {
	__gong__unstage(stage.TimeInstantTypes, stage.TimeInstantTypes_mapString, timeinstanttype, timeinstanttype.Name)
	return timeinstanttype
}

// UnstageVoid removes timeinstanttype off the model stage
func (timeinstanttype *TimeInstantType) UnstageVoid(stage *Stage) {
	timeinstanttype.Unstage(stage)
}

func (timeinstanttype *TimeInstantType) StageVoid(stage *Stage) {
	timeinstanttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timeinstanttype *TimeInstantType) GetName() (res string) {
	return timeinstanttype.Name
}

// for satisfaction of GongStruct interface
func (timeinstanttype *TimeInstantType) SetName(name string) {
	timeinstanttype.Name = name
}

// Stage puts timeperiodtype to the model stage
func (timeperiodtype *TimePeriodType) Stage(stage *Stage) *TimePeriodType {
	__gong__stage(stage.TimePeriodTypes, stage.TimePeriodType_stagedOrder, stage.TimePeriodType_orderStaged, &stage.TimePeriodTypeOrder, stage.TimePeriodTypes_mapString, timeperiodtype, timeperiodtype.Name)
	return timeperiodtype
}

// StagePreserveOrder puts timeperiodtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimePeriodTypeOrder
// - update stage.TimePeriodTypeOrder accordingly
func (timeperiodtype *TimePeriodType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimePeriodTypes, stage.TimePeriodType_stagedOrder, stage.TimePeriodType_orderStaged, &stage.TimePeriodTypeOrder, stage.TimePeriodTypes_mapString, timeperiodtype, order, timeperiodtype.Name)
}

// Unstage removes timeperiodtype off the model stage
func (timeperiodtype *TimePeriodType) Unstage(stage *Stage) *TimePeriodType {
	__gong__unstage(stage.TimePeriodTypes, stage.TimePeriodTypes_mapString, timeperiodtype, timeperiodtype.Name)
	return timeperiodtype
}

// UnstageVoid removes timeperiodtype off the model stage
func (timeperiodtype *TimePeriodType) UnstageVoid(stage *Stage) {
	timeperiodtype.Unstage(stage)
}

func (timeperiodtype *TimePeriodType) StageVoid(stage *Stage) {
	timeperiodtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timeperiodtype *TimePeriodType) GetName() (res string) {
	return timeperiodtype.Name
}

// for satisfaction of GongStruct interface
func (timeperiodtype *TimePeriodType) SetName(name string) {
	timeperiodtype.Name = name
}

// Stage puts timepositiontype to the model stage
func (timepositiontype *TimePositionType) Stage(stage *Stage) *TimePositionType {
	__gong__stage(stage.TimePositionTypes, stage.TimePositionType_stagedOrder, stage.TimePositionType_orderStaged, &stage.TimePositionTypeOrder, stage.TimePositionTypes_mapString, timepositiontype, timepositiontype.Name)
	return timepositiontype
}

// StagePreserveOrder puts timepositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimePositionTypeOrder
// - update stage.TimePositionTypeOrder accordingly
func (timepositiontype *TimePositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimePositionTypes, stage.TimePositionType_stagedOrder, stage.TimePositionType_orderStaged, &stage.TimePositionTypeOrder, stage.TimePositionTypes_mapString, timepositiontype, order, timepositiontype.Name)
}

// Unstage removes timepositiontype off the model stage
func (timepositiontype *TimePositionType) Unstage(stage *Stage) *TimePositionType {
	__gong__unstage(stage.TimePositionTypes, stage.TimePositionTypes_mapString, timepositiontype, timepositiontype.Name)
	return timepositiontype
}

// UnstageVoid removes timepositiontype off the model stage
func (timepositiontype *TimePositionType) UnstageVoid(stage *Stage) {
	timepositiontype.Unstage(stage)
}

func (timepositiontype *TimePositionType) StageVoid(stage *Stage) {
	timepositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timepositiontype *TimePositionType) GetName() (res string) {
	return timepositiontype.Name
}

// for satisfaction of GongStruct interface
func (timepositiontype *TimePositionType) SetName(name string) {
	timepositiontype.Name = name
}

// Stage puts timeprimitivepropertytype to the model stage
func (timeprimitivepropertytype *TimePrimitivePropertyType) Stage(stage *Stage) *TimePrimitivePropertyType {
	__gong__stage(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyType_stagedOrder, stage.TimePrimitivePropertyType_orderStaged, &stage.TimePrimitivePropertyTypeOrder, stage.TimePrimitivePropertyTypes_mapString, timeprimitivepropertytype, timeprimitivepropertytype.Name)
	return timeprimitivepropertytype
}

// StagePreserveOrder puts timeprimitivepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimePrimitivePropertyTypeOrder
// - update stage.TimePrimitivePropertyTypeOrder accordingly
func (timeprimitivepropertytype *TimePrimitivePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyType_stagedOrder, stage.TimePrimitivePropertyType_orderStaged, &stage.TimePrimitivePropertyTypeOrder, stage.TimePrimitivePropertyTypes_mapString, timeprimitivepropertytype, order, timeprimitivepropertytype.Name)
}

// Unstage removes timeprimitivepropertytype off the model stage
func (timeprimitivepropertytype *TimePrimitivePropertyType) Unstage(stage *Stage) *TimePrimitivePropertyType {
	__gong__unstage(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyTypes_mapString, timeprimitivepropertytype, timeprimitivepropertytype.Name)
	return timeprimitivepropertytype
}

// UnstageVoid removes timeprimitivepropertytype off the model stage
func (timeprimitivepropertytype *TimePrimitivePropertyType) UnstageVoid(stage *Stage) {
	timeprimitivepropertytype.Unstage(stage)
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) StageVoid(stage *Stage) {
	timeprimitivepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timeprimitivepropertytype *TimePrimitivePropertyType) GetName() (res string) {
	return timeprimitivepropertytype.Name
}

// for satisfaction of GongStruct interface
func (timeprimitivepropertytype *TimePrimitivePropertyType) SetName(name string) {
	timeprimitivepropertytype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractFeatureMemberTypes, &stage.AbstractFeatureMemberTypes_mapString, &stage.AbstractFeatureMemberType_stagedOrder, &stage.AbstractFeatureMemberTypeOrder)

	__gong__resetStageType(&stage.AbstractFeatureTypes, &stage.AbstractFeatureTypes_mapString, &stage.AbstractFeatureType_stagedOrder, &stage.AbstractFeatureTypeOrder)

	__gong__resetStageType(&stage.AbstractGMLTypes, &stage.AbstractGMLTypes_mapString, &stage.AbstractGMLType_stagedOrder, &stage.AbstractGMLTypeOrder)

	__gong__resetStageType(&stage.AbstractGeometricPrimitiveTypes, &stage.AbstractGeometricPrimitiveTypes_mapString, &stage.AbstractGeometricPrimitiveType_stagedOrder, &stage.AbstractGeometricPrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractGeometryTypes, &stage.AbstractGeometryTypes_mapString, &stage.AbstractGeometryType_stagedOrder, &stage.AbstractGeometryTypeOrder)

	__gong__resetStageType(&stage.AbstractMetadataPropertyTypes, &stage.AbstractMetadataPropertyTypes_mapString, &stage.AbstractMetadataPropertyType_stagedOrder, &stage.AbstractMetadataPropertyTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeGeometricPrimitiveTypes, &stage.AbstractTimeGeometricPrimitiveTypes_mapString, &stage.AbstractTimeGeometricPrimitiveType_stagedOrder, &stage.AbstractTimeGeometricPrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeObjectTypes, &stage.AbstractTimeObjectTypes_mapString, &stage.AbstractTimeObjectType_stagedOrder, &stage.AbstractTimeObjectTypeOrder)

	__gong__resetStageType(&stage.AbstractTimePrimitiveTypes, &stage.AbstractTimePrimitiveTypes_mapString, &stage.AbstractTimePrimitiveType_stagedOrder, &stage.AbstractTimePrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeSliceTypes, &stage.AbstractTimeSliceTypes_mapString, &stage.AbstractTimeSliceType_stagedOrder, &stage.AbstractTimeSliceTypeOrder)

	__gong__resetStageType(&stage.BoundingShapeTypes, &stage.BoundingShapeTypes_mapString, &stage.BoundingShapeType_stagedOrder, &stage.BoundingShapeTypeOrder)

	__gong__resetStageType(&stage.CodeTypes, &stage.CodeTypes_mapString, &stage.CodeType_stagedOrder, &stage.CodeTypeOrder)

	__gong__resetStageType(&stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString, &stage.CodeWithAuthorityType_stagedOrder, &stage.CodeWithAuthorityTypeOrder)

	__gong__resetStageType(&stage.DirectPositionTypes, &stage.DirectPositionTypes_mapString, &stage.DirectPositionType_stagedOrder, &stage.DirectPositionTypeOrder)

	__gong__resetStageType(&stage.DynamicFeatureTypes, &stage.DynamicFeatureTypes_mapString, &stage.DynamicFeatureType_stagedOrder, &stage.DynamicFeatureTypeOrder)

	__gong__resetStageType(&stage.EnvelopeTypes, &stage.EnvelopeTypes_mapString, &stage.EnvelopeType_stagedOrder, &stage.EnvelopeTypeOrder)

	__gong__resetStageType(&stage.PointTypes, &stage.PointTypes_mapString, &stage.PointType_stagedOrder, &stage.PointTypeOrder)

	__gong__resetStageType(&stage.TimeInstantTypes, &stage.TimeInstantTypes_mapString, &stage.TimeInstantType_stagedOrder, &stage.TimeInstantTypeOrder)

	__gong__resetStageType(&stage.TimePeriodTypes, &stage.TimePeriodTypes_mapString, &stage.TimePeriodType_stagedOrder, &stage.TimePeriodTypeOrder)

	__gong__resetStageType(&stage.TimePositionTypes, &stage.TimePositionTypes_mapString, &stage.TimePositionType_stagedOrder, &stage.TimePositionTypeOrder)

	__gong__resetStageType(&stage.TimePrimitivePropertyTypes, &stage.TimePrimitivePropertyTypes_mapString, &stage.TimePrimitivePropertyType_stagedOrder, &stage.TimePrimitivePropertyTypeOrder)

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
func (AbstractFeatureMemberType) GongGetAssociationName() any {
	return &AbstractFeatureMemberType{
	}
}

func (AbstractFeatureType) GongGetAssociationName() any {
	return &AbstractFeatureType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			BoundedBy: &BoundingShapeType{Name: "BoundedBy"},
	}
}

func (AbstractGMLType) GongGetAssociationName() any {
	return &AbstractGMLType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractGeometricPrimitiveType) GongGetAssociationName() any {
	return &AbstractGeometricPrimitiveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractGeometryType) GongGetAssociationName() any {
	return &AbstractGeometryType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractMetadataPropertyType) GongGetAssociationName() any {
	return &AbstractMetadataPropertyType{
	}
}

func (AbstractTimeGeometricPrimitiveType) GongGetAssociationName() any {
	return &AbstractTimeGeometricPrimitiveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractTimeObjectType) GongGetAssociationName() any {
	return &AbstractTimeObjectType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractTimePrimitiveType) GongGetAssociationName() any {
	return &AbstractTimePrimitiveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

func (AbstractTimeSliceType) GongGetAssociationName() any {
	return &AbstractTimeSliceType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			ValidTime: &TimePrimitivePropertyType{Name: "ValidTime"},
	}
}

func (BoundingShapeType) GongGetAssociationName() any {
	return &BoundingShapeType{
			Envelope: &EnvelopeType{Name: "Envelope"},
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

func (DirectPositionType) GongGetAssociationName() any {
	return &DirectPositionType{
	}
}

func (DynamicFeatureType) GongGetAssociationName() any {
	return &DynamicFeatureType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			BoundedBy: &BoundingShapeType{Name: "BoundedBy"},
	}
}

func (EnvelopeType) GongGetAssociationName() any {
	return &EnvelopeType{
			LowerCorner: &DirectPositionType{Name: "LowerCorner"},
			UpperCorner: &DirectPositionType{Name: "UpperCorner"},
	}
}

func (PointType) GongGetAssociationName() any {
	return &PointType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			Pos: &DirectPositionType{Name: "Pos"},
	}
}

func (TimeInstantType) GongGetAssociationName() any {
	return &TimeInstantType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			TimePosition: &TimePositionType{Name: "TimePosition"},
	}
}

func (TimePeriodType) GongGetAssociationName() any {
	return &TimePeriodType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			BeginPosition: &TimePositionType{Name: "BeginPosition"},
			EndPosition: &TimePositionType{Name: "EndPosition"},
	}
}

func (TimePositionType) GongGetAssociationName() any {
	return &TimePositionType{
	}
}

func (TimePrimitivePropertyType) GongGetAssociationName() any {
	return &TimePrimitivePropertyType{
			TimeInstant: &TimeInstantType{Name: "TimeInstant"},
			TimePeriod: &TimePeriodType{Name: "TimePeriod"},
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
	// reverse maps of direct associations of AbstractFeatureMemberType
	case AbstractFeatureMemberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractFeatureType
	case AbstractFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractFeatureType)
			for abstractfeaturetype := range stage.AbstractFeatureTypes {
				if abstractfeaturetype.Identifier != nil {
					codewithauthoritytype_ := abstractfeaturetype.Identifier
					var abstractfeaturetypes []*AbstractFeatureType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractfeaturetypes = res[codewithauthoritytype_]
					} else {
						abstractfeaturetypes = make([]*AbstractFeatureType, 0)
					}
					abstractfeaturetypes = append(abstractfeaturetypes, abstractfeaturetype)
					res[codewithauthoritytype_] = abstractfeaturetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "BoundedBy":
			res := make(map[*BoundingShapeType][]*AbstractFeatureType)
			for abstractfeaturetype := range stage.AbstractFeatureTypes {
				if abstractfeaturetype.BoundedBy != nil {
					boundingshapetype_ := abstractfeaturetype.BoundedBy
					var abstractfeaturetypes []*AbstractFeatureType
					_, ok := res[boundingshapetype_]
					if ok {
						abstractfeaturetypes = res[boundingshapetype_]
					} else {
						abstractfeaturetypes = make([]*AbstractFeatureType, 0)
					}
					abstractfeaturetypes = append(abstractfeaturetypes, abstractfeaturetype)
					res[boundingshapetype_] = abstractfeaturetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGMLType
	case AbstractGMLType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractGMLType)
			for abstractgmltype := range stage.AbstractGMLTypes {
				if abstractgmltype.Identifier != nil {
					codewithauthoritytype_ := abstractgmltype.Identifier
					var abstractgmltypes []*AbstractGMLType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractgmltypes = res[codewithauthoritytype_]
					} else {
						abstractgmltypes = make([]*AbstractGMLType, 0)
					}
					abstractgmltypes = append(abstractgmltypes, abstractgmltype)
					res[codewithauthoritytype_] = abstractgmltypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGeometricPrimitiveType
	case AbstractGeometricPrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractGeometricPrimitiveType)
			for abstractgeometricprimitivetype := range stage.AbstractGeometricPrimitiveTypes {
				if abstractgeometricprimitivetype.Identifier != nil {
					codewithauthoritytype_ := abstractgeometricprimitivetype.Identifier
					var abstractgeometricprimitivetypes []*AbstractGeometricPrimitiveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractgeometricprimitivetypes = res[codewithauthoritytype_]
					} else {
						abstractgeometricprimitivetypes = make([]*AbstractGeometricPrimitiveType, 0)
					}
					abstractgeometricprimitivetypes = append(abstractgeometricprimitivetypes, abstractgeometricprimitivetype)
					res[codewithauthoritytype_] = abstractgeometricprimitivetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGeometryType
	case AbstractGeometryType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractGeometryType)
			for abstractgeometrytype := range stage.AbstractGeometryTypes {
				if abstractgeometrytype.Identifier != nil {
					codewithauthoritytype_ := abstractgeometrytype.Identifier
					var abstractgeometrytypes []*AbstractGeometryType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractgeometrytypes = res[codewithauthoritytype_]
					} else {
						abstractgeometrytypes = make([]*AbstractGeometryType, 0)
					}
					abstractgeometrytypes = append(abstractgeometrytypes, abstractgeometrytype)
					res[codewithauthoritytype_] = abstractgeometrytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractMetadataPropertyType
	case AbstractMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractTimeGeometricPrimitiveType
	case AbstractTimeGeometricPrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractTimeGeometricPrimitiveType)
			for abstracttimegeometricprimitivetype := range stage.AbstractTimeGeometricPrimitiveTypes {
				if abstracttimegeometricprimitivetype.Identifier != nil {
					codewithauthoritytype_ := abstracttimegeometricprimitivetype.Identifier
					var abstracttimegeometricprimitivetypes []*AbstractTimeGeometricPrimitiveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstracttimegeometricprimitivetypes = res[codewithauthoritytype_]
					} else {
						abstracttimegeometricprimitivetypes = make([]*AbstractTimeGeometricPrimitiveType, 0)
					}
					abstracttimegeometricprimitivetypes = append(abstracttimegeometricprimitivetypes, abstracttimegeometricprimitivetype)
					res[codewithauthoritytype_] = abstracttimegeometricprimitivetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimeObjectType
	case AbstractTimeObjectType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractTimeObjectType)
			for abstracttimeobjecttype := range stage.AbstractTimeObjectTypes {
				if abstracttimeobjecttype.Identifier != nil {
					codewithauthoritytype_ := abstracttimeobjecttype.Identifier
					var abstracttimeobjecttypes []*AbstractTimeObjectType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstracttimeobjecttypes = res[codewithauthoritytype_]
					} else {
						abstracttimeobjecttypes = make([]*AbstractTimeObjectType, 0)
					}
					abstracttimeobjecttypes = append(abstracttimeobjecttypes, abstracttimeobjecttype)
					res[codewithauthoritytype_] = abstracttimeobjecttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimePrimitiveType
	case AbstractTimePrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractTimePrimitiveType)
			for abstracttimeprimitivetype := range stage.AbstractTimePrimitiveTypes {
				if abstracttimeprimitivetype.Identifier != nil {
					codewithauthoritytype_ := abstracttimeprimitivetype.Identifier
					var abstracttimeprimitivetypes []*AbstractTimePrimitiveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstracttimeprimitivetypes = res[codewithauthoritytype_]
					} else {
						abstracttimeprimitivetypes = make([]*AbstractTimePrimitiveType, 0)
					}
					abstracttimeprimitivetypes = append(abstracttimeprimitivetypes, abstracttimeprimitivetype)
					res[codewithauthoritytype_] = abstracttimeprimitivetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimeSliceType
	case AbstractTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractTimeSliceType)
			for abstracttimeslicetype := range stage.AbstractTimeSliceTypes {
				if abstracttimeslicetype.Identifier != nil {
					codewithauthoritytype_ := abstracttimeslicetype.Identifier
					var abstracttimeslicetypes []*AbstractTimeSliceType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstracttimeslicetypes = res[codewithauthoritytype_]
					} else {
						abstracttimeslicetypes = make([]*AbstractTimeSliceType, 0)
					}
					abstracttimeslicetypes = append(abstracttimeslicetypes, abstracttimeslicetype)
					res[codewithauthoritytype_] = abstracttimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ValidTime":
			res := make(map[*TimePrimitivePropertyType][]*AbstractTimeSliceType)
			for abstracttimeslicetype := range stage.AbstractTimeSliceTypes {
				if abstracttimeslicetype.ValidTime != nil {
					timeprimitivepropertytype_ := abstracttimeslicetype.ValidTime
					var abstracttimeslicetypes []*AbstractTimeSliceType
					_, ok := res[timeprimitivepropertytype_]
					if ok {
						abstracttimeslicetypes = res[timeprimitivepropertytype_]
					} else {
						abstracttimeslicetypes = make([]*AbstractTimeSliceType, 0)
					}
					abstracttimeslicetypes = append(abstracttimeslicetypes, abstracttimeslicetype)
					res[timeprimitivepropertytype_] = abstracttimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BoundingShapeType
	case BoundingShapeType:
		switch fieldname {
		// insertion point for per direct association field
		case "Envelope":
			res := make(map[*EnvelopeType][]*BoundingShapeType)
			for boundingshapetype := range stage.BoundingShapeTypes {
				if boundingshapetype.Envelope != nil {
					envelopetype_ := boundingshapetype.Envelope
					var boundingshapetypes []*BoundingShapeType
					_, ok := res[envelopetype_]
					if ok {
						boundingshapetypes = res[envelopetype_]
					} else {
						boundingshapetypes = make([]*BoundingShapeType, 0)
					}
					boundingshapetypes = append(boundingshapetypes, boundingshapetype)
					res[envelopetype_] = boundingshapetypes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of DirectPositionType
	case DirectPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DynamicFeatureType
	case DynamicFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*DynamicFeatureType)
			for dynamicfeaturetype := range stage.DynamicFeatureTypes {
				if dynamicfeaturetype.Identifier != nil {
					codewithauthoritytype_ := dynamicfeaturetype.Identifier
					var dynamicfeaturetypes []*DynamicFeatureType
					_, ok := res[codewithauthoritytype_]
					if ok {
						dynamicfeaturetypes = res[codewithauthoritytype_]
					} else {
						dynamicfeaturetypes = make([]*DynamicFeatureType, 0)
					}
					dynamicfeaturetypes = append(dynamicfeaturetypes, dynamicfeaturetype)
					res[codewithauthoritytype_] = dynamicfeaturetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "BoundedBy":
			res := make(map[*BoundingShapeType][]*DynamicFeatureType)
			for dynamicfeaturetype := range stage.DynamicFeatureTypes {
				if dynamicfeaturetype.BoundedBy != nil {
					boundingshapetype_ := dynamicfeaturetype.BoundedBy
					var dynamicfeaturetypes []*DynamicFeatureType
					_, ok := res[boundingshapetype_]
					if ok {
						dynamicfeaturetypes = res[boundingshapetype_]
					} else {
						dynamicfeaturetypes = make([]*DynamicFeatureType, 0)
					}
					dynamicfeaturetypes = append(dynamicfeaturetypes, dynamicfeaturetype)
					res[boundingshapetype_] = dynamicfeaturetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of EnvelopeType
	case EnvelopeType:
		switch fieldname {
		// insertion point for per direct association field
		case "LowerCorner":
			res := make(map[*DirectPositionType][]*EnvelopeType)
			for envelopetype := range stage.EnvelopeTypes {
				if envelopetype.LowerCorner != nil {
					directpositiontype_ := envelopetype.LowerCorner
					var envelopetypes []*EnvelopeType
					_, ok := res[directpositiontype_]
					if ok {
						envelopetypes = res[directpositiontype_]
					} else {
						envelopetypes = make([]*EnvelopeType, 0)
					}
					envelopetypes = append(envelopetypes, envelopetype)
					res[directpositiontype_] = envelopetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "UpperCorner":
			res := make(map[*DirectPositionType][]*EnvelopeType)
			for envelopetype := range stage.EnvelopeTypes {
				if envelopetype.UpperCorner != nil {
					directpositiontype_ := envelopetype.UpperCorner
					var envelopetypes []*EnvelopeType
					_, ok := res[directpositiontype_]
					if ok {
						envelopetypes = res[directpositiontype_]
					} else {
						envelopetypes = make([]*EnvelopeType, 0)
					}
					envelopetypes = append(envelopetypes, envelopetype)
					res[directpositiontype_] = envelopetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PointType
	case PointType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*PointType)
			for pointtype := range stage.PointTypes {
				if pointtype.Identifier != nil {
					codewithauthoritytype_ := pointtype.Identifier
					var pointtypes []*PointType
					_, ok := res[codewithauthoritytype_]
					if ok {
						pointtypes = res[codewithauthoritytype_]
					} else {
						pointtypes = make([]*PointType, 0)
					}
					pointtypes = append(pointtypes, pointtype)
					res[codewithauthoritytype_] = pointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pos":
			res := make(map[*DirectPositionType][]*PointType)
			for pointtype := range stage.PointTypes {
				if pointtype.Pos != nil {
					directpositiontype_ := pointtype.Pos
					var pointtypes []*PointType
					_, ok := res[directpositiontype_]
					if ok {
						pointtypes = res[directpositiontype_]
					} else {
						pointtypes = make([]*PointType, 0)
					}
					pointtypes = append(pointtypes, pointtype)
					res[directpositiontype_] = pointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimeInstantType
	case TimeInstantType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*TimeInstantType)
			for timeinstanttype := range stage.TimeInstantTypes {
				if timeinstanttype.Identifier != nil {
					codewithauthoritytype_ := timeinstanttype.Identifier
					var timeinstanttypes []*TimeInstantType
					_, ok := res[codewithauthoritytype_]
					if ok {
						timeinstanttypes = res[codewithauthoritytype_]
					} else {
						timeinstanttypes = make([]*TimeInstantType, 0)
					}
					timeinstanttypes = append(timeinstanttypes, timeinstanttype)
					res[codewithauthoritytype_] = timeinstanttypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "TimePosition":
			res := make(map[*TimePositionType][]*TimeInstantType)
			for timeinstanttype := range stage.TimeInstantTypes {
				if timeinstanttype.TimePosition != nil {
					timepositiontype_ := timeinstanttype.TimePosition
					var timeinstanttypes []*TimeInstantType
					_, ok := res[timepositiontype_]
					if ok {
						timeinstanttypes = res[timepositiontype_]
					} else {
						timeinstanttypes = make([]*TimeInstantType, 0)
					}
					timeinstanttypes = append(timeinstanttypes, timeinstanttype)
					res[timepositiontype_] = timeinstanttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimePeriodType
	case TimePeriodType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*TimePeriodType)
			for timeperiodtype := range stage.TimePeriodTypes {
				if timeperiodtype.Identifier != nil {
					codewithauthoritytype_ := timeperiodtype.Identifier
					var timeperiodtypes []*TimePeriodType
					_, ok := res[codewithauthoritytype_]
					if ok {
						timeperiodtypes = res[codewithauthoritytype_]
					} else {
						timeperiodtypes = make([]*TimePeriodType, 0)
					}
					timeperiodtypes = append(timeperiodtypes, timeperiodtype)
					res[codewithauthoritytype_] = timeperiodtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "BeginPosition":
			res := make(map[*TimePositionType][]*TimePeriodType)
			for timeperiodtype := range stage.TimePeriodTypes {
				if timeperiodtype.BeginPosition != nil {
					timepositiontype_ := timeperiodtype.BeginPosition
					var timeperiodtypes []*TimePeriodType
					_, ok := res[timepositiontype_]
					if ok {
						timeperiodtypes = res[timepositiontype_]
					} else {
						timeperiodtypes = make([]*TimePeriodType, 0)
					}
					timeperiodtypes = append(timeperiodtypes, timeperiodtype)
					res[timepositiontype_] = timeperiodtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndPosition":
			res := make(map[*TimePositionType][]*TimePeriodType)
			for timeperiodtype := range stage.TimePeriodTypes {
				if timeperiodtype.EndPosition != nil {
					timepositiontype_ := timeperiodtype.EndPosition
					var timeperiodtypes []*TimePeriodType
					_, ok := res[timepositiontype_]
					if ok {
						timeperiodtypes = res[timepositiontype_]
					} else {
						timeperiodtypes = make([]*TimePeriodType, 0)
					}
					timeperiodtypes = append(timeperiodtypes, timeperiodtype)
					res[timepositiontype_] = timeperiodtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimePositionType
	case TimePositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimePrimitivePropertyType
	case TimePrimitivePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInstant":
			res := make(map[*TimeInstantType][]*TimePrimitivePropertyType)
			for timeprimitivepropertytype := range stage.TimePrimitivePropertyTypes {
				if timeprimitivepropertytype.TimeInstant != nil {
					timeinstanttype_ := timeprimitivepropertytype.TimeInstant
					var timeprimitivepropertytypes []*TimePrimitivePropertyType
					_, ok := res[timeinstanttype_]
					if ok {
						timeprimitivepropertytypes = res[timeinstanttype_]
					} else {
						timeprimitivepropertytypes = make([]*TimePrimitivePropertyType, 0)
					}
					timeprimitivepropertytypes = append(timeprimitivepropertytypes, timeprimitivepropertytype)
					res[timeinstanttype_] = timeprimitivepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "TimePeriod":
			res := make(map[*TimePeriodType][]*TimePrimitivePropertyType)
			for timeprimitivepropertytype := range stage.TimePrimitivePropertyTypes {
				if timeprimitivepropertytype.TimePeriod != nil {
					timeperiodtype_ := timeprimitivepropertytype.TimePeriod
					var timeprimitivepropertytypes []*TimePrimitivePropertyType
					_, ok := res[timeperiodtype_]
					if ok {
						timeprimitivepropertytypes = res[timeperiodtype_]
					} else {
						timeprimitivepropertytypes = make([]*TimePrimitivePropertyType, 0)
					}
					timeprimitivepropertytypes = append(timeprimitivepropertytypes, timeprimitivepropertytype)
					res[timeperiodtype_] = timeprimitivepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AbstractFeatureMemberType
	case AbstractFeatureMemberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractFeatureType
	case AbstractFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractFeatureType)
			for abstractfeaturetype := range stage.AbstractFeatureTypes {
				for _, codetype_ := range abstractfeaturetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractfeaturetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGMLType
	case AbstractGMLType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractGMLType)
			for abstractgmltype := range stage.AbstractGMLTypes {
				for _, codetype_ := range abstractgmltype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractgmltype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGeometricPrimitiveType
	case AbstractGeometricPrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractGeometricPrimitiveType)
			for abstractgeometricprimitivetype := range stage.AbstractGeometricPrimitiveTypes {
				for _, codetype_ := range abstractgeometricprimitivetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractgeometricprimitivetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractGeometryType
	case AbstractGeometryType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractGeometryType)
			for abstractgeometrytype := range stage.AbstractGeometryTypes {
				for _, codetype_ := range abstractgeometrytype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractgeometrytype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractMetadataPropertyType
	case AbstractMetadataPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractTimeGeometricPrimitiveType
	case AbstractTimeGeometricPrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractTimeGeometricPrimitiveType)
			for abstracttimegeometricprimitivetype := range stage.AbstractTimeGeometricPrimitiveTypes {
				for _, codetype_ := range abstracttimegeometricprimitivetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstracttimegeometricprimitivetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimeObjectType
	case AbstractTimeObjectType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractTimeObjectType)
			for abstracttimeobjecttype := range stage.AbstractTimeObjectTypes {
				for _, codetype_ := range abstracttimeobjecttype.NameXSD {
					res[codetype_] = append(res[codetype_], abstracttimeobjecttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimePrimitiveType
	case AbstractTimePrimitiveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractTimePrimitiveType)
			for abstracttimeprimitivetype := range stage.AbstractTimePrimitiveTypes {
				for _, codetype_ := range abstracttimeprimitivetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstracttimeprimitivetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractTimeSliceType
	case AbstractTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractTimeSliceType)
			for abstracttimeslicetype := range stage.AbstractTimeSliceTypes {
				for _, codetype_ := range abstracttimeslicetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstracttimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BoundingShapeType
	case BoundingShapeType:
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
	// reverse maps of direct associations of DirectPositionType
	case DirectPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DynamicFeatureType
	case DynamicFeatureType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*DynamicFeatureType)
			for dynamicfeaturetype := range stage.DynamicFeatureTypes {
				for _, codetype_ := range dynamicfeaturetype.NameXSD {
					res[codetype_] = append(res[codetype_], dynamicfeaturetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of EnvelopeType
	case EnvelopeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointType
	case PointType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*PointType)
			for pointtype := range stage.PointTypes {
				for _, codetype_ := range pointtype.NameXSD {
					res[codetype_] = append(res[codetype_], pointtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimeInstantType
	case TimeInstantType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*TimeInstantType)
			for timeinstanttype := range stage.TimeInstantTypes {
				for _, codetype_ := range timeinstanttype.NameXSD {
					res[codetype_] = append(res[codetype_], timeinstanttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimePeriodType
	case TimePeriodType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*TimePeriodType)
			for timeperiodtype := range stage.TimePeriodTypes {
				for _, codetype_ := range timeperiodtype.NameXSD {
					res[codetype_] = append(res[codetype_], timeperiodtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimePositionType
	case TimePositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimePrimitivePropertyType
	case TimePrimitivePropertyType:
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
func (*AbstractFeatureMemberType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractFeatureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractGMLType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractGeometricPrimitiveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractGeometryType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractMetadataPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractTimeGeometricPrimitiveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractTimeObjectType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractTimePrimitiveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*BoundingShapeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AbstractFeatureType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractGMLType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractGeometricPrimitiveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractGeometryType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractTimeGeometricPrimitiveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractTimeObjectType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractTimePrimitiveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "AbstractTimeSliceType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "DynamicFeatureType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "PointType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "TimeInstantType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "TimePeriodType",
			Fieldname: "NameXSD",
		},
	}
}

func (*CodeWithAuthorityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DirectPositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DynamicFeatureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EnvelopeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimeInstantType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimePeriodType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimePositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimePrimitivePropertyType) GongGetReverseFields() []GongReverseField {
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
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractfeaturetype *AbstractFeatureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:                 "BoundedBy",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BoundingShapeType",
		},
	}
	return
}

func (abstractgmltype *AbstractGMLType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
	}
	return
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractgeometrytype *AbstractGeometryType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "Frame",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
	}
	return
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
	}
	return
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:                 "ValidTime",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimePrimitivePropertyType",
		},
	}
	return
}

func (boundingshapetype *BoundingShapeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Envelope",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EnvelopeType",
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

func (directpositiontype *DirectPositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:                 "BoundedBy",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BoundingShapeType",
		},
	}
	return
}

func (envelopetype *EnvelopeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "LowerCorner",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "UpperCorner",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
	}
	return
}

func (pointtype *PointType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
	}
	return
}

func (timeinstanttype *TimeInstantType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "Frame",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TimePosition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimePositionType",
		},
	}
	return
}

func (timeperiodtype *TimePeriodType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Identifier",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CodeWithAuthorityType",
		},
		{
			Name:                 "NameXSD",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CodeType",
		},
		{
			Name:               "Frame",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BeginPosition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimePositionType",
		},
		{
			Name:                 "EndPosition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimePositionType",
		},
	}
	return
}

func (timepositiontype *TimePositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Frame",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "IndeterminatePosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_TimeIndeterminateValueType",
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeInstant",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimeInstantType",
		},
		{
			Name:                 "TimePeriod",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimePeriodType",
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
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractfeaturemembertype.Name
	}
	return
}

func (abstractfeaturetype *AbstractFeatureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractfeaturetype.Name
	case "Id":
		res.valueString = abstractfeaturetype.Id
	case "Description":
		res.valueString = abstractfeaturetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractfeaturetype.Identifier != nil {
			res.valueString = abstractfeaturetype.Identifier.Name
			res.ids = abstractfeaturetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractfeaturetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "BoundedBy":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractfeaturetype.BoundedBy != nil {
			res.valueString = abstractfeaturetype.BoundedBy.Name
			res.ids = abstractfeaturetype.BoundedBy.GongGetUUID(stage)
		}
	}
	return
}

func (abstractgmltype *AbstractGMLType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractgmltype.Name
	case "Id":
		res.valueString = abstractgmltype.Id
	case "Description":
		res.valueString = abstractgmltype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractgmltype.Identifier != nil {
			res.valueString = abstractgmltype.Identifier.Name
			res.ids = abstractgmltype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractgmltype.NameXSD {
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

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractgeometricprimitivetype.Name
	case "Id":
		res.valueString = abstractgeometricprimitivetype.Id
	case "Description":
		res.valueString = abstractgeometricprimitivetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractgeometricprimitivetype.Identifier != nil {
			res.valueString = abstractgeometricprimitivetype.Identifier.Name
			res.ids = abstractgeometricprimitivetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractgeometricprimitivetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = abstractgeometricprimitivetype.SrsName
	}
	return
}

func (abstractgeometrytype *AbstractGeometryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractgeometrytype.Name
	case "Id":
		res.valueString = abstractgeometrytype.Id
	case "Description":
		res.valueString = abstractgeometrytype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractgeometrytype.Identifier != nil {
			res.valueString = abstractgeometrytype.Identifier.Name
			res.ids = abstractgeometrytype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractgeometrytype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = abstractgeometrytype.SrsName
	}
	return
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractmetadatapropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", abstractmetadatapropertytype.Owns)
		res.valueBool = abstractmetadatapropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstracttimegeometricprimitivetype.Name
	case "Id":
		res.valueString = abstracttimegeometricprimitivetype.Id
	case "Description":
		res.valueString = abstracttimegeometricprimitivetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstracttimegeometricprimitivetype.Identifier != nil {
			res.valueString = abstracttimegeometricprimitivetype.Identifier.Name
			res.ids = abstracttimegeometricprimitivetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstracttimegeometricprimitivetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Frame":
		res.valueString = abstracttimegeometricprimitivetype.Frame
	}
	return
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstracttimeobjecttype.Name
	case "Id":
		res.valueString = abstracttimeobjecttype.Id
	case "Description":
		res.valueString = abstracttimeobjecttype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstracttimeobjecttype.Identifier != nil {
			res.valueString = abstracttimeobjecttype.Identifier.Name
			res.ids = abstracttimeobjecttype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstracttimeobjecttype.NameXSD {
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

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstracttimeprimitivetype.Name
	case "Id":
		res.valueString = abstracttimeprimitivetype.Id
	case "Description":
		res.valueString = abstracttimeprimitivetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstracttimeprimitivetype.Identifier != nil {
			res.valueString = abstracttimeprimitivetype.Identifier.Name
			res.ids = abstracttimeprimitivetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstracttimeprimitivetype.NameXSD {
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

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstracttimeslicetype.Name
	case "Id":
		res.valueString = abstracttimeslicetype.Id
	case "Description":
		res.valueString = abstracttimeslicetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstracttimeslicetype.Identifier != nil {
			res.valueString = abstracttimeslicetype.Identifier.Name
			res.ids = abstracttimeslicetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstracttimeslicetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ValidTime":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstracttimeslicetype.ValidTime != nil {
			res.valueString = abstracttimeslicetype.ValidTime.Name
			res.ids = abstracttimeslicetype.ValidTime.GongGetUUID(stage)
		}
	}
	return
}

func (boundingshapetype *BoundingShapeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = boundingshapetype.Name
	case "NilReason":
		res.valueString = boundingshapetype.NilReason
	case "Envelope":
		res.GongFieldValueType = GongFieldValueTypePointer
		if boundingshapetype.Envelope != nil {
			res.valueString = boundingshapetype.Envelope.Name
			res.ids = boundingshapetype.Envelope.GongGetUUID(stage)
		}
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

func (directpositiontype *DirectPositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = directpositiontype.Name
	case "SrsName":
		res.valueString = directpositiontype.SrsName
	case "EnclosedText":
		res.valueString = directpositiontype.EnclosedText
	}
	return
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dynamicfeaturetype.Name
	case "Id":
		res.valueString = dynamicfeaturetype.Id
	case "Description":
		res.valueString = dynamicfeaturetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dynamicfeaturetype.Identifier != nil {
			res.valueString = dynamicfeaturetype.Identifier.Name
			res.ids = dynamicfeaturetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamicfeaturetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "BoundedBy":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dynamicfeaturetype.BoundedBy != nil {
			res.valueString = dynamicfeaturetype.BoundedBy.Name
			res.ids = dynamicfeaturetype.BoundedBy.GongGetUUID(stage)
		}
	}
	return
}

func (envelopetype *EnvelopeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = envelopetype.Name
	case "SrsName":
		res.valueString = envelopetype.SrsName
	case "LowerCorner":
		res.GongFieldValueType = GongFieldValueTypePointer
		if envelopetype.LowerCorner != nil {
			res.valueString = envelopetype.LowerCorner.Name
			res.ids = envelopetype.LowerCorner.GongGetUUID(stage)
		}
	case "UpperCorner":
		res.GongFieldValueType = GongFieldValueTypePointer
		if envelopetype.UpperCorner != nil {
			res.valueString = envelopetype.UpperCorner.Name
			res.ids = envelopetype.UpperCorner.GongGetUUID(stage)
		}
	}
	return
}

func (pointtype *PointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pointtype.Name
	case "Id":
		res.valueString = pointtype.Id
	case "Description":
		res.valueString = pointtype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pointtype.Identifier != nil {
			res.valueString = pointtype.Identifier.Name
			res.ids = pointtype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range pointtype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = pointtype.SrsName
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pointtype.Pos != nil {
			res.valueString = pointtype.Pos.Name
			res.ids = pointtype.Pos.GongGetUUID(stage)
		}
	}
	return
}

func (timeinstanttype *TimeInstantType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timeinstanttype.Name
	case "Id":
		res.valueString = timeinstanttype.Id
	case "Description":
		res.valueString = timeinstanttype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeinstanttype.Identifier != nil {
			res.valueString = timeinstanttype.Identifier.Name
			res.ids = timeinstanttype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range timeinstanttype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Frame":
		res.valueString = timeinstanttype.Frame
	case "TimePosition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeinstanttype.TimePosition != nil {
			res.valueString = timeinstanttype.TimePosition.Name
			res.ids = timeinstanttype.TimePosition.GongGetUUID(stage)
		}
	}
	return
}

func (timeperiodtype *TimePeriodType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timeperiodtype.Name
	case "Id":
		res.valueString = timeperiodtype.Id
	case "Description":
		res.valueString = timeperiodtype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeperiodtype.Identifier != nil {
			res.valueString = timeperiodtype.Identifier.Name
			res.ids = timeperiodtype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range timeperiodtype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Frame":
		res.valueString = timeperiodtype.Frame
	case "BeginPosition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeperiodtype.BeginPosition != nil {
			res.valueString = timeperiodtype.BeginPosition.Name
			res.ids = timeperiodtype.BeginPosition.GongGetUUID(stage)
		}
	case "EndPosition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeperiodtype.EndPosition != nil {
			res.valueString = timeperiodtype.EndPosition.Name
			res.ids = timeperiodtype.EndPosition.GongGetUUID(stage)
		}
	}
	return
}

func (timepositiontype *TimePositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timepositiontype.Name
	case "Frame":
		res.valueString = timepositiontype.Frame
	case "IndeterminatePosition":
		enum := timepositiontype.IndeterminatePosition
		res.valueString = enum.ToCodeString()
	case "EnclosedText":
		res.valueString = timepositiontype.EnclosedText
	}
	return
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timeprimitivepropertytype.Name
	case "NilReason":
		res.valueString = timeprimitivepropertytype.NilReason
	case "Type_XSD":
		res.valueString = timeprimitivepropertytype.Type_XSD
	case "Href":
		res.valueString = timeprimitivepropertytype.Href
	case "Role":
		res.valueString = timeprimitivepropertytype.Role
	case "Arcrole":
		res.valueString = timeprimitivepropertytype.Arcrole
	case "Title":
		res.valueString = timeprimitivepropertytype.Title
	case "Show":
		res.valueString = timeprimitivepropertytype.Show
	case "Actuate":
		res.valueString = timeprimitivepropertytype.Actuate
	case "TimeInstant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeprimitivepropertytype.TimeInstant != nil {
			res.valueString = timeprimitivepropertytype.TimeInstant.Name
			res.ids = timeprimitivepropertytype.TimeInstant.GongGetUUID(stage)
		}
	case "TimePeriod":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timeprimitivepropertytype.TimePeriod != nil {
			res.valueString = timeprimitivepropertytype.TimePeriod.Name
			res.ids = timeprimitivepropertytype.TimePeriod.GongGetUUID(stage)
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
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetGongstructName() string {
	return "AbstractFeatureMemberType"
}

func (abstractfeaturetype *AbstractFeatureType) GongGetGongstructName() string {
	return "AbstractFeatureType"
}

func (abstractgmltype *AbstractGMLType) GongGetGongstructName() string {
	return "AbstractGMLType"
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetGongstructName() string {
	return "AbstractGeometricPrimitiveType"
}

func (abstractgeometrytype *AbstractGeometryType) GongGetGongstructName() string {
	return "AbstractGeometryType"
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetGongstructName() string {
	return "AbstractMetadataPropertyType"
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetGongstructName() string {
	return "AbstractTimeGeometricPrimitiveType"
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetGongstructName() string {
	return "AbstractTimeObjectType"
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetGongstructName() string {
	return "AbstractTimePrimitiveType"
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetGongstructName() string {
	return "AbstractTimeSliceType"
}

func (boundingshapetype *BoundingShapeType) GongGetGongstructName() string {
	return "BoundingShapeType"
}

func (codetype *CodeType) GongGetGongstructName() string {
	return "CodeType"
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetGongstructName() string {
	return "CodeWithAuthorityType"
}

func (directpositiontype *DirectPositionType) GongGetGongstructName() string {
	return "DirectPositionType"
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetGongstructName() string {
	return "DynamicFeatureType"
}

func (envelopetype *EnvelopeType) GongGetGongstructName() string {
	return "EnvelopeType"
}

func (pointtype *PointType) GongGetGongstructName() string {
	return "PointType"
}

func (timeinstanttype *TimeInstantType) GongGetGongstructName() string {
	return "TimeInstantType"
}

func (timeperiodtype *TimePeriodType) GongGetGongstructName() string {
	return "TimePeriodType"
}

func (timepositiontype *TimePositionType) GongGetGongstructName() string {
	return "TimePositionType"
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetGongstructName() string {
	return "TimePrimitivePropertyType"
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
	__gong__rebuildMapString(stage.AbstractFeatureMemberTypes, &stage.AbstractFeatureMemberTypes_mapString)

	__gong__rebuildMapString(stage.AbstractFeatureTypes, &stage.AbstractFeatureTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGMLTypes, &stage.AbstractGMLTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGeometricPrimitiveTypes, &stage.AbstractGeometricPrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGeometryTypes, &stage.AbstractGeometryTypes_mapString)

	__gong__rebuildMapString(stage.AbstractMetadataPropertyTypes, &stage.AbstractMetadataPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeGeometricPrimitiveTypes, &stage.AbstractTimeGeometricPrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeObjectTypes, &stage.AbstractTimeObjectTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimePrimitiveTypes, &stage.AbstractTimePrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeSliceTypes, &stage.AbstractTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.BoundingShapeTypes, &stage.BoundingShapeTypes_mapString)

	__gong__rebuildMapString(stage.CodeTypes, &stage.CodeTypes_mapString)

	__gong__rebuildMapString(stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString)

	__gong__rebuildMapString(stage.DirectPositionTypes, &stage.DirectPositionTypes_mapString)

	__gong__rebuildMapString(stage.DynamicFeatureTypes, &stage.DynamicFeatureTypes_mapString)

	__gong__rebuildMapString(stage.EnvelopeTypes, &stage.EnvelopeTypes_mapString)

	__gong__rebuildMapString(stage.PointTypes, &stage.PointTypes_mapString)

	__gong__rebuildMapString(stage.TimeInstantTypes, &stage.TimeInstantTypes_mapString)

	__gong__rebuildMapString(stage.TimePeriodTypes, &stage.TimePeriodTypes_mapString)

	__gong__rebuildMapString(stage.TimePositionTypes, &stage.TimePositionTypes_mapString)

	__gong__rebuildMapString(stage.TimePrimitivePropertyTypes, &stage.TimePrimitivePropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
