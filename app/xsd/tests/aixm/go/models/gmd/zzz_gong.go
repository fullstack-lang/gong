// generated code - do not edit
package gmd

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
	AbstractDS_Aggregate_Types                map[*AbstractDS_Aggregate_Type]struct{}
	AbstractDS_Aggregate_Types_instance       map[*AbstractDS_Aggregate_Type]*AbstractDS_Aggregate_Type
	AbstractDS_Aggregate_Types_mapString      map[string]*AbstractDS_Aggregate_Type
	AbstractDS_Aggregate_TypeOrder            uint
	AbstractDS_Aggregate_Type_stagedOrder     map[*AbstractDS_Aggregate_Type]uint
	AbstractDS_Aggregate_Type_orderStaged     map[uint]*AbstractDS_Aggregate_Type
	AbstractDS_Aggregate_Types_reference      map[*AbstractDS_Aggregate_Type]*AbstractDS_Aggregate_Type
	AbstractDS_Aggregate_Types_referenceOrder map[*AbstractDS_Aggregate_Type]uint

	// insertion point for slice of pointers maps
	AbstractDS_Aggregate_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*AbstractDS_Aggregate_Type

	AbstractDS_Aggregate_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*AbstractDS_Aggregate_Type

	AbstractDS_Aggregate_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*AbstractDS_Aggregate_Type

	AbstractDS_Aggregate_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*AbstractDS_Aggregate_Type

	OnAfterAbstractDS_Aggregate_TypeCreateCallback GongOnAfterCreateInterface[AbstractDS_Aggregate_Type]
	OnAfterAbstractDS_Aggregate_TypeUpdateCallback GongOnAfterUpdateInterface[AbstractDS_Aggregate_Type]
	OnAfterAbstractDS_Aggregate_TypeDeleteCallback GongOnAfterDeleteInterface[AbstractDS_Aggregate_Type]

	DS_Aggregate_PropertyTypes                map[*DS_Aggregate_PropertyType]struct{}
	DS_Aggregate_PropertyTypes_instance       map[*DS_Aggregate_PropertyType]*DS_Aggregate_PropertyType
	DS_Aggregate_PropertyTypes_mapString      map[string]*DS_Aggregate_PropertyType
	DS_Aggregate_PropertyTypeOrder            uint
	DS_Aggregate_PropertyType_stagedOrder     map[*DS_Aggregate_PropertyType]uint
	DS_Aggregate_PropertyType_orderStaged     map[uint]*DS_Aggregate_PropertyType
	DS_Aggregate_PropertyTypes_reference      map[*DS_Aggregate_PropertyType]*DS_Aggregate_PropertyType
	DS_Aggregate_PropertyTypes_referenceOrder map[*DS_Aggregate_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_Aggregate_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_Aggregate_PropertyType]
	OnAfterDS_Aggregate_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_Aggregate_PropertyType]
	OnAfterDS_Aggregate_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_Aggregate_PropertyType]

	DS_DataSet_PropertyTypes                map[*DS_DataSet_PropertyType]struct{}
	DS_DataSet_PropertyTypes_instance       map[*DS_DataSet_PropertyType]*DS_DataSet_PropertyType
	DS_DataSet_PropertyTypes_mapString      map[string]*DS_DataSet_PropertyType
	DS_DataSet_PropertyTypeOrder            uint
	DS_DataSet_PropertyType_stagedOrder     map[*DS_DataSet_PropertyType]uint
	DS_DataSet_PropertyType_orderStaged     map[uint]*DS_DataSet_PropertyType
	DS_DataSet_PropertyTypes_reference      map[*DS_DataSet_PropertyType]*DS_DataSet_PropertyType
	DS_DataSet_PropertyTypes_referenceOrder map[*DS_DataSet_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_DataSet_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_DataSet_PropertyType]
	OnAfterDS_DataSet_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_DataSet_PropertyType]
	OnAfterDS_DataSet_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_DataSet_PropertyType]

	DS_DataSet_Types                map[*DS_DataSet_Type]struct{}
	DS_DataSet_Types_instance       map[*DS_DataSet_Type]*DS_DataSet_Type
	DS_DataSet_Types_mapString      map[string]*DS_DataSet_Type
	DS_DataSet_TypeOrder            uint
	DS_DataSet_Type_stagedOrder     map[*DS_DataSet_Type]uint
	DS_DataSet_Type_orderStaged     map[uint]*DS_DataSet_Type
	DS_DataSet_Types_reference      map[*DS_DataSet_Type]*DS_DataSet_Type
	DS_DataSet_Types_referenceOrder map[*DS_DataSet_Type]uint

	// insertion point for slice of pointers maps
	DS_DataSet_Type_Has_reverseMap map[*MD_Metadata_PropertyType]*DS_DataSet_Type

	DS_DataSet_Type_PartOf_reverseMap map[*DS_Aggregate_PropertyType]*DS_DataSet_Type

	OnAfterDS_DataSet_TypeCreateCallback GongOnAfterCreateInterface[DS_DataSet_Type]
	OnAfterDS_DataSet_TypeUpdateCallback GongOnAfterUpdateInterface[DS_DataSet_Type]
	OnAfterDS_DataSet_TypeDeleteCallback GongOnAfterDeleteInterface[DS_DataSet_Type]

	DS_Initiative_PropertyTypes                map[*DS_Initiative_PropertyType]struct{}
	DS_Initiative_PropertyTypes_instance       map[*DS_Initiative_PropertyType]*DS_Initiative_PropertyType
	DS_Initiative_PropertyTypes_mapString      map[string]*DS_Initiative_PropertyType
	DS_Initiative_PropertyTypeOrder            uint
	DS_Initiative_PropertyType_stagedOrder     map[*DS_Initiative_PropertyType]uint
	DS_Initiative_PropertyType_orderStaged     map[uint]*DS_Initiative_PropertyType
	DS_Initiative_PropertyTypes_reference      map[*DS_Initiative_PropertyType]*DS_Initiative_PropertyType
	DS_Initiative_PropertyTypes_referenceOrder map[*DS_Initiative_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_Initiative_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_Initiative_PropertyType]
	OnAfterDS_Initiative_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_Initiative_PropertyType]
	OnAfterDS_Initiative_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_Initiative_PropertyType]

	DS_Initiative_Types                map[*DS_Initiative_Type]struct{}
	DS_Initiative_Types_instance       map[*DS_Initiative_Type]*DS_Initiative_Type
	DS_Initiative_Types_mapString      map[string]*DS_Initiative_Type
	DS_Initiative_TypeOrder            uint
	DS_Initiative_Type_stagedOrder     map[*DS_Initiative_Type]uint
	DS_Initiative_Type_orderStaged     map[uint]*DS_Initiative_Type
	DS_Initiative_Types_reference      map[*DS_Initiative_Type]*DS_Initiative_Type
	DS_Initiative_Types_referenceOrder map[*DS_Initiative_Type]uint

	// insertion point for slice of pointers maps
	DS_Initiative_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_Initiative_Type

	DS_Initiative_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_Initiative_Type

	DS_Initiative_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Initiative_Type

	DS_Initiative_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Initiative_Type

	OnAfterDS_Initiative_TypeCreateCallback GongOnAfterCreateInterface[DS_Initiative_Type]
	OnAfterDS_Initiative_TypeUpdateCallback GongOnAfterUpdateInterface[DS_Initiative_Type]
	OnAfterDS_Initiative_TypeDeleteCallback GongOnAfterDeleteInterface[DS_Initiative_Type]

	DS_OtherAggregate_PropertyTypes                map[*DS_OtherAggregate_PropertyType]struct{}
	DS_OtherAggregate_PropertyTypes_instance       map[*DS_OtherAggregate_PropertyType]*DS_OtherAggregate_PropertyType
	DS_OtherAggregate_PropertyTypes_mapString      map[string]*DS_OtherAggregate_PropertyType
	DS_OtherAggregate_PropertyTypeOrder            uint
	DS_OtherAggregate_PropertyType_stagedOrder     map[*DS_OtherAggregate_PropertyType]uint
	DS_OtherAggregate_PropertyType_orderStaged     map[uint]*DS_OtherAggregate_PropertyType
	DS_OtherAggregate_PropertyTypes_reference      map[*DS_OtherAggregate_PropertyType]*DS_OtherAggregate_PropertyType
	DS_OtherAggregate_PropertyTypes_referenceOrder map[*DS_OtherAggregate_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_OtherAggregate_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_OtherAggregate_PropertyType]
	OnAfterDS_OtherAggregate_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_OtherAggregate_PropertyType]
	OnAfterDS_OtherAggregate_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_OtherAggregate_PropertyType]

	DS_OtherAggregate_Types                map[*DS_OtherAggregate_Type]struct{}
	DS_OtherAggregate_Types_instance       map[*DS_OtherAggregate_Type]*DS_OtherAggregate_Type
	DS_OtherAggregate_Types_mapString      map[string]*DS_OtherAggregate_Type
	DS_OtherAggregate_TypeOrder            uint
	DS_OtherAggregate_Type_stagedOrder     map[*DS_OtherAggregate_Type]uint
	DS_OtherAggregate_Type_orderStaged     map[uint]*DS_OtherAggregate_Type
	DS_OtherAggregate_Types_reference      map[*DS_OtherAggregate_Type]*DS_OtherAggregate_Type
	DS_OtherAggregate_Types_referenceOrder map[*DS_OtherAggregate_Type]uint

	// insertion point for slice of pointers maps
	DS_OtherAggregate_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_OtherAggregate_Type

	DS_OtherAggregate_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_OtherAggregate_Type

	DS_OtherAggregate_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_OtherAggregate_Type

	DS_OtherAggregate_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_OtherAggregate_Type

	OnAfterDS_OtherAggregate_TypeCreateCallback GongOnAfterCreateInterface[DS_OtherAggregate_Type]
	OnAfterDS_OtherAggregate_TypeUpdateCallback GongOnAfterUpdateInterface[DS_OtherAggregate_Type]
	OnAfterDS_OtherAggregate_TypeDeleteCallback GongOnAfterDeleteInterface[DS_OtherAggregate_Type]

	DS_Platform_PropertyTypes                map[*DS_Platform_PropertyType]struct{}
	DS_Platform_PropertyTypes_instance       map[*DS_Platform_PropertyType]*DS_Platform_PropertyType
	DS_Platform_PropertyTypes_mapString      map[string]*DS_Platform_PropertyType
	DS_Platform_PropertyTypeOrder            uint
	DS_Platform_PropertyType_stagedOrder     map[*DS_Platform_PropertyType]uint
	DS_Platform_PropertyType_orderStaged     map[uint]*DS_Platform_PropertyType
	DS_Platform_PropertyTypes_reference      map[*DS_Platform_PropertyType]*DS_Platform_PropertyType
	DS_Platform_PropertyTypes_referenceOrder map[*DS_Platform_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_Platform_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_Platform_PropertyType]
	OnAfterDS_Platform_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_Platform_PropertyType]
	OnAfterDS_Platform_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_Platform_PropertyType]

	DS_Platform_Types                map[*DS_Platform_Type]struct{}
	DS_Platform_Types_instance       map[*DS_Platform_Type]*DS_Platform_Type
	DS_Platform_Types_mapString      map[string]*DS_Platform_Type
	DS_Platform_TypeOrder            uint
	DS_Platform_Type_stagedOrder     map[*DS_Platform_Type]uint
	DS_Platform_Type_orderStaged     map[uint]*DS_Platform_Type
	DS_Platform_Types_reference      map[*DS_Platform_Type]*DS_Platform_Type
	DS_Platform_Types_referenceOrder map[*DS_Platform_Type]uint

	// insertion point for slice of pointers maps
	DS_Platform_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_Platform_Type

	DS_Platform_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_Platform_Type

	DS_Platform_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Platform_Type

	DS_Platform_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Platform_Type

	OnAfterDS_Platform_TypeCreateCallback GongOnAfterCreateInterface[DS_Platform_Type]
	OnAfterDS_Platform_TypeUpdateCallback GongOnAfterUpdateInterface[DS_Platform_Type]
	OnAfterDS_Platform_TypeDeleteCallback GongOnAfterDeleteInterface[DS_Platform_Type]

	DS_ProductionSeries_PropertyTypes                map[*DS_ProductionSeries_PropertyType]struct{}
	DS_ProductionSeries_PropertyTypes_instance       map[*DS_ProductionSeries_PropertyType]*DS_ProductionSeries_PropertyType
	DS_ProductionSeries_PropertyTypes_mapString      map[string]*DS_ProductionSeries_PropertyType
	DS_ProductionSeries_PropertyTypeOrder            uint
	DS_ProductionSeries_PropertyType_stagedOrder     map[*DS_ProductionSeries_PropertyType]uint
	DS_ProductionSeries_PropertyType_orderStaged     map[uint]*DS_ProductionSeries_PropertyType
	DS_ProductionSeries_PropertyTypes_reference      map[*DS_ProductionSeries_PropertyType]*DS_ProductionSeries_PropertyType
	DS_ProductionSeries_PropertyTypes_referenceOrder map[*DS_ProductionSeries_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_ProductionSeries_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_ProductionSeries_PropertyType]
	OnAfterDS_ProductionSeries_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_ProductionSeries_PropertyType]
	OnAfterDS_ProductionSeries_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_ProductionSeries_PropertyType]

	DS_ProductionSeries_Types                map[*DS_ProductionSeries_Type]struct{}
	DS_ProductionSeries_Types_instance       map[*DS_ProductionSeries_Type]*DS_ProductionSeries_Type
	DS_ProductionSeries_Types_mapString      map[string]*DS_ProductionSeries_Type
	DS_ProductionSeries_TypeOrder            uint
	DS_ProductionSeries_Type_stagedOrder     map[*DS_ProductionSeries_Type]uint
	DS_ProductionSeries_Type_orderStaged     map[uint]*DS_ProductionSeries_Type
	DS_ProductionSeries_Types_reference      map[*DS_ProductionSeries_Type]*DS_ProductionSeries_Type
	DS_ProductionSeries_Types_referenceOrder map[*DS_ProductionSeries_Type]uint

	// insertion point for slice of pointers maps
	DS_ProductionSeries_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_ProductionSeries_Type

	DS_ProductionSeries_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_ProductionSeries_Type

	DS_ProductionSeries_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_ProductionSeries_Type

	DS_ProductionSeries_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_ProductionSeries_Type

	OnAfterDS_ProductionSeries_TypeCreateCallback GongOnAfterCreateInterface[DS_ProductionSeries_Type]
	OnAfterDS_ProductionSeries_TypeUpdateCallback GongOnAfterUpdateInterface[DS_ProductionSeries_Type]
	OnAfterDS_ProductionSeries_TypeDeleteCallback GongOnAfterDeleteInterface[DS_ProductionSeries_Type]

	DS_Sensor_PropertyTypes                map[*DS_Sensor_PropertyType]struct{}
	DS_Sensor_PropertyTypes_instance       map[*DS_Sensor_PropertyType]*DS_Sensor_PropertyType
	DS_Sensor_PropertyTypes_mapString      map[string]*DS_Sensor_PropertyType
	DS_Sensor_PropertyTypeOrder            uint
	DS_Sensor_PropertyType_stagedOrder     map[*DS_Sensor_PropertyType]uint
	DS_Sensor_PropertyType_orderStaged     map[uint]*DS_Sensor_PropertyType
	DS_Sensor_PropertyTypes_reference      map[*DS_Sensor_PropertyType]*DS_Sensor_PropertyType
	DS_Sensor_PropertyTypes_referenceOrder map[*DS_Sensor_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_Sensor_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_Sensor_PropertyType]
	OnAfterDS_Sensor_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_Sensor_PropertyType]
	OnAfterDS_Sensor_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_Sensor_PropertyType]

	DS_Sensor_Types                map[*DS_Sensor_Type]struct{}
	DS_Sensor_Types_instance       map[*DS_Sensor_Type]*DS_Sensor_Type
	DS_Sensor_Types_mapString      map[string]*DS_Sensor_Type
	DS_Sensor_TypeOrder            uint
	DS_Sensor_Type_stagedOrder     map[*DS_Sensor_Type]uint
	DS_Sensor_Type_orderStaged     map[uint]*DS_Sensor_Type
	DS_Sensor_Types_reference      map[*DS_Sensor_Type]*DS_Sensor_Type
	DS_Sensor_Types_referenceOrder map[*DS_Sensor_Type]uint

	// insertion point for slice of pointers maps
	DS_Sensor_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_Sensor_Type

	DS_Sensor_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_Sensor_Type

	DS_Sensor_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Sensor_Type

	DS_Sensor_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Sensor_Type

	OnAfterDS_Sensor_TypeCreateCallback GongOnAfterCreateInterface[DS_Sensor_Type]
	OnAfterDS_Sensor_TypeUpdateCallback GongOnAfterUpdateInterface[DS_Sensor_Type]
	OnAfterDS_Sensor_TypeDeleteCallback GongOnAfterDeleteInterface[DS_Sensor_Type]

	DS_Series_PropertyTypes                map[*DS_Series_PropertyType]struct{}
	DS_Series_PropertyTypes_instance       map[*DS_Series_PropertyType]*DS_Series_PropertyType
	DS_Series_PropertyTypes_mapString      map[string]*DS_Series_PropertyType
	DS_Series_PropertyTypeOrder            uint
	DS_Series_PropertyType_stagedOrder     map[*DS_Series_PropertyType]uint
	DS_Series_PropertyType_orderStaged     map[uint]*DS_Series_PropertyType
	DS_Series_PropertyTypes_reference      map[*DS_Series_PropertyType]*DS_Series_PropertyType
	DS_Series_PropertyTypes_referenceOrder map[*DS_Series_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_Series_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_Series_PropertyType]
	OnAfterDS_Series_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_Series_PropertyType]
	OnAfterDS_Series_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_Series_PropertyType]

	DS_Series_Types                map[*DS_Series_Type]struct{}
	DS_Series_Types_instance       map[*DS_Series_Type]*DS_Series_Type
	DS_Series_Types_mapString      map[string]*DS_Series_Type
	DS_Series_TypeOrder            uint
	DS_Series_Type_stagedOrder     map[*DS_Series_Type]uint
	DS_Series_Type_orderStaged     map[uint]*DS_Series_Type
	DS_Series_Types_reference      map[*DS_Series_Type]*DS_Series_Type
	DS_Series_Types_referenceOrder map[*DS_Series_Type]uint

	// insertion point for slice of pointers maps
	DS_Series_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_Series_Type

	DS_Series_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_Series_Type

	DS_Series_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Series_Type

	DS_Series_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_Series_Type

	OnAfterDS_Series_TypeCreateCallback GongOnAfterCreateInterface[DS_Series_Type]
	OnAfterDS_Series_TypeUpdateCallback GongOnAfterUpdateInterface[DS_Series_Type]
	OnAfterDS_Series_TypeDeleteCallback GongOnAfterDeleteInterface[DS_Series_Type]

	DS_StereoMate_PropertyTypes                map[*DS_StereoMate_PropertyType]struct{}
	DS_StereoMate_PropertyTypes_instance       map[*DS_StereoMate_PropertyType]*DS_StereoMate_PropertyType
	DS_StereoMate_PropertyTypes_mapString      map[string]*DS_StereoMate_PropertyType
	DS_StereoMate_PropertyTypeOrder            uint
	DS_StereoMate_PropertyType_stagedOrder     map[*DS_StereoMate_PropertyType]uint
	DS_StereoMate_PropertyType_orderStaged     map[uint]*DS_StereoMate_PropertyType
	DS_StereoMate_PropertyTypes_reference      map[*DS_StereoMate_PropertyType]*DS_StereoMate_PropertyType
	DS_StereoMate_PropertyTypes_referenceOrder map[*DS_StereoMate_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterDS_StereoMate_PropertyTypeCreateCallback GongOnAfterCreateInterface[DS_StereoMate_PropertyType]
	OnAfterDS_StereoMate_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[DS_StereoMate_PropertyType]
	OnAfterDS_StereoMate_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[DS_StereoMate_PropertyType]

	DS_StereoMate_Types                map[*DS_StereoMate_Type]struct{}
	DS_StereoMate_Types_instance       map[*DS_StereoMate_Type]*DS_StereoMate_Type
	DS_StereoMate_Types_mapString      map[string]*DS_StereoMate_Type
	DS_StereoMate_TypeOrder            uint
	DS_StereoMate_Type_stagedOrder     map[*DS_StereoMate_Type]uint
	DS_StereoMate_Type_orderStaged     map[uint]*DS_StereoMate_Type
	DS_StereoMate_Types_reference      map[*DS_StereoMate_Type]*DS_StereoMate_Type
	DS_StereoMate_Types_referenceOrder map[*DS_StereoMate_Type]uint

	// insertion point for slice of pointers maps
	DS_StereoMate_Type_ComposedOf_reverseMap map[*DS_DataSet_PropertyType]*DS_StereoMate_Type

	DS_StereoMate_Type_SeriesMetadata_reverseMap map[*MD_Metadata_PropertyType]*DS_StereoMate_Type

	DS_StereoMate_Type_Subset_reverseMap map[*DS_Aggregate_PropertyType]*DS_StereoMate_Type

	DS_StereoMate_Type_Superset_reverseMap map[*DS_Aggregate_PropertyType]*DS_StereoMate_Type

	OnAfterDS_StereoMate_TypeCreateCallback GongOnAfterCreateInterface[DS_StereoMate_Type]
	OnAfterDS_StereoMate_TypeUpdateCallback GongOnAfterUpdateInterface[DS_StereoMate_Type]
	OnAfterDS_StereoMate_TypeDeleteCallback GongOnAfterDeleteInterface[DS_StereoMate_Type]

	MD_Metadata_PropertyTypes                map[*MD_Metadata_PropertyType]struct{}
	MD_Metadata_PropertyTypes_instance       map[*MD_Metadata_PropertyType]*MD_Metadata_PropertyType
	MD_Metadata_PropertyTypes_mapString      map[string]*MD_Metadata_PropertyType
	MD_Metadata_PropertyTypeOrder            uint
	MD_Metadata_PropertyType_stagedOrder     map[*MD_Metadata_PropertyType]uint
	MD_Metadata_PropertyType_orderStaged     map[uint]*MD_Metadata_PropertyType
	MD_Metadata_PropertyTypes_reference      map[*MD_Metadata_PropertyType]*MD_Metadata_PropertyType
	MD_Metadata_PropertyTypes_referenceOrder map[*MD_Metadata_PropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMD_Metadata_PropertyTypeCreateCallback GongOnAfterCreateInterface[MD_Metadata_PropertyType]
	OnAfterMD_Metadata_PropertyTypeUpdateCallback GongOnAfterUpdateInterface[MD_Metadata_PropertyType]
	OnAfterMD_Metadata_PropertyTypeDeleteCallback GongOnAfterDeleteInterface[MD_Metadata_PropertyType]

	MD_Metadata_Types                map[*MD_Metadata_Type]struct{}
	MD_Metadata_Types_instance       map[*MD_Metadata_Type]*MD_Metadata_Type
	MD_Metadata_Types_mapString      map[string]*MD_Metadata_Type
	MD_Metadata_TypeOrder            uint
	MD_Metadata_Type_stagedOrder     map[*MD_Metadata_Type]uint
	MD_Metadata_Type_orderStaged     map[uint]*MD_Metadata_Type
	MD_Metadata_Types_reference      map[*MD_Metadata_Type]*MD_Metadata_Type
	MD_Metadata_Types_referenceOrder map[*MD_Metadata_Type]uint

	// insertion point for slice of pointers maps
	MD_Metadata_Type_Series_reverseMap map[*DS_Aggregate_PropertyType]*MD_Metadata_Type

	MD_Metadata_Type_Describes_reverseMap map[*DS_DataSet_PropertyType]*MD_Metadata_Type

	OnAfterMD_Metadata_TypeCreateCallback GongOnAfterCreateInterface[MD_Metadata_Type]
	OnAfterMD_Metadata_TypeUpdateCallback GongOnAfterUpdateInterface[MD_Metadata_Type]
	OnAfterMD_Metadata_TypeDeleteCallback GongOnAfterDeleteInterface[MD_Metadata_Type]

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
	__gong__clearReferences(&stage.AbstractDS_Aggregate_Types_reference, &stage.AbstractDS_Aggregate_Types_instance, &stage.AbstractDS_Aggregate_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_Aggregate_PropertyTypes_reference, &stage.DS_Aggregate_PropertyTypes_instance, &stage.DS_Aggregate_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_DataSet_PropertyTypes_reference, &stage.DS_DataSet_PropertyTypes_instance, &stage.DS_DataSet_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_DataSet_Types_reference, &stage.DS_DataSet_Types_instance, &stage.DS_DataSet_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_Initiative_PropertyTypes_reference, &stage.DS_Initiative_PropertyTypes_instance, &stage.DS_Initiative_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_Initiative_Types_reference, &stage.DS_Initiative_Types_instance, &stage.DS_Initiative_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_OtherAggregate_PropertyTypes_reference, &stage.DS_OtherAggregate_PropertyTypes_instance, &stage.DS_OtherAggregate_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_OtherAggregate_Types_reference, &stage.DS_OtherAggregate_Types_instance, &stage.DS_OtherAggregate_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_Platform_PropertyTypes_reference, &stage.DS_Platform_PropertyTypes_instance, &stage.DS_Platform_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_Platform_Types_reference, &stage.DS_Platform_Types_instance, &stage.DS_Platform_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_ProductionSeries_PropertyTypes_reference, &stage.DS_ProductionSeries_PropertyTypes_instance, &stage.DS_ProductionSeries_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_ProductionSeries_Types_reference, &stage.DS_ProductionSeries_Types_instance, &stage.DS_ProductionSeries_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_Sensor_PropertyTypes_reference, &stage.DS_Sensor_PropertyTypes_instance, &stage.DS_Sensor_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_Sensor_Types_reference, &stage.DS_Sensor_Types_instance, &stage.DS_Sensor_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_Series_PropertyTypes_reference, &stage.DS_Series_PropertyTypes_instance, &stage.DS_Series_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_Series_Types_reference, &stage.DS_Series_Types_instance, &stage.DS_Series_Types_referenceOrder)

	__gong__clearReferences(&stage.DS_StereoMate_PropertyTypes_reference, &stage.DS_StereoMate_PropertyTypes_instance, &stage.DS_StereoMate_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.DS_StereoMate_Types_reference, &stage.DS_StereoMate_Types_instance, &stage.DS_StereoMate_Types_referenceOrder)

	__gong__clearReferences(&stage.MD_Metadata_PropertyTypes_reference, &stage.MD_Metadata_PropertyTypes_instance, &stage.MD_Metadata_PropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.MD_Metadata_Types_reference, &stage.MD_Metadata_Types_instance, &stage.MD_Metadata_Types_referenceOrder)

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
	stage.AbstractDS_Aggregate_TypeOrder = __gong__recomputeOrder(stage.AbstractDS_Aggregate_Type_stagedOrder)

	stage.DS_Aggregate_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_Aggregate_PropertyType_stagedOrder)

	stage.DS_DataSet_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_DataSet_PropertyType_stagedOrder)

	stage.DS_DataSet_TypeOrder = __gong__recomputeOrder(stage.DS_DataSet_Type_stagedOrder)

	stage.DS_Initiative_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_Initiative_PropertyType_stagedOrder)

	stage.DS_Initiative_TypeOrder = __gong__recomputeOrder(stage.DS_Initiative_Type_stagedOrder)

	stage.DS_OtherAggregate_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_OtherAggregate_PropertyType_stagedOrder)

	stage.DS_OtherAggregate_TypeOrder = __gong__recomputeOrder(stage.DS_OtherAggregate_Type_stagedOrder)

	stage.DS_Platform_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_Platform_PropertyType_stagedOrder)

	stage.DS_Platform_TypeOrder = __gong__recomputeOrder(stage.DS_Platform_Type_stagedOrder)

	stage.DS_ProductionSeries_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_ProductionSeries_PropertyType_stagedOrder)

	stage.DS_ProductionSeries_TypeOrder = __gong__recomputeOrder(stage.DS_ProductionSeries_Type_stagedOrder)

	stage.DS_Sensor_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_Sensor_PropertyType_stagedOrder)

	stage.DS_Sensor_TypeOrder = __gong__recomputeOrder(stage.DS_Sensor_Type_stagedOrder)

	stage.DS_Series_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_Series_PropertyType_stagedOrder)

	stage.DS_Series_TypeOrder = __gong__recomputeOrder(stage.DS_Series_Type_stagedOrder)

	stage.DS_StereoMate_PropertyTypeOrder = __gong__recomputeOrder(stage.DS_StereoMate_PropertyType_stagedOrder)

	stage.DS_StereoMate_TypeOrder = __gong__recomputeOrder(stage.DS_StereoMate_Type_stagedOrder)

	stage.MD_Metadata_PropertyTypeOrder = __gong__recomputeOrder(stage.MD_Metadata_PropertyType_stagedOrder)

	stage.MD_Metadata_TypeOrder = __gong__recomputeOrder(stage.MD_Metadata_Type_stagedOrder)

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
func (*AbstractDS_Aggregate_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Type_stagedOrder)
}

