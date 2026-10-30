// generated code - do not edit
package spatialrepresentation

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
	AbstractMD_SpatialRepresentation_Types                map[*AbstractMD_SpatialRepresentation_Type]struct{}
	AbstractMD_SpatialRepresentation_Types_instance       map[*AbstractMD_SpatialRepresentation_Type]*AbstractMD_SpatialRepresentation_Type
	AbstractMD_SpatialRepresentation_Types_mapString      map[string]*AbstractMD_SpatialRepresentation_Type
	AbstractMD_SpatialRepresentation_TypeOrder            uint
	AbstractMD_SpatialRepresentation_Type_stagedOrder     map[*AbstractMD_SpatialRepresentation_Type]uint
	AbstractMD_SpatialRepresentation_Type_orderStaged     map[uint]*AbstractMD_SpatialRepresentation_Type
	AbstractMD_SpatialRepresentation_Types_reference      map[*AbstractMD_SpatialRepresentation_Type]*AbstractMD_SpatialRepresentation_Type
	AbstractMD_SpatialRepresentation_Types_referenceOrder map[*AbstractMD_SpatialRepresentation_Type]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractMD_SpatialRepresentation_TypeCreateCallback GongOnAfterCreateInterface[AbstractMD_SpatialRepresentation_Type]
	OnAfterAbstractMD_SpatialRepresentation_TypeUpdateCallback GongOnAfterUpdateInterface[AbstractMD_SpatialRepresentation_Type]
	OnAfterAbstractMD_SpatialRepresentation_TypeDeleteCallback GongOnAfterDeleteInterface[AbstractMD_SpatialRepresentation_Type]

	MD_CellGeometryCode_PropertyTypes                map[*MD_CellGeometryCode_PropertyType]struct{}
	MD_CellGeometryCode_PropertyTypes_instance       map[*MD_CellGeometryCode_PropertyType]*MD_CellGeometryCode_PropertyType
	MD_CellGeometryCode_PropertyTypes_mapString      map[string]*MD_CellGeometryCode_PropertyType
	MD_CellGeometryCode_PropertyTypeOrder            uint
	MD_CellGeometryCode_PropertyType_stagedOrder     map[*MD_CellGeometryCode_PropertyType]uint
	MD_CellGeometryCode_PropertyType_orderStaged     map[uint]*MD_CellGeometryCode_PropertyType
	MD_CellGeometryCode_PropertyTypes_reference      map[*MD_CellGeometryCode_PropertyType]*MD_CellGeometryCode_PropertyType
	MD_CellGeometryCode_PropertyTypes_referenceOrder map[*MD_CellGeometryCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_CellGeometryCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_CellGeometryCode_PropertyType]
	OnAfterMD_CellGeometryCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_CellGeometryCode_PropertyType]
	OnAfterMD_CellGeometryCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_CellGeometryCode_PropertyType]

	MD_DimensionNameTypeCode_PropertyTypes                map[*MD_DimensionNameTypeCode_PropertyType]struct{}
	MD_DimensionNameTypeCode_PropertyTypes_instance       map[*MD_DimensionNameTypeCode_PropertyType]*MD_DimensionNameTypeCode_PropertyType
	MD_DimensionNameTypeCode_PropertyTypes_mapString      map[string]*MD_DimensionNameTypeCode_PropertyType
	MD_DimensionNameTypeCode_PropertyTypeOrder            uint
	MD_DimensionNameTypeCode_PropertyType_stagedOrder     map[*MD_DimensionNameTypeCode_PropertyType]uint
	MD_DimensionNameTypeCode_PropertyType_orderStaged     map[uint]*MD_DimensionNameTypeCode_PropertyType
	MD_DimensionNameTypeCode_PropertyTypes_reference      map[*MD_DimensionNameTypeCode_PropertyType]*MD_DimensionNameTypeCode_PropertyType
	MD_DimensionNameTypeCode_PropertyTypes_referenceOrder map[*MD_DimensionNameTypeCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_DimensionNameTypeCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_DimensionNameTypeCode_PropertyType]
	OnAfterMD_DimensionNameTypeCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_DimensionNameTypeCode_PropertyType]
	OnAfterMD_DimensionNameTypeCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_DimensionNameTypeCode_PropertyType]

	MD_Dimension_PropertyTypes                map[*MD_Dimension_PropertyType]struct{}
	MD_Dimension_PropertyTypes_instance       map[*MD_Dimension_PropertyType]*MD_Dimension_PropertyType
	MD_Dimension_PropertyTypes_mapString      map[string]*MD_Dimension_PropertyType
	MD_Dimension_PropertyTypeOrder            uint
	MD_Dimension_PropertyType_stagedOrder     map[*MD_Dimension_PropertyType]uint
	MD_Dimension_PropertyType_orderStaged     map[uint]*MD_Dimension_PropertyType
	MD_Dimension_PropertyTypes_reference      map[*MD_Dimension_PropertyType]*MD_Dimension_PropertyType
	MD_Dimension_PropertyTypes_referenceOrder map[*MD_Dimension_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Dimension_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_Dimension_PropertyType]
	OnAfterMD_Dimension_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_Dimension_PropertyType]
	OnAfterMD_Dimension_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_Dimension_PropertyType]

	MD_Dimension_Types                map[*MD_Dimension_Type]struct{}
	MD_Dimension_Types_instance       map[*MD_Dimension_Type]*MD_Dimension_Type
	MD_Dimension_Types_mapString      map[string]*MD_Dimension_Type
	MD_Dimension_TypeOrder            uint
	MD_Dimension_Type_stagedOrder     map[*MD_Dimension_Type]uint
	MD_Dimension_Type_orderStaged     map[uint]*MD_Dimension_Type
	MD_Dimension_Types_reference      map[*MD_Dimension_Type]*MD_Dimension_Type
	MD_Dimension_Types_referenceOrder map[*MD_Dimension_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Dimension_TypeCreateCallback GongOnAfterCreateInterface[MD_Dimension_Type]
	OnAfterMD_Dimension_TypeUpdateCallback GongOnAfterUpdateInterface[MD_Dimension_Type]
	OnAfterMD_Dimension_TypeDeleteCallback GongOnAfterDeleteInterface[MD_Dimension_Type]

	MD_GeometricObjectTypeCode_PropertyTypes                map[*MD_GeometricObjectTypeCode_PropertyType]struct{}
	MD_GeometricObjectTypeCode_PropertyTypes_instance       map[*MD_GeometricObjectTypeCode_PropertyType]*MD_GeometricObjectTypeCode_PropertyType
	MD_GeometricObjectTypeCode_PropertyTypes_mapString      map[string]*MD_GeometricObjectTypeCode_PropertyType
	MD_GeometricObjectTypeCode_PropertyTypeOrder            uint
	MD_GeometricObjectTypeCode_PropertyType_stagedOrder     map[*MD_GeometricObjectTypeCode_PropertyType]uint
	MD_GeometricObjectTypeCode_PropertyType_orderStaged     map[uint]*MD_GeometricObjectTypeCode_PropertyType
	MD_GeometricObjectTypeCode_PropertyTypes_reference      map[*MD_GeometricObjectTypeCode_PropertyType]*MD_GeometricObjectTypeCode_PropertyType
	MD_GeometricObjectTypeCode_PropertyTypes_referenceOrder map[*MD_GeometricObjectTypeCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_GeometricObjectTypeCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_GeometricObjectTypeCode_PropertyType]
	OnAfterMD_GeometricObjectTypeCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_GeometricObjectTypeCode_PropertyType]
	OnAfterMD_GeometricObjectTypeCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_GeometricObjectTypeCode_PropertyType]

	MD_GeometricObjects_PropertyTypes                map[*MD_GeometricObjects_PropertyType]struct{}
	MD_GeometricObjects_PropertyTypes_instance       map[*MD_GeometricObjects_PropertyType]*MD_GeometricObjects_PropertyType
	MD_GeometricObjects_PropertyTypes_mapString      map[string]*MD_GeometricObjects_PropertyType
	MD_GeometricObjects_PropertyTypeOrder            uint
	MD_GeometricObjects_PropertyType_stagedOrder     map[*MD_GeometricObjects_PropertyType]uint
	MD_GeometricObjects_PropertyType_orderStaged     map[uint]*MD_GeometricObjects_PropertyType
	MD_GeometricObjects_PropertyTypes_reference      map[*MD_GeometricObjects_PropertyType]*MD_GeometricObjects_PropertyType
	MD_GeometricObjects_PropertyTypes_referenceOrder map[*MD_GeometricObjects_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_GeometricObjects_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_GeometricObjects_PropertyType]
	OnAfterMD_GeometricObjects_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_GeometricObjects_PropertyType]
	OnAfterMD_GeometricObjects_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_GeometricObjects_PropertyType]

	MD_GeometricObjects_Types                map[*MD_GeometricObjects_Type]struct{}
	MD_GeometricObjects_Types_instance       map[*MD_GeometricObjects_Type]*MD_GeometricObjects_Type
	MD_GeometricObjects_Types_mapString      map[string]*MD_GeometricObjects_Type
	MD_GeometricObjects_TypeOrder            uint
	MD_GeometricObjects_Type_stagedOrder     map[*MD_GeometricObjects_Type]uint
	MD_GeometricObjects_Type_orderStaged     map[uint]*MD_GeometricObjects_Type
	MD_GeometricObjects_Types_reference      map[*MD_GeometricObjects_Type]*MD_GeometricObjects_Type
	MD_GeometricObjects_Types_referenceOrder map[*MD_GeometricObjects_Type]uint

	// insertion point for slice of pointers maps
	OnAfterMD_GeometricObjects_TypeCreateCallback GongOnAfterCreateInterface[MD_GeometricObjects_Type]
	OnAfterMD_GeometricObjects_TypeUpdateCallback GongOnAfterUpdateInterface[MD_GeometricObjects_Type]
	OnAfterMD_GeometricObjects_TypeDeleteCallback GongOnAfterDeleteInterface[MD_GeometricObjects_Type]

	MD_Georectified_PropertyTypes                map[*MD_Georectified_PropertyType]struct{}
	MD_Georectified_PropertyTypes_instance       map[*MD_Georectified_PropertyType]*MD_Georectified_PropertyType
	MD_Georectified_PropertyTypes_mapString      map[string]*MD_Georectified_PropertyType
	MD_Georectified_PropertyTypeOrder            uint
	MD_Georectified_PropertyType_stagedOrder     map[*MD_Georectified_PropertyType]uint
	MD_Georectified_PropertyType_orderStaged     map[uint]*MD_Georectified_PropertyType
	MD_Georectified_PropertyTypes_reference      map[*MD_Georectified_PropertyType]*MD_Georectified_PropertyType
	MD_Georectified_PropertyTypes_referenceOrder map[*MD_Georectified_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Georectified_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_Georectified_PropertyType]
	OnAfterMD_Georectified_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_Georectified_PropertyType]
	OnAfterMD_Georectified_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_Georectified_PropertyType]

	MD_Georectified_Types                map[*MD_Georectified_Type]struct{}
	MD_Georectified_Types_instance       map[*MD_Georectified_Type]*MD_Georectified_Type
	MD_Georectified_Types_mapString      map[string]*MD_Georectified_Type
	MD_Georectified_TypeOrder            uint
	MD_Georectified_Type_stagedOrder     map[*MD_Georectified_Type]uint
	MD_Georectified_Type_orderStaged     map[uint]*MD_Georectified_Type
	MD_Georectified_Types_reference      map[*MD_Georectified_Type]*MD_Georectified_Type
	MD_Georectified_Types_referenceOrder map[*MD_Georectified_Type]uint

	// insertion point for slice of pointers maps
	MD_Georectified_Type_AxisDimensionProperties_reverseMap map[*MD_Dimension_PropertyType]*MD_Georectified_Type

	OnAfterMD_Georectified_TypeCreateCallback GongOnAfterCreateInterface[MD_Georectified_Type]
	OnAfterMD_Georectified_TypeUpdateCallback GongOnAfterUpdateInterface[MD_Georectified_Type]
	OnAfterMD_Georectified_TypeDeleteCallback GongOnAfterDeleteInterface[MD_Georectified_Type]

	MD_Georeferenceable_PropertyTypes                map[*MD_Georeferenceable_PropertyType]struct{}
	MD_Georeferenceable_PropertyTypes_instance       map[*MD_Georeferenceable_PropertyType]*MD_Georeferenceable_PropertyType
	MD_Georeferenceable_PropertyTypes_mapString      map[string]*MD_Georeferenceable_PropertyType
	MD_Georeferenceable_PropertyTypeOrder            uint
	MD_Georeferenceable_PropertyType_stagedOrder     map[*MD_Georeferenceable_PropertyType]uint
	MD_Georeferenceable_PropertyType_orderStaged     map[uint]*MD_Georeferenceable_PropertyType
	MD_Georeferenceable_PropertyTypes_reference      map[*MD_Georeferenceable_PropertyType]*MD_Georeferenceable_PropertyType
	MD_Georeferenceable_PropertyTypes_referenceOrder map[*MD_Georeferenceable_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Georeferenceable_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_Georeferenceable_PropertyType]
	OnAfterMD_Georeferenceable_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_Georeferenceable_PropertyType]
	OnAfterMD_Georeferenceable_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_Georeferenceable_PropertyType]

	MD_Georeferenceable_Types                map[*MD_Georeferenceable_Type]struct{}
	MD_Georeferenceable_Types_instance       map[*MD_Georeferenceable_Type]*MD_Georeferenceable_Type
	MD_Georeferenceable_Types_mapString      map[string]*MD_Georeferenceable_Type
	MD_Georeferenceable_TypeOrder            uint
	MD_Georeferenceable_Type_stagedOrder     map[*MD_Georeferenceable_Type]uint
	MD_Georeferenceable_Type_orderStaged     map[uint]*MD_Georeferenceable_Type
	MD_Georeferenceable_Types_reference      map[*MD_Georeferenceable_Type]*MD_Georeferenceable_Type
	MD_Georeferenceable_Types_referenceOrder map[*MD_Georeferenceable_Type]uint

	// insertion point for slice of pointers maps
	MD_Georeferenceable_Type_AxisDimensionProperties_reverseMap map[*MD_Dimension_PropertyType]*MD_Georeferenceable_Type

	OnAfterMD_Georeferenceable_TypeCreateCallback GongOnAfterCreateInterface[MD_Georeferenceable_Type]
	OnAfterMD_Georeferenceable_TypeUpdateCallback GongOnAfterUpdateInterface[MD_Georeferenceable_Type]
	OnAfterMD_Georeferenceable_TypeDeleteCallback GongOnAfterDeleteInterface[MD_Georeferenceable_Type]

	MD_GridSpatialRepresentation_PropertyTypes                map[*MD_GridSpatialRepresentation_PropertyType]struct{}
	MD_GridSpatialRepresentation_PropertyTypes_instance       map[*MD_GridSpatialRepresentation_PropertyType]*MD_GridSpatialRepresentation_PropertyType
	MD_GridSpatialRepresentation_PropertyTypes_mapString      map[string]*MD_GridSpatialRepresentation_PropertyType
	MD_GridSpatialRepresentation_PropertyTypeOrder            uint
	MD_GridSpatialRepresentation_PropertyType_stagedOrder     map[*MD_GridSpatialRepresentation_PropertyType]uint
	MD_GridSpatialRepresentation_PropertyType_orderStaged     map[uint]*MD_GridSpatialRepresentation_PropertyType
	MD_GridSpatialRepresentation_PropertyTypes_reference      map[*MD_GridSpatialRepresentation_PropertyType]*MD_GridSpatialRepresentation_PropertyType
	MD_GridSpatialRepresentation_PropertyTypes_referenceOrder map[*MD_GridSpatialRepresentation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_GridSpatialRepresentation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_GridSpatialRepresentation_PropertyType]
	OnAfterMD_GridSpatialRepresentation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_GridSpatialRepresentation_PropertyType]
	OnAfterMD_GridSpatialRepresentation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_GridSpatialRepresentation_PropertyType]

	MD_GridSpatialRepresentation_Types                map[*MD_GridSpatialRepresentation_Type]struct{}
	MD_GridSpatialRepresentation_Types_instance       map[*MD_GridSpatialRepresentation_Type]*MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation_Types_mapString      map[string]*MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation_TypeOrder            uint
	MD_GridSpatialRepresentation_Type_stagedOrder     map[*MD_GridSpatialRepresentation_Type]uint
	MD_GridSpatialRepresentation_Type_orderStaged     map[uint]*MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation_Types_reference      map[*MD_GridSpatialRepresentation_Type]*MD_GridSpatialRepresentation_Type
	MD_GridSpatialRepresentation_Types_referenceOrder map[*MD_GridSpatialRepresentation_Type]uint

	// insertion point for slice of pointers maps
	MD_GridSpatialRepresentation_Type_AxisDimensionProperties_reverseMap map[*MD_Dimension_PropertyType]*MD_GridSpatialRepresentation_Type

	OnAfterMD_GridSpatialRepresentation_TypeCreateCallback GongOnAfterCreateInterface[MD_GridSpatialRepresentation_Type]
	OnAfterMD_GridSpatialRepresentation_TypeUpdateCallback GongOnAfterUpdateInterface[MD_GridSpatialRepresentation_Type]
	OnAfterMD_GridSpatialRepresentation_TypeDeleteCallback GongOnAfterDeleteInterface[MD_GridSpatialRepresentation_Type]

	MD_PixelOrientationCode_PropertyTypes                map[*MD_PixelOrientationCode_PropertyType]struct{}
	MD_PixelOrientationCode_PropertyTypes_instance       map[*MD_PixelOrientationCode_PropertyType]*MD_PixelOrientationCode_PropertyType
	MD_PixelOrientationCode_PropertyTypes_mapString      map[string]*MD_PixelOrientationCode_PropertyType
	MD_PixelOrientationCode_PropertyTypeOrder            uint
	MD_PixelOrientationCode_PropertyType_stagedOrder     map[*MD_PixelOrientationCode_PropertyType]uint
	MD_PixelOrientationCode_PropertyType_orderStaged     map[uint]*MD_PixelOrientationCode_PropertyType
	MD_PixelOrientationCode_PropertyTypes_reference      map[*MD_PixelOrientationCode_PropertyType]*MD_PixelOrientationCode_PropertyType
	MD_PixelOrientationCode_PropertyTypes_referenceOrder map[*MD_PixelOrientationCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_PixelOrientationCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_PixelOrientationCode_PropertyType]
	OnAfterMD_PixelOrientationCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_PixelOrientationCode_PropertyType]
	OnAfterMD_PixelOrientationCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_PixelOrientationCode_PropertyType]

	MD_SpatialRepresentation_PropertyTypes                map[*MD_SpatialRepresentation_PropertyType]struct{}
	MD_SpatialRepresentation_PropertyTypes_instance       map[*MD_SpatialRepresentation_PropertyType]*MD_SpatialRepresentation_PropertyType
	MD_SpatialRepresentation_PropertyTypes_mapString      map[string]*MD_SpatialRepresentation_PropertyType
	MD_SpatialRepresentation_PropertyTypeOrder            uint
	MD_SpatialRepresentation_PropertyType_stagedOrder     map[*MD_SpatialRepresentation_PropertyType]uint
	MD_SpatialRepresentation_PropertyType_orderStaged     map[uint]*MD_SpatialRepresentation_PropertyType
	MD_SpatialRepresentation_PropertyTypes_reference      map[*MD_SpatialRepresentation_PropertyType]*MD_SpatialRepresentation_PropertyType
	MD_SpatialRepresentation_PropertyTypes_referenceOrder map[*MD_SpatialRepresentation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_SpatialRepresentation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_SpatialRepresentation_PropertyType]
	OnAfterMD_SpatialRepresentation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_SpatialRepresentation_PropertyType]
	OnAfterMD_SpatialRepresentation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_SpatialRepresentation_PropertyType]

	MD_TopologyLevelCode_PropertyTypes                map[*MD_TopologyLevelCode_PropertyType]struct{}
	MD_TopologyLevelCode_PropertyTypes_instance       map[*MD_TopologyLevelCode_PropertyType]*MD_TopologyLevelCode_PropertyType
	MD_TopologyLevelCode_PropertyTypes_mapString      map[string]*MD_TopologyLevelCode_PropertyType
	MD_TopologyLevelCode_PropertyTypeOrder            uint
	MD_TopologyLevelCode_PropertyType_stagedOrder     map[*MD_TopologyLevelCode_PropertyType]uint
	MD_TopologyLevelCode_PropertyType_orderStaged     map[uint]*MD_TopologyLevelCode_PropertyType
	MD_TopologyLevelCode_PropertyTypes_reference      map[*MD_TopologyLevelCode_PropertyType]*MD_TopologyLevelCode_PropertyType
	MD_TopologyLevelCode_PropertyTypes_referenceOrder map[*MD_TopologyLevelCode_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_TopologyLevelCode_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_TopologyLevelCode_PropertyType]
	OnAfterMD_TopologyLevelCode_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_TopologyLevelCode_PropertyType]
	OnAfterMD_TopologyLevelCode_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_TopologyLevelCode_PropertyType]

	MD_VectorSpatialRepresentation_PropertyTypes                map[*MD_VectorSpatialRepresentation_PropertyType]struct{}
	MD_VectorSpatialRepresentation_PropertyTypes_instance       map[*MD_VectorSpatialRepresentation_PropertyType]*MD_VectorSpatialRepresentation_PropertyType
	MD_VectorSpatialRepresentation_PropertyTypes_mapString      map[string]*MD_VectorSpatialRepresentation_PropertyType
	MD_VectorSpatialRepresentation_PropertyTypeOrder            uint
	MD_VectorSpatialRepresentation_PropertyType_stagedOrder     map[*MD_VectorSpatialRepresentation_PropertyType]uint
	MD_VectorSpatialRepresentation_PropertyType_orderStaged     map[uint]*MD_VectorSpatialRepresentation_PropertyType
	MD_VectorSpatialRepresentation_PropertyTypes_reference      map[*MD_VectorSpatialRepresentation_PropertyType]*MD_VectorSpatialRepresentation_PropertyType
	MD_VectorSpatialRepresentation_PropertyTypes_referenceOrder map[*MD_VectorSpatialRepresentation_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_VectorSpatialRepresentation_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_VectorSpatialRepresentation_PropertyType]
	OnAfterMD_VectorSpatialRepresentation_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_VectorSpatialRepresentation_PropertyType]
	OnAfterMD_VectorSpatialRepresentation_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_VectorSpatialRepresentation_PropertyType]

	MD_VectorSpatialRepresentation_Types                map[*MD_VectorSpatialRepresentation_Type]struct{}
	MD_VectorSpatialRepresentation_Types_instance       map[*MD_VectorSpatialRepresentation_Type]*MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation_Types_mapString      map[string]*MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation_TypeOrder            uint
	MD_VectorSpatialRepresentation_Type_stagedOrder     map[*MD_VectorSpatialRepresentation_Type]uint
	MD_VectorSpatialRepresentation_Type_orderStaged     map[uint]*MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation_Types_reference      map[*MD_VectorSpatialRepresentation_Type]*MD_VectorSpatialRepresentation_Type
	MD_VectorSpatialRepresentation_Types_referenceOrder map[*MD_VectorSpatialRepresentation_Type]uint

	// insertion point for slice of pointers maps
	MD_VectorSpatialRepresentation_Type_GeometricObjects_reverseMap map[*MD_GeometricObjects_PropertyType]*MD_VectorSpatialRepresentation_Type

	OnAfterMD_VectorSpatialRepresentation_TypeCreateCallback GongOnAfterCreateInterface[MD_VectorSpatialRepresentation_Type]
	OnAfterMD_VectorSpatialRepresentation_TypeUpdateCallback GongOnAfterUpdateInterface[MD_VectorSpatialRepresentation_Type]
	OnAfterMD_VectorSpatialRepresentation_TypeDeleteCallback GongOnAfterDeleteInterface[MD_VectorSpatialRepresentation_Type]

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
	__gong__clearReferences(&stage.AbstractMD_SpatialRepresentation_Types_reference, &stage.AbstractMD_SpatialRepresentation_Types_instance, &stage.AbstractMD_SpatialRepresentation_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_CellGeometryCode_PropertyTypes_reference, &stage.MD_CellGeometryCode_PropertyTypes_instance, &stage.MD_CellGeometryCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_DimensionNameTypeCode_PropertyTypes_reference, &stage.MD_DimensionNameTypeCode_PropertyTypes_instance, &stage.MD_DimensionNameTypeCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Dimension_PropertyTypes_reference, &stage.MD_Dimension_PropertyTypes_instance, &stage.MD_Dimension_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Dimension_Types_reference, &stage.MD_Dimension_Types_instance, &stage.MD_Dimension_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_GeometricObjectTypeCode_PropertyTypes_reference, &stage.MD_GeometricObjectTypeCode_PropertyTypes_instance, &stage.MD_GeometricObjectTypeCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_GeometricObjects_PropertyTypes_reference, &stage.MD_GeometricObjects_PropertyTypes_instance, &stage.MD_GeometricObjects_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_GeometricObjects_Types_reference, &stage.MD_GeometricObjects_Types_instance, &stage.MD_GeometricObjects_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_Georectified_PropertyTypes_reference, &stage.MD_Georectified_PropertyTypes_instance, &stage.MD_Georectified_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Georectified_Types_reference, &stage.MD_Georectified_Types_instance, &stage.MD_Georectified_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_Georeferenceable_PropertyTypes_reference, &stage.MD_Georeferenceable_PropertyTypes_instance, &stage.MD_Georeferenceable_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Georeferenceable_Types_reference, &stage.MD_Georeferenceable_Types_instance, &stage.MD_Georeferenceable_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_GridSpatialRepresentation_PropertyTypes_reference, &stage.MD_GridSpatialRepresentation_PropertyTypes_instance, &stage.MD_GridSpatialRepresentation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_GridSpatialRepresentation_Types_reference, &stage.MD_GridSpatialRepresentation_Types_instance, &stage.MD_GridSpatialRepresentation_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_PixelOrientationCode_PropertyTypes_reference, &stage.MD_PixelOrientationCode_PropertyTypes_instance, &stage.MD_PixelOrientationCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_SpatialRepresentation_PropertyTypes_reference, &stage.MD_SpatialRepresentation_PropertyTypes_instance, &stage.MD_SpatialRepresentation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_TopologyLevelCode_PropertyTypes_reference, &stage.MD_TopologyLevelCode_PropertyTypes_instance, &stage.MD_TopologyLevelCode_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_VectorSpatialRepresentation_PropertyTypes_reference, &stage.MD_VectorSpatialRepresentation_PropertyTypes_instance, &stage.MD_VectorSpatialRepresentation_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_VectorSpatialRepresentation_Types_reference, &stage.MD_VectorSpatialRepresentation_Types_instance, &stage.MD_VectorSpatialRepresentation_Types_referenceOrder)

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
	stage.AbstractMD_SpatialRepresentation_TypeOrder = __gong__recomputeOrder(stage.AbstractMD_SpatialRepresentation_Type_stagedOrder)

	stage.MD_CellGeometryCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_CellGeometryCode_PropertyType_stagedOrder)

	stage.MD_DimensionNameTypeCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder)

	stage.MD_Dimension_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_Dimension_PropertyType_stagedOrder)

	stage.MD_Dimension_TypeOrder = __gong__recomputeOrder(stage.MD_Dimension_Type_stagedOrder)

	stage.MD_GeometricObjectTypeCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder)

	stage.MD_GeometricObjects_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_GeometricObjects_PropertyType_stagedOrder)

	stage.MD_GeometricObjects_TypeOrder = __gong__recomputeOrder(stage.MD_GeometricObjects_Type_stagedOrder)

	stage.MD_Georectified_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_Georectified_PropertyType_stagedOrder)

	stage.MD_Georectified_TypeOrder = __gong__recomputeOrder(stage.MD_Georectified_Type_stagedOrder)

	stage.MD_Georeferenceable_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_Georeferenceable_PropertyType_stagedOrder)

	stage.MD_Georeferenceable_TypeOrder = __gong__recomputeOrder(stage.MD_Georeferenceable_Type_stagedOrder)

	stage.MD_GridSpatialRepresentation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder)

	stage.MD_GridSpatialRepresentation_TypeOrder = __gong__recomputeOrder(stage.MD_GridSpatialRepresentation_Type_stagedOrder)

	stage.MD_PixelOrientationCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_PixelOrientationCode_PropertyType_stagedOrder)

	stage.MD_SpatialRepresentation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_SpatialRepresentation_PropertyType_stagedOrder)

	stage.MD_TopologyLevelCode_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_TopologyLevelCode_PropertyType_stagedOrder)

	stage.MD_VectorSpatialRepresentation_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder)

	stage.MD_VectorSpatialRepresentation_TypeOrder = __gong__recomputeOrder(stage.MD_VectorSpatialRepresentation_Type_stagedOrder)

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
func (*AbstractMD_SpatialRepresentation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Type_stagedOrder)
}

