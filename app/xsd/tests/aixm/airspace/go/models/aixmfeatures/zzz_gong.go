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

	A_extension_13s                map[*A_extension_13]struct{}
	A_extension_13s_instance       map[*A_extension_13]*A_extension_13
	A_extension_13s_mapString      map[string]*A_extension_13
	A_extension_13Order            uint
	A_extension_13_stagedOrder     map[*A_extension_13]uint
	A_extension_13_orderStaged     map[uint]*A_extension_13
	A_extension_13s_reference      map[*A_extension_13]*A_extension_13
	A_extension_13s_referenceOrder map[*A_extension_13]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_13CreateCallback GongOnAfterCreateInterface[A_extension_13]
	OnAfterA_extension_13UpdateCallback GongOnAfterUpdateInterface[A_extension_13]
	OnAfterA_extension_13DeleteCallback GongOnAfterDeleteInterface[A_extension_13]

	A_extension_14s                map[*A_extension_14]struct{}
	A_extension_14s_instance       map[*A_extension_14]*A_extension_14
	A_extension_14s_mapString      map[string]*A_extension_14
	A_extension_14Order            uint
	A_extension_14_stagedOrder     map[*A_extension_14]uint
	A_extension_14_orderStaged     map[uint]*A_extension_14
	A_extension_14s_reference      map[*A_extension_14]*A_extension_14
	A_extension_14s_referenceOrder map[*A_extension_14]uint

	// insertion point for slice of pointers maps
	OnAfterA_extension_14CreateCallback GongOnAfterCreateInterface[A_extension_14]
	OnAfterA_extension_14UpdateCallback GongOnAfterUpdateInterface[A_extension_14]
	OnAfterA_extension_14DeleteCallback GongOnAfterDeleteInterface[A_extension_14]

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

	AircraftCharacteristicPropertyTypes                map[*AircraftCharacteristicPropertyType]struct{}
	AircraftCharacteristicPropertyTypes_instance       map[*AircraftCharacteristicPropertyType]*AircraftCharacteristicPropertyType
	AircraftCharacteristicPropertyTypes_mapString      map[string]*AircraftCharacteristicPropertyType
	AircraftCharacteristicPropertyTypeOrder            uint
	AircraftCharacteristicPropertyType_stagedOrder     map[*AircraftCharacteristicPropertyType]uint
	AircraftCharacteristicPropertyType_orderStaged     map[uint]*AircraftCharacteristicPropertyType
	AircraftCharacteristicPropertyTypes_reference      map[*AircraftCharacteristicPropertyType]*AircraftCharacteristicPropertyType
	AircraftCharacteristicPropertyTypes_referenceOrder map[*AircraftCharacteristicPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAircraftCharacteristicPropertyTypeCreateCallback GongOnAfterCreateInterface[AircraftCharacteristicPropertyType]
	OnAfterAircraftCharacteristicPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AircraftCharacteristicPropertyType]
	OnAfterAircraftCharacteristicPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AircraftCharacteristicPropertyType]

	AircraftCharacteristicTypes                map[*AircraftCharacteristicType]struct{}
	AircraftCharacteristicTypes_instance       map[*AircraftCharacteristicType]*AircraftCharacteristicType
	AircraftCharacteristicTypes_mapString      map[string]*AircraftCharacteristicType
	AircraftCharacteristicTypeOrder            uint
	AircraftCharacteristicType_stagedOrder     map[*AircraftCharacteristicType]uint
	AircraftCharacteristicType_orderStaged     map[uint]*AircraftCharacteristicType
	AircraftCharacteristicTypes_reference      map[*AircraftCharacteristicType]*AircraftCharacteristicType
	AircraftCharacteristicTypes_referenceOrder map[*AircraftCharacteristicType]uint

	// insertion point for slice of pointers maps
	AircraftCharacteristicType_Annotation_reverseMap map[*NotePropertyType]*AircraftCharacteristicType

	AircraftCharacteristicType_Extension_reverseMap map[*A_extension_12]*AircraftCharacteristicType

	OnAfterAircraftCharacteristicTypeCreateCallback GongOnAfterCreateInterface[AircraftCharacteristicType]
	OnAfterAircraftCharacteristicTypeUpdateCallback GongOnAfterUpdateInterface[AircraftCharacteristicType]
	OnAfterAircraftCharacteristicTypeDeleteCallback GongOnAfterDeleteInterface[AircraftCharacteristicType]

	AirspaceActivationPropertyTypes                map[*AirspaceActivationPropertyType]struct{}
	AirspaceActivationPropertyTypes_instance       map[*AirspaceActivationPropertyType]*AirspaceActivationPropertyType
	AirspaceActivationPropertyTypes_mapString      map[string]*AirspaceActivationPropertyType
	AirspaceActivationPropertyTypeOrder            uint
	AirspaceActivationPropertyType_stagedOrder     map[*AirspaceActivationPropertyType]uint
	AirspaceActivationPropertyType_orderStaged     map[uint]*AirspaceActivationPropertyType
	AirspaceActivationPropertyTypes_reference      map[*AirspaceActivationPropertyType]*AirspaceActivationPropertyType
	AirspaceActivationPropertyTypes_referenceOrder map[*AirspaceActivationPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceActivationPropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceActivationPropertyType]
	OnAfterAirspaceActivationPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceActivationPropertyType]
	OnAfterAirspaceActivationPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceActivationPropertyType]

	AirspaceActivationTypes                map[*AirspaceActivationType]struct{}
	AirspaceActivationTypes_instance       map[*AirspaceActivationType]*AirspaceActivationType
	AirspaceActivationTypes_mapString      map[string]*AirspaceActivationType
	AirspaceActivationTypeOrder            uint
	AirspaceActivationType_stagedOrder     map[*AirspaceActivationType]uint
	AirspaceActivationType_orderStaged     map[uint]*AirspaceActivationType
	AirspaceActivationTypes_reference      map[*AirspaceActivationType]*AirspaceActivationType
	AirspaceActivationTypes_referenceOrder map[*AirspaceActivationType]uint

	// insertion point for slice of pointers maps
	AirspaceActivationType_TimeInterval_reverseMap map[*TimesheetPropertyType]*AirspaceActivationType

	AirspaceActivationType_Annotation_reverseMap map[*NotePropertyType]*AirspaceActivationType

	AirspaceActivationType_SpecialDateAuthority_reverseMap map[*OrganisationAuthorityPropertyType]*AirspaceActivationType

	AirspaceActivationType_Levels_reverseMap map[*AirspaceLayerPropertyType]*AirspaceActivationType

	AirspaceActivationType_User_reverseMap map[*OrganisationAuthorityPropertyType]*AirspaceActivationType

	AirspaceActivationType_Aircraft_reverseMap map[*AircraftCharacteristicPropertyType]*AirspaceActivationType

	AirspaceActivationType_Extension_reverseMap map[*A_extension_1]*AirspaceActivationType

	OnAfterAirspaceActivationTypeCreateCallback GongOnAfterCreateInterface[AirspaceActivationType]
	OnAfterAirspaceActivationTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceActivationType]
	OnAfterAirspaceActivationTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceActivationType]

	AirspaceGeometryComponentPropertyTypes                map[*AirspaceGeometryComponentPropertyType]struct{}
	AirspaceGeometryComponentPropertyTypes_instance       map[*AirspaceGeometryComponentPropertyType]*AirspaceGeometryComponentPropertyType
	AirspaceGeometryComponentPropertyTypes_mapString      map[string]*AirspaceGeometryComponentPropertyType
	AirspaceGeometryComponentPropertyTypeOrder            uint
	AirspaceGeometryComponentPropertyType_stagedOrder     map[*AirspaceGeometryComponentPropertyType]uint
	AirspaceGeometryComponentPropertyType_orderStaged     map[uint]*AirspaceGeometryComponentPropertyType
	AirspaceGeometryComponentPropertyTypes_reference      map[*AirspaceGeometryComponentPropertyType]*AirspaceGeometryComponentPropertyType
	AirspaceGeometryComponentPropertyTypes_referenceOrder map[*AirspaceGeometryComponentPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceGeometryComponentPropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceGeometryComponentPropertyType]
	OnAfterAirspaceGeometryComponentPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceGeometryComponentPropertyType]
	OnAfterAirspaceGeometryComponentPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceGeometryComponentPropertyType]

	AirspaceGeometryComponentTypes                map[*AirspaceGeometryComponentType]struct{}
	AirspaceGeometryComponentTypes_instance       map[*AirspaceGeometryComponentType]*AirspaceGeometryComponentType
	AirspaceGeometryComponentTypes_mapString      map[string]*AirspaceGeometryComponentType
	AirspaceGeometryComponentTypeOrder            uint
	AirspaceGeometryComponentType_stagedOrder     map[*AirspaceGeometryComponentType]uint
	AirspaceGeometryComponentType_orderStaged     map[uint]*AirspaceGeometryComponentType
	AirspaceGeometryComponentTypes_reference      map[*AirspaceGeometryComponentType]*AirspaceGeometryComponentType
	AirspaceGeometryComponentTypes_referenceOrder map[*AirspaceGeometryComponentType]uint

	// insertion point for slice of pointers maps
	AirspaceGeometryComponentType_Annotation_reverseMap map[*NotePropertyType]*AirspaceGeometryComponentType

	AirspaceGeometryComponentType_Extension_reverseMap map[*A_extension_2]*AirspaceGeometryComponentType

	OnAfterAirspaceGeometryComponentTypeCreateCallback GongOnAfterCreateInterface[AirspaceGeometryComponentType]
	OnAfterAirspaceGeometryComponentTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceGeometryComponentType]
	OnAfterAirspaceGeometryComponentTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceGeometryComponentType]

	AirspaceLayerClassPropertyTypes                map[*AirspaceLayerClassPropertyType]struct{}
	AirspaceLayerClassPropertyTypes_instance       map[*AirspaceLayerClassPropertyType]*AirspaceLayerClassPropertyType
	AirspaceLayerClassPropertyTypes_mapString      map[string]*AirspaceLayerClassPropertyType
	AirspaceLayerClassPropertyTypeOrder            uint
	AirspaceLayerClassPropertyType_stagedOrder     map[*AirspaceLayerClassPropertyType]uint
	AirspaceLayerClassPropertyType_orderStaged     map[uint]*AirspaceLayerClassPropertyType
	AirspaceLayerClassPropertyTypes_reference      map[*AirspaceLayerClassPropertyType]*AirspaceLayerClassPropertyType
	AirspaceLayerClassPropertyTypes_referenceOrder map[*AirspaceLayerClassPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceLayerClassPropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceLayerClassPropertyType]
	OnAfterAirspaceLayerClassPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceLayerClassPropertyType]
	OnAfterAirspaceLayerClassPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceLayerClassPropertyType]

	AirspaceLayerClassTypes                map[*AirspaceLayerClassType]struct{}
	AirspaceLayerClassTypes_instance       map[*AirspaceLayerClassType]*AirspaceLayerClassType
	AirspaceLayerClassTypes_mapString      map[string]*AirspaceLayerClassType
	AirspaceLayerClassTypeOrder            uint
	AirspaceLayerClassType_stagedOrder     map[*AirspaceLayerClassType]uint
	AirspaceLayerClassType_orderStaged     map[uint]*AirspaceLayerClassType
	AirspaceLayerClassTypes_reference      map[*AirspaceLayerClassType]*AirspaceLayerClassType
	AirspaceLayerClassTypes_referenceOrder map[*AirspaceLayerClassType]uint

	// insertion point for slice of pointers maps
	AirspaceLayerClassType_TimeInterval_reverseMap map[*TimesheetPropertyType]*AirspaceLayerClassType

	AirspaceLayerClassType_Annotation_reverseMap map[*NotePropertyType]*AirspaceLayerClassType

	AirspaceLayerClassType_SpecialDateAuthority_reverseMap map[*OrganisationAuthorityPropertyType]*AirspaceLayerClassType

	AirspaceLayerClassType_AssociatedLevels_reverseMap map[*AirspaceLayerPropertyType]*AirspaceLayerClassType

	AirspaceLayerClassType_Extension_reverseMap map[*A_extension_3]*AirspaceLayerClassType

	OnAfterAirspaceLayerClassTypeCreateCallback GongOnAfterCreateInterface[AirspaceLayerClassType]
	OnAfterAirspaceLayerClassTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceLayerClassType]
	OnAfterAirspaceLayerClassTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceLayerClassType]

	AirspaceLayerPropertyTypes                map[*AirspaceLayerPropertyType]struct{}
	AirspaceLayerPropertyTypes_instance       map[*AirspaceLayerPropertyType]*AirspaceLayerPropertyType
	AirspaceLayerPropertyTypes_mapString      map[string]*AirspaceLayerPropertyType
	AirspaceLayerPropertyTypeOrder            uint
	AirspaceLayerPropertyType_stagedOrder     map[*AirspaceLayerPropertyType]uint
	AirspaceLayerPropertyType_orderStaged     map[uint]*AirspaceLayerPropertyType
	AirspaceLayerPropertyTypes_reference      map[*AirspaceLayerPropertyType]*AirspaceLayerPropertyType
	AirspaceLayerPropertyTypes_referenceOrder map[*AirspaceLayerPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceLayerPropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceLayerPropertyType]
	OnAfterAirspaceLayerPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceLayerPropertyType]
	OnAfterAirspaceLayerPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceLayerPropertyType]

	AirspaceLayerTypes                map[*AirspaceLayerType]struct{}
	AirspaceLayerTypes_instance       map[*AirspaceLayerType]*AirspaceLayerType
	AirspaceLayerTypes_mapString      map[string]*AirspaceLayerType
	AirspaceLayerTypeOrder            uint
	AirspaceLayerType_stagedOrder     map[*AirspaceLayerType]uint
	AirspaceLayerType_orderStaged     map[uint]*AirspaceLayerType
	AirspaceLayerTypes_reference      map[*AirspaceLayerType]*AirspaceLayerType
	AirspaceLayerTypes_referenceOrder map[*AirspaceLayerType]uint

	// insertion point for slice of pointers maps
	AirspaceLayerType_Annotation_reverseMap map[*NotePropertyType]*AirspaceLayerType

	AirspaceLayerType_Extension_reverseMap map[*A_extension_13]*AirspaceLayerType

	OnAfterAirspaceLayerTypeCreateCallback GongOnAfterCreateInterface[AirspaceLayerType]
	OnAfterAirspaceLayerTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceLayerType]
	OnAfterAirspaceLayerTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceLayerType]

	AirspacePropertyTypes                map[*AirspacePropertyType]struct{}
	AirspacePropertyTypes_instance       map[*AirspacePropertyType]*AirspacePropertyType
	AirspacePropertyTypes_mapString      map[string]*AirspacePropertyType
	AirspacePropertyTypeOrder            uint
	AirspacePropertyType_stagedOrder     map[*AirspacePropertyType]uint
	AirspacePropertyType_orderStaged     map[uint]*AirspacePropertyType
	AirspacePropertyTypes_reference      map[*AirspacePropertyType]*AirspacePropertyType
	AirspacePropertyTypes_referenceOrder map[*AirspacePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspacePropertyTypeCreateCallback GongOnAfterCreateInterface[AirspacePropertyType]
	OnAfterAirspacePropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspacePropertyType]
	OnAfterAirspacePropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspacePropertyType]

	AirspaceTimeSlicePropertyTypes                map[*AirspaceTimeSlicePropertyType]struct{}
	AirspaceTimeSlicePropertyTypes_instance       map[*AirspaceTimeSlicePropertyType]*AirspaceTimeSlicePropertyType
	AirspaceTimeSlicePropertyTypes_mapString      map[string]*AirspaceTimeSlicePropertyType
	AirspaceTimeSlicePropertyTypeOrder            uint
	AirspaceTimeSlicePropertyType_stagedOrder     map[*AirspaceTimeSlicePropertyType]uint
	AirspaceTimeSlicePropertyType_orderStaged     map[uint]*AirspaceTimeSlicePropertyType
	AirspaceTimeSlicePropertyTypes_reference      map[*AirspaceTimeSlicePropertyType]*AirspaceTimeSlicePropertyType
	AirspaceTimeSlicePropertyTypes_referenceOrder map[*AirspaceTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceTimeSlicePropertyType]
	OnAfterAirspaceTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceTimeSlicePropertyType]
	OnAfterAirspaceTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceTimeSlicePropertyType]

	AirspaceTimeSliceTypes                map[*AirspaceTimeSliceType]struct{}
	AirspaceTimeSliceTypes_instance       map[*AirspaceTimeSliceType]*AirspaceTimeSliceType
	AirspaceTimeSliceTypes_mapString      map[string]*AirspaceTimeSliceType
	AirspaceTimeSliceTypeOrder            uint
	AirspaceTimeSliceType_stagedOrder     map[*AirspaceTimeSliceType]uint
	AirspaceTimeSliceType_orderStaged     map[uint]*AirspaceTimeSliceType
	AirspaceTimeSliceTypes_reference      map[*AirspaceTimeSliceType]*AirspaceTimeSliceType
	AirspaceTimeSliceTypes_referenceOrder map[*AirspaceTimeSliceType]uint

	// insertion point for slice of pointers maps
	AirspaceTimeSliceType_Class_reverseMap map[*AirspaceLayerClassPropertyType]*AirspaceTimeSliceType

	AirspaceTimeSliceType_GeometryComponent_reverseMap map[*AirspaceGeometryComponentPropertyType]*AirspaceTimeSliceType

	AirspaceTimeSliceType_Activation_reverseMap map[*AirspaceActivationPropertyType]*AirspaceTimeSliceType

	AirspaceTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*AirspaceTimeSliceType

	AirspaceTimeSliceType_Extension_reverseMap map[*A_extension]*AirspaceTimeSliceType

	OnAfterAirspaceTimeSliceTypeCreateCallback GongOnAfterCreateInterface[AirspaceTimeSliceType]
	OnAfterAirspaceTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceTimeSliceType]
	OnAfterAirspaceTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceTimeSliceType]

	AirspaceTypes                map[*AirspaceType]struct{}
	AirspaceTypes_instance       map[*AirspaceType]*AirspaceType
	AirspaceTypes_mapString      map[string]*AirspaceType
	AirspaceTypeOrder            uint
	AirspaceType_stagedOrder     map[*AirspaceType]uint
	AirspaceType_orderStaged     map[uint]*AirspaceType
	AirspaceTypes_reference      map[*AirspaceType]*AirspaceType
	AirspaceTypes_referenceOrder map[*AirspaceType]uint

	// insertion point for slice of pointers maps
	AirspaceType_TimeSlice_reverseMap map[*AirspaceTimeSlicePropertyType]*AirspaceType

	OnAfterAirspaceTypeCreateCallback GongOnAfterCreateInterface[AirspaceType]
	OnAfterAirspaceTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceType]
	OnAfterAirspaceTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceType]

	AirspaceVolumeDependencyPropertyTypes                map[*AirspaceVolumeDependencyPropertyType]struct{}
	AirspaceVolumeDependencyPropertyTypes_instance       map[*AirspaceVolumeDependencyPropertyType]*AirspaceVolumeDependencyPropertyType
	AirspaceVolumeDependencyPropertyTypes_mapString      map[string]*AirspaceVolumeDependencyPropertyType
	AirspaceVolumeDependencyPropertyTypeOrder            uint
	AirspaceVolumeDependencyPropertyType_stagedOrder     map[*AirspaceVolumeDependencyPropertyType]uint
	AirspaceVolumeDependencyPropertyType_orderStaged     map[uint]*AirspaceVolumeDependencyPropertyType
	AirspaceVolumeDependencyPropertyTypes_reference      map[*AirspaceVolumeDependencyPropertyType]*AirspaceVolumeDependencyPropertyType
	AirspaceVolumeDependencyPropertyTypes_referenceOrder map[*AirspaceVolumeDependencyPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceVolumeDependencyPropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceVolumeDependencyPropertyType]
	OnAfterAirspaceVolumeDependencyPropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceVolumeDependencyPropertyType]
	OnAfterAirspaceVolumeDependencyPropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceVolumeDependencyPropertyType]

	AirspaceVolumeDependencyTypes                map[*AirspaceVolumeDependencyType]struct{}
	AirspaceVolumeDependencyTypes_instance       map[*AirspaceVolumeDependencyType]*AirspaceVolumeDependencyType
	AirspaceVolumeDependencyTypes_mapString      map[string]*AirspaceVolumeDependencyType
	AirspaceVolumeDependencyTypeOrder            uint
	AirspaceVolumeDependencyType_stagedOrder     map[*AirspaceVolumeDependencyType]uint
	AirspaceVolumeDependencyType_orderStaged     map[uint]*AirspaceVolumeDependencyType
	AirspaceVolumeDependencyTypes_reference      map[*AirspaceVolumeDependencyType]*AirspaceVolumeDependencyType
	AirspaceVolumeDependencyTypes_referenceOrder map[*AirspaceVolumeDependencyType]uint

	// insertion point for slice of pointers maps
	AirspaceVolumeDependencyType_Annotation_reverseMap map[*NotePropertyType]*AirspaceVolumeDependencyType

	AirspaceVolumeDependencyType_Extension_reverseMap map[*A_extension_5]*AirspaceVolumeDependencyType

	OnAfterAirspaceVolumeDependencyTypeCreateCallback GongOnAfterCreateInterface[AirspaceVolumeDependencyType]
	OnAfterAirspaceVolumeDependencyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceVolumeDependencyType]
	OnAfterAirspaceVolumeDependencyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceVolumeDependencyType]

	AirspaceVolumePropertyTypes                map[*AirspaceVolumePropertyType]struct{}
	AirspaceVolumePropertyTypes_instance       map[*AirspaceVolumePropertyType]*AirspaceVolumePropertyType
	AirspaceVolumePropertyTypes_mapString      map[string]*AirspaceVolumePropertyType
	AirspaceVolumePropertyTypeOrder            uint
	AirspaceVolumePropertyType_stagedOrder     map[*AirspaceVolumePropertyType]uint
	AirspaceVolumePropertyType_orderStaged     map[uint]*AirspaceVolumePropertyType
	AirspaceVolumePropertyTypes_reference      map[*AirspaceVolumePropertyType]*AirspaceVolumePropertyType
	AirspaceVolumePropertyTypes_referenceOrder map[*AirspaceVolumePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAirspaceVolumePropertyTypeCreateCallback GongOnAfterCreateInterface[AirspaceVolumePropertyType]
	OnAfterAirspaceVolumePropertyTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceVolumePropertyType]
	OnAfterAirspaceVolumePropertyTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceVolumePropertyType]

	AirspaceVolumeTypes                map[*AirspaceVolumeType]struct{}
	AirspaceVolumeTypes_instance       map[*AirspaceVolumeType]*AirspaceVolumeType
	AirspaceVolumeTypes_mapString      map[string]*AirspaceVolumeType
	AirspaceVolumeTypeOrder            uint
	AirspaceVolumeType_stagedOrder     map[*AirspaceVolumeType]uint
	AirspaceVolumeType_orderStaged     map[uint]*AirspaceVolumeType
	AirspaceVolumeTypes_reference      map[*AirspaceVolumeType]*AirspaceVolumeType
	AirspaceVolumeTypes_referenceOrder map[*AirspaceVolumeType]uint

	// insertion point for slice of pointers maps
	AirspaceVolumeType_Annotation_reverseMap map[*NotePropertyType]*AirspaceVolumeType

	AirspaceVolumeType_Extension_reverseMap map[*A_extension_4]*AirspaceVolumeType

	OnAfterAirspaceVolumeTypeCreateCallback GongOnAfterCreateInterface[AirspaceVolumeType]
	OnAfterAirspaceVolumeTypeUpdateCallback GongOnAfterUpdateInterface[AirspaceVolumeType]
	OnAfterAirspaceVolumeTypeDeleteCallback GongOnAfterDeleteInterface[AirspaceVolumeType]

	AuthorityForAirspaceTimeSlicePropertyTypes                map[*AuthorityForAirspaceTimeSlicePropertyType]struct{}
	AuthorityForAirspaceTimeSlicePropertyTypes_instance       map[*AuthorityForAirspaceTimeSlicePropertyType]*AuthorityForAirspaceTimeSlicePropertyType
	AuthorityForAirspaceTimeSlicePropertyTypes_mapString      map[string]*AuthorityForAirspaceTimeSlicePropertyType
	AuthorityForAirspaceTimeSlicePropertyTypeOrder            uint
	AuthorityForAirspaceTimeSlicePropertyType_stagedOrder     map[*AuthorityForAirspaceTimeSlicePropertyType]uint
	AuthorityForAirspaceTimeSlicePropertyType_orderStaged     map[uint]*AuthorityForAirspaceTimeSlicePropertyType
	AuthorityForAirspaceTimeSlicePropertyTypes_reference      map[*AuthorityForAirspaceTimeSlicePropertyType]*AuthorityForAirspaceTimeSlicePropertyType
	AuthorityForAirspaceTimeSlicePropertyTypes_referenceOrder map[*AuthorityForAirspaceTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterAuthorityForAirspaceTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[AuthorityForAirspaceTimeSlicePropertyType]
	OnAfterAuthorityForAirspaceTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[AuthorityForAirspaceTimeSlicePropertyType]
	OnAfterAuthorityForAirspaceTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[AuthorityForAirspaceTimeSlicePropertyType]

	AuthorityForAirspaceTimeSliceTypes                map[*AuthorityForAirspaceTimeSliceType]struct{}
	AuthorityForAirspaceTimeSliceTypes_instance       map[*AuthorityForAirspaceTimeSliceType]*AuthorityForAirspaceTimeSliceType
	AuthorityForAirspaceTimeSliceTypes_mapString      map[string]*AuthorityForAirspaceTimeSliceType
	AuthorityForAirspaceTimeSliceTypeOrder            uint
	AuthorityForAirspaceTimeSliceType_stagedOrder     map[*AuthorityForAirspaceTimeSliceType]uint
	AuthorityForAirspaceTimeSliceType_orderStaged     map[uint]*AuthorityForAirspaceTimeSliceType
	AuthorityForAirspaceTimeSliceTypes_reference      map[*AuthorityForAirspaceTimeSliceType]*AuthorityForAirspaceTimeSliceType
	AuthorityForAirspaceTimeSliceTypes_referenceOrder map[*AuthorityForAirspaceTimeSliceType]uint

	// insertion point for slice of pointers maps
	AuthorityForAirspaceTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*AuthorityForAirspaceTimeSliceType

	AuthorityForAirspaceTimeSliceType_Extension_reverseMap map[*A_extension_6]*AuthorityForAirspaceTimeSliceType

	OnAfterAuthorityForAirspaceTimeSliceTypeCreateCallback GongOnAfterCreateInterface[AuthorityForAirspaceTimeSliceType]
	OnAfterAuthorityForAirspaceTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[AuthorityForAirspaceTimeSliceType]
	OnAfterAuthorityForAirspaceTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[AuthorityForAirspaceTimeSliceType]

	AuthorityForAirspaceTypes                map[*AuthorityForAirspaceType]struct{}
	AuthorityForAirspaceTypes_instance       map[*AuthorityForAirspaceType]*AuthorityForAirspaceType
	AuthorityForAirspaceTypes_mapString      map[string]*AuthorityForAirspaceType
	AuthorityForAirspaceTypeOrder            uint
	AuthorityForAirspaceType_stagedOrder     map[*AuthorityForAirspaceType]uint
	AuthorityForAirspaceType_orderStaged     map[uint]*AuthorityForAirspaceType
	AuthorityForAirspaceTypes_reference      map[*AuthorityForAirspaceType]*AuthorityForAirspaceType
	AuthorityForAirspaceTypes_referenceOrder map[*AuthorityForAirspaceType]uint

	// insertion point for slice of pointers maps
	AuthorityForAirspaceType_TimeSlice_reverseMap map[*AuthorityForAirspaceTimeSlicePropertyType]*AuthorityForAirspaceType

	OnAfterAuthorityForAirspaceTypeCreateCallback GongOnAfterCreateInterface[AuthorityForAirspaceType]
	OnAfterAuthorityForAirspaceTypeUpdateCallback GongOnAfterUpdateInterface[AuthorityForAirspaceType]
	OnAfterAuthorityForAirspaceTypeDeleteCallback GongOnAfterDeleteInterface[AuthorityForAirspaceType]

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

	ElevatedCurveType_Extension_reverseMap map[*A_extension_8]*ElevatedCurveType

	OnAfterElevatedCurveTypeCreateCallback GongOnAfterCreateInterface[ElevatedCurveType]
	OnAfterElevatedCurveTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedCurveType]
	OnAfterElevatedCurveTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedCurveType]

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

	ElevatedSurfaceType_Extension_reverseMap map[*A_extension_9]*ElevatedSurfaceType

	OnAfterElevatedSurfaceTypeCreateCallback GongOnAfterCreateInterface[ElevatedSurfaceType]
	OnAfterElevatedSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[ElevatedSurfaceType]
	OnAfterElevatedSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[ElevatedSurfaceType]

	GeoBorderTimeSlicePropertyTypes                map[*GeoBorderTimeSlicePropertyType]struct{}
	GeoBorderTimeSlicePropertyTypes_instance       map[*GeoBorderTimeSlicePropertyType]*GeoBorderTimeSlicePropertyType
	GeoBorderTimeSlicePropertyTypes_mapString      map[string]*GeoBorderTimeSlicePropertyType
	GeoBorderTimeSlicePropertyTypeOrder            uint
	GeoBorderTimeSlicePropertyType_stagedOrder     map[*GeoBorderTimeSlicePropertyType]uint
	GeoBorderTimeSlicePropertyType_orderStaged     map[uint]*GeoBorderTimeSlicePropertyType
	GeoBorderTimeSlicePropertyTypes_reference      map[*GeoBorderTimeSlicePropertyType]*GeoBorderTimeSlicePropertyType
	GeoBorderTimeSlicePropertyTypes_referenceOrder map[*GeoBorderTimeSlicePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterGeoBorderTimeSlicePropertyTypeCreateCallback GongOnAfterCreateInterface[GeoBorderTimeSlicePropertyType]
	OnAfterGeoBorderTimeSlicePropertyTypeUpdateCallback GongOnAfterUpdateInterface[GeoBorderTimeSlicePropertyType]
	OnAfterGeoBorderTimeSlicePropertyTypeDeleteCallback GongOnAfterDeleteInterface[GeoBorderTimeSlicePropertyType]

	GeoBorderTimeSliceTypes                map[*GeoBorderTimeSliceType]struct{}
	GeoBorderTimeSliceTypes_instance       map[*GeoBorderTimeSliceType]*GeoBorderTimeSliceType
	GeoBorderTimeSliceTypes_mapString      map[string]*GeoBorderTimeSliceType
	GeoBorderTimeSliceTypeOrder            uint
	GeoBorderTimeSliceType_stagedOrder     map[*GeoBorderTimeSliceType]uint
	GeoBorderTimeSliceType_orderStaged     map[uint]*GeoBorderTimeSliceType
	GeoBorderTimeSliceTypes_reference      map[*GeoBorderTimeSliceType]*GeoBorderTimeSliceType
	GeoBorderTimeSliceTypes_referenceOrder map[*GeoBorderTimeSliceType]uint

	// insertion point for slice of pointers maps
	GeoBorderTimeSliceType_Annotation_reverseMap map[*NotePropertyType]*GeoBorderTimeSliceType

	GeoBorderTimeSliceType_Extension_reverseMap map[*A_extension_7]*GeoBorderTimeSliceType

	OnAfterGeoBorderTimeSliceTypeCreateCallback GongOnAfterCreateInterface[GeoBorderTimeSliceType]
	OnAfterGeoBorderTimeSliceTypeUpdateCallback GongOnAfterUpdateInterface[GeoBorderTimeSliceType]
	OnAfterGeoBorderTimeSliceTypeDeleteCallback GongOnAfterDeleteInterface[GeoBorderTimeSliceType]

	GeoBorderTypes                map[*GeoBorderType]struct{}
	GeoBorderTypes_instance       map[*GeoBorderType]*GeoBorderType
	GeoBorderTypes_mapString      map[string]*GeoBorderType
	GeoBorderTypeOrder            uint
	GeoBorderType_stagedOrder     map[*GeoBorderType]uint
	GeoBorderType_orderStaged     map[uint]*GeoBorderType
	GeoBorderTypes_reference      map[*GeoBorderType]*GeoBorderType
	GeoBorderTypes_referenceOrder map[*GeoBorderType]uint

	// insertion point for slice of pointers maps
	GeoBorderType_TimeSlice_reverseMap map[*GeoBorderTimeSlicePropertyType]*GeoBorderType

	OnAfterGeoBorderTypeCreateCallback GongOnAfterCreateInterface[GeoBorderType]
	OnAfterGeoBorderTypeUpdateCallback GongOnAfterUpdateInterface[GeoBorderType]
	OnAfterGeoBorderTypeDeleteCallback GongOnAfterDeleteInterface[GeoBorderType]

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
	LinguisticNoteType_Extension_reverseMap map[*A_extension_10]*LinguisticNoteType

	OnAfterLinguisticNoteTypeCreateCallback GongOnAfterCreateInterface[LinguisticNoteType]
	OnAfterLinguisticNoteTypeUpdateCallback GongOnAfterUpdateInterface[LinguisticNoteType]
	OnAfterLinguisticNoteTypeDeleteCallback GongOnAfterDeleteInterface[LinguisticNoteType]

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

	NoteType_Extension_reverseMap map[*A_extension_11]*NoteType

	OnAfterNoteTypeCreateCallback GongOnAfterCreateInterface[NoteType]
	OnAfterNoteTypeUpdateCallback GongOnAfterUpdateInterface[NoteType]
	OnAfterNoteTypeDeleteCallback GongOnAfterDeleteInterface[NoteType]

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

	RoutePropertyTypes                map[*RoutePropertyType]struct{}
	RoutePropertyTypes_instance       map[*RoutePropertyType]*RoutePropertyType
	RoutePropertyTypes_mapString      map[string]*RoutePropertyType
	RoutePropertyTypeOrder            uint
	RoutePropertyType_stagedOrder     map[*RoutePropertyType]uint
	RoutePropertyType_orderStaged     map[uint]*RoutePropertyType
	RoutePropertyTypes_reference      map[*RoutePropertyType]*RoutePropertyType
	RoutePropertyTypes_referenceOrder map[*RoutePropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterRoutePropertyTypeCreateCallback GongOnAfterCreateInterface[RoutePropertyType]
	OnAfterRoutePropertyTypeUpdateCallback GongOnAfterUpdateInterface[RoutePropertyType]
	OnAfterRoutePropertyTypeDeleteCallback GongOnAfterDeleteInterface[RoutePropertyType]

	StandardLevelColumnPropertyTypes                map[*StandardLevelColumnPropertyType]struct{}
	StandardLevelColumnPropertyTypes_instance       map[*StandardLevelColumnPropertyType]*StandardLevelColumnPropertyType
	StandardLevelColumnPropertyTypes_mapString      map[string]*StandardLevelColumnPropertyType
	StandardLevelColumnPropertyTypeOrder            uint
	StandardLevelColumnPropertyType_stagedOrder     map[*StandardLevelColumnPropertyType]uint
	StandardLevelColumnPropertyType_orderStaged     map[uint]*StandardLevelColumnPropertyType
	StandardLevelColumnPropertyTypes_reference      map[*StandardLevelColumnPropertyType]*StandardLevelColumnPropertyType
	StandardLevelColumnPropertyTypes_referenceOrder map[*StandardLevelColumnPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterStandardLevelColumnPropertyTypeCreateCallback GongOnAfterCreateInterface[StandardLevelColumnPropertyType]
	OnAfterStandardLevelColumnPropertyTypeUpdateCallback GongOnAfterUpdateInterface[StandardLevelColumnPropertyType]
	OnAfterStandardLevelColumnPropertyTypeDeleteCallback GongOnAfterDeleteInterface[StandardLevelColumnPropertyType]

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

	TimesheetType_Extension_reverseMap map[*A_extension_14]*TimesheetType

	OnAfterTimesheetTypeCreateCallback GongOnAfterCreateInterface[TimesheetType]
	OnAfterTimesheetTypeUpdateCallback GongOnAfterUpdateInterface[TimesheetType]
	OnAfterTimesheetTypeDeleteCallback GongOnAfterDeleteInterface[TimesheetType]

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

	__gong__clearReferences(&stage.A_extension_13s_reference, &stage.A_extension_13s_instance, &stage.A_extension_13s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_14s_reference, &stage.A_extension_14s_instance, &stage.A_extension_14s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_2s_reference, &stage.A_extension_2s_instance, &stage.A_extension_2s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_3s_reference, &stage.A_extension_3s_instance, &stage.A_extension_3s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_4s_reference, &stage.A_extension_4s_instance, &stage.A_extension_4s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_5s_reference, &stage.A_extension_5s_instance, &stage.A_extension_5s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_6s_reference, &stage.A_extension_6s_instance, &stage.A_extension_6s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_7s_reference, &stage.A_extension_7s_instance, &stage.A_extension_7s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_8s_reference, &stage.A_extension_8s_instance, &stage.A_extension_8s_referenceOrder)

	__gong__clearReferences(&stage.A_extension_9s_reference, &stage.A_extension_9s_instance, &stage.A_extension_9s_referenceOrder)

	__gong__clearReferences(&stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_instance, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder)

	__gong__clearReferences(&stage.AircraftCharacteristicPropertyTypes_reference, &stage.AircraftCharacteristicPropertyTypes_instance, &stage.AircraftCharacteristicPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AircraftCharacteristicTypes_reference, &stage.AircraftCharacteristicTypes_instance, &stage.AircraftCharacteristicTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceActivationPropertyTypes_reference, &stage.AirspaceActivationPropertyTypes_instance, &stage.AirspaceActivationPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceActivationTypes_reference, &stage.AirspaceActivationTypes_instance, &stage.AirspaceActivationTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceGeometryComponentPropertyTypes_reference, &stage.AirspaceGeometryComponentPropertyTypes_instance, &stage.AirspaceGeometryComponentPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceGeometryComponentTypes_reference, &stage.AirspaceGeometryComponentTypes_instance, &stage.AirspaceGeometryComponentTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceLayerClassPropertyTypes_reference, &stage.AirspaceLayerClassPropertyTypes_instance, &stage.AirspaceLayerClassPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceLayerClassTypes_reference, &stage.AirspaceLayerClassTypes_instance, &stage.AirspaceLayerClassTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceLayerPropertyTypes_reference, &stage.AirspaceLayerPropertyTypes_instance, &stage.AirspaceLayerPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceLayerTypes_reference, &stage.AirspaceLayerTypes_instance, &stage.AirspaceLayerTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspacePropertyTypes_reference, &stage.AirspacePropertyTypes_instance, &stage.AirspacePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceTimeSlicePropertyTypes_reference, &stage.AirspaceTimeSlicePropertyTypes_instance, &stage.AirspaceTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceTimeSliceTypes_reference, &stage.AirspaceTimeSliceTypes_instance, &stage.AirspaceTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceTypes_reference, &stage.AirspaceTypes_instance, &stage.AirspaceTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceVolumeDependencyPropertyTypes_reference, &stage.AirspaceVolumeDependencyPropertyTypes_instance, &stage.AirspaceVolumeDependencyPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceVolumeDependencyTypes_reference, &stage.AirspaceVolumeDependencyTypes_instance, &stage.AirspaceVolumeDependencyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceVolumePropertyTypes_reference, &stage.AirspaceVolumePropertyTypes_instance, &stage.AirspaceVolumePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AirspaceVolumeTypes_reference, &stage.AirspaceVolumeTypes_instance, &stage.AirspaceVolumeTypes_referenceOrder)

	__gong__clearReferences(&stage.AuthorityForAirspaceTimeSlicePropertyTypes_reference, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_instance, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.AuthorityForAirspaceTimeSliceTypes_reference, &stage.AuthorityForAirspaceTimeSliceTypes_instance, &stage.AuthorityForAirspaceTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.AuthorityForAirspaceTypes_reference, &stage.AuthorityForAirspaceTypes_instance, &stage.AuthorityForAirspaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CurvePropertyTypes_reference, &stage.CurvePropertyTypes_instance, &stage.CurvePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.CurveTypes_reference, &stage.CurveTypes_instance, &stage.CurveTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_instance, &stage.ElevatedCurveTypes_referenceOrder)

	__gong__clearReferences(&stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_instance, &stage.ElevatedSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.GeoBorderTimeSlicePropertyTypes_reference, &stage.GeoBorderTimeSlicePropertyTypes_instance, &stage.GeoBorderTimeSlicePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.GeoBorderTimeSliceTypes_reference, &stage.GeoBorderTimeSliceTypes_instance, &stage.GeoBorderTimeSliceTypes_referenceOrder)

	__gong__clearReferences(&stage.GeoBorderTypes_reference, &stage.GeoBorderTypes_instance, &stage.GeoBorderTypes_referenceOrder)

	__gong__clearReferences(&stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_instance, &stage.LinguisticNotePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_instance, &stage.LinguisticNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_instance, &stage.NotePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.NoteTypes_reference, &stage.NoteTypes_instance, &stage.NoteTypes_referenceOrder)

	__gong__clearReferences(&stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_instance, &stage.OrganisationAuthorityPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.RoutePropertyTypes_reference, &stage.RoutePropertyTypes_instance, &stage.RoutePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.StandardLevelColumnPropertyTypes_reference, &stage.StandardLevelColumnPropertyTypes_instance, &stage.StandardLevelColumnPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_instance, &stage.SurfacePropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.SurfaceTypes_reference, &stage.SurfaceTypes_instance, &stage.SurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_instance, &stage.TimesheetPropertyTypes_referenceOrder)

	__gong__clearReferences(&stage.TimesheetTypes_reference, &stage.TimesheetTypes_instance, &stage.TimesheetTypes_referenceOrder)

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

	stage.A_extension_13Order = __gong__recomputeOrder(stage.A_extension_13_stagedOrder)

	stage.A_extension_14Order = __gong__recomputeOrder(stage.A_extension_14_stagedOrder)

	stage.A_extension_2Order = __gong__recomputeOrder(stage.A_extension_2_stagedOrder)

	stage.A_extension_3Order = __gong__recomputeOrder(stage.A_extension_3_stagedOrder)

	stage.A_extension_4Order = __gong__recomputeOrder(stage.A_extension_4_stagedOrder)

	stage.A_extension_5Order = __gong__recomputeOrder(stage.A_extension_5_stagedOrder)

	stage.A_extension_6Order = __gong__recomputeOrder(stage.A_extension_6_stagedOrder)

	stage.A_extension_7Order = __gong__recomputeOrder(stage.A_extension_7_stagedOrder)

	stage.A_extension_8Order = __gong__recomputeOrder(stage.A_extension_8_stagedOrder)

	stage.A_extension_9Order = __gong__recomputeOrder(stage.A_extension_9_stagedOrder)

	stage.AbstractPropertiesWithScheduleTypeOrder = __gong__recomputeOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder)

	stage.AircraftCharacteristicPropertyTypeOrder = __gong__recomputeOrder(stage.AircraftCharacteristicPropertyType_stagedOrder)

	stage.AircraftCharacteristicTypeOrder = __gong__recomputeOrder(stage.AircraftCharacteristicType_stagedOrder)

	stage.AirspaceActivationPropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceActivationPropertyType_stagedOrder)

	stage.AirspaceActivationTypeOrder = __gong__recomputeOrder(stage.AirspaceActivationType_stagedOrder)

	stage.AirspaceGeometryComponentPropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceGeometryComponentPropertyType_stagedOrder)

	stage.AirspaceGeometryComponentTypeOrder = __gong__recomputeOrder(stage.AirspaceGeometryComponentType_stagedOrder)

	stage.AirspaceLayerClassPropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceLayerClassPropertyType_stagedOrder)

	stage.AirspaceLayerClassTypeOrder = __gong__recomputeOrder(stage.AirspaceLayerClassType_stagedOrder)

	stage.AirspaceLayerPropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceLayerPropertyType_stagedOrder)

	stage.AirspaceLayerTypeOrder = __gong__recomputeOrder(stage.AirspaceLayerType_stagedOrder)

	stage.AirspacePropertyTypeOrder = __gong__recomputeOrder(stage.AirspacePropertyType_stagedOrder)

	stage.AirspaceTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceTimeSlicePropertyType_stagedOrder)

	stage.AirspaceTimeSliceTypeOrder = __gong__recomputeOrder(stage.AirspaceTimeSliceType_stagedOrder)

	stage.AirspaceTypeOrder = __gong__recomputeOrder(stage.AirspaceType_stagedOrder)

	stage.AirspaceVolumeDependencyPropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceVolumeDependencyPropertyType_stagedOrder)

	stage.AirspaceVolumeDependencyTypeOrder = __gong__recomputeOrder(stage.AirspaceVolumeDependencyType_stagedOrder)

	stage.AirspaceVolumePropertyTypeOrder = __gong__recomputeOrder(stage.AirspaceVolumePropertyType_stagedOrder)

	stage.AirspaceVolumeTypeOrder = __gong__recomputeOrder(stage.AirspaceVolumeType_stagedOrder)

	stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder)

	stage.AuthorityForAirspaceTimeSliceTypeOrder = __gong__recomputeOrder(stage.AuthorityForAirspaceTimeSliceType_stagedOrder)

	stage.AuthorityForAirspaceTypeOrder = __gong__recomputeOrder(stage.AuthorityForAirspaceType_stagedOrder)

	stage.CurvePropertyTypeOrder = __gong__recomputeOrder(stage.CurvePropertyType_stagedOrder)

	stage.CurveTypeOrder = __gong__recomputeOrder(stage.CurveType_stagedOrder)

	stage.ElevatedCurveTypeOrder = __gong__recomputeOrder(stage.ElevatedCurveType_stagedOrder)

	stage.ElevatedSurfaceTypeOrder = __gong__recomputeOrder(stage.ElevatedSurfaceType_stagedOrder)

	stage.GeoBorderTimeSlicePropertyTypeOrder = __gong__recomputeOrder(stage.GeoBorderTimeSlicePropertyType_stagedOrder)

	stage.GeoBorderTimeSliceTypeOrder = __gong__recomputeOrder(stage.GeoBorderTimeSliceType_stagedOrder)

	stage.GeoBorderTypeOrder = __gong__recomputeOrder(stage.GeoBorderType_stagedOrder)

	stage.LinguisticNotePropertyTypeOrder = __gong__recomputeOrder(stage.LinguisticNotePropertyType_stagedOrder)

	stage.LinguisticNoteTypeOrder = __gong__recomputeOrder(stage.LinguisticNoteType_stagedOrder)

	stage.NotePropertyTypeOrder = __gong__recomputeOrder(stage.NotePropertyType_stagedOrder)

	stage.NoteTypeOrder = __gong__recomputeOrder(stage.NoteType_stagedOrder)

	stage.OrganisationAuthorityPropertyTypeOrder = __gong__recomputeOrder(stage.OrganisationAuthorityPropertyType_stagedOrder)

	stage.RoutePropertyTypeOrder = __gong__recomputeOrder(stage.RoutePropertyType_stagedOrder)

	stage.StandardLevelColumnPropertyTypeOrder = __gong__recomputeOrder(stage.StandardLevelColumnPropertyType_stagedOrder)

	stage.SurfacePropertyTypeOrder = __gong__recomputeOrder(stage.SurfacePropertyType_stagedOrder)

	stage.SurfaceTypeOrder = __gong__recomputeOrder(stage.SurfaceType_stagedOrder)

	stage.TimesheetPropertyTypeOrder = __gong__recomputeOrder(stage.TimesheetPropertyType_stagedOrder)

	stage.TimesheetTypeOrder = __gong__recomputeOrder(stage.TimesheetType_stagedOrder)

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