func (*AbstractDS_Aggregate_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractDS_Aggregate_Type_orderStaged[order]
}

func (*AbstractDS_Aggregate_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractDS_Aggregate_Types_mapString
}

func (*AbstractDS_Aggregate_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractDS_Aggregate_Types
}

func (*AbstractDS_Aggregate_Type) GongNewInstance() any {
	return new(AbstractDS_Aggregate_Type)
}

func (*DS_Aggregate_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyType_stagedOrder)
}

func (*DS_Aggregate_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Aggregate_PropertyType_orderStaged[order]
}

func (*DS_Aggregate_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Aggregate_PropertyTypes_mapString
}

func (*DS_Aggregate_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Aggregate_PropertyTypes
}

func (*DS_Aggregate_PropertyType) GongNewInstance() any {
	return new(DS_Aggregate_PropertyType)
}

func (*DS_DataSet_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyType_stagedOrder)
}

func (*DS_DataSet_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_DataSet_PropertyType_orderStaged[order]
}

func (*DS_DataSet_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_DataSet_PropertyTypes_mapString
}

func (*DS_DataSet_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_DataSet_PropertyTypes
}

func (*DS_DataSet_PropertyType) GongNewInstance() any {
	return new(DS_DataSet_PropertyType)
}

func (*DS_DataSet_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_DataSet_Types, stage.DS_DataSet_Type_stagedOrder)
}

func (*DS_DataSet_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_DataSet_Type_orderStaged[order]
}

func (*DS_DataSet_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_DataSet_Types_mapString
}

func (*DS_DataSet_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_DataSet_Types
}

func (*DS_DataSet_Type) GongNewInstance() any {
	return new(DS_DataSet_Type)
}

func (*DS_Initiative_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyType_stagedOrder)
}

func (*DS_Initiative_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Initiative_PropertyType_orderStaged[order]
}

func (*DS_Initiative_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Initiative_PropertyTypes_mapString
}

func (*DS_Initiative_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Initiative_PropertyTypes
}

func (*DS_Initiative_PropertyType) GongNewInstance() any {
	return new(DS_Initiative_PropertyType)
}

func (*DS_Initiative_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Initiative_Types, stage.DS_Initiative_Type_stagedOrder)
}

func (*DS_Initiative_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Initiative_Type_orderStaged[order]
}

func (*DS_Initiative_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Initiative_Types_mapString
}

func (*DS_Initiative_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Initiative_Types
}

func (*DS_Initiative_Type) GongNewInstance() any {
	return new(DS_Initiative_Type)
}

func (*DS_OtherAggregate_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyType_stagedOrder)
}

func (*DS_OtherAggregate_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_OtherAggregate_PropertyType_orderStaged[order]
}

func (*DS_OtherAggregate_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_OtherAggregate_PropertyTypes_mapString
}

func (*DS_OtherAggregate_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_OtherAggregate_PropertyTypes
}

func (*DS_OtherAggregate_PropertyType) GongNewInstance() any {
	return new(DS_OtherAggregate_PropertyType)
}

func (*DS_OtherAggregate_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Type_stagedOrder)
}

func (*DS_OtherAggregate_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_OtherAggregate_Type_orderStaged[order]
}

func (*DS_OtherAggregate_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_OtherAggregate_Types_mapString
}

func (*DS_OtherAggregate_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_OtherAggregate_Types
}

func (*DS_OtherAggregate_Type) GongNewInstance() any {
	return new(DS_OtherAggregate_Type)
}

func (*DS_Platform_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyType_stagedOrder)
}

func (*DS_Platform_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Platform_PropertyType_orderStaged[order]
}

func (*DS_Platform_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Platform_PropertyTypes_mapString
}

func (*DS_Platform_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Platform_PropertyTypes
}

func (*DS_Platform_PropertyType) GongNewInstance() any {
	return new(DS_Platform_PropertyType)
}

func (*DS_Platform_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Platform_Types, stage.DS_Platform_Type_stagedOrder)
}

func (*DS_Platform_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Platform_Type_orderStaged[order]
}

func (*DS_Platform_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Platform_Types_mapString
}

func (*DS_Platform_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Platform_Types
}

func (*DS_Platform_Type) GongNewInstance() any {
	return new(DS_Platform_Type)
}

func (*DS_ProductionSeries_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyType_stagedOrder)
}

func (*DS_ProductionSeries_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_ProductionSeries_PropertyType_orderStaged[order]
}

func (*DS_ProductionSeries_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_ProductionSeries_PropertyTypes_mapString
}

func (*DS_ProductionSeries_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_ProductionSeries_PropertyTypes
}

func (*DS_ProductionSeries_PropertyType) GongNewInstance() any {
	return new(DS_ProductionSeries_PropertyType)
}

func (*DS_ProductionSeries_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Type_stagedOrder)
}

func (*DS_ProductionSeries_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_ProductionSeries_Type_orderStaged[order]
}

func (*DS_ProductionSeries_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_ProductionSeries_Types_mapString
}

func (*DS_ProductionSeries_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_ProductionSeries_Types
}

func (*DS_ProductionSeries_Type) GongNewInstance() any {
	return new(DS_ProductionSeries_Type)
}

func (*DS_Sensor_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyType_stagedOrder)
}

func (*DS_Sensor_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Sensor_PropertyType_orderStaged[order]
}

func (*DS_Sensor_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Sensor_PropertyTypes_mapString
}

func (*DS_Sensor_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Sensor_PropertyTypes
}

func (*DS_Sensor_PropertyType) GongNewInstance() any {
	return new(DS_Sensor_PropertyType)
}

func (*DS_Sensor_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Sensor_Types, stage.DS_Sensor_Type_stagedOrder)
}

func (*DS_Sensor_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Sensor_Type_orderStaged[order]
}

func (*DS_Sensor_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Sensor_Types_mapString
}

func (*DS_Sensor_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Sensor_Types
}

func (*DS_Sensor_Type) GongNewInstance() any {
	return new(DS_Sensor_Type)
}

func (*DS_Series_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyType_stagedOrder)
}

