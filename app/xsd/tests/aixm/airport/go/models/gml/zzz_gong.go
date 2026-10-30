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
	AbstractCurveSegmentTypes                map[*AbstractCurveSegmentType]struct{}
	AbstractCurveSegmentTypes_instance       map[*AbstractCurveSegmentType]*AbstractCurveSegmentType
	AbstractCurveSegmentTypes_mapString      map[string]*AbstractCurveSegmentType
	AbstractCurveSegmentTypeOrder            uint
	AbstractCurveSegmentType_stagedOrder     map[*AbstractCurveSegmentType]uint
	AbstractCurveSegmentType_orderStaged     map[uint]*AbstractCurveSegmentType
	AbstractCurveSegmentTypes_reference      map[*AbstractCurveSegmentType]*AbstractCurveSegmentType
	AbstractCurveSegmentTypes_referenceOrder map[*AbstractCurveSegmentType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractCurveSegmentTypeCreateCallback GongOnAfterCreateInterface[AbstractCurveSegmentType]
	OnAfterAbstractCurveSegmentTypeUpdateCallback GongOnAfterUpdateInterface[AbstractCurveSegmentType]
	OnAfterAbstractCurveSegmentTypeDeleteCallback GongOnAfterDeleteInterface[AbstractCurveSegmentType]

	AbstractCurveTypes                map[*AbstractCurveType]struct{}
	AbstractCurveTypes_instance       map[*AbstractCurveType]*AbstractCurveType
	AbstractCurveTypes_mapString      map[string]*AbstractCurveType
	AbstractCurveTypeOrder            uint
	AbstractCurveType_stagedOrder     map[*AbstractCurveType]uint
	AbstractCurveType_orderStaged     map[uint]*AbstractCurveType
	AbstractCurveTypes_reference      map[*AbstractCurveType]*AbstractCurveType
	AbstractCurveTypes_referenceOrder map[*AbstractCurveType]uint

	// insertion point for slice of pointers maps
	AbstractCurveType_NameXSD_reverseMap map[*CodeType]*AbstractCurveType

	OnAfterAbstractCurveTypeCreateCallback GongOnAfterCreateInterface[AbstractCurveType]
	OnAfterAbstractCurveTypeUpdateCallback GongOnAfterUpdateInterface[AbstractCurveType]
	OnAfterAbstractCurveTypeDeleteCallback GongOnAfterDeleteInterface[AbstractCurveType]

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

	AbstractRingPropertyTypes                map[*AbstractRingPropertyType]struct{}
	AbstractRingPropertyTypes_instance       map[*AbstractRingPropertyType]*AbstractRingPropertyType
	AbstractRingPropertyTypes_mapString      map[string]*AbstractRingPropertyType
	AbstractRingPropertyTypeOrder            uint
	AbstractRingPropertyType_stagedOrder     map[*AbstractRingPropertyType]uint
	AbstractRingPropertyType_orderStaged     map[uint]*AbstractRingPropertyType
	AbstractRingPropertyTypes_reference      map[*AbstractRingPropertyType]*AbstractRingPropertyType
	AbstractRingPropertyTypes_referenceOrder map[*AbstractRingPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractRingPropertyTypeCreateCallback GongOnAfterCreateInterface[AbstractRingPropertyType]
	OnAfterAbstractRingPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AbstractRingPropertyType]
	OnAfterAbstractRingPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AbstractRingPropertyType]

	AbstractRingTypes                map[*AbstractRingType]struct{}
	AbstractRingTypes_instance       map[*AbstractRingType]*AbstractRingType
	AbstractRingTypes_mapString      map[string]*AbstractRingType
	AbstractRingTypeOrder            uint
	AbstractRingType_stagedOrder     map[*AbstractRingType]uint
	AbstractRingType_orderStaged     map[uint]*AbstractRingType
	AbstractRingTypes_reference      map[*AbstractRingType]*AbstractRingType
	AbstractRingTypes_referenceOrder map[*AbstractRingType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractRingTypeCreateCallback GongOnAfterCreateInterface[AbstractRingType]
	OnAfterAbstractRingTypeUpdateCallback GongOnAfterUpdateInterface[AbstractRingType]
	OnAfterAbstractRingTypeDeleteCallback GongOnAfterDeleteInterface[AbstractRingType]

	AbstractSurfacePatchTypes                map[*AbstractSurfacePatchType]struct{}
	AbstractSurfacePatchTypes_instance       map[*AbstractSurfacePatchType]*AbstractSurfacePatchType
	AbstractSurfacePatchTypes_mapString      map[string]*AbstractSurfacePatchType
	AbstractSurfacePatchTypeOrder            uint
	AbstractSurfacePatchType_stagedOrder     map[*AbstractSurfacePatchType]uint
	AbstractSurfacePatchType_orderStaged     map[uint]*AbstractSurfacePatchType
	AbstractSurfacePatchTypes_reference      map[*AbstractSurfacePatchType]*AbstractSurfacePatchType
	AbstractSurfacePatchTypes_referenceOrder map[*AbstractSurfacePatchType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractSurfacePatchTypeCreateCallback GongOnAfterCreateInterface[AbstractSurfacePatchType]
	OnAfterAbstractSurfacePatchTypeUpdateCallback GongOnAfterUpdateInterface[AbstractSurfacePatchType]
	OnAfterAbstractSurfacePatchTypeDeleteCallback GongOnAfterDeleteInterface[AbstractSurfacePatchType]

	AbstractSurfaceTypes                map[*AbstractSurfaceType]struct{}
	AbstractSurfaceTypes_instance       map[*AbstractSurfaceType]*AbstractSurfaceType
	AbstractSurfaceTypes_mapString      map[string]*AbstractSurfaceType
	AbstractSurfaceTypeOrder            uint
	AbstractSurfaceType_stagedOrder     map[*AbstractSurfaceType]uint
	AbstractSurfaceType_orderStaged     map[uint]*AbstractSurfaceType
	AbstractSurfaceTypes_reference      map[*AbstractSurfaceType]*AbstractSurfaceType
	AbstractSurfaceTypes_referenceOrder map[*AbstractSurfaceType]uint

	// insertion point for slice of pointers maps
	AbstractSurfaceType_NameXSD_reverseMap map[*CodeType]*AbstractSurfaceType

	OnAfterAbstractSurfaceTypeCreateCallback GongOnAfterCreateInterface[AbstractSurfaceType]
	OnAfterAbstractSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[AbstractSurfaceType]
	OnAfterAbstractSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[AbstractSurfaceType]

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

	ArcByCenterPointTypes                map[*ArcByCenterPointType]struct{}
	ArcByCenterPointTypes_instance       map[*ArcByCenterPointType]*ArcByCenterPointType
	ArcByCenterPointTypes_mapString      map[string]*ArcByCenterPointType
	ArcByCenterPointTypeOrder            uint
	ArcByCenterPointType_stagedOrder     map[*ArcByCenterPointType]uint
	ArcByCenterPointType_orderStaged     map[uint]*ArcByCenterPointType
	ArcByCenterPointTypes_reference      map[*ArcByCenterPointType]*ArcByCenterPointType
	ArcByCenterPointTypes_referenceOrder map[*ArcByCenterPointType]uint

	// insertion point for slice of pointers maps
	OnAfterArcByCenterPointTypeCreateCallback GongOnAfterCreateInterface[ArcByCenterPointType]
	OnAfterArcByCenterPointTypeUpdateCallback GongOnAfterUpdateInterface[ArcByCenterPointType]
	OnAfterArcByCenterPointTypeDeleteCallback GongOnAfterDeleteInterface[ArcByCenterPointType]

	ArcStringTypes                map[*ArcStringType]struct{}
	ArcStringTypes_instance       map[*ArcStringType]*ArcStringType
	ArcStringTypes_mapString      map[string]*ArcStringType
	ArcStringTypeOrder            uint
	ArcStringType_stagedOrder     map[*ArcStringType]uint
	ArcStringType_orderStaged     map[uint]*ArcStringType
	ArcStringTypes_reference      map[*ArcStringType]*ArcStringType
	ArcStringTypes_referenceOrder map[*ArcStringType]uint

	// insertion point for slice of pointers maps
	ArcStringType_Pos_reverseMap map[*DirectPositionType]*ArcStringType

	ArcStringType_PointProperty_reverseMap map[*PointPropertyType]*ArcStringType

	OnAfterArcStringTypeCreateCallback GongOnAfterCreateInterface[ArcStringType]
	OnAfterArcStringTypeUpdateCallback GongOnAfterUpdateInterface[ArcStringType]
	OnAfterArcStringTypeDeleteCallback GongOnAfterDeleteInterface[ArcStringType]

	ArcTypes                map[*ArcType]struct{}
	ArcTypes_instance       map[*ArcType]*ArcType
	ArcTypes_mapString      map[string]*ArcType
	ArcTypeOrder            uint
	ArcType_stagedOrder     map[*ArcType]uint
	ArcType_orderStaged     map[uint]*ArcType
	ArcTypes_reference      map[*ArcType]*ArcType
	ArcTypes_referenceOrder map[*ArcType]uint

	// insertion point for slice of pointers maps
	OnAfterArcTypeCreateCallback GongOnAfterCreateInterface[ArcType]
	OnAfterArcTypeUpdateCallback GongOnAfterUpdateInterface[ArcType]
	OnAfterArcTypeDeleteCallback GongOnAfterDeleteInterface[ArcType]

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

	CircleByCenterPointTypes                map[*CircleByCenterPointType]struct{}
	CircleByCenterPointTypes_instance       map[*CircleByCenterPointType]*CircleByCenterPointType
	CircleByCenterPointTypes_mapString      map[string]*CircleByCenterPointType
	CircleByCenterPointTypeOrder            uint
	CircleByCenterPointType_stagedOrder     map[*CircleByCenterPointType]uint
	CircleByCenterPointType_orderStaged     map[uint]*CircleByCenterPointType
	CircleByCenterPointTypes_reference      map[*CircleByCenterPointType]*CircleByCenterPointType
	CircleByCenterPointTypes_referenceOrder map[*CircleByCenterPointType]uint

	// insertion point for slice of pointers maps
	OnAfterCircleByCenterPointTypeCreateCallback GongOnAfterCreateInterface[CircleByCenterPointType]
	OnAfterCircleByCenterPointTypeUpdateCallback GongOnAfterUpdateInterface[CircleByCenterPointType]
	OnAfterCircleByCenterPointTypeDeleteCallback GongOnAfterDeleteInterface[CircleByCenterPointType]

	CircleTypes                map[*CircleType]struct{}
	CircleTypes_instance       map[*CircleType]*CircleType
	CircleTypes_mapString      map[string]*CircleType
	CircleTypeOrder            uint
	CircleType_stagedOrder     map[*CircleType]uint
	CircleType_orderStaged     map[uint]*CircleType
	CircleTypes_reference      map[*CircleType]*CircleType
	CircleTypes_referenceOrder map[*CircleType]uint

	// insertion point for slice of pointers maps
	CircleType_Pos_reverseMap map[*DirectPositionType]*CircleType

	CircleType_PointProperty_reverseMap map[*PointPropertyType]*CircleType

	OnAfterCircleTypeCreateCallback GongOnAfterCreateInterface[CircleType]
	OnAfterCircleTypeUpdateCallback GongOnAfterUpdateInterface[CircleType]
	OnAfterCircleTypeDeleteCallback GongOnAfterDeleteInterface[CircleType]

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

	CompositeCurveTypes                map[*CompositeCurveType]struct{}
	CompositeCurveTypes_instance       map[*CompositeCurveType]*CompositeCurveType
	CompositeCurveTypes_mapString      map[string]*CompositeCurveType
	CompositeCurveTypeOrder            uint
	CompositeCurveType_stagedOrder     map[*CompositeCurveType]uint
	CompositeCurveType_orderStaged     map[uint]*CompositeCurveType
	CompositeCurveTypes_reference      map[*CompositeCurveType]*CompositeCurveType
	CompositeCurveTypes_referenceOrder map[*CompositeCurveType]uint

	// insertion point for slice of pointers maps
	CompositeCurveType_NameXSD_reverseMap map[*CodeType]*CompositeCurveType

	CompositeCurveType_CurveMember_reverseMap map[*CurvePropertyType]*CompositeCurveType

	OnAfterCompositeCurveTypeCreateCallback GongOnAfterCreateInterface[CompositeCurveType]
	OnAfterCompositeCurveTypeUpdateCallback GongOnAfterUpdateInterface[CompositeCurveType]
	OnAfterCompositeCurveTypeDeleteCallback GongOnAfterDeleteInterface[CompositeCurveType]

	CurvePropertyTypes                map[*CurvePropertyType]struct{}
	CurvePropertyTypes_instance       map[*CurvePropertyType]*CurvePropertyType
	CurvePropertyTypes_mapString      map[string]*CurvePropertyType
	CurvePropertyTypeOrder            uint
	CurvePropertyType_stagedOrder     map[*CurvePropertyType]uint
	CurvePropertyType_orderStaged     map[uint]*CurvePropertyType
	CurvePropertyTypes_reference      map[*CurvePropertyType]*CurvePropertyType
	CurvePropertyTypes_referenceOrder map[*CurvePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterCurvePropertyTypeCreateCallback GongOnAfterCreateInterface[CurvePropertyType]
	OnAfterCurvePropertyTypeUpdateCallback GongOnAfterUpdateInterface[CurvePropertyType]
	OnAfterCurvePropertyTypeDeleteCallback GongOnAfterDeleteInterface[CurvePropertyType]

	CurveSegmentArrayPropertyTypes                map[*CurveSegmentArrayPropertyType]struct{}
	CurveSegmentArrayPropertyTypes_instance       map[*CurveSegmentArrayPropertyType]*CurveSegmentArrayPropertyType
	CurveSegmentArrayPropertyTypes_mapString      map[string]*CurveSegmentArrayPropertyType
	CurveSegmentArrayPropertyTypeOrder            uint
	CurveSegmentArrayPropertyType_stagedOrder     map[*CurveSegmentArrayPropertyType]uint
	CurveSegmentArrayPropertyType_orderStaged     map[uint]*CurveSegmentArrayPropertyType
	CurveSegmentArrayPropertyTypes_reference      map[*CurveSegmentArrayPropertyType]*CurveSegmentArrayPropertyType
	CurveSegmentArrayPropertyTypes_referenceOrder map[*CurveSegmentArrayPropertyType]uint

	// insertion point for slice of pointers maps
	CurveSegmentArrayPropertyType_LineStringSegment_reverseMap map[*LineStringSegmentType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_ArcString_reverseMap map[*ArcStringType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_Arc_reverseMap map[*ArcType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_Circle_reverseMap map[*CircleType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_ArcByCenterPoint_reverseMap map[*ArcByCenterPointType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_CircleByCenterPoint_reverseMap map[*CircleByCenterPointType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_GeodesicString_reverseMap map[*GeodesicStringType]*CurveSegmentArrayPropertyType

	CurveSegmentArrayPropertyType_Geodesic_reverseMap map[*GeodesicType]*CurveSegmentArrayPropertyType

	OnAfterCurveSegmentArrayPropertyTypeCreateCallback GongOnAfterCreateInterface[CurveSegmentArrayPropertyType]
	OnAfterCurveSegmentArrayPropertyTypeUpdateCallback GongOnAfterUpdateInterface[CurveSegmentArrayPropertyType]
	OnAfterCurveSegmentArrayPropertyTypeDeleteCallback GongOnAfterDeleteInterface[CurveSegmentArrayPropertyType]

	CurveTypes                map[*CurveType]struct{}
	CurveTypes_instance       map[*CurveType]*CurveType
	CurveTypes_mapString      map[string]*CurveType
	CurveTypeOrder            uint
	CurveType_stagedOrder     map[*CurveType]uint
	CurveType_orderStaged     map[uint]*CurveType
	CurveTypes_reference      map[*CurveType]*CurveType
	CurveTypes_referenceOrder map[*CurveType]uint

	// insertion point for slice of pointers maps
	CurveType_NameXSD_reverseMap map[*CodeType]*CurveType

	OnAfterCurveTypeCreateCallback GongOnAfterCreateInterface[CurveType]
	OnAfterCurveTypeUpdateCallback GongOnAfterUpdateInterface[CurveType]
	OnAfterCurveTypeDeleteCallback GongOnAfterDeleteInterface[CurveType]

	DirectPositionListTypes                map[*DirectPositionListType]struct{}
	DirectPositionListTypes_instance       map[*DirectPositionListType]*DirectPositionListType
	DirectPositionListTypes_mapString      map[string]*DirectPositionListType
	DirectPositionListTypeOrder            uint
	DirectPositionListType_stagedOrder     map[*DirectPositionListType]uint
	DirectPositionListType_orderStaged     map[uint]*DirectPositionListType
	DirectPositionListTypes_reference      map[*DirectPositionListType]*DirectPositionListType
	DirectPositionListTypes_referenceOrder map[*DirectPositionListType]uint

	// insertion point for slice of pointers maps
	OnAfterDirectPositionListTypeCreateCallback GongOnAfterCreateInterface[DirectPositionListType]
	OnAfterDirectPositionListTypeUpdateCallback GongOnAfterUpdateInterface[DirectPositionListType]
	OnAfterDirectPositionListTypeDeleteCallback GongOnAfterDeleteInterface[DirectPositionListType]

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

	GeodesicStringTypes                map[*GeodesicStringType]struct{}
	GeodesicStringTypes_instance       map[*GeodesicStringType]*GeodesicStringType
	GeodesicStringTypes_mapString      map[string]*GeodesicStringType
	GeodesicStringTypeOrder            uint
	GeodesicStringType_stagedOrder     map[*GeodesicStringType]uint
	GeodesicStringType_orderStaged     map[uint]*GeodesicStringType
	GeodesicStringTypes_reference      map[*GeodesicStringType]*GeodesicStringType
	GeodesicStringTypes_referenceOrder map[*GeodesicStringType]uint

	// insertion point for slice of pointers maps
	OnAfterGeodesicStringTypeCreateCallback GongOnAfterCreateInterface[GeodesicStringType]
	OnAfterGeodesicStringTypeUpdateCallback GongOnAfterUpdateInterface[GeodesicStringType]
	OnAfterGeodesicStringTypeDeleteCallback GongOnAfterDeleteInterface[GeodesicStringType]

	GeodesicTypes                map[*GeodesicType]struct{}
	GeodesicTypes_instance       map[*GeodesicType]*GeodesicType
	GeodesicTypes_mapString      map[string]*GeodesicType
	GeodesicTypeOrder            uint
	GeodesicType_stagedOrder     map[*GeodesicType]uint
	GeodesicType_orderStaged     map[uint]*GeodesicType
	GeodesicTypes_reference      map[*GeodesicType]*GeodesicType
	GeodesicTypes_referenceOrder map[*GeodesicType]uint

	// insertion point for slice of pointers maps
	OnAfterGeodesicTypeCreateCallback GongOnAfterCreateInterface[GeodesicType]
	OnAfterGeodesicTypeUpdateCallback GongOnAfterUpdateInterface[GeodesicType]
	OnAfterGeodesicTypeDeleteCallback GongOnAfterDeleteInterface[GeodesicType]

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

	LineStringSegmentTypes                map[*LineStringSegmentType]struct{}
	LineStringSegmentTypes_instance       map[*LineStringSegmentType]*LineStringSegmentType
	LineStringSegmentTypes_mapString      map[string]*LineStringSegmentType
	LineStringSegmentTypeOrder            uint
	LineStringSegmentType_stagedOrder     map[*LineStringSegmentType]uint
	LineStringSegmentType_orderStaged     map[uint]*LineStringSegmentType
	LineStringSegmentTypes_reference      map[*LineStringSegmentType]*LineStringSegmentType
	LineStringSegmentTypes_referenceOrder map[*LineStringSegmentType]uint

	// insertion point for slice of pointers maps
	LineStringSegmentType_Pos_reverseMap map[*DirectPositionType]*LineStringSegmentType

	LineStringSegmentType_PointProperty_reverseMap map[*PointPropertyType]*LineStringSegmentType

	OnAfterLineStringSegmentTypeCreateCallback GongOnAfterCreateInterface[LineStringSegmentType]
	OnAfterLineStringSegmentTypeUpdateCallback GongOnAfterUpdateInterface[LineStringSegmentType]
	OnAfterLineStringSegmentTypeDeleteCallback GongOnAfterDeleteInterface[LineStringSegmentType]

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

	OrientableCurveTypes                map[*OrientableCurveType]struct{}
	OrientableCurveTypes_instance       map[*OrientableCurveType]*OrientableCurveType
	OrientableCurveTypes_mapString      map[string]*OrientableCurveType
	OrientableCurveTypeOrder            uint
	OrientableCurveType_stagedOrder     map[*OrientableCurveType]uint
	OrientableCurveType_orderStaged     map[uint]*OrientableCurveType
	OrientableCurveTypes_reference      map[*OrientableCurveType]*OrientableCurveType
	OrientableCurveTypes_referenceOrder map[*OrientableCurveType]uint

	// insertion point for slice of pointers maps
	OrientableCurveType_NameXSD_reverseMap map[*CodeType]*OrientableCurveType

	OnAfterOrientableCurveTypeCreateCallback GongOnAfterCreateInterface[OrientableCurveType]
	OnAfterOrientableCurveTypeUpdateCallback GongOnAfterUpdateInterface[OrientableCurveType]
	OnAfterOrientableCurveTypeDeleteCallback GongOnAfterDeleteInterface[OrientableCurveType]

	PointPropertyTypes                map[*PointPropertyType]struct{}
	PointPropertyTypes_instance       map[*PointPropertyType]*PointPropertyType
	PointPropertyTypes_mapString      map[string]*PointPropertyType
	PointPropertyTypeOrder            uint
	PointPropertyType_stagedOrder     map[*PointPropertyType]uint
	PointPropertyType_orderStaged     map[uint]*PointPropertyType
	PointPropertyTypes_reference      map[*PointPropertyType]*PointPropertyType
	PointPropertyTypes_referenceOrder map[*PointPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterPointPropertyTypeCreateCallback GongOnAfterCreateInterface[PointPropertyType]
	OnAfterPointPropertyTypeUpdateCallback GongOnAfterUpdateInterface[PointPropertyType]
	OnAfterPointPropertyTypeDeleteCallback GongOnAfterDeleteInterface[PointPropertyType]

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

	PolygonPatchTypes                map[*PolygonPatchType]struct{}
	PolygonPatchTypes_instance       map[*PolygonPatchType]*PolygonPatchType
	PolygonPatchTypes_mapString      map[string]*PolygonPatchType
	PolygonPatchTypeOrder            uint
	PolygonPatchType_stagedOrder     map[*PolygonPatchType]uint
	PolygonPatchType_orderStaged     map[uint]*PolygonPatchType
	PolygonPatchTypes_reference      map[*PolygonPatchType]*PolygonPatchType
	PolygonPatchTypes_referenceOrder map[*PolygonPatchType]uint

	// insertion point for slice of pointers maps
	OnAfterPolygonPatchTypeCreateCallback GongOnAfterCreateInterface[PolygonPatchType]
	OnAfterPolygonPatchTypeUpdateCallback GongOnAfterUpdateInterface[PolygonPatchType]
	OnAfterPolygonPatchTypeDeleteCallback GongOnAfterDeleteInterface[PolygonPatchType]

	RingTypes                map[*RingType]struct{}
	RingTypes_instance       map[*RingType]*RingType
	RingTypes_mapString      map[string]*RingType
	RingTypeOrder            uint
	RingType_stagedOrder     map[*RingType]uint
	RingType_orderStaged     map[uint]*RingType
	RingTypes_reference      map[*RingType]*RingType
	RingTypes_referenceOrder map[*RingType]uint

	// insertion point for slice of pointers maps
	RingType_CurveMember_reverseMap map[*CurvePropertyType]*RingType

	OnAfterRingTypeCreateCallback GongOnAfterCreateInterface[RingType]
	OnAfterRingTypeUpdateCallback GongOnAfterUpdateInterface[RingType]
	OnAfterRingTypeDeleteCallback GongOnAfterDeleteInterface[RingType]

	SurfacePatchArrayPropertyTypes                map[*SurfacePatchArrayPropertyType]struct{}
	SurfacePatchArrayPropertyTypes_instance       map[*SurfacePatchArrayPropertyType]*SurfacePatchArrayPropertyType
	SurfacePatchArrayPropertyTypes_mapString      map[string]*SurfacePatchArrayPropertyType
	SurfacePatchArrayPropertyTypeOrder            uint
	SurfacePatchArrayPropertyType_stagedOrder     map[*SurfacePatchArrayPropertyType]uint
	SurfacePatchArrayPropertyType_orderStaged     map[uint]*SurfacePatchArrayPropertyType
	SurfacePatchArrayPropertyTypes_reference      map[*SurfacePatchArrayPropertyType]*SurfacePatchArrayPropertyType
	SurfacePatchArrayPropertyTypes_referenceOrder map[*SurfacePatchArrayPropertyType]uint

	// insertion point for slice of pointers maps
	SurfacePatchArrayPropertyType_PolygonPatch_reverseMap map[*PolygonPatchType]*SurfacePatchArrayPropertyType

	OnAfterSurfacePatchArrayPropertyTypeCreateCallback GongOnAfterCreateInterface[SurfacePatchArrayPropertyType]
	OnAfterSurfacePatchArrayPropertyTypeUpdateCallback GongOnAfterUpdateInterface[SurfacePatchArrayPropertyType]
	OnAfterSurfacePatchArrayPropertyTypeDeleteCallback GongOnAfterDeleteInterface[SurfacePatchArrayPropertyType]

	SurfaceTypes                map[*SurfaceType]struct{}
	SurfaceTypes_instance       map[*SurfaceType]*SurfaceType
	SurfaceTypes_mapString      map[string]*SurfaceType
	SurfaceTypeOrder            uint
	SurfaceType_stagedOrder     map[*SurfaceType]uint
	SurfaceType_orderStaged     map[uint]*SurfaceType
	SurfaceTypes_reference      map[*SurfaceType]*SurfaceType
	SurfaceTypes_referenceOrder map[*SurfaceType]uint

	// insertion point for slice of pointers maps
	SurfaceType_NameXSD_reverseMap map[*CodeType]*SurfaceType

	OnAfterSurfaceTypeCreateCallback GongOnAfterCreateInterface[SurfaceType]
	OnAfterSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[SurfaceType]
	OnAfterSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[SurfaceType]

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
	__gong__clearReferences(&stage.AbstractCurveSegmentTypes_reference, &stage.AbstractCurveSegmentTypes_instance, &stage.AbstractCurveSegmentTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractCurveTypes_reference, &stage.AbstractCurveTypes_instance, &stage.AbstractCurveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractFeatureMemberTypes_reference, &stage.AbstractFeatureMemberTypes_instance, &stage.AbstractFeatureMemberTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractFeatureTypes_reference, &stage.AbstractFeatureTypes_instance, &stage.AbstractFeatureTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGMLTypes_reference, &stage.AbstractGMLTypes_instance, &stage.AbstractGMLTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGeometricPrimitiveTypes_reference, &stage.AbstractGeometricPrimitiveTypes_instance, &stage.AbstractGeometricPrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractGeometryTypes_reference, &stage.AbstractGeometryTypes_instance, &stage.AbstractGeometryTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractMetadataPropertyTypes_reference, &stage.AbstractMetadataPropertyTypes_instance, &stage.AbstractMetadataPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractRingPropertyTypes_reference, &stage.AbstractRingPropertyTypes_instance, &stage.AbstractRingPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractRingTypes_reference, &stage.AbstractRingTypes_instance, &stage.AbstractRingTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractSurfacePatchTypes_reference, &stage.AbstractSurfacePatchTypes_instance, &stage.AbstractSurfacePatchTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractSurfaceTypes_reference, &stage.AbstractSurfaceTypes_instance, &stage.AbstractSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeGeometricPrimitiveTypes_reference, &stage.AbstractTimeGeometricPrimitiveTypes_instance, &stage.AbstractTimeGeometricPrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeObjectTypes_reference, &stage.AbstractTimeObjectTypes_instance, &stage.AbstractTimeObjectTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimePrimitiveTypes_reference, &stage.AbstractTimePrimitiveTypes_instance, &stage.AbstractTimePrimitiveTypes_referenceOrder)

	__gong__clearReferences(&stage.AbstractTimeSliceTypes_reference, &stage.AbstractTimeSliceTypes_instance, &stage.AbstractTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.AngleTypes_reference, &stage.AngleTypes_instance, &stage.AngleTypes_referenceOrder)

	__gong__clearReferences(&stage.ArcByCenterPointTypes_reference, &stage.ArcByCenterPointTypes_instance, &stage.ArcByCenterPointTypes_referenceOrder)

	__gong__clearReferences(&stage.ArcStringTypes_reference, &stage.ArcStringTypes_instance, &stage.ArcStringTypes_referenceOrder)

	__gong__clearReferences(&stage.ArcTypes_reference, &stage.ArcTypes_instance, &stage.ArcTypes_referenceOrder)

	__gong__clearReferences(&stage.BoundingShapeTypes_reference, &stage.BoundingShapeTypes_instance, &stage.BoundingShapeTypes_referenceOrder)

	__gong__clearReferences(&stage.CircleByCenterPointTypes_reference, &stage.CircleByCenterPointTypes_instance, &stage.CircleByCenterPointTypes_referenceOrder)

	__gong__clearReferences(&stage.CircleTypes_reference, &stage.CircleTypes_instance, &stage.CircleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTypes_reference, &stage.CodeTypes_instance, &stage.CodeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWithAuthorityTypes_reference, &stage.CodeWithAuthorityTypes_instance, &stage.CodeWithAuthorityTypes_referenceOrder)

	__gong__clearReferences(&stage.CompositeCurveTypes_reference, &stage.CompositeCurveTypes_instance, &stage.CompositeCurveTypes_referenceOrder)

	__gong__clearReferences(&stage.CurvePropertyTypes_reference, &stage.CurvePropertyTypes_instance, &stage.CurvePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.CurveSegmentArrayPropertyTypes_reference, &stage.CurveSegmentArrayPropertyTypes_instance, &stage.CurveSegmentArrayPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.CurveTypes_reference, &stage.CurveTypes_instance, &stage.CurveTypes_referenceOrder)

	__gong__clearReferences(&stage.DirectPositionListTypes_reference, &stage.DirectPositionListTypes_instance, &stage.DirectPositionListTypes_referenceOrder)

	__gong__clearReferences(&stage.DirectPositionTypes_reference, &stage.DirectPositionTypes_instance, &stage.DirectPositionTypes_referenceOrder)

	__gong__clearReferences(&stage.DynamicFeatureTypes_reference, &stage.DynamicFeatureTypes_instance, &stage.DynamicFeatureTypes_referenceOrder)

	__gong__clearReferences(&stage.EnvelopeTypes_reference, &stage.EnvelopeTypes_instance, &stage.EnvelopeTypes_referenceOrder)

	__gong__clearReferences(&stage.GeodesicStringTypes_reference, &stage.GeodesicStringTypes_instance, &stage.GeodesicStringTypes_referenceOrder)

	__gong__clearReferences(&stage.GeodesicTypes_reference, &stage.GeodesicTypes_instance, &stage.GeodesicTypes_referenceOrder)

	__gong__clearReferences(&stage.LengthTypes_reference, &stage.LengthTypes_instance, &stage.LengthTypes_referenceOrder)

	__gong__clearReferences(&stage.LineStringSegmentTypes_reference, &stage.LineStringSegmentTypes_instance, &stage.LineStringSegmentTypes_referenceOrder)

	__gong__clearReferences(&stage.MeasureTypes_reference, &stage.MeasureTypes_instance, &stage.MeasureTypes_referenceOrder)

	__gong__clearReferences(&stage.OrientableCurveTypes_reference, &stage.OrientableCurveTypes_instance, &stage.OrientableCurveTypes_referenceOrder)

	__gong__clearReferences(&stage.PointPropertyTypes_reference, &stage.PointPropertyTypes_instance, &stage.PointPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.PointTypes_reference, &stage.PointTypes_instance, &stage.PointTypes_referenceOrder)

	__gong__clearReferences(&stage.PolygonPatchTypes_reference, &stage.PolygonPatchTypes_instance, &stage.PolygonPatchTypes_referenceOrder)

	__gong__clearReferences(&stage.RingTypes_reference, &stage.RingTypes_instance, &stage.RingTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfacePatchArrayPropertyTypes_reference, &stage.SurfacePatchArrayPropertyTypes_instance, &stage.SurfacePatchArrayPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfaceTypes_reference, &stage.SurfaceTypes_instance, &stage.SurfaceTypes_referenceOrder)

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
	stage.AbstractCurveSegmentTypeOrder = __gong__recomputeOrder(stage.AbstractCurveSegmentType_stagedOrder)

	stage.AbstractCurveTypeOrder = __gong__recomputeOrder(stage.AbstractCurveType_stagedOrder)

	stage.AbstractFeatureMemberTypeOrder = __gong__recomputeOrder(stage.AbstractFeatureMemberType_stagedOrder)

	stage.AbstractFeatureTypeOrder = __gong__recomputeOrder(stage.AbstractFeatureType_stagedOrder)

	stage.AbstractGMLTypeOrder = __gong__recomputeOrder(stage.AbstractGMLType_stagedOrder)

	stage.AbstractGeometricPrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractGeometricPrimitiveType_stagedOrder)

	stage.AbstractGeometryTypeOrder = __gong__recomputeOrder(stage.AbstractGeometryType_stagedOrder)

	stage.AbstractMetadataPropertyTypeOrder = __gong__recomputeOrder(stage.AbstractMetadataPropertyType_stagedOrder)

	stage.AbstractRingPropertyTypeOrder = __gong__recomputeOrder(stage.AbstractRingPropertyType_stagedOrder)

	stage.AbstractRingTypeOrder = __gong__recomputeOrder(stage.AbstractRingType_stagedOrder)

	stage.AbstractSurfacePatchTypeOrder = __gong__recomputeOrder(stage.AbstractSurfacePatchType_stagedOrder)

	stage.AbstractSurfaceTypeOrder = __gong__recomputeOrder(stage.AbstractSurfaceType_stagedOrder)

	stage.AbstractTimeGeometricPrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractTimeGeometricPrimitiveType_stagedOrder)

	stage.AbstractTimeObjectTypeOrder = __gong__recomputeOrder(stage.AbstractTimeObjectType_stagedOrder)

	stage.AbstractTimePrimitiveTypeOrder = __gong__recomputeOrder(stage.AbstractTimePrimitiveType_stagedOrder)

	stage.AbstractTimeSliceTypeOrder = __gong__recomputeOrder(stage.AbstractTimeSliceType_stagedOrder)

	stage.AngleTypeOrder = __gong__recomputeOrder(stage.AngleType_stagedOrder)

	stage.ArcByCenterPointTypeOrder = __gong__recomputeOrder(stage.ArcByCenterPointType_stagedOrder)

	stage.ArcStringTypeOrder = __gong__recomputeOrder(stage.ArcStringType_stagedOrder)

	stage.ArcTypeOrder = __gong__recomputeOrder(stage.ArcType_stagedOrder)

	stage.BoundingShapeTypeOrder = __gong__recomputeOrder(stage.BoundingShapeType_stagedOrder)

	stage.CircleByCenterPointTypeOrder = __gong__recomputeOrder(stage.CircleByCenterPointType_stagedOrder)

	stage.CircleTypeOrder = __gong__recomputeOrder(stage.CircleType_stagedOrder)

	stage.CodeTypeOrder = __gong__recomputeOrder(stage.CodeType_stagedOrder)

	stage.CodeWithAuthorityTypeOrder = __gong__recomputeOrder(stage.CodeWithAuthorityType_stagedOrder)

	stage.CompositeCurveTypeOrder = __gong__recomputeOrder(stage.CompositeCurveType_stagedOrder)

	stage.CurvePropertyTypeOrder = __gong__recomputeOrder(stage.CurvePropertyType_stagedOrder)

	stage.CurveSegmentArrayPropertyTypeOrder = __gong__recomputeOrder(stage.CurveSegmentArrayPropertyType_stagedOrder)

	stage.CurveTypeOrder = __gong__recomputeOrder(stage.CurveType_stagedOrder)

	stage.DirectPositionListTypeOrder = __gong__recomputeOrder(stage.DirectPositionListType_stagedOrder)

	stage.DirectPositionTypeOrder = __gong__recomputeOrder(stage.DirectPositionType_stagedOrder)

	stage.DynamicFeatureTypeOrder = __gong__recomputeOrder(stage.DynamicFeatureType_stagedOrder)

	stage.EnvelopeTypeOrder = __gong__recomputeOrder(stage.EnvelopeType_stagedOrder)

	stage.GeodesicStringTypeOrder = __gong__recomputeOrder(stage.GeodesicStringType_stagedOrder)

	stage.GeodesicTypeOrder = __gong__recomputeOrder(stage.GeodesicType_stagedOrder)

	stage.LengthTypeOrder = __gong__recomputeOrder(stage.LengthType_stagedOrder)

	stage.LineStringSegmentTypeOrder = __gong__recomputeOrder(stage.LineStringSegmentType_stagedOrder)

	stage.MeasureTypeOrder = __gong__recomputeOrder(stage.MeasureType_stagedOrder)

	stage.OrientableCurveTypeOrder = __gong__recomputeOrder(stage.OrientableCurveType_stagedOrder)

	stage.PointPropertyTypeOrder = __gong__recomputeOrder(stage.PointPropertyType_stagedOrder)

	stage.PointTypeOrder = __gong__recomputeOrder(stage.PointType_stagedOrder)

	stage.PolygonPatchTypeOrder = __gong__recomputeOrder(stage.PolygonPatchType_stagedOrder)

	stage.RingTypeOrder = __gong__recomputeOrder(stage.RingType_stagedOrder)

	stage.SurfacePatchArrayPropertyTypeOrder = __gong__recomputeOrder(stage.SurfacePatchArrayPropertyType_stagedOrder)

	stage.SurfaceTypeOrder = __gong__recomputeOrder(stage.SurfaceType_stagedOrder)

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
func (*AbstractCurveSegmentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractCurveSegmentTypes, stage.AbstractCurveSegmentType_stagedOrder)
}

func (*AbstractCurveSegmentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractCurveSegmentType_orderStaged[order]
}

func (*AbstractCurveSegmentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractCurveSegmentTypes_mapString
}

func (*AbstractCurveSegmentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractCurveSegmentTypes
}

func (*AbstractCurveSegmentType) GongNewInstance() any {
	return new(AbstractCurveSegmentType)
}

func (*AbstractCurveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractCurveTypes, stage.AbstractCurveType_stagedOrder)
}

func (*AbstractCurveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractCurveType_orderStaged[order]
}

func (*AbstractCurveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractCurveTypes_mapString
}

func (*AbstractCurveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractCurveTypes
}

func (*AbstractCurveType) GongNewInstance() any {
	return new(AbstractCurveType)
}

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

func (*AbstractRingPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractRingPropertyTypes, stage.AbstractRingPropertyType_stagedOrder)
}

func (*AbstractRingPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractRingPropertyType_orderStaged[order]
}

func (*AbstractRingPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractRingPropertyTypes_mapString
}

func (*AbstractRingPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractRingPropertyTypes
}

func (*AbstractRingPropertyType) GongNewInstance() any {
	return new(AbstractRingPropertyType)
}

func (*AbstractRingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractRingTypes, stage.AbstractRingType_stagedOrder)
}

func (*AbstractRingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractRingType_orderStaged[order]
}

func (*AbstractRingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractRingTypes_mapString
}

func (*AbstractRingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractRingTypes
}

func (*AbstractRingType) GongNewInstance() any {
	return new(AbstractRingType)
}

func (*AbstractSurfacePatchType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractSurfacePatchTypes, stage.AbstractSurfacePatchType_stagedOrder)
}

func (*AbstractSurfacePatchType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractSurfacePatchType_orderStaged[order]
}

func (*AbstractSurfacePatchType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractSurfacePatchTypes_mapString
}

func (*AbstractSurfacePatchType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractSurfacePatchTypes
}

func (*AbstractSurfacePatchType) GongNewInstance() any {
	return new(AbstractSurfacePatchType)
}

func (*AbstractSurfaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractSurfaceTypes, stage.AbstractSurfaceType_stagedOrder)
}

func (*AbstractSurfaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractSurfaceType_orderStaged[order]
}

func (*AbstractSurfaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractSurfaceTypes_mapString
}

func (*AbstractSurfaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractSurfaceTypes
}

func (*AbstractSurfaceType) GongNewInstance() any {
	return new(AbstractSurfaceType)
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

func (*ArcByCenterPointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ArcByCenterPointTypes, stage.ArcByCenterPointType_stagedOrder)
}

func (*ArcByCenterPointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ArcByCenterPointType_orderStaged[order]
}

func (*ArcByCenterPointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ArcByCenterPointTypes_mapString
}

func (*ArcByCenterPointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ArcByCenterPointTypes
}

func (*ArcByCenterPointType) GongNewInstance() any {
	return new(ArcByCenterPointType)
}

func (*ArcStringType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ArcStringTypes, stage.ArcStringType_stagedOrder)
}

func (*ArcStringType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ArcStringType_orderStaged[order]
}

func (*ArcStringType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ArcStringTypes_mapString
}

func (*ArcStringType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ArcStringTypes
}

func (*ArcStringType) GongNewInstance() any {
	return new(ArcStringType)
}

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

func (*CircleByCenterPointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CircleByCenterPointTypes, stage.CircleByCenterPointType_stagedOrder)
}

