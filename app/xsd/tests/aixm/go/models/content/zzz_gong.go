// generated code - do not edit
package content

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
	AbstractMD_ContentInformation_Types                map[*AbstractMD_ContentInformation_Type]struct{}
	AbstractMD_ContentInformation_Types_instance       map[*AbstractMD_ContentInformation_Type]*AbstractMD_ContentInformation_Type
	AbstractMD_ContentInformation_Types_mapString      map[string]*AbstractMD_ContentInformation_Type
	AbstractMD_ContentInformation_TypeOrder            uint
	AbstractMD_ContentInformation_Type_stagedOrder     map[*AbstractMD_ContentInformation_Type]uint
	AbstractMD_ContentInformation_Type_orderStaged     map[uint]*AbstractMD_ContentInformation_Type
	AbstractMD_ContentInformation_Types_reference      map[*AbstractMD_ContentInformation_Type]*AbstractMD_ContentInformation_Type
	AbstractMD_ContentInformation_Types_referenceOrder map[*AbstractMD_ContentInformation_Type]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractMD_ContentInformation_TypeCreateCallback GongOnAfterCreateInterface[AbstractMD_ContentInformation_Type]
	OnAfterAbstractMD_ContentInformation_TypeUpdateCallback GongOnAfterUpdateInterface[AbstractMD_ContentInformation_Type]
	OnAfterAbstractMD_ContentInformation_TypeDeleteCallback GongOnAfterDeleteInterface[AbstractMD_ContentInformation_Type]

	MD_Band_PropertyTypes                map[*MD_Band_PropertyType]struct{}
	MD_Band_PropertyTypes_instance       map[*MD_Band_PropertyType]*MD_Band_PropertyType
	MD_Band_PropertyTypes_mapString      map[string]*MD_Band_PropertyType
	MD_Band_PropertyTypeOrder            uint
	MD_Band_PropertyType_stagedOrder     map[*MD_Band_PropertyType]uint
	MD_Band_PropertyType_orderStaged     map[uint]*MD_Band_PropertyType
	MD_Band_PropertyTypes_reference      map[*MD_Band_PropertyType]*MD_Band_PropertyType
	MD_Band_PropertyTypes_referenceOrder map[*MD_Band_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Band_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_Band_PropertyType]
	OnAfterMD_Band_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_Band_PropertyType]
	OnAfterMD_Band_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_Band_PropertyType]

	MD_Band_Types                map[*MD_Band_Type]struct{}
	MD_Band_Types_instance       map[*MD_Band_Type]*MD_Band_Type
	MD_Band_Types_mapString      map[string]*MD_Band_Type
	MD_Band_TypeOrder            uint
	MD_Band_Type_stagedOrder     map[*MD_Band_Type]uint
	MD_Band_Type_orderStaged     map[uint]*MD_Band_Type
	MD_Band_Types_reference      map[*MD_Band_Type]*MD_Band_Type
	MD_Band_Types_referenceOrder map[*MD_Band_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Band_TypeCreateCallback GongOnAfterCreateInterface[MD_Band_Type]
	OnAfterMD_Band_TypeUpdateCallback GongOnAfterUpdateInterface[MD_Band_Type]
	OnAfterMD_Band_TypeDeleteCallback GongOnAfterDeleteInterface[MD_Band_Type]

	MD_ContentInformation_PropertyTypes                map[*MD_ContentInformation_PropertyType]struct{}
	MD_ContentInformation_PropertyTypes_instance       map[*MD_ContentInformation_PropertyType]*MD_ContentInformation_PropertyType
	MD_ContentInformation_PropertyTypes_mapString      map[string]*MD_ContentInformation_PropertyType
	MD_ContentInformation_PropertyTypeOrder            uint
	MD_ContentInformation_PropertyType_stagedOrder     map[*MD_ContentInformation_PropertyType]uint
	MD_ContentInformation_PropertyType_orderStaged     map[uint]*MD_ContentInformation_PropertyType
	MD_ContentInformation_PropertyTypes_reference      map[*MD_ContentInformation_PropertyType]*MD_ContentInformation_PropertyType
	MD_ContentInformation_PropertyTypes_referenceOrder map[*MD_ContentInformation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ContentInformation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ContentInformation_PropertyType]
	OnAfterMD_ContentInformation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ContentInformation_PropertyType]
	OnAfterMD_ContentInformation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ContentInformation_PropertyType]

	MD_CoverageContentTypeCode_PropertyTypes                map[*MD_CoverageContentTypeCode_PropertyType]struct{}
	MD_CoverageContentTypeCode_PropertyTypes_instance       map[*MD_CoverageContentTypeCode_PropertyType]*MD_CoverageContentTypeCode_PropertyType
	MD_CoverageContentTypeCode_PropertyTypes_mapString      map[string]*MD_CoverageContentTypeCode_PropertyType
	MD_CoverageContentTypeCode_PropertyTypeOrder            uint
	MD_CoverageContentTypeCode_PropertyType_stagedOrder     map[*MD_CoverageContentTypeCode_PropertyType]uint
	MD_CoverageContentTypeCode_PropertyType_orderStaged     map[uint]*MD_CoverageContentTypeCode_PropertyType
	MD_CoverageContentTypeCode_PropertyTypes_reference      map[*MD_CoverageContentTypeCode_PropertyType]*MD_CoverageContentTypeCode_PropertyType
	MD_CoverageContentTypeCode_PropertyTypes_referenceOrder map[*MD_CoverageContentTypeCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_CoverageContentTypeCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_CoverageContentTypeCode_PropertyType]
	OnAfterMD_CoverageContentTypeCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_CoverageContentTypeCode_PropertyType]
	OnAfterMD_CoverageContentTypeCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_CoverageContentTypeCode_PropertyType]

	MD_CoverageDescription_PropertyTypes                map[*MD_CoverageDescription_PropertyType]struct{}
	MD_CoverageDescription_PropertyTypes_instance       map[*MD_CoverageDescription_PropertyType]*MD_CoverageDescription_PropertyType
	MD_CoverageDescription_PropertyTypes_mapString      map[string]*MD_CoverageDescription_PropertyType
	MD_CoverageDescription_PropertyTypeOrder            uint
	MD_CoverageDescription_PropertyType_stagedOrder     map[*MD_CoverageDescription_PropertyType]uint
	MD_CoverageDescription_PropertyType_orderStaged     map[uint]*MD_CoverageDescription_PropertyType
	MD_CoverageDescription_PropertyTypes_reference      map[*MD_CoverageDescription_PropertyType]*MD_CoverageDescription_PropertyType
	MD_CoverageDescription_PropertyTypes_referenceOrder map[*MD_CoverageDescription_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_CoverageDescription_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_CoverageDescription_PropertyType]
	OnAfterMD_CoverageDescription_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_CoverageDescription_PropertyType]
	OnAfterMD_CoverageDescription_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_CoverageDescription_PropertyType]

	MD_CoverageDescription_Types                map[*MD_CoverageDescription_Type]struct{}
	MD_CoverageDescription_Types_instance       map[*MD_CoverageDescription_Type]*MD_CoverageDescription_Type
	MD_CoverageDescription_Types_mapString      map[string]*MD_CoverageDescription_Type
	MD_CoverageDescription_TypeOrder            uint
	MD_CoverageDescription_Type_stagedOrder     map[*MD_CoverageDescription_Type]uint
	MD_CoverageDescription_Type_orderStaged     map[uint]*MD_CoverageDescription_Type
	MD_CoverageDescription_Types_reference      map[*MD_CoverageDescription_Type]*MD_CoverageDescription_Type
	MD_CoverageDescription_Types_referenceOrder map[*MD_CoverageDescription_Type]uint

	// insertion point for slice of pointers maps
	MD_CoverageDescription_Type_Dimension_reverseMap map[*MD_RangeDimension_PropertyType]*MD_CoverageDescription_Type

	OnAfterMD_CoverageDescription_TypeCreateCallback GongOnAfterCreateInterface[MD_CoverageDescription_Type]
	OnAfterMD_CoverageDescription_TypeUpdateCallback GongOnAfterUpdateInterface[MD_CoverageDescription_Type]
	OnAfterMD_CoverageDescription_TypeDeleteCallback GongOnAfterDeleteInterface[MD_CoverageDescription_Type]

	MD_FeatureCatalogueDescription_PropertyTypes                map[*MD_FeatureCatalogueDescription_PropertyType]struct{}
	MD_FeatureCatalogueDescription_PropertyTypes_instance       map[*MD_FeatureCatalogueDescription_PropertyType]*MD_FeatureCatalogueDescription_PropertyType
	MD_FeatureCatalogueDescription_PropertyTypes_mapString      map[string]*MD_FeatureCatalogueDescription_PropertyType
	MD_FeatureCatalogueDescription_PropertyTypeOrder            uint
	MD_FeatureCatalogueDescription_PropertyType_stagedOrder     map[*MD_FeatureCatalogueDescription_PropertyType]uint
	MD_FeatureCatalogueDescription_PropertyType_orderStaged     map[uint]*MD_FeatureCatalogueDescription_PropertyType
	MD_FeatureCatalogueDescription_PropertyTypes_reference      map[*MD_FeatureCatalogueDescription_PropertyType]*MD_FeatureCatalogueDescription_PropertyType
	MD_FeatureCatalogueDescription_PropertyTypes_referenceOrder map[*MD_FeatureCatalogueDescription_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_FeatureCatalogueDescription_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_FeatureCatalogueDescription_PropertyType]
	OnAfterMD_FeatureCatalogueDescription_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_FeatureCatalogueDescription_PropertyType]
	OnAfterMD_FeatureCatalogueDescription_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_FeatureCatalogueDescription_PropertyType]

	MD_FeatureCatalogueDescription_Types                map[*MD_FeatureCatalogueDescription_Type]struct{}
	MD_FeatureCatalogueDescription_Types_instance       map[*MD_FeatureCatalogueDescription_Type]*MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription_Types_mapString      map[string]*MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription_TypeOrder            uint
	MD_FeatureCatalogueDescription_Type_stagedOrder     map[*MD_FeatureCatalogueDescription_Type]uint
	MD_FeatureCatalogueDescription_Type_orderStaged     map[uint]*MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription_Types_reference      map[*MD_FeatureCatalogueDescription_Type]*MD_FeatureCatalogueDescription_Type
	MD_FeatureCatalogueDescription_Types_referenceOrder map[*MD_FeatureCatalogueDescription_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_FeatureCatalogueDescription_TypeCreateCallback GongOnAfterCreateInterface[MD_FeatureCatalogueDescription_Type]
	OnAfterMD_FeatureCatalogueDescription_TypeUpdateCallback GongOnAfterUpdateInterface[MD_FeatureCatalogueDescription_Type]
	OnAfterMD_FeatureCatalogueDescription_TypeDeleteCallback GongOnAfterDeleteInterface[MD_FeatureCatalogueDescription_Type]

	MD_ImageDescription_PropertyTypes                map[*MD_ImageDescription_PropertyType]struct{}
	MD_ImageDescription_PropertyTypes_instance       map[*MD_ImageDescription_PropertyType]*MD_ImageDescription_PropertyType
	MD_ImageDescription_PropertyTypes_mapString      map[string]*MD_ImageDescription_PropertyType
	MD_ImageDescription_PropertyTypeOrder            uint
	MD_ImageDescription_PropertyType_stagedOrder     map[*MD_ImageDescription_PropertyType]uint
	MD_ImageDescription_PropertyType_orderStaged     map[uint]*MD_ImageDescription_PropertyType
	MD_ImageDescription_PropertyTypes_reference      map[*MD_ImageDescription_PropertyType]*MD_ImageDescription_PropertyType
	MD_ImageDescription_PropertyTypes_referenceOrder map[*MD_ImageDescription_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ImageDescription_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ImageDescription_PropertyType]
	OnAfterMD_ImageDescription_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ImageDescription_PropertyType]
	OnAfterMD_ImageDescription_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ImageDescription_PropertyType]

	MD_ImageDescription_Types                map[*MD_ImageDescription_Type]struct{}
	MD_ImageDescription_Types_instance       map[*MD_ImageDescription_Type]*MD_ImageDescription_Type
	MD_ImageDescription_Types_mapString      map[string]*MD_ImageDescription_Type
	MD_ImageDescription_TypeOrder            uint
	MD_ImageDescription_Type_stagedOrder     map[*MD_ImageDescription_Type]uint
	MD_ImageDescription_Type_orderStaged     map[uint]*MD_ImageDescription_Type
	MD_ImageDescription_Types_reference      map[*MD_ImageDescription_Type]*MD_ImageDescription_Type
	MD_ImageDescription_Types_referenceOrder map[*MD_ImageDescription_Type]uint

	// insertion point for slice of pointers maps
	MD_ImageDescription_Type_Dimension_reverseMap map[*MD_RangeDimension_PropertyType]*MD_ImageDescription_Type

	OnAfterMD_ImageDescription_TypeCreateCallback GongOnAfterCreateInterface[MD_ImageDescription_Type]
	OnAfterMD_ImageDescription_TypeUpdateCallback GongOnAfterUpdateInterface[MD_ImageDescription_Type]
	OnAfterMD_ImageDescription_TypeDeleteCallback GongOnAfterDeleteInterface[MD_ImageDescription_Type]

	MD_ImagingConditionCode_PropertyTypes                map[*MD_ImagingConditionCode_PropertyType]struct{}
	MD_ImagingConditionCode_PropertyTypes_instance       map[*MD_ImagingConditionCode_PropertyType]*MD_ImagingConditionCode_PropertyType
	MD_ImagingConditionCode_PropertyTypes_mapString      map[string]*MD_ImagingConditionCode_PropertyType
	MD_ImagingConditionCode_PropertyTypeOrder            uint
	MD_ImagingConditionCode_PropertyType_stagedOrder     map[*MD_ImagingConditionCode_PropertyType]uint
	MD_ImagingConditionCode_PropertyType_orderStaged     map[uint]*MD_ImagingConditionCode_PropertyType
	MD_ImagingConditionCode_PropertyTypes_reference      map[*MD_ImagingConditionCode_PropertyType]*MD_ImagingConditionCode_PropertyType
	MD_ImagingConditionCode_PropertyTypes_referenceOrder map[*MD_ImagingConditionCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_ImagingConditionCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_ImagingConditionCode_PropertyType]
	OnAfterMD_ImagingConditionCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_ImagingConditionCode_PropertyType]
	OnAfterMD_ImagingConditionCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_ImagingConditionCode_PropertyType]

	MD_RangeDimension_PropertyTypes                map[*MD_RangeDimension_PropertyType]struct{}
	MD_RangeDimension_PropertyTypes_instance       map[*MD_RangeDimension_PropertyType]*MD_RangeDimension_PropertyType
	MD_RangeDimension_PropertyTypes_mapString      map[string]*MD_RangeDimension_PropertyType
	MD_RangeDimension_PropertyTypeOrder            uint
	MD_RangeDimension_PropertyType_stagedOrder     map[*MD_RangeDimension_PropertyType]uint
	MD_RangeDimension_PropertyType_orderStaged     map[uint]*MD_RangeDimension_PropertyType
	MD_RangeDimension_PropertyTypes_reference      map[*MD_RangeDimension_PropertyType]*MD_RangeDimension_PropertyType
	MD_RangeDimension_PropertyTypes_referenceOrder map[*MD_RangeDimension_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_RangeDimension_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_RangeDimension_PropertyType]
	OnAfterMD_RangeDimension_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_RangeDimension_PropertyType]
	OnAfterMD_RangeDimension_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_RangeDimension_PropertyType]

	MD_RangeDimension_Types                map[*MD_RangeDimension_Type]struct{}
	MD_RangeDimension_Types_instance       map[*MD_RangeDimension_Type]*MD_RangeDimension_Type
	MD_RangeDimension_Types_mapString      map[string]*MD_RangeDimension_Type
	MD_RangeDimension_TypeOrder            uint
	MD_RangeDimension_Type_stagedOrder     map[*MD_RangeDimension_Type]uint
	MD_RangeDimension_Type_orderStaged     map[uint]*MD_RangeDimension_Type
	MD_RangeDimension_Types_reference      map[*MD_RangeDimension_Type]*MD_RangeDimension_Type
	MD_RangeDimension_Types_referenceOrder map[*MD_RangeDimension_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_RangeDimension_TypeCreateCallback GongOnAfterCreateInterface[MD_RangeDimension_Type]
	OnAfterMD_RangeDimension_TypeUpdateCallback GongOnAfterUpdateInterface[MD_RangeDimension_Type]
	OnAfterMD_RangeDimension_TypeDeleteCallback GongOnAfterDeleteInterface[MD_RangeDimension_Type]

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
	__gong__clearReferences(&stage.AbstractMD_ContentInformation_Types_reference, &stage.AbstractMD_ContentInformation_Types_instance, &stage.AbstractMD_ContentInformation_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_Band_PropertyTypes_reference, &stage.MD_Band_PropertyTypes_instance, &stage.MD_Band_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Band_Types_reference, &stage.MD_Band_Types_instance, &stage.MD_Band_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_ContentInformation_PropertyTypes_reference, &stage.MD_ContentInformation_PropertyTypes_instance, &stage.MD_ContentInformation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_CoverageContentTypeCode_PropertyTypes_reference, &stage.MD_CoverageContentTypeCode_PropertyTypes_instance, &stage.MD_CoverageContentTypeCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_CoverageDescription_PropertyTypes_reference, &stage.MD_CoverageDescription_PropertyTypes_instance, &stage.MD_CoverageDescription_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_CoverageDescription_Types_reference, &stage.MD_CoverageDescription_Types_instance, &stage.MD_CoverageDescription_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_FeatureCatalogueDescription_PropertyTypes_reference, &stage.MD_FeatureCatalogueDescription_PropertyTypes_instance, &stage.MD_FeatureCatalogueDescription_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_FeatureCatalogueDescription_Types_reference, &stage.MD_FeatureCatalogueDescription_Types_instance, &stage.MD_FeatureCatalogueDescription_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_ImageDescription_PropertyTypes_reference, &stage.MD_ImageDescription_PropertyTypes_instance, &stage.MD_ImageDescription_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_ImageDescription_Types_reference, &stage.MD_ImageDescription_Types_instance, &stage.MD_ImageDescription_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_ImagingConditionCode_PropertyTypes_reference, &stage.MD_ImagingConditionCode_PropertyTypes_instance, &stage.MD_ImagingConditionCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_RangeDimension_PropertyTypes_reference, &stage.MD_RangeDimension_PropertyTypes_instance, &stage.MD_RangeDimension_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_RangeDimension_Types_reference, &stage.MD_RangeDimension_Types_instance, &stage.MD_RangeDimension_Types_referenceOrder)

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
	stage.AbstractMD_ContentInformation_TypeOrder = __gong__recomputeOrder(stage.AbstractMD_ContentInformation_Type_stagedOrder)

	stage.MD_Band_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_Band_PropertyType_stagedOrder)

	stage.MD_Band_TypeOrder = __gong__recomputeOrder(stage.MD_Band_Type_stagedOrder)

	stage.MD_ContentInformation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ContentInformation_PropertyType_stagedOrder)

	stage.MD_CoverageContentTypeCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder)

	stage.MD_CoverageDescription_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_CoverageDescription_PropertyType_stagedOrder)

	stage.MD_CoverageDescription_TypeOrder = __gong__recomputeOrder(stage.MD_CoverageDescription_Type_stagedOrder)

	stage.MD_FeatureCatalogueDescription_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder)

	stage.MD_FeatureCatalogueDescription_TypeOrder = __gong__recomputeOrder(stage.MD_FeatureCatalogueDescription_Type_stagedOrder)

	stage.MD_ImageDescription_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ImageDescription_PropertyType_stagedOrder)

	stage.MD_ImageDescription_TypeOrder = __gong__recomputeOrder(stage.MD_ImageDescription_Type_stagedOrder)

	stage.MD_ImagingConditionCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_ImagingConditionCode_PropertyType_stagedOrder)

	stage.MD_RangeDimension_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_RangeDimension_PropertyType_stagedOrder)

	stage.MD_RangeDimension_TypeOrder = __gong__recomputeOrder(stage.MD_RangeDimension_Type_stagedOrder)

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
func (*AbstractMD_ContentInformation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Type_stagedOrder)
}

