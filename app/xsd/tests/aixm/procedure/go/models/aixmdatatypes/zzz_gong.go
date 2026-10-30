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

	CodeApproachEquipmentAdditionalTypes                map[*CodeApproachEquipmentAdditionalType]struct{}
	CodeApproachEquipmentAdditionalTypes_instance       map[*CodeApproachEquipmentAdditionalType]*CodeApproachEquipmentAdditionalType
	CodeApproachEquipmentAdditionalTypes_mapString      map[string]*CodeApproachEquipmentAdditionalType
	CodeApproachEquipmentAdditionalTypeOrder            uint
	CodeApproachEquipmentAdditionalType_stagedOrder     map[*CodeApproachEquipmentAdditionalType]uint
	CodeApproachEquipmentAdditionalType_orderStaged     map[uint]*CodeApproachEquipmentAdditionalType
	CodeApproachEquipmentAdditionalTypes_reference      map[*CodeApproachEquipmentAdditionalType]*CodeApproachEquipmentAdditionalType
	CodeApproachEquipmentAdditionalTypes_referenceOrder map[*CodeApproachEquipmentAdditionalType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachEquipmentAdditionalTypeCreateCallback GongOnAfterCreateInterface[CodeApproachEquipmentAdditionalType]
	OnAfterCodeApproachEquipmentAdditionalTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachEquipmentAdditionalType]
	OnAfterCodeApproachEquipmentAdditionalTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachEquipmentAdditionalType]

	CodeApproachGuidanceTypes                map[*CodeApproachGuidanceType]struct{}
	CodeApproachGuidanceTypes_instance       map[*CodeApproachGuidanceType]*CodeApproachGuidanceType
	CodeApproachGuidanceTypes_mapString      map[string]*CodeApproachGuidanceType
	CodeApproachGuidanceTypeOrder            uint
	CodeApproachGuidanceType_stagedOrder     map[*CodeApproachGuidanceType]uint
	CodeApproachGuidanceType_orderStaged     map[uint]*CodeApproachGuidanceType
	CodeApproachGuidanceTypes_reference      map[*CodeApproachGuidanceType]*CodeApproachGuidanceType
	CodeApproachGuidanceTypes_referenceOrder map[*CodeApproachGuidanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachGuidanceTypeCreateCallback GongOnAfterCreateInterface[CodeApproachGuidanceType]
	OnAfterCodeApproachGuidanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachGuidanceType]
	OnAfterCodeApproachGuidanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachGuidanceType]

	CodeApproachPrefixTypes                map[*CodeApproachPrefixType]struct{}
	CodeApproachPrefixTypes_instance       map[*CodeApproachPrefixType]*CodeApproachPrefixType
	CodeApproachPrefixTypes_mapString      map[string]*CodeApproachPrefixType
	CodeApproachPrefixTypeOrder            uint
	CodeApproachPrefixType_stagedOrder     map[*CodeApproachPrefixType]uint
	CodeApproachPrefixType_orderStaged     map[uint]*CodeApproachPrefixType
	CodeApproachPrefixTypes_reference      map[*CodeApproachPrefixType]*CodeApproachPrefixType
	CodeApproachPrefixTypes_referenceOrder map[*CodeApproachPrefixType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachPrefixTypeCreateCallback GongOnAfterCreateInterface[CodeApproachPrefixType]
	OnAfterCodeApproachPrefixTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachPrefixType]
	OnAfterCodeApproachPrefixTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachPrefixType]

	CodeApproachTypes                map[*CodeApproachType]struct{}
	CodeApproachTypes_instance       map[*CodeApproachType]*CodeApproachType
	CodeApproachTypes_mapString      map[string]*CodeApproachType
	CodeApproachTypeOrder            uint
	CodeApproachType_stagedOrder     map[*CodeApproachType]uint
	CodeApproachType_orderStaged     map[uint]*CodeApproachType
	CodeApproachTypes_reference      map[*CodeApproachType]*CodeApproachType
	CodeApproachTypes_referenceOrder map[*CodeApproachType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachTypeCreateCallback GongOnAfterCreateInterface[CodeApproachType]
	OnAfterCodeApproachTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachType]
	OnAfterCodeApproachTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachType]

	CodeApprovalTypes                map[*CodeApprovalType]struct{}
	CodeApprovalTypes_instance       map[*CodeApprovalType]*CodeApprovalType
	CodeApprovalTypes_mapString      map[string]*CodeApprovalType
	CodeApprovalTypeOrder            uint
	CodeApprovalType_stagedOrder     map[*CodeApprovalType]uint
	CodeApprovalType_orderStaged     map[uint]*CodeApprovalType
	CodeApprovalTypes_reference      map[*CodeApprovalType]*CodeApprovalType
	CodeApprovalTypes_referenceOrder map[*CodeApprovalType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApprovalTypeCreateCallback GongOnAfterCreateInterface[CodeApprovalType]
	OnAfterCodeApprovalTypeUpdateCallback GongOnAfterUpdateInterface[CodeApprovalType]
	OnAfterCodeApprovalTypeDeleteCallback GongOnAfterDeleteInterface[CodeApprovalType]

	CodeArcDirectionTypes                map[*CodeArcDirectionType]struct{}
	CodeArcDirectionTypes_instance       map[*CodeArcDirectionType]*CodeArcDirectionType
	CodeArcDirectionTypes_mapString      map[string]*CodeArcDirectionType
	CodeArcDirectionTypeOrder            uint
	CodeArcDirectionType_stagedOrder     map[*CodeArcDirectionType]uint
	CodeArcDirectionType_orderStaged     map[uint]*CodeArcDirectionType
	CodeArcDirectionTypes_reference      map[*CodeArcDirectionType]*CodeArcDirectionType
	CodeArcDirectionTypes_referenceOrder map[*CodeArcDirectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeArcDirectionTypeCreateCallback GongOnAfterCreateInterface[CodeArcDirectionType]
	OnAfterCodeArcDirectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeArcDirectionType]
	OnAfterCodeArcDirectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeArcDirectionType]

	CodeBearingTypes                map[*CodeBearingType]struct{}
	CodeBearingTypes_instance       map[*CodeBearingType]*CodeBearingType
	CodeBearingTypes_mapString      map[string]*CodeBearingType
	CodeBearingTypeOrder            uint
	CodeBearingType_stagedOrder     map[*CodeBearingType]uint
	CodeBearingType_orderStaged     map[uint]*CodeBearingType
	CodeBearingTypes_reference      map[*CodeBearingType]*CodeBearingType
	CodeBearingTypes_referenceOrder map[*CodeBearingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeBearingTypeCreateCallback GongOnAfterCreateInterface[CodeBearingType]
	OnAfterCodeBearingTypeUpdateCallback GongOnAfterUpdateInterface[CodeBearingType]
	OnAfterCodeBearingTypeDeleteCallback GongOnAfterDeleteInterface[CodeBearingType]

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

	CodeDesignStandardTypes                map[*CodeDesignStandardType]struct{}
	CodeDesignStandardTypes_instance       map[*CodeDesignStandardType]*CodeDesignStandardType
	CodeDesignStandardTypes_mapString      map[string]*CodeDesignStandardType
	CodeDesignStandardTypeOrder            uint
	CodeDesignStandardType_stagedOrder     map[*CodeDesignStandardType]uint
	CodeDesignStandardType_orderStaged     map[uint]*CodeDesignStandardType
	CodeDesignStandardTypes_reference      map[*CodeDesignStandardType]*CodeDesignStandardType
	CodeDesignStandardTypes_referenceOrder map[*CodeDesignStandardType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDesignStandardTypeCreateCallback GongOnAfterCreateInterface[CodeDesignStandardType]
	OnAfterCodeDesignStandardTypeUpdateCallback GongOnAfterUpdateInterface[CodeDesignStandardType]
	OnAfterCodeDesignStandardTypeDeleteCallback GongOnAfterDeleteInterface[CodeDesignStandardType]

	CodeDesignatedPointDesignatorTypes                map[*CodeDesignatedPointDesignatorType]struct{}
	CodeDesignatedPointDesignatorTypes_instance       map[*CodeDesignatedPointDesignatorType]*CodeDesignatedPointDesignatorType
	CodeDesignatedPointDesignatorTypes_mapString      map[string]*CodeDesignatedPointDesignatorType
	CodeDesignatedPointDesignatorTypeOrder            uint
	CodeDesignatedPointDesignatorType_stagedOrder     map[*CodeDesignatedPointDesignatorType]uint
	CodeDesignatedPointDesignatorType_orderStaged     map[uint]*CodeDesignatedPointDesignatorType
	CodeDesignatedPointDesignatorTypes_reference      map[*CodeDesignatedPointDesignatorType]*CodeDesignatedPointDesignatorType
	CodeDesignatedPointDesignatorTypes_referenceOrder map[*CodeDesignatedPointDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDesignatedPointDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeDesignatedPointDesignatorType]
	OnAfterCodeDesignatedPointDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeDesignatedPointDesignatorType]
	OnAfterCodeDesignatedPointDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeDesignatedPointDesignatorType]

	CodeDirectionReferenceTypes                map[*CodeDirectionReferenceType]struct{}
	CodeDirectionReferenceTypes_instance       map[*CodeDirectionReferenceType]*CodeDirectionReferenceType
	CodeDirectionReferenceTypes_mapString      map[string]*CodeDirectionReferenceType
	CodeDirectionReferenceTypeOrder            uint
	CodeDirectionReferenceType_stagedOrder     map[*CodeDirectionReferenceType]uint
	CodeDirectionReferenceType_orderStaged     map[uint]*CodeDirectionReferenceType
	CodeDirectionReferenceTypes_reference      map[*CodeDirectionReferenceType]*CodeDirectionReferenceType
	CodeDirectionReferenceTypes_referenceOrder map[*CodeDirectionReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDirectionReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeDirectionReferenceType]
	OnAfterCodeDirectionReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeDirectionReferenceType]
	OnAfterCodeDirectionReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeDirectionReferenceType]

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

	CodeEquipmentUnavailableTypes                map[*CodeEquipmentUnavailableType]struct{}
	CodeEquipmentUnavailableTypes_instance       map[*CodeEquipmentUnavailableType]*CodeEquipmentUnavailableType
	CodeEquipmentUnavailableTypes_mapString      map[string]*CodeEquipmentUnavailableType
	CodeEquipmentUnavailableTypeOrder            uint
	CodeEquipmentUnavailableType_stagedOrder     map[*CodeEquipmentUnavailableType]uint
	CodeEquipmentUnavailableType_orderStaged     map[uint]*CodeEquipmentUnavailableType
	CodeEquipmentUnavailableTypes_reference      map[*CodeEquipmentUnavailableType]*CodeEquipmentUnavailableType
	CodeEquipmentUnavailableTypes_referenceOrder map[*CodeEquipmentUnavailableType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeEquipmentUnavailableTypeCreateCallback GongOnAfterCreateInterface[CodeEquipmentUnavailableType]
	OnAfterCodeEquipmentUnavailableTypeUpdateCallback GongOnAfterUpdateInterface[CodeEquipmentUnavailableType]
	OnAfterCodeEquipmentUnavailableTypeDeleteCallback GongOnAfterDeleteInterface[CodeEquipmentUnavailableType]

	CodeFinalGuidanceTypes                map[*CodeFinalGuidanceType]struct{}
	CodeFinalGuidanceTypes_instance       map[*CodeFinalGuidanceType]*CodeFinalGuidanceType
	CodeFinalGuidanceTypes_mapString      map[string]*CodeFinalGuidanceType
	CodeFinalGuidanceTypeOrder            uint
	CodeFinalGuidanceType_stagedOrder     map[*CodeFinalGuidanceType]uint
	CodeFinalGuidanceType_orderStaged     map[uint]*CodeFinalGuidanceType
	CodeFinalGuidanceTypes_reference      map[*CodeFinalGuidanceType]*CodeFinalGuidanceType
	CodeFinalGuidanceTypes_referenceOrder map[*CodeFinalGuidanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFinalGuidanceTypeCreateCallback GongOnAfterCreateInterface[CodeFinalGuidanceType]
	OnAfterCodeFinalGuidanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeFinalGuidanceType]
	OnAfterCodeFinalGuidanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeFinalGuidanceType]

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

	CodeHeightReferenceTypes                map[*CodeHeightReferenceType]struct{}
	CodeHeightReferenceTypes_instance       map[*CodeHeightReferenceType]*CodeHeightReferenceType
	CodeHeightReferenceTypes_mapString      map[string]*CodeHeightReferenceType
	CodeHeightReferenceTypeOrder            uint
	CodeHeightReferenceType_stagedOrder     map[*CodeHeightReferenceType]uint
	CodeHeightReferenceType_orderStaged     map[uint]*CodeHeightReferenceType
	CodeHeightReferenceTypes_reference      map[*CodeHeightReferenceType]*CodeHeightReferenceType
	CodeHeightReferenceTypes_referenceOrder map[*CodeHeightReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeHeightReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeHeightReferenceType]
	OnAfterCodeHeightReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeHeightReferenceType]
	OnAfterCodeHeightReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeHeightReferenceType]

	CodeHoldingUsageTypes                map[*CodeHoldingUsageType]struct{}
	CodeHoldingUsageTypes_instance       map[*CodeHoldingUsageType]*CodeHoldingUsageType
	CodeHoldingUsageTypes_mapString      map[string]*CodeHoldingUsageType
	CodeHoldingUsageTypeOrder            uint
	CodeHoldingUsageType_stagedOrder     map[*CodeHoldingUsageType]uint
	CodeHoldingUsageType_orderStaged     map[uint]*CodeHoldingUsageType
	CodeHoldingUsageTypes_reference      map[*CodeHoldingUsageType]*CodeHoldingUsageType
	CodeHoldingUsageTypes_referenceOrder map[*CodeHoldingUsageType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeHoldingUsageTypeCreateCallback GongOnAfterCreateInterface[CodeHoldingUsageType]
	OnAfterCodeHoldingUsageTypeUpdateCallback GongOnAfterUpdateInterface[CodeHoldingUsageType]
	OnAfterCodeHoldingUsageTypeDeleteCallback GongOnAfterDeleteInterface[CodeHoldingUsageType]

	CodeHoldingUseTypes                map[*CodeHoldingUseType]struct{}
	CodeHoldingUseTypes_instance       map[*CodeHoldingUseType]*CodeHoldingUseType
	CodeHoldingUseTypes_mapString      map[string]*CodeHoldingUseType
	CodeHoldingUseTypeOrder            uint
	CodeHoldingUseType_stagedOrder     map[*CodeHoldingUseType]uint
	CodeHoldingUseType_orderStaged     map[uint]*CodeHoldingUseType
	CodeHoldingUseTypes_reference      map[*CodeHoldingUseType]*CodeHoldingUseType
	CodeHoldingUseTypes_referenceOrder map[*CodeHoldingUseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeHoldingUseTypeCreateCallback GongOnAfterCreateInterface[CodeHoldingUseType]
	OnAfterCodeHoldingUseTypeUpdateCallback GongOnAfterUpdateInterface[CodeHoldingUseType]
	OnAfterCodeHoldingUseTypeDeleteCallback GongOnAfterDeleteInterface[CodeHoldingUseType]

	CodeICAOCountryTypes                map[*CodeICAOCountryType]struct{}
	CodeICAOCountryTypes_instance       map[*CodeICAOCountryType]*CodeICAOCountryType
	CodeICAOCountryTypes_mapString      map[string]*CodeICAOCountryType
	CodeICAOCountryTypeOrder            uint
	CodeICAOCountryType_stagedOrder     map[*CodeICAOCountryType]uint
	CodeICAOCountryType_orderStaged     map[uint]*CodeICAOCountryType
	CodeICAOCountryTypes_reference      map[*CodeICAOCountryType]*CodeICAOCountryType
	CodeICAOCountryTypes_referenceOrder map[*CodeICAOCountryType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeICAOCountryTypeCreateCallback GongOnAfterCreateInterface[CodeICAOCountryType]
	OnAfterCodeICAOCountryTypeUpdateCallback GongOnAfterUpdateInterface[CodeICAOCountryType]
	OnAfterCodeICAOCountryTypeDeleteCallback GongOnAfterDeleteInterface[CodeICAOCountryType]

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

	CodeMinimaFinalApproachPathTypes                map[*CodeMinimaFinalApproachPathType]struct{}
	CodeMinimaFinalApproachPathTypes_instance       map[*CodeMinimaFinalApproachPathType]*CodeMinimaFinalApproachPathType
	CodeMinimaFinalApproachPathTypes_mapString      map[string]*CodeMinimaFinalApproachPathType
	CodeMinimaFinalApproachPathTypeOrder            uint
	CodeMinimaFinalApproachPathType_stagedOrder     map[*CodeMinimaFinalApproachPathType]uint
	CodeMinimaFinalApproachPathType_orderStaged     map[uint]*CodeMinimaFinalApproachPathType
	CodeMinimaFinalApproachPathTypes_reference      map[*CodeMinimaFinalApproachPathType]*CodeMinimaFinalApproachPathType
	CodeMinimaFinalApproachPathTypes_referenceOrder map[*CodeMinimaFinalApproachPathType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMinimaFinalApproachPathTypeCreateCallback GongOnAfterCreateInterface[CodeMinimaFinalApproachPathType]
	OnAfterCodeMinimaFinalApproachPathTypeUpdateCallback GongOnAfterUpdateInterface[CodeMinimaFinalApproachPathType]
	OnAfterCodeMinimaFinalApproachPathTypeDeleteCallback GongOnAfterDeleteInterface[CodeMinimaFinalApproachPathType]

	CodeMinimumAltitudeTypes                map[*CodeMinimumAltitudeType]struct{}
	CodeMinimumAltitudeTypes_instance       map[*CodeMinimumAltitudeType]*CodeMinimumAltitudeType
	CodeMinimumAltitudeTypes_mapString      map[string]*CodeMinimumAltitudeType
	CodeMinimumAltitudeTypeOrder            uint
	CodeMinimumAltitudeType_stagedOrder     map[*CodeMinimumAltitudeType]uint
	CodeMinimumAltitudeType_orderStaged     map[uint]*CodeMinimumAltitudeType
	CodeMinimumAltitudeTypes_reference      map[*CodeMinimumAltitudeType]*CodeMinimumAltitudeType
	CodeMinimumAltitudeTypes_referenceOrder map[*CodeMinimumAltitudeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMinimumAltitudeTypeCreateCallback GongOnAfterCreateInterface[CodeMinimumAltitudeType]
	OnAfterCodeMinimumAltitudeTypeUpdateCallback GongOnAfterUpdateInterface[CodeMinimumAltitudeType]
	OnAfterCodeMinimumAltitudeTypeDeleteCallback GongOnAfterDeleteInterface[CodeMinimumAltitudeType]

	CodeMinimumHeightTypes                map[*CodeMinimumHeightType]struct{}
	CodeMinimumHeightTypes_instance       map[*CodeMinimumHeightType]*CodeMinimumHeightType
	CodeMinimumHeightTypes_mapString      map[string]*CodeMinimumHeightType
	CodeMinimumHeightTypeOrder            uint
	CodeMinimumHeightType_stagedOrder     map[*CodeMinimumHeightType]uint
	CodeMinimumHeightType_orderStaged     map[uint]*CodeMinimumHeightType
	CodeMinimumHeightTypes_reference      map[*CodeMinimumHeightType]*CodeMinimumHeightType
	CodeMinimumHeightTypes_referenceOrder map[*CodeMinimumHeightType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMinimumHeightTypeCreateCallback GongOnAfterCreateInterface[CodeMinimumHeightType]
	OnAfterCodeMinimumHeightTypeUpdateCallback GongOnAfterUpdateInterface[CodeMinimumHeightType]
	OnAfterCodeMinimumHeightTypeDeleteCallback GongOnAfterDeleteInterface[CodeMinimumHeightType]

	CodeMissedApproachTypes                map[*CodeMissedApproachType]struct{}
	CodeMissedApproachTypes_instance       map[*CodeMissedApproachType]*CodeMissedApproachType
	CodeMissedApproachTypes_mapString      map[string]*CodeMissedApproachType
	CodeMissedApproachTypeOrder            uint
	CodeMissedApproachType_stagedOrder     map[*CodeMissedApproachType]uint
	CodeMissedApproachType_orderStaged     map[uint]*CodeMissedApproachType
	CodeMissedApproachTypes_reference      map[*CodeMissedApproachType]*CodeMissedApproachType
	CodeMissedApproachTypes_referenceOrder map[*CodeMissedApproachType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMissedApproachTypeCreateCallback GongOnAfterCreateInterface[CodeMissedApproachType]
	OnAfterCodeMissedApproachTypeUpdateCallback GongOnAfterUpdateInterface[CodeMissedApproachType]
	OnAfterCodeMissedApproachTypeDeleteCallback GongOnAfterDeleteInterface[CodeMissedApproachType]

	CodeNavigationAreaRestrictionTypes                map[*CodeNavigationAreaRestrictionType]struct{}
	CodeNavigationAreaRestrictionTypes_instance       map[*CodeNavigationAreaRestrictionType]*CodeNavigationAreaRestrictionType
	CodeNavigationAreaRestrictionTypes_mapString      map[string]*CodeNavigationAreaRestrictionType
	CodeNavigationAreaRestrictionTypeOrder            uint
	CodeNavigationAreaRestrictionType_stagedOrder     map[*CodeNavigationAreaRestrictionType]uint
	CodeNavigationAreaRestrictionType_orderStaged     map[uint]*CodeNavigationAreaRestrictionType
	CodeNavigationAreaRestrictionTypes_reference      map[*CodeNavigationAreaRestrictionType]*CodeNavigationAreaRestrictionType
	CodeNavigationAreaRestrictionTypes_referenceOrder map[*CodeNavigationAreaRestrictionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavigationAreaRestrictionTypeCreateCallback GongOnAfterCreateInterface[CodeNavigationAreaRestrictionType]
	OnAfterCodeNavigationAreaRestrictionTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavigationAreaRestrictionType]
	OnAfterCodeNavigationAreaRestrictionTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavigationAreaRestrictionType]

	CodeNavigationAreaTypes                map[*CodeNavigationAreaType]struct{}
	CodeNavigationAreaTypes_instance       map[*CodeNavigationAreaType]*CodeNavigationAreaType
	CodeNavigationAreaTypes_mapString      map[string]*CodeNavigationAreaType
	CodeNavigationAreaTypeOrder            uint
	CodeNavigationAreaType_stagedOrder     map[*CodeNavigationAreaType]uint
	CodeNavigationAreaType_orderStaged     map[uint]*CodeNavigationAreaType
	CodeNavigationAreaTypes_reference      map[*CodeNavigationAreaType]*CodeNavigationAreaType
	CodeNavigationAreaTypes_referenceOrder map[*CodeNavigationAreaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavigationAreaTypeCreateCallback GongOnAfterCreateInterface[CodeNavigationAreaType]
	OnAfterCodeNavigationAreaTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavigationAreaType]
	OnAfterCodeNavigationAreaTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavigationAreaType]

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

	CodeProcedureAvailabilityTypes                map[*CodeProcedureAvailabilityType]struct{}
	CodeProcedureAvailabilityTypes_instance       map[*CodeProcedureAvailabilityType]*CodeProcedureAvailabilityType
	CodeProcedureAvailabilityTypes_mapString      map[string]*CodeProcedureAvailabilityType
	CodeProcedureAvailabilityTypeOrder            uint
	CodeProcedureAvailabilityType_stagedOrder     map[*CodeProcedureAvailabilityType]uint
	CodeProcedureAvailabilityType_orderStaged     map[uint]*CodeProcedureAvailabilityType
	CodeProcedureAvailabilityTypes_reference      map[*CodeProcedureAvailabilityType]*CodeProcedureAvailabilityType
	CodeProcedureAvailabilityTypes_referenceOrder map[*CodeProcedureAvailabilityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProcedureAvailabilityTypeCreateCallback GongOnAfterCreateInterface[CodeProcedureAvailabilityType]
	OnAfterCodeProcedureAvailabilityTypeUpdateCallback GongOnAfterUpdateInterface[CodeProcedureAvailabilityType]
	OnAfterCodeProcedureAvailabilityTypeDeleteCallback GongOnAfterDeleteInterface[CodeProcedureAvailabilityType]

	CodeProcedureCodingStandardTypes                map[*CodeProcedureCodingStandardType]struct{}
	CodeProcedureCodingStandardTypes_instance       map[*CodeProcedureCodingStandardType]*CodeProcedureCodingStandardType
	CodeProcedureCodingStandardTypes_mapString      map[string]*CodeProcedureCodingStandardType
	CodeProcedureCodingStandardTypeOrder            uint
	CodeProcedureCodingStandardType_stagedOrder     map[*CodeProcedureCodingStandardType]uint
	CodeProcedureCodingStandardType_orderStaged     map[uint]*CodeProcedureCodingStandardType
	CodeProcedureCodingStandardTypes_reference      map[*CodeProcedureCodingStandardType]*CodeProcedureCodingStandardType
	CodeProcedureCodingStandardTypes_referenceOrder map[*CodeProcedureCodingStandardType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProcedureCodingStandardTypeCreateCallback GongOnAfterCreateInterface[CodeProcedureCodingStandardType]
	OnAfterCodeProcedureCodingStandardTypeUpdateCallback GongOnAfterUpdateInterface[CodeProcedureCodingStandardType]
	OnAfterCodeProcedureCodingStandardTypeDeleteCallback GongOnAfterDeleteInterface[CodeProcedureCodingStandardType]

	CodeProcedureDistanceTypes                map[*CodeProcedureDistanceType]struct{}
	CodeProcedureDistanceTypes_instance       map[*CodeProcedureDistanceType]*CodeProcedureDistanceType
	CodeProcedureDistanceTypes_mapString      map[string]*CodeProcedureDistanceType
	CodeProcedureDistanceTypeOrder            uint
	CodeProcedureDistanceType_stagedOrder     map[*CodeProcedureDistanceType]uint
	CodeProcedureDistanceType_orderStaged     map[uint]*CodeProcedureDistanceType
	CodeProcedureDistanceTypes_reference      map[*CodeProcedureDistanceType]*CodeProcedureDistanceType
	CodeProcedureDistanceTypes_referenceOrder map[*CodeProcedureDistanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProcedureDistanceTypeCreateCallback GongOnAfterCreateInterface[CodeProcedureDistanceType]
	OnAfterCodeProcedureDistanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeProcedureDistanceType]
	OnAfterCodeProcedureDistanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeProcedureDistanceType]

	CodeProcedureFixRoleTypes                map[*CodeProcedureFixRoleType]struct{}
	CodeProcedureFixRoleTypes_instance       map[*CodeProcedureFixRoleType]*CodeProcedureFixRoleType
	CodeProcedureFixRoleTypes_mapString      map[string]*CodeProcedureFixRoleType
	CodeProcedureFixRoleTypeOrder            uint
	CodeProcedureFixRoleType_stagedOrder     map[*CodeProcedureFixRoleType]uint
	CodeProcedureFixRoleType_orderStaged     map[uint]*CodeProcedureFixRoleType
	CodeProcedureFixRoleTypes_reference      map[*CodeProcedureFixRoleType]*CodeProcedureFixRoleType
	CodeProcedureFixRoleTypes_referenceOrder map[*CodeProcedureFixRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProcedureFixRoleTypeCreateCallback GongOnAfterCreateInterface[CodeProcedureFixRoleType]
	OnAfterCodeProcedureFixRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeProcedureFixRoleType]
	OnAfterCodeProcedureFixRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeProcedureFixRoleType]

	CodeProcedurePhaseTypes                map[*CodeProcedurePhaseType]struct{}
	CodeProcedurePhaseTypes_instance       map[*CodeProcedurePhaseType]*CodeProcedurePhaseType
	CodeProcedurePhaseTypes_mapString      map[string]*CodeProcedurePhaseType
	CodeProcedurePhaseTypeOrder            uint
	CodeProcedurePhaseType_stagedOrder     map[*CodeProcedurePhaseType]uint
	CodeProcedurePhaseType_orderStaged     map[uint]*CodeProcedurePhaseType
	CodeProcedurePhaseTypes_reference      map[*CodeProcedurePhaseType]*CodeProcedurePhaseType
	CodeProcedurePhaseTypes_referenceOrder map[*CodeProcedurePhaseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProcedurePhaseTypeCreateCallback GongOnAfterCreateInterface[CodeProcedurePhaseType]
	OnAfterCodeProcedurePhaseTypeUpdateCallback GongOnAfterUpdateInterface[CodeProcedurePhaseType]
	OnAfterCodeProcedurePhaseTypeDeleteCallback GongOnAfterDeleteInterface[CodeProcedurePhaseType]

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

	CodeReferencePathIdentifierTypes                map[*CodeReferencePathIdentifierType]struct{}
	CodeReferencePathIdentifierTypes_instance       map[*CodeReferencePathIdentifierType]*CodeReferencePathIdentifierType
	CodeReferencePathIdentifierTypes_mapString      map[string]*CodeReferencePathIdentifierType
	CodeReferencePathIdentifierTypeOrder            uint
	CodeReferencePathIdentifierType_stagedOrder     map[*CodeReferencePathIdentifierType]uint
	CodeReferencePathIdentifierType_orderStaged     map[uint]*CodeReferencePathIdentifierType
	CodeReferencePathIdentifierTypes_reference      map[*CodeReferencePathIdentifierType]*CodeReferencePathIdentifierType
	CodeReferencePathIdentifierTypes_referenceOrder map[*CodeReferencePathIdentifierType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeReferencePathIdentifierTypeCreateCallback GongOnAfterCreateInterface[CodeReferencePathIdentifierType]
	OnAfterCodeReferencePathIdentifierTypeUpdateCallback GongOnAfterUpdateInterface[CodeReferencePathIdentifierType]
	OnAfterCodeReferencePathIdentifierTypeDeleteCallback GongOnAfterDeleteInterface[CodeReferencePathIdentifierType]

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

	CodeRelativePositionTypes                map[*CodeRelativePositionType]struct{}
	CodeRelativePositionTypes_instance       map[*CodeRelativePositionType]*CodeRelativePositionType
	CodeRelativePositionTypes_mapString      map[string]*CodeRelativePositionType
	CodeRelativePositionTypeOrder            uint
	CodeRelativePositionType_stagedOrder     map[*CodeRelativePositionType]uint
	CodeRelativePositionType_orderStaged     map[uint]*CodeRelativePositionType
	CodeRelativePositionTypes_reference      map[*CodeRelativePositionType]*CodeRelativePositionType
	CodeRelativePositionTypes_referenceOrder map[*CodeRelativePositionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRelativePositionTypeCreateCallback GongOnAfterCreateInterface[CodeRelativePositionType]
	OnAfterCodeRelativePositionTypeUpdateCallback GongOnAfterUpdateInterface[CodeRelativePositionType]
	OnAfterCodeRelativePositionTypeDeleteCallback GongOnAfterDeleteInterface[CodeRelativePositionType]

	CodeRouteIndicatorTypes                map[*CodeRouteIndicatorType]struct{}
	CodeRouteIndicatorTypes_instance       map[*CodeRouteIndicatorType]*CodeRouteIndicatorType
	CodeRouteIndicatorTypes_mapString      map[string]*CodeRouteIndicatorType
	CodeRouteIndicatorTypeOrder            uint
	CodeRouteIndicatorType_stagedOrder     map[*CodeRouteIndicatorType]uint
	CodeRouteIndicatorType_orderStaged     map[uint]*CodeRouteIndicatorType
	CodeRouteIndicatorTypes_reference      map[*CodeRouteIndicatorType]*CodeRouteIndicatorType
	CodeRouteIndicatorTypes_referenceOrder map[*CodeRouteIndicatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRouteIndicatorTypeCreateCallback GongOnAfterCreateInterface[CodeRouteIndicatorType]
	OnAfterCodeRouteIndicatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeRouteIndicatorType]
	OnAfterCodeRouteIndicatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeRouteIndicatorType]

	CodeSafeAltitudeTypes                map[*CodeSafeAltitudeType]struct{}
	CodeSafeAltitudeTypes_instance       map[*CodeSafeAltitudeType]*CodeSafeAltitudeType
	CodeSafeAltitudeTypes_mapString      map[string]*CodeSafeAltitudeType
	CodeSafeAltitudeTypeOrder            uint
	CodeSafeAltitudeType_stagedOrder     map[*CodeSafeAltitudeType]uint
	CodeSafeAltitudeType_orderStaged     map[uint]*CodeSafeAltitudeType
	CodeSafeAltitudeTypes_reference      map[*CodeSafeAltitudeType]*CodeSafeAltitudeType
	CodeSafeAltitudeTypes_referenceOrder map[*CodeSafeAltitudeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSafeAltitudeTypeCreateCallback GongOnAfterCreateInterface[CodeSafeAltitudeType]
	OnAfterCodeSafeAltitudeTypeUpdateCallback GongOnAfterUpdateInterface[CodeSafeAltitudeType]
	OnAfterCodeSafeAltitudeTypeDeleteCallback GongOnAfterDeleteInterface[CodeSafeAltitudeType]

	CodeSegmentPathTypes                map[*CodeSegmentPathType]struct{}
	CodeSegmentPathTypes_instance       map[*CodeSegmentPathType]*CodeSegmentPathType
	CodeSegmentPathTypes_mapString      map[string]*CodeSegmentPathType
	CodeSegmentPathTypeOrder            uint
	CodeSegmentPathType_stagedOrder     map[*CodeSegmentPathType]uint
	CodeSegmentPathType_orderStaged     map[uint]*CodeSegmentPathType
	CodeSegmentPathTypes_reference      map[*CodeSegmentPathType]*CodeSegmentPathType
	CodeSegmentPathTypes_referenceOrder map[*CodeSegmentPathType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSegmentPathTypeCreateCallback GongOnAfterCreateInterface[CodeSegmentPathType]
	OnAfterCodeSegmentPathTypeUpdateCallback GongOnAfterUpdateInterface[CodeSegmentPathType]
	OnAfterCodeSegmentPathTypeDeleteCallback GongOnAfterDeleteInterface[CodeSegmentPathType]

	CodeSegmentTerminationTypes                map[*CodeSegmentTerminationType]struct{}
	CodeSegmentTerminationTypes_instance       map[*CodeSegmentTerminationType]*CodeSegmentTerminationType
	CodeSegmentTerminationTypes_mapString      map[string]*CodeSegmentTerminationType
	CodeSegmentTerminationTypeOrder            uint
	CodeSegmentTerminationType_stagedOrder     map[*CodeSegmentTerminationType]uint
	CodeSegmentTerminationType_orderStaged     map[uint]*CodeSegmentTerminationType
	CodeSegmentTerminationTypes_reference      map[*CodeSegmentTerminationType]*CodeSegmentTerminationType
	CodeSegmentTerminationTypes_referenceOrder map[*CodeSegmentTerminationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSegmentTerminationTypeCreateCallback GongOnAfterCreateInterface[CodeSegmentTerminationType]
	OnAfterCodeSegmentTerminationTypeUpdateCallback GongOnAfterUpdateInterface[CodeSegmentTerminationType]
	OnAfterCodeSegmentTerminationTypeDeleteCallback GongOnAfterDeleteInterface[CodeSegmentTerminationType]

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

	CodeTAATypes                map[*CodeTAAType]struct{}
	CodeTAATypes_instance       map[*CodeTAAType]*CodeTAAType
	CodeTAATypes_mapString      map[string]*CodeTAAType
	CodeTAATypeOrder            uint
	CodeTAAType_stagedOrder     map[*CodeTAAType]uint
	CodeTAAType_orderStaged     map[uint]*CodeTAAType
	CodeTAATypes_reference      map[*CodeTAAType]*CodeTAAType
	CodeTAATypes_referenceOrder map[*CodeTAAType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTAATypeCreateCallback GongOnAfterCreateInterface[CodeTAAType]
	OnAfterCodeTAATypeUpdateCallback GongOnAfterUpdateInterface[CodeTAAType]
	OnAfterCodeTAATypeDeleteCallback GongOnAfterDeleteInterface[CodeTAAType]

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

	CodeTrajectoryTypes                map[*CodeTrajectoryType]struct{}
	CodeTrajectoryTypes_instance       map[*CodeTrajectoryType]*CodeTrajectoryType
	CodeTrajectoryTypes_mapString      map[string]*CodeTrajectoryType
	CodeTrajectoryTypeOrder            uint
	CodeTrajectoryType_stagedOrder     map[*CodeTrajectoryType]uint
	CodeTrajectoryType_orderStaged     map[uint]*CodeTrajectoryType
	CodeTrajectoryTypes_reference      map[*CodeTrajectoryType]*CodeTrajectoryType
	CodeTrajectoryTypes_referenceOrder map[*CodeTrajectoryType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTrajectoryTypeCreateCallback GongOnAfterCreateInterface[CodeTrajectoryType]
	OnAfterCodeTrajectoryTypeUpdateCallback GongOnAfterUpdateInterface[CodeTrajectoryType]
	OnAfterCodeTrajectoryTypeDeleteCallback GongOnAfterDeleteInterface[CodeTrajectoryType]

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

	TextSIDSTARDesignatorTypes                map[*TextSIDSTARDesignatorType]struct{}
	TextSIDSTARDesignatorTypes_instance       map[*TextSIDSTARDesignatorType]*TextSIDSTARDesignatorType
	TextSIDSTARDesignatorTypes_mapString      map[string]*TextSIDSTARDesignatorType
	TextSIDSTARDesignatorTypeOrder            uint
	TextSIDSTARDesignatorType_stagedOrder     map[*TextSIDSTARDesignatorType]uint
	TextSIDSTARDesignatorType_orderStaged     map[uint]*TextSIDSTARDesignatorType
	TextSIDSTARDesignatorTypes_reference      map[*TextSIDSTARDesignatorType]*TextSIDSTARDesignatorType
	TextSIDSTARDesignatorTypes_referenceOrder map[*TextSIDSTARDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterTextSIDSTARDesignatorTypeCreateCallback GongOnAfterCreateInterface[TextSIDSTARDesignatorType]
	OnAfterTextSIDSTARDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[TextSIDSTARDesignatorType]
	OnAfterTextSIDSTARDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[TextSIDSTARDesignatorType]

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

	ValAlarmLimitTypes                map[*ValAlarmLimitType]struct{}
	ValAlarmLimitTypes_instance       map[*ValAlarmLimitType]*ValAlarmLimitType
	ValAlarmLimitTypes_mapString      map[string]*ValAlarmLimitType
	ValAlarmLimitTypeOrder            uint
	ValAlarmLimitType_stagedOrder     map[*ValAlarmLimitType]uint
	ValAlarmLimitType_orderStaged     map[uint]*ValAlarmLimitType
	ValAlarmLimitTypes_reference      map[*ValAlarmLimitType]*ValAlarmLimitType
	ValAlarmLimitTypes_referenceOrder map[*ValAlarmLimitType]uint

	// insertion point for slice of pointers maps
	OnAfterValAlarmLimitTypeCreateCallback GongOnAfterCreateInterface[ValAlarmLimitType]
	OnAfterValAlarmLimitTypeUpdateCallback GongOnAfterUpdateInterface[ValAlarmLimitType]
	OnAfterValAlarmLimitTypeDeleteCallback GongOnAfterDeleteInterface[ValAlarmLimitType]

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

	ValChannelNumberTypes                map[*ValChannelNumberType]struct{}
	ValChannelNumberTypes_instance       map[*ValChannelNumberType]*ValChannelNumberType
	ValChannelNumberTypes_mapString      map[string]*ValChannelNumberType
	ValChannelNumberTypeOrder            uint
	ValChannelNumberType_stagedOrder     map[*ValChannelNumberType]uint
	ValChannelNumberType_orderStaged     map[uint]*ValChannelNumberType
	ValChannelNumberTypes_reference      map[*ValChannelNumberType]*ValChannelNumberType
	ValChannelNumberTypes_referenceOrder map[*ValChannelNumberType]uint

	// insertion point for slice of pointers maps
	OnAfterValChannelNumberTypeCreateCallback GongOnAfterCreateInterface[ValChannelNumberType]
	OnAfterValChannelNumberTypeUpdateCallback GongOnAfterUpdateInterface[ValChannelNumberType]
	OnAfterValChannelNumberTypeDeleteCallback GongOnAfterDeleteInterface[ValChannelNumberType]

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

	ValHexTypes                map[*ValHexType]struct{}
	ValHexTypes_instance       map[*ValHexType]*ValHexType
	ValHexTypes_mapString      map[string]*ValHexType
	ValHexTypeOrder            uint
	ValHexType_stagedOrder     map[*ValHexType]uint
	ValHexType_orderStaged     map[uint]*ValHexType
	ValHexTypes_reference      map[*ValHexType]*ValHexType
	ValHexTypes_referenceOrder map[*ValHexType]uint

	// insertion point for slice of pointers maps
	OnAfterValHexTypeCreateCallback GongOnAfterCreateInterface[ValHexType]
	OnAfterValHexTypeUpdateCallback GongOnAfterUpdateInterface[ValHexType]
	OnAfterValHexTypeDeleteCallback GongOnAfterDeleteInterface[ValHexType]

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

	ValTemperatureTypes                map[*ValTemperatureType]struct{}
	ValTemperatureTypes_instance       map[*ValTemperatureType]*ValTemperatureType
	ValTemperatureTypes_mapString      map[string]*ValTemperatureType
	ValTemperatureTypeOrder            uint
	ValTemperatureType_stagedOrder     map[*ValTemperatureType]uint
	ValTemperatureType_orderStaged     map[uint]*ValTemperatureType
	ValTemperatureTypes_reference      map[*ValTemperatureType]*ValTemperatureType
	ValTemperatureTypes_referenceOrder map[*ValTemperatureType]uint

	// insertion point for slice of pointers maps
	OnAfterValTemperatureTypeCreateCallback GongOnAfterCreateInterface[ValTemperatureType]
	OnAfterValTemperatureTypeUpdateCallback GongOnAfterUpdateInterface[ValTemperatureType]
	OnAfterValTemperatureTypeDeleteCallback GongOnAfterDeleteInterface[ValTemperatureType]

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

	__gong__clearReferences(&stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_instance, &stage.CodeAircraftCategoryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_instance, &stage.CodeAircraftEngineNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_instance, &stage.CodeAircraftEngineTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_instance, &stage.CodeAircraftICAOTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_instance, &stage.CodeAircraftTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_instance, &stage.CodeAircraftWingspanClassTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeAdjustmentTypes_reference, &stage.CodeAltitudeAdjustmentTypes_instance, &stage.CodeAltitudeAdjustmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_instance, &stage.CodeAltitudeUseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachEquipmentAdditionalTypes_reference, &stage.CodeApproachEquipmentAdditionalTypes_instance, &stage.CodeApproachEquipmentAdditionalTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachGuidanceTypes_reference, &stage.CodeApproachGuidanceTypes_instance, &stage.CodeApproachGuidanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachPrefixTypes_reference, &stage.CodeApproachPrefixTypes_instance, &stage.CodeApproachPrefixTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachTypes_reference, &stage.CodeApproachTypes_instance, &stage.CodeApproachTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApprovalTypes_reference, &stage.CodeApprovalTypes_instance, &stage.CodeApprovalTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeArcDirectionTypes_reference, &stage.CodeArcDirectionTypes_instance, &stage.CodeArcDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeBearingTypes_reference, &stage.CodeBearingTypes_instance, &stage.CodeBearingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_instance, &stage.CodeCommunicationModeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCourseTypes_reference, &stage.CodeCourseTypes_instance, &stage.CodeCourseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDesignStandardTypes_reference, &stage.CodeDesignStandardTypes_instance, &stage.CodeDesignStandardTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDesignatedPointDesignatorTypes_reference, &stage.CodeDesignatedPointDesignatorTypes_instance, &stage.CodeDesignatedPointDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionReferenceTypes_reference, &stage.CodeDirectionReferenceTypes_instance, &stage.CodeDirectionReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_instance, &stage.CodeDirectionTurnTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_instance, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEquipmentUnavailableTypes_reference, &stage.CodeEquipmentUnavailableTypes_instance, &stage.CodeEquipmentUnavailableTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFinalGuidanceTypes_reference, &stage.CodeFinalGuidanceTypes_instance, &stage.CodeFinalGuidanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFreeFlightTypes_reference, &stage.CodeFreeFlightTypes_instance, &stage.CodeFreeFlightTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeHeightReferenceTypes_reference, &stage.CodeHeightReferenceTypes_instance, &stage.CodeHeightReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeHoldingUsageTypes_reference, &stage.CodeHoldingUsageTypes_instance, &stage.CodeHoldingUsageTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeHoldingUseTypes_reference, &stage.CodeHoldingUseTypes_instance, &stage.CodeHoldingUseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeICAOCountryTypes_reference, &stage.CodeICAOCountryTypes_instance, &stage.CodeICAOCountryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryRoutePointTypes_reference, &stage.CodeMilitaryRoutePointTypes_instance, &stage.CodeMilitaryRoutePointTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMinimaFinalApproachPathTypes_reference, &stage.CodeMinimaFinalApproachPathTypes_instance, &stage.CodeMinimaFinalApproachPathTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMinimumAltitudeTypes_reference, &stage.CodeMinimumAltitudeTypes_instance, &stage.CodeMinimumAltitudeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMinimumHeightTypes_reference, &stage.CodeMinimumHeightTypes_instance, &stage.CodeMinimumHeightTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMissedApproachTypes_reference, &stage.CodeMissedApproachTypes_instance, &stage.CodeMissedApproachTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationAreaRestrictionTypes_reference, &stage.CodeNavigationAreaRestrictionTypes_instance, &stage.CodeNavigationAreaRestrictionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationAreaTypes_reference, &stage.CodeNavigationAreaTypes_instance, &stage.CodeNavigationAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_instance, &stage.CodeNavigationEquipmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_instance, &stage.CodeNavigationSpecificationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_instance, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstructionIdSurfaceZoneTypes_reference, &stage.CodeObstructionIdSurfaceZoneTypes_instance, &stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProcedureAvailabilityTypes_reference, &stage.CodeProcedureAvailabilityTypes_instance, &stage.CodeProcedureAvailabilityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProcedureCodingStandardTypes_reference, &stage.CodeProcedureCodingStandardTypes_instance, &stage.CodeProcedureCodingStandardTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProcedureDistanceTypes_reference, &stage.CodeProcedureDistanceTypes_instance, &stage.CodeProcedureDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProcedureFixRoleTypes_reference, &stage.CodeProcedureFixRoleTypes_instance, &stage.CodeProcedureFixRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProcedurePhaseTypes_reference, &stage.CodeProcedurePhaseTypes_instance, &stage.CodeProcedurePhaseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRNPTypes_reference, &stage.CodeRNPTypes_instance, &stage.CodeRNPTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMPointRoleTypes_reference, &stage.CodeRVSMPointRoleTypes_instance, &stage.CodeRVSMPointRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_instance, &stage.CodeRVSMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeReferencePathIdentifierTypes_reference, &stage.CodeReferencePathIdentifierTypes_instance, &stage.CodeReferencePathIdentifierTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeReferenceRoleTypes_reference, &stage.CodeReferenceRoleTypes_instance, &stage.CodeReferenceRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRelativePositionTypes_reference, &stage.CodeRelativePositionTypes_instance, &stage.CodeRelativePositionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRouteIndicatorTypes_reference, &stage.CodeRouteIndicatorTypes_instance, &stage.CodeRouteIndicatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSafeAltitudeTypes_reference, &stage.CodeSafeAltitudeTypes_instance, &stage.CodeSafeAltitudeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSegmentPathTypes_reference, &stage.CodeSegmentPathTypes_instance, &stage.CodeSegmentPathTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSegmentTerminationTypes_reference, &stage.CodeSegmentTerminationTypes_instance, &stage.CodeSegmentTerminationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSideTypes_reference, &stage.CodeSideTypes_instance, &stage.CodeSideTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpeedReferenceTypes_reference, &stage.CodeSpeedReferenceTypes_instance, &stage.CodeSpeedReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTAATypes_reference, &stage.CodeTAATypes_instance, &stage.CodeTAATypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTrajectoryTypes_reference, &stage.CodeTrajectoryTypes_instance, &stage.CodeTrajectoryTypes_referenceOrder)

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

	__gong__clearReferences(&stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_instance, &stage.TextInstructionTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNameTypes_reference, &stage.TextNameTypes_instance, &stage.TextNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNoteTypes_reference, &stage.TextNoteTypes_instance, &stage.TextNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_instance, &stage.TextPropertyNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextSIDSTARDesignatorTypes_reference, &stage.TextSIDSTARDesignatorTypes_instance, &stage.TextSIDSTARDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValAlarmLimitTypes_reference, &stage.ValAlarmLimitTypes_instance, &stage.ValAlarmLimitTypes_referenceOrder)

	__gong__clearReferences(&stage.ValAngleTypes_reference, &stage.ValAngleTypes_instance, &stage.ValAngleTypes_referenceOrder)

	__gong__clearReferences(&stage.ValBearingTypes_reference, &stage.ValBearingTypes_instance, &stage.ValBearingTypes_referenceOrder)

	__gong__clearReferences(&stage.ValChannelNumberTypes_reference, &stage.ValChannelNumberTypes_instance, &stage.ValChannelNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValHexTypes_reference, &stage.ValHexTypes_instance, &stage.ValHexTypes_referenceOrder)

	__gong__clearReferences(&stage.ValSlopeTypes_reference, &stage.ValSlopeTypes_instance, &stage.ValSlopeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_instance, &stage.ValSpeedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValTemperatureTypes_reference, &stage.ValTemperatureTypes_instance, &stage.ValTemperatureTypes_referenceOrder)

	__gong__clearReferences(&stage.ValWeightTypes_reference, &stage.ValWeightTypes_instance, &stage.ValWeightTypes_referenceOrder)

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

	stage.CodeAircraftCategoryTypeOrder = __gong__recomputeOrder(stage.CodeAircraftCategoryType_stagedOrder)

	stage.CodeAircraftEngineNumberTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineNumberType_stagedOrder)

	stage.CodeAircraftEngineTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineType_stagedOrder)

	stage.CodeAircraftICAOTypeOrder = __gong__recomputeOrder(stage.CodeAircraftICAOType_stagedOrder)

	stage.CodeAircraftTypeOrder = __gong__recomputeOrder(stage.CodeAircraftType_stagedOrder)

	stage.CodeAircraftWingspanClassTypeOrder = __gong__recomputeOrder(stage.CodeAircraftWingspanClassType_stagedOrder)

	stage.CodeAltitudeAdjustmentTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeAdjustmentType_stagedOrder)

	stage.CodeAltitudeUseTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeUseType_stagedOrder)

	stage.CodeApproachEquipmentAdditionalTypeOrder = __gong__recomputeOrder(stage.CodeApproachEquipmentAdditionalType_stagedOrder)

	stage.CodeApproachGuidanceTypeOrder = __gong__recomputeOrder(stage.CodeApproachGuidanceType_stagedOrder)

	stage.CodeApproachPrefixTypeOrder = __gong__recomputeOrder(stage.CodeApproachPrefixType_stagedOrder)

	stage.CodeApproachTypeOrder = __gong__recomputeOrder(stage.CodeApproachType_stagedOrder)

	stage.CodeApprovalTypeOrder = __gong__recomputeOrder(stage.CodeApprovalType_stagedOrder)

	stage.CodeArcDirectionTypeOrder = __gong__recomputeOrder(stage.CodeArcDirectionType_stagedOrder)

	stage.CodeBearingTypeOrder = __gong__recomputeOrder(stage.CodeBearingType_stagedOrder)

	stage.CodeCommunicationModeTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationModeType_stagedOrder)

	stage.CodeCourseTypeOrder = __gong__recomputeOrder(stage.CodeCourseType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeDesignStandardTypeOrder = __gong__recomputeOrder(stage.CodeDesignStandardType_stagedOrder)

	stage.CodeDesignatedPointDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeDesignatedPointDesignatorType_stagedOrder)

	stage.CodeDirectionReferenceTypeOrder = __gong__recomputeOrder(stage.CodeDirectionReferenceType_stagedOrder)

	stage.CodeDirectionTurnTypeOrder = __gong__recomputeOrder(stage.CodeDirectionTurnType_stagedOrder)

	stage.CodeEquipmentAntiCollisionTypeOrder = __gong__recomputeOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder)

	stage.CodeEquipmentUnavailableTypeOrder = __gong__recomputeOrder(stage.CodeEquipmentUnavailableType_stagedOrder)

	stage.CodeFinalGuidanceTypeOrder = __gong__recomputeOrder(stage.CodeFinalGuidanceType_stagedOrder)

	stage.CodeFreeFlightTypeOrder = __gong__recomputeOrder(stage.CodeFreeFlightType_stagedOrder)

	stage.CodeHeightReferenceTypeOrder = __gong__recomputeOrder(stage.CodeHeightReferenceType_stagedOrder)

	stage.CodeHoldingUsageTypeOrder = __gong__recomputeOrder(stage.CodeHoldingUsageType_stagedOrder)

	stage.CodeHoldingUseTypeOrder = __gong__recomputeOrder(stage.CodeHoldingUseType_stagedOrder)

	stage.CodeICAOCountryTypeOrder = __gong__recomputeOrder(stage.CodeICAOCountryType_stagedOrder)

	stage.CodeMilitaryRoutePointTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryRoutePointType_stagedOrder)

	stage.CodeMinimaFinalApproachPathTypeOrder = __gong__recomputeOrder(stage.CodeMinimaFinalApproachPathType_stagedOrder)

	stage.CodeMinimumAltitudeTypeOrder = __gong__recomputeOrder(stage.CodeMinimumAltitudeType_stagedOrder)

	stage.CodeMinimumHeightTypeOrder = __gong__recomputeOrder(stage.CodeMinimumHeightType_stagedOrder)

	stage.CodeMissedApproachTypeOrder = __gong__recomputeOrder(stage.CodeMissedApproachType_stagedOrder)

	stage.CodeNavigationAreaRestrictionTypeOrder = __gong__recomputeOrder(stage.CodeNavigationAreaRestrictionType_stagedOrder)

	stage.CodeNavigationAreaTypeOrder = __gong__recomputeOrder(stage.CodeNavigationAreaType_stagedOrder)

	stage.CodeNavigationEquipmentTypeOrder = __gong__recomputeOrder(stage.CodeNavigationEquipmentType_stagedOrder)

	stage.CodeNavigationSpecificationTypeOrder = __gong__recomputeOrder(stage.CodeNavigationSpecificationType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeObstacleAssessmentSurfaceTypeOrder = __gong__recomputeOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder)

	stage.CodeObstructionIdSurfaceZoneTypeOrder = __gong__recomputeOrder(stage.CodeObstructionIdSurfaceZoneType_stagedOrder)

	stage.CodeProcedureAvailabilityTypeOrder = __gong__recomputeOrder(stage.CodeProcedureAvailabilityType_stagedOrder)

	stage.CodeProcedureCodingStandardTypeOrder = __gong__recomputeOrder(stage.CodeProcedureCodingStandardType_stagedOrder)

	stage.CodeProcedureDistanceTypeOrder = __gong__recomputeOrder(stage.CodeProcedureDistanceType_stagedOrder)

	stage.CodeProcedureFixRoleTypeOrder = __gong__recomputeOrder(stage.CodeProcedureFixRoleType_stagedOrder)

	stage.CodeProcedurePhaseTypeOrder = __gong__recomputeOrder(stage.CodeProcedurePhaseType_stagedOrder)

	stage.CodeRNPTypeOrder = __gong__recomputeOrder(stage.CodeRNPType_stagedOrder)

	stage.CodeRVSMPointRoleTypeOrder = __gong__recomputeOrder(stage.CodeRVSMPointRoleType_stagedOrder)

	stage.CodeRVSMTypeOrder = __gong__recomputeOrder(stage.CodeRVSMType_stagedOrder)

	stage.CodeReferencePathIdentifierTypeOrder = __gong__recomputeOrder(stage.CodeReferencePathIdentifierType_stagedOrder)

	stage.CodeReferenceRoleTypeOrder = __gong__recomputeOrder(stage.CodeReferenceRoleType_stagedOrder)

	stage.CodeRelativePositionTypeOrder = __gong__recomputeOrder(stage.CodeRelativePositionType_stagedOrder)

	stage.CodeRouteIndicatorTypeOrder = __gong__recomputeOrder(stage.CodeRouteIndicatorType_stagedOrder)

	stage.CodeSafeAltitudeTypeOrder = __gong__recomputeOrder(stage.CodeSafeAltitudeType_stagedOrder)

	stage.CodeSegmentPathTypeOrder = __gong__recomputeOrder(stage.CodeSegmentPathType_stagedOrder)

	stage.CodeSegmentTerminationTypeOrder = __gong__recomputeOrder(stage.CodeSegmentTerminationType_stagedOrder)

	stage.CodeSideTypeOrder = __gong__recomputeOrder(stage.CodeSideType_stagedOrder)

	stage.CodeSpeedReferenceTypeOrder = __gong__recomputeOrder(stage.CodeSpeedReferenceType_stagedOrder)

	stage.CodeTAATypeOrder = __gong__recomputeOrder(stage.CodeTAAType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeTrajectoryTypeOrder = __gong__recomputeOrder(stage.CodeTrajectoryType_stagedOrder)

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

	stage.TextInstructionTypeOrder = __gong__recomputeOrder(stage.TextInstructionType_stagedOrder)

	stage.TextNameTypeOrder = __gong__recomputeOrder(stage.TextNameType_stagedOrder)

	stage.TextNoteTypeOrder = __gong__recomputeOrder(stage.TextNoteType_stagedOrder)

	stage.TextPropertyNameTypeOrder = __gong__recomputeOrder(stage.TextPropertyNameType_stagedOrder)

	stage.TextSIDSTARDesignatorTypeOrder = __gong__recomputeOrder(stage.TextSIDSTARDesignatorType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.ValAlarmLimitTypeOrder = __gong__recomputeOrder(stage.ValAlarmLimitType_stagedOrder)

	stage.ValAngleTypeOrder = __gong__recomputeOrder(stage.ValAngleType_stagedOrder)

	stage.ValBearingTypeOrder = __gong__recomputeOrder(stage.ValBearingType_stagedOrder)

	stage.ValChannelNumberTypeOrder = __gong__recomputeOrder(stage.ValChannelNumberType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValHexTypeOrder = __gong__recomputeOrder(stage.ValHexType_stagedOrder)

	stage.ValSlopeTypeOrder = __gong__recomputeOrder(stage.ValSlopeType_stagedOrder)

	stage.ValSpeedTypeOrder = __gong__recomputeOrder(stage.ValSpeedType_stagedOrder)

	stage.ValTemperatureTypeOrder = __gong__recomputeOrder(stage.ValTemperatureType_stagedOrder)

	stage.ValWeightTypeOrder = __gong__recomputeOrder(stage.ValWeightType_stagedOrder)

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

func (*CodeApproachEquipmentAdditionalType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalType_stagedOrder)
}

func (*CodeApproachEquipmentAdditionalType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachEquipmentAdditionalType_orderStaged[order]
}

func (*CodeApproachEquipmentAdditionalType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachEquipmentAdditionalTypes_mapString
}

func (*CodeApproachEquipmentAdditionalType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachEquipmentAdditionalTypes
}

func (*CodeApproachEquipmentAdditionalType) GongNewInstance() any {
	return new(CodeApproachEquipmentAdditionalType)
}

func (*CodeApproachGuidanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceType_stagedOrder)
}

func (*CodeApproachGuidanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachGuidanceType_orderStaged[order]
}

func (*CodeApproachGuidanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachGuidanceTypes_mapString
}

func (*CodeApproachGuidanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachGuidanceTypes
}

func (*CodeApproachGuidanceType) GongNewInstance() any {
	return new(CodeApproachGuidanceType)
}

func (*CodeApproachPrefixType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixType_stagedOrder)
}

func (*CodeApproachPrefixType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachPrefixType_orderStaged[order]
}

func (*CodeApproachPrefixType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachPrefixTypes_mapString
}

func (*CodeApproachPrefixType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachPrefixTypes
}

func (*CodeApproachPrefixType) GongNewInstance() any {
	return new(CodeApproachPrefixType)
}

func (*CodeApproachType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachTypes, stage.CodeApproachType_stagedOrder)
}

func (*CodeApproachType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachType_orderStaged[order]
}

func (*CodeApproachType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachTypes_mapString
}

func (*CodeApproachType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachTypes
}

func (*CodeApproachType) GongNewInstance() any {
	return new(CodeApproachType)
}

func (*CodeApprovalType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApprovalTypes, stage.CodeApprovalType_stagedOrder)
}

func (*CodeApprovalType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApprovalType_orderStaged[order]
}

func (*CodeApprovalType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApprovalTypes_mapString
}

func (*CodeApprovalType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApprovalTypes
}

func (*CodeApprovalType) GongNewInstance() any {
	return new(CodeApprovalType)
}

func (*CodeArcDirectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeArcDirectionTypes, stage.CodeArcDirectionType_stagedOrder)
}

func (*CodeArcDirectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeArcDirectionType_orderStaged[order]
}

func (*CodeArcDirectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeArcDirectionTypes_mapString
}

func (*CodeArcDirectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeArcDirectionTypes
}

func (*CodeArcDirectionType) GongNewInstance() any {
	return new(CodeArcDirectionType)
}

func (*CodeBearingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeBearingTypes, stage.CodeBearingType_stagedOrder)
}

func (*CodeBearingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeBearingType_orderStaged[order]
}

func (*CodeBearingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeBearingTypes_mapString
}

func (*CodeBearingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeBearingTypes
}

func (*CodeBearingType) GongNewInstance() any {
	return new(CodeBearingType)
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

func (*CodeDesignStandardType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDesignStandardTypes, stage.CodeDesignStandardType_stagedOrder)
}

func (*CodeDesignStandardType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDesignStandardType_orderStaged[order]
}

func (*CodeDesignStandardType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDesignStandardTypes_mapString
}

func (*CodeDesignStandardType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDesignStandardTypes
}

func (*CodeDesignStandardType) GongNewInstance() any {
	return new(CodeDesignStandardType)
}

func (*CodeDesignatedPointDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorType_stagedOrder)
}

func (*CodeDesignatedPointDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDesignatedPointDesignatorType_orderStaged[order]
}

func (*CodeDesignatedPointDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDesignatedPointDesignatorTypes_mapString
}

func (*CodeDesignatedPointDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDesignatedPointDesignatorTypes
}

func (*CodeDesignatedPointDesignatorType) GongNewInstance() any {
	return new(CodeDesignatedPointDesignatorType)
}

func (*CodeDirectionReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceType_stagedOrder)
}