func (*A_extension_13) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_13s, stage.A_extension_13_stagedOrder)
}

func (*A_extension_13) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_13_orderStaged[order]
}

func (*A_extension_13) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_13s_mapString
}

func (*A_extension_13) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_13s
}

func (*A_extension_13) GongNewInstance() any {
	return new(A_extension_13)
}

func (*A_extension_14) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.A_extension_14s, stage.A_extension_14_stagedOrder)
}

func (*A_extension_14) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.A_extension_14_orderStaged[order]
}

func (*A_extension_14) GongGetInstancesMapByName(stage *Stage) any {
	return stage.A_extension_14s_mapString
}

func (*A_extension_14) GongGetInstancesSet(stage *Stage) any {
	return &stage.A_extension_14s
}

func (*A_extension_14) GongNewInstance() any {
	return new(A_extension_14)
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

func (*AircraftCharacteristicPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyType_stagedOrder)
}

func (*AircraftCharacteristicPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AircraftCharacteristicPropertyType_orderStaged[order]
}

func (*AircraftCharacteristicPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AircraftCharacteristicPropertyTypes_mapString
}

func (*AircraftCharacteristicPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AircraftCharacteristicPropertyTypes
}

func (*AircraftCharacteristicPropertyType) GongNewInstance() any {
	return new(AircraftCharacteristicPropertyType)
}

func (*AircraftCharacteristicType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicType_stagedOrder)
}

func (*AircraftCharacteristicType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AircraftCharacteristicType_orderStaged[order]
}

func (*AircraftCharacteristicType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AircraftCharacteristicTypes_mapString
}

func (*AircraftCharacteristicType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AircraftCharacteristicTypes
}

func (*AircraftCharacteristicType) GongNewInstance() any {
	return new(AircraftCharacteristicType)
}

func (*AirspaceActivationPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyType_stagedOrder)
}

func (*AirspaceActivationPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceActivationPropertyType_orderStaged[order]
}

func (*AirspaceActivationPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceActivationPropertyTypes_mapString
}

func (*AirspaceActivationPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceActivationPropertyTypes
}

func (*AirspaceActivationPropertyType) GongNewInstance() any {
	return new(AirspaceActivationPropertyType)
}

func (*AirspaceActivationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceActivationTypes, stage.AirspaceActivationType_stagedOrder)
}

func (*AirspaceActivationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceActivationType_orderStaged[order]
}

func (*AirspaceActivationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceActivationTypes_mapString
}

func (*AirspaceActivationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceActivationTypes
}

func (*AirspaceActivationType) GongNewInstance() any {
	return new(AirspaceActivationType)
}

func (*AirspaceGeometryComponentPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyType_stagedOrder)
}

func (*AirspaceGeometryComponentPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceGeometryComponentPropertyType_orderStaged[order]
}

func (*AirspaceGeometryComponentPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceGeometryComponentPropertyTypes_mapString
}

func (*AirspaceGeometryComponentPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceGeometryComponentPropertyTypes
}

func (*AirspaceGeometryComponentPropertyType) GongNewInstance() any {
	return new(AirspaceGeometryComponentPropertyType)
}

func (*AirspaceGeometryComponentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentType_stagedOrder)
}

func (*AirspaceGeometryComponentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceGeometryComponentType_orderStaged[order]
}

func (*AirspaceGeometryComponentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceGeometryComponentTypes_mapString
}

func (*AirspaceGeometryComponentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceGeometryComponentTypes
}

func (*AirspaceGeometryComponentType) GongNewInstance() any {
	return new(AirspaceGeometryComponentType)
}

func (*AirspaceLayerClassPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyType_stagedOrder)
}

func (*AirspaceLayerClassPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceLayerClassPropertyType_orderStaged[order]
}

func (*AirspaceLayerClassPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceLayerClassPropertyTypes_mapString
}

func (*AirspaceLayerClassPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceLayerClassPropertyTypes
}

func (*AirspaceLayerClassPropertyType) GongNewInstance() any {
	return new(AirspaceLayerClassPropertyType)
}

func (*AirspaceLayerClassType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassType_stagedOrder)
}

func (*AirspaceLayerClassType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceLayerClassType_orderStaged[order]
}

func (*AirspaceLayerClassType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceLayerClassTypes_mapString
}

func (*AirspaceLayerClassType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceLayerClassTypes
}

func (*AirspaceLayerClassType) GongNewInstance() any {
	return new(AirspaceLayerClassType)
}

func (*AirspaceLayerPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyType_stagedOrder)
}

func (*AirspaceLayerPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceLayerPropertyType_orderStaged[order]
}

func (*AirspaceLayerPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceLayerPropertyTypes_mapString
}

func (*AirspaceLayerPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceLayerPropertyTypes
}

func (*AirspaceLayerPropertyType) GongNewInstance() any {
	return new(AirspaceLayerPropertyType)
}

func (*AirspaceLayerType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceLayerTypes, stage.AirspaceLayerType_stagedOrder)
}

func (*AirspaceLayerType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceLayerType_orderStaged[order]
}

func (*AirspaceLayerType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceLayerTypes_mapString
}

func (*AirspaceLayerType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceLayerTypes
}

func (*AirspaceLayerType) GongNewInstance() any {
	return new(AirspaceLayerType)
}

func (*AirspacePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspacePropertyTypes, stage.AirspacePropertyType_stagedOrder)
}

func (*AirspacePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspacePropertyType_orderStaged[order]
}

func (*AirspacePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspacePropertyTypes_mapString
}

func (*AirspacePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspacePropertyTypes
}

func (*AirspacePropertyType) GongNewInstance() any {
	return new(AirspacePropertyType)
}

func (*AirspaceTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyType_stagedOrder)
}

func (*AirspaceTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceTimeSlicePropertyType_orderStaged[order]
}

func (*AirspaceTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceTimeSlicePropertyTypes_mapString
}

func (*AirspaceTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceTimeSlicePropertyTypes
}