func (*AbstractMD_ContentInformation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractMD_ContentInformation_Type_orderStaged[order]
}

func (*AbstractMD_ContentInformation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractMD_ContentInformation_Types_mapString
}

func (*AbstractMD_ContentInformation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractMD_ContentInformation_Types
}

func (*AbstractMD_ContentInformation_Type) GongNewInstance() any {
	return new(AbstractMD_ContentInformation_Type)
}

func (*MD_Band_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyType_stagedOrder)
}

func (*MD_Band_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Band_PropertyType_orderStaged[order]
}

func (*MD_Band_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Band_PropertyTypes_mapString
}

func (*MD_Band_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Band_PropertyTypes
}

func (*MD_Band_PropertyType) GongNewInstance() any {
	return new(MD_Band_PropertyType)
}

func (*MD_Band_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Band_Types, stage.MD_Band_Type_stagedOrder)
}

func (*MD_Band_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Band_Type_orderStaged[order]
}

func (*MD_Band_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Band_Types_mapString
}

func (*MD_Band_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Band_Types
}

func (*MD_Band_Type) GongNewInstance() any {
	return new(MD_Band_Type)
}

func (*MD_ContentInformation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyType_stagedOrder)
}

func (*MD_ContentInformation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ContentInformation_PropertyType_orderStaged[order]
}

func (*MD_ContentInformation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ContentInformation_PropertyTypes_mapString
}

func (*MD_ContentInformation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ContentInformation_PropertyTypes
}

func (*MD_ContentInformation_PropertyType) GongNewInstance() any {
	return new(MD_ContentInformation_PropertyType)
}

func (*MD_CoverageContentTypeCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder)
}

func (*MD_CoverageContentTypeCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_CoverageContentTypeCode_PropertyType_orderStaged[order]
}

func (*MD_CoverageContentTypeCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_CoverageContentTypeCode_PropertyTypes_mapString
}

func (*MD_CoverageContentTypeCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_CoverageContentTypeCode_PropertyTypes
}

func (*MD_CoverageContentTypeCode_PropertyType) GongNewInstance() any {
	return new(MD_CoverageContentTypeCode_PropertyType)
}

func (*MD_CoverageDescription_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyType_stagedOrder)
}

func (*MD_CoverageDescription_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_CoverageDescription_PropertyType_orderStaged[order]
}

func (*MD_CoverageDescription_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_CoverageDescription_PropertyTypes_mapString
}

func (*MD_CoverageDescription_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_CoverageDescription_PropertyTypes
}

func (*MD_CoverageDescription_PropertyType) GongNewInstance() any {
	return new(MD_CoverageDescription_PropertyType)
}

func (*MD_CoverageDescription_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Type_stagedOrder)
}

func (*MD_CoverageDescription_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_CoverageDescription_Type_orderStaged[order]
}

func (*MD_CoverageDescription_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_CoverageDescription_Types_mapString
}

func (*MD_CoverageDescription_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_CoverageDescription_Types
}

func (*MD_CoverageDescription_Type) GongNewInstance() any {
	return new(MD_CoverageDescription_Type)
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder)
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_FeatureCatalogueDescription_PropertyType_orderStaged[order]
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_FeatureCatalogueDescription_PropertyTypes
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongNewInstance() any {
	return new(MD_FeatureCatalogueDescription_PropertyType)
}

func (*MD_FeatureCatalogueDescription_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Type_stagedOrder)
}

func (*MD_FeatureCatalogueDescription_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_FeatureCatalogueDescription_Type_orderStaged[order]
}

func (*MD_FeatureCatalogueDescription_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_FeatureCatalogueDescription_Types_mapString
}