func (*CodeDirectionReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDirectionReferenceType_orderStaged[order]
}

func (*CodeDirectionReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDirectionReferenceTypes_mapString
}

func (*CodeDirectionReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDirectionReferenceTypes
}

func (*CodeDirectionReferenceType) GongNewInstance() any {
	return new(CodeDirectionReferenceType)
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

func (*CodeEquipmentUnavailableType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableType_stagedOrder)
}

func (*CodeEquipmentUnavailableType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeEquipmentUnavailableType_orderStaged[order]
}

func (*CodeEquipmentUnavailableType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeEquipmentUnavailableTypes_mapString
}

func (*CodeEquipmentUnavailableType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeEquipmentUnavailableTypes
}

func (*CodeEquipmentUnavailableType) GongNewInstance() any {
	return new(CodeEquipmentUnavailableType)
}

func (*CodeFinalGuidanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceType_stagedOrder)
}

func (*CodeFinalGuidanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFinalGuidanceType_orderStaged[order]
}

func (*CodeFinalGuidanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFinalGuidanceTypes_mapString
}

func (*CodeFinalGuidanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFinalGuidanceTypes
}

func (*CodeFinalGuidanceType) GongNewInstance() any {
	return new(CodeFinalGuidanceType)
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

func (*CodeHeightReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceType_stagedOrder)
}

