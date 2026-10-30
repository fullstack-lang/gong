// generated code - do not edit
package aixmfeatures

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
	A_extensions                map[*A_extension]struct{}
	A_extensions_instance       map[*A_extension]*A_extension
	A_extensions_mapString      map[string]*A_extension
	A_extensionOrder            uint
	A_extension_stagedOrder     map[*A_extension]uint
	A_extension_orderStaged     map[uint]*A_extension
	A_extensions_reference      map[*A_extension]*A_extension
	A_extensions_referenceOrder map[*A_extension]uint

	// insertion point for slice of pointers maps
	OnAfterA_extensionCreateCallback GongOnAfterCreateInterface[A_extension]
	OnAfterA_extensionUpdateCallback GongOnAfterUpdateInterface[A_extension]
	OnAfterA_extensionDeleteCallback GongOnAfterDeleteInterface[A_extension]

	A_extension_1s                map[*A_extension_1]struct{}
	A_extension_1s_instance       map[*A_extension_1]*A_extension_1
	A_extension_1s_mapString      map[string]*A_extension_1
	A_extension_1Order            uint
	A_extension_1_stagedOrder     map[*A_extension_1]uint
	A_extension_1_orderStaged     map[uint]*A_extension_1
	A_extension_1s_reference      map[*A_extension_1]*A_extension_1
	A_extension_1s_referenceOrder map[*A_extension_1]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_1CreateCallback GongOnAfterCreateInterface[A_extension_1]
	OnAfterA_extension_1UpdateCallback GongOnAfterUpdateInterface[A_extension_1]
	OnAfterA_extension_1DeleteCallback GongOnAfterDeleteInterface[A_extension_1]

	A_extension_10s                map[*A_extension_10]struct{}
	A_extension_10s_instance       map[*A_extension_10]*A_extension_10
	A_extension_10s_mapString      map[string]*A_extension_10
	A_extension_10Order            uint
	A_extension_10_stagedOrder     map[*A_extension_10]uint
	A_extension_10_orderStaged     map[uint]*A_extension_10
	A_extension_10s_reference      map[*A_extension_10]*A_extension_10
	A_extension_10s_referenceOrder map[*A_extension_10]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_10CreateCallback GongOnAfterCreateInterface[A_extension_10]
	OnAfterA_extension_10UpdateCallback GongOnAfterUpdateInterface[A_extension_10]
	OnAfterA_extension_10DeleteCallback GongOnAfterDeleteInterface[A_extension_10]

	A_extension_11s                map[*A_extension_11]struct{}
	A_extension_11s_instance       map[*A_extension_11]*A_extension_11
	A_extension_11s_mapString      map[string]*A_extension_11
	A_extension_11Order            uint
	A_extension_11_stagedOrder     map[*A_extension_11]uint
	A_extension_11_orderStaged     map[uint]*A_extension_11
	A_extension_11s_reference      map[*A_extension_11]*A_extension_11
	A_extension_11s_referenceOrder map[*A_extension_11]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_11CreateCallback GongOnAfterCreateInterface[A_extension_11]
	OnAfterA_extension_11UpdateCallback GongOnAfterUpdateInterface[A_extension_11]
	OnAfterA_extension_11DeleteCallback GongOnAfterDeleteInterface[A_extension_11]

	A_extension_12s                map[*A_extension_12]struct{}
	A_extension_12s_instance       map[*A_extension_12]*A_extension_12
	A_extension_12s_mapString      map[string]*A_extension_12
	A_extension_12Order            uint
	A_extension_12_stagedOrder     map[*A_extension_12]uint
	A_extension_12_orderStaged     map[uint]*A_extension_12
	A_extension_12s_reference      map[*A_extension_12]*A_extension_12
	A_extension_12s_referenceOrder map[*A_extension_12]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_12CreateCallback GongOnAfterCreateInterface[A_extension_12]
	OnAfterA_extension_12UpdateCallback GongOnAfterUpdateInterface[A_extension_12]
	OnAfterA_extension_12DeleteCallback GongOnAfterDeleteInterface[A_extension_12]

	A_extension_2s                map[*A_extension_2]struct{}
	A_extension_2s_instance       map[*A_extension_2]*A_extension_2
	A_extension_2s_mapString      map[string]*A_extension_2
	A_extension_2Order            uint
	A_extension_2_stagedOrder     map[*A_extension_2]uint
	A_extension_2_orderStaged     map[uint]*A_extension_2
	A_extension_2s_reference      map[*A_extension_2]*A_extension_2
	A_extension_2s_referenceOrder map[*A_extension_2]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_2CreateCallback GongOnAfterCreateInterface[A_extension_2]
	OnAfterA_extension_2UpdateCallback GongOnAfterUpdateInterface[A_extension_2]
	OnAfterA_extension_2DeleteCallback GongOnAfterDeleteInterface[A_extension_2]

	A_extension_3s                map[*A_extension_3]struct{}
	A_extension_3s_instance       map[*A_extension_3]*A_extension_3
	A_extension_3s_mapString      map[string]*A_extension_3
	A_extension_3Order            uint
	A_extension_3_stagedOrder     map[*A_extension_3]uint
	A_extension_3_orderStaged     map[uint]*A_extension_3
	A_extension_3s_reference      map[*A_extension_3]*A_extension_3
	A_extension_3s_referenceOrder map[*A_extension_3]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_3CreateCallback GongOnAfterCreateInterface[A_extension_3]
	OnAfterA_extension_3UpdateCallback GongOnAfterUpdateInterface[A_extension_3]
	OnAfterA_extension_3DeleteCallback GongOnAfterDeleteInterface[A_extension_3]

	A_extension_4s                map[*A_extension_4]struct{}
	A_extension_4s_instance       map[*A_extension_4]*A_extension_4
	A_extension_4s_mapString      map[string]*A_extension_4
	A_extension_4Order            uint
	A_extension_4_stagedOrder     map[*A_extension_4]uint
	A_extension_4_orderStaged     map[uint]*A_extension_4
	A_extension_4s_reference      map[*A_extension_4]*A_extension_4
	A_extension_4s_referenceOrder map[*A_extension_4]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_4CreateCallback GongOnAfterCreateInterface[A_extension_4]
	OnAfterA_extension_4UpdateCallback GongOnAfterUpdateInterface[A_extension_4]
	OnAfterA_extension_4DeleteCallback GongOnAfterDeleteInterface[A_extension_4]

	A_extension_5s                map[*A_extension_5]struct{}
	A_extension_5s_instance       map[*A_extension_5]*A_extension_5
	A_extension_5s_mapString      map[string]*A_extension_5
	A_extension_5Order            uint
	A_extension_5_stagedOrder     map[*A_extension_5]uint
	A_extension_5_orderStaged     map[uint]*A_extension_5
	A_extension_5s_reference      map[*A_extension_5]*A_extension_5
	A_extension_5s_referenceOrder map[*A_extension_5]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_5CreateCallback GongOnAfterCreateInterface[A_extension_5]
	OnAfterA_extension_5UpdateCallback GongOnAfterUpdateInterface[A_extension_5]
	OnAfterA_extension_5DeleteCallback GongOnAfterDeleteInterface[A_extension_5]

	A_extension_6s                map[*A_extension_6]struct{}
	A_extension_6s_instance       map[*A_extension_6]*A_extension_6
	A_extension_6s_mapString      map[string]*A_extension_6
	A_extension_6Order            uint
	A_extension_6_stagedOrder     map[*A_extension_6]uint
	A_extension_6_orderStaged     map[uint]*A_extension_6
	A_extension_6s_reference      map[*A_extension_6]*A_extension_6
	A_extension_6s_referenceOrder map[*A_extension_6]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_6CreateCallback GongOnAfterCreateInterface[A_extension_6]
	OnAfterA_extension_6UpdateCallback GongOnAfterUpdateInterface[A_extension_6]
	OnAfterA_extension_6DeleteCallback GongOnAfterDeleteInterface[A_extension_6]

	A_extension_7s                map[*A_extension_7]struct{}
	A_extension_7s_instance       map[*A_extension_7]*A_extension_7
	A_extension_7s_mapString      map[string]*A_extension_7
	A_extension_7Order            uint
	A_extension_7_stagedOrder     map[*A_extension_7]uint
	A_extension_7_orderStaged     map[uint]*A_extension_7
	A_extension_7s_reference      map[*A_extension_7]*A_extension_7
	A_extension_7s_referenceOrder map[*A_extension_7]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_7CreateCallback GongOnAfterCreateInterface[A_extension_7]
	OnAfterA_extension_7UpdateCallback GongOnAfterUpdateInterface[A_extension_7]
	OnAfterA_extension_7DeleteCallback GongOnAfterDeleteInterface[A_extension_7]

	A_extension_8s                map[*A_extension_8]struct{}
	A_extension_8s_instance       map[*A_extension_8]*A_extension_8
	A_extension_8s_mapString      map[string]*A_extension_8
	A_extension_8Order            uint
	A_extension_8_stagedOrder     map[*A_extension_8]uint
	A_extension_8_orderStaged     map[uint]*A_extension_8
	A_extension_8s_reference      map[*A_extension_8]*A_extension_8
	A_extension_8s_referenceOrder map[*A_extension_8]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_8CreateCallback GongOnAfterCreateInterface[A_extension_8]
	OnAfterA_extension_8UpdateCallback GongOnAfterUpdateInterface[A_extension_8]
	OnAfterA_extension_8DeleteCallback GongOnAfterDeleteInterface[A_extension_8]

	A_extension_9s                map[*A_extension_9]struct{}
	A_extension_9s_instance       map[*A_extension_9]*A_extension_9
	A_extension_9s_mapString      map[string]*A_extension_9
	A_extension_9Order            uint
	A_extension_9_stagedOrder     map[*A_extension_9]uint
	A_extension_9_orderStaged     map[uint]*A_extension_9
	A_extension_9s_reference      map[*A_extension_9]*A_extension_9
	A_extension_9s_referenceOrder map[*A_extension_9]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_9CreateCallback GongOnAfterCreateInterface[A_extension_9]
	OnAfterA_extension_9UpdateCallback GongOnAfterUpdateInterface[A_extension_9]
	OnAfterA_extension_9DeleteCallback GongOnAfterDeleteInterface[A_extension_9]

	AbstractPropertiesWithScheduleTypes                map[*AbstractPropertiesWithScheduleType]struct{}
	AbstractPropertiesWithScheduleTypes_instance       map[*AbstractPropertiesWithScheduleType]*AbstractPropertiesWithScheduleType
	AbstractPropertiesWithScheduleTypes_mapString      map[string]*AbstractPropertiesWithScheduleType
	AbstractPropertiesWithScheduleTypeOrder            uint
	AbstractPropertiesWithScheduleType_stagedOrder     map[*AbstractPropertiesWithScheduleType]uint
	AbstractPropertiesWithScheduleType_orderStaged     map[uint]*AbstractPropertiesWithScheduleType
	AbstractPropertiesWithScheduleTypes_reference      map[*AbstractPropertiesWithScheduleType]*AbstractPropertiesWithScheduleType
	AbstractPropertiesWithScheduleTypes_referenceOrder map[*AbstractPropertiesWithScheduleType]uint

	// insertion point for slice of pointers maps
	OnAfterAbstractPropertiesWithScheduleTypeCreateCallback GongOnAfterCreateInterface[AbstractPropertiesWithScheduleType]
	OnAfterAbstractPropertiesWithScheduleTypeUpdateCallback GongOnAfterUpdateInterface[AbstractPropertiesWithScheduleType]
	OnAfterAbstractPropertiesWithScheduleTypeDeleteCallback GongOnAfterDeleteInterface[AbstractPropertiesWithScheduleType]

	AeronauticalGroundLightTimeSlicePropertyTypes                map[*AeronauticalGroundLightTimeSlicePropertyType]struct{}
	AeronauticalGroundLightTimeSlicePropertyTypes_instance       map[*AeronauticalGroundLightTimeSlicePropertyType]*AeronauticalGroundLightTimeSlicePropertyType
	AeronauticalGroundLightTimeSlicePropertyTypes_mapString      map[string]*AeronauticalGroundLightTimeSlicePropertyType
	AeronauticalGroundLightTimeSlicePropertyTypeOrder            uint
	AeronauticalGroundLightTimeSlicePropertyType_stagedOrder     map[*AeronauticalGroundLightTimeSlicePropertyType]uint
	AeronauticalGroundLightTimeSlicePropertyType_orderStaged     map[uint]*AeronauticalGroundLightTimeSlicePropertyType
	AeronauticalGroundLightTimeSlicePropertyTypes_reference      map[*AeronauticalGroundLightTimeSlicePropertyType]*AeronauticalGroundLightTimeSlicePropertyType
	AeronauticalGroundLightTimeSlicePropertyTypes_referenceOrder map[*AeronauticalGroundLightTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAeronauticalGroundLightTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[AeronauticalGroundLightTimeSlicePropertyType]
	OnAfterAeronauticalGroundLightTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[AeronauticalGroundLightTimeSlicePropertyType]
	OnAfterAeronauticalGroundLightTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[AeronauticalGroundLightTimeSlicePropertyType]

	AeronauticalGroundLightTimeSliceTypes                map[*AeronauticalGroundLightTimeSliceType]struct{}
	AeronauticalGroundLightTimeSliceTypes_instance       map[*AeronauticalGroundLightTimeSliceType]*AeronauticalGroundLightTimeSliceType
	AeronauticalGroundLightTimeSliceTypes_mapString      map[string]*AeronauticalGroundLightTimeSliceType
	AeronauticalGroundLightTimeSliceTypeOrder            uint
	AeronauticalGroundLightTimeSliceType_stagedOrder     map[*AeronauticalGroundLightTimeSliceType]uint
	AeronauticalGroundLightTimeSliceType_orderStaged     map[uint]*AeronauticalGroundLightTimeSliceType
	AeronauticalGroundLightTimeSliceTypes_reference      map[*AeronauticalGroundLightTimeSliceType]*AeronauticalGroundLightTimeSliceType
	AeronauticalGroundLightTimeSliceTypes_referenceOrder map[*AeronauticalGroundLightTimeSliceType]uint

	// insertion point for slice of pointers maps
	AeronauticalGroundLightTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*AeronauticalGroundLightTimeSliceType

	AeronauticalGroundLightTimeSliceType_Extension_reverseMap map[*A_extension_3]*AeronauticalGroundLightTimeSliceType

	OnAfterAeronauticalGroundLightTimeSliceTypeCreateCallback GongOnAfterCreateInterface[AeronauticalGroundLightTimeSliceType]
	OnAfterAeronauticalGroundLightTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[AeronauticalGroundLightTimeSliceType]
	OnAfterAeronauticalGroundLightTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[AeronauticalGroundLightTimeSliceType]

	AeronauticalGroundLightTypes                map[*AeronauticalGroundLightType]struct{}
	AeronauticalGroundLightTypes_instance       map[*AeronauticalGroundLightType]*AeronauticalGroundLightType
	AeronauticalGroundLightTypes_mapString      map[string]*AeronauticalGroundLightType
	AeronauticalGroundLightTypeOrder            uint
	AeronauticalGroundLightType_stagedOrder     map[*AeronauticalGroundLightType]uint
	AeronauticalGroundLightType_orderStaged     map[uint]*AeronauticalGroundLightType
	AeronauticalGroundLightTypes_reference      map[*AeronauticalGroundLightType]*AeronauticalGroundLightType
	AeronauticalGroundLightTypes_referenceOrder map[*AeronauticalGroundLightType]uint

	// insertion point for slice of pointers maps
	AeronauticalGroundLightType_TimeSlice_reverseMap map[*AeronauticalGroundLightTimeSlicePropertyType]*AeronauticalGroundLightType

	OnAfterAeronauticalGroundLightTypeCreateCallback GongOnAfterCreateInterface[AeronauticalGroundLightType]
	OnAfterAeronauticalGroundLightTypeUpdateCallback GongOnAfterUpdateInterface[AeronauticalGroundLightType]
	OnAfterAeronauticalGroundLightTypeDeleteCallback GongOnAfterDeleteInterface[AeronauticalGroundLightType]

	AirportHeliportPropertyTypes                map[*AirportHeliportPropertyType]struct{}
	AirportHeliportPropertyTypes_instance       map[*AirportHeliportPropertyType]*AirportHeliportPropertyType
	AirportHeliportPropertyTypes_mapString      map[string]*AirportHeliportPropertyType
	AirportHeliportPropertyTypeOrder            uint
	AirportHeliportPropertyType_stagedOrder     map[*AirportHeliportPropertyType]uint
	AirportHeliportPropertyType_orderStaged     map[uint]*AirportHeliportPropertyType
	AirportHeliportPropertyTypes_reference      map[*AirportHeliportPropertyType]*AirportHeliportPropertyType
	AirportHeliportPropertyTypes_referenceOrder map[*AirportHeliportPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirportHeliportPropertyTypeCreateCallback GongOnAfterCreateInterface[AirportHeliportPropertyType]
	OnAfterAirportHeliportPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirportHeliportPropertyType]
	OnAfterAirportHeliportPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirportHeliportPropertyType]

	CurveTypes                map[*CurveType]struct{}
	CurveTypes_instance       map[*CurveType]*CurveType
	CurveTypes_mapString      map[string]*CurveType
	CurveTypeOrder            uint
	CurveType_stagedOrder     map[*CurveType]uint
	CurveType_orderStaged     map[uint]*CurveType
	CurveTypes_reference      map[*CurveType]*CurveType
	CurveTypes_referenceOrder map[*CurveType]uint

	// insertion point for slice of pointers maps
	CurveType_Annotation_reverseMap map[*NotePropertyType]*CurveType

	OnAfterCurveTypeCreateCallback GongOnAfterCreateInterface[CurveType]
	OnAfterCurveTypeUpdateCallback GongOnAfterUpdateInterface[CurveType]
	OnAfterCurveTypeDeleteCallback GongOnAfterDeleteInterface[CurveType]

	ElevatedCurvePropertyTypes                map[*ElevatedCurvePropertyType]struct{}
	ElevatedCurvePropertyTypes_instance       map[*ElevatedCurvePropertyType]*ElevatedCurvePropertyType
	ElevatedCurvePropertyTypes_mapString      map[string]*ElevatedCurvePropertyType
	ElevatedCurvePropertyTypeOrder            uint
	ElevatedCurvePropertyType_stagedOrder     map[*ElevatedCurvePropertyType]uint
	ElevatedCurvePropertyType_orderStaged     map[uint]*ElevatedCurvePropertyType
	ElevatedCurvePropertyTypes_reference      map[*ElevatedCurvePropertyType]*ElevatedCurvePropertyType
	ElevatedCurvePropertyTypes_referenceOrder map[*ElevatedCurvePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterElevatedCurvePropertyTypeCreateCallback GongOnAfterCreateInterface[ElevatedCurvePropertyType]
	OnAfterElevatedCurvePropertyTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedCurvePropertyType]
	OnAfterElevatedCurvePropertyTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedCurvePropertyType]

	ElevatedCurveTypes                map[*ElevatedCurveType]struct{}
	ElevatedCurveTypes_instance       map[*ElevatedCurveType]*ElevatedCurveType
	ElevatedCurveTypes_mapString      map[string]*ElevatedCurveType
	ElevatedCurveTypeOrder            uint
	ElevatedCurveType_stagedOrder     map[*ElevatedCurveType]uint
	ElevatedCurveType_orderStaged     map[uint]*ElevatedCurveType
	ElevatedCurveTypes_reference      map[*ElevatedCurveType]*ElevatedCurveType
	ElevatedCurveTypes_referenceOrder map[*ElevatedCurveType]uint

	// insertion point for slice of pointers maps
	ElevatedCurveType_Annotation_reverseMap map[*NotePropertyType]*ElevatedCurveType

	ElevatedCurveType_Extension_reverseMap map[*A_extension]*ElevatedCurveType

	OnAfterElevatedCurveTypeCreateCallback GongOnAfterCreateInterface[ElevatedCurveType]
	OnAfterElevatedCurveTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedCurveType]
	OnAfterElevatedCurveTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedCurveType]

	ElevatedPointPropertyTypes                map[*ElevatedPointPropertyType]struct{}
	ElevatedPointPropertyTypes_instance       map[*ElevatedPointPropertyType]*ElevatedPointPropertyType
	ElevatedPointPropertyTypes_mapString      map[string]*ElevatedPointPropertyType
	ElevatedPointPropertyTypeOrder            uint
	ElevatedPointPropertyType_stagedOrder     map[*ElevatedPointPropertyType]uint
	ElevatedPointPropertyType_orderStaged     map[uint]*ElevatedPointPropertyType
	ElevatedPointPropertyTypes_reference      map[*ElevatedPointPropertyType]*ElevatedPointPropertyType
	ElevatedPointPropertyTypes_referenceOrder map[*ElevatedPointPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterElevatedPointPropertyTypeCreateCallback GongOnAfterCreateInterface[ElevatedPointPropertyType]
	OnAfterElevatedPointPropertyTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedPointPropertyType]
	OnAfterElevatedPointPropertyTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedPointPropertyType]

	ElevatedPointTypes                map[*ElevatedPointType]struct{}
	ElevatedPointTypes_instance       map[*ElevatedPointType]*ElevatedPointType
	ElevatedPointTypes_mapString      map[string]*ElevatedPointType
	ElevatedPointTypeOrder            uint
	ElevatedPointType_stagedOrder     map[*ElevatedPointType]uint
	ElevatedPointType_orderStaged     map[uint]*ElevatedPointType
	ElevatedPointTypes_reference      map[*ElevatedPointType]*ElevatedPointType
	ElevatedPointTypes_referenceOrder map[*ElevatedPointType]uint

	// insertion point for slice of pointers maps
	ElevatedPointType_Annotation_reverseMap map[*NotePropertyType]*ElevatedPointType

	ElevatedPointType_Extension_reverseMap map[*A_extension_1]*ElevatedPointType

	OnAfterElevatedPointTypeCreateCallback GongOnAfterCreateInterface[ElevatedPointType]
	OnAfterElevatedPointTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedPointType]
	OnAfterElevatedPointTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedPointType]

	ElevatedSurfacePropertyTypes                map[*ElevatedSurfacePropertyType]struct{}
	ElevatedSurfacePropertyTypes_instance       map[*ElevatedSurfacePropertyType]*ElevatedSurfacePropertyType
	ElevatedSurfacePropertyTypes_mapString      map[string]*ElevatedSurfacePropertyType
	ElevatedSurfacePropertyTypeOrder            uint
	ElevatedSurfacePropertyType_stagedOrder     map[*ElevatedSurfacePropertyType]uint
	ElevatedSurfacePropertyType_orderStaged     map[uint]*ElevatedSurfacePropertyType
	ElevatedSurfacePropertyTypes_reference      map[*ElevatedSurfacePropertyType]*ElevatedSurfacePropertyType
	ElevatedSurfacePropertyTypes_referenceOrder map[*ElevatedSurfacePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterElevatedSurfacePropertyTypeCreateCallback GongOnAfterCreateInterface[ElevatedSurfacePropertyType]
	OnAfterElevatedSurfacePropertyTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedSurfacePropertyType]
	OnAfterElevatedSurfacePropertyTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedSurfacePropertyType]

	ElevatedSurfaceTypes                map[*ElevatedSurfaceType]struct{}
	ElevatedSurfaceTypes_instance       map[*ElevatedSurfaceType]*ElevatedSurfaceType
	ElevatedSurfaceTypes_mapString      map[string]*ElevatedSurfaceType
	ElevatedSurfaceTypeOrder            uint
	ElevatedSurfaceType_stagedOrder     map[*ElevatedSurfaceType]uint
	ElevatedSurfaceType_orderStaged     map[uint]*ElevatedSurfaceType
	ElevatedSurfaceTypes_reference      map[*ElevatedSurfaceType]*ElevatedSurfaceType
	ElevatedSurfaceTypes_referenceOrder map[*ElevatedSurfaceType]uint

	// insertion point for slice of pointers maps
	ElevatedSurfaceType_Annotation_reverseMap map[*NotePropertyType]*ElevatedSurfaceType

	ElevatedSurfaceType_Extension_reverseMap map[*A_extension_2]*ElevatedSurfaceType

	OnAfterElevatedSurfaceTypeCreateCallback GongOnAfterCreateInterface[ElevatedSurfaceType]
	OnAfterElevatedSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedSurfaceType]
	OnAfterElevatedSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedSurfaceType]

	GroundLightSystemPropertyTypes                map[*GroundLightSystemPropertyType]struct{}
	GroundLightSystemPropertyTypes_instance       map[*GroundLightSystemPropertyType]*GroundLightSystemPropertyType
	GroundLightSystemPropertyTypes_mapString      map[string]*GroundLightSystemPropertyType
	GroundLightSystemPropertyTypeOrder            uint
	GroundLightSystemPropertyType_stagedOrder     map[*GroundLightSystemPropertyType]uint
	GroundLightSystemPropertyType_orderStaged     map[uint]*GroundLightSystemPropertyType
	GroundLightSystemPropertyTypes_reference      map[*GroundLightSystemPropertyType]*GroundLightSystemPropertyType
	GroundLightSystemPropertyTypes_referenceOrder map[*GroundLightSystemPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterGroundLightSystemPropertyTypeCreateCallback GongOnAfterCreateInterface[GroundLightSystemPropertyType]
	OnAfterGroundLightSystemPropertyTypeUpdateCallback GongOnAfterUpdateInterface[GroundLightSystemPropertyType]
	OnAfterGroundLightSystemPropertyTypeDeleteCallback GongOnAfterDeleteInterface[GroundLightSystemPropertyType]

	LightElementPropertyTypes                map[*LightElementPropertyType]struct{}
	LightElementPropertyTypes_instance       map[*LightElementPropertyType]*LightElementPropertyType
	LightElementPropertyTypes_mapString      map[string]*LightElementPropertyType
	LightElementPropertyTypeOrder            uint
	LightElementPropertyType_stagedOrder     map[*LightElementPropertyType]uint
	LightElementPropertyType_orderStaged     map[uint]*LightElementPropertyType
	LightElementPropertyTypes_reference      map[*LightElementPropertyType]*LightElementPropertyType
	LightElementPropertyTypes_referenceOrder map[*LightElementPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterLightElementPropertyTypeCreateCallback GongOnAfterCreateInterface[LightElementPropertyType]
	OnAfterLightElementPropertyTypeUpdateCallback GongOnAfterUpdateInterface[LightElementPropertyType]
	OnAfterLightElementPropertyTypeDeleteCallback GongOnAfterDeleteInterface[LightElementPropertyType]

	LightElementStatusPropertyTypes                map[*LightElementStatusPropertyType]struct{}
	LightElementStatusPropertyTypes_instance       map[*LightElementStatusPropertyType]*LightElementStatusPropertyType
	LightElementStatusPropertyTypes_mapString      map[string]*LightElementStatusPropertyType
	LightElementStatusPropertyTypeOrder            uint
	LightElementStatusPropertyType_stagedOrder     map[*LightElementStatusPropertyType]uint
	LightElementStatusPropertyType_orderStaged     map[uint]*LightElementStatusPropertyType
	LightElementStatusPropertyTypes_reference      map[*LightElementStatusPropertyType]*LightElementStatusPropertyType
	LightElementStatusPropertyTypes_referenceOrder map[*LightElementStatusPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterLightElementStatusPropertyTypeCreateCallback GongOnAfterCreateInterface[LightElementStatusPropertyType]
	OnAfterLightElementStatusPropertyTypeUpdateCallback GongOnAfterUpdateInterface[LightElementStatusPropertyType]
	OnAfterLightElementStatusPropertyTypeDeleteCallback GongOnAfterDeleteInterface[LightElementStatusPropertyType]

	LightElementStatusTypes                map[*LightElementStatusType]struct{}
	LightElementStatusTypes_instance       map[*LightElementStatusType]*LightElementStatusType
	LightElementStatusTypes_mapString      map[string]*LightElementStatusType
	LightElementStatusTypeOrder            uint
	LightElementStatusType_stagedOrder     map[*LightElementStatusType]uint
	LightElementStatusType_orderStaged     map[uint]*LightElementStatusType
	LightElementStatusTypes_reference      map[*LightElementStatusType]*LightElementStatusType
	LightElementStatusTypes_referenceOrder map[*LightElementStatusType]uint

	// insertion point for slice of pointers maps
	LightElementStatusType_TimeInterval_reverseMap map[*TimesheetPropertyType]*LightElementStatusType

	LightElementStatusType_Annotation_reverseMap map[*NotePropertyType]*LightElementStatusType

	LightElementStatusType_SpecialDateAuthority_reverseMap map[*OrganisationAuthorityPropertyType]*LightElementStatusType

	LightElementStatusType_Extension_reverseMap map[*A_extension_11]*LightElementStatusType

	OnAfterLightElementStatusTypeCreateCallback GongOnAfterCreateInterface[LightElementStatusType]
	OnAfterLightElementStatusTypeUpdateCallback GongOnAfterUpdateInterface[LightElementStatusType]
	OnAfterLightElementStatusTypeDeleteCallback GongOnAfterDeleteInterface[LightElementStatusType]

	LightElementTypes                map[*LightElementType]struct{}
	LightElementTypes_instance       map[*LightElementType]*LightElementType
	LightElementTypes_mapString      map[string]*LightElementType
	LightElementTypeOrder            uint
	LightElementType_stagedOrder     map[*LightElementType]uint
	LightElementType_orderStaged     map[uint]*LightElementType
	LightElementTypes_reference      map[*LightElementType]*LightElementType
	LightElementTypes_referenceOrder map[*LightElementType]uint

	// insertion point for slice of pointers maps
	LightElementType_Annotation_reverseMap map[*NotePropertyType]*LightElementType

	LightElementType_Availability_reverseMap map[*LightElementStatusPropertyType]*LightElementType

	LightElementType_Extension_reverseMap map[*A_extension_10]*LightElementType

	OnAfterLightElementTypeCreateCallback GongOnAfterCreateInterface[LightElementType]
	OnAfterLightElementTypeUpdateCallback GongOnAfterUpdateInterface[LightElementType]
	OnAfterLightElementTypeDeleteCallback GongOnAfterDeleteInterface[LightElementType]

	LinguisticNotePropertyTypes                map[*LinguisticNotePropertyType]struct{}
	LinguisticNotePropertyTypes_instance       map[*LinguisticNotePropertyType]*LinguisticNotePropertyType
	LinguisticNotePropertyTypes_mapString      map[string]*LinguisticNotePropertyType
	LinguisticNotePropertyTypeOrder            uint
	LinguisticNotePropertyType_stagedOrder     map[*LinguisticNotePropertyType]uint
	LinguisticNotePropertyType_orderStaged     map[uint]*LinguisticNotePropertyType
	LinguisticNotePropertyTypes_reference      map[*LinguisticNotePropertyType]*LinguisticNotePropertyType
	LinguisticNotePropertyTypes_referenceOrder map[*LinguisticNotePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterLinguisticNotePropertyTypeCreateCallback GongOnAfterCreateInterface[LinguisticNotePropertyType]
	OnAfterLinguisticNotePropertyTypeUpdateCallback GongOnAfterUpdateInterface[LinguisticNotePropertyType]
	OnAfterLinguisticNotePropertyTypeDeleteCallback GongOnAfterDeleteInterface[LinguisticNotePropertyType]

	LinguisticNoteTypes                map[*LinguisticNoteType]struct{}
	LinguisticNoteTypes_instance       map[*LinguisticNoteType]*LinguisticNoteType
	LinguisticNoteTypes_mapString      map[string]*LinguisticNoteType
	LinguisticNoteTypeOrder            uint
	LinguisticNoteType_stagedOrder     map[*LinguisticNoteType]uint
	LinguisticNoteType_orderStaged     map[uint]*LinguisticNoteType
	LinguisticNoteTypes_reference      map[*LinguisticNoteType]*LinguisticNoteType
	LinguisticNoteTypes_referenceOrder map[*LinguisticNoteType]uint

	// insertion point for slice of pointers maps
	LinguisticNoteType_Extension_reverseMap map[*A_extension_4]*LinguisticNoteType

	OnAfterLinguisticNoteTypeCreateCallback GongOnAfterCreateInterface[LinguisticNoteType]
	OnAfterLinguisticNoteTypeUpdateCallback GongOnAfterUpdateInterface[LinguisticNoteType]
	OnAfterLinguisticNoteTypeDeleteCallback GongOnAfterDeleteInterface[LinguisticNoteType]

	MarkerBeaconPropertyTypes                map[*MarkerBeaconPropertyType]struct{}
	MarkerBeaconPropertyTypes_instance       map[*MarkerBeaconPropertyType]*MarkerBeaconPropertyType
	MarkerBeaconPropertyTypes_mapString      map[string]*MarkerBeaconPropertyType
	MarkerBeaconPropertyTypeOrder            uint
	MarkerBeaconPropertyType_stagedOrder     map[*MarkerBeaconPropertyType]uint
	MarkerBeaconPropertyType_orderStaged     map[uint]*MarkerBeaconPropertyType
	MarkerBeaconPropertyTypes_reference      map[*MarkerBeaconPropertyType]*MarkerBeaconPropertyType
	MarkerBeaconPropertyTypes_referenceOrder map[*MarkerBeaconPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterMarkerBeaconPropertyTypeCreateCallback GongOnAfterCreateInterface[MarkerBeaconPropertyType]
	OnAfterMarkerBeaconPropertyTypeUpdateCallback GongOnAfterUpdateInterface[MarkerBeaconPropertyType]
	OnAfterMarkerBeaconPropertyTypeDeleteCallback GongOnAfterDeleteInterface[MarkerBeaconPropertyType]

	NavaidEquipmentPropertyTypes                map[*NavaidEquipmentPropertyType]struct{}
	NavaidEquipmentPropertyTypes_instance       map[*NavaidEquipmentPropertyType]*NavaidEquipmentPropertyType
	NavaidEquipmentPropertyTypes_mapString      map[string]*NavaidEquipmentPropertyType
	NavaidEquipmentPropertyTypeOrder            uint
	NavaidEquipmentPropertyType_stagedOrder     map[*NavaidEquipmentPropertyType]uint
	NavaidEquipmentPropertyType_orderStaged     map[uint]*NavaidEquipmentPropertyType
	NavaidEquipmentPropertyTypes_reference      map[*NavaidEquipmentPropertyType]*NavaidEquipmentPropertyType
	NavaidEquipmentPropertyTypes_referenceOrder map[*NavaidEquipmentPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterNavaidEquipmentPropertyTypeCreateCallback GongOnAfterCreateInterface[NavaidEquipmentPropertyType]
	OnAfterNavaidEquipmentPropertyTypeUpdateCallback GongOnAfterUpdateInterface[NavaidEquipmentPropertyType]
	OnAfterNavaidEquipmentPropertyTypeDeleteCallback GongOnAfterDeleteInterface[NavaidEquipmentPropertyType]

	NotePropertyTypes                map[*NotePropertyType]struct{}
	NotePropertyTypes_instance       map[*NotePropertyType]*NotePropertyType
	NotePropertyTypes_mapString      map[string]*NotePropertyType
	NotePropertyTypeOrder            uint
	NotePropertyType_stagedOrder     map[*NotePropertyType]uint
	NotePropertyType_orderStaged     map[uint]*NotePropertyType
	NotePropertyTypes_reference      map[*NotePropertyType]*NotePropertyType
	NotePropertyTypes_referenceOrder map[*NotePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterNotePropertyTypeCreateCallback GongOnAfterCreateInterface[NotePropertyType]
	OnAfterNotePropertyTypeUpdateCallback GongOnAfterUpdateInterface[NotePropertyType]
	OnAfterNotePropertyTypeDeleteCallback GongOnAfterDeleteInterface[NotePropertyType]

	NoteTypes                map[*NoteType]struct{}
	NoteTypes_instance       map[*NoteType]*NoteType
	NoteTypes_mapString      map[string]*NoteType
	NoteTypeOrder            uint
	NoteType_stagedOrder     map[*NoteType]uint
	NoteType_orderStaged     map[uint]*NoteType
	NoteTypes_reference      map[*NoteType]*NoteType
	NoteTypes_referenceOrder map[*NoteType]uint

	// insertion point for slice of pointers maps
	NoteType_TranslatedNote_reverseMap map[*LinguisticNotePropertyType]*NoteType

	NoteType_Extension_reverseMap map[*A_extension_5]*NoteType

	OnAfterNoteTypeCreateCallback GongOnAfterCreateInterface[NoteType]
	OnAfterNoteTypeUpdateCallback GongOnAfterUpdateInterface[NoteType]
	OnAfterNoteTypeDeleteCallback GongOnAfterDeleteInterface[NoteType]

	ObstacleAreaTimeSlicePropertyTypes                map[*ObstacleAreaTimeSlicePropertyType]struct{}
	ObstacleAreaTimeSlicePropertyTypes_instance       map[*ObstacleAreaTimeSlicePropertyType]*ObstacleAreaTimeSlicePropertyType
	ObstacleAreaTimeSlicePropertyTypes_mapString      map[string]*ObstacleAreaTimeSlicePropertyType
	ObstacleAreaTimeSlicePropertyTypeOrder            uint
	ObstacleAreaTimeSlicePropertyType_stagedOrder     map[*ObstacleAreaTimeSlicePropertyType]uint
	ObstacleAreaTimeSlicePropertyType_orderStaged     map[uint]*ObstacleAreaTimeSlicePropertyType
	ObstacleAreaTimeSlicePropertyTypes_reference      map[*ObstacleAreaTimeSlicePropertyType]*ObstacleAreaTimeSlicePropertyType
	ObstacleAreaTimeSlicePropertyTypes_referenceOrder map[*ObstacleAreaTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterObstacleAreaTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[ObstacleAreaTimeSlicePropertyType]
	OnAfterObstacleAreaTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[ObstacleAreaTimeSlicePropertyType]
	OnAfterObstacleAreaTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[ObstacleAreaTimeSlicePropertyType]

	ObstacleAreaTimeSliceTypes                map[*ObstacleAreaTimeSliceType]struct{}
	ObstacleAreaTimeSliceTypes_instance       map[*ObstacleAreaTimeSliceType]*ObstacleAreaTimeSliceType
	ObstacleAreaTimeSliceTypes_mapString      map[string]*ObstacleAreaTimeSliceType
	ObstacleAreaTimeSliceTypeOrder            uint
	ObstacleAreaTimeSliceType_stagedOrder     map[*ObstacleAreaTimeSliceType]uint
	ObstacleAreaTimeSliceType_orderStaged     map[uint]*ObstacleAreaTimeSliceType
	ObstacleAreaTimeSliceTypes_reference      map[*ObstacleAreaTimeSliceType]*ObstacleAreaTimeSliceType
	ObstacleAreaTimeSliceTypes_referenceOrder map[*ObstacleAreaTimeSliceType]uint

	// insertion point for slice of pointers maps
	ObstacleAreaTimeSliceType_Obstacle_reverseMap map[*VerticalStructurePropertyType]*ObstacleAreaTimeSliceType

	ObstacleAreaTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*ObstacleAreaTimeSliceType

	ObstacleAreaTimeSliceType_Extension_reverseMap map[*A_extension_6]*ObstacleAreaTimeSliceType

	OnAfterObstacleAreaTimeSliceTypeCreateCallback GongOnAfterCreateInterface[ObstacleAreaTimeSliceType]
	OnAfterObstacleAreaTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[ObstacleAreaTimeSliceType]
	OnAfterObstacleAreaTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[ObstacleAreaTimeSliceType]

	ObstacleAreaTypes                map[*ObstacleAreaType]struct{}
	ObstacleAreaTypes_instance       map[*ObstacleAreaType]*ObstacleAreaType
	ObstacleAreaTypes_mapString      map[string]*ObstacleAreaType
	ObstacleAreaTypeOrder            uint
	ObstacleAreaType_stagedOrder     map[*ObstacleAreaType]uint
	ObstacleAreaType_orderStaged     map[uint]*ObstacleAreaType
	ObstacleAreaTypes_reference      map[*ObstacleAreaType]*ObstacleAreaType
	ObstacleAreaTypes_referenceOrder map[*ObstacleAreaType]uint

	// insertion point for slice of pointers maps
	ObstacleAreaType_TimeSlice_reverseMap map[*ObstacleAreaTimeSlicePropertyType]*ObstacleAreaType

	OnAfterObstacleAreaTypeCreateCallback GongOnAfterCreateInterface[ObstacleAreaType]
	OnAfterObstacleAreaTypeUpdateCallback GongOnAfterUpdateInterface[ObstacleAreaType]
	OnAfterObstacleAreaTypeDeleteCallback GongOnAfterDeleteInterface[ObstacleAreaType]

	OrganisationAuthorityPropertyTypes                map[*OrganisationAuthorityPropertyType]struct{}
	OrganisationAuthorityPropertyTypes_instance       map[*OrganisationAuthorityPropertyType]*OrganisationAuthorityPropertyType
	OrganisationAuthorityPropertyTypes_mapString      map[string]*OrganisationAuthorityPropertyType
	OrganisationAuthorityPropertyTypeOrder            uint
	OrganisationAuthorityPropertyType_stagedOrder     map[*OrganisationAuthorityPropertyType]uint
	OrganisationAuthorityPropertyType_orderStaged     map[uint]*OrganisationAuthorityPropertyType
	OrganisationAuthorityPropertyTypes_reference      map[*OrganisationAuthorityPropertyType]*OrganisationAuthorityPropertyType
	OrganisationAuthorityPropertyTypes_referenceOrder map[*OrganisationAuthorityPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterOrganisationAuthorityPropertyTypeCreateCallback GongOnAfterCreateInterface[OrganisationAuthorityPropertyType]
	OnAfterOrganisationAuthorityPropertyTypeUpdateCallback GongOnAfterUpdateInterface[OrganisationAuthorityPropertyType]
	OnAfterOrganisationAuthorityPropertyTypeDeleteCallback GongOnAfterDeleteInterface[OrganisationAuthorityPropertyType]

	PassengerServicePropertyTypes                map[*PassengerServicePropertyType]struct{}
	PassengerServicePropertyTypes_instance       map[*PassengerServicePropertyType]*PassengerServicePropertyType
	PassengerServicePropertyTypes_mapString      map[string]*PassengerServicePropertyType
	PassengerServicePropertyTypeOrder            uint
	PassengerServicePropertyType_stagedOrder     map[*PassengerServicePropertyType]uint
	PassengerServicePropertyType_orderStaged     map[uint]*PassengerServicePropertyType
	PassengerServicePropertyTypes_reference      map[*PassengerServicePropertyType]*PassengerServicePropertyType
	PassengerServicePropertyTypes_referenceOrder map[*PassengerServicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterPassengerServicePropertyTypeCreateCallback GongOnAfterCreateInterface[PassengerServicePropertyType]
	OnAfterPassengerServicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[PassengerServicePropertyType]
	OnAfterPassengerServicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[PassengerServicePropertyType]

	PointTypes                map[*PointType]struct{}
	PointTypes_instance       map[*PointType]*PointType
	PointTypes_mapString      map[string]*PointType
	PointTypeOrder            uint
	PointType_stagedOrder     map[*PointType]uint
	PointType_orderStaged     map[uint]*PointType
	PointTypes_reference      map[*PointType]*PointType
	PointTypes_referenceOrder map[*PointType]uint

	// insertion point for slice of pointers maps
	PointType_Annotation_reverseMap map[*NotePropertyType]*PointType

	OnAfterPointTypeCreateCallback GongOnAfterCreateInterface[PointType]
	OnAfterPointTypeUpdateCallback GongOnAfterUpdateInterface[PointType]
	OnAfterPointTypeDeleteCallback GongOnAfterDeleteInterface[PointType]

	RunwayDirectionPropertyTypes                map[*RunwayDirectionPropertyType]struct{}
	RunwayDirectionPropertyTypes_instance       map[*RunwayDirectionPropertyType]*RunwayDirectionPropertyType
	RunwayDirectionPropertyTypes_mapString      map[string]*RunwayDirectionPropertyType
	RunwayDirectionPropertyTypeOrder            uint
	RunwayDirectionPropertyType_stagedOrder     map[*RunwayDirectionPropertyType]uint
	RunwayDirectionPropertyType_orderStaged     map[uint]*RunwayDirectionPropertyType
	RunwayDirectionPropertyTypes_reference      map[*RunwayDirectionPropertyType]*RunwayDirectionPropertyType
	RunwayDirectionPropertyTypes_referenceOrder map[*RunwayDirectionPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterRunwayDirectionPropertyTypeCreateCallback GongOnAfterCreateInterface[RunwayDirectionPropertyType]
	OnAfterRunwayDirectionPropertyTypeUpdateCallback GongOnAfterUpdateInterface[RunwayDirectionPropertyType]
	OnAfterRunwayDirectionPropertyTypeDeleteCallback GongOnAfterDeleteInterface[RunwayDirectionPropertyType]

	ServicePropertyTypes                map[*ServicePropertyType]struct{}
	ServicePropertyTypes_instance       map[*ServicePropertyType]*ServicePropertyType
	ServicePropertyTypes_mapString      map[string]*ServicePropertyType
	ServicePropertyTypeOrder            uint
	ServicePropertyType_stagedOrder     map[*ServicePropertyType]uint
	ServicePropertyType_orderStaged     map[uint]*ServicePropertyType
	ServicePropertyTypes_reference      map[*ServicePropertyType]*ServicePropertyType
	ServicePropertyTypes_referenceOrder map[*ServicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterServicePropertyTypeCreateCallback GongOnAfterCreateInterface[ServicePropertyType]
	OnAfterServicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[ServicePropertyType]
	OnAfterServicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[ServicePropertyType]

	SpecialNavigationStationPropertyTypes                map[*SpecialNavigationStationPropertyType]struct{}
	SpecialNavigationStationPropertyTypes_instance       map[*SpecialNavigationStationPropertyType]*SpecialNavigationStationPropertyType
	SpecialNavigationStationPropertyTypes_mapString      map[string]*SpecialNavigationStationPropertyType
	SpecialNavigationStationPropertyTypeOrder            uint
	SpecialNavigationStationPropertyType_stagedOrder     map[*SpecialNavigationStationPropertyType]uint
	SpecialNavigationStationPropertyType_orderStaged     map[uint]*SpecialNavigationStationPropertyType
	SpecialNavigationStationPropertyTypes_reference      map[*SpecialNavigationStationPropertyType]*SpecialNavigationStationPropertyType
	SpecialNavigationStationPropertyTypes_referenceOrder map[*SpecialNavigationStationPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterSpecialNavigationStationPropertyTypeCreateCallback GongOnAfterCreateInterface[SpecialNavigationStationPropertyType]
	OnAfterSpecialNavigationStationPropertyTypeUpdateCallback GongOnAfterUpdateInterface[SpecialNavigationStationPropertyType]
	OnAfterSpecialNavigationStationPropertyTypeDeleteCallback GongOnAfterDeleteInterface[SpecialNavigationStationPropertyType]

	SurfacePropertyTypes                map[*SurfacePropertyType]struct{}
	SurfacePropertyTypes_instance       map[*SurfacePropertyType]*SurfacePropertyType
	SurfacePropertyTypes_mapString      map[string]*SurfacePropertyType
	SurfacePropertyTypeOrder            uint
	SurfacePropertyType_stagedOrder     map[*SurfacePropertyType]uint
	SurfacePropertyType_orderStaged     map[uint]*SurfacePropertyType
	SurfacePropertyTypes_reference      map[*SurfacePropertyType]*SurfacePropertyType
	SurfacePropertyTypes_referenceOrder map[*SurfacePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterSurfacePropertyTypeCreateCallback GongOnAfterCreateInterface[SurfacePropertyType]
	OnAfterSurfacePropertyTypeUpdateCallback GongOnAfterUpdateInterface[SurfacePropertyType]
	OnAfterSurfacePropertyTypeDeleteCallback GongOnAfterDeleteInterface[SurfacePropertyType]

	SurfaceTypes                map[*SurfaceType]struct{}
	SurfaceTypes_instance       map[*SurfaceType]*SurfaceType
	SurfaceTypes_mapString      map[string]*SurfaceType
	SurfaceTypeOrder            uint
	SurfaceType_stagedOrder     map[*SurfaceType]uint
	SurfaceType_orderStaged     map[uint]*SurfaceType
	SurfaceTypes_reference      map[*SurfaceType]*SurfaceType
	SurfaceTypes_referenceOrder map[*SurfaceType]uint

	// insertion point for slice of pointers maps
	SurfaceType_Annotation_reverseMap map[*NotePropertyType]*SurfaceType

	OnAfterSurfaceTypeCreateCallback GongOnAfterCreateInterface[SurfaceType]
	OnAfterSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[SurfaceType]
	OnAfterSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[SurfaceType]

	TimesheetPropertyTypes                map[*TimesheetPropertyType]struct{}
	TimesheetPropertyTypes_instance       map[*TimesheetPropertyType]*TimesheetPropertyType
	TimesheetPropertyTypes_mapString      map[string]*TimesheetPropertyType
	TimesheetPropertyTypeOrder            uint
	TimesheetPropertyType_stagedOrder     map[*TimesheetPropertyType]uint
	TimesheetPropertyType_orderStaged     map[uint]*TimesheetPropertyType
	TimesheetPropertyTypes_reference      map[*TimesheetPropertyType]*TimesheetPropertyType
	TimesheetPropertyTypes_referenceOrder map[*TimesheetPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterTimesheetPropertyTypeCreateCallback GongOnAfterCreateInterface[TimesheetPropertyType]
	OnAfterTimesheetPropertyTypeUpdateCallback GongOnAfterUpdateInterface[TimesheetPropertyType]
	OnAfterTimesheetPropertyTypeDeleteCallback GongOnAfterDeleteInterface[TimesheetPropertyType]

	TimesheetTypes                map[*TimesheetType]struct{}
	TimesheetTypes_instance       map[*TimesheetType]*TimesheetType
	TimesheetTypes_mapString      map[string]*TimesheetType
	TimesheetTypeOrder            uint
	TimesheetType_stagedOrder     map[*TimesheetType]uint
	TimesheetType_orderStaged     map[uint]*TimesheetType
	TimesheetTypes_reference      map[*TimesheetType]*TimesheetType
	TimesheetTypes_referenceOrder map[*TimesheetType]uint

	// insertion point for slice of pointers maps
	TimesheetType_Annotation_reverseMap map[*NotePropertyType]*TimesheetType

	TimesheetType_Extension_reverseMap map[*A_extension_12]*TimesheetType

	OnAfterTimesheetTypeCreateCallback GongOnAfterCreateInterface[TimesheetType]
	OnAfterTimesheetTypeUpdateCallback GongOnAfterUpdateInterface[TimesheetType]
	OnAfterTimesheetTypeDeleteCallback GongOnAfterDeleteInterface[TimesheetType]

	UnitPropertyTypes                map[*UnitPropertyType]struct{}
	UnitPropertyTypes_instance       map[*UnitPropertyType]*UnitPropertyType
	UnitPropertyTypes_mapString      map[string]*UnitPropertyType
	UnitPropertyTypeOrder            uint
	UnitPropertyType_stagedOrder     map[*UnitPropertyType]uint
	UnitPropertyType_orderStaged     map[uint]*UnitPropertyType
	UnitPropertyTypes_reference      map[*UnitPropertyType]*UnitPropertyType
	UnitPropertyTypes_referenceOrder map[*UnitPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterUnitPropertyTypeCreateCallback GongOnAfterCreateInterface[UnitPropertyType]
	OnAfterUnitPropertyTypeUpdateCallback GongOnAfterUpdateInterface[UnitPropertyType]
	OnAfterUnitPropertyTypeDeleteCallback GongOnAfterDeleteInterface[UnitPropertyType]

	VerticalStructureLightingStatusPropertyTypes                map[*VerticalStructureLightingStatusPropertyType]struct{}
	VerticalStructureLightingStatusPropertyTypes_instance       map[*VerticalStructureLightingStatusPropertyType]*VerticalStructureLightingStatusPropertyType
	VerticalStructureLightingStatusPropertyTypes_mapString      map[string]*VerticalStructureLightingStatusPropertyType
	VerticalStructureLightingStatusPropertyTypeOrder            uint
	VerticalStructureLightingStatusPropertyType_stagedOrder     map[*VerticalStructureLightingStatusPropertyType]uint
	VerticalStructureLightingStatusPropertyType_orderStaged     map[uint]*VerticalStructureLightingStatusPropertyType
	VerticalStructureLightingStatusPropertyTypes_reference      map[*VerticalStructureLightingStatusPropertyType]*VerticalStructureLightingStatusPropertyType
	VerticalStructureLightingStatusPropertyTypes_referenceOrder map[*VerticalStructureLightingStatusPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterVerticalStructureLightingStatusPropertyTypeCreateCallback GongOnAfterCreateInterface[VerticalStructureLightingStatusPropertyType]
	OnAfterVerticalStructureLightingStatusPropertyTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructureLightingStatusPropertyType]
	OnAfterVerticalStructureLightingStatusPropertyTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructureLightingStatusPropertyType]

	VerticalStructureLightingStatusTypes                map[*VerticalStructureLightingStatusType]struct{}
	VerticalStructureLightingStatusTypes_instance       map[*VerticalStructureLightingStatusType]*VerticalStructureLightingStatusType
	VerticalStructureLightingStatusTypes_mapString      map[string]*VerticalStructureLightingStatusType
	VerticalStructureLightingStatusTypeOrder            uint
	VerticalStructureLightingStatusType_stagedOrder     map[*VerticalStructureLightingStatusType]uint
	VerticalStructureLightingStatusType_orderStaged     map[uint]*VerticalStructureLightingStatusType
	VerticalStructureLightingStatusTypes_reference      map[*VerticalStructureLightingStatusType]*VerticalStructureLightingStatusType
	VerticalStructureLightingStatusTypes_referenceOrder map[*VerticalStructureLightingStatusType]uint

	// insertion point for slice of pointers maps
	VerticalStructureLightingStatusType_TimeInterval_reverseMap map[*TimesheetPropertyType]*VerticalStructureLightingStatusType

	VerticalStructureLightingStatusType_Annotation_reverseMap map[*NotePropertyType]*VerticalStructureLightingStatusType

	VerticalStructureLightingStatusType_SpecialDateAuthority_reverseMap map[*OrganisationAuthorityPropertyType]*VerticalStructureLightingStatusType

	VerticalStructureLightingStatusType_Extension_reverseMap map[*A_extension_8]*VerticalStructureLightingStatusType

	OnAfterVerticalStructureLightingStatusTypeCreateCallback GongOnAfterCreateInterface[VerticalStructureLightingStatusType]
	OnAfterVerticalStructureLightingStatusTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructureLightingStatusType]
	OnAfterVerticalStructureLightingStatusTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructureLightingStatusType]

	VerticalStructurePartPropertyTypes                map[*VerticalStructurePartPropertyType]struct{}
	VerticalStructurePartPropertyTypes_instance       map[*VerticalStructurePartPropertyType]*VerticalStructurePartPropertyType
	VerticalStructurePartPropertyTypes_mapString      map[string]*VerticalStructurePartPropertyType
	VerticalStructurePartPropertyTypeOrder            uint
	VerticalStructurePartPropertyType_stagedOrder     map[*VerticalStructurePartPropertyType]uint
	VerticalStructurePartPropertyType_orderStaged     map[uint]*VerticalStructurePartPropertyType
	VerticalStructurePartPropertyTypes_reference      map[*VerticalStructurePartPropertyType]*VerticalStructurePartPropertyType
	VerticalStructurePartPropertyTypes_referenceOrder map[*VerticalStructurePartPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterVerticalStructurePartPropertyTypeCreateCallback GongOnAfterCreateInterface[VerticalStructurePartPropertyType]
	OnAfterVerticalStructurePartPropertyTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructurePartPropertyType]
	OnAfterVerticalStructurePartPropertyTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructurePartPropertyType]

	VerticalStructurePartTypes                map[*VerticalStructurePartType]struct{}
	VerticalStructurePartTypes_instance       map[*VerticalStructurePartType]*VerticalStructurePartType
	VerticalStructurePartTypes_mapString      map[string]*VerticalStructurePartType
	VerticalStructurePartTypeOrder            uint
	VerticalStructurePartType_stagedOrder     map[*VerticalStructurePartType]uint
	VerticalStructurePartType_orderStaged     map[uint]*VerticalStructurePartType
	VerticalStructurePartTypes_reference      map[*VerticalStructurePartType]*VerticalStructurePartType
	VerticalStructurePartTypes_referenceOrder map[*VerticalStructurePartType]uint

	// insertion point for slice of pointers maps
	VerticalStructurePartType_TimeInterval_reverseMap map[*TimesheetPropertyType]*VerticalStructurePartType

	VerticalStructurePartType_Annotation_reverseMap map[*NotePropertyType]*VerticalStructurePartType

	VerticalStructurePartType_SpecialDateAuthority_reverseMap map[*OrganisationAuthorityPropertyType]*VerticalStructurePartType

	VerticalStructurePartType_Lighting_reverseMap map[*LightElementPropertyType]*VerticalStructurePartType

	VerticalStructurePartType_Extension_reverseMap map[*A_extension_9]*VerticalStructurePartType

	OnAfterVerticalStructurePartTypeCreateCallback GongOnAfterCreateInterface[VerticalStructurePartType]
	OnAfterVerticalStructurePartTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructurePartType]
	OnAfterVerticalStructurePartTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructurePartType]

	VerticalStructurePropertyTypes                map[*VerticalStructurePropertyType]struct{}
	VerticalStructurePropertyTypes_instance       map[*VerticalStructurePropertyType]*VerticalStructurePropertyType
	VerticalStructurePropertyTypes_mapString      map[string]*VerticalStructurePropertyType
	VerticalStructurePropertyTypeOrder            uint
	VerticalStructurePropertyType_stagedOrder     map[*VerticalStructurePropertyType]uint
	VerticalStructurePropertyType_orderStaged     map[uint]*VerticalStructurePropertyType
	VerticalStructurePropertyTypes_reference      map[*VerticalStructurePropertyType]*VerticalStructurePropertyType
	VerticalStructurePropertyTypes_referenceOrder map[*VerticalStructurePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterVerticalStructurePropertyTypeCreateCallback GongOnAfterCreateInterface[VerticalStructurePropertyType]
	OnAfterVerticalStructurePropertyTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructurePropertyType]
	OnAfterVerticalStructurePropertyTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructurePropertyType]

	VerticalStructureTimeSlicePropertyTypes                map[*VerticalStructureTimeSlicePropertyType]struct{}
	VerticalStructureTimeSlicePropertyTypes_instance       map[*VerticalStructureTimeSlicePropertyType]*VerticalStructureTimeSlicePropertyType
	VerticalStructureTimeSlicePropertyTypes_mapString      map[string]*VerticalStructureTimeSlicePropertyType
	VerticalStructureTimeSlicePropertyTypeOrder            uint
	VerticalStructureTimeSlicePropertyType_stagedOrder     map[*VerticalStructureTimeSlicePropertyType]uint
	VerticalStructureTimeSlicePropertyType_orderStaged     map[uint]*VerticalStructureTimeSlicePropertyType
	VerticalStructureTimeSlicePropertyTypes_reference      map[*VerticalStructureTimeSlicePropertyType]*VerticalStructureTimeSlicePropertyType
	VerticalStructureTimeSlicePropertyTypes_referenceOrder map[*VerticalStructureTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterVerticalStructureTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[VerticalStructureTimeSlicePropertyType]
	OnAfterVerticalStructureTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructureTimeSlicePropertyType]
	OnAfterVerticalStructureTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructureTimeSlicePropertyType]

	VerticalStructureTimeSliceTypes                map[*VerticalStructureTimeSliceType]struct{}
	VerticalStructureTimeSliceTypes_instance       map[*VerticalStructureTimeSliceType]*VerticalStructureTimeSliceType
	VerticalStructureTimeSliceTypes_mapString      map[string]*VerticalStructureTimeSliceType
	VerticalStructureTimeSliceTypeOrder            uint
	VerticalStructureTimeSliceType_stagedOrder     map[*VerticalStructureTimeSliceType]uint
	VerticalStructureTimeSliceType_orderStaged     map[uint]*VerticalStructureTimeSliceType
	VerticalStructureTimeSliceTypes_reference      map[*VerticalStructureTimeSliceType]*VerticalStructureTimeSliceType
	VerticalStructureTimeSliceTypes_referenceOrder map[*VerticalStructureTimeSliceType]uint

	// insertion point for slice of pointers maps
	VerticalStructureTimeSliceType_Part_reverseMap map[*VerticalStructurePartPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_HostedPassengerService_reverseMap map[*PassengerServicePropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_SupportedGroundLight_reverseMap map[*GroundLightSystemPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_HostedNavaidEquipment_reverseMap map[*NavaidEquipmentPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_HostedSpecialNavStation_reverseMap map[*SpecialNavigationStationPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_HostedUnit_reverseMap map[*UnitPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_HostedOrganisation_reverseMap map[*OrganisationAuthorityPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_SupportedService_reverseMap map[*ServicePropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_LightingAvailability_reverseMap map[*VerticalStructureLightingStatusPropertyType]*VerticalStructureTimeSliceType

	VerticalStructureTimeSliceType_Extension_reverseMap map[*A_extension_7]*VerticalStructureTimeSliceType

	OnAfterVerticalStructureTimeSliceTypeCreateCallback GongOnAfterCreateInterface[VerticalStructureTimeSliceType]
	OnAfterVerticalStructureTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructureTimeSliceType]
	OnAfterVerticalStructureTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructureTimeSliceType]

	VerticalStructureTypes                map[*VerticalStructureType]struct{}
	VerticalStructureTypes_instance       map[*VerticalStructureType]*VerticalStructureType
	VerticalStructureTypes_mapString      map[string]*VerticalStructureType
	VerticalStructureTypeOrder            uint
	VerticalStructureType_stagedOrder     map[*VerticalStructureType]uint
	VerticalStructureType_orderStaged     map[uint]*VerticalStructureType
	VerticalStructureTypes_reference      map[*VerticalStructureType]*VerticalStructureType
	VerticalStructureTypes_referenceOrder map[*VerticalStructureType]uint

	// insertion point for slice of pointers maps
	VerticalStructureType_TimeSlice_reverseMap map[*VerticalStructureTimeSlicePropertyType]*VerticalStructureType

	OnAfterVerticalStructureTypeCreateCallback GongOnAfterCreateInterface[VerticalStructureType]
	OnAfterVerticalStructureTypeUpdateCallback GongOnAfterUpdateInterface[VerticalStructureType]
	OnAfterVerticalStructureTypeDeleteCallback GongOnAfterDeleteInterface[VerticalStructureType]

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
	__gong__clearReferences(&stage.A_extensions_reference, &stage.A_extensions_instance, &stage.A_extensions_referenceOrder)

	__gong__clearReferences(&stage.A_extension_1s_reference, &stage.A_extension_1s_instance, &stage.A_extension_1s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_10s_reference, &stage.A_extension_10s_instance, &stage.A_extension_10s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_11s_reference, &stage.A_extension_11s_instance, &stage.A_extension_11s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_12s_reference, &stage.A_extension_12s_instance, &stage.A_extension_12s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_2s_reference, &stage.A_extension_2s_instance, &stage.A_extension_2s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_3s_reference, &stage.A_extension_3s_instance, &stage.A_extension_3s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_4s_reference, &stage.A_extension_4s_instance, &stage.A_extension_4s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_5s_reference, &stage.A_extension_5s_instance, &stage.A_extension_5s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_6s_reference, &stage.A_extension_6s_instance, &stage.A_extension_6s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_7s_reference, &stage.A_extension_7s_instance, &stage.A_extension_7s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_8s_reference, &stage.A_extension_8s_instance, &stage.A_extension_8s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_9s_reference, &stage.A_extension_9s_instance, &stage.A_extension_9s_referenceOrder)

	__gong__clearReferences(&stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_instance, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder)

	__gong__clearReferences(&stage.AeronauticalGroundLightTimeSlicePropertyTypes_reference, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_instance, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AeronauticalGroundLightTimeSliceTypes_reference, &stage.AeronauticalGroundLightTimeSliceTypes_instance, &stage.AeronauticalGroundLightTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.AeronauticalGroundLightTypes_reference, &stage.AeronauticalGroundLightTypes_instance, &stage.AeronauticalGroundLightTypes_referenceOrder)

	__gong__clearReferences(&stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_instance, &stage.AirportHeliportPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.CurveTypes_reference, &stage.CurveTypes_instance, &stage.CurveTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedCurvePropertyTypes_reference, &stage.ElevatedCurvePropertyTypes_instance, &stage.ElevatedCurvePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_instance, &stage.ElevatedCurveTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedPointPropertyTypes_reference, &stage.ElevatedPointPropertyTypes_instance, &stage.ElevatedPointPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_instance, &stage.ElevatedPointTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedSurfacePropertyTypes_reference, &stage.ElevatedSurfacePropertyTypes_instance, &stage.ElevatedSurfacePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_instance, &stage.ElevatedSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.GroundLightSystemPropertyTypes_reference, &stage.GroundLightSystemPropertyTypes_instance, &stage.GroundLightSystemPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.LightElementPropertyTypes_reference, &stage.LightElementPropertyTypes_instance, &stage.LightElementPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.LightElementStatusPropertyTypes_reference, &stage.LightElementStatusPropertyTypes_instance, &stage.LightElementStatusPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.LightElementStatusTypes_reference, &stage.LightElementStatusTypes_instance, &stage.LightElementStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.LightElementTypes_reference, &stage.LightElementTypes_instance, &stage.LightElementTypes_referenceOrder)

	__gong__clearReferences(&stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_instance, &stage.LinguisticNotePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_instance, &stage.LinguisticNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.MarkerBeaconPropertyTypes_reference, &stage.MarkerBeaconPropertyTypes_instance, &stage.MarkerBeaconPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.NavaidEquipmentPropertyTypes_reference, &stage.NavaidEquipmentPropertyTypes_instance, &stage.NavaidEquipmentPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_instance, &stage.NotePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.NoteTypes_reference, &stage.NoteTypes_instance, &stage.NoteTypes_referenceOrder)

	__gong__clearReferences(&stage.ObstacleAreaTimeSlicePropertyTypes_reference, &stage.ObstacleAreaTimeSlicePropertyTypes_instance, &stage.ObstacleAreaTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.ObstacleAreaTimeSliceTypes_reference, &stage.ObstacleAreaTimeSliceTypes_instance, &stage.ObstacleAreaTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.ObstacleAreaTypes_reference, &stage.ObstacleAreaTypes_instance, &stage.ObstacleAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_instance, &stage.OrganisationAuthorityPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.PassengerServicePropertyTypes_reference, &stage.PassengerServicePropertyTypes_instance, &stage.PassengerServicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.PointTypes_reference, &stage.PointTypes_instance, &stage.PointTypes_referenceOrder)

	__gong__clearReferences(&stage.RunwayDirectionPropertyTypes_reference, &stage.RunwayDirectionPropertyTypes_instance, &stage.RunwayDirectionPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.ServicePropertyTypes_reference, &stage.ServicePropertyTypes_instance, &stage.ServicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SpecialNavigationStationPropertyTypes_reference, &stage.SpecialNavigationStationPropertyTypes_instance, &stage.SpecialNavigationStationPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_instance, &stage.SurfacePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfaceTypes_reference, &stage.SurfaceTypes_instance, &stage.SurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_instance, &stage.TimesheetPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.TimesheetTypes_reference, &stage.TimesheetTypes_instance, &stage.TimesheetTypes_referenceOrder)

	__gong__clearReferences(&stage.UnitPropertyTypes_reference, &stage.UnitPropertyTypes_instance, &stage.UnitPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructureLightingStatusPropertyTypes_reference, &stage.VerticalStructureLightingStatusPropertyTypes_instance, &stage.VerticalStructureLightingStatusPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructureLightingStatusTypes_reference, &stage.VerticalStructureLightingStatusTypes_instance, &stage.VerticalStructureLightingStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructurePartPropertyTypes_reference, &stage.VerticalStructurePartPropertyTypes_instance, &stage.VerticalStructurePartPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructurePartTypes_reference, &stage.VerticalStructurePartTypes_instance, &stage.VerticalStructurePartTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructurePropertyTypes_reference, &stage.VerticalStructurePropertyTypes_instance, &stage.VerticalStructurePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructureTimeSlicePropertyTypes_reference, &stage.VerticalStructureTimeSlicePropertyTypes_instance, &stage.VerticalStructureTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructureTimeSliceTypes_reference, &stage.VerticalStructureTimeSliceTypes_instance, &stage.VerticalStructureTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.VerticalStructureTypes_reference, &stage.VerticalStructureTypes_instance, &stage.VerticalStructureTypes_referenceOrder)

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
	stage.A_extensionOrder = __gong__recomputeOrder(stage.A_extension_stagedOrder)

	stage.A_extension_1Order = __gong__recomputeOrder(stage.A_extension_1_stagedOrder)

	stage.A_extension_10Order = __gong__recomputeOrder(stage.A_extension_10_stagedOrder)

	stage.A_extension_11Order = __gong__recomputeOrder(stage.A_extension_11_stagedOrder)

	stage.A_extension_12Order = __gong__recomputeOrder(stage.A_extension_12_stagedOrder)

	stage.A_extension_2Order = __gong__recomputeOrder(stage.A_extension_2_stagedOrder)

	stage.A_extension_3Order = __gong__recomputeOrder(stage.A_extension_3_stagedOrder)

	stage.A_extension_4Order = __gong__recomputeOrder(stage.A_extension_4_stagedOrder)

	stage.A_extension_5Order = __gong__recomputeOrder(stage.A_extension_5_stagedOrder)

	stage.A_extension_6Order = __gong__recomputeOrder(stage.A_extension_6_stagedOrder)

	stage.A_extension_7Order = __gong__recomputeOrder(stage.A_extension_7_stagedOrder)

	stage.A_extension_8Order = __gong__recomputeOrder(stage.A_extension_8_stagedOrder)

	stage.A_extension_9Order = __gong__recomputeOrder(stage.A_extension_9_stagedOrder)

	stage.AbstractPropertiesWithScheduleTypeOrder = __gong__recomputeOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder)

	stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder)

	stage.AeronauticalGroundLightTimeSliceTypeOrder = __gong__recomputeOrder(stage.AeronauticalGroundLightTimeSliceType_stagedOrder)

	stage.AeronauticalGroundLightTypeOrder = __gong__recomputeOrder(stage.AeronauticalGroundLightType_stagedOrder)

	stage.AirportHeliportPropertyTypeOrder = __gong__recomputeOrder(stage.AirportHeliportPropertyType_stagedOrder)

	stage.CurveTypeOrder = __gong__recomputeOrder(stage.CurveType_stagedOrder)

	stage.ElevatedCurvePropertyTypeOrder = __gong__recomputeOrder(stage.ElevatedCurvePropertyType_stagedOrder)

	stage.ElevatedCurveTypeOrder = __gong__recomputeOrder(stage.ElevatedCurveType_stagedOrder)

	stage.ElevatedPointPropertyTypeOrder = __gong__recomputeOrder(stage.ElevatedPointPropertyType_stagedOrder)

	stage.ElevatedPointTypeOrder = __gong__recomputeOrder(stage.ElevatedPointType_stagedOrder)

	stage.ElevatedSurfacePropertyTypeOrder = __gong__recomputeOrder(stage.ElevatedSurfacePropertyType_stagedOrder)

	stage.ElevatedSurfaceTypeOrder = __gong__recomputeOrder(stage.ElevatedSurfaceType_stagedOrder)

	stage.GroundLightSystemPropertyTypeOrder = __gong__recomputeOrder(stage.GroundLightSystemPropertyType_stagedOrder)

	stage.LightElementPropertyTypeOrder = __gong__recomputeOrder(stage.LightElementPropertyType_stagedOrder)

	stage.LightElementStatusPropertyTypeOrder = __gong__recomputeOrder(stage.LightElementStatusPropertyType_stagedOrder)

	stage.LightElementStatusTypeOrder = __gong__recomputeOrder(stage.LightElementStatusType_stagedOrder)

	stage.LightElementTypeOrder = __gong__recomputeOrder(stage.LightElementType_stagedOrder)

	stage.LinguisticNotePropertyTypeOrder = __gong__recomputeOrder(stage.LinguisticNotePropertyType_stagedOrder)

	stage.LinguisticNoteTypeOrder = __gong__recomputeOrder(stage.LinguisticNoteType_stagedOrder)

	stage.MarkerBeaconPropertyTypeOrder = __gong__recomputeOrder(stage.MarkerBeaconPropertyType_stagedOrder)

	stage.NavaidEquipmentPropertyTypeOrder = __gong__recomputeOrder(stage.NavaidEquipmentPropertyType_stagedOrder)

	stage.NotePropertyTypeOrder = __gong__recomputeOrder(stage.NotePropertyType_stagedOrder)

	stage.NoteTypeOrder = __gong__recomputeOrder(stage.NoteType_stagedOrder)

	stage.ObstacleAreaTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.ObstacleAreaTimeSlicePropertyType_stagedOrder)

	stage.ObstacleAreaTimeSliceTypeOrder = __gong__recomputeOrder(stage.ObstacleAreaTimeSliceType_stagedOrder)

	stage.ObstacleAreaTypeOrder = __gong__recomputeOrder(stage.ObstacleAreaType_stagedOrder)

	stage.OrganisationAuthorityPropertyTypeOrder = __gong__recomputeOrder(stage.OrganisationAuthorityPropertyType_stagedOrder)

	stage.PassengerServicePropertyTypeOrder = __gong__recomputeOrder(stage.PassengerServicePropertyType_stagedOrder)

	stage.PointTypeOrder = __gong__recomputeOrder(stage.PointType_stagedOrder)

	stage.RunwayDirectionPropertyTypeOrder = __gong__recomputeOrder(stage.RunwayDirectionPropertyType_stagedOrder)

	stage.ServicePropertyTypeOrder = __gong__recomputeOrder(stage.ServicePropertyType_stagedOrder)

	stage.SpecialNavigationStationPropertyTypeOrder = __gong__recomputeOrder(stage.SpecialNavigationStationPropertyType_stagedOrder)

	stage.SurfacePropertyTypeOrder = __gong__recomputeOrder(stage.SurfacePropertyType_stagedOrder)

	stage.SurfaceTypeOrder = __gong__recomputeOrder(stage.SurfaceType_stagedOrder)

	stage.TimesheetPropertyTypeOrder = __gong__recomputeOrder(stage.TimesheetPropertyType_stagedOrder)

	stage.TimesheetTypeOrder = __gong__recomputeOrder(stage.TimesheetType_stagedOrder)

	stage.UnitPropertyTypeOrder = __gong__recomputeOrder(stage.UnitPropertyType_stagedOrder)

	stage.VerticalStructureLightingStatusPropertyTypeOrder = __gong__recomputeOrder(stage.VerticalStructureLightingStatusPropertyType_stagedOrder)

	stage.VerticalStructureLightingStatusTypeOrder = __gong__recomputeOrder(stage.VerticalStructureLightingStatusType_stagedOrder)

	stage.VerticalStructurePartPropertyTypeOrder = __gong__recomputeOrder(stage.VerticalStructurePartPropertyType_stagedOrder)

	stage.VerticalStructurePartTypeOrder = __gong__recomputeOrder(stage.VerticalStructurePartType_stagedOrder)

	stage.VerticalStructurePropertyTypeOrder = __gong__recomputeOrder(stage.VerticalStructurePropertyType_stagedOrder)

	stage.VerticalStructureTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.VerticalStructureTimeSlicePropertyType_stagedOrder)

	stage.VerticalStructureTimeSliceTypeOrder = __gong__recomputeOrder(stage.VerticalStructureTimeSliceType_stagedOrder)

	stage.VerticalStructureTypeOrder = __gong__recomputeOrder(stage.VerticalStructureType_stagedOrder)

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
func (*A_extension) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extensions, stage.A_extension_stagedOrder)
}

func (*A_extension) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_orderStaged[order]
}

func (*A_extension) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extensions_mapString
}

func (*A_extension) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extensions
}

func (*A_extension) GongNewInstance() any {
	return new(A_extension)
}

func (*A_extension_1) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_1s, stage.A_extension_1_stagedOrder)
}

func (*A_extension_1) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_1_orderStaged[order]
}

func (*A_extension_1) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_1s_mapString
}

func (*A_extension_1) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_1s
}

func (*A_extension_1) GongNewInstance() any {
	return new(A_extension_1)
}

func (*A_extension_10) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_10s, stage.A_extension_10_stagedOrder)
}

func (*A_extension_10) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_10_orderStaged[order]
}

func (*A_extension_10) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_10s_mapString
}

func (*A_extension_10) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_10s
}

func (*A_extension_10) GongNewInstance() any {
	return new(A_extension_10)
}

func (*A_extension_11) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_11s, stage.A_extension_11_stagedOrder)
}

func (*A_extension_11) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_11_orderStaged[order]
}

func (*A_extension_11) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_11s_mapString
}

func (*A_extension_11) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_11s
}

func (*A_extension_11) GongNewInstance() any {
	return new(A_extension_11)
}

func (*A_extension_12) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_12s, stage.A_extension_12_stagedOrder)
}

func (*A_extension_12) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_12_orderStaged[order]
}

func (*A_extension_12) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_12s_mapString
}

func (*A_extension_12) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_12s
}

func (*A_extension_12) GongNewInstance() any {
	return new(A_extension_12)
}

func (*A_extension_2) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_2s, stage.A_extension_2_stagedOrder)
}

func (*A_extension_2) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_2_orderStaged[order]
}

func (*A_extension_2) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_2s_mapString
}

func (*A_extension_2) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_2s
}

func (*A_extension_2) GongNewInstance() any {
	return new(A_extension_2)
}

func (*A_extension_3) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_3s, stage.A_extension_3_stagedOrder)
}

func (*A_extension_3) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_3_orderStaged[order]
}

func (*A_extension_3) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_3s_mapString
}

func (*A_extension_3) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_3s
}

func (*A_extension_3) GongNewInstance() any {
	return new(A_extension_3)
}

func (*A_extension_4) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_4s, stage.A_extension_4_stagedOrder)
}

func (*A_extension_4) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_4_orderStaged[order]
}

func (*A_extension_4) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_4s_mapString
}

func (*A_extension_4) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_4s
}

func (*A_extension_4) GongNewInstance() any {
	return new(A_extension_4)
}

func (*A_extension_5) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_5s, stage.A_extension_5_stagedOrder)
}

func (*A_extension_5) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_5_orderStaged[order]
}

func (*A_extension_5) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_5s_mapString
}

func (*A_extension_5) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_5s
}

func (*A_extension_5) GongNewInstance() any {
	return new(A_extension_5)
}

func (*A_extension_6) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_6s, stage.A_extension_6_stagedOrder)
}

func (*A_extension_6) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_6_orderStaged[order]
}

func (*A_extension_6) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_6s_mapString
}

func (*A_extension_6) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_6s
}

func (*A_extension_6) GongNewInstance() any {
	return new(A_extension_6)
}

func (*A_extension_7) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_7s, stage.A_extension_7_stagedOrder)
}

func (*A_extension_7) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_7_orderStaged[order]
}

func (*A_extension_7) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_7s_mapString
}

func (*A_extension_7) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_7s
}

func (*A_extension_7) GongNewInstance() any {
	return new(A_extension_7)
}

func (*A_extension_8) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_8s, stage.A_extension_8_stagedOrder)
}

func (*A_extension_8) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_8_orderStaged[order]
}

func (*A_extension_8) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_8s_mapString
}

func (*A_extension_8) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_8s
}

func (*A_extension_8) GongNewInstance() any {
	return new(A_extension_8)
}

func (*A_extension_9) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_9s, stage.A_extension_9_stagedOrder)
}

func (*A_extension_9) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_9_orderStaged[order]
}

func (*A_extension_9) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_9s_mapString
}

func (*A_extension_9) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_9s
}

func (*A_extension_9) GongNewInstance() any {
	return new(A_extension_9)
}

func (*AbstractPropertiesWithScheduleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleType_stagedOrder)
}

func (*AbstractPropertiesWithScheduleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AbstractPropertiesWithScheduleType_orderStaged[order]
}

func (*AbstractPropertiesWithScheduleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AbstractPropertiesWithScheduleTypes_mapString
}

func (*AbstractPropertiesWithScheduleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AbstractPropertiesWithScheduleTypes
}

func (*AbstractPropertiesWithScheduleType) GongNewInstance() any {
	return new(AbstractPropertiesWithScheduleType)
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder)
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AeronauticalGroundLightTimeSlicePropertyType_orderStaged[order]
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AeronauticalGroundLightTimeSlicePropertyTypes
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongNewInstance() any {
	return new(AeronauticalGroundLightTimeSlicePropertyType)
}

func (*AeronauticalGroundLightTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceType_stagedOrder)
}

func (*AeronauticalGroundLightTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AeronauticalGroundLightTimeSliceType_orderStaged[order]
}

func (*AeronauticalGroundLightTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AeronauticalGroundLightTimeSliceTypes_mapString
}

func (*AeronauticalGroundLightTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AeronauticalGroundLightTimeSliceTypes
}

func (*AeronauticalGroundLightTimeSliceType) GongNewInstance() any {
	return new(AeronauticalGroundLightTimeSliceType)
}

func (*AeronauticalGroundLightType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightType_stagedOrder)
}

func (*AeronauticalGroundLightType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AeronauticalGroundLightType_orderStaged[order]
}

func (*AeronauticalGroundLightType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AeronauticalGroundLightTypes_mapString
}

func (*AeronauticalGroundLightType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AeronauticalGroundLightTypes
}

func (*AeronauticalGroundLightType) GongNewInstance() any {
	return new(AeronauticalGroundLightType)
}

func (*AirportHeliportPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyType_stagedOrder)
}

func (*AirportHeliportPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirportHeliportPropertyType_orderStaged[order]
}

func (*AirportHeliportPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirportHeliportPropertyTypes_mapString
}

func (*AirportHeliportPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirportHeliportPropertyTypes
}

func (*AirportHeliportPropertyType) GongNewInstance() any {
	return new(AirportHeliportPropertyType)
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

func (*ElevatedCurvePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyType_stagedOrder)
}

func (*ElevatedCurvePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedCurvePropertyType_orderStaged[order]
}

func (*ElevatedCurvePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedCurvePropertyTypes_mapString
}

func (*ElevatedCurvePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedCurvePropertyTypes
}

func (*ElevatedCurvePropertyType) GongNewInstance() any {
	return new(ElevatedCurvePropertyType)
}

func (*ElevatedCurveType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedCurveTypes, stage.ElevatedCurveType_stagedOrder)
}

func (*ElevatedCurveType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedCurveType_orderStaged[order]
}

func (*ElevatedCurveType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedCurveTypes_mapString
}

func (*ElevatedCurveType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedCurveTypes
}

func (*ElevatedCurveType) GongNewInstance() any {
	return new(ElevatedCurveType)
}

func (*ElevatedPointPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyType_stagedOrder)
}

func (*ElevatedPointPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedPointPropertyType_orderStaged[order]
}

func (*ElevatedPointPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedPointPropertyTypes_mapString
}

func (*ElevatedPointPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedPointPropertyTypes
}

func (*ElevatedPointPropertyType) GongNewInstance() any {
	return new(ElevatedPointPropertyType)
}

func (*ElevatedPointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedPointTypes, stage.ElevatedPointType_stagedOrder)
}

func (*ElevatedPointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedPointType_orderStaged[order]
}

func (*ElevatedPointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedPointTypes_mapString
}

func (*ElevatedPointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedPointTypes
}

func (*ElevatedPointType) GongNewInstance() any {
	return new(ElevatedPointType)
}

func (*ElevatedSurfacePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyType_stagedOrder)
}

func (*ElevatedSurfacePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedSurfacePropertyType_orderStaged[order]
}

func (*ElevatedSurfacePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedSurfacePropertyTypes_mapString
}

func (*ElevatedSurfacePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedSurfacePropertyTypes
}

func (*ElevatedSurfacePropertyType) GongNewInstance() any {
	return new(ElevatedSurfacePropertyType)
}

func (*ElevatedSurfaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceType_stagedOrder)
}

func (*ElevatedSurfaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ElevatedSurfaceType_orderStaged[order]
}

func (*ElevatedSurfaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ElevatedSurfaceTypes_mapString
}

func (*ElevatedSurfaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ElevatedSurfaceTypes
}

func (*ElevatedSurfaceType) GongNewInstance() any {
	return new(ElevatedSurfaceType)
}

func (*GroundLightSystemPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyType_stagedOrder)
}

func (*GroundLightSystemPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GroundLightSystemPropertyType_orderStaged[order]
}

func (*GroundLightSystemPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GroundLightSystemPropertyTypes_mapString
}

func (*GroundLightSystemPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GroundLightSystemPropertyTypes
}

func (*GroundLightSystemPropertyType) GongNewInstance() any {
	return new(GroundLightSystemPropertyType)
}

func (*LightElementPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LightElementPropertyTypes, stage.LightElementPropertyType_stagedOrder)
}

func (*LightElementPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LightElementPropertyType_orderStaged[order]
}

func (*LightElementPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LightElementPropertyTypes_mapString
}

func (*LightElementPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LightElementPropertyTypes
}

func (*LightElementPropertyType) GongNewInstance() any {
	return new(LightElementPropertyType)
}

func (*LightElementStatusPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyType_stagedOrder)
}

func (*LightElementStatusPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LightElementStatusPropertyType_orderStaged[order]
}

func (*LightElementStatusPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LightElementStatusPropertyTypes_mapString
}

func (*LightElementStatusPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LightElementStatusPropertyTypes
}

func (*LightElementStatusPropertyType) GongNewInstance() any {
	return new(LightElementStatusPropertyType)
}

func (*LightElementStatusType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LightElementStatusTypes, stage.LightElementStatusType_stagedOrder)
}

func (*LightElementStatusType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LightElementStatusType_orderStaged[order]
}

func (*LightElementStatusType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LightElementStatusTypes_mapString
}

func (*LightElementStatusType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LightElementStatusTypes
}

func (*LightElementStatusType) GongNewInstance() any {
	return new(LightElementStatusType)
}

func (*LightElementType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LightElementTypes, stage.LightElementType_stagedOrder)
}

func (*LightElementType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LightElementType_orderStaged[order]
}

func (*LightElementType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LightElementTypes_mapString
}

func (*LightElementType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LightElementTypes
}

func (*LightElementType) GongNewInstance() any {
	return new(LightElementType)
}

func (*LinguisticNotePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyType_stagedOrder)
}

func (*LinguisticNotePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LinguisticNotePropertyType_orderStaged[order]
}

func (*LinguisticNotePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LinguisticNotePropertyTypes_mapString
}

func (*LinguisticNotePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LinguisticNotePropertyTypes
}

func (*LinguisticNotePropertyType) GongNewInstance() any {
	return new(LinguisticNotePropertyType)
}

func (*LinguisticNoteType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LinguisticNoteTypes, stage.LinguisticNoteType_stagedOrder)
}

func (*LinguisticNoteType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LinguisticNoteType_orderStaged[order]
}

func (*LinguisticNoteType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LinguisticNoteTypes_mapString
}

func (*LinguisticNoteType) GongGetInstancesSet(stage *Stage) any {
	return &stage.LinguisticNoteTypes
}

func (*LinguisticNoteType) GongNewInstance() any {
	return new(LinguisticNoteType)
}

func (*MarkerBeaconPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyType_stagedOrder)
}

func (*MarkerBeaconPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MarkerBeaconPropertyType_orderStaged[order]
}

func (*MarkerBeaconPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MarkerBeaconPropertyTypes_mapString
}

func (*MarkerBeaconPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.MarkerBeaconPropertyTypes
}

func (*MarkerBeaconPropertyType) GongNewInstance() any {
	return new(MarkerBeaconPropertyType)
}

func (*NavaidEquipmentPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyType_stagedOrder)
}

func (*NavaidEquipmentPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NavaidEquipmentPropertyType_orderStaged[order]
}

func (*NavaidEquipmentPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NavaidEquipmentPropertyTypes_mapString
}

func (*NavaidEquipmentPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NavaidEquipmentPropertyTypes
}

func (*NavaidEquipmentPropertyType) GongNewInstance() any {
	return new(NavaidEquipmentPropertyType)
}

func (*NotePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NotePropertyTypes, stage.NotePropertyType_stagedOrder)
}

func (*NotePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NotePropertyType_orderStaged[order]
}

func (*NotePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NotePropertyTypes_mapString
}

func (*NotePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NotePropertyTypes
}

func (*NotePropertyType) GongNewInstance() any {
	return new(NotePropertyType)
}

func (*NoteType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoteTypes, stage.NoteType_stagedOrder)
}

func (*NoteType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoteType_orderStaged[order]
}

func (*NoteType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoteTypes_mapString
}

func (*NoteType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoteTypes
}

func (*NoteType) GongNewInstance() any {
	return new(NoteType)
}

func (*ObstacleAreaTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyType_stagedOrder)
}

func (*ObstacleAreaTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ObstacleAreaTimeSlicePropertyType_orderStaged[order]
}

func (*ObstacleAreaTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ObstacleAreaTimeSlicePropertyTypes_mapString
}

func (*ObstacleAreaTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ObstacleAreaTimeSlicePropertyTypes
}

func (*ObstacleAreaTimeSlicePropertyType) GongNewInstance() any {
	return new(ObstacleAreaTimeSlicePropertyType)
}

func (*ObstacleAreaTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceType_stagedOrder)
}

func (*ObstacleAreaTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ObstacleAreaTimeSliceType_orderStaged[order]
}

func (*ObstacleAreaTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ObstacleAreaTimeSliceTypes_mapString
}

func (*ObstacleAreaTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ObstacleAreaTimeSliceTypes
}

func (*ObstacleAreaTimeSliceType) GongNewInstance() any {
	return new(ObstacleAreaTimeSliceType)
}

func (*ObstacleAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ObstacleAreaTypes, stage.ObstacleAreaType_stagedOrder)
}

func (*ObstacleAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ObstacleAreaType_orderStaged[order]
}

func (*ObstacleAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ObstacleAreaTypes_mapString
}

func (*ObstacleAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ObstacleAreaTypes
}

func (*ObstacleAreaType) GongNewInstance() any {
	return new(ObstacleAreaType)
}

func (*OrganisationAuthorityPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyType_stagedOrder)
}

func (*OrganisationAuthorityPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OrganisationAuthorityPropertyType_orderStaged[order]
}

func (*OrganisationAuthorityPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OrganisationAuthorityPropertyTypes_mapString
}

func (*OrganisationAuthorityPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.OrganisationAuthorityPropertyTypes
}

func (*OrganisationAuthorityPropertyType) GongNewInstance() any {
	return new(OrganisationAuthorityPropertyType)
}

func (*PassengerServicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyType_stagedOrder)
}

func (*PassengerServicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PassengerServicePropertyType_orderStaged[order]
}

func (*PassengerServicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PassengerServicePropertyTypes_mapString
}

func (*PassengerServicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.PassengerServicePropertyTypes
}

func (*PassengerServicePropertyType) GongNewInstance() any {
	return new(PassengerServicePropertyType)
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

func (*RunwayDirectionPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyType_stagedOrder)
}

func (*RunwayDirectionPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.RunwayDirectionPropertyType_orderStaged[order]
}

func (*RunwayDirectionPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.RunwayDirectionPropertyTypes_mapString
}

func (*RunwayDirectionPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.RunwayDirectionPropertyTypes
}

func (*RunwayDirectionPropertyType) GongNewInstance() any {
	return new(RunwayDirectionPropertyType)
}

func (*ServicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ServicePropertyTypes, stage.ServicePropertyType_stagedOrder)
}

func (*ServicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ServicePropertyType_orderStaged[order]
}

func (*ServicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ServicePropertyTypes_mapString
}

func (*ServicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ServicePropertyTypes
}

func (*ServicePropertyType) GongNewInstance() any {
	return new(ServicePropertyType)
}

func (*SpecialNavigationStationPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyType_stagedOrder)
}

func (*SpecialNavigationStationPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SpecialNavigationStationPropertyType_orderStaged[order]
}

func (*SpecialNavigationStationPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SpecialNavigationStationPropertyTypes_mapString
}

func (*SpecialNavigationStationPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.SpecialNavigationStationPropertyTypes
}

func (*SpecialNavigationStationPropertyType) GongNewInstance() any {
	return new(SpecialNavigationStationPropertyType)
}

func (*SurfacePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SurfacePropertyTypes, stage.SurfacePropertyType_stagedOrder)
}

func (*SurfacePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SurfacePropertyType_orderStaged[order]
}

func (*SurfacePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SurfacePropertyTypes_mapString
}

func (*SurfacePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.SurfacePropertyTypes
}

func (*SurfacePropertyType) GongNewInstance() any {
	return new(SurfacePropertyType)
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

func (*TimesheetPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimesheetPropertyTypes, stage.TimesheetPropertyType_stagedOrder)
}

func (*TimesheetPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimesheetPropertyType_orderStaged[order]
}

func (*TimesheetPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimesheetPropertyTypes_mapString
}

func (*TimesheetPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimesheetPropertyTypes
}

func (*TimesheetPropertyType) GongNewInstance() any {
	return new(TimesheetPropertyType)
}

func (*TimesheetType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimesheetTypes, stage.TimesheetType_stagedOrder)
}

func (*TimesheetType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimesheetType_orderStaged[order]
}

func (*TimesheetType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimesheetTypes_mapString
}

func (*TimesheetType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimesheetTypes
}

func (*TimesheetType) GongNewInstance() any {
	return new(TimesheetType)
}

func (*UnitPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.UnitPropertyTypes, stage.UnitPropertyType_stagedOrder)
}

func (*UnitPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.UnitPropertyType_orderStaged[order]
}

func (*UnitPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.UnitPropertyTypes_mapString
}

func (*UnitPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.UnitPropertyTypes
}

func (*UnitPropertyType) GongNewInstance() any {
	return new(UnitPropertyType)
}

func (*VerticalStructureLightingStatusPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyType_stagedOrder)
}

func (*VerticalStructureLightingStatusPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructureLightingStatusPropertyType_orderStaged[order]
}

func (*VerticalStructureLightingStatusPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructureLightingStatusPropertyTypes_mapString
}

func (*VerticalStructureLightingStatusPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructureLightingStatusPropertyTypes
}

func (*VerticalStructureLightingStatusPropertyType) GongNewInstance() any {
	return new(VerticalStructureLightingStatusPropertyType)
}

func (*VerticalStructureLightingStatusType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusType_stagedOrder)
}

func (*VerticalStructureLightingStatusType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructureLightingStatusType_orderStaged[order]
}

func (*VerticalStructureLightingStatusType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructureLightingStatusTypes_mapString
}

func (*VerticalStructureLightingStatusType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructureLightingStatusTypes
}

func (*VerticalStructureLightingStatusType) GongNewInstance() any {
	return new(VerticalStructureLightingStatusType)
}

func (*VerticalStructurePartPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyType_stagedOrder)
}

func (*VerticalStructurePartPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructurePartPropertyType_orderStaged[order]
}

func (*VerticalStructurePartPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructurePartPropertyTypes_mapString
}

func (*VerticalStructurePartPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructurePartPropertyTypes
}

func (*VerticalStructurePartPropertyType) GongNewInstance() any {
	return new(VerticalStructurePartPropertyType)
}

func (*VerticalStructurePartType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructurePartTypes, stage.VerticalStructurePartType_stagedOrder)
}

func (*VerticalStructurePartType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructurePartType_orderStaged[order]
}

func (*VerticalStructurePartType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructurePartTypes_mapString
}

func (*VerticalStructurePartType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructurePartTypes
}

func (*VerticalStructurePartType) GongNewInstance() any {
	return new(VerticalStructurePartType)
}

func (*VerticalStructurePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyType_stagedOrder)
}

func (*VerticalStructurePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructurePropertyType_orderStaged[order]
}

func (*VerticalStructurePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructurePropertyTypes_mapString
}

func (*VerticalStructurePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructurePropertyTypes
}

func (*VerticalStructurePropertyType) GongNewInstance() any {
	return new(VerticalStructurePropertyType)
}

func (*VerticalStructureTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyType_stagedOrder)
}

func (*VerticalStructureTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructureTimeSlicePropertyType_orderStaged[order]
}

func (*VerticalStructureTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructureTimeSlicePropertyTypes_mapString
}

func (*VerticalStructureTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructureTimeSlicePropertyTypes
}

func (*VerticalStructureTimeSlicePropertyType) GongNewInstance() any {
	return new(VerticalStructureTimeSlicePropertyType)
}

func (*VerticalStructureTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceType_stagedOrder)
}

func (*VerticalStructureTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructureTimeSliceType_orderStaged[order]
}

func (*VerticalStructureTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructureTimeSliceTypes_mapString
}

func (*VerticalStructureTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructureTimeSliceTypes
}

func (*VerticalStructureTimeSliceType) GongNewInstance() any {
	return new(VerticalStructureTimeSliceType)
}

func (*VerticalStructureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VerticalStructureTypes, stage.VerticalStructureType_stagedOrder)
}

func (*VerticalStructureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VerticalStructureType_orderStaged[order]
}

func (*VerticalStructureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VerticalStructureTypes_mapString
}

func (*VerticalStructureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.VerticalStructureTypes
}

func (*VerticalStructureType) GongNewInstance() any {
	return new(VerticalStructureType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
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
		A_extensions:           make(map[*A_extension]struct{}),
		A_extensions_mapString: make(map[string]*A_extension),

		A_extension_1s:           make(map[*A_extension_1]struct{}),
		A_extension_1s_mapString: make(map[string]*A_extension_1),

		A_extension_10s:           make(map[*A_extension_10]struct{}),
		A_extension_10s_mapString: make(map[string]*A_extension_10),

		A_extension_11s:           make(map[*A_extension_11]struct{}),
		A_extension_11s_mapString: make(map[string]*A_extension_11),

		A_extension_12s:           make(map[*A_extension_12]struct{}),
		A_extension_12s_mapString: make(map[string]*A_extension_12),

		A_extension_2s:           make(map[*A_extension_2]struct{}),
		A_extension_2s_mapString: make(map[string]*A_extension_2),

		A_extension_3s:           make(map[*A_extension_3]struct{}),
		A_extension_3s_mapString: make(map[string]*A_extension_3),

		A_extension_4s:           make(map[*A_extension_4]struct{}),
		A_extension_4s_mapString: make(map[string]*A_extension_4),

		A_extension_5s:           make(map[*A_extension_5]struct{}),
		A_extension_5s_mapString: make(map[string]*A_extension_5),

		A_extension_6s:           make(map[*A_extension_6]struct{}),
		A_extension_6s_mapString: make(map[string]*A_extension_6),

		A_extension_7s:           make(map[*A_extension_7]struct{}),
		A_extension_7s_mapString: make(map[string]*A_extension_7),

		A_extension_8s:           make(map[*A_extension_8]struct{}),
		A_extension_8s_mapString: make(map[string]*A_extension_8),

		A_extension_9s:           make(map[*A_extension_9]struct{}),
		A_extension_9s_mapString: make(map[string]*A_extension_9),

		AbstractPropertiesWithScheduleTypes:           make(map[*AbstractPropertiesWithScheduleType]struct{}),
		AbstractPropertiesWithScheduleTypes_mapString: make(map[string]*AbstractPropertiesWithScheduleType),

		AeronauticalGroundLightTimeSlicePropertyTypes:           make(map[*AeronauticalGroundLightTimeSlicePropertyType]struct{}),
		AeronauticalGroundLightTimeSlicePropertyTypes_mapString: make(map[string]*AeronauticalGroundLightTimeSlicePropertyType),

		AeronauticalGroundLightTimeSliceTypes:           make(map[*AeronauticalGroundLightTimeSliceType]struct{}),
		AeronauticalGroundLightTimeSliceTypes_mapString: make(map[string]*AeronauticalGroundLightTimeSliceType),

		AeronauticalGroundLightTypes:           make(map[*AeronauticalGroundLightType]struct{}),
		AeronauticalGroundLightTypes_mapString: make(map[string]*AeronauticalGroundLightType),

		AirportHeliportPropertyTypes:           make(map[*AirportHeliportPropertyType]struct{}),
		AirportHeliportPropertyTypes_mapString: make(map[string]*AirportHeliportPropertyType),

		CurveTypes:           make(map[*CurveType]struct{}),
		CurveTypes_mapString: make(map[string]*CurveType),

		ElevatedCurvePropertyTypes:           make(map[*ElevatedCurvePropertyType]struct{}),
		ElevatedCurvePropertyTypes_mapString: make(map[string]*ElevatedCurvePropertyType),

		ElevatedCurveTypes:           make(map[*ElevatedCurveType]struct{}),
		ElevatedCurveTypes_mapString: make(map[string]*ElevatedCurveType),

		ElevatedPointPropertyTypes:           make(map[*ElevatedPointPropertyType]struct{}),
		ElevatedPointPropertyTypes_mapString: make(map[string]*ElevatedPointPropertyType),

		ElevatedPointTypes:           make(map[*ElevatedPointType]struct{}),
		ElevatedPointTypes_mapString: make(map[string]*ElevatedPointType),

		ElevatedSurfacePropertyTypes:           make(map[*ElevatedSurfacePropertyType]struct{}),
		ElevatedSurfacePropertyTypes_mapString: make(map[string]*ElevatedSurfacePropertyType),

		ElevatedSurfaceTypes:           make(map[*ElevatedSurfaceType]struct{}),
		ElevatedSurfaceTypes_mapString: make(map[string]*ElevatedSurfaceType),

		GroundLightSystemPropertyTypes:           make(map[*GroundLightSystemPropertyType]struct{}),
		GroundLightSystemPropertyTypes_mapString: make(map[string]*GroundLightSystemPropertyType),

		LightElementPropertyTypes:           make(map[*LightElementPropertyType]struct{}),
		LightElementPropertyTypes_mapString: make(map[string]*LightElementPropertyType),

		LightElementStatusPropertyTypes:           make(map[*LightElementStatusPropertyType]struct{}),
		LightElementStatusPropertyTypes_mapString: make(map[string]*LightElementStatusPropertyType),

		LightElementStatusTypes:           make(map[*LightElementStatusType]struct{}),
		LightElementStatusTypes_mapString: make(map[string]*LightElementStatusType),

		LightElementTypes:           make(map[*LightElementType]struct{}),
		LightElementTypes_mapString: make(map[string]*LightElementType),

		LinguisticNotePropertyTypes:           make(map[*LinguisticNotePropertyType]struct{}),
		LinguisticNotePropertyTypes_mapString: make(map[string]*LinguisticNotePropertyType),

		LinguisticNoteTypes:           make(map[*LinguisticNoteType]struct{}),
		LinguisticNoteTypes_mapString: make(map[string]*LinguisticNoteType),

		MarkerBeaconPropertyTypes:           make(map[*MarkerBeaconPropertyType]struct{}),
		MarkerBeaconPropertyTypes_mapString: make(map[string]*MarkerBeaconPropertyType),

		NavaidEquipmentPropertyTypes:           make(map[*NavaidEquipmentPropertyType]struct{}),
		NavaidEquipmentPropertyTypes_mapString: make(map[string]*NavaidEquipmentPropertyType),

		NotePropertyTypes:           make(map[*NotePropertyType]struct{}),
		NotePropertyTypes_mapString: make(map[string]*NotePropertyType),

		NoteTypes:           make(map[*NoteType]struct{}),
		NoteTypes_mapString: make(map[string]*NoteType),

		ObstacleAreaTimeSlicePropertyTypes:           make(map[*ObstacleAreaTimeSlicePropertyType]struct{}),
		ObstacleAreaTimeSlicePropertyTypes_mapString: make(map[string]*ObstacleAreaTimeSlicePropertyType),

		ObstacleAreaTimeSliceTypes:           make(map[*ObstacleAreaTimeSliceType]struct{}),
		ObstacleAreaTimeSliceTypes_mapString: make(map[string]*ObstacleAreaTimeSliceType),

		ObstacleAreaTypes:           make(map[*ObstacleAreaType]struct{}),
		ObstacleAreaTypes_mapString: make(map[string]*ObstacleAreaType),

		OrganisationAuthorityPropertyTypes:           make(map[*OrganisationAuthorityPropertyType]struct{}),
		OrganisationAuthorityPropertyTypes_mapString: make(map[string]*OrganisationAuthorityPropertyType),

		PassengerServicePropertyTypes:           make(map[*PassengerServicePropertyType]struct{}),
		PassengerServicePropertyTypes_mapString: make(map[string]*PassengerServicePropertyType),

		PointTypes:           make(map[*PointType]struct{}),
		PointTypes_mapString: make(map[string]*PointType),

		RunwayDirectionPropertyTypes:           make(map[*RunwayDirectionPropertyType]struct{}),
		RunwayDirectionPropertyTypes_mapString: make(map[string]*RunwayDirectionPropertyType),

		ServicePropertyTypes:           make(map[*ServicePropertyType]struct{}),
		ServicePropertyTypes_mapString: make(map[string]*ServicePropertyType),

		SpecialNavigationStationPropertyTypes:           make(map[*SpecialNavigationStationPropertyType]struct{}),
		SpecialNavigationStationPropertyTypes_mapString: make(map[string]*SpecialNavigationStationPropertyType),

		SurfacePropertyTypes:           make(map[*SurfacePropertyType]struct{}),
		SurfacePropertyTypes_mapString: make(map[string]*SurfacePropertyType),

		SurfaceTypes:           make(map[*SurfaceType]struct{}),
		SurfaceTypes_mapString: make(map[string]*SurfaceType),

		TimesheetPropertyTypes:           make(map[*TimesheetPropertyType]struct{}),
		TimesheetPropertyTypes_mapString: make(map[string]*TimesheetPropertyType),

		TimesheetTypes:           make(map[*TimesheetType]struct{}),
		TimesheetTypes_mapString: make(map[string]*TimesheetType),

		UnitPropertyTypes:           make(map[*UnitPropertyType]struct{}),
		UnitPropertyTypes_mapString: make(map[string]*UnitPropertyType),

		VerticalStructureLightingStatusPropertyTypes:           make(map[*VerticalStructureLightingStatusPropertyType]struct{}),
		VerticalStructureLightingStatusPropertyTypes_mapString: make(map[string]*VerticalStructureLightingStatusPropertyType),

		VerticalStructureLightingStatusTypes:           make(map[*VerticalStructureLightingStatusType]struct{}),
		VerticalStructureLightingStatusTypes_mapString: make(map[string]*VerticalStructureLightingStatusType),

		VerticalStructurePartPropertyTypes:           make(map[*VerticalStructurePartPropertyType]struct{}),
		VerticalStructurePartPropertyTypes_mapString: make(map[string]*VerticalStructurePartPropertyType),

		VerticalStructurePartTypes:           make(map[*VerticalStructurePartType]struct{}),
		VerticalStructurePartTypes_mapString: make(map[string]*VerticalStructurePartType),

		VerticalStructurePropertyTypes:           make(map[*VerticalStructurePropertyType]struct{}),
		VerticalStructurePropertyTypes_mapString: make(map[string]*VerticalStructurePropertyType),

		VerticalStructureTimeSlicePropertyTypes:           make(map[*VerticalStructureTimeSlicePropertyType]struct{}),
		VerticalStructureTimeSlicePropertyTypes_mapString: make(map[string]*VerticalStructureTimeSlicePropertyType),

		VerticalStructureTimeSliceTypes:           make(map[*VerticalStructureTimeSliceType]struct{}),
		VerticalStructureTimeSliceTypes_mapString: make(map[string]*VerticalStructureTimeSliceType),

		VerticalStructureTypes:           make(map[*VerticalStructureType]struct{}),
		VerticalStructureTypes_mapString: make(map[string]*VerticalStructureType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		A_extension_stagedOrder: make(map[*A_extension]uint),
		A_extension_orderStaged: make(map[uint]*A_extension),
		A_extensions_reference:  make(map[*A_extension]*A_extension),

		A_extension_1_stagedOrder: make(map[*A_extension_1]uint),
		A_extension_1_orderStaged: make(map[uint]*A_extension_1),
		A_extension_1s_reference:  make(map[*A_extension_1]*A_extension_1),

		A_extension_10_stagedOrder: make(map[*A_extension_10]uint),
		A_extension_10_orderStaged: make(map[uint]*A_extension_10),
		A_extension_10s_reference:  make(map[*A_extension_10]*A_extension_10),

		A_extension_11_stagedOrder: make(map[*A_extension_11]uint),
		A_extension_11_orderStaged: make(map[uint]*A_extension_11),
		A_extension_11s_reference:  make(map[*A_extension_11]*A_extension_11),

		A_extension_12_stagedOrder: make(map[*A_extension_12]uint),
		A_extension_12_orderStaged: make(map[uint]*A_extension_12),
		A_extension_12s_reference:  make(map[*A_extension_12]*A_extension_12),

		A_extension_2_stagedOrder: make(map[*A_extension_2]uint),
		A_extension_2_orderStaged: make(map[uint]*A_extension_2),
		A_extension_2s_reference:  make(map[*A_extension_2]*A_extension_2),

		A_extension_3_stagedOrder: make(map[*A_extension_3]uint),
		A_extension_3_orderStaged: make(map[uint]*A_extension_3),
		A_extension_3s_reference:  make(map[*A_extension_3]*A_extension_3),

		A_extension_4_stagedOrder: make(map[*A_extension_4]uint),
		A_extension_4_orderStaged: make(map[uint]*A_extension_4),
		A_extension_4s_reference:  make(map[*A_extension_4]*A_extension_4),

		A_extension_5_stagedOrder: make(map[*A_extension_5]uint),
		A_extension_5_orderStaged: make(map[uint]*A_extension_5),
		A_extension_5s_reference:  make(map[*A_extension_5]*A_extension_5),

		A_extension_6_stagedOrder: make(map[*A_extension_6]uint),
		A_extension_6_orderStaged: make(map[uint]*A_extension_6),
		A_extension_6s_reference:  make(map[*A_extension_6]*A_extension_6),

		A_extension_7_stagedOrder: make(map[*A_extension_7]uint),
		A_extension_7_orderStaged: make(map[uint]*A_extension_7),
		A_extension_7s_reference:  make(map[*A_extension_7]*A_extension_7),

		A_extension_8_stagedOrder: make(map[*A_extension_8]uint),
		A_extension_8_orderStaged: make(map[uint]*A_extension_8),
		A_extension_8s_reference:  make(map[*A_extension_8]*A_extension_8),

		A_extension_9_stagedOrder: make(map[*A_extension_9]uint),
		A_extension_9_orderStaged: make(map[uint]*A_extension_9),
		A_extension_9s_reference:  make(map[*A_extension_9]*A_extension_9),

		AbstractPropertiesWithScheduleType_stagedOrder: make(map[*AbstractPropertiesWithScheduleType]uint),
		AbstractPropertiesWithScheduleType_orderStaged: make(map[uint]*AbstractPropertiesWithScheduleType),
		AbstractPropertiesWithScheduleTypes_reference:  make(map[*AbstractPropertiesWithScheduleType]*AbstractPropertiesWithScheduleType),

		AeronauticalGroundLightTimeSlicePropertyType_stagedOrder: make(map[*AeronauticalGroundLightTimeSlicePropertyType]uint),
		AeronauticalGroundLightTimeSlicePropertyType_orderStaged: make(map[uint]*AeronauticalGroundLightTimeSlicePropertyType),
		AeronauticalGroundLightTimeSlicePropertyTypes_reference:  make(map[*AeronauticalGroundLightTimeSlicePropertyType]*AeronauticalGroundLightTimeSlicePropertyType),

		AeronauticalGroundLightTimeSliceType_stagedOrder: make(map[*AeronauticalGroundLightTimeSliceType]uint),
		AeronauticalGroundLightTimeSliceType_orderStaged: make(map[uint]*AeronauticalGroundLightTimeSliceType),
		AeronauticalGroundLightTimeSliceTypes_reference:  make(map[*AeronauticalGroundLightTimeSliceType]*AeronauticalGroundLightTimeSliceType),

		AeronauticalGroundLightType_stagedOrder: make(map[*AeronauticalGroundLightType]uint),
		AeronauticalGroundLightType_orderStaged: make(map[uint]*AeronauticalGroundLightType),
		AeronauticalGroundLightTypes_reference:  make(map[*AeronauticalGroundLightType]*AeronauticalGroundLightType),

		AirportHeliportPropertyType_stagedOrder: make(map[*AirportHeliportPropertyType]uint),
		AirportHeliportPropertyType_orderStaged: make(map[uint]*AirportHeliportPropertyType),
		AirportHeliportPropertyTypes_reference:  make(map[*AirportHeliportPropertyType]*AirportHeliportPropertyType),

		CurveType_stagedOrder: make(map[*CurveType]uint),
		CurveType_orderStaged: make(map[uint]*CurveType),
		CurveTypes_reference:  make(map[*CurveType]*CurveType),

		ElevatedCurvePropertyType_stagedOrder: make(map[*ElevatedCurvePropertyType]uint),
		ElevatedCurvePropertyType_orderStaged: make(map[uint]*ElevatedCurvePropertyType),
		ElevatedCurvePropertyTypes_reference:  make(map[*ElevatedCurvePropertyType]*ElevatedCurvePropertyType),

		ElevatedCurveType_stagedOrder: make(map[*ElevatedCurveType]uint),
		ElevatedCurveType_orderStaged: make(map[uint]*ElevatedCurveType),
		ElevatedCurveTypes_reference:  make(map[*ElevatedCurveType]*ElevatedCurveType),

		ElevatedPointPropertyType_stagedOrder: make(map[*ElevatedPointPropertyType]uint),
		ElevatedPointPropertyType_orderStaged: make(map[uint]*ElevatedPointPropertyType),
		ElevatedPointPropertyTypes_reference:  make(map[*ElevatedPointPropertyType]*ElevatedPointPropertyType),

		ElevatedPointType_stagedOrder: make(map[*ElevatedPointType]uint),
		ElevatedPointType_orderStaged: make(map[uint]*ElevatedPointType),
		ElevatedPointTypes_reference:  make(map[*ElevatedPointType]*ElevatedPointType),

		ElevatedSurfacePropertyType_stagedOrder: make(map[*ElevatedSurfacePropertyType]uint),
		ElevatedSurfacePropertyType_orderStaged: make(map[uint]*ElevatedSurfacePropertyType),
		ElevatedSurfacePropertyTypes_reference:  make(map[*ElevatedSurfacePropertyType]*ElevatedSurfacePropertyType),

		ElevatedSurfaceType_stagedOrder: make(map[*ElevatedSurfaceType]uint),
		ElevatedSurfaceType_orderStaged: make(map[uint]*ElevatedSurfaceType),
		ElevatedSurfaceTypes_reference:  make(map[*ElevatedSurfaceType]*ElevatedSurfaceType),

		GroundLightSystemPropertyType_stagedOrder: make(map[*GroundLightSystemPropertyType]uint),
		GroundLightSystemPropertyType_orderStaged: make(map[uint]*GroundLightSystemPropertyType),
		GroundLightSystemPropertyTypes_reference:  make(map[*GroundLightSystemPropertyType]*GroundLightSystemPropertyType),

		LightElementPropertyType_stagedOrder: make(map[*LightElementPropertyType]uint),
		LightElementPropertyType_orderStaged: make(map[uint]*LightElementPropertyType),
		LightElementPropertyTypes_reference:  make(map[*LightElementPropertyType]*LightElementPropertyType),

		LightElementStatusPropertyType_stagedOrder: make(map[*LightElementStatusPropertyType]uint),
		LightElementStatusPropertyType_orderStaged: make(map[uint]*LightElementStatusPropertyType),
		LightElementStatusPropertyTypes_reference:  make(map[*LightElementStatusPropertyType]*LightElementStatusPropertyType),

		LightElementStatusType_stagedOrder: make(map[*LightElementStatusType]uint),
		LightElementStatusType_orderStaged: make(map[uint]*LightElementStatusType),
		LightElementStatusTypes_reference:  make(map[*LightElementStatusType]*LightElementStatusType),

		LightElementType_stagedOrder: make(map[*LightElementType]uint),
		LightElementType_orderStaged: make(map[uint]*LightElementType),
		LightElementTypes_reference:  make(map[*LightElementType]*LightElementType),

		LinguisticNotePropertyType_stagedOrder: make(map[*LinguisticNotePropertyType]uint),
		LinguisticNotePropertyType_orderStaged: make(map[uint]*LinguisticNotePropertyType),
		LinguisticNotePropertyTypes_reference:  make(map[*LinguisticNotePropertyType]*LinguisticNotePropertyType),

		LinguisticNoteType_stagedOrder: make(map[*LinguisticNoteType]uint),
		LinguisticNoteType_orderStaged: make(map[uint]*LinguisticNoteType),
		LinguisticNoteTypes_reference:  make(map[*LinguisticNoteType]*LinguisticNoteType),

		MarkerBeaconPropertyType_stagedOrder: make(map[*MarkerBeaconPropertyType]uint),
		MarkerBeaconPropertyType_orderStaged: make(map[uint]*MarkerBeaconPropertyType),
		MarkerBeaconPropertyTypes_reference:  make(map[*MarkerBeaconPropertyType]*MarkerBeaconPropertyType),

		NavaidEquipmentPropertyType_stagedOrder: make(map[*NavaidEquipmentPropertyType]uint),
		NavaidEquipmentPropertyType_orderStaged: make(map[uint]*NavaidEquipmentPropertyType),
		NavaidEquipmentPropertyTypes_reference:  make(map[*NavaidEquipmentPropertyType]*NavaidEquipmentPropertyType),

		NotePropertyType_stagedOrder: make(map[*NotePropertyType]uint),
		NotePropertyType_orderStaged: make(map[uint]*NotePropertyType),
		NotePropertyTypes_reference:  make(map[*NotePropertyType]*NotePropertyType),

		NoteType_stagedOrder: make(map[*NoteType]uint),
		NoteType_orderStaged: make(map[uint]*NoteType),
		NoteTypes_reference:  make(map[*NoteType]*NoteType),

		ObstacleAreaTimeSlicePropertyType_stagedOrder: make(map[*ObstacleAreaTimeSlicePropertyType]uint),
		ObstacleAreaTimeSlicePropertyType_orderStaged: make(map[uint]*ObstacleAreaTimeSlicePropertyType),
		ObstacleAreaTimeSlicePropertyTypes_reference:  make(map[*ObstacleAreaTimeSlicePropertyType]*ObstacleAreaTimeSlicePropertyType),

		ObstacleAreaTimeSliceType_stagedOrder: make(map[*ObstacleAreaTimeSliceType]uint),
		ObstacleAreaTimeSliceType_orderStaged: make(map[uint]*ObstacleAreaTimeSliceType),
		ObstacleAreaTimeSliceTypes_reference:  make(map[*ObstacleAreaTimeSliceType]*ObstacleAreaTimeSliceType),

		ObstacleAreaType_stagedOrder: make(map[*ObstacleAreaType]uint),
		ObstacleAreaType_orderStaged: make(map[uint]*ObstacleAreaType),
		ObstacleAreaTypes_reference:  make(map[*ObstacleAreaType]*ObstacleAreaType),

		OrganisationAuthorityPropertyType_stagedOrder: make(map[*OrganisationAuthorityPropertyType]uint),
		OrganisationAuthorityPropertyType_orderStaged: make(map[uint]*OrganisationAuthorityPropertyType),
		OrganisationAuthorityPropertyTypes_reference:  make(map[*OrganisationAuthorityPropertyType]*OrganisationAuthorityPropertyType),

		PassengerServicePropertyType_stagedOrder: make(map[*PassengerServicePropertyType]uint),
		PassengerServicePropertyType_orderStaged: make(map[uint]*PassengerServicePropertyType),
		PassengerServicePropertyTypes_reference:  make(map[*PassengerServicePropertyType]*PassengerServicePropertyType),

		PointType_stagedOrder: make(map[*PointType]uint),
		PointType_orderStaged: make(map[uint]*PointType),
		PointTypes_reference:  make(map[*PointType]*PointType),

		RunwayDirectionPropertyType_stagedOrder: make(map[*RunwayDirectionPropertyType]uint),
		RunwayDirectionPropertyType_orderStaged: make(map[uint]*RunwayDirectionPropertyType),
		RunwayDirectionPropertyTypes_reference:  make(map[*RunwayDirectionPropertyType]*RunwayDirectionPropertyType),

		ServicePropertyType_stagedOrder: make(map[*ServicePropertyType]uint),
		ServicePropertyType_orderStaged: make(map[uint]*ServicePropertyType),
		ServicePropertyTypes_reference:  make(map[*ServicePropertyType]*ServicePropertyType),

		SpecialNavigationStationPropertyType_stagedOrder: make(map[*SpecialNavigationStationPropertyType]uint),
		SpecialNavigationStationPropertyType_orderStaged: make(map[uint]*SpecialNavigationStationPropertyType),
		SpecialNavigationStationPropertyTypes_reference:  make(map[*SpecialNavigationStationPropertyType]*SpecialNavigationStationPropertyType),

		SurfacePropertyType_stagedOrder: make(map[*SurfacePropertyType]uint),
		SurfacePropertyType_orderStaged: make(map[uint]*SurfacePropertyType),
		SurfacePropertyTypes_reference:  make(map[*SurfacePropertyType]*SurfacePropertyType),

		SurfaceType_stagedOrder: make(map[*SurfaceType]uint),
		SurfaceType_orderStaged: make(map[uint]*SurfaceType),
		SurfaceTypes_reference:  make(map[*SurfaceType]*SurfaceType),

		TimesheetPropertyType_stagedOrder: make(map[*TimesheetPropertyType]uint),
		TimesheetPropertyType_orderStaged: make(map[uint]*TimesheetPropertyType),
		TimesheetPropertyTypes_reference:  make(map[*TimesheetPropertyType]*TimesheetPropertyType),

		TimesheetType_stagedOrder: make(map[*TimesheetType]uint),
		TimesheetType_orderStaged: make(map[uint]*TimesheetType),
		TimesheetTypes_reference:  make(map[*TimesheetType]*TimesheetType),

		UnitPropertyType_stagedOrder: make(map[*UnitPropertyType]uint),
		UnitPropertyType_orderStaged: make(map[uint]*UnitPropertyType),
		UnitPropertyTypes_reference:  make(map[*UnitPropertyType]*UnitPropertyType),

		VerticalStructureLightingStatusPropertyType_stagedOrder: make(map[*VerticalStructureLightingStatusPropertyType]uint),
		VerticalStructureLightingStatusPropertyType_orderStaged: make(map[uint]*VerticalStructureLightingStatusPropertyType),
		VerticalStructureLightingStatusPropertyTypes_reference:  make(map[*VerticalStructureLightingStatusPropertyType]*VerticalStructureLightingStatusPropertyType),

		VerticalStructureLightingStatusType_stagedOrder: make(map[*VerticalStructureLightingStatusType]uint),
		VerticalStructureLightingStatusType_orderStaged: make(map[uint]*VerticalStructureLightingStatusType),
		VerticalStructureLightingStatusTypes_reference:  make(map[*VerticalStructureLightingStatusType]*VerticalStructureLightingStatusType),

		VerticalStructurePartPropertyType_stagedOrder: make(map[*VerticalStructurePartPropertyType]uint),
		VerticalStructurePartPropertyType_orderStaged: make(map[uint]*VerticalStructurePartPropertyType),
		VerticalStructurePartPropertyTypes_reference:  make(map[*VerticalStructurePartPropertyType]*VerticalStructurePartPropertyType),

		VerticalStructurePartType_stagedOrder: make(map[*VerticalStructurePartType]uint),
		VerticalStructurePartType_orderStaged: make(map[uint]*VerticalStructurePartType),
		VerticalStructurePartTypes_reference:  make(map[*VerticalStructurePartType]*VerticalStructurePartType),

		VerticalStructurePropertyType_stagedOrder: make(map[*VerticalStructurePropertyType]uint),
		VerticalStructurePropertyType_orderStaged: make(map[uint]*VerticalStructurePropertyType),
		VerticalStructurePropertyTypes_reference:  make(map[*VerticalStructurePropertyType]*VerticalStructurePropertyType),

		VerticalStructureTimeSlicePropertyType_stagedOrder: make(map[*VerticalStructureTimeSlicePropertyType]uint),
		VerticalStructureTimeSlicePropertyType_orderStaged: make(map[uint]*VerticalStructureTimeSlicePropertyType),
		VerticalStructureTimeSlicePropertyTypes_reference:  make(map[*VerticalStructureTimeSlicePropertyType]*VerticalStructureTimeSlicePropertyType),

		VerticalStructureTimeSliceType_stagedOrder: make(map[*VerticalStructureTimeSliceType]uint),
		VerticalStructureTimeSliceType_orderStaged: make(map[uint]*VerticalStructureTimeSliceType),
		VerticalStructureTimeSliceTypes_reference:  make(map[*VerticalStructureTimeSliceType]*VerticalStructureTimeSliceType),

		VerticalStructureType_stagedOrder: make(map[*VerticalStructureType]uint),
		VerticalStructureType_orderStaged: make(map[uint]*VerticalStructureType),
		VerticalStructureTypes_reference:  make(map[*VerticalStructureType]*VerticalStructureType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"A_extension": &A_extensionUnmarshaller{},

			"A_extension_1": &A_extension_1Unmarshaller{},

			"A_extension_10": &A_extension_10Unmarshaller{},

			"A_extension_11": &A_extension_11Unmarshaller{},

			"A_extension_12": &A_extension_12Unmarshaller{},

			"A_extension_2": &A_extension_2Unmarshaller{},

			"A_extension_3": &A_extension_3Unmarshaller{},

			"A_extension_4": &A_extension_4Unmarshaller{},

			"A_extension_5": &A_extension_5Unmarshaller{},

			"A_extension_6": &A_extension_6Unmarshaller{},

			"A_extension_7": &A_extension_7Unmarshaller{},

			"A_extension_8": &A_extension_8Unmarshaller{},

			"A_extension_9": &A_extension_9Unmarshaller{},

			"AbstractPropertiesWithScheduleType": &AbstractPropertiesWithScheduleTypeUnmarshaller{},

			"AeronauticalGroundLightTimeSlicePropertyType": &AeronauticalGroundLightTimeSlicePropertyTypeUnmarshaller{},

			"AeronauticalGroundLightTimeSliceType": &AeronauticalGroundLightTimeSliceTypeUnmarshaller{},

			"AeronauticalGroundLightType": &AeronauticalGroundLightTypeUnmarshaller{},

			"AirportHeliportPropertyType": &AirportHeliportPropertyTypeUnmarshaller{},

			"CurveType": &CurveTypeUnmarshaller{},

			"ElevatedCurvePropertyType": &ElevatedCurvePropertyTypeUnmarshaller{},

			"ElevatedCurveType": &ElevatedCurveTypeUnmarshaller{},

			"ElevatedPointPropertyType": &ElevatedPointPropertyTypeUnmarshaller{},

			"ElevatedPointType": &ElevatedPointTypeUnmarshaller{},

			"ElevatedSurfacePropertyType": &ElevatedSurfacePropertyTypeUnmarshaller{},

			"ElevatedSurfaceType": &ElevatedSurfaceTypeUnmarshaller{},

			"GroundLightSystemPropertyType": &GroundLightSystemPropertyTypeUnmarshaller{},

			"LightElementPropertyType": &LightElementPropertyTypeUnmarshaller{},

			"LightElementStatusPropertyType": &LightElementStatusPropertyTypeUnmarshaller{},

			"LightElementStatusType": &LightElementStatusTypeUnmarshaller{},

			"LightElementType": &LightElementTypeUnmarshaller{},

			"LinguisticNotePropertyType": &LinguisticNotePropertyTypeUnmarshaller{},

			"LinguisticNoteType": &LinguisticNoteTypeUnmarshaller{},

			"MarkerBeaconPropertyType": &MarkerBeaconPropertyTypeUnmarshaller{},

			"NavaidEquipmentPropertyType": &NavaidEquipmentPropertyTypeUnmarshaller{},

			"NotePropertyType": &NotePropertyTypeUnmarshaller{},

			"NoteType": &NoteTypeUnmarshaller{},

			"ObstacleAreaTimeSlicePropertyType": &ObstacleAreaTimeSlicePropertyTypeUnmarshaller{},

			"ObstacleAreaTimeSliceType": &ObstacleAreaTimeSliceTypeUnmarshaller{},

			"ObstacleAreaType": &ObstacleAreaTypeUnmarshaller{},

			"OrganisationAuthorityPropertyType": &OrganisationAuthorityPropertyTypeUnmarshaller{},

			"PassengerServicePropertyType": &PassengerServicePropertyTypeUnmarshaller{},

			"PointType": &PointTypeUnmarshaller{},

			"RunwayDirectionPropertyType": &RunwayDirectionPropertyTypeUnmarshaller{},

			"ServicePropertyType": &ServicePropertyTypeUnmarshaller{},

			"SpecialNavigationStationPropertyType": &SpecialNavigationStationPropertyTypeUnmarshaller{},

			"SurfacePropertyType": &SurfacePropertyTypeUnmarshaller{},

			"SurfaceType": &SurfaceTypeUnmarshaller{},

			"TimesheetPropertyType": &TimesheetPropertyTypeUnmarshaller{},

			"TimesheetType": &TimesheetTypeUnmarshaller{},

			"UnitPropertyType": &UnitPropertyTypeUnmarshaller{},

			"VerticalStructureLightingStatusPropertyType": &VerticalStructureLightingStatusPropertyTypeUnmarshaller{},

			"VerticalStructureLightingStatusType": &VerticalStructureLightingStatusTypeUnmarshaller{},

			"VerticalStructurePartPropertyType": &VerticalStructurePartPropertyTypeUnmarshaller{},

			"VerticalStructurePartType": &VerticalStructurePartTypeUnmarshaller{},

			"VerticalStructurePropertyType": &VerticalStructurePropertyTypeUnmarshaller{},

			"VerticalStructureTimeSlicePropertyType": &VerticalStructureTimeSlicePropertyTypeUnmarshaller{},

			"VerticalStructureTimeSliceType": &VerticalStructureTimeSliceTypeUnmarshaller{},

			"VerticalStructureType": &VerticalStructureTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["A_extension"] = len(stage.A_extensions)
	stage.Map_GongStructName_InstancesNb["A_extension_1"] = len(stage.A_extension_1s)
	stage.Map_GongStructName_InstancesNb["A_extension_10"] = len(stage.A_extension_10s)
	stage.Map_GongStructName_InstancesNb["A_extension_11"] = len(stage.A_extension_11s)
	stage.Map_GongStructName_InstancesNb["A_extension_12"] = len(stage.A_extension_12s)
	stage.Map_GongStructName_InstancesNb["A_extension_2"] = len(stage.A_extension_2s)
	stage.Map_GongStructName_InstancesNb["A_extension_3"] = len(stage.A_extension_3s)
	stage.Map_GongStructName_InstancesNb["A_extension_4"] = len(stage.A_extension_4s)
	stage.Map_GongStructName_InstancesNb["A_extension_5"] = len(stage.A_extension_5s)
	stage.Map_GongStructName_InstancesNb["A_extension_6"] = len(stage.A_extension_6s)
	stage.Map_GongStructName_InstancesNb["A_extension_7"] = len(stage.A_extension_7s)
	stage.Map_GongStructName_InstancesNb["A_extension_8"] = len(stage.A_extension_8s)
	stage.Map_GongStructName_InstancesNb["A_extension_9"] = len(stage.A_extension_9s)
	stage.Map_GongStructName_InstancesNb["AbstractPropertiesWithScheduleType"] = len(stage.AbstractPropertiesWithScheduleTypes)
	stage.Map_GongStructName_InstancesNb["AeronauticalGroundLightTimeSlicePropertyType"] = len(stage.AeronauticalGroundLightTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["AeronauticalGroundLightTimeSliceType"] = len(stage.AeronauticalGroundLightTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["AeronauticalGroundLightType"] = len(stage.AeronauticalGroundLightTypes)
	stage.Map_GongStructName_InstancesNb["AirportHeliportPropertyType"] = len(stage.AirportHeliportPropertyTypes)
	stage.Map_GongStructName_InstancesNb["CurveType"] = len(stage.CurveTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedCurvePropertyType"] = len(stage.ElevatedCurvePropertyTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedCurveType"] = len(stage.ElevatedCurveTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedPointPropertyType"] = len(stage.ElevatedPointPropertyTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedPointType"] = len(stage.ElevatedPointTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedSurfacePropertyType"] = len(stage.ElevatedSurfacePropertyTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedSurfaceType"] = len(stage.ElevatedSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["GroundLightSystemPropertyType"] = len(stage.GroundLightSystemPropertyTypes)
	stage.Map_GongStructName_InstancesNb["LightElementPropertyType"] = len(stage.LightElementPropertyTypes)
	stage.Map_GongStructName_InstancesNb["LightElementStatusPropertyType"] = len(stage.LightElementStatusPropertyTypes)
	stage.Map_GongStructName_InstancesNb["LightElementStatusType"] = len(stage.LightElementStatusTypes)
	stage.Map_GongStructName_InstancesNb["LightElementType"] = len(stage.LightElementTypes)
	stage.Map_GongStructName_InstancesNb["LinguisticNotePropertyType"] = len(stage.LinguisticNotePropertyTypes)
	stage.Map_GongStructName_InstancesNb["LinguisticNoteType"] = len(stage.LinguisticNoteTypes)
	stage.Map_GongStructName_InstancesNb["MarkerBeaconPropertyType"] = len(stage.MarkerBeaconPropertyTypes)
	stage.Map_GongStructName_InstancesNb["NavaidEquipmentPropertyType"] = len(stage.NavaidEquipmentPropertyTypes)
	stage.Map_GongStructName_InstancesNb["NotePropertyType"] = len(stage.NotePropertyTypes)
	stage.Map_GongStructName_InstancesNb["NoteType"] = len(stage.NoteTypes)
	stage.Map_GongStructName_InstancesNb["ObstacleAreaTimeSlicePropertyType"] = len(stage.ObstacleAreaTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["ObstacleAreaTimeSliceType"] = len(stage.ObstacleAreaTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["ObstacleAreaType"] = len(stage.ObstacleAreaTypes)
	stage.Map_GongStructName_InstancesNb["OrganisationAuthorityPropertyType"] = len(stage.OrganisationAuthorityPropertyTypes)
	stage.Map_GongStructName_InstancesNb["PassengerServicePropertyType"] = len(stage.PassengerServicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["PointType"] = len(stage.PointTypes)
	stage.Map_GongStructName_InstancesNb["RunwayDirectionPropertyType"] = len(stage.RunwayDirectionPropertyTypes)
	stage.Map_GongStructName_InstancesNb["ServicePropertyType"] = len(stage.ServicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["SpecialNavigationStationPropertyType"] = len(stage.SpecialNavigationStationPropertyTypes)
	stage.Map_GongStructName_InstancesNb["SurfacePropertyType"] = len(stage.SurfacePropertyTypes)
	stage.Map_GongStructName_InstancesNb["SurfaceType"] = len(stage.SurfaceTypes)
	stage.Map_GongStructName_InstancesNb["TimesheetPropertyType"] = len(stage.TimesheetPropertyTypes)
	stage.Map_GongStructName_InstancesNb["TimesheetType"] = len(stage.TimesheetTypes)
	stage.Map_GongStructName_InstancesNb["UnitPropertyType"] = len(stage.UnitPropertyTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructureLightingStatusPropertyType"] = len(stage.VerticalStructureLightingStatusPropertyTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructureLightingStatusType"] = len(stage.VerticalStructureLightingStatusTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructurePartPropertyType"] = len(stage.VerticalStructurePartPropertyTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructurePartType"] = len(stage.VerticalStructurePartTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructurePropertyType"] = len(stage.VerticalStructurePropertyTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructureTimeSlicePropertyType"] = len(stage.VerticalStructureTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructureTimeSliceType"] = len(stage.VerticalStructureTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["VerticalStructureType"] = len(stage.VerticalStructureTypes)
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
// Stage puts a_extension to the model stage
func (a_extension *A_extension) Stage(stage *Stage) *A_extension {
	__gong__stage(stage.A_extensions, stage.A_extension_stagedOrder, stage.A_extension_orderStaged, &stage.A_extensionOrder, stage.A_extensions_mapString, a_extension, a_extension.Name)
	return a_extension
}

// StagePreserveOrder puts a_extension to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extensionOrder
// - update stage.A_extensionOrder accordingly
func (a_extension *A_extension) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extensions, stage.A_extension_stagedOrder, stage.A_extension_orderStaged, &stage.A_extensionOrder, stage.A_extensions_mapString, a_extension, order, a_extension.Name)
}

// Unstage removes a_extension off the model stage
func (a_extension *A_extension) Unstage(stage *Stage) *A_extension {
	__gong__unstage(stage.A_extensions, stage.A_extensions_mapString, a_extension, a_extension.Name)
	return a_extension
}

// UnstageVoid removes a_extension off the model stage
func (a_extension *A_extension) UnstageVoid(stage *Stage) {
	a_extension.Unstage(stage)
}

func (a_extension *A_extension) StageVoid(stage *Stage) {
	a_extension.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension *A_extension) GetName() (res string) {
	return a_extension.Name
}

// for satisfaction of GongStruct interface
func (a_extension *A_extension) SetName(name string) {
	a_extension.Name = name
}

// Stage puts a_extension_1 to the model stage
func (a_extension_1 *A_extension_1) Stage(stage *Stage) *A_extension_1 {
	__gong__stage(stage.A_extension_1s, stage.A_extension_1_stagedOrder, stage.A_extension_1_orderStaged, &stage.A_extension_1Order, stage.A_extension_1s_mapString, a_extension_1, a_extension_1.Name)
	return a_extension_1
}

// StagePreserveOrder puts a_extension_1 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_1Order
// - update stage.A_extension_1Order accordingly
func (a_extension_1 *A_extension_1) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_1s, stage.A_extension_1_stagedOrder, stage.A_extension_1_orderStaged, &stage.A_extension_1Order, stage.A_extension_1s_mapString, a_extension_1, order, a_extension_1.Name)
}

// Unstage removes a_extension_1 off the model stage
func (a_extension_1 *A_extension_1) Unstage(stage *Stage) *A_extension_1 {
	__gong__unstage(stage.A_extension_1s, stage.A_extension_1s_mapString, a_extension_1, a_extension_1.Name)
	return a_extension_1
}

// UnstageVoid removes a_extension_1 off the model stage
func (a_extension_1 *A_extension_1) UnstageVoid(stage *Stage) {
	a_extension_1.Unstage(stage)
}

func (a_extension_1 *A_extension_1) StageVoid(stage *Stage) {
	a_extension_1.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_1 *A_extension_1) GetName() (res string) {
	return a_extension_1.Name
}

// for satisfaction of GongStruct interface
func (a_extension_1 *A_extension_1) SetName(name string) {
	a_extension_1.Name = name
}

// Stage puts a_extension_10 to the model stage
func (a_extension_10 *A_extension_10) Stage(stage *Stage) *A_extension_10 {
	__gong__stage(stage.A_extension_10s, stage.A_extension_10_stagedOrder, stage.A_extension_10_orderStaged, &stage.A_extension_10Order, stage.A_extension_10s_mapString, a_extension_10, a_extension_10.Name)
	return a_extension_10
}

// StagePreserveOrder puts a_extension_10 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_10Order
// - update stage.A_extension_10Order accordingly
func (a_extension_10 *A_extension_10) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_10s, stage.A_extension_10_stagedOrder, stage.A_extension_10_orderStaged, &stage.A_extension_10Order, stage.A_extension_10s_mapString, a_extension_10, order, a_extension_10.Name)
}

// Unstage removes a_extension_10 off the model stage
func (a_extension_10 *A_extension_10) Unstage(stage *Stage) *A_extension_10 {
	__gong__unstage(stage.A_extension_10s, stage.A_extension_10s_mapString, a_extension_10, a_extension_10.Name)
	return a_extension_10
}

// UnstageVoid removes a_extension_10 off the model stage
func (a_extension_10 *A_extension_10) UnstageVoid(stage *Stage) {
	a_extension_10.Unstage(stage)
}

func (a_extension_10 *A_extension_10) StageVoid(stage *Stage) {
	a_extension_10.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_10 *A_extension_10) GetName() (res string) {
	return a_extension_10.Name
}

// for satisfaction of GongStruct interface
func (a_extension_10 *A_extension_10) SetName(name string) {
	a_extension_10.Name = name
}

// Stage puts a_extension_11 to the model stage
func (a_extension_11 *A_extension_11) Stage(stage *Stage) *A_extension_11 {
	__gong__stage(stage.A_extension_11s, stage.A_extension_11_stagedOrder, stage.A_extension_11_orderStaged, &stage.A_extension_11Order, stage.A_extension_11s_mapString, a_extension_11, a_extension_11.Name)
	return a_extension_11
}

// StagePreserveOrder puts a_extension_11 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_11Order
// - update stage.A_extension_11Order accordingly
func (a_extension_11 *A_extension_11) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_11s, stage.A_extension_11_stagedOrder, stage.A_extension_11_orderStaged, &stage.A_extension_11Order, stage.A_extension_11s_mapString, a_extension_11, order, a_extension_11.Name)
}

// Unstage removes a_extension_11 off the model stage
func (a_extension_11 *A_extension_11) Unstage(stage *Stage) *A_extension_11 {
	__gong__unstage(stage.A_extension_11s, stage.A_extension_11s_mapString, a_extension_11, a_extension_11.Name)
	return a_extension_11
}

// UnstageVoid removes a_extension_11 off the model stage
func (a_extension_11 *A_extension_11) UnstageVoid(stage *Stage) {
	a_extension_11.Unstage(stage)
}

func (a_extension_11 *A_extension_11) StageVoid(stage *Stage) {
	a_extension_11.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_11 *A_extension_11) GetName() (res string) {
	return a_extension_11.Name
}

// for satisfaction of GongStruct interface
func (a_extension_11 *A_extension_11) SetName(name string) {
	a_extension_11.Name = name
}

// Stage puts a_extension_12 to the model stage
func (a_extension_12 *A_extension_12) Stage(stage *Stage) *A_extension_12 {
	__gong__stage(stage.A_extension_12s, stage.A_extension_12_stagedOrder, stage.A_extension_12_orderStaged, &stage.A_extension_12Order, stage.A_extension_12s_mapString, a_extension_12, a_extension_12.Name)
	return a_extension_12
}

// StagePreserveOrder puts a_extension_12 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_12Order
// - update stage.A_extension_12Order accordingly
func (a_extension_12 *A_extension_12) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_12s, stage.A_extension_12_stagedOrder, stage.A_extension_12_orderStaged, &stage.A_extension_12Order, stage.A_extension_12s_mapString, a_extension_12, order, a_extension_12.Name)
}

// Unstage removes a_extension_12 off the model stage
func (a_extension_12 *A_extension_12) Unstage(stage *Stage) *A_extension_12 {
	__gong__unstage(stage.A_extension_12s, stage.A_extension_12s_mapString, a_extension_12, a_extension_12.Name)
	return a_extension_12
}

// UnstageVoid removes a_extension_12 off the model stage
func (a_extension_12 *A_extension_12) UnstageVoid(stage *Stage) {
	a_extension_12.Unstage(stage)
}

func (a_extension_12 *A_extension_12) StageVoid(stage *Stage) {
	a_extension_12.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_12 *A_extension_12) GetName() (res string) {
	return a_extension_12.Name
}

// for satisfaction of GongStruct interface
func (a_extension_12 *A_extension_12) SetName(name string) {
	a_extension_12.Name = name
}

// Stage puts a_extension_2 to the model stage
func (a_extension_2 *A_extension_2) Stage(stage *Stage) *A_extension_2 {
	__gong__stage(stage.A_extension_2s, stage.A_extension_2_stagedOrder, stage.A_extension_2_orderStaged, &stage.A_extension_2Order, stage.A_extension_2s_mapString, a_extension_2, a_extension_2.Name)
	return a_extension_2
}

// StagePreserveOrder puts a_extension_2 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_2Order
// - update stage.A_extension_2Order accordingly
func (a_extension_2 *A_extension_2) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_2s, stage.A_extension_2_stagedOrder, stage.A_extension_2_orderStaged, &stage.A_extension_2Order, stage.A_extension_2s_mapString, a_extension_2, order, a_extension_2.Name)
}

// Unstage removes a_extension_2 off the model stage
func (a_extension_2 *A_extension_2) Unstage(stage *Stage) *A_extension_2 {
	__gong__unstage(stage.A_extension_2s, stage.A_extension_2s_mapString, a_extension_2, a_extension_2.Name)
	return a_extension_2
}

// UnstageVoid removes a_extension_2 off the model stage
func (a_extension_2 *A_extension_2) UnstageVoid(stage *Stage) {
	a_extension_2.Unstage(stage)
}

func (a_extension_2 *A_extension_2) StageVoid(stage *Stage) {
	a_extension_2.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_2 *A_extension_2) GetName() (res string) {
	return a_extension_2.Name
}

// for satisfaction of GongStruct interface
func (a_extension_2 *A_extension_2) SetName(name string) {
	a_extension_2.Name = name
}

// Stage puts a_extension_3 to the model stage
func (a_extension_3 *A_extension_3) Stage(stage *Stage) *A_extension_3 {
	__gong__stage(stage.A_extension_3s, stage.A_extension_3_stagedOrder, stage.A_extension_3_orderStaged, &stage.A_extension_3Order, stage.A_extension_3s_mapString, a_extension_3, a_extension_3.Name)
	return a_extension_3
}

// StagePreserveOrder puts a_extension_3 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_3Order
// - update stage.A_extension_3Order accordingly
func (a_extension_3 *A_extension_3) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_3s, stage.A_extension_3_stagedOrder, stage.A_extension_3_orderStaged, &stage.A_extension_3Order, stage.A_extension_3s_mapString, a_extension_3, order, a_extension_3.Name)
}

// Unstage removes a_extension_3 off the model stage
func (a_extension_3 *A_extension_3) Unstage(stage *Stage) *A_extension_3 {
	__gong__unstage(stage.A_extension_3s, stage.A_extension_3s_mapString, a_extension_3, a_extension_3.Name)
	return a_extension_3
}

// UnstageVoid removes a_extension_3 off the model stage
func (a_extension_3 *A_extension_3) UnstageVoid(stage *Stage) {
	a_extension_3.Unstage(stage)
}

func (a_extension_3 *A_extension_3) StageVoid(stage *Stage) {
	a_extension_3.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_3 *A_extension_3) GetName() (res string) {
	return a_extension_3.Name
}

// for satisfaction of GongStruct interface
func (a_extension_3 *A_extension_3) SetName(name string) {
	a_extension_3.Name = name
}

// Stage puts a_extension_4 to the model stage
func (a_extension_4 *A_extension_4) Stage(stage *Stage) *A_extension_4 {
	__gong__stage(stage.A_extension_4s, stage.A_extension_4_stagedOrder, stage.A_extension_4_orderStaged, &stage.A_extension_4Order, stage.A_extension_4s_mapString, a_extension_4, a_extension_4.Name)
	return a_extension_4
}

// StagePreserveOrder puts a_extension_4 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_4Order
// - update stage.A_extension_4Order accordingly
func (a_extension_4 *A_extension_4) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_4s, stage.A_extension_4_stagedOrder, stage.A_extension_4_orderStaged, &stage.A_extension_4Order, stage.A_extension_4s_mapString, a_extension_4, order, a_extension_4.Name)
}

// Unstage removes a_extension_4 off the model stage
func (a_extension_4 *A_extension_4) Unstage(stage *Stage) *A_extension_4 {
	__gong__unstage(stage.A_extension_4s, stage.A_extension_4s_mapString, a_extension_4, a_extension_4.Name)
	return a_extension_4
}

// UnstageVoid removes a_extension_4 off the model stage
func (a_extension_4 *A_extension_4) UnstageVoid(stage *Stage) {
	a_extension_4.Unstage(stage)
}

func (a_extension_4 *A_extension_4) StageVoid(stage *Stage) {
	a_extension_4.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_4 *A_extension_4) GetName() (res string) {
	return a_extension_4.Name
}

// for satisfaction of GongStruct interface
func (a_extension_4 *A_extension_4) SetName(name string) {
	a_extension_4.Name = name
}

// Stage puts a_extension_5 to the model stage
func (a_extension_5 *A_extension_5) Stage(stage *Stage) *A_extension_5 {
	__gong__stage(stage.A_extension_5s, stage.A_extension_5_stagedOrder, stage.A_extension_5_orderStaged, &stage.A_extension_5Order, stage.A_extension_5s_mapString, a_extension_5, a_extension_5.Name)
	return a_extension_5
}

// StagePreserveOrder puts a_extension_5 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_5Order
// - update stage.A_extension_5Order accordingly
func (a_extension_5 *A_extension_5) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_5s, stage.A_extension_5_stagedOrder, stage.A_extension_5_orderStaged, &stage.A_extension_5Order, stage.A_extension_5s_mapString, a_extension_5, order, a_extension_5.Name)
}

// Unstage removes a_extension_5 off the model stage
func (a_extension_5 *A_extension_5) Unstage(stage *Stage) *A_extension_5 {
	__gong__unstage(stage.A_extension_5s, stage.A_extension_5s_mapString, a_extension_5, a_extension_5.Name)
	return a_extension_5
}

// UnstageVoid removes a_extension_5 off the model stage
func (a_extension_5 *A_extension_5) UnstageVoid(stage *Stage) {
	a_extension_5.Unstage(stage)
}

func (a_extension_5 *A_extension_5) StageVoid(stage *Stage) {
	a_extension_5.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_5 *A_extension_5) GetName() (res string) {
	return a_extension_5.Name
}

// for satisfaction of GongStruct interface
func (a_extension_5 *A_extension_5) SetName(name string) {
	a_extension_5.Name = name
}

// Stage puts a_extension_6 to the model stage
func (a_extension_6 *A_extension_6) Stage(stage *Stage) *A_extension_6 {
	__gong__stage(stage.A_extension_6s, stage.A_extension_6_stagedOrder, stage.A_extension_6_orderStaged, &stage.A_extension_6Order, stage.A_extension_6s_mapString, a_extension_6, a_extension_6.Name)
	return a_extension_6
}

// StagePreserveOrder puts a_extension_6 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_6Order
// - update stage.A_extension_6Order accordingly
func (a_extension_6 *A_extension_6) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_6s, stage.A_extension_6_stagedOrder, stage.A_extension_6_orderStaged, &stage.A_extension_6Order, stage.A_extension_6s_mapString, a_extension_6, order, a_extension_6.Name)
}

// Unstage removes a_extension_6 off the model stage
func (a_extension_6 *A_extension_6) Unstage(stage *Stage) *A_extension_6 {
	__gong__unstage(stage.A_extension_6s, stage.A_extension_6s_mapString, a_extension_6, a_extension_6.Name)
	return a_extension_6
}

// UnstageVoid removes a_extension_6 off the model stage
func (a_extension_6 *A_extension_6) UnstageVoid(stage *Stage) {
	a_extension_6.Unstage(stage)
}

func (a_extension_6 *A_extension_6) StageVoid(stage *Stage) {
	a_extension_6.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_6 *A_extension_6) GetName() (res string) {
	return a_extension_6.Name
}

// for satisfaction of GongStruct interface
func (a_extension_6 *A_extension_6) SetName(name string) {
	a_extension_6.Name = name
}

// Stage puts a_extension_7 to the model stage
func (a_extension_7 *A_extension_7) Stage(stage *Stage) *A_extension_7 {
	__gong__stage(stage.A_extension_7s, stage.A_extension_7_stagedOrder, stage.A_extension_7_orderStaged, &stage.A_extension_7Order, stage.A_extension_7s_mapString, a_extension_7, a_extension_7.Name)
	return a_extension_7
}

// StagePreserveOrder puts a_extension_7 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_7Order
// - update stage.A_extension_7Order accordingly
func (a_extension_7 *A_extension_7) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_7s, stage.A_extension_7_stagedOrder, stage.A_extension_7_orderStaged, &stage.A_extension_7Order, stage.A_extension_7s_mapString, a_extension_7, order, a_extension_7.Name)
}

// Unstage removes a_extension_7 off the model stage
func (a_extension_7 *A_extension_7) Unstage(stage *Stage) *A_extension_7 {
	__gong__unstage(stage.A_extension_7s, stage.A_extension_7s_mapString, a_extension_7, a_extension_7.Name)
	return a_extension_7
}

// UnstageVoid removes a_extension_7 off the model stage
func (a_extension_7 *A_extension_7) UnstageVoid(stage *Stage) {
	a_extension_7.Unstage(stage)
}

func (a_extension_7 *A_extension_7) StageVoid(stage *Stage) {
	a_extension_7.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_7 *A_extension_7) GetName() (res string) {
	return a_extension_7.Name
}

// for satisfaction of GongStruct interface
func (a_extension_7 *A_extension_7) SetName(name string) {
	a_extension_7.Name = name
}

// Stage puts a_extension_8 to the model stage
func (a_extension_8 *A_extension_8) Stage(stage *Stage) *A_extension_8 {
	__gong__stage(stage.A_extension_8s, stage.A_extension_8_stagedOrder, stage.A_extension_8_orderStaged, &stage.A_extension_8Order, stage.A_extension_8s_mapString, a_extension_8, a_extension_8.Name)
	return a_extension_8
}

// StagePreserveOrder puts a_extension_8 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_8Order
// - update stage.A_extension_8Order accordingly
func (a_extension_8 *A_extension_8) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_8s, stage.A_extension_8_stagedOrder, stage.A_extension_8_orderStaged, &stage.A_extension_8Order, stage.A_extension_8s_mapString, a_extension_8, order, a_extension_8.Name)
}

// Unstage removes a_extension_8 off the model stage
func (a_extension_8 *A_extension_8) Unstage(stage *Stage) *A_extension_8 {
	__gong__unstage(stage.A_extension_8s, stage.A_extension_8s_mapString, a_extension_8, a_extension_8.Name)
	return a_extension_8
}

// UnstageVoid removes a_extension_8 off the model stage
func (a_extension_8 *A_extension_8) UnstageVoid(stage *Stage) {
	a_extension_8.Unstage(stage)
}

func (a_extension_8 *A_extension_8) StageVoid(stage *Stage) {
	a_extension_8.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_8 *A_extension_8) GetName() (res string) {
	return a_extension_8.Name
}

// for satisfaction of GongStruct interface
func (a_extension_8 *A_extension_8) SetName(name string) {
	a_extension_8.Name = name
}

// Stage puts a_extension_9 to the model stage
func (a_extension_9 *A_extension_9) Stage(stage *Stage) *A_extension_9 {
	__gong__stage(stage.A_extension_9s, stage.A_extension_9_stagedOrder, stage.A_extension_9_orderStaged, &stage.A_extension_9Order, stage.A_extension_9s_mapString, a_extension_9, a_extension_9.Name)
	return a_extension_9
}

// StagePreserveOrder puts a_extension_9 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_9Order
// - update stage.A_extension_9Order accordingly
func (a_extension_9 *A_extension_9) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_9s, stage.A_extension_9_stagedOrder, stage.A_extension_9_orderStaged, &stage.A_extension_9Order, stage.A_extension_9s_mapString, a_extension_9, order, a_extension_9.Name)
}

// Unstage removes a_extension_9 off the model stage
func (a_extension_9 *A_extension_9) Unstage(stage *Stage) *A_extension_9 {
	__gong__unstage(stage.A_extension_9s, stage.A_extension_9s_mapString, a_extension_9, a_extension_9.Name)
	return a_extension_9
}

// UnstageVoid removes a_extension_9 off the model stage
func (a_extension_9 *A_extension_9) UnstageVoid(stage *Stage) {
	a_extension_9.Unstage(stage)
}

func (a_extension_9 *A_extension_9) StageVoid(stage *Stage) {
	a_extension_9.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_9 *A_extension_9) GetName() (res string) {
	return a_extension_9.Name
}

// for satisfaction of GongStruct interface
func (a_extension_9 *A_extension_9) SetName(name string) {
	a_extension_9.Name = name
}

// Stage puts abstractpropertieswithscheduletype to the model stage
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) Stage(stage *Stage) *AbstractPropertiesWithScheduleType {
	__gong__stage(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleType_orderStaged, &stage.AbstractPropertiesWithScheduleTypeOrder, stage.AbstractPropertiesWithScheduleTypes_mapString, abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.Name)
	return abstractpropertieswithscheduletype
}

// StagePreserveOrder puts abstractpropertieswithscheduletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AbstractPropertiesWithScheduleTypeOrder
// - update stage.AbstractPropertiesWithScheduleTypeOrder accordingly
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleType_orderStaged, &stage.AbstractPropertiesWithScheduleTypeOrder, stage.AbstractPropertiesWithScheduleTypes_mapString, abstractpropertieswithscheduletype, order, abstractpropertieswithscheduletype.Name)
}

// Unstage removes abstractpropertieswithscheduletype off the model stage
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) Unstage(stage *Stage) *AbstractPropertiesWithScheduleType {
	__gong__unstage(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_mapString, abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.Name)
	return abstractpropertieswithscheduletype
}

// UnstageVoid removes abstractpropertieswithscheduletype off the model stage
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) UnstageVoid(stage *Stage) {
	abstractpropertieswithscheduletype.Unstage(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) StageVoid(stage *Stage) {
	abstractpropertieswithscheduletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GetName() (res string) {
	return abstractpropertieswithscheduletype.Name
}

// for satisfaction of GongStruct interface
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) SetName(name string) {
	abstractpropertieswithscheduletype.Name = name
}

// Stage puts aeronauticalgroundlighttimeslicepropertytype to the model stage
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) Stage(stage *Stage) *AeronauticalGroundLightTimeSlicePropertyType {
	__gong__stage(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder, stage.AeronauticalGroundLightTimeSlicePropertyType_orderStaged, &stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder, stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString, aeronauticalgroundlighttimeslicepropertytype, aeronauticalgroundlighttimeslicepropertytype.Name)
	return aeronauticalgroundlighttimeslicepropertytype
}

// StagePreserveOrder puts aeronauticalgroundlighttimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder
// - update stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder accordingly
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder, stage.AeronauticalGroundLightTimeSlicePropertyType_orderStaged, &stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder, stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString, aeronauticalgroundlighttimeslicepropertytype, order, aeronauticalgroundlighttimeslicepropertytype.Name)
}

// Unstage removes aeronauticalgroundlighttimeslicepropertytype off the model stage
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) Unstage(stage *Stage) *AeronauticalGroundLightTimeSlicePropertyType {
	__gong__unstage(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString, aeronauticalgroundlighttimeslicepropertytype, aeronauticalgroundlighttimeslicepropertytype.Name)
	return aeronauticalgroundlighttimeslicepropertytype
}

// UnstageVoid removes aeronauticalgroundlighttimeslicepropertytype off the model stage
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	aeronauticalgroundlighttimeslicepropertytype.Unstage(stage)
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) StageVoid(stage *Stage) {
	aeronauticalgroundlighttimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GetName() (res string) {
	return aeronauticalgroundlighttimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) SetName(name string) {
	aeronauticalgroundlighttimeslicepropertytype.Name = name
}

// Stage puts aeronauticalgroundlighttimeslicetype to the model stage
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) Stage(stage *Stage) *AeronauticalGroundLightTimeSliceType {
	__gong__stage(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceType_stagedOrder, stage.AeronauticalGroundLightTimeSliceType_orderStaged, &stage.AeronauticalGroundLightTimeSliceTypeOrder, stage.AeronauticalGroundLightTimeSliceTypes_mapString, aeronauticalgroundlighttimeslicetype, aeronauticalgroundlighttimeslicetype.Name)
	return aeronauticalgroundlighttimeslicetype
}

// StagePreserveOrder puts aeronauticalgroundlighttimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AeronauticalGroundLightTimeSliceTypeOrder
// - update stage.AeronauticalGroundLightTimeSliceTypeOrder accordingly
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceType_stagedOrder, stage.AeronauticalGroundLightTimeSliceType_orderStaged, &stage.AeronauticalGroundLightTimeSliceTypeOrder, stage.AeronauticalGroundLightTimeSliceTypes_mapString, aeronauticalgroundlighttimeslicetype, order, aeronauticalgroundlighttimeslicetype.Name)
}

// Unstage removes aeronauticalgroundlighttimeslicetype off the model stage
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) Unstage(stage *Stage) *AeronauticalGroundLightTimeSliceType {
	__gong__unstage(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceTypes_mapString, aeronauticalgroundlighttimeslicetype, aeronauticalgroundlighttimeslicetype.Name)
	return aeronauticalgroundlighttimeslicetype
}

// UnstageVoid removes aeronauticalgroundlighttimeslicetype off the model stage
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) UnstageVoid(stage *Stage) {
	aeronauticalgroundlighttimeslicetype.Unstage(stage)
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) StageVoid(stage *Stage) {
	aeronauticalgroundlighttimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GetName() (res string) {
	return aeronauticalgroundlighttimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) SetName(name string) {
	aeronauticalgroundlighttimeslicetype.Name = name
}

// Stage puts aeronauticalgroundlighttype to the model stage
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) Stage(stage *Stage) *AeronauticalGroundLightType {
	__gong__stage(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightType_stagedOrder, stage.AeronauticalGroundLightType_orderStaged, &stage.AeronauticalGroundLightTypeOrder, stage.AeronauticalGroundLightTypes_mapString, aeronauticalgroundlighttype, aeronauticalgroundlighttype.Name)
	return aeronauticalgroundlighttype
}

// StagePreserveOrder puts aeronauticalgroundlighttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AeronauticalGroundLightTypeOrder
// - update stage.AeronauticalGroundLightTypeOrder accordingly
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightType_stagedOrder, stage.AeronauticalGroundLightType_orderStaged, &stage.AeronauticalGroundLightTypeOrder, stage.AeronauticalGroundLightTypes_mapString, aeronauticalgroundlighttype, order, aeronauticalgroundlighttype.Name)
}

// Unstage removes aeronauticalgroundlighttype off the model stage
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) Unstage(stage *Stage) *AeronauticalGroundLightType {
	__gong__unstage(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightTypes_mapString, aeronauticalgroundlighttype, aeronauticalgroundlighttype.Name)
	return aeronauticalgroundlighttype
}

// UnstageVoid removes aeronauticalgroundlighttype off the model stage
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) UnstageVoid(stage *Stage) {
	aeronauticalgroundlighttype.Unstage(stage)
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) StageVoid(stage *Stage) {
	aeronauticalgroundlighttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GetName() (res string) {
	return aeronauticalgroundlighttype.Name
}

// for satisfaction of GongStruct interface
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) SetName(name string) {
	aeronauticalgroundlighttype.Name = name
}

// Stage puts airportheliportpropertytype to the model stage
func (airportheliportpropertytype *AirportHeliportPropertyType) Stage(stage *Stage) *AirportHeliportPropertyType {
	__gong__stage(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyType_orderStaged, &stage.AirportHeliportPropertyTypeOrder, stage.AirportHeliportPropertyTypes_mapString, airportheliportpropertytype, airportheliportpropertytype.Name)
	return airportheliportpropertytype
}

// StagePreserveOrder puts airportheliportpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirportHeliportPropertyTypeOrder
// - update stage.AirportHeliportPropertyTypeOrder accordingly
func (airportheliportpropertytype *AirportHeliportPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyType_orderStaged, &stage.AirportHeliportPropertyTypeOrder, stage.AirportHeliportPropertyTypes_mapString, airportheliportpropertytype, order, airportheliportpropertytype.Name)
}

// Unstage removes airportheliportpropertytype off the model stage
func (airportheliportpropertytype *AirportHeliportPropertyType) Unstage(stage *Stage) *AirportHeliportPropertyType {
	__gong__unstage(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_mapString, airportheliportpropertytype, airportheliportpropertytype.Name)
	return airportheliportpropertytype
}

// UnstageVoid removes airportheliportpropertytype off the model stage
func (airportheliportpropertytype *AirportHeliportPropertyType) UnstageVoid(stage *Stage) {
	airportheliportpropertytype.Unstage(stage)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) StageVoid(stage *Stage) {
	airportheliportpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airportheliportpropertytype *AirportHeliportPropertyType) GetName() (res string) {
	return airportheliportpropertytype.Name
}

// for satisfaction of GongStruct interface
func (airportheliportpropertytype *AirportHeliportPropertyType) SetName(name string) {
	airportheliportpropertytype.Name = name
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

// Stage puts elevatedcurvepropertytype to the model stage
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) Stage(stage *Stage) *ElevatedCurvePropertyType {
	__gong__stage(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyType_stagedOrder, stage.ElevatedCurvePropertyType_orderStaged, &stage.ElevatedCurvePropertyTypeOrder, stage.ElevatedCurvePropertyTypes_mapString, elevatedcurvepropertytype, elevatedcurvepropertytype.Name)
	return elevatedcurvepropertytype
}

// StagePreserveOrder puts elevatedcurvepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedCurvePropertyTypeOrder
// - update stage.ElevatedCurvePropertyTypeOrder accordingly
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyType_stagedOrder, stage.ElevatedCurvePropertyType_orderStaged, &stage.ElevatedCurvePropertyTypeOrder, stage.ElevatedCurvePropertyTypes_mapString, elevatedcurvepropertytype, order, elevatedcurvepropertytype.Name)
}

// Unstage removes elevatedcurvepropertytype off the model stage
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) Unstage(stage *Stage) *ElevatedCurvePropertyType {
	__gong__unstage(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyTypes_mapString, elevatedcurvepropertytype, elevatedcurvepropertytype.Name)
	return elevatedcurvepropertytype
}

// UnstageVoid removes elevatedcurvepropertytype off the model stage
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) UnstageVoid(stage *Stage) {
	elevatedcurvepropertytype.Unstage(stage)
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) StageVoid(stage *Stage) {
	elevatedcurvepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GetName() (res string) {
	return elevatedcurvepropertytype.Name
}

// for satisfaction of GongStruct interface
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) SetName(name string) {
	elevatedcurvepropertytype.Name = name
}

// Stage puts elevatedcurvetype to the model stage
func (elevatedcurvetype *ElevatedCurveType) Stage(stage *Stage) *ElevatedCurveType {
	__gong__stage(stage.ElevatedCurveTypes, stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveType_orderStaged, &stage.ElevatedCurveTypeOrder, stage.ElevatedCurveTypes_mapString, elevatedcurvetype, elevatedcurvetype.Name)
	return elevatedcurvetype
}

// StagePreserveOrder puts elevatedcurvetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedCurveTypeOrder
// - update stage.ElevatedCurveTypeOrder accordingly
func (elevatedcurvetype *ElevatedCurveType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedCurveTypes, stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveType_orderStaged, &stage.ElevatedCurveTypeOrder, stage.ElevatedCurveTypes_mapString, elevatedcurvetype, order, elevatedcurvetype.Name)
}

// Unstage removes elevatedcurvetype off the model stage
func (elevatedcurvetype *ElevatedCurveType) Unstage(stage *Stage) *ElevatedCurveType {
	__gong__unstage(stage.ElevatedCurveTypes, stage.ElevatedCurveTypes_mapString, elevatedcurvetype, elevatedcurvetype.Name)
	return elevatedcurvetype
}

// UnstageVoid removes elevatedcurvetype off the model stage
func (elevatedcurvetype *ElevatedCurveType) UnstageVoid(stage *Stage) {
	elevatedcurvetype.Unstage(stage)
}

func (elevatedcurvetype *ElevatedCurveType) StageVoid(stage *Stage) {
	elevatedcurvetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedcurvetype *ElevatedCurveType) GetName() (res string) {
	return elevatedcurvetype.Name
}

// for satisfaction of GongStruct interface
func (elevatedcurvetype *ElevatedCurveType) SetName(name string) {
	elevatedcurvetype.Name = name
}

// Stage puts elevatedpointpropertytype to the model stage
func (elevatedpointpropertytype *ElevatedPointPropertyType) Stage(stage *Stage) *ElevatedPointPropertyType {
	__gong__stage(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyType_stagedOrder, stage.ElevatedPointPropertyType_orderStaged, &stage.ElevatedPointPropertyTypeOrder, stage.ElevatedPointPropertyTypes_mapString, elevatedpointpropertytype, elevatedpointpropertytype.Name)
	return elevatedpointpropertytype
}

// StagePreserveOrder puts elevatedpointpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedPointPropertyTypeOrder
// - update stage.ElevatedPointPropertyTypeOrder accordingly
func (elevatedpointpropertytype *ElevatedPointPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyType_stagedOrder, stage.ElevatedPointPropertyType_orderStaged, &stage.ElevatedPointPropertyTypeOrder, stage.ElevatedPointPropertyTypes_mapString, elevatedpointpropertytype, order, elevatedpointpropertytype.Name)
}

// Unstage removes elevatedpointpropertytype off the model stage
func (elevatedpointpropertytype *ElevatedPointPropertyType) Unstage(stage *Stage) *ElevatedPointPropertyType {
	__gong__unstage(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyTypes_mapString, elevatedpointpropertytype, elevatedpointpropertytype.Name)
	return elevatedpointpropertytype
}

// UnstageVoid removes elevatedpointpropertytype off the model stage
func (elevatedpointpropertytype *ElevatedPointPropertyType) UnstageVoid(stage *Stage) {
	elevatedpointpropertytype.Unstage(stage)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) StageVoid(stage *Stage) {
	elevatedpointpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedpointpropertytype *ElevatedPointPropertyType) GetName() (res string) {
	return elevatedpointpropertytype.Name
}

// for satisfaction of GongStruct interface
func (elevatedpointpropertytype *ElevatedPointPropertyType) SetName(name string) {
	elevatedpointpropertytype.Name = name
}

// Stage puts elevatedpointtype to the model stage
func (elevatedpointtype *ElevatedPointType) Stage(stage *Stage) *ElevatedPointType {
	__gong__stage(stage.ElevatedPointTypes, stage.ElevatedPointType_stagedOrder, stage.ElevatedPointType_orderStaged, &stage.ElevatedPointTypeOrder, stage.ElevatedPointTypes_mapString, elevatedpointtype, elevatedpointtype.Name)
	return elevatedpointtype
}

// StagePreserveOrder puts elevatedpointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedPointTypeOrder
// - update stage.ElevatedPointTypeOrder accordingly
func (elevatedpointtype *ElevatedPointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedPointTypes, stage.ElevatedPointType_stagedOrder, stage.ElevatedPointType_orderStaged, &stage.ElevatedPointTypeOrder, stage.ElevatedPointTypes_mapString, elevatedpointtype, order, elevatedpointtype.Name)
}

// Unstage removes elevatedpointtype off the model stage
func (elevatedpointtype *ElevatedPointType) Unstage(stage *Stage) *ElevatedPointType {
	__gong__unstage(stage.ElevatedPointTypes, stage.ElevatedPointTypes_mapString, elevatedpointtype, elevatedpointtype.Name)
	return elevatedpointtype
}

// UnstageVoid removes elevatedpointtype off the model stage
func (elevatedpointtype *ElevatedPointType) UnstageVoid(stage *Stage) {
	elevatedpointtype.Unstage(stage)
}

func (elevatedpointtype *ElevatedPointType) StageVoid(stage *Stage) {
	elevatedpointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedpointtype *ElevatedPointType) GetName() (res string) {
	return elevatedpointtype.Name
}

// for satisfaction of GongStruct interface
func (elevatedpointtype *ElevatedPointType) SetName(name string) {
	elevatedpointtype.Name = name
}

// Stage puts elevatedsurfacepropertytype to the model stage
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) Stage(stage *Stage) *ElevatedSurfacePropertyType {
	__gong__stage(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyType_stagedOrder, stage.ElevatedSurfacePropertyType_orderStaged, &stage.ElevatedSurfacePropertyTypeOrder, stage.ElevatedSurfacePropertyTypes_mapString, elevatedsurfacepropertytype, elevatedsurfacepropertytype.Name)
	return elevatedsurfacepropertytype
}

// StagePreserveOrder puts elevatedsurfacepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedSurfacePropertyTypeOrder
// - update stage.ElevatedSurfacePropertyTypeOrder accordingly
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyType_stagedOrder, stage.ElevatedSurfacePropertyType_orderStaged, &stage.ElevatedSurfacePropertyTypeOrder, stage.ElevatedSurfacePropertyTypes_mapString, elevatedsurfacepropertytype, order, elevatedsurfacepropertytype.Name)
}

// Unstage removes elevatedsurfacepropertytype off the model stage
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) Unstage(stage *Stage) *ElevatedSurfacePropertyType {
	__gong__unstage(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyTypes_mapString, elevatedsurfacepropertytype, elevatedsurfacepropertytype.Name)
	return elevatedsurfacepropertytype
}

// UnstageVoid removes elevatedsurfacepropertytype off the model stage
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) UnstageVoid(stage *Stage) {
	elevatedsurfacepropertytype.Unstage(stage)
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) StageVoid(stage *Stage) {
	elevatedsurfacepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GetName() (res string) {
	return elevatedsurfacepropertytype.Name
}

// for satisfaction of GongStruct interface
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) SetName(name string) {
	elevatedsurfacepropertytype.Name = name
}

// Stage puts elevatedsurfacetype to the model stage
func (elevatedsurfacetype *ElevatedSurfaceType) Stage(stage *Stage) *ElevatedSurfaceType {
	__gong__stage(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceType_orderStaged, &stage.ElevatedSurfaceTypeOrder, stage.ElevatedSurfaceTypes_mapString, elevatedsurfacetype, elevatedsurfacetype.Name)
	return elevatedsurfacetype
}

// StagePreserveOrder puts elevatedsurfacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElevatedSurfaceTypeOrder
// - update stage.ElevatedSurfaceTypeOrder accordingly
func (elevatedsurfacetype *ElevatedSurfaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceType_orderStaged, &stage.ElevatedSurfaceTypeOrder, stage.ElevatedSurfaceTypes_mapString, elevatedsurfacetype, order, elevatedsurfacetype.Name)
}

// Unstage removes elevatedsurfacetype off the model stage
func (elevatedsurfacetype *ElevatedSurfaceType) Unstage(stage *Stage) *ElevatedSurfaceType {
	__gong__unstage(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceTypes_mapString, elevatedsurfacetype, elevatedsurfacetype.Name)
	return elevatedsurfacetype
}

// UnstageVoid removes elevatedsurfacetype off the model stage
func (elevatedsurfacetype *ElevatedSurfaceType) UnstageVoid(stage *Stage) {
	elevatedsurfacetype.Unstage(stage)
}

func (elevatedsurfacetype *ElevatedSurfaceType) StageVoid(stage *Stage) {
	elevatedsurfacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elevatedsurfacetype *ElevatedSurfaceType) GetName() (res string) {
	return elevatedsurfacetype.Name
}

// for satisfaction of GongStruct interface
func (elevatedsurfacetype *ElevatedSurfaceType) SetName(name string) {
	elevatedsurfacetype.Name = name
}

// Stage puts groundlightsystempropertytype to the model stage
func (groundlightsystempropertytype *GroundLightSystemPropertyType) Stage(stage *Stage) *GroundLightSystemPropertyType {
	__gong__stage(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyType_stagedOrder, stage.GroundLightSystemPropertyType_orderStaged, &stage.GroundLightSystemPropertyTypeOrder, stage.GroundLightSystemPropertyTypes_mapString, groundlightsystempropertytype, groundlightsystempropertytype.Name)
	return groundlightsystempropertytype
}

// StagePreserveOrder puts groundlightsystempropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroundLightSystemPropertyTypeOrder
// - update stage.GroundLightSystemPropertyTypeOrder accordingly
func (groundlightsystempropertytype *GroundLightSystemPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyType_stagedOrder, stage.GroundLightSystemPropertyType_orderStaged, &stage.GroundLightSystemPropertyTypeOrder, stage.GroundLightSystemPropertyTypes_mapString, groundlightsystempropertytype, order, groundlightsystempropertytype.Name)
}

// Unstage removes groundlightsystempropertytype off the model stage
func (groundlightsystempropertytype *GroundLightSystemPropertyType) Unstage(stage *Stage) *GroundLightSystemPropertyType {
	__gong__unstage(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyTypes_mapString, groundlightsystempropertytype, groundlightsystempropertytype.Name)
	return groundlightsystempropertytype
}

// UnstageVoid removes groundlightsystempropertytype off the model stage
func (groundlightsystempropertytype *GroundLightSystemPropertyType) UnstageVoid(stage *Stage) {
	groundlightsystempropertytype.Unstage(stage)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) StageVoid(stage *Stage) {
	groundlightsystempropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GetName() (res string) {
	return groundlightsystempropertytype.Name
}

// for satisfaction of GongStruct interface
func (groundlightsystempropertytype *GroundLightSystemPropertyType) SetName(name string) {
	groundlightsystempropertytype.Name = name
}

// Stage puts lightelementpropertytype to the model stage
func (lightelementpropertytype *LightElementPropertyType) Stage(stage *Stage) *LightElementPropertyType {
	__gong__stage(stage.LightElementPropertyTypes, stage.LightElementPropertyType_stagedOrder, stage.LightElementPropertyType_orderStaged, &stage.LightElementPropertyTypeOrder, stage.LightElementPropertyTypes_mapString, lightelementpropertytype, lightelementpropertytype.Name)
	return lightelementpropertytype
}

// StagePreserveOrder puts lightelementpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LightElementPropertyTypeOrder
// - update stage.LightElementPropertyTypeOrder accordingly
func (lightelementpropertytype *LightElementPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LightElementPropertyTypes, stage.LightElementPropertyType_stagedOrder, stage.LightElementPropertyType_orderStaged, &stage.LightElementPropertyTypeOrder, stage.LightElementPropertyTypes_mapString, lightelementpropertytype, order, lightelementpropertytype.Name)
}

// Unstage removes lightelementpropertytype off the model stage
func (lightelementpropertytype *LightElementPropertyType) Unstage(stage *Stage) *LightElementPropertyType {
	__gong__unstage(stage.LightElementPropertyTypes, stage.LightElementPropertyTypes_mapString, lightelementpropertytype, lightelementpropertytype.Name)
	return lightelementpropertytype
}

// UnstageVoid removes lightelementpropertytype off the model stage
func (lightelementpropertytype *LightElementPropertyType) UnstageVoid(stage *Stage) {
	lightelementpropertytype.Unstage(stage)
}

func (lightelementpropertytype *LightElementPropertyType) StageVoid(stage *Stage) {
	lightelementpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lightelementpropertytype *LightElementPropertyType) GetName() (res string) {
	return lightelementpropertytype.Name
}

// for satisfaction of GongStruct interface
func (lightelementpropertytype *LightElementPropertyType) SetName(name string) {
	lightelementpropertytype.Name = name
}

// Stage puts lightelementstatuspropertytype to the model stage
func (lightelementstatuspropertytype *LightElementStatusPropertyType) Stage(stage *Stage) *LightElementStatusPropertyType {
	__gong__stage(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyType_stagedOrder, stage.LightElementStatusPropertyType_orderStaged, &stage.LightElementStatusPropertyTypeOrder, stage.LightElementStatusPropertyTypes_mapString, lightelementstatuspropertytype, lightelementstatuspropertytype.Name)
	return lightelementstatuspropertytype
}

// StagePreserveOrder puts lightelementstatuspropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LightElementStatusPropertyTypeOrder
// - update stage.LightElementStatusPropertyTypeOrder accordingly
func (lightelementstatuspropertytype *LightElementStatusPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyType_stagedOrder, stage.LightElementStatusPropertyType_orderStaged, &stage.LightElementStatusPropertyTypeOrder, stage.LightElementStatusPropertyTypes_mapString, lightelementstatuspropertytype, order, lightelementstatuspropertytype.Name)
}

// Unstage removes lightelementstatuspropertytype off the model stage
func (lightelementstatuspropertytype *LightElementStatusPropertyType) Unstage(stage *Stage) *LightElementStatusPropertyType {
	__gong__unstage(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyTypes_mapString, lightelementstatuspropertytype, lightelementstatuspropertytype.Name)
	return lightelementstatuspropertytype
}

// UnstageVoid removes lightelementstatuspropertytype off the model stage
func (lightelementstatuspropertytype *LightElementStatusPropertyType) UnstageVoid(stage *Stage) {
	lightelementstatuspropertytype.Unstage(stage)
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) StageVoid(stage *Stage) {
	lightelementstatuspropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lightelementstatuspropertytype *LightElementStatusPropertyType) GetName() (res string) {
	return lightelementstatuspropertytype.Name
}

// for satisfaction of GongStruct interface
func (lightelementstatuspropertytype *LightElementStatusPropertyType) SetName(name string) {
	lightelementstatuspropertytype.Name = name
}

// Stage puts lightelementstatustype to the model stage
func (lightelementstatustype *LightElementStatusType) Stage(stage *Stage) *LightElementStatusType {
	__gong__stage(stage.LightElementStatusTypes, stage.LightElementStatusType_stagedOrder, stage.LightElementStatusType_orderStaged, &stage.LightElementStatusTypeOrder, stage.LightElementStatusTypes_mapString, lightelementstatustype, lightelementstatustype.Name)
	return lightelementstatustype
}

// StagePreserveOrder puts lightelementstatustype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LightElementStatusTypeOrder
// - update stage.LightElementStatusTypeOrder accordingly
func (lightelementstatustype *LightElementStatusType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LightElementStatusTypes, stage.LightElementStatusType_stagedOrder, stage.LightElementStatusType_orderStaged, &stage.LightElementStatusTypeOrder, stage.LightElementStatusTypes_mapString, lightelementstatustype, order, lightelementstatustype.Name)
}

// Unstage removes lightelementstatustype off the model stage
func (lightelementstatustype *LightElementStatusType) Unstage(stage *Stage) *LightElementStatusType {
	__gong__unstage(stage.LightElementStatusTypes, stage.LightElementStatusTypes_mapString, lightelementstatustype, lightelementstatustype.Name)
	return lightelementstatustype
}

// UnstageVoid removes lightelementstatustype off the model stage
func (lightelementstatustype *LightElementStatusType) UnstageVoid(stage *Stage) {
	lightelementstatustype.Unstage(stage)
}

func (lightelementstatustype *LightElementStatusType) StageVoid(stage *Stage) {
	lightelementstatustype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lightelementstatustype *LightElementStatusType) GetName() (res string) {
	return lightelementstatustype.Name
}

// for satisfaction of GongStruct interface
func (lightelementstatustype *LightElementStatusType) SetName(name string) {
	lightelementstatustype.Name = name
}

// Stage puts lightelementtype to the model stage
func (lightelementtype *LightElementType) Stage(stage *Stage) *LightElementType {
	__gong__stage(stage.LightElementTypes, stage.LightElementType_stagedOrder, stage.LightElementType_orderStaged, &stage.LightElementTypeOrder, stage.LightElementTypes_mapString, lightelementtype, lightelementtype.Name)
	return lightelementtype
}

// StagePreserveOrder puts lightelementtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LightElementTypeOrder
// - update stage.LightElementTypeOrder accordingly
func (lightelementtype *LightElementType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LightElementTypes, stage.LightElementType_stagedOrder, stage.LightElementType_orderStaged, &stage.LightElementTypeOrder, stage.LightElementTypes_mapString, lightelementtype, order, lightelementtype.Name)
}

// Unstage removes lightelementtype off the model stage
func (lightelementtype *LightElementType) Unstage(stage *Stage) *LightElementType {
	__gong__unstage(stage.LightElementTypes, stage.LightElementTypes_mapString, lightelementtype, lightelementtype.Name)
	return lightelementtype
}

// UnstageVoid removes lightelementtype off the model stage
func (lightelementtype *LightElementType) UnstageVoid(stage *Stage) {
	lightelementtype.Unstage(stage)
}

func (lightelementtype *LightElementType) StageVoid(stage *Stage) {
	lightelementtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lightelementtype *LightElementType) GetName() (res string) {
	return lightelementtype.Name
}

// for satisfaction of GongStruct interface
func (lightelementtype *LightElementType) SetName(name string) {
	lightelementtype.Name = name
}

// Stage puts linguisticnotepropertytype to the model stage
func (linguisticnotepropertytype *LinguisticNotePropertyType) Stage(stage *Stage) *LinguisticNotePropertyType {
	__gong__stage(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyType_orderStaged, &stage.LinguisticNotePropertyTypeOrder, stage.LinguisticNotePropertyTypes_mapString, linguisticnotepropertytype, linguisticnotepropertytype.Name)
	return linguisticnotepropertytype
}

// StagePreserveOrder puts linguisticnotepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinguisticNotePropertyTypeOrder
// - update stage.LinguisticNotePropertyTypeOrder accordingly
func (linguisticnotepropertytype *LinguisticNotePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyType_orderStaged, &stage.LinguisticNotePropertyTypeOrder, stage.LinguisticNotePropertyTypes_mapString, linguisticnotepropertytype, order, linguisticnotepropertytype.Name)
}

// Unstage removes linguisticnotepropertytype off the model stage
func (linguisticnotepropertytype *LinguisticNotePropertyType) Unstage(stage *Stage) *LinguisticNotePropertyType {
	__gong__unstage(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_mapString, linguisticnotepropertytype, linguisticnotepropertytype.Name)
	return linguisticnotepropertytype
}

// UnstageVoid removes linguisticnotepropertytype off the model stage
func (linguisticnotepropertytype *LinguisticNotePropertyType) UnstageVoid(stage *Stage) {
	linguisticnotepropertytype.Unstage(stage)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) StageVoid(stage *Stage) {
	linguisticnotepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linguisticnotepropertytype *LinguisticNotePropertyType) GetName() (res string) {
	return linguisticnotepropertytype.Name
}

// for satisfaction of GongStruct interface
func (linguisticnotepropertytype *LinguisticNotePropertyType) SetName(name string) {
	linguisticnotepropertytype.Name = name
}

// Stage puts linguisticnotetype to the model stage
func (linguisticnotetype *LinguisticNoteType) Stage(stage *Stage) *LinguisticNoteType {
	__gong__stage(stage.LinguisticNoteTypes, stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteType_orderStaged, &stage.LinguisticNoteTypeOrder, stage.LinguisticNoteTypes_mapString, linguisticnotetype, linguisticnotetype.Name)
	return linguisticnotetype
}

// StagePreserveOrder puts linguisticnotetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinguisticNoteTypeOrder
// - update stage.LinguisticNoteTypeOrder accordingly
func (linguisticnotetype *LinguisticNoteType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LinguisticNoteTypes, stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteType_orderStaged, &stage.LinguisticNoteTypeOrder, stage.LinguisticNoteTypes_mapString, linguisticnotetype, order, linguisticnotetype.Name)
}

// Unstage removes linguisticnotetype off the model stage
func (linguisticnotetype *LinguisticNoteType) Unstage(stage *Stage) *LinguisticNoteType {
	__gong__unstage(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_mapString, linguisticnotetype, linguisticnotetype.Name)
	return linguisticnotetype
}

// UnstageVoid removes linguisticnotetype off the model stage
func (linguisticnotetype *LinguisticNoteType) UnstageVoid(stage *Stage) {
	linguisticnotetype.Unstage(stage)
}

func (linguisticnotetype *LinguisticNoteType) StageVoid(stage *Stage) {
	linguisticnotetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linguisticnotetype *LinguisticNoteType) GetName() (res string) {
	return linguisticnotetype.Name
}

// for satisfaction of GongStruct interface
func (linguisticnotetype *LinguisticNoteType) SetName(name string) {
	linguisticnotetype.Name = name
}

// Stage puts markerbeaconpropertytype to the model stage
func (markerbeaconpropertytype *MarkerBeaconPropertyType) Stage(stage *Stage) *MarkerBeaconPropertyType {
	__gong__stage(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyType_stagedOrder, stage.MarkerBeaconPropertyType_orderStaged, &stage.MarkerBeaconPropertyTypeOrder, stage.MarkerBeaconPropertyTypes_mapString, markerbeaconpropertytype, markerbeaconpropertytype.Name)
	return markerbeaconpropertytype
}

// StagePreserveOrder puts markerbeaconpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MarkerBeaconPropertyTypeOrder
// - update stage.MarkerBeaconPropertyTypeOrder accordingly
func (markerbeaconpropertytype *MarkerBeaconPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyType_stagedOrder, stage.MarkerBeaconPropertyType_orderStaged, &stage.MarkerBeaconPropertyTypeOrder, stage.MarkerBeaconPropertyTypes_mapString, markerbeaconpropertytype, order, markerbeaconpropertytype.Name)
}

// Unstage removes markerbeaconpropertytype off the model stage
func (markerbeaconpropertytype *MarkerBeaconPropertyType) Unstage(stage *Stage) *MarkerBeaconPropertyType {
	__gong__unstage(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyTypes_mapString, markerbeaconpropertytype, markerbeaconpropertytype.Name)
	return markerbeaconpropertytype
}

// UnstageVoid removes markerbeaconpropertytype off the model stage
func (markerbeaconpropertytype *MarkerBeaconPropertyType) UnstageVoid(stage *Stage) {
	markerbeaconpropertytype.Unstage(stage)
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) StageVoid(stage *Stage) {
	markerbeaconpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (markerbeaconpropertytype *MarkerBeaconPropertyType) GetName() (res string) {
	return markerbeaconpropertytype.Name
}

// for satisfaction of GongStruct interface
func (markerbeaconpropertytype *MarkerBeaconPropertyType) SetName(name string) {
	markerbeaconpropertytype.Name = name
}

// Stage puts navaidequipmentpropertytype to the model stage
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) Stage(stage *Stage) *NavaidEquipmentPropertyType {
	__gong__stage(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyType_stagedOrder, stage.NavaidEquipmentPropertyType_orderStaged, &stage.NavaidEquipmentPropertyTypeOrder, stage.NavaidEquipmentPropertyTypes_mapString, navaidequipmentpropertytype, navaidequipmentpropertytype.Name)
	return navaidequipmentpropertytype
}

// StagePreserveOrder puts navaidequipmentpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NavaidEquipmentPropertyTypeOrder
// - update stage.NavaidEquipmentPropertyTypeOrder accordingly
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyType_stagedOrder, stage.NavaidEquipmentPropertyType_orderStaged, &stage.NavaidEquipmentPropertyTypeOrder, stage.NavaidEquipmentPropertyTypes_mapString, navaidequipmentpropertytype, order, navaidequipmentpropertytype.Name)
}

// Unstage removes navaidequipmentpropertytype off the model stage
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) Unstage(stage *Stage) *NavaidEquipmentPropertyType {
	__gong__unstage(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyTypes_mapString, navaidequipmentpropertytype, navaidequipmentpropertytype.Name)
	return navaidequipmentpropertytype
}

// UnstageVoid removes navaidequipmentpropertytype off the model stage
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) UnstageVoid(stage *Stage) {
	navaidequipmentpropertytype.Unstage(stage)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) StageVoid(stage *Stage) {
	navaidequipmentpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GetName() (res string) {
	return navaidequipmentpropertytype.Name
}

// for satisfaction of GongStruct interface
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) SetName(name string) {
	navaidequipmentpropertytype.Name = name
}

// Stage puts notepropertytype to the model stage
func (notepropertytype *NotePropertyType) Stage(stage *Stage) *NotePropertyType {
	__gong__stage(stage.NotePropertyTypes, stage.NotePropertyType_stagedOrder, stage.NotePropertyType_orderStaged, &stage.NotePropertyTypeOrder, stage.NotePropertyTypes_mapString, notepropertytype, notepropertytype.Name)
	return notepropertytype
}

// StagePreserveOrder puts notepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NotePropertyTypeOrder
// - update stage.NotePropertyTypeOrder accordingly
func (notepropertytype *NotePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NotePropertyTypes, stage.NotePropertyType_stagedOrder, stage.NotePropertyType_orderStaged, &stage.NotePropertyTypeOrder, stage.NotePropertyTypes_mapString, notepropertytype, order, notepropertytype.Name)
}

// Unstage removes notepropertytype off the model stage
func (notepropertytype *NotePropertyType) Unstage(stage *Stage) *NotePropertyType {
	__gong__unstage(stage.NotePropertyTypes, stage.NotePropertyTypes_mapString, notepropertytype, notepropertytype.Name)
	return notepropertytype
}

// UnstageVoid removes notepropertytype off the model stage
func (notepropertytype *NotePropertyType) UnstageVoid(stage *Stage) {
	notepropertytype.Unstage(stage)
}

func (notepropertytype *NotePropertyType) StageVoid(stage *Stage) {
	notepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notepropertytype *NotePropertyType) GetName() (res string) {
	return notepropertytype.Name
}

// for satisfaction of GongStruct interface
func (notepropertytype *NotePropertyType) SetName(name string) {
	notepropertytype.Name = name
}

// Stage puts notetype to the model stage
func (notetype *NoteType) Stage(stage *Stage) *NoteType {
	__gong__stage(stage.NoteTypes, stage.NoteType_stagedOrder, stage.NoteType_orderStaged, &stage.NoteTypeOrder, stage.NoteTypes_mapString, notetype, notetype.Name)
	return notetype
}

// StagePreserveOrder puts notetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteTypeOrder
// - update stage.NoteTypeOrder accordingly
func (notetype *NoteType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteTypes, stage.NoteType_stagedOrder, stage.NoteType_orderStaged, &stage.NoteTypeOrder, stage.NoteTypes_mapString, notetype, order, notetype.Name)
}

// Unstage removes notetype off the model stage
func (notetype *NoteType) Unstage(stage *Stage) *NoteType {
	__gong__unstage(stage.NoteTypes, stage.NoteTypes_mapString, notetype, notetype.Name)
	return notetype
}

// UnstageVoid removes notetype off the model stage
func (notetype *NoteType) UnstageVoid(stage *Stage) {
	notetype.Unstage(stage)
}

func (notetype *NoteType) StageVoid(stage *Stage) {
	notetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notetype *NoteType) GetName() (res string) {
	return notetype.Name
}

// for satisfaction of GongStruct interface
func (notetype *NoteType) SetName(name string) {
	notetype.Name = name
}

// Stage puts obstacleareatimeslicepropertytype to the model stage
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) Stage(stage *Stage) *ObstacleAreaTimeSlicePropertyType {
	__gong__stage(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyType_stagedOrder, stage.ObstacleAreaTimeSlicePropertyType_orderStaged, &stage.ObstacleAreaTimeSlicePropertyTypeOrder, stage.ObstacleAreaTimeSlicePropertyTypes_mapString, obstacleareatimeslicepropertytype, obstacleareatimeslicepropertytype.Name)
	return obstacleareatimeslicepropertytype
}

// StagePreserveOrder puts obstacleareatimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObstacleAreaTimeSlicePropertyTypeOrder
// - update stage.ObstacleAreaTimeSlicePropertyTypeOrder accordingly
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyType_stagedOrder, stage.ObstacleAreaTimeSlicePropertyType_orderStaged, &stage.ObstacleAreaTimeSlicePropertyTypeOrder, stage.ObstacleAreaTimeSlicePropertyTypes_mapString, obstacleareatimeslicepropertytype, order, obstacleareatimeslicepropertytype.Name)
}

// Unstage removes obstacleareatimeslicepropertytype off the model stage
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) Unstage(stage *Stage) *ObstacleAreaTimeSlicePropertyType {
	__gong__unstage(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyTypes_mapString, obstacleareatimeslicepropertytype, obstacleareatimeslicepropertytype.Name)
	return obstacleareatimeslicepropertytype
}

// UnstageVoid removes obstacleareatimeslicepropertytype off the model stage
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	obstacleareatimeslicepropertytype.Unstage(stage)
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) StageVoid(stage *Stage) {
	obstacleareatimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GetName() (res string) {
	return obstacleareatimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) SetName(name string) {
	obstacleareatimeslicepropertytype.Name = name
}

// Stage puts obstacleareatimeslicetype to the model stage
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) Stage(stage *Stage) *ObstacleAreaTimeSliceType {
	__gong__stage(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceType_stagedOrder, stage.ObstacleAreaTimeSliceType_orderStaged, &stage.ObstacleAreaTimeSliceTypeOrder, stage.ObstacleAreaTimeSliceTypes_mapString, obstacleareatimeslicetype, obstacleareatimeslicetype.Name)
	return obstacleareatimeslicetype
}

// StagePreserveOrder puts obstacleareatimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObstacleAreaTimeSliceTypeOrder
// - update stage.ObstacleAreaTimeSliceTypeOrder accordingly
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceType_stagedOrder, stage.ObstacleAreaTimeSliceType_orderStaged, &stage.ObstacleAreaTimeSliceTypeOrder, stage.ObstacleAreaTimeSliceTypes_mapString, obstacleareatimeslicetype, order, obstacleareatimeslicetype.Name)
}

// Unstage removes obstacleareatimeslicetype off the model stage
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) Unstage(stage *Stage) *ObstacleAreaTimeSliceType {
	__gong__unstage(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceTypes_mapString, obstacleareatimeslicetype, obstacleareatimeslicetype.Name)
	return obstacleareatimeslicetype
}

// UnstageVoid removes obstacleareatimeslicetype off the model stage
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) UnstageVoid(stage *Stage) {
	obstacleareatimeslicetype.Unstage(stage)
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) StageVoid(stage *Stage) {
	obstacleareatimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GetName() (res string) {
	return obstacleareatimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) SetName(name string) {
	obstacleareatimeslicetype.Name = name
}

// Stage puts obstacleareatype to the model stage
func (obstacleareatype *ObstacleAreaType) Stage(stage *Stage) *ObstacleAreaType {
	__gong__stage(stage.ObstacleAreaTypes, stage.ObstacleAreaType_stagedOrder, stage.ObstacleAreaType_orderStaged, &stage.ObstacleAreaTypeOrder, stage.ObstacleAreaTypes_mapString, obstacleareatype, obstacleareatype.Name)
	return obstacleareatype
}

// StagePreserveOrder puts obstacleareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObstacleAreaTypeOrder
// - update stage.ObstacleAreaTypeOrder accordingly
func (obstacleareatype *ObstacleAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ObstacleAreaTypes, stage.ObstacleAreaType_stagedOrder, stage.ObstacleAreaType_orderStaged, &stage.ObstacleAreaTypeOrder, stage.ObstacleAreaTypes_mapString, obstacleareatype, order, obstacleareatype.Name)
}

// Unstage removes obstacleareatype off the model stage
func (obstacleareatype *ObstacleAreaType) Unstage(stage *Stage) *ObstacleAreaType {
	__gong__unstage(stage.ObstacleAreaTypes, stage.ObstacleAreaTypes_mapString, obstacleareatype, obstacleareatype.Name)
	return obstacleareatype
}

// UnstageVoid removes obstacleareatype off the model stage
func (obstacleareatype *ObstacleAreaType) UnstageVoid(stage *Stage) {
	obstacleareatype.Unstage(stage)
}

func (obstacleareatype *ObstacleAreaType) StageVoid(stage *Stage) {
	obstacleareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (obstacleareatype *ObstacleAreaType) GetName() (res string) {
	return obstacleareatype.Name
}

// for satisfaction of GongStruct interface
func (obstacleareatype *ObstacleAreaType) SetName(name string) {
	obstacleareatype.Name = name
}

// Stage puts organisationauthoritypropertytype to the model stage
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) Stage(stage *Stage) *OrganisationAuthorityPropertyType {
	__gong__stage(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyType_orderStaged, &stage.OrganisationAuthorityPropertyTypeOrder, stage.OrganisationAuthorityPropertyTypes_mapString, organisationauthoritypropertytype, organisationauthoritypropertytype.Name)
	return organisationauthoritypropertytype
}

// StagePreserveOrder puts organisationauthoritypropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OrganisationAuthorityPropertyTypeOrder
// - update stage.OrganisationAuthorityPropertyTypeOrder accordingly
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyType_orderStaged, &stage.OrganisationAuthorityPropertyTypeOrder, stage.OrganisationAuthorityPropertyTypes_mapString, organisationauthoritypropertytype, order, organisationauthoritypropertytype.Name)
}

// Unstage removes organisationauthoritypropertytype off the model stage
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) Unstage(stage *Stage) *OrganisationAuthorityPropertyType {
	__gong__unstage(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_mapString, organisationauthoritypropertytype, organisationauthoritypropertytype.Name)
	return organisationauthoritypropertytype
}

// UnstageVoid removes organisationauthoritypropertytype off the model stage
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) UnstageVoid(stage *Stage) {
	organisationauthoritypropertytype.Unstage(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) StageVoid(stage *Stage) {
	organisationauthoritypropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GetName() (res string) {
	return organisationauthoritypropertytype.Name
}

// for satisfaction of GongStruct interface
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) SetName(name string) {
	organisationauthoritypropertytype.Name = name
}

// Stage puts passengerservicepropertytype to the model stage
func (passengerservicepropertytype *PassengerServicePropertyType) Stage(stage *Stage) *PassengerServicePropertyType {
	__gong__stage(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyType_stagedOrder, stage.PassengerServicePropertyType_orderStaged, &stage.PassengerServicePropertyTypeOrder, stage.PassengerServicePropertyTypes_mapString, passengerservicepropertytype, passengerservicepropertytype.Name)
	return passengerservicepropertytype
}

// StagePreserveOrder puts passengerservicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PassengerServicePropertyTypeOrder
// - update stage.PassengerServicePropertyTypeOrder accordingly
func (passengerservicepropertytype *PassengerServicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyType_stagedOrder, stage.PassengerServicePropertyType_orderStaged, &stage.PassengerServicePropertyTypeOrder, stage.PassengerServicePropertyTypes_mapString, passengerservicepropertytype, order, passengerservicepropertytype.Name)
}

// Unstage removes passengerservicepropertytype off the model stage
func (passengerservicepropertytype *PassengerServicePropertyType) Unstage(stage *Stage) *PassengerServicePropertyType {
	__gong__unstage(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyTypes_mapString, passengerservicepropertytype, passengerservicepropertytype.Name)
	return passengerservicepropertytype
}

// UnstageVoid removes passengerservicepropertytype off the model stage
func (passengerservicepropertytype *PassengerServicePropertyType) UnstageVoid(stage *Stage) {
	passengerservicepropertytype.Unstage(stage)
}

func (passengerservicepropertytype *PassengerServicePropertyType) StageVoid(stage *Stage) {
	passengerservicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (passengerservicepropertytype *PassengerServicePropertyType) GetName() (res string) {
	return passengerservicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (passengerservicepropertytype *PassengerServicePropertyType) SetName(name string) {
	passengerservicepropertytype.Name = name
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

// Stage puts runwaydirectionpropertytype to the model stage
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) Stage(stage *Stage) *RunwayDirectionPropertyType {
	__gong__stage(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyType_stagedOrder, stage.RunwayDirectionPropertyType_orderStaged, &stage.RunwayDirectionPropertyTypeOrder, stage.RunwayDirectionPropertyTypes_mapString, runwaydirectionpropertytype, runwaydirectionpropertytype.Name)
	return runwaydirectionpropertytype
}

// StagePreserveOrder puts runwaydirectionpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RunwayDirectionPropertyTypeOrder
// - update stage.RunwayDirectionPropertyTypeOrder accordingly
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyType_stagedOrder, stage.RunwayDirectionPropertyType_orderStaged, &stage.RunwayDirectionPropertyTypeOrder, stage.RunwayDirectionPropertyTypes_mapString, runwaydirectionpropertytype, order, runwaydirectionpropertytype.Name)
}

// Unstage removes runwaydirectionpropertytype off the model stage
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) Unstage(stage *Stage) *RunwayDirectionPropertyType {
	__gong__unstage(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyTypes_mapString, runwaydirectionpropertytype, runwaydirectionpropertytype.Name)
	return runwaydirectionpropertytype
}

// UnstageVoid removes runwaydirectionpropertytype off the model stage
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) UnstageVoid(stage *Stage) {
	runwaydirectionpropertytype.Unstage(stage)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) StageVoid(stage *Stage) {
	runwaydirectionpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GetName() (res string) {
	return runwaydirectionpropertytype.Name
}

// for satisfaction of GongStruct interface
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) SetName(name string) {
	runwaydirectionpropertytype.Name = name
}

// Stage puts servicepropertytype to the model stage
func (servicepropertytype *ServicePropertyType) Stage(stage *Stage) *ServicePropertyType {
	__gong__stage(stage.ServicePropertyTypes, stage.ServicePropertyType_stagedOrder, stage.ServicePropertyType_orderStaged, &stage.ServicePropertyTypeOrder, stage.ServicePropertyTypes_mapString, servicepropertytype, servicepropertytype.Name)
	return servicepropertytype
}

// StagePreserveOrder puts servicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ServicePropertyTypeOrder
// - update stage.ServicePropertyTypeOrder accordingly
func (servicepropertytype *ServicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ServicePropertyTypes, stage.ServicePropertyType_stagedOrder, stage.ServicePropertyType_orderStaged, &stage.ServicePropertyTypeOrder, stage.ServicePropertyTypes_mapString, servicepropertytype, order, servicepropertytype.Name)
}

// Unstage removes servicepropertytype off the model stage
func (servicepropertytype *ServicePropertyType) Unstage(stage *Stage) *ServicePropertyType {
	__gong__unstage(stage.ServicePropertyTypes, stage.ServicePropertyTypes_mapString, servicepropertytype, servicepropertytype.Name)
	return servicepropertytype
}

// UnstageVoid removes servicepropertytype off the model stage
func (servicepropertytype *ServicePropertyType) UnstageVoid(stage *Stage) {
	servicepropertytype.Unstage(stage)
}

func (servicepropertytype *ServicePropertyType) StageVoid(stage *Stage) {
	servicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (servicepropertytype *ServicePropertyType) GetName() (res string) {
	return servicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (servicepropertytype *ServicePropertyType) SetName(name string) {
	servicepropertytype.Name = name
}

// Stage puts specialnavigationstationpropertytype to the model stage
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) Stage(stage *Stage) *SpecialNavigationStationPropertyType {
	__gong__stage(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyType_stagedOrder, stage.SpecialNavigationStationPropertyType_orderStaged, &stage.SpecialNavigationStationPropertyTypeOrder, stage.SpecialNavigationStationPropertyTypes_mapString, specialnavigationstationpropertytype, specialnavigationstationpropertytype.Name)
	return specialnavigationstationpropertytype
}

// StagePreserveOrder puts specialnavigationstationpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SpecialNavigationStationPropertyTypeOrder
// - update stage.SpecialNavigationStationPropertyTypeOrder accordingly
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyType_stagedOrder, stage.SpecialNavigationStationPropertyType_orderStaged, &stage.SpecialNavigationStationPropertyTypeOrder, stage.SpecialNavigationStationPropertyTypes_mapString, specialnavigationstationpropertytype, order, specialnavigationstationpropertytype.Name)
}

// Unstage removes specialnavigationstationpropertytype off the model stage
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) Unstage(stage *Stage) *SpecialNavigationStationPropertyType {
	__gong__unstage(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyTypes_mapString, specialnavigationstationpropertytype, specialnavigationstationpropertytype.Name)
	return specialnavigationstationpropertytype
}

// UnstageVoid removes specialnavigationstationpropertytype off the model stage
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) UnstageVoid(stage *Stage) {
	specialnavigationstationpropertytype.Unstage(stage)
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) StageVoid(stage *Stage) {
	specialnavigationstationpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GetName() (res string) {
	return specialnavigationstationpropertytype.Name
}

// for satisfaction of GongStruct interface
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) SetName(name string) {
	specialnavigationstationpropertytype.Name = name
}

// Stage puts surfacepropertytype to the model stage
func (surfacepropertytype *SurfacePropertyType) Stage(stage *Stage) *SurfacePropertyType {
	__gong__stage(stage.SurfacePropertyTypes, stage.SurfacePropertyType_stagedOrder, stage.SurfacePropertyType_orderStaged, &stage.SurfacePropertyTypeOrder, stage.SurfacePropertyTypes_mapString, surfacepropertytype, surfacepropertytype.Name)
	return surfacepropertytype
}

// StagePreserveOrder puts surfacepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SurfacePropertyTypeOrder
// - update stage.SurfacePropertyTypeOrder accordingly
func (surfacepropertytype *SurfacePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SurfacePropertyTypes, stage.SurfacePropertyType_stagedOrder, stage.SurfacePropertyType_orderStaged, &stage.SurfacePropertyTypeOrder, stage.SurfacePropertyTypes_mapString, surfacepropertytype, order, surfacepropertytype.Name)
}

// Unstage removes surfacepropertytype off the model stage
func (surfacepropertytype *SurfacePropertyType) Unstage(stage *Stage) *SurfacePropertyType {
	__gong__unstage(stage.SurfacePropertyTypes, stage.SurfacePropertyTypes_mapString, surfacepropertytype, surfacepropertytype.Name)
	return surfacepropertytype
}

// UnstageVoid removes surfacepropertytype off the model stage
func (surfacepropertytype *SurfacePropertyType) UnstageVoid(stage *Stage) {
	surfacepropertytype.Unstage(stage)
}

func (surfacepropertytype *SurfacePropertyType) StageVoid(stage *Stage) {
	surfacepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (surfacepropertytype *SurfacePropertyType) GetName() (res string) {
	return surfacepropertytype.Name
}

// for satisfaction of GongStruct interface
func (surfacepropertytype *SurfacePropertyType) SetName(name string) {
	surfacepropertytype.Name = name
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

// Stage puts timesheetpropertytype to the model stage
func (timesheetpropertytype *TimesheetPropertyType) Stage(stage *Stage) *TimesheetPropertyType {
	__gong__stage(stage.TimesheetPropertyTypes, stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyType_orderStaged, &stage.TimesheetPropertyTypeOrder, stage.TimesheetPropertyTypes_mapString, timesheetpropertytype, timesheetpropertytype.Name)
	return timesheetpropertytype
}

// StagePreserveOrder puts timesheetpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimesheetPropertyTypeOrder
// - update stage.TimesheetPropertyTypeOrder accordingly
func (timesheetpropertytype *TimesheetPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimesheetPropertyTypes, stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyType_orderStaged, &stage.TimesheetPropertyTypeOrder, stage.TimesheetPropertyTypes_mapString, timesheetpropertytype, order, timesheetpropertytype.Name)
}

// Unstage removes timesheetpropertytype off the model stage
func (timesheetpropertytype *TimesheetPropertyType) Unstage(stage *Stage) *TimesheetPropertyType {
	__gong__unstage(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_mapString, timesheetpropertytype, timesheetpropertytype.Name)
	return timesheetpropertytype
}

// UnstageVoid removes timesheetpropertytype off the model stage
func (timesheetpropertytype *TimesheetPropertyType) UnstageVoid(stage *Stage) {
	timesheetpropertytype.Unstage(stage)
}

func (timesheetpropertytype *TimesheetPropertyType) StageVoid(stage *Stage) {
	timesheetpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timesheetpropertytype *TimesheetPropertyType) GetName() (res string) {
	return timesheetpropertytype.Name
}

// for satisfaction of GongStruct interface
func (timesheetpropertytype *TimesheetPropertyType) SetName(name string) {
	timesheetpropertytype.Name = name
}

// Stage puts timesheettype to the model stage
func (timesheettype *TimesheetType) Stage(stage *Stage) *TimesheetType {
	__gong__stage(stage.TimesheetTypes, stage.TimesheetType_stagedOrder, stage.TimesheetType_orderStaged, &stage.TimesheetTypeOrder, stage.TimesheetTypes_mapString, timesheettype, timesheettype.Name)
	return timesheettype
}

// StagePreserveOrder puts timesheettype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimesheetTypeOrder
// - update stage.TimesheetTypeOrder accordingly
func (timesheettype *TimesheetType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimesheetTypes, stage.TimesheetType_stagedOrder, stage.TimesheetType_orderStaged, &stage.TimesheetTypeOrder, stage.TimesheetTypes_mapString, timesheettype, order, timesheettype.Name)
}

// Unstage removes timesheettype off the model stage
func (timesheettype *TimesheetType) Unstage(stage *Stage) *TimesheetType {
	__gong__unstage(stage.TimesheetTypes, stage.TimesheetTypes_mapString, timesheettype, timesheettype.Name)
	return timesheettype
}

// UnstageVoid removes timesheettype off the model stage
func (timesheettype *TimesheetType) UnstageVoid(stage *Stage) {
	timesheettype.Unstage(stage)
}

func (timesheettype *TimesheetType) StageVoid(stage *Stage) {
	timesheettype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timesheettype *TimesheetType) GetName() (res string) {
	return timesheettype.Name
}

// for satisfaction of GongStruct interface
func (timesheettype *TimesheetType) SetName(name string) {
	timesheettype.Name = name
}

// Stage puts unitpropertytype to the model stage
func (unitpropertytype *UnitPropertyType) Stage(stage *Stage) *UnitPropertyType {
	__gong__stage(stage.UnitPropertyTypes, stage.UnitPropertyType_stagedOrder, stage.UnitPropertyType_orderStaged, &stage.UnitPropertyTypeOrder, stage.UnitPropertyTypes_mapString, unitpropertytype, unitpropertytype.Name)
	return unitpropertytype
}

// StagePreserveOrder puts unitpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UnitPropertyTypeOrder
// - update stage.UnitPropertyTypeOrder accordingly
func (unitpropertytype *UnitPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.UnitPropertyTypes, stage.UnitPropertyType_stagedOrder, stage.UnitPropertyType_orderStaged, &stage.UnitPropertyTypeOrder, stage.UnitPropertyTypes_mapString, unitpropertytype, order, unitpropertytype.Name)
}

// Unstage removes unitpropertytype off the model stage
func (unitpropertytype *UnitPropertyType) Unstage(stage *Stage) *UnitPropertyType {
	__gong__unstage(stage.UnitPropertyTypes, stage.UnitPropertyTypes_mapString, unitpropertytype, unitpropertytype.Name)
	return unitpropertytype
}

// UnstageVoid removes unitpropertytype off the model stage
func (unitpropertytype *UnitPropertyType) UnstageVoid(stage *Stage) {
	unitpropertytype.Unstage(stage)
}

func (unitpropertytype *UnitPropertyType) StageVoid(stage *Stage) {
	unitpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (unitpropertytype *UnitPropertyType) GetName() (res string) {
	return unitpropertytype.Name
}

// for satisfaction of GongStruct interface
func (unitpropertytype *UnitPropertyType) SetName(name string) {
	unitpropertytype.Name = name
}

// Stage puts verticalstructurelightingstatuspropertytype to the model stage
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) Stage(stage *Stage) *VerticalStructureLightingStatusPropertyType {
	__gong__stage(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyType_stagedOrder, stage.VerticalStructureLightingStatusPropertyType_orderStaged, &stage.VerticalStructureLightingStatusPropertyTypeOrder, stage.VerticalStructureLightingStatusPropertyTypes_mapString, verticalstructurelightingstatuspropertytype, verticalstructurelightingstatuspropertytype.Name)
	return verticalstructurelightingstatuspropertytype
}

// StagePreserveOrder puts verticalstructurelightingstatuspropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructureLightingStatusPropertyTypeOrder
// - update stage.VerticalStructureLightingStatusPropertyTypeOrder accordingly
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyType_stagedOrder, stage.VerticalStructureLightingStatusPropertyType_orderStaged, &stage.VerticalStructureLightingStatusPropertyTypeOrder, stage.VerticalStructureLightingStatusPropertyTypes_mapString, verticalstructurelightingstatuspropertytype, order, verticalstructurelightingstatuspropertytype.Name)
}

// Unstage removes verticalstructurelightingstatuspropertytype off the model stage
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) Unstage(stage *Stage) *VerticalStructureLightingStatusPropertyType {
	__gong__unstage(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyTypes_mapString, verticalstructurelightingstatuspropertytype, verticalstructurelightingstatuspropertytype.Name)
	return verticalstructurelightingstatuspropertytype
}

// UnstageVoid removes verticalstructurelightingstatuspropertytype off the model stage
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) UnstageVoid(stage *Stage) {
	verticalstructurelightingstatuspropertytype.Unstage(stage)
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) StageVoid(stage *Stage) {
	verticalstructurelightingstatuspropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GetName() (res string) {
	return verticalstructurelightingstatuspropertytype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) SetName(name string) {
	verticalstructurelightingstatuspropertytype.Name = name
}

// Stage puts verticalstructurelightingstatustype to the model stage
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) Stage(stage *Stage) *VerticalStructureLightingStatusType {
	__gong__stage(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusType_stagedOrder, stage.VerticalStructureLightingStatusType_orderStaged, &stage.VerticalStructureLightingStatusTypeOrder, stage.VerticalStructureLightingStatusTypes_mapString, verticalstructurelightingstatustype, verticalstructurelightingstatustype.Name)
	return verticalstructurelightingstatustype
}

// StagePreserveOrder puts verticalstructurelightingstatustype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructureLightingStatusTypeOrder
// - update stage.VerticalStructureLightingStatusTypeOrder accordingly
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusType_stagedOrder, stage.VerticalStructureLightingStatusType_orderStaged, &stage.VerticalStructureLightingStatusTypeOrder, stage.VerticalStructureLightingStatusTypes_mapString, verticalstructurelightingstatustype, order, verticalstructurelightingstatustype.Name)
}

// Unstage removes verticalstructurelightingstatustype off the model stage
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) Unstage(stage *Stage) *VerticalStructureLightingStatusType {
	__gong__unstage(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusTypes_mapString, verticalstructurelightingstatustype, verticalstructurelightingstatustype.Name)
	return verticalstructurelightingstatustype
}

// UnstageVoid removes verticalstructurelightingstatustype off the model stage
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) UnstageVoid(stage *Stage) {
	verticalstructurelightingstatustype.Unstage(stage)
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) StageVoid(stage *Stage) {
	verticalstructurelightingstatustype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GetName() (res string) {
	return verticalstructurelightingstatustype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) SetName(name string) {
	verticalstructurelightingstatustype.Name = name
}

// Stage puts verticalstructurepartpropertytype to the model stage
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) Stage(stage *Stage) *VerticalStructurePartPropertyType {
	__gong__stage(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyType_stagedOrder, stage.VerticalStructurePartPropertyType_orderStaged, &stage.VerticalStructurePartPropertyTypeOrder, stage.VerticalStructurePartPropertyTypes_mapString, verticalstructurepartpropertytype, verticalstructurepartpropertytype.Name)
	return verticalstructurepartpropertytype
}

// StagePreserveOrder puts verticalstructurepartpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructurePartPropertyTypeOrder
// - update stage.VerticalStructurePartPropertyTypeOrder accordingly
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyType_stagedOrder, stage.VerticalStructurePartPropertyType_orderStaged, &stage.VerticalStructurePartPropertyTypeOrder, stage.VerticalStructurePartPropertyTypes_mapString, verticalstructurepartpropertytype, order, verticalstructurepartpropertytype.Name)
}

// Unstage removes verticalstructurepartpropertytype off the model stage
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) Unstage(stage *Stage) *VerticalStructurePartPropertyType {
	__gong__unstage(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyTypes_mapString, verticalstructurepartpropertytype, verticalstructurepartpropertytype.Name)
	return verticalstructurepartpropertytype
}

// UnstageVoid removes verticalstructurepartpropertytype off the model stage
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) UnstageVoid(stage *Stage) {
	verticalstructurepartpropertytype.Unstage(stage)
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) StageVoid(stage *Stage) {
	verticalstructurepartpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GetName() (res string) {
	return verticalstructurepartpropertytype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) SetName(name string) {
	verticalstructurepartpropertytype.Name = name
}

// Stage puts verticalstructureparttype to the model stage
func (verticalstructureparttype *VerticalStructurePartType) Stage(stage *Stage) *VerticalStructurePartType {
	__gong__stage(stage.VerticalStructurePartTypes, stage.VerticalStructurePartType_stagedOrder, stage.VerticalStructurePartType_orderStaged, &stage.VerticalStructurePartTypeOrder, stage.VerticalStructurePartTypes_mapString, verticalstructureparttype, verticalstructureparttype.Name)
	return verticalstructureparttype
}

// StagePreserveOrder puts verticalstructureparttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructurePartTypeOrder
// - update stage.VerticalStructurePartTypeOrder accordingly
func (verticalstructureparttype *VerticalStructurePartType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructurePartTypes, stage.VerticalStructurePartType_stagedOrder, stage.VerticalStructurePartType_orderStaged, &stage.VerticalStructurePartTypeOrder, stage.VerticalStructurePartTypes_mapString, verticalstructureparttype, order, verticalstructureparttype.Name)
}

// Unstage removes verticalstructureparttype off the model stage
func (verticalstructureparttype *VerticalStructurePartType) Unstage(stage *Stage) *VerticalStructurePartType {
	__gong__unstage(stage.VerticalStructurePartTypes, stage.VerticalStructurePartTypes_mapString, verticalstructureparttype, verticalstructureparttype.Name)
	return verticalstructureparttype
}

// UnstageVoid removes verticalstructureparttype off the model stage
func (verticalstructureparttype *VerticalStructurePartType) UnstageVoid(stage *Stage) {
	verticalstructureparttype.Unstage(stage)
}

func (verticalstructureparttype *VerticalStructurePartType) StageVoid(stage *Stage) {
	verticalstructureparttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructureparttype *VerticalStructurePartType) GetName() (res string) {
	return verticalstructureparttype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructureparttype *VerticalStructurePartType) SetName(name string) {
	verticalstructureparttype.Name = name
}

// Stage puts verticalstructurepropertytype to the model stage
func (verticalstructurepropertytype *VerticalStructurePropertyType) Stage(stage *Stage) *VerticalStructurePropertyType {
	__gong__stage(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyType_stagedOrder, stage.VerticalStructurePropertyType_orderStaged, &stage.VerticalStructurePropertyTypeOrder, stage.VerticalStructurePropertyTypes_mapString, verticalstructurepropertytype, verticalstructurepropertytype.Name)
	return verticalstructurepropertytype
}

// StagePreserveOrder puts verticalstructurepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructurePropertyTypeOrder
// - update stage.VerticalStructurePropertyTypeOrder accordingly
func (verticalstructurepropertytype *VerticalStructurePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyType_stagedOrder, stage.VerticalStructurePropertyType_orderStaged, &stage.VerticalStructurePropertyTypeOrder, stage.VerticalStructurePropertyTypes_mapString, verticalstructurepropertytype, order, verticalstructurepropertytype.Name)
}

// Unstage removes verticalstructurepropertytype off the model stage
func (verticalstructurepropertytype *VerticalStructurePropertyType) Unstage(stage *Stage) *VerticalStructurePropertyType {
	__gong__unstage(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyTypes_mapString, verticalstructurepropertytype, verticalstructurepropertytype.Name)
	return verticalstructurepropertytype
}

// UnstageVoid removes verticalstructurepropertytype off the model stage
func (verticalstructurepropertytype *VerticalStructurePropertyType) UnstageVoid(stage *Stage) {
	verticalstructurepropertytype.Unstage(stage)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) StageVoid(stage *Stage) {
	verticalstructurepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructurepropertytype *VerticalStructurePropertyType) GetName() (res string) {
	return verticalstructurepropertytype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructurepropertytype *VerticalStructurePropertyType) SetName(name string) {
	verticalstructurepropertytype.Name = name
}

// Stage puts verticalstructuretimeslicepropertytype to the model stage
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) Stage(stage *Stage) *VerticalStructureTimeSlicePropertyType {
	__gong__stage(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyType_stagedOrder, stage.VerticalStructureTimeSlicePropertyType_orderStaged, &stage.VerticalStructureTimeSlicePropertyTypeOrder, stage.VerticalStructureTimeSlicePropertyTypes_mapString, verticalstructuretimeslicepropertytype, verticalstructuretimeslicepropertytype.Name)
	return verticalstructuretimeslicepropertytype
}

// StagePreserveOrder puts verticalstructuretimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructureTimeSlicePropertyTypeOrder
// - update stage.VerticalStructureTimeSlicePropertyTypeOrder accordingly
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyType_stagedOrder, stage.VerticalStructureTimeSlicePropertyType_orderStaged, &stage.VerticalStructureTimeSlicePropertyTypeOrder, stage.VerticalStructureTimeSlicePropertyTypes_mapString, verticalstructuretimeslicepropertytype, order, verticalstructuretimeslicepropertytype.Name)
}

// Unstage removes verticalstructuretimeslicepropertytype off the model stage
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) Unstage(stage *Stage) *VerticalStructureTimeSlicePropertyType {
	__gong__unstage(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyTypes_mapString, verticalstructuretimeslicepropertytype, verticalstructuretimeslicepropertytype.Name)
	return verticalstructuretimeslicepropertytype
}

// UnstageVoid removes verticalstructuretimeslicepropertytype off the model stage
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	verticalstructuretimeslicepropertytype.Unstage(stage)
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) StageVoid(stage *Stage) {
	verticalstructuretimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GetName() (res string) {
	return verticalstructuretimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) SetName(name string) {
	verticalstructuretimeslicepropertytype.Name = name
}

// Stage puts verticalstructuretimeslicetype to the model stage
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) Stage(stage *Stage) *VerticalStructureTimeSliceType {
	__gong__stage(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceType_stagedOrder, stage.VerticalStructureTimeSliceType_orderStaged, &stage.VerticalStructureTimeSliceTypeOrder, stage.VerticalStructureTimeSliceTypes_mapString, verticalstructuretimeslicetype, verticalstructuretimeslicetype.Name)
	return verticalstructuretimeslicetype
}

// StagePreserveOrder puts verticalstructuretimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructureTimeSliceTypeOrder
// - update stage.VerticalStructureTimeSliceTypeOrder accordingly
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceType_stagedOrder, stage.VerticalStructureTimeSliceType_orderStaged, &stage.VerticalStructureTimeSliceTypeOrder, stage.VerticalStructureTimeSliceTypes_mapString, verticalstructuretimeslicetype, order, verticalstructuretimeslicetype.Name)
}

// Unstage removes verticalstructuretimeslicetype off the model stage
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) Unstage(stage *Stage) *VerticalStructureTimeSliceType {
	__gong__unstage(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceTypes_mapString, verticalstructuretimeslicetype, verticalstructuretimeslicetype.Name)
	return verticalstructuretimeslicetype
}

// UnstageVoid removes verticalstructuretimeslicetype off the model stage
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) UnstageVoid(stage *Stage) {
	verticalstructuretimeslicetype.Unstage(stage)
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) StageVoid(stage *Stage) {
	verticalstructuretimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GetName() (res string) {
	return verticalstructuretimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) SetName(name string) {
	verticalstructuretimeslicetype.Name = name
}

// Stage puts verticalstructuretype to the model stage
func (verticalstructuretype *VerticalStructureType) Stage(stage *Stage) *VerticalStructureType {
	__gong__stage(stage.VerticalStructureTypes, stage.VerticalStructureType_stagedOrder, stage.VerticalStructureType_orderStaged, &stage.VerticalStructureTypeOrder, stage.VerticalStructureTypes_mapString, verticalstructuretype, verticalstructuretype.Name)
	return verticalstructuretype
}

// StagePreserveOrder puts verticalstructuretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VerticalStructureTypeOrder
// - update stage.VerticalStructureTypeOrder accordingly
func (verticalstructuretype *VerticalStructureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VerticalStructureTypes, stage.VerticalStructureType_stagedOrder, stage.VerticalStructureType_orderStaged, &stage.VerticalStructureTypeOrder, stage.VerticalStructureTypes_mapString, verticalstructuretype, order, verticalstructuretype.Name)
}

// Unstage removes verticalstructuretype off the model stage
func (verticalstructuretype *VerticalStructureType) Unstage(stage *Stage) *VerticalStructureType {
	__gong__unstage(stage.VerticalStructureTypes, stage.VerticalStructureTypes_mapString, verticalstructuretype, verticalstructuretype.Name)
	return verticalstructuretype
}

// UnstageVoid removes verticalstructuretype off the model stage
func (verticalstructuretype *VerticalStructureType) UnstageVoid(stage *Stage) {
	verticalstructuretype.Unstage(stage)
}

func (verticalstructuretype *VerticalStructureType) StageVoid(stage *Stage) {
	verticalstructuretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (verticalstructuretype *VerticalStructureType) GetName() (res string) {
	return verticalstructuretype.Name
}

// for satisfaction of GongStruct interface
func (verticalstructuretype *VerticalStructureType) SetName(name string) {
	verticalstructuretype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.A_extensions, &stage.A_extensions_mapString, &stage.A_extension_stagedOrder, &stage.A_extensionOrder)

	__gong__resetStageType(&stage.A_extension_1s, &stage.A_extension_1s_mapString, &stage.A_extension_1_stagedOrder, &stage.A_extension_1Order)

	__gong__resetStageType(&stage.A_extension_10s, &stage.A_extension_10s_mapString, &stage.A_extension_10_stagedOrder, &stage.A_extension_10Order)

	__gong__resetStageType(&stage.A_extension_11s, &stage.A_extension_11s_mapString, &stage.A_extension_11_stagedOrder, &stage.A_extension_11Order)

	__gong__resetStageType(&stage.A_extension_12s, &stage.A_extension_12s_mapString, &stage.A_extension_12_stagedOrder, &stage.A_extension_12Order)

	__gong__resetStageType(&stage.A_extension_2s, &stage.A_extension_2s_mapString, &stage.A_extension_2_stagedOrder, &stage.A_extension_2Order)

	__gong__resetStageType(&stage.A_extension_3s, &stage.A_extension_3s_mapString, &stage.A_extension_3_stagedOrder, &stage.A_extension_3Order)

	__gong__resetStageType(&stage.A_extension_4s, &stage.A_extension_4s_mapString, &stage.A_extension_4_stagedOrder, &stage.A_extension_4Order)

	__gong__resetStageType(&stage.A_extension_5s, &stage.A_extension_5s_mapString, &stage.A_extension_5_stagedOrder, &stage.A_extension_5Order)

	__gong__resetStageType(&stage.A_extension_6s, &stage.A_extension_6s_mapString, &stage.A_extension_6_stagedOrder, &stage.A_extension_6Order)

	__gong__resetStageType(&stage.A_extension_7s, &stage.A_extension_7s_mapString, &stage.A_extension_7_stagedOrder, &stage.A_extension_7Order)

	__gong__resetStageType(&stage.A_extension_8s, &stage.A_extension_8s_mapString, &stage.A_extension_8_stagedOrder, &stage.A_extension_8Order)

	__gong__resetStageType(&stage.A_extension_9s, &stage.A_extension_9s_mapString, &stage.A_extension_9_stagedOrder, &stage.A_extension_9Order)

	__gong__resetStageType(&stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_mapString, &stage.AbstractPropertiesWithScheduleType_stagedOrder, &stage.AbstractPropertiesWithScheduleTypeOrder)

	__gong__resetStageType(&stage.AeronauticalGroundLightTimeSlicePropertyTypes, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString, &stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder, &stage.AeronauticalGroundLightTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.AeronauticalGroundLightTimeSliceTypes, &stage.AeronauticalGroundLightTimeSliceTypes_mapString, &stage.AeronauticalGroundLightTimeSliceType_stagedOrder, &stage.AeronauticalGroundLightTimeSliceTypeOrder)

	__gong__resetStageType(&stage.AeronauticalGroundLightTypes, &stage.AeronauticalGroundLightTypes_mapString, &stage.AeronauticalGroundLightType_stagedOrder, &stage.AeronauticalGroundLightTypeOrder)

	__gong__resetStageType(&stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_mapString, &stage.AirportHeliportPropertyType_stagedOrder, &stage.AirportHeliportPropertyTypeOrder)

	__gong__resetStageType(&stage.CurveTypes, &stage.CurveTypes_mapString, &stage.CurveType_stagedOrder, &stage.CurveTypeOrder)

	__gong__resetStageType(&stage.ElevatedCurvePropertyTypes, &stage.ElevatedCurvePropertyTypes_mapString, &stage.ElevatedCurvePropertyType_stagedOrder, &stage.ElevatedCurvePropertyTypeOrder)

	__gong__resetStageType(&stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_mapString, &stage.ElevatedCurveType_stagedOrder, &stage.ElevatedCurveTypeOrder)

	__gong__resetStageType(&stage.ElevatedPointPropertyTypes, &stage.ElevatedPointPropertyTypes_mapString, &stage.ElevatedPointPropertyType_stagedOrder, &stage.ElevatedPointPropertyTypeOrder)

	__gong__resetStageType(&stage.ElevatedPointTypes, &stage.ElevatedPointTypes_mapString, &stage.ElevatedPointType_stagedOrder, &stage.ElevatedPointTypeOrder)

	__gong__resetStageType(&stage.ElevatedSurfacePropertyTypes, &stage.ElevatedSurfacePropertyTypes_mapString, &stage.ElevatedSurfacePropertyType_stagedOrder, &stage.ElevatedSurfacePropertyTypeOrder)

	__gong__resetStageType(&stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_mapString, &stage.ElevatedSurfaceType_stagedOrder, &stage.ElevatedSurfaceTypeOrder)

	__gong__resetStageType(&stage.GroundLightSystemPropertyTypes, &stage.GroundLightSystemPropertyTypes_mapString, &stage.GroundLightSystemPropertyType_stagedOrder, &stage.GroundLightSystemPropertyTypeOrder)

	__gong__resetStageType(&stage.LightElementPropertyTypes, &stage.LightElementPropertyTypes_mapString, &stage.LightElementPropertyType_stagedOrder, &stage.LightElementPropertyTypeOrder)

	__gong__resetStageType(&stage.LightElementStatusPropertyTypes, &stage.LightElementStatusPropertyTypes_mapString, &stage.LightElementStatusPropertyType_stagedOrder, &stage.LightElementStatusPropertyTypeOrder)

	__gong__resetStageType(&stage.LightElementStatusTypes, &stage.LightElementStatusTypes_mapString, &stage.LightElementStatusType_stagedOrder, &stage.LightElementStatusTypeOrder)

	__gong__resetStageType(&stage.LightElementTypes, &stage.LightElementTypes_mapString, &stage.LightElementType_stagedOrder, &stage.LightElementTypeOrder)

	__gong__resetStageType(&stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_mapString, &stage.LinguisticNotePropertyType_stagedOrder, &stage.LinguisticNotePropertyTypeOrder)

	__gong__resetStageType(&stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_mapString, &stage.LinguisticNoteType_stagedOrder, &stage.LinguisticNoteTypeOrder)

	__gong__resetStageType(&stage.MarkerBeaconPropertyTypes, &stage.MarkerBeaconPropertyTypes_mapString, &stage.MarkerBeaconPropertyType_stagedOrder, &stage.MarkerBeaconPropertyTypeOrder)

	__gong__resetStageType(&stage.NavaidEquipmentPropertyTypes, &stage.NavaidEquipmentPropertyTypes_mapString, &stage.NavaidEquipmentPropertyType_stagedOrder, &stage.NavaidEquipmentPropertyTypeOrder)

	__gong__resetStageType(&stage.NotePropertyTypes, &stage.NotePropertyTypes_mapString, &stage.NotePropertyType_stagedOrder, &stage.NotePropertyTypeOrder)

	__gong__resetStageType(&stage.NoteTypes, &stage.NoteTypes_mapString, &stage.NoteType_stagedOrder, &stage.NoteTypeOrder)

	__gong__resetStageType(&stage.ObstacleAreaTimeSlicePropertyTypes, &stage.ObstacleAreaTimeSlicePropertyTypes_mapString, &stage.ObstacleAreaTimeSlicePropertyType_stagedOrder, &stage.ObstacleAreaTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.ObstacleAreaTimeSliceTypes, &stage.ObstacleAreaTimeSliceTypes_mapString, &stage.ObstacleAreaTimeSliceType_stagedOrder, &stage.ObstacleAreaTimeSliceTypeOrder)

	__gong__resetStageType(&stage.ObstacleAreaTypes, &stage.ObstacleAreaTypes_mapString, &stage.ObstacleAreaType_stagedOrder, &stage.ObstacleAreaTypeOrder)

	__gong__resetStageType(&stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_mapString, &stage.OrganisationAuthorityPropertyType_stagedOrder, &stage.OrganisationAuthorityPropertyTypeOrder)

	__gong__resetStageType(&stage.PassengerServicePropertyTypes, &stage.PassengerServicePropertyTypes_mapString, &stage.PassengerServicePropertyType_stagedOrder, &stage.PassengerServicePropertyTypeOrder)

	__gong__resetStageType(&stage.PointTypes, &stage.PointTypes_mapString, &stage.PointType_stagedOrder, &stage.PointTypeOrder)

	__gong__resetStageType(&stage.RunwayDirectionPropertyTypes, &stage.RunwayDirectionPropertyTypes_mapString, &stage.RunwayDirectionPropertyType_stagedOrder, &stage.RunwayDirectionPropertyTypeOrder)

	__gong__resetStageType(&stage.ServicePropertyTypes, &stage.ServicePropertyTypes_mapString, &stage.ServicePropertyType_stagedOrder, &stage.ServicePropertyTypeOrder)

	__gong__resetStageType(&stage.SpecialNavigationStationPropertyTypes, &stage.SpecialNavigationStationPropertyTypes_mapString, &stage.SpecialNavigationStationPropertyType_stagedOrder, &stage.SpecialNavigationStationPropertyTypeOrder)

	__gong__resetStageType(&stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_mapString, &stage.SurfacePropertyType_stagedOrder, &stage.SurfacePropertyTypeOrder)

	__gong__resetStageType(&stage.SurfaceTypes, &stage.SurfaceTypes_mapString, &stage.SurfaceType_stagedOrder, &stage.SurfaceTypeOrder)

	__gong__resetStageType(&stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_mapString, &stage.TimesheetPropertyType_stagedOrder, &stage.TimesheetPropertyTypeOrder)

	__gong__resetStageType(&stage.TimesheetTypes, &stage.TimesheetTypes_mapString, &stage.TimesheetType_stagedOrder, &stage.TimesheetTypeOrder)

	__gong__resetStageType(&stage.UnitPropertyTypes, &stage.UnitPropertyTypes_mapString, &stage.UnitPropertyType_stagedOrder, &stage.UnitPropertyTypeOrder)

	__gong__resetStageType(&stage.VerticalStructureLightingStatusPropertyTypes, &stage.VerticalStructureLightingStatusPropertyTypes_mapString, &stage.VerticalStructureLightingStatusPropertyType_stagedOrder, &stage.VerticalStructureLightingStatusPropertyTypeOrder)

	__gong__resetStageType(&stage.VerticalStructureLightingStatusTypes, &stage.VerticalStructureLightingStatusTypes_mapString, &stage.VerticalStructureLightingStatusType_stagedOrder, &stage.VerticalStructureLightingStatusTypeOrder)

	__gong__resetStageType(&stage.VerticalStructurePartPropertyTypes, &stage.VerticalStructurePartPropertyTypes_mapString, &stage.VerticalStructurePartPropertyType_stagedOrder, &stage.VerticalStructurePartPropertyTypeOrder)

	__gong__resetStageType(&stage.VerticalStructurePartTypes, &stage.VerticalStructurePartTypes_mapString, &stage.VerticalStructurePartType_stagedOrder, &stage.VerticalStructurePartTypeOrder)

	__gong__resetStageType(&stage.VerticalStructurePropertyTypes, &stage.VerticalStructurePropertyTypes_mapString, &stage.VerticalStructurePropertyType_stagedOrder, &stage.VerticalStructurePropertyTypeOrder)

	__gong__resetStageType(&stage.VerticalStructureTimeSlicePropertyTypes, &stage.VerticalStructureTimeSlicePropertyTypes_mapString, &stage.VerticalStructureTimeSlicePropertyType_stagedOrder, &stage.VerticalStructureTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.VerticalStructureTimeSliceTypes, &stage.VerticalStructureTimeSliceTypes_mapString, &stage.VerticalStructureTimeSliceType_stagedOrder, &stage.VerticalStructureTimeSliceTypeOrder)

	__gong__resetStageType(&stage.VerticalStructureTypes, &stage.VerticalStructureTypes_mapString, &stage.VerticalStructureType_stagedOrder, &stage.VerticalStructureTypeOrder)

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
func (A_extension) GongGetAssociationName() any {
	return &A_extension{
	}
}

func (A_extension_1) GongGetAssociationName() any {
	return &A_extension_1{
	}
}

func (A_extension_10) GongGetAssociationName() any {
	return &A_extension_10{
	}
}

func (A_extension_11) GongGetAssociationName() any {
	return &A_extension_11{
	}
}

func (A_extension_12) GongGetAssociationName() any {
	return &A_extension_12{
	}
}

func (A_extension_2) GongGetAssociationName() any {
	return &A_extension_2{
	}
}

func (A_extension_3) GongGetAssociationName() any {
	return &A_extension_3{
	}
}

func (A_extension_4) GongGetAssociationName() any {
	return &A_extension_4{
	}
}

func (A_extension_5) GongGetAssociationName() any {
	return &A_extension_5{
	}
}

func (A_extension_6) GongGetAssociationName() any {
	return &A_extension_6{
	}
}

func (A_extension_7) GongGetAssociationName() any {
	return &A_extension_7{
	}
}

func (A_extension_8) GongGetAssociationName() any {
	return &A_extension_8{
	}
}

func (A_extension_9) GongGetAssociationName() any {
	return &A_extension_9{
	}
}

func (AbstractPropertiesWithScheduleType) GongGetAssociationName() any {
	return &AbstractPropertiesWithScheduleType{
	}
}

func (AeronauticalGroundLightTimeSlicePropertyType) GongGetAssociationName() any {
	return &AeronauticalGroundLightTimeSlicePropertyType{
			AeronauticalGroundLightTimeSlice: &AeronauticalGroundLightTimeSliceType{Name: "AeronauticalGroundLightTimeSlice"},
	}
}

func (AeronauticalGroundLightTimeSliceType) GongGetAssociationName() any {
	return &AeronauticalGroundLightTimeSliceType{
			StructureBeacon: &VerticalStructurePropertyType{Name: "StructureBeacon"},
			AerodromeBeacon: &AirportHeliportPropertyType{Name: "AerodromeBeacon"},
			Location: &ElevatedPointPropertyType{Name: "Location"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_3{{Name: "Extension"}},
	}
}

func (AeronauticalGroundLightType) GongGetAssociationName() any {
	return &AeronauticalGroundLightType{
			TimeSlice: []*AeronauticalGroundLightTimeSlicePropertyType{{Name: "TimeSlice"}},
	}
}

func (AirportHeliportPropertyType) GongGetAssociationName() any {
	return &AirportHeliportPropertyType{
	}
}

func (CurveType) GongGetAssociationName() any {
	return &CurveType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
	}
}

func (ElevatedCurvePropertyType) GongGetAssociationName() any {
	return &ElevatedCurvePropertyType{
			ElevatedCurve: &ElevatedCurveType{Name: "ElevatedCurve"},
	}
}

func (ElevatedCurveType) GongGetAssociationName() any {
	return &ElevatedCurveType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension{{Name: "Extension"}},
	}
}

func (ElevatedPointPropertyType) GongGetAssociationName() any {
	return &ElevatedPointPropertyType{
			ElevatedPoint: &ElevatedPointType{Name: "ElevatedPoint"},
	}
}

func (ElevatedPointType) GongGetAssociationName() any {
	return &ElevatedPointType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_1{{Name: "Extension"}},
	}
}

func (ElevatedSurfacePropertyType) GongGetAssociationName() any {
	return &ElevatedSurfacePropertyType{
			ElevatedSurface: &ElevatedSurfaceType{Name: "ElevatedSurface"},
	}
}

func (ElevatedSurfaceType) GongGetAssociationName() any {
	return &ElevatedSurfaceType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_2{{Name: "Extension"}},
	}
}

func (GroundLightSystemPropertyType) GongGetAssociationName() any {
	return &GroundLightSystemPropertyType{
	}
}

func (LightElementPropertyType) GongGetAssociationName() any {
	return &LightElementPropertyType{
			LightElement: &LightElementType{Name: "LightElement"},
	}
}

func (LightElementStatusPropertyType) GongGetAssociationName() any {
	return &LightElementStatusPropertyType{
			LightElementStatus: &LightElementStatusType{Name: "LightElementStatus"},
	}
}

func (LightElementStatusType) GongGetAssociationName() any {
	return &LightElementStatusType{
			TimeInterval: []*TimesheetPropertyType{{Name: "TimeInterval"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			SpecialDateAuthority: []*OrganisationAuthorityPropertyType{{Name: "SpecialDateAuthority"}},
			Extension: []*A_extension_11{{Name: "Extension"}},
	}
}

func (LightElementType) GongGetAssociationName() any {
	return &LightElementType{
			Location: &ElevatedPointPropertyType{Name: "Location"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Availability: []*LightElementStatusPropertyType{{Name: "Availability"}},
			Extension: []*A_extension_10{{Name: "Extension"}},
	}
}

func (LinguisticNotePropertyType) GongGetAssociationName() any {
	return &LinguisticNotePropertyType{
			LinguisticNote: &LinguisticNoteType{Name: "LinguisticNote"},
	}
}

func (LinguisticNoteType) GongGetAssociationName() any {
	return &LinguisticNoteType{
			Extension: []*A_extension_4{{Name: "Extension"}},
	}
}

func (MarkerBeaconPropertyType) GongGetAssociationName() any {
	return &MarkerBeaconPropertyType{
	}
}

func (NavaidEquipmentPropertyType) GongGetAssociationName() any {
	return &NavaidEquipmentPropertyType{
	}
}

func (NotePropertyType) GongGetAssociationName() any {
	return &NotePropertyType{
			Note: &NoteType{Name: "Note"},
	}
}

func (NoteType) GongGetAssociationName() any {
	return &NoteType{
			TranslatedNote: []*LinguisticNotePropertyType{{Name: "TranslatedNote"}},
			Extension: []*A_extension_5{{Name: "Extension"}},
	}
}

func (ObstacleAreaTimeSlicePropertyType) GongGetAssociationName() any {
	return &ObstacleAreaTimeSlicePropertyType{
			ObstacleAreaTimeSlice: &ObstacleAreaTimeSliceType{Name: "ObstacleAreaTimeSlice"},
	}
}

func (ObstacleAreaTimeSliceType) GongGetAssociationName() any {
	return &ObstacleAreaTimeSliceType{
			Reference_ownerOrganisation: &OrganisationAuthorityPropertyType{Name: "Reference_ownerOrganisation"},
			Reference_ownerRunway: &RunwayDirectionPropertyType{Name: "Reference_ownerRunway"},
			Reference_ownerAirport: &AirportHeliportPropertyType{Name: "Reference_ownerAirport"},
			SurfaceExtent: &SurfacePropertyType{Name: "SurfaceExtent"},
			Obstacle: []*VerticalStructurePropertyType{{Name: "Obstacle"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_6{{Name: "Extension"}},
	}
}

func (ObstacleAreaType) GongGetAssociationName() any {
	return &ObstacleAreaType{
			TimeSlice: []*ObstacleAreaTimeSlicePropertyType{{Name: "TimeSlice"}},
	}
}

func (OrganisationAuthorityPropertyType) GongGetAssociationName() any {
	return &OrganisationAuthorityPropertyType{
	}
}

func (PassengerServicePropertyType) GongGetAssociationName() any {
	return &PassengerServicePropertyType{
	}
}

func (PointType) GongGetAssociationName() any {
	return &PointType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
	}
}

func (RunwayDirectionPropertyType) GongGetAssociationName() any {
	return &RunwayDirectionPropertyType{
	}
}

func (ServicePropertyType) GongGetAssociationName() any {
	return &ServicePropertyType{
	}
}

func (SpecialNavigationStationPropertyType) GongGetAssociationName() any {
	return &SpecialNavigationStationPropertyType{
	}
}

func (SurfacePropertyType) GongGetAssociationName() any {
	return &SurfacePropertyType{
			Surface: &SurfaceType{Name: "Surface"},
			ElevatedSurface: &ElevatedSurfaceType{Name: "ElevatedSurface"},
	}
}

func (SurfaceType) GongGetAssociationName() any {
	return &SurfaceType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
	}
}

func (TimesheetPropertyType) GongGetAssociationName() any {
	return &TimesheetPropertyType{
			Timesheet: &TimesheetType{Name: "Timesheet"},
	}
}

func (TimesheetType) GongGetAssociationName() any {
	return &TimesheetType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_12{{Name: "Extension"}},
	}
}

func (UnitPropertyType) GongGetAssociationName() any {
	return &UnitPropertyType{
	}
}

func (VerticalStructureLightingStatusPropertyType) GongGetAssociationName() any {
	return &VerticalStructureLightingStatusPropertyType{
			VerticalStructureLightingStatus: &VerticalStructureLightingStatusType{Name: "VerticalStructureLightingStatus"},
	}
}

func (VerticalStructureLightingStatusType) GongGetAssociationName() any {
	return &VerticalStructureLightingStatusType{
			TimeInterval: []*TimesheetPropertyType{{Name: "TimeInterval"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			SpecialDateAuthority: []*OrganisationAuthorityPropertyType{{Name: "SpecialDateAuthority"}},
			Extension: []*A_extension_8{{Name: "Extension"}},
	}
}

func (VerticalStructurePartPropertyType) GongGetAssociationName() any {
	return &VerticalStructurePartPropertyType{
			VerticalStructurePart: &VerticalStructurePartType{Name: "VerticalStructurePart"},
	}
}

func (VerticalStructurePartType) GongGetAssociationName() any {
	return &VerticalStructurePartType{
			TimeInterval: []*TimesheetPropertyType{{Name: "TimeInterval"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			SpecialDateAuthority: []*OrganisationAuthorityPropertyType{{Name: "SpecialDateAuthority"}},
			HorizontalProjection_surfaceExtent: &ElevatedSurfacePropertyType{Name: "HorizontalProjection_surfaceExtent"},
			HorizontalProjection_linearExtent: &ElevatedCurvePropertyType{Name: "HorizontalProjection_linearExtent"},
			HorizontalProjection_location: &ElevatedPointPropertyType{Name: "HorizontalProjection_location"},
			Lighting: []*LightElementPropertyType{{Name: "Lighting"}},
			Extension: []*A_extension_9{{Name: "Extension"}},
	}
}

func (VerticalStructurePropertyType) GongGetAssociationName() any {
	return &VerticalStructurePropertyType{
	}
}

func (VerticalStructureTimeSlicePropertyType) GongGetAssociationName() any {
	return &VerticalStructureTimeSlicePropertyType{
			VerticalStructureTimeSlice: &VerticalStructureTimeSliceType{Name: "VerticalStructureTimeSlice"},
	}
}

func (VerticalStructureTimeSliceType) GongGetAssociationName() any {
	return &VerticalStructureTimeSliceType{
			Marker: &MarkerBeaconPropertyType{Name: "Marker"},
			Part: []*VerticalStructurePartPropertyType{{Name: "Part"}},
			HostedPassengerService: []*PassengerServicePropertyType{{Name: "HostedPassengerService"}},
			SupportedGroundLight: []*GroundLightSystemPropertyType{{Name: "SupportedGroundLight"}},
			HostedNavaidEquipment: []*NavaidEquipmentPropertyType{{Name: "HostedNavaidEquipment"}},
			HostedSpecialNavStation: []*SpecialNavigationStationPropertyType{{Name: "HostedSpecialNavStation"}},
			HostedUnit: []*UnitPropertyType{{Name: "HostedUnit"}},
			HostedOrganisation: []*OrganisationAuthorityPropertyType{{Name: "HostedOrganisation"}},
			SupportedService: []*ServicePropertyType{{Name: "SupportedService"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			LightingAvailability: []*VerticalStructureLightingStatusPropertyType{{Name: "LightingAvailability"}},
			Extension: []*A_extension_7{{Name: "Extension"}},
	}
}

func (VerticalStructureType) GongGetAssociationName() any {
	return &VerticalStructureType{
			TimeSlice: []*VerticalStructureTimeSlicePropertyType{{Name: "TimeSlice"}},
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
	// reverse maps of direct associations of A_extension
	case A_extension:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_1
	case A_extension_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_10
	case A_extension_10:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_11
	case A_extension_11:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_12
	case A_extension_12:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_2
	case A_extension_2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_3
	case A_extension_3:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_4
	case A_extension_4:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_5
	case A_extension_5:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_6
	case A_extension_6:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_7
	case A_extension_7:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_8
	case A_extension_8:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_9
	case A_extension_9:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractPropertiesWithScheduleType
	case AbstractPropertiesWithScheduleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AeronauticalGroundLightTimeSlicePropertyType
	case AeronauticalGroundLightTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AeronauticalGroundLightTimeSlice":
			res := make(map[*AeronauticalGroundLightTimeSliceType][]*AeronauticalGroundLightTimeSlicePropertyType)
			for aeronauticalgroundlighttimeslicepropertytype := range stage.AeronauticalGroundLightTimeSlicePropertyTypes {
				if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
					aeronauticalgroundlighttimeslicetype_ := aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice
					var aeronauticalgroundlighttimeslicepropertytypes []*AeronauticalGroundLightTimeSlicePropertyType
					_, ok := res[aeronauticalgroundlighttimeslicetype_]
					if ok {
						aeronauticalgroundlighttimeslicepropertytypes = res[aeronauticalgroundlighttimeslicetype_]
					} else {
						aeronauticalgroundlighttimeslicepropertytypes = make([]*AeronauticalGroundLightTimeSlicePropertyType, 0)
					}
					aeronauticalgroundlighttimeslicepropertytypes = append(aeronauticalgroundlighttimeslicepropertytypes, aeronauticalgroundlighttimeslicepropertytype)
					res[aeronauticalgroundlighttimeslicetype_] = aeronauticalgroundlighttimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AeronauticalGroundLightTimeSliceType
	case AeronauticalGroundLightTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "StructureBeacon":
			res := make(map[*VerticalStructurePropertyType][]*AeronauticalGroundLightTimeSliceType)
			for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
				if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
					verticalstructurepropertytype_ := aeronauticalgroundlighttimeslicetype.StructureBeacon
					var aeronauticalgroundlighttimeslicetypes []*AeronauticalGroundLightTimeSliceType
					_, ok := res[verticalstructurepropertytype_]
					if ok {
						aeronauticalgroundlighttimeslicetypes = res[verticalstructurepropertytype_]
					} else {
						aeronauticalgroundlighttimeslicetypes = make([]*AeronauticalGroundLightTimeSliceType, 0)
					}
					aeronauticalgroundlighttimeslicetypes = append(aeronauticalgroundlighttimeslicetypes, aeronauticalgroundlighttimeslicetype)
					res[verticalstructurepropertytype_] = aeronauticalgroundlighttimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "AerodromeBeacon":
			res := make(map[*AirportHeliportPropertyType][]*AeronauticalGroundLightTimeSliceType)
			for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
				if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
					airportheliportpropertytype_ := aeronauticalgroundlighttimeslicetype.AerodromeBeacon
					var aeronauticalgroundlighttimeslicetypes []*AeronauticalGroundLightTimeSliceType
					_, ok := res[airportheliportpropertytype_]
					if ok {
						aeronauticalgroundlighttimeslicetypes = res[airportheliportpropertytype_]
					} else {
						aeronauticalgroundlighttimeslicetypes = make([]*AeronauticalGroundLightTimeSliceType, 0)
					}
					aeronauticalgroundlighttimeslicetypes = append(aeronauticalgroundlighttimeslicetypes, aeronauticalgroundlighttimeslicetype)
					res[airportheliportpropertytype_] = aeronauticalgroundlighttimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Location":
			res := make(map[*ElevatedPointPropertyType][]*AeronauticalGroundLightTimeSliceType)
			for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
				if aeronauticalgroundlighttimeslicetype.Location != nil {
					elevatedpointpropertytype_ := aeronauticalgroundlighttimeslicetype.Location
					var aeronauticalgroundlighttimeslicetypes []*AeronauticalGroundLightTimeSliceType
					_, ok := res[elevatedpointpropertytype_]
					if ok {
						aeronauticalgroundlighttimeslicetypes = res[elevatedpointpropertytype_]
					} else {
						aeronauticalgroundlighttimeslicetypes = make([]*AeronauticalGroundLightTimeSliceType, 0)
					}
					aeronauticalgroundlighttimeslicetypes = append(aeronauticalgroundlighttimeslicetypes, aeronauticalgroundlighttimeslicetype)
					res[elevatedpointpropertytype_] = aeronauticalgroundlighttimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AeronauticalGroundLightType
	case AeronauticalGroundLightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirportHeliportPropertyType
	case AirportHeliportPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CurveType
	case CurveType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedCurvePropertyType
	case ElevatedCurvePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "ElevatedCurve":
			res := make(map[*ElevatedCurveType][]*ElevatedCurvePropertyType)
			for elevatedcurvepropertytype := range stage.ElevatedCurvePropertyTypes {
				if elevatedcurvepropertytype.ElevatedCurve != nil {
					elevatedcurvetype_ := elevatedcurvepropertytype.ElevatedCurve
					var elevatedcurvepropertytypes []*ElevatedCurvePropertyType
					_, ok := res[elevatedcurvetype_]
					if ok {
						elevatedcurvepropertytypes = res[elevatedcurvetype_]
					} else {
						elevatedcurvepropertytypes = make([]*ElevatedCurvePropertyType, 0)
					}
					elevatedcurvepropertytypes = append(elevatedcurvepropertytypes, elevatedcurvepropertytype)
					res[elevatedcurvetype_] = elevatedcurvepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedCurveType
	case ElevatedCurveType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedPointPropertyType
	case ElevatedPointPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "ElevatedPoint":
			res := make(map[*ElevatedPointType][]*ElevatedPointPropertyType)
			for elevatedpointpropertytype := range stage.ElevatedPointPropertyTypes {
				if elevatedpointpropertytype.ElevatedPoint != nil {
					elevatedpointtype_ := elevatedpointpropertytype.ElevatedPoint
					var elevatedpointpropertytypes []*ElevatedPointPropertyType
					_, ok := res[elevatedpointtype_]
					if ok {
						elevatedpointpropertytypes = res[elevatedpointtype_]
					} else {
						elevatedpointpropertytypes = make([]*ElevatedPointPropertyType, 0)
					}
					elevatedpointpropertytypes = append(elevatedpointpropertytypes, elevatedpointpropertytype)
					res[elevatedpointtype_] = elevatedpointpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedPointType
	case ElevatedPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedSurfacePropertyType
	case ElevatedSurfacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "ElevatedSurface":
			res := make(map[*ElevatedSurfaceType][]*ElevatedSurfacePropertyType)
			for elevatedsurfacepropertytype := range stage.ElevatedSurfacePropertyTypes {
				if elevatedsurfacepropertytype.ElevatedSurface != nil {
					elevatedsurfacetype_ := elevatedsurfacepropertytype.ElevatedSurface
					var elevatedsurfacepropertytypes []*ElevatedSurfacePropertyType
					_, ok := res[elevatedsurfacetype_]
					if ok {
						elevatedsurfacepropertytypes = res[elevatedsurfacetype_]
					} else {
						elevatedsurfacepropertytypes = make([]*ElevatedSurfacePropertyType, 0)
					}
					elevatedsurfacepropertytypes = append(elevatedsurfacepropertytypes, elevatedsurfacepropertytype)
					res[elevatedsurfacetype_] = elevatedsurfacepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedSurfaceType
	case ElevatedSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroundLightSystemPropertyType
	case GroundLightSystemPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LightElementPropertyType
	case LightElementPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "LightElement":
			res := make(map[*LightElementType][]*LightElementPropertyType)
			for lightelementpropertytype := range stage.LightElementPropertyTypes {
				if lightelementpropertytype.LightElement != nil {
					lightelementtype_ := lightelementpropertytype.LightElement
					var lightelementpropertytypes []*LightElementPropertyType
					_, ok := res[lightelementtype_]
					if ok {
						lightelementpropertytypes = res[lightelementtype_]
					} else {
						lightelementpropertytypes = make([]*LightElementPropertyType, 0)
					}
					lightelementpropertytypes = append(lightelementpropertytypes, lightelementpropertytype)
					res[lightelementtype_] = lightelementpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LightElementStatusPropertyType
	case LightElementStatusPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "LightElementStatus":
			res := make(map[*LightElementStatusType][]*LightElementStatusPropertyType)
			for lightelementstatuspropertytype := range stage.LightElementStatusPropertyTypes {
				if lightelementstatuspropertytype.LightElementStatus != nil {
					lightelementstatustype_ := lightelementstatuspropertytype.LightElementStatus
					var lightelementstatuspropertytypes []*LightElementStatusPropertyType
					_, ok := res[lightelementstatustype_]
					if ok {
						lightelementstatuspropertytypes = res[lightelementstatustype_]
					} else {
						lightelementstatuspropertytypes = make([]*LightElementStatusPropertyType, 0)
					}
					lightelementstatuspropertytypes = append(lightelementstatuspropertytypes, lightelementstatuspropertytype)
					res[lightelementstatustype_] = lightelementstatuspropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LightElementStatusType
	case LightElementStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LightElementType
	case LightElementType:
		switch fieldname {
		// insertion point for per direct association field
		case "Location":
			res := make(map[*ElevatedPointPropertyType][]*LightElementType)
			for lightelementtype := range stage.LightElementTypes {
				if lightelementtype.Location != nil {
					elevatedpointpropertytype_ := lightelementtype.Location
					var lightelementtypes []*LightElementType
					_, ok := res[elevatedpointpropertytype_]
					if ok {
						lightelementtypes = res[elevatedpointpropertytype_]
					} else {
						lightelementtypes = make([]*LightElementType, 0)
					}
					lightelementtypes = append(lightelementtypes, lightelementtype)
					res[elevatedpointpropertytype_] = lightelementtypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinguisticNotePropertyType
	case LinguisticNotePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "LinguisticNote":
			res := make(map[*LinguisticNoteType][]*LinguisticNotePropertyType)
			for linguisticnotepropertytype := range stage.LinguisticNotePropertyTypes {
				if linguisticnotepropertytype.LinguisticNote != nil {
					linguisticnotetype_ := linguisticnotepropertytype.LinguisticNote
					var linguisticnotepropertytypes []*LinguisticNotePropertyType
					_, ok := res[linguisticnotetype_]
					if ok {
						linguisticnotepropertytypes = res[linguisticnotetype_]
					} else {
						linguisticnotepropertytypes = make([]*LinguisticNotePropertyType, 0)
					}
					linguisticnotepropertytypes = append(linguisticnotepropertytypes, linguisticnotepropertytype)
					res[linguisticnotetype_] = linguisticnotepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinguisticNoteType
	case LinguisticNoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MarkerBeaconPropertyType
	case MarkerBeaconPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NavaidEquipmentPropertyType
	case NavaidEquipmentPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NotePropertyType
	case NotePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*NoteType][]*NotePropertyType)
			for notepropertytype := range stage.NotePropertyTypes {
				if notepropertytype.Note != nil {
					notetype_ := notepropertytype.Note
					var notepropertytypes []*NotePropertyType
					_, ok := res[notetype_]
					if ok {
						notepropertytypes = res[notetype_]
					} else {
						notepropertytypes = make([]*NotePropertyType, 0)
					}
					notepropertytypes = append(notepropertytypes, notepropertytype)
					res[notetype_] = notepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteType
	case NoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ObstacleAreaTimeSlicePropertyType
	case ObstacleAreaTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "ObstacleAreaTimeSlice":
			res := make(map[*ObstacleAreaTimeSliceType][]*ObstacleAreaTimeSlicePropertyType)
			for obstacleareatimeslicepropertytype := range stage.ObstacleAreaTimeSlicePropertyTypes {
				if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
					obstacleareatimeslicetype_ := obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice
					var obstacleareatimeslicepropertytypes []*ObstacleAreaTimeSlicePropertyType
					_, ok := res[obstacleareatimeslicetype_]
					if ok {
						obstacleareatimeslicepropertytypes = res[obstacleareatimeslicetype_]
					} else {
						obstacleareatimeslicepropertytypes = make([]*ObstacleAreaTimeSlicePropertyType, 0)
					}
					obstacleareatimeslicepropertytypes = append(obstacleareatimeslicepropertytypes, obstacleareatimeslicepropertytype)
					res[obstacleareatimeslicetype_] = obstacleareatimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ObstacleAreaTimeSliceType
	case ObstacleAreaTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Reference_ownerOrganisation":
			res := make(map[*OrganisationAuthorityPropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
					organisationauthoritypropertytype_ := obstacleareatimeslicetype.Reference_ownerOrganisation
					var obstacleareatimeslicetypes []*ObstacleAreaTimeSliceType
					_, ok := res[organisationauthoritypropertytype_]
					if ok {
						obstacleareatimeslicetypes = res[organisationauthoritypropertytype_]
					} else {
						obstacleareatimeslicetypes = make([]*ObstacleAreaTimeSliceType, 0)
					}
					obstacleareatimeslicetypes = append(obstacleareatimeslicetypes, obstacleareatimeslicetype)
					res[organisationauthoritypropertytype_] = obstacleareatimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Reference_ownerRunway":
			res := make(map[*RunwayDirectionPropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				if obstacleareatimeslicetype.Reference_ownerRunway != nil {
					runwaydirectionpropertytype_ := obstacleareatimeslicetype.Reference_ownerRunway
					var obstacleareatimeslicetypes []*ObstacleAreaTimeSliceType
					_, ok := res[runwaydirectionpropertytype_]
					if ok {
						obstacleareatimeslicetypes = res[runwaydirectionpropertytype_]
					} else {
						obstacleareatimeslicetypes = make([]*ObstacleAreaTimeSliceType, 0)
					}
					obstacleareatimeslicetypes = append(obstacleareatimeslicetypes, obstacleareatimeslicetype)
					res[runwaydirectionpropertytype_] = obstacleareatimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Reference_ownerAirport":
			res := make(map[*AirportHeliportPropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				if obstacleareatimeslicetype.Reference_ownerAirport != nil {
					airportheliportpropertytype_ := obstacleareatimeslicetype.Reference_ownerAirport
					var obstacleareatimeslicetypes []*ObstacleAreaTimeSliceType
					_, ok := res[airportheliportpropertytype_]
					if ok {
						obstacleareatimeslicetypes = res[airportheliportpropertytype_]
					} else {
						obstacleareatimeslicetypes = make([]*ObstacleAreaTimeSliceType, 0)
					}
					obstacleareatimeslicetypes = append(obstacleareatimeslicetypes, obstacleareatimeslicetype)
					res[airportheliportpropertytype_] = obstacleareatimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SurfaceExtent":
			res := make(map[*SurfacePropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				if obstacleareatimeslicetype.SurfaceExtent != nil {
					surfacepropertytype_ := obstacleareatimeslicetype.SurfaceExtent
					var obstacleareatimeslicetypes []*ObstacleAreaTimeSliceType
					_, ok := res[surfacepropertytype_]
					if ok {
						obstacleareatimeslicetypes = res[surfacepropertytype_]
					} else {
						obstacleareatimeslicetypes = make([]*ObstacleAreaTimeSliceType, 0)
					}
					obstacleareatimeslicetypes = append(obstacleareatimeslicetypes, obstacleareatimeslicetype)
					res[surfacepropertytype_] = obstacleareatimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ObstacleAreaType
	case ObstacleAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OrganisationAuthorityPropertyType
	case OrganisationAuthorityPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PassengerServicePropertyType
	case PassengerServicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointType
	case PointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RunwayDirectionPropertyType
	case RunwayDirectionPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ServicePropertyType
	case ServicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpecialNavigationStationPropertyType
	case SpecialNavigationStationPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SurfacePropertyType
	case SurfacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Surface":
			res := make(map[*SurfaceType][]*SurfacePropertyType)
			for surfacepropertytype := range stage.SurfacePropertyTypes {
				if surfacepropertytype.Surface != nil {
					surfacetype_ := surfacepropertytype.Surface
					var surfacepropertytypes []*SurfacePropertyType
					_, ok := res[surfacetype_]
					if ok {
						surfacepropertytypes = res[surfacetype_]
					} else {
						surfacepropertytypes = make([]*SurfacePropertyType, 0)
					}
					surfacepropertytypes = append(surfacepropertytypes, surfacepropertytype)
					res[surfacetype_] = surfacepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ElevatedSurface":
			res := make(map[*ElevatedSurfaceType][]*SurfacePropertyType)
			for surfacepropertytype := range stage.SurfacePropertyTypes {
				if surfacepropertytype.ElevatedSurface != nil {
					elevatedsurfacetype_ := surfacepropertytype.ElevatedSurface
					var surfacepropertytypes []*SurfacePropertyType
					_, ok := res[elevatedsurfacetype_]
					if ok {
						surfacepropertytypes = res[elevatedsurfacetype_]
					} else {
						surfacepropertytypes = make([]*SurfacePropertyType, 0)
					}
					surfacepropertytypes = append(surfacepropertytypes, surfacepropertytype)
					res[elevatedsurfacetype_] = surfacepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SurfaceType
	case SurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimesheetPropertyType
	case TimesheetPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Timesheet":
			res := make(map[*TimesheetType][]*TimesheetPropertyType)
			for timesheetpropertytype := range stage.TimesheetPropertyTypes {
				if timesheetpropertytype.Timesheet != nil {
					timesheettype_ := timesheetpropertytype.Timesheet
					var timesheetpropertytypes []*TimesheetPropertyType
					_, ok := res[timesheettype_]
					if ok {
						timesheetpropertytypes = res[timesheettype_]
					} else {
						timesheetpropertytypes = make([]*TimesheetPropertyType, 0)
					}
					timesheetpropertytypes = append(timesheetpropertytypes, timesheetpropertytype)
					res[timesheettype_] = timesheetpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimesheetType
	case TimesheetType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UnitPropertyType
	case UnitPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureLightingStatusPropertyType
	case VerticalStructureLightingStatusPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "VerticalStructureLightingStatus":
			res := make(map[*VerticalStructureLightingStatusType][]*VerticalStructureLightingStatusPropertyType)
			for verticalstructurelightingstatuspropertytype := range stage.VerticalStructureLightingStatusPropertyTypes {
				if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
					verticalstructurelightingstatustype_ := verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus
					var verticalstructurelightingstatuspropertytypes []*VerticalStructureLightingStatusPropertyType
					_, ok := res[verticalstructurelightingstatustype_]
					if ok {
						verticalstructurelightingstatuspropertytypes = res[verticalstructurelightingstatustype_]
					} else {
						verticalstructurelightingstatuspropertytypes = make([]*VerticalStructureLightingStatusPropertyType, 0)
					}
					verticalstructurelightingstatuspropertytypes = append(verticalstructurelightingstatuspropertytypes, verticalstructurelightingstatuspropertytype)
					res[verticalstructurelightingstatustype_] = verticalstructurelightingstatuspropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructureLightingStatusType
	case VerticalStructureLightingStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructurePartPropertyType
	case VerticalStructurePartPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "VerticalStructurePart":
			res := make(map[*VerticalStructurePartType][]*VerticalStructurePartPropertyType)
			for verticalstructurepartpropertytype := range stage.VerticalStructurePartPropertyTypes {
				if verticalstructurepartpropertytype.VerticalStructurePart != nil {
					verticalstructureparttype_ := verticalstructurepartpropertytype.VerticalStructurePart
					var verticalstructurepartpropertytypes []*VerticalStructurePartPropertyType
					_, ok := res[verticalstructureparttype_]
					if ok {
						verticalstructurepartpropertytypes = res[verticalstructureparttype_]
					} else {
						verticalstructurepartpropertytypes = make([]*VerticalStructurePartPropertyType, 0)
					}
					verticalstructurepartpropertytypes = append(verticalstructurepartpropertytypes, verticalstructurepartpropertytype)
					res[verticalstructureparttype_] = verticalstructurepartpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructurePartType
	case VerticalStructurePartType:
		switch fieldname {
		// insertion point for per direct association field
		case "HorizontalProjection_surfaceExtent":
			res := make(map[*ElevatedSurfacePropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
					elevatedsurfacepropertytype_ := verticalstructureparttype.HorizontalProjection_surfaceExtent
					var verticalstructureparttypes []*VerticalStructurePartType
					_, ok := res[elevatedsurfacepropertytype_]
					if ok {
						verticalstructureparttypes = res[elevatedsurfacepropertytype_]
					} else {
						verticalstructureparttypes = make([]*VerticalStructurePartType, 0)
					}
					verticalstructureparttypes = append(verticalstructureparttypes, verticalstructureparttype)
					res[elevatedsurfacepropertytype_] = verticalstructureparttypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "HorizontalProjection_linearExtent":
			res := make(map[*ElevatedCurvePropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
					elevatedcurvepropertytype_ := verticalstructureparttype.HorizontalProjection_linearExtent
					var verticalstructureparttypes []*VerticalStructurePartType
					_, ok := res[elevatedcurvepropertytype_]
					if ok {
						verticalstructureparttypes = res[elevatedcurvepropertytype_]
					} else {
						verticalstructureparttypes = make([]*VerticalStructurePartType, 0)
					}
					verticalstructureparttypes = append(verticalstructureparttypes, verticalstructureparttype)
					res[elevatedcurvepropertytype_] = verticalstructureparttypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "HorizontalProjection_location":
			res := make(map[*ElevatedPointPropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				if verticalstructureparttype.HorizontalProjection_location != nil {
					elevatedpointpropertytype_ := verticalstructureparttype.HorizontalProjection_location
					var verticalstructureparttypes []*VerticalStructurePartType
					_, ok := res[elevatedpointpropertytype_]
					if ok {
						verticalstructureparttypes = res[elevatedpointpropertytype_]
					} else {
						verticalstructureparttypes = make([]*VerticalStructurePartType, 0)
					}
					verticalstructureparttypes = append(verticalstructureparttypes, verticalstructureparttype)
					res[elevatedpointpropertytype_] = verticalstructureparttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructurePropertyType
	case VerticalStructurePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureTimeSlicePropertyType
	case VerticalStructureTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "VerticalStructureTimeSlice":
			res := make(map[*VerticalStructureTimeSliceType][]*VerticalStructureTimeSlicePropertyType)
			for verticalstructuretimeslicepropertytype := range stage.VerticalStructureTimeSlicePropertyTypes {
				if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
					verticalstructuretimeslicetype_ := verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice
					var verticalstructuretimeslicepropertytypes []*VerticalStructureTimeSlicePropertyType
					_, ok := res[verticalstructuretimeslicetype_]
					if ok {
						verticalstructuretimeslicepropertytypes = res[verticalstructuretimeslicetype_]
					} else {
						verticalstructuretimeslicepropertytypes = make([]*VerticalStructureTimeSlicePropertyType, 0)
					}
					verticalstructuretimeslicepropertytypes = append(verticalstructuretimeslicepropertytypes, verticalstructuretimeslicepropertytype)
					res[verticalstructuretimeslicetype_] = verticalstructuretimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructureTimeSliceType
	case VerticalStructureTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Marker":
			res := make(map[*MarkerBeaconPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				if verticalstructuretimeslicetype.Marker != nil {
					markerbeaconpropertytype_ := verticalstructuretimeslicetype.Marker
					var verticalstructuretimeslicetypes []*VerticalStructureTimeSliceType
					_, ok := res[markerbeaconpropertytype_]
					if ok {
						verticalstructuretimeslicetypes = res[markerbeaconpropertytype_]
					} else {
						verticalstructuretimeslicetypes = make([]*VerticalStructureTimeSliceType, 0)
					}
					verticalstructuretimeslicetypes = append(verticalstructuretimeslicetypes, verticalstructuretimeslicetype)
					res[markerbeaconpropertytype_] = verticalstructuretimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructureType
	case VerticalStructureType:
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
	// reverse maps of direct associations of A_extension
	case A_extension:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_1
	case A_extension_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_10
	case A_extension_10:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_11
	case A_extension_11:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_12
	case A_extension_12:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_2
	case A_extension_2:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_3
	case A_extension_3:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_4
	case A_extension_4:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_5
	case A_extension_5:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_6
	case A_extension_6:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_7
	case A_extension_7:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_8
	case A_extension_8:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_9
	case A_extension_9:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AbstractPropertiesWithScheduleType
	case AbstractPropertiesWithScheduleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AeronauticalGroundLightTimeSlicePropertyType
	case AeronauticalGroundLightTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AeronauticalGroundLightTimeSliceType
	case AeronauticalGroundLightTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AeronauticalGroundLightTimeSliceType)
			for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
				for _, notepropertytype_ := range aeronauticalgroundlighttimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], aeronauticalgroundlighttimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_3][]*AeronauticalGroundLightTimeSliceType)
			for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
				for _, a_extension_3_ := range aeronauticalgroundlighttimeslicetype.Extension {
					res[a_extension_3_] = append(res[a_extension_3_], aeronauticalgroundlighttimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AeronauticalGroundLightType
	case AeronauticalGroundLightType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*AeronauticalGroundLightTimeSlicePropertyType][]*AeronauticalGroundLightType)
			for aeronauticalgroundlighttype := range stage.AeronauticalGroundLightTypes {
				for _, aeronauticalgroundlighttimeslicepropertytype_ := range aeronauticalgroundlighttype.TimeSlice {
					res[aeronauticalgroundlighttimeslicepropertytype_] = append(res[aeronauticalgroundlighttimeslicepropertytype_], aeronauticalgroundlighttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirportHeliportPropertyType
	case AirportHeliportPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CurveType
	case CurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*CurveType)
			for curvetype := range stage.CurveTypes {
				for _, notepropertytype_ := range curvetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], curvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedCurvePropertyType
	case ElevatedCurvePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedCurveType
	case ElevatedCurveType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*ElevatedCurveType)
			for elevatedcurvetype := range stage.ElevatedCurveTypes {
				for _, notepropertytype_ := range elevatedcurvetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], elevatedcurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension][]*ElevatedCurveType)
			for elevatedcurvetype := range stage.ElevatedCurveTypes {
				for _, a_extension_ := range elevatedcurvetype.Extension {
					res[a_extension_] = append(res[a_extension_], elevatedcurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedPointPropertyType
	case ElevatedPointPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedPointType
	case ElevatedPointType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*ElevatedPointType)
			for elevatedpointtype := range stage.ElevatedPointTypes {
				for _, notepropertytype_ := range elevatedpointtype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], elevatedpointtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_1][]*ElevatedPointType)
			for elevatedpointtype := range stage.ElevatedPointTypes {
				for _, a_extension_1_ := range elevatedpointtype.Extension {
					res[a_extension_1_] = append(res[a_extension_1_], elevatedpointtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ElevatedSurfacePropertyType
	case ElevatedSurfacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedSurfaceType
	case ElevatedSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*ElevatedSurfaceType)
			for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
				for _, notepropertytype_ := range elevatedsurfacetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], elevatedsurfacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_2][]*ElevatedSurfaceType)
			for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
				for _, a_extension_2_ := range elevatedsurfacetype.Extension {
					res[a_extension_2_] = append(res[a_extension_2_], elevatedsurfacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GroundLightSystemPropertyType
	case GroundLightSystemPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LightElementPropertyType
	case LightElementPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LightElementStatusPropertyType
	case LightElementStatusPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LightElementStatusType
	case LightElementStatusType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInterval":
			res := make(map[*TimesheetPropertyType][]*LightElementStatusType)
			for lightelementstatustype := range stage.LightElementStatusTypes {
				for _, timesheetpropertytype_ := range lightelementstatustype.TimeInterval {
					res[timesheetpropertytype_] = append(res[timesheetpropertytype_], lightelementstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*LightElementStatusType)
			for lightelementstatustype := range stage.LightElementStatusTypes {
				for _, notepropertytype_ := range lightelementstatustype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], lightelementstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpecialDateAuthority":
			res := make(map[*OrganisationAuthorityPropertyType][]*LightElementStatusType)
			for lightelementstatustype := range stage.LightElementStatusTypes {
				for _, organisationauthoritypropertytype_ := range lightelementstatustype.SpecialDateAuthority {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], lightelementstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_11][]*LightElementStatusType)
			for lightelementstatustype := range stage.LightElementStatusTypes {
				for _, a_extension_11_ := range lightelementstatustype.Extension {
					res[a_extension_11_] = append(res[a_extension_11_], lightelementstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LightElementType
	case LightElementType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*LightElementType)
			for lightelementtype := range stage.LightElementTypes {
				for _, notepropertytype_ := range lightelementtype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], lightelementtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Availability":
			res := make(map[*LightElementStatusPropertyType][]*LightElementType)
			for lightelementtype := range stage.LightElementTypes {
				for _, lightelementstatuspropertytype_ := range lightelementtype.Availability {
					res[lightelementstatuspropertytype_] = append(res[lightelementstatuspropertytype_], lightelementtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_10][]*LightElementType)
			for lightelementtype := range stage.LightElementTypes {
				for _, a_extension_10_ := range lightelementtype.Extension {
					res[a_extension_10_] = append(res[a_extension_10_], lightelementtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LinguisticNotePropertyType
	case LinguisticNotePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LinguisticNoteType
	case LinguisticNoteType:
		switch fieldname {
		// insertion point for per direct association field
		case "Extension":
			res := make(map[*A_extension_4][]*LinguisticNoteType)
			for linguisticnotetype := range stage.LinguisticNoteTypes {
				for _, a_extension_4_ := range linguisticnotetype.Extension {
					res[a_extension_4_] = append(res[a_extension_4_], linguisticnotetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MarkerBeaconPropertyType
	case MarkerBeaconPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NavaidEquipmentPropertyType
	case NavaidEquipmentPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NotePropertyType
	case NotePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteType
	case NoteType:
		switch fieldname {
		// insertion point for per direct association field
		case "TranslatedNote":
			res := make(map[*LinguisticNotePropertyType][]*NoteType)
			for notetype := range stage.NoteTypes {
				for _, linguisticnotepropertytype_ := range notetype.TranslatedNote {
					res[linguisticnotepropertytype_] = append(res[linguisticnotepropertytype_], notetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_5][]*NoteType)
			for notetype := range stage.NoteTypes {
				for _, a_extension_5_ := range notetype.Extension {
					res[a_extension_5_] = append(res[a_extension_5_], notetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ObstacleAreaTimeSlicePropertyType
	case ObstacleAreaTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ObstacleAreaTimeSliceType
	case ObstacleAreaTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Obstacle":
			res := make(map[*VerticalStructurePropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				for _, verticalstructurepropertytype_ := range obstacleareatimeslicetype.Obstacle {
					res[verticalstructurepropertytype_] = append(res[verticalstructurepropertytype_], obstacleareatimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				for _, notepropertytype_ := range obstacleareatimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], obstacleareatimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_6][]*ObstacleAreaTimeSliceType)
			for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
				for _, a_extension_6_ := range obstacleareatimeslicetype.Extension {
					res[a_extension_6_] = append(res[a_extension_6_], obstacleareatimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ObstacleAreaType
	case ObstacleAreaType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*ObstacleAreaTimeSlicePropertyType][]*ObstacleAreaType)
			for obstacleareatype := range stage.ObstacleAreaTypes {
				for _, obstacleareatimeslicepropertytype_ := range obstacleareatype.TimeSlice {
					res[obstacleareatimeslicepropertytype_] = append(res[obstacleareatimeslicepropertytype_], obstacleareatype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of OrganisationAuthorityPropertyType
	case OrganisationAuthorityPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PassengerServicePropertyType
	case PassengerServicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointType
	case PointType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*PointType)
			for pointtype := range stage.PointTypes {
				for _, notepropertytype_ := range pointtype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], pointtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RunwayDirectionPropertyType
	case RunwayDirectionPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ServicePropertyType
	case ServicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SpecialNavigationStationPropertyType
	case SpecialNavigationStationPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SurfacePropertyType
	case SurfacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SurfaceType
	case SurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*SurfaceType)
			for surfacetype := range stage.SurfaceTypes {
				for _, notepropertytype_ := range surfacetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], surfacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TimesheetPropertyType
	case TimesheetPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimesheetType
	case TimesheetType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*TimesheetType)
			for timesheettype := range stage.TimesheetTypes {
				for _, notepropertytype_ := range timesheettype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], timesheettype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_12][]*TimesheetType)
			for timesheettype := range stage.TimesheetTypes {
				for _, a_extension_12_ := range timesheettype.Extension {
					res[a_extension_12_] = append(res[a_extension_12_], timesheettype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of UnitPropertyType
	case UnitPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureLightingStatusPropertyType
	case VerticalStructureLightingStatusPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureLightingStatusType
	case VerticalStructureLightingStatusType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInterval":
			res := make(map[*TimesheetPropertyType][]*VerticalStructureLightingStatusType)
			for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
				for _, timesheetpropertytype_ := range verticalstructurelightingstatustype.TimeInterval {
					res[timesheetpropertytype_] = append(res[timesheetpropertytype_], verticalstructurelightingstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*VerticalStructureLightingStatusType)
			for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
				for _, notepropertytype_ := range verticalstructurelightingstatustype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], verticalstructurelightingstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpecialDateAuthority":
			res := make(map[*OrganisationAuthorityPropertyType][]*VerticalStructureLightingStatusType)
			for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
				for _, organisationauthoritypropertytype_ := range verticalstructurelightingstatustype.SpecialDateAuthority {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], verticalstructurelightingstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_8][]*VerticalStructureLightingStatusType)
			for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
				for _, a_extension_8_ := range verticalstructurelightingstatustype.Extension {
					res[a_extension_8_] = append(res[a_extension_8_], verticalstructurelightingstatustype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructurePartPropertyType
	case VerticalStructurePartPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructurePartType
	case VerticalStructurePartType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInterval":
			res := make(map[*TimesheetPropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				for _, timesheetpropertytype_ := range verticalstructureparttype.TimeInterval {
					res[timesheetpropertytype_] = append(res[timesheetpropertytype_], verticalstructureparttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				for _, notepropertytype_ := range verticalstructureparttype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], verticalstructureparttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpecialDateAuthority":
			res := make(map[*OrganisationAuthorityPropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				for _, organisationauthoritypropertytype_ := range verticalstructureparttype.SpecialDateAuthority {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], verticalstructureparttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Lighting":
			res := make(map[*LightElementPropertyType][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				for _, lightelementpropertytype_ := range verticalstructureparttype.Lighting {
					res[lightelementpropertytype_] = append(res[lightelementpropertytype_], verticalstructureparttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_9][]*VerticalStructurePartType)
			for verticalstructureparttype := range stage.VerticalStructurePartTypes {
				for _, a_extension_9_ := range verticalstructureparttype.Extension {
					res[a_extension_9_] = append(res[a_extension_9_], verticalstructureparttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructurePropertyType
	case VerticalStructurePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureTimeSlicePropertyType
	case VerticalStructureTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VerticalStructureTimeSliceType
	case VerticalStructureTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*VerticalStructurePartPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, verticalstructurepartpropertytype_ := range verticalstructuretimeslicetype.Part {
					res[verticalstructurepartpropertytype_] = append(res[verticalstructurepartpropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HostedPassengerService":
			res := make(map[*PassengerServicePropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, passengerservicepropertytype_ := range verticalstructuretimeslicetype.HostedPassengerService {
					res[passengerservicepropertytype_] = append(res[passengerservicepropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SupportedGroundLight":
			res := make(map[*GroundLightSystemPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, groundlightsystempropertytype_ := range verticalstructuretimeslicetype.SupportedGroundLight {
					res[groundlightsystempropertytype_] = append(res[groundlightsystempropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HostedNavaidEquipment":
			res := make(map[*NavaidEquipmentPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, navaidequipmentpropertytype_ := range verticalstructuretimeslicetype.HostedNavaidEquipment {
					res[navaidequipmentpropertytype_] = append(res[navaidequipmentpropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HostedSpecialNavStation":
			res := make(map[*SpecialNavigationStationPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, specialnavigationstationpropertytype_ := range verticalstructuretimeslicetype.HostedSpecialNavStation {
					res[specialnavigationstationpropertytype_] = append(res[specialnavigationstationpropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HostedUnit":
			res := make(map[*UnitPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, unitpropertytype_ := range verticalstructuretimeslicetype.HostedUnit {
					res[unitpropertytype_] = append(res[unitpropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HostedOrganisation":
			res := make(map[*OrganisationAuthorityPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, organisationauthoritypropertytype_ := range verticalstructuretimeslicetype.HostedOrganisation {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SupportedService":
			res := make(map[*ServicePropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, servicepropertytype_ := range verticalstructuretimeslicetype.SupportedService {
					res[servicepropertytype_] = append(res[servicepropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, notepropertytype_ := range verticalstructuretimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "LightingAvailability":
			res := make(map[*VerticalStructureLightingStatusPropertyType][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, verticalstructurelightingstatuspropertytype_ := range verticalstructuretimeslicetype.LightingAvailability {
					res[verticalstructurelightingstatuspropertytype_] = append(res[verticalstructurelightingstatuspropertytype_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_7][]*VerticalStructureTimeSliceType)
			for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
				for _, a_extension_7_ := range verticalstructuretimeslicetype.Extension {
					res[a_extension_7_] = append(res[a_extension_7_], verticalstructuretimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of VerticalStructureType
	case VerticalStructureType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*VerticalStructureTimeSlicePropertyType][]*VerticalStructureType)
			for verticalstructuretype := range stage.VerticalStructureTypes {
				for _, verticalstructuretimeslicepropertytype_ := range verticalstructuretype.TimeSlice {
					res[verticalstructuretimeslicepropertytype_] = append(res[verticalstructuretimeslicepropertytype_], verticalstructuretype)
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
func (*A_extension) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ElevatedCurveType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_1) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ElevatedPointType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_10) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LightElementType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_11) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LightElementStatusType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_12) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "TimesheetType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_2) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ElevatedSurfaceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_3) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AeronauticalGroundLightTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_4) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LinguisticNoteType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_5) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "NoteType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_6) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ObstacleAreaTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_7) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_8) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureLightingStatusType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_9) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructurePartType",
			Fieldname: "Extension",
		},
	}
}

func (*AbstractPropertiesWithScheduleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AeronauticalGroundLightTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AeronauticalGroundLightType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*AeronauticalGroundLightTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AeronauticalGroundLightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirportHeliportPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedCurvePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedCurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedPointPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedPointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedSurfacePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedSurfaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GroundLightSystemPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "SupportedGroundLight",
		},
	}
}

func (*LightElementPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructurePartType",
			Fieldname: "Lighting",
		},
	}
}

func (*LightElementStatusPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LightElementType",
			Fieldname: "Availability",
		},
	}
}

func (*LightElementStatusType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LightElementType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LinguisticNotePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "NoteType",
			Fieldname: "TranslatedNote",
		},
	}
}

func (*LinguisticNoteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MarkerBeaconPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NavaidEquipmentPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "HostedNavaidEquipment",
		},
	}
}

func (*NotePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AeronauticalGroundLightTimeSliceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "CurveType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "ElevatedCurveType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "ElevatedPointType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "ElevatedSurfaceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "LightElementStatusType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "LightElementType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "ObstacleAreaTimeSliceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "PointType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "SurfaceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "TimesheetType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "VerticalStructureLightingStatusType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "VerticalStructurePartType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "Annotation",
		},
	}
}

func (*NoteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ObstacleAreaTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ObstacleAreaType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*ObstacleAreaTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ObstacleAreaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OrganisationAuthorityPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LightElementStatusType",
			Fieldname: "SpecialDateAuthority",
		},
		{
			GongstructName: "VerticalStructureLightingStatusType",
			Fieldname: "SpecialDateAuthority",
		},
		{
			GongstructName: "VerticalStructurePartType",
			Fieldname: "SpecialDateAuthority",
		},
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "HostedOrganisation",
		},
	}
}

func (*PassengerServicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "HostedPassengerService",
		},
	}
}

func (*PointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*RunwayDirectionPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ServicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "SupportedService",
		},
	}
}

func (*SpecialNavigationStationPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "HostedSpecialNavStation",
		},
	}
}

func (*SurfacePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SurfaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimesheetPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LightElementStatusType",
			Fieldname: "TimeInterval",
		},
		{
			GongstructName: "VerticalStructureLightingStatusType",
			Fieldname: "TimeInterval",
		},
		{
			GongstructName: "VerticalStructurePartType",
			Fieldname: "TimeInterval",
		},
	}
}

func (*TimesheetType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*UnitPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "HostedUnit",
		},
	}
}

func (*VerticalStructureLightingStatusPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "LightingAvailability",
		},
	}
}

func (*VerticalStructureLightingStatusType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*VerticalStructurePartPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureTimeSliceType",
			Fieldname: "Part",
		},
	}
}

func (*VerticalStructurePartType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*VerticalStructurePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ObstacleAreaTimeSliceType",
			Fieldname: "Obstacle",
		},
	}
}

func (*VerticalStructureTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "VerticalStructureType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*VerticalStructureTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*VerticalStructureType) GongGetReverseFields() []GongReverseField {
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
func (a_extension *A_extension) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_1 *A_extension_1) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_10 *A_extension_10) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_11 *A_extension_11) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_12 *A_extension_12) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_2 *A_extension_2) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_3 *A_extension_3) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_4 *A_extension_4) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_5 *A_extension_5) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_6 *A_extension_6) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_7 *A_extension_7) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_8 *A_extension_8) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_9 *A_extension_9) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "AeronauticalGroundLightTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AeronauticalGroundLightTimeSliceType",
		},
	}
	return
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Interpretation",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:                 "StructureBeacon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VerticalStructurePropertyType",
		},
		{
			Name:                 "AerodromeBeacon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirportHeliportPropertyType",
		},
		{
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedPointPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_3",
		},
	}
	return
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeSlice",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AeronauticalGroundLightTimeSlicePropertyType",
		},
	}
	return
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
	}
	return
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ElevatedCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedCurveType",
		},
	}
	return
}

func (elevatedcurvetype *ElevatedCurveType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension",
		},
	}
	return
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ElevatedPoint",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedPointType",
		},
	}
	return
}

func (elevatedpointtype *ElevatedPointType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_1",
		},
	}
	return
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ElevatedSurface",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedSurfaceType",
		},
	}
	return
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_2",
		},
	}
	return
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (lightelementpropertytype *LightElementPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "LightElement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LightElementType",
		},
	}
	return
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "LightElementStatus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LightElementStatusType",
		},
	}
	return
}

func (lightelementstatustype *LightElementStatusType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeInterval",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TimesheetPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "SpecialDateAuthority",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_11",
		},
	}
	return
}

func (lightelementtype *LightElementType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedPointPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Availability",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LightElementStatusPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_10",
		},
	}
	return
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "LinguisticNote",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "LinguisticNoteType",
		},
	}
	return
}

func (linguisticnotetype *LinguisticNoteType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_4",
		},
	}
	return
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (notepropertytype *NotePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "NoteType",
		},
	}
	return
}

func (notetype *NoteType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TranslatedNote",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LinguisticNotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_5",
		},
	}
	return
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "ObstacleAreaTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ObstacleAreaTimeSliceType",
		},
	}
	return
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Interpretation",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:                 "Reference_ownerOrganisation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "Reference_ownerRunway",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RunwayDirectionPropertyType",
		},
		{
			Name:                 "Reference_ownerAirport",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirportHeliportPropertyType",
		},
		{
			Name:                 "SurfaceExtent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SurfacePropertyType",
		},
		{
			Name:                 "Obstacle",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "VerticalStructurePropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_6",
		},
	}
	return
}

func (obstacleareatype *ObstacleAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeSlice",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ObstacleAreaTimeSlicePropertyType",
		},
	}
	return
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
	}
	return
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (servicepropertytype *ServicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (surfacepropertytype *SurfacePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Surface",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SurfaceType",
		},
		{
			Name:                 "ElevatedSurface",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedSurfaceType",
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
			Name:               "SrsName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
	}
	return
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Timesheet",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TimesheetType",
		},
	}
	return
}

func (timesheettype *TimesheetType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_12",
		},
	}
	return
}

func (unitpropertytype *UnitPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "VerticalStructureLightingStatus",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VerticalStructureLightingStatusType",
		},
	}
	return
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeInterval",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TimesheetPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "SpecialDateAuthority",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_8",
		},
	}
	return
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "VerticalStructurePart",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VerticalStructurePartType",
		},
	}
	return
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeInterval",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TimesheetPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "SpecialDateAuthority",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "HorizontalProjection_surfaceExtent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedSurfacePropertyType",
		},
		{
			Name:                 "HorizontalProjection_linearExtent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedCurvePropertyType",
		},
		{
			Name:                 "HorizontalProjection_location",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedPointPropertyType",
		},
		{
			Name:                 "Lighting",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LightElementPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_9",
		},
	}
	return
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "VerticalStructureTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VerticalStructureTimeSliceType",
		},
	}
	return
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Interpretation",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:                 "Marker",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MarkerBeaconPropertyType",
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "VerticalStructurePartPropertyType",
		},
		{
			Name:                 "HostedPassengerService",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PassengerServicePropertyType",
		},
		{
			Name:                 "SupportedGroundLight",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroundLightSystemPropertyType",
		},
		{
			Name:                 "HostedNavaidEquipment",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NavaidEquipmentPropertyType",
		},
		{
			Name:                 "HostedSpecialNavStation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SpecialNavigationStationPropertyType",
		},
		{
			Name:                 "HostedUnit",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "UnitPropertyType",
		},
		{
			Name:                 "HostedOrganisation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "SupportedService",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ServicePropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "LightingAvailability",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "VerticalStructureLightingStatusPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_7",
		},
	}
	return
}

func (verticalstructuretype *VerticalStructureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TimeSlice",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "VerticalStructureTimeSlicePropertyType",
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
func (a_extension *A_extension) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension.Owns)
		res.valueBool = a_extension.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_1 *A_extension_1) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_1.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_1.Owns)
		res.valueBool = a_extension_1.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_10 *A_extension_10) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_10.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_10.Owns)
		res.valueBool = a_extension_10.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_11 *A_extension_11) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_11.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_11.Owns)
		res.valueBool = a_extension_11.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_12 *A_extension_12) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_12.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_12.Owns)
		res.valueBool = a_extension_12.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_2 *A_extension_2) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_2.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_2.Owns)
		res.valueBool = a_extension_2.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_3 *A_extension_3) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_3.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_3.Owns)
		res.valueBool = a_extension_3.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_4 *A_extension_4) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_4.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_4.Owns)
		res.valueBool = a_extension_4.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_5 *A_extension_5) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_5.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_5.Owns)
		res.valueBool = a_extension_5.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_6 *A_extension_6) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_6.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_6.Owns)
		res.valueBool = a_extension_6.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_7 *A_extension_7) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_7.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_7.Owns)
		res.valueBool = a_extension_7.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_8 *A_extension_8) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_8.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_8.Owns)
		res.valueBool = a_extension_8.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_9 *A_extension_9) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_9.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_9.Owns)
		res.valueBool = a_extension_9.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = abstractpropertieswithscheduletype.Name
	case "Id":
		res.valueString = abstractpropertieswithscheduletype.Id
	case "Description":
		res.valueString = abstractpropertieswithscheduletype.Description
	}
	return
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aeronauticalgroundlighttimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", aeronauticalgroundlighttimeslicepropertytype.Owns)
		res.valueBool = aeronauticalgroundlighttimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AeronauticalGroundLightTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
			res.valueString = aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice.Name
			res.ids = aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aeronauticalgroundlighttimeslicetype.Name
	case "Id":
		res.valueString = aeronauticalgroundlighttimeslicetype.Id
	case "Description":
		res.valueString = aeronauticalgroundlighttimeslicetype.Description
	case "Interpretation":
		res.valueString = aeronauticalgroundlighttimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", aeronauticalgroundlighttimeslicetype.SequenceNumber)
		res.valueInt = aeronauticalgroundlighttimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", aeronauticalgroundlighttimeslicetype.CorrectionNumber)
		res.valueInt = aeronauticalgroundlighttimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "StructureBeacon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
			res.valueString = aeronauticalgroundlighttimeslicetype.StructureBeacon.Name
			res.ids = aeronauticalgroundlighttimeslicetype.StructureBeacon.GongGetUUID(stage)
		}
	case "AerodromeBeacon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
			res.valueString = aeronauticalgroundlighttimeslicetype.AerodromeBeacon.Name
			res.ids = aeronauticalgroundlighttimeslicetype.AerodromeBeacon.GongGetUUID(stage)
		}
	case "Location":
		res.GongFieldValueType = GongFieldValueTypePointer
		if aeronauticalgroundlighttimeslicetype.Location != nil {
			res.valueString = aeronauticalgroundlighttimeslicetype.Location.Name
			res.ids = aeronauticalgroundlighttimeslicetype.Location.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aeronauticalgroundlighttimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aeronauticalgroundlighttimeslicetype.Extension {
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

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aeronauticalgroundlighttype.Name
	case "Id":
		res.valueString = aeronauticalgroundlighttype.Id
	case "Description":
		res.valueString = aeronauticalgroundlighttype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aeronauticalgroundlighttype.TimeSlice {
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

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airportheliportpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", airportheliportpropertytype.Owns)
		res.valueBool = airportheliportpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = airportheliportpropertytype.NilReason
	case "Type_XSD":
		res.valueString = airportheliportpropertytype.Type_XSD
	case "Href":
		res.valueString = airportheliportpropertytype.Href
	case "Role":
		res.valueString = airportheliportpropertytype.Role
	case "Arcrole":
		res.valueString = airportheliportpropertytype.Arcrole
	case "Title":
		res.valueString = airportheliportpropertytype.Title
	case "Show":
		res.valueString = airportheliportpropertytype.Show
	case "Actuate":
		res.valueString = airportheliportpropertytype.Actuate
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
	case "SrsName":
		res.valueString = curvetype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range curvetype.Annotation {
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

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedcurvepropertytype.Name
	case "NilReason":
		res.valueString = elevatedcurvepropertytype.NilReason
	case "ElevatedCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if elevatedcurvepropertytype.ElevatedCurve != nil {
			res.valueString = elevatedcurvepropertytype.ElevatedCurve.Name
			res.ids = elevatedcurvepropertytype.ElevatedCurve.GongGetUUID(stage)
		}
	}
	return
}

func (elevatedcurvetype *ElevatedCurveType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedcurvetype.Name
	case "Id":
		res.valueString = elevatedcurvetype.Id
	case "Description":
		res.valueString = elevatedcurvetype.Description
	case "SrsName":
		res.valueString = elevatedcurvetype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedcurvetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedcurvetype.Extension {
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

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedpointpropertytype.Name
	case "NilReason":
		res.valueString = elevatedpointpropertytype.NilReason
	case "ElevatedPoint":
		res.GongFieldValueType = GongFieldValueTypePointer
		if elevatedpointpropertytype.ElevatedPoint != nil {
			res.valueString = elevatedpointpropertytype.ElevatedPoint.Name
			res.ids = elevatedpointpropertytype.ElevatedPoint.GongGetUUID(stage)
		}
	}
	return
}

func (elevatedpointtype *ElevatedPointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedpointtype.Name
	case "Id":
		res.valueString = elevatedpointtype.Id
	case "Description":
		res.valueString = elevatedpointtype.Description
	case "SrsName":
		res.valueString = elevatedpointtype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedpointtype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedpointtype.Extension {
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

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedsurfacepropertytype.Name
	case "NilReason":
		res.valueString = elevatedsurfacepropertytype.NilReason
	case "ElevatedSurface":
		res.GongFieldValueType = GongFieldValueTypePointer
		if elevatedsurfacepropertytype.ElevatedSurface != nil {
			res.valueString = elevatedsurfacepropertytype.ElevatedSurface.Name
			res.ids = elevatedsurfacepropertytype.ElevatedSurface.GongGetUUID(stage)
		}
	}
	return
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elevatedsurfacetype.Name
	case "Id":
		res.valueString = elevatedsurfacetype.Id
	case "Description":
		res.valueString = elevatedsurfacetype.Description
	case "SrsName":
		res.valueString = elevatedsurfacetype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedsurfacetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range elevatedsurfacetype.Extension {
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

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = groundlightsystempropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", groundlightsystempropertytype.Owns)
		res.valueBool = groundlightsystempropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = groundlightsystempropertytype.NilReason
	case "Type_XSD":
		res.valueString = groundlightsystempropertytype.Type_XSD
	case "Href":
		res.valueString = groundlightsystempropertytype.Href
	case "Role":
		res.valueString = groundlightsystempropertytype.Role
	case "Arcrole":
		res.valueString = groundlightsystempropertytype.Arcrole
	case "Title":
		res.valueString = groundlightsystempropertytype.Title
	case "Show":
		res.valueString = groundlightsystempropertytype.Show
	case "Actuate":
		res.valueString = groundlightsystempropertytype.Actuate
	}
	return
}

func (lightelementpropertytype *LightElementPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lightelementpropertytype.Name
	case "NilReason":
		res.valueString = lightelementpropertytype.NilReason
	case "LightElement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if lightelementpropertytype.LightElement != nil {
			res.valueString = lightelementpropertytype.LightElement.Name
			res.ids = lightelementpropertytype.LightElement.GongGetUUID(stage)
		}
	}
	return
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lightelementstatuspropertytype.Name
	case "NilReason":
		res.valueString = lightelementstatuspropertytype.NilReason
	case "LightElementStatus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if lightelementstatuspropertytype.LightElementStatus != nil {
			res.valueString = lightelementstatuspropertytype.LightElementStatus.Name
			res.ids = lightelementstatuspropertytype.LightElementStatus.GongGetUUID(stage)
		}
	}
	return
}

func (lightelementstatustype *LightElementStatusType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lightelementstatustype.Name
	case "Id":
		res.valueString = lightelementstatustype.Id
	case "Description":
		res.valueString = lightelementstatustype.Description
	case "TimeInterval":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementstatustype.TimeInterval {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementstatustype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpecialDateAuthority":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementstatustype.SpecialDateAuthority {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementstatustype.Extension {
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

func (lightelementtype *LightElementType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lightelementtype.Name
	case "Id":
		res.valueString = lightelementtype.Id
	case "Description":
		res.valueString = lightelementtype.Description
	case "Location":
		res.GongFieldValueType = GongFieldValueTypePointer
		if lightelementtype.Location != nil {
			res.valueString = lightelementtype.Location.Name
			res.ids = lightelementtype.Location.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementtype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Availability":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementtype.Availability {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lightelementtype.Extension {
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

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linguisticnotepropertytype.Name
	case "NilReason":
		res.valueString = linguisticnotepropertytype.NilReason
	case "LinguisticNote":
		res.GongFieldValueType = GongFieldValueTypePointer
		if linguisticnotepropertytype.LinguisticNote != nil {
			res.valueString = linguisticnotepropertytype.LinguisticNote.Name
			res.ids = linguisticnotepropertytype.LinguisticNote.GongGetUUID(stage)
		}
	}
	return
}

func (linguisticnotetype *LinguisticNoteType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linguisticnotetype.Name
	case "Id":
		res.valueString = linguisticnotetype.Id
	case "Description":
		res.valueString = linguisticnotetype.Description
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range linguisticnotetype.Extension {
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

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = markerbeaconpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", markerbeaconpropertytype.Owns)
		res.valueBool = markerbeaconpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = markerbeaconpropertytype.NilReason
	case "Type_XSD":
		res.valueString = markerbeaconpropertytype.Type_XSD
	case "Href":
		res.valueString = markerbeaconpropertytype.Href
	case "Role":
		res.valueString = markerbeaconpropertytype.Role
	case "Arcrole":
		res.valueString = markerbeaconpropertytype.Arcrole
	case "Title":
		res.valueString = markerbeaconpropertytype.Title
	case "Show":
		res.valueString = markerbeaconpropertytype.Show
	case "Actuate":
		res.valueString = markerbeaconpropertytype.Actuate
	}
	return
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = navaidequipmentpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", navaidequipmentpropertytype.Owns)
		res.valueBool = navaidequipmentpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = navaidequipmentpropertytype.NilReason
	case "Type_XSD":
		res.valueString = navaidequipmentpropertytype.Type_XSD
	case "Href":
		res.valueString = navaidequipmentpropertytype.Href
	case "Role":
		res.valueString = navaidequipmentpropertytype.Role
	case "Arcrole":
		res.valueString = navaidequipmentpropertytype.Arcrole
	case "Title":
		res.valueString = navaidequipmentpropertytype.Title
	case "Show":
		res.valueString = navaidequipmentpropertytype.Show
	case "Actuate":
		res.valueString = navaidequipmentpropertytype.Actuate
	}
	return
}

func (notepropertytype *NotePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notepropertytype.Name
	case "NilReason":
		res.valueString = notepropertytype.NilReason
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notepropertytype.Note != nil {
			res.valueString = notepropertytype.Note.Name
			res.ids = notepropertytype.Note.GongGetUUID(stage)
		}
	}
	return
}

func (notetype *NoteType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notetype.Name
	case "Id":
		res.valueString = notetype.Id
	case "Description":
		res.valueString = notetype.Description
	case "TranslatedNote":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notetype.TranslatedNote {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notetype.Extension {
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

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = obstacleareatimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", obstacleareatimeslicepropertytype.Owns)
		res.valueBool = obstacleareatimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ObstacleAreaTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
			res.valueString = obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice.Name
			res.ids = obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = obstacleareatimeslicetype.Name
	case "Id":
		res.valueString = obstacleareatimeslicetype.Id
	case "Description":
		res.valueString = obstacleareatimeslicetype.Description
	case "Interpretation":
		res.valueString = obstacleareatimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", obstacleareatimeslicetype.SequenceNumber)
		res.valueInt = obstacleareatimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", obstacleareatimeslicetype.CorrectionNumber)
		res.valueInt = obstacleareatimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Reference_ownerOrganisation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
			res.valueString = obstacleareatimeslicetype.Reference_ownerOrganisation.Name
			res.ids = obstacleareatimeslicetype.Reference_ownerOrganisation.GongGetUUID(stage)
		}
	case "Reference_ownerRunway":
		res.GongFieldValueType = GongFieldValueTypePointer
		if obstacleareatimeslicetype.Reference_ownerRunway != nil {
			res.valueString = obstacleareatimeslicetype.Reference_ownerRunway.Name
			res.ids = obstacleareatimeslicetype.Reference_ownerRunway.GongGetUUID(stage)
		}
	case "Reference_ownerAirport":
		res.GongFieldValueType = GongFieldValueTypePointer
		if obstacleareatimeslicetype.Reference_ownerAirport != nil {
			res.valueString = obstacleareatimeslicetype.Reference_ownerAirport.Name
			res.ids = obstacleareatimeslicetype.Reference_ownerAirport.GongGetUUID(stage)
		}
	case "SurfaceExtent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if obstacleareatimeslicetype.SurfaceExtent != nil {
			res.valueString = obstacleareatimeslicetype.SurfaceExtent.Name
			res.ids = obstacleareatimeslicetype.SurfaceExtent.GongGetUUID(stage)
		}
	case "Obstacle":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range obstacleareatimeslicetype.Obstacle {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range obstacleareatimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range obstacleareatimeslicetype.Extension {
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

func (obstacleareatype *ObstacleAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = obstacleareatype.Name
	case "Id":
		res.valueString = obstacleareatype.Id
	case "Description":
		res.valueString = obstacleareatype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range obstacleareatype.TimeSlice {
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

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = organisationauthoritypropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", organisationauthoritypropertytype.Owns)
		res.valueBool = organisationauthoritypropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = organisationauthoritypropertytype.NilReason
	case "Type_XSD":
		res.valueString = organisationauthoritypropertytype.Type_XSD
	case "Href":
		res.valueString = organisationauthoritypropertytype.Href
	case "Role":
		res.valueString = organisationauthoritypropertytype.Role
	case "Arcrole":
		res.valueString = organisationauthoritypropertytype.Arcrole
	case "Title":
		res.valueString = organisationauthoritypropertytype.Title
	case "Show":
		res.valueString = organisationauthoritypropertytype.Show
	case "Actuate":
		res.valueString = organisationauthoritypropertytype.Actuate
	}
	return
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = passengerservicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", passengerservicepropertytype.Owns)
		res.valueBool = passengerservicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = passengerservicepropertytype.NilReason
	case "Type_XSD":
		res.valueString = passengerservicepropertytype.Type_XSD
	case "Href":
		res.valueString = passengerservicepropertytype.Href
	case "Role":
		res.valueString = passengerservicepropertytype.Role
	case "Arcrole":
		res.valueString = passengerservicepropertytype.Arcrole
	case "Title":
		res.valueString = passengerservicepropertytype.Title
	case "Show":
		res.valueString = passengerservicepropertytype.Show
	case "Actuate":
		res.valueString = passengerservicepropertytype.Actuate
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
	case "SrsName":
		res.valueString = pointtype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range pointtype.Annotation {
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

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = runwaydirectionpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", runwaydirectionpropertytype.Owns)
		res.valueBool = runwaydirectionpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = runwaydirectionpropertytype.NilReason
	case "Type_XSD":
		res.valueString = runwaydirectionpropertytype.Type_XSD
	case "Href":
		res.valueString = runwaydirectionpropertytype.Href
	case "Role":
		res.valueString = runwaydirectionpropertytype.Role
	case "Arcrole":
		res.valueString = runwaydirectionpropertytype.Arcrole
	case "Title":
		res.valueString = runwaydirectionpropertytype.Title
	case "Show":
		res.valueString = runwaydirectionpropertytype.Show
	case "Actuate":
		res.valueString = runwaydirectionpropertytype.Actuate
	}
	return
}

func (servicepropertytype *ServicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = servicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", servicepropertytype.Owns)
		res.valueBool = servicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = servicepropertytype.NilReason
	case "Type_XSD":
		res.valueString = servicepropertytype.Type_XSD
	case "Href":
		res.valueString = servicepropertytype.Href
	case "Role":
		res.valueString = servicepropertytype.Role
	case "Arcrole":
		res.valueString = servicepropertytype.Arcrole
	case "Title":
		res.valueString = servicepropertytype.Title
	case "Show":
		res.valueString = servicepropertytype.Show
	case "Actuate":
		res.valueString = servicepropertytype.Actuate
	}
	return
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = specialnavigationstationpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", specialnavigationstationpropertytype.Owns)
		res.valueBool = specialnavigationstationpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = specialnavigationstationpropertytype.NilReason
	case "Type_XSD":
		res.valueString = specialnavigationstationpropertytype.Type_XSD
	case "Href":
		res.valueString = specialnavigationstationpropertytype.Href
	case "Role":
		res.valueString = specialnavigationstationpropertytype.Role
	case "Arcrole":
		res.valueString = specialnavigationstationpropertytype.Arcrole
	case "Title":
		res.valueString = specialnavigationstationpropertytype.Title
	case "Show":
		res.valueString = specialnavigationstationpropertytype.Show
	case "Actuate":
		res.valueString = specialnavigationstationpropertytype.Actuate
	}
	return
}

func (surfacepropertytype *SurfacePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = surfacepropertytype.Name
	case "NilReason":
		res.valueString = surfacepropertytype.NilReason
	case "Surface":
		res.GongFieldValueType = GongFieldValueTypePointer
		if surfacepropertytype.Surface != nil {
			res.valueString = surfacepropertytype.Surface.Name
			res.ids = surfacepropertytype.Surface.GongGetUUID(stage)
		}
	case "ElevatedSurface":
		res.GongFieldValueType = GongFieldValueTypePointer
		if surfacepropertytype.ElevatedSurface != nil {
			res.valueString = surfacepropertytype.ElevatedSurface.Name
			res.ids = surfacepropertytype.ElevatedSurface.GongGetUUID(stage)
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
	case "SrsName":
		res.valueString = surfacetype.SrsName
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range surfacetype.Annotation {
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

func (timesheetpropertytype *TimesheetPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timesheetpropertytype.Name
	case "NilReason":
		res.valueString = timesheetpropertytype.NilReason
	case "Timesheet":
		res.GongFieldValueType = GongFieldValueTypePointer
		if timesheetpropertytype.Timesheet != nil {
			res.valueString = timesheetpropertytype.Timesheet.Name
			res.ids = timesheetpropertytype.Timesheet.GongGetUUID(stage)
		}
	}
	return
}

func (timesheettype *TimesheetType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timesheettype.Name
	case "Id":
		res.valueString = timesheettype.Id
	case "Description":
		res.valueString = timesheettype.Description
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range timesheettype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range timesheettype.Extension {
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

func (unitpropertytype *UnitPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = unitpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", unitpropertytype.Owns)
		res.valueBool = unitpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = unitpropertytype.NilReason
	case "Type_XSD":
		res.valueString = unitpropertytype.Type_XSD
	case "Href":
		res.valueString = unitpropertytype.Href
	case "Role":
		res.valueString = unitpropertytype.Role
	case "Arcrole":
		res.valueString = unitpropertytype.Arcrole
	case "Title":
		res.valueString = unitpropertytype.Title
	case "Show":
		res.valueString = unitpropertytype.Show
	case "Actuate":
		res.valueString = unitpropertytype.Actuate
	}
	return
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructurelightingstatuspropertytype.Name
	case "NilReason":
		res.valueString = verticalstructurelightingstatuspropertytype.NilReason
	case "VerticalStructureLightingStatus":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
			res.valueString = verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus.Name
			res.ids = verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus.GongGetUUID(stage)
		}
	}
	return
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructurelightingstatustype.Name
	case "Id":
		res.valueString = verticalstructurelightingstatustype.Id
	case "Description":
		res.valueString = verticalstructurelightingstatustype.Description
	case "TimeInterval":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructurelightingstatustype.TimeInterval {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructurelightingstatustype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpecialDateAuthority":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructurelightingstatustype.SpecialDateAuthority {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructurelightingstatustype.Extension {
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

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructurepartpropertytype.Name
	case "NilReason":
		res.valueString = verticalstructurepartpropertytype.NilReason
	case "VerticalStructurePart":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructurepartpropertytype.VerticalStructurePart != nil {
			res.valueString = verticalstructurepartpropertytype.VerticalStructurePart.Name
			res.ids = verticalstructurepartpropertytype.VerticalStructurePart.GongGetUUID(stage)
		}
	}
	return
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructureparttype.Name
	case "Id":
		res.valueString = verticalstructureparttype.Id
	case "Description":
		res.valueString = verticalstructureparttype.Description
	case "TimeInterval":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructureparttype.TimeInterval {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructureparttype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpecialDateAuthority":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructureparttype.SpecialDateAuthority {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HorizontalProjection_surfaceExtent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
			res.valueString = verticalstructureparttype.HorizontalProjection_surfaceExtent.Name
			res.ids = verticalstructureparttype.HorizontalProjection_surfaceExtent.GongGetUUID(stage)
		}
	case "HorizontalProjection_linearExtent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
			res.valueString = verticalstructureparttype.HorizontalProjection_linearExtent.Name
			res.ids = verticalstructureparttype.HorizontalProjection_linearExtent.GongGetUUID(stage)
		}
	case "HorizontalProjection_location":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructureparttype.HorizontalProjection_location != nil {
			res.valueString = verticalstructureparttype.HorizontalProjection_location.Name
			res.ids = verticalstructureparttype.HorizontalProjection_location.GongGetUUID(stage)
		}
	case "Lighting":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructureparttype.Lighting {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructureparttype.Extension {
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

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructurepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", verticalstructurepropertytype.Owns)
		res.valueBool = verticalstructurepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = verticalstructurepropertytype.NilReason
	case "Type_XSD":
		res.valueString = verticalstructurepropertytype.Type_XSD
	case "Href":
		res.valueString = verticalstructurepropertytype.Href
	case "Role":
		res.valueString = verticalstructurepropertytype.Role
	case "Arcrole":
		res.valueString = verticalstructurepropertytype.Arcrole
	case "Title":
		res.valueString = verticalstructurepropertytype.Title
	case "Show":
		res.valueString = verticalstructurepropertytype.Show
	case "Actuate":
		res.valueString = verticalstructurepropertytype.Actuate
	}
	return
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructuretimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", verticalstructuretimeslicepropertytype.Owns)
		res.valueBool = verticalstructuretimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "VerticalStructureTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
			res.valueString = verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice.Name
			res.ids = verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructuretimeslicetype.Name
	case "Id":
		res.valueString = verticalstructuretimeslicetype.Id
	case "Description":
		res.valueString = verticalstructuretimeslicetype.Description
	case "Interpretation":
		res.valueString = verticalstructuretimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", verticalstructuretimeslicetype.SequenceNumber)
		res.valueInt = verticalstructuretimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", verticalstructuretimeslicetype.CorrectionNumber)
		res.valueInt = verticalstructuretimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Marker":
		res.GongFieldValueType = GongFieldValueTypePointer
		if verticalstructuretimeslicetype.Marker != nil {
			res.valueString = verticalstructuretimeslicetype.Marker.Name
			res.ids = verticalstructuretimeslicetype.Marker.GongGetUUID(stage)
		}
	case "Part":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.Part {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HostedPassengerService":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.HostedPassengerService {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SupportedGroundLight":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.SupportedGroundLight {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HostedNavaidEquipment":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.HostedNavaidEquipment {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HostedSpecialNavStation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.HostedSpecialNavStation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HostedUnit":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.HostedUnit {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HostedOrganisation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.HostedOrganisation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SupportedService":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.SupportedService {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "LightingAvailability":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.LightingAvailability {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretimeslicetype.Extension {
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

func (verticalstructuretype *VerticalStructureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = verticalstructuretype.Name
	case "Id":
		res.valueString = verticalstructuretype.Id
	case "Description":
		res.valueString = verticalstructuretype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range verticalstructuretype.TimeSlice {
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
func (a_extension *A_extension) GongGetGongstructName() string {
	return "A_extension"
}

func (a_extension_1 *A_extension_1) GongGetGongstructName() string {
	return "A_extension_1"
}

func (a_extension_10 *A_extension_10) GongGetGongstructName() string {
	return "A_extension_10"
}

func (a_extension_11 *A_extension_11) GongGetGongstructName() string {
	return "A_extension_11"
}

func (a_extension_12 *A_extension_12) GongGetGongstructName() string {
	return "A_extension_12"
}

func (a_extension_2 *A_extension_2) GongGetGongstructName() string {
	return "A_extension_2"
}

func (a_extension_3 *A_extension_3) GongGetGongstructName() string {
	return "A_extension_3"
}

func (a_extension_4 *A_extension_4) GongGetGongstructName() string {
	return "A_extension_4"
}

func (a_extension_5 *A_extension_5) GongGetGongstructName() string {
	return "A_extension_5"
}

func (a_extension_6 *A_extension_6) GongGetGongstructName() string {
	return "A_extension_6"
}

func (a_extension_7 *A_extension_7) GongGetGongstructName() string {
	return "A_extension_7"
}

func (a_extension_8 *A_extension_8) GongGetGongstructName() string {
	return "A_extension_8"
}

func (a_extension_9 *A_extension_9) GongGetGongstructName() string {
	return "A_extension_9"
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetGongstructName() string {
	return "AbstractPropertiesWithScheduleType"
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetGongstructName() string {
	return "AeronauticalGroundLightTimeSlicePropertyType"
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetGongstructName() string {
	return "AeronauticalGroundLightTimeSliceType"
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetGongstructName() string {
	return "AeronauticalGroundLightType"
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetGongstructName() string {
	return "AirportHeliportPropertyType"
}

func (curvetype *CurveType) GongGetGongstructName() string {
	return "CurveType"
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetGongstructName() string {
	return "ElevatedCurvePropertyType"
}

func (elevatedcurvetype *ElevatedCurveType) GongGetGongstructName() string {
	return "ElevatedCurveType"
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetGongstructName() string {
	return "ElevatedPointPropertyType"
}

func (elevatedpointtype *ElevatedPointType) GongGetGongstructName() string {
	return "ElevatedPointType"
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetGongstructName() string {
	return "ElevatedSurfacePropertyType"
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetGongstructName() string {
	return "ElevatedSurfaceType"
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetGongstructName() string {
	return "GroundLightSystemPropertyType"
}

func (lightelementpropertytype *LightElementPropertyType) GongGetGongstructName() string {
	return "LightElementPropertyType"
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetGongstructName() string {
	return "LightElementStatusPropertyType"
}

func (lightelementstatustype *LightElementStatusType) GongGetGongstructName() string {
	return "LightElementStatusType"
}

func (lightelementtype *LightElementType) GongGetGongstructName() string {
	return "LightElementType"
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetGongstructName() string {
	return "LinguisticNotePropertyType"
}

func (linguisticnotetype *LinguisticNoteType) GongGetGongstructName() string {
	return "LinguisticNoteType"
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetGongstructName() string {
	return "MarkerBeaconPropertyType"
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetGongstructName() string {
	return "NavaidEquipmentPropertyType"
}

func (notepropertytype *NotePropertyType) GongGetGongstructName() string {
	return "NotePropertyType"
}

func (notetype *NoteType) GongGetGongstructName() string {
	return "NoteType"
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetGongstructName() string {
	return "ObstacleAreaTimeSlicePropertyType"
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetGongstructName() string {
	return "ObstacleAreaTimeSliceType"
}

func (obstacleareatype *ObstacleAreaType) GongGetGongstructName() string {
	return "ObstacleAreaType"
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetGongstructName() string {
	return "OrganisationAuthorityPropertyType"
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetGongstructName() string {
	return "PassengerServicePropertyType"
}

func (pointtype *PointType) GongGetGongstructName() string {
	return "PointType"
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetGongstructName() string {
	return "RunwayDirectionPropertyType"
}

func (servicepropertytype *ServicePropertyType) GongGetGongstructName() string {
	return "ServicePropertyType"
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetGongstructName() string {
	return "SpecialNavigationStationPropertyType"
}

func (surfacepropertytype *SurfacePropertyType) GongGetGongstructName() string {
	return "SurfacePropertyType"
}

func (surfacetype *SurfaceType) GongGetGongstructName() string {
	return "SurfaceType"
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetGongstructName() string {
	return "TimesheetPropertyType"
}

func (timesheettype *TimesheetType) GongGetGongstructName() string {
	return "TimesheetType"
}

func (unitpropertytype *UnitPropertyType) GongGetGongstructName() string {
	return "UnitPropertyType"
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetGongstructName() string {
	return "VerticalStructureLightingStatusPropertyType"
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetGongstructName() string {
	return "VerticalStructureLightingStatusType"
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetGongstructName() string {
	return "VerticalStructurePartPropertyType"
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetGongstructName() string {
	return "VerticalStructurePartType"
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetGongstructName() string {
	return "VerticalStructurePropertyType"
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetGongstructName() string {
	return "VerticalStructureTimeSlicePropertyType"
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetGongstructName() string {
	return "VerticalStructureTimeSliceType"
}

func (verticalstructuretype *VerticalStructureType) GongGetGongstructName() string {
	return "VerticalStructureType"
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
	__gong__rebuildMapString(stage.A_extensions, &stage.A_extensions_mapString)

	__gong__rebuildMapString(stage.A_extension_1s, &stage.A_extension_1s_mapString)

	__gong__rebuildMapString(stage.A_extension_10s, &stage.A_extension_10s_mapString)

	__gong__rebuildMapString(stage.A_extension_11s, &stage.A_extension_11s_mapString)

	__gong__rebuildMapString(stage.A_extension_12s, &stage.A_extension_12s_mapString)

	__gong__rebuildMapString(stage.A_extension_2s, &stage.A_extension_2s_mapString)

	__gong__rebuildMapString(stage.A_extension_3s, &stage.A_extension_3s_mapString)

	__gong__rebuildMapString(stage.A_extension_4s, &stage.A_extension_4s_mapString)

	__gong__rebuildMapString(stage.A_extension_5s, &stage.A_extension_5s_mapString)

	__gong__rebuildMapString(stage.A_extension_6s, &stage.A_extension_6s_mapString)

	__gong__rebuildMapString(stage.A_extension_7s, &stage.A_extension_7s_mapString)

	__gong__rebuildMapString(stage.A_extension_8s, &stage.A_extension_8s_mapString)

	__gong__rebuildMapString(stage.A_extension_9s, &stage.A_extension_9s_mapString)

	__gong__rebuildMapString(stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_mapString)

	__gong__rebuildMapString(stage.AeronauticalGroundLightTimeSlicePropertyTypes, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.AeronauticalGroundLightTimeSliceTypes, &stage.AeronauticalGroundLightTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.AeronauticalGroundLightTypes, &stage.AeronauticalGroundLightTypes_mapString)

	__gong__rebuildMapString(stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_mapString)

	__gong__rebuildMapString(stage.CurveTypes, &stage.CurveTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedCurvePropertyTypes, &stage.ElevatedCurvePropertyTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedPointPropertyTypes, &stage.ElevatedPointPropertyTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedPointTypes, &stage.ElevatedPointTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedSurfacePropertyTypes, &stage.ElevatedSurfacePropertyTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.GroundLightSystemPropertyTypes, &stage.GroundLightSystemPropertyTypes_mapString)

	__gong__rebuildMapString(stage.LightElementPropertyTypes, &stage.LightElementPropertyTypes_mapString)

	__gong__rebuildMapString(stage.LightElementStatusPropertyTypes, &stage.LightElementStatusPropertyTypes_mapString)

	__gong__rebuildMapString(stage.LightElementStatusTypes, &stage.LightElementStatusTypes_mapString)

	__gong__rebuildMapString(stage.LightElementTypes, &stage.LightElementTypes_mapString)

	__gong__rebuildMapString(stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_mapString)

	__gong__rebuildMapString(stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_mapString)

	__gong__rebuildMapString(stage.MarkerBeaconPropertyTypes, &stage.MarkerBeaconPropertyTypes_mapString)

	__gong__rebuildMapString(stage.NavaidEquipmentPropertyTypes, &stage.NavaidEquipmentPropertyTypes_mapString)

	__gong__rebuildMapString(stage.NotePropertyTypes, &stage.NotePropertyTypes_mapString)

	__gong__rebuildMapString(stage.NoteTypes, &stage.NoteTypes_mapString)

	__gong__rebuildMapString(stage.ObstacleAreaTimeSlicePropertyTypes, &stage.ObstacleAreaTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.ObstacleAreaTimeSliceTypes, &stage.ObstacleAreaTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.ObstacleAreaTypes, &stage.ObstacleAreaTypes_mapString)

	__gong__rebuildMapString(stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_mapString)

	__gong__rebuildMapString(stage.PassengerServicePropertyTypes, &stage.PassengerServicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.PointTypes, &stage.PointTypes_mapString)

	__gong__rebuildMapString(stage.RunwayDirectionPropertyTypes, &stage.RunwayDirectionPropertyTypes_mapString)

	__gong__rebuildMapString(stage.ServicePropertyTypes, &stage.ServicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.SpecialNavigationStationPropertyTypes, &stage.SpecialNavigationStationPropertyTypes_mapString)

	__gong__rebuildMapString(stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_mapString)

	__gong__rebuildMapString(stage.SurfaceTypes, &stage.SurfaceTypes_mapString)

	__gong__rebuildMapString(stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_mapString)

	__gong__rebuildMapString(stage.TimesheetTypes, &stage.TimesheetTypes_mapString)

	__gong__rebuildMapString(stage.UnitPropertyTypes, &stage.UnitPropertyTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructureLightingStatusPropertyTypes, &stage.VerticalStructureLightingStatusPropertyTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructureLightingStatusTypes, &stage.VerticalStructureLightingStatusTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructurePartPropertyTypes, &stage.VerticalStructurePartPropertyTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructurePartTypes, &stage.VerticalStructurePartTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructurePropertyTypes, &stage.VerticalStructurePropertyTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructureTimeSlicePropertyTypes, &stage.VerticalStructureTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructureTimeSliceTypes, &stage.VerticalStructureTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.VerticalStructureTypes, &stage.VerticalStructureTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