func (*DS_Series_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Series_PropertyType_orderStaged[order]
}

func (*DS_Series_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Series_PropertyTypes_mapString
}

func (*DS_Series_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Series_PropertyTypes
}

func (*DS_Series_PropertyType) GongNewInstance() any {
	return new(DS_Series_PropertyType)
}

func (*DS_Series_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_Series_Types, stage.DS_Series_Type_stagedOrder)
}

func (*DS_Series_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_Series_Type_orderStaged[order]
}

func (*DS_Series_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_Series_Types_mapString
}

func (*DS_Series_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_Series_Types
}

func (*DS_Series_Type) GongNewInstance() any {
	return new(DS_Series_Type)
}

func (*DS_StereoMate_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyType_stagedOrder)
}

func (*DS_StereoMate_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_StereoMate_PropertyType_orderStaged[order]
}

func (*DS_StereoMate_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_StereoMate_PropertyTypes_mapString
}

func (*DS_StereoMate_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_StereoMate_PropertyTypes
}

func (*DS_StereoMate_PropertyType) GongNewInstance() any {
	return new(DS_StereoMate_PropertyType)
}

func (*DS_StereoMate_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DS_StereoMate_Types, stage.DS_StereoMate_Type_stagedOrder)
}

func (*DS_StereoMate_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DS_StereoMate_Type_orderStaged[order]
}

func (*DS_StereoMate_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DS_StereoMate_Types_mapString
}

func (*DS_StereoMate_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.DS_StereoMate_Types
}

func (*DS_StereoMate_Type) GongNewInstance() any {
	return new(DS_StereoMate_Type)
}

func (*MD_Metadata_PropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyType_stagedOrder)
}

func (*MD_Metadata_PropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Metadata_PropertyType_orderStaged[order]
}

func (*MD_Metadata_PropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Metadata_PropertyTypes_mapString
}

func (*MD_Metadata_PropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Metadata_PropertyTypes
}

func (*MD_Metadata_PropertyType) GongNewInstance() any {
	return new(MD_Metadata_PropertyType)
}

func (*MD_Metadata_Type) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MD_Metadata_Types, stage.MD_Metadata_Type_stagedOrder)
}

func (*MD_Metadata_Type) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MD_Metadata_Type_orderStaged[order]
}

func (*MD_Metadata_Type) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MD_Metadata_Types_mapString
}

func (*MD_Metadata_Type) GongGetInstancesSet(stage *Stage) any {
	return &stage.MD_Metadata_Types
}

func (*MD_Metadata_Type) GongNewInstance() any {
	return new(MD_Metadata_Type)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd"
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
		AbstractDS_Aggregate_Types:           make(map[*AbstractDS_Aggregate_Type]struct{}),
		AbstractDS_Aggregate_Types_mapString: make(map[string]*AbstractDS_Aggregate_Type),

		DS_Aggregate_PropertyTypes:           make(map[*DS_Aggregate_PropertyType]struct{}),
		DS_Aggregate_PropertyTypes_mapString: make(map[string]*DS_Aggregate_PropertyType),

		DS_DataSet_PropertyTypes:           make(map[*DS_DataSet_PropertyType]struct{}),
		DS_DataSet_PropertyTypes_mapString: make(map[string]*DS_DataSet_PropertyType),

		DS_DataSet_Types:           make(map[*DS_DataSet_Type]struct{}),
		DS_DataSet_Types_mapString: make(map[string]*DS_DataSet_Type),

		DS_Initiative_PropertyTypes:           make(map[*DS_Initiative_PropertyType]struct{}),
		DS_Initiative_PropertyTypes_mapString: make(map[string]*DS_Initiative_PropertyType),

		DS_Initiative_Types:           make(map[*DS_Initiative_Type]struct{}),
		DS_Initiative_Types_mapString: make(map[string]*DS_Initiative_Type),

		DS_OtherAggregate_PropertyTypes:           make(map[*DS_OtherAggregate_PropertyType]struct{}),
		DS_OtherAggregate_PropertyTypes_mapString: make(map[string]*DS_OtherAggregate_PropertyType),

		DS_OtherAggregate_Types:           make(map[*DS_OtherAggregate_Type]struct{}),
		DS_OtherAggregate_Types_mapString: make(map[string]*DS_OtherAggregate_Type),

		DS_Platform_PropertyTypes:           make(map[*DS_Platform_PropertyType]struct{}),
		DS_Platform_PropertyTypes_mapString: make(map[string]*DS_Platform_PropertyType),

		DS_Platform_Types:           make(map[*DS_Platform_Type]struct{}),
		DS_Platform_Types_mapString: make(map[string]*DS_Platform_Type),

		DS_ProductionSeries_PropertyTypes:           make(map[*DS_ProductionSeries_PropertyType]struct{}),
		DS_ProductionSeries_PropertyTypes_mapString: make(map[string]*DS_ProductionSeries_PropertyType),

		DS_ProductionSeries_Types:           make(map[*DS_ProductionSeries_Type]struct{}),
		DS_ProductionSeries_Types_mapString: make(map[string]*DS_ProductionSeries_Type),

		DS_Sensor_PropertyTypes:           make(map[*DS_Sensor_PropertyType]struct{}),
		DS_Sensor_PropertyTypes_mapString: make(map[string]*DS_Sensor_PropertyType),

		DS_Sensor_Types:           make(map[*DS_Sensor_Type]struct{}),
		DS_Sensor_Types_mapString: make(map[string]*DS_Sensor_Type),

		DS_Series_PropertyTypes:           make(map[*DS_Series_PropertyType]struct{}),
		DS_Series_PropertyTypes_mapString: make(map[string]*DS_Series_PropertyType),

		DS_Series_Types:           make(map[*DS_Series_Type]struct{}),
		DS_Series_Types_mapString: make(map[string]*DS_Series_Type),

		DS_StereoMate_PropertyTypes:           make(map[*DS_StereoMate_PropertyType]struct{}),
		DS_StereoMate_PropertyTypes_mapString: make(map[string]*DS_StereoMate_PropertyType),

		DS_StereoMate_Types:           make(map[*DS_StereoMate_Type]struct{}),
		DS_StereoMate_Types_mapString: make(map[string]*DS_StereoMate_Type),

		MD_Metadata_PropertyTypes:           make(map[*MD_Metadata_PropertyType]struct{}),
		MD_Metadata_PropertyTypes_mapString: make(map[string]*MD_Metadata_PropertyType),

		MD_Metadata_Types:           make(map[*MD_Metadata_Type]struct{}),
		MD_Metadata_Types_mapString: make(map[string]*MD_Metadata_Type),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AbstractDS_Aggregate_Type_stagedOrder: make(map[*AbstractDS_Aggregate_Type]uint),
		AbstractDS_Aggregate_Type_orderStaged: make(map[uint]*AbstractDS_Aggregate_Type),
		AbstractDS_Aggregate_Types_reference:  make(map[*AbstractDS_Aggregate_Type]*AbstractDS_Aggregate_Type),

		DS_Aggregate_PropertyType_stagedOrder: make(map[*DS_Aggregate_PropertyType]uint),
		DS_Aggregate_PropertyType_orderStaged: make(map[uint]*DS_Aggregate_PropertyType),
		DS_Aggregate_PropertyTypes_reference:  make(map[*DS_Aggregate_PropertyType]*DS_Aggregate_PropertyType),

		DS_DataSet_PropertyType_stagedOrder: make(map[*DS_DataSet_PropertyType]uint),
		DS_DataSet_PropertyType_orderStaged: make(map[uint]*DS_DataSet_PropertyType),
		DS_DataSet_PropertyTypes_reference:  make(map[*DS_DataSet_PropertyType]*DS_DataSet_PropertyType),

		DS_DataSet_Type_stagedOrder: make(map[*DS_DataSet_Type]uint),
		DS_DataSet_Type_orderStaged: make(map[uint]*DS_DataSet_Type),
		DS_DataSet_Types_reference:  make(map[*DS_DataSet_Type]*DS_DataSet_Type),

		DS_Initiative_PropertyType_stagedOrder: make(map[*DS_Initiative_PropertyType]uint),
		DS_Initiative_PropertyType_orderStaged: make(map[uint]*DS_Initiative_PropertyType),
		DS_Initiative_PropertyTypes_reference:  make(map[*DS_Initiative_PropertyType]*DS_Initiative_PropertyType),

		DS_Initiative_Type_stagedOrder: make(map[*DS_Initiative_Type]uint),
		DS_Initiative_Type_orderStaged: make(map[uint]*DS_Initiative_Type),
		DS_Initiative_Types_reference:  make(map[*DS_Initiative_Type]*DS_Initiative_Type),

		DS_OtherAggregate_PropertyType_stagedOrder: make(map[*DS_OtherAggregate_PropertyType]uint),
		DS_OtherAggregate_PropertyType_orderStaged: make(map[uint]*DS_OtherAggregate_PropertyType),
		DS_OtherAggregate_PropertyTypes_reference:  make(map[*DS_OtherAggregate_PropertyType]*DS_OtherAggregate_PropertyType),

		DS_OtherAggregate_Type_stagedOrder: make(map[*DS_OtherAggregate_Type]uint),
		DS_OtherAggregate_Type_orderStaged: make(map[uint]*DS_OtherAggregate_Type),
		DS_OtherAggregate_Types_reference:  make(map[*DS_OtherAggregate_Type]*DS_OtherAggregate_Type),

		DS_Platform_PropertyType_stagedOrder: make(map[*DS_Platform_PropertyType]uint),
		DS_Platform_PropertyType_orderStaged: make(map[uint]*DS_Platform_PropertyType),
		DS_Platform_PropertyTypes_reference:  make(map[*DS_Platform_PropertyType]*DS_Platform_PropertyType),

		DS_Platform_Type_stagedOrder: make(map[*DS_Platform_Type]uint),
		DS_Platform_Type_orderStaged: make(map[uint]*DS_Platform_Type),
		DS_Platform_Types_reference:  make(map[*DS_Platform_Type]*DS_Platform_Type),

		DS_ProductionSeries_PropertyType_stagedOrder: make(map[*DS_ProductionSeries_PropertyType]uint),
		DS_ProductionSeries_PropertyType_orderStaged: make(map[uint]*DS_ProductionSeries_PropertyType),
		DS_ProductionSeries_PropertyTypes_reference:  make(map[*DS_ProductionSeries_PropertyType]*DS_ProductionSeries_PropertyType),

		DS_ProductionSeries_Type_stagedOrder: make(map[*DS_ProductionSeries_Type]uint),
		DS_ProductionSeries_Type_orderStaged: make(map[uint]*DS_ProductionSeries_Type),
		DS_ProductionSeries_Types_reference:  make(map[*DS_ProductionSeries_Type]*DS_ProductionSeries_Type),

		DS_Sensor_PropertyType_stagedOrder: make(map[*DS_Sensor_PropertyType]uint),
		DS_Sensor_PropertyType_orderStaged: make(map[uint]*DS_Sensor_PropertyType),
		DS_Sensor_PropertyTypes_reference:  make(map[*DS_Sensor_PropertyType]*DS_Sensor_PropertyType),

		DS_Sensor_Type_stagedOrder: make(map[*DS_Sensor_Type]uint),
		DS_Sensor_Type_orderStaged: make(map[uint]*DS_Sensor_Type),
		DS_Sensor_Types_reference:  make(map[*DS_Sensor_Type]*DS_Sensor_Type),

		DS_Series_PropertyType_stagedOrder: make(map[*DS_Series_PropertyType]uint),
		DS_Series_PropertyType_orderStaged: make(map[uint]*DS_Series_PropertyType),
		DS_Series_PropertyTypes_reference:  make(map[*DS_Series_PropertyType]*DS_Series_PropertyType),

		DS_Series_Type_stagedOrder: make(map[*DS_Series_Type]uint),
		DS_Series_Type_orderStaged: make(map[uint]*DS_Series_Type),
		DS_Series_Types_reference:  make(map[*DS_Series_Type]*DS_Series_Type),

		DS_StereoMate_PropertyType_stagedOrder: make(map[*DS_StereoMate_PropertyType]uint),
		DS_StereoMate_PropertyType_orderStaged: make(map[uint]*DS_StereoMate_PropertyType),
		DS_StereoMate_PropertyTypes_reference:  make(map[*DS_StereoMate_PropertyType]*DS_StereoMate_PropertyType),

		DS_StereoMate_Type_stagedOrder: make(map[*DS_StereoMate_Type]uint),
		DS_StereoMate_Type_orderStaged: make(map[uint]*DS_StereoMate_Type),
		DS_StereoMate_Types_reference:  make(map[*DS_StereoMate_Type]*DS_StereoMate_Type),

		MD_Metadata_PropertyType_stagedOrder: make(map[*MD_Metadata_PropertyType]uint),
		MD_Metadata_PropertyType_orderStaged: make(map[uint]*MD_Metadata_PropertyType),
		MD_Metadata_PropertyTypes_reference:  make(map[*MD_Metadata_PropertyType]*MD_Metadata_PropertyType),

		MD_Metadata_Type_stagedOrder: make(map[*MD_Metadata_Type]uint),
		MD_Metadata_Type_orderStaged: make(map[uint]*MD_Metadata_Type),
		MD_Metadata_Types_reference:  make(map[*MD_Metadata_Type]*MD_Metadata_Type),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AbstractDS_Aggregate_Type": &AbstractDS_Aggregate_TypeUnmarshaller{},

			"DS_Aggregate_PropertyType": &DS_Aggregate_PropertyTypeUnmarshaller{},

			"DS_DataSet_PropertyType": &DS_DataSet_PropertyTypeUnmarshaller{},

			"DS_DataSet_Type": &DS_DataSet_TypeUnmarshaller{},

			"DS_Initiative_PropertyType": &DS_Initiative_PropertyTypeUnmarshaller{},

			"DS_Initiative_Type": &DS_Initiative_TypeUnmarshaller{},

			"DS_OtherAggregate_PropertyType": &DS_OtherAggregate_PropertyTypeUnmarshaller{},

			"DS_OtherAggregate_Type": &DS_OtherAggregate_TypeUnmarshaller{},

			"DS_Platform_PropertyType": &DS_Platform_PropertyTypeUnmarshaller{},

			"DS_Platform_Type": &DS_Platform_TypeUnmarshaller{},

			"DS_ProductionSeries_PropertyType": &DS_ProductionSeries_PropertyTypeUnmarshaller{},

			"DS_ProductionSeries_Type": &DS_ProductionSeries_TypeUnmarshaller{},

			"DS_Sensor_PropertyType": &DS_Sensor_PropertyTypeUnmarshaller{},

			"DS_Sensor_Type": &DS_Sensor_TypeUnmarshaller{},

			"DS_Series_PropertyType": &DS_Series_PropertyTypeUnmarshaller{},

			"DS_Series_Type": &DS_Series_TypeUnmarshaller{},

			"DS_StereoMate_PropertyType": &DS_StereoMate_PropertyTypeUnmarshaller{},

			"DS_StereoMate_Type": &DS_StereoMate_TypeUnmarshaller{},

			"MD_Metadata_PropertyType": &MD_Metadata_PropertyTypeUnmarshaller{},

			"MD_Metadata_Type": &MD_Metadata_TypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AbstractDS_Aggregate_Type"] = len(stage.AbstractDS_Aggregate_Types)
	stage.Map_GongStructName_InstancesNb["DS_Aggregate_PropertyType"] = len(stage.DS_Aggregate_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_DataSet_PropertyType"] = len(stage.DS_DataSet_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_DataSet_Type"] = len(stage.DS_DataSet_Types)
	stage.Map_GongStructName_InstancesNb["DS_Initiative_PropertyType"] = len(stage.DS_Initiative_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_Initiative_Type"] = len(stage.DS_Initiative_Types)
	stage.Map_GongStructName_InstancesNb["DS_OtherAggregate_PropertyType"] = len(stage.DS_OtherAggregate_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_OtherAggregate_Type"] = len(stage.DS_OtherAggregate_Types)
	stage.Map_GongStructName_InstancesNb["DS_Platform_PropertyType"] = len(stage.DS_Platform_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_Platform_Type"] = len(stage.DS_Platform_Types)
	stage.Map_GongStructName_InstancesNb["DS_ProductionSeries_PropertyType"] = len(stage.DS_ProductionSeries_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_ProductionSeries_Type"] = len(stage.DS_ProductionSeries_Types)
	stage.Map_GongStructName_InstancesNb["DS_Sensor_PropertyType"] = len(stage.DS_Sensor_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_Sensor_Type"] = len(stage.DS_Sensor_Types)
	stage.Map_GongStructName_InstancesNb["DS_Series_PropertyType"] = len(stage.DS_Series_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_Series_Type"] = len(stage.DS_Series_Types)
	stage.Map_GongStructName_InstancesNb["DS_StereoMate_PropertyType"] = len(stage.DS_StereoMate_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["DS_StereoMate_Type"] = len(stage.DS_StereoMate_Types)
	stage.Map_GongStructName_InstancesNb["MD_Metadata_PropertyType"] = len(stage.MD_Metadata_PropertyTypes)
	stage.Map_GongStructName_InstancesNb["MD_Metadata_Type"] = len(stage.MD_Metadata_Types)
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
// Stage puts abstractds_aggregate_type to the model stage
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) Stage(stage *Stage) *AbstractDS_Aggregate_Type {
	__gong__stage(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Type_stagedOrder, stage.AbstractDS_Aggregate_Type_orderStaged, &stage.AbstractDS_Aggregate_TypeOrder, stage.AbstractDS_Aggregate_Types_mapString, abstractds_aggregate_type, abstractds_aggregate_type.Name)
	return abstractds_aggregate_type
}

// StagePreserveOrder puts abstractds_aggregate_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractDS_Aggregate_TypeOrder
// - update stage.AbstractDS_Aggregate_TypeOrder accordingly
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Type_stagedOrder, stage.AbstractDS_Aggregate_Type_orderStaged, &stage.AbstractDS_Aggregate_TypeOrder, stage.AbstractDS_Aggregate_Types_mapString, abstractds_aggregate_type, order, abstractds_aggregate_type.Name)
}

// Unstage removes abstractds_aggregate_type off the model stage
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) Unstage(stage *Stage) *AbstractDS_Aggregate_Type {
	__gong__unstage(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Types_mapString, abstractds_aggregate_type, abstractds_aggregate_type.Name)
	return abstractds_aggregate_type
}

// UnstageVoid removes abstractds_aggregate_type off the model stage
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) UnstageVoid(stage *Stage) {
	abstractds_aggregate_type.Unstage(stage)
}

func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) StageVoid(stage *Stage) {
	abstractds_aggregate_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GetName() (res string) {
	return abstractds_aggregate_type.Name
}

// for satisfaction of GongStruct interface
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) SetName(name string) {
	abstractds_aggregate_type.Name = name
}

// Stage puts ds_aggregate_propertytype to the model stage
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) Stage(stage *Stage) *DS_Aggregate_PropertyType {
	__gong__stage(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyType_stagedOrder, stage.DS_Aggregate_PropertyType_orderStaged, &stage.DS_Aggregate_PropertyTypeOrder, stage.DS_Aggregate_PropertyTypes_mapString, ds_aggregate_propertytype, ds_aggregate_propertytype.Name)
	return ds_aggregate_propertytype
}

// StagePreserveOrder puts ds_aggregate_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Aggregate_PropertyTypeOrder
// - update stage.DS_Aggregate_PropertyTypeOrder accordingly
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyType_stagedOrder, stage.DS_Aggregate_PropertyType_orderStaged, &stage.DS_Aggregate_PropertyTypeOrder, stage.DS_Aggregate_PropertyTypes_mapString, ds_aggregate_propertytype, order, ds_aggregate_propertytype.Name)
}