func (*AbstractMD_SpatialRepresentation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractMD_SpatialRepresentation_Type_orderStaged[order]
}

func (*AbstractMD_SpatialRepresentation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractMD_SpatialRepresentation_Types_mapString
}

func (*AbstractMD_SpatialRepresentation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractMD_SpatialRepresentation_Types
}

func (*AbstractMD_SpatialRepresentation_Type) GongNewInstance() any {
	return new(AbstractMD_SpatialRepresentation_Type)
}

func (*MD_CellGeometryCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyType_stagedOrder)
}

func (*MD_CellGeometryCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_CellGeometryCode_PropertyType_orderStaged[order]
}

func (*MD_CellGeometryCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_CellGeometryCode_PropertyTypes_mapString
}

func (*MD_CellGeometryCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_CellGeometryCode_PropertyTypes
}

func (*MD_CellGeometryCode_PropertyType) GongNewInstance() any {
	return new(MD_CellGeometryCode_PropertyType)
}

func (*MD_DimensionNameTypeCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder)
}

func (*MD_DimensionNameTypeCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_DimensionNameTypeCode_PropertyType_orderStaged[order]
}

func (*MD_DimensionNameTypeCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_DimensionNameTypeCode_PropertyTypes_mapString
}

func (*MD_DimensionNameTypeCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_DimensionNameTypeCode_PropertyTypes
}

func (*MD_DimensionNameTypeCode_PropertyType) GongNewInstance() any {
	return new(MD_DimensionNameTypeCode_PropertyType)
}

func (*MD_Dimension_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyType_stagedOrder)
}

func (*MD_Dimension_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Dimension_PropertyType_orderStaged[order]
}

func (*MD_Dimension_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Dimension_PropertyTypes_mapString
}

func (*MD_Dimension_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Dimension_PropertyTypes
}

func (*MD_Dimension_PropertyType) GongNewInstance() any {
	return new(MD_Dimension_PropertyType)
}

func (*MD_Dimension_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Dimension_Types, stage.MD_Dimension_Type_stagedOrder)
}

func (*MD_Dimension_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Dimension_Type_orderStaged[order]
}

func (*MD_Dimension_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Dimension_Types_mapString
}

func (*MD_Dimension_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Dimension_Types
}

func (*MD_Dimension_Type) GongNewInstance() any {
	return new(MD_Dimension_Type)
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder)
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_GeometricObjectTypeCode_PropertyType_orderStaged[order]
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_GeometricObjectTypeCode_PropertyTypes
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongNewInstance() any {
	return new(MD_GeometricObjectTypeCode_PropertyType)
}

func (*MD_GeometricObjects_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyType_stagedOrder)
}

func (*MD_GeometricObjects_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_GeometricObjects_PropertyType_orderStaged[order]
}

func (*MD_GeometricObjects_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_GeometricObjects_PropertyTypes_mapString
}

func (*MD_GeometricObjects_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_GeometricObjects_PropertyTypes
}

func (*MD_GeometricObjects_PropertyType) GongNewInstance() any {
	return new(MD_GeometricObjects_PropertyType)
}

func (*MD_GeometricObjects_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Type_stagedOrder)
}

func (*MD_GeometricObjects_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_GeometricObjects_Type_orderStaged[order]
}

func (*MD_GeometricObjects_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_GeometricObjects_Types_mapString
}

func (*MD_GeometricObjects_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_GeometricObjects_Types
}

func (*MD_GeometricObjects_Type) GongNewInstance() any {
	return new(MD_GeometricObjects_Type)
}

func (*MD_Georectified_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyType_stagedOrder)
}

func (*MD_Georectified_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Georectified_PropertyType_orderStaged[order]
}

func (*MD_Georectified_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Georectified_PropertyTypes_mapString
}

func (*MD_Georectified_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Georectified_PropertyTypes
}

func (*MD_Georectified_PropertyType) GongNewInstance() any {
	return new(MD_Georectified_PropertyType)
}

func (*MD_Georectified_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Georectified_Types, stage.MD_Georectified_Type_stagedOrder)
}

func (*MD_Georectified_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Georectified_Type_orderStaged[order]
}

func (*MD_Georectified_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Georectified_Types_mapString
}

func (*MD_Georectified_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Georectified_Types
}

func (*MD_Georectified_Type) GongNewInstance() any {
	return new(MD_Georectified_Type)
}

func (*MD_Georeferenceable_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyType_stagedOrder)
}

func (*MD_Georeferenceable_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Georeferenceable_PropertyType_orderStaged[order]
}

func (*MD_Georeferenceable_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Georeferenceable_PropertyTypes_mapString
}

func (*MD_Georeferenceable_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Georeferenceable_PropertyTypes
}

func (*MD_Georeferenceable_PropertyType) GongNewInstance() any {
	return new(MD_Georeferenceable_PropertyType)
}

func (*MD_Georeferenceable_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Type_stagedOrder)
}

func (*MD_Georeferenceable_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Georeferenceable_Type_orderStaged[order]
}

func (*MD_Georeferenceable_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Georeferenceable_Types_mapString
}

func (*MD_Georeferenceable_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Georeferenceable_Types
}

func (*MD_Georeferenceable_Type) GongNewInstance() any {
	return new(MD_Georeferenceable_Type)
}

func (*MD_GridSpatialRepresentation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder)
}

func (*MD_GridSpatialRepresentation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_GridSpatialRepresentation_PropertyType_orderStaged[order]
}