func (*MD_FeatureCatalogueDescription_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_FeatureCatalogueDescription_Types
}

func (*MD_FeatureCatalogueDescription_Type) GongNewInstance() any {
	return new(MD_FeatureCatalogueDescription_Type)
}

func (*MD_ImageDescription_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyType_stagedOrder)
}

func (*MD_ImageDescription_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ImageDescription_PropertyType_orderStaged[order]
}

func (*MD_ImageDescription_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ImageDescription_PropertyTypes_mapString
}

func (*MD_ImageDescription_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ImageDescription_PropertyTypes
}

func (*MD_ImageDescription_PropertyType) GongNewInstance() any {
	return new(MD_ImageDescription_PropertyType)
}

func (*MD_ImageDescription_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Type_stagedOrder)
}

func (*MD_ImageDescription_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ImageDescription_Type_orderStaged[order]
}

func (*MD_ImageDescription_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ImageDescription_Types_mapString
}

func (*MD_ImageDescription_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ImageDescription_Types
}

func (*MD_ImageDescription_Type) GongNewInstance() any {
	return new(MD_ImageDescription_Type)
}

func (*MD_ImagingConditionCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyType_stagedOrder)
}

func (*MD_ImagingConditionCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_ImagingConditionCode_PropertyType_orderStaged[order]
}

func (*MD_ImagingConditionCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_ImagingConditionCode_PropertyTypes_mapString
}

func (*MD_ImagingConditionCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_ImagingConditionCode_PropertyTypes
}

func (*MD_ImagingConditionCode_PropertyType) GongNewInstance() any {
	return new(MD_ImagingConditionCode_PropertyType)
}

func (*MD_RangeDimension_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyType_stagedOrder)
}

func (*MD_RangeDimension_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_RangeDimension_PropertyType_orderStaged[order]
}

func (*MD_RangeDimension_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_RangeDimension_PropertyTypes_mapString
}

func (*MD_RangeDimension_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_RangeDimension_PropertyTypes
}

func (*MD_RangeDimension_PropertyType) GongNewInstance() any {
	return new(MD_RangeDimension_PropertyType)
}

func (*MD_RangeDimension_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Type_stagedOrder)
}

func (*MD_RangeDimension_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_RangeDimension_Type_orderStaged[order]
}

func (*MD_RangeDimension_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_RangeDimension_Types_mapString
}

func (*MD_RangeDimension_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_RangeDimension_Types
}

func (*MD_RangeDimension_Type) GongNewInstance() any {
	return new(MD_RangeDimension_Type)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content"
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
		AbstractMD_ContentInformation_Types:           make(map[*AbstractMD_ContentInformation_Type]struct{}),
		AbstractMD_ContentInformation_Types_mapString: make(map[string]*AbstractMD_ContentInformation_Type),

		MD_Band_PropertyTypes:           make(map[*MD_Band_PropertyType]struct{}),
		MD_Band_PropertyTypes_mapString: make(map[string]*MD_Band_PropertyType),

		MD_Band_Types:           make(map[*MD_Band_Type]struct{}),
		MD_Band_Types_mapString: make(map[string]*MD_Band_Type),

		MD_ContentInformation_PropertyTypes:           make(map[*MD_ContentInformation_PropertyType]struct{}),
		MD_ContentInformation_PropertyTypes_mapString: make(map[string]*MD_ContentInformation_PropertyType),

		MD_CoverageContentTypeCode_PropertyTypes:           make(map[*MD_CoverageContentTypeCode_PropertyType]struct{}),
		MD_CoverageContentTypeCode_PropertyTypes_mapString: make(map[string]*MD_CoverageContentTypeCode_PropertyType),

		MD_CoverageDescription_PropertyTypes:           make(map[*MD_CoverageDescription_PropertyType]struct{}),
		MD_CoverageDescription_PropertyTypes_mapString: make(map[string]*MD_CoverageDescription_PropertyType),

		MD_CoverageDescription_Types:           make(map[*MD_CoverageDescription_Type]struct{}),
		MD_CoverageDescription_Types_mapString: make(map[string]*MD_CoverageDescription_Type),

		MD_FeatureCatalogueDescription_PropertyTypes:           make(map[*MD_FeatureCatalogueDescription_PropertyType]struct{}),
		MD_FeatureCatalogueDescription_PropertyTypes_mapString: make(map[string]*MD_FeatureCatalogueDescription_PropertyType),

		MD_FeatureCatalogueDescription_Types:           make(map[*MD_FeatureCatalogueDescription_Type]struct{}),
		MD_FeatureCatalogueDescription_Types_mapString: make(map[string]*MD_FeatureCatalogueDescription_Type),

		MD_ImageDescription_PropertyTypes:           make(map[*MD_ImageDescription_PropertyType]struct{}),
		MD_ImageDescription_PropertyTypes_mapString: make(map[string]*MD_ImageDescription_PropertyType),

		MD_ImageDescription_Types:           make(map[*MD_ImageDescription_Type]struct{}),
		MD_ImageDescription_Types_mapString: make(map[string]*MD_ImageDescription_Type),

		MD_ImagingConditionCode_PropertyTypes:           make(map[*MD_ImagingConditionCode_PropertyType]struct{}),
		MD_ImagingConditionCode_PropertyTypes_mapString: make(map[string]*MD_ImagingConditionCode_PropertyType),

		MD_RangeDimension_PropertyTypes:           make(map[*MD_RangeDimension_PropertyType]struct{}),
		MD_RangeDimension_PropertyTypes_mapString: make(map[string]*MD_RangeDimension_PropertyType),

		MD_RangeDimension_Types:           make(map[*MD_RangeDimension_Type]struct{}),
		MD_RangeDimension_Types_mapString: make(map[string]*MD_RangeDimension_Type),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractMD_ContentInformation_Type_stagedOrder: make(map[*AbstractMD_ContentInformation_Type]uint),
		AbstractMD_ContentInformation_Type_orderStaged: make(map[uint]*AbstractMD_ContentInformation_Type),
		AbstractMD_ContentInformation_Types_reference:  make(map[*AbstractMD_ContentInformation_Type]*AbstractMD_ContentInformation_Type),

		MD_Band_PropertyType_stagedOrder: make(map[*MD_Band_PropertyType]uint),
		MD_Band_PropertyType_orderStaged: make(map[uint]*MD_Band_PropertyType),
		MD_Band_PropertyTypes_reference:  make(map[*MD_Band_PropertyType]*MD_Band_PropertyType),

		MD_Band_Type_stagedOrder: make(map[*MD_Band_Type]uint),
		MD_Band_Type_orderStaged: make(map[uint]*MD_Band_Type),
		MD_Band_Types_reference:  make(map[*MD_Band_Type]*MD_Band_Type),

		MD_ContentInformation_PropertyType_stagedOrder: make(map[*MD_ContentInformation_PropertyType]uint),
		MD_ContentInformation_PropertyType_orderStaged: make(map[uint]*MD_ContentInformation_PropertyType),
		MD_ContentInformation_PropertyTypes_reference:  make(map[*MD_ContentInformation_PropertyType]*MD_ContentInformation_PropertyType),

		MD_CoverageContentTypeCode_PropertyType_stagedOrder: make(map[*MD_CoverageContentTypeCode_PropertyType]uint),
		MD_CoverageContentTypeCode_PropertyType_orderStaged: make(map[uint]*MD_CoverageContentTypeCode_PropertyType),
		MD_CoverageContentTypeCode_PropertyTypes_reference:  make(map[*MD_CoverageContentTypeCode_PropertyType]*MD_CoverageContentTypeCode_PropertyType),

		MD_CoverageDescription_PropertyType_stagedOrder: make(map[*MD_CoverageDescription_PropertyType]uint),
		MD_CoverageDescription_PropertyType_orderStaged: make(map[uint]*MD_CoverageDescription_PropertyType),
		MD_CoverageDescription_PropertyTypes_reference:  make(map[*MD_CoverageDescription_PropertyType]*MD_CoverageDescription_PropertyType),

		MD_CoverageDescription_Type_stagedOrder: make(map[*MD_CoverageDescription_Type]uint),
		MD_CoverageDescription_Type_orderStaged: make(map[uint]*MD_CoverageDescription_Type),
		MD_CoverageDescription_Types_reference:  make(map[*MD_CoverageDescription_Type]*MD_CoverageDescription_Type),

		MD_FeatureCatalogueDescription_PropertyType_stagedOrder: make(map[*MD_FeatureCatalogueDescription_PropertyType]uint),
		MD_FeatureCatalogueDescription_PropertyType_orderStaged: make(map[uint]*MD_FeatureCatalogueDescription_PropertyType),
		MD_FeatureCatalogueDescription_PropertyTypes_reference:  make(map[*MD_FeatureCatalogueDescription_PropertyType]*MD_FeatureCatalogueDescription_PropertyType),

		MD_FeatureCatalogueDescription_Type_stagedOrder: make(map[*MD_FeatureCatalogueDescription_Type]uint),
		MD_FeatureCatalogueDescription_Type_orderStaged: make(map[uint]*MD_FeatureCatalogueDescription_Type),
		MD_FeatureCatalogueDescription_Types_reference:  make(map[*MD_FeatureCatalogueDescription_Type]*MD_FeatureCatalogueDescription_Type),

		MD_ImageDescription_PropertyType_stagedOrder: make(map[*MD_ImageDescription_PropertyType]uint),
		MD_ImageDescription_PropertyType_orderStaged: make(map[uint]*MD_ImageDescription_PropertyType),
		MD_ImageDescription_PropertyTypes_reference:  make(map[*MD_ImageDescription_PropertyType]*MD_ImageDescription_PropertyType),

		MD_ImageDescription_Type_stagedOrder: make(map[*MD_ImageDescription_Type]uint),
		MD_ImageDescription_Type_orderStaged: make(map[uint]*MD_ImageDescription_Type),
		MD_ImageDescription_Types_reference:  make(map[*MD_ImageDescription_Type]*MD_ImageDescription_Type),

		MD_ImagingConditionCode_PropertyType_stagedOrder: make(map[*MD_ImagingConditionCode_PropertyType]uint),
		MD_ImagingConditionCode_PropertyType_orderStaged: make(map[uint]*MD_ImagingConditionCode_PropertyType),
		MD_ImagingConditionCode_PropertyTypes_reference:  make(map[*MD_ImagingConditionCode_PropertyType]*MD_ImagingConditionCode_PropertyType),

		MD_RangeDimension_PropertyType_stagedOrder: make(map[*MD_RangeDimension_PropertyType]uint),
		MD_RangeDimension_PropertyType_orderStaged: make(map[uint]*MD_RangeDimension_PropertyType),
		MD_RangeDimension_PropertyTypes_reference:  make(map[*MD_RangeDimension_PropertyType]*MD_RangeDimension_PropertyType),

		MD_RangeDimension_Type_stagedOrder: make(map[*MD_RangeDimension_Type]uint),
		MD_RangeDimension_Type_orderStaged: make(map[uint]*MD_RangeDimension_Type),
		MD_RangeDimension_Types_reference:  make(map[*MD_RangeDimension_Type]*MD_RangeDimension_Type),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractMD_ContentInformation_Type": &AbstractMD_ContentInformation_TypeUnmarshaller{},

			"MD_Band_PropertyType": &MD_Band_PropertyTypeUnmarshaller{},

			"MD_Band_Type": &MD_Band_TypeUnmarshaller{},

			"MD_ContentInformation_PropertyType": &MD_ContentInformation_PropertyTypeUnmarshaller{},

			"MD_CoverageContentTypeCode_PropertyType": &MD_CoverageContentTypeCode_PropertyTypeUnmarshaller{},

			"MD_CoverageDescription_PropertyType": &MD_CoverageDescription_PropertyTypeUnmarshaller{},

			"MD_CoverageDescription_Type": &MD_CoverageDescription_TypeUnmarshaller{},

			"MD_FeatureCatalogueDescription_PropertyType": &MD_FeatureCatalogueDescription_PropertyTypeUnmarshaller{},

			"MD_FeatureCatalogueDescription_Type": &MD_FeatureCatalogueDescription_TypeUnmarshaller{},

			"MD_ImageDescription_PropertyType": &MD_ImageDescription_PropertyTypeUnmarshaller{},

			"MD_ImageDescription_Type": &MD_ImageDescription_TypeUnmarshaller{},

			"MD_ImagingConditionCode_PropertyType": &MD_ImagingConditionCode_PropertyTypeUnmarshaller{},

			"MD_RangeDimension_PropertyType": &MD_RangeDimension_PropertyTypeUnmarshaller{},

			"MD_RangeDimension_Type": &MD_RangeDimension_TypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractMD_ContentInformation_Type"] = len(stage.AbstractMD_ContentInformation_Types)
	stage.Map_GongStructName_InstancesNb["MD_Band_PropertyType"] = len(stage.MD_Band_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Band_Type"] = len(stage.MD_Band_Types)
	stage.Map_GongStructName_InstancesNb["MD_ContentInformation_PropertyType"] = len(stage.MD_ContentInformation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_CoverageContentTypeCode_PropertyType"] = len(stage.MD_CoverageContentTypeCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_CoverageDescription_PropertyType"] = len(stage.MD_CoverageDescription_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_CoverageDescription_Type"] = len(stage.MD_CoverageDescription_Types)
	stage.Map_GongStructName_InstancesNb["MD_FeatureCatalogueDescription_PropertyType"] = len(stage.MD_FeatureCatalogueDescription_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_FeatureCatalogueDescription_Type"] = len(stage.MD_FeatureCatalogueDescription_Types)
	stage.Map_GongStructName_InstancesNb["MD_ImageDescription_PropertyType"] = len(stage.MD_ImageDescription_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_ImageDescription_Type"] = len(stage.MD_ImageDescription_Types)
	stage.Map_GongStructName_InstancesNb["MD_ImagingConditionCode_PropertyType"] = len(stage.MD_ImagingConditionCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_RangeDimension_PropertyType"] = len(stage.MD_RangeDimension_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_RangeDimension_Type"] = len(stage.MD_RangeDimension_Types)
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
// Stage puts abstractmd_contentinformation_type to the model stage
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) Stage(stage *Stage) *AbstractMD_ContentInformation_Type {
	__gong__stage(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Type_stagedOrder, stage.AbstractMD_ContentInformation_Type_orderStaged, &stage.AbstractMD_ContentInformation_TypeOrder, stage.AbstractMD_ContentInformation_Types_mapString, abstractmd_contentinformation_type, abstractmd_contentinformation_type.Name)
	return abstractmd_contentinformation_type
}

// StagePreserveOrder puts abstractmd_contentinformation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractMD_ContentInformation_TypeOrder
// - update stage.AbstractMD_ContentInformation_TypeOrder accordingly
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Type_stagedOrder, stage.AbstractMD_ContentInformation_Type_orderStaged, &stage.AbstractMD_ContentInformation_TypeOrder, stage.AbstractMD_ContentInformation_Types_mapString, abstractmd_contentinformation_type, order, abstractmd_contentinformation_type.Name)
}

// Unstage removes abstractmd_contentinformation_type off the model stage
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) Unstage(stage *Stage) *AbstractMD_ContentInformation_Type {
	__gong__unstage(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Types_mapString, abstractmd_contentinformation_type, abstractmd_contentinformation_type.Name)
	return abstractmd_contentinformation_type
}

// UnstageVoid removes abstractmd_contentinformation_type off the model stage
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) UnstageVoid(stage *Stage) {
	abstractmd_contentinformation_type.Unstage(stage)
}