func (*CodeHeightReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeHeightReferenceType_orderStaged[order]
}

func (*CodeHeightReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeHeightReferenceTypes_mapString
}

func (*CodeHeightReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeHeightReferenceTypes
}

func (*CodeHeightReferenceType) GongNewInstance() any {
	return new(CodeHeightReferenceType)
}

func (*CodeHoldingUsageType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageType_stagedOrder)
}

func (*CodeHoldingUsageType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeHoldingUsageType_orderStaged[order]
}

func (*CodeHoldingUsageType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeHoldingUsageTypes_mapString
}

func (*CodeHoldingUsageType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeHoldingUsageTypes
}

func (*CodeHoldingUsageType) GongNewInstance() any {
	return new(CodeHoldingUsageType)
}

func (*CodeHoldingUseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeHoldingUseTypes, stage.CodeHoldingUseType_stagedOrder)
}

func (*CodeHoldingUseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeHoldingUseType_orderStaged[order]
}

func (*CodeHoldingUseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeHoldingUseTypes_mapString
}

func (*CodeHoldingUseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeHoldingUseTypes
}

func (*CodeHoldingUseType) GongNewInstance() any {
	return new(CodeHoldingUseType)
}

func (*CodeICAOCountryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeICAOCountryTypes, stage.CodeICAOCountryType_stagedOrder)
}

func (*CodeICAOCountryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeICAOCountryType_orderStaged[order]
}

func (*CodeICAOCountryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeICAOCountryTypes_mapString
}

func (*CodeICAOCountryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeICAOCountryTypes
}

func (*CodeICAOCountryType) GongNewInstance() any {
	return new(CodeICAOCountryType)
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

func (*CodeMinimaFinalApproachPathType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathType_stagedOrder)
}

func (*CodeMinimaFinalApproachPathType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMinimaFinalApproachPathType_orderStaged[order]
}

func (*CodeMinimaFinalApproachPathType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMinimaFinalApproachPathTypes_mapString
}

func (*CodeMinimaFinalApproachPathType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMinimaFinalApproachPathTypes
}

func (*CodeMinimaFinalApproachPathType) GongNewInstance() any {
	return new(CodeMinimaFinalApproachPathType)
}

func (*CodeMinimumAltitudeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeType_stagedOrder)
}

func (*CodeMinimumAltitudeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMinimumAltitudeType_orderStaged[order]
}

func (*CodeMinimumAltitudeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMinimumAltitudeTypes_mapString
}

func (*CodeMinimumAltitudeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMinimumAltitudeTypes
}

func (*CodeMinimumAltitudeType) GongNewInstance() any {
	return new(CodeMinimumAltitudeType)
}

func (*CodeMinimumHeightType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightType_stagedOrder)
}

func (*CodeMinimumHeightType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMinimumHeightType_orderStaged[order]
}

func (*CodeMinimumHeightType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMinimumHeightTypes_mapString
}

func (*CodeMinimumHeightType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMinimumHeightTypes
}

func (*CodeMinimumHeightType) GongNewInstance() any {
	return new(CodeMinimumHeightType)
}

func (*CodeMissedApproachType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMissedApproachTypes, stage.CodeMissedApproachType_stagedOrder)
}

func (*CodeMissedApproachType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMissedApproachType_orderStaged[order]
}

func (*CodeMissedApproachType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMissedApproachTypes_mapString
}

func (*CodeMissedApproachType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMissedApproachTypes
}

func (*CodeMissedApproachType) GongNewInstance() any {
	return new(CodeMissedApproachType)
}

func (*CodeNavigationAreaRestrictionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionType_stagedOrder)
}

func (*CodeNavigationAreaRestrictionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavigationAreaRestrictionType_orderStaged[order]
}

func (*CodeNavigationAreaRestrictionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavigationAreaRestrictionTypes_mapString
}

func (*CodeNavigationAreaRestrictionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavigationAreaRestrictionTypes
}

func (*CodeNavigationAreaRestrictionType) GongNewInstance() any {
	return new(CodeNavigationAreaRestrictionType)
}

func (*CodeNavigationAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaType_stagedOrder)
}

func (*CodeNavigationAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavigationAreaType_orderStaged[order]
}

func (*CodeNavigationAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavigationAreaTypes_mapString
}

func (*CodeNavigationAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavigationAreaTypes
}

func (*CodeNavigationAreaType) GongNewInstance() any {
	return new(CodeNavigationAreaType)
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

func (*CodeProcedureAvailabilityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityType_stagedOrder)
}

func (*CodeProcedureAvailabilityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProcedureAvailabilityType_orderStaged[order]
}

func (*CodeProcedureAvailabilityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProcedureAvailabilityTypes_mapString
}

func (*CodeProcedureAvailabilityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProcedureAvailabilityTypes
}

func (*CodeProcedureAvailabilityType) GongNewInstance() any {
	return new(CodeProcedureAvailabilityType)
}

func (*CodeProcedureCodingStandardType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardType_stagedOrder)
}

func (*CodeProcedureCodingStandardType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProcedureCodingStandardType_orderStaged[order]
}

func (*CodeProcedureCodingStandardType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProcedureCodingStandardTypes_mapString
}

func (*CodeProcedureCodingStandardType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProcedureCodingStandardTypes
}

func (*CodeProcedureCodingStandardType) GongNewInstance() any {
	return new(CodeProcedureCodingStandardType)
}

func (*CodeProcedureDistanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceType_stagedOrder)
}

func (*CodeProcedureDistanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProcedureDistanceType_orderStaged[order]
}

func (*CodeProcedureDistanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProcedureDistanceTypes_mapString
}

func (*CodeProcedureDistanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProcedureDistanceTypes
}

func (*CodeProcedureDistanceType) GongNewInstance() any {
	return new(CodeProcedureDistanceType)
}

func (*CodeProcedureFixRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleType_stagedOrder)
}

func (*CodeProcedureFixRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProcedureFixRoleType_orderStaged[order]
}

func (*CodeProcedureFixRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProcedureFixRoleTypes_mapString
}

func (*CodeProcedureFixRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProcedureFixRoleTypes
}

func (*CodeProcedureFixRoleType) GongNewInstance() any {
	return new(CodeProcedureFixRoleType)
}

func (*CodeProcedurePhaseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseType_stagedOrder)
}

func (*CodeProcedurePhaseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProcedurePhaseType_orderStaged[order]
}

func (*CodeProcedurePhaseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProcedurePhaseTypes_mapString
}

func (*CodeProcedurePhaseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProcedurePhaseTypes
}

func (*CodeProcedurePhaseType) GongNewInstance() any {
	return new(CodeProcedurePhaseType)
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

func (*CodeReferencePathIdentifierType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierType_stagedOrder)
}

func (*CodeReferencePathIdentifierType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeReferencePathIdentifierType_orderStaged[order]
}

func (*CodeReferencePathIdentifierType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeReferencePathIdentifierTypes_mapString
}

func (*CodeReferencePathIdentifierType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeReferencePathIdentifierTypes
}

func (*CodeReferencePathIdentifierType) GongNewInstance() any {
	return new(CodeReferencePathIdentifierType)
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

func (*CodeRelativePositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRelativePositionTypes, stage.CodeRelativePositionType_stagedOrder)
}

func (*CodeRelativePositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRelativePositionType_orderStaged[order]
}

func (*CodeRelativePositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRelativePositionTypes_mapString
}

func (*CodeRelativePositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRelativePositionTypes
}

func (*CodeRelativePositionType) GongNewInstance() any {
	return new(CodeRelativePositionType)
}

func (*CodeRouteIndicatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorType_stagedOrder)
}

func (*CodeRouteIndicatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRouteIndicatorType_orderStaged[order]
}

func (*CodeRouteIndicatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRouteIndicatorTypes_mapString
}

func (*CodeRouteIndicatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRouteIndicatorTypes
}

func (*CodeRouteIndicatorType) GongNewInstance() any {
	return new(CodeRouteIndicatorType)
}

func (*CodeSafeAltitudeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeType_stagedOrder)
}

func (*CodeSafeAltitudeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSafeAltitudeType_orderStaged[order]
}

func (*CodeSafeAltitudeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSafeAltitudeTypes_mapString
}

func (*CodeSafeAltitudeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSafeAltitudeTypes
}

func (*CodeSafeAltitudeType) GongNewInstance() any {
	return new(CodeSafeAltitudeType)
}

func (*CodeSegmentPathType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSegmentPathTypes, stage.CodeSegmentPathType_stagedOrder)
}

func (*CodeSegmentPathType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSegmentPathType_orderStaged[order]
}

func (*CodeSegmentPathType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSegmentPathTypes_mapString
}

func (*CodeSegmentPathType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSegmentPathTypes
}

func (*CodeSegmentPathType) GongNewInstance() any {
	return new(CodeSegmentPathType)
}

func (*CodeSegmentTerminationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationType_stagedOrder)
}

func (*CodeSegmentTerminationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSegmentTerminationType_orderStaged[order]
}

func (*CodeSegmentTerminationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSegmentTerminationTypes_mapString
}

func (*CodeSegmentTerminationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSegmentTerminationTypes
}

func (*CodeSegmentTerminationType) GongNewInstance() any {
	return new(CodeSegmentTerminationType)
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

func (*CodeTAAType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTAATypes, stage.CodeTAAType_stagedOrder)
}

func (*CodeTAAType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTAAType_orderStaged[order]
}

func (*CodeTAAType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTAATypes_mapString
}

func (*CodeTAAType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTAATypes
}

func (*CodeTAAType) GongNewInstance() any {
	return new(CodeTAAType)
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

func (*CodeTrajectoryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTrajectoryTypes, stage.CodeTrajectoryType_stagedOrder)
}

func (*CodeTrajectoryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTrajectoryType_orderStaged[order]
}

func (*CodeTrajectoryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTrajectoryTypes_mapString
}

func (*CodeTrajectoryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTrajectoryTypes
}

func (*CodeTrajectoryType) GongNewInstance() any {
	return new(CodeTrajectoryType)
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

func (*TextSIDSTARDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorType_stagedOrder)
}

func (*TextSIDSTARDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextSIDSTARDesignatorType_orderStaged[order]
}

func (*TextSIDSTARDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextSIDSTARDesignatorTypes_mapString
}

func (*TextSIDSTARDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextSIDSTARDesignatorTypes
}

func (*TextSIDSTARDesignatorType) GongNewInstance() any {
	return new(TextSIDSTARDesignatorType)
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

func (*ValAlarmLimitType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValAlarmLimitTypes, stage.ValAlarmLimitType_stagedOrder)
}

func (*ValAlarmLimitType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValAlarmLimitType_orderStaged[order]
}

func (*ValAlarmLimitType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValAlarmLimitTypes_mapString
}

func (*ValAlarmLimitType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValAlarmLimitTypes
}

func (*ValAlarmLimitType) GongNewInstance() any {
	return new(ValAlarmLimitType)
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

func (*ValChannelNumberType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValChannelNumberTypes, stage.ValChannelNumberType_stagedOrder)
}

func (*ValChannelNumberType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValChannelNumberType_orderStaged[order]
}

func (*ValChannelNumberType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValChannelNumberTypes_mapString
}

func (*ValChannelNumberType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValChannelNumberTypes
}

func (*ValChannelNumberType) GongNewInstance() any {
	return new(ValChannelNumberType)
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

func (*ValHexType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValHexTypes, stage.ValHexType_stagedOrder)
}

func (*ValHexType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValHexType_orderStaged[order]
}

func (*ValHexType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValHexTypes_mapString
}

func (*ValHexType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValHexTypes
}

func (*ValHexType) GongNewInstance() any {
	return new(ValHexType)
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

func (*ValTemperatureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValTemperatureTypes, stage.ValTemperatureType_stagedOrder)
}

func (*ValTemperatureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValTemperatureType_orderStaged[order]
}

func (*ValTemperatureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValTemperatureTypes_mapString
}

func (*ValTemperatureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValTemperatureTypes
}

