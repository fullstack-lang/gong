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
	CodeAirspacePointPositionTypes                map[*CodeAirspacePointPositionType]struct{}
	CodeAirspacePointPositionTypes_instance       map[*CodeAirspacePointPositionType]*CodeAirspacePointPositionType
	CodeAirspacePointPositionTypes_mapString      map[string]*CodeAirspacePointPositionType
	CodeAirspacePointPositionTypeOrder            uint
	CodeAirspacePointPositionType_stagedOrder     map[*CodeAirspacePointPositionType]uint
	CodeAirspacePointPositionType_orderStaged     map[uint]*CodeAirspacePointPositionType
	CodeAirspacePointPositionTypes_reference      map[*CodeAirspacePointPositionType]*CodeAirspacePointPositionType
	CodeAirspacePointPositionTypes_referenceOrder map[*CodeAirspacePointPositionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspacePointPositionTypeCreateCallback GongOnAfterCreateInterface[CodeAirspacePointPositionType]
	OnAfterCodeAirspacePointPositionTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspacePointPositionType]
	OnAfterCodeAirspacePointPositionTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspacePointPositionType]

	CodeAirspacePointRoleTypes                map[*CodeAirspacePointRoleType]struct{}
	CodeAirspacePointRoleTypes_instance       map[*CodeAirspacePointRoleType]*CodeAirspacePointRoleType
	CodeAirspacePointRoleTypes_mapString      map[string]*CodeAirspacePointRoleType
	CodeAirspacePointRoleTypeOrder            uint
	CodeAirspacePointRoleType_stagedOrder     map[*CodeAirspacePointRoleType]uint
	CodeAirspacePointRoleType_orderStaged     map[uint]*CodeAirspacePointRoleType
	CodeAirspacePointRoleTypes_reference      map[*CodeAirspacePointRoleType]*CodeAirspacePointRoleType
	CodeAirspacePointRoleTypes_referenceOrder map[*CodeAirspacePointRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspacePointRoleTypeCreateCallback GongOnAfterCreateInterface[CodeAirspacePointRoleType]
	OnAfterCodeAirspacePointRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspacePointRoleType]
	OnAfterCodeAirspacePointRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspacePointRoleType]

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

	CodeAuralMorseTypes                map[*CodeAuralMorseType]struct{}
	CodeAuralMorseTypes_instance       map[*CodeAuralMorseType]*CodeAuralMorseType
	CodeAuralMorseTypes_mapString      map[string]*CodeAuralMorseType
	CodeAuralMorseTypeOrder            uint
	CodeAuralMorseType_stagedOrder     map[*CodeAuralMorseType]uint
	CodeAuralMorseType_orderStaged     map[uint]*CodeAuralMorseType
	CodeAuralMorseTypes_reference      map[*CodeAuralMorseType]*CodeAuralMorseType
	CodeAuralMorseTypes_referenceOrder map[*CodeAuralMorseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAuralMorseTypeCreateCallback GongOnAfterCreateInterface[CodeAuralMorseType]
	OnAfterCodeAuralMorseTypeUpdateCallback GongOnAfterUpdateInterface[CodeAuralMorseType]
	OnAfterCodeAuralMorseTypeDeleteCallback GongOnAfterDeleteInterface[CodeAuralMorseType]

	CodeAuthorityRoleTypes                map[*CodeAuthorityRoleType]struct{}
	CodeAuthorityRoleTypes_instance       map[*CodeAuthorityRoleType]*CodeAuthorityRoleType
	CodeAuthorityRoleTypes_mapString      map[string]*CodeAuthorityRoleType
	CodeAuthorityRoleTypeOrder            uint
	CodeAuthorityRoleType_stagedOrder     map[*CodeAuthorityRoleType]uint
	CodeAuthorityRoleType_orderStaged     map[uint]*CodeAuthorityRoleType
	CodeAuthorityRoleTypes_reference      map[*CodeAuthorityRoleType]*CodeAuthorityRoleType
	CodeAuthorityRoleTypes_referenceOrder map[*CodeAuthorityRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAuthorityRoleTypeCreateCallback GongOnAfterCreateInterface[CodeAuthorityRoleType]
	OnAfterCodeAuthorityRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeAuthorityRoleType]
	OnAfterCodeAuthorityRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeAuthorityRoleType]

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

	CodeCheckpointCategoryTypes                map[*CodeCheckpointCategoryType]struct{}
	CodeCheckpointCategoryTypes_instance       map[*CodeCheckpointCategoryType]*CodeCheckpointCategoryType
	CodeCheckpointCategoryTypes_mapString      map[string]*CodeCheckpointCategoryType
	CodeCheckpointCategoryTypeOrder            uint
	CodeCheckpointCategoryType_stagedOrder     map[*CodeCheckpointCategoryType]uint
	CodeCheckpointCategoryType_orderStaged     map[uint]*CodeCheckpointCategoryType
	CodeCheckpointCategoryTypes_reference      map[*CodeCheckpointCategoryType]*CodeCheckpointCategoryType
	CodeCheckpointCategoryTypes_referenceOrder map[*CodeCheckpointCategoryType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCheckpointCategoryTypeCreateCallback GongOnAfterCreateInterface[CodeCheckpointCategoryType]
	OnAfterCodeCheckpointCategoryTypeUpdateCallback GongOnAfterUpdateInterface[CodeCheckpointCategoryType]
	OnAfterCodeCheckpointCategoryTypeDeleteCallback GongOnAfterDeleteInterface[CodeCheckpointCategoryType]

	CodeCourseQualityILSTypes                map[*CodeCourseQualityILSType]struct{}
	CodeCourseQualityILSTypes_instance       map[*CodeCourseQualityILSType]*CodeCourseQualityILSType
	CodeCourseQualityILSTypes_mapString      map[string]*CodeCourseQualityILSType
	CodeCourseQualityILSTypeOrder            uint
	CodeCourseQualityILSType_stagedOrder     map[*CodeCourseQualityILSType]uint
	CodeCourseQualityILSType_orderStaged     map[uint]*CodeCourseQualityILSType
	CodeCourseQualityILSTypes_reference      map[*CodeCourseQualityILSType]*CodeCourseQualityILSType
	CodeCourseQualityILSTypes_referenceOrder map[*CodeCourseQualityILSType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCourseQualityILSTypeCreateCallback GongOnAfterCreateInterface[CodeCourseQualityILSType]
	OnAfterCodeCourseQualityILSTypeUpdateCallback GongOnAfterUpdateInterface[CodeCourseQualityILSType]
	OnAfterCodeCourseQualityILSTypeDeleteCallback GongOnAfterDeleteInterface[CodeCourseQualityILSType]

	CodeDMEChannelTypes                map[*CodeDMEChannelType]struct{}
	CodeDMEChannelTypes_instance       map[*CodeDMEChannelType]*CodeDMEChannelType
	CodeDMEChannelTypes_mapString      map[string]*CodeDMEChannelType
	CodeDMEChannelTypeOrder            uint
	CodeDMEChannelType_stagedOrder     map[*CodeDMEChannelType]uint
	CodeDMEChannelType_orderStaged     map[uint]*CodeDMEChannelType
	CodeDMEChannelTypes_reference      map[*CodeDMEChannelType]*CodeDMEChannelType
	CodeDMEChannelTypes_referenceOrder map[*CodeDMEChannelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDMEChannelTypeCreateCallback GongOnAfterCreateInterface[CodeDMEChannelType]
	OnAfterCodeDMEChannelTypeUpdateCallback GongOnAfterUpdateInterface[CodeDMEChannelType]
	OnAfterCodeDMEChannelTypeDeleteCallback GongOnAfterDeleteInterface[CodeDMEChannelType]

	CodeDMETypes                map[*CodeDMEType]struct{}
	CodeDMETypes_instance       map[*CodeDMEType]*CodeDMEType
	CodeDMETypes_mapString      map[string]*CodeDMEType
	CodeDMETypeOrder            uint
	CodeDMEType_stagedOrder     map[*CodeDMEType]uint
	CodeDMEType_orderStaged     map[uint]*CodeDMEType
	CodeDMETypes_reference      map[*CodeDMEType]*CodeDMEType
	CodeDMETypes_referenceOrder map[*CodeDMEType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDMETypeCreateCallback GongOnAfterCreateInterface[CodeDMEType]
	OnAfterCodeDMETypeUpdateCallback GongOnAfterUpdateInterface[CodeDMEType]
	OnAfterCodeDMETypeDeleteCallback GongOnAfterDeleteInterface[CodeDMEType]

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

	CodeDesignatedPointTypes                map[*CodeDesignatedPointType]struct{}
	CodeDesignatedPointTypes_instance       map[*CodeDesignatedPointType]*CodeDesignatedPointType
	CodeDesignatedPointTypes_mapString      map[string]*CodeDesignatedPointType
	CodeDesignatedPointTypeOrder            uint
	CodeDesignatedPointType_stagedOrder     map[*CodeDesignatedPointType]uint
	CodeDesignatedPointType_orderStaged     map[uint]*CodeDesignatedPointType
	CodeDesignatedPointTypes_reference      map[*CodeDesignatedPointType]*CodeDesignatedPointType
	CodeDesignatedPointTypes_referenceOrder map[*CodeDesignatedPointType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDesignatedPointTypeCreateCallback GongOnAfterCreateInterface[CodeDesignatedPointType]
	OnAfterCodeDesignatedPointTypeUpdateCallback GongOnAfterUpdateInterface[CodeDesignatedPointType]
	OnAfterCodeDesignatedPointTypeDeleteCallback GongOnAfterDeleteInterface[CodeDesignatedPointType]

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

	CodeDistanceIndicationTypes                map[*CodeDistanceIndicationType]struct{}
	CodeDistanceIndicationTypes_instance       map[*CodeDistanceIndicationType]*CodeDistanceIndicationType
	CodeDistanceIndicationTypes_mapString      map[string]*CodeDistanceIndicationType
	CodeDistanceIndicationTypeOrder            uint
	CodeDistanceIndicationType_stagedOrder     map[*CodeDistanceIndicationType]uint
	CodeDistanceIndicationType_orderStaged     map[uint]*CodeDistanceIndicationType
	CodeDistanceIndicationTypes_reference      map[*CodeDistanceIndicationType]*CodeDistanceIndicationType
	CodeDistanceIndicationTypes_referenceOrder map[*CodeDistanceIndicationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDistanceIndicationTypeCreateCallback GongOnAfterCreateInterface[CodeDistanceIndicationType]
	OnAfterCodeDistanceIndicationTypeUpdateCallback GongOnAfterUpdateInterface[CodeDistanceIndicationType]
	OnAfterCodeDistanceIndicationTypeDeleteCallback GongOnAfterDeleteInterface[CodeDistanceIndicationType]

	CodeEmissionBandTypes                map[*CodeEmissionBandType]struct{}
	CodeEmissionBandTypes_instance       map[*CodeEmissionBandType]*CodeEmissionBandType
	CodeEmissionBandTypes_mapString      map[string]*CodeEmissionBandType
	CodeEmissionBandTypeOrder            uint
	CodeEmissionBandType_stagedOrder     map[*CodeEmissionBandType]uint
	CodeEmissionBandType_orderStaged     map[uint]*CodeEmissionBandType
	CodeEmissionBandTypes_reference      map[*CodeEmissionBandType]*CodeEmissionBandType
	CodeEmissionBandTypes_referenceOrder map[*CodeEmissionBandType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeEmissionBandTypeCreateCallback GongOnAfterCreateInterface[CodeEmissionBandType]
	OnAfterCodeEmissionBandTypeUpdateCallback GongOnAfterUpdateInterface[CodeEmissionBandType]
	OnAfterCodeEmissionBandTypeDeleteCallback GongOnAfterDeleteInterface[CodeEmissionBandType]

	CodeILSBackCourseTypes                map[*CodeILSBackCourseType]struct{}
	CodeILSBackCourseTypes_instance       map[*CodeILSBackCourseType]*CodeILSBackCourseType
	CodeILSBackCourseTypes_mapString      map[string]*CodeILSBackCourseType
	CodeILSBackCourseTypeOrder            uint
	CodeILSBackCourseType_stagedOrder     map[*CodeILSBackCourseType]uint
	CodeILSBackCourseType_orderStaged     map[uint]*CodeILSBackCourseType
	CodeILSBackCourseTypes_reference      map[*CodeILSBackCourseType]*CodeILSBackCourseType
	CodeILSBackCourseTypes_referenceOrder map[*CodeILSBackCourseType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeILSBackCourseTypeCreateCallback GongOnAfterCreateInterface[CodeILSBackCourseType]
	OnAfterCodeILSBackCourseTypeUpdateCallback GongOnAfterUpdateInterface[CodeILSBackCourseType]
	OnAfterCodeILSBackCourseTypeDeleteCallback GongOnAfterDeleteInterface[CodeILSBackCourseType]

	CodeIntegrityLevelILSTypes                map[*CodeIntegrityLevelILSType]struct{}
	CodeIntegrityLevelILSTypes_instance       map[*CodeIntegrityLevelILSType]*CodeIntegrityLevelILSType
	CodeIntegrityLevelILSTypes_mapString      map[string]*CodeIntegrityLevelILSType
	CodeIntegrityLevelILSTypeOrder            uint
	CodeIntegrityLevelILSType_stagedOrder     map[*CodeIntegrityLevelILSType]uint
	CodeIntegrityLevelILSType_orderStaged     map[uint]*CodeIntegrityLevelILSType
	CodeIntegrityLevelILSTypes_reference      map[*CodeIntegrityLevelILSType]*CodeIntegrityLevelILSType
	CodeIntegrityLevelILSTypes_referenceOrder map[*CodeIntegrityLevelILSType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeIntegrityLevelILSTypeCreateCallback GongOnAfterCreateInterface[CodeIntegrityLevelILSType]
	OnAfterCodeIntegrityLevelILSTypeUpdateCallback GongOnAfterUpdateInterface[CodeIntegrityLevelILSType]
	OnAfterCodeIntegrityLevelILSTypeDeleteCallback GongOnAfterDeleteInterface[CodeIntegrityLevelILSType]

	CodeMLSAzimuthTypes                map[*CodeMLSAzimuthType]struct{}
	CodeMLSAzimuthTypes_instance       map[*CodeMLSAzimuthType]*CodeMLSAzimuthType
	CodeMLSAzimuthTypes_mapString      map[string]*CodeMLSAzimuthType
	CodeMLSAzimuthTypeOrder            uint
	CodeMLSAzimuthType_stagedOrder     map[*CodeMLSAzimuthType]uint
	CodeMLSAzimuthType_orderStaged     map[uint]*CodeMLSAzimuthType
	CodeMLSAzimuthTypes_reference      map[*CodeMLSAzimuthType]*CodeMLSAzimuthType
	CodeMLSAzimuthTypes_referenceOrder map[*CodeMLSAzimuthType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMLSAzimuthTypeCreateCallback GongOnAfterCreateInterface[CodeMLSAzimuthType]
	OnAfterCodeMLSAzimuthTypeUpdateCallback GongOnAfterUpdateInterface[CodeMLSAzimuthType]
	OnAfterCodeMLSAzimuthTypeDeleteCallback GongOnAfterDeleteInterface[CodeMLSAzimuthType]

	CodeMLSChannelTypes                map[*CodeMLSChannelType]struct{}
	CodeMLSChannelTypes_instance       map[*CodeMLSChannelType]*CodeMLSChannelType
	CodeMLSChannelTypes_mapString      map[string]*CodeMLSChannelType
	CodeMLSChannelTypeOrder            uint
	CodeMLSChannelType_stagedOrder     map[*CodeMLSChannelType]uint
	CodeMLSChannelType_orderStaged     map[uint]*CodeMLSChannelType
	CodeMLSChannelTypes_reference      map[*CodeMLSChannelType]*CodeMLSChannelType
	CodeMLSChannelTypes_referenceOrder map[*CodeMLSChannelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMLSChannelTypeCreateCallback GongOnAfterCreateInterface[CodeMLSChannelType]
	OnAfterCodeMLSChannelTypeUpdateCallback GongOnAfterUpdateInterface[CodeMLSChannelType]
	OnAfterCodeMLSChannelTypeDeleteCallback GongOnAfterDeleteInterface[CodeMLSChannelType]

	CodeMarkerBeaconSignalTypes                map[*CodeMarkerBeaconSignalType]struct{}
	CodeMarkerBeaconSignalTypes_instance       map[*CodeMarkerBeaconSignalType]*CodeMarkerBeaconSignalType
	CodeMarkerBeaconSignalTypes_mapString      map[string]*CodeMarkerBeaconSignalType
	CodeMarkerBeaconSignalTypeOrder            uint
	CodeMarkerBeaconSignalType_stagedOrder     map[*CodeMarkerBeaconSignalType]uint
	CodeMarkerBeaconSignalType_orderStaged     map[uint]*CodeMarkerBeaconSignalType
	CodeMarkerBeaconSignalTypes_reference      map[*CodeMarkerBeaconSignalType]*CodeMarkerBeaconSignalType
	CodeMarkerBeaconSignalTypes_referenceOrder map[*CodeMarkerBeaconSignalType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMarkerBeaconSignalTypeCreateCallback GongOnAfterCreateInterface[CodeMarkerBeaconSignalType]
	OnAfterCodeMarkerBeaconSignalTypeUpdateCallback GongOnAfterUpdateInterface[CodeMarkerBeaconSignalType]
	OnAfterCodeMarkerBeaconSignalTypeDeleteCallback GongOnAfterDeleteInterface[CodeMarkerBeaconSignalType]

	CodeNDBUsageTypes                map[*CodeNDBUsageType]struct{}
	CodeNDBUsageTypes_instance       map[*CodeNDBUsageType]*CodeNDBUsageType
	CodeNDBUsageTypes_mapString      map[string]*CodeNDBUsageType
	CodeNDBUsageTypeOrder            uint
	CodeNDBUsageType_stagedOrder     map[*CodeNDBUsageType]uint
	CodeNDBUsageType_orderStaged     map[uint]*CodeNDBUsageType
	CodeNDBUsageTypes_reference      map[*CodeNDBUsageType]*CodeNDBUsageType
	CodeNDBUsageTypes_referenceOrder map[*CodeNDBUsageType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNDBUsageTypeCreateCallback GongOnAfterCreateInterface[CodeNDBUsageType]
	OnAfterCodeNDBUsageTypeUpdateCallback GongOnAfterUpdateInterface[CodeNDBUsageType]
	OnAfterCodeNDBUsageTypeDeleteCallback GongOnAfterDeleteInterface[CodeNDBUsageType]

	CodeNavaidDesignatorTypes                map[*CodeNavaidDesignatorType]struct{}
	CodeNavaidDesignatorTypes_instance       map[*CodeNavaidDesignatorType]*CodeNavaidDesignatorType
	CodeNavaidDesignatorTypes_mapString      map[string]*CodeNavaidDesignatorType
	CodeNavaidDesignatorTypeOrder            uint
	CodeNavaidDesignatorType_stagedOrder     map[*CodeNavaidDesignatorType]uint
	CodeNavaidDesignatorType_orderStaged     map[uint]*CodeNavaidDesignatorType
	CodeNavaidDesignatorTypes_reference      map[*CodeNavaidDesignatorType]*CodeNavaidDesignatorType
	CodeNavaidDesignatorTypes_referenceOrder map[*CodeNavaidDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavaidDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeNavaidDesignatorType]
	OnAfterCodeNavaidDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavaidDesignatorType]
	OnAfterCodeNavaidDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavaidDesignatorType]

	CodeNavaidPurposeTypes                map[*CodeNavaidPurposeType]struct{}
	CodeNavaidPurposeTypes_instance       map[*CodeNavaidPurposeType]*CodeNavaidPurposeType
	CodeNavaidPurposeTypes_mapString      map[string]*CodeNavaidPurposeType
	CodeNavaidPurposeTypeOrder            uint
	CodeNavaidPurposeType_stagedOrder     map[*CodeNavaidPurposeType]uint
	CodeNavaidPurposeType_orderStaged     map[uint]*CodeNavaidPurposeType
	CodeNavaidPurposeTypes_reference      map[*CodeNavaidPurposeType]*CodeNavaidPurposeType
	CodeNavaidPurposeTypes_referenceOrder map[*CodeNavaidPurposeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavaidPurposeTypeCreateCallback GongOnAfterCreateInterface[CodeNavaidPurposeType]
	OnAfterCodeNavaidPurposeTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavaidPurposeType]
	OnAfterCodeNavaidPurposeTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavaidPurposeType]

	CodeNavaidServiceTypes                map[*CodeNavaidServiceType]struct{}
	CodeNavaidServiceTypes_instance       map[*CodeNavaidServiceType]*CodeNavaidServiceType
	CodeNavaidServiceTypes_mapString      map[string]*CodeNavaidServiceType
	CodeNavaidServiceTypeOrder            uint
	CodeNavaidServiceType_stagedOrder     map[*CodeNavaidServiceType]uint
	CodeNavaidServiceType_orderStaged     map[uint]*CodeNavaidServiceType
	CodeNavaidServiceTypes_reference      map[*CodeNavaidServiceType]*CodeNavaidServiceType
	CodeNavaidServiceTypes_referenceOrder map[*CodeNavaidServiceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNavaidServiceTypeCreateCallback GongOnAfterCreateInterface[CodeNavaidServiceType]
	OnAfterCodeNavaidServiceTypeUpdateCallback GongOnAfterUpdateInterface[CodeNavaidServiceType]
	OnAfterCodeNavaidServiceTypeDeleteCallback GongOnAfterDeleteInterface[CodeNavaidServiceType]

	CodeNorthReferenceTypes                map[*CodeNorthReferenceType]struct{}
	CodeNorthReferenceTypes_instance       map[*CodeNorthReferenceType]*CodeNorthReferenceType
	CodeNorthReferenceTypes_mapString      map[string]*CodeNorthReferenceType
	CodeNorthReferenceTypeOrder            uint
	CodeNorthReferenceType_stagedOrder     map[*CodeNorthReferenceType]uint
	CodeNorthReferenceType_orderStaged     map[uint]*CodeNorthReferenceType
	CodeNorthReferenceTypes_reference      map[*CodeNorthReferenceType]*CodeNorthReferenceType
	CodeNorthReferenceTypes_referenceOrder map[*CodeNorthReferenceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNorthReferenceTypeCreateCallback GongOnAfterCreateInterface[CodeNorthReferenceType]
	OnAfterCodeNorthReferenceTypeUpdateCallback GongOnAfterUpdateInterface[CodeNorthReferenceType]
	OnAfterCodeNorthReferenceTypeDeleteCallback GongOnAfterDeleteInterface[CodeNorthReferenceType]

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

	CodePARTypes                map[*CodePARType]struct{}
	CodePARTypes_instance       map[*CodePARType]*CodePARType
	CodePARTypes_mapString      map[string]*CodePARType
	CodePARTypeOrder            uint
	CodePARType_stagedOrder     map[*CodePARType]uint
	CodePARType_orderStaged     map[uint]*CodePARType
	CodePARTypes_reference      map[*CodePARType]*CodePARType
	CodePARTypes_referenceOrder map[*CodePARType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePARTypeCreateCallback GongOnAfterCreateInterface[CodePARType]
	OnAfterCodePARTypeUpdateCallback GongOnAfterUpdateInterface[CodePARType]
	OnAfterCodePARTypeDeleteCallback GongOnAfterDeleteInterface[CodePARType]

	CodePositionInILSTypes                map[*CodePositionInILSType]struct{}
	CodePositionInILSTypes_instance       map[*CodePositionInILSType]*CodePositionInILSType
	CodePositionInILSTypes_mapString      map[string]*CodePositionInILSType
	CodePositionInILSTypeOrder            uint
	CodePositionInILSType_stagedOrder     map[*CodePositionInILSType]uint
	CodePositionInILSType_orderStaged     map[uint]*CodePositionInILSType
	CodePositionInILSTypes_reference      map[*CodePositionInILSType]*CodePositionInILSType
	CodePositionInILSTypes_referenceOrder map[*CodePositionInILSType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePositionInILSTypeCreateCallback GongOnAfterCreateInterface[CodePositionInILSType]
	OnAfterCodePositionInILSTypeUpdateCallback GongOnAfterUpdateInterface[CodePositionInILSType]
	OnAfterCodePositionInILSTypeDeleteCallback GongOnAfterDeleteInterface[CodePositionInILSType]

	CodePrimaryRadarTypes                map[*CodePrimaryRadarType]struct{}
	CodePrimaryRadarTypes_instance       map[*CodePrimaryRadarType]*CodePrimaryRadarType
	CodePrimaryRadarTypes_mapString      map[string]*CodePrimaryRadarType
	CodePrimaryRadarTypeOrder            uint
	CodePrimaryRadarType_stagedOrder     map[*CodePrimaryRadarType]uint
	CodePrimaryRadarType_orderStaged     map[uint]*CodePrimaryRadarType
	CodePrimaryRadarTypes_reference      map[*CodePrimaryRadarType]*CodePrimaryRadarType
	CodePrimaryRadarTypes_referenceOrder map[*CodePrimaryRadarType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePrimaryRadarTypeCreateCallback GongOnAfterCreateInterface[CodePrimaryRadarType]
	OnAfterCodePrimaryRadarTypeUpdateCallback GongOnAfterUpdateInterface[CodePrimaryRadarType]
	OnAfterCodePrimaryRadarTypeDeleteCallback GongOnAfterDeleteInterface[CodePrimaryRadarType]

	CodeRadarServiceTypes                map[*CodeRadarServiceType]struct{}
	CodeRadarServiceTypes_instance       map[*CodeRadarServiceType]*CodeRadarServiceType
	CodeRadarServiceTypes_mapString      map[string]*CodeRadarServiceType
	CodeRadarServiceTypeOrder            uint
	CodeRadarServiceType_stagedOrder     map[*CodeRadarServiceType]uint
	CodeRadarServiceType_orderStaged     map[uint]*CodeRadarServiceType
	CodeRadarServiceTypes_reference      map[*CodeRadarServiceType]*CodeRadarServiceType
	CodeRadarServiceTypes_referenceOrder map[*CodeRadarServiceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRadarServiceTypeCreateCallback GongOnAfterCreateInterface[CodeRadarServiceType]
	OnAfterCodeRadarServiceTypeUpdateCallback GongOnAfterUpdateInterface[CodeRadarServiceType]
	OnAfterCodeRadarServiceTypeDeleteCallback GongOnAfterDeleteInterface[CodeRadarServiceType]

	CodeRadioEmissionTypes                map[*CodeRadioEmissionType]struct{}
	CodeRadioEmissionTypes_instance       map[*CodeRadioEmissionType]*CodeRadioEmissionType
	CodeRadioEmissionTypes_mapString      map[string]*CodeRadioEmissionType
	CodeRadioEmissionTypeOrder            uint
	CodeRadioEmissionType_stagedOrder     map[*CodeRadioEmissionType]uint
	CodeRadioEmissionType_orderStaged     map[uint]*CodeRadioEmissionType
	CodeRadioEmissionTypes_reference      map[*CodeRadioEmissionType]*CodeRadioEmissionType
	CodeRadioEmissionTypes_referenceOrder map[*CodeRadioEmissionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRadioEmissionTypeCreateCallback GongOnAfterCreateInterface[CodeRadioEmissionType]
	OnAfterCodeRadioEmissionTypeUpdateCallback GongOnAfterUpdateInterface[CodeRadioEmissionType]
	OnAfterCodeRadioEmissionTypeDeleteCallback GongOnAfterDeleteInterface[CodeRadioEmissionType]

	CodeRadioSignalTypes                map[*CodeRadioSignalType]struct{}
	CodeRadioSignalTypes_instance       map[*CodeRadioSignalType]*CodeRadioSignalType
	CodeRadioSignalTypes_mapString      map[string]*CodeRadioSignalType
	CodeRadioSignalTypeOrder            uint
	CodeRadioSignalType_stagedOrder     map[*CodeRadioSignalType]uint
	CodeRadioSignalType_orderStaged     map[uint]*CodeRadioSignalType
	CodeRadioSignalTypes_reference      map[*CodeRadioSignalType]*CodeRadioSignalType
	CodeRadioSignalTypes_referenceOrder map[*CodeRadioSignalType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRadioSignalTypeCreateCallback GongOnAfterCreateInterface[CodeRadioSignalType]
	OnAfterCodeRadioSignalTypeUpdateCallback GongOnAfterUpdateInterface[CodeRadioSignalType]
	OnAfterCodeRadioSignalTypeDeleteCallback GongOnAfterDeleteInterface[CodeRadioSignalType]

	CodeReflectorTypes                map[*CodeReflectorType]struct{}
	CodeReflectorTypes_instance       map[*CodeReflectorType]*CodeReflectorType
	CodeReflectorTypes_mapString      map[string]*CodeReflectorType
	CodeReflectorTypeOrder            uint
	CodeReflectorType_stagedOrder     map[*CodeReflectorType]uint
	CodeReflectorType_orderStaged     map[uint]*CodeReflectorType
	CodeReflectorTypes_reference      map[*CodeReflectorType]*CodeReflectorType
	CodeReflectorTypes_referenceOrder map[*CodeReflectorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeReflectorTypeCreateCallback GongOnAfterCreateInterface[CodeReflectorType]
	OnAfterCodeReflectorTypeUpdateCallback GongOnAfterUpdateInterface[CodeReflectorType]
	OnAfterCodeReflectorTypeDeleteCallback GongOnAfterDeleteInterface[CodeReflectorType]

	CodeSignalPerformanceILSTypes                map[*CodeSignalPerformanceILSType]struct{}
	CodeSignalPerformanceILSTypes_instance       map[*CodeSignalPerformanceILSType]*CodeSignalPerformanceILSType
	CodeSignalPerformanceILSTypes_mapString      map[string]*CodeSignalPerformanceILSType
	CodeSignalPerformanceILSTypeOrder            uint
	CodeSignalPerformanceILSType_stagedOrder     map[*CodeSignalPerformanceILSType]uint
	CodeSignalPerformanceILSType_orderStaged     map[uint]*CodeSignalPerformanceILSType
	CodeSignalPerformanceILSTypes_reference      map[*CodeSignalPerformanceILSType]*CodeSignalPerformanceILSType
	CodeSignalPerformanceILSTypes_referenceOrder map[*CodeSignalPerformanceILSType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSignalPerformanceILSTypeCreateCallback GongOnAfterCreateInterface[CodeSignalPerformanceILSType]
	OnAfterCodeSignalPerformanceILSTypeUpdateCallback GongOnAfterUpdateInterface[CodeSignalPerformanceILSType]
	OnAfterCodeSignalPerformanceILSTypeDeleteCallback GongOnAfterDeleteInterface[CodeSignalPerformanceILSType]

	CodeSpecialNavigationChainDesignatorTypes                map[*CodeSpecialNavigationChainDesignatorType]struct{}
	CodeSpecialNavigationChainDesignatorTypes_instance       map[*CodeSpecialNavigationChainDesignatorType]*CodeSpecialNavigationChainDesignatorType
	CodeSpecialNavigationChainDesignatorTypes_mapString      map[string]*CodeSpecialNavigationChainDesignatorType
	CodeSpecialNavigationChainDesignatorTypeOrder            uint
	CodeSpecialNavigationChainDesignatorType_stagedOrder     map[*CodeSpecialNavigationChainDesignatorType]uint
	CodeSpecialNavigationChainDesignatorType_orderStaged     map[uint]*CodeSpecialNavigationChainDesignatorType
	CodeSpecialNavigationChainDesignatorTypes_reference      map[*CodeSpecialNavigationChainDesignatorType]*CodeSpecialNavigationChainDesignatorType
	CodeSpecialNavigationChainDesignatorTypes_referenceOrder map[*CodeSpecialNavigationChainDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSpecialNavigationChainDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeSpecialNavigationChainDesignatorType]
	OnAfterCodeSpecialNavigationChainDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeSpecialNavigationChainDesignatorType]
	OnAfterCodeSpecialNavigationChainDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeSpecialNavigationChainDesignatorType]

	CodeSpecialNavigationStationTypes                map[*CodeSpecialNavigationStationType]struct{}
	CodeSpecialNavigationStationTypes_instance       map[*CodeSpecialNavigationStationType]*CodeSpecialNavigationStationType
	CodeSpecialNavigationStationTypes_mapString      map[string]*CodeSpecialNavigationStationType
	CodeSpecialNavigationStationTypeOrder            uint
	CodeSpecialNavigationStationType_stagedOrder     map[*CodeSpecialNavigationStationType]uint
	CodeSpecialNavigationStationType_orderStaged     map[uint]*CodeSpecialNavigationStationType
	CodeSpecialNavigationStationTypes_reference      map[*CodeSpecialNavigationStationType]*CodeSpecialNavigationStationType
	CodeSpecialNavigationStationTypes_referenceOrder map[*CodeSpecialNavigationStationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSpecialNavigationStationTypeCreateCallback GongOnAfterCreateInterface[CodeSpecialNavigationStationType]
	OnAfterCodeSpecialNavigationStationTypeUpdateCallback GongOnAfterUpdateInterface[CodeSpecialNavigationStationType]
	OnAfterCodeSpecialNavigationStationTypeDeleteCallback GongOnAfterDeleteInterface[CodeSpecialNavigationStationType]

	CodeSpecialNavigationSystemTypes                map[*CodeSpecialNavigationSystemType]struct{}
	CodeSpecialNavigationSystemTypes_instance       map[*CodeSpecialNavigationSystemType]*CodeSpecialNavigationSystemType
	CodeSpecialNavigationSystemTypes_mapString      map[string]*CodeSpecialNavigationSystemType
	CodeSpecialNavigationSystemTypeOrder            uint
	CodeSpecialNavigationSystemType_stagedOrder     map[*CodeSpecialNavigationSystemType]uint
	CodeSpecialNavigationSystemType_orderStaged     map[uint]*CodeSpecialNavigationSystemType
	CodeSpecialNavigationSystemTypes_reference      map[*CodeSpecialNavigationSystemType]*CodeSpecialNavigationSystemType
	CodeSpecialNavigationSystemTypes_referenceOrder map[*CodeSpecialNavigationSystemType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSpecialNavigationSystemTypeCreateCallback GongOnAfterCreateInterface[CodeSpecialNavigationSystemType]
	OnAfterCodeSpecialNavigationSystemTypeUpdateCallback GongOnAfterUpdateInterface[CodeSpecialNavigationSystemType]
	OnAfterCodeSpecialNavigationSystemTypeDeleteCallback GongOnAfterDeleteInterface[CodeSpecialNavigationSystemType]

	CodeStandbyPowerTypes                map[*CodeStandbyPowerType]struct{}
	CodeStandbyPowerTypes_instance       map[*CodeStandbyPowerType]*CodeStandbyPowerType
	CodeStandbyPowerTypes_mapString      map[string]*CodeStandbyPowerType
	CodeStandbyPowerTypeOrder            uint
	CodeStandbyPowerType_stagedOrder     map[*CodeStandbyPowerType]uint
	CodeStandbyPowerType_orderStaged     map[uint]*CodeStandbyPowerType
	CodeStandbyPowerTypes_reference      map[*CodeStandbyPowerType]*CodeStandbyPowerType
	CodeStandbyPowerTypes_referenceOrder map[*CodeStandbyPowerType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStandbyPowerTypeCreateCallback GongOnAfterCreateInterface[CodeStandbyPowerType]
	OnAfterCodeStandbyPowerTypeUpdateCallback GongOnAfterUpdateInterface[CodeStandbyPowerType]
	OnAfterCodeStandbyPowerTypeDeleteCallback GongOnAfterDeleteInterface[CodeStandbyPowerType]

	CodeStatusNavaidTypes                map[*CodeStatusNavaidType]struct{}
	CodeStatusNavaidTypes_instance       map[*CodeStatusNavaidType]*CodeStatusNavaidType
	CodeStatusNavaidTypes_mapString      map[string]*CodeStatusNavaidType
	CodeStatusNavaidTypeOrder            uint
	CodeStatusNavaidType_stagedOrder     map[*CodeStatusNavaidType]uint
	CodeStatusNavaidType_orderStaged     map[uint]*CodeStatusNavaidType
	CodeStatusNavaidTypes_reference      map[*CodeStatusNavaidType]*CodeStatusNavaidType
	CodeStatusNavaidTypes_referenceOrder map[*CodeStatusNavaidType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusNavaidTypeCreateCallback GongOnAfterCreateInterface[CodeStatusNavaidType]
	OnAfterCodeStatusNavaidTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusNavaidType]
	OnAfterCodeStatusNavaidTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusNavaidType]

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

	CodeVORTypes                map[*CodeVORType]struct{}
	CodeVORTypes_instance       map[*CodeVORType]*CodeVORType
	CodeVORTypes_mapString      map[string]*CodeVORType
	CodeVORTypeOrder            uint
	CodeVORType_stagedOrder     map[*CodeVORType]uint
	CodeVORType_orderStaged     map[uint]*CodeVORType
	CodeVORTypes_reference      map[*CodeVORType]*CodeVORType
	CodeVORTypes_referenceOrder map[*CodeVORType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVORTypeCreateCallback GongOnAfterCreateInterface[CodeVORType]
	OnAfterCodeVORTypeUpdateCallback GongOnAfterUpdateInterface[CodeVORType]
	OnAfterCodeVORTypeDeleteCallback GongOnAfterDeleteInterface[CodeVORType]

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

	DateYearTypes                map[*DateYearType]struct{}
	DateYearTypes_instance       map[*DateYearType]*DateYearType
	DateYearTypes_mapString      map[string]*DateYearType
	DateYearTypeOrder            uint
	DateYearType_stagedOrder     map[*DateYearType]uint
	DateYearType_orderStaged     map[uint]*DateYearType
	DateYearTypes_reference      map[*DateYearType]*DateYearType
	DateYearTypes_referenceOrder map[*DateYearType]uint

	// insertion point for slice of pointers maps
	OnAfterDateYearTypeCreateCallback GongOnAfterCreateInterface[DateYearType]
	OnAfterDateYearTypeUpdateCallback GongOnAfterUpdateInterface[DateYearType]
	OnAfterDateYearTypeDeleteCallback GongOnAfterDeleteInterface[DateYearType]

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

	ValFrequencyTypes                map[*ValFrequencyType]struct{}
	ValFrequencyTypes_instance       map[*ValFrequencyType]*ValFrequencyType
	ValFrequencyTypes_mapString      map[string]*ValFrequencyType
	ValFrequencyTypeOrder            uint
	ValFrequencyType_stagedOrder     map[*ValFrequencyType]uint
	ValFrequencyType_orderStaged     map[uint]*ValFrequencyType
	ValFrequencyTypes_reference      map[*ValFrequencyType]*ValFrequencyType
	ValFrequencyTypes_referenceOrder map[*ValFrequencyType]uint

	// insertion point for slice of pointers maps
	OnAfterValFrequencyTypeCreateCallback GongOnAfterCreateInterface[ValFrequencyType]
	OnAfterValFrequencyTypeUpdateCallback GongOnAfterUpdateInterface[ValFrequencyType]
	OnAfterValFrequencyTypeDeleteCallback GongOnAfterDeleteInterface[ValFrequencyType]

	ValMagneticVariationTypes                map[*ValMagneticVariationType]struct{}
	ValMagneticVariationTypes_instance       map[*ValMagneticVariationType]*ValMagneticVariationType
	ValMagneticVariationTypes_mapString      map[string]*ValMagneticVariationType
	ValMagneticVariationTypeOrder            uint
	ValMagneticVariationType_stagedOrder     map[*ValMagneticVariationType]uint
	ValMagneticVariationType_orderStaged     map[uint]*ValMagneticVariationType
	ValMagneticVariationTypes_reference      map[*ValMagneticVariationType]*ValMagneticVariationType
	ValMagneticVariationTypes_referenceOrder map[*ValMagneticVariationType]uint

	// insertion point for slice of pointers maps
	OnAfterValMagneticVariationTypeCreateCallback GongOnAfterCreateInterface[ValMagneticVariationType]
	OnAfterValMagneticVariationTypeUpdateCallback GongOnAfterUpdateInterface[ValMagneticVariationType]
	OnAfterValMagneticVariationTypeDeleteCallback GongOnAfterDeleteInterface[ValMagneticVariationType]

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
	__gong__clearReferences(&stage.CodeAirspacePointPositionTypes_reference, &stage.CodeAirspacePointPositionTypes_instance, &stage.CodeAirspacePointPositionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspacePointRoleTypes_reference, &stage.CodeAirspacePointRoleTypes_instance, &stage.CodeAirspacePointRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_instance, &stage.CodeAltitudeUseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAuralMorseTypes_reference, &stage.CodeAuralMorseTypes_instance, &stage.CodeAuralMorseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAuthorityRoleTypes_reference, &stage.CodeAuthorityRoleTypes_instance, &stage.CodeAuthorityRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeBearingTypes_reference, &stage.CodeBearingTypes_instance, &stage.CodeBearingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCardinalDirectionTypes_reference, &stage.CodeCardinalDirectionTypes_instance, &stage.CodeCardinalDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCheckpointCategoryTypes_reference, &stage.CodeCheckpointCategoryTypes_instance, &stage.CodeCheckpointCategoryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCourseQualityILSTypes_reference, &stage.CodeCourseQualityILSTypes_instance, &stage.CodeCourseQualityILSTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDMEChannelTypes_reference, &stage.CodeDMEChannelTypes_instance, &stage.CodeDMEChannelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDMETypes_reference, &stage.CodeDMETypes_instance, &stage.CodeDMETypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDesignatedPointDesignatorTypes_reference, &stage.CodeDesignatedPointDesignatorTypes_instance, &stage.CodeDesignatedPointDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDesignatedPointTypes_reference, &stage.CodeDesignatedPointTypes_instance, &stage.CodeDesignatedPointTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionReferenceTypes_reference, &stage.CodeDirectionReferenceTypes_instance, &stage.CodeDirectionReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDistanceIndicationTypes_reference, &stage.CodeDistanceIndicationTypes_instance, &stage.CodeDistanceIndicationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEmissionBandTypes_reference, &stage.CodeEmissionBandTypes_instance, &stage.CodeEmissionBandTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeILSBackCourseTypes_reference, &stage.CodeILSBackCourseTypes_instance, &stage.CodeILSBackCourseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeIntegrityLevelILSTypes_reference, &stage.CodeIntegrityLevelILSTypes_instance, &stage.CodeIntegrityLevelILSTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMLSAzimuthTypes_reference, &stage.CodeMLSAzimuthTypes_instance, &stage.CodeMLSAzimuthTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMLSChannelTypes_reference, &stage.CodeMLSChannelTypes_instance, &stage.CodeMLSChannelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMarkerBeaconSignalTypes_reference, &stage.CodeMarkerBeaconSignalTypes_instance, &stage.CodeMarkerBeaconSignalTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNDBUsageTypes_reference, &stage.CodeNDBUsageTypes_instance, &stage.CodeNDBUsageTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavaidDesignatorTypes_reference, &stage.CodeNavaidDesignatorTypes_instance, &stage.CodeNavaidDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavaidPurposeTypes_reference, &stage.CodeNavaidPurposeTypes_instance, &stage.CodeNavaidPurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavaidServiceTypes_reference, &stage.CodeNavaidServiceTypes_instance, &stage.CodeNavaidServiceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNorthReferenceTypes_reference, &stage.CodeNorthReferenceTypes_instance, &stage.CodeNorthReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePARTypes_reference, &stage.CodePARTypes_instance, &stage.CodePARTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePositionInILSTypes_reference, &stage.CodePositionInILSTypes_instance, &stage.CodePositionInILSTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePrimaryRadarTypes_reference, &stage.CodePrimaryRadarTypes_instance, &stage.CodePrimaryRadarTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRadarServiceTypes_reference, &stage.CodeRadarServiceTypes_instance, &stage.CodeRadarServiceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRadioEmissionTypes_reference, &stage.CodeRadioEmissionTypes_instance, &stage.CodeRadioEmissionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRadioSignalTypes_reference, &stage.CodeRadioSignalTypes_instance, &stage.CodeRadioSignalTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeReflectorTypes_reference, &stage.CodeReflectorTypes_instance, &stage.CodeReflectorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSignalPerformanceILSTypes_reference, &stage.CodeSignalPerformanceILSTypes_instance, &stage.CodeSignalPerformanceILSTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpecialNavigationChainDesignatorTypes_reference, &stage.CodeSpecialNavigationChainDesignatorTypes_instance, &stage.CodeSpecialNavigationChainDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpecialNavigationStationTypes_reference, &stage.CodeSpecialNavigationStationTypes_instance, &stage.CodeSpecialNavigationStationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpecialNavigationSystemTypes_reference, &stage.CodeSpecialNavigationSystemTypes_instance, &stage.CodeSpecialNavigationSystemTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStandbyPowerTypes_reference, &stage.CodeStandbyPowerTypes_instance, &stage.CodeStandbyPowerTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusNavaidTypes_reference, &stage.CodeStatusNavaidTypes_instance, &stage.CodeStatusNavaidTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTACANChannelTypes_reference, &stage.CodeTACANChannelTypes_instance, &stage.CodeTACANChannelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_instance, &stage.CodeTelecomNetworkTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_instance, &stage.CodeTransponderTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVORTypes_reference, &stage.CodeVORTypes_instance, &stage.CodeVORTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_instance, &stage.CodeVerticalReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.DateYearTypes_reference, &stage.DateYearTypes_instance, &stage.DateYearTypes_referenceOrder)

	__gong__clearReferences(&stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_instance, &stage.NoSequenceTypes_referenceOrder)

	__gong__clearReferences(&stage.TextAddressTypes_reference, &stage.TextAddressTypes_instance, &stage.TextAddressTypes_referenceOrder)

	__gong__clearReferences(&stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_instance, &stage.TextDesignatorTypes_referenceOrder)

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

	__gong__clearReferences(&stage.ValFrequencyTypes_reference, &stage.ValFrequencyTypes_instance, &stage.ValFrequencyTypes_referenceOrder)

	__gong__clearReferences(&stage.ValMagneticVariationTypes_reference, &stage.ValMagneticVariationTypes_instance, &stage.ValMagneticVariationTypes_referenceOrder)

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
	stage.CodeAirspacePointPositionTypeOrder = __gong__recomputeOrder(stage.CodeAirspacePointPositionType_stagedOrder)

	stage.CodeAirspacePointRoleTypeOrder = __gong__recomputeOrder(stage.CodeAirspacePointRoleType_stagedOrder)

	stage.CodeAltitudeUseTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeUseType_stagedOrder)

	stage.CodeAuralMorseTypeOrder = __gong__recomputeOrder(stage.CodeAuralMorseType_stagedOrder)

	stage.CodeAuthorityRoleTypeOrder = __gong__recomputeOrder(stage.CodeAuthorityRoleType_stagedOrder)

	stage.CodeBearingTypeOrder = __gong__recomputeOrder(stage.CodeBearingType_stagedOrder)

	stage.CodeCardinalDirectionTypeOrder = __gong__recomputeOrder(stage.CodeCardinalDirectionType_stagedOrder)

	stage.CodeCheckpointCategoryTypeOrder = __gong__recomputeOrder(stage.CodeCheckpointCategoryType_stagedOrder)

	stage.CodeCourseQualityILSTypeOrder = __gong__recomputeOrder(stage.CodeCourseQualityILSType_stagedOrder)

	stage.CodeDMEChannelTypeOrder = __gong__recomputeOrder(stage.CodeDMEChannelType_stagedOrder)

	stage.CodeDMETypeOrder = __gong__recomputeOrder(stage.CodeDMEType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeDesignatedPointDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeDesignatedPointDesignatorType_stagedOrder)

	stage.CodeDesignatedPointTypeOrder = __gong__recomputeOrder(stage.CodeDesignatedPointType_stagedOrder)

	stage.CodeDirectionReferenceTypeOrder = __gong__recomputeOrder(stage.CodeDirectionReferenceType_stagedOrder)

	stage.CodeDistanceIndicationTypeOrder = __gong__recomputeOrder(stage.CodeDistanceIndicationType_stagedOrder)

	stage.CodeEmissionBandTypeOrder = __gong__recomputeOrder(stage.CodeEmissionBandType_stagedOrder)

	stage.CodeILSBackCourseTypeOrder = __gong__recomputeOrder(stage.CodeILSBackCourseType_stagedOrder)

	stage.CodeIntegrityLevelILSTypeOrder = __gong__recomputeOrder(stage.CodeIntegrityLevelILSType_stagedOrder)

	stage.CodeMLSAzimuthTypeOrder = __gong__recomputeOrder(stage.CodeMLSAzimuthType_stagedOrder)

	stage.CodeMLSChannelTypeOrder = __gong__recomputeOrder(stage.CodeMLSChannelType_stagedOrder)

	stage.CodeMarkerBeaconSignalTypeOrder = __gong__recomputeOrder(stage.CodeMarkerBeaconSignalType_stagedOrder)

	stage.CodeNDBUsageTypeOrder = __gong__recomputeOrder(stage.CodeNDBUsageType_stagedOrder)

	stage.CodeNavaidDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeNavaidDesignatorType_stagedOrder)

	stage.CodeNavaidPurposeTypeOrder = __gong__recomputeOrder(stage.CodeNavaidPurposeType_stagedOrder)

	stage.CodeNavaidServiceTypeOrder = __gong__recomputeOrder(stage.CodeNavaidServiceType_stagedOrder)

	stage.CodeNorthReferenceTypeOrder = __gong__recomputeOrder(stage.CodeNorthReferenceType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodePARTypeOrder = __gong__recomputeOrder(stage.CodePARType_stagedOrder)

	stage.CodePositionInILSTypeOrder = __gong__recomputeOrder(stage.CodePositionInILSType_stagedOrder)

	stage.CodePrimaryRadarTypeOrder = __gong__recomputeOrder(stage.CodePrimaryRadarType_stagedOrder)

	stage.CodeRadarServiceTypeOrder = __gong__recomputeOrder(stage.CodeRadarServiceType_stagedOrder)

	stage.CodeRadioEmissionTypeOrder = __gong__recomputeOrder(stage.CodeRadioEmissionType_stagedOrder)

	stage.CodeRadioSignalTypeOrder = __gong__recomputeOrder(stage.CodeRadioSignalType_stagedOrder)

	stage.CodeReflectorTypeOrder = __gong__recomputeOrder(stage.CodeReflectorType_stagedOrder)

	stage.CodeSignalPerformanceILSTypeOrder = __gong__recomputeOrder(stage.CodeSignalPerformanceILSType_stagedOrder)

	stage.CodeSpecialNavigationChainDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeSpecialNavigationChainDesignatorType_stagedOrder)

	stage.CodeSpecialNavigationStationTypeOrder = __gong__recomputeOrder(stage.CodeSpecialNavigationStationType_stagedOrder)

	stage.CodeSpecialNavigationSystemTypeOrder = __gong__recomputeOrder(stage.CodeSpecialNavigationSystemType_stagedOrder)

	stage.CodeStandbyPowerTypeOrder = __gong__recomputeOrder(stage.CodeStandbyPowerType_stagedOrder)

	stage.CodeStatusNavaidTypeOrder = __gong__recomputeOrder(stage.CodeStatusNavaidType_stagedOrder)

	stage.CodeTACANChannelTypeOrder = __gong__recomputeOrder(stage.CodeTACANChannelType_stagedOrder)

	stage.CodeTelecomNetworkTypeOrder = __gong__recomputeOrder(stage.CodeTelecomNetworkType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeTransponderTypeOrder = __gong__recomputeOrder(stage.CodeTransponderType_stagedOrder)

	stage.CodeVORTypeOrder = __gong__recomputeOrder(stage.CodeVORType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeVerticalReferenceTypeOrder = __gong__recomputeOrder(stage.CodeVerticalReferenceType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.DateYearTypeOrder = __gong__recomputeOrder(stage.DateYearType_stagedOrder)

	stage.NoSequenceTypeOrder = __gong__recomputeOrder(stage.NoSequenceType_stagedOrder)

	stage.TextAddressTypeOrder = __gong__recomputeOrder(stage.TextAddressType_stagedOrder)

	stage.TextDesignatorTypeOrder = __gong__recomputeOrder(stage.TextDesignatorType_stagedOrder)

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

	stage.ValFrequencyTypeOrder = __gong__recomputeOrder(stage.ValFrequencyType_stagedOrder)

	stage.ValMagneticVariationTypeOrder = __gong__recomputeOrder(stage.ValMagneticVariationType_stagedOrder)

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
func (*CodeAirspacePointPositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionType_stagedOrder)
}

func (*CodeAirspacePointPositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspacePointPositionType_orderStaged[order]
}

func (*CodeAirspacePointPositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspacePointPositionTypes_mapString
}

func (*CodeAirspacePointPositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspacePointPositionTypes
}

func (*CodeAirspacePointPositionType) GongNewInstance() any {
	return new(CodeAirspacePointPositionType)
}

func (*CodeAirspacePointRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleType_stagedOrder)
}

func (*CodeAirspacePointRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspacePointRoleType_orderStaged[order]
}

func (*CodeAirspacePointRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspacePointRoleTypes_mapString
}

func (*CodeAirspacePointRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspacePointRoleTypes
}

func (*CodeAirspacePointRoleType) GongNewInstance() any {
	return new(CodeAirspacePointRoleType)
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

func (*CodeAuralMorseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAuralMorseTypes, stage.CodeAuralMorseType_stagedOrder)
}

func (*CodeAuralMorseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAuralMorseType_orderStaged[order]
}

func (*CodeAuralMorseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAuralMorseTypes_mapString
}

func (*CodeAuralMorseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAuralMorseTypes
}

func (*CodeAuralMorseType) GongNewInstance() any {
	return new(CodeAuralMorseType)
}

func (*CodeAuthorityRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleType_stagedOrder)
}

func (*CodeAuthorityRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAuthorityRoleType_orderStaged[order]
}

func (*CodeAuthorityRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAuthorityRoleTypes_mapString
}

func (*CodeAuthorityRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAuthorityRoleTypes
}

func (*CodeAuthorityRoleType) GongNewInstance() any {
	return new(CodeAuthorityRoleType)
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

func (*CodeCheckpointCategoryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryType_stagedOrder)
}

func (*CodeCheckpointCategoryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCheckpointCategoryType_orderStaged[order]
}

func (*CodeCheckpointCategoryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCheckpointCategoryTypes_mapString
}

func (*CodeCheckpointCategoryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCheckpointCategoryTypes
}

func (*CodeCheckpointCategoryType) GongNewInstance() any {
	return new(CodeCheckpointCategoryType)
}

func (*CodeCourseQualityILSType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSType_stagedOrder)
}

func (*CodeCourseQualityILSType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCourseQualityILSType_orderStaged[order]
}

func (*CodeCourseQualityILSType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCourseQualityILSTypes_mapString
}

func (*CodeCourseQualityILSType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCourseQualityILSTypes
}

func (*CodeCourseQualityILSType) GongNewInstance() any {
	return new(CodeCourseQualityILSType)
}

func (*CodeDMEChannelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDMEChannelTypes, stage.CodeDMEChannelType_stagedOrder)
}

func (*CodeDMEChannelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDMEChannelType_orderStaged[order]
}

func (*CodeDMEChannelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDMEChannelTypes_mapString
}

func (*CodeDMEChannelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDMEChannelTypes
}

func (*CodeDMEChannelType) GongNewInstance() any {
	return new(CodeDMEChannelType)
}

func (*CodeDMEType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDMETypes, stage.CodeDMEType_stagedOrder)
}

func (*CodeDMEType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDMEType_orderStaged[order]
}

func (*CodeDMEType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDMETypes_mapString
}

func (*CodeDMEType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDMETypes
}

func (*CodeDMEType) GongNewInstance() any {
	return new(CodeDMEType)
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

func (*CodeDesignatedPointType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointType_stagedOrder)
}

func (*CodeDesignatedPointType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDesignatedPointType_orderStaged[order]
}

func (*CodeDesignatedPointType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDesignatedPointTypes_mapString
}

func (*CodeDesignatedPointType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDesignatedPointTypes
}

func (*CodeDesignatedPointType) GongNewInstance() any {
	return new(CodeDesignatedPointType)
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

func (*CodeDistanceIndicationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationType_stagedOrder)
}

func (*CodeDistanceIndicationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDistanceIndicationType_orderStaged[order]
}

func (*CodeDistanceIndicationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDistanceIndicationTypes_mapString
}

func (*CodeDistanceIndicationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDistanceIndicationTypes
}

func (*CodeDistanceIndicationType) GongNewInstance() any {
	return new(CodeDistanceIndicationType)
}

func (*CodeEmissionBandType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeEmissionBandTypes, stage.CodeEmissionBandType_stagedOrder)
}

func (*CodeEmissionBandType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeEmissionBandType_orderStaged[order]
}

func (*CodeEmissionBandType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeEmissionBandTypes_mapString
}

func (*CodeEmissionBandType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeEmissionBandTypes
}

func (*CodeEmissionBandType) GongNewInstance() any {
	return new(CodeEmissionBandType)
}

func (*CodeILSBackCourseType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseType_stagedOrder)
}

func (*CodeILSBackCourseType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeILSBackCourseType_orderStaged[order]
}

func (*CodeILSBackCourseType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeILSBackCourseTypes_mapString
}

func (*CodeILSBackCourseType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeILSBackCourseTypes
}

func (*CodeILSBackCourseType) GongNewInstance() any {
	return new(CodeILSBackCourseType)
}

func (*CodeIntegrityLevelILSType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSType_stagedOrder)
}

func (*CodeIntegrityLevelILSType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeIntegrityLevelILSType_orderStaged[order]
}

func (*CodeIntegrityLevelILSType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeIntegrityLevelILSTypes_mapString
}

func (*CodeIntegrityLevelILSType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeIntegrityLevelILSTypes
}

func (*CodeIntegrityLevelILSType) GongNewInstance() any {
	return new(CodeIntegrityLevelILSType)
}

func (*CodeMLSAzimuthType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthType_stagedOrder)
}

func (*CodeMLSAzimuthType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMLSAzimuthType_orderStaged[order]
}

func (*CodeMLSAzimuthType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMLSAzimuthTypes_mapString
}

func (*CodeMLSAzimuthType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMLSAzimuthTypes
}

func (*CodeMLSAzimuthType) GongNewInstance() any {
	return new(CodeMLSAzimuthType)
}

func (*CodeMLSChannelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMLSChannelTypes, stage.CodeMLSChannelType_stagedOrder)
}

func (*CodeMLSChannelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMLSChannelType_orderStaged[order]
}

func (*CodeMLSChannelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMLSChannelTypes_mapString
}

func (*CodeMLSChannelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMLSChannelTypes
}

func (*CodeMLSChannelType) GongNewInstance() any {
	return new(CodeMLSChannelType)
}

func (*CodeMarkerBeaconSignalType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalType_stagedOrder)
}

func (*CodeMarkerBeaconSignalType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMarkerBeaconSignalType_orderStaged[order]
}

func (*CodeMarkerBeaconSignalType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMarkerBeaconSignalTypes_mapString
}

func (*CodeMarkerBeaconSignalType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMarkerBeaconSignalTypes
}

func (*CodeMarkerBeaconSignalType) GongNewInstance() any {
	return new(CodeMarkerBeaconSignalType)
}

func (*CodeNDBUsageType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNDBUsageTypes, stage.CodeNDBUsageType_stagedOrder)
}

func (*CodeNDBUsageType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNDBUsageType_orderStaged[order]
}

func (*CodeNDBUsageType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNDBUsageTypes_mapString
}

func (*CodeNDBUsageType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNDBUsageTypes
}

func (*CodeNDBUsageType) GongNewInstance() any {
	return new(CodeNDBUsageType)
}

func (*CodeNavaidDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorType_stagedOrder)
}

func (*CodeNavaidDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavaidDesignatorType_orderStaged[order]
}

func (*CodeNavaidDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavaidDesignatorTypes_mapString
}

func (*CodeNavaidDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavaidDesignatorTypes
}

func (*CodeNavaidDesignatorType) GongNewInstance() any {
	return new(CodeNavaidDesignatorType)
}

func (*CodeNavaidPurposeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeType_stagedOrder)
}

func (*CodeNavaidPurposeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavaidPurposeType_orderStaged[order]
}

func (*CodeNavaidPurposeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavaidPurposeTypes_mapString
}

func (*CodeNavaidPurposeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavaidPurposeTypes
}

func (*CodeNavaidPurposeType) GongNewInstance() any {
	return new(CodeNavaidPurposeType)
}

func (*CodeNavaidServiceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceType_stagedOrder)
}

func (*CodeNavaidServiceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNavaidServiceType_orderStaged[order]
}

func (*CodeNavaidServiceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNavaidServiceTypes_mapString
}

func (*CodeNavaidServiceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNavaidServiceTypes
}

func (*CodeNavaidServiceType) GongNewInstance() any {
	return new(CodeNavaidServiceType)
}

func (*CodeNorthReferenceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceType_stagedOrder)
}

func (*CodeNorthReferenceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNorthReferenceType_orderStaged[order]
}

func (*CodeNorthReferenceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNorthReferenceTypes_mapString
}

func (*CodeNorthReferenceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNorthReferenceTypes
}

func (*CodeNorthReferenceType) GongNewInstance() any {
	return new(CodeNorthReferenceType)
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

func (*CodePARType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePARTypes, stage.CodePARType_stagedOrder)
}

func (*CodePARType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePARType_orderStaged[order]
}

func (*CodePARType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePARTypes_mapString
}

func (*CodePARType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePARTypes
}

func (*CodePARType) GongNewInstance() any {
	return new(CodePARType)
}

func (*CodePositionInILSType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePositionInILSTypes, stage.CodePositionInILSType_stagedOrder)
}

func (*CodePositionInILSType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePositionInILSType_orderStaged[order]
}

func (*CodePositionInILSType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePositionInILSTypes_mapString
}

func (*CodePositionInILSType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePositionInILSTypes
}

func (*CodePositionInILSType) GongNewInstance() any {
	return new(CodePositionInILSType)
}

func (*CodePrimaryRadarType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarType_stagedOrder)
}

func (*CodePrimaryRadarType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePrimaryRadarType_orderStaged[order]
}

func (*CodePrimaryRadarType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePrimaryRadarTypes_mapString
}

func (*CodePrimaryRadarType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePrimaryRadarTypes
}

func (*CodePrimaryRadarType) GongNewInstance() any {
	return new(CodePrimaryRadarType)
}

func (*CodeRadarServiceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRadarServiceTypes, stage.CodeRadarServiceType_stagedOrder)
}

func (*CodeRadarServiceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRadarServiceType_orderStaged[order]
}

func (*CodeRadarServiceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRadarServiceTypes_mapString
}

func (*CodeRadarServiceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRadarServiceTypes
}

func (*CodeRadarServiceType) GongNewInstance() any {
	return new(CodeRadarServiceType)
}

func (*CodeRadioEmissionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionType_stagedOrder)
}

func (*CodeRadioEmissionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRadioEmissionType_orderStaged[order]
}

func (*CodeRadioEmissionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRadioEmissionTypes_mapString
}

func (*CodeRadioEmissionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRadioEmissionTypes
}

func (*CodeRadioEmissionType) GongNewInstance() any {
	return new(CodeRadioEmissionType)
}

func (*CodeRadioSignalType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRadioSignalTypes, stage.CodeRadioSignalType_stagedOrder)
}