func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) StageVoid(stage *Stage) {
	abstractmd_contentinformation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GetName() (res string) {
	return abstractmd_contentinformation_type.Name
}

// for satisfaction of GongStruct interface
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) SetName(name string) {
	abstractmd_contentinformation_type.Name = name
}

// Stage puts md_band_propertytype to the model stage
func (md_band_propertytype *MD_Band_PropertyType) Stage(stage *Stage) *MD_Band_PropertyType {
	__gong__stage(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyType_stagedOrder, stage.MD_Band_PropertyType_orderStaged, &stage.MD_Band_PropertyTypeOrder, stage.MD_Band_PropertyTypes_mapString, md_band_propertytype, md_band_propertytype.Name)
	return md_band_propertytype
}

// StagePreserveOrder puts md_band_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Band_PropertyTypeOrder
// - update stage.MD_Band_PropertyTypeOrder accordingly
func (md_band_propertytype *MD_Band_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyType_stagedOrder, stage.MD_Band_PropertyType_orderStaged, &stage.MD_Band_PropertyTypeOrder, stage.MD_Band_PropertyTypes_mapString, md_band_propertytype, order, md_band_propertytype.Name)
}

// Unstage removes md_band_propertytype off the model stage
func (md_band_propertytype *MD_Band_PropertyType) Unstage(stage *Stage) *MD_Band_PropertyType {
	__gong__unstage(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyTypes_mapString, md_band_propertytype, md_band_propertytype.Name)
	return md_band_propertytype
}

// UnstageVoid removes md_band_propertytype off the model stage
func (md_band_propertytype *MD_Band_PropertyType) UnstageVoid(stage *Stage) {
	md_band_propertytype.Unstage(stage)
}

func (md_band_propertytype *MD_Band_PropertyType) StageVoid(stage *Stage) {
	md_band_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_band_propertytype *MD_Band_PropertyType) GetName() (res string) {
	return md_band_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_band_propertytype *MD_Band_PropertyType) SetName(name string) {
	md_band_propertytype.Name = name
}

// Stage puts md_band_type to the model stage
func (md_band_type *MD_Band_Type) Stage(stage *Stage) *MD_Band_Type {
	__gong__stage(stage.MD_Band_Types, stage.MD_Band_Type_stagedOrder, stage.MD_Band_Type_orderStaged, &stage.MD_Band_TypeOrder, stage.MD_Band_Types_mapString, md_band_type, md_band_type.Name)
	return md_band_type
}

// StagePreserveOrder puts md_band_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Band_TypeOrder
// - update stage.MD_Band_TypeOrder accordingly
func (md_band_type *MD_Band_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Band_Types, stage.MD_Band_Type_stagedOrder, stage.MD_Band_Type_orderStaged, &stage.MD_Band_TypeOrder, stage.MD_Band_Types_mapString, md_band_type, order, md_band_type.Name)
}

// Unstage removes md_band_type off the model stage
func (md_band_type *MD_Band_Type) Unstage(stage *Stage) *MD_Band_Type {
	__gong__unstage(stage.MD_Band_Types, stage.MD_Band_Types_mapString, md_band_type, md_band_type.Name)
	return md_band_type
}

// UnstageVoid removes md_band_type off the model stage
func (md_band_type *MD_Band_Type) UnstageVoid(stage *Stage) {
	md_band_type.Unstage(stage)
}

func (md_band_type *MD_Band_Type) StageVoid(stage *Stage) {
	md_band_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_band_type *MD_Band_Type) GetName() (res string) {
	return md_band_type.Name
}

// for satisfaction of GongStruct interface
func (md_band_type *MD_Band_Type) SetName(name string) {
	md_band_type.Name = name
}

// Stage puts md_contentinformation_propertytype to the model stage
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) Stage(stage *Stage) *MD_ContentInformation_PropertyType {
	__gong__stage(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyType_stagedOrder, stage.MD_ContentInformation_PropertyType_orderStaged, &stage.MD_ContentInformation_PropertyTypeOrder, stage.MD_ContentInformation_PropertyTypes_mapString, md_contentinformation_propertytype, md_contentinformation_propertytype.Name)
	return md_contentinformation_propertytype
}

// StagePreserveOrder puts md_contentinformation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ContentInformation_PropertyTypeOrder
// - update stage.MD_ContentInformation_PropertyTypeOrder accordingly
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyType_stagedOrder, stage.MD_ContentInformation_PropertyType_orderStaged, &stage.MD_ContentInformation_PropertyTypeOrder, stage.MD_ContentInformation_PropertyTypes_mapString, md_contentinformation_propertytype, order, md_contentinformation_propertytype.Name)
}

// Unstage removes md_contentinformation_propertytype off the model stage
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) Unstage(stage *Stage) *MD_ContentInformation_PropertyType {
	__gong__unstage(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyTypes_mapString, md_contentinformation_propertytype, md_contentinformation_propertytype.Name)
	return md_contentinformation_propertytype
}

// UnstageVoid removes md_contentinformation_propertytype off the model stage
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) UnstageVoid(stage *Stage) {
	md_contentinformation_propertytype.Unstage(stage)
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) StageVoid(stage *Stage) {
	md_contentinformation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GetName() (res string) {
	return md_contentinformation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) SetName(name string) {
	md_contentinformation_propertytype.Name = name
}

// Stage puts md_coveragecontenttypecode_propertytype to the model stage
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) Stage(stage *Stage) *MD_CoverageContentTypeCode_PropertyType {
	__gong__stage(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder, stage.MD_CoverageContentTypeCode_PropertyType_orderStaged, &stage.MD_CoverageContentTypeCode_PropertyTypeOrder, stage.MD_CoverageContentTypeCode_PropertyTypes_mapString, md_coveragecontenttypecode_propertytype, md_coveragecontenttypecode_propertytype.Name)
	return md_coveragecontenttypecode_propertytype
}

// StagePreserveOrder puts md_coveragecontenttypecode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_CoverageContentTypeCode_PropertyTypeOrder
// - update stage.MD_CoverageContentTypeCode_PropertyTypeOrder accordingly
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder, stage.MD_CoverageContentTypeCode_PropertyType_orderStaged, &stage.MD_CoverageContentTypeCode_PropertyTypeOrder, stage.MD_CoverageContentTypeCode_PropertyTypes_mapString, md_coveragecontenttypecode_propertytype, order, md_coveragecontenttypecode_propertytype.Name)
}

// Unstage removes md_coveragecontenttypecode_propertytype off the model stage
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) Unstage(stage *Stage) *MD_CoverageContentTypeCode_PropertyType {
	__gong__unstage(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyTypes_mapString, md_coveragecontenttypecode_propertytype, md_coveragecontenttypecode_propertytype.Name)
	return md_coveragecontenttypecode_propertytype
}

// UnstageVoid removes md_coveragecontenttypecode_propertytype off the model stage
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) UnstageVoid(stage *Stage) {
	md_coveragecontenttypecode_propertytype.Unstage(stage)
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) StageVoid(stage *Stage) {
	md_coveragecontenttypecode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GetName() (res string) {
	return md_coveragecontenttypecode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) SetName(name string) {
	md_coveragecontenttypecode_propertytype.Name = name
}

// Stage puts md_coveragedescription_propertytype to the model stage
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) Stage(stage *Stage) *MD_CoverageDescription_PropertyType {
	__gong__stage(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyType_stagedOrder, stage.MD_CoverageDescription_PropertyType_orderStaged, &stage.MD_CoverageDescription_PropertyTypeOrder, stage.MD_CoverageDescription_PropertyTypes_mapString, md_coveragedescription_propertytype, md_coveragedescription_propertytype.Name)
	return md_coveragedescription_propertytype
}

// StagePreserveOrder puts md_coveragedescription_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_CoverageDescription_PropertyTypeOrder
// - update stage.MD_CoverageDescription_PropertyTypeOrder accordingly
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyType_stagedOrder, stage.MD_CoverageDescription_PropertyType_orderStaged, &stage.MD_CoverageDescription_PropertyTypeOrder, stage.MD_CoverageDescription_PropertyTypes_mapString, md_coveragedescription_propertytype, order, md_coveragedescription_propertytype.Name)
}

// Unstage removes md_coveragedescription_propertytype off the model stage
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) Unstage(stage *Stage) *MD_CoverageDescription_PropertyType {
	__gong__unstage(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyTypes_mapString, md_coveragedescription_propertytype, md_coveragedescription_propertytype.Name)
	return md_coveragedescription_propertytype
}