func (*CircleByCenterPointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CircleByCenterPointType_orderStaged[order]
}

func (*CircleByCenterPointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CircleByCenterPointTypes_mapString
}

func (*CircleByCenterPointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CircleByCenterPointTypes
}

func (*CircleByCenterPointType) GongNewInstance() any {
	return new(CircleByCenterPointType)
}

func (*CircleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CircleTypes, stage.CircleType_stagedOrder)
}

func (*CircleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CircleType_orderStaged[order]
}

func (*CircleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CircleTypes_mapString
}

func (*CircleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CircleTypes
}

func (*CircleType) GongNewInstance() any {
	return new(CircleType)
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

func (*CompositeCurveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CompositeCurveTypes, stage.CompositeCurveType_stagedOrder)
}

func (*CompositeCurveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CompositeCurveType_orderStaged[order]
}

func (*CompositeCurveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CompositeCurveTypes_mapString
}

func (*CompositeCurveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CompositeCurveTypes
}

func (*CompositeCurveType) GongNewInstance() any {
	return new(CompositeCurveType)
}

func (*CurvePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder)
}

func (*CurvePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CurvePropertyType_orderStaged[order]
}

func (*CurvePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CurvePropertyTypes_mapString
}

func (*CurvePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CurvePropertyTypes
}

func (*CurvePropertyType) GongNewInstance() any {
	return new(CurvePropertyType)
}

func (*CurveSegmentArrayPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CurveSegmentArrayPropertyTypes, stage.CurveSegmentArrayPropertyType_stagedOrder)
}

func (*CurveSegmentArrayPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CurveSegmentArrayPropertyType_orderStaged[order]
}

func (*CurveSegmentArrayPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CurveSegmentArrayPropertyTypes_mapString
}

func (*CurveSegmentArrayPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CurveSegmentArrayPropertyTypes
}

func (*CurveSegmentArrayPropertyType) GongNewInstance() any {
	return new(CurveSegmentArrayPropertyType)
}

func (*CurveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CurveTypes, stage.CurveType_stagedOrder)
}

func (*CurveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CurveType_orderStaged[order]
}

func (*CurveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CurveTypes_mapString
}

func (*CurveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CurveTypes
}

func (*CurveType) GongNewInstance() any {
	return new(CurveType)
}

func (*DirectPositionListType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DirectPositionListTypes, stage.DirectPositionListType_stagedOrder)
}

func (*DirectPositionListType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DirectPositionListType_orderStaged[order]
}

func (*DirectPositionListType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DirectPositionListTypes_mapString
}

func (*DirectPositionListType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DirectPositionListTypes
}

func (*DirectPositionListType) GongNewInstance() any {
	return new(DirectPositionListType)
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

func (*GeodesicStringType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GeodesicStringTypes, stage.GeodesicStringType_stagedOrder)
}

func (*GeodesicStringType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GeodesicStringType_orderStaged[order]
}

func (*GeodesicStringType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GeodesicStringTypes_mapString
}

func (*GeodesicStringType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GeodesicStringTypes
}

func (*GeodesicStringType) GongNewInstance() any {
	return new(GeodesicStringType)
}

func (*GeodesicType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GeodesicTypes, stage.GeodesicType_stagedOrder)
}

func (*GeodesicType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GeodesicType_orderStaged[order]
}

func (*GeodesicType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GeodesicTypes_mapString
}

func (*GeodesicType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GeodesicTypes
}

func (*GeodesicType) GongNewInstance() any {
	return new(GeodesicType)
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

func (*LineStringSegmentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LineStringSegmentTypes, stage.LineStringSegmentType_stagedOrder)
}

func (*LineStringSegmentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LineStringSegmentType_orderStaged[order]
}

func (*LineStringSegmentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LineStringSegmentTypes_mapString
}

func (*LineStringSegmentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LineStringSegmentTypes
}

func (*LineStringSegmentType) GongNewInstance() any {
	return new(LineStringSegmentType)
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

func (*OrientableCurveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OrientableCurveTypes, stage.OrientableCurveType_stagedOrder)
}

func (*OrientableCurveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OrientableCurveType_orderStaged[order]
}

func (*OrientableCurveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OrientableCurveTypes_mapString
}

func (*OrientableCurveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.OrientableCurveTypes
}

func (*OrientableCurveType) GongNewInstance() any {
	return new(OrientableCurveType)
}

func (*PointPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PointPropertyTypes, stage.PointPropertyType_stagedOrder)
}

func (*PointPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PointPropertyType_orderStaged[order]
}

func (*PointPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PointPropertyTypes_mapString
}

func (*PointPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.PointPropertyTypes
}

func (*PointPropertyType) GongNewInstance() any {
	return new(PointPropertyType)
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

func (*PolygonPatchType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PolygonPatchTypes, stage.PolygonPatchType_stagedOrder)
}

func (*PolygonPatchType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PolygonPatchType_orderStaged[order]
}

func (*PolygonPatchType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PolygonPatchTypes_mapString
}

func (*PolygonPatchType) GongGetInstancesSet(stage *Stage) any {
	return &stage.PolygonPatchTypes
}

func (*PolygonPatchType) GongNewInstance() any {
	return new(PolygonPatchType)
}

func (*RingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.RingTypes, stage.RingType_stagedOrder)
}

func (*RingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.RingType_orderStaged[order]
}

func (*RingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.RingTypes_mapString
}

func (*RingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.RingTypes
}

func (*RingType) GongNewInstance() any {
	return new(RingType)
}

func (*SurfacePatchArrayPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SurfacePatchArrayPropertyTypes, stage.SurfacePatchArrayPropertyType_stagedOrder)
}

