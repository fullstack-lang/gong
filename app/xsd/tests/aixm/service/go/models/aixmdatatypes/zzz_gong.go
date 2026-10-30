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
	CodeAircraftGroundServiceTypes                map[*CodeAircraftGroundServiceType]struct{}
	CodeAircraftGroundServiceTypes_instance       map[*CodeAircraftGroundServiceType]*CodeAircraftGroundServiceType
	CodeAircraftGroundServiceTypes_mapString      map[string]*CodeAircraftGroundServiceType
	CodeAircraftGroundServiceTypeOrder            uint
	CodeAircraftGroundServiceType_stagedOrder     map[*CodeAircraftGroundServiceType]uint
	CodeAircraftGroundServiceType_orderStaged     map[uint]*CodeAircraftGroundServiceType
	CodeAircraftGroundServiceTypes_reference      map[*CodeAircraftGroundServiceType]*CodeAircraftGroundServiceType
	CodeAircraftGroundServiceTypes_referenceOrder map[*CodeAircraftGroundServiceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftGroundServiceTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftGroundServiceType]
	OnAfterCodeAircraftGroundServiceTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftGroundServiceType]
	OnAfterCodeAircraftGroundServiceTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftGroundServiceType]

	CodeAviationStandardsTypes                map[*CodeAviationStandardsType]struct{}
	CodeAviationStandardsTypes_instance       map[*CodeAviationStandardsType]*CodeAviationStandardsType
	CodeAviationStandardsTypes_mapString      map[string]*CodeAviationStandardsType
	CodeAviationStandardsTypeOrder            uint
	CodeAviationStandardsType_stagedOrder     map[*CodeAviationStandardsType]uint
	CodeAviationStandardsType_orderStaged     map[uint]*CodeAviationStandardsType
	CodeAviationStandardsTypes_reference      map[*CodeAviationStandardsType]*CodeAviationStandardsType
	CodeAviationStandardsTypes_referenceOrder map[*CodeAviationStandardsType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAviationStandardsTypeCreateCallback GongOnAfterCreateInterface[CodeAviationStandardsType]
	OnAfterCodeAviationStandardsTypeUpdateCallback GongOnAfterUpdateInterface[CodeAviationStandardsType]
	OnAfterCodeAviationStandardsTypeDeleteCallback GongOnAfterDeleteInterface[CodeAviationStandardsType]

	CodeCommunicationChannelTypes                map[*CodeCommunicationChannelType]struct{}
	CodeCommunicationChannelTypes_instance       map[*CodeCommunicationChannelType]*CodeCommunicationChannelType
	CodeCommunicationChannelTypes_mapString      map[string]*CodeCommunicationChannelType
	CodeCommunicationChannelTypeOrder            uint
	CodeCommunicationChannelType_stagedOrder     map[*CodeCommunicationChannelType]uint
	CodeCommunicationChannelType_orderStaged     map[uint]*CodeCommunicationChannelType
	CodeCommunicationChannelTypes_reference      map[*CodeCommunicationChannelType]*CodeCommunicationChannelType
	CodeCommunicationChannelTypes_referenceOrder map[*CodeCommunicationChannelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCommunicationChannelTypeCreateCallback GongOnAfterCreateInterface[CodeCommunicationChannelType]
	OnAfterCodeCommunicationChannelTypeUpdateCallback GongOnAfterUpdateInterface[CodeCommunicationChannelType]
	OnAfterCodeCommunicationChannelTypeDeleteCallback GongOnAfterDeleteInterface[CodeCommunicationChannelType]

	CodeCommunicationDirectionTypes                map[*CodeCommunicationDirectionType]struct{}
	CodeCommunicationDirectionTypes_instance       map[*CodeCommunicationDirectionType]*CodeCommunicationDirectionType
	CodeCommunicationDirectionTypes_mapString      map[string]*CodeCommunicationDirectionType
	CodeCommunicationDirectionTypeOrder            uint
	CodeCommunicationDirectionType_stagedOrder     map[*CodeCommunicationDirectionType]uint
	CodeCommunicationDirectionType_orderStaged     map[uint]*CodeCommunicationDirectionType
	CodeCommunicationDirectionTypes_reference      map[*CodeCommunicationDirectionType]*CodeCommunicationDirectionType
	CodeCommunicationDirectionTypes_referenceOrder map[*CodeCommunicationDirectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeCommunicationDirectionTypeCreateCallback GongOnAfterCreateInterface[CodeCommunicationDirectionType]
	OnAfterCodeCommunicationDirectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeCommunicationDirectionType]
	OnAfterCodeCommunicationDirectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeCommunicationDirectionType]

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

	CodeFacilityRankingTypes                map[*CodeFacilityRankingType]struct{}
	CodeFacilityRankingTypes_instance       map[*CodeFacilityRankingType]*CodeFacilityRankingType
	CodeFacilityRankingTypes_mapString      map[string]*CodeFacilityRankingType
	CodeFacilityRankingTypeOrder            uint
	CodeFacilityRankingType_stagedOrder     map[*CodeFacilityRankingType]uint
	CodeFacilityRankingType_orderStaged     map[uint]*CodeFacilityRankingType
	CodeFacilityRankingTypes_reference      map[*CodeFacilityRankingType]*CodeFacilityRankingType
	CodeFacilityRankingTypes_referenceOrder map[*CodeFacilityRankingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFacilityRankingTypeCreateCallback GongOnAfterCreateInterface[CodeFacilityRankingType]
	OnAfterCodeFacilityRankingTypeUpdateCallback GongOnAfterUpdateInterface[CodeFacilityRankingType]
	OnAfterCodeFacilityRankingTypeDeleteCallback GongOnAfterDeleteInterface[CodeFacilityRankingType]

	CodeFireFightingTypes                map[*CodeFireFightingType]struct{}
	CodeFireFightingTypes_instance       map[*CodeFireFightingType]*CodeFireFightingType
	CodeFireFightingTypes_mapString      map[string]*CodeFireFightingType
	CodeFireFightingTypeOrder            uint
	CodeFireFightingType_stagedOrder     map[*CodeFireFightingType]uint
	CodeFireFightingType_orderStaged     map[uint]*CodeFireFightingType
	CodeFireFightingTypes_reference      map[*CodeFireFightingType]*CodeFireFightingType
	CodeFireFightingTypes_referenceOrder map[*CodeFireFightingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFireFightingTypeCreateCallback GongOnAfterCreateInterface[CodeFireFightingType]
	OnAfterCodeFireFightingTypeUpdateCallback GongOnAfterUpdateInterface[CodeFireFightingType]
	OnAfterCodeFireFightingTypeDeleteCallback GongOnAfterDeleteInterface[CodeFireFightingType]

	CodeFlightDestinationTypes                map[*CodeFlightDestinationType]struct{}
	CodeFlightDestinationTypes_instance       map[*CodeFlightDestinationType]*CodeFlightDestinationType
	CodeFlightDestinationTypes_mapString      map[string]*CodeFlightDestinationType
	CodeFlightDestinationTypeOrder            uint
	CodeFlightDestinationType_stagedOrder     map[*CodeFlightDestinationType]uint
	CodeFlightDestinationType_orderStaged     map[uint]*CodeFlightDestinationType
	CodeFlightDestinationTypes_reference      map[*CodeFlightDestinationType]*CodeFlightDestinationType
	CodeFlightDestinationTypes_referenceOrder map[*CodeFlightDestinationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFlightDestinationTypeCreateCallback GongOnAfterCreateInterface[CodeFlightDestinationType]
	OnAfterCodeFlightDestinationTypeUpdateCallback GongOnAfterUpdateInterface[CodeFlightDestinationType]
	OnAfterCodeFlightDestinationTypeDeleteCallback GongOnAfterDeleteInterface[CodeFlightDestinationType]

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

	CodeFuelTypes                map[*CodeFuelType]struct{}
	CodeFuelTypes_instance       map[*CodeFuelType]*CodeFuelType
	CodeFuelTypes_mapString      map[string]*CodeFuelType
	CodeFuelTypeOrder            uint
	CodeFuelType_stagedOrder     map[*CodeFuelType]uint
	CodeFuelType_orderStaged     map[uint]*CodeFuelType
	CodeFuelTypes_reference      map[*CodeFuelType]*CodeFuelType
	CodeFuelTypes_referenceOrder map[*CodeFuelType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFuelTypeCreateCallback GongOnAfterCreateInterface[CodeFuelType]
	OnAfterCodeFuelTypeUpdateCallback GongOnAfterUpdateInterface[CodeFuelType]
	OnAfterCodeFuelTypeDeleteCallback GongOnAfterDeleteInterface[CodeFuelType]

	CodeIntensityStandByTypes                map[*CodeIntensityStandByType]struct{}
	CodeIntensityStandByTypes_instance       map[*CodeIntensityStandByType]*CodeIntensityStandByType
	CodeIntensityStandByTypes_mapString      map[string]*CodeIntensityStandByType
	CodeIntensityStandByTypeOrder            uint
	CodeIntensityStandByType_stagedOrder     map[*CodeIntensityStandByType]uint
	CodeIntensityStandByType_orderStaged     map[uint]*CodeIntensityStandByType
	CodeIntensityStandByTypes_reference      map[*CodeIntensityStandByType]*CodeIntensityStandByType
	CodeIntensityStandByTypes_referenceOrder map[*CodeIntensityStandByType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeIntensityStandByTypeCreateCallback GongOnAfterCreateInterface[CodeIntensityStandByType]
	OnAfterCodeIntensityStandByTypeUpdateCallback GongOnAfterUpdateInterface[CodeIntensityStandByType]
	OnAfterCodeIntensityStandByTypeDeleteCallback GongOnAfterDeleteInterface[CodeIntensityStandByType]

	CodeLanguageTypes                map[*CodeLanguageType]struct{}
	CodeLanguageTypes_instance       map[*CodeLanguageType]*CodeLanguageType
	CodeLanguageTypes_mapString      map[string]*CodeLanguageType
	CodeLanguageTypeOrder            uint
	CodeLanguageType_stagedOrder     map[*CodeLanguageType]uint
	CodeLanguageType_orderStaged     map[uint]*CodeLanguageType
	CodeLanguageTypes_reference      map[*CodeLanguageType]*CodeLanguageType
	CodeLanguageTypes_referenceOrder map[*CodeLanguageType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLanguageTypeCreateCallback GongOnAfterCreateInterface[CodeLanguageType]
	OnAfterCodeLanguageTypeUpdateCallback GongOnAfterUpdateInterface[CodeLanguageType]
	OnAfterCodeLanguageTypeDeleteCallback GongOnAfterDeleteInterface[CodeLanguageType]

	CodeLevelSeriesTypes                map[*CodeLevelSeriesType]struct{}
	CodeLevelSeriesTypes_instance       map[*CodeLevelSeriesType]*CodeLevelSeriesType
	CodeLevelSeriesTypes_mapString      map[string]*CodeLevelSeriesType
	CodeLevelSeriesTypeOrder            uint
	CodeLevelSeriesType_stagedOrder     map[*CodeLevelSeriesType]uint
	CodeLevelSeriesType_orderStaged     map[uint]*CodeLevelSeriesType
	CodeLevelSeriesTypes_reference      map[*CodeLevelSeriesType]*CodeLevelSeriesType
	CodeLevelSeriesTypes_referenceOrder map[*CodeLevelSeriesType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLevelSeriesTypeCreateCallback GongOnAfterCreateInterface[CodeLevelSeriesType]
	OnAfterCodeLevelSeriesTypeUpdateCallback GongOnAfterUpdateInterface[CodeLevelSeriesType]
	OnAfterCodeLevelSeriesTypeDeleteCallback GongOnAfterDeleteInterface[CodeLevelSeriesType]

	CodeLevelTableDesignatorTypes                map[*CodeLevelTableDesignatorType]struct{}
	CodeLevelTableDesignatorTypes_instance       map[*CodeLevelTableDesignatorType]*CodeLevelTableDesignatorType
	CodeLevelTableDesignatorTypes_mapString      map[string]*CodeLevelTableDesignatorType
	CodeLevelTableDesignatorTypeOrder            uint
	CodeLevelTableDesignatorType_stagedOrder     map[*CodeLevelTableDesignatorType]uint
	CodeLevelTableDesignatorType_orderStaged     map[uint]*CodeLevelTableDesignatorType
	CodeLevelTableDesignatorTypes_reference      map[*CodeLevelTableDesignatorType]*CodeLevelTableDesignatorType
	CodeLevelTableDesignatorTypes_referenceOrder map[*CodeLevelTableDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLevelTableDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeLevelTableDesignatorType]
	OnAfterCodeLevelTableDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeLevelTableDesignatorType]
	OnAfterCodeLevelTableDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeLevelTableDesignatorType]

	CodeLightIntensityTypes                map[*CodeLightIntensityType]struct{}
	CodeLightIntensityTypes_instance       map[*CodeLightIntensityType]*CodeLightIntensityType
	CodeLightIntensityTypes_mapString      map[string]*CodeLightIntensityType
	CodeLightIntensityTypeOrder            uint
	CodeLightIntensityType_stagedOrder     map[*CodeLightIntensityType]uint
	CodeLightIntensityType_orderStaged     map[uint]*CodeLightIntensityType
	CodeLightIntensityTypes_reference      map[*CodeLightIntensityType]*CodeLightIntensityType
	CodeLightIntensityTypes_referenceOrder map[*CodeLightIntensityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLightIntensityTypeCreateCallback GongOnAfterCreateInterface[CodeLightIntensityType]
	OnAfterCodeLightIntensityTypeUpdateCallback GongOnAfterUpdateInterface[CodeLightIntensityType]
	OnAfterCodeLightIntensityTypeDeleteCallback GongOnAfterDeleteInterface[CodeLightIntensityType]

	CodeMilitaryOperationsTypes                map[*CodeMilitaryOperationsType]struct{}
	CodeMilitaryOperationsTypes_instance       map[*CodeMilitaryOperationsType]*CodeMilitaryOperationsType
	CodeMilitaryOperationsTypes_mapString      map[string]*CodeMilitaryOperationsType
	CodeMilitaryOperationsTypeOrder            uint
	CodeMilitaryOperationsType_stagedOrder     map[*CodeMilitaryOperationsType]uint
	CodeMilitaryOperationsType_orderStaged     map[uint]*CodeMilitaryOperationsType
	CodeMilitaryOperationsTypes_reference      map[*CodeMilitaryOperationsType]*CodeMilitaryOperationsType
	CodeMilitaryOperationsTypes_referenceOrder map[*CodeMilitaryOperationsType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMilitaryOperationsTypeCreateCallback GongOnAfterCreateInterface[CodeMilitaryOperationsType]
	OnAfterCodeMilitaryOperationsTypeUpdateCallback GongOnAfterUpdateInterface[CodeMilitaryOperationsType]
	OnAfterCodeMilitaryOperationsTypeDeleteCallback GongOnAfterDeleteInterface[CodeMilitaryOperationsType]

	CodeNitrogenTypes                map[*CodeNitrogenType]struct{}
	CodeNitrogenTypes_instance       map[*CodeNitrogenType]*CodeNitrogenType
	CodeNitrogenTypes_mapString      map[string]*CodeNitrogenType
	CodeNitrogenTypeOrder            uint
	CodeNitrogenType_stagedOrder     map[*CodeNitrogenType]uint
	CodeNitrogenType_orderStaged     map[uint]*CodeNitrogenType
	CodeNitrogenTypes_reference      map[*CodeNitrogenType]*CodeNitrogenType
	CodeNitrogenTypes_referenceOrder map[*CodeNitrogenType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeNitrogenTypeCreateCallback GongOnAfterCreateInterface[CodeNitrogenType]
	OnAfterCodeNitrogenTypeUpdateCallback GongOnAfterUpdateInterface[CodeNitrogenType]
	OnAfterCodeNitrogenTypeDeleteCallback GongOnAfterDeleteInterface[CodeNitrogenType]

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

	CodeOilTypes                map[*CodeOilType]struct{}
	CodeOilTypes_instance       map[*CodeOilType]*CodeOilType
	CodeOilTypes_mapString      map[string]*CodeOilType
	CodeOilTypeOrder            uint
	CodeOilType_stagedOrder     map[*CodeOilType]uint
	CodeOilType_orderStaged     map[uint]*CodeOilType
	CodeOilTypes_reference      map[*CodeOilType]*CodeOilType
	CodeOilTypes_referenceOrder map[*CodeOilType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOilTypeCreateCallback GongOnAfterCreateInterface[CodeOilType]
	OnAfterCodeOilTypeUpdateCallback GongOnAfterUpdateInterface[CodeOilType]
	OnAfterCodeOilTypeDeleteCallback GongOnAfterDeleteInterface[CodeOilType]

	CodeOrganisationDesignatorTypes                map[*CodeOrganisationDesignatorType]struct{}
	CodeOrganisationDesignatorTypes_instance       map[*CodeOrganisationDesignatorType]*CodeOrganisationDesignatorType
	CodeOrganisationDesignatorTypes_mapString      map[string]*CodeOrganisationDesignatorType
	CodeOrganisationDesignatorTypeOrder            uint
	CodeOrganisationDesignatorType_stagedOrder     map[*CodeOrganisationDesignatorType]uint
	CodeOrganisationDesignatorType_orderStaged     map[uint]*CodeOrganisationDesignatorType
	CodeOrganisationDesignatorTypes_reference      map[*CodeOrganisationDesignatorType]*CodeOrganisationDesignatorType
	CodeOrganisationDesignatorTypes_referenceOrder map[*CodeOrganisationDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOrganisationDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeOrganisationDesignatorType]
	OnAfterCodeOrganisationDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeOrganisationDesignatorType]
	OnAfterCodeOrganisationDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeOrganisationDesignatorType]

	CodeOrganisationHierarchyTypes                map[*CodeOrganisationHierarchyType]struct{}
	CodeOrganisationHierarchyTypes_instance       map[*CodeOrganisationHierarchyType]*CodeOrganisationHierarchyType
	CodeOrganisationHierarchyTypes_mapString      map[string]*CodeOrganisationHierarchyType
	CodeOrganisationHierarchyTypeOrder            uint
	CodeOrganisationHierarchyType_stagedOrder     map[*CodeOrganisationHierarchyType]uint
	CodeOrganisationHierarchyType_orderStaged     map[uint]*CodeOrganisationHierarchyType
	CodeOrganisationHierarchyTypes_reference      map[*CodeOrganisationHierarchyType]*CodeOrganisationHierarchyType
	CodeOrganisationHierarchyTypes_referenceOrder map[*CodeOrganisationHierarchyType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOrganisationHierarchyTypeCreateCallback GongOnAfterCreateInterface[CodeOrganisationHierarchyType]
	OnAfterCodeOrganisationHierarchyTypeUpdateCallback GongOnAfterUpdateInterface[CodeOrganisationHierarchyType]
	OnAfterCodeOrganisationHierarchyTypeDeleteCallback GongOnAfterDeleteInterface[CodeOrganisationHierarchyType]

	CodeOrganisationTypes                map[*CodeOrganisationType]struct{}
	CodeOrganisationTypes_instance       map[*CodeOrganisationType]*CodeOrganisationType
	CodeOrganisationTypes_mapString      map[string]*CodeOrganisationType
	CodeOrganisationTypeOrder            uint
	CodeOrganisationType_stagedOrder     map[*CodeOrganisationType]uint
	CodeOrganisationType_orderStaged     map[uint]*CodeOrganisationType
	CodeOrganisationTypes_reference      map[*CodeOrganisationType]*CodeOrganisationType
	CodeOrganisationTypes_referenceOrder map[*CodeOrganisationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOrganisationTypeCreateCallback GongOnAfterCreateInterface[CodeOrganisationType]
	OnAfterCodeOrganisationTypeUpdateCallback GongOnAfterUpdateInterface[CodeOrganisationType]
	OnAfterCodeOrganisationTypeDeleteCallback GongOnAfterDeleteInterface[CodeOrganisationType]

	CodeOxygenTypes                map[*CodeOxygenType]struct{}
	CodeOxygenTypes_instance       map[*CodeOxygenType]*CodeOxygenType
	CodeOxygenTypes_mapString      map[string]*CodeOxygenType
	CodeOxygenTypeOrder            uint
	CodeOxygenType_stagedOrder     map[*CodeOxygenType]uint
	CodeOxygenType_orderStaged     map[uint]*CodeOxygenType
	CodeOxygenTypes_reference      map[*CodeOxygenType]*CodeOxygenType
	CodeOxygenTypes_referenceOrder map[*CodeOxygenType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOxygenTypeCreateCallback GongOnAfterCreateInterface[CodeOxygenType]
	OnAfterCodeOxygenTypeUpdateCallback GongOnAfterUpdateInterface[CodeOxygenType]
	OnAfterCodeOxygenTypeDeleteCallback GongOnAfterDeleteInterface[CodeOxygenType]

	CodePassengerServiceTypes                map[*CodePassengerServiceType]struct{}
	CodePassengerServiceTypes_instance       map[*CodePassengerServiceType]*CodePassengerServiceType
	CodePassengerServiceTypes_mapString      map[string]*CodePassengerServiceType
	CodePassengerServiceTypeOrder            uint
	CodePassengerServiceType_stagedOrder     map[*CodePassengerServiceType]uint
	CodePassengerServiceType_orderStaged     map[uint]*CodePassengerServiceType
	CodePassengerServiceTypes_reference      map[*CodePassengerServiceType]*CodePassengerServiceType
	CodePassengerServiceTypes_referenceOrder map[*CodePassengerServiceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePassengerServiceTypeCreateCallback GongOnAfterCreateInterface[CodePassengerServiceType]
	OnAfterCodePassengerServiceTypeUpdateCallback GongOnAfterUpdateInterface[CodePassengerServiceType]
	OnAfterCodePassengerServiceTypeDeleteCallback GongOnAfterDeleteInterface[CodePassengerServiceType]

	CodePilotControlledLightingTypes                map[*CodePilotControlledLightingType]struct{}
	CodePilotControlledLightingTypes_instance       map[*CodePilotControlledLightingType]*CodePilotControlledLightingType
	CodePilotControlledLightingTypes_mapString      map[string]*CodePilotControlledLightingType
	CodePilotControlledLightingTypeOrder            uint
	CodePilotControlledLightingType_stagedOrder     map[*CodePilotControlledLightingType]uint
	CodePilotControlledLightingType_orderStaged     map[uint]*CodePilotControlledLightingType
	CodePilotControlledLightingTypes_reference      map[*CodePilotControlledLightingType]*CodePilotControlledLightingType
	CodePilotControlledLightingTypes_referenceOrder map[*CodePilotControlledLightingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePilotControlledLightingTypeCreateCallback GongOnAfterCreateInterface[CodePilotControlledLightingType]
	OnAfterCodePilotControlledLightingTypeUpdateCallback GongOnAfterUpdateInterface[CodePilotControlledLightingType]
	OnAfterCodePilotControlledLightingTypeDeleteCallback GongOnAfterDeleteInterface[CodePilotControlledLightingType]

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

	CodeServiceATCTypes                map[*CodeServiceATCType]struct{}
	CodeServiceATCTypes_instance       map[*CodeServiceATCType]*CodeServiceATCType
	CodeServiceATCTypes_mapString      map[string]*CodeServiceATCType
	CodeServiceATCTypeOrder            uint
	CodeServiceATCType_stagedOrder     map[*CodeServiceATCType]uint
	CodeServiceATCType_orderStaged     map[uint]*CodeServiceATCType
	CodeServiceATCTypes_reference      map[*CodeServiceATCType]*CodeServiceATCType
	CodeServiceATCTypes_referenceOrder map[*CodeServiceATCType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeServiceATCTypeCreateCallback GongOnAfterCreateInterface[CodeServiceATCType]
	OnAfterCodeServiceATCTypeUpdateCallback GongOnAfterUpdateInterface[CodeServiceATCType]
	OnAfterCodeServiceATCTypeDeleteCallback GongOnAfterDeleteInterface[CodeServiceATCType]

	CodeServiceATFMTypes                map[*CodeServiceATFMType]struct{}
	CodeServiceATFMTypes_instance       map[*CodeServiceATFMType]*CodeServiceATFMType
	CodeServiceATFMTypes_mapString      map[string]*CodeServiceATFMType
	CodeServiceATFMTypeOrder            uint
	CodeServiceATFMType_stagedOrder     map[*CodeServiceATFMType]uint
	CodeServiceATFMType_orderStaged     map[uint]*CodeServiceATFMType
	CodeServiceATFMTypes_reference      map[*CodeServiceATFMType]*CodeServiceATFMType
	CodeServiceATFMTypes_referenceOrder map[*CodeServiceATFMType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeServiceATFMTypeCreateCallback GongOnAfterCreateInterface[CodeServiceATFMType]
	OnAfterCodeServiceATFMTypeUpdateCallback GongOnAfterUpdateInterface[CodeServiceATFMType]
	OnAfterCodeServiceATFMTypeDeleteCallback GongOnAfterDeleteInterface[CodeServiceATFMType]

	CodeServiceGroundControlTypes                map[*CodeServiceGroundControlType]struct{}
	CodeServiceGroundControlTypes_instance       map[*CodeServiceGroundControlType]*CodeServiceGroundControlType
	CodeServiceGroundControlTypes_mapString      map[string]*CodeServiceGroundControlType
	CodeServiceGroundControlTypeOrder            uint
	CodeServiceGroundControlType_stagedOrder     map[*CodeServiceGroundControlType]uint
	CodeServiceGroundControlType_orderStaged     map[uint]*CodeServiceGroundControlType
	CodeServiceGroundControlTypes_reference      map[*CodeServiceGroundControlType]*CodeServiceGroundControlType
	CodeServiceGroundControlTypes_referenceOrder map[*CodeServiceGroundControlType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeServiceGroundControlTypeCreateCallback GongOnAfterCreateInterface[CodeServiceGroundControlType]
	OnAfterCodeServiceGroundControlTypeUpdateCallback GongOnAfterUpdateInterface[CodeServiceGroundControlType]
	OnAfterCodeServiceGroundControlTypeDeleteCallback GongOnAfterDeleteInterface[CodeServiceGroundControlType]

	CodeServiceInformationTypes                map[*CodeServiceInformationType]struct{}
	CodeServiceInformationTypes_instance       map[*CodeServiceInformationType]*CodeServiceInformationType
	CodeServiceInformationTypes_mapString      map[string]*CodeServiceInformationType
	CodeServiceInformationTypeOrder            uint
	CodeServiceInformationType_stagedOrder     map[*CodeServiceInformationType]uint
	CodeServiceInformationType_orderStaged     map[uint]*CodeServiceInformationType
	CodeServiceInformationTypes_reference      map[*CodeServiceInformationType]*CodeServiceInformationType
	CodeServiceInformationTypes_referenceOrder map[*CodeServiceInformationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeServiceInformationTypeCreateCallback GongOnAfterCreateInterface[CodeServiceInformationType]
	OnAfterCodeServiceInformationTypeUpdateCallback GongOnAfterUpdateInterface[CodeServiceInformationType]
	OnAfterCodeServiceInformationTypeDeleteCallback GongOnAfterDeleteInterface[CodeServiceInformationType]

	CodeServiceSARTypes                map[*CodeServiceSARType]struct{}
	CodeServiceSARTypes_instance       map[*CodeServiceSARType]*CodeServiceSARType
	CodeServiceSARTypes_mapString      map[string]*CodeServiceSARType
	CodeServiceSARTypeOrder            uint
	CodeServiceSARType_stagedOrder     map[*CodeServiceSARType]uint
	CodeServiceSARType_orderStaged     map[uint]*CodeServiceSARType
	CodeServiceSARTypes_reference      map[*CodeServiceSARType]*CodeServiceSARType
	CodeServiceSARTypes_referenceOrder map[*CodeServiceSARType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeServiceSARTypeCreateCallback GongOnAfterCreateInterface[CodeServiceSARType]
	OnAfterCodeServiceSARTypeUpdateCallback GongOnAfterUpdateInterface[CodeServiceSARType]
	OnAfterCodeServiceSARTypeDeleteCallback GongOnAfterDeleteInterface[CodeServiceSARType]

	CodeSpecialDateTypes                map[*CodeSpecialDateType]struct{}
	CodeSpecialDateTypes_instance       map[*CodeSpecialDateType]*CodeSpecialDateType
	CodeSpecialDateTypes_mapString      map[string]*CodeSpecialDateType
	CodeSpecialDateTypeOrder            uint
	CodeSpecialDateType_stagedOrder     map[*CodeSpecialDateType]uint
	CodeSpecialDateType_orderStaged     map[uint]*CodeSpecialDateType
	CodeSpecialDateTypes_reference      map[*CodeSpecialDateType]*CodeSpecialDateType
	CodeSpecialDateTypes_referenceOrder map[*CodeSpecialDateType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSpecialDateTypeCreateCallback GongOnAfterCreateInterface[CodeSpecialDateType]
	OnAfterCodeSpecialDateTypeUpdateCallback GongOnAfterUpdateInterface[CodeSpecialDateType]
	OnAfterCodeSpecialDateTypeDeleteCallback GongOnAfterDeleteInterface[CodeSpecialDateType]

	CodeStatusOperationsTypes                map[*CodeStatusOperationsType]struct{}
	CodeStatusOperationsTypes_instance       map[*CodeStatusOperationsType]*CodeStatusOperationsType
	CodeStatusOperationsTypes_mapString      map[string]*CodeStatusOperationsType
	CodeStatusOperationsTypeOrder            uint
	CodeStatusOperationsType_stagedOrder     map[*CodeStatusOperationsType]uint
	CodeStatusOperationsType_orderStaged     map[uint]*CodeStatusOperationsType
	CodeStatusOperationsTypes_reference      map[*CodeStatusOperationsType]*CodeStatusOperationsType
	CodeStatusOperationsTypes_referenceOrder map[*CodeStatusOperationsType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusOperationsTypeCreateCallback GongOnAfterCreateInterface[CodeStatusOperationsType]
	OnAfterCodeStatusOperationsTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusOperationsType]
	OnAfterCodeStatusOperationsTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusOperationsType]

	CodeStatusServiceTypes                map[*CodeStatusServiceType]struct{}
	CodeStatusServiceTypes_instance       map[*CodeStatusServiceType]*CodeStatusServiceType
	CodeStatusServiceTypes_mapString      map[string]*CodeStatusServiceType
	CodeStatusServiceTypeOrder            uint
	CodeStatusServiceType_stagedOrder     map[*CodeStatusServiceType]uint
	CodeStatusServiceType_orderStaged     map[uint]*CodeStatusServiceType
	CodeStatusServiceTypes_reference      map[*CodeStatusServiceType]*CodeStatusServiceType
	CodeStatusServiceTypes_referenceOrder map[*CodeStatusServiceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusServiceTypeCreateCallback GongOnAfterCreateInterface[CodeStatusServiceType]
	OnAfterCodeStatusServiceTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusServiceType]
	OnAfterCodeStatusServiceTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusServiceType]

	CodeSystemActivationTypes                map[*CodeSystemActivationType]struct{}
	CodeSystemActivationTypes_instance       map[*CodeSystemActivationType]*CodeSystemActivationType
	CodeSystemActivationTypes_mapString      map[string]*CodeSystemActivationType
	CodeSystemActivationTypeOrder            uint
	CodeSystemActivationType_stagedOrder     map[*CodeSystemActivationType]uint
	CodeSystemActivationType_orderStaged     map[uint]*CodeSystemActivationType
	CodeSystemActivationTypes_reference      map[*CodeSystemActivationType]*CodeSystemActivationType
	CodeSystemActivationTypes_referenceOrder map[*CodeSystemActivationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSystemActivationTypeCreateCallback GongOnAfterCreateInterface[CodeSystemActivationType]
	OnAfterCodeSystemActivationTypeUpdateCallback GongOnAfterUpdateInterface[CodeSystemActivationType]
	OnAfterCodeSystemActivationTypeDeleteCallback GongOnAfterDeleteInterface[CodeSystemActivationType]

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

	CodeUnitDependencyTypes                map[*CodeUnitDependencyType]struct{}
	CodeUnitDependencyTypes_instance       map[*CodeUnitDependencyType]*CodeUnitDependencyType
	CodeUnitDependencyTypes_mapString      map[string]*CodeUnitDependencyType
	CodeUnitDependencyTypeOrder            uint
	CodeUnitDependencyType_stagedOrder     map[*CodeUnitDependencyType]uint
	CodeUnitDependencyType_orderStaged     map[uint]*CodeUnitDependencyType
	CodeUnitDependencyTypes_reference      map[*CodeUnitDependencyType]*CodeUnitDependencyType
	CodeUnitDependencyTypes_referenceOrder map[*CodeUnitDependencyType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeUnitDependencyTypeCreateCallback GongOnAfterCreateInterface[CodeUnitDependencyType]
	OnAfterCodeUnitDependencyTypeUpdateCallback GongOnAfterUpdateInterface[CodeUnitDependencyType]
	OnAfterCodeUnitDependencyTypeDeleteCallback GongOnAfterDeleteInterface[CodeUnitDependencyType]

	CodeUnitTypes                map[*CodeUnitType]struct{}
	CodeUnitTypes_instance       map[*CodeUnitType]*CodeUnitType
	CodeUnitTypes_mapString      map[string]*CodeUnitType
	CodeUnitTypeOrder            uint
	CodeUnitType_stagedOrder     map[*CodeUnitType]uint
	CodeUnitType_orderStaged     map[uint]*CodeUnitType
	CodeUnitTypes_reference      map[*CodeUnitType]*CodeUnitType
	CodeUnitTypes_referenceOrder map[*CodeUnitType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeUnitTypeCreateCallback GongOnAfterCreateInterface[CodeUnitType]
	OnAfterCodeUnitTypeUpdateCallback GongOnAfterUpdateInterface[CodeUnitType]
	OnAfterCodeUnitTypeDeleteCallback GongOnAfterDeleteInterface[CodeUnitType]

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
	__gong__clearReferences(&stage.CodeAircraftGroundServiceTypes_reference, &stage.CodeAircraftGroundServiceTypes_instance, &stage.CodeAircraftGroundServiceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAviationStandardsTypes_reference, &stage.CodeAviationStandardsTypes_instance, &stage.CodeAviationStandardsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationChannelTypes_reference, &stage.CodeCommunicationChannelTypes_instance, &stage.CodeCommunicationChannelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationDirectionTypes_reference, &stage.CodeCommunicationDirectionTypes_instance, &stage.CodeCommunicationDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_instance, &stage.CodeCommunicationModeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFacilityRankingTypes_reference, &stage.CodeFacilityRankingTypes_instance, &stage.CodeFacilityRankingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFireFightingTypes_reference, &stage.CodeFireFightingTypes_instance, &stage.CodeFireFightingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightDestinationTypes_reference, &stage.CodeFlightDestinationTypes_instance, &stage.CodeFlightDestinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_instance, &stage.CodeFlightRuleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFuelTypes_reference, &stage.CodeFuelTypes_instance, &stage.CodeFuelTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeIntensityStandByTypes_reference, &stage.CodeIntensityStandByTypes_instance, &stage.CodeIntensityStandByTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLanguageTypes_reference, &stage.CodeLanguageTypes_instance, &stage.CodeLanguageTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLevelSeriesTypes_reference, &stage.CodeLevelSeriesTypes_instance, &stage.CodeLevelSeriesTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLevelTableDesignatorTypes_reference, &stage.CodeLevelTableDesignatorTypes_instance, &stage.CodeLevelTableDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_instance, &stage.CodeLightIntensityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_instance, &stage.CodeMilitaryOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNitrogenTypes_reference, &stage.CodeNitrogenTypes_instance, &stage.CodeNitrogenTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNorthReferenceTypes_reference, &stage.CodeNorthReferenceTypes_instance, &stage.CodeNorthReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOilTypes_reference, &stage.CodeOilTypes_instance, &stage.CodeOilTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOrganisationDesignatorTypes_reference, &stage.CodeOrganisationDesignatorTypes_instance, &stage.CodeOrganisationDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOrganisationHierarchyTypes_reference, &stage.CodeOrganisationHierarchyTypes_instance, &stage.CodeOrganisationHierarchyTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOrganisationTypes_reference, &stage.CodeOrganisationTypes_instance, &stage.CodeOrganisationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOxygenTypes_reference, &stage.CodeOxygenTypes_instance, &stage.CodeOxygenTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePassengerServiceTypes_reference, &stage.CodePassengerServiceTypes_instance, &stage.CodePassengerServiceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePilotControlledLightingTypes_reference, &stage.CodePilotControlledLightingTypes_instance, &stage.CodePilotControlledLightingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_instance, &stage.CodeRVSMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRadioEmissionTypes_reference, &stage.CodeRadioEmissionTypes_instance, &stage.CodeRadioEmissionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeServiceATCTypes_reference, &stage.CodeServiceATCTypes_instance, &stage.CodeServiceATCTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeServiceATFMTypes_reference, &stage.CodeServiceATFMTypes_instance, &stage.CodeServiceATFMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeServiceGroundControlTypes_reference, &stage.CodeServiceGroundControlTypes_instance, &stage.CodeServiceGroundControlTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeServiceInformationTypes_reference, &stage.CodeServiceInformationTypes_instance, &stage.CodeServiceInformationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeServiceSARTypes_reference, &stage.CodeServiceSARTypes_instance, &stage.CodeServiceSARTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSpecialDateTypes_reference, &stage.CodeSpecialDateTypes_instance, &stage.CodeSpecialDateTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_instance, &stage.CodeStatusOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusServiceTypes_reference, &stage.CodeStatusServiceTypes_instance, &stage.CodeStatusServiceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSystemActivationTypes_reference, &stage.CodeSystemActivationTypes_instance, &stage.CodeSystemActivationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_instance, &stage.CodeTelecomNetworkTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeUnitDependencyTypes_reference, &stage.CodeUnitDependencyTypes_instance, &stage.CodeUnitDependencyTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeUnitTypes_reference, &stage.CodeUnitTypes_instance, &stage.CodeUnitTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.DateYearTypes_reference, &stage.DateYearTypes_instance, &stage.DateYearTypes_referenceOrder)

	__gong__clearReferences(&stage.NoNumberTypes_reference, &stage.NoNumberTypes_instance, &stage.NoNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.TextAddressTypes_reference, &stage.TextAddressTypes_instance, &stage.TextAddressTypes_referenceOrder)

	__gong__clearReferences(&stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_instance, &stage.TextDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_instance, &stage.TextInstructionTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNameTypes_reference, &stage.TextNameTypes_instance, &stage.TextNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNoteTypes_reference, &stage.TextNoteTypes_instance, &stage.TextNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPhoneTypes_reference, &stage.TextPhoneTypes_instance, &stage.TextPhoneTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_instance, &stage.TextPropertyNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValBearingTypes_reference, &stage.ValBearingTypes_instance, &stage.ValBearingTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValFrequencyTypes_reference, &stage.ValFrequencyTypes_instance, &stage.ValFrequencyTypes_referenceOrder)

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
	stage.CodeAircraftGroundServiceTypeOrder = __gong__recomputeOrder(stage.CodeAircraftGroundServiceType_stagedOrder)

	stage.CodeAviationStandardsTypeOrder = __gong__recomputeOrder(stage.CodeAviationStandardsType_stagedOrder)

	stage.CodeCommunicationChannelTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationChannelType_stagedOrder)

	stage.CodeCommunicationDirectionTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationDirectionType_stagedOrder)

	stage.CodeCommunicationModeTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationModeType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeFacilityRankingTypeOrder = __gong__recomputeOrder(stage.CodeFacilityRankingType_stagedOrder)

	stage.CodeFireFightingTypeOrder = __gong__recomputeOrder(stage.CodeFireFightingType_stagedOrder)

	stage.CodeFlightDestinationTypeOrder = __gong__recomputeOrder(stage.CodeFlightDestinationType_stagedOrder)

	stage.CodeFlightRuleTypeOrder = __gong__recomputeOrder(stage.CodeFlightRuleType_stagedOrder)

	stage.CodeFuelTypeOrder = __gong__recomputeOrder(stage.CodeFuelType_stagedOrder)

	stage.CodeIntensityStandByTypeOrder = __gong__recomputeOrder(stage.CodeIntensityStandByType_stagedOrder)

	stage.CodeLanguageTypeOrder = __gong__recomputeOrder(stage.CodeLanguageType_stagedOrder)

	stage.CodeLevelSeriesTypeOrder = __gong__recomputeOrder(stage.CodeLevelSeriesType_stagedOrder)

	stage.CodeLevelTableDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeLevelTableDesignatorType_stagedOrder)

	stage.CodeLightIntensityTypeOrder = __gong__recomputeOrder(stage.CodeLightIntensityType_stagedOrder)

	stage.CodeMilitaryOperationsTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryOperationsType_stagedOrder)

	stage.CodeNitrogenTypeOrder = __gong__recomputeOrder(stage.CodeNitrogenType_stagedOrder)

	stage.CodeNorthReferenceTypeOrder = __gong__recomputeOrder(stage.CodeNorthReferenceType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeOilTypeOrder = __gong__recomputeOrder(stage.CodeOilType_stagedOrder)

	stage.CodeOrganisationDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeOrganisationDesignatorType_stagedOrder)

	stage.CodeOrganisationHierarchyTypeOrder = __gong__recomputeOrder(stage.CodeOrganisationHierarchyType_stagedOrder)

	stage.CodeOrganisationTypeOrder = __gong__recomputeOrder(stage.CodeOrganisationType_stagedOrder)

	stage.CodeOxygenTypeOrder = __gong__recomputeOrder(stage.CodeOxygenType_stagedOrder)

	stage.CodePassengerServiceTypeOrder = __gong__recomputeOrder(stage.CodePassengerServiceType_stagedOrder)

	stage.CodePilotControlledLightingTypeOrder = __gong__recomputeOrder(stage.CodePilotControlledLightingType_stagedOrder)

	stage.CodeRVSMTypeOrder = __gong__recomputeOrder(stage.CodeRVSMType_stagedOrder)

	stage.CodeRadioEmissionTypeOrder = __gong__recomputeOrder(stage.CodeRadioEmissionType_stagedOrder)

	stage.CodeServiceATCTypeOrder = __gong__recomputeOrder(stage.CodeServiceATCType_stagedOrder)

	stage.CodeServiceATFMTypeOrder = __gong__recomputeOrder(stage.CodeServiceATFMType_stagedOrder)

	stage.CodeServiceGroundControlTypeOrder = __gong__recomputeOrder(stage.CodeServiceGroundControlType_stagedOrder)

	stage.CodeServiceInformationTypeOrder = __gong__recomputeOrder(stage.CodeServiceInformationType_stagedOrder)

	stage.CodeServiceSARTypeOrder = __gong__recomputeOrder(stage.CodeServiceSARType_stagedOrder)

	stage.CodeSpecialDateTypeOrder = __gong__recomputeOrder(stage.CodeSpecialDateType_stagedOrder)

	stage.CodeStatusOperationsTypeOrder = __gong__recomputeOrder(stage.CodeStatusOperationsType_stagedOrder)

	stage.CodeStatusServiceTypeOrder = __gong__recomputeOrder(stage.CodeStatusServiceType_stagedOrder)

	stage.CodeSystemActivationTypeOrder = __gong__recomputeOrder(stage.CodeSystemActivationType_stagedOrder)

	stage.CodeTelecomNetworkTypeOrder = __gong__recomputeOrder(stage.CodeTelecomNetworkType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeUnitDependencyTypeOrder = __gong__recomputeOrder(stage.CodeUnitDependencyType_stagedOrder)

	stage.CodeUnitTypeOrder = __gong__recomputeOrder(stage.CodeUnitType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.DateYearTypeOrder = __gong__recomputeOrder(stage.DateYearType_stagedOrder)

	stage.NoNumberTypeOrder = __gong__recomputeOrder(stage.NoNumberType_stagedOrder)

	stage.TextAddressTypeOrder = __gong__recomputeOrder(stage.TextAddressType_stagedOrder)

	stage.TextDesignatorTypeOrder = __gong__recomputeOrder(stage.TextDesignatorType_stagedOrder)

	stage.TextInstructionTypeOrder = __gong__recomputeOrder(stage.TextInstructionType_stagedOrder)

	stage.TextNameTypeOrder = __gong__recomputeOrder(stage.TextNameType_stagedOrder)

	stage.TextNoteTypeOrder = __gong__recomputeOrder(stage.TextNoteType_stagedOrder)

	stage.TextPhoneTypeOrder = __gong__recomputeOrder(stage.TextPhoneType_stagedOrder)

	stage.TextPropertyNameTypeOrder = __gong__recomputeOrder(stage.TextPropertyNameType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.ValBearingTypeOrder = __gong__recomputeOrder(stage.ValBearingType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValFrequencyTypeOrder = __gong__recomputeOrder(stage.ValFrequencyType_stagedOrder)

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
func (*CodeAircraftGroundServiceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceType_stagedOrder)
}

func (*CodeAircraftGroundServiceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftGroundServiceType_orderStaged[order]
}

func (*CodeAircraftGroundServiceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftGroundServiceTypes_mapString
}

func (*CodeAircraftGroundServiceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftGroundServiceTypes
}

func (*CodeAircraftGroundServiceType) GongNewInstance() any {
	return new(CodeAircraftGroundServiceType)
}

func (*CodeAviationStandardsType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsType_stagedOrder)
}

func (*CodeAviationStandardsType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAviationStandardsType_orderStaged[order]
}

func (*CodeAviationStandardsType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAviationStandardsTypes_mapString
}

func (*CodeAviationStandardsType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAviationStandardsTypes
}

func (*CodeAviationStandardsType) GongNewInstance() any {
	return new(CodeAviationStandardsType)
}

func (*CodeCommunicationChannelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelType_stagedOrder)
}

func (*CodeCommunicationChannelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCommunicationChannelType_orderStaged[order]
}

func (*CodeCommunicationChannelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCommunicationChannelTypes_mapString
}

func (*CodeCommunicationChannelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCommunicationChannelTypes
}

func (*CodeCommunicationChannelType) GongNewInstance() any {
	return new(CodeCommunicationChannelType)
}

func (*CodeCommunicationDirectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionType_stagedOrder)
}

func (*CodeCommunicationDirectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeCommunicationDirectionType_orderStaged[order]
}

func (*CodeCommunicationDirectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeCommunicationDirectionTypes_mapString
}

func (*CodeCommunicationDirectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeCommunicationDirectionTypes
}

func (*CodeCommunicationDirectionType) GongNewInstance() any {
	return new(CodeCommunicationDirectionType)
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

func (*CodeFacilityRankingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingType_stagedOrder)
}

func (*CodeFacilityRankingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFacilityRankingType_orderStaged[order]
}

func (*CodeFacilityRankingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFacilityRankingTypes_mapString
}

func (*CodeFacilityRankingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFacilityRankingTypes
}

func (*CodeFacilityRankingType) GongNewInstance() any {
	return new(CodeFacilityRankingType)
}

func (*CodeFireFightingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFireFightingTypes, stage.CodeFireFightingType_stagedOrder)
}

func (*CodeFireFightingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFireFightingType_orderStaged[order]
}

func (*CodeFireFightingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFireFightingTypes_mapString
}

func (*CodeFireFightingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFireFightingTypes
}

func (*CodeFireFightingType) GongNewInstance() any {
	return new(CodeFireFightingType)
}

func (*CodeFlightDestinationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationType_stagedOrder)
}

func (*CodeFlightDestinationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFlightDestinationType_orderStaged[order]
}

func (*CodeFlightDestinationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFlightDestinationTypes_mapString
}

func (*CodeFlightDestinationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFlightDestinationTypes
}

func (*CodeFlightDestinationType) GongNewInstance() any {
	return new(CodeFlightDestinationType)
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

func (*CodeFuelType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFuelTypes, stage.CodeFuelType_stagedOrder)
}

func (*CodeFuelType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFuelType_orderStaged[order]
}

func (*CodeFuelType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFuelTypes_mapString
}

func (*CodeFuelType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFuelTypes
}

func (*CodeFuelType) GongNewInstance() any {
	return new(CodeFuelType)
}

func (*CodeIntensityStandByType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByType_stagedOrder)
}

func (*CodeIntensityStandByType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeIntensityStandByType_orderStaged[order]
}

func (*CodeIntensityStandByType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeIntensityStandByTypes_mapString
}

func (*CodeIntensityStandByType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeIntensityStandByTypes
}

func (*CodeIntensityStandByType) GongNewInstance() any {
	return new(CodeIntensityStandByType)
}

func (*CodeLanguageType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLanguageTypes, stage.CodeLanguageType_stagedOrder)
}

func (*CodeLanguageType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLanguageType_orderStaged[order]
}

func (*CodeLanguageType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLanguageTypes_mapString
}

func (*CodeLanguageType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLanguageTypes
}

func (*CodeLanguageType) GongNewInstance() any {
	return new(CodeLanguageType)
}

func (*CodeLevelSeriesType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesType_stagedOrder)
}

func (*CodeLevelSeriesType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLevelSeriesType_orderStaged[order]
}

func (*CodeLevelSeriesType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLevelSeriesTypes_mapString
}

func (*CodeLevelSeriesType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLevelSeriesTypes
}

func (*CodeLevelSeriesType) GongNewInstance() any {
	return new(CodeLevelSeriesType)
}

func (*CodeLevelTableDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorType_stagedOrder)
}

func (*CodeLevelTableDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLevelTableDesignatorType_orderStaged[order]
}

func (*CodeLevelTableDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLevelTableDesignatorTypes_mapString
}

func (*CodeLevelTableDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLevelTableDesignatorTypes
}

func (*CodeLevelTableDesignatorType) GongNewInstance() any {
	return new(CodeLevelTableDesignatorType)
}

func (*CodeLightIntensityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLightIntensityTypes, stage.CodeLightIntensityType_stagedOrder)
}

func (*CodeLightIntensityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLightIntensityType_orderStaged[order]
}

func (*CodeLightIntensityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLightIntensityTypes_mapString
}

func (*CodeLightIntensityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLightIntensityTypes
}

func (*CodeLightIntensityType) GongNewInstance() any {
	return new(CodeLightIntensityType)
}

func (*CodeMilitaryOperationsType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsType_stagedOrder)
}

func (*CodeMilitaryOperationsType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMilitaryOperationsType_orderStaged[order]
}

func (*CodeMilitaryOperationsType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMilitaryOperationsTypes_mapString
}

func (*CodeMilitaryOperationsType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMilitaryOperationsTypes
}

func (*CodeMilitaryOperationsType) GongNewInstance() any {
	return new(CodeMilitaryOperationsType)
}

func (*CodeNitrogenType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeNitrogenTypes, stage.CodeNitrogenType_stagedOrder)
}

func (*CodeNitrogenType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeNitrogenType_orderStaged[order]
}

func (*CodeNitrogenType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeNitrogenTypes_mapString
}

func (*CodeNitrogenType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeNitrogenTypes
}

func (*CodeNitrogenType) GongNewInstance() any {
	return new(CodeNitrogenType)
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

func (*CodeOilType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOilTypes, stage.CodeOilType_stagedOrder)
}

func (*CodeOilType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOilType_orderStaged[order]
}

func (*CodeOilType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOilTypes_mapString
}

func (*CodeOilType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOilTypes
}

func (*CodeOilType) GongNewInstance() any {
	return new(CodeOilType)
}

func (*CodeOrganisationDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorType_stagedOrder)
}

func (*CodeOrganisationDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOrganisationDesignatorType_orderStaged[order]
}

func (*CodeOrganisationDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOrganisationDesignatorTypes_mapString
}

func (*CodeOrganisationDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOrganisationDesignatorTypes
}

func (*CodeOrganisationDesignatorType) GongNewInstance() any {
	return new(CodeOrganisationDesignatorType)
}

func (*CodeOrganisationHierarchyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyType_stagedOrder)
}

func (*CodeOrganisationHierarchyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOrganisationHierarchyType_orderStaged[order]
}

func (*CodeOrganisationHierarchyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOrganisationHierarchyTypes_mapString
}

func (*CodeOrganisationHierarchyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOrganisationHierarchyTypes
}

func (*CodeOrganisationHierarchyType) GongNewInstance() any {
	return new(CodeOrganisationHierarchyType)
}

func (*CodeOrganisationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOrganisationTypes, stage.CodeOrganisationType_stagedOrder)
}

func (*CodeOrganisationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOrganisationType_orderStaged[order]
}

func (*CodeOrganisationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOrganisationTypes_mapString
}

func (*CodeOrganisationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOrganisationTypes
}

func (*CodeOrganisationType) GongNewInstance() any {
	return new(CodeOrganisationType)
}

func (*CodeOxygenType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOxygenTypes, stage.CodeOxygenType_stagedOrder)
}

func (*CodeOxygenType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOxygenType_orderStaged[order]
}

func (*CodeOxygenType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOxygenTypes_mapString
}

func (*CodeOxygenType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOxygenTypes
}

func (*CodeOxygenType) GongNewInstance() any {
	return new(CodeOxygenType)
}

func (*CodePassengerServiceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePassengerServiceTypes, stage.CodePassengerServiceType_stagedOrder)
}

func (*CodePassengerServiceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePassengerServiceType_orderStaged[order]
}

func (*CodePassengerServiceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePassengerServiceTypes_mapString
}

func (*CodePassengerServiceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePassengerServiceTypes
}

func (*CodePassengerServiceType) GongNewInstance() any {
	return new(CodePassengerServiceType)
}

func (*CodePilotControlledLightingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingType_stagedOrder)
}

func (*CodePilotControlledLightingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePilotControlledLightingType_orderStaged[order]
}

func (*CodePilotControlledLightingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePilotControlledLightingTypes_mapString
}

func (*CodePilotControlledLightingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePilotControlledLightingTypes
}

func (*CodePilotControlledLightingType) GongNewInstance() any {
	return new(CodePilotControlledLightingType)
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

func (*CodeServiceATCType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeServiceATCTypes, stage.CodeServiceATCType_stagedOrder)
}

func (*CodeServiceATCType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeServiceATCType_orderStaged[order]
}

func (*CodeServiceATCType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeServiceATCTypes_mapString
}

func (*CodeServiceATCType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeServiceATCTypes
}

func (*CodeServiceATCType) GongNewInstance() any {
	return new(CodeServiceATCType)
}

func (*CodeServiceATFMType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeServiceATFMTypes, stage.CodeServiceATFMType_stagedOrder)
}