func (*CodeRadioSignalType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRadioSignalType_orderStaged[order]
}

func (*CodeRadioSignalType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRadioSignalTypes_mapString
}

func (*CodeRadioSignalType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRadioSignalTypes
}

func (*CodeRadioSignalType) GongNewInstance() any {
	return new(CodeRadioSignalType)
}

func (*CodeReflectorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeReflectorTypes, stage.CodeReflectorType_stagedOrder)
}

func (*CodeReflectorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeReflectorType_orderStaged[order]
}

func (*CodeReflectorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeReflectorTypes_mapString
}

func (*CodeReflectorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeReflectorTypes
}

func (*CodeReflectorType) GongNewInstance() any {
	return new(CodeReflectorType)
}

func (*CodeSignalPerformanceILSType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSType_stagedOrder)
}

func (*CodeSignalPerformanceILSType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSignalPerformanceILSType_orderStaged[order]
}

func (*CodeSignalPerformanceILSType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSignalPerformanceILSTypes_mapString
}

func (*CodeSignalPerformanceILSType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSignalPerformanceILSTypes
}

func (*CodeSignalPerformanceILSType) GongNewInstance() any {
	return new(CodeSignalPerformanceILSType)
}

func (*CodeSpecialNavigationChainDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorType_stagedOrder)
}

func (*CodeSpecialNavigationChainDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSpecialNavigationChainDesignatorType_orderStaged[order]
}

func (*CodeSpecialNavigationChainDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSpecialNavigationChainDesignatorTypes_mapString
}

func (*CodeSpecialNavigationChainDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSpecialNavigationChainDesignatorTypes
}

func (*CodeSpecialNavigationChainDesignatorType) GongNewInstance() any {
	return new(CodeSpecialNavigationChainDesignatorType)
}

func (*CodeSpecialNavigationStationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationType_stagedOrder)
}

func (*CodeSpecialNavigationStationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSpecialNavigationStationType_orderStaged[order]
}

func (*CodeSpecialNavigationStationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSpecialNavigationStationTypes_mapString
}

func (*CodeSpecialNavigationStationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSpecialNavigationStationTypes
}

func (*CodeSpecialNavigationStationType) GongNewInstance() any {
	return new(CodeSpecialNavigationStationType)
}

func (*CodeSpecialNavigationSystemType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemType_stagedOrder)
}

func (*CodeSpecialNavigationSystemType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSpecialNavigationSystemType_orderStaged[order]
}

func (*CodeSpecialNavigationSystemType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSpecialNavigationSystemTypes_mapString
}

func (*CodeSpecialNavigationSystemType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSpecialNavigationSystemTypes
}

func (*CodeSpecialNavigationSystemType) GongNewInstance() any {
	return new(CodeSpecialNavigationSystemType)
}

func (*CodeStandbyPowerType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerType_stagedOrder)
}

func (*CodeStandbyPowerType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStandbyPowerType_orderStaged[order]
}

func (*CodeStandbyPowerType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStandbyPowerTypes_mapString
}

func (*CodeStandbyPowerType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStandbyPowerTypes
}

func (*CodeStandbyPowerType) GongNewInstance() any {
	return new(CodeStandbyPowerType)
}

func (*CodeStatusNavaidType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidType_stagedOrder)
}

func (*CodeStatusNavaidType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusNavaidType_orderStaged[order]
}

func (*CodeStatusNavaidType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusNavaidTypes_mapString
}

func (*CodeStatusNavaidType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusNavaidTypes
}

func (*CodeStatusNavaidType) GongNewInstance() any {
	return new(CodeStatusNavaidType)
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

func (*CodeVORType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVORTypes, stage.CodeVORType_stagedOrder)
}

func (*CodeVORType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVORType_orderStaged[order]
}

func (*CodeVORType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVORTypes_mapString
}

func (*CodeVORType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVORTypes
}

func (*CodeVORType) GongNewInstance() any {
	return new(CodeVORType)
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

func (*DateYearType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DateYearTypes, stage.DateYearType_stagedOrder)
}

func (*DateYearType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DateYearType_orderStaged[order]
}

func (*DateYearType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DateYearTypes_mapString
}

func (*DateYearType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DateYearTypes
}

func (*DateYearType) GongNewInstance() any {
	return new(DateYearType)
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

func (*ValFrequencyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValFrequencyTypes, stage.ValFrequencyType_stagedOrder)
}

func (*ValFrequencyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValFrequencyType_orderStaged[order]
}

func (*ValFrequencyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValFrequencyTypes_mapString
}

func (*ValFrequencyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValFrequencyTypes
}

func (*ValFrequencyType) GongNewInstance() any {
	return new(ValFrequencyType)
}

func (*ValMagneticVariationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValMagneticVariationTypes, stage.ValMagneticVariationType_stagedOrder)
}

func (*ValMagneticVariationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValMagneticVariationType_orderStaged[order]
}

func (*ValMagneticVariationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValMagneticVariationTypes_mapString
}

func (*ValMagneticVariationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValMagneticVariationTypes
}