func (*AirspaceTimeSlicePropertyType) GongNewInstance() any {
	return new(AirspaceTimeSlicePropertyType)
}

func (*AirspaceTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceType_stagedOrder)
}

func (*AirspaceTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceTimeSliceType_orderStaged[order]
}

func (*AirspaceTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceTimeSliceTypes_mapString
}

func (*AirspaceTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceTimeSliceTypes
}

func (*AirspaceTimeSliceType) GongNewInstance() any {
	return new(AirspaceTimeSliceType)
}

func (*AirspaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceTypes, stage.AirspaceType_stagedOrder)
}

func (*AirspaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceType_orderStaged[order]
}

func (*AirspaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceTypes_mapString
}

func (*AirspaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceTypes
}

func (*AirspaceType) GongNewInstance() any {
	return new(AirspaceType)
}

func (*AirspaceVolumeDependencyPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyType_stagedOrder)
}

func (*AirspaceVolumeDependencyPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceVolumeDependencyPropertyType_orderStaged[order]
}

func (*AirspaceVolumeDependencyPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceVolumeDependencyPropertyTypes_mapString
}

func (*AirspaceVolumeDependencyPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceVolumeDependencyPropertyTypes
}

func (*AirspaceVolumeDependencyPropertyType) GongNewInstance() any {
	return new(AirspaceVolumeDependencyPropertyType)
}

func (*AirspaceVolumeDependencyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyType_stagedOrder)
}

func (*AirspaceVolumeDependencyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceVolumeDependencyType_orderStaged[order]
}

func (*AirspaceVolumeDependencyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceVolumeDependencyTypes_mapString
}

func (*AirspaceVolumeDependencyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceVolumeDependencyTypes
}

func (*AirspaceVolumeDependencyType) GongNewInstance() any {
	return new(AirspaceVolumeDependencyType)
}

func (*AirspaceVolumePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyType_stagedOrder)
}

func (*AirspaceVolumePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceVolumePropertyType_orderStaged[order]
}

func (*AirspaceVolumePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceVolumePropertyTypes_mapString
}

func (*AirspaceVolumePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceVolumePropertyTypes
}

func (*AirspaceVolumePropertyType) GongNewInstance() any {
	return new(AirspaceVolumePropertyType)
}

func (*AirspaceVolumeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AirspaceVolumeTypes, stage.AirspaceVolumeType_stagedOrder)
}

func (*AirspaceVolumeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AirspaceVolumeType_orderStaged[order]
}

func (*AirspaceVolumeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AirspaceVolumeTypes_mapString
}

func (*AirspaceVolumeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AirspaceVolumeTypes
}

func (*AirspaceVolumeType) GongNewInstance() any {
	return new(AirspaceVolumeType)
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder)
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AuthorityForAirspaceTimeSlicePropertyType_orderStaged[order]
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AuthorityForAirspaceTimeSlicePropertyTypes
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongNewInstance() any {
	return new(AuthorityForAirspaceTimeSlicePropertyType)
}

func (*AuthorityForAirspaceTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceType_stagedOrder)
}

func (*AuthorityForAirspaceTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AuthorityForAirspaceTimeSliceType_orderStaged[order]
}

func (*AuthorityForAirspaceTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AuthorityForAirspaceTimeSliceTypes_mapString
}

func (*AuthorityForAirspaceTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AuthorityForAirspaceTimeSliceTypes
}

func (*AuthorityForAirspaceTimeSliceType) GongNewInstance() any {
	return new(AuthorityForAirspaceTimeSliceType)
}

func (*AuthorityForAirspaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceType_stagedOrder)
}

func (*AuthorityForAirspaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AuthorityForAirspaceType_orderStaged[order]
}

func (*AuthorityForAirspaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AuthorityForAirspaceTypes_mapString
}

func (*AuthorityForAirspaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AuthorityForAirspaceTypes
}

func (*AuthorityForAirspaceType) GongNewInstance() any {
	return new(AuthorityForAirspaceType)
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

func (*GeoBorderTimeSlicePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyType_stagedOrder)
}

func (*GeoBorderTimeSlicePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GeoBorderTimeSlicePropertyType_orderStaged[order]
}

func (*GeoBorderTimeSlicePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GeoBorderTimeSlicePropertyTypes_mapString
}

func (*GeoBorderTimeSlicePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GeoBorderTimeSlicePropertyTypes
}

func (*GeoBorderTimeSlicePropertyType) GongNewInstance() any {
	return new(GeoBorderTimeSlicePropertyType)
}

func (*GeoBorderTimeSliceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceType_stagedOrder)
}

func (*GeoBorderTimeSliceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GeoBorderTimeSliceType_orderStaged[order]
}

func (*GeoBorderTimeSliceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GeoBorderTimeSliceTypes_mapString
}

func (*GeoBorderTimeSliceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GeoBorderTimeSliceTypes
}

func (*GeoBorderTimeSliceType) GongNewInstance() any {
	return new(GeoBorderTimeSliceType)
}

func (*GeoBorderType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.GeoBorderTypes, stage.GeoBorderType_stagedOrder)
}

func (*GeoBorderType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.GeoBorderType_orderStaged[order]
}

func (*GeoBorderType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.GeoBorderTypes_mapString
}

func (*GeoBorderType) GongGetInstancesSet(stage *Stage) any {
	return &stage.GeoBorderTypes
}

func (*GeoBorderType) GongNewInstance() any {
	return new(GeoBorderType)
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

func (*RoutePropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.RoutePropertyTypes, stage.RoutePropertyType_stagedOrder)
}

func (*RoutePropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.RoutePropertyType_orderStaged[order]
}

func (*RoutePropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.RoutePropertyTypes_mapString
}

func (*RoutePropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.RoutePropertyTypes
}

func (*RoutePropertyType) GongNewInstance() any {
	return new(RoutePropertyType)
}

func (*StandardLevelColumnPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyType_stagedOrder)
}

func (*StandardLevelColumnPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StandardLevelColumnPropertyType_orderStaged[order]
}

func (*StandardLevelColumnPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StandardLevelColumnPropertyTypes_mapString
}

func (*StandardLevelColumnPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.StandardLevelColumnPropertyTypes
}

func (*StandardLevelColumnPropertyType) GongNewInstance() any {
	return new(StandardLevelColumnPropertyType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
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

		A_extension_13s:           make(map[*A_extension_13]struct{}),
		A_extension_13s_mapString: make(map[string]*A_extension_13),

		A_extension_14s:           make(map[*A_extension_14]struct{}),
		A_extension_14s_mapString: make(map[string]*A_extension_14),

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

		AircraftCharacteristicPropertyTypes:           make(map[*AircraftCharacteristicPropertyType]struct{}),
		AircraftCharacteristicPropertyTypes_mapString: make(map[string]*AircraftCharacteristicPropertyType),

		AircraftCharacteristicTypes:           make(map[*AircraftCharacteristicType]struct{}),
		AircraftCharacteristicTypes_mapString: make(map[string]*AircraftCharacteristicType),

		AirspaceActivationPropertyTypes:           make(map[*AirspaceActivationPropertyType]struct{}),
		AirspaceActivationPropertyTypes_mapString: make(map[string]*AirspaceActivationPropertyType),

		AirspaceActivationTypes:           make(map[*AirspaceActivationType]struct{}),
		AirspaceActivationTypes_mapString: make(map[string]*AirspaceActivationType),

		AirspaceGeometryComponentPropertyTypes:           make(map[*AirspaceGeometryComponentPropertyType]struct{}),
		AirspaceGeometryComponentPropertyTypes_mapString: make(map[string]*AirspaceGeometryComponentPropertyType),

		AirspaceGeometryComponentTypes:           make(map[*AirspaceGeometryComponentType]struct{}),
		AirspaceGeometryComponentTypes_mapString: make(map[string]*AirspaceGeometryComponentType),

		AirspaceLayerClassPropertyTypes:           make(map[*AirspaceLayerClassPropertyType]struct{}),
		AirspaceLayerClassPropertyTypes_mapString: make(map[string]*AirspaceLayerClassPropertyType),

		AirspaceLayerClassTypes:           make(map[*AirspaceLayerClassType]struct{}),
		AirspaceLayerClassTypes_mapString: make(map[string]*AirspaceLayerClassType),

		AirspaceLayerPropertyTypes:           make(map[*AirspaceLayerPropertyType]struct{}),
		AirspaceLayerPropertyTypes_mapString: make(map[string]*AirspaceLayerPropertyType),

		AirspaceLayerTypes:           make(map[*AirspaceLayerType]struct{}),
		AirspaceLayerTypes_mapString: make(map[string]*AirspaceLayerType),

		AirspacePropertyTypes:           make(map[*AirspacePropertyType]struct{}),
		AirspacePropertyTypes_mapString: make(map[string]*AirspacePropertyType),

		AirspaceTimeSlicePropertyTypes:           make(map[*AirspaceTimeSlicePropertyType]struct{}),
		AirspaceTimeSlicePropertyTypes_mapString: make(map[string]*AirspaceTimeSlicePropertyType),

		AirspaceTimeSliceTypes:           make(map[*AirspaceTimeSliceType]struct{}),
		AirspaceTimeSliceTypes_mapString: make(map[string]*AirspaceTimeSliceType),

		AirspaceTypes:           make(map[*AirspaceType]struct{}),
		AirspaceTypes_mapString: make(map[string]*AirspaceType),

		AirspaceVolumeDependencyPropertyTypes:           make(map[*AirspaceVolumeDependencyPropertyType]struct{}),
		AirspaceVolumeDependencyPropertyTypes_mapString: make(map[string]*AirspaceVolumeDependencyPropertyType),

		AirspaceVolumeDependencyTypes:           make(map[*AirspaceVolumeDependencyType]struct{}),
		AirspaceVolumeDependencyTypes_mapString: make(map[string]*AirspaceVolumeDependencyType),

		AirspaceVolumePropertyTypes:           make(map[*AirspaceVolumePropertyType]struct{}),
		AirspaceVolumePropertyTypes_mapString: make(map[string]*AirspaceVolumePropertyType),

		AirspaceVolumeTypes:           make(map[*AirspaceVolumeType]struct{}),
		AirspaceVolumeTypes_mapString: make(map[string]*AirspaceVolumeType),

		AuthorityForAirspaceTimeSlicePropertyTypes:           make(map[*AuthorityForAirspaceTimeSlicePropertyType]struct{}),
		AuthorityForAirspaceTimeSlicePropertyTypes_mapString: make(map[string]*AuthorityForAirspaceTimeSlicePropertyType),

		AuthorityForAirspaceTimeSliceTypes:           make(map[*AuthorityForAirspaceTimeSliceType]struct{}),
		AuthorityForAirspaceTimeSliceTypes_mapString: make(map[string]*AuthorityForAirspaceTimeSliceType),

		AuthorityForAirspaceTypes:           make(map[*AuthorityForAirspaceType]struct{}),
		AuthorityForAirspaceTypes_mapString: make(map[string]*AuthorityForAirspaceType),

		CurvePropertyTypes:           make(map[*CurvePropertyType]struct{}),
		CurvePropertyTypes_mapString: make(map[string]*CurvePropertyType),

		CurveTypes:           make(map[*CurveType]struct{}),
		CurveTypes_mapString: make(map[string]*CurveType),

		ElevatedCurveTypes:           make(map[*ElevatedCurveType]struct{}),
		ElevatedCurveTypes_mapString: make(map[string]*ElevatedCurveType),

		ElevatedSurfaceTypes:           make(map[*ElevatedSurfaceType]struct{}),
		ElevatedSurfaceTypes_mapString: make(map[string]*ElevatedSurfaceType),

		GeoBorderTimeSlicePropertyTypes:           make(map[*GeoBorderTimeSlicePropertyType]struct{}),
		GeoBorderTimeSlicePropertyTypes_mapString: make(map[string]*GeoBorderTimeSlicePropertyType),

		GeoBorderTimeSliceTypes:           make(map[*GeoBorderTimeSliceType]struct{}),
		GeoBorderTimeSliceTypes_mapString: make(map[string]*GeoBorderTimeSliceType),

		GeoBorderTypes:           make(map[*GeoBorderType]struct{}),
		GeoBorderTypes_mapString: make(map[string]*GeoBorderType),

		LinguisticNotePropertyTypes:           make(map[*LinguisticNotePropertyType]struct{}),
		LinguisticNotePropertyTypes_mapString: make(map[string]*LinguisticNotePropertyType),

		LinguisticNoteTypes:           make(map[*LinguisticNoteType]struct{}),
		LinguisticNoteTypes_mapString: make(map[string]*LinguisticNoteType),

		NotePropertyTypes:           make(map[*NotePropertyType]struct{}),
		NotePropertyTypes_mapString: make(map[string]*NotePropertyType),

		NoteTypes:           make(map[*NoteType]struct{}),
		NoteTypes_mapString: make(map[string]*NoteType),

		OrganisationAuthorityPropertyTypes:           make(map[*OrganisationAuthorityPropertyType]struct{}),
		OrganisationAuthorityPropertyTypes_mapString: make(map[string]*OrganisationAuthorityPropertyType),

		RoutePropertyTypes:           make(map[*RoutePropertyType]struct{}),
		RoutePropertyTypes_mapString: make(map[string]*RoutePropertyType),

		StandardLevelColumnPropertyTypes:           make(map[*StandardLevelColumnPropertyType]struct{}),
		StandardLevelColumnPropertyTypes_mapString: make(map[string]*StandardLevelColumnPropertyType),

		SurfacePropertyTypes:           make(map[*SurfacePropertyType]struct{}),
		SurfacePropertyTypes_mapString: make(map[string]*SurfacePropertyType),

		SurfaceTypes:           make(map[*SurfaceType]struct{}),
		SurfaceTypes_mapString: make(map[string]*SurfaceType),

		TimesheetPropertyTypes:           make(map[*TimesheetPropertyType]struct{}),
		TimesheetPropertyTypes_mapString: make(map[string]*TimesheetPropertyType),

		TimesheetTypes:           make(map[*TimesheetType]struct{}),
		TimesheetTypes_mapString: make(map[string]*TimesheetType),

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

		A_extension_13_stagedOrder: make(map[*A_extension_13]uint),
		A_extension_13_orderStaged: make(map[uint]*A_extension_13),
		A_extension_13s_reference:  make(map[*A_extension_13]*A_extension_13),

		A_extension_14_stagedOrder: make(map[*A_extension_14]uint),
		A_extension_14_orderStaged: make(map[uint]*A_extension_14),
		A_extension_14s_reference:  make(map[*A_extension_14]*A_extension_14),

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

		AircraftCharacteristicPropertyType_stagedOrder: make(map[*AircraftCharacteristicPropertyType]uint),
		AircraftCharacteristicPropertyType_orderStaged: make(map[uint]*AircraftCharacteristicPropertyType),
		AircraftCharacteristicPropertyTypes_reference:  make(map[*AircraftCharacteristicPropertyType]*AircraftCharacteristicPropertyType),

		AircraftCharacteristicType_stagedOrder: make(map[*AircraftCharacteristicType]uint),
		AircraftCharacteristicType_orderStaged: make(map[uint]*AircraftCharacteristicType),
		AircraftCharacteristicTypes_reference:  make(map[*AircraftCharacteristicType]*AircraftCharacteristicType),

		AirspaceActivationPropertyType_stagedOrder: make(map[*AirspaceActivationPropertyType]uint),
		AirspaceActivationPropertyType_orderStaged: make(map[uint]*AirspaceActivationPropertyType),
		AirspaceActivationPropertyTypes_reference:  make(map[*AirspaceActivationPropertyType]*AirspaceActivationPropertyType),

		AirspaceActivationType_stagedOrder: make(map[*AirspaceActivationType]uint),
		AirspaceActivationType_orderStaged: make(map[uint]*AirspaceActivationType),
		AirspaceActivationTypes_reference:  make(map[*AirspaceActivationType]*AirspaceActivationType),

		AirspaceGeometryComponentPropertyType_stagedOrder: make(map[*AirspaceGeometryComponentPropertyType]uint),
		AirspaceGeometryComponentPropertyType_orderStaged: make(map[uint]*AirspaceGeometryComponentPropertyType),
		AirspaceGeometryComponentPropertyTypes_reference:  make(map[*AirspaceGeometryComponentPropertyType]*AirspaceGeometryComponentPropertyType),

		AirspaceGeometryComponentType_stagedOrder: make(map[*AirspaceGeometryComponentType]uint),
		AirspaceGeometryComponentType_orderStaged: make(map[uint]*AirspaceGeometryComponentType),
		AirspaceGeometryComponentTypes_reference:  make(map[*AirspaceGeometryComponentType]*AirspaceGeometryComponentType),

		AirspaceLayerClassPropertyType_stagedOrder: make(map[*AirspaceLayerClassPropertyType]uint),
		AirspaceLayerClassPropertyType_orderStaged: make(map[uint]*AirspaceLayerClassPropertyType),
		AirspaceLayerClassPropertyTypes_reference:  make(map[*AirspaceLayerClassPropertyType]*AirspaceLayerClassPropertyType),

		AirspaceLayerClassType_stagedOrder: make(map[*AirspaceLayerClassType]uint),
		AirspaceLayerClassType_orderStaged: make(map[uint]*AirspaceLayerClassType),
		AirspaceLayerClassTypes_reference:  make(map[*AirspaceLayerClassType]*AirspaceLayerClassType),

		AirspaceLayerPropertyType_stagedOrder: make(map[*AirspaceLayerPropertyType]uint),
		AirspaceLayerPropertyType_orderStaged: make(map[uint]*AirspaceLayerPropertyType),
		AirspaceLayerPropertyTypes_reference:  make(map[*AirspaceLayerPropertyType]*AirspaceLayerPropertyType),

		AirspaceLayerType_stagedOrder: make(map[*AirspaceLayerType]uint),
		AirspaceLayerType_orderStaged: make(map[uint]*AirspaceLayerType),
		AirspaceLayerTypes_reference:  make(map[*AirspaceLayerType]*AirspaceLayerType),

		AirspacePropertyType_stagedOrder: make(map[*AirspacePropertyType]uint),
		AirspacePropertyType_orderStaged: make(map[uint]*AirspacePropertyType),
		AirspacePropertyTypes_reference:  make(map[*AirspacePropertyType]*AirspacePropertyType),

		AirspaceTimeSlicePropertyType_stagedOrder: make(map[*AirspaceTimeSlicePropertyType]uint),
		AirspaceTimeSlicePropertyType_orderStaged: make(map[uint]*AirspaceTimeSlicePropertyType),
		AirspaceTimeSlicePropertyTypes_reference:  make(map[*AirspaceTimeSlicePropertyType]*AirspaceTimeSlicePropertyType),

		AirspaceTimeSliceType_stagedOrder: make(map[*AirspaceTimeSliceType]uint),
		AirspaceTimeSliceType_orderStaged: make(map[uint]*AirspaceTimeSliceType),
		AirspaceTimeSliceTypes_reference:  make(map[*AirspaceTimeSliceType]*AirspaceTimeSliceType),

		AirspaceType_stagedOrder: make(map[*AirspaceType]uint),
		AirspaceType_orderStaged: make(map[uint]*AirspaceType),
		AirspaceTypes_reference:  make(map[*AirspaceType]*AirspaceType),

		AirspaceVolumeDependencyPropertyType_stagedOrder: make(map[*AirspaceVolumeDependencyPropertyType]uint),
		AirspaceVolumeDependencyPropertyType_orderStaged: make(map[uint]*AirspaceVolumeDependencyPropertyType),
		AirspaceVolumeDependencyPropertyTypes_reference:  make(map[*AirspaceVolumeDependencyPropertyType]*AirspaceVolumeDependencyPropertyType),

		AirspaceVolumeDependencyType_stagedOrder: make(map[*AirspaceVolumeDependencyType]uint),
		AirspaceVolumeDependencyType_orderStaged: make(map[uint]*AirspaceVolumeDependencyType),
		AirspaceVolumeDependencyTypes_reference:  make(map[*AirspaceVolumeDependencyType]*AirspaceVolumeDependencyType),

		AirspaceVolumePropertyType_stagedOrder: make(map[*AirspaceVolumePropertyType]uint),
		AirspaceVolumePropertyType_orderStaged: make(map[uint]*AirspaceVolumePropertyType),
		AirspaceVolumePropertyTypes_reference:  make(map[*AirspaceVolumePropertyType]*AirspaceVolumePropertyType),

		AirspaceVolumeType_stagedOrder: make(map[*AirspaceVolumeType]uint),
		AirspaceVolumeType_orderStaged: make(map[uint]*AirspaceVolumeType),
		AirspaceVolumeTypes_reference:  make(map[*AirspaceVolumeType]*AirspaceVolumeType),

		AuthorityForAirspaceTimeSlicePropertyType_stagedOrder: make(map[*AuthorityForAirspaceTimeSlicePropertyType]uint),
		AuthorityForAirspaceTimeSlicePropertyType_orderStaged: make(map[uint]*AuthorityForAirspaceTimeSlicePropertyType),
		AuthorityForAirspaceTimeSlicePropertyTypes_reference:  make(map[*AuthorityForAirspaceTimeSlicePropertyType]*AuthorityForAirspaceTimeSlicePropertyType),

		AuthorityForAirspaceTimeSliceType_stagedOrder: make(map[*AuthorityForAirspaceTimeSliceType]uint),
		AuthorityForAirspaceTimeSliceType_orderStaged: make(map[uint]*AuthorityForAirspaceTimeSliceType),
		AuthorityForAirspaceTimeSliceTypes_reference:  make(map[*AuthorityForAirspaceTimeSliceType]*AuthorityForAirspaceTimeSliceType),

		AuthorityForAirspaceType_stagedOrder: make(map[*AuthorityForAirspaceType]uint),
		AuthorityForAirspaceType_orderStaged: make(map[uint]*AuthorityForAirspaceType),
		AuthorityForAirspaceTypes_reference:  make(map[*AuthorityForAirspaceType]*AuthorityForAirspaceType),

		CurvePropertyType_stagedOrder: make(map[*CurvePropertyType]uint),
		CurvePropertyType_orderStaged: make(map[uint]*CurvePropertyType),
		CurvePropertyTypes_reference:  make(map[*CurvePropertyType]*CurvePropertyType),

		CurveType_stagedOrder: make(map[*CurveType]uint),
		CurveType_orderStaged: make(map[uint]*CurveType),
		CurveTypes_reference:  make(map[*CurveType]*CurveType),

		ElevatedCurveType_stagedOrder: make(map[*ElevatedCurveType]uint),
		ElevatedCurveType_orderStaged: make(map[uint]*ElevatedCurveType),
		ElevatedCurveTypes_reference:  make(map[*ElevatedCurveType]*ElevatedCurveType),

		ElevatedSurfaceType_stagedOrder: make(map[*ElevatedSurfaceType]uint),
		ElevatedSurfaceType_orderStaged: make(map[uint]*ElevatedSurfaceType),
		ElevatedSurfaceTypes_reference:  make(map[*ElevatedSurfaceType]*ElevatedSurfaceType),

		GeoBorderTimeSlicePropertyType_stagedOrder: make(map[*GeoBorderTimeSlicePropertyType]uint),
		GeoBorderTimeSlicePropertyType_orderStaged: make(map[uint]*GeoBorderTimeSlicePropertyType),
		GeoBorderTimeSlicePropertyTypes_reference:  make(map[*GeoBorderTimeSlicePropertyType]*GeoBorderTimeSlicePropertyType),

		GeoBorderTimeSliceType_stagedOrder: make(map[*GeoBorderTimeSliceType]uint),
		GeoBorderTimeSliceType_orderStaged: make(map[uint]*GeoBorderTimeSliceType),
		GeoBorderTimeSliceTypes_reference:  make(map[*GeoBorderTimeSliceType]*GeoBorderTimeSliceType),

		GeoBorderType_stagedOrder: make(map[*GeoBorderType]uint),
		GeoBorderType_orderStaged: make(map[uint]*GeoBorderType),
		GeoBorderTypes_reference:  make(map[*GeoBorderType]*GeoBorderType),

		LinguisticNotePropertyType_stagedOrder: make(map[*LinguisticNotePropertyType]uint),
		LinguisticNotePropertyType_orderStaged: make(map[uint]*LinguisticNotePropertyType),
		LinguisticNotePropertyTypes_reference:  make(map[*LinguisticNotePropertyType]*LinguisticNotePropertyType),

		LinguisticNoteType_stagedOrder: make(map[*LinguisticNoteType]uint),
		LinguisticNoteType_orderStaged: make(map[uint]*LinguisticNoteType),
		LinguisticNoteTypes_reference:  make(map[*LinguisticNoteType]*LinguisticNoteType),

		NotePropertyType_stagedOrder: make(map[*NotePropertyType]uint),
		NotePropertyType_orderStaged: make(map[uint]*NotePropertyType),
		NotePropertyTypes_reference:  make(map[*NotePropertyType]*NotePropertyType),

		NoteType_stagedOrder: make(map[*NoteType]uint),
		NoteType_orderStaged: make(map[uint]*NoteType),
		NoteTypes_reference:  make(map[*NoteType]*NoteType),

		OrganisationAuthorityPropertyType_stagedOrder: make(map[*OrganisationAuthorityPropertyType]uint),
		OrganisationAuthorityPropertyType_orderStaged: make(map[uint]*OrganisationAuthorityPropertyType),
		OrganisationAuthorityPropertyTypes_reference:  make(map[*OrganisationAuthorityPropertyType]*OrganisationAuthorityPropertyType),

		RoutePropertyType_stagedOrder: make(map[*RoutePropertyType]uint),
		RoutePropertyType_orderStaged: make(map[uint]*RoutePropertyType),
		RoutePropertyTypes_reference:  make(map[*RoutePropertyType]*RoutePropertyType),

		StandardLevelColumnPropertyType_stagedOrder: make(map[*StandardLevelColumnPropertyType]uint),
		StandardLevelColumnPropertyType_orderStaged: make(map[uint]*StandardLevelColumnPropertyType),
		StandardLevelColumnPropertyTypes_reference:  make(map[*StandardLevelColumnPropertyType]*StandardLevelColumnPropertyType),

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

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"A_extension": &A_extensionUnmarshaller{},

			"A_extension_1": &A_extension_1Unmarshaller{},

			"A_extension_10": &A_extension_10Unmarshaller{},

			"A_extension_11": &A_extension_11Unmarshaller{},

			"A_extension_12": &A_extension_12Unmarshaller{},

			"A_extension_13": &A_extension_13Unmarshaller{},

			"A_extension_14": &A_extension_14Unmarshaller{},

			"A_extension_2": &A_extension_2Unmarshaller{},

			"A_extension_3": &A_extension_3Unmarshaller{},

			"A_extension_4": &A_extension_4Unmarshaller{},

			"A_extension_5": &A_extension_5Unmarshaller{},

			"A_extension_6": &A_extension_6Unmarshaller{},

			"A_extension_7": &A_extension_7Unmarshaller{},

			"A_extension_8": &A_extension_8Unmarshaller{},

			"A_extension_9": &A_extension_9Unmarshaller{},

			"AbstractPropertiesWithScheduleType": &AbstractPropertiesWithScheduleTypeUnmarshaller{},

			"AircraftCharacteristicPropertyType": &AircraftCharacteristicPropertyTypeUnmarshaller{},

			"AircraftCharacteristicType": &AircraftCharacteristicTypeUnmarshaller{},

			"AirspaceActivationPropertyType": &AirspaceActivationPropertyTypeUnmarshaller{},

			"AirspaceActivationType": &AirspaceActivationTypeUnmarshaller{},

			"AirspaceGeometryComponentPropertyType": &AirspaceGeometryComponentPropertyTypeUnmarshaller{},

			"AirspaceGeometryComponentType": &AirspaceGeometryComponentTypeUnmarshaller{},

			"AirspaceLayerClassPropertyType": &AirspaceLayerClassPropertyTypeUnmarshaller{},

			"AirspaceLayerClassType": &AirspaceLayerClassTypeUnmarshaller{},

			"AirspaceLayerPropertyType": &AirspaceLayerPropertyTypeUnmarshaller{},

			"AirspaceLayerType": &AirspaceLayerTypeUnmarshaller{},

			"AirspacePropertyType": &AirspacePropertyTypeUnmarshaller{},

			"AirspaceTimeSlicePropertyType": &AirspaceTimeSlicePropertyTypeUnmarshaller{},

			"AirspaceTimeSliceType": &AirspaceTimeSliceTypeUnmarshaller{},

			"AirspaceType": &AirspaceTypeUnmarshaller{},

			"AirspaceVolumeDependencyPropertyType": &AirspaceVolumeDependencyPropertyTypeUnmarshaller{},

			"AirspaceVolumeDependencyType": &AirspaceVolumeDependencyTypeUnmarshaller{},

			"AirspaceVolumePropertyType": &AirspaceVolumePropertyTypeUnmarshaller{},

			"AirspaceVolumeType": &AirspaceVolumeTypeUnmarshaller{},

			"AuthorityForAirspaceTimeSlicePropertyType": &AuthorityForAirspaceTimeSlicePropertyTypeUnmarshaller{},

			"AuthorityForAirspaceTimeSliceType": &AuthorityForAirspaceTimeSliceTypeUnmarshaller{},

			"AuthorityForAirspaceType": &AuthorityForAirspaceTypeUnmarshaller{},

			"CurvePropertyType": &CurvePropertyTypeUnmarshaller{},

			"CurveType": &CurveTypeUnmarshaller{},

			"ElevatedCurveType": &ElevatedCurveTypeUnmarshaller{},

			"ElevatedSurfaceType": &ElevatedSurfaceTypeUnmarshaller{},

			"GeoBorderTimeSlicePropertyType": &GeoBorderTimeSlicePropertyTypeUnmarshaller{},

			"GeoBorderTimeSliceType": &GeoBorderTimeSliceTypeUnmarshaller{},

			"GeoBorderType": &GeoBorderTypeUnmarshaller{},

			"LinguisticNotePropertyType": &LinguisticNotePropertyTypeUnmarshaller{},

			"LinguisticNoteType": &LinguisticNoteTypeUnmarshaller{},

			"NotePropertyType": &NotePropertyTypeUnmarshaller{},

			"NoteType": &NoteTypeUnmarshaller{},

			"OrganisationAuthorityPropertyType": &OrganisationAuthorityPropertyTypeUnmarshaller{},

			"RoutePropertyType": &RoutePropertyTypeUnmarshaller{},

			"StandardLevelColumnPropertyType": &StandardLevelColumnPropertyTypeUnmarshaller{},

			"SurfacePropertyType": &SurfacePropertyTypeUnmarshaller{},

			"SurfaceType": &SurfaceTypeUnmarshaller{},

			"TimesheetPropertyType": &TimesheetPropertyTypeUnmarshaller{},

			"TimesheetType": &TimesheetTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["A_extension_13"] = len(stage.A_extension_13s)
	stage.Map_GongStructName_InstancesNb["A_extension_14"] = len(stage.A_extension_14s)
	stage.Map_GongStructName_InstancesNb["A_extension_2"] = len(stage.A_extension_2s)
	stage.Map_GongStructName_InstancesNb["A_extension_3"] = len(stage.A_extension_3s)
	stage.Map_GongStructName_InstancesNb["A_extension_4"] = len(stage.A_extension_4s)
	stage.Map_GongStructName_InstancesNb["A_extension_5"] = len(stage.A_extension_5s)
	stage.Map_GongStructName_InstancesNb["A_extension_6"] = len(stage.A_extension_6s)
	stage.Map_GongStructName_InstancesNb["A_extension_7"] = len(stage.A_extension_7s)
	stage.Map_GongStructName_InstancesNb["A_extension_8"] = len(stage.A_extension_8s)
	stage.Map_GongStructName_InstancesNb["A_extension_9"] = len(stage.A_extension_9s)
	stage.Map_GongStructName_InstancesNb["AbstractPropertiesWithScheduleType"] = len(stage.AbstractPropertiesWithScheduleTypes)
	stage.Map_GongStructName_InstancesNb["AircraftCharacteristicPropertyType"] = len(stage.AircraftCharacteristicPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AircraftCharacteristicType"] = len(stage.AircraftCharacteristicTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceActivationPropertyType"] = len(stage.AirspaceActivationPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceActivationType"] = len(stage.AirspaceActivationTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceGeometryComponentPropertyType"] = len(stage.AirspaceGeometryComponentPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceGeometryComponentType"] = len(stage.AirspaceGeometryComponentTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceLayerClassPropertyType"] = len(stage.AirspaceLayerClassPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceLayerClassType"] = len(stage.AirspaceLayerClassTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceLayerPropertyType"] = len(stage.AirspaceLayerPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceLayerType"] = len(stage.AirspaceLayerTypes)
	stage.Map_GongStructName_InstancesNb["AirspacePropertyType"] = len(stage.AirspacePropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceTimeSlicePropertyType"] = len(stage.AirspaceTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceTimeSliceType"] = len(stage.AirspaceTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceType"] = len(stage.AirspaceTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceVolumeDependencyPropertyType"] = len(stage.AirspaceVolumeDependencyPropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceVolumeDependencyType"] = len(stage.AirspaceVolumeDependencyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceVolumePropertyType"] = len(stage.AirspaceVolumePropertyTypes)
	stage.Map_GongStructName_InstancesNb["AirspaceVolumeType"] = len(stage.AirspaceVolumeTypes)
	stage.Map_GongStructName_InstancesNb["AuthorityForAirspaceTimeSlicePropertyType"] = len(stage.AuthorityForAirspaceTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["AuthorityForAirspaceTimeSliceType"] = len(stage.AuthorityForAirspaceTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["AuthorityForAirspaceType"] = len(stage.AuthorityForAirspaceTypes)
	stage.Map_GongStructName_InstancesNb["CurvePropertyType"] = len(stage.CurvePropertyTypes)
	stage.Map_GongStructName_InstancesNb["CurveType"] = len(stage.CurveTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedCurveType"] = len(stage.ElevatedCurveTypes)
	stage.Map_GongStructName_InstancesNb["ElevatedSurfaceType"] = len(stage.ElevatedSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["GeoBorderTimeSlicePropertyType"] = len(stage.GeoBorderTimeSlicePropertyTypes)
	stage.Map_GongStructName_InstancesNb["GeoBorderTimeSliceType"] = len(stage.GeoBorderTimeSliceTypes)
	stage.Map_GongStructName_InstancesNb["GeoBorderType"] = len(stage.GeoBorderTypes)
	stage.Map_GongStructName_InstancesNb["LinguisticNotePropertyType"] = len(stage.LinguisticNotePropertyTypes)
	stage.Map_GongStructName_InstancesNb["LinguisticNoteType"] = len(stage.LinguisticNoteTypes)
	stage.Map_GongStructName_InstancesNb["NotePropertyType"] = len(stage.NotePropertyTypes)
	stage.Map_GongStructName_InstancesNb["NoteType"] = len(stage.NoteTypes)
	stage.Map_GongStructName_InstancesNb["OrganisationAuthorityPropertyType"] = len(stage.OrganisationAuthorityPropertyTypes)
	stage.Map_GongStructName_InstancesNb["RoutePropertyType"] = len(stage.RoutePropertyTypes)
	stage.Map_GongStructName_InstancesNb["StandardLevelColumnPropertyType"] = len(stage.StandardLevelColumnPropertyTypes)
	stage.Map_GongStructName_InstancesNb["SurfacePropertyType"] = len(stage.SurfacePropertyTypes)
	stage.Map_GongStructName_InstancesNb["SurfaceType"] = len(stage.SurfaceTypes)
	stage.Map_GongStructName_InstancesNb["TimesheetPropertyType"] = len(stage.TimesheetPropertyTypes)
	stage.Map_GongStructName_InstancesNb["TimesheetType"] = len(stage.TimesheetTypes)
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

// Stage puts a_extension_13 to the model stage
func (a_extension_13 *A_extension_13) Stage(stage *Stage) *A_extension_13 {
	__gong__stage(stage.A_extension_13s, stage.A_extension_13_stagedOrder, stage.A_extension_13_orderStaged, &stage.A_extension_13Order, stage.A_extension_13s_mapString, a_extension_13, a_extension_13.Name)
	return a_extension_13
}

// StagePreserveOrder puts a_extension_13 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_13Order
// - update stage.A_extension_13Order accordingly
func (a_extension_13 *A_extension_13) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_13s, stage.A_extension_13_stagedOrder, stage.A_extension_13_orderStaged, &stage.A_extension_13Order, stage.A_extension_13s_mapString, a_extension_13, order, a_extension_13.Name)
}

// Unstage removes a_extension_13 off the model stage
func (a_extension_13 *A_extension_13) Unstage(stage *Stage) *A_extension_13 {
	__gong__unstage(stage.A_extension_13s, stage.A_extension_13s_mapString, a_extension_13, a_extension_13.Name)
	return a_extension_13
}

// UnstageVoid removes a_extension_13 off the model stage
func (a_extension_13 *A_extension_13) UnstageVoid(stage *Stage) {
	a_extension_13.Unstage(stage)
}

func (a_extension_13 *A_extension_13) StageVoid(stage *Stage) {
	a_extension_13.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_13 *A_extension_13) GetName() (res string) {
	return a_extension_13.Name
}

// for satisfaction of GongStruct interface
func (a_extension_13 *A_extension_13) SetName(name string) {
	a_extension_13.Name = name
}

// Stage puts a_extension_14 to the model stage
func (a_extension_14 *A_extension_14) Stage(stage *Stage) *A_extension_14 {
	__gong__stage(stage.A_extension_14s, stage.A_extension_14_stagedOrder, stage.A_extension_14_orderStaged, &stage.A_extension_14Order, stage.A_extension_14s_mapString, a_extension_14, a_extension_14.Name)
	return a_extension_14
}

// StagePreserveOrder puts a_extension_14 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.A_extension_14Order
// - update stage.A_extension_14Order accordingly
func (a_extension_14 *A_extension_14) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.A_extension_14s, stage.A_extension_14_stagedOrder, stage.A_extension_14_orderStaged, &stage.A_extension_14Order, stage.A_extension_14s_mapString, a_extension_14, order, a_extension_14.Name)
}

// Unstage removes a_extension_14 off the model stage
func (a_extension_14 *A_extension_14) Unstage(stage *Stage) *A_extension_14 {
	__gong__unstage(stage.A_extension_14s, stage.A_extension_14s_mapString, a_extension_14, a_extension_14.Name)
	return a_extension_14
}

// UnstageVoid removes a_extension_14 off the model stage
func (a_extension_14 *A_extension_14) UnstageVoid(stage *Stage) {
	a_extension_14.Unstage(stage)
}

func (a_extension_14 *A_extension_14) StageVoid(stage *Stage) {
	a_extension_14.Stage(stage)
}

// for satisfaction of GongStruct interface
func (a_extension_14 *A_extension_14) GetName() (res string) {
	return a_extension_14.Name
}

// for satisfaction of GongStruct interface
func (a_extension_14 *A_extension_14) SetName(name string) {
	a_extension_14.Name = name
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

// Stage puts aircraftcharacteristicpropertytype to the model stage
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) Stage(stage *Stage) *AircraftCharacteristicPropertyType {
	__gong__stage(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyType_stagedOrder, stage.AircraftCharacteristicPropertyType_orderStaged, &stage.AircraftCharacteristicPropertyTypeOrder, stage.AircraftCharacteristicPropertyTypes_mapString, aircraftcharacteristicpropertytype, aircraftcharacteristicpropertytype.Name)
	return aircraftcharacteristicpropertytype
}

// StagePreserveOrder puts aircraftcharacteristicpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AircraftCharacteristicPropertyTypeOrder
// - update stage.AircraftCharacteristicPropertyTypeOrder accordingly
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyType_stagedOrder, stage.AircraftCharacteristicPropertyType_orderStaged, &stage.AircraftCharacteristicPropertyTypeOrder, stage.AircraftCharacteristicPropertyTypes_mapString, aircraftcharacteristicpropertytype, order, aircraftcharacteristicpropertytype.Name)
}

// Unstage removes aircraftcharacteristicpropertytype off the model stage
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) Unstage(stage *Stage) *AircraftCharacteristicPropertyType {
	__gong__unstage(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyTypes_mapString, aircraftcharacteristicpropertytype, aircraftcharacteristicpropertytype.Name)
	return aircraftcharacteristicpropertytype
}