// UnstageVoid removes md_coveragedescription_propertytype off the model stage
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) UnstageVoid(stage *Stage) {
	md_coveragedescription_propertytype.Unstage(stage)
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) StageVoid(stage *Stage) {
	md_coveragedescription_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GetName() (res string) {
	return md_coveragedescription_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) SetName(name string) {
	md_coveragedescription_propertytype.Name = name
}

// Stage puts md_coveragedescription_type to the model stage
func (md_coveragedescription_type *MD_CoverageDescription_Type) Stage(stage *Stage) *MD_CoverageDescription_Type {
	__gong__stage(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Type_stagedOrder, stage.MD_CoverageDescription_Type_orderStaged, &stage.MD_CoverageDescription_TypeOrder, stage.MD_CoverageDescription_Types_mapString, md_coveragedescription_type, md_coveragedescription_type.Name)
	return md_coveragedescription_type
}

// StagePreserveOrder puts md_coveragedescription_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_CoverageDescription_TypeOrder
// - update stage.MD_CoverageDescription_TypeOrder accordingly
func (md_coveragedescription_type *MD_CoverageDescription_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Type_stagedOrder, stage.MD_CoverageDescription_Type_orderStaged, &stage.MD_CoverageDescription_TypeOrder, stage.MD_CoverageDescription_Types_mapString, md_coveragedescription_type, order, md_coveragedescription_type.Name)
}

// Unstage removes md_coveragedescription_type off the model stage
func (md_coveragedescription_type *MD_CoverageDescription_Type) Unstage(stage *Stage) *MD_CoverageDescription_Type {
	__gong__unstage(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Types_mapString, md_coveragedescription_type, md_coveragedescription_type.Name)
	return md_coveragedescription_type
}

// UnstageVoid removes md_coveragedescription_type off the model stage
func (md_coveragedescription_type *MD_CoverageDescription_Type) UnstageVoid(stage *Stage) {
	md_coveragedescription_type.Unstage(stage)
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) StageVoid(stage *Stage) {
	md_coveragedescription_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_coveragedescription_type *MD_CoverageDescription_Type) GetName() (res string) {
	return md_coveragedescription_type.Name
}

// for satisfaction of GongStruct interface
func (md_coveragedescription_type *MD_CoverageDescription_Type) SetName(name string) {
	md_coveragedescription_type.Name = name
}

// Stage puts md_featurecataloguedescription_propertytype to the model stage
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) Stage(stage *Stage) *MD_FeatureCatalogueDescription_PropertyType {
	__gong__stage(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder, stage.MD_FeatureCatalogueDescription_PropertyType_orderStaged, &stage.MD_FeatureCatalogueDescription_PropertyTypeOrder, stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString, md_featurecataloguedescription_propertytype, md_featurecataloguedescription_propertytype.Name)
	return md_featurecataloguedescription_propertytype
}

// StagePreserveOrder puts md_featurecataloguedescription_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_FeatureCatalogueDescription_PropertyTypeOrder
// - update stage.MD_FeatureCatalogueDescription_PropertyTypeOrder accordingly
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder, stage.MD_FeatureCatalogueDescription_PropertyType_orderStaged, &stage.MD_FeatureCatalogueDescription_PropertyTypeOrder, stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString, md_featurecataloguedescription_propertytype, order, md_featurecataloguedescription_propertytype.Name)
}

// Unstage removes md_featurecataloguedescription_propertytype off the model stage
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) Unstage(stage *Stage) *MD_FeatureCatalogueDescription_PropertyType {
	__gong__unstage(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString, md_featurecataloguedescription_propertytype, md_featurecataloguedescription_propertytype.Name)
	return md_featurecataloguedescription_propertytype
}

// UnstageVoid removes md_featurecataloguedescription_propertytype off the model stage
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) UnstageVoid(stage *Stage) {
	md_featurecataloguedescription_propertytype.Unstage(stage)
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) StageVoid(stage *Stage) {
	md_featurecataloguedescription_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GetName() (res string) {
	return md_featurecataloguedescription_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) SetName(name string) {
	md_featurecataloguedescription_propertytype.Name = name
}

// Stage puts md_featurecataloguedescription_type to the model stage
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) Stage(stage *Stage) *MD_FeatureCatalogueDescription_Type {
	__gong__stage(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Type_stagedOrder, stage.MD_FeatureCatalogueDescription_Type_orderStaged, &stage.MD_FeatureCatalogueDescription_TypeOrder, stage.MD_FeatureCatalogueDescription_Types_mapString, md_featurecataloguedescription_type, md_featurecataloguedescription_type.Name)
	return md_featurecataloguedescription_type
}

// StagePreserveOrder puts md_featurecataloguedescription_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_FeatureCatalogueDescription_TypeOrder
// - update stage.MD_FeatureCatalogueDescription_TypeOrder accordingly
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Type_stagedOrder, stage.MD_FeatureCatalogueDescription_Type_orderStaged, &stage.MD_FeatureCatalogueDescription_TypeOrder, stage.MD_FeatureCatalogueDescription_Types_mapString, md_featurecataloguedescription_type, order, md_featurecataloguedescription_type.Name)
}

// Unstage removes md_featurecataloguedescription_type off the model stage
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) Unstage(stage *Stage) *MD_FeatureCatalogueDescription_Type {
	__gong__unstage(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Types_mapString, md_featurecataloguedescription_type, md_featurecataloguedescription_type.Name)
	return md_featurecataloguedescription_type
}

// UnstageVoid removes md_featurecataloguedescription_type off the model stage
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) UnstageVoid(stage *Stage) {
	md_featurecataloguedescription_type.Unstage(stage)
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) StageVoid(stage *Stage) {
	md_featurecataloguedescription_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GetName() (res string) {
	return md_featurecataloguedescription_type.Name
}

// for satisfaction of GongStruct interface
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) SetName(name string) {
	md_featurecataloguedescription_type.Name = name
}

// Stage puts md_imagedescription_propertytype to the model stage
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) Stage(stage *Stage) *MD_ImageDescription_PropertyType {
	__gong__stage(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyType_stagedOrder, stage.MD_ImageDescription_PropertyType_orderStaged, &stage.MD_ImageDescription_PropertyTypeOrder, stage.MD_ImageDescription_PropertyTypes_mapString, md_imagedescription_propertytype, md_imagedescription_propertytype.Name)
	return md_imagedescription_propertytype
}

// StagePreserveOrder puts md_imagedescription_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ImageDescription_PropertyTypeOrder
// - update stage.MD_ImageDescription_PropertyTypeOrder accordingly
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyType_stagedOrder, stage.MD_ImageDescription_PropertyType_orderStaged, &stage.MD_ImageDescription_PropertyTypeOrder, stage.MD_ImageDescription_PropertyTypes_mapString, md_imagedescription_propertytype, order, md_imagedescription_propertytype.Name)
}

// Unstage removes md_imagedescription_propertytype off the model stage
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) Unstage(stage *Stage) *MD_ImageDescription_PropertyType {
	__gong__unstage(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyTypes_mapString, md_imagedescription_propertytype, md_imagedescription_propertytype.Name)
	return md_imagedescription_propertytype
}

// UnstageVoid removes md_imagedescription_propertytype off the model stage
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) UnstageVoid(stage *Stage) {
	md_imagedescription_propertytype.Unstage(stage)
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) StageVoid(stage *Stage) {
	md_imagedescription_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GetName() (res string) {
	return md_imagedescription_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) SetName(name string) {
	md_imagedescription_propertytype.Name = name
}

// Stage puts md_imagedescription_type to the model stage
func (md_imagedescription_type *MD_ImageDescription_Type) Stage(stage *Stage) *MD_ImageDescription_Type {
	__gong__stage(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Type_stagedOrder, stage.MD_ImageDescription_Type_orderStaged, &stage.MD_ImageDescription_TypeOrder, stage.MD_ImageDescription_Types_mapString, md_imagedescription_type, md_imagedescription_type.Name)
	return md_imagedescription_type
}

// StagePreserveOrder puts md_imagedescription_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ImageDescription_TypeOrder
// - update stage.MD_ImageDescription_TypeOrder accordingly
func (md_imagedescription_type *MD_ImageDescription_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Type_stagedOrder, stage.MD_ImageDescription_Type_orderStaged, &stage.MD_ImageDescription_TypeOrder, stage.MD_ImageDescription_Types_mapString, md_imagedescription_type, order, md_imagedescription_type.Name)
}

// Unstage removes md_imagedescription_type off the model stage
func (md_imagedescription_type *MD_ImageDescription_Type) Unstage(stage *Stage) *MD_ImageDescription_Type {
	__gong__unstage(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Types_mapString, md_imagedescription_type, md_imagedescription_type.Name)
	return md_imagedescription_type
}

// UnstageVoid removes md_imagedescription_type off the model stage
func (md_imagedescription_type *MD_ImageDescription_Type) UnstageVoid(stage *Stage) {
	md_imagedescription_type.Unstage(stage)
}

func (md_imagedescription_type *MD_ImageDescription_Type) StageVoid(stage *Stage) {
	md_imagedescription_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_imagedescription_type *MD_ImageDescription_Type) GetName() (res string) {
	return md_imagedescription_type.Name
}

// for satisfaction of GongStruct interface
func (md_imagedescription_type *MD_ImageDescription_Type) SetName(name string) {
	md_imagedescription_type.Name = name
}

// Stage puts md_imagingconditioncode_propertytype to the model stage
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) Stage(stage *Stage) *MD_ImagingConditionCode_PropertyType {
	__gong__stage(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyType_stagedOrder, stage.MD_ImagingConditionCode_PropertyType_orderStaged, &stage.MD_ImagingConditionCode_PropertyTypeOrder, stage.MD_ImagingConditionCode_PropertyTypes_mapString, md_imagingconditioncode_propertytype, md_imagingconditioncode_propertytype.Name)
	return md_imagingconditioncode_propertytype
}

// StagePreserveOrder puts md_imagingconditioncode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_ImagingConditionCode_PropertyTypeOrder
// - update stage.MD_ImagingConditionCode_PropertyTypeOrder accordingly
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyType_stagedOrder, stage.MD_ImagingConditionCode_PropertyType_orderStaged, &stage.MD_ImagingConditionCode_PropertyTypeOrder, stage.MD_ImagingConditionCode_PropertyTypes_mapString, md_imagingconditioncode_propertytype, order, md_imagingconditioncode_propertytype.Name)
}

// Unstage removes md_imagingconditioncode_propertytype off the model stage
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) Unstage(stage *Stage) *MD_ImagingConditionCode_PropertyType {
	__gong__unstage(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyTypes_mapString, md_imagingconditioncode_propertytype, md_imagingconditioncode_propertytype.Name)
	return md_imagingconditioncode_propertytype
}

// UnstageVoid removes md_imagingconditioncode_propertytype off the model stage
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) UnstageVoid(stage *Stage) {
	md_imagingconditioncode_propertytype.Unstage(stage)
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) StageVoid(stage *Stage) {
	md_imagingconditioncode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GetName() (res string) {
	return md_imagingconditioncode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) SetName(name string) {
	md_imagingconditioncode_propertytype.Name = name
}

// Stage puts md_rangedimension_propertytype to the model stage
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) Stage(stage *Stage) *MD_RangeDimension_PropertyType {
	__gong__stage(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyType_stagedOrder, stage.MD_RangeDimension_PropertyType_orderStaged, &stage.MD_RangeDimension_PropertyTypeOrder, stage.MD_RangeDimension_PropertyTypes_mapString, md_rangedimension_propertytype, md_rangedimension_propertytype.Name)
	return md_rangedimension_propertytype
}

// StagePreserveOrder puts md_rangedimension_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_RangeDimension_PropertyTypeOrder
// - update stage.MD_RangeDimension_PropertyTypeOrder accordingly
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyType_stagedOrder, stage.MD_RangeDimension_PropertyType_orderStaged, &stage.MD_RangeDimension_PropertyTypeOrder, stage.MD_RangeDimension_PropertyTypes_mapString, md_rangedimension_propertytype, order, md_rangedimension_propertytype.Name)
}