func (*SurfacePatchArrayPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SurfacePatchArrayPropertyType_orderStaged[order]
}

func (*SurfacePatchArrayPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SurfacePatchArrayPropertyTypes_mapString
}

func (*SurfacePatchArrayPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.SurfacePatchArrayPropertyTypes
}

func (*SurfacePatchArrayPropertyType) GongNewInstance() any {
	return new(SurfacePatchArrayPropertyType)
}

func (*SurfaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SurfaceTypes, stage.SurfaceType_stagedOrder)
}

func (*SurfaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SurfaceType_orderStaged[order]
}

func (*SurfaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SurfaceTypes_mapString
}

func (*SurfaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.SurfaceTypes
}

func (*SurfaceType) GongNewInstance() any {
	return new(SurfaceType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml"
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
		AbstractCurveSegmentTypes:           make(map[*AbstractCurveSegmentType]struct{}),
		AbstractCurveSegmentTypes_mapString: make(map[string]*AbstractCurveSegmentType),

		AbstractCurveTypes:           make(map[*AbstractCurveType]struct{}),
		AbstractCurveTypes_mapString: make(map[string]*AbstractCurveType),

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

		AbstractRingPropertyTypes:           make(map[*AbstractRingPropertyType]struct{}),
		AbstractRingPropertyTypes_mapString: make(map[string]*AbstractRingPropertyType),

		AbstractRingTypes:           make(map[*AbstractRingType]struct{}),
		AbstractRingTypes_mapString: make(map[string]*AbstractRingType),

		AbstractSurfacePatchTypes:           make(map[*AbstractSurfacePatchType]struct{}),
		AbstractSurfacePatchTypes_mapString: make(map[string]*AbstractSurfacePatchType),

		AbstractSurfaceTypes:           make(map[*AbstractSurfaceType]struct{}),
		AbstractSurfaceTypes_mapString: make(map[string]*AbstractSurfaceType),

		AbstractTimeGeometricPrimitiveTypes:           make(map[*AbstractTimeGeometricPrimitiveType]struct{}),
		AbstractTimeGeometricPrimitiveTypes_mapString: make(map[string]*AbstractTimeGeometricPrimitiveType),

		AbstractTimeObjectTypes:           make(map[*AbstractTimeObjectType]struct{}),
		AbstractTimeObjectTypes_mapString: make(map[string]*AbstractTimeObjectType),

		AbstractTimePrimitiveTypes:           make(map[*AbstractTimePrimitiveType]struct{}),
		AbstractTimePrimitiveTypes_mapString: make(map[string]*AbstractTimePrimitiveType),

		AbstractTimeSliceTypes:           make(map[*AbstractTimeSliceType]struct{}),
		AbstractTimeSliceTypes_mapString: make(map[string]*AbstractTimeSliceType),

		AngleTypes:           make(map[*AngleType]struct{}),
		AngleTypes_mapString: make(map[string]*AngleType),

		ArcByCenterPointTypes:           make(map[*ArcByCenterPointType]struct{}),
		ArcByCenterPointTypes_mapString: make(map[string]*ArcByCenterPointType),

		ArcStringTypes:           make(map[*ArcStringType]struct{}),
		ArcStringTypes_mapString: make(map[string]*ArcStringType),

		ArcTypes:           make(map[*ArcType]struct{}),
		ArcTypes_mapString: make(map[string]*ArcType),

		BoundingShapeTypes:           make(map[*BoundingShapeType]struct{}),
		BoundingShapeTypes_mapString: make(map[string]*BoundingShapeType),

		CircleByCenterPointTypes:           make(map[*CircleByCenterPointType]struct{}),
		CircleByCenterPointTypes_mapString: make(map[string]*CircleByCenterPointType),

		CircleTypes:           make(map[*CircleType]struct{}),
		CircleTypes_mapString: make(map[string]*CircleType),

		CodeTypes:           make(map[*CodeType]struct{}),
		CodeTypes_mapString: make(map[string]*CodeType),

		CodeWithAuthorityTypes:           make(map[*CodeWithAuthorityType]struct{}),
		CodeWithAuthorityTypes_mapString: make(map[string]*CodeWithAuthorityType),

		CompositeCurveTypes:           make(map[*CompositeCurveType]struct{}),
		CompositeCurveTypes_mapString: make(map[string]*CompositeCurveType),

		CurvePropertyTypes:           make(map[*CurvePropertyType]struct{}),
		CurvePropertyTypes_mapString: make(map[string]*CurvePropertyType),

		CurveSegmentArrayPropertyTypes:           make(map[*CurveSegmentArrayPropertyType]struct{}),
		CurveSegmentArrayPropertyTypes_mapString: make(map[string]*CurveSegmentArrayPropertyType),

		CurveTypes:           make(map[*CurveType]struct{}),
		CurveTypes_mapString: make(map[string]*CurveType),

		DirectPositionListTypes:           make(map[*DirectPositionListType]struct{}),
		DirectPositionListTypes_mapString: make(map[string]*DirectPositionListType),

		DirectPositionTypes:           make(map[*DirectPositionType]struct{}),
		DirectPositionTypes_mapString: make(map[string]*DirectPositionType),

		DynamicFeatureTypes:           make(map[*DynamicFeatureType]struct{}),
		DynamicFeatureTypes_mapString: make(map[string]*DynamicFeatureType),

		EnvelopeTypes:           make(map[*EnvelopeType]struct{}),
		EnvelopeTypes_mapString: make(map[string]*EnvelopeType),

		GeodesicStringTypes:           make(map[*GeodesicStringType]struct{}),
		GeodesicStringTypes_mapString: make(map[string]*GeodesicStringType),

		GeodesicTypes:           make(map[*GeodesicType]struct{}),
		GeodesicTypes_mapString: make(map[string]*GeodesicType),

		LengthTypes:           make(map[*LengthType]struct{}),
		LengthTypes_mapString: make(map[string]*LengthType),

		LineStringSegmentTypes:           make(map[*LineStringSegmentType]struct{}),
		LineStringSegmentTypes_mapString: make(map[string]*LineStringSegmentType),

		MeasureTypes:           make(map[*MeasureType]struct{}),
		MeasureTypes_mapString: make(map[string]*MeasureType),

		OrientableCurveTypes:           make(map[*OrientableCurveType]struct{}),
		OrientableCurveTypes_mapString: make(map[string]*OrientableCurveType),

		PointPropertyTypes:           make(map[*PointPropertyType]struct{}),
		PointPropertyTypes_mapString: make(map[string]*PointPropertyType),

		PointTypes:           make(map[*PointType]struct{}),
		PointTypes_mapString: make(map[string]*PointType),

		PolygonPatchTypes:           make(map[*PolygonPatchType]struct{}),
		PolygonPatchTypes_mapString: make(map[string]*PolygonPatchType),

		RingTypes:           make(map[*RingType]struct{}),
		RingTypes_mapString: make(map[string]*RingType),

		SurfacePatchArrayPropertyTypes:           make(map[*SurfacePatchArrayPropertyType]struct{}),
		SurfacePatchArrayPropertyTypes_mapString: make(map[string]*SurfacePatchArrayPropertyType),

		SurfaceTypes:           make(map[*SurfaceType]struct{}),
		SurfaceTypes_mapString: make(map[string]*SurfaceType),

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
		AbstractCurveSegmentType_stagedOrder: make(map[*AbstractCurveSegmentType]uint),
		AbstractCurveSegmentType_orderStaged: make(map[uint]*AbstractCurveSegmentType),
		AbstractCurveSegmentTypes_reference:  make(map[*AbstractCurveSegmentType]*AbstractCurveSegmentType),

		AbstractCurveType_stagedOrder: make(map[*AbstractCurveType]uint),
		AbstractCurveType_orderStaged: make(map[uint]*AbstractCurveType),
		AbstractCurveTypes_reference:  make(map[*AbstractCurveType]*AbstractCurveType),

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

		AbstractRingPropertyType_stagedOrder: make(map[*AbstractRingPropertyType]uint),
		AbstractRingPropertyType_orderStaged: make(map[uint]*AbstractRingPropertyType),
		AbstractRingPropertyTypes_reference:  make(map[*AbstractRingPropertyType]*AbstractRingPropertyType),

		AbstractRingType_stagedOrder: make(map[*AbstractRingType]uint),
		AbstractRingType_orderStaged: make(map[uint]*AbstractRingType),
		AbstractRingTypes_reference:  make(map[*AbstractRingType]*AbstractRingType),

		AbstractSurfacePatchType_stagedOrder: make(map[*AbstractSurfacePatchType]uint),
		AbstractSurfacePatchType_orderStaged: make(map[uint]*AbstractSurfacePatchType),
		AbstractSurfacePatchTypes_reference:  make(map[*AbstractSurfacePatchType]*AbstractSurfacePatchType),

		AbstractSurfaceType_stagedOrder: make(map[*AbstractSurfaceType]uint),
		AbstractSurfaceType_orderStaged: make(map[uint]*AbstractSurfaceType),
		AbstractSurfaceTypes_reference:  make(map[*AbstractSurfaceType]*AbstractSurfaceType),

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

		AngleType_stagedOrder: make(map[*AngleType]uint),
		AngleType_orderStaged: make(map[uint]*AngleType),
		AngleTypes_reference:  make(map[*AngleType]*AngleType),

		ArcByCenterPointType_stagedOrder: make(map[*ArcByCenterPointType]uint),
		ArcByCenterPointType_orderStaged: make(map[uint]*ArcByCenterPointType),
		ArcByCenterPointTypes_reference:  make(map[*ArcByCenterPointType]*ArcByCenterPointType),

		ArcStringType_stagedOrder: make(map[*ArcStringType]uint),
		ArcStringType_orderStaged: make(map[uint]*ArcStringType),
		ArcStringTypes_reference:  make(map[*ArcStringType]*ArcStringType),

		ArcType_stagedOrder: make(map[*ArcType]uint),
		ArcType_orderStaged: make(map[uint]*ArcType),
		ArcTypes_reference:  make(map[*ArcType]*ArcType),

		BoundingShapeType_stagedOrder: make(map[*BoundingShapeType]uint),
		BoundingShapeType_orderStaged: make(map[uint]*BoundingShapeType),
		BoundingShapeTypes_reference:  make(map[*BoundingShapeType]*BoundingShapeType),

		CircleByCenterPointType_stagedOrder: make(map[*CircleByCenterPointType]uint),
		CircleByCenterPointType_orderStaged: make(map[uint]*CircleByCenterPointType),
		CircleByCenterPointTypes_reference:  make(map[*CircleByCenterPointType]*CircleByCenterPointType),

		CircleType_stagedOrder: make(map[*CircleType]uint),
		CircleType_orderStaged: make(map[uint]*CircleType),
		CircleTypes_reference:  make(map[*CircleType]*CircleType),

		CodeType_stagedOrder: make(map[*CodeType]uint),
		CodeType_orderStaged: make(map[uint]*CodeType),
		CodeTypes_reference:  make(map[*CodeType]*CodeType),

		CodeWithAuthorityType_stagedOrder: make(map[*CodeWithAuthorityType]uint),
		CodeWithAuthorityType_orderStaged: make(map[uint]*CodeWithAuthorityType),
		CodeWithAuthorityTypes_reference:  make(map[*CodeWithAuthorityType]*CodeWithAuthorityType),

		CompositeCurveType_stagedOrder: make(map[*CompositeCurveType]uint),
		CompositeCurveType_orderStaged: make(map[uint]*CompositeCurveType),
		CompositeCurveTypes_reference:  make(map[*CompositeCurveType]*CompositeCurveType),

		CurvePropertyType_stagedOrder: make(map[*CurvePropertyType]uint),
		CurvePropertyType_orderStaged: make(map[uint]*CurvePropertyType),
		CurvePropertyTypes_reference:  make(map[*CurvePropertyType]*CurvePropertyType),

		CurveSegmentArrayPropertyType_stagedOrder: make(map[*CurveSegmentArrayPropertyType]uint),
		CurveSegmentArrayPropertyType_orderStaged: make(map[uint]*CurveSegmentArrayPropertyType),
		CurveSegmentArrayPropertyTypes_reference:  make(map[*CurveSegmentArrayPropertyType]*CurveSegmentArrayPropertyType),

		CurveType_stagedOrder: make(map[*CurveType]uint),
		CurveType_orderStaged: make(map[uint]*CurveType),
		CurveTypes_reference:  make(map[*CurveType]*CurveType),

		DirectPositionListType_stagedOrder: make(map[*DirectPositionListType]uint),
		DirectPositionListType_orderStaged: make(map[uint]*DirectPositionListType),
		DirectPositionListTypes_reference:  make(map[*DirectPositionListType]*DirectPositionListType),

		DirectPositionType_stagedOrder: make(map[*DirectPositionType]uint),
		DirectPositionType_orderStaged: make(map[uint]*DirectPositionType),
		DirectPositionTypes_reference:  make(map[*DirectPositionType]*DirectPositionType),

		DynamicFeatureType_stagedOrder: make(map[*DynamicFeatureType]uint),
		DynamicFeatureType_orderStaged: make(map[uint]*DynamicFeatureType),
		DynamicFeatureTypes_reference:  make(map[*DynamicFeatureType]*DynamicFeatureType),

		EnvelopeType_stagedOrder: make(map[*EnvelopeType]uint),
		EnvelopeType_orderStaged: make(map[uint]*EnvelopeType),
		EnvelopeTypes_reference:  make(map[*EnvelopeType]*EnvelopeType),

		GeodesicStringType_stagedOrder: make(map[*GeodesicStringType]uint),
		GeodesicStringType_orderStaged: make(map[uint]*GeodesicStringType),
		GeodesicStringTypes_reference:  make(map[*GeodesicStringType]*GeodesicStringType),

		GeodesicType_stagedOrder: make(map[*GeodesicType]uint),
		GeodesicType_orderStaged: make(map[uint]*GeodesicType),
		GeodesicTypes_reference:  make(map[*GeodesicType]*GeodesicType),

		LengthType_stagedOrder: make(map[*LengthType]uint),
		LengthType_orderStaged: make(map[uint]*LengthType),
		LengthTypes_reference:  make(map[*LengthType]*LengthType),

		LineStringSegmentType_stagedOrder: make(map[*LineStringSegmentType]uint),
		LineStringSegmentType_orderStaged: make(map[uint]*LineStringSegmentType),
		LineStringSegmentTypes_reference:  make(map[*LineStringSegmentType]*LineStringSegmentType),

		MeasureType_stagedOrder: make(map[*MeasureType]uint),
		MeasureType_orderStaged: make(map[uint]*MeasureType),
		MeasureTypes_reference:  make(map[*MeasureType]*MeasureType),

		OrientableCurveType_stagedOrder: make(map[*OrientableCurveType]uint),
		OrientableCurveType_orderStaged: make(map[uint]*OrientableCurveType),
		OrientableCurveTypes_reference:  make(map[*OrientableCurveType]*OrientableCurveType),

		PointPropertyType_stagedOrder: make(map[*PointPropertyType]uint),
		PointPropertyType_orderStaged: make(map[uint]*PointPropertyType),
		PointPropertyTypes_reference:  make(map[*PointPropertyType]*PointPropertyType),

		PointType_stagedOrder: make(map[*PointType]uint),
		PointType_orderStaged: make(map[uint]*PointType),
		PointTypes_reference:  make(map[*PointType]*PointType),

		PolygonPatchType_stagedOrder: make(map[*PolygonPatchType]uint),
		PolygonPatchType_orderStaged: make(map[uint]*PolygonPatchType),
		PolygonPatchTypes_reference:  make(map[*PolygonPatchType]*PolygonPatchType),

		RingType_stagedOrder: make(map[*RingType]uint),
		RingType_orderStaged: make(map[uint]*RingType),
		RingTypes_reference:  make(map[*RingType]*RingType),

		SurfacePatchArrayPropertyType_stagedOrder: make(map[*SurfacePatchArrayPropertyType]uint),
		SurfacePatchArrayPropertyType_orderStaged: make(map[uint]*SurfacePatchArrayPropertyType),
		SurfacePatchArrayPropertyTypes_reference:  make(map[*SurfacePatchArrayPropertyType]*SurfacePatchArrayPropertyType),

		SurfaceType_stagedOrder: make(map[*SurfaceType]uint),
		SurfaceType_orderStaged: make(map[uint]*SurfaceType),
		SurfaceTypes_reference:  make(map[*SurfaceType]*SurfaceType),

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
			"AbstractCurveSegmentType": &AbstractCurveSegmentTypeUnmarshaller{},

			"AbstractCurveType": &AbstractCurveTypeUnmarshaller{},

			"AbstractFeatureMemberType": &AbstractFeatureMemberTypeUnmarshaller{},

			"AbstractFeatureType": &AbstractFeatureTypeUnmarshaller{},

			"AbstractGMLType": &AbstractGMLTypeUnmarshaller{},

			"AbstractGeometricPrimitiveType": &AbstractGeometricPrimitiveTypeUnmarshaller{},

			"AbstractGeometryType": &AbstractGeometryTypeUnmarshaller{},

			"AbstractMetadataPropertyType": &AbstractMetadataPropertyTypeUnmarshaller{},

			"AbstractRingPropertyType": &AbstractRingPropertyTypeUnmarshaller{},

			"AbstractRingType": &AbstractRingTypeUnmarshaller{},

			"AbstractSurfacePatchType": &AbstractSurfacePatchTypeUnmarshaller{},

			"AbstractSurfaceType": &AbstractSurfaceTypeUnmarshaller{},

			"AbstractTimeGeometricPrimitiveType": &AbstractTimeGeometricPrimitiveTypeUnmarshaller{},

			"AbstractTimeObjectType": &AbstractTimeObjectTypeUnmarshaller{},

			"AbstractTimePrimitiveType": &AbstractTimePrimitiveTypeUnmarshaller{},

			"AbstractTimeSliceType": &AbstractTimeSliceTypeUnmarshaller{},

			"AngleType": &AngleTypeUnmarshaller{},

			"ArcByCenterPointType": &ArcByCenterPointTypeUnmarshaller{},

			"ArcStringType": &ArcStringTypeUnmarshaller{},

			"ArcType": &ArcTypeUnmarshaller{},

			"BoundingShapeType": &BoundingShapeTypeUnmarshaller{},

			"CircleByCenterPointType": &CircleByCenterPointTypeUnmarshaller{},

			"CircleType": &CircleTypeUnmarshaller{},

			"CodeType": &CodeTypeUnmarshaller{},

			"CodeWithAuthorityType": &CodeWithAuthorityTypeUnmarshaller{},

			"CompositeCurveType": &CompositeCurveTypeUnmarshaller{},

			"CurvePropertyType": &CurvePropertyTypeUnmarshaller{},

			"CurveSegmentArrayPropertyType": &CurveSegmentArrayPropertyTypeUnmarshaller{},

			"CurveType": &CurveTypeUnmarshaller{},

			"DirectPositionListType": &DirectPositionListTypeUnmarshaller{},

			"DirectPositionType": &DirectPositionTypeUnmarshaller{},

			"DynamicFeatureType": &DynamicFeatureTypeUnmarshaller{},

			"EnvelopeType": &EnvelopeTypeUnmarshaller{},

			"GeodesicStringType": &GeodesicStringTypeUnmarshaller{},

			"GeodesicType": &GeodesicTypeUnmarshaller{},

			"LengthType": &LengthTypeUnmarshaller{},

			"LineStringSegmentType": &LineStringSegmentTypeUnmarshaller{},

			"MeasureType": &MeasureTypeUnmarshaller{},

			"OrientableCurveType": &OrientableCurveTypeUnmarshaller{},

			"PointPropertyType": &PointPropertyTypeUnmarshaller{},

			"PointType": &PointTypeUnmarshaller{},

			"PolygonPatchType": &PolygonPatchTypeUnmarshaller{},

			"RingType": &RingTypeUnmarshaller{},

			"SurfacePatchArrayPropertyType": &SurfacePatchArrayPropertyTypeUnmarshaller{},

			"SurfaceType": &SurfaceTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractCurveSegmentType"] = len(stage.AbstractCurveSegmentTypes)
	stage.Map_GongStructName_InstancesNb["AbstractCurveType"] = len(stage.AbstractCurveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractFeatureMemberType"] = len(stage.AbstractFeatureMemberTypes)
	stage.Map_GongStructName_InstancesNb["AbstractFeatureType"] = len(stage.AbstractFeatureTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGMLType"] = len(stage.AbstractGMLTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGeometricPrimitiveType"] = len(stage.AbstractGeometricPrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractGeometryType"] = len(stage.AbstractGeometryTypes)
	stage.Map_GongStructName_InstancesNb["AbstractMetadataPropertyType"] = len(stage.AbstractMetadataPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AbstractRingPropertyType"] = len(stage.AbstractRingPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AbstractRingType"] = len(stage.AbstractRingTypes)
	stage.Map_GongStructName_InstancesNb["AbstractSurfacePatchType"] = len(stage.AbstractSurfacePatchTypes)
	stage.Map_GongStructName_InstancesNb["AbstractSurfaceType"] = len(stage.AbstractSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeGeometricPrimitiveType"] = len(stage.AbstractTimeGeometricPrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeObjectType"] = len(stage.AbstractTimeObjectTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimePrimitiveType"] = len(stage.AbstractTimePrimitiveTypes)
	stage.Map_GongStructName_InstancesNb["AbstractTimeSliceType"] = len(stage.AbstractTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["AngleType"] = len(stage.AngleTypes)
	stage.Map_GongStructName_InstancesNb["ArcByCenterPointType"] = len(stage.ArcByCenterPointTypes)
	stage.Map_GongStructName_InstancesNb["ArcStringType"] = len(stage.ArcStringTypes)
	stage.Map_GongStructName_InstancesNb["ArcType"] = len(stage.ArcTypes)
	stage.Map_GongStructName_InstancesNb["BoundingShapeType"] = len(stage.BoundingShapeTypes)
	stage.Map_GongStructName_InstancesNb["CircleByCenterPointType"] = len(stage.CircleByCenterPointTypes)
	stage.Map_GongStructName_InstancesNb["CircleType"] = len(stage.CircleTypes)
	stage.Map_GongStructName_InstancesNb["CodeType"] = len(stage.CodeTypes)
	stage.Map_GongStructName_InstancesNb["CodeWithAuthorityType"] = len(stage.CodeWithAuthorityTypes)
	stage.Map_GongStructName_InstancesNb["CompositeCurveType"] = len(stage.CompositeCurveTypes)
	stage.Map_GongStructName_InstancesNb["CurvePropertyType"] = len(stage.CurvePropertyTypes)
	stage.Map_GongStructName_InstancesNb["CurveSegmentArrayPropertyType"] = len(stage.CurveSegmentArrayPropertyTypes)
	stage.Map_GongStructName_InstancesNb["CurveType"] = len(stage.CurveTypes)
	stage.Map_GongStructName_InstancesNb["DirectPositionListType"] = len(stage.DirectPositionListTypes)
	stage.Map_GongStructName_InstancesNb["DirectPositionType"] = len(stage.DirectPositionTypes)
	stage.Map_GongStructName_InstancesNb["DynamicFeatureType"] = len(stage.DynamicFeatureTypes)
	stage.Map_GongStructName_InstancesNb["EnvelopeType"] = len(stage.EnvelopeTypes)
	stage.Map_GongStructName_InstancesNb["GeodesicStringType"] = len(stage.GeodesicStringTypes)
	stage.Map_GongStructName_InstancesNb["GeodesicType"] = len(stage.GeodesicTypes)
	stage.Map_GongStructName_InstancesNb["LengthType"] = len(stage.LengthTypes)
	stage.Map_GongStructName_InstancesNb["LineStringSegmentType"] = len(stage.LineStringSegmentTypes)
	stage.Map_GongStructName_InstancesNb["MeasureType"] = len(stage.MeasureTypes)
	stage.Map_GongStructName_InstancesNb["OrientableCurveType"] = len(stage.OrientableCurveTypes)
	stage.Map_GongStructName_InstancesNb["PointPropertyType"] = len(stage.PointPropertyTypes)
	stage.Map_GongStructName_InstancesNb["PointType"] = len(stage.PointTypes)
	stage.Map_GongStructName_InstancesNb["PolygonPatchType"] = len(stage.PolygonPatchTypes)
	stage.Map_GongStructName_InstancesNb["RingType"] = len(stage.RingTypes)
	stage.Map_GongStructName_InstancesNb["SurfacePatchArrayPropertyType"] = len(stage.SurfacePatchArrayPropertyTypes)
	stage.Map_GongStructName_InstancesNb["SurfaceType"] = len(stage.SurfaceTypes)
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
// Stage puts abstractcurvesegmenttype to the model stage
func (abstractcurvesegmenttype *AbstractCurveSegmentType) Stage(stage *Stage) *AbstractCurveSegmentType {
	__gong__stage(stage.AbstractCurveSegmentTypes, stage.AbstractCurveSegmentType_stagedOrder, stage.AbstractCurveSegmentType_orderStaged, &stage.AbstractCurveSegmentTypeOrder, stage.AbstractCurveSegmentTypes_mapString, abstractcurvesegmenttype, abstractcurvesegmenttype.Name)
	return abstractcurvesegmenttype
}

// StagePreserveOrder puts abstractcurvesegmenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractCurveSegmentTypeOrder
// - update stage.AbstractCurveSegmentTypeOrder accordingly
func (abstractcurvesegmenttype *AbstractCurveSegmentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractCurveSegmentTypes, stage.AbstractCurveSegmentType_stagedOrder, stage.AbstractCurveSegmentType_orderStaged, &stage.AbstractCurveSegmentTypeOrder, stage.AbstractCurveSegmentTypes_mapString, abstractcurvesegmenttype, order, abstractcurvesegmenttype.Name)
}

// Unstage removes abstractcurvesegmenttype off the model stage
func (abstractcurvesegmenttype *AbstractCurveSegmentType) Unstage(stage *Stage) *AbstractCurveSegmentType {
	__gong__unstage(stage.AbstractCurveSegmentTypes, stage.AbstractCurveSegmentTypes_mapString, abstractcurvesegmenttype, abstractcurvesegmenttype.Name)
	return abstractcurvesegmenttype
}

// UnstageVoid removes abstractcurvesegmenttype off the model stage
func (abstractcurvesegmenttype *AbstractCurveSegmentType) UnstageVoid(stage *Stage) {
	abstractcurvesegmenttype.Unstage(stage)
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) StageVoid(stage *Stage) {
	abstractcurvesegmenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GetName() (res string) {
	return abstractcurvesegmenttype.Name
}

// for satisfaction of GongStruct interface
func (abstractcurvesegmenttype *AbstractCurveSegmentType) SetName(name string) {
	abstractcurvesegmenttype.Name = name
}

// Stage puts abstractcurvetype to the model stage
func (abstractcurvetype *AbstractCurveType) Stage(stage *Stage) *AbstractCurveType {
	__gong__stage(stage.AbstractCurveTypes, stage.AbstractCurveType_stagedOrder, stage.AbstractCurveType_orderStaged, &stage.AbstractCurveTypeOrder, stage.AbstractCurveTypes_mapString, abstractcurvetype, abstractcurvetype.Name)
	return abstractcurvetype
}

// StagePreserveOrder puts abstractcurvetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractCurveTypeOrder
// - update stage.AbstractCurveTypeOrder accordingly
func (abstractcurvetype *AbstractCurveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractCurveTypes, stage.AbstractCurveType_stagedOrder, stage.AbstractCurveType_orderStaged, &stage.AbstractCurveTypeOrder, stage.AbstractCurveTypes_mapString, abstractcurvetype, order, abstractcurvetype.Name)
}

// Unstage removes abstractcurvetype off the model stage
func (abstractcurvetype *AbstractCurveType) Unstage(stage *Stage) *AbstractCurveType {
	__gong__unstage(stage.AbstractCurveTypes, stage.AbstractCurveTypes_mapString, abstractcurvetype, abstractcurvetype.Name)
	return abstractcurvetype
}

// UnstageVoid removes abstractcurvetype off the model stage
func (abstractcurvetype *AbstractCurveType) UnstageVoid(stage *Stage) {
	abstractcurvetype.Unstage(stage)
}

func (abstractcurvetype *AbstractCurveType) StageVoid(stage *Stage) {
	abstractcurvetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractcurvetype *AbstractCurveType) GetName() (res string) {
	return abstractcurvetype.Name
}

// for satisfaction of GongStruct interface
func (abstractcurvetype *AbstractCurveType) SetName(name string) {
	abstractcurvetype.Name = name
}

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

// Stage puts abstractringpropertytype to the model stage
func (abstractringpropertytype *AbstractRingPropertyType) Stage(stage *Stage) *AbstractRingPropertyType {
	__gong__stage(stage.AbstractRingPropertyTypes, stage.AbstractRingPropertyType_stagedOrder, stage.AbstractRingPropertyType_orderStaged, &stage.AbstractRingPropertyTypeOrder, stage.AbstractRingPropertyTypes_mapString, abstractringpropertytype, abstractringpropertytype.Name)
	return abstractringpropertytype
}

// StagePreserveOrder puts abstractringpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractRingPropertyTypeOrder
// - update stage.AbstractRingPropertyTypeOrder accordingly
func (abstractringpropertytype *AbstractRingPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractRingPropertyTypes, stage.AbstractRingPropertyType_stagedOrder, stage.AbstractRingPropertyType_orderStaged, &stage.AbstractRingPropertyTypeOrder, stage.AbstractRingPropertyTypes_mapString, abstractringpropertytype, order, abstractringpropertytype.Name)
}

// Unstage removes abstractringpropertytype off the model stage
func (abstractringpropertytype *AbstractRingPropertyType) Unstage(stage *Stage) *AbstractRingPropertyType {
	__gong__unstage(stage.AbstractRingPropertyTypes, stage.AbstractRingPropertyTypes_mapString, abstractringpropertytype, abstractringpropertytype.Name)
	return abstractringpropertytype
}

// UnstageVoid removes abstractringpropertytype off the model stage
func (abstractringpropertytype *AbstractRingPropertyType) UnstageVoid(stage *Stage) {
	abstractringpropertytype.Unstage(stage)
}

func (abstractringpropertytype *AbstractRingPropertyType) StageVoid(stage *Stage) {
	abstractringpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractringpropertytype *AbstractRingPropertyType) GetName() (res string) {
	return abstractringpropertytype.Name
}

// for satisfaction of GongStruct interface
func (abstractringpropertytype *AbstractRingPropertyType) SetName(name string) {
	abstractringpropertytype.Name = name
}

// Stage puts abstractringtype to the model stage
func (abstractringtype *AbstractRingType) Stage(stage *Stage) *AbstractRingType {
	__gong__stage(stage.AbstractRingTypes, stage.AbstractRingType_stagedOrder, stage.AbstractRingType_orderStaged, &stage.AbstractRingTypeOrder, stage.AbstractRingTypes_mapString, abstractringtype, abstractringtype.Name)
	return abstractringtype
}

// StagePreserveOrder puts abstractringtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractRingTypeOrder
// - update stage.AbstractRingTypeOrder accordingly
func (abstractringtype *AbstractRingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractRingTypes, stage.AbstractRingType_stagedOrder, stage.AbstractRingType_orderStaged, &stage.AbstractRingTypeOrder, stage.AbstractRingTypes_mapString, abstractringtype, order, abstractringtype.Name)
}

// Unstage removes abstractringtype off the model stage
func (abstractringtype *AbstractRingType) Unstage(stage *Stage) *AbstractRingType {
	__gong__unstage(stage.AbstractRingTypes, stage.AbstractRingTypes_mapString, abstractringtype, abstractringtype.Name)
	return abstractringtype
}

// UnstageVoid removes abstractringtype off the model stage
func (abstractringtype *AbstractRingType) UnstageVoid(stage *Stage) {
	abstractringtype.Unstage(stage)
}

func (abstractringtype *AbstractRingType) StageVoid(stage *Stage) {
	abstractringtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractringtype *AbstractRingType) GetName() (res string) {
	return abstractringtype.Name
}

// for satisfaction of GongStruct interface
func (abstractringtype *AbstractRingType) SetName(name string) {
	abstractringtype.Name = name
}

// Stage puts abstractsurfacepatchtype to the model stage
func (abstractsurfacepatchtype *AbstractSurfacePatchType) Stage(stage *Stage) *AbstractSurfacePatchType {
	__gong__stage(stage.AbstractSurfacePatchTypes, stage.AbstractSurfacePatchType_stagedOrder, stage.AbstractSurfacePatchType_orderStaged, &stage.AbstractSurfacePatchTypeOrder, stage.AbstractSurfacePatchTypes_mapString, abstractsurfacepatchtype, abstractsurfacepatchtype.Name)
	return abstractsurfacepatchtype
}

// StagePreserveOrder puts abstractsurfacepatchtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractSurfacePatchTypeOrder
// - update stage.AbstractSurfacePatchTypeOrder accordingly
func (abstractsurfacepatchtype *AbstractSurfacePatchType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractSurfacePatchTypes, stage.AbstractSurfacePatchType_stagedOrder, stage.AbstractSurfacePatchType_orderStaged, &stage.AbstractSurfacePatchTypeOrder, stage.AbstractSurfacePatchTypes_mapString, abstractsurfacepatchtype, order, abstractsurfacepatchtype.Name)
}

// Unstage removes abstractsurfacepatchtype off the model stage
func (abstractsurfacepatchtype *AbstractSurfacePatchType) Unstage(stage *Stage) *AbstractSurfacePatchType {
	__gong__unstage(stage.AbstractSurfacePatchTypes, stage.AbstractSurfacePatchTypes_mapString, abstractsurfacepatchtype, abstractsurfacepatchtype.Name)
	return abstractsurfacepatchtype
}

// UnstageVoid removes abstractsurfacepatchtype off the model stage
func (abstractsurfacepatchtype *AbstractSurfacePatchType) UnstageVoid(stage *Stage) {
	abstractsurfacepatchtype.Unstage(stage)
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) StageVoid(stage *Stage) {
	abstractsurfacepatchtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractsurfacepatchtype *AbstractSurfacePatchType) GetName() (res string) {
	return abstractsurfacepatchtype.Name
}

// for satisfaction of GongStruct interface
func (abstractsurfacepatchtype *AbstractSurfacePatchType) SetName(name string) {
	abstractsurfacepatchtype.Name = name
}

// Stage puts abstractsurfacetype to the model stage
func (abstractsurfacetype *AbstractSurfaceType) Stage(stage *Stage) *AbstractSurfaceType {
	__gong__stage(stage.AbstractSurfaceTypes, stage.AbstractSurfaceType_stagedOrder, stage.AbstractSurfaceType_orderStaged, &stage.AbstractSurfaceTypeOrder, stage.AbstractSurfaceTypes_mapString, abstractsurfacetype, abstractsurfacetype.Name)
	return abstractsurfacetype
}

// StagePreserveOrder puts abstractsurfacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractSurfaceTypeOrder
// - update stage.AbstractSurfaceTypeOrder accordingly
func (abstractsurfacetype *AbstractSurfaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractSurfaceTypes, stage.AbstractSurfaceType_stagedOrder, stage.AbstractSurfaceType_orderStaged, &stage.AbstractSurfaceTypeOrder, stage.AbstractSurfaceTypes_mapString, abstractsurfacetype, order, abstractsurfacetype.Name)
}

// Unstage removes abstractsurfacetype off the model stage
func (abstractsurfacetype *AbstractSurfaceType) Unstage(stage *Stage) *AbstractSurfaceType {
	__gong__unstage(stage.AbstractSurfaceTypes, stage.AbstractSurfaceTypes_mapString, abstractsurfacetype, abstractsurfacetype.Name)
	return abstractsurfacetype
}

// UnstageVoid removes abstractsurfacetype off the model stage
func (abstractsurfacetype *AbstractSurfaceType) UnstageVoid(stage *Stage) {
	abstractsurfacetype.Unstage(stage)
}

func (abstractsurfacetype *AbstractSurfaceType) StageVoid(stage *Stage) {
	abstractsurfacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractsurfacetype *AbstractSurfaceType) GetName() (res string) {
	return abstractsurfacetype.Name
}

// for satisfaction of GongStruct interface
func (abstractsurfacetype *AbstractSurfaceType) SetName(name string) {
	abstractsurfacetype.Name = name
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

// Stage puts arcbycenterpointtype to the model stage
func (arcbycenterpointtype *ArcByCenterPointType) Stage(stage *Stage) *ArcByCenterPointType {
	__gong__stage(stage.ArcByCenterPointTypes, stage.ArcByCenterPointType_stagedOrder, stage.ArcByCenterPointType_orderStaged, &stage.ArcByCenterPointTypeOrder, stage.ArcByCenterPointTypes_mapString, arcbycenterpointtype, arcbycenterpointtype.Name)
	return arcbycenterpointtype
}

// StagePreserveOrder puts arcbycenterpointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArcByCenterPointTypeOrder
// - update stage.ArcByCenterPointTypeOrder accordingly
func (arcbycenterpointtype *ArcByCenterPointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArcByCenterPointTypes, stage.ArcByCenterPointType_stagedOrder, stage.ArcByCenterPointType_orderStaged, &stage.ArcByCenterPointTypeOrder, stage.ArcByCenterPointTypes_mapString, arcbycenterpointtype, order, arcbycenterpointtype.Name)
}

// Unstage removes arcbycenterpointtype off the model stage
func (arcbycenterpointtype *ArcByCenterPointType) Unstage(stage *Stage) *ArcByCenterPointType {
	__gong__unstage(stage.ArcByCenterPointTypes, stage.ArcByCenterPointTypes_mapString, arcbycenterpointtype, arcbycenterpointtype.Name)
	return arcbycenterpointtype
}

// UnstageVoid removes arcbycenterpointtype off the model stage
func (arcbycenterpointtype *ArcByCenterPointType) UnstageVoid(stage *Stage) {
	arcbycenterpointtype.Unstage(stage)
}

func (arcbycenterpointtype *ArcByCenterPointType) StageVoid(stage *Stage) {
	arcbycenterpointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arcbycenterpointtype *ArcByCenterPointType) GetName() (res string) {
	return arcbycenterpointtype.Name
}

// for satisfaction of GongStruct interface
func (arcbycenterpointtype *ArcByCenterPointType) SetName(name string) {
	arcbycenterpointtype.Name = name
}

// Stage puts arcstringtype to the model stage
func (arcstringtype *ArcStringType) Stage(stage *Stage) *ArcStringType {
	__gong__stage(stage.ArcStringTypes, stage.ArcStringType_stagedOrder, stage.ArcStringType_orderStaged, &stage.ArcStringTypeOrder, stage.ArcStringTypes_mapString, arcstringtype, arcstringtype.Name)
	return arcstringtype
}

// StagePreserveOrder puts arcstringtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArcStringTypeOrder
// - update stage.ArcStringTypeOrder accordingly
func (arcstringtype *ArcStringType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ArcStringTypes, stage.ArcStringType_stagedOrder, stage.ArcStringType_orderStaged, &stage.ArcStringTypeOrder, stage.ArcStringTypes_mapString, arcstringtype, order, arcstringtype.Name)
}