func (*MD_GridSpatialRepresentation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_GridSpatialRepresentation_PropertyTypes_mapString
}

func (*MD_GridSpatialRepresentation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_GridSpatialRepresentation_PropertyTypes
}

func (*MD_GridSpatialRepresentation_PropertyType) GongNewInstance() any {
	return new(MD_GridSpatialRepresentation_PropertyType)
}

func (*MD_GridSpatialRepresentation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Type_stagedOrder)
}

func (*MD_GridSpatialRepresentation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_GridSpatialRepresentation_Type_orderStaged[order]
}

func (*MD_GridSpatialRepresentation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_GridSpatialRepresentation_Types_mapString
}

func (*MD_GridSpatialRepresentation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_GridSpatialRepresentation_Types
}

func (*MD_GridSpatialRepresentation_Type) GongNewInstance() any {
	return new(MD_GridSpatialRepresentation_Type)
}

func (*MD_PixelOrientationCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyType_stagedOrder)
}

func (*MD_PixelOrientationCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_PixelOrientationCode_PropertyType_orderStaged[order]
}

func (*MD_PixelOrientationCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_PixelOrientationCode_PropertyTypes_mapString
}

func (*MD_PixelOrientationCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_PixelOrientationCode_PropertyTypes
}

func (*MD_PixelOrientationCode_PropertyType) GongNewInstance() any {
	return new(MD_PixelOrientationCode_PropertyType)
}

func (*MD_SpatialRepresentation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyType_stagedOrder)
}

func (*MD_SpatialRepresentation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_SpatialRepresentation_PropertyType_orderStaged[order]
}

func (*MD_SpatialRepresentation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_SpatialRepresentation_PropertyTypes_mapString
}

func (*MD_SpatialRepresentation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_SpatialRepresentation_PropertyTypes
}

func (*MD_SpatialRepresentation_PropertyType) GongNewInstance() any {
	return new(MD_SpatialRepresentation_PropertyType)
}

func (*MD_TopologyLevelCode_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyType_stagedOrder)
}

func (*MD_TopologyLevelCode_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_TopologyLevelCode_PropertyType_orderStaged[order]
}

func (*MD_TopologyLevelCode_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_TopologyLevelCode_PropertyTypes_mapString
}

func (*MD_TopologyLevelCode_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_TopologyLevelCode_PropertyTypes
}

func (*MD_TopologyLevelCode_PropertyType) GongNewInstance() any {
	return new(MD_TopologyLevelCode_PropertyType)
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder)
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_VectorSpatialRepresentation_PropertyType_orderStaged[order]
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_VectorSpatialRepresentation_PropertyTypes
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongNewInstance() any {
	return new(MD_VectorSpatialRepresentation_PropertyType)
}

func (*MD_VectorSpatialRepresentation_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Type_stagedOrder)
}

func (*MD_VectorSpatialRepresentation_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_VectorSpatialRepresentation_Type_orderStaged[order]
}

func (*MD_VectorSpatialRepresentation_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_VectorSpatialRepresentation_Types_mapString
}

func (*MD_VectorSpatialRepresentation_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_VectorSpatialRepresentation_Types
}

func (*MD_VectorSpatialRepresentation_Type) GongNewInstance() any {
	return new(MD_VectorSpatialRepresentation_Type)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation"
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
		AbstractMD_SpatialRepresentation_Types:           make(map[*AbstractMD_SpatialRepresentation_Type]struct{}),
		AbstractMD_SpatialRepresentation_Types_mapString: make(map[string]*AbstractMD_SpatialRepresentation_Type),

		MD_CellGeometryCode_PropertyTypes:           make(map[*MD_CellGeometryCode_PropertyType]struct{}),
		MD_CellGeometryCode_PropertyTypes_mapString: make(map[string]*MD_CellGeometryCode_PropertyType),

		MD_DimensionNameTypeCode_PropertyTypes:           make(map[*MD_DimensionNameTypeCode_PropertyType]struct{}),
		MD_DimensionNameTypeCode_PropertyTypes_mapString: make(map[string]*MD_DimensionNameTypeCode_PropertyType),

		MD_Dimension_PropertyTypes:           make(map[*MD_Dimension_PropertyType]struct{}),
		MD_Dimension_PropertyTypes_mapString: make(map[string]*MD_Dimension_PropertyType),

		MD_Dimension_Types:           make(map[*MD_Dimension_Type]struct{}),
		MD_Dimension_Types_mapString: make(map[string]*MD_Dimension_Type),

		MD_GeometricObjectTypeCode_PropertyTypes:           make(map[*MD_GeometricObjectTypeCode_PropertyType]struct{}),
		MD_GeometricObjectTypeCode_PropertyTypes_mapString: make(map[string]*MD_GeometricObjectTypeCode_PropertyType),

		MD_GeometricObjects_PropertyTypes:           make(map[*MD_GeometricObjects_PropertyType]struct{}),
		MD_GeometricObjects_PropertyTypes_mapString: make(map[string]*MD_GeometricObjects_PropertyType),

		MD_GeometricObjects_Types:           make(map[*MD_GeometricObjects_Type]struct{}),
		MD_GeometricObjects_Types_mapString: make(map[string]*MD_GeometricObjects_Type),

		MD_Georectified_PropertyTypes:           make(map[*MD_Georectified_PropertyType]struct{}),
		MD_Georectified_PropertyTypes_mapString: make(map[string]*MD_Georectified_PropertyType),

		MD_Georectified_Types:           make(map[*MD_Georectified_Type]struct{}),
		MD_Georectified_Types_mapString: make(map[string]*MD_Georectified_Type),

		MD_Georeferenceable_PropertyTypes:           make(map[*MD_Georeferenceable_PropertyType]struct{}),
		MD_Georeferenceable_PropertyTypes_mapString: make(map[string]*MD_Georeferenceable_PropertyType),

		MD_Georeferenceable_Types:           make(map[*MD_Georeferenceable_Type]struct{}),
		MD_Georeferenceable_Types_mapString: make(map[string]*MD_Georeferenceable_Type),

		MD_GridSpatialRepresentation_PropertyTypes:           make(map[*MD_GridSpatialRepresentation_PropertyType]struct{}),
		MD_GridSpatialRepresentation_PropertyTypes_mapString: make(map[string]*MD_GridSpatialRepresentation_PropertyType),

		MD_GridSpatialRepresentation_Types:           make(map[*MD_GridSpatialRepresentation_Type]struct{}),
		MD_GridSpatialRepresentation_Types_mapString: make(map[string]*MD_GridSpatialRepresentation_Type),

		MD_PixelOrientationCode_PropertyTypes:           make(map[*MD_PixelOrientationCode_PropertyType]struct{}),
		MD_PixelOrientationCode_PropertyTypes_mapString: make(map[string]*MD_PixelOrientationCode_PropertyType),

		MD_SpatialRepresentation_PropertyTypes:           make(map[*MD_SpatialRepresentation_PropertyType]struct{}),
		MD_SpatialRepresentation_PropertyTypes_mapString: make(map[string]*MD_SpatialRepresentation_PropertyType),

		MD_TopologyLevelCode_PropertyTypes:           make(map[*MD_TopologyLevelCode_PropertyType]struct{}),
		MD_TopologyLevelCode_PropertyTypes_mapString: make(map[string]*MD_TopologyLevelCode_PropertyType),

		MD_VectorSpatialRepresentation_PropertyTypes:           make(map[*MD_VectorSpatialRepresentation_PropertyType]struct{}),
		MD_VectorSpatialRepresentation_PropertyTypes_mapString: make(map[string]*MD_VectorSpatialRepresentation_PropertyType),

		MD_VectorSpatialRepresentation_Types:           make(map[*MD_VectorSpatialRepresentation_Type]struct{}),
		MD_VectorSpatialRepresentation_Types_mapString: make(map[string]*MD_VectorSpatialRepresentation_Type),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractMD_SpatialRepresentation_Type_stagedOrder: make(map[*AbstractMD_SpatialRepresentation_Type]uint),
		AbstractMD_SpatialRepresentation_Type_orderStaged: make(map[uint]*AbstractMD_SpatialRepresentation_Type),
		AbstractMD_SpatialRepresentation_Types_reference:  make(map[*AbstractMD_SpatialRepresentation_Type]*AbstractMD_SpatialRepresentation_Type),

		MD_CellGeometryCode_PropertyType_stagedOrder: make(map[*MD_CellGeometryCode_PropertyType]uint),
		MD_CellGeometryCode_PropertyType_orderStaged: make(map[uint]*MD_CellGeometryCode_PropertyType),
		MD_CellGeometryCode_PropertyTypes_reference:  make(map[*MD_CellGeometryCode_PropertyType]*MD_CellGeometryCode_PropertyType),

		MD_DimensionNameTypeCode_PropertyType_stagedOrder: make(map[*MD_DimensionNameTypeCode_PropertyType]uint),
		MD_DimensionNameTypeCode_PropertyType_orderStaged: make(map[uint]*MD_DimensionNameTypeCode_PropertyType),
		MD_DimensionNameTypeCode_PropertyTypes_reference:  make(map[*MD_DimensionNameTypeCode_PropertyType]*MD_DimensionNameTypeCode_PropertyType),

		MD_Dimension_PropertyType_stagedOrder: make(map[*MD_Dimension_PropertyType]uint),
		MD_Dimension_PropertyType_orderStaged: make(map[uint]*MD_Dimension_PropertyType),
		MD_Dimension_PropertyTypes_reference:  make(map[*MD_Dimension_PropertyType]*MD_Dimension_PropertyType),

		MD_Dimension_Type_stagedOrder: make(map[*MD_Dimension_Type]uint),
		MD_Dimension_Type_orderStaged: make(map[uint]*MD_Dimension_Type),
		MD_Dimension_Types_reference:  make(map[*MD_Dimension_Type]*MD_Dimension_Type),

		MD_GeometricObjectTypeCode_PropertyType_stagedOrder: make(map[*MD_GeometricObjectTypeCode_PropertyType]uint),
		MD_GeometricObjectTypeCode_PropertyType_orderStaged: make(map[uint]*MD_GeometricObjectTypeCode_PropertyType),
		MD_GeometricObjectTypeCode_PropertyTypes_reference:  make(map[*MD_GeometricObjectTypeCode_PropertyType]*MD_GeometricObjectTypeCode_PropertyType),

		MD_GeometricObjects_PropertyType_stagedOrder: make(map[*MD_GeometricObjects_PropertyType]uint),
		MD_GeometricObjects_PropertyType_orderStaged: make(map[uint]*MD_GeometricObjects_PropertyType),
		MD_GeometricObjects_PropertyTypes_reference:  make(map[*MD_GeometricObjects_PropertyType]*MD_GeometricObjects_PropertyType),

		MD_GeometricObjects_Type_stagedOrder: make(map[*MD_GeometricObjects_Type]uint),
		MD_GeometricObjects_Type_orderStaged: make(map[uint]*MD_GeometricObjects_Type),
		MD_GeometricObjects_Types_reference:  make(map[*MD_GeometricObjects_Type]*MD_GeometricObjects_Type),

		MD_Georectified_PropertyType_stagedOrder: make(map[*MD_Georectified_PropertyType]uint),
		MD_Georectified_PropertyType_orderStaged: make(map[uint]*MD_Georectified_PropertyType),
		MD_Georectified_PropertyTypes_reference:  make(map[*MD_Georectified_PropertyType]*MD_Georectified_PropertyType),

		MD_Georectified_Type_stagedOrder: make(map[*MD_Georectified_Type]uint),
		MD_Georectified_Type_orderStaged: make(map[uint]*MD_Georectified_Type),
		MD_Georectified_Types_reference:  make(map[*MD_Georectified_Type]*MD_Georectified_Type),

		MD_Georeferenceable_PropertyType_stagedOrder: make(map[*MD_Georeferenceable_PropertyType]uint),
		MD_Georeferenceable_PropertyType_orderStaged: make(map[uint]*MD_Georeferenceable_PropertyType),
		MD_Georeferenceable_PropertyTypes_reference:  make(map[*MD_Georeferenceable_PropertyType]*MD_Georeferenceable_PropertyType),

		MD_Georeferenceable_Type_stagedOrder: make(map[*MD_Georeferenceable_Type]uint),
		MD_Georeferenceable_Type_orderStaged: make(map[uint]*MD_Georeferenceable_Type),
		MD_Georeferenceable_Types_reference:  make(map[*MD_Georeferenceable_Type]*MD_Georeferenceable_Type),

		MD_GridSpatialRepresentation_PropertyType_stagedOrder: make(map[*MD_GridSpatialRepresentation_PropertyType]uint),
		MD_GridSpatialRepresentation_PropertyType_orderStaged: make(map[uint]*MD_GridSpatialRepresentation_PropertyType),
		MD_GridSpatialRepresentation_PropertyTypes_reference:  make(map[*MD_GridSpatialRepresentation_PropertyType]*MD_GridSpatialRepresentation_PropertyType),

		MD_GridSpatialRepresentation_Type_stagedOrder: make(map[*MD_GridSpatialRepresentation_Type]uint),
		MD_GridSpatialRepresentation_Type_orderStaged: make(map[uint]*MD_GridSpatialRepresentation_Type),
		MD_GridSpatialRepresentation_Types_reference:  make(map[*MD_GridSpatialRepresentation_Type]*MD_GridSpatialRepresentation_Type),

		MD_PixelOrientationCode_PropertyType_stagedOrder: make(map[*MD_PixelOrientationCode_PropertyType]uint),
		MD_PixelOrientationCode_PropertyType_orderStaged: make(map[uint]*MD_PixelOrientationCode_PropertyType),
		MD_PixelOrientationCode_PropertyTypes_reference:  make(map[*MD_PixelOrientationCode_PropertyType]*MD_PixelOrientationCode_PropertyType),

		MD_SpatialRepresentation_PropertyType_stagedOrder: make(map[*MD_SpatialRepresentation_PropertyType]uint),
		MD_SpatialRepresentation_PropertyType_orderStaged: make(map[uint]*MD_SpatialRepresentation_PropertyType),
		MD_SpatialRepresentation_PropertyTypes_reference:  make(map[*MD_SpatialRepresentation_PropertyType]*MD_SpatialRepresentation_PropertyType),

		MD_TopologyLevelCode_PropertyType_stagedOrder: make(map[*MD_TopologyLevelCode_PropertyType]uint),
		MD_TopologyLevelCode_PropertyType_orderStaged: make(map[uint]*MD_TopologyLevelCode_PropertyType),
		MD_TopologyLevelCode_PropertyTypes_reference:  make(map[*MD_TopologyLevelCode_PropertyType]*MD_TopologyLevelCode_PropertyType),

		MD_VectorSpatialRepresentation_PropertyType_stagedOrder: make(map[*MD_VectorSpatialRepresentation_PropertyType]uint),
		MD_VectorSpatialRepresentation_PropertyType_orderStaged: make(map[uint]*MD_VectorSpatialRepresentation_PropertyType),
		MD_VectorSpatialRepresentation_PropertyTypes_reference:  make(map[*MD_VectorSpatialRepresentation_PropertyType]*MD_VectorSpatialRepresentation_PropertyType),

		MD_VectorSpatialRepresentation_Type_stagedOrder: make(map[*MD_VectorSpatialRepresentation_Type]uint),
		MD_VectorSpatialRepresentation_Type_orderStaged: make(map[uint]*MD_VectorSpatialRepresentation_Type),
		MD_VectorSpatialRepresentation_Types_reference:  make(map[*MD_VectorSpatialRepresentation_Type]*MD_VectorSpatialRepresentation_Type),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractMD_SpatialRepresentation_Type": &AbstractMD_SpatialRepresentation_TypeUnmarshaller{},

			"MD_CellGeometryCode_PropertyType": &MD_CellGeometryCode_PropertyTypeUnmarshaller{},

			"MD_DimensionNameTypeCode_PropertyType": &MD_DimensionNameTypeCode_PropertyTypeUnmarshaller{},

			"MD_Dimension_PropertyType": &MD_Dimension_PropertyTypeUnmarshaller{},

			"MD_Dimension_Type": &MD_Dimension_TypeUnmarshaller{},

			"MD_GeometricObjectTypeCode_PropertyType": &MD_GeometricObjectTypeCode_PropertyTypeUnmarshaller{},

			"MD_GeometricObjects_PropertyType": &MD_GeometricObjects_PropertyTypeUnmarshaller{},

			"MD_GeometricObjects_Type": &MD_GeometricObjects_TypeUnmarshaller{},

			"MD_Georectified_PropertyType": &MD_Georectified_PropertyTypeUnmarshaller{},

			"MD_Georectified_Type": &MD_Georectified_TypeUnmarshaller{},

			"MD_Georeferenceable_PropertyType": &MD_Georeferenceable_PropertyTypeUnmarshaller{},

			"MD_Georeferenceable_Type": &MD_Georeferenceable_TypeUnmarshaller{},

			"MD_GridSpatialRepresentation_PropertyType": &MD_GridSpatialRepresentation_PropertyTypeUnmarshaller{},

			"MD_GridSpatialRepresentation_Type": &MD_GridSpatialRepresentation_TypeUnmarshaller{},

			"MD_PixelOrientationCode_PropertyType": &MD_PixelOrientationCode_PropertyTypeUnmarshaller{},

			"MD_SpatialRepresentation_PropertyType": &MD_SpatialRepresentation_PropertyTypeUnmarshaller{},

			"MD_TopologyLevelCode_PropertyType": &MD_TopologyLevelCode_PropertyTypeUnmarshaller{},

			"MD_VectorSpatialRepresentation_PropertyType": &MD_VectorSpatialRepresentation_PropertyTypeUnmarshaller{},

			"MD_VectorSpatialRepresentation_Type": &MD_VectorSpatialRepresentation_TypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractMD_SpatialRepresentation_Type"] = len(stage.AbstractMD_SpatialRepresentation_Types)
	stage.Map_GongStructName_InstancesNb["MD_CellGeometryCode_PropertyType"] = len(stage.MD_CellGeometryCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_DimensionNameTypeCode_PropertyType"] = len(stage.MD_DimensionNameTypeCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Dimension_PropertyType"] = len(stage.MD_Dimension_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Dimension_Type"] = len(stage.MD_Dimension_Types)
	stage.Map_GongStructName_InstancesNb["MD_GeometricObjectTypeCode_PropertyType"] = len(stage.MD_GeometricObjectTypeCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_GeometricObjects_PropertyType"] = len(stage.MD_GeometricObjects_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_GeometricObjects_Type"] = len(stage.MD_GeometricObjects_Types)
	stage.Map_GongStructName_InstancesNb["MD_Georectified_PropertyType"] = len(stage.MD_Georectified_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Georectified_Type"] = len(stage.MD_Georectified_Types)
	stage.Map_GongStructName_InstancesNb["MD_Georeferenceable_PropertyType"] = len(stage.MD_Georeferenceable_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Georeferenceable_Type"] = len(stage.MD_Georeferenceable_Types)
	stage.Map_GongStructName_InstancesNb["MD_GridSpatialRepresentation_PropertyType"] = len(stage.MD_GridSpatialRepresentation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_GridSpatialRepresentation_Type"] = len(stage.MD_GridSpatialRepresentation_Types)
	stage.Map_GongStructName_InstancesNb["MD_PixelOrientationCode_PropertyType"] = len(stage.MD_PixelOrientationCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_SpatialRepresentation_PropertyType"] = len(stage.MD_SpatialRepresentation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_TopologyLevelCode_PropertyType"] = len(stage.MD_TopologyLevelCode_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_VectorSpatialRepresentation_PropertyType"] = len(stage.MD_VectorSpatialRepresentation_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_VectorSpatialRepresentation_Type"] = len(stage.MD_VectorSpatialRepresentation_Types)
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
// Stage puts abstractmd_spatialrepresentation_type to the model stage
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) Stage(stage *Stage) *AbstractMD_SpatialRepresentation_Type {
	__gong__stage(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Type_stagedOrder, stage.AbstractMD_SpatialRepresentation_Type_orderStaged, &stage.AbstractMD_SpatialRepresentation_TypeOrder, stage.AbstractMD_SpatialRepresentation_Types_mapString, abstractmd_spatialrepresentation_type, abstractmd_spatialrepresentation_type.Name)
	return abstractmd_spatialrepresentation_type
}

// StagePreserveOrder puts abstractmd_spatialrepresentation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractMD_SpatialRepresentation_TypeOrder
// - update stage.AbstractMD_SpatialRepresentation_TypeOrder accordingly
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Type_stagedOrder, stage.AbstractMD_SpatialRepresentation_Type_orderStaged, &stage.AbstractMD_SpatialRepresentation_TypeOrder, stage.AbstractMD_SpatialRepresentation_Types_mapString, abstractmd_spatialrepresentation_type, order, abstractmd_spatialrepresentation_type.Name)
}

// Unstage removes abstractmd_spatialrepresentation_type off the model stage
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) Unstage(stage *Stage) *AbstractMD_SpatialRepresentation_Type {
	__gong__unstage(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Types_mapString, abstractmd_spatialrepresentation_type, abstractmd_spatialrepresentation_type.Name)
	return abstractmd_spatialrepresentation_type
}

// UnstageVoid removes abstractmd_spatialrepresentation_type off the model stage
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) UnstageVoid(stage *Stage) {
	abstractmd_spatialrepresentation_type.Unstage(stage)
}