func (*ValTemperatureType) GongNewInstance() any {
	return new(ValTemperatureType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
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

		CodeApproachEquipmentAdditionalTypes:           make(map[*CodeApproachEquipmentAdditionalType]struct{}),
		CodeApproachEquipmentAdditionalTypes_mapString: make(map[string]*CodeApproachEquipmentAdditionalType),

		CodeApproachGuidanceTypes:           make(map[*CodeApproachGuidanceType]struct{}),
		CodeApproachGuidanceTypes_mapString: make(map[string]*CodeApproachGuidanceType),

		CodeApproachPrefixTypes:           make(map[*CodeApproachPrefixType]struct{}),
		CodeApproachPrefixTypes_mapString: make(map[string]*CodeApproachPrefixType),

		CodeApproachTypes:           make(map[*CodeApproachType]struct{}),
		CodeApproachTypes_mapString: make(map[string]*CodeApproachType),

		CodeApprovalTypes:           make(map[*CodeApprovalType]struct{}),
		CodeApprovalTypes_mapString: make(map[string]*CodeApprovalType),

		CodeArcDirectionTypes:           make(map[*CodeArcDirectionType]struct{}),
		CodeArcDirectionTypes_mapString: make(map[string]*CodeArcDirectionType),

		CodeBearingTypes:           make(map[*CodeBearingType]struct{}),
		CodeBearingTypes_mapString: make(map[string]*CodeBearingType),

		CodeCommunicationModeTypes:           make(map[*CodeCommunicationModeType]struct{}),
		CodeCommunicationModeTypes_mapString: make(map[string]*CodeCommunicationModeType),

		CodeCourseTypes:           make(map[*CodeCourseType]struct{}),
		CodeCourseTypes_mapString: make(map[string]*CodeCourseType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeDesignStandardTypes:           make(map[*CodeDesignStandardType]struct{}),
		CodeDesignStandardTypes_mapString: make(map[string]*CodeDesignStandardType),

		CodeDesignatedPointDesignatorTypes:           make(map[*CodeDesignatedPointDesignatorType]struct{}),
		CodeDesignatedPointDesignatorTypes_mapString: make(map[string]*CodeDesignatedPointDesignatorType),

		CodeDirectionReferenceTypes:           make(map[*CodeDirectionReferenceType]struct{}),
		CodeDirectionReferenceTypes_mapString: make(map[string]*CodeDirectionReferenceType),

		CodeDirectionTurnTypes:           make(map[*CodeDirectionTurnType]struct{}),
		CodeDirectionTurnTypes_mapString: make(map[string]*CodeDirectionTurnType),

		CodeEquipmentAntiCollisionTypes:           make(map[*CodeEquipmentAntiCollisionType]struct{}),
		CodeEquipmentAntiCollisionTypes_mapString: make(map[string]*CodeEquipmentAntiCollisionType),

		CodeEquipmentUnavailableTypes:           make(map[*CodeEquipmentUnavailableType]struct{}),
		CodeEquipmentUnavailableTypes_mapString: make(map[string]*CodeEquipmentUnavailableType),

		CodeFinalGuidanceTypes:           make(map[*CodeFinalGuidanceType]struct{}),
		CodeFinalGuidanceTypes_mapString: make(map[string]*CodeFinalGuidanceType),

		CodeFreeFlightTypes:           make(map[*CodeFreeFlightType]struct{}),
		CodeFreeFlightTypes_mapString: make(map[string]*CodeFreeFlightType),

		CodeHeightReferenceTypes:           make(map[*CodeHeightReferenceType]struct{}),
		CodeHeightReferenceTypes_mapString: make(map[string]*CodeHeightReferenceType),

		CodeHoldingUsageTypes:           make(map[*CodeHoldingUsageType]struct{}),
		CodeHoldingUsageTypes_mapString: make(map[string]*CodeHoldingUsageType),

		CodeHoldingUseTypes:           make(map[*CodeHoldingUseType]struct{}),
		CodeHoldingUseTypes_mapString: make(map[string]*CodeHoldingUseType),

		CodeICAOCountryTypes:           make(map[*CodeICAOCountryType]struct{}),
		CodeICAOCountryTypes_mapString: make(map[string]*CodeICAOCountryType),

		CodeMilitaryRoutePointTypes:           make(map[*CodeMilitaryRoutePointType]struct{}),
		CodeMilitaryRoutePointTypes_mapString: make(map[string]*CodeMilitaryRoutePointType),

		CodeMinimaFinalApproachPathTypes:           make(map[*CodeMinimaFinalApproachPathType]struct{}),
		CodeMinimaFinalApproachPathTypes_mapString: make(map[string]*CodeMinimaFinalApproachPathType),

		CodeMinimumAltitudeTypes:           make(map[*CodeMinimumAltitudeType]struct{}),
		CodeMinimumAltitudeTypes_mapString: make(map[string]*CodeMinimumAltitudeType),

		CodeMinimumHeightTypes:           make(map[*CodeMinimumHeightType]struct{}),
		CodeMinimumHeightTypes_mapString: make(map[string]*CodeMinimumHeightType),

		CodeMissedApproachTypes:           make(map[*CodeMissedApproachType]struct{}),
		CodeMissedApproachTypes_mapString: make(map[string]*CodeMissedApproachType),

		CodeNavigationAreaRestrictionTypes:           make(map[*CodeNavigationAreaRestrictionType]struct{}),
		CodeNavigationAreaRestrictionTypes_mapString: make(map[string]*CodeNavigationAreaRestrictionType),

		CodeNavigationAreaTypes:           make(map[*CodeNavigationAreaType]struct{}),
		CodeNavigationAreaTypes_mapString: make(map[string]*CodeNavigationAreaType),

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

		CodeProcedureAvailabilityTypes:           make(map[*CodeProcedureAvailabilityType]struct{}),
		CodeProcedureAvailabilityTypes_mapString: make(map[string]*CodeProcedureAvailabilityType),

		CodeProcedureCodingStandardTypes:           make(map[*CodeProcedureCodingStandardType]struct{}),
		CodeProcedureCodingStandardTypes_mapString: make(map[string]*CodeProcedureCodingStandardType),

		CodeProcedureDistanceTypes:           make(map[*CodeProcedureDistanceType]struct{}),
		CodeProcedureDistanceTypes_mapString: make(map[string]*CodeProcedureDistanceType),

		CodeProcedureFixRoleTypes:           make(map[*CodeProcedureFixRoleType]struct{}),
		CodeProcedureFixRoleTypes_mapString: make(map[string]*CodeProcedureFixRoleType),

		CodeProcedurePhaseTypes:           make(map[*CodeProcedurePhaseType]struct{}),
		CodeProcedurePhaseTypes_mapString: make(map[string]*CodeProcedurePhaseType),

		CodeRNPTypes:           make(map[*CodeRNPType]struct{}),
		CodeRNPTypes_mapString: make(map[string]*CodeRNPType),

		CodeRVSMPointRoleTypes:           make(map[*CodeRVSMPointRoleType]struct{}),
		CodeRVSMPointRoleTypes_mapString: make(map[string]*CodeRVSMPointRoleType),

		CodeRVSMTypes:           make(map[*CodeRVSMType]struct{}),
		CodeRVSMTypes_mapString: make(map[string]*CodeRVSMType),

		CodeReferencePathIdentifierTypes:           make(map[*CodeReferencePathIdentifierType]struct{}),
		CodeReferencePathIdentifierTypes_mapString: make(map[string]*CodeReferencePathIdentifierType),

		CodeReferenceRoleTypes:           make(map[*CodeReferenceRoleType]struct{}),
		CodeReferenceRoleTypes_mapString: make(map[string]*CodeReferenceRoleType),

		CodeRelativePositionTypes:           make(map[*CodeRelativePositionType]struct{}),
		CodeRelativePositionTypes_mapString: make(map[string]*CodeRelativePositionType),

		CodeRouteIndicatorTypes:           make(map[*CodeRouteIndicatorType]struct{}),
		CodeRouteIndicatorTypes_mapString: make(map[string]*CodeRouteIndicatorType),

		CodeSafeAltitudeTypes:           make(map[*CodeSafeAltitudeType]struct{}),
		CodeSafeAltitudeTypes_mapString: make(map[string]*CodeSafeAltitudeType),

		CodeSegmentPathTypes:           make(map[*CodeSegmentPathType]struct{}),
		CodeSegmentPathTypes_mapString: make(map[string]*CodeSegmentPathType),

		CodeSegmentTerminationTypes:           make(map[*CodeSegmentTerminationType]struct{}),
		CodeSegmentTerminationTypes_mapString: make(map[string]*CodeSegmentTerminationType),

		CodeSideTypes:           make(map[*CodeSideType]struct{}),
		CodeSideTypes_mapString: make(map[string]*CodeSideType),

		CodeSpeedReferenceTypes:           make(map[*CodeSpeedReferenceType]struct{}),
		CodeSpeedReferenceTypes_mapString: make(map[string]*CodeSpeedReferenceType),

		CodeTAATypes:           make(map[*CodeTAAType]struct{}),
		CodeTAATypes_mapString: make(map[string]*CodeTAAType),

		CodeTimeEventCombinationTypes:           make(map[*CodeTimeEventCombinationType]struct{}),
		CodeTimeEventCombinationTypes_mapString: make(map[string]*CodeTimeEventCombinationType),

		CodeTimeEventTypes:           make(map[*CodeTimeEventType]struct{}),
		CodeTimeEventTypes_mapString: make(map[string]*CodeTimeEventType),

		CodeTimeReferenceTypes:           make(map[*CodeTimeReferenceType]struct{}),
		CodeTimeReferenceTypes_mapString: make(map[string]*CodeTimeReferenceType),

		CodeTrajectoryTypes:           make(map[*CodeTrajectoryType]struct{}),
		CodeTrajectoryTypes_mapString: make(map[string]*CodeTrajectoryType),

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

		TextInstructionTypes:           make(map[*TextInstructionType]struct{}),
		TextInstructionTypes_mapString: make(map[string]*TextInstructionType),

		TextNameTypes:           make(map[*TextNameType]struct{}),
		TextNameTypes_mapString: make(map[string]*TextNameType),

		TextNoteTypes:           make(map[*TextNoteType]struct{}),
		TextNoteTypes_mapString: make(map[string]*TextNoteType),

		TextPropertyNameTypes:           make(map[*TextPropertyNameType]struct{}),
		TextPropertyNameTypes_mapString: make(map[string]*TextPropertyNameType),

		TextSIDSTARDesignatorTypes:           make(map[*TextSIDSTARDesignatorType]struct{}),
		TextSIDSTARDesignatorTypes_mapString: make(map[string]*TextSIDSTARDesignatorType),

		TimeTypes:           make(map[*TimeType]struct{}),
		TimeTypes_mapString: make(map[string]*TimeType),

		ValAlarmLimitTypes:           make(map[*ValAlarmLimitType]struct{}),
		ValAlarmLimitTypes_mapString: make(map[string]*ValAlarmLimitType),

		ValAngleTypes:           make(map[*ValAngleType]struct{}),
		ValAngleTypes_mapString: make(map[string]*ValAngleType),

		ValBearingTypes:           make(map[*ValBearingType]struct{}),
		ValBearingTypes_mapString: make(map[string]*ValBearingType),

		ValChannelNumberTypes:           make(map[*ValChannelNumberType]struct{}),
		ValChannelNumberTypes_mapString: make(map[string]*ValChannelNumberType),

		ValDistanceSignedTypes:           make(map[*ValDistanceSignedType]struct{}),
		ValDistanceSignedTypes_mapString: make(map[string]*ValDistanceSignedType),

		ValDistanceTypes:           make(map[*ValDistanceType]struct{}),
		ValDistanceTypes_mapString: make(map[string]*ValDistanceType),

		ValDistanceVerticalTypes:           make(map[*ValDistanceVerticalType]struct{}),
		ValDistanceVerticalTypes_mapString: make(map[string]*ValDistanceVerticalType),

		ValDurationTypes:           make(map[*ValDurationType]struct{}),
		ValDurationTypes_mapString: make(map[string]*ValDurationType),

		ValHexTypes:           make(map[*ValHexType]struct{}),
		ValHexTypes_mapString: make(map[string]*ValHexType),

		ValSlopeTypes:           make(map[*ValSlopeType]struct{}),
		ValSlopeTypes_mapString: make(map[string]*ValSlopeType),

		ValSpeedTypes:           make(map[*ValSpeedType]struct{}),
		ValSpeedTypes_mapString: make(map[string]*ValSpeedType),

		ValTemperatureTypes:           make(map[*ValTemperatureType]struct{}),
		ValTemperatureTypes_mapString: make(map[string]*ValTemperatureType),

		ValWeightTypes:           make(map[*ValWeightType]struct{}),
		ValWeightTypes_mapString: make(map[string]*ValWeightType),

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

		CodeApproachEquipmentAdditionalType_stagedOrder: make(map[*CodeApproachEquipmentAdditionalType]uint),
		CodeApproachEquipmentAdditionalType_orderStaged: make(map[uint]*CodeApproachEquipmentAdditionalType),
		CodeApproachEquipmentAdditionalTypes_reference:  make(map[*CodeApproachEquipmentAdditionalType]*CodeApproachEquipmentAdditionalType),

		CodeApproachGuidanceType_stagedOrder: make(map[*CodeApproachGuidanceType]uint),
		CodeApproachGuidanceType_orderStaged: make(map[uint]*CodeApproachGuidanceType),
		CodeApproachGuidanceTypes_reference:  make(map[*CodeApproachGuidanceType]*CodeApproachGuidanceType),

		CodeApproachPrefixType_stagedOrder: make(map[*CodeApproachPrefixType]uint),
		CodeApproachPrefixType_orderStaged: make(map[uint]*CodeApproachPrefixType),
		CodeApproachPrefixTypes_reference:  make(map[*CodeApproachPrefixType]*CodeApproachPrefixType),

		CodeApproachType_stagedOrder: make(map[*CodeApproachType]uint),
		CodeApproachType_orderStaged: make(map[uint]*CodeApproachType),
		CodeApproachTypes_reference:  make(map[*CodeApproachType]*CodeApproachType),

		CodeApprovalType_stagedOrder: make(map[*CodeApprovalType]uint),
		CodeApprovalType_orderStaged: make(map[uint]*CodeApprovalType),
		CodeApprovalTypes_reference:  make(map[*CodeApprovalType]*CodeApprovalType),

		CodeArcDirectionType_stagedOrder: make(map[*CodeArcDirectionType]uint),
		CodeArcDirectionType_orderStaged: make(map[uint]*CodeArcDirectionType),
		CodeArcDirectionTypes_reference:  make(map[*CodeArcDirectionType]*CodeArcDirectionType),

		CodeBearingType_stagedOrder: make(map[*CodeBearingType]uint),
		CodeBearingType_orderStaged: make(map[uint]*CodeBearingType),
		CodeBearingTypes_reference:  make(map[*CodeBearingType]*CodeBearingType),

		CodeCommunicationModeType_stagedOrder: make(map[*CodeCommunicationModeType]uint),
		CodeCommunicationModeType_orderStaged: make(map[uint]*CodeCommunicationModeType),
		CodeCommunicationModeTypes_reference:  make(map[*CodeCommunicationModeType]*CodeCommunicationModeType),

		CodeCourseType_stagedOrder: make(map[*CodeCourseType]uint),
		CodeCourseType_orderStaged: make(map[uint]*CodeCourseType),
		CodeCourseTypes_reference:  make(map[*CodeCourseType]*CodeCourseType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeDesignStandardType_stagedOrder: make(map[*CodeDesignStandardType]uint),
		CodeDesignStandardType_orderStaged: make(map[uint]*CodeDesignStandardType),
		CodeDesignStandardTypes_reference:  make(map[*CodeDesignStandardType]*CodeDesignStandardType),

		CodeDesignatedPointDesignatorType_stagedOrder: make(map[*CodeDesignatedPointDesignatorType]uint),
		CodeDesignatedPointDesignatorType_orderStaged: make(map[uint]*CodeDesignatedPointDesignatorType),
		CodeDesignatedPointDesignatorTypes_reference:  make(map[*CodeDesignatedPointDesignatorType]*CodeDesignatedPointDesignatorType),

		CodeDirectionReferenceType_stagedOrder: make(map[*CodeDirectionReferenceType]uint),
		CodeDirectionReferenceType_orderStaged: make(map[uint]*CodeDirectionReferenceType),
		CodeDirectionReferenceTypes_reference:  make(map[*CodeDirectionReferenceType]*CodeDirectionReferenceType),

		CodeDirectionTurnType_stagedOrder: make(map[*CodeDirectionTurnType]uint),
		CodeDirectionTurnType_orderStaged: make(map[uint]*CodeDirectionTurnType),
		CodeDirectionTurnTypes_reference:  make(map[*CodeDirectionTurnType]*CodeDirectionTurnType),

		CodeEquipmentAntiCollisionType_stagedOrder: make(map[*CodeEquipmentAntiCollisionType]uint),
		CodeEquipmentAntiCollisionType_orderStaged: make(map[uint]*CodeEquipmentAntiCollisionType),
		CodeEquipmentAntiCollisionTypes_reference:  make(map[*CodeEquipmentAntiCollisionType]*CodeEquipmentAntiCollisionType),

		CodeEquipmentUnavailableType_stagedOrder: make(map[*CodeEquipmentUnavailableType]uint),
		CodeEquipmentUnavailableType_orderStaged: make(map[uint]*CodeEquipmentUnavailableType),
		CodeEquipmentUnavailableTypes_reference:  make(map[*CodeEquipmentUnavailableType]*CodeEquipmentUnavailableType),

		CodeFinalGuidanceType_stagedOrder: make(map[*CodeFinalGuidanceType]uint),
		CodeFinalGuidanceType_orderStaged: make(map[uint]*CodeFinalGuidanceType),
		CodeFinalGuidanceTypes_reference:  make(map[*CodeFinalGuidanceType]*CodeFinalGuidanceType),

		CodeFreeFlightType_stagedOrder: make(map[*CodeFreeFlightType]uint),
		CodeFreeFlightType_orderStaged: make(map[uint]*CodeFreeFlightType),
		CodeFreeFlightTypes_reference:  make(map[*CodeFreeFlightType]*CodeFreeFlightType),

		CodeHeightReferenceType_stagedOrder: make(map[*CodeHeightReferenceType]uint),
		CodeHeightReferenceType_orderStaged: make(map[uint]*CodeHeightReferenceType),
		CodeHeightReferenceTypes_reference:  make(map[*CodeHeightReferenceType]*CodeHeightReferenceType),

		CodeHoldingUsageType_stagedOrder: make(map[*CodeHoldingUsageType]uint),
		CodeHoldingUsageType_orderStaged: make(map[uint]*CodeHoldingUsageType),
		CodeHoldingUsageTypes_reference:  make(map[*CodeHoldingUsageType]*CodeHoldingUsageType),

		CodeHoldingUseType_stagedOrder: make(map[*CodeHoldingUseType]uint),
		CodeHoldingUseType_orderStaged: make(map[uint]*CodeHoldingUseType),
		CodeHoldingUseTypes_reference:  make(map[*CodeHoldingUseType]*CodeHoldingUseType),

		CodeICAOCountryType_stagedOrder: make(map[*CodeICAOCountryType]uint),
		CodeICAOCountryType_orderStaged: make(map[uint]*CodeICAOCountryType),
		CodeICAOCountryTypes_reference:  make(map[*CodeICAOCountryType]*CodeICAOCountryType),

		CodeMilitaryRoutePointType_stagedOrder: make(map[*CodeMilitaryRoutePointType]uint),
		CodeMilitaryRoutePointType_orderStaged: make(map[uint]*CodeMilitaryRoutePointType),
		CodeMilitaryRoutePointTypes_reference:  make(map[*CodeMilitaryRoutePointType]*CodeMilitaryRoutePointType),

		CodeMinimaFinalApproachPathType_stagedOrder: make(map[*CodeMinimaFinalApproachPathType]uint),
		CodeMinimaFinalApproachPathType_orderStaged: make(map[uint]*CodeMinimaFinalApproachPathType),
		CodeMinimaFinalApproachPathTypes_reference:  make(map[*CodeMinimaFinalApproachPathType]*CodeMinimaFinalApproachPathType),

		CodeMinimumAltitudeType_stagedOrder: make(map[*CodeMinimumAltitudeType]uint),
		CodeMinimumAltitudeType_orderStaged: make(map[uint]*CodeMinimumAltitudeType),
		CodeMinimumAltitudeTypes_reference:  make(map[*CodeMinimumAltitudeType]*CodeMinimumAltitudeType),

		CodeMinimumHeightType_stagedOrder: make(map[*CodeMinimumHeightType]uint),
		CodeMinimumHeightType_orderStaged: make(map[uint]*CodeMinimumHeightType),
		CodeMinimumHeightTypes_reference:  make(map[*CodeMinimumHeightType]*CodeMinimumHeightType),

		CodeMissedApproachType_stagedOrder: make(map[*CodeMissedApproachType]uint),
		CodeMissedApproachType_orderStaged: make(map[uint]*CodeMissedApproachType),
		CodeMissedApproachTypes_reference:  make(map[*CodeMissedApproachType]*CodeMissedApproachType),

		CodeNavigationAreaRestrictionType_stagedOrder: make(map[*CodeNavigationAreaRestrictionType]uint),
		CodeNavigationAreaRestrictionType_orderStaged: make(map[uint]*CodeNavigationAreaRestrictionType),
		CodeNavigationAreaRestrictionTypes_reference:  make(map[*CodeNavigationAreaRestrictionType]*CodeNavigationAreaRestrictionType),

		CodeNavigationAreaType_stagedOrder: make(map[*CodeNavigationAreaType]uint),
		CodeNavigationAreaType_orderStaged: make(map[uint]*CodeNavigationAreaType),
		CodeNavigationAreaTypes_reference:  make(map[*CodeNavigationAreaType]*CodeNavigationAreaType),

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

		CodeProcedureAvailabilityType_stagedOrder: make(map[*CodeProcedureAvailabilityType]uint),
		CodeProcedureAvailabilityType_orderStaged: make(map[uint]*CodeProcedureAvailabilityType),
		CodeProcedureAvailabilityTypes_reference:  make(map[*CodeProcedureAvailabilityType]*CodeProcedureAvailabilityType),

		CodeProcedureCodingStandardType_stagedOrder: make(map[*CodeProcedureCodingStandardType]uint),
		CodeProcedureCodingStandardType_orderStaged: make(map[uint]*CodeProcedureCodingStandardType),
		CodeProcedureCodingStandardTypes_reference:  make(map[*CodeProcedureCodingStandardType]*CodeProcedureCodingStandardType),

		CodeProcedureDistanceType_stagedOrder: make(map[*CodeProcedureDistanceType]uint),
		CodeProcedureDistanceType_orderStaged: make(map[uint]*CodeProcedureDistanceType),
		CodeProcedureDistanceTypes_reference:  make(map[*CodeProcedureDistanceType]*CodeProcedureDistanceType),

		CodeProcedureFixRoleType_stagedOrder: make(map[*CodeProcedureFixRoleType]uint),
		CodeProcedureFixRoleType_orderStaged: make(map[uint]*CodeProcedureFixRoleType),
		CodeProcedureFixRoleTypes_reference:  make(map[*CodeProcedureFixRoleType]*CodeProcedureFixRoleType),

		CodeProcedurePhaseType_stagedOrder: make(map[*CodeProcedurePhaseType]uint),
		CodeProcedurePhaseType_orderStaged: make(map[uint]*CodeProcedurePhaseType),
		CodeProcedurePhaseTypes_reference:  make(map[*CodeProcedurePhaseType]*CodeProcedurePhaseType),

		CodeRNPType_stagedOrder: make(map[*CodeRNPType]uint),
		CodeRNPType_orderStaged: make(map[uint]*CodeRNPType),
		CodeRNPTypes_reference:  make(map[*CodeRNPType]*CodeRNPType),

		CodeRVSMPointRoleType_stagedOrder: make(map[*CodeRVSMPointRoleType]uint),
		CodeRVSMPointRoleType_orderStaged: make(map[uint]*CodeRVSMPointRoleType),
		CodeRVSMPointRoleTypes_reference:  make(map[*CodeRVSMPointRoleType]*CodeRVSMPointRoleType),

		CodeRVSMType_stagedOrder: make(map[*CodeRVSMType]uint),
		CodeRVSMType_orderStaged: make(map[uint]*CodeRVSMType),
		CodeRVSMTypes_reference:  make(map[*CodeRVSMType]*CodeRVSMType),

		CodeReferencePathIdentifierType_stagedOrder: make(map[*CodeReferencePathIdentifierType]uint),
		CodeReferencePathIdentifierType_orderStaged: make(map[uint]*CodeReferencePathIdentifierType),
		CodeReferencePathIdentifierTypes_reference:  make(map[*CodeReferencePathIdentifierType]*CodeReferencePathIdentifierType),

		CodeReferenceRoleType_stagedOrder: make(map[*CodeReferenceRoleType]uint),
		CodeReferenceRoleType_orderStaged: make(map[uint]*CodeReferenceRoleType),
		CodeReferenceRoleTypes_reference:  make(map[*CodeReferenceRoleType]*CodeReferenceRoleType),

		CodeRelativePositionType_stagedOrder: make(map[*CodeRelativePositionType]uint),
		CodeRelativePositionType_orderStaged: make(map[uint]*CodeRelativePositionType),
		CodeRelativePositionTypes_reference:  make(map[*CodeRelativePositionType]*CodeRelativePositionType),

		CodeRouteIndicatorType_stagedOrder: make(map[*CodeRouteIndicatorType]uint),
		CodeRouteIndicatorType_orderStaged: make(map[uint]*CodeRouteIndicatorType),
		CodeRouteIndicatorTypes_reference:  make(map[*CodeRouteIndicatorType]*CodeRouteIndicatorType),

		CodeSafeAltitudeType_stagedOrder: make(map[*CodeSafeAltitudeType]uint),
		CodeSafeAltitudeType_orderStaged: make(map[uint]*CodeSafeAltitudeType),
		CodeSafeAltitudeTypes_reference:  make(map[*CodeSafeAltitudeType]*CodeSafeAltitudeType),

		CodeSegmentPathType_stagedOrder: make(map[*CodeSegmentPathType]uint),
		CodeSegmentPathType_orderStaged: make(map[uint]*CodeSegmentPathType),
		CodeSegmentPathTypes_reference:  make(map[*CodeSegmentPathType]*CodeSegmentPathType),

		CodeSegmentTerminationType_stagedOrder: make(map[*CodeSegmentTerminationType]uint),
		CodeSegmentTerminationType_orderStaged: make(map[uint]*CodeSegmentTerminationType),
		CodeSegmentTerminationTypes_reference:  make(map[*CodeSegmentTerminationType]*CodeSegmentTerminationType),

		CodeSideType_stagedOrder: make(map[*CodeSideType]uint),
		CodeSideType_orderStaged: make(map[uint]*CodeSideType),
		CodeSideTypes_reference:  make(map[*CodeSideType]*CodeSideType),

		CodeSpeedReferenceType_stagedOrder: make(map[*CodeSpeedReferenceType]uint),
		CodeSpeedReferenceType_orderStaged: make(map[uint]*CodeSpeedReferenceType),
		CodeSpeedReferenceTypes_reference:  make(map[*CodeSpeedReferenceType]*CodeSpeedReferenceType),

		CodeTAAType_stagedOrder: make(map[*CodeTAAType]uint),
		CodeTAAType_orderStaged: make(map[uint]*CodeTAAType),
		CodeTAATypes_reference:  make(map[*CodeTAAType]*CodeTAAType),

		CodeTimeEventCombinationType_stagedOrder: make(map[*CodeTimeEventCombinationType]uint),
		CodeTimeEventCombinationType_orderStaged: make(map[uint]*CodeTimeEventCombinationType),
		CodeTimeEventCombinationTypes_reference:  make(map[*CodeTimeEventCombinationType]*CodeTimeEventCombinationType),

		CodeTimeEventType_stagedOrder: make(map[*CodeTimeEventType]uint),
		CodeTimeEventType_orderStaged: make(map[uint]*CodeTimeEventType),
		CodeTimeEventTypes_reference:  make(map[*CodeTimeEventType]*CodeTimeEventType),

		CodeTimeReferenceType_stagedOrder: make(map[*CodeTimeReferenceType]uint),
		CodeTimeReferenceType_orderStaged: make(map[uint]*CodeTimeReferenceType),
		CodeTimeReferenceTypes_reference:  make(map[*CodeTimeReferenceType]*CodeTimeReferenceType),

		CodeTrajectoryType_stagedOrder: make(map[*CodeTrajectoryType]uint),
		CodeTrajectoryType_orderStaged: make(map[uint]*CodeTrajectoryType),
		CodeTrajectoryTypes_reference:  make(map[*CodeTrajectoryType]*CodeTrajectoryType),

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

		TextInstructionType_stagedOrder: make(map[*TextInstructionType]uint),
		TextInstructionType_orderStaged: make(map[uint]*TextInstructionType),
		TextInstructionTypes_reference:  make(map[*TextInstructionType]*TextInstructionType),

		TextNameType_stagedOrder: make(map[*TextNameType]uint),
		TextNameType_orderStaged: make(map[uint]*TextNameType),
		TextNameTypes_reference:  make(map[*TextNameType]*TextNameType),

		TextNoteType_stagedOrder: make(map[*TextNoteType]uint),
		TextNoteType_orderStaged: make(map[uint]*TextNoteType),
		TextNoteTypes_reference:  make(map[*TextNoteType]*TextNoteType),

		TextPropertyNameType_stagedOrder: make(map[*TextPropertyNameType]uint),
		TextPropertyNameType_orderStaged: make(map[uint]*TextPropertyNameType),
		TextPropertyNameTypes_reference:  make(map[*TextPropertyNameType]*TextPropertyNameType),

		TextSIDSTARDesignatorType_stagedOrder: make(map[*TextSIDSTARDesignatorType]uint),
		TextSIDSTARDesignatorType_orderStaged: make(map[uint]*TextSIDSTARDesignatorType),
		TextSIDSTARDesignatorTypes_reference:  make(map[*TextSIDSTARDesignatorType]*TextSIDSTARDesignatorType),

		TimeType_stagedOrder: make(map[*TimeType]uint),
		TimeType_orderStaged: make(map[uint]*TimeType),
		TimeTypes_reference:  make(map[*TimeType]*TimeType),

		ValAlarmLimitType_stagedOrder: make(map[*ValAlarmLimitType]uint),
		ValAlarmLimitType_orderStaged: make(map[uint]*ValAlarmLimitType),
		ValAlarmLimitTypes_reference:  make(map[*ValAlarmLimitType]*ValAlarmLimitType),

		ValAngleType_stagedOrder: make(map[*ValAngleType]uint),
		ValAngleType_orderStaged: make(map[uint]*ValAngleType),
		ValAngleTypes_reference:  make(map[*ValAngleType]*ValAngleType),

		ValBearingType_stagedOrder: make(map[*ValBearingType]uint),
		ValBearingType_orderStaged: make(map[uint]*ValBearingType),
		ValBearingTypes_reference:  make(map[*ValBearingType]*ValBearingType),

		ValChannelNumberType_stagedOrder: make(map[*ValChannelNumberType]uint),
		ValChannelNumberType_orderStaged: make(map[uint]*ValChannelNumberType),
		ValChannelNumberTypes_reference:  make(map[*ValChannelNumberType]*ValChannelNumberType),

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

		ValHexType_stagedOrder: make(map[*ValHexType]uint),
		ValHexType_orderStaged: make(map[uint]*ValHexType),
		ValHexTypes_reference:  make(map[*ValHexType]*ValHexType),

		ValSlopeType_stagedOrder: make(map[*ValSlopeType]uint),
		ValSlopeType_orderStaged: make(map[uint]*ValSlopeType),
		ValSlopeTypes_reference:  make(map[*ValSlopeType]*ValSlopeType),

		ValSpeedType_stagedOrder: make(map[*ValSpeedType]uint),
		ValSpeedType_orderStaged: make(map[uint]*ValSpeedType),
		ValSpeedTypes_reference:  make(map[*ValSpeedType]*ValSpeedType),

		ValTemperatureType_stagedOrder: make(map[*ValTemperatureType]uint),
		ValTemperatureType_orderStaged: make(map[uint]*ValTemperatureType),
		ValTemperatureTypes_reference:  make(map[*ValTemperatureType]*ValTemperatureType),

		ValWeightType_stagedOrder: make(map[*ValWeightType]uint),
		ValWeightType_orderStaged: make(map[uint]*ValWeightType),
		ValWeightTypes_reference:  make(map[*ValWeightType]*ValWeightType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeATCReportingType": &CodeATCReportingTypeUnmarshaller{},

			"CodeAerialRefuellingPointType": &CodeAerialRefuellingPointTypeUnmarshaller{},

			"CodeAircraftCategoryType": &CodeAircraftCategoryTypeUnmarshaller{},

			"CodeAircraftEngineNumberType": &CodeAircraftEngineNumberTypeUnmarshaller{},

			"CodeAircraftEngineType": &CodeAircraftEngineTypeUnmarshaller{},

			"CodeAircraftICAOType": &CodeAircraftICAOTypeUnmarshaller{},

			"CodeAircraftType": &CodeAircraftTypeUnmarshaller{},

			"CodeAircraftWingspanClassType": &CodeAircraftWingspanClassTypeUnmarshaller{},

			"CodeAltitudeAdjustmentType": &CodeAltitudeAdjustmentTypeUnmarshaller{},

			"CodeAltitudeUseType": &CodeAltitudeUseTypeUnmarshaller{},

			"CodeApproachEquipmentAdditionalType": &CodeApproachEquipmentAdditionalTypeUnmarshaller{},

			"CodeApproachGuidanceType": &CodeApproachGuidanceTypeUnmarshaller{},

			"CodeApproachPrefixType": &CodeApproachPrefixTypeUnmarshaller{},

			"CodeApproachType": &CodeApproachTypeUnmarshaller{},

			"CodeApprovalType": &CodeApprovalTypeUnmarshaller{},

			"CodeArcDirectionType": &CodeArcDirectionTypeUnmarshaller{},

			"CodeBearingType": &CodeBearingTypeUnmarshaller{},

			"CodeCommunicationModeType": &CodeCommunicationModeTypeUnmarshaller{},

			"CodeCourseType": &CodeCourseTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeDesignStandardType": &CodeDesignStandardTypeUnmarshaller{},

			"CodeDesignatedPointDesignatorType": &CodeDesignatedPointDesignatorTypeUnmarshaller{},

			"CodeDirectionReferenceType": &CodeDirectionReferenceTypeUnmarshaller{},

			"CodeDirectionTurnType": &CodeDirectionTurnTypeUnmarshaller{},

			"CodeEquipmentAntiCollisionType": &CodeEquipmentAntiCollisionTypeUnmarshaller{},

			"CodeEquipmentUnavailableType": &CodeEquipmentUnavailableTypeUnmarshaller{},

			"CodeFinalGuidanceType": &CodeFinalGuidanceTypeUnmarshaller{},

			"CodeFreeFlightType": &CodeFreeFlightTypeUnmarshaller{},

			"CodeHeightReferenceType": &CodeHeightReferenceTypeUnmarshaller{},

			"CodeHoldingUsageType": &CodeHoldingUsageTypeUnmarshaller{},

			"CodeHoldingUseType": &CodeHoldingUseTypeUnmarshaller{},

			"CodeICAOCountryType": &CodeICAOCountryTypeUnmarshaller{},

			"CodeMilitaryRoutePointType": &CodeMilitaryRoutePointTypeUnmarshaller{},

			"CodeMinimaFinalApproachPathType": &CodeMinimaFinalApproachPathTypeUnmarshaller{},

			"CodeMinimumAltitudeType": &CodeMinimumAltitudeTypeUnmarshaller{},

			"CodeMinimumHeightType": &CodeMinimumHeightTypeUnmarshaller{},

			"CodeMissedApproachType": &CodeMissedApproachTypeUnmarshaller{},

			"CodeNavigationAreaRestrictionType": &CodeNavigationAreaRestrictionTypeUnmarshaller{},

			"CodeNavigationAreaType": &CodeNavigationAreaTypeUnmarshaller{},

			"CodeNavigationEquipmentType": &CodeNavigationEquipmentTypeUnmarshaller{},

			"CodeNavigationSpecificationType": &CodeNavigationSpecificationTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeObstacleAssessmentSurfaceType": &CodeObstacleAssessmentSurfaceTypeUnmarshaller{},

			"CodeObstructionIdSurfaceZoneType": &CodeObstructionIdSurfaceZoneTypeUnmarshaller{},

			"CodeProcedureAvailabilityType": &CodeProcedureAvailabilityTypeUnmarshaller{},

			"CodeProcedureCodingStandardType": &CodeProcedureCodingStandardTypeUnmarshaller{},

			"CodeProcedureDistanceType": &CodeProcedureDistanceTypeUnmarshaller{},

			"CodeProcedureFixRoleType": &CodeProcedureFixRoleTypeUnmarshaller{},

			"CodeProcedurePhaseType": &CodeProcedurePhaseTypeUnmarshaller{},

			"CodeRNPType": &CodeRNPTypeUnmarshaller{},

			"CodeRVSMPointRoleType": &CodeRVSMPointRoleTypeUnmarshaller{},

			"CodeRVSMType": &CodeRVSMTypeUnmarshaller{},

			"CodeReferencePathIdentifierType": &CodeReferencePathIdentifierTypeUnmarshaller{},

			"CodeReferenceRoleType": &CodeReferenceRoleTypeUnmarshaller{},

			"CodeRelativePositionType": &CodeRelativePositionTypeUnmarshaller{},

			"CodeRouteIndicatorType": &CodeRouteIndicatorTypeUnmarshaller{},

			"CodeSafeAltitudeType": &CodeSafeAltitudeTypeUnmarshaller{},

			"CodeSegmentPathType": &CodeSegmentPathTypeUnmarshaller{},

			"CodeSegmentTerminationType": &CodeSegmentTerminationTypeUnmarshaller{},

			"CodeSideType": &CodeSideTypeUnmarshaller{},

			"CodeSpeedReferenceType": &CodeSpeedReferenceTypeUnmarshaller{},

			"CodeTAAType": &CodeTAATypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeTrajectoryType": &CodeTrajectoryTypeUnmarshaller{},

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

			"TextInstructionType": &TextInstructionTypeUnmarshaller{},

			"TextNameType": &TextNameTypeUnmarshaller{},

			"TextNoteType": &TextNoteTypeUnmarshaller{},

			"TextPropertyNameType": &TextPropertyNameTypeUnmarshaller{},

			"TextSIDSTARDesignatorType": &TextSIDSTARDesignatorTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"ValAlarmLimitType": &ValAlarmLimitTypeUnmarshaller{},

			"ValAngleType": &ValAngleTypeUnmarshaller{},

			"ValBearingType": &ValBearingTypeUnmarshaller{},

			"ValChannelNumberType": &ValChannelNumberTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValHexType": &ValHexTypeUnmarshaller{},

			"ValSlopeType": &ValSlopeTypeUnmarshaller{},

			"ValSpeedType": &ValSpeedTypeUnmarshaller{},

			"ValTemperatureType": &ValTemperatureTypeUnmarshaller{},

			"ValWeightType": &ValWeightTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeAircraftCategoryType"] = len(stage.CodeAircraftCategoryTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineNumberType"] = len(stage.CodeAircraftEngineNumberTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineType"] = len(stage.CodeAircraftEngineTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftICAOType"] = len(stage.CodeAircraftICAOTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftType"] = len(stage.CodeAircraftTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftWingspanClassType"] = len(stage.CodeAircraftWingspanClassTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeAdjustmentType"] = len(stage.CodeAltitudeAdjustmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeUseType"] = len(stage.CodeAltitudeUseTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachEquipmentAdditionalType"] = len(stage.CodeApproachEquipmentAdditionalTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachGuidanceType"] = len(stage.CodeApproachGuidanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachPrefixType"] = len(stage.CodeApproachPrefixTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachType"] = len(stage.CodeApproachTypes)
	stage.Map_GongStructName_InstancesNb["CodeApprovalType"] = len(stage.CodeApprovalTypes)
	stage.Map_GongStructName_InstancesNb["CodeArcDirectionType"] = len(stage.CodeArcDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeBearingType"] = len(stage.CodeBearingTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationModeType"] = len(stage.CodeCommunicationModeTypes)
	stage.Map_GongStructName_InstancesNb["CodeCourseType"] = len(stage.CodeCourseTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeDesignStandardType"] = len(stage.CodeDesignStandardTypes)
	stage.Map_GongStructName_InstancesNb["CodeDesignatedPointDesignatorType"] = len(stage.CodeDesignatedPointDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionReferenceType"] = len(stage.CodeDirectionReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionTurnType"] = len(stage.CodeDirectionTurnTypes)
	stage.Map_GongStructName_InstancesNb["CodeEquipmentAntiCollisionType"] = len(stage.CodeEquipmentAntiCollisionTypes)
	stage.Map_GongStructName_InstancesNb["CodeEquipmentUnavailableType"] = len(stage.CodeEquipmentUnavailableTypes)
	stage.Map_GongStructName_InstancesNb["CodeFinalGuidanceType"] = len(stage.CodeFinalGuidanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeFreeFlightType"] = len(stage.CodeFreeFlightTypes)
	stage.Map_GongStructName_InstancesNb["CodeHeightReferenceType"] = len(stage.CodeHeightReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeHoldingUsageType"] = len(stage.CodeHoldingUsageTypes)
	stage.Map_GongStructName_InstancesNb["CodeHoldingUseType"] = len(stage.CodeHoldingUseTypes)
	stage.Map_GongStructName_InstancesNb["CodeICAOCountryType"] = len(stage.CodeICAOCountryTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryRoutePointType"] = len(stage.CodeMilitaryRoutePointTypes)
	stage.Map_GongStructName_InstancesNb["CodeMinimaFinalApproachPathType"] = len(stage.CodeMinimaFinalApproachPathTypes)
	stage.Map_GongStructName_InstancesNb["CodeMinimumAltitudeType"] = len(stage.CodeMinimumAltitudeTypes)
	stage.Map_GongStructName_InstancesNb["CodeMinimumHeightType"] = len(stage.CodeMinimumHeightTypes)
	stage.Map_GongStructName_InstancesNb["CodeMissedApproachType"] = len(stage.CodeMissedApproachTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationAreaRestrictionType"] = len(stage.CodeNavigationAreaRestrictionTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationAreaType"] = len(stage.CodeNavigationAreaTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationEquipmentType"] = len(stage.CodeNavigationEquipmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationSpecificationType"] = len(stage.CodeNavigationSpecificationTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstacleAssessmentSurfaceType"] = len(stage.CodeObstacleAssessmentSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstructionIdSurfaceZoneType"] = len(stage.CodeObstructionIdSurfaceZoneTypes)
	stage.Map_GongStructName_InstancesNb["CodeProcedureAvailabilityType"] = len(stage.CodeProcedureAvailabilityTypes)
	stage.Map_GongStructName_InstancesNb["CodeProcedureCodingStandardType"] = len(stage.CodeProcedureCodingStandardTypes)
	stage.Map_GongStructName_InstancesNb["CodeProcedureDistanceType"] = len(stage.CodeProcedureDistanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeProcedureFixRoleType"] = len(stage.CodeProcedureFixRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeProcedurePhaseType"] = len(stage.CodeProcedurePhaseTypes)
	stage.Map_GongStructName_InstancesNb["CodeRNPType"] = len(stage.CodeRNPTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMPointRoleType"] = len(stage.CodeRVSMPointRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMType"] = len(stage.CodeRVSMTypes)
	stage.Map_GongStructName_InstancesNb["CodeReferencePathIdentifierType"] = len(stage.CodeReferencePathIdentifierTypes)
	stage.Map_GongStructName_InstancesNb["CodeReferenceRoleType"] = len(stage.CodeReferenceRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRelativePositionType"] = len(stage.CodeRelativePositionTypes)
	stage.Map_GongStructName_InstancesNb["CodeRouteIndicatorType"] = len(stage.CodeRouteIndicatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeSafeAltitudeType"] = len(stage.CodeSafeAltitudeTypes)
	stage.Map_GongStructName_InstancesNb["CodeSegmentPathType"] = len(stage.CodeSegmentPathTypes)
	stage.Map_GongStructName_InstancesNb["CodeSegmentTerminationType"] = len(stage.CodeSegmentTerminationTypes)
	stage.Map_GongStructName_InstancesNb["CodeSideType"] = len(stage.CodeSideTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpeedReferenceType"] = len(stage.CodeSpeedReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTAAType"] = len(stage.CodeTAATypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTrajectoryType"] = len(stage.CodeTrajectoryTypes)
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
	stage.Map_GongStructName_InstancesNb["TextInstructionType"] = len(stage.TextInstructionTypes)
	stage.Map_GongStructName_InstancesNb["TextNameType"] = len(stage.TextNameTypes)
	stage.Map_GongStructName_InstancesNb["TextNoteType"] = len(stage.TextNoteTypes)
	stage.Map_GongStructName_InstancesNb["TextPropertyNameType"] = len(stage.TextPropertyNameTypes)
	stage.Map_GongStructName_InstancesNb["TextSIDSTARDesignatorType"] = len(stage.TextSIDSTARDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["ValAlarmLimitType"] = len(stage.ValAlarmLimitTypes)
	stage.Map_GongStructName_InstancesNb["ValAngleType"] = len(stage.ValAngleTypes)
	stage.Map_GongStructName_InstancesNb["ValBearingType"] = len(stage.ValBearingTypes)
	stage.Map_GongStructName_InstancesNb["ValChannelNumberType"] = len(stage.ValChannelNumberTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValHexType"] = len(stage.ValHexTypes)
	stage.Map_GongStructName_InstancesNb["ValSlopeType"] = len(stage.ValSlopeTypes)
	stage.Map_GongStructName_InstancesNb["ValSpeedType"] = len(stage.ValSpeedTypes)
	stage.Map_GongStructName_InstancesNb["ValTemperatureType"] = len(stage.ValTemperatureTypes)
	stage.Map_GongStructName_InstancesNb["ValWeightType"] = len(stage.ValWeightTypes)
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

// Stage puts codeapproachequipmentadditionaltype to the model stage
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) Stage(stage *Stage) *CodeApproachEquipmentAdditionalType {
	__gong__stage(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalType_stagedOrder, stage.CodeApproachEquipmentAdditionalType_orderStaged, &stage.CodeApproachEquipmentAdditionalTypeOrder, stage.CodeApproachEquipmentAdditionalTypes_mapString, codeapproachequipmentadditionaltype, codeapproachequipmentadditionaltype.Name)
	return codeapproachequipmentadditionaltype
}

// StagePreserveOrder puts codeapproachequipmentadditionaltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachEquipmentAdditionalTypeOrder
// - update stage.CodeApproachEquipmentAdditionalTypeOrder accordingly
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalType_stagedOrder, stage.CodeApproachEquipmentAdditionalType_orderStaged, &stage.CodeApproachEquipmentAdditionalTypeOrder, stage.CodeApproachEquipmentAdditionalTypes_mapString, codeapproachequipmentadditionaltype, order, codeapproachequipmentadditionaltype.Name)
}

// Unstage removes codeapproachequipmentadditionaltype off the model stage
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) Unstage(stage *Stage) *CodeApproachEquipmentAdditionalType {
	__gong__unstage(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalTypes_mapString, codeapproachequipmentadditionaltype, codeapproachequipmentadditionaltype.Name)
	return codeapproachequipmentadditionaltype
}

// UnstageVoid removes codeapproachequipmentadditionaltype off the model stage
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) UnstageVoid(stage *Stage) {
	codeapproachequipmentadditionaltype.Unstage(stage)
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) StageVoid(stage *Stage) {
	codeapproachequipmentadditionaltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GetName() (res string) {
	return codeapproachequipmentadditionaltype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) SetName(name string) {
	codeapproachequipmentadditionaltype.Name = name
}

// Stage puts codeapproachguidancetype to the model stage
func (codeapproachguidancetype *CodeApproachGuidanceType) Stage(stage *Stage) *CodeApproachGuidanceType {
	__gong__stage(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceType_stagedOrder, stage.CodeApproachGuidanceType_orderStaged, &stage.CodeApproachGuidanceTypeOrder, stage.CodeApproachGuidanceTypes_mapString, codeapproachguidancetype, codeapproachguidancetype.Name)
	return codeapproachguidancetype
}

// StagePreserveOrder puts codeapproachguidancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachGuidanceTypeOrder
// - update stage.CodeApproachGuidanceTypeOrder accordingly
func (codeapproachguidancetype *CodeApproachGuidanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceType_stagedOrder, stage.CodeApproachGuidanceType_orderStaged, &stage.CodeApproachGuidanceTypeOrder, stage.CodeApproachGuidanceTypes_mapString, codeapproachguidancetype, order, codeapproachguidancetype.Name)
}

// Unstage removes codeapproachguidancetype off the model stage
func (codeapproachguidancetype *CodeApproachGuidanceType) Unstage(stage *Stage) *CodeApproachGuidanceType {
	__gong__unstage(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceTypes_mapString, codeapproachguidancetype, codeapproachguidancetype.Name)
	return codeapproachguidancetype
}

// UnstageVoid removes codeapproachguidancetype off the model stage
func (codeapproachguidancetype *CodeApproachGuidanceType) UnstageVoid(stage *Stage) {
	codeapproachguidancetype.Unstage(stage)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) StageVoid(stage *Stage) {
	codeapproachguidancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachguidancetype *CodeApproachGuidanceType) GetName() (res string) {
	return codeapproachguidancetype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachguidancetype *CodeApproachGuidanceType) SetName(name string) {
	codeapproachguidancetype.Name = name
}

// Stage puts codeapproachprefixtype to the model stage
func (codeapproachprefixtype *CodeApproachPrefixType) Stage(stage *Stage) *CodeApproachPrefixType {
	__gong__stage(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixType_stagedOrder, stage.CodeApproachPrefixType_orderStaged, &stage.CodeApproachPrefixTypeOrder, stage.CodeApproachPrefixTypes_mapString, codeapproachprefixtype, codeapproachprefixtype.Name)
	return codeapproachprefixtype
}

// StagePreserveOrder puts codeapproachprefixtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachPrefixTypeOrder
// - update stage.CodeApproachPrefixTypeOrder accordingly
func (codeapproachprefixtype *CodeApproachPrefixType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixType_stagedOrder, stage.CodeApproachPrefixType_orderStaged, &stage.CodeApproachPrefixTypeOrder, stage.CodeApproachPrefixTypes_mapString, codeapproachprefixtype, order, codeapproachprefixtype.Name)
}

// Unstage removes codeapproachprefixtype off the model stage
func (codeapproachprefixtype *CodeApproachPrefixType) Unstage(stage *Stage) *CodeApproachPrefixType {
	__gong__unstage(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixTypes_mapString, codeapproachprefixtype, codeapproachprefixtype.Name)
	return codeapproachprefixtype
}

// UnstageVoid removes codeapproachprefixtype off the model stage
func (codeapproachprefixtype *CodeApproachPrefixType) UnstageVoid(stage *Stage) {
	codeapproachprefixtype.Unstage(stage)
}

func (codeapproachprefixtype *CodeApproachPrefixType) StageVoid(stage *Stage) {
	codeapproachprefixtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachprefixtype *CodeApproachPrefixType) GetName() (res string) {
	return codeapproachprefixtype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachprefixtype *CodeApproachPrefixType) SetName(name string) {
	codeapproachprefixtype.Name = name
}

// Stage puts codeapproachtype to the model stage
func (codeapproachtype *CodeApproachType) Stage(stage *Stage) *CodeApproachType {
	__gong__stage(stage.CodeApproachTypes, stage.CodeApproachType_stagedOrder, stage.CodeApproachType_orderStaged, &stage.CodeApproachTypeOrder, stage.CodeApproachTypes_mapString, codeapproachtype, codeapproachtype.Name)
	return codeapproachtype
}

// StagePreserveOrder puts codeapproachtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachTypeOrder
// - update stage.CodeApproachTypeOrder accordingly
func (codeapproachtype *CodeApproachType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachTypes, stage.CodeApproachType_stagedOrder, stage.CodeApproachType_orderStaged, &stage.CodeApproachTypeOrder, stage.CodeApproachTypes_mapString, codeapproachtype, order, codeapproachtype.Name)
}

// Unstage removes codeapproachtype off the model stage
func (codeapproachtype *CodeApproachType) Unstage(stage *Stage) *CodeApproachType {
	__gong__unstage(stage.CodeApproachTypes, stage.CodeApproachTypes_mapString, codeapproachtype, codeapproachtype.Name)
	return codeapproachtype
}

// UnstageVoid removes codeapproachtype off the model stage
func (codeapproachtype *CodeApproachType) UnstageVoid(stage *Stage) {
	codeapproachtype.Unstage(stage)
}

func (codeapproachtype *CodeApproachType) StageVoid(stage *Stage) {
	codeapproachtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachtype *CodeApproachType) GetName() (res string) {
	return codeapproachtype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachtype *CodeApproachType) SetName(name string) {
	codeapproachtype.Name = name
}

// Stage puts codeapprovaltype to the model stage
func (codeapprovaltype *CodeApprovalType) Stage(stage *Stage) *CodeApprovalType {
	__gong__stage(stage.CodeApprovalTypes, stage.CodeApprovalType_stagedOrder, stage.CodeApprovalType_orderStaged, &stage.CodeApprovalTypeOrder, stage.CodeApprovalTypes_mapString, codeapprovaltype, codeapprovaltype.Name)
	return codeapprovaltype
}

// StagePreserveOrder puts codeapprovaltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApprovalTypeOrder
// - update stage.CodeApprovalTypeOrder accordingly
func (codeapprovaltype *CodeApprovalType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApprovalTypes, stage.CodeApprovalType_stagedOrder, stage.CodeApprovalType_orderStaged, &stage.CodeApprovalTypeOrder, stage.CodeApprovalTypes_mapString, codeapprovaltype, order, codeapprovaltype.Name)
}

// Unstage removes codeapprovaltype off the model stage
func (codeapprovaltype *CodeApprovalType) Unstage(stage *Stage) *CodeApprovalType {
	__gong__unstage(stage.CodeApprovalTypes, stage.CodeApprovalTypes_mapString, codeapprovaltype, codeapprovaltype.Name)
	return codeapprovaltype
}

// UnstageVoid removes codeapprovaltype off the model stage
func (codeapprovaltype *CodeApprovalType) UnstageVoid(stage *Stage) {
	codeapprovaltype.Unstage(stage)
}

func (codeapprovaltype *CodeApprovalType) StageVoid(stage *Stage) {
	codeapprovaltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapprovaltype *CodeApprovalType) GetName() (res string) {
	return codeapprovaltype.Name
}

// for satisfaction of GongStruct interface
func (codeapprovaltype *CodeApprovalType) SetName(name string) {
	codeapprovaltype.Name = name
}

// Stage puts codearcdirectiontype to the model stage
func (codearcdirectiontype *CodeArcDirectionType) Stage(stage *Stage) *CodeArcDirectionType {
	__gong__stage(stage.CodeArcDirectionTypes, stage.CodeArcDirectionType_stagedOrder, stage.CodeArcDirectionType_orderStaged, &stage.CodeArcDirectionTypeOrder, stage.CodeArcDirectionTypes_mapString, codearcdirectiontype, codearcdirectiontype.Name)
	return codearcdirectiontype
}

// StagePreserveOrder puts codearcdirectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeArcDirectionTypeOrder
// - update stage.CodeArcDirectionTypeOrder accordingly
func (codearcdirectiontype *CodeArcDirectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeArcDirectionTypes, stage.CodeArcDirectionType_stagedOrder, stage.CodeArcDirectionType_orderStaged, &stage.CodeArcDirectionTypeOrder, stage.CodeArcDirectionTypes_mapString, codearcdirectiontype, order, codearcdirectiontype.Name)
}

// Unstage removes codearcdirectiontype off the model stage
func (codearcdirectiontype *CodeArcDirectionType) Unstage(stage *Stage) *CodeArcDirectionType {
	__gong__unstage(stage.CodeArcDirectionTypes, stage.CodeArcDirectionTypes_mapString, codearcdirectiontype, codearcdirectiontype.Name)
	return codearcdirectiontype
}

// UnstageVoid removes codearcdirectiontype off the model stage
func (codearcdirectiontype *CodeArcDirectionType) UnstageVoid(stage *Stage) {
	codearcdirectiontype.Unstage(stage)
}

func (codearcdirectiontype *CodeArcDirectionType) StageVoid(stage *Stage) {
	codearcdirectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codearcdirectiontype *CodeArcDirectionType) GetName() (res string) {
	return codearcdirectiontype.Name
}

// for satisfaction of GongStruct interface
func (codearcdirectiontype *CodeArcDirectionType) SetName(name string) {
	codearcdirectiontype.Name = name
}

// Stage puts codebearingtype to the model stage
func (codebearingtype *CodeBearingType) Stage(stage *Stage) *CodeBearingType {
	__gong__stage(stage.CodeBearingTypes, stage.CodeBearingType_stagedOrder, stage.CodeBearingType_orderStaged, &stage.CodeBearingTypeOrder, stage.CodeBearingTypes_mapString, codebearingtype, codebearingtype.Name)
	return codebearingtype
}

// StagePreserveOrder puts codebearingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeBearingTypeOrder
// - update stage.CodeBearingTypeOrder accordingly
func (codebearingtype *CodeBearingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeBearingTypes, stage.CodeBearingType_stagedOrder, stage.CodeBearingType_orderStaged, &stage.CodeBearingTypeOrder, stage.CodeBearingTypes_mapString, codebearingtype, order, codebearingtype.Name)
}

// Unstage removes codebearingtype off the model stage
func (codebearingtype *CodeBearingType) Unstage(stage *Stage) *CodeBearingType {
	__gong__unstage(stage.CodeBearingTypes, stage.CodeBearingTypes_mapString, codebearingtype, codebearingtype.Name)
	return codebearingtype
}

// UnstageVoid removes codebearingtype off the model stage
func (codebearingtype *CodeBearingType) UnstageVoid(stage *Stage) {
	codebearingtype.Unstage(stage)
}

func (codebearingtype *CodeBearingType) StageVoid(stage *Stage) {
	codebearingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codebearingtype *CodeBearingType) GetName() (res string) {
	return codebearingtype.Name
}

// for satisfaction of GongStruct interface
func (codebearingtype *CodeBearingType) SetName(name string) {
	codebearingtype.Name = name
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

// Stage puts codedesignstandardtype to the model stage
func (codedesignstandardtype *CodeDesignStandardType) Stage(stage *Stage) *CodeDesignStandardType {
	__gong__stage(stage.CodeDesignStandardTypes, stage.CodeDesignStandardType_stagedOrder, stage.CodeDesignStandardType_orderStaged, &stage.CodeDesignStandardTypeOrder, stage.CodeDesignStandardTypes_mapString, codedesignstandardtype, codedesignstandardtype.Name)
	return codedesignstandardtype
}

// StagePreserveOrder puts codedesignstandardtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDesignStandardTypeOrder
// - update stage.CodeDesignStandardTypeOrder accordingly
func (codedesignstandardtype *CodeDesignStandardType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDesignStandardTypes, stage.CodeDesignStandardType_stagedOrder, stage.CodeDesignStandardType_orderStaged, &stage.CodeDesignStandardTypeOrder, stage.CodeDesignStandardTypes_mapString, codedesignstandardtype, order, codedesignstandardtype.Name)
}

// Unstage removes codedesignstandardtype off the model stage
func (codedesignstandardtype *CodeDesignStandardType) Unstage(stage *Stage) *CodeDesignStandardType {
	__gong__unstage(stage.CodeDesignStandardTypes, stage.CodeDesignStandardTypes_mapString, codedesignstandardtype, codedesignstandardtype.Name)
	return codedesignstandardtype
}

// UnstageVoid removes codedesignstandardtype off the model stage
func (codedesignstandardtype *CodeDesignStandardType) UnstageVoid(stage *Stage) {
	codedesignstandardtype.Unstage(stage)
}

func (codedesignstandardtype *CodeDesignStandardType) StageVoid(stage *Stage) {
	codedesignstandardtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedesignstandardtype *CodeDesignStandardType) GetName() (res string) {
	return codedesignstandardtype.Name
}

// for satisfaction of GongStruct interface
func (codedesignstandardtype *CodeDesignStandardType) SetName(name string) {
	codedesignstandardtype.Name = name
}

// Stage puts codedesignatedpointdesignatortype to the model stage
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) Stage(stage *Stage) *CodeDesignatedPointDesignatorType {
	__gong__stage(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorType_stagedOrder, stage.CodeDesignatedPointDesignatorType_orderStaged, &stage.CodeDesignatedPointDesignatorTypeOrder, stage.CodeDesignatedPointDesignatorTypes_mapString, codedesignatedpointdesignatortype, codedesignatedpointdesignatortype.Name)
	return codedesignatedpointdesignatortype
}

// StagePreserveOrder puts codedesignatedpointdesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDesignatedPointDesignatorTypeOrder
// - update stage.CodeDesignatedPointDesignatorTypeOrder accordingly
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorType_stagedOrder, stage.CodeDesignatedPointDesignatorType_orderStaged, &stage.CodeDesignatedPointDesignatorTypeOrder, stage.CodeDesignatedPointDesignatorTypes_mapString, codedesignatedpointdesignatortype, order, codedesignatedpointdesignatortype.Name)
}