func (*CodeServiceATFMType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeServiceATFMType_orderStaged[order]
}

func (*CodeServiceATFMType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeServiceATFMTypes_mapString
}

func (*CodeServiceATFMType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeServiceATFMTypes
}

func (*CodeServiceATFMType) GongNewInstance() any {
	return new(CodeServiceATFMType)
}

func (*CodeServiceGroundControlType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlType_stagedOrder)
}

func (*CodeServiceGroundControlType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeServiceGroundControlType_orderStaged[order]
}

func (*CodeServiceGroundControlType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeServiceGroundControlTypes_mapString
}

func (*CodeServiceGroundControlType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeServiceGroundControlTypes
}

func (*CodeServiceGroundControlType) GongNewInstance() any {
	return new(CodeServiceGroundControlType)
}

func (*CodeServiceInformationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeServiceInformationTypes, stage.CodeServiceInformationType_stagedOrder)
}

func (*CodeServiceInformationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeServiceInformationType_orderStaged[order]
}

func (*CodeServiceInformationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeServiceInformationTypes_mapString
}

func (*CodeServiceInformationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeServiceInformationTypes
}

func (*CodeServiceInformationType) GongNewInstance() any {
	return new(CodeServiceInformationType)
}

func (*CodeServiceSARType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeServiceSARTypes, stage.CodeServiceSARType_stagedOrder)
}