func (*ValMagneticVariationType) GongNewInstance() any {
	return new(ValMagneticVariationType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
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
		CodeAirspacePointPositionTypes:           make(map[*CodeAirspacePointPositionType]struct{}),
		CodeAirspacePointPositionTypes_mapString: make(map[string]*CodeAirspacePointPositionType),

		CodeAirspacePointRoleTypes:           make(map[*CodeAirspacePointRoleType]struct{}),
		CodeAirspacePointRoleTypes_mapString: make(map[string]*CodeAirspacePointRoleType),

		CodeAltitudeUseTypes:           make(map[*CodeAltitudeUseType]struct{}),
		CodeAltitudeUseTypes_mapString: make(map[string]*CodeAltitudeUseType),

		CodeAuralMorseTypes:           make(map[*CodeAuralMorseType]struct{}),
		CodeAuralMorseTypes_mapString: make(map[string]*CodeAuralMorseType),

		CodeAuthorityRoleTypes:           make(map[*CodeAuthorityRoleType]struct{}),
		CodeAuthorityRoleTypes_mapString: make(map[string]*CodeAuthorityRoleType),

		CodeBearingTypes:           make(map[*CodeBearingType]struct{}),
		CodeBearingTypes_mapString: make(map[string]*CodeBearingType),

		CodeCardinalDirectionTypes:           make(map[*CodeCardinalDirectionType]struct{}),
		CodeCardinalDirectionTypes_mapString: make(map[string]*CodeCardinalDirectionType),

		CodeCheckpointCategoryTypes:           make(map[*CodeCheckpointCategoryType]struct{}),
		CodeCheckpointCategoryTypes_mapString: make(map[string]*CodeCheckpointCategoryType),

		CodeCourseQualityILSTypes:           make(map[*CodeCourseQualityILSType]struct{}),
		CodeCourseQualityILSTypes_mapString: make(map[string]*CodeCourseQualityILSType),

		CodeDMEChannelTypes:           make(map[*CodeDMEChannelType]struct{}),
		CodeDMEChannelTypes_mapString: make(map[string]*CodeDMEChannelType),

		CodeDMETypes:           make(map[*CodeDMEType]struct{}),
		CodeDMETypes_mapString: make(map[string]*CodeDMEType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeDesignatedPointDesignatorTypes:           make(map[*CodeDesignatedPointDesignatorType]struct{}),
		CodeDesignatedPointDesignatorTypes_mapString: make(map[string]*CodeDesignatedPointDesignatorType),

		CodeDesignatedPointTypes:           make(map[*CodeDesignatedPointType]struct{}),
		CodeDesignatedPointTypes_mapString: make(map[string]*CodeDesignatedPointType),

		CodeDirectionReferenceTypes:           make(map[*CodeDirectionReferenceType]struct{}),
		CodeDirectionReferenceTypes_mapString: make(map[string]*CodeDirectionReferenceType),

		CodeDistanceIndicationTypes:           make(map[*CodeDistanceIndicationType]struct{}),
		CodeDistanceIndicationTypes_mapString: make(map[string]*CodeDistanceIndicationType),

		CodeEmissionBandTypes:           make(map[*CodeEmissionBandType]struct{}),
		CodeEmissionBandTypes_mapString: make(map[string]*CodeEmissionBandType),

		CodeILSBackCourseTypes:           make(map[*CodeILSBackCourseType]struct{}),
		CodeILSBackCourseTypes_mapString: make(map[string]*CodeILSBackCourseType),

		CodeIntegrityLevelILSTypes:           make(map[*CodeIntegrityLevelILSType]struct{}),
		CodeIntegrityLevelILSTypes_mapString: make(map[string]*CodeIntegrityLevelILSType),

		CodeMLSAzimuthTypes:           make(map[*CodeMLSAzimuthType]struct{}),
		CodeMLSAzimuthTypes_mapString: make(map[string]*CodeMLSAzimuthType),

		CodeMLSChannelTypes:           make(map[*CodeMLSChannelType]struct{}),
		CodeMLSChannelTypes_mapString: make(map[string]*CodeMLSChannelType),

		CodeMarkerBeaconSignalTypes:           make(map[*CodeMarkerBeaconSignalType]struct{}),
		CodeMarkerBeaconSignalTypes_mapString: make(map[string]*CodeMarkerBeaconSignalType),

		CodeNDBUsageTypes:           make(map[*CodeNDBUsageType]struct{}),
		CodeNDBUsageTypes_mapString: make(map[string]*CodeNDBUsageType),

		CodeNavaidDesignatorTypes:           make(map[*CodeNavaidDesignatorType]struct{}),
		CodeNavaidDesignatorTypes_mapString: make(map[string]*CodeNavaidDesignatorType),

		CodeNavaidPurposeTypes:           make(map[*CodeNavaidPurposeType]struct{}),
		CodeNavaidPurposeTypes_mapString: make(map[string]*CodeNavaidPurposeType),

		CodeNavaidServiceTypes:           make(map[*CodeNavaidServiceType]struct{}),
		CodeNavaidServiceTypes_mapString: make(map[string]*CodeNavaidServiceType),

		CodeNorthReferenceTypes:           make(map[*CodeNorthReferenceType]struct{}),
		CodeNorthReferenceTypes_mapString: make(map[string]*CodeNorthReferenceType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodePARTypes:           make(map[*CodePARType]struct{}),
		CodePARTypes_mapString: make(map[string]*CodePARType),

		CodePositionInILSTypes:           make(map[*CodePositionInILSType]struct{}),
		CodePositionInILSTypes_mapString: make(map[string]*CodePositionInILSType),

		CodePrimaryRadarTypes:           make(map[*CodePrimaryRadarType]struct{}),
		CodePrimaryRadarTypes_mapString: make(map[string]*CodePrimaryRadarType),

		CodeRadarServiceTypes:           make(map[*CodeRadarServiceType]struct{}),
		CodeRadarServiceTypes_mapString: make(map[string]*CodeRadarServiceType),

		CodeRadioEmissionTypes:           make(map[*CodeRadioEmissionType]struct{}),
		CodeRadioEmissionTypes_mapString: make(map[string]*CodeRadioEmissionType),

		CodeRadioSignalTypes:           make(map[*CodeRadioSignalType]struct{}),
		CodeRadioSignalTypes_mapString: make(map[string]*CodeRadioSignalType),

		CodeReflectorTypes:           make(map[*CodeReflectorType]struct{}),
		CodeReflectorTypes_mapString: make(map[string]*CodeReflectorType),

		CodeSignalPerformanceILSTypes:           make(map[*CodeSignalPerformanceILSType]struct{}),
		CodeSignalPerformanceILSTypes_mapString: make(map[string]*CodeSignalPerformanceILSType),

		CodeSpecialNavigationChainDesignatorTypes:           make(map[*CodeSpecialNavigationChainDesignatorType]struct{}),
		CodeSpecialNavigationChainDesignatorTypes_mapString: make(map[string]*CodeSpecialNavigationChainDesignatorType),

		CodeSpecialNavigationStationTypes:           make(map[*CodeSpecialNavigationStationType]struct{}),
		CodeSpecialNavigationStationTypes_mapString: make(map[string]*CodeSpecialNavigationStationType),

		CodeSpecialNavigationSystemTypes:           make(map[*CodeSpecialNavigationSystemType]struct{}),
		CodeSpecialNavigationSystemTypes_mapString: make(map[string]*CodeSpecialNavigationSystemType),

		CodeStandbyPowerTypes:           make(map[*CodeStandbyPowerType]struct{}),
		CodeStandbyPowerTypes_mapString: make(map[string]*CodeStandbyPowerType),

		CodeStatusNavaidTypes:           make(map[*CodeStatusNavaidType]struct{}),
		CodeStatusNavaidTypes_mapString: make(map[string]*CodeStatusNavaidType),

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

		CodeVORTypes:           make(map[*CodeVORType]struct{}),
		CodeVORTypes_mapString: make(map[string]*CodeVORType),

		CodeVerticalDatumTypes:           make(map[*CodeVerticalDatumType]struct{}),
		CodeVerticalDatumTypes_mapString: make(map[string]*CodeVerticalDatumType),

		CodeVerticalReferenceTypes:           make(map[*CodeVerticalReferenceType]struct{}),
		CodeVerticalReferenceTypes_mapString: make(map[string]*CodeVerticalReferenceType),

		CodeYesNoTypes:           make(map[*CodeYesNoType]struct{}),
		CodeYesNoTypes_mapString: make(map[string]*CodeYesNoType),

		DateMonthDayTypes:           make(map[*DateMonthDayType]struct{}),
		DateMonthDayTypes_mapString: make(map[string]*DateMonthDayType),

		DateYearTypes:           make(map[*DateYearType]struct{}),
		DateYearTypes_mapString: make(map[string]*DateYearType),

		NoSequenceTypes:           make(map[*NoSequenceType]struct{}),
		NoSequenceTypes_mapString: make(map[string]*NoSequenceType),

		TextAddressTypes:           make(map[*TextAddressType]struct{}),
		TextAddressTypes_mapString: make(map[string]*TextAddressType),

		TextDesignatorTypes:           make(map[*TextDesignatorType]struct{}),
		TextDesignatorTypes_mapString: make(map[string]*TextDesignatorType),

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

		ValFrequencyTypes:           make(map[*ValFrequencyType]struct{}),
		ValFrequencyTypes_mapString: make(map[string]*ValFrequencyType),

		ValMagneticVariationTypes:           make(map[*ValMagneticVariationType]struct{}),
		ValMagneticVariationTypes_mapString: make(map[string]*ValMagneticVariationType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CodeAirspacePointPositionType_stagedOrder: make(map[*CodeAirspacePointPositionType]uint),
		CodeAirspacePointPositionType_orderStaged: make(map[uint]*CodeAirspacePointPositionType),
		CodeAirspacePointPositionTypes_reference:  make(map[*CodeAirspacePointPositionType]*CodeAirspacePointPositionType),

		CodeAirspacePointRoleType_stagedOrder: make(map[*CodeAirspacePointRoleType]uint),
		CodeAirspacePointRoleType_orderStaged: make(map[uint]*CodeAirspacePointRoleType),
		CodeAirspacePointRoleTypes_reference:  make(map[*CodeAirspacePointRoleType]*CodeAirspacePointRoleType),

		CodeAltitudeUseType_stagedOrder: make(map[*CodeAltitudeUseType]uint),
		CodeAltitudeUseType_orderStaged: make(map[uint]*CodeAltitudeUseType),
		CodeAltitudeUseTypes_reference:  make(map[*CodeAltitudeUseType]*CodeAltitudeUseType),

		CodeAuralMorseType_stagedOrder: make(map[*CodeAuralMorseType]uint),
		CodeAuralMorseType_orderStaged: make(map[uint]*CodeAuralMorseType),
		CodeAuralMorseTypes_reference:  make(map[*CodeAuralMorseType]*CodeAuralMorseType),

		CodeAuthorityRoleType_stagedOrder: make(map[*CodeAuthorityRoleType]uint),
		CodeAuthorityRoleType_orderStaged: make(map[uint]*CodeAuthorityRoleType),
		CodeAuthorityRoleTypes_reference:  make(map[*CodeAuthorityRoleType]*CodeAuthorityRoleType),

		CodeBearingType_stagedOrder: make(map[*CodeBearingType]uint),
		CodeBearingType_orderStaged: make(map[uint]*CodeBearingType),
		CodeBearingTypes_reference:  make(map[*CodeBearingType]*CodeBearingType),

		CodeCardinalDirectionType_stagedOrder: make(map[*CodeCardinalDirectionType]uint),
		CodeCardinalDirectionType_orderStaged: make(map[uint]*CodeCardinalDirectionType),
		CodeCardinalDirectionTypes_reference:  make(map[*CodeCardinalDirectionType]*CodeCardinalDirectionType),

		CodeCheckpointCategoryType_stagedOrder: make(map[*CodeCheckpointCategoryType]uint),
		CodeCheckpointCategoryType_orderStaged: make(map[uint]*CodeCheckpointCategoryType),
		CodeCheckpointCategoryTypes_reference:  make(map[*CodeCheckpointCategoryType]*CodeCheckpointCategoryType),

		CodeCourseQualityILSType_stagedOrder: make(map[*CodeCourseQualityILSType]uint),
		CodeCourseQualityILSType_orderStaged: make(map[uint]*CodeCourseQualityILSType),
		CodeCourseQualityILSTypes_reference:  make(map[*CodeCourseQualityILSType]*CodeCourseQualityILSType),

		CodeDMEChannelType_stagedOrder: make(map[*CodeDMEChannelType]uint),
		CodeDMEChannelType_orderStaged: make(map[uint]*CodeDMEChannelType),
		CodeDMEChannelTypes_reference:  make(map[*CodeDMEChannelType]*CodeDMEChannelType),

		CodeDMEType_stagedOrder: make(map[*CodeDMEType]uint),
		CodeDMEType_orderStaged: make(map[uint]*CodeDMEType),
		CodeDMETypes_reference:  make(map[*CodeDMEType]*CodeDMEType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeDesignatedPointDesignatorType_stagedOrder: make(map[*CodeDesignatedPointDesignatorType]uint),
		CodeDesignatedPointDesignatorType_orderStaged: make(map[uint]*CodeDesignatedPointDesignatorType),
		CodeDesignatedPointDesignatorTypes_reference:  make(map[*CodeDesignatedPointDesignatorType]*CodeDesignatedPointDesignatorType),

		CodeDesignatedPointType_stagedOrder: make(map[*CodeDesignatedPointType]uint),
		CodeDesignatedPointType_orderStaged: make(map[uint]*CodeDesignatedPointType),
		CodeDesignatedPointTypes_reference:  make(map[*CodeDesignatedPointType]*CodeDesignatedPointType),

		CodeDirectionReferenceType_stagedOrder: make(map[*CodeDirectionReferenceType]uint),
		CodeDirectionReferenceType_orderStaged: make(map[uint]*CodeDirectionReferenceType),
		CodeDirectionReferenceTypes_reference:  make(map[*CodeDirectionReferenceType]*CodeDirectionReferenceType),

		CodeDistanceIndicationType_stagedOrder: make(map[*CodeDistanceIndicationType]uint),
		CodeDistanceIndicationType_orderStaged: make(map[uint]*CodeDistanceIndicationType),
		CodeDistanceIndicationTypes_reference:  make(map[*CodeDistanceIndicationType]*CodeDistanceIndicationType),

		CodeEmissionBandType_stagedOrder: make(map[*CodeEmissionBandType]uint),
		CodeEmissionBandType_orderStaged: make(map[uint]*CodeEmissionBandType),
		CodeEmissionBandTypes_reference:  make(map[*CodeEmissionBandType]*CodeEmissionBandType),

		CodeILSBackCourseType_stagedOrder: make(map[*CodeILSBackCourseType]uint),
		CodeILSBackCourseType_orderStaged: make(map[uint]*CodeILSBackCourseType),
		CodeILSBackCourseTypes_reference:  make(map[*CodeILSBackCourseType]*CodeILSBackCourseType),

		CodeIntegrityLevelILSType_stagedOrder: make(map[*CodeIntegrityLevelILSType]uint),
		CodeIntegrityLevelILSType_orderStaged: make(map[uint]*CodeIntegrityLevelILSType),
		CodeIntegrityLevelILSTypes_reference:  make(map[*CodeIntegrityLevelILSType]*CodeIntegrityLevelILSType),

		CodeMLSAzimuthType_stagedOrder: make(map[*CodeMLSAzimuthType]uint),
		CodeMLSAzimuthType_orderStaged: make(map[uint]*CodeMLSAzimuthType),
		CodeMLSAzimuthTypes_reference:  make(map[*CodeMLSAzimuthType]*CodeMLSAzimuthType),

		CodeMLSChannelType_stagedOrder: make(map[*CodeMLSChannelType]uint),
		CodeMLSChannelType_orderStaged: make(map[uint]*CodeMLSChannelType),
		CodeMLSChannelTypes_reference:  make(map[*CodeMLSChannelType]*CodeMLSChannelType),

		CodeMarkerBeaconSignalType_stagedOrder: make(map[*CodeMarkerBeaconSignalType]uint),
		CodeMarkerBeaconSignalType_orderStaged: make(map[uint]*CodeMarkerBeaconSignalType),
		CodeMarkerBeaconSignalTypes_reference:  make(map[*CodeMarkerBeaconSignalType]*CodeMarkerBeaconSignalType),

		CodeNDBUsageType_stagedOrder: make(map[*CodeNDBUsageType]uint),
		CodeNDBUsageType_orderStaged: make(map[uint]*CodeNDBUsageType),
		CodeNDBUsageTypes_reference:  make(map[*CodeNDBUsageType]*CodeNDBUsageType),

		CodeNavaidDesignatorType_stagedOrder: make(map[*CodeNavaidDesignatorType]uint),
		CodeNavaidDesignatorType_orderStaged: make(map[uint]*CodeNavaidDesignatorType),
		CodeNavaidDesignatorTypes_reference:  make(map[*CodeNavaidDesignatorType]*CodeNavaidDesignatorType),

		CodeNavaidPurposeType_stagedOrder: make(map[*CodeNavaidPurposeType]uint),
		CodeNavaidPurposeType_orderStaged: make(map[uint]*CodeNavaidPurposeType),
		CodeNavaidPurposeTypes_reference:  make(map[*CodeNavaidPurposeType]*CodeNavaidPurposeType),

		CodeNavaidServiceType_stagedOrder: make(map[*CodeNavaidServiceType]uint),
		CodeNavaidServiceType_orderStaged: make(map[uint]*CodeNavaidServiceType),
		CodeNavaidServiceTypes_reference:  make(map[*CodeNavaidServiceType]*CodeNavaidServiceType),

		CodeNorthReferenceType_stagedOrder: make(map[*CodeNorthReferenceType]uint),
		CodeNorthReferenceType_orderStaged: make(map[uint]*CodeNorthReferenceType),
		CodeNorthReferenceTypes_reference:  make(map[*CodeNorthReferenceType]*CodeNorthReferenceType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodePARType_stagedOrder: make(map[*CodePARType]uint),
		CodePARType_orderStaged: make(map[uint]*CodePARType),
		CodePARTypes_reference:  make(map[*CodePARType]*CodePARType),

		CodePositionInILSType_stagedOrder: make(map[*CodePositionInILSType]uint),
		CodePositionInILSType_orderStaged: make(map[uint]*CodePositionInILSType),
		CodePositionInILSTypes_reference:  make(map[*CodePositionInILSType]*CodePositionInILSType),

		CodePrimaryRadarType_stagedOrder: make(map[*CodePrimaryRadarType]uint),
		CodePrimaryRadarType_orderStaged: make(map[uint]*CodePrimaryRadarType),
		CodePrimaryRadarTypes_reference:  make(map[*CodePrimaryRadarType]*CodePrimaryRadarType),

		CodeRadarServiceType_stagedOrder: make(map[*CodeRadarServiceType]uint),
		CodeRadarServiceType_orderStaged: make(map[uint]*CodeRadarServiceType),
		CodeRadarServiceTypes_reference:  make(map[*CodeRadarServiceType]*CodeRadarServiceType),

		CodeRadioEmissionType_stagedOrder: make(map[*CodeRadioEmissionType]uint),
		CodeRadioEmissionType_orderStaged: make(map[uint]*CodeRadioEmissionType),
		CodeRadioEmissionTypes_reference:  make(map[*CodeRadioEmissionType]*CodeRadioEmissionType),

		CodeRadioSignalType_stagedOrder: make(map[*CodeRadioSignalType]uint),
		CodeRadioSignalType_orderStaged: make(map[uint]*CodeRadioSignalType),
		CodeRadioSignalTypes_reference:  make(map[*CodeRadioSignalType]*CodeRadioSignalType),

		CodeReflectorType_stagedOrder: make(map[*CodeReflectorType]uint),
		CodeReflectorType_orderStaged: make(map[uint]*CodeReflectorType),
		CodeReflectorTypes_reference:  make(map[*CodeReflectorType]*CodeReflectorType),

		CodeSignalPerformanceILSType_stagedOrder: make(map[*CodeSignalPerformanceILSType]uint),
		CodeSignalPerformanceILSType_orderStaged: make(map[uint]*CodeSignalPerformanceILSType),
		CodeSignalPerformanceILSTypes_reference:  make(map[*CodeSignalPerformanceILSType]*CodeSignalPerformanceILSType),

		CodeSpecialNavigationChainDesignatorType_stagedOrder: make(map[*CodeSpecialNavigationChainDesignatorType]uint),
		CodeSpecialNavigationChainDesignatorType_orderStaged: make(map[uint]*CodeSpecialNavigationChainDesignatorType),
		CodeSpecialNavigationChainDesignatorTypes_reference:  make(map[*CodeSpecialNavigationChainDesignatorType]*CodeSpecialNavigationChainDesignatorType),

		CodeSpecialNavigationStationType_stagedOrder: make(map[*CodeSpecialNavigationStationType]uint),
		CodeSpecialNavigationStationType_orderStaged: make(map[uint]*CodeSpecialNavigationStationType),
		CodeSpecialNavigationStationTypes_reference:  make(map[*CodeSpecialNavigationStationType]*CodeSpecialNavigationStationType),

		CodeSpecialNavigationSystemType_stagedOrder: make(map[*CodeSpecialNavigationSystemType]uint),
		CodeSpecialNavigationSystemType_orderStaged: make(map[uint]*CodeSpecialNavigationSystemType),
		CodeSpecialNavigationSystemTypes_reference:  make(map[*CodeSpecialNavigationSystemType]*CodeSpecialNavigationSystemType),

		CodeStandbyPowerType_stagedOrder: make(map[*CodeStandbyPowerType]uint),
		CodeStandbyPowerType_orderStaged: make(map[uint]*CodeStandbyPowerType),
		CodeStandbyPowerTypes_reference:  make(map[*CodeStandbyPowerType]*CodeStandbyPowerType),

		CodeStatusNavaidType_stagedOrder: make(map[*CodeStatusNavaidType]uint),
		CodeStatusNavaidType_orderStaged: make(map[uint]*CodeStatusNavaidType),
		CodeStatusNavaidTypes_reference:  make(map[*CodeStatusNavaidType]*CodeStatusNavaidType),

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

		CodeVORType_stagedOrder: make(map[*CodeVORType]uint),
		CodeVORType_orderStaged: make(map[uint]*CodeVORType),
		CodeVORTypes_reference:  make(map[*CodeVORType]*CodeVORType),

		CodeVerticalDatumType_stagedOrder: make(map[*CodeVerticalDatumType]uint),
		CodeVerticalDatumType_orderStaged: make(map[uint]*CodeVerticalDatumType),
		CodeVerticalDatumTypes_reference:  make(map[*CodeVerticalDatumType]*CodeVerticalDatumType),

		CodeVerticalReferenceType_stagedOrder: make(map[*CodeVerticalReferenceType]uint),
		CodeVerticalReferenceType_orderStaged: make(map[uint]*CodeVerticalReferenceType),
		CodeVerticalReferenceTypes_reference:  make(map[*CodeVerticalReferenceType]*CodeVerticalReferenceType),

		CodeYesNoType_stagedOrder: make(map[*CodeYesNoType]uint),
		CodeYesNoType_orderStaged: make(map[uint]*CodeYesNoType),
		CodeYesNoTypes_reference:  make(map[*CodeYesNoType]*CodeYesNoType),

		DateMonthDayType_stagedOrder: make(map[*DateMonthDayType]uint),
		DateMonthDayType_orderStaged: make(map[uint]*DateMonthDayType),
		DateMonthDayTypes_reference:  make(map[*DateMonthDayType]*DateMonthDayType),

		DateYearType_stagedOrder: make(map[*DateYearType]uint),
		DateYearType_orderStaged: make(map[uint]*DateYearType),
		DateYearTypes_reference:  make(map[*DateYearType]*DateYearType),

		NoSequenceType_stagedOrder: make(map[*NoSequenceType]uint),
		NoSequenceType_orderStaged: make(map[uint]*NoSequenceType),
		NoSequenceTypes_reference:  make(map[*NoSequenceType]*NoSequenceType),

		TextAddressType_stagedOrder: make(map[*TextAddressType]uint),
		TextAddressType_orderStaged: make(map[uint]*TextAddressType),
		TextAddressTypes_reference:  make(map[*TextAddressType]*TextAddressType),

		TextDesignatorType_stagedOrder: make(map[*TextDesignatorType]uint),
		TextDesignatorType_orderStaged: make(map[uint]*TextDesignatorType),
		TextDesignatorTypes_reference:  make(map[*TextDesignatorType]*TextDesignatorType),

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

		ValFrequencyType_stagedOrder: make(map[*ValFrequencyType]uint),
		ValFrequencyType_orderStaged: make(map[uint]*ValFrequencyType),
		ValFrequencyTypes_reference:  make(map[*ValFrequencyType]*ValFrequencyType),

		ValMagneticVariationType_stagedOrder: make(map[*ValMagneticVariationType]uint),
		ValMagneticVariationType_orderStaged: make(map[uint]*ValMagneticVariationType),
		ValMagneticVariationTypes_reference:  make(map[*ValMagneticVariationType]*ValMagneticVariationType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeAirspacePointPositionType": &CodeAirspacePointPositionTypeUnmarshaller{},

			"CodeAirspacePointRoleType": &CodeAirspacePointRoleTypeUnmarshaller{},

			"CodeAltitudeUseType": &CodeAltitudeUseTypeUnmarshaller{},

			"CodeAuralMorseType": &CodeAuralMorseTypeUnmarshaller{},

			"CodeAuthorityRoleType": &CodeAuthorityRoleTypeUnmarshaller{},

			"CodeBearingType": &CodeBearingTypeUnmarshaller{},

			"CodeCardinalDirectionType": &CodeCardinalDirectionTypeUnmarshaller{},

			"CodeCheckpointCategoryType": &CodeCheckpointCategoryTypeUnmarshaller{},

			"CodeCourseQualityILSType": &CodeCourseQualityILSTypeUnmarshaller{},

			"CodeDMEChannelType": &CodeDMEChannelTypeUnmarshaller{},

			"CodeDMEType": &CodeDMETypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeDesignatedPointDesignatorType": &CodeDesignatedPointDesignatorTypeUnmarshaller{},

			"CodeDesignatedPointType": &CodeDesignatedPointTypeUnmarshaller{},

			"CodeDirectionReferenceType": &CodeDirectionReferenceTypeUnmarshaller{},

			"CodeDistanceIndicationType": &CodeDistanceIndicationTypeUnmarshaller{},

			"CodeEmissionBandType": &CodeEmissionBandTypeUnmarshaller{},

			"CodeILSBackCourseType": &CodeILSBackCourseTypeUnmarshaller{},

			"CodeIntegrityLevelILSType": &CodeIntegrityLevelILSTypeUnmarshaller{},

			"CodeMLSAzimuthType": &CodeMLSAzimuthTypeUnmarshaller{},

			"CodeMLSChannelType": &CodeMLSChannelTypeUnmarshaller{},

			"CodeMarkerBeaconSignalType": &CodeMarkerBeaconSignalTypeUnmarshaller{},

			"CodeNDBUsageType": &CodeNDBUsageTypeUnmarshaller{},

			"CodeNavaidDesignatorType": &CodeNavaidDesignatorTypeUnmarshaller{},

			"CodeNavaidPurposeType": &CodeNavaidPurposeTypeUnmarshaller{},

			"CodeNavaidServiceType": &CodeNavaidServiceTypeUnmarshaller{},

			"CodeNorthReferenceType": &CodeNorthReferenceTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodePARType": &CodePARTypeUnmarshaller{},

			"CodePositionInILSType": &CodePositionInILSTypeUnmarshaller{},

			"CodePrimaryRadarType": &CodePrimaryRadarTypeUnmarshaller{},

			"CodeRadarServiceType": &CodeRadarServiceTypeUnmarshaller{},

			"CodeRadioEmissionType": &CodeRadioEmissionTypeUnmarshaller{},

			"CodeRadioSignalType": &CodeRadioSignalTypeUnmarshaller{},

			"CodeReflectorType": &CodeReflectorTypeUnmarshaller{},

			"CodeSignalPerformanceILSType": &CodeSignalPerformanceILSTypeUnmarshaller{},

			"CodeSpecialNavigationChainDesignatorType": &CodeSpecialNavigationChainDesignatorTypeUnmarshaller{},

			"CodeSpecialNavigationStationType": &CodeSpecialNavigationStationTypeUnmarshaller{},

			"CodeSpecialNavigationSystemType": &CodeSpecialNavigationSystemTypeUnmarshaller{},

			"CodeStandbyPowerType": &CodeStandbyPowerTypeUnmarshaller{},

			"CodeStatusNavaidType": &CodeStatusNavaidTypeUnmarshaller{},

			"CodeTACANChannelType": &CodeTACANChannelTypeUnmarshaller{},

			"CodeTelecomNetworkType": &CodeTelecomNetworkTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeTransponderType": &CodeTransponderTypeUnmarshaller{},

			"CodeVORType": &CodeVORTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeVerticalReferenceType": &CodeVerticalReferenceTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"DateYearType": &DateYearTypeUnmarshaller{},

			"NoSequenceType": &NoSequenceTypeUnmarshaller{},

			"TextAddressType": &TextAddressTypeUnmarshaller{},

			"TextDesignatorType": &TextDesignatorTypeUnmarshaller{},

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

			"ValFrequencyType": &ValFrequencyTypeUnmarshaller{},

			"ValMagneticVariationType": &ValMagneticVariationTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeAirspacePointPositionType"] = len(stage.CodeAirspacePointPositionTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspacePointRoleType"] = len(stage.CodeAirspacePointRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeUseType"] = len(stage.CodeAltitudeUseTypes)
	stage.Map_GongStructName_InstancesNb["CodeAuralMorseType"] = len(stage.CodeAuralMorseTypes)
	stage.Map_GongStructName_InstancesNb["CodeAuthorityRoleType"] = len(stage.CodeAuthorityRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeBearingType"] = len(stage.CodeBearingTypes)
	stage.Map_GongStructName_InstancesNb["CodeCardinalDirectionType"] = len(stage.CodeCardinalDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeCheckpointCategoryType"] = len(stage.CodeCheckpointCategoryTypes)
	stage.Map_GongStructName_InstancesNb["CodeCourseQualityILSType"] = len(stage.CodeCourseQualityILSTypes)
	stage.Map_GongStructName_InstancesNb["CodeDMEChannelType"] = len(stage.CodeDMEChannelTypes)
	stage.Map_GongStructName_InstancesNb["CodeDMEType"] = len(stage.CodeDMETypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeDesignatedPointDesignatorType"] = len(stage.CodeDesignatedPointDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeDesignatedPointType"] = len(stage.CodeDesignatedPointTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionReferenceType"] = len(stage.CodeDirectionReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeDistanceIndicationType"] = len(stage.CodeDistanceIndicationTypes)
	stage.Map_GongStructName_InstancesNb["CodeEmissionBandType"] = len(stage.CodeEmissionBandTypes)
	stage.Map_GongStructName_InstancesNb["CodeILSBackCourseType"] = len(stage.CodeILSBackCourseTypes)
	stage.Map_GongStructName_InstancesNb["CodeIntegrityLevelILSType"] = len(stage.CodeIntegrityLevelILSTypes)
	stage.Map_GongStructName_InstancesNb["CodeMLSAzimuthType"] = len(stage.CodeMLSAzimuthTypes)
	stage.Map_GongStructName_InstancesNb["CodeMLSChannelType"] = len(stage.CodeMLSChannelTypes)
	stage.Map_GongStructName_InstancesNb["CodeMarkerBeaconSignalType"] = len(stage.CodeMarkerBeaconSignalTypes)
	stage.Map_GongStructName_InstancesNb["CodeNDBUsageType"] = len(stage.CodeNDBUsageTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavaidDesignatorType"] = len(stage.CodeNavaidDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavaidPurposeType"] = len(stage.CodeNavaidPurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavaidServiceType"] = len(stage.CodeNavaidServiceTypes)
	stage.Map_GongStructName_InstancesNb["CodeNorthReferenceType"] = len(stage.CodeNorthReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodePARType"] = len(stage.CodePARTypes)
	stage.Map_GongStructName_InstancesNb["CodePositionInILSType"] = len(stage.CodePositionInILSTypes)
	stage.Map_GongStructName_InstancesNb["CodePrimaryRadarType"] = len(stage.CodePrimaryRadarTypes)
	stage.Map_GongStructName_InstancesNb["CodeRadarServiceType"] = len(stage.CodeRadarServiceTypes)
	stage.Map_GongStructName_InstancesNb["CodeRadioEmissionType"] = len(stage.CodeRadioEmissionTypes)
	stage.Map_GongStructName_InstancesNb["CodeRadioSignalType"] = len(stage.CodeRadioSignalTypes)
	stage.Map_GongStructName_InstancesNb["CodeReflectorType"] = len(stage.CodeReflectorTypes)
	stage.Map_GongStructName_InstancesNb["CodeSignalPerformanceILSType"] = len(stage.CodeSignalPerformanceILSTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpecialNavigationChainDesignatorType"] = len(stage.CodeSpecialNavigationChainDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpecialNavigationStationType"] = len(stage.CodeSpecialNavigationStationTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpecialNavigationSystemType"] = len(stage.CodeSpecialNavigationSystemTypes)
	stage.Map_GongStructName_InstancesNb["CodeStandbyPowerType"] = len(stage.CodeStandbyPowerTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusNavaidType"] = len(stage.CodeStatusNavaidTypes)
	stage.Map_GongStructName_InstancesNb["CodeTACANChannelType"] = len(stage.CodeTACANChannelTypes)
	stage.Map_GongStructName_InstancesNb["CodeTelecomNetworkType"] = len(stage.CodeTelecomNetworkTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTransponderType"] = len(stage.CodeTransponderTypes)
	stage.Map_GongStructName_InstancesNb["CodeVORType"] = len(stage.CodeVORTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalReferenceType"] = len(stage.CodeVerticalReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["DateYearType"] = len(stage.DateYearTypes)
	stage.Map_GongStructName_InstancesNb["NoSequenceType"] = len(stage.NoSequenceTypes)
	stage.Map_GongStructName_InstancesNb["TextAddressType"] = len(stage.TextAddressTypes)
	stage.Map_GongStructName_InstancesNb["TextDesignatorType"] = len(stage.TextDesignatorTypes)
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
	stage.Map_GongStructName_InstancesNb["ValFrequencyType"] = len(stage.ValFrequencyTypes)
	stage.Map_GongStructName_InstancesNb["ValMagneticVariationType"] = len(stage.ValMagneticVariationTypes)
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
// Stage puts codeairspacepointpositiontype to the model stage
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) Stage(stage *Stage) *CodeAirspacePointPositionType {
	__gong__stage(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionType_stagedOrder, stage.CodeAirspacePointPositionType_orderStaged, &stage.CodeAirspacePointPositionTypeOrder, stage.CodeAirspacePointPositionTypes_mapString, codeairspacepointpositiontype, codeairspacepointpositiontype.Name)
	return codeairspacepointpositiontype
}

// StagePreserveOrder puts codeairspacepointpositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspacePointPositionTypeOrder
// - update stage.CodeAirspacePointPositionTypeOrder accordingly
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionType_stagedOrder, stage.CodeAirspacePointPositionType_orderStaged, &stage.CodeAirspacePointPositionTypeOrder, stage.CodeAirspacePointPositionTypes_mapString, codeairspacepointpositiontype, order, codeairspacepointpositiontype.Name)
}

// Unstage removes codeairspacepointpositiontype off the model stage
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) Unstage(stage *Stage) *CodeAirspacePointPositionType {
	__gong__unstage(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionTypes_mapString, codeairspacepointpositiontype, codeairspacepointpositiontype.Name)
	return codeairspacepointpositiontype
}

// UnstageVoid removes codeairspacepointpositiontype off the model stage
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) UnstageVoid(stage *Stage) {
	codeairspacepointpositiontype.Unstage(stage)
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) StageVoid(stage *Stage) {
	codeairspacepointpositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GetName() (res string) {
	return codeairspacepointpositiontype.Name
}

// for satisfaction of GongStruct interface
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) SetName(name string) {
	codeairspacepointpositiontype.Name = name
}

// Stage puts codeairspacepointroletype to the model stage
func (codeairspacepointroletype *CodeAirspacePointRoleType) Stage(stage *Stage) *CodeAirspacePointRoleType {
	__gong__stage(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleType_stagedOrder, stage.CodeAirspacePointRoleType_orderStaged, &stage.CodeAirspacePointRoleTypeOrder, stage.CodeAirspacePointRoleTypes_mapString, codeairspacepointroletype, codeairspacepointroletype.Name)
	return codeairspacepointroletype
}

// StagePreserveOrder puts codeairspacepointroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspacePointRoleTypeOrder
// - update stage.CodeAirspacePointRoleTypeOrder accordingly
func (codeairspacepointroletype *CodeAirspacePointRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleType_stagedOrder, stage.CodeAirspacePointRoleType_orderStaged, &stage.CodeAirspacePointRoleTypeOrder, stage.CodeAirspacePointRoleTypes_mapString, codeairspacepointroletype, order, codeairspacepointroletype.Name)
}

// Unstage removes codeairspacepointroletype off the model stage
func (codeairspacepointroletype *CodeAirspacePointRoleType) Unstage(stage *Stage) *CodeAirspacePointRoleType {
	__gong__unstage(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleTypes_mapString, codeairspacepointroletype, codeairspacepointroletype.Name)
	return codeairspacepointroletype
}

// UnstageVoid removes codeairspacepointroletype off the model stage
func (codeairspacepointroletype *CodeAirspacePointRoleType) UnstageVoid(stage *Stage) {
	codeairspacepointroletype.Unstage(stage)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) StageVoid(stage *Stage) {
	codeairspacepointroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspacepointroletype *CodeAirspacePointRoleType) GetName() (res string) {
	return codeairspacepointroletype.Name
}

// for satisfaction of GongStruct interface
func (codeairspacepointroletype *CodeAirspacePointRoleType) SetName(name string) {
	codeairspacepointroletype.Name = name
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

// Stage puts codeauralmorsetype to the model stage
func (codeauralmorsetype *CodeAuralMorseType) Stage(stage *Stage) *CodeAuralMorseType {
	__gong__stage(stage.CodeAuralMorseTypes, stage.CodeAuralMorseType_stagedOrder, stage.CodeAuralMorseType_orderStaged, &stage.CodeAuralMorseTypeOrder, stage.CodeAuralMorseTypes_mapString, codeauralmorsetype, codeauralmorsetype.Name)
	return codeauralmorsetype
}

// StagePreserveOrder puts codeauralmorsetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAuralMorseTypeOrder
// - update stage.CodeAuralMorseTypeOrder accordingly
func (codeauralmorsetype *CodeAuralMorseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAuralMorseTypes, stage.CodeAuralMorseType_stagedOrder, stage.CodeAuralMorseType_orderStaged, &stage.CodeAuralMorseTypeOrder, stage.CodeAuralMorseTypes_mapString, codeauralmorsetype, order, codeauralmorsetype.Name)
}

// Unstage removes codeauralmorsetype off the model stage
func (codeauralmorsetype *CodeAuralMorseType) Unstage(stage *Stage) *CodeAuralMorseType {
	__gong__unstage(stage.CodeAuralMorseTypes, stage.CodeAuralMorseTypes_mapString, codeauralmorsetype, codeauralmorsetype.Name)
	return codeauralmorsetype
}

// UnstageVoid removes codeauralmorsetype off the model stage
func (codeauralmorsetype *CodeAuralMorseType) UnstageVoid(stage *Stage) {
	codeauralmorsetype.Unstage(stage)
}

func (codeauralmorsetype *CodeAuralMorseType) StageVoid(stage *Stage) {
	codeauralmorsetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeauralmorsetype *CodeAuralMorseType) GetName() (res string) {
	return codeauralmorsetype.Name
}

// for satisfaction of GongStruct interface
func (codeauralmorsetype *CodeAuralMorseType) SetName(name string) {
	codeauralmorsetype.Name = name
}

// Stage puts codeauthorityroletype to the model stage
func (codeauthorityroletype *CodeAuthorityRoleType) Stage(stage *Stage) *CodeAuthorityRoleType {
	__gong__stage(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleType_stagedOrder, stage.CodeAuthorityRoleType_orderStaged, &stage.CodeAuthorityRoleTypeOrder, stage.CodeAuthorityRoleTypes_mapString, codeauthorityroletype, codeauthorityroletype.Name)
	return codeauthorityroletype
}

// StagePreserveOrder puts codeauthorityroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAuthorityRoleTypeOrder
// - update stage.CodeAuthorityRoleTypeOrder accordingly
func (codeauthorityroletype *CodeAuthorityRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleType_stagedOrder, stage.CodeAuthorityRoleType_orderStaged, &stage.CodeAuthorityRoleTypeOrder, stage.CodeAuthorityRoleTypes_mapString, codeauthorityroletype, order, codeauthorityroletype.Name)
}

// Unstage removes codeauthorityroletype off the model stage
func (codeauthorityroletype *CodeAuthorityRoleType) Unstage(stage *Stage) *CodeAuthorityRoleType {
	__gong__unstage(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleTypes_mapString, codeauthorityroletype, codeauthorityroletype.Name)
	return codeauthorityroletype
}

// UnstageVoid removes codeauthorityroletype off the model stage
func (codeauthorityroletype *CodeAuthorityRoleType) UnstageVoid(stage *Stage) {
	codeauthorityroletype.Unstage(stage)
}

func (codeauthorityroletype *CodeAuthorityRoleType) StageVoid(stage *Stage) {
	codeauthorityroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeauthorityroletype *CodeAuthorityRoleType) GetName() (res string) {
	return codeauthorityroletype.Name
}

// for satisfaction of GongStruct interface
func (codeauthorityroletype *CodeAuthorityRoleType) SetName(name string) {
	codeauthorityroletype.Name = name
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

// Stage puts codecheckpointcategorytype to the model stage
func (codecheckpointcategorytype *CodeCheckpointCategoryType) Stage(stage *Stage) *CodeCheckpointCategoryType {
	__gong__stage(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryType_stagedOrder, stage.CodeCheckpointCategoryType_orderStaged, &stage.CodeCheckpointCategoryTypeOrder, stage.CodeCheckpointCategoryTypes_mapString, codecheckpointcategorytype, codecheckpointcategorytype.Name)
	return codecheckpointcategorytype
}

// StagePreserveOrder puts codecheckpointcategorytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCheckpointCategoryTypeOrder
// - update stage.CodeCheckpointCategoryTypeOrder accordingly
func (codecheckpointcategorytype *CodeCheckpointCategoryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryType_stagedOrder, stage.CodeCheckpointCategoryType_orderStaged, &stage.CodeCheckpointCategoryTypeOrder, stage.CodeCheckpointCategoryTypes_mapString, codecheckpointcategorytype, order, codecheckpointcategorytype.Name)
}

// Unstage removes codecheckpointcategorytype off the model stage
func (codecheckpointcategorytype *CodeCheckpointCategoryType) Unstage(stage *Stage) *CodeCheckpointCategoryType {
	__gong__unstage(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryTypes_mapString, codecheckpointcategorytype, codecheckpointcategorytype.Name)
	return codecheckpointcategorytype
}

// UnstageVoid removes codecheckpointcategorytype off the model stage
func (codecheckpointcategorytype *CodeCheckpointCategoryType) UnstageVoid(stage *Stage) {
	codecheckpointcategorytype.Unstage(stage)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) StageVoid(stage *Stage) {
	codecheckpointcategorytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecheckpointcategorytype *CodeCheckpointCategoryType) GetName() (res string) {
	return codecheckpointcategorytype.Name
}

// for satisfaction of GongStruct interface
func (codecheckpointcategorytype *CodeCheckpointCategoryType) SetName(name string) {
	codecheckpointcategorytype.Name = name
}

// Stage puts codecoursequalityilstype to the model stage
func (codecoursequalityilstype *CodeCourseQualityILSType) Stage(stage *Stage) *CodeCourseQualityILSType {
	__gong__stage(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSType_stagedOrder, stage.CodeCourseQualityILSType_orderStaged, &stage.CodeCourseQualityILSTypeOrder, stage.CodeCourseQualityILSTypes_mapString, codecoursequalityilstype, codecoursequalityilstype.Name)
	return codecoursequalityilstype
}

// StagePreserveOrder puts codecoursequalityilstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCourseQualityILSTypeOrder
// - update stage.CodeCourseQualityILSTypeOrder accordingly
func (codecoursequalityilstype *CodeCourseQualityILSType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSType_stagedOrder, stage.CodeCourseQualityILSType_orderStaged, &stage.CodeCourseQualityILSTypeOrder, stage.CodeCourseQualityILSTypes_mapString, codecoursequalityilstype, order, codecoursequalityilstype.Name)
}

// Unstage removes codecoursequalityilstype off the model stage
func (codecoursequalityilstype *CodeCourseQualityILSType) Unstage(stage *Stage) *CodeCourseQualityILSType {
	__gong__unstage(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSTypes_mapString, codecoursequalityilstype, codecoursequalityilstype.Name)
	return codecoursequalityilstype
}

// UnstageVoid removes codecoursequalityilstype off the model stage
func (codecoursequalityilstype *CodeCourseQualityILSType) UnstageVoid(stage *Stage) {
	codecoursequalityilstype.Unstage(stage)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) StageVoid(stage *Stage) {
	codecoursequalityilstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecoursequalityilstype *CodeCourseQualityILSType) GetName() (res string) {
	return codecoursequalityilstype.Name
}

// for satisfaction of GongStruct interface
func (codecoursequalityilstype *CodeCourseQualityILSType) SetName(name string) {
	codecoursequalityilstype.Name = name
}

// Stage puts codedmechanneltype to the model stage
func (codedmechanneltype *CodeDMEChannelType) Stage(stage *Stage) *CodeDMEChannelType {
	__gong__stage(stage.CodeDMEChannelTypes, stage.CodeDMEChannelType_stagedOrder, stage.CodeDMEChannelType_orderStaged, &stage.CodeDMEChannelTypeOrder, stage.CodeDMEChannelTypes_mapString, codedmechanneltype, codedmechanneltype.Name)
	return codedmechanneltype
}

// StagePreserveOrder puts codedmechanneltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDMEChannelTypeOrder
// - update stage.CodeDMEChannelTypeOrder accordingly
func (codedmechanneltype *CodeDMEChannelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDMEChannelTypes, stage.CodeDMEChannelType_stagedOrder, stage.CodeDMEChannelType_orderStaged, &stage.CodeDMEChannelTypeOrder, stage.CodeDMEChannelTypes_mapString, codedmechanneltype, order, codedmechanneltype.Name)
}

// Unstage removes codedmechanneltype off the model stage
func (codedmechanneltype *CodeDMEChannelType) Unstage(stage *Stage) *CodeDMEChannelType {
	__gong__unstage(stage.CodeDMEChannelTypes, stage.CodeDMEChannelTypes_mapString, codedmechanneltype, codedmechanneltype.Name)
	return codedmechanneltype
}

// UnstageVoid removes codedmechanneltype off the model stage
func (codedmechanneltype *CodeDMEChannelType) UnstageVoid(stage *Stage) {
	codedmechanneltype.Unstage(stage)
}

func (codedmechanneltype *CodeDMEChannelType) StageVoid(stage *Stage) {
	codedmechanneltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedmechanneltype *CodeDMEChannelType) GetName() (res string) {
	return codedmechanneltype.Name
}

// for satisfaction of GongStruct interface
func (codedmechanneltype *CodeDMEChannelType) SetName(name string) {
	codedmechanneltype.Name = name
}

// Stage puts codedmetype to the model stage
func (codedmetype *CodeDMEType) Stage(stage *Stage) *CodeDMEType {
	__gong__stage(stage.CodeDMETypes, stage.CodeDMEType_stagedOrder, stage.CodeDMEType_orderStaged, &stage.CodeDMETypeOrder, stage.CodeDMETypes_mapString, codedmetype, codedmetype.Name)
	return codedmetype
}

// StagePreserveOrder puts codedmetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDMETypeOrder
// - update stage.CodeDMETypeOrder accordingly
func (codedmetype *CodeDMEType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDMETypes, stage.CodeDMEType_stagedOrder, stage.CodeDMEType_orderStaged, &stage.CodeDMETypeOrder, stage.CodeDMETypes_mapString, codedmetype, order, codedmetype.Name)
}

// Unstage removes codedmetype off the model stage
func (codedmetype *CodeDMEType) Unstage(stage *Stage) *CodeDMEType {
	__gong__unstage(stage.CodeDMETypes, stage.CodeDMETypes_mapString, codedmetype, codedmetype.Name)
	return codedmetype
}

// UnstageVoid removes codedmetype off the model stage
func (codedmetype *CodeDMEType) UnstageVoid(stage *Stage) {
	codedmetype.Unstage(stage)
}

func (codedmetype *CodeDMEType) StageVoid(stage *Stage) {
	codedmetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedmetype *CodeDMEType) GetName() (res string) {
	return codedmetype.Name
}

// for satisfaction of GongStruct interface
func (codedmetype *CodeDMEType) SetName(name string) {
	codedmetype.Name = name
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

// Stage puts codedesignatedpointtype to the model stage
func (codedesignatedpointtype *CodeDesignatedPointType) Stage(stage *Stage) *CodeDesignatedPointType {
	__gong__stage(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointType_stagedOrder, stage.CodeDesignatedPointType_orderStaged, &stage.CodeDesignatedPointTypeOrder, stage.CodeDesignatedPointTypes_mapString, codedesignatedpointtype, codedesignatedpointtype.Name)
	return codedesignatedpointtype
}

// StagePreserveOrder puts codedesignatedpointtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDesignatedPointTypeOrder
// - update stage.CodeDesignatedPointTypeOrder accordingly
func (codedesignatedpointtype *CodeDesignatedPointType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointType_stagedOrder, stage.CodeDesignatedPointType_orderStaged, &stage.CodeDesignatedPointTypeOrder, stage.CodeDesignatedPointTypes_mapString, codedesignatedpointtype, order, codedesignatedpointtype.Name)
}

// Unstage removes codedesignatedpointtype off the model stage
func (codedesignatedpointtype *CodeDesignatedPointType) Unstage(stage *Stage) *CodeDesignatedPointType {
	__gong__unstage(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointTypes_mapString, codedesignatedpointtype, codedesignatedpointtype.Name)
	return codedesignatedpointtype
}

// UnstageVoid removes codedesignatedpointtype off the model stage
func (codedesignatedpointtype *CodeDesignatedPointType) UnstageVoid(stage *Stage) {
	codedesignatedpointtype.Unstage(stage)
}

func (codedesignatedpointtype *CodeDesignatedPointType) StageVoid(stage *Stage) {
	codedesignatedpointtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedesignatedpointtype *CodeDesignatedPointType) GetName() (res string) {
	return codedesignatedpointtype.Name
}

// for satisfaction of GongStruct interface
func (codedesignatedpointtype *CodeDesignatedPointType) SetName(name string) {
	codedesignatedpointtype.Name = name
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

// Stage puts codedistanceindicationtype to the model stage
func (codedistanceindicationtype *CodeDistanceIndicationType) Stage(stage *Stage) *CodeDistanceIndicationType {
	__gong__stage(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationType_stagedOrder, stage.CodeDistanceIndicationType_orderStaged, &stage.CodeDistanceIndicationTypeOrder, stage.CodeDistanceIndicationTypes_mapString, codedistanceindicationtype, codedistanceindicationtype.Name)
	return codedistanceindicationtype
}

// StagePreserveOrder puts codedistanceindicationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDistanceIndicationTypeOrder
// - update stage.CodeDistanceIndicationTypeOrder accordingly
func (codedistanceindicationtype *CodeDistanceIndicationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationType_stagedOrder, stage.CodeDistanceIndicationType_orderStaged, &stage.CodeDistanceIndicationTypeOrder, stage.CodeDistanceIndicationTypes_mapString, codedistanceindicationtype, order, codedistanceindicationtype.Name)
}

// Unstage removes codedistanceindicationtype off the model stage
func (codedistanceindicationtype *CodeDistanceIndicationType) Unstage(stage *Stage) *CodeDistanceIndicationType {
	__gong__unstage(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationTypes_mapString, codedistanceindicationtype, codedistanceindicationtype.Name)
	return codedistanceindicationtype
}

// UnstageVoid removes codedistanceindicationtype off the model stage
func (codedistanceindicationtype *CodeDistanceIndicationType) UnstageVoid(stage *Stage) {
	codedistanceindicationtype.Unstage(stage)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) StageVoid(stage *Stage) {
	codedistanceindicationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedistanceindicationtype *CodeDistanceIndicationType) GetName() (res string) {
	return codedistanceindicationtype.Name
}

// for satisfaction of GongStruct interface
func (codedistanceindicationtype *CodeDistanceIndicationType) SetName(name string) {
	codedistanceindicationtype.Name = name
}

// Stage puts codeemissionbandtype to the model stage
func (codeemissionbandtype *CodeEmissionBandType) Stage(stage *Stage) *CodeEmissionBandType {
	__gong__stage(stage.CodeEmissionBandTypes, stage.CodeEmissionBandType_stagedOrder, stage.CodeEmissionBandType_orderStaged, &stage.CodeEmissionBandTypeOrder, stage.CodeEmissionBandTypes_mapString, codeemissionbandtype, codeemissionbandtype.Name)
	return codeemissionbandtype
}

// StagePreserveOrder puts codeemissionbandtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeEmissionBandTypeOrder
// - update stage.CodeEmissionBandTypeOrder accordingly
func (codeemissionbandtype *CodeEmissionBandType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeEmissionBandTypes, stage.CodeEmissionBandType_stagedOrder, stage.CodeEmissionBandType_orderStaged, &stage.CodeEmissionBandTypeOrder, stage.CodeEmissionBandTypes_mapString, codeemissionbandtype, order, codeemissionbandtype.Name)
}

// Unstage removes codeemissionbandtype off the model stage
func (codeemissionbandtype *CodeEmissionBandType) Unstage(stage *Stage) *CodeEmissionBandType {
	__gong__unstage(stage.CodeEmissionBandTypes, stage.CodeEmissionBandTypes_mapString, codeemissionbandtype, codeemissionbandtype.Name)
	return codeemissionbandtype
}

// UnstageVoid removes codeemissionbandtype off the model stage
func (codeemissionbandtype *CodeEmissionBandType) UnstageVoid(stage *Stage) {
	codeemissionbandtype.Unstage(stage)
}

func (codeemissionbandtype *CodeEmissionBandType) StageVoid(stage *Stage) {
	codeemissionbandtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeemissionbandtype *CodeEmissionBandType) GetName() (res string) {
	return codeemissionbandtype.Name
}

// for satisfaction of GongStruct interface
func (codeemissionbandtype *CodeEmissionBandType) SetName(name string) {
	codeemissionbandtype.Name = name
}

// Stage puts codeilsbackcoursetype to the model stage
func (codeilsbackcoursetype *CodeILSBackCourseType) Stage(stage *Stage) *CodeILSBackCourseType {
	__gong__stage(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseType_stagedOrder, stage.CodeILSBackCourseType_orderStaged, &stage.CodeILSBackCourseTypeOrder, stage.CodeILSBackCourseTypes_mapString, codeilsbackcoursetype, codeilsbackcoursetype.Name)
	return codeilsbackcoursetype
}

// StagePreserveOrder puts codeilsbackcoursetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeILSBackCourseTypeOrder
// - update stage.CodeILSBackCourseTypeOrder accordingly
func (codeilsbackcoursetype *CodeILSBackCourseType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseType_stagedOrder, stage.CodeILSBackCourseType_orderStaged, &stage.CodeILSBackCourseTypeOrder, stage.CodeILSBackCourseTypes_mapString, codeilsbackcoursetype, order, codeilsbackcoursetype.Name)
}

// Unstage removes codeilsbackcoursetype off the model stage
func (codeilsbackcoursetype *CodeILSBackCourseType) Unstage(stage *Stage) *CodeILSBackCourseType {
	__gong__unstage(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseTypes_mapString, codeilsbackcoursetype, codeilsbackcoursetype.Name)
	return codeilsbackcoursetype
}

// UnstageVoid removes codeilsbackcoursetype off the model stage
func (codeilsbackcoursetype *CodeILSBackCourseType) UnstageVoid(stage *Stage) {
	codeilsbackcoursetype.Unstage(stage)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) StageVoid(stage *Stage) {
	codeilsbackcoursetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeilsbackcoursetype *CodeILSBackCourseType) GetName() (res string) {
	return codeilsbackcoursetype.Name
}

// for satisfaction of GongStruct interface
func (codeilsbackcoursetype *CodeILSBackCourseType) SetName(name string) {
	codeilsbackcoursetype.Name = name
}

// Stage puts codeintegritylevelilstype to the model stage
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) Stage(stage *Stage) *CodeIntegrityLevelILSType {
	__gong__stage(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSType_stagedOrder, stage.CodeIntegrityLevelILSType_orderStaged, &stage.CodeIntegrityLevelILSTypeOrder, stage.CodeIntegrityLevelILSTypes_mapString, codeintegritylevelilstype, codeintegritylevelilstype.Name)
	return codeintegritylevelilstype
}

// StagePreserveOrder puts codeintegritylevelilstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeIntegrityLevelILSTypeOrder
// - update stage.CodeIntegrityLevelILSTypeOrder accordingly
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSType_stagedOrder, stage.CodeIntegrityLevelILSType_orderStaged, &stage.CodeIntegrityLevelILSTypeOrder, stage.CodeIntegrityLevelILSTypes_mapString, codeintegritylevelilstype, order, codeintegritylevelilstype.Name)
}

// Unstage removes codeintegritylevelilstype off the model stage
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) Unstage(stage *Stage) *CodeIntegrityLevelILSType {
	__gong__unstage(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSTypes_mapString, codeintegritylevelilstype, codeintegritylevelilstype.Name)
	return codeintegritylevelilstype
}