// Unstage removes codedesignatedpointdesignatortype off the model stage
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) Unstage(stage *Stage) *CodeDesignatedPointDesignatorType {
	__gong__unstage(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorTypes_mapString, codedesignatedpointdesignatortype, codedesignatedpointdesignatortype.Name)
	return codedesignatedpointdesignatortype
}

// UnstageVoid removes codedesignatedpointdesignatortype off the model stage
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) UnstageVoid(stage *Stage) {
	codedesignatedpointdesignatortype.Unstage(stage)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) StageVoid(stage *Stage) {
	codedesignatedpointdesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GetName() (res string) {
	return codedesignatedpointdesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) SetName(name string) {
	codedesignatedpointdesignatortype.Name = name
}

// Stage puts codedirectionreferencetype to the model stage
func (codedirectionreferencetype *CodeDirectionReferenceType) Stage(stage *Stage) *CodeDirectionReferenceType {
	__gong__stage(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceType_stagedOrder, stage.CodeDirectionReferenceType_orderStaged, &stage.CodeDirectionReferenceTypeOrder, stage.CodeDirectionReferenceTypes_mapString, codedirectionreferencetype, codedirectionreferencetype.Name)
	return codedirectionreferencetype
}

// StagePreserveOrder puts codedirectionreferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDirectionReferenceTypeOrder
// - update stage.CodeDirectionReferenceTypeOrder accordingly
func (codedirectionreferencetype *CodeDirectionReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceType_stagedOrder, stage.CodeDirectionReferenceType_orderStaged, &stage.CodeDirectionReferenceTypeOrder, stage.CodeDirectionReferenceTypes_mapString, codedirectionreferencetype, order, codedirectionreferencetype.Name)
}

// Unstage removes codedirectionreferencetype off the model stage
func (codedirectionreferencetype *CodeDirectionReferenceType) Unstage(stage *Stage) *CodeDirectionReferenceType {
	__gong__unstage(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceTypes_mapString, codedirectionreferencetype, codedirectionreferencetype.Name)
	return codedirectionreferencetype
}

// UnstageVoid removes codedirectionreferencetype off the model stage
func (codedirectionreferencetype *CodeDirectionReferenceType) UnstageVoid(stage *Stage) {
	codedirectionreferencetype.Unstage(stage)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) StageVoid(stage *Stage) {
	codedirectionreferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedirectionreferencetype *CodeDirectionReferenceType) GetName() (res string) {
	return codedirectionreferencetype.Name
}

// for satisfaction of GongStruct interface
func (codedirectionreferencetype *CodeDirectionReferenceType) SetName(name string) {
	codedirectionreferencetype.Name = name
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

// Stage puts codeequipmentunavailabletype to the model stage
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) Stage(stage *Stage) *CodeEquipmentUnavailableType {
	__gong__stage(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableType_stagedOrder, stage.CodeEquipmentUnavailableType_orderStaged, &stage.CodeEquipmentUnavailableTypeOrder, stage.CodeEquipmentUnavailableTypes_mapString, codeequipmentunavailabletype, codeequipmentunavailabletype.Name)
	return codeequipmentunavailabletype
}

// StagePreserveOrder puts codeequipmentunavailabletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeEquipmentUnavailableTypeOrder
// - update stage.CodeEquipmentUnavailableTypeOrder accordingly
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableType_stagedOrder, stage.CodeEquipmentUnavailableType_orderStaged, &stage.CodeEquipmentUnavailableTypeOrder, stage.CodeEquipmentUnavailableTypes_mapString, codeequipmentunavailabletype, order, codeequipmentunavailabletype.Name)
}

// Unstage removes codeequipmentunavailabletype off the model stage
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) Unstage(stage *Stage) *CodeEquipmentUnavailableType {
	__gong__unstage(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableTypes_mapString, codeequipmentunavailabletype, codeequipmentunavailabletype.Name)
	return codeequipmentunavailabletype
}

// UnstageVoid removes codeequipmentunavailabletype off the model stage
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) UnstageVoid(stage *Stage) {
	codeequipmentunavailabletype.Unstage(stage)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) StageVoid(stage *Stage) {
	codeequipmentunavailabletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GetName() (res string) {
	return codeequipmentunavailabletype.Name
}

// for satisfaction of GongStruct interface
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) SetName(name string) {
	codeequipmentunavailabletype.Name = name
}

// Stage puts codefinalguidancetype to the model stage
func (codefinalguidancetype *CodeFinalGuidanceType) Stage(stage *Stage) *CodeFinalGuidanceType {
	__gong__stage(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceType_stagedOrder, stage.CodeFinalGuidanceType_orderStaged, &stage.CodeFinalGuidanceTypeOrder, stage.CodeFinalGuidanceTypes_mapString, codefinalguidancetype, codefinalguidancetype.Name)
	return codefinalguidancetype
}

// StagePreserveOrder puts codefinalguidancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFinalGuidanceTypeOrder
// - update stage.CodeFinalGuidanceTypeOrder accordingly
func (codefinalguidancetype *CodeFinalGuidanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceType_stagedOrder, stage.CodeFinalGuidanceType_orderStaged, &stage.CodeFinalGuidanceTypeOrder, stage.CodeFinalGuidanceTypes_mapString, codefinalguidancetype, order, codefinalguidancetype.Name)
}

// Unstage removes codefinalguidancetype off the model stage
func (codefinalguidancetype *CodeFinalGuidanceType) Unstage(stage *Stage) *CodeFinalGuidanceType {
	__gong__unstage(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceTypes_mapString, codefinalguidancetype, codefinalguidancetype.Name)
	return codefinalguidancetype
}

// UnstageVoid removes codefinalguidancetype off the model stage
func (codefinalguidancetype *CodeFinalGuidanceType) UnstageVoid(stage *Stage) {
	codefinalguidancetype.Unstage(stage)
}

func (codefinalguidancetype *CodeFinalGuidanceType) StageVoid(stage *Stage) {
	codefinalguidancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefinalguidancetype *CodeFinalGuidanceType) GetName() (res string) {
	return codefinalguidancetype.Name
}

// for satisfaction of GongStruct interface
func (codefinalguidancetype *CodeFinalGuidanceType) SetName(name string) {
	codefinalguidancetype.Name = name
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

// Stage puts codeheightreferencetype to the model stage
func (codeheightreferencetype *CodeHeightReferenceType) Stage(stage *Stage) *CodeHeightReferenceType {
	__gong__stage(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceType_stagedOrder, stage.CodeHeightReferenceType_orderStaged, &stage.CodeHeightReferenceTypeOrder, stage.CodeHeightReferenceTypes_mapString, codeheightreferencetype, codeheightreferencetype.Name)
	return codeheightreferencetype
}

// StagePreserveOrder puts codeheightreferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeHeightReferenceTypeOrder
// - update stage.CodeHeightReferenceTypeOrder accordingly
func (codeheightreferencetype *CodeHeightReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceType_stagedOrder, stage.CodeHeightReferenceType_orderStaged, &stage.CodeHeightReferenceTypeOrder, stage.CodeHeightReferenceTypes_mapString, codeheightreferencetype, order, codeheightreferencetype.Name)
}

// Unstage removes codeheightreferencetype off the model stage
func (codeheightreferencetype *CodeHeightReferenceType) Unstage(stage *Stage) *CodeHeightReferenceType {
	__gong__unstage(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceTypes_mapString, codeheightreferencetype, codeheightreferencetype.Name)
	return codeheightreferencetype
}

// UnstageVoid removes codeheightreferencetype off the model stage
func (codeheightreferencetype *CodeHeightReferenceType) UnstageVoid(stage *Stage) {
	codeheightreferencetype.Unstage(stage)
}

func (codeheightreferencetype *CodeHeightReferenceType) StageVoid(stage *Stage) {
	codeheightreferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeheightreferencetype *CodeHeightReferenceType) GetName() (res string) {
	return codeheightreferencetype.Name
}

// for satisfaction of GongStruct interface
func (codeheightreferencetype *CodeHeightReferenceType) SetName(name string) {
	codeheightreferencetype.Name = name
}

// Stage puts codeholdingusagetype to the model stage
func (codeholdingusagetype *CodeHoldingUsageType) Stage(stage *Stage) *CodeHoldingUsageType {
	__gong__stage(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageType_stagedOrder, stage.CodeHoldingUsageType_orderStaged, &stage.CodeHoldingUsageTypeOrder, stage.CodeHoldingUsageTypes_mapString, codeholdingusagetype, codeholdingusagetype.Name)
	return codeholdingusagetype
}

// StagePreserveOrder puts codeholdingusagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeHoldingUsageTypeOrder
// - update stage.CodeHoldingUsageTypeOrder accordingly
func (codeholdingusagetype *CodeHoldingUsageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageType_stagedOrder, stage.CodeHoldingUsageType_orderStaged, &stage.CodeHoldingUsageTypeOrder, stage.CodeHoldingUsageTypes_mapString, codeholdingusagetype, order, codeholdingusagetype.Name)
}

// Unstage removes codeholdingusagetype off the model stage
func (codeholdingusagetype *CodeHoldingUsageType) Unstage(stage *Stage) *CodeHoldingUsageType {
	__gong__unstage(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageTypes_mapString, codeholdingusagetype, codeholdingusagetype.Name)
	return codeholdingusagetype
}

// UnstageVoid removes codeholdingusagetype off the model stage
func (codeholdingusagetype *CodeHoldingUsageType) UnstageVoid(stage *Stage) {
	codeholdingusagetype.Unstage(stage)
}

func (codeholdingusagetype *CodeHoldingUsageType) StageVoid(stage *Stage) {
	codeholdingusagetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeholdingusagetype *CodeHoldingUsageType) GetName() (res string) {
	return codeholdingusagetype.Name
}

// for satisfaction of GongStruct interface
func (codeholdingusagetype *CodeHoldingUsageType) SetName(name string) {
	codeholdingusagetype.Name = name
}

// Stage puts codeholdingusetype to the model stage
func (codeholdingusetype *CodeHoldingUseType) Stage(stage *Stage) *CodeHoldingUseType {
	__gong__stage(stage.CodeHoldingUseTypes, stage.CodeHoldingUseType_stagedOrder, stage.CodeHoldingUseType_orderStaged, &stage.CodeHoldingUseTypeOrder, stage.CodeHoldingUseTypes_mapString, codeholdingusetype, codeholdingusetype.Name)
	return codeholdingusetype
}

// StagePreserveOrder puts codeholdingusetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeHoldingUseTypeOrder
// - update stage.CodeHoldingUseTypeOrder accordingly
func (codeholdingusetype *CodeHoldingUseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeHoldingUseTypes, stage.CodeHoldingUseType_stagedOrder, stage.CodeHoldingUseType_orderStaged, &stage.CodeHoldingUseTypeOrder, stage.CodeHoldingUseTypes_mapString, codeholdingusetype, order, codeholdingusetype.Name)
}

// Unstage removes codeholdingusetype off the model stage
func (codeholdingusetype *CodeHoldingUseType) Unstage(stage *Stage) *CodeHoldingUseType {
	__gong__unstage(stage.CodeHoldingUseTypes, stage.CodeHoldingUseTypes_mapString, codeholdingusetype, codeholdingusetype.Name)
	return codeholdingusetype
}

// UnstageVoid removes codeholdingusetype off the model stage
func (codeholdingusetype *CodeHoldingUseType) UnstageVoid(stage *Stage) {
	codeholdingusetype.Unstage(stage)
}

func (codeholdingusetype *CodeHoldingUseType) StageVoid(stage *Stage) {
	codeholdingusetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeholdingusetype *CodeHoldingUseType) GetName() (res string) {
	return codeholdingusetype.Name
}

// for satisfaction of GongStruct interface
func (codeholdingusetype *CodeHoldingUseType) SetName(name string) {
	codeholdingusetype.Name = name
}

// Stage puts codeicaocountrytype to the model stage
func (codeicaocountrytype *CodeICAOCountryType) Stage(stage *Stage) *CodeICAOCountryType {
	__gong__stage(stage.CodeICAOCountryTypes, stage.CodeICAOCountryType_stagedOrder, stage.CodeICAOCountryType_orderStaged, &stage.CodeICAOCountryTypeOrder, stage.CodeICAOCountryTypes_mapString, codeicaocountrytype, codeicaocountrytype.Name)
	return codeicaocountrytype
}

// StagePreserveOrder puts codeicaocountrytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeICAOCountryTypeOrder
// - update stage.CodeICAOCountryTypeOrder accordingly
func (codeicaocountrytype *CodeICAOCountryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeICAOCountryTypes, stage.CodeICAOCountryType_stagedOrder, stage.CodeICAOCountryType_orderStaged, &stage.CodeICAOCountryTypeOrder, stage.CodeICAOCountryTypes_mapString, codeicaocountrytype, order, codeicaocountrytype.Name)
}

// Unstage removes codeicaocountrytype off the model stage
func (codeicaocountrytype *CodeICAOCountryType) Unstage(stage *Stage) *CodeICAOCountryType {
	__gong__unstage(stage.CodeICAOCountryTypes, stage.CodeICAOCountryTypes_mapString, codeicaocountrytype, codeicaocountrytype.Name)
	return codeicaocountrytype
}

// UnstageVoid removes codeicaocountrytype off the model stage
func (codeicaocountrytype *CodeICAOCountryType) UnstageVoid(stage *Stage) {
	codeicaocountrytype.Unstage(stage)
}

func (codeicaocountrytype *CodeICAOCountryType) StageVoid(stage *Stage) {
	codeicaocountrytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeicaocountrytype *CodeICAOCountryType) GetName() (res string) {
	return codeicaocountrytype.Name
}