func (*CodeServiceSARType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeServiceSARType_orderStaged[order]
}

func (*CodeServiceSARType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeServiceSARTypes_mapString
}

func (*CodeServiceSARType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeServiceSARTypes
}

func (*CodeServiceSARType) GongNewInstance() any {
	return new(CodeServiceSARType)
}

func (*CodeSpecialDateType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSpecialDateTypes, stage.CodeSpecialDateType_stagedOrder)
}

func (*CodeSpecialDateType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSpecialDateType_orderStaged[order]
}

func (*CodeSpecialDateType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSpecialDateTypes_mapString
}

func (*CodeSpecialDateType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSpecialDateTypes
}

func (*CodeSpecialDateType) GongNewInstance() any {
	return new(CodeSpecialDateType)
}

func (*CodeStatusOperationsType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsType_stagedOrder)
}

func (*CodeStatusOperationsType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusOperationsType_orderStaged[order]
}

func (*CodeStatusOperationsType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusOperationsTypes_mapString
}

func (*CodeStatusOperationsType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusOperationsTypes
}

func (*CodeStatusOperationsType) GongNewInstance() any {
	return new(CodeStatusOperationsType)
}

func (*CodeStatusServiceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusServiceTypes, stage.CodeStatusServiceType_stagedOrder)
}

func (*CodeStatusServiceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusServiceType_orderStaged[order]
}

func (*CodeStatusServiceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusServiceTypes_mapString
}

func (*CodeStatusServiceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusServiceTypes
}

func (*CodeStatusServiceType) GongNewInstance() any {
	return new(CodeStatusServiceType)
}

func (*CodeSystemActivationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSystemActivationTypes, stage.CodeSystemActivationType_stagedOrder)
}

func (*CodeSystemActivationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSystemActivationType_orderStaged[order]
}

func (*CodeSystemActivationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSystemActivationTypes_mapString
}

func (*CodeSystemActivationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSystemActivationTypes
}

func (*CodeSystemActivationType) GongNewInstance() any {
	return new(CodeSystemActivationType)
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

func (*CodeUnitDependencyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyType_stagedOrder)
}

func (*CodeUnitDependencyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeUnitDependencyType_orderStaged[order]
}

func (*CodeUnitDependencyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeUnitDependencyTypes_mapString
}

func (*CodeUnitDependencyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeUnitDependencyTypes
}

func (*CodeUnitDependencyType) GongNewInstance() any {
	return new(CodeUnitDependencyType)
}

func (*CodeUnitType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeUnitTypes, stage.CodeUnitType_stagedOrder)
}

func (*CodeUnitType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeUnitType_orderStaged[order]
}

func (*CodeUnitType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeUnitTypes_mapString
}

func (*CodeUnitType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeUnitTypes
}