// UnstageVoid removes codeintegritylevelilstype off the model stage
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) UnstageVoid(stage *Stage) {
	codeintegritylevelilstype.Unstage(stage)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) StageVoid(stage *Stage) {
	codeintegritylevelilstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GetName() (res string) {
	return codeintegritylevelilstype.Name
}

// for satisfaction of GongStruct interface
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) SetName(name string) {
	codeintegritylevelilstype.Name = name
}

// Stage puts codemlsazimuthtype to the model stage
func (codemlsazimuthtype *CodeMLSAzimuthType) Stage(stage *Stage) *CodeMLSAzimuthType {
	__gong__stage(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthType_stagedOrder, stage.CodeMLSAzimuthType_orderStaged, &stage.CodeMLSAzimuthTypeOrder, stage.CodeMLSAzimuthTypes_mapString, codemlsazimuthtype, codemlsazimuthtype.Name)
	return codemlsazimuthtype
}

// StagePreserveOrder puts codemlsazimuthtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMLSAzimuthTypeOrder
// - update stage.CodeMLSAzimuthTypeOrder accordingly
func (codemlsazimuthtype *CodeMLSAzimuthType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthType_stagedOrder, stage.CodeMLSAzimuthType_orderStaged, &stage.CodeMLSAzimuthTypeOrder, stage.CodeMLSAzimuthTypes_mapString, codemlsazimuthtype, order, codemlsazimuthtype.Name)
}

// Unstage removes codemlsazimuthtype off the model stage
func (codemlsazimuthtype *CodeMLSAzimuthType) Unstage(stage *Stage) *CodeMLSAzimuthType {
	__gong__unstage(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthTypes_mapString, codemlsazimuthtype, codemlsazimuthtype.Name)
	return codemlsazimuthtype
}

// UnstageVoid removes codemlsazimuthtype off the model stage
func (codemlsazimuthtype *CodeMLSAzimuthType) UnstageVoid(stage *Stage) {
	codemlsazimuthtype.Unstage(stage)
}

func (codemlsazimuthtype *CodeMLSAzimuthType) StageVoid(stage *Stage) {
	codemlsazimuthtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemlsazimuthtype *CodeMLSAzimuthType) GetName() (res string) {
	return codemlsazimuthtype.Name
}

// for satisfaction of GongStruct interface
func (codemlsazimuthtype *CodeMLSAzimuthType) SetName(name string) {
	codemlsazimuthtype.Name = name
}

// Stage puts codemlschanneltype to the model stage
func (codemlschanneltype *CodeMLSChannelType) Stage(stage *Stage) *CodeMLSChannelType {
	__gong__stage(stage.CodeMLSChannelTypes, stage.CodeMLSChannelType_stagedOrder, stage.CodeMLSChannelType_orderStaged, &stage.CodeMLSChannelTypeOrder, stage.CodeMLSChannelTypes_mapString, codemlschanneltype, codemlschanneltype.Name)
	return codemlschanneltype
}

// StagePreserveOrder puts codemlschanneltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMLSChannelTypeOrder
// - update stage.CodeMLSChannelTypeOrder accordingly
func (codemlschanneltype *CodeMLSChannelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMLSChannelTypes, stage.CodeMLSChannelType_stagedOrder, stage.CodeMLSChannelType_orderStaged, &stage.CodeMLSChannelTypeOrder, stage.CodeMLSChannelTypes_mapString, codemlschanneltype, order, codemlschanneltype.Name)
}

// Unstage removes codemlschanneltype off the model stage
func (codemlschanneltype *CodeMLSChannelType) Unstage(stage *Stage) *CodeMLSChannelType {
	__gong__unstage(stage.CodeMLSChannelTypes, stage.CodeMLSChannelTypes_mapString, codemlschanneltype, codemlschanneltype.Name)
	return codemlschanneltype
}

// UnstageVoid removes codemlschanneltype off the model stage
func (codemlschanneltype *CodeMLSChannelType) UnstageVoid(stage *Stage) {
	codemlschanneltype.Unstage(stage)
}

func (codemlschanneltype *CodeMLSChannelType) StageVoid(stage *Stage) {
	codemlschanneltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemlschanneltype *CodeMLSChannelType) GetName() (res string) {
	return codemlschanneltype.Name
}

// for satisfaction of GongStruct interface
func (codemlschanneltype *CodeMLSChannelType) SetName(name string) {
	codemlschanneltype.Name = name
}

// Stage puts codemarkerbeaconsignaltype to the model stage
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) Stage(stage *Stage) *CodeMarkerBeaconSignalType {
	__gong__stage(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalType_stagedOrder, stage.CodeMarkerBeaconSignalType_orderStaged, &stage.CodeMarkerBeaconSignalTypeOrder, stage.CodeMarkerBeaconSignalTypes_mapString, codemarkerbeaconsignaltype, codemarkerbeaconsignaltype.Name)
	return codemarkerbeaconsignaltype
}