// Unstage removes arcstringtype off the model stage
func (arcstringtype *ArcStringType) Unstage(stage *Stage) *ArcStringType {
	__gong__unstage(stage.ArcStringTypes, stage.ArcStringTypes_mapString, arcstringtype, arcstringtype.Name)
	return arcstringtype
}

// UnstageVoid removes arcstringtype off the model stage
func (arcstringtype *ArcStringType) UnstageVoid(stage *Stage) {
	arcstringtype.Unstage(stage)
}

func (arcstringtype *ArcStringType) StageVoid(stage *Stage) {
	arcstringtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arcstringtype *ArcStringType) GetName() (res string) {
	return arcstringtype.Name
}

// for satisfaction of GongStruct interface
func (arcstringtype *ArcStringType) SetName(name string) {
	arcstringtype.Name = name
}

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

// Stage puts circlebycenterpointtype to the model stage
func (circlebycenterpointtype *CircleByCenterPointType) Stage(stage *Stage) *CircleByCenterPointType {
	__gong__stage(stage.CircleByCenterPointTypes, stage.CircleByCenterPointType_stagedOrder, stage.CircleByCenterPointType_orderStaged, &stage.CircleByCenterPointTypeOrder, stage.CircleByCenterPointTypes_mapString, circlebycenterpointtype, circlebycenterpointtype.Name)
	return circlebycenterpointtype
}

// StagePreserveOrder puts circlebycenterpointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleByCenterPointTypeOrder
// - update stage.CircleByCenterPointTypeOrder accordingly
func (circlebycenterpointtype *CircleByCenterPointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CircleByCenterPointTypes, stage.CircleByCenterPointType_stagedOrder, stage.CircleByCenterPointType_orderStaged, &stage.CircleByCenterPointTypeOrder, stage.CircleByCenterPointTypes_mapString, circlebycenterpointtype, order, circlebycenterpointtype.Name)
}

// Unstage removes circlebycenterpointtype off the model stage
func (circlebycenterpointtype *CircleByCenterPointType) Unstage(stage *Stage) *CircleByCenterPointType {
	__gong__unstage(stage.CircleByCenterPointTypes, stage.CircleByCenterPointTypes_mapString, circlebycenterpointtype, circlebycenterpointtype.Name)
	return circlebycenterpointtype
}

// UnstageVoid removes circlebycenterpointtype off the model stage
func (circlebycenterpointtype *CircleByCenterPointType) UnstageVoid(stage *Stage) {
	circlebycenterpointtype.Unstage(stage)
}

func (circlebycenterpointtype *CircleByCenterPointType) StageVoid(stage *Stage) {
	circlebycenterpointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (circlebycenterpointtype *CircleByCenterPointType) GetName() (res string) {
	return circlebycenterpointtype.Name
}

// for satisfaction of GongStruct interface
func (circlebycenterpointtype *CircleByCenterPointType) SetName(name string) {
	circlebycenterpointtype.Name = name
}

// Stage puts circletype to the model stage
func (circletype *CircleType) Stage(stage *Stage) *CircleType {
	__gong__stage(stage.CircleTypes, stage.CircleType_stagedOrder, stage.CircleType_orderStaged, &stage.CircleTypeOrder, stage.CircleTypes_mapString, circletype, circletype.Name)
	return circletype
}

// StagePreserveOrder puts circletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CircleTypeOrder
// - update stage.CircleTypeOrder accordingly
func (circletype *CircleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CircleTypes, stage.CircleType_stagedOrder, stage.CircleType_orderStaged, &stage.CircleTypeOrder, stage.CircleTypes_mapString, circletype, order, circletype.Name)
}

// Unstage removes circletype off the model stage
func (circletype *CircleType) Unstage(stage *Stage) *CircleType {
	__gong__unstage(stage.CircleTypes, stage.CircleTypes_mapString, circletype, circletype.Name)
	return circletype
}

// UnstageVoid removes circletype off the model stage
func (circletype *CircleType) UnstageVoid(stage *Stage) {
	circletype.Unstage(stage)
}

func (circletype *CircleType) StageVoid(stage *Stage) {
	circletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (circletype *CircleType) GetName() (res string) {
	return circletype.Name
}

// for satisfaction of GongStruct interface
func (circletype *CircleType) SetName(name string) {
	circletype.Name = name
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

// Stage puts compositecurvetype to the model stage
func (compositecurvetype *CompositeCurveType) Stage(stage *Stage) *CompositeCurveType {
	__gong__stage(stage.CompositeCurveTypes, stage.CompositeCurveType_stagedOrder, stage.CompositeCurveType_orderStaged, &stage.CompositeCurveTypeOrder, stage.CompositeCurveTypes_mapString, compositecurvetype, compositecurvetype.Name)
	return compositecurvetype
}

// StagePreserveOrder puts compositecurvetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CompositeCurveTypeOrder
// - update stage.CompositeCurveTypeOrder accordingly
func (compositecurvetype *CompositeCurveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CompositeCurveTypes, stage.CompositeCurveType_stagedOrder, stage.CompositeCurveType_orderStaged, &stage.CompositeCurveTypeOrder, stage.CompositeCurveTypes_mapString, compositecurvetype, order, compositecurvetype.Name)
}

// Unstage removes compositecurvetype off the model stage
func (compositecurvetype *CompositeCurveType) Unstage(stage *Stage) *CompositeCurveType {
	__gong__unstage(stage.CompositeCurveTypes, stage.CompositeCurveTypes_mapString, compositecurvetype, compositecurvetype.Name)
	return compositecurvetype
}

// UnstageVoid removes compositecurvetype off the model stage
func (compositecurvetype *CompositeCurveType) UnstageVoid(stage *Stage) {
	compositecurvetype.Unstage(stage)
}

func (compositecurvetype *CompositeCurveType) StageVoid(stage *Stage) {
	compositecurvetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (compositecurvetype *CompositeCurveType) GetName() (res string) {
	return compositecurvetype.Name
}

// for satisfaction of GongStruct interface
func (compositecurvetype *CompositeCurveType) SetName(name string) {
	compositecurvetype.Name = name
}

// Stage puts curvepropertytype to the model stage
func (curvepropertytype *CurvePropertyType) Stage(stage *Stage) *CurvePropertyType {
	__gong__stage(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder, stage.CurvePropertyType_orderStaged, &stage.CurvePropertyTypeOrder, stage.CurvePropertyTypes_mapString, curvepropertytype, curvepropertytype.Name)
	return curvepropertytype
}

// StagePreserveOrder puts curvepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CurvePropertyTypeOrder
// - update stage.CurvePropertyTypeOrder accordingly
func (curvepropertytype *CurvePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder, stage.CurvePropertyType_orderStaged, &stage.CurvePropertyTypeOrder, stage.CurvePropertyTypes_mapString, curvepropertytype, order, curvepropertytype.Name)
}

// Unstage removes curvepropertytype off the model stage
func (curvepropertytype *CurvePropertyType) Unstage(stage *Stage) *CurvePropertyType {
	__gong__unstage(stage.CurvePropertyTypes, stage.CurvePropertyTypes_mapString, curvepropertytype, curvepropertytype.Name)
	return curvepropertytype
}

// UnstageVoid removes curvepropertytype off the model stage
func (curvepropertytype *CurvePropertyType) UnstageVoid(stage *Stage) {
	curvepropertytype.Unstage(stage)
}

func (curvepropertytype *CurvePropertyType) StageVoid(stage *Stage) {
	curvepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (curvepropertytype *CurvePropertyType) GetName() (res string) {
	return curvepropertytype.Name
}

// for satisfaction of GongStruct interface
func (curvepropertytype *CurvePropertyType) SetName(name string) {
	curvepropertytype.Name = name
}

// Stage puts curvesegmentarraypropertytype to the model stage
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) Stage(stage *Stage) *CurveSegmentArrayPropertyType {
	__gong__stage(stage.CurveSegmentArrayPropertyTypes, stage.CurveSegmentArrayPropertyType_stagedOrder, stage.CurveSegmentArrayPropertyType_orderStaged, &stage.CurveSegmentArrayPropertyTypeOrder, stage.CurveSegmentArrayPropertyTypes_mapString, curvesegmentarraypropertytype, curvesegmentarraypropertytype.Name)
	return curvesegmentarraypropertytype
}

// StagePreserveOrder puts curvesegmentarraypropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CurveSegmentArrayPropertyTypeOrder
// - update stage.CurveSegmentArrayPropertyTypeOrder accordingly
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CurveSegmentArrayPropertyTypes, stage.CurveSegmentArrayPropertyType_stagedOrder, stage.CurveSegmentArrayPropertyType_orderStaged, &stage.CurveSegmentArrayPropertyTypeOrder, stage.CurveSegmentArrayPropertyTypes_mapString, curvesegmentarraypropertytype, order, curvesegmentarraypropertytype.Name)
}