func (*CodeUnitType) GongNewInstance() any {
	return new(CodeUnitType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
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
		CodeAircraftGroundServiceTypes:           make(map[*CodeAircraftGroundServiceType]struct{}),
		CodeAircraftGroundServiceTypes_mapString: make(map[string]*CodeAircraftGroundServiceType),

		CodeAviationStandardsTypes:           make(map[*CodeAviationStandardsType]struct{}),
		CodeAviationStandardsTypes_mapString: make(map[string]*CodeAviationStandardsType),

		CodeCommunicationChannelTypes:           make(map[*CodeCommunicationChannelType]struct{}),
		CodeCommunicationChannelTypes_mapString: make(map[string]*CodeCommunicationChannelType),

		CodeCommunicationDirectionTypes:           make(map[*CodeCommunicationDirectionType]struct{}),
		CodeCommunicationDirectionTypes_mapString: make(map[string]*CodeCommunicationDirectionType),

		CodeCommunicationModeTypes:           make(map[*CodeCommunicationModeType]struct{}),
		CodeCommunicationModeTypes_mapString: make(map[string]*CodeCommunicationModeType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeFacilityRankingTypes:           make(map[*CodeFacilityRankingType]struct{}),
		CodeFacilityRankingTypes_mapString: make(map[string]*CodeFacilityRankingType),

		CodeFireFightingTypes:           make(map[*CodeFireFightingType]struct{}),
		CodeFireFightingTypes_mapString: make(map[string]*CodeFireFightingType),

		CodeFlightDestinationTypes:           make(map[*CodeFlightDestinationType]struct{}),
		CodeFlightDestinationTypes_mapString: make(map[string]*CodeFlightDestinationType),

		CodeFlightRuleTypes:           make(map[*CodeFlightRuleType]struct{}),
		CodeFlightRuleTypes_mapString: make(map[string]*CodeFlightRuleType),

		CodeFuelTypes:           make(map[*CodeFuelType]struct{}),
		CodeFuelTypes_mapString: make(map[string]*CodeFuelType),

		CodeIntensityStandByTypes:           make(map[*CodeIntensityStandByType]struct{}),
		CodeIntensityStandByTypes_mapString: make(map[string]*CodeIntensityStandByType),

		CodeLanguageTypes:           make(map[*CodeLanguageType]struct{}),
		CodeLanguageTypes_mapString: make(map[string]*CodeLanguageType),

		CodeLevelSeriesTypes:           make(map[*CodeLevelSeriesType]struct{}),
		CodeLevelSeriesTypes_mapString: make(map[string]*CodeLevelSeriesType),

		CodeLevelTableDesignatorTypes:           make(map[*CodeLevelTableDesignatorType]struct{}),
		CodeLevelTableDesignatorTypes_mapString: make(map[string]*CodeLevelTableDesignatorType),

		CodeLightIntensityTypes:           make(map[*CodeLightIntensityType]struct{}),
		CodeLightIntensityTypes_mapString: make(map[string]*CodeLightIntensityType),

		CodeMilitaryOperationsTypes:           make(map[*CodeMilitaryOperationsType]struct{}),
		CodeMilitaryOperationsTypes_mapString: make(map[string]*CodeMilitaryOperationsType),

		CodeNitrogenTypes:           make(map[*CodeNitrogenType]struct{}),
		CodeNitrogenTypes_mapString: make(map[string]*CodeNitrogenType),

		CodeNorthReferenceTypes:           make(map[*CodeNorthReferenceType]struct{}),
		CodeNorthReferenceTypes_mapString: make(map[string]*CodeNorthReferenceType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodeOilTypes:           make(map[*CodeOilType]struct{}),
		CodeOilTypes_mapString: make(map[string]*CodeOilType),

		CodeOrganisationDesignatorTypes:           make(map[*CodeOrganisationDesignatorType]struct{}),
		CodeOrganisationDesignatorTypes_mapString: make(map[string]*CodeOrganisationDesignatorType),

		CodeOrganisationHierarchyTypes:           make(map[*CodeOrganisationHierarchyType]struct{}),
		CodeOrganisationHierarchyTypes_mapString: make(map[string]*CodeOrganisationHierarchyType),

		CodeOrganisationTypes:           make(map[*CodeOrganisationType]struct{}),
		CodeOrganisationTypes_mapString: make(map[string]*CodeOrganisationType),

		CodeOxygenTypes:           make(map[*CodeOxygenType]struct{}),
		CodeOxygenTypes_mapString: make(map[string]*CodeOxygenType),

		CodePassengerServiceTypes:           make(map[*CodePassengerServiceType]struct{}),
		CodePassengerServiceTypes_mapString: make(map[string]*CodePassengerServiceType),

		CodePilotControlledLightingTypes:           make(map[*CodePilotControlledLightingType]struct{}),
		CodePilotControlledLightingTypes_mapString: make(map[string]*CodePilotControlledLightingType),

		CodeRVSMTypes:           make(map[*CodeRVSMType]struct{}),
		CodeRVSMTypes_mapString: make(map[string]*CodeRVSMType),

		CodeRadioEmissionTypes:           make(map[*CodeRadioEmissionType]struct{}),
		CodeRadioEmissionTypes_mapString: make(map[string]*CodeRadioEmissionType),

		CodeServiceATCTypes:           make(map[*CodeServiceATCType]struct{}),
		CodeServiceATCTypes_mapString: make(map[string]*CodeServiceATCType),

		CodeServiceATFMTypes:           make(map[*CodeServiceATFMType]struct{}),
		CodeServiceATFMTypes_mapString: make(map[string]*CodeServiceATFMType),

		CodeServiceGroundControlTypes:           make(map[*CodeServiceGroundControlType]struct{}),
		CodeServiceGroundControlTypes_mapString: make(map[string]*CodeServiceGroundControlType),

		CodeServiceInformationTypes:           make(map[*CodeServiceInformationType]struct{}),
		CodeServiceInformationTypes_mapString: make(map[string]*CodeServiceInformationType),

		CodeServiceSARTypes:           make(map[*CodeServiceSARType]struct{}),
		CodeServiceSARTypes_mapString: make(map[string]*CodeServiceSARType),

		CodeSpecialDateTypes:           make(map[*CodeSpecialDateType]struct{}),
		CodeSpecialDateTypes_mapString: make(map[string]*CodeSpecialDateType),

		CodeStatusOperationsTypes:           make(map[*CodeStatusOperationsType]struct{}),
		CodeStatusOperationsTypes_mapString: make(map[string]*CodeStatusOperationsType),

		CodeStatusServiceTypes:           make(map[*CodeStatusServiceType]struct{}),
		CodeStatusServiceTypes_mapString: make(map[string]*CodeStatusServiceType),

		CodeSystemActivationTypes:           make(map[*CodeSystemActivationType]struct{}),
		CodeSystemActivationTypes_mapString: make(map[string]*CodeSystemActivationType),

		CodeTelecomNetworkTypes:           make(map[*CodeTelecomNetworkType]struct{}),
		CodeTelecomNetworkTypes_mapString: make(map[string]*CodeTelecomNetworkType),

		CodeTimeEventCombinationTypes:           make(map[*CodeTimeEventCombinationType]struct{}),
		CodeTimeEventCombinationTypes_mapString: make(map[string]*CodeTimeEventCombinationType),

		CodeTimeEventTypes:           make(map[*CodeTimeEventType]struct{}),
		CodeTimeEventTypes_mapString: make(map[string]*CodeTimeEventType),

		CodeTimeReferenceTypes:           make(map[*CodeTimeReferenceType]struct{}),
		CodeTimeReferenceTypes_mapString: make(map[string]*CodeTimeReferenceType),

		CodeUnitDependencyTypes:           make(map[*CodeUnitDependencyType]struct{}),
		CodeUnitDependencyTypes_mapString: make(map[string]*CodeUnitDependencyType),

		CodeUnitTypes:           make(map[*CodeUnitType]struct{}),
		CodeUnitTypes_mapString: make(map[string]*CodeUnitType),

		CodeVerticalDatumTypes:           make(map[*CodeVerticalDatumType]struct{}),
		CodeVerticalDatumTypes_mapString: make(map[string]*CodeVerticalDatumType),

		CodeYesNoTypes:           make(map[*CodeYesNoType]struct{}),
		CodeYesNoTypes_mapString: make(map[string]*CodeYesNoType),

		DateMonthDayTypes:           make(map[*DateMonthDayType]struct{}),
		DateMonthDayTypes_mapString: make(map[string]*DateMonthDayType),

		DateYearTypes:           make(map[*DateYearType]struct{}),
		DateYearTypes_mapString: make(map[string]*DateYearType),

		NoNumberTypes:           make(map[*NoNumberType]struct{}),
		NoNumberTypes_mapString: make(map[string]*NoNumberType),

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

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CodeAircraftGroundServiceType_stagedOrder: make(map[*CodeAircraftGroundServiceType]uint),
		CodeAircraftGroundServiceType_orderStaged: make(map[uint]*CodeAircraftGroundServiceType),
		CodeAircraftGroundServiceTypes_reference:  make(map[*CodeAircraftGroundServiceType]*CodeAircraftGroundServiceType),

		CodeAviationStandardsType_stagedOrder: make(map[*CodeAviationStandardsType]uint),
		CodeAviationStandardsType_orderStaged: make(map[uint]*CodeAviationStandardsType),
		CodeAviationStandardsTypes_reference:  make(map[*CodeAviationStandardsType]*CodeAviationStandardsType),

		CodeCommunicationChannelType_stagedOrder: make(map[*CodeCommunicationChannelType]uint),
		CodeCommunicationChannelType_orderStaged: make(map[uint]*CodeCommunicationChannelType),
		CodeCommunicationChannelTypes_reference:  make(map[*CodeCommunicationChannelType]*CodeCommunicationChannelType),

		CodeCommunicationDirectionType_stagedOrder: make(map[*CodeCommunicationDirectionType]uint),
		CodeCommunicationDirectionType_orderStaged: make(map[uint]*CodeCommunicationDirectionType),
		CodeCommunicationDirectionTypes_reference:  make(map[*CodeCommunicationDirectionType]*CodeCommunicationDirectionType),

		CodeCommunicationModeType_stagedOrder: make(map[*CodeCommunicationModeType]uint),
		CodeCommunicationModeType_orderStaged: make(map[uint]*CodeCommunicationModeType),
		CodeCommunicationModeTypes_reference:  make(map[*CodeCommunicationModeType]*CodeCommunicationModeType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeFacilityRankingType_stagedOrder: make(map[*CodeFacilityRankingType]uint),
		CodeFacilityRankingType_orderStaged: make(map[uint]*CodeFacilityRankingType),
		CodeFacilityRankingTypes_reference:  make(map[*CodeFacilityRankingType]*CodeFacilityRankingType),

		CodeFireFightingType_stagedOrder: make(map[*CodeFireFightingType]uint),
		CodeFireFightingType_orderStaged: make(map[uint]*CodeFireFightingType),
		CodeFireFightingTypes_reference:  make(map[*CodeFireFightingType]*CodeFireFightingType),

		CodeFlightDestinationType_stagedOrder: make(map[*CodeFlightDestinationType]uint),
		CodeFlightDestinationType_orderStaged: make(map[uint]*CodeFlightDestinationType),
		CodeFlightDestinationTypes_reference:  make(map[*CodeFlightDestinationType]*CodeFlightDestinationType),

		CodeFlightRuleType_stagedOrder: make(map[*CodeFlightRuleType]uint),
		CodeFlightRuleType_orderStaged: make(map[uint]*CodeFlightRuleType),
		CodeFlightRuleTypes_reference:  make(map[*CodeFlightRuleType]*CodeFlightRuleType),

		CodeFuelType_stagedOrder: make(map[*CodeFuelType]uint),
		CodeFuelType_orderStaged: make(map[uint]*CodeFuelType),
		CodeFuelTypes_reference:  make(map[*CodeFuelType]*CodeFuelType),

		CodeIntensityStandByType_stagedOrder: make(map[*CodeIntensityStandByType]uint),
		CodeIntensityStandByType_orderStaged: make(map[uint]*CodeIntensityStandByType),
		CodeIntensityStandByTypes_reference:  make(map[*CodeIntensityStandByType]*CodeIntensityStandByType),

		CodeLanguageType_stagedOrder: make(map[*CodeLanguageType]uint),
		CodeLanguageType_orderStaged: make(map[uint]*CodeLanguageType),
		CodeLanguageTypes_reference:  make(map[*CodeLanguageType]*CodeLanguageType),

		CodeLevelSeriesType_stagedOrder: make(map[*CodeLevelSeriesType]uint),
		CodeLevelSeriesType_orderStaged: make(map[uint]*CodeLevelSeriesType),
		CodeLevelSeriesTypes_reference:  make(map[*CodeLevelSeriesType]*CodeLevelSeriesType),

		CodeLevelTableDesignatorType_stagedOrder: make(map[*CodeLevelTableDesignatorType]uint),
		CodeLevelTableDesignatorType_orderStaged: make(map[uint]*CodeLevelTableDesignatorType),
		CodeLevelTableDesignatorTypes_reference:  make(map[*CodeLevelTableDesignatorType]*CodeLevelTableDesignatorType),

		CodeLightIntensityType_stagedOrder: make(map[*CodeLightIntensityType]uint),
		CodeLightIntensityType_orderStaged: make(map[uint]*CodeLightIntensityType),
		CodeLightIntensityTypes_reference:  make(map[*CodeLightIntensityType]*CodeLightIntensityType),

		CodeMilitaryOperationsType_stagedOrder: make(map[*CodeMilitaryOperationsType]uint),
		CodeMilitaryOperationsType_orderStaged: make(map[uint]*CodeMilitaryOperationsType),
		CodeMilitaryOperationsTypes_reference:  make(map[*CodeMilitaryOperationsType]*CodeMilitaryOperationsType),

		CodeNitrogenType_stagedOrder: make(map[*CodeNitrogenType]uint),
		CodeNitrogenType_orderStaged: make(map[uint]*CodeNitrogenType),
		CodeNitrogenTypes_reference:  make(map[*CodeNitrogenType]*CodeNitrogenType),

		CodeNorthReferenceType_stagedOrder: make(map[*CodeNorthReferenceType]uint),
		CodeNorthReferenceType_orderStaged: make(map[uint]*CodeNorthReferenceType),
		CodeNorthReferenceTypes_reference:  make(map[*CodeNorthReferenceType]*CodeNorthReferenceType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodeOilType_stagedOrder: make(map[*CodeOilType]uint),
		CodeOilType_orderStaged: make(map[uint]*CodeOilType),
		CodeOilTypes_reference:  make(map[*CodeOilType]*CodeOilType),

		CodeOrganisationDesignatorType_stagedOrder: make(map[*CodeOrganisationDesignatorType]uint),
		CodeOrganisationDesignatorType_orderStaged: make(map[uint]*CodeOrganisationDesignatorType),
		CodeOrganisationDesignatorTypes_reference:  make(map[*CodeOrganisationDesignatorType]*CodeOrganisationDesignatorType),

		CodeOrganisationHierarchyType_stagedOrder: make(map[*CodeOrganisationHierarchyType]uint),
		CodeOrganisationHierarchyType_orderStaged: make(map[uint]*CodeOrganisationHierarchyType),
		CodeOrganisationHierarchyTypes_reference:  make(map[*CodeOrganisationHierarchyType]*CodeOrganisationHierarchyType),

		CodeOrganisationType_stagedOrder: make(map[*CodeOrganisationType]uint),
		CodeOrganisationType_orderStaged: make(map[uint]*CodeOrganisationType),
		CodeOrganisationTypes_reference:  make(map[*CodeOrganisationType]*CodeOrganisationType),

		CodeOxygenType_stagedOrder: make(map[*CodeOxygenType]uint),
		CodeOxygenType_orderStaged: make(map[uint]*CodeOxygenType),
		CodeOxygenTypes_reference:  make(map[*CodeOxygenType]*CodeOxygenType),

		CodePassengerServiceType_stagedOrder: make(map[*CodePassengerServiceType]uint),
		CodePassengerServiceType_orderStaged: make(map[uint]*CodePassengerServiceType),
		CodePassengerServiceTypes_reference:  make(map[*CodePassengerServiceType]*CodePassengerServiceType),

		CodePilotControlledLightingType_stagedOrder: make(map[*CodePilotControlledLightingType]uint),
		CodePilotControlledLightingType_orderStaged: make(map[uint]*CodePilotControlledLightingType),
		CodePilotControlledLightingTypes_reference:  make(map[*CodePilotControlledLightingType]*CodePilotControlledLightingType),

		CodeRVSMType_stagedOrder: make(map[*CodeRVSMType]uint),
		CodeRVSMType_orderStaged: make(map[uint]*CodeRVSMType),
		CodeRVSMTypes_reference:  make(map[*CodeRVSMType]*CodeRVSMType),

		CodeRadioEmissionType_stagedOrder: make(map[*CodeRadioEmissionType]uint),
		CodeRadioEmissionType_orderStaged: make(map[uint]*CodeRadioEmissionType),
		CodeRadioEmissionTypes_reference:  make(map[*CodeRadioEmissionType]*CodeRadioEmissionType),

		CodeServiceATCType_stagedOrder: make(map[*CodeServiceATCType]uint),
		CodeServiceATCType_orderStaged: make(map[uint]*CodeServiceATCType),
		CodeServiceATCTypes_reference:  make(map[*CodeServiceATCType]*CodeServiceATCType),

		CodeServiceATFMType_stagedOrder: make(map[*CodeServiceATFMType]uint),
		CodeServiceATFMType_orderStaged: make(map[uint]*CodeServiceATFMType),
		CodeServiceATFMTypes_reference:  make(map[*CodeServiceATFMType]*CodeServiceATFMType),

		CodeServiceGroundControlType_stagedOrder: make(map[*CodeServiceGroundControlType]uint),
		CodeServiceGroundControlType_orderStaged: make(map[uint]*CodeServiceGroundControlType),
		CodeServiceGroundControlTypes_reference:  make(map[*CodeServiceGroundControlType]*CodeServiceGroundControlType),

		CodeServiceInformationType_stagedOrder: make(map[*CodeServiceInformationType]uint),
		CodeServiceInformationType_orderStaged: make(map[uint]*CodeServiceInformationType),
		CodeServiceInformationTypes_reference:  make(map[*CodeServiceInformationType]*CodeServiceInformationType),

		CodeServiceSARType_stagedOrder: make(map[*CodeServiceSARType]uint),
		CodeServiceSARType_orderStaged: make(map[uint]*CodeServiceSARType),
		CodeServiceSARTypes_reference:  make(map[*CodeServiceSARType]*CodeServiceSARType),

		CodeSpecialDateType_stagedOrder: make(map[*CodeSpecialDateType]uint),
		CodeSpecialDateType_orderStaged: make(map[uint]*CodeSpecialDateType),
		CodeSpecialDateTypes_reference:  make(map[*CodeSpecialDateType]*CodeSpecialDateType),

		CodeStatusOperationsType_stagedOrder: make(map[*CodeStatusOperationsType]uint),
		CodeStatusOperationsType_orderStaged: make(map[uint]*CodeStatusOperationsType),
		CodeStatusOperationsTypes_reference:  make(map[*CodeStatusOperationsType]*CodeStatusOperationsType),

		CodeStatusServiceType_stagedOrder: make(map[*CodeStatusServiceType]uint),
		CodeStatusServiceType_orderStaged: make(map[uint]*CodeStatusServiceType),
		CodeStatusServiceTypes_reference:  make(map[*CodeStatusServiceType]*CodeStatusServiceType),

		CodeSystemActivationType_stagedOrder: make(map[*CodeSystemActivationType]uint),
		CodeSystemActivationType_orderStaged: make(map[uint]*CodeSystemActivationType),
		CodeSystemActivationTypes_reference:  make(map[*CodeSystemActivationType]*CodeSystemActivationType),

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

		CodeUnitDependencyType_stagedOrder: make(map[*CodeUnitDependencyType]uint),
		CodeUnitDependencyType_orderStaged: make(map[uint]*CodeUnitDependencyType),
		CodeUnitDependencyTypes_reference:  make(map[*CodeUnitDependencyType]*CodeUnitDependencyType),

		CodeUnitType_stagedOrder: make(map[*CodeUnitType]uint),
		CodeUnitType_orderStaged: make(map[uint]*CodeUnitType),
		CodeUnitTypes_reference:  make(map[*CodeUnitType]*CodeUnitType),

		CodeVerticalDatumType_stagedOrder: make(map[*CodeVerticalDatumType]uint),
		CodeVerticalDatumType_orderStaged: make(map[uint]*CodeVerticalDatumType),
		CodeVerticalDatumTypes_reference:  make(map[*CodeVerticalDatumType]*CodeVerticalDatumType),

		CodeYesNoType_stagedOrder: make(map[*CodeYesNoType]uint),
		CodeYesNoType_orderStaged: make(map[uint]*CodeYesNoType),
		CodeYesNoTypes_reference:  make(map[*CodeYesNoType]*CodeYesNoType),

		DateMonthDayType_stagedOrder: make(map[*DateMonthDayType]uint),
		DateMonthDayType_orderStaged: make(map[uint]*DateMonthDayType),
		DateMonthDayTypes_reference:  make(map[*DateMonthDayType]*DateMonthDayType),

		DateYearType_stagedOrder: make(map[*DateYearType]uint),
		DateYearType_orderStaged: make(map[uint]*DateYearType),
		DateYearTypes_reference:  make(map[*DateYearType]*DateYearType),

		NoNumberType_stagedOrder: make(map[*NoNumberType]uint),
		NoNumberType_orderStaged: make(map[uint]*NoNumberType),
		NoNumberTypes_reference:  make(map[*NoNumberType]*NoNumberType),

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

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeAircraftGroundServiceType": &CodeAircraftGroundServiceTypeUnmarshaller{},

			"CodeAviationStandardsType": &CodeAviationStandardsTypeUnmarshaller{},

			"CodeCommunicationChannelType": &CodeCommunicationChannelTypeUnmarshaller{},

			"CodeCommunicationDirectionType": &CodeCommunicationDirectionTypeUnmarshaller{},

			"CodeCommunicationModeType": &CodeCommunicationModeTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeFacilityRankingType": &CodeFacilityRankingTypeUnmarshaller{},

			"CodeFireFightingType": &CodeFireFightingTypeUnmarshaller{},

			"CodeFlightDestinationType": &CodeFlightDestinationTypeUnmarshaller{},

			"CodeFlightRuleType": &CodeFlightRuleTypeUnmarshaller{},

			"CodeFuelType": &CodeFuelTypeUnmarshaller{},

			"CodeIntensityStandByType": &CodeIntensityStandByTypeUnmarshaller{},

			"CodeLanguageType": &CodeLanguageTypeUnmarshaller{},

			"CodeLevelSeriesType": &CodeLevelSeriesTypeUnmarshaller{},

			"CodeLevelTableDesignatorType": &CodeLevelTableDesignatorTypeUnmarshaller{},

			"CodeLightIntensityType": &CodeLightIntensityTypeUnmarshaller{},

			"CodeMilitaryOperationsType": &CodeMilitaryOperationsTypeUnmarshaller{},

			"CodeNitrogenType": &CodeNitrogenTypeUnmarshaller{},

			"CodeNorthReferenceType": &CodeNorthReferenceTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeOilType": &CodeOilTypeUnmarshaller{},

			"CodeOrganisationDesignatorType": &CodeOrganisationDesignatorTypeUnmarshaller{},

			"CodeOrganisationHierarchyType": &CodeOrganisationHierarchyTypeUnmarshaller{},

			"CodeOrganisationType": &CodeOrganisationTypeUnmarshaller{},

			"CodeOxygenType": &CodeOxygenTypeUnmarshaller{},

			"CodePassengerServiceType": &CodePassengerServiceTypeUnmarshaller{},

			"CodePilotControlledLightingType": &CodePilotControlledLightingTypeUnmarshaller{},

			"CodeRVSMType": &CodeRVSMTypeUnmarshaller{},

			"CodeRadioEmissionType": &CodeRadioEmissionTypeUnmarshaller{},

			"CodeServiceATCType": &CodeServiceATCTypeUnmarshaller{},

			"CodeServiceATFMType": &CodeServiceATFMTypeUnmarshaller{},

			"CodeServiceGroundControlType": &CodeServiceGroundControlTypeUnmarshaller{},

			"CodeServiceInformationType": &CodeServiceInformationTypeUnmarshaller{},

			"CodeServiceSARType": &CodeServiceSARTypeUnmarshaller{},

			"CodeSpecialDateType": &CodeSpecialDateTypeUnmarshaller{},

			"CodeStatusOperationsType": &CodeStatusOperationsTypeUnmarshaller{},

			"CodeStatusServiceType": &CodeStatusServiceTypeUnmarshaller{},

			"CodeSystemActivationType": &CodeSystemActivationTypeUnmarshaller{},

			"CodeTelecomNetworkType": &CodeTelecomNetworkTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeUnitDependencyType": &CodeUnitDependencyTypeUnmarshaller{},

			"CodeUnitType": &CodeUnitTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"DateYearType": &DateYearTypeUnmarshaller{},

			"NoNumberType": &NoNumberTypeUnmarshaller{},

			"TextAddressType": &TextAddressTypeUnmarshaller{},

			"TextDesignatorType": &TextDesignatorTypeUnmarshaller{},

			"TextInstructionType": &TextInstructionTypeUnmarshaller{},

			"TextNameType": &TextNameTypeUnmarshaller{},

			"TextNoteType": &TextNoteTypeUnmarshaller{},

			"TextPhoneType": &TextPhoneTypeUnmarshaller{},

			"TextPropertyNameType": &TextPropertyNameTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"ValBearingType": &ValBearingTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValFrequencyType": &ValFrequencyTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeAircraftGroundServiceType"] = len(stage.CodeAircraftGroundServiceTypes)
	stage.Map_GongStructName_InstancesNb["CodeAviationStandardsType"] = len(stage.CodeAviationStandardsTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationChannelType"] = len(stage.CodeCommunicationChannelTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationDirectionType"] = len(stage.CodeCommunicationDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationModeType"] = len(stage.CodeCommunicationModeTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeFacilityRankingType"] = len(stage.CodeFacilityRankingTypes)
	stage.Map_GongStructName_InstancesNb["CodeFireFightingType"] = len(stage.CodeFireFightingTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightDestinationType"] = len(stage.CodeFlightDestinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightRuleType"] = len(stage.CodeFlightRuleTypes)
	stage.Map_GongStructName_InstancesNb["CodeFuelType"] = len(stage.CodeFuelTypes)
	stage.Map_GongStructName_InstancesNb["CodeIntensityStandByType"] = len(stage.CodeIntensityStandByTypes)
	stage.Map_GongStructName_InstancesNb["CodeLanguageType"] = len(stage.CodeLanguageTypes)
	stage.Map_GongStructName_InstancesNb["CodeLevelSeriesType"] = len(stage.CodeLevelSeriesTypes)
	stage.Map_GongStructName_InstancesNb["CodeLevelTableDesignatorType"] = len(stage.CodeLevelTableDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightIntensityType"] = len(stage.CodeLightIntensityTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryOperationsType"] = len(stage.CodeMilitaryOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeNitrogenType"] = len(stage.CodeNitrogenTypes)
	stage.Map_GongStructName_InstancesNb["CodeNorthReferenceType"] = len(stage.CodeNorthReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeOilType"] = len(stage.CodeOilTypes)
	stage.Map_GongStructName_InstancesNb["CodeOrganisationDesignatorType"] = len(stage.CodeOrganisationDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeOrganisationHierarchyType"] = len(stage.CodeOrganisationHierarchyTypes)
	stage.Map_GongStructName_InstancesNb["CodeOrganisationType"] = len(stage.CodeOrganisationTypes)
	stage.Map_GongStructName_InstancesNb["CodeOxygenType"] = len(stage.CodeOxygenTypes)
	stage.Map_GongStructName_InstancesNb["CodePassengerServiceType"] = len(stage.CodePassengerServiceTypes)
	stage.Map_GongStructName_InstancesNb["CodePilotControlledLightingType"] = len(stage.CodePilotControlledLightingTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMType"] = len(stage.CodeRVSMTypes)
	stage.Map_GongStructName_InstancesNb["CodeRadioEmissionType"] = len(stage.CodeRadioEmissionTypes)
	stage.Map_GongStructName_InstancesNb["CodeServiceATCType"] = len(stage.CodeServiceATCTypes)
	stage.Map_GongStructName_InstancesNb["CodeServiceATFMType"] = len(stage.CodeServiceATFMTypes)
	stage.Map_GongStructName_InstancesNb["CodeServiceGroundControlType"] = len(stage.CodeServiceGroundControlTypes)
	stage.Map_GongStructName_InstancesNb["CodeServiceInformationType"] = len(stage.CodeServiceInformationTypes)
	stage.Map_GongStructName_InstancesNb["CodeServiceSARType"] = len(stage.CodeServiceSARTypes)
	stage.Map_GongStructName_InstancesNb["CodeSpecialDateType"] = len(stage.CodeSpecialDateTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusOperationsType"] = len(stage.CodeStatusOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusServiceType"] = len(stage.CodeStatusServiceTypes)
	stage.Map_GongStructName_InstancesNb["CodeSystemActivationType"] = len(stage.CodeSystemActivationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTelecomNetworkType"] = len(stage.CodeTelecomNetworkTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeUnitDependencyType"] = len(stage.CodeUnitDependencyTypes)
	stage.Map_GongStructName_InstancesNb["CodeUnitType"] = len(stage.CodeUnitTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["DateYearType"] = len(stage.DateYearTypes)
	stage.Map_GongStructName_InstancesNb["NoNumberType"] = len(stage.NoNumberTypes)
	stage.Map_GongStructName_InstancesNb["TextAddressType"] = len(stage.TextAddressTypes)
	stage.Map_GongStructName_InstancesNb["TextDesignatorType"] = len(stage.TextDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["TextInstructionType"] = len(stage.TextInstructionTypes)
	stage.Map_GongStructName_InstancesNb["TextNameType"] = len(stage.TextNameTypes)
	stage.Map_GongStructName_InstancesNb["TextNoteType"] = len(stage.TextNoteTypes)
	stage.Map_GongStructName_InstancesNb["TextPhoneType"] = len(stage.TextPhoneTypes)
	stage.Map_GongStructName_InstancesNb["TextPropertyNameType"] = len(stage.TextPropertyNameTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["ValBearingType"] = len(stage.ValBearingTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValFrequencyType"] = len(stage.ValFrequencyTypes)
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
// Stage puts codeaircraftgroundservicetype to the model stage
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) Stage(stage *Stage) *CodeAircraftGroundServiceType {
	__gong__stage(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceType_stagedOrder, stage.CodeAircraftGroundServiceType_orderStaged, &stage.CodeAircraftGroundServiceTypeOrder, stage.CodeAircraftGroundServiceTypes_mapString, codeaircraftgroundservicetype, codeaircraftgroundservicetype.Name)
	return codeaircraftgroundservicetype
}

// StagePreserveOrder puts codeaircraftgroundservicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftGroundServiceTypeOrder
// - update stage.CodeAircraftGroundServiceTypeOrder accordingly
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceType_stagedOrder, stage.CodeAircraftGroundServiceType_orderStaged, &stage.CodeAircraftGroundServiceTypeOrder, stage.CodeAircraftGroundServiceTypes_mapString, codeaircraftgroundservicetype, order, codeaircraftgroundservicetype.Name)
}

// Unstage removes codeaircraftgroundservicetype off the model stage
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) Unstage(stage *Stage) *CodeAircraftGroundServiceType {
	__gong__unstage(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceTypes_mapString, codeaircraftgroundservicetype, codeaircraftgroundservicetype.Name)
	return codeaircraftgroundservicetype
}

// UnstageVoid removes codeaircraftgroundservicetype off the model stage
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) UnstageVoid(stage *Stage) {
	codeaircraftgroundservicetype.Unstage(stage)
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) StageVoid(stage *Stage) {
	codeaircraftgroundservicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GetName() (res string) {
	return codeaircraftgroundservicetype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) SetName(name string) {
	codeaircraftgroundservicetype.Name = name
}

// Stage puts codeaviationstandardstype to the model stage
func (codeaviationstandardstype *CodeAviationStandardsType) Stage(stage *Stage) *CodeAviationStandardsType {
	__gong__stage(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsType_stagedOrder, stage.CodeAviationStandardsType_orderStaged, &stage.CodeAviationStandardsTypeOrder, stage.CodeAviationStandardsTypes_mapString, codeaviationstandardstype, codeaviationstandardstype.Name)
	return codeaviationstandardstype
}

// StagePreserveOrder puts codeaviationstandardstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAviationStandardsTypeOrder
// - update stage.CodeAviationStandardsTypeOrder accordingly
func (codeaviationstandardstype *CodeAviationStandardsType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsType_stagedOrder, stage.CodeAviationStandardsType_orderStaged, &stage.CodeAviationStandardsTypeOrder, stage.CodeAviationStandardsTypes_mapString, codeaviationstandardstype, order, codeaviationstandardstype.Name)
}

// Unstage removes codeaviationstandardstype off the model stage
func (codeaviationstandardstype *CodeAviationStandardsType) Unstage(stage *Stage) *CodeAviationStandardsType {
	__gong__unstage(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsTypes_mapString, codeaviationstandardstype, codeaviationstandardstype.Name)
	return codeaviationstandardstype
}

// UnstageVoid removes codeaviationstandardstype off the model stage
func (codeaviationstandardstype *CodeAviationStandardsType) UnstageVoid(stage *Stage) {
	codeaviationstandardstype.Unstage(stage)
}

func (codeaviationstandardstype *CodeAviationStandardsType) StageVoid(stage *Stage) {
	codeaviationstandardstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaviationstandardstype *CodeAviationStandardsType) GetName() (res string) {
	return codeaviationstandardstype.Name
}

// for satisfaction of GongStruct interface
func (codeaviationstandardstype *CodeAviationStandardsType) SetName(name string) {
	codeaviationstandardstype.Name = name
}

// Stage puts codecommunicationchanneltype to the model stage
func (codecommunicationchanneltype *CodeCommunicationChannelType) Stage(stage *Stage) *CodeCommunicationChannelType {
	__gong__stage(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelType_stagedOrder, stage.CodeCommunicationChannelType_orderStaged, &stage.CodeCommunicationChannelTypeOrder, stage.CodeCommunicationChannelTypes_mapString, codecommunicationchanneltype, codecommunicationchanneltype.Name)
	return codecommunicationchanneltype
}

// StagePreserveOrder puts codecommunicationchanneltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCommunicationChannelTypeOrder
// - update stage.CodeCommunicationChannelTypeOrder accordingly
func (codecommunicationchanneltype *CodeCommunicationChannelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelType_stagedOrder, stage.CodeCommunicationChannelType_orderStaged, &stage.CodeCommunicationChannelTypeOrder, stage.CodeCommunicationChannelTypes_mapString, codecommunicationchanneltype, order, codecommunicationchanneltype.Name)
}

// Unstage removes codecommunicationchanneltype off the model stage
func (codecommunicationchanneltype *CodeCommunicationChannelType) Unstage(stage *Stage) *CodeCommunicationChannelType {
	__gong__unstage(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelTypes_mapString, codecommunicationchanneltype, codecommunicationchanneltype.Name)
	return codecommunicationchanneltype
}

// UnstageVoid removes codecommunicationchanneltype off the model stage
func (codecommunicationchanneltype *CodeCommunicationChannelType) UnstageVoid(stage *Stage) {
	codecommunicationchanneltype.Unstage(stage)
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) StageVoid(stage *Stage) {
	codecommunicationchanneltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecommunicationchanneltype *CodeCommunicationChannelType) GetName() (res string) {
	return codecommunicationchanneltype.Name
}

// for satisfaction of GongStruct interface
func (codecommunicationchanneltype *CodeCommunicationChannelType) SetName(name string) {
	codecommunicationchanneltype.Name = name
}

// Stage puts codecommunicationdirectiontype to the model stage
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) Stage(stage *Stage) *CodeCommunicationDirectionType {
	__gong__stage(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionType_stagedOrder, stage.CodeCommunicationDirectionType_orderStaged, &stage.CodeCommunicationDirectionTypeOrder, stage.CodeCommunicationDirectionTypes_mapString, codecommunicationdirectiontype, codecommunicationdirectiontype.Name)
	return codecommunicationdirectiontype
}

// StagePreserveOrder puts codecommunicationdirectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeCommunicationDirectionTypeOrder
// - update stage.CodeCommunicationDirectionTypeOrder accordingly
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionType_stagedOrder, stage.CodeCommunicationDirectionType_orderStaged, &stage.CodeCommunicationDirectionTypeOrder, stage.CodeCommunicationDirectionTypes_mapString, codecommunicationdirectiontype, order, codecommunicationdirectiontype.Name)
}

// Unstage removes codecommunicationdirectiontype off the model stage
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) Unstage(stage *Stage) *CodeCommunicationDirectionType {
	__gong__unstage(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionTypes_mapString, codecommunicationdirectiontype, codecommunicationdirectiontype.Name)
	return codecommunicationdirectiontype
}

// UnstageVoid removes codecommunicationdirectiontype off the model stage
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) UnstageVoid(stage *Stage) {
	codecommunicationdirectiontype.Unstage(stage)
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) StageVoid(stage *Stage) {
	codecommunicationdirectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GetName() (res string) {
	return codecommunicationdirectiontype.Name
}

// for satisfaction of GongStruct interface
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) SetName(name string) {
	codecommunicationdirectiontype.Name = name
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

// Stage puts codefacilityrankingtype to the model stage
func (codefacilityrankingtype *CodeFacilityRankingType) Stage(stage *Stage) *CodeFacilityRankingType {
	__gong__stage(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingType_stagedOrder, stage.CodeFacilityRankingType_orderStaged, &stage.CodeFacilityRankingTypeOrder, stage.CodeFacilityRankingTypes_mapString, codefacilityrankingtype, codefacilityrankingtype.Name)
	return codefacilityrankingtype
}

// StagePreserveOrder puts codefacilityrankingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFacilityRankingTypeOrder
// - update stage.CodeFacilityRankingTypeOrder accordingly
func (codefacilityrankingtype *CodeFacilityRankingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingType_stagedOrder, stage.CodeFacilityRankingType_orderStaged, &stage.CodeFacilityRankingTypeOrder, stage.CodeFacilityRankingTypes_mapString, codefacilityrankingtype, order, codefacilityrankingtype.Name)
}

// Unstage removes codefacilityrankingtype off the model stage
func (codefacilityrankingtype *CodeFacilityRankingType) Unstage(stage *Stage) *CodeFacilityRankingType {
	__gong__unstage(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingTypes_mapString, codefacilityrankingtype, codefacilityrankingtype.Name)
	return codefacilityrankingtype
}

// UnstageVoid removes codefacilityrankingtype off the model stage
func (codefacilityrankingtype *CodeFacilityRankingType) UnstageVoid(stage *Stage) {
	codefacilityrankingtype.Unstage(stage)
}

func (codefacilityrankingtype *CodeFacilityRankingType) StageVoid(stage *Stage) {
	codefacilityrankingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefacilityrankingtype *CodeFacilityRankingType) GetName() (res string) {
	return codefacilityrankingtype.Name
}

// for satisfaction of GongStruct interface
func (codefacilityrankingtype *CodeFacilityRankingType) SetName(name string) {
	codefacilityrankingtype.Name = name
}

// Stage puts codefirefightingtype to the model stage
func (codefirefightingtype *CodeFireFightingType) Stage(stage *Stage) *CodeFireFightingType {
	__gong__stage(stage.CodeFireFightingTypes, stage.CodeFireFightingType_stagedOrder, stage.CodeFireFightingType_orderStaged, &stage.CodeFireFightingTypeOrder, stage.CodeFireFightingTypes_mapString, codefirefightingtype, codefirefightingtype.Name)
	return codefirefightingtype
}

// StagePreserveOrder puts codefirefightingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFireFightingTypeOrder
// - update stage.CodeFireFightingTypeOrder accordingly
func (codefirefightingtype *CodeFireFightingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFireFightingTypes, stage.CodeFireFightingType_stagedOrder, stage.CodeFireFightingType_orderStaged, &stage.CodeFireFightingTypeOrder, stage.CodeFireFightingTypes_mapString, codefirefightingtype, order, codefirefightingtype.Name)
}

// Unstage removes codefirefightingtype off the model stage
func (codefirefightingtype *CodeFireFightingType) Unstage(stage *Stage) *CodeFireFightingType {
	__gong__unstage(stage.CodeFireFightingTypes, stage.CodeFireFightingTypes_mapString, codefirefightingtype, codefirefightingtype.Name)
	return codefirefightingtype
}

// UnstageVoid removes codefirefightingtype off the model stage
func (codefirefightingtype *CodeFireFightingType) UnstageVoid(stage *Stage) {
	codefirefightingtype.Unstage(stage)
}

func (codefirefightingtype *CodeFireFightingType) StageVoid(stage *Stage) {
	codefirefightingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefirefightingtype *CodeFireFightingType) GetName() (res string) {
	return codefirefightingtype.Name
}

// for satisfaction of GongStruct interface
func (codefirefightingtype *CodeFireFightingType) SetName(name string) {
	codefirefightingtype.Name = name
}

// Stage puts codeflightdestinationtype to the model stage
func (codeflightdestinationtype *CodeFlightDestinationType) Stage(stage *Stage) *CodeFlightDestinationType {
	__gong__stage(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationType_stagedOrder, stage.CodeFlightDestinationType_orderStaged, &stage.CodeFlightDestinationTypeOrder, stage.CodeFlightDestinationTypes_mapString, codeflightdestinationtype, codeflightdestinationtype.Name)
	return codeflightdestinationtype
}

// StagePreserveOrder puts codeflightdestinationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFlightDestinationTypeOrder
// - update stage.CodeFlightDestinationTypeOrder accordingly
func (codeflightdestinationtype *CodeFlightDestinationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationType_stagedOrder, stage.CodeFlightDestinationType_orderStaged, &stage.CodeFlightDestinationTypeOrder, stage.CodeFlightDestinationTypes_mapString, codeflightdestinationtype, order, codeflightdestinationtype.Name)
}