// for satisfaction of GongStruct interface
func (codeicaocountrytype *CodeICAOCountryType) SetName(name string) {
	codeicaocountrytype.Name = name
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

// Stage puts codeminimafinalapproachpathtype to the model stage
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) Stage(stage *Stage) *CodeMinimaFinalApproachPathType {
	__gong__stage(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathType_stagedOrder, stage.CodeMinimaFinalApproachPathType_orderStaged, &stage.CodeMinimaFinalApproachPathTypeOrder, stage.CodeMinimaFinalApproachPathTypes_mapString, codeminimafinalapproachpathtype, codeminimafinalapproachpathtype.Name)
	return codeminimafinalapproachpathtype
}

// StagePreserveOrder puts codeminimafinalapproachpathtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMinimaFinalApproachPathTypeOrder
// - update stage.CodeMinimaFinalApproachPathTypeOrder accordingly
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathType_stagedOrder, stage.CodeMinimaFinalApproachPathType_orderStaged, &stage.CodeMinimaFinalApproachPathTypeOrder, stage.CodeMinimaFinalApproachPathTypes_mapString, codeminimafinalapproachpathtype, order, codeminimafinalapproachpathtype.Name)
}

// Unstage removes codeminimafinalapproachpathtype off the model stage
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) Unstage(stage *Stage) *CodeMinimaFinalApproachPathType {
	__gong__unstage(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathTypes_mapString, codeminimafinalapproachpathtype, codeminimafinalapproachpathtype.Name)
	return codeminimafinalapproachpathtype
}

// UnstageVoid removes codeminimafinalapproachpathtype off the model stage
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) UnstageVoid(stage *Stage) {
	codeminimafinalapproachpathtype.Unstage(stage)
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) StageVoid(stage *Stage) {
	codeminimafinalapproachpathtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GetName() (res string) {
	return codeminimafinalapproachpathtype.Name
}

// for satisfaction of GongStruct interface
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) SetName(name string) {
	codeminimafinalapproachpathtype.Name = name
}

// Stage puts codeminimumaltitudetype to the model stage
func (codeminimumaltitudetype *CodeMinimumAltitudeType) Stage(stage *Stage) *CodeMinimumAltitudeType {
	__gong__stage(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeType_stagedOrder, stage.CodeMinimumAltitudeType_orderStaged, &stage.CodeMinimumAltitudeTypeOrder, stage.CodeMinimumAltitudeTypes_mapString, codeminimumaltitudetype, codeminimumaltitudetype.Name)
	return codeminimumaltitudetype
}

// StagePreserveOrder puts codeminimumaltitudetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMinimumAltitudeTypeOrder
// - update stage.CodeMinimumAltitudeTypeOrder accordingly
func (codeminimumaltitudetype *CodeMinimumAltitudeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeType_stagedOrder, stage.CodeMinimumAltitudeType_orderStaged, &stage.CodeMinimumAltitudeTypeOrder, stage.CodeMinimumAltitudeTypes_mapString, codeminimumaltitudetype, order, codeminimumaltitudetype.Name)
}

// Unstage removes codeminimumaltitudetype off the model stage
func (codeminimumaltitudetype *CodeMinimumAltitudeType) Unstage(stage *Stage) *CodeMinimumAltitudeType {
	__gong__unstage(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeTypes_mapString, codeminimumaltitudetype, codeminimumaltitudetype.Name)
	return codeminimumaltitudetype
}

// UnstageVoid removes codeminimumaltitudetype off the model stage
func (codeminimumaltitudetype *CodeMinimumAltitudeType) UnstageVoid(stage *Stage) {
	codeminimumaltitudetype.Unstage(stage)
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) StageVoid(stage *Stage) {
	codeminimumaltitudetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeminimumaltitudetype *CodeMinimumAltitudeType) GetName() (res string) {
	return codeminimumaltitudetype.Name
}

// for satisfaction of GongStruct interface
func (codeminimumaltitudetype *CodeMinimumAltitudeType) SetName(name string) {
	codeminimumaltitudetype.Name = name
}

// Stage puts codeminimumheighttype to the model stage
func (codeminimumheighttype *CodeMinimumHeightType) Stage(stage *Stage) *CodeMinimumHeightType {
	__gong__stage(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightType_stagedOrder, stage.CodeMinimumHeightType_orderStaged, &stage.CodeMinimumHeightTypeOrder, stage.CodeMinimumHeightTypes_mapString, codeminimumheighttype, codeminimumheighttype.Name)
	return codeminimumheighttype
}

// StagePreserveOrder puts codeminimumheighttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMinimumHeightTypeOrder
// - update stage.CodeMinimumHeightTypeOrder accordingly
func (codeminimumheighttype *CodeMinimumHeightType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightType_stagedOrder, stage.CodeMinimumHeightType_orderStaged, &stage.CodeMinimumHeightTypeOrder, stage.CodeMinimumHeightTypes_mapString, codeminimumheighttype, order, codeminimumheighttype.Name)
}

// Unstage removes codeminimumheighttype off the model stage
func (codeminimumheighttype *CodeMinimumHeightType) Unstage(stage *Stage) *CodeMinimumHeightType {
	__gong__unstage(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightTypes_mapString, codeminimumheighttype, codeminimumheighttype.Name)
	return codeminimumheighttype
}

// UnstageVoid removes codeminimumheighttype off the model stage
func (codeminimumheighttype *CodeMinimumHeightType) UnstageVoid(stage *Stage) {
	codeminimumheighttype.Unstage(stage)
}

func (codeminimumheighttype *CodeMinimumHeightType) StageVoid(stage *Stage) {
	codeminimumheighttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeminimumheighttype *CodeMinimumHeightType) GetName() (res string) {
	return codeminimumheighttype.Name
}

// for satisfaction of GongStruct interface
func (codeminimumheighttype *CodeMinimumHeightType) SetName(name string) {
	codeminimumheighttype.Name = name
}

// Stage puts codemissedapproachtype to the model stage
func (codemissedapproachtype *CodeMissedApproachType) Stage(stage *Stage) *CodeMissedApproachType {
	__gong__stage(stage.CodeMissedApproachTypes, stage.CodeMissedApproachType_stagedOrder, stage.CodeMissedApproachType_orderStaged, &stage.CodeMissedApproachTypeOrder, stage.CodeMissedApproachTypes_mapString, codemissedapproachtype, codemissedapproachtype.Name)
	return codemissedapproachtype
}

// StagePreserveOrder puts codemissedapproachtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMissedApproachTypeOrder
// - update stage.CodeMissedApproachTypeOrder accordingly
func (codemissedapproachtype *CodeMissedApproachType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMissedApproachTypes, stage.CodeMissedApproachType_stagedOrder, stage.CodeMissedApproachType_orderStaged, &stage.CodeMissedApproachTypeOrder, stage.CodeMissedApproachTypes_mapString, codemissedapproachtype, order, codemissedapproachtype.Name)
}

// Unstage removes codemissedapproachtype off the model stage
func (codemissedapproachtype *CodeMissedApproachType) Unstage(stage *Stage) *CodeMissedApproachType {
	__gong__unstage(stage.CodeMissedApproachTypes, stage.CodeMissedApproachTypes_mapString, codemissedapproachtype, codemissedapproachtype.Name)
	return codemissedapproachtype
}

// UnstageVoid removes codemissedapproachtype off the model stage
func (codemissedapproachtype *CodeMissedApproachType) UnstageVoid(stage *Stage) {
	codemissedapproachtype.Unstage(stage)
}

func (codemissedapproachtype *CodeMissedApproachType) StageVoid(stage *Stage) {
	codemissedapproachtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemissedapproachtype *CodeMissedApproachType) GetName() (res string) {
	return codemissedapproachtype.Name
}

// for satisfaction of GongStruct interface
func (codemissedapproachtype *CodeMissedApproachType) SetName(name string) {
	codemissedapproachtype.Name = name
}

// Stage puts codenavigationarearestrictiontype to the model stage
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) Stage(stage *Stage) *CodeNavigationAreaRestrictionType {
	__gong__stage(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionType_stagedOrder, stage.CodeNavigationAreaRestrictionType_orderStaged, &stage.CodeNavigationAreaRestrictionTypeOrder, stage.CodeNavigationAreaRestrictionTypes_mapString, codenavigationarearestrictiontype, codenavigationarearestrictiontype.Name)
	return codenavigationarearestrictiontype
}

// StagePreserveOrder puts codenavigationarearestrictiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavigationAreaRestrictionTypeOrder
// - update stage.CodeNavigationAreaRestrictionTypeOrder accordingly
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionType_stagedOrder, stage.CodeNavigationAreaRestrictionType_orderStaged, &stage.CodeNavigationAreaRestrictionTypeOrder, stage.CodeNavigationAreaRestrictionTypes_mapString, codenavigationarearestrictiontype, order, codenavigationarearestrictiontype.Name)
}

// Unstage removes codenavigationarearestrictiontype off the model stage
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) Unstage(stage *Stage) *CodeNavigationAreaRestrictionType {
	__gong__unstage(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionTypes_mapString, codenavigationarearestrictiontype, codenavigationarearestrictiontype.Name)
	return codenavigationarearestrictiontype
}

// UnstageVoid removes codenavigationarearestrictiontype off the model stage
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) UnstageVoid(stage *Stage) {
	codenavigationarearestrictiontype.Unstage(stage)
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) StageVoid(stage *Stage) {
	codenavigationarearestrictiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GetName() (res string) {
	return codenavigationarearestrictiontype.Name
}

// for satisfaction of GongStruct interface
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) SetName(name string) {
	codenavigationarearestrictiontype.Name = name
}

// Stage puts codenavigationareatype to the model stage
func (codenavigationareatype *CodeNavigationAreaType) Stage(stage *Stage) *CodeNavigationAreaType {
	__gong__stage(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaType_stagedOrder, stage.CodeNavigationAreaType_orderStaged, &stage.CodeNavigationAreaTypeOrder, stage.CodeNavigationAreaTypes_mapString, codenavigationareatype, codenavigationareatype.Name)
	return codenavigationareatype
}

// StagePreserveOrder puts codenavigationareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavigationAreaTypeOrder
// - update stage.CodeNavigationAreaTypeOrder accordingly
func (codenavigationareatype *CodeNavigationAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaType_stagedOrder, stage.CodeNavigationAreaType_orderStaged, &stage.CodeNavigationAreaTypeOrder, stage.CodeNavigationAreaTypes_mapString, codenavigationareatype, order, codenavigationareatype.Name)
}

// Unstage removes codenavigationareatype off the model stage
func (codenavigationareatype *CodeNavigationAreaType) Unstage(stage *Stage) *CodeNavigationAreaType {
	__gong__unstage(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaTypes_mapString, codenavigationareatype, codenavigationareatype.Name)
	return codenavigationareatype
}

// UnstageVoid removes codenavigationareatype off the model stage
func (codenavigationareatype *CodeNavigationAreaType) UnstageVoid(stage *Stage) {
	codenavigationareatype.Unstage(stage)
}

func (codenavigationareatype *CodeNavigationAreaType) StageVoid(stage *Stage) {
	codenavigationareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavigationareatype *CodeNavigationAreaType) GetName() (res string) {
	return codenavigationareatype.Name
}

// for satisfaction of GongStruct interface
func (codenavigationareatype *CodeNavigationAreaType) SetName(name string) {
	codenavigationareatype.Name = name
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

// Stage puts codeprocedureavailabilitytype to the model stage
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) Stage(stage *Stage) *CodeProcedureAvailabilityType {
	__gong__stage(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityType_stagedOrder, stage.CodeProcedureAvailabilityType_orderStaged, &stage.CodeProcedureAvailabilityTypeOrder, stage.CodeProcedureAvailabilityTypes_mapString, codeprocedureavailabilitytype, codeprocedureavailabilitytype.Name)
	return codeprocedureavailabilitytype
}

// StagePreserveOrder puts codeprocedureavailabilitytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProcedureAvailabilityTypeOrder
// - update stage.CodeProcedureAvailabilityTypeOrder accordingly
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityType_stagedOrder, stage.CodeProcedureAvailabilityType_orderStaged, &stage.CodeProcedureAvailabilityTypeOrder, stage.CodeProcedureAvailabilityTypes_mapString, codeprocedureavailabilitytype, order, codeprocedureavailabilitytype.Name)
}

// Unstage removes codeprocedureavailabilitytype off the model stage
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) Unstage(stage *Stage) *CodeProcedureAvailabilityType {
	__gong__unstage(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityTypes_mapString, codeprocedureavailabilitytype, codeprocedureavailabilitytype.Name)
	return codeprocedureavailabilitytype
}

// UnstageVoid removes codeprocedureavailabilitytype off the model stage
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) UnstageVoid(stage *Stage) {
	codeprocedureavailabilitytype.Unstage(stage)
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) StageVoid(stage *Stage) {
	codeprocedureavailabilitytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GetName() (res string) {
	return codeprocedureavailabilitytype.Name
}

// for satisfaction of GongStruct interface
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) SetName(name string) {
	codeprocedureavailabilitytype.Name = name
}

// Stage puts codeprocedurecodingstandardtype to the model stage
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) Stage(stage *Stage) *CodeProcedureCodingStandardType {
	__gong__stage(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardType_stagedOrder, stage.CodeProcedureCodingStandardType_orderStaged, &stage.CodeProcedureCodingStandardTypeOrder, stage.CodeProcedureCodingStandardTypes_mapString, codeprocedurecodingstandardtype, codeprocedurecodingstandardtype.Name)
	return codeprocedurecodingstandardtype
}

// StagePreserveOrder puts codeprocedurecodingstandardtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProcedureCodingStandardTypeOrder
// - update stage.CodeProcedureCodingStandardTypeOrder accordingly
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardType_stagedOrder, stage.CodeProcedureCodingStandardType_orderStaged, &stage.CodeProcedureCodingStandardTypeOrder, stage.CodeProcedureCodingStandardTypes_mapString, codeprocedurecodingstandardtype, order, codeprocedurecodingstandardtype.Name)
}

// Unstage removes codeprocedurecodingstandardtype off the model stage
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) Unstage(stage *Stage) *CodeProcedureCodingStandardType {
	__gong__unstage(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardTypes_mapString, codeprocedurecodingstandardtype, codeprocedurecodingstandardtype.Name)
	return codeprocedurecodingstandardtype
}

// UnstageVoid removes codeprocedurecodingstandardtype off the model stage
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) UnstageVoid(stage *Stage) {
	codeprocedurecodingstandardtype.Unstage(stage)
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) StageVoid(stage *Stage) {
	codeprocedurecodingstandardtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GetName() (res string) {
	return codeprocedurecodingstandardtype.Name
}

// for satisfaction of GongStruct interface
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) SetName(name string) {
	codeprocedurecodingstandardtype.Name = name
}

// Stage puts codeproceduredistancetype to the model stage
func (codeproceduredistancetype *CodeProcedureDistanceType) Stage(stage *Stage) *CodeProcedureDistanceType {
	__gong__stage(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceType_stagedOrder, stage.CodeProcedureDistanceType_orderStaged, &stage.CodeProcedureDistanceTypeOrder, stage.CodeProcedureDistanceTypes_mapString, codeproceduredistancetype, codeproceduredistancetype.Name)
	return codeproceduredistancetype
}

// StagePreserveOrder puts codeproceduredistancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProcedureDistanceTypeOrder
// - update stage.CodeProcedureDistanceTypeOrder accordingly
func (codeproceduredistancetype *CodeProcedureDistanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceType_stagedOrder, stage.CodeProcedureDistanceType_orderStaged, &stage.CodeProcedureDistanceTypeOrder, stage.CodeProcedureDistanceTypes_mapString, codeproceduredistancetype, order, codeproceduredistancetype.Name)
}

// Unstage removes codeproceduredistancetype off the model stage
func (codeproceduredistancetype *CodeProcedureDistanceType) Unstage(stage *Stage) *CodeProcedureDistanceType {
	__gong__unstage(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceTypes_mapString, codeproceduredistancetype, codeproceduredistancetype.Name)
	return codeproceduredistancetype
}

// UnstageVoid removes codeproceduredistancetype off the model stage
func (codeproceduredistancetype *CodeProcedureDistanceType) UnstageVoid(stage *Stage) {
	codeproceduredistancetype.Unstage(stage)
}

func (codeproceduredistancetype *CodeProcedureDistanceType) StageVoid(stage *Stage) {
	codeproceduredistancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeproceduredistancetype *CodeProcedureDistanceType) GetName() (res string) {
	return codeproceduredistancetype.Name
}

// for satisfaction of GongStruct interface
func (codeproceduredistancetype *CodeProcedureDistanceType) SetName(name string) {
	codeproceduredistancetype.Name = name
}

// Stage puts codeprocedurefixroletype to the model stage
func (codeprocedurefixroletype *CodeProcedureFixRoleType) Stage(stage *Stage) *CodeProcedureFixRoleType {
	__gong__stage(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleType_stagedOrder, stage.CodeProcedureFixRoleType_orderStaged, &stage.CodeProcedureFixRoleTypeOrder, stage.CodeProcedureFixRoleTypes_mapString, codeprocedurefixroletype, codeprocedurefixroletype.Name)
	return codeprocedurefixroletype
}

// StagePreserveOrder puts codeprocedurefixroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProcedureFixRoleTypeOrder
// - update stage.CodeProcedureFixRoleTypeOrder accordingly
func (codeprocedurefixroletype *CodeProcedureFixRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleType_stagedOrder, stage.CodeProcedureFixRoleType_orderStaged, &stage.CodeProcedureFixRoleTypeOrder, stage.CodeProcedureFixRoleTypes_mapString, codeprocedurefixroletype, order, codeprocedurefixroletype.Name)
}

// Unstage removes codeprocedurefixroletype off the model stage
func (codeprocedurefixroletype *CodeProcedureFixRoleType) Unstage(stage *Stage) *CodeProcedureFixRoleType {
	__gong__unstage(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleTypes_mapString, codeprocedurefixroletype, codeprocedurefixroletype.Name)
	return codeprocedurefixroletype
}

// UnstageVoid removes codeprocedurefixroletype off the model stage
func (codeprocedurefixroletype *CodeProcedureFixRoleType) UnstageVoid(stage *Stage) {
	codeprocedurefixroletype.Unstage(stage)
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) StageVoid(stage *Stage) {
	codeprocedurefixroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprocedurefixroletype *CodeProcedureFixRoleType) GetName() (res string) {
	return codeprocedurefixroletype.Name
}

// for satisfaction of GongStruct interface
func (codeprocedurefixroletype *CodeProcedureFixRoleType) SetName(name string) {
	codeprocedurefixroletype.Name = name
}

// Stage puts codeprocedurephasetype to the model stage
func (codeprocedurephasetype *CodeProcedurePhaseType) Stage(stage *Stage) *CodeProcedurePhaseType {
	__gong__stage(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseType_stagedOrder, stage.CodeProcedurePhaseType_orderStaged, &stage.CodeProcedurePhaseTypeOrder, stage.CodeProcedurePhaseTypes_mapString, codeprocedurephasetype, codeprocedurephasetype.Name)
	return codeprocedurephasetype
}

// StagePreserveOrder puts codeprocedurephasetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProcedurePhaseTypeOrder
// - update stage.CodeProcedurePhaseTypeOrder accordingly
func (codeprocedurephasetype *CodeProcedurePhaseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseType_stagedOrder, stage.CodeProcedurePhaseType_orderStaged, &stage.CodeProcedurePhaseTypeOrder, stage.CodeProcedurePhaseTypes_mapString, codeprocedurephasetype, order, codeprocedurephasetype.Name)
}

// Unstage removes codeprocedurephasetype off the model stage
func (codeprocedurephasetype *CodeProcedurePhaseType) Unstage(stage *Stage) *CodeProcedurePhaseType {
	__gong__unstage(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseTypes_mapString, codeprocedurephasetype, codeprocedurephasetype.Name)
	return codeprocedurephasetype
}

// UnstageVoid removes codeprocedurephasetype off the model stage
func (codeprocedurephasetype *CodeProcedurePhaseType) UnstageVoid(stage *Stage) {
	codeprocedurephasetype.Unstage(stage)
}

func (codeprocedurephasetype *CodeProcedurePhaseType) StageVoid(stage *Stage) {
	codeprocedurephasetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprocedurephasetype *CodeProcedurePhaseType) GetName() (res string) {
	return codeprocedurephasetype.Name
}

// for satisfaction of GongStruct interface
func (codeprocedurephasetype *CodeProcedurePhaseType) SetName(name string) {
	codeprocedurephasetype.Name = name
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

// Stage puts codereferencepathidentifiertype to the model stage
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) Stage(stage *Stage) *CodeReferencePathIdentifierType {
	__gong__stage(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierType_stagedOrder, stage.CodeReferencePathIdentifierType_orderStaged, &stage.CodeReferencePathIdentifierTypeOrder, stage.CodeReferencePathIdentifierTypes_mapString, codereferencepathidentifiertype, codereferencepathidentifiertype.Name)
	return codereferencepathidentifiertype
}

// StagePreserveOrder puts codereferencepathidentifiertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeReferencePathIdentifierTypeOrder
// - update stage.CodeReferencePathIdentifierTypeOrder accordingly
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierType_stagedOrder, stage.CodeReferencePathIdentifierType_orderStaged, &stage.CodeReferencePathIdentifierTypeOrder, stage.CodeReferencePathIdentifierTypes_mapString, codereferencepathidentifiertype, order, codereferencepathidentifiertype.Name)
}

// Unstage removes codereferencepathidentifiertype off the model stage
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) Unstage(stage *Stage) *CodeReferencePathIdentifierType {
	__gong__unstage(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierTypes_mapString, codereferencepathidentifiertype, codereferencepathidentifiertype.Name)
	return codereferencepathidentifiertype
}

// UnstageVoid removes codereferencepathidentifiertype off the model stage
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) UnstageVoid(stage *Stage) {
	codereferencepathidentifiertype.Unstage(stage)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) StageVoid(stage *Stage) {
	codereferencepathidentifiertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GetName() (res string) {
	return codereferencepathidentifiertype.Name
}

// for satisfaction of GongStruct interface
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) SetName(name string) {
	codereferencepathidentifiertype.Name = name
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

// Stage puts coderelativepositiontype to the model stage
func (coderelativepositiontype *CodeRelativePositionType) Stage(stage *Stage) *CodeRelativePositionType {
	__gong__stage(stage.CodeRelativePositionTypes, stage.CodeRelativePositionType_stagedOrder, stage.CodeRelativePositionType_orderStaged, &stage.CodeRelativePositionTypeOrder, stage.CodeRelativePositionTypes_mapString, coderelativepositiontype, coderelativepositiontype.Name)
	return coderelativepositiontype
}

// StagePreserveOrder puts coderelativepositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRelativePositionTypeOrder
// - update stage.CodeRelativePositionTypeOrder accordingly
func (coderelativepositiontype *CodeRelativePositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRelativePositionTypes, stage.CodeRelativePositionType_stagedOrder, stage.CodeRelativePositionType_orderStaged, &stage.CodeRelativePositionTypeOrder, stage.CodeRelativePositionTypes_mapString, coderelativepositiontype, order, coderelativepositiontype.Name)
}

// Unstage removes coderelativepositiontype off the model stage
func (coderelativepositiontype *CodeRelativePositionType) Unstage(stage *Stage) *CodeRelativePositionType {
	__gong__unstage(stage.CodeRelativePositionTypes, stage.CodeRelativePositionTypes_mapString, coderelativepositiontype, coderelativepositiontype.Name)
	return coderelativepositiontype
}

// UnstageVoid removes coderelativepositiontype off the model stage
func (coderelativepositiontype *CodeRelativePositionType) UnstageVoid(stage *Stage) {
	coderelativepositiontype.Unstage(stage)
}

func (coderelativepositiontype *CodeRelativePositionType) StageVoid(stage *Stage) {
	coderelativepositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderelativepositiontype *CodeRelativePositionType) GetName() (res string) {
	return coderelativepositiontype.Name
}

// for satisfaction of GongStruct interface
func (coderelativepositiontype *CodeRelativePositionType) SetName(name string) {
	coderelativepositiontype.Name = name
}

// Stage puts coderouteindicatortype to the model stage
func (coderouteindicatortype *CodeRouteIndicatorType) Stage(stage *Stage) *CodeRouteIndicatorType {
	__gong__stage(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorType_stagedOrder, stage.CodeRouteIndicatorType_orderStaged, &stage.CodeRouteIndicatorTypeOrder, stage.CodeRouteIndicatorTypes_mapString, coderouteindicatortype, coderouteindicatortype.Name)
	return coderouteindicatortype
}

// StagePreserveOrder puts coderouteindicatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRouteIndicatorTypeOrder
// - update stage.CodeRouteIndicatorTypeOrder accordingly
func (coderouteindicatortype *CodeRouteIndicatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorType_stagedOrder, stage.CodeRouteIndicatorType_orderStaged, &stage.CodeRouteIndicatorTypeOrder, stage.CodeRouteIndicatorTypes_mapString, coderouteindicatortype, order, coderouteindicatortype.Name)
}

// Unstage removes coderouteindicatortype off the model stage
func (coderouteindicatortype *CodeRouteIndicatorType) Unstage(stage *Stage) *CodeRouteIndicatorType {
	__gong__unstage(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorTypes_mapString, coderouteindicatortype, coderouteindicatortype.Name)
	return coderouteindicatortype
}

// UnstageVoid removes coderouteindicatortype off the model stage
func (coderouteindicatortype *CodeRouteIndicatorType) UnstageVoid(stage *Stage) {
	coderouteindicatortype.Unstage(stage)
}

func (coderouteindicatortype *CodeRouteIndicatorType) StageVoid(stage *Stage) {
	coderouteindicatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderouteindicatortype *CodeRouteIndicatorType) GetName() (res string) {
	return coderouteindicatortype.Name
}

// for satisfaction of GongStruct interface
func (coderouteindicatortype *CodeRouteIndicatorType) SetName(name string) {
	coderouteindicatortype.Name = name
}

// Stage puts codesafealtitudetype to the model stage
func (codesafealtitudetype *CodeSafeAltitudeType) Stage(stage *Stage) *CodeSafeAltitudeType {
	__gong__stage(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeType_stagedOrder, stage.CodeSafeAltitudeType_orderStaged, &stage.CodeSafeAltitudeTypeOrder, stage.CodeSafeAltitudeTypes_mapString, codesafealtitudetype, codesafealtitudetype.Name)
	return codesafealtitudetype
}

// StagePreserveOrder puts codesafealtitudetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSafeAltitudeTypeOrder
// - update stage.CodeSafeAltitudeTypeOrder accordingly
func (codesafealtitudetype *CodeSafeAltitudeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeType_stagedOrder, stage.CodeSafeAltitudeType_orderStaged, &stage.CodeSafeAltitudeTypeOrder, stage.CodeSafeAltitudeTypes_mapString, codesafealtitudetype, order, codesafealtitudetype.Name)
}

// Unstage removes codesafealtitudetype off the model stage
func (codesafealtitudetype *CodeSafeAltitudeType) Unstage(stage *Stage) *CodeSafeAltitudeType {
	__gong__unstage(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeTypes_mapString, codesafealtitudetype, codesafealtitudetype.Name)
	return codesafealtitudetype
}

// UnstageVoid removes codesafealtitudetype off the model stage
func (codesafealtitudetype *CodeSafeAltitudeType) UnstageVoid(stage *Stage) {
	codesafealtitudetype.Unstage(stage)
}

func (codesafealtitudetype *CodeSafeAltitudeType) StageVoid(stage *Stage) {
	codesafealtitudetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesafealtitudetype *CodeSafeAltitudeType) GetName() (res string) {
	return codesafealtitudetype.Name
}

// for satisfaction of GongStruct interface
func (codesafealtitudetype *CodeSafeAltitudeType) SetName(name string) {
	codesafealtitudetype.Name = name
}

// Stage puts codesegmentpathtype to the model stage
func (codesegmentpathtype *CodeSegmentPathType) Stage(stage *Stage) *CodeSegmentPathType {
	__gong__stage(stage.CodeSegmentPathTypes, stage.CodeSegmentPathType_stagedOrder, stage.CodeSegmentPathType_orderStaged, &stage.CodeSegmentPathTypeOrder, stage.CodeSegmentPathTypes_mapString, codesegmentpathtype, codesegmentpathtype.Name)
	return codesegmentpathtype
}

// StagePreserveOrder puts codesegmentpathtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSegmentPathTypeOrder
// - update stage.CodeSegmentPathTypeOrder accordingly
func (codesegmentpathtype *CodeSegmentPathType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSegmentPathTypes, stage.CodeSegmentPathType_stagedOrder, stage.CodeSegmentPathType_orderStaged, &stage.CodeSegmentPathTypeOrder, stage.CodeSegmentPathTypes_mapString, codesegmentpathtype, order, codesegmentpathtype.Name)
}

// Unstage removes codesegmentpathtype off the model stage
func (codesegmentpathtype *CodeSegmentPathType) Unstage(stage *Stage) *CodeSegmentPathType {
	__gong__unstage(stage.CodeSegmentPathTypes, stage.CodeSegmentPathTypes_mapString, codesegmentpathtype, codesegmentpathtype.Name)
	return codesegmentpathtype
}

// UnstageVoid removes codesegmentpathtype off the model stage
func (codesegmentpathtype *CodeSegmentPathType) UnstageVoid(stage *Stage) {
	codesegmentpathtype.Unstage(stage)
}

func (codesegmentpathtype *CodeSegmentPathType) StageVoid(stage *Stage) {
	codesegmentpathtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesegmentpathtype *CodeSegmentPathType) GetName() (res string) {
	return codesegmentpathtype.Name
}

// for satisfaction of GongStruct interface
func (codesegmentpathtype *CodeSegmentPathType) SetName(name string) {
	codesegmentpathtype.Name = name
}

// Stage puts codesegmentterminationtype to the model stage
func (codesegmentterminationtype *CodeSegmentTerminationType) Stage(stage *Stage) *CodeSegmentTerminationType {
	__gong__stage(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationType_stagedOrder, stage.CodeSegmentTerminationType_orderStaged, &stage.CodeSegmentTerminationTypeOrder, stage.CodeSegmentTerminationTypes_mapString, codesegmentterminationtype, codesegmentterminationtype.Name)
	return codesegmentterminationtype
}