// Unstage removes ds_aggregate_propertytype off the model stage
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) Unstage(stage *Stage) *DS_Aggregate_PropertyType {
	__gong__unstage(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyTypes_mapString, ds_aggregate_propertytype, ds_aggregate_propertytype.Name)
	return ds_aggregate_propertytype
}

// UnstageVoid removes ds_aggregate_propertytype off the model stage
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) UnstageVoid(stage *Stage) {
	ds_aggregate_propertytype.Unstage(stage)
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) StageVoid(stage *Stage) {
	ds_aggregate_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GetName() (res string) {
	return ds_aggregate_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) SetName(name string) {
	ds_aggregate_propertytype.Name = name
}

// Stage puts ds_dataset_propertytype to the model stage
func (ds_dataset_propertytype *DS_DataSet_PropertyType) Stage(stage *Stage) *DS_DataSet_PropertyType {
	__gong__stage(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyType_stagedOrder, stage.DS_DataSet_PropertyType_orderStaged, &stage.DS_DataSet_PropertyTypeOrder, stage.DS_DataSet_PropertyTypes_mapString, ds_dataset_propertytype, ds_dataset_propertytype.Name)
	return ds_dataset_propertytype
}

// StagePreserveOrder puts ds_dataset_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_DataSet_PropertyTypeOrder
// - update stage.DS_DataSet_PropertyTypeOrder accordingly
func (ds_dataset_propertytype *DS_DataSet_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyType_stagedOrder, stage.DS_DataSet_PropertyType_orderStaged, &stage.DS_DataSet_PropertyTypeOrder, stage.DS_DataSet_PropertyTypes_mapString, ds_dataset_propertytype, order, ds_dataset_propertytype.Name)
}

// Unstage removes ds_dataset_propertytype off the model stage
func (ds_dataset_propertytype *DS_DataSet_PropertyType) Unstage(stage *Stage) *DS_DataSet_PropertyType {
	__gong__unstage(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyTypes_mapString, ds_dataset_propertytype, ds_dataset_propertytype.Name)
	return ds_dataset_propertytype
}

// UnstageVoid removes ds_dataset_propertytype off the model stage
func (ds_dataset_propertytype *DS_DataSet_PropertyType) UnstageVoid(stage *Stage) {
	ds_dataset_propertytype.Unstage(stage)
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) StageVoid(stage *Stage) {
	ds_dataset_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_dataset_propertytype *DS_DataSet_PropertyType) GetName() (res string) {
	return ds_dataset_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_dataset_propertytype *DS_DataSet_PropertyType) SetName(name string) {
	ds_dataset_propertytype.Name = name
}

// Stage puts ds_dataset_type to the model stage
func (ds_dataset_type *DS_DataSet_Type) Stage(stage *Stage) *DS_DataSet_Type {
	__gong__stage(stage.DS_DataSet_Types, stage.DS_DataSet_Type_stagedOrder, stage.DS_DataSet_Type_orderStaged, &stage.DS_DataSet_TypeOrder, stage.DS_DataSet_Types_mapString, ds_dataset_type, ds_dataset_type.Name)
	return ds_dataset_type
}

// StagePreserveOrder puts ds_dataset_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_DataSet_TypeOrder
// - update stage.DS_DataSet_TypeOrder accordingly
func (ds_dataset_type *DS_DataSet_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_DataSet_Types, stage.DS_DataSet_Type_stagedOrder, stage.DS_DataSet_Type_orderStaged, &stage.DS_DataSet_TypeOrder, stage.DS_DataSet_Types_mapString, ds_dataset_type, order, ds_dataset_type.Name)
}

// Unstage removes ds_dataset_type off the model stage
func (ds_dataset_type *DS_DataSet_Type) Unstage(stage *Stage) *DS_DataSet_Type {
	__gong__unstage(stage.DS_DataSet_Types, stage.DS_DataSet_Types_mapString, ds_dataset_type, ds_dataset_type.Name)
	return ds_dataset_type
}

// UnstageVoid removes ds_dataset_type off the model stage
func (ds_dataset_type *DS_DataSet_Type) UnstageVoid(stage *Stage) {
	ds_dataset_type.Unstage(stage)
}

func (ds_dataset_type *DS_DataSet_Type) StageVoid(stage *Stage) {
	ds_dataset_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_dataset_type *DS_DataSet_Type) GetName() (res string) {
	return ds_dataset_type.Name
}

// for satisfaction of GongStruct interface
func (ds_dataset_type *DS_DataSet_Type) SetName(name string) {
	ds_dataset_type.Name = name
}

// Stage puts ds_initiative_propertytype to the model stage
func (ds_initiative_propertytype *DS_Initiative_PropertyType) Stage(stage *Stage) *DS_Initiative_PropertyType {
	__gong__stage(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyType_stagedOrder, stage.DS_Initiative_PropertyType_orderStaged, &stage.DS_Initiative_PropertyTypeOrder, stage.DS_Initiative_PropertyTypes_mapString, ds_initiative_propertytype, ds_initiative_propertytype.Name)
	return ds_initiative_propertytype
}

// StagePreserveOrder puts ds_initiative_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Initiative_PropertyTypeOrder
// - update stage.DS_Initiative_PropertyTypeOrder accordingly
func (ds_initiative_propertytype *DS_Initiative_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyType_stagedOrder, stage.DS_Initiative_PropertyType_orderStaged, &stage.DS_Initiative_PropertyTypeOrder, stage.DS_Initiative_PropertyTypes_mapString, ds_initiative_propertytype, order, ds_initiative_propertytype.Name)
}

// Unstage removes ds_initiative_propertytype off the model stage
func (ds_initiative_propertytype *DS_Initiative_PropertyType) Unstage(stage *Stage) *DS_Initiative_PropertyType {
	__gong__unstage(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyTypes_mapString, ds_initiative_propertytype, ds_initiative_propertytype.Name)
	return ds_initiative_propertytype
}

// UnstageVoid removes ds_initiative_propertytype off the model stage
func (ds_initiative_propertytype *DS_Initiative_PropertyType) UnstageVoid(stage *Stage) {
	ds_initiative_propertytype.Unstage(stage)
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) StageVoid(stage *Stage) {
	ds_initiative_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_initiative_propertytype *DS_Initiative_PropertyType) GetName() (res string) {
	return ds_initiative_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_initiative_propertytype *DS_Initiative_PropertyType) SetName(name string) {
	ds_initiative_propertytype.Name = name
}

// Stage puts ds_initiative_type to the model stage
func (ds_initiative_type *DS_Initiative_Type) Stage(stage *Stage) *DS_Initiative_Type {
	__gong__stage(stage.DS_Initiative_Types, stage.DS_Initiative_Type_stagedOrder, stage.DS_Initiative_Type_orderStaged, &stage.DS_Initiative_TypeOrder, stage.DS_Initiative_Types_mapString, ds_initiative_type, ds_initiative_type.Name)
	return ds_initiative_type
}

// StagePreserveOrder puts ds_initiative_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Initiative_TypeOrder
// - update stage.DS_Initiative_TypeOrder accordingly
func (ds_initiative_type *DS_Initiative_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Initiative_Types, stage.DS_Initiative_Type_stagedOrder, stage.DS_Initiative_Type_orderStaged, &stage.DS_Initiative_TypeOrder, stage.DS_Initiative_Types_mapString, ds_initiative_type, order, ds_initiative_type.Name)
}

// Unstage removes ds_initiative_type off the model stage
func (ds_initiative_type *DS_Initiative_Type) Unstage(stage *Stage) *DS_Initiative_Type {
	__gong__unstage(stage.DS_Initiative_Types, stage.DS_Initiative_Types_mapString, ds_initiative_type, ds_initiative_type.Name)
	return ds_initiative_type
}

// UnstageVoid removes ds_initiative_type off the model stage
func (ds_initiative_type *DS_Initiative_Type) UnstageVoid(stage *Stage) {
	ds_initiative_type.Unstage(stage)
}

func (ds_initiative_type *DS_Initiative_Type) StageVoid(stage *Stage) {
	ds_initiative_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_initiative_type *DS_Initiative_Type) GetName() (res string) {
	return ds_initiative_type.Name
}

// for satisfaction of GongStruct interface
func (ds_initiative_type *DS_Initiative_Type) SetName(name string) {
	ds_initiative_type.Name = name
}

// Stage puts ds_otheraggregate_propertytype to the model stage
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) Stage(stage *Stage) *DS_OtherAggregate_PropertyType {
	__gong__stage(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyType_stagedOrder, stage.DS_OtherAggregate_PropertyType_orderStaged, &stage.DS_OtherAggregate_PropertyTypeOrder, stage.DS_OtherAggregate_PropertyTypes_mapString, ds_otheraggregate_propertytype, ds_otheraggregate_propertytype.Name)
	return ds_otheraggregate_propertytype
}

// StagePreserveOrder puts ds_otheraggregate_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_OtherAggregate_PropertyTypeOrder
// - update stage.DS_OtherAggregate_PropertyTypeOrder accordingly
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyType_stagedOrder, stage.DS_OtherAggregate_PropertyType_orderStaged, &stage.DS_OtherAggregate_PropertyTypeOrder, stage.DS_OtherAggregate_PropertyTypes_mapString, ds_otheraggregate_propertytype, order, ds_otheraggregate_propertytype.Name)
}

// Unstage removes ds_otheraggregate_propertytype off the model stage
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) Unstage(stage *Stage) *DS_OtherAggregate_PropertyType {
	__gong__unstage(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyTypes_mapString, ds_otheraggregate_propertytype, ds_otheraggregate_propertytype.Name)
	return ds_otheraggregate_propertytype
}

// UnstageVoid removes ds_otheraggregate_propertytype off the model stage
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) UnstageVoid(stage *Stage) {
	ds_otheraggregate_propertytype.Unstage(stage)
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) StageVoid(stage *Stage) {
	ds_otheraggregate_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GetName() (res string) {
	return ds_otheraggregate_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) SetName(name string) {
	ds_otheraggregate_propertytype.Name = name
}

// Stage puts ds_otheraggregate_type to the model stage
func (ds_otheraggregate_type *DS_OtherAggregate_Type) Stage(stage *Stage) *DS_OtherAggregate_Type {
	__gong__stage(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Type_stagedOrder, stage.DS_OtherAggregate_Type_orderStaged, &stage.DS_OtherAggregate_TypeOrder, stage.DS_OtherAggregate_Types_mapString, ds_otheraggregate_type, ds_otheraggregate_type.Name)
	return ds_otheraggregate_type
}

// StagePreserveOrder puts ds_otheraggregate_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_OtherAggregate_TypeOrder
// - update stage.DS_OtherAggregate_TypeOrder accordingly
func (ds_otheraggregate_type *DS_OtherAggregate_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Type_stagedOrder, stage.DS_OtherAggregate_Type_orderStaged, &stage.DS_OtherAggregate_TypeOrder, stage.DS_OtherAggregate_Types_mapString, ds_otheraggregate_type, order, ds_otheraggregate_type.Name)
}

// Unstage removes ds_otheraggregate_type off the model stage
func (ds_otheraggregate_type *DS_OtherAggregate_Type) Unstage(stage *Stage) *DS_OtherAggregate_Type {
	__gong__unstage(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Types_mapString, ds_otheraggregate_type, ds_otheraggregate_type.Name)
	return ds_otheraggregate_type
}

// UnstageVoid removes ds_otheraggregate_type off the model stage
func (ds_otheraggregate_type *DS_OtherAggregate_Type) UnstageVoid(stage *Stage) {
	ds_otheraggregate_type.Unstage(stage)
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) StageVoid(stage *Stage) {
	ds_otheraggregate_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_otheraggregate_type *DS_OtherAggregate_Type) GetName() (res string) {
	return ds_otheraggregate_type.Name
}

// for satisfaction of GongStruct interface
func (ds_otheraggregate_type *DS_OtherAggregate_Type) SetName(name string) {
	ds_otheraggregate_type.Name = name
}

// Stage puts ds_platform_propertytype to the model stage
func (ds_platform_propertytype *DS_Platform_PropertyType) Stage(stage *Stage) *DS_Platform_PropertyType {
	__gong__stage(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyType_stagedOrder, stage.DS_Platform_PropertyType_orderStaged, &stage.DS_Platform_PropertyTypeOrder, stage.DS_Platform_PropertyTypes_mapString, ds_platform_propertytype, ds_platform_propertytype.Name)
	return ds_platform_propertytype
}

// StagePreserveOrder puts ds_platform_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Platform_PropertyTypeOrder
// - update stage.DS_Platform_PropertyTypeOrder accordingly
func (ds_platform_propertytype *DS_Platform_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyType_stagedOrder, stage.DS_Platform_PropertyType_orderStaged, &stage.DS_Platform_PropertyTypeOrder, stage.DS_Platform_PropertyTypes_mapString, ds_platform_propertytype, order, ds_platform_propertytype.Name)
}

// Unstage removes ds_platform_propertytype off the model stage
func (ds_platform_propertytype *DS_Platform_PropertyType) Unstage(stage *Stage) *DS_Platform_PropertyType {
	__gong__unstage(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyTypes_mapString, ds_platform_propertytype, ds_platform_propertytype.Name)
	return ds_platform_propertytype
}

// UnstageVoid removes ds_platform_propertytype off the model stage
func (ds_platform_propertytype *DS_Platform_PropertyType) UnstageVoid(stage *Stage) {
	ds_platform_propertytype.Unstage(stage)
}

func (ds_platform_propertytype *DS_Platform_PropertyType) StageVoid(stage *Stage) {
	ds_platform_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_platform_propertytype *DS_Platform_PropertyType) GetName() (res string) {
	return ds_platform_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_platform_propertytype *DS_Platform_PropertyType) SetName(name string) {
	ds_platform_propertytype.Name = name
}

// Stage puts ds_platform_type to the model stage
func (ds_platform_type *DS_Platform_Type) Stage(stage *Stage) *DS_Platform_Type {
	__gong__stage(stage.DS_Platform_Types, stage.DS_Platform_Type_stagedOrder, stage.DS_Platform_Type_orderStaged, &stage.DS_Platform_TypeOrder, stage.DS_Platform_Types_mapString, ds_platform_type, ds_platform_type.Name)
	return ds_platform_type
}

// StagePreserveOrder puts ds_platform_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Platform_TypeOrder
// - update stage.DS_Platform_TypeOrder accordingly
func (ds_platform_type *DS_Platform_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Platform_Types, stage.DS_Platform_Type_stagedOrder, stage.DS_Platform_Type_orderStaged, &stage.DS_Platform_TypeOrder, stage.DS_Platform_Types_mapString, ds_platform_type, order, ds_platform_type.Name)
}

// Unstage removes ds_platform_type off the model stage
func (ds_platform_type *DS_Platform_Type) Unstage(stage *Stage) *DS_Platform_Type {
	__gong__unstage(stage.DS_Platform_Types, stage.DS_Platform_Types_mapString, ds_platform_type, ds_platform_type.Name)
	return ds_platform_type
}

// UnstageVoid removes ds_platform_type off the model stage
func (ds_platform_type *DS_Platform_Type) UnstageVoid(stage *Stage) {
	ds_platform_type.Unstage(stage)
}

func (ds_platform_type *DS_Platform_Type) StageVoid(stage *Stage) {
	ds_platform_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_platform_type *DS_Platform_Type) GetName() (res string) {
	return ds_platform_type.Name
}

// for satisfaction of GongStruct interface
func (ds_platform_type *DS_Platform_Type) SetName(name string) {
	ds_platform_type.Name = name
}