// Unstage removes codeflightdestinationtype off the model stage
func (codeflightdestinationtype *CodeFlightDestinationType) Unstage(stage *Stage) *CodeFlightDestinationType {
	__gong__unstage(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationTypes_mapString, codeflightdestinationtype, codeflightdestinationtype.Name)
	return codeflightdestinationtype
}

// UnstageVoid removes codeflightdestinationtype off the model stage
func (codeflightdestinationtype *CodeFlightDestinationType) UnstageVoid(stage *Stage) {
	codeflightdestinationtype.Unstage(stage)
}

func (codeflightdestinationtype *CodeFlightDestinationType) StageVoid(stage *Stage) {
	codeflightdestinationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeflightdestinationtype *CodeFlightDestinationType) GetName() (res string) {
	return codeflightdestinationtype.Name
}

// for satisfaction of GongStruct interface
func (codeflightdestinationtype *CodeFlightDestinationType) SetName(name string) {
	codeflightdestinationtype.Name = name
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

// Stage puts codefueltype to the model stage
func (codefueltype *CodeFuelType) Stage(stage *Stage) *CodeFuelType {
	__gong__stage(stage.CodeFuelTypes, stage.CodeFuelType_stagedOrder, stage.CodeFuelType_orderStaged, &stage.CodeFuelTypeOrder, stage.CodeFuelTypes_mapString, codefueltype, codefueltype.Name)
	return codefueltype
}

// StagePreserveOrder puts codefueltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFuelTypeOrder
// - update stage.CodeFuelTypeOrder accordingly
func (codefueltype *CodeFuelType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFuelTypes, stage.CodeFuelType_stagedOrder, stage.CodeFuelType_orderStaged, &stage.CodeFuelTypeOrder, stage.CodeFuelTypes_mapString, codefueltype, order, codefueltype.Name)
}

// Unstage removes codefueltype off the model stage
func (codefueltype *CodeFuelType) Unstage(stage *Stage) *CodeFuelType {
	__gong__unstage(stage.CodeFuelTypes, stage.CodeFuelTypes_mapString, codefueltype, codefueltype.Name)
	return codefueltype
}

// UnstageVoid removes codefueltype off the model stage
func (codefueltype *CodeFuelType) UnstageVoid(stage *Stage) {
	codefueltype.Unstage(stage)
}

func (codefueltype *CodeFuelType) StageVoid(stage *Stage) {
	codefueltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefueltype *CodeFuelType) GetName() (res string) {
	return codefueltype.Name
}

// for satisfaction of GongStruct interface
func (codefueltype *CodeFuelType) SetName(name string) {
	codefueltype.Name = name
}

// Stage puts codeintensitystandbytype to the model stage
func (codeintensitystandbytype *CodeIntensityStandByType) Stage(stage *Stage) *CodeIntensityStandByType {
	__gong__stage(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByType_stagedOrder, stage.CodeIntensityStandByType_orderStaged, &stage.CodeIntensityStandByTypeOrder, stage.CodeIntensityStandByTypes_mapString, codeintensitystandbytype, codeintensitystandbytype.Name)
	return codeintensitystandbytype
}

// StagePreserveOrder puts codeintensitystandbytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeIntensityStandByTypeOrder
// - update stage.CodeIntensityStandByTypeOrder accordingly
func (codeintensitystandbytype *CodeIntensityStandByType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByType_stagedOrder, stage.CodeIntensityStandByType_orderStaged, &stage.CodeIntensityStandByTypeOrder, stage.CodeIntensityStandByTypes_mapString, codeintensitystandbytype, order, codeintensitystandbytype.Name)
}

// Unstage removes codeintensitystandbytype off the model stage
func (codeintensitystandbytype *CodeIntensityStandByType) Unstage(stage *Stage) *CodeIntensityStandByType {
	__gong__unstage(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByTypes_mapString, codeintensitystandbytype, codeintensitystandbytype.Name)
	return codeintensitystandbytype
}

// UnstageVoid removes codeintensitystandbytype off the model stage
func (codeintensitystandbytype *CodeIntensityStandByType) UnstageVoid(stage *Stage) {
	codeintensitystandbytype.Unstage(stage)
}

func (codeintensitystandbytype *CodeIntensityStandByType) StageVoid(stage *Stage) {
	codeintensitystandbytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeintensitystandbytype *CodeIntensityStandByType) GetName() (res string) {
	return codeintensitystandbytype.Name
}

// for satisfaction of GongStruct interface
func (codeintensitystandbytype *CodeIntensityStandByType) SetName(name string) {
	codeintensitystandbytype.Name = name
}

// Stage puts codelanguagetype to the model stage
func (codelanguagetype *CodeLanguageType) Stage(stage *Stage) *CodeLanguageType {
	__gong__stage(stage.CodeLanguageTypes, stage.CodeLanguageType_stagedOrder, stage.CodeLanguageType_orderStaged, &stage.CodeLanguageTypeOrder, stage.CodeLanguageTypes_mapString, codelanguagetype, codelanguagetype.Name)
	return codelanguagetype
}

// StagePreserveOrder puts codelanguagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLanguageTypeOrder
// - update stage.CodeLanguageTypeOrder accordingly
func (codelanguagetype *CodeLanguageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLanguageTypes, stage.CodeLanguageType_stagedOrder, stage.CodeLanguageType_orderStaged, &stage.CodeLanguageTypeOrder, stage.CodeLanguageTypes_mapString, codelanguagetype, order, codelanguagetype.Name)
}

// Unstage removes codelanguagetype off the model stage
func (codelanguagetype *CodeLanguageType) Unstage(stage *Stage) *CodeLanguageType {
	__gong__unstage(stage.CodeLanguageTypes, stage.CodeLanguageTypes_mapString, codelanguagetype, codelanguagetype.Name)
	return codelanguagetype
}

// UnstageVoid removes codelanguagetype off the model stage
func (codelanguagetype *CodeLanguageType) UnstageVoid(stage *Stage) {
	codelanguagetype.Unstage(stage)
}

func (codelanguagetype *CodeLanguageType) StageVoid(stage *Stage) {
	codelanguagetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelanguagetype *CodeLanguageType) GetName() (res string) {
	return codelanguagetype.Name
}

// for satisfaction of GongStruct interface
func (codelanguagetype *CodeLanguageType) SetName(name string) {
	codelanguagetype.Name = name
}

// Stage puts codelevelseriestype to the model stage
func (codelevelseriestype *CodeLevelSeriesType) Stage(stage *Stage) *CodeLevelSeriesType {
	__gong__stage(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesType_stagedOrder, stage.CodeLevelSeriesType_orderStaged, &stage.CodeLevelSeriesTypeOrder, stage.CodeLevelSeriesTypes_mapString, codelevelseriestype, codelevelseriestype.Name)
	return codelevelseriestype
}

// StagePreserveOrder puts codelevelseriestype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLevelSeriesTypeOrder
// - update stage.CodeLevelSeriesTypeOrder accordingly
func (codelevelseriestype *CodeLevelSeriesType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesType_stagedOrder, stage.CodeLevelSeriesType_orderStaged, &stage.CodeLevelSeriesTypeOrder, stage.CodeLevelSeriesTypes_mapString, codelevelseriestype, order, codelevelseriestype.Name)
}

// Unstage removes codelevelseriestype off the model stage
func (codelevelseriestype *CodeLevelSeriesType) Unstage(stage *Stage) *CodeLevelSeriesType {
	__gong__unstage(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesTypes_mapString, codelevelseriestype, codelevelseriestype.Name)
	return codelevelseriestype
}

// UnstageVoid removes codelevelseriestype off the model stage
func (codelevelseriestype *CodeLevelSeriesType) UnstageVoid(stage *Stage) {
	codelevelseriestype.Unstage(stage)
}

func (codelevelseriestype *CodeLevelSeriesType) StageVoid(stage *Stage) {
	codelevelseriestype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelevelseriestype *CodeLevelSeriesType) GetName() (res string) {
	return codelevelseriestype.Name
}

// for satisfaction of GongStruct interface
func (codelevelseriestype *CodeLevelSeriesType) SetName(name string) {
	codelevelseriestype.Name = name
}

// Stage puts codeleveltabledesignatortype to the model stage
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) Stage(stage *Stage) *CodeLevelTableDesignatorType {
	__gong__stage(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorType_stagedOrder, stage.CodeLevelTableDesignatorType_orderStaged, &stage.CodeLevelTableDesignatorTypeOrder, stage.CodeLevelTableDesignatorTypes_mapString, codeleveltabledesignatortype, codeleveltabledesignatortype.Name)
	return codeleveltabledesignatortype
}

// StagePreserveOrder puts codeleveltabledesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLevelTableDesignatorTypeOrder
// - update stage.CodeLevelTableDesignatorTypeOrder accordingly
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorType_stagedOrder, stage.CodeLevelTableDesignatorType_orderStaged, &stage.CodeLevelTableDesignatorTypeOrder, stage.CodeLevelTableDesignatorTypes_mapString, codeleveltabledesignatortype, order, codeleveltabledesignatortype.Name)
}

// Unstage removes codeleveltabledesignatortype off the model stage
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) Unstage(stage *Stage) *CodeLevelTableDesignatorType {
	__gong__unstage(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorTypes_mapString, codeleveltabledesignatortype, codeleveltabledesignatortype.Name)
	return codeleveltabledesignatortype
}

// UnstageVoid removes codeleveltabledesignatortype off the model stage
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) UnstageVoid(stage *Stage) {
	codeleveltabledesignatortype.Unstage(stage)
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) StageVoid(stage *Stage) {
	codeleveltabledesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GetName() (res string) {
	return codeleveltabledesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) SetName(name string) {
	codeleveltabledesignatortype.Name = name
}

// Stage puts codelightintensitytype to the model stage
func (codelightintensitytype *CodeLightIntensityType) Stage(stage *Stage) *CodeLightIntensityType {
	__gong__stage(stage.CodeLightIntensityTypes, stage.CodeLightIntensityType_stagedOrder, stage.CodeLightIntensityType_orderStaged, &stage.CodeLightIntensityTypeOrder, stage.CodeLightIntensityTypes_mapString, codelightintensitytype, codelightintensitytype.Name)
	return codelightintensitytype
}

// StagePreserveOrder puts codelightintensitytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLightIntensityTypeOrder
// - update stage.CodeLightIntensityTypeOrder accordingly
func (codelightintensitytype *CodeLightIntensityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLightIntensityTypes, stage.CodeLightIntensityType_stagedOrder, stage.CodeLightIntensityType_orderStaged, &stage.CodeLightIntensityTypeOrder, stage.CodeLightIntensityTypes_mapString, codelightintensitytype, order, codelightintensitytype.Name)
}

// Unstage removes codelightintensitytype off the model stage
func (codelightintensitytype *CodeLightIntensityType) Unstage(stage *Stage) *CodeLightIntensityType {
	__gong__unstage(stage.CodeLightIntensityTypes, stage.CodeLightIntensityTypes_mapString, codelightintensitytype, codelightintensitytype.Name)
	return codelightintensitytype
}

// UnstageVoid removes codelightintensitytype off the model stage
func (codelightintensitytype *CodeLightIntensityType) UnstageVoid(stage *Stage) {
	codelightintensitytype.Unstage(stage)
}

func (codelightintensitytype *CodeLightIntensityType) StageVoid(stage *Stage) {
	codelightintensitytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelightintensitytype *CodeLightIntensityType) GetName() (res string) {
	return codelightintensitytype.Name
}

// for satisfaction of GongStruct interface
func (codelightintensitytype *CodeLightIntensityType) SetName(name string) {
	codelightintensitytype.Name = name
}

// Stage puts codemilitaryoperationstype to the model stage
func (codemilitaryoperationstype *CodeMilitaryOperationsType) Stage(stage *Stage) *CodeMilitaryOperationsType {
	__gong__stage(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsType_stagedOrder, stage.CodeMilitaryOperationsType_orderStaged, &stage.CodeMilitaryOperationsTypeOrder, stage.CodeMilitaryOperationsTypes_mapString, codemilitaryoperationstype, codemilitaryoperationstype.Name)
	return codemilitaryoperationstype
}

// StagePreserveOrder puts codemilitaryoperationstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMilitaryOperationsTypeOrder
// - update stage.CodeMilitaryOperationsTypeOrder accordingly
func (codemilitaryoperationstype *CodeMilitaryOperationsType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsType_stagedOrder, stage.CodeMilitaryOperationsType_orderStaged, &stage.CodeMilitaryOperationsTypeOrder, stage.CodeMilitaryOperationsTypes_mapString, codemilitaryoperationstype, order, codemilitaryoperationstype.Name)
}

// Unstage removes codemilitaryoperationstype off the model stage
func (codemilitaryoperationstype *CodeMilitaryOperationsType) Unstage(stage *Stage) *CodeMilitaryOperationsType {
	__gong__unstage(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsTypes_mapString, codemilitaryoperationstype, codemilitaryoperationstype.Name)
	return codemilitaryoperationstype
}

// UnstageVoid removes codemilitaryoperationstype off the model stage
func (codemilitaryoperationstype *CodeMilitaryOperationsType) UnstageVoid(stage *Stage) {
	codemilitaryoperationstype.Unstage(stage)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) StageVoid(stage *Stage) {
	codemilitaryoperationstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GetName() (res string) {
	return codemilitaryoperationstype.Name
}

// for satisfaction of GongStruct interface
func (codemilitaryoperationstype *CodeMilitaryOperationsType) SetName(name string) {
	codemilitaryoperationstype.Name = name
}

// Stage puts codenitrogentype to the model stage
func (codenitrogentype *CodeNitrogenType) Stage(stage *Stage) *CodeNitrogenType {
	__gong__stage(stage.CodeNitrogenTypes, stage.CodeNitrogenType_stagedOrder, stage.CodeNitrogenType_orderStaged, &stage.CodeNitrogenTypeOrder, stage.CodeNitrogenTypes_mapString, codenitrogentype, codenitrogentype.Name)
	return codenitrogentype
}

// StagePreserveOrder puts codenitrogentype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeNitrogenTypeOrder
// - update stage.CodeNitrogenTypeOrder accordingly
func (codenitrogentype *CodeNitrogenType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeNitrogenTypes, stage.CodeNitrogenType_stagedOrder, stage.CodeNitrogenType_orderStaged, &stage.CodeNitrogenTypeOrder, stage.CodeNitrogenTypes_mapString, codenitrogentype, order, codenitrogentype.Name)
}