// StagePreserveOrder puts codesegmentterminationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSegmentTerminationTypeOrder
// - update stage.CodeSegmentTerminationTypeOrder accordingly
func (codesegmentterminationtype *CodeSegmentTerminationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationType_stagedOrder, stage.CodeSegmentTerminationType_orderStaged, &stage.CodeSegmentTerminationTypeOrder, stage.CodeSegmentTerminationTypes_mapString, codesegmentterminationtype, order, codesegmentterminationtype.Name)
}

// Unstage removes codesegmentterminationtype off the model stage
func (codesegmentterminationtype *CodeSegmentTerminationType) Unstage(stage *Stage) *CodeSegmentTerminationType {
	__gong__unstage(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationTypes_mapString, codesegmentterminationtype, codesegmentterminationtype.Name)
	return codesegmentterminationtype
}

// UnstageVoid removes codesegmentterminationtype off the model stage
func (codesegmentterminationtype *CodeSegmentTerminationType) UnstageVoid(stage *Stage) {
	codesegmentterminationtype.Unstage(stage)
}

func (codesegmentterminationtype *CodeSegmentTerminationType) StageVoid(stage *Stage) {
	codesegmentterminationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesegmentterminationtype *CodeSegmentTerminationType) GetName() (res string) {
	return codesegmentterminationtype.Name
}

// for satisfaction of GongStruct interface
func (codesegmentterminationtype *CodeSegmentTerminationType) SetName(name string) {
	codesegmentterminationtype.Name = name
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

// Stage puts codetaatype to the model stage
func (codetaatype *CodeTAAType) Stage(stage *Stage) *CodeTAAType {
	__gong__stage(stage.CodeTAATypes, stage.CodeTAAType_stagedOrder, stage.CodeTAAType_orderStaged, &stage.CodeTAATypeOrder, stage.CodeTAATypes_mapString, codetaatype, codetaatype.Name)
	return codetaatype
}

// StagePreserveOrder puts codetaatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTAATypeOrder
// - update stage.CodeTAATypeOrder accordingly
func (codetaatype *CodeTAAType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTAATypes, stage.CodeTAAType_stagedOrder, stage.CodeTAAType_orderStaged, &stage.CodeTAATypeOrder, stage.CodeTAATypes_mapString, codetaatype, order, codetaatype.Name)
}

// Unstage removes codetaatype off the model stage
func (codetaatype *CodeTAAType) Unstage(stage *Stage) *CodeTAAType {
	__gong__unstage(stage.CodeTAATypes, stage.CodeTAATypes_mapString, codetaatype, codetaatype.Name)
	return codetaatype
}

// UnstageVoid removes codetaatype off the model stage
func (codetaatype *CodeTAAType) UnstageVoid(stage *Stage) {
	codetaatype.Unstage(stage)
}

func (codetaatype *CodeTAAType) StageVoid(stage *Stage) {
	codetaatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetaatype *CodeTAAType) GetName() (res string) {
	return codetaatype.Name
}

// for satisfaction of GongStruct interface
func (codetaatype *CodeTAAType) SetName(name string) {
	codetaatype.Name = name
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

// Stage puts codetrajectorytype to the model stage
func (codetrajectorytype *CodeTrajectoryType) Stage(stage *Stage) *CodeTrajectoryType {
	__gong__stage(stage.CodeTrajectoryTypes, stage.CodeTrajectoryType_stagedOrder, stage.CodeTrajectoryType_orderStaged, &stage.CodeTrajectoryTypeOrder, stage.CodeTrajectoryTypes_mapString, codetrajectorytype, codetrajectorytype.Name)
	return codetrajectorytype
}

// StagePreserveOrder puts codetrajectorytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTrajectoryTypeOrder
// - update stage.CodeTrajectoryTypeOrder accordingly
func (codetrajectorytype *CodeTrajectoryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTrajectoryTypes, stage.CodeTrajectoryType_stagedOrder, stage.CodeTrajectoryType_orderStaged, &stage.CodeTrajectoryTypeOrder, stage.CodeTrajectoryTypes_mapString, codetrajectorytype, order, codetrajectorytype.Name)
}

// Unstage removes codetrajectorytype off the model stage
func (codetrajectorytype *CodeTrajectoryType) Unstage(stage *Stage) *CodeTrajectoryType {
	__gong__unstage(stage.CodeTrajectoryTypes, stage.CodeTrajectoryTypes_mapString, codetrajectorytype, codetrajectorytype.Name)
	return codetrajectorytype
}

// UnstageVoid removes codetrajectorytype off the model stage
func (codetrajectorytype *CodeTrajectoryType) UnstageVoid(stage *Stage) {
	codetrajectorytype.Unstage(stage)
}

func (codetrajectorytype *CodeTrajectoryType) StageVoid(stage *Stage) {
	codetrajectorytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetrajectorytype *CodeTrajectoryType) GetName() (res string) {
	return codetrajectorytype.Name
}

// for satisfaction of GongStruct interface
func (codetrajectorytype *CodeTrajectoryType) SetName(name string) {
	codetrajectorytype.Name = name
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

// Stage puts textsidstardesignatortype to the model stage
func (textsidstardesignatortype *TextSIDSTARDesignatorType) Stage(stage *Stage) *TextSIDSTARDesignatorType {
	__gong__stage(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorType_stagedOrder, stage.TextSIDSTARDesignatorType_orderStaged, &stage.TextSIDSTARDesignatorTypeOrder, stage.TextSIDSTARDesignatorTypes_mapString, textsidstardesignatortype, textsidstardesignatortype.Name)
	return textsidstardesignatortype
}

// StagePreserveOrder puts textsidstardesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextSIDSTARDesignatorTypeOrder
// - update stage.TextSIDSTARDesignatorTypeOrder accordingly
func (textsidstardesignatortype *TextSIDSTARDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorType_stagedOrder, stage.TextSIDSTARDesignatorType_orderStaged, &stage.TextSIDSTARDesignatorTypeOrder, stage.TextSIDSTARDesignatorTypes_mapString, textsidstardesignatortype, order, textsidstardesignatortype.Name)
}

// Unstage removes textsidstardesignatortype off the model stage
func (textsidstardesignatortype *TextSIDSTARDesignatorType) Unstage(stage *Stage) *TextSIDSTARDesignatorType {
	__gong__unstage(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorTypes_mapString, textsidstardesignatortype, textsidstardesignatortype.Name)
	return textsidstardesignatortype
}

// UnstageVoid removes textsidstardesignatortype off the model stage
func (textsidstardesignatortype *TextSIDSTARDesignatorType) UnstageVoid(stage *Stage) {
	textsidstardesignatortype.Unstage(stage)
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) StageVoid(stage *Stage) {
	textsidstardesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textsidstardesignatortype *TextSIDSTARDesignatorType) GetName() (res string) {
	return textsidstardesignatortype.Name
}

// for satisfaction of GongStruct interface
func (textsidstardesignatortype *TextSIDSTARDesignatorType) SetName(name string) {
	textsidstardesignatortype.Name = name
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

// Stage puts valalarmlimittype to the model stage
func (valalarmlimittype *ValAlarmLimitType) Stage(stage *Stage) *ValAlarmLimitType {
	__gong__stage(stage.ValAlarmLimitTypes, stage.ValAlarmLimitType_stagedOrder, stage.ValAlarmLimitType_orderStaged, &stage.ValAlarmLimitTypeOrder, stage.ValAlarmLimitTypes_mapString, valalarmlimittype, valalarmlimittype.Name)
	return valalarmlimittype
}

// StagePreserveOrder puts valalarmlimittype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValAlarmLimitTypeOrder
// - update stage.ValAlarmLimitTypeOrder accordingly
func (valalarmlimittype *ValAlarmLimitType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValAlarmLimitTypes, stage.ValAlarmLimitType_stagedOrder, stage.ValAlarmLimitType_orderStaged, &stage.ValAlarmLimitTypeOrder, stage.ValAlarmLimitTypes_mapString, valalarmlimittype, order, valalarmlimittype.Name)
}

// Unstage removes valalarmlimittype off the model stage
func (valalarmlimittype *ValAlarmLimitType) Unstage(stage *Stage) *ValAlarmLimitType {
	__gong__unstage(stage.ValAlarmLimitTypes, stage.ValAlarmLimitTypes_mapString, valalarmlimittype, valalarmlimittype.Name)
	return valalarmlimittype
}

// UnstageVoid removes valalarmlimittype off the model stage
func (valalarmlimittype *ValAlarmLimitType) UnstageVoid(stage *Stage) {
	valalarmlimittype.Unstage(stage)
}

func (valalarmlimittype *ValAlarmLimitType) StageVoid(stage *Stage) {
	valalarmlimittype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valalarmlimittype *ValAlarmLimitType) GetName() (res string) {
	return valalarmlimittype.Name
}

// for satisfaction of GongStruct interface
func (valalarmlimittype *ValAlarmLimitType) SetName(name string) {
	valalarmlimittype.Name = name
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

// Stage puts valchannelnumbertype to the model stage
func (valchannelnumbertype *ValChannelNumberType) Stage(stage *Stage) *ValChannelNumberType {
	__gong__stage(stage.ValChannelNumberTypes, stage.ValChannelNumberType_stagedOrder, stage.ValChannelNumberType_orderStaged, &stage.ValChannelNumberTypeOrder, stage.ValChannelNumberTypes_mapString, valchannelnumbertype, valchannelnumbertype.Name)
	return valchannelnumbertype
}

// StagePreserveOrder puts valchannelnumbertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValChannelNumberTypeOrder
// - update stage.ValChannelNumberTypeOrder accordingly
func (valchannelnumbertype *ValChannelNumberType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValChannelNumberTypes, stage.ValChannelNumberType_stagedOrder, stage.ValChannelNumberType_orderStaged, &stage.ValChannelNumberTypeOrder, stage.ValChannelNumberTypes_mapString, valchannelnumbertype, order, valchannelnumbertype.Name)
}

// Unstage removes valchannelnumbertype off the model stage
func (valchannelnumbertype *ValChannelNumberType) Unstage(stage *Stage) *ValChannelNumberType {
	__gong__unstage(stage.ValChannelNumberTypes, stage.ValChannelNumberTypes_mapString, valchannelnumbertype, valchannelnumbertype.Name)
	return valchannelnumbertype
}

// UnstageVoid removes valchannelnumbertype off the model stage
func (valchannelnumbertype *ValChannelNumberType) UnstageVoid(stage *Stage) {
	valchannelnumbertype.Unstage(stage)
}

func (valchannelnumbertype *ValChannelNumberType) StageVoid(stage *Stage) {
	valchannelnumbertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valchannelnumbertype *ValChannelNumberType) GetName() (res string) {
	return valchannelnumbertype.Name
}

// for satisfaction of GongStruct interface
func (valchannelnumbertype *ValChannelNumberType) SetName(name string) {
	valchannelnumbertype.Name = name
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

// Stage puts valhextype to the model stage
func (valhextype *ValHexType) Stage(stage *Stage) *ValHexType {
	__gong__stage(stage.ValHexTypes, stage.ValHexType_stagedOrder, stage.ValHexType_orderStaged, &stage.ValHexTypeOrder, stage.ValHexTypes_mapString, valhextype, valhextype.Name)
	return valhextype
}

// StagePreserveOrder puts valhextype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValHexTypeOrder
// - update stage.ValHexTypeOrder accordingly
func (valhextype *ValHexType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValHexTypes, stage.ValHexType_stagedOrder, stage.ValHexType_orderStaged, &stage.ValHexTypeOrder, stage.ValHexTypes_mapString, valhextype, order, valhextype.Name)
}

// Unstage removes valhextype off the model stage
func (valhextype *ValHexType) Unstage(stage *Stage) *ValHexType {
	__gong__unstage(stage.ValHexTypes, stage.ValHexTypes_mapString, valhextype, valhextype.Name)
	return valhextype
}

// UnstageVoid removes valhextype off the model stage
func (valhextype *ValHexType) UnstageVoid(stage *Stage) {
	valhextype.Unstage(stage)
}

func (valhextype *ValHexType) StageVoid(stage *Stage) {
	valhextype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valhextype *ValHexType) GetName() (res string) {
	return valhextype.Name
}

// for satisfaction of GongStruct interface
func (valhextype *ValHexType) SetName(name string) {
	valhextype.Name = name
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

// Stage puts valtemperaturetype to the model stage
func (valtemperaturetype *ValTemperatureType) Stage(stage *Stage) *ValTemperatureType {
	__gong__stage(stage.ValTemperatureTypes, stage.ValTemperatureType_stagedOrder, stage.ValTemperatureType_orderStaged, &stage.ValTemperatureTypeOrder, stage.ValTemperatureTypes_mapString, valtemperaturetype, valtemperaturetype.Name)
	return valtemperaturetype
}

// StagePreserveOrder puts valtemperaturetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValTemperatureTypeOrder
// - update stage.ValTemperatureTypeOrder accordingly
func (valtemperaturetype *ValTemperatureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValTemperatureTypes, stage.ValTemperatureType_stagedOrder, stage.ValTemperatureType_orderStaged, &stage.ValTemperatureTypeOrder, stage.ValTemperatureTypes_mapString, valtemperaturetype, order, valtemperaturetype.Name)
}

// Unstage removes valtemperaturetype off the model stage
func (valtemperaturetype *ValTemperatureType) Unstage(stage *Stage) *ValTemperatureType {
	__gong__unstage(stage.ValTemperatureTypes, stage.ValTemperatureTypes_mapString, valtemperaturetype, valtemperaturetype.Name)
	return valtemperaturetype
}

// UnstageVoid removes valtemperaturetype off the model stage
func (valtemperaturetype *ValTemperatureType) UnstageVoid(stage *Stage) {
	valtemperaturetype.Unstage(stage)
}

func (valtemperaturetype *ValTemperatureType) StageVoid(stage *Stage) {
	valtemperaturetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valtemperaturetype *ValTemperatureType) GetName() (res string) {
	return valtemperaturetype.Name
}

// for satisfaction of GongStruct interface
func (valtemperaturetype *ValTemperatureType) SetName(name string) {
	valtemperaturetype.Name = name
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

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeATCReportingTypes, &stage.CodeATCReportingTypes_mapString, &stage.CodeATCReportingType_stagedOrder, &stage.CodeATCReportingTypeOrder)

	__gong__resetStageType(&stage.CodeAerialRefuellingPointTypes, &stage.CodeAerialRefuellingPointTypes_mapString, &stage.CodeAerialRefuellingPointType_stagedOrder, &stage.CodeAerialRefuellingPointTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString, &stage.CodeAircraftCategoryType_stagedOrder, &stage.CodeAircraftCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString, &stage.CodeAircraftEngineNumberType_stagedOrder, &stage.CodeAircraftEngineNumberTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString, &stage.CodeAircraftEngineType_stagedOrder, &stage.CodeAircraftEngineTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString, &stage.CodeAircraftICAOType_stagedOrder, &stage.CodeAircraftICAOTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString, &stage.CodeAircraftType_stagedOrder, &stage.CodeAircraftTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString, &stage.CodeAircraftWingspanClassType_stagedOrder, &stage.CodeAircraftWingspanClassTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_mapString, &stage.CodeAltitudeAdjustmentType_stagedOrder, &stage.CodeAltitudeAdjustmentTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString, &stage.CodeAltitudeUseType_stagedOrder, &stage.CodeAltitudeUseTypeOrder)

	__gong__resetStageType(&stage.CodeApproachEquipmentAdditionalTypes, &stage.CodeApproachEquipmentAdditionalTypes_mapString, &stage.CodeApproachEquipmentAdditionalType_stagedOrder, &stage.CodeApproachEquipmentAdditionalTypeOrder)

	__gong__resetStageType(&stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_mapString, &stage.CodeApproachGuidanceType_stagedOrder, &stage.CodeApproachGuidanceTypeOrder)

	__gong__resetStageType(&stage.CodeApproachPrefixTypes, &stage.CodeApproachPrefixTypes_mapString, &stage.CodeApproachPrefixType_stagedOrder, &stage.CodeApproachPrefixTypeOrder)

	__gong__resetStageType(&stage.CodeApproachTypes, &stage.CodeApproachTypes_mapString, &stage.CodeApproachType_stagedOrder, &stage.CodeApproachTypeOrder)

	__gong__resetStageType(&stage.CodeApprovalTypes, &stage.CodeApprovalTypes_mapString, &stage.CodeApprovalType_stagedOrder, &stage.CodeApprovalTypeOrder)

	__gong__resetStageType(&stage.CodeArcDirectionTypes, &stage.CodeArcDirectionTypes_mapString, &stage.CodeArcDirectionType_stagedOrder, &stage.CodeArcDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeBearingTypes, &stage.CodeBearingTypes_mapString, &stage.CodeBearingType_stagedOrder, &stage.CodeBearingTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString, &stage.CodeCommunicationModeType_stagedOrder, &stage.CodeCommunicationModeTypeOrder)

	__gong__resetStageType(&stage.CodeCourseTypes, &stage.CodeCourseTypes_mapString, &stage.CodeCourseType_stagedOrder, &stage.CodeCourseTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeDesignStandardTypes, &stage.CodeDesignStandardTypes_mapString, &stage.CodeDesignStandardType_stagedOrder, &stage.CodeDesignStandardTypeOrder)

	__gong__resetStageType(&stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_mapString, &stage.CodeDesignatedPointDesignatorType_stagedOrder, &stage.CodeDesignatedPointDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_mapString, &stage.CodeDirectionReferenceType_stagedOrder, &stage.CodeDirectionReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString, &stage.CodeDirectionTurnType_stagedOrder, &stage.CodeDirectionTurnTypeOrder)

	__gong__resetStageType(&stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString, &stage.CodeEquipmentAntiCollisionType_stagedOrder, &stage.CodeEquipmentAntiCollisionTypeOrder)

	__gong__resetStageType(&stage.CodeEquipmentUnavailableTypes, &stage.CodeEquipmentUnavailableTypes_mapString, &stage.CodeEquipmentUnavailableType_stagedOrder, &stage.CodeEquipmentUnavailableTypeOrder)

	__gong__resetStageType(&stage.CodeFinalGuidanceTypes, &stage.CodeFinalGuidanceTypes_mapString, &stage.CodeFinalGuidanceType_stagedOrder, &stage.CodeFinalGuidanceTypeOrder)

	__gong__resetStageType(&stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_mapString, &stage.CodeFreeFlightType_stagedOrder, &stage.CodeFreeFlightTypeOrder)

	__gong__resetStageType(&stage.CodeHeightReferenceTypes, &stage.CodeHeightReferenceTypes_mapString, &stage.CodeHeightReferenceType_stagedOrder, &stage.CodeHeightReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeHoldingUsageTypes, &stage.CodeHoldingUsageTypes_mapString, &stage.CodeHoldingUsageType_stagedOrder, &stage.CodeHoldingUsageTypeOrder)

	__gong__resetStageType(&stage.CodeHoldingUseTypes, &stage.CodeHoldingUseTypes_mapString, &stage.CodeHoldingUseType_stagedOrder, &stage.CodeHoldingUseTypeOrder)

	__gong__resetStageType(&stage.CodeICAOCountryTypes, &stage.CodeICAOCountryTypes_mapString, &stage.CodeICAOCountryType_stagedOrder, &stage.CodeICAOCountryTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_mapString, &stage.CodeMilitaryRoutePointType_stagedOrder, &stage.CodeMilitaryRoutePointTypeOrder)

	__gong__resetStageType(&stage.CodeMinimaFinalApproachPathTypes, &stage.CodeMinimaFinalApproachPathTypes_mapString, &stage.CodeMinimaFinalApproachPathType_stagedOrder, &stage.CodeMinimaFinalApproachPathTypeOrder)

	__gong__resetStageType(&stage.CodeMinimumAltitudeTypes, &stage.CodeMinimumAltitudeTypes_mapString, &stage.CodeMinimumAltitudeType_stagedOrder, &stage.CodeMinimumAltitudeTypeOrder)

	__gong__resetStageType(&stage.CodeMinimumHeightTypes, &stage.CodeMinimumHeightTypes_mapString, &stage.CodeMinimumHeightType_stagedOrder, &stage.CodeMinimumHeightTypeOrder)

	__gong__resetStageType(&stage.CodeMissedApproachTypes, &stage.CodeMissedApproachTypes_mapString, &stage.CodeMissedApproachType_stagedOrder, &stage.CodeMissedApproachTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationAreaRestrictionTypes, &stage.CodeNavigationAreaRestrictionTypes_mapString, &stage.CodeNavigationAreaRestrictionType_stagedOrder, &stage.CodeNavigationAreaRestrictionTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationAreaTypes, &stage.CodeNavigationAreaTypes_mapString, &stage.CodeNavigationAreaType_stagedOrder, &stage.CodeNavigationAreaTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString, &stage.CodeNavigationEquipmentType_stagedOrder, &stage.CodeNavigationEquipmentTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString, &stage.CodeNavigationSpecificationType_stagedOrder, &stage.CodeNavigationSpecificationTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString, &stage.CodeObstacleAssessmentSurfaceType_stagedOrder, &stage.CodeObstacleAssessmentSurfaceTypeOrder)

	__gong__resetStageType(&stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_mapString, &stage.CodeObstructionIdSurfaceZoneType_stagedOrder, &stage.CodeObstructionIdSurfaceZoneTypeOrder)

	__gong__resetStageType(&stage.CodeProcedureAvailabilityTypes, &stage.CodeProcedureAvailabilityTypes_mapString, &stage.CodeProcedureAvailabilityType_stagedOrder, &stage.CodeProcedureAvailabilityTypeOrder)

	__gong__resetStageType(&stage.CodeProcedureCodingStandardTypes, &stage.CodeProcedureCodingStandardTypes_mapString, &stage.CodeProcedureCodingStandardType_stagedOrder, &stage.CodeProcedureCodingStandardTypeOrder)

	__gong__resetStageType(&stage.CodeProcedureDistanceTypes, &stage.CodeProcedureDistanceTypes_mapString, &stage.CodeProcedureDistanceType_stagedOrder, &stage.CodeProcedureDistanceTypeOrder)

	__gong__resetStageType(&stage.CodeProcedureFixRoleTypes, &stage.CodeProcedureFixRoleTypes_mapString, &stage.CodeProcedureFixRoleType_stagedOrder, &stage.CodeProcedureFixRoleTypeOrder)

	__gong__resetStageType(&stage.CodeProcedurePhaseTypes, &stage.CodeProcedurePhaseTypes_mapString, &stage.CodeProcedurePhaseType_stagedOrder, &stage.CodeProcedurePhaseTypeOrder)

	__gong__resetStageType(&stage.CodeRNPTypes, &stage.CodeRNPTypes_mapString, &stage.CodeRNPType_stagedOrder, &stage.CodeRNPTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_mapString, &stage.CodeRVSMPointRoleType_stagedOrder, &stage.CodeRVSMPointRoleTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString, &stage.CodeRVSMType_stagedOrder, &stage.CodeRVSMTypeOrder)

	__gong__resetStageType(&stage.CodeReferencePathIdentifierTypes, &stage.CodeReferencePathIdentifierTypes_mapString, &stage.CodeReferencePathIdentifierType_stagedOrder, &stage.CodeReferencePathIdentifierTypeOrder)

	__gong__resetStageType(&stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_mapString, &stage.CodeReferenceRoleType_stagedOrder, &stage.CodeReferenceRoleTypeOrder)

	__gong__resetStageType(&stage.CodeRelativePositionTypes, &stage.CodeRelativePositionTypes_mapString, &stage.CodeRelativePositionType_stagedOrder, &stage.CodeRelativePositionTypeOrder)

	__gong__resetStageType(&stage.CodeRouteIndicatorTypes, &stage.CodeRouteIndicatorTypes_mapString, &stage.CodeRouteIndicatorType_stagedOrder, &stage.CodeRouteIndicatorTypeOrder)

	__gong__resetStageType(&stage.CodeSafeAltitudeTypes, &stage.CodeSafeAltitudeTypes_mapString, &stage.CodeSafeAltitudeType_stagedOrder, &stage.CodeSafeAltitudeTypeOrder)

	__gong__resetStageType(&stage.CodeSegmentPathTypes, &stage.CodeSegmentPathTypes_mapString, &stage.CodeSegmentPathType_stagedOrder, &stage.CodeSegmentPathTypeOrder)

	__gong__resetStageType(&stage.CodeSegmentTerminationTypes, &stage.CodeSegmentTerminationTypes_mapString, &stage.CodeSegmentTerminationType_stagedOrder, &stage.CodeSegmentTerminationTypeOrder)

	__gong__resetStageType(&stage.CodeSideTypes, &stage.CodeSideTypes_mapString, &stage.CodeSideType_stagedOrder, &stage.CodeSideTypeOrder)

	__gong__resetStageType(&stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_mapString, &stage.CodeSpeedReferenceType_stagedOrder, &stage.CodeSpeedReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTAATypes, &stage.CodeTAATypes_mapString, &stage.CodeTAAType_stagedOrder, &stage.CodeTAATypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTrajectoryTypes, &stage.CodeTrajectoryTypes_mapString, &stage.CodeTrajectoryType_stagedOrder, &stage.CodeTrajectoryTypeOrder)

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

	__gong__resetStageType(&stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString, &stage.TextInstructionType_stagedOrder, &stage.TextInstructionTypeOrder)

	__gong__resetStageType(&stage.TextNameTypes, &stage.TextNameTypes_mapString, &stage.TextNameType_stagedOrder, &stage.TextNameTypeOrder)

	__gong__resetStageType(&stage.TextNoteTypes, &stage.TextNoteTypes_mapString, &stage.TextNoteType_stagedOrder, &stage.TextNoteTypeOrder)

	__gong__resetStageType(&stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString, &stage.TextPropertyNameType_stagedOrder, &stage.TextPropertyNameTypeOrder)

	__gong__resetStageType(&stage.TextSIDSTARDesignatorTypes, &stage.TextSIDSTARDesignatorTypes_mapString, &stage.TextSIDSTARDesignatorType_stagedOrder, &stage.TextSIDSTARDesignatorTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.ValAlarmLimitTypes, &stage.ValAlarmLimitTypes_mapString, &stage.ValAlarmLimitType_stagedOrder, &stage.ValAlarmLimitTypeOrder)

	__gong__resetStageType(&stage.ValAngleTypes, &stage.ValAngleTypes_mapString, &stage.ValAngleType_stagedOrder, &stage.ValAngleTypeOrder)

	__gong__resetStageType(&stage.ValBearingTypes, &stage.ValBearingTypes_mapString, &stage.ValBearingType_stagedOrder, &stage.ValBearingTypeOrder)

	__gong__resetStageType(&stage.ValChannelNumberTypes, &stage.ValChannelNumberTypes_mapString, &stage.ValChannelNumberType_stagedOrder, &stage.ValChannelNumberTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValHexTypes, &stage.ValHexTypes_mapString, &stage.ValHexType_stagedOrder, &stage.ValHexTypeOrder)

	__gong__resetStageType(&stage.ValSlopeTypes, &stage.ValSlopeTypes_mapString, &stage.ValSlopeType_stagedOrder, &stage.ValSlopeTypeOrder)

	__gong__resetStageType(&stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString, &stage.ValSpeedType_stagedOrder, &stage.ValSpeedTypeOrder)

	__gong__resetStageType(&stage.ValTemperatureTypes, &stage.ValTemperatureTypes_mapString, &stage.ValTemperatureType_stagedOrder, &stage.ValTemperatureTypeOrder)

	__gong__resetStageType(&stage.ValWeightTypes, &stage.ValWeightTypes_mapString, &stage.ValWeightType_stagedOrder, &stage.ValWeightTypeOrder)

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

func (CodeApproachEquipmentAdditionalType) GongGetAssociationName() any {
	return &CodeApproachEquipmentAdditionalType{
	}
}

func (CodeApproachGuidanceType) GongGetAssociationName() any {
	return &CodeApproachGuidanceType{
	}
}

func (CodeApproachPrefixType) GongGetAssociationName() any {
	return &CodeApproachPrefixType{
	}
}

func (CodeApproachType) GongGetAssociationName() any {
	return &CodeApproachType{
	}
}

func (CodeApprovalType) GongGetAssociationName() any {
	return &CodeApprovalType{
	}
}

func (CodeArcDirectionType) GongGetAssociationName() any {
	return &CodeArcDirectionType{
	}
}

func (CodeBearingType) GongGetAssociationName() any {
	return &CodeBearingType{
	}
}