// UnstageVoid removes aircraftcharacteristicpropertytype off the model stage
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) UnstageVoid(stage *Stage) {
	aircraftcharacteristicpropertytype.Unstage(stage)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) StageVoid(stage *Stage) {
	aircraftcharacteristicpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GetName() (res string) {
	return aircraftcharacteristicpropertytype.Name
}

// for satisfaction of GongStruct interface
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) SetName(name string) {
	aircraftcharacteristicpropertytype.Name = name
}

// Stage puts aircraftcharacteristictype to the model stage
func (aircraftcharacteristictype *AircraftCharacteristicType) Stage(stage *Stage) *AircraftCharacteristicType {
	__gong__stage(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicType_stagedOrder, stage.AircraftCharacteristicType_orderStaged, &stage.AircraftCharacteristicTypeOrder, stage.AircraftCharacteristicTypes_mapString, aircraftcharacteristictype, aircraftcharacteristictype.Name)
	return aircraftcharacteristictype
}

// StagePreserveOrder puts aircraftcharacteristictype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AircraftCharacteristicTypeOrder
// - update stage.AircraftCharacteristicTypeOrder accordingly
func (aircraftcharacteristictype *AircraftCharacteristicType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicType_stagedOrder, stage.AircraftCharacteristicType_orderStaged, &stage.AircraftCharacteristicTypeOrder, stage.AircraftCharacteristicTypes_mapString, aircraftcharacteristictype, order, aircraftcharacteristictype.Name)
}

// Unstage removes aircraftcharacteristictype off the model stage
func (aircraftcharacteristictype *AircraftCharacteristicType) Unstage(stage *Stage) *AircraftCharacteristicType {
	__gong__unstage(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicTypes_mapString, aircraftcharacteristictype, aircraftcharacteristictype.Name)
	return aircraftcharacteristictype
}

// UnstageVoid removes aircraftcharacteristictype off the model stage
func (aircraftcharacteristictype *AircraftCharacteristicType) UnstageVoid(stage *Stage) {
	aircraftcharacteristictype.Unstage(stage)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) StageVoid(stage *Stage) {
	aircraftcharacteristictype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aircraftcharacteristictype *AircraftCharacteristicType) GetName() (res string) {
	return aircraftcharacteristictype.Name
}

// for satisfaction of GongStruct interface
func (aircraftcharacteristictype *AircraftCharacteristicType) SetName(name string) {
	aircraftcharacteristictype.Name = name
}

// Stage puts airspaceactivationpropertytype to the model stage
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) Stage(stage *Stage) *AirspaceActivationPropertyType {
	__gong__stage(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyType_stagedOrder, stage.AirspaceActivationPropertyType_orderStaged, &stage.AirspaceActivationPropertyTypeOrder, stage.AirspaceActivationPropertyTypes_mapString, airspaceactivationpropertytype, airspaceactivationpropertytype.Name)
	return airspaceactivationpropertytype
}

// StagePreserveOrder puts airspaceactivationpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceActivationPropertyTypeOrder
// - update stage.AirspaceActivationPropertyTypeOrder accordingly
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyType_stagedOrder, stage.AirspaceActivationPropertyType_orderStaged, &stage.AirspaceActivationPropertyTypeOrder, stage.AirspaceActivationPropertyTypes_mapString, airspaceactivationpropertytype, order, airspaceactivationpropertytype.Name)
}

// Unstage removes airspaceactivationpropertytype off the model stage
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) Unstage(stage *Stage) *AirspaceActivationPropertyType {
	__gong__unstage(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyTypes_mapString, airspaceactivationpropertytype, airspaceactivationpropertytype.Name)
	return airspaceactivationpropertytype
}

// UnstageVoid removes airspaceactivationpropertytype off the model stage
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) UnstageVoid(stage *Stage) {
	airspaceactivationpropertytype.Unstage(stage)
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) StageVoid(stage *Stage) {
	airspaceactivationpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GetName() (res string) {
	return airspaceactivationpropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) SetName(name string) {
	airspaceactivationpropertytype.Name = name
}

// Stage puts airspaceactivationtype to the model stage
func (airspaceactivationtype *AirspaceActivationType) Stage(stage *Stage) *AirspaceActivationType {
	__gong__stage(stage.AirspaceActivationTypes, stage.AirspaceActivationType_stagedOrder, stage.AirspaceActivationType_orderStaged, &stage.AirspaceActivationTypeOrder, stage.AirspaceActivationTypes_mapString, airspaceactivationtype, airspaceactivationtype.Name)
	return airspaceactivationtype
}

// StagePreserveOrder puts airspaceactivationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceActivationTypeOrder
// - update stage.AirspaceActivationTypeOrder accordingly
func (airspaceactivationtype *AirspaceActivationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceActivationTypes, stage.AirspaceActivationType_stagedOrder, stage.AirspaceActivationType_orderStaged, &stage.AirspaceActivationTypeOrder, stage.AirspaceActivationTypes_mapString, airspaceactivationtype, order, airspaceactivationtype.Name)
}

// Unstage removes airspaceactivationtype off the model stage
func (airspaceactivationtype *AirspaceActivationType) Unstage(stage *Stage) *AirspaceActivationType {
	__gong__unstage(stage.AirspaceActivationTypes, stage.AirspaceActivationTypes_mapString, airspaceactivationtype, airspaceactivationtype.Name)
	return airspaceactivationtype
}

// UnstageVoid removes airspaceactivationtype off the model stage
func (airspaceactivationtype *AirspaceActivationType) UnstageVoid(stage *Stage) {
	airspaceactivationtype.Unstage(stage)
}

func (airspaceactivationtype *AirspaceActivationType) StageVoid(stage *Stage) {
	airspaceactivationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspaceactivationtype *AirspaceActivationType) GetName() (res string) {
	return airspaceactivationtype.Name
}

// for satisfaction of GongStruct interface
func (airspaceactivationtype *AirspaceActivationType) SetName(name string) {
	airspaceactivationtype.Name = name
}

// Stage puts airspacegeometrycomponentpropertytype to the model stage
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) Stage(stage *Stage) *AirspaceGeometryComponentPropertyType {
	__gong__stage(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyType_stagedOrder, stage.AirspaceGeometryComponentPropertyType_orderStaged, &stage.AirspaceGeometryComponentPropertyTypeOrder, stage.AirspaceGeometryComponentPropertyTypes_mapString, airspacegeometrycomponentpropertytype, airspacegeometrycomponentpropertytype.Name)
	return airspacegeometrycomponentpropertytype
}

// StagePreserveOrder puts airspacegeometrycomponentpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceGeometryComponentPropertyTypeOrder
// - update stage.AirspaceGeometryComponentPropertyTypeOrder accordingly
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyType_stagedOrder, stage.AirspaceGeometryComponentPropertyType_orderStaged, &stage.AirspaceGeometryComponentPropertyTypeOrder, stage.AirspaceGeometryComponentPropertyTypes_mapString, airspacegeometrycomponentpropertytype, order, airspacegeometrycomponentpropertytype.Name)
}

// Unstage removes airspacegeometrycomponentpropertytype off the model stage
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) Unstage(stage *Stage) *AirspaceGeometryComponentPropertyType {
	__gong__unstage(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyTypes_mapString, airspacegeometrycomponentpropertytype, airspacegeometrycomponentpropertytype.Name)
	return airspacegeometrycomponentpropertytype
}

// UnstageVoid removes airspacegeometrycomponentpropertytype off the model stage
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) UnstageVoid(stage *Stage) {
	airspacegeometrycomponentpropertytype.Unstage(stage)
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) StageVoid(stage *Stage) {
	airspacegeometrycomponentpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GetName() (res string) {
	return airspacegeometrycomponentpropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) SetName(name string) {
	airspacegeometrycomponentpropertytype.Name = name
}

// Stage puts airspacegeometrycomponenttype to the model stage
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) Stage(stage *Stage) *AirspaceGeometryComponentType {
	__gong__stage(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentType_stagedOrder, stage.AirspaceGeometryComponentType_orderStaged, &stage.AirspaceGeometryComponentTypeOrder, stage.AirspaceGeometryComponentTypes_mapString, airspacegeometrycomponenttype, airspacegeometrycomponenttype.Name)
	return airspacegeometrycomponenttype
}

// StagePreserveOrder puts airspacegeometrycomponenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceGeometryComponentTypeOrder
// - update stage.AirspaceGeometryComponentTypeOrder accordingly
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentType_stagedOrder, stage.AirspaceGeometryComponentType_orderStaged, &stage.AirspaceGeometryComponentTypeOrder, stage.AirspaceGeometryComponentTypes_mapString, airspacegeometrycomponenttype, order, airspacegeometrycomponenttype.Name)
}

// Unstage removes airspacegeometrycomponenttype off the model stage
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) Unstage(stage *Stage) *AirspaceGeometryComponentType {
	__gong__unstage(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentTypes_mapString, airspacegeometrycomponenttype, airspacegeometrycomponenttype.Name)
	return airspacegeometrycomponenttype
}

// UnstageVoid removes airspacegeometrycomponenttype off the model stage
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) UnstageVoid(stage *Stage) {
	airspacegeometrycomponenttype.Unstage(stage)
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) StageVoid(stage *Stage) {
	airspacegeometrycomponenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GetName() (res string) {
	return airspacegeometrycomponenttype.Name
}

// for satisfaction of GongStruct interface
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) SetName(name string) {
	airspacegeometrycomponenttype.Name = name
}

// Stage puts airspacelayerclasspropertytype to the model stage
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) Stage(stage *Stage) *AirspaceLayerClassPropertyType {
	__gong__stage(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyType_stagedOrder, stage.AirspaceLayerClassPropertyType_orderStaged, &stage.AirspaceLayerClassPropertyTypeOrder, stage.AirspaceLayerClassPropertyTypes_mapString, airspacelayerclasspropertytype, airspacelayerclasspropertytype.Name)
	return airspacelayerclasspropertytype
}

// StagePreserveOrder puts airspacelayerclasspropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceLayerClassPropertyTypeOrder
// - update stage.AirspaceLayerClassPropertyTypeOrder accordingly
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyType_stagedOrder, stage.AirspaceLayerClassPropertyType_orderStaged, &stage.AirspaceLayerClassPropertyTypeOrder, stage.AirspaceLayerClassPropertyTypes_mapString, airspacelayerclasspropertytype, order, airspacelayerclasspropertytype.Name)
}

// Unstage removes airspacelayerclasspropertytype off the model stage
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) Unstage(stage *Stage) *AirspaceLayerClassPropertyType {
	__gong__unstage(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyTypes_mapString, airspacelayerclasspropertytype, airspacelayerclasspropertytype.Name)
	return airspacelayerclasspropertytype
}

// UnstageVoid removes airspacelayerclasspropertytype off the model stage
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) UnstageVoid(stage *Stage) {
	airspacelayerclasspropertytype.Unstage(stage)
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) StageVoid(stage *Stage) {
	airspacelayerclasspropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GetName() (res string) {
	return airspacelayerclasspropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) SetName(name string) {
	airspacelayerclasspropertytype.Name = name
}

// Stage puts airspacelayerclasstype to the model stage
func (airspacelayerclasstype *AirspaceLayerClassType) Stage(stage *Stage) *AirspaceLayerClassType {
	__gong__stage(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassType_stagedOrder, stage.AirspaceLayerClassType_orderStaged, &stage.AirspaceLayerClassTypeOrder, stage.AirspaceLayerClassTypes_mapString, airspacelayerclasstype, airspacelayerclasstype.Name)
	return airspacelayerclasstype
}

// StagePreserveOrder puts airspacelayerclasstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceLayerClassTypeOrder
// - update stage.AirspaceLayerClassTypeOrder accordingly
func (airspacelayerclasstype *AirspaceLayerClassType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassType_stagedOrder, stage.AirspaceLayerClassType_orderStaged, &stage.AirspaceLayerClassTypeOrder, stage.AirspaceLayerClassTypes_mapString, airspacelayerclasstype, order, airspacelayerclasstype.Name)
}

// Unstage removes airspacelayerclasstype off the model stage
func (airspacelayerclasstype *AirspaceLayerClassType) Unstage(stage *Stage) *AirspaceLayerClassType {
	__gong__unstage(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassTypes_mapString, airspacelayerclasstype, airspacelayerclasstype.Name)
	return airspacelayerclasstype
}

// UnstageVoid removes airspacelayerclasstype off the model stage
func (airspacelayerclasstype *AirspaceLayerClassType) UnstageVoid(stage *Stage) {
	airspacelayerclasstype.Unstage(stage)
}

func (airspacelayerclasstype *AirspaceLayerClassType) StageVoid(stage *Stage) {
	airspacelayerclasstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacelayerclasstype *AirspaceLayerClassType) GetName() (res string) {
	return airspacelayerclasstype.Name
}

// for satisfaction of GongStruct interface
func (airspacelayerclasstype *AirspaceLayerClassType) SetName(name string) {
	airspacelayerclasstype.Name = name
}

// Stage puts airspacelayerpropertytype to the model stage
func (airspacelayerpropertytype *AirspaceLayerPropertyType) Stage(stage *Stage) *AirspaceLayerPropertyType {
	__gong__stage(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyType_stagedOrder, stage.AirspaceLayerPropertyType_orderStaged, &stage.AirspaceLayerPropertyTypeOrder, stage.AirspaceLayerPropertyTypes_mapString, airspacelayerpropertytype, airspacelayerpropertytype.Name)
	return airspacelayerpropertytype
}