// StagePreserveOrder puts codemarkerbeaconsignaltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMarkerBeaconSignalTypeOrder
// - update stage.CodeMarkerBeaconSignalTypeOrder accordingly
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalType_stagedOrder, stage.CodeMarkerBeaconSignalType_orderStaged, &stage.CodeMarkerBeaconSignalTypeOrder, stage.CodeMarkerBeaconSignalTypes_mapString, codemarkerbeaconsignaltype, order, codemarkerbeaconsignaltype.Name)
}

// Unstage removes codemarkerbeaconsignaltype off the model stage
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) Unstage(stage *Stage) *CodeMarkerBeaconSignalType {
	__gong__unstage(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalTypes_mapString, codemarkerbeaconsignaltype, codemarkerbeaconsignaltype.Name)
	return codemarkerbeaconsignaltype
}

// UnstageVoid removes codemarkerbeaconsignaltype off the model stage
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) UnstageVoid(stage *Stage) {
	codemarkerbeaconsignaltype.Unstage(stage)
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) StageVoid(stage *Stage) {
	codemarkerbeaconsignaltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GetName() (res string) {
	return codemarkerbeaconsignaltype.Name
}

// for satisfaction of GongStruct interface
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) SetName(name string) {
	codemarkerbeaconsignaltype.Name = name
}

// Stage puts codendbusagetype to the model stage
func (codendbusagetype *CodeNDBUsageType) Stage(stage *Stage) *CodeNDBUsageType {
	__gong__stage(stage.CodeNDBUsageTypes, stage.CodeNDBUsageType_stagedOrder, stage.CodeNDBUsageType_orderStaged, &stage.CodeNDBUsageTypeOrder, stage.CodeNDBUsageTypes_mapString, codendbusagetype, codendbusagetype.Name)
	return codendbusagetype
}

// StagePreserveOrder puts codendbusagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNDBUsageTypeOrder
// - update stage.CodeNDBUsageTypeOrder accordingly
func (codendbusagetype *CodeNDBUsageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNDBUsageTypes, stage.CodeNDBUsageType_stagedOrder, stage.CodeNDBUsageType_orderStaged, &stage.CodeNDBUsageTypeOrder, stage.CodeNDBUsageTypes_mapString, codendbusagetype, order, codendbusagetype.Name)
}

// Unstage removes codendbusagetype off the model stage
func (codendbusagetype *CodeNDBUsageType) Unstage(stage *Stage) *CodeNDBUsageType {
	__gong__unstage(stage.CodeNDBUsageTypes, stage.CodeNDBUsageTypes_mapString, codendbusagetype, codendbusagetype.Name)
	return codendbusagetype
}

// UnstageVoid removes codendbusagetype off the model stage
func (codendbusagetype *CodeNDBUsageType) UnstageVoid(stage *Stage) {
	codendbusagetype.Unstage(stage)
}

func (codendbusagetype *CodeNDBUsageType) StageVoid(stage *Stage) {
	codendbusagetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codendbusagetype *CodeNDBUsageType) GetName() (res string) {
	return codendbusagetype.Name
}

// for satisfaction of GongStruct interface
func (codendbusagetype *CodeNDBUsageType) SetName(name string) {
	codendbusagetype.Name = name
}

// Stage puts codenavaiddesignatortype to the model stage
func (codenavaiddesignatortype *CodeNavaidDesignatorType) Stage(stage *Stage) *CodeNavaidDesignatorType {
	__gong__stage(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorType_stagedOrder, stage.CodeNavaidDesignatorType_orderStaged, &stage.CodeNavaidDesignatorTypeOrder, stage.CodeNavaidDesignatorTypes_mapString, codenavaiddesignatortype, codenavaiddesignatortype.Name)
	return codenavaiddesignatortype
}

// StagePreserveOrder puts codenavaiddesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavaidDesignatorTypeOrder
// - update stage.CodeNavaidDesignatorTypeOrder accordingly
func (codenavaiddesignatortype *CodeNavaidDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorType_stagedOrder, stage.CodeNavaidDesignatorType_orderStaged, &stage.CodeNavaidDesignatorTypeOrder, stage.CodeNavaidDesignatorTypes_mapString, codenavaiddesignatortype, order, codenavaiddesignatortype.Name)
}

// Unstage removes codenavaiddesignatortype off the model stage
func (codenavaiddesignatortype *CodeNavaidDesignatorType) Unstage(stage *Stage) *CodeNavaidDesignatorType {
	__gong__unstage(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorTypes_mapString, codenavaiddesignatortype, codenavaiddesignatortype.Name)
	return codenavaiddesignatortype
}

// UnstageVoid removes codenavaiddesignatortype off the model stage
func (codenavaiddesignatortype *CodeNavaidDesignatorType) UnstageVoid(stage *Stage) {
	codenavaiddesignatortype.Unstage(stage)
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) StageVoid(stage *Stage) {
	codenavaiddesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavaiddesignatortype *CodeNavaidDesignatorType) GetName() (res string) {
	return codenavaiddesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codenavaiddesignatortype *CodeNavaidDesignatorType) SetName(name string) {
	codenavaiddesignatortype.Name = name
}

// Stage puts codenavaidpurposetype to the model stage
func (codenavaidpurposetype *CodeNavaidPurposeType) Stage(stage *Stage) *CodeNavaidPurposeType {
	__gong__stage(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeType_stagedOrder, stage.CodeNavaidPurposeType_orderStaged, &stage.CodeNavaidPurposeTypeOrder, stage.CodeNavaidPurposeTypes_mapString, codenavaidpurposetype, codenavaidpurposetype.Name)
	return codenavaidpurposetype
}

// StagePreserveOrder puts codenavaidpurposetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavaidPurposeTypeOrder
// - update stage.CodeNavaidPurposeTypeOrder accordingly
func (codenavaidpurposetype *CodeNavaidPurposeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeType_stagedOrder, stage.CodeNavaidPurposeType_orderStaged, &stage.CodeNavaidPurposeTypeOrder, stage.CodeNavaidPurposeTypes_mapString, codenavaidpurposetype, order, codenavaidpurposetype.Name)
}

// Unstage removes codenavaidpurposetype off the model stage
func (codenavaidpurposetype *CodeNavaidPurposeType) Unstage(stage *Stage) *CodeNavaidPurposeType {
	__gong__unstage(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeTypes_mapString, codenavaidpurposetype, codenavaidpurposetype.Name)
	return codenavaidpurposetype
}

// UnstageVoid removes codenavaidpurposetype off the model stage
func (codenavaidpurposetype *CodeNavaidPurposeType) UnstageVoid(stage *Stage) {
	codenavaidpurposetype.Unstage(stage)
}

func (codenavaidpurposetype *CodeNavaidPurposeType) StageVoid(stage *Stage) {
	codenavaidpurposetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavaidpurposetype *CodeNavaidPurposeType) GetName() (res string) {
	return codenavaidpurposetype.Name
}

// for satisfaction of GongStruct interface
func (codenavaidpurposetype *CodeNavaidPurposeType) SetName(name string) {
	codenavaidpurposetype.Name = name
}

// Stage puts codenavaidservicetype to the model stage
func (codenavaidservicetype *CodeNavaidServiceType) Stage(stage *Stage) *CodeNavaidServiceType {
	__gong__stage(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceType_stagedOrder, stage.CodeNavaidServiceType_orderStaged, &stage.CodeNavaidServiceTypeOrder, stage.CodeNavaidServiceTypes_mapString, codenavaidservicetype, codenavaidservicetype.Name)
	return codenavaidservicetype
}

// StagePreserveOrder puts codenavaidservicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNavaidServiceTypeOrder
// - update stage.CodeNavaidServiceTypeOrder accordingly
func (codenavaidservicetype *CodeNavaidServiceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceType_stagedOrder, stage.CodeNavaidServiceType_orderStaged, &stage.CodeNavaidServiceTypeOrder, stage.CodeNavaidServiceTypes_mapString, codenavaidservicetype, order, codenavaidservicetype.Name)
}

// Unstage removes codenavaidservicetype off the model stage
func (codenavaidservicetype *CodeNavaidServiceType) Unstage(stage *Stage) *CodeNavaidServiceType {
	__gong__unstage(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceTypes_mapString, codenavaidservicetype, codenavaidservicetype.Name)
	return codenavaidservicetype
}

// UnstageVoid removes codenavaidservicetype off the model stage
func (codenavaidservicetype *CodeNavaidServiceType) UnstageVoid(stage *Stage) {
	codenavaidservicetype.Unstage(stage)
}

func (codenavaidservicetype *CodeNavaidServiceType) StageVoid(stage *Stage) {
	codenavaidservicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenavaidservicetype *CodeNavaidServiceType) GetName() (res string) {
	return codenavaidservicetype.Name
}

// for satisfaction of GongStruct interface
func (codenavaidservicetype *CodeNavaidServiceType) SetName(name string) {
	codenavaidservicetype.Name = name
}

// Stage puts codenorthreferencetype to the model stage
func (codenorthreferencetype *CodeNorthReferenceType) Stage(stage *Stage) *CodeNorthReferenceType {
	__gong__stage(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceType_stagedOrder, stage.CodeNorthReferenceType_orderStaged, &stage.CodeNorthReferenceTypeOrder, stage.CodeNorthReferenceTypes_mapString, codenorthreferencetype, codenorthreferencetype.Name)
	return codenorthreferencetype
}

// StagePreserveOrder puts codenorthreferencetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNorthReferenceTypeOrder
// - update stage.CodeNorthReferenceTypeOrder accordingly
func (codenorthreferencetype *CodeNorthReferenceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceType_stagedOrder, stage.CodeNorthReferenceType_orderStaged, &stage.CodeNorthReferenceTypeOrder, stage.CodeNorthReferenceTypes_mapString, codenorthreferencetype, order, codenorthreferencetype.Name)
}

// Unstage removes codenorthreferencetype off the model stage
func (codenorthreferencetype *CodeNorthReferenceType) Unstage(stage *Stage) *CodeNorthReferenceType {
	__gong__unstage(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceTypes_mapString, codenorthreferencetype, codenorthreferencetype.Name)
	return codenorthreferencetype
}

// UnstageVoid removes codenorthreferencetype off the model stage
func (codenorthreferencetype *CodeNorthReferenceType) UnstageVoid(stage *Stage) {
	codenorthreferencetype.Unstage(stage)
}

func (codenorthreferencetype *CodeNorthReferenceType) StageVoid(stage *Stage) {
	codenorthreferencetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenorthreferencetype *CodeNorthReferenceType) GetName() (res string) {
	return codenorthreferencetype.Name
}

// for satisfaction of GongStruct interface
func (codenorthreferencetype *CodeNorthReferenceType) SetName(name string) {
	codenorthreferencetype.Name = name
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

// Stage puts codepartype to the model stage
func (codepartype *CodePARType) Stage(stage *Stage) *CodePARType {
	__gong__stage(stage.CodePARTypes, stage.CodePARType_stagedOrder, stage.CodePARType_orderStaged, &stage.CodePARTypeOrder, stage.CodePARTypes_mapString, codepartype, codepartype.Name)
	return codepartype
}

// StagePreserveOrder puts codepartype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePARTypeOrder
// - update stage.CodePARTypeOrder accordingly
func (codepartype *CodePARType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePARTypes, stage.CodePARType_stagedOrder, stage.CodePARType_orderStaged, &stage.CodePARTypeOrder, stage.CodePARTypes_mapString, codepartype, order, codepartype.Name)
}

// Unstage removes codepartype off the model stage
func (codepartype *CodePARType) Unstage(stage *Stage) *CodePARType {
	__gong__unstage(stage.CodePARTypes, stage.CodePARTypes_mapString, codepartype, codepartype.Name)
	return codepartype
}

// UnstageVoid removes codepartype off the model stage
func (codepartype *CodePARType) UnstageVoid(stage *Stage) {
	codepartype.Unstage(stage)
}

func (codepartype *CodePARType) StageVoid(stage *Stage) {
	codepartype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepartype *CodePARType) GetName() (res string) {
	return codepartype.Name
}

// for satisfaction of GongStruct interface
func (codepartype *CodePARType) SetName(name string) {
	codepartype.Name = name
}

// Stage puts codepositioninilstype to the model stage
func (codepositioninilstype *CodePositionInILSType) Stage(stage *Stage) *CodePositionInILSType {
	__gong__stage(stage.CodePositionInILSTypes, stage.CodePositionInILSType_stagedOrder, stage.CodePositionInILSType_orderStaged, &stage.CodePositionInILSTypeOrder, stage.CodePositionInILSTypes_mapString, codepositioninilstype, codepositioninilstype.Name)
	return codepositioninilstype
}

// StagePreserveOrder puts codepositioninilstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePositionInILSTypeOrder
// - update stage.CodePositionInILSTypeOrder accordingly
func (codepositioninilstype *CodePositionInILSType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePositionInILSTypes, stage.CodePositionInILSType_stagedOrder, stage.CodePositionInILSType_orderStaged, &stage.CodePositionInILSTypeOrder, stage.CodePositionInILSTypes_mapString, codepositioninilstype, order, codepositioninilstype.Name)
}

// Unstage removes codepositioninilstype off the model stage
func (codepositioninilstype *CodePositionInILSType) Unstage(stage *Stage) *CodePositionInILSType {
	__gong__unstage(stage.CodePositionInILSTypes, stage.CodePositionInILSTypes_mapString, codepositioninilstype, codepositioninilstype.Name)
	return codepositioninilstype
}

// UnstageVoid removes codepositioninilstype off the model stage
func (codepositioninilstype *CodePositionInILSType) UnstageVoid(stage *Stage) {
	codepositioninilstype.Unstage(stage)
}

func (codepositioninilstype *CodePositionInILSType) StageVoid(stage *Stage) {
	codepositioninilstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepositioninilstype *CodePositionInILSType) GetName() (res string) {
	return codepositioninilstype.Name
}

// for satisfaction of GongStruct interface
func (codepositioninilstype *CodePositionInILSType) SetName(name string) {
	codepositioninilstype.Name = name
}

// Stage puts codeprimaryradartype to the model stage
func (codeprimaryradartype *CodePrimaryRadarType) Stage(stage *Stage) *CodePrimaryRadarType {
	__gong__stage(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarType_stagedOrder, stage.CodePrimaryRadarType_orderStaged, &stage.CodePrimaryRadarTypeOrder, stage.CodePrimaryRadarTypes_mapString, codeprimaryradartype, codeprimaryradartype.Name)
	return codeprimaryradartype
}

// StagePreserveOrder puts codeprimaryradartype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePrimaryRadarTypeOrder
// - update stage.CodePrimaryRadarTypeOrder accordingly
func (codeprimaryradartype *CodePrimaryRadarType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarType_stagedOrder, stage.CodePrimaryRadarType_orderStaged, &stage.CodePrimaryRadarTypeOrder, stage.CodePrimaryRadarTypes_mapString, codeprimaryradartype, order, codeprimaryradartype.Name)
}

// Unstage removes codeprimaryradartype off the model stage
func (codeprimaryradartype *CodePrimaryRadarType) Unstage(stage *Stage) *CodePrimaryRadarType {
	__gong__unstage(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarTypes_mapString, codeprimaryradartype, codeprimaryradartype.Name)
	return codeprimaryradartype
}

// UnstageVoid removes codeprimaryradartype off the model stage
func (codeprimaryradartype *CodePrimaryRadarType) UnstageVoid(stage *Stage) {
	codeprimaryradartype.Unstage(stage)
}

func (codeprimaryradartype *CodePrimaryRadarType) StageVoid(stage *Stage) {
	codeprimaryradartype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprimaryradartype *CodePrimaryRadarType) GetName() (res string) {
	return codeprimaryradartype.Name
}

// for satisfaction of GongStruct interface
func (codeprimaryradartype *CodePrimaryRadarType) SetName(name string) {
	codeprimaryradartype.Name = name
}

// Stage puts coderadarservicetype to the model stage
func (coderadarservicetype *CodeRadarServiceType) Stage(stage *Stage) *CodeRadarServiceType {
	__gong__stage(stage.CodeRadarServiceTypes, stage.CodeRadarServiceType_stagedOrder, stage.CodeRadarServiceType_orderStaged, &stage.CodeRadarServiceTypeOrder, stage.CodeRadarServiceTypes_mapString, coderadarservicetype, coderadarservicetype.Name)
	return coderadarservicetype
}

// StagePreserveOrder puts coderadarservicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRadarServiceTypeOrder
// - update stage.CodeRadarServiceTypeOrder accordingly
func (coderadarservicetype *CodeRadarServiceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRadarServiceTypes, stage.CodeRadarServiceType_stagedOrder, stage.CodeRadarServiceType_orderStaged, &stage.CodeRadarServiceTypeOrder, stage.CodeRadarServiceTypes_mapString, coderadarservicetype, order, coderadarservicetype.Name)
}

// Unstage removes coderadarservicetype off the model stage
func (coderadarservicetype *CodeRadarServiceType) Unstage(stage *Stage) *CodeRadarServiceType {
	__gong__unstage(stage.CodeRadarServiceTypes, stage.CodeRadarServiceTypes_mapString, coderadarservicetype, coderadarservicetype.Name)
	return coderadarservicetype
}

// UnstageVoid removes coderadarservicetype off the model stage
func (coderadarservicetype *CodeRadarServiceType) UnstageVoid(stage *Stage) {
	coderadarservicetype.Unstage(stage)
}

func (coderadarservicetype *CodeRadarServiceType) StageVoid(stage *Stage) {
	coderadarservicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderadarservicetype *CodeRadarServiceType) GetName() (res string) {
	return coderadarservicetype.Name
}

// for satisfaction of GongStruct interface
func (coderadarservicetype *CodeRadarServiceType) SetName(name string) {
	coderadarservicetype.Name = name
}

// Stage puts coderadioemissiontype to the model stage
func (coderadioemissiontype *CodeRadioEmissionType) Stage(stage *Stage) *CodeRadioEmissionType {
	__gong__stage(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionType_stagedOrder, stage.CodeRadioEmissionType_orderStaged, &stage.CodeRadioEmissionTypeOrder, stage.CodeRadioEmissionTypes_mapString, coderadioemissiontype, coderadioemissiontype.Name)
	return coderadioemissiontype
}

// StagePreserveOrder puts coderadioemissiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRadioEmissionTypeOrder
// - update stage.CodeRadioEmissionTypeOrder accordingly
func (coderadioemissiontype *CodeRadioEmissionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionType_stagedOrder, stage.CodeRadioEmissionType_orderStaged, &stage.CodeRadioEmissionTypeOrder, stage.CodeRadioEmissionTypes_mapString, coderadioemissiontype, order, coderadioemissiontype.Name)
}

// Unstage removes coderadioemissiontype off the model stage
func (coderadioemissiontype *CodeRadioEmissionType) Unstage(stage *Stage) *CodeRadioEmissionType {
	__gong__unstage(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionTypes_mapString, coderadioemissiontype, coderadioemissiontype.Name)
	return coderadioemissiontype
}

// UnstageVoid removes coderadioemissiontype off the model stage
func (coderadioemissiontype *CodeRadioEmissionType) UnstageVoid(stage *Stage) {
	coderadioemissiontype.Unstage(stage)
}

func (coderadioemissiontype *CodeRadioEmissionType) StageVoid(stage *Stage) {
	coderadioemissiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderadioemissiontype *CodeRadioEmissionType) GetName() (res string) {
	return coderadioemissiontype.Name
}

// for satisfaction of GongStruct interface
func (coderadioemissiontype *CodeRadioEmissionType) SetName(name string) {
	coderadioemissiontype.Name = name
}

// Stage puts coderadiosignaltype to the model stage
func (coderadiosignaltype *CodeRadioSignalType) Stage(stage *Stage) *CodeRadioSignalType {
	__gong__stage(stage.CodeRadioSignalTypes, stage.CodeRadioSignalType_stagedOrder, stage.CodeRadioSignalType_orderStaged, &stage.CodeRadioSignalTypeOrder, stage.CodeRadioSignalTypes_mapString, coderadiosignaltype, coderadiosignaltype.Name)
	return coderadiosignaltype
}

// StagePreserveOrder puts coderadiosignaltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRadioSignalTypeOrder
// - update stage.CodeRadioSignalTypeOrder accordingly
func (coderadiosignaltype *CodeRadioSignalType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRadioSignalTypes, stage.CodeRadioSignalType_stagedOrder, stage.CodeRadioSignalType_orderStaged, &stage.CodeRadioSignalTypeOrder, stage.CodeRadioSignalTypes_mapString, coderadiosignaltype, order, coderadiosignaltype.Name)
}

// Unstage removes coderadiosignaltype off the model stage
func (coderadiosignaltype *CodeRadioSignalType) Unstage(stage *Stage) *CodeRadioSignalType {
	__gong__unstage(stage.CodeRadioSignalTypes, stage.CodeRadioSignalTypes_mapString, coderadiosignaltype, coderadiosignaltype.Name)
	return coderadiosignaltype
}

// UnstageVoid removes coderadiosignaltype off the model stage
func (coderadiosignaltype *CodeRadioSignalType) UnstageVoid(stage *Stage) {
	coderadiosignaltype.Unstage(stage)
}

func (coderadiosignaltype *CodeRadioSignalType) StageVoid(stage *Stage) {
	coderadiosignaltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderadiosignaltype *CodeRadioSignalType) GetName() (res string) {
	return coderadiosignaltype.Name
}

// for satisfaction of GongStruct interface
func (coderadiosignaltype *CodeRadioSignalType) SetName(name string) {
	coderadiosignaltype.Name = name
}

// Stage puts codereflectortype to the model stage
func (codereflectortype *CodeReflectorType) Stage(stage *Stage) *CodeReflectorType {
	__gong__stage(stage.CodeReflectorTypes, stage.CodeReflectorType_stagedOrder, stage.CodeReflectorType_orderStaged, &stage.CodeReflectorTypeOrder, stage.CodeReflectorTypes_mapString, codereflectortype, codereflectortype.Name)
	return codereflectortype
}

// StagePreserveOrder puts codereflectortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeReflectorTypeOrder
// - update stage.CodeReflectorTypeOrder accordingly
func (codereflectortype *CodeReflectorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeReflectorTypes, stage.CodeReflectorType_stagedOrder, stage.CodeReflectorType_orderStaged, &stage.CodeReflectorTypeOrder, stage.CodeReflectorTypes_mapString, codereflectortype, order, codereflectortype.Name)
}

// Unstage removes codereflectortype off the model stage
func (codereflectortype *CodeReflectorType) Unstage(stage *Stage) *CodeReflectorType {
	__gong__unstage(stage.CodeReflectorTypes, stage.CodeReflectorTypes_mapString, codereflectortype, codereflectortype.Name)
	return codereflectortype
}

// UnstageVoid removes codereflectortype off the model stage
func (codereflectortype *CodeReflectorType) UnstageVoid(stage *Stage) {
	codereflectortype.Unstage(stage)
}

func (codereflectortype *CodeReflectorType) StageVoid(stage *Stage) {
	codereflectortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codereflectortype *CodeReflectorType) GetName() (res string) {
	return codereflectortype.Name
}

// for satisfaction of GongStruct interface
func (codereflectortype *CodeReflectorType) SetName(name string) {
	codereflectortype.Name = name
}

// Stage puts codesignalperformanceilstype to the model stage
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) Stage(stage *Stage) *CodeSignalPerformanceILSType {
	__gong__stage(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSType_stagedOrder, stage.CodeSignalPerformanceILSType_orderStaged, &stage.CodeSignalPerformanceILSTypeOrder, stage.CodeSignalPerformanceILSTypes_mapString, codesignalperformanceilstype, codesignalperformanceilstype.Name)
	return codesignalperformanceilstype
}

// StagePreserveOrder puts codesignalperformanceilstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSignalPerformanceILSTypeOrder
// - update stage.CodeSignalPerformanceILSTypeOrder accordingly
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSType_stagedOrder, stage.CodeSignalPerformanceILSType_orderStaged, &stage.CodeSignalPerformanceILSTypeOrder, stage.CodeSignalPerformanceILSTypes_mapString, codesignalperformanceilstype, order, codesignalperformanceilstype.Name)
}

// Unstage removes codesignalperformanceilstype off the model stage
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) Unstage(stage *Stage) *CodeSignalPerformanceILSType {
	__gong__unstage(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSTypes_mapString, codesignalperformanceilstype, codesignalperformanceilstype.Name)
	return codesignalperformanceilstype
}

// UnstageVoid removes codesignalperformanceilstype off the model stage
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) UnstageVoid(stage *Stage) {
	codesignalperformanceilstype.Unstage(stage)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) StageVoid(stage *Stage) {
	codesignalperformanceilstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GetName() (res string) {
	return codesignalperformanceilstype.Name
}

// for satisfaction of GongStruct interface
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) SetName(name string) {
	codesignalperformanceilstype.Name = name
}

// Stage puts codespecialnavigationchaindesignatortype to the model stage
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) Stage(stage *Stage) *CodeSpecialNavigationChainDesignatorType {
	__gong__stage(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorType_stagedOrder, stage.CodeSpecialNavigationChainDesignatorType_orderStaged, &stage.CodeSpecialNavigationChainDesignatorTypeOrder, stage.CodeSpecialNavigationChainDesignatorTypes_mapString, codespecialnavigationchaindesignatortype, codespecialnavigationchaindesignatortype.Name)
	return codespecialnavigationchaindesignatortype
}

// StagePreserveOrder puts codespecialnavigationchaindesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSpecialNavigationChainDesignatorTypeOrder
// - update stage.CodeSpecialNavigationChainDesignatorTypeOrder accordingly
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorType_stagedOrder, stage.CodeSpecialNavigationChainDesignatorType_orderStaged, &stage.CodeSpecialNavigationChainDesignatorTypeOrder, stage.CodeSpecialNavigationChainDesignatorTypes_mapString, codespecialnavigationchaindesignatortype, order, codespecialnavigationchaindesignatortype.Name)
}

// Unstage removes codespecialnavigationchaindesignatortype off the model stage
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) Unstage(stage *Stage) *CodeSpecialNavigationChainDesignatorType {
	__gong__unstage(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorTypes_mapString, codespecialnavigationchaindesignatortype, codespecialnavigationchaindesignatortype.Name)
	return codespecialnavigationchaindesignatortype
}