// Unstage removes md_rangedimension_propertytype off the model stage
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) Unstage(stage *Stage) *MD_RangeDimension_PropertyType {
	__gong__unstage(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyTypes_mapString, md_rangedimension_propertytype, md_rangedimension_propertytype.Name)
	return md_rangedimension_propertytype
}

// UnstageVoid removes md_rangedimension_propertytype off the model stage
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) UnstageVoid(stage *Stage) {
	md_rangedimension_propertytype.Unstage(stage)
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) StageVoid(stage *Stage) {
	md_rangedimension_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GetName() (res string) {
	return md_rangedimension_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) SetName(name string) {
	md_rangedimension_propertytype.Name = name
}

// Stage puts md_rangedimension_type to the model stage
func (md_rangedimension_type *MD_RangeDimension_Type) Stage(stage *Stage) *MD_RangeDimension_Type {
	__gong__stage(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Type_stagedOrder, stage.MD_RangeDimension_Type_orderStaged, &stage.MD_RangeDimension_TypeOrder, stage.MD_RangeDimension_Types_mapString, md_rangedimension_type, md_rangedimension_type.Name)
	return md_rangedimension_type
}

// StagePreserveOrder puts md_rangedimension_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_RangeDimension_TypeOrder
// - update stage.MD_RangeDimension_TypeOrder accordingly
func (md_rangedimension_type *MD_RangeDimension_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Type_stagedOrder, stage.MD_RangeDimension_Type_orderStaged, &stage.MD_RangeDimension_TypeOrder, stage.MD_RangeDimension_Types_mapString, md_rangedimension_type, order, md_rangedimension_type.Name)
}

// Unstage removes md_rangedimension_type off the model stage
func (md_rangedimension_type *MD_RangeDimension_Type) Unstage(stage *Stage) *MD_RangeDimension_Type {
	__gong__unstage(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Types_mapString, md_rangedimension_type, md_rangedimension_type.Name)
	return md_rangedimension_type
}

// UnstageVoid removes md_rangedimension_type off the model stage
func (md_rangedimension_type *MD_RangeDimension_Type) UnstageVoid(stage *Stage) {
	md_rangedimension_type.Unstage(stage)
}

func (md_rangedimension_type *MD_RangeDimension_Type) StageVoid(stage *Stage) {
	md_rangedimension_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_rangedimension_type *MD_RangeDimension_Type) GetName() (res string) {
	return md_rangedimension_type.Name
}

// for satisfaction of GongStruct interface
func (md_rangedimension_type *MD_RangeDimension_Type) SetName(name string) {
	md_rangedimension_type.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractMD_ContentInformation_Types, &stage.AbstractMD_ContentInformation_Types_mapString, &stage.AbstractMD_ContentInformation_Type_stagedOrder, &stage.AbstractMD_ContentInformation_TypeOrder)

	__gong__resetStageType(&stage.MD_Band_PropertyTypes, &stage.MD_Band_PropertyTypes_mapString, &stage.MD_Band_PropertyType_stagedOrder, &stage.MD_Band_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Band_Types, &stage.MD_Band_Types_mapString, &stage.MD_Band_Type_stagedOrder, &stage.MD_Band_TypeOrder)

	__gong__resetStageType(&stage.MD_ContentInformation_PropertyTypes, &stage.MD_ContentInformation_PropertyTypes_mapString, &stage.MD_ContentInformation_PropertyType_stagedOrder, &stage.MD_ContentInformation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_CoverageContentTypeCode_PropertyTypes, &stage.MD_CoverageContentTypeCode_PropertyTypes_mapString, &stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder, &stage.MD_CoverageContentTypeCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_CoverageDescription_PropertyTypes, &stage.MD_CoverageDescription_PropertyTypes_mapString, &stage.MD_CoverageDescription_PropertyType_stagedOrder, &stage.MD_CoverageDescription_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_CoverageDescription_Types, &stage.MD_CoverageDescription_Types_mapString, &stage.MD_CoverageDescription_Type_stagedOrder, &stage.MD_CoverageDescription_TypeOrder)

	__gong__resetStageType(&stage.MD_FeatureCatalogueDescription_PropertyTypes, &stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString, &stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder, &stage.MD_FeatureCatalogueDescription_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_FeatureCatalogueDescription_Types, &stage.MD_FeatureCatalogueDescription_Types_mapString, &stage.MD_FeatureCatalogueDescription_Type_stagedOrder, &stage.MD_FeatureCatalogueDescription_TypeOrder)

	__gong__resetStageType(&stage.MD_ImageDescription_PropertyTypes, &stage.MD_ImageDescription_PropertyTypes_mapString, &stage.MD_ImageDescription_PropertyType_stagedOrder, &stage.MD_ImageDescription_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_ImageDescription_Types, &stage.MD_ImageDescription_Types_mapString, &stage.MD_ImageDescription_Type_stagedOrder, &stage.MD_ImageDescription_TypeOrder)

	__gong__resetStageType(&stage.MD_ImagingConditionCode_PropertyTypes, &stage.MD_ImagingConditionCode_PropertyTypes_mapString, &stage.MD_ImagingConditionCode_PropertyType_stagedOrder, &stage.MD_ImagingConditionCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_RangeDimension_PropertyTypes, &stage.MD_RangeDimension_PropertyTypes_mapString, &stage.MD_RangeDimension_PropertyType_stagedOrder, &stage.MD_RangeDimension_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_RangeDimension_Types, &stage.MD_RangeDimension_Types_mapString, &stage.MD_RangeDimension_Type_stagedOrder, &stage.MD_RangeDimension_TypeOrder)

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
func (AbstractMD_ContentInformation_Type) GongGetAssociationName() any {
	return &AbstractMD_ContentInformation_Type{
	}
}

func (MD_Band_PropertyType) GongGetAssociationName() any {
	return &MD_Band_PropertyType{
			MD_Band: &MD_Band_Type{Name: "MD_Band"},
	}
}

func (MD_Band_Type) GongGetAssociationName() any {
	return &MD_Band_Type{
	}
}

func (MD_ContentInformation_PropertyType) GongGetAssociationName() any {
	return &MD_ContentInformation_PropertyType{
			MD_FeatureCatalogueDescription: &MD_FeatureCatalogueDescription_Type{Name: "MD_FeatureCatalogueDescription"},
			MD_CoverageDescription: &MD_CoverageDescription_Type{Name: "MD_CoverageDescription"},
			MD_ImageDescription: &MD_ImageDescription_Type{Name: "MD_ImageDescription"},
	}
}

func (MD_CoverageContentTypeCode_PropertyType) GongGetAssociationName() any {
	return &MD_CoverageContentTypeCode_PropertyType{
	}
}

func (MD_CoverageDescription_PropertyType) GongGetAssociationName() any {
	return &MD_CoverageDescription_PropertyType{
			MD_CoverageDescription: &MD_CoverageDescription_Type{Name: "MD_CoverageDescription"},
			MD_ImageDescription: &MD_ImageDescription_Type{Name: "MD_ImageDescription"},
	}
}

func (MD_CoverageDescription_Type) GongGetAssociationName() any {
	return &MD_CoverageDescription_Type{
			ContentType: &MD_CoverageContentTypeCode_PropertyType{Name: "ContentType"},
			Dimension: []*MD_RangeDimension_PropertyType{{Name: "Dimension"}},
	}
}

func (MD_FeatureCatalogueDescription_PropertyType) GongGetAssociationName() any {
	return &MD_FeatureCatalogueDescription_PropertyType{
			MD_FeatureCatalogueDescription: &MD_FeatureCatalogueDescription_Type{Name: "MD_FeatureCatalogueDescription"},
	}
}

func (MD_FeatureCatalogueDescription_Type) GongGetAssociationName() any {
	return &MD_FeatureCatalogueDescription_Type{
	}
}

func (MD_ImageDescription_PropertyType) GongGetAssociationName() any {
	return &MD_ImageDescription_PropertyType{
			MD_ImageDescription: &MD_ImageDescription_Type{Name: "MD_ImageDescription"},
	}
}

func (MD_ImageDescription_Type) GongGetAssociationName() any {
	return &MD_ImageDescription_Type{
			ContentType: &MD_CoverageContentTypeCode_PropertyType{Name: "ContentType"},
			Dimension: []*MD_RangeDimension_PropertyType{{Name: "Dimension"}},
			ImagingCondition: &MD_ImagingConditionCode_PropertyType{Name: "ImagingCondition"},
	}
}

func (MD_ImagingConditionCode_PropertyType) GongGetAssociationName() any {
	return &MD_ImagingConditionCode_PropertyType{
	}
}

func (MD_RangeDimension_PropertyType) GongGetAssociationName() any {
	return &MD_RangeDimension_PropertyType{
			MD_RangeDimension: &MD_RangeDimension_Type{Name: "MD_RangeDimension"},
			MD_Band: &MD_Band_Type{Name: "MD_Band"},
	}
}