func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) StageVoid(stage *Stage) {
	abstractmd_spatialrepresentation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GetName() (res string) {
	return abstractmd_spatialrepresentation_type.Name
}

// for satisfaction of GongStruct interface
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) SetName(name string) {
	abstractmd_spatialrepresentation_type.Name = name
}

// Stage puts md_cellgeometrycode_propertytype to the model stage
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) Stage(stage *Stage) *MD_CellGeometryCode_PropertyType {
	__gong__stage(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyType_stagedOrder, stage.MD_CellGeometryCode_PropertyType_orderStaged, &stage.MD_CellGeometryCode_PropertyTypeOrder, stage.MD_CellGeometryCode_PropertyTypes_mapString, md_cellgeometrycode_propertytype, md_cellgeometrycode_propertytype.Name)
	return md_cellgeometrycode_propertytype
}

// StagePreserveOrder puts md_cellgeometrycode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_CellGeometryCode_PropertyTypeOrder
// - update stage.MD_CellGeometryCode_PropertyTypeOrder accordingly
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyType_stagedOrder, stage.MD_CellGeometryCode_PropertyType_orderStaged, &stage.MD_CellGeometryCode_PropertyTypeOrder, stage.MD_CellGeometryCode_PropertyTypes_mapString, md_cellgeometrycode_propertytype, order, md_cellgeometrycode_propertytype.Name)
}

// Unstage removes md_cellgeometrycode_propertytype off the model stage
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) Unstage(stage *Stage) *MD_CellGeometryCode_PropertyType {
	__gong__unstage(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyTypes_mapString, md_cellgeometrycode_propertytype, md_cellgeometrycode_propertytype.Name)
	return md_cellgeometrycode_propertytype
}

// UnstageVoid removes md_cellgeometrycode_propertytype off the model stage
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) UnstageVoid(stage *Stage) {
	md_cellgeometrycode_propertytype.Unstage(stage)
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) StageVoid(stage *Stage) {
	md_cellgeometrycode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GetName() (res string) {
	return md_cellgeometrycode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) SetName(name string) {
	md_cellgeometrycode_propertytype.Name = name
}

// Stage puts md_dimensionnametypecode_propertytype to the model stage
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) Stage(stage *Stage) *MD_DimensionNameTypeCode_PropertyType {
	__gong__stage(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder, stage.MD_DimensionNameTypeCode_PropertyType_orderStaged, &stage.MD_DimensionNameTypeCode_PropertyTypeOrder, stage.MD_DimensionNameTypeCode_PropertyTypes_mapString, md_dimensionnametypecode_propertytype, md_dimensionnametypecode_propertytype.Name)
	return md_dimensionnametypecode_propertytype
}

// StagePreserveOrder puts md_dimensionnametypecode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_DimensionNameTypeCode_PropertyTypeOrder
// - update stage.MD_DimensionNameTypeCode_PropertyTypeOrder accordingly
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder, stage.MD_DimensionNameTypeCode_PropertyType_orderStaged, &stage.MD_DimensionNameTypeCode_PropertyTypeOrder, stage.MD_DimensionNameTypeCode_PropertyTypes_mapString, md_dimensionnametypecode_propertytype, order, md_dimensionnametypecode_propertytype.Name)
}

// Unstage removes md_dimensionnametypecode_propertytype off the model stage
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) Unstage(stage *Stage) *MD_DimensionNameTypeCode_PropertyType {
	__gong__unstage(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyTypes_mapString, md_dimensionnametypecode_propertytype, md_dimensionnametypecode_propertytype.Name)
	return md_dimensionnametypecode_propertytype
}

// UnstageVoid removes md_dimensionnametypecode_propertytype off the model stage
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) UnstageVoid(stage *Stage) {
	md_dimensionnametypecode_propertytype.Unstage(stage)
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) StageVoid(stage *Stage) {
	md_dimensionnametypecode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GetName() (res string) {
	return md_dimensionnametypecode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) SetName(name string) {
	md_dimensionnametypecode_propertytype.Name = name
}

// Stage puts md_dimension_propertytype to the model stage
func (md_dimension_propertytype *MD_Dimension_PropertyType) Stage(stage *Stage) *MD_Dimension_PropertyType {
	__gong__stage(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyType_stagedOrder, stage.MD_Dimension_PropertyType_orderStaged, &stage.MD_Dimension_PropertyTypeOrder, stage.MD_Dimension_PropertyTypes_mapString, md_dimension_propertytype, md_dimension_propertytype.Name)
	return md_dimension_propertytype
}

// StagePreserveOrder puts md_dimension_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Dimension_PropertyTypeOrder
// - update stage.MD_Dimension_PropertyTypeOrder accordingly
func (md_dimension_propertytype *MD_Dimension_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyType_stagedOrder, stage.MD_Dimension_PropertyType_orderStaged, &stage.MD_Dimension_PropertyTypeOrder, stage.MD_Dimension_PropertyTypes_mapString, md_dimension_propertytype, order, md_dimension_propertytype.Name)
}

// Unstage removes md_dimension_propertytype off the model stage
func (md_dimension_propertytype *MD_Dimension_PropertyType) Unstage(stage *Stage) *MD_Dimension_PropertyType {
	__gong__unstage(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyTypes_mapString, md_dimension_propertytype, md_dimension_propertytype.Name)
	return md_dimension_propertytype
}

// UnstageVoid removes md_dimension_propertytype off the model stage
func (md_dimension_propertytype *MD_Dimension_PropertyType) UnstageVoid(stage *Stage) {
	md_dimension_propertytype.Unstage(stage)
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) StageVoid(stage *Stage) {
	md_dimension_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_dimension_propertytype *MD_Dimension_PropertyType) GetName() (res string) {
	return md_dimension_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_dimension_propertytype *MD_Dimension_PropertyType) SetName(name string) {
	md_dimension_propertytype.Name = name
}

// Stage puts md_dimension_type to the model stage
func (md_dimension_type *MD_Dimension_Type) Stage(stage *Stage) *MD_Dimension_Type {
	__gong__stage(stage.MD_Dimension_Types, stage.MD_Dimension_Type_stagedOrder, stage.MD_Dimension_Type_orderStaged, &stage.MD_Dimension_TypeOrder, stage.MD_Dimension_Types_mapString, md_dimension_type, md_dimension_type.Name)
	return md_dimension_type
}

// StagePreserveOrder puts md_dimension_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Dimension_TypeOrder
// - update stage.MD_Dimension_TypeOrder accordingly
func (md_dimension_type *MD_Dimension_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Dimension_Types, stage.MD_Dimension_Type_stagedOrder, stage.MD_Dimension_Type_orderStaged, &stage.MD_Dimension_TypeOrder, stage.MD_Dimension_Types_mapString, md_dimension_type, order, md_dimension_type.Name)
}

// Unstage removes md_dimension_type off the model stage
func (md_dimension_type *MD_Dimension_Type) Unstage(stage *Stage) *MD_Dimension_Type {
	__gong__unstage(stage.MD_Dimension_Types, stage.MD_Dimension_Types_mapString, md_dimension_type, md_dimension_type.Name)
	return md_dimension_type
}

// UnstageVoid removes md_dimension_type off the model stage
func (md_dimension_type *MD_Dimension_Type) UnstageVoid(stage *Stage) {
	md_dimension_type.Unstage(stage)
}

func (md_dimension_type *MD_Dimension_Type) StageVoid(stage *Stage) {
	md_dimension_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_dimension_type *MD_Dimension_Type) GetName() (res string) {
	return md_dimension_type.Name
}

// for satisfaction of GongStruct interface
func (md_dimension_type *MD_Dimension_Type) SetName(name string) {
	md_dimension_type.Name = name
}

// Stage puts md_geometricobjecttypecode_propertytype to the model stage
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) Stage(stage *Stage) *MD_GeometricObjectTypeCode_PropertyType {
	__gong__stage(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder, stage.MD_GeometricObjectTypeCode_PropertyType_orderStaged, &stage.MD_GeometricObjectTypeCode_PropertyTypeOrder, stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString, md_geometricobjecttypecode_propertytype, md_geometricobjecttypecode_propertytype.Name)
	return md_geometricobjecttypecode_propertytype
}

// StagePreserveOrder puts md_geometricobjecttypecode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_GeometricObjectTypeCode_PropertyTypeOrder
// - update stage.MD_GeometricObjectTypeCode_PropertyTypeOrder accordingly
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder, stage.MD_GeometricObjectTypeCode_PropertyType_orderStaged, &stage.MD_GeometricObjectTypeCode_PropertyTypeOrder, stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString, md_geometricobjecttypecode_propertytype, order, md_geometricobjecttypecode_propertytype.Name)
}

// Unstage removes md_geometricobjecttypecode_propertytype off the model stage
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) Unstage(stage *Stage) *MD_GeometricObjectTypeCode_PropertyType {
	__gong__unstage(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString, md_geometricobjecttypecode_propertytype, md_geometricobjecttypecode_propertytype.Name)
	return md_geometricobjecttypecode_propertytype
}

// UnstageVoid removes md_geometricobjecttypecode_propertytype off the model stage
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) UnstageVoid(stage *Stage) {
	md_geometricobjecttypecode_propertytype.Unstage(stage)
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) StageVoid(stage *Stage) {
	md_geometricobjecttypecode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GetName() (res string) {
	return md_geometricobjecttypecode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) SetName(name string) {
	md_geometricobjecttypecode_propertytype.Name = name
}

// Stage puts md_geometricobjects_propertytype to the model stage
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) Stage(stage *Stage) *MD_GeometricObjects_PropertyType {
	__gong__stage(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyType_stagedOrder, stage.MD_GeometricObjects_PropertyType_orderStaged, &stage.MD_GeometricObjects_PropertyTypeOrder, stage.MD_GeometricObjects_PropertyTypes_mapString, md_geometricobjects_propertytype, md_geometricobjects_propertytype.Name)
	return md_geometricobjects_propertytype
}

// StagePreserveOrder puts md_geometricobjects_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_GeometricObjects_PropertyTypeOrder
// - update stage.MD_GeometricObjects_PropertyTypeOrder accordingly
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyType_stagedOrder, stage.MD_GeometricObjects_PropertyType_orderStaged, &stage.MD_GeometricObjects_PropertyTypeOrder, stage.MD_GeometricObjects_PropertyTypes_mapString, md_geometricobjects_propertytype, order, md_geometricobjects_propertytype.Name)
}

// Unstage removes md_geometricobjects_propertytype off the model stage
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) Unstage(stage *Stage) *MD_GeometricObjects_PropertyType {
	__gong__unstage(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyTypes_mapString, md_geometricobjects_propertytype, md_geometricobjects_propertytype.Name)
	return md_geometricobjects_propertytype
}

// UnstageVoid removes md_geometricobjects_propertytype off the model stage
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) UnstageVoid(stage *Stage) {
	md_geometricobjects_propertytype.Unstage(stage)
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) StageVoid(stage *Stage) {
	md_geometricobjects_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GetName() (res string) {
	return md_geometricobjects_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) SetName(name string) {
	md_geometricobjects_propertytype.Name = name
}

// Stage puts md_geometricobjects_type to the model stage
func (md_geometricobjects_type *MD_GeometricObjects_Type) Stage(stage *Stage) *MD_GeometricObjects_Type {
	__gong__stage(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Type_stagedOrder, stage.MD_GeometricObjects_Type_orderStaged, &stage.MD_GeometricObjects_TypeOrder, stage.MD_GeometricObjects_Types_mapString, md_geometricobjects_type, md_geometricobjects_type.Name)
	return md_geometricobjects_type
}

// StagePreserveOrder puts md_geometricobjects_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_GeometricObjects_TypeOrder
// - update stage.MD_GeometricObjects_TypeOrder accordingly
func (md_geometricobjects_type *MD_GeometricObjects_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Type_stagedOrder, stage.MD_GeometricObjects_Type_orderStaged, &stage.MD_GeometricObjects_TypeOrder, stage.MD_GeometricObjects_Types_mapString, md_geometricobjects_type, order, md_geometricobjects_type.Name)
}

// Unstage removes md_geometricobjects_type off the model stage
func (md_geometricobjects_type *MD_GeometricObjects_Type) Unstage(stage *Stage) *MD_GeometricObjects_Type {
	__gong__unstage(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Types_mapString, md_geometricobjects_type, md_geometricobjects_type.Name)
	return md_geometricobjects_type
}

// UnstageVoid removes md_geometricobjects_type off the model stage
func (md_geometricobjects_type *MD_GeometricObjects_Type) UnstageVoid(stage *Stage) {
	md_geometricobjects_type.Unstage(stage)
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) StageVoid(stage *Stage) {
	md_geometricobjects_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_geometricobjects_type *MD_GeometricObjects_Type) GetName() (res string) {
	return md_geometricobjects_type.Name
}

// for satisfaction of GongStruct interface
func (md_geometricobjects_type *MD_GeometricObjects_Type) SetName(name string) {
	md_geometricobjects_type.Name = name
}