// Unstage removes curvesegmentarraypropertytype off the model stage
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) Unstage(stage *Stage) *CurveSegmentArrayPropertyType {
	__gong__unstage(stage.CurveSegmentArrayPropertyTypes, stage.CurveSegmentArrayPropertyTypes_mapString, curvesegmentarraypropertytype, curvesegmentarraypropertytype.Name)
	return curvesegmentarraypropertytype
}

// UnstageVoid removes curvesegmentarraypropertytype off the model stage
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) UnstageVoid(stage *Stage) {
	curvesegmentarraypropertytype.Unstage(stage)
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) StageVoid(stage *Stage) {
	curvesegmentarraypropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GetName() (res string) {
	return curvesegmentarraypropertytype.Name
}

// for satisfaction of GongStruct interface
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) SetName(name string) {
	curvesegmentarraypropertytype.Name = name
}

// Stage puts curvetype to the model stage
func (curvetype *CurveType) Stage(stage *Stage) *CurveType {
	__gong__stage(stage.CurveTypes, stage.CurveType_stagedOrder, stage.CurveType_orderStaged, &stage.CurveTypeOrder, stage.CurveTypes_mapString, curvetype, curvetype.Name)
	return curvetype
}

// StagePreserveOrder puts curvetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CurveTypeOrder
// - update stage.CurveTypeOrder accordingly
func (curvetype *CurveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CurveTypes, stage.CurveType_stagedOrder, stage.CurveType_orderStaged, &stage.CurveTypeOrder, stage.CurveTypes_mapString, curvetype, order, curvetype.Name)
}

// Unstage removes curvetype off the model stage
func (curvetype *CurveType) Unstage(stage *Stage) *CurveType {
	__gong__unstage(stage.CurveTypes, stage.CurveTypes_mapString, curvetype, curvetype.Name)
	return curvetype
}

// UnstageVoid removes curvetype off the model stage
func (curvetype *CurveType) UnstageVoid(stage *Stage) {
	curvetype.Unstage(stage)
}

func (curvetype *CurveType) StageVoid(stage *Stage) {
	curvetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (curvetype *CurveType) GetName() (res string) {
	return curvetype.Name
}

// for satisfaction of GongStruct interface
func (curvetype *CurveType) SetName(name string) {
	curvetype.Name = name
}

// Stage puts directpositionlisttype to the model stage
func (directpositionlisttype *DirectPositionListType) Stage(stage *Stage) *DirectPositionListType {
	__gong__stage(stage.DirectPositionListTypes, stage.DirectPositionListType_stagedOrder, stage.DirectPositionListType_orderStaged, &stage.DirectPositionListTypeOrder, stage.DirectPositionListTypes_mapString, directpositionlisttype, directpositionlisttype.Name)
	return directpositionlisttype
}

// StagePreserveOrder puts directpositionlisttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectPositionListTypeOrder
// - update stage.DirectPositionListTypeOrder accordingly
func (directpositionlisttype *DirectPositionListType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DirectPositionListTypes, stage.DirectPositionListType_stagedOrder, stage.DirectPositionListType_orderStaged, &stage.DirectPositionListTypeOrder, stage.DirectPositionListTypes_mapString, directpositionlisttype, order, directpositionlisttype.Name)
}

// Unstage removes directpositionlisttype off the model stage
func (directpositionlisttype *DirectPositionListType) Unstage(stage *Stage) *DirectPositionListType {
	__gong__unstage(stage.DirectPositionListTypes, stage.DirectPositionListTypes_mapString, directpositionlisttype, directpositionlisttype.Name)
	return directpositionlisttype
}

// UnstageVoid removes directpositionlisttype off the model stage
func (directpositionlisttype *DirectPositionListType) UnstageVoid(stage *Stage) {
	directpositionlisttype.Unstage(stage)
}

func (directpositionlisttype *DirectPositionListType) StageVoid(stage *Stage) {
	directpositionlisttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (directpositionlisttype *DirectPositionListType) GetName() (res string) {
	return directpositionlisttype.Name
}

// for satisfaction of GongStruct interface
func (directpositionlisttype *DirectPositionListType) SetName(name string) {
	directpositionlisttype.Name = name
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

// Stage puts geodesicstringtype to the model stage
func (geodesicstringtype *GeodesicStringType) Stage(stage *Stage) *GeodesicStringType {
	__gong__stage(stage.GeodesicStringTypes, stage.GeodesicStringType_stagedOrder, stage.GeodesicStringType_orderStaged, &stage.GeodesicStringTypeOrder, stage.GeodesicStringTypes_mapString, geodesicstringtype, geodesicstringtype.Name)
	return geodesicstringtype
}

// StagePreserveOrder puts geodesicstringtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeodesicStringTypeOrder
// - update stage.GeodesicStringTypeOrder accordingly
func (geodesicstringtype *GeodesicStringType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeodesicStringTypes, stage.GeodesicStringType_stagedOrder, stage.GeodesicStringType_orderStaged, &stage.GeodesicStringTypeOrder, stage.GeodesicStringTypes_mapString, geodesicstringtype, order, geodesicstringtype.Name)
}

// Unstage removes geodesicstringtype off the model stage
func (geodesicstringtype *GeodesicStringType) Unstage(stage *Stage) *GeodesicStringType {
	__gong__unstage(stage.GeodesicStringTypes, stage.GeodesicStringTypes_mapString, geodesicstringtype, geodesicstringtype.Name)
	return geodesicstringtype
}

// UnstageVoid removes geodesicstringtype off the model stage
func (geodesicstringtype *GeodesicStringType) UnstageVoid(stage *Stage) {
	geodesicstringtype.Unstage(stage)
}

func (geodesicstringtype *GeodesicStringType) StageVoid(stage *Stage) {
	geodesicstringtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geodesicstringtype *GeodesicStringType) GetName() (res string) {
	return geodesicstringtype.Name
}

// for satisfaction of GongStruct interface
func (geodesicstringtype *GeodesicStringType) SetName(name string) {
	geodesicstringtype.Name = name
}

// Stage puts geodesictype to the model stage
func (geodesictype *GeodesicType) Stage(stage *Stage) *GeodesicType {
	__gong__stage(stage.GeodesicTypes, stage.GeodesicType_stagedOrder, stage.GeodesicType_orderStaged, &stage.GeodesicTypeOrder, stage.GeodesicTypes_mapString, geodesictype, geodesictype.Name)
	return geodesictype
}

// StagePreserveOrder puts geodesictype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeodesicTypeOrder
// - update stage.GeodesicTypeOrder accordingly
func (geodesictype *GeodesicType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeodesicTypes, stage.GeodesicType_stagedOrder, stage.GeodesicType_orderStaged, &stage.GeodesicTypeOrder, stage.GeodesicTypes_mapString, geodesictype, order, geodesictype.Name)
}

// Unstage removes geodesictype off the model stage
func (geodesictype *GeodesicType) Unstage(stage *Stage) *GeodesicType {
	__gong__unstage(stage.GeodesicTypes, stage.GeodesicTypes_mapString, geodesictype, geodesictype.Name)
	return geodesictype
}

// UnstageVoid removes geodesictype off the model stage
func (geodesictype *GeodesicType) UnstageVoid(stage *Stage) {
	geodesictype.Unstage(stage)
}

func (geodesictype *GeodesicType) StageVoid(stage *Stage) {
	geodesictype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geodesictype *GeodesicType) GetName() (res string) {
	return geodesictype.Name
}

// for satisfaction of GongStruct interface
func (geodesictype *GeodesicType) SetName(name string) {
	geodesictype.Name = name
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

// Stage puts linestringsegmenttype to the model stage
func (linestringsegmenttype *LineStringSegmentType) Stage(stage *Stage) *LineStringSegmentType {
	__gong__stage(stage.LineStringSegmentTypes, stage.LineStringSegmentType_stagedOrder, stage.LineStringSegmentType_orderStaged, &stage.LineStringSegmentTypeOrder, stage.LineStringSegmentTypes_mapString, linestringsegmenttype, linestringsegmenttype.Name)
	return linestringsegmenttype
}

// StagePreserveOrder puts linestringsegmenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LineStringSegmentTypeOrder
// - update stage.LineStringSegmentTypeOrder accordingly
func (linestringsegmenttype *LineStringSegmentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LineStringSegmentTypes, stage.LineStringSegmentType_stagedOrder, stage.LineStringSegmentType_orderStaged, &stage.LineStringSegmentTypeOrder, stage.LineStringSegmentTypes_mapString, linestringsegmenttype, order, linestringsegmenttype.Name)
}

// Unstage removes linestringsegmenttype off the model stage
func (linestringsegmenttype *LineStringSegmentType) Unstage(stage *Stage) *LineStringSegmentType {
	__gong__unstage(stage.LineStringSegmentTypes, stage.LineStringSegmentTypes_mapString, linestringsegmenttype, linestringsegmenttype.Name)
	return linestringsegmenttype
}

// UnstageVoid removes linestringsegmenttype off the model stage
func (linestringsegmenttype *LineStringSegmentType) UnstageVoid(stage *Stage) {
	linestringsegmenttype.Unstage(stage)
}

func (linestringsegmenttype *LineStringSegmentType) StageVoid(stage *Stage) {
	linestringsegmenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linestringsegmenttype *LineStringSegmentType) GetName() (res string) {
	return linestringsegmenttype.Name
}

// for satisfaction of GongStruct interface
func (linestringsegmenttype *LineStringSegmentType) SetName(name string) {
	linestringsegmenttype.Name = name
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

// Stage puts orientablecurvetype to the model stage
func (orientablecurvetype *OrientableCurveType) Stage(stage *Stage) *OrientableCurveType {
	__gong__stage(stage.OrientableCurveTypes, stage.OrientableCurveType_stagedOrder, stage.OrientableCurveType_orderStaged, &stage.OrientableCurveTypeOrder, stage.OrientableCurveTypes_mapString, orientablecurvetype, orientablecurvetype.Name)
	return orientablecurvetype
}

// StagePreserveOrder puts orientablecurvetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OrientableCurveTypeOrder
// - update stage.OrientableCurveTypeOrder accordingly
func (orientablecurvetype *OrientableCurveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OrientableCurveTypes, stage.OrientableCurveType_stagedOrder, stage.OrientableCurveType_orderStaged, &stage.OrientableCurveTypeOrder, stage.OrientableCurveTypes_mapString, orientablecurvetype, order, orientablecurvetype.Name)
}

// Unstage removes orientablecurvetype off the model stage
func (orientablecurvetype *OrientableCurveType) Unstage(stage *Stage) *OrientableCurveType {
	__gong__unstage(stage.OrientableCurveTypes, stage.OrientableCurveTypes_mapString, orientablecurvetype, orientablecurvetype.Name)
	return orientablecurvetype
}

// UnstageVoid removes orientablecurvetype off the model stage
func (orientablecurvetype *OrientableCurveType) UnstageVoid(stage *Stage) {
	orientablecurvetype.Unstage(stage)
}

func (orientablecurvetype *OrientableCurveType) StageVoid(stage *Stage) {
	orientablecurvetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (orientablecurvetype *OrientableCurveType) GetName() (res string) {
	return orientablecurvetype.Name
}

// for satisfaction of GongStruct interface
func (orientablecurvetype *OrientableCurveType) SetName(name string) {
	orientablecurvetype.Name = name
}

// Stage puts pointpropertytype to the model stage
func (pointpropertytype *PointPropertyType) Stage(stage *Stage) *PointPropertyType {
	__gong__stage(stage.PointPropertyTypes, stage.PointPropertyType_stagedOrder, stage.PointPropertyType_orderStaged, &stage.PointPropertyTypeOrder, stage.PointPropertyTypes_mapString, pointpropertytype, pointpropertytype.Name)
	return pointpropertytype
}

// StagePreserveOrder puts pointpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PointPropertyTypeOrder
// - update stage.PointPropertyTypeOrder accordingly
func (pointpropertytype *PointPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PointPropertyTypes, stage.PointPropertyType_stagedOrder, stage.PointPropertyType_orderStaged, &stage.PointPropertyTypeOrder, stage.PointPropertyTypes_mapString, pointpropertytype, order, pointpropertytype.Name)
}

// Unstage removes pointpropertytype off the model stage
func (pointpropertytype *PointPropertyType) Unstage(stage *Stage) *PointPropertyType {
	__gong__unstage(stage.PointPropertyTypes, stage.PointPropertyTypes_mapString, pointpropertytype, pointpropertytype.Name)
	return pointpropertytype
}

// UnstageVoid removes pointpropertytype off the model stage
func (pointpropertytype *PointPropertyType) UnstageVoid(stage *Stage) {
	pointpropertytype.Unstage(stage)
}

func (pointpropertytype *PointPropertyType) StageVoid(stage *Stage) {
	pointpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pointpropertytype *PointPropertyType) GetName() (res string) {
	return pointpropertytype.Name
}

// for satisfaction of GongStruct interface
func (pointpropertytype *PointPropertyType) SetName(name string) {
	pointpropertytype.Name = name
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

// Stage puts polygonpatchtype to the model stage
func (polygonpatchtype *PolygonPatchType) Stage(stage *Stage) *PolygonPatchType {
	__gong__stage(stage.PolygonPatchTypes, stage.PolygonPatchType_stagedOrder, stage.PolygonPatchType_orderStaged, &stage.PolygonPatchTypeOrder, stage.PolygonPatchTypes_mapString, polygonpatchtype, polygonpatchtype.Name)
	return polygonpatchtype
}

// StagePreserveOrder puts polygonpatchtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PolygonPatchTypeOrder
// - update stage.PolygonPatchTypeOrder accordingly
func (polygonpatchtype *PolygonPatchType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PolygonPatchTypes, stage.PolygonPatchType_stagedOrder, stage.PolygonPatchType_orderStaged, &stage.PolygonPatchTypeOrder, stage.PolygonPatchTypes_mapString, polygonpatchtype, order, polygonpatchtype.Name)
}

// Unstage removes polygonpatchtype off the model stage
func (polygonpatchtype *PolygonPatchType) Unstage(stage *Stage) *PolygonPatchType {
	__gong__unstage(stage.PolygonPatchTypes, stage.PolygonPatchTypes_mapString, polygonpatchtype, polygonpatchtype.Name)
	return polygonpatchtype
}

// UnstageVoid removes polygonpatchtype off the model stage
func (polygonpatchtype *PolygonPatchType) UnstageVoid(stage *Stage) {
	polygonpatchtype.Unstage(stage)
}

func (polygonpatchtype *PolygonPatchType) StageVoid(stage *Stage) {
	polygonpatchtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (polygonpatchtype *PolygonPatchType) GetName() (res string) {
	return polygonpatchtype.Name
}

// for satisfaction of GongStruct interface
func (polygonpatchtype *PolygonPatchType) SetName(name string) {
	polygonpatchtype.Name = name
}

// Stage puts ringtype to the model stage
func (ringtype *RingType) Stage(stage *Stage) *RingType {
	__gong__stage(stage.RingTypes, stage.RingType_stagedOrder, stage.RingType_orderStaged, &stage.RingTypeOrder, stage.RingTypes_mapString, ringtype, ringtype.Name)
	return ringtype
}

// StagePreserveOrder puts ringtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RingTypeOrder
// - update stage.RingTypeOrder accordingly
func (ringtype *RingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RingTypes, stage.RingType_stagedOrder, stage.RingType_orderStaged, &stage.RingTypeOrder, stage.RingTypes_mapString, ringtype, order, ringtype.Name)
}

// Unstage removes ringtype off the model stage
func (ringtype *RingType) Unstage(stage *Stage) *RingType {
	__gong__unstage(stage.RingTypes, stage.RingTypes_mapString, ringtype, ringtype.Name)
	return ringtype
}

// UnstageVoid removes ringtype off the model stage
func (ringtype *RingType) UnstageVoid(stage *Stage) {
	ringtype.Unstage(stage)
}

func (ringtype *RingType) StageVoid(stage *Stage) {
	ringtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ringtype *RingType) GetName() (res string) {
	return ringtype.Name
}

// for satisfaction of GongStruct interface
func (ringtype *RingType) SetName(name string) {
	ringtype.Name = name
}

// Stage puts surfacepatcharraypropertytype to the model stage
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) Stage(stage *Stage) *SurfacePatchArrayPropertyType {
	__gong__stage(stage.SurfacePatchArrayPropertyTypes, stage.SurfacePatchArrayPropertyType_stagedOrder, stage.SurfacePatchArrayPropertyType_orderStaged, &stage.SurfacePatchArrayPropertyTypeOrder, stage.SurfacePatchArrayPropertyTypes_mapString, surfacepatcharraypropertytype, surfacepatcharraypropertytype.Name)
	return surfacepatcharraypropertytype
}

// StagePreserveOrder puts surfacepatcharraypropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SurfacePatchArrayPropertyTypeOrder
// - update stage.SurfacePatchArrayPropertyTypeOrder accordingly
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SurfacePatchArrayPropertyTypes, stage.SurfacePatchArrayPropertyType_stagedOrder, stage.SurfacePatchArrayPropertyType_orderStaged, &stage.SurfacePatchArrayPropertyTypeOrder, stage.SurfacePatchArrayPropertyTypes_mapString, surfacepatcharraypropertytype, order, surfacepatcharraypropertytype.Name)
}

// Unstage removes surfacepatcharraypropertytype off the model stage
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) Unstage(stage *Stage) *SurfacePatchArrayPropertyType {
	__gong__unstage(stage.SurfacePatchArrayPropertyTypes, stage.SurfacePatchArrayPropertyTypes_mapString, surfacepatcharraypropertytype, surfacepatcharraypropertytype.Name)
	return surfacepatcharraypropertytype
}

// UnstageVoid removes surfacepatcharraypropertytype off the model stage
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) UnstageVoid(stage *Stage) {
	surfacepatcharraypropertytype.Unstage(stage)
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) StageVoid(stage *Stage) {
	surfacepatcharraypropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GetName() (res string) {
	return surfacepatcharraypropertytype.Name
}

// for satisfaction of GongStruct interface
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) SetName(name string) {
	surfacepatcharraypropertytype.Name = name
}

// Stage puts surfacetype to the model stage
func (surfacetype *SurfaceType) Stage(stage *Stage) *SurfaceType {
	__gong__stage(stage.SurfaceTypes, stage.SurfaceType_stagedOrder, stage.SurfaceType_orderStaged, &stage.SurfaceTypeOrder, stage.SurfaceTypes_mapString, surfacetype, surfacetype.Name)
	return surfacetype
}

// StagePreserveOrder puts surfacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SurfaceTypeOrder
// - update stage.SurfaceTypeOrder accordingly
func (surfacetype *SurfaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SurfaceTypes, stage.SurfaceType_stagedOrder, stage.SurfaceType_orderStaged, &stage.SurfaceTypeOrder, stage.SurfaceTypes_mapString, surfacetype, order, surfacetype.Name)
}

// Unstage removes surfacetype off the model stage
func (surfacetype *SurfaceType) Unstage(stage *Stage) *SurfaceType {
	__gong__unstage(stage.SurfaceTypes, stage.SurfaceTypes_mapString, surfacetype, surfacetype.Name)
	return surfacetype
}

// UnstageVoid removes surfacetype off the model stage
func (surfacetype *SurfaceType) UnstageVoid(stage *Stage) {
	surfacetype.Unstage(stage)
}