// StagePreserveOrder puts airspacelayerpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceLayerPropertyTypeOrder
// - update stage.AirspaceLayerPropertyTypeOrder accordingly
func (airspacelayerpropertytype *AirspaceLayerPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyType_stagedOrder, stage.AirspaceLayerPropertyType_orderStaged, &stage.AirspaceLayerPropertyTypeOrder, stage.AirspaceLayerPropertyTypes_mapString, airspacelayerpropertytype, order, airspacelayerpropertytype.Name)
}

// Unstage removes airspacelayerpropertytype off the model stage
func (airspacelayerpropertytype *AirspaceLayerPropertyType) Unstage(stage *Stage) *AirspaceLayerPropertyType {
	__gong__unstage(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyTypes_mapString, airspacelayerpropertytype, airspacelayerpropertytype.Name)
	return airspacelayerpropertytype
}

// UnstageVoid removes airspacelayerpropertytype off the model stage
func (airspacelayerpropertytype *AirspaceLayerPropertyType) UnstageVoid(stage *Stage) {
	airspacelayerpropertytype.Unstage(stage)
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) StageVoid(stage *Stage) {
	airspacelayerpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GetName() (res string) {
	return airspacelayerpropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacelayerpropertytype *AirspaceLayerPropertyType) SetName(name string) {
	airspacelayerpropertytype.Name = name
}

// Stage puts airspacelayertype to the model stage
func (airspacelayertype *AirspaceLayerType) Stage(stage *Stage) *AirspaceLayerType {
	__gong__stage(stage.AirspaceLayerTypes, stage.AirspaceLayerType_stagedOrder, stage.AirspaceLayerType_orderStaged, &stage.AirspaceLayerTypeOrder, stage.AirspaceLayerTypes_mapString, airspacelayertype, airspacelayertype.Name)
	return airspacelayertype
}

// StagePreserveOrder puts airspacelayertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceLayerTypeOrder
// - update stage.AirspaceLayerTypeOrder accordingly
func (airspacelayertype *AirspaceLayerType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceLayerTypes, stage.AirspaceLayerType_stagedOrder, stage.AirspaceLayerType_orderStaged, &stage.AirspaceLayerTypeOrder, stage.AirspaceLayerTypes_mapString, airspacelayertype, order, airspacelayertype.Name)
}

// Unstage removes airspacelayertype off the model stage
func (airspacelayertype *AirspaceLayerType) Unstage(stage *Stage) *AirspaceLayerType {
	__gong__unstage(stage.AirspaceLayerTypes, stage.AirspaceLayerTypes_mapString, airspacelayertype, airspacelayertype.Name)
	return airspacelayertype
}

// UnstageVoid removes airspacelayertype off the model stage
func (airspacelayertype *AirspaceLayerType) UnstageVoid(stage *Stage) {
	airspacelayertype.Unstage(stage)
}

func (airspacelayertype *AirspaceLayerType) StageVoid(stage *Stage) {
	airspacelayertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacelayertype *AirspaceLayerType) GetName() (res string) {
	return airspacelayertype.Name
}

// for satisfaction of GongStruct interface
func (airspacelayertype *AirspaceLayerType) SetName(name string) {
	airspacelayertype.Name = name
}

// Stage puts airspacepropertytype to the model stage
func (airspacepropertytype *AirspacePropertyType) Stage(stage *Stage) *AirspacePropertyType {
	__gong__stage(stage.AirspacePropertyTypes, stage.AirspacePropertyType_stagedOrder, stage.AirspacePropertyType_orderStaged, &stage.AirspacePropertyTypeOrder, stage.AirspacePropertyTypes_mapString, airspacepropertytype, airspacepropertytype.Name)
	return airspacepropertytype
}

// StagePreserveOrder puts airspacepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspacePropertyTypeOrder
// - update stage.AirspacePropertyTypeOrder accordingly
func (airspacepropertytype *AirspacePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspacePropertyTypes, stage.AirspacePropertyType_stagedOrder, stage.AirspacePropertyType_orderStaged, &stage.AirspacePropertyTypeOrder, stage.AirspacePropertyTypes_mapString, airspacepropertytype, order, airspacepropertytype.Name)
}

// Unstage removes airspacepropertytype off the model stage
func (airspacepropertytype *AirspacePropertyType) Unstage(stage *Stage) *AirspacePropertyType {
	__gong__unstage(stage.AirspacePropertyTypes, stage.AirspacePropertyTypes_mapString, airspacepropertytype, airspacepropertytype.Name)
	return airspacepropertytype
}

// UnstageVoid removes airspacepropertytype off the model stage
func (airspacepropertytype *AirspacePropertyType) UnstageVoid(stage *Stage) {
	airspacepropertytype.Unstage(stage)
}

func (airspacepropertytype *AirspacePropertyType) StageVoid(stage *Stage) {
	airspacepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacepropertytype *AirspacePropertyType) GetName() (res string) {
	return airspacepropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacepropertytype *AirspacePropertyType) SetName(name string) {
	airspacepropertytype.Name = name
}

// Stage puts airspacetimeslicepropertytype to the model stage
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) Stage(stage *Stage) *AirspaceTimeSlicePropertyType {
	__gong__stage(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyType_stagedOrder, stage.AirspaceTimeSlicePropertyType_orderStaged, &stage.AirspaceTimeSlicePropertyTypeOrder, stage.AirspaceTimeSlicePropertyTypes_mapString, airspacetimeslicepropertytype, airspacetimeslicepropertytype.Name)
	return airspacetimeslicepropertytype
}

// StagePreserveOrder puts airspacetimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceTimeSlicePropertyTypeOrder
// - update stage.AirspaceTimeSlicePropertyTypeOrder accordingly
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyType_stagedOrder, stage.AirspaceTimeSlicePropertyType_orderStaged, &stage.AirspaceTimeSlicePropertyTypeOrder, stage.AirspaceTimeSlicePropertyTypes_mapString, airspacetimeslicepropertytype, order, airspacetimeslicepropertytype.Name)
}

// Unstage removes airspacetimeslicepropertytype off the model stage
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) Unstage(stage *Stage) *AirspaceTimeSlicePropertyType {
	__gong__unstage(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyTypes_mapString, airspacetimeslicepropertytype, airspacetimeslicepropertytype.Name)
	return airspacetimeslicepropertytype
}

// UnstageVoid removes airspacetimeslicepropertytype off the model stage
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	airspacetimeslicepropertytype.Unstage(stage)
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) StageVoid(stage *Stage) {
	airspacetimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GetName() (res string) {
	return airspacetimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) SetName(name string) {
	airspacetimeslicepropertytype.Name = name
}

// Stage puts airspacetimeslicetype to the model stage
func (airspacetimeslicetype *AirspaceTimeSliceType) Stage(stage *Stage) *AirspaceTimeSliceType {
	__gong__stage(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceType_stagedOrder, stage.AirspaceTimeSliceType_orderStaged, &stage.AirspaceTimeSliceTypeOrder, stage.AirspaceTimeSliceTypes_mapString, airspacetimeslicetype, airspacetimeslicetype.Name)
	return airspacetimeslicetype
}

// StagePreserveOrder puts airspacetimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceTimeSliceTypeOrder
// - update stage.AirspaceTimeSliceTypeOrder accordingly
func (airspacetimeslicetype *AirspaceTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceType_stagedOrder, stage.AirspaceTimeSliceType_orderStaged, &stage.AirspaceTimeSliceTypeOrder, stage.AirspaceTimeSliceTypes_mapString, airspacetimeslicetype, order, airspacetimeslicetype.Name)
}

// Unstage removes airspacetimeslicetype off the model stage
func (airspacetimeslicetype *AirspaceTimeSliceType) Unstage(stage *Stage) *AirspaceTimeSliceType {
	__gong__unstage(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceTypes_mapString, airspacetimeslicetype, airspacetimeslicetype.Name)
	return airspacetimeslicetype
}

// UnstageVoid removes airspacetimeslicetype off the model stage
func (airspacetimeslicetype *AirspaceTimeSliceType) UnstageVoid(stage *Stage) {
	airspacetimeslicetype.Unstage(stage)
}

func (airspacetimeslicetype *AirspaceTimeSliceType) StageVoid(stage *Stage) {
	airspacetimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacetimeslicetype *AirspaceTimeSliceType) GetName() (res string) {
	return airspacetimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (airspacetimeslicetype *AirspaceTimeSliceType) SetName(name string) {
	airspacetimeslicetype.Name = name
}

// Stage puts airspacetype to the model stage
func (airspacetype *AirspaceType) Stage(stage *Stage) *AirspaceType {
	__gong__stage(stage.AirspaceTypes, stage.AirspaceType_stagedOrder, stage.AirspaceType_orderStaged, &stage.AirspaceTypeOrder, stage.AirspaceTypes_mapString, airspacetype, airspacetype.Name)
	return airspacetype
}

// StagePreserveOrder puts airspacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceTypeOrder
// - update stage.AirspaceTypeOrder accordingly
func (airspacetype *AirspaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceTypes, stage.AirspaceType_stagedOrder, stage.AirspaceType_orderStaged, &stage.AirspaceTypeOrder, stage.AirspaceTypes_mapString, airspacetype, order, airspacetype.Name)
}

// Unstage removes airspacetype off the model stage
func (airspacetype *AirspaceType) Unstage(stage *Stage) *AirspaceType {
	__gong__unstage(stage.AirspaceTypes, stage.AirspaceTypes_mapString, airspacetype, airspacetype.Name)
	return airspacetype
}

// UnstageVoid removes airspacetype off the model stage
func (airspacetype *AirspaceType) UnstageVoid(stage *Stage) {
	airspacetype.Unstage(stage)
}

func (airspacetype *AirspaceType) StageVoid(stage *Stage) {
	airspacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacetype *AirspaceType) GetName() (res string) {
	return airspacetype.Name
}

// for satisfaction of GongStruct interface
func (airspacetype *AirspaceType) SetName(name string) {
	airspacetype.Name = name
}

// Stage puts airspacevolumedependencypropertytype to the model stage
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) Stage(stage *Stage) *AirspaceVolumeDependencyPropertyType {
	__gong__stage(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyType_stagedOrder, stage.AirspaceVolumeDependencyPropertyType_orderStaged, &stage.AirspaceVolumeDependencyPropertyTypeOrder, stage.AirspaceVolumeDependencyPropertyTypes_mapString, airspacevolumedependencypropertytype, airspacevolumedependencypropertytype.Name)
	return airspacevolumedependencypropertytype
}

// StagePreserveOrder puts airspacevolumedependencypropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceVolumeDependencyPropertyTypeOrder
// - update stage.AirspaceVolumeDependencyPropertyTypeOrder accordingly
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyType_stagedOrder, stage.AirspaceVolumeDependencyPropertyType_orderStaged, &stage.AirspaceVolumeDependencyPropertyTypeOrder, stage.AirspaceVolumeDependencyPropertyTypes_mapString, airspacevolumedependencypropertytype, order, airspacevolumedependencypropertytype.Name)
}

// Unstage removes airspacevolumedependencypropertytype off the model stage
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) Unstage(stage *Stage) *AirspaceVolumeDependencyPropertyType {
	__gong__unstage(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyTypes_mapString, airspacevolumedependencypropertytype, airspacevolumedependencypropertytype.Name)
	return airspacevolumedependencypropertytype
}

// UnstageVoid removes airspacevolumedependencypropertytype off the model stage
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) UnstageVoid(stage *Stage) {
	airspacevolumedependencypropertytype.Unstage(stage)
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) StageVoid(stage *Stage) {
	airspacevolumedependencypropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GetName() (res string) {
	return airspacevolumedependencypropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) SetName(name string) {
	airspacevolumedependencypropertytype.Name = name
}

// Stage puts airspacevolumedependencytype to the model stage
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) Stage(stage *Stage) *AirspaceVolumeDependencyType {
	__gong__stage(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyType_stagedOrder, stage.AirspaceVolumeDependencyType_orderStaged, &stage.AirspaceVolumeDependencyTypeOrder, stage.AirspaceVolumeDependencyTypes_mapString, airspacevolumedependencytype, airspacevolumedependencytype.Name)
	return airspacevolumedependencytype
}

// StagePreserveOrder puts airspacevolumedependencytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceVolumeDependencyTypeOrder
// - update stage.AirspaceVolumeDependencyTypeOrder accordingly
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyType_stagedOrder, stage.AirspaceVolumeDependencyType_orderStaged, &stage.AirspaceVolumeDependencyTypeOrder, stage.AirspaceVolumeDependencyTypes_mapString, airspacevolumedependencytype, order, airspacevolumedependencytype.Name)
}

// Unstage removes airspacevolumedependencytype off the model stage
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) Unstage(stage *Stage) *AirspaceVolumeDependencyType {
	__gong__unstage(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyTypes_mapString, airspacevolumedependencytype, airspacevolumedependencytype.Name)
	return airspacevolumedependencytype
}

// UnstageVoid removes airspacevolumedependencytype off the model stage
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) UnstageVoid(stage *Stage) {
	airspacevolumedependencytype.Unstage(stage)
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) StageVoid(stage *Stage) {
	airspacevolumedependencytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GetName() (res string) {
	return airspacevolumedependencytype.Name
}

// for satisfaction of GongStruct interface
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) SetName(name string) {
	airspacevolumedependencytype.Name = name
}

// Stage puts airspacevolumepropertytype to the model stage
func (airspacevolumepropertytype *AirspaceVolumePropertyType) Stage(stage *Stage) *AirspaceVolumePropertyType {
	__gong__stage(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyType_stagedOrder, stage.AirspaceVolumePropertyType_orderStaged, &stage.AirspaceVolumePropertyTypeOrder, stage.AirspaceVolumePropertyTypes_mapString, airspacevolumepropertytype, airspacevolumepropertytype.Name)
	return airspacevolumepropertytype
}

// StagePreserveOrder puts airspacevolumepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceVolumePropertyTypeOrder
// - update stage.AirspaceVolumePropertyTypeOrder accordingly
func (airspacevolumepropertytype *AirspaceVolumePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyType_stagedOrder, stage.AirspaceVolumePropertyType_orderStaged, &stage.AirspaceVolumePropertyTypeOrder, stage.AirspaceVolumePropertyTypes_mapString, airspacevolumepropertytype, order, airspacevolumepropertytype.Name)
}

// Unstage removes airspacevolumepropertytype off the model stage
func (airspacevolumepropertytype *AirspaceVolumePropertyType) Unstage(stage *Stage) *AirspaceVolumePropertyType {
	__gong__unstage(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyTypes_mapString, airspacevolumepropertytype, airspacevolumepropertytype.Name)
	return airspacevolumepropertytype
}

// UnstageVoid removes airspacevolumepropertytype off the model stage
func (airspacevolumepropertytype *AirspaceVolumePropertyType) UnstageVoid(stage *Stage) {
	airspacevolumepropertytype.Unstage(stage)
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) StageVoid(stage *Stage) {
	airspacevolumepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacevolumepropertytype *AirspaceVolumePropertyType) GetName() (res string) {
	return airspacevolumepropertytype.Name
}

// for satisfaction of GongStruct interface
func (airspacevolumepropertytype *AirspaceVolumePropertyType) SetName(name string) {
	airspacevolumepropertytype.Name = name
}

// Stage puts airspacevolumetype to the model stage
func (airspacevolumetype *AirspaceVolumeType) Stage(stage *Stage) *AirspaceVolumeType {
	__gong__stage(stage.AirspaceVolumeTypes, stage.AirspaceVolumeType_stagedOrder, stage.AirspaceVolumeType_orderStaged, &stage.AirspaceVolumeTypeOrder, stage.AirspaceVolumeTypes_mapString, airspacevolumetype, airspacevolumetype.Name)
	return airspacevolumetype
}

// StagePreserveOrder puts airspacevolumetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AirspaceVolumeTypeOrder
// - update stage.AirspaceVolumeTypeOrder accordingly
func (airspacevolumetype *AirspaceVolumeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AirspaceVolumeTypes, stage.AirspaceVolumeType_stagedOrder, stage.AirspaceVolumeType_orderStaged, &stage.AirspaceVolumeTypeOrder, stage.AirspaceVolumeTypes_mapString, airspacevolumetype, order, airspacevolumetype.Name)
}

// Unstage removes airspacevolumetype off the model stage
func (airspacevolumetype *AirspaceVolumeType) Unstage(stage *Stage) *AirspaceVolumeType {
	__gong__unstage(stage.AirspaceVolumeTypes, stage.AirspaceVolumeTypes_mapString, airspacevolumetype, airspacevolumetype.Name)
	return airspacevolumetype
}

// UnstageVoid removes airspacevolumetype off the model stage
func (airspacevolumetype *AirspaceVolumeType) UnstageVoid(stage *Stage) {
	airspacevolumetype.Unstage(stage)
}

func (airspacevolumetype *AirspaceVolumeType) StageVoid(stage *Stage) {
	airspacevolumetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (airspacevolumetype *AirspaceVolumeType) GetName() (res string) {
	return airspacevolumetype.Name
}

// for satisfaction of GongStruct interface
func (airspacevolumetype *AirspaceVolumeType) SetName(name string) {
	airspacevolumetype.Name = name
}

// Stage puts authorityforairspacetimeslicepropertytype to the model stage
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) Stage(stage *Stage) *AuthorityForAirspaceTimeSlicePropertyType {
	__gong__stage(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder, stage.AuthorityForAirspaceTimeSlicePropertyType_orderStaged, &stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder, stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString, authorityforairspacetimeslicepropertytype, authorityforairspacetimeslicepropertytype.Name)
	return authorityforairspacetimeslicepropertytype
}

// StagePreserveOrder puts authorityforairspacetimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder
// - update stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder accordingly
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder, stage.AuthorityForAirspaceTimeSlicePropertyType_orderStaged, &stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder, stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString, authorityforairspacetimeslicepropertytype, order, authorityforairspacetimeslicepropertytype.Name)
}

// Unstage removes authorityforairspacetimeslicepropertytype off the model stage
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) Unstage(stage *Stage) *AuthorityForAirspaceTimeSlicePropertyType {
	__gong__unstage(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString, authorityforairspacetimeslicepropertytype, authorityforairspacetimeslicepropertytype.Name)
	return authorityforairspacetimeslicepropertytype
}

// UnstageVoid removes authorityforairspacetimeslicepropertytype off the model stage
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	authorityforairspacetimeslicepropertytype.Unstage(stage)
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) StageVoid(stage *Stage) {
	authorityforairspacetimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GetName() (res string) {
	return authorityforairspacetimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) SetName(name string) {
	authorityforairspacetimeslicepropertytype.Name = name
}

// Stage puts authorityforairspacetimeslicetype to the model stage
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) Stage(stage *Stage) *AuthorityForAirspaceTimeSliceType {
	__gong__stage(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceType_stagedOrder, stage.AuthorityForAirspaceTimeSliceType_orderStaged, &stage.AuthorityForAirspaceTimeSliceTypeOrder, stage.AuthorityForAirspaceTimeSliceTypes_mapString, authorityforairspacetimeslicetype, authorityforairspacetimeslicetype.Name)
	return authorityforairspacetimeslicetype
}

// StagePreserveOrder puts authorityforairspacetimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AuthorityForAirspaceTimeSliceTypeOrder
// - update stage.AuthorityForAirspaceTimeSliceTypeOrder accordingly
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceType_stagedOrder, stage.AuthorityForAirspaceTimeSliceType_orderStaged, &stage.AuthorityForAirspaceTimeSliceTypeOrder, stage.AuthorityForAirspaceTimeSliceTypes_mapString, authorityforairspacetimeslicetype, order, authorityforairspacetimeslicetype.Name)
}

// Unstage removes authorityforairspacetimeslicetype off the model stage
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) Unstage(stage *Stage) *AuthorityForAirspaceTimeSliceType {
	__gong__unstage(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceTypes_mapString, authorityforairspacetimeslicetype, authorityforairspacetimeslicetype.Name)
	return authorityforairspacetimeslicetype
}

// UnstageVoid removes authorityforairspacetimeslicetype off the model stage
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) UnstageVoid(stage *Stage) {
	authorityforairspacetimeslicetype.Unstage(stage)
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) StageVoid(stage *Stage) {
	authorityforairspacetimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GetName() (res string) {
	return authorityforairspacetimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) SetName(name string) {
	authorityforairspacetimeslicetype.Name = name
}

// Stage puts authorityforairspacetype to the model stage
func (authorityforairspacetype *AuthorityForAirspaceType) Stage(stage *Stage) *AuthorityForAirspaceType {
	__gong__stage(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceType_stagedOrder, stage.AuthorityForAirspaceType_orderStaged, &stage.AuthorityForAirspaceTypeOrder, stage.AuthorityForAirspaceTypes_mapString, authorityforairspacetype, authorityforairspacetype.Name)
	return authorityforairspacetype
}

// StagePreserveOrder puts authorityforairspacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AuthorityForAirspaceTypeOrder
// - update stage.AuthorityForAirspaceTypeOrder accordingly
func (authorityforairspacetype *AuthorityForAirspaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceType_stagedOrder, stage.AuthorityForAirspaceType_orderStaged, &stage.AuthorityForAirspaceTypeOrder, stage.AuthorityForAirspaceTypes_mapString, authorityforairspacetype, order, authorityforairspacetype.Name)
}

// Unstage removes authorityforairspacetype off the model stage
func (authorityforairspacetype *AuthorityForAirspaceType) Unstage(stage *Stage) *AuthorityForAirspaceType {
	__gong__unstage(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceTypes_mapString, authorityforairspacetype, authorityforairspacetype.Name)
	return authorityforairspacetype
}

// UnstageVoid removes authorityforairspacetype off the model stage
func (authorityforairspacetype *AuthorityForAirspaceType) UnstageVoid(stage *Stage) {
	authorityforairspacetype.Unstage(stage)
}

func (authorityforairspacetype *AuthorityForAirspaceType) StageVoid(stage *Stage) {
	authorityforairspacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (authorityforairspacetype *AuthorityForAirspaceType) GetName() (res string) {
	return authorityforairspacetype.Name
}

// for satisfaction of GongStruct interface
func (authorityforairspacetype *AuthorityForAirspaceType) SetName(name string) {
	authorityforairspacetype.Name = name
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

// Stage puts geobordertimeslicepropertytype to the model stage
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) Stage(stage *Stage) *GeoBorderTimeSlicePropertyType {
	__gong__stage(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyType_stagedOrder, stage.GeoBorderTimeSlicePropertyType_orderStaged, &stage.GeoBorderTimeSlicePropertyTypeOrder, stage.GeoBorderTimeSlicePropertyTypes_mapString, geobordertimeslicepropertytype, geobordertimeslicepropertytype.Name)
	return geobordertimeslicepropertytype
}

// StagePreserveOrder puts geobordertimeslicepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoBorderTimeSlicePropertyTypeOrder
// - update stage.GeoBorderTimeSlicePropertyTypeOrder accordingly
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyType_stagedOrder, stage.GeoBorderTimeSlicePropertyType_orderStaged, &stage.GeoBorderTimeSlicePropertyTypeOrder, stage.GeoBorderTimeSlicePropertyTypes_mapString, geobordertimeslicepropertytype, order, geobordertimeslicepropertytype.Name)
}

// Unstage removes geobordertimeslicepropertytype off the model stage
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) Unstage(stage *Stage) *GeoBorderTimeSlicePropertyType {
	__gong__unstage(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyTypes_mapString, geobordertimeslicepropertytype, geobordertimeslicepropertytype.Name)
	return geobordertimeslicepropertytype
}

// UnstageVoid removes geobordertimeslicepropertytype off the model stage
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) UnstageVoid(stage *Stage) {
	geobordertimeslicepropertytype.Unstage(stage)
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) StageVoid(stage *Stage) {
	geobordertimeslicepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GetName() (res string) {
	return geobordertimeslicepropertytype.Name
}

// for satisfaction of GongStruct interface
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) SetName(name string) {
	geobordertimeslicepropertytype.Name = name
}

// Stage puts geobordertimeslicetype to the model stage
func (geobordertimeslicetype *GeoBorderTimeSliceType) Stage(stage *Stage) *GeoBorderTimeSliceType {
	__gong__stage(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceType_stagedOrder, stage.GeoBorderTimeSliceType_orderStaged, &stage.GeoBorderTimeSliceTypeOrder, stage.GeoBorderTimeSliceTypes_mapString, geobordertimeslicetype, geobordertimeslicetype.Name)
	return geobordertimeslicetype
}

// StagePreserveOrder puts geobordertimeslicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoBorderTimeSliceTypeOrder
// - update stage.GeoBorderTimeSliceTypeOrder accordingly
func (geobordertimeslicetype *GeoBorderTimeSliceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceType_stagedOrder, stage.GeoBorderTimeSliceType_orderStaged, &stage.GeoBorderTimeSliceTypeOrder, stage.GeoBorderTimeSliceTypes_mapString, geobordertimeslicetype, order, geobordertimeslicetype.Name)
}

// Unstage removes geobordertimeslicetype off the model stage
func (geobordertimeslicetype *GeoBorderTimeSliceType) Unstage(stage *Stage) *GeoBorderTimeSliceType {
	__gong__unstage(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceTypes_mapString, geobordertimeslicetype, geobordertimeslicetype.Name)
	return geobordertimeslicetype
}

// UnstageVoid removes geobordertimeslicetype off the model stage
func (geobordertimeslicetype *GeoBorderTimeSliceType) UnstageVoid(stage *Stage) {
	geobordertimeslicetype.Unstage(stage)
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) StageVoid(stage *Stage) {
	geobordertimeslicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geobordertimeslicetype *GeoBorderTimeSliceType) GetName() (res string) {
	return geobordertimeslicetype.Name
}

// for satisfaction of GongStruct interface
func (geobordertimeslicetype *GeoBorderTimeSliceType) SetName(name string) {
	geobordertimeslicetype.Name = name
}

// Stage puts geobordertype to the model stage
func (geobordertype *GeoBorderType) Stage(stage *Stage) *GeoBorderType {
	__gong__stage(stage.GeoBorderTypes, stage.GeoBorderType_stagedOrder, stage.GeoBorderType_orderStaged, &stage.GeoBorderTypeOrder, stage.GeoBorderTypes_mapString, geobordertype, geobordertype.Name)
	return geobordertype
}

// StagePreserveOrder puts geobordertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GeoBorderTypeOrder
// - update stage.GeoBorderTypeOrder accordingly
func (geobordertype *GeoBorderType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.GeoBorderTypes, stage.GeoBorderType_stagedOrder, stage.GeoBorderType_orderStaged, &stage.GeoBorderTypeOrder, stage.GeoBorderTypes_mapString, geobordertype, order, geobordertype.Name)
}

// Unstage removes geobordertype off the model stage
func (geobordertype *GeoBorderType) Unstage(stage *Stage) *GeoBorderType {
	__gong__unstage(stage.GeoBorderTypes, stage.GeoBorderTypes_mapString, geobordertype, geobordertype.Name)
	return geobordertype
}