// Stage puts md_georectified_propertytype to the model stage
func (md_georectified_propertytype *MD_Georectified_PropertyType) Stage(stage *Stage) *MD_Georectified_PropertyType {
	__gong__stage(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyType_stagedOrder, stage.MD_Georectified_PropertyType_orderStaged, &stage.MD_Georectified_PropertyTypeOrder, stage.MD_Georectified_PropertyTypes_mapString, md_georectified_propertytype, md_georectified_propertytype.Name)
	return md_georectified_propertytype
}

// StagePreserveOrder puts md_georectified_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Georectified_PropertyTypeOrder
// - update stage.MD_Georectified_PropertyTypeOrder accordingly
func (md_georectified_propertytype *MD_Georectified_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyType_stagedOrder, stage.MD_Georectified_PropertyType_orderStaged, &stage.MD_Georectified_PropertyTypeOrder, stage.MD_Georectified_PropertyTypes_mapString, md_georectified_propertytype, order, md_georectified_propertytype.Name)
}

// Unstage removes md_georectified_propertytype off the model stage
func (md_georectified_propertytype *MD_Georectified_PropertyType) Unstage(stage *Stage) *MD_Georectified_PropertyType {
	__gong__unstage(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyTypes_mapString, md_georectified_propertytype, md_georectified_propertytype.Name)
	return md_georectified_propertytype
}

// UnstageVoid removes md_georectified_propertytype off the model stage
func (md_georectified_propertytype *MD_Georectified_PropertyType) UnstageVoid(stage *Stage) {
	md_georectified_propertytype.Unstage(stage)
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) StageVoid(stage *Stage) {
	md_georectified_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_georectified_propertytype *MD_Georectified_PropertyType) GetName() (res string) {
	return md_georectified_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_georectified_propertytype *MD_Georectified_PropertyType) SetName(name string) {
	md_georectified_propertytype.Name = name
}

// Stage puts md_georectified_type to the model stage
func (md_georectified_type *MD_Georectified_Type) Stage(stage *Stage) *MD_Georectified_Type {
	__gong__stage(stage.MD_Georectified_Types, stage.MD_Georectified_Type_stagedOrder, stage.MD_Georectified_Type_orderStaged, &stage.MD_Georectified_TypeOrder, stage.MD_Georectified_Types_mapString, md_georectified_type, md_georectified_type.Name)
	return md_georectified_type
}

// StagePreserveOrder puts md_georectified_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Georectified_TypeOrder
// - update stage.MD_Georectified_TypeOrder accordingly
func (md_georectified_type *MD_Georectified_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Georectified_Types, stage.MD_Georectified_Type_stagedOrder, stage.MD_Georectified_Type_orderStaged, &stage.MD_Georectified_TypeOrder, stage.MD_Georectified_Types_mapString, md_georectified_type, order, md_georectified_type.Name)
}

// Unstage removes md_georectified_type off the model stage
func (md_georectified_type *MD_Georectified_Type) Unstage(stage *Stage) *MD_Georectified_Type {
	__gong__unstage(stage.MD_Georectified_Types, stage.MD_Georectified_Types_mapString, md_georectified_type, md_georectified_type.Name)
	return md_georectified_type
}

// UnstageVoid removes md_georectified_type off the model stage
func (md_georectified_type *MD_Georectified_Type) UnstageVoid(stage *Stage) {
	md_georectified_type.Unstage(stage)
}

func (md_georectified_type *MD_Georectified_Type) StageVoid(stage *Stage) {
	md_georectified_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_georectified_type *MD_Georectified_Type) GetName() (res string) {
	return md_georectified_type.Name
}

// for satisfaction of GongStruct interface
func (md_georectified_type *MD_Georectified_Type) SetName(name string) {
	md_georectified_type.Name = name
}

// Stage puts md_georeferenceable_propertytype to the model stage
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) Stage(stage *Stage) *MD_Georeferenceable_PropertyType {
	__gong__stage(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyType_stagedOrder, stage.MD_Georeferenceable_PropertyType_orderStaged, &stage.MD_Georeferenceable_PropertyTypeOrder, stage.MD_Georeferenceable_PropertyTypes_mapString, md_georeferenceable_propertytype, md_georeferenceable_propertytype.Name)
	return md_georeferenceable_propertytype
}

// StagePreserveOrder puts md_georeferenceable_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Georeferenceable_PropertyTypeOrder
// - update stage.MD_Georeferenceable_PropertyTypeOrder accordingly
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyType_stagedOrder, stage.MD_Georeferenceable_PropertyType_orderStaged, &stage.MD_Georeferenceable_PropertyTypeOrder, stage.MD_Georeferenceable_PropertyTypes_mapString, md_georeferenceable_propertytype, order, md_georeferenceable_propertytype.Name)
}

// Unstage removes md_georeferenceable_propertytype off the model stage
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) Unstage(stage *Stage) *MD_Georeferenceable_PropertyType {
	__gong__unstage(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyTypes_mapString, md_georeferenceable_propertytype, md_georeferenceable_propertytype.Name)
	return md_georeferenceable_propertytype
}

// UnstageVoid removes md_georeferenceable_propertytype off the model stage
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) UnstageVoid(stage *Stage) {
	md_georeferenceable_propertytype.Unstage(stage)
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) StageVoid(stage *Stage) {
	md_georeferenceable_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GetName() (res string) {
	return md_georeferenceable_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) SetName(name string) {
	md_georeferenceable_propertytype.Name = name
}

// Stage puts md_georeferenceable_type to the model stage
func (md_georeferenceable_type *MD_Georeferenceable_Type) Stage(stage *Stage) *MD_Georeferenceable_Type {
	__gong__stage(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Type_stagedOrder, stage.MD_Georeferenceable_Type_orderStaged, &stage.MD_Georeferenceable_TypeOrder, stage.MD_Georeferenceable_Types_mapString, md_georeferenceable_type, md_georeferenceable_type.Name)
	return md_georeferenceable_type
}

// StagePreserveOrder puts md_georeferenceable_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Georeferenceable_TypeOrder
// - update stage.MD_Georeferenceable_TypeOrder accordingly
func (md_georeferenceable_type *MD_Georeferenceable_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Type_stagedOrder, stage.MD_Georeferenceable_Type_orderStaged, &stage.MD_Georeferenceable_TypeOrder, stage.MD_Georeferenceable_Types_mapString, md_georeferenceable_type, order, md_georeferenceable_type.Name)
}

// Unstage removes md_georeferenceable_type off the model stage
func (md_georeferenceable_type *MD_Georeferenceable_Type) Unstage(stage *Stage) *MD_Georeferenceable_Type {
	__gong__unstage(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Types_mapString, md_georeferenceable_type, md_georeferenceable_type.Name)
	return md_georeferenceable_type
}

// UnstageVoid removes md_georeferenceable_type off the model stage
func (md_georeferenceable_type *MD_Georeferenceable_Type) UnstageVoid(stage *Stage) {
	md_georeferenceable_type.Unstage(stage)
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) StageVoid(stage *Stage) {
	md_georeferenceable_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_georeferenceable_type *MD_Georeferenceable_Type) GetName() (res string) {
	return md_georeferenceable_type.Name
}

// for satisfaction of GongStruct interface
func (md_georeferenceable_type *MD_Georeferenceable_Type) SetName(name string) {
	md_georeferenceable_type.Name = name
}

// Stage puts md_gridspatialrepresentation_propertytype to the model stage
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) Stage(stage *Stage) *MD_GridSpatialRepresentation_PropertyType {
	__gong__stage(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder, stage.MD_GridSpatialRepresentation_PropertyType_orderStaged, &stage.MD_GridSpatialRepresentation_PropertyTypeOrder, stage.MD_GridSpatialRepresentation_PropertyTypes_mapString, md_gridspatialrepresentation_propertytype, md_gridspatialrepresentation_propertytype.Name)
	return md_gridspatialrepresentation_propertytype
}

// StagePreserveOrder puts md_gridspatialrepresentation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_GridSpatialRepresentation_PropertyTypeOrder
// - update stage.MD_GridSpatialRepresentation_PropertyTypeOrder accordingly
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder, stage.MD_GridSpatialRepresentation_PropertyType_orderStaged, &stage.MD_GridSpatialRepresentation_PropertyTypeOrder, stage.MD_GridSpatialRepresentation_PropertyTypes_mapString, md_gridspatialrepresentation_propertytype, order, md_gridspatialrepresentation_propertytype.Name)
}

// Unstage removes md_gridspatialrepresentation_propertytype off the model stage
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) Unstage(stage *Stage) *MD_GridSpatialRepresentation_PropertyType {
	__gong__unstage(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyTypes_mapString, md_gridspatialrepresentation_propertytype, md_gridspatialrepresentation_propertytype.Name)
	return md_gridspatialrepresentation_propertytype
}

// UnstageVoid removes md_gridspatialrepresentation_propertytype off the model stage
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) UnstageVoid(stage *Stage) {
	md_gridspatialrepresentation_propertytype.Unstage(stage)
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) StageVoid(stage *Stage) {
	md_gridspatialrepresentation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GetName() (res string) {
	return md_gridspatialrepresentation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) SetName(name string) {
	md_gridspatialrepresentation_propertytype.Name = name
}

// Stage puts md_gridspatialrepresentation_type to the model stage
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) Stage(stage *Stage) *MD_GridSpatialRepresentation_Type {
	__gong__stage(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Type_stagedOrder, stage.MD_GridSpatialRepresentation_Type_orderStaged, &stage.MD_GridSpatialRepresentation_TypeOrder, stage.MD_GridSpatialRepresentation_Types_mapString, md_gridspatialrepresentation_type, md_gridspatialrepresentation_type.Name)
	return md_gridspatialrepresentation_type
}

// StagePreserveOrder puts md_gridspatialrepresentation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_GridSpatialRepresentation_TypeOrder
// - update stage.MD_GridSpatialRepresentation_TypeOrder accordingly
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Type_stagedOrder, stage.MD_GridSpatialRepresentation_Type_orderStaged, &stage.MD_GridSpatialRepresentation_TypeOrder, stage.MD_GridSpatialRepresentation_Types_mapString, md_gridspatialrepresentation_type, order, md_gridspatialrepresentation_type.Name)
}

// Unstage removes md_gridspatialrepresentation_type off the model stage
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) Unstage(stage *Stage) *MD_GridSpatialRepresentation_Type {
	__gong__unstage(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Types_mapString, md_gridspatialrepresentation_type, md_gridspatialrepresentation_type.Name)
	return md_gridspatialrepresentation_type
}

// UnstageVoid removes md_gridspatialrepresentation_type off the model stage
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) UnstageVoid(stage *Stage) {
	md_gridspatialrepresentation_type.Unstage(stage)
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) StageVoid(stage *Stage) {
	md_gridspatialrepresentation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GetName() (res string) {
	return md_gridspatialrepresentation_type.Name
}

// for satisfaction of GongStruct interface
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) SetName(name string) {
	md_gridspatialrepresentation_type.Name = name
}

// Stage puts md_pixelorientationcode_propertytype to the model stage
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) Stage(stage *Stage) *MD_PixelOrientationCode_PropertyType {
	__gong__stage(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyType_stagedOrder, stage.MD_PixelOrientationCode_PropertyType_orderStaged, &stage.MD_PixelOrientationCode_PropertyTypeOrder, stage.MD_PixelOrientationCode_PropertyTypes_mapString, md_pixelorientationcode_propertytype, md_pixelorientationcode_propertytype.Name)
	return md_pixelorientationcode_propertytype
}

// StagePreserveOrder puts md_pixelorientationcode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_PixelOrientationCode_PropertyTypeOrder
// - update stage.MD_PixelOrientationCode_PropertyTypeOrder accordingly
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyType_stagedOrder, stage.MD_PixelOrientationCode_PropertyType_orderStaged, &stage.MD_PixelOrientationCode_PropertyTypeOrder, stage.MD_PixelOrientationCode_PropertyTypes_mapString, md_pixelorientationcode_propertytype, order, md_pixelorientationcode_propertytype.Name)
}

// Unstage removes md_pixelorientationcode_propertytype off the model stage
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) Unstage(stage *Stage) *MD_PixelOrientationCode_PropertyType {
	__gong__unstage(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyTypes_mapString, md_pixelorientationcode_propertytype, md_pixelorientationcode_propertytype.Name)
	return md_pixelorientationcode_propertytype
}

// UnstageVoid removes md_pixelorientationcode_propertytype off the model stage
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) UnstageVoid(stage *Stage) {
	md_pixelorientationcode_propertytype.Unstage(stage)
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) StageVoid(stage *Stage) {
	md_pixelorientationcode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GetName() (res string) {
	return md_pixelorientationcode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) SetName(name string) {
	md_pixelorientationcode_propertytype.Name = name
}

// Stage puts md_spatialrepresentation_propertytype to the model stage
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) Stage(stage *Stage) *MD_SpatialRepresentation_PropertyType {
	__gong__stage(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyType_stagedOrder, stage.MD_SpatialRepresentation_PropertyType_orderStaged, &stage.MD_SpatialRepresentation_PropertyTypeOrder, stage.MD_SpatialRepresentation_PropertyTypes_mapString, md_spatialrepresentation_propertytype, md_spatialrepresentation_propertytype.Name)
	return md_spatialrepresentation_propertytype
}

// StagePreserveOrder puts md_spatialrepresentation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_SpatialRepresentation_PropertyTypeOrder
// - update stage.MD_SpatialRepresentation_PropertyTypeOrder accordingly
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyType_stagedOrder, stage.MD_SpatialRepresentation_PropertyType_orderStaged, &stage.MD_SpatialRepresentation_PropertyTypeOrder, stage.MD_SpatialRepresentation_PropertyTypes_mapString, md_spatialrepresentation_propertytype, order, md_spatialrepresentation_propertytype.Name)
}

// Unstage removes md_spatialrepresentation_propertytype off the model stage
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) Unstage(stage *Stage) *MD_SpatialRepresentation_PropertyType {
	__gong__unstage(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyTypes_mapString, md_spatialrepresentation_propertytype, md_spatialrepresentation_propertytype.Name)
	return md_spatialrepresentation_propertytype
}

// UnstageVoid removes md_spatialrepresentation_propertytype off the model stage
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) UnstageVoid(stage *Stage) {
	md_spatialrepresentation_propertytype.Unstage(stage)
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) StageVoid(stage *Stage) {
	md_spatialrepresentation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GetName() (res string) {
	return md_spatialrepresentation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) SetName(name string) {
	md_spatialrepresentation_propertytype.Name = name
}

// Stage puts md_topologylevelcode_propertytype to the model stage
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) Stage(stage *Stage) *MD_TopologyLevelCode_PropertyType {
	__gong__stage(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyType_stagedOrder, stage.MD_TopologyLevelCode_PropertyType_orderStaged, &stage.MD_TopologyLevelCode_PropertyTypeOrder, stage.MD_TopologyLevelCode_PropertyTypes_mapString, md_topologylevelcode_propertytype, md_topologylevelcode_propertytype.Name)
	return md_topologylevelcode_propertytype
}

// StagePreserveOrder puts md_topologylevelcode_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_TopologyLevelCode_PropertyTypeOrder
// - update stage.MD_TopologyLevelCode_PropertyTypeOrder accordingly
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyType_stagedOrder, stage.MD_TopologyLevelCode_PropertyType_orderStaged, &stage.MD_TopologyLevelCode_PropertyTypeOrder, stage.MD_TopologyLevelCode_PropertyTypes_mapString, md_topologylevelcode_propertytype, order, md_topologylevelcode_propertytype.Name)
}

// Unstage removes md_topologylevelcode_propertytype off the model stage
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) Unstage(stage *Stage) *MD_TopologyLevelCode_PropertyType {
	__gong__unstage(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyTypes_mapString, md_topologylevelcode_propertytype, md_topologylevelcode_propertytype.Name)
	return md_topologylevelcode_propertytype
}

// UnstageVoid removes md_topologylevelcode_propertytype off the model stage
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) UnstageVoid(stage *Stage) {
	md_topologylevelcode_propertytype.Unstage(stage)
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) StageVoid(stage *Stage) {
	md_topologylevelcode_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GetName() (res string) {
	return md_topologylevelcode_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) SetName(name string) {
	md_topologylevelcode_propertytype.Name = name
}

// Stage puts md_vectorspatialrepresentation_propertytype to the model stage
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) Stage(stage *Stage) *MD_VectorSpatialRepresentation_PropertyType {
	__gong__stage(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder, stage.MD_VectorSpatialRepresentation_PropertyType_orderStaged, &stage.MD_VectorSpatialRepresentation_PropertyTypeOrder, stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString, md_vectorspatialrepresentation_propertytype, md_vectorspatialrepresentation_propertytype.Name)
	return md_vectorspatialrepresentation_propertytype
}

// StagePreserveOrder puts md_vectorspatialrepresentation_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_VectorSpatialRepresentation_PropertyTypeOrder
// - update stage.MD_VectorSpatialRepresentation_PropertyTypeOrder accordingly
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder, stage.MD_VectorSpatialRepresentation_PropertyType_orderStaged, &stage.MD_VectorSpatialRepresentation_PropertyTypeOrder, stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString, md_vectorspatialrepresentation_propertytype, order, md_vectorspatialrepresentation_propertytype.Name)
}