// UnstageVoid removes codespecialnavigationchaindesignatortype off the model stage
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) UnstageVoid(stage *Stage) {
	codespecialnavigationchaindesignatortype.Unstage(stage)
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) StageVoid(stage *Stage) {
	codespecialnavigationchaindesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GetName() (res string) {
	return codespecialnavigationchaindesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) SetName(name string) {
	codespecialnavigationchaindesignatortype.Name = name
}

// Stage puts codespecialnavigationstationtype to the model stage
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) Stage(stage *Stage) *CodeSpecialNavigationStationType {
	__gong__stage(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationType_stagedOrder, stage.CodeSpecialNavigationStationType_orderStaged, &stage.CodeSpecialNavigationStationTypeOrder, stage.CodeSpecialNavigationStationTypes_mapString, codespecialnavigationstationtype, codespecialnavigationstationtype.Name)
	return codespecialnavigationstationtype
}

// StagePreserveOrder puts codespecialnavigationstationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSpecialNavigationStationTypeOrder
// - update stage.CodeSpecialNavigationStationTypeOrder accordingly
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationType_stagedOrder, stage.CodeSpecialNavigationStationType_orderStaged, &stage.CodeSpecialNavigationStationTypeOrder, stage.CodeSpecialNavigationStationTypes_mapString, codespecialnavigationstationtype, order, codespecialnavigationstationtype.Name)
}

// Unstage removes codespecialnavigationstationtype off the model stage
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) Unstage(stage *Stage) *CodeSpecialNavigationStationType {
	__gong__unstage(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationTypes_mapString, codespecialnavigationstationtype, codespecialnavigationstationtype.Name)
	return codespecialnavigationstationtype
}

// UnstageVoid removes codespecialnavigationstationtype off the model stage
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) UnstageVoid(stage *Stage) {
	codespecialnavigationstationtype.Unstage(stage)
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) StageVoid(stage *Stage) {
	codespecialnavigationstationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GetName() (res string) {
	return codespecialnavigationstationtype.Name
}

// for satisfaction of GongStruct interface
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) SetName(name string) {
	codespecialnavigationstationtype.Name = name
}

// Stage puts codespecialnavigationsystemtype to the model stage
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) Stage(stage *Stage) *CodeSpecialNavigationSystemType {
	__gong__stage(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemType_stagedOrder, stage.CodeSpecialNavigationSystemType_orderStaged, &stage.CodeSpecialNavigationSystemTypeOrder, stage.CodeSpecialNavigationSystemTypes_mapString, codespecialnavigationsystemtype, codespecialnavigationsystemtype.Name)
	return codespecialnavigationsystemtype
}

// StagePreserveOrder puts codespecialnavigationsystemtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSpecialNavigationSystemTypeOrder
// - update stage.CodeSpecialNavigationSystemTypeOrder accordingly
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemType_stagedOrder, stage.CodeSpecialNavigationSystemType_orderStaged, &stage.CodeSpecialNavigationSystemTypeOrder, stage.CodeSpecialNavigationSystemTypes_mapString, codespecialnavigationsystemtype, order, codespecialnavigationsystemtype.Name)
}

// Unstage removes codespecialnavigationsystemtype off the model stage
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) Unstage(stage *Stage) *CodeSpecialNavigationSystemType {
	__gong__unstage(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemTypes_mapString, codespecialnavigationsystemtype, codespecialnavigationsystemtype.Name)
	return codespecialnavigationsystemtype
}

// UnstageVoid removes codespecialnavigationsystemtype off the model stage
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) UnstageVoid(stage *Stage) {
	codespecialnavigationsystemtype.Unstage(stage)
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) StageVoid(stage *Stage) {
	codespecialnavigationsystemtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GetName() (res string) {
	return codespecialnavigationsystemtype.Name
}

// for satisfaction of GongStruct interface
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) SetName(name string) {
	codespecialnavigationsystemtype.Name = name
}

// Stage puts codestandbypowertype to the model stage
func (codestandbypowertype *CodeStandbyPowerType) Stage(stage *Stage) *CodeStandbyPowerType {
	__gong__stage(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerType_stagedOrder, stage.CodeStandbyPowerType_orderStaged, &stage.CodeStandbyPowerTypeOrder, stage.CodeStandbyPowerTypes_mapString, codestandbypowertype, codestandbypowertype.Name)
	return codestandbypowertype
}

// StagePreserveOrder puts codestandbypowertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStandbyPowerTypeOrder
// - update stage.CodeStandbyPowerTypeOrder accordingly
func (codestandbypowertype *CodeStandbyPowerType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerType_stagedOrder, stage.CodeStandbyPowerType_orderStaged, &stage.CodeStandbyPowerTypeOrder, stage.CodeStandbyPowerTypes_mapString, codestandbypowertype, order, codestandbypowertype.Name)
}

// Unstage removes codestandbypowertype off the model stage
func (codestandbypowertype *CodeStandbyPowerType) Unstage(stage *Stage) *CodeStandbyPowerType {
	__gong__unstage(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerTypes_mapString, codestandbypowertype, codestandbypowertype.Name)
	return codestandbypowertype
}

// UnstageVoid removes codestandbypowertype off the model stage
func (codestandbypowertype *CodeStandbyPowerType) UnstageVoid(stage *Stage) {
	codestandbypowertype.Unstage(stage)
}

func (codestandbypowertype *CodeStandbyPowerType) StageVoid(stage *Stage) {
	codestandbypowertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestandbypowertype *CodeStandbyPowerType) GetName() (res string) {
	return codestandbypowertype.Name
}

// for satisfaction of GongStruct interface
func (codestandbypowertype *CodeStandbyPowerType) SetName(name string) {
	codestandbypowertype.Name = name
}

// Stage puts codestatusnavaidtype to the model stage
func (codestatusnavaidtype *CodeStatusNavaidType) Stage(stage *Stage) *CodeStatusNavaidType {
	__gong__stage(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidType_stagedOrder, stage.CodeStatusNavaidType_orderStaged, &stage.CodeStatusNavaidTypeOrder, stage.CodeStatusNavaidTypes_mapString, codestatusnavaidtype, codestatusnavaidtype.Name)
	return codestatusnavaidtype
}

// StagePreserveOrder puts codestatusnavaidtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusNavaidTypeOrder
// - update stage.CodeStatusNavaidTypeOrder accordingly
func (codestatusnavaidtype *CodeStatusNavaidType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidType_stagedOrder, stage.CodeStatusNavaidType_orderStaged, &stage.CodeStatusNavaidTypeOrder, stage.CodeStatusNavaidTypes_mapString, codestatusnavaidtype, order, codestatusnavaidtype.Name)
}

// Unstage removes codestatusnavaidtype off the model stage
func (codestatusnavaidtype *CodeStatusNavaidType) Unstage(stage *Stage) *CodeStatusNavaidType {
	__gong__unstage(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidTypes_mapString, codestatusnavaidtype, codestatusnavaidtype.Name)
	return codestatusnavaidtype
}

// UnstageVoid removes codestatusnavaidtype off the model stage
func (codestatusnavaidtype *CodeStatusNavaidType) UnstageVoid(stage *Stage) {
	codestatusnavaidtype.Unstage(stage)
}

func (codestatusnavaidtype *CodeStatusNavaidType) StageVoid(stage *Stage) {
	codestatusnavaidtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusnavaidtype *CodeStatusNavaidType) GetName() (res string) {
	return codestatusnavaidtype.Name
}

// for satisfaction of GongStruct interface
func (codestatusnavaidtype *CodeStatusNavaidType) SetName(name string) {
	codestatusnavaidtype.Name = name
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

// Stage puts codevortype to the model stage
func (codevortype *CodeVORType) Stage(stage *Stage) *CodeVORType {
	__gong__stage(stage.CodeVORTypes, stage.CodeVORType_stagedOrder, stage.CodeVORType_orderStaged, &stage.CodeVORTypeOrder, stage.CodeVORTypes_mapString, codevortype, codevortype.Name)
	return codevortype
}

// StagePreserveOrder puts codevortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVORTypeOrder
// - update stage.CodeVORTypeOrder accordingly
func (codevortype *CodeVORType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVORTypes, stage.CodeVORType_stagedOrder, stage.CodeVORType_orderStaged, &stage.CodeVORTypeOrder, stage.CodeVORTypes_mapString, codevortype, order, codevortype.Name)
}

// Unstage removes codevortype off the model stage
func (codevortype *CodeVORType) Unstage(stage *Stage) *CodeVORType {
	__gong__unstage(stage.CodeVORTypes, stage.CodeVORTypes_mapString, codevortype, codevortype.Name)
	return codevortype
}

// UnstageVoid removes codevortype off the model stage
func (codevortype *CodeVORType) UnstageVoid(stage *Stage) {
	codevortype.Unstage(stage)
}

func (codevortype *CodeVORType) StageVoid(stage *Stage) {
	codevortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codevortype *CodeVORType) GetName() (res string) {
	return codevortype.Name
}

// for satisfaction of GongStruct interface
func (codevortype *CodeVORType) SetName(name string) {
	codevortype.Name = name
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

// Stage puts dateyeartype to the model stage
func (dateyeartype *DateYearType) Stage(stage *Stage) *DateYearType {
	__gong__stage(stage.DateYearTypes, stage.DateYearType_stagedOrder, stage.DateYearType_orderStaged, &stage.DateYearTypeOrder, stage.DateYearTypes_mapString, dateyeartype, dateyeartype.Name)
	return dateyeartype
}

// StagePreserveOrder puts dateyeartype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DateYearTypeOrder
// - update stage.DateYearTypeOrder accordingly
func (dateyeartype *DateYearType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DateYearTypes, stage.DateYearType_stagedOrder, stage.DateYearType_orderStaged, &stage.DateYearTypeOrder, stage.DateYearTypes_mapString, dateyeartype, order, dateyeartype.Name)
}

// Unstage removes dateyeartype off the model stage
func (dateyeartype *DateYearType) Unstage(stage *Stage) *DateYearType {
	__gong__unstage(stage.DateYearTypes, stage.DateYearTypes_mapString, dateyeartype, dateyeartype.Name)
	return dateyeartype
}

// UnstageVoid removes dateyeartype off the model stage
func (dateyeartype *DateYearType) UnstageVoid(stage *Stage) {
	dateyeartype.Unstage(stage)
}

func (dateyeartype *DateYearType) StageVoid(stage *Stage) {
	dateyeartype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dateyeartype *DateYearType) GetName() (res string) {
	return dateyeartype.Name
}

// for satisfaction of GongStruct interface
func (dateyeartype *DateYearType) SetName(name string) {
	dateyeartype.Name = name
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

// Stage puts valfrequencytype to the model stage
func (valfrequencytype *ValFrequencyType) Stage(stage *Stage) *ValFrequencyType {
	__gong__stage(stage.ValFrequencyTypes, stage.ValFrequencyType_stagedOrder, stage.ValFrequencyType_orderStaged, &stage.ValFrequencyTypeOrder, stage.ValFrequencyTypes_mapString, valfrequencytype, valfrequencytype.Name)
	return valfrequencytype
}

// StagePreserveOrder puts valfrequencytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValFrequencyTypeOrder
// - update stage.ValFrequencyTypeOrder accordingly
func (valfrequencytype *ValFrequencyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValFrequencyTypes, stage.ValFrequencyType_stagedOrder, stage.ValFrequencyType_orderStaged, &stage.ValFrequencyTypeOrder, stage.ValFrequencyTypes_mapString, valfrequencytype, order, valfrequencytype.Name)
}

// Unstage removes valfrequencytype off the model stage
func (valfrequencytype *ValFrequencyType) Unstage(stage *Stage) *ValFrequencyType {
	__gong__unstage(stage.ValFrequencyTypes, stage.ValFrequencyTypes_mapString, valfrequencytype, valfrequencytype.Name)
	return valfrequencytype
}

// UnstageVoid removes valfrequencytype off the model stage
func (valfrequencytype *ValFrequencyType) UnstageVoid(stage *Stage) {
	valfrequencytype.Unstage(stage)
}

func (valfrequencytype *ValFrequencyType) StageVoid(stage *Stage) {
	valfrequencytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valfrequencytype *ValFrequencyType) GetName() (res string) {
	return valfrequencytype.Name
}

// for satisfaction of GongStruct interface
func (valfrequencytype *ValFrequencyType) SetName(name string) {
	valfrequencytype.Name = name
}

// Stage puts valmagneticvariationtype to the model stage
func (valmagneticvariationtype *ValMagneticVariationType) Stage(stage *Stage) *ValMagneticVariationType {
	__gong__stage(stage.ValMagneticVariationTypes, stage.ValMagneticVariationType_stagedOrder, stage.ValMagneticVariationType_orderStaged, &stage.ValMagneticVariationTypeOrder, stage.ValMagneticVariationTypes_mapString, valmagneticvariationtype, valmagneticvariationtype.Name)
	return valmagneticvariationtype
}

// StagePreserveOrder puts valmagneticvariationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValMagneticVariationTypeOrder
// - update stage.ValMagneticVariationTypeOrder accordingly
func (valmagneticvariationtype *ValMagneticVariationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValMagneticVariationTypes, stage.ValMagneticVariationType_stagedOrder, stage.ValMagneticVariationType_orderStaged, &stage.ValMagneticVariationTypeOrder, stage.ValMagneticVariationTypes_mapString, valmagneticvariationtype, order, valmagneticvariationtype.Name)
}

// Unstage removes valmagneticvariationtype off the model stage
func (valmagneticvariationtype *ValMagneticVariationType) Unstage(stage *Stage) *ValMagneticVariationType {
	__gong__unstage(stage.ValMagneticVariationTypes, stage.ValMagneticVariationTypes_mapString, valmagneticvariationtype, valmagneticvariationtype.Name)
	return valmagneticvariationtype
}

// UnstageVoid removes valmagneticvariationtype off the model stage
func (valmagneticvariationtype *ValMagneticVariationType) UnstageVoid(stage *Stage) {
	valmagneticvariationtype.Unstage(stage)
}

func (valmagneticvariationtype *ValMagneticVariationType) StageVoid(stage *Stage) {
	valmagneticvariationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valmagneticvariationtype *ValMagneticVariationType) GetName() (res string) {
	return valmagneticvariationtype.Name
}

// for satisfaction of GongStruct interface
func (valmagneticvariationtype *ValMagneticVariationType) SetName(name string) {
	valmagneticvariationtype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeAirspacePointPositionTypes, &stage.CodeAirspacePointPositionTypes_mapString, &stage.CodeAirspacePointPositionType_stagedOrder, &stage.CodeAirspacePointPositionTypeOrder)

	__gong__resetStageType(&stage.CodeAirspacePointRoleTypes, &stage.CodeAirspacePointRoleTypes_mapString, &stage.CodeAirspacePointRoleType_stagedOrder, &stage.CodeAirspacePointRoleTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString, &stage.CodeAltitudeUseType_stagedOrder, &stage.CodeAltitudeUseTypeOrder)

	__gong__resetStageType(&stage.CodeAuralMorseTypes, &stage.CodeAuralMorseTypes_mapString, &stage.CodeAuralMorseType_stagedOrder, &stage.CodeAuralMorseTypeOrder)

	__gong__resetStageType(&stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_mapString, &stage.CodeAuthorityRoleType_stagedOrder, &stage.CodeAuthorityRoleTypeOrder)

	__gong__resetStageType(&stage.CodeBearingTypes, &stage.CodeBearingTypes_mapString, &stage.CodeBearingType_stagedOrder, &stage.CodeBearingTypeOrder)

	__gong__resetStageType(&stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_mapString, &stage.CodeCardinalDirectionType_stagedOrder, &stage.CodeCardinalDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeCheckpointCategoryTypes, &stage.CodeCheckpointCategoryTypes_mapString, &stage.CodeCheckpointCategoryType_stagedOrder, &stage.CodeCheckpointCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeCourseQualityILSTypes, &stage.CodeCourseQualityILSTypes_mapString, &stage.CodeCourseQualityILSType_stagedOrder, &stage.CodeCourseQualityILSTypeOrder)

	__gong__resetStageType(&stage.CodeDMEChannelTypes, &stage.CodeDMEChannelTypes_mapString, &stage.CodeDMEChannelType_stagedOrder, &stage.CodeDMEChannelTypeOrder)

	__gong__resetStageType(&stage.CodeDMETypes, &stage.CodeDMETypes_mapString, &stage.CodeDMEType_stagedOrder, &stage.CodeDMETypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_mapString, &stage.CodeDesignatedPointDesignatorType_stagedOrder, &stage.CodeDesignatedPointDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeDesignatedPointTypes, &stage.CodeDesignatedPointTypes_mapString, &stage.CodeDesignatedPointType_stagedOrder, &stage.CodeDesignatedPointTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_mapString, &stage.CodeDirectionReferenceType_stagedOrder, &stage.CodeDirectionReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeDistanceIndicationTypes, &stage.CodeDistanceIndicationTypes_mapString, &stage.CodeDistanceIndicationType_stagedOrder, &stage.CodeDistanceIndicationTypeOrder)

	__gong__resetStageType(&stage.CodeEmissionBandTypes, &stage.CodeEmissionBandTypes_mapString, &stage.CodeEmissionBandType_stagedOrder, &stage.CodeEmissionBandTypeOrder)

	__gong__resetStageType(&stage.CodeILSBackCourseTypes, &stage.CodeILSBackCourseTypes_mapString, &stage.CodeILSBackCourseType_stagedOrder, &stage.CodeILSBackCourseTypeOrder)

	__gong__resetStageType(&stage.CodeIntegrityLevelILSTypes, &stage.CodeIntegrityLevelILSTypes_mapString, &stage.CodeIntegrityLevelILSType_stagedOrder, &stage.CodeIntegrityLevelILSTypeOrder)

	__gong__resetStageType(&stage.CodeMLSAzimuthTypes, &stage.CodeMLSAzimuthTypes_mapString, &stage.CodeMLSAzimuthType_stagedOrder, &stage.CodeMLSAzimuthTypeOrder)

	__gong__resetStageType(&stage.CodeMLSChannelTypes, &stage.CodeMLSChannelTypes_mapString, &stage.CodeMLSChannelType_stagedOrder, &stage.CodeMLSChannelTypeOrder)

	__gong__resetStageType(&stage.CodeMarkerBeaconSignalTypes, &stage.CodeMarkerBeaconSignalTypes_mapString, &stage.CodeMarkerBeaconSignalType_stagedOrder, &stage.CodeMarkerBeaconSignalTypeOrder)

	__gong__resetStageType(&stage.CodeNDBUsageTypes, &stage.CodeNDBUsageTypes_mapString, &stage.CodeNDBUsageType_stagedOrder, &stage.CodeNDBUsageTypeOrder)

	__gong__resetStageType(&stage.CodeNavaidDesignatorTypes, &stage.CodeNavaidDesignatorTypes_mapString, &stage.CodeNavaidDesignatorType_stagedOrder, &stage.CodeNavaidDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeNavaidPurposeTypes, &stage.CodeNavaidPurposeTypes_mapString, &stage.CodeNavaidPurposeType_stagedOrder, &stage.CodeNavaidPurposeTypeOrder)

	__gong__resetStageType(&stage.CodeNavaidServiceTypes, &stage.CodeNavaidServiceTypes_mapString, &stage.CodeNavaidServiceType_stagedOrder, &stage.CodeNavaidServiceTypeOrder)

	__gong__resetStageType(&stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_mapString, &stage.CodeNorthReferenceType_stagedOrder, &stage.CodeNorthReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodePARTypes, &stage.CodePARTypes_mapString, &stage.CodePARType_stagedOrder, &stage.CodePARTypeOrder)

	__gong__resetStageType(&stage.CodePositionInILSTypes, &stage.CodePositionInILSTypes_mapString, &stage.CodePositionInILSType_stagedOrder, &stage.CodePositionInILSTypeOrder)

	__gong__resetStageType(&stage.CodePrimaryRadarTypes, &stage.CodePrimaryRadarTypes_mapString, &stage.CodePrimaryRadarType_stagedOrder, &stage.CodePrimaryRadarTypeOrder)

	__gong__resetStageType(&stage.CodeRadarServiceTypes, &stage.CodeRadarServiceTypes_mapString, &stage.CodeRadarServiceType_stagedOrder, &stage.CodeRadarServiceTypeOrder)

	__gong__resetStageType(&stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_mapString, &stage.CodeRadioEmissionType_stagedOrder, &stage.CodeRadioEmissionTypeOrder)

	__gong__resetStageType(&stage.CodeRadioSignalTypes, &stage.CodeRadioSignalTypes_mapString, &stage.CodeRadioSignalType_stagedOrder, &stage.CodeRadioSignalTypeOrder)

	__gong__resetStageType(&stage.CodeReflectorTypes, &stage.CodeReflectorTypes_mapString, &stage.CodeReflectorType_stagedOrder, &stage.CodeReflectorTypeOrder)

	__gong__resetStageType(&stage.CodeSignalPerformanceILSTypes, &stage.CodeSignalPerformanceILSTypes_mapString, &stage.CodeSignalPerformanceILSType_stagedOrder, &stage.CodeSignalPerformanceILSTypeOrder)

	__gong__resetStageType(&stage.CodeSpecialNavigationChainDesignatorTypes, &stage.CodeSpecialNavigationChainDesignatorTypes_mapString, &stage.CodeSpecialNavigationChainDesignatorType_stagedOrder, &stage.CodeSpecialNavigationChainDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeSpecialNavigationStationTypes, &stage.CodeSpecialNavigationStationTypes_mapString, &stage.CodeSpecialNavigationStationType_stagedOrder, &stage.CodeSpecialNavigationStationTypeOrder)

	__gong__resetStageType(&stage.CodeSpecialNavigationSystemTypes, &stage.CodeSpecialNavigationSystemTypes_mapString, &stage.CodeSpecialNavigationSystemType_stagedOrder, &stage.CodeSpecialNavigationSystemTypeOrder)

	__gong__resetStageType(&stage.CodeStandbyPowerTypes, &stage.CodeStandbyPowerTypes_mapString, &stage.CodeStandbyPowerType_stagedOrder, &stage.CodeStandbyPowerTypeOrder)

	__gong__resetStageType(&stage.CodeStatusNavaidTypes, &stage.CodeStatusNavaidTypes_mapString, &stage.CodeStatusNavaidType_stagedOrder, &stage.CodeStatusNavaidTypeOrder)

	__gong__resetStageType(&stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_mapString, &stage.CodeTACANChannelType_stagedOrder, &stage.CodeTACANChannelTypeOrder)

	__gong__resetStageType(&stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString, &stage.CodeTelecomNetworkType_stagedOrder, &stage.CodeTelecomNetworkTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString, &stage.CodeTransponderType_stagedOrder, &stage.CodeTransponderTypeOrder)

	__gong__resetStageType(&stage.CodeVORTypes, &stage.CodeVORTypes_mapString, &stage.CodeVORType_stagedOrder, &stage.CodeVORTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString, &stage.CodeVerticalReferenceType_stagedOrder, &stage.CodeVerticalReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.DateYearTypes, &stage.DateYearTypes_mapString, &stage.DateYearType_stagedOrder, &stage.DateYearTypeOrder)

	__gong__resetStageType(&stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString, &stage.NoSequenceType_stagedOrder, &stage.NoSequenceTypeOrder)

	__gong__resetStageType(&stage.TextAddressTypes, &stage.TextAddressTypes_mapString, &stage.TextAddressType_stagedOrder, &stage.TextAddressTypeOrder)

	__gong__resetStageType(&stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString, &stage.TextDesignatorType_stagedOrder, &stage.TextDesignatorTypeOrder)

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

	__gong__resetStageType(&stage.ValFrequencyTypes, &stage.ValFrequencyTypes_mapString, &stage.ValFrequencyType_stagedOrder, &stage.ValFrequencyTypeOrder)

	__gong__resetStageType(&stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_mapString, &stage.ValMagneticVariationType_stagedOrder, &stage.ValMagneticVariationTypeOrder)

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
func (CodeAirspacePointPositionType) GongGetAssociationName() any {
	return &CodeAirspacePointPositionType{
	}
}

func (CodeAirspacePointRoleType) GongGetAssociationName() any {
	return &CodeAirspacePointRoleType{
	}
}

func (CodeAltitudeUseType) GongGetAssociationName() any {
	return &CodeAltitudeUseType{
	}
}

func (CodeAuralMorseType) GongGetAssociationName() any {
	return &CodeAuralMorseType{
	}
}

func (CodeAuthorityRoleType) GongGetAssociationName() any {
	return &CodeAuthorityRoleType{
	}
}

func (CodeBearingType) GongGetAssociationName() any {
	return &CodeBearingType{
	}
}

func (CodeCardinalDirectionType) GongGetAssociationName() any {
	return &CodeCardinalDirectionType{
	}
}

func (CodeCheckpointCategoryType) GongGetAssociationName() any {
	return &CodeCheckpointCategoryType{
	}
}

func (CodeCourseQualityILSType) GongGetAssociationName() any {
	return &CodeCourseQualityILSType{
	}
}

func (CodeDMEChannelType) GongGetAssociationName() any {
	return &CodeDMEChannelType{
	}
}

func (CodeDMEType) GongGetAssociationName() any {
	return &CodeDMEType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeDesignatedPointDesignatorType) GongGetAssociationName() any {
	return &CodeDesignatedPointDesignatorType{
	}
}

func (CodeDesignatedPointType) GongGetAssociationName() any {
	return &CodeDesignatedPointType{
	}
}

func (CodeDirectionReferenceType) GongGetAssociationName() any {
	return &CodeDirectionReferenceType{
	}
}

func (CodeDistanceIndicationType) GongGetAssociationName() any {
	return &CodeDistanceIndicationType{
	}
}

func (CodeEmissionBandType) GongGetAssociationName() any {
	return &CodeEmissionBandType{
	}
}

func (CodeILSBackCourseType) GongGetAssociationName() any {
	return &CodeILSBackCourseType{
	}
}

func (CodeIntegrityLevelILSType) GongGetAssociationName() any {
	return &CodeIntegrityLevelILSType{
	}
}

func (CodeMLSAzimuthType) GongGetAssociationName() any {
	return &CodeMLSAzimuthType{
	}
}

func (CodeMLSChannelType) GongGetAssociationName() any {
	return &CodeMLSChannelType{
	}
}