// Unstage removes codenitrogentype off the model stage
func (codenitrogentype *CodeNitrogenType) Unstage(stage *Stage) *CodeNitrogenType {
	__gong__unstage(stage.CodeNitrogenTypes, stage.CodeNitrogenTypes_mapString, codenitrogentype, codenitrogentype.Name)
	return codenitrogentype
}

// UnstageVoid removes codenitrogentype off the model stage
func (codenitrogentype *CodeNitrogenType) UnstageVoid(stage *Stage) {
	codenitrogentype.Unstage(stage)
}

func (codenitrogentype *CodeNitrogenType) StageVoid(stage *Stage) {
	codenitrogentype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codenitrogentype *CodeNitrogenType) GetName() (res string) {
	return codenitrogentype.Name
}

// for satisfaction of GongStruct interface
func (codenitrogentype *CodeNitrogenType) SetName(name string) {
	codenitrogentype.Name = name
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

// Stage puts codeoiltype to the model stage
func (codeoiltype *CodeOilType) Stage(stage *Stage) *CodeOilType {
	__gong__stage(stage.CodeOilTypes, stage.CodeOilType_stagedOrder, stage.CodeOilType_orderStaged, &stage.CodeOilTypeOrder, stage.CodeOilTypes_mapString, codeoiltype, codeoiltype.Name)
	return codeoiltype
}

// StagePreserveOrder puts codeoiltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOilTypeOrder
// - update stage.CodeOilTypeOrder accordingly
func (codeoiltype *CodeOilType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOilTypes, stage.CodeOilType_stagedOrder, stage.CodeOilType_orderStaged, &stage.CodeOilTypeOrder, stage.CodeOilTypes_mapString, codeoiltype, order, codeoiltype.Name)
}

// Unstage removes codeoiltype off the model stage
func (codeoiltype *CodeOilType) Unstage(stage *Stage) *CodeOilType {
	__gong__unstage(stage.CodeOilTypes, stage.CodeOilTypes_mapString, codeoiltype, codeoiltype.Name)
	return codeoiltype
}

// UnstageVoid removes codeoiltype off the model stage
func (codeoiltype *CodeOilType) UnstageVoid(stage *Stage) {
	codeoiltype.Unstage(stage)
}

func (codeoiltype *CodeOilType) StageVoid(stage *Stage) {
	codeoiltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeoiltype *CodeOilType) GetName() (res string) {
	return codeoiltype.Name
}

// for satisfaction of GongStruct interface
func (codeoiltype *CodeOilType) SetName(name string) {
	codeoiltype.Name = name
}

// Stage puts codeorganisationdesignatortype to the model stage
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) Stage(stage *Stage) *CodeOrganisationDesignatorType {
	__gong__stage(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorType_stagedOrder, stage.CodeOrganisationDesignatorType_orderStaged, &stage.CodeOrganisationDesignatorTypeOrder, stage.CodeOrganisationDesignatorTypes_mapString, codeorganisationdesignatortype, codeorganisationdesignatortype.Name)
	return codeorganisationdesignatortype
}

// StagePreserveOrder puts codeorganisationdesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOrganisationDesignatorTypeOrder
// - update stage.CodeOrganisationDesignatorTypeOrder accordingly
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorType_stagedOrder, stage.CodeOrganisationDesignatorType_orderStaged, &stage.CodeOrganisationDesignatorTypeOrder, stage.CodeOrganisationDesignatorTypes_mapString, codeorganisationdesignatortype, order, codeorganisationdesignatortype.Name)
}

// Unstage removes codeorganisationdesignatortype off the model stage
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) Unstage(stage *Stage) *CodeOrganisationDesignatorType {
	__gong__unstage(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorTypes_mapString, codeorganisationdesignatortype, codeorganisationdesignatortype.Name)
	return codeorganisationdesignatortype
}

// UnstageVoid removes codeorganisationdesignatortype off the model stage
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) UnstageVoid(stage *Stage) {
	codeorganisationdesignatortype.Unstage(stage)
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) StageVoid(stage *Stage) {
	codeorganisationdesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GetName() (res string) {
	return codeorganisationdesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) SetName(name string) {
	codeorganisationdesignatortype.Name = name
}

// Stage puts codeorganisationhierarchytype to the model stage
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) Stage(stage *Stage) *CodeOrganisationHierarchyType {
	__gong__stage(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyType_stagedOrder, stage.CodeOrganisationHierarchyType_orderStaged, &stage.CodeOrganisationHierarchyTypeOrder, stage.CodeOrganisationHierarchyTypes_mapString, codeorganisationhierarchytype, codeorganisationhierarchytype.Name)
	return codeorganisationhierarchytype
}

// StagePreserveOrder puts codeorganisationhierarchytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOrganisationHierarchyTypeOrder
// - update stage.CodeOrganisationHierarchyTypeOrder accordingly
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyType_stagedOrder, stage.CodeOrganisationHierarchyType_orderStaged, &stage.CodeOrganisationHierarchyTypeOrder, stage.CodeOrganisationHierarchyTypes_mapString, codeorganisationhierarchytype, order, codeorganisationhierarchytype.Name)
}

// Unstage removes codeorganisationhierarchytype off the model stage
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) Unstage(stage *Stage) *CodeOrganisationHierarchyType {
	__gong__unstage(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyTypes_mapString, codeorganisationhierarchytype, codeorganisationhierarchytype.Name)
	return codeorganisationhierarchytype
}

// UnstageVoid removes codeorganisationhierarchytype off the model stage
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) UnstageVoid(stage *Stage) {
	codeorganisationhierarchytype.Unstage(stage)
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) StageVoid(stage *Stage) {
	codeorganisationhierarchytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GetName() (res string) {
	return codeorganisationhierarchytype.Name
}

// for satisfaction of GongStruct interface
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) SetName(name string) {
	codeorganisationhierarchytype.Name = name
}

// Stage puts codeorganisationtype to the model stage
func (codeorganisationtype *CodeOrganisationType) Stage(stage *Stage) *CodeOrganisationType {
	__gong__stage(stage.CodeOrganisationTypes, stage.CodeOrganisationType_stagedOrder, stage.CodeOrganisationType_orderStaged, &stage.CodeOrganisationTypeOrder, stage.CodeOrganisationTypes_mapString, codeorganisationtype, codeorganisationtype.Name)
	return codeorganisationtype
}

// StagePreserveOrder puts codeorganisationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOrganisationTypeOrder
// - update stage.CodeOrganisationTypeOrder accordingly
func (codeorganisationtype *CodeOrganisationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOrganisationTypes, stage.CodeOrganisationType_stagedOrder, stage.CodeOrganisationType_orderStaged, &stage.CodeOrganisationTypeOrder, stage.CodeOrganisationTypes_mapString, codeorganisationtype, order, codeorganisationtype.Name)
}

// Unstage removes codeorganisationtype off the model stage
func (codeorganisationtype *CodeOrganisationType) Unstage(stage *Stage) *CodeOrganisationType {
	__gong__unstage(stage.CodeOrganisationTypes, stage.CodeOrganisationTypes_mapString, codeorganisationtype, codeorganisationtype.Name)
	return codeorganisationtype
}

// UnstageVoid removes codeorganisationtype off the model stage
func (codeorganisationtype *CodeOrganisationType) UnstageVoid(stage *Stage) {
	codeorganisationtype.Unstage(stage)
}

func (codeorganisationtype *CodeOrganisationType) StageVoid(stage *Stage) {
	codeorganisationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeorganisationtype *CodeOrganisationType) GetName() (res string) {
	return codeorganisationtype.Name
}

// for satisfaction of GongStruct interface
func (codeorganisationtype *CodeOrganisationType) SetName(name string) {
	codeorganisationtype.Name = name
}

// Stage puts codeoxygentype to the model stage
func (codeoxygentype *CodeOxygenType) Stage(stage *Stage) *CodeOxygenType {
	__gong__stage(stage.CodeOxygenTypes, stage.CodeOxygenType_stagedOrder, stage.CodeOxygenType_orderStaged, &stage.CodeOxygenTypeOrder, stage.CodeOxygenTypes_mapString, codeoxygentype, codeoxygentype.Name)
	return codeoxygentype
}

// StagePreserveOrder puts codeoxygentype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOxygenTypeOrder
// - update stage.CodeOxygenTypeOrder accordingly
func (codeoxygentype *CodeOxygenType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOxygenTypes, stage.CodeOxygenType_stagedOrder, stage.CodeOxygenType_orderStaged, &stage.CodeOxygenTypeOrder, stage.CodeOxygenTypes_mapString, codeoxygentype, order, codeoxygentype.Name)
}

// Unstage removes codeoxygentype off the model stage
func (codeoxygentype *CodeOxygenType) Unstage(stage *Stage) *CodeOxygenType {
	__gong__unstage(stage.CodeOxygenTypes, stage.CodeOxygenTypes_mapString, codeoxygentype, codeoxygentype.Name)
	return codeoxygentype
}

// UnstageVoid removes codeoxygentype off the model stage
func (codeoxygentype *CodeOxygenType) UnstageVoid(stage *Stage) {
	codeoxygentype.Unstage(stage)
}

func (codeoxygentype *CodeOxygenType) StageVoid(stage *Stage) {
	codeoxygentype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeoxygentype *CodeOxygenType) GetName() (res string) {
	return codeoxygentype.Name
}

// for satisfaction of GongStruct interface
func (codeoxygentype *CodeOxygenType) SetName(name string) {
	codeoxygentype.Name = name
}

// Stage puts codepassengerservicetype to the model stage
func (codepassengerservicetype *CodePassengerServiceType) Stage(stage *Stage) *CodePassengerServiceType {
	__gong__stage(stage.CodePassengerServiceTypes, stage.CodePassengerServiceType_stagedOrder, stage.CodePassengerServiceType_orderStaged, &stage.CodePassengerServiceTypeOrder, stage.CodePassengerServiceTypes_mapString, codepassengerservicetype, codepassengerservicetype.Name)
	return codepassengerservicetype
}

// StagePreserveOrder puts codepassengerservicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePassengerServiceTypeOrder
// - update stage.CodePassengerServiceTypeOrder accordingly
func (codepassengerservicetype *CodePassengerServiceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePassengerServiceTypes, stage.CodePassengerServiceType_stagedOrder, stage.CodePassengerServiceType_orderStaged, &stage.CodePassengerServiceTypeOrder, stage.CodePassengerServiceTypes_mapString, codepassengerservicetype, order, codepassengerservicetype.Name)
}

// Unstage removes codepassengerservicetype off the model stage
func (codepassengerservicetype *CodePassengerServiceType) Unstage(stage *Stage) *CodePassengerServiceType {
	__gong__unstage(stage.CodePassengerServiceTypes, stage.CodePassengerServiceTypes_mapString, codepassengerservicetype, codepassengerservicetype.Name)
	return codepassengerservicetype
}

// UnstageVoid removes codepassengerservicetype off the model stage
func (codepassengerservicetype *CodePassengerServiceType) UnstageVoid(stage *Stage) {
	codepassengerservicetype.Unstage(stage)
}

func (codepassengerservicetype *CodePassengerServiceType) StageVoid(stage *Stage) {
	codepassengerservicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepassengerservicetype *CodePassengerServiceType) GetName() (res string) {
	return codepassengerservicetype.Name
}

// for satisfaction of GongStruct interface
func (codepassengerservicetype *CodePassengerServiceType) SetName(name string) {
	codepassengerservicetype.Name = name
}

// Stage puts codepilotcontrolledlightingtype to the model stage
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) Stage(stage *Stage) *CodePilotControlledLightingType {
	__gong__stage(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingType_stagedOrder, stage.CodePilotControlledLightingType_orderStaged, &stage.CodePilotControlledLightingTypeOrder, stage.CodePilotControlledLightingTypes_mapString, codepilotcontrolledlightingtype, codepilotcontrolledlightingtype.Name)
	return codepilotcontrolledlightingtype
}

// StagePreserveOrder puts codepilotcontrolledlightingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePilotControlledLightingTypeOrder
// - update stage.CodePilotControlledLightingTypeOrder accordingly
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingType_stagedOrder, stage.CodePilotControlledLightingType_orderStaged, &stage.CodePilotControlledLightingTypeOrder, stage.CodePilotControlledLightingTypes_mapString, codepilotcontrolledlightingtype, order, codepilotcontrolledlightingtype.Name)
}

// Unstage removes codepilotcontrolledlightingtype off the model stage
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) Unstage(stage *Stage) *CodePilotControlledLightingType {
	__gong__unstage(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingTypes_mapString, codepilotcontrolledlightingtype, codepilotcontrolledlightingtype.Name)
	return codepilotcontrolledlightingtype
}

// UnstageVoid removes codepilotcontrolledlightingtype off the model stage
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) UnstageVoid(stage *Stage) {
	codepilotcontrolledlightingtype.Unstage(stage)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) StageVoid(stage *Stage) {
	codepilotcontrolledlightingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GetName() (res string) {
	return codepilotcontrolledlightingtype.Name
}

// for satisfaction of GongStruct interface
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) SetName(name string) {
	codepilotcontrolledlightingtype.Name = name
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

// Stage puts codeserviceatctype to the model stage
func (codeserviceatctype *CodeServiceATCType) Stage(stage *Stage) *CodeServiceATCType {
	__gong__stage(stage.CodeServiceATCTypes, stage.CodeServiceATCType_stagedOrder, stage.CodeServiceATCType_orderStaged, &stage.CodeServiceATCTypeOrder, stage.CodeServiceATCTypes_mapString, codeserviceatctype, codeserviceatctype.Name)
	return codeserviceatctype
}

// StagePreserveOrder puts codeserviceatctype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeServiceATCTypeOrder
// - update stage.CodeServiceATCTypeOrder accordingly
func (codeserviceatctype *CodeServiceATCType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeServiceATCTypes, stage.CodeServiceATCType_stagedOrder, stage.CodeServiceATCType_orderStaged, &stage.CodeServiceATCTypeOrder, stage.CodeServiceATCTypes_mapString, codeserviceatctype, order, codeserviceatctype.Name)
}

// Unstage removes codeserviceatctype off the model stage
func (codeserviceatctype *CodeServiceATCType) Unstage(stage *Stage) *CodeServiceATCType {
	__gong__unstage(stage.CodeServiceATCTypes, stage.CodeServiceATCTypes_mapString, codeserviceatctype, codeserviceatctype.Name)
	return codeserviceatctype
}

// UnstageVoid removes codeserviceatctype off the model stage
func (codeserviceatctype *CodeServiceATCType) UnstageVoid(stage *Stage) {
	codeserviceatctype.Unstage(stage)
}

func (codeserviceatctype *CodeServiceATCType) StageVoid(stage *Stage) {
	codeserviceatctype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeserviceatctype *CodeServiceATCType) GetName() (res string) {
	return codeserviceatctype.Name
}

// for satisfaction of GongStruct interface
func (codeserviceatctype *CodeServiceATCType) SetName(name string) {
	codeserviceatctype.Name = name
}

// Stage puts codeserviceatfmtype to the model stage
func (codeserviceatfmtype *CodeServiceATFMType) Stage(stage *Stage) *CodeServiceATFMType {
	__gong__stage(stage.CodeServiceATFMTypes, stage.CodeServiceATFMType_stagedOrder, stage.CodeServiceATFMType_orderStaged, &stage.CodeServiceATFMTypeOrder, stage.CodeServiceATFMTypes_mapString, codeserviceatfmtype, codeserviceatfmtype.Name)
	return codeserviceatfmtype
}

// StagePreserveOrder puts codeserviceatfmtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeServiceATFMTypeOrder
// - update stage.CodeServiceATFMTypeOrder accordingly
func (codeserviceatfmtype *CodeServiceATFMType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeServiceATFMTypes, stage.CodeServiceATFMType_stagedOrder, stage.CodeServiceATFMType_orderStaged, &stage.CodeServiceATFMTypeOrder, stage.CodeServiceATFMTypes_mapString, codeserviceatfmtype, order, codeserviceatfmtype.Name)
}

// Unstage removes codeserviceatfmtype off the model stage
func (codeserviceatfmtype *CodeServiceATFMType) Unstage(stage *Stage) *CodeServiceATFMType {
	__gong__unstage(stage.CodeServiceATFMTypes, stage.CodeServiceATFMTypes_mapString, codeserviceatfmtype, codeserviceatfmtype.Name)
	return codeserviceatfmtype
}

// UnstageVoid removes codeserviceatfmtype off the model stage
func (codeserviceatfmtype *CodeServiceATFMType) UnstageVoid(stage *Stage) {
	codeserviceatfmtype.Unstage(stage)
}

func (codeserviceatfmtype *CodeServiceATFMType) StageVoid(stage *Stage) {
	codeserviceatfmtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeserviceatfmtype *CodeServiceATFMType) GetName() (res string) {
	return codeserviceatfmtype.Name
}

// for satisfaction of GongStruct interface
func (codeserviceatfmtype *CodeServiceATFMType) SetName(name string) {
	codeserviceatfmtype.Name = name
}

// Stage puts codeservicegroundcontroltype to the model stage
func (codeservicegroundcontroltype *CodeServiceGroundControlType) Stage(stage *Stage) *CodeServiceGroundControlType {
	__gong__stage(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlType_stagedOrder, stage.CodeServiceGroundControlType_orderStaged, &stage.CodeServiceGroundControlTypeOrder, stage.CodeServiceGroundControlTypes_mapString, codeservicegroundcontroltype, codeservicegroundcontroltype.Name)
	return codeservicegroundcontroltype
}

// StagePreserveOrder puts codeservicegroundcontroltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeServiceGroundControlTypeOrder
// - update stage.CodeServiceGroundControlTypeOrder accordingly
func (codeservicegroundcontroltype *CodeServiceGroundControlType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlType_stagedOrder, stage.CodeServiceGroundControlType_orderStaged, &stage.CodeServiceGroundControlTypeOrder, stage.CodeServiceGroundControlTypes_mapString, codeservicegroundcontroltype, order, codeservicegroundcontroltype.Name)
}

// Unstage removes codeservicegroundcontroltype off the model stage
func (codeservicegroundcontroltype *CodeServiceGroundControlType) Unstage(stage *Stage) *CodeServiceGroundControlType {
	__gong__unstage(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlTypes_mapString, codeservicegroundcontroltype, codeservicegroundcontroltype.Name)
	return codeservicegroundcontroltype
}

// UnstageVoid removes codeservicegroundcontroltype off the model stage
func (codeservicegroundcontroltype *CodeServiceGroundControlType) UnstageVoid(stage *Stage) {
	codeservicegroundcontroltype.Unstage(stage)
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) StageVoid(stage *Stage) {
	codeservicegroundcontroltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeservicegroundcontroltype *CodeServiceGroundControlType) GetName() (res string) {
	return codeservicegroundcontroltype.Name
}

// for satisfaction of GongStruct interface
func (codeservicegroundcontroltype *CodeServiceGroundControlType) SetName(name string) {
	codeservicegroundcontroltype.Name = name
}

// Stage puts codeserviceinformationtype to the model stage
func (codeserviceinformationtype *CodeServiceInformationType) Stage(stage *Stage) *CodeServiceInformationType {
	__gong__stage(stage.CodeServiceInformationTypes, stage.CodeServiceInformationType_stagedOrder, stage.CodeServiceInformationType_orderStaged, &stage.CodeServiceInformationTypeOrder, stage.CodeServiceInformationTypes_mapString, codeserviceinformationtype, codeserviceinformationtype.Name)
	return codeserviceinformationtype
}

// StagePreserveOrder puts codeserviceinformationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeServiceInformationTypeOrder
// - update stage.CodeServiceInformationTypeOrder accordingly
func (codeserviceinformationtype *CodeServiceInformationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeServiceInformationTypes, stage.CodeServiceInformationType_stagedOrder, stage.CodeServiceInformationType_orderStaged, &stage.CodeServiceInformationTypeOrder, stage.CodeServiceInformationTypes_mapString, codeserviceinformationtype, order, codeserviceinformationtype.Name)
}

// Unstage removes codeserviceinformationtype off the model stage
func (codeserviceinformationtype *CodeServiceInformationType) Unstage(stage *Stage) *CodeServiceInformationType {
	__gong__unstage(stage.CodeServiceInformationTypes, stage.CodeServiceInformationTypes_mapString, codeserviceinformationtype, codeserviceinformationtype.Name)
	return codeserviceinformationtype
}

// UnstageVoid removes codeserviceinformationtype off the model stage
func (codeserviceinformationtype *CodeServiceInformationType) UnstageVoid(stage *Stage) {
	codeserviceinformationtype.Unstage(stage)
}

func (codeserviceinformationtype *CodeServiceInformationType) StageVoid(stage *Stage) {
	codeserviceinformationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeserviceinformationtype *CodeServiceInformationType) GetName() (res string) {
	return codeserviceinformationtype.Name
}

// for satisfaction of GongStruct interface
func (codeserviceinformationtype *CodeServiceInformationType) SetName(name string) {
	codeserviceinformationtype.Name = name
}

// Stage puts codeservicesartype to the model stage
func (codeservicesartype *CodeServiceSARType) Stage(stage *Stage) *CodeServiceSARType {
	__gong__stage(stage.CodeServiceSARTypes, stage.CodeServiceSARType_stagedOrder, stage.CodeServiceSARType_orderStaged, &stage.CodeServiceSARTypeOrder, stage.CodeServiceSARTypes_mapString, codeservicesartype, codeservicesartype.Name)
	return codeservicesartype
}

// StagePreserveOrder puts codeservicesartype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeServiceSARTypeOrder
// - update stage.CodeServiceSARTypeOrder accordingly
func (codeservicesartype *CodeServiceSARType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeServiceSARTypes, stage.CodeServiceSARType_stagedOrder, stage.CodeServiceSARType_orderStaged, &stage.CodeServiceSARTypeOrder, stage.CodeServiceSARTypes_mapString, codeservicesartype, order, codeservicesartype.Name)
}

// Unstage removes codeservicesartype off the model stage
func (codeservicesartype *CodeServiceSARType) Unstage(stage *Stage) *CodeServiceSARType {
	__gong__unstage(stage.CodeServiceSARTypes, stage.CodeServiceSARTypes_mapString, codeservicesartype, codeservicesartype.Name)
	return codeservicesartype
}