// Stage puts ds_productionseries_propertytype to the model stage
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) Stage(stage *Stage) *DS_ProductionSeries_PropertyType {
	__gong__stage(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyType_stagedOrder, stage.DS_ProductionSeries_PropertyType_orderStaged, &stage.DS_ProductionSeries_PropertyTypeOrder, stage.DS_ProductionSeries_PropertyTypes_mapString, ds_productionseries_propertytype, ds_productionseries_propertytype.Name)
	return ds_productionseries_propertytype
}

// StagePreserveOrder puts ds_productionseries_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_ProductionSeries_PropertyTypeOrder
// - update stage.DS_ProductionSeries_PropertyTypeOrder accordingly
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyType_stagedOrder, stage.DS_ProductionSeries_PropertyType_orderStaged, &stage.DS_ProductionSeries_PropertyTypeOrder, stage.DS_ProductionSeries_PropertyTypes_mapString, ds_productionseries_propertytype, order, ds_productionseries_propertytype.Name)
}

// Unstage removes ds_productionseries_propertytype off the model stage
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) Unstage(stage *Stage) *DS_ProductionSeries_PropertyType {
	__gong__unstage(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyTypes_mapString, ds_productionseries_propertytype, ds_productionseries_propertytype.Name)
	return ds_productionseries_propertytype
}

// UnstageVoid removes ds_productionseries_propertytype off the model stage
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) UnstageVoid(stage *Stage) {
	ds_productionseries_propertytype.Unstage(stage)
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) StageVoid(stage *Stage) {
	ds_productionseries_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GetName() (res string) {
	return ds_productionseries_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) SetName(name string) {
	ds_productionseries_propertytype.Name = name
}

// Stage puts ds_productionseries_type to the model stage
func (ds_productionseries_type *DS_ProductionSeries_Type) Stage(stage *Stage) *DS_ProductionSeries_Type {
	__gong__stage(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Type_stagedOrder, stage.DS_ProductionSeries_Type_orderStaged, &stage.DS_ProductionSeries_TypeOrder, stage.DS_ProductionSeries_Types_mapString, ds_productionseries_type, ds_productionseries_type.Name)
	return ds_productionseries_type
}

// StagePreserveOrder puts ds_productionseries_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_ProductionSeries_TypeOrder
// - update stage.DS_ProductionSeries_TypeOrder accordingly
func (ds_productionseries_type *DS_ProductionSeries_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Type_stagedOrder, stage.DS_ProductionSeries_Type_orderStaged, &stage.DS_ProductionSeries_TypeOrder, stage.DS_ProductionSeries_Types_mapString, ds_productionseries_type, order, ds_productionseries_type.Name)
}

// Unstage removes ds_productionseries_type off the model stage
func (ds_productionseries_type *DS_ProductionSeries_Type) Unstage(stage *Stage) *DS_ProductionSeries_Type {
	__gong__unstage(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Types_mapString, ds_productionseries_type, ds_productionseries_type.Name)
	return ds_productionseries_type
}

// UnstageVoid removes ds_productionseries_type off the model stage
func (ds_productionseries_type *DS_ProductionSeries_Type) UnstageVoid(stage *Stage) {
	ds_productionseries_type.Unstage(stage)
}

func (ds_productionseries_type *DS_ProductionSeries_Type) StageVoid(stage *Stage) {
	ds_productionseries_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_productionseries_type *DS_ProductionSeries_Type) GetName() (res string) {
	return ds_productionseries_type.Name
}

// for satisfaction of GongStruct interface
func (ds_productionseries_type *DS_ProductionSeries_Type) SetName(name string) {
	ds_productionseries_type.Name = name
}

// Stage puts ds_sensor_propertytype to the model stage
func (ds_sensor_propertytype *DS_Sensor_PropertyType) Stage(stage *Stage) *DS_Sensor_PropertyType {
	__gong__stage(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyType_stagedOrder, stage.DS_Sensor_PropertyType_orderStaged, &stage.DS_Sensor_PropertyTypeOrder, stage.DS_Sensor_PropertyTypes_mapString, ds_sensor_propertytype, ds_sensor_propertytype.Name)
	return ds_sensor_propertytype
}

// StagePreserveOrder puts ds_sensor_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Sensor_PropertyTypeOrder
// - update stage.DS_Sensor_PropertyTypeOrder accordingly
func (ds_sensor_propertytype *DS_Sensor_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyType_stagedOrder, stage.DS_Sensor_PropertyType_orderStaged, &stage.DS_Sensor_PropertyTypeOrder, stage.DS_Sensor_PropertyTypes_mapString, ds_sensor_propertytype, order, ds_sensor_propertytype.Name)
}

// Unstage removes ds_sensor_propertytype off the model stage
func (ds_sensor_propertytype *DS_Sensor_PropertyType) Unstage(stage *Stage) *DS_Sensor_PropertyType {
	__gong__unstage(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyTypes_mapString, ds_sensor_propertytype, ds_sensor_propertytype.Name)
	return ds_sensor_propertytype
}

// UnstageVoid removes ds_sensor_propertytype off the model stage
func (ds_sensor_propertytype *DS_Sensor_PropertyType) UnstageVoid(stage *Stage) {
	ds_sensor_propertytype.Unstage(stage)
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) StageVoid(stage *Stage) {
	ds_sensor_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_sensor_propertytype *DS_Sensor_PropertyType) GetName() (res string) {
	return ds_sensor_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_sensor_propertytype *DS_Sensor_PropertyType) SetName(name string) {
	ds_sensor_propertytype.Name = name
}

// Stage puts ds_sensor_type to the model stage
func (ds_sensor_type *DS_Sensor_Type) Stage(stage *Stage) *DS_Sensor_Type {
	__gong__stage(stage.DS_Sensor_Types, stage.DS_Sensor_Type_stagedOrder, stage.DS_Sensor_Type_orderStaged, &stage.DS_Sensor_TypeOrder, stage.DS_Sensor_Types_mapString, ds_sensor_type, ds_sensor_type.Name)
	return ds_sensor_type
}

// StagePreserveOrder puts ds_sensor_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Sensor_TypeOrder
// - update stage.DS_Sensor_TypeOrder accordingly
func (ds_sensor_type *DS_Sensor_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Sensor_Types, stage.DS_Sensor_Type_stagedOrder, stage.DS_Sensor_Type_orderStaged, &stage.DS_Sensor_TypeOrder, stage.DS_Sensor_Types_mapString, ds_sensor_type, order, ds_sensor_type.Name)
}

// Unstage removes ds_sensor_type off the model stage
func (ds_sensor_type *DS_Sensor_Type) Unstage(stage *Stage) *DS_Sensor_Type {
	__gong__unstage(stage.DS_Sensor_Types, stage.DS_Sensor_Types_mapString, ds_sensor_type, ds_sensor_type.Name)
	return ds_sensor_type
}

// UnstageVoid removes ds_sensor_type off the model stage
func (ds_sensor_type *DS_Sensor_Type) UnstageVoid(stage *Stage) {
	ds_sensor_type.Unstage(stage)
}

func (ds_sensor_type *DS_Sensor_Type) StageVoid(stage *Stage) {
	ds_sensor_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_sensor_type *DS_Sensor_Type) GetName() (res string) {
	return ds_sensor_type.Name
}

// for satisfaction of GongStruct interface
func (ds_sensor_type *DS_Sensor_Type) SetName(name string) {
	ds_sensor_type.Name = name
}

// Stage puts ds_series_propertytype to the model stage
func (ds_series_propertytype *DS_Series_PropertyType) Stage(stage *Stage) *DS_Series_PropertyType {
	__gong__stage(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyType_stagedOrder, stage.DS_Series_PropertyType_orderStaged, &stage.DS_Series_PropertyTypeOrder, stage.DS_Series_PropertyTypes_mapString, ds_series_propertytype, ds_series_propertytype.Name)
	return ds_series_propertytype
}

// StagePreserveOrder puts ds_series_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Series_PropertyTypeOrder
// - update stage.DS_Series_PropertyTypeOrder accordingly
func (ds_series_propertytype *DS_Series_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyType_stagedOrder, stage.DS_Series_PropertyType_orderStaged, &stage.DS_Series_PropertyTypeOrder, stage.DS_Series_PropertyTypes_mapString, ds_series_propertytype, order, ds_series_propertytype.Name)
}

// Unstage removes ds_series_propertytype off the model stage
func (ds_series_propertytype *DS_Series_PropertyType) Unstage(stage *Stage) *DS_Series_PropertyType {
	__gong__unstage(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyTypes_mapString, ds_series_propertytype, ds_series_propertytype.Name)
	return ds_series_propertytype
}

// UnstageVoid removes ds_series_propertytype off the model stage
func (ds_series_propertytype *DS_Series_PropertyType) UnstageVoid(stage *Stage) {
	ds_series_propertytype.Unstage(stage)
}

func (ds_series_propertytype *DS_Series_PropertyType) StageVoid(stage *Stage) {
	ds_series_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_series_propertytype *DS_Series_PropertyType) GetName() (res string) {
	return ds_series_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_series_propertytype *DS_Series_PropertyType) SetName(name string) {
	ds_series_propertytype.Name = name
}

// Stage puts ds_series_type to the model stage
func (ds_series_type *DS_Series_Type) Stage(stage *Stage) *DS_Series_Type {
	__gong__stage(stage.DS_Series_Types, stage.DS_Series_Type_stagedOrder, stage.DS_Series_Type_orderStaged, &stage.DS_Series_TypeOrder, stage.DS_Series_Types_mapString, ds_series_type, ds_series_type.Name)
	return ds_series_type
}

// StagePreserveOrder puts ds_series_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_Series_TypeOrder
// - update stage.DS_Series_TypeOrder accordingly
func (ds_series_type *DS_Series_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_Series_Types, stage.DS_Series_Type_stagedOrder, stage.DS_Series_Type_orderStaged, &stage.DS_Series_TypeOrder, stage.DS_Series_Types_mapString, ds_series_type, order, ds_series_type.Name)
}

// Unstage removes ds_series_type off the model stage
func (ds_series_type *DS_Series_Type) Unstage(stage *Stage) *DS_Series_Type {
	__gong__unstage(stage.DS_Series_Types, stage.DS_Series_Types_mapString, ds_series_type, ds_series_type.Name)
	return ds_series_type
}

// UnstageVoid removes ds_series_type off the model stage
func (ds_series_type *DS_Series_Type) UnstageVoid(stage *Stage) {
	ds_series_type.Unstage(stage)
}

func (ds_series_type *DS_Series_Type) StageVoid(stage *Stage) {
	ds_series_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_series_type *DS_Series_Type) GetName() (res string) {
	return ds_series_type.Name
}

// for satisfaction of GongStruct interface
func (ds_series_type *DS_Series_Type) SetName(name string) {
	ds_series_type.Name = name
}

// Stage puts ds_stereomate_propertytype to the model stage
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) Stage(stage *Stage) *DS_StereoMate_PropertyType {
	__gong__stage(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyType_stagedOrder, stage.DS_StereoMate_PropertyType_orderStaged, &stage.DS_StereoMate_PropertyTypeOrder, stage.DS_StereoMate_PropertyTypes_mapString, ds_stereomate_propertytype, ds_stereomate_propertytype.Name)
	return ds_stereomate_propertytype
}

// StagePreserveOrder puts ds_stereomate_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_StereoMate_PropertyTypeOrder
// - update stage.DS_StereoMate_PropertyTypeOrder accordingly
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyType_stagedOrder, stage.DS_StereoMate_PropertyType_orderStaged, &stage.DS_StereoMate_PropertyTypeOrder, stage.DS_StereoMate_PropertyTypes_mapString, ds_stereomate_propertytype, order, ds_stereomate_propertytype.Name)
}

// Unstage removes ds_stereomate_propertytype off the model stage
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) Unstage(stage *Stage) *DS_StereoMate_PropertyType {
	__gong__unstage(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyTypes_mapString, ds_stereomate_propertytype, ds_stereomate_propertytype.Name)
	return ds_stereomate_propertytype
}

// UnstageVoid removes ds_stereomate_propertytype off the model stage
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) UnstageVoid(stage *Stage) {
	ds_stereomate_propertytype.Unstage(stage)
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) StageVoid(stage *Stage) {
	ds_stereomate_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GetName() (res string) {
	return ds_stereomate_propertytype.Name
}

// for satisfaction of GongStruct interface
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) SetName(name string) {
	ds_stereomate_propertytype.Name = name
}

// Stage puts ds_stereomate_type to the model stage
func (ds_stereomate_type *DS_StereoMate_Type) Stage(stage *Stage) *DS_StereoMate_Type {
	__gong__stage(stage.DS_StereoMate_Types, stage.DS_StereoMate_Type_stagedOrder, stage.DS_StereoMate_Type_orderStaged, &stage.DS_StereoMate_TypeOrder, stage.DS_StereoMate_Types_mapString, ds_stereomate_type, ds_stereomate_type.Name)
	return ds_stereomate_type
}

// StagePreserveOrder puts ds_stereomate_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DS_StereoMate_TypeOrder
// - update stage.DS_StereoMate_TypeOrder accordingly
func (ds_stereomate_type *DS_StereoMate_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DS_StereoMate_Types, stage.DS_StereoMate_Type_stagedOrder, stage.DS_StereoMate_Type_orderStaged, &stage.DS_StereoMate_TypeOrder, stage.DS_StereoMate_Types_mapString, ds_stereomate_type, order, ds_stereomate_type.Name)
}

// Unstage removes ds_stereomate_type off the model stage
func (ds_stereomate_type *DS_StereoMate_Type) Unstage(stage *Stage) *DS_StereoMate_Type {
	__gong__unstage(stage.DS_StereoMate_Types, stage.DS_StereoMate_Types_mapString, ds_stereomate_type, ds_stereomate_type.Name)
	return ds_stereomate_type
}

// UnstageVoid removes ds_stereomate_type off the model stage
func (ds_stereomate_type *DS_StereoMate_Type) UnstageVoid(stage *Stage) {
	ds_stereomate_type.Unstage(stage)
}

func (ds_stereomate_type *DS_StereoMate_Type) StageVoid(stage *Stage) {
	ds_stereomate_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ds_stereomate_type *DS_StereoMate_Type) GetName() (res string) {
	return ds_stereomate_type.Name
}

// for satisfaction of GongStruct interface
func (ds_stereomate_type *DS_StereoMate_Type) SetName(name string) {
	ds_stereomate_type.Name = name
}

// Stage puts md_metadata_propertytype to the model stage
func (md_metadata_propertytype *MD_Metadata_PropertyType) Stage(stage *Stage) *MD_Metadata_PropertyType {
	__gong__stage(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyType_stagedOrder, stage.MD_Metadata_PropertyType_orderStaged, &stage.MD_Metadata_PropertyTypeOrder, stage.MD_Metadata_PropertyTypes_mapString, md_metadata_propertytype, md_metadata_propertytype.Name)
	return md_metadata_propertytype
}

// StagePreserveOrder puts md_metadata_propertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Metadata_PropertyTypeOrder
// - update stage.MD_Metadata_PropertyTypeOrder accordingly
func (md_metadata_propertytype *MD_Metadata_PropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyType_stagedOrder, stage.MD_Metadata_PropertyType_orderStaged, &stage.MD_Metadata_PropertyTypeOrder, stage.MD_Metadata_PropertyTypes_mapString, md_metadata_propertytype, order, md_metadata_propertytype.Name)
}

// Unstage removes md_metadata_propertytype off the model stage
func (md_metadata_propertytype *MD_Metadata_PropertyType) Unstage(stage *Stage) *MD_Metadata_PropertyType {
	__gong__unstage(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyTypes_mapString, md_metadata_propertytype, md_metadata_propertytype.Name)
	return md_metadata_propertytype
}

// UnstageVoid removes md_metadata_propertytype off the model stage
func (md_metadata_propertytype *MD_Metadata_PropertyType) UnstageVoid(stage *Stage) {
	md_metadata_propertytype.Unstage(stage)
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) StageVoid(stage *Stage) {
	md_metadata_propertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_metadata_propertytype *MD_Metadata_PropertyType) GetName() (res string) {
	return md_metadata_propertytype.Name
}

// for satisfaction of GongStruct interface
func (md_metadata_propertytype *MD_Metadata_PropertyType) SetName(name string) {
	md_metadata_propertytype.Name = name
}

// Stage puts md_metadata_type to the model stage
func (md_metadata_type *MD_Metadata_Type) Stage(stage *Stage) *MD_Metadata_Type {
	__gong__stage(stage.MD_Metadata_Types, stage.MD_Metadata_Type_stagedOrder, stage.MD_Metadata_Type_orderStaged, &stage.MD_Metadata_TypeOrder, stage.MD_Metadata_Types_mapString, md_metadata_type, md_metadata_type.Name)
	return md_metadata_type
}

// StagePreserveOrder puts md_metadata_type to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MD_Metadata_TypeOrder
// - update stage.MD_Metadata_TypeOrder accordingly
func (md_metadata_type *MD_Metadata_Type) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MD_Metadata_Types, stage.MD_Metadata_Type_stagedOrder, stage.MD_Metadata_Type_orderStaged, &stage.MD_Metadata_TypeOrder, stage.MD_Metadata_Types_mapString, md_metadata_type, order, md_metadata_type.Name)
}

// Unstage removes md_metadata_type off the model stage
func (md_metadata_type *MD_Metadata_Type) Unstage(stage *Stage) *MD_Metadata_Type {
	__gong__unstage(stage.MD_Metadata_Types, stage.MD_Metadata_Types_mapString, md_metadata_type, md_metadata_type.Name)
	return md_metadata_type
}