func (MD_RangeDimension_Type) GongGetAssociationName() any {
	return &MD_RangeDimension_Type{
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
	// reverse maps of direct associations of AbstractMD_ContentInformation_Type
	case AbstractMD_ContentInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Band_PropertyType
	case MD_Band_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_Band":
			res := make(map[*MD_Band_Type][]*MD_Band_PropertyType)
			for md_band_propertytype := range stage.MD_Band_PropertyTypes {
				if md_band_propertytype.MD_Band != nil {
					md_band_type_ := md_band_propertytype.MD_Band
					var md_band_propertytypes []*MD_Band_PropertyType
					_, ok := res[md_band_type_]
					if ok {
						md_band_propertytypes = res[md_band_type_]
					} else {
						md_band_propertytypes = make([]*MD_Band_PropertyType, 0)
					}
					md_band_propertytypes = append(md_band_propertytypes, md_band_propertytype)
					res[md_band_type_] = md_band_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Band_Type
	case MD_Band_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ContentInformation_PropertyType
	case MD_ContentInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_FeatureCatalogueDescription":
			res := make(map[*MD_FeatureCatalogueDescription_Type][]*MD_ContentInformation_PropertyType)
			for md_contentinformation_propertytype := range stage.MD_ContentInformation_PropertyTypes {
				if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
					md_featurecataloguedescription_type_ := md_contentinformation_propertytype.MD_FeatureCatalogueDescription
					var md_contentinformation_propertytypes []*MD_ContentInformation_PropertyType
					_, ok := res[md_featurecataloguedescription_type_]
					if ok {
						md_contentinformation_propertytypes = res[md_featurecataloguedescription_type_]
					} else {
						md_contentinformation_propertytypes = make([]*MD_ContentInformation_PropertyType, 0)
					}
					md_contentinformation_propertytypes = append(md_contentinformation_propertytypes, md_contentinformation_propertytype)
					res[md_featurecataloguedescription_type_] = md_contentinformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_CoverageDescription":
			res := make(map[*MD_CoverageDescription_Type][]*MD_ContentInformation_PropertyType)
			for md_contentinformation_propertytype := range stage.MD_ContentInformation_PropertyTypes {
				if md_contentinformation_propertytype.MD_CoverageDescription != nil {
					md_coveragedescription_type_ := md_contentinformation_propertytype.MD_CoverageDescription
					var md_contentinformation_propertytypes []*MD_ContentInformation_PropertyType
					_, ok := res[md_coveragedescription_type_]
					if ok {
						md_contentinformation_propertytypes = res[md_coveragedescription_type_]
					} else {
						md_contentinformation_propertytypes = make([]*MD_ContentInformation_PropertyType, 0)
					}
					md_contentinformation_propertytypes = append(md_contentinformation_propertytypes, md_contentinformation_propertytype)
					res[md_coveragedescription_type_] = md_contentinformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_ImageDescription":
			res := make(map[*MD_ImageDescription_Type][]*MD_ContentInformation_PropertyType)
			for md_contentinformation_propertytype := range stage.MD_ContentInformation_PropertyTypes {
				if md_contentinformation_propertytype.MD_ImageDescription != nil {
					md_imagedescription_type_ := md_contentinformation_propertytype.MD_ImageDescription
					var md_contentinformation_propertytypes []*MD_ContentInformation_PropertyType
					_, ok := res[md_imagedescription_type_]
					if ok {
						md_contentinformation_propertytypes = res[md_imagedescription_type_]
					} else {
						md_contentinformation_propertytypes = make([]*MD_ContentInformation_PropertyType, 0)
					}
					md_contentinformation_propertytypes = append(md_contentinformation_propertytypes, md_contentinformation_propertytype)
					res[md_imagedescription_type_] = md_contentinformation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_CoverageContentTypeCode_PropertyType
	case MD_CoverageContentTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CoverageDescription_PropertyType
	case MD_CoverageDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_CoverageDescription":
			res := make(map[*MD_CoverageDescription_Type][]*MD_CoverageDescription_PropertyType)
			for md_coveragedescription_propertytype := range stage.MD_CoverageDescription_PropertyTypes {
				if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
					md_coveragedescription_type_ := md_coveragedescription_propertytype.MD_CoverageDescription
					var md_coveragedescription_propertytypes []*MD_CoverageDescription_PropertyType
					_, ok := res[md_coveragedescription_type_]
					if ok {
						md_coveragedescription_propertytypes = res[md_coveragedescription_type_]
					} else {
						md_coveragedescription_propertytypes = make([]*MD_CoverageDescription_PropertyType, 0)
					}
					md_coveragedescription_propertytypes = append(md_coveragedescription_propertytypes, md_coveragedescription_propertytype)
					res[md_coveragedescription_type_] = md_coveragedescription_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_ImageDescription":
			res := make(map[*MD_ImageDescription_Type][]*MD_CoverageDescription_PropertyType)
			for md_coveragedescription_propertytype := range stage.MD_CoverageDescription_PropertyTypes {
				if md_coveragedescription_propertytype.MD_ImageDescription != nil {
					md_imagedescription_type_ := md_coveragedescription_propertytype.MD_ImageDescription
					var md_coveragedescription_propertytypes []*MD_CoverageDescription_PropertyType
					_, ok := res[md_imagedescription_type_]
					if ok {
						md_coveragedescription_propertytypes = res[md_imagedescription_type_]
					} else {
						md_coveragedescription_propertytypes = make([]*MD_CoverageDescription_PropertyType, 0)
					}
					md_coveragedescription_propertytypes = append(md_coveragedescription_propertytypes, md_coveragedescription_propertytype)
					res[md_imagedescription_type_] = md_coveragedescription_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_CoverageDescription_Type
	case MD_CoverageDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ContentType":
			res := make(map[*MD_CoverageContentTypeCode_PropertyType][]*MD_CoverageDescription_Type)
			for md_coveragedescription_type := range stage.MD_CoverageDescription_Types {
				if md_coveragedescription_type.ContentType != nil {
					md_coveragecontenttypecode_propertytype_ := md_coveragedescription_type.ContentType
					var md_coveragedescription_types []*MD_CoverageDescription_Type
					_, ok := res[md_coveragecontenttypecode_propertytype_]
					if ok {
						md_coveragedescription_types = res[md_coveragecontenttypecode_propertytype_]
					} else {
						md_coveragedescription_types = make([]*MD_CoverageDescription_Type, 0)
					}
					md_coveragedescription_types = append(md_coveragedescription_types, md_coveragedescription_type)
					res[md_coveragecontenttypecode_propertytype_] = md_coveragedescription_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_FeatureCatalogueDescription_PropertyType
	case MD_FeatureCatalogueDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_FeatureCatalogueDescription":
			res := make(map[*MD_FeatureCatalogueDescription_Type][]*MD_FeatureCatalogueDescription_PropertyType)
			for md_featurecataloguedescription_propertytype := range stage.MD_FeatureCatalogueDescription_PropertyTypes {
				if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
					md_featurecataloguedescription_type_ := md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription
					var md_featurecataloguedescription_propertytypes []*MD_FeatureCatalogueDescription_PropertyType
					_, ok := res[md_featurecataloguedescription_type_]
					if ok {
						md_featurecataloguedescription_propertytypes = res[md_featurecataloguedescription_type_]
					} else {
						md_featurecataloguedescription_propertytypes = make([]*MD_FeatureCatalogueDescription_PropertyType, 0)
					}
					md_featurecataloguedescription_propertytypes = append(md_featurecataloguedescription_propertytypes, md_featurecataloguedescription_propertytype)
					res[md_featurecataloguedescription_type_] = md_featurecataloguedescription_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_FeatureCatalogueDescription_Type
	case MD_FeatureCatalogueDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ImageDescription_PropertyType
	case MD_ImageDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_ImageDescription":
			res := make(map[*MD_ImageDescription_Type][]*MD_ImageDescription_PropertyType)
			for md_imagedescription_propertytype := range stage.MD_ImageDescription_PropertyTypes {
				if md_imagedescription_propertytype.MD_ImageDescription != nil {
					md_imagedescription_type_ := md_imagedescription_propertytype.MD_ImageDescription
					var md_imagedescription_propertytypes []*MD_ImageDescription_PropertyType
					_, ok := res[md_imagedescription_type_]
					if ok {
						md_imagedescription_propertytypes = res[md_imagedescription_type_]
					} else {
						md_imagedescription_propertytypes = make([]*MD_ImageDescription_PropertyType, 0)
					}
					md_imagedescription_propertytypes = append(md_imagedescription_propertytypes, md_imagedescription_propertytype)
					res[md_imagedescription_type_] = md_imagedescription_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ImageDescription_Type
	case MD_ImageDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ContentType":
			res := make(map[*MD_CoverageContentTypeCode_PropertyType][]*MD_ImageDescription_Type)
			for md_imagedescription_type := range stage.MD_ImageDescription_Types {
				if md_imagedescription_type.ContentType != nil {
					md_coveragecontenttypecode_propertytype_ := md_imagedescription_type.ContentType
					var md_imagedescription_types []*MD_ImageDescription_Type
					_, ok := res[md_coveragecontenttypecode_propertytype_]
					if ok {
						md_imagedescription_types = res[md_coveragecontenttypecode_propertytype_]
					} else {
						md_imagedescription_types = make([]*MD_ImageDescription_Type, 0)
					}
					md_imagedescription_types = append(md_imagedescription_types, md_imagedescription_type)
					res[md_coveragecontenttypecode_propertytype_] = md_imagedescription_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "ImagingCondition":
			res := make(map[*MD_ImagingConditionCode_PropertyType][]*MD_ImageDescription_Type)
			for md_imagedescription_type := range stage.MD_ImageDescription_Types {
				if md_imagedescription_type.ImagingCondition != nil {
					md_imagingconditioncode_propertytype_ := md_imagedescription_type.ImagingCondition
					var md_imagedescription_types []*MD_ImageDescription_Type
					_, ok := res[md_imagingconditioncode_propertytype_]
					if ok {
						md_imagedescription_types = res[md_imagingconditioncode_propertytype_]
					} else {
						md_imagedescription_types = make([]*MD_ImageDescription_Type, 0)
					}
					md_imagedescription_types = append(md_imagedescription_types, md_imagedescription_type)
					res[md_imagingconditioncode_propertytype_] = md_imagedescription_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ImagingConditionCode_PropertyType
	case MD_ImagingConditionCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_RangeDimension_PropertyType
	case MD_RangeDimension_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_RangeDimension":
			res := make(map[*MD_RangeDimension_Type][]*MD_RangeDimension_PropertyType)
			for md_rangedimension_propertytype := range stage.MD_RangeDimension_PropertyTypes {
				if md_rangedimension_propertytype.MD_RangeDimension != nil {
					md_rangedimension_type_ := md_rangedimension_propertytype.MD_RangeDimension
					var md_rangedimension_propertytypes []*MD_RangeDimension_PropertyType
					_, ok := res[md_rangedimension_type_]
					if ok {
						md_rangedimension_propertytypes = res[md_rangedimension_type_]
					} else {
						md_rangedimension_propertytypes = make([]*MD_RangeDimension_PropertyType, 0)
					}
					md_rangedimension_propertytypes = append(md_rangedimension_propertytypes, md_rangedimension_propertytype)
					res[md_rangedimension_type_] = md_rangedimension_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_Band":
			res := make(map[*MD_Band_Type][]*MD_RangeDimension_PropertyType)
			for md_rangedimension_propertytype := range stage.MD_RangeDimension_PropertyTypes {
				if md_rangedimension_propertytype.MD_Band != nil {
					md_band_type_ := md_rangedimension_propertytype.MD_Band
					var md_rangedimension_propertytypes []*MD_RangeDimension_PropertyType
					_, ok := res[md_band_type_]
					if ok {
						md_rangedimension_propertytypes = res[md_band_type_]
					} else {
						md_rangedimension_propertytypes = make([]*MD_RangeDimension_PropertyType, 0)
					}
					md_rangedimension_propertytypes = append(md_rangedimension_propertytypes, md_rangedimension_propertytype)
					res[md_band_type_] = md_rangedimension_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_RangeDimension_Type
	case MD_RangeDimension_Type:
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
	// reverse maps of direct associations of AbstractMD_ContentInformation_Type
	case AbstractMD_ContentInformation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Band_PropertyType
	case MD_Band_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Band_Type
	case MD_Band_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ContentInformation_PropertyType
	case MD_ContentInformation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CoverageContentTypeCode_PropertyType
	case MD_CoverageContentTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CoverageDescription_PropertyType
	case MD_CoverageDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CoverageDescription_Type
	case MD_CoverageDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "Dimension":
			res := make(map[*MD_RangeDimension_PropertyType][]*MD_CoverageDescription_Type)
			for md_coveragedescription_type := range stage.MD_CoverageDescription_Types {
				for _, md_rangedimension_propertytype_ := range md_coveragedescription_type.Dimension {
					res[md_rangedimension_propertytype_] = append(res[md_rangedimension_propertytype_], md_coveragedescription_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_FeatureCatalogueDescription_PropertyType
	case MD_FeatureCatalogueDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_FeatureCatalogueDescription_Type
	case MD_FeatureCatalogueDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ImageDescription_PropertyType
	case MD_ImageDescription_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_ImageDescription_Type
	case MD_ImageDescription_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "Dimension":
			res := make(map[*MD_RangeDimension_PropertyType][]*MD_ImageDescription_Type)
			for md_imagedescription_type := range stage.MD_ImageDescription_Types {
				for _, md_rangedimension_propertytype_ := range md_imagedescription_type.Dimension {
					res[md_rangedimension_propertytype_] = append(res[md_rangedimension_propertytype_], md_imagedescription_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_ImagingConditionCode_PropertyType
	case MD_ImagingConditionCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_RangeDimension_PropertyType
	case MD_RangeDimension_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_RangeDimension_Type
	case MD_RangeDimension_Type:
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
func (*AbstractMD_ContentInformation_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Band_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Band_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ContentInformation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_CoverageContentTypeCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_CoverageDescription_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_CoverageDescription_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_FeatureCatalogueDescription_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_FeatureCatalogueDescription_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ImageDescription_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ImageDescription_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_ImagingConditionCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_RangeDimension_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MD_CoverageDescription_Type",
			Fieldname: "Dimension",
		},
		{
			GongstructName: "MD_ImageDescription_Type",
			Fieldname: "Dimension",
		},
	}
}

func (*MD_RangeDimension_Type) GongGetReverseFields() []GongReverseField {
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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_band_propertytype *MD_Band_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_Band",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Band_Type",
		},
	}
	return
}

func (md_band_type *MD_Band_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_FeatureCatalogueDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_FeatureCatalogueDescription_Type",
		},
		{
			Name:                 "MD_CoverageDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CoverageDescription_Type",
		},
		{
			Name:                 "MD_ImageDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ImageDescription_Type",
		},
	}
	return
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_CoverageDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CoverageDescription_Type",
		},
		{
			Name:                 "MD_ImageDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ImageDescription_Type",
		},
	}
	return
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ContentType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CoverageContentTypeCode_PropertyType",
		},
		{
			Name:                 "Dimension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_RangeDimension_PropertyType",
		},
	}
	return
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_FeatureCatalogueDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_FeatureCatalogueDescription_Type",
		},
	}
	return
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_ImageDescription",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ImageDescription_Type",
		},
	}
	return
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ContentType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CoverageContentTypeCode_PropertyType",
		},
		{
			Name:                 "Dimension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_RangeDimension_PropertyType",
		},
		{
			Name:                 "ImagingCondition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_ImagingConditionCode_PropertyType",
		},
	}
	return
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_RangeDimension",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_RangeDimension_Type",
		},
		{
			Name:                 "MD_Band",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Band_Type",
		},
	}
	return
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractmd_contentinformation_type.Name
	case "Id":
		res.valueString = abstractmd_contentinformation_type.Id
	case "Uuid":
		res.valueString = abstractmd_contentinformation_type.Uuid
	}
	return
}

func (md_band_propertytype *MD_Band_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_band_propertytype.Name
	case "NilReason":
		res.valueString = md_band_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_band_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_band_propertytype.Type_XSD
	case "Href":
		res.valueString = md_band_propertytype.Href
	case "Role":
		res.valueString = md_band_propertytype.Role
	case "Arcrole":
		res.valueString = md_band_propertytype.Arcrole
	case "Title":
		res.valueString = md_band_propertytype.Title
	case "Show":
		res.valueString = md_band_propertytype.Show
	case "Actuate":
		res.valueString = md_band_propertytype.Actuate
	case "MD_Band":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_band_propertytype.MD_Band != nil {
			res.valueString = md_band_propertytype.MD_Band.Name
			res.ids = md_band_propertytype.MD_Band.GongGetUUID(stage)
		}
	}
	return
}

func (md_band_type *MD_Band_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_band_type.Name
	case "Id":
		res.valueString = md_band_type.Id
	case "Uuid":
		res.valueString = md_band_type.Uuid
	}
	return
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_contentinformation_propertytype.Name
	case "NilReason":
		res.valueString = md_contentinformation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_contentinformation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_contentinformation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_contentinformation_propertytype.Href
	case "Role":
		res.valueString = md_contentinformation_propertytype.Role
	case "Arcrole":
		res.valueString = md_contentinformation_propertytype.Arcrole
	case "Title":
		res.valueString = md_contentinformation_propertytype.Title
	case "Show":
		res.valueString = md_contentinformation_propertytype.Show
	case "Actuate":
		res.valueString = md_contentinformation_propertytype.Actuate
	case "MD_FeatureCatalogueDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
			res.valueString = md_contentinformation_propertytype.MD_FeatureCatalogueDescription.Name
			res.ids = md_contentinformation_propertytype.MD_FeatureCatalogueDescription.GongGetUUID(stage)
		}
	case "MD_CoverageDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_contentinformation_propertytype.MD_CoverageDescription != nil {
			res.valueString = md_contentinformation_propertytype.MD_CoverageDescription.Name
			res.ids = md_contentinformation_propertytype.MD_CoverageDescription.GongGetUUID(stage)
		}
	case "MD_ImageDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_contentinformation_propertytype.MD_ImageDescription != nil {
			res.valueString = md_contentinformation_propertytype.MD_ImageDescription.Name
			res.ids = md_contentinformation_propertytype.MD_ImageDescription.GongGetUUID(stage)
		}
	}
	return
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_coveragecontenttypecode_propertytype.Name
	case "NilReason":
		res.valueString = md_coveragecontenttypecode_propertytype.NilReason
	}
	return
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_coveragedescription_propertytype.Name
	case "NilReason":
		res.valueString = md_coveragedescription_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_coveragedescription_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_coveragedescription_propertytype.Type_XSD
	case "Href":
		res.valueString = md_coveragedescription_propertytype.Href
	case "Role":
		res.valueString = md_coveragedescription_propertytype.Role
	case "Arcrole":
		res.valueString = md_coveragedescription_propertytype.Arcrole
	case "Title":
		res.valueString = md_coveragedescription_propertytype.Title
	case "Show":
		res.valueString = md_coveragedescription_propertytype.Show
	case "Actuate":
		res.valueString = md_coveragedescription_propertytype.Actuate
	case "MD_CoverageDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
			res.valueString = md_coveragedescription_propertytype.MD_CoverageDescription.Name
			res.ids = md_coveragedescription_propertytype.MD_CoverageDescription.GongGetUUID(stage)
		}
	case "MD_ImageDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_coveragedescription_propertytype.MD_ImageDescription != nil {
			res.valueString = md_coveragedescription_propertytype.MD_ImageDescription.Name
			res.ids = md_coveragedescription_propertytype.MD_ImageDescription.GongGetUUID(stage)
		}
	}
	return
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_coveragedescription_type.Name
	case "Id":
		res.valueString = md_coveragedescription_type.Id
	case "Uuid":
		res.valueString = md_coveragedescription_type.Uuid
	case "ContentType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_coveragedescription_type.ContentType != nil {
			res.valueString = md_coveragedescription_type.ContentType.Name
			res.ids = md_coveragedescription_type.ContentType.GongGetUUID(stage)
		}
	case "Dimension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_coveragedescription_type.Dimension {
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

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_featurecataloguedescription_propertytype.Name
	case "NilReason":
		res.valueString = md_featurecataloguedescription_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_featurecataloguedescription_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_featurecataloguedescription_propertytype.Type_XSD
	case "Href":
		res.valueString = md_featurecataloguedescription_propertytype.Href
	case "Role":
		res.valueString = md_featurecataloguedescription_propertytype.Role
	case "Arcrole":
		res.valueString = md_featurecataloguedescription_propertytype.Arcrole
	case "Title":
		res.valueString = md_featurecataloguedescription_propertytype.Title
	case "Show":
		res.valueString = md_featurecataloguedescription_propertytype.Show
	case "Actuate":
		res.valueString = md_featurecataloguedescription_propertytype.Actuate
	case "MD_FeatureCatalogueDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
			res.valueString = md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription.Name
			res.ids = md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription.GongGetUUID(stage)
		}
	}
	return
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_featurecataloguedescription_type.Name
	case "Id":
		res.valueString = md_featurecataloguedescription_type.Id
	case "Uuid":
		res.valueString = md_featurecataloguedescription_type.Uuid
	}
	return
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_imagedescription_propertytype.Name
	case "NilReason":
		res.valueString = md_imagedescription_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_imagedescription_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_imagedescription_propertytype.Type_XSD
	case "Href":
		res.valueString = md_imagedescription_propertytype.Href
	case "Role":
		res.valueString = md_imagedescription_propertytype.Role
	case "Arcrole":
		res.valueString = md_imagedescription_propertytype.Arcrole
	case "Title":
		res.valueString = md_imagedescription_propertytype.Title
	case "Show":
		res.valueString = md_imagedescription_propertytype.Show
	case "Actuate":
		res.valueString = md_imagedescription_propertytype.Actuate
	case "MD_ImageDescription":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_imagedescription_propertytype.MD_ImageDescription != nil {
			res.valueString = md_imagedescription_propertytype.MD_ImageDescription.Name
			res.ids = md_imagedescription_propertytype.MD_ImageDescription.GongGetUUID(stage)
		}
	}
	return
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_imagedescription_type.Name
	case "Id":
		res.valueString = md_imagedescription_type.Id
	case "Uuid":
		res.valueString = md_imagedescription_type.Uuid
	case "ContentType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_imagedescription_type.ContentType != nil {
			res.valueString = md_imagedescription_type.ContentType.Name
			res.ids = md_imagedescription_type.ContentType.GongGetUUID(stage)
		}
	case "Dimension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_imagedescription_type.Dimension {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ImagingCondition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_imagedescription_type.ImagingCondition != nil {
			res.valueString = md_imagedescription_type.ImagingCondition.Name
			res.ids = md_imagedescription_type.ImagingCondition.GongGetUUID(stage)
		}
	}
	return
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_imagingconditioncode_propertytype.Name
	case "NilReason":
		res.valueString = md_imagingconditioncode_propertytype.NilReason
	}
	return
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_rangedimension_propertytype.Name
	case "NilReason":
		res.valueString = md_rangedimension_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_rangedimension_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_rangedimension_propertytype.Type_XSD
	case "Href":
		res.valueString = md_rangedimension_propertytype.Href
	case "Role":
		res.valueString = md_rangedimension_propertytype.Role
	case "Arcrole":
		res.valueString = md_rangedimension_propertytype.Arcrole
	case "Title":
		res.valueString = md_rangedimension_propertytype.Title
	case "Show":
		res.valueString = md_rangedimension_propertytype.Show
	case "Actuate":
		res.valueString = md_rangedimension_propertytype.Actuate
	case "MD_RangeDimension":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_rangedimension_propertytype.MD_RangeDimension != nil {
			res.valueString = md_rangedimension_propertytype.MD_RangeDimension.Name
			res.ids = md_rangedimension_propertytype.MD_RangeDimension.GongGetUUID(stage)
		}
	case "MD_Band":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_rangedimension_propertytype.MD_Band != nil {
			res.valueString = md_rangedimension_propertytype.MD_Band.Name
			res.ids = md_rangedimension_propertytype.MD_Band.GongGetUUID(stage)
		}
	}
	return
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_rangedimension_type.Name
	case "Id":
		res.valueString = md_rangedimension_type.Id
	case "Uuid":
		res.valueString = md_rangedimension_type.Uuid
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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetGongstructName() string {
	return "AbstractMD_ContentInformation_Type"
}