// Unstage removes md_vectorspatialrepresentation_propertytype off the model stage
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) Unstage(stage *Stage) *MD_VectorSpatialRepresentation_PropertyType {
	__gong__unstage(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString, md_vectorspatialrepresentation_propertytype, md_vectorspatialrepresentation_propertytype.Name)
	return md_vectorspatialrepresentation_propertytype
}

// UnstageVoid removes md_vectorspatialrepresentation_propertytype off the model stage
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) UnstageVoid(stage *Stage) {
	md_vectorspatialrepresentation_propertytype.Unstage(stage)
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) StageVoid(stage *Stage) {
	md_vectorspatialrepresentation_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GetName() (res string) {
	return md_vectorspatialrepresentation_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) SetName(name string) {
	md_vectorspatialrepresentation_propertytype.Name = name
}

// Stage puts md_vectorspatialrepresentation_type to the model stage
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) Stage(stage *Stage) *MD_VectorSpatialRepresentation_Type {
	__gong__stage(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Type_stagedOrder, stage.MD_VectorSpatialRepresentation_Type_orderStaged, &stage.MD_VectorSpatialRepresentation_TypeOrder, stage.MD_VectorSpatialRepresentation_Types_mapString, md_vectorspatialrepresentation_type, md_vectorspatialrepresentation_type.Name)
	return md_vectorspatialrepresentation_type
}

// StagePreserveOrder puts md_vectorspatialrepresentation_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_VectorSpatialRepresentation_TypeOrder
// - update stage.MD_VectorSpatialRepresentation_TypeOrder accordingly
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Type_stagedOrder, stage.MD_VectorSpatialRepresentation_Type_orderStaged, &stage.MD_VectorSpatialRepresentation_TypeOrder, stage.MD_VectorSpatialRepresentation_Types_mapString, md_vectorspatialrepresentation_type, order, md_vectorspatialrepresentation_type.Name)
}

// Unstage removes md_vectorspatialrepresentation_type off the model stage
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) Unstage(stage *Stage) *MD_VectorSpatialRepresentation_Type {
	__gong__unstage(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Types_mapString, md_vectorspatialrepresentation_type, md_vectorspatialrepresentation_type.Name)
	return md_vectorspatialrepresentation_type
}

// UnstageVoid removes md_vectorspatialrepresentation_type off the model stage
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) UnstageVoid(stage *Stage) {
	md_vectorspatialrepresentation_type.Unstage(stage)
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) StageVoid(stage *Stage) {
	md_vectorspatialrepresentation_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GetName() (res string) {
	return md_vectorspatialrepresentation_type.Name
}

// for satisfaction of GongStruct interface
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) SetName(name string) {
	md_vectorspatialrepresentation_type.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractMD_SpatialRepresentation_Types, &stage.AbstractMD_SpatialRepresentation_Types_mapString, &stage.AbstractMD_SpatialRepresentation_Type_stagedOrder, &stage.AbstractMD_SpatialRepresentation_TypeOrder)

	__gong__resetStageType(&stage.MD_CellGeometryCode_PropertyTypes, &stage.MD_CellGeometryCode_PropertyTypes_mapString, &stage.MD_CellGeometryCode_PropertyType_stagedOrder, &stage.MD_CellGeometryCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_DimensionNameTypeCode_PropertyTypes, &stage.MD_DimensionNameTypeCode_PropertyTypes_mapString, &stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder, &stage.MD_DimensionNameTypeCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Dimension_PropertyTypes, &stage.MD_Dimension_PropertyTypes_mapString, &stage.MD_Dimension_PropertyType_stagedOrder, &stage.MD_Dimension_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Dimension_Types, &stage.MD_Dimension_Types_mapString, &stage.MD_Dimension_Type_stagedOrder, &stage.MD_Dimension_TypeOrder)

	__gong__resetStageType(&stage.MD_GeometricObjectTypeCode_PropertyTypes, &stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString, &stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder, &stage.MD_GeometricObjectTypeCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_GeometricObjects_PropertyTypes, &stage.MD_GeometricObjects_PropertyTypes_mapString, &stage.MD_GeometricObjects_PropertyType_stagedOrder, &stage.MD_GeometricObjects_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_GeometricObjects_Types, &stage.MD_GeometricObjects_Types_mapString, &stage.MD_GeometricObjects_Type_stagedOrder, &stage.MD_GeometricObjects_TypeOrder)

	__gong__resetStageType(&stage.MD_Georectified_PropertyTypes, &stage.MD_Georectified_PropertyTypes_mapString, &stage.MD_Georectified_PropertyType_stagedOrder, &stage.MD_Georectified_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Georectified_Types, &stage.MD_Georectified_Types_mapString, &stage.MD_Georectified_Type_stagedOrder, &stage.MD_Georectified_TypeOrder)

	__gong__resetStageType(&stage.MD_Georeferenceable_PropertyTypes, &stage.MD_Georeferenceable_PropertyTypes_mapString, &stage.MD_Georeferenceable_PropertyType_stagedOrder, &stage.MD_Georeferenceable_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Georeferenceable_Types, &stage.MD_Georeferenceable_Types_mapString, &stage.MD_Georeferenceable_Type_stagedOrder, &stage.MD_Georeferenceable_TypeOrder)

	__gong__resetStageType(&stage.MD_GridSpatialRepresentation_PropertyTypes, &stage.MD_GridSpatialRepresentation_PropertyTypes_mapString, &stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder, &stage.MD_GridSpatialRepresentation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_GridSpatialRepresentation_Types, &stage.MD_GridSpatialRepresentation_Types_mapString, &stage.MD_GridSpatialRepresentation_Type_stagedOrder, &stage.MD_GridSpatialRepresentation_TypeOrder)

	__gong__resetStageType(&stage.MD_PixelOrientationCode_PropertyTypes, &stage.MD_PixelOrientationCode_PropertyTypes_mapString, &stage.MD_PixelOrientationCode_PropertyType_stagedOrder, &stage.MD_PixelOrientationCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_SpatialRepresentation_PropertyTypes, &stage.MD_SpatialRepresentation_PropertyTypes_mapString, &stage.MD_SpatialRepresentation_PropertyType_stagedOrder, &stage.MD_SpatialRepresentation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_TopologyLevelCode_PropertyTypes, &stage.MD_TopologyLevelCode_PropertyTypes_mapString, &stage.MD_TopologyLevelCode_PropertyType_stagedOrder, &stage.MD_TopologyLevelCode_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_VectorSpatialRepresentation_PropertyTypes, &stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString, &stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder, &stage.MD_VectorSpatialRepresentation_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_VectorSpatialRepresentation_Types, &stage.MD_VectorSpatialRepresentation_Types_mapString, &stage.MD_VectorSpatialRepresentation_Type_stagedOrder, &stage.MD_VectorSpatialRepresentation_TypeOrder)

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
func (AbstractMD_SpatialRepresentation_Type) GongGetAssociationName() any {
	return &AbstractMD_SpatialRepresentation_Type{
	}
}

func (MD_CellGeometryCode_PropertyType) GongGetAssociationName() any {
	return &MD_CellGeometryCode_PropertyType{
	}
}

func (MD_DimensionNameTypeCode_PropertyType) GongGetAssociationName() any {
	return &MD_DimensionNameTypeCode_PropertyType{
	}
}

func (MD_Dimension_PropertyType) GongGetAssociationName() any {
	return &MD_Dimension_PropertyType{
			MD_Dimension: &MD_Dimension_Type{Name: "MD_Dimension"},
	}
}

func (MD_Dimension_Type) GongGetAssociationName() any {
	return &MD_Dimension_Type{
			DimensionName: &MD_DimensionNameTypeCode_PropertyType{Name: "DimensionName"},
	}
}

func (MD_GeometricObjectTypeCode_PropertyType) GongGetAssociationName() any {
	return &MD_GeometricObjectTypeCode_PropertyType{
	}
}

func (MD_GeometricObjects_PropertyType) GongGetAssociationName() any {
	return &MD_GeometricObjects_PropertyType{
			MD_GeometricObjects: &MD_GeometricObjects_Type{Name: "MD_GeometricObjects"},
	}
}

func (MD_GeometricObjects_Type) GongGetAssociationName() any {
	return &MD_GeometricObjects_Type{
			GeometricObjectType: &MD_GeometricObjectTypeCode_PropertyType{Name: "GeometricObjectType"},
	}
}

func (MD_Georectified_PropertyType) GongGetAssociationName() any {
	return &MD_Georectified_PropertyType{
			MD_Georectified: &MD_Georectified_Type{Name: "MD_Georectified"},
	}
}

func (MD_Georectified_Type) GongGetAssociationName() any {
	return &MD_Georectified_Type{
			AxisDimensionProperties: []*MD_Dimension_PropertyType{{Name: "AxisDimensionProperties"}},
			CellGeometry: &MD_CellGeometryCode_PropertyType{Name: "CellGeometry"},
			PointInPixel: &MD_PixelOrientationCode_PropertyType{Name: "PointInPixel"},
	}
}

func (MD_Georeferenceable_PropertyType) GongGetAssociationName() any {
	return &MD_Georeferenceable_PropertyType{
			MD_Georeferenceable: &MD_Georeferenceable_Type{Name: "MD_Georeferenceable"},
	}
}

func (MD_Georeferenceable_Type) GongGetAssociationName() any {
	return &MD_Georeferenceable_Type{
			AxisDimensionProperties: []*MD_Dimension_PropertyType{{Name: "AxisDimensionProperties"}},
			CellGeometry: &MD_CellGeometryCode_PropertyType{Name: "CellGeometry"},
	}
}

func (MD_GridSpatialRepresentation_PropertyType) GongGetAssociationName() any {
	return &MD_GridSpatialRepresentation_PropertyType{
			MD_GridSpatialRepresentation: &MD_GridSpatialRepresentation_Type{Name: "MD_GridSpatialRepresentation"},
			MD_Georeferenceable: &MD_Georeferenceable_Type{Name: "MD_Georeferenceable"},
			MD_Georectified: &MD_Georectified_Type{Name: "MD_Georectified"},
	}
}

func (MD_GridSpatialRepresentation_Type) GongGetAssociationName() any {
	return &MD_GridSpatialRepresentation_Type{
			AxisDimensionProperties: []*MD_Dimension_PropertyType{{Name: "AxisDimensionProperties"}},
			CellGeometry: &MD_CellGeometryCode_PropertyType{Name: "CellGeometry"},
	}
}

func (MD_PixelOrientationCode_PropertyType) GongGetAssociationName() any {
	return &MD_PixelOrientationCode_PropertyType{
	}
}

func (MD_SpatialRepresentation_PropertyType) GongGetAssociationName() any {
	return &MD_SpatialRepresentation_PropertyType{
			MD_GridSpatialRepresentation: &MD_GridSpatialRepresentation_Type{Name: "MD_GridSpatialRepresentation"},
			MD_Georeferenceable: &MD_Georeferenceable_Type{Name: "MD_Georeferenceable"},
			MD_Georectified: &MD_Georectified_Type{Name: "MD_Georectified"},
			MD_VectorSpatialRepresentation: &MD_VectorSpatialRepresentation_Type{Name: "MD_VectorSpatialRepresentation"},
	}
}

func (MD_TopologyLevelCode_PropertyType) GongGetAssociationName() any {
	return &MD_TopologyLevelCode_PropertyType{
	}
}

func (MD_VectorSpatialRepresentation_PropertyType) GongGetAssociationName() any {
	return &MD_VectorSpatialRepresentation_PropertyType{
			MD_VectorSpatialRepresentation: &MD_VectorSpatialRepresentation_Type{Name: "MD_VectorSpatialRepresentation"},
	}
}