func (surfacetype *SurfaceType) StageVoid(stage *Stage) {
	surfacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (surfacetype *SurfaceType) GetName() (res string) {
	return surfacetype.Name
}

// for satisfaction of GongStruct interface
func (surfacetype *SurfaceType) SetName(name string) {
	surfacetype.Name = name
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
	__gong__resetStageType(&stage.AbstractCurveSegmentTypes, &stage.AbstractCurveSegmentTypes_mapString, &stage.AbstractCurveSegmentType_stagedOrder, &stage.AbstractCurveSegmentTypeOrder)

	__gong__resetStageType(&stage.AbstractCurveTypes, &stage.AbstractCurveTypes_mapString, &stage.AbstractCurveType_stagedOrder, &stage.AbstractCurveTypeOrder)

	__gong__resetStageType(&stage.AbstractFeatureMemberTypes, &stage.AbstractFeatureMemberTypes_mapString, &stage.AbstractFeatureMemberType_stagedOrder, &stage.AbstractFeatureMemberTypeOrder)

	__gong__resetStageType(&stage.AbstractFeatureTypes, &stage.AbstractFeatureTypes_mapString, &stage.AbstractFeatureType_stagedOrder, &stage.AbstractFeatureTypeOrder)

	__gong__resetStageType(&stage.AbstractGMLTypes, &stage.AbstractGMLTypes_mapString, &stage.AbstractGMLType_stagedOrder, &stage.AbstractGMLTypeOrder)

	__gong__resetStageType(&stage.AbstractGeometricPrimitiveTypes, &stage.AbstractGeometricPrimitiveTypes_mapString, &stage.AbstractGeometricPrimitiveType_stagedOrder, &stage.AbstractGeometricPrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractGeometryTypes, &stage.AbstractGeometryTypes_mapString, &stage.AbstractGeometryType_stagedOrder, &stage.AbstractGeometryTypeOrder)

	__gong__resetStageType(&stage.AbstractMetadataPropertyTypes, &stage.AbstractMetadataPropertyTypes_mapString, &stage.AbstractMetadataPropertyType_stagedOrder, &stage.AbstractMetadataPropertyTypeOrder)

	__gong__resetStageType(&stage.AbstractRingPropertyTypes, &stage.AbstractRingPropertyTypes_mapString, &stage.AbstractRingPropertyType_stagedOrder, &stage.AbstractRingPropertyTypeOrder)

	__gong__resetStageType(&stage.AbstractRingTypes, &stage.AbstractRingTypes_mapString, &stage.AbstractRingType_stagedOrder, &stage.AbstractRingTypeOrder)

	__gong__resetStageType(&stage.AbstractSurfacePatchTypes, &stage.AbstractSurfacePatchTypes_mapString, &stage.AbstractSurfacePatchType_stagedOrder, &stage.AbstractSurfacePatchTypeOrder)

	__gong__resetStageType(&stage.AbstractSurfaceTypes, &stage.AbstractSurfaceTypes_mapString, &stage.AbstractSurfaceType_stagedOrder, &stage.AbstractSurfaceTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeGeometricPrimitiveTypes, &stage.AbstractTimeGeometricPrimitiveTypes_mapString, &stage.AbstractTimeGeometricPrimitiveType_stagedOrder, &stage.AbstractTimeGeometricPrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeObjectTypes, &stage.AbstractTimeObjectTypes_mapString, &stage.AbstractTimeObjectType_stagedOrder, &stage.AbstractTimeObjectTypeOrder)

	__gong__resetStageType(&stage.AbstractTimePrimitiveTypes, &stage.AbstractTimePrimitiveTypes_mapString, &stage.AbstractTimePrimitiveType_stagedOrder, &stage.AbstractTimePrimitiveTypeOrder)

	__gong__resetStageType(&stage.AbstractTimeSliceTypes, &stage.AbstractTimeSliceTypes_mapString, &stage.AbstractTimeSliceType_stagedOrder, &stage.AbstractTimeSliceTypeOrder)

	__gong__resetStageType(&stage.AngleTypes, &stage.AngleTypes_mapString, &stage.AngleType_stagedOrder, &stage.AngleTypeOrder)

	__gong__resetStageType(&stage.ArcByCenterPointTypes, &stage.ArcByCenterPointTypes_mapString, &stage.ArcByCenterPointType_stagedOrder, &stage.ArcByCenterPointTypeOrder)

	__gong__resetStageType(&stage.ArcStringTypes, &stage.ArcStringTypes_mapString, &stage.ArcStringType_stagedOrder, &stage.ArcStringTypeOrder)

	__gong__resetStageType(&stage.ArcTypes, &stage.ArcTypes_mapString, &stage.ArcType_stagedOrder, &stage.ArcTypeOrder)

	__gong__resetStageType(&stage.BoundingShapeTypes, &stage.BoundingShapeTypes_mapString, &stage.BoundingShapeType_stagedOrder, &stage.BoundingShapeTypeOrder)

	__gong__resetStageType(&stage.CircleByCenterPointTypes, &stage.CircleByCenterPointTypes_mapString, &stage.CircleByCenterPointType_stagedOrder, &stage.CircleByCenterPointTypeOrder)

	__gong__resetStageType(&stage.CircleTypes, &stage.CircleTypes_mapString, &stage.CircleType_stagedOrder, &stage.CircleTypeOrder)

	__gong__resetStageType(&stage.CodeTypes, &stage.CodeTypes_mapString, &stage.CodeType_stagedOrder, &stage.CodeTypeOrder)

	__gong__resetStageType(&stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString, &stage.CodeWithAuthorityType_stagedOrder, &stage.CodeWithAuthorityTypeOrder)

	__gong__resetStageType(&stage.CompositeCurveTypes, &stage.CompositeCurveTypes_mapString, &stage.CompositeCurveType_stagedOrder, &stage.CompositeCurveTypeOrder)

	__gong__resetStageType(&stage.CurvePropertyTypes, &stage.CurvePropertyTypes_mapString, &stage.CurvePropertyType_stagedOrder, &stage.CurvePropertyTypeOrder)

	__gong__resetStageType(&stage.CurveSegmentArrayPropertyTypes, &stage.CurveSegmentArrayPropertyTypes_mapString, &stage.CurveSegmentArrayPropertyType_stagedOrder, &stage.CurveSegmentArrayPropertyTypeOrder)

	__gong__resetStageType(&stage.CurveTypes, &stage.CurveTypes_mapString, &stage.CurveType_stagedOrder, &stage.CurveTypeOrder)

	__gong__resetStageType(&stage.DirectPositionListTypes, &stage.DirectPositionListTypes_mapString, &stage.DirectPositionListType_stagedOrder, &stage.DirectPositionListTypeOrder)

	__gong__resetStageType(&stage.DirectPositionTypes, &stage.DirectPositionTypes_mapString, &stage.DirectPositionType_stagedOrder, &stage.DirectPositionTypeOrder)

	__gong__resetStageType(&stage.DynamicFeatureTypes, &stage.DynamicFeatureTypes_mapString, &stage.DynamicFeatureType_stagedOrder, &stage.DynamicFeatureTypeOrder)

	__gong__resetStageType(&stage.EnvelopeTypes, &stage.EnvelopeTypes_mapString, &stage.EnvelopeType_stagedOrder, &stage.EnvelopeTypeOrder)

	__gong__resetStageType(&stage.GeodesicStringTypes, &stage.GeodesicStringTypes_mapString, &stage.GeodesicStringType_stagedOrder, &stage.GeodesicStringTypeOrder)

	__gong__resetStageType(&stage.GeodesicTypes, &stage.GeodesicTypes_mapString, &stage.GeodesicType_stagedOrder, &stage.GeodesicTypeOrder)

	__gong__resetStageType(&stage.LengthTypes, &stage.LengthTypes_mapString, &stage.LengthType_stagedOrder, &stage.LengthTypeOrder)

	__gong__resetStageType(&stage.LineStringSegmentTypes, &stage.LineStringSegmentTypes_mapString, &stage.LineStringSegmentType_stagedOrder, &stage.LineStringSegmentTypeOrder)

	__gong__resetStageType(&stage.MeasureTypes, &stage.MeasureTypes_mapString, &stage.MeasureType_stagedOrder, &stage.MeasureTypeOrder)

	__gong__resetStageType(&stage.OrientableCurveTypes, &stage.OrientableCurveTypes_mapString, &stage.OrientableCurveType_stagedOrder, &stage.OrientableCurveTypeOrder)

	__gong__resetStageType(&stage.PointPropertyTypes, &stage.PointPropertyTypes_mapString, &stage.PointPropertyType_stagedOrder, &stage.PointPropertyTypeOrder)

	__gong__resetStageType(&stage.PointTypes, &stage.PointTypes_mapString, &stage.PointType_stagedOrder, &stage.PointTypeOrder)

	__gong__resetStageType(&stage.PolygonPatchTypes, &stage.PolygonPatchTypes_mapString, &stage.PolygonPatchType_stagedOrder, &stage.PolygonPatchTypeOrder)

	__gong__resetStageType(&stage.RingTypes, &stage.RingTypes_mapString, &stage.RingType_stagedOrder, &stage.RingTypeOrder)

	__gong__resetStageType(&stage.SurfacePatchArrayPropertyTypes, &stage.SurfacePatchArrayPropertyTypes_mapString, &stage.SurfacePatchArrayPropertyType_stagedOrder, &stage.SurfacePatchArrayPropertyTypeOrder)

	__gong__resetStageType(&stage.SurfaceTypes, &stage.SurfaceTypes_mapString, &stage.SurfaceType_stagedOrder, &stage.SurfaceTypeOrder)

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
func (AbstractCurveSegmentType) GongGetAssociationName() any {
	return &AbstractCurveSegmentType{
	}
}

func (AbstractCurveType) GongGetAssociationName() any {
	return &AbstractCurveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
	}
}

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

func (AbstractRingPropertyType) GongGetAssociationName() any {
	return &AbstractRingPropertyType{
			Ring: &RingType{Name: "Ring"},
	}
}

func (AbstractRingType) GongGetAssociationName() any {
	return &AbstractRingType{
	}
}

func (AbstractSurfacePatchType) GongGetAssociationName() any {
	return &AbstractSurfacePatchType{
	}
}

func (AbstractSurfaceType) GongGetAssociationName() any {
	return &AbstractSurfaceType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
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

func (AngleType) GongGetAssociationName() any {
	return &AngleType{
	}
}

func (ArcByCenterPointType) GongGetAssociationName() any {
	return &ArcByCenterPointType{
			Pos: &DirectPositionType{Name: "Pos"},
			PointProperty: &PointPropertyType{Name: "PointProperty"},
			PosList: &DirectPositionListType{Name: "PosList"},
			Radius: &LengthType{Name: "Radius"},
			StartAngle: &AngleType{Name: "StartAngle"},
			EndAngle: &AngleType{Name: "EndAngle"},
	}
}

func (ArcStringType) GongGetAssociationName() any {
	return &ArcStringType{
			Pos: []*DirectPositionType{{Name: "Pos"}},
			PointProperty: []*PointPropertyType{{Name: "PointProperty"}},
			PosList: &DirectPositionListType{Name: "PosList"},
	}
}

func (ArcType) GongGetAssociationName() any {
	return &ArcType{
			Pos: &DirectPositionType{Name: "Pos"},
			PointProperty: &PointPropertyType{Name: "PointProperty"},
			PosList: &DirectPositionListType{Name: "PosList"},
	}
}

func (BoundingShapeType) GongGetAssociationName() any {
	return &BoundingShapeType{
			Envelope: &EnvelopeType{Name: "Envelope"},
	}
}

func (CircleByCenterPointType) GongGetAssociationName() any {
	return &CircleByCenterPointType{
			Pos: &DirectPositionType{Name: "Pos"},
			PointProperty: &PointPropertyType{Name: "PointProperty"},
			PosList: &DirectPositionListType{Name: "PosList"},
			Radius: &LengthType{Name: "Radius"},
	}
}

func (CircleType) GongGetAssociationName() any {
	return &CircleType{
			Pos: []*DirectPositionType{{Name: "Pos"}},
			PointProperty: []*PointPropertyType{{Name: "PointProperty"}},
			PosList: &DirectPositionListType{Name: "PosList"},
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

func (CompositeCurveType) GongGetAssociationName() any {
	return &CompositeCurveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			CurveMember: []*CurvePropertyType{{Name: "CurveMember"}},
	}
}

func (CurvePropertyType) GongGetAssociationName() any {
	return &CurvePropertyType{
			Curve: &CurveType{Name: "Curve"},
			OrientableCurve: &OrientableCurveType{Name: "OrientableCurve"},
			CompositeCurve: &CompositeCurveType{Name: "CompositeCurve"},
	}
}

func (CurveSegmentArrayPropertyType) GongGetAssociationName() any {
	return &CurveSegmentArrayPropertyType{
			LineStringSegment: []*LineStringSegmentType{{Name: "LineStringSegment"}},
			ArcString: []*ArcStringType{{Name: "ArcString"}},
			Arc: []*ArcType{{Name: "Arc"}},
			Circle: []*CircleType{{Name: "Circle"}},
			ArcByCenterPoint: []*ArcByCenterPointType{{Name: "ArcByCenterPoint"}},
			CircleByCenterPoint: []*CircleByCenterPointType{{Name: "CircleByCenterPoint"}},
			GeodesicString: []*GeodesicStringType{{Name: "GeodesicString"}},
			Geodesic: []*GeodesicType{{Name: "Geodesic"}},
	}
}

func (CurveType) GongGetAssociationName() any {
	return &CurveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			Segments: &CurveSegmentArrayPropertyType{Name: "Segments"},
	}
}

func (DirectPositionListType) GongGetAssociationName() any {
	return &DirectPositionListType{
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

func (GeodesicStringType) GongGetAssociationName() any {
	return &GeodesicStringType{
			PosList: &DirectPositionListType{Name: "PosList"},
			Pos: &DirectPositionType{Name: "Pos"},
			PointProperty: &PointPropertyType{Name: "PointProperty"},
	}
}

func (GeodesicType) GongGetAssociationName() any {
	return &GeodesicType{
			PosList: &DirectPositionListType{Name: "PosList"},
			Pos: &DirectPositionType{Name: "Pos"},
			PointProperty: &PointPropertyType{Name: "PointProperty"},
	}
}

func (LengthType) GongGetAssociationName() any {
	return &LengthType{
	}
}

func (LineStringSegmentType) GongGetAssociationName() any {
	return &LineStringSegmentType{
			Pos: []*DirectPositionType{{Name: "Pos"}},
			PointProperty: []*PointPropertyType{{Name: "PointProperty"}},
			PosList: &DirectPositionListType{Name: "PosList"},
	}
}

func (MeasureType) GongGetAssociationName() any {
	return &MeasureType{
	}
}

func (OrientableCurveType) GongGetAssociationName() any {
	return &OrientableCurveType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			BaseCurve: &CurvePropertyType{Name: "BaseCurve"},
	}
}

func (PointPropertyType) GongGetAssociationName() any {
	return &PointPropertyType{
			Point: &PointType{Name: "Point"},
	}
}

func (PointType) GongGetAssociationName() any {
	return &PointType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			Pos: &DirectPositionType{Name: "Pos"},
	}
}

func (PolygonPatchType) GongGetAssociationName() any {
	return &PolygonPatchType{
			Exterior: &AbstractRingPropertyType{Name: "Exterior"},
	}
}

func (RingType) GongGetAssociationName() any {
	return &RingType{
			CurveMember: []*CurvePropertyType{{Name: "CurveMember"}},
	}
}

func (SurfacePatchArrayPropertyType) GongGetAssociationName() any {
	return &SurfacePatchArrayPropertyType{
			PolygonPatch: []*PolygonPatchType{{Name: "PolygonPatch"}},
	}
}

func (SurfaceType) GongGetAssociationName() any {
	return &SurfaceType{
			Identifier: &CodeWithAuthorityType{Name: "Identifier"},
			NameXSD: []*CodeType{{Name: "NameXSD"}},
			Patches: &SurfacePatchArrayPropertyType{Name: "Patches"},
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
	// reverse maps of direct associations of AbstractCurveSegmentType
	case AbstractCurveSegmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractCurveType
	case AbstractCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractCurveType)
			for abstractcurvetype := range stage.AbstractCurveTypes {
				if abstractcurvetype.Identifier != nil {
					codewithauthoritytype_ := abstractcurvetype.Identifier
					var abstractcurvetypes []*AbstractCurveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractcurvetypes = res[codewithauthoritytype_]
					} else {
						abstractcurvetypes = make([]*AbstractCurveType, 0)
					}
					abstractcurvetypes = append(abstractcurvetypes, abstractcurvetype)
					res[codewithauthoritytype_] = abstractcurvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
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
	// reverse maps of direct associations of AbstractRingPropertyType
	case AbstractRingPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Ring":
			res := make(map[*RingType][]*AbstractRingPropertyType)
			for abstractringpropertytype := range stage.AbstractRingPropertyTypes {
				if abstractringpropertytype.Ring != nil {
					ringtype_ := abstractringpropertytype.Ring
					var abstractringpropertytypes []*AbstractRingPropertyType
					_, ok := res[ringtype_]
					if ok {
						abstractringpropertytypes = res[ringtype_]
					} else {
						abstractringpropertytypes = make([]*AbstractRingPropertyType, 0)
					}
					abstractringpropertytypes = append(abstractringpropertytypes, abstractringpropertytype)
					res[ringtype_] = abstractringpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AbstractRingType
	case AbstractRingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractSurfacePatchType
	case AbstractSurfacePatchType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractSurfaceType
	case AbstractSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*AbstractSurfaceType)
			for abstractsurfacetype := range stage.AbstractSurfaceTypes {
				if abstractsurfacetype.Identifier != nil {
					codewithauthoritytype_ := abstractsurfacetype.Identifier
					var abstractsurfacetypes []*AbstractSurfaceType
					_, ok := res[codewithauthoritytype_]
					if ok {
						abstractsurfacetypes = res[codewithauthoritytype_]
					} else {
						abstractsurfacetypes = make([]*AbstractSurfaceType, 0)
					}
					abstractsurfacetypes = append(abstractsurfacetypes, abstractsurfacetype)
					res[codewithauthoritytype_] = abstractsurfacetypes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of AngleType
	case AngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArcByCenterPointType
	case ArcByCenterPointType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.Pos != nil {
					directpositiontype_ := arcbycenterpointtype.Pos
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[directpositiontype_]
					if ok {
						arcbycenterpointtypes = res[directpositiontype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[directpositiontype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.PointProperty != nil {
					pointpropertytype_ := arcbycenterpointtype.PointProperty
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[pointpropertytype_]
					if ok {
						arcbycenterpointtypes = res[pointpropertytype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[pointpropertytype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PosList":
			res := make(map[*DirectPositionListType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.PosList != nil {
					directpositionlisttype_ := arcbycenterpointtype.PosList
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[directpositionlisttype_]
					if ok {
						arcbycenterpointtypes = res[directpositionlisttype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[directpositionlisttype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Radius":
			res := make(map[*LengthType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.Radius != nil {
					lengthtype_ := arcbycenterpointtype.Radius
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[lengthtype_]
					if ok {
						arcbycenterpointtypes = res[lengthtype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[lengthtype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "StartAngle":
			res := make(map[*AngleType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.StartAngle != nil {
					angletype_ := arcbycenterpointtype.StartAngle
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[angletype_]
					if ok {
						arcbycenterpointtypes = res[angletype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[angletype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndAngle":
			res := make(map[*AngleType][]*ArcByCenterPointType)
			for arcbycenterpointtype := range stage.ArcByCenterPointTypes {
				if arcbycenterpointtype.EndAngle != nil {
					angletype_ := arcbycenterpointtype.EndAngle
					var arcbycenterpointtypes []*ArcByCenterPointType
					_, ok := res[angletype_]
					if ok {
						arcbycenterpointtypes = res[angletype_]
					} else {
						arcbycenterpointtypes = make([]*ArcByCenterPointType, 0)
					}
					arcbycenterpointtypes = append(arcbycenterpointtypes, arcbycenterpointtype)
					res[angletype_] = arcbycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ArcStringType
	case ArcStringType:
		switch fieldname {
		// insertion point for per direct association field
		case "PosList":
			res := make(map[*DirectPositionListType][]*ArcStringType)
			for arcstringtype := range stage.ArcStringTypes {
				if arcstringtype.PosList != nil {
					directpositionlisttype_ := arcstringtype.PosList
					var arcstringtypes []*ArcStringType
					_, ok := res[directpositionlisttype_]
					if ok {
						arcstringtypes = res[directpositionlisttype_]
					} else {
						arcstringtypes = make([]*ArcStringType, 0)
					}
					arcstringtypes = append(arcstringtypes, arcstringtype)
					res[directpositionlisttype_] = arcstringtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ArcType
	case ArcType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*ArcType)
			for arctype := range stage.ArcTypes {
				if arctype.Pos != nil {
					directpositiontype_ := arctype.Pos
					var arctypes []*ArcType
					_, ok := res[directpositiontype_]
					if ok {
						arctypes = res[directpositiontype_]
					} else {
						arctypes = make([]*ArcType, 0)
					}
					arctypes = append(arctypes, arctype)
					res[directpositiontype_] = arctypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*ArcType)
			for arctype := range stage.ArcTypes {
				if arctype.PointProperty != nil {
					pointpropertytype_ := arctype.PointProperty
					var arctypes []*ArcType
					_, ok := res[pointpropertytype_]
					if ok {
						arctypes = res[pointpropertytype_]
					} else {
						arctypes = make([]*ArcType, 0)
					}
					arctypes = append(arctypes, arctype)
					res[pointpropertytype_] = arctypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PosList":
			res := make(map[*DirectPositionListType][]*ArcType)
			for arctype := range stage.ArcTypes {
				if arctype.PosList != nil {
					directpositionlisttype_ := arctype.PosList
					var arctypes []*ArcType
					_, ok := res[directpositionlisttype_]
					if ok {
						arctypes = res[directpositionlisttype_]
					} else {
						arctypes = make([]*ArcType, 0)
					}
					arctypes = append(arctypes, arctype)
					res[directpositionlisttype_] = arctypes
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
	// reverse maps of direct associations of CircleByCenterPointType
	case CircleByCenterPointType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*CircleByCenterPointType)
			for circlebycenterpointtype := range stage.CircleByCenterPointTypes {
				if circlebycenterpointtype.Pos != nil {
					directpositiontype_ := circlebycenterpointtype.Pos
					var circlebycenterpointtypes []*CircleByCenterPointType
					_, ok := res[directpositiontype_]
					if ok {
						circlebycenterpointtypes = res[directpositiontype_]
					} else {
						circlebycenterpointtypes = make([]*CircleByCenterPointType, 0)
					}
					circlebycenterpointtypes = append(circlebycenterpointtypes, circlebycenterpointtype)
					res[directpositiontype_] = circlebycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*CircleByCenterPointType)
			for circlebycenterpointtype := range stage.CircleByCenterPointTypes {
				if circlebycenterpointtype.PointProperty != nil {
					pointpropertytype_ := circlebycenterpointtype.PointProperty
					var circlebycenterpointtypes []*CircleByCenterPointType
					_, ok := res[pointpropertytype_]
					if ok {
						circlebycenterpointtypes = res[pointpropertytype_]
					} else {
						circlebycenterpointtypes = make([]*CircleByCenterPointType, 0)
					}
					circlebycenterpointtypes = append(circlebycenterpointtypes, circlebycenterpointtype)
					res[pointpropertytype_] = circlebycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PosList":
			res := make(map[*DirectPositionListType][]*CircleByCenterPointType)
			for circlebycenterpointtype := range stage.CircleByCenterPointTypes {
				if circlebycenterpointtype.PosList != nil {
					directpositionlisttype_ := circlebycenterpointtype.PosList
					var circlebycenterpointtypes []*CircleByCenterPointType
					_, ok := res[directpositionlisttype_]
					if ok {
						circlebycenterpointtypes = res[directpositionlisttype_]
					} else {
						circlebycenterpointtypes = make([]*CircleByCenterPointType, 0)
					}
					circlebycenterpointtypes = append(circlebycenterpointtypes, circlebycenterpointtype)
					res[directpositionlisttype_] = circlebycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Radius":
			res := make(map[*LengthType][]*CircleByCenterPointType)
			for circlebycenterpointtype := range stage.CircleByCenterPointTypes {
				if circlebycenterpointtype.Radius != nil {
					lengthtype_ := circlebycenterpointtype.Radius
					var circlebycenterpointtypes []*CircleByCenterPointType
					_, ok := res[lengthtype_]
					if ok {
						circlebycenterpointtypes = res[lengthtype_]
					} else {
						circlebycenterpointtypes = make([]*CircleByCenterPointType, 0)
					}
					circlebycenterpointtypes = append(circlebycenterpointtypes, circlebycenterpointtype)
					res[lengthtype_] = circlebycenterpointtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CircleType
	case CircleType:
		switch fieldname {
		// insertion point for per direct association field
		case "PosList":
			res := make(map[*DirectPositionListType][]*CircleType)
			for circletype := range stage.CircleTypes {
				if circletype.PosList != nil {
					directpositionlisttype_ := circletype.PosList
					var circletypes []*CircleType
					_, ok := res[directpositionlisttype_]
					if ok {
						circletypes = res[directpositionlisttype_]
					} else {
						circletypes = make([]*CircleType, 0)
					}
					circletypes = append(circletypes, circletype)
					res[directpositionlisttype_] = circletypes
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
	// reverse maps of direct associations of CompositeCurveType
	case CompositeCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*CompositeCurveType)
			for compositecurvetype := range stage.CompositeCurveTypes {
				if compositecurvetype.Identifier != nil {
					codewithauthoritytype_ := compositecurvetype.Identifier
					var compositecurvetypes []*CompositeCurveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						compositecurvetypes = res[codewithauthoritytype_]
					} else {
						compositecurvetypes = make([]*CompositeCurveType, 0)
					}
					compositecurvetypes = append(compositecurvetypes, compositecurvetype)
					res[codewithauthoritytype_] = compositecurvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurvePropertyType
	case CurvePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Curve":
			res := make(map[*CurveType][]*CurvePropertyType)
			for curvepropertytype := range stage.CurvePropertyTypes {
				if curvepropertytype.Curve != nil {
					curvetype_ := curvepropertytype.Curve
					var curvepropertytypes []*CurvePropertyType
					_, ok := res[curvetype_]
					if ok {
						curvepropertytypes = res[curvetype_]
					} else {
						curvepropertytypes = make([]*CurvePropertyType, 0)
					}
					curvepropertytypes = append(curvepropertytypes, curvepropertytype)
					res[curvetype_] = curvepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "OrientableCurve":
			res := make(map[*OrientableCurveType][]*CurvePropertyType)
			for curvepropertytype := range stage.CurvePropertyTypes {
				if curvepropertytype.OrientableCurve != nil {
					orientablecurvetype_ := curvepropertytype.OrientableCurve
					var curvepropertytypes []*CurvePropertyType
					_, ok := res[orientablecurvetype_]
					if ok {
						curvepropertytypes = res[orientablecurvetype_]
					} else {
						curvepropertytypes = make([]*CurvePropertyType, 0)
					}
					curvepropertytypes = append(curvepropertytypes, curvepropertytype)
					res[orientablecurvetype_] = curvepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "CompositeCurve":
			res := make(map[*CompositeCurveType][]*CurvePropertyType)
			for curvepropertytype := range stage.CurvePropertyTypes {
				if curvepropertytype.CompositeCurve != nil {
					compositecurvetype_ := curvepropertytype.CompositeCurve
					var curvepropertytypes []*CurvePropertyType
					_, ok := res[compositecurvetype_]
					if ok {
						curvepropertytypes = res[compositecurvetype_]
					} else {
						curvepropertytypes = make([]*CurvePropertyType, 0)
					}
					curvepropertytypes = append(curvepropertytypes, curvepropertytype)
					res[compositecurvetype_] = curvepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurveSegmentArrayPropertyType
	case CurveSegmentArrayPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CurveType
	case CurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*CurveType)
			for curvetype := range stage.CurveTypes {
				if curvetype.Identifier != nil {
					codewithauthoritytype_ := curvetype.Identifier
					var curvetypes []*CurveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						curvetypes = res[codewithauthoritytype_]
					} else {
						curvetypes = make([]*CurveType, 0)
					}
					curvetypes = append(curvetypes, curvetype)
					res[codewithauthoritytype_] = curvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Segments":
			res := make(map[*CurveSegmentArrayPropertyType][]*CurveType)
			for curvetype := range stage.CurveTypes {
				if curvetype.Segments != nil {
					curvesegmentarraypropertytype_ := curvetype.Segments
					var curvetypes []*CurveType
					_, ok := res[curvesegmentarraypropertytype_]
					if ok {
						curvetypes = res[curvesegmentarraypropertytype_]
					} else {
						curvetypes = make([]*CurveType, 0)
					}
					curvetypes = append(curvetypes, curvetype)
					res[curvesegmentarraypropertytype_] = curvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DirectPositionListType
	case DirectPositionListType:
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
	// reverse maps of direct associations of GeodesicStringType
	case GeodesicStringType:
		switch fieldname {
		// insertion point for per direct association field
		case "PosList":
			res := make(map[*DirectPositionListType][]*GeodesicStringType)
			for geodesicstringtype := range stage.GeodesicStringTypes {
				if geodesicstringtype.PosList != nil {
					directpositionlisttype_ := geodesicstringtype.PosList
					var geodesicstringtypes []*GeodesicStringType
					_, ok := res[directpositionlisttype_]
					if ok {
						geodesicstringtypes = res[directpositionlisttype_]
					} else {
						geodesicstringtypes = make([]*GeodesicStringType, 0)
					}
					geodesicstringtypes = append(geodesicstringtypes, geodesicstringtype)
					res[directpositionlisttype_] = geodesicstringtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pos":
			res := make(map[*DirectPositionType][]*GeodesicStringType)
			for geodesicstringtype := range stage.GeodesicStringTypes {
				if geodesicstringtype.Pos != nil {
					directpositiontype_ := geodesicstringtype.Pos
					var geodesicstringtypes []*GeodesicStringType
					_, ok := res[directpositiontype_]
					if ok {
						geodesicstringtypes = res[directpositiontype_]
					} else {
						geodesicstringtypes = make([]*GeodesicStringType, 0)
					}
					geodesicstringtypes = append(geodesicstringtypes, geodesicstringtype)
					res[directpositiontype_] = geodesicstringtypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*GeodesicStringType)
			for geodesicstringtype := range stage.GeodesicStringTypes {
				if geodesicstringtype.PointProperty != nil {
					pointpropertytype_ := geodesicstringtype.PointProperty
					var geodesicstringtypes []*GeodesicStringType
					_, ok := res[pointpropertytype_]
					if ok {
						geodesicstringtypes = res[pointpropertytype_]
					} else {
						geodesicstringtypes = make([]*GeodesicStringType, 0)
					}
					geodesicstringtypes = append(geodesicstringtypes, geodesicstringtype)
					res[pointpropertytype_] = geodesicstringtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GeodesicType
	case GeodesicType:
		switch fieldname {
		// insertion point for per direct association field
		case "PosList":
			res := make(map[*DirectPositionListType][]*GeodesicType)
			for geodesictype := range stage.GeodesicTypes {
				if geodesictype.PosList != nil {
					directpositionlisttype_ := geodesictype.PosList
					var geodesictypes []*GeodesicType
					_, ok := res[directpositionlisttype_]
					if ok {
						geodesictypes = res[directpositionlisttype_]
					} else {
						geodesictypes = make([]*GeodesicType, 0)
					}
					geodesictypes = append(geodesictypes, geodesictype)
					res[directpositionlisttype_] = geodesictypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pos":
			res := make(map[*DirectPositionType][]*GeodesicType)
			for geodesictype := range stage.GeodesicTypes {
				if geodesictype.Pos != nil {
					directpositiontype_ := geodesictype.Pos
					var geodesictypes []*GeodesicType
					_, ok := res[directpositiontype_]
					if ok {
						geodesictypes = res[directpositiontype_]
					} else {
						geodesictypes = make([]*GeodesicType, 0)
					}
					geodesictypes = append(geodesictypes, geodesictype)
					res[directpositiontype_] = geodesictypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*GeodesicType)
			for geodesictype := range stage.GeodesicTypes {
				if geodesictype.PointProperty != nil {
					pointpropertytype_ := geodesictype.PointProperty
					var geodesictypes []*GeodesicType
					_, ok := res[pointpropertytype_]
					if ok {
						geodesictypes = res[pointpropertytype_]
					} else {
						geodesictypes = make([]*GeodesicType, 0)
					}
					geodesictypes = append(geodesictypes, geodesictype)
					res[pointpropertytype_] = geodesictypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LengthType
	case LengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineStringSegmentType
	case LineStringSegmentType:
		switch fieldname {
		// insertion point for per direct association field
		case "PosList":
			res := make(map[*DirectPositionListType][]*LineStringSegmentType)
			for linestringsegmenttype := range stage.LineStringSegmentTypes {
				if linestringsegmenttype.PosList != nil {
					directpositionlisttype_ := linestringsegmenttype.PosList
					var linestringsegmenttypes []*LineStringSegmentType
					_, ok := res[directpositionlisttype_]
					if ok {
						linestringsegmenttypes = res[directpositionlisttype_]
					} else {
						linestringsegmenttypes = make([]*LineStringSegmentType, 0)
					}
					linestringsegmenttypes = append(linestringsegmenttypes, linestringsegmenttype)
					res[directpositionlisttype_] = linestringsegmenttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeasureType
	case MeasureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OrientableCurveType
	case OrientableCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*OrientableCurveType)
			for orientablecurvetype := range stage.OrientableCurveTypes {
				if orientablecurvetype.Identifier != nil {
					codewithauthoritytype_ := orientablecurvetype.Identifier
					var orientablecurvetypes []*OrientableCurveType
					_, ok := res[codewithauthoritytype_]
					if ok {
						orientablecurvetypes = res[codewithauthoritytype_]
					} else {
						orientablecurvetypes = make([]*OrientableCurveType, 0)
					}
					orientablecurvetypes = append(orientablecurvetypes, orientablecurvetype)
					res[codewithauthoritytype_] = orientablecurvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "BaseCurve":
			res := make(map[*CurvePropertyType][]*OrientableCurveType)
			for orientablecurvetype := range stage.OrientableCurveTypes {
				if orientablecurvetype.BaseCurve != nil {
					curvepropertytype_ := orientablecurvetype.BaseCurve
					var orientablecurvetypes []*OrientableCurveType
					_, ok := res[curvepropertytype_]
					if ok {
						orientablecurvetypes = res[curvepropertytype_]
					} else {
						orientablecurvetypes = make([]*OrientableCurveType, 0)
					}
					orientablecurvetypes = append(orientablecurvetypes, orientablecurvetype)
					res[curvepropertytype_] = orientablecurvetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PointPropertyType
	case PointPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Point":
			res := make(map[*PointType][]*PointPropertyType)
			for pointpropertytype := range stage.PointPropertyTypes {
				if pointpropertytype.Point != nil {
					pointtype_ := pointpropertytype.Point
					var pointpropertytypes []*PointPropertyType
					_, ok := res[pointtype_]
					if ok {
						pointpropertytypes = res[pointtype_]
					} else {
						pointpropertytypes = make([]*PointPropertyType, 0)
					}
					pointpropertytypes = append(pointpropertytypes, pointpropertytype)
					res[pointtype_] = pointpropertytypes
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
	// reverse maps of direct associations of PolygonPatchType
	case PolygonPatchType:
		switch fieldname {
		// insertion point for per direct association field
		case "Exterior":
			res := make(map[*AbstractRingPropertyType][]*PolygonPatchType)
			for polygonpatchtype := range stage.PolygonPatchTypes {
				if polygonpatchtype.Exterior != nil {
					abstractringpropertytype_ := polygonpatchtype.Exterior
					var polygonpatchtypes []*PolygonPatchType
					_, ok := res[abstractringpropertytype_]
					if ok {
						polygonpatchtypes = res[abstractringpropertytype_]
					} else {
						polygonpatchtypes = make([]*PolygonPatchType, 0)
					}
					polygonpatchtypes = append(polygonpatchtypes, polygonpatchtype)
					res[abstractringpropertytype_] = polygonpatchtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RingType
	case RingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SurfacePatchArrayPropertyType
	case SurfacePatchArrayPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SurfaceType
	case SurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Identifier":
			res := make(map[*CodeWithAuthorityType][]*SurfaceType)
			for surfacetype := range stage.SurfaceTypes {
				if surfacetype.Identifier != nil {
					codewithauthoritytype_ := surfacetype.Identifier
					var surfacetypes []*SurfaceType
					_, ok := res[codewithauthoritytype_]
					if ok {
						surfacetypes = res[codewithauthoritytype_]
					} else {
						surfacetypes = make([]*SurfaceType, 0)
					}
					surfacetypes = append(surfacetypes, surfacetype)
					res[codewithauthoritytype_] = surfacetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Patches":
			res := make(map[*SurfacePatchArrayPropertyType][]*SurfaceType)
			for surfacetype := range stage.SurfaceTypes {
				if surfacetype.Patches != nil {
					surfacepatcharraypropertytype_ := surfacetype.Patches
					var surfacetypes []*SurfaceType
					_, ok := res[surfacepatcharraypropertytype_]
					if ok {
						surfacetypes = res[surfacepatcharraypropertytype_]
					} else {
						surfacetypes = make([]*SurfaceType, 0)
					}
					surfacetypes = append(surfacetypes, surfacetype)
					res[surfacepatcharraypropertytype_] = surfacetypes
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
	// reverse maps of direct associations of AbstractCurveSegmentType
	case AbstractCurveSegmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractCurveType
	case AbstractCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractCurveType)
			for abstractcurvetype := range stage.AbstractCurveTypes {
				for _, codetype_ := range abstractcurvetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractcurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
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
	// reverse maps of direct associations of AbstractRingPropertyType
	case AbstractRingPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractRingType
	case AbstractRingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractSurfacePatchType
	case AbstractSurfacePatchType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractSurfaceType
	case AbstractSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*AbstractSurfaceType)
			for abstractsurfacetype := range stage.AbstractSurfaceTypes {
				for _, codetype_ := range abstractsurfacetype.NameXSD {
					res[codetype_] = append(res[codetype_], abstractsurfacetype)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of AngleType
	case AngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArcByCenterPointType
	case ArcByCenterPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ArcStringType
	case ArcStringType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*ArcStringType)
			for arcstringtype := range stage.ArcStringTypes {
				for _, directpositiontype_ := range arcstringtype.Pos {
					res[directpositiontype_] = append(res[directpositiontype_], arcstringtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*ArcStringType)
			for arcstringtype := range stage.ArcStringTypes {
				for _, pointpropertytype_ := range arcstringtype.PointProperty {
					res[pointpropertytype_] = append(res[pointpropertytype_], arcstringtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ArcType
	case ArcType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BoundingShapeType
	case BoundingShapeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleByCenterPointType
	case CircleByCenterPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CircleType
	case CircleType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*CircleType)
			for circletype := range stage.CircleTypes {
				for _, directpositiontype_ := range circletype.Pos {
					res[directpositiontype_] = append(res[directpositiontype_], circletype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*CircleType)
			for circletype := range stage.CircleTypes {
				for _, pointpropertytype_ := range circletype.PointProperty {
					res[pointpropertytype_] = append(res[pointpropertytype_], circletype)
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
	// reverse maps of direct associations of CompositeCurveType
	case CompositeCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*CompositeCurveType)
			for compositecurvetype := range stage.CompositeCurveTypes {
				for _, codetype_ := range compositecurvetype.NameXSD {
					res[codetype_] = append(res[codetype_], compositecurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "CurveMember":
			res := make(map[*CurvePropertyType][]*CompositeCurveType)
			for compositecurvetype := range stage.CompositeCurveTypes {
				for _, curvepropertytype_ := range compositecurvetype.CurveMember {
					res[curvepropertytype_] = append(res[curvepropertytype_], compositecurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurvePropertyType
	case CurvePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CurveSegmentArrayPropertyType
	case CurveSegmentArrayPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "LineStringSegment":
			res := make(map[*LineStringSegmentType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, linestringsegmenttype_ := range curvesegmentarraypropertytype.LineStringSegment {
					res[linestringsegmenttype_] = append(res[linestringsegmenttype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ArcString":
			res := make(map[*ArcStringType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, arcstringtype_ := range curvesegmentarraypropertytype.ArcString {
					res[arcstringtype_] = append(res[arcstringtype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arc":
			res := make(map[*ArcType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, arctype_ := range curvesegmentarraypropertytype.Arc {
					res[arctype_] = append(res[arctype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Circle":
			res := make(map[*CircleType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, circletype_ := range curvesegmentarraypropertytype.Circle {
					res[circletype_] = append(res[circletype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ArcByCenterPoint":
			res := make(map[*ArcByCenterPointType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, arcbycenterpointtype_ := range curvesegmentarraypropertytype.ArcByCenterPoint {
					res[arcbycenterpointtype_] = append(res[arcbycenterpointtype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "CircleByCenterPoint":
			res := make(map[*CircleByCenterPointType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, circlebycenterpointtype_ := range curvesegmentarraypropertytype.CircleByCenterPoint {
					res[circlebycenterpointtype_] = append(res[circlebycenterpointtype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GeodesicString":
			res := make(map[*GeodesicStringType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, geodesicstringtype_ := range curvesegmentarraypropertytype.GeodesicString {
					res[geodesicstringtype_] = append(res[geodesicstringtype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Geodesic":
			res := make(map[*GeodesicType][]*CurveSegmentArrayPropertyType)
			for curvesegmentarraypropertytype := range stage.CurveSegmentArrayPropertyTypes {
				for _, geodesictype_ := range curvesegmentarraypropertytype.Geodesic {
					res[geodesictype_] = append(res[geodesictype_], curvesegmentarraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurveType
	case CurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*CurveType)
			for curvetype := range stage.CurveTypes {
				for _, codetype_ := range curvetype.NameXSD {
					res[codetype_] = append(res[codetype_], curvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DirectPositionListType
	case DirectPositionListType:
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
	// reverse maps of direct associations of GeodesicStringType
	case GeodesicStringType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeodesicType
	case GeodesicType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LengthType
	case LengthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineStringSegmentType
	case LineStringSegmentType:
		switch fieldname {
		// insertion point for per direct association field
		case "Pos":
			res := make(map[*DirectPositionType][]*LineStringSegmentType)
			for linestringsegmenttype := range stage.LineStringSegmentTypes {
				for _, directpositiontype_ := range linestringsegmenttype.Pos {
					res[directpositiontype_] = append(res[directpositiontype_], linestringsegmenttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointProperty":
			res := make(map[*PointPropertyType][]*LineStringSegmentType)
			for linestringsegmenttype := range stage.LineStringSegmentTypes {
				for _, pointpropertytype_ := range linestringsegmenttype.PointProperty {
					res[pointpropertytype_] = append(res[pointpropertytype_], linestringsegmenttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeasureType
	case MeasureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OrientableCurveType
	case OrientableCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*OrientableCurveType)
			for orientablecurvetype := range stage.OrientableCurveTypes {
				for _, codetype_ := range orientablecurvetype.NameXSD {
					res[codetype_] = append(res[codetype_], orientablecurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PointPropertyType
	case PointPropertyType:
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
	// reverse maps of direct associations of PolygonPatchType
	case PolygonPatchType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RingType
	case RingType:
		switch fieldname {
		// insertion point for per direct association field
		case "CurveMember":
			res := make(map[*CurvePropertyType][]*RingType)
			for ringtype := range stage.RingTypes {
				for _, curvepropertytype_ := range ringtype.CurveMember {
					res[curvepropertytype_] = append(res[curvepropertytype_], ringtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SurfacePatchArrayPropertyType
	case SurfacePatchArrayPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "PolygonPatch":
			res := make(map[*PolygonPatchType][]*SurfacePatchArrayPropertyType)
			for surfacepatcharraypropertytype := range stage.SurfacePatchArrayPropertyTypes {
				for _, polygonpatchtype_ := range surfacepatcharraypropertytype.PolygonPatch {
					res[polygonpatchtype_] = append(res[polygonpatchtype_], surfacepatcharraypropertytype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SurfaceType
	case SurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "NameXSD":
			res := make(map[*CodeType][]*SurfaceType)
			for surfacetype := range stage.SurfaceTypes {
				for _, codetype_ := range surfacetype.NameXSD {
					res[codetype_] = append(res[codetype_], surfacetype)
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
func (*AbstractCurveSegmentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractCurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

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

func (*AbstractRingPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractRingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractSurfacePatchType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AbstractSurfaceType) GongGetReverseFields() []GongReverseField {
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

func (*AngleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ArcByCenterPointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "ArcByCenterPoint",
		},
	}
}

func (*ArcStringType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "ArcString",
		},
	}
}

func (*ArcType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "Arc",
		},
	}
}

func (*BoundingShapeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CircleByCenterPointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "CircleByCenterPoint",
		},
	}
}

func (*CircleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "Circle",
		},
	}
}

func (*CodeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AbstractCurveType",
			Fieldname: "NameXSD",
		},
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
			GongstructName: "AbstractSurfaceType",
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
			GongstructName: "CompositeCurveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "CurveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "DynamicFeatureType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "OrientableCurveType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "PointType",
			Fieldname: "NameXSD",
		},
		{
			GongstructName: "SurfaceType",
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

func (*CompositeCurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CurvePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CompositeCurveType",
			Fieldname: "CurveMember",
		},
		{
			GongstructName: "RingType",
			Fieldname: "CurveMember",
		},
	}
}

func (*CurveSegmentArrayPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DirectPositionListType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DirectPositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ArcStringType",
			Fieldname: "Pos",
		},
		{
			GongstructName: "CircleType",
			Fieldname: "Pos",
		},
		{
			GongstructName: "LineStringSegmentType",
			Fieldname: "Pos",
		},
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

func (*GeodesicStringType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "GeodesicString",
		},
	}
}

func (*GeodesicType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "Geodesic",
		},
	}
}

func (*LengthType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LineStringSegmentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "CurveSegmentArrayPropertyType",
			Fieldname: "LineStringSegment",
		},
	}
}

func (*MeasureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OrientableCurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PointPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ArcStringType",
			Fieldname: "PointProperty",
		},
		{
			GongstructName: "CircleType",
			Fieldname: "PointProperty",
		},
		{
			GongstructName: "LineStringSegmentType",
			Fieldname: "PointProperty",
		},
	}
}

func (*PointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PolygonPatchType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "SurfacePatchArrayPropertyType",
			Fieldname: "PolygonPatch",
		},
	}
}

func (*RingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SurfacePatchArrayPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SurfaceType) GongGetReverseFields() []GongReverseField {
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
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (abstractcurvetype *AbstractCurveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (abstractringpropertytype *AbstractRingPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Ring",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RingType",
		},
	}
	return
}

func (abstractringtype *AbstractRingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (abstractsurfacetype *AbstractSurfaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (arcbycenterpointtype *ArcByCenterPointType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:               "NumArc",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
		},
		{
			Name:                 "Radius",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LengthType",
		},
		{
			Name:                 "StartAngle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AngleType",
		},
		{
			Name:                 "EndAngle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AngleType",
		},
	}
	return
}

func (arcstringtype *ArcStringType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:               "NumArc",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
		},
	}
	return
}

func (arctype *ArcType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:               "NumArc",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
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

func (circlebycenterpointtype *CircleByCenterPointType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:               "NumArc",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
		},
		{
			Name:                 "Radius",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LengthType",
		},
	}
	return
}

func (circletype *CircleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:               "NumArc",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
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

func (compositecurvetype *CompositeCurveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "AggregationType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CurveMember",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CurvePropertyType",
		},
	}
	return
}

func (curvepropertytype *CurvePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Curve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurveType",
		},
		{
			Name:                 "OrientableCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "OrientableCurveType",
		},
		{
			Name:                 "CompositeCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CompositeCurveType",
		},
	}
	return
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "LineStringSegment",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LineStringSegmentType",
		},
		{
			Name:                 "ArcString",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArcStringType",
		},
		{
			Name:                 "Arc",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArcType",
		},
		{
			Name:                 "Circle",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CircleType",
		},
		{
			Name:                 "ArcByCenterPoint",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ArcByCenterPointType",
		},
		{
			Name:                 "CircleByCenterPoint",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CircleByCenterPointType",
		},
		{
			Name:                 "GeodesicString",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GeodesicStringType",
		},
		{
			Name:                 "Geodesic",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GeodesicType",
		},
	}
	return
}

func (curvetype *CurveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Segments",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurveSegmentArrayPropertyType",
		},
	}
	return
}

func (directpositionlisttype *DirectPositionListType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (geodesicstringtype *GeodesicStringType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointPropertyType",
		},
	}
	return
}

func (geodesictype *GeodesicType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointPropertyType",
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

func (linestringsegmenttype *LineStringSegmentType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumDerivativesAtStart",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativesAtEnd",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NumDerivativeInterior",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_CurveInterpolationType",
		},
		{
			Name:                 "Pos",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DirectPositionType",
		},
		{
			Name:                 "PointProperty",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PointPropertyType",
		},
		{
			Name:                 "PosList",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DirectPositionListType",
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

func (orientablecurvetype *OrientableCurveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Orientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_SignType",
		},
		{
			Name:                 "BaseCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurvePropertyType",
		},
	}
	return
}

func (pointpropertytype *PointPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Point",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PointType",
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

func (polygonpatchtype *PolygonPatchType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Interpolation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_SurfaceInterpolationType",
		},
		{
			Name:                 "Exterior",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AbstractRingPropertyType",
		},
	}
	return
}

func (ringtype *RingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "AggregationType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CurveMember",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CurvePropertyType",
		},
	}
	return
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PolygonPatch",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PolygonPatchType",
		},
	}
	return
}

func (surfacetype *SurfaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Patches",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SurfacePatchArrayPropertyType",
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
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractcurvesegmenttype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", abstractcurvesegmenttype.NumDerivativesAtStart)
		res.valueInt = abstractcurvesegmenttype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", abstractcurvesegmenttype.NumDerivativesAtEnd)
		res.valueInt = abstractcurvesegmenttype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", abstractcurvesegmenttype.NumDerivativeInterior)
		res.valueInt = abstractcurvesegmenttype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (abstractcurvetype *AbstractCurveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractcurvetype.Name
	case "Id":
		res.valueString = abstractcurvetype.Id
	case "Description":
		res.valueString = abstractcurvetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractcurvetype.Identifier != nil {
			res.valueString = abstractcurvetype.Identifier.Name
			res.ids = abstractcurvetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractcurvetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = abstractcurvetype.SrsName
	}
	return
}

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

func (abstractringpropertytype *AbstractRingPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractringpropertytype.Name
	case "Ring":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractringpropertytype.Ring != nil {
			res.valueString = abstractringpropertytype.Ring.Name
			res.ids = abstractringpropertytype.Ring.GongGetUUID(stage)
		}
	}
	return
}

func (abstractringtype *AbstractRingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractringtype.Name
	}
	return
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractsurfacepatchtype.Name
	}
	return
}

func (abstractsurfacetype *AbstractSurfaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractsurfacetype.Name
	case "Id":
		res.valueString = abstractsurfacetype.Id
	case "Description":
		res.valueString = abstractsurfacetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if abstractsurfacetype.Identifier != nil {
			res.valueString = abstractsurfacetype.Identifier.Name
			res.ids = abstractsurfacetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractsurfacetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = abstractsurfacetype.SrsName
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

func (arcbycenterpointtype *ArcByCenterPointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arcbycenterpointtype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", arcbycenterpointtype.NumDerivativesAtStart)
		res.valueInt = arcbycenterpointtype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", arcbycenterpointtype.NumDerivativesAtEnd)
		res.valueInt = arcbycenterpointtype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", arcbycenterpointtype.NumDerivativeInterior)
		res.valueInt = arcbycenterpointtype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := arcbycenterpointtype.Interpolation
		res.valueString = enum.ToCodeString()
	case "NumArc":
		res.valueString = fmt.Sprintf("%d", arcbycenterpointtype.NumArc)
		res.valueInt = arcbycenterpointtype.NumArc
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.Pos != nil {
			res.valueString = arcbycenterpointtype.Pos.Name
			res.ids = arcbycenterpointtype.Pos.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.PointProperty != nil {
			res.valueString = arcbycenterpointtype.PointProperty.Name
			res.ids = arcbycenterpointtype.PointProperty.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.PosList != nil {
			res.valueString = arcbycenterpointtype.PosList.Name
			res.ids = arcbycenterpointtype.PosList.GongGetUUID(stage)
		}
	case "Radius":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.Radius != nil {
			res.valueString = arcbycenterpointtype.Radius.Name
			res.ids = arcbycenterpointtype.Radius.GongGetUUID(stage)
		}
	case "StartAngle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.StartAngle != nil {
			res.valueString = arcbycenterpointtype.StartAngle.Name
			res.ids = arcbycenterpointtype.StartAngle.GongGetUUID(stage)
		}
	case "EndAngle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcbycenterpointtype.EndAngle != nil {
			res.valueString = arcbycenterpointtype.EndAngle.Name
			res.ids = arcbycenterpointtype.EndAngle.GongGetUUID(stage)
		}
	}
	return
}

func (arcstringtype *ArcStringType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arcstringtype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", arcstringtype.NumDerivativesAtStart)
		res.valueInt = arcstringtype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", arcstringtype.NumDerivativesAtEnd)
		res.valueInt = arcstringtype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", arcstringtype.NumDerivativeInterior)
		res.valueInt = arcstringtype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := arcstringtype.Interpolation
		res.valueString = enum.ToCodeString()
	case "NumArc":
		res.valueString = fmt.Sprintf("%d", arcstringtype.NumArc)
		res.valueInt = arcstringtype.NumArc
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range arcstringtype.Pos {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range arcstringtype.PointProperty {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arcstringtype.PosList != nil {
			res.valueString = arcstringtype.PosList.Name
			res.ids = arcstringtype.PosList.GongGetUUID(stage)
		}
	}
	return
}

func (arctype *ArcType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arctype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", arctype.NumDerivativesAtStart)
		res.valueInt = arctype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", arctype.NumDerivativesAtEnd)
		res.valueInt = arctype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", arctype.NumDerivativeInterior)
		res.valueInt = arctype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := arctype.Interpolation
		res.valueString = enum.ToCodeString()
	case "NumArc":
		res.valueString = fmt.Sprintf("%d", arctype.NumArc)
		res.valueInt = arctype.NumArc
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arctype.Pos != nil {
			res.valueString = arctype.Pos.Name
			res.ids = arctype.Pos.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arctype.PointProperty != nil {
			res.valueString = arctype.PointProperty.Name
			res.ids = arctype.PointProperty.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arctype.PosList != nil {
			res.valueString = arctype.PosList.Name
			res.ids = arctype.PosList.GongGetUUID(stage)
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

func (circlebycenterpointtype *CircleByCenterPointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circlebycenterpointtype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", circlebycenterpointtype.NumDerivativesAtStart)
		res.valueInt = circlebycenterpointtype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", circlebycenterpointtype.NumDerivativesAtEnd)
		res.valueInt = circlebycenterpointtype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", circlebycenterpointtype.NumDerivativeInterior)
		res.valueInt = circlebycenterpointtype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := circlebycenterpointtype.Interpolation
		res.valueString = enum.ToCodeString()
	case "NumArc":
		res.valueString = fmt.Sprintf("%d", circlebycenterpointtype.NumArc)
		res.valueInt = circlebycenterpointtype.NumArc
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlebycenterpointtype.Pos != nil {
			res.valueString = circlebycenterpointtype.Pos.Name
			res.ids = circlebycenterpointtype.Pos.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlebycenterpointtype.PointProperty != nil {
			res.valueString = circlebycenterpointtype.PointProperty.Name
			res.ids = circlebycenterpointtype.PointProperty.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlebycenterpointtype.PosList != nil {
			res.valueString = circlebycenterpointtype.PosList.Name
			res.ids = circlebycenterpointtype.PosList.GongGetUUID(stage)
		}
	case "Radius":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circlebycenterpointtype.Radius != nil {
			res.valueString = circlebycenterpointtype.Radius.Name
			res.ids = circlebycenterpointtype.Radius.GongGetUUID(stage)
		}
	}
	return
}

func (circletype *CircleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = circletype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", circletype.NumDerivativesAtStart)
		res.valueInt = circletype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", circletype.NumDerivativesAtEnd)
		res.valueInt = circletype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", circletype.NumDerivativeInterior)
		res.valueInt = circletype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := circletype.Interpolation
		res.valueString = enum.ToCodeString()
	case "NumArc":
		res.valueString = fmt.Sprintf("%d", circletype.NumArc)
		res.valueInt = circletype.NumArc
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range circletype.Pos {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range circletype.PointProperty {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if circletype.PosList != nil {
			res.valueString = circletype.PosList.Name
			res.ids = circletype.PosList.GongGetUUID(stage)
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

func (compositecurvetype *CompositeCurveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = compositecurvetype.Name
	case "Id":
		res.valueString = compositecurvetype.Id
	case "Description":
		res.valueString = compositecurvetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if compositecurvetype.Identifier != nil {
			res.valueString = compositecurvetype.Identifier.Name
			res.ids = compositecurvetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range compositecurvetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = compositecurvetype.SrsName
	case "AggregationType":
		res.valueString = compositecurvetype.AggregationType
	case "CurveMember":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range compositecurvetype.CurveMember {
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

func (curvepropertytype *CurvePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = curvepropertytype.Name
	case "NilReason":
		res.valueString = curvepropertytype.NilReason
	case "Type_XSD":
		res.valueString = curvepropertytype.Type_XSD
	case "Href":
		res.valueString = curvepropertytype.Href
	case "Role":
		res.valueString = curvepropertytype.Role
	case "Arcrole":
		res.valueString = curvepropertytype.Arcrole
	case "Title":
		res.valueString = curvepropertytype.Title
	case "Show":
		res.valueString = curvepropertytype.Show
	case "Actuate":
		res.valueString = curvepropertytype.Actuate
	case "Curve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvepropertytype.Curve != nil {
			res.valueString = curvepropertytype.Curve.Name
			res.ids = curvepropertytype.Curve.GongGetUUID(stage)
		}
	case "OrientableCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvepropertytype.OrientableCurve != nil {
			res.valueString = curvepropertytype.OrientableCurve.Name
			res.ids = curvepropertytype.OrientableCurve.GongGetUUID(stage)
		}
	case "CompositeCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvepropertytype.CompositeCurve != nil {
			res.valueString = curvepropertytype.CompositeCurve.Name
			res.ids = curvepropertytype.CompositeCurve.GongGetUUID(stage)
		}
	}
	return
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = curvesegmentarraypropertytype.Name
	case "LineStringSegment":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.LineStringSegment {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ArcString":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.ArcString {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Arc":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.Arc {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Circle":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.Circle {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ArcByCenterPoint":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.ArcByCenterPoint {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CircleByCenterPoint":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.CircleByCenterPoint {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GeodesicString":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.GeodesicString {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Geodesic":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvesegmentarraypropertytype.Geodesic {
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

func (curvetype *CurveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = curvetype.Name
	case "Id":
		res.valueString = curvetype.Id
	case "Description":
		res.valueString = curvetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvetype.Identifier != nil {
			res.valueString = curvetype.Identifier.Name
			res.ids = curvetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = curvetype.SrsName
	case "Segments":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvetype.Segments != nil {
			res.valueString = curvetype.Segments.Name
			res.ids = curvetype.Segments.GongGetUUID(stage)
		}
	}
	return
}

func (directpositionlisttype *DirectPositionListType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = directpositionlisttype.Name
	case "SrsName":
		res.valueString = directpositionlisttype.SrsName
	case "EnclosedText":
		res.valueString = directpositionlisttype.EnclosedText
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

func (geodesicstringtype *GeodesicStringType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geodesicstringtype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", geodesicstringtype.NumDerivativesAtStart)
		res.valueInt = geodesicstringtype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", geodesicstringtype.NumDerivativesAtEnd)
		res.valueInt = geodesicstringtype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", geodesicstringtype.NumDerivativeInterior)
		res.valueInt = geodesicstringtype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := geodesicstringtype.Interpolation
		res.valueString = enum.ToCodeString()
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesicstringtype.PosList != nil {
			res.valueString = geodesicstringtype.PosList.Name
			res.ids = geodesicstringtype.PosList.GongGetUUID(stage)
		}
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesicstringtype.Pos != nil {
			res.valueString = geodesicstringtype.Pos.Name
			res.ids = geodesicstringtype.Pos.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesicstringtype.PointProperty != nil {
			res.valueString = geodesicstringtype.PointProperty.Name
			res.ids = geodesicstringtype.PointProperty.GongGetUUID(stage)
		}
	}
	return
}

func (geodesictype *GeodesicType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geodesictype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", geodesictype.NumDerivativesAtStart)
		res.valueInt = geodesictype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", geodesictype.NumDerivativesAtEnd)
		res.valueInt = geodesictype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", geodesictype.NumDerivativeInterior)
		res.valueInt = geodesictype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := geodesictype.Interpolation
		res.valueString = enum.ToCodeString()
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesictype.PosList != nil {
			res.valueString = geodesictype.PosList.Name
			res.ids = geodesictype.PosList.GongGetUUID(stage)
		}
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesictype.Pos != nil {
			res.valueString = geodesictype.Pos.Name
			res.ids = geodesictype.Pos.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geodesictype.PointProperty != nil {
			res.valueString = geodesictype.PointProperty.Name
			res.ids = geodesictype.PointProperty.GongGetUUID(stage)
		}
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

func (linestringsegmenttype *LineStringSegmentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linestringsegmenttype.Name
	case "NumDerivativesAtStart":
		res.valueString = fmt.Sprintf("%d", linestringsegmenttype.NumDerivativesAtStart)
		res.valueInt = linestringsegmenttype.NumDerivativesAtStart
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativesAtEnd":
		res.valueString = fmt.Sprintf("%d", linestringsegmenttype.NumDerivativesAtEnd)
		res.valueInt = linestringsegmenttype.NumDerivativesAtEnd
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumDerivativeInterior":
		res.valueString = fmt.Sprintf("%d", linestringsegmenttype.NumDerivativeInterior)
		res.valueInt = linestringsegmenttype.NumDerivativeInterior
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Interpolation":
		enum := linestringsegmenttype.Interpolation
		res.valueString = enum.ToCodeString()
	case "Pos":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range linestringsegmenttype.Pos {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PointProperty":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range linestringsegmenttype.PointProperty {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PosList":
		res.GongFieldValueType = GongFieldValueTypePointer
		if linestringsegmenttype.PosList != nil {
			res.valueString = linestringsegmenttype.PosList.Name
			res.ids = linestringsegmenttype.PosList.GongGetUUID(stage)
		}
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

func (orientablecurvetype *OrientableCurveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = orientablecurvetype.Name
	case "Id":
		res.valueString = orientablecurvetype.Id
	case "Description":
		res.valueString = orientablecurvetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if orientablecurvetype.Identifier != nil {
			res.valueString = orientablecurvetype.Identifier.Name
			res.ids = orientablecurvetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range orientablecurvetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = orientablecurvetype.SrsName
	case "Orientation":
		enum := orientablecurvetype.Orientation
		res.valueString = enum.ToCodeString()
	case "BaseCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if orientablecurvetype.BaseCurve != nil {
			res.valueString = orientablecurvetype.BaseCurve.Name
			res.ids = orientablecurvetype.BaseCurve.GongGetUUID(stage)
		}
	}
	return
}

func (pointpropertytype *PointPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pointpropertytype.Name
	case "NilReason":
		res.valueString = pointpropertytype.NilReason
	case "Type_XSD":
		res.valueString = pointpropertytype.Type_XSD
	case "Href":
		res.valueString = pointpropertytype.Href
	case "Role":
		res.valueString = pointpropertytype.Role
	case "Arcrole":
		res.valueString = pointpropertytype.Arcrole
	case "Title":
		res.valueString = pointpropertytype.Title
	case "Show":
		res.valueString = pointpropertytype.Show
	case "Actuate":
		res.valueString = pointpropertytype.Actuate
	case "Point":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pointpropertytype.Point != nil {
			res.valueString = pointpropertytype.Point.Name
			res.ids = pointpropertytype.Point.GongGetUUID(stage)
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

func (polygonpatchtype *PolygonPatchType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = polygonpatchtype.Name
	case "Interpolation":
		enum := polygonpatchtype.Interpolation
		res.valueString = enum.ToCodeString()
	case "Exterior":
		res.GongFieldValueType = GongFieldValueTypePointer
		if polygonpatchtype.Exterior != nil {
			res.valueString = polygonpatchtype.Exterior.Name
			res.ids = polygonpatchtype.Exterior.GongGetUUID(stage)
		}
	}
	return
}

func (ringtype *RingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ringtype.Name
	case "AggregationType":
		res.valueString = ringtype.AggregationType
	case "CurveMember":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ringtype.CurveMember {
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

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = surfacepatcharraypropertytype.Name
	case "PolygonPatch":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range surfacepatcharraypropertytype.PolygonPatch {
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

func (surfacetype *SurfaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = surfacetype.Name
	case "Id":
		res.valueString = surfacetype.Id
	case "Description":
		res.valueString = surfacetype.Description
	case "Identifier":
		res.GongFieldValueType = GongFieldValueTypePointer
		if surfacetype.Identifier != nil {
			res.valueString = surfacetype.Identifier.Name
			res.ids = surfacetype.Identifier.GongGetUUID(stage)
		}
	case "NameXSD":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range surfacetype.NameXSD {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SrsName":
		res.valueString = surfacetype.SrsName
	case "Patches":
		res.GongFieldValueType = GongFieldValueTypePointer
		if surfacetype.Patches != nil {
			res.valueString = surfacetype.Patches.Name
			res.ids = surfacetype.Patches.GongGetUUID(stage)
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
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongGetGongstructName() string {
	return "AbstractCurveSegmentType"
}

func (abstractcurvetype *AbstractCurveType) GongGetGongstructName() string {
	return "AbstractCurveType"
}

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

func (abstractringpropertytype *AbstractRingPropertyType) GongGetGongstructName() string {
	return "AbstractRingPropertyType"
}

func (abstractringtype *AbstractRingType) GongGetGongstructName() string {
	return "AbstractRingType"
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongGetGongstructName() string {
	return "AbstractSurfacePatchType"
}

func (abstractsurfacetype *AbstractSurfaceType) GongGetGongstructName() string {
	return "AbstractSurfaceType"
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

func (angletype *AngleType) GongGetGongstructName() string {
	return "AngleType"
}

func (arcbycenterpointtype *ArcByCenterPointType) GongGetGongstructName() string {
	return "ArcByCenterPointType"
}

func (arcstringtype *ArcStringType) GongGetGongstructName() string {
	return "ArcStringType"
}

func (arctype *ArcType) GongGetGongstructName() string {
	return "ArcType"
}

func (boundingshapetype *BoundingShapeType) GongGetGongstructName() string {
	return "BoundingShapeType"
}

func (circlebycenterpointtype *CircleByCenterPointType) GongGetGongstructName() string {
	return "CircleByCenterPointType"
}

func (circletype *CircleType) GongGetGongstructName() string {
	return "CircleType"
}

func (codetype *CodeType) GongGetGongstructName() string {
	return "CodeType"
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetGongstructName() string {
	return "CodeWithAuthorityType"
}

func (compositecurvetype *CompositeCurveType) GongGetGongstructName() string {
	return "CompositeCurveType"
}

func (curvepropertytype *CurvePropertyType) GongGetGongstructName() string {
	return "CurvePropertyType"
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongGetGongstructName() string {
	return "CurveSegmentArrayPropertyType"
}

func (curvetype *CurveType) GongGetGongstructName() string {
	return "CurveType"
}

func (directpositionlisttype *DirectPositionListType) GongGetGongstructName() string {
	return "DirectPositionListType"
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

func (geodesicstringtype *GeodesicStringType) GongGetGongstructName() string {
	return "GeodesicStringType"
}

func (geodesictype *GeodesicType) GongGetGongstructName() string {
	return "GeodesicType"
}

func (lengthtype *LengthType) GongGetGongstructName() string {
	return "LengthType"
}

func (linestringsegmenttype *LineStringSegmentType) GongGetGongstructName() string {
	return "LineStringSegmentType"
}

func (measuretype *MeasureType) GongGetGongstructName() string {
	return "MeasureType"
}

func (orientablecurvetype *OrientableCurveType) GongGetGongstructName() string {
	return "OrientableCurveType"
}

func (pointpropertytype *PointPropertyType) GongGetGongstructName() string {
	return "PointPropertyType"
}

func (pointtype *PointType) GongGetGongstructName() string {
	return "PointType"
}

func (polygonpatchtype *PolygonPatchType) GongGetGongstructName() string {
	return "PolygonPatchType"
}

func (ringtype *RingType) GongGetGongstructName() string {
	return "RingType"
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongGetGongstructName() string {
	return "SurfacePatchArrayPropertyType"
}

func (surfacetype *SurfaceType) GongGetGongstructName() string {
	return "SurfaceType"
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
	__gong__rebuildMapString(stage.AbstractCurveSegmentTypes, &stage.AbstractCurveSegmentTypes_mapString)

	__gong__rebuildMapString(stage.AbstractCurveTypes, &stage.AbstractCurveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractFeatureMemberTypes, &stage.AbstractFeatureMemberTypes_mapString)

	__gong__rebuildMapString(stage.AbstractFeatureTypes, &stage.AbstractFeatureTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGMLTypes, &stage.AbstractGMLTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGeometricPrimitiveTypes, &stage.AbstractGeometricPrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractGeometryTypes, &stage.AbstractGeometryTypes_mapString)

	__gong__rebuildMapString(stage.AbstractMetadataPropertyTypes, &stage.AbstractMetadataPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AbstractRingPropertyTypes, &stage.AbstractRingPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AbstractRingTypes, &stage.AbstractRingTypes_mapString)

	__gong__rebuildMapString(stage.AbstractSurfacePatchTypes, &stage.AbstractSurfacePatchTypes_mapString)

	__gong__rebuildMapString(stage.AbstractSurfaceTypes, &stage.AbstractSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeGeometricPrimitiveTypes, &stage.AbstractTimeGeometricPrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeObjectTypes, &stage.AbstractTimeObjectTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimePrimitiveTypes, &stage.AbstractTimePrimitiveTypes_mapString)

	__gong__rebuildMapString(stage.AbstractTimeSliceTypes, &stage.AbstractTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.AngleTypes, &stage.AngleTypes_mapString)

	__gong__rebuildMapString(stage.ArcByCenterPointTypes, &stage.ArcByCenterPointTypes_mapString)

	__gong__rebuildMapString(stage.ArcStringTypes, &stage.ArcStringTypes_mapString)

	__gong__rebuildMapString(stage.ArcTypes, &stage.ArcTypes_mapString)

	__gong__rebuildMapString(stage.BoundingShapeTypes, &stage.BoundingShapeTypes_mapString)

	__gong__rebuildMapString(stage.CircleByCenterPointTypes, &stage.CircleByCenterPointTypes_mapString)

	__gong__rebuildMapString(stage.CircleTypes, &stage.CircleTypes_mapString)

	__gong__rebuildMapString(stage.CodeTypes, &stage.CodeTypes_mapString)

	__gong__rebuildMapString(stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_mapString)

	__gong__rebuildMapString(stage.CompositeCurveTypes, &stage.CompositeCurveTypes_mapString)

	__gong__rebuildMapString(stage.CurvePropertyTypes, &stage.CurvePropertyTypes_mapString)

	__gong__rebuildMapString(stage.CurveSegmentArrayPropertyTypes, &stage.CurveSegmentArrayPropertyTypes_mapString)

	__gong__rebuildMapString(stage.CurveTypes, &stage.CurveTypes_mapString)

	__gong__rebuildMapString(stage.DirectPositionListTypes, &stage.DirectPositionListTypes_mapString)

	__gong__rebuildMapString(stage.DirectPositionTypes, &stage.DirectPositionTypes_mapString)

	__gong__rebuildMapString(stage.DynamicFeatureTypes, &stage.DynamicFeatureTypes_mapString)

	__gong__rebuildMapString(stage.EnvelopeTypes, &stage.EnvelopeTypes_mapString)

	__gong__rebuildMapString(stage.GeodesicStringTypes, &stage.GeodesicStringTypes_mapString)

	__gong__rebuildMapString(stage.GeodesicTypes, &stage.GeodesicTypes_mapString)

	__gong__rebuildMapString(stage.LengthTypes, &stage.LengthTypes_mapString)

	__gong__rebuildMapString(stage.LineStringSegmentTypes, &stage.LineStringSegmentTypes_mapString)

	__gong__rebuildMapString(stage.MeasureTypes, &stage.MeasureTypes_mapString)

	__gong__rebuildMapString(stage.OrientableCurveTypes, &stage.OrientableCurveTypes_mapString)

	__gong__rebuildMapString(stage.PointPropertyTypes, &stage.PointPropertyTypes_mapString)

	__gong__rebuildMapString(stage.PointTypes, &stage.PointTypes_mapString)

	__gong__rebuildMapString(stage.PolygonPatchTypes, &stage.PolygonPatchTypes_mapString)

	__gong__rebuildMapString(stage.RingTypes, &stage.RingTypes_mapString)

	__gong__rebuildMapString(stage.SurfacePatchArrayPropertyTypes, &stage.SurfacePatchArrayPropertyTypes_mapString)

	__gong__rebuildMapString(stage.SurfaceTypes, &stage.SurfaceTypes_mapString)

	__gong__rebuildMapString(stage.TimeInstantTypes, &stage.TimeInstantTypes_mapString)

	__gong__rebuildMapString(stage.TimePeriodTypes, &stage.TimePeriodTypes_mapString)

	__gong__rebuildMapString(stage.TimePositionTypes, &stage.TimePositionTypes_mapString)

	__gong__rebuildMapString(stage.TimePrimitivePropertyTypes, &stage.TimePrimitivePropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