func (md_band_propertytype *MD_Band_PropertyType) GongGetGongstructName() string {
	return "MD_Band_PropertyType"
}

func (md_band_type *MD_Band_Type) GongGetGongstructName() string {
	return "MD_Band_Type"
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetGongstructName() string {
	return "MD_ContentInformation_PropertyType"
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetGongstructName() string {
	return "MD_CoverageContentTypeCode_PropertyType"
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetGongstructName() string {
	return "MD_CoverageDescription_PropertyType"
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetGongstructName() string {
	return "MD_CoverageDescription_Type"
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetGongstructName() string {
	return "MD_FeatureCatalogueDescription_PropertyType"
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetGongstructName() string {
	return "MD_FeatureCatalogueDescription_Type"
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetGongstructName() string {
	return "MD_ImageDescription_PropertyType"
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetGongstructName() string {
	return "MD_ImageDescription_Type"
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetGongstructName() string {
	return "MD_ImagingConditionCode_PropertyType"
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetGongstructName() string {
	return "MD_RangeDimension_PropertyType"
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetGongstructName() string {
	return "MD_RangeDimension_Type"
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
	__gong__rebuildMapString(stage.AbstractMD_ContentInformation_Types, &stage.AbstractMD_ContentInformation_Types_mapString)

	__gong__rebuildMapString(stage.MD_Band_PropertyTypes, &stage.MD_Band_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Band_Types, &stage.MD_Band_Types_mapString)

	__gong__rebuildMapString(stage.MD_ContentInformation_PropertyTypes, &stage.MD_ContentInformation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_CoverageContentTypeCode_PropertyTypes, &stage.MD_CoverageContentTypeCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_CoverageDescription_PropertyTypes, &stage.MD_CoverageDescription_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_CoverageDescription_Types, &stage.MD_CoverageDescription_Types_mapString)

	__gong__rebuildMapString(stage.MD_FeatureCatalogueDescription_PropertyTypes, &stage.MD_FeatureCatalogueDescription_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_FeatureCatalogueDescription_Types, &stage.MD_FeatureCatalogueDescription_Types_mapString)

	__gong__rebuildMapString(stage.MD_ImageDescription_PropertyTypes, &stage.MD_ImageDescription_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_ImageDescription_Types, &stage.MD_ImageDescription_Types_mapString)

	__gong__rebuildMapString(stage.MD_ImagingConditionCode_PropertyTypes, &stage.MD_ImagingConditionCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_RangeDimension_PropertyTypes, &stage.MD_RangeDimension_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_RangeDimension_Types, &stage.MD_RangeDimension_Types_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