// UnstageVoid removes md_metadata_type off the model stage
func (md_metadata_type *MD_Metadata_Type) UnstageVoid(stage *Stage) {
	md_metadata_type.Unstage(stage)
}

func (md_metadata_type *MD_Metadata_Type) StageVoid(stage *Stage) {
	md_metadata_type.Stage(stage)
}

// for satisfaction of GongStruct interface
func (md_metadata_type *MD_Metadata_Type) GetName() (res string) {
	return md_metadata_type.Name
}

// for satisfaction of GongStruct interface
func (md_metadata_type *MD_Metadata_Type) SetName(name string) {
	md_metadata_type.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AbstractDS_Aggregate_Types, &stage.AbstractDS_Aggregate_Types_mapString, &stage.AbstractDS_Aggregate_Type_stagedOrder, &stage.AbstractDS_Aggregate_TypeOrder)

	__gong__resetStageType(&stage.DS_Aggregate_PropertyTypes, &stage.DS_Aggregate_PropertyTypes_mapString, &stage.DS_Aggregate_PropertyType_stagedOrder, &stage.DS_Aggregate_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_DataSet_PropertyTypes, &stage.DS_DataSet_PropertyTypes_mapString, &stage.DS_DataSet_PropertyType_stagedOrder, &stage.DS_DataSet_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_DataSet_Types, &stage.DS_DataSet_Types_mapString, &stage.DS_DataSet_Type_stagedOrder, &stage.DS_DataSet_TypeOrder)

	__gong__resetStageType(&stage.DS_Initiative_PropertyTypes, &stage.DS_Initiative_PropertyTypes_mapString, &stage.DS_Initiative_PropertyType_stagedOrder, &stage.DS_Initiative_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_Initiative_Types, &stage.DS_Initiative_Types_mapString, &stage.DS_Initiative_Type_stagedOrder, &stage.DS_Initiative_TypeOrder)

	__gong__resetStageType(&stage.DS_OtherAggregate_PropertyTypes, &stage.DS_OtherAggregate_PropertyTypes_mapString, &stage.DS_OtherAggregate_PropertyType_stagedOrder, &stage.DS_OtherAggregate_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_OtherAggregate_Types, &stage.DS_OtherAggregate_Types_mapString, &stage.DS_OtherAggregate_Type_stagedOrder, &stage.DS_OtherAggregate_TypeOrder)

	__gong__resetStageType(&stage.DS_Platform_PropertyTypes, &stage.DS_Platform_PropertyTypes_mapString, &stage.DS_Platform_PropertyType_stagedOrder, &stage.DS_Platform_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_Platform_Types, &stage.DS_Platform_Types_mapString, &stage.DS_Platform_Type_stagedOrder, &stage.DS_Platform_TypeOrder)

	__gong__resetStageType(&stage.DS_ProductionSeries_PropertyTypes, &stage.DS_ProductionSeries_PropertyTypes_mapString, &stage.DS_ProductionSeries_PropertyType_stagedOrder, &stage.DS_ProductionSeries_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_ProductionSeries_Types, &stage.DS_ProductionSeries_Types_mapString, &stage.DS_ProductionSeries_Type_stagedOrder, &stage.DS_ProductionSeries_TypeOrder)

	__gong__resetStageType(&stage.DS_Sensor_PropertyTypes, &stage.DS_Sensor_PropertyTypes_mapString, &stage.DS_Sensor_PropertyType_stagedOrder, &stage.DS_Sensor_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_Sensor_Types, &stage.DS_Sensor_Types_mapString, &stage.DS_Sensor_Type_stagedOrder, &stage.DS_Sensor_TypeOrder)

	__gong__resetStageType(&stage.DS_Series_PropertyTypes, &stage.DS_Series_PropertyTypes_mapString, &stage.DS_Series_PropertyType_stagedOrder, &stage.DS_Series_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_Series_Types, &stage.DS_Series_Types_mapString, &stage.DS_Series_Type_stagedOrder, &stage.DS_Series_TypeOrder)

	__gong__resetStageType(&stage.DS_StereoMate_PropertyTypes, &stage.DS_StereoMate_PropertyTypes_mapString, &stage.DS_StereoMate_PropertyType_stagedOrder, &stage.DS_StereoMate_PropertyTypeOrder)

	__gong__resetStageType(&stage.DS_StereoMate_Types, &stage.DS_StereoMate_Types_mapString, &stage.DS_StereoMate_Type_stagedOrder, &stage.DS_StereoMate_TypeOrder)

	__gong__resetStageType(&stage.MD_Metadata_PropertyTypes, &stage.MD_Metadata_PropertyTypes_mapString, &stage.MD_Metadata_PropertyType_stagedOrder, &stage.MD_Metadata_PropertyTypeOrder)

	__gong__resetStageType(&stage.MD_Metadata_Types, &stage.MD_Metadata_Types_mapString, &stage.MD_Metadata_Type_stagedOrder, &stage.MD_Metadata_TypeOrder)

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
func (AbstractDS_Aggregate_Type) GongGetAssociationName() any {
	return &AbstractDS_Aggregate_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_Aggregate_PropertyType) GongGetAssociationName() any {
	return &DS_Aggregate_PropertyType{
			DS_OtherAggregate: &DS_OtherAggregate_Type{Name: "DS_OtherAggregate"},
			DS_StereoMate: &DS_StereoMate_Type{Name: "DS_StereoMate"},
			DS_Series: &DS_Series_Type{Name: "DS_Series"},
			DS_Platform: &DS_Platform_Type{Name: "DS_Platform"},
			DS_Sensor: &DS_Sensor_Type{Name: "DS_Sensor"},
			DS_ProductionSeries: &DS_ProductionSeries_Type{Name: "DS_ProductionSeries"},
			DS_Initiative: &DS_Initiative_Type{Name: "DS_Initiative"},
	}
}

func (DS_DataSet_PropertyType) GongGetAssociationName() any {
	return &DS_DataSet_PropertyType{
			DS_DataSet: &DS_DataSet_Type{Name: "DS_DataSet"},
	}
}

func (DS_DataSet_Type) GongGetAssociationName() any {
	return &DS_DataSet_Type{
			Has: []*MD_Metadata_PropertyType{{Name: "Has"}},
			PartOf: []*DS_Aggregate_PropertyType{{Name: "PartOf"}},
	}
}

func (DS_Initiative_PropertyType) GongGetAssociationName() any {
	return &DS_Initiative_PropertyType{
			DS_Initiative: &DS_Initiative_Type{Name: "DS_Initiative"},
	}
}

func (DS_Initiative_Type) GongGetAssociationName() any {
	return &DS_Initiative_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_OtherAggregate_PropertyType) GongGetAssociationName() any {
	return &DS_OtherAggregate_PropertyType{
			DS_OtherAggregate: &DS_OtherAggregate_Type{Name: "DS_OtherAggregate"},
			DS_StereoMate: &DS_StereoMate_Type{Name: "DS_StereoMate"},
	}
}

func (DS_OtherAggregate_Type) GongGetAssociationName() any {
	return &DS_OtherAggregate_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_Platform_PropertyType) GongGetAssociationName() any {
	return &DS_Platform_PropertyType{
			DS_Platform: &DS_Platform_Type{Name: "DS_Platform"},
	}
}

func (DS_Platform_Type) GongGetAssociationName() any {
	return &DS_Platform_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_ProductionSeries_PropertyType) GongGetAssociationName() any {
	return &DS_ProductionSeries_PropertyType{
			DS_ProductionSeries: &DS_ProductionSeries_Type{Name: "DS_ProductionSeries"},
	}
}

func (DS_ProductionSeries_Type) GongGetAssociationName() any {
	return &DS_ProductionSeries_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_Sensor_PropertyType) GongGetAssociationName() any {
	return &DS_Sensor_PropertyType{
			DS_Sensor: &DS_Sensor_Type{Name: "DS_Sensor"},
	}
}

func (DS_Sensor_Type) GongGetAssociationName() any {
	return &DS_Sensor_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_Series_PropertyType) GongGetAssociationName() any {
	return &DS_Series_PropertyType{
			DS_Series: &DS_Series_Type{Name: "DS_Series"},
			DS_Platform: &DS_Platform_Type{Name: "DS_Platform"},
			DS_Sensor: &DS_Sensor_Type{Name: "DS_Sensor"},
			DS_ProductionSeries: &DS_ProductionSeries_Type{Name: "DS_ProductionSeries"},
	}
}