// UnstageVoid removes geobordertype off the model stage
func (geobordertype *GeoBorderType) UnstageVoid(stage *Stage) {
	geobordertype.Unstage(stage)
}

func (geobordertype *GeoBorderType) StageVoid(stage *Stage) {
	geobordertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (geobordertype *GeoBorderType) GetName() (res string) {
	return geobordertype.Name
}

// for satisfaction of GongStruct interface
func (geobordertype *GeoBorderType) SetName(name string) {
	geobordertype.Name = name
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

// Stage puts routepropertytype to the model stage
func (routepropertytype *RoutePropertyType) Stage(stage *Stage) *RoutePropertyType {
	__gong__stage(stage.RoutePropertyTypes, stage.RoutePropertyType_stagedOrder, stage.RoutePropertyType_orderStaged, &stage.RoutePropertyTypeOrder, stage.RoutePropertyTypes_mapString, routepropertytype, routepropertytype.Name)
	return routepropertytype
}

// StagePreserveOrder puts routepropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RoutePropertyTypeOrder
// - update stage.RoutePropertyTypeOrder accordingly
func (routepropertytype *RoutePropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RoutePropertyTypes, stage.RoutePropertyType_stagedOrder, stage.RoutePropertyType_orderStaged, &stage.RoutePropertyTypeOrder, stage.RoutePropertyTypes_mapString, routepropertytype, order, routepropertytype.Name)
}

// Unstage removes routepropertytype off the model stage
func (routepropertytype *RoutePropertyType) Unstage(stage *Stage) *RoutePropertyType {
	__gong__unstage(stage.RoutePropertyTypes, stage.RoutePropertyTypes_mapString, routepropertytype, routepropertytype.Name)
	return routepropertytype
}

// UnstageVoid removes routepropertytype off the model stage
func (routepropertytype *RoutePropertyType) UnstageVoid(stage *Stage) {
	routepropertytype.Unstage(stage)
}

func (routepropertytype *RoutePropertyType) StageVoid(stage *Stage) {
	routepropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (routepropertytype *RoutePropertyType) GetName() (res string) {
	return routepropertytype.Name
}

// for satisfaction of GongStruct interface
func (routepropertytype *RoutePropertyType) SetName(name string) {
	routepropertytype.Name = name
}

// Stage puts standardlevelcolumnpropertytype to the model stage
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) Stage(stage *Stage) *StandardLevelColumnPropertyType {
	__gong__stage(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyType_stagedOrder, stage.StandardLevelColumnPropertyType_orderStaged, &stage.StandardLevelColumnPropertyTypeOrder, stage.StandardLevelColumnPropertyTypes_mapString, standardlevelcolumnpropertytype, standardlevelcolumnpropertytype.Name)
	return standardlevelcolumnpropertytype
}

// StagePreserveOrder puts standardlevelcolumnpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StandardLevelColumnPropertyTypeOrder
// - update stage.StandardLevelColumnPropertyTypeOrder accordingly
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyType_stagedOrder, stage.StandardLevelColumnPropertyType_orderStaged, &stage.StandardLevelColumnPropertyTypeOrder, stage.StandardLevelColumnPropertyTypes_mapString, standardlevelcolumnpropertytype, order, standardlevelcolumnpropertytype.Name)
}

// Unstage removes standardlevelcolumnpropertytype off the model stage
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) Unstage(stage *Stage) *StandardLevelColumnPropertyType {
	__gong__unstage(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyTypes_mapString, standardlevelcolumnpropertytype, standardlevelcolumnpropertytype.Name)
	return standardlevelcolumnpropertytype
}

// UnstageVoid removes standardlevelcolumnpropertytype off the model stage
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) UnstageVoid(stage *Stage) {
	standardlevelcolumnpropertytype.Unstage(stage)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) StageVoid(stage *Stage) {
	standardlevelcolumnpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GetName() (res string) {
	return standardlevelcolumnpropertytype.Name
}

// for satisfaction of GongStruct interface
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) SetName(name string) {
	standardlevelcolumnpropertytype.Name = name
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

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.A_extensions, &stage.A_extensions_mapString, &stage.A_extension_stagedOrder, &stage.A_extensionOrder)

	__gong__resetStageType(&stage.A_extension_1s, &stage.A_extension_1s_mapString, &stage.A_extension_1_stagedOrder, &stage.A_extension_1Order)

	__gong__resetStageType(&stage.A_extension_10s, &stage.A_extension_10s_mapString, &stage.A_extension_10_stagedOrder, &stage.A_extension_10Order)

	__gong__resetStageType(&stage.A_extension_11s, &stage.A_extension_11s_mapString, &stage.A_extension_11_stagedOrder, &stage.A_extension_11Order)

	__gong__resetStageType(&stage.A_extension_12s, &stage.A_extension_12s_mapString, &stage.A_extension_12_stagedOrder, &stage.A_extension_12Order)

	__gong__resetStageType(&stage.A_extension_13s, &stage.A_extension_13s_mapString, &stage.A_extension_13_stagedOrder, &stage.A_extension_13Order)

	__gong__resetStageType(&stage.A_extension_14s, &stage.A_extension_14s_mapString, &stage.A_extension_14_stagedOrder, &stage.A_extension_14Order)

	__gong__resetStageType(&stage.A_extension_2s, &stage.A_extension_2s_mapString, &stage.A_extension_2_stagedOrder, &stage.A_extension_2Order)

	__gong__resetStageType(&stage.A_extension_3s, &stage.A_extension_3s_mapString, &stage.A_extension_3_stagedOrder, &stage.A_extension_3Order)

	__gong__resetStageType(&stage.A_extension_4s, &stage.A_extension_4s_mapString, &stage.A_extension_4_stagedOrder, &stage.A_extension_4Order)

	__gong__resetStageType(&stage.A_extension_5s, &stage.A_extension_5s_mapString, &stage.A_extension_5_stagedOrder, &stage.A_extension_5Order)

	__gong__resetStageType(&stage.A_extension_6s, &stage.A_extension_6s_mapString, &stage.A_extension_6_stagedOrder, &stage.A_extension_6Order)

	__gong__resetStageType(&stage.A_extension_7s, &stage.A_extension_7s_mapString, &stage.A_extension_7_stagedOrder, &stage.A_extension_7Order)

	__gong__resetStageType(&stage.A_extension_8s, &stage.A_extension_8s_mapString, &stage.A_extension_8_stagedOrder, &stage.A_extension_8Order)

	__gong__resetStageType(&stage.A_extension_9s, &stage.A_extension_9s_mapString, &stage.A_extension_9_stagedOrder, &stage.A_extension_9Order)

	__gong__resetStageType(&stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_mapString, &stage.AbstractPropertiesWithScheduleType_stagedOrder, &stage.AbstractPropertiesWithScheduleTypeOrder)

	__gong__resetStageType(&stage.AircraftCharacteristicPropertyTypes, &stage.AircraftCharacteristicPropertyTypes_mapString, &stage.AircraftCharacteristicPropertyType_stagedOrder, &stage.AircraftCharacteristicPropertyTypeOrder)

	__gong__resetStageType(&stage.AircraftCharacteristicTypes, &stage.AircraftCharacteristicTypes_mapString, &stage.AircraftCharacteristicType_stagedOrder, &stage.AircraftCharacteristicTypeOrder)

	__gong__resetStageType(&stage.AirspaceActivationPropertyTypes, &stage.AirspaceActivationPropertyTypes_mapString, &stage.AirspaceActivationPropertyType_stagedOrder, &stage.AirspaceActivationPropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceActivationTypes, &stage.AirspaceActivationTypes_mapString, &stage.AirspaceActivationType_stagedOrder, &stage.AirspaceActivationTypeOrder)

	__gong__resetStageType(&stage.AirspaceGeometryComponentPropertyTypes, &stage.AirspaceGeometryComponentPropertyTypes_mapString, &stage.AirspaceGeometryComponentPropertyType_stagedOrder, &stage.AirspaceGeometryComponentPropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceGeometryComponentTypes, &stage.AirspaceGeometryComponentTypes_mapString, &stage.AirspaceGeometryComponentType_stagedOrder, &stage.AirspaceGeometryComponentTypeOrder)

	__gong__resetStageType(&stage.AirspaceLayerClassPropertyTypes, &stage.AirspaceLayerClassPropertyTypes_mapString, &stage.AirspaceLayerClassPropertyType_stagedOrder, &stage.AirspaceLayerClassPropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceLayerClassTypes, &stage.AirspaceLayerClassTypes_mapString, &stage.AirspaceLayerClassType_stagedOrder, &stage.AirspaceLayerClassTypeOrder)

	__gong__resetStageType(&stage.AirspaceLayerPropertyTypes, &stage.AirspaceLayerPropertyTypes_mapString, &stage.AirspaceLayerPropertyType_stagedOrder, &stage.AirspaceLayerPropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceLayerTypes, &stage.AirspaceLayerTypes_mapString, &stage.AirspaceLayerType_stagedOrder, &stage.AirspaceLayerTypeOrder)

	__gong__resetStageType(&stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_mapString, &stage.AirspacePropertyType_stagedOrder, &stage.AirspacePropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceTimeSlicePropertyTypes, &stage.AirspaceTimeSlicePropertyTypes_mapString, &stage.AirspaceTimeSlicePropertyType_stagedOrder, &stage.AirspaceTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceTimeSliceTypes, &stage.AirspaceTimeSliceTypes_mapString, &stage.AirspaceTimeSliceType_stagedOrder, &stage.AirspaceTimeSliceTypeOrder)

	__gong__resetStageType(&stage.AirspaceTypes, &stage.AirspaceTypes_mapString, &stage.AirspaceType_stagedOrder, &stage.AirspaceTypeOrder)

	__gong__resetStageType(&stage.AirspaceVolumeDependencyPropertyTypes, &stage.AirspaceVolumeDependencyPropertyTypes_mapString, &stage.AirspaceVolumeDependencyPropertyType_stagedOrder, &stage.AirspaceVolumeDependencyPropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceVolumeDependencyTypes, &stage.AirspaceVolumeDependencyTypes_mapString, &stage.AirspaceVolumeDependencyType_stagedOrder, &stage.AirspaceVolumeDependencyTypeOrder)

	__gong__resetStageType(&stage.AirspaceVolumePropertyTypes, &stage.AirspaceVolumePropertyTypes_mapString, &stage.AirspaceVolumePropertyType_stagedOrder, &stage.AirspaceVolumePropertyTypeOrder)

	__gong__resetStageType(&stage.AirspaceVolumeTypes, &stage.AirspaceVolumeTypes_mapString, &stage.AirspaceVolumeType_stagedOrder, &stage.AirspaceVolumeTypeOrder)

	__gong__resetStageType(&stage.AuthorityForAirspaceTimeSlicePropertyTypes, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString, &stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder, &stage.AuthorityForAirspaceTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.AuthorityForAirspaceTimeSliceTypes, &stage.AuthorityForAirspaceTimeSliceTypes_mapString, &stage.AuthorityForAirspaceTimeSliceType_stagedOrder, &stage.AuthorityForAirspaceTimeSliceTypeOrder)

	__gong__resetStageType(&stage.AuthorityForAirspaceTypes, &stage.AuthorityForAirspaceTypes_mapString, &stage.AuthorityForAirspaceType_stagedOrder, &stage.AuthorityForAirspaceTypeOrder)

	__gong__resetStageType(&stage.CurvePropertyTypes, &stage.CurvePropertyTypes_mapString, &stage.CurvePropertyType_stagedOrder, &stage.CurvePropertyTypeOrder)

	__gong__resetStageType(&stage.CurveTypes, &stage.CurveTypes_mapString, &stage.CurveType_stagedOrder, &stage.CurveTypeOrder)

	__gong__resetStageType(&stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_mapString, &stage.ElevatedCurveType_stagedOrder, &stage.ElevatedCurveTypeOrder)

	__gong__resetStageType(&stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_mapString, &stage.ElevatedSurfaceType_stagedOrder, &stage.ElevatedSurfaceTypeOrder)

	__gong__resetStageType(&stage.GeoBorderTimeSlicePropertyTypes, &stage.GeoBorderTimeSlicePropertyTypes_mapString, &stage.GeoBorderTimeSlicePropertyType_stagedOrder, &stage.GeoBorderTimeSlicePropertyTypeOrder)

	__gong__resetStageType(&stage.GeoBorderTimeSliceTypes, &stage.GeoBorderTimeSliceTypes_mapString, &stage.GeoBorderTimeSliceType_stagedOrder, &stage.GeoBorderTimeSliceTypeOrder)

	__gong__resetStageType(&stage.GeoBorderTypes, &stage.GeoBorderTypes_mapString, &stage.GeoBorderType_stagedOrder, &stage.GeoBorderTypeOrder)

	__gong__resetStageType(&stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_mapString, &stage.LinguisticNotePropertyType_stagedOrder, &stage.LinguisticNotePropertyTypeOrder)

	__gong__resetStageType(&stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_mapString, &stage.LinguisticNoteType_stagedOrder, &stage.LinguisticNoteTypeOrder)

	__gong__resetStageType(&stage.NotePropertyTypes, &stage.NotePropertyTypes_mapString, &stage.NotePropertyType_stagedOrder, &stage.NotePropertyTypeOrder)

	__gong__resetStageType(&stage.NoteTypes, &stage.NoteTypes_mapString, &stage.NoteType_stagedOrder, &stage.NoteTypeOrder)

	__gong__resetStageType(&stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_mapString, &stage.OrganisationAuthorityPropertyType_stagedOrder, &stage.OrganisationAuthorityPropertyTypeOrder)

	__gong__resetStageType(&stage.RoutePropertyTypes, &stage.RoutePropertyTypes_mapString, &stage.RoutePropertyType_stagedOrder, &stage.RoutePropertyTypeOrder)

	__gong__resetStageType(&stage.StandardLevelColumnPropertyTypes, &stage.StandardLevelColumnPropertyTypes_mapString, &stage.StandardLevelColumnPropertyType_stagedOrder, &stage.StandardLevelColumnPropertyTypeOrder)

	__gong__resetStageType(&stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_mapString, &stage.SurfacePropertyType_stagedOrder, &stage.SurfacePropertyTypeOrder)

	__gong__resetStageType(&stage.SurfaceTypes, &stage.SurfaceTypes_mapString, &stage.SurfaceType_stagedOrder, &stage.SurfaceTypeOrder)

	__gong__resetStageType(&stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_mapString, &stage.TimesheetPropertyType_stagedOrder, &stage.TimesheetPropertyTypeOrder)

	__gong__resetStageType(&stage.TimesheetTypes, &stage.TimesheetTypes_mapString, &stage.TimesheetType_stagedOrder, &stage.TimesheetTypeOrder)

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

func (A_extension_13) GongGetAssociationName() any {
	return &A_extension_13{
	}
}