func (MD_VectorSpatialRepresentation_Type) GongGetAssociationName() any {
	return &MD_VectorSpatialRepresentation_Type{
			TopologyLevel: &MD_TopologyLevelCode_PropertyType{Name: "TopologyLevel"},
			GeometricObjects: []*MD_GeometricObjects_PropertyType{{Name: "GeometricObjects"}},
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
	// reverse maps of direct associations of AbstractMD_SpatialRepresentation_Type
	case AbstractMD_SpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CellGeometryCode_PropertyType
	case MD_CellGeometryCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_DimensionNameTypeCode_PropertyType
	case MD_DimensionNameTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Dimension_PropertyType
	case MD_Dimension_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_Dimension":
			res := make(map[*MD_Dimension_Type][]*MD_Dimension_PropertyType)
			for md_dimension_propertytype := range stage.MD_Dimension_PropertyTypes {
				if md_dimension_propertytype.MD_Dimension != nil {
					md_dimension_type_ := md_dimension_propertytype.MD_Dimension
					var md_dimension_propertytypes []*MD_Dimension_PropertyType
					_, ok := res[md_dimension_type_]
					if ok {
						md_dimension_propertytypes = res[md_dimension_type_]
					} else {
						md_dimension_propertytypes = make([]*MD_Dimension_PropertyType, 0)
					}
					md_dimension_propertytypes = append(md_dimension_propertytypes, md_dimension_propertytype)
					res[md_dimension_type_] = md_dimension_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Dimension_Type
	case MD_Dimension_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "DimensionName":
			res := make(map[*MD_DimensionNameTypeCode_PropertyType][]*MD_Dimension_Type)
			for md_dimension_type := range stage.MD_Dimension_Types {
				if md_dimension_type.DimensionName != nil {
					md_dimensionnametypecode_propertytype_ := md_dimension_type.DimensionName
					var md_dimension_types []*MD_Dimension_Type
					_, ok := res[md_dimensionnametypecode_propertytype_]
					if ok {
						md_dimension_types = res[md_dimensionnametypecode_propertytype_]
					} else {
						md_dimension_types = make([]*MD_Dimension_Type, 0)
					}
					md_dimension_types = append(md_dimension_types, md_dimension_type)
					res[md_dimensionnametypecode_propertytype_] = md_dimension_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_GeometricObjectTypeCode_PropertyType
	case MD_GeometricObjectTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_GeometricObjects_PropertyType
	case MD_GeometricObjects_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_GeometricObjects":
			res := make(map[*MD_GeometricObjects_Type][]*MD_GeometricObjects_PropertyType)
			for md_geometricobjects_propertytype := range stage.MD_GeometricObjects_PropertyTypes {
				if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
					md_geometricobjects_type_ := md_geometricobjects_propertytype.MD_GeometricObjects
					var md_geometricobjects_propertytypes []*MD_GeometricObjects_PropertyType
					_, ok := res[md_geometricobjects_type_]
					if ok {
						md_geometricobjects_propertytypes = res[md_geometricobjects_type_]
					} else {
						md_geometricobjects_propertytypes = make([]*MD_GeometricObjects_PropertyType, 0)
					}
					md_geometricobjects_propertytypes = append(md_geometricobjects_propertytypes, md_geometricobjects_propertytype)
					res[md_geometricobjects_type_] = md_geometricobjects_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_GeometricObjects_Type
	case MD_GeometricObjects_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "GeometricObjectType":
			res := make(map[*MD_GeometricObjectTypeCode_PropertyType][]*MD_GeometricObjects_Type)
			for md_geometricobjects_type := range stage.MD_GeometricObjects_Types {
				if md_geometricobjects_type.GeometricObjectType != nil {
					md_geometricobjecttypecode_propertytype_ := md_geometricobjects_type.GeometricObjectType
					var md_geometricobjects_types []*MD_GeometricObjects_Type
					_, ok := res[md_geometricobjecttypecode_propertytype_]
					if ok {
						md_geometricobjects_types = res[md_geometricobjecttypecode_propertytype_]
					} else {
						md_geometricobjects_types = make([]*MD_GeometricObjects_Type, 0)
					}
					md_geometricobjects_types = append(md_geometricobjects_types, md_geometricobjects_type)
					res[md_geometricobjecttypecode_propertytype_] = md_geometricobjects_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Georectified_PropertyType
	case MD_Georectified_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_Georectified":
			res := make(map[*MD_Georectified_Type][]*MD_Georectified_PropertyType)
			for md_georectified_propertytype := range stage.MD_Georectified_PropertyTypes {
				if md_georectified_propertytype.MD_Georectified != nil {
					md_georectified_type_ := md_georectified_propertytype.MD_Georectified
					var md_georectified_propertytypes []*MD_Georectified_PropertyType
					_, ok := res[md_georectified_type_]
					if ok {
						md_georectified_propertytypes = res[md_georectified_type_]
					} else {
						md_georectified_propertytypes = make([]*MD_Georectified_PropertyType, 0)
					}
					md_georectified_propertytypes = append(md_georectified_propertytypes, md_georectified_propertytype)
					res[md_georectified_type_] = md_georectified_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Georectified_Type
	case MD_Georectified_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "CellGeometry":
			res := make(map[*MD_CellGeometryCode_PropertyType][]*MD_Georectified_Type)
			for md_georectified_type := range stage.MD_Georectified_Types {
				if md_georectified_type.CellGeometry != nil {
					md_cellgeometrycode_propertytype_ := md_georectified_type.CellGeometry
					var md_georectified_types []*MD_Georectified_Type
					_, ok := res[md_cellgeometrycode_propertytype_]
					if ok {
						md_georectified_types = res[md_cellgeometrycode_propertytype_]
					} else {
						md_georectified_types = make([]*MD_Georectified_Type, 0)
					}
					md_georectified_types = append(md_georectified_types, md_georectified_type)
					res[md_cellgeometrycode_propertytype_] = md_georectified_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointInPixel":
			res := make(map[*MD_PixelOrientationCode_PropertyType][]*MD_Georectified_Type)
			for md_georectified_type := range stage.MD_Georectified_Types {
				if md_georectified_type.PointInPixel != nil {
					md_pixelorientationcode_propertytype_ := md_georectified_type.PointInPixel
					var md_georectified_types []*MD_Georectified_Type
					_, ok := res[md_pixelorientationcode_propertytype_]
					if ok {
						md_georectified_types = res[md_pixelorientationcode_propertytype_]
					} else {
						md_georectified_types = make([]*MD_Georectified_Type, 0)
					}
					md_georectified_types = append(md_georectified_types, md_georectified_type)
					res[md_pixelorientationcode_propertytype_] = md_georectified_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Georeferenceable_PropertyType
	case MD_Georeferenceable_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_Georeferenceable":
			res := make(map[*MD_Georeferenceable_Type][]*MD_Georeferenceable_PropertyType)
			for md_georeferenceable_propertytype := range stage.MD_Georeferenceable_PropertyTypes {
				if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
					md_georeferenceable_type_ := md_georeferenceable_propertytype.MD_Georeferenceable
					var md_georeferenceable_propertytypes []*MD_Georeferenceable_PropertyType
					_, ok := res[md_georeferenceable_type_]
					if ok {
						md_georeferenceable_propertytypes = res[md_georeferenceable_type_]
					} else {
						md_georeferenceable_propertytypes = make([]*MD_Georeferenceable_PropertyType, 0)
					}
					md_georeferenceable_propertytypes = append(md_georeferenceable_propertytypes, md_georeferenceable_propertytype)
					res[md_georeferenceable_type_] = md_georeferenceable_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Georeferenceable_Type
	case MD_Georeferenceable_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "CellGeometry":
			res := make(map[*MD_CellGeometryCode_PropertyType][]*MD_Georeferenceable_Type)
			for md_georeferenceable_type := range stage.MD_Georeferenceable_Types {
				if md_georeferenceable_type.CellGeometry != nil {
					md_cellgeometrycode_propertytype_ := md_georeferenceable_type.CellGeometry
					var md_georeferenceable_types []*MD_Georeferenceable_Type
					_, ok := res[md_cellgeometrycode_propertytype_]
					if ok {
						md_georeferenceable_types = res[md_cellgeometrycode_propertytype_]
					} else {
						md_georeferenceable_types = make([]*MD_Georeferenceable_Type, 0)
					}
					md_georeferenceable_types = append(md_georeferenceable_types, md_georeferenceable_type)
					res[md_cellgeometrycode_propertytype_] = md_georeferenceable_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_GridSpatialRepresentation_PropertyType
	case MD_GridSpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_GridSpatialRepresentation":
			res := make(map[*MD_GridSpatialRepresentation_Type][]*MD_GridSpatialRepresentation_PropertyType)
			for md_gridspatialrepresentation_propertytype := range stage.MD_GridSpatialRepresentation_PropertyTypes {
				if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
					md_gridspatialrepresentation_type_ := md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation
					var md_gridspatialrepresentation_propertytypes []*MD_GridSpatialRepresentation_PropertyType
					_, ok := res[md_gridspatialrepresentation_type_]
					if ok {
						md_gridspatialrepresentation_propertytypes = res[md_gridspatialrepresentation_type_]
					} else {
						md_gridspatialrepresentation_propertytypes = make([]*MD_GridSpatialRepresentation_PropertyType, 0)
					}
					md_gridspatialrepresentation_propertytypes = append(md_gridspatialrepresentation_propertytypes, md_gridspatialrepresentation_propertytype)
					res[md_gridspatialrepresentation_type_] = md_gridspatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_Georeferenceable":
			res := make(map[*MD_Georeferenceable_Type][]*MD_GridSpatialRepresentation_PropertyType)
			for md_gridspatialrepresentation_propertytype := range stage.MD_GridSpatialRepresentation_PropertyTypes {
				if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
					md_georeferenceable_type_ := md_gridspatialrepresentation_propertytype.MD_Georeferenceable
					var md_gridspatialrepresentation_propertytypes []*MD_GridSpatialRepresentation_PropertyType
					_, ok := res[md_georeferenceable_type_]
					if ok {
						md_gridspatialrepresentation_propertytypes = res[md_georeferenceable_type_]
					} else {
						md_gridspatialrepresentation_propertytypes = make([]*MD_GridSpatialRepresentation_PropertyType, 0)
					}
					md_gridspatialrepresentation_propertytypes = append(md_gridspatialrepresentation_propertytypes, md_gridspatialrepresentation_propertytype)
					res[md_georeferenceable_type_] = md_gridspatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_Georectified":
			res := make(map[*MD_Georectified_Type][]*MD_GridSpatialRepresentation_PropertyType)
			for md_gridspatialrepresentation_propertytype := range stage.MD_GridSpatialRepresentation_PropertyTypes {
				if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
					md_georectified_type_ := md_gridspatialrepresentation_propertytype.MD_Georectified
					var md_gridspatialrepresentation_propertytypes []*MD_GridSpatialRepresentation_PropertyType
					_, ok := res[md_georectified_type_]
					if ok {
						md_gridspatialrepresentation_propertytypes = res[md_georectified_type_]
					} else {
						md_gridspatialrepresentation_propertytypes = make([]*MD_GridSpatialRepresentation_PropertyType, 0)
					}
					md_gridspatialrepresentation_propertytypes = append(md_gridspatialrepresentation_propertytypes, md_gridspatialrepresentation_propertytype)
					res[md_georectified_type_] = md_gridspatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_GridSpatialRepresentation_Type
	case MD_GridSpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "CellGeometry":
			res := make(map[*MD_CellGeometryCode_PropertyType][]*MD_GridSpatialRepresentation_Type)
			for md_gridspatialrepresentation_type := range stage.MD_GridSpatialRepresentation_Types {
				if md_gridspatialrepresentation_type.CellGeometry != nil {
					md_cellgeometrycode_propertytype_ := md_gridspatialrepresentation_type.CellGeometry
					var md_gridspatialrepresentation_types []*MD_GridSpatialRepresentation_Type
					_, ok := res[md_cellgeometrycode_propertytype_]
					if ok {
						md_gridspatialrepresentation_types = res[md_cellgeometrycode_propertytype_]
					} else {
						md_gridspatialrepresentation_types = make([]*MD_GridSpatialRepresentation_Type, 0)
					}
					md_gridspatialrepresentation_types = append(md_gridspatialrepresentation_types, md_gridspatialrepresentation_type)
					res[md_cellgeometrycode_propertytype_] = md_gridspatialrepresentation_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_PixelOrientationCode_PropertyType
	case MD_PixelOrientationCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_SpatialRepresentation_PropertyType
	case MD_SpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_GridSpatialRepresentation":
			res := make(map[*MD_GridSpatialRepresentation_Type][]*MD_SpatialRepresentation_PropertyType)
			for md_spatialrepresentation_propertytype := range stage.MD_SpatialRepresentation_PropertyTypes {
				if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
					md_gridspatialrepresentation_type_ := md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation
					var md_spatialrepresentation_propertytypes []*MD_SpatialRepresentation_PropertyType
					_, ok := res[md_gridspatialrepresentation_type_]
					if ok {
						md_spatialrepresentation_propertytypes = res[md_gridspatialrepresentation_type_]
					} else {
						md_spatialrepresentation_propertytypes = make([]*MD_SpatialRepresentation_PropertyType, 0)
					}
					md_spatialrepresentation_propertytypes = append(md_spatialrepresentation_propertytypes, md_spatialrepresentation_propertytype)
					res[md_gridspatialrepresentation_type_] = md_spatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_Georeferenceable":
			res := make(map[*MD_Georeferenceable_Type][]*MD_SpatialRepresentation_PropertyType)
			for md_spatialrepresentation_propertytype := range stage.MD_SpatialRepresentation_PropertyTypes {
				if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
					md_georeferenceable_type_ := md_spatialrepresentation_propertytype.MD_Georeferenceable
					var md_spatialrepresentation_propertytypes []*MD_SpatialRepresentation_PropertyType
					_, ok := res[md_georeferenceable_type_]
					if ok {
						md_spatialrepresentation_propertytypes = res[md_georeferenceable_type_]
					} else {
						md_spatialrepresentation_propertytypes = make([]*MD_SpatialRepresentation_PropertyType, 0)
					}
					md_spatialrepresentation_propertytypes = append(md_spatialrepresentation_propertytypes, md_spatialrepresentation_propertytype)
					res[md_georeferenceable_type_] = md_spatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_Georectified":
			res := make(map[*MD_Georectified_Type][]*MD_SpatialRepresentation_PropertyType)
			for md_spatialrepresentation_propertytype := range stage.MD_SpatialRepresentation_PropertyTypes {
				if md_spatialrepresentation_propertytype.MD_Georectified != nil {
					md_georectified_type_ := md_spatialrepresentation_propertytype.MD_Georectified
					var md_spatialrepresentation_propertytypes []*MD_SpatialRepresentation_PropertyType
					_, ok := res[md_georectified_type_]
					if ok {
						md_spatialrepresentation_propertytypes = res[md_georectified_type_]
					} else {
						md_spatialrepresentation_propertytypes = make([]*MD_SpatialRepresentation_PropertyType, 0)
					}
					md_spatialrepresentation_propertytypes = append(md_spatialrepresentation_propertytypes, md_spatialrepresentation_propertytype)
					res[md_georectified_type_] = md_spatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MD_VectorSpatialRepresentation":
			res := make(map[*MD_VectorSpatialRepresentation_Type][]*MD_SpatialRepresentation_PropertyType)
			for md_spatialrepresentation_propertytype := range stage.MD_SpatialRepresentation_PropertyTypes {
				if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
					md_vectorspatialrepresentation_type_ := md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation
					var md_spatialrepresentation_propertytypes []*MD_SpatialRepresentation_PropertyType
					_, ok := res[md_vectorspatialrepresentation_type_]
					if ok {
						md_spatialrepresentation_propertytypes = res[md_vectorspatialrepresentation_type_]
					} else {
						md_spatialrepresentation_propertytypes = make([]*MD_SpatialRepresentation_PropertyType, 0)
					}
					md_spatialrepresentation_propertytypes = append(md_spatialrepresentation_propertytypes, md_spatialrepresentation_propertytype)
					res[md_vectorspatialrepresentation_type_] = md_spatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_TopologyLevelCode_PropertyType
	case MD_TopologyLevelCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_VectorSpatialRepresentation_PropertyType
	case MD_VectorSpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_VectorSpatialRepresentation":
			res := make(map[*MD_VectorSpatialRepresentation_Type][]*MD_VectorSpatialRepresentation_PropertyType)
			for md_vectorspatialrepresentation_propertytype := range stage.MD_VectorSpatialRepresentation_PropertyTypes {
				if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
					md_vectorspatialrepresentation_type_ := md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation
					var md_vectorspatialrepresentation_propertytypes []*MD_VectorSpatialRepresentation_PropertyType
					_, ok := res[md_vectorspatialrepresentation_type_]
					if ok {
						md_vectorspatialrepresentation_propertytypes = res[md_vectorspatialrepresentation_type_]
					} else {
						md_vectorspatialrepresentation_propertytypes = make([]*MD_VectorSpatialRepresentation_PropertyType, 0)
					}
					md_vectorspatialrepresentation_propertytypes = append(md_vectorspatialrepresentation_propertytypes, md_vectorspatialrepresentation_propertytype)
					res[md_vectorspatialrepresentation_type_] = md_vectorspatialrepresentation_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_VectorSpatialRepresentation_Type
	case MD_VectorSpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "TopologyLevel":
			res := make(map[*MD_TopologyLevelCode_PropertyType][]*MD_VectorSpatialRepresentation_Type)
			for md_vectorspatialrepresentation_type := range stage.MD_VectorSpatialRepresentation_Types {
				if md_vectorspatialrepresentation_type.TopologyLevel != nil {
					md_topologylevelcode_propertytype_ := md_vectorspatialrepresentation_type.TopologyLevel
					var md_vectorspatialrepresentation_types []*MD_VectorSpatialRepresentation_Type
					_, ok := res[md_topologylevelcode_propertytype_]
					if ok {
						md_vectorspatialrepresentation_types = res[md_topologylevelcode_propertytype_]
					} else {
						md_vectorspatialrepresentation_types = make([]*MD_VectorSpatialRepresentation_Type, 0)
					}
					md_vectorspatialrepresentation_types = append(md_vectorspatialrepresentation_types, md_vectorspatialrepresentation_type)
					res[md_topologylevelcode_propertytype_] = md_vectorspatialrepresentation_types
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
	// reverse maps of direct associations of AbstractMD_SpatialRepresentation_Type
	case AbstractMD_SpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_CellGeometryCode_PropertyType
	case MD_CellGeometryCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_DimensionNameTypeCode_PropertyType
	case MD_DimensionNameTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Dimension_PropertyType
	case MD_Dimension_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Dimension_Type
	case MD_Dimension_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_GeometricObjectTypeCode_PropertyType
	case MD_GeometricObjectTypeCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_GeometricObjects_PropertyType
	case MD_GeometricObjects_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_GeometricObjects_Type
	case MD_GeometricObjects_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Georectified_PropertyType
	case MD_Georectified_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Georectified_Type
	case MD_Georectified_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "AxisDimensionProperties":
			res := make(map[*MD_Dimension_PropertyType][]*MD_Georectified_Type)
			for md_georectified_type := range stage.MD_Georectified_Types {
				for _, md_dimension_propertytype_ := range md_georectified_type.AxisDimensionProperties {
					res[md_dimension_propertytype_] = append(res[md_dimension_propertytype_], md_georectified_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Georeferenceable_PropertyType
	case MD_Georeferenceable_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Georeferenceable_Type
	case MD_Georeferenceable_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "AxisDimensionProperties":
			res := make(map[*MD_Dimension_PropertyType][]*MD_Georeferenceable_Type)
			for md_georeferenceable_type := range stage.MD_Georeferenceable_Types {
				for _, md_dimension_propertytype_ := range md_georeferenceable_type.AxisDimensionProperties {
					res[md_dimension_propertytype_] = append(res[md_dimension_propertytype_], md_georeferenceable_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_GridSpatialRepresentation_PropertyType
	case MD_GridSpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_GridSpatialRepresentation_Type
	case MD_GridSpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "AxisDimensionProperties":
			res := make(map[*MD_Dimension_PropertyType][]*MD_GridSpatialRepresentation_Type)
			for md_gridspatialrepresentation_type := range stage.MD_GridSpatialRepresentation_Types {
				for _, md_dimension_propertytype_ := range md_gridspatialrepresentation_type.AxisDimensionProperties {
					res[md_dimension_propertytype_] = append(res[md_dimension_propertytype_], md_gridspatialrepresentation_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_PixelOrientationCode_PropertyType
	case MD_PixelOrientationCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_SpatialRepresentation_PropertyType
	case MD_SpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_TopologyLevelCode_PropertyType
	case MD_TopologyLevelCode_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_VectorSpatialRepresentation_PropertyType
	case MD_VectorSpatialRepresentation_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_VectorSpatialRepresentation_Type
	case MD_VectorSpatialRepresentation_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "GeometricObjects":
			res := make(map[*MD_GeometricObjects_PropertyType][]*MD_VectorSpatialRepresentation_Type)
			for md_vectorspatialrepresentation_type := range stage.MD_VectorSpatialRepresentation_Types {
				for _, md_geometricobjects_propertytype_ := range md_vectorspatialrepresentation_type.GeometricObjects {
					res[md_geometricobjects_propertytype_] = append(res[md_geometricobjects_propertytype_], md_vectorspatialrepresentation_type)
				}
			}
			return any(res).(map[*End][]*Start)
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
func (*AbstractMD_SpatialRepresentation_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_CellGeometryCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_DimensionNameTypeCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Dimension_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MD_Georectified_Type",
			Fieldname: "AxisDimensionProperties",
		},
		{
			GongstructName: "MD_Georeferenceable_Type",
			Fieldname: "AxisDimensionProperties",
		},
		{
			GongstructName: "MD_GridSpatialRepresentation_Type",
			Fieldname: "AxisDimensionProperties",
		},
	}
}

func (*MD_Dimension_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_GeometricObjectTypeCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_GeometricObjects_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MD_VectorSpatialRepresentation_Type",
			Fieldname: "GeometricObjects",
		},
	}
}

func (*MD_GeometricObjects_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Georectified_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Georectified_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Georeferenceable_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Georeferenceable_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_GridSpatialRepresentation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_GridSpatialRepresentation_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_PixelOrientationCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_SpatialRepresentation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_TopologyLevelCode_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_VectorSpatialRepresentation_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_VectorSpatialRepresentation_Type) GongGetReverseFields() []GongReverseField {
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
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_Dimension",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Dimension_Type",
		},
	}
	return
}

func (md_dimension_type *MD_Dimension_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DimensionName",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_DimensionNameTypeCode_PropertyType",
		},
	}
	return
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_GeometricObjects",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_GeometricObjects_Type",
		},
	}
	return
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "GeometricObjectType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_GeometricObjectTypeCode_PropertyType",
		},
	}
	return
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_Georectified",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georectified_Type",
		},
	}
	return
}

func (md_georectified_type *MD_Georectified_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AxisDimensionProperties",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Dimension_PropertyType",
		},
		{
			Name:                 "CellGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CellGeometryCode_PropertyType",
		},
		{
			Name:                 "PointInPixel",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_PixelOrientationCode_PropertyType",
		},
	}
	return
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_Georeferenceable",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georeferenceable_Type",
		},
	}
	return
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AxisDimensionProperties",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Dimension_PropertyType",
		},
		{
			Name:                 "CellGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CellGeometryCode_PropertyType",
		},
	}
	return
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_GridSpatialRepresentation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_GridSpatialRepresentation_Type",
		},
		{
			Name:                 "MD_Georeferenceable",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georeferenceable_Type",
		},
		{
			Name:                 "MD_Georectified",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georectified_Type",
		},
	}
	return
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AxisDimensionProperties",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Dimension_PropertyType",
		},
		{
			Name:                 "CellGeometry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_CellGeometryCode_PropertyType",
		},
	}
	return
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_PixelOrientationCode",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Enum_MD_PixelOrientationCode_Type",
		},
	}
	return
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_GridSpatialRepresentation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_GridSpatialRepresentation_Type",
		},
		{
			Name:                 "MD_Georeferenceable",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georeferenceable_Type",
		},
		{
			Name:                 "MD_Georectified",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Georectified_Type",
		},
		{
			Name:                 "MD_VectorSpatialRepresentation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_VectorSpatialRepresentation_Type",
		},
	}
	return
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_VectorSpatialRepresentation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_VectorSpatialRepresentation_Type",
		},
	}
	return
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TopologyLevel",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_TopologyLevelCode_PropertyType",
		},
		{
			Name:                 "GeometricObjects",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_GeometricObjects_PropertyType",
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
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractmd_spatialrepresentation_type.Name
	case "Id":
		res.valueString = abstractmd_spatialrepresentation_type.Id
	case "Uuid":
		res.valueString = abstractmd_spatialrepresentation_type.Uuid
	}
	return
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_cellgeometrycode_propertytype.Name
	case "NilReason":
		res.valueString = md_cellgeometrycode_propertytype.NilReason
	}
	return
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_dimensionnametypecode_propertytype.Name
	case "NilReason":
		res.valueString = md_dimensionnametypecode_propertytype.NilReason
	}
	return
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_dimension_propertytype.Name
	case "NilReason":
		res.valueString = md_dimension_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_dimension_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_dimension_propertytype.Type_XSD
	case "Href":
		res.valueString = md_dimension_propertytype.Href
	case "Role":
		res.valueString = md_dimension_propertytype.Role
	case "Arcrole":
		res.valueString = md_dimension_propertytype.Arcrole
	case "Title":
		res.valueString = md_dimension_propertytype.Title
	case "Show":
		res.valueString = md_dimension_propertytype.Show
	case "Actuate":
		res.valueString = md_dimension_propertytype.Actuate
	case "MD_Dimension":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_dimension_propertytype.MD_Dimension != nil {
			res.valueString = md_dimension_propertytype.MD_Dimension.Name
			res.ids = md_dimension_propertytype.MD_Dimension.GongGetUUID(stage)
		}
	}
	return
}

func (md_dimension_type *MD_Dimension_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_dimension_type.Name
	case "Id":
		res.valueString = md_dimension_type.Id
	case "Uuid":
		res.valueString = md_dimension_type.Uuid
	case "DimensionName":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_dimension_type.DimensionName != nil {
			res.valueString = md_dimension_type.DimensionName.Name
			res.ids = md_dimension_type.DimensionName.GongGetUUID(stage)
		}
	}
	return
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_geometricobjecttypecode_propertytype.Name
	case "NilReason":
		res.valueString = md_geometricobjecttypecode_propertytype.NilReason
	}
	return
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_geometricobjects_propertytype.Name
	case "NilReason":
		res.valueString = md_geometricobjects_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_geometricobjects_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_geometricobjects_propertytype.Type_XSD
	case "Href":
		res.valueString = md_geometricobjects_propertytype.Href
	case "Role":
		res.valueString = md_geometricobjects_propertytype.Role
	case "Arcrole":
		res.valueString = md_geometricobjects_propertytype.Arcrole
	case "Title":
		res.valueString = md_geometricobjects_propertytype.Title
	case "Show":
		res.valueString = md_geometricobjects_propertytype.Show
	case "Actuate":
		res.valueString = md_geometricobjects_propertytype.Actuate
	case "MD_GeometricObjects":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
			res.valueString = md_geometricobjects_propertytype.MD_GeometricObjects.Name
			res.ids = md_geometricobjects_propertytype.MD_GeometricObjects.GongGetUUID(stage)
		}
	}
	return
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_geometricobjects_type.Name
	case "Id":
		res.valueString = md_geometricobjects_type.Id
	case "Uuid":
		res.valueString = md_geometricobjects_type.Uuid
	case "GeometricObjectType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_geometricobjects_type.GeometricObjectType != nil {
			res.valueString = md_geometricobjects_type.GeometricObjectType.Name
			res.ids = md_geometricobjects_type.GeometricObjectType.GongGetUUID(stage)
		}
	}
	return
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_georectified_propertytype.Name
	case "NilReason":
		res.valueString = md_georectified_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_georectified_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_georectified_propertytype.Type_XSD
	case "Href":
		res.valueString = md_georectified_propertytype.Href
	case "Role":
		res.valueString = md_georectified_propertytype.Role
	case "Arcrole":
		res.valueString = md_georectified_propertytype.Arcrole
	case "Title":
		res.valueString = md_georectified_propertytype.Title
	case "Show":
		res.valueString = md_georectified_propertytype.Show
	case "Actuate":
		res.valueString = md_georectified_propertytype.Actuate
	case "MD_Georectified":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_georectified_propertytype.MD_Georectified != nil {
			res.valueString = md_georectified_propertytype.MD_Georectified.Name
			res.ids = md_georectified_propertytype.MD_Georectified.GongGetUUID(stage)
		}
	}
	return
}