// UnstageVoid removes codeservicesartype off the model stage
func (codeservicesartype *CodeServiceSARType) UnstageVoid(stage *Stage) {
	codeservicesartype.Unstage(stage)
}

func (codeservicesartype *CodeServiceSARType) StageVoid(stage *Stage) {
	codeservicesartype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeservicesartype *CodeServiceSARType) GetName() (res string) {
	return codeservicesartype.Name
}

// for satisfaction of GongStruct interface
func (codeservicesartype *CodeServiceSARType) SetName(name string) {
	codeservicesartype.Name = name
}

// Stage puts codespecialdatetype to the model stage
func (codespecialdatetype *CodeSpecialDateType) Stage(stage *Stage) *CodeSpecialDateType {
	__gong__stage(stage.CodeSpecialDateTypes, stage.CodeSpecialDateType_stagedOrder, stage.CodeSpecialDateType_orderStaged, &stage.CodeSpecialDateTypeOrder, stage.CodeSpecialDateTypes_mapString, codespecialdatetype, codespecialdatetype.Name)
	return codespecialdatetype
}

// StagePreserveOrder puts codespecialdatetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSpecialDateTypeOrder
// - update stage.CodeSpecialDateTypeOrder accordingly
func (codespecialdatetype *CodeSpecialDateType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSpecialDateTypes, stage.CodeSpecialDateType_stagedOrder, stage.CodeSpecialDateType_orderStaged, &stage.CodeSpecialDateTypeOrder, stage.CodeSpecialDateTypes_mapString, codespecialdatetype, order, codespecialdatetype.Name)
}

// Unstage removes codespecialdatetype off the model stage
func (codespecialdatetype *CodeSpecialDateType) Unstage(stage *Stage) *CodeSpecialDateType {
	__gong__unstage(stage.CodeSpecialDateTypes, stage.CodeSpecialDateTypes_mapString, codespecialdatetype, codespecialdatetype.Name)
	return codespecialdatetype
}

// UnstageVoid removes codespecialdatetype off the model stage
func (codespecialdatetype *CodeSpecialDateType) UnstageVoid(stage *Stage) {
	codespecialdatetype.Unstage(stage)
}

func (codespecialdatetype *CodeSpecialDateType) StageVoid(stage *Stage) {
	codespecialdatetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codespecialdatetype *CodeSpecialDateType) GetName() (res string) {
	return codespecialdatetype.Name
}

// for satisfaction of GongStruct interface
func (codespecialdatetype *CodeSpecialDateType) SetName(name string) {
	codespecialdatetype.Name = name
}

// Stage puts codestatusoperationstype to the model stage
func (codestatusoperationstype *CodeStatusOperationsType) Stage(stage *Stage) *CodeStatusOperationsType {
	__gong__stage(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsType_orderStaged, &stage.CodeStatusOperationsTypeOrder, stage.CodeStatusOperationsTypes_mapString, codestatusoperationstype, codestatusoperationstype.Name)
	return codestatusoperationstype
}

// StagePreserveOrder puts codestatusoperationstype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusOperationsTypeOrder
// - update stage.CodeStatusOperationsTypeOrder accordingly
func (codestatusoperationstype *CodeStatusOperationsType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsType_orderStaged, &stage.CodeStatusOperationsTypeOrder, stage.CodeStatusOperationsTypes_mapString, codestatusoperationstype, order, codestatusoperationstype.Name)
}

// Unstage removes codestatusoperationstype off the model stage
func (codestatusoperationstype *CodeStatusOperationsType) Unstage(stage *Stage) *CodeStatusOperationsType {
	__gong__unstage(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsTypes_mapString, codestatusoperationstype, codestatusoperationstype.Name)
	return codestatusoperationstype
}

// UnstageVoid removes codestatusoperationstype off the model stage
func (codestatusoperationstype *CodeStatusOperationsType) UnstageVoid(stage *Stage) {
	codestatusoperationstype.Unstage(stage)
}

func (codestatusoperationstype *CodeStatusOperationsType) StageVoid(stage *Stage) {
	codestatusoperationstype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusoperationstype *CodeStatusOperationsType) GetName() (res string) {
	return codestatusoperationstype.Name
}

// for satisfaction of GongStruct interface
func (codestatusoperationstype *CodeStatusOperationsType) SetName(name string) {
	codestatusoperationstype.Name = name
}

// Stage puts codestatusservicetype to the model stage
func (codestatusservicetype *CodeStatusServiceType) Stage(stage *Stage) *CodeStatusServiceType {
	__gong__stage(stage.CodeStatusServiceTypes, stage.CodeStatusServiceType_stagedOrder, stage.CodeStatusServiceType_orderStaged, &stage.CodeStatusServiceTypeOrder, stage.CodeStatusServiceTypes_mapString, codestatusservicetype, codestatusservicetype.Name)
	return codestatusservicetype
}

// StagePreserveOrder puts codestatusservicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusServiceTypeOrder
// - update stage.CodeStatusServiceTypeOrder accordingly
func (codestatusservicetype *CodeStatusServiceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusServiceTypes, stage.CodeStatusServiceType_stagedOrder, stage.CodeStatusServiceType_orderStaged, &stage.CodeStatusServiceTypeOrder, stage.CodeStatusServiceTypes_mapString, codestatusservicetype, order, codestatusservicetype.Name)
}

// Unstage removes codestatusservicetype off the model stage
func (codestatusservicetype *CodeStatusServiceType) Unstage(stage *Stage) *CodeStatusServiceType {
	__gong__unstage(stage.CodeStatusServiceTypes, stage.CodeStatusServiceTypes_mapString, codestatusservicetype, codestatusservicetype.Name)
	return codestatusservicetype
}

// UnstageVoid removes codestatusservicetype off the model stage
func (codestatusservicetype *CodeStatusServiceType) UnstageVoid(stage *Stage) {
	codestatusservicetype.Unstage(stage)
}

func (codestatusservicetype *CodeStatusServiceType) StageVoid(stage *Stage) {
	codestatusservicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusservicetype *CodeStatusServiceType) GetName() (res string) {
	return codestatusservicetype.Name
}

// for satisfaction of GongStruct interface
func (codestatusservicetype *CodeStatusServiceType) SetName(name string) {
	codestatusservicetype.Name = name
}

// Stage puts codesystemactivationtype to the model stage
func (codesystemactivationtype *CodeSystemActivationType) Stage(stage *Stage) *CodeSystemActivationType {
	__gong__stage(stage.CodeSystemActivationTypes, stage.CodeSystemActivationType_stagedOrder, stage.CodeSystemActivationType_orderStaged, &stage.CodeSystemActivationTypeOrder, stage.CodeSystemActivationTypes_mapString, codesystemactivationtype, codesystemactivationtype.Name)
	return codesystemactivationtype
}

// StagePreserveOrder puts codesystemactivationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSystemActivationTypeOrder
// - update stage.CodeSystemActivationTypeOrder accordingly
func (codesystemactivationtype *CodeSystemActivationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSystemActivationTypes, stage.CodeSystemActivationType_stagedOrder, stage.CodeSystemActivationType_orderStaged, &stage.CodeSystemActivationTypeOrder, stage.CodeSystemActivationTypes_mapString, codesystemactivationtype, order, codesystemactivationtype.Name)
}

// Unstage removes codesystemactivationtype off the model stage
func (codesystemactivationtype *CodeSystemActivationType) Unstage(stage *Stage) *CodeSystemActivationType {
	__gong__unstage(stage.CodeSystemActivationTypes, stage.CodeSystemActivationTypes_mapString, codesystemactivationtype, codesystemactivationtype.Name)
	return codesystemactivationtype
}

// UnstageVoid removes codesystemactivationtype off the model stage
func (codesystemactivationtype *CodeSystemActivationType) UnstageVoid(stage *Stage) {
	codesystemactivationtype.Unstage(stage)
}

func (codesystemactivationtype *CodeSystemActivationType) StageVoid(stage *Stage) {
	codesystemactivationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesystemactivationtype *CodeSystemActivationType) GetName() (res string) {
	return codesystemactivationtype.Name
}

// for satisfaction of GongStruct interface
func (codesystemactivationtype *CodeSystemActivationType) SetName(name string) {
	codesystemactivationtype.Name = name
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

// Stage puts codeunitdependencytype to the model stage
func (codeunitdependencytype *CodeUnitDependencyType) Stage(stage *Stage) *CodeUnitDependencyType {
	__gong__stage(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyType_stagedOrder, stage.CodeUnitDependencyType_orderStaged, &stage.CodeUnitDependencyTypeOrder, stage.CodeUnitDependencyTypes_mapString, codeunitdependencytype, codeunitdependencytype.Name)
	return codeunitdependencytype
}

// StagePreserveOrder puts codeunitdependencytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeUnitDependencyTypeOrder
// - update stage.CodeUnitDependencyTypeOrder accordingly
func (codeunitdependencytype *CodeUnitDependencyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyType_stagedOrder, stage.CodeUnitDependencyType_orderStaged, &stage.CodeUnitDependencyTypeOrder, stage.CodeUnitDependencyTypes_mapString, codeunitdependencytype, order, codeunitdependencytype.Name)
}

// Unstage removes codeunitdependencytype off the model stage
func (codeunitdependencytype *CodeUnitDependencyType) Unstage(stage *Stage) *CodeUnitDependencyType {
	__gong__unstage(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyTypes_mapString, codeunitdependencytype, codeunitdependencytype.Name)
	return codeunitdependencytype
}

// UnstageVoid removes codeunitdependencytype off the model stage
func (codeunitdependencytype *CodeUnitDependencyType) UnstageVoid(stage *Stage) {
	codeunitdependencytype.Unstage(stage)
}

func (codeunitdependencytype *CodeUnitDependencyType) StageVoid(stage *Stage) {
	codeunitdependencytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeunitdependencytype *CodeUnitDependencyType) GetName() (res string) {
	return codeunitdependencytype.Name
}

// for satisfaction of GongStruct interface
func (codeunitdependencytype *CodeUnitDependencyType) SetName(name string) {
	codeunitdependencytype.Name = name
}

// Stage puts codeunittype to the model stage
func (codeunittype *CodeUnitType) Stage(stage *Stage) *CodeUnitType {
	__gong__stage(stage.CodeUnitTypes, stage.CodeUnitType_stagedOrder, stage.CodeUnitType_orderStaged, &stage.CodeUnitTypeOrder, stage.CodeUnitTypes_mapString, codeunittype, codeunittype.Name)
	return codeunittype
}

// StagePreserveOrder puts codeunittype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeUnitTypeOrder
// - update stage.CodeUnitTypeOrder accordingly
func (codeunittype *CodeUnitType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeUnitTypes, stage.CodeUnitType_stagedOrder, stage.CodeUnitType_orderStaged, &stage.CodeUnitTypeOrder, stage.CodeUnitTypes_mapString, codeunittype, order, codeunittype.Name)
}

// Unstage removes codeunittype off the model stage
func (codeunittype *CodeUnitType) Unstage(stage *Stage) *CodeUnitType {
	__gong__unstage(stage.CodeUnitTypes, stage.CodeUnitTypes_mapString, codeunittype, codeunittype.Name)
	return codeunittype
}

// UnstageVoid removes codeunittype off the model stage
func (codeunittype *CodeUnitType) UnstageVoid(stage *Stage) {
	codeunittype.Unstage(stage)
}

func (codeunittype *CodeUnitType) StageVoid(stage *Stage) {
	codeunittype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeunittype *CodeUnitType) GetName() (res string) {
	return codeunittype.Name
}

// for satisfaction of GongStruct interface
func (codeunittype *CodeUnitType) SetName(name string) {
	codeunittype.Name = name
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

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeAircraftGroundServiceTypes, &stage.CodeAircraftGroundServiceTypes_mapString, &stage.CodeAircraftGroundServiceType_stagedOrder, &stage.CodeAircraftGroundServiceTypeOrder)

	__gong__resetStageType(&stage.CodeAviationStandardsTypes, &stage.CodeAviationStandardsTypes_mapString, &stage.CodeAviationStandardsType_stagedOrder, &stage.CodeAviationStandardsTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationChannelTypes, &stage.CodeCommunicationChannelTypes_mapString, &stage.CodeCommunicationChannelType_stagedOrder, &stage.CodeCommunicationChannelTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationDirectionTypes, &stage.CodeCommunicationDirectionTypes_mapString, &stage.CodeCommunicationDirectionType_stagedOrder, &stage.CodeCommunicationDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString, &stage.CodeCommunicationModeType_stagedOrder, &stage.CodeCommunicationModeTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeFacilityRankingTypes, &stage.CodeFacilityRankingTypes_mapString, &stage.CodeFacilityRankingType_stagedOrder, &stage.CodeFacilityRankingTypeOrder)

	__gong__resetStageType(&stage.CodeFireFightingTypes, &stage.CodeFireFightingTypes_mapString, &stage.CodeFireFightingType_stagedOrder, &stage.CodeFireFightingTypeOrder)

	__gong__resetStageType(&stage.CodeFlightDestinationTypes, &stage.CodeFlightDestinationTypes_mapString, &stage.CodeFlightDestinationType_stagedOrder, &stage.CodeFlightDestinationTypeOrder)

	__gong__resetStageType(&stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString, &stage.CodeFlightRuleType_stagedOrder, &stage.CodeFlightRuleTypeOrder)

	__gong__resetStageType(&stage.CodeFuelTypes, &stage.CodeFuelTypes_mapString, &stage.CodeFuelType_stagedOrder, &stage.CodeFuelTypeOrder)

	__gong__resetStageType(&stage.CodeIntensityStandByTypes, &stage.CodeIntensityStandByTypes_mapString, &stage.CodeIntensityStandByType_stagedOrder, &stage.CodeIntensityStandByTypeOrder)

	__gong__resetStageType(&stage.CodeLanguageTypes, &stage.CodeLanguageTypes_mapString, &stage.CodeLanguageType_stagedOrder, &stage.CodeLanguageTypeOrder)

	__gong__resetStageType(&stage.CodeLevelSeriesTypes, &stage.CodeLevelSeriesTypes_mapString, &stage.CodeLevelSeriesType_stagedOrder, &stage.CodeLevelSeriesTypeOrder)

	__gong__resetStageType(&stage.CodeLevelTableDesignatorTypes, &stage.CodeLevelTableDesignatorTypes_mapString, &stage.CodeLevelTableDesignatorType_stagedOrder, &stage.CodeLevelTableDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString, &stage.CodeLightIntensityType_stagedOrder, &stage.CodeLightIntensityTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString, &stage.CodeMilitaryOperationsType_stagedOrder, &stage.CodeMilitaryOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeNitrogenTypes, &stage.CodeNitrogenTypes_mapString, &stage.CodeNitrogenType_stagedOrder, &stage.CodeNitrogenTypeOrder)

	__gong__resetStageType(&stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_mapString, &stage.CodeNorthReferenceType_stagedOrder, &stage.CodeNorthReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeOilTypes, &stage.CodeOilTypes_mapString, &stage.CodeOilType_stagedOrder, &stage.CodeOilTypeOrder)

	__gong__resetStageType(&stage.CodeOrganisationDesignatorTypes, &stage.CodeOrganisationDesignatorTypes_mapString, &stage.CodeOrganisationDesignatorType_stagedOrder, &stage.CodeOrganisationDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeOrganisationHierarchyTypes, &stage.CodeOrganisationHierarchyTypes_mapString, &stage.CodeOrganisationHierarchyType_stagedOrder, &stage.CodeOrganisationHierarchyTypeOrder)

	__gong__resetStageType(&stage.CodeOrganisationTypes, &stage.CodeOrganisationTypes_mapString, &stage.CodeOrganisationType_stagedOrder, &stage.CodeOrganisationTypeOrder)

	__gong__resetStageType(&stage.CodeOxygenTypes, &stage.CodeOxygenTypes_mapString, &stage.CodeOxygenType_stagedOrder, &stage.CodeOxygenTypeOrder)

	__gong__resetStageType(&stage.CodePassengerServiceTypes, &stage.CodePassengerServiceTypes_mapString, &stage.CodePassengerServiceType_stagedOrder, &stage.CodePassengerServiceTypeOrder)

	__gong__resetStageType(&stage.CodePilotControlledLightingTypes, &stage.CodePilotControlledLightingTypes_mapString, &stage.CodePilotControlledLightingType_stagedOrder, &stage.CodePilotControlledLightingTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString, &stage.CodeRVSMType_stagedOrder, &stage.CodeRVSMTypeOrder)

	__gong__resetStageType(&stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_mapString, &stage.CodeRadioEmissionType_stagedOrder, &stage.CodeRadioEmissionTypeOrder)

	__gong__resetStageType(&stage.CodeServiceATCTypes, &stage.CodeServiceATCTypes_mapString, &stage.CodeServiceATCType_stagedOrder, &stage.CodeServiceATCTypeOrder)

	__gong__resetStageType(&stage.CodeServiceATFMTypes, &stage.CodeServiceATFMTypes_mapString, &stage.CodeServiceATFMType_stagedOrder, &stage.CodeServiceATFMTypeOrder)

	__gong__resetStageType(&stage.CodeServiceGroundControlTypes, &stage.CodeServiceGroundControlTypes_mapString, &stage.CodeServiceGroundControlType_stagedOrder, &stage.CodeServiceGroundControlTypeOrder)

	__gong__resetStageType(&stage.CodeServiceInformationTypes, &stage.CodeServiceInformationTypes_mapString, &stage.CodeServiceInformationType_stagedOrder, &stage.CodeServiceInformationTypeOrder)

	__gong__resetStageType(&stage.CodeServiceSARTypes, &stage.CodeServiceSARTypes_mapString, &stage.CodeServiceSARType_stagedOrder, &stage.CodeServiceSARTypeOrder)

	__gong__resetStageType(&stage.CodeSpecialDateTypes, &stage.CodeSpecialDateTypes_mapString, &stage.CodeSpecialDateType_stagedOrder, &stage.CodeSpecialDateTypeOrder)

	__gong__resetStageType(&stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString, &stage.CodeStatusOperationsType_stagedOrder, &stage.CodeStatusOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeStatusServiceTypes, &stage.CodeStatusServiceTypes_mapString, &stage.CodeStatusServiceType_stagedOrder, &stage.CodeStatusServiceTypeOrder)

	__gong__resetStageType(&stage.CodeSystemActivationTypes, &stage.CodeSystemActivationTypes_mapString, &stage.CodeSystemActivationType_stagedOrder, &stage.CodeSystemActivationTypeOrder)

	__gong__resetStageType(&stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString, &stage.CodeTelecomNetworkType_stagedOrder, &stage.CodeTelecomNetworkTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeUnitDependencyTypes, &stage.CodeUnitDependencyTypes_mapString, &stage.CodeUnitDependencyType_stagedOrder, &stage.CodeUnitDependencyTypeOrder)

	__gong__resetStageType(&stage.CodeUnitTypes, &stage.CodeUnitTypes_mapString, &stage.CodeUnitType_stagedOrder, &stage.CodeUnitTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.DateYearTypes, &stage.DateYearTypes_mapString, &stage.DateYearType_stagedOrder, &stage.DateYearTypeOrder)

	__gong__resetStageType(&stage.NoNumberTypes, &stage.NoNumberTypes_mapString, &stage.NoNumberType_stagedOrder, &stage.NoNumberTypeOrder)

	__gong__resetStageType(&stage.TextAddressTypes, &stage.TextAddressTypes_mapString, &stage.TextAddressType_stagedOrder, &stage.TextAddressTypeOrder)

	__gong__resetStageType(&stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString, &stage.TextDesignatorType_stagedOrder, &stage.TextDesignatorTypeOrder)

	__gong__resetStageType(&stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString, &stage.TextInstructionType_stagedOrder, &stage.TextInstructionTypeOrder)

	__gong__resetStageType(&stage.TextNameTypes, &stage.TextNameTypes_mapString, &stage.TextNameType_stagedOrder, &stage.TextNameTypeOrder)

	__gong__resetStageType(&stage.TextNoteTypes, &stage.TextNoteTypes_mapString, &stage.TextNoteType_stagedOrder, &stage.TextNoteTypeOrder)

	__gong__resetStageType(&stage.TextPhoneTypes, &stage.TextPhoneTypes_mapString, &stage.TextPhoneType_stagedOrder, &stage.TextPhoneTypeOrder)

	__gong__resetStageType(&stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString, &stage.TextPropertyNameType_stagedOrder, &stage.TextPropertyNameTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.ValBearingTypes, &stage.ValBearingTypes_mapString, &stage.ValBearingType_stagedOrder, &stage.ValBearingTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValFrequencyTypes, &stage.ValFrequencyTypes_mapString, &stage.ValFrequencyType_stagedOrder, &stage.ValFrequencyTypeOrder)

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
func (CodeAircraftGroundServiceType) GongGetAssociationName() any {
	return &CodeAircraftGroundServiceType{
	}
}

func (CodeAviationStandardsType) GongGetAssociationName() any {
	return &CodeAviationStandardsType{
	}
}

func (CodeCommunicationChannelType) GongGetAssociationName() any {
	return &CodeCommunicationChannelType{
	}
}

func (CodeCommunicationDirectionType) GongGetAssociationName() any {
	return &CodeCommunicationDirectionType{
	}
}

func (CodeCommunicationModeType) GongGetAssociationName() any {
	return &CodeCommunicationModeType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeFacilityRankingType) GongGetAssociationName() any {
	return &CodeFacilityRankingType{
	}
}

func (CodeFireFightingType) GongGetAssociationName() any {
	return &CodeFireFightingType{
	}
}

func (CodeFlightDestinationType) GongGetAssociationName() any {
	return &CodeFlightDestinationType{
	}
}

func (CodeFlightRuleType) GongGetAssociationName() any {
	return &CodeFlightRuleType{
	}
}

func (CodeFuelType) GongGetAssociationName() any {
	return &CodeFuelType{
	}
}

func (CodeIntensityStandByType) GongGetAssociationName() any {
	return &CodeIntensityStandByType{
	}
}

func (CodeLanguageType) GongGetAssociationName() any {
	return &CodeLanguageType{
	}
}

func (CodeLevelSeriesType) GongGetAssociationName() any {
	return &CodeLevelSeriesType{
	}
}

func (CodeLevelTableDesignatorType) GongGetAssociationName() any {
	return &CodeLevelTableDesignatorType{
	}
}

func (CodeLightIntensityType) GongGetAssociationName() any {
	return &CodeLightIntensityType{
	}
}

func (CodeMilitaryOperationsType) GongGetAssociationName() any {
	return &CodeMilitaryOperationsType{
	}
}