func (CodeMarkerBeaconSignalType) GongGetAssociationName() any {
	return &CodeMarkerBeaconSignalType{
	}
}

func (CodeNDBUsageType) GongGetAssociationName() any {
	return &CodeNDBUsageType{
	}
}

func (CodeNavaidDesignatorType) GongGetAssociationName() any {
	return &CodeNavaidDesignatorType{
	}
}

func (CodeNavaidPurposeType) GongGetAssociationName() any {
	return &CodeNavaidPurposeType{
	}
}

func (CodeNavaidServiceType) GongGetAssociationName() any {
	return &CodeNavaidServiceType{
	}
}

func (CodeNorthReferenceType) GongGetAssociationName() any {
	return &CodeNorthReferenceType{
	}
}

func (CodeNotePurposeType) GongGetAssociationName() any {
	return &CodeNotePurposeType{
	}
}

func (CodePARType) GongGetAssociationName() any {
	return &CodePARType{
	}
}

func (CodePositionInILSType) GongGetAssociationName() any {
	return &CodePositionInILSType{
	}
}

func (CodePrimaryRadarType) GongGetAssociationName() any {
	return &CodePrimaryRadarType{
	}
}

func (CodeRadarServiceType) GongGetAssociationName() any {
	return &CodeRadarServiceType{
	}
}

func (CodeRadioEmissionType) GongGetAssociationName() any {
	return &CodeRadioEmissionType{
	}
}

func (CodeRadioSignalType) GongGetAssociationName() any {
	return &CodeRadioSignalType{
	}
}

func (CodeReflectorType) GongGetAssociationName() any {
	return &CodeReflectorType{
	}
}

func (CodeSignalPerformanceILSType) GongGetAssociationName() any {
	return &CodeSignalPerformanceILSType{
	}
}

func (CodeSpecialNavigationChainDesignatorType) GongGetAssociationName() any {
	return &CodeSpecialNavigationChainDesignatorType{
	}
}

func (CodeSpecialNavigationStationType) GongGetAssociationName() any {
	return &CodeSpecialNavigationStationType{
	}
}

func (CodeSpecialNavigationSystemType) GongGetAssociationName() any {
	return &CodeSpecialNavigationSystemType{
	}
}

func (CodeStandbyPowerType) GongGetAssociationName() any {
	return &CodeStandbyPowerType{
	}
}

func (CodeStatusNavaidType) GongGetAssociationName() any {
	return &CodeStatusNavaidType{
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

func (CodeVORType) GongGetAssociationName() any {
	return &CodeVORType{
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

func (CodeYesNoType) GongGetAssociationName() any {
	return &CodeYesNoType{
	}
}

func (DateMonthDayType) GongGetAssociationName() any {
	return &DateMonthDayType{
	}
}

func (DateYearType) GongGetAssociationName() any {
	return &DateYearType{
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

func (ValFrequencyType) GongGetAssociationName() any {
	return &ValFrequencyType{
	}
}

func (ValMagneticVariationType) GongGetAssociationName() any {
	return &ValMagneticVariationType{
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
	// reverse maps of direct associations of CodeAirspacePointPositionType
	case CodeAirspacePointPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspacePointRoleType
	case CodeAirspacePointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeUseType
	case CodeAltitudeUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuralMorseType
	case CodeAuralMorseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityRoleType
	case CodeAuthorityRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBearingType
	case CodeBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCardinalDirectionType
	case CodeCardinalDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCheckpointCategoryType
	case CodeCheckpointCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCourseQualityILSType
	case CodeCourseQualityILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDMEChannelType
	case CodeDMEChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDMEType
	case CodeDMEType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointDesignatorType
	case CodeDesignatedPointDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointType
	case CodeDesignatedPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionReferenceType
	case CodeDirectionReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDistanceIndicationType
	case CodeDistanceIndicationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEmissionBandType
	case CodeEmissionBandType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeILSBackCourseType
	case CodeILSBackCourseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIntegrityLevelILSType
	case CodeIntegrityLevelILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMLSAzimuthType
	case CodeMLSAzimuthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMLSChannelType
	case CodeMLSChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkerBeaconSignalType
	case CodeMarkerBeaconSignalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNDBUsageType
	case CodeNDBUsageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidDesignatorType
	case CodeNavaidDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidPurposeType
	case CodeNavaidPurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidServiceType
	case CodeNavaidServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNorthReferenceType
	case CodeNorthReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePARType
	case CodePARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePositionInILSType
	case CodePositionInILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePrimaryRadarType
	case CodePrimaryRadarType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadarServiceType
	case CodeRadarServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioEmissionType
	case CodeRadioEmissionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioSignalType
	case CodeRadioSignalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReflectorType
	case CodeReflectorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSignalPerformanceILSType
	case CodeSignalPerformanceILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationChainDesignatorType
	case CodeSpecialNavigationChainDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationStationType
	case CodeSpecialNavigationStationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationSystemType
	case CodeSpecialNavigationSystemType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStandbyPowerType
	case CodeStandbyPowerType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusNavaidType
	case CodeStatusNavaidType:
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
	// reverse maps of direct associations of CodeVORType
	case CodeVORType:
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
	// reverse maps of direct associations of DateYearType
	case DateYearType:
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
	// reverse maps of direct associations of ValFrequencyType
	case ValFrequencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationType
	case ValMagneticVariationType:
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
	// reverse maps of direct associations of CodeAirspacePointPositionType
	case CodeAirspacePointPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspacePointRoleType
	case CodeAirspacePointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAltitudeUseType
	case CodeAltitudeUseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuralMorseType
	case CodeAuralMorseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityRoleType
	case CodeAuthorityRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBearingType
	case CodeBearingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCardinalDirectionType
	case CodeCardinalDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCheckpointCategoryType
	case CodeCheckpointCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCourseQualityILSType
	case CodeCourseQualityILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDMEChannelType
	case CodeDMEChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDMEType
	case CodeDMEType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointDesignatorType
	case CodeDesignatedPointDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDesignatedPointType
	case CodeDesignatedPointType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDirectionReferenceType
	case CodeDirectionReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDistanceIndicationType
	case CodeDistanceIndicationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeEmissionBandType
	case CodeEmissionBandType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeILSBackCourseType
	case CodeILSBackCourseType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIntegrityLevelILSType
	case CodeIntegrityLevelILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMLSAzimuthType
	case CodeMLSAzimuthType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMLSChannelType
	case CodeMLSChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkerBeaconSignalType
	case CodeMarkerBeaconSignalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNDBUsageType
	case CodeNDBUsageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidDesignatorType
	case CodeNavaidDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidPurposeType
	case CodeNavaidPurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNavaidServiceType
	case CodeNavaidServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNorthReferenceType
	case CodeNorthReferenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePARType
	case CodePARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePositionInILSType
	case CodePositionInILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePrimaryRadarType
	case CodePrimaryRadarType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadarServiceType
	case CodeRadarServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioEmissionType
	case CodeRadioEmissionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioSignalType
	case CodeRadioSignalType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeReflectorType
	case CodeReflectorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSignalPerformanceILSType
	case CodeSignalPerformanceILSType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationChainDesignatorType
	case CodeSpecialNavigationChainDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationStationType
	case CodeSpecialNavigationStationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialNavigationSystemType
	case CodeSpecialNavigationSystemType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStandbyPowerType
	case CodeStandbyPowerType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusNavaidType
	case CodeStatusNavaidType:
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
	// reverse maps of direct associations of CodeVORType
	case CodeVORType:
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
	// reverse maps of direct associations of DateYearType
	case DateYearType:
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
	// reverse maps of direct associations of ValFrequencyType
	case ValFrequencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationType
	case ValMagneticVariationType:
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
func (*CodeAirspacePointPositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspacePointRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAltitudeUseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAuralMorseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAuthorityRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeBearingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCardinalDirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCheckpointCategoryType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCourseQualityILSType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDMEChannelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDMEType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDesignatedPointDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDesignatedPointType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDirectionReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDistanceIndicationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeEmissionBandType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeILSBackCourseType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeIntegrityLevelILSType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMLSAzimuthType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMLSChannelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMarkerBeaconSignalType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNDBUsageType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavaidDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavaidPurposeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNavaidServiceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNorthReferenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNotePurposeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePARType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePositionInILSType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePrimaryRadarType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRadarServiceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRadioEmissionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRadioSignalType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeReflectorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSignalPerformanceILSType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSpecialNavigationChainDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSpecialNavigationStationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSpecialNavigationSystemType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStandbyPowerType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusNavaidType) GongGetReverseFields() []GongReverseField {
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

func (*CodeVORType) GongGetReverseFields() []GongReverseField {
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

func (*CodeYesNoType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DateMonthDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DateYearType) GongGetReverseFields() []GongReverseField {
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

func (*ValFrequencyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValMagneticVariationType) GongGetReverseFields() []GongReverseField {
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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeauralmorsetype *CodeAuralMorseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedmechanneltype *CodeDMEChannelType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedmetype *CodeDMEType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeemissionbandtype *CodeEmissionBandType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codemlschanneltype *CodeMLSChannelType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codendbusagetype *CodeNDBUsageType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenavaidservicetype *CodeNavaidServiceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codenorthreferencetype *CodeNorthReferenceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codepartype *CodePARType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codepositioninilstype *CodePositionInILSType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeprimaryradartype *CodePrimaryRadarType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (coderadarservicetype *CodeRadarServiceType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (coderadioemissiontype *CodeRadioEmissionType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (coderadiosignaltype *CodeRadioSignalType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codereflectortype *CodeReflectorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codestandbypowertype *CodeStandbyPowerType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codevortype *CodeVORType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (dateyeartype *DateYearType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valfrequencytype *ValFrequencyType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valmagneticvariationtype *ValMagneticVariationType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspacepointpositiontype.Name
	case "NilReason":
		res.valueString = codeairspacepointpositiontype.NilReason
	case "EnclosedText":
		res.valueString = codeairspacepointpositiontype.EnclosedText
	}
	return
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspacepointroletype.Name
	case "NilReason":
		res.valueString = codeairspacepointroletype.NilReason
	case "EnclosedText":
		res.valueString = codeairspacepointroletype.EnclosedText
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

func (codeauralmorsetype *CodeAuralMorseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeauralmorsetype.Name
	case "NilReason":
		res.valueString = codeauralmorsetype.NilReason
	case "EnclosedText":
		res.valueString = codeauralmorsetype.EnclosedText
	}
	return
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeauthorityroletype.Name
	case "NilReason":
		res.valueString = codeauthorityroletype.NilReason
	case "EnclosedText":
		res.valueString = codeauthorityroletype.EnclosedText
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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecheckpointcategorytype.Name
	case "NilReason":
		res.valueString = codecheckpointcategorytype.NilReason
	case "EnclosedText":
		res.valueString = codecheckpointcategorytype.EnclosedText
	}
	return
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecoursequalityilstype.Name
	case "NilReason":
		res.valueString = codecoursequalityilstype.NilReason
	case "EnclosedText":
		res.valueString = codecoursequalityilstype.EnclosedText
	}
	return
}

func (codedmechanneltype *CodeDMEChannelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedmechanneltype.Name
	case "NilReason":
		res.valueString = codedmechanneltype.NilReason
	case "EnclosedText":
		res.valueString = codedmechanneltype.EnclosedText
	}
	return
}

func (codedmetype *CodeDMEType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedmetype.Name
	case "NilReason":
		res.valueString = codedmetype.NilReason
	case "EnclosedText":
		res.valueString = codedmetype.EnclosedText
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

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedesignatedpointtype.Name
	case "NilReason":
		res.valueString = codedesignatedpointtype.NilReason
	case "EnclosedText":
		res.valueString = codedesignatedpointtype.EnclosedText
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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedistanceindicationtype.Name
	case "NilReason":
		res.valueString = codedistanceindicationtype.NilReason
	case "EnclosedText":
		res.valueString = codedistanceindicationtype.EnclosedText
	}
	return
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeemissionbandtype.Name
	case "NilReason":
		res.valueString = codeemissionbandtype.NilReason
	case "EnclosedText":
		res.valueString = codeemissionbandtype.EnclosedText
	}
	return
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeilsbackcoursetype.Name
	case "NilReason":
		res.valueString = codeilsbackcoursetype.NilReason
	case "EnclosedText":
		res.valueString = codeilsbackcoursetype.EnclosedText
	}
	return
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeintegritylevelilstype.Name
	case "NilReason":
		res.valueString = codeintegritylevelilstype.NilReason
	case "EnclosedText":
		res.valueString = codeintegritylevelilstype.EnclosedText
	}
	return
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemlsazimuthtype.Name
	case "NilReason":
		res.valueString = codemlsazimuthtype.NilReason
	case "EnclosedText":
		res.valueString = codemlsazimuthtype.EnclosedText
	}
	return
}

func (codemlschanneltype *CodeMLSChannelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemlschanneltype.Name
	case "NilReason":
		res.valueString = codemlschanneltype.NilReason
	case "EnclosedText":
		res.valueString = codemlschanneltype.EnclosedText
	}
	return
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemarkerbeaconsignaltype.Name
	case "NilReason":
		res.valueString = codemarkerbeaconsignaltype.NilReason
	case "EnclosedText":
		res.valueString = codemarkerbeaconsignaltype.EnclosedText
	}
	return
}

func (codendbusagetype *CodeNDBUsageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codendbusagetype.Name
	case "NilReason":
		res.valueString = codendbusagetype.NilReason
	case "EnclosedText":
		res.valueString = codendbusagetype.EnclosedText
	}
	return
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavaiddesignatortype.Name
	case "NilReason":
		res.valueString = codenavaiddesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codenavaiddesignatortype.EnclosedText
	}
	return
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavaidpurposetype.Name
	case "NilReason":
		res.valueString = codenavaidpurposetype.NilReason
	case "EnclosedText":
		res.valueString = codenavaidpurposetype.EnclosedText
	}
	return
}

func (codenavaidservicetype *CodeNavaidServiceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenavaidservicetype.Name
	case "NilReason":
		res.valueString = codenavaidservicetype.NilReason
	case "EnclosedText":
		res.valueString = codenavaidservicetype.EnclosedText
	}
	return
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenorthreferencetype.Name
	case "NilReason":
		res.valueString = codenorthreferencetype.NilReason
	case "EnclosedText":
		res.valueString = codenorthreferencetype.EnclosedText
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

func (codepartype *CodePARType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepartype.Name
	case "NilReason":
		res.valueString = codepartype.NilReason
	case "EnclosedText":
		res.valueString = codepartype.EnclosedText
	}
	return
}

func (codepositioninilstype *CodePositionInILSType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepositioninilstype.Name
	case "NilReason":
		res.valueString = codepositioninilstype.NilReason
	case "EnclosedText":
		res.valueString = codepositioninilstype.EnclosedText
	}
	return
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprimaryradartype.Name
	case "NilReason":
		res.valueString = codeprimaryradartype.NilReason
	case "EnclosedText":
		res.valueString = codeprimaryradartype.EnclosedText
	}
	return
}

func (coderadarservicetype *CodeRadarServiceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderadarservicetype.Name
	case "NilReason":
		res.valueString = coderadarservicetype.NilReason
	case "EnclosedText":
		res.valueString = coderadarservicetype.EnclosedText
	}
	return
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderadioemissiontype.Name
	case "NilReason":
		res.valueString = coderadioemissiontype.NilReason
	case "EnclosedText":
		res.valueString = coderadioemissiontype.EnclosedText
	}
	return
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderadiosignaltype.Name
	case "NilReason":
		res.valueString = coderadiosignaltype.NilReason
	case "EnclosedText":
		res.valueString = coderadiosignaltype.EnclosedText
	}
	return
}

func (codereflectortype *CodeReflectorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codereflectortype.Name
	case "NilReason":
		res.valueString = codereflectortype.NilReason
	case "EnclosedText":
		res.valueString = codereflectortype.EnclosedText
	}
	return
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesignalperformanceilstype.Name
	case "NilReason":
		res.valueString = codesignalperformanceilstype.NilReason
	case "EnclosedText":
		res.valueString = codesignalperformanceilstype.EnclosedText
	}
	return
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codespecialnavigationchaindesignatortype.Name
	case "NilReason":
		res.valueString = codespecialnavigationchaindesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codespecialnavigationchaindesignatortype.EnclosedText
	}
	return
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codespecialnavigationstationtype.Name
	case "NilReason":
		res.valueString = codespecialnavigationstationtype.NilReason
	case "EnclosedText":
		res.valueString = codespecialnavigationstationtype.EnclosedText
	}
	return
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codespecialnavigationsystemtype.Name
	case "NilReason":
		res.valueString = codespecialnavigationsystemtype.NilReason
	case "EnclosedText":
		res.valueString = codespecialnavigationsystemtype.EnclosedText
	}
	return
}

func (codestandbypowertype *CodeStandbyPowerType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestandbypowertype.Name
	case "NilReason":
		res.valueString = codestandbypowertype.NilReason
	case "EnclosedText":
		res.valueString = codestandbypowertype.EnclosedText
	}
	return
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusnavaidtype.Name
	case "NilReason":
		res.valueString = codestatusnavaidtype.NilReason
	case "EnclosedText":
		res.valueString = codestatusnavaidtype.EnclosedText
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

func (codevortype *CodeVORType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codevortype.Name
	case "NilReason":
		res.valueString = codevortype.NilReason
	case "EnclosedText":
		res.valueString = codevortype.EnclosedText
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

func (dateyeartype *DateYearType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dateyeartype.Name
	case "NilReason":
		res.valueString = dateyeartype.NilReason
	case "EnclosedText":
		res.valueString = dateyeartype.EnclosedText
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

func (valfrequencytype *ValFrequencyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valfrequencytype.Name
	case "Uom":
		res.valueString = valfrequencytype.Uom
	case "NilReason":
		res.valueString = valfrequencytype.NilReason
	case "EnclosedText":
		res.valueString = valfrequencytype.EnclosedText
	}
	return
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valmagneticvariationtype.Name
	case "NilReason":
		res.valueString = valmagneticvariationtype.NilReason
	case "EnclosedText":
		res.valueString = valmagneticvariationtype.EnclosedText
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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetGongstructName() string {
	return "CodeAirspacePointPositionType"
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetGongstructName() string {
	return "CodeAirspacePointRoleType"
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetGongstructName() string {
	return "CodeAltitudeUseType"
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetGongstructName() string {
	return "CodeAuralMorseType"
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetGongstructName() string {
	return "CodeAuthorityRoleType"
}

func (codebearingtype *CodeBearingType) GongGetGongstructName() string {
	return "CodeBearingType"
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetGongstructName() string {
	return "CodeCardinalDirectionType"
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetGongstructName() string {
	return "CodeCheckpointCategoryType"
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetGongstructName() string {
	return "CodeCourseQualityILSType"
}

func (codedmechanneltype *CodeDMEChannelType) GongGetGongstructName() string {
	return "CodeDMEChannelType"
}

func (codedmetype *CodeDMEType) GongGetGongstructName() string {
	return "CodeDMEType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetGongstructName() string {
	return "CodeDesignatedPointDesignatorType"
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetGongstructName() string {
	return "CodeDesignatedPointType"
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetGongstructName() string {
	return "CodeDirectionReferenceType"
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetGongstructName() string {
	return "CodeDistanceIndicationType"
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetGongstructName() string {
	return "CodeEmissionBandType"
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetGongstructName() string {
	return "CodeILSBackCourseType"
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetGongstructName() string {
	return "CodeIntegrityLevelILSType"
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetGongstructName() string {
	return "CodeMLSAzimuthType"
}

func (codemlschanneltype *CodeMLSChannelType) GongGetGongstructName() string {
	return "CodeMLSChannelType"
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetGongstructName() string {
	return "CodeMarkerBeaconSignalType"
}

func (codendbusagetype *CodeNDBUsageType) GongGetGongstructName() string {
	return "CodeNDBUsageType"
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetGongstructName() string {
	return "CodeNavaidDesignatorType"
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetGongstructName() string {
	return "CodeNavaidPurposeType"
}

func (codenavaidservicetype *CodeNavaidServiceType) GongGetGongstructName() string {
	return "CodeNavaidServiceType"
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetGongstructName() string {
	return "CodeNorthReferenceType"
}

func (codenotepurposetype *CodeNotePurposeType) GongGetGongstructName() string {
	return "CodeNotePurposeType"
}

func (codepartype *CodePARType) GongGetGongstructName() string {
	return "CodePARType"
}

func (codepositioninilstype *CodePositionInILSType) GongGetGongstructName() string {
	return "CodePositionInILSType"
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetGongstructName() string {
	return "CodePrimaryRadarType"
}

func (coderadarservicetype *CodeRadarServiceType) GongGetGongstructName() string {
	return "CodeRadarServiceType"
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetGongstructName() string {
	return "CodeRadioEmissionType"
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetGongstructName() string {
	return "CodeRadioSignalType"
}

func (codereflectortype *CodeReflectorType) GongGetGongstructName() string {
	return "CodeReflectorType"
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetGongstructName() string {
	return "CodeSignalPerformanceILSType"
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetGongstructName() string {
	return "CodeSpecialNavigationChainDesignatorType"
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetGongstructName() string {
	return "CodeSpecialNavigationStationType"
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetGongstructName() string {
	return "CodeSpecialNavigationSystemType"
}

func (codestandbypowertype *CodeStandbyPowerType) GongGetGongstructName() string {
	return "CodeStandbyPowerType"
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetGongstructName() string {
	return "CodeStatusNavaidType"
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

func (codevortype *CodeVORType) GongGetGongstructName() string {
	return "CodeVORType"
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetGongstructName() string {
	return "CodeVerticalDatumType"
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetGongstructName() string {
	return "CodeVerticalReferenceType"
}

func (codeyesnotype *CodeYesNoType) GongGetGongstructName() string {
	return "CodeYesNoType"
}

func (datemonthdaytype *DateMonthDayType) GongGetGongstructName() string {
	return "DateMonthDayType"
}

func (dateyeartype *DateYearType) GongGetGongstructName() string {
	return "DateYearType"
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

func (valfrequencytype *ValFrequencyType) GongGetGongstructName() string {
	return "ValFrequencyType"
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetGongstructName() string {
	return "ValMagneticVariationType"
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
	__gong__rebuildMapString(stage.CodeAirspacePointPositionTypes, &stage.CodeAirspacePointPositionTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspacePointRoleTypes, &stage.CodeAirspacePointRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString)

	__gong__rebuildMapString(stage.CodeAuralMorseTypes, &stage.CodeAuralMorseTypes_mapString)

	__gong__rebuildMapString(stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeBearingTypes, &stage.CodeBearingTypes_mapString)

	__gong__rebuildMapString(stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeCheckpointCategoryTypes, &stage.CodeCheckpointCategoryTypes_mapString)

	__gong__rebuildMapString(stage.CodeCourseQualityILSTypes, &stage.CodeCourseQualityILSTypes_mapString)

	__gong__rebuildMapString(stage.CodeDMEChannelTypes, &stage.CodeDMEChannelTypes_mapString)

	__gong__rebuildMapString(stage.CodeDMETypes, &stage.CodeDMETypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeDesignatedPointTypes, &stage.CodeDesignatedPointTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeDistanceIndicationTypes, &stage.CodeDistanceIndicationTypes_mapString)

	__gong__rebuildMapString(stage.CodeEmissionBandTypes, &stage.CodeEmissionBandTypes_mapString)

	__gong__rebuildMapString(stage.CodeILSBackCourseTypes, &stage.CodeILSBackCourseTypes_mapString)

	__gong__rebuildMapString(stage.CodeIntegrityLevelILSTypes, &stage.CodeIntegrityLevelILSTypes_mapString)

	__gong__rebuildMapString(stage.CodeMLSAzimuthTypes, &stage.CodeMLSAzimuthTypes_mapString)

	__gong__rebuildMapString(stage.CodeMLSChannelTypes, &stage.CodeMLSChannelTypes_mapString)

	__gong__rebuildMapString(stage.CodeMarkerBeaconSignalTypes, &stage.CodeMarkerBeaconSignalTypes_mapString)

	__gong__rebuildMapString(stage.CodeNDBUsageTypes, &stage.CodeNDBUsageTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavaidDesignatorTypes, &stage.CodeNavaidDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavaidPurposeTypes, &stage.CodeNavaidPurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavaidServiceTypes, &stage.CodeNavaidServiceTypes_mapString)

	__gong__rebuildMapString(stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodePARTypes, &stage.CodePARTypes_mapString)

	__gong__rebuildMapString(stage.CodePositionInILSTypes, &stage.CodePositionInILSTypes_mapString)

	__gong__rebuildMapString(stage.CodePrimaryRadarTypes, &stage.CodePrimaryRadarTypes_mapString)

	__gong__rebuildMapString(stage.CodeRadarServiceTypes, &stage.CodeRadarServiceTypes_mapString)

	__gong__rebuildMapString(stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_mapString)

	__gong__rebuildMapString(stage.CodeRadioSignalTypes, &stage.CodeRadioSignalTypes_mapString)

	__gong__rebuildMapString(stage.CodeReflectorTypes, &stage.CodeReflectorTypes_mapString)

	__gong__rebuildMapString(stage.CodeSignalPerformanceILSTypes, &stage.CodeSignalPerformanceILSTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpecialNavigationChainDesignatorTypes, &stage.CodeSpecialNavigationChainDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpecialNavigationStationTypes, &stage.CodeSpecialNavigationStationTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpecialNavigationSystemTypes, &stage.CodeSpecialNavigationSystemTypes_mapString)

	__gong__rebuildMapString(stage.CodeStandbyPowerTypes, &stage.CodeStandbyPowerTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusNavaidTypes, &stage.CodeStatusNavaidTypes_mapString)

	__gong__rebuildMapString(stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_mapString)

	__gong__rebuildMapString(stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString)

	__gong__rebuildMapString(stage.CodeVORTypes, &stage.CodeVORTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.DateYearTypes, &stage.DateYearTypes_mapString)

	__gong__rebuildMapString(stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString)

	__gong__rebuildMapString(stage.TextAddressTypes, &stage.TextAddressTypes_mapString)

	__gong__rebuildMapString(stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString)

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

	__gong__rebuildMapString(stage.ValFrequencyTypes, &stage.ValFrequencyTypes_mapString)

	__gong__rebuildMapString(stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