func (md_georectified_type *MD_Georectified_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_georectified_type.Name
	case "Id":
		res.valueString = md_georectified_type.Id
	case "Uuid":
		res.valueString = md_georectified_type.Uuid
	case "AxisDimensionProperties":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_georectified_type.AxisDimensionProperties {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CellGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_georectified_type.CellGeometry != nil {
			res.valueString = md_georectified_type.CellGeometry.Name
			res.ids = md_georectified_type.CellGeometry.GongGetUUID(stage)
		}
	case "PointInPixel":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_georectified_type.PointInPixel != nil {
			res.valueString = md_georectified_type.PointInPixel.Name
			res.ids = md_georectified_type.PointInPixel.GongGetUUID(stage)
		}
	}
	return
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_georeferenceable_propertytype.Name
	case "NilReason":
		res.valueString = md_georeferenceable_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_georeferenceable_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_georeferenceable_propertytype.Type_XSD
	case "Href":
		res.valueString = md_georeferenceable_propertytype.Href
	case "Role":
		res.valueString = md_georeferenceable_propertytype.Role
	case "Arcrole":
		res.valueString = md_georeferenceable_propertytype.Arcrole
	case "Title":
		res.valueString = md_georeferenceable_propertytype.Title
	case "Show":
		res.valueString = md_georeferenceable_propertytype.Show
	case "Actuate":
		res.valueString = md_georeferenceable_propertytype.Actuate
	case "MD_Georeferenceable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
			res.valueString = md_georeferenceable_propertytype.MD_Georeferenceable.Name
			res.ids = md_georeferenceable_propertytype.MD_Georeferenceable.GongGetUUID(stage)
		}
	}
	return
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_georeferenceable_type.Name
	case "Id":
		res.valueString = md_georeferenceable_type.Id
	case "Uuid":
		res.valueString = md_georeferenceable_type.Uuid
	case "AxisDimensionProperties":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_georeferenceable_type.AxisDimensionProperties {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CellGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_georeferenceable_type.CellGeometry != nil {
			res.valueString = md_georeferenceable_type.CellGeometry.Name
			res.ids = md_georeferenceable_type.CellGeometry.GongGetUUID(stage)
		}
	}
	return
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_gridspatialrepresentation_propertytype.Name
	case "NilReason":
		res.valueString = md_gridspatialrepresentation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_gridspatialrepresentation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_gridspatialrepresentation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_gridspatialrepresentation_propertytype.Href
	case "Role":
		res.valueString = md_gridspatialrepresentation_propertytype.Role
	case "Arcrole":
		res.valueString = md_gridspatialrepresentation_propertytype.Arcrole
	case "Title":
		res.valueString = md_gridspatialrepresentation_propertytype.Title
	case "Show":
		res.valueString = md_gridspatialrepresentation_propertytype.Show
	case "Actuate":
		res.valueString = md_gridspatialrepresentation_propertytype.Actuate
	case "MD_GridSpatialRepresentation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
			res.valueString = md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation.Name
			res.ids = md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetUUID(stage)
		}
	case "MD_Georeferenceable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
			res.valueString = md_gridspatialrepresentation_propertytype.MD_Georeferenceable.Name
			res.ids = md_gridspatialrepresentation_propertytype.MD_Georeferenceable.GongGetUUID(stage)
		}
	case "MD_Georectified":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
			res.valueString = md_gridspatialrepresentation_propertytype.MD_Georectified.Name
			res.ids = md_gridspatialrepresentation_propertytype.MD_Georectified.GongGetUUID(stage)
		}
	}
	return
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_gridspatialrepresentation_type.Name
	case "Id":
		res.valueString = md_gridspatialrepresentation_type.Id
	case "Uuid":
		res.valueString = md_gridspatialrepresentation_type.Uuid
	case "AxisDimensionProperties":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_gridspatialrepresentation_type.AxisDimensionProperties {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CellGeometry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_gridspatialrepresentation_type.CellGeometry != nil {
			res.valueString = md_gridspatialrepresentation_type.CellGeometry.Name
			res.ids = md_gridspatialrepresentation_type.CellGeometry.GongGetUUID(stage)
		}
	}
	return
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_pixelorientationcode_propertytype.Name
	case "NilReason":
		res.valueString = md_pixelorientationcode_propertytype.NilReason
	case "MD_PixelOrientationCode":
		enum := md_pixelorientationcode_propertytype.MD_PixelOrientationCode
		res.valueString = enum.ToCodeString()
	}
	return
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_spatialrepresentation_propertytype.Name
	case "NilReason":
		res.valueString = md_spatialrepresentation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_spatialrepresentation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_spatialrepresentation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_spatialrepresentation_propertytype.Href
	case "Role":
		res.valueString = md_spatialrepresentation_propertytype.Role
	case "Arcrole":
		res.valueString = md_spatialrepresentation_propertytype.Arcrole
	case "Title":
		res.valueString = md_spatialrepresentation_propertytype.Title
	case "Show":
		res.valueString = md_spatialrepresentation_propertytype.Show
	case "Actuate":
		res.valueString = md_spatialrepresentation_propertytype.Actuate
	case "MD_GridSpatialRepresentation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
			res.valueString = md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation.Name
			res.ids = md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetUUID(stage)
		}
	case "MD_Georeferenceable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
			res.valueString = md_spatialrepresentation_propertytype.MD_Georeferenceable.Name
			res.ids = md_spatialrepresentation_propertytype.MD_Georeferenceable.GongGetUUID(stage)
		}
	case "MD_Georectified":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_spatialrepresentation_propertytype.MD_Georectified != nil {
			res.valueString = md_spatialrepresentation_propertytype.MD_Georectified.Name
			res.ids = md_spatialrepresentation_propertytype.MD_Georectified.GongGetUUID(stage)
		}
	case "MD_VectorSpatialRepresentation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
			res.valueString = md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation.Name
			res.ids = md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetUUID(stage)
		}
	}
	return
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_topologylevelcode_propertytype.Name
	case "NilReason":
		res.valueString = md_topologylevelcode_propertytype.NilReason
	}
	return
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_vectorspatialrepresentation_propertytype.Name
	case "NilReason":
		res.valueString = md_vectorspatialrepresentation_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_vectorspatialrepresentation_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_vectorspatialrepresentation_propertytype.Type_XSD
	case "Href":
		res.valueString = md_vectorspatialrepresentation_propertytype.Href
	case "Role":
		res.valueString = md_vectorspatialrepresentation_propertytype.Role
	case "Arcrole":
		res.valueString = md_vectorspatialrepresentation_propertytype.Arcrole
	case "Title":
		res.valueString = md_vectorspatialrepresentation_propertytype.Title
	case "Show":
		res.valueString = md_vectorspatialrepresentation_propertytype.Show
	case "Actuate":
		res.valueString = md_vectorspatialrepresentation_propertytype.Actuate
	case "MD_VectorSpatialRepresentation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
			res.valueString = md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation.Name
			res.ids = md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetUUID(stage)
		}
	}
	return
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_vectorspatialrepresentation_type.Name
	case "Id":
		res.valueString = md_vectorspatialrepresentation_type.Id
	case "Uuid":
		res.valueString = md_vectorspatialrepresentation_type.Uuid
	case "TopologyLevel":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_vectorspatialrepresentation_type.TopologyLevel != nil {
			res.valueString = md_vectorspatialrepresentation_type.TopologyLevel.Name
			res.ids = md_vectorspatialrepresentation_type.TopologyLevel.GongGetUUID(stage)
		}
	case "GeometricObjects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_vectorspatialrepresentation_type.GeometricObjects {
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetGongstructName() string {
	return "AbstractMD_SpatialRepresentation_Type"
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetGongstructName() string {
	return "MD_CellGeometryCode_PropertyType"
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetGongstructName() string {
	return "MD_DimensionNameTypeCode_PropertyType"
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetGongstructName() string {
	return "MD_Dimension_PropertyType"
}

func (md_dimension_type *MD_Dimension_Type) GongGetGongstructName() string {
	return "MD_Dimension_Type"
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetGongstructName() string {
	return "MD_GeometricObjectTypeCode_PropertyType"
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetGongstructName() string {
	return "MD_GeometricObjects_PropertyType"
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetGongstructName() string {
	return "MD_GeometricObjects_Type"
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetGongstructName() string {
	return "MD_Georectified_PropertyType"
}

func (md_georectified_type *MD_Georectified_Type) GongGetGongstructName() string {
	return "MD_Georectified_Type"
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetGongstructName() string {
	return "MD_Georeferenceable_PropertyType"
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetGongstructName() string {
	return "MD_Georeferenceable_Type"
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetGongstructName() string {
	return "MD_GridSpatialRepresentation_PropertyType"
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetGongstructName() string {
	return "MD_GridSpatialRepresentation_Type"
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetGongstructName() string {
	return "MD_PixelOrientationCode_PropertyType"
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetGongstructName() string {
	return "MD_SpatialRepresentation_PropertyType"
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetGongstructName() string {
	return "MD_TopologyLevelCode_PropertyType"
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetGongstructName() string {
	return "MD_VectorSpatialRepresentation_PropertyType"
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetGongstructName() string {
	return "MD_VectorSpatialRepresentation_Type"
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
	__gong__rebuildMapString(stage.AbstractMD_SpatialRepresentation_Types, &stage.AbstractMD_SpatialRepresentation_Types_mapString)

	__gong__rebuildMapString(stage.MD_CellGeometryCode_PropertyTypes, &stage.MD_CellGeometryCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_DimensionNameTypeCode_PropertyTypes, &stage.MD_DimensionNameTypeCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Dimension_PropertyTypes, &stage.MD_Dimension_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Dimension_Types, &stage.MD_Dimension_Types_mapString)

	__gong__rebuildMapString(stage.MD_GeometricObjectTypeCode_PropertyTypes, &stage.MD_GeometricObjectTypeCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_GeometricObjects_PropertyTypes, &stage.MD_GeometricObjects_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_GeometricObjects_Types, &stage.MD_GeometricObjects_Types_mapString)

	__gong__rebuildMapString(stage.MD_Georectified_PropertyTypes, &stage.MD_Georectified_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Georectified_Types, &stage.MD_Georectified_Types_mapString)

	__gong__rebuildMapString(stage.MD_Georeferenceable_PropertyTypes, &stage.MD_Georeferenceable_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Georeferenceable_Types, &stage.MD_Georeferenceable_Types_mapString)

	__gong__rebuildMapString(stage.MD_GridSpatialRepresentation_PropertyTypes, &stage.MD_GridSpatialRepresentation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_GridSpatialRepresentation_Types, &stage.MD_GridSpatialRepresentation_Types_mapString)

	__gong__rebuildMapString(stage.MD_PixelOrientationCode_PropertyTypes, &stage.MD_PixelOrientationCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_SpatialRepresentation_PropertyTypes, &stage.MD_SpatialRepresentation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_TopologyLevelCode_PropertyTypes, &stage.MD_TopologyLevelCode_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_VectorSpatialRepresentation_PropertyTypes, &stage.MD_VectorSpatialRepresentation_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_VectorSpatialRepresentation_Types, &stage.MD_VectorSpatialRepresentation_Types_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