func (CodeNitrogenType) GongGetAssociationName() any {
	return &CodeNitrogenType{
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

func (CodeOilType) GongGetAssociationName() any {
	return &CodeOilType{
	}
}

func (CodeOrganisationDesignatorType) GongGetAssociationName() any {
	return &CodeOrganisationDesignatorType{
	}
}

func (CodeOrganisationHierarchyType) GongGetAssociationName() any {
	return &CodeOrganisationHierarchyType{
	}
}

func (CodeOrganisationType) GongGetAssociationName() any {
	return &CodeOrganisationType{
	}
}

func (CodeOxygenType) GongGetAssociationName() any {
	return &CodeOxygenType{
	}
}

func (CodePassengerServiceType) GongGetAssociationName() any {
	return &CodePassengerServiceType{
	}
}

func (CodePilotControlledLightingType) GongGetAssociationName() any {
	return &CodePilotControlledLightingType{
	}
}

func (CodeRVSMType) GongGetAssociationName() any {
	return &CodeRVSMType{
	}
}

func (CodeRadioEmissionType) GongGetAssociationName() any {
	return &CodeRadioEmissionType{
	}
}

func (CodeServiceATCType) GongGetAssociationName() any {
	return &CodeServiceATCType{
	}
}

func (CodeServiceATFMType) GongGetAssociationName() any {
	return &CodeServiceATFMType{
	}
}

func (CodeServiceGroundControlType) GongGetAssociationName() any {
	return &CodeServiceGroundControlType{
	}
}

func (CodeServiceInformationType) GongGetAssociationName() any {
	return &CodeServiceInformationType{
	}
}

func (CodeServiceSARType) GongGetAssociationName() any {
	return &CodeServiceSARType{
	}
}

func (CodeSpecialDateType) GongGetAssociationName() any {
	return &CodeSpecialDateType{
	}
}

func (CodeStatusOperationsType) GongGetAssociationName() any {
	return &CodeStatusOperationsType{
	}
}

func (CodeStatusServiceType) GongGetAssociationName() any {
	return &CodeStatusServiceType{
	}
}

func (CodeSystemActivationType) GongGetAssociationName() any {
	return &CodeSystemActivationType{
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

func (CodeUnitDependencyType) GongGetAssociationName() any {
	return &CodeUnitDependencyType{
	}
}

func (CodeUnitType) GongGetAssociationName() any {
	return &CodeUnitType{
	}
}

func (CodeVerticalDatumType) GongGetAssociationName() any {
	return &CodeVerticalDatumType{
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

func (NoNumberType) GongGetAssociationName() any {
	return &NoNumberType{
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
	// reverse maps of direct associations of CodeAircraftGroundServiceType
	case CodeAircraftGroundServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAviationStandardsType
	case CodeAviationStandardsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationChannelType
	case CodeCommunicationChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationDirectionType
	case CodeCommunicationDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFacilityRankingType
	case CodeFacilityRankingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFireFightingType
	case CodeFireFightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightDestinationType
	case CodeFlightDestinationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRuleType
	case CodeFlightRuleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFuelType
	case CodeFuelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIntensityStandByType
	case CodeIntensityStandByType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLanguageType
	case CodeLanguageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelSeriesType
	case CodeLevelSeriesType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelTableDesignatorType
	case CodeLevelTableDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightIntensityType
	case CodeLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNitrogenType
	case CodeNitrogenType:
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
	// reverse maps of direct associations of CodeOilType
	case CodeOilType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationDesignatorType
	case CodeOrganisationDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationHierarchyType
	case CodeOrganisationHierarchyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationType
	case CodeOrganisationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOxygenType
	case CodeOxygenType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePassengerServiceType
	case CodePassengerServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePilotControlledLightingType
	case CodePilotControlledLightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioEmissionType
	case CodeRadioEmissionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceATCType
	case CodeServiceATCType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceATFMType
	case CodeServiceATFMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceGroundControlType
	case CodeServiceGroundControlType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceInformationType
	case CodeServiceInformationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceSARType
	case CodeServiceSARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialDateType
	case CodeSpecialDateType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusServiceType
	case CodeStatusServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSystemActivationType
	case CodeSystemActivationType:
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
	// reverse maps of direct associations of CodeUnitDependencyType
	case CodeUnitDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeUnitType
	case CodeUnitType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
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
	// reverse maps of direct associations of NoNumberType
	case NoNumberType:
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
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CodeAircraftGroundServiceType
	case CodeAircraftGroundServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAviationStandardsType
	case CodeAviationStandardsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationChannelType
	case CodeCommunicationChannelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationDirectionType
	case CodeCommunicationDirectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFacilityRankingType
	case CodeFacilityRankingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFireFightingType
	case CodeFireFightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightDestinationType
	case CodeFlightDestinationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFlightRuleType
	case CodeFlightRuleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFuelType
	case CodeFuelType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIntensityStandByType
	case CodeIntensityStandByType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLanguageType
	case CodeLanguageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelSeriesType
	case CodeLevelSeriesType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLevelTableDesignatorType
	case CodeLevelTableDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightIntensityType
	case CodeLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeNitrogenType
	case CodeNitrogenType:
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
	// reverse maps of direct associations of CodeOilType
	case CodeOilType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationDesignatorType
	case CodeOrganisationDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationHierarchyType
	case CodeOrganisationHierarchyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOrganisationType
	case CodeOrganisationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOxygenType
	case CodeOxygenType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePassengerServiceType
	case CodePassengerServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePilotControlledLightingType
	case CodePilotControlledLightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRadioEmissionType
	case CodeRadioEmissionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceATCType
	case CodeServiceATCType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceATFMType
	case CodeServiceATFMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceGroundControlType
	case CodeServiceGroundControlType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceInformationType
	case CodeServiceInformationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeServiceSARType
	case CodeServiceSARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSpecialDateType
	case CodeSpecialDateType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusServiceType
	case CodeStatusServiceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSystemActivationType
	case CodeSystemActivationType:
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
	// reverse maps of direct associations of CodeUnitDependencyType
	case CodeUnitDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeUnitType
	case CodeUnitType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
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
	// reverse maps of direct associations of NoNumberType
	case NoNumberType:
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
func (*CodeAircraftGroundServiceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAviationStandardsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationChannelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationDirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationModeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFacilityRankingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFireFightingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightDestinationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFlightRuleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFuelType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeIntensityStandByType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLanguageType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLevelSeriesType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLevelTableDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLightIntensityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryOperationsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeNitrogenType) GongGetReverseFields() []GongReverseField {
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

func (*CodeOilType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOrganisationDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOrganisationHierarchyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOrganisationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOxygenType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePassengerServiceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePilotControlledLightingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRVSMType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRadioEmissionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeServiceATCType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeServiceATFMType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeServiceGroundControlType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeServiceInformationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeServiceSARType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSpecialDateType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusOperationsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusServiceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSystemActivationType) GongGetReverseFields() []GongReverseField {
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

func (*CodeUnitDependencyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeUnitType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalDatumType) GongGetReverseFields() []GongReverseField {
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

func (*NoNumberType) GongGetReverseFields() []GongReverseField {
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
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codefirefightingtype *CodeFireFightingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codefueltype *CodeFuelType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codelanguagetype *CodeLanguageType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codelevelseriestype *CodeLevelSeriesType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codelightintensitytype *CodeLightIntensityType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codenitrogentype *CodeNitrogenType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeoiltype *CodeOilType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeorganisationtype *CodeOrganisationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeoxygentype *CodeOxygenType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codepassengerservicetype *CodePassengerServiceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeserviceatctype *CodeServiceATCType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeserviceatfmtype *CodeServiceATFMType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeserviceinformationtype *CodeServiceInformationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeservicesartype *CodeServiceSARType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codespecialdatetype *CodeSpecialDateType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codestatusoperationstype *CodeStatusOperationsType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codestatusservicetype *CodeStatusServiceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codesystemactivationtype *CodeSystemActivationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeunitdependencytype *CodeUnitDependencyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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

func (codeunittype *CodeUnitType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
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
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftgroundservicetype.Name
	case "NilReason":
		res.valueString = codeaircraftgroundservicetype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftgroundservicetype.EnclosedText
	}
	return
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaviationstandardstype.Name
	case "NilReason":
		res.valueString = codeaviationstandardstype.NilReason
	case "EnclosedText":
		res.valueString = codeaviationstandardstype.EnclosedText
	}
	return
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecommunicationchanneltype.Name
	case "NilReason":
		res.valueString = codecommunicationchanneltype.NilReason
	case "EnclosedText":
		res.valueString = codecommunicationchanneltype.EnclosedText
	}
	return
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecommunicationdirectiontype.Name
	case "NilReason":
		res.valueString = codecommunicationdirectiontype.NilReason
	case "EnclosedText":
		res.valueString = codecommunicationdirectiontype.EnclosedText
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

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefacilityrankingtype.Name
	case "NilReason":
		res.valueString = codefacilityrankingtype.NilReason
	case "EnclosedText":
		res.valueString = codefacilityrankingtype.EnclosedText
	}
	return
}

func (codefirefightingtype *CodeFireFightingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefirefightingtype.Name
	case "NilReason":
		res.valueString = codefirefightingtype.NilReason
	case "EnclosedText":
		res.valueString = codefirefightingtype.EnclosedText
	}
	return
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeflightdestinationtype.Name
	case "NilReason":
		res.valueString = codeflightdestinationtype.NilReason
	case "EnclosedText":
		res.valueString = codeflightdestinationtype.EnclosedText
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

func (codefueltype *CodeFuelType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefueltype.Name
	case "NilReason":
		res.valueString = codefueltype.NilReason
	case "EnclosedText":
		res.valueString = codefueltype.EnclosedText
	}
	return
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeintensitystandbytype.Name
	case "NilReason":
		res.valueString = codeintensitystandbytype.NilReason
	case "EnclosedText":
		res.valueString = codeintensitystandbytype.EnclosedText
	}
	return
}

func (codelanguagetype *CodeLanguageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelanguagetype.Name
	case "NilReason":
		res.valueString = codelanguagetype.NilReason
	case "EnclosedText":
		res.valueString = codelanguagetype.EnclosedText
	}
	return
}

func (codelevelseriestype *CodeLevelSeriesType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelevelseriestype.Name
	case "NilReason":
		res.valueString = codelevelseriestype.NilReason
	case "EnclosedText":
		res.valueString = codelevelseriestype.EnclosedText
	}
	return
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeleveltabledesignatortype.Name
	case "NilReason":
		res.valueString = codeleveltabledesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codeleveltabledesignatortype.EnclosedText
	}
	return
}

func (codelightintensitytype *CodeLightIntensityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelightintensitytype.Name
	case "NilReason":
		res.valueString = codelightintensitytype.NilReason
	case "EnclosedText":
		res.valueString = codelightintensitytype.EnclosedText
	}
	return
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemilitaryoperationstype.Name
	case "NilReason":
		res.valueString = codemilitaryoperationstype.NilReason
	case "EnclosedText":
		res.valueString = codemilitaryoperationstype.EnclosedText
	}
	return
}

func (codenitrogentype *CodeNitrogenType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codenitrogentype.Name
	case "NilReason":
		res.valueString = codenitrogentype.NilReason
	case "EnclosedText":
		res.valueString = codenitrogentype.EnclosedText
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

func (codeoiltype *CodeOilType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeoiltype.Name
	case "NilReason":
		res.valueString = codeoiltype.NilReason
	case "EnclosedText":
		res.valueString = codeoiltype.EnclosedText
	}
	return
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeorganisationdesignatortype.Name
	case "NilReason":
		res.valueString = codeorganisationdesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codeorganisationdesignatortype.EnclosedText
	}
	return
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeorganisationhierarchytype.Name
	case "NilReason":
		res.valueString = codeorganisationhierarchytype.NilReason
	case "EnclosedText":
		res.valueString = codeorganisationhierarchytype.EnclosedText
	}
	return
}

func (codeorganisationtype *CodeOrganisationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeorganisationtype.Name
	case "NilReason":
		res.valueString = codeorganisationtype.NilReason
	case "EnclosedText":
		res.valueString = codeorganisationtype.EnclosedText
	}
	return
}

func (codeoxygentype *CodeOxygenType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeoxygentype.Name
	case "NilReason":
		res.valueString = codeoxygentype.NilReason
	case "EnclosedText":
		res.valueString = codeoxygentype.EnclosedText
	}
	return
}

func (codepassengerservicetype *CodePassengerServiceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepassengerservicetype.Name
	case "NilReason":
		res.valueString = codepassengerservicetype.NilReason
	case "EnclosedText":
		res.valueString = codepassengerservicetype.EnclosedText
	}
	return
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepilotcontrolledlightingtype.Name
	case "NilReason":
		res.valueString = codepilotcontrolledlightingtype.NilReason
	case "EnclosedText":
		res.valueString = codepilotcontrolledlightingtype.EnclosedText
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

func (codeserviceatctype *CodeServiceATCType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeserviceatctype.Name
	case "NilReason":
		res.valueString = codeserviceatctype.NilReason
	case "EnclosedText":
		res.valueString = codeserviceatctype.EnclosedText
	}
	return
}

func (codeserviceatfmtype *CodeServiceATFMType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeserviceatfmtype.Name
	case "NilReason":
		res.valueString = codeserviceatfmtype.NilReason
	case "EnclosedText":
		res.valueString = codeserviceatfmtype.EnclosedText
	}
	return
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeservicegroundcontroltype.Name
	case "NilReason":
		res.valueString = codeservicegroundcontroltype.NilReason
	case "EnclosedText":
		res.valueString = codeservicegroundcontroltype.EnclosedText
	}
	return
}

func (codeserviceinformationtype *CodeServiceInformationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeserviceinformationtype.Name
	case "NilReason":
		res.valueString = codeserviceinformationtype.NilReason
	case "EnclosedText":
		res.valueString = codeserviceinformationtype.EnclosedText
	}
	return
}

func (codeservicesartype *CodeServiceSARType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeservicesartype.Name
	case "NilReason":
		res.valueString = codeservicesartype.NilReason
	case "EnclosedText":
		res.valueString = codeservicesartype.EnclosedText
	}
	return
}

func (codespecialdatetype *CodeSpecialDateType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codespecialdatetype.Name
	case "NilReason":
		res.valueString = codespecialdatetype.NilReason
	case "EnclosedText":
		res.valueString = codespecialdatetype.EnclosedText
	}
	return
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusoperationstype.Name
	case "NilReason":
		res.valueString = codestatusoperationstype.NilReason
	case "EnclosedText":
		res.valueString = codestatusoperationstype.EnclosedText
	}
	return
}

func (codestatusservicetype *CodeStatusServiceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusservicetype.Name
	case "NilReason":
		res.valueString = codestatusservicetype.NilReason
	case "EnclosedText":
		res.valueString = codestatusservicetype.EnclosedText
	}
	return
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesystemactivationtype.Name
	case "NilReason":
		res.valueString = codesystemactivationtype.NilReason
	case "EnclosedText":
		res.valueString = codesystemactivationtype.EnclosedText
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

func (codeunitdependencytype *CodeUnitDependencyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeunitdependencytype.Name
	case "NilReason":
		res.valueString = codeunitdependencytype.NilReason
	case "EnclosedText":
		res.valueString = codeunitdependencytype.EnclosedText
	}
	return
}

func (codeunittype *CodeUnitType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeunittype.Name
	case "NilReason":
		res.valueString = codeunittype.NilReason
	case "EnclosedText":
		res.valueString = codeunittype.EnclosedText
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetGongstructName() string {
	return "CodeAircraftGroundServiceType"
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetGongstructName() string {
	return "CodeAviationStandardsType"
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetGongstructName() string {
	return "CodeCommunicationChannelType"
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetGongstructName() string {
	return "CodeCommunicationDirectionType"
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetGongstructName() string {
	return "CodeCommunicationModeType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetGongstructName() string {
	return "CodeFacilityRankingType"
}

func (codefirefightingtype *CodeFireFightingType) GongGetGongstructName() string {
	return "CodeFireFightingType"
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetGongstructName() string {
	return "CodeFlightDestinationType"
}

func (codeflightruletype *CodeFlightRuleType) GongGetGongstructName() string {
	return "CodeFlightRuleType"
}

func (codefueltype *CodeFuelType) GongGetGongstructName() string {
	return "CodeFuelType"
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetGongstructName() string {
	return "CodeIntensityStandByType"
}

func (codelanguagetype *CodeLanguageType) GongGetGongstructName() string {
	return "CodeLanguageType"
}

func (codelevelseriestype *CodeLevelSeriesType) GongGetGongstructName() string {
	return "CodeLevelSeriesType"
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetGongstructName() string {
	return "CodeLevelTableDesignatorType"
}

func (codelightintensitytype *CodeLightIntensityType) GongGetGongstructName() string {
	return "CodeLightIntensityType"
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetGongstructName() string {
	return "CodeMilitaryOperationsType"
}

func (codenitrogentype *CodeNitrogenType) GongGetGongstructName() string {
	return "CodeNitrogenType"
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetGongstructName() string {
	return "CodeNorthReferenceType"
}

func (codenotepurposetype *CodeNotePurposeType) GongGetGongstructName() string {
	return "CodeNotePurposeType"
}

func (codeoiltype *CodeOilType) GongGetGongstructName() string {
	return "CodeOilType"
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetGongstructName() string {
	return "CodeOrganisationDesignatorType"
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetGongstructName() string {
	return "CodeOrganisationHierarchyType"
}

func (codeorganisationtype *CodeOrganisationType) GongGetGongstructName() string {
	return "CodeOrganisationType"
}

func (codeoxygentype *CodeOxygenType) GongGetGongstructName() string {
	return "CodeOxygenType"
}

func (codepassengerservicetype *CodePassengerServiceType) GongGetGongstructName() string {
	return "CodePassengerServiceType"
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetGongstructName() string {
	return "CodePilotControlledLightingType"
}

func (codervsmtype *CodeRVSMType) GongGetGongstructName() string {
	return "CodeRVSMType"
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetGongstructName() string {
	return "CodeRadioEmissionType"
}

func (codeserviceatctype *CodeServiceATCType) GongGetGongstructName() string {
	return "CodeServiceATCType"
}

func (codeserviceatfmtype *CodeServiceATFMType) GongGetGongstructName() string {
	return "CodeServiceATFMType"
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetGongstructName() string {
	return "CodeServiceGroundControlType"
}

func (codeserviceinformationtype *CodeServiceInformationType) GongGetGongstructName() string {
	return "CodeServiceInformationType"
}

func (codeservicesartype *CodeServiceSARType) GongGetGongstructName() string {
	return "CodeServiceSARType"
}

func (codespecialdatetype *CodeSpecialDateType) GongGetGongstructName() string {
	return "CodeSpecialDateType"
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetGongstructName() string {
	return "CodeStatusOperationsType"
}

func (codestatusservicetype *CodeStatusServiceType) GongGetGongstructName() string {
	return "CodeStatusServiceType"
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetGongstructName() string {
	return "CodeSystemActivationType"
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

func (codeunitdependencytype *CodeUnitDependencyType) GongGetGongstructName() string {
	return "CodeUnitDependencyType"
}

func (codeunittype *CodeUnitType) GongGetGongstructName() string {
	return "CodeUnitType"
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetGongstructName() string {
	return "CodeVerticalDatumType"
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

func (nonumbertype *NoNumberType) GongGetGongstructName() string {
	return "NoNumberType"
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

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.CodeAircraftGroundServiceTypes, &stage.CodeAircraftGroundServiceTypes_mapString)

	__gong__rebuildMapString(stage.CodeAviationStandardsTypes, &stage.CodeAviationStandardsTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationChannelTypes, &stage.CodeCommunicationChannelTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationDirectionTypes, &stage.CodeCommunicationDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeFacilityRankingTypes, &stage.CodeFacilityRankingTypes_mapString)

	__gong__rebuildMapString(stage.CodeFireFightingTypes, &stage.CodeFireFightingTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightDestinationTypes, &stage.CodeFlightDestinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString)

	__gong__rebuildMapString(stage.CodeFuelTypes, &stage.CodeFuelTypes_mapString)

	__gong__rebuildMapString(stage.CodeIntensityStandByTypes, &stage.CodeIntensityStandByTypes_mapString)

	__gong__rebuildMapString(stage.CodeLanguageTypes, &stage.CodeLanguageTypes_mapString)

	__gong__rebuildMapString(stage.CodeLevelSeriesTypes, &stage.CodeLevelSeriesTypes_mapString)

	__gong__rebuildMapString(stage.CodeLevelTableDesignatorTypes, &stage.CodeLevelTableDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeNitrogenTypes, &stage.CodeNitrogenTypes_mapString)

	__gong__rebuildMapString(stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeOilTypes, &stage.CodeOilTypes_mapString)

	__gong__rebuildMapString(stage.CodeOrganisationDesignatorTypes, &stage.CodeOrganisationDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeOrganisationHierarchyTypes, &stage.CodeOrganisationHierarchyTypes_mapString)

	__gong__rebuildMapString(stage.CodeOrganisationTypes, &stage.CodeOrganisationTypes_mapString)

	__gong__rebuildMapString(stage.CodeOxygenTypes, &stage.CodeOxygenTypes_mapString)

	__gong__rebuildMapString(stage.CodePassengerServiceTypes, &stage.CodePassengerServiceTypes_mapString)

	__gong__rebuildMapString(stage.CodePilotControlledLightingTypes, &stage.CodePilotControlledLightingTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString)

	__gong__rebuildMapString(stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_mapString)

	__gong__rebuildMapString(stage.CodeServiceATCTypes, &stage.CodeServiceATCTypes_mapString)

	__gong__rebuildMapString(stage.CodeServiceATFMTypes, &stage.CodeServiceATFMTypes_mapString)

	__gong__rebuildMapString(stage.CodeServiceGroundControlTypes, &stage.CodeServiceGroundControlTypes_mapString)

	__gong__rebuildMapString(stage.CodeServiceInformationTypes, &stage.CodeServiceInformationTypes_mapString)

	__gong__rebuildMapString(stage.CodeServiceSARTypes, &stage.CodeServiceSARTypes_mapString)

	__gong__rebuildMapString(stage.CodeSpecialDateTypes, &stage.CodeSpecialDateTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusServiceTypes, &stage.CodeStatusServiceTypes_mapString)

	__gong__rebuildMapString(stage.CodeSystemActivationTypes, &stage.CodeSystemActivationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeUnitDependencyTypes, &stage.CodeUnitDependencyTypes_mapString)

	__gong__rebuildMapString(stage.CodeUnitTypes, &stage.CodeUnitTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.DateYearTypes, &stage.DateYearTypes_mapString)

	__gong__rebuildMapString(stage.NoNumberTypes, &stage.NoNumberTypes_mapString)

	__gong__rebuildMapString(stage.TextAddressTypes, &stage.TextAddressTypes_mapString)

	__gong__rebuildMapString(stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.TextInstructionTypes, &stage.TextInstructionTypes_mapString)

	__gong__rebuildMapString(stage.TextNameTypes, &stage.TextNameTypes_mapString)

	__gong__rebuildMapString(stage.TextNoteTypes, &stage.TextNoteTypes_mapString)

	__gong__rebuildMapString(stage.TextPhoneTypes, &stage.TextPhoneTypes_mapString)

	__gong__rebuildMapString(stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.ValBearingTypes, &stage.ValBearingTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValFrequencyTypes, &stage.ValFrequencyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