func (CodeCommunicationModeType) GongGetAssociationName() any {
	return &CodeCommunicationModeType{
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

func (CodeDesignStandardType) GongGetAssociationName() any {
	return &CodeDesignStandardType{
	}
}

func (CodeDesignatedPointDesignatorType) GongGetAssociationName() any {
	return &CodeDesignatedPointDesignatorType{
	}
}

func (CodeDirectionReferenceType) GongGetAssociationName() any {
	return &CodeDirectionReferenceType{
	}
}

func (CodeDirectionTurnType) GongGetAssociationName() any {
	return &CodeDirectionTurnType{
	}
}

func (CodeEquipmentAntiCollisionType) GongGetAssociationName() any {
	return &CodeEquipmentAntiCollisionType{
	}
}

func (CodeEquipmentUnavailableType) GongGetAssociationName() any {
	return &CodeEquipmentUnavailableType{
	}
}

func (CodeFinalGuidanceType) GongGetAssociationName() any {
	return &CodeFinalGuidanceType{
	}
}

func (CodeFreeFlightType) GongGetAssociationName() any {
	return &CodeFreeFlightType{
	}
}

func (CodeHeightReferenceType) GongGetAssociationName() any {
	return &CodeHeightReferenceType{
	}
}

func (CodeHoldingUsageType) GongGetAssociationName() any {
	return &CodeHoldingUsageType{
	}
}

func (CodeHoldingUseType) GongGetAssociationName() any {
	return &CodeHoldingUseType{
	}
}

func (CodeICAOCountryType) GongGetAssociationName() any {
	return &CodeICAOCountryType{
	}
}

func (CodeMilitaryRoutePointType) GongGetAssociationName() any {
	return &CodeMilitaryRoutePointType{
	}
}

func (CodeMinimaFinalApproachPathType) GongGetAssociationName() any {
	return &CodeMinimaFinalApproachPathType{
	}
}

func (CodeMinimumAltitudeType) GongGetAssociationName() any {
	return &CodeMinimumAltitudeType{
	}
}

func (CodeMinimumHeightType) GongGetAssociationName() any {
	return &CodeMinimumHeightType{
	}
}

func (CodeMissedApproachType) GongGetAssociationName() any {
	return &CodeMissedApproachType{
	}
}

func (CodeNavigationAreaRestrictionType) GongGetAssociationName() any {
	return &CodeNavigationAreaRestrictionType{
	}
}

func (CodeNavigationAreaType) GongGetAssociationName() any {
	return &CodeNavigationAreaType{
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

func (CodeProcedureAvailabilityType) GongGetAssociationName() any {
	return &CodeProcedureAvailabilityType{
	}
}

func (CodeProcedureCodingStandardType) GongGetAssociationName() any {
	return &CodeProcedureCodingStandardType{
	}
}

func (CodeProcedureDistanceType) GongGetAssociationName() any {
	return &CodeProcedureDistanceType{
	}
}

func (CodeProcedureFixRoleType) GongGetAssociationName() any {
	return &CodeProcedureFixRoleType{
	}
}

func (CodeProcedurePhaseType) GongGetAssociationName() any {
	return &CodeProcedurePhaseType{
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

func (CodeReferencePathIdentifierType) GongGetAssociationName() any {
	return &CodeReferencePathIdentifierType{
	}
}

func (CodeReferenceRoleType) GongGetAssociationName() any {
	return &CodeReferenceRoleType{
	}
}

func (CodeRelativePositionType) GongGetAssociationName() any {
	return &CodeRelativePositionType{
	}
}

func (CodeRouteIndicatorType) GongGetAssociationName() any {
	return &CodeRouteIndicatorType{
	}
}

func (CodeSafeAltitudeType) GongGetAssociationName() any {
	return &CodeSafeAltitudeType{
	}
}

func (CodeSegmentPathType) GongGetAssociationName() any {
	return &CodeSegmentPathType{
	}
}

func (CodeSegmentTerminationType) GongGetAssociationName() any {
	return &CodeSegmentTerminationType{
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

func (CodeTAAType) GongGetAssociationName() any {
	return &CodeTAAType{
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

func (CodeTrajectoryType) GongGetAssociationName() any {
	return &CodeTrajectoryType{
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

func (TextPropertyNameType) GongGetAssociationName() any {
	return &TextPropertyNameType{
	}
}

func (TextSIDSTARDesignatorType) GongGetAssociationName() any {
	return &TextSIDSTARDesignatorType{
	}
}

func (TimeType) GongGetAssociationName() any {
	return &TimeType{
	}
}

func (ValAlarmLimitType) GongGetAssociationName() any {
	return &ValAlarmLimitType{
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

func (ValChannelNumberType) GongGetAssociationName() any {
	return &ValChannelNumberType{
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

func (ValHexType) GongGetAssociationName() any {
	return &ValHexType{
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

func (ValTemperatureType) GongGetAssociationName() any {
	return &ValTemperatureType{
	}
}

func (ValWeightType) GongGetAssociationName() any {
	return &ValWeightType{
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
	// reverse maps of direct associations of CodeApproachEquipmentAdditionalType
	case CodeApproachEquipmentAdditionalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachGuidanceType
	case CodeApproachGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachPrefixType
	case CodeApproachPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachType
	case CodeApproachType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApprovalType
	case CodeApprovalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArcDirectionType
	case CodeArcDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBearingType
	case CodeBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
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
	// reverse maps of direct associations of CodeDesignStandardType
	case CodeDesignStandardType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointDesignatorType
	case CodeDesignatedPointDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionReferenceType
	case CodeDirectionReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionTurnType
	case CodeDirectionTurnType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentUnavailableType
	case CodeEquipmentUnavailableType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFinalGuidanceType
	case CodeFinalGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFreeFlightType
	case CodeFreeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHeightReferenceType
	case CodeHeightReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingUsageType
	case CodeHoldingUsageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingUseType
	case CodeHoldingUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeICAOCountryType
	case CodeICAOCountryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryRoutePointType
	case CodeMilitaryRoutePointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimaFinalApproachPathType
	case CodeMinimaFinalApproachPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimumAltitudeType
	case CodeMinimumAltitudeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimumHeightType
	case CodeMinimumHeightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMissedApproachType
	case CodeMissedApproachType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationAreaRestrictionType
	case CodeNavigationAreaRestrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationAreaType
	case CodeNavigationAreaType:
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
	// reverse maps of direct associations of CodeProcedureAvailabilityType
	case CodeProcedureAvailabilityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureCodingStandardType
	case CodeProcedureCodingStandardType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureDistanceType
	case CodeProcedureDistanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureFixRoleType
	case CodeProcedureFixRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedurePhaseType
	case CodeProcedurePhaseType:
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
	// reverse maps of direct associations of CodeReferencePathIdentifierType
	case CodeReferencePathIdentifierType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReferenceRoleType
	case CodeReferenceRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRelativePositionType
	case CodeRelativePositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteIndicatorType
	case CodeRouteIndicatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSafeAltitudeType
	case CodeSafeAltitudeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSegmentPathType
	case CodeSegmentPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSegmentTerminationType
	case CodeSegmentTerminationType:
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
	// reverse maps of direct associations of CodeTAAType
	case CodeTAAType:
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
	// reverse maps of direct associations of CodeTrajectoryType
	case CodeTrajectoryType:
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
	// reverse maps of direct associations of TextPropertyNameType
	case TextPropertyNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextSIDSTARDesignatorType
	case TextSIDSTARDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValAlarmLimitType
	case ValAlarmLimitType:
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
	// reverse maps of direct associations of ValChannelNumberType
	case ValChannelNumberType:
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
	// reverse maps of direct associations of ValHexType
	case ValHexType:
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
	// reverse maps of direct associations of ValTemperatureType
	case ValTemperatureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValWeightType
	case ValWeightType:
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
	// reverse maps of direct associations of CodeApproachEquipmentAdditionalType
	case CodeApproachEquipmentAdditionalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachGuidanceType
	case CodeApproachGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachPrefixType
	case CodeApproachPrefixType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachType
	case CodeApproachType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApprovalType
	case CodeApprovalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArcDirectionType
	case CodeArcDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBearingType
	case CodeBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
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
	// reverse maps of direct associations of CodeDesignStandardType
	case CodeDesignStandardType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointDesignatorType
	case CodeDesignatedPointDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionReferenceType
	case CodeDirectionReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionTurnType
	case CodeDirectionTurnType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEquipmentUnavailableType
	case CodeEquipmentUnavailableType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFinalGuidanceType
	case CodeFinalGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFreeFlightType
	case CodeFreeFlightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHeightReferenceType
	case CodeHeightReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingUsageType
	case CodeHoldingUsageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingUseType
	case CodeHoldingUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeICAOCountryType
	case CodeICAOCountryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryRoutePointType
	case CodeMilitaryRoutePointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimaFinalApproachPathType
	case CodeMinimaFinalApproachPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimumAltitudeType
	case CodeMinimumAltitudeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMinimumHeightType
	case CodeMinimumHeightType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMissedApproachType
	case CodeMissedApproachType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationAreaRestrictionType
	case CodeNavigationAreaRestrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavigationAreaType
	case CodeNavigationAreaType:
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
	// reverse maps of direct associations of CodeProcedureAvailabilityType
	case CodeProcedureAvailabilityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureCodingStandardType
	case CodeProcedureCodingStandardType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureDistanceType
	case CodeProcedureDistanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedureFixRoleType
	case CodeProcedureFixRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProcedurePhaseType
	case CodeProcedurePhaseType:
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
	// reverse maps of direct associations of CodeReferencePathIdentifierType
	case CodeReferencePathIdentifierType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReferenceRoleType
	case CodeReferenceRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRelativePositionType
	case CodeRelativePositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRouteIndicatorType
	case CodeRouteIndicatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSafeAltitudeType
	case CodeSafeAltitudeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSegmentPathType
	case CodeSegmentPathType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSegmentTerminationType
	case CodeSegmentTerminationType:
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
	// reverse maps of direct associations of CodeTAAType
	case CodeTAAType:
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
	// reverse maps of direct associations of CodeTrajectoryType
	case CodeTrajectoryType:
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
	// reverse maps of direct associations of TextPropertyNameType
	case TextPropertyNameType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextSIDSTARDesignatorType
	case TextSIDSTARDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeType
	case TimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValAlarmLimitType
	case ValAlarmLimitType:
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
	// reverse maps of direct associations of ValChannelNumberType
	case ValChannelNumberType:
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
	// reverse maps of direct associations of ValHexType
	case ValHexType:
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
	// reverse maps of direct associations of ValTemperatureType
	case ValTemperatureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValWeightType
	case ValWeightType:
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

func (*CodeApproachEquipmentAdditionalType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachGuidanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachPrefixType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApprovalType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeArcDirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeBearingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationModeType) GongGetReverseFields() []GongReverseField {
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

func (*CodeDesignStandardType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDesignatedPointDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDirectionReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDirectionTurnType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeEquipmentAntiCollisionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeEquipmentUnavailableType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFinalGuidanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFreeFlightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeHeightReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeHoldingUsageType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeHoldingUseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeICAOCountryType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryRoutePointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMinimaFinalApproachPathType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMinimumAltitudeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMinimumHeightType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMissedApproachType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavigationAreaRestrictionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavigationAreaType) GongGetReverseFields() []GongReverseField {
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

func (*CodeProcedureAvailabilityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeProcedureCodingStandardType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeProcedureDistanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeProcedureFixRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeProcedurePhaseType) GongGetReverseFields() []GongReverseField {
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

func (*CodeReferencePathIdentifierType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeReferenceRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRelativePositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRouteIndicatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSafeAltitudeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSegmentPathType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSegmentTerminationType) GongGetReverseFields() []GongReverseField {
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

func (*CodeTAAType) GongGetReverseFields() []GongReverseField {
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

func (*CodeTrajectoryType) GongGetReverseFields() []GongReverseField {
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

func (*TextPropertyNameType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TextSIDSTARDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TimeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValAlarmLimitType) GongGetReverseFields() []GongReverseField {
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

func (*ValChannelNumberType) GongGetReverseFields() []GongReverseField {
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

func (*ValHexType) GongGetReverseFields() []GongReverseField {
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

func (*ValTemperatureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValWeightType) GongGetReverseFields() []GongReverseField {
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeapproachtype *CodeApproachType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeapprovaltype *CodeApprovalType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codearcdirectiontype *CodeArcDirectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codebearingtype *CodeBearingType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedesignstandardtype *CodeDesignStandardType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeheightreferencetype *CodeHeightReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeholdingusagetype *CodeHoldingUsageType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeholdingusetype *CodeHoldingUseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeicaocountrytype *CodeICAOCountryType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeminimumheighttype *CodeMinimumHeightType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codemissedapproachtype *CodeMissedApproachType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenavigationareatype *CodeNavigationAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (coderelativepositiontype *CodeRelativePositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codesegmentpathtype *CodeSegmentPathType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codetaatype *CodeTAAType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codetrajectorytype *CodeTrajectoryType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valalarmlimittype *ValAlarmLimitType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valchannelnumbertype *ValChannelNumberType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valhextype *ValHexType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valtemperaturetype *ValTemperatureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachequipmentadditionaltype.Name
	case "NilReason":
		res.valueString = codeapproachequipmentadditionaltype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachequipmentadditionaltype.EnclosedText
	}
	return
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachguidancetype.Name
	case "NilReason":
		res.valueString = codeapproachguidancetype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachguidancetype.EnclosedText
	}
	return
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachprefixtype.Name
	case "NilReason":
		res.valueString = codeapproachprefixtype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachprefixtype.EnclosedText
	}
	return
}

func (codeapproachtype *CodeApproachType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachtype.Name
	case "NilReason":
		res.valueString = codeapproachtype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachtype.EnclosedText
	}
	return
}

func (codeapprovaltype *CodeApprovalType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapprovaltype.Name
	case "NilReason":
		res.valueString = codeapprovaltype.NilReason
	case "EnclosedText":
		res.valueString = codeapprovaltype.EnclosedText
	}
	return
}

func (codearcdirectiontype *CodeArcDirectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codearcdirectiontype.Name
	case "NilReason":
		res.valueString = codearcdirectiontype.NilReason
	case "EnclosedText":
		res.valueString = codearcdirectiontype.EnclosedText
	}
	return
}

func (codebearingtype *CodeBearingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codebearingtype.Name
	case "NilReason":
		res.valueString = codebearingtype.NilReason
	case "EnclosedText":
		res.valueString = codebearingtype.EnclosedText
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

func (codedesignstandardtype *CodeDesignStandardType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedesignstandardtype.Name
	case "NilReason":
		res.valueString = codedesignstandardtype.NilReason
	case "EnclosedText":
		res.valueString = codedesignstandardtype.EnclosedText
	}
	return
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedesignatedpointdesignatortype.Name
	case "NilReason":
		res.valueString = codedesignatedpointdesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codedesignatedpointdesignatortype.EnclosedText
	}
	return
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedirectionreferencetype.Name
	case "NilReason":
		res.valueString = codedirectionreferencetype.NilReason
	case "EnclosedText":
		res.valueString = codedirectionreferencetype.EnclosedText
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

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeequipmentunavailabletype.Name
	case "NilReason":
		res.valueString = codeequipmentunavailabletype.NilReason
	case "EnclosedText":
		res.valueString = codeequipmentunavailabletype.EnclosedText
	}
	return
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefinalguidancetype.Name
	case "NilReason":
		res.valueString = codefinalguidancetype.NilReason
	case "EnclosedText":
		res.valueString = codefinalguidancetype.EnclosedText
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

func (codeheightreferencetype *CodeHeightReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeheightreferencetype.Name
	case "NilReason":
		res.valueString = codeheightreferencetype.NilReason
	case "EnclosedText":
		res.valueString = codeheightreferencetype.EnclosedText
	}
	return
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeholdingusagetype.Name
	case "NilReason":
		res.valueString = codeholdingusagetype.NilReason
	case "EnclosedText":
		res.valueString = codeholdingusagetype.EnclosedText
	}
	return
}

func (codeholdingusetype *CodeHoldingUseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeholdingusetype.Name
	case "NilReason":
		res.valueString = codeholdingusetype.NilReason
	case "EnclosedText":
		res.valueString = codeholdingusetype.EnclosedText
	}
	return
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeicaocountrytype.Name
	case "NilReason":
		res.valueString = codeicaocountrytype.NilReason
	case "EnclosedText":
		res.valueString = codeicaocountrytype.EnclosedText
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

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeminimafinalapproachpathtype.Name
	case "NilReason":
		res.valueString = codeminimafinalapproachpathtype.NilReason
	case "EnclosedText":
		res.valueString = codeminimafinalapproachpathtype.EnclosedText
	}
	return
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeminimumaltitudetype.Name
	case "NilReason":
		res.valueString = codeminimumaltitudetype.NilReason
	case "EnclosedText":
		res.valueString = codeminimumaltitudetype.EnclosedText
	}
	return
}

func (codeminimumheighttype *CodeMinimumHeightType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeminimumheighttype.Name
	case "NilReason":
		res.valueString = codeminimumheighttype.NilReason
	case "EnclosedText":
		res.valueString = codeminimumheighttype.EnclosedText
	}
	return
}

func (codemissedapproachtype *CodeMissedApproachType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemissedapproachtype.Name
	case "NilReason":
		res.valueString = codemissedapproachtype.NilReason
	case "EnclosedText":
		res.valueString = codemissedapproachtype.EnclosedText
	}
	return
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavigationarearestrictiontype.Name
	case "NilReason":
		res.valueString = codenavigationarearestrictiontype.NilReason
	case "EnclosedText":
		res.valueString = codenavigationarearestrictiontype.EnclosedText
	}
	return
}

func (codenavigationareatype *CodeNavigationAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavigationareatype.Name
	case "NilReason":
		res.valueString = codenavigationareatype.NilReason
	case "EnclosedText":
		res.valueString = codenavigationareatype.EnclosedText
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

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprocedureavailabilitytype.Name
	case "NilReason":
		res.valueString = codeprocedureavailabilitytype.NilReason
	case "EnclosedText":
		res.valueString = codeprocedureavailabilitytype.EnclosedText
	}
	return
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprocedurecodingstandardtype.Name
	case "NilReason":
		res.valueString = codeprocedurecodingstandardtype.NilReason
	case "EnclosedText":
		res.valueString = codeprocedurecodingstandardtype.EnclosedText
	}
	return
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeproceduredistancetype.Name
	case "NilReason":
		res.valueString = codeproceduredistancetype.NilReason
	case "EnclosedText":
		res.valueString = codeproceduredistancetype.EnclosedText
	}
	return
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprocedurefixroletype.Name
	case "NilReason":
		res.valueString = codeprocedurefixroletype.NilReason
	case "EnclosedText":
		res.valueString = codeprocedurefixroletype.EnclosedText
	}
	return
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprocedurephasetype.Name
	case "NilReason":
		res.valueString = codeprocedurephasetype.NilReason
	case "EnclosedText":
		res.valueString = codeprocedurephasetype.EnclosedText
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codereferencepathidentifiertype.Name
	case "NilReason":
		res.valueString = codereferencepathidentifiertype.NilReason
	case "EnclosedText":
		res.valueString = codereferencepathidentifiertype.EnclosedText
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

func (coderelativepositiontype *CodeRelativePositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderelativepositiontype.Name
	case "NilReason":
		res.valueString = coderelativepositiontype.NilReason
	case "EnclosedText":
		res.valueString = coderelativepositiontype.EnclosedText
	}
	return
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderouteindicatortype.Name
	case "NilReason":
		res.valueString = coderouteindicatortype.NilReason
	case "EnclosedText":
		res.valueString = coderouteindicatortype.EnclosedText
	}
	return
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesafealtitudetype.Name
	case "NilReason":
		res.valueString = codesafealtitudetype.NilReason
	case "EnclosedText":
		res.valueString = codesafealtitudetype.EnclosedText
	}
	return
}

func (codesegmentpathtype *CodeSegmentPathType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesegmentpathtype.Name
	case "NilReason":
		res.valueString = codesegmentpathtype.NilReason
	case "EnclosedText":
		res.valueString = codesegmentpathtype.EnclosedText
	}
	return
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesegmentterminationtype.Name
	case "NilReason":
		res.valueString = codesegmentterminationtype.NilReason
	case "EnclosedText":
		res.valueString = codesegmentterminationtype.EnclosedText
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

func (codetaatype *CodeTAAType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetaatype.Name
	case "NilReason":
		res.valueString = codetaatype.NilReason
	case "EnclosedText":
		res.valueString = codetaatype.EnclosedText
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

func (codetrajectorytype *CodeTrajectoryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetrajectorytype.Name
	case "NilReason":
		res.valueString = codetrajectorytype.NilReason
	case "EnclosedText":
		res.valueString = codetrajectorytype.EnclosedText
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

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textsidstardesignatortype.Name
	case "NilReason":
		res.valueString = textsidstardesignatortype.NilReason
	case "EnclosedText":
		res.valueString = textsidstardesignatortype.EnclosedText
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

func (valalarmlimittype *ValAlarmLimitType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valalarmlimittype.Name
	case "NilReason":
		res.valueString = valalarmlimittype.NilReason
	case "EnclosedText":
		res.valueString = valalarmlimittype.EnclosedText
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

func (valchannelnumbertype *ValChannelNumberType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valchannelnumbertype.Name
	case "NilReason":
		res.valueString = valchannelnumbertype.NilReason
	case "EnclosedText":
		res.valueString = valchannelnumbertype.EnclosedText
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

func (valhextype *ValHexType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valhextype.Name
	case "NilReason":
		res.valueString = valhextype.NilReason
	case "EnclosedText":
		res.valueString = valhextype.EnclosedText
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

func (valtemperaturetype *ValTemperatureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valtemperaturetype.Name
	case "Uom":
		res.valueString = valtemperaturetype.Uom
	case "NilReason":
		res.valueString = valtemperaturetype.NilReason
	case "EnclosedText":
		res.valueString = valtemperaturetype.EnclosedText
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetGongstructName() string {
	return "CodeApproachEquipmentAdditionalType"
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetGongstructName() string {
	return "CodeApproachGuidanceType"
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetGongstructName() string {
	return "CodeApproachPrefixType"
}

func (codeapproachtype *CodeApproachType) GongGetGongstructName() string {
	return "CodeApproachType"
}

func (codeapprovaltype *CodeApprovalType) GongGetGongstructName() string {
	return "CodeApprovalType"
}

func (codearcdirectiontype *CodeArcDirectionType) GongGetGongstructName() string {
	return "CodeArcDirectionType"
}

func (codebearingtype *CodeBearingType) GongGetGongstructName() string {
	return "CodeBearingType"
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetGongstructName() string {
	return "CodeCommunicationModeType"
}

func (codecoursetype *CodeCourseType) GongGetGongstructName() string {
	return "CodeCourseType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetGongstructName() string {
	return "CodeDesignStandardType"
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetGongstructName() string {
	return "CodeDesignatedPointDesignatorType"
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetGongstructName() string {
	return "CodeDirectionReferenceType"
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetGongstructName() string {
	return "CodeDirectionTurnType"
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetGongstructName() string {
	return "CodeEquipmentAntiCollisionType"
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetGongstructName() string {
	return "CodeEquipmentUnavailableType"
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetGongstructName() string {
	return "CodeFinalGuidanceType"
}

func (codefreeflighttype *CodeFreeFlightType) GongGetGongstructName() string {
	return "CodeFreeFlightType"
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetGongstructName() string {
	return "CodeHeightReferenceType"
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetGongstructName() string {
	return "CodeHoldingUsageType"
}

func (codeholdingusetype *CodeHoldingUseType) GongGetGongstructName() string {
	return "CodeHoldingUseType"
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetGongstructName() string {
	return "CodeICAOCountryType"
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetGongstructName() string {
	return "CodeMilitaryRoutePointType"
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetGongstructName() string {
	return "CodeMinimaFinalApproachPathType"
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetGongstructName() string {
	return "CodeMinimumAltitudeType"
}

func (codeminimumheighttype *CodeMinimumHeightType) GongGetGongstructName() string {
	return "CodeMinimumHeightType"
}

func (codemissedapproachtype *CodeMissedApproachType) GongGetGongstructName() string {
	return "CodeMissedApproachType"
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetGongstructName() string {
	return "CodeNavigationAreaRestrictionType"
}

func (codenavigationareatype *CodeNavigationAreaType) GongGetGongstructName() string {
	return "CodeNavigationAreaType"
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

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetGongstructName() string {
	return "CodeProcedureAvailabilityType"
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetGongstructName() string {
	return "CodeProcedureCodingStandardType"
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetGongstructName() string {
	return "CodeProcedureDistanceType"
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetGongstructName() string {
	return "CodeProcedureFixRoleType"
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetGongstructName() string {
	return "CodeProcedurePhaseType"
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetGongstructName() string {
	return "CodeReferencePathIdentifierType"
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetGongstructName() string {
	return "CodeReferenceRoleType"
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetGongstructName() string {
	return "CodeRelativePositionType"
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetGongstructName() string {
	return "CodeRouteIndicatorType"
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetGongstructName() string {
	return "CodeSafeAltitudeType"
}

func (codesegmentpathtype *CodeSegmentPathType) GongGetGongstructName() string {
	return "CodeSegmentPathType"
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetGongstructName() string {
	return "CodeSegmentTerminationType"
}

func (codesidetype *CodeSideType) GongGetGongstructName() string {
	return "CodeSideType"
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetGongstructName() string {
	return "CodeSpeedReferenceType"
}

func (codetaatype *CodeTAAType) GongGetGongstructName() string {
	return "CodeTAAType"
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

func (codetrajectorytype *CodeTrajectoryType) GongGetGongstructName() string {
	return "CodeTrajectoryType"
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

func (textinstructiontype *TextInstructionType) GongGetGongstructName() string {
	return "TextInstructionType"
}

func (textnametype *TextNameType) GongGetGongstructName() string {
	return "TextNameType"
}

func (textnotetype *TextNoteType) GongGetGongstructName() string {
	return "TextNoteType"
}

func (textpropertynametype *TextPropertyNameType) GongGetGongstructName() string {
	return "TextPropertyNameType"
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetGongstructName() string {
	return "TextSIDSTARDesignatorType"
}

func (timetype *TimeType) GongGetGongstructName() string {
	return "TimeType"
}

func (valalarmlimittype *ValAlarmLimitType) GongGetGongstructName() string {
	return "ValAlarmLimitType"
}

func (valangletype *ValAngleType) GongGetGongstructName() string {
	return "ValAngleType"
}

func (valbearingtype *ValBearingType) GongGetGongstructName() string {
	return "ValBearingType"
}

func (valchannelnumbertype *ValChannelNumberType) GongGetGongstructName() string {
	return "ValChannelNumberType"
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

func (valhextype *ValHexType) GongGetGongstructName() string {
	return "ValHexType"
}

func (valslopetype *ValSlopeType) GongGetGongstructName() string {
	return "ValSlopeType"
}

func (valspeedtype *ValSpeedType) GongGetGongstructName() string {
	return "ValSpeedType"
}

func (valtemperaturetype *ValTemperatureType) GongGetGongstructName() string {
	return "ValTemperatureType"
}

func (valweighttype *ValWeightType) GongGetGongstructName() string {
	return "ValWeightType"
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

	__gong__rebuildMapString(stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachEquipmentAdditionalTypes, &stage.CodeApproachEquipmentAdditionalTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachPrefixTypes, &stage.CodeApproachPrefixTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachTypes, &stage.CodeApproachTypes_mapString)

	__gong__rebuildMapString(stage.CodeApprovalTypes, &stage.CodeApprovalTypes_mapString)

	__gong__rebuildMapString(stage.CodeArcDirectionTypes, &stage.CodeArcDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeBearingTypes, &stage.CodeBearingTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString)

	__gong__rebuildMapString(stage.CodeCourseTypes, &stage.CodeCourseTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeDesignStandardTypes, &stage.CodeDesignStandardTypes_mapString)

	__gong__rebuildMapString(stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString)

	__gong__rebuildMapString(stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString)

	__gong__rebuildMapString(stage.CodeEquipmentUnavailableTypes, &stage.CodeEquipmentUnavailableTypes_mapString)

	__gong__rebuildMapString(stage.CodeFinalGuidanceTypes, &stage.CodeFinalGuidanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_mapString)

	__gong__rebuildMapString(stage.CodeHeightReferenceTypes, &stage.CodeHeightReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeHoldingUsageTypes, &stage.CodeHoldingUsageTypes_mapString)

	__gong__rebuildMapString(stage.CodeHoldingUseTypes, &stage.CodeHoldingUseTypes_mapString)

	__gong__rebuildMapString(stage.CodeICAOCountryTypes, &stage.CodeICAOCountryTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_mapString)

	__gong__rebuildMapString(stage.CodeMinimaFinalApproachPathTypes, &stage.CodeMinimaFinalApproachPathTypes_mapString)

	__gong__rebuildMapString(stage.CodeMinimumAltitudeTypes, &stage.CodeMinimumAltitudeTypes_mapString)

	__gong__rebuildMapString(stage.CodeMinimumHeightTypes, &stage.CodeMinimumHeightTypes_mapString)

	__gong__rebuildMapString(stage.CodeMissedApproachTypes, &stage.CodeMissedApproachTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationAreaRestrictionTypes, &stage.CodeNavigationAreaRestrictionTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationAreaTypes, &stage.CodeNavigationAreaTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_mapString)

	__gong__rebuildMapString(stage.CodeProcedureAvailabilityTypes, &stage.CodeProcedureAvailabilityTypes_mapString)

	__gong__rebuildMapString(stage.CodeProcedureCodingStandardTypes, &stage.CodeProcedureCodingStandardTypes_mapString)

	__gong__rebuildMapString(stage.CodeProcedureDistanceTypes, &stage.CodeProcedureDistanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeProcedureFixRoleTypes, &stage.CodeProcedureFixRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeProcedurePhaseTypes, &stage.CodeProcedurePhaseTypes_mapString)

	__gong__rebuildMapString(stage.CodeRNPTypes, &stage.CodeRNPTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString)

	__gong__rebuildMapString(stage.CodeReferencePathIdentifierTypes, &stage.CodeReferencePathIdentifierTypes_mapString)

	__gong__rebuildMapString(stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRelativePositionTypes, &stage.CodeRelativePositionTypes_mapString)

	__gong__rebuildMapString(stage.CodeRouteIndicatorTypes, &stage.CodeRouteIndicatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeSafeAltitudeTypes, &stage.CodeSafeAltitudeTypes_mapString)

	__gong__rebuildMapString(stage.CodeSegmentPathTypes, &stage.CodeSegmentPathTypes_mapString)

	__gong__rebuildMapString(stage.CodeSegmentTerminationTypes, &stage.CodeSegmentTerminationTypes_mapString)

	__gong__rebuildMapString(stage.CodeSideTypes, &stage.CodeSideTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTAATypes, &stage.CodeTAATypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTrajectoryTypes, &stage.CodeTrajectoryTypes_mapString)

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

	__gong__rebuildMapString(stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString)

	__gong__rebuildMapString(stage.TextNameTypes, &stage.TextNameTypes_mapString)

	__gong__rebuildMapString(stage.TextNoteTypes, &stage.TextNoteTypes_mapString)

	__gong__rebuildMapString(stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString)

	__gong__rebuildMapString(stage.TextSIDSTARDesignatorTypes, &stage.TextSIDSTARDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.ValAlarmLimitTypes, &stage.ValAlarmLimitTypes_mapString)

	__gong__rebuildMapString(stage.ValAngleTypes, &stage.ValAngleTypes_mapString)

	__gong__rebuildMapString(stage.ValBearingTypes, &stage.ValBearingTypes_mapString)

	__gong__rebuildMapString(stage.ValChannelNumberTypes, &stage.ValChannelNumberTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValHexTypes, &stage.ValHexTypes_mapString)

	__gong__rebuildMapString(stage.ValSlopeTypes, &stage.ValSlopeTypes_mapString)

	__gong__rebuildMapString(stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString)

	__gong__rebuildMapString(stage.ValTemperatureTypes, &stage.ValTemperatureTypes_mapString)

	__gong__rebuildMapString(stage.ValWeightTypes, &stage.ValWeightTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
