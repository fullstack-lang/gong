// generated code - do not edit
package aixmdatatypes

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
	CodeATCReportingTypes                map[*CodeATCReportingType]struct{}
	CodeATCReportingTypes_instance       map[*CodeATCReportingType]*CodeATCReportingType
	CodeATCReportingTypes_mapString      map[string]*CodeATCReportingType
	CodeATCReportingTypeOrder            uint
	CodeATCReportingType_stagedOrder     map[*CodeATCReportingType]uint
	CodeATCReportingType_orderStaged     map[uint]*CodeATCReportingType
	CodeATCReportingTypes_reference      map[*CodeATCReportingType]*CodeATCReportingType
	CodeATCReportingTypes_referenceOrder map[*CodeATCReportingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeATCReportingTypeCreateCallback GongOnAfterCreateInterface[CodeATCReportingType]
	OnAfterCodeATCReportingTypeUpdateCallback GongOnAfterUpdateInterface[CodeATCReportingType]
	OnAfterCodeATCReportingTypeDeleteCallback GongOnAfterDeleteInterface[CodeATCReportingType]

	CodeAerialRefuellingPointTypes                map[*CodeAerialRefuellingPointType]struct{}
	CodeAerialRefuellingPointTypes_instance       map[*CodeAerialRefuellingPointType]*CodeAerialRefuellingPointType
	CodeAerialRefuellingPointTypes_mapString      map[string]*CodeAerialRefuellingPointType
	CodeAerialRefuellingPointTypeOrder            uint
	CodeAerialRefuellingPointType_stagedOrder     map[*CodeAerialRefuellingPointType]uint
	CodeAerialRefuellingPointType_orderStaged     map[uint]*CodeAerialRefuellingPointType
	CodeAerialRefuellingPointTypes_reference      map[*CodeAerialRefuellingPointType]*CodeAerialRefuellingPointType
	CodeAerialRefuellingPointTypes_referenceOrder map[*CodeAerialRefuellingPointType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAerialRefuellingPointTypeCreateCallback GongOnAfterCreateInterface[CodeAerialRefuellingPointType]
	OnAfterCodeAerialRefuellingPointTypeUpdateCallback GongOnAfterUpdateInterface[CodeAerialRefuellingPointType]
	OnAfterCodeAerialRefuellingPointTypeDeleteCallback GongOnAfterDeleteInterface[CodeAerialRefuellingPointType]

	CodeAerialRefuellingPrefixTypes                map[*CodeAerialRefuellingPrefixType]struct{}
	CodeAerialRefuellingPrefixTypes_instance       map[*CodeAerialRefuellingPrefixType]*CodeAerialRefuellingPrefixType
	CodeAerialRefuellingPrefixTypes_mapString      map[string]*CodeAerialRefuellingPrefixType
	CodeAerialRefuellingPrefixTypeOrder            uint
	CodeAerialRefuellingPrefixType_stagedOrder     map[*CodeAerialRefuellingPrefixType]uint
	CodeAerialRefuellingPrefixType_orderStaged     map[uint]*CodeAerialRefuellingPrefixType
	CodeAerialRefuellingPrefixTypes_reference      map[*CodeAerialRefuellingPrefixType]*CodeAerialRefuellingPrefixType
	CodeAerialRefuellingPrefixTypes_referenceOrder map[*CodeAerialRefuellingPrefixType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAerialRefuellingPrefixTypeCreateCallback GongOnAfterCreateInterface[CodeAerialRefuellingPrefixType]
	OnAfterCodeAerialRefuellingPrefixTypeUpdateCallback GongOnAfterUpdateInterface[CodeAerialRefuellingPrefixType]
	OnAfterCodeAerialRefuellingPrefixTypeDeleteCallback GongOnAfterDeleteInterface[CodeAerialRefuellingPrefixType]

	CodeAerialRefuellingTypes                map[*CodeAerialRefuellingType]struct{}
	CodeAerialRefuellingTypes_instance       map[*CodeAerialRefuellingType]*CodeAerialRefuellingType
	CodeAerialRefuellingTypes_mapString      map[string]*CodeAerialRefuellingType
	CodeAerialRefuellingTypeOrder            uint
	CodeAerialRefuellingType_stagedOrder     map[*CodeAerialRefuellingType]uint
	CodeAerialRefuellingType_orderStaged     map[uint]*CodeAerialRefuellingType
	CodeAerialRefuellingTypes_reference      map[*CodeAerialRefuellingType]*CodeAerialRefuellingType
	CodeAerialRefuellingTypes_referenceOrder map[*CodeAerialRefuellingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAerialRefuellingTypeCreateCallback GongOnAfterCreateInterface[CodeAerialRefuellingType]
	OnAfterCodeAerialRefuellingTypeUpdateCallback GongOnAfterUpdateInterface[CodeAerialRefuellingType]
	OnAfterCodeAerialRefuellingTypeDeleteCallback GongOnAfterDeleteInterface[CodeAerialRefuellingType]

	CodeAircraftCategoryTypes                map[*CodeAircraftCategoryType]struct{}
	CodeAircraftCategoryTypes_instance       map[*CodeAircraftCategoryType]*CodeAircraftCategoryType
	CodeAircraftCategoryTypes_mapString      map[string]*CodeAircraftCategoryType
	CodeAircraftCategoryTypeOrder            uint
	CodeAircraftCategoryType_stagedOrder     map[*CodeAircraftCategoryType]uint
	CodeAircraftCategoryType_orderStaged     map[uint]*CodeAircraftCategoryType
	CodeAircraftCategoryTypes_reference      map[*CodeAircraftCategoryType]*CodeAircraftCategoryType
	CodeAircraftCategoryTypes_referenceOrder map[*CodeAircraftCategoryType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftCategoryTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftCategoryType]
	OnAfterCodeAircraftCategoryTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftCategoryType]
	OnAfterCodeAircraftCategoryTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftCategoryType]

	CodeAircraftEngineNumberTypes                map[*CodeAircraftEngineNumberType]struct{}
	CodeAircraftEngineNumberTypes_instance       map[*CodeAircraftEngineNumberType]*CodeAircraftEngineNumberType
	CodeAircraftEngineNumberTypes_mapString      map[string]*CodeAircraftEngineNumberType
	CodeAircraftEngineNumberTypeOrder            uint
	CodeAircraftEngineNumberType_stagedOrder     map[*CodeAircraftEngineNumberType]uint
	CodeAircraftEngineNumberType_orderStaged     map[uint]*CodeAircraftEngineNumberType
	CodeAircraftEngineNumberTypes_reference      map[*CodeAircraftEngineNumberType]*CodeAircraftEngineNumberType
	CodeAircraftEngineNumberTypes_referenceOrder map[*CodeAircraftEngineNumberType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftEngineNumberTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftEngineNumberType]
	OnAfterCodeAircraftEngineNumberTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftEngineNumberType]
	OnAfterCodeAircraftEngineNumberTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftEngineNumberType]

	CodeAircraftEngineTypes                map[*CodeAircraftEngineType]struct{}
	CodeAircraftEngineTypes_instance       map[*CodeAircraftEngineType]*CodeAircraftEngineType
	CodeAircraftEngineTypes_mapString      map[string]*CodeAircraftEngineType
	CodeAircraftEngineTypeOrder            uint
	CodeAircraftEngineType_stagedOrder     map[*CodeAircraftEngineType]uint
	CodeAircraftEngineType_orderStaged     map[uint]*CodeAircraftEngineType
	CodeAircraftEngineTypes_reference      map[*CodeAircraftEngineType]*CodeAircraftEngineType
	CodeAircraftEngineTypes_referenceOrder map[*CodeAircraftEngineType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftEngineTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftEngineType]
	OnAfterCodeAircraftEngineTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftEngineType]
	OnAfterCodeAircraftEngineTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftEngineType]

	CodeAircraftICAOTypes                map[*CodeAircraftICAOType]struct{}
	CodeAircraftICAOTypes_instance       map[*CodeAircraftICAOType]*CodeAircraftICAOType
	CodeAircraftICAOTypes_mapString      map[string]*CodeAircraftICAOType
	CodeAircraftICAOTypeOrder            uint
	CodeAircraftICAOType_stagedOrder     map[*CodeAircraftICAOType]uint
	CodeAircraftICAOType_orderStaged     map[uint]*CodeAircraftICAOType
	CodeAircraftICAOTypes_reference      map[*CodeAircraftICAOType]*CodeAircraftICAOType
	CodeAircraftICAOTypes_referenceOrder map[*CodeAircraftICAOType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftICAOTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftICAOType]
	OnAfterCodeAircraftICAOTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftICAOType]
	OnAfterCodeAircraftICAOTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftICAOType]

	CodeAircraftTypes                map[*CodeAircraftType]struct{}
	CodeAircraftTypes_instance       map[*CodeAircraftType]*CodeAircraftType
	CodeAircraftTypes_mapString      map[string]*CodeAircraftType
	CodeAircraftTypeOrder            uint
	CodeAircraftType_stagedOrder     map[*CodeAircraftType]uint
	CodeAircraftType_orderStaged     map[uint]*CodeAircraftType
	CodeAircraftTypes_reference      map[*CodeAircraftType]*CodeAircraftType
	CodeAircraftTypes_referenceOrder map[*CodeAircraftType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftType]
	OnAfterCodeAircraftTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftType]
	OnAfterCodeAircraftTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftType]

	CodeAircraftWingspanClassTypes                map[*CodeAircraftWingspanClassType]struct{}
	CodeAircraftWingspanClassTypes_instance       map[*CodeAircraftWingspanClassType]*CodeAircraftWingspanClassType
	CodeAircraftWingspanClassTypes_mapString      map[string]*CodeAircraftWingspanClassType
	CodeAircraftWingspanClassTypeOrder            uint
	CodeAircraftWingspanClassType_stagedOrder     map[*CodeAircraftWingspanClassType]uint
	CodeAircraftWingspanClassType_orderStaged     map[uint]*CodeAircraftWingspanClassType
	CodeAircraftWingspanClassTypes_reference      map[*CodeAircraftWingspanClassType]*CodeAircraftWingspanClassType
	CodeAircraftWingspanClassTypes_referenceOrder map[*CodeAircraftWingspanClassType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftWingspanClassTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftWingspanClassType]
	OnAfterCodeAircraftWingspanClassTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftWingspanClassType]
	OnAfterCodeAircraftWingspanClassTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftWingspanClassType]

	CodeAltitudeAdjustmentTypes                map[*CodeAltitudeAdjustmentType]struct{}
	CodeAltitudeAdjustmentTypes_instance       map[*CodeAltitudeAdjustmentType]*CodeAltitudeAdjustmentType
	CodeAltitudeAdjustmentTypes_mapString      map[string]*CodeAltitudeAdjustmentType
	CodeAltitudeAdjustmentTypeOrder            uint
	CodeAltitudeAdjustmentType_stagedOrder     map[*CodeAltitudeAdjustmentType]uint
	CodeAltitudeAdjustmentType_orderStaged     map[uint]*CodeAltitudeAdjustmentType
	CodeAltitudeAdjustmentTypes_reference      map[*CodeAltitudeAdjustmentType]*CodeAltitudeAdjustmentType
	CodeAltitudeAdjustmentTypes_referenceOrder map[*CodeAltitudeAdjustmentType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAltitudeAdjustmentTypeCreateCallback GongOnAfterCreateInterface[CodeAltitudeAdjustmentType]
	OnAfterCodeAltitudeAdjustmentTypeUpdateCallback GongOnAfterUpdateInterface[CodeAltitudeAdjustmentType]
	OnAfterCodeAltitudeAdjustmentTypeDeleteCallback GongOnAfterDeleteInterface[CodeAltitudeAdjustmentType]

	CodeAltitudeUseTypes                map[*CodeAltitudeUseType]struct{}
	CodeAltitudeUseTypes_instance       map[*CodeAltitudeUseType]*CodeAltitudeUseType
	CodeAltitudeUseTypes_mapString      map[string]*CodeAltitudeUseType
	CodeAltitudeUseTypeOrder            uint
	CodeAltitudeUseType_stagedOrder     map[*CodeAltitudeUseType]uint
	CodeAltitudeUseType_orderStaged     map[uint]*CodeAltitudeUseType
	CodeAltitudeUseTypes_reference      map[*CodeAltitudeUseType]*CodeAltitudeUseType
	CodeAltitudeUseTypes_referenceOrder map[*CodeAltitudeUseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAltitudeUseTypeCreateCallback GongOnAfterCreateInterface[CodeAltitudeUseType]
	OnAfterCodeAltitudeUseTypeUpdateCallback GongOnAfterUpdateInterface[CodeAltitudeUseType]
	OnAfterCodeAltitudeUseTypeDeleteCallback GongOnAfterDeleteInterface[CodeAltitudeUseType]

	CodeAuthorityTypes                map[*CodeAuthorityType]struct{}
	CodeAuthorityTypes_instance       map[*CodeAuthorityType]*CodeAuthorityType
	CodeAuthorityTypes_mapString      map[string]*CodeAuthorityType
	CodeAuthorityTypeOrder            uint
	CodeAuthorityType_stagedOrder     map[*CodeAuthorityType]uint
	CodeAuthorityType_orderStaged     map[uint]*CodeAuthorityType
	CodeAuthorityTypes_reference      map[*CodeAuthorityType]*CodeAuthorityType
	CodeAuthorityTypes_referenceOrder map[*CodeAuthorityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAuthorityTypeCreateCallback GongOnAfterCreateInterface[CodeAuthorityType]
	OnAfterCodeAuthorityTypeUpdateCallback GongOnAfterUpdateInterface[CodeAuthorityType]
	OnAfterCodeAuthorityTypeDeleteCallback GongOnAfterDeleteInterface[CodeAuthorityType]

	CodeCardinalDirectionTypes                map[*CodeCardinalDirectionType]struct{}
	CodeCardinalDirectionTypes_instance       map[*CodeCardinalDirectionType]*CodeCardinalDirectionType
	CodeCardinalDirectionTypes_mapString      map[string]*CodeCardinalDirectionType
	CodeCardinalDirectionTypeOrder            uint
	CodeCardinalDirectionType_stagedOrder     map[*CodeCardinalDirectionType]uint
	CodeCardinalDirectionType_orderStaged     map[uint]*CodeCardinalDirectionType
	CodeCardinalDirectionTypes_reference      map[*CodeCardinalDirectionType]*CodeCardinalDirectionType
	CodeCardinalDirectionTypes_referenceOrder map[*CodeCardinalDirectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCardinalDirectionTypeCreateCallback GongOnAfterCreateInterface[CodeCardinalDirectionType]
	OnAfterCodeCardinalDirectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeCardinalDirectionType]
	OnAfterCodeCardinalDirectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeCardinalDirectionType]

	CodeCommunicationModeTypes                map[*CodeCommunicationModeType]struct{}
	CodeCommunicationModeTypes_instance       map[*CodeCommunicationModeType]*CodeCommunicationModeType
	CodeCommunicationModeTypes_mapString      map[string]*CodeCommunicationModeType
	CodeCommunicationModeTypeOrder            uint
	CodeCommunicationModeType_stagedOrder     map[*CodeCommunicationModeType]uint
	CodeCommunicationModeType_orderStaged     map[uint]*CodeCommunicationModeType
	CodeCommunicationModeTypes_reference      map[*CodeCommunicationModeType]*CodeCommunicationModeType
	CodeCommunicationModeTypes_referenceOrder map[*CodeCommunicationModeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCommunicationModeTypeCreateCallback GongOnAfterCreateInterface[CodeCommunicationModeType]
	OnAfterCodeCommunicationModeTypeUpdateCallback GongOnAfterUpdateInterface[CodeCommunicationModeType]
	OnAfterCodeCommunicationModeTypeDeleteCallback GongOnAfterDeleteInterface[CodeCommunicationModeType]

	CodeComparisonTypes                map[*CodeComparisonType]struct{}
	CodeComparisonTypes_instance       map[*CodeComparisonType]*CodeComparisonType
	CodeComparisonTypes_mapString      map[string]*CodeComparisonType
	CodeComparisonTypeOrder            uint
	CodeComparisonType_stagedOrder     map[*CodeComparisonType]uint
	CodeComparisonType_orderStaged     map[uint]*CodeComparisonType
	CodeComparisonTypes_reference      map[*CodeComparisonType]*CodeComparisonType
	CodeComparisonTypes_referenceOrder map[*CodeComparisonType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeComparisonTypeCreateCallback GongOnAfterCreateInterface[CodeComparisonType]
	OnAfterCodeComparisonTypeUpdateCallback GongOnAfterUpdateInterface[CodeComparisonType]
	OnAfterCodeComparisonTypeDeleteCallback GongOnAfterDeleteInterface[CodeComparisonType]

	CodeCourseTypes                map[*CodeCourseType]struct{}
	CodeCourseTypes_instance       map[*CodeCourseType]*CodeCourseType
	CodeCourseTypes_mapString      map[string]*CodeCourseType
	CodeCourseTypeOrder            uint
	CodeCourseType_stagedOrder     map[*CodeCourseType]uint
	CodeCourseType_orderStaged     map[uint]*CodeCourseType
	CodeCourseTypes_reference      map[*CodeCourseType]*CodeCourseType
	CodeCourseTypes_referenceOrder map[*CodeCourseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCourseTypeCreateCallback GongOnAfterCreateInterface[CodeCourseType]
	OnAfterCodeCourseTypeUpdateCallback GongOnAfterUpdateInterface[CodeCourseType]
	OnAfterCodeCourseTypeDeleteCallback GongOnAfterDeleteInterface[CodeCourseType]

	CodeDayTypes                map[*CodeDayType]struct{}
	CodeDayTypes_instance       map[*CodeDayType]*CodeDayType
	CodeDayTypes_mapString      map[string]*CodeDayType
	CodeDayTypeOrder            uint
	CodeDayType_stagedOrder     map[*CodeDayType]uint
	CodeDayType_orderStaged     map[uint]*CodeDayType
	CodeDayTypes_reference      map[*CodeDayType]*CodeDayType
	CodeDayTypes_referenceOrder map[*CodeDayType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDayTypeCreateCallback GongOnAfterCreateInterface[CodeDayType]
	OnAfterCodeDayTypeUpdateCallback GongOnAfterUpdateInterface[CodeDayType]
	OnAfterCodeDayTypeDeleteCallback GongOnAfterDeleteInterface[CodeDayType]

	CodeDirectionTurnTypes                map[*CodeDirectionTurnType]struct{}
	CodeDirectionTurnTypes_instance       map[*CodeDirectionTurnType]*CodeDirectionTurnType
	CodeDirectionTurnTypes_mapString      map[string]*CodeDirectionTurnType
	CodeDirectionTurnTypeOrder            uint
	CodeDirectionTurnType_stagedOrder     map[*CodeDirectionTurnType]uint
	CodeDirectionTurnType_orderStaged     map[uint]*CodeDirectionTurnType
	CodeDirectionTurnTypes_reference      map[*CodeDirectionTurnType]*CodeDirectionTurnType
	CodeDirectionTurnTypes_referenceOrder map[*CodeDirectionTurnType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDirectionTurnTypeCreateCallback GongOnAfterCreateInterface[CodeDirectionTurnType]
	OnAfterCodeDirectionTurnTypeUpdateCallback GongOnAfterUpdateInterface[CodeDirectionTurnType]
	OnAfterCodeDirectionTurnTypeDeleteCallback GongOnAfterDeleteInterface[CodeDirectionTurnType]

	CodeDirectionTypes                map[*CodeDirectionType]struct{}
	CodeDirectionTypes_instance       map[*CodeDirectionType]*CodeDirectionType
	CodeDirectionTypes_mapString      map[string]*CodeDirectionType
	CodeDirectionTypeOrder            uint
	CodeDirectionType_stagedOrder     map[*CodeDirectionType]uint
	CodeDirectionType_orderStaged     map[uint]*CodeDirectionType
	CodeDirectionTypes_reference      map[*CodeDirectionType]*CodeDirectionType
	CodeDirectionTypes_referenceOrder map[*CodeDirectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDirectionTypeCreateCallback GongOnAfterCreateInterface[CodeDirectionType]
	OnAfterCodeDirectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeDirectionType]
	OnAfterCodeDirectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeDirectionType]

	CodeEquipmentAntiCollisionTypes                map[*CodeEquipmentAntiCollisionType]struct{}
	CodeEquipmentAntiCollisionTypes_instance       map[*CodeEquipmentAntiCollisionType]*CodeEquipmentAntiCollisionType
	CodeEquipmentAntiCollisionTypes_mapString      map[string]*CodeEquipmentAntiCollisionType
	CodeEquipmentAntiCollisionTypeOrder            uint
	CodeEquipmentAntiCollisionType_stagedOrder     map[*CodeEquipmentAntiCollisionType]uint
	CodeEquipmentAntiCollisionType_orderStaged     map[uint]*CodeEquipmentAntiCollisionType
	CodeEquipmentAntiCollisionTypes_reference      map[*CodeEquipmentAntiCollisionType]*CodeEquipmentAntiCollisionType
	CodeEquipmentAntiCollisionTypes_referenceOrder map[*CodeEquipmentAntiCollisionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeEquipmentAntiCollisionTypeCreateCallback GongOnAfterCreateInterface[CodeEquipmentAntiCollisionType]
	OnAfterCodeEquipmentAntiCollisionTypeUpdateCallback GongOnAfterUpdateInterface[CodeEquipmentAntiCollisionType]
	OnAfterCodeEquipmentAntiCollisionTypeDeleteCallback GongOnAfterDeleteInterface[CodeEquipmentAntiCollisionType]

	CodeFlightOriginTypes                map[*CodeFlightOriginType]struct{}
	CodeFlightOriginTypes_instance       map[*CodeFlightOriginType]*CodeFlightOriginType
	CodeFlightOriginTypes_mapString      map[string]*CodeFlightOriginType
	CodeFlightOriginTypeOrder            uint
	CodeFlightOriginType_stagedOrder     map[*CodeFlightOriginType]uint
	CodeFlightOriginType_orderStaged     map[uint]*CodeFlightOriginType
	CodeFlightOriginTypes_reference      map[*CodeFlightOriginType]*CodeFlightOriginType
	CodeFlightOriginTypes_referenceOrder map[*CodeFlightOriginType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightOriginTypeCreateCallback GongOnAfterCreateInterface[CodeFlightOriginType]
	OnAfterCodeFlightOriginTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightOriginType]
	OnAfterCodeFlightOriginTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightOriginType]

	CodeFlightPurposeTypes                map[*CodeFlightPurposeType]struct{}
	CodeFlightPurposeTypes_instance       map[*CodeFlightPurposeType]*CodeFlightPurposeType
	CodeFlightPurposeTypes_mapString      map[string]*CodeFlightPurposeType
	CodeFlightPurposeTypeOrder            uint
	CodeFlightPurposeType_stagedOrder     map[*CodeFlightPurposeType]uint
	CodeFlightPurposeType_orderStaged     map[uint]*CodeFlightPurposeType
	CodeFlightPurposeTypes_reference      map[*CodeFlightPurposeType]*CodeFlightPurposeType
	CodeFlightPurposeTypes_referenceOrder map[*CodeFlightPurposeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightPurposeTypeCreateCallback GongOnAfterCreateInterface[CodeFlightPurposeType]
	OnAfterCodeFlightPurposeTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightPurposeType]
	OnAfterCodeFlightPurposeTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightPurposeType]

	CodeFlightRestrictionDesignatorTypes                map[*CodeFlightRestrictionDesignatorType]struct{}
	CodeFlightRestrictionDesignatorTypes_instance       map[*CodeFlightRestrictionDesignatorType]*CodeFlightRestrictionDesignatorType
	CodeFlightRestrictionDesignatorTypes_mapString      map[string]*CodeFlightRestrictionDesignatorType
	CodeFlightRestrictionDesignatorTypeOrder            uint
	CodeFlightRestrictionDesignatorType_stagedOrder     map[*CodeFlightRestrictionDesignatorType]uint
	CodeFlightRestrictionDesignatorType_orderStaged     map[uint]*CodeFlightRestrictionDesignatorType
	CodeFlightRestrictionDesignatorTypes_reference      map[*CodeFlightRestrictionDesignatorType]*CodeFlightRestrictionDesignatorType
	CodeFlightRestrictionDesignatorTypes_referenceOrder map[*CodeFlightRestrictionDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightRestrictionDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeFlightRestrictionDesignatorType]
	OnAfterCodeFlightRestrictionDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightRestrictionDesignatorType]
	OnAfterCodeFlightRestrictionDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightRestrictionDesignatorType]

	CodeFlightRestrictionTypes                map[*CodeFlightRestrictionType]struct{}
	CodeFlightRestrictionTypes_instance       map[*CodeFlightRestrictionType]*CodeFlightRestrictionType
	CodeFlightRestrictionTypes_mapString      map[string]*CodeFlightRestrictionType
	CodeFlightRestrictionTypeOrder            uint
	CodeFlightRestrictionType_stagedOrder     map[*CodeFlightRestrictionType]uint
	CodeFlightRestrictionType_orderStaged     map[uint]*CodeFlightRestrictionType
	CodeFlightRestrictionTypes_reference      map[*CodeFlightRestrictionType]*CodeFlightRestrictionType
	CodeFlightRestrictionTypes_referenceOrder map[*CodeFlightRestrictionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightRestrictionTypeCreateCallback GongOnAfterCreateInterface[CodeFlightRestrictionType]
	OnAfterCodeFlightRestrictionTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightRestrictionType]
	OnAfterCodeFlightRestrictionTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightRestrictionType]

	CodeFlightRuleTypes                map[*CodeFlightRuleType]struct{}
	CodeFlightRuleTypes_instance       map[*CodeFlightRuleType]*CodeFlightRuleType
	CodeFlightRuleTypes_mapString      map[string]*CodeFlightRuleType
	CodeFlightRuleTypeOrder            uint
	CodeFlightRuleType_stagedOrder     map[*CodeFlightRuleType]uint
	CodeFlightRuleType_orderStaged     map[uint]*CodeFlightRuleType
	CodeFlightRuleTypes_reference      map[*CodeFlightRuleType]*CodeFlightRuleType
	CodeFlightRuleTypes_referenceOrder map[*CodeFlightRuleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightRuleTypeCreateCallback GongOnAfterCreateInterface[CodeFlightRuleType]
	OnAfterCodeFlightRuleTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightRuleType]
	OnAfterCodeFlightRuleTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightRuleType]

	CodeFlightStatusTypes                map[*CodeFlightStatusType]struct{}
	CodeFlightStatusTypes_instance       map[*CodeFlightStatusType]*CodeFlightStatusType
	CodeFlightStatusTypes_mapString      map[string]*CodeFlightStatusType
	CodeFlightStatusTypeOrder            uint
	CodeFlightStatusType_stagedOrder     map[*CodeFlightStatusType]uint
	CodeFlightStatusType_orderStaged     map[uint]*CodeFlightStatusType
	CodeFlightStatusTypes_reference      map[*CodeFlightStatusType]*CodeFlightStatusType
	CodeFlightStatusTypes_referenceOrder map[*CodeFlightStatusType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightStatusTypeCreateCallback GongOnAfterCreateInterface[CodeFlightStatusType]
	OnAfterCodeFlightStatusTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightStatusType]
	OnAfterCodeFlightStatusTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightStatusType]

	CodeFlightTypes                map[*CodeFlightType]struct{}
	CodeFlightTypes_instance       map[*CodeFlightType]*CodeFlightType
	CodeFlightTypes_mapString      map[string]*CodeFlightType
	CodeFlightTypeOrder            uint
	CodeFlightType_stagedOrder     map[*CodeFlightType]uint
	CodeFlightType_orderStaged     map[uint]*CodeFlightType
	CodeFlightTypes_reference      map[*CodeFlightType]*CodeFlightType
	CodeFlightTypes_referenceOrder map[*CodeFlightType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightTypeCreateCallback GongOnAfterCreateInterface[CodeFlightType]
	OnAfterCodeFlightTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightType]
	OnAfterCodeFlightTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightType]

	CodeFlowConditionOperationTypes                map[*CodeFlowConditionOperationType]struct{}
	CodeFlowConditionOperationTypes_instance       map[*CodeFlowConditionOperationType]*CodeFlowConditionOperationType
	CodeFlowConditionOperationTypes_mapString      map[string]*CodeFlowConditionOperationType
	CodeFlowConditionOperationTypeOrder            uint
	CodeFlowConditionOperationType_stagedOrder     map[*CodeFlowConditionOperationType]uint
	CodeFlowConditionOperationType_orderStaged     map[uint]*CodeFlowConditionOperationType
	CodeFlowConditionOperationTypes_reference      map[*CodeFlowConditionOperationType]*CodeFlowConditionOperationType
	CodeFlowConditionOperationTypes_referenceOrder map[*CodeFlowConditionOperationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlowConditionOperationTypeCreateCallback GongOnAfterCreateInterface[CodeFlowConditionOperationType]
	OnAfterCodeFlowConditionOperationTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlowConditionOperationType]
	OnAfterCodeFlowConditionOperationTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlowConditionOperationType]

	CodeFreeFlightTypes                map[*CodeFreeFlightType]struct{}
	CodeFreeFlightTypes_instance       map[*CodeFreeFlightType]*CodeFreeFlightType
	CodeFreeFlightTypes_mapString      map[string]*CodeFreeFlightType
	CodeFreeFlightTypeOrder            uint
	CodeFreeFlightType_stagedOrder     map[*CodeFreeFlightType]uint
	CodeFreeFlightType_orderStaged     map[uint]*CodeFreeFlightType
	CodeFreeFlightTypes_reference      map[*CodeFreeFlightType]*CodeFreeFlightType
	CodeFreeFlightTypes_referenceOrder map[*CodeFreeFlightType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFreeFlightTypeCreateCallback GongOnAfterCreateInterface[CodeFreeFlightType]
	OnAfterCodeFreeFlightTypeUpdateCallback GongOnAfterUpdateInterface[CodeFreeFlightType]
	OnAfterCodeFreeFlightTypeDeleteCallback GongOnAfterDeleteInterface[CodeFreeFlightType]

	CodeLevelTypes                map[*CodeLevelType]struct{}
	CodeLevelTypes_instance       map[*CodeLevelType]*CodeLevelType
	CodeLevelTypes_mapString      map[string]*CodeLevelType
	CodeLevelTypeOrder            uint
	CodeLevelType_stagedOrder     map[*CodeLevelType]uint
	CodeLevelType_orderStaged     map[uint]*CodeLevelType
	CodeLevelTypes_reference      map[*CodeLevelType]*CodeLevelType
	CodeLevelTypes_referenceOrder map[*CodeLevelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLevelTypeCreateCallback GongOnAfterCreateInterface[CodeLevelType]
	OnAfterCodeLevelTypeUpdateCallback GongOnAfterUpdateInterface[CodeLevelType]
	OnAfterCodeLevelTypeDeleteCallback GongOnAfterDeleteInterface[CodeLevelType]

	CodeLocationQualifierTypes                map[*CodeLocationQualifierType]struct{}
	CodeLocationQualifierTypes_instance       map[*CodeLocationQualifierType]*CodeLocationQualifierType
	CodeLocationQualifierTypes_mapString      map[string]*CodeLocationQualifierType
	CodeLocationQualifierTypeOrder            uint
	CodeLocationQualifierType_stagedOrder     map[*CodeLocationQualifierType]uint
	CodeLocationQualifierType_orderStaged     map[uint]*CodeLocationQualifierType
	CodeLocationQualifierTypes_reference      map[*CodeLocationQualifierType]*CodeLocationQualifierType
	CodeLocationQualifierTypes_referenceOrder map[*CodeLocationQualifierType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLocationQualifierTypeCreateCallback GongOnAfterCreateInterface[CodeLocationQualifierType]
	OnAfterCodeLocationQualifierTypeUpdateCallback GongOnAfterUpdateInterface[CodeLocationQualifierType]
	OnAfterCodeLocationQualifierTypeDeleteCallback GongOnAfterDeleteInterface[CodeLocationQualifierType]

	CodeMeteoConditionsTypes                map[*CodeMeteoConditionsType]struct{}
	CodeMeteoConditionsTypes_instance       map[*CodeMeteoConditionsType]*CodeMeteoConditionsType
	CodeMeteoConditionsTypes_mapString      map[string]*CodeMeteoConditionsType
	CodeMeteoConditionsTypeOrder            uint
	CodeMeteoConditionsType_stagedOrder     map[*CodeMeteoConditionsType]uint
	CodeMeteoConditionsType_orderStaged     map[uint]*CodeMeteoConditionsType
	CodeMeteoConditionsTypes_reference      map[*CodeMeteoConditionsType]*CodeMeteoConditionsType
	CodeMeteoConditionsTypes_referenceOrder map[*CodeMeteoConditionsType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMeteoConditionsTypeCreateCallback GongOnAfterCreateInterface[CodeMeteoConditionsType]
	OnAfterCodeMeteoConditionsTypeUpdateCallback GongOnAfterUpdateInterface[CodeMeteoConditionsType]
	OnAfterCodeMeteoConditionsTypeDeleteCallback GongOnAfterDeleteInterface[CodeMeteoConditionsType]

	CodeMilitaryRoutePointTypes                map[*CodeMilitaryRoutePointType]struct{}
	CodeMilitaryRoutePointTypes_instance       map[*CodeMilitaryRoutePointType]*CodeMilitaryRoutePointType
	CodeMilitaryRoutePointTypes_mapString      map[string]*CodeMilitaryRoutePointType
	CodeMilitaryRoutePointTypeOrder            uint
	CodeMilitaryRoutePointType_stagedOrder     map[*CodeMilitaryRoutePointType]uint
	CodeMilitaryRoutePointType_orderStaged     map[uint]*CodeMilitaryRoutePointType
	CodeMilitaryRoutePointTypes_reference      map[*CodeMilitaryRoutePointType]*CodeMilitaryRoutePointType
	CodeMilitaryRoutePointTypes_referenceOrder map[*CodeMilitaryRoutePointType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMilitaryRoutePointTypeCreateCallback GongOnAfterCreateInterface[CodeMilitaryRoutePointType]
	OnAfterCodeMilitaryRoutePointTypeUpdateCallback GongOnAfterUpdateInterface[CodeMilitaryRoutePointType]
	OnAfterCodeMilitaryRoutePointTypeDeleteCallback GongOnAfterDeleteInterface[CodeMilitaryRoutePointType]

	CodeMilitaryStatusTypes                map[*CodeMilitaryStatusType]struct{}
	CodeMilitaryStatusTypes_instance       map[*CodeMilitaryStatusType]*CodeMilitaryStatusType
	CodeMilitaryStatusTypes_mapString      map[string]*CodeMilitaryStatusType
	CodeMilitaryStatusTypeOrder            uint
	CodeMilitaryStatusType_stagedOrder     map[*CodeMilitaryStatusType]uint
	CodeMilitaryStatusType_orderStaged     map[uint]*CodeMilitaryStatusType
	CodeMilitaryStatusTypes_reference      map[*CodeMilitaryStatusType]*CodeMilitaryStatusType
	CodeMilitaryStatusTypes_referenceOrder map[*CodeMilitaryStatusType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMilitaryStatusTypeCreateCallback GongOnAfterCreateInterface[CodeMilitaryStatusType]
	OnAfterCodeMilitaryStatusTypeUpdateCallback GongOnAfterUpdateInterface[CodeMilitaryStatusType]
	OnAfterCodeMilitaryStatusTypeDeleteCallback GongOnAfterDeleteInterface[CodeMilitaryStatusType]

	CodeMilitaryTrainingTypes                map[*CodeMilitaryTrainingType]struct{}
	CodeMilitaryTrainingTypes_instance       map[*CodeMilitaryTrainingType]*CodeMilitaryTrainingType
	CodeMilitaryTrainingTypes_mapString      map[string]*CodeMilitaryTrainingType
	CodeMilitaryTrainingTypeOrder            uint
	CodeMilitaryTrainingType_stagedOrder     map[*CodeMilitaryTrainingType]uint
	CodeMilitaryTrainingType_orderStaged     map[uint]*CodeMilitaryTrainingType
	CodeMilitaryTrainingTypes_reference      map[*CodeMilitaryTrainingType]*CodeMilitaryTrainingType
	CodeMilitaryTrainingTypes_referenceOrder map[*CodeMilitaryTrainingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMilitaryTrainingTypeCreateCallback GongOnAfterCreateInterface[CodeMilitaryTrainingType]
	OnAfterCodeMilitaryTrainingTypeUpdateCallback GongOnAfterUpdateInterface[CodeMilitaryTrainingType]
	OnAfterCodeMilitaryTrainingTypeDeleteCallback GongOnAfterDeleteInterface[CodeMilitaryTrainingType]

	CodeNavigationEquipmentTypes                map[*CodeNavigationEquipmentType]struct{}
	CodeNavigationEquipmentTypes_instance       map[*CodeNavigationEquipmentType]*CodeNavigationEquipmentType
	CodeNavigationEquipmentTypes_mapString      map[string]*CodeNavigationEquipmentType
	CodeNavigationEquipmentTypeOrder            uint
	CodeNavigationEquipmentType_stagedOrder     map[*CodeNavigationEquipmentType]uint
	CodeNavigationEquipmentType_orderStaged     map[uint]*CodeNavigationEquipmentType
	CodeNavigationEquipmentTypes_reference      map[*CodeNavigationEquipmentType]*CodeNavigationEquipmentType
	CodeNavigationEquipmentTypes_referenceOrder map[*CodeNavigationEquipmentType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavigationEquipmentTypeCreateCallback GongOnAfterCreateInterface[CodeNavigationEquipmentType]
	OnAfterCodeNavigationEquipmentTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavigationEquipmentType]
	OnAfterCodeNavigationEquipmentTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavigationEquipmentType]

	CodeNavigationSpecificationTypes                map[*CodeNavigationSpecificationType]struct{}
	CodeNavigationSpecificationTypes_instance       map[*CodeNavigationSpecificationType]*CodeNavigationSpecificationType
	CodeNavigationSpecificationTypes_mapString      map[string]*CodeNavigationSpecificationType
	CodeNavigationSpecificationTypeOrder            uint
	CodeNavigationSpecificationType_stagedOrder     map[*CodeNavigationSpecificationType]uint
	CodeNavigationSpecificationType_orderStaged     map[uint]*CodeNavigationSpecificationType
	CodeNavigationSpecificationTypes_reference      map[*CodeNavigationSpecificationType]*CodeNavigationSpecificationType
	CodeNavigationSpecificationTypes_referenceOrder map[*CodeNavigationSpecificationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavigationSpecificationTypeCreateCallback GongOnAfterCreateInterface[CodeNavigationSpecificationType]
	OnAfterCodeNavigationSpecificationTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavigationSpecificationType]
	OnAfterCodeNavigationSpecificationTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavigationSpecificationType]

	CodeNotePurposeTypes                map[*CodeNotePurposeType]struct{}
	CodeNotePurposeTypes_instance       map[*CodeNotePurposeType]*CodeNotePurposeType
	CodeNotePurposeTypes_mapString      map[string]*CodeNotePurposeType
	CodeNotePurposeTypeOrder            uint
	CodeNotePurposeType_stagedOrder     map[*CodeNotePurposeType]uint
	CodeNotePurposeType_orderStaged     map[uint]*CodeNotePurposeType
	CodeNotePurposeTypes_reference      map[*CodeNotePurposeType]*CodeNotePurposeType
	CodeNotePurposeTypes_referenceOrder map[*CodeNotePurposeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNotePurposeTypeCreateCallback GongOnAfterCreateInterface[CodeNotePurposeType]
	OnAfterCodeNotePurposeTypeUpdateCallback GongOnAfterUpdateInterface[CodeNotePurposeType]
	OnAfterCodeNotePurposeTypeDeleteCallback GongOnAfterDeleteInterface[CodeNotePurposeType]

	CodeObstacleAssessmentSurfaceTypes                map[*CodeObstacleAssessmentSurfaceType]struct{}
	CodeObstacleAssessmentSurfaceTypes_instance       map[*CodeObstacleAssessmentSurfaceType]*CodeObstacleAssessmentSurfaceType
	CodeObstacleAssessmentSurfaceTypes_mapString      map[string]*CodeObstacleAssessmentSurfaceType
	CodeObstacleAssessmentSurfaceTypeOrder            uint
	CodeObstacleAssessmentSurfaceType_stagedOrder     map[*CodeObstacleAssessmentSurfaceType]uint
	CodeObstacleAssessmentSurfaceType_orderStaged     map[uint]*CodeObstacleAssessmentSurfaceType
	CodeObstacleAssessmentSurfaceTypes_reference      map[*CodeObstacleAssessmentSurfaceType]*CodeObstacleAssessmentSurfaceType
	CodeObstacleAssessmentSurfaceTypes_referenceOrder map[*CodeObstacleAssessmentSurfaceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeObstacleAssessmentSurfaceTypeCreateCallback GongOnAfterCreateInterface[CodeObstacleAssessmentSurfaceType]
	OnAfterCodeObstacleAssessmentSurfaceTypeUpdateCallback GongOnAfterUpdateInterface[CodeObstacleAssessmentSurfaceType]
	OnAfterCodeObstacleAssessmentSurfaceTypeDeleteCallback GongOnAfterDeleteInterface[CodeObstacleAssessmentSurfaceType]

	CodeObstructionIdSurfaceZoneTypes                map[*CodeObstructionIdSurfaceZoneType]struct{}
	CodeObstructionIdSurfaceZoneTypes_instance       map[*CodeObstructionIdSurfaceZoneType]*CodeObstructionIdSurfaceZoneType
	CodeObstructionIdSurfaceZoneTypes_mapString      map[string]*CodeObstructionIdSurfaceZoneType
	CodeObstructionIdSurfaceZoneTypeOrder            uint
	CodeObstructionIdSurfaceZoneType_stagedOrder     map[*CodeObstructionIdSurfaceZoneType]uint
	CodeObstructionIdSurfaceZoneType_orderStaged     map[uint]*CodeObstructionIdSurfaceZoneType
	CodeObstructionIdSurfaceZoneTypes_reference      map[*CodeObstructionIdSurfaceZoneType]*CodeObstructionIdSurfaceZoneType
	CodeObstructionIdSurfaceZoneTypes_referenceOrder map[*CodeObstructionIdSurfaceZoneType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeObstructionIdSurfaceZoneTypeCreateCallback GongOnAfterCreateInterface[CodeObstructionIdSurfaceZoneType]
	OnAfterCodeObstructionIdSurfaceZoneTypeUpdateCallback GongOnAfterUpdateInterface[CodeObstructionIdSurfaceZoneType]
	OnAfterCodeObstructionIdSurfaceZoneTypeDeleteCallback GongOnAfterDeleteInterface[CodeObstructionIdSurfaceZoneType]

	CodeRNPTypes                map[*CodeRNPType]struct{}
	CodeRNPTypes_instance       map[*CodeRNPType]*CodeRNPType
	CodeRNPTypes_mapString      map[string]*CodeRNPType
	CodeRNPTypeOrder            uint
	CodeRNPType_stagedOrder     map[*CodeRNPType]uint
	CodeRNPType_orderStaged     map[uint]*CodeRNPType
	CodeRNPTypes_reference      map[*CodeRNPType]*CodeRNPType
	CodeRNPTypes_referenceOrder map[*CodeRNPType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRNPTypeCreateCallback GongOnAfterCreateInterface[CodeRNPType]
	OnAfterCodeRNPTypeUpdateCallback GongOnAfterUpdateInterface[CodeRNPType]
	OnAfterCodeRNPTypeDeleteCallback GongOnAfterDeleteInterface[CodeRNPType]

	CodeRVSMPointRoleTypes                map[*CodeRVSMPointRoleType]struct{}
	CodeRVSMPointRoleTypes_instance       map[*CodeRVSMPointRoleType]*CodeRVSMPointRoleType
	CodeRVSMPointRoleTypes_mapString      map[string]*CodeRVSMPointRoleType
	CodeRVSMPointRoleTypeOrder            uint
	CodeRVSMPointRoleType_stagedOrder     map[*CodeRVSMPointRoleType]uint
	CodeRVSMPointRoleType_orderStaged     map[uint]*CodeRVSMPointRoleType
	CodeRVSMPointRoleTypes_reference      map[*CodeRVSMPointRoleType]*CodeRVSMPointRoleType
	CodeRVSMPointRoleTypes_referenceOrder map[*CodeRVSMPointRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRVSMPointRoleTypeCreateCallback GongOnAfterCreateInterface[CodeRVSMPointRoleType]
	OnAfterCodeRVSMPointRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeRVSMPointRoleType]
	OnAfterCodeRVSMPointRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeRVSMPointRoleType]

	CodeRVSMTypes                map[*CodeRVSMType]struct{}
	CodeRVSMTypes_instance       map[*CodeRVSMType]*CodeRVSMType
	CodeRVSMTypes_mapString      map[string]*CodeRVSMType
	CodeRVSMTypeOrder            uint
	CodeRVSMType_stagedOrder     map[*CodeRVSMType]uint
	CodeRVSMType_orderStaged     map[uint]*CodeRVSMType
	CodeRVSMTypes_reference      map[*CodeRVSMType]*CodeRVSMType
	CodeRVSMTypes_referenceOrder map[*CodeRVSMType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRVSMTypeCreateCallback GongOnAfterCreateInterface[CodeRVSMType]
	OnAfterCodeRVSMTypeUpdateCallback GongOnAfterUpdateInterface[CodeRVSMType]
	OnAfterCodeRVSMTypeDeleteCallback GongOnAfterDeleteInterface[CodeRVSMType]

	CodeReferenceRoleTypes                map[*CodeReferenceRoleType]struct{}
	CodeReferenceRoleTypes_instance       map[*CodeReferenceRoleType]*CodeReferenceRoleType
	CodeReferenceRoleTypes_mapString      map[string]*CodeReferenceRoleType
	CodeReferenceRoleTypeOrder            uint
	CodeReferenceRoleType_stagedOrder     map[*CodeReferenceRoleType]uint
	CodeReferenceRoleType_orderStaged     map[uint]*CodeReferenceRoleType
	CodeReferenceRoleTypes_reference      map[*CodeReferenceRoleType]*CodeReferenceRoleType
	CodeReferenceRoleTypes_referenceOrder map[*CodeReferenceRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeReferenceRoleTypeCreateCallback GongOnAfterCreateInterface[CodeReferenceRoleType]
	OnAfterCodeReferenceRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeReferenceRoleType]
	OnAfterCodeReferenceRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeReferenceRoleType]

	CodeRouteAvailabilityTypes                map[*CodeRouteAvailabilityType]struct{}
	CodeRouteAvailabilityTypes_instance       map[*CodeRouteAvailabilityType]*CodeRouteAvailabilityType
	CodeRouteAvailabilityTypes_mapString      map[string]*CodeRouteAvailabilityType
	CodeRouteAvailabilityTypeOrder            uint
	CodeRouteAvailabilityType_stagedOrder     map[*CodeRouteAvailabilityType]uint
	CodeRouteAvailabilityType_orderStaged     map[uint]*CodeRouteAvailabilityType
	CodeRouteAvailabilityTypes_reference      map[*CodeRouteAvailabilityType]*CodeRouteAvailabilityType
	CodeRouteAvailabilityTypes_referenceOrder map[*CodeRouteAvailabilityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteAvailabilityTypeCreateCallback GongOnAfterCreateInterface[CodeRouteAvailabilityType]
	OnAfterCodeRouteAvailabilityTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteAvailabilityType]
	OnAfterCodeRouteAvailabilityTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteAvailabilityType]

	CodeRouteDesignatorLetterTypes                map[*CodeRouteDesignatorLetterType]struct{}
	CodeRouteDesignatorLetterTypes_instance       map[*CodeRouteDesignatorLetterType]*CodeRouteDesignatorLetterType
	CodeRouteDesignatorLetterTypes_mapString      map[string]*CodeRouteDesignatorLetterType
	CodeRouteDesignatorLetterTypeOrder            uint
	CodeRouteDesignatorLetterType_stagedOrder     map[*CodeRouteDesignatorLetterType]uint
	CodeRouteDesignatorLetterType_orderStaged     map[uint]*CodeRouteDesignatorLetterType
	CodeRouteDesignatorLetterTypes_reference      map[*CodeRouteDesignatorLetterType]*CodeRouteDesignatorLetterType
	CodeRouteDesignatorLetterTypes_referenceOrder map[*CodeRouteDesignatorLetterType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteDesignatorLetterTypeCreateCallback GongOnAfterCreateInterface[CodeRouteDesignatorLetterType]
	OnAfterCodeRouteDesignatorLetterTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteDesignatorLetterType]
	OnAfterCodeRouteDesignatorLetterTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteDesignatorLetterType]

	CodeRouteDesignatorPrefixTypes                map[*CodeRouteDesignatorPrefixType]struct{}
	CodeRouteDesignatorPrefixTypes_instance       map[*CodeRouteDesignatorPrefixType]*CodeRouteDesignatorPrefixType
	CodeRouteDesignatorPrefixTypes_mapString      map[string]*CodeRouteDesignatorPrefixType
	CodeRouteDesignatorPrefixTypeOrder            uint
	CodeRouteDesignatorPrefixType_stagedOrder     map[*CodeRouteDesignatorPrefixType]uint
	CodeRouteDesignatorPrefixType_orderStaged     map[uint]*CodeRouteDesignatorPrefixType
	CodeRouteDesignatorPrefixTypes_reference      map[*CodeRouteDesignatorPrefixType]*CodeRouteDesignatorPrefixType
	CodeRouteDesignatorPrefixTypes_referenceOrder map[*CodeRouteDesignatorPrefixType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteDesignatorPrefixTypeCreateCallback GongOnAfterCreateInterface[CodeRouteDesignatorPrefixType]
	OnAfterCodeRouteDesignatorPrefixTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteDesignatorPrefixType]
	OnAfterCodeRouteDesignatorPrefixTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteDesignatorPrefixType]

	CodeRouteDesignatorSuffixTypes                map[*CodeRouteDesignatorSuffixType]struct{}
	CodeRouteDesignatorSuffixTypes_instance       map[*CodeRouteDesignatorSuffixType]*CodeRouteDesignatorSuffixType
	CodeRouteDesignatorSuffixTypes_mapString      map[string]*CodeRouteDesignatorSuffixType
	CodeRouteDesignatorSuffixTypeOrder            uint
	CodeRouteDesignatorSuffixType_stagedOrder     map[*CodeRouteDesignatorSuffixType]uint
	CodeRouteDesignatorSuffixType_orderStaged     map[uint]*CodeRouteDesignatorSuffixType
	CodeRouteDesignatorSuffixTypes_reference      map[*CodeRouteDesignatorSuffixType]*CodeRouteDesignatorSuffixType
	CodeRouteDesignatorSuffixTypes_referenceOrder map[*CodeRouteDesignatorSuffixType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteDesignatorSuffixTypeCreateCallback GongOnAfterCreateInterface[CodeRouteDesignatorSuffixType]
	OnAfterCodeRouteDesignatorSuffixTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteDesignatorSuffixType]
	OnAfterCodeRouteDesignatorSuffixTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteDesignatorSuffixType]

	CodeRouteNavigationTypes                map[*CodeRouteNavigationType]struct{}
	CodeRouteNavigationTypes_instance       map[*CodeRouteNavigationType]*CodeRouteNavigationType
	CodeRouteNavigationTypes_mapString      map[string]*CodeRouteNavigationType
	CodeRouteNavigationTypeOrder            uint
	CodeRouteNavigationType_stagedOrder     map[*CodeRouteNavigationType]uint
	CodeRouteNavigationType_orderStaged     map[uint]*CodeRouteNavigationType
	CodeRouteNavigationTypes_reference      map[*CodeRouteNavigationType]*CodeRouteNavigationType
	CodeRouteNavigationTypes_referenceOrder map[*CodeRouteNavigationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteNavigationTypeCreateCallback GongOnAfterCreateInterface[CodeRouteNavigationType]
	OnAfterCodeRouteNavigationTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteNavigationType]
	OnAfterCodeRouteNavigationTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteNavigationType]

	CodeRouteOriginTypes                map[*CodeRouteOriginType]struct{}
	CodeRouteOriginTypes_instance       map[*CodeRouteOriginType]*CodeRouteOriginType
	CodeRouteOriginTypes_mapString      map[string]*CodeRouteOriginType
	CodeRouteOriginTypeOrder            uint
	CodeRouteOriginType_stagedOrder     map[*CodeRouteOriginType]uint
	CodeRouteOriginType_orderStaged     map[uint]*CodeRouteOriginType
	CodeRouteOriginTypes_reference      map[*CodeRouteOriginType]*CodeRouteOriginType
	CodeRouteOriginTypes_referenceOrder map[*CodeRouteOriginType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteOriginTypeCreateCallback GongOnAfterCreateInterface[CodeRouteOriginType]
	OnAfterCodeRouteOriginTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteOriginType]
	OnAfterCodeRouteOriginTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteOriginType]

	CodeRouteSegmentPathTypes                map[*CodeRouteSegmentPathType]struct{}
	CodeRouteSegmentPathTypes_instance       map[*CodeRouteSegmentPathType]*CodeRouteSegmentPathType
	CodeRouteSegmentPathTypes_mapString      map[string]*CodeRouteSegmentPathType
	CodeRouteSegmentPathTypeOrder            uint
	CodeRouteSegmentPathType_stagedOrder     map[*CodeRouteSegmentPathType]uint
	CodeRouteSegmentPathType_orderStaged     map[uint]*CodeRouteSegmentPathType
	CodeRouteSegmentPathTypes_reference      map[*CodeRouteSegmentPathType]*CodeRouteSegmentPathType
	CodeRouteSegmentPathTypes_referenceOrder map[*CodeRouteSegmentPathType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteSegmentPathTypeCreateCallback GongOnAfterCreateInterface[CodeRouteSegmentPathType]
	OnAfterCodeRouteSegmentPathTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteSegmentPathType]
	OnAfterCodeRouteSegmentPathTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteSegmentPathType]

	CodeRouteTypes                map[*CodeRouteType]struct{}
	CodeRouteTypes_instance       map[*CodeRouteType]*CodeRouteType
	CodeRouteTypes_mapString      map[string]*CodeRouteType
	CodeRouteTypeOrder            uint
	CodeRouteType_stagedOrder     map[*CodeRouteType]uint
	CodeRouteType_orderStaged     map[uint]*CodeRouteType
	CodeRouteTypes_reference      map[*CodeRouteType]*CodeRouteType
	CodeRouteTypes_referenceOrder map[*CodeRouteType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteTypeCreateCallback GongOnAfterCreateInterface[CodeRouteType]
	OnAfterCodeRouteTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteType]
	OnAfterCodeRouteTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteType]

	CodeRuleProcedureTitleTypes                map[*CodeRuleProcedureTitleType]struct{}
	CodeRuleProcedureTitleTypes_instance       map[*CodeRuleProcedureTitleType]*CodeRuleProcedureTitleType
	CodeRuleProcedureTitleTypes_mapString      map[string]*CodeRuleProcedureTitleType
	CodeRuleProcedureTitleTypeOrder            uint
	CodeRuleProcedureTitleType_stagedOrder     map[*CodeRuleProcedureTitleType]uint
	CodeRuleProcedureTitleType_orderStaged     map[uint]*CodeRuleProcedureTitleType
	CodeRuleProcedureTitleTypes_reference      map[*CodeRuleProcedureTitleType]*CodeRuleProcedureTitleType
	CodeRuleProcedureTitleTypes_referenceOrder map[*CodeRuleProcedureTitleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRuleProcedureTitleTypeCreateCallback GongOnAfterCreateInterface[CodeRuleProcedureTitleType]
	OnAfterCodeRuleProcedureTitleTypeUpdateCallback GongOnAfterUpdateInterface[CodeRuleProcedureTitleType]
	OnAfterCodeRuleProcedureTitleTypeDeleteCallback GongOnAfterDeleteInterface[CodeRuleProcedureTitleType]

	CodeRuleProcedureTypes                map[*CodeRuleProcedureType]struct{}
	CodeRuleProcedureTypes_instance       map[*CodeRuleProcedureType]*CodeRuleProcedureType
	CodeRuleProcedureTypes_mapString      map[string]*CodeRuleProcedureType
	CodeRuleProcedureTypeOrder            uint
	CodeRuleProcedureType_stagedOrder     map[*CodeRuleProcedureType]uint
	CodeRuleProcedureType_orderStaged     map[uint]*CodeRuleProcedureType
	CodeRuleProcedureTypes_reference      map[*CodeRuleProcedureType]*CodeRuleProcedureType
	CodeRuleProcedureTypes_referenceOrder map[*CodeRuleProcedureType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRuleProcedureTypeCreateCallback GongOnAfterCreateInterface[CodeRuleProcedureType]
	OnAfterCodeRuleProcedureTypeUpdateCallback GongOnAfterUpdateInterface[CodeRuleProcedureType]
	OnAfterCodeRuleProcedureTypeDeleteCallback GongOnAfterDeleteInterface[CodeRuleProcedureType]

	CodeSideTypes                map[*CodeSideType]struct{}
	CodeSideTypes_instance       map[*CodeSideType]*CodeSideType
	CodeSideTypes_mapString      map[string]*CodeSideType
	CodeSideTypeOrder            uint
	CodeSideType_stagedOrder     map[*CodeSideType]uint
	CodeSideType_orderStaged     map[uint]*CodeSideType
	CodeSideTypes_reference      map[*CodeSideType]*CodeSideType
	CodeSideTypes_referenceOrder map[*CodeSideType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSideTypeCreateCallback GongOnAfterCreateInterface[CodeSideType]
	OnAfterCodeSideTypeUpdateCallback GongOnAfterUpdateInterface[CodeSideType]
	OnAfterCodeSideTypeDeleteCallback GongOnAfterDeleteInterface[CodeSideType]

	CodeSpeedReferenceTypes                map[*CodeSpeedReferenceType]struct{}
	CodeSpeedReferenceTypes_instance       map[*CodeSpeedReferenceType]*CodeSpeedReferenceType
	CodeSpeedReferenceTypes_mapString      map[string]*CodeSpeedReferenceType
	CodeSpeedReferenceTypeOrder            uint
	CodeSpeedReferenceType_stagedOrder     map[*CodeSpeedReferenceType]uint
	CodeSpeedReferenceType_orderStaged     map[uint]*CodeSpeedReferenceType
	CodeSpeedReferenceTypes_reference      map[*CodeSpeedReferenceType]*CodeSpeedReferenceType
	CodeSpeedReferenceTypes_referenceOrder map[*CodeSpeedReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSpeedReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeSpeedReferenceType]
	OnAfterCodeSpeedReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeSpeedReferenceType]
	OnAfterCodeSpeedReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeSpeedReferenceType]

	CodeTACANChannelTypes                map[*CodeTACANChannelType]struct{}
	CodeTACANChannelTypes_instance       map[*CodeTACANChannelType]*CodeTACANChannelType
	CodeTACANChannelTypes_mapString      map[string]*CodeTACANChannelType
	CodeTACANChannelTypeOrder            uint
	CodeTACANChannelType_stagedOrder     map[*CodeTACANChannelType]uint
	CodeTACANChannelType_orderStaged     map[uint]*CodeTACANChannelType
	CodeTACANChannelTypes_reference      map[*CodeTACANChannelType]*CodeTACANChannelType
	CodeTACANChannelTypes_referenceOrder map[*CodeTACANChannelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTACANChannelTypeCreateCallback GongOnAfterCreateInterface[CodeTACANChannelType]
	OnAfterCodeTACANChannelTypeUpdateCallback GongOnAfterUpdateInterface[CodeTACANChannelType]
	OnAfterCodeTACANChannelTypeDeleteCallback GongOnAfterDeleteInterface[CodeTACANChannelType]

	CodeTelecomNetworkTypes                map[*CodeTelecomNetworkType]struct{}
	CodeTelecomNetworkTypes_instance       map[*CodeTelecomNetworkType]*CodeTelecomNetworkType
	CodeTelecomNetworkTypes_mapString      map[string]*CodeTelecomNetworkType
	CodeTelecomNetworkTypeOrder            uint
	CodeTelecomNetworkType_stagedOrder     map[*CodeTelecomNetworkType]uint
	CodeTelecomNetworkType_orderStaged     map[uint]*CodeTelecomNetworkType
	CodeTelecomNetworkTypes_reference      map[*CodeTelecomNetworkType]*CodeTelecomNetworkType
	CodeTelecomNetworkTypes_referenceOrder map[*CodeTelecomNetworkType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTelecomNetworkTypeCreateCallback GongOnAfterCreateInterface[CodeTelecomNetworkType]
	OnAfterCodeTelecomNetworkTypeUpdateCallback GongOnAfterUpdateInterface[CodeTelecomNetworkType]
	OnAfterCodeTelecomNetworkTypeDeleteCallback GongOnAfterDeleteInterface[CodeTelecomNetworkType]

	CodeTimeEventCombinationTypes                map[*CodeTimeEventCombinationType]struct{}
	CodeTimeEventCombinationTypes_instance       map[*CodeTimeEventCombinationType]*CodeTimeEventCombinationType
	CodeTimeEventCombinationTypes_mapString      map[string]*CodeTimeEventCombinationType
	CodeTimeEventCombinationTypeOrder            uint
	CodeTimeEventCombinationType_stagedOrder     map[*CodeTimeEventCombinationType]uint
	CodeTimeEventCombinationType_orderStaged     map[uint]*CodeTimeEventCombinationType
	CodeTimeEventCombinationTypes_reference      map[*CodeTimeEventCombinationType]*CodeTimeEventCombinationType
	CodeTimeEventCombinationTypes_referenceOrder map[*CodeTimeEventCombinationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTimeEventCombinationTypeCreateCallback GongOnAfterCreateInterface[CodeTimeEventCombinationType]
	OnAfterCodeTimeEventCombinationTypeUpdateCallback GongOnAfterUpdateInterface[CodeTimeEventCombinationType]
	OnAfterCodeTimeEventCombinationTypeDeleteCallback GongOnAfterDeleteInterface[CodeTimeEventCombinationType]

	CodeTimeEventTypes                map[*CodeTimeEventType]struct{}
	CodeTimeEventTypes_instance       map[*CodeTimeEventType]*CodeTimeEventType
	CodeTimeEventTypes_mapString      map[string]*CodeTimeEventType
	CodeTimeEventTypeOrder            uint
	CodeTimeEventType_stagedOrder     map[*CodeTimeEventType]uint
	CodeTimeEventType_orderStaged     map[uint]*CodeTimeEventType
	CodeTimeEventTypes_reference      map[*CodeTimeEventType]*CodeTimeEventType
	CodeTimeEventTypes_referenceOrder map[*CodeTimeEventType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTimeEventTypeCreateCallback GongOnAfterCreateInterface[CodeTimeEventType]
	OnAfterCodeTimeEventTypeUpdateCallback GongOnAfterUpdateInterface[CodeTimeEventType]
	OnAfterCodeTimeEventTypeDeleteCallback GongOnAfterDeleteInterface[CodeTimeEventType]

	CodeTimeReferenceTypes                map[*CodeTimeReferenceType]struct{}
	CodeTimeReferenceTypes_instance       map[*CodeTimeReferenceType]*CodeTimeReferenceType
	CodeTimeReferenceTypes_mapString      map[string]*CodeTimeReferenceType
	CodeTimeReferenceTypeOrder            uint
	CodeTimeReferenceType_stagedOrder     map[*CodeTimeReferenceType]uint
	CodeTimeReferenceType_orderStaged     map[uint]*CodeTimeReferenceType
	CodeTimeReferenceTypes_reference      map[*CodeTimeReferenceType]*CodeTimeReferenceType
	CodeTimeReferenceTypes_referenceOrder map[*CodeTimeReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTimeReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeTimeReferenceType]
	OnAfterCodeTimeReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeTimeReferenceType]
	OnAfterCodeTimeReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeTimeReferenceType]

	CodeTransponderTypes                map[*CodeTransponderType]struct{}
	CodeTransponderTypes_instance       map[*CodeTransponderType]*CodeTransponderType
	CodeTransponderTypes_mapString      map[string]*CodeTransponderType
	CodeTransponderTypeOrder            uint
	CodeTransponderType_stagedOrder     map[*CodeTransponderType]uint
	CodeTransponderType_orderStaged     map[uint]*CodeTransponderType
	CodeTransponderTypes_reference      map[*CodeTransponderType]*CodeTransponderType
	CodeTransponderTypes_referenceOrder map[*CodeTransponderType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTransponderTypeCreateCallback GongOnAfterCreateInterface[CodeTransponderType]
	OnAfterCodeTransponderTypeUpdateCallback GongOnAfterUpdateInterface[CodeTransponderType]
	OnAfterCodeTransponderTypeDeleteCallback GongOnAfterDeleteInterface[CodeTransponderType]

	CodeUpperAlphaTypes                map[*CodeUpperAlphaType]struct{}
	CodeUpperAlphaTypes_instance       map[*CodeUpperAlphaType]*CodeUpperAlphaType
	CodeUpperAlphaTypes_mapString      map[string]*CodeUpperAlphaType
	CodeUpperAlphaTypeOrder            uint
	CodeUpperAlphaType_stagedOrder     map[*CodeUpperAlphaType]uint
	CodeUpperAlphaType_orderStaged     map[uint]*CodeUpperAlphaType
	CodeUpperAlphaTypes_reference      map[*CodeUpperAlphaType]*CodeUpperAlphaType
	CodeUpperAlphaTypes_referenceOrder map[*CodeUpperAlphaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeUpperAlphaTypeCreateCallback GongOnAfterCreateInterface[CodeUpperAlphaType]
	OnAfterCodeUpperAlphaTypeUpdateCallback GongOnAfterUpdateInterface[CodeUpperAlphaType]
	OnAfterCodeUpperAlphaTypeDeleteCallback GongOnAfterDeleteInterface[CodeUpperAlphaType]

	CodeValueInterpretationTypes                map[*CodeValueInterpretationType]struct{}
	CodeValueInterpretationTypes_instance       map[*CodeValueInterpretationType]*CodeValueInterpretationType
	CodeValueInterpretationTypes_mapString      map[string]*CodeValueInterpretationType
	CodeValueInterpretationTypeOrder            uint
	CodeValueInterpretationType_stagedOrder     map[*CodeValueInterpretationType]uint
	CodeValueInterpretationType_orderStaged     map[uint]*CodeValueInterpretationType
	CodeValueInterpretationTypes_reference      map[*CodeValueInterpretationType]*CodeValueInterpretationType
	CodeValueInterpretationTypes_referenceOrder map[*CodeValueInterpretationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeValueInterpretationTypeCreateCallback GongOnAfterCreateInterface[CodeValueInterpretationType]
	OnAfterCodeValueInterpretationTypeUpdateCallback GongOnAfterUpdateInterface[CodeValueInterpretationType]
	OnAfterCodeValueInterpretationTypeDeleteCallback GongOnAfterDeleteInterface[CodeValueInterpretationType]

	CodeVerticalDatumTypes                map[*CodeVerticalDatumType]struct{}
	CodeVerticalDatumTypes_instance       map[*CodeVerticalDatumType]*CodeVerticalDatumType
	CodeVerticalDatumTypes_mapString      map[string]*CodeVerticalDatumType
	CodeVerticalDatumTypeOrder            uint
	CodeVerticalDatumType_stagedOrder     map[*CodeVerticalDatumType]uint
	CodeVerticalDatumType_orderStaged     map[uint]*CodeVerticalDatumType
	CodeVerticalDatumTypes_reference      map[*CodeVerticalDatumType]*CodeVerticalDatumType
	CodeVerticalDatumTypes_referenceOrder map[*CodeVerticalDatumType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVerticalDatumTypeCreateCallback GongOnAfterCreateInterface[CodeVerticalDatumType]
	OnAfterCodeVerticalDatumTypeUpdateCallback GongOnAfterUpdateInterface[CodeVerticalDatumType]
	OnAfterCodeVerticalDatumTypeDeleteCallback GongOnAfterDeleteInterface[CodeVerticalDatumType]

	CodeVerticalReferenceTypes                map[*CodeVerticalReferenceType]struct{}
	CodeVerticalReferenceTypes_instance       map[*CodeVerticalReferenceType]*CodeVerticalReferenceType
	CodeVerticalReferenceTypes_mapString      map[string]*CodeVerticalReferenceType
	CodeVerticalReferenceTypeOrder            uint
	CodeVerticalReferenceType_stagedOrder     map[*CodeVerticalReferenceType]uint
	CodeVerticalReferenceType_orderStaged     map[uint]*CodeVerticalReferenceType
	CodeVerticalReferenceTypes_reference      map[*CodeVerticalReferenceType]*CodeVerticalReferenceType
	CodeVerticalReferenceTypes_referenceOrder map[*CodeVerticalReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVerticalReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeVerticalReferenceType]
	OnAfterCodeVerticalReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeVerticalReferenceType]
	OnAfterCodeVerticalReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeVerticalReferenceType]

	CodeWakeTurbulenceTypes                map[*CodeWakeTurbulenceType]struct{}
	CodeWakeTurbulenceTypes_instance       map[*CodeWakeTurbulenceType]*CodeWakeTurbulenceType
	CodeWakeTurbulenceTypes_mapString      map[string]*CodeWakeTurbulenceType
	CodeWakeTurbulenceTypeOrder            uint
	CodeWakeTurbulenceType_stagedOrder     map[*CodeWakeTurbulenceType]uint
	CodeWakeTurbulenceType_orderStaged     map[uint]*CodeWakeTurbulenceType
	CodeWakeTurbulenceTypes_reference      map[*CodeWakeTurbulenceType]*CodeWakeTurbulenceType
	CodeWakeTurbulenceTypes_referenceOrder map[*CodeWakeTurbulenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeWakeTurbulenceTypeCreateCallback GongOnAfterCreateInterface[CodeWakeTurbulenceType]
	OnAfterCodeWakeTurbulenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeWakeTurbulenceType]
	OnAfterCodeWakeTurbulenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeWakeTurbulenceType]

	CodeYesNoTypes                map[*CodeYesNoType]struct{}
	CodeYesNoTypes_instance       map[*CodeYesNoType]*CodeYesNoType
	CodeYesNoTypes_mapString      map[string]*CodeYesNoType
	CodeYesNoTypeOrder            uint
	CodeYesNoType_stagedOrder     map[*CodeYesNoType]uint
	CodeYesNoType_orderStaged     map[uint]*CodeYesNoType
	CodeYesNoTypes_reference      map[*CodeYesNoType]*CodeYesNoType
	CodeYesNoTypes_referenceOrder map[*CodeYesNoType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeYesNoTypeCreateCallback GongOnAfterCreateInterface[CodeYesNoType]
	OnAfterCodeYesNoTypeUpdateCallback GongOnAfterUpdateInterface[CodeYesNoType]
	OnAfterCodeYesNoTypeDeleteCallback GongOnAfterDeleteInterface[CodeYesNoType]

	DateMonthDayTypes                map[*DateMonthDayType]struct{}
	DateMonthDayTypes_instance       map[*DateMonthDayType]*DateMonthDayType
	DateMonthDayTypes_mapString      map[string]*DateMonthDayType
	DateMonthDayTypeOrder            uint
	DateMonthDayType_stagedOrder     map[*DateMonthDayType]uint
	DateMonthDayType_orderStaged     map[uint]*DateMonthDayType
	DateMonthDayTypes_reference      map[*DateMonthDayType]*DateMonthDayType
	DateMonthDayTypes_referenceOrder map[*DateMonthDayType]uint

	// insertion point for slice of pointers maps
	OnAfterDateMonthDayTypeCreateCallback GongOnAfterCreateInterface[DateMonthDayType]
	OnAfterDateMonthDayTypeUpdateCallback GongOnAfterUpdateInterface[DateMonthDayType]
	OnAfterDateMonthDayTypeDeleteCallback GongOnAfterDeleteInterface[DateMonthDayType]

	NoNumberTypes                map[*NoNumberType]struct{}
	NoNumberTypes_instance       map[*NoNumberType]*NoNumberType
	NoNumberTypes_mapString      map[string]*NoNumberType
	NoNumberTypeOrder            uint
	NoNumberType_stagedOrder     map[*NoNumberType]uint
	NoNumberType_orderStaged     map[uint]*NoNumberType
	NoNumberTypes_reference      map[*NoNumberType]*NoNumberType
	NoNumberTypes_referenceOrder map[*NoNumberType]uint

	// insertion point for slice of pointers maps
	OnAfterNoNumberTypeCreateCallback GongOnAfterCreateInterface[NoNumberType]
	OnAfterNoNumberTypeUpdateCallback GongOnAfterUpdateInterface[NoNumberType]
	OnAfterNoNumberTypeDeleteCallback GongOnAfterDeleteInterface[NoNumberType]

	NoSequenceTypes                map[*NoSequenceType]struct{}
	NoSequenceTypes_instance       map[*NoSequenceType]*NoSequenceType
	NoSequenceTypes_mapString      map[string]*NoSequenceType
	NoSequenceTypeOrder            uint
	NoSequenceType_stagedOrder     map[*NoSequenceType]uint
	NoSequenceType_orderStaged     map[uint]*NoSequenceType
	NoSequenceTypes_reference      map[*NoSequenceType]*NoSequenceType
	NoSequenceTypes_referenceOrder map[*NoSequenceType]uint

	// insertion point for slice of pointers maps
	OnAfterNoSequenceTypeCreateCallback GongOnAfterCreateInterface[NoSequenceType]
	OnAfterNoSequenceTypeUpdateCallback GongOnAfterUpdateInterface[NoSequenceType]
	OnAfterNoSequenceTypeDeleteCallback GongOnAfterDeleteInterface[NoSequenceType]

	TextAddressTypes                map[*TextAddressType]struct{}
	TextAddressTypes_instance       map[*TextAddressType]*TextAddressType
	TextAddressTypes_mapString      map[string]*TextAddressType
	TextAddressTypeOrder            uint
	TextAddressType_stagedOrder     map[*TextAddressType]uint
	TextAddressType_orderStaged     map[uint]*TextAddressType
	TextAddressTypes_reference      map[*TextAddressType]*TextAddressType
	TextAddressTypes_referenceOrder map[*TextAddressType]uint

	// insertion point for slice of pointers maps
	OnAfterTextAddressTypeCreateCallback GongOnAfterCreateInterface[TextAddressType]
	OnAfterTextAddressTypeUpdateCallback GongOnAfterUpdateInterface[TextAddressType]
	OnAfterTextAddressTypeDeleteCallback GongOnAfterDeleteInterface[TextAddressType]

	TextDesignatorTypes                map[*TextDesignatorType]struct{}
	TextDesignatorTypes_instance       map[*TextDesignatorType]*TextDesignatorType
	TextDesignatorTypes_mapString      map[string]*TextDesignatorType
	TextDesignatorTypeOrder            uint
	TextDesignatorType_stagedOrder     map[*TextDesignatorType]uint
	TextDesignatorType_orderStaged     map[uint]*TextDesignatorType
	TextDesignatorTypes_reference      map[*TextDesignatorType]*TextDesignatorType
	TextDesignatorTypes_referenceOrder map[*TextDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterTextDesignatorTypeCreateCallback GongOnAfterCreateInterface[TextDesignatorType]
	OnAfterTextDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[TextDesignatorType]
	OnAfterTextDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[TextDesignatorType]

	TextInstructionTypes                map[*TextInstructionType]struct{}
	TextInstructionTypes_instance       map[*TextInstructionType]*TextInstructionType
	TextInstructionTypes_mapString      map[string]*TextInstructionType
	TextInstructionTypeOrder            uint
	TextInstructionType_stagedOrder     map[*TextInstructionType]uint
	TextInstructionType_orderStaged     map[uint]*TextInstructionType
	TextInstructionTypes_reference      map[*TextInstructionType]*TextInstructionType
	TextInstructionTypes_referenceOrder map[*TextInstructionType]uint

	// insertion point for slice of pointers maps
	OnAfterTextInstructionTypeCreateCallback GongOnAfterCreateInterface[TextInstructionType]
	OnAfterTextInstructionTypeUpdateCallback GongOnAfterUpdateInterface[TextInstructionType]
	OnAfterTextInstructionTypeDeleteCallback GongOnAfterDeleteInterface[TextInstructionType]

	TextNameTypes                map[*TextNameType]struct{}
	TextNameTypes_instance       map[*TextNameType]*TextNameType
	TextNameTypes_mapString      map[string]*TextNameType
	TextNameTypeOrder            uint
	TextNameType_stagedOrder     map[*TextNameType]uint
	TextNameType_orderStaged     map[uint]*TextNameType
	TextNameTypes_reference      map[*TextNameType]*TextNameType
	TextNameTypes_referenceOrder map[*TextNameType]uint

	// insertion point for slice of pointers maps
	OnAfterTextNameTypeCreateCallback GongOnAfterCreateInterface[TextNameType]
	OnAfterTextNameTypeUpdateCallback GongOnAfterUpdateInterface[TextNameType]
	OnAfterTextNameTypeDeleteCallback GongOnAfterDeleteInterface[TextNameType]

	TextNoteTypes                map[*TextNoteType]struct{}
	TextNoteTypes_instance       map[*TextNoteType]*TextNoteType
	TextNoteTypes_mapString      map[string]*TextNoteType
	TextNoteTypeOrder            uint
	TextNoteType_stagedOrder     map[*TextNoteType]uint
	TextNoteType_orderStaged     map[uint]*TextNoteType
	TextNoteTypes_reference      map[*TextNoteType]*TextNoteType
	TextNoteTypes_referenceOrder map[*TextNoteType]uint

	// insertion point for slice of pointers maps
	OnAfterTextNoteTypeCreateCallback GongOnAfterCreateInterface[TextNoteType]
	OnAfterTextNoteTypeUpdateCallback GongOnAfterUpdateInterface[TextNoteType]
	OnAfterTextNoteTypeDeleteCallback GongOnAfterDeleteInterface[TextNoteType]

	TextPhoneTypes                map[*TextPhoneType]struct{}
	TextPhoneTypes_instance       map[*TextPhoneType]*TextPhoneType
	TextPhoneTypes_mapString      map[string]*TextPhoneType
	TextPhoneTypeOrder            uint
	TextPhoneType_stagedOrder     map[*TextPhoneType]uint
	TextPhoneType_orderStaged     map[uint]*TextPhoneType
	TextPhoneTypes_reference      map[*TextPhoneType]*TextPhoneType
	TextPhoneTypes_referenceOrder map[*TextPhoneType]uint

	// insertion point for slice of pointers maps
	OnAfterTextPhoneTypeCreateCallback GongOnAfterCreateInterface[TextPhoneType]
	OnAfterTextPhoneTypeUpdateCallback GongOnAfterUpdateInterface[TextPhoneType]
	OnAfterTextPhoneTypeDeleteCallback GongOnAfterDeleteInterface[TextPhoneType]

	TextPropertyNameTypes                map[*TextPropertyNameType]struct{}
	TextPropertyNameTypes_instance       map[*TextPropertyNameType]*TextPropertyNameType
	TextPropertyNameTypes_mapString      map[string]*TextPropertyNameType
	TextPropertyNameTypeOrder            uint
	TextPropertyNameType_stagedOrder     map[*TextPropertyNameType]uint
	TextPropertyNameType_orderStaged     map[uint]*TextPropertyNameType
	TextPropertyNameTypes_reference      map[*TextPropertyNameType]*TextPropertyNameType
	TextPropertyNameTypes_referenceOrder map[*TextPropertyNameType]uint

	// insertion point for slice of pointers maps
	OnAfterTextPropertyNameTypeCreateCallback GongOnAfterCreateInterface[TextPropertyNameType]
	OnAfterTextPropertyNameTypeUpdateCallback GongOnAfterUpdateInterface[TextPropertyNameType]
	OnAfterTextPropertyNameTypeDeleteCallback GongOnAfterDeleteInterface[TextPropertyNameType]

	TimeTypes                map[*TimeType]struct{}
	TimeTypes_instance       map[*TimeType]*TimeType
	TimeTypes_mapString      map[string]*TimeType
	TimeTypeOrder            uint
	TimeType_stagedOrder     map[*TimeType]uint
	TimeType_orderStaged     map[uint]*TimeType
	TimeTypes_reference      map[*TimeType]*TimeType
	TimeTypes_referenceOrder map[*TimeType]uint

	// insertion point for slice of pointers maps
	OnAfterTimeTypeCreateCallback GongOnAfterCreateInterface[TimeType]
	OnAfterTimeTypeUpdateCallback GongOnAfterUpdateInterface[TimeType]
	OnAfterTimeTypeDeleteCallback GongOnAfterDeleteInterface[TimeType]

	ValAngleTypes                map[*ValAngleType]struct{}
	ValAngleTypes_instance       map[*ValAngleType]*ValAngleType
	ValAngleTypes_mapString      map[string]*ValAngleType
	ValAngleTypeOrder            uint
	ValAngleType_stagedOrder     map[*ValAngleType]uint
	ValAngleType_orderStaged     map[uint]*ValAngleType
	ValAngleTypes_reference      map[*ValAngleType]*ValAngleType
	ValAngleTypes_referenceOrder map[*ValAngleType]uint

	// insertion point for slice of pointers maps
	OnAfterValAngleTypeCreateCallback GongOnAfterCreateInterface[ValAngleType]
	OnAfterValAngleTypeUpdateCallback GongOnAfterUpdateInterface[ValAngleType]
	OnAfterValAngleTypeDeleteCallback GongOnAfterDeleteInterface[ValAngleType]

	ValBearingTypes                map[*ValBearingType]struct{}
	ValBearingTypes_instance       map[*ValBearingType]*ValBearingType
	ValBearingTypes_mapString      map[string]*ValBearingType
	ValBearingTypeOrder            uint
	ValBearingType_stagedOrder     map[*ValBearingType]uint
	ValBearingType_orderStaged     map[uint]*ValBearingType
	ValBearingTypes_reference      map[*ValBearingType]*ValBearingType
	ValBearingTypes_referenceOrder map[*ValBearingType]uint

	// insertion point for slice of pointers maps
	OnAfterValBearingTypeCreateCallback GongOnAfterCreateInterface[ValBearingType]
	OnAfterValBearingTypeUpdateCallback GongOnAfterUpdateInterface[ValBearingType]
	OnAfterValBearingTypeDeleteCallback GongOnAfterDeleteInterface[ValBearingType]

	ValDistanceSignedTypes                map[*ValDistanceSignedType]struct{}
	ValDistanceSignedTypes_instance       map[*ValDistanceSignedType]*ValDistanceSignedType
	ValDistanceSignedTypes_mapString      map[string]*ValDistanceSignedType
	ValDistanceSignedTypeOrder            uint
	ValDistanceSignedType_stagedOrder     map[*ValDistanceSignedType]uint
	ValDistanceSignedType_orderStaged     map[uint]*ValDistanceSignedType
	ValDistanceSignedTypes_reference      map[*ValDistanceSignedType]*ValDistanceSignedType
	ValDistanceSignedTypes_referenceOrder map[*ValDistanceSignedType]uint

	// insertion point for slice of pointers maps
	OnAfterValDistanceSignedTypeCreateCallback GongOnAfterCreateInterface[ValDistanceSignedType]
	OnAfterValDistanceSignedTypeUpdateCallback GongOnAfterUpdateInterface[ValDistanceSignedType]
	OnAfterValDistanceSignedTypeDeleteCallback GongOnAfterDeleteInterface[ValDistanceSignedType]

	ValDistanceTypes                map[*ValDistanceType]struct{}
	ValDistanceTypes_instance       map[*ValDistanceType]*ValDistanceType
	ValDistanceTypes_mapString      map[string]*ValDistanceType
	ValDistanceTypeOrder            uint
	ValDistanceType_stagedOrder     map[*ValDistanceType]uint
	ValDistanceType_orderStaged     map[uint]*ValDistanceType
	ValDistanceTypes_reference      map[*ValDistanceType]*ValDistanceType
	ValDistanceTypes_referenceOrder map[*ValDistanceType]uint

	// insertion point for slice of pointers maps
	OnAfterValDistanceTypeCreateCallback GongOnAfterCreateInterface[ValDistanceType]
	OnAfterValDistanceTypeUpdateCallback GongOnAfterUpdateInterface[ValDistanceType]
	OnAfterValDistanceTypeDeleteCallback GongOnAfterDeleteInterface[ValDistanceType]

	ValDistanceVerticalTypes                map[*ValDistanceVerticalType]struct{}
	ValDistanceVerticalTypes_instance       map[*ValDistanceVerticalType]*ValDistanceVerticalType
	ValDistanceVerticalTypes_mapString      map[string]*ValDistanceVerticalType
	ValDistanceVerticalTypeOrder            uint
	ValDistanceVerticalType_stagedOrder     map[*ValDistanceVerticalType]uint
	ValDistanceVerticalType_orderStaged     map[uint]*ValDistanceVerticalType
	ValDistanceVerticalTypes_reference      map[*ValDistanceVerticalType]*ValDistanceVerticalType
	ValDistanceVerticalTypes_referenceOrder map[*ValDistanceVerticalType]uint

	// insertion point for slice of pointers maps
	OnAfterValDistanceVerticalTypeCreateCallback GongOnAfterCreateInterface[ValDistanceVerticalType]
	OnAfterValDistanceVerticalTypeUpdateCallback GongOnAfterUpdateInterface[ValDistanceVerticalType]
	OnAfterValDistanceVerticalTypeDeleteCallback GongOnAfterDeleteInterface[ValDistanceVerticalType]

	ValDurationTypes                map[*ValDurationType]struct{}
	ValDurationTypes_instance       map[*ValDurationType]*ValDurationType
	ValDurationTypes_mapString      map[string]*ValDurationType
	ValDurationTypeOrder            uint
	ValDurationType_stagedOrder     map[*ValDurationType]uint
	ValDurationType_orderStaged     map[uint]*ValDurationType
	ValDurationTypes_reference      map[*ValDurationType]*ValDurationType
	ValDurationTypes_referenceOrder map[*ValDurationType]uint

	// insertion point for slice of pointers maps
	OnAfterValDurationTypeCreateCallback GongOnAfterCreateInterface[ValDurationType]
	OnAfterValDurationTypeUpdateCallback GongOnAfterUpdateInterface[ValDurationType]
	OnAfterValDurationTypeDeleteCallback GongOnAfterDeleteInterface[ValDurationType]

	ValSlopeTypes                map[*ValSlopeType]struct{}
	ValSlopeTypes_instance       map[*ValSlopeType]*ValSlopeType
	ValSlopeTypes_mapString      map[string]*ValSlopeType
	ValSlopeTypeOrder            uint
	ValSlopeType_stagedOrder     map[*ValSlopeType]uint
	ValSlopeType_orderStaged     map[uint]*ValSlopeType
	ValSlopeTypes_reference      map[*ValSlopeType]*ValSlopeType
	ValSlopeTypes_referenceOrder map[*ValSlopeType]uint

	// insertion point for slice of pointers maps
	OnAfterValSlopeTypeCreateCallback GongOnAfterCreateInterface[ValSlopeType]
	OnAfterValSlopeTypeUpdateCallback GongOnAfterUpdateInterface[ValSlopeType]
	OnAfterValSlopeTypeDeleteCallback GongOnAfterDeleteInterface[ValSlopeType]

	ValSpeedTypes                map[*ValSpeedType]struct{}
	ValSpeedTypes_instance       map[*ValSpeedType]*ValSpeedType
	ValSpeedTypes_mapString      map[string]*ValSpeedType
	ValSpeedTypeOrder            uint
	ValSpeedType_stagedOrder     map[*ValSpeedType]uint
	ValSpeedType_orderStaged     map[uint]*ValSpeedType
	ValSpeedTypes_reference      map[*ValSpeedType]*ValSpeedType
	ValSpeedTypes_referenceOrder map[*ValSpeedType]uint

	// insertion point for slice of pointers maps
	OnAfterValSpeedTypeCreateCallback GongOnAfterCreateInterface[ValSpeedType]
	OnAfterValSpeedTypeUpdateCallback GongOnAfterUpdateInterface[ValSpeedType]
	OnAfterValSpeedTypeDeleteCallback GongOnAfterDeleteInterface[ValSpeedType]

	ValWeightTypes                map[*ValWeightType]struct{}
	ValWeightTypes_instance       map[*ValWeightType]*ValWeightType
	ValWeightTypes_mapString      map[string]*ValWeightType
	ValWeightTypeOrder            uint
	ValWeightType_stagedOrder     map[*ValWeightType]uint
	ValWeightType_orderStaged     map[uint]*ValWeightType
	ValWeightTypes_reference      map[*ValWeightType]*ValWeightType
	ValWeightTypes_referenceOrder map[*ValWeightType]uint

	// insertion point for slice of pointers maps
	OnAfterValWeightTypeCreateCallback GongOnAfterCreateInterface[ValWeightType]
	OnAfterValWeightTypeUpdateCallback GongOnAfterUpdateInterface[ValWeightType]
	OnAfterValWeightTypeDeleteCallback GongOnAfterDeleteInterface[ValWeightType]

	XHTMLTypes                map[*XHTMLType]struct{}
	XHTMLTypes_instance       map[*XHTMLType]*XHTMLType
	XHTMLTypes_mapString      map[string]*XHTMLType
	XHTMLTypeOrder            uint
	XHTMLType_stagedOrder     map[*XHTMLType]uint
	XHTMLType_orderStaged     map[uint]*XHTMLType
	XHTMLTypes_reference      map[*XHTMLType]*XHTMLType
	XHTMLTypes_referenceOrder map[*XHTMLType]uint

	// insertion point for slice of pointers maps
	OnAfterXHTMLTypeCreateCallback GongOnAfterCreateInterface[XHTMLType]
	OnAfterXHTMLTypeUpdateCallback GongOnAfterUpdateInterface[XHTMLType]
	OnAfterXHTMLTypeDeleteCallback GongOnAfterDeleteInterface[XHTMLType]

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
	__gong__clearReferences(&stage.CodeATCReportingTypes_reference, &stage.CodeATCReportingTypes_instance, &stage.CodeATCReportingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAerialRefuellingPointTypes_reference, &stage.CodeAerialRefuellingPointTypes_instance, &stage.CodeAerialRefuellingPointTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAerialRefuellingPrefixTypes_reference, &stage.CodeAerialRefuellingPrefixTypes_instance, &stage.CodeAerialRefuellingPrefixTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAerialRefuellingTypes_reference, &stage.CodeAerialRefuellingTypes_instance, &stage.CodeAerialRefuellingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_instance, &stage.CodeAircraftCategoryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_instance, &stage.CodeAircraftEngineNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_instance, &stage.CodeAircraftEngineTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_instance, &stage.CodeAircraftICAOTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_instance, &stage.CodeAircraftTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_instance, &stage.CodeAircraftWingspanClassTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeAdjustmentTypes_reference, &stage.CodeAltitudeAdjustmentTypes_instance, &stage.CodeAltitudeAdjustmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_instance, &stage.CodeAltitudeUseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAuthorityTypes_reference, &stage.CodeAuthorityTypes_instance, &stage.CodeAuthorityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCardinalDirectionTypes_reference, &stage.CodeCardinalDirectionTypes_instance, &stage.CodeCardinalDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_instance, &stage.CodeCommunicationModeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeComparisonTypes_reference, &stage.CodeComparisonTypes_instance, &stage.CodeComparisonTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCourseTypes_reference, &stage.CodeCourseTypes_instance, &stage.CodeCourseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_instance, &stage.CodeDirectionTurnTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionTypes_reference, &stage.CodeDirectionTypes_instance, &stage.CodeDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_instance, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightOriginTypes_reference, &stage.CodeFlightOriginTypes_instance, &stage.CodeFlightOriginTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightPurposeTypes_reference, &stage.CodeFlightPurposeTypes_instance, &stage.CodeFlightPurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightRestrictionDesignatorTypes_reference, &stage.CodeFlightRestrictionDesignatorTypes_instance, &stage.CodeFlightRestrictionDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightRestrictionTypes_reference, &stage.CodeFlightRestrictionTypes_instance, &stage.CodeFlightRestrictionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_instance, &stage.CodeFlightRuleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightStatusTypes_reference, &stage.CodeFlightStatusTypes_instance, &stage.CodeFlightStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightTypes_reference, &stage.CodeFlightTypes_instance, &stage.CodeFlightTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlowConditionOperationTypes_reference, &stage.CodeFlowConditionOperationTypes_instance, &stage.CodeFlowConditionOperationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFreeFlightTypes_reference, &stage.CodeFreeFlightTypes_instance, &stage.CodeFreeFlightTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLevelTypes_reference, &stage.CodeLevelTypes_instance, &stage.CodeLevelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLocationQualifierTypes_reference, &stage.CodeLocationQualifierTypes_instance, &stage.CodeLocationQualifierTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMeteoConditionsTypes_reference, &stage.CodeMeteoConditionsTypes_instance, &stage.CodeMeteoConditionsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryRoutePointTypes_reference, &stage.CodeMilitaryRoutePointTypes_instance, &stage.CodeMilitaryRoutePointTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryStatusTypes_reference, &stage.CodeMilitaryStatusTypes_instance, &stage.CodeMilitaryStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryTrainingTypes_reference, &stage.CodeMilitaryTrainingTypes_instance, &stage.CodeMilitaryTrainingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_instance, &stage.CodeNavigationEquipmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_instance, &stage.CodeNavigationSpecificationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_instance, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstructionIdSurfaceZoneTypes_reference, &stage.CodeObstructionIdSurfaceZoneTypes_instance, &stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRNPTypes_reference, &stage.CodeRNPTypes_instance, &stage.CodeRNPTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMPointRoleTypes_reference, &stage.CodeRVSMPointRoleTypes_instance, &stage.CodeRVSMPointRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_instance, &stage.CodeRVSMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeReferenceRoleTypes_reference, &stage.CodeReferenceRoleTypes_instance, &stage.CodeReferenceRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteAvailabilityTypes_reference, &stage.CodeRouteAvailabilityTypes_instance, &stage.CodeRouteAvailabilityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteDesignatorLetterTypes_reference, &stage.CodeRouteDesignatorLetterTypes_instance, &stage.CodeRouteDesignatorLetterTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteDesignatorPrefixTypes_reference, &stage.CodeRouteDesignatorPrefixTypes_instance, &stage.CodeRouteDesignatorPrefixTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteDesignatorSuffixTypes_reference, &stage.CodeRouteDesignatorSuffixTypes_instance, &stage.CodeRouteDesignatorSuffixTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteNavigationTypes_reference, &stage.CodeRouteNavigationTypes_instance, &stage.CodeRouteNavigationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteOriginTypes_reference, &stage.CodeRouteOriginTypes_instance, &stage.CodeRouteOriginTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteSegmentPathTypes_reference, &stage.CodeRouteSegmentPathTypes_instance, &stage.CodeRouteSegmentPathTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteTypes_reference, &stage.CodeRouteTypes_instance, &stage.CodeRouteTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRuleProcedureTitleTypes_reference, &stage.CodeRuleProcedureTitleTypes_instance, &stage.CodeRuleProcedureTitleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRuleProcedureTypes_reference, &stage.CodeRuleProcedureTypes_instance, &stage.CodeRuleProcedureTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSideTypes_reference, &stage.CodeSideTypes_instance, &stage.CodeSideTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpeedReferenceTypes_reference, &stage.CodeSpeedReferenceTypes_instance, &stage.CodeSpeedReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTACANChannelTypes_reference, &stage.CodeTACANChannelTypes_instance, &stage.CodeTACANChannelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_instance, &stage.CodeTelecomNetworkTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_instance, &stage.CodeTransponderTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeUpperAlphaTypes_reference, &stage.CodeUpperAlphaTypes_instance, &stage.CodeUpperAlphaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_instance, &stage.CodeValueInterpretationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_instance, &stage.CodeVerticalReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_instance, &stage.CodeWakeTurbulenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.NoNumberTypes_reference, &stage.NoNumberTypes_instance, &stage.NoNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_instance, &stage.NoSequenceTypes_referenceOrder)

	__gong__clearReferences(&stage.TextAddressTypes_reference, &stage.TextAddressTypes_instance, &stage.TextAddressTypes_referenceOrder)

	__gong__clearReferences(&stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_instance, &stage.TextDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_instance, &stage.TextInstructionTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNameTypes_reference, &stage.TextNameTypes_instance, &stage.TextNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNoteTypes_reference, &stage.TextNoteTypes_instance, &stage.TextNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPhoneTypes_reference, &stage.TextPhoneTypes_instance, &stage.TextPhoneTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_instance, &stage.TextPropertyNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValAngleTypes_reference, &stage.ValAngleTypes_instance, &stage.ValAngleTypes_referenceOrder)

	__gong__clearReferences(&stage.ValBearingTypes_reference, &stage.ValBearingTypes_instance, &stage.ValBearingTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValSlopeTypes_reference, &stage.ValSlopeTypes_instance, &stage.ValSlopeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_instance, &stage.ValSpeedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValWeightTypes_reference, &stage.ValWeightTypes_instance, &stage.ValWeightTypes_referenceOrder)

	__gong__clearReferences(&stage.XHTMLTypes_reference, &stage.XHTMLTypes_instance, &stage.XHTMLTypes_referenceOrder)

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
	stage.CodeATCReportingTypeOrder = __gong__recomputeOrder(stage.CodeATCReportingType_stagedOrder)

	stage.CodeAerialRefuellingPointTypeOrder = __gong__recomputeOrder(stage.CodeAerialRefuellingPointType_stagedOrder)

	stage.CodeAerialRefuellingPrefixTypeOrder = __gong__recomputeOrder(stage.CodeAerialRefuellingPrefixType_stagedOrder)

	stage.CodeAerialRefuellingTypeOrder = __gong__recomputeOrder(stage.CodeAerialRefuellingType_stagedOrder)

	stage.CodeAircraftCategoryTypeOrder = __gong__recomputeOrder(stage.CodeAircraftCategoryType_stagedOrder)

	stage.CodeAircraftEngineNumberTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineNumberType_stagedOrder)

	stage.CodeAircraftEngineTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineType_stagedOrder)

	stage.CodeAircraftICAOTypeOrder = __gong__recomputeOrder(stage.CodeAircraftICAOType_stagedOrder)

	stage.CodeAircraftTypeOrder = __gong__recomputeOrder(stage.CodeAircraftType_stagedOrder)

	stage.CodeAircraftWingspanClassTypeOrder = __gong__recomputeOrder(stage.CodeAircraftWingspanClassType_stagedOrder)

	stage.CodeAltitudeAdjustmentTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeAdjustmentType_stagedOrder)

	stage.CodeAltitudeUseTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeUseType_stagedOrder)

	stage.CodeAuthorityTypeOrder = __gong__recomputeOrder(stage.CodeAuthorityType_stagedOrder)

	stage.CodeCardinalDirectionTypeOrder = __gong__recomputeOrder(stage.CodeCardinalDirectionType_stagedOrder)

	stage.CodeCommunicationModeTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationModeType_stagedOrder)

	stage.CodeComparisonTypeOrder = __gong__recomputeOrder(stage.CodeComparisonType_stagedOrder)

	stage.CodeCourseTypeOrder = __gong__recomputeOrder(stage.CodeCourseType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeDirectionTurnTypeOrder = __gong__recomputeOrder(stage.CodeDirectionTurnType_stagedOrder)

	stage.CodeDirectionTypeOrder = __gong__recomputeOrder(stage.CodeDirectionType_stagedOrder)

	stage.CodeEquipmentAntiCollisionTypeOrder = __gong__recomputeOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder)

	stage.CodeFlightOriginTypeOrder = __gong__recomputeOrder(stage.CodeFlightOriginType_stagedOrder)

	stage.CodeFlightPurposeTypeOrder = __gong__recomputeOrder(stage.CodeFlightPurposeType_stagedOrder)

	stage.CodeFlightRestrictionDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeFlightRestrictionDesignatorType_stagedOrder)

	stage.CodeFlightRestrictionTypeOrder = __gong__recomputeOrder(stage.CodeFlightRestrictionType_stagedOrder)

	stage.CodeFlightRuleTypeOrder = __gong__recomputeOrder(stage.CodeFlightRuleType_stagedOrder)

	stage.CodeFlightStatusTypeOrder = __gong__recomputeOrder(stage.CodeFlightStatusType_stagedOrder)

	stage.CodeFlightTypeOrder = __gong__recomputeOrder(stage.CodeFlightType_stagedOrder)

	stage.CodeFlowConditionOperationTypeOrder = __gong__recomputeOrder(stage.CodeFlowConditionOperationType_stagedOrder)

	stage.CodeFreeFlightTypeOrder = __gong__recomputeOrder(stage.CodeFreeFlightType_stagedOrder)

	stage.CodeLevelTypeOrder = __gong__recomputeOrder(stage.CodeLevelType_stagedOrder)

	stage.CodeLocationQualifierTypeOrder = __gong__recomputeOrder(stage.CodeLocationQualifierType_stagedOrder)

	stage.CodeMeteoConditionsTypeOrder = __gong__recomputeOrder(stage.CodeMeteoConditionsType_stagedOrder)

	stage.CodeMilitaryRoutePointTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryRoutePointType_stagedOrder)

	stage.CodeMilitaryStatusTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryStatusType_stagedOrder)

	stage.CodeMilitaryTrainingTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryTrainingType_stagedOrder)

	stage.CodeNavigationEquipmentTypeOrder = __gong__recomputeOrder(stage.CodeNavigationEquipmentType_stagedOrder)

	stage.CodeNavigationSpecificationTypeOrder = __gong__recomputeOrder(stage.CodeNavigationSpecificationType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeObstacleAssessmentSurfaceTypeOrder = __gong__recomputeOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder)

	stage.CodeObstructionIdSurfaceZoneTypeOrder = __gong__recomputeOrder(stage.CodeObstructionIdSurfaceZoneType_stagedOrder)

	stage.CodeRNPTypeOrder = __gong__recomputeOrder(stage.CodeRNPType_stagedOrder)

	stage.CodeRVSMPointRoleTypeOrder = __gong__recomputeOrder(stage.CodeRVSMPointRoleType_stagedOrder)

	stage.CodeRVSMTypeOrder = __gong__recomputeOrder(stage.CodeRVSMType_stagedOrder)

	stage.CodeReferenceRoleTypeOrder = __gong__recomputeOrder(stage.CodeReferenceRoleType_stagedOrder)

	stage.CodeRouteAvailabilityTypeOrder = __gong__recomputeOrder(stage.CodeRouteAvailabilityType_stagedOrder)

	stage.CodeRouteDesignatorLetterTypeOrder = __gong__recomputeOrder(stage.CodeRouteDesignatorLetterType_stagedOrder)

	stage.CodeRouteDesignatorPrefixTypeOrder = __gong__recomputeOrder(stage.CodeRouteDesignatorPrefixType_stagedOrder)

	stage.CodeRouteDesignatorSuffixTypeOrder = __gong__recomputeOrder(stage.CodeRouteDesignatorSuffixType_stagedOrder)

	stage.CodeRouteNavigationTypeOrder = __gong__recomputeOrder(stage.CodeRouteNavigationType_stagedOrder)

	stage.CodeRouteOriginTypeOrder = __gong__recomputeOrder(stage.CodeRouteOriginType_stagedOrder)

	stage.CodeRouteSegmentPathTypeOrder = __gong__recomputeOrder(stage.CodeRouteSegmentPathType_stagedOrder)

	stage.CodeRouteTypeOrder = __gong__recomputeOrder(stage.CodeRouteType_stagedOrder)

	stage.CodeRuleProcedureTitleTypeOrder = __gong__recomputeOrder(stage.CodeRuleProcedureTitleType_stagedOrder)

	stage.CodeRuleProcedureTypeOrder = __gong__recomputeOrder(stage.CodeRuleProcedureType_stagedOrder)

	stage.CodeSideTypeOrder = __gong__recomputeOrder(stage.CodeSideType_stagedOrder)

	stage.CodeSpeedReferenceTypeOrder = __gong__recomputeOrder(stage.CodeSpeedReferenceType_stagedOrder)

	stage.CodeTACANChannelTypeOrder = __gong__recomputeOrder(stage.CodeTACANChannelType_stagedOrder)

	stage.CodeTelecomNetworkTypeOrder = __gong__recomputeOrder(stage.CodeTelecomNetworkType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeTransponderTypeOrder = __gong__recomputeOrder(stage.CodeTransponderType_stagedOrder)

	stage.CodeUpperAlphaTypeOrder = __gong__recomputeOrder(stage.CodeUpperAlphaType_stagedOrder)

	stage.CodeValueInterpretationTypeOrder = __gong__recomputeOrder(stage.CodeValueInterpretationType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeVerticalReferenceTypeOrder = __gong__recomputeOrder(stage.CodeVerticalReferenceType_stagedOrder)

	stage.CodeWakeTurbulenceTypeOrder = __gong__recomputeOrder(stage.CodeWakeTurbulenceType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.NoNumberTypeOrder = __gong__recomputeOrder(stage.NoNumberType_stagedOrder)

	stage.NoSequenceTypeOrder = __gong__recomputeOrder(stage.NoSequenceType_stagedOrder)

	stage.TextAddressTypeOrder = __gong__recomputeOrder(stage.TextAddressType_stagedOrder)

	stage.TextDesignatorTypeOrder = __gong__recomputeOrder(stage.TextDesignatorType_stagedOrder)

	stage.TextInstructionTypeOrder = __gong__recomputeOrder(stage.TextInstructionType_stagedOrder)

	stage.TextNameTypeOrder = __gong__recomputeOrder(stage.TextNameType_stagedOrder)

	stage.TextNoteTypeOrder = __gong__recomputeOrder(stage.TextNoteType_stagedOrder)

	stage.TextPhoneTypeOrder = __gong__recomputeOrder(stage.TextPhoneType_stagedOrder)

	stage.TextPropertyNameTypeOrder = __gong__recomputeOrder(stage.TextPropertyNameType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.ValAngleTypeOrder = __gong__recomputeOrder(stage.ValAngleType_stagedOrder)

	stage.ValBearingTypeOrder = __gong__recomputeOrder(stage.ValBearingType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValSlopeTypeOrder = __gong__recomputeOrder(stage.ValSlopeType_stagedOrder)

	stage.ValSpeedTypeOrder = __gong__recomputeOrder(stage.ValSpeedType_stagedOrder)

	stage.ValWeightTypeOrder = __gong__recomputeOrder(stage.ValWeightType_stagedOrder)

	stage.XHTMLTypeOrder = __gong__recomputeOrder(stage.XHTMLType_stagedOrder)

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
func (*CodeATCReportingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeATCReportingTypes, stage.CodeATCReportingType_stagedOrder)
}

func (*CodeATCReportingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeATCReportingType_orderStaged[order]
}

func (*CodeATCReportingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeATCReportingTypes_mapString
}

func (*CodeATCReportingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeATCReportingTypes
}

func (*CodeATCReportingType) GongNewInstance() any {
	return new(CodeATCReportingType)
}

func (*CodeAerialRefuellingPointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointType_stagedOrder)
}

func (*CodeAerialRefuellingPointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAerialRefuellingPointType_orderStaged[order]
}

func (*CodeAerialRefuellingPointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAerialRefuellingPointTypes_mapString
}

func (*CodeAerialRefuellingPointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAerialRefuellingPointTypes
}

func (*CodeAerialRefuellingPointType) GongNewInstance() any {
	return new(CodeAerialRefuellingPointType)
}

func (*CodeAerialRefuellingPrefixType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixType_stagedOrder)
}

func (*CodeAerialRefuellingPrefixType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAerialRefuellingPrefixType_orderStaged[order]
}

func (*CodeAerialRefuellingPrefixType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAerialRefuellingPrefixTypes_mapString
}

func (*CodeAerialRefuellingPrefixType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAerialRefuellingPrefixTypes
}

func (*CodeAerialRefuellingPrefixType) GongNewInstance() any {
	return new(CodeAerialRefuellingPrefixType)
}

func (*CodeAerialRefuellingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingType_stagedOrder)
}

func (*CodeAerialRefuellingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAerialRefuellingType_orderStaged[order]
}

func (*CodeAerialRefuellingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAerialRefuellingTypes_mapString
}

func (*CodeAerialRefuellingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAerialRefuellingTypes
}

func (*CodeAerialRefuellingType) GongNewInstance() any {
	return new(CodeAerialRefuellingType)
}

func (*CodeAircraftCategoryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryType_stagedOrder)
}

func (*CodeAircraftCategoryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftCategoryType_orderStaged[order]
}

func (*CodeAircraftCategoryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftCategoryTypes_mapString
}

func (*CodeAircraftCategoryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftCategoryTypes
}

func (*CodeAircraftCategoryType) GongNewInstance() any {
	return new(CodeAircraftCategoryType)
}

func (*CodeAircraftEngineNumberType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberType_stagedOrder)
}

func (*CodeAircraftEngineNumberType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftEngineNumberType_orderStaged[order]
}

func (*CodeAircraftEngineNumberType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftEngineNumberTypes_mapString
}

func (*CodeAircraftEngineNumberType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftEngineNumberTypes
}

func (*CodeAircraftEngineNumberType) GongNewInstance() any {
	return new(CodeAircraftEngineNumberType)
}

func (*CodeAircraftEngineType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineType_stagedOrder)
}

func (*CodeAircraftEngineType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftEngineType_orderStaged[order]
}

func (*CodeAircraftEngineType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftEngineTypes_mapString
}

func (*CodeAircraftEngineType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftEngineTypes
}

func (*CodeAircraftEngineType) GongNewInstance() any {
	return new(CodeAircraftEngineType)
}

func (*CodeAircraftICAOType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOType_stagedOrder)
}

func (*CodeAircraftICAOType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftICAOType_orderStaged[order]
}

func (*CodeAircraftICAOType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftICAOTypes_mapString
}

func (*CodeAircraftICAOType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftICAOTypes
}

func (*CodeAircraftICAOType) GongNewInstance() any {
	return new(CodeAircraftICAOType)
}

func (*CodeAircraftType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftTypes, stage.CodeAircraftType_stagedOrder)
}

func (*CodeAircraftType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftType_orderStaged[order]
}

func (*CodeAircraftType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftTypes_mapString
}

func (*CodeAircraftType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftTypes
}

func (*CodeAircraftType) GongNewInstance() any {
	return new(CodeAircraftType)
}

func (*CodeAircraftWingspanClassType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassType_stagedOrder)
}

func (*CodeAircraftWingspanClassType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftWingspanClassType_orderStaged[order]
}

func (*CodeAircraftWingspanClassType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftWingspanClassTypes_mapString
}

func (*CodeAircraftWingspanClassType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftWingspanClassTypes
}

func (*CodeAircraftWingspanClassType) GongNewInstance() any {
	return new(CodeAircraftWingspanClassType)
}

func (*CodeAltitudeAdjustmentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentType_stagedOrder)
}

func (*CodeAltitudeAdjustmentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAltitudeAdjustmentType_orderStaged[order]
}

func (*CodeAltitudeAdjustmentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAltitudeAdjustmentTypes_mapString
}

func (*CodeAltitudeAdjustmentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAltitudeAdjustmentTypes
}

func (*CodeAltitudeAdjustmentType) GongNewInstance() any {
	return new(CodeAltitudeAdjustmentType)
}

func (*CodeAltitudeUseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseType_stagedOrder)
}

func (*CodeAltitudeUseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAltitudeUseType_orderStaged[order]
}

func (*CodeAltitudeUseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAltitudeUseTypes_mapString
}

func (*CodeAltitudeUseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAltitudeUseTypes
}

func (*CodeAltitudeUseType) GongNewInstance() any {
	return new(CodeAltitudeUseType)
}

func (*CodeAuthorityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAuthorityTypes, stage.CodeAuthorityType_stagedOrder)
}

func (*CodeAuthorityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAuthorityType_orderStaged[order]
}

func (*CodeAuthorityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAuthorityTypes_mapString
}

func (*CodeAuthorityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAuthorityTypes
}

func (*CodeAuthorityType) GongNewInstance() any {
	return new(CodeAuthorityType)
}

func (*CodeCardinalDirectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionType_stagedOrder)
}

func (*CodeCardinalDirectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCardinalDirectionType_orderStaged[order]
}

func (*CodeCardinalDirectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCardinalDirectionTypes_mapString
}

func (*CodeCardinalDirectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCardinalDirectionTypes
}

func (*CodeCardinalDirectionType) GongNewInstance() any {
	return new(CodeCardinalDirectionType)
}

func (*CodeCommunicationModeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeType_stagedOrder)
}

func (*CodeCommunicationModeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCommunicationModeType_orderStaged[order]
}

func (*CodeCommunicationModeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCommunicationModeTypes_mapString
}

func (*CodeCommunicationModeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCommunicationModeTypes
}

func (*CodeCommunicationModeType) GongNewInstance() any {
	return new(CodeCommunicationModeType)
}

func (*CodeComparisonType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeComparisonTypes, stage.CodeComparisonType_stagedOrder)
}

func (*CodeComparisonType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeComparisonType_orderStaged[order]
}

func (*CodeComparisonType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeComparisonTypes_mapString
}

func (*CodeComparisonType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeComparisonTypes
}

func (*CodeComparisonType) GongNewInstance() any {
	return new(CodeComparisonType)
}

func (*CodeCourseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCourseTypes, stage.CodeCourseType_stagedOrder)
}

func (*CodeCourseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCourseType_orderStaged[order]
}

func (*CodeCourseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCourseTypes_mapString
}

func (*CodeCourseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCourseTypes
}

func (*CodeCourseType) GongNewInstance() any {
	return new(CodeCourseType)
}

func (*CodeDayType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDayTypes, stage.CodeDayType_stagedOrder)
}

func (*CodeDayType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDayType_orderStaged[order]
}

func (*CodeDayType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDayTypes_mapString
}

func (*CodeDayType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDayTypes
}

func (*CodeDayType) GongNewInstance() any {
	return new(CodeDayType)
}

func (*CodeDirectionTurnType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnType_stagedOrder)
}

func (*CodeDirectionTurnType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDirectionTurnType_orderStaged[order]
}

func (*CodeDirectionTurnType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDirectionTurnTypes_mapString
}

func (*CodeDirectionTurnType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDirectionTurnTypes
}

func (*CodeDirectionTurnType) GongNewInstance() any {
	return new(CodeDirectionTurnType)
}

func (*CodeDirectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDirectionTypes, stage.CodeDirectionType_stagedOrder)
}

func (*CodeDirectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDirectionType_orderStaged[order]
}

func (*CodeDirectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDirectionTypes_mapString
}

func (*CodeDirectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDirectionTypes
}

func (*CodeDirectionType) GongNewInstance() any {
	return new(CodeDirectionType)
}

func (*CodeEquipmentAntiCollisionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionType_stagedOrder)
}

func (*CodeEquipmentAntiCollisionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeEquipmentAntiCollisionType_orderStaged[order]
}

func (*CodeEquipmentAntiCollisionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeEquipmentAntiCollisionTypes_mapString
}

func (*CodeEquipmentAntiCollisionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeEquipmentAntiCollisionTypes
}

func (*CodeEquipmentAntiCollisionType) GongNewInstance() any {
	return new(CodeEquipmentAntiCollisionType)
}

func (*CodeFlightOriginType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightOriginTypes, stage.CodeFlightOriginType_stagedOrder)
}

func (*CodeFlightOriginType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightOriginType_orderStaged[order]
}

func (*CodeFlightOriginType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightOriginTypes_mapString
}

func (*CodeFlightOriginType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightOriginTypes
}

func (*CodeFlightOriginType) GongNewInstance() any {
	return new(CodeFlightOriginType)
}

func (*CodeFlightPurposeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeType_stagedOrder)
}

func (*CodeFlightPurposeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightPurposeType_orderStaged[order]
}

func (*CodeFlightPurposeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightPurposeTypes_mapString
}

func (*CodeFlightPurposeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightPurposeTypes
}

func (*CodeFlightPurposeType) GongNewInstance() any {
	return new(CodeFlightPurposeType)
}

func (*CodeFlightRestrictionDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorType_stagedOrder)
}

func (*CodeFlightRestrictionDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightRestrictionDesignatorType_orderStaged[order]
}

func (*CodeFlightRestrictionDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightRestrictionDesignatorTypes_mapString
}

func (*CodeFlightRestrictionDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightRestrictionDesignatorTypes
}

func (*CodeFlightRestrictionDesignatorType) GongNewInstance() any {
	return new(CodeFlightRestrictionDesignatorType)
}

func (*CodeFlightRestrictionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionType_stagedOrder)
}

func (*CodeFlightRestrictionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightRestrictionType_orderStaged[order]
}

func (*CodeFlightRestrictionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightRestrictionTypes_mapString
}

func (*CodeFlightRestrictionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightRestrictionTypes
}

func (*CodeFlightRestrictionType) GongNewInstance() any {
	return new(CodeFlightRestrictionType)
}

func (*CodeFlightRuleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightRuleTypes, stage.CodeFlightRuleType_stagedOrder)
}

func (*CodeFlightRuleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightRuleType_orderStaged[order]
}

func (*CodeFlightRuleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightRuleTypes_mapString
}

func (*CodeFlightRuleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightRuleTypes
}

func (*CodeFlightRuleType) GongNewInstance() any {
	return new(CodeFlightRuleType)
}

func (*CodeFlightStatusType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightStatusTypes, stage.CodeFlightStatusType_stagedOrder)
}

func (*CodeFlightStatusType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightStatusType_orderStaged[order]
}

func (*CodeFlightStatusType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightStatusTypes_mapString
}

func (*CodeFlightStatusType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightStatusTypes
}

func (*CodeFlightStatusType) GongNewInstance() any {
	return new(CodeFlightStatusType)
}

func (*CodeFlightType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightTypes, stage.CodeFlightType_stagedOrder)
}

func (*CodeFlightType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightType_orderStaged[order]
}

func (*CodeFlightType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightTypes_mapString
}

func (*CodeFlightType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightTypes
}

func (*CodeFlightType) GongNewInstance() any {
	return new(CodeFlightType)
}

func (*CodeFlowConditionOperationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationType_stagedOrder)
}

func (*CodeFlowConditionOperationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlowConditionOperationType_orderStaged[order]
}

func (*CodeFlowConditionOperationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlowConditionOperationTypes_mapString
}

func (*CodeFlowConditionOperationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlowConditionOperationTypes
}

func (*CodeFlowConditionOperationType) GongNewInstance() any {
	return new(CodeFlowConditionOperationType)
}

func (*CodeFreeFlightType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFreeFlightTypes, stage.CodeFreeFlightType_stagedOrder)
}

func (*CodeFreeFlightType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFreeFlightType_orderStaged[order]
}

func (*CodeFreeFlightType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFreeFlightTypes_mapString
}

func (*CodeFreeFlightType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFreeFlightTypes
}

func (*CodeFreeFlightType) GongNewInstance() any {
	return new(CodeFreeFlightType)
}

func (*CodeLevelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLevelTypes, stage.CodeLevelType_stagedOrder)
}

func (*CodeLevelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLevelType_orderStaged[order]
}

func (*CodeLevelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLevelTypes_mapString
}

func (*CodeLevelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLevelTypes
}

func (*CodeLevelType) GongNewInstance() any {
	return new(CodeLevelType)
}

func (*CodeLocationQualifierType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierType_stagedOrder)
}

func (*CodeLocationQualifierType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLocationQualifierType_orderStaged[order]
}

func (*CodeLocationQualifierType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLocationQualifierTypes_mapString
}

func (*CodeLocationQualifierType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLocationQualifierTypes
}

func (*CodeLocationQualifierType) GongNewInstance() any {
	return new(CodeLocationQualifierType)
}

func (*CodeMeteoConditionsType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsType_stagedOrder)
}

func (*CodeMeteoConditionsType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMeteoConditionsType_orderStaged[order]
}

func (*CodeMeteoConditionsType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMeteoConditionsTypes_mapString
}

func (*CodeMeteoConditionsType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMeteoConditionsTypes
}

func (*CodeMeteoConditionsType) GongNewInstance() any {
	return new(CodeMeteoConditionsType)
}

func (*CodeMilitaryRoutePointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointType_stagedOrder)
}

func (*CodeMilitaryRoutePointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMilitaryRoutePointType_orderStaged[order]
}

func (*CodeMilitaryRoutePointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMilitaryRoutePointTypes_mapString
}

func (*CodeMilitaryRoutePointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMilitaryRoutePointTypes
}

func (*CodeMilitaryRoutePointType) GongNewInstance() any {
	return new(CodeMilitaryRoutePointType)
}

func (*CodeMilitaryStatusType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusType_stagedOrder)
}

func (*CodeMilitaryStatusType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMilitaryStatusType_orderStaged[order]
}

func (*CodeMilitaryStatusType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMilitaryStatusTypes_mapString
}

func (*CodeMilitaryStatusType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMilitaryStatusTypes
}

func (*CodeMilitaryStatusType) GongNewInstance() any {
	return new(CodeMilitaryStatusType)
}

func (*CodeMilitaryTrainingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingType_stagedOrder)
}

func (*CodeMilitaryTrainingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMilitaryTrainingType_orderStaged[order]
}

func (*CodeMilitaryTrainingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMilitaryTrainingTypes_mapString
}

func (*CodeMilitaryTrainingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMilitaryTrainingTypes
}

func (*CodeMilitaryTrainingType) GongNewInstance() any {
	return new(CodeMilitaryTrainingType)
}

func (*CodeNavigationEquipmentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentType_stagedOrder)
}

func (*CodeNavigationEquipmentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavigationEquipmentType_orderStaged[order]
}

func (*CodeNavigationEquipmentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavigationEquipmentTypes_mapString
}

func (*CodeNavigationEquipmentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavigationEquipmentTypes
}

func (*CodeNavigationEquipmentType) GongNewInstance() any {
	return new(CodeNavigationEquipmentType)
}

func (*CodeNavigationSpecificationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationType_stagedOrder)
}

func (*CodeNavigationSpecificationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavigationSpecificationType_orderStaged[order]
}

func (*CodeNavigationSpecificationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavigationSpecificationTypes_mapString
}

func (*CodeNavigationSpecificationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavigationSpecificationTypes
}

func (*CodeNavigationSpecificationType) GongNewInstance() any {
	return new(CodeNavigationSpecificationType)
}

func (*CodeNotePurposeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNotePurposeTypes, stage.CodeNotePurposeType_stagedOrder)
}

func (*CodeNotePurposeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNotePurposeType_orderStaged[order]
}

func (*CodeNotePurposeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNotePurposeTypes_mapString
}

func (*CodeNotePurposeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNotePurposeTypes
}

func (*CodeNotePurposeType) GongNewInstance() any {
	return new(CodeNotePurposeType)
}

func (*CodeObstacleAssessmentSurfaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceType_stagedOrder)
}

func (*CodeObstacleAssessmentSurfaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeObstacleAssessmentSurfaceType_orderStaged[order]
}

func (*CodeObstacleAssessmentSurfaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeObstacleAssessmentSurfaceTypes_mapString
}

func (*CodeObstacleAssessmentSurfaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeObstacleAssessmentSurfaceTypes
}

func (*CodeObstacleAssessmentSurfaceType) GongNewInstance() any {
	return new(CodeObstacleAssessmentSurfaceType)
}

func (*CodeObstructionIdSurfaceZoneType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneType_stagedOrder)
}

func (*CodeObstructionIdSurfaceZoneType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeObstructionIdSurfaceZoneType_orderStaged[order]
}

func (*CodeObstructionIdSurfaceZoneType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeObstructionIdSurfaceZoneTypes_mapString
}

func (*CodeObstructionIdSurfaceZoneType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeObstructionIdSurfaceZoneTypes
}

func (*CodeObstructionIdSurfaceZoneType) GongNewInstance() any {
	return new(CodeObstructionIdSurfaceZoneType)
}

func (*CodeRNPType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRNPTypes, stage.CodeRNPType_stagedOrder)
}

func (*CodeRNPType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRNPType_orderStaged[order]
}

func (*CodeRNPType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRNPTypes_mapString
}

func (*CodeRNPType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRNPTypes
}

func (*CodeRNPType) GongNewInstance() any {
	return new(CodeRNPType)
}

func (*CodeRVSMPointRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleType_stagedOrder)
}

func (*CodeRVSMPointRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRVSMPointRoleType_orderStaged[order]
}

func (*CodeRVSMPointRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRVSMPointRoleTypes_mapString
}

func (*CodeRVSMPointRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRVSMPointRoleTypes
}

func (*CodeRVSMPointRoleType) GongNewInstance() any {
	return new(CodeRVSMPointRoleType)
}

func (*CodeRVSMType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRVSMTypes, stage.CodeRVSMType_stagedOrder)
}

func (*CodeRVSMType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRVSMType_orderStaged[order]
}

func (*CodeRVSMType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRVSMTypes_mapString
}

func (*CodeRVSMType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRVSMTypes
}

func (*CodeRVSMType) GongNewInstance() any {
	return new(CodeRVSMType)
}

func (*CodeReferenceRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleType_stagedOrder)
}

func (*CodeReferenceRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeReferenceRoleType_orderStaged[order]
}

func (*CodeReferenceRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeReferenceRoleTypes_mapString
}

func (*CodeReferenceRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeReferenceRoleTypes
}

func (*CodeReferenceRoleType) GongNewInstance() any {
	return new(CodeReferenceRoleType)
}

func (*CodeRouteAvailabilityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityType_stagedOrder)
}

func (*CodeRouteAvailabilityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteAvailabilityType_orderStaged[order]
}

func (*CodeRouteAvailabilityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteAvailabilityTypes_mapString
}

func (*CodeRouteAvailabilityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteAvailabilityTypes
}

func (*CodeRouteAvailabilityType) GongNewInstance() any {
	return new(CodeRouteAvailabilityType)
}

func (*CodeRouteDesignatorLetterType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterType_stagedOrder)
}

func (*CodeRouteDesignatorLetterType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteDesignatorLetterType_orderStaged[order]
}

func (*CodeRouteDesignatorLetterType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteDesignatorLetterTypes_mapString
}

func (*CodeRouteDesignatorLetterType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteDesignatorLetterTypes
}

func (*CodeRouteDesignatorLetterType) GongNewInstance() any {
	return new(CodeRouteDesignatorLetterType)
}

func (*CodeRouteDesignatorPrefixType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixType_stagedOrder)
}

func (*CodeRouteDesignatorPrefixType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteDesignatorPrefixType_orderStaged[order]
}

func (*CodeRouteDesignatorPrefixType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteDesignatorPrefixTypes_mapString
}

func (*CodeRouteDesignatorPrefixType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteDesignatorPrefixTypes
}

func (*CodeRouteDesignatorPrefixType) GongNewInstance() any {
	return new(CodeRouteDesignatorPrefixType)
}

func (*CodeRouteDesignatorSuffixType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixType_stagedOrder)
}

func (*CodeRouteDesignatorSuffixType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteDesignatorSuffixType_orderStaged[order]
}

func (*CodeRouteDesignatorSuffixType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteDesignatorSuffixTypes_mapString
}

func (*CodeRouteDesignatorSuffixType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteDesignatorSuffixTypes
}

func (*CodeRouteDesignatorSuffixType) GongNewInstance() any {
	return new(CodeRouteDesignatorSuffixType)
}

func (*CodeRouteNavigationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationType_stagedOrder)
}

func (*CodeRouteNavigationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteNavigationType_orderStaged[order]
}

func (*CodeRouteNavigationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteNavigationTypes_mapString
}

func (*CodeRouteNavigationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteNavigationTypes
}

func (*CodeRouteNavigationType) GongNewInstance() any {
	return new(CodeRouteNavigationType)
}

func (*CodeRouteOriginType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteOriginTypes, stage.CodeRouteOriginType_stagedOrder)
}

func (*CodeRouteOriginType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteOriginType_orderStaged[order]
}

func (*CodeRouteOriginType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteOriginTypes_mapString
}

func (*CodeRouteOriginType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteOriginTypes
}

func (*CodeRouteOriginType) GongNewInstance() any {
	return new(CodeRouteOriginType)
}

func (*CodeRouteSegmentPathType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathType_stagedOrder)
}

func (*CodeRouteSegmentPathType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteSegmentPathType_orderStaged[order]
}

func (*CodeRouteSegmentPathType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteSegmentPathTypes_mapString
}

func (*CodeRouteSegmentPathType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteSegmentPathTypes
}

func (*CodeRouteSegmentPathType) GongNewInstance() any {
	return new(CodeRouteSegmentPathType)
}

func (*CodeRouteType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteTypes, stage.CodeRouteType_stagedOrder)
}

func (*CodeRouteType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteType_orderStaged[order]
}

func (*CodeRouteType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteTypes_mapString
}

func (*CodeRouteType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteTypes
}

func (*CodeRouteType) GongNewInstance() any {
	return new(CodeRouteType)
}

func (*CodeRuleProcedureTitleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleType_stagedOrder)
}

func (*CodeRuleProcedureTitleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRuleProcedureTitleType_orderStaged[order]
}

func (*CodeRuleProcedureTitleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRuleProcedureTitleTypes_mapString
}

func (*CodeRuleProcedureTitleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRuleProcedureTitleTypes
}

func (*CodeRuleProcedureTitleType) GongNewInstance() any {
	return new(CodeRuleProcedureTitleType)
}

func (*CodeRuleProcedureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureType_stagedOrder)
}

func (*CodeRuleProcedureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRuleProcedureType_orderStaged[order]
}

func (*CodeRuleProcedureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRuleProcedureTypes_mapString
}

func (*CodeRuleProcedureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRuleProcedureTypes
}

func (*CodeRuleProcedureType) GongNewInstance() any {
	return new(CodeRuleProcedureType)
}

func (*CodeSideType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSideTypes, stage.CodeSideType_stagedOrder)
}

func (*CodeSideType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSideType_orderStaged[order]
}

func (*CodeSideType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSideTypes_mapString
}

func (*CodeSideType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSideTypes
}

func (*CodeSideType) GongNewInstance() any {
	return new(CodeSideType)
}

func (*CodeSpeedReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceType_stagedOrder)
}

func (*CodeSpeedReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSpeedReferenceType_orderStaged[order]
}

func (*CodeSpeedReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSpeedReferenceTypes_mapString
}

func (*CodeSpeedReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSpeedReferenceTypes
}

func (*CodeSpeedReferenceType) GongNewInstance() any {
	return new(CodeSpeedReferenceType)
}

func (*CodeTACANChannelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTACANChannelTypes, stage.CodeTACANChannelType_stagedOrder)
}

func (*CodeTACANChannelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTACANChannelType_orderStaged[order]
}

func (*CodeTACANChannelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTACANChannelTypes_mapString
}

func (*CodeTACANChannelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTACANChannelTypes
}

func (*CodeTACANChannelType) GongNewInstance() any {
	return new(CodeTACANChannelType)
}

func (*CodeTelecomNetworkType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkType_stagedOrder)
}

func (*CodeTelecomNetworkType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTelecomNetworkType_orderStaged[order]
}

func (*CodeTelecomNetworkType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTelecomNetworkTypes_mapString
}

func (*CodeTelecomNetworkType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTelecomNetworkTypes
}

func (*CodeTelecomNetworkType) GongNewInstance() any {
	return new(CodeTelecomNetworkType)
}

func (*CodeTimeEventCombinationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationType_stagedOrder)
}

func (*CodeTimeEventCombinationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTimeEventCombinationType_orderStaged[order]
}

func (*CodeTimeEventCombinationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTimeEventCombinationTypes_mapString
}

func (*CodeTimeEventCombinationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTimeEventCombinationTypes
}

func (*CodeTimeEventCombinationType) GongNewInstance() any {
	return new(CodeTimeEventCombinationType)
}

func (*CodeTimeEventType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTimeEventTypes, stage.CodeTimeEventType_stagedOrder)
}

func (*CodeTimeEventType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTimeEventType_orderStaged[order]
}

func (*CodeTimeEventType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTimeEventTypes_mapString
}

func (*CodeTimeEventType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTimeEventTypes
}

func (*CodeTimeEventType) GongNewInstance() any {
	return new(CodeTimeEventType)
}

func (*CodeTimeReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceType_stagedOrder)
}

func (*CodeTimeReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTimeReferenceType_orderStaged[order]
}

func (*CodeTimeReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTimeReferenceTypes_mapString
}

func (*CodeTimeReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTimeReferenceTypes
}

func (*CodeTimeReferenceType) GongNewInstance() any {
	return new(CodeTimeReferenceType)
}

func (*CodeTransponderType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTransponderTypes, stage.CodeTransponderType_stagedOrder)
}

func (*CodeTransponderType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTransponderType_orderStaged[order]
}

func (*CodeTransponderType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTransponderTypes_mapString
}

func (*CodeTransponderType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTransponderTypes
}

func (*CodeTransponderType) GongNewInstance() any {
	return new(CodeTransponderType)
}

func (*CodeUpperAlphaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaType_stagedOrder)
}

func (*CodeUpperAlphaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeUpperAlphaType_orderStaged[order]
}

func (*CodeUpperAlphaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeUpperAlphaTypes_mapString
}

func (*CodeUpperAlphaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeUpperAlphaTypes
}

func (*CodeUpperAlphaType) GongNewInstance() any {
	return new(CodeUpperAlphaType)
}

func (*CodeValueInterpretationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationType_stagedOrder)
}

func (*CodeValueInterpretationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeValueInterpretationType_orderStaged[order]
}

func (*CodeValueInterpretationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeValueInterpretationTypes_mapString
}

func (*CodeValueInterpretationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeValueInterpretationTypes
}

func (*CodeValueInterpretationType) GongNewInstance() any {
	return new(CodeValueInterpretationType)
}

func (*CodeVerticalDatumType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumType_stagedOrder)
}

func (*CodeVerticalDatumType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVerticalDatumType_orderStaged[order]
}

func (*CodeVerticalDatumType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVerticalDatumTypes_mapString
}

func (*CodeVerticalDatumType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVerticalDatumTypes
}

func (*CodeVerticalDatumType) GongNewInstance() any {
	return new(CodeVerticalDatumType)
}

func (*CodeVerticalReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceType_stagedOrder)
}

func (*CodeVerticalReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVerticalReferenceType_orderStaged[order]
}

func (*CodeVerticalReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVerticalReferenceTypes_mapString
}

func (*CodeVerticalReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVerticalReferenceTypes
}

func (*CodeVerticalReferenceType) GongNewInstance() any {
	return new(CodeVerticalReferenceType)
}

func (*CodeWakeTurbulenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceType_stagedOrder)
}

func (*CodeWakeTurbulenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeWakeTurbulenceType_orderStaged[order]
}

func (*CodeWakeTurbulenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeWakeTurbulenceTypes_mapString
}

func (*CodeWakeTurbulenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeWakeTurbulenceTypes
}

func (*CodeWakeTurbulenceType) GongNewInstance() any {
	return new(CodeWakeTurbulenceType)
}

func (*CodeYesNoType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeYesNoTypes, stage.CodeYesNoType_stagedOrder)
}

func (*CodeYesNoType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeYesNoType_orderStaged[order]
}

func (*CodeYesNoType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeYesNoTypes_mapString
}

func (*CodeYesNoType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeYesNoTypes
}

func (*CodeYesNoType) GongNewInstance() any {
	return new(CodeYesNoType)
}

func (*DateMonthDayType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DateMonthDayTypes, stage.DateMonthDayType_stagedOrder)
}

func (*DateMonthDayType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DateMonthDayType_orderStaged[order]
}

func (*DateMonthDayType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DateMonthDayTypes_mapString
}

func (*DateMonthDayType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DateMonthDayTypes
}

func (*DateMonthDayType) GongNewInstance() any {
	return new(DateMonthDayType)
}

func (*NoNumberType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoNumberTypes, stage.NoNumberType_stagedOrder)
}

func (*NoNumberType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoNumberType_orderStaged[order]
}

func (*NoNumberType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoNumberTypes_mapString
}

func (*NoNumberType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoNumberTypes
}

func (*NoNumberType) GongNewInstance() any {
	return new(NoNumberType)
}

func (*NoSequenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoSequenceTypes, stage.NoSequenceType_stagedOrder)
}

func (*NoSequenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoSequenceType_orderStaged[order]
}

func (*NoSequenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoSequenceTypes_mapString
}

func (*NoSequenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoSequenceTypes
}

func (*NoSequenceType) GongNewInstance() any {
	return new(NoSequenceType)
}

func (*TextAddressType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextAddressTypes, stage.TextAddressType_stagedOrder)
}

func (*TextAddressType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextAddressType_orderStaged[order]
}

func (*TextAddressType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextAddressTypes_mapString
}

func (*TextAddressType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextAddressTypes
}

func (*TextAddressType) GongNewInstance() any {
	return new(TextAddressType)
}

func (*TextDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextDesignatorTypes, stage.TextDesignatorType_stagedOrder)
}

func (*TextDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextDesignatorType_orderStaged[order]
}

func (*TextDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextDesignatorTypes_mapString
}

func (*TextDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextDesignatorTypes
}

func (*TextDesignatorType) GongNewInstance() any {
	return new(TextDesignatorType)
}

func (*TextInstructionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextInstructionTypes, stage.TextInstructionType_stagedOrder)
}

func (*TextInstructionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextInstructionType_orderStaged[order]
}

func (*TextInstructionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextInstructionTypes_mapString
}

func (*TextInstructionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextInstructionTypes
}

func (*TextInstructionType) GongNewInstance() any {
	return new(TextInstructionType)
}

func (*TextNameType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextNameTypes, stage.TextNameType_stagedOrder)
}

func (*TextNameType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextNameType_orderStaged[order]
}

func (*TextNameType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextNameTypes_mapString
}

func (*TextNameType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextNameTypes
}

func (*TextNameType) GongNewInstance() any {
	return new(TextNameType)
}

func (*TextNoteType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextNoteTypes, stage.TextNoteType_stagedOrder)
}

func (*TextNoteType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextNoteType_orderStaged[order]
}

func (*TextNoteType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextNoteTypes_mapString
}

func (*TextNoteType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextNoteTypes
}

func (*TextNoteType) GongNewInstance() any {
	return new(TextNoteType)
}

func (*TextPhoneType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextPhoneTypes, stage.TextPhoneType_stagedOrder)
}

func (*TextPhoneType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextPhoneType_orderStaged[order]
}

func (*TextPhoneType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextPhoneTypes_mapString
}

func (*TextPhoneType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextPhoneTypes
}

func (*TextPhoneType) GongNewInstance() any {
	return new(TextPhoneType)
}

func (*TextPropertyNameType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextPropertyNameTypes, stage.TextPropertyNameType_stagedOrder)
}

func (*TextPropertyNameType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextPropertyNameType_orderStaged[order]
}

func (*TextPropertyNameType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextPropertyNameTypes_mapString
}

func (*TextPropertyNameType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextPropertyNameTypes
}

func (*TextPropertyNameType) GongNewInstance() any {
	return new(TextPropertyNameType)
}

func (*TimeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimeTypes, stage.TimeType_stagedOrder)
}

func (*TimeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimeType_orderStaged[order]
}

func (*TimeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimeTypes_mapString
}

func (*TimeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimeTypes
}

func (*TimeType) GongNewInstance() any {
	return new(TimeType)
}

func (*ValAngleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValAngleTypes, stage.ValAngleType_stagedOrder)
}

func (*ValAngleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValAngleType_orderStaged[order]
}

func (*ValAngleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValAngleTypes_mapString
}

func (*ValAngleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValAngleTypes
}

func (*ValAngleType) GongNewInstance() any {
	return new(ValAngleType)
}

func (*ValBearingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValBearingTypes, stage.ValBearingType_stagedOrder)
}

func (*ValBearingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValBearingType_orderStaged[order]
}

func (*ValBearingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValBearingTypes_mapString
}

func (*ValBearingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValBearingTypes
}

func (*ValBearingType) GongNewInstance() any {
	return new(ValBearingType)
}

func (*ValDistanceSignedType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValDistanceSignedTypes, stage.ValDistanceSignedType_stagedOrder)
}

func (*ValDistanceSignedType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValDistanceSignedType_orderStaged[order]
}

func (*ValDistanceSignedType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValDistanceSignedTypes_mapString
}

func (*ValDistanceSignedType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValDistanceSignedTypes
}

func (*ValDistanceSignedType) GongNewInstance() any {
	return new(ValDistanceSignedType)
}

func (*ValDistanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValDistanceTypes, stage.ValDistanceType_stagedOrder)
}

func (*ValDistanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValDistanceType_orderStaged[order]
}

func (*ValDistanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValDistanceTypes_mapString
}

func (*ValDistanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValDistanceTypes
}

func (*ValDistanceType) GongNewInstance() any {
	return new(ValDistanceType)
}

func (*ValDistanceVerticalType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalType_stagedOrder)
}

func (*ValDistanceVerticalType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValDistanceVerticalType_orderStaged[order]
}

func (*ValDistanceVerticalType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValDistanceVerticalTypes_mapString
}

func (*ValDistanceVerticalType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValDistanceVerticalTypes
}

func (*ValDistanceVerticalType) GongNewInstance() any {
	return new(ValDistanceVerticalType)
}

func (*ValDurationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValDurationTypes, stage.ValDurationType_stagedOrder)
}

func (*ValDurationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValDurationType_orderStaged[order]
}

func (*ValDurationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValDurationTypes_mapString
}

func (*ValDurationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValDurationTypes
}

func (*ValDurationType) GongNewInstance() any {
	return new(ValDurationType)
}

func (*ValSlopeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValSlopeTypes, stage.ValSlopeType_stagedOrder)
}

func (*ValSlopeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValSlopeType_orderStaged[order]
}

func (*ValSlopeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValSlopeTypes_mapString
}

func (*ValSlopeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValSlopeTypes
}

func (*ValSlopeType) GongNewInstance() any {
	return new(ValSlopeType)
}

func (*ValSpeedType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValSpeedTypes, stage.ValSpeedType_stagedOrder)
}

func (*ValSpeedType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValSpeedType_orderStaged[order]
}

func (*ValSpeedType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValSpeedTypes_mapString
}

func (*ValSpeedType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValSpeedTypes
}

func (*ValSpeedType) GongNewInstance() any {
	return new(ValSpeedType)
}

func (*ValWeightType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValWeightTypes, stage.ValWeightType_stagedOrder)
}

func (*ValWeightType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValWeightType_orderStaged[order]
}

func (*ValWeightType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValWeightTypes_mapString
}

func (*ValWeightType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValWeightTypes
}

func (*ValWeightType) GongNewInstance() any {
	return new(ValWeightType)
}

func (*XHTMLType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.XHTMLTypes, stage.XHTMLType_stagedOrder)
}

func (*XHTMLType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.XHTMLType_orderStaged[order]
}

func (*XHTMLType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.XHTMLTypes_mapString
}

func (*XHTMLType) GongGetInstancesSet(stage *Stage) any {
	return &stage.XHTMLTypes
}

func (*XHTMLType) GongNewInstance() any {
	return new(XHTMLType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
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
		CodeATCReportingTypes:           make(map[*CodeATCReportingType]struct{}),
		CodeATCReportingTypes_mapString: make(map[string]*CodeATCReportingType),

		CodeAerialRefuellingPointTypes:           make(map[*CodeAerialRefuellingPointType]struct{}),
		CodeAerialRefuellingPointTypes_mapString: make(map[string]*CodeAerialRefuellingPointType),

		CodeAerialRefuellingPrefixTypes:           make(map[*CodeAerialRefuellingPrefixType]struct{}),
		CodeAerialRefuellingPrefixTypes_mapString: make(map[string]*CodeAerialRefuellingPrefixType),

		CodeAerialRefuellingTypes:           make(map[*CodeAerialRefuellingType]struct{}),
		CodeAerialRefuellingTypes_mapString: make(map[string]*CodeAerialRefuellingType),

		CodeAircraftCategoryTypes:           make(map[*CodeAircraftCategoryType]struct{}),
		CodeAircraftCategoryTypes_mapString: make(map[string]*CodeAircraftCategoryType),

		CodeAircraftEngineNumberTypes:           make(map[*CodeAircraftEngineNumberType]struct{}),
		CodeAircraftEngineNumberTypes_mapString: make(map[string]*CodeAircraftEngineNumberType),

		CodeAircraftEngineTypes:           make(map[*CodeAircraftEngineType]struct{}),
		CodeAircraftEngineTypes_mapString: make(map[string]*CodeAircraftEngineType),

		CodeAircraftICAOTypes:           make(map[*CodeAircraftICAOType]struct{}),
		CodeAircraftICAOTypes_mapString: make(map[string]*CodeAircraftICAOType),

		CodeAircraftTypes:           make(map[*CodeAircraftType]struct{}),
		CodeAircraftTypes_mapString: make(map[string]*CodeAircraftType),

		CodeAircraftWingspanClassTypes:           make(map[*CodeAircraftWingspanClassType]struct{}),
		CodeAircraftWingspanClassTypes_mapString: make(map[string]*CodeAircraftWingspanClassType),

		CodeAltitudeAdjustmentTypes:           make(map[*CodeAltitudeAdjustmentType]struct{}),
		CodeAltitudeAdjustmentTypes_mapString: make(map[string]*CodeAltitudeAdjustmentType),

		CodeAltitudeUseTypes:           make(map[*CodeAltitudeUseType]struct{}),
		CodeAltitudeUseTypes_mapString: make(map[string]*CodeAltitudeUseType),

		CodeAuthorityTypes:           make(map[*CodeAuthorityType]struct{}),
		CodeAuthorityTypes_mapString: make(map[string]*CodeAuthorityType),

		CodeCardinalDirectionTypes:           make(map[*CodeCardinalDirectionType]struct{}),
		CodeCardinalDirectionTypes_mapString: make(map[string]*CodeCardinalDirectionType),

		CodeCommunicationModeTypes:           make(map[*CodeCommunicationModeType]struct{}),
		CodeCommunicationModeTypes_mapString: make(map[string]*CodeCommunicationModeType),

		CodeComparisonTypes:           make(map[*CodeComparisonType]struct{}),
		CodeComparisonTypes_mapString: make(map[string]*CodeComparisonType),

		CodeCourseTypes:           make(map[*CodeCourseType]struct{}),
		CodeCourseTypes_mapString: make(map[string]*CodeCourseType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeDirectionTurnTypes:           make(map[*CodeDirectionTurnType]struct{}),
		CodeDirectionTurnTypes_mapString: make(map[string]*CodeDirectionTurnType),

		CodeDirectionTypes:           make(map[*CodeDirectionType]struct{}),
		CodeDirectionTypes_mapString: make(map[string]*CodeDirectionType),

		CodeEquipmentAntiCollisionTypes:           make(map[*CodeEquipmentAntiCollisionType]struct{}),
		CodeEquipmentAntiCollisionTypes_mapString: make(map[string]*CodeEquipmentAntiCollisionType),

		CodeFlightOriginTypes:           make(map[*CodeFlightOriginType]struct{}),
		CodeFlightOriginTypes_mapString: make(map[string]*CodeFlightOriginType),

		CodeFlightPurposeTypes:           make(map[*CodeFlightPurposeType]struct{}),
		CodeFlightPurposeTypes_mapString: make(map[string]*CodeFlightPurposeType),

		CodeFlightRestrictionDesignatorTypes:           make(map[*CodeFlightRestrictionDesignatorType]struct{}),
		CodeFlightRestrictionDesignatorTypes_mapString: make(map[string]*CodeFlightRestrictionDesignatorType),

		CodeFlightRestrictionTypes:           make(map[*CodeFlightRestrictionType]struct{}),
		CodeFlightRestrictionTypes_mapString: make(map[string]*CodeFlightRestrictionType),

		CodeFlightRuleTypes:           make(map[*CodeFlightRuleType]struct{}),
		CodeFlightRuleTypes_mapString: make(map[string]*CodeFlightRuleType),

		CodeFlightStatusTypes:           make(map[*CodeFlightStatusType]struct{}),
		CodeFlightStatusTypes_mapString: make(map[string]*CodeFlightStatusType),

		CodeFlightTypes:           make(map[*CodeFlightType]struct{}),
		CodeFlightTypes_mapString: make(map[string]*CodeFlightType),

		CodeFlowConditionOperationTypes:           make(map[*CodeFlowConditionOperationType]struct{}),
		CodeFlowConditionOperationTypes_mapString: make(map[string]*CodeFlowConditionOperationType),

		CodeFreeFlightTypes:           make(map[*CodeFreeFlightType]struct{}),
		CodeFreeFlightTypes_mapString: make(map[string]*CodeFreeFlightType),

		CodeLevelTypes:           make(map[*CodeLevelType]struct{}),
		CodeLevelTypes_mapString: make(map[string]*CodeLevelType),

		CodeLocationQualifierTypes:           make(map[*CodeLocationQualifierType]struct{}),
		CodeLocationQualifierTypes_mapString: make(map[string]*CodeLocationQualifierType),

		CodeMeteoConditionsTypes:           make(map[*CodeMeteoConditionsType]struct{}),
		CodeMeteoConditionsTypes_mapString: make(map[string]*CodeMeteoConditionsType),

		CodeMilitaryRoutePointTypes:           make(map[*CodeMilitaryRoutePointType]struct{}),
		CodeMilitaryRoutePointTypes_mapString: make(map[string]*CodeMilitaryRoutePointType),

		CodeMilitaryStatusTypes:           make(map[*CodeMilitaryStatusType]struct{}),
		CodeMilitaryStatusTypes_mapString: make(map[string]*CodeMilitaryStatusType),

		CodeMilitaryTrainingTypes:           make(map[*CodeMilitaryTrainingType]struct{}),
		CodeMilitaryTrainingTypes_mapString: make(map[string]*CodeMilitaryTrainingType),

		CodeNavigationEquipmentTypes:           make(map[*CodeNavigationEquipmentType]struct{}),
		CodeNavigationEquipmentTypes_mapString: make(map[string]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationTypes:           make(map[*CodeNavigationSpecificationType]struct{}),
		CodeNavigationSpecificationTypes_mapString: make(map[string]*CodeNavigationSpecificationType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodeObstacleAssessmentSurfaceTypes:           make(map[*CodeObstacleAssessmentSurfaceType]struct{}),
		CodeObstacleAssessmentSurfaceTypes_mapString: make(map[string]*CodeObstacleAssessmentSurfaceType),

		CodeObstructionIdSurfaceZoneTypes:           make(map[*CodeObstructionIdSurfaceZoneType]struct{}),
		CodeObstructionIdSurfaceZoneTypes_mapString: make(map[string]*CodeObstructionIdSurfaceZoneType),

		CodeRNPTypes:           make(map[*CodeRNPType]struct{}),
		CodeRNPTypes_mapString: make(map[string]*CodeRNPType),

		CodeRVSMPointRoleTypes:           make(map[*CodeRVSMPointRoleType]struct{}),
		CodeRVSMPointRoleTypes_mapString: make(map[string]*CodeRVSMPointRoleType),

		CodeRVSMTypes:           make(map[*CodeRVSMType]struct{}),
		CodeRVSMTypes_mapString: make(map[string]*CodeRVSMType),

		CodeReferenceRoleTypes:           make(map[*CodeReferenceRoleType]struct{}),
		CodeReferenceRoleTypes_mapString: make(map[string]*CodeReferenceRoleType),

		CodeRouteAvailabilityTypes:           make(map[*CodeRouteAvailabilityType]struct{}),
		CodeRouteAvailabilityTypes_mapString: make(map[string]*CodeRouteAvailabilityType),

		CodeRouteDesignatorLetterTypes:           make(map[*CodeRouteDesignatorLetterType]struct{}),
		CodeRouteDesignatorLetterTypes_mapString: make(map[string]*CodeRouteDesignatorLetterType),

		CodeRouteDesignatorPrefixTypes:           make(map[*CodeRouteDesignatorPrefixType]struct{}),
		CodeRouteDesignatorPrefixTypes_mapString: make(map[string]*CodeRouteDesignatorPrefixType),

		CodeRouteDesignatorSuffixTypes:           make(map[*CodeRouteDesignatorSuffixType]struct{}),
		CodeRouteDesignatorSuffixTypes_mapString: make(map[string]*CodeRouteDesignatorSuffixType),

		CodeRouteNavigationTypes:           make(map[*CodeRouteNavigationType]struct{}),
		CodeRouteNavigationTypes_mapString: make(map[string]*CodeRouteNavigationType),

		CodeRouteOriginTypes:           make(map[*CodeRouteOriginType]struct{}),
		CodeRouteOriginTypes_mapString: make(map[string]*CodeRouteOriginType),

		CodeRouteSegmentPathTypes:           make(map[*CodeRouteSegmentPathType]struct{}),
		CodeRouteSegmentPathTypes_mapString: make(map[string]*CodeRouteSegmentPathType),

		CodeRouteTypes:           make(map[*CodeRouteType]struct{}),
		CodeRouteTypes_mapString: make(map[string]*CodeRouteType),

		CodeRuleProcedureTitleTypes:           make(map[*CodeRuleProcedureTitleType]struct{}),
		CodeRuleProcedureTitleTypes_mapString: make(map[string]*CodeRuleProcedureTitleType),

		CodeRuleProcedureTypes:           make(map[*CodeRuleProcedureType]struct{}),
		CodeRuleProcedureTypes_mapString: make(map[string]*CodeRuleProcedureType),

		CodeSideTypes:           make(map[*CodeSideType]struct{}),
		CodeSideTypes_mapString: make(map[string]*CodeSideType),

		CodeSpeedReferenceTypes:           make(map[*CodeSpeedReferenceType]struct{}),
		CodeSpeedReferenceTypes_mapString: make(map[string]*CodeSpeedReferenceType),

		CodeTACANChannelTypes:           make(map[*CodeTACANChannelType]struct{}),
		CodeTACANChannelTypes_mapString: make(map[string]*CodeTACANChannelType),

		CodeTelecomNetworkTypes:           make(map[*CodeTelecomNetworkType]struct{}),
		CodeTelecomNetworkTypes_mapString: make(map[string]*CodeTelecomNetworkType),

		CodeTimeEventCombinationTypes:           make(map[*CodeTimeEventCombinationType]struct{}),
		CodeTimeEventCombinationTypes_mapString: make(map[string]*CodeTimeEventCombinationType),

		CodeTimeEventTypes:           make(map[*CodeTimeEventType]struct{}),
		CodeTimeEventTypes_mapString: make(map[string]*CodeTimeEventType),

		CodeTimeReferenceTypes:           make(map[*CodeTimeReferenceType]struct{}),
		CodeTimeReferenceTypes_mapString: make(map[string]*CodeTimeReferenceType),

		CodeTransponderTypes:           make(map[*CodeTransponderType]struct{}),
		CodeTransponderTypes_mapString: make(map[string]*CodeTransponderType),

		CodeUpperAlphaTypes:           make(map[*CodeUpperAlphaType]struct{}),
		CodeUpperAlphaTypes_mapString: make(map[string]*CodeUpperAlphaType),

		CodeValueInterpretationTypes:           make(map[*CodeValueInterpretationType]struct{}),
		CodeValueInterpretationTypes_mapString: make(map[string]*CodeValueInterpretationType),

		CodeVerticalDatumTypes:           make(map[*CodeVerticalDatumType]struct{}),
		CodeVerticalDatumTypes_mapString: make(map[string]*CodeVerticalDatumType),

		CodeVerticalReferenceTypes:           make(map[*CodeVerticalReferenceType]struct{}),
		CodeVerticalReferenceTypes_mapString: make(map[string]*CodeVerticalReferenceType),

		CodeWakeTurbulenceTypes:           make(map[*CodeWakeTurbulenceType]struct{}),
		CodeWakeTurbulenceTypes_mapString: make(map[string]*CodeWakeTurbulenceType),

		CodeYesNoTypes:           make(map[*CodeYesNoType]struct{}),
		CodeYesNoTypes_mapString: make(map[string]*CodeYesNoType),

		DateMonthDayTypes:           make(map[*DateMonthDayType]struct{}),
		DateMonthDayTypes_mapString: make(map[string]*DateMonthDayType),

		NoNumberTypes:           make(map[*NoNumberType]struct{}),
		NoNumberTypes_mapString: make(map[string]*NoNumberType),

		NoSequenceTypes:           make(map[*NoSequenceType]struct{}),
		NoSequenceTypes_mapString: make(map[string]*NoSequenceType),

		TextAddressTypes:           make(map[*TextAddressType]struct{}),
		TextAddressTypes_mapString: make(map[string]*TextAddressType),

		TextDesignatorTypes:           make(map[*TextDesignatorType]struct{}),
		TextDesignatorTypes_mapString: make(map[string]*TextDesignatorType),

		TextInstructionTypes:           make(map[*TextInstructionType]struct{}),
		TextInstructionTypes_mapString: make(map[string]*TextInstructionType),

		TextNameTypes:           make(map[*TextNameType]struct{}),
		TextNameTypes_mapString: make(map[string]*TextNameType),

		TextNoteTypes:           make(map[*TextNoteType]struct{}),
		TextNoteTypes_mapString: make(map[string]*TextNoteType),

		TextPhoneTypes:           make(map[*TextPhoneType]struct{}),
		TextPhoneTypes_mapString: make(map[string]*TextPhoneType),

		TextPropertyNameTypes:           make(map[*TextPropertyNameType]struct{}),
		TextPropertyNameTypes_mapString: make(map[string]*TextPropertyNameType),

		TimeTypes:           make(map[*TimeType]struct{}),
		TimeTypes_mapString: make(map[string]*TimeType),

		ValAngleTypes:           make(map[*ValAngleType]struct{}),
		ValAngleTypes_mapString: make(map[string]*ValAngleType),

		ValBearingTypes:           make(map[*ValBearingType]struct{}),
		ValBearingTypes_mapString: make(map[string]*ValBearingType),

		ValDistanceSignedTypes:           make(map[*ValDistanceSignedType]struct{}),
		ValDistanceSignedTypes_mapString: make(map[string]*ValDistanceSignedType),

		ValDistanceTypes:           make(map[*ValDistanceType]struct{}),
		ValDistanceTypes_mapString: make(map[string]*ValDistanceType),

		ValDistanceVerticalTypes:           make(map[*ValDistanceVerticalType]struct{}),
		ValDistanceVerticalTypes_mapString: make(map[string]*ValDistanceVerticalType),

		ValDurationTypes:           make(map[*ValDurationType]struct{}),
		ValDurationTypes_mapString: make(map[string]*ValDurationType),

		ValSlopeTypes:           make(map[*ValSlopeType]struct{}),
		ValSlopeTypes_mapString: make(map[string]*ValSlopeType),

		ValSpeedTypes:           make(map[*ValSpeedType]struct{}),
		ValSpeedTypes_mapString: make(map[string]*ValSpeedType),

		ValWeightTypes:           make(map[*ValWeightType]struct{}),
		ValWeightTypes_mapString: make(map[string]*ValWeightType),

		XHTMLTypes:           make(map[*XHTMLType]struct{}),
		XHTMLTypes_mapString: make(map[string]*XHTMLType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CodeATCReportingType_stagedOrder: make(map[*CodeATCReportingType]uint),
		CodeATCReportingType_orderStaged: make(map[uint]*CodeATCReportingType),
		CodeATCReportingTypes_reference:  make(map[*CodeATCReportingType]*CodeATCReportingType),

		CodeAerialRefuellingPointType_stagedOrder: make(map[*CodeAerialRefuellingPointType]uint),
		CodeAerialRefuellingPointType_orderStaged: make(map[uint]*CodeAerialRefuellingPointType),
		CodeAerialRefuellingPointTypes_reference:  make(map[*CodeAerialRefuellingPointType]*CodeAerialRefuellingPointType),

		CodeAerialRefuellingPrefixType_stagedOrder: make(map[*CodeAerialRefuellingPrefixType]uint),
		CodeAerialRefuellingPrefixType_orderStaged: make(map[uint]*CodeAerialRefuellingPrefixType),
		CodeAerialRefuellingPrefixTypes_reference:  make(map[*CodeAerialRefuellingPrefixType]*CodeAerialRefuellingPrefixType),

		CodeAerialRefuellingType_stagedOrder: make(map[*CodeAerialRefuellingType]uint),
		CodeAerialRefuellingType_orderStaged: make(map[uint]*CodeAerialRefuellingType),
		CodeAerialRefuellingTypes_reference:  make(map[*CodeAerialRefuellingType]*CodeAerialRefuellingType),

		CodeAircraftCategoryType_stagedOrder: make(map[*CodeAircraftCategoryType]uint),
		CodeAircraftCategoryType_orderStaged: make(map[uint]*CodeAircraftCategoryType),
		CodeAircraftCategoryTypes_reference:  make(map[*CodeAircraftCategoryType]*CodeAircraftCategoryType),

		CodeAircraftEngineNumberType_stagedOrder: make(map[*CodeAircraftEngineNumberType]uint),
		CodeAircraftEngineNumberType_orderStaged: make(map[uint]*CodeAircraftEngineNumberType),
		CodeAircraftEngineNumberTypes_reference:  make(map[*CodeAircraftEngineNumberType]*CodeAircraftEngineNumberType),

		CodeAircraftEngineType_stagedOrder: make(map[*CodeAircraftEngineType]uint),
		CodeAircraftEngineType_orderStaged: make(map[uint]*CodeAircraftEngineType),
		CodeAircraftEngineTypes_reference:  make(map[*CodeAircraftEngineType]*CodeAircraftEngineType),

		CodeAircraftICAOType_stagedOrder: make(map[*CodeAircraftICAOType]uint),
		CodeAircraftICAOType_orderStaged: make(map[uint]*CodeAircraftICAOType),
		CodeAircraftICAOTypes_reference:  make(map[*CodeAircraftICAOType]*CodeAircraftICAOType),

		CodeAircraftType_stagedOrder: make(map[*CodeAircraftType]uint),
		CodeAircraftType_orderStaged: make(map[uint]*CodeAircraftType),
		CodeAircraftTypes_reference:  make(map[*CodeAircraftType]*CodeAircraftType),

		CodeAircraftWingspanClassType_stagedOrder: make(map[*CodeAircraftWingspanClassType]uint),
		CodeAircraftWingspanClassType_orderStaged: make(map[uint]*CodeAircraftWingspanClassType),
		CodeAircraftWingspanClassTypes_reference:  make(map[*CodeAircraftWingspanClassType]*CodeAircraftWingspanClassType),

		CodeAltitudeAdjustmentType_stagedOrder: make(map[*CodeAltitudeAdjustmentType]uint),
		CodeAltitudeAdjustmentType_orderStaged: make(map[uint]*CodeAltitudeAdjustmentType),
		CodeAltitudeAdjustmentTypes_reference:  make(map[*CodeAltitudeAdjustmentType]*CodeAltitudeAdjustmentType),

		CodeAltitudeUseType_stagedOrder: make(map[*CodeAltitudeUseType]uint),
		CodeAltitudeUseType_orderStaged: make(map[uint]*CodeAltitudeUseType),
		CodeAltitudeUseTypes_reference:  make(map[*CodeAltitudeUseType]*CodeAltitudeUseType),

		CodeAuthorityType_stagedOrder: make(map[*CodeAuthorityType]uint),
		CodeAuthorityType_orderStaged: make(map[uint]*CodeAuthorityType),
		CodeAuthorityTypes_reference:  make(map[*CodeAuthorityType]*CodeAuthorityType),

		CodeCardinalDirectionType_stagedOrder: make(map[*CodeCardinalDirectionType]uint),
		CodeCardinalDirectionType_orderStaged: make(map[uint]*CodeCardinalDirectionType),
		CodeCardinalDirectionTypes_reference:  make(map[*CodeCardinalDirectionType]*CodeCardinalDirectionType),

		CodeCommunicationModeType_stagedOrder: make(map[*CodeCommunicationModeType]uint),
		CodeCommunicationModeType_orderStaged: make(map[uint]*CodeCommunicationModeType),
		CodeCommunicationModeTypes_reference:  make(map[*CodeCommunicationModeType]*CodeCommunicationModeType),

		CodeComparisonType_stagedOrder: make(map[*CodeComparisonType]uint),
		CodeComparisonType_orderStaged: make(map[uint]*CodeComparisonType),
		CodeComparisonTypes_reference:  make(map[*CodeComparisonType]*CodeComparisonType),

		CodeCourseType_stagedOrder: make(map[*CodeCourseType]uint),
		CodeCourseType_orderStaged: make(map[uint]*CodeCourseType),
		CodeCourseTypes_reference:  make(map[*CodeCourseType]*CodeCourseType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeDirectionTurnType_stagedOrder: make(map[*CodeDirectionTurnType]uint),
		CodeDirectionTurnType_orderStaged: make(map[uint]*CodeDirectionTurnType),
		CodeDirectionTurnTypes_reference:  make(map[*CodeDirectionTurnType]*CodeDirectionTurnType),

		CodeDirectionType_stagedOrder: make(map[*CodeDirectionType]uint),
		CodeDirectionType_orderStaged: make(map[uint]*CodeDirectionType),
		CodeDirectionTypes_reference:  make(map[*CodeDirectionType]*CodeDirectionType),

		CodeEquipmentAntiCollisionType_stagedOrder: make(map[*CodeEquipmentAntiCollisionType]uint),
		CodeEquipmentAntiCollisionType_orderStaged: make(map[uint]*CodeEquipmentAntiCollisionType),
		CodeEquipmentAntiCollisionTypes_reference:  make(map[*CodeEquipmentAntiCollisionType]*CodeEquipmentAntiCollisionType),

		CodeFlightOriginType_stagedOrder: make(map[*CodeFlightOriginType]uint),
		CodeFlightOriginType_orderStaged: make(map[uint]*CodeFlightOriginType),
		CodeFlightOriginTypes_reference:  make(map[*CodeFlightOriginType]*CodeFlightOriginType),

		CodeFlightPurposeType_stagedOrder: make(map[*CodeFlightPurposeType]uint),
		CodeFlightPurposeType_orderStaged: make(map[uint]*CodeFlightPurposeType),
		CodeFlightPurposeTypes_reference:  make(map[*CodeFlightPurposeType]*CodeFlightPurposeType),

		CodeFlightRestrictionDesignatorType_stagedOrder: make(map[*CodeFlightRestrictionDesignatorType]uint),
		CodeFlightRestrictionDesignatorType_orderStaged: make(map[uint]*CodeFlightRestrictionDesignatorType),
		CodeFlightRestrictionDesignatorTypes_reference:  make(map[*CodeFlightRestrictionDesignatorType]*CodeFlightRestrictionDesignatorType),

		CodeFlightRestrictionType_stagedOrder: make(map[*CodeFlightRestrictionType]uint),
		CodeFlightRestrictionType_orderStaged: make(map[uint]*CodeFlightRestrictionType),
		CodeFlightRestrictionTypes_reference:  make(map[*CodeFlightRestrictionType]*CodeFlightRestrictionType),

		CodeFlightRuleType_stagedOrder: make(map[*CodeFlightRuleType]uint),
		CodeFlightRuleType_orderStaged: make(map[uint]*CodeFlightRuleType),
		CodeFlightRuleTypes_reference:  make(map[*CodeFlightRuleType]*CodeFlightRuleType),

		CodeFlightStatusType_stagedOrder: make(map[*CodeFlightStatusType]uint),
		CodeFlightStatusType_orderStaged: make(map[uint]*CodeFlightStatusType),
		CodeFlightStatusTypes_reference:  make(map[*CodeFlightStatusType]*CodeFlightStatusType),

		CodeFlightType_stagedOrder: make(map[*CodeFlightType]uint),
		CodeFlightType_orderStaged: make(map[uint]*CodeFlightType),
		CodeFlightTypes_reference:  make(map[*CodeFlightType]*CodeFlightType),

		CodeFlowConditionOperationType_stagedOrder: make(map[*CodeFlowConditionOperationType]uint),
		CodeFlowConditionOperationType_orderStaged: make(map[uint]*CodeFlowConditionOperationType),
		CodeFlowConditionOperationTypes_reference:  make(map[*CodeFlowConditionOperationType]*CodeFlowConditionOperationType),

		CodeFreeFlightType_stagedOrder: make(map[*CodeFreeFlightType]uint),
		CodeFreeFlightType_orderStaged: make(map[uint]*CodeFreeFlightType),
		CodeFreeFlightTypes_reference:  make(map[*CodeFreeFlightType]*CodeFreeFlightType),

		CodeLevelType_stagedOrder: make(map[*CodeLevelType]uint),
		CodeLevelType_orderStaged: make(map[uint]*CodeLevelType),
		CodeLevelTypes_reference:  make(map[*CodeLevelType]*CodeLevelType),

		CodeLocationQualifierType_stagedOrder: make(map[*CodeLocationQualifierType]uint),
		CodeLocationQualifierType_orderStaged: make(map[uint]*CodeLocationQualifierType),
		CodeLocationQualifierTypes_reference:  make(map[*CodeLocationQualifierType]*CodeLocationQualifierType),

		CodeMeteoConditionsType_stagedOrder: make(map[*CodeMeteoConditionsType]uint),
		CodeMeteoConditionsType_orderStaged: make(map[uint]*CodeMeteoConditionsType),
		CodeMeteoConditionsTypes_reference:  make(map[*CodeMeteoConditionsType]*CodeMeteoConditionsType),

		CodeMilitaryRoutePointType_stagedOrder: make(map[*CodeMilitaryRoutePointType]uint),
		CodeMilitaryRoutePointType_orderStaged: make(map[uint]*CodeMilitaryRoutePointType),
		CodeMilitaryRoutePointTypes_reference:  make(map[*CodeMilitaryRoutePointType]*CodeMilitaryRoutePointType),

		CodeMilitaryStatusType_stagedOrder: make(map[*CodeMilitaryStatusType]uint),
		CodeMilitaryStatusType_orderStaged: make(map[uint]*CodeMilitaryStatusType),
		CodeMilitaryStatusTypes_reference:  make(map[*CodeMilitaryStatusType]*CodeMilitaryStatusType),

		CodeMilitaryTrainingType_stagedOrder: make(map[*CodeMilitaryTrainingType]uint),
		CodeMilitaryTrainingType_orderStaged: make(map[uint]*CodeMilitaryTrainingType),
		CodeMilitaryTrainingTypes_reference:  make(map[*CodeMilitaryTrainingType]*CodeMilitaryTrainingType),

		CodeNavigationEquipmentType_stagedOrder: make(map[*CodeNavigationEquipmentType]uint),
		CodeNavigationEquipmentType_orderStaged: make(map[uint]*CodeNavigationEquipmentType),
		CodeNavigationEquipmentTypes_reference:  make(map[*CodeNavigationEquipmentType]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationType_stagedOrder: make(map[*CodeNavigationSpecificationType]uint),
		CodeNavigationSpecificationType_orderStaged: make(map[uint]*CodeNavigationSpecificationType),
		CodeNavigationSpecificationTypes_reference:  make(map[*CodeNavigationSpecificationType]*CodeNavigationSpecificationType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodeObstacleAssessmentSurfaceType_stagedOrder: make(map[*CodeObstacleAssessmentSurfaceType]uint),
		CodeObstacleAssessmentSurfaceType_orderStaged: make(map[uint]*CodeObstacleAssessmentSurfaceType),
		CodeObstacleAssessmentSurfaceTypes_reference:  make(map[*CodeObstacleAssessmentSurfaceType]*CodeObstacleAssessmentSurfaceType),

		CodeObstructionIdSurfaceZoneType_stagedOrder: make(map[*CodeObstructionIdSurfaceZoneType]uint),
		CodeObstructionIdSurfaceZoneType_orderStaged: make(map[uint]*CodeObstructionIdSurfaceZoneType),
		CodeObstructionIdSurfaceZoneTypes_reference:  make(map[*CodeObstructionIdSurfaceZoneType]*CodeObstructionIdSurfaceZoneType),

		CodeRNPType_stagedOrder: make(map[*CodeRNPType]uint),
		CodeRNPType_orderStaged: make(map[uint]*CodeRNPType),
		CodeRNPTypes_reference:  make(map[*CodeRNPType]*CodeRNPType),

		CodeRVSMPointRoleType_stagedOrder: make(map[*CodeRVSMPointRoleType]uint),
		CodeRVSMPointRoleType_orderStaged: make(map[uint]*CodeRVSMPointRoleType),
		CodeRVSMPointRoleTypes_reference:  make(map[*CodeRVSMPointRoleType]*CodeRVSMPointRoleType),

		CodeRVSMType_stagedOrder: make(map[*CodeRVSMType]uint),
		CodeRVSMType_orderStaged: make(map[uint]*CodeRVSMType),
		CodeRVSMTypes_reference:  make(map[*CodeRVSMType]*CodeRVSMType),

		CodeReferenceRoleType_stagedOrder: make(map[*CodeReferenceRoleType]uint),
		CodeReferenceRoleType_orderStaged: make(map[uint]*CodeReferenceRoleType),
		CodeReferenceRoleTypes_reference:  make(map[*CodeReferenceRoleType]*CodeReferenceRoleType),

		CodeRouteAvailabilityType_stagedOrder: make(map[*CodeRouteAvailabilityType]uint),
		CodeRouteAvailabilityType_orderStaged: make(map[uint]*CodeRouteAvailabilityType),
		CodeRouteAvailabilityTypes_reference:  make(map[*CodeRouteAvailabilityType]*CodeRouteAvailabilityType),

		CodeRouteDesignatorLetterType_stagedOrder: make(map[*CodeRouteDesignatorLetterType]uint),
		CodeRouteDesignatorLetterType_orderStaged: make(map[uint]*CodeRouteDesignatorLetterType),
		CodeRouteDesignatorLetterTypes_reference:  make(map[*CodeRouteDesignatorLetterType]*CodeRouteDesignatorLetterType),

		CodeRouteDesignatorPrefixType_stagedOrder: make(map[*CodeRouteDesignatorPrefixType]uint),
		CodeRouteDesignatorPrefixType_orderStaged: make(map[uint]*CodeRouteDesignatorPrefixType),
		CodeRouteDesignatorPrefixTypes_reference:  make(map[*CodeRouteDesignatorPrefixType]*CodeRouteDesignatorPrefixType),

		CodeRouteDesignatorSuffixType_stagedOrder: make(map[*CodeRouteDesignatorSuffixType]uint),
		CodeRouteDesignatorSuffixType_orderStaged: make(map[uint]*CodeRouteDesignatorSuffixType),
		CodeRouteDesignatorSuffixTypes_reference:  make(map[*CodeRouteDesignatorSuffixType]*CodeRouteDesignatorSuffixType),

		CodeRouteNavigationType_stagedOrder: make(map[*CodeRouteNavigationType]uint),
		CodeRouteNavigationType_orderStaged: make(map[uint]*CodeRouteNavigationType),
		CodeRouteNavigationTypes_reference:  make(map[*CodeRouteNavigationType]*CodeRouteNavigationType),

		CodeRouteOriginType_stagedOrder: make(map[*CodeRouteOriginType]uint),
		CodeRouteOriginType_orderStaged: make(map[uint]*CodeRouteOriginType),
		CodeRouteOriginTypes_reference:  make(map[*CodeRouteOriginType]*CodeRouteOriginType),

		CodeRouteSegmentPathType_stagedOrder: make(map[*CodeRouteSegmentPathType]uint),
		CodeRouteSegmentPathType_orderStaged: make(map[uint]*CodeRouteSegmentPathType),
		CodeRouteSegmentPathTypes_reference:  make(map[*CodeRouteSegmentPathType]*CodeRouteSegmentPathType),

		CodeRouteType_stagedOrder: make(map[*CodeRouteType]uint),
		CodeRouteType_orderStaged: make(map[uint]*CodeRouteType),
		CodeRouteTypes_reference:  make(map[*CodeRouteType]*CodeRouteType),

		CodeRuleProcedureTitleType_stagedOrder: make(map[*CodeRuleProcedureTitleType]uint),
		CodeRuleProcedureTitleType_orderStaged: make(map[uint]*CodeRuleProcedureTitleType),
		CodeRuleProcedureTitleTypes_reference:  make(map[*CodeRuleProcedureTitleType]*CodeRuleProcedureTitleType),

		CodeRuleProcedureType_stagedOrder: make(map[*CodeRuleProcedureType]uint),
		CodeRuleProcedureType_orderStaged: make(map[uint]*CodeRuleProcedureType),
		CodeRuleProcedureTypes_reference:  make(map[*CodeRuleProcedureType]*CodeRuleProcedureType),

		CodeSideType_stagedOrder: make(map[*CodeSideType]uint),
		CodeSideType_orderStaged: make(map[uint]*CodeSideType),
		CodeSideTypes_reference:  make(map[*CodeSideType]*CodeSideType),

		CodeSpeedReferenceType_stagedOrder: make(map[*CodeSpeedReferenceType]uint),
		CodeSpeedReferenceType_orderStaged: make(map[uint]*CodeSpeedReferenceType),
		CodeSpeedReferenceTypes_reference:  make(map[*CodeSpeedReferenceType]*CodeSpeedReferenceType),

		CodeTACANChannelType_stagedOrder: make(map[*CodeTACANChannelType]uint),
		CodeTACANChannelType_orderStaged: make(map[uint]*CodeTACANChannelType),
		CodeTACANChannelTypes_reference:  make(map[*CodeTACANChannelType]*CodeTACANChannelType),

		CodeTelecomNetworkType_stagedOrder: make(map[*CodeTelecomNetworkType]uint),
		CodeTelecomNetworkType_orderStaged: make(map[uint]*CodeTelecomNetworkType),
		CodeTelecomNetworkTypes_reference:  make(map[*CodeTelecomNetworkType]*CodeTelecomNetworkType),

		CodeTimeEventCombinationType_stagedOrder: make(map[*CodeTimeEventCombinationType]uint),
		CodeTimeEventCombinationType_orderStaged: make(map[uint]*CodeTimeEventCombinationType),
		CodeTimeEventCombinationTypes_reference:  make(map[*CodeTimeEventCombinationType]*CodeTimeEventCombinationType),

		CodeTimeEventType_stagedOrder: make(map[*CodeTimeEventType]uint),
		CodeTimeEventType_orderStaged: make(map[uint]*CodeTimeEventType),
		CodeTimeEventTypes_reference:  make(map[*CodeTimeEventType]*CodeTimeEventType),

		CodeTimeReferenceType_stagedOrder: make(map[*CodeTimeReferenceType]uint),
		CodeTimeReferenceType_orderStaged: make(map[uint]*CodeTimeReferenceType),
		CodeTimeReferenceTypes_reference:  make(map[*CodeTimeReferenceType]*CodeTimeReferenceType),

		CodeTransponderType_stagedOrder: make(map[*CodeTransponderType]uint),
		CodeTransponderType_orderStaged: make(map[uint]*CodeTransponderType),
		CodeTransponderTypes_reference:  make(map[*CodeTransponderType]*CodeTransponderType),

		CodeUpperAlphaType_stagedOrder: make(map[*CodeUpperAlphaType]uint),
		CodeUpperAlphaType_orderStaged: make(map[uint]*CodeUpperAlphaType),
		CodeUpperAlphaTypes_reference:  make(map[*CodeUpperAlphaType]*CodeUpperAlphaType),

		CodeValueInterpretationType_stagedOrder: make(map[*CodeValueInterpretationType]uint),
		CodeValueInterpretationType_orderStaged: make(map[uint]*CodeValueInterpretationType),
		CodeValueInterpretationTypes_reference:  make(map[*CodeValueInterpretationType]*CodeValueInterpretationType),

		CodeVerticalDatumType_stagedOrder: make(map[*CodeVerticalDatumType]uint),
		CodeVerticalDatumType_orderStaged: make(map[uint]*CodeVerticalDatumType),
		CodeVerticalDatumTypes_reference:  make(map[*CodeVerticalDatumType]*CodeVerticalDatumType),

		CodeVerticalReferenceType_stagedOrder: make(map[*CodeVerticalReferenceType]uint),
		CodeVerticalReferenceType_orderStaged: make(map[uint]*CodeVerticalReferenceType),
		CodeVerticalReferenceTypes_reference:  make(map[*CodeVerticalReferenceType]*CodeVerticalReferenceType),

		CodeWakeTurbulenceType_stagedOrder: make(map[*CodeWakeTurbulenceType]uint),
		CodeWakeTurbulenceType_orderStaged: make(map[uint]*CodeWakeTurbulenceType),
		CodeWakeTurbulenceTypes_reference:  make(map[*CodeWakeTurbulenceType]*CodeWakeTurbulenceType),

		CodeYesNoType_stagedOrder: make(map[*CodeYesNoType]uint),
		CodeYesNoType_orderStaged: make(map[uint]*CodeYesNoType),
		CodeYesNoTypes_reference:  make(map[*CodeYesNoType]*CodeYesNoType),

		DateMonthDayType_stagedOrder: make(map[*DateMonthDayType]uint),
		DateMonthDayType_orderStaged: make(map[uint]*DateMonthDayType),
		DateMonthDayTypes_reference:  make(map[*DateMonthDayType]*DateMonthDayType),

		NoNumberType_stagedOrder: make(map[*NoNumberType]uint),
		NoNumberType_orderStaged: make(map[uint]*NoNumberType),
		NoNumberTypes_reference:  make(map[*NoNumberType]*NoNumberType),

		NoSequenceType_stagedOrder: make(map[*NoSequenceType]uint),
		NoSequenceType_orderStaged: make(map[uint]*NoSequenceType),
		NoSequenceTypes_reference:  make(map[*NoSequenceType]*NoSequenceType),

		TextAddressType_stagedOrder: make(map[*TextAddressType]uint),
		TextAddressType_orderStaged: make(map[uint]*TextAddressType),
		TextAddressTypes_reference:  make(map[*TextAddressType]*TextAddressType),

		TextDesignatorType_stagedOrder: make(map[*TextDesignatorType]uint),
		TextDesignatorType_orderStaged: make(map[uint]*TextDesignatorType),
		TextDesignatorTypes_reference:  make(map[*TextDesignatorType]*TextDesignatorType),

		TextInstructionType_stagedOrder: make(map[*TextInstructionType]uint),
		TextInstructionType_orderStaged: make(map[uint]*TextInstructionType),
		TextInstructionTypes_reference:  make(map[*TextInstructionType]*TextInstructionType),

		TextNameType_stagedOrder: make(map[*TextNameType]uint),
		TextNameType_orderStaged: make(map[uint]*TextNameType),
		TextNameTypes_reference:  make(map[*TextNameType]*TextNameType),

		TextNoteType_stagedOrder: make(map[*TextNoteType]uint),
		TextNoteType_orderStaged: make(map[uint]*TextNoteType),
		TextNoteTypes_reference:  make(map[*TextNoteType]*TextNoteType),

		TextPhoneType_stagedOrder: make(map[*TextPhoneType]uint),
		TextPhoneType_orderStaged: make(map[uint]*TextPhoneType),
		TextPhoneTypes_reference:  make(map[*TextPhoneType]*TextPhoneType),

		TextPropertyNameType_stagedOrder: make(map[*TextPropertyNameType]uint),
		TextPropertyNameType_orderStaged: make(map[uint]*TextPropertyNameType),
		TextPropertyNameTypes_reference:  make(map[*TextPropertyNameType]*TextPropertyNameType),

		TimeType_stagedOrder: make(map[*TimeType]uint),
		TimeType_orderStaged: make(map[uint]*TimeType),
		TimeTypes_reference:  make(map[*TimeType]*TimeType),

		ValAngleType_stagedOrder: make(map[*ValAngleType]uint),
		ValAngleType_orderStaged: make(map[uint]*ValAngleType),
		ValAngleTypes_reference:  make(map[*ValAngleType]*ValAngleType),

		ValBearingType_stagedOrder: make(map[*ValBearingType]uint),
		ValBearingType_orderStaged: make(map[uint]*ValBearingType),
		ValBearingTypes_reference:  make(map[*ValBearingType]*ValBearingType),

		ValDistanceSignedType_stagedOrder: make(map[*ValDistanceSignedType]uint),
		ValDistanceSignedType_orderStaged: make(map[uint]*ValDistanceSignedType),
		ValDistanceSignedTypes_reference:  make(map[*ValDistanceSignedType]*ValDistanceSignedType),

		ValDistanceType_stagedOrder: make(map[*ValDistanceType]uint),
		ValDistanceType_orderStaged: make(map[uint]*ValDistanceType),
		ValDistanceTypes_reference:  make(map[*ValDistanceType]*ValDistanceType),

		ValDistanceVerticalType_stagedOrder: make(map[*ValDistanceVerticalType]uint),
		ValDistanceVerticalType_orderStaged: make(map[uint]*ValDistanceVerticalType),
		ValDistanceVerticalTypes_reference:  make(map[*ValDistanceVerticalType]*ValDistanceVerticalType),

		ValDurationType_stagedOrder: make(map[*ValDurationType]uint),
		ValDurationType_orderStaged: make(map[uint]*ValDurationType),
		ValDurationTypes_reference:  make(map[*ValDurationType]*ValDurationType),

		ValSlopeType_stagedOrder: make(map[*ValSlopeType]uint),
		ValSlopeType_orderStaged: make(map[uint]*ValSlopeType),
		ValSlopeTypes_reference:  make(map[*ValSlopeType]*ValSlopeType),

		ValSpeedType_stagedOrder: make(map[*ValSpeedType]uint),
		ValSpeedType_orderStaged: make(map[uint]*ValSpeedType),
		ValSpeedTypes_reference:  make(map[*ValSpeedType]*ValSpeedType),

		ValWeightType_stagedOrder: make(map[*ValWeightType]uint),
		ValWeightType_orderStaged: make(map[uint]*ValWeightType),
		ValWeightTypes_reference:  make(map[*ValWeightType]*ValWeightType),

		XHTMLType_stagedOrder: make(map[*XHTMLType]uint),
		XHTMLType_orderStaged: make(map[uint]*XHTMLType),
		XHTMLTypes_reference:  make(map[*XHTMLType]*XHTMLType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeATCReportingType": &CodeATCReportingTypeUnmarshaller{},

			"CodeAerialRefuellingPointType": &CodeAerialRefuellingPointTypeUnmarshaller{},

			"CodeAerialRefuellingPrefixType": &CodeAerialRefuellingPrefixTypeUnmarshaller{},

			"CodeAerialRefuellingType": &CodeAerialRefuellingTypeUnmarshaller{},

			"CodeAircraftCategoryType": &CodeAircraftCategoryTypeUnmarshaller{},

			"CodeAircraftEngineNumberType": &CodeAircraftEngineNumberTypeUnmarshaller{},

			"CodeAircraftEngineType": &CodeAircraftEngineTypeUnmarshaller{},

			"CodeAircraftICAOType": &CodeAircraftICAOTypeUnmarshaller{},

			"CodeAircraftType": &CodeAircraftTypeUnmarshaller{},

			"CodeAircraftWingspanClassType": &CodeAircraftWingspanClassTypeUnmarshaller{},

			"CodeAltitudeAdjustmentType": &CodeAltitudeAdjustmentTypeUnmarshaller{},

			"CodeAltitudeUseType": &CodeAltitudeUseTypeUnmarshaller{},

			"CodeAuthorityType": &CodeAuthorityTypeUnmarshaller{},

			"CodeCardinalDirectionType": &CodeCardinalDirectionTypeUnmarshaller{},

			"CodeCommunicationModeType": &CodeCommunicationModeTypeUnmarshaller{},

			"CodeComparisonType": &CodeComparisonTypeUnmarshaller{},

			"CodeCourseType": &CodeCourseTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeDirectionTurnType": &CodeDirectionTurnTypeUnmarshaller{},

			"CodeDirectionType": &CodeDirectionTypeUnmarshaller{},

			"CodeEquipmentAntiCollisionType": &CodeEquipmentAntiCollisionTypeUnmarshaller{},

			"CodeFlightOriginType": &CodeFlightOriginTypeUnmarshaller{},

			"CodeFlightPurposeType": &CodeFlightPurposeTypeUnmarshaller{},

			"CodeFlightRestrictionDesignatorType": &CodeFlightRestrictionDesignatorTypeUnmarshaller{},

			"CodeFlightRestrictionType": &CodeFlightRestrictionTypeUnmarshaller{},

			"CodeFlightRuleType": &CodeFlightRuleTypeUnmarshaller{},

			"CodeFlightStatusType": &CodeFlightStatusTypeUnmarshaller{},

			"CodeFlightType": &CodeFlightTypeUnmarshaller{},

			"CodeFlowConditionOperationType": &CodeFlowConditionOperationTypeUnmarshaller{},

			"CodeFreeFlightType": &CodeFreeFlightTypeUnmarshaller{},

			"CodeLevelType": &CodeLevelTypeUnmarshaller{},

			"CodeLocationQualifierType": &CodeLocationQualifierTypeUnmarshaller{},

			"CodeMeteoConditionsType": &CodeMeteoConditionsTypeUnmarshaller{},

			"CodeMilitaryRoutePointType": &CodeMilitaryRoutePointTypeUnmarshaller{},

			"CodeMilitaryStatusType": &CodeMilitaryStatusTypeUnmarshaller{},

			"CodeMilitaryTrainingType": &CodeMilitaryTrainingTypeUnmarshaller{},

			"CodeNavigationEquipmentType": &CodeNavigationEquipmentTypeUnmarshaller{},

			"CodeNavigationSpecificationType": &CodeNavigationSpecificationTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeObstacleAssessmentSurfaceType": &CodeObstacleAssessmentSurfaceTypeUnmarshaller{},

			"CodeObstructionIdSurfaceZoneType": &CodeObstructionIdSurfaceZoneTypeUnmarshaller{},

			"CodeRNPType": &CodeRNPTypeUnmarshaller{},

			"CodeRVSMPointRoleType": &CodeRVSMPointRoleTypeUnmarshaller{},

			"CodeRVSMType": &CodeRVSMTypeUnmarshaller{},

			"CodeReferenceRoleType": &CodeReferenceRoleTypeUnmarshaller{},

			"CodeRouteAvailabilityType": &CodeRouteAvailabilityTypeUnmarshaller{},

			"CodeRouteDesignatorLetterType": &CodeRouteDesignatorLetterTypeUnmarshaller{},

			"CodeRouteDesignatorPrefixType": &CodeRouteDesignatorPrefixTypeUnmarshaller{},

			"CodeRouteDesignatorSuffixType": &CodeRouteDesignatorSuffixTypeUnmarshaller{},

			"CodeRouteNavigationType": &CodeRouteNavigationTypeUnmarshaller{},

			"CodeRouteOriginType": &CodeRouteOriginTypeUnmarshaller{},

			"CodeRouteSegmentPathType": &CodeRouteSegmentPathTypeUnmarshaller{},

			"CodeRouteType": &CodeRouteTypeUnmarshaller{},

			"CodeRuleProcedureTitleType": &CodeRuleProcedureTitleTypeUnmarshaller{},

			"CodeRuleProcedureType": &CodeRuleProcedureTypeUnmarshaller{},

			"CodeSideType": &CodeSideTypeUnmarshaller{},

			"CodeSpeedReferenceType": &CodeSpeedReferenceTypeUnmarshaller{},

			"CodeTACANChannelType": &CodeTACANChannelTypeUnmarshaller{},

			"CodeTelecomNetworkType": &CodeTelecomNetworkTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeTransponderType": &CodeTransponderTypeUnmarshaller{},

			"CodeUpperAlphaType": &CodeUpperAlphaTypeUnmarshaller{},

			"CodeValueInterpretationType": &CodeValueInterpretationTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeVerticalReferenceType": &CodeVerticalReferenceTypeUnmarshaller{},

			"CodeWakeTurbulenceType": &CodeWakeTurbulenceTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"NoNumberType": &NoNumberTypeUnmarshaller{},

			"NoSequenceType": &NoSequenceTypeUnmarshaller{},

			"TextAddressType": &TextAddressTypeUnmarshaller{},

			"TextDesignatorType": &TextDesignatorTypeUnmarshaller{},

			"TextInstructionType": &TextInstructionTypeUnmarshaller{},

			"TextNameType": &TextNameTypeUnmarshaller{},

			"TextNoteType": &TextNoteTypeUnmarshaller{},

			"TextPhoneType": &TextPhoneTypeUnmarshaller{},

			"TextPropertyNameType": &TextPropertyNameTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"ValAngleType": &ValAngleTypeUnmarshaller{},

			"ValBearingType": &ValBearingTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValSlopeType": &ValSlopeTypeUnmarshaller{},

			"ValSpeedType": &ValSpeedTypeUnmarshaller{},

			"ValWeightType": &ValWeightTypeUnmarshaller{},

			"XHTMLType": &XHTMLTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeATCReportingType"] = len(stage.CodeATCReportingTypes)
	stage.Map_GongStructName_InstancesNb["CodeAerialRefuellingPointType"] = len(stage.CodeAerialRefuellingPointTypes)
	stage.Map_GongStructName_InstancesNb["CodeAerialRefuellingPrefixType"] = len(stage.CodeAerialRefuellingPrefixTypes)
	stage.Map_GongStructName_InstancesNb["CodeAerialRefuellingType"] = len(stage.CodeAerialRefuellingTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftCategoryType"] = len(stage.CodeAircraftCategoryTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineNumberType"] = len(stage.CodeAircraftEngineNumberTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineType"] = len(stage.CodeAircraftEngineTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftICAOType"] = len(stage.CodeAircraftICAOTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftType"] = len(stage.CodeAircraftTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftWingspanClassType"] = len(stage.CodeAircraftWingspanClassTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeAdjustmentType"] = len(stage.CodeAltitudeAdjustmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeUseType"] = len(stage.CodeAltitudeUseTypes)
	stage.Map_GongStructName_InstancesNb["CodeAuthorityType"] = len(stage.CodeAuthorityTypes)
	stage.Map_GongStructName_InstancesNb["CodeCardinalDirectionType"] = len(stage.CodeCardinalDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationModeType"] = len(stage.CodeCommunicationModeTypes)
	stage.Map_GongStructName_InstancesNb["CodeComparisonType"] = len(stage.CodeComparisonTypes)
	stage.Map_GongStructName_InstancesNb["CodeCourseType"] = len(stage.CodeCourseTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionTurnType"] = len(stage.CodeDirectionTurnTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionType"] = len(stage.CodeDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeEquipmentAntiCollisionType"] = len(stage.CodeEquipmentAntiCollisionTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightOriginType"] = len(stage.CodeFlightOriginTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightPurposeType"] = len(stage.CodeFlightPurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightRestrictionDesignatorType"] = len(stage.CodeFlightRestrictionDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightRestrictionType"] = len(stage.CodeFlightRestrictionTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightRuleType"] = len(stage.CodeFlightRuleTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightStatusType"] = len(stage.CodeFlightStatusTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightType"] = len(stage.CodeFlightTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlowConditionOperationType"] = len(stage.CodeFlowConditionOperationTypes)
	stage.Map_GongStructName_InstancesNb["CodeFreeFlightType"] = len(stage.CodeFreeFlightTypes)
	stage.Map_GongStructName_InstancesNb["CodeLevelType"] = len(stage.CodeLevelTypes)
	stage.Map_GongStructName_InstancesNb["CodeLocationQualifierType"] = len(stage.CodeLocationQualifierTypes)
	stage.Map_GongStructName_InstancesNb["CodeMeteoConditionsType"] = len(stage.CodeMeteoConditionsTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryRoutePointType"] = len(stage.CodeMilitaryRoutePointTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryStatusType"] = len(stage.CodeMilitaryStatusTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryTrainingType"] = len(stage.CodeMilitaryTrainingTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationEquipmentType"] = len(stage.CodeNavigationEquipmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationSpecificationType"] = len(stage.CodeNavigationSpecificationTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstacleAssessmentSurfaceType"] = len(stage.CodeObstacleAssessmentSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstructionIdSurfaceZoneType"] = len(stage.CodeObstructionIdSurfaceZoneTypes)
	stage.Map_GongStructName_InstancesNb["CodeRNPType"] = len(stage.CodeRNPTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMPointRoleType"] = len(stage.CodeRVSMPointRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMType"] = len(stage.CodeRVSMTypes)
	stage.Map_GongStructName_InstancesNb["CodeReferenceRoleType"] = len(stage.CodeReferenceRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteAvailabilityType"] = len(stage.CodeRouteAvailabilityTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteDesignatorLetterType"] = len(stage.CodeRouteDesignatorLetterTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteDesignatorPrefixType"] = len(stage.CodeRouteDesignatorPrefixTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteDesignatorSuffixType"] = len(stage.CodeRouteDesignatorSuffixTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteNavigationType"] = len(stage.CodeRouteNavigationTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteOriginType"] = len(stage.CodeRouteOriginTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteSegmentPathType"] = len(stage.CodeRouteSegmentPathTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteType"] = len(stage.CodeRouteTypes)
	stage.Map_GongStructName_InstancesNb["CodeRuleProcedureTitleType"] = len(stage.CodeRuleProcedureTitleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRuleProcedureType"] = len(stage.CodeRuleProcedureTypes)
	stage.Map_GongStructName_InstancesNb["CodeSideType"] = len(stage.CodeSideTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpeedReferenceType"] = len(stage.CodeSpeedReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTACANChannelType"] = len(stage.CodeTACANChannelTypes)
	stage.Map_GongStructName_InstancesNb["CodeTelecomNetworkType"] = len(stage.CodeTelecomNetworkTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTransponderType"] = len(stage.CodeTransponderTypes)
	stage.Map_GongStructName_InstancesNb["CodeUpperAlphaType"] = len(stage.CodeUpperAlphaTypes)
	stage.Map_GongStructName_InstancesNb["CodeValueInterpretationType"] = len(stage.CodeValueInterpretationTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalReferenceType"] = len(stage.CodeVerticalReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeWakeTurbulenceType"] = len(stage.CodeWakeTurbulenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["NoNumberType"] = len(stage.NoNumberTypes)
	stage.Map_GongStructName_InstancesNb["NoSequenceType"] = len(stage.NoSequenceTypes)
	stage.Map_GongStructName_InstancesNb["TextAddressType"] = len(stage.TextAddressTypes)
	stage.Map_GongStructName_InstancesNb["TextDesignatorType"] = len(stage.TextDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["TextInstructionType"] = len(stage.TextInstructionTypes)
	stage.Map_GongStructName_InstancesNb["TextNameType"] = len(stage.TextNameTypes)
	stage.Map_GongStructName_InstancesNb["TextNoteType"] = len(stage.TextNoteTypes)
	stage.Map_GongStructName_InstancesNb["TextPhoneType"] = len(stage.TextPhoneTypes)
	stage.Map_GongStructName_InstancesNb["TextPropertyNameType"] = len(stage.TextPropertyNameTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["ValAngleType"] = len(stage.ValAngleTypes)
	stage.Map_GongStructName_InstancesNb["ValBearingType"] = len(stage.ValBearingTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValSlopeType"] = len(stage.ValSlopeTypes)
	stage.Map_GongStructName_InstancesNb["ValSpeedType"] = len(stage.ValSpeedTypes)
	stage.Map_GongStructName_InstancesNb["ValWeightType"] = len(stage.ValWeightTypes)
	stage.Map_GongStructName_InstancesNb["XHTMLType"] = len(stage.XHTMLTypes)
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
// Stage puts codeatcreportingtype to the model stage
func (codeatcreportingtype *CodeATCReportingType) Stage(stage *Stage) *CodeATCReportingType {
	__gong__stage(stage.CodeATCReportingTypes, stage.CodeATCReportingType_stagedOrder, stage.CodeATCReportingType_orderStaged, &stage.CodeATCReportingTypeOrder, stage.CodeATCReportingTypes_mapString, codeatcreportingtype, codeatcreportingtype.Name)
	return codeatcreportingtype
}

// StagePreserveOrder puts codeatcreportingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeATCReportingTypeOrder
// - update stage.CodeATCReportingTypeOrder accordingly
func (codeatcreportingtype *CodeATCReportingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeATCReportingTypes, stage.CodeATCReportingType_stagedOrder, stage.CodeATCReportingType_orderStaged, &stage.CodeATCReportingTypeOrder, stage.CodeATCReportingTypes_mapString, codeatcreportingtype, order, codeatcreportingtype.Name)
}

// Unstage removes codeatcreportingtype off the model stage
func (codeatcreportingtype *CodeATCReportingType) Unstage(stage *Stage) *CodeATCReportingType {
	__gong__unstage(stage.CodeATCReportingTypes, stage.CodeATCReportingTypes_mapString, codeatcreportingtype, codeatcreportingtype.Name)
	return codeatcreportingtype
}

// UnstageVoid removes codeatcreportingtype off the model stage
func (codeatcreportingtype *CodeATCReportingType) UnstageVoid(stage *Stage) {
	codeatcreportingtype.Unstage(stage)
}

func (codeatcreportingtype *CodeATCReportingType) StageVoid(stage *Stage) {
	codeatcreportingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeatcreportingtype *CodeATCReportingType) GetName() (res string) {
	return codeatcreportingtype.Name
}

// for satisfaction of GongStruct interface
func (codeatcreportingtype *CodeATCReportingType) SetName(name string) {
	codeatcreportingtype.Name = name
}

// Stage puts codeaerialrefuellingpointtype to the model stage
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) Stage(stage *Stage) *CodeAerialRefuellingPointType {
	__gong__stage(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointType_stagedOrder, stage.CodeAerialRefuellingPointType_orderStaged, &stage.CodeAerialRefuellingPointTypeOrder, stage.CodeAerialRefuellingPointTypes_mapString, codeaerialrefuellingpointtype, codeaerialrefuellingpointtype.Name)
	return codeaerialrefuellingpointtype
}

// StagePreserveOrder puts codeaerialrefuellingpointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAerialRefuellingPointTypeOrder
// - update stage.CodeAerialRefuellingPointTypeOrder accordingly
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointType_stagedOrder, stage.CodeAerialRefuellingPointType_orderStaged, &stage.CodeAerialRefuellingPointTypeOrder, stage.CodeAerialRefuellingPointTypes_mapString, codeaerialrefuellingpointtype, order, codeaerialrefuellingpointtype.Name)
}

// Unstage removes codeaerialrefuellingpointtype off the model stage
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) Unstage(stage *Stage) *CodeAerialRefuellingPointType {
	__gong__unstage(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointTypes_mapString, codeaerialrefuellingpointtype, codeaerialrefuellingpointtype.Name)
	return codeaerialrefuellingpointtype
}

// UnstageVoid removes codeaerialrefuellingpointtype off the model stage
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) UnstageVoid(stage *Stage) {
	codeaerialrefuellingpointtype.Unstage(stage)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) StageVoid(stage *Stage) {
	codeaerialrefuellingpointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GetName() (res string) {
	return codeaerialrefuellingpointtype.Name
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) SetName(name string) {
	codeaerialrefuellingpointtype.Name = name
}

// Stage puts codeaerialrefuellingprefixtype to the model stage
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) Stage(stage *Stage) *CodeAerialRefuellingPrefixType {
	__gong__stage(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixType_stagedOrder, stage.CodeAerialRefuellingPrefixType_orderStaged, &stage.CodeAerialRefuellingPrefixTypeOrder, stage.CodeAerialRefuellingPrefixTypes_mapString, codeaerialrefuellingprefixtype, codeaerialrefuellingprefixtype.Name)
	return codeaerialrefuellingprefixtype
}

// StagePreserveOrder puts codeaerialrefuellingprefixtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAerialRefuellingPrefixTypeOrder
// - update stage.CodeAerialRefuellingPrefixTypeOrder accordingly
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixType_stagedOrder, stage.CodeAerialRefuellingPrefixType_orderStaged, &stage.CodeAerialRefuellingPrefixTypeOrder, stage.CodeAerialRefuellingPrefixTypes_mapString, codeaerialrefuellingprefixtype, order, codeaerialrefuellingprefixtype.Name)
}

// Unstage removes codeaerialrefuellingprefixtype off the model stage
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) Unstage(stage *Stage) *CodeAerialRefuellingPrefixType {
	__gong__unstage(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixTypes_mapString, codeaerialrefuellingprefixtype, codeaerialrefuellingprefixtype.Name)
	return codeaerialrefuellingprefixtype
}

// UnstageVoid removes codeaerialrefuellingprefixtype off the model stage
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) UnstageVoid(stage *Stage) {
	codeaerialrefuellingprefixtype.Unstage(stage)
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) StageVoid(stage *Stage) {
	codeaerialrefuellingprefixtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GetName() (res string) {
	return codeaerialrefuellingprefixtype.Name
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) SetName(name string) {
	codeaerialrefuellingprefixtype.Name = name
}

// Stage puts codeaerialrefuellingtype to the model stage
func (codeaerialrefuellingtype *CodeAerialRefuellingType) Stage(stage *Stage) *CodeAerialRefuellingType {
	__gong__stage(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingType_stagedOrder, stage.CodeAerialRefuellingType_orderStaged, &stage.CodeAerialRefuellingTypeOrder, stage.CodeAerialRefuellingTypes_mapString, codeaerialrefuellingtype, codeaerialrefuellingtype.Name)
	return codeaerialrefuellingtype
}

// StagePreserveOrder puts codeaerialrefuellingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAerialRefuellingTypeOrder
// - update stage.CodeAerialRefuellingTypeOrder accordingly
func (codeaerialrefuellingtype *CodeAerialRefuellingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingType_stagedOrder, stage.CodeAerialRefuellingType_orderStaged, &stage.CodeAerialRefuellingTypeOrder, stage.CodeAerialRefuellingTypes_mapString, codeaerialrefuellingtype, order, codeaerialrefuellingtype.Name)
}

// Unstage removes codeaerialrefuellingtype off the model stage
func (codeaerialrefuellingtype *CodeAerialRefuellingType) Unstage(stage *Stage) *CodeAerialRefuellingType {
	__gong__unstage(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingTypes_mapString, codeaerialrefuellingtype, codeaerialrefuellingtype.Name)
	return codeaerialrefuellingtype
}

// UnstageVoid removes codeaerialrefuellingtype off the model stage
func (codeaerialrefuellingtype *CodeAerialRefuellingType) UnstageVoid(stage *Stage) {
	codeaerialrefuellingtype.Unstage(stage)
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) StageVoid(stage *Stage) {
	codeaerialrefuellingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingtype *CodeAerialRefuellingType) GetName() (res string) {
	return codeaerialrefuellingtype.Name
}

// for satisfaction of GongStruct interface
func (codeaerialrefuellingtype *CodeAerialRefuellingType) SetName(name string) {
	codeaerialrefuellingtype.Name = name
}

// Stage puts codeaircraftcategorytype to the model stage
func (codeaircraftcategorytype *CodeAircraftCategoryType) Stage(stage *Stage) *CodeAircraftCategoryType {
	__gong__stage(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryType_stagedOrder, stage.CodeAircraftCategoryType_orderStaged, &stage.CodeAircraftCategoryTypeOrder, stage.CodeAircraftCategoryTypes_mapString, codeaircraftcategorytype, codeaircraftcategorytype.Name)
	return codeaircraftcategorytype
}

// StagePreserveOrder puts codeaircraftcategorytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftCategoryTypeOrder
// - update stage.CodeAircraftCategoryTypeOrder accordingly
func (codeaircraftcategorytype *CodeAircraftCategoryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryType_stagedOrder, stage.CodeAircraftCategoryType_orderStaged, &stage.CodeAircraftCategoryTypeOrder, stage.CodeAircraftCategoryTypes_mapString, codeaircraftcategorytype, order, codeaircraftcategorytype.Name)
}

// Unstage removes codeaircraftcategorytype off the model stage
func (codeaircraftcategorytype *CodeAircraftCategoryType) Unstage(stage *Stage) *CodeAircraftCategoryType {
	__gong__unstage(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryTypes_mapString, codeaircraftcategorytype, codeaircraftcategorytype.Name)
	return codeaircraftcategorytype
}

// UnstageVoid removes codeaircraftcategorytype off the model stage
func (codeaircraftcategorytype *CodeAircraftCategoryType) UnstageVoid(stage *Stage) {
	codeaircraftcategorytype.Unstage(stage)
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) StageVoid(stage *Stage) {
	codeaircraftcategorytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftcategorytype *CodeAircraftCategoryType) GetName() (res string) {
	return codeaircraftcategorytype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftcategorytype *CodeAircraftCategoryType) SetName(name string) {
	codeaircraftcategorytype.Name = name
}

// Stage puts codeaircraftenginenumbertype to the model stage
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) Stage(stage *Stage) *CodeAircraftEngineNumberType {
	__gong__stage(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberType_stagedOrder, stage.CodeAircraftEngineNumberType_orderStaged, &stage.CodeAircraftEngineNumberTypeOrder, stage.CodeAircraftEngineNumberTypes_mapString, codeaircraftenginenumbertype, codeaircraftenginenumbertype.Name)
	return codeaircraftenginenumbertype
}

// StagePreserveOrder puts codeaircraftenginenumbertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftEngineNumberTypeOrder
// - update stage.CodeAircraftEngineNumberTypeOrder accordingly
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberType_stagedOrder, stage.CodeAircraftEngineNumberType_orderStaged, &stage.CodeAircraftEngineNumberTypeOrder, stage.CodeAircraftEngineNumberTypes_mapString, codeaircraftenginenumbertype, order, codeaircraftenginenumbertype.Name)
}

// Unstage removes codeaircraftenginenumbertype off the model stage
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) Unstage(stage *Stage) *CodeAircraftEngineNumberType {
	__gong__unstage(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberTypes_mapString, codeaircraftenginenumbertype, codeaircraftenginenumbertype.Name)
	return codeaircraftenginenumbertype
}

// UnstageVoid removes codeaircraftenginenumbertype off the model stage
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) UnstageVoid(stage *Stage) {
	codeaircraftenginenumbertype.Unstage(stage)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) StageVoid(stage *Stage) {
	codeaircraftenginenumbertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GetName() (res string) {
	return codeaircraftenginenumbertype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) SetName(name string) {
	codeaircraftenginenumbertype.Name = name
}

// Stage puts codeaircraftenginetype to the model stage
func (codeaircraftenginetype *CodeAircraftEngineType) Stage(stage *Stage) *CodeAircraftEngineType {
	__gong__stage(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineType_stagedOrder, stage.CodeAircraftEngineType_orderStaged, &stage.CodeAircraftEngineTypeOrder, stage.CodeAircraftEngineTypes_mapString, codeaircraftenginetype, codeaircraftenginetype.Name)
	return codeaircraftenginetype
}

// StagePreserveOrder puts codeaircraftenginetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftEngineTypeOrder
// - update stage.CodeAircraftEngineTypeOrder accordingly
func (codeaircraftenginetype *CodeAircraftEngineType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineType_stagedOrder, stage.CodeAircraftEngineType_orderStaged, &stage.CodeAircraftEngineTypeOrder, stage.CodeAircraftEngineTypes_mapString, codeaircraftenginetype, order, codeaircraftenginetype.Name)
}

// Unstage removes codeaircraftenginetype off the model stage
func (codeaircraftenginetype *CodeAircraftEngineType) Unstage(stage *Stage) *CodeAircraftEngineType {
	__gong__unstage(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineTypes_mapString, codeaircraftenginetype, codeaircraftenginetype.Name)
	return codeaircraftenginetype
}

// UnstageVoid removes codeaircraftenginetype off the model stage
func (codeaircraftenginetype *CodeAircraftEngineType) UnstageVoid(stage *Stage) {
	codeaircraftenginetype.Unstage(stage)
}

func (codeaircraftenginetype *CodeAircraftEngineType) StageVoid(stage *Stage) {
	codeaircraftenginetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftenginetype *CodeAircraftEngineType) GetName() (res string) {
	return codeaircraftenginetype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftenginetype *CodeAircraftEngineType) SetName(name string) {
	codeaircraftenginetype.Name = name
}

// Stage puts codeaircrafticaotype to the model stage
func (codeaircrafticaotype *CodeAircraftICAOType) Stage(stage *Stage) *CodeAircraftICAOType {
	__gong__stage(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOType_stagedOrder, stage.CodeAircraftICAOType_orderStaged, &stage.CodeAircraftICAOTypeOrder, stage.CodeAircraftICAOTypes_mapString, codeaircrafticaotype, codeaircrafticaotype.Name)
	return codeaircrafticaotype
}

// StagePreserveOrder puts codeaircrafticaotype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftICAOTypeOrder
// - update stage.CodeAircraftICAOTypeOrder accordingly
func (codeaircrafticaotype *CodeAircraftICAOType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOType_stagedOrder, stage.CodeAircraftICAOType_orderStaged, &stage.CodeAircraftICAOTypeOrder, stage.CodeAircraftICAOTypes_mapString, codeaircrafticaotype, order, codeaircrafticaotype.Name)
}

// Unstage removes codeaircrafticaotype off the model stage
func (codeaircrafticaotype *CodeAircraftICAOType) Unstage(stage *Stage) *CodeAircraftICAOType {
	__gong__unstage(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_mapString, codeaircrafticaotype, codeaircrafticaotype.Name)
	return codeaircrafticaotype
}

// UnstageVoid removes codeaircrafticaotype off the model stage
func (codeaircrafticaotype *CodeAircraftICAOType) UnstageVoid(stage *Stage) {
	codeaircrafticaotype.Unstage(stage)
}

func (codeaircrafticaotype *CodeAircraftICAOType) StageVoid(stage *Stage) {
	codeaircrafticaotype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircrafticaotype *CodeAircraftICAOType) GetName() (res string) {
	return codeaircrafticaotype.Name
}

// for satisfaction of GongStruct interface
func (codeaircrafticaotype *CodeAircraftICAOType) SetName(name string) {
	codeaircrafticaotype.Name = name
}

// Stage puts codeaircrafttype to the model stage
func (codeaircrafttype *CodeAircraftType) Stage(stage *Stage) *CodeAircraftType {
	__gong__stage(stage.CodeAircraftTypes, stage.CodeAircraftType_stagedOrder, stage.CodeAircraftType_orderStaged, &stage.CodeAircraftTypeOrder, stage.CodeAircraftTypes_mapString, codeaircrafttype, codeaircrafttype.Name)
	return codeaircrafttype
}

// StagePreserveOrder puts codeaircrafttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftTypeOrder
// - update stage.CodeAircraftTypeOrder accordingly
func (codeaircrafttype *CodeAircraftType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftTypes, stage.CodeAircraftType_stagedOrder, stage.CodeAircraftType_orderStaged, &stage.CodeAircraftTypeOrder, stage.CodeAircraftTypes_mapString, codeaircrafttype, order, codeaircrafttype.Name)
}

// Unstage removes codeaircrafttype off the model stage
func (codeaircrafttype *CodeAircraftType) Unstage(stage *Stage) *CodeAircraftType {
	__gong__unstage(stage.CodeAircraftTypes, stage.CodeAircraftTypes_mapString, codeaircrafttype, codeaircrafttype.Name)
	return codeaircrafttype
}

// UnstageVoid removes codeaircrafttype off the model stage
func (codeaircrafttype *CodeAircraftType) UnstageVoid(stage *Stage) {
	codeaircrafttype.Unstage(stage)
}

func (codeaircrafttype *CodeAircraftType) StageVoid(stage *Stage) {
	codeaircrafttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircrafttype *CodeAircraftType) GetName() (res string) {
	return codeaircrafttype.Name
}

// for satisfaction of GongStruct interface
func (codeaircrafttype *CodeAircraftType) SetName(name string) {
	codeaircrafttype.Name = name
}

// Stage puts codeaircraftwingspanclasstype to the model stage
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) Stage(stage *Stage) *CodeAircraftWingspanClassType {
	__gong__stage(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassType_stagedOrder, stage.CodeAircraftWingspanClassType_orderStaged, &stage.CodeAircraftWingspanClassTypeOrder, stage.CodeAircraftWingspanClassTypes_mapString, codeaircraftwingspanclasstype, codeaircraftwingspanclasstype.Name)
	return codeaircraftwingspanclasstype
}

// StagePreserveOrder puts codeaircraftwingspanclasstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftWingspanClassTypeOrder
// - update stage.CodeAircraftWingspanClassTypeOrder accordingly
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassType_stagedOrder, stage.CodeAircraftWingspanClassType_orderStaged, &stage.CodeAircraftWingspanClassTypeOrder, stage.CodeAircraftWingspanClassTypes_mapString, codeaircraftwingspanclasstype, order, codeaircraftwingspanclasstype.Name)
}

// Unstage removes codeaircraftwingspanclasstype off the model stage
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) Unstage(stage *Stage) *CodeAircraftWingspanClassType {
	__gong__unstage(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_mapString, codeaircraftwingspanclasstype, codeaircraftwingspanclasstype.Name)
	return codeaircraftwingspanclasstype
}

// UnstageVoid removes codeaircraftwingspanclasstype off the model stage
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) UnstageVoid(stage *Stage) {
	codeaircraftwingspanclasstype.Unstage(stage)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) StageVoid(stage *Stage) {
	codeaircraftwingspanclasstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GetName() (res string) {
	return codeaircraftwingspanclasstype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) SetName(name string) {
	codeaircraftwingspanclasstype.Name = name
}

// Stage puts codealtitudeadjustmenttype to the model stage
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) Stage(stage *Stage) *CodeAltitudeAdjustmentType {
	__gong__stage(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentType_stagedOrder, stage.CodeAltitudeAdjustmentType_orderStaged, &stage.CodeAltitudeAdjustmentTypeOrder, stage.CodeAltitudeAdjustmentTypes_mapString, codealtitudeadjustmenttype, codealtitudeadjustmenttype.Name)
	return codealtitudeadjustmenttype
}

// StagePreserveOrder puts codealtitudeadjustmenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAltitudeAdjustmentTypeOrder
// - update stage.CodeAltitudeAdjustmentTypeOrder accordingly
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentType_stagedOrder, stage.CodeAltitudeAdjustmentType_orderStaged, &stage.CodeAltitudeAdjustmentTypeOrder, stage.CodeAltitudeAdjustmentTypes_mapString, codealtitudeadjustmenttype, order, codealtitudeadjustmenttype.Name)
}

// Unstage removes codealtitudeadjustmenttype off the model stage
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) Unstage(stage *Stage) *CodeAltitudeAdjustmentType {
	__gong__unstage(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentTypes_mapString, codealtitudeadjustmenttype, codealtitudeadjustmenttype.Name)
	return codealtitudeadjustmenttype
}

// UnstageVoid removes codealtitudeadjustmenttype off the model stage
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) UnstageVoid(stage *Stage) {
	codealtitudeadjustmenttype.Unstage(stage)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) StageVoid(stage *Stage) {
	codealtitudeadjustmenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GetName() (res string) {
	return codealtitudeadjustmenttype.Name
}

// for satisfaction of GongStruct interface
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) SetName(name string) {
	codealtitudeadjustmenttype.Name = name
}

// Stage puts codealtitudeusetype to the model stage
func (codealtitudeusetype *CodeAltitudeUseType) Stage(stage *Stage) *CodeAltitudeUseType {
	__gong__stage(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseType_orderStaged, &stage.CodeAltitudeUseTypeOrder, stage.CodeAltitudeUseTypes_mapString, codealtitudeusetype, codealtitudeusetype.Name)
	return codealtitudeusetype
}

// StagePreserveOrder puts codealtitudeusetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAltitudeUseTypeOrder
// - update stage.CodeAltitudeUseTypeOrder accordingly
func (codealtitudeusetype *CodeAltitudeUseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseType_orderStaged, &stage.CodeAltitudeUseTypeOrder, stage.CodeAltitudeUseTypes_mapString, codealtitudeusetype, order, codealtitudeusetype.Name)
}

// Unstage removes codealtitudeusetype off the model stage
func (codealtitudeusetype *CodeAltitudeUseType) Unstage(stage *Stage) *CodeAltitudeUseType {
	__gong__unstage(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_mapString, codealtitudeusetype, codealtitudeusetype.Name)
	return codealtitudeusetype
}

// UnstageVoid removes codealtitudeusetype off the model stage
func (codealtitudeusetype *CodeAltitudeUseType) UnstageVoid(stage *Stage) {
	codealtitudeusetype.Unstage(stage)
}

func (codealtitudeusetype *CodeAltitudeUseType) StageVoid(stage *Stage) {
	codealtitudeusetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codealtitudeusetype *CodeAltitudeUseType) GetName() (res string) {
	return codealtitudeusetype.Name
}

// for satisfaction of GongStruct interface
func (codealtitudeusetype *CodeAltitudeUseType) SetName(name string) {
	codealtitudeusetype.Name = name
}

// Stage puts codeauthoritytype to the model stage
func (codeauthoritytype *CodeAuthorityType) Stage(stage *Stage) *CodeAuthorityType {
	__gong__stage(stage.CodeAuthorityTypes, stage.CodeAuthorityType_stagedOrder, stage.CodeAuthorityType_orderStaged, &stage.CodeAuthorityTypeOrder, stage.CodeAuthorityTypes_mapString, codeauthoritytype, codeauthoritytype.Name)
	return codeauthoritytype
}

// StagePreserveOrder puts codeauthoritytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAuthorityTypeOrder
// - update stage.CodeAuthorityTypeOrder accordingly
func (codeauthoritytype *CodeAuthorityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAuthorityTypes, stage.CodeAuthorityType_stagedOrder, stage.CodeAuthorityType_orderStaged, &stage.CodeAuthorityTypeOrder, stage.CodeAuthorityTypes_mapString, codeauthoritytype, order, codeauthoritytype.Name)
}

// Unstage removes codeauthoritytype off the model stage
func (codeauthoritytype *CodeAuthorityType) Unstage(stage *Stage) *CodeAuthorityType {
	__gong__unstage(stage.CodeAuthorityTypes, stage.CodeAuthorityTypes_mapString, codeauthoritytype, codeauthoritytype.Name)
	return codeauthoritytype
}

// UnstageVoid removes codeauthoritytype off the model stage
func (codeauthoritytype *CodeAuthorityType) UnstageVoid(stage *Stage) {
	codeauthoritytype.Unstage(stage)
}

func (codeauthoritytype *CodeAuthorityType) StageVoid(stage *Stage) {
	codeauthoritytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeauthoritytype *CodeAuthorityType) GetName() (res string) {
	return codeauthoritytype.Name
}

// for satisfaction of GongStruct interface
func (codeauthoritytype *CodeAuthorityType) SetName(name string) {
	codeauthoritytype.Name = name
}

// Stage puts codecardinaldirectiontype to the model stage
func (codecardinaldirectiontype *CodeCardinalDirectionType) Stage(stage *Stage) *CodeCardinalDirectionType {
	__gong__stage(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionType_stagedOrder, stage.CodeCardinalDirectionType_orderStaged, &stage.CodeCardinalDirectionTypeOrder, stage.CodeCardinalDirectionTypes_mapString, codecardinaldirectiontype, codecardinaldirectiontype.Name)
	return codecardinaldirectiontype
}

// StagePreserveOrder puts codecardinaldirectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCardinalDirectionTypeOrder
// - update stage.CodeCardinalDirectionTypeOrder accordingly
func (codecardinaldirectiontype *CodeCardinalDirectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionType_stagedOrder, stage.CodeCardinalDirectionType_orderStaged, &stage.CodeCardinalDirectionTypeOrder, stage.CodeCardinalDirectionTypes_mapString, codecardinaldirectiontype, order, codecardinaldirectiontype.Name)
}

// Unstage removes codecardinaldirectiontype off the model stage
func (codecardinaldirectiontype *CodeCardinalDirectionType) Unstage(stage *Stage) *CodeCardinalDirectionType {
	__gong__unstage(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionTypes_mapString, codecardinaldirectiontype, codecardinaldirectiontype.Name)
	return codecardinaldirectiontype
}

// UnstageVoid removes codecardinaldirectiontype off the model stage
func (codecardinaldirectiontype *CodeCardinalDirectionType) UnstageVoid(stage *Stage) {
	codecardinaldirectiontype.Unstage(stage)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) StageVoid(stage *Stage) {
	codecardinaldirectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecardinaldirectiontype *CodeCardinalDirectionType) GetName() (res string) {
	return codecardinaldirectiontype.Name
}

// for satisfaction of GongStruct interface
func (codecardinaldirectiontype *CodeCardinalDirectionType) SetName(name string) {
	codecardinaldirectiontype.Name = name
}

// Stage puts codecommunicationmodetype to the model stage
func (codecommunicationmodetype *CodeCommunicationModeType) Stage(stage *Stage) *CodeCommunicationModeType {
	__gong__stage(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeType_orderStaged, &stage.CodeCommunicationModeTypeOrder, stage.CodeCommunicationModeTypes_mapString, codecommunicationmodetype, codecommunicationmodetype.Name)
	return codecommunicationmodetype
}

// StagePreserveOrder puts codecommunicationmodetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCommunicationModeTypeOrder
// - update stage.CodeCommunicationModeTypeOrder accordingly
func (codecommunicationmodetype *CodeCommunicationModeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeType_orderStaged, &stage.CodeCommunicationModeTypeOrder, stage.CodeCommunicationModeTypes_mapString, codecommunicationmodetype, order, codecommunicationmodetype.Name)
}

// Unstage removes codecommunicationmodetype off the model stage
func (codecommunicationmodetype *CodeCommunicationModeType) Unstage(stage *Stage) *CodeCommunicationModeType {
	__gong__unstage(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_mapString, codecommunicationmodetype, codecommunicationmodetype.Name)
	return codecommunicationmodetype
}

// UnstageVoid removes codecommunicationmodetype off the model stage
func (codecommunicationmodetype *CodeCommunicationModeType) UnstageVoid(stage *Stage) {
	codecommunicationmodetype.Unstage(stage)
}

func (codecommunicationmodetype *CodeCommunicationModeType) StageVoid(stage *Stage) {
	codecommunicationmodetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecommunicationmodetype *CodeCommunicationModeType) GetName() (res string) {
	return codecommunicationmodetype.Name
}

// for satisfaction of GongStruct interface
func (codecommunicationmodetype *CodeCommunicationModeType) SetName(name string) {
	codecommunicationmodetype.Name = name
}

// Stage puts codecomparisontype to the model stage
func (codecomparisontype *CodeComparisonType) Stage(stage *Stage) *CodeComparisonType {
	__gong__stage(stage.CodeComparisonTypes, stage.CodeComparisonType_stagedOrder, stage.CodeComparisonType_orderStaged, &stage.CodeComparisonTypeOrder, stage.CodeComparisonTypes_mapString, codecomparisontype, codecomparisontype.Name)
	return codecomparisontype
}

// StagePreserveOrder puts codecomparisontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeComparisonTypeOrder
// - update stage.CodeComparisonTypeOrder accordingly
func (codecomparisontype *CodeComparisonType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeComparisonTypes, stage.CodeComparisonType_stagedOrder, stage.CodeComparisonType_orderStaged, &stage.CodeComparisonTypeOrder, stage.CodeComparisonTypes_mapString, codecomparisontype, order, codecomparisontype.Name)
}

// Unstage removes codecomparisontype off the model stage
func (codecomparisontype *CodeComparisonType) Unstage(stage *Stage) *CodeComparisonType {
	__gong__unstage(stage.CodeComparisonTypes, stage.CodeComparisonTypes_mapString, codecomparisontype, codecomparisontype.Name)
	return codecomparisontype
}

// UnstageVoid removes codecomparisontype off the model stage
func (codecomparisontype *CodeComparisonType) UnstageVoid(stage *Stage) {
	codecomparisontype.Unstage(stage)
}

func (codecomparisontype *CodeComparisonType) StageVoid(stage *Stage) {
	codecomparisontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecomparisontype *CodeComparisonType) GetName() (res string) {
	return codecomparisontype.Name
}

// for satisfaction of GongStruct interface
func (codecomparisontype *CodeComparisonType) SetName(name string) {
	codecomparisontype.Name = name
}

// Stage puts codecoursetype to the model stage
func (codecoursetype *CodeCourseType) Stage(stage *Stage) *CodeCourseType {
	__gong__stage(stage.CodeCourseTypes, stage.CodeCourseType_stagedOrder, stage.CodeCourseType_orderStaged, &stage.CodeCourseTypeOrder, stage.CodeCourseTypes_mapString, codecoursetype, codecoursetype.Name)
	return codecoursetype
}

// StagePreserveOrder puts codecoursetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCourseTypeOrder
// - update stage.CodeCourseTypeOrder accordingly
func (codecoursetype *CodeCourseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCourseTypes, stage.CodeCourseType_stagedOrder, stage.CodeCourseType_orderStaged, &stage.CodeCourseTypeOrder, stage.CodeCourseTypes_mapString, codecoursetype, order, codecoursetype.Name)
}

// Unstage removes codecoursetype off the model stage
func (codecoursetype *CodeCourseType) Unstage(stage *Stage) *CodeCourseType {
	__gong__unstage(stage.CodeCourseTypes, stage.CodeCourseTypes_mapString, codecoursetype, codecoursetype.Name)
	return codecoursetype
}

// UnstageVoid removes codecoursetype off the model stage
func (codecoursetype *CodeCourseType) UnstageVoid(stage *Stage) {
	codecoursetype.Unstage(stage)
}

func (codecoursetype *CodeCourseType) StageVoid(stage *Stage) {
	codecoursetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecoursetype *CodeCourseType) GetName() (res string) {
	return codecoursetype.Name
}

// for satisfaction of GongStruct interface
func (codecoursetype *CodeCourseType) SetName(name string) {
	codecoursetype.Name = name
}

// Stage puts codedaytype to the model stage
func (codedaytype *CodeDayType) Stage(stage *Stage) *CodeDayType {
	__gong__stage(stage.CodeDayTypes, stage.CodeDayType_stagedOrder, stage.CodeDayType_orderStaged, &stage.CodeDayTypeOrder, stage.CodeDayTypes_mapString, codedaytype, codedaytype.Name)
	return codedaytype
}

// StagePreserveOrder puts codedaytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDayTypeOrder
// - update stage.CodeDayTypeOrder accordingly
func (codedaytype *CodeDayType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDayTypes, stage.CodeDayType_stagedOrder, stage.CodeDayType_orderStaged, &stage.CodeDayTypeOrder, stage.CodeDayTypes_mapString, codedaytype, order, codedaytype.Name)
}

// Unstage removes codedaytype off the model stage
func (codedaytype *CodeDayType) Unstage(stage *Stage) *CodeDayType {
	__gong__unstage(stage.CodeDayTypes, stage.CodeDayTypes_mapString, codedaytype, codedaytype.Name)
	return codedaytype
}

// UnstageVoid removes codedaytype off the model stage
func (codedaytype *CodeDayType) UnstageVoid(stage *Stage) {
	codedaytype.Unstage(stage)
}

func (codedaytype *CodeDayType) StageVoid(stage *Stage) {
	codedaytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedaytype *CodeDayType) GetName() (res string) {
	return codedaytype.Name
}

// for satisfaction of GongStruct interface
func (codedaytype *CodeDayType) SetName(name string) {
	codedaytype.Name = name
}

// Stage puts codedirectionturntype to the model stage
func (codedirectionturntype *CodeDirectionTurnType) Stage(stage *Stage) *CodeDirectionTurnType {
	__gong__stage(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnType_stagedOrder, stage.CodeDirectionTurnType_orderStaged, &stage.CodeDirectionTurnTypeOrder, stage.CodeDirectionTurnTypes_mapString, codedirectionturntype, codedirectionturntype.Name)
	return codedirectionturntype
}

// StagePreserveOrder puts codedirectionturntype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDirectionTurnTypeOrder
// - update stage.CodeDirectionTurnTypeOrder accordingly
func (codedirectionturntype *CodeDirectionTurnType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnType_stagedOrder, stage.CodeDirectionTurnType_orderStaged, &stage.CodeDirectionTurnTypeOrder, stage.CodeDirectionTurnTypes_mapString, codedirectionturntype, order, codedirectionturntype.Name)
}

// Unstage removes codedirectionturntype off the model stage
func (codedirectionturntype *CodeDirectionTurnType) Unstage(stage *Stage) *CodeDirectionTurnType {
	__gong__unstage(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnTypes_mapString, codedirectionturntype, codedirectionturntype.Name)
	return codedirectionturntype
}

// UnstageVoid removes codedirectionturntype off the model stage
func (codedirectionturntype *CodeDirectionTurnType) UnstageVoid(stage *Stage) {
	codedirectionturntype.Unstage(stage)
}

func (codedirectionturntype *CodeDirectionTurnType) StageVoid(stage *Stage) {
	codedirectionturntype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedirectionturntype *CodeDirectionTurnType) GetName() (res string) {
	return codedirectionturntype.Name
}

// for satisfaction of GongStruct interface
func (codedirectionturntype *CodeDirectionTurnType) SetName(name string) {
	codedirectionturntype.Name = name
}

// Stage puts codedirectiontype to the model stage
func (codedirectiontype *CodeDirectionType) Stage(stage *Stage) *CodeDirectionType {
	__gong__stage(stage.CodeDirectionTypes, stage.CodeDirectionType_stagedOrder, stage.CodeDirectionType_orderStaged, &stage.CodeDirectionTypeOrder, stage.CodeDirectionTypes_mapString, codedirectiontype, codedirectiontype.Name)
	return codedirectiontype
}

// StagePreserveOrder puts codedirectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDirectionTypeOrder
// - update stage.CodeDirectionTypeOrder accordingly
func (codedirectiontype *CodeDirectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDirectionTypes, stage.CodeDirectionType_stagedOrder, stage.CodeDirectionType_orderStaged, &stage.CodeDirectionTypeOrder, stage.CodeDirectionTypes_mapString, codedirectiontype, order, codedirectiontype.Name)
}

// Unstage removes codedirectiontype off the model stage
func (codedirectiontype *CodeDirectionType) Unstage(stage *Stage) *CodeDirectionType {
	__gong__unstage(stage.CodeDirectionTypes, stage.CodeDirectionTypes_mapString, codedirectiontype, codedirectiontype.Name)
	return codedirectiontype
}

// UnstageVoid removes codedirectiontype off the model stage
func (codedirectiontype *CodeDirectionType) UnstageVoid(stage *Stage) {
	codedirectiontype.Unstage(stage)
}

func (codedirectiontype *CodeDirectionType) StageVoid(stage *Stage) {
	codedirectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedirectiontype *CodeDirectionType) GetName() (res string) {
	return codedirectiontype.Name
}

// for satisfaction of GongStruct interface
func (codedirectiontype *CodeDirectionType) SetName(name string) {
	codedirectiontype.Name = name
}

// Stage puts codeequipmentanticollisiontype to the model stage
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) Stage(stage *Stage) *CodeEquipmentAntiCollisionType {
	__gong__stage(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionType_orderStaged, &stage.CodeEquipmentAntiCollisionTypeOrder, stage.CodeEquipmentAntiCollisionTypes_mapString, codeequipmentanticollisiontype, codeequipmentanticollisiontype.Name)
	return codeequipmentanticollisiontype
}

// StagePreserveOrder puts codeequipmentanticollisiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeEquipmentAntiCollisionTypeOrder
// - update stage.CodeEquipmentAntiCollisionTypeOrder accordingly
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionType_orderStaged, &stage.CodeEquipmentAntiCollisionTypeOrder, stage.CodeEquipmentAntiCollisionTypes_mapString, codeequipmentanticollisiontype, order, codeequipmentanticollisiontype.Name)
}

// Unstage removes codeequipmentanticollisiontype off the model stage
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) Unstage(stage *Stage) *CodeEquipmentAntiCollisionType {
	__gong__unstage(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_mapString, codeequipmentanticollisiontype, codeequipmentanticollisiontype.Name)
	return codeequipmentanticollisiontype
}

// UnstageVoid removes codeequipmentanticollisiontype off the model stage
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) UnstageVoid(stage *Stage) {
	codeequipmentanticollisiontype.Unstage(stage)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) StageVoid(stage *Stage) {
	codeequipmentanticollisiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GetName() (res string) {
	return codeequipmentanticollisiontype.Name
}

// for satisfaction of GongStruct interface
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) SetName(name string) {
	codeequipmentanticollisiontype.Name = name
}

// Stage puts codeflightorigintype to the model stage
func (codeflightorigintype *CodeFlightOriginType) Stage(stage *Stage) *CodeFlightOriginType {
	__gong__stage(stage.CodeFlightOriginTypes, stage.CodeFlightOriginType_stagedOrder, stage.CodeFlightOriginType_orderStaged, &stage.CodeFlightOriginTypeOrder, stage.CodeFlightOriginTypes_mapString, codeflightorigintype, codeflightorigintype.Name)
	return codeflightorigintype
}

// StagePreserveOrder puts codeflightorigintype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightOriginTypeOrder
// - update stage.CodeFlightOriginTypeOrder accordingly
func (codeflightorigintype *CodeFlightOriginType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightOriginTypes, stage.CodeFlightOriginType_stagedOrder, stage.CodeFlightOriginType_orderStaged, &stage.CodeFlightOriginTypeOrder, stage.CodeFlightOriginTypes_mapString, codeflightorigintype, order, codeflightorigintype.Name)
}

// Unstage removes codeflightorigintype off the model stage
func (codeflightorigintype *CodeFlightOriginType) Unstage(stage *Stage) *CodeFlightOriginType {
	__gong__unstage(stage.CodeFlightOriginTypes, stage.CodeFlightOriginTypes_mapString, codeflightorigintype, codeflightorigintype.Name)
	return codeflightorigintype
}

// UnstageVoid removes codeflightorigintype off the model stage
func (codeflightorigintype *CodeFlightOriginType) UnstageVoid(stage *Stage) {
	codeflightorigintype.Unstage(stage)
}

func (codeflightorigintype *CodeFlightOriginType) StageVoid(stage *Stage) {
	codeflightorigintype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightorigintype *CodeFlightOriginType) GetName() (res string) {
	return codeflightorigintype.Name
}

// for satisfaction of GongStruct interface
func (codeflightorigintype *CodeFlightOriginType) SetName(name string) {
	codeflightorigintype.Name = name
}

// Stage puts codeflightpurposetype to the model stage
func (codeflightpurposetype *CodeFlightPurposeType) Stage(stage *Stage) *CodeFlightPurposeType {
	__gong__stage(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeType_stagedOrder, stage.CodeFlightPurposeType_orderStaged, &stage.CodeFlightPurposeTypeOrder, stage.CodeFlightPurposeTypes_mapString, codeflightpurposetype, codeflightpurposetype.Name)
	return codeflightpurposetype
}

// StagePreserveOrder puts codeflightpurposetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightPurposeTypeOrder
// - update stage.CodeFlightPurposeTypeOrder accordingly
func (codeflightpurposetype *CodeFlightPurposeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeType_stagedOrder, stage.CodeFlightPurposeType_orderStaged, &stage.CodeFlightPurposeTypeOrder, stage.CodeFlightPurposeTypes_mapString, codeflightpurposetype, order, codeflightpurposetype.Name)
}

// Unstage removes codeflightpurposetype off the model stage
func (codeflightpurposetype *CodeFlightPurposeType) Unstage(stage *Stage) *CodeFlightPurposeType {
	__gong__unstage(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeTypes_mapString, codeflightpurposetype, codeflightpurposetype.Name)
	return codeflightpurposetype
}

// UnstageVoid removes codeflightpurposetype off the model stage
func (codeflightpurposetype *CodeFlightPurposeType) UnstageVoid(stage *Stage) {
	codeflightpurposetype.Unstage(stage)
}

func (codeflightpurposetype *CodeFlightPurposeType) StageVoid(stage *Stage) {
	codeflightpurposetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightpurposetype *CodeFlightPurposeType) GetName() (res string) {
	return codeflightpurposetype.Name
}

// for satisfaction of GongStruct interface
func (codeflightpurposetype *CodeFlightPurposeType) SetName(name string) {
	codeflightpurposetype.Name = name
}

// Stage puts codeflightrestrictiondesignatortype to the model stage
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) Stage(stage *Stage) *CodeFlightRestrictionDesignatorType {
	__gong__stage(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorType_stagedOrder, stage.CodeFlightRestrictionDesignatorType_orderStaged, &stage.CodeFlightRestrictionDesignatorTypeOrder, stage.CodeFlightRestrictionDesignatorTypes_mapString, codeflightrestrictiondesignatortype, codeflightrestrictiondesignatortype.Name)
	return codeflightrestrictiondesignatortype
}

// StagePreserveOrder puts codeflightrestrictiondesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightRestrictionDesignatorTypeOrder
// - update stage.CodeFlightRestrictionDesignatorTypeOrder accordingly
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorType_stagedOrder, stage.CodeFlightRestrictionDesignatorType_orderStaged, &stage.CodeFlightRestrictionDesignatorTypeOrder, stage.CodeFlightRestrictionDesignatorTypes_mapString, codeflightrestrictiondesignatortype, order, codeflightrestrictiondesignatortype.Name)
}

// Unstage removes codeflightrestrictiondesignatortype off the model stage
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) Unstage(stage *Stage) *CodeFlightRestrictionDesignatorType {
	__gong__unstage(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorTypes_mapString, codeflightrestrictiondesignatortype, codeflightrestrictiondesignatortype.Name)
	return codeflightrestrictiondesignatortype
}

// UnstageVoid removes codeflightrestrictiondesignatortype off the model stage
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) UnstageVoid(stage *Stage) {
	codeflightrestrictiondesignatortype.Unstage(stage)
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) StageVoid(stage *Stage) {
	codeflightrestrictiondesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GetName() (res string) {
	return codeflightrestrictiondesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) SetName(name string) {
	codeflightrestrictiondesignatortype.Name = name
}

// Stage puts codeflightrestrictiontype to the model stage
func (codeflightrestrictiontype *CodeFlightRestrictionType) Stage(stage *Stage) *CodeFlightRestrictionType {
	__gong__stage(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionType_stagedOrder, stage.CodeFlightRestrictionType_orderStaged, &stage.CodeFlightRestrictionTypeOrder, stage.CodeFlightRestrictionTypes_mapString, codeflightrestrictiontype, codeflightrestrictiontype.Name)
	return codeflightrestrictiontype
}

// StagePreserveOrder puts codeflightrestrictiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightRestrictionTypeOrder
// - update stage.CodeFlightRestrictionTypeOrder accordingly
func (codeflightrestrictiontype *CodeFlightRestrictionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionType_stagedOrder, stage.CodeFlightRestrictionType_orderStaged, &stage.CodeFlightRestrictionTypeOrder, stage.CodeFlightRestrictionTypes_mapString, codeflightrestrictiontype, order, codeflightrestrictiontype.Name)
}

// Unstage removes codeflightrestrictiontype off the model stage
func (codeflightrestrictiontype *CodeFlightRestrictionType) Unstage(stage *Stage) *CodeFlightRestrictionType {
	__gong__unstage(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionTypes_mapString, codeflightrestrictiontype, codeflightrestrictiontype.Name)
	return codeflightrestrictiontype
}

// UnstageVoid removes codeflightrestrictiontype off the model stage
func (codeflightrestrictiontype *CodeFlightRestrictionType) UnstageVoid(stage *Stage) {
	codeflightrestrictiontype.Unstage(stage)
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) StageVoid(stage *Stage) {
	codeflightrestrictiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightrestrictiontype *CodeFlightRestrictionType) GetName() (res string) {
	return codeflightrestrictiontype.Name
}

// for satisfaction of GongStruct interface
func (codeflightrestrictiontype *CodeFlightRestrictionType) SetName(name string) {
	codeflightrestrictiontype.Name = name
}

// Stage puts codeflightruletype to the model stage
func (codeflightruletype *CodeFlightRuleType) Stage(stage *Stage) *CodeFlightRuleType {
	__gong__stage(stage.CodeFlightRuleTypes, stage.CodeFlightRuleType_stagedOrder, stage.CodeFlightRuleType_orderStaged, &stage.CodeFlightRuleTypeOrder, stage.CodeFlightRuleTypes_mapString, codeflightruletype, codeflightruletype.Name)
	return codeflightruletype
}

// StagePreserveOrder puts codeflightruletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightRuleTypeOrder
// - update stage.CodeFlightRuleTypeOrder accordingly
func (codeflightruletype *CodeFlightRuleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightRuleTypes, stage.CodeFlightRuleType_stagedOrder, stage.CodeFlightRuleType_orderStaged, &stage.CodeFlightRuleTypeOrder, stage.CodeFlightRuleTypes_mapString, codeflightruletype, order, codeflightruletype.Name)
}

// Unstage removes codeflightruletype off the model stage
func (codeflightruletype *CodeFlightRuleType) Unstage(stage *Stage) *CodeFlightRuleType {
	__gong__unstage(stage.CodeFlightRuleTypes, stage.CodeFlightRuleTypes_mapString, codeflightruletype, codeflightruletype.Name)
	return codeflightruletype
}

// UnstageVoid removes codeflightruletype off the model stage
func (codeflightruletype *CodeFlightRuleType) UnstageVoid(stage *Stage) {
	codeflightruletype.Unstage(stage)
}

func (codeflightruletype *CodeFlightRuleType) StageVoid(stage *Stage) {
	codeflightruletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightruletype *CodeFlightRuleType) GetName() (res string) {
	return codeflightruletype.Name
}

// for satisfaction of GongStruct interface
func (codeflightruletype *CodeFlightRuleType) SetName(name string) {
	codeflightruletype.Name = name
}

// Stage puts codeflightstatustype to the model stage
func (codeflightstatustype *CodeFlightStatusType) Stage(stage *Stage) *CodeFlightStatusType {
	__gong__stage(stage.CodeFlightStatusTypes, stage.CodeFlightStatusType_stagedOrder, stage.CodeFlightStatusType_orderStaged, &stage.CodeFlightStatusTypeOrder, stage.CodeFlightStatusTypes_mapString, codeflightstatustype, codeflightstatustype.Name)
	return codeflightstatustype
}

// StagePreserveOrder puts codeflightstatustype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightStatusTypeOrder
// - update stage.CodeFlightStatusTypeOrder accordingly
func (codeflightstatustype *CodeFlightStatusType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightStatusTypes, stage.CodeFlightStatusType_stagedOrder, stage.CodeFlightStatusType_orderStaged, &stage.CodeFlightStatusTypeOrder, stage.CodeFlightStatusTypes_mapString, codeflightstatustype, order, codeflightstatustype.Name)
}

// Unstage removes codeflightstatustype off the model stage
func (codeflightstatustype *CodeFlightStatusType) Unstage(stage *Stage) *CodeFlightStatusType {
	__gong__unstage(stage.CodeFlightStatusTypes, stage.CodeFlightStatusTypes_mapString, codeflightstatustype, codeflightstatustype.Name)
	return codeflightstatustype
}

// UnstageVoid removes codeflightstatustype off the model stage
func (codeflightstatustype *CodeFlightStatusType) UnstageVoid(stage *Stage) {
	codeflightstatustype.Unstage(stage)
}

func (codeflightstatustype *CodeFlightStatusType) StageVoid(stage *Stage) {
	codeflightstatustype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightstatustype *CodeFlightStatusType) GetName() (res string) {
	return codeflightstatustype.Name
}

// for satisfaction of GongStruct interface
func (codeflightstatustype *CodeFlightStatusType) SetName(name string) {
	codeflightstatustype.Name = name
}

// Stage puts codeflighttype to the model stage
func (codeflighttype *CodeFlightType) Stage(stage *Stage) *CodeFlightType {
	__gong__stage(stage.CodeFlightTypes, stage.CodeFlightType_stagedOrder, stage.CodeFlightType_orderStaged, &stage.CodeFlightTypeOrder, stage.CodeFlightTypes_mapString, codeflighttype, codeflighttype.Name)
	return codeflighttype
}

// StagePreserveOrder puts codeflighttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightTypeOrder
// - update stage.CodeFlightTypeOrder accordingly
func (codeflighttype *CodeFlightType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightTypes, stage.CodeFlightType_stagedOrder, stage.CodeFlightType_orderStaged, &stage.CodeFlightTypeOrder, stage.CodeFlightTypes_mapString, codeflighttype, order, codeflighttype.Name)
}

// Unstage removes codeflighttype off the model stage
func (codeflighttype *CodeFlightType) Unstage(stage *Stage) *CodeFlightType {
	__gong__unstage(stage.CodeFlightTypes, stage.CodeFlightTypes_mapString, codeflighttype, codeflighttype.Name)
	return codeflighttype
}

// UnstageVoid removes codeflighttype off the model stage
func (codeflighttype *CodeFlightType) UnstageVoid(stage *Stage) {
	codeflighttype.Unstage(stage)
}

func (codeflighttype *CodeFlightType) StageVoid(stage *Stage) {
	codeflighttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflighttype *CodeFlightType) GetName() (res string) {
	return codeflighttype.Name
}

// for satisfaction of GongStruct interface
func (codeflighttype *CodeFlightType) SetName(name string) {
	codeflighttype.Name = name
}

// Stage puts codeflowconditionoperationtype to the model stage
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) Stage(stage *Stage) *CodeFlowConditionOperationType {
	__gong__stage(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationType_stagedOrder, stage.CodeFlowConditionOperationType_orderStaged, &stage.CodeFlowConditionOperationTypeOrder, stage.CodeFlowConditionOperationTypes_mapString, codeflowconditionoperationtype, codeflowconditionoperationtype.Name)
	return codeflowconditionoperationtype
}

// StagePreserveOrder puts codeflowconditionoperationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlowConditionOperationTypeOrder
// - update stage.CodeFlowConditionOperationTypeOrder accordingly
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationType_stagedOrder, stage.CodeFlowConditionOperationType_orderStaged, &stage.CodeFlowConditionOperationTypeOrder, stage.CodeFlowConditionOperationTypes_mapString, codeflowconditionoperationtype, order, codeflowconditionoperationtype.Name)
}

// Unstage removes codeflowconditionoperationtype off the model stage
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) Unstage(stage *Stage) *CodeFlowConditionOperationType {
	__gong__unstage(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationTypes_mapString, codeflowconditionoperationtype, codeflowconditionoperationtype.Name)
	return codeflowconditionoperationtype
}

// UnstageVoid removes codeflowconditionoperationtype off the model stage
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) UnstageVoid(stage *Stage) {
	codeflowconditionoperationtype.Unstage(stage)
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) StageVoid(stage *Stage) {
	codeflowconditionoperationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GetName() (res string) {
	return codeflowconditionoperationtype.Name
}

// for satisfaction of GongStruct interface
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) SetName(name string) {
	codeflowconditionoperationtype.Name = name
}

// Stage puts codefreeflighttype to the model stage
func (codefreeflighttype *CodeFreeFlightType) Stage(stage *Stage) *CodeFreeFlightType {
	__gong__stage(stage.CodeFreeFlightTypes, stage.CodeFreeFlightType_stagedOrder, stage.CodeFreeFlightType_orderStaged, &stage.CodeFreeFlightTypeOrder, stage.CodeFreeFlightTypes_mapString, codefreeflighttype, codefreeflighttype.Name)
	return codefreeflighttype
}

// StagePreserveOrder puts codefreeflighttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFreeFlightTypeOrder
// - update stage.CodeFreeFlightTypeOrder accordingly
func (codefreeflighttype *CodeFreeFlightType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFreeFlightTypes, stage.CodeFreeFlightType_stagedOrder, stage.CodeFreeFlightType_orderStaged, &stage.CodeFreeFlightTypeOrder, stage.CodeFreeFlightTypes_mapString, codefreeflighttype, order, codefreeflighttype.Name)
}

// Unstage removes codefreeflighttype off the model stage
func (codefreeflighttype *CodeFreeFlightType) Unstage(stage *Stage) *CodeFreeFlightType {
	__gong__unstage(stage.CodeFreeFlightTypes, stage.CodeFreeFlightTypes_mapString, codefreeflighttype, codefreeflighttype.Name)
	return codefreeflighttype
}

// UnstageVoid removes codefreeflighttype off the model stage
func (codefreeflighttype *CodeFreeFlightType) UnstageVoid(stage *Stage) {
	codefreeflighttype.Unstage(stage)
}

func (codefreeflighttype *CodeFreeFlightType) StageVoid(stage *Stage) {
	codefreeflighttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefreeflighttype *CodeFreeFlightType) GetName() (res string) {
	return codefreeflighttype.Name
}

// for satisfaction of GongStruct interface
func (codefreeflighttype *CodeFreeFlightType) SetName(name string) {
	codefreeflighttype.Name = name
}

// Stage puts codeleveltype to the model stage
func (codeleveltype *CodeLevelType) Stage(stage *Stage) *CodeLevelType {
	__gong__stage(stage.CodeLevelTypes, stage.CodeLevelType_stagedOrder, stage.CodeLevelType_orderStaged, &stage.CodeLevelTypeOrder, stage.CodeLevelTypes_mapString, codeleveltype, codeleveltype.Name)
	return codeleveltype
}

// StagePreserveOrder puts codeleveltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLevelTypeOrder
// - update stage.CodeLevelTypeOrder accordingly
func (codeleveltype *CodeLevelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLevelTypes, stage.CodeLevelType_stagedOrder, stage.CodeLevelType_orderStaged, &stage.CodeLevelTypeOrder, stage.CodeLevelTypes_mapString, codeleveltype, order, codeleveltype.Name)
}

// Unstage removes codeleveltype off the model stage
func (codeleveltype *CodeLevelType) Unstage(stage *Stage) *CodeLevelType {
	__gong__unstage(stage.CodeLevelTypes, stage.CodeLevelTypes_mapString, codeleveltype, codeleveltype.Name)
	return codeleveltype
}

// UnstageVoid removes codeleveltype off the model stage
func (codeleveltype *CodeLevelType) UnstageVoid(stage *Stage) {
	codeleveltype.Unstage(stage)
}

func (codeleveltype *CodeLevelType) StageVoid(stage *Stage) {
	codeleveltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeleveltype *CodeLevelType) GetName() (res string) {
	return codeleveltype.Name
}

// for satisfaction of GongStruct interface
func (codeleveltype *CodeLevelType) SetName(name string) {
	codeleveltype.Name = name
}

// Stage puts codelocationqualifiertype to the model stage
func (codelocationqualifiertype *CodeLocationQualifierType) Stage(stage *Stage) *CodeLocationQualifierType {
	__gong__stage(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierType_stagedOrder, stage.CodeLocationQualifierType_orderStaged, &stage.CodeLocationQualifierTypeOrder, stage.CodeLocationQualifierTypes_mapString, codelocationqualifiertype, codelocationqualifiertype.Name)
	return codelocationqualifiertype
}

// StagePreserveOrder puts codelocationqualifiertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLocationQualifierTypeOrder
// - update stage.CodeLocationQualifierTypeOrder accordingly
func (codelocationqualifiertype *CodeLocationQualifierType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierType_stagedOrder, stage.CodeLocationQualifierType_orderStaged, &stage.CodeLocationQualifierTypeOrder, stage.CodeLocationQualifierTypes_mapString, codelocationqualifiertype, order, codelocationqualifiertype.Name)
}

// Unstage removes codelocationqualifiertype off the model stage
func (codelocationqualifiertype *CodeLocationQualifierType) Unstage(stage *Stage) *CodeLocationQualifierType {
	__gong__unstage(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierTypes_mapString, codelocationqualifiertype, codelocationqualifiertype.Name)
	return codelocationqualifiertype
}

// UnstageVoid removes codelocationqualifiertype off the model stage
func (codelocationqualifiertype *CodeLocationQualifierType) UnstageVoid(stage *Stage) {
	codelocationqualifiertype.Unstage(stage)
}

func (codelocationqualifiertype *CodeLocationQualifierType) StageVoid(stage *Stage) {
	codelocationqualifiertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelocationqualifiertype *CodeLocationQualifierType) GetName() (res string) {
	return codelocationqualifiertype.Name
}

// for satisfaction of GongStruct interface
func (codelocationqualifiertype *CodeLocationQualifierType) SetName(name string) {
	codelocationqualifiertype.Name = name
}

// Stage puts codemeteoconditionstype to the model stage
func (codemeteoconditionstype *CodeMeteoConditionsType) Stage(stage *Stage) *CodeMeteoConditionsType {
	__gong__stage(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsType_stagedOrder, stage.CodeMeteoConditionsType_orderStaged, &stage.CodeMeteoConditionsTypeOrder, stage.CodeMeteoConditionsTypes_mapString, codemeteoconditionstype, codemeteoconditionstype.Name)
	return codemeteoconditionstype
}

// StagePreserveOrder puts codemeteoconditionstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMeteoConditionsTypeOrder
// - update stage.CodeMeteoConditionsTypeOrder accordingly
func (codemeteoconditionstype *CodeMeteoConditionsType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsType_stagedOrder, stage.CodeMeteoConditionsType_orderStaged, &stage.CodeMeteoConditionsTypeOrder, stage.CodeMeteoConditionsTypes_mapString, codemeteoconditionstype, order, codemeteoconditionstype.Name)
}

// Unstage removes codemeteoconditionstype off the model stage
func (codemeteoconditionstype *CodeMeteoConditionsType) Unstage(stage *Stage) *CodeMeteoConditionsType {
	__gong__unstage(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsTypes_mapString, codemeteoconditionstype, codemeteoconditionstype.Name)
	return codemeteoconditionstype
}

// UnstageVoid removes codemeteoconditionstype off the model stage
func (codemeteoconditionstype *CodeMeteoConditionsType) UnstageVoid(stage *Stage) {
	codemeteoconditionstype.Unstage(stage)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) StageVoid(stage *Stage) {
	codemeteoconditionstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemeteoconditionstype *CodeMeteoConditionsType) GetName() (res string) {
	return codemeteoconditionstype.Name
}

// for satisfaction of GongStruct interface
func (codemeteoconditionstype *CodeMeteoConditionsType) SetName(name string) {
	codemeteoconditionstype.Name = name
}

// Stage puts codemilitaryroutepointtype to the model stage
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) Stage(stage *Stage) *CodeMilitaryRoutePointType {
	__gong__stage(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointType_stagedOrder, stage.CodeMilitaryRoutePointType_orderStaged, &stage.CodeMilitaryRoutePointTypeOrder, stage.CodeMilitaryRoutePointTypes_mapString, codemilitaryroutepointtype, codemilitaryroutepointtype.Name)
	return codemilitaryroutepointtype
}

// StagePreserveOrder puts codemilitaryroutepointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMilitaryRoutePointTypeOrder
// - update stage.CodeMilitaryRoutePointTypeOrder accordingly
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointType_stagedOrder, stage.CodeMilitaryRoutePointType_orderStaged, &stage.CodeMilitaryRoutePointTypeOrder, stage.CodeMilitaryRoutePointTypes_mapString, codemilitaryroutepointtype, order, codemilitaryroutepointtype.Name)
}

// Unstage removes codemilitaryroutepointtype off the model stage
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) Unstage(stage *Stage) *CodeMilitaryRoutePointType {
	__gong__unstage(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointTypes_mapString, codemilitaryroutepointtype, codemilitaryroutepointtype.Name)
	return codemilitaryroutepointtype
}

// UnstageVoid removes codemilitaryroutepointtype off the model stage
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) UnstageVoid(stage *Stage) {
	codemilitaryroutepointtype.Unstage(stage)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) StageVoid(stage *Stage) {
	codemilitaryroutepointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GetName() (res string) {
	return codemilitaryroutepointtype.Name
}

// for satisfaction of GongStruct interface
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) SetName(name string) {
	codemilitaryroutepointtype.Name = name
}

// Stage puts codemilitarystatustype to the model stage
func (codemilitarystatustype *CodeMilitaryStatusType) Stage(stage *Stage) *CodeMilitaryStatusType {
	__gong__stage(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusType_stagedOrder, stage.CodeMilitaryStatusType_orderStaged, &stage.CodeMilitaryStatusTypeOrder, stage.CodeMilitaryStatusTypes_mapString, codemilitarystatustype, codemilitarystatustype.Name)
	return codemilitarystatustype
}

// StagePreserveOrder puts codemilitarystatustype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMilitaryStatusTypeOrder
// - update stage.CodeMilitaryStatusTypeOrder accordingly
func (codemilitarystatustype *CodeMilitaryStatusType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusType_stagedOrder, stage.CodeMilitaryStatusType_orderStaged, &stage.CodeMilitaryStatusTypeOrder, stage.CodeMilitaryStatusTypes_mapString, codemilitarystatustype, order, codemilitarystatustype.Name)
}

// Unstage removes codemilitarystatustype off the model stage
func (codemilitarystatustype *CodeMilitaryStatusType) Unstage(stage *Stage) *CodeMilitaryStatusType {
	__gong__unstage(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusTypes_mapString, codemilitarystatustype, codemilitarystatustype.Name)
	return codemilitarystatustype
}

// UnstageVoid removes codemilitarystatustype off the model stage
func (codemilitarystatustype *CodeMilitaryStatusType) UnstageVoid(stage *Stage) {
	codemilitarystatustype.Unstage(stage)
}

func (codemilitarystatustype *CodeMilitaryStatusType) StageVoid(stage *Stage) {
	codemilitarystatustype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemilitarystatustype *CodeMilitaryStatusType) GetName() (res string) {
	return codemilitarystatustype.Name
}

// for satisfaction of GongStruct interface
func (codemilitarystatustype *CodeMilitaryStatusType) SetName(name string) {
	codemilitarystatustype.Name = name
}

// Stage puts codemilitarytrainingtype to the model stage
func (codemilitarytrainingtype *CodeMilitaryTrainingType) Stage(stage *Stage) *CodeMilitaryTrainingType {
	__gong__stage(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingType_stagedOrder, stage.CodeMilitaryTrainingType_orderStaged, &stage.CodeMilitaryTrainingTypeOrder, stage.CodeMilitaryTrainingTypes_mapString, codemilitarytrainingtype, codemilitarytrainingtype.Name)
	return codemilitarytrainingtype
}

// StagePreserveOrder puts codemilitarytrainingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMilitaryTrainingTypeOrder
// - update stage.CodeMilitaryTrainingTypeOrder accordingly
func (codemilitarytrainingtype *CodeMilitaryTrainingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingType_stagedOrder, stage.CodeMilitaryTrainingType_orderStaged, &stage.CodeMilitaryTrainingTypeOrder, stage.CodeMilitaryTrainingTypes_mapString, codemilitarytrainingtype, order, codemilitarytrainingtype.Name)
}

// Unstage removes codemilitarytrainingtype off the model stage
func (codemilitarytrainingtype *CodeMilitaryTrainingType) Unstage(stage *Stage) *CodeMilitaryTrainingType {
	__gong__unstage(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingTypes_mapString, codemilitarytrainingtype, codemilitarytrainingtype.Name)
	return codemilitarytrainingtype
}

// UnstageVoid removes codemilitarytrainingtype off the model stage
func (codemilitarytrainingtype *CodeMilitaryTrainingType) UnstageVoid(stage *Stage) {
	codemilitarytrainingtype.Unstage(stage)
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) StageVoid(stage *Stage) {
	codemilitarytrainingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemilitarytrainingtype *CodeMilitaryTrainingType) GetName() (res string) {
	return codemilitarytrainingtype.Name
}

// for satisfaction of GongStruct interface
func (codemilitarytrainingtype *CodeMilitaryTrainingType) SetName(name string) {
	codemilitarytrainingtype.Name = name
}

// Stage puts codenavigationequipmenttype to the model stage
func (codenavigationequipmenttype *CodeNavigationEquipmentType) Stage(stage *Stage) *CodeNavigationEquipmentType {
	__gong__stage(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentType_orderStaged, &stage.CodeNavigationEquipmentTypeOrder, stage.CodeNavigationEquipmentTypes_mapString, codenavigationequipmenttype, codenavigationequipmenttype.Name)
	return codenavigationequipmenttype
}

// StagePreserveOrder puts codenavigationequipmenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavigationEquipmentTypeOrder
// - update stage.CodeNavigationEquipmentTypeOrder accordingly
func (codenavigationequipmenttype *CodeNavigationEquipmentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentType_orderStaged, &stage.CodeNavigationEquipmentTypeOrder, stage.CodeNavigationEquipmentTypes_mapString, codenavigationequipmenttype, order, codenavigationequipmenttype.Name)
}

// Unstage removes codenavigationequipmenttype off the model stage
func (codenavigationequipmenttype *CodeNavigationEquipmentType) Unstage(stage *Stage) *CodeNavigationEquipmentType {
	__gong__unstage(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_mapString, codenavigationequipmenttype, codenavigationequipmenttype.Name)
	return codenavigationequipmenttype
}

// UnstageVoid removes codenavigationequipmenttype off the model stage
func (codenavigationequipmenttype *CodeNavigationEquipmentType) UnstageVoid(stage *Stage) {
	codenavigationequipmenttype.Unstage(stage)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) StageVoid(stage *Stage) {
	codenavigationequipmenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavigationequipmenttype *CodeNavigationEquipmentType) GetName() (res string) {
	return codenavigationequipmenttype.Name
}

// for satisfaction of GongStruct interface
func (codenavigationequipmenttype *CodeNavigationEquipmentType) SetName(name string) {
	codenavigationequipmenttype.Name = name
}

// Stage puts codenavigationspecificationtype to the model stage
func (codenavigationspecificationtype *CodeNavigationSpecificationType) Stage(stage *Stage) *CodeNavigationSpecificationType {
	__gong__stage(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationType_orderStaged, &stage.CodeNavigationSpecificationTypeOrder, stage.CodeNavigationSpecificationTypes_mapString, codenavigationspecificationtype, codenavigationspecificationtype.Name)
	return codenavigationspecificationtype
}

// StagePreserveOrder puts codenavigationspecificationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavigationSpecificationTypeOrder
// - update stage.CodeNavigationSpecificationTypeOrder accordingly
func (codenavigationspecificationtype *CodeNavigationSpecificationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationType_orderStaged, &stage.CodeNavigationSpecificationTypeOrder, stage.CodeNavigationSpecificationTypes_mapString, codenavigationspecificationtype, order, codenavigationspecificationtype.Name)
}

// Unstage removes codenavigationspecificationtype off the model stage
func (codenavigationspecificationtype *CodeNavigationSpecificationType) Unstage(stage *Stage) *CodeNavigationSpecificationType {
	__gong__unstage(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_mapString, codenavigationspecificationtype, codenavigationspecificationtype.Name)
	return codenavigationspecificationtype
}

// UnstageVoid removes codenavigationspecificationtype off the model stage
func (codenavigationspecificationtype *CodeNavigationSpecificationType) UnstageVoid(stage *Stage) {
	codenavigationspecificationtype.Unstage(stage)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) StageVoid(stage *Stage) {
	codenavigationspecificationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavigationspecificationtype *CodeNavigationSpecificationType) GetName() (res string) {
	return codenavigationspecificationtype.Name
}

// for satisfaction of GongStruct interface
func (codenavigationspecificationtype *CodeNavigationSpecificationType) SetName(name string) {
	codenavigationspecificationtype.Name = name
}

// Stage puts codenotepurposetype to the model stage
func (codenotepurposetype *CodeNotePurposeType) Stage(stage *Stage) *CodeNotePurposeType {
	__gong__stage(stage.CodeNotePurposeTypes, stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeType_orderStaged, &stage.CodeNotePurposeTypeOrder, stage.CodeNotePurposeTypes_mapString, codenotepurposetype, codenotepurposetype.Name)
	return codenotepurposetype
}

// StagePreserveOrder puts codenotepurposetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNotePurposeTypeOrder
// - update stage.CodeNotePurposeTypeOrder accordingly
func (codenotepurposetype *CodeNotePurposeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNotePurposeTypes, stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeType_orderStaged, &stage.CodeNotePurposeTypeOrder, stage.CodeNotePurposeTypes_mapString, codenotepurposetype, order, codenotepurposetype.Name)
}

// Unstage removes codenotepurposetype off the model stage
func (codenotepurposetype *CodeNotePurposeType) Unstage(stage *Stage) *CodeNotePurposeType {
	__gong__unstage(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_mapString, codenotepurposetype, codenotepurposetype.Name)
	return codenotepurposetype
}

// UnstageVoid removes codenotepurposetype off the model stage
func (codenotepurposetype *CodeNotePurposeType) UnstageVoid(stage *Stage) {
	codenotepurposetype.Unstage(stage)
}

func (codenotepurposetype *CodeNotePurposeType) StageVoid(stage *Stage) {
	codenotepurposetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenotepurposetype *CodeNotePurposeType) GetName() (res string) {
	return codenotepurposetype.Name
}

// for satisfaction of GongStruct interface
func (codenotepurposetype *CodeNotePurposeType) SetName(name string) {
	codenotepurposetype.Name = name
}

// Stage puts codeobstacleassessmentsurfacetype to the model stage
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) Stage(stage *Stage) *CodeObstacleAssessmentSurfaceType {
	__gong__stage(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceType_orderStaged, &stage.CodeObstacleAssessmentSurfaceTypeOrder, stage.CodeObstacleAssessmentSurfaceTypes_mapString, codeobstacleassessmentsurfacetype, codeobstacleassessmentsurfacetype.Name)
	return codeobstacleassessmentsurfacetype
}

// StagePreserveOrder puts codeobstacleassessmentsurfacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeObstacleAssessmentSurfaceTypeOrder
// - update stage.CodeObstacleAssessmentSurfaceTypeOrder accordingly
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceType_orderStaged, &stage.CodeObstacleAssessmentSurfaceTypeOrder, stage.CodeObstacleAssessmentSurfaceTypes_mapString, codeobstacleassessmentsurfacetype, order, codeobstacleassessmentsurfacetype.Name)
}

// Unstage removes codeobstacleassessmentsurfacetype off the model stage
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) Unstage(stage *Stage) *CodeObstacleAssessmentSurfaceType {
	__gong__unstage(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceTypes_mapString, codeobstacleassessmentsurfacetype, codeobstacleassessmentsurfacetype.Name)
	return codeobstacleassessmentsurfacetype
}

// UnstageVoid removes codeobstacleassessmentsurfacetype off the model stage
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) UnstageVoid(stage *Stage) {
	codeobstacleassessmentsurfacetype.Unstage(stage)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) StageVoid(stage *Stage) {
	codeobstacleassessmentsurfacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GetName() (res string) {
	return codeobstacleassessmentsurfacetype.Name
}

// for satisfaction of GongStruct interface
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) SetName(name string) {
	codeobstacleassessmentsurfacetype.Name = name
}

// Stage puts codeobstructionidsurfacezonetype to the model stage
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) Stage(stage *Stage) *CodeObstructionIdSurfaceZoneType {
	__gong__stage(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneType_stagedOrder, stage.CodeObstructionIdSurfaceZoneType_orderStaged, &stage.CodeObstructionIdSurfaceZoneTypeOrder, stage.CodeObstructionIdSurfaceZoneTypes_mapString, codeobstructionidsurfacezonetype, codeobstructionidsurfacezonetype.Name)
	return codeobstructionidsurfacezonetype
}

// StagePreserveOrder puts codeobstructionidsurfacezonetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeObstructionIdSurfaceZoneTypeOrder
// - update stage.CodeObstructionIdSurfaceZoneTypeOrder accordingly
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneType_stagedOrder, stage.CodeObstructionIdSurfaceZoneType_orderStaged, &stage.CodeObstructionIdSurfaceZoneTypeOrder, stage.CodeObstructionIdSurfaceZoneTypes_mapString, codeobstructionidsurfacezonetype, order, codeobstructionidsurfacezonetype.Name)
}

// Unstage removes codeobstructionidsurfacezonetype off the model stage
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) Unstage(stage *Stage) *CodeObstructionIdSurfaceZoneType {
	__gong__unstage(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneTypes_mapString, codeobstructionidsurfacezonetype, codeobstructionidsurfacezonetype.Name)
	return codeobstructionidsurfacezonetype
}

// UnstageVoid removes codeobstructionidsurfacezonetype off the model stage
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) UnstageVoid(stage *Stage) {
	codeobstructionidsurfacezonetype.Unstage(stage)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) StageVoid(stage *Stage) {
	codeobstructionidsurfacezonetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GetName() (res string) {
	return codeobstructionidsurfacezonetype.Name
}

// for satisfaction of GongStruct interface
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) SetName(name string) {
	codeobstructionidsurfacezonetype.Name = name
}

// Stage puts codernptype to the model stage
func (codernptype *CodeRNPType) Stage(stage *Stage) *CodeRNPType {
	__gong__stage(stage.CodeRNPTypes, stage.CodeRNPType_stagedOrder, stage.CodeRNPType_orderStaged, &stage.CodeRNPTypeOrder, stage.CodeRNPTypes_mapString, codernptype, codernptype.Name)
	return codernptype
}

// StagePreserveOrder puts codernptype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRNPTypeOrder
// - update stage.CodeRNPTypeOrder accordingly
func (codernptype *CodeRNPType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRNPTypes, stage.CodeRNPType_stagedOrder, stage.CodeRNPType_orderStaged, &stage.CodeRNPTypeOrder, stage.CodeRNPTypes_mapString, codernptype, order, codernptype.Name)
}

// Unstage removes codernptype off the model stage
func (codernptype *CodeRNPType) Unstage(stage *Stage) *CodeRNPType {
	__gong__unstage(stage.CodeRNPTypes, stage.CodeRNPTypes_mapString, codernptype, codernptype.Name)
	return codernptype
}

// UnstageVoid removes codernptype off the model stage
func (codernptype *CodeRNPType) UnstageVoid(stage *Stage) {
	codernptype.Unstage(stage)
}

func (codernptype *CodeRNPType) StageVoid(stage *Stage) {
	codernptype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codernptype *CodeRNPType) GetName() (res string) {
	return codernptype.Name
}

// for satisfaction of GongStruct interface
func (codernptype *CodeRNPType) SetName(name string) {
	codernptype.Name = name
}

// Stage puts codervsmpointroletype to the model stage
func (codervsmpointroletype *CodeRVSMPointRoleType) Stage(stage *Stage) *CodeRVSMPointRoleType {
	__gong__stage(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleType_stagedOrder, stage.CodeRVSMPointRoleType_orderStaged, &stage.CodeRVSMPointRoleTypeOrder, stage.CodeRVSMPointRoleTypes_mapString, codervsmpointroletype, codervsmpointroletype.Name)
	return codervsmpointroletype
}

// StagePreserveOrder puts codervsmpointroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRVSMPointRoleTypeOrder
// - update stage.CodeRVSMPointRoleTypeOrder accordingly
func (codervsmpointroletype *CodeRVSMPointRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleType_stagedOrder, stage.CodeRVSMPointRoleType_orderStaged, &stage.CodeRVSMPointRoleTypeOrder, stage.CodeRVSMPointRoleTypes_mapString, codervsmpointroletype, order, codervsmpointroletype.Name)
}

// Unstage removes codervsmpointroletype off the model stage
func (codervsmpointroletype *CodeRVSMPointRoleType) Unstage(stage *Stage) *CodeRVSMPointRoleType {
	__gong__unstage(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleTypes_mapString, codervsmpointroletype, codervsmpointroletype.Name)
	return codervsmpointroletype
}

// UnstageVoid removes codervsmpointroletype off the model stage
func (codervsmpointroletype *CodeRVSMPointRoleType) UnstageVoid(stage *Stage) {
	codervsmpointroletype.Unstage(stage)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) StageVoid(stage *Stage) {
	codervsmpointroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codervsmpointroletype *CodeRVSMPointRoleType) GetName() (res string) {
	return codervsmpointroletype.Name
}

// for satisfaction of GongStruct interface
func (codervsmpointroletype *CodeRVSMPointRoleType) SetName(name string) {
	codervsmpointroletype.Name = name
}

// Stage puts codervsmtype to the model stage
func (codervsmtype *CodeRVSMType) Stage(stage *Stage) *CodeRVSMType {
	__gong__stage(stage.CodeRVSMTypes, stage.CodeRVSMType_stagedOrder, stage.CodeRVSMType_orderStaged, &stage.CodeRVSMTypeOrder, stage.CodeRVSMTypes_mapString, codervsmtype, codervsmtype.Name)
	return codervsmtype
}

// StagePreserveOrder puts codervsmtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRVSMTypeOrder
// - update stage.CodeRVSMTypeOrder accordingly
func (codervsmtype *CodeRVSMType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRVSMTypes, stage.CodeRVSMType_stagedOrder, stage.CodeRVSMType_orderStaged, &stage.CodeRVSMTypeOrder, stage.CodeRVSMTypes_mapString, codervsmtype, order, codervsmtype.Name)
}

// Unstage removes codervsmtype off the model stage
func (codervsmtype *CodeRVSMType) Unstage(stage *Stage) *CodeRVSMType {
	__gong__unstage(stage.CodeRVSMTypes, stage.CodeRVSMTypes_mapString, codervsmtype, codervsmtype.Name)
	return codervsmtype
}

// UnstageVoid removes codervsmtype off the model stage
func (codervsmtype *CodeRVSMType) UnstageVoid(stage *Stage) {
	codervsmtype.Unstage(stage)
}

func (codervsmtype *CodeRVSMType) StageVoid(stage *Stage) {
	codervsmtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codervsmtype *CodeRVSMType) GetName() (res string) {
	return codervsmtype.Name
}

// for satisfaction of GongStruct interface
func (codervsmtype *CodeRVSMType) SetName(name string) {
	codervsmtype.Name = name
}

// Stage puts codereferenceroletype to the model stage
func (codereferenceroletype *CodeReferenceRoleType) Stage(stage *Stage) *CodeReferenceRoleType {
	__gong__stage(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleType_stagedOrder, stage.CodeReferenceRoleType_orderStaged, &stage.CodeReferenceRoleTypeOrder, stage.CodeReferenceRoleTypes_mapString, codereferenceroletype, codereferenceroletype.Name)
	return codereferenceroletype
}

// StagePreserveOrder puts codereferenceroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeReferenceRoleTypeOrder
// - update stage.CodeReferenceRoleTypeOrder accordingly
func (codereferenceroletype *CodeReferenceRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleType_stagedOrder, stage.CodeReferenceRoleType_orderStaged, &stage.CodeReferenceRoleTypeOrder, stage.CodeReferenceRoleTypes_mapString, codereferenceroletype, order, codereferenceroletype.Name)
}

// Unstage removes codereferenceroletype off the model stage
func (codereferenceroletype *CodeReferenceRoleType) Unstage(stage *Stage) *CodeReferenceRoleType {
	__gong__unstage(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleTypes_mapString, codereferenceroletype, codereferenceroletype.Name)
	return codereferenceroletype
}

// UnstageVoid removes codereferenceroletype off the model stage
func (codereferenceroletype *CodeReferenceRoleType) UnstageVoid(stage *Stage) {
	codereferenceroletype.Unstage(stage)
}

func (codereferenceroletype *CodeReferenceRoleType) StageVoid(stage *Stage) {
	codereferenceroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codereferenceroletype *CodeReferenceRoleType) GetName() (res string) {
	return codereferenceroletype.Name
}

// for satisfaction of GongStruct interface
func (codereferenceroletype *CodeReferenceRoleType) SetName(name string) {
	codereferenceroletype.Name = name
}

// Stage puts coderouteavailabilitytype to the model stage
func (coderouteavailabilitytype *CodeRouteAvailabilityType) Stage(stage *Stage) *CodeRouteAvailabilityType {
	__gong__stage(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityType_stagedOrder, stage.CodeRouteAvailabilityType_orderStaged, &stage.CodeRouteAvailabilityTypeOrder, stage.CodeRouteAvailabilityTypes_mapString, coderouteavailabilitytype, coderouteavailabilitytype.Name)
	return coderouteavailabilitytype
}

// StagePreserveOrder puts coderouteavailabilitytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteAvailabilityTypeOrder
// - update stage.CodeRouteAvailabilityTypeOrder accordingly
func (coderouteavailabilitytype *CodeRouteAvailabilityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityType_stagedOrder, stage.CodeRouteAvailabilityType_orderStaged, &stage.CodeRouteAvailabilityTypeOrder, stage.CodeRouteAvailabilityTypes_mapString, coderouteavailabilitytype, order, coderouteavailabilitytype.Name)
}

// Unstage removes coderouteavailabilitytype off the model stage
func (coderouteavailabilitytype *CodeRouteAvailabilityType) Unstage(stage *Stage) *CodeRouteAvailabilityType {
	__gong__unstage(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityTypes_mapString, coderouteavailabilitytype, coderouteavailabilitytype.Name)
	return coderouteavailabilitytype
}

// UnstageVoid removes coderouteavailabilitytype off the model stage
func (coderouteavailabilitytype *CodeRouteAvailabilityType) UnstageVoid(stage *Stage) {
	coderouteavailabilitytype.Unstage(stage)
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) StageVoid(stage *Stage) {
	coderouteavailabilitytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderouteavailabilitytype *CodeRouteAvailabilityType) GetName() (res string) {
	return coderouteavailabilitytype.Name
}

// for satisfaction of GongStruct interface
func (coderouteavailabilitytype *CodeRouteAvailabilityType) SetName(name string) {
	coderouteavailabilitytype.Name = name
}

// Stage puts coderoutedesignatorlettertype to the model stage
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) Stage(stage *Stage) *CodeRouteDesignatorLetterType {
	__gong__stage(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterType_stagedOrder, stage.CodeRouteDesignatorLetterType_orderStaged, &stage.CodeRouteDesignatorLetterTypeOrder, stage.CodeRouteDesignatorLetterTypes_mapString, coderoutedesignatorlettertype, coderoutedesignatorlettertype.Name)
	return coderoutedesignatorlettertype
}

// StagePreserveOrder puts coderoutedesignatorlettertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteDesignatorLetterTypeOrder
// - update stage.CodeRouteDesignatorLetterTypeOrder accordingly
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterType_stagedOrder, stage.CodeRouteDesignatorLetterType_orderStaged, &stage.CodeRouteDesignatorLetterTypeOrder, stage.CodeRouteDesignatorLetterTypes_mapString, coderoutedesignatorlettertype, order, coderoutedesignatorlettertype.Name)
}

// Unstage removes coderoutedesignatorlettertype off the model stage
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) Unstage(stage *Stage) *CodeRouteDesignatorLetterType {
	__gong__unstage(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterTypes_mapString, coderoutedesignatorlettertype, coderoutedesignatorlettertype.Name)
	return coderoutedesignatorlettertype
}

// UnstageVoid removes coderoutedesignatorlettertype off the model stage
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) UnstageVoid(stage *Stage) {
	coderoutedesignatorlettertype.Unstage(stage)
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) StageVoid(stage *Stage) {
	coderoutedesignatorlettertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GetName() (res string) {
	return coderoutedesignatorlettertype.Name
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) SetName(name string) {
	coderoutedesignatorlettertype.Name = name
}

// Stage puts coderoutedesignatorprefixtype to the model stage
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) Stage(stage *Stage) *CodeRouteDesignatorPrefixType {
	__gong__stage(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixType_stagedOrder, stage.CodeRouteDesignatorPrefixType_orderStaged, &stage.CodeRouteDesignatorPrefixTypeOrder, stage.CodeRouteDesignatorPrefixTypes_mapString, coderoutedesignatorprefixtype, coderoutedesignatorprefixtype.Name)
	return coderoutedesignatorprefixtype
}

// StagePreserveOrder puts coderoutedesignatorprefixtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteDesignatorPrefixTypeOrder
// - update stage.CodeRouteDesignatorPrefixTypeOrder accordingly
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixType_stagedOrder, stage.CodeRouteDesignatorPrefixType_orderStaged, &stage.CodeRouteDesignatorPrefixTypeOrder, stage.CodeRouteDesignatorPrefixTypes_mapString, coderoutedesignatorprefixtype, order, coderoutedesignatorprefixtype.Name)
}

// Unstage removes coderoutedesignatorprefixtype off the model stage
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) Unstage(stage *Stage) *CodeRouteDesignatorPrefixType {
	__gong__unstage(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixTypes_mapString, coderoutedesignatorprefixtype, coderoutedesignatorprefixtype.Name)
	return coderoutedesignatorprefixtype
}

// UnstageVoid removes coderoutedesignatorprefixtype off the model stage
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) UnstageVoid(stage *Stage) {
	coderoutedesignatorprefixtype.Unstage(stage)
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) StageVoid(stage *Stage) {
	coderoutedesignatorprefixtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GetName() (res string) {
	return coderoutedesignatorprefixtype.Name
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) SetName(name string) {
	coderoutedesignatorprefixtype.Name = name
}

// Stage puts coderoutedesignatorsuffixtype to the model stage
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) Stage(stage *Stage) *CodeRouteDesignatorSuffixType {
	__gong__stage(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixType_stagedOrder, stage.CodeRouteDesignatorSuffixType_orderStaged, &stage.CodeRouteDesignatorSuffixTypeOrder, stage.CodeRouteDesignatorSuffixTypes_mapString, coderoutedesignatorsuffixtype, coderoutedesignatorsuffixtype.Name)
	return coderoutedesignatorsuffixtype
}

// StagePreserveOrder puts coderoutedesignatorsuffixtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteDesignatorSuffixTypeOrder
// - update stage.CodeRouteDesignatorSuffixTypeOrder accordingly
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixType_stagedOrder, stage.CodeRouteDesignatorSuffixType_orderStaged, &stage.CodeRouteDesignatorSuffixTypeOrder, stage.CodeRouteDesignatorSuffixTypes_mapString, coderoutedesignatorsuffixtype, order, coderoutedesignatorsuffixtype.Name)
}

// Unstage removes coderoutedesignatorsuffixtype off the model stage
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) Unstage(stage *Stage) *CodeRouteDesignatorSuffixType {
	__gong__unstage(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixTypes_mapString, coderoutedesignatorsuffixtype, coderoutedesignatorsuffixtype.Name)
	return coderoutedesignatorsuffixtype
}

// UnstageVoid removes coderoutedesignatorsuffixtype off the model stage
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) UnstageVoid(stage *Stage) {
	coderoutedesignatorsuffixtype.Unstage(stage)
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) StageVoid(stage *Stage) {
	coderoutedesignatorsuffixtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GetName() (res string) {
	return coderoutedesignatorsuffixtype.Name
}

// for satisfaction of GongStruct interface
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) SetName(name string) {
	coderoutedesignatorsuffixtype.Name = name
}

// Stage puts coderoutenavigationtype to the model stage
func (coderoutenavigationtype *CodeRouteNavigationType) Stage(stage *Stage) *CodeRouteNavigationType {
	__gong__stage(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationType_stagedOrder, stage.CodeRouteNavigationType_orderStaged, &stage.CodeRouteNavigationTypeOrder, stage.CodeRouteNavigationTypes_mapString, coderoutenavigationtype, coderoutenavigationtype.Name)
	return coderoutenavigationtype
}

// StagePreserveOrder puts coderoutenavigationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteNavigationTypeOrder
// - update stage.CodeRouteNavigationTypeOrder accordingly
func (coderoutenavigationtype *CodeRouteNavigationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationType_stagedOrder, stage.CodeRouteNavigationType_orderStaged, &stage.CodeRouteNavigationTypeOrder, stage.CodeRouteNavigationTypes_mapString, coderoutenavigationtype, order, coderoutenavigationtype.Name)
}

// Unstage removes coderoutenavigationtype off the model stage
func (coderoutenavigationtype *CodeRouteNavigationType) Unstage(stage *Stage) *CodeRouteNavigationType {
	__gong__unstage(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationTypes_mapString, coderoutenavigationtype, coderoutenavigationtype.Name)
	return coderoutenavigationtype
}

// UnstageVoid removes coderoutenavigationtype off the model stage
func (coderoutenavigationtype *CodeRouteNavigationType) UnstageVoid(stage *Stage) {
	coderoutenavigationtype.Unstage(stage)
}

func (coderoutenavigationtype *CodeRouteNavigationType) StageVoid(stage *Stage) {
	coderoutenavigationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutenavigationtype *CodeRouteNavigationType) GetName() (res string) {
	return coderoutenavigationtype.Name
}

// for satisfaction of GongStruct interface
func (coderoutenavigationtype *CodeRouteNavigationType) SetName(name string) {
	coderoutenavigationtype.Name = name
}

// Stage puts coderouteorigintype to the model stage
func (coderouteorigintype *CodeRouteOriginType) Stage(stage *Stage) *CodeRouteOriginType {
	__gong__stage(stage.CodeRouteOriginTypes, stage.CodeRouteOriginType_stagedOrder, stage.CodeRouteOriginType_orderStaged, &stage.CodeRouteOriginTypeOrder, stage.CodeRouteOriginTypes_mapString, coderouteorigintype, coderouteorigintype.Name)
	return coderouteorigintype
}

// StagePreserveOrder puts coderouteorigintype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteOriginTypeOrder
// - update stage.CodeRouteOriginTypeOrder accordingly
func (coderouteorigintype *CodeRouteOriginType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteOriginTypes, stage.CodeRouteOriginType_stagedOrder, stage.CodeRouteOriginType_orderStaged, &stage.CodeRouteOriginTypeOrder, stage.CodeRouteOriginTypes_mapString, coderouteorigintype, order, coderouteorigintype.Name)
}

// Unstage removes coderouteorigintype off the model stage
func (coderouteorigintype *CodeRouteOriginType) Unstage(stage *Stage) *CodeRouteOriginType {
	__gong__unstage(stage.CodeRouteOriginTypes, stage.CodeRouteOriginTypes_mapString, coderouteorigintype, coderouteorigintype.Name)
	return coderouteorigintype
}

// UnstageVoid removes coderouteorigintype off the model stage
func (coderouteorigintype *CodeRouteOriginType) UnstageVoid(stage *Stage) {
	coderouteorigintype.Unstage(stage)
}

func (coderouteorigintype *CodeRouteOriginType) StageVoid(stage *Stage) {
	coderouteorigintype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderouteorigintype *CodeRouteOriginType) GetName() (res string) {
	return coderouteorigintype.Name
}

// for satisfaction of GongStruct interface
func (coderouteorigintype *CodeRouteOriginType) SetName(name string) {
	coderouteorigintype.Name = name
}

// Stage puts coderoutesegmentpathtype to the model stage
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) Stage(stage *Stage) *CodeRouteSegmentPathType {
	__gong__stage(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathType_stagedOrder, stage.CodeRouteSegmentPathType_orderStaged, &stage.CodeRouteSegmentPathTypeOrder, stage.CodeRouteSegmentPathTypes_mapString, coderoutesegmentpathtype, coderoutesegmentpathtype.Name)
	return coderoutesegmentpathtype
}

// StagePreserveOrder puts coderoutesegmentpathtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteSegmentPathTypeOrder
// - update stage.CodeRouteSegmentPathTypeOrder accordingly
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathType_stagedOrder, stage.CodeRouteSegmentPathType_orderStaged, &stage.CodeRouteSegmentPathTypeOrder, stage.CodeRouteSegmentPathTypes_mapString, coderoutesegmentpathtype, order, coderoutesegmentpathtype.Name)
}

// Unstage removes coderoutesegmentpathtype off the model stage
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) Unstage(stage *Stage) *CodeRouteSegmentPathType {
	__gong__unstage(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathTypes_mapString, coderoutesegmentpathtype, coderoutesegmentpathtype.Name)
	return coderoutesegmentpathtype
}

// UnstageVoid removes coderoutesegmentpathtype off the model stage
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) UnstageVoid(stage *Stage) {
	coderoutesegmentpathtype.Unstage(stage)
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) StageVoid(stage *Stage) {
	coderoutesegmentpathtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GetName() (res string) {
	return coderoutesegmentpathtype.Name
}

// for satisfaction of GongStruct interface
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) SetName(name string) {
	coderoutesegmentpathtype.Name = name
}

// Stage puts coderoutetype to the model stage
func (coderoutetype *CodeRouteType) Stage(stage *Stage) *CodeRouteType {
	__gong__stage(stage.CodeRouteTypes, stage.CodeRouteType_stagedOrder, stage.CodeRouteType_orderStaged, &stage.CodeRouteTypeOrder, stage.CodeRouteTypes_mapString, coderoutetype, coderoutetype.Name)
	return coderoutetype
}

// StagePreserveOrder puts coderoutetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteTypeOrder
// - update stage.CodeRouteTypeOrder accordingly
func (coderoutetype *CodeRouteType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteTypes, stage.CodeRouteType_stagedOrder, stage.CodeRouteType_orderStaged, &stage.CodeRouteTypeOrder, stage.CodeRouteTypes_mapString, coderoutetype, order, coderoutetype.Name)
}

// Unstage removes coderoutetype off the model stage
func (coderoutetype *CodeRouteType) Unstage(stage *Stage) *CodeRouteType {
	__gong__unstage(stage.CodeRouteTypes, stage.CodeRouteTypes_mapString, coderoutetype, coderoutetype.Name)
	return coderoutetype
}

// UnstageVoid removes coderoutetype off the model stage
func (coderoutetype *CodeRouteType) UnstageVoid(stage *Stage) {
	coderoutetype.Unstage(stage)
}

func (coderoutetype *CodeRouteType) StageVoid(stage *Stage) {
	coderoutetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoutetype *CodeRouteType) GetName() (res string) {
	return coderoutetype.Name
}

// for satisfaction of GongStruct interface
func (coderoutetype *CodeRouteType) SetName(name string) {
	coderoutetype.Name = name
}

// Stage puts coderuleproceduretitletype to the model stage
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) Stage(stage *Stage) *CodeRuleProcedureTitleType {
	__gong__stage(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleType_stagedOrder, stage.CodeRuleProcedureTitleType_orderStaged, &stage.CodeRuleProcedureTitleTypeOrder, stage.CodeRuleProcedureTitleTypes_mapString, coderuleproceduretitletype, coderuleproceduretitletype.Name)
	return coderuleproceduretitletype
}

// StagePreserveOrder puts coderuleproceduretitletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRuleProcedureTitleTypeOrder
// - update stage.CodeRuleProcedureTitleTypeOrder accordingly
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleType_stagedOrder, stage.CodeRuleProcedureTitleType_orderStaged, &stage.CodeRuleProcedureTitleTypeOrder, stage.CodeRuleProcedureTitleTypes_mapString, coderuleproceduretitletype, order, coderuleproceduretitletype.Name)
}

// Unstage removes coderuleproceduretitletype off the model stage
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) Unstage(stage *Stage) *CodeRuleProcedureTitleType {
	__gong__unstage(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleTypes_mapString, coderuleproceduretitletype, coderuleproceduretitletype.Name)
	return coderuleproceduretitletype
}

// UnstageVoid removes coderuleproceduretitletype off the model stage
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) UnstageVoid(stage *Stage) {
	coderuleproceduretitletype.Unstage(stage)
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) StageVoid(stage *Stage) {
	coderuleproceduretitletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GetName() (res string) {
	return coderuleproceduretitletype.Name
}

// for satisfaction of GongStruct interface
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) SetName(name string) {
	coderuleproceduretitletype.Name = name
}

// Stage puts coderuleproceduretype to the model stage
func (coderuleproceduretype *CodeRuleProcedureType) Stage(stage *Stage) *CodeRuleProcedureType {
	__gong__stage(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureType_stagedOrder, stage.CodeRuleProcedureType_orderStaged, &stage.CodeRuleProcedureTypeOrder, stage.CodeRuleProcedureTypes_mapString, coderuleproceduretype, coderuleproceduretype.Name)
	return coderuleproceduretype
}

// StagePreserveOrder puts coderuleproceduretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRuleProcedureTypeOrder
// - update stage.CodeRuleProcedureTypeOrder accordingly
func (coderuleproceduretype *CodeRuleProcedureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureType_stagedOrder, stage.CodeRuleProcedureType_orderStaged, &stage.CodeRuleProcedureTypeOrder, stage.CodeRuleProcedureTypes_mapString, coderuleproceduretype, order, coderuleproceduretype.Name)
}

// Unstage removes coderuleproceduretype off the model stage
func (coderuleproceduretype *CodeRuleProcedureType) Unstage(stage *Stage) *CodeRuleProcedureType {
	__gong__unstage(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureTypes_mapString, coderuleproceduretype, coderuleproceduretype.Name)
	return coderuleproceduretype
}

// UnstageVoid removes coderuleproceduretype off the model stage
func (coderuleproceduretype *CodeRuleProcedureType) UnstageVoid(stage *Stage) {
	coderuleproceduretype.Unstage(stage)
}

func (coderuleproceduretype *CodeRuleProcedureType) StageVoid(stage *Stage) {
	coderuleproceduretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderuleproceduretype *CodeRuleProcedureType) GetName() (res string) {
	return coderuleproceduretype.Name
}

// for satisfaction of GongStruct interface
func (coderuleproceduretype *CodeRuleProcedureType) SetName(name string) {
	coderuleproceduretype.Name = name
}

// Stage puts codesidetype to the model stage
func (codesidetype *CodeSideType) Stage(stage *Stage) *CodeSideType {
	__gong__stage(stage.CodeSideTypes, stage.CodeSideType_stagedOrder, stage.CodeSideType_orderStaged, &stage.CodeSideTypeOrder, stage.CodeSideTypes_mapString, codesidetype, codesidetype.Name)
	return codesidetype
}

// StagePreserveOrder puts codesidetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSideTypeOrder
// - update stage.CodeSideTypeOrder accordingly
func (codesidetype *CodeSideType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSideTypes, stage.CodeSideType_stagedOrder, stage.CodeSideType_orderStaged, &stage.CodeSideTypeOrder, stage.CodeSideTypes_mapString, codesidetype, order, codesidetype.Name)
}

// Unstage removes codesidetype off the model stage
func (codesidetype *CodeSideType) Unstage(stage *Stage) *CodeSideType {
	__gong__unstage(stage.CodeSideTypes, stage.CodeSideTypes_mapString, codesidetype, codesidetype.Name)
	return codesidetype
}

// UnstageVoid removes codesidetype off the model stage
func (codesidetype *CodeSideType) UnstageVoid(stage *Stage) {
	codesidetype.Unstage(stage)
}

func (codesidetype *CodeSideType) StageVoid(stage *Stage) {
	codesidetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesidetype *CodeSideType) GetName() (res string) {
	return codesidetype.Name
}

// for satisfaction of GongStruct interface
func (codesidetype *CodeSideType) SetName(name string) {
	codesidetype.Name = name
}

// Stage puts codespeedreferencetype to the model stage
func (codespeedreferencetype *CodeSpeedReferenceType) Stage(stage *Stage) *CodeSpeedReferenceType {
	__gong__stage(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceType_stagedOrder, stage.CodeSpeedReferenceType_orderStaged, &stage.CodeSpeedReferenceTypeOrder, stage.CodeSpeedReferenceTypes_mapString, codespeedreferencetype, codespeedreferencetype.Name)
	return codespeedreferencetype
}

// StagePreserveOrder puts codespeedreferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSpeedReferenceTypeOrder
// - update stage.CodeSpeedReferenceTypeOrder accordingly
func (codespeedreferencetype *CodeSpeedReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceType_stagedOrder, stage.CodeSpeedReferenceType_orderStaged, &stage.CodeSpeedReferenceTypeOrder, stage.CodeSpeedReferenceTypes_mapString, codespeedreferencetype, order, codespeedreferencetype.Name)
}

// Unstage removes codespeedreferencetype off the model stage
func (codespeedreferencetype *CodeSpeedReferenceType) Unstage(stage *Stage) *CodeSpeedReferenceType {
	__gong__unstage(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceTypes_mapString, codespeedreferencetype, codespeedreferencetype.Name)
	return codespeedreferencetype
}

// UnstageVoid removes codespeedreferencetype off the model stage
func (codespeedreferencetype *CodeSpeedReferenceType) UnstageVoid(stage *Stage) {
	codespeedreferencetype.Unstage(stage)
}

func (codespeedreferencetype *CodeSpeedReferenceType) StageVoid(stage *Stage) {
	codespeedreferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codespeedreferencetype *CodeSpeedReferenceType) GetName() (res string) {
	return codespeedreferencetype.Name
}

// for satisfaction of GongStruct interface
func (codespeedreferencetype *CodeSpeedReferenceType) SetName(name string) {
	codespeedreferencetype.Name = name
}

// Stage puts codetacanchanneltype to the model stage
func (codetacanchanneltype *CodeTACANChannelType) Stage(stage *Stage) *CodeTACANChannelType {
	__gong__stage(stage.CodeTACANChannelTypes, stage.CodeTACANChannelType_stagedOrder, stage.CodeTACANChannelType_orderStaged, &stage.CodeTACANChannelTypeOrder, stage.CodeTACANChannelTypes_mapString, codetacanchanneltype, codetacanchanneltype.Name)
	return codetacanchanneltype
}

// StagePreserveOrder puts codetacanchanneltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTACANChannelTypeOrder
// - update stage.CodeTACANChannelTypeOrder accordingly
func (codetacanchanneltype *CodeTACANChannelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTACANChannelTypes, stage.CodeTACANChannelType_stagedOrder, stage.CodeTACANChannelType_orderStaged, &stage.CodeTACANChannelTypeOrder, stage.CodeTACANChannelTypes_mapString, codetacanchanneltype, order, codetacanchanneltype.Name)
}

// Unstage removes codetacanchanneltype off the model stage
func (codetacanchanneltype *CodeTACANChannelType) Unstage(stage *Stage) *CodeTACANChannelType {
	__gong__unstage(stage.CodeTACANChannelTypes, stage.CodeTACANChannelTypes_mapString, codetacanchanneltype, codetacanchanneltype.Name)
	return codetacanchanneltype
}

// UnstageVoid removes codetacanchanneltype off the model stage
func (codetacanchanneltype *CodeTACANChannelType) UnstageVoid(stage *Stage) {
	codetacanchanneltype.Unstage(stage)
}

func (codetacanchanneltype *CodeTACANChannelType) StageVoid(stage *Stage) {
	codetacanchanneltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetacanchanneltype *CodeTACANChannelType) GetName() (res string) {
	return codetacanchanneltype.Name
}

// for satisfaction of GongStruct interface
func (codetacanchanneltype *CodeTACANChannelType) SetName(name string) {
	codetacanchanneltype.Name = name
}

// Stage puts codetelecomnetworktype to the model stage
func (codetelecomnetworktype *CodeTelecomNetworkType) Stage(stage *Stage) *CodeTelecomNetworkType {
	__gong__stage(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkType_stagedOrder, stage.CodeTelecomNetworkType_orderStaged, &stage.CodeTelecomNetworkTypeOrder, stage.CodeTelecomNetworkTypes_mapString, codetelecomnetworktype, codetelecomnetworktype.Name)
	return codetelecomnetworktype
}

// StagePreserveOrder puts codetelecomnetworktype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTelecomNetworkTypeOrder
// - update stage.CodeTelecomNetworkTypeOrder accordingly
func (codetelecomnetworktype *CodeTelecomNetworkType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkType_stagedOrder, stage.CodeTelecomNetworkType_orderStaged, &stage.CodeTelecomNetworkTypeOrder, stage.CodeTelecomNetworkTypes_mapString, codetelecomnetworktype, order, codetelecomnetworktype.Name)
}

// Unstage removes codetelecomnetworktype off the model stage
func (codetelecomnetworktype *CodeTelecomNetworkType) Unstage(stage *Stage) *CodeTelecomNetworkType {
	__gong__unstage(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_mapString, codetelecomnetworktype, codetelecomnetworktype.Name)
	return codetelecomnetworktype
}

// UnstageVoid removes codetelecomnetworktype off the model stage
func (codetelecomnetworktype *CodeTelecomNetworkType) UnstageVoid(stage *Stage) {
	codetelecomnetworktype.Unstage(stage)
}

func (codetelecomnetworktype *CodeTelecomNetworkType) StageVoid(stage *Stage) {
	codetelecomnetworktype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetelecomnetworktype *CodeTelecomNetworkType) GetName() (res string) {
	return codetelecomnetworktype.Name
}

// for satisfaction of GongStruct interface
func (codetelecomnetworktype *CodeTelecomNetworkType) SetName(name string) {
	codetelecomnetworktype.Name = name
}

// Stage puts codetimeeventcombinationtype to the model stage
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) Stage(stage *Stage) *CodeTimeEventCombinationType {
	__gong__stage(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationType_stagedOrder, stage.CodeTimeEventCombinationType_orderStaged, &stage.CodeTimeEventCombinationTypeOrder, stage.CodeTimeEventCombinationTypes_mapString, codetimeeventcombinationtype, codetimeeventcombinationtype.Name)
	return codetimeeventcombinationtype
}

// StagePreserveOrder puts codetimeeventcombinationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTimeEventCombinationTypeOrder
// - update stage.CodeTimeEventCombinationTypeOrder accordingly
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationType_stagedOrder, stage.CodeTimeEventCombinationType_orderStaged, &stage.CodeTimeEventCombinationTypeOrder, stage.CodeTimeEventCombinationTypes_mapString, codetimeeventcombinationtype, order, codetimeeventcombinationtype.Name)
}

// Unstage removes codetimeeventcombinationtype off the model stage
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) Unstage(stage *Stage) *CodeTimeEventCombinationType {
	__gong__unstage(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_mapString, codetimeeventcombinationtype, codetimeeventcombinationtype.Name)
	return codetimeeventcombinationtype
}

// UnstageVoid removes codetimeeventcombinationtype off the model stage
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) UnstageVoid(stage *Stage) {
	codetimeeventcombinationtype.Unstage(stage)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) StageVoid(stage *Stage) {
	codetimeeventcombinationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GetName() (res string) {
	return codetimeeventcombinationtype.Name
}

// for satisfaction of GongStruct interface
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) SetName(name string) {
	codetimeeventcombinationtype.Name = name
}

// Stage puts codetimeeventtype to the model stage
func (codetimeeventtype *CodeTimeEventType) Stage(stage *Stage) *CodeTimeEventType {
	__gong__stage(stage.CodeTimeEventTypes, stage.CodeTimeEventType_stagedOrder, stage.CodeTimeEventType_orderStaged, &stage.CodeTimeEventTypeOrder, stage.CodeTimeEventTypes_mapString, codetimeeventtype, codetimeeventtype.Name)
	return codetimeeventtype
}

// StagePreserveOrder puts codetimeeventtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTimeEventTypeOrder
// - update stage.CodeTimeEventTypeOrder accordingly
func (codetimeeventtype *CodeTimeEventType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTimeEventTypes, stage.CodeTimeEventType_stagedOrder, stage.CodeTimeEventType_orderStaged, &stage.CodeTimeEventTypeOrder, stage.CodeTimeEventTypes_mapString, codetimeeventtype, order, codetimeeventtype.Name)
}

// Unstage removes codetimeeventtype off the model stage
func (codetimeeventtype *CodeTimeEventType) Unstage(stage *Stage) *CodeTimeEventType {
	__gong__unstage(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_mapString, codetimeeventtype, codetimeeventtype.Name)
	return codetimeeventtype
}

// UnstageVoid removes codetimeeventtype off the model stage
func (codetimeeventtype *CodeTimeEventType) UnstageVoid(stage *Stage) {
	codetimeeventtype.Unstage(stage)
}

func (codetimeeventtype *CodeTimeEventType) StageVoid(stage *Stage) {
	codetimeeventtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetimeeventtype *CodeTimeEventType) GetName() (res string) {
	return codetimeeventtype.Name
}

// for satisfaction of GongStruct interface
func (codetimeeventtype *CodeTimeEventType) SetName(name string) {
	codetimeeventtype.Name = name
}

// Stage puts codetimereferencetype to the model stage
func (codetimereferencetype *CodeTimeReferenceType) Stage(stage *Stage) *CodeTimeReferenceType {
	__gong__stage(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceType_stagedOrder, stage.CodeTimeReferenceType_orderStaged, &stage.CodeTimeReferenceTypeOrder, stage.CodeTimeReferenceTypes_mapString, codetimereferencetype, codetimereferencetype.Name)
	return codetimereferencetype
}

// StagePreserveOrder puts codetimereferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTimeReferenceTypeOrder
// - update stage.CodeTimeReferenceTypeOrder accordingly
func (codetimereferencetype *CodeTimeReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceType_stagedOrder, stage.CodeTimeReferenceType_orderStaged, &stage.CodeTimeReferenceTypeOrder, stage.CodeTimeReferenceTypes_mapString, codetimereferencetype, order, codetimereferencetype.Name)
}

// Unstage removes codetimereferencetype off the model stage
func (codetimereferencetype *CodeTimeReferenceType) Unstage(stage *Stage) *CodeTimeReferenceType {
	__gong__unstage(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_mapString, codetimereferencetype, codetimereferencetype.Name)
	return codetimereferencetype
}

// UnstageVoid removes codetimereferencetype off the model stage
func (codetimereferencetype *CodeTimeReferenceType) UnstageVoid(stage *Stage) {
	codetimereferencetype.Unstage(stage)
}

func (codetimereferencetype *CodeTimeReferenceType) StageVoid(stage *Stage) {
	codetimereferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetimereferencetype *CodeTimeReferenceType) GetName() (res string) {
	return codetimereferencetype.Name
}

// for satisfaction of GongStruct interface
func (codetimereferencetype *CodeTimeReferenceType) SetName(name string) {
	codetimereferencetype.Name = name
}

// Stage puts codetranspondertype to the model stage
func (codetranspondertype *CodeTransponderType) Stage(stage *Stage) *CodeTransponderType {
	__gong__stage(stage.CodeTransponderTypes, stage.CodeTransponderType_stagedOrder, stage.CodeTransponderType_orderStaged, &stage.CodeTransponderTypeOrder, stage.CodeTransponderTypes_mapString, codetranspondertype, codetranspondertype.Name)
	return codetranspondertype
}

// StagePreserveOrder puts codetranspondertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTransponderTypeOrder
// - update stage.CodeTransponderTypeOrder accordingly
func (codetranspondertype *CodeTransponderType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTransponderTypes, stage.CodeTransponderType_stagedOrder, stage.CodeTransponderType_orderStaged, &stage.CodeTransponderTypeOrder, stage.CodeTransponderTypes_mapString, codetranspondertype, order, codetranspondertype.Name)
}

// Unstage removes codetranspondertype off the model stage
func (codetranspondertype *CodeTransponderType) Unstage(stage *Stage) *CodeTransponderType {
	__gong__unstage(stage.CodeTransponderTypes, stage.CodeTransponderTypes_mapString, codetranspondertype, codetranspondertype.Name)
	return codetranspondertype
}

// UnstageVoid removes codetranspondertype off the model stage
func (codetranspondertype *CodeTransponderType) UnstageVoid(stage *Stage) {
	codetranspondertype.Unstage(stage)
}

func (codetranspondertype *CodeTransponderType) StageVoid(stage *Stage) {
	codetranspondertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetranspondertype *CodeTransponderType) GetName() (res string) {
	return codetranspondertype.Name
}

// for satisfaction of GongStruct interface
func (codetranspondertype *CodeTransponderType) SetName(name string) {
	codetranspondertype.Name = name
}

// Stage puts codeupperalphatype to the model stage
func (codeupperalphatype *CodeUpperAlphaType) Stage(stage *Stage) *CodeUpperAlphaType {
	__gong__stage(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaType_stagedOrder, stage.CodeUpperAlphaType_orderStaged, &stage.CodeUpperAlphaTypeOrder, stage.CodeUpperAlphaTypes_mapString, codeupperalphatype, codeupperalphatype.Name)
	return codeupperalphatype
}

// StagePreserveOrder puts codeupperalphatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeUpperAlphaTypeOrder
// - update stage.CodeUpperAlphaTypeOrder accordingly
func (codeupperalphatype *CodeUpperAlphaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaType_stagedOrder, stage.CodeUpperAlphaType_orderStaged, &stage.CodeUpperAlphaTypeOrder, stage.CodeUpperAlphaTypes_mapString, codeupperalphatype, order, codeupperalphatype.Name)
}

// Unstage removes codeupperalphatype off the model stage
func (codeupperalphatype *CodeUpperAlphaType) Unstage(stage *Stage) *CodeUpperAlphaType {
	__gong__unstage(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaTypes_mapString, codeupperalphatype, codeupperalphatype.Name)
	return codeupperalphatype
}

// UnstageVoid removes codeupperalphatype off the model stage
func (codeupperalphatype *CodeUpperAlphaType) UnstageVoid(stage *Stage) {
	codeupperalphatype.Unstage(stage)
}

func (codeupperalphatype *CodeUpperAlphaType) StageVoid(stage *Stage) {
	codeupperalphatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeupperalphatype *CodeUpperAlphaType) GetName() (res string) {
	return codeupperalphatype.Name
}

// for satisfaction of GongStruct interface
func (codeupperalphatype *CodeUpperAlphaType) SetName(name string) {
	codeupperalphatype.Name = name
}

// Stage puts codevalueinterpretationtype to the model stage
func (codevalueinterpretationtype *CodeValueInterpretationType) Stage(stage *Stage) *CodeValueInterpretationType {
	__gong__stage(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationType_stagedOrder, stage.CodeValueInterpretationType_orderStaged, &stage.CodeValueInterpretationTypeOrder, stage.CodeValueInterpretationTypes_mapString, codevalueinterpretationtype, codevalueinterpretationtype.Name)
	return codevalueinterpretationtype
}

// StagePreserveOrder puts codevalueinterpretationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeValueInterpretationTypeOrder
// - update stage.CodeValueInterpretationTypeOrder accordingly
func (codevalueinterpretationtype *CodeValueInterpretationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationType_stagedOrder, stage.CodeValueInterpretationType_orderStaged, &stage.CodeValueInterpretationTypeOrder, stage.CodeValueInterpretationTypes_mapString, codevalueinterpretationtype, order, codevalueinterpretationtype.Name)
}

// Unstage removes codevalueinterpretationtype off the model stage
func (codevalueinterpretationtype *CodeValueInterpretationType) Unstage(stage *Stage) *CodeValueInterpretationType {
	__gong__unstage(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationTypes_mapString, codevalueinterpretationtype, codevalueinterpretationtype.Name)
	return codevalueinterpretationtype
}

// UnstageVoid removes codevalueinterpretationtype off the model stage
func (codevalueinterpretationtype *CodeValueInterpretationType) UnstageVoid(stage *Stage) {
	codevalueinterpretationtype.Unstage(stage)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) StageVoid(stage *Stage) {
	codevalueinterpretationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codevalueinterpretationtype *CodeValueInterpretationType) GetName() (res string) {
	return codevalueinterpretationtype.Name
}

// for satisfaction of GongStruct interface
func (codevalueinterpretationtype *CodeValueInterpretationType) SetName(name string) {
	codevalueinterpretationtype.Name = name
}

// Stage puts codeverticaldatumtype to the model stage
func (codeverticaldatumtype *CodeVerticalDatumType) Stage(stage *Stage) *CodeVerticalDatumType {
	__gong__stage(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumType_orderStaged, &stage.CodeVerticalDatumTypeOrder, stage.CodeVerticalDatumTypes_mapString, codeverticaldatumtype, codeverticaldatumtype.Name)
	return codeverticaldatumtype
}

// StagePreserveOrder puts codeverticaldatumtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVerticalDatumTypeOrder
// - update stage.CodeVerticalDatumTypeOrder accordingly
func (codeverticaldatumtype *CodeVerticalDatumType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumType_orderStaged, &stage.CodeVerticalDatumTypeOrder, stage.CodeVerticalDatumTypes_mapString, codeverticaldatumtype, order, codeverticaldatumtype.Name)
}

// Unstage removes codeverticaldatumtype off the model stage
func (codeverticaldatumtype *CodeVerticalDatumType) Unstage(stage *Stage) *CodeVerticalDatumType {
	__gong__unstage(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_mapString, codeverticaldatumtype, codeverticaldatumtype.Name)
	return codeverticaldatumtype
}

// UnstageVoid removes codeverticaldatumtype off the model stage
func (codeverticaldatumtype *CodeVerticalDatumType) UnstageVoid(stage *Stage) {
	codeverticaldatumtype.Unstage(stage)
}

func (codeverticaldatumtype *CodeVerticalDatumType) StageVoid(stage *Stage) {
	codeverticaldatumtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeverticaldatumtype *CodeVerticalDatumType) GetName() (res string) {
	return codeverticaldatumtype.Name
}

// for satisfaction of GongStruct interface
func (codeverticaldatumtype *CodeVerticalDatumType) SetName(name string) {
	codeverticaldatumtype.Name = name
}

// Stage puts codeverticalreferencetype to the model stage
func (codeverticalreferencetype *CodeVerticalReferenceType) Stage(stage *Stage) *CodeVerticalReferenceType {
	__gong__stage(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceType_stagedOrder, stage.CodeVerticalReferenceType_orderStaged, &stage.CodeVerticalReferenceTypeOrder, stage.CodeVerticalReferenceTypes_mapString, codeverticalreferencetype, codeverticalreferencetype.Name)
	return codeverticalreferencetype
}

// StagePreserveOrder puts codeverticalreferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVerticalReferenceTypeOrder
// - update stage.CodeVerticalReferenceTypeOrder accordingly
func (codeverticalreferencetype *CodeVerticalReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceType_stagedOrder, stage.CodeVerticalReferenceType_orderStaged, &stage.CodeVerticalReferenceTypeOrder, stage.CodeVerticalReferenceTypes_mapString, codeverticalreferencetype, order, codeverticalreferencetype.Name)
}

// Unstage removes codeverticalreferencetype off the model stage
func (codeverticalreferencetype *CodeVerticalReferenceType) Unstage(stage *Stage) *CodeVerticalReferenceType {
	__gong__unstage(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceTypes_mapString, codeverticalreferencetype, codeverticalreferencetype.Name)
	return codeverticalreferencetype
}

// UnstageVoid removes codeverticalreferencetype off the model stage
func (codeverticalreferencetype *CodeVerticalReferenceType) UnstageVoid(stage *Stage) {
	codeverticalreferencetype.Unstage(stage)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) StageVoid(stage *Stage) {
	codeverticalreferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeverticalreferencetype *CodeVerticalReferenceType) GetName() (res string) {
	return codeverticalreferencetype.Name
}

// for satisfaction of GongStruct interface
func (codeverticalreferencetype *CodeVerticalReferenceType) SetName(name string) {
	codeverticalreferencetype.Name = name
}

// Stage puts codewaketurbulencetype to the model stage
func (codewaketurbulencetype *CodeWakeTurbulenceType) Stage(stage *Stage) *CodeWakeTurbulenceType {
	__gong__stage(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceType_stagedOrder, stage.CodeWakeTurbulenceType_orderStaged, &stage.CodeWakeTurbulenceTypeOrder, stage.CodeWakeTurbulenceTypes_mapString, codewaketurbulencetype, codewaketurbulencetype.Name)
	return codewaketurbulencetype
}

// StagePreserveOrder puts codewaketurbulencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeWakeTurbulenceTypeOrder
// - update stage.CodeWakeTurbulenceTypeOrder accordingly
func (codewaketurbulencetype *CodeWakeTurbulenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceType_stagedOrder, stage.CodeWakeTurbulenceType_orderStaged, &stage.CodeWakeTurbulenceTypeOrder, stage.CodeWakeTurbulenceTypes_mapString, codewaketurbulencetype, order, codewaketurbulencetype.Name)
}

// Unstage removes codewaketurbulencetype off the model stage
func (codewaketurbulencetype *CodeWakeTurbulenceType) Unstage(stage *Stage) *CodeWakeTurbulenceType {
	__gong__unstage(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceTypes_mapString, codewaketurbulencetype, codewaketurbulencetype.Name)
	return codewaketurbulencetype
}

// UnstageVoid removes codewaketurbulencetype off the model stage
func (codewaketurbulencetype *CodeWakeTurbulenceType) UnstageVoid(stage *Stage) {
	codewaketurbulencetype.Unstage(stage)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) StageVoid(stage *Stage) {
	codewaketurbulencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codewaketurbulencetype *CodeWakeTurbulenceType) GetName() (res string) {
	return codewaketurbulencetype.Name
}

// for satisfaction of GongStruct interface
func (codewaketurbulencetype *CodeWakeTurbulenceType) SetName(name string) {
	codewaketurbulencetype.Name = name
}

// Stage puts codeyesnotype to the model stage
func (codeyesnotype *CodeYesNoType) Stage(stage *Stage) *CodeYesNoType {
	__gong__stage(stage.CodeYesNoTypes, stage.CodeYesNoType_stagedOrder, stage.CodeYesNoType_orderStaged, &stage.CodeYesNoTypeOrder, stage.CodeYesNoTypes_mapString, codeyesnotype, codeyesnotype.Name)
	return codeyesnotype
}

// StagePreserveOrder puts codeyesnotype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeYesNoTypeOrder
// - update stage.CodeYesNoTypeOrder accordingly
func (codeyesnotype *CodeYesNoType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeYesNoTypes, stage.CodeYesNoType_stagedOrder, stage.CodeYesNoType_orderStaged, &stage.CodeYesNoTypeOrder, stage.CodeYesNoTypes_mapString, codeyesnotype, order, codeyesnotype.Name)
}

// Unstage removes codeyesnotype off the model stage
func (codeyesnotype *CodeYesNoType) Unstage(stage *Stage) *CodeYesNoType {
	__gong__unstage(stage.CodeYesNoTypes, stage.CodeYesNoTypes_mapString, codeyesnotype, codeyesnotype.Name)
	return codeyesnotype
}

// UnstageVoid removes codeyesnotype off the model stage
func (codeyesnotype *CodeYesNoType) UnstageVoid(stage *Stage) {
	codeyesnotype.Unstage(stage)
}

func (codeyesnotype *CodeYesNoType) StageVoid(stage *Stage) {
	codeyesnotype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeyesnotype *CodeYesNoType) GetName() (res string) {
	return codeyesnotype.Name
}

// for satisfaction of GongStruct interface
func (codeyesnotype *CodeYesNoType) SetName(name string) {
	codeyesnotype.Name = name
}

// Stage puts datemonthdaytype to the model stage
func (datemonthdaytype *DateMonthDayType) Stage(stage *Stage) *DateMonthDayType {
	__gong__stage(stage.DateMonthDayTypes, stage.DateMonthDayType_stagedOrder, stage.DateMonthDayType_orderStaged, &stage.DateMonthDayTypeOrder, stage.DateMonthDayTypes_mapString, datemonthdaytype, datemonthdaytype.Name)
	return datemonthdaytype
}

// StagePreserveOrder puts datemonthdaytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DateMonthDayTypeOrder
// - update stage.DateMonthDayTypeOrder accordingly
func (datemonthdaytype *DateMonthDayType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DateMonthDayTypes, stage.DateMonthDayType_stagedOrder, stage.DateMonthDayType_orderStaged, &stage.DateMonthDayTypeOrder, stage.DateMonthDayTypes_mapString, datemonthdaytype, order, datemonthdaytype.Name)
}

// Unstage removes datemonthdaytype off the model stage
func (datemonthdaytype *DateMonthDayType) Unstage(stage *Stage) *DateMonthDayType {
	__gong__unstage(stage.DateMonthDayTypes, stage.DateMonthDayTypes_mapString, datemonthdaytype, datemonthdaytype.Name)
	return datemonthdaytype
}

// UnstageVoid removes datemonthdaytype off the model stage
func (datemonthdaytype *DateMonthDayType) UnstageVoid(stage *Stage) {
	datemonthdaytype.Unstage(stage)
}

func (datemonthdaytype *DateMonthDayType) StageVoid(stage *Stage) {
	datemonthdaytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (datemonthdaytype *DateMonthDayType) GetName() (res string) {
	return datemonthdaytype.Name
}

// for satisfaction of GongStruct interface
func (datemonthdaytype *DateMonthDayType) SetName(name string) {
	datemonthdaytype.Name = name
}

// Stage puts nonumbertype to the model stage
func (nonumbertype *NoNumberType) Stage(stage *Stage) *NoNumberType {
	__gong__stage(stage.NoNumberTypes, stage.NoNumberType_stagedOrder, stage.NoNumberType_orderStaged, &stage.NoNumberTypeOrder, stage.NoNumberTypes_mapString, nonumbertype, nonumbertype.Name)
	return nonumbertype
}

// StagePreserveOrder puts nonumbertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoNumberTypeOrder
// - update stage.NoNumberTypeOrder accordingly
func (nonumbertype *NoNumberType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoNumberTypes, stage.NoNumberType_stagedOrder, stage.NoNumberType_orderStaged, &stage.NoNumberTypeOrder, stage.NoNumberTypes_mapString, nonumbertype, order, nonumbertype.Name)
}

// Unstage removes nonumbertype off the model stage
func (nonumbertype *NoNumberType) Unstage(stage *Stage) *NoNumberType {
	__gong__unstage(stage.NoNumberTypes, stage.NoNumberTypes_mapString, nonumbertype, nonumbertype.Name)
	return nonumbertype
}

// UnstageVoid removes nonumbertype off the model stage
func (nonumbertype *NoNumberType) UnstageVoid(stage *Stage) {
	nonumbertype.Unstage(stage)
}

func (nonumbertype *NoNumberType) StageVoid(stage *Stage) {
	nonumbertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (nonumbertype *NoNumberType) GetName() (res string) {
	return nonumbertype.Name
}

// for satisfaction of GongStruct interface
func (nonumbertype *NoNumberType) SetName(name string) {
	nonumbertype.Name = name
}

// Stage puts nosequencetype to the model stage
func (nosequencetype *NoSequenceType) Stage(stage *Stage) *NoSequenceType {
	__gong__stage(stage.NoSequenceTypes, stage.NoSequenceType_stagedOrder, stage.NoSequenceType_orderStaged, &stage.NoSequenceTypeOrder, stage.NoSequenceTypes_mapString, nosequencetype, nosequencetype.Name)
	return nosequencetype
}

// StagePreserveOrder puts nosequencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoSequenceTypeOrder
// - update stage.NoSequenceTypeOrder accordingly
func (nosequencetype *NoSequenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoSequenceTypes, stage.NoSequenceType_stagedOrder, stage.NoSequenceType_orderStaged, &stage.NoSequenceTypeOrder, stage.NoSequenceTypes_mapString, nosequencetype, order, nosequencetype.Name)
}

// Unstage removes nosequencetype off the model stage
func (nosequencetype *NoSequenceType) Unstage(stage *Stage) *NoSequenceType {
	__gong__unstage(stage.NoSequenceTypes, stage.NoSequenceTypes_mapString, nosequencetype, nosequencetype.Name)
	return nosequencetype
}

// UnstageVoid removes nosequencetype off the model stage
func (nosequencetype *NoSequenceType) UnstageVoid(stage *Stage) {
	nosequencetype.Unstage(stage)
}

func (nosequencetype *NoSequenceType) StageVoid(stage *Stage) {
	nosequencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (nosequencetype *NoSequenceType) GetName() (res string) {
	return nosequencetype.Name
}

// for satisfaction of GongStruct interface
func (nosequencetype *NoSequenceType) SetName(name string) {
	nosequencetype.Name = name
}

// Stage puts textaddresstype to the model stage
func (textaddresstype *TextAddressType) Stage(stage *Stage) *TextAddressType {
	__gong__stage(stage.TextAddressTypes, stage.TextAddressType_stagedOrder, stage.TextAddressType_orderStaged, &stage.TextAddressTypeOrder, stage.TextAddressTypes_mapString, textaddresstype, textaddresstype.Name)
	return textaddresstype
}

// StagePreserveOrder puts textaddresstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextAddressTypeOrder
// - update stage.TextAddressTypeOrder accordingly
func (textaddresstype *TextAddressType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextAddressTypes, stage.TextAddressType_stagedOrder, stage.TextAddressType_orderStaged, &stage.TextAddressTypeOrder, stage.TextAddressTypes_mapString, textaddresstype, order, textaddresstype.Name)
}

// Unstage removes textaddresstype off the model stage
func (textaddresstype *TextAddressType) Unstage(stage *Stage) *TextAddressType {
	__gong__unstage(stage.TextAddressTypes, stage.TextAddressTypes_mapString, textaddresstype, textaddresstype.Name)
	return textaddresstype
}

// UnstageVoid removes textaddresstype off the model stage
func (textaddresstype *TextAddressType) UnstageVoid(stage *Stage) {
	textaddresstype.Unstage(stage)
}

func (textaddresstype *TextAddressType) StageVoid(stage *Stage) {
	textaddresstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textaddresstype *TextAddressType) GetName() (res string) {
	return textaddresstype.Name
}

// for satisfaction of GongStruct interface
func (textaddresstype *TextAddressType) SetName(name string) {
	textaddresstype.Name = name
}

// Stage puts textdesignatortype to the model stage
func (textdesignatortype *TextDesignatorType) Stage(stage *Stage) *TextDesignatorType {
	__gong__stage(stage.TextDesignatorTypes, stage.TextDesignatorType_stagedOrder, stage.TextDesignatorType_orderStaged, &stage.TextDesignatorTypeOrder, stage.TextDesignatorTypes_mapString, textdesignatortype, textdesignatortype.Name)
	return textdesignatortype
}

// StagePreserveOrder puts textdesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextDesignatorTypeOrder
// - update stage.TextDesignatorTypeOrder accordingly
func (textdesignatortype *TextDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextDesignatorTypes, stage.TextDesignatorType_stagedOrder, stage.TextDesignatorType_orderStaged, &stage.TextDesignatorTypeOrder, stage.TextDesignatorTypes_mapString, textdesignatortype, order, textdesignatortype.Name)
}

// Unstage removes textdesignatortype off the model stage
func (textdesignatortype *TextDesignatorType) Unstage(stage *Stage) *TextDesignatorType {
	__gong__unstage(stage.TextDesignatorTypes, stage.TextDesignatorTypes_mapString, textdesignatortype, textdesignatortype.Name)
	return textdesignatortype
}

// UnstageVoid removes textdesignatortype off the model stage
func (textdesignatortype *TextDesignatorType) UnstageVoid(stage *Stage) {
	textdesignatortype.Unstage(stage)
}

func (textdesignatortype *TextDesignatorType) StageVoid(stage *Stage) {
	textdesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textdesignatortype *TextDesignatorType) GetName() (res string) {
	return textdesignatortype.Name
}

// for satisfaction of GongStruct interface
func (textdesignatortype *TextDesignatorType) SetName(name string) {
	textdesignatortype.Name = name
}

// Stage puts textinstructiontype to the model stage
func (textinstructiontype *TextInstructionType) Stage(stage *Stage) *TextInstructionType {
	__gong__stage(stage.TextInstructionTypes, stage.TextInstructionType_stagedOrder, stage.TextInstructionType_orderStaged, &stage.TextInstructionTypeOrder, stage.TextInstructionTypes_mapString, textinstructiontype, textinstructiontype.Name)
	return textinstructiontype
}

// StagePreserveOrder puts textinstructiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextInstructionTypeOrder
// - update stage.TextInstructionTypeOrder accordingly
func (textinstructiontype *TextInstructionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextInstructionTypes, stage.TextInstructionType_stagedOrder, stage.TextInstructionType_orderStaged, &stage.TextInstructionTypeOrder, stage.TextInstructionTypes_mapString, textinstructiontype, order, textinstructiontype.Name)
}

// Unstage removes textinstructiontype off the model stage
func (textinstructiontype *TextInstructionType) Unstage(stage *Stage) *TextInstructionType {
	__gong__unstage(stage.TextInstructionTypes, stage.TextInstructionTypes_mapString, textinstructiontype, textinstructiontype.Name)
	return textinstructiontype
}

// UnstageVoid removes textinstructiontype off the model stage
func (textinstructiontype *TextInstructionType) UnstageVoid(stage *Stage) {
	textinstructiontype.Unstage(stage)
}

func (textinstructiontype *TextInstructionType) StageVoid(stage *Stage) {
	textinstructiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textinstructiontype *TextInstructionType) GetName() (res string) {
	return textinstructiontype.Name
}

// for satisfaction of GongStruct interface
func (textinstructiontype *TextInstructionType) SetName(name string) {
	textinstructiontype.Name = name
}

// Stage puts textnametype to the model stage
func (textnametype *TextNameType) Stage(stage *Stage) *TextNameType {
	__gong__stage(stage.TextNameTypes, stage.TextNameType_stagedOrder, stage.TextNameType_orderStaged, &stage.TextNameTypeOrder, stage.TextNameTypes_mapString, textnametype, textnametype.Name)
	return textnametype
}

// StagePreserveOrder puts textnametype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextNameTypeOrder
// - update stage.TextNameTypeOrder accordingly
func (textnametype *TextNameType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextNameTypes, stage.TextNameType_stagedOrder, stage.TextNameType_orderStaged, &stage.TextNameTypeOrder, stage.TextNameTypes_mapString, textnametype, order, textnametype.Name)
}

// Unstage removes textnametype off the model stage
func (textnametype *TextNameType) Unstage(stage *Stage) *TextNameType {
	__gong__unstage(stage.TextNameTypes, stage.TextNameTypes_mapString, textnametype, textnametype.Name)
	return textnametype
}

// UnstageVoid removes textnametype off the model stage
func (textnametype *TextNameType) UnstageVoid(stage *Stage) {
	textnametype.Unstage(stage)
}

func (textnametype *TextNameType) StageVoid(stage *Stage) {
	textnametype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textnametype *TextNameType) GetName() (res string) {
	return textnametype.Name
}

// for satisfaction of GongStruct interface
func (textnametype *TextNameType) SetName(name string) {
	textnametype.Name = name
}

// Stage puts textnotetype to the model stage
func (textnotetype *TextNoteType) Stage(stage *Stage) *TextNoteType {
	__gong__stage(stage.TextNoteTypes, stage.TextNoteType_stagedOrder, stage.TextNoteType_orderStaged, &stage.TextNoteTypeOrder, stage.TextNoteTypes_mapString, textnotetype, textnotetype.Name)
	return textnotetype
}

// StagePreserveOrder puts textnotetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextNoteTypeOrder
// - update stage.TextNoteTypeOrder accordingly
func (textnotetype *TextNoteType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextNoteTypes, stage.TextNoteType_stagedOrder, stage.TextNoteType_orderStaged, &stage.TextNoteTypeOrder, stage.TextNoteTypes_mapString, textnotetype, order, textnotetype.Name)
}

// Unstage removes textnotetype off the model stage
func (textnotetype *TextNoteType) Unstage(stage *Stage) *TextNoteType {
	__gong__unstage(stage.TextNoteTypes, stage.TextNoteTypes_mapString, textnotetype, textnotetype.Name)
	return textnotetype
}

// UnstageVoid removes textnotetype off the model stage
func (textnotetype *TextNoteType) UnstageVoid(stage *Stage) {
	textnotetype.Unstage(stage)
}

func (textnotetype *TextNoteType) StageVoid(stage *Stage) {
	textnotetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textnotetype *TextNoteType) GetName() (res string) {
	return textnotetype.Name
}

// for satisfaction of GongStruct interface
func (textnotetype *TextNoteType) SetName(name string) {
	textnotetype.Name = name
}

// Stage puts textphonetype to the model stage
func (textphonetype *TextPhoneType) Stage(stage *Stage) *TextPhoneType {
	__gong__stage(stage.TextPhoneTypes, stage.TextPhoneType_stagedOrder, stage.TextPhoneType_orderStaged, &stage.TextPhoneTypeOrder, stage.TextPhoneTypes_mapString, textphonetype, textphonetype.Name)
	return textphonetype
}

// StagePreserveOrder puts textphonetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextPhoneTypeOrder
// - update stage.TextPhoneTypeOrder accordingly
func (textphonetype *TextPhoneType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextPhoneTypes, stage.TextPhoneType_stagedOrder, stage.TextPhoneType_orderStaged, &stage.TextPhoneTypeOrder, stage.TextPhoneTypes_mapString, textphonetype, order, textphonetype.Name)
}

// Unstage removes textphonetype off the model stage
func (textphonetype *TextPhoneType) Unstage(stage *Stage) *TextPhoneType {
	__gong__unstage(stage.TextPhoneTypes, stage.TextPhoneTypes_mapString, textphonetype, textphonetype.Name)
	return textphonetype
}

// UnstageVoid removes textphonetype off the model stage
func (textphonetype *TextPhoneType) UnstageVoid(stage *Stage) {
	textphonetype.Unstage(stage)
}

func (textphonetype *TextPhoneType) StageVoid(stage *Stage) {
	textphonetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textphonetype *TextPhoneType) GetName() (res string) {
	return textphonetype.Name
}

// for satisfaction of GongStruct interface
func (textphonetype *TextPhoneType) SetName(name string) {
	textphonetype.Name = name
}

// Stage puts textpropertynametype to the model stage
func (textpropertynametype *TextPropertyNameType) Stage(stage *Stage) *TextPropertyNameType {
	__gong__stage(stage.TextPropertyNameTypes, stage.TextPropertyNameType_stagedOrder, stage.TextPropertyNameType_orderStaged, &stage.TextPropertyNameTypeOrder, stage.TextPropertyNameTypes_mapString, textpropertynametype, textpropertynametype.Name)
	return textpropertynametype
}

// StagePreserveOrder puts textpropertynametype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextPropertyNameTypeOrder
// - update stage.TextPropertyNameTypeOrder accordingly
func (textpropertynametype *TextPropertyNameType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextPropertyNameTypes, stage.TextPropertyNameType_stagedOrder, stage.TextPropertyNameType_orderStaged, &stage.TextPropertyNameTypeOrder, stage.TextPropertyNameTypes_mapString, textpropertynametype, order, textpropertynametype.Name)
}

// Unstage removes textpropertynametype off the model stage
func (textpropertynametype *TextPropertyNameType) Unstage(stage *Stage) *TextPropertyNameType {
	__gong__unstage(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_mapString, textpropertynametype, textpropertynametype.Name)
	return textpropertynametype
}

// UnstageVoid removes textpropertynametype off the model stage
func (textpropertynametype *TextPropertyNameType) UnstageVoid(stage *Stage) {
	textpropertynametype.Unstage(stage)
}

func (textpropertynametype *TextPropertyNameType) StageVoid(stage *Stage) {
	textpropertynametype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textpropertynametype *TextPropertyNameType) GetName() (res string) {
	return textpropertynametype.Name
}

// for satisfaction of GongStruct interface
func (textpropertynametype *TextPropertyNameType) SetName(name string) {
	textpropertynametype.Name = name
}

// Stage puts timetype to the model stage
func (timetype *TimeType) Stage(stage *Stage) *TimeType {
	__gong__stage(stage.TimeTypes, stage.TimeType_stagedOrder, stage.TimeType_orderStaged, &stage.TimeTypeOrder, stage.TimeTypes_mapString, timetype, timetype.Name)
	return timetype
}

// StagePreserveOrder puts timetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimeTypeOrder
// - update stage.TimeTypeOrder accordingly
func (timetype *TimeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimeTypes, stage.TimeType_stagedOrder, stage.TimeType_orderStaged, &stage.TimeTypeOrder, stage.TimeTypes_mapString, timetype, order, timetype.Name)
}

// Unstage removes timetype off the model stage
func (timetype *TimeType) Unstage(stage *Stage) *TimeType {
	__gong__unstage(stage.TimeTypes, stage.TimeTypes_mapString, timetype, timetype.Name)
	return timetype
}

// UnstageVoid removes timetype off the model stage
func (timetype *TimeType) UnstageVoid(stage *Stage) {
	timetype.Unstage(stage)
}

func (timetype *TimeType) StageVoid(stage *Stage) {
	timetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timetype *TimeType) GetName() (res string) {
	return timetype.Name
}

// for satisfaction of GongStruct interface
func (timetype *TimeType) SetName(name string) {
	timetype.Name = name
}

// Stage puts valangletype to the model stage
func (valangletype *ValAngleType) Stage(stage *Stage) *ValAngleType {
	__gong__stage(stage.ValAngleTypes, stage.ValAngleType_stagedOrder, stage.ValAngleType_orderStaged, &stage.ValAngleTypeOrder, stage.ValAngleTypes_mapString, valangletype, valangletype.Name)
	return valangletype
}

// StagePreserveOrder puts valangletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValAngleTypeOrder
// - update stage.ValAngleTypeOrder accordingly
func (valangletype *ValAngleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValAngleTypes, stage.ValAngleType_stagedOrder, stage.ValAngleType_orderStaged, &stage.ValAngleTypeOrder, stage.ValAngleTypes_mapString, valangletype, order, valangletype.Name)
}

// Unstage removes valangletype off the model stage
func (valangletype *ValAngleType) Unstage(stage *Stage) *ValAngleType {
	__gong__unstage(stage.ValAngleTypes, stage.ValAngleTypes_mapString, valangletype, valangletype.Name)
	return valangletype
}

// UnstageVoid removes valangletype off the model stage
func (valangletype *ValAngleType) UnstageVoid(stage *Stage) {
	valangletype.Unstage(stage)
}

func (valangletype *ValAngleType) StageVoid(stage *Stage) {
	valangletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valangletype *ValAngleType) GetName() (res string) {
	return valangletype.Name
}

// for satisfaction of GongStruct interface
func (valangletype *ValAngleType) SetName(name string) {
	valangletype.Name = name
}

// Stage puts valbearingtype to the model stage
func (valbearingtype *ValBearingType) Stage(stage *Stage) *ValBearingType {
	__gong__stage(stage.ValBearingTypes, stage.ValBearingType_stagedOrder, stage.ValBearingType_orderStaged, &stage.ValBearingTypeOrder, stage.ValBearingTypes_mapString, valbearingtype, valbearingtype.Name)
	return valbearingtype
}

// StagePreserveOrder puts valbearingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValBearingTypeOrder
// - update stage.ValBearingTypeOrder accordingly
func (valbearingtype *ValBearingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValBearingTypes, stage.ValBearingType_stagedOrder, stage.ValBearingType_orderStaged, &stage.ValBearingTypeOrder, stage.ValBearingTypes_mapString, valbearingtype, order, valbearingtype.Name)
}

// Unstage removes valbearingtype off the model stage
func (valbearingtype *ValBearingType) Unstage(stage *Stage) *ValBearingType {
	__gong__unstage(stage.ValBearingTypes, stage.ValBearingTypes_mapString, valbearingtype, valbearingtype.Name)
	return valbearingtype
}

// UnstageVoid removes valbearingtype off the model stage
func (valbearingtype *ValBearingType) UnstageVoid(stage *Stage) {
	valbearingtype.Unstage(stage)
}

func (valbearingtype *ValBearingType) StageVoid(stage *Stage) {
	valbearingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valbearingtype *ValBearingType) GetName() (res string) {
	return valbearingtype.Name
}

// for satisfaction of GongStruct interface
func (valbearingtype *ValBearingType) SetName(name string) {
	valbearingtype.Name = name
}

// Stage puts valdistancesignedtype to the model stage
func (valdistancesignedtype *ValDistanceSignedType) Stage(stage *Stage) *ValDistanceSignedType {
	__gong__stage(stage.ValDistanceSignedTypes, stage.ValDistanceSignedType_stagedOrder, stage.ValDistanceSignedType_orderStaged, &stage.ValDistanceSignedTypeOrder, stage.ValDistanceSignedTypes_mapString, valdistancesignedtype, valdistancesignedtype.Name)
	return valdistancesignedtype
}

// StagePreserveOrder puts valdistancesignedtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValDistanceSignedTypeOrder
// - update stage.ValDistanceSignedTypeOrder accordingly
func (valdistancesignedtype *ValDistanceSignedType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValDistanceSignedTypes, stage.ValDistanceSignedType_stagedOrder, stage.ValDistanceSignedType_orderStaged, &stage.ValDistanceSignedTypeOrder, stage.ValDistanceSignedTypes_mapString, valdistancesignedtype, order, valdistancesignedtype.Name)
}

// Unstage removes valdistancesignedtype off the model stage
func (valdistancesignedtype *ValDistanceSignedType) Unstage(stage *Stage) *ValDistanceSignedType {
	__gong__unstage(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_mapString, valdistancesignedtype, valdistancesignedtype.Name)
	return valdistancesignedtype
}

// UnstageVoid removes valdistancesignedtype off the model stage
func (valdistancesignedtype *ValDistanceSignedType) UnstageVoid(stage *Stage) {
	valdistancesignedtype.Unstage(stage)
}

func (valdistancesignedtype *ValDistanceSignedType) StageVoid(stage *Stage) {
	valdistancesignedtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valdistancesignedtype *ValDistanceSignedType) GetName() (res string) {
	return valdistancesignedtype.Name
}

// for satisfaction of GongStruct interface
func (valdistancesignedtype *ValDistanceSignedType) SetName(name string) {
	valdistancesignedtype.Name = name
}

// Stage puts valdistancetype to the model stage
func (valdistancetype *ValDistanceType) Stage(stage *Stage) *ValDistanceType {
	__gong__stage(stage.ValDistanceTypes, stage.ValDistanceType_stagedOrder, stage.ValDistanceType_orderStaged, &stage.ValDistanceTypeOrder, stage.ValDistanceTypes_mapString, valdistancetype, valdistancetype.Name)
	return valdistancetype
}

// StagePreserveOrder puts valdistancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValDistanceTypeOrder
// - update stage.ValDistanceTypeOrder accordingly
func (valdistancetype *ValDistanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValDistanceTypes, stage.ValDistanceType_stagedOrder, stage.ValDistanceType_orderStaged, &stage.ValDistanceTypeOrder, stage.ValDistanceTypes_mapString, valdistancetype, order, valdistancetype.Name)
}

// Unstage removes valdistancetype off the model stage
func (valdistancetype *ValDistanceType) Unstage(stage *Stage) *ValDistanceType {
	__gong__unstage(stage.ValDistanceTypes, stage.ValDistanceTypes_mapString, valdistancetype, valdistancetype.Name)
	return valdistancetype
}

// UnstageVoid removes valdistancetype off the model stage
func (valdistancetype *ValDistanceType) UnstageVoid(stage *Stage) {
	valdistancetype.Unstage(stage)
}

func (valdistancetype *ValDistanceType) StageVoid(stage *Stage) {
	valdistancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valdistancetype *ValDistanceType) GetName() (res string) {
	return valdistancetype.Name
}

// for satisfaction of GongStruct interface
func (valdistancetype *ValDistanceType) SetName(name string) {
	valdistancetype.Name = name
}

// Stage puts valdistanceverticaltype to the model stage
func (valdistanceverticaltype *ValDistanceVerticalType) Stage(stage *Stage) *ValDistanceVerticalType {
	__gong__stage(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalType_stagedOrder, stage.ValDistanceVerticalType_orderStaged, &stage.ValDistanceVerticalTypeOrder, stage.ValDistanceVerticalTypes_mapString, valdistanceverticaltype, valdistanceverticaltype.Name)
	return valdistanceverticaltype
}

// StagePreserveOrder puts valdistanceverticaltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValDistanceVerticalTypeOrder
// - update stage.ValDistanceVerticalTypeOrder accordingly
func (valdistanceverticaltype *ValDistanceVerticalType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalType_stagedOrder, stage.ValDistanceVerticalType_orderStaged, &stage.ValDistanceVerticalTypeOrder, stage.ValDistanceVerticalTypes_mapString, valdistanceverticaltype, order, valdistanceverticaltype.Name)
}

// Unstage removes valdistanceverticaltype off the model stage
func (valdistanceverticaltype *ValDistanceVerticalType) Unstage(stage *Stage) *ValDistanceVerticalType {
	__gong__unstage(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_mapString, valdistanceverticaltype, valdistanceverticaltype.Name)
	return valdistanceverticaltype
}

// UnstageVoid removes valdistanceverticaltype off the model stage
func (valdistanceverticaltype *ValDistanceVerticalType) UnstageVoid(stage *Stage) {
	valdistanceverticaltype.Unstage(stage)
}

func (valdistanceverticaltype *ValDistanceVerticalType) StageVoid(stage *Stage) {
	valdistanceverticaltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valdistanceverticaltype *ValDistanceVerticalType) GetName() (res string) {
	return valdistanceverticaltype.Name
}

// for satisfaction of GongStruct interface
func (valdistanceverticaltype *ValDistanceVerticalType) SetName(name string) {
	valdistanceverticaltype.Name = name
}

// Stage puts valdurationtype to the model stage
func (valdurationtype *ValDurationType) Stage(stage *Stage) *ValDurationType {
	__gong__stage(stage.ValDurationTypes, stage.ValDurationType_stagedOrder, stage.ValDurationType_orderStaged, &stage.ValDurationTypeOrder, stage.ValDurationTypes_mapString, valdurationtype, valdurationtype.Name)
	return valdurationtype
}

// StagePreserveOrder puts valdurationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValDurationTypeOrder
// - update stage.ValDurationTypeOrder accordingly
func (valdurationtype *ValDurationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValDurationTypes, stage.ValDurationType_stagedOrder, stage.ValDurationType_orderStaged, &stage.ValDurationTypeOrder, stage.ValDurationTypes_mapString, valdurationtype, order, valdurationtype.Name)
}

// Unstage removes valdurationtype off the model stage
func (valdurationtype *ValDurationType) Unstage(stage *Stage) *ValDurationType {
	__gong__unstage(stage.ValDurationTypes, stage.ValDurationTypes_mapString, valdurationtype, valdurationtype.Name)
	return valdurationtype
}

// UnstageVoid removes valdurationtype off the model stage
func (valdurationtype *ValDurationType) UnstageVoid(stage *Stage) {
	valdurationtype.Unstage(stage)
}

func (valdurationtype *ValDurationType) StageVoid(stage *Stage) {
	valdurationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valdurationtype *ValDurationType) GetName() (res string) {
	return valdurationtype.Name
}

// for satisfaction of GongStruct interface
func (valdurationtype *ValDurationType) SetName(name string) {
	valdurationtype.Name = name
}

// Stage puts valslopetype to the model stage
func (valslopetype *ValSlopeType) Stage(stage *Stage) *ValSlopeType {
	__gong__stage(stage.ValSlopeTypes, stage.ValSlopeType_stagedOrder, stage.ValSlopeType_orderStaged, &stage.ValSlopeTypeOrder, stage.ValSlopeTypes_mapString, valslopetype, valslopetype.Name)
	return valslopetype
}

// StagePreserveOrder puts valslopetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValSlopeTypeOrder
// - update stage.ValSlopeTypeOrder accordingly
func (valslopetype *ValSlopeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValSlopeTypes, stage.ValSlopeType_stagedOrder, stage.ValSlopeType_orderStaged, &stage.ValSlopeTypeOrder, stage.ValSlopeTypes_mapString, valslopetype, order, valslopetype.Name)
}

// Unstage removes valslopetype off the model stage
func (valslopetype *ValSlopeType) Unstage(stage *Stage) *ValSlopeType {
	__gong__unstage(stage.ValSlopeTypes, stage.ValSlopeTypes_mapString, valslopetype, valslopetype.Name)
	return valslopetype
}

// UnstageVoid removes valslopetype off the model stage
func (valslopetype *ValSlopeType) UnstageVoid(stage *Stage) {
	valslopetype.Unstage(stage)
}

func (valslopetype *ValSlopeType) StageVoid(stage *Stage) {
	valslopetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valslopetype *ValSlopeType) GetName() (res string) {
	return valslopetype.Name
}

// for satisfaction of GongStruct interface
func (valslopetype *ValSlopeType) SetName(name string) {
	valslopetype.Name = name
}

// Stage puts valspeedtype to the model stage
func (valspeedtype *ValSpeedType) Stage(stage *Stage) *ValSpeedType {
	__gong__stage(stage.ValSpeedTypes, stage.ValSpeedType_stagedOrder, stage.ValSpeedType_orderStaged, &stage.ValSpeedTypeOrder, stage.ValSpeedTypes_mapString, valspeedtype, valspeedtype.Name)
	return valspeedtype
}

// StagePreserveOrder puts valspeedtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValSpeedTypeOrder
// - update stage.ValSpeedTypeOrder accordingly
func (valspeedtype *ValSpeedType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValSpeedTypes, stage.ValSpeedType_stagedOrder, stage.ValSpeedType_orderStaged, &stage.ValSpeedTypeOrder, stage.ValSpeedTypes_mapString, valspeedtype, order, valspeedtype.Name)
}

// Unstage removes valspeedtype off the model stage
func (valspeedtype *ValSpeedType) Unstage(stage *Stage) *ValSpeedType {
	__gong__unstage(stage.ValSpeedTypes, stage.ValSpeedTypes_mapString, valspeedtype, valspeedtype.Name)
	return valspeedtype
}

// UnstageVoid removes valspeedtype off the model stage
func (valspeedtype *ValSpeedType) UnstageVoid(stage *Stage) {
	valspeedtype.Unstage(stage)
}

func (valspeedtype *ValSpeedType) StageVoid(stage *Stage) {
	valspeedtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valspeedtype *ValSpeedType) GetName() (res string) {
	return valspeedtype.Name
}

// for satisfaction of GongStruct interface
func (valspeedtype *ValSpeedType) SetName(name string) {
	valspeedtype.Name = name
}

// Stage puts valweighttype to the model stage
func (valweighttype *ValWeightType) Stage(stage *Stage) *ValWeightType {
	__gong__stage(stage.ValWeightTypes, stage.ValWeightType_stagedOrder, stage.ValWeightType_orderStaged, &stage.ValWeightTypeOrder, stage.ValWeightTypes_mapString, valweighttype, valweighttype.Name)
	return valweighttype
}

// StagePreserveOrder puts valweighttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValWeightTypeOrder
// - update stage.ValWeightTypeOrder accordingly
func (valweighttype *ValWeightType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValWeightTypes, stage.ValWeightType_stagedOrder, stage.ValWeightType_orderStaged, &stage.ValWeightTypeOrder, stage.ValWeightTypes_mapString, valweighttype, order, valweighttype.Name)
}

// Unstage removes valweighttype off the model stage
func (valweighttype *ValWeightType) Unstage(stage *Stage) *ValWeightType {
	__gong__unstage(stage.ValWeightTypes, stage.ValWeightTypes_mapString, valweighttype, valweighttype.Name)
	return valweighttype
}

// UnstageVoid removes valweighttype off the model stage
func (valweighttype *ValWeightType) UnstageVoid(stage *Stage) {
	valweighttype.Unstage(stage)
}

func (valweighttype *ValWeightType) StageVoid(stage *Stage) {
	valweighttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valweighttype *ValWeightType) GetName() (res string) {
	return valweighttype.Name
}

// for satisfaction of GongStruct interface
func (valweighttype *ValWeightType) SetName(name string) {
	valweighttype.Name = name
}

// Stage puts xhtmltype to the model stage
func (xhtmltype *XHTMLType) Stage(stage *Stage) *XHTMLType {
	__gong__stage(stage.XHTMLTypes, stage.XHTMLType_stagedOrder, stage.XHTMLType_orderStaged, &stage.XHTMLTypeOrder, stage.XHTMLTypes_mapString, xhtmltype, xhtmltype.Name)
	return xhtmltype
}

// StagePreserveOrder puts xhtmltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XHTMLTypeOrder
// - update stage.XHTMLTypeOrder accordingly
func (xhtmltype *XHTMLType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.XHTMLTypes, stage.XHTMLType_stagedOrder, stage.XHTMLType_orderStaged, &stage.XHTMLTypeOrder, stage.XHTMLTypes_mapString, xhtmltype, order, xhtmltype.Name)
}

// Unstage removes xhtmltype off the model stage
func (xhtmltype *XHTMLType) Unstage(stage *Stage) *XHTMLType {
	__gong__unstage(stage.XHTMLTypes, stage.XHTMLTypes_mapString, xhtmltype, xhtmltype.Name)
	return xhtmltype
}

// UnstageVoid removes xhtmltype off the model stage
func (xhtmltype *XHTMLType) UnstageVoid(stage *Stage) {
	xhtmltype.Unstage(stage)
}

func (xhtmltype *XHTMLType) StageVoid(stage *Stage) {
	xhtmltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xhtmltype *XHTMLType) GetName() (res string) {
	return xhtmltype.Name
}

// for satisfaction of GongStruct interface
func (xhtmltype *XHTMLType) SetName(name string) {
	xhtmltype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeATCReportingTypes, &stage.CodeATCReportingTypes_mapString, &stage.CodeATCReportingType_stagedOrder, &stage.CodeATCReportingTypeOrder)

	__gong__resetStageType(&stage.CodeAerialRefuellingPointTypes, &stage.CodeAerialRefuellingPointTypes_mapString, &stage.CodeAerialRefuellingPointType_stagedOrder, &stage.CodeAerialRefuellingPointTypeOrder)

	__gong__resetStageType(&stage.CodeAerialRefuellingPrefixTypes, &stage.CodeAerialRefuellingPrefixTypes_mapString, &stage.CodeAerialRefuellingPrefixType_stagedOrder, &stage.CodeAerialRefuellingPrefixTypeOrder)

	__gong__resetStageType(&stage.CodeAerialRefuellingTypes, &stage.CodeAerialRefuellingTypes_mapString, &stage.CodeAerialRefuellingType_stagedOrder, &stage.CodeAerialRefuellingTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString, &stage.CodeAircraftCategoryType_stagedOrder, &stage.CodeAircraftCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString, &stage.CodeAircraftEngineNumberType_stagedOrder, &stage.CodeAircraftEngineNumberTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString, &stage.CodeAircraftEngineType_stagedOrder, &stage.CodeAircraftEngineTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString, &stage.CodeAircraftICAOType_stagedOrder, &stage.CodeAircraftICAOTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString, &stage.CodeAircraftType_stagedOrder, &stage.CodeAircraftTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString, &stage.CodeAircraftWingspanClassType_stagedOrder, &stage.CodeAircraftWingspanClassTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_mapString, &stage.CodeAltitudeAdjustmentType_stagedOrder, &stage.CodeAltitudeAdjustmentTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString, &stage.CodeAltitudeUseType_stagedOrder, &stage.CodeAltitudeUseTypeOrder)

	__gong__resetStageType(&stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_mapString, &stage.CodeAuthorityType_stagedOrder, &stage.CodeAuthorityTypeOrder)

	__gong__resetStageType(&stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_mapString, &stage.CodeCardinalDirectionType_stagedOrder, &stage.CodeCardinalDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString, &stage.CodeCommunicationModeType_stagedOrder, &stage.CodeCommunicationModeTypeOrder)

	__gong__resetStageType(&stage.CodeComparisonTypes, &stage.CodeComparisonTypes_mapString, &stage.CodeComparisonType_stagedOrder, &stage.CodeComparisonTypeOrder)

	__gong__resetStageType(&stage.CodeCourseTypes, &stage.CodeCourseTypes_mapString, &stage.CodeCourseType_stagedOrder, &stage.CodeCourseTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString, &stage.CodeDirectionTurnType_stagedOrder, &stage.CodeDirectionTurnTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionTypes, &stage.CodeDirectionTypes_mapString, &stage.CodeDirectionType_stagedOrder, &stage.CodeDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString, &stage.CodeEquipmentAntiCollisionType_stagedOrder, &stage.CodeEquipmentAntiCollisionTypeOrder)

	__gong__resetStageType(&stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_mapString, &stage.CodeFlightOriginType_stagedOrder, &stage.CodeFlightOriginTypeOrder)

	__gong__resetStageType(&stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_mapString, &stage.CodeFlightPurposeType_stagedOrder, &stage.CodeFlightPurposeTypeOrder)

	__gong__resetStageType(&stage.CodeFlightRestrictionDesignatorTypes, &stage.CodeFlightRestrictionDesignatorTypes_mapString, &stage.CodeFlightRestrictionDesignatorType_stagedOrder, &stage.CodeFlightRestrictionDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeFlightRestrictionTypes, &stage.CodeFlightRestrictionTypes_mapString, &stage.CodeFlightRestrictionType_stagedOrder, &stage.CodeFlightRestrictionTypeOrder)

	__gong__resetStageType(&stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString, &stage.CodeFlightRuleType_stagedOrder, &stage.CodeFlightRuleTypeOrder)

	__gong__resetStageType(&stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_mapString, &stage.CodeFlightStatusType_stagedOrder, &stage.CodeFlightStatusTypeOrder)

	__gong__resetStageType(&stage.CodeFlightTypes, &stage.CodeFlightTypes_mapString, &stage.CodeFlightType_stagedOrder, &stage.CodeFlightTypeOrder)

	__gong__resetStageType(&stage.CodeFlowConditionOperationTypes, &stage.CodeFlowConditionOperationTypes_mapString, &stage.CodeFlowConditionOperationType_stagedOrder, &stage.CodeFlowConditionOperationTypeOrder)

	__gong__resetStageType(&stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_mapString, &stage.CodeFreeFlightType_stagedOrder, &stage.CodeFreeFlightTypeOrder)

	__gong__resetStageType(&stage.CodeLevelTypes, &stage.CodeLevelTypes_mapString, &stage.CodeLevelType_stagedOrder, &stage.CodeLevelTypeOrder)

	__gong__resetStageType(&stage.CodeLocationQualifierTypes, &stage.CodeLocationQualifierTypes_mapString, &stage.CodeLocationQualifierType_stagedOrder, &stage.CodeLocationQualifierTypeOrder)

	__gong__resetStageType(&stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_mapString, &stage.CodeMeteoConditionsType_stagedOrder, &stage.CodeMeteoConditionsTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_mapString, &stage.CodeMilitaryRoutePointType_stagedOrder, &stage.CodeMilitaryRoutePointTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_mapString, &stage.CodeMilitaryStatusType_stagedOrder, &stage.CodeMilitaryStatusTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryTrainingTypes, &stage.CodeMilitaryTrainingTypes_mapString, &stage.CodeMilitaryTrainingType_stagedOrder, &stage.CodeMilitaryTrainingTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString, &stage.CodeNavigationEquipmentType_stagedOrder, &stage.CodeNavigationEquipmentTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString, &stage.CodeNavigationSpecificationType_stagedOrder, &stage.CodeNavigationSpecificationTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString, &stage.CodeObstacleAssessmentSurfaceType_stagedOrder, &stage.CodeObstacleAssessmentSurfaceTypeOrder)

	__gong__resetStageType(&stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_mapString, &stage.CodeObstructionIdSurfaceZoneType_stagedOrder, &stage.CodeObstructionIdSurfaceZoneTypeOrder)

	__gong__resetStageType(&stage.CodeRNPTypes, &stage.CodeRNPTypes_mapString, &stage.CodeRNPType_stagedOrder, &stage.CodeRNPTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_mapString, &stage.CodeRVSMPointRoleType_stagedOrder, &stage.CodeRVSMPointRoleTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString, &stage.CodeRVSMType_stagedOrder, &stage.CodeRVSMTypeOrder)

	__gong__resetStageType(&stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_mapString, &stage.CodeReferenceRoleType_stagedOrder, &stage.CodeReferenceRoleTypeOrder)

	__gong__resetStageType(&stage.CodeRouteAvailabilityTypes, &stage.CodeRouteAvailabilityTypes_mapString, &stage.CodeRouteAvailabilityType_stagedOrder, &stage.CodeRouteAvailabilityTypeOrder)

	__gong__resetStageType(&stage.CodeRouteDesignatorLetterTypes, &stage.CodeRouteDesignatorLetterTypes_mapString, &stage.CodeRouteDesignatorLetterType_stagedOrder, &stage.CodeRouteDesignatorLetterTypeOrder)

	__gong__resetStageType(&stage.CodeRouteDesignatorPrefixTypes, &stage.CodeRouteDesignatorPrefixTypes_mapString, &stage.CodeRouteDesignatorPrefixType_stagedOrder, &stage.CodeRouteDesignatorPrefixTypeOrder)

	__gong__resetStageType(&stage.CodeRouteDesignatorSuffixTypes, &stage.CodeRouteDesignatorSuffixTypes_mapString, &stage.CodeRouteDesignatorSuffixType_stagedOrder, &stage.CodeRouteDesignatorSuffixTypeOrder)

	__gong__resetStageType(&stage.CodeRouteNavigationTypes, &stage.CodeRouteNavigationTypes_mapString, &stage.CodeRouteNavigationType_stagedOrder, &stage.CodeRouteNavigationTypeOrder)

	__gong__resetStageType(&stage.CodeRouteOriginTypes, &stage.CodeRouteOriginTypes_mapString, &stage.CodeRouteOriginType_stagedOrder, &stage.CodeRouteOriginTypeOrder)

	__gong__resetStageType(&stage.CodeRouteSegmentPathTypes, &stage.CodeRouteSegmentPathTypes_mapString, &stage.CodeRouteSegmentPathType_stagedOrder, &stage.CodeRouteSegmentPathTypeOrder)

	__gong__resetStageType(&stage.CodeRouteTypes, &stage.CodeRouteTypes_mapString, &stage.CodeRouteType_stagedOrder, &stage.CodeRouteTypeOrder)

	__gong__resetStageType(&stage.CodeRuleProcedureTitleTypes, &stage.CodeRuleProcedureTitleTypes_mapString, &stage.CodeRuleProcedureTitleType_stagedOrder, &stage.CodeRuleProcedureTitleTypeOrder)

	__gong__resetStageType(&stage.CodeRuleProcedureTypes, &stage.CodeRuleProcedureTypes_mapString, &stage.CodeRuleProcedureType_stagedOrder, &stage.CodeRuleProcedureTypeOrder)

	__gong__resetStageType(&stage.CodeSideTypes, &stage.CodeSideTypes_mapString, &stage.CodeSideType_stagedOrder, &stage.CodeSideTypeOrder)

	__gong__resetStageType(&stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_mapString, &stage.CodeSpeedReferenceType_stagedOrder, &stage.CodeSpeedReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_mapString, &stage.CodeTACANChannelType_stagedOrder, &stage.CodeTACANChannelTypeOrder)

	__gong__resetStageType(&stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString, &stage.CodeTelecomNetworkType_stagedOrder, &stage.CodeTelecomNetworkTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString, &stage.CodeTransponderType_stagedOrder, &stage.CodeTransponderTypeOrder)

	__gong__resetStageType(&stage.CodeUpperAlphaTypes, &stage.CodeUpperAlphaTypes_mapString, &stage.CodeUpperAlphaType_stagedOrder, &stage.CodeUpperAlphaTypeOrder)

	__gong__resetStageType(&stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString, &stage.CodeValueInterpretationType_stagedOrder, &stage.CodeValueInterpretationTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString, &stage.CodeVerticalReferenceType_stagedOrder, &stage.CodeVerticalReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString, &stage.CodeWakeTurbulenceType_stagedOrder, &stage.CodeWakeTurbulenceTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.NoNumberTypes, &stage.NoNumberTypes_mapString, &stage.NoNumberType_stagedOrder, &stage.NoNumberTypeOrder)

	__gong__resetStageType(&stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString, &stage.NoSequenceType_stagedOrder, &stage.NoSequenceTypeOrder)

	__gong__resetStageType(&stage.TextAddressTypes, &stage.TextAddressTypes_mapString, &stage.TextAddressType_stagedOrder, &stage.TextAddressTypeOrder)

	__gong__resetStageType(&stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString, &stage.TextDesignatorType_stagedOrder, &stage.TextDesignatorTypeOrder)

	__gong__resetStageType(&stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString, &stage.TextInstructionType_stagedOrder, &stage.TextInstructionTypeOrder)

	__gong__resetStageType(&stage.TextNameTypes, &stage.TextNameTypes_mapString, &stage.TextNameType_stagedOrder, &stage.TextNameTypeOrder)

	__gong__resetStageType(&stage.TextNoteTypes, &stage.TextNoteTypes_mapString, &stage.TextNoteType_stagedOrder, &stage.TextNoteTypeOrder)

	__gong__resetStageType(&stage.TextPhoneTypes, &stage.TextPhoneTypes_mapString, &stage.TextPhoneType_stagedOrder, &stage.TextPhoneTypeOrder)

	__gong__resetStageType(&stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString, &stage.TextPropertyNameType_stagedOrder, &stage.TextPropertyNameTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.ValAngleTypes, &stage.ValAngleTypes_mapString, &stage.ValAngleType_stagedOrder, &stage.ValAngleTypeOrder)

	__gong__resetStageType(&stage.ValBearingTypes, &stage.ValBearingTypes_mapString, &stage.ValBearingType_stagedOrder, &stage.ValBearingTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValSlopeTypes, &stage.ValSlopeTypes_mapString, &stage.ValSlopeType_stagedOrder, &stage.ValSlopeTypeOrder)

	__gong__resetStageType(&stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString, &stage.ValSpeedType_stagedOrder, &stage.ValSpeedTypeOrder)

	__gong__resetStageType(&stage.ValWeightTypes, &stage.ValWeightTypes_mapString, &stage.ValWeightType_stagedOrder, &stage.ValWeightTypeOrder)

	__gong__resetStageType(&stage.XHTMLTypes, &stage.XHTMLTypes_mapString, &stage.XHTMLType_stagedOrder, &stage.XHTMLTypeOrder)

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
func (CodeATCReportingType) GongGetAssociationName() any {
	return &CodeATCReportingType{
	}
}

func (CodeAerialRefuellingPointType) GongGetAssociationName() any {
	return &CodeAerialRefuellingPointType{
	}
}

func (CodeAerialRefuellingPrefixType) GongGetAssociationName() any {
	return &CodeAerialRefuellingPrefixType{
	}
}

func (CodeAerialRefuellingType) GongGetAssociationName() any {
	return &CodeAerialRefuellingType{
	}
}

func (CodeAircraftCategoryType) GongGetAssociationName() any {
	return &CodeAircraftCategoryType{
	}
}

func (CodeAircraftEngineNumberType) GongGetAssociationName() any {
	return &CodeAircraftEngineNumberType{
	}
}

func (CodeAircraftEngineType) GongGetAssociationName() any {
	return &CodeAircraftEngineType{
	}
}

func (CodeAircraftICAOType) GongGetAssociationName() any {
	return &CodeAircraftICAOType{
	}
}

func (CodeAircraftType) GongGetAssociationName() any {
	return &CodeAircraftType{
	}
}

func (CodeAircraftWingspanClassType) GongGetAssociationName() any {
	return &CodeAircraftWingspanClassType{
	}
}

func (CodeAltitudeAdjustmentType) GongGetAssociationName() any {
	return &CodeAltitudeAdjustmentType{
	}
}

func (CodeAltitudeUseType) GongGetAssociationName() any {
	return &CodeAltitudeUseType{
	}
}

func (CodeAuthorityType) GongGetAssociationName() any {
	return &CodeAuthorityType{
	}
}

func (CodeCardinalDirectionType) GongGetAssociationName() any {
	return &CodeCardinalDirectionType{
	}
}

func (CodeCommunicationModeType) GongGetAssociationName() any {
	return &CodeCommunicationModeType{
	}
}

func (CodeComparisonType) GongGetAssociationName() any {
	return &CodeComparisonType{
	}
}

func (CodeCourseType) GongGetAssociationName() any {
	return &CodeCourseType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeDirectionTurnType) GongGetAssociationName() any {
	return &CodeDirectionTurnType{
	}
}

func (CodeDirectionType) GongGetAssociationName() any {
	return &CodeDirectionType{
	}
}

func (CodeEquipmentAntiCollisionType) GongGetAssociationName() any {
	return &CodeEquipmentAntiCollisionType{
	}
}

func (CodeFlightOriginType) GongGetAssociationName() any {
	return &CodeFlightOriginType{
	}
}

func (CodeFlightPurposeType) GongGetAssociationName() any {
	return &CodeFlightPurposeType{
	}
}

func (CodeFlightRestrictionDesignatorType) GongGetAssociationName() any {
	return &CodeFlightRestrictionDesignatorType{
	}
}

func (CodeFlightRestrictionType) GongGetAssociationName() any {
	return &CodeFlightRestrictionType{
	}
}

func (CodeFlightRuleType) GongGetAssociationName() any {
	return &CodeFlightRuleType{
	}
}

func (CodeFlightStatusType) GongGetAssociationName() any {
	return &CodeFlightStatusType{
	}
}

func (CodeFlightType) GongGetAssociationName() any {
	return &CodeFlightType{
	}
}

func (CodeFlowConditionOperationType) GongGetAssociationName() any {
	return &CodeFlowConditionOperationType{
	}
}

func (CodeFreeFlightType) GongGetAssociationName() any {
	return &CodeFreeFlightType{
	}
}

func (CodeLevelType) GongGetAssociationName() any {
	return &CodeLevelType{
	}
}

func (CodeLocationQualifierType) GongGetAssociationName() any {
	return &CodeLocationQualifierType{
	}
}

func (CodeMeteoConditionsType) GongGetAssociationName() any {
	return &CodeMeteoConditionsType{
	}
}

func (CodeMilitaryRoutePointType) GongGetAssociationName() any {
	return &CodeMilitaryRoutePointType{
	}
}

func (CodeMilitaryStatusType) GongGetAssociationName() any {
	return &CodeMilitaryStatusType{
	}
}

func (CodeMilitaryTrainingType) GongGetAssociationName() any {
	return &CodeMilitaryTrainingType{
	}
}

func (CodeNavigationEquipmentType) GongGetAssociationName() any {
	return &CodeNavigationEquipmentType{
	}
}

func (CodeNavigationSpecificationType) GongGetAssociationName() any {
	return &CodeNavigationSpecificationType{
	}
}

func (CodeNotePurposeType) GongGetAssociationName() any {
	return &CodeNotePurposeType{
	}
}

func (CodeObstacleAssessmentSurfaceType) GongGetAssociationName() any {
	return &CodeObstacleAssessmentSurfaceType{
	}
}

func (CodeObstructionIdSurfaceZoneType) GongGetAssociationName() any {
	return &CodeObstructionIdSurfaceZoneType{
	}
}

func (CodeRNPType) GongGetAssociationName() any {
	return &CodeRNPType{
	}
}

func (CodeRVSMPointRoleType) GongGetAssociationName() any {
	return &CodeRVSMPointRoleType{
	}
}

func (CodeRVSMType) GongGetAssociationName() any {
	return &CodeRVSMType{
	}
}

func (CodeReferenceRoleType) GongGetAssociationName() any {
	return &CodeReferenceRoleType{
	}
}

func (CodeRouteAvailabilityType) GongGetAssociationName() any {
	return &CodeRouteAvailabilityType{
	}
}

func (CodeRouteDesignatorLetterType) GongGetAssociationName() any {
	return &CodeRouteDesignatorLetterType{
	}
}

func (CodeRouteDesignatorPrefixType) GongGetAssociationName() any {
	return &CodeRouteDesignatorPrefixType{
	}
}

func (CodeRouteDesignatorSuffixType) GongGetAssociationName() any {
	return &CodeRouteDesignatorSuffixType{
	}
}

func (CodeRouteNavigationType) GongGetAssociationName() any {
	return &CodeRouteNavigationType{
	}
}

func (CodeRouteOriginType) GongGetAssociationName() any {
	return &CodeRouteOriginType{
	}
}

func (CodeRouteSegmentPathType) GongGetAssociationName() any {
	return &CodeRouteSegmentPathType{
	}
}

func (CodeRouteType) GongGetAssociationName() any {
	return &CodeRouteType{
	}
}

func (CodeRuleProcedureTitleType) GongGetAssociationName() any {
	return &CodeRuleProcedureTitleType{
	}
}

func (CodeRuleProcedureType) GongGetAssociationName() any {
	return &CodeRuleProcedureType{
	}
}

func (CodeSideType) GongGetAssociationName() any {
	return &CodeSideType{
	}
}

func (CodeSpeedReferenceType) GongGetAssociationName() any {
	return &CodeSpeedReferenceType{
	}
}

func (CodeTACANChannelType) GongGetAssociationName() any {
	return &CodeTACANChannelType{
	}
}

func (CodeTelecomNetworkType) GongGetAssociationName() any {
	return &CodeTelecomNetworkType{
	}
}

func (CodeTimeEventCombinationType) GongGetAssociationName() any {
	return &CodeTimeEventCombinationType{
	}
}

func (CodeTimeEventType) GongGetAssociationName() any {
	return &CodeTimeEventType{
	}
}

func (CodeTimeReferenceType) GongGetAssociationName() any {
	return &CodeTimeReferenceType{
	}
}

func (CodeTransponderType) GongGetAssociationName() any {
	return &CodeTransponderType{
	}
}

func (CodeUpperAlphaType) GongGetAssociationName() any {
	return &CodeUpperAlphaType{
	}
}

func (CodeValueInterpretationType) GongGetAssociationName() any {
	return &CodeValueInterpretationType{
	}
}

func (CodeVerticalDatumType) GongGetAssociationName() any {
	return &CodeVerticalDatumType{
	}
}

func (CodeVerticalReferenceType) GongGetAssociationName() any {
	return &CodeVerticalReferenceType{
	}
}

func (CodeWakeTurbulenceType) GongGetAssociationName() any {
	return &CodeWakeTurbulenceType{
	}
}

func (CodeYesNoType) GongGetAssociationName() any {
	return &CodeYesNoType{
	}
}

func (DateMonthDayType) GongGetAssociationName() any {
	return &DateMonthDayType{
	}
}

func (NoNumberType) GongGetAssociationName() any {
	return &NoNumberType{
	}
}

func (NoSequenceType) GongGetAssociationName() any {
	return &NoSequenceType{
	}
}

func (TextAddressType) GongGetAssociationName() any {
	return &TextAddressType{
	}
}

func (TextDesignatorType) GongGetAssociationName() any {
	return &TextDesignatorType{
	}
}

func (TextInstructionType) GongGetAssociationName() any {
	return &TextInstructionType{
	}
}

func (TextNameType) GongGetAssociationName() any {
	return &TextNameType{
	}
}

func (TextNoteType) GongGetAssociationName() any {
	return &TextNoteType{
	}
}

func (TextPhoneType) GongGetAssociationName() any {
	return &TextPhoneType{
	}
}

func (TextPropertyNameType) GongGetAssociationName() any {
	return &TextPropertyNameType{
	}
}

func (TimeType) GongGetAssociationName() any {
	return &TimeType{
	}
}

func (ValAngleType) GongGetAssociationName() any {
	return &ValAngleType{
	}
}

func (ValBearingType) GongGetAssociationName() any {
	return &ValBearingType{
	}
}

func (ValDistanceSignedType) GongGetAssociationName() any {
	return &ValDistanceSignedType{
	}
}

func (ValDistanceType) GongGetAssociationName() any {
	return &ValDistanceType{
	}
}

func (ValDistanceVerticalType) GongGetAssociationName() any {
	return &ValDistanceVerticalType{
	}
}

func (ValDurationType) GongGetAssociationName() any {
	return &ValDurationType{
	}
}

func (ValSlopeType) GongGetAssociationName() any {
	return &ValSlopeType{
	}
}

func (ValSpeedType) GongGetAssociationName() any {
	return &ValSpeedType{
	}
}

func (ValWeightType) GongGetAssociationName() any {
	return &ValWeightType{
	}
}

func (XHTMLType) GongGetAssociationName() any {
	return &XHTMLType{
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
	// reverse maps of direct associations of CodeATCReportingType
	case CodeATCReportingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingPointType
	case CodeAerialRefuellingPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingPrefixType
	case CodeAerialRefuellingPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingType
	case CodeAerialRefuellingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftCategoryType
	case CodeAircraftCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftEngineNumberType
	case CodeAircraftEngineNumberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftEngineType
	case CodeAircraftEngineType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftICAOType
	case CodeAircraftICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftType
	case CodeAircraftType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftWingspanClassType
	case CodeAircraftWingspanClassType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeAdjustmentType
	case CodeAltitudeAdjustmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeUseType
	case CodeAltitudeUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityType
	case CodeAuthorityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCardinalDirectionType
	case CodeCardinalDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeComparisonType
	case CodeComparisonType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCourseType
	case CodeCourseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionTurnType
	case CodeDirectionTurnType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionType
	case CodeDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightOriginType
	case CodeFlightOriginType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightPurposeType
	case CodeFlightPurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRestrictionDesignatorType
	case CodeFlightRestrictionDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRestrictionType
	case CodeFlightRestrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRuleType
	case CodeFlightRuleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightStatusType
	case CodeFlightStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightType
	case CodeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlowConditionOperationType
	case CodeFlowConditionOperationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFreeFlightType
	case CodeFreeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelType
	case CodeLevelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLocationQualifierType
	case CodeLocationQualifierType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMeteoConditionsType
	case CodeMeteoConditionsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryRoutePointType
	case CodeMilitaryRoutePointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryStatusType
	case CodeMilitaryStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryTrainingType
	case CodeMilitaryTrainingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationEquipmentType
	case CodeNavigationEquipmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationSpecificationType
	case CodeNavigationSpecificationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAssessmentSurfaceType
	case CodeObstacleAssessmentSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstructionIdSurfaceZoneType
	case CodeObstructionIdSurfaceZoneType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRNPType
	case CodeRNPType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMPointRoleType
	case CodeRVSMPointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReferenceRoleType
	case CodeReferenceRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteAvailabilityType
	case CodeRouteAvailabilityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorLetterType
	case CodeRouteDesignatorLetterType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorPrefixType
	case CodeRouteDesignatorPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorSuffixType
	case CodeRouteDesignatorSuffixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteNavigationType
	case CodeRouteNavigationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteOriginType
	case CodeRouteOriginType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteSegmentPathType
	case CodeRouteSegmentPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteType
	case CodeRouteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRuleProcedureTitleType
	case CodeRuleProcedureTitleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRuleProcedureType
	case CodeRuleProcedureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSideType
	case CodeSideType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpeedReferenceType
	case CodeSpeedReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTACANChannelType
	case CodeTACANChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTelecomNetworkType
	case CodeTelecomNetworkType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeEventCombinationType
	case CodeTimeEventCombinationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeEventType
	case CodeTimeEventType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeReferenceType
	case CodeTimeReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTransponderType
	case CodeTransponderType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeUpperAlphaType
	case CodeUpperAlphaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeValueInterpretationType
	case CodeValueInterpretationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalReferenceType
	case CodeVerticalReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWakeTurbulenceType
	case CodeWakeTurbulenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeYesNoType
	case CodeYesNoType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DateMonthDayType
	case DateMonthDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoNumberType
	case NoNumberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoSequenceType
	case NoSequenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextAddressType
	case TextAddressType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextDesignatorType
	case TextDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextInstructionType
	case TextInstructionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextNameType
	case TextNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextNoteType
	case TextNoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextPhoneType
	case TextPhoneType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextPropertyNameType
	case TextPropertyNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValAngleType
	case ValAngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValBearingType
	case ValBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceSignedType
	case ValDistanceSignedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceType
	case ValDistanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceVerticalType
	case ValDistanceVerticalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDurationType
	case ValDurationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValSlopeType
	case ValSlopeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValSpeedType
	case ValSpeedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValWeightType
	case ValWeightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTMLType
	case XHTMLType:
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
	// reverse maps of direct associations of CodeATCReportingType
	case CodeATCReportingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingPointType
	case CodeAerialRefuellingPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingPrefixType
	case CodeAerialRefuellingPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAerialRefuellingType
	case CodeAerialRefuellingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftCategoryType
	case CodeAircraftCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftEngineNumberType
	case CodeAircraftEngineNumberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftEngineType
	case CodeAircraftEngineType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftICAOType
	case CodeAircraftICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftType
	case CodeAircraftType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAircraftWingspanClassType
	case CodeAircraftWingspanClassType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeAdjustmentType
	case CodeAltitudeAdjustmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeUseType
	case CodeAltitudeUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityType
	case CodeAuthorityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCardinalDirectionType
	case CodeCardinalDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeComparisonType
	case CodeComparisonType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCourseType
	case CodeCourseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionTurnType
	case CodeDirectionTurnType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionType
	case CodeDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightOriginType
	case CodeFlightOriginType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightPurposeType
	case CodeFlightPurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRestrictionDesignatorType
	case CodeFlightRestrictionDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRestrictionType
	case CodeFlightRestrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRuleType
	case CodeFlightRuleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightStatusType
	case CodeFlightStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightType
	case CodeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlowConditionOperationType
	case CodeFlowConditionOperationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFreeFlightType
	case CodeFreeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelType
	case CodeLevelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLocationQualifierType
	case CodeLocationQualifierType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMeteoConditionsType
	case CodeMeteoConditionsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryRoutePointType
	case CodeMilitaryRoutePointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryStatusType
	case CodeMilitaryStatusType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryTrainingType
	case CodeMilitaryTrainingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationEquipmentType
	case CodeNavigationEquipmentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationSpecificationType
	case CodeNavigationSpecificationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAssessmentSurfaceType
	case CodeObstacleAssessmentSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstructionIdSurfaceZoneType
	case CodeObstructionIdSurfaceZoneType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRNPType
	case CodeRNPType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMPointRoleType
	case CodeRVSMPointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReferenceRoleType
	case CodeReferenceRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteAvailabilityType
	case CodeRouteAvailabilityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorLetterType
	case CodeRouteDesignatorLetterType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorPrefixType
	case CodeRouteDesignatorPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteDesignatorSuffixType
	case CodeRouteDesignatorSuffixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteNavigationType
	case CodeRouteNavigationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteOriginType
	case CodeRouteOriginType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteSegmentPathType
	case CodeRouteSegmentPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteType
	case CodeRouteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRuleProcedureTitleType
	case CodeRuleProcedureTitleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRuleProcedureType
	case CodeRuleProcedureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSideType
	case CodeSideType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpeedReferenceType
	case CodeSpeedReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTACANChannelType
	case CodeTACANChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTelecomNetworkType
	case CodeTelecomNetworkType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeEventCombinationType
	case CodeTimeEventCombinationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeEventType
	case CodeTimeEventType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTimeReferenceType
	case CodeTimeReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTransponderType
	case CodeTransponderType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeUpperAlphaType
	case CodeUpperAlphaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeValueInterpretationType
	case CodeValueInterpretationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalReferenceType
	case CodeVerticalReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWakeTurbulenceType
	case CodeWakeTurbulenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeYesNoType
	case CodeYesNoType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DateMonthDayType
	case DateMonthDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoNumberType
	case NoNumberType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoSequenceType
	case NoSequenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextAddressType
	case TextAddressType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextDesignatorType
	case TextDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextInstructionType
	case TextInstructionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextNameType
	case TextNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextNoteType
	case TextNoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextPhoneType
	case TextPhoneType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextPropertyNameType
	case TextPropertyNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValAngleType
	case ValAngleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValBearingType
	case ValBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceSignedType
	case ValDistanceSignedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceType
	case ValDistanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDistanceVerticalType
	case ValDistanceVerticalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValDurationType
	case ValDurationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValSlopeType
	case ValSlopeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValSpeedType
	case ValSpeedType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValWeightType
	case ValWeightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTMLType
	case XHTMLType:
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
func (*CodeATCReportingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAerialRefuellingPointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAerialRefuellingPrefixType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAerialRefuellingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftCategoryType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftEngineNumberType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftEngineType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftICAOType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftWingspanClassType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAltitudeAdjustmentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAltitudeUseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAuthorityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCardinalDirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationModeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeComparisonType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCourseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDirectionTurnType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeEquipmentAntiCollisionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightOriginType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightPurposeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightRestrictionDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightRestrictionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightRuleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightStatusType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlowConditionOperationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFreeFlightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLevelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLocationQualifierType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMeteoConditionsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryRoutePointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryStatusType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryTrainingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavigationEquipmentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavigationSpecificationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNotePurposeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeObstacleAssessmentSurfaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeObstructionIdSurfaceZoneType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRNPType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRVSMPointRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRVSMType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeReferenceRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteAvailabilityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteDesignatorLetterType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteDesignatorPrefixType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteDesignatorSuffixType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteNavigationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteOriginType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteSegmentPathType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRuleProcedureTitleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRuleProcedureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSideType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSpeedReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTACANChannelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTelecomNetworkType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTimeEventCombinationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTimeEventType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTimeReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTransponderType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeUpperAlphaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeValueInterpretationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalDatumType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeWakeTurbulenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeYesNoType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DateMonthDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NoNumberType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NoSequenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextAddressType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextInstructionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextNameType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextNoteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextPhoneType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextPropertyNameType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValAngleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValBearingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValDistanceSignedType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValDistanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValDistanceVerticalType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValDurationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValSlopeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValSpeedType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValWeightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*XHTMLType) GongGetReverseFields() []GongReverseField {
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
func (codeatcreportingtype *CodeATCReportingType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircrafttype *CodeAircraftType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeauthoritytype *CodeAuthorityType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codecomparisontype *CodeComparisonType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codecoursetype *CodeCourseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codedaytype *CodeDayType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codedirectiontype *CodeDirectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightorigintype *CodeFlightOriginType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightruletype *CodeFlightRuleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflightstatustype *CodeFlightStatusType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflighttype *CodeFlightType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codefreeflighttype *CodeFreeFlightType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeleveltype *CodeLevelType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codenotepurposetype *CodeNotePurposeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codernptype *CodeRNPType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codervsmtype *CodeRVSMType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderouteorigintype *CodeRouteOriginType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderoutetype *CodeRouteType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codesidetype *CodeSideType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetimeeventtype *CodeTimeEventType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetranspondertype *CodeTransponderType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeyesnotype *CodeYesNoType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (datemonthdaytype *DateMonthDayType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (nonumbertype *NoNumberType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (nosequencetype *NoSequenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (textaddresstype *TextAddressType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textdesignatortype *TextDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textinstructiontype *TextInstructionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textnametype *TextNameType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textnotetype *TextNoteType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Lang",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textphonetype *TextPhoneType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (textpropertynametype *TextPropertyNameType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (timetype *TimeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valangletype *ValAngleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valbearingtype *ValBearingType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valdistancetype *ValDistanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valdurationtype *ValDurationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valslopetype *ValSlopeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valspeedtype *ValSpeedType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valweighttype *ValWeightType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (xhtmltype *XHTMLType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
func (codeatcreportingtype *CodeATCReportingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeatcreportingtype.Name
	case "NilReason":
		res.valueString = codeatcreportingtype.NilReason
	case "EnclosedText":
		res.valueString = codeatcreportingtype.EnclosedText
	}
	return
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaerialrefuellingpointtype.Name
	case "NilReason":
		res.valueString = codeaerialrefuellingpointtype.NilReason
	case "EnclosedText":
		res.valueString = codeaerialrefuellingpointtype.EnclosedText
	}
	return
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaerialrefuellingprefixtype.Name
	case "NilReason":
		res.valueString = codeaerialrefuellingprefixtype.NilReason
	case "EnclosedText":
		res.valueString = codeaerialrefuellingprefixtype.EnclosedText
	}
	return
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaerialrefuellingtype.Name
	case "NilReason":
		res.valueString = codeaerialrefuellingtype.NilReason
	case "EnclosedText":
		res.valueString = codeaerialrefuellingtype.EnclosedText
	}
	return
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftcategorytype.Name
	case "NilReason":
		res.valueString = codeaircraftcategorytype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftcategorytype.EnclosedText
	}
	return
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftenginenumbertype.Name
	case "NilReason":
		res.valueString = codeaircraftenginenumbertype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftenginenumbertype.EnclosedText
	}
	return
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftenginetype.Name
	case "NilReason":
		res.valueString = codeaircraftenginetype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftenginetype.EnclosedText
	}
	return
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircrafticaotype.Name
	case "NilReason":
		res.valueString = codeaircrafticaotype.NilReason
	case "EnclosedText":
		res.valueString = codeaircrafticaotype.EnclosedText
	}
	return
}

func (codeaircrafttype *CodeAircraftType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircrafttype.Name
	case "NilReason":
		res.valueString = codeaircrafttype.NilReason
	case "EnclosedText":
		res.valueString = codeaircrafttype.EnclosedText
	}
	return
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftwingspanclasstype.Name
	case "NilReason":
		res.valueString = codeaircraftwingspanclasstype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftwingspanclasstype.EnclosedText
	}
	return
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codealtitudeadjustmenttype.Name
	case "NilReason":
		res.valueString = codealtitudeadjustmenttype.NilReason
	case "EnclosedText":
		res.valueString = codealtitudeadjustmenttype.EnclosedText
	}
	return
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codealtitudeusetype.Name
	case "NilReason":
		res.valueString = codealtitudeusetype.NilReason
	case "EnclosedText":
		res.valueString = codealtitudeusetype.EnclosedText
	}
	return
}

func (codeauthoritytype *CodeAuthorityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeauthoritytype.Name
	case "NilReason":
		res.valueString = codeauthoritytype.NilReason
	case "EnclosedText":
		res.valueString = codeauthoritytype.EnclosedText
	}
	return
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecardinaldirectiontype.Name
	case "NilReason":
		res.valueString = codecardinaldirectiontype.NilReason
	case "EnclosedText":
		res.valueString = codecardinaldirectiontype.EnclosedText
	}
	return
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecommunicationmodetype.Name
	case "NilReason":
		res.valueString = codecommunicationmodetype.NilReason
	case "EnclosedText":
		res.valueString = codecommunicationmodetype.EnclosedText
	}
	return
}

func (codecomparisontype *CodeComparisonType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecomparisontype.Name
	case "NilReason":
		res.valueString = codecomparisontype.NilReason
	case "EnclosedText":
		res.valueString = codecomparisontype.EnclosedText
	}
	return
}

func (codecoursetype *CodeCourseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecoursetype.Name
	case "NilReason":
		res.valueString = codecoursetype.NilReason
	case "EnclosedText":
		res.valueString = codecoursetype.EnclosedText
	}
	return
}

func (codedaytype *CodeDayType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedaytype.Name
	case "NilReason":
		res.valueString = codedaytype.NilReason
	case "EnclosedText":
		res.valueString = codedaytype.EnclosedText
	}
	return
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedirectionturntype.Name
	case "NilReason":
		res.valueString = codedirectionturntype.NilReason
	case "EnclosedText":
		res.valueString = codedirectionturntype.EnclosedText
	}
	return
}

func (codedirectiontype *CodeDirectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedirectiontype.Name
	case "NilReason":
		res.valueString = codedirectiontype.NilReason
	case "EnclosedText":
		res.valueString = codedirectiontype.EnclosedText
	}
	return
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeequipmentanticollisiontype.Name
	case "NilReason":
		res.valueString = codeequipmentanticollisiontype.NilReason
	case "EnclosedText":
		res.valueString = codeequipmentanticollisiontype.EnclosedText
	}
	return
}

func (codeflightorigintype *CodeFlightOriginType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightorigintype.Name
	case "NilReason":
		res.valueString = codeflightorigintype.NilReason
	case "EnclosedText":
		res.valueString = codeflightorigintype.EnclosedText
	}
	return
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightpurposetype.Name
	case "NilReason":
		res.valueString = codeflightpurposetype.NilReason
	case "EnclosedText":
		res.valueString = codeflightpurposetype.EnclosedText
	}
	return
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightrestrictiondesignatortype.Name
	case "NilReason":
		res.valueString = codeflightrestrictiondesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codeflightrestrictiondesignatortype.EnclosedText
	}
	return
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightrestrictiontype.Name
	case "NilReason":
		res.valueString = codeflightrestrictiontype.NilReason
	case "EnclosedText":
		res.valueString = codeflightrestrictiontype.EnclosedText
	}
	return
}

func (codeflightruletype *CodeFlightRuleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightruletype.Name
	case "NilReason":
		res.valueString = codeflightruletype.NilReason
	case "EnclosedText":
		res.valueString = codeflightruletype.EnclosedText
	}
	return
}

func (codeflightstatustype *CodeFlightStatusType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightstatustype.Name
	case "NilReason":
		res.valueString = codeflightstatustype.NilReason
	case "EnclosedText":
		res.valueString = codeflightstatustype.EnclosedText
	}
	return
}

func (codeflighttype *CodeFlightType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflighttype.Name
	case "NilReason":
		res.valueString = codeflighttype.NilReason
	case "EnclosedText":
		res.valueString = codeflighttype.EnclosedText
	}
	return
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflowconditionoperationtype.Name
	case "NilReason":
		res.valueString = codeflowconditionoperationtype.NilReason
	case "EnclosedText":
		res.valueString = codeflowconditionoperationtype.EnclosedText
	}
	return
}

func (codefreeflighttype *CodeFreeFlightType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefreeflighttype.Name
	case "NilReason":
		res.valueString = codefreeflighttype.NilReason
	case "EnclosedText":
		res.valueString = codefreeflighttype.EnclosedText
	}
	return
}

func (codeleveltype *CodeLevelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeleveltype.Name
	case "NilReason":
		res.valueString = codeleveltype.NilReason
	case "EnclosedText":
		res.valueString = codeleveltype.EnclosedText
	}
	return
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelocationqualifiertype.Name
	case "NilReason":
		res.valueString = codelocationqualifiertype.NilReason
	case "EnclosedText":
		res.valueString = codelocationqualifiertype.EnclosedText
	}
	return
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemeteoconditionstype.Name
	case "NilReason":
		res.valueString = codemeteoconditionstype.NilReason
	case "EnclosedText":
		res.valueString = codemeteoconditionstype.EnclosedText
	}
	return
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemilitaryroutepointtype.Name
	case "NilReason":
		res.valueString = codemilitaryroutepointtype.NilReason
	case "EnclosedText":
		res.valueString = codemilitaryroutepointtype.EnclosedText
	}
	return
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemilitarystatustype.Name
	case "NilReason":
		res.valueString = codemilitarystatustype.NilReason
	case "EnclosedText":
		res.valueString = codemilitarystatustype.EnclosedText
	}
	return
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemilitarytrainingtype.Name
	case "NilReason":
		res.valueString = codemilitarytrainingtype.NilReason
	case "EnclosedText":
		res.valueString = codemilitarytrainingtype.EnclosedText
	}
	return
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavigationequipmenttype.Name
	case "NilReason":
		res.valueString = codenavigationequipmenttype.NilReason
	case "EnclosedText":
		res.valueString = codenavigationequipmenttype.EnclosedText
	}
	return
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavigationspecificationtype.Name
	case "NilReason":
		res.valueString = codenavigationspecificationtype.NilReason
	case "EnclosedText":
		res.valueString = codenavigationspecificationtype.EnclosedText
	}
	return
}

func (codenotepurposetype *CodeNotePurposeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenotepurposetype.Name
	case "NilReason":
		res.valueString = codenotepurposetype.NilReason
	case "EnclosedText":
		res.valueString = codenotepurposetype.EnclosedText
	}
	return
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeobstacleassessmentsurfacetype.Name
	case "NilReason":
		res.valueString = codeobstacleassessmentsurfacetype.NilReason
	case "EnclosedText":
		res.valueString = codeobstacleassessmentsurfacetype.EnclosedText
	}
	return
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeobstructionidsurfacezonetype.Name
	case "NilReason":
		res.valueString = codeobstructionidsurfacezonetype.NilReason
	case "EnclosedText":
		res.valueString = codeobstructionidsurfacezonetype.EnclosedText
	}
	return
}

func (codernptype *CodeRNPType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codernptype.Name
	case "NilReason":
		res.valueString = codernptype.NilReason
	case "EnclosedText":
		res.valueString = codernptype.EnclosedText
	}
	return
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codervsmpointroletype.Name
	case "NilReason":
		res.valueString = codervsmpointroletype.NilReason
	case "EnclosedText":
		res.valueString = codervsmpointroletype.EnclosedText
	}
	return
}

func (codervsmtype *CodeRVSMType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codervsmtype.Name
	case "NilReason":
		res.valueString = codervsmtype.NilReason
	case "EnclosedText":
		res.valueString = codervsmtype.EnclosedText
	}
	return
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codereferenceroletype.Name
	case "NilReason":
		res.valueString = codereferenceroletype.NilReason
	case "EnclosedText":
		res.valueString = codereferenceroletype.EnclosedText
	}
	return
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderouteavailabilitytype.Name
	case "NilReason":
		res.valueString = coderouteavailabilitytype.NilReason
	case "EnclosedText":
		res.valueString = coderouteavailabilitytype.EnclosedText
	}
	return
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutedesignatorlettertype.Name
	case "NilReason":
		res.valueString = coderoutedesignatorlettertype.NilReason
	case "EnclosedText":
		res.valueString = coderoutedesignatorlettertype.EnclosedText
	}
	return
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutedesignatorprefixtype.Name
	case "NilReason":
		res.valueString = coderoutedesignatorprefixtype.NilReason
	case "EnclosedText":
		res.valueString = coderoutedesignatorprefixtype.EnclosedText
	}
	return
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutedesignatorsuffixtype.Name
	case "NilReason":
		res.valueString = coderoutedesignatorsuffixtype.NilReason
	case "EnclosedText":
		res.valueString = coderoutedesignatorsuffixtype.EnclosedText
	}
	return
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutenavigationtype.Name
	case "NilReason":
		res.valueString = coderoutenavigationtype.NilReason
	case "EnclosedText":
		res.valueString = coderoutenavigationtype.EnclosedText
	}
	return
}

func (coderouteorigintype *CodeRouteOriginType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderouteorigintype.Name
	case "NilReason":
		res.valueString = coderouteorigintype.NilReason
	case "EnclosedText":
		res.valueString = coderouteorigintype.EnclosedText
	}
	return
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutesegmentpathtype.Name
	case "NilReason":
		res.valueString = coderoutesegmentpathtype.NilReason
	case "EnclosedText":
		res.valueString = coderoutesegmentpathtype.EnclosedText
	}
	return
}

func (coderoutetype *CodeRouteType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoutetype.Name
	case "NilReason":
		res.valueString = coderoutetype.NilReason
	case "EnclosedText":
		res.valueString = coderoutetype.EnclosedText
	}
	return
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderuleproceduretitletype.Name
	case "NilReason":
		res.valueString = coderuleproceduretitletype.NilReason
	case "EnclosedText":
		res.valueString = coderuleproceduretitletype.EnclosedText
	}
	return
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderuleproceduretype.Name
	case "NilReason":
		res.valueString = coderuleproceduretype.NilReason
	case "EnclosedText":
		res.valueString = coderuleproceduretype.EnclosedText
	}
	return
}

func (codesidetype *CodeSideType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesidetype.Name
	case "NilReason":
		res.valueString = codesidetype.NilReason
	case "EnclosedText":
		res.valueString = codesidetype.EnclosedText
	}
	return
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codespeedreferencetype.Name
	case "NilReason":
		res.valueString = codespeedreferencetype.NilReason
	case "EnclosedText":
		res.valueString = codespeedreferencetype.EnclosedText
	}
	return
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetacanchanneltype.Name
	case "NilReason":
		res.valueString = codetacanchanneltype.NilReason
	case "EnclosedText":
		res.valueString = codetacanchanneltype.EnclosedText
	}
	return
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetelecomnetworktype.Name
	case "NilReason":
		res.valueString = codetelecomnetworktype.NilReason
	case "EnclosedText":
		res.valueString = codetelecomnetworktype.EnclosedText
	}
	return
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetimeeventcombinationtype.Name
	case "NilReason":
		res.valueString = codetimeeventcombinationtype.NilReason
	case "EnclosedText":
		res.valueString = codetimeeventcombinationtype.EnclosedText
	}
	return
}

func (codetimeeventtype *CodeTimeEventType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetimeeventtype.Name
	case "NilReason":
		res.valueString = codetimeeventtype.NilReason
	case "EnclosedText":
		res.valueString = codetimeeventtype.EnclosedText
	}
	return
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetimereferencetype.Name
	case "NilReason":
		res.valueString = codetimereferencetype.NilReason
	case "EnclosedText":
		res.valueString = codetimereferencetype.EnclosedText
	}
	return
}

func (codetranspondertype *CodeTransponderType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetranspondertype.Name
	case "NilReason":
		res.valueString = codetranspondertype.NilReason
	case "EnclosedText":
		res.valueString = codetranspondertype.EnclosedText
	}
	return
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeupperalphatype.Name
	case "NilReason":
		res.valueString = codeupperalphatype.NilReason
	case "EnclosedText":
		res.valueString = codeupperalphatype.EnclosedText
	}
	return
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codevalueinterpretationtype.Name
	case "NilReason":
		res.valueString = codevalueinterpretationtype.NilReason
	case "EnclosedText":
		res.valueString = codevalueinterpretationtype.EnclosedText
	}
	return
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeverticaldatumtype.Name
	case "NilReason":
		res.valueString = codeverticaldatumtype.NilReason
	case "EnclosedText":
		res.valueString = codeverticaldatumtype.EnclosedText
	}
	return
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeverticalreferencetype.Name
	case "NilReason":
		res.valueString = codeverticalreferencetype.NilReason
	case "EnclosedText":
		res.valueString = codeverticalreferencetype.EnclosedText
	}
	return
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codewaketurbulencetype.Name
	case "NilReason":
		res.valueString = codewaketurbulencetype.NilReason
	case "EnclosedText":
		res.valueString = codewaketurbulencetype.EnclosedText
	}
	return
}

func (codeyesnotype *CodeYesNoType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeyesnotype.Name
	case "NilReason":
		res.valueString = codeyesnotype.NilReason
	case "EnclosedText":
		res.valueString = codeyesnotype.EnclosedText
	}
	return
}

func (datemonthdaytype *DateMonthDayType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = datemonthdaytype.Name
	case "NilReason":
		res.valueString = datemonthdaytype.NilReason
	case "EnclosedText":
		res.valueString = datemonthdaytype.EnclosedText
	}
	return
}

func (nonumbertype *NoNumberType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = nonumbertype.Name
	case "NilReason":
		res.valueString = nonumbertype.NilReason
	case "EnclosedText":
		res.valueString = fmt.Sprintf("%d", nonumbertype.EnclosedText)
		res.valueInt = nonumbertype.EnclosedText
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (nosequencetype *NoSequenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = nosequencetype.Name
	case "NilReason":
		res.valueString = nosequencetype.NilReason
	case "EnclosedText":
		res.valueString = fmt.Sprintf("%d", nosequencetype.EnclosedText)
		res.valueInt = nosequencetype.EnclosedText
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (textaddresstype *TextAddressType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textaddresstype.Name
	case "NilReason":
		res.valueString = textaddresstype.NilReason
	case "EnclosedText":
		res.valueString = textaddresstype.EnclosedText
	}
	return
}

func (textdesignatortype *TextDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textdesignatortype.Name
	case "NilReason":
		res.valueString = textdesignatortype.NilReason
	case "EnclosedText":
		res.valueString = textdesignatortype.EnclosedText
	}
	return
}

func (textinstructiontype *TextInstructionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textinstructiontype.Name
	case "NilReason":
		res.valueString = textinstructiontype.NilReason
	case "EnclosedText":
		res.valueString = textinstructiontype.EnclosedText
	}
	return
}

func (textnametype *TextNameType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textnametype.Name
	case "NilReason":
		res.valueString = textnametype.NilReason
	case "EnclosedText":
		res.valueString = textnametype.EnclosedText
	}
	return
}

func (textnotetype *TextNoteType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textnotetype.Name
	case "NilReason":
		res.valueString = textnotetype.NilReason
	case "Lang":
		res.valueString = textnotetype.Lang
	case "EnclosedText":
		res.valueString = textnotetype.EnclosedText
	}
	return
}

func (textphonetype *TextPhoneType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textphonetype.Name
	case "NilReason":
		res.valueString = textphonetype.NilReason
	case "EnclosedText":
		res.valueString = textphonetype.EnclosedText
	}
	return
}

func (textpropertynametype *TextPropertyNameType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textpropertynametype.Name
	case "NilReason":
		res.valueString = textpropertynametype.NilReason
	case "EnclosedText":
		res.valueString = textpropertynametype.EnclosedText
	}
	return
}

func (timetype *TimeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timetype.Name
	case "NilReason":
		res.valueString = timetype.NilReason
	case "EnclosedText":
		res.valueString = timetype.EnclosedText
	}
	return
}

func (valangletype *ValAngleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valangletype.Name
	case "NilReason":
		res.valueString = valangletype.NilReason
	case "EnclosedText":
		res.valueString = valangletype.EnclosedText
	}
	return
}

func (valbearingtype *ValBearingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valbearingtype.Name
	case "NilReason":
		res.valueString = valbearingtype.NilReason
	case "EnclosedText":
		res.valueString = valbearingtype.EnclosedText
	}
	return
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valdistancesignedtype.Name
	case "Uom":
		res.valueString = valdistancesignedtype.Uom
	case "NilReason":
		res.valueString = valdistancesignedtype.NilReason
	case "EnclosedText":
		res.valueString = valdistancesignedtype.EnclosedText
	}
	return
}

func (valdistancetype *ValDistanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valdistancetype.Name
	case "Uom":
		res.valueString = valdistancetype.Uom
	case "NilReason":
		res.valueString = valdistancetype.NilReason
	case "EnclosedText":
		res.valueString = valdistancetype.EnclosedText
	}
	return
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valdistanceverticaltype.Name
	case "Uom":
		res.valueString = valdistanceverticaltype.Uom
	case "NilReason":
		res.valueString = valdistanceverticaltype.NilReason
	case "EnclosedText":
		res.valueString = valdistanceverticaltype.EnclosedText
	}
	return
}

func (valdurationtype *ValDurationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valdurationtype.Name
	case "Uom":
		res.valueString = valdurationtype.Uom
	case "NilReason":
		res.valueString = valdurationtype.NilReason
	case "EnclosedText":
		res.valueString = valdurationtype.EnclosedText
	}
	return
}

func (valslopetype *ValSlopeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valslopetype.Name
	case "NilReason":
		res.valueString = valslopetype.NilReason
	case "EnclosedText":
		res.valueString = valslopetype.EnclosedText
	}
	return
}

func (valspeedtype *ValSpeedType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valspeedtype.Name
	case "Uom":
		res.valueString = valspeedtype.Uom
	case "NilReason":
		res.valueString = valspeedtype.NilReason
	case "EnclosedText":
		res.valueString = valspeedtype.EnclosedText
	}
	return
}

func (valweighttype *ValWeightType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valweighttype.Name
	case "Uom":
		res.valueString = valweighttype.Uom
	case "NilReason":
		res.valueString = valweighttype.NilReason
	case "EnclosedText":
		res.valueString = valweighttype.EnclosedText
	}
	return
}

func (xhtmltype *XHTMLType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xhtmltype.Name
	case "NilReason":
		res.valueString = xhtmltype.NilReason
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
func (codeatcreportingtype *CodeATCReportingType) GongGetGongstructName() string {
	return "CodeATCReportingType"
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetGongstructName() string {
	return "CodeAerialRefuellingPointType"
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetGongstructName() string {
	return "CodeAerialRefuellingPrefixType"
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetGongstructName() string {
	return "CodeAerialRefuellingType"
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetGongstructName() string {
	return "CodeAircraftCategoryType"
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetGongstructName() string {
	return "CodeAircraftEngineNumberType"
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetGongstructName() string {
	return "CodeAircraftEngineType"
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetGongstructName() string {
	return "CodeAircraftICAOType"
}

func (codeaircrafttype *CodeAircraftType) GongGetGongstructName() string {
	return "CodeAircraftType"
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetGongstructName() string {
	return "CodeAircraftWingspanClassType"
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetGongstructName() string {
	return "CodeAltitudeAdjustmentType"
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetGongstructName() string {
	return "CodeAltitudeUseType"
}

func (codeauthoritytype *CodeAuthorityType) GongGetGongstructName() string {
	return "CodeAuthorityType"
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetGongstructName() string {
	return "CodeCardinalDirectionType"
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetGongstructName() string {
	return "CodeCommunicationModeType"
}

func (codecomparisontype *CodeComparisonType) GongGetGongstructName() string {
	return "CodeComparisonType"
}

func (codecoursetype *CodeCourseType) GongGetGongstructName() string {
	return "CodeCourseType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetGongstructName() string {
	return "CodeDirectionTurnType"
}

func (codedirectiontype *CodeDirectionType) GongGetGongstructName() string {
	return "CodeDirectionType"
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetGongstructName() string {
	return "CodeEquipmentAntiCollisionType"
}

func (codeflightorigintype *CodeFlightOriginType) GongGetGongstructName() string {
	return "CodeFlightOriginType"
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetGongstructName() string {
	return "CodeFlightPurposeType"
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetGongstructName() string {
	return "CodeFlightRestrictionDesignatorType"
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetGongstructName() string {
	return "CodeFlightRestrictionType"
}

func (codeflightruletype *CodeFlightRuleType) GongGetGongstructName() string {
	return "CodeFlightRuleType"
}

func (codeflightstatustype *CodeFlightStatusType) GongGetGongstructName() string {
	return "CodeFlightStatusType"
}

func (codeflighttype *CodeFlightType) GongGetGongstructName() string {
	return "CodeFlightType"
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetGongstructName() string {
	return "CodeFlowConditionOperationType"
}

func (codefreeflighttype *CodeFreeFlightType) GongGetGongstructName() string {
	return "CodeFreeFlightType"
}

func (codeleveltype *CodeLevelType) GongGetGongstructName() string {
	return "CodeLevelType"
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetGongstructName() string {
	return "CodeLocationQualifierType"
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetGongstructName() string {
	return "CodeMeteoConditionsType"
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetGongstructName() string {
	return "CodeMilitaryRoutePointType"
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetGongstructName() string {
	return "CodeMilitaryStatusType"
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetGongstructName() string {
	return "CodeMilitaryTrainingType"
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetGongstructName() string {
	return "CodeNavigationEquipmentType"
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetGongstructName() string {
	return "CodeNavigationSpecificationType"
}

func (codenotepurposetype *CodeNotePurposeType) GongGetGongstructName() string {
	return "CodeNotePurposeType"
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetGongstructName() string {
	return "CodeObstacleAssessmentSurfaceType"
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetGongstructName() string {
	return "CodeObstructionIdSurfaceZoneType"
}

func (codernptype *CodeRNPType) GongGetGongstructName() string {
	return "CodeRNPType"
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetGongstructName() string {
	return "CodeRVSMPointRoleType"
}

func (codervsmtype *CodeRVSMType) GongGetGongstructName() string {
	return "CodeRVSMType"
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetGongstructName() string {
	return "CodeReferenceRoleType"
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetGongstructName() string {
	return "CodeRouteAvailabilityType"
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetGongstructName() string {
	return "CodeRouteDesignatorLetterType"
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetGongstructName() string {
	return "CodeRouteDesignatorPrefixType"
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetGongstructName() string {
	return "CodeRouteDesignatorSuffixType"
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetGongstructName() string {
	return "CodeRouteNavigationType"
}

func (coderouteorigintype *CodeRouteOriginType) GongGetGongstructName() string {
	return "CodeRouteOriginType"
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetGongstructName() string {
	return "CodeRouteSegmentPathType"
}

func (coderoutetype *CodeRouteType) GongGetGongstructName() string {
	return "CodeRouteType"
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetGongstructName() string {
	return "CodeRuleProcedureTitleType"
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetGongstructName() string {
	return "CodeRuleProcedureType"
}

func (codesidetype *CodeSideType) GongGetGongstructName() string {
	return "CodeSideType"
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetGongstructName() string {
	return "CodeSpeedReferenceType"
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetGongstructName() string {
	return "CodeTACANChannelType"
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetGongstructName() string {
	return "CodeTelecomNetworkType"
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetGongstructName() string {
	return "CodeTimeEventCombinationType"
}

func (codetimeeventtype *CodeTimeEventType) GongGetGongstructName() string {
	return "CodeTimeEventType"
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetGongstructName() string {
	return "CodeTimeReferenceType"
}

func (codetranspondertype *CodeTransponderType) GongGetGongstructName() string {
	return "CodeTransponderType"
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetGongstructName() string {
	return "CodeUpperAlphaType"
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetGongstructName() string {
	return "CodeValueInterpretationType"
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetGongstructName() string {
	return "CodeVerticalDatumType"
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetGongstructName() string {
	return "CodeVerticalReferenceType"
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetGongstructName() string {
	return "CodeWakeTurbulenceType"
}

func (codeyesnotype *CodeYesNoType) GongGetGongstructName() string {
	return "CodeYesNoType"
}

func (datemonthdaytype *DateMonthDayType) GongGetGongstructName() string {
	return "DateMonthDayType"
}

func (nonumbertype *NoNumberType) GongGetGongstructName() string {
	return "NoNumberType"
}

func (nosequencetype *NoSequenceType) GongGetGongstructName() string {
	return "NoSequenceType"
}

func (textaddresstype *TextAddressType) GongGetGongstructName() string {
	return "TextAddressType"
}

func (textdesignatortype *TextDesignatorType) GongGetGongstructName() string {
	return "TextDesignatorType"
}

func (textinstructiontype *TextInstructionType) GongGetGongstructName() string {
	return "TextInstructionType"
}

func (textnametype *TextNameType) GongGetGongstructName() string {
	return "TextNameType"
}

func (textnotetype *TextNoteType) GongGetGongstructName() string {
	return "TextNoteType"
}

func (textphonetype *TextPhoneType) GongGetGongstructName() string {
	return "TextPhoneType"
}

func (textpropertynametype *TextPropertyNameType) GongGetGongstructName() string {
	return "TextPropertyNameType"
}

func (timetype *TimeType) GongGetGongstructName() string {
	return "TimeType"
}

func (valangletype *ValAngleType) GongGetGongstructName() string {
	return "ValAngleType"
}

func (valbearingtype *ValBearingType) GongGetGongstructName() string {
	return "ValBearingType"
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetGongstructName() string {
	return "ValDistanceSignedType"
}

func (valdistancetype *ValDistanceType) GongGetGongstructName() string {
	return "ValDistanceType"
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetGongstructName() string {
	return "ValDistanceVerticalType"
}

func (valdurationtype *ValDurationType) GongGetGongstructName() string {
	return "ValDurationType"
}

func (valslopetype *ValSlopeType) GongGetGongstructName() string {
	return "ValSlopeType"
}

func (valspeedtype *ValSpeedType) GongGetGongstructName() string {
	return "ValSpeedType"
}

func (valweighttype *ValWeightType) GongGetGongstructName() string {
	return "ValWeightType"
}

func (xhtmltype *XHTMLType) GongGetGongstructName() string {
	return "XHTMLType"
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
	__gong__rebuildMapString(stage.CodeATCReportingTypes, &stage.CodeATCReportingTypes_mapString)

	__gong__rebuildMapString(stage.CodeAerialRefuellingPointTypes, &stage.CodeAerialRefuellingPointTypes_mapString)

	__gong__rebuildMapString(stage.CodeAerialRefuellingPrefixTypes, &stage.CodeAerialRefuellingPrefixTypes_mapString)

	__gong__rebuildMapString(stage.CodeAerialRefuellingTypes, &stage.CodeAerialRefuellingTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString)

	__gong__rebuildMapString(stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_mapString)

	__gong__rebuildMapString(stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString)

	__gong__rebuildMapString(stage.CodeComparisonTypes, &stage.CodeComparisonTypes_mapString)

	__gong__rebuildMapString(stage.CodeCourseTypes, &stage.CodeCourseTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionTypes, &stage.CodeDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightRestrictionDesignatorTypes, &stage.CodeFlightRestrictionDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightRestrictionTypes, &stage.CodeFlightRestrictionTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightTypes, &stage.CodeFlightTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlowConditionOperationTypes, &stage.CodeFlowConditionOperationTypes_mapString)

	__gong__rebuildMapString(stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_mapString)

	__gong__rebuildMapString(stage.CodeLevelTypes, &stage.CodeLevelTypes_mapString)

	__gong__rebuildMapString(stage.CodeLocationQualifierTypes, &stage.CodeLocationQualifierTypes_mapString)

	__gong__rebuildMapString(stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryTrainingTypes, &stage.CodeMilitaryTrainingTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_mapString)

	__gong__rebuildMapString(stage.CodeRNPTypes, &stage.CodeRNPTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString)

	__gong__rebuildMapString(stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteAvailabilityTypes, &stage.CodeRouteAvailabilityTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteDesignatorLetterTypes, &stage.CodeRouteDesignatorLetterTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteDesignatorPrefixTypes, &stage.CodeRouteDesignatorPrefixTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteDesignatorSuffixTypes, &stage.CodeRouteDesignatorSuffixTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteNavigationTypes, &stage.CodeRouteNavigationTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteOriginTypes, &stage.CodeRouteOriginTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteSegmentPathTypes, &stage.CodeRouteSegmentPathTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteTypes, &stage.CodeRouteTypes_mapString)

	__gong__rebuildMapString(stage.CodeRuleProcedureTitleTypes, &stage.CodeRuleProcedureTitleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRuleProcedureTypes, &stage.CodeRuleProcedureTypes_mapString)

	__gong__rebuildMapString(stage.CodeSideTypes, &stage.CodeSideTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_mapString)

	__gong__rebuildMapString(stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString)

	__gong__rebuildMapString(stage.CodeUpperAlphaTypes, &stage.CodeUpperAlphaTypes_mapString)

	__gong__rebuildMapString(stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.NoNumberTypes, &stage.NoNumberTypes_mapString)

	__gong__rebuildMapString(stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString)

	__gong__rebuildMapString(stage.TextAddressTypes, &stage.TextAddressTypes_mapString)

	__gong__rebuildMapString(stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString)

	__gong__rebuildMapString(stage.TextNameTypes, &stage.TextNameTypes_mapString)

	__gong__rebuildMapString(stage.TextNoteTypes, &stage.TextNoteTypes_mapString)

	__gong__rebuildMapString(stage.TextPhoneTypes, &stage.TextPhoneTypes_mapString)

	__gong__rebuildMapString(stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.ValAngleTypes, &stage.ValAngleTypes_mapString)

	__gong__rebuildMapString(stage.ValBearingTypes, &stage.ValBearingTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValSlopeTypes, &stage.ValSlopeTypes_mapString)

	__gong__rebuildMapString(stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString)

	__gong__rebuildMapString(stage.ValWeightTypes, &stage.ValWeightTypes_mapString)

	__gong__rebuildMapString(stage.XHTMLTypes, &stage.XHTMLTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