func (A_extension_14) GongGetAssociationName() any {
	return &A_extension_14{
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

func (AircraftCharacteristicPropertyType) GongGetAssociationName() any {
	return &AircraftCharacteristicPropertyType{
			AircraftCharacteristic: &AircraftCharacteristicType{Name: "AircraftCharacteristic"},
	}
}

func (AircraftCharacteristicType) GongGetAssociationName() any {
	return &AircraftCharacteristicType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_12{{Name: "Extension"}},
	}
}

func (AirspaceActivationPropertyType) GongGetAssociationName() any {
	return &AirspaceActivationPropertyType{
			AirspaceActivation: &AirspaceActivationType{Name: "AirspaceActivation"},
	}
}

func (AirspaceActivationType) GongGetAssociationName() any {
	return &AirspaceActivationType{
			TimeInterval: []*TimesheetPropertyType{{Name: "TimeInterval"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			SpecialDateAuthority: []*OrganisationAuthorityPropertyType{{Name: "SpecialDateAuthority"}},
			Levels: []*AirspaceLayerPropertyType{{Name: "Levels"}},
			User: []*OrganisationAuthorityPropertyType{{Name: "User"}},
			Aircraft: []*AircraftCharacteristicPropertyType{{Name: "Aircraft"}},
			Extension: []*A_extension_1{{Name: "Extension"}},
	}
}

func (AirspaceGeometryComponentPropertyType) GongGetAssociationName() any {
	return &AirspaceGeometryComponentPropertyType{
			AirspaceGeometryComponent: &AirspaceGeometryComponentType{Name: "AirspaceGeometryComponent"},
	}
}

func (AirspaceGeometryComponentType) GongGetAssociationName() any {
	return &AirspaceGeometryComponentType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			TheAirspaceVolume: &AirspaceVolumePropertyType{Name: "TheAirspaceVolume"},
			Extension: []*A_extension_2{{Name: "Extension"}},
	}
}

func (AirspaceLayerClassPropertyType) GongGetAssociationName() any {
	return &AirspaceLayerClassPropertyType{
			AirspaceLayerClass: &AirspaceLayerClassType{Name: "AirspaceLayerClass"},
	}
}

func (AirspaceLayerClassType) GongGetAssociationName() any {
	return &AirspaceLayerClassType{
			TimeInterval: []*TimesheetPropertyType{{Name: "TimeInterval"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			SpecialDateAuthority: []*OrganisationAuthorityPropertyType{{Name: "SpecialDateAuthority"}},
			AssociatedLevels: []*AirspaceLayerPropertyType{{Name: "AssociatedLevels"}},
			Extension: []*A_extension_3{{Name: "Extension"}},
	}
}

func (AirspaceLayerPropertyType) GongGetAssociationName() any {
	return &AirspaceLayerPropertyType{
			AirspaceLayer: &AirspaceLayerType{Name: "AirspaceLayer"},
	}
}

func (AirspaceLayerType) GongGetAssociationName() any {
	return &AirspaceLayerType{
			DiscreteLevelSeries: &StandardLevelColumnPropertyType{Name: "DiscreteLevelSeries"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_13{{Name: "Extension"}},
	}
}

func (AirspacePropertyType) GongGetAssociationName() any {
	return &AirspacePropertyType{
	}
}

func (AirspaceTimeSlicePropertyType) GongGetAssociationName() any {
	return &AirspaceTimeSlicePropertyType{
			AirspaceTimeSlice: &AirspaceTimeSliceType{Name: "AirspaceTimeSlice"},
	}
}

func (AirspaceTimeSliceType) GongGetAssociationName() any {
	return &AirspaceTimeSliceType{
			Class: []*AirspaceLayerClassPropertyType{{Name: "Class"}},
			ProtectedRoute: &RoutePropertyType{Name: "ProtectedRoute"},
			GeometryComponent: []*AirspaceGeometryComponentPropertyType{{Name: "GeometryComponent"}},
			Activation: []*AirspaceActivationPropertyType{{Name: "Activation"}},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension{{Name: "Extension"}},
	}
}

func (AirspaceType) GongGetAssociationName() any {
	return &AirspaceType{
			TimeSlice: []*AirspaceTimeSlicePropertyType{{Name: "TimeSlice"}},
	}
}

func (AirspaceVolumeDependencyPropertyType) GongGetAssociationName() any {
	return &AirspaceVolumeDependencyPropertyType{
			AirspaceVolumeDependency: &AirspaceVolumeDependencyType{Name: "AirspaceVolumeDependency"},
	}
}

func (AirspaceVolumeDependencyType) GongGetAssociationName() any {
	return &AirspaceVolumeDependencyType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			TheAirspace: &AirspacePropertyType{Name: "TheAirspace"},
			Extension: []*A_extension_5{{Name: "Extension"}},
	}
}

func (AirspaceVolumePropertyType) GongGetAssociationName() any {
	return &AirspaceVolumePropertyType{
			AirspaceVolume: &AirspaceVolumeType{Name: "AirspaceVolume"},
	}
}

func (AirspaceVolumeType) GongGetAssociationName() any {
	return &AirspaceVolumeType{
			HorizontalProjection: &SurfacePropertyType{Name: "HorizontalProjection"},
			Centreline: &CurvePropertyType{Name: "Centreline"},
			ContributorAirspace: &AirspaceVolumeDependencyPropertyType{Name: "ContributorAirspace"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_4{{Name: "Extension"}},
	}
}

func (AuthorityForAirspaceTimeSlicePropertyType) GongGetAssociationName() any {
	return &AuthorityForAirspaceTimeSlicePropertyType{
			AuthorityForAirspaceTimeSlice: &AuthorityForAirspaceTimeSliceType{Name: "AuthorityForAirspaceTimeSlice"},
	}
}

func (AuthorityForAirspaceTimeSliceType) GongGetAssociationName() any {
	return &AuthorityForAirspaceTimeSliceType{
			ResponsibleOrganisation: &OrganisationAuthorityPropertyType{Name: "ResponsibleOrganisation"},
			AssignedAirspace: &AirspacePropertyType{Name: "AssignedAirspace"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_6{{Name: "Extension"}},
	}
}

func (AuthorityForAirspaceType) GongGetAssociationName() any {
	return &AuthorityForAirspaceType{
			TimeSlice: []*AuthorityForAirspaceTimeSlicePropertyType{{Name: "TimeSlice"}},
	}
}

func (CurvePropertyType) GongGetAssociationName() any {
	return &CurvePropertyType{
			Curve: &CurveType{Name: "Curve"},
			ElevatedCurve: &ElevatedCurveType{Name: "ElevatedCurve"},
	}
}

func (CurveType) GongGetAssociationName() any {
	return &CurveType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
	}
}

func (ElevatedCurveType) GongGetAssociationName() any {
	return &ElevatedCurveType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_8{{Name: "Extension"}},
	}
}

func (ElevatedSurfaceType) GongGetAssociationName() any {
	return &ElevatedSurfaceType{
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_9{{Name: "Extension"}},
	}
}

func (GeoBorderTimeSlicePropertyType) GongGetAssociationName() any {
	return &GeoBorderTimeSlicePropertyType{
			GeoBorderTimeSlice: &GeoBorderTimeSliceType{Name: "GeoBorderTimeSlice"},
	}
}

func (GeoBorderTimeSliceType) GongGetAssociationName() any {
	return &GeoBorderTimeSliceType{
			Border: &CurvePropertyType{Name: "Border"},
			Annotation: []*NotePropertyType{{Name: "Annotation"}},
			Extension: []*A_extension_7{{Name: "Extension"}},
	}
}

func (GeoBorderType) GongGetAssociationName() any {
	return &GeoBorderType{
			TimeSlice: []*GeoBorderTimeSlicePropertyType{{Name: "TimeSlice"}},
	}
}

func (LinguisticNotePropertyType) GongGetAssociationName() any {
	return &LinguisticNotePropertyType{
			LinguisticNote: &LinguisticNoteType{Name: "LinguisticNote"},
	}
}

func (LinguisticNoteType) GongGetAssociationName() any {
	return &LinguisticNoteType{
			Extension: []*A_extension_10{{Name: "Extension"}},
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
			Extension: []*A_extension_11{{Name: "Extension"}},
	}
}

func (OrganisationAuthorityPropertyType) GongGetAssociationName() any {
	return &OrganisationAuthorityPropertyType{
	}
}

func (RoutePropertyType) GongGetAssociationName() any {
	return &RoutePropertyType{
	}
}

func (StandardLevelColumnPropertyType) GongGetAssociationName() any {
	return &StandardLevelColumnPropertyType{
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
			Extension: []*A_extension_14{{Name: "Extension"}},
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
	// reverse maps of direct associations of A_extension_13
	case A_extension_13:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_14
	case A_extension_14:
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
	// reverse maps of direct associations of AircraftCharacteristicPropertyType
	case AircraftCharacteristicPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AircraftCharacteristic":
			res := make(map[*AircraftCharacteristicType][]*AircraftCharacteristicPropertyType)
			for aircraftcharacteristicpropertytype := range stage.AircraftCharacteristicPropertyTypes {
				if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
					aircraftcharacteristictype_ := aircraftcharacteristicpropertytype.AircraftCharacteristic
					var aircraftcharacteristicpropertytypes []*AircraftCharacteristicPropertyType
					_, ok := res[aircraftcharacteristictype_]
					if ok {
						aircraftcharacteristicpropertytypes = res[aircraftcharacteristictype_]
					} else {
						aircraftcharacteristicpropertytypes = make([]*AircraftCharacteristicPropertyType, 0)
					}
					aircraftcharacteristicpropertytypes = append(aircraftcharacteristicpropertytypes, aircraftcharacteristicpropertytype)
					res[aircraftcharacteristictype_] = aircraftcharacteristicpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AircraftCharacteristicType
	case AircraftCharacteristicType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceActivationPropertyType
	case AirspaceActivationPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceActivation":
			res := make(map[*AirspaceActivationType][]*AirspaceActivationPropertyType)
			for airspaceactivationpropertytype := range stage.AirspaceActivationPropertyTypes {
				if airspaceactivationpropertytype.AirspaceActivation != nil {
					airspaceactivationtype_ := airspaceactivationpropertytype.AirspaceActivation
					var airspaceactivationpropertytypes []*AirspaceActivationPropertyType
					_, ok := res[airspaceactivationtype_]
					if ok {
						airspaceactivationpropertytypes = res[airspaceactivationtype_]
					} else {
						airspaceactivationpropertytypes = make([]*AirspaceActivationPropertyType, 0)
					}
					airspaceactivationpropertytypes = append(airspaceactivationpropertytypes, airspaceactivationpropertytype)
					res[airspaceactivationtype_] = airspaceactivationpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceActivationType
	case AirspaceActivationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceGeometryComponentPropertyType
	case AirspaceGeometryComponentPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceGeometryComponent":
			res := make(map[*AirspaceGeometryComponentType][]*AirspaceGeometryComponentPropertyType)
			for airspacegeometrycomponentpropertytype := range stage.AirspaceGeometryComponentPropertyTypes {
				if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
					airspacegeometrycomponenttype_ := airspacegeometrycomponentpropertytype.AirspaceGeometryComponent
					var airspacegeometrycomponentpropertytypes []*AirspaceGeometryComponentPropertyType
					_, ok := res[airspacegeometrycomponenttype_]
					if ok {
						airspacegeometrycomponentpropertytypes = res[airspacegeometrycomponenttype_]
					} else {
						airspacegeometrycomponentpropertytypes = make([]*AirspaceGeometryComponentPropertyType, 0)
					}
					airspacegeometrycomponentpropertytypes = append(airspacegeometrycomponentpropertytypes, airspacegeometrycomponentpropertytype)
					res[airspacegeometrycomponenttype_] = airspacegeometrycomponentpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceGeometryComponentType
	case AirspaceGeometryComponentType:
		switch fieldname {
		// insertion point for per direct association field
		case "TheAirspaceVolume":
			res := make(map[*AirspaceVolumePropertyType][]*AirspaceGeometryComponentType)
			for airspacegeometrycomponenttype := range stage.AirspaceGeometryComponentTypes {
				if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
					airspacevolumepropertytype_ := airspacegeometrycomponenttype.TheAirspaceVolume
					var airspacegeometrycomponenttypes []*AirspaceGeometryComponentType
					_, ok := res[airspacevolumepropertytype_]
					if ok {
						airspacegeometrycomponenttypes = res[airspacevolumepropertytype_]
					} else {
						airspacegeometrycomponenttypes = make([]*AirspaceGeometryComponentType, 0)
					}
					airspacegeometrycomponenttypes = append(airspacegeometrycomponenttypes, airspacegeometrycomponenttype)
					res[airspacevolumepropertytype_] = airspacegeometrycomponenttypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceLayerClassPropertyType
	case AirspaceLayerClassPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceLayerClass":
			res := make(map[*AirspaceLayerClassType][]*AirspaceLayerClassPropertyType)
			for airspacelayerclasspropertytype := range stage.AirspaceLayerClassPropertyTypes {
				if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
					airspacelayerclasstype_ := airspacelayerclasspropertytype.AirspaceLayerClass
					var airspacelayerclasspropertytypes []*AirspaceLayerClassPropertyType
					_, ok := res[airspacelayerclasstype_]
					if ok {
						airspacelayerclasspropertytypes = res[airspacelayerclasstype_]
					} else {
						airspacelayerclasspropertytypes = make([]*AirspaceLayerClassPropertyType, 0)
					}
					airspacelayerclasspropertytypes = append(airspacelayerclasspropertytypes, airspacelayerclasspropertytype)
					res[airspacelayerclasstype_] = airspacelayerclasspropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceLayerClassType
	case AirspaceLayerClassType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceLayerPropertyType
	case AirspaceLayerPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceLayer":
			res := make(map[*AirspaceLayerType][]*AirspaceLayerPropertyType)
			for airspacelayerpropertytype := range stage.AirspaceLayerPropertyTypes {
				if airspacelayerpropertytype.AirspaceLayer != nil {
					airspacelayertype_ := airspacelayerpropertytype.AirspaceLayer
					var airspacelayerpropertytypes []*AirspaceLayerPropertyType
					_, ok := res[airspacelayertype_]
					if ok {
						airspacelayerpropertytypes = res[airspacelayertype_]
					} else {
						airspacelayerpropertytypes = make([]*AirspaceLayerPropertyType, 0)
					}
					airspacelayerpropertytypes = append(airspacelayerpropertytypes, airspacelayerpropertytype)
					res[airspacelayertype_] = airspacelayerpropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceLayerType
	case AirspaceLayerType:
		switch fieldname {
		// insertion point for per direct association field
		case "DiscreteLevelSeries":
			res := make(map[*StandardLevelColumnPropertyType][]*AirspaceLayerType)
			for airspacelayertype := range stage.AirspaceLayerTypes {
				if airspacelayertype.DiscreteLevelSeries != nil {
					standardlevelcolumnpropertytype_ := airspacelayertype.DiscreteLevelSeries
					var airspacelayertypes []*AirspaceLayerType
					_, ok := res[standardlevelcolumnpropertytype_]
					if ok {
						airspacelayertypes = res[standardlevelcolumnpropertytype_]
					} else {
						airspacelayertypes = make([]*AirspaceLayerType, 0)
					}
					airspacelayertypes = append(airspacelayertypes, airspacelayertype)
					res[standardlevelcolumnpropertytype_] = airspacelayertypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspacePropertyType
	case AirspacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceTimeSlicePropertyType
	case AirspaceTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceTimeSlice":
			res := make(map[*AirspaceTimeSliceType][]*AirspaceTimeSlicePropertyType)
			for airspacetimeslicepropertytype := range stage.AirspaceTimeSlicePropertyTypes {
				if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
					airspacetimeslicetype_ := airspacetimeslicepropertytype.AirspaceTimeSlice
					var airspacetimeslicepropertytypes []*AirspaceTimeSlicePropertyType
					_, ok := res[airspacetimeslicetype_]
					if ok {
						airspacetimeslicepropertytypes = res[airspacetimeslicetype_]
					} else {
						airspacetimeslicepropertytypes = make([]*AirspaceTimeSlicePropertyType, 0)
					}
					airspacetimeslicepropertytypes = append(airspacetimeslicepropertytypes, airspacetimeslicepropertytype)
					res[airspacetimeslicetype_] = airspacetimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceTimeSliceType
	case AirspaceTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "ProtectedRoute":
			res := make(map[*RoutePropertyType][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				if airspacetimeslicetype.ProtectedRoute != nil {
					routepropertytype_ := airspacetimeslicetype.ProtectedRoute
					var airspacetimeslicetypes []*AirspaceTimeSliceType
					_, ok := res[routepropertytype_]
					if ok {
						airspacetimeslicetypes = res[routepropertytype_]
					} else {
						airspacetimeslicetypes = make([]*AirspaceTimeSliceType, 0)
					}
					airspacetimeslicetypes = append(airspacetimeslicetypes, airspacetimeslicetype)
					res[routepropertytype_] = airspacetimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceType
	case AirspaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceVolumeDependencyPropertyType
	case AirspaceVolumeDependencyPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceVolumeDependency":
			res := make(map[*AirspaceVolumeDependencyType][]*AirspaceVolumeDependencyPropertyType)
			for airspacevolumedependencypropertytype := range stage.AirspaceVolumeDependencyPropertyTypes {
				if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
					airspacevolumedependencytype_ := airspacevolumedependencypropertytype.AirspaceVolumeDependency
					var airspacevolumedependencypropertytypes []*AirspaceVolumeDependencyPropertyType
					_, ok := res[airspacevolumedependencytype_]
					if ok {
						airspacevolumedependencypropertytypes = res[airspacevolumedependencytype_]
					} else {
						airspacevolumedependencypropertytypes = make([]*AirspaceVolumeDependencyPropertyType, 0)
					}
					airspacevolumedependencypropertytypes = append(airspacevolumedependencypropertytypes, airspacevolumedependencypropertytype)
					res[airspacevolumedependencytype_] = airspacevolumedependencypropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceVolumeDependencyType
	case AirspaceVolumeDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		case "TheAirspace":
			res := make(map[*AirspacePropertyType][]*AirspaceVolumeDependencyType)
			for airspacevolumedependencytype := range stage.AirspaceVolumeDependencyTypes {
				if airspacevolumedependencytype.TheAirspace != nil {
					airspacepropertytype_ := airspacevolumedependencytype.TheAirspace
					var airspacevolumedependencytypes []*AirspaceVolumeDependencyType
					_, ok := res[airspacepropertytype_]
					if ok {
						airspacevolumedependencytypes = res[airspacepropertytype_]
					} else {
						airspacevolumedependencytypes = make([]*AirspaceVolumeDependencyType, 0)
					}
					airspacevolumedependencytypes = append(airspacevolumedependencytypes, airspacevolumedependencytype)
					res[airspacepropertytype_] = airspacevolumedependencytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceVolumePropertyType
	case AirspaceVolumePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AirspaceVolume":
			res := make(map[*AirspaceVolumeType][]*AirspaceVolumePropertyType)
			for airspacevolumepropertytype := range stage.AirspaceVolumePropertyTypes {
				if airspacevolumepropertytype.AirspaceVolume != nil {
					airspacevolumetype_ := airspacevolumepropertytype.AirspaceVolume
					var airspacevolumepropertytypes []*AirspaceVolumePropertyType
					_, ok := res[airspacevolumetype_]
					if ok {
						airspacevolumepropertytypes = res[airspacevolumetype_]
					} else {
						airspacevolumepropertytypes = make([]*AirspaceVolumePropertyType, 0)
					}
					airspacevolumepropertytypes = append(airspacevolumepropertytypes, airspacevolumepropertytype)
					res[airspacevolumetype_] = airspacevolumepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceVolumeType
	case AirspaceVolumeType:
		switch fieldname {
		// insertion point for per direct association field
		case "HorizontalProjection":
			res := make(map[*SurfacePropertyType][]*AirspaceVolumeType)
			for airspacevolumetype := range stage.AirspaceVolumeTypes {
				if airspacevolumetype.HorizontalProjection != nil {
					surfacepropertytype_ := airspacevolumetype.HorizontalProjection
					var airspacevolumetypes []*AirspaceVolumeType
					_, ok := res[surfacepropertytype_]
					if ok {
						airspacevolumetypes = res[surfacepropertytype_]
					} else {
						airspacevolumetypes = make([]*AirspaceVolumeType, 0)
					}
					airspacevolumetypes = append(airspacevolumetypes, airspacevolumetype)
					res[surfacepropertytype_] = airspacevolumetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Centreline":
			res := make(map[*CurvePropertyType][]*AirspaceVolumeType)
			for airspacevolumetype := range stage.AirspaceVolumeTypes {
				if airspacevolumetype.Centreline != nil {
					curvepropertytype_ := airspacevolumetype.Centreline
					var airspacevolumetypes []*AirspaceVolumeType
					_, ok := res[curvepropertytype_]
					if ok {
						airspacevolumetypes = res[curvepropertytype_]
					} else {
						airspacevolumetypes = make([]*AirspaceVolumeType, 0)
					}
					airspacevolumetypes = append(airspacevolumetypes, airspacevolumetype)
					res[curvepropertytype_] = airspacevolumetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ContributorAirspace":
			res := make(map[*AirspaceVolumeDependencyPropertyType][]*AirspaceVolumeType)
			for airspacevolumetype := range stage.AirspaceVolumeTypes {
				if airspacevolumetype.ContributorAirspace != nil {
					airspacevolumedependencypropertytype_ := airspacevolumetype.ContributorAirspace
					var airspacevolumetypes []*AirspaceVolumeType
					_, ok := res[airspacevolumedependencypropertytype_]
					if ok {
						airspacevolumetypes = res[airspacevolumedependencypropertytype_]
					} else {
						airspacevolumetypes = make([]*AirspaceVolumeType, 0)
					}
					airspacevolumetypes = append(airspacevolumetypes, airspacevolumetype)
					res[airspacevolumedependencypropertytype_] = airspacevolumetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AuthorityForAirspaceTimeSlicePropertyType
	case AuthorityForAirspaceTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "AuthorityForAirspaceTimeSlice":
			res := make(map[*AuthorityForAirspaceTimeSliceType][]*AuthorityForAirspaceTimeSlicePropertyType)
			for authorityforairspacetimeslicepropertytype := range stage.AuthorityForAirspaceTimeSlicePropertyTypes {
				if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
					authorityforairspacetimeslicetype_ := authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice
					var authorityforairspacetimeslicepropertytypes []*AuthorityForAirspaceTimeSlicePropertyType
					_, ok := res[authorityforairspacetimeslicetype_]
					if ok {
						authorityforairspacetimeslicepropertytypes = res[authorityforairspacetimeslicetype_]
					} else {
						authorityforairspacetimeslicepropertytypes = make([]*AuthorityForAirspaceTimeSlicePropertyType, 0)
					}
					authorityforairspacetimeslicepropertytypes = append(authorityforairspacetimeslicepropertytypes, authorityforairspacetimeslicepropertytype)
					res[authorityforairspacetimeslicetype_] = authorityforairspacetimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AuthorityForAirspaceTimeSliceType
	case AuthorityForAirspaceTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "ResponsibleOrganisation":
			res := make(map[*OrganisationAuthorityPropertyType][]*AuthorityForAirspaceTimeSliceType)
			for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
				if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
					organisationauthoritypropertytype_ := authorityforairspacetimeslicetype.ResponsibleOrganisation
					var authorityforairspacetimeslicetypes []*AuthorityForAirspaceTimeSliceType
					_, ok := res[organisationauthoritypropertytype_]
					if ok {
						authorityforairspacetimeslicetypes = res[organisationauthoritypropertytype_]
					} else {
						authorityforairspacetimeslicetypes = make([]*AuthorityForAirspaceTimeSliceType, 0)
					}
					authorityforairspacetimeslicetypes = append(authorityforairspacetimeslicetypes, authorityforairspacetimeslicetype)
					res[organisationauthoritypropertytype_] = authorityforairspacetimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "AssignedAirspace":
			res := make(map[*AirspacePropertyType][]*AuthorityForAirspaceTimeSliceType)
			for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
				if authorityforairspacetimeslicetype.AssignedAirspace != nil {
					airspacepropertytype_ := authorityforairspacetimeslicetype.AssignedAirspace
					var authorityforairspacetimeslicetypes []*AuthorityForAirspaceTimeSliceType
					_, ok := res[airspacepropertytype_]
					if ok {
						authorityforairspacetimeslicetypes = res[airspacepropertytype_]
					} else {
						authorityforairspacetimeslicetypes = make([]*AuthorityForAirspaceTimeSliceType, 0)
					}
					authorityforairspacetimeslicetypes = append(authorityforairspacetimeslicetypes, authorityforairspacetimeslicetype)
					res[airspacepropertytype_] = authorityforairspacetimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AuthorityForAirspaceType
	case AuthorityForAirspaceType:
		switch fieldname {
		// insertion point for per direct association field
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
		case "ElevatedCurve":
			res := make(map[*ElevatedCurveType][]*CurvePropertyType)
			for curvepropertytype := range stage.CurvePropertyTypes {
				if curvepropertytype.ElevatedCurve != nil {
					elevatedcurvetype_ := curvepropertytype.ElevatedCurve
					var curvepropertytypes []*CurvePropertyType
					_, ok := res[elevatedcurvetype_]
					if ok {
						curvepropertytypes = res[elevatedcurvetype_]
					} else {
						curvepropertytypes = make([]*CurvePropertyType, 0)
					}
					curvepropertytypes = append(curvepropertytypes, curvepropertytype)
					res[elevatedcurvetype_] = curvepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurveType
	case CurveType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedCurveType
	case ElevatedCurveType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ElevatedSurfaceType
	case ElevatedSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoBorderTimeSlicePropertyType
	case GeoBorderTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		case "GeoBorderTimeSlice":
			res := make(map[*GeoBorderTimeSliceType][]*GeoBorderTimeSlicePropertyType)
			for geobordertimeslicepropertytype := range stage.GeoBorderTimeSlicePropertyTypes {
				if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
					geobordertimeslicetype_ := geobordertimeslicepropertytype.GeoBorderTimeSlice
					var geobordertimeslicepropertytypes []*GeoBorderTimeSlicePropertyType
					_, ok := res[geobordertimeslicetype_]
					if ok {
						geobordertimeslicepropertytypes = res[geobordertimeslicetype_]
					} else {
						geobordertimeslicepropertytypes = make([]*GeoBorderTimeSlicePropertyType, 0)
					}
					geobordertimeslicepropertytypes = append(geobordertimeslicepropertytypes, geobordertimeslicepropertytype)
					res[geobordertimeslicetype_] = geobordertimeslicepropertytypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GeoBorderTimeSliceType
	case GeoBorderTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Border":
			res := make(map[*CurvePropertyType][]*GeoBorderTimeSliceType)
			for geobordertimeslicetype := range stage.GeoBorderTimeSliceTypes {
				if geobordertimeslicetype.Border != nil {
					curvepropertytype_ := geobordertimeslicetype.Border
					var geobordertimeslicetypes []*GeoBorderTimeSliceType
					_, ok := res[curvepropertytype_]
					if ok {
						geobordertimeslicetypes = res[curvepropertytype_]
					} else {
						geobordertimeslicetypes = make([]*GeoBorderTimeSliceType, 0)
					}
					geobordertimeslicetypes = append(geobordertimeslicetypes, geobordertimeslicetype)
					res[curvepropertytype_] = geobordertimeslicetypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GeoBorderType
	case GeoBorderType:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of OrganisationAuthorityPropertyType
	case OrganisationAuthorityPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RoutePropertyType
	case RoutePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StandardLevelColumnPropertyType
	case StandardLevelColumnPropertyType:
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
	// reverse maps of direct associations of A_extension_13
	case A_extension_13:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_extension_14
	case A_extension_14:
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
	// reverse maps of direct associations of AircraftCharacteristicPropertyType
	case AircraftCharacteristicPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AircraftCharacteristicType
	case AircraftCharacteristicType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AircraftCharacteristicType)
			for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
				for _, notepropertytype_ := range aircraftcharacteristictype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], aircraftcharacteristictype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_12][]*AircraftCharacteristicType)
			for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
				for _, a_extension_12_ := range aircraftcharacteristictype.Extension {
					res[a_extension_12_] = append(res[a_extension_12_], aircraftcharacteristictype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceActivationPropertyType
	case AirspaceActivationPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceActivationType
	case AirspaceActivationType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInterval":
			res := make(map[*TimesheetPropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, timesheetpropertytype_ := range airspaceactivationtype.TimeInterval {
					res[timesheetpropertytype_] = append(res[timesheetpropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, notepropertytype_ := range airspaceactivationtype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpecialDateAuthority":
			res := make(map[*OrganisationAuthorityPropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, organisationauthoritypropertytype_ := range airspaceactivationtype.SpecialDateAuthority {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Levels":
			res := make(map[*AirspaceLayerPropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, airspacelayerpropertytype_ := range airspaceactivationtype.Levels {
					res[airspacelayerpropertytype_] = append(res[airspacelayerpropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "User":
			res := make(map[*OrganisationAuthorityPropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, organisationauthoritypropertytype_ := range airspaceactivationtype.User {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Aircraft":
			res := make(map[*AircraftCharacteristicPropertyType][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, aircraftcharacteristicpropertytype_ := range airspaceactivationtype.Aircraft {
					res[aircraftcharacteristicpropertytype_] = append(res[aircraftcharacteristicpropertytype_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_1][]*AirspaceActivationType)
			for airspaceactivationtype := range stage.AirspaceActivationTypes {
				for _, a_extension_1_ := range airspaceactivationtype.Extension {
					res[a_extension_1_] = append(res[a_extension_1_], airspaceactivationtype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceGeometryComponentPropertyType
	case AirspaceGeometryComponentPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceGeometryComponentType
	case AirspaceGeometryComponentType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceGeometryComponentType)
			for airspacegeometrycomponenttype := range stage.AirspaceGeometryComponentTypes {
				for _, notepropertytype_ := range airspacegeometrycomponenttype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacegeometrycomponenttype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_2][]*AirspaceGeometryComponentType)
			for airspacegeometrycomponenttype := range stage.AirspaceGeometryComponentTypes {
				for _, a_extension_2_ := range airspacegeometrycomponenttype.Extension {
					res[a_extension_2_] = append(res[a_extension_2_], airspacegeometrycomponenttype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceLayerClassPropertyType
	case AirspaceLayerClassPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceLayerClassType
	case AirspaceLayerClassType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeInterval":
			res := make(map[*TimesheetPropertyType][]*AirspaceLayerClassType)
			for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
				for _, timesheetpropertytype_ := range airspacelayerclasstype.TimeInterval {
					res[timesheetpropertytype_] = append(res[timesheetpropertytype_], airspacelayerclasstype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceLayerClassType)
			for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
				for _, notepropertytype_ := range airspacelayerclasstype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacelayerclasstype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SpecialDateAuthority":
			res := make(map[*OrganisationAuthorityPropertyType][]*AirspaceLayerClassType)
			for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
				for _, organisationauthoritypropertytype_ := range airspacelayerclasstype.SpecialDateAuthority {
					res[organisationauthoritypropertytype_] = append(res[organisationauthoritypropertytype_], airspacelayerclasstype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AssociatedLevels":
			res := make(map[*AirspaceLayerPropertyType][]*AirspaceLayerClassType)
			for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
				for _, airspacelayerpropertytype_ := range airspacelayerclasstype.AssociatedLevels {
					res[airspacelayerpropertytype_] = append(res[airspacelayerpropertytype_], airspacelayerclasstype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_3][]*AirspaceLayerClassType)
			for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
				for _, a_extension_3_ := range airspacelayerclasstype.Extension {
					res[a_extension_3_] = append(res[a_extension_3_], airspacelayerclasstype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceLayerPropertyType
	case AirspaceLayerPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceLayerType
	case AirspaceLayerType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceLayerType)
			for airspacelayertype := range stage.AirspaceLayerTypes {
				for _, notepropertytype_ := range airspacelayertype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacelayertype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_13][]*AirspaceLayerType)
			for airspacelayertype := range stage.AirspaceLayerTypes {
				for _, a_extension_13_ := range airspacelayertype.Extension {
					res[a_extension_13_] = append(res[a_extension_13_], airspacelayertype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspacePropertyType
	case AirspacePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceTimeSlicePropertyType
	case AirspaceTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceTimeSliceType
	case AirspaceTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Class":
			res := make(map[*AirspaceLayerClassPropertyType][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				for _, airspacelayerclasspropertytype_ := range airspacetimeslicetype.Class {
					res[airspacelayerclasspropertytype_] = append(res[airspacelayerclasspropertytype_], airspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GeometryComponent":
			res := make(map[*AirspaceGeometryComponentPropertyType][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				for _, airspacegeometrycomponentpropertytype_ := range airspacetimeslicetype.GeometryComponent {
					res[airspacegeometrycomponentpropertytype_] = append(res[airspacegeometrycomponentpropertytype_], airspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Activation":
			res := make(map[*AirspaceActivationPropertyType][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				for _, airspaceactivationpropertytype_ := range airspacetimeslicetype.Activation {
					res[airspaceactivationpropertytype_] = append(res[airspaceactivationpropertytype_], airspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				for _, notepropertytype_ := range airspacetimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension][]*AirspaceTimeSliceType)
			for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
				for _, a_extension_ := range airspacetimeslicetype.Extension {
					res[a_extension_] = append(res[a_extension_], airspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceType
	case AirspaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*AirspaceTimeSlicePropertyType][]*AirspaceType)
			for airspacetype := range stage.AirspaceTypes {
				for _, airspacetimeslicepropertytype_ := range airspacetype.TimeSlice {
					res[airspacetimeslicepropertytype_] = append(res[airspacetimeslicepropertytype_], airspacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceVolumeDependencyPropertyType
	case AirspaceVolumeDependencyPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceVolumeDependencyType
	case AirspaceVolumeDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceVolumeDependencyType)
			for airspacevolumedependencytype := range stage.AirspaceVolumeDependencyTypes {
				for _, notepropertytype_ := range airspacevolumedependencytype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacevolumedependencytype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_5][]*AirspaceVolumeDependencyType)
			for airspacevolumedependencytype := range stage.AirspaceVolumeDependencyTypes {
				for _, a_extension_5_ := range airspacevolumedependencytype.Extension {
					res[a_extension_5_] = append(res[a_extension_5_], airspacevolumedependencytype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AirspaceVolumePropertyType
	case AirspaceVolumePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AirspaceVolumeType
	case AirspaceVolumeType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AirspaceVolumeType)
			for airspacevolumetype := range stage.AirspaceVolumeTypes {
				for _, notepropertytype_ := range airspacevolumetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], airspacevolumetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_4][]*AirspaceVolumeType)
			for airspacevolumetype := range stage.AirspaceVolumeTypes {
				for _, a_extension_4_ := range airspacevolumetype.Extension {
					res[a_extension_4_] = append(res[a_extension_4_], airspacevolumetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AuthorityForAirspaceTimeSlicePropertyType
	case AuthorityForAirspaceTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AuthorityForAirspaceTimeSliceType
	case AuthorityForAirspaceTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*AuthorityForAirspaceTimeSliceType)
			for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
				for _, notepropertytype_ := range authorityforairspacetimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], authorityforairspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_6][]*AuthorityForAirspaceTimeSliceType)
			for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
				for _, a_extension_6_ := range authorityforairspacetimeslicetype.Extension {
					res[a_extension_6_] = append(res[a_extension_6_], authorityforairspacetimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AuthorityForAirspaceType
	case AuthorityForAirspaceType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*AuthorityForAirspaceTimeSlicePropertyType][]*AuthorityForAirspaceType)
			for authorityforairspacetype := range stage.AuthorityForAirspaceTypes {
				for _, authorityforairspacetimeslicepropertytype_ := range authorityforairspacetype.TimeSlice {
					res[authorityforairspacetimeslicepropertytype_] = append(res[authorityforairspacetimeslicepropertytype_], authorityforairspacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CurvePropertyType
	case CurvePropertyType:
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
			res := make(map[*A_extension_8][]*ElevatedCurveType)
			for elevatedcurvetype := range stage.ElevatedCurveTypes {
				for _, a_extension_8_ := range elevatedcurvetype.Extension {
					res[a_extension_8_] = append(res[a_extension_8_], elevatedcurvetype)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*A_extension_9][]*ElevatedSurfaceType)
			for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
				for _, a_extension_9_ := range elevatedsurfacetype.Extension {
					res[a_extension_9_] = append(res[a_extension_9_], elevatedsurfacetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GeoBorderTimeSlicePropertyType
	case GeoBorderTimeSlicePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GeoBorderTimeSliceType
	case GeoBorderTimeSliceType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*NotePropertyType][]*GeoBorderTimeSliceType)
			for geobordertimeslicetype := range stage.GeoBorderTimeSliceTypes {
				for _, notepropertytype_ := range geobordertimeslicetype.Annotation {
					res[notepropertytype_] = append(res[notepropertytype_], geobordertimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*A_extension_7][]*GeoBorderTimeSliceType)
			for geobordertimeslicetype := range stage.GeoBorderTimeSliceTypes {
				for _, a_extension_7_ := range geobordertimeslicetype.Extension {
					res[a_extension_7_] = append(res[a_extension_7_], geobordertimeslicetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GeoBorderType
	case GeoBorderType:
		switch fieldname {
		// insertion point for per direct association field
		case "TimeSlice":
			res := make(map[*GeoBorderTimeSlicePropertyType][]*GeoBorderType)
			for geobordertype := range stage.GeoBorderTypes {
				for _, geobordertimeslicepropertytype_ := range geobordertype.TimeSlice {
					res[geobordertimeslicepropertytype_] = append(res[geobordertimeslicepropertytype_], geobordertype)
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
			res := make(map[*A_extension_10][]*LinguisticNoteType)
			for linguisticnotetype := range stage.LinguisticNoteTypes {
				for _, a_extension_10_ := range linguisticnotetype.Extension {
					res[a_extension_10_] = append(res[a_extension_10_], linguisticnotetype)
				}
			}
			return any(res).(map[*End][]*Start)
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
			res := make(map[*A_extension_11][]*NoteType)
			for notetype := range stage.NoteTypes {
				for _, a_extension_11_ := range notetype.Extension {
					res[a_extension_11_] = append(res[a_extension_11_], notetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of OrganisationAuthorityPropertyType
	case OrganisationAuthorityPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RoutePropertyType
	case RoutePropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StandardLevelColumnPropertyType
	case StandardLevelColumnPropertyType:
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
			res := make(map[*A_extension_14][]*TimesheetType)
			for timesheettype := range stage.TimesheetTypes {
				for _, a_extension_14_ := range timesheettype.Extension {
					res[a_extension_14_] = append(res[a_extension_14_], timesheettype)
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
			GongstructName: "AirspaceTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_1) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_10) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "LinguisticNoteType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_11) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "NoteType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_12) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AircraftCharacteristicType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_13) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceLayerType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_14) GongGetReverseFields() []GongReverseField {
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
			GongstructName: "AirspaceGeometryComponentType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_3) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceLayerClassType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_4) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceVolumeType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_5) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceVolumeDependencyType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_6) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AuthorityForAirspaceTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_7) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "GeoBorderTimeSliceType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_8) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ElevatedCurveType",
			Fieldname: "Extension",
		},
	}
}

func (*A_extension_9) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "ElevatedSurfaceType",
			Fieldname: "Extension",
		},
	}
}

func (*AbstractPropertiesWithScheduleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AircraftCharacteristicPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "Aircraft",
		},
	}
}

func (*AircraftCharacteristicType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceActivationPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceTimeSliceType",
			Fieldname: "Activation",
		},
	}
}

func (*AirspaceActivationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceGeometryComponentPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceTimeSliceType",
			Fieldname: "GeometryComponent",
		},
	}
}

func (*AirspaceGeometryComponentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceLayerClassPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceTimeSliceType",
			Fieldname: "Class",
		},
	}
}

func (*AirspaceLayerClassType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceLayerPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "Levels",
		},
		{
			GongstructName: "AirspaceLayerClassType",
			Fieldname: "AssociatedLevels",
		},
	}
}

func (*AirspaceLayerType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspacePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*AirspaceTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceVolumeDependencyPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceVolumeDependencyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceVolumePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AirspaceVolumeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AuthorityForAirspaceTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AuthorityForAirspaceType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*AuthorityForAirspaceTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AuthorityForAirspaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CurvePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedCurveType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ElevatedSurfaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GeoBorderTimeSlicePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "GeoBorderType",
			Fieldname: "TimeSlice",
		},
	}
}

func (*GeoBorderTimeSliceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*GeoBorderType) GongGetReverseFields() []GongReverseField {
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

func (*NotePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AircraftCharacteristicType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceGeometryComponentType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceLayerClassType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceLayerType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceTimeSliceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceVolumeDependencyType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AirspaceVolumeType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "AuthorityForAirspaceTimeSliceType",
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
			GongstructName: "ElevatedSurfaceType",
			Fieldname: "Annotation",
		},
		{
			GongstructName: "GeoBorderTimeSliceType",
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
	}
}

func (*NoteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OrganisationAuthorityPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "SpecialDateAuthority",
		},
		{
			GongstructName: "AirspaceActivationType",
			Fieldname: "User",
		},
		{
			GongstructName: "AirspaceLayerClassType",
			Fieldname: "SpecialDateAuthority",
		},
	}
}

func (*RoutePropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StandardLevelColumnPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
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
			GongstructName: "AirspaceActivationType",
			Fieldname: "TimeInterval",
		},
		{
			GongstructName: "AirspaceLayerClassType",
			Fieldname: "TimeInterval",
		},
	}
}

func (*TimesheetType) GongGetReverseFields() []GongReverseField {
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

func (a_extension_13 *A_extension_13) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (a_extension_14 *A_extension_14) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AircraftCharacteristic",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AircraftCharacteristicType",
		},
	}
	return
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceActivation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceActivationType",
		},
	}
	return
}

func (airspaceactivationtype *AirspaceActivationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Levels",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AirspaceLayerPropertyType",
		},
		{
			Name:                 "User",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "Aircraft",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AircraftCharacteristicPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_1",
		},
	}
	return
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceGeometryComponent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceGeometryComponentType",
		},
	}
	return
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TheAirspaceVolume",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceVolumePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_2",
		},
	}
	return
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceLayerClass",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceLayerClassType",
		},
	}
	return
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AssociatedLevels",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AirspaceLayerPropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_3",
		},
	}
	return
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceLayer",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceLayerType",
		},
	}
	return
}

func (airspacelayertype *AirspaceLayerType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DiscreteLevelSeries",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StandardLevelColumnPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_13",
		},
	}
	return
}

func (airspacepropertytype *AirspacePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceTimeSliceType",
		},
	}
	return
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Class",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AirspaceLayerClassPropertyType",
		},
		{
			Name:                 "ProtectedRoute",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RoutePropertyType",
		},
		{
			Name:                 "GeometryComponent",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AirspaceGeometryComponentPropertyType",
		},
		{
			Name:                 "Activation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AirspaceActivationPropertyType",
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

func (airspacetype *AirspaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			TargetGongstructName: "AirspaceTimeSlicePropertyType",
		},
	}
	return
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceVolumeDependency",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceVolumeDependencyType",
		},
	}
	return
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "TheAirspace",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspacePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_5",
		},
	}
	return
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AirspaceVolume",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceVolumeType",
		},
	}
	return
}

func (airspacevolumetype *AirspaceVolumeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "HorizontalProjection",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SurfacePropertyType",
		},
		{
			Name:                 "Centreline",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurvePropertyType",
		},
		{
			Name:                 "ContributorAirspace",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspaceVolumeDependencyPropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_4",
		},
	}
	return
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "AuthorityForAirspaceTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AuthorityForAirspaceTimeSliceType",
		},
	}
	return
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ResponsibleOrganisation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "OrganisationAuthorityPropertyType",
		},
		{
			Name:                 "AssignedAirspace",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AirspacePropertyType",
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

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			TargetGongstructName: "AuthorityForAirspaceTimeSlicePropertyType",
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
			Name:                 "Curve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurveType",
		},
		{
			Name:                 "ElevatedCurve",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ElevatedCurveType",
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
			TargetGongstructName: "A_extension_8",
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
			TargetGongstructName: "A_extension_9",
		},
	}
	return
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "GeoBorderTimeSlice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GeoBorderTimeSliceType",
		},
	}
	return
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Border",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CurvePropertyType",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePropertyType",
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "A_extension_7",
		},
	}
	return
}

func (geobordertype *GeoBorderType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			TargetGongstructName: "GeoBorderTimeSlicePropertyType",
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
			TargetGongstructName: "A_extension_10",
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
			TargetGongstructName: "A_extension_11",
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

func (routepropertytype *RoutePropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			TargetGongstructName: "A_extension_14",
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

func (a_extension_13 *A_extension_13) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_13.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_13.Owns)
		res.valueBool = a_extension_13.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (a_extension_14 *A_extension_14) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a_extension_14.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", a_extension_14.Owns)
		res.valueBool = a_extension_14.Owns
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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aircraftcharacteristicpropertytype.Name
	case "NilReason":
		res.valueString = aircraftcharacteristicpropertytype.NilReason
	case "AircraftCharacteristic":
		res.GongFieldValueType = GongFieldValueTypePointer
		if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
			res.valueString = aircraftcharacteristicpropertytype.AircraftCharacteristic.Name
			res.ids = aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetUUID(stage)
		}
	}
	return
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aircraftcharacteristictype.Name
	case "Id":
		res.valueString = aircraftcharacteristictype.Id
	case "Description":
		res.valueString = aircraftcharacteristictype.Description
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aircraftcharacteristictype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aircraftcharacteristictype.Extension {
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

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspaceactivationpropertytype.Name
	case "NilReason":
		res.valueString = airspaceactivationpropertytype.NilReason
	case "AirspaceActivation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspaceactivationpropertytype.AirspaceActivation != nil {
			res.valueString = airspaceactivationpropertytype.AirspaceActivation.Name
			res.ids = airspaceactivationpropertytype.AirspaceActivation.GongGetUUID(stage)
		}
	}
	return
}

func (airspaceactivationtype *AirspaceActivationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspaceactivationtype.Name
	case "Id":
		res.valueString = airspaceactivationtype.Id
	case "Description":
		res.valueString = airspaceactivationtype.Description
	case "TimeInterval":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.TimeInterval {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpecialDateAuthority":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.SpecialDateAuthority {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Levels":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.Levels {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "User":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.User {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Aircraft":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.Aircraft {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspaceactivationtype.Extension {
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

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacegeometrycomponentpropertytype.Name
	case "NilReason":
		res.valueString = airspacegeometrycomponentpropertytype.NilReason
	case "AirspaceGeometryComponent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
			res.valueString = airspacegeometrycomponentpropertytype.AirspaceGeometryComponent.Name
			res.ids = airspacegeometrycomponentpropertytype.AirspaceGeometryComponent.GongGetUUID(stage)
		}
	}
	return
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacegeometrycomponenttype.Name
	case "Id":
		res.valueString = airspacegeometrycomponenttype.Id
	case "Description":
		res.valueString = airspacegeometrycomponenttype.Description
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacegeometrycomponenttype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TheAirspaceVolume":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
			res.valueString = airspacegeometrycomponenttype.TheAirspaceVolume.Name
			res.ids = airspacegeometrycomponenttype.TheAirspaceVolume.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacegeometrycomponenttype.Extension {
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

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacelayerclasspropertytype.Name
	case "NilReason":
		res.valueString = airspacelayerclasspropertytype.NilReason
	case "AirspaceLayerClass":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
			res.valueString = airspacelayerclasspropertytype.AirspaceLayerClass.Name
			res.ids = airspacelayerclasspropertytype.AirspaceLayerClass.GongGetUUID(stage)
		}
	}
	return
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacelayerclasstype.Name
	case "Id":
		res.valueString = airspacelayerclasstype.Id
	case "Description":
		res.valueString = airspacelayerclasstype.Description
	case "TimeInterval":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayerclasstype.TimeInterval {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayerclasstype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SpecialDateAuthority":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayerclasstype.SpecialDateAuthority {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AssociatedLevels":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayerclasstype.AssociatedLevels {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayerclasstype.Extension {
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

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacelayerpropertytype.Name
	case "NilReason":
		res.valueString = airspacelayerpropertytype.NilReason
	case "AirspaceLayer":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacelayerpropertytype.AirspaceLayer != nil {
			res.valueString = airspacelayerpropertytype.AirspaceLayer.Name
			res.ids = airspacelayerpropertytype.AirspaceLayer.GongGetUUID(stage)
		}
	}
	return
}

func (airspacelayertype *AirspaceLayerType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacelayertype.Name
	case "Id":
		res.valueString = airspacelayertype.Id
	case "Description":
		res.valueString = airspacelayertype.Description
	case "DiscreteLevelSeries":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacelayertype.DiscreteLevelSeries != nil {
			res.valueString = airspacelayertype.DiscreteLevelSeries.Name
			res.ids = airspacelayertype.DiscreteLevelSeries.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayertype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacelayertype.Extension {
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

func (airspacepropertytype *AirspacePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", airspacepropertytype.Owns)
		res.valueBool = airspacepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = airspacepropertytype.NilReason
	case "Type_XSD":
		res.valueString = airspacepropertytype.Type_XSD
	case "Href":
		res.valueString = airspacepropertytype.Href
	case "Role":
		res.valueString = airspacepropertytype.Role
	case "Arcrole":
		res.valueString = airspacepropertytype.Arcrole
	case "Title":
		res.valueString = airspacepropertytype.Title
	case "Show":
		res.valueString = airspacepropertytype.Show
	case "Actuate":
		res.valueString = airspacepropertytype.Actuate
	}
	return
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacetimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", airspacetimeslicepropertytype.Owns)
		res.valueBool = airspacetimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AirspaceTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
			res.valueString = airspacetimeslicepropertytype.AirspaceTimeSlice.Name
			res.ids = airspacetimeslicepropertytype.AirspaceTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacetimeslicetype.Name
	case "Id":
		res.valueString = airspacetimeslicetype.Id
	case "Description":
		res.valueString = airspacetimeslicetype.Description
	case "Interpretation":
		res.valueString = airspacetimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", airspacetimeslicetype.SequenceNumber)
		res.valueInt = airspacetimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", airspacetimeslicetype.CorrectionNumber)
		res.valueInt = airspacetimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Class":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetimeslicetype.Class {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ProtectedRoute":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacetimeslicetype.ProtectedRoute != nil {
			res.valueString = airspacetimeslicetype.ProtectedRoute.Name
			res.ids = airspacetimeslicetype.ProtectedRoute.GongGetUUID(stage)
		}
	case "GeometryComponent":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetimeslicetype.GeometryComponent {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Activation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetimeslicetype.Activation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetimeslicetype.Extension {
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

func (airspacetype *AirspaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacetype.Name
	case "Id":
		res.valueString = airspacetype.Id
	case "Description":
		res.valueString = airspacetype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacetype.TimeSlice {
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

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacevolumedependencypropertytype.Name
	case "NilReason":
		res.valueString = airspacevolumedependencypropertytype.NilReason
	case "AirspaceVolumeDependency":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
			res.valueString = airspacevolumedependencypropertytype.AirspaceVolumeDependency.Name
			res.ids = airspacevolumedependencypropertytype.AirspaceVolumeDependency.GongGetUUID(stage)
		}
	}
	return
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacevolumedependencytype.Name
	case "Id":
		res.valueString = airspacevolumedependencytype.Id
	case "Description":
		res.valueString = airspacevolumedependencytype.Description
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacevolumedependencytype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TheAirspace":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumedependencytype.TheAirspace != nil {
			res.valueString = airspacevolumedependencytype.TheAirspace.Name
			res.ids = airspacevolumedependencytype.TheAirspace.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacevolumedependencytype.Extension {
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

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacevolumepropertytype.Name
	case "NilReason":
		res.valueString = airspacevolumepropertytype.NilReason
	case "AirspaceVolume":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumepropertytype.AirspaceVolume != nil {
			res.valueString = airspacevolumepropertytype.AirspaceVolume.Name
			res.ids = airspacevolumepropertytype.AirspaceVolume.GongGetUUID(stage)
		}
	}
	return
}

func (airspacevolumetype *AirspaceVolumeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = airspacevolumetype.Name
	case "Id":
		res.valueString = airspacevolumetype.Id
	case "Description":
		res.valueString = airspacevolumetype.Description
	case "HorizontalProjection":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumetype.HorizontalProjection != nil {
			res.valueString = airspacevolumetype.HorizontalProjection.Name
			res.ids = airspacevolumetype.HorizontalProjection.GongGetUUID(stage)
		}
	case "Centreline":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumetype.Centreline != nil {
			res.valueString = airspacevolumetype.Centreline.Name
			res.ids = airspacevolumetype.Centreline.GongGetUUID(stage)
		}
	case "ContributorAirspace":
		res.GongFieldValueType = GongFieldValueTypePointer
		if airspacevolumetype.ContributorAirspace != nil {
			res.valueString = airspacevolumetype.ContributorAirspace.Name
			res.ids = airspacevolumetype.ContributorAirspace.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacevolumetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range airspacevolumetype.Extension {
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

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = authorityforairspacetimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", authorityforairspacetimeslicepropertytype.Owns)
		res.valueBool = authorityforairspacetimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AuthorityForAirspaceTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
			res.valueString = authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice.Name
			res.ids = authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = authorityforairspacetimeslicetype.Name
	case "Id":
		res.valueString = authorityforairspacetimeslicetype.Id
	case "Description":
		res.valueString = authorityforairspacetimeslicetype.Description
	case "Interpretation":
		res.valueString = authorityforairspacetimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", authorityforairspacetimeslicetype.SequenceNumber)
		res.valueInt = authorityforairspacetimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", authorityforairspacetimeslicetype.CorrectionNumber)
		res.valueInt = authorityforairspacetimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ResponsibleOrganisation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
			res.valueString = authorityforairspacetimeslicetype.ResponsibleOrganisation.Name
			res.ids = authorityforairspacetimeslicetype.ResponsibleOrganisation.GongGetUUID(stage)
		}
	case "AssignedAirspace":
		res.GongFieldValueType = GongFieldValueTypePointer
		if authorityforairspacetimeslicetype.AssignedAirspace != nil {
			res.valueString = authorityforairspacetimeslicetype.AssignedAirspace.Name
			res.ids = authorityforairspacetimeslicetype.AssignedAirspace.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range authorityforairspacetimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range authorityforairspacetimeslicetype.Extension {
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

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = authorityforairspacetype.Name
	case "Id":
		res.valueString = authorityforairspacetype.Id
	case "Description":
		res.valueString = authorityforairspacetype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range authorityforairspacetype.TimeSlice {
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
	case "Curve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvepropertytype.Curve != nil {
			res.valueString = curvepropertytype.Curve.Name
			res.ids = curvepropertytype.Curve.GongGetUUID(stage)
		}
	case "ElevatedCurve":
		res.GongFieldValueType = GongFieldValueTypePointer
		if curvepropertytype.ElevatedCurve != nil {
			res.valueString = curvepropertytype.ElevatedCurve.Name
			res.ids = curvepropertytype.ElevatedCurve.GongGetUUID(stage)
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

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geobordertimeslicepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", geobordertimeslicepropertytype.Owns)
		res.valueBool = geobordertimeslicepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GeoBorderTimeSlice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
			res.valueString = geobordertimeslicepropertytype.GeoBorderTimeSlice.Name
			res.ids = geobordertimeslicepropertytype.GeoBorderTimeSlice.GongGetUUID(stage)
		}
	}
	return
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geobordertimeslicetype.Name
	case "Id":
		res.valueString = geobordertimeslicetype.Id
	case "Description":
		res.valueString = geobordertimeslicetype.Description
	case "Interpretation":
		res.valueString = geobordertimeslicetype.Interpretation
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", geobordertimeslicetype.SequenceNumber)
		res.valueInt = geobordertimeslicetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CorrectionNumber":
		res.valueString = fmt.Sprintf("%d", geobordertimeslicetype.CorrectionNumber)
		res.valueInt = geobordertimeslicetype.CorrectionNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Border":
		res.GongFieldValueType = GongFieldValueTypePointer
		if geobordertimeslicetype.Border != nil {
			res.valueString = geobordertimeslicetype.Border.Name
			res.ids = geobordertimeslicetype.Border.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range geobordertimeslicetype.Annotation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range geobordertimeslicetype.Extension {
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

func (geobordertype *GeoBorderType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = geobordertype.Name
	case "Id":
		res.valueString = geobordertype.Id
	case "Description":
		res.valueString = geobordertype.Description
	case "TimeSlice":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range geobordertype.TimeSlice {
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

func (routepropertytype *RoutePropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = routepropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", routepropertytype.Owns)
		res.valueBool = routepropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = routepropertytype.NilReason
	case "Type_XSD":
		res.valueString = routepropertytype.Type_XSD
	case "Href":
		res.valueString = routepropertytype.Href
	case "Role":
		res.valueString = routepropertytype.Role
	case "Arcrole":
		res.valueString = routepropertytype.Arcrole
	case "Title":
		res.valueString = routepropertytype.Title
	case "Show":
		res.valueString = routepropertytype.Show
	case "Actuate":
		res.valueString = routepropertytype.Actuate
	}
	return
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = standardlevelcolumnpropertytype.Name
	case "Owns":
		res.valueString = fmt.Sprintf("%t", standardlevelcolumnpropertytype.Owns)
		res.valueBool = standardlevelcolumnpropertytype.Owns
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NilReason":
		res.valueString = standardlevelcolumnpropertytype.NilReason
	case "Type_XSD":
		res.valueString = standardlevelcolumnpropertytype.Type_XSD
	case "Href":
		res.valueString = standardlevelcolumnpropertytype.Href
	case "Role":
		res.valueString = standardlevelcolumnpropertytype.Role
	case "Arcrole":
		res.valueString = standardlevelcolumnpropertytype.Arcrole
	case "Title":
		res.valueString = standardlevelcolumnpropertytype.Title
	case "Show":
		res.valueString = standardlevelcolumnpropertytype.Show
	case "Actuate":
		res.valueString = standardlevelcolumnpropertytype.Actuate
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

func (a_extension_13 *A_extension_13) GongGetGongstructName() string {
	return "A_extension_13"
}

func (a_extension_14 *A_extension_14) GongGetGongstructName() string {
	return "A_extension_14"
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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetGongstructName() string {
	return "AircraftCharacteristicPropertyType"
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetGongstructName() string {
	return "AircraftCharacteristicType"
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetGongstructName() string {
	return "AirspaceActivationPropertyType"
}

func (airspaceactivationtype *AirspaceActivationType) GongGetGongstructName() string {
	return "AirspaceActivationType"
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetGongstructName() string {
	return "AirspaceGeometryComponentPropertyType"
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetGongstructName() string {
	return "AirspaceGeometryComponentType"
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetGongstructName() string {
	return "AirspaceLayerClassPropertyType"
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetGongstructName() string {
	return "AirspaceLayerClassType"
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetGongstructName() string {
	return "AirspaceLayerPropertyType"
}

func (airspacelayertype *AirspaceLayerType) GongGetGongstructName() string {
	return "AirspaceLayerType"
}

func (airspacepropertytype *AirspacePropertyType) GongGetGongstructName() string {
	return "AirspacePropertyType"
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetGongstructName() string {
	return "AirspaceTimeSlicePropertyType"
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetGongstructName() string {
	return "AirspaceTimeSliceType"
}

func (airspacetype *AirspaceType) GongGetGongstructName() string {
	return "AirspaceType"
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetGongstructName() string {
	return "AirspaceVolumeDependencyPropertyType"
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetGongstructName() string {
	return "AirspaceVolumeDependencyType"
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetGongstructName() string {
	return "AirspaceVolumePropertyType"
}

func (airspacevolumetype *AirspaceVolumeType) GongGetGongstructName() string {
	return "AirspaceVolumeType"
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetGongstructName() string {
	return "AuthorityForAirspaceTimeSlicePropertyType"
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetGongstructName() string {
	return "AuthorityForAirspaceTimeSliceType"
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetGongstructName() string {
	return "AuthorityForAirspaceType"
}

func (curvepropertytype *CurvePropertyType) GongGetGongstructName() string {
	return "CurvePropertyType"
}

func (curvetype *CurveType) GongGetGongstructName() string {
	return "CurveType"
}

func (elevatedcurvetype *ElevatedCurveType) GongGetGongstructName() string {
	return "ElevatedCurveType"
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetGongstructName() string {
	return "ElevatedSurfaceType"
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetGongstructName() string {
	return "GeoBorderTimeSlicePropertyType"
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetGongstructName() string {
	return "GeoBorderTimeSliceType"
}

func (geobordertype *GeoBorderType) GongGetGongstructName() string {
	return "GeoBorderType"
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetGongstructName() string {
	return "LinguisticNotePropertyType"
}

func (linguisticnotetype *LinguisticNoteType) GongGetGongstructName() string {
	return "LinguisticNoteType"
}

func (notepropertytype *NotePropertyType) GongGetGongstructName() string {
	return "NotePropertyType"
}

func (notetype *NoteType) GongGetGongstructName() string {
	return "NoteType"
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetGongstructName() string {
	return "OrganisationAuthorityPropertyType"
}

func (routepropertytype *RoutePropertyType) GongGetGongstructName() string {
	return "RoutePropertyType"
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetGongstructName() string {
	return "StandardLevelColumnPropertyType"
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

	__gong__rebuildMapString(stage.A_extension_13s, &stage.A_extension_13s_mapString)

	__gong__rebuildMapString(stage.A_extension_14s, &stage.A_extension_14s_mapString)

	__gong__rebuildMapString(stage.A_extension_2s, &stage.A_extension_2s_mapString)

	__gong__rebuildMapString(stage.A_extension_3s, &stage.A_extension_3s_mapString)

	__gong__rebuildMapString(stage.A_extension_4s, &stage.A_extension_4s_mapString)

	__gong__rebuildMapString(stage.A_extension_5s, &stage.A_extension_5s_mapString)

	__gong__rebuildMapString(stage.A_extension_6s, &stage.A_extension_6s_mapString)

	__gong__rebuildMapString(stage.A_extension_7s, &stage.A_extension_7s_mapString)

	__gong__rebuildMapString(stage.A_extension_8s, &stage.A_extension_8s_mapString)

	__gong__rebuildMapString(stage.A_extension_9s, &stage.A_extension_9s_mapString)

	__gong__rebuildMapString(stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_mapString)

	__gong__rebuildMapString(stage.AircraftCharacteristicPropertyTypes, &stage.AircraftCharacteristicPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AircraftCharacteristicTypes, &stage.AircraftCharacteristicTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceActivationPropertyTypes, &stage.AirspaceActivationPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceActivationTypes, &stage.AirspaceActivationTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceGeometryComponentPropertyTypes, &stage.AirspaceGeometryComponentPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceGeometryComponentTypes, &stage.AirspaceGeometryComponentTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceLayerClassPropertyTypes, &stage.AirspaceLayerClassPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceLayerClassTypes, &stage.AirspaceLayerClassTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceLayerPropertyTypes, &stage.AirspaceLayerPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceLayerTypes, &stage.AirspaceLayerTypes_mapString)

	__gong__rebuildMapString(stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceTimeSlicePropertyTypes, &stage.AirspaceTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceTimeSliceTypes, &stage.AirspaceTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceTypes, &stage.AirspaceTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceVolumeDependencyPropertyTypes, &stage.AirspaceVolumeDependencyPropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceVolumeDependencyTypes, &stage.AirspaceVolumeDependencyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceVolumePropertyTypes, &stage.AirspaceVolumePropertyTypes_mapString)

	__gong__rebuildMapString(stage.AirspaceVolumeTypes, &stage.AirspaceVolumeTypes_mapString)

	__gong__rebuildMapString(stage.AuthorityForAirspaceTimeSlicePropertyTypes, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.AuthorityForAirspaceTimeSliceTypes, &stage.AuthorityForAirspaceTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.AuthorityForAirspaceTypes, &stage.AuthorityForAirspaceTypes_mapString)

	__gong__rebuildMapString(stage.CurvePropertyTypes, &stage.CurvePropertyTypes_mapString)

	__gong__rebuildMapString(stage.CurveTypes, &stage.CurveTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_mapString)

	__gong__rebuildMapString(stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.GeoBorderTimeSlicePropertyTypes, &stage.GeoBorderTimeSlicePropertyTypes_mapString)

	__gong__rebuildMapString(stage.GeoBorderTimeSliceTypes, &stage.GeoBorderTimeSliceTypes_mapString)

	__gong__rebuildMapString(stage.GeoBorderTypes, &stage.GeoBorderTypes_mapString)

	__gong__rebuildMapString(stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_mapString)

	__gong__rebuildMapString(stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_mapString)

	__gong__rebuildMapString(stage.NotePropertyTypes, &stage.NotePropertyTypes_mapString)

	__gong__rebuildMapString(stage.NoteTypes, &stage.NoteTypes_mapString)

	__gong__rebuildMapString(stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_mapString)

	__gong__rebuildMapString(stage.RoutePropertyTypes, &stage.RoutePropertyTypes_mapString)

	__gong__rebuildMapString(stage.StandardLevelColumnPropertyTypes, &stage.StandardLevelColumnPropertyTypes_mapString)

	__gong__rebuildMapString(stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_mapString)

	__gong__rebuildMapString(stage.SurfaceTypes, &stage.SurfaceTypes_mapString)

	__gong__rebuildMapString(stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_mapString)

	__gong__rebuildMapString(stage.TimesheetTypes, &stage.TimesheetTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