func (DS_Series_Type) GongGetAssociationName() any {
	return &DS_Series_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (DS_StereoMate_PropertyType) GongGetAssociationName() any {
	return &DS_StereoMate_PropertyType{
			DS_StereoMate: &DS_StereoMate_Type{Name: "DS_StereoMate"},
	}
}

func (DS_StereoMate_Type) GongGetAssociationName() any {
	return &DS_StereoMate_Type{
			ComposedOf: []*DS_DataSet_PropertyType{{Name: "ComposedOf"}},
			SeriesMetadata: []*MD_Metadata_PropertyType{{Name: "SeriesMetadata"}},
			Subset: []*DS_Aggregate_PropertyType{{Name: "Subset"}},
			Superset: []*DS_Aggregate_PropertyType{{Name: "Superset"}},
	}
}

func (MD_Metadata_PropertyType) GongGetAssociationName() any {
	return &MD_Metadata_PropertyType{
			MD_Metadata: &MD_Metadata_Type{Name: "MD_Metadata"},
	}
}

func (MD_Metadata_Type) GongGetAssociationName() any {
	return &MD_Metadata_Type{
			Series: []*DS_Aggregate_PropertyType{{Name: "Series"}},
			Describes: []*DS_DataSet_PropertyType{{Name: "Describes"}},
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
	// reverse maps of direct associations of AbstractDS_Aggregate_Type
	case AbstractDS_Aggregate_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Aggregate_PropertyType
	case DS_Aggregate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_OtherAggregate":
			res := make(map[*DS_OtherAggregate_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_OtherAggregate != nil {
					ds_otheraggregate_type_ := ds_aggregate_propertytype.DS_OtherAggregate
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_otheraggregate_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_otheraggregate_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_otheraggregate_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_StereoMate":
			res := make(map[*DS_StereoMate_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_StereoMate != nil {
					ds_stereomate_type_ := ds_aggregate_propertytype.DS_StereoMate
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_stereomate_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_stereomate_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_stereomate_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Series":
			res := make(map[*DS_Series_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_Series != nil {
					ds_series_type_ := ds_aggregate_propertytype.DS_Series
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_series_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_series_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_series_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Platform":
			res := make(map[*DS_Platform_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_Platform != nil {
					ds_platform_type_ := ds_aggregate_propertytype.DS_Platform
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_platform_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_platform_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_platform_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Sensor":
			res := make(map[*DS_Sensor_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_Sensor != nil {
					ds_sensor_type_ := ds_aggregate_propertytype.DS_Sensor
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_sensor_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_sensor_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_sensor_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_ProductionSeries":
			res := make(map[*DS_ProductionSeries_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_ProductionSeries != nil {
					ds_productionseries_type_ := ds_aggregate_propertytype.DS_ProductionSeries
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_productionseries_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_productionseries_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_productionseries_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Initiative":
			res := make(map[*DS_Initiative_Type][]*DS_Aggregate_PropertyType)
			for ds_aggregate_propertytype := range stage.DS_Aggregate_PropertyTypes {
				if ds_aggregate_propertytype.DS_Initiative != nil {
					ds_initiative_type_ := ds_aggregate_propertytype.DS_Initiative
					var ds_aggregate_propertytypes []*DS_Aggregate_PropertyType
					_, ok := res[ds_initiative_type_]
					if ok {
						ds_aggregate_propertytypes = res[ds_initiative_type_]
					} else {
						ds_aggregate_propertytypes = make([]*DS_Aggregate_PropertyType, 0)
					}
					ds_aggregate_propertytypes = append(ds_aggregate_propertytypes, ds_aggregate_propertytype)
					res[ds_initiative_type_] = ds_aggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_DataSet_PropertyType
	case DS_DataSet_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_DataSet":
			res := make(map[*DS_DataSet_Type][]*DS_DataSet_PropertyType)
			for ds_dataset_propertytype := range stage.DS_DataSet_PropertyTypes {
				if ds_dataset_propertytype.DS_DataSet != nil {
					ds_dataset_type_ := ds_dataset_propertytype.DS_DataSet
					var ds_dataset_propertytypes []*DS_DataSet_PropertyType
					_, ok := res[ds_dataset_type_]
					if ok {
						ds_dataset_propertytypes = res[ds_dataset_type_]
					} else {
						ds_dataset_propertytypes = make([]*DS_DataSet_PropertyType, 0)
					}
					ds_dataset_propertytypes = append(ds_dataset_propertytypes, ds_dataset_propertytype)
					res[ds_dataset_type_] = ds_dataset_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_DataSet_Type
	case DS_DataSet_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Initiative_PropertyType
	case DS_Initiative_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_Initiative":
			res := make(map[*DS_Initiative_Type][]*DS_Initiative_PropertyType)
			for ds_initiative_propertytype := range stage.DS_Initiative_PropertyTypes {
				if ds_initiative_propertytype.DS_Initiative != nil {
					ds_initiative_type_ := ds_initiative_propertytype.DS_Initiative
					var ds_initiative_propertytypes []*DS_Initiative_PropertyType
					_, ok := res[ds_initiative_type_]
					if ok {
						ds_initiative_propertytypes = res[ds_initiative_type_]
					} else {
						ds_initiative_propertytypes = make([]*DS_Initiative_PropertyType, 0)
					}
					ds_initiative_propertytypes = append(ds_initiative_propertytypes, ds_initiative_propertytype)
					res[ds_initiative_type_] = ds_initiative_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Initiative_Type
	case DS_Initiative_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_OtherAggregate_PropertyType
	case DS_OtherAggregate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_OtherAggregate":
			res := make(map[*DS_OtherAggregate_Type][]*DS_OtherAggregate_PropertyType)
			for ds_otheraggregate_propertytype := range stage.DS_OtherAggregate_PropertyTypes {
				if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
					ds_otheraggregate_type_ := ds_otheraggregate_propertytype.DS_OtherAggregate
					var ds_otheraggregate_propertytypes []*DS_OtherAggregate_PropertyType
					_, ok := res[ds_otheraggregate_type_]
					if ok {
						ds_otheraggregate_propertytypes = res[ds_otheraggregate_type_]
					} else {
						ds_otheraggregate_propertytypes = make([]*DS_OtherAggregate_PropertyType, 0)
					}
					ds_otheraggregate_propertytypes = append(ds_otheraggregate_propertytypes, ds_otheraggregate_propertytype)
					res[ds_otheraggregate_type_] = ds_otheraggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_StereoMate":
			res := make(map[*DS_StereoMate_Type][]*DS_OtherAggregate_PropertyType)
			for ds_otheraggregate_propertytype := range stage.DS_OtherAggregate_PropertyTypes {
				if ds_otheraggregate_propertytype.DS_StereoMate != nil {
					ds_stereomate_type_ := ds_otheraggregate_propertytype.DS_StereoMate
					var ds_otheraggregate_propertytypes []*DS_OtherAggregate_PropertyType
					_, ok := res[ds_stereomate_type_]
					if ok {
						ds_otheraggregate_propertytypes = res[ds_stereomate_type_]
					} else {
						ds_otheraggregate_propertytypes = make([]*DS_OtherAggregate_PropertyType, 0)
					}
					ds_otheraggregate_propertytypes = append(ds_otheraggregate_propertytypes, ds_otheraggregate_propertytype)
					res[ds_stereomate_type_] = ds_otheraggregate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_OtherAggregate_Type
	case DS_OtherAggregate_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Platform_PropertyType
	case DS_Platform_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_Platform":
			res := make(map[*DS_Platform_Type][]*DS_Platform_PropertyType)
			for ds_platform_propertytype := range stage.DS_Platform_PropertyTypes {
				if ds_platform_propertytype.DS_Platform != nil {
					ds_platform_type_ := ds_platform_propertytype.DS_Platform
					var ds_platform_propertytypes []*DS_Platform_PropertyType
					_, ok := res[ds_platform_type_]
					if ok {
						ds_platform_propertytypes = res[ds_platform_type_]
					} else {
						ds_platform_propertytypes = make([]*DS_Platform_PropertyType, 0)
					}
					ds_platform_propertytypes = append(ds_platform_propertytypes, ds_platform_propertytype)
					res[ds_platform_type_] = ds_platform_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Platform_Type
	case DS_Platform_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_ProductionSeries_PropertyType
	case DS_ProductionSeries_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_ProductionSeries":
			res := make(map[*DS_ProductionSeries_Type][]*DS_ProductionSeries_PropertyType)
			for ds_productionseries_propertytype := range stage.DS_ProductionSeries_PropertyTypes {
				if ds_productionseries_propertytype.DS_ProductionSeries != nil {
					ds_productionseries_type_ := ds_productionseries_propertytype.DS_ProductionSeries
					var ds_productionseries_propertytypes []*DS_ProductionSeries_PropertyType
					_, ok := res[ds_productionseries_type_]
					if ok {
						ds_productionseries_propertytypes = res[ds_productionseries_type_]
					} else {
						ds_productionseries_propertytypes = make([]*DS_ProductionSeries_PropertyType, 0)
					}
					ds_productionseries_propertytypes = append(ds_productionseries_propertytypes, ds_productionseries_propertytype)
					res[ds_productionseries_type_] = ds_productionseries_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_ProductionSeries_Type
	case DS_ProductionSeries_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Sensor_PropertyType
	case DS_Sensor_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_Sensor":
			res := make(map[*DS_Sensor_Type][]*DS_Sensor_PropertyType)
			for ds_sensor_propertytype := range stage.DS_Sensor_PropertyTypes {
				if ds_sensor_propertytype.DS_Sensor != nil {
					ds_sensor_type_ := ds_sensor_propertytype.DS_Sensor
					var ds_sensor_propertytypes []*DS_Sensor_PropertyType
					_, ok := res[ds_sensor_type_]
					if ok {
						ds_sensor_propertytypes = res[ds_sensor_type_]
					} else {
						ds_sensor_propertytypes = make([]*DS_Sensor_PropertyType, 0)
					}
					ds_sensor_propertytypes = append(ds_sensor_propertytypes, ds_sensor_propertytype)
					res[ds_sensor_type_] = ds_sensor_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Sensor_Type
	case DS_Sensor_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Series_PropertyType
	case DS_Series_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_Series":
			res := make(map[*DS_Series_Type][]*DS_Series_PropertyType)
			for ds_series_propertytype := range stage.DS_Series_PropertyTypes {
				if ds_series_propertytype.DS_Series != nil {
					ds_series_type_ := ds_series_propertytype.DS_Series
					var ds_series_propertytypes []*DS_Series_PropertyType
					_, ok := res[ds_series_type_]
					if ok {
						ds_series_propertytypes = res[ds_series_type_]
					} else {
						ds_series_propertytypes = make([]*DS_Series_PropertyType, 0)
					}
					ds_series_propertytypes = append(ds_series_propertytypes, ds_series_propertytype)
					res[ds_series_type_] = ds_series_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Platform":
			res := make(map[*DS_Platform_Type][]*DS_Series_PropertyType)
			for ds_series_propertytype := range stage.DS_Series_PropertyTypes {
				if ds_series_propertytype.DS_Platform != nil {
					ds_platform_type_ := ds_series_propertytype.DS_Platform
					var ds_series_propertytypes []*DS_Series_PropertyType
					_, ok := res[ds_platform_type_]
					if ok {
						ds_series_propertytypes = res[ds_platform_type_]
					} else {
						ds_series_propertytypes = make([]*DS_Series_PropertyType, 0)
					}
					ds_series_propertytypes = append(ds_series_propertytypes, ds_series_propertytype)
					res[ds_platform_type_] = ds_series_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_Sensor":
			res := make(map[*DS_Sensor_Type][]*DS_Series_PropertyType)
			for ds_series_propertytype := range stage.DS_Series_PropertyTypes {
				if ds_series_propertytype.DS_Sensor != nil {
					ds_sensor_type_ := ds_series_propertytype.DS_Sensor
					var ds_series_propertytypes []*DS_Series_PropertyType
					_, ok := res[ds_sensor_type_]
					if ok {
						ds_series_propertytypes = res[ds_sensor_type_]
					} else {
						ds_series_propertytypes = make([]*DS_Series_PropertyType, 0)
					}
					ds_series_propertytypes = append(ds_series_propertytypes, ds_series_propertytype)
					res[ds_sensor_type_] = ds_series_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DS_ProductionSeries":
			res := make(map[*DS_ProductionSeries_Type][]*DS_Series_PropertyType)
			for ds_series_propertytype := range stage.DS_Series_PropertyTypes {
				if ds_series_propertytype.DS_ProductionSeries != nil {
					ds_productionseries_type_ := ds_series_propertytype.DS_ProductionSeries
					var ds_series_propertytypes []*DS_Series_PropertyType
					_, ok := res[ds_productionseries_type_]
					if ok {
						ds_series_propertytypes = res[ds_productionseries_type_]
					} else {
						ds_series_propertytypes = make([]*DS_Series_PropertyType, 0)
					}
					ds_series_propertytypes = append(ds_series_propertytypes, ds_series_propertytype)
					res[ds_productionseries_type_] = ds_series_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Series_Type
	case DS_Series_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_StereoMate_PropertyType
	case DS_StereoMate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "DS_StereoMate":
			res := make(map[*DS_StereoMate_Type][]*DS_StereoMate_PropertyType)
			for ds_stereomate_propertytype := range stage.DS_StereoMate_PropertyTypes {
				if ds_stereomate_propertytype.DS_StereoMate != nil {
					ds_stereomate_type_ := ds_stereomate_propertytype.DS_StereoMate
					var ds_stereomate_propertytypes []*DS_StereoMate_PropertyType
					_, ok := res[ds_stereomate_type_]
					if ok {
						ds_stereomate_propertytypes = res[ds_stereomate_type_]
					} else {
						ds_stereomate_propertytypes = make([]*DS_StereoMate_PropertyType, 0)
					}
					ds_stereomate_propertytypes = append(ds_stereomate_propertytypes, ds_stereomate_propertytype)
					res[ds_stereomate_type_] = ds_stereomate_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_StereoMate_Type
	case DS_StereoMate_Type:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Metadata_PropertyType
	case MD_Metadata_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "MD_Metadata":
			res := make(map[*MD_Metadata_Type][]*MD_Metadata_PropertyType)
			for md_metadata_propertytype := range stage.MD_Metadata_PropertyTypes {
				if md_metadata_propertytype.MD_Metadata != nil {
					md_metadata_type_ := md_metadata_propertytype.MD_Metadata
					var md_metadata_propertytypes []*MD_Metadata_PropertyType
					_, ok := res[md_metadata_type_]
					if ok {
						md_metadata_propertytypes = res[md_metadata_type_]
					} else {
						md_metadata_propertytypes = make([]*MD_Metadata_PropertyType, 0)
					}
					md_metadata_propertytypes = append(md_metadata_propertytypes, md_metadata_propertytype)
					res[md_metadata_type_] = md_metadata_propertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Metadata_Type
	case MD_Metadata_Type:
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
	// reverse maps of direct associations of AbstractDS_Aggregate_Type
	case AbstractDS_Aggregate_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*AbstractDS_Aggregate_Type)
			for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
				for _, ds_dataset_propertytype_ := range abstractds_aggregate_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], abstractds_aggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*AbstractDS_Aggregate_Type)
			for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
				for _, md_metadata_propertytype_ := range abstractds_aggregate_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], abstractds_aggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*AbstractDS_Aggregate_Type)
			for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
				for _, ds_aggregate_propertytype_ := range abstractds_aggregate_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], abstractds_aggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*AbstractDS_Aggregate_Type)
			for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
				for _, ds_aggregate_propertytype_ := range abstractds_aggregate_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], abstractds_aggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Aggregate_PropertyType
	case DS_Aggregate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_DataSet_PropertyType
	case DS_DataSet_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_DataSet_Type
	case DS_DataSet_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "Has":
			res := make(map[*MD_Metadata_PropertyType][]*DS_DataSet_Type)
			for ds_dataset_type := range stage.DS_DataSet_Types {
				for _, md_metadata_propertytype_ := range ds_dataset_type.Has {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_dataset_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartOf":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_DataSet_Type)
			for ds_dataset_type := range stage.DS_DataSet_Types {
				for _, ds_aggregate_propertytype_ := range ds_dataset_type.PartOf {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_dataset_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Initiative_PropertyType
	case DS_Initiative_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Initiative_Type
	case DS_Initiative_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_Initiative_Type)
			for ds_initiative_type := range stage.DS_Initiative_Types {
				for _, ds_dataset_propertytype_ := range ds_initiative_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_initiative_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_Initiative_Type)
			for ds_initiative_type := range stage.DS_Initiative_Types {
				for _, md_metadata_propertytype_ := range ds_initiative_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_initiative_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Initiative_Type)
			for ds_initiative_type := range stage.DS_Initiative_Types {
				for _, ds_aggregate_propertytype_ := range ds_initiative_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_initiative_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Initiative_Type)
			for ds_initiative_type := range stage.DS_Initiative_Types {
				for _, ds_aggregate_propertytype_ := range ds_initiative_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_initiative_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_OtherAggregate_PropertyType
	case DS_OtherAggregate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_OtherAggregate_Type
	case DS_OtherAggregate_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_OtherAggregate_Type)
			for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
				for _, ds_dataset_propertytype_ := range ds_otheraggregate_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_otheraggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_OtherAggregate_Type)
			for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
				for _, md_metadata_propertytype_ := range ds_otheraggregate_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_otheraggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_OtherAggregate_Type)
			for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
				for _, ds_aggregate_propertytype_ := range ds_otheraggregate_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_otheraggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_OtherAggregate_Type)
			for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
				for _, ds_aggregate_propertytype_ := range ds_otheraggregate_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_otheraggregate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Platform_PropertyType
	case DS_Platform_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Platform_Type
	case DS_Platform_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_Platform_Type)
			for ds_platform_type := range stage.DS_Platform_Types {
				for _, ds_dataset_propertytype_ := range ds_platform_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_platform_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_Platform_Type)
			for ds_platform_type := range stage.DS_Platform_Types {
				for _, md_metadata_propertytype_ := range ds_platform_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_platform_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Platform_Type)
			for ds_platform_type := range stage.DS_Platform_Types {
				for _, ds_aggregate_propertytype_ := range ds_platform_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_platform_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Platform_Type)
			for ds_platform_type := range stage.DS_Platform_Types {
				for _, ds_aggregate_propertytype_ := range ds_platform_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_platform_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_ProductionSeries_PropertyType
	case DS_ProductionSeries_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_ProductionSeries_Type
	case DS_ProductionSeries_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_ProductionSeries_Type)
			for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
				for _, ds_dataset_propertytype_ := range ds_productionseries_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_productionseries_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_ProductionSeries_Type)
			for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
				for _, md_metadata_propertytype_ := range ds_productionseries_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_productionseries_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_ProductionSeries_Type)
			for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
				for _, ds_aggregate_propertytype_ := range ds_productionseries_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_productionseries_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_ProductionSeries_Type)
			for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
				for _, ds_aggregate_propertytype_ := range ds_productionseries_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_productionseries_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Sensor_PropertyType
	case DS_Sensor_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Sensor_Type
	case DS_Sensor_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_Sensor_Type)
			for ds_sensor_type := range stage.DS_Sensor_Types {
				for _, ds_dataset_propertytype_ := range ds_sensor_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_sensor_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_Sensor_Type)
			for ds_sensor_type := range stage.DS_Sensor_Types {
				for _, md_metadata_propertytype_ := range ds_sensor_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_sensor_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Sensor_Type)
			for ds_sensor_type := range stage.DS_Sensor_Types {
				for _, ds_aggregate_propertytype_ := range ds_sensor_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_sensor_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Sensor_Type)
			for ds_sensor_type := range stage.DS_Sensor_Types {
				for _, ds_aggregate_propertytype_ := range ds_sensor_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_sensor_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_Series_PropertyType
	case DS_Series_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_Series_Type
	case DS_Series_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_Series_Type)
			for ds_series_type := range stage.DS_Series_Types {
				for _, ds_dataset_propertytype_ := range ds_series_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_series_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_Series_Type)
			for ds_series_type := range stage.DS_Series_Types {
				for _, md_metadata_propertytype_ := range ds_series_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_series_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Series_Type)
			for ds_series_type := range stage.DS_Series_Types {
				for _, ds_aggregate_propertytype_ := range ds_series_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_series_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_Series_Type)
			for ds_series_type := range stage.DS_Series_Types {
				for _, ds_aggregate_propertytype_ := range ds_series_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_series_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DS_StereoMate_PropertyType
	case DS_StereoMate_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DS_StereoMate_Type
	case DS_StereoMate_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "ComposedOf":
			res := make(map[*DS_DataSet_PropertyType][]*DS_StereoMate_Type)
			for ds_stereomate_type := range stage.DS_StereoMate_Types {
				for _, ds_dataset_propertytype_ := range ds_stereomate_type.ComposedOf {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], ds_stereomate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SeriesMetadata":
			res := make(map[*MD_Metadata_PropertyType][]*DS_StereoMate_Type)
			for ds_stereomate_type := range stage.DS_StereoMate_Types {
				for _, md_metadata_propertytype_ := range ds_stereomate_type.SeriesMetadata {
					res[md_metadata_propertytype_] = append(res[md_metadata_propertytype_], ds_stereomate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Subset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_StereoMate_Type)
			for ds_stereomate_type := range stage.DS_StereoMate_Types {
				for _, ds_aggregate_propertytype_ := range ds_stereomate_type.Subset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_stereomate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Superset":
			res := make(map[*DS_Aggregate_PropertyType][]*DS_StereoMate_Type)
			for ds_stereomate_type := range stage.DS_StereoMate_Types {
				for _, ds_aggregate_propertytype_ := range ds_stereomate_type.Superset {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], ds_stereomate_type)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MD_Metadata_PropertyType
	case MD_Metadata_PropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MD_Metadata_Type
	case MD_Metadata_Type:
		switch fieldname {
		// insertion point for per direct association field
		case "Series":
			res := make(map[*DS_Aggregate_PropertyType][]*MD_Metadata_Type)
			for md_metadata_type := range stage.MD_Metadata_Types {
				for _, ds_aggregate_propertytype_ := range md_metadata_type.Series {
					res[ds_aggregate_propertytype_] = append(res[ds_aggregate_propertytype_], md_metadata_type)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Describes":
			res := make(map[*DS_DataSet_PropertyType][]*MD_Metadata_Type)
			for md_metadata_type := range stage.MD_Metadata_Types {
				for _, ds_dataset_propertytype_ := range md_metadata_type.Describes {
					res[ds_dataset_propertytype_] = append(res[ds_dataset_propertytype_], md_metadata_type)
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
func (*AbstractDS_Aggregate_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Aggregate_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AbstractDS_Aggregate_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "AbstractDS_Aggregate_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_DataSet_Type",
			Fieldname: "PartOf",
		},
		{
			GongstructName: "DS_Initiative_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_Initiative_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_OtherAggregate_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_OtherAggregate_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_Platform_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_Platform_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_ProductionSeries_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_ProductionSeries_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_Sensor_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_Sensor_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_Series_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_Series_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "DS_StereoMate_Type",
			Fieldname: "Subset",
		},
		{
			GongstructName: "DS_StereoMate_Type",
			Fieldname: "Superset",
		},
		{
			GongstructName: "MD_Metadata_Type",
			Fieldname: "Series",
		},
	}
}

func (*DS_DataSet_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AbstractDS_Aggregate_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_Initiative_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_OtherAggregate_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_Platform_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_ProductionSeries_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_Sensor_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_Series_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "DS_StereoMate_Type",
			Fieldname: "ComposedOf",
		},
		{
			GongstructName: "MD_Metadata_Type",
			Fieldname: "Describes",
		},
	}
}

func (*DS_DataSet_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Initiative_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Initiative_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_OtherAggregate_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_OtherAggregate_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Platform_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Platform_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_ProductionSeries_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_ProductionSeries_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Sensor_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Sensor_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Series_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_Series_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_StereoMate_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DS_StereoMate_Type) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MD_Metadata_PropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AbstractDS_Aggregate_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_DataSet_Type",
			Fieldname: "Has",
		},
		{
			GongstructName: "DS_Initiative_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_OtherAggregate_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_Platform_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_ProductionSeries_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_Sensor_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_Series_Type",
			Fieldname: "SeriesMetadata",
		},
		{
			GongstructName: "DS_StereoMate_Type",
			Fieldname: "SeriesMetadata",
		},
	}
}

func (*MD_Metadata_Type) GongGetReverseFields() []GongReverseField {
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
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_OtherAggregate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_OtherAggregate_Type",
		},
		{
			Name:                 "DS_StereoMate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_StereoMate_Type",
		},
		{
			Name:                 "DS_Series",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Series_Type",
		},
		{
			Name:                 "DS_Platform",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Platform_Type",
		},
		{
			Name:                 "DS_Sensor",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Sensor_Type",
		},
		{
			Name:                 "DS_ProductionSeries",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_ProductionSeries_Type",
		},
		{
			Name:                 "DS_Initiative",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Initiative_Type",
		},
	}
	return
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_DataSet",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_DataSet_Type",
		},
	}
	return
}

func (ds_dataset_type *DS_DataSet_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Has",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "PartOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_Initiative",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Initiative_Type",
		},
	}
	return
}

func (ds_initiative_type *DS_Initiative_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_OtherAggregate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_OtherAggregate_Type",
		},
		{
			Name:                 "DS_StereoMate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_StereoMate_Type",
		},
	}
	return
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_Platform",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Platform_Type",
		},
	}
	return
}

func (ds_platform_type *DS_Platform_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_ProductionSeries",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_ProductionSeries_Type",
		},
	}
	return
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_Sensor",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Sensor_Type",
		},
	}
	return
}

func (ds_sensor_type *DS_Sensor_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_series_propertytype *DS_Series_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_Series",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Series_Type",
		},
		{
			Name:                 "DS_Platform",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Platform_Type",
		},
		{
			Name:                 "DS_Sensor",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_Sensor_Type",
		},
		{
			Name:                 "DS_ProductionSeries",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_ProductionSeries_Type",
		},
	}
	return
}

func (ds_series_type *DS_Series_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DS_StereoMate",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DS_StereoMate_Type",
		},
	}
	return
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ComposedOf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
		},
		{
			Name:                 "SeriesMetadata",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MD_Metadata_PropertyType",
		},
		{
			Name:                 "Subset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Superset",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
	}
	return
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "MD_Metadata",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MD_Metadata_Type",
		},
	}
	return
}

func (md_metadata_type *MD_Metadata_Type) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Series",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_Aggregate_PropertyType",
		},
		{
			Name:                 "Describes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DS_DataSet_PropertyType",
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
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractds_aggregate_type.Name
	case "Id":
		res.valueString = abstractds_aggregate_type.Id
	case "Uuid":
		res.valueString = abstractds_aggregate_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractds_aggregate_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractds_aggregate_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractds_aggregate_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range abstractds_aggregate_type.Superset {
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

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_aggregate_propertytype.Name
	case "NilReason":
		res.valueString = ds_aggregate_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_aggregate_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_aggregate_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_aggregate_propertytype.Href
	case "Role":
		res.valueString = ds_aggregate_propertytype.Role
	case "Arcrole":
		res.valueString = ds_aggregate_propertytype.Arcrole
	case "Title":
		res.valueString = ds_aggregate_propertytype.Title
	case "Show":
		res.valueString = ds_aggregate_propertytype.Show
	case "Actuate":
		res.valueString = ds_aggregate_propertytype.Actuate
	case "DS_OtherAggregate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_OtherAggregate != nil {
			res.valueString = ds_aggregate_propertytype.DS_OtherAggregate.Name
			res.ids = ds_aggregate_propertytype.DS_OtherAggregate.GongGetUUID(stage)
		}
	case "DS_StereoMate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_StereoMate != nil {
			res.valueString = ds_aggregate_propertytype.DS_StereoMate.Name
			res.ids = ds_aggregate_propertytype.DS_StereoMate.GongGetUUID(stage)
		}
	case "DS_Series":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_Series != nil {
			res.valueString = ds_aggregate_propertytype.DS_Series.Name
			res.ids = ds_aggregate_propertytype.DS_Series.GongGetUUID(stage)
		}
	case "DS_Platform":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_Platform != nil {
			res.valueString = ds_aggregate_propertytype.DS_Platform.Name
			res.ids = ds_aggregate_propertytype.DS_Platform.GongGetUUID(stage)
		}
	case "DS_Sensor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_Sensor != nil {
			res.valueString = ds_aggregate_propertytype.DS_Sensor.Name
			res.ids = ds_aggregate_propertytype.DS_Sensor.GongGetUUID(stage)
		}
	case "DS_ProductionSeries":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_ProductionSeries != nil {
			res.valueString = ds_aggregate_propertytype.DS_ProductionSeries.Name
			res.ids = ds_aggregate_propertytype.DS_ProductionSeries.GongGetUUID(stage)
		}
	case "DS_Initiative":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_aggregate_propertytype.DS_Initiative != nil {
			res.valueString = ds_aggregate_propertytype.DS_Initiative.Name
			res.ids = ds_aggregate_propertytype.DS_Initiative.GongGetUUID(stage)
		}
	}
	return
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_dataset_propertytype.Name
	case "NilReason":
		res.valueString = ds_dataset_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_dataset_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_dataset_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_dataset_propertytype.Href
	case "Role":
		res.valueString = ds_dataset_propertytype.Role
	case "Arcrole":
		res.valueString = ds_dataset_propertytype.Arcrole
	case "Title":
		res.valueString = ds_dataset_propertytype.Title
	case "Show":
		res.valueString = ds_dataset_propertytype.Show
	case "Actuate":
		res.valueString = ds_dataset_propertytype.Actuate
	case "DS_DataSet":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_dataset_propertytype.DS_DataSet != nil {
			res.valueString = ds_dataset_propertytype.DS_DataSet.Name
			res.ids = ds_dataset_propertytype.DS_DataSet.GongGetUUID(stage)
		}
	}
	return
}

func (ds_dataset_type *DS_DataSet_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_dataset_type.Name
	case "Id":
		res.valueString = ds_dataset_type.Id
	case "Uuid":
		res.valueString = ds_dataset_type.Uuid
	case "Has":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_dataset_type.Has {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_dataset_type.PartOf {
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

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_initiative_propertytype.Name
	case "NilReason":
		res.valueString = ds_initiative_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_initiative_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_initiative_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_initiative_propertytype.Href
	case "Role":
		res.valueString = ds_initiative_propertytype.Role
	case "Arcrole":
		res.valueString = ds_initiative_propertytype.Arcrole
	case "Title":
		res.valueString = ds_initiative_propertytype.Title
	case "Show":
		res.valueString = ds_initiative_propertytype.Show
	case "Actuate":
		res.valueString = ds_initiative_propertytype.Actuate
	case "DS_Initiative":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_initiative_propertytype.DS_Initiative != nil {
			res.valueString = ds_initiative_propertytype.DS_Initiative.Name
			res.ids = ds_initiative_propertytype.DS_Initiative.GongGetUUID(stage)
		}
	}
	return
}

func (ds_initiative_type *DS_Initiative_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_initiative_type.Name
	case "Id":
		res.valueString = ds_initiative_type.Id
	case "Uuid":
		res.valueString = ds_initiative_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_initiative_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_initiative_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_initiative_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_initiative_type.Superset {
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

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_otheraggregate_propertytype.Name
	case "NilReason":
		res.valueString = ds_otheraggregate_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_otheraggregate_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_otheraggregate_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_otheraggregate_propertytype.Href
	case "Role":
		res.valueString = ds_otheraggregate_propertytype.Role
	case "Arcrole":
		res.valueString = ds_otheraggregate_propertytype.Arcrole
	case "Title":
		res.valueString = ds_otheraggregate_propertytype.Title
	case "Show":
		res.valueString = ds_otheraggregate_propertytype.Show
	case "Actuate":
		res.valueString = ds_otheraggregate_propertytype.Actuate
	case "DS_OtherAggregate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
			res.valueString = ds_otheraggregate_propertytype.DS_OtherAggregate.Name
			res.ids = ds_otheraggregate_propertytype.DS_OtherAggregate.GongGetUUID(stage)
		}
	case "DS_StereoMate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_otheraggregate_propertytype.DS_StereoMate != nil {
			res.valueString = ds_otheraggregate_propertytype.DS_StereoMate.Name
			res.ids = ds_otheraggregate_propertytype.DS_StereoMate.GongGetUUID(stage)
		}
	}
	return
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_otheraggregate_type.Name
	case "Id":
		res.valueString = ds_otheraggregate_type.Id
	case "Uuid":
		res.valueString = ds_otheraggregate_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_otheraggregate_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_otheraggregate_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_otheraggregate_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_otheraggregate_type.Superset {
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

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_platform_propertytype.Name
	case "NilReason":
		res.valueString = ds_platform_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_platform_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_platform_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_platform_propertytype.Href
	case "Role":
		res.valueString = ds_platform_propertytype.Role
	case "Arcrole":
		res.valueString = ds_platform_propertytype.Arcrole
	case "Title":
		res.valueString = ds_platform_propertytype.Title
	case "Show":
		res.valueString = ds_platform_propertytype.Show
	case "Actuate":
		res.valueString = ds_platform_propertytype.Actuate
	case "DS_Platform":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_platform_propertytype.DS_Platform != nil {
			res.valueString = ds_platform_propertytype.DS_Platform.Name
			res.ids = ds_platform_propertytype.DS_Platform.GongGetUUID(stage)
		}
	}
	return
}

func (ds_platform_type *DS_Platform_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_platform_type.Name
	case "Id":
		res.valueString = ds_platform_type.Id
	case "Uuid":
		res.valueString = ds_platform_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_platform_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_platform_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_platform_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_platform_type.Superset {
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

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_productionseries_propertytype.Name
	case "NilReason":
		res.valueString = ds_productionseries_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_productionseries_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_productionseries_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_productionseries_propertytype.Href
	case "Role":
		res.valueString = ds_productionseries_propertytype.Role
	case "Arcrole":
		res.valueString = ds_productionseries_propertytype.Arcrole
	case "Title":
		res.valueString = ds_productionseries_propertytype.Title
	case "Show":
		res.valueString = ds_productionseries_propertytype.Show
	case "Actuate":
		res.valueString = ds_productionseries_propertytype.Actuate
	case "DS_ProductionSeries":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_productionseries_propertytype.DS_ProductionSeries != nil {
			res.valueString = ds_productionseries_propertytype.DS_ProductionSeries.Name
			res.ids = ds_productionseries_propertytype.DS_ProductionSeries.GongGetUUID(stage)
		}
	}
	return
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_productionseries_type.Name
	case "Id":
		res.valueString = ds_productionseries_type.Id
	case "Uuid":
		res.valueString = ds_productionseries_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_productionseries_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_productionseries_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_productionseries_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_productionseries_type.Superset {
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

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_sensor_propertytype.Name
	case "NilReason":
		res.valueString = ds_sensor_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_sensor_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_sensor_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_sensor_propertytype.Href
	case "Role":
		res.valueString = ds_sensor_propertytype.Role
	case "Arcrole":
		res.valueString = ds_sensor_propertytype.Arcrole
	case "Title":
		res.valueString = ds_sensor_propertytype.Title
	case "Show":
		res.valueString = ds_sensor_propertytype.Show
	case "Actuate":
		res.valueString = ds_sensor_propertytype.Actuate
	case "DS_Sensor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_sensor_propertytype.DS_Sensor != nil {
			res.valueString = ds_sensor_propertytype.DS_Sensor.Name
			res.ids = ds_sensor_propertytype.DS_Sensor.GongGetUUID(stage)
		}
	}
	return
}

func (ds_sensor_type *DS_Sensor_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_sensor_type.Name
	case "Id":
		res.valueString = ds_sensor_type.Id
	case "Uuid":
		res.valueString = ds_sensor_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_sensor_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_sensor_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_sensor_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_sensor_type.Superset {
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

func (ds_series_propertytype *DS_Series_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_series_propertytype.Name
	case "NilReason":
		res.valueString = ds_series_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_series_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_series_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_series_propertytype.Href
	case "Role":
		res.valueString = ds_series_propertytype.Role
	case "Arcrole":
		res.valueString = ds_series_propertytype.Arcrole
	case "Title":
		res.valueString = ds_series_propertytype.Title
	case "Show":
		res.valueString = ds_series_propertytype.Show
	case "Actuate":
		res.valueString = ds_series_propertytype.Actuate
	case "DS_Series":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_series_propertytype.DS_Series != nil {
			res.valueString = ds_series_propertytype.DS_Series.Name
			res.ids = ds_series_propertytype.DS_Series.GongGetUUID(stage)
		}
	case "DS_Platform":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_series_propertytype.DS_Platform != nil {
			res.valueString = ds_series_propertytype.DS_Platform.Name
			res.ids = ds_series_propertytype.DS_Platform.GongGetUUID(stage)
		}
	case "DS_Sensor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_series_propertytype.DS_Sensor != nil {
			res.valueString = ds_series_propertytype.DS_Sensor.Name
			res.ids = ds_series_propertytype.DS_Sensor.GongGetUUID(stage)
		}
	case "DS_ProductionSeries":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_series_propertytype.DS_ProductionSeries != nil {
			res.valueString = ds_series_propertytype.DS_ProductionSeries.Name
			res.ids = ds_series_propertytype.DS_ProductionSeries.GongGetUUID(stage)
		}
	}
	return
}

func (ds_series_type *DS_Series_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_series_type.Name
	case "Id":
		res.valueString = ds_series_type.Id
	case "Uuid":
		res.valueString = ds_series_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_series_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_series_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_series_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_series_type.Superset {
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

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_stereomate_propertytype.Name
	case "NilReason":
		res.valueString = ds_stereomate_propertytype.NilReason
	case "Uuidref":
		res.valueString = ds_stereomate_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = ds_stereomate_propertytype.Type_XSD
	case "Href":
		res.valueString = ds_stereomate_propertytype.Href
	case "Role":
		res.valueString = ds_stereomate_propertytype.Role
	case "Arcrole":
		res.valueString = ds_stereomate_propertytype.Arcrole
	case "Title":
		res.valueString = ds_stereomate_propertytype.Title
	case "Show":
		res.valueString = ds_stereomate_propertytype.Show
	case "Actuate":
		res.valueString = ds_stereomate_propertytype.Actuate
	case "DS_StereoMate":
		res.GongFieldValueType = GongFieldValueTypePointer
		if ds_stereomate_propertytype.DS_StereoMate != nil {
			res.valueString = ds_stereomate_propertytype.DS_StereoMate.Name
			res.ids = ds_stereomate_propertytype.DS_StereoMate.GongGetUUID(stage)
		}
	}
	return
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ds_stereomate_type.Name
	case "Id":
		res.valueString = ds_stereomate_type.Id
	case "Uuid":
		res.valueString = ds_stereomate_type.Uuid
	case "ComposedOf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_stereomate_type.ComposedOf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SeriesMetadata":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_stereomate_type.SeriesMetadata {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Subset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_stereomate_type.Subset {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Superset":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ds_stereomate_type.Superset {
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

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_metadata_propertytype.Name
	case "NilReason":
		res.valueString = md_metadata_propertytype.NilReason
	case "Uuidref":
		res.valueString = md_metadata_propertytype.Uuidref
	case "Type_XSD":
		res.valueString = md_metadata_propertytype.Type_XSD
	case "Href":
		res.valueString = md_metadata_propertytype.Href
	case "Role":
		res.valueString = md_metadata_propertytype.Role
	case "Arcrole":
		res.valueString = md_metadata_propertytype.Arcrole
	case "Title":
		res.valueString = md_metadata_propertytype.Title
	case "Show":
		res.valueString = md_metadata_propertytype.Show
	case "Actuate":
		res.valueString = md_metadata_propertytype.Actuate
	case "MD_Metadata":
		res.GongFieldValueType = GongFieldValueTypePointer
		if md_metadata_propertytype.MD_Metadata != nil {
			res.valueString = md_metadata_propertytype.MD_Metadata.Name
			res.ids = md_metadata_propertytype.MD_Metadata.GongGetUUID(stage)
		}
	}
	return
}

func (md_metadata_type *MD_Metadata_Type) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = md_metadata_type.Name
	case "Id":
		res.valueString = md_metadata_type.Id
	case "Uuid":
		res.valueString = md_metadata_type.Uuid
	case "Series":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_metadata_type.Series {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Describes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range md_metadata_type.Describes {
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
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetGongstructName() string {
	return "AbstractDS_Aggregate_Type"
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetGongstructName() string {
	return "DS_Aggregate_PropertyType"
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetGongstructName() string {
	return "DS_DataSet_PropertyType"
}

func (ds_dataset_type *DS_DataSet_Type) GongGetGongstructName() string {
	return "DS_DataSet_Type"
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetGongstructName() string {
	return "DS_Initiative_PropertyType"
}

func (ds_initiative_type *DS_Initiative_Type) GongGetGongstructName() string {
	return "DS_Initiative_Type"
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetGongstructName() string {
	return "DS_OtherAggregate_PropertyType"
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetGongstructName() string {
	return "DS_OtherAggregate_Type"
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetGongstructName() string {
	return "DS_Platform_PropertyType"
}

func (ds_platform_type *DS_Platform_Type) GongGetGongstructName() string {
	return "DS_Platform_Type"
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetGongstructName() string {
	return "DS_ProductionSeries_PropertyType"
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetGongstructName() string {
	return "DS_ProductionSeries_Type"
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetGongstructName() string {
	return "DS_Sensor_PropertyType"
}

func (ds_sensor_type *DS_Sensor_Type) GongGetGongstructName() string {
	return "DS_Sensor_Type"
}

func (ds_series_propertytype *DS_Series_PropertyType) GongGetGongstructName() string {
	return "DS_Series_PropertyType"
}

func (ds_series_type *DS_Series_Type) GongGetGongstructName() string {
	return "DS_Series_Type"
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetGongstructName() string {
	return "DS_StereoMate_PropertyType"
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetGongstructName() string {
	return "DS_StereoMate_Type"
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetGongstructName() string {
	return "MD_Metadata_PropertyType"
}

func (md_metadata_type *MD_Metadata_Type) GongGetGongstructName() string {
	return "MD_Metadata_Type"
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
	__gong__rebuildMapString(stage.AbstractDS_Aggregate_Types, &stage.AbstractDS_Aggregate_Types_mapString)

	__gong__rebuildMapString(stage.DS_Aggregate_PropertyTypes, &stage.DS_Aggregate_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_DataSet_PropertyTypes, &stage.DS_DataSet_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_DataSet_Types, &stage.DS_DataSet_Types_mapString)

	__gong__rebuildMapString(stage.DS_Initiative_PropertyTypes, &stage.DS_Initiative_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_Initiative_Types, &stage.DS_Initiative_Types_mapString)

	__gong__rebuildMapString(stage.DS_OtherAggregate_PropertyTypes, &stage.DS_OtherAggregate_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_OtherAggregate_Types, &stage.DS_OtherAggregate_Types_mapString)

	__gong__rebuildMapString(stage.DS_Platform_PropertyTypes, &stage.DS_Platform_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_Platform_Types, &stage.DS_Platform_Types_mapString)

	__gong__rebuildMapString(stage.DS_ProductionSeries_PropertyTypes, &stage.DS_ProductionSeries_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_ProductionSeries_Types, &stage.DS_ProductionSeries_Types_mapString)

	__gong__rebuildMapString(stage.DS_Sensor_PropertyTypes, &stage.DS_Sensor_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_Sensor_Types, &stage.DS_Sensor_Types_mapString)

	__gong__rebuildMapString(stage.DS_Series_PropertyTypes, &stage.DS_Series_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_Series_Types, &stage.DS_Series_Types_mapString)

	__gong__rebuildMapString(stage.DS_StereoMate_PropertyTypes, &stage.DS_StereoMate_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.DS_StereoMate_Types, &stage.DS_StereoMate_Types_mapString)

	__gong__rebuildMapString(stage.MD_Metadata_PropertyTypes, &stage.MD_Metadata_PropertyTypes_mapString)

	__gong__rebuildMapString(stage.MD_Metadata_Types, &stage.MD_Metadata_Types_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
