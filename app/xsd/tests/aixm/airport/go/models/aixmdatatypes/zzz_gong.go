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

	CodeAircraftStandTypes                map[*CodeAircraftStandType]struct{}
	CodeAircraftStandTypes_instance       map[*CodeAircraftStandType]*CodeAircraftStandType
	CodeAircraftStandTypes_mapString      map[string]*CodeAircraftStandType
	CodeAircraftStandTypeOrder            uint
	CodeAircraftStandType_stagedOrder     map[*CodeAircraftStandType]uint
	CodeAircraftStandType_orderStaged     map[uint]*CodeAircraftStandType
	CodeAircraftStandTypes_reference      map[*CodeAircraftStandType]*CodeAircraftStandType
	CodeAircraftStandTypes_referenceOrder map[*CodeAircraftStandType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAircraftStandTypeCreateCallback GongOnAfterCreateInterface[CodeAircraftStandType]
	OnAfterCodeAircraftStandTypeUpdateCallback GongOnAfterUpdateInterface[CodeAircraftStandType]
	OnAfterCodeAircraftStandTypeDeleteCallback GongOnAfterDeleteInterface[CodeAircraftStandType]

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

	CodeAirportHeliportCollocationTypes                map[*CodeAirportHeliportCollocationType]struct{}
	CodeAirportHeliportCollocationTypes_instance       map[*CodeAirportHeliportCollocationType]*CodeAirportHeliportCollocationType
	CodeAirportHeliportCollocationTypes_mapString      map[string]*CodeAirportHeliportCollocationType
	CodeAirportHeliportCollocationTypeOrder            uint
	CodeAirportHeliportCollocationType_stagedOrder     map[*CodeAirportHeliportCollocationType]uint
	CodeAirportHeliportCollocationType_orderStaged     map[uint]*CodeAirportHeliportCollocationType
	CodeAirportHeliportCollocationTypes_reference      map[*CodeAirportHeliportCollocationType]*CodeAirportHeliportCollocationType
	CodeAirportHeliportCollocationTypes_referenceOrder map[*CodeAirportHeliportCollocationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirportHeliportCollocationTypeCreateCallback GongOnAfterCreateInterface[CodeAirportHeliportCollocationType]
	OnAfterCodeAirportHeliportCollocationTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirportHeliportCollocationType]
	OnAfterCodeAirportHeliportCollocationTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirportHeliportCollocationType]

	CodeAirportHeliportDesignatorTypes                map[*CodeAirportHeliportDesignatorType]struct{}
	CodeAirportHeliportDesignatorTypes_instance       map[*CodeAirportHeliportDesignatorType]*CodeAirportHeliportDesignatorType
	CodeAirportHeliportDesignatorTypes_mapString      map[string]*CodeAirportHeliportDesignatorType
	CodeAirportHeliportDesignatorTypeOrder            uint
	CodeAirportHeliportDesignatorType_stagedOrder     map[*CodeAirportHeliportDesignatorType]uint
	CodeAirportHeliportDesignatorType_orderStaged     map[uint]*CodeAirportHeliportDesignatorType
	CodeAirportHeliportDesignatorTypes_reference      map[*CodeAirportHeliportDesignatorType]*CodeAirportHeliportDesignatorType
	CodeAirportHeliportDesignatorTypes_referenceOrder map[*CodeAirportHeliportDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirportHeliportDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeAirportHeliportDesignatorType]
	OnAfterCodeAirportHeliportDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirportHeliportDesignatorType]
	OnAfterCodeAirportHeliportDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirportHeliportDesignatorType]

	CodeAirportHeliportTypes                map[*CodeAirportHeliportType]struct{}
	CodeAirportHeliportTypes_instance       map[*CodeAirportHeliportType]*CodeAirportHeliportType
	CodeAirportHeliportTypes_mapString      map[string]*CodeAirportHeliportType
	CodeAirportHeliportTypeOrder            uint
	CodeAirportHeliportType_stagedOrder     map[*CodeAirportHeliportType]uint
	CodeAirportHeliportType_orderStaged     map[uint]*CodeAirportHeliportType
	CodeAirportHeliportTypes_reference      map[*CodeAirportHeliportType]*CodeAirportHeliportType
	CodeAirportHeliportTypes_referenceOrder map[*CodeAirportHeliportType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirportHeliportTypeCreateCallback GongOnAfterCreateInterface[CodeAirportHeliportType]
	OnAfterCodeAirportHeliportTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirportHeliportType]
	OnAfterCodeAirportHeliportTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirportHeliportType]

	CodeAirportWarningTypes                map[*CodeAirportWarningType]struct{}
	CodeAirportWarningTypes_instance       map[*CodeAirportWarningType]*CodeAirportWarningType
	CodeAirportWarningTypes_mapString      map[string]*CodeAirportWarningType
	CodeAirportWarningTypeOrder            uint
	CodeAirportWarningType_stagedOrder     map[*CodeAirportWarningType]uint
	CodeAirportWarningType_orderStaged     map[uint]*CodeAirportWarningType
	CodeAirportWarningTypes_reference      map[*CodeAirportWarningType]*CodeAirportWarningType
	CodeAirportWarningTypes_referenceOrder map[*CodeAirportWarningType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirportWarningTypeCreateCallback GongOnAfterCreateInterface[CodeAirportWarningType]
	OnAfterCodeAirportWarningTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirportWarningType]
	OnAfterCodeAirportWarningTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirportWarningType]

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

	CodeApproachLightingICAOTypes                map[*CodeApproachLightingICAOType]struct{}
	CodeApproachLightingICAOTypes_instance       map[*CodeApproachLightingICAOType]*CodeApproachLightingICAOType
	CodeApproachLightingICAOTypes_mapString      map[string]*CodeApproachLightingICAOType
	CodeApproachLightingICAOTypeOrder            uint
	CodeApproachLightingICAOType_stagedOrder     map[*CodeApproachLightingICAOType]uint
	CodeApproachLightingICAOType_orderStaged     map[uint]*CodeApproachLightingICAOType
	CodeApproachLightingICAOTypes_reference      map[*CodeApproachLightingICAOType]*CodeApproachLightingICAOType
	CodeApproachLightingICAOTypes_referenceOrder map[*CodeApproachLightingICAOType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachLightingICAOTypeCreateCallback GongOnAfterCreateInterface[CodeApproachLightingICAOType]
	OnAfterCodeApproachLightingICAOTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachLightingICAOType]
	OnAfterCodeApproachLightingICAOTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachLightingICAOType]

	CodeApproachLightingTypes                map[*CodeApproachLightingType]struct{}
	CodeApproachLightingTypes_instance       map[*CodeApproachLightingType]*CodeApproachLightingType
	CodeApproachLightingTypes_mapString      map[string]*CodeApproachLightingType
	CodeApproachLightingTypeOrder            uint
	CodeApproachLightingType_stagedOrder     map[*CodeApproachLightingType]uint
	CodeApproachLightingType_orderStaged     map[uint]*CodeApproachLightingType
	CodeApproachLightingTypes_reference      map[*CodeApproachLightingType]*CodeApproachLightingType
	CodeApproachLightingTypes_referenceOrder map[*CodeApproachLightingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApproachLightingTypeCreateCallback GongOnAfterCreateInterface[CodeApproachLightingType]
	OnAfterCodeApproachLightingTypeUpdateCallback GongOnAfterUpdateInterface[CodeApproachLightingType]
	OnAfterCodeApproachLightingTypeDeleteCallback GongOnAfterDeleteInterface[CodeApproachLightingType]

	CodeApronElementTypes                map[*CodeApronElementType]struct{}
	CodeApronElementTypes_instance       map[*CodeApronElementType]*CodeApronElementType
	CodeApronElementTypes_mapString      map[string]*CodeApronElementType
	CodeApronElementTypeOrder            uint
	CodeApronElementType_stagedOrder     map[*CodeApronElementType]uint
	CodeApronElementType_orderStaged     map[uint]*CodeApronElementType
	CodeApronElementTypes_reference      map[*CodeApronElementType]*CodeApronElementType
	CodeApronElementTypes_referenceOrder map[*CodeApronElementType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApronElementTypeCreateCallback GongOnAfterCreateInterface[CodeApronElementType]
	OnAfterCodeApronElementTypeUpdateCallback GongOnAfterUpdateInterface[CodeApronElementType]
	OnAfterCodeApronElementTypeDeleteCallback GongOnAfterDeleteInterface[CodeApronElementType]

	CodeApronSectionTypes                map[*CodeApronSectionType]struct{}
	CodeApronSectionTypes_instance       map[*CodeApronSectionType]*CodeApronSectionType
	CodeApronSectionTypes_mapString      map[string]*CodeApronSectionType
	CodeApronSectionTypeOrder            uint
	CodeApronSectionType_stagedOrder     map[*CodeApronSectionType]uint
	CodeApronSectionType_orderStaged     map[uint]*CodeApronSectionType
	CodeApronSectionTypes_reference      map[*CodeApronSectionType]*CodeApronSectionType
	CodeApronSectionTypes_referenceOrder map[*CodeApronSectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeApronSectionTypeCreateCallback GongOnAfterCreateInterface[CodeApronSectionType]
	OnAfterCodeApronSectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeApronSectionType]
	OnAfterCodeApronSectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeApronSectionType]

	CodeArrestingGearEnergyAbsorbTypes                map[*CodeArrestingGearEnergyAbsorbType]struct{}
	CodeArrestingGearEnergyAbsorbTypes_instance       map[*CodeArrestingGearEnergyAbsorbType]*CodeArrestingGearEnergyAbsorbType
	CodeArrestingGearEnergyAbsorbTypes_mapString      map[string]*CodeArrestingGearEnergyAbsorbType
	CodeArrestingGearEnergyAbsorbTypeOrder            uint
	CodeArrestingGearEnergyAbsorbType_stagedOrder     map[*CodeArrestingGearEnergyAbsorbType]uint
	CodeArrestingGearEnergyAbsorbType_orderStaged     map[uint]*CodeArrestingGearEnergyAbsorbType
	CodeArrestingGearEnergyAbsorbTypes_reference      map[*CodeArrestingGearEnergyAbsorbType]*CodeArrestingGearEnergyAbsorbType
	CodeArrestingGearEnergyAbsorbTypes_referenceOrder map[*CodeArrestingGearEnergyAbsorbType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeArrestingGearEnergyAbsorbTypeCreateCallback GongOnAfterCreateInterface[CodeArrestingGearEnergyAbsorbType]
	OnAfterCodeArrestingGearEnergyAbsorbTypeUpdateCallback GongOnAfterUpdateInterface[CodeArrestingGearEnergyAbsorbType]
	OnAfterCodeArrestingGearEnergyAbsorbTypeDeleteCallback GongOnAfterDeleteInterface[CodeArrestingGearEnergyAbsorbType]

	CodeArrestingGearEngageDeviceTypes                map[*CodeArrestingGearEngageDeviceType]struct{}
	CodeArrestingGearEngageDeviceTypes_instance       map[*CodeArrestingGearEngageDeviceType]*CodeArrestingGearEngageDeviceType
	CodeArrestingGearEngageDeviceTypes_mapString      map[string]*CodeArrestingGearEngageDeviceType
	CodeArrestingGearEngageDeviceTypeOrder            uint
	CodeArrestingGearEngageDeviceType_stagedOrder     map[*CodeArrestingGearEngageDeviceType]uint
	CodeArrestingGearEngageDeviceType_orderStaged     map[uint]*CodeArrestingGearEngageDeviceType
	CodeArrestingGearEngageDeviceTypes_reference      map[*CodeArrestingGearEngageDeviceType]*CodeArrestingGearEngageDeviceType
	CodeArrestingGearEngageDeviceTypes_referenceOrder map[*CodeArrestingGearEngageDeviceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeArrestingGearEngageDeviceTypeCreateCallback GongOnAfterCreateInterface[CodeArrestingGearEngageDeviceType]
	OnAfterCodeArrestingGearEngageDeviceTypeUpdateCallback GongOnAfterUpdateInterface[CodeArrestingGearEngageDeviceType]
	OnAfterCodeArrestingGearEngageDeviceTypeDeleteCallback GongOnAfterDeleteInterface[CodeArrestingGearEngageDeviceType]

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

	CodeBuoyDesignatorTypes                map[*CodeBuoyDesignatorType]struct{}
	CodeBuoyDesignatorTypes_instance       map[*CodeBuoyDesignatorType]*CodeBuoyDesignatorType
	CodeBuoyDesignatorTypes_mapString      map[string]*CodeBuoyDesignatorType
	CodeBuoyDesignatorTypeOrder            uint
	CodeBuoyDesignatorType_stagedOrder     map[*CodeBuoyDesignatorType]uint
	CodeBuoyDesignatorType_orderStaged     map[uint]*CodeBuoyDesignatorType
	CodeBuoyDesignatorTypes_reference      map[*CodeBuoyDesignatorType]*CodeBuoyDesignatorType
	CodeBuoyDesignatorTypes_referenceOrder map[*CodeBuoyDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeBuoyDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeBuoyDesignatorType]
	OnAfterCodeBuoyDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeBuoyDesignatorType]
	OnAfterCodeBuoyDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeBuoyDesignatorType]

	CodeBuoyTypes                map[*CodeBuoyType]struct{}
	CodeBuoyTypes_instance       map[*CodeBuoyType]*CodeBuoyType
	CodeBuoyTypes_mapString      map[string]*CodeBuoyType
	CodeBuoyTypeOrder            uint
	CodeBuoyType_stagedOrder     map[*CodeBuoyType]uint
	CodeBuoyType_orderStaged     map[uint]*CodeBuoyType
	CodeBuoyTypes_reference      map[*CodeBuoyType]*CodeBuoyType
	CodeBuoyTypes_referenceOrder map[*CodeBuoyType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeBuoyTypeCreateCallback GongOnAfterCreateInterface[CodeBuoyType]
	OnAfterCodeBuoyTypeUpdateCallback GongOnAfterUpdateInterface[CodeBuoyType]
	OnAfterCodeBuoyTypeDeleteCallback GongOnAfterDeleteInterface[CodeBuoyType]

	CodeColourTypes                map[*CodeColourType]struct{}
	CodeColourTypes_instance       map[*CodeColourType]*CodeColourType
	CodeColourTypes_mapString      map[string]*CodeColourType
	CodeColourTypeOrder            uint
	CodeColourType_stagedOrder     map[*CodeColourType]uint
	CodeColourType_orderStaged     map[uint]*CodeColourType
	CodeColourTypes_reference      map[*CodeColourType]*CodeColourType
	CodeColourTypes_referenceOrder map[*CodeColourType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeColourTypeCreateCallback GongOnAfterCreateInterface[CodeColourType]
	OnAfterCodeColourTypeUpdateCallback GongOnAfterUpdateInterface[CodeColourType]
	OnAfterCodeColourTypeDeleteCallback GongOnAfterDeleteInterface[CodeColourType]

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

	CodeContaminationTypes                map[*CodeContaminationType]struct{}
	CodeContaminationTypes_instance       map[*CodeContaminationType]*CodeContaminationType
	CodeContaminationTypes_mapString      map[string]*CodeContaminationType
	CodeContaminationTypeOrder            uint
	CodeContaminationType_stagedOrder     map[*CodeContaminationType]uint
	CodeContaminationType_orderStaged     map[uint]*CodeContaminationType
	CodeContaminationTypes_reference      map[*CodeContaminationType]*CodeContaminationType
	CodeContaminationTypes_referenceOrder map[*CodeContaminationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeContaminationTypeCreateCallback GongOnAfterCreateInterface[CodeContaminationType]
	OnAfterCodeContaminationTypeUpdateCallback GongOnAfterUpdateInterface[CodeContaminationType]
	OnAfterCodeContaminationTypeDeleteCallback GongOnAfterDeleteInterface[CodeContaminationType]

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

	CodeDeclaredDistanceTypes                map[*CodeDeclaredDistanceType]struct{}
	CodeDeclaredDistanceTypes_instance       map[*CodeDeclaredDistanceType]*CodeDeclaredDistanceType
	CodeDeclaredDistanceTypes_mapString      map[string]*CodeDeclaredDistanceType
	CodeDeclaredDistanceTypeOrder            uint
	CodeDeclaredDistanceType_stagedOrder     map[*CodeDeclaredDistanceType]uint
	CodeDeclaredDistanceType_orderStaged     map[uint]*CodeDeclaredDistanceType
	CodeDeclaredDistanceTypes_reference      map[*CodeDeclaredDistanceType]*CodeDeclaredDistanceType
	CodeDeclaredDistanceTypes_referenceOrder map[*CodeDeclaredDistanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeDeclaredDistanceTypeCreateCallback GongOnAfterCreateInterface[CodeDeclaredDistanceType]
	OnAfterCodeDeclaredDistanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeDeclaredDistanceType]
	OnAfterCodeDeclaredDistanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeDeclaredDistanceType]

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

	CodeFrictionDeviceTypes                map[*CodeFrictionDeviceType]struct{}
	CodeFrictionDeviceTypes_instance       map[*CodeFrictionDeviceType]*CodeFrictionDeviceType
	CodeFrictionDeviceTypes_mapString      map[string]*CodeFrictionDeviceType
	CodeFrictionDeviceTypeOrder            uint
	CodeFrictionDeviceType_stagedOrder     map[*CodeFrictionDeviceType]uint
	CodeFrictionDeviceType_orderStaged     map[uint]*CodeFrictionDeviceType
	CodeFrictionDeviceTypes_reference      map[*CodeFrictionDeviceType]*CodeFrictionDeviceType
	CodeFrictionDeviceTypes_referenceOrder map[*CodeFrictionDeviceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFrictionDeviceTypeCreateCallback GongOnAfterCreateInterface[CodeFrictionDeviceType]
	OnAfterCodeFrictionDeviceTypeUpdateCallback GongOnAfterUpdateInterface[CodeFrictionDeviceType]
	OnAfterCodeFrictionDeviceTypeDeleteCallback GongOnAfterDeleteInterface[CodeFrictionDeviceType]

	CodeFrictionEstimateTypes                map[*CodeFrictionEstimateType]struct{}
	CodeFrictionEstimateTypes_instance       map[*CodeFrictionEstimateType]*CodeFrictionEstimateType
	CodeFrictionEstimateTypes_mapString      map[string]*CodeFrictionEstimateType
	CodeFrictionEstimateTypeOrder            uint
	CodeFrictionEstimateType_stagedOrder     map[*CodeFrictionEstimateType]uint
	CodeFrictionEstimateType_orderStaged     map[uint]*CodeFrictionEstimateType
	CodeFrictionEstimateTypes_reference      map[*CodeFrictionEstimateType]*CodeFrictionEstimateType
	CodeFrictionEstimateTypes_referenceOrder map[*CodeFrictionEstimateType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeFrictionEstimateTypeCreateCallback GongOnAfterCreateInterface[CodeFrictionEstimateType]
	OnAfterCodeFrictionEstimateTypeUpdateCallback GongOnAfterUpdateInterface[CodeFrictionEstimateType]
	OnAfterCodeFrictionEstimateTypeDeleteCallback GongOnAfterDeleteInterface[CodeFrictionEstimateType]

	CodeGradeSeparationTypes                map[*CodeGradeSeparationType]struct{}
	CodeGradeSeparationTypes_instance       map[*CodeGradeSeparationType]*CodeGradeSeparationType
	CodeGradeSeparationTypes_mapString      map[string]*CodeGradeSeparationType
	CodeGradeSeparationTypeOrder            uint
	CodeGradeSeparationType_stagedOrder     map[*CodeGradeSeparationType]uint
	CodeGradeSeparationType_orderStaged     map[uint]*CodeGradeSeparationType
	CodeGradeSeparationTypes_reference      map[*CodeGradeSeparationType]*CodeGradeSeparationType
	CodeGradeSeparationTypes_referenceOrder map[*CodeGradeSeparationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeGradeSeparationTypeCreateCallback GongOnAfterCreateInterface[CodeGradeSeparationType]
	OnAfterCodeGradeSeparationTypeUpdateCallback GongOnAfterUpdateInterface[CodeGradeSeparationType]
	OnAfterCodeGradeSeparationTypeDeleteCallback GongOnAfterDeleteInterface[CodeGradeSeparationType]

	CodeGuidanceLineTypes                map[*CodeGuidanceLineType]struct{}
	CodeGuidanceLineTypes_instance       map[*CodeGuidanceLineType]*CodeGuidanceLineType
	CodeGuidanceLineTypes_mapString      map[string]*CodeGuidanceLineType
	CodeGuidanceLineTypeOrder            uint
	CodeGuidanceLineType_stagedOrder     map[*CodeGuidanceLineType]uint
	CodeGuidanceLineType_orderStaged     map[uint]*CodeGuidanceLineType
	CodeGuidanceLineTypes_reference      map[*CodeGuidanceLineType]*CodeGuidanceLineType
	CodeGuidanceLineTypes_referenceOrder map[*CodeGuidanceLineType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeGuidanceLineTypeCreateCallback GongOnAfterCreateInterface[CodeGuidanceLineType]
	OnAfterCodeGuidanceLineTypeUpdateCallback GongOnAfterUpdateInterface[CodeGuidanceLineType]
	OnAfterCodeGuidanceLineTypeDeleteCallback GongOnAfterDeleteInterface[CodeGuidanceLineType]

	CodeHelicopterPerformanceTypes                map[*CodeHelicopterPerformanceType]struct{}
	CodeHelicopterPerformanceTypes_instance       map[*CodeHelicopterPerformanceType]*CodeHelicopterPerformanceType
	CodeHelicopterPerformanceTypes_mapString      map[string]*CodeHelicopterPerformanceType
	CodeHelicopterPerformanceTypeOrder            uint
	CodeHelicopterPerformanceType_stagedOrder     map[*CodeHelicopterPerformanceType]uint
	CodeHelicopterPerformanceType_orderStaged     map[uint]*CodeHelicopterPerformanceType
	CodeHelicopterPerformanceTypes_reference      map[*CodeHelicopterPerformanceType]*CodeHelicopterPerformanceType
	CodeHelicopterPerformanceTypes_referenceOrder map[*CodeHelicopterPerformanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeHelicopterPerformanceTypeCreateCallback GongOnAfterCreateInterface[CodeHelicopterPerformanceType]
	OnAfterCodeHelicopterPerformanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeHelicopterPerformanceType]
	OnAfterCodeHelicopterPerformanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeHelicopterPerformanceType]

	CodeHoldingCategoryTypes                map[*CodeHoldingCategoryType]struct{}
	CodeHoldingCategoryTypes_instance       map[*CodeHoldingCategoryType]*CodeHoldingCategoryType
	CodeHoldingCategoryTypes_mapString      map[string]*CodeHoldingCategoryType
	CodeHoldingCategoryTypeOrder            uint
	CodeHoldingCategoryType_stagedOrder     map[*CodeHoldingCategoryType]uint
	CodeHoldingCategoryType_orderStaged     map[uint]*CodeHoldingCategoryType
	CodeHoldingCategoryTypes_reference      map[*CodeHoldingCategoryType]*CodeHoldingCategoryType
	CodeHoldingCategoryTypes_referenceOrder map[*CodeHoldingCategoryType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeHoldingCategoryTypeCreateCallback GongOnAfterCreateInterface[CodeHoldingCategoryType]
	OnAfterCodeHoldingCategoryTypeUpdateCallback GongOnAfterUpdateInterface[CodeHoldingCategoryType]
	OnAfterCodeHoldingCategoryTypeDeleteCallback GongOnAfterDeleteInterface[CodeHoldingCategoryType]

	CodeIATATypes                map[*CodeIATAType]struct{}
	CodeIATATypes_instance       map[*CodeIATAType]*CodeIATAType
	CodeIATATypes_mapString      map[string]*CodeIATAType
	CodeIATATypeOrder            uint
	CodeIATAType_stagedOrder     map[*CodeIATAType]uint
	CodeIATAType_orderStaged     map[uint]*CodeIATAType
	CodeIATATypes_reference      map[*CodeIATAType]*CodeIATAType
	CodeIATATypes_referenceOrder map[*CodeIATAType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeIATATypeCreateCallback GongOnAfterCreateInterface[CodeIATAType]
	OnAfterCodeIATATypeUpdateCallback GongOnAfterUpdateInterface[CodeIATAType]
	OnAfterCodeIATATypeDeleteCallback GongOnAfterDeleteInterface[CodeIATAType]

	CodeICAOTypes                map[*CodeICAOType]struct{}
	CodeICAOTypes_instance       map[*CodeICAOType]*CodeICAOType
	CodeICAOTypes_mapString      map[string]*CodeICAOType
	CodeICAOTypeOrder            uint
	CodeICAOType_stagedOrder     map[*CodeICAOType]uint
	CodeICAOType_orderStaged     map[uint]*CodeICAOType
	CodeICAOTypes_reference      map[*CodeICAOType]*CodeICAOType
	CodeICAOTypes_referenceOrder map[*CodeICAOType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeICAOTypeCreateCallback GongOnAfterCreateInterface[CodeICAOType]
	OnAfterCodeICAOTypeUpdateCallback GongOnAfterUpdateInterface[CodeICAOType]
	OnAfterCodeICAOTypeDeleteCallback GongOnAfterDeleteInterface[CodeICAOType]

	CodeLightHoldingPositionTypes                map[*CodeLightHoldingPositionType]struct{}
	CodeLightHoldingPositionTypes_instance       map[*CodeLightHoldingPositionType]*CodeLightHoldingPositionType
	CodeLightHoldingPositionTypes_mapString      map[string]*CodeLightHoldingPositionType
	CodeLightHoldingPositionTypeOrder            uint
	CodeLightHoldingPositionType_stagedOrder     map[*CodeLightHoldingPositionType]uint
	CodeLightHoldingPositionType_orderStaged     map[uint]*CodeLightHoldingPositionType
	CodeLightHoldingPositionTypes_reference      map[*CodeLightHoldingPositionType]*CodeLightHoldingPositionType
	CodeLightHoldingPositionTypes_referenceOrder map[*CodeLightHoldingPositionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLightHoldingPositionTypeCreateCallback GongOnAfterCreateInterface[CodeLightHoldingPositionType]
	OnAfterCodeLightHoldingPositionTypeUpdateCallback GongOnAfterUpdateInterface[CodeLightHoldingPositionType]
	OnAfterCodeLightHoldingPositionTypeDeleteCallback GongOnAfterDeleteInterface[CodeLightHoldingPositionType]

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

	CodeLightSourceTypes                map[*CodeLightSourceType]struct{}
	CodeLightSourceTypes_instance       map[*CodeLightSourceType]*CodeLightSourceType
	CodeLightSourceTypes_mapString      map[string]*CodeLightSourceType
	CodeLightSourceTypeOrder            uint
	CodeLightSourceType_stagedOrder     map[*CodeLightSourceType]uint
	CodeLightSourceType_orderStaged     map[uint]*CodeLightSourceType
	CodeLightSourceTypes_reference      map[*CodeLightSourceType]*CodeLightSourceType
	CodeLightSourceTypes_referenceOrder map[*CodeLightSourceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLightSourceTypeCreateCallback GongOnAfterCreateInterface[CodeLightSourceType]
	OnAfterCodeLightSourceTypeUpdateCallback GongOnAfterUpdateInterface[CodeLightSourceType]
	OnAfterCodeLightSourceTypeDeleteCallback GongOnAfterDeleteInterface[CodeLightSourceType]

	CodeLightingJARTypes                map[*CodeLightingJARType]struct{}
	CodeLightingJARTypes_instance       map[*CodeLightingJARType]*CodeLightingJARType
	CodeLightingJARTypes_mapString      map[string]*CodeLightingJARType
	CodeLightingJARTypeOrder            uint
	CodeLightingJARType_stagedOrder     map[*CodeLightingJARType]uint
	CodeLightingJARType_orderStaged     map[uint]*CodeLightingJARType
	CodeLightingJARTypes_reference      map[*CodeLightingJARType]*CodeLightingJARType
	CodeLightingJARTypes_referenceOrder map[*CodeLightingJARType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLightingJARTypeCreateCallback GongOnAfterCreateInterface[CodeLightingJARType]
	OnAfterCodeLightingJARTypeUpdateCallback GongOnAfterUpdateInterface[CodeLightingJARType]
	OnAfterCodeLightingJARTypeDeleteCallback GongOnAfterDeleteInterface[CodeLightingJARType]

	CodeLoadingBridgeTypes                map[*CodeLoadingBridgeType]struct{}
	CodeLoadingBridgeTypes_instance       map[*CodeLoadingBridgeType]*CodeLoadingBridgeType
	CodeLoadingBridgeTypes_mapString      map[string]*CodeLoadingBridgeType
	CodeLoadingBridgeTypeOrder            uint
	CodeLoadingBridgeType_stagedOrder     map[*CodeLoadingBridgeType]uint
	CodeLoadingBridgeType_orderStaged     map[uint]*CodeLoadingBridgeType
	CodeLoadingBridgeTypes_reference      map[*CodeLoadingBridgeType]*CodeLoadingBridgeType
	CodeLoadingBridgeTypes_referenceOrder map[*CodeLoadingBridgeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLoadingBridgeTypeCreateCallback GongOnAfterCreateInterface[CodeLoadingBridgeType]
	OnAfterCodeLoadingBridgeTypeUpdateCallback GongOnAfterUpdateInterface[CodeLoadingBridgeType]
	OnAfterCodeLoadingBridgeTypeDeleteCallback GongOnAfterDeleteInterface[CodeLoadingBridgeType]

	CodeLogicalOperatorTypes                map[*CodeLogicalOperatorType]struct{}
	CodeLogicalOperatorTypes_instance       map[*CodeLogicalOperatorType]*CodeLogicalOperatorType
	CodeLogicalOperatorTypes_mapString      map[string]*CodeLogicalOperatorType
	CodeLogicalOperatorTypeOrder            uint
	CodeLogicalOperatorType_stagedOrder     map[*CodeLogicalOperatorType]uint
	CodeLogicalOperatorType_orderStaged     map[uint]*CodeLogicalOperatorType
	CodeLogicalOperatorTypes_reference      map[*CodeLogicalOperatorType]*CodeLogicalOperatorType
	CodeLogicalOperatorTypes_referenceOrder map[*CodeLogicalOperatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeLogicalOperatorTypeCreateCallback GongOnAfterCreateInterface[CodeLogicalOperatorType]
	OnAfterCodeLogicalOperatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeLogicalOperatorType]
	OnAfterCodeLogicalOperatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeLogicalOperatorType]

	CodeMarkingConditionTypes                map[*CodeMarkingConditionType]struct{}
	CodeMarkingConditionTypes_instance       map[*CodeMarkingConditionType]*CodeMarkingConditionType
	CodeMarkingConditionTypes_mapString      map[string]*CodeMarkingConditionType
	CodeMarkingConditionTypeOrder            uint
	CodeMarkingConditionType_stagedOrder     map[*CodeMarkingConditionType]uint
	CodeMarkingConditionType_orderStaged     map[uint]*CodeMarkingConditionType
	CodeMarkingConditionTypes_reference      map[*CodeMarkingConditionType]*CodeMarkingConditionType
	CodeMarkingConditionTypes_referenceOrder map[*CodeMarkingConditionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMarkingConditionTypeCreateCallback GongOnAfterCreateInterface[CodeMarkingConditionType]
	OnAfterCodeMarkingConditionTypeUpdateCallback GongOnAfterUpdateInterface[CodeMarkingConditionType]
	OnAfterCodeMarkingConditionTypeDeleteCallback GongOnAfterDeleteInterface[CodeMarkingConditionType]

	CodeMarkingStyleTypes                map[*CodeMarkingStyleType]struct{}
	CodeMarkingStyleTypes_instance       map[*CodeMarkingStyleType]*CodeMarkingStyleType
	CodeMarkingStyleTypes_mapString      map[string]*CodeMarkingStyleType
	CodeMarkingStyleTypeOrder            uint
	CodeMarkingStyleType_stagedOrder     map[*CodeMarkingStyleType]uint
	CodeMarkingStyleType_orderStaged     map[uint]*CodeMarkingStyleType
	CodeMarkingStyleTypes_reference      map[*CodeMarkingStyleType]*CodeMarkingStyleType
	CodeMarkingStyleTypes_referenceOrder map[*CodeMarkingStyleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeMarkingStyleTypeCreateCallback GongOnAfterCreateInterface[CodeMarkingStyleType]
	OnAfterCodeMarkingStyleTypeUpdateCallback GongOnAfterUpdateInterface[CodeMarkingStyleType]
	OnAfterCodeMarkingStyleTypeDeleteCallback GongOnAfterDeleteInterface[CodeMarkingStyleType]

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

	CodeOperationAirportHeliportTypes                map[*CodeOperationAirportHeliportType]struct{}
	CodeOperationAirportHeliportTypes_instance       map[*CodeOperationAirportHeliportType]*CodeOperationAirportHeliportType
	CodeOperationAirportHeliportTypes_mapString      map[string]*CodeOperationAirportHeliportType
	CodeOperationAirportHeliportTypeOrder            uint
	CodeOperationAirportHeliportType_stagedOrder     map[*CodeOperationAirportHeliportType]uint
	CodeOperationAirportHeliportType_orderStaged     map[uint]*CodeOperationAirportHeliportType
	CodeOperationAirportHeliportTypes_reference      map[*CodeOperationAirportHeliportType]*CodeOperationAirportHeliportType
	CodeOperationAirportHeliportTypes_referenceOrder map[*CodeOperationAirportHeliportType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOperationAirportHeliportTypeCreateCallback GongOnAfterCreateInterface[CodeOperationAirportHeliportType]
	OnAfterCodeOperationAirportHeliportTypeUpdateCallback GongOnAfterUpdateInterface[CodeOperationAirportHeliportType]
	OnAfterCodeOperationAirportHeliportTypeDeleteCallback GongOnAfterDeleteInterface[CodeOperationAirportHeliportType]

	CodeOperationManoeuvringAreaTypes                map[*CodeOperationManoeuvringAreaType]struct{}
	CodeOperationManoeuvringAreaTypes_instance       map[*CodeOperationManoeuvringAreaType]*CodeOperationManoeuvringAreaType
	CodeOperationManoeuvringAreaTypes_mapString      map[string]*CodeOperationManoeuvringAreaType
	CodeOperationManoeuvringAreaTypeOrder            uint
	CodeOperationManoeuvringAreaType_stagedOrder     map[*CodeOperationManoeuvringAreaType]uint
	CodeOperationManoeuvringAreaType_orderStaged     map[uint]*CodeOperationManoeuvringAreaType
	CodeOperationManoeuvringAreaTypes_reference      map[*CodeOperationManoeuvringAreaType]*CodeOperationManoeuvringAreaType
	CodeOperationManoeuvringAreaTypes_referenceOrder map[*CodeOperationManoeuvringAreaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeOperationManoeuvringAreaTypeCreateCallback GongOnAfterCreateInterface[CodeOperationManoeuvringAreaType]
	OnAfterCodeOperationManoeuvringAreaTypeUpdateCallback GongOnAfterUpdateInterface[CodeOperationManoeuvringAreaType]
	OnAfterCodeOperationManoeuvringAreaTypeDeleteCallback GongOnAfterDeleteInterface[CodeOperationManoeuvringAreaType]

	CodePCNMethodTypes                map[*CodePCNMethodType]struct{}
	CodePCNMethodTypes_instance       map[*CodePCNMethodType]*CodePCNMethodType
	CodePCNMethodTypes_mapString      map[string]*CodePCNMethodType
	CodePCNMethodTypeOrder            uint
	CodePCNMethodType_stagedOrder     map[*CodePCNMethodType]uint
	CodePCNMethodType_orderStaged     map[uint]*CodePCNMethodType
	CodePCNMethodTypes_reference      map[*CodePCNMethodType]*CodePCNMethodType
	CodePCNMethodTypes_referenceOrder map[*CodePCNMethodType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePCNMethodTypeCreateCallback GongOnAfterCreateInterface[CodePCNMethodType]
	OnAfterCodePCNMethodTypeUpdateCallback GongOnAfterUpdateInterface[CodePCNMethodType]
	OnAfterCodePCNMethodTypeDeleteCallback GongOnAfterDeleteInterface[CodePCNMethodType]

	CodePCNPavementTypes                map[*CodePCNPavementType]struct{}
	CodePCNPavementTypes_instance       map[*CodePCNPavementType]*CodePCNPavementType
	CodePCNPavementTypes_mapString      map[string]*CodePCNPavementType
	CodePCNPavementTypeOrder            uint
	CodePCNPavementType_stagedOrder     map[*CodePCNPavementType]uint
	CodePCNPavementType_orderStaged     map[uint]*CodePCNPavementType
	CodePCNPavementTypes_reference      map[*CodePCNPavementType]*CodePCNPavementType
	CodePCNPavementTypes_referenceOrder map[*CodePCNPavementType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePCNPavementTypeCreateCallback GongOnAfterCreateInterface[CodePCNPavementType]
	OnAfterCodePCNPavementTypeUpdateCallback GongOnAfterUpdateInterface[CodePCNPavementType]
	OnAfterCodePCNPavementTypeDeleteCallback GongOnAfterDeleteInterface[CodePCNPavementType]

	CodePCNSubgradeTypes                map[*CodePCNSubgradeType]struct{}
	CodePCNSubgradeTypes_instance       map[*CodePCNSubgradeType]*CodePCNSubgradeType
	CodePCNSubgradeTypes_mapString      map[string]*CodePCNSubgradeType
	CodePCNSubgradeTypeOrder            uint
	CodePCNSubgradeType_stagedOrder     map[*CodePCNSubgradeType]uint
	CodePCNSubgradeType_orderStaged     map[uint]*CodePCNSubgradeType
	CodePCNSubgradeTypes_reference      map[*CodePCNSubgradeType]*CodePCNSubgradeType
	CodePCNSubgradeTypes_referenceOrder map[*CodePCNSubgradeType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePCNSubgradeTypeCreateCallback GongOnAfterCreateInterface[CodePCNSubgradeType]
	OnAfterCodePCNSubgradeTypeUpdateCallback GongOnAfterUpdateInterface[CodePCNSubgradeType]
	OnAfterCodePCNSubgradeTypeDeleteCallback GongOnAfterDeleteInterface[CodePCNSubgradeType]

	CodePCNTyrePressureTypes                map[*CodePCNTyrePressureType]struct{}
	CodePCNTyrePressureTypes_instance       map[*CodePCNTyrePressureType]*CodePCNTyrePressureType
	CodePCNTyrePressureTypes_mapString      map[string]*CodePCNTyrePressureType
	CodePCNTyrePressureTypeOrder            uint
	CodePCNTyrePressureType_stagedOrder     map[*CodePCNTyrePressureType]uint
	CodePCNTyrePressureType_orderStaged     map[uint]*CodePCNTyrePressureType
	CodePCNTyrePressureTypes_reference      map[*CodePCNTyrePressureType]*CodePCNTyrePressureType
	CodePCNTyrePressureTypes_referenceOrder map[*CodePCNTyrePressureType]uint

	// insertion point for slice of pointers maps
	OnAfterCodePCNTyrePressureTypeCreateCallback GongOnAfterCreateInterface[CodePCNTyrePressureType]
	OnAfterCodePCNTyrePressureTypeUpdateCallback GongOnAfterUpdateInterface[CodePCNTyrePressureType]
	OnAfterCodePCNTyrePressureTypeDeleteCallback GongOnAfterDeleteInterface[CodePCNTyrePressureType]

	CodeProtectAreaSectionTypes                map[*CodeProtectAreaSectionType]struct{}
	CodeProtectAreaSectionTypes_instance       map[*CodeProtectAreaSectionType]*CodeProtectAreaSectionType
	CodeProtectAreaSectionTypes_mapString      map[string]*CodeProtectAreaSectionType
	CodeProtectAreaSectionTypeOrder            uint
	CodeProtectAreaSectionType_stagedOrder     map[*CodeProtectAreaSectionType]uint
	CodeProtectAreaSectionType_orderStaged     map[uint]*CodeProtectAreaSectionType
	CodeProtectAreaSectionTypes_reference      map[*CodeProtectAreaSectionType]*CodeProtectAreaSectionType
	CodeProtectAreaSectionTypes_referenceOrder map[*CodeProtectAreaSectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeProtectAreaSectionTypeCreateCallback GongOnAfterCreateInterface[CodeProtectAreaSectionType]
	OnAfterCodeProtectAreaSectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeProtectAreaSectionType]
	OnAfterCodeProtectAreaSectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeProtectAreaSectionType]

	CodeRVRReadingTypes                map[*CodeRVRReadingType]struct{}
	CodeRVRReadingTypes_instance       map[*CodeRVRReadingType]*CodeRVRReadingType
	CodeRVRReadingTypes_mapString      map[string]*CodeRVRReadingType
	CodeRVRReadingTypeOrder            uint
	CodeRVRReadingType_stagedOrder     map[*CodeRVRReadingType]uint
	CodeRVRReadingType_orderStaged     map[uint]*CodeRVRReadingType
	CodeRVRReadingTypes_reference      map[*CodeRVRReadingType]*CodeRVRReadingType
	CodeRVRReadingTypes_referenceOrder map[*CodeRVRReadingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRVRReadingTypeCreateCallback GongOnAfterCreateInterface[CodeRVRReadingType]
	OnAfterCodeRVRReadingTypeUpdateCallback GongOnAfterUpdateInterface[CodeRVRReadingType]
	OnAfterCodeRVRReadingTypeDeleteCallback GongOnAfterDeleteInterface[CodeRVRReadingType]

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

	CodeRoadTypes                map[*CodeRoadType]struct{}
	CodeRoadTypes_instance       map[*CodeRoadType]*CodeRoadType
	CodeRoadTypes_mapString      map[string]*CodeRoadType
	CodeRoadTypeOrder            uint
	CodeRoadType_stagedOrder     map[*CodeRoadType]uint
	CodeRoadType_orderStaged     map[uint]*CodeRoadType
	CodeRoadTypes_reference      map[*CodeRoadType]*CodeRoadType
	CodeRoadTypes_referenceOrder map[*CodeRoadType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRoadTypeCreateCallback GongOnAfterCreateInterface[CodeRoadType]
	OnAfterCodeRoadTypeUpdateCallback GongOnAfterUpdateInterface[CodeRoadType]
	OnAfterCodeRoadTypeDeleteCallback GongOnAfterDeleteInterface[CodeRoadType]

	CodeRunwayElementTypes                map[*CodeRunwayElementType]struct{}
	CodeRunwayElementTypes_instance       map[*CodeRunwayElementType]*CodeRunwayElementType
	CodeRunwayElementTypes_mapString      map[string]*CodeRunwayElementType
	CodeRunwayElementTypeOrder            uint
	CodeRunwayElementType_stagedOrder     map[*CodeRunwayElementType]uint
	CodeRunwayElementType_orderStaged     map[uint]*CodeRunwayElementType
	CodeRunwayElementTypes_reference      map[*CodeRunwayElementType]*CodeRunwayElementType
	CodeRunwayElementTypes_referenceOrder map[*CodeRunwayElementType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwayElementTypeCreateCallback GongOnAfterCreateInterface[CodeRunwayElementType]
	OnAfterCodeRunwayElementTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwayElementType]
	OnAfterCodeRunwayElementTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwayElementType]

	CodeRunwayMarkingTypes                map[*CodeRunwayMarkingType]struct{}
	CodeRunwayMarkingTypes_instance       map[*CodeRunwayMarkingType]*CodeRunwayMarkingType
	CodeRunwayMarkingTypes_mapString      map[string]*CodeRunwayMarkingType
	CodeRunwayMarkingTypeOrder            uint
	CodeRunwayMarkingType_stagedOrder     map[*CodeRunwayMarkingType]uint
	CodeRunwayMarkingType_orderStaged     map[uint]*CodeRunwayMarkingType
	CodeRunwayMarkingTypes_reference      map[*CodeRunwayMarkingType]*CodeRunwayMarkingType
	CodeRunwayMarkingTypes_referenceOrder map[*CodeRunwayMarkingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwayMarkingTypeCreateCallback GongOnAfterCreateInterface[CodeRunwayMarkingType]
	OnAfterCodeRunwayMarkingTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwayMarkingType]
	OnAfterCodeRunwayMarkingTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwayMarkingType]

	CodeRunwayPointRoleTypes                map[*CodeRunwayPointRoleType]struct{}
	CodeRunwayPointRoleTypes_instance       map[*CodeRunwayPointRoleType]*CodeRunwayPointRoleType
	CodeRunwayPointRoleTypes_mapString      map[string]*CodeRunwayPointRoleType
	CodeRunwayPointRoleTypeOrder            uint
	CodeRunwayPointRoleType_stagedOrder     map[*CodeRunwayPointRoleType]uint
	CodeRunwayPointRoleType_orderStaged     map[uint]*CodeRunwayPointRoleType
	CodeRunwayPointRoleTypes_reference      map[*CodeRunwayPointRoleType]*CodeRunwayPointRoleType
	CodeRunwayPointRoleTypes_referenceOrder map[*CodeRunwayPointRoleType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwayPointRoleTypeCreateCallback GongOnAfterCreateInterface[CodeRunwayPointRoleType]
	OnAfterCodeRunwayPointRoleTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwayPointRoleType]
	OnAfterCodeRunwayPointRoleTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwayPointRoleType]

	CodeRunwayProtectionAreaTypes                map[*CodeRunwayProtectionAreaType]struct{}
	CodeRunwayProtectionAreaTypes_instance       map[*CodeRunwayProtectionAreaType]*CodeRunwayProtectionAreaType
	CodeRunwayProtectionAreaTypes_mapString      map[string]*CodeRunwayProtectionAreaType
	CodeRunwayProtectionAreaTypeOrder            uint
	CodeRunwayProtectionAreaType_stagedOrder     map[*CodeRunwayProtectionAreaType]uint
	CodeRunwayProtectionAreaType_orderStaged     map[uint]*CodeRunwayProtectionAreaType
	CodeRunwayProtectionAreaTypes_reference      map[*CodeRunwayProtectionAreaType]*CodeRunwayProtectionAreaType
	CodeRunwayProtectionAreaTypes_referenceOrder map[*CodeRunwayProtectionAreaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwayProtectionAreaTypeCreateCallback GongOnAfterCreateInterface[CodeRunwayProtectionAreaType]
	OnAfterCodeRunwayProtectionAreaTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwayProtectionAreaType]
	OnAfterCodeRunwayProtectionAreaTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwayProtectionAreaType]

	CodeRunwaySectionTypes                map[*CodeRunwaySectionType]struct{}
	CodeRunwaySectionTypes_instance       map[*CodeRunwaySectionType]*CodeRunwaySectionType
	CodeRunwaySectionTypes_mapString      map[string]*CodeRunwaySectionType
	CodeRunwaySectionTypeOrder            uint
	CodeRunwaySectionType_stagedOrder     map[*CodeRunwaySectionType]uint
	CodeRunwaySectionType_orderStaged     map[uint]*CodeRunwaySectionType
	CodeRunwaySectionTypes_reference      map[*CodeRunwaySectionType]*CodeRunwaySectionType
	CodeRunwaySectionTypes_referenceOrder map[*CodeRunwaySectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwaySectionTypeCreateCallback GongOnAfterCreateInterface[CodeRunwaySectionType]
	OnAfterCodeRunwaySectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwaySectionType]
	OnAfterCodeRunwaySectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwaySectionType]

	CodeRunwayTypes                map[*CodeRunwayType]struct{}
	CodeRunwayTypes_instance       map[*CodeRunwayType]*CodeRunwayType
	CodeRunwayTypes_mapString      map[string]*CodeRunwayType
	CodeRunwayTypeOrder            uint
	CodeRunwayType_stagedOrder     map[*CodeRunwayType]uint
	CodeRunwayType_orderStaged     map[uint]*CodeRunwayType
	CodeRunwayTypes_reference      map[*CodeRunwayType]*CodeRunwayType
	CodeRunwayTypes_referenceOrder map[*CodeRunwayType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeRunwayTypeCreateCallback GongOnAfterCreateInterface[CodeRunwayType]
	OnAfterCodeRunwayTypeUpdateCallback GongOnAfterUpdateInterface[CodeRunwayType]
	OnAfterCodeRunwayTypeDeleteCallback GongOnAfterDeleteInterface[CodeRunwayType]

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

	CodeStatusAirportTypes                map[*CodeStatusAirportType]struct{}
	CodeStatusAirportTypes_instance       map[*CodeStatusAirportType]*CodeStatusAirportType
	CodeStatusAirportTypes_mapString      map[string]*CodeStatusAirportType
	CodeStatusAirportTypeOrder            uint
	CodeStatusAirportType_stagedOrder     map[*CodeStatusAirportType]uint
	CodeStatusAirportType_orderStaged     map[uint]*CodeStatusAirportType
	CodeStatusAirportTypes_reference      map[*CodeStatusAirportType]*CodeStatusAirportType
	CodeStatusAirportTypes_referenceOrder map[*CodeStatusAirportType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusAirportTypeCreateCallback GongOnAfterCreateInterface[CodeStatusAirportType]
	OnAfterCodeStatusAirportTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusAirportType]
	OnAfterCodeStatusAirportTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusAirportType]

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

	CodeSurfaceCompositionTypes                map[*CodeSurfaceCompositionType]struct{}
	CodeSurfaceCompositionTypes_instance       map[*CodeSurfaceCompositionType]*CodeSurfaceCompositionType
	CodeSurfaceCompositionTypes_mapString      map[string]*CodeSurfaceCompositionType
	CodeSurfaceCompositionTypeOrder            uint
	CodeSurfaceCompositionType_stagedOrder     map[*CodeSurfaceCompositionType]uint
	CodeSurfaceCompositionType_orderStaged     map[uint]*CodeSurfaceCompositionType
	CodeSurfaceCompositionTypes_reference      map[*CodeSurfaceCompositionType]*CodeSurfaceCompositionType
	CodeSurfaceCompositionTypes_referenceOrder map[*CodeSurfaceCompositionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSurfaceCompositionTypeCreateCallback GongOnAfterCreateInterface[CodeSurfaceCompositionType]
	OnAfterCodeSurfaceCompositionTypeUpdateCallback GongOnAfterUpdateInterface[CodeSurfaceCompositionType]
	OnAfterCodeSurfaceCompositionTypeDeleteCallback GongOnAfterDeleteInterface[CodeSurfaceCompositionType]

	CodeSurfaceConditionTypes                map[*CodeSurfaceConditionType]struct{}
	CodeSurfaceConditionTypes_instance       map[*CodeSurfaceConditionType]*CodeSurfaceConditionType
	CodeSurfaceConditionTypes_mapString      map[string]*CodeSurfaceConditionType
	CodeSurfaceConditionTypeOrder            uint
	CodeSurfaceConditionType_stagedOrder     map[*CodeSurfaceConditionType]uint
	CodeSurfaceConditionType_orderStaged     map[uint]*CodeSurfaceConditionType
	CodeSurfaceConditionTypes_reference      map[*CodeSurfaceConditionType]*CodeSurfaceConditionType
	CodeSurfaceConditionTypes_referenceOrder map[*CodeSurfaceConditionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSurfaceConditionTypeCreateCallback GongOnAfterCreateInterface[CodeSurfaceConditionType]
	OnAfterCodeSurfaceConditionTypeUpdateCallback GongOnAfterUpdateInterface[CodeSurfaceConditionType]
	OnAfterCodeSurfaceConditionTypeDeleteCallback GongOnAfterDeleteInterface[CodeSurfaceConditionType]

	CodeSurfacePreparationTypes                map[*CodeSurfacePreparationType]struct{}
	CodeSurfacePreparationTypes_instance       map[*CodeSurfacePreparationType]*CodeSurfacePreparationType
	CodeSurfacePreparationTypes_mapString      map[string]*CodeSurfacePreparationType
	CodeSurfacePreparationTypeOrder            uint
	CodeSurfacePreparationType_stagedOrder     map[*CodeSurfacePreparationType]uint
	CodeSurfacePreparationType_orderStaged     map[uint]*CodeSurfacePreparationType
	CodeSurfacePreparationTypes_reference      map[*CodeSurfacePreparationType]*CodeSurfacePreparationType
	CodeSurfacePreparationTypes_referenceOrder map[*CodeSurfacePreparationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeSurfacePreparationTypeCreateCallback GongOnAfterCreateInterface[CodeSurfacePreparationType]
	OnAfterCodeSurfacePreparationTypeUpdateCallback GongOnAfterUpdateInterface[CodeSurfacePreparationType]
	OnAfterCodeSurfacePreparationTypeDeleteCallback GongOnAfterDeleteInterface[CodeSurfacePreparationType]

	CodeTLOFSectionTypes                map[*CodeTLOFSectionType]struct{}
	CodeTLOFSectionTypes_instance       map[*CodeTLOFSectionType]*CodeTLOFSectionType
	CodeTLOFSectionTypes_mapString      map[string]*CodeTLOFSectionType
	CodeTLOFSectionTypeOrder            uint
	CodeTLOFSectionType_stagedOrder     map[*CodeTLOFSectionType]uint
	CodeTLOFSectionType_orderStaged     map[uint]*CodeTLOFSectionType
	CodeTLOFSectionTypes_reference      map[*CodeTLOFSectionType]*CodeTLOFSectionType
	CodeTLOFSectionTypes_referenceOrder map[*CodeTLOFSectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTLOFSectionTypeCreateCallback GongOnAfterCreateInterface[CodeTLOFSectionType]
	OnAfterCodeTLOFSectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeTLOFSectionType]
	OnAfterCodeTLOFSectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeTLOFSectionType]

	CodeTaxiwayElementTypes                map[*CodeTaxiwayElementType]struct{}
	CodeTaxiwayElementTypes_instance       map[*CodeTaxiwayElementType]*CodeTaxiwayElementType
	CodeTaxiwayElementTypes_mapString      map[string]*CodeTaxiwayElementType
	CodeTaxiwayElementTypeOrder            uint
	CodeTaxiwayElementType_stagedOrder     map[*CodeTaxiwayElementType]uint
	CodeTaxiwayElementType_orderStaged     map[uint]*CodeTaxiwayElementType
	CodeTaxiwayElementTypes_reference      map[*CodeTaxiwayElementType]*CodeTaxiwayElementType
	CodeTaxiwayElementTypes_referenceOrder map[*CodeTaxiwayElementType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTaxiwayElementTypeCreateCallback GongOnAfterCreateInterface[CodeTaxiwayElementType]
	OnAfterCodeTaxiwayElementTypeUpdateCallback GongOnAfterUpdateInterface[CodeTaxiwayElementType]
	OnAfterCodeTaxiwayElementTypeDeleteCallback GongOnAfterDeleteInterface[CodeTaxiwayElementType]

	CodeTaxiwaySectionTypes                map[*CodeTaxiwaySectionType]struct{}
	CodeTaxiwaySectionTypes_instance       map[*CodeTaxiwaySectionType]*CodeTaxiwaySectionType
	CodeTaxiwaySectionTypes_mapString      map[string]*CodeTaxiwaySectionType
	CodeTaxiwaySectionTypeOrder            uint
	CodeTaxiwaySectionType_stagedOrder     map[*CodeTaxiwaySectionType]uint
	CodeTaxiwaySectionType_orderStaged     map[uint]*CodeTaxiwaySectionType
	CodeTaxiwaySectionTypes_reference      map[*CodeTaxiwaySectionType]*CodeTaxiwaySectionType
	CodeTaxiwaySectionTypes_referenceOrder map[*CodeTaxiwaySectionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTaxiwaySectionTypeCreateCallback GongOnAfterCreateInterface[CodeTaxiwaySectionType]
	OnAfterCodeTaxiwaySectionTypeUpdateCallback GongOnAfterUpdateInterface[CodeTaxiwaySectionType]
	OnAfterCodeTaxiwaySectionTypeDeleteCallback GongOnAfterDeleteInterface[CodeTaxiwaySectionType]

	CodeTaxiwayTypes                map[*CodeTaxiwayType]struct{}
	CodeTaxiwayTypes_instance       map[*CodeTaxiwayType]*CodeTaxiwayType
	CodeTaxiwayTypes_mapString      map[string]*CodeTaxiwayType
	CodeTaxiwayTypeOrder            uint
	CodeTaxiwayType_stagedOrder     map[*CodeTaxiwayType]uint
	CodeTaxiwayType_orderStaged     map[uint]*CodeTaxiwayType
	CodeTaxiwayTypes_reference      map[*CodeTaxiwayType]*CodeTaxiwayType
	CodeTaxiwayTypes_referenceOrder map[*CodeTaxiwayType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeTaxiwayTypeCreateCallback GongOnAfterCreateInterface[CodeTaxiwayType]
	OnAfterCodeTaxiwayTypeUpdateCallback GongOnAfterUpdateInterface[CodeTaxiwayType]
	OnAfterCodeTaxiwayTypeDeleteCallback GongOnAfterDeleteInterface[CodeTaxiwayType]

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

	CodeUsageLimitationTypes                map[*CodeUsageLimitationType]struct{}
	CodeUsageLimitationTypes_instance       map[*CodeUsageLimitationType]*CodeUsageLimitationType
	CodeUsageLimitationTypes_mapString      map[string]*CodeUsageLimitationType
	CodeUsageLimitationTypeOrder            uint
	CodeUsageLimitationType_stagedOrder     map[*CodeUsageLimitationType]uint
	CodeUsageLimitationType_orderStaged     map[uint]*CodeUsageLimitationType
	CodeUsageLimitationTypes_reference      map[*CodeUsageLimitationType]*CodeUsageLimitationType
	CodeUsageLimitationTypes_referenceOrder map[*CodeUsageLimitationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeUsageLimitationTypeCreateCallback GongOnAfterCreateInterface[CodeUsageLimitationType]
	OnAfterCodeUsageLimitationTypeUpdateCallback GongOnAfterUpdateInterface[CodeUsageLimitationType]
	OnAfterCodeUsageLimitationTypeDeleteCallback GongOnAfterDeleteInterface[CodeUsageLimitationType]

	CodeVASISTypes                map[*CodeVASISType]struct{}
	CodeVASISTypes_instance       map[*CodeVASISType]*CodeVASISType
	CodeVASISTypes_mapString      map[string]*CodeVASISType
	CodeVASISTypeOrder            uint
	CodeVASISType_stagedOrder     map[*CodeVASISType]uint
	CodeVASISType_orderStaged     map[uint]*CodeVASISType
	CodeVASISTypes_reference      map[*CodeVASISType]*CodeVASISType
	CodeVASISTypes_referenceOrder map[*CodeVASISType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVASISTypeCreateCallback GongOnAfterCreateInterface[CodeVASISType]
	OnAfterCodeVASISTypeUpdateCallback GongOnAfterUpdateInterface[CodeVASISType]
	OnAfterCodeVASISTypeDeleteCallback GongOnAfterDeleteInterface[CodeVASISType]

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

	CodeVisualDockingGuidanceTypes                map[*CodeVisualDockingGuidanceType]struct{}
	CodeVisualDockingGuidanceTypes_instance       map[*CodeVisualDockingGuidanceType]*CodeVisualDockingGuidanceType
	CodeVisualDockingGuidanceTypes_mapString      map[string]*CodeVisualDockingGuidanceType
	CodeVisualDockingGuidanceTypeOrder            uint
	CodeVisualDockingGuidanceType_stagedOrder     map[*CodeVisualDockingGuidanceType]uint
	CodeVisualDockingGuidanceType_orderStaged     map[uint]*CodeVisualDockingGuidanceType
	CodeVisualDockingGuidanceTypes_reference      map[*CodeVisualDockingGuidanceType]*CodeVisualDockingGuidanceType
	CodeVisualDockingGuidanceTypes_referenceOrder map[*CodeVisualDockingGuidanceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVisualDockingGuidanceTypeCreateCallback GongOnAfterCreateInterface[CodeVisualDockingGuidanceType]
	OnAfterCodeVisualDockingGuidanceTypeUpdateCallback GongOnAfterUpdateInterface[CodeVisualDockingGuidanceType]
	OnAfterCodeVisualDockingGuidanceTypeDeleteCallback GongOnAfterDeleteInterface[CodeVisualDockingGuidanceType]

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

	CodeWorkAreaTypes                map[*CodeWorkAreaType]struct{}
	CodeWorkAreaTypes_instance       map[*CodeWorkAreaType]*CodeWorkAreaType
	CodeWorkAreaTypes_mapString      map[string]*CodeWorkAreaType
	CodeWorkAreaTypeOrder            uint
	CodeWorkAreaType_stagedOrder     map[*CodeWorkAreaType]uint
	CodeWorkAreaType_orderStaged     map[uint]*CodeWorkAreaType
	CodeWorkAreaTypes_reference      map[*CodeWorkAreaType]*CodeWorkAreaType
	CodeWorkAreaTypes_referenceOrder map[*CodeWorkAreaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeWorkAreaTypeCreateCallback GongOnAfterCreateInterface[CodeWorkAreaType]
	OnAfterCodeWorkAreaTypeUpdateCallback GongOnAfterUpdateInterface[CodeWorkAreaType]
	OnAfterCodeWorkAreaTypeDeleteCallback GongOnAfterDeleteInterface[CodeWorkAreaType]

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

	DateTimeTypes                map[*DateTimeType]struct{}
	DateTimeTypes_instance       map[*DateTimeType]*DateTimeType
	DateTimeTypes_mapString      map[string]*DateTimeType
	DateTimeTypeOrder            uint
	DateTimeType_stagedOrder     map[*DateTimeType]uint
	DateTimeType_orderStaged     map[uint]*DateTimeType
	DateTimeTypes_reference      map[*DateTimeType]*DateTimeType
	DateTimeTypes_referenceOrder map[*DateTimeType]uint

	// insertion point for slice of pointers maps
	OnAfterDateTimeTypeCreateCallback GongOnAfterCreateInterface[DateTimeType]
	OnAfterDateTimeTypeUpdateCallback GongOnAfterUpdateInterface[DateTimeType]
	OnAfterDateTimeTypeDeleteCallback GongOnAfterDeleteInterface[DateTimeType]

	DateTypes                map[*DateType]struct{}
	DateTypes_instance       map[*DateType]*DateType
	DateTypes_mapString      map[string]*DateType
	DateTypeOrder            uint
	DateType_stagedOrder     map[*DateType]uint
	DateType_orderStaged     map[uint]*DateType
	DateTypes_reference      map[*DateType]*DateType
	DateTypes_referenceOrder map[*DateType]uint

	// insertion point for slice of pointers maps
	OnAfterDateTypeCreateCallback GongOnAfterCreateInterface[DateType]
	OnAfterDateTypeUpdateCallback GongOnAfterUpdateInterface[DateType]
	OnAfterDateTypeDeleteCallback GongOnAfterDeleteInterface[DateType]

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

	ValDepthTypes                map[*ValDepthType]struct{}
	ValDepthTypes_instance       map[*ValDepthType]*ValDepthType
	ValDepthTypes_mapString      map[string]*ValDepthType
	ValDepthTypeOrder            uint
	ValDepthType_stagedOrder     map[*ValDepthType]uint
	ValDepthType_orderStaged     map[uint]*ValDepthType
	ValDepthTypes_reference      map[*ValDepthType]*ValDepthType
	ValDepthTypes_referenceOrder map[*ValDepthType]uint

	// insertion point for slice of pointers maps
	OnAfterValDepthTypeCreateCallback GongOnAfterCreateInterface[ValDepthType]
	OnAfterValDepthTypeUpdateCallback GongOnAfterUpdateInterface[ValDepthType]
	OnAfterValDepthTypeDeleteCallback GongOnAfterDeleteInterface[ValDepthType]

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

	ValFLTypes                map[*ValFLType]struct{}
	ValFLTypes_instance       map[*ValFLType]*ValFLType
	ValFLTypes_mapString      map[string]*ValFLType
	ValFLTypeOrder            uint
	ValFLType_stagedOrder     map[*ValFLType]uint
	ValFLType_orderStaged     map[uint]*ValFLType
	ValFLTypes_reference      map[*ValFLType]*ValFLType
	ValFLTypes_referenceOrder map[*ValFLType]uint

	// insertion point for slice of pointers maps
	OnAfterValFLTypeCreateCallback GongOnAfterCreateInterface[ValFLType]
	OnAfterValFLTypeUpdateCallback GongOnAfterUpdateInterface[ValFLType]
	OnAfterValFLTypeDeleteCallback GongOnAfterDeleteInterface[ValFLType]

	ValFrictionTypes                map[*ValFrictionType]struct{}
	ValFrictionTypes_instance       map[*ValFrictionType]*ValFrictionType
	ValFrictionTypes_mapString      map[string]*ValFrictionType
	ValFrictionTypeOrder            uint
	ValFrictionType_stagedOrder     map[*ValFrictionType]uint
	ValFrictionType_orderStaged     map[uint]*ValFrictionType
	ValFrictionTypes_reference      map[*ValFrictionType]*ValFrictionType
	ValFrictionTypes_referenceOrder map[*ValFrictionType]uint

	// insertion point for slice of pointers maps
	OnAfterValFrictionTypeCreateCallback GongOnAfterCreateInterface[ValFrictionType]
	OnAfterValFrictionTypeUpdateCallback GongOnAfterUpdateInterface[ValFrictionType]
	OnAfterValFrictionTypeDeleteCallback GongOnAfterDeleteInterface[ValFrictionType]

	ValLCNTypes                map[*ValLCNType]struct{}
	ValLCNTypes_instance       map[*ValLCNType]*ValLCNType
	ValLCNTypes_mapString      map[string]*ValLCNType
	ValLCNTypeOrder            uint
	ValLCNType_stagedOrder     map[*ValLCNType]uint
	ValLCNType_orderStaged     map[uint]*ValLCNType
	ValLCNTypes_reference      map[*ValLCNType]*ValLCNType
	ValLCNTypes_referenceOrder map[*ValLCNType]uint

	// insertion point for slice of pointers maps
	OnAfterValLCNTypeCreateCallback GongOnAfterCreateInterface[ValLCNType]
	OnAfterValLCNTypeUpdateCallback GongOnAfterUpdateInterface[ValLCNType]
	OnAfterValLCNTypeDeleteCallback GongOnAfterDeleteInterface[ValLCNType]

	ValLightIntensityTypes                map[*ValLightIntensityType]struct{}
	ValLightIntensityTypes_instance       map[*ValLightIntensityType]*ValLightIntensityType
	ValLightIntensityTypes_mapString      map[string]*ValLightIntensityType
	ValLightIntensityTypeOrder            uint
	ValLightIntensityType_stagedOrder     map[*ValLightIntensityType]uint
	ValLightIntensityType_orderStaged     map[uint]*ValLightIntensityType
	ValLightIntensityTypes_reference      map[*ValLightIntensityType]*ValLightIntensityType
	ValLightIntensityTypes_referenceOrder map[*ValLightIntensityType]uint

	// insertion point for slice of pointers maps
	OnAfterValLightIntensityTypeCreateCallback GongOnAfterCreateInterface[ValLightIntensityType]
	OnAfterValLightIntensityTypeUpdateCallback GongOnAfterUpdateInterface[ValLightIntensityType]
	OnAfterValLightIntensityTypeDeleteCallback GongOnAfterDeleteInterface[ValLightIntensityType]

	ValMagneticVariationChangeTypes                map[*ValMagneticVariationChangeType]struct{}
	ValMagneticVariationChangeTypes_instance       map[*ValMagneticVariationChangeType]*ValMagneticVariationChangeType
	ValMagneticVariationChangeTypes_mapString      map[string]*ValMagneticVariationChangeType
	ValMagneticVariationChangeTypeOrder            uint
	ValMagneticVariationChangeType_stagedOrder     map[*ValMagneticVariationChangeType]uint
	ValMagneticVariationChangeType_orderStaged     map[uint]*ValMagneticVariationChangeType
	ValMagneticVariationChangeTypes_reference      map[*ValMagneticVariationChangeType]*ValMagneticVariationChangeType
	ValMagneticVariationChangeTypes_referenceOrder map[*ValMagneticVariationChangeType]uint

	// insertion point for slice of pointers maps
	OnAfterValMagneticVariationChangeTypeCreateCallback GongOnAfterCreateInterface[ValMagneticVariationChangeType]
	OnAfterValMagneticVariationChangeTypeUpdateCallback GongOnAfterUpdateInterface[ValMagneticVariationChangeType]
	OnAfterValMagneticVariationChangeTypeDeleteCallback GongOnAfterDeleteInterface[ValMagneticVariationChangeType]

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

	ValPCNTypes                map[*ValPCNType]struct{}
	ValPCNTypes_instance       map[*ValPCNType]*ValPCNType
	ValPCNTypes_mapString      map[string]*ValPCNType
	ValPCNTypeOrder            uint
	ValPCNType_stagedOrder     map[*ValPCNType]uint
	ValPCNType_orderStaged     map[uint]*ValPCNType
	ValPCNTypes_reference      map[*ValPCNType]*ValPCNType
	ValPCNTypes_referenceOrder map[*ValPCNType]uint

	// insertion point for slice of pointers maps
	OnAfterValPCNTypeCreateCallback GongOnAfterCreateInterface[ValPCNType]
	OnAfterValPCNTypeUpdateCallback GongOnAfterUpdateInterface[ValPCNType]
	OnAfterValPCNTypeDeleteCallback GongOnAfterDeleteInterface[ValPCNType]

	ValPercentTypes                map[*ValPercentType]struct{}
	ValPercentTypes_instance       map[*ValPercentType]*ValPercentType
	ValPercentTypes_mapString      map[string]*ValPercentType
	ValPercentTypeOrder            uint
	ValPercentType_stagedOrder     map[*ValPercentType]uint
	ValPercentType_orderStaged     map[uint]*ValPercentType
	ValPercentTypes_reference      map[*ValPercentType]*ValPercentType
	ValPercentTypes_referenceOrder map[*ValPercentType]uint

	// insertion point for slice of pointers maps
	OnAfterValPercentTypeCreateCallback GongOnAfterCreateInterface[ValPercentType]
	OnAfterValPercentTypeUpdateCallback GongOnAfterUpdateInterface[ValPercentType]
	OnAfterValPercentTypeDeleteCallback GongOnAfterDeleteInterface[ValPercentType]

	ValPressureTypes                map[*ValPressureType]struct{}
	ValPressureTypes_instance       map[*ValPressureType]*ValPressureType
	ValPressureTypes_mapString      map[string]*ValPressureType
	ValPressureTypeOrder            uint
	ValPressureType_stagedOrder     map[*ValPressureType]uint
	ValPressureType_orderStaged     map[uint]*ValPressureType
	ValPressureTypes_reference      map[*ValPressureType]*ValPressureType
	ValPressureTypes_referenceOrder map[*ValPressureType]uint

	// insertion point for slice of pointers maps
	OnAfterValPressureTypeCreateCallback GongOnAfterCreateInterface[ValPressureType]
	OnAfterValPressureTypeUpdateCallback GongOnAfterUpdateInterface[ValPressureType]
	OnAfterValPressureTypeDeleteCallback GongOnAfterDeleteInterface[ValPressureType]

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
	__gong__clearReferences(&stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_instance, &stage.CodeAircraftCategoryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_instance, &stage.CodeAircraftEngineNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_instance, &stage.CodeAircraftEngineTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_instance, &stage.CodeAircraftICAOTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftStandTypes_reference, &stage.CodeAircraftStandTypes_instance, &stage.CodeAircraftStandTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_instance, &stage.CodeAircraftTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_instance, &stage.CodeAircraftWingspanClassTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirportHeliportCollocationTypes_reference, &stage.CodeAirportHeliportCollocationTypes_instance, &stage.CodeAirportHeliportCollocationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirportHeliportDesignatorTypes_reference, &stage.CodeAirportHeliportDesignatorTypes_instance, &stage.CodeAirportHeliportDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirportHeliportTypes_reference, &stage.CodeAirportHeliportTypes_instance, &stage.CodeAirportHeliportTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirportWarningTypes_reference, &stage.CodeAirportWarningTypes_instance, &stage.CodeAirportWarningTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachGuidanceTypes_reference, &stage.CodeApproachGuidanceTypes_instance, &stage.CodeApproachGuidanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachLightingICAOTypes_reference, &stage.CodeApproachLightingICAOTypes_instance, &stage.CodeApproachLightingICAOTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApproachLightingTypes_reference, &stage.CodeApproachLightingTypes_instance, &stage.CodeApproachLightingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApronElementTypes_reference, &stage.CodeApronElementTypes_instance, &stage.CodeApronElementTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeApronSectionTypes_reference, &stage.CodeApronSectionTypes_instance, &stage.CodeApronSectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeArrestingGearEnergyAbsorbTypes_reference, &stage.CodeArrestingGearEnergyAbsorbTypes_instance, &stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeArrestingGearEngageDeviceTypes_reference, &stage.CodeArrestingGearEngageDeviceTypes_instance, &stage.CodeArrestingGearEngageDeviceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAuthorityRoleTypes_reference, &stage.CodeAuthorityRoleTypes_instance, &stage.CodeAuthorityRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeBuoyDesignatorTypes_reference, &stage.CodeBuoyDesignatorTypes_instance, &stage.CodeBuoyDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeBuoyTypes_reference, &stage.CodeBuoyTypes_instance, &stage.CodeBuoyTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeColourTypes_reference, &stage.CodeColourTypes_instance, &stage.CodeColourTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_instance, &stage.CodeCommunicationModeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeContaminationTypes_reference, &stage.CodeContaminationTypes_instance, &stage.CodeContaminationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDeclaredDistanceTypes_reference, &stage.CodeDeclaredDistanceTypes_instance, &stage.CodeDeclaredDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_instance, &stage.CodeDirectionTurnTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDirectionTypes_reference, &stage.CodeDirectionTypes_instance, &stage.CodeDirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_instance, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightOriginTypes_reference, &stage.CodeFlightOriginTypes_instance, &stage.CodeFlightOriginTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightPurposeTypes_reference, &stage.CodeFlightPurposeTypes_instance, &stage.CodeFlightPurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_instance, &stage.CodeFlightRuleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightStatusTypes_reference, &stage.CodeFlightStatusTypes_instance, &stage.CodeFlightStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFlightTypes_reference, &stage.CodeFlightTypes_instance, &stage.CodeFlightTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFrictionDeviceTypes_reference, &stage.CodeFrictionDeviceTypes_instance, &stage.CodeFrictionDeviceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeFrictionEstimateTypes_reference, &stage.CodeFrictionEstimateTypes_instance, &stage.CodeFrictionEstimateTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeGradeSeparationTypes_reference, &stage.CodeGradeSeparationTypes_instance, &stage.CodeGradeSeparationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeGuidanceLineTypes_reference, &stage.CodeGuidanceLineTypes_instance, &stage.CodeGuidanceLineTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeHelicopterPerformanceTypes_reference, &stage.CodeHelicopterPerformanceTypes_instance, &stage.CodeHelicopterPerformanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeHoldingCategoryTypes_reference, &stage.CodeHoldingCategoryTypes_instance, &stage.CodeHoldingCategoryTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeIATATypes_reference, &stage.CodeIATATypes_instance, &stage.CodeIATATypes_referenceOrder)

	__gong__clearReferences(&stage.CodeICAOTypes_reference, &stage.CodeICAOTypes_instance, &stage.CodeICAOTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightHoldingPositionTypes_reference, &stage.CodeLightHoldingPositionTypes_instance, &stage.CodeLightHoldingPositionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_instance, &stage.CodeLightIntensityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightSourceTypes_reference, &stage.CodeLightSourceTypes_instance, &stage.CodeLightSourceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightingJARTypes_reference, &stage.CodeLightingJARTypes_instance, &stage.CodeLightingJARTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLoadingBridgeTypes_reference, &stage.CodeLoadingBridgeTypes_instance, &stage.CodeLoadingBridgeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLogicalOperatorTypes_reference, &stage.CodeLogicalOperatorTypes_instance, &stage.CodeLogicalOperatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMarkingConditionTypes_reference, &stage.CodeMarkingConditionTypes_instance, &stage.CodeMarkingConditionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMarkingStyleTypes_reference, &stage.CodeMarkingStyleTypes_instance, &stage.CodeMarkingStyleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMeteoConditionsTypes_reference, &stage.CodeMeteoConditionsTypes_instance, &stage.CodeMeteoConditionsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_instance, &stage.CodeMilitaryOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryStatusTypes_reference, &stage.CodeMilitaryStatusTypes_instance, &stage.CodeMilitaryStatusTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_instance, &stage.CodeNavigationEquipmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_instance, &stage.CodeNavigationSpecificationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOperationAirportHeliportTypes_reference, &stage.CodeOperationAirportHeliportTypes_instance, &stage.CodeOperationAirportHeliportTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeOperationManoeuvringAreaTypes_reference, &stage.CodeOperationManoeuvringAreaTypes_instance, &stage.CodeOperationManoeuvringAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePCNMethodTypes_reference, &stage.CodePCNMethodTypes_instance, &stage.CodePCNMethodTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePCNPavementTypes_reference, &stage.CodePCNPavementTypes_instance, &stage.CodePCNPavementTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePCNSubgradeTypes_reference, &stage.CodePCNSubgradeTypes_instance, &stage.CodePCNSubgradeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodePCNTyrePressureTypes_reference, &stage.CodePCNTyrePressureTypes_instance, &stage.CodePCNTyrePressureTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeProtectAreaSectionTypes_reference, &stage.CodeProtectAreaSectionTypes_instance, &stage.CodeProtectAreaSectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVRReadingTypes_reference, &stage.CodeRVRReadingTypes_instance, &stage.CodeRVRReadingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_instance, &stage.CodeRVSMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRoadTypes_reference, &stage.CodeRoadTypes_instance, &stage.CodeRoadTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwayElementTypes_reference, &stage.CodeRunwayElementTypes_instance, &stage.CodeRunwayElementTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwayMarkingTypes_reference, &stage.CodeRunwayMarkingTypes_instance, &stage.CodeRunwayMarkingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwayPointRoleTypes_reference, &stage.CodeRunwayPointRoleTypes_instance, &stage.CodeRunwayPointRoleTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwayProtectionAreaTypes_reference, &stage.CodeRunwayProtectionAreaTypes_instance, &stage.CodeRunwayProtectionAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwaySectionTypes_reference, &stage.CodeRunwaySectionTypes_instance, &stage.CodeRunwaySectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRunwayTypes_reference, &stage.CodeRunwayTypes_instance, &stage.CodeRunwayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSideTypes_reference, &stage.CodeSideTypes_instance, &stage.CodeSideTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusAirportTypes_reference, &stage.CodeStatusAirportTypes_instance, &stage.CodeStatusAirportTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_instance, &stage.CodeStatusOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSurfaceCompositionTypes_reference, &stage.CodeSurfaceCompositionTypes_instance, &stage.CodeSurfaceCompositionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSurfaceConditionTypes_reference, &stage.CodeSurfaceConditionTypes_instance, &stage.CodeSurfaceConditionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeSurfacePreparationTypes_reference, &stage.CodeSurfacePreparationTypes_instance, &stage.CodeSurfacePreparationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTLOFSectionTypes_reference, &stage.CodeTLOFSectionTypes_instance, &stage.CodeTLOFSectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTaxiwayElementTypes_reference, &stage.CodeTaxiwayElementTypes_instance, &stage.CodeTaxiwayElementTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTaxiwaySectionTypes_reference, &stage.CodeTaxiwaySectionTypes_instance, &stage.CodeTaxiwaySectionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTaxiwayTypes_reference, &stage.CodeTaxiwayTypes_instance, &stage.CodeTaxiwayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_instance, &stage.CodeTelecomNetworkTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_instance, &stage.CodeTransponderTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeUsageLimitationTypes_reference, &stage.CodeUsageLimitationTypes_instance, &stage.CodeUsageLimitationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVASISTypes_reference, &stage.CodeVASISTypes_instance, &stage.CodeVASISTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_instance, &stage.CodeValueInterpretationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVisualDockingGuidanceTypes_reference, &stage.CodeVisualDockingGuidanceTypes_instance, &stage.CodeVisualDockingGuidanceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_instance, &stage.CodeWakeTurbulenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWorkAreaTypes_reference, &stage.CodeWorkAreaTypes_instance, &stage.CodeWorkAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.DateTimeTypes_reference, &stage.DateTimeTypes_instance, &stage.DateTimeTypes_referenceOrder)

	__gong__clearReferences(&stage.DateTypes_reference, &stage.DateTypes_instance, &stage.DateTypes_referenceOrder)

	__gong__clearReferences(&stage.DateYearTypes_reference, &stage.DateYearTypes_instance, &stage.DateYearTypes_referenceOrder)

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

	__gong__clearReferences(&stage.ValDepthTypes_reference, &stage.ValDepthTypes_instance, &stage.ValDepthTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValFLTypes_reference, &stage.ValFLTypes_instance, &stage.ValFLTypes_referenceOrder)

	__gong__clearReferences(&stage.ValFrictionTypes_reference, &stage.ValFrictionTypes_instance, &stage.ValFrictionTypes_referenceOrder)

	__gong__clearReferences(&stage.ValLCNTypes_reference, &stage.ValLCNTypes_instance, &stage.ValLCNTypes_referenceOrder)

	__gong__clearReferences(&stage.ValLightIntensityTypes_reference, &stage.ValLightIntensityTypes_instance, &stage.ValLightIntensityTypes_referenceOrder)

	__gong__clearReferences(&stage.ValMagneticVariationChangeTypes_reference, &stage.ValMagneticVariationChangeTypes_instance, &stage.ValMagneticVariationChangeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValMagneticVariationTypes_reference, &stage.ValMagneticVariationTypes_instance, &stage.ValMagneticVariationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValPCNTypes_reference, &stage.ValPCNTypes_instance, &stage.ValPCNTypes_referenceOrder)

	__gong__clearReferences(&stage.ValPercentTypes_reference, &stage.ValPercentTypes_instance, &stage.ValPercentTypes_referenceOrder)

	__gong__clearReferences(&stage.ValPressureTypes_reference, &stage.ValPressureTypes_instance, &stage.ValPressureTypes_referenceOrder)

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
	stage.CodeAircraftCategoryTypeOrder = __gong__recomputeOrder(stage.CodeAircraftCategoryType_stagedOrder)

	stage.CodeAircraftEngineNumberTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineNumberType_stagedOrder)

	stage.CodeAircraftEngineTypeOrder = __gong__recomputeOrder(stage.CodeAircraftEngineType_stagedOrder)

	stage.CodeAircraftICAOTypeOrder = __gong__recomputeOrder(stage.CodeAircraftICAOType_stagedOrder)

	stage.CodeAircraftStandTypeOrder = __gong__recomputeOrder(stage.CodeAircraftStandType_stagedOrder)

	stage.CodeAircraftTypeOrder = __gong__recomputeOrder(stage.CodeAircraftType_stagedOrder)

	stage.CodeAircraftWingspanClassTypeOrder = __gong__recomputeOrder(stage.CodeAircraftWingspanClassType_stagedOrder)

	stage.CodeAirportHeliportCollocationTypeOrder = __gong__recomputeOrder(stage.CodeAirportHeliportCollocationType_stagedOrder)

	stage.CodeAirportHeliportDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeAirportHeliportDesignatorType_stagedOrder)

	stage.CodeAirportHeliportTypeOrder = __gong__recomputeOrder(stage.CodeAirportHeliportType_stagedOrder)

	stage.CodeAirportWarningTypeOrder = __gong__recomputeOrder(stage.CodeAirportWarningType_stagedOrder)

	stage.CodeApproachGuidanceTypeOrder = __gong__recomputeOrder(stage.CodeApproachGuidanceType_stagedOrder)

	stage.CodeApproachLightingICAOTypeOrder = __gong__recomputeOrder(stage.CodeApproachLightingICAOType_stagedOrder)

	stage.CodeApproachLightingTypeOrder = __gong__recomputeOrder(stage.CodeApproachLightingType_stagedOrder)

	stage.CodeApronElementTypeOrder = __gong__recomputeOrder(stage.CodeApronElementType_stagedOrder)

	stage.CodeApronSectionTypeOrder = __gong__recomputeOrder(stage.CodeApronSectionType_stagedOrder)

	stage.CodeArrestingGearEnergyAbsorbTypeOrder = __gong__recomputeOrder(stage.CodeArrestingGearEnergyAbsorbType_stagedOrder)

	stage.CodeArrestingGearEngageDeviceTypeOrder = __gong__recomputeOrder(stage.CodeArrestingGearEngageDeviceType_stagedOrder)

	stage.CodeAuthorityRoleTypeOrder = __gong__recomputeOrder(stage.CodeAuthorityRoleType_stagedOrder)

	stage.CodeBuoyDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeBuoyDesignatorType_stagedOrder)

	stage.CodeBuoyTypeOrder = __gong__recomputeOrder(stage.CodeBuoyType_stagedOrder)

	stage.CodeColourTypeOrder = __gong__recomputeOrder(stage.CodeColourType_stagedOrder)

	stage.CodeCommunicationModeTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationModeType_stagedOrder)

	stage.CodeContaminationTypeOrder = __gong__recomputeOrder(stage.CodeContaminationType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeDeclaredDistanceTypeOrder = __gong__recomputeOrder(stage.CodeDeclaredDistanceType_stagedOrder)

	stage.CodeDirectionTurnTypeOrder = __gong__recomputeOrder(stage.CodeDirectionTurnType_stagedOrder)

	stage.CodeDirectionTypeOrder = __gong__recomputeOrder(stage.CodeDirectionType_stagedOrder)

	stage.CodeEquipmentAntiCollisionTypeOrder = __gong__recomputeOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder)

	stage.CodeFlightOriginTypeOrder = __gong__recomputeOrder(stage.CodeFlightOriginType_stagedOrder)

	stage.CodeFlightPurposeTypeOrder = __gong__recomputeOrder(stage.CodeFlightPurposeType_stagedOrder)

	stage.CodeFlightRuleTypeOrder = __gong__recomputeOrder(stage.CodeFlightRuleType_stagedOrder)

	stage.CodeFlightStatusTypeOrder = __gong__recomputeOrder(stage.CodeFlightStatusType_stagedOrder)

	stage.CodeFlightTypeOrder = __gong__recomputeOrder(stage.CodeFlightType_stagedOrder)

	stage.CodeFrictionDeviceTypeOrder = __gong__recomputeOrder(stage.CodeFrictionDeviceType_stagedOrder)

	stage.CodeFrictionEstimateTypeOrder = __gong__recomputeOrder(stage.CodeFrictionEstimateType_stagedOrder)

	stage.CodeGradeSeparationTypeOrder = __gong__recomputeOrder(stage.CodeGradeSeparationType_stagedOrder)

	stage.CodeGuidanceLineTypeOrder = __gong__recomputeOrder(stage.CodeGuidanceLineType_stagedOrder)

	stage.CodeHelicopterPerformanceTypeOrder = __gong__recomputeOrder(stage.CodeHelicopterPerformanceType_stagedOrder)

	stage.CodeHoldingCategoryTypeOrder = __gong__recomputeOrder(stage.CodeHoldingCategoryType_stagedOrder)

	stage.CodeIATATypeOrder = __gong__recomputeOrder(stage.CodeIATAType_stagedOrder)

	stage.CodeICAOTypeOrder = __gong__recomputeOrder(stage.CodeICAOType_stagedOrder)

	stage.CodeLightHoldingPositionTypeOrder = __gong__recomputeOrder(stage.CodeLightHoldingPositionType_stagedOrder)

	stage.CodeLightIntensityTypeOrder = __gong__recomputeOrder(stage.CodeLightIntensityType_stagedOrder)

	stage.CodeLightSourceTypeOrder = __gong__recomputeOrder(stage.CodeLightSourceType_stagedOrder)

	stage.CodeLightingJARTypeOrder = __gong__recomputeOrder(stage.CodeLightingJARType_stagedOrder)

	stage.CodeLoadingBridgeTypeOrder = __gong__recomputeOrder(stage.CodeLoadingBridgeType_stagedOrder)

	stage.CodeLogicalOperatorTypeOrder = __gong__recomputeOrder(stage.CodeLogicalOperatorType_stagedOrder)

	stage.CodeMarkingConditionTypeOrder = __gong__recomputeOrder(stage.CodeMarkingConditionType_stagedOrder)

	stage.CodeMarkingStyleTypeOrder = __gong__recomputeOrder(stage.CodeMarkingStyleType_stagedOrder)

	stage.CodeMeteoConditionsTypeOrder = __gong__recomputeOrder(stage.CodeMeteoConditionsType_stagedOrder)

	stage.CodeMilitaryOperationsTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryOperationsType_stagedOrder)

	stage.CodeMilitaryStatusTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryStatusType_stagedOrder)

	stage.CodeNavigationEquipmentTypeOrder = __gong__recomputeOrder(stage.CodeNavigationEquipmentType_stagedOrder)

	stage.CodeNavigationSpecificationTypeOrder = __gong__recomputeOrder(stage.CodeNavigationSpecificationType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeOperationAirportHeliportTypeOrder = __gong__recomputeOrder(stage.CodeOperationAirportHeliportType_stagedOrder)

	stage.CodeOperationManoeuvringAreaTypeOrder = __gong__recomputeOrder(stage.CodeOperationManoeuvringAreaType_stagedOrder)

	stage.CodePCNMethodTypeOrder = __gong__recomputeOrder(stage.CodePCNMethodType_stagedOrder)

	stage.CodePCNPavementTypeOrder = __gong__recomputeOrder(stage.CodePCNPavementType_stagedOrder)

	stage.CodePCNSubgradeTypeOrder = __gong__recomputeOrder(stage.CodePCNSubgradeType_stagedOrder)

	stage.CodePCNTyrePressureTypeOrder = __gong__recomputeOrder(stage.CodePCNTyrePressureType_stagedOrder)

	stage.CodeProtectAreaSectionTypeOrder = __gong__recomputeOrder(stage.CodeProtectAreaSectionType_stagedOrder)

	stage.CodeRVRReadingTypeOrder = __gong__recomputeOrder(stage.CodeRVRReadingType_stagedOrder)

	stage.CodeRVSMTypeOrder = __gong__recomputeOrder(stage.CodeRVSMType_stagedOrder)

	stage.CodeRoadTypeOrder = __gong__recomputeOrder(stage.CodeRoadType_stagedOrder)

	stage.CodeRunwayElementTypeOrder = __gong__recomputeOrder(stage.CodeRunwayElementType_stagedOrder)

	stage.CodeRunwayMarkingTypeOrder = __gong__recomputeOrder(stage.CodeRunwayMarkingType_stagedOrder)

	stage.CodeRunwayPointRoleTypeOrder = __gong__recomputeOrder(stage.CodeRunwayPointRoleType_stagedOrder)

	stage.CodeRunwayProtectionAreaTypeOrder = __gong__recomputeOrder(stage.CodeRunwayProtectionAreaType_stagedOrder)

	stage.CodeRunwaySectionTypeOrder = __gong__recomputeOrder(stage.CodeRunwaySectionType_stagedOrder)

	stage.CodeRunwayTypeOrder = __gong__recomputeOrder(stage.CodeRunwayType_stagedOrder)

	stage.CodeSideTypeOrder = __gong__recomputeOrder(stage.CodeSideType_stagedOrder)

	stage.CodeStatusAirportTypeOrder = __gong__recomputeOrder(stage.CodeStatusAirportType_stagedOrder)

	stage.CodeStatusOperationsTypeOrder = __gong__recomputeOrder(stage.CodeStatusOperationsType_stagedOrder)

	stage.CodeSurfaceCompositionTypeOrder = __gong__recomputeOrder(stage.CodeSurfaceCompositionType_stagedOrder)

	stage.CodeSurfaceConditionTypeOrder = __gong__recomputeOrder(stage.CodeSurfaceConditionType_stagedOrder)

	stage.CodeSurfacePreparationTypeOrder = __gong__recomputeOrder(stage.CodeSurfacePreparationType_stagedOrder)

	stage.CodeTLOFSectionTypeOrder = __gong__recomputeOrder(stage.CodeTLOFSectionType_stagedOrder)

	stage.CodeTaxiwayElementTypeOrder = __gong__recomputeOrder(stage.CodeTaxiwayElementType_stagedOrder)

	stage.CodeTaxiwaySectionTypeOrder = __gong__recomputeOrder(stage.CodeTaxiwaySectionType_stagedOrder)

	stage.CodeTaxiwayTypeOrder = __gong__recomputeOrder(stage.CodeTaxiwayType_stagedOrder)

	stage.CodeTelecomNetworkTypeOrder = __gong__recomputeOrder(stage.CodeTelecomNetworkType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeTransponderTypeOrder = __gong__recomputeOrder(stage.CodeTransponderType_stagedOrder)

	stage.CodeUsageLimitationTypeOrder = __gong__recomputeOrder(stage.CodeUsageLimitationType_stagedOrder)

	stage.CodeVASISTypeOrder = __gong__recomputeOrder(stage.CodeVASISType_stagedOrder)

	stage.CodeValueInterpretationTypeOrder = __gong__recomputeOrder(stage.CodeValueInterpretationType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeVisualDockingGuidanceTypeOrder = __gong__recomputeOrder(stage.CodeVisualDockingGuidanceType_stagedOrder)

	stage.CodeWakeTurbulenceTypeOrder = __gong__recomputeOrder(stage.CodeWakeTurbulenceType_stagedOrder)

	stage.CodeWorkAreaTypeOrder = __gong__recomputeOrder(stage.CodeWorkAreaType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.DateTimeTypeOrder = __gong__recomputeOrder(stage.DateTimeType_stagedOrder)

	stage.DateTypeOrder = __gong__recomputeOrder(stage.DateType_stagedOrder)

	stage.DateYearTypeOrder = __gong__recomputeOrder(stage.DateYearType_stagedOrder)

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

	stage.ValDepthTypeOrder = __gong__recomputeOrder(stage.ValDepthType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValFLTypeOrder = __gong__recomputeOrder(stage.ValFLType_stagedOrder)

	stage.ValFrictionTypeOrder = __gong__recomputeOrder(stage.ValFrictionType_stagedOrder)

	stage.ValLCNTypeOrder = __gong__recomputeOrder(stage.ValLCNType_stagedOrder)

	stage.ValLightIntensityTypeOrder = __gong__recomputeOrder(stage.ValLightIntensityType_stagedOrder)

	stage.ValMagneticVariationChangeTypeOrder = __gong__recomputeOrder(stage.ValMagneticVariationChangeType_stagedOrder)

	stage.ValMagneticVariationTypeOrder = __gong__recomputeOrder(stage.ValMagneticVariationType_stagedOrder)

	stage.ValPCNTypeOrder = __gong__recomputeOrder(stage.ValPCNType_stagedOrder)

	stage.ValPercentTypeOrder = __gong__recomputeOrder(stage.ValPercentType_stagedOrder)

	stage.ValPressureTypeOrder = __gong__recomputeOrder(stage.ValPressureType_stagedOrder)

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

func (*CodeAircraftStandType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAircraftStandTypes, stage.CodeAircraftStandType_stagedOrder)
}

func (*CodeAircraftStandType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAircraftStandType_orderStaged[order]
}

func (*CodeAircraftStandType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAircraftStandTypes_mapString
}

func (*CodeAircraftStandType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAircraftStandTypes
}

func (*CodeAircraftStandType) GongNewInstance() any {
	return new(CodeAircraftStandType)
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

func (*CodeAirportHeliportCollocationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationType_stagedOrder)
}

func (*CodeAirportHeliportCollocationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirportHeliportCollocationType_orderStaged[order]
}

func (*CodeAirportHeliportCollocationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirportHeliportCollocationTypes_mapString
}

func (*CodeAirportHeliportCollocationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirportHeliportCollocationTypes
}

func (*CodeAirportHeliportCollocationType) GongNewInstance() any {
	return new(CodeAirportHeliportCollocationType)
}

func (*CodeAirportHeliportDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorType_stagedOrder)
}

func (*CodeAirportHeliportDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirportHeliportDesignatorType_orderStaged[order]
}

func (*CodeAirportHeliportDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirportHeliportDesignatorTypes_mapString
}

func (*CodeAirportHeliportDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirportHeliportDesignatorTypes
}

func (*CodeAirportHeliportDesignatorType) GongNewInstance() any {
	return new(CodeAirportHeliportDesignatorType)
}

func (*CodeAirportHeliportType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportType_stagedOrder)
}

func (*CodeAirportHeliportType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirportHeliportType_orderStaged[order]
}

func (*CodeAirportHeliportType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirportHeliportTypes_mapString
}

func (*CodeAirportHeliportType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirportHeliportTypes
}

func (*CodeAirportHeliportType) GongNewInstance() any {
	return new(CodeAirportHeliportType)
}

func (*CodeAirportWarningType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirportWarningTypes, stage.CodeAirportWarningType_stagedOrder)
}

func (*CodeAirportWarningType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirportWarningType_orderStaged[order]
}

func (*CodeAirportWarningType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirportWarningTypes_mapString
}

func (*CodeAirportWarningType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirportWarningTypes
}

func (*CodeAirportWarningType) GongNewInstance() any {
	return new(CodeAirportWarningType)
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

func (*CodeApproachLightingICAOType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOType_stagedOrder)
}

func (*CodeApproachLightingICAOType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachLightingICAOType_orderStaged[order]
}

func (*CodeApproachLightingICAOType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachLightingICAOTypes_mapString
}

func (*CodeApproachLightingICAOType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachLightingICAOTypes
}

func (*CodeApproachLightingICAOType) GongNewInstance() any {
	return new(CodeApproachLightingICAOType)
}

func (*CodeApproachLightingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApproachLightingTypes, stage.CodeApproachLightingType_stagedOrder)
}

func (*CodeApproachLightingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApproachLightingType_orderStaged[order]
}

func (*CodeApproachLightingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApproachLightingTypes_mapString
}

func (*CodeApproachLightingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApproachLightingTypes
}

func (*CodeApproachLightingType) GongNewInstance() any {
	return new(CodeApproachLightingType)
}

func (*CodeApronElementType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApronElementTypes, stage.CodeApronElementType_stagedOrder)
}

func (*CodeApronElementType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApronElementType_orderStaged[order]
}

func (*CodeApronElementType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApronElementTypes_mapString
}

func (*CodeApronElementType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApronElementTypes
}

func (*CodeApronElementType) GongNewInstance() any {
	return new(CodeApronElementType)
}

func (*CodeApronSectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeApronSectionTypes, stage.CodeApronSectionType_stagedOrder)
}

func (*CodeApronSectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeApronSectionType_orderStaged[order]
}

func (*CodeApronSectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeApronSectionTypes_mapString
}

func (*CodeApronSectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeApronSectionTypes
}

func (*CodeApronSectionType) GongNewInstance() any {
	return new(CodeApronSectionType)
}

func (*CodeArrestingGearEnergyAbsorbType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbType_stagedOrder)
}

func (*CodeArrestingGearEnergyAbsorbType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeArrestingGearEnergyAbsorbType_orderStaged[order]
}

func (*CodeArrestingGearEnergyAbsorbType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeArrestingGearEnergyAbsorbTypes_mapString
}

func (*CodeArrestingGearEnergyAbsorbType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeArrestingGearEnergyAbsorbTypes
}

func (*CodeArrestingGearEnergyAbsorbType) GongNewInstance() any {
	return new(CodeArrestingGearEnergyAbsorbType)
}

func (*CodeArrestingGearEngageDeviceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceType_stagedOrder)
}

func (*CodeArrestingGearEngageDeviceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeArrestingGearEngageDeviceType_orderStaged[order]
}

func (*CodeArrestingGearEngageDeviceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeArrestingGearEngageDeviceTypes_mapString
}

func (*CodeArrestingGearEngageDeviceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeArrestingGearEngageDeviceTypes
}

func (*CodeArrestingGearEngageDeviceType) GongNewInstance() any {
	return new(CodeArrestingGearEngageDeviceType)
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

func (*CodeBuoyDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorType_stagedOrder)
}

func (*CodeBuoyDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeBuoyDesignatorType_orderStaged[order]
}

func (*CodeBuoyDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeBuoyDesignatorTypes_mapString
}

func (*CodeBuoyDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeBuoyDesignatorTypes
}

func (*CodeBuoyDesignatorType) GongNewInstance() any {
	return new(CodeBuoyDesignatorType)
}

func (*CodeBuoyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeBuoyTypes, stage.CodeBuoyType_stagedOrder)
}

func (*CodeBuoyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeBuoyType_orderStaged[order]
}

func (*CodeBuoyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeBuoyTypes_mapString
}

func (*CodeBuoyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeBuoyTypes
}

func (*CodeBuoyType) GongNewInstance() any {
	return new(CodeBuoyType)
}

func (*CodeColourType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeColourTypes, stage.CodeColourType_stagedOrder)
}

func (*CodeColourType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeColourType_orderStaged[order]
}

func (*CodeColourType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeColourTypes_mapString
}

func (*CodeColourType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeColourTypes
}

func (*CodeColourType) GongNewInstance() any {
	return new(CodeColourType)
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

func (*CodeContaminationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeContaminationTypes, stage.CodeContaminationType_stagedOrder)
}

func (*CodeContaminationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeContaminationType_orderStaged[order]
}

func (*CodeContaminationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeContaminationTypes_mapString
}

func (*CodeContaminationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeContaminationTypes
}

func (*CodeContaminationType) GongNewInstance() any {
	return new(CodeContaminationType)
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

func (*CodeDeclaredDistanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceType_stagedOrder)
}

func (*CodeDeclaredDistanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeDeclaredDistanceType_orderStaged[order]
}

func (*CodeDeclaredDistanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeDeclaredDistanceTypes_mapString
}

func (*CodeDeclaredDistanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeDeclaredDistanceTypes
}

func (*CodeDeclaredDistanceType) GongNewInstance() any {
	return new(CodeDeclaredDistanceType)
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

func (*CodeFrictionDeviceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceType_stagedOrder)
}

func (*CodeFrictionDeviceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFrictionDeviceType_orderStaged[order]
}

func (*CodeFrictionDeviceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFrictionDeviceTypes_mapString
}

func (*CodeFrictionDeviceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFrictionDeviceTypes
}

func (*CodeFrictionDeviceType) GongNewInstance() any {
	return new(CodeFrictionDeviceType)
}

func (*CodeFrictionEstimateType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateType_stagedOrder)
}

func (*CodeFrictionEstimateType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeFrictionEstimateType_orderStaged[order]
}

func (*CodeFrictionEstimateType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeFrictionEstimateTypes_mapString
}

func (*CodeFrictionEstimateType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeFrictionEstimateTypes
}

func (*CodeFrictionEstimateType) GongNewInstance() any {
	return new(CodeFrictionEstimateType)
}

func (*CodeGradeSeparationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationType_stagedOrder)
}

func (*CodeGradeSeparationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeGradeSeparationType_orderStaged[order]
}

func (*CodeGradeSeparationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeGradeSeparationTypes_mapString
}

func (*CodeGradeSeparationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeGradeSeparationTypes
}

func (*CodeGradeSeparationType) GongNewInstance() any {
	return new(CodeGradeSeparationType)
}

func (*CodeGuidanceLineType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineType_stagedOrder)
}

func (*CodeGuidanceLineType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeGuidanceLineType_orderStaged[order]
}

func (*CodeGuidanceLineType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeGuidanceLineTypes_mapString
}

func (*CodeGuidanceLineType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeGuidanceLineTypes
}

func (*CodeGuidanceLineType) GongNewInstance() any {
	return new(CodeGuidanceLineType)
}

func (*CodeHelicopterPerformanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceType_stagedOrder)
}

func (*CodeHelicopterPerformanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeHelicopterPerformanceType_orderStaged[order]
}

func (*CodeHelicopterPerformanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeHelicopterPerformanceTypes_mapString
}

func (*CodeHelicopterPerformanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeHelicopterPerformanceTypes
}

func (*CodeHelicopterPerformanceType) GongNewInstance() any {
	return new(CodeHelicopterPerformanceType)
}

func (*CodeHoldingCategoryType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryType_stagedOrder)
}

func (*CodeHoldingCategoryType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeHoldingCategoryType_orderStaged[order]
}

func (*CodeHoldingCategoryType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeHoldingCategoryTypes_mapString
}

func (*CodeHoldingCategoryType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeHoldingCategoryTypes
}

func (*CodeHoldingCategoryType) GongNewInstance() any {
	return new(CodeHoldingCategoryType)
}

func (*CodeIATAType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeIATATypes, stage.CodeIATAType_stagedOrder)
}

func (*CodeIATAType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeIATAType_orderStaged[order]
}

func (*CodeIATAType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeIATATypes_mapString
}

func (*CodeIATAType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeIATATypes
}

func (*CodeIATAType) GongNewInstance() any {
	return new(CodeIATAType)
}

func (*CodeICAOType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeICAOTypes, stage.CodeICAOType_stagedOrder)
}

func (*CodeICAOType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeICAOType_orderStaged[order]
}

func (*CodeICAOType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeICAOTypes_mapString
}

func (*CodeICAOType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeICAOTypes
}

func (*CodeICAOType) GongNewInstance() any {
	return new(CodeICAOType)
}

func (*CodeLightHoldingPositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionType_stagedOrder)
}

func (*CodeLightHoldingPositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLightHoldingPositionType_orderStaged[order]
}

func (*CodeLightHoldingPositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLightHoldingPositionTypes_mapString
}

func (*CodeLightHoldingPositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLightHoldingPositionTypes
}

func (*CodeLightHoldingPositionType) GongNewInstance() any {
	return new(CodeLightHoldingPositionType)
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

func (*CodeLightSourceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLightSourceTypes, stage.CodeLightSourceType_stagedOrder)
}

func (*CodeLightSourceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLightSourceType_orderStaged[order]
}

func (*CodeLightSourceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLightSourceTypes_mapString
}

func (*CodeLightSourceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLightSourceTypes
}

func (*CodeLightSourceType) GongNewInstance() any {
	return new(CodeLightSourceType)
}

func (*CodeLightingJARType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLightingJARTypes, stage.CodeLightingJARType_stagedOrder)
}

func (*CodeLightingJARType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLightingJARType_orderStaged[order]
}

func (*CodeLightingJARType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLightingJARTypes_mapString
}

func (*CodeLightingJARType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLightingJARTypes
}

func (*CodeLightingJARType) GongNewInstance() any {
	return new(CodeLightingJARType)
}

func (*CodeLoadingBridgeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeType_stagedOrder)
}

func (*CodeLoadingBridgeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLoadingBridgeType_orderStaged[order]
}

func (*CodeLoadingBridgeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLoadingBridgeTypes_mapString
}

func (*CodeLoadingBridgeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLoadingBridgeTypes
}

func (*CodeLoadingBridgeType) GongNewInstance() any {
	return new(CodeLoadingBridgeType)
}

func (*CodeLogicalOperatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorType_stagedOrder)
}

func (*CodeLogicalOperatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeLogicalOperatorType_orderStaged[order]
}

func (*CodeLogicalOperatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeLogicalOperatorTypes_mapString
}

func (*CodeLogicalOperatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeLogicalOperatorTypes
}

func (*CodeLogicalOperatorType) GongNewInstance() any {
	return new(CodeLogicalOperatorType)
}

func (*CodeMarkingConditionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionType_stagedOrder)
}

func (*CodeMarkingConditionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMarkingConditionType_orderStaged[order]
}

func (*CodeMarkingConditionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMarkingConditionTypes_mapString
}

func (*CodeMarkingConditionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMarkingConditionTypes
}

func (*CodeMarkingConditionType) GongNewInstance() any {
	return new(CodeMarkingConditionType)
}

func (*CodeMarkingStyleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleType_stagedOrder)
}

func (*CodeMarkingStyleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeMarkingStyleType_orderStaged[order]
}

func (*CodeMarkingStyleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeMarkingStyleTypes_mapString
}

func (*CodeMarkingStyleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeMarkingStyleTypes
}

func (*CodeMarkingStyleType) GongNewInstance() any {
	return new(CodeMarkingStyleType)
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

func (*CodeOperationAirportHeliportType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportType_stagedOrder)
}

func (*CodeOperationAirportHeliportType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOperationAirportHeliportType_orderStaged[order]
}

func (*CodeOperationAirportHeliportType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOperationAirportHeliportTypes_mapString
}

func (*CodeOperationAirportHeliportType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOperationAirportHeliportTypes
}

func (*CodeOperationAirportHeliportType) GongNewInstance() any {
	return new(CodeOperationAirportHeliportType)
}

func (*CodeOperationManoeuvringAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaType_stagedOrder)
}

func (*CodeOperationManoeuvringAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeOperationManoeuvringAreaType_orderStaged[order]
}

func (*CodeOperationManoeuvringAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeOperationManoeuvringAreaTypes_mapString
}

func (*CodeOperationManoeuvringAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeOperationManoeuvringAreaTypes
}

func (*CodeOperationManoeuvringAreaType) GongNewInstance() any {
	return new(CodeOperationManoeuvringAreaType)
}

func (*CodePCNMethodType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePCNMethodTypes, stage.CodePCNMethodType_stagedOrder)
}

func (*CodePCNMethodType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePCNMethodType_orderStaged[order]
}

func (*CodePCNMethodType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePCNMethodTypes_mapString
}

func (*CodePCNMethodType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePCNMethodTypes
}

func (*CodePCNMethodType) GongNewInstance() any {
	return new(CodePCNMethodType)
}

func (*CodePCNPavementType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePCNPavementTypes, stage.CodePCNPavementType_stagedOrder)
}

func (*CodePCNPavementType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePCNPavementType_orderStaged[order]
}

func (*CodePCNPavementType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePCNPavementTypes_mapString
}

func (*CodePCNPavementType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePCNPavementTypes
}

func (*CodePCNPavementType) GongNewInstance() any {
	return new(CodePCNPavementType)
}

func (*CodePCNSubgradeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeType_stagedOrder)
}

func (*CodePCNSubgradeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePCNSubgradeType_orderStaged[order]
}

func (*CodePCNSubgradeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePCNSubgradeTypes_mapString
}

func (*CodePCNSubgradeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePCNSubgradeTypes
}

func (*CodePCNSubgradeType) GongNewInstance() any {
	return new(CodePCNSubgradeType)
}

func (*CodePCNTyrePressureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureType_stagedOrder)
}

func (*CodePCNTyrePressureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodePCNTyrePressureType_orderStaged[order]
}

func (*CodePCNTyrePressureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodePCNTyrePressureTypes_mapString
}

func (*CodePCNTyrePressureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodePCNTyrePressureTypes
}

func (*CodePCNTyrePressureType) GongNewInstance() any {
	return new(CodePCNTyrePressureType)
}

func (*CodeProtectAreaSectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionType_stagedOrder)
}

func (*CodeProtectAreaSectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeProtectAreaSectionType_orderStaged[order]
}

func (*CodeProtectAreaSectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeProtectAreaSectionTypes_mapString
}

func (*CodeProtectAreaSectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeProtectAreaSectionTypes
}

func (*CodeProtectAreaSectionType) GongNewInstance() any {
	return new(CodeProtectAreaSectionType)
}

func (*CodeRVRReadingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRVRReadingTypes, stage.CodeRVRReadingType_stagedOrder)
}

func (*CodeRVRReadingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRVRReadingType_orderStaged[order]
}

func (*CodeRVRReadingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRVRReadingTypes_mapString
}

func (*CodeRVRReadingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRVRReadingTypes
}

func (*CodeRVRReadingType) GongNewInstance() any {
	return new(CodeRVRReadingType)
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

func (*CodeRoadType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRoadTypes, stage.CodeRoadType_stagedOrder)
}

func (*CodeRoadType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRoadType_orderStaged[order]
}

func (*CodeRoadType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRoadTypes_mapString
}

func (*CodeRoadType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRoadTypes
}

func (*CodeRoadType) GongNewInstance() any {
	return new(CodeRoadType)
}

func (*CodeRunwayElementType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwayElementTypes, stage.CodeRunwayElementType_stagedOrder)
}

func (*CodeRunwayElementType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwayElementType_orderStaged[order]
}

func (*CodeRunwayElementType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwayElementTypes_mapString
}

func (*CodeRunwayElementType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwayElementTypes
}

func (*CodeRunwayElementType) GongNewInstance() any {
	return new(CodeRunwayElementType)
}

func (*CodeRunwayMarkingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingType_stagedOrder)
}

func (*CodeRunwayMarkingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwayMarkingType_orderStaged[order]
}

func (*CodeRunwayMarkingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwayMarkingTypes_mapString
}

func (*CodeRunwayMarkingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwayMarkingTypes
}

func (*CodeRunwayMarkingType) GongNewInstance() any {
	return new(CodeRunwayMarkingType)
}

func (*CodeRunwayPointRoleType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleType_stagedOrder)
}

func (*CodeRunwayPointRoleType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwayPointRoleType_orderStaged[order]
}

func (*CodeRunwayPointRoleType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwayPointRoleTypes_mapString
}

func (*CodeRunwayPointRoleType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwayPointRoleTypes
}

func (*CodeRunwayPointRoleType) GongNewInstance() any {
	return new(CodeRunwayPointRoleType)
}

func (*CodeRunwayProtectionAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaType_stagedOrder)
}

func (*CodeRunwayProtectionAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwayProtectionAreaType_orderStaged[order]
}

func (*CodeRunwayProtectionAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwayProtectionAreaTypes_mapString
}

func (*CodeRunwayProtectionAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwayProtectionAreaTypes
}

func (*CodeRunwayProtectionAreaType) GongNewInstance() any {
	return new(CodeRunwayProtectionAreaType)
}

func (*CodeRunwaySectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionType_stagedOrder)
}

func (*CodeRunwaySectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwaySectionType_orderStaged[order]
}

func (*CodeRunwaySectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwaySectionTypes_mapString
}

func (*CodeRunwaySectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwaySectionTypes
}

func (*CodeRunwaySectionType) GongNewInstance() any {
	return new(CodeRunwaySectionType)
}

func (*CodeRunwayType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeRunwayTypes, stage.CodeRunwayType_stagedOrder)
}

func (*CodeRunwayType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeRunwayType_orderStaged[order]
}

func (*CodeRunwayType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeRunwayTypes_mapString
}

func (*CodeRunwayType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeRunwayTypes
}

func (*CodeRunwayType) GongNewInstance() any {
	return new(CodeRunwayType)
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

func (*CodeStatusAirportType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusAirportTypes, stage.CodeStatusAirportType_stagedOrder)
}

func (*CodeStatusAirportType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusAirportType_orderStaged[order]
}

func (*CodeStatusAirportType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusAirportTypes_mapString
}

func (*CodeStatusAirportType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusAirportTypes
}

func (*CodeStatusAirportType) GongNewInstance() any {
	return new(CodeStatusAirportType)
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

func (*CodeSurfaceCompositionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionType_stagedOrder)
}

func (*CodeSurfaceCompositionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSurfaceCompositionType_orderStaged[order]
}

func (*CodeSurfaceCompositionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSurfaceCompositionTypes_mapString
}

func (*CodeSurfaceCompositionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSurfaceCompositionTypes
}

func (*CodeSurfaceCompositionType) GongNewInstance() any {
	return new(CodeSurfaceCompositionType)
}

func (*CodeSurfaceConditionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionType_stagedOrder)
}

func (*CodeSurfaceConditionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSurfaceConditionType_orderStaged[order]
}

func (*CodeSurfaceConditionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSurfaceConditionTypes_mapString
}

func (*CodeSurfaceConditionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSurfaceConditionTypes
}

func (*CodeSurfaceConditionType) GongNewInstance() any {
	return new(CodeSurfaceConditionType)
}

func (*CodeSurfacePreparationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationType_stagedOrder)
}

func (*CodeSurfacePreparationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeSurfacePreparationType_orderStaged[order]
}

func (*CodeSurfacePreparationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeSurfacePreparationTypes_mapString
}

func (*CodeSurfacePreparationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeSurfacePreparationTypes
}

func (*CodeSurfacePreparationType) GongNewInstance() any {
	return new(CodeSurfacePreparationType)
}

func (*CodeTLOFSectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionType_stagedOrder)
}

func (*CodeTLOFSectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTLOFSectionType_orderStaged[order]
}

func (*CodeTLOFSectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTLOFSectionTypes_mapString
}

func (*CodeTLOFSectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTLOFSectionTypes
}

func (*CodeTLOFSectionType) GongNewInstance() any {
	return new(CodeTLOFSectionType)
}

func (*CodeTaxiwayElementType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementType_stagedOrder)
}

func (*CodeTaxiwayElementType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTaxiwayElementType_orderStaged[order]
}

func (*CodeTaxiwayElementType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTaxiwayElementTypes_mapString
}

func (*CodeTaxiwayElementType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTaxiwayElementTypes
}

func (*CodeTaxiwayElementType) GongNewInstance() any {
	return new(CodeTaxiwayElementType)
}

func (*CodeTaxiwaySectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionType_stagedOrder)
}

func (*CodeTaxiwaySectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTaxiwaySectionType_orderStaged[order]
}

func (*CodeTaxiwaySectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTaxiwaySectionTypes_mapString
}

func (*CodeTaxiwaySectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTaxiwaySectionTypes
}

func (*CodeTaxiwaySectionType) GongNewInstance() any {
	return new(CodeTaxiwaySectionType)
}

func (*CodeTaxiwayType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeTaxiwayTypes, stage.CodeTaxiwayType_stagedOrder)
}

func (*CodeTaxiwayType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeTaxiwayType_orderStaged[order]
}

func (*CodeTaxiwayType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeTaxiwayTypes_mapString
}

func (*CodeTaxiwayType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeTaxiwayTypes
}

func (*CodeTaxiwayType) GongNewInstance() any {
	return new(CodeTaxiwayType)
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

func (*CodeUsageLimitationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationType_stagedOrder)
}

func (*CodeUsageLimitationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeUsageLimitationType_orderStaged[order]
}

func (*CodeUsageLimitationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeUsageLimitationTypes_mapString
}

func (*CodeUsageLimitationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeUsageLimitationTypes
}

func (*CodeUsageLimitationType) GongNewInstance() any {
	return new(CodeUsageLimitationType)
}

func (*CodeVASISType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVASISTypes, stage.CodeVASISType_stagedOrder)
}

func (*CodeVASISType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVASISType_orderStaged[order]
}

func (*CodeVASISType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVASISTypes_mapString
}

func (*CodeVASISType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVASISTypes
}

func (*CodeVASISType) GongNewInstance() any {
	return new(CodeVASISType)
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

func (*CodeVisualDockingGuidanceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceType_stagedOrder)
}

func (*CodeVisualDockingGuidanceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVisualDockingGuidanceType_orderStaged[order]
}

func (*CodeVisualDockingGuidanceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVisualDockingGuidanceTypes_mapString
}

func (*CodeVisualDockingGuidanceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVisualDockingGuidanceTypes
}

func (*CodeVisualDockingGuidanceType) GongNewInstance() any {
	return new(CodeVisualDockingGuidanceType)
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

func (*CodeWorkAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeWorkAreaTypes, stage.CodeWorkAreaType_stagedOrder)
}

func (*CodeWorkAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeWorkAreaType_orderStaged[order]
}

func (*CodeWorkAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeWorkAreaTypes_mapString
}

func (*CodeWorkAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeWorkAreaTypes
}

func (*CodeWorkAreaType) GongNewInstance() any {
	return new(CodeWorkAreaType)
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

func (*DateTimeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DateTimeTypes, stage.DateTimeType_stagedOrder)
}

func (*DateTimeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DateTimeType_orderStaged[order]
}

func (*DateTimeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DateTimeTypes_mapString
}

func (*DateTimeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DateTimeTypes
}

func (*DateTimeType) GongNewInstance() any {
	return new(DateTimeType)
}

func (*DateType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DateTypes, stage.DateType_stagedOrder)
}

func (*DateType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DateType_orderStaged[order]
}

func (*DateType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DateTypes_mapString
}

func (*DateType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DateTypes
}

func (*DateType) GongNewInstance() any {
	return new(DateType)
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

func (*ValDepthType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValDepthTypes, stage.ValDepthType_stagedOrder)
}

func (*ValDepthType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValDepthType_orderStaged[order]
}

func (*ValDepthType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValDepthTypes_mapString
}

func (*ValDepthType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValDepthTypes
}

func (*ValDepthType) GongNewInstance() any {
	return new(ValDepthType)
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

func (*ValFLType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValFLTypes, stage.ValFLType_stagedOrder)
}

func (*ValFLType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValFLType_orderStaged[order]
}

func (*ValFLType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValFLTypes_mapString
}

func (*ValFLType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValFLTypes
}

func (*ValFLType) GongNewInstance() any {
	return new(ValFLType)
}

func (*ValFrictionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValFrictionTypes, stage.ValFrictionType_stagedOrder)
}

func (*ValFrictionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValFrictionType_orderStaged[order]
}

func (*ValFrictionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValFrictionTypes_mapString
}

func (*ValFrictionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValFrictionTypes
}

func (*ValFrictionType) GongNewInstance() any {
	return new(ValFrictionType)
}

func (*ValLCNType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValLCNTypes, stage.ValLCNType_stagedOrder)
}

func (*ValLCNType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValLCNType_orderStaged[order]
}

func (*ValLCNType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValLCNTypes_mapString
}

func (*ValLCNType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValLCNTypes
}

func (*ValLCNType) GongNewInstance() any {
	return new(ValLCNType)
}

func (*ValLightIntensityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValLightIntensityTypes, stage.ValLightIntensityType_stagedOrder)
}

func (*ValLightIntensityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValLightIntensityType_orderStaged[order]
}

func (*ValLightIntensityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValLightIntensityTypes_mapString
}

func (*ValLightIntensityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValLightIntensityTypes
}

func (*ValLightIntensityType) GongNewInstance() any {
	return new(ValLightIntensityType)
}

func (*ValMagneticVariationChangeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeType_stagedOrder)
}

func (*ValMagneticVariationChangeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValMagneticVariationChangeType_orderStaged[order]
}

func (*ValMagneticVariationChangeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValMagneticVariationChangeTypes_mapString
}

func (*ValMagneticVariationChangeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValMagneticVariationChangeTypes
}

func (*ValMagneticVariationChangeType) GongNewInstance() any {
	return new(ValMagneticVariationChangeType)
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

func (*ValPCNType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValPCNTypes, stage.ValPCNType_stagedOrder)
}

func (*ValPCNType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValPCNType_orderStaged[order]
}

func (*ValPCNType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValPCNTypes_mapString
}

func (*ValPCNType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValPCNTypes
}

func (*ValPCNType) GongNewInstance() any {
	return new(ValPCNType)
}

func (*ValPercentType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValPercentTypes, stage.ValPercentType_stagedOrder)
}

func (*ValPercentType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValPercentType_orderStaged[order]
}

func (*ValPercentType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValPercentTypes_mapString
}

func (*ValPercentType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValPercentTypes
}

func (*ValPercentType) GongNewInstance() any {
	return new(ValPercentType)
}

func (*ValPressureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.ValPressureTypes, stage.ValPressureType_stagedOrder)
}

func (*ValPressureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.ValPressureType_orderStaged[order]
}

func (*ValPressureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.ValPressureTypes_mapString
}

func (*ValPressureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.ValPressureTypes
}

func (*ValPressureType) GongNewInstance() any {
	return new(ValPressureType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
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
		CodeAircraftCategoryTypes:           make(map[*CodeAircraftCategoryType]struct{}),
		CodeAircraftCategoryTypes_mapString: make(map[string]*CodeAircraftCategoryType),

		CodeAircraftEngineNumberTypes:           make(map[*CodeAircraftEngineNumberType]struct{}),
		CodeAircraftEngineNumberTypes_mapString: make(map[string]*CodeAircraftEngineNumberType),

		CodeAircraftEngineTypes:           make(map[*CodeAircraftEngineType]struct{}),
		CodeAircraftEngineTypes_mapString: make(map[string]*CodeAircraftEngineType),

		CodeAircraftICAOTypes:           make(map[*CodeAircraftICAOType]struct{}),
		CodeAircraftICAOTypes_mapString: make(map[string]*CodeAircraftICAOType),

		CodeAircraftStandTypes:           make(map[*CodeAircraftStandType]struct{}),
		CodeAircraftStandTypes_mapString: make(map[string]*CodeAircraftStandType),

		CodeAircraftTypes:           make(map[*CodeAircraftType]struct{}),
		CodeAircraftTypes_mapString: make(map[string]*CodeAircraftType),

		CodeAircraftWingspanClassTypes:           make(map[*CodeAircraftWingspanClassType]struct{}),
		CodeAircraftWingspanClassTypes_mapString: make(map[string]*CodeAircraftWingspanClassType),

		CodeAirportHeliportCollocationTypes:           make(map[*CodeAirportHeliportCollocationType]struct{}),
		CodeAirportHeliportCollocationTypes_mapString: make(map[string]*CodeAirportHeliportCollocationType),

		CodeAirportHeliportDesignatorTypes:           make(map[*CodeAirportHeliportDesignatorType]struct{}),
		CodeAirportHeliportDesignatorTypes_mapString: make(map[string]*CodeAirportHeliportDesignatorType),

		CodeAirportHeliportTypes:           make(map[*CodeAirportHeliportType]struct{}),
		CodeAirportHeliportTypes_mapString: make(map[string]*CodeAirportHeliportType),

		CodeAirportWarningTypes:           make(map[*CodeAirportWarningType]struct{}),
		CodeAirportWarningTypes_mapString: make(map[string]*CodeAirportWarningType),

		CodeApproachGuidanceTypes:           make(map[*CodeApproachGuidanceType]struct{}),
		CodeApproachGuidanceTypes_mapString: make(map[string]*CodeApproachGuidanceType),

		CodeApproachLightingICAOTypes:           make(map[*CodeApproachLightingICAOType]struct{}),
		CodeApproachLightingICAOTypes_mapString: make(map[string]*CodeApproachLightingICAOType),

		CodeApproachLightingTypes:           make(map[*CodeApproachLightingType]struct{}),
		CodeApproachLightingTypes_mapString: make(map[string]*CodeApproachLightingType),

		CodeApronElementTypes:           make(map[*CodeApronElementType]struct{}),
		CodeApronElementTypes_mapString: make(map[string]*CodeApronElementType),

		CodeApronSectionTypes:           make(map[*CodeApronSectionType]struct{}),
		CodeApronSectionTypes_mapString: make(map[string]*CodeApronSectionType),

		CodeArrestingGearEnergyAbsorbTypes:           make(map[*CodeArrestingGearEnergyAbsorbType]struct{}),
		CodeArrestingGearEnergyAbsorbTypes_mapString: make(map[string]*CodeArrestingGearEnergyAbsorbType),

		CodeArrestingGearEngageDeviceTypes:           make(map[*CodeArrestingGearEngageDeviceType]struct{}),
		CodeArrestingGearEngageDeviceTypes_mapString: make(map[string]*CodeArrestingGearEngageDeviceType),

		CodeAuthorityRoleTypes:           make(map[*CodeAuthorityRoleType]struct{}),
		CodeAuthorityRoleTypes_mapString: make(map[string]*CodeAuthorityRoleType),

		CodeBuoyDesignatorTypes:           make(map[*CodeBuoyDesignatorType]struct{}),
		CodeBuoyDesignatorTypes_mapString: make(map[string]*CodeBuoyDesignatorType),

		CodeBuoyTypes:           make(map[*CodeBuoyType]struct{}),
		CodeBuoyTypes_mapString: make(map[string]*CodeBuoyType),

		CodeColourTypes:           make(map[*CodeColourType]struct{}),
		CodeColourTypes_mapString: make(map[string]*CodeColourType),

		CodeCommunicationModeTypes:           make(map[*CodeCommunicationModeType]struct{}),
		CodeCommunicationModeTypes_mapString: make(map[string]*CodeCommunicationModeType),

		CodeContaminationTypes:           make(map[*CodeContaminationType]struct{}),
		CodeContaminationTypes_mapString: make(map[string]*CodeContaminationType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeDeclaredDistanceTypes:           make(map[*CodeDeclaredDistanceType]struct{}),
		CodeDeclaredDistanceTypes_mapString: make(map[string]*CodeDeclaredDistanceType),

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

		CodeFlightRuleTypes:           make(map[*CodeFlightRuleType]struct{}),
		CodeFlightRuleTypes_mapString: make(map[string]*CodeFlightRuleType),

		CodeFlightStatusTypes:           make(map[*CodeFlightStatusType]struct{}),
		CodeFlightStatusTypes_mapString: make(map[string]*CodeFlightStatusType),

		CodeFlightTypes:           make(map[*CodeFlightType]struct{}),
		CodeFlightTypes_mapString: make(map[string]*CodeFlightType),

		CodeFrictionDeviceTypes:           make(map[*CodeFrictionDeviceType]struct{}),
		CodeFrictionDeviceTypes_mapString: make(map[string]*CodeFrictionDeviceType),

		CodeFrictionEstimateTypes:           make(map[*CodeFrictionEstimateType]struct{}),
		CodeFrictionEstimateTypes_mapString: make(map[string]*CodeFrictionEstimateType),

		CodeGradeSeparationTypes:           make(map[*CodeGradeSeparationType]struct{}),
		CodeGradeSeparationTypes_mapString: make(map[string]*CodeGradeSeparationType),

		CodeGuidanceLineTypes:           make(map[*CodeGuidanceLineType]struct{}),
		CodeGuidanceLineTypes_mapString: make(map[string]*CodeGuidanceLineType),

		CodeHelicopterPerformanceTypes:           make(map[*CodeHelicopterPerformanceType]struct{}),
		CodeHelicopterPerformanceTypes_mapString: make(map[string]*CodeHelicopterPerformanceType),

		CodeHoldingCategoryTypes:           make(map[*CodeHoldingCategoryType]struct{}),
		CodeHoldingCategoryTypes_mapString: make(map[string]*CodeHoldingCategoryType),

		CodeIATATypes:           make(map[*CodeIATAType]struct{}),
		CodeIATATypes_mapString: make(map[string]*CodeIATAType),

		CodeICAOTypes:           make(map[*CodeICAOType]struct{}),
		CodeICAOTypes_mapString: make(map[string]*CodeICAOType),

		CodeLightHoldingPositionTypes:           make(map[*CodeLightHoldingPositionType]struct{}),
		CodeLightHoldingPositionTypes_mapString: make(map[string]*CodeLightHoldingPositionType),

		CodeLightIntensityTypes:           make(map[*CodeLightIntensityType]struct{}),
		CodeLightIntensityTypes_mapString: make(map[string]*CodeLightIntensityType),

		CodeLightSourceTypes:           make(map[*CodeLightSourceType]struct{}),
		CodeLightSourceTypes_mapString: make(map[string]*CodeLightSourceType),

		CodeLightingJARTypes:           make(map[*CodeLightingJARType]struct{}),
		CodeLightingJARTypes_mapString: make(map[string]*CodeLightingJARType),

		CodeLoadingBridgeTypes:           make(map[*CodeLoadingBridgeType]struct{}),
		CodeLoadingBridgeTypes_mapString: make(map[string]*CodeLoadingBridgeType),

		CodeLogicalOperatorTypes:           make(map[*CodeLogicalOperatorType]struct{}),
		CodeLogicalOperatorTypes_mapString: make(map[string]*CodeLogicalOperatorType),

		CodeMarkingConditionTypes:           make(map[*CodeMarkingConditionType]struct{}),
		CodeMarkingConditionTypes_mapString: make(map[string]*CodeMarkingConditionType),

		CodeMarkingStyleTypes:           make(map[*CodeMarkingStyleType]struct{}),
		CodeMarkingStyleTypes_mapString: make(map[string]*CodeMarkingStyleType),

		CodeMeteoConditionsTypes:           make(map[*CodeMeteoConditionsType]struct{}),
		CodeMeteoConditionsTypes_mapString: make(map[string]*CodeMeteoConditionsType),

		CodeMilitaryOperationsTypes:           make(map[*CodeMilitaryOperationsType]struct{}),
		CodeMilitaryOperationsTypes_mapString: make(map[string]*CodeMilitaryOperationsType),

		CodeMilitaryStatusTypes:           make(map[*CodeMilitaryStatusType]struct{}),
		CodeMilitaryStatusTypes_mapString: make(map[string]*CodeMilitaryStatusType),

		CodeNavigationEquipmentTypes:           make(map[*CodeNavigationEquipmentType]struct{}),
		CodeNavigationEquipmentTypes_mapString: make(map[string]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationTypes:           make(map[*CodeNavigationSpecificationType]struct{}),
		CodeNavigationSpecificationTypes_mapString: make(map[string]*CodeNavigationSpecificationType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodeOperationAirportHeliportTypes:           make(map[*CodeOperationAirportHeliportType]struct{}),
		CodeOperationAirportHeliportTypes_mapString: make(map[string]*CodeOperationAirportHeliportType),

		CodeOperationManoeuvringAreaTypes:           make(map[*CodeOperationManoeuvringAreaType]struct{}),
		CodeOperationManoeuvringAreaTypes_mapString: make(map[string]*CodeOperationManoeuvringAreaType),

		CodePCNMethodTypes:           make(map[*CodePCNMethodType]struct{}),
		CodePCNMethodTypes_mapString: make(map[string]*CodePCNMethodType),

		CodePCNPavementTypes:           make(map[*CodePCNPavementType]struct{}),
		CodePCNPavementTypes_mapString: make(map[string]*CodePCNPavementType),

		CodePCNSubgradeTypes:           make(map[*CodePCNSubgradeType]struct{}),
		CodePCNSubgradeTypes_mapString: make(map[string]*CodePCNSubgradeType),

		CodePCNTyrePressureTypes:           make(map[*CodePCNTyrePressureType]struct{}),
		CodePCNTyrePressureTypes_mapString: make(map[string]*CodePCNTyrePressureType),

		CodeProtectAreaSectionTypes:           make(map[*CodeProtectAreaSectionType]struct{}),
		CodeProtectAreaSectionTypes_mapString: make(map[string]*CodeProtectAreaSectionType),

		CodeRVRReadingTypes:           make(map[*CodeRVRReadingType]struct{}),
		CodeRVRReadingTypes_mapString: make(map[string]*CodeRVRReadingType),

		CodeRVSMTypes:           make(map[*CodeRVSMType]struct{}),
		CodeRVSMTypes_mapString: make(map[string]*CodeRVSMType),

		CodeRoadTypes:           make(map[*CodeRoadType]struct{}),
		CodeRoadTypes_mapString: make(map[string]*CodeRoadType),

		CodeRunwayElementTypes:           make(map[*CodeRunwayElementType]struct{}),
		CodeRunwayElementTypes_mapString: make(map[string]*CodeRunwayElementType),

		CodeRunwayMarkingTypes:           make(map[*CodeRunwayMarkingType]struct{}),
		CodeRunwayMarkingTypes_mapString: make(map[string]*CodeRunwayMarkingType),

		CodeRunwayPointRoleTypes:           make(map[*CodeRunwayPointRoleType]struct{}),
		CodeRunwayPointRoleTypes_mapString: make(map[string]*CodeRunwayPointRoleType),

		CodeRunwayProtectionAreaTypes:           make(map[*CodeRunwayProtectionAreaType]struct{}),
		CodeRunwayProtectionAreaTypes_mapString: make(map[string]*CodeRunwayProtectionAreaType),

		CodeRunwaySectionTypes:           make(map[*CodeRunwaySectionType]struct{}),
		CodeRunwaySectionTypes_mapString: make(map[string]*CodeRunwaySectionType),

		CodeRunwayTypes:           make(map[*CodeRunwayType]struct{}),
		CodeRunwayTypes_mapString: make(map[string]*CodeRunwayType),

		CodeSideTypes:           make(map[*CodeSideType]struct{}),
		CodeSideTypes_mapString: make(map[string]*CodeSideType),

		CodeStatusAirportTypes:           make(map[*CodeStatusAirportType]struct{}),
		CodeStatusAirportTypes_mapString: make(map[string]*CodeStatusAirportType),

		CodeStatusOperationsTypes:           make(map[*CodeStatusOperationsType]struct{}),
		CodeStatusOperationsTypes_mapString: make(map[string]*CodeStatusOperationsType),

		CodeSurfaceCompositionTypes:           make(map[*CodeSurfaceCompositionType]struct{}),
		CodeSurfaceCompositionTypes_mapString: make(map[string]*CodeSurfaceCompositionType),

		CodeSurfaceConditionTypes:           make(map[*CodeSurfaceConditionType]struct{}),
		CodeSurfaceConditionTypes_mapString: make(map[string]*CodeSurfaceConditionType),

		CodeSurfacePreparationTypes:           make(map[*CodeSurfacePreparationType]struct{}),
		CodeSurfacePreparationTypes_mapString: make(map[string]*CodeSurfacePreparationType),

		CodeTLOFSectionTypes:           make(map[*CodeTLOFSectionType]struct{}),
		CodeTLOFSectionTypes_mapString: make(map[string]*CodeTLOFSectionType),

		CodeTaxiwayElementTypes:           make(map[*CodeTaxiwayElementType]struct{}),
		CodeTaxiwayElementTypes_mapString: make(map[string]*CodeTaxiwayElementType),

		CodeTaxiwaySectionTypes:           make(map[*CodeTaxiwaySectionType]struct{}),
		CodeTaxiwaySectionTypes_mapString: make(map[string]*CodeTaxiwaySectionType),

		CodeTaxiwayTypes:           make(map[*CodeTaxiwayType]struct{}),
		CodeTaxiwayTypes_mapString: make(map[string]*CodeTaxiwayType),

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

		CodeUsageLimitationTypes:           make(map[*CodeUsageLimitationType]struct{}),
		CodeUsageLimitationTypes_mapString: make(map[string]*CodeUsageLimitationType),

		CodeVASISTypes:           make(map[*CodeVASISType]struct{}),
		CodeVASISTypes_mapString: make(map[string]*CodeVASISType),

		CodeValueInterpretationTypes:           make(map[*CodeValueInterpretationType]struct{}),
		CodeValueInterpretationTypes_mapString: make(map[string]*CodeValueInterpretationType),

		CodeVerticalDatumTypes:           make(map[*CodeVerticalDatumType]struct{}),
		CodeVerticalDatumTypes_mapString: make(map[string]*CodeVerticalDatumType),

		CodeVisualDockingGuidanceTypes:           make(map[*CodeVisualDockingGuidanceType]struct{}),
		CodeVisualDockingGuidanceTypes_mapString: make(map[string]*CodeVisualDockingGuidanceType),

		CodeWakeTurbulenceTypes:           make(map[*CodeWakeTurbulenceType]struct{}),
		CodeWakeTurbulenceTypes_mapString: make(map[string]*CodeWakeTurbulenceType),

		CodeWorkAreaTypes:           make(map[*CodeWorkAreaType]struct{}),
		CodeWorkAreaTypes_mapString: make(map[string]*CodeWorkAreaType),

		CodeYesNoTypes:           make(map[*CodeYesNoType]struct{}),
		CodeYesNoTypes_mapString: make(map[string]*CodeYesNoType),

		DateMonthDayTypes:           make(map[*DateMonthDayType]struct{}),
		DateMonthDayTypes_mapString: make(map[string]*DateMonthDayType),

		DateTimeTypes:           make(map[*DateTimeType]struct{}),
		DateTimeTypes_mapString: make(map[string]*DateTimeType),

		DateTypes:           make(map[*DateType]struct{}),
		DateTypes_mapString: make(map[string]*DateType),

		DateYearTypes:           make(map[*DateYearType]struct{}),
		DateYearTypes_mapString: make(map[string]*DateYearType),

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

		ValDepthTypes:           make(map[*ValDepthType]struct{}),
		ValDepthTypes_mapString: make(map[string]*ValDepthType),

		ValDistanceSignedTypes:           make(map[*ValDistanceSignedType]struct{}),
		ValDistanceSignedTypes_mapString: make(map[string]*ValDistanceSignedType),

		ValDistanceTypes:           make(map[*ValDistanceType]struct{}),
		ValDistanceTypes_mapString: make(map[string]*ValDistanceType),

		ValDistanceVerticalTypes:           make(map[*ValDistanceVerticalType]struct{}),
		ValDistanceVerticalTypes_mapString: make(map[string]*ValDistanceVerticalType),

		ValDurationTypes:           make(map[*ValDurationType]struct{}),
		ValDurationTypes_mapString: make(map[string]*ValDurationType),

		ValFLTypes:           make(map[*ValFLType]struct{}),
		ValFLTypes_mapString: make(map[string]*ValFLType),

		ValFrictionTypes:           make(map[*ValFrictionType]struct{}),
		ValFrictionTypes_mapString: make(map[string]*ValFrictionType),

		ValLCNTypes:           make(map[*ValLCNType]struct{}),
		ValLCNTypes_mapString: make(map[string]*ValLCNType),

		ValLightIntensityTypes:           make(map[*ValLightIntensityType]struct{}),
		ValLightIntensityTypes_mapString: make(map[string]*ValLightIntensityType),

		ValMagneticVariationChangeTypes:           make(map[*ValMagneticVariationChangeType]struct{}),
		ValMagneticVariationChangeTypes_mapString: make(map[string]*ValMagneticVariationChangeType),

		ValMagneticVariationTypes:           make(map[*ValMagneticVariationType]struct{}),
		ValMagneticVariationTypes_mapString: make(map[string]*ValMagneticVariationType),

		ValPCNTypes:           make(map[*ValPCNType]struct{}),
		ValPCNTypes_mapString: make(map[string]*ValPCNType),

		ValPercentTypes:           make(map[*ValPercentType]struct{}),
		ValPercentTypes_mapString: make(map[string]*ValPercentType),

		ValPressureTypes:           make(map[*ValPressureType]struct{}),
		ValPressureTypes_mapString: make(map[string]*ValPressureType),

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

		CodeAircraftStandType_stagedOrder: make(map[*CodeAircraftStandType]uint),
		CodeAircraftStandType_orderStaged: make(map[uint]*CodeAircraftStandType),
		CodeAircraftStandTypes_reference:  make(map[*CodeAircraftStandType]*CodeAircraftStandType),

		CodeAircraftType_stagedOrder: make(map[*CodeAircraftType]uint),
		CodeAircraftType_orderStaged: make(map[uint]*CodeAircraftType),
		CodeAircraftTypes_reference:  make(map[*CodeAircraftType]*CodeAircraftType),

		CodeAircraftWingspanClassType_stagedOrder: make(map[*CodeAircraftWingspanClassType]uint),
		CodeAircraftWingspanClassType_orderStaged: make(map[uint]*CodeAircraftWingspanClassType),
		CodeAircraftWingspanClassTypes_reference:  make(map[*CodeAircraftWingspanClassType]*CodeAircraftWingspanClassType),

		CodeAirportHeliportCollocationType_stagedOrder: make(map[*CodeAirportHeliportCollocationType]uint),
		CodeAirportHeliportCollocationType_orderStaged: make(map[uint]*CodeAirportHeliportCollocationType),
		CodeAirportHeliportCollocationTypes_reference:  make(map[*CodeAirportHeliportCollocationType]*CodeAirportHeliportCollocationType),

		CodeAirportHeliportDesignatorType_stagedOrder: make(map[*CodeAirportHeliportDesignatorType]uint),
		CodeAirportHeliportDesignatorType_orderStaged: make(map[uint]*CodeAirportHeliportDesignatorType),
		CodeAirportHeliportDesignatorTypes_reference:  make(map[*CodeAirportHeliportDesignatorType]*CodeAirportHeliportDesignatorType),

		CodeAirportHeliportType_stagedOrder: make(map[*CodeAirportHeliportType]uint),
		CodeAirportHeliportType_orderStaged: make(map[uint]*CodeAirportHeliportType),
		CodeAirportHeliportTypes_reference:  make(map[*CodeAirportHeliportType]*CodeAirportHeliportType),

		CodeAirportWarningType_stagedOrder: make(map[*CodeAirportWarningType]uint),
		CodeAirportWarningType_orderStaged: make(map[uint]*CodeAirportWarningType),
		CodeAirportWarningTypes_reference:  make(map[*CodeAirportWarningType]*CodeAirportWarningType),

		CodeApproachGuidanceType_stagedOrder: make(map[*CodeApproachGuidanceType]uint),
		CodeApproachGuidanceType_orderStaged: make(map[uint]*CodeApproachGuidanceType),
		CodeApproachGuidanceTypes_reference:  make(map[*CodeApproachGuidanceType]*CodeApproachGuidanceType),

		CodeApproachLightingICAOType_stagedOrder: make(map[*CodeApproachLightingICAOType]uint),
		CodeApproachLightingICAOType_orderStaged: make(map[uint]*CodeApproachLightingICAOType),
		CodeApproachLightingICAOTypes_reference:  make(map[*CodeApproachLightingICAOType]*CodeApproachLightingICAOType),

		CodeApproachLightingType_stagedOrder: make(map[*CodeApproachLightingType]uint),
		CodeApproachLightingType_orderStaged: make(map[uint]*CodeApproachLightingType),
		CodeApproachLightingTypes_reference:  make(map[*CodeApproachLightingType]*CodeApproachLightingType),

		CodeApronElementType_stagedOrder: make(map[*CodeApronElementType]uint),
		CodeApronElementType_orderStaged: make(map[uint]*CodeApronElementType),
		CodeApronElementTypes_reference:  make(map[*CodeApronElementType]*CodeApronElementType),

		CodeApronSectionType_stagedOrder: make(map[*CodeApronSectionType]uint),
		CodeApronSectionType_orderStaged: make(map[uint]*CodeApronSectionType),
		CodeApronSectionTypes_reference:  make(map[*CodeApronSectionType]*CodeApronSectionType),

		CodeArrestingGearEnergyAbsorbType_stagedOrder: make(map[*CodeArrestingGearEnergyAbsorbType]uint),
		CodeArrestingGearEnergyAbsorbType_orderStaged: make(map[uint]*CodeArrestingGearEnergyAbsorbType),
		CodeArrestingGearEnergyAbsorbTypes_reference:  make(map[*CodeArrestingGearEnergyAbsorbType]*CodeArrestingGearEnergyAbsorbType),

		CodeArrestingGearEngageDeviceType_stagedOrder: make(map[*CodeArrestingGearEngageDeviceType]uint),
		CodeArrestingGearEngageDeviceType_orderStaged: make(map[uint]*CodeArrestingGearEngageDeviceType),
		CodeArrestingGearEngageDeviceTypes_reference:  make(map[*CodeArrestingGearEngageDeviceType]*CodeArrestingGearEngageDeviceType),

		CodeAuthorityRoleType_stagedOrder: make(map[*CodeAuthorityRoleType]uint),
		CodeAuthorityRoleType_orderStaged: make(map[uint]*CodeAuthorityRoleType),
		CodeAuthorityRoleTypes_reference:  make(map[*CodeAuthorityRoleType]*CodeAuthorityRoleType),

		CodeBuoyDesignatorType_stagedOrder: make(map[*CodeBuoyDesignatorType]uint),
		CodeBuoyDesignatorType_orderStaged: make(map[uint]*CodeBuoyDesignatorType),
		CodeBuoyDesignatorTypes_reference:  make(map[*CodeBuoyDesignatorType]*CodeBuoyDesignatorType),

		CodeBuoyType_stagedOrder: make(map[*CodeBuoyType]uint),
		CodeBuoyType_orderStaged: make(map[uint]*CodeBuoyType),
		CodeBuoyTypes_reference:  make(map[*CodeBuoyType]*CodeBuoyType),

		CodeColourType_stagedOrder: make(map[*CodeColourType]uint),
		CodeColourType_orderStaged: make(map[uint]*CodeColourType),
		CodeColourTypes_reference:  make(map[*CodeColourType]*CodeColourType),

		CodeCommunicationModeType_stagedOrder: make(map[*CodeCommunicationModeType]uint),
		CodeCommunicationModeType_orderStaged: make(map[uint]*CodeCommunicationModeType),
		CodeCommunicationModeTypes_reference:  make(map[*CodeCommunicationModeType]*CodeCommunicationModeType),

		CodeContaminationType_stagedOrder: make(map[*CodeContaminationType]uint),
		CodeContaminationType_orderStaged: make(map[uint]*CodeContaminationType),
		CodeContaminationTypes_reference:  make(map[*CodeContaminationType]*CodeContaminationType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeDeclaredDistanceType_stagedOrder: make(map[*CodeDeclaredDistanceType]uint),
		CodeDeclaredDistanceType_orderStaged: make(map[uint]*CodeDeclaredDistanceType),
		CodeDeclaredDistanceTypes_reference:  make(map[*CodeDeclaredDistanceType]*CodeDeclaredDistanceType),

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

		CodeFlightRuleType_stagedOrder: make(map[*CodeFlightRuleType]uint),
		CodeFlightRuleType_orderStaged: make(map[uint]*CodeFlightRuleType),
		CodeFlightRuleTypes_reference:  make(map[*CodeFlightRuleType]*CodeFlightRuleType),

		CodeFlightStatusType_stagedOrder: make(map[*CodeFlightStatusType]uint),
		CodeFlightStatusType_orderStaged: make(map[uint]*CodeFlightStatusType),
		CodeFlightStatusTypes_reference:  make(map[*CodeFlightStatusType]*CodeFlightStatusType),

		CodeFlightType_stagedOrder: make(map[*CodeFlightType]uint),
		CodeFlightType_orderStaged: make(map[uint]*CodeFlightType),
		CodeFlightTypes_reference:  make(map[*CodeFlightType]*CodeFlightType),

		CodeFrictionDeviceType_stagedOrder: make(map[*CodeFrictionDeviceType]uint),
		CodeFrictionDeviceType_orderStaged: make(map[uint]*CodeFrictionDeviceType),
		CodeFrictionDeviceTypes_reference:  make(map[*CodeFrictionDeviceType]*CodeFrictionDeviceType),

		CodeFrictionEstimateType_stagedOrder: make(map[*CodeFrictionEstimateType]uint),
		CodeFrictionEstimateType_orderStaged: make(map[uint]*CodeFrictionEstimateType),
		CodeFrictionEstimateTypes_reference:  make(map[*CodeFrictionEstimateType]*CodeFrictionEstimateType),

		CodeGradeSeparationType_stagedOrder: make(map[*CodeGradeSeparationType]uint),
		CodeGradeSeparationType_orderStaged: make(map[uint]*CodeGradeSeparationType),
		CodeGradeSeparationTypes_reference:  make(map[*CodeGradeSeparationType]*CodeGradeSeparationType),

		CodeGuidanceLineType_stagedOrder: make(map[*CodeGuidanceLineType]uint),
		CodeGuidanceLineType_orderStaged: make(map[uint]*CodeGuidanceLineType),
		CodeGuidanceLineTypes_reference:  make(map[*CodeGuidanceLineType]*CodeGuidanceLineType),

		CodeHelicopterPerformanceType_stagedOrder: make(map[*CodeHelicopterPerformanceType]uint),
		CodeHelicopterPerformanceType_orderStaged: make(map[uint]*CodeHelicopterPerformanceType),
		CodeHelicopterPerformanceTypes_reference:  make(map[*CodeHelicopterPerformanceType]*CodeHelicopterPerformanceType),

		CodeHoldingCategoryType_stagedOrder: make(map[*CodeHoldingCategoryType]uint),
		CodeHoldingCategoryType_orderStaged: make(map[uint]*CodeHoldingCategoryType),
		CodeHoldingCategoryTypes_reference:  make(map[*CodeHoldingCategoryType]*CodeHoldingCategoryType),

		CodeIATAType_stagedOrder: make(map[*CodeIATAType]uint),
		CodeIATAType_orderStaged: make(map[uint]*CodeIATAType),
		CodeIATATypes_reference:  make(map[*CodeIATAType]*CodeIATAType),

		CodeICAOType_stagedOrder: make(map[*CodeICAOType]uint),
		CodeICAOType_orderStaged: make(map[uint]*CodeICAOType),
		CodeICAOTypes_reference:  make(map[*CodeICAOType]*CodeICAOType),

		CodeLightHoldingPositionType_stagedOrder: make(map[*CodeLightHoldingPositionType]uint),
		CodeLightHoldingPositionType_orderStaged: make(map[uint]*CodeLightHoldingPositionType),
		CodeLightHoldingPositionTypes_reference:  make(map[*CodeLightHoldingPositionType]*CodeLightHoldingPositionType),

		CodeLightIntensityType_stagedOrder: make(map[*CodeLightIntensityType]uint),
		CodeLightIntensityType_orderStaged: make(map[uint]*CodeLightIntensityType),
		CodeLightIntensityTypes_reference:  make(map[*CodeLightIntensityType]*CodeLightIntensityType),

		CodeLightSourceType_stagedOrder: make(map[*CodeLightSourceType]uint),
		CodeLightSourceType_orderStaged: make(map[uint]*CodeLightSourceType),
		CodeLightSourceTypes_reference:  make(map[*CodeLightSourceType]*CodeLightSourceType),

		CodeLightingJARType_stagedOrder: make(map[*CodeLightingJARType]uint),
		CodeLightingJARType_orderStaged: make(map[uint]*CodeLightingJARType),
		CodeLightingJARTypes_reference:  make(map[*CodeLightingJARType]*CodeLightingJARType),

		CodeLoadingBridgeType_stagedOrder: make(map[*CodeLoadingBridgeType]uint),
		CodeLoadingBridgeType_orderStaged: make(map[uint]*CodeLoadingBridgeType),
		CodeLoadingBridgeTypes_reference:  make(map[*CodeLoadingBridgeType]*CodeLoadingBridgeType),

		CodeLogicalOperatorType_stagedOrder: make(map[*CodeLogicalOperatorType]uint),
		CodeLogicalOperatorType_orderStaged: make(map[uint]*CodeLogicalOperatorType),
		CodeLogicalOperatorTypes_reference:  make(map[*CodeLogicalOperatorType]*CodeLogicalOperatorType),

		CodeMarkingConditionType_stagedOrder: make(map[*CodeMarkingConditionType]uint),
		CodeMarkingConditionType_orderStaged: make(map[uint]*CodeMarkingConditionType),
		CodeMarkingConditionTypes_reference:  make(map[*CodeMarkingConditionType]*CodeMarkingConditionType),

		CodeMarkingStyleType_stagedOrder: make(map[*CodeMarkingStyleType]uint),
		CodeMarkingStyleType_orderStaged: make(map[uint]*CodeMarkingStyleType),
		CodeMarkingStyleTypes_reference:  make(map[*CodeMarkingStyleType]*CodeMarkingStyleType),

		CodeMeteoConditionsType_stagedOrder: make(map[*CodeMeteoConditionsType]uint),
		CodeMeteoConditionsType_orderStaged: make(map[uint]*CodeMeteoConditionsType),
		CodeMeteoConditionsTypes_reference:  make(map[*CodeMeteoConditionsType]*CodeMeteoConditionsType),

		CodeMilitaryOperationsType_stagedOrder: make(map[*CodeMilitaryOperationsType]uint),
		CodeMilitaryOperationsType_orderStaged: make(map[uint]*CodeMilitaryOperationsType),
		CodeMilitaryOperationsTypes_reference:  make(map[*CodeMilitaryOperationsType]*CodeMilitaryOperationsType),

		CodeMilitaryStatusType_stagedOrder: make(map[*CodeMilitaryStatusType]uint),
		CodeMilitaryStatusType_orderStaged: make(map[uint]*CodeMilitaryStatusType),
		CodeMilitaryStatusTypes_reference:  make(map[*CodeMilitaryStatusType]*CodeMilitaryStatusType),

		CodeNavigationEquipmentType_stagedOrder: make(map[*CodeNavigationEquipmentType]uint),
		CodeNavigationEquipmentType_orderStaged: make(map[uint]*CodeNavigationEquipmentType),
		CodeNavigationEquipmentTypes_reference:  make(map[*CodeNavigationEquipmentType]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationType_stagedOrder: make(map[*CodeNavigationSpecificationType]uint),
		CodeNavigationSpecificationType_orderStaged: make(map[uint]*CodeNavigationSpecificationType),
		CodeNavigationSpecificationTypes_reference:  make(map[*CodeNavigationSpecificationType]*CodeNavigationSpecificationType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodeOperationAirportHeliportType_stagedOrder: make(map[*CodeOperationAirportHeliportType]uint),
		CodeOperationAirportHeliportType_orderStaged: make(map[uint]*CodeOperationAirportHeliportType),
		CodeOperationAirportHeliportTypes_reference:  make(map[*CodeOperationAirportHeliportType]*CodeOperationAirportHeliportType),

		CodeOperationManoeuvringAreaType_stagedOrder: make(map[*CodeOperationManoeuvringAreaType]uint),
		CodeOperationManoeuvringAreaType_orderStaged: make(map[uint]*CodeOperationManoeuvringAreaType),
		CodeOperationManoeuvringAreaTypes_reference:  make(map[*CodeOperationManoeuvringAreaType]*CodeOperationManoeuvringAreaType),

		CodePCNMethodType_stagedOrder: make(map[*CodePCNMethodType]uint),
		CodePCNMethodType_orderStaged: make(map[uint]*CodePCNMethodType),
		CodePCNMethodTypes_reference:  make(map[*CodePCNMethodType]*CodePCNMethodType),

		CodePCNPavementType_stagedOrder: make(map[*CodePCNPavementType]uint),
		CodePCNPavementType_orderStaged: make(map[uint]*CodePCNPavementType),
		CodePCNPavementTypes_reference:  make(map[*CodePCNPavementType]*CodePCNPavementType),

		CodePCNSubgradeType_stagedOrder: make(map[*CodePCNSubgradeType]uint),
		CodePCNSubgradeType_orderStaged: make(map[uint]*CodePCNSubgradeType),
		CodePCNSubgradeTypes_reference:  make(map[*CodePCNSubgradeType]*CodePCNSubgradeType),

		CodePCNTyrePressureType_stagedOrder: make(map[*CodePCNTyrePressureType]uint),
		CodePCNTyrePressureType_orderStaged: make(map[uint]*CodePCNTyrePressureType),
		CodePCNTyrePressureTypes_reference:  make(map[*CodePCNTyrePressureType]*CodePCNTyrePressureType),

		CodeProtectAreaSectionType_stagedOrder: make(map[*CodeProtectAreaSectionType]uint),
		CodeProtectAreaSectionType_orderStaged: make(map[uint]*CodeProtectAreaSectionType),
		CodeProtectAreaSectionTypes_reference:  make(map[*CodeProtectAreaSectionType]*CodeProtectAreaSectionType),

		CodeRVRReadingType_stagedOrder: make(map[*CodeRVRReadingType]uint),
		CodeRVRReadingType_orderStaged: make(map[uint]*CodeRVRReadingType),
		CodeRVRReadingTypes_reference:  make(map[*CodeRVRReadingType]*CodeRVRReadingType),

		CodeRVSMType_stagedOrder: make(map[*CodeRVSMType]uint),
		CodeRVSMType_orderStaged: make(map[uint]*CodeRVSMType),
		CodeRVSMTypes_reference:  make(map[*CodeRVSMType]*CodeRVSMType),

		CodeRoadType_stagedOrder: make(map[*CodeRoadType]uint),
		CodeRoadType_orderStaged: make(map[uint]*CodeRoadType),
		CodeRoadTypes_reference:  make(map[*CodeRoadType]*CodeRoadType),

		CodeRunwayElementType_stagedOrder: make(map[*CodeRunwayElementType]uint),
		CodeRunwayElementType_orderStaged: make(map[uint]*CodeRunwayElementType),
		CodeRunwayElementTypes_reference:  make(map[*CodeRunwayElementType]*CodeRunwayElementType),

		CodeRunwayMarkingType_stagedOrder: make(map[*CodeRunwayMarkingType]uint),
		CodeRunwayMarkingType_orderStaged: make(map[uint]*CodeRunwayMarkingType),
		CodeRunwayMarkingTypes_reference:  make(map[*CodeRunwayMarkingType]*CodeRunwayMarkingType),

		CodeRunwayPointRoleType_stagedOrder: make(map[*CodeRunwayPointRoleType]uint),
		CodeRunwayPointRoleType_orderStaged: make(map[uint]*CodeRunwayPointRoleType),
		CodeRunwayPointRoleTypes_reference:  make(map[*CodeRunwayPointRoleType]*CodeRunwayPointRoleType),

		CodeRunwayProtectionAreaType_stagedOrder: make(map[*CodeRunwayProtectionAreaType]uint),
		CodeRunwayProtectionAreaType_orderStaged: make(map[uint]*CodeRunwayProtectionAreaType),
		CodeRunwayProtectionAreaTypes_reference:  make(map[*CodeRunwayProtectionAreaType]*CodeRunwayProtectionAreaType),

		CodeRunwaySectionType_stagedOrder: make(map[*CodeRunwaySectionType]uint),
		CodeRunwaySectionType_orderStaged: make(map[uint]*CodeRunwaySectionType),
		CodeRunwaySectionTypes_reference:  make(map[*CodeRunwaySectionType]*CodeRunwaySectionType),

		CodeRunwayType_stagedOrder: make(map[*CodeRunwayType]uint),
		CodeRunwayType_orderStaged: make(map[uint]*CodeRunwayType),
		CodeRunwayTypes_reference:  make(map[*CodeRunwayType]*CodeRunwayType),

		CodeSideType_stagedOrder: make(map[*CodeSideType]uint),
		CodeSideType_orderStaged: make(map[uint]*CodeSideType),
		CodeSideTypes_reference:  make(map[*CodeSideType]*CodeSideType),

		CodeStatusAirportType_stagedOrder: make(map[*CodeStatusAirportType]uint),
		CodeStatusAirportType_orderStaged: make(map[uint]*CodeStatusAirportType),
		CodeStatusAirportTypes_reference:  make(map[*CodeStatusAirportType]*CodeStatusAirportType),

		CodeStatusOperationsType_stagedOrder: make(map[*CodeStatusOperationsType]uint),
		CodeStatusOperationsType_orderStaged: make(map[uint]*CodeStatusOperationsType),
		CodeStatusOperationsTypes_reference:  make(map[*CodeStatusOperationsType]*CodeStatusOperationsType),

		CodeSurfaceCompositionType_stagedOrder: make(map[*CodeSurfaceCompositionType]uint),
		CodeSurfaceCompositionType_orderStaged: make(map[uint]*CodeSurfaceCompositionType),
		CodeSurfaceCompositionTypes_reference:  make(map[*CodeSurfaceCompositionType]*CodeSurfaceCompositionType),

		CodeSurfaceConditionType_stagedOrder: make(map[*CodeSurfaceConditionType]uint),
		CodeSurfaceConditionType_orderStaged: make(map[uint]*CodeSurfaceConditionType),
		CodeSurfaceConditionTypes_reference:  make(map[*CodeSurfaceConditionType]*CodeSurfaceConditionType),

		CodeSurfacePreparationType_stagedOrder: make(map[*CodeSurfacePreparationType]uint),
		CodeSurfacePreparationType_orderStaged: make(map[uint]*CodeSurfacePreparationType),
		CodeSurfacePreparationTypes_reference:  make(map[*CodeSurfacePreparationType]*CodeSurfacePreparationType),

		CodeTLOFSectionType_stagedOrder: make(map[*CodeTLOFSectionType]uint),
		CodeTLOFSectionType_orderStaged: make(map[uint]*CodeTLOFSectionType),
		CodeTLOFSectionTypes_reference:  make(map[*CodeTLOFSectionType]*CodeTLOFSectionType),

		CodeTaxiwayElementType_stagedOrder: make(map[*CodeTaxiwayElementType]uint),
		CodeTaxiwayElementType_orderStaged: make(map[uint]*CodeTaxiwayElementType),
		CodeTaxiwayElementTypes_reference:  make(map[*CodeTaxiwayElementType]*CodeTaxiwayElementType),

		CodeTaxiwaySectionType_stagedOrder: make(map[*CodeTaxiwaySectionType]uint),
		CodeTaxiwaySectionType_orderStaged: make(map[uint]*CodeTaxiwaySectionType),
		CodeTaxiwaySectionTypes_reference:  make(map[*CodeTaxiwaySectionType]*CodeTaxiwaySectionType),

		CodeTaxiwayType_stagedOrder: make(map[*CodeTaxiwayType]uint),
		CodeTaxiwayType_orderStaged: make(map[uint]*CodeTaxiwayType),
		CodeTaxiwayTypes_reference:  make(map[*CodeTaxiwayType]*CodeTaxiwayType),

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

		CodeUsageLimitationType_stagedOrder: make(map[*CodeUsageLimitationType]uint),
		CodeUsageLimitationType_orderStaged: make(map[uint]*CodeUsageLimitationType),
		CodeUsageLimitationTypes_reference:  make(map[*CodeUsageLimitationType]*CodeUsageLimitationType),

		CodeVASISType_stagedOrder: make(map[*CodeVASISType]uint),
		CodeVASISType_orderStaged: make(map[uint]*CodeVASISType),
		CodeVASISTypes_reference:  make(map[*CodeVASISType]*CodeVASISType),

		CodeValueInterpretationType_stagedOrder: make(map[*CodeValueInterpretationType]uint),
		CodeValueInterpretationType_orderStaged: make(map[uint]*CodeValueInterpretationType),
		CodeValueInterpretationTypes_reference:  make(map[*CodeValueInterpretationType]*CodeValueInterpretationType),

		CodeVerticalDatumType_stagedOrder: make(map[*CodeVerticalDatumType]uint),
		CodeVerticalDatumType_orderStaged: make(map[uint]*CodeVerticalDatumType),
		CodeVerticalDatumTypes_reference:  make(map[*CodeVerticalDatumType]*CodeVerticalDatumType),

		CodeVisualDockingGuidanceType_stagedOrder: make(map[*CodeVisualDockingGuidanceType]uint),
		CodeVisualDockingGuidanceType_orderStaged: make(map[uint]*CodeVisualDockingGuidanceType),
		CodeVisualDockingGuidanceTypes_reference:  make(map[*CodeVisualDockingGuidanceType]*CodeVisualDockingGuidanceType),

		CodeWakeTurbulenceType_stagedOrder: make(map[*CodeWakeTurbulenceType]uint),
		CodeWakeTurbulenceType_orderStaged: make(map[uint]*CodeWakeTurbulenceType),
		CodeWakeTurbulenceTypes_reference:  make(map[*CodeWakeTurbulenceType]*CodeWakeTurbulenceType),

		CodeWorkAreaType_stagedOrder: make(map[*CodeWorkAreaType]uint),
		CodeWorkAreaType_orderStaged: make(map[uint]*CodeWorkAreaType),
		CodeWorkAreaTypes_reference:  make(map[*CodeWorkAreaType]*CodeWorkAreaType),

		CodeYesNoType_stagedOrder: make(map[*CodeYesNoType]uint),
		CodeYesNoType_orderStaged: make(map[uint]*CodeYesNoType),
		CodeYesNoTypes_reference:  make(map[*CodeYesNoType]*CodeYesNoType),

		DateMonthDayType_stagedOrder: make(map[*DateMonthDayType]uint),
		DateMonthDayType_orderStaged: make(map[uint]*DateMonthDayType),
		DateMonthDayTypes_reference:  make(map[*DateMonthDayType]*DateMonthDayType),

		DateTimeType_stagedOrder: make(map[*DateTimeType]uint),
		DateTimeType_orderStaged: make(map[uint]*DateTimeType),
		DateTimeTypes_reference:  make(map[*DateTimeType]*DateTimeType),

		DateType_stagedOrder: make(map[*DateType]uint),
		DateType_orderStaged: make(map[uint]*DateType),
		DateTypes_reference:  make(map[*DateType]*DateType),

		DateYearType_stagedOrder: make(map[*DateYearType]uint),
		DateYearType_orderStaged: make(map[uint]*DateYearType),
		DateYearTypes_reference:  make(map[*DateYearType]*DateYearType),

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

		ValDepthType_stagedOrder: make(map[*ValDepthType]uint),
		ValDepthType_orderStaged: make(map[uint]*ValDepthType),
		ValDepthTypes_reference:  make(map[*ValDepthType]*ValDepthType),

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

		ValFLType_stagedOrder: make(map[*ValFLType]uint),
		ValFLType_orderStaged: make(map[uint]*ValFLType),
		ValFLTypes_reference:  make(map[*ValFLType]*ValFLType),

		ValFrictionType_stagedOrder: make(map[*ValFrictionType]uint),
		ValFrictionType_orderStaged: make(map[uint]*ValFrictionType),
		ValFrictionTypes_reference:  make(map[*ValFrictionType]*ValFrictionType),

		ValLCNType_stagedOrder: make(map[*ValLCNType]uint),
		ValLCNType_orderStaged: make(map[uint]*ValLCNType),
		ValLCNTypes_reference:  make(map[*ValLCNType]*ValLCNType),

		ValLightIntensityType_stagedOrder: make(map[*ValLightIntensityType]uint),
		ValLightIntensityType_orderStaged: make(map[uint]*ValLightIntensityType),
		ValLightIntensityTypes_reference:  make(map[*ValLightIntensityType]*ValLightIntensityType),

		ValMagneticVariationChangeType_stagedOrder: make(map[*ValMagneticVariationChangeType]uint),
		ValMagneticVariationChangeType_orderStaged: make(map[uint]*ValMagneticVariationChangeType),
		ValMagneticVariationChangeTypes_reference:  make(map[*ValMagneticVariationChangeType]*ValMagneticVariationChangeType),

		ValMagneticVariationType_stagedOrder: make(map[*ValMagneticVariationType]uint),
		ValMagneticVariationType_orderStaged: make(map[uint]*ValMagneticVariationType),
		ValMagneticVariationTypes_reference:  make(map[*ValMagneticVariationType]*ValMagneticVariationType),

		ValPCNType_stagedOrder: make(map[*ValPCNType]uint),
		ValPCNType_orderStaged: make(map[uint]*ValPCNType),
		ValPCNTypes_reference:  make(map[*ValPCNType]*ValPCNType),

		ValPercentType_stagedOrder: make(map[*ValPercentType]uint),
		ValPercentType_orderStaged: make(map[uint]*ValPercentType),
		ValPercentTypes_reference:  make(map[*ValPercentType]*ValPercentType),

		ValPressureType_stagedOrder: make(map[*ValPressureType]uint),
		ValPressureType_orderStaged: make(map[uint]*ValPressureType),
		ValPressureTypes_reference:  make(map[*ValPressureType]*ValPressureType),

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
			"CodeAircraftCategoryType": &CodeAircraftCategoryTypeUnmarshaller{},

			"CodeAircraftEngineNumberType": &CodeAircraftEngineNumberTypeUnmarshaller{},

			"CodeAircraftEngineType": &CodeAircraftEngineTypeUnmarshaller{},

			"CodeAircraftICAOType": &CodeAircraftICAOTypeUnmarshaller{},

			"CodeAircraftStandType": &CodeAircraftStandTypeUnmarshaller{},

			"CodeAircraftType": &CodeAircraftTypeUnmarshaller{},

			"CodeAircraftWingspanClassType": &CodeAircraftWingspanClassTypeUnmarshaller{},

			"CodeAirportHeliportCollocationType": &CodeAirportHeliportCollocationTypeUnmarshaller{},

			"CodeAirportHeliportDesignatorType": &CodeAirportHeliportDesignatorTypeUnmarshaller{},

			"CodeAirportHeliportType": &CodeAirportHeliportTypeUnmarshaller{},

			"CodeAirportWarningType": &CodeAirportWarningTypeUnmarshaller{},

			"CodeApproachGuidanceType": &CodeApproachGuidanceTypeUnmarshaller{},

			"CodeApproachLightingICAOType": &CodeApproachLightingICAOTypeUnmarshaller{},

			"CodeApproachLightingType": &CodeApproachLightingTypeUnmarshaller{},

			"CodeApronElementType": &CodeApronElementTypeUnmarshaller{},

			"CodeApronSectionType": &CodeApronSectionTypeUnmarshaller{},

			"CodeArrestingGearEnergyAbsorbType": &CodeArrestingGearEnergyAbsorbTypeUnmarshaller{},

			"CodeArrestingGearEngageDeviceType": &CodeArrestingGearEngageDeviceTypeUnmarshaller{},

			"CodeAuthorityRoleType": &CodeAuthorityRoleTypeUnmarshaller{},

			"CodeBuoyDesignatorType": &CodeBuoyDesignatorTypeUnmarshaller{},

			"CodeBuoyType": &CodeBuoyTypeUnmarshaller{},

			"CodeColourType": &CodeColourTypeUnmarshaller{},

			"CodeCommunicationModeType": &CodeCommunicationModeTypeUnmarshaller{},

			"CodeContaminationType": &CodeContaminationTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeDeclaredDistanceType": &CodeDeclaredDistanceTypeUnmarshaller{},

			"CodeDirectionTurnType": &CodeDirectionTurnTypeUnmarshaller{},

			"CodeDirectionType": &CodeDirectionTypeUnmarshaller{},

			"CodeEquipmentAntiCollisionType": &CodeEquipmentAntiCollisionTypeUnmarshaller{},

			"CodeFlightOriginType": &CodeFlightOriginTypeUnmarshaller{},

			"CodeFlightPurposeType": &CodeFlightPurposeTypeUnmarshaller{},

			"CodeFlightRuleType": &CodeFlightRuleTypeUnmarshaller{},

			"CodeFlightStatusType": &CodeFlightStatusTypeUnmarshaller{},

			"CodeFlightType": &CodeFlightTypeUnmarshaller{},

			"CodeFrictionDeviceType": &CodeFrictionDeviceTypeUnmarshaller{},

			"CodeFrictionEstimateType": &CodeFrictionEstimateTypeUnmarshaller{},

			"CodeGradeSeparationType": &CodeGradeSeparationTypeUnmarshaller{},

			"CodeGuidanceLineType": &CodeGuidanceLineTypeUnmarshaller{},

			"CodeHelicopterPerformanceType": &CodeHelicopterPerformanceTypeUnmarshaller{},

			"CodeHoldingCategoryType": &CodeHoldingCategoryTypeUnmarshaller{},

			"CodeIATAType": &CodeIATATypeUnmarshaller{},

			"CodeICAOType": &CodeICAOTypeUnmarshaller{},

			"CodeLightHoldingPositionType": &CodeLightHoldingPositionTypeUnmarshaller{},

			"CodeLightIntensityType": &CodeLightIntensityTypeUnmarshaller{},

			"CodeLightSourceType": &CodeLightSourceTypeUnmarshaller{},

			"CodeLightingJARType": &CodeLightingJARTypeUnmarshaller{},

			"CodeLoadingBridgeType": &CodeLoadingBridgeTypeUnmarshaller{},

			"CodeLogicalOperatorType": &CodeLogicalOperatorTypeUnmarshaller{},

			"CodeMarkingConditionType": &CodeMarkingConditionTypeUnmarshaller{},

			"CodeMarkingStyleType": &CodeMarkingStyleTypeUnmarshaller{},

			"CodeMeteoConditionsType": &CodeMeteoConditionsTypeUnmarshaller{},

			"CodeMilitaryOperationsType": &CodeMilitaryOperationsTypeUnmarshaller{},

			"CodeMilitaryStatusType": &CodeMilitaryStatusTypeUnmarshaller{},

			"CodeNavigationEquipmentType": &CodeNavigationEquipmentTypeUnmarshaller{},

			"CodeNavigationSpecificationType": &CodeNavigationSpecificationTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeOperationAirportHeliportType": &CodeOperationAirportHeliportTypeUnmarshaller{},

			"CodeOperationManoeuvringAreaType": &CodeOperationManoeuvringAreaTypeUnmarshaller{},

			"CodePCNMethodType": &CodePCNMethodTypeUnmarshaller{},

			"CodePCNPavementType": &CodePCNPavementTypeUnmarshaller{},

			"CodePCNSubgradeType": &CodePCNSubgradeTypeUnmarshaller{},

			"CodePCNTyrePressureType": &CodePCNTyrePressureTypeUnmarshaller{},

			"CodeProtectAreaSectionType": &CodeProtectAreaSectionTypeUnmarshaller{},

			"CodeRVRReadingType": &CodeRVRReadingTypeUnmarshaller{},

			"CodeRVSMType": &CodeRVSMTypeUnmarshaller{},

			"CodeRoadType": &CodeRoadTypeUnmarshaller{},

			"CodeRunwayElementType": &CodeRunwayElementTypeUnmarshaller{},

			"CodeRunwayMarkingType": &CodeRunwayMarkingTypeUnmarshaller{},

			"CodeRunwayPointRoleType": &CodeRunwayPointRoleTypeUnmarshaller{},

			"CodeRunwayProtectionAreaType": &CodeRunwayProtectionAreaTypeUnmarshaller{},

			"CodeRunwaySectionType": &CodeRunwaySectionTypeUnmarshaller{},

			"CodeRunwayType": &CodeRunwayTypeUnmarshaller{},

			"CodeSideType": &CodeSideTypeUnmarshaller{},

			"CodeStatusAirportType": &CodeStatusAirportTypeUnmarshaller{},

			"CodeStatusOperationsType": &CodeStatusOperationsTypeUnmarshaller{},

			"CodeSurfaceCompositionType": &CodeSurfaceCompositionTypeUnmarshaller{},

			"CodeSurfaceConditionType": &CodeSurfaceConditionTypeUnmarshaller{},

			"CodeSurfacePreparationType": &CodeSurfacePreparationTypeUnmarshaller{},

			"CodeTLOFSectionType": &CodeTLOFSectionTypeUnmarshaller{},

			"CodeTaxiwayElementType": &CodeTaxiwayElementTypeUnmarshaller{},

			"CodeTaxiwaySectionType": &CodeTaxiwaySectionTypeUnmarshaller{},

			"CodeTaxiwayType": &CodeTaxiwayTypeUnmarshaller{},

			"CodeTelecomNetworkType": &CodeTelecomNetworkTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeTransponderType": &CodeTransponderTypeUnmarshaller{},

			"CodeUsageLimitationType": &CodeUsageLimitationTypeUnmarshaller{},

			"CodeVASISType": &CodeVASISTypeUnmarshaller{},

			"CodeValueInterpretationType": &CodeValueInterpretationTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeVisualDockingGuidanceType": &CodeVisualDockingGuidanceTypeUnmarshaller{},

			"CodeWakeTurbulenceType": &CodeWakeTurbulenceTypeUnmarshaller{},

			"CodeWorkAreaType": &CodeWorkAreaTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"DateTimeType": &DateTimeTypeUnmarshaller{},

			"DateType": &DateTypeUnmarshaller{},

			"DateYearType": &DateYearTypeUnmarshaller{},

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

			"ValDepthType": &ValDepthTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValFLType": &ValFLTypeUnmarshaller{},

			"ValFrictionType": &ValFrictionTypeUnmarshaller{},

			"ValLCNType": &ValLCNTypeUnmarshaller{},

			"ValLightIntensityType": &ValLightIntensityTypeUnmarshaller{},

			"ValMagneticVariationChangeType": &ValMagneticVariationChangeTypeUnmarshaller{},

			"ValMagneticVariationType": &ValMagneticVariationTypeUnmarshaller{},

			"ValPCNType": &ValPCNTypeUnmarshaller{},

			"ValPercentType": &ValPercentTypeUnmarshaller{},

			"ValPressureType": &ValPressureTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeAircraftCategoryType"] = len(stage.CodeAircraftCategoryTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineNumberType"] = len(stage.CodeAircraftEngineNumberTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftEngineType"] = len(stage.CodeAircraftEngineTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftICAOType"] = len(stage.CodeAircraftICAOTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftStandType"] = len(stage.CodeAircraftStandTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftType"] = len(stage.CodeAircraftTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftWingspanClassType"] = len(stage.CodeAircraftWingspanClassTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirportHeliportCollocationType"] = len(stage.CodeAirportHeliportCollocationTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirportHeliportDesignatorType"] = len(stage.CodeAirportHeliportDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirportHeliportType"] = len(stage.CodeAirportHeliportTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirportWarningType"] = len(stage.CodeAirportWarningTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachGuidanceType"] = len(stage.CodeApproachGuidanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachLightingICAOType"] = len(stage.CodeApproachLightingICAOTypes)
	stage.Map_GongStructName_InstancesNb["CodeApproachLightingType"] = len(stage.CodeApproachLightingTypes)
	stage.Map_GongStructName_InstancesNb["CodeApronElementType"] = len(stage.CodeApronElementTypes)
	stage.Map_GongStructName_InstancesNb["CodeApronSectionType"] = len(stage.CodeApronSectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeArrestingGearEnergyAbsorbType"] = len(stage.CodeArrestingGearEnergyAbsorbTypes)
	stage.Map_GongStructName_InstancesNb["CodeArrestingGearEngageDeviceType"] = len(stage.CodeArrestingGearEngageDeviceTypes)
	stage.Map_GongStructName_InstancesNb["CodeAuthorityRoleType"] = len(stage.CodeAuthorityRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeBuoyDesignatorType"] = len(stage.CodeBuoyDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeBuoyType"] = len(stage.CodeBuoyTypes)
	stage.Map_GongStructName_InstancesNb["CodeColourType"] = len(stage.CodeColourTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationModeType"] = len(stage.CodeCommunicationModeTypes)
	stage.Map_GongStructName_InstancesNb["CodeContaminationType"] = len(stage.CodeContaminationTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeDeclaredDistanceType"] = len(stage.CodeDeclaredDistanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionTurnType"] = len(stage.CodeDirectionTurnTypes)
	stage.Map_GongStructName_InstancesNb["CodeDirectionType"] = len(stage.CodeDirectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeEquipmentAntiCollisionType"] = len(stage.CodeEquipmentAntiCollisionTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightOriginType"] = len(stage.CodeFlightOriginTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightPurposeType"] = len(stage.CodeFlightPurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightRuleType"] = len(stage.CodeFlightRuleTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightStatusType"] = len(stage.CodeFlightStatusTypes)
	stage.Map_GongStructName_InstancesNb["CodeFlightType"] = len(stage.CodeFlightTypes)
	stage.Map_GongStructName_InstancesNb["CodeFrictionDeviceType"] = len(stage.CodeFrictionDeviceTypes)
	stage.Map_GongStructName_InstancesNb["CodeFrictionEstimateType"] = len(stage.CodeFrictionEstimateTypes)
	stage.Map_GongStructName_InstancesNb["CodeGradeSeparationType"] = len(stage.CodeGradeSeparationTypes)
	stage.Map_GongStructName_InstancesNb["CodeGuidanceLineType"] = len(stage.CodeGuidanceLineTypes)
	stage.Map_GongStructName_InstancesNb["CodeHelicopterPerformanceType"] = len(stage.CodeHelicopterPerformanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeHoldingCategoryType"] = len(stage.CodeHoldingCategoryTypes)
	stage.Map_GongStructName_InstancesNb["CodeIATAType"] = len(stage.CodeIATATypes)
	stage.Map_GongStructName_InstancesNb["CodeICAOType"] = len(stage.CodeICAOTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightHoldingPositionType"] = len(stage.CodeLightHoldingPositionTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightIntensityType"] = len(stage.CodeLightIntensityTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightSourceType"] = len(stage.CodeLightSourceTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightingJARType"] = len(stage.CodeLightingJARTypes)
	stage.Map_GongStructName_InstancesNb["CodeLoadingBridgeType"] = len(stage.CodeLoadingBridgeTypes)
	stage.Map_GongStructName_InstancesNb["CodeLogicalOperatorType"] = len(stage.CodeLogicalOperatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeMarkingConditionType"] = len(stage.CodeMarkingConditionTypes)
	stage.Map_GongStructName_InstancesNb["CodeMarkingStyleType"] = len(stage.CodeMarkingStyleTypes)
	stage.Map_GongStructName_InstancesNb["CodeMeteoConditionsType"] = len(stage.CodeMeteoConditionsTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryOperationsType"] = len(stage.CodeMilitaryOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryStatusType"] = len(stage.CodeMilitaryStatusTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationEquipmentType"] = len(stage.CodeNavigationEquipmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationSpecificationType"] = len(stage.CodeNavigationSpecificationTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeOperationAirportHeliportType"] = len(stage.CodeOperationAirportHeliportTypes)
	stage.Map_GongStructName_InstancesNb["CodeOperationManoeuvringAreaType"] = len(stage.CodeOperationManoeuvringAreaTypes)
	stage.Map_GongStructName_InstancesNb["CodePCNMethodType"] = len(stage.CodePCNMethodTypes)
	stage.Map_GongStructName_InstancesNb["CodePCNPavementType"] = len(stage.CodePCNPavementTypes)
	stage.Map_GongStructName_InstancesNb["CodePCNSubgradeType"] = len(stage.CodePCNSubgradeTypes)
	stage.Map_GongStructName_InstancesNb["CodePCNTyrePressureType"] = len(stage.CodePCNTyrePressureTypes)
	stage.Map_GongStructName_InstancesNb["CodeProtectAreaSectionType"] = len(stage.CodeProtectAreaSectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVRReadingType"] = len(stage.CodeRVRReadingTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMType"] = len(stage.CodeRVSMTypes)
	stage.Map_GongStructName_InstancesNb["CodeRoadType"] = len(stage.CodeRoadTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwayElementType"] = len(stage.CodeRunwayElementTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwayMarkingType"] = len(stage.CodeRunwayMarkingTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwayPointRoleType"] = len(stage.CodeRunwayPointRoleTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwayProtectionAreaType"] = len(stage.CodeRunwayProtectionAreaTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwaySectionType"] = len(stage.CodeRunwaySectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeRunwayType"] = len(stage.CodeRunwayTypes)
	stage.Map_GongStructName_InstancesNb["CodeSideType"] = len(stage.CodeSideTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusAirportType"] = len(stage.CodeStatusAirportTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusOperationsType"] = len(stage.CodeStatusOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeSurfaceCompositionType"] = len(stage.CodeSurfaceCompositionTypes)
	stage.Map_GongStructName_InstancesNb["CodeSurfaceConditionType"] = len(stage.CodeSurfaceConditionTypes)
	stage.Map_GongStructName_InstancesNb["CodeSurfacePreparationType"] = len(stage.CodeSurfacePreparationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTLOFSectionType"] = len(stage.CodeTLOFSectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeTaxiwayElementType"] = len(stage.CodeTaxiwayElementTypes)
	stage.Map_GongStructName_InstancesNb["CodeTaxiwaySectionType"] = len(stage.CodeTaxiwaySectionTypes)
	stage.Map_GongStructName_InstancesNb["CodeTaxiwayType"] = len(stage.CodeTaxiwayTypes)
	stage.Map_GongStructName_InstancesNb["CodeTelecomNetworkType"] = len(stage.CodeTelecomNetworkTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTransponderType"] = len(stage.CodeTransponderTypes)
	stage.Map_GongStructName_InstancesNb["CodeUsageLimitationType"] = len(stage.CodeUsageLimitationTypes)
	stage.Map_GongStructName_InstancesNb["CodeVASISType"] = len(stage.CodeVASISTypes)
	stage.Map_GongStructName_InstancesNb["CodeValueInterpretationType"] = len(stage.CodeValueInterpretationTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeVisualDockingGuidanceType"] = len(stage.CodeVisualDockingGuidanceTypes)
	stage.Map_GongStructName_InstancesNb["CodeWakeTurbulenceType"] = len(stage.CodeWakeTurbulenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeWorkAreaType"] = len(stage.CodeWorkAreaTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["DateTimeType"] = len(stage.DateTimeTypes)
	stage.Map_GongStructName_InstancesNb["DateType"] = len(stage.DateTypes)
	stage.Map_GongStructName_InstancesNb["DateYearType"] = len(stage.DateYearTypes)
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
	stage.Map_GongStructName_InstancesNb["ValDepthType"] = len(stage.ValDepthTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValFLType"] = len(stage.ValFLTypes)
	stage.Map_GongStructName_InstancesNb["ValFrictionType"] = len(stage.ValFrictionTypes)
	stage.Map_GongStructName_InstancesNb["ValLCNType"] = len(stage.ValLCNTypes)
	stage.Map_GongStructName_InstancesNb["ValLightIntensityType"] = len(stage.ValLightIntensityTypes)
	stage.Map_GongStructName_InstancesNb["ValMagneticVariationChangeType"] = len(stage.ValMagneticVariationChangeTypes)
	stage.Map_GongStructName_InstancesNb["ValMagneticVariationType"] = len(stage.ValMagneticVariationTypes)
	stage.Map_GongStructName_InstancesNb["ValPCNType"] = len(stage.ValPCNTypes)
	stage.Map_GongStructName_InstancesNb["ValPercentType"] = len(stage.ValPercentTypes)
	stage.Map_GongStructName_InstancesNb["ValPressureType"] = len(stage.ValPressureTypes)
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

// Stage puts codeaircraftstandtype to the model stage
func (codeaircraftstandtype *CodeAircraftStandType) Stage(stage *Stage) *CodeAircraftStandType {
	__gong__stage(stage.CodeAircraftStandTypes, stage.CodeAircraftStandType_stagedOrder, stage.CodeAircraftStandType_orderStaged, &stage.CodeAircraftStandTypeOrder, stage.CodeAircraftStandTypes_mapString, codeaircraftstandtype, codeaircraftstandtype.Name)
	return codeaircraftstandtype
}

// StagePreserveOrder puts codeaircraftstandtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAircraftStandTypeOrder
// - update stage.CodeAircraftStandTypeOrder accordingly
func (codeaircraftstandtype *CodeAircraftStandType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAircraftStandTypes, stage.CodeAircraftStandType_stagedOrder, stage.CodeAircraftStandType_orderStaged, &stage.CodeAircraftStandTypeOrder, stage.CodeAircraftStandTypes_mapString, codeaircraftstandtype, order, codeaircraftstandtype.Name)
}

// Unstage removes codeaircraftstandtype off the model stage
func (codeaircraftstandtype *CodeAircraftStandType) Unstage(stage *Stage) *CodeAircraftStandType {
	__gong__unstage(stage.CodeAircraftStandTypes, stage.CodeAircraftStandTypes_mapString, codeaircraftstandtype, codeaircraftstandtype.Name)
	return codeaircraftstandtype
}

// UnstageVoid removes codeaircraftstandtype off the model stage
func (codeaircraftstandtype *CodeAircraftStandType) UnstageVoid(stage *Stage) {
	codeaircraftstandtype.Unstage(stage)
}

func (codeaircraftstandtype *CodeAircraftStandType) StageVoid(stage *Stage) {
	codeaircraftstandtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeaircraftstandtype *CodeAircraftStandType) GetName() (res string) {
	return codeaircraftstandtype.Name
}

// for satisfaction of GongStruct interface
func (codeaircraftstandtype *CodeAircraftStandType) SetName(name string) {
	codeaircraftstandtype.Name = name
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

// Stage puts codeairportheliportcollocationtype to the model stage
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) Stage(stage *Stage) *CodeAirportHeliportCollocationType {
	__gong__stage(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationType_stagedOrder, stage.CodeAirportHeliportCollocationType_orderStaged, &stage.CodeAirportHeliportCollocationTypeOrder, stage.CodeAirportHeliportCollocationTypes_mapString, codeairportheliportcollocationtype, codeairportheliportcollocationtype.Name)
	return codeairportheliportcollocationtype
}

// StagePreserveOrder puts codeairportheliportcollocationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirportHeliportCollocationTypeOrder
// - update stage.CodeAirportHeliportCollocationTypeOrder accordingly
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationType_stagedOrder, stage.CodeAirportHeliportCollocationType_orderStaged, &stage.CodeAirportHeliportCollocationTypeOrder, stage.CodeAirportHeliportCollocationTypes_mapString, codeairportheliportcollocationtype, order, codeairportheliportcollocationtype.Name)
}

// Unstage removes codeairportheliportcollocationtype off the model stage
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) Unstage(stage *Stage) *CodeAirportHeliportCollocationType {
	__gong__unstage(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationTypes_mapString, codeairportheliportcollocationtype, codeairportheliportcollocationtype.Name)
	return codeairportheliportcollocationtype
}

// UnstageVoid removes codeairportheliportcollocationtype off the model stage
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) UnstageVoid(stage *Stage) {
	codeairportheliportcollocationtype.Unstage(stage)
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) StageVoid(stage *Stage) {
	codeairportheliportcollocationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GetName() (res string) {
	return codeairportheliportcollocationtype.Name
}

// for satisfaction of GongStruct interface
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) SetName(name string) {
	codeairportheliportcollocationtype.Name = name
}

// Stage puts codeairportheliportdesignatortype to the model stage
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) Stage(stage *Stage) *CodeAirportHeliportDesignatorType {
	__gong__stage(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorType_stagedOrder, stage.CodeAirportHeliportDesignatorType_orderStaged, &stage.CodeAirportHeliportDesignatorTypeOrder, stage.CodeAirportHeliportDesignatorTypes_mapString, codeairportheliportdesignatortype, codeairportheliportdesignatortype.Name)
	return codeairportheliportdesignatortype
}

// StagePreserveOrder puts codeairportheliportdesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirportHeliportDesignatorTypeOrder
// - update stage.CodeAirportHeliportDesignatorTypeOrder accordingly
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorType_stagedOrder, stage.CodeAirportHeliportDesignatorType_orderStaged, &stage.CodeAirportHeliportDesignatorTypeOrder, stage.CodeAirportHeliportDesignatorTypes_mapString, codeairportheliportdesignatortype, order, codeairportheliportdesignatortype.Name)
}

// Unstage removes codeairportheliportdesignatortype off the model stage
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) Unstage(stage *Stage) *CodeAirportHeliportDesignatorType {
	__gong__unstage(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorTypes_mapString, codeairportheliportdesignatortype, codeairportheliportdesignatortype.Name)
	return codeairportheliportdesignatortype
}

// UnstageVoid removes codeairportheliportdesignatortype off the model stage
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) UnstageVoid(stage *Stage) {
	codeairportheliportdesignatortype.Unstage(stage)
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) StageVoid(stage *Stage) {
	codeairportheliportdesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GetName() (res string) {
	return codeairportheliportdesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) SetName(name string) {
	codeairportheliportdesignatortype.Name = name
}

// Stage puts codeairportheliporttype to the model stage
func (codeairportheliporttype *CodeAirportHeliportType) Stage(stage *Stage) *CodeAirportHeliportType {
	__gong__stage(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportType_stagedOrder, stage.CodeAirportHeliportType_orderStaged, &stage.CodeAirportHeliportTypeOrder, stage.CodeAirportHeliportTypes_mapString, codeairportheliporttype, codeairportheliporttype.Name)
	return codeairportheliporttype
}

// StagePreserveOrder puts codeairportheliporttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirportHeliportTypeOrder
// - update stage.CodeAirportHeliportTypeOrder accordingly
func (codeairportheliporttype *CodeAirportHeliportType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportType_stagedOrder, stage.CodeAirportHeliportType_orderStaged, &stage.CodeAirportHeliportTypeOrder, stage.CodeAirportHeliportTypes_mapString, codeairportheliporttype, order, codeairportheliporttype.Name)
}

// Unstage removes codeairportheliporttype off the model stage
func (codeairportheliporttype *CodeAirportHeliportType) Unstage(stage *Stage) *CodeAirportHeliportType {
	__gong__unstage(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportTypes_mapString, codeairportheliporttype, codeairportheliporttype.Name)
	return codeairportheliporttype
}

// UnstageVoid removes codeairportheliporttype off the model stage
func (codeairportheliporttype *CodeAirportHeliportType) UnstageVoid(stage *Stage) {
	codeairportheliporttype.Unstage(stage)
}

func (codeairportheliporttype *CodeAirportHeliportType) StageVoid(stage *Stage) {
	codeairportheliporttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairportheliporttype *CodeAirportHeliportType) GetName() (res string) {
	return codeairportheliporttype.Name
}

// for satisfaction of GongStruct interface
func (codeairportheliporttype *CodeAirportHeliportType) SetName(name string) {
	codeairportheliporttype.Name = name
}

// Stage puts codeairportwarningtype to the model stage
func (codeairportwarningtype *CodeAirportWarningType) Stage(stage *Stage) *CodeAirportWarningType {
	__gong__stage(stage.CodeAirportWarningTypes, stage.CodeAirportWarningType_stagedOrder, stage.CodeAirportWarningType_orderStaged, &stage.CodeAirportWarningTypeOrder, stage.CodeAirportWarningTypes_mapString, codeairportwarningtype, codeairportwarningtype.Name)
	return codeairportwarningtype
}

// StagePreserveOrder puts codeairportwarningtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirportWarningTypeOrder
// - update stage.CodeAirportWarningTypeOrder accordingly
func (codeairportwarningtype *CodeAirportWarningType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirportWarningTypes, stage.CodeAirportWarningType_stagedOrder, stage.CodeAirportWarningType_orderStaged, &stage.CodeAirportWarningTypeOrder, stage.CodeAirportWarningTypes_mapString, codeairportwarningtype, order, codeairportwarningtype.Name)
}

// Unstage removes codeairportwarningtype off the model stage
func (codeairportwarningtype *CodeAirportWarningType) Unstage(stage *Stage) *CodeAirportWarningType {
	__gong__unstage(stage.CodeAirportWarningTypes, stage.CodeAirportWarningTypes_mapString, codeairportwarningtype, codeairportwarningtype.Name)
	return codeairportwarningtype
}

// UnstageVoid removes codeairportwarningtype off the model stage
func (codeairportwarningtype *CodeAirportWarningType) UnstageVoid(stage *Stage) {
	codeairportwarningtype.Unstage(stage)
}

func (codeairportwarningtype *CodeAirportWarningType) StageVoid(stage *Stage) {
	codeairportwarningtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairportwarningtype *CodeAirportWarningType) GetName() (res string) {
	return codeairportwarningtype.Name
}

// for satisfaction of GongStruct interface
func (codeairportwarningtype *CodeAirportWarningType) SetName(name string) {
	codeairportwarningtype.Name = name
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

// Stage puts codeapproachlightingicaotype to the model stage
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) Stage(stage *Stage) *CodeApproachLightingICAOType {
	__gong__stage(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOType_stagedOrder, stage.CodeApproachLightingICAOType_orderStaged, &stage.CodeApproachLightingICAOTypeOrder, stage.CodeApproachLightingICAOTypes_mapString, codeapproachlightingicaotype, codeapproachlightingicaotype.Name)
	return codeapproachlightingicaotype
}

// StagePreserveOrder puts codeapproachlightingicaotype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachLightingICAOTypeOrder
// - update stage.CodeApproachLightingICAOTypeOrder accordingly
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOType_stagedOrder, stage.CodeApproachLightingICAOType_orderStaged, &stage.CodeApproachLightingICAOTypeOrder, stage.CodeApproachLightingICAOTypes_mapString, codeapproachlightingicaotype, order, codeapproachlightingicaotype.Name)
}

// Unstage removes codeapproachlightingicaotype off the model stage
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) Unstage(stage *Stage) *CodeApproachLightingICAOType {
	__gong__unstage(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOTypes_mapString, codeapproachlightingicaotype, codeapproachlightingicaotype.Name)
	return codeapproachlightingicaotype
}

// UnstageVoid removes codeapproachlightingicaotype off the model stage
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) UnstageVoid(stage *Stage) {
	codeapproachlightingicaotype.Unstage(stage)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) StageVoid(stage *Stage) {
	codeapproachlightingicaotype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GetName() (res string) {
	return codeapproachlightingicaotype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) SetName(name string) {
	codeapproachlightingicaotype.Name = name
}

// Stage puts codeapproachlightingtype to the model stage
func (codeapproachlightingtype *CodeApproachLightingType) Stage(stage *Stage) *CodeApproachLightingType {
	__gong__stage(stage.CodeApproachLightingTypes, stage.CodeApproachLightingType_stagedOrder, stage.CodeApproachLightingType_orderStaged, &stage.CodeApproachLightingTypeOrder, stage.CodeApproachLightingTypes_mapString, codeapproachlightingtype, codeapproachlightingtype.Name)
	return codeapproachlightingtype
}

// StagePreserveOrder puts codeapproachlightingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApproachLightingTypeOrder
// - update stage.CodeApproachLightingTypeOrder accordingly
func (codeapproachlightingtype *CodeApproachLightingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApproachLightingTypes, stage.CodeApproachLightingType_stagedOrder, stage.CodeApproachLightingType_orderStaged, &stage.CodeApproachLightingTypeOrder, stage.CodeApproachLightingTypes_mapString, codeapproachlightingtype, order, codeapproachlightingtype.Name)
}

// Unstage removes codeapproachlightingtype off the model stage
func (codeapproachlightingtype *CodeApproachLightingType) Unstage(stage *Stage) *CodeApproachLightingType {
	__gong__unstage(stage.CodeApproachLightingTypes, stage.CodeApproachLightingTypes_mapString, codeapproachlightingtype, codeapproachlightingtype.Name)
	return codeapproachlightingtype
}

// UnstageVoid removes codeapproachlightingtype off the model stage
func (codeapproachlightingtype *CodeApproachLightingType) UnstageVoid(stage *Stage) {
	codeapproachlightingtype.Unstage(stage)
}

func (codeapproachlightingtype *CodeApproachLightingType) StageVoid(stage *Stage) {
	codeapproachlightingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapproachlightingtype *CodeApproachLightingType) GetName() (res string) {
	return codeapproachlightingtype.Name
}

// for satisfaction of GongStruct interface
func (codeapproachlightingtype *CodeApproachLightingType) SetName(name string) {
	codeapproachlightingtype.Name = name
}

// Stage puts codeapronelementtype to the model stage
func (codeapronelementtype *CodeApronElementType) Stage(stage *Stage) *CodeApronElementType {
	__gong__stage(stage.CodeApronElementTypes, stage.CodeApronElementType_stagedOrder, stage.CodeApronElementType_orderStaged, &stage.CodeApronElementTypeOrder, stage.CodeApronElementTypes_mapString, codeapronelementtype, codeapronelementtype.Name)
	return codeapronelementtype
}

// StagePreserveOrder puts codeapronelementtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApronElementTypeOrder
// - update stage.CodeApronElementTypeOrder accordingly
func (codeapronelementtype *CodeApronElementType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApronElementTypes, stage.CodeApronElementType_stagedOrder, stage.CodeApronElementType_orderStaged, &stage.CodeApronElementTypeOrder, stage.CodeApronElementTypes_mapString, codeapronelementtype, order, codeapronelementtype.Name)
}

// Unstage removes codeapronelementtype off the model stage
func (codeapronelementtype *CodeApronElementType) Unstage(stage *Stage) *CodeApronElementType {
	__gong__unstage(stage.CodeApronElementTypes, stage.CodeApronElementTypes_mapString, codeapronelementtype, codeapronelementtype.Name)
	return codeapronelementtype
}

// UnstageVoid removes codeapronelementtype off the model stage
func (codeapronelementtype *CodeApronElementType) UnstageVoid(stage *Stage) {
	codeapronelementtype.Unstage(stage)
}

func (codeapronelementtype *CodeApronElementType) StageVoid(stage *Stage) {
	codeapronelementtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapronelementtype *CodeApronElementType) GetName() (res string) {
	return codeapronelementtype.Name
}

// for satisfaction of GongStruct interface
func (codeapronelementtype *CodeApronElementType) SetName(name string) {
	codeapronelementtype.Name = name
}

// Stage puts codeapronsectiontype to the model stage
func (codeapronsectiontype *CodeApronSectionType) Stage(stage *Stage) *CodeApronSectionType {
	__gong__stage(stage.CodeApronSectionTypes, stage.CodeApronSectionType_stagedOrder, stage.CodeApronSectionType_orderStaged, &stage.CodeApronSectionTypeOrder, stage.CodeApronSectionTypes_mapString, codeapronsectiontype, codeapronsectiontype.Name)
	return codeapronsectiontype
}

// StagePreserveOrder puts codeapronsectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeApronSectionTypeOrder
// - update stage.CodeApronSectionTypeOrder accordingly
func (codeapronsectiontype *CodeApronSectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeApronSectionTypes, stage.CodeApronSectionType_stagedOrder, stage.CodeApronSectionType_orderStaged, &stage.CodeApronSectionTypeOrder, stage.CodeApronSectionTypes_mapString, codeapronsectiontype, order, codeapronsectiontype.Name)
}

// Unstage removes codeapronsectiontype off the model stage
func (codeapronsectiontype *CodeApronSectionType) Unstage(stage *Stage) *CodeApronSectionType {
	__gong__unstage(stage.CodeApronSectionTypes, stage.CodeApronSectionTypes_mapString, codeapronsectiontype, codeapronsectiontype.Name)
	return codeapronsectiontype
}

// UnstageVoid removes codeapronsectiontype off the model stage
func (codeapronsectiontype *CodeApronSectionType) UnstageVoid(stage *Stage) {
	codeapronsectiontype.Unstage(stage)
}

func (codeapronsectiontype *CodeApronSectionType) StageVoid(stage *Stage) {
	codeapronsectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeapronsectiontype *CodeApronSectionType) GetName() (res string) {
	return codeapronsectiontype.Name
}

// for satisfaction of GongStruct interface
func (codeapronsectiontype *CodeApronSectionType) SetName(name string) {
	codeapronsectiontype.Name = name
}

// Stage puts codearrestinggearenergyabsorbtype to the model stage
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) Stage(stage *Stage) *CodeArrestingGearEnergyAbsorbType {
	__gong__stage(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbType_stagedOrder, stage.CodeArrestingGearEnergyAbsorbType_orderStaged, &stage.CodeArrestingGearEnergyAbsorbTypeOrder, stage.CodeArrestingGearEnergyAbsorbTypes_mapString, codearrestinggearenergyabsorbtype, codearrestinggearenergyabsorbtype.Name)
	return codearrestinggearenergyabsorbtype
}

// StagePreserveOrder puts codearrestinggearenergyabsorbtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeArrestingGearEnergyAbsorbTypeOrder
// - update stage.CodeArrestingGearEnergyAbsorbTypeOrder accordingly
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbType_stagedOrder, stage.CodeArrestingGearEnergyAbsorbType_orderStaged, &stage.CodeArrestingGearEnergyAbsorbTypeOrder, stage.CodeArrestingGearEnergyAbsorbTypes_mapString, codearrestinggearenergyabsorbtype, order, codearrestinggearenergyabsorbtype.Name)
}

// Unstage removes codearrestinggearenergyabsorbtype off the model stage
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) Unstage(stage *Stage) *CodeArrestingGearEnergyAbsorbType {
	__gong__unstage(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbTypes_mapString, codearrestinggearenergyabsorbtype, codearrestinggearenergyabsorbtype.Name)
	return codearrestinggearenergyabsorbtype
}

// UnstageVoid removes codearrestinggearenergyabsorbtype off the model stage
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) UnstageVoid(stage *Stage) {
	codearrestinggearenergyabsorbtype.Unstage(stage)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) StageVoid(stage *Stage) {
	codearrestinggearenergyabsorbtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GetName() (res string) {
	return codearrestinggearenergyabsorbtype.Name
}

// for satisfaction of GongStruct interface
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) SetName(name string) {
	codearrestinggearenergyabsorbtype.Name = name
}

// Stage puts codearrestinggearengagedevicetype to the model stage
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) Stage(stage *Stage) *CodeArrestingGearEngageDeviceType {
	__gong__stage(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceType_stagedOrder, stage.CodeArrestingGearEngageDeviceType_orderStaged, &stage.CodeArrestingGearEngageDeviceTypeOrder, stage.CodeArrestingGearEngageDeviceTypes_mapString, codearrestinggearengagedevicetype, codearrestinggearengagedevicetype.Name)
	return codearrestinggearengagedevicetype
}

// StagePreserveOrder puts codearrestinggearengagedevicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeArrestingGearEngageDeviceTypeOrder
// - update stage.CodeArrestingGearEngageDeviceTypeOrder accordingly
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceType_stagedOrder, stage.CodeArrestingGearEngageDeviceType_orderStaged, &stage.CodeArrestingGearEngageDeviceTypeOrder, stage.CodeArrestingGearEngageDeviceTypes_mapString, codearrestinggearengagedevicetype, order, codearrestinggearengagedevicetype.Name)
}

// Unstage removes codearrestinggearengagedevicetype off the model stage
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) Unstage(stage *Stage) *CodeArrestingGearEngageDeviceType {
	__gong__unstage(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceTypes_mapString, codearrestinggearengagedevicetype, codearrestinggearengagedevicetype.Name)
	return codearrestinggearengagedevicetype
}

// UnstageVoid removes codearrestinggearengagedevicetype off the model stage
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) UnstageVoid(stage *Stage) {
	codearrestinggearengagedevicetype.Unstage(stage)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) StageVoid(stage *Stage) {
	codearrestinggearengagedevicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GetName() (res string) {
	return codearrestinggearengagedevicetype.Name
}

// for satisfaction of GongStruct interface
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) SetName(name string) {
	codearrestinggearengagedevicetype.Name = name
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

// Stage puts codebuoydesignatortype to the model stage
func (codebuoydesignatortype *CodeBuoyDesignatorType) Stage(stage *Stage) *CodeBuoyDesignatorType {
	__gong__stage(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorType_stagedOrder, stage.CodeBuoyDesignatorType_orderStaged, &stage.CodeBuoyDesignatorTypeOrder, stage.CodeBuoyDesignatorTypes_mapString, codebuoydesignatortype, codebuoydesignatortype.Name)
	return codebuoydesignatortype
}

// StagePreserveOrder puts codebuoydesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeBuoyDesignatorTypeOrder
// - update stage.CodeBuoyDesignatorTypeOrder accordingly
func (codebuoydesignatortype *CodeBuoyDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorType_stagedOrder, stage.CodeBuoyDesignatorType_orderStaged, &stage.CodeBuoyDesignatorTypeOrder, stage.CodeBuoyDesignatorTypes_mapString, codebuoydesignatortype, order, codebuoydesignatortype.Name)
}

// Unstage removes codebuoydesignatortype off the model stage
func (codebuoydesignatortype *CodeBuoyDesignatorType) Unstage(stage *Stage) *CodeBuoyDesignatorType {
	__gong__unstage(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorTypes_mapString, codebuoydesignatortype, codebuoydesignatortype.Name)
	return codebuoydesignatortype
}

// UnstageVoid removes codebuoydesignatortype off the model stage
func (codebuoydesignatortype *CodeBuoyDesignatorType) UnstageVoid(stage *Stage) {
	codebuoydesignatortype.Unstage(stage)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) StageVoid(stage *Stage) {
	codebuoydesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codebuoydesignatortype *CodeBuoyDesignatorType) GetName() (res string) {
	return codebuoydesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codebuoydesignatortype *CodeBuoyDesignatorType) SetName(name string) {
	codebuoydesignatortype.Name = name
}

// Stage puts codebuoytype to the model stage
func (codebuoytype *CodeBuoyType) Stage(stage *Stage) *CodeBuoyType {
	__gong__stage(stage.CodeBuoyTypes, stage.CodeBuoyType_stagedOrder, stage.CodeBuoyType_orderStaged, &stage.CodeBuoyTypeOrder, stage.CodeBuoyTypes_mapString, codebuoytype, codebuoytype.Name)
	return codebuoytype
}

// StagePreserveOrder puts codebuoytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeBuoyTypeOrder
// - update stage.CodeBuoyTypeOrder accordingly
func (codebuoytype *CodeBuoyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeBuoyTypes, stage.CodeBuoyType_stagedOrder, stage.CodeBuoyType_orderStaged, &stage.CodeBuoyTypeOrder, stage.CodeBuoyTypes_mapString, codebuoytype, order, codebuoytype.Name)
}

// Unstage removes codebuoytype off the model stage
func (codebuoytype *CodeBuoyType) Unstage(stage *Stage) *CodeBuoyType {
	__gong__unstage(stage.CodeBuoyTypes, stage.CodeBuoyTypes_mapString, codebuoytype, codebuoytype.Name)
	return codebuoytype
}

// UnstageVoid removes codebuoytype off the model stage
func (codebuoytype *CodeBuoyType) UnstageVoid(stage *Stage) {
	codebuoytype.Unstage(stage)
}

func (codebuoytype *CodeBuoyType) StageVoid(stage *Stage) {
	codebuoytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codebuoytype *CodeBuoyType) GetName() (res string) {
	return codebuoytype.Name
}

// for satisfaction of GongStruct interface
func (codebuoytype *CodeBuoyType) SetName(name string) {
	codebuoytype.Name = name
}

// Stage puts codecolourtype to the model stage
func (codecolourtype *CodeColourType) Stage(stage *Stage) *CodeColourType {
	__gong__stage(stage.CodeColourTypes, stage.CodeColourType_stagedOrder, stage.CodeColourType_orderStaged, &stage.CodeColourTypeOrder, stage.CodeColourTypes_mapString, codecolourtype, codecolourtype.Name)
	return codecolourtype
}

// StagePreserveOrder puts codecolourtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeColourTypeOrder
// - update stage.CodeColourTypeOrder accordingly
func (codecolourtype *CodeColourType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeColourTypes, stage.CodeColourType_stagedOrder, stage.CodeColourType_orderStaged, &stage.CodeColourTypeOrder, stage.CodeColourTypes_mapString, codecolourtype, order, codecolourtype.Name)
}

// Unstage removes codecolourtype off the model stage
func (codecolourtype *CodeColourType) Unstage(stage *Stage) *CodeColourType {
	__gong__unstage(stage.CodeColourTypes, stage.CodeColourTypes_mapString, codecolourtype, codecolourtype.Name)
	return codecolourtype
}

// UnstageVoid removes codecolourtype off the model stage
func (codecolourtype *CodeColourType) UnstageVoid(stage *Stage) {
	codecolourtype.Unstage(stage)
}

func (codecolourtype *CodeColourType) StageVoid(stage *Stage) {
	codecolourtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecolourtype *CodeColourType) GetName() (res string) {
	return codecolourtype.Name
}

// for satisfaction of GongStruct interface
func (codecolourtype *CodeColourType) SetName(name string) {
	codecolourtype.Name = name
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

// Stage puts codecontaminationtype to the model stage
func (codecontaminationtype *CodeContaminationType) Stage(stage *Stage) *CodeContaminationType {
	__gong__stage(stage.CodeContaminationTypes, stage.CodeContaminationType_stagedOrder, stage.CodeContaminationType_orderStaged, &stage.CodeContaminationTypeOrder, stage.CodeContaminationTypes_mapString, codecontaminationtype, codecontaminationtype.Name)
	return codecontaminationtype
}

// StagePreserveOrder puts codecontaminationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeContaminationTypeOrder
// - update stage.CodeContaminationTypeOrder accordingly
func (codecontaminationtype *CodeContaminationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeContaminationTypes, stage.CodeContaminationType_stagedOrder, stage.CodeContaminationType_orderStaged, &stage.CodeContaminationTypeOrder, stage.CodeContaminationTypes_mapString, codecontaminationtype, order, codecontaminationtype.Name)
}

// Unstage removes codecontaminationtype off the model stage
func (codecontaminationtype *CodeContaminationType) Unstage(stage *Stage) *CodeContaminationType {
	__gong__unstage(stage.CodeContaminationTypes, stage.CodeContaminationTypes_mapString, codecontaminationtype, codecontaminationtype.Name)
	return codecontaminationtype
}

// UnstageVoid removes codecontaminationtype off the model stage
func (codecontaminationtype *CodeContaminationType) UnstageVoid(stage *Stage) {
	codecontaminationtype.Unstage(stage)
}

func (codecontaminationtype *CodeContaminationType) StageVoid(stage *Stage) {
	codecontaminationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codecontaminationtype *CodeContaminationType) GetName() (res string) {
	return codecontaminationtype.Name
}

// for satisfaction of GongStruct interface
func (codecontaminationtype *CodeContaminationType) SetName(name string) {
	codecontaminationtype.Name = name
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

// Stage puts codedeclareddistancetype to the model stage
func (codedeclareddistancetype *CodeDeclaredDistanceType) Stage(stage *Stage) *CodeDeclaredDistanceType {
	__gong__stage(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceType_stagedOrder, stage.CodeDeclaredDistanceType_orderStaged, &stage.CodeDeclaredDistanceTypeOrder, stage.CodeDeclaredDistanceTypes_mapString, codedeclareddistancetype, codedeclareddistancetype.Name)
	return codedeclareddistancetype
}

// StagePreserveOrder puts codedeclareddistancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeDeclaredDistanceTypeOrder
// - update stage.CodeDeclaredDistanceTypeOrder accordingly
func (codedeclareddistancetype *CodeDeclaredDistanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceType_stagedOrder, stage.CodeDeclaredDistanceType_orderStaged, &stage.CodeDeclaredDistanceTypeOrder, stage.CodeDeclaredDistanceTypes_mapString, codedeclareddistancetype, order, codedeclareddistancetype.Name)
}

// Unstage removes codedeclareddistancetype off the model stage
func (codedeclareddistancetype *CodeDeclaredDistanceType) Unstage(stage *Stage) *CodeDeclaredDistanceType {
	__gong__unstage(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceTypes_mapString, codedeclareddistancetype, codedeclareddistancetype.Name)
	return codedeclareddistancetype
}

// UnstageVoid removes codedeclareddistancetype off the model stage
func (codedeclareddistancetype *CodeDeclaredDistanceType) UnstageVoid(stage *Stage) {
	codedeclareddistancetype.Unstage(stage)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) StageVoid(stage *Stage) {
	codedeclareddistancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codedeclareddistancetype *CodeDeclaredDistanceType) GetName() (res string) {
	return codedeclareddistancetype.Name
}

// for satisfaction of GongStruct interface
func (codedeclareddistancetype *CodeDeclaredDistanceType) SetName(name string) {
	codedeclareddistancetype.Name = name
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

// Stage puts codefrictiondevicetype to the model stage
func (codefrictiondevicetype *CodeFrictionDeviceType) Stage(stage *Stage) *CodeFrictionDeviceType {
	__gong__stage(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceType_stagedOrder, stage.CodeFrictionDeviceType_orderStaged, &stage.CodeFrictionDeviceTypeOrder, stage.CodeFrictionDeviceTypes_mapString, codefrictiondevicetype, codefrictiondevicetype.Name)
	return codefrictiondevicetype
}

// StagePreserveOrder puts codefrictiondevicetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFrictionDeviceTypeOrder
// - update stage.CodeFrictionDeviceTypeOrder accordingly
func (codefrictiondevicetype *CodeFrictionDeviceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceType_stagedOrder, stage.CodeFrictionDeviceType_orderStaged, &stage.CodeFrictionDeviceTypeOrder, stage.CodeFrictionDeviceTypes_mapString, codefrictiondevicetype, order, codefrictiondevicetype.Name)
}

// Unstage removes codefrictiondevicetype off the model stage
func (codefrictiondevicetype *CodeFrictionDeviceType) Unstage(stage *Stage) *CodeFrictionDeviceType {
	__gong__unstage(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceTypes_mapString, codefrictiondevicetype, codefrictiondevicetype.Name)
	return codefrictiondevicetype
}

// UnstageVoid removes codefrictiondevicetype off the model stage
func (codefrictiondevicetype *CodeFrictionDeviceType) UnstageVoid(stage *Stage) {
	codefrictiondevicetype.Unstage(stage)
}

func (codefrictiondevicetype *CodeFrictionDeviceType) StageVoid(stage *Stage) {
	codefrictiondevicetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefrictiondevicetype *CodeFrictionDeviceType) GetName() (res string) {
	return codefrictiondevicetype.Name
}

// for satisfaction of GongStruct interface
func (codefrictiondevicetype *CodeFrictionDeviceType) SetName(name string) {
	codefrictiondevicetype.Name = name
}

// Stage puts codefrictionestimatetype to the model stage
func (codefrictionestimatetype *CodeFrictionEstimateType) Stage(stage *Stage) *CodeFrictionEstimateType {
	__gong__stage(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateType_stagedOrder, stage.CodeFrictionEstimateType_orderStaged, &stage.CodeFrictionEstimateTypeOrder, stage.CodeFrictionEstimateTypes_mapString, codefrictionestimatetype, codefrictionestimatetype.Name)
	return codefrictionestimatetype
}

// StagePreserveOrder puts codefrictionestimatetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeFrictionEstimateTypeOrder
// - update stage.CodeFrictionEstimateTypeOrder accordingly
func (codefrictionestimatetype *CodeFrictionEstimateType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateType_stagedOrder, stage.CodeFrictionEstimateType_orderStaged, &stage.CodeFrictionEstimateTypeOrder, stage.CodeFrictionEstimateTypes_mapString, codefrictionestimatetype, order, codefrictionestimatetype.Name)
}

// Unstage removes codefrictionestimatetype off the model stage
func (codefrictionestimatetype *CodeFrictionEstimateType) Unstage(stage *Stage) *CodeFrictionEstimateType {
	__gong__unstage(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateTypes_mapString, codefrictionestimatetype, codefrictionestimatetype.Name)
	return codefrictionestimatetype
}

// UnstageVoid removes codefrictionestimatetype off the model stage
func (codefrictionestimatetype *CodeFrictionEstimateType) UnstageVoid(stage *Stage) {
	codefrictionestimatetype.Unstage(stage)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) StageVoid(stage *Stage) {
	codefrictionestimatetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codefrictionestimatetype *CodeFrictionEstimateType) GetName() (res string) {
	return codefrictionestimatetype.Name
}

// for satisfaction of GongStruct interface
func (codefrictionestimatetype *CodeFrictionEstimateType) SetName(name string) {
	codefrictionestimatetype.Name = name
}

// Stage puts codegradeseparationtype to the model stage
func (codegradeseparationtype *CodeGradeSeparationType) Stage(stage *Stage) *CodeGradeSeparationType {
	__gong__stage(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationType_stagedOrder, stage.CodeGradeSeparationType_orderStaged, &stage.CodeGradeSeparationTypeOrder, stage.CodeGradeSeparationTypes_mapString, codegradeseparationtype, codegradeseparationtype.Name)
	return codegradeseparationtype
}

// StagePreserveOrder puts codegradeseparationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeGradeSeparationTypeOrder
// - update stage.CodeGradeSeparationTypeOrder accordingly
func (codegradeseparationtype *CodeGradeSeparationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationType_stagedOrder, stage.CodeGradeSeparationType_orderStaged, &stage.CodeGradeSeparationTypeOrder, stage.CodeGradeSeparationTypes_mapString, codegradeseparationtype, order, codegradeseparationtype.Name)
}

// Unstage removes codegradeseparationtype off the model stage
func (codegradeseparationtype *CodeGradeSeparationType) Unstage(stage *Stage) *CodeGradeSeparationType {
	__gong__unstage(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationTypes_mapString, codegradeseparationtype, codegradeseparationtype.Name)
	return codegradeseparationtype
}

// UnstageVoid removes codegradeseparationtype off the model stage
func (codegradeseparationtype *CodeGradeSeparationType) UnstageVoid(stage *Stage) {
	codegradeseparationtype.Unstage(stage)
}

func (codegradeseparationtype *CodeGradeSeparationType) StageVoid(stage *Stage) {
	codegradeseparationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codegradeseparationtype *CodeGradeSeparationType) GetName() (res string) {
	return codegradeseparationtype.Name
}

// for satisfaction of GongStruct interface
func (codegradeseparationtype *CodeGradeSeparationType) SetName(name string) {
	codegradeseparationtype.Name = name
}

// Stage puts codeguidancelinetype to the model stage
func (codeguidancelinetype *CodeGuidanceLineType) Stage(stage *Stage) *CodeGuidanceLineType {
	__gong__stage(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineType_stagedOrder, stage.CodeGuidanceLineType_orderStaged, &stage.CodeGuidanceLineTypeOrder, stage.CodeGuidanceLineTypes_mapString, codeguidancelinetype, codeguidancelinetype.Name)
	return codeguidancelinetype
}

// StagePreserveOrder puts codeguidancelinetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeGuidanceLineTypeOrder
// - update stage.CodeGuidanceLineTypeOrder accordingly
func (codeguidancelinetype *CodeGuidanceLineType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineType_stagedOrder, stage.CodeGuidanceLineType_orderStaged, &stage.CodeGuidanceLineTypeOrder, stage.CodeGuidanceLineTypes_mapString, codeguidancelinetype, order, codeguidancelinetype.Name)
}

// Unstage removes codeguidancelinetype off the model stage
func (codeguidancelinetype *CodeGuidanceLineType) Unstage(stage *Stage) *CodeGuidanceLineType {
	__gong__unstage(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineTypes_mapString, codeguidancelinetype, codeguidancelinetype.Name)
	return codeguidancelinetype
}

// UnstageVoid removes codeguidancelinetype off the model stage
func (codeguidancelinetype *CodeGuidanceLineType) UnstageVoid(stage *Stage) {
	codeguidancelinetype.Unstage(stage)
}

func (codeguidancelinetype *CodeGuidanceLineType) StageVoid(stage *Stage) {
	codeguidancelinetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeguidancelinetype *CodeGuidanceLineType) GetName() (res string) {
	return codeguidancelinetype.Name
}

// for satisfaction of GongStruct interface
func (codeguidancelinetype *CodeGuidanceLineType) SetName(name string) {
	codeguidancelinetype.Name = name
}

// Stage puts codehelicopterperformancetype to the model stage
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) Stage(stage *Stage) *CodeHelicopterPerformanceType {
	__gong__stage(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceType_stagedOrder, stage.CodeHelicopterPerformanceType_orderStaged, &stage.CodeHelicopterPerformanceTypeOrder, stage.CodeHelicopterPerformanceTypes_mapString, codehelicopterperformancetype, codehelicopterperformancetype.Name)
	return codehelicopterperformancetype
}

// StagePreserveOrder puts codehelicopterperformancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeHelicopterPerformanceTypeOrder
// - update stage.CodeHelicopterPerformanceTypeOrder accordingly
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceType_stagedOrder, stage.CodeHelicopterPerformanceType_orderStaged, &stage.CodeHelicopterPerformanceTypeOrder, stage.CodeHelicopterPerformanceTypes_mapString, codehelicopterperformancetype, order, codehelicopterperformancetype.Name)
}

// Unstage removes codehelicopterperformancetype off the model stage
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) Unstage(stage *Stage) *CodeHelicopterPerformanceType {
	__gong__unstage(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceTypes_mapString, codehelicopterperformancetype, codehelicopterperformancetype.Name)
	return codehelicopterperformancetype
}

// UnstageVoid removes codehelicopterperformancetype off the model stage
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) UnstageVoid(stage *Stage) {
	codehelicopterperformancetype.Unstage(stage)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) StageVoid(stage *Stage) {
	codehelicopterperformancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GetName() (res string) {
	return codehelicopterperformancetype.Name
}

// for satisfaction of GongStruct interface
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) SetName(name string) {
	codehelicopterperformancetype.Name = name
}

// Stage puts codeholdingcategorytype to the model stage
func (codeholdingcategorytype *CodeHoldingCategoryType) Stage(stage *Stage) *CodeHoldingCategoryType {
	__gong__stage(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryType_stagedOrder, stage.CodeHoldingCategoryType_orderStaged, &stage.CodeHoldingCategoryTypeOrder, stage.CodeHoldingCategoryTypes_mapString, codeholdingcategorytype, codeholdingcategorytype.Name)
	return codeholdingcategorytype
}

// StagePreserveOrder puts codeholdingcategorytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeHoldingCategoryTypeOrder
// - update stage.CodeHoldingCategoryTypeOrder accordingly
func (codeholdingcategorytype *CodeHoldingCategoryType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryType_stagedOrder, stage.CodeHoldingCategoryType_orderStaged, &stage.CodeHoldingCategoryTypeOrder, stage.CodeHoldingCategoryTypes_mapString, codeholdingcategorytype, order, codeholdingcategorytype.Name)
}

// Unstage removes codeholdingcategorytype off the model stage
func (codeholdingcategorytype *CodeHoldingCategoryType) Unstage(stage *Stage) *CodeHoldingCategoryType {
	__gong__unstage(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryTypes_mapString, codeholdingcategorytype, codeholdingcategorytype.Name)
	return codeholdingcategorytype
}

// UnstageVoid removes codeholdingcategorytype off the model stage
func (codeholdingcategorytype *CodeHoldingCategoryType) UnstageVoid(stage *Stage) {
	codeholdingcategorytype.Unstage(stage)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) StageVoid(stage *Stage) {
	codeholdingcategorytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeholdingcategorytype *CodeHoldingCategoryType) GetName() (res string) {
	return codeholdingcategorytype.Name
}

// for satisfaction of GongStruct interface
func (codeholdingcategorytype *CodeHoldingCategoryType) SetName(name string) {
	codeholdingcategorytype.Name = name
}

// Stage puts codeiatatype to the model stage
func (codeiatatype *CodeIATAType) Stage(stage *Stage) *CodeIATAType {
	__gong__stage(stage.CodeIATATypes, stage.CodeIATAType_stagedOrder, stage.CodeIATAType_orderStaged, &stage.CodeIATATypeOrder, stage.CodeIATATypes_mapString, codeiatatype, codeiatatype.Name)
	return codeiatatype
}

// StagePreserveOrder puts codeiatatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeIATATypeOrder
// - update stage.CodeIATATypeOrder accordingly
func (codeiatatype *CodeIATAType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeIATATypes, stage.CodeIATAType_stagedOrder, stage.CodeIATAType_orderStaged, &stage.CodeIATATypeOrder, stage.CodeIATATypes_mapString, codeiatatype, order, codeiatatype.Name)
}

// Unstage removes codeiatatype off the model stage
func (codeiatatype *CodeIATAType) Unstage(stage *Stage) *CodeIATAType {
	__gong__unstage(stage.CodeIATATypes, stage.CodeIATATypes_mapString, codeiatatype, codeiatatype.Name)
	return codeiatatype
}

// UnstageVoid removes codeiatatype off the model stage
func (codeiatatype *CodeIATAType) UnstageVoid(stage *Stage) {
	codeiatatype.Unstage(stage)
}

func (codeiatatype *CodeIATAType) StageVoid(stage *Stage) {
	codeiatatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeiatatype *CodeIATAType) GetName() (res string) {
	return codeiatatype.Name
}

// for satisfaction of GongStruct interface
func (codeiatatype *CodeIATAType) SetName(name string) {
	codeiatatype.Name = name
}

// Stage puts codeicaotype to the model stage
func (codeicaotype *CodeICAOType) Stage(stage *Stage) *CodeICAOType {
	__gong__stage(stage.CodeICAOTypes, stage.CodeICAOType_stagedOrder, stage.CodeICAOType_orderStaged, &stage.CodeICAOTypeOrder, stage.CodeICAOTypes_mapString, codeicaotype, codeicaotype.Name)
	return codeicaotype
}

// StagePreserveOrder puts codeicaotype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeICAOTypeOrder
// - update stage.CodeICAOTypeOrder accordingly
func (codeicaotype *CodeICAOType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeICAOTypes, stage.CodeICAOType_stagedOrder, stage.CodeICAOType_orderStaged, &stage.CodeICAOTypeOrder, stage.CodeICAOTypes_mapString, codeicaotype, order, codeicaotype.Name)
}

// Unstage removes codeicaotype off the model stage
func (codeicaotype *CodeICAOType) Unstage(stage *Stage) *CodeICAOType {
	__gong__unstage(stage.CodeICAOTypes, stage.CodeICAOTypes_mapString, codeicaotype, codeicaotype.Name)
	return codeicaotype
}

// UnstageVoid removes codeicaotype off the model stage
func (codeicaotype *CodeICAOType) UnstageVoid(stage *Stage) {
	codeicaotype.Unstage(stage)
}

func (codeicaotype *CodeICAOType) StageVoid(stage *Stage) {
	codeicaotype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeicaotype *CodeICAOType) GetName() (res string) {
	return codeicaotype.Name
}

// for satisfaction of GongStruct interface
func (codeicaotype *CodeICAOType) SetName(name string) {
	codeicaotype.Name = name
}

// Stage puts codelightholdingpositiontype to the model stage
func (codelightholdingpositiontype *CodeLightHoldingPositionType) Stage(stage *Stage) *CodeLightHoldingPositionType {
	__gong__stage(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionType_stagedOrder, stage.CodeLightHoldingPositionType_orderStaged, &stage.CodeLightHoldingPositionTypeOrder, stage.CodeLightHoldingPositionTypes_mapString, codelightholdingpositiontype, codelightholdingpositiontype.Name)
	return codelightholdingpositiontype
}

// StagePreserveOrder puts codelightholdingpositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLightHoldingPositionTypeOrder
// - update stage.CodeLightHoldingPositionTypeOrder accordingly
func (codelightholdingpositiontype *CodeLightHoldingPositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionType_stagedOrder, stage.CodeLightHoldingPositionType_orderStaged, &stage.CodeLightHoldingPositionTypeOrder, stage.CodeLightHoldingPositionTypes_mapString, codelightholdingpositiontype, order, codelightholdingpositiontype.Name)
}

// Unstage removes codelightholdingpositiontype off the model stage
func (codelightholdingpositiontype *CodeLightHoldingPositionType) Unstage(stage *Stage) *CodeLightHoldingPositionType {
	__gong__unstage(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionTypes_mapString, codelightholdingpositiontype, codelightholdingpositiontype.Name)
	return codelightholdingpositiontype
}

// UnstageVoid removes codelightholdingpositiontype off the model stage
func (codelightholdingpositiontype *CodeLightHoldingPositionType) UnstageVoid(stage *Stage) {
	codelightholdingpositiontype.Unstage(stage)
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) StageVoid(stage *Stage) {
	codelightholdingpositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelightholdingpositiontype *CodeLightHoldingPositionType) GetName() (res string) {
	return codelightholdingpositiontype.Name
}

// for satisfaction of GongStruct interface
func (codelightholdingpositiontype *CodeLightHoldingPositionType) SetName(name string) {
	codelightholdingpositiontype.Name = name
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

// Stage puts codelightsourcetype to the model stage
func (codelightsourcetype *CodeLightSourceType) Stage(stage *Stage) *CodeLightSourceType {
	__gong__stage(stage.CodeLightSourceTypes, stage.CodeLightSourceType_stagedOrder, stage.CodeLightSourceType_orderStaged, &stage.CodeLightSourceTypeOrder, stage.CodeLightSourceTypes_mapString, codelightsourcetype, codelightsourcetype.Name)
	return codelightsourcetype
}

// StagePreserveOrder puts codelightsourcetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLightSourceTypeOrder
// - update stage.CodeLightSourceTypeOrder accordingly
func (codelightsourcetype *CodeLightSourceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLightSourceTypes, stage.CodeLightSourceType_stagedOrder, stage.CodeLightSourceType_orderStaged, &stage.CodeLightSourceTypeOrder, stage.CodeLightSourceTypes_mapString, codelightsourcetype, order, codelightsourcetype.Name)
}

// Unstage removes codelightsourcetype off the model stage
func (codelightsourcetype *CodeLightSourceType) Unstage(stage *Stage) *CodeLightSourceType {
	__gong__unstage(stage.CodeLightSourceTypes, stage.CodeLightSourceTypes_mapString, codelightsourcetype, codelightsourcetype.Name)
	return codelightsourcetype
}

// UnstageVoid removes codelightsourcetype off the model stage
func (codelightsourcetype *CodeLightSourceType) UnstageVoid(stage *Stage) {
	codelightsourcetype.Unstage(stage)
}

func (codelightsourcetype *CodeLightSourceType) StageVoid(stage *Stage) {
	codelightsourcetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelightsourcetype *CodeLightSourceType) GetName() (res string) {
	return codelightsourcetype.Name
}

// for satisfaction of GongStruct interface
func (codelightsourcetype *CodeLightSourceType) SetName(name string) {
	codelightsourcetype.Name = name
}

// Stage puts codelightingjartype to the model stage
func (codelightingjartype *CodeLightingJARType) Stage(stage *Stage) *CodeLightingJARType {
	__gong__stage(stage.CodeLightingJARTypes, stage.CodeLightingJARType_stagedOrder, stage.CodeLightingJARType_orderStaged, &stage.CodeLightingJARTypeOrder, stage.CodeLightingJARTypes_mapString, codelightingjartype, codelightingjartype.Name)
	return codelightingjartype
}

// StagePreserveOrder puts codelightingjartype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLightingJARTypeOrder
// - update stage.CodeLightingJARTypeOrder accordingly
func (codelightingjartype *CodeLightingJARType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLightingJARTypes, stage.CodeLightingJARType_stagedOrder, stage.CodeLightingJARType_orderStaged, &stage.CodeLightingJARTypeOrder, stage.CodeLightingJARTypes_mapString, codelightingjartype, order, codelightingjartype.Name)
}

// Unstage removes codelightingjartype off the model stage
func (codelightingjartype *CodeLightingJARType) Unstage(stage *Stage) *CodeLightingJARType {
	__gong__unstage(stage.CodeLightingJARTypes, stage.CodeLightingJARTypes_mapString, codelightingjartype, codelightingjartype.Name)
	return codelightingjartype
}

// UnstageVoid removes codelightingjartype off the model stage
func (codelightingjartype *CodeLightingJARType) UnstageVoid(stage *Stage) {
	codelightingjartype.Unstage(stage)
}

func (codelightingjartype *CodeLightingJARType) StageVoid(stage *Stage) {
	codelightingjartype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelightingjartype *CodeLightingJARType) GetName() (res string) {
	return codelightingjartype.Name
}

// for satisfaction of GongStruct interface
func (codelightingjartype *CodeLightingJARType) SetName(name string) {
	codelightingjartype.Name = name
}

// Stage puts codeloadingbridgetype to the model stage
func (codeloadingbridgetype *CodeLoadingBridgeType) Stage(stage *Stage) *CodeLoadingBridgeType {
	__gong__stage(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeType_stagedOrder, stage.CodeLoadingBridgeType_orderStaged, &stage.CodeLoadingBridgeTypeOrder, stage.CodeLoadingBridgeTypes_mapString, codeloadingbridgetype, codeloadingbridgetype.Name)
	return codeloadingbridgetype
}

// StagePreserveOrder puts codeloadingbridgetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLoadingBridgeTypeOrder
// - update stage.CodeLoadingBridgeTypeOrder accordingly
func (codeloadingbridgetype *CodeLoadingBridgeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeType_stagedOrder, stage.CodeLoadingBridgeType_orderStaged, &stage.CodeLoadingBridgeTypeOrder, stage.CodeLoadingBridgeTypes_mapString, codeloadingbridgetype, order, codeloadingbridgetype.Name)
}

// Unstage removes codeloadingbridgetype off the model stage
func (codeloadingbridgetype *CodeLoadingBridgeType) Unstage(stage *Stage) *CodeLoadingBridgeType {
	__gong__unstage(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeTypes_mapString, codeloadingbridgetype, codeloadingbridgetype.Name)
	return codeloadingbridgetype
}

// UnstageVoid removes codeloadingbridgetype off the model stage
func (codeloadingbridgetype *CodeLoadingBridgeType) UnstageVoid(stage *Stage) {
	codeloadingbridgetype.Unstage(stage)
}

func (codeloadingbridgetype *CodeLoadingBridgeType) StageVoid(stage *Stage) {
	codeloadingbridgetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeloadingbridgetype *CodeLoadingBridgeType) GetName() (res string) {
	return codeloadingbridgetype.Name
}

// for satisfaction of GongStruct interface
func (codeloadingbridgetype *CodeLoadingBridgeType) SetName(name string) {
	codeloadingbridgetype.Name = name
}

// Stage puts codelogicaloperatortype to the model stage
func (codelogicaloperatortype *CodeLogicalOperatorType) Stage(stage *Stage) *CodeLogicalOperatorType {
	__gong__stage(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorType_stagedOrder, stage.CodeLogicalOperatorType_orderStaged, &stage.CodeLogicalOperatorTypeOrder, stage.CodeLogicalOperatorTypes_mapString, codelogicaloperatortype, codelogicaloperatortype.Name)
	return codelogicaloperatortype
}

// StagePreserveOrder puts codelogicaloperatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeLogicalOperatorTypeOrder
// - update stage.CodeLogicalOperatorTypeOrder accordingly
func (codelogicaloperatortype *CodeLogicalOperatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorType_stagedOrder, stage.CodeLogicalOperatorType_orderStaged, &stage.CodeLogicalOperatorTypeOrder, stage.CodeLogicalOperatorTypes_mapString, codelogicaloperatortype, order, codelogicaloperatortype.Name)
}

// Unstage removes codelogicaloperatortype off the model stage
func (codelogicaloperatortype *CodeLogicalOperatorType) Unstage(stage *Stage) *CodeLogicalOperatorType {
	__gong__unstage(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorTypes_mapString, codelogicaloperatortype, codelogicaloperatortype.Name)
	return codelogicaloperatortype
}

// UnstageVoid removes codelogicaloperatortype off the model stage
func (codelogicaloperatortype *CodeLogicalOperatorType) UnstageVoid(stage *Stage) {
	codelogicaloperatortype.Unstage(stage)
}

func (codelogicaloperatortype *CodeLogicalOperatorType) StageVoid(stage *Stage) {
	codelogicaloperatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codelogicaloperatortype *CodeLogicalOperatorType) GetName() (res string) {
	return codelogicaloperatortype.Name
}

// for satisfaction of GongStruct interface
func (codelogicaloperatortype *CodeLogicalOperatorType) SetName(name string) {
	codelogicaloperatortype.Name = name
}

// Stage puts codemarkingconditiontype to the model stage
func (codemarkingconditiontype *CodeMarkingConditionType) Stage(stage *Stage) *CodeMarkingConditionType {
	__gong__stage(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionType_stagedOrder, stage.CodeMarkingConditionType_orderStaged, &stage.CodeMarkingConditionTypeOrder, stage.CodeMarkingConditionTypes_mapString, codemarkingconditiontype, codemarkingconditiontype.Name)
	return codemarkingconditiontype
}

// StagePreserveOrder puts codemarkingconditiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMarkingConditionTypeOrder
// - update stage.CodeMarkingConditionTypeOrder accordingly
func (codemarkingconditiontype *CodeMarkingConditionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionType_stagedOrder, stage.CodeMarkingConditionType_orderStaged, &stage.CodeMarkingConditionTypeOrder, stage.CodeMarkingConditionTypes_mapString, codemarkingconditiontype, order, codemarkingconditiontype.Name)
}

// Unstage removes codemarkingconditiontype off the model stage
func (codemarkingconditiontype *CodeMarkingConditionType) Unstage(stage *Stage) *CodeMarkingConditionType {
	__gong__unstage(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionTypes_mapString, codemarkingconditiontype, codemarkingconditiontype.Name)
	return codemarkingconditiontype
}

// UnstageVoid removes codemarkingconditiontype off the model stage
func (codemarkingconditiontype *CodeMarkingConditionType) UnstageVoid(stage *Stage) {
	codemarkingconditiontype.Unstage(stage)
}

func (codemarkingconditiontype *CodeMarkingConditionType) StageVoid(stage *Stage) {
	codemarkingconditiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemarkingconditiontype *CodeMarkingConditionType) GetName() (res string) {
	return codemarkingconditiontype.Name
}

// for satisfaction of GongStruct interface
func (codemarkingconditiontype *CodeMarkingConditionType) SetName(name string) {
	codemarkingconditiontype.Name = name
}

// Stage puts codemarkingstyletype to the model stage
func (codemarkingstyletype *CodeMarkingStyleType) Stage(stage *Stage) *CodeMarkingStyleType {
	__gong__stage(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleType_stagedOrder, stage.CodeMarkingStyleType_orderStaged, &stage.CodeMarkingStyleTypeOrder, stage.CodeMarkingStyleTypes_mapString, codemarkingstyletype, codemarkingstyletype.Name)
	return codemarkingstyletype
}

// StagePreserveOrder puts codemarkingstyletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeMarkingStyleTypeOrder
// - update stage.CodeMarkingStyleTypeOrder accordingly
func (codemarkingstyletype *CodeMarkingStyleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleType_stagedOrder, stage.CodeMarkingStyleType_orderStaged, &stage.CodeMarkingStyleTypeOrder, stage.CodeMarkingStyleTypes_mapString, codemarkingstyletype, order, codemarkingstyletype.Name)
}

// Unstage removes codemarkingstyletype off the model stage
func (codemarkingstyletype *CodeMarkingStyleType) Unstage(stage *Stage) *CodeMarkingStyleType {
	__gong__unstage(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleTypes_mapString, codemarkingstyletype, codemarkingstyletype.Name)
	return codemarkingstyletype
}

// UnstageVoid removes codemarkingstyletype off the model stage
func (codemarkingstyletype *CodeMarkingStyleType) UnstageVoid(stage *Stage) {
	codemarkingstyletype.Unstage(stage)
}

func (codemarkingstyletype *CodeMarkingStyleType) StageVoid(stage *Stage) {
	codemarkingstyletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codemarkingstyletype *CodeMarkingStyleType) GetName() (res string) {
	return codemarkingstyletype.Name
}

// for satisfaction of GongStruct interface
func (codemarkingstyletype *CodeMarkingStyleType) SetName(name string) {
	codemarkingstyletype.Name = name
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

// Stage puts codeoperationairportheliporttype to the model stage
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) Stage(stage *Stage) *CodeOperationAirportHeliportType {
	__gong__stage(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportType_stagedOrder, stage.CodeOperationAirportHeliportType_orderStaged, &stage.CodeOperationAirportHeliportTypeOrder, stage.CodeOperationAirportHeliportTypes_mapString, codeoperationairportheliporttype, codeoperationairportheliporttype.Name)
	return codeoperationairportheliporttype
}

// StagePreserveOrder puts codeoperationairportheliporttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOperationAirportHeliportTypeOrder
// - update stage.CodeOperationAirportHeliportTypeOrder accordingly
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportType_stagedOrder, stage.CodeOperationAirportHeliportType_orderStaged, &stage.CodeOperationAirportHeliportTypeOrder, stage.CodeOperationAirportHeliportTypes_mapString, codeoperationairportheliporttype, order, codeoperationairportheliporttype.Name)
}

// Unstage removes codeoperationairportheliporttype off the model stage
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) Unstage(stage *Stage) *CodeOperationAirportHeliportType {
	__gong__unstage(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportTypes_mapString, codeoperationairportheliporttype, codeoperationairportheliporttype.Name)
	return codeoperationairportheliporttype
}

// UnstageVoid removes codeoperationairportheliporttype off the model stage
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) UnstageVoid(stage *Stage) {
	codeoperationairportheliporttype.Unstage(stage)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) StageVoid(stage *Stage) {
	codeoperationairportheliporttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GetName() (res string) {
	return codeoperationairportheliporttype.Name
}

// for satisfaction of GongStruct interface
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) SetName(name string) {
	codeoperationairportheliporttype.Name = name
}

// Stage puts codeoperationmanoeuvringareatype to the model stage
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) Stage(stage *Stage) *CodeOperationManoeuvringAreaType {
	__gong__stage(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaType_stagedOrder, stage.CodeOperationManoeuvringAreaType_orderStaged, &stage.CodeOperationManoeuvringAreaTypeOrder, stage.CodeOperationManoeuvringAreaTypes_mapString, codeoperationmanoeuvringareatype, codeoperationmanoeuvringareatype.Name)
	return codeoperationmanoeuvringareatype
}

// StagePreserveOrder puts codeoperationmanoeuvringareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeOperationManoeuvringAreaTypeOrder
// - update stage.CodeOperationManoeuvringAreaTypeOrder accordingly
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaType_stagedOrder, stage.CodeOperationManoeuvringAreaType_orderStaged, &stage.CodeOperationManoeuvringAreaTypeOrder, stage.CodeOperationManoeuvringAreaTypes_mapString, codeoperationmanoeuvringareatype, order, codeoperationmanoeuvringareatype.Name)
}

// Unstage removes codeoperationmanoeuvringareatype off the model stage
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) Unstage(stage *Stage) *CodeOperationManoeuvringAreaType {
	__gong__unstage(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaTypes_mapString, codeoperationmanoeuvringareatype, codeoperationmanoeuvringareatype.Name)
	return codeoperationmanoeuvringareatype
}

// UnstageVoid removes codeoperationmanoeuvringareatype off the model stage
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) UnstageVoid(stage *Stage) {
	codeoperationmanoeuvringareatype.Unstage(stage)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) StageVoid(stage *Stage) {
	codeoperationmanoeuvringareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GetName() (res string) {
	return codeoperationmanoeuvringareatype.Name
}

// for satisfaction of GongStruct interface
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) SetName(name string) {
	codeoperationmanoeuvringareatype.Name = name
}

// Stage puts codepcnmethodtype to the model stage
func (codepcnmethodtype *CodePCNMethodType) Stage(stage *Stage) *CodePCNMethodType {
	__gong__stage(stage.CodePCNMethodTypes, stage.CodePCNMethodType_stagedOrder, stage.CodePCNMethodType_orderStaged, &stage.CodePCNMethodTypeOrder, stage.CodePCNMethodTypes_mapString, codepcnmethodtype, codepcnmethodtype.Name)
	return codepcnmethodtype
}

// StagePreserveOrder puts codepcnmethodtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePCNMethodTypeOrder
// - update stage.CodePCNMethodTypeOrder accordingly
func (codepcnmethodtype *CodePCNMethodType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePCNMethodTypes, stage.CodePCNMethodType_stagedOrder, stage.CodePCNMethodType_orderStaged, &stage.CodePCNMethodTypeOrder, stage.CodePCNMethodTypes_mapString, codepcnmethodtype, order, codepcnmethodtype.Name)
}

// Unstage removes codepcnmethodtype off the model stage
func (codepcnmethodtype *CodePCNMethodType) Unstage(stage *Stage) *CodePCNMethodType {
	__gong__unstage(stage.CodePCNMethodTypes, stage.CodePCNMethodTypes_mapString, codepcnmethodtype, codepcnmethodtype.Name)
	return codepcnmethodtype
}

// UnstageVoid removes codepcnmethodtype off the model stage
func (codepcnmethodtype *CodePCNMethodType) UnstageVoid(stage *Stage) {
	codepcnmethodtype.Unstage(stage)
}

func (codepcnmethodtype *CodePCNMethodType) StageVoid(stage *Stage) {
	codepcnmethodtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepcnmethodtype *CodePCNMethodType) GetName() (res string) {
	return codepcnmethodtype.Name
}

// for satisfaction of GongStruct interface
func (codepcnmethodtype *CodePCNMethodType) SetName(name string) {
	codepcnmethodtype.Name = name
}

// Stage puts codepcnpavementtype to the model stage
func (codepcnpavementtype *CodePCNPavementType) Stage(stage *Stage) *CodePCNPavementType {
	__gong__stage(stage.CodePCNPavementTypes, stage.CodePCNPavementType_stagedOrder, stage.CodePCNPavementType_orderStaged, &stage.CodePCNPavementTypeOrder, stage.CodePCNPavementTypes_mapString, codepcnpavementtype, codepcnpavementtype.Name)
	return codepcnpavementtype
}

// StagePreserveOrder puts codepcnpavementtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePCNPavementTypeOrder
// - update stage.CodePCNPavementTypeOrder accordingly
func (codepcnpavementtype *CodePCNPavementType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePCNPavementTypes, stage.CodePCNPavementType_stagedOrder, stage.CodePCNPavementType_orderStaged, &stage.CodePCNPavementTypeOrder, stage.CodePCNPavementTypes_mapString, codepcnpavementtype, order, codepcnpavementtype.Name)
}

// Unstage removes codepcnpavementtype off the model stage
func (codepcnpavementtype *CodePCNPavementType) Unstage(stage *Stage) *CodePCNPavementType {
	__gong__unstage(stage.CodePCNPavementTypes, stage.CodePCNPavementTypes_mapString, codepcnpavementtype, codepcnpavementtype.Name)
	return codepcnpavementtype
}

// UnstageVoid removes codepcnpavementtype off the model stage
func (codepcnpavementtype *CodePCNPavementType) UnstageVoid(stage *Stage) {
	codepcnpavementtype.Unstage(stage)
}

func (codepcnpavementtype *CodePCNPavementType) StageVoid(stage *Stage) {
	codepcnpavementtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepcnpavementtype *CodePCNPavementType) GetName() (res string) {
	return codepcnpavementtype.Name
}

// for satisfaction of GongStruct interface
func (codepcnpavementtype *CodePCNPavementType) SetName(name string) {
	codepcnpavementtype.Name = name
}

// Stage puts codepcnsubgradetype to the model stage
func (codepcnsubgradetype *CodePCNSubgradeType) Stage(stage *Stage) *CodePCNSubgradeType {
	__gong__stage(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeType_stagedOrder, stage.CodePCNSubgradeType_orderStaged, &stage.CodePCNSubgradeTypeOrder, stage.CodePCNSubgradeTypes_mapString, codepcnsubgradetype, codepcnsubgradetype.Name)
	return codepcnsubgradetype
}

// StagePreserveOrder puts codepcnsubgradetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePCNSubgradeTypeOrder
// - update stage.CodePCNSubgradeTypeOrder accordingly
func (codepcnsubgradetype *CodePCNSubgradeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeType_stagedOrder, stage.CodePCNSubgradeType_orderStaged, &stage.CodePCNSubgradeTypeOrder, stage.CodePCNSubgradeTypes_mapString, codepcnsubgradetype, order, codepcnsubgradetype.Name)
}

// Unstage removes codepcnsubgradetype off the model stage
func (codepcnsubgradetype *CodePCNSubgradeType) Unstage(stage *Stage) *CodePCNSubgradeType {
	__gong__unstage(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeTypes_mapString, codepcnsubgradetype, codepcnsubgradetype.Name)
	return codepcnsubgradetype
}

// UnstageVoid removes codepcnsubgradetype off the model stage
func (codepcnsubgradetype *CodePCNSubgradeType) UnstageVoid(stage *Stage) {
	codepcnsubgradetype.Unstage(stage)
}

func (codepcnsubgradetype *CodePCNSubgradeType) StageVoid(stage *Stage) {
	codepcnsubgradetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepcnsubgradetype *CodePCNSubgradeType) GetName() (res string) {
	return codepcnsubgradetype.Name
}

// for satisfaction of GongStruct interface
func (codepcnsubgradetype *CodePCNSubgradeType) SetName(name string) {
	codepcnsubgradetype.Name = name
}

// Stage puts codepcntyrepressuretype to the model stage
func (codepcntyrepressuretype *CodePCNTyrePressureType) Stage(stage *Stage) *CodePCNTyrePressureType {
	__gong__stage(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureType_stagedOrder, stage.CodePCNTyrePressureType_orderStaged, &stage.CodePCNTyrePressureTypeOrder, stage.CodePCNTyrePressureTypes_mapString, codepcntyrepressuretype, codepcntyrepressuretype.Name)
	return codepcntyrepressuretype
}

// StagePreserveOrder puts codepcntyrepressuretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodePCNTyrePressureTypeOrder
// - update stage.CodePCNTyrePressureTypeOrder accordingly
func (codepcntyrepressuretype *CodePCNTyrePressureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureType_stagedOrder, stage.CodePCNTyrePressureType_orderStaged, &stage.CodePCNTyrePressureTypeOrder, stage.CodePCNTyrePressureTypes_mapString, codepcntyrepressuretype, order, codepcntyrepressuretype.Name)
}

// Unstage removes codepcntyrepressuretype off the model stage
func (codepcntyrepressuretype *CodePCNTyrePressureType) Unstage(stage *Stage) *CodePCNTyrePressureType {
	__gong__unstage(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureTypes_mapString, codepcntyrepressuretype, codepcntyrepressuretype.Name)
	return codepcntyrepressuretype
}

// UnstageVoid removes codepcntyrepressuretype off the model stage
func (codepcntyrepressuretype *CodePCNTyrePressureType) UnstageVoid(stage *Stage) {
	codepcntyrepressuretype.Unstage(stage)
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) StageVoid(stage *Stage) {
	codepcntyrepressuretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codepcntyrepressuretype *CodePCNTyrePressureType) GetName() (res string) {
	return codepcntyrepressuretype.Name
}

// for satisfaction of GongStruct interface
func (codepcntyrepressuretype *CodePCNTyrePressureType) SetName(name string) {
	codepcntyrepressuretype.Name = name
}

// Stage puts codeprotectareasectiontype to the model stage
func (codeprotectareasectiontype *CodeProtectAreaSectionType) Stage(stage *Stage) *CodeProtectAreaSectionType {
	__gong__stage(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionType_stagedOrder, stage.CodeProtectAreaSectionType_orderStaged, &stage.CodeProtectAreaSectionTypeOrder, stage.CodeProtectAreaSectionTypes_mapString, codeprotectareasectiontype, codeprotectareasectiontype.Name)
	return codeprotectareasectiontype
}

// StagePreserveOrder puts codeprotectareasectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeProtectAreaSectionTypeOrder
// - update stage.CodeProtectAreaSectionTypeOrder accordingly
func (codeprotectareasectiontype *CodeProtectAreaSectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionType_stagedOrder, stage.CodeProtectAreaSectionType_orderStaged, &stage.CodeProtectAreaSectionTypeOrder, stage.CodeProtectAreaSectionTypes_mapString, codeprotectareasectiontype, order, codeprotectareasectiontype.Name)
}

// Unstage removes codeprotectareasectiontype off the model stage
func (codeprotectareasectiontype *CodeProtectAreaSectionType) Unstage(stage *Stage) *CodeProtectAreaSectionType {
	__gong__unstage(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionTypes_mapString, codeprotectareasectiontype, codeprotectareasectiontype.Name)
	return codeprotectareasectiontype
}

// UnstageVoid removes codeprotectareasectiontype off the model stage
func (codeprotectareasectiontype *CodeProtectAreaSectionType) UnstageVoid(stage *Stage) {
	codeprotectareasectiontype.Unstage(stage)
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) StageVoid(stage *Stage) {
	codeprotectareasectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeprotectareasectiontype *CodeProtectAreaSectionType) GetName() (res string) {
	return codeprotectareasectiontype.Name
}

// for satisfaction of GongStruct interface
func (codeprotectareasectiontype *CodeProtectAreaSectionType) SetName(name string) {
	codeprotectareasectiontype.Name = name
}

// Stage puts codervrreadingtype to the model stage
func (codervrreadingtype *CodeRVRReadingType) Stage(stage *Stage) *CodeRVRReadingType {
	__gong__stage(stage.CodeRVRReadingTypes, stage.CodeRVRReadingType_stagedOrder, stage.CodeRVRReadingType_orderStaged, &stage.CodeRVRReadingTypeOrder, stage.CodeRVRReadingTypes_mapString, codervrreadingtype, codervrreadingtype.Name)
	return codervrreadingtype
}

// StagePreserveOrder puts codervrreadingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRVRReadingTypeOrder
// - update stage.CodeRVRReadingTypeOrder accordingly
func (codervrreadingtype *CodeRVRReadingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRVRReadingTypes, stage.CodeRVRReadingType_stagedOrder, stage.CodeRVRReadingType_orderStaged, &stage.CodeRVRReadingTypeOrder, stage.CodeRVRReadingTypes_mapString, codervrreadingtype, order, codervrreadingtype.Name)
}

// Unstage removes codervrreadingtype off the model stage
func (codervrreadingtype *CodeRVRReadingType) Unstage(stage *Stage) *CodeRVRReadingType {
	__gong__unstage(stage.CodeRVRReadingTypes, stage.CodeRVRReadingTypes_mapString, codervrreadingtype, codervrreadingtype.Name)
	return codervrreadingtype
}

// UnstageVoid removes codervrreadingtype off the model stage
func (codervrreadingtype *CodeRVRReadingType) UnstageVoid(stage *Stage) {
	codervrreadingtype.Unstage(stage)
}

func (codervrreadingtype *CodeRVRReadingType) StageVoid(stage *Stage) {
	codervrreadingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codervrreadingtype *CodeRVRReadingType) GetName() (res string) {
	return codervrreadingtype.Name
}

// for satisfaction of GongStruct interface
func (codervrreadingtype *CodeRVRReadingType) SetName(name string) {
	codervrreadingtype.Name = name
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

// Stage puts coderoadtype to the model stage
func (coderoadtype *CodeRoadType) Stage(stage *Stage) *CodeRoadType {
	__gong__stage(stage.CodeRoadTypes, stage.CodeRoadType_stagedOrder, stage.CodeRoadType_orderStaged, &stage.CodeRoadTypeOrder, stage.CodeRoadTypes_mapString, coderoadtype, coderoadtype.Name)
	return coderoadtype
}

// StagePreserveOrder puts coderoadtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRoadTypeOrder
// - update stage.CodeRoadTypeOrder accordingly
func (coderoadtype *CodeRoadType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRoadTypes, stage.CodeRoadType_stagedOrder, stage.CodeRoadType_orderStaged, &stage.CodeRoadTypeOrder, stage.CodeRoadTypes_mapString, coderoadtype, order, coderoadtype.Name)
}

// Unstage removes coderoadtype off the model stage
func (coderoadtype *CodeRoadType) Unstage(stage *Stage) *CodeRoadType {
	__gong__unstage(stage.CodeRoadTypes, stage.CodeRoadTypes_mapString, coderoadtype, coderoadtype.Name)
	return coderoadtype
}

// UnstageVoid removes coderoadtype off the model stage
func (coderoadtype *CodeRoadType) UnstageVoid(stage *Stage) {
	coderoadtype.Unstage(stage)
}

func (coderoadtype *CodeRoadType) StageVoid(stage *Stage) {
	coderoadtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderoadtype *CodeRoadType) GetName() (res string) {
	return coderoadtype.Name
}

// for satisfaction of GongStruct interface
func (coderoadtype *CodeRoadType) SetName(name string) {
	coderoadtype.Name = name
}

// Stage puts coderunwayelementtype to the model stage
func (coderunwayelementtype *CodeRunwayElementType) Stage(stage *Stage) *CodeRunwayElementType {
	__gong__stage(stage.CodeRunwayElementTypes, stage.CodeRunwayElementType_stagedOrder, stage.CodeRunwayElementType_orderStaged, &stage.CodeRunwayElementTypeOrder, stage.CodeRunwayElementTypes_mapString, coderunwayelementtype, coderunwayelementtype.Name)
	return coderunwayelementtype
}

// StagePreserveOrder puts coderunwayelementtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwayElementTypeOrder
// - update stage.CodeRunwayElementTypeOrder accordingly
func (coderunwayelementtype *CodeRunwayElementType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwayElementTypes, stage.CodeRunwayElementType_stagedOrder, stage.CodeRunwayElementType_orderStaged, &stage.CodeRunwayElementTypeOrder, stage.CodeRunwayElementTypes_mapString, coderunwayelementtype, order, coderunwayelementtype.Name)
}

// Unstage removes coderunwayelementtype off the model stage
func (coderunwayelementtype *CodeRunwayElementType) Unstage(stage *Stage) *CodeRunwayElementType {
	__gong__unstage(stage.CodeRunwayElementTypes, stage.CodeRunwayElementTypes_mapString, coderunwayelementtype, coderunwayelementtype.Name)
	return coderunwayelementtype
}

// UnstageVoid removes coderunwayelementtype off the model stage
func (coderunwayelementtype *CodeRunwayElementType) UnstageVoid(stage *Stage) {
	coderunwayelementtype.Unstage(stage)
}

func (coderunwayelementtype *CodeRunwayElementType) StageVoid(stage *Stage) {
	coderunwayelementtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwayelementtype *CodeRunwayElementType) GetName() (res string) {
	return coderunwayelementtype.Name
}

// for satisfaction of GongStruct interface
func (coderunwayelementtype *CodeRunwayElementType) SetName(name string) {
	coderunwayelementtype.Name = name
}

// Stage puts coderunwaymarkingtype to the model stage
func (coderunwaymarkingtype *CodeRunwayMarkingType) Stage(stage *Stage) *CodeRunwayMarkingType {
	__gong__stage(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingType_stagedOrder, stage.CodeRunwayMarkingType_orderStaged, &stage.CodeRunwayMarkingTypeOrder, stage.CodeRunwayMarkingTypes_mapString, coderunwaymarkingtype, coderunwaymarkingtype.Name)
	return coderunwaymarkingtype
}

// StagePreserveOrder puts coderunwaymarkingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwayMarkingTypeOrder
// - update stage.CodeRunwayMarkingTypeOrder accordingly
func (coderunwaymarkingtype *CodeRunwayMarkingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingType_stagedOrder, stage.CodeRunwayMarkingType_orderStaged, &stage.CodeRunwayMarkingTypeOrder, stage.CodeRunwayMarkingTypes_mapString, coderunwaymarkingtype, order, coderunwaymarkingtype.Name)
}

// Unstage removes coderunwaymarkingtype off the model stage
func (coderunwaymarkingtype *CodeRunwayMarkingType) Unstage(stage *Stage) *CodeRunwayMarkingType {
	__gong__unstage(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingTypes_mapString, coderunwaymarkingtype, coderunwaymarkingtype.Name)
	return coderunwaymarkingtype
}

// UnstageVoid removes coderunwaymarkingtype off the model stage
func (coderunwaymarkingtype *CodeRunwayMarkingType) UnstageVoid(stage *Stage) {
	coderunwaymarkingtype.Unstage(stage)
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) StageVoid(stage *Stage) {
	coderunwaymarkingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwaymarkingtype *CodeRunwayMarkingType) GetName() (res string) {
	return coderunwaymarkingtype.Name
}

// for satisfaction of GongStruct interface
func (coderunwaymarkingtype *CodeRunwayMarkingType) SetName(name string) {
	coderunwaymarkingtype.Name = name
}

// Stage puts coderunwaypointroletype to the model stage
func (coderunwaypointroletype *CodeRunwayPointRoleType) Stage(stage *Stage) *CodeRunwayPointRoleType {
	__gong__stage(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleType_stagedOrder, stage.CodeRunwayPointRoleType_orderStaged, &stage.CodeRunwayPointRoleTypeOrder, stage.CodeRunwayPointRoleTypes_mapString, coderunwaypointroletype, coderunwaypointroletype.Name)
	return coderunwaypointroletype
}

// StagePreserveOrder puts coderunwaypointroletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwayPointRoleTypeOrder
// - update stage.CodeRunwayPointRoleTypeOrder accordingly
func (coderunwaypointroletype *CodeRunwayPointRoleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleType_stagedOrder, stage.CodeRunwayPointRoleType_orderStaged, &stage.CodeRunwayPointRoleTypeOrder, stage.CodeRunwayPointRoleTypes_mapString, coderunwaypointroletype, order, coderunwaypointroletype.Name)
}

// Unstage removes coderunwaypointroletype off the model stage
func (coderunwaypointroletype *CodeRunwayPointRoleType) Unstage(stage *Stage) *CodeRunwayPointRoleType {
	__gong__unstage(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleTypes_mapString, coderunwaypointroletype, coderunwaypointroletype.Name)
	return coderunwaypointroletype
}

// UnstageVoid removes coderunwaypointroletype off the model stage
func (coderunwaypointroletype *CodeRunwayPointRoleType) UnstageVoid(stage *Stage) {
	coderunwaypointroletype.Unstage(stage)
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) StageVoid(stage *Stage) {
	coderunwaypointroletype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwaypointroletype *CodeRunwayPointRoleType) GetName() (res string) {
	return coderunwaypointroletype.Name
}

// for satisfaction of GongStruct interface
func (coderunwaypointroletype *CodeRunwayPointRoleType) SetName(name string) {
	coderunwaypointroletype.Name = name
}

// Stage puts coderunwayprotectionareatype to the model stage
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) Stage(stage *Stage) *CodeRunwayProtectionAreaType {
	__gong__stage(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaType_stagedOrder, stage.CodeRunwayProtectionAreaType_orderStaged, &stage.CodeRunwayProtectionAreaTypeOrder, stage.CodeRunwayProtectionAreaTypes_mapString, coderunwayprotectionareatype, coderunwayprotectionareatype.Name)
	return coderunwayprotectionareatype
}

// StagePreserveOrder puts coderunwayprotectionareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwayProtectionAreaTypeOrder
// - update stage.CodeRunwayProtectionAreaTypeOrder accordingly
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaType_stagedOrder, stage.CodeRunwayProtectionAreaType_orderStaged, &stage.CodeRunwayProtectionAreaTypeOrder, stage.CodeRunwayProtectionAreaTypes_mapString, coderunwayprotectionareatype, order, coderunwayprotectionareatype.Name)
}

// Unstage removes coderunwayprotectionareatype off the model stage
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) Unstage(stage *Stage) *CodeRunwayProtectionAreaType {
	__gong__unstage(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaTypes_mapString, coderunwayprotectionareatype, coderunwayprotectionareatype.Name)
	return coderunwayprotectionareatype
}

// UnstageVoid removes coderunwayprotectionareatype off the model stage
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) UnstageVoid(stage *Stage) {
	coderunwayprotectionareatype.Unstage(stage)
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) StageVoid(stage *Stage) {
	coderunwayprotectionareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GetName() (res string) {
	return coderunwayprotectionareatype.Name
}

// for satisfaction of GongStruct interface
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) SetName(name string) {
	coderunwayprotectionareatype.Name = name
}

// Stage puts coderunwaysectiontype to the model stage
func (coderunwaysectiontype *CodeRunwaySectionType) Stage(stage *Stage) *CodeRunwaySectionType {
	__gong__stage(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionType_stagedOrder, stage.CodeRunwaySectionType_orderStaged, &stage.CodeRunwaySectionTypeOrder, stage.CodeRunwaySectionTypes_mapString, coderunwaysectiontype, coderunwaysectiontype.Name)
	return coderunwaysectiontype
}

// StagePreserveOrder puts coderunwaysectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwaySectionTypeOrder
// - update stage.CodeRunwaySectionTypeOrder accordingly
func (coderunwaysectiontype *CodeRunwaySectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionType_stagedOrder, stage.CodeRunwaySectionType_orderStaged, &stage.CodeRunwaySectionTypeOrder, stage.CodeRunwaySectionTypes_mapString, coderunwaysectiontype, order, coderunwaysectiontype.Name)
}

// Unstage removes coderunwaysectiontype off the model stage
func (coderunwaysectiontype *CodeRunwaySectionType) Unstage(stage *Stage) *CodeRunwaySectionType {
	__gong__unstage(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionTypes_mapString, coderunwaysectiontype, coderunwaysectiontype.Name)
	return coderunwaysectiontype
}

// UnstageVoid removes coderunwaysectiontype off the model stage
func (coderunwaysectiontype *CodeRunwaySectionType) UnstageVoid(stage *Stage) {
	coderunwaysectiontype.Unstage(stage)
}

func (coderunwaysectiontype *CodeRunwaySectionType) StageVoid(stage *Stage) {
	coderunwaysectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwaysectiontype *CodeRunwaySectionType) GetName() (res string) {
	return coderunwaysectiontype.Name
}

// for satisfaction of GongStruct interface
func (coderunwaysectiontype *CodeRunwaySectionType) SetName(name string) {
	coderunwaysectiontype.Name = name
}

// Stage puts coderunwaytype to the model stage
func (coderunwaytype *CodeRunwayType) Stage(stage *Stage) *CodeRunwayType {
	__gong__stage(stage.CodeRunwayTypes, stage.CodeRunwayType_stagedOrder, stage.CodeRunwayType_orderStaged, &stage.CodeRunwayTypeOrder, stage.CodeRunwayTypes_mapString, coderunwaytype, coderunwaytype.Name)
	return coderunwaytype
}

// StagePreserveOrder puts coderunwaytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeRunwayTypeOrder
// - update stage.CodeRunwayTypeOrder accordingly
func (coderunwaytype *CodeRunwayType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeRunwayTypes, stage.CodeRunwayType_stagedOrder, stage.CodeRunwayType_orderStaged, &stage.CodeRunwayTypeOrder, stage.CodeRunwayTypes_mapString, coderunwaytype, order, coderunwaytype.Name)
}

// Unstage removes coderunwaytype off the model stage
func (coderunwaytype *CodeRunwayType) Unstage(stage *Stage) *CodeRunwayType {
	__gong__unstage(stage.CodeRunwayTypes, stage.CodeRunwayTypes_mapString, coderunwaytype, coderunwaytype.Name)
	return coderunwaytype
}

// UnstageVoid removes coderunwaytype off the model stage
func (coderunwaytype *CodeRunwayType) UnstageVoid(stage *Stage) {
	coderunwaytype.Unstage(stage)
}

func (coderunwaytype *CodeRunwayType) StageVoid(stage *Stage) {
	coderunwaytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coderunwaytype *CodeRunwayType) GetName() (res string) {
	return coderunwaytype.Name
}

// for satisfaction of GongStruct interface
func (coderunwaytype *CodeRunwayType) SetName(name string) {
	coderunwaytype.Name = name
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

// Stage puts codestatusairporttype to the model stage
func (codestatusairporttype *CodeStatusAirportType) Stage(stage *Stage) *CodeStatusAirportType {
	__gong__stage(stage.CodeStatusAirportTypes, stage.CodeStatusAirportType_stagedOrder, stage.CodeStatusAirportType_orderStaged, &stage.CodeStatusAirportTypeOrder, stage.CodeStatusAirportTypes_mapString, codestatusairporttype, codestatusairporttype.Name)
	return codestatusairporttype
}

// StagePreserveOrder puts codestatusairporttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusAirportTypeOrder
// - update stage.CodeStatusAirportTypeOrder accordingly
func (codestatusairporttype *CodeStatusAirportType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusAirportTypes, stage.CodeStatusAirportType_stagedOrder, stage.CodeStatusAirportType_orderStaged, &stage.CodeStatusAirportTypeOrder, stage.CodeStatusAirportTypes_mapString, codestatusairporttype, order, codestatusairporttype.Name)
}

// Unstage removes codestatusairporttype off the model stage
func (codestatusairporttype *CodeStatusAirportType) Unstage(stage *Stage) *CodeStatusAirportType {
	__gong__unstage(stage.CodeStatusAirportTypes, stage.CodeStatusAirportTypes_mapString, codestatusairporttype, codestatusairporttype.Name)
	return codestatusairporttype
}

// UnstageVoid removes codestatusairporttype off the model stage
func (codestatusairporttype *CodeStatusAirportType) UnstageVoid(stage *Stage) {
	codestatusairporttype.Unstage(stage)
}

func (codestatusairporttype *CodeStatusAirportType) StageVoid(stage *Stage) {
	codestatusairporttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusairporttype *CodeStatusAirportType) GetName() (res string) {
	return codestatusairporttype.Name
}

// for satisfaction of GongStruct interface
func (codestatusairporttype *CodeStatusAirportType) SetName(name string) {
	codestatusairporttype.Name = name
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

// Stage puts codesurfacecompositiontype to the model stage
func (codesurfacecompositiontype *CodeSurfaceCompositionType) Stage(stage *Stage) *CodeSurfaceCompositionType {
	__gong__stage(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionType_stagedOrder, stage.CodeSurfaceCompositionType_orderStaged, &stage.CodeSurfaceCompositionTypeOrder, stage.CodeSurfaceCompositionTypes_mapString, codesurfacecompositiontype, codesurfacecompositiontype.Name)
	return codesurfacecompositiontype
}

// StagePreserveOrder puts codesurfacecompositiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSurfaceCompositionTypeOrder
// - update stage.CodeSurfaceCompositionTypeOrder accordingly
func (codesurfacecompositiontype *CodeSurfaceCompositionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionType_stagedOrder, stage.CodeSurfaceCompositionType_orderStaged, &stage.CodeSurfaceCompositionTypeOrder, stage.CodeSurfaceCompositionTypes_mapString, codesurfacecompositiontype, order, codesurfacecompositiontype.Name)
}

// Unstage removes codesurfacecompositiontype off the model stage
func (codesurfacecompositiontype *CodeSurfaceCompositionType) Unstage(stage *Stage) *CodeSurfaceCompositionType {
	__gong__unstage(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionTypes_mapString, codesurfacecompositiontype, codesurfacecompositiontype.Name)
	return codesurfacecompositiontype
}

// UnstageVoid removes codesurfacecompositiontype off the model stage
func (codesurfacecompositiontype *CodeSurfaceCompositionType) UnstageVoid(stage *Stage) {
	codesurfacecompositiontype.Unstage(stage)
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) StageVoid(stage *Stage) {
	codesurfacecompositiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesurfacecompositiontype *CodeSurfaceCompositionType) GetName() (res string) {
	return codesurfacecompositiontype.Name
}

// for satisfaction of GongStruct interface
func (codesurfacecompositiontype *CodeSurfaceCompositionType) SetName(name string) {
	codesurfacecompositiontype.Name = name
}

// Stage puts codesurfaceconditiontype to the model stage
func (codesurfaceconditiontype *CodeSurfaceConditionType) Stage(stage *Stage) *CodeSurfaceConditionType {
	__gong__stage(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionType_stagedOrder, stage.CodeSurfaceConditionType_orderStaged, &stage.CodeSurfaceConditionTypeOrder, stage.CodeSurfaceConditionTypes_mapString, codesurfaceconditiontype, codesurfaceconditiontype.Name)
	return codesurfaceconditiontype
}

// StagePreserveOrder puts codesurfaceconditiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSurfaceConditionTypeOrder
// - update stage.CodeSurfaceConditionTypeOrder accordingly
func (codesurfaceconditiontype *CodeSurfaceConditionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionType_stagedOrder, stage.CodeSurfaceConditionType_orderStaged, &stage.CodeSurfaceConditionTypeOrder, stage.CodeSurfaceConditionTypes_mapString, codesurfaceconditiontype, order, codesurfaceconditiontype.Name)
}

// Unstage removes codesurfaceconditiontype off the model stage
func (codesurfaceconditiontype *CodeSurfaceConditionType) Unstage(stage *Stage) *CodeSurfaceConditionType {
	__gong__unstage(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionTypes_mapString, codesurfaceconditiontype, codesurfaceconditiontype.Name)
	return codesurfaceconditiontype
}

// UnstageVoid removes codesurfaceconditiontype off the model stage
func (codesurfaceconditiontype *CodeSurfaceConditionType) UnstageVoid(stage *Stage) {
	codesurfaceconditiontype.Unstage(stage)
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) StageVoid(stage *Stage) {
	codesurfaceconditiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesurfaceconditiontype *CodeSurfaceConditionType) GetName() (res string) {
	return codesurfaceconditiontype.Name
}

// for satisfaction of GongStruct interface
func (codesurfaceconditiontype *CodeSurfaceConditionType) SetName(name string) {
	codesurfaceconditiontype.Name = name
}

// Stage puts codesurfacepreparationtype to the model stage
func (codesurfacepreparationtype *CodeSurfacePreparationType) Stage(stage *Stage) *CodeSurfacePreparationType {
	__gong__stage(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationType_stagedOrder, stage.CodeSurfacePreparationType_orderStaged, &stage.CodeSurfacePreparationTypeOrder, stage.CodeSurfacePreparationTypes_mapString, codesurfacepreparationtype, codesurfacepreparationtype.Name)
	return codesurfacepreparationtype
}

// StagePreserveOrder puts codesurfacepreparationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeSurfacePreparationTypeOrder
// - update stage.CodeSurfacePreparationTypeOrder accordingly
func (codesurfacepreparationtype *CodeSurfacePreparationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationType_stagedOrder, stage.CodeSurfacePreparationType_orderStaged, &stage.CodeSurfacePreparationTypeOrder, stage.CodeSurfacePreparationTypes_mapString, codesurfacepreparationtype, order, codesurfacepreparationtype.Name)
}

// Unstage removes codesurfacepreparationtype off the model stage
func (codesurfacepreparationtype *CodeSurfacePreparationType) Unstage(stage *Stage) *CodeSurfacePreparationType {
	__gong__unstage(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationTypes_mapString, codesurfacepreparationtype, codesurfacepreparationtype.Name)
	return codesurfacepreparationtype
}

// UnstageVoid removes codesurfacepreparationtype off the model stage
func (codesurfacepreparationtype *CodeSurfacePreparationType) UnstageVoid(stage *Stage) {
	codesurfacepreparationtype.Unstage(stage)
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) StageVoid(stage *Stage) {
	codesurfacepreparationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codesurfacepreparationtype *CodeSurfacePreparationType) GetName() (res string) {
	return codesurfacepreparationtype.Name
}

// for satisfaction of GongStruct interface
func (codesurfacepreparationtype *CodeSurfacePreparationType) SetName(name string) {
	codesurfacepreparationtype.Name = name
}

// Stage puts codetlofsectiontype to the model stage
func (codetlofsectiontype *CodeTLOFSectionType) Stage(stage *Stage) *CodeTLOFSectionType {
	__gong__stage(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionType_stagedOrder, stage.CodeTLOFSectionType_orderStaged, &stage.CodeTLOFSectionTypeOrder, stage.CodeTLOFSectionTypes_mapString, codetlofsectiontype, codetlofsectiontype.Name)
	return codetlofsectiontype
}

// StagePreserveOrder puts codetlofsectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTLOFSectionTypeOrder
// - update stage.CodeTLOFSectionTypeOrder accordingly
func (codetlofsectiontype *CodeTLOFSectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionType_stagedOrder, stage.CodeTLOFSectionType_orderStaged, &stage.CodeTLOFSectionTypeOrder, stage.CodeTLOFSectionTypes_mapString, codetlofsectiontype, order, codetlofsectiontype.Name)
}

// Unstage removes codetlofsectiontype off the model stage
func (codetlofsectiontype *CodeTLOFSectionType) Unstage(stage *Stage) *CodeTLOFSectionType {
	__gong__unstage(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionTypes_mapString, codetlofsectiontype, codetlofsectiontype.Name)
	return codetlofsectiontype
}

// UnstageVoid removes codetlofsectiontype off the model stage
func (codetlofsectiontype *CodeTLOFSectionType) UnstageVoid(stage *Stage) {
	codetlofsectiontype.Unstage(stage)
}

func (codetlofsectiontype *CodeTLOFSectionType) StageVoid(stage *Stage) {
	codetlofsectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetlofsectiontype *CodeTLOFSectionType) GetName() (res string) {
	return codetlofsectiontype.Name
}

// for satisfaction of GongStruct interface
func (codetlofsectiontype *CodeTLOFSectionType) SetName(name string) {
	codetlofsectiontype.Name = name
}

// Stage puts codetaxiwayelementtype to the model stage
func (codetaxiwayelementtype *CodeTaxiwayElementType) Stage(stage *Stage) *CodeTaxiwayElementType {
	__gong__stage(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementType_stagedOrder, stage.CodeTaxiwayElementType_orderStaged, &stage.CodeTaxiwayElementTypeOrder, stage.CodeTaxiwayElementTypes_mapString, codetaxiwayelementtype, codetaxiwayelementtype.Name)
	return codetaxiwayelementtype
}

// StagePreserveOrder puts codetaxiwayelementtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTaxiwayElementTypeOrder
// - update stage.CodeTaxiwayElementTypeOrder accordingly
func (codetaxiwayelementtype *CodeTaxiwayElementType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementType_stagedOrder, stage.CodeTaxiwayElementType_orderStaged, &stage.CodeTaxiwayElementTypeOrder, stage.CodeTaxiwayElementTypes_mapString, codetaxiwayelementtype, order, codetaxiwayelementtype.Name)
}

// Unstage removes codetaxiwayelementtype off the model stage
func (codetaxiwayelementtype *CodeTaxiwayElementType) Unstage(stage *Stage) *CodeTaxiwayElementType {
	__gong__unstage(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementTypes_mapString, codetaxiwayelementtype, codetaxiwayelementtype.Name)
	return codetaxiwayelementtype
}

// UnstageVoid removes codetaxiwayelementtype off the model stage
func (codetaxiwayelementtype *CodeTaxiwayElementType) UnstageVoid(stage *Stage) {
	codetaxiwayelementtype.Unstage(stage)
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) StageVoid(stage *Stage) {
	codetaxiwayelementtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetaxiwayelementtype *CodeTaxiwayElementType) GetName() (res string) {
	return codetaxiwayelementtype.Name
}

// for satisfaction of GongStruct interface
func (codetaxiwayelementtype *CodeTaxiwayElementType) SetName(name string) {
	codetaxiwayelementtype.Name = name
}

// Stage puts codetaxiwaysectiontype to the model stage
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) Stage(stage *Stage) *CodeTaxiwaySectionType {
	__gong__stage(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionType_stagedOrder, stage.CodeTaxiwaySectionType_orderStaged, &stage.CodeTaxiwaySectionTypeOrder, stage.CodeTaxiwaySectionTypes_mapString, codetaxiwaysectiontype, codetaxiwaysectiontype.Name)
	return codetaxiwaysectiontype
}

// StagePreserveOrder puts codetaxiwaysectiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTaxiwaySectionTypeOrder
// - update stage.CodeTaxiwaySectionTypeOrder accordingly
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionType_stagedOrder, stage.CodeTaxiwaySectionType_orderStaged, &stage.CodeTaxiwaySectionTypeOrder, stage.CodeTaxiwaySectionTypes_mapString, codetaxiwaysectiontype, order, codetaxiwaysectiontype.Name)
}

// Unstage removes codetaxiwaysectiontype off the model stage
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) Unstage(stage *Stage) *CodeTaxiwaySectionType {
	__gong__unstage(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionTypes_mapString, codetaxiwaysectiontype, codetaxiwaysectiontype.Name)
	return codetaxiwaysectiontype
}

// UnstageVoid removes codetaxiwaysectiontype off the model stage
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) UnstageVoid(stage *Stage) {
	codetaxiwaysectiontype.Unstage(stage)
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) StageVoid(stage *Stage) {
	codetaxiwaysectiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GetName() (res string) {
	return codetaxiwaysectiontype.Name
}

// for satisfaction of GongStruct interface
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) SetName(name string) {
	codetaxiwaysectiontype.Name = name
}

// Stage puts codetaxiwaytype to the model stage
func (codetaxiwaytype *CodeTaxiwayType) Stage(stage *Stage) *CodeTaxiwayType {
	__gong__stage(stage.CodeTaxiwayTypes, stage.CodeTaxiwayType_stagedOrder, stage.CodeTaxiwayType_orderStaged, &stage.CodeTaxiwayTypeOrder, stage.CodeTaxiwayTypes_mapString, codetaxiwaytype, codetaxiwaytype.Name)
	return codetaxiwaytype
}

// StagePreserveOrder puts codetaxiwaytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeTaxiwayTypeOrder
// - update stage.CodeTaxiwayTypeOrder accordingly
func (codetaxiwaytype *CodeTaxiwayType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeTaxiwayTypes, stage.CodeTaxiwayType_stagedOrder, stage.CodeTaxiwayType_orderStaged, &stage.CodeTaxiwayTypeOrder, stage.CodeTaxiwayTypes_mapString, codetaxiwaytype, order, codetaxiwaytype.Name)
}

// Unstage removes codetaxiwaytype off the model stage
func (codetaxiwaytype *CodeTaxiwayType) Unstage(stage *Stage) *CodeTaxiwayType {
	__gong__unstage(stage.CodeTaxiwayTypes, stage.CodeTaxiwayTypes_mapString, codetaxiwaytype, codetaxiwaytype.Name)
	return codetaxiwaytype
}

// UnstageVoid removes codetaxiwaytype off the model stage
func (codetaxiwaytype *CodeTaxiwayType) UnstageVoid(stage *Stage) {
	codetaxiwaytype.Unstage(stage)
}

func (codetaxiwaytype *CodeTaxiwayType) StageVoid(stage *Stage) {
	codetaxiwaytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codetaxiwaytype *CodeTaxiwayType) GetName() (res string) {
	return codetaxiwaytype.Name
}

// for satisfaction of GongStruct interface
func (codetaxiwaytype *CodeTaxiwayType) SetName(name string) {
	codetaxiwaytype.Name = name
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

// Stage puts codeusagelimitationtype to the model stage
func (codeusagelimitationtype *CodeUsageLimitationType) Stage(stage *Stage) *CodeUsageLimitationType {
	__gong__stage(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationType_stagedOrder, stage.CodeUsageLimitationType_orderStaged, &stage.CodeUsageLimitationTypeOrder, stage.CodeUsageLimitationTypes_mapString, codeusagelimitationtype, codeusagelimitationtype.Name)
	return codeusagelimitationtype
}

// StagePreserveOrder puts codeusagelimitationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeUsageLimitationTypeOrder
// - update stage.CodeUsageLimitationTypeOrder accordingly
func (codeusagelimitationtype *CodeUsageLimitationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationType_stagedOrder, stage.CodeUsageLimitationType_orderStaged, &stage.CodeUsageLimitationTypeOrder, stage.CodeUsageLimitationTypes_mapString, codeusagelimitationtype, order, codeusagelimitationtype.Name)
}

// Unstage removes codeusagelimitationtype off the model stage
func (codeusagelimitationtype *CodeUsageLimitationType) Unstage(stage *Stage) *CodeUsageLimitationType {
	__gong__unstage(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationTypes_mapString, codeusagelimitationtype, codeusagelimitationtype.Name)
	return codeusagelimitationtype
}

// UnstageVoid removes codeusagelimitationtype off the model stage
func (codeusagelimitationtype *CodeUsageLimitationType) UnstageVoid(stage *Stage) {
	codeusagelimitationtype.Unstage(stage)
}

func (codeusagelimitationtype *CodeUsageLimitationType) StageVoid(stage *Stage) {
	codeusagelimitationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeusagelimitationtype *CodeUsageLimitationType) GetName() (res string) {
	return codeusagelimitationtype.Name
}

// for satisfaction of GongStruct interface
func (codeusagelimitationtype *CodeUsageLimitationType) SetName(name string) {
	codeusagelimitationtype.Name = name
}

// Stage puts codevasistype to the model stage
func (codevasistype *CodeVASISType) Stage(stage *Stage) *CodeVASISType {
	__gong__stage(stage.CodeVASISTypes, stage.CodeVASISType_stagedOrder, stage.CodeVASISType_orderStaged, &stage.CodeVASISTypeOrder, stage.CodeVASISTypes_mapString, codevasistype, codevasistype.Name)
	return codevasistype
}

// StagePreserveOrder puts codevasistype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVASISTypeOrder
// - update stage.CodeVASISTypeOrder accordingly
func (codevasistype *CodeVASISType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVASISTypes, stage.CodeVASISType_stagedOrder, stage.CodeVASISType_orderStaged, &stage.CodeVASISTypeOrder, stage.CodeVASISTypes_mapString, codevasistype, order, codevasistype.Name)
}

// Unstage removes codevasistype off the model stage
func (codevasistype *CodeVASISType) Unstage(stage *Stage) *CodeVASISType {
	__gong__unstage(stage.CodeVASISTypes, stage.CodeVASISTypes_mapString, codevasistype, codevasistype.Name)
	return codevasistype
}

// UnstageVoid removes codevasistype off the model stage
func (codevasistype *CodeVASISType) UnstageVoid(stage *Stage) {
	codevasistype.Unstage(stage)
}

func (codevasistype *CodeVASISType) StageVoid(stage *Stage) {
	codevasistype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codevasistype *CodeVASISType) GetName() (res string) {
	return codevasistype.Name
}

// for satisfaction of GongStruct interface
func (codevasistype *CodeVASISType) SetName(name string) {
	codevasistype.Name = name
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

// Stage puts codevisualdockingguidancetype to the model stage
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) Stage(stage *Stage) *CodeVisualDockingGuidanceType {
	__gong__stage(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceType_stagedOrder, stage.CodeVisualDockingGuidanceType_orderStaged, &stage.CodeVisualDockingGuidanceTypeOrder, stage.CodeVisualDockingGuidanceTypes_mapString, codevisualdockingguidancetype, codevisualdockingguidancetype.Name)
	return codevisualdockingguidancetype
}

// StagePreserveOrder puts codevisualdockingguidancetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVisualDockingGuidanceTypeOrder
// - update stage.CodeVisualDockingGuidanceTypeOrder accordingly
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceType_stagedOrder, stage.CodeVisualDockingGuidanceType_orderStaged, &stage.CodeVisualDockingGuidanceTypeOrder, stage.CodeVisualDockingGuidanceTypes_mapString, codevisualdockingguidancetype, order, codevisualdockingguidancetype.Name)
}

// Unstage removes codevisualdockingguidancetype off the model stage
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) Unstage(stage *Stage) *CodeVisualDockingGuidanceType {
	__gong__unstage(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceTypes_mapString, codevisualdockingguidancetype, codevisualdockingguidancetype.Name)
	return codevisualdockingguidancetype
}

// UnstageVoid removes codevisualdockingguidancetype off the model stage
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) UnstageVoid(stage *Stage) {
	codevisualdockingguidancetype.Unstage(stage)
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) StageVoid(stage *Stage) {
	codevisualdockingguidancetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GetName() (res string) {
	return codevisualdockingguidancetype.Name
}

// for satisfaction of GongStruct interface
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) SetName(name string) {
	codevisualdockingguidancetype.Name = name
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

// Stage puts codeworkareatype to the model stage
func (codeworkareatype *CodeWorkAreaType) Stage(stage *Stage) *CodeWorkAreaType {
	__gong__stage(stage.CodeWorkAreaTypes, stage.CodeWorkAreaType_stagedOrder, stage.CodeWorkAreaType_orderStaged, &stage.CodeWorkAreaTypeOrder, stage.CodeWorkAreaTypes_mapString, codeworkareatype, codeworkareatype.Name)
	return codeworkareatype
}

// StagePreserveOrder puts codeworkareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeWorkAreaTypeOrder
// - update stage.CodeWorkAreaTypeOrder accordingly
func (codeworkareatype *CodeWorkAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeWorkAreaTypes, stage.CodeWorkAreaType_stagedOrder, stage.CodeWorkAreaType_orderStaged, &stage.CodeWorkAreaTypeOrder, stage.CodeWorkAreaTypes_mapString, codeworkareatype, order, codeworkareatype.Name)
}

// Unstage removes codeworkareatype off the model stage
func (codeworkareatype *CodeWorkAreaType) Unstage(stage *Stage) *CodeWorkAreaType {
	__gong__unstage(stage.CodeWorkAreaTypes, stage.CodeWorkAreaTypes_mapString, codeworkareatype, codeworkareatype.Name)
	return codeworkareatype
}

// UnstageVoid removes codeworkareatype off the model stage
func (codeworkareatype *CodeWorkAreaType) UnstageVoid(stage *Stage) {
	codeworkareatype.Unstage(stage)
}

func (codeworkareatype *CodeWorkAreaType) StageVoid(stage *Stage) {
	codeworkareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeworkareatype *CodeWorkAreaType) GetName() (res string) {
	return codeworkareatype.Name
}

// for satisfaction of GongStruct interface
func (codeworkareatype *CodeWorkAreaType) SetName(name string) {
	codeworkareatype.Name = name
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

// Stage puts datetimetype to the model stage
func (datetimetype *DateTimeType) Stage(stage *Stage) *DateTimeType {
	__gong__stage(stage.DateTimeTypes, stage.DateTimeType_stagedOrder, stage.DateTimeType_orderStaged, &stage.DateTimeTypeOrder, stage.DateTimeTypes_mapString, datetimetype, datetimetype.Name)
	return datetimetype
}

// StagePreserveOrder puts datetimetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DateTimeTypeOrder
// - update stage.DateTimeTypeOrder accordingly
func (datetimetype *DateTimeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DateTimeTypes, stage.DateTimeType_stagedOrder, stage.DateTimeType_orderStaged, &stage.DateTimeTypeOrder, stage.DateTimeTypes_mapString, datetimetype, order, datetimetype.Name)
}

// Unstage removes datetimetype off the model stage
func (datetimetype *DateTimeType) Unstage(stage *Stage) *DateTimeType {
	__gong__unstage(stage.DateTimeTypes, stage.DateTimeTypes_mapString, datetimetype, datetimetype.Name)
	return datetimetype
}

// UnstageVoid removes datetimetype off the model stage
func (datetimetype *DateTimeType) UnstageVoid(stage *Stage) {
	datetimetype.Unstage(stage)
}

func (datetimetype *DateTimeType) StageVoid(stage *Stage) {
	datetimetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (datetimetype *DateTimeType) GetName() (res string) {
	return datetimetype.Name
}

// for satisfaction of GongStruct interface
func (datetimetype *DateTimeType) SetName(name string) {
	datetimetype.Name = name
}

// Stage puts datetype to the model stage
func (datetype *DateType) Stage(stage *Stage) *DateType {
	__gong__stage(stage.DateTypes, stage.DateType_stagedOrder, stage.DateType_orderStaged, &stage.DateTypeOrder, stage.DateTypes_mapString, datetype, datetype.Name)
	return datetype
}

// StagePreserveOrder puts datetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DateTypeOrder
// - update stage.DateTypeOrder accordingly
func (datetype *DateType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DateTypes, stage.DateType_stagedOrder, stage.DateType_orderStaged, &stage.DateTypeOrder, stage.DateTypes_mapString, datetype, order, datetype.Name)
}

// Unstage removes datetype off the model stage
func (datetype *DateType) Unstage(stage *Stage) *DateType {
	__gong__unstage(stage.DateTypes, stage.DateTypes_mapString, datetype, datetype.Name)
	return datetype
}

// UnstageVoid removes datetype off the model stage
func (datetype *DateType) UnstageVoid(stage *Stage) {
	datetype.Unstage(stage)
}

func (datetype *DateType) StageVoid(stage *Stage) {
	datetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (datetype *DateType) GetName() (res string) {
	return datetype.Name
}

// for satisfaction of GongStruct interface
func (datetype *DateType) SetName(name string) {
	datetype.Name = name
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

// Stage puts valdepthtype to the model stage
func (valdepthtype *ValDepthType) Stage(stage *Stage) *ValDepthType {
	__gong__stage(stage.ValDepthTypes, stage.ValDepthType_stagedOrder, stage.ValDepthType_orderStaged, &stage.ValDepthTypeOrder, stage.ValDepthTypes_mapString, valdepthtype, valdepthtype.Name)
	return valdepthtype
}

// StagePreserveOrder puts valdepthtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValDepthTypeOrder
// - update stage.ValDepthTypeOrder accordingly
func (valdepthtype *ValDepthType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValDepthTypes, stage.ValDepthType_stagedOrder, stage.ValDepthType_orderStaged, &stage.ValDepthTypeOrder, stage.ValDepthTypes_mapString, valdepthtype, order, valdepthtype.Name)
}

// Unstage removes valdepthtype off the model stage
func (valdepthtype *ValDepthType) Unstage(stage *Stage) *ValDepthType {
	__gong__unstage(stage.ValDepthTypes, stage.ValDepthTypes_mapString, valdepthtype, valdepthtype.Name)
	return valdepthtype
}

// UnstageVoid removes valdepthtype off the model stage
func (valdepthtype *ValDepthType) UnstageVoid(stage *Stage) {
	valdepthtype.Unstage(stage)
}

func (valdepthtype *ValDepthType) StageVoid(stage *Stage) {
	valdepthtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valdepthtype *ValDepthType) GetName() (res string) {
	return valdepthtype.Name
}

// for satisfaction of GongStruct interface
func (valdepthtype *ValDepthType) SetName(name string) {
	valdepthtype.Name = name
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

// Stage puts valfltype to the model stage
func (valfltype *ValFLType) Stage(stage *Stage) *ValFLType {
	__gong__stage(stage.ValFLTypes, stage.ValFLType_stagedOrder, stage.ValFLType_orderStaged, &stage.ValFLTypeOrder, stage.ValFLTypes_mapString, valfltype, valfltype.Name)
	return valfltype
}

// StagePreserveOrder puts valfltype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValFLTypeOrder
// - update stage.ValFLTypeOrder accordingly
func (valfltype *ValFLType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValFLTypes, stage.ValFLType_stagedOrder, stage.ValFLType_orderStaged, &stage.ValFLTypeOrder, stage.ValFLTypes_mapString, valfltype, order, valfltype.Name)
}

// Unstage removes valfltype off the model stage
func (valfltype *ValFLType) Unstage(stage *Stage) *ValFLType {
	__gong__unstage(stage.ValFLTypes, stage.ValFLTypes_mapString, valfltype, valfltype.Name)
	return valfltype
}

// UnstageVoid removes valfltype off the model stage
func (valfltype *ValFLType) UnstageVoid(stage *Stage) {
	valfltype.Unstage(stage)
}

func (valfltype *ValFLType) StageVoid(stage *Stage) {
	valfltype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valfltype *ValFLType) GetName() (res string) {
	return valfltype.Name
}

// for satisfaction of GongStruct interface
func (valfltype *ValFLType) SetName(name string) {
	valfltype.Name = name
}

// Stage puts valfrictiontype to the model stage
func (valfrictiontype *ValFrictionType) Stage(stage *Stage) *ValFrictionType {
	__gong__stage(stage.ValFrictionTypes, stage.ValFrictionType_stagedOrder, stage.ValFrictionType_orderStaged, &stage.ValFrictionTypeOrder, stage.ValFrictionTypes_mapString, valfrictiontype, valfrictiontype.Name)
	return valfrictiontype
}

// StagePreserveOrder puts valfrictiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValFrictionTypeOrder
// - update stage.ValFrictionTypeOrder accordingly
func (valfrictiontype *ValFrictionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValFrictionTypes, stage.ValFrictionType_stagedOrder, stage.ValFrictionType_orderStaged, &stage.ValFrictionTypeOrder, stage.ValFrictionTypes_mapString, valfrictiontype, order, valfrictiontype.Name)
}

// Unstage removes valfrictiontype off the model stage
func (valfrictiontype *ValFrictionType) Unstage(stage *Stage) *ValFrictionType {
	__gong__unstage(stage.ValFrictionTypes, stage.ValFrictionTypes_mapString, valfrictiontype, valfrictiontype.Name)
	return valfrictiontype
}

// UnstageVoid removes valfrictiontype off the model stage
func (valfrictiontype *ValFrictionType) UnstageVoid(stage *Stage) {
	valfrictiontype.Unstage(stage)
}

func (valfrictiontype *ValFrictionType) StageVoid(stage *Stage) {
	valfrictiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valfrictiontype *ValFrictionType) GetName() (res string) {
	return valfrictiontype.Name
}

// for satisfaction of GongStruct interface
func (valfrictiontype *ValFrictionType) SetName(name string) {
	valfrictiontype.Name = name
}

// Stage puts vallcntype to the model stage
func (vallcntype *ValLCNType) Stage(stage *Stage) *ValLCNType {
	__gong__stage(stage.ValLCNTypes, stage.ValLCNType_stagedOrder, stage.ValLCNType_orderStaged, &stage.ValLCNTypeOrder, stage.ValLCNTypes_mapString, vallcntype, vallcntype.Name)
	return vallcntype
}

// StagePreserveOrder puts vallcntype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValLCNTypeOrder
// - update stage.ValLCNTypeOrder accordingly
func (vallcntype *ValLCNType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValLCNTypes, stage.ValLCNType_stagedOrder, stage.ValLCNType_orderStaged, &stage.ValLCNTypeOrder, stage.ValLCNTypes_mapString, vallcntype, order, vallcntype.Name)
}

// Unstage removes vallcntype off the model stage
func (vallcntype *ValLCNType) Unstage(stage *Stage) *ValLCNType {
	__gong__unstage(stage.ValLCNTypes, stage.ValLCNTypes_mapString, vallcntype, vallcntype.Name)
	return vallcntype
}

// UnstageVoid removes vallcntype off the model stage
func (vallcntype *ValLCNType) UnstageVoid(stage *Stage) {
	vallcntype.Unstage(stage)
}

func (vallcntype *ValLCNType) StageVoid(stage *Stage) {
	vallcntype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (vallcntype *ValLCNType) GetName() (res string) {
	return vallcntype.Name
}

// for satisfaction of GongStruct interface
func (vallcntype *ValLCNType) SetName(name string) {
	vallcntype.Name = name
}

// Stage puts vallightintensitytype to the model stage
func (vallightintensitytype *ValLightIntensityType) Stage(stage *Stage) *ValLightIntensityType {
	__gong__stage(stage.ValLightIntensityTypes, stage.ValLightIntensityType_stagedOrder, stage.ValLightIntensityType_orderStaged, &stage.ValLightIntensityTypeOrder, stage.ValLightIntensityTypes_mapString, vallightintensitytype, vallightintensitytype.Name)
	return vallightintensitytype
}

// StagePreserveOrder puts vallightintensitytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValLightIntensityTypeOrder
// - update stage.ValLightIntensityTypeOrder accordingly
func (vallightintensitytype *ValLightIntensityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValLightIntensityTypes, stage.ValLightIntensityType_stagedOrder, stage.ValLightIntensityType_orderStaged, &stage.ValLightIntensityTypeOrder, stage.ValLightIntensityTypes_mapString, vallightintensitytype, order, vallightintensitytype.Name)
}

// Unstage removes vallightintensitytype off the model stage
func (vallightintensitytype *ValLightIntensityType) Unstage(stage *Stage) *ValLightIntensityType {
	__gong__unstage(stage.ValLightIntensityTypes, stage.ValLightIntensityTypes_mapString, vallightintensitytype, vallightintensitytype.Name)
	return vallightintensitytype
}

// UnstageVoid removes vallightintensitytype off the model stage
func (vallightintensitytype *ValLightIntensityType) UnstageVoid(stage *Stage) {
	vallightintensitytype.Unstage(stage)
}

func (vallightintensitytype *ValLightIntensityType) StageVoid(stage *Stage) {
	vallightintensitytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (vallightintensitytype *ValLightIntensityType) GetName() (res string) {
	return vallightintensitytype.Name
}

// for satisfaction of GongStruct interface
func (vallightintensitytype *ValLightIntensityType) SetName(name string) {
	vallightintensitytype.Name = name
}

// Stage puts valmagneticvariationchangetype to the model stage
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) Stage(stage *Stage) *ValMagneticVariationChangeType {
	__gong__stage(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeType_stagedOrder, stage.ValMagneticVariationChangeType_orderStaged, &stage.ValMagneticVariationChangeTypeOrder, stage.ValMagneticVariationChangeTypes_mapString, valmagneticvariationchangetype, valmagneticvariationchangetype.Name)
	return valmagneticvariationchangetype
}

// StagePreserveOrder puts valmagneticvariationchangetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValMagneticVariationChangeTypeOrder
// - update stage.ValMagneticVariationChangeTypeOrder accordingly
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeType_stagedOrder, stage.ValMagneticVariationChangeType_orderStaged, &stage.ValMagneticVariationChangeTypeOrder, stage.ValMagneticVariationChangeTypes_mapString, valmagneticvariationchangetype, order, valmagneticvariationchangetype.Name)
}

// Unstage removes valmagneticvariationchangetype off the model stage
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) Unstage(stage *Stage) *ValMagneticVariationChangeType {
	__gong__unstage(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeTypes_mapString, valmagneticvariationchangetype, valmagneticvariationchangetype.Name)
	return valmagneticvariationchangetype
}

// UnstageVoid removes valmagneticvariationchangetype off the model stage
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) UnstageVoid(stage *Stage) {
	valmagneticvariationchangetype.Unstage(stage)
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) StageVoid(stage *Stage) {
	valmagneticvariationchangetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GetName() (res string) {
	return valmagneticvariationchangetype.Name
}

// for satisfaction of GongStruct interface
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) SetName(name string) {
	valmagneticvariationchangetype.Name = name
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

// Stage puts valpcntype to the model stage
func (valpcntype *ValPCNType) Stage(stage *Stage) *ValPCNType {
	__gong__stage(stage.ValPCNTypes, stage.ValPCNType_stagedOrder, stage.ValPCNType_orderStaged, &stage.ValPCNTypeOrder, stage.ValPCNTypes_mapString, valpcntype, valpcntype.Name)
	return valpcntype
}

// StagePreserveOrder puts valpcntype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValPCNTypeOrder
// - update stage.ValPCNTypeOrder accordingly
func (valpcntype *ValPCNType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValPCNTypes, stage.ValPCNType_stagedOrder, stage.ValPCNType_orderStaged, &stage.ValPCNTypeOrder, stage.ValPCNTypes_mapString, valpcntype, order, valpcntype.Name)
}

// Unstage removes valpcntype off the model stage
func (valpcntype *ValPCNType) Unstage(stage *Stage) *ValPCNType {
	__gong__unstage(stage.ValPCNTypes, stage.ValPCNTypes_mapString, valpcntype, valpcntype.Name)
	return valpcntype
}

// UnstageVoid removes valpcntype off the model stage
func (valpcntype *ValPCNType) UnstageVoid(stage *Stage) {
	valpcntype.Unstage(stage)
}

func (valpcntype *ValPCNType) StageVoid(stage *Stage) {
	valpcntype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valpcntype *ValPCNType) GetName() (res string) {
	return valpcntype.Name
}

// for satisfaction of GongStruct interface
func (valpcntype *ValPCNType) SetName(name string) {
	valpcntype.Name = name
}

// Stage puts valpercenttype to the model stage
func (valpercenttype *ValPercentType) Stage(stage *Stage) *ValPercentType {
	__gong__stage(stage.ValPercentTypes, stage.ValPercentType_stagedOrder, stage.ValPercentType_orderStaged, &stage.ValPercentTypeOrder, stage.ValPercentTypes_mapString, valpercenttype, valpercenttype.Name)
	return valpercenttype
}

// StagePreserveOrder puts valpercenttype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValPercentTypeOrder
// - update stage.ValPercentTypeOrder accordingly
func (valpercenttype *ValPercentType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValPercentTypes, stage.ValPercentType_stagedOrder, stage.ValPercentType_orderStaged, &stage.ValPercentTypeOrder, stage.ValPercentTypes_mapString, valpercenttype, order, valpercenttype.Name)
}

// Unstage removes valpercenttype off the model stage
func (valpercenttype *ValPercentType) Unstage(stage *Stage) *ValPercentType {
	__gong__unstage(stage.ValPercentTypes, stage.ValPercentTypes_mapString, valpercenttype, valpercenttype.Name)
	return valpercenttype
}

// UnstageVoid removes valpercenttype off the model stage
func (valpercenttype *ValPercentType) UnstageVoid(stage *Stage) {
	valpercenttype.Unstage(stage)
}

func (valpercenttype *ValPercentType) StageVoid(stage *Stage) {
	valpercenttype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valpercenttype *ValPercentType) GetName() (res string) {
	return valpercenttype.Name
}

// for satisfaction of GongStruct interface
func (valpercenttype *ValPercentType) SetName(name string) {
	valpercenttype.Name = name
}

// Stage puts valpressuretype to the model stage
func (valpressuretype *ValPressureType) Stage(stage *Stage) *ValPressureType {
	__gong__stage(stage.ValPressureTypes, stage.ValPressureType_stagedOrder, stage.ValPressureType_orderStaged, &stage.ValPressureTypeOrder, stage.ValPressureTypes_mapString, valpressuretype, valpressuretype.Name)
	return valpressuretype
}

// StagePreserveOrder puts valpressuretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ValPressureTypeOrder
// - update stage.ValPressureTypeOrder accordingly
func (valpressuretype *ValPressureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ValPressureTypes, stage.ValPressureType_stagedOrder, stage.ValPressureType_orderStaged, &stage.ValPressureTypeOrder, stage.ValPressureTypes_mapString, valpressuretype, order, valpressuretype.Name)
}

// Unstage removes valpressuretype off the model stage
func (valpressuretype *ValPressureType) Unstage(stage *Stage) *ValPressureType {
	__gong__unstage(stage.ValPressureTypes, stage.ValPressureTypes_mapString, valpressuretype, valpressuretype.Name)
	return valpressuretype
}

// UnstageVoid removes valpressuretype off the model stage
func (valpressuretype *ValPressureType) UnstageVoid(stage *Stage) {
	valpressuretype.Unstage(stage)
}

func (valpressuretype *ValPressureType) StageVoid(stage *Stage) {
	valpressuretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (valpressuretype *ValPressureType) GetName() (res string) {
	return valpressuretype.Name
}

// for satisfaction of GongStruct interface
func (valpressuretype *ValPressureType) SetName(name string) {
	valpressuretype.Name = name
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
	__gong__resetStageType(&stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString, &stage.CodeAircraftCategoryType_stagedOrder, &stage.CodeAircraftCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString, &stage.CodeAircraftEngineNumberType_stagedOrder, &stage.CodeAircraftEngineNumberTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString, &stage.CodeAircraftEngineType_stagedOrder, &stage.CodeAircraftEngineTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString, &stage.CodeAircraftICAOType_stagedOrder, &stage.CodeAircraftICAOTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftStandTypes, &stage.CodeAircraftStandTypes_mapString, &stage.CodeAircraftStandType_stagedOrder, &stage.CodeAircraftStandTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString, &stage.CodeAircraftType_stagedOrder, &stage.CodeAircraftTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString, &stage.CodeAircraftWingspanClassType_stagedOrder, &stage.CodeAircraftWingspanClassTypeOrder)

	__gong__resetStageType(&stage.CodeAirportHeliportCollocationTypes, &stage.CodeAirportHeliportCollocationTypes_mapString, &stage.CodeAirportHeliportCollocationType_stagedOrder, &stage.CodeAirportHeliportCollocationTypeOrder)

	__gong__resetStageType(&stage.CodeAirportHeliportDesignatorTypes, &stage.CodeAirportHeliportDesignatorTypes_mapString, &stage.CodeAirportHeliportDesignatorType_stagedOrder, &stage.CodeAirportHeliportDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeAirportHeliportTypes, &stage.CodeAirportHeliportTypes_mapString, &stage.CodeAirportHeliportType_stagedOrder, &stage.CodeAirportHeliportTypeOrder)

	__gong__resetStageType(&stage.CodeAirportWarningTypes, &stage.CodeAirportWarningTypes_mapString, &stage.CodeAirportWarningType_stagedOrder, &stage.CodeAirportWarningTypeOrder)

	__gong__resetStageType(&stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_mapString, &stage.CodeApproachGuidanceType_stagedOrder, &stage.CodeApproachGuidanceTypeOrder)

	__gong__resetStageType(&stage.CodeApproachLightingICAOTypes, &stage.CodeApproachLightingICAOTypes_mapString, &stage.CodeApproachLightingICAOType_stagedOrder, &stage.CodeApproachLightingICAOTypeOrder)

	__gong__resetStageType(&stage.CodeApproachLightingTypes, &stage.CodeApproachLightingTypes_mapString, &stage.CodeApproachLightingType_stagedOrder, &stage.CodeApproachLightingTypeOrder)

	__gong__resetStageType(&stage.CodeApronElementTypes, &stage.CodeApronElementTypes_mapString, &stage.CodeApronElementType_stagedOrder, &stage.CodeApronElementTypeOrder)

	__gong__resetStageType(&stage.CodeApronSectionTypes, &stage.CodeApronSectionTypes_mapString, &stage.CodeApronSectionType_stagedOrder, &stage.CodeApronSectionTypeOrder)

	__gong__resetStageType(&stage.CodeArrestingGearEnergyAbsorbTypes, &stage.CodeArrestingGearEnergyAbsorbTypes_mapString, &stage.CodeArrestingGearEnergyAbsorbType_stagedOrder, &stage.CodeArrestingGearEnergyAbsorbTypeOrder)

	__gong__resetStageType(&stage.CodeArrestingGearEngageDeviceTypes, &stage.CodeArrestingGearEngageDeviceTypes_mapString, &stage.CodeArrestingGearEngageDeviceType_stagedOrder, &stage.CodeArrestingGearEngageDeviceTypeOrder)

	__gong__resetStageType(&stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_mapString, &stage.CodeAuthorityRoleType_stagedOrder, &stage.CodeAuthorityRoleTypeOrder)

	__gong__resetStageType(&stage.CodeBuoyDesignatorTypes, &stage.CodeBuoyDesignatorTypes_mapString, &stage.CodeBuoyDesignatorType_stagedOrder, &stage.CodeBuoyDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeBuoyTypes, &stage.CodeBuoyTypes_mapString, &stage.CodeBuoyType_stagedOrder, &stage.CodeBuoyTypeOrder)

	__gong__resetStageType(&stage.CodeColourTypes, &stage.CodeColourTypes_mapString, &stage.CodeColourType_stagedOrder, &stage.CodeColourTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString, &stage.CodeCommunicationModeType_stagedOrder, &stage.CodeCommunicationModeTypeOrder)

	__gong__resetStageType(&stage.CodeContaminationTypes, &stage.CodeContaminationTypes_mapString, &stage.CodeContaminationType_stagedOrder, &stage.CodeContaminationTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeDeclaredDistanceTypes, &stage.CodeDeclaredDistanceTypes_mapString, &stage.CodeDeclaredDistanceType_stagedOrder, &stage.CodeDeclaredDistanceTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString, &stage.CodeDirectionTurnType_stagedOrder, &stage.CodeDirectionTurnTypeOrder)

	__gong__resetStageType(&stage.CodeDirectionTypes, &stage.CodeDirectionTypes_mapString, &stage.CodeDirectionType_stagedOrder, &stage.CodeDirectionTypeOrder)

	__gong__resetStageType(&stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString, &stage.CodeEquipmentAntiCollisionType_stagedOrder, &stage.CodeEquipmentAntiCollisionTypeOrder)

	__gong__resetStageType(&stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_mapString, &stage.CodeFlightOriginType_stagedOrder, &stage.CodeFlightOriginTypeOrder)

	__gong__resetStageType(&stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_mapString, &stage.CodeFlightPurposeType_stagedOrder, &stage.CodeFlightPurposeTypeOrder)

	__gong__resetStageType(&stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString, &stage.CodeFlightRuleType_stagedOrder, &stage.CodeFlightRuleTypeOrder)

	__gong__resetStageType(&stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_mapString, &stage.CodeFlightStatusType_stagedOrder, &stage.CodeFlightStatusTypeOrder)

	__gong__resetStageType(&stage.CodeFlightTypes, &stage.CodeFlightTypes_mapString, &stage.CodeFlightType_stagedOrder, &stage.CodeFlightTypeOrder)

	__gong__resetStageType(&stage.CodeFrictionDeviceTypes, &stage.CodeFrictionDeviceTypes_mapString, &stage.CodeFrictionDeviceType_stagedOrder, &stage.CodeFrictionDeviceTypeOrder)

	__gong__resetStageType(&stage.CodeFrictionEstimateTypes, &stage.CodeFrictionEstimateTypes_mapString, &stage.CodeFrictionEstimateType_stagedOrder, &stage.CodeFrictionEstimateTypeOrder)

	__gong__resetStageType(&stage.CodeGradeSeparationTypes, &stage.CodeGradeSeparationTypes_mapString, &stage.CodeGradeSeparationType_stagedOrder, &stage.CodeGradeSeparationTypeOrder)

	__gong__resetStageType(&stage.CodeGuidanceLineTypes, &stage.CodeGuidanceLineTypes_mapString, &stage.CodeGuidanceLineType_stagedOrder, &stage.CodeGuidanceLineTypeOrder)

	__gong__resetStageType(&stage.CodeHelicopterPerformanceTypes, &stage.CodeHelicopterPerformanceTypes_mapString, &stage.CodeHelicopterPerformanceType_stagedOrder, &stage.CodeHelicopterPerformanceTypeOrder)

	__gong__resetStageType(&stage.CodeHoldingCategoryTypes, &stage.CodeHoldingCategoryTypes_mapString, &stage.CodeHoldingCategoryType_stagedOrder, &stage.CodeHoldingCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeIATATypes, &stage.CodeIATATypes_mapString, &stage.CodeIATAType_stagedOrder, &stage.CodeIATATypeOrder)

	__gong__resetStageType(&stage.CodeICAOTypes, &stage.CodeICAOTypes_mapString, &stage.CodeICAOType_stagedOrder, &stage.CodeICAOTypeOrder)

	__gong__resetStageType(&stage.CodeLightHoldingPositionTypes, &stage.CodeLightHoldingPositionTypes_mapString, &stage.CodeLightHoldingPositionType_stagedOrder, &stage.CodeLightHoldingPositionTypeOrder)

	__gong__resetStageType(&stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString, &stage.CodeLightIntensityType_stagedOrder, &stage.CodeLightIntensityTypeOrder)

	__gong__resetStageType(&stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_mapString, &stage.CodeLightSourceType_stagedOrder, &stage.CodeLightSourceTypeOrder)

	__gong__resetStageType(&stage.CodeLightingJARTypes, &stage.CodeLightingJARTypes_mapString, &stage.CodeLightingJARType_stagedOrder, &stage.CodeLightingJARTypeOrder)

	__gong__resetStageType(&stage.CodeLoadingBridgeTypes, &stage.CodeLoadingBridgeTypes_mapString, &stage.CodeLoadingBridgeType_stagedOrder, &stage.CodeLoadingBridgeTypeOrder)

	__gong__resetStageType(&stage.CodeLogicalOperatorTypes, &stage.CodeLogicalOperatorTypes_mapString, &stage.CodeLogicalOperatorType_stagedOrder, &stage.CodeLogicalOperatorTypeOrder)

	__gong__resetStageType(&stage.CodeMarkingConditionTypes, &stage.CodeMarkingConditionTypes_mapString, &stage.CodeMarkingConditionType_stagedOrder, &stage.CodeMarkingConditionTypeOrder)

	__gong__resetStageType(&stage.CodeMarkingStyleTypes, &stage.CodeMarkingStyleTypes_mapString, &stage.CodeMarkingStyleType_stagedOrder, &stage.CodeMarkingStyleTypeOrder)

	__gong__resetStageType(&stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_mapString, &stage.CodeMeteoConditionsType_stagedOrder, &stage.CodeMeteoConditionsTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString, &stage.CodeMilitaryOperationsType_stagedOrder, &stage.CodeMilitaryOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_mapString, &stage.CodeMilitaryStatusType_stagedOrder, &stage.CodeMilitaryStatusTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString, &stage.CodeNavigationEquipmentType_stagedOrder, &stage.CodeNavigationEquipmentTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString, &stage.CodeNavigationSpecificationType_stagedOrder, &stage.CodeNavigationSpecificationTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeOperationAirportHeliportTypes, &stage.CodeOperationAirportHeliportTypes_mapString, &stage.CodeOperationAirportHeliportType_stagedOrder, &stage.CodeOperationAirportHeliportTypeOrder)

	__gong__resetStageType(&stage.CodeOperationManoeuvringAreaTypes, &stage.CodeOperationManoeuvringAreaTypes_mapString, &stage.CodeOperationManoeuvringAreaType_stagedOrder, &stage.CodeOperationManoeuvringAreaTypeOrder)

	__gong__resetStageType(&stage.CodePCNMethodTypes, &stage.CodePCNMethodTypes_mapString, &stage.CodePCNMethodType_stagedOrder, &stage.CodePCNMethodTypeOrder)

	__gong__resetStageType(&stage.CodePCNPavementTypes, &stage.CodePCNPavementTypes_mapString, &stage.CodePCNPavementType_stagedOrder, &stage.CodePCNPavementTypeOrder)

	__gong__resetStageType(&stage.CodePCNSubgradeTypes, &stage.CodePCNSubgradeTypes_mapString, &stage.CodePCNSubgradeType_stagedOrder, &stage.CodePCNSubgradeTypeOrder)

	__gong__resetStageType(&stage.CodePCNTyrePressureTypes, &stage.CodePCNTyrePressureTypes_mapString, &stage.CodePCNTyrePressureType_stagedOrder, &stage.CodePCNTyrePressureTypeOrder)

	__gong__resetStageType(&stage.CodeProtectAreaSectionTypes, &stage.CodeProtectAreaSectionTypes_mapString, &stage.CodeProtectAreaSectionType_stagedOrder, &stage.CodeProtectAreaSectionTypeOrder)

	__gong__resetStageType(&stage.CodeRVRReadingTypes, &stage.CodeRVRReadingTypes_mapString, &stage.CodeRVRReadingType_stagedOrder, &stage.CodeRVRReadingTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString, &stage.CodeRVSMType_stagedOrder, &stage.CodeRVSMTypeOrder)

	__gong__resetStageType(&stage.CodeRoadTypes, &stage.CodeRoadTypes_mapString, &stage.CodeRoadType_stagedOrder, &stage.CodeRoadTypeOrder)

	__gong__resetStageType(&stage.CodeRunwayElementTypes, &stage.CodeRunwayElementTypes_mapString, &stage.CodeRunwayElementType_stagedOrder, &stage.CodeRunwayElementTypeOrder)

	__gong__resetStageType(&stage.CodeRunwayMarkingTypes, &stage.CodeRunwayMarkingTypes_mapString, &stage.CodeRunwayMarkingType_stagedOrder, &stage.CodeRunwayMarkingTypeOrder)

	__gong__resetStageType(&stage.CodeRunwayPointRoleTypes, &stage.CodeRunwayPointRoleTypes_mapString, &stage.CodeRunwayPointRoleType_stagedOrder, &stage.CodeRunwayPointRoleTypeOrder)

	__gong__resetStageType(&stage.CodeRunwayProtectionAreaTypes, &stage.CodeRunwayProtectionAreaTypes_mapString, &stage.CodeRunwayProtectionAreaType_stagedOrder, &stage.CodeRunwayProtectionAreaTypeOrder)

	__gong__resetStageType(&stage.CodeRunwaySectionTypes, &stage.CodeRunwaySectionTypes_mapString, &stage.CodeRunwaySectionType_stagedOrder, &stage.CodeRunwaySectionTypeOrder)

	__gong__resetStageType(&stage.CodeRunwayTypes, &stage.CodeRunwayTypes_mapString, &stage.CodeRunwayType_stagedOrder, &stage.CodeRunwayTypeOrder)

	__gong__resetStageType(&stage.CodeSideTypes, &stage.CodeSideTypes_mapString, &stage.CodeSideType_stagedOrder, &stage.CodeSideTypeOrder)

	__gong__resetStageType(&stage.CodeStatusAirportTypes, &stage.CodeStatusAirportTypes_mapString, &stage.CodeStatusAirportType_stagedOrder, &stage.CodeStatusAirportTypeOrder)

	__gong__resetStageType(&stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString, &stage.CodeStatusOperationsType_stagedOrder, &stage.CodeStatusOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeSurfaceCompositionTypes, &stage.CodeSurfaceCompositionTypes_mapString, &stage.CodeSurfaceCompositionType_stagedOrder, &stage.CodeSurfaceCompositionTypeOrder)

	__gong__resetStageType(&stage.CodeSurfaceConditionTypes, &stage.CodeSurfaceConditionTypes_mapString, &stage.CodeSurfaceConditionType_stagedOrder, &stage.CodeSurfaceConditionTypeOrder)

	__gong__resetStageType(&stage.CodeSurfacePreparationTypes, &stage.CodeSurfacePreparationTypes_mapString, &stage.CodeSurfacePreparationType_stagedOrder, &stage.CodeSurfacePreparationTypeOrder)

	__gong__resetStageType(&stage.CodeTLOFSectionTypes, &stage.CodeTLOFSectionTypes_mapString, &stage.CodeTLOFSectionType_stagedOrder, &stage.CodeTLOFSectionTypeOrder)

	__gong__resetStageType(&stage.CodeTaxiwayElementTypes, &stage.CodeTaxiwayElementTypes_mapString, &stage.CodeTaxiwayElementType_stagedOrder, &stage.CodeTaxiwayElementTypeOrder)

	__gong__resetStageType(&stage.CodeTaxiwaySectionTypes, &stage.CodeTaxiwaySectionTypes_mapString, &stage.CodeTaxiwaySectionType_stagedOrder, &stage.CodeTaxiwaySectionTypeOrder)

	__gong__resetStageType(&stage.CodeTaxiwayTypes, &stage.CodeTaxiwayTypes_mapString, &stage.CodeTaxiwayType_stagedOrder, &stage.CodeTaxiwayTypeOrder)

	__gong__resetStageType(&stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString, &stage.CodeTelecomNetworkType_stagedOrder, &stage.CodeTelecomNetworkTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString, &stage.CodeTransponderType_stagedOrder, &stage.CodeTransponderTypeOrder)

	__gong__resetStageType(&stage.CodeUsageLimitationTypes, &stage.CodeUsageLimitationTypes_mapString, &stage.CodeUsageLimitationType_stagedOrder, &stage.CodeUsageLimitationTypeOrder)

	__gong__resetStageType(&stage.CodeVASISTypes, &stage.CodeVASISTypes_mapString, &stage.CodeVASISType_stagedOrder, &stage.CodeVASISTypeOrder)

	__gong__resetStageType(&stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString, &stage.CodeValueInterpretationType_stagedOrder, &stage.CodeValueInterpretationTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeVisualDockingGuidanceTypes, &stage.CodeVisualDockingGuidanceTypes_mapString, &stage.CodeVisualDockingGuidanceType_stagedOrder, &stage.CodeVisualDockingGuidanceTypeOrder)

	__gong__resetStageType(&stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString, &stage.CodeWakeTurbulenceType_stagedOrder, &stage.CodeWakeTurbulenceTypeOrder)

	__gong__resetStageType(&stage.CodeWorkAreaTypes, &stage.CodeWorkAreaTypes_mapString, &stage.CodeWorkAreaType_stagedOrder, &stage.CodeWorkAreaTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.DateTimeTypes, &stage.DateTimeTypes_mapString, &stage.DateTimeType_stagedOrder, &stage.DateTimeTypeOrder)

	__gong__resetStageType(&stage.DateTypes, &stage.DateTypes_mapString, &stage.DateType_stagedOrder, &stage.DateTypeOrder)

	__gong__resetStageType(&stage.DateYearTypes, &stage.DateYearTypes_mapString, &stage.DateYearType_stagedOrder, &stage.DateYearTypeOrder)

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

	__gong__resetStageType(&stage.ValDepthTypes, &stage.ValDepthTypes_mapString, &stage.ValDepthType_stagedOrder, &stage.ValDepthTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValFLTypes, &stage.ValFLTypes_mapString, &stage.ValFLType_stagedOrder, &stage.ValFLTypeOrder)

	__gong__resetStageType(&stage.ValFrictionTypes, &stage.ValFrictionTypes_mapString, &stage.ValFrictionType_stagedOrder, &stage.ValFrictionTypeOrder)

	__gong__resetStageType(&stage.ValLCNTypes, &stage.ValLCNTypes_mapString, &stage.ValLCNType_stagedOrder, &stage.ValLCNTypeOrder)

	__gong__resetStageType(&stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_mapString, &stage.ValLightIntensityType_stagedOrder, &stage.ValLightIntensityTypeOrder)

	__gong__resetStageType(&stage.ValMagneticVariationChangeTypes, &stage.ValMagneticVariationChangeTypes_mapString, &stage.ValMagneticVariationChangeType_stagedOrder, &stage.ValMagneticVariationChangeTypeOrder)

	__gong__resetStageType(&stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_mapString, &stage.ValMagneticVariationType_stagedOrder, &stage.ValMagneticVariationTypeOrder)

	__gong__resetStageType(&stage.ValPCNTypes, &stage.ValPCNTypes_mapString, &stage.ValPCNType_stagedOrder, &stage.ValPCNTypeOrder)

	__gong__resetStageType(&stage.ValPercentTypes, &stage.ValPercentTypes_mapString, &stage.ValPercentType_stagedOrder, &stage.ValPercentTypeOrder)

	__gong__resetStageType(&stage.ValPressureTypes, &stage.ValPressureTypes_mapString, &stage.ValPressureType_stagedOrder, &stage.ValPressureTypeOrder)

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

func (CodeAircraftStandType) GongGetAssociationName() any {
	return &CodeAircraftStandType{
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

func (CodeAirportHeliportCollocationType) GongGetAssociationName() any {
	return &CodeAirportHeliportCollocationType{
	}
}

func (CodeAirportHeliportDesignatorType) GongGetAssociationName() any {
	return &CodeAirportHeliportDesignatorType{
	}
}

func (CodeAirportHeliportType) GongGetAssociationName() any {
	return &CodeAirportHeliportType{
	}
}

func (CodeAirportWarningType) GongGetAssociationName() any {
	return &CodeAirportWarningType{
	}
}

func (CodeApproachGuidanceType) GongGetAssociationName() any {
	return &CodeApproachGuidanceType{
	}
}

func (CodeApproachLightingICAOType) GongGetAssociationName() any {
	return &CodeApproachLightingICAOType{
	}
}

func (CodeApproachLightingType) GongGetAssociationName() any {
	return &CodeApproachLightingType{
	}
}

func (CodeApronElementType) GongGetAssociationName() any {
	return &CodeApronElementType{
	}
}

func (CodeApronSectionType) GongGetAssociationName() any {
	return &CodeApronSectionType{
	}
}

func (CodeArrestingGearEnergyAbsorbType) GongGetAssociationName() any {
	return &CodeArrestingGearEnergyAbsorbType{
	}
}

func (CodeArrestingGearEngageDeviceType) GongGetAssociationName() any {
	return &CodeArrestingGearEngageDeviceType{
	}
}

func (CodeAuthorityRoleType) GongGetAssociationName() any {
	return &CodeAuthorityRoleType{
	}
}

func (CodeBuoyDesignatorType) GongGetAssociationName() any {
	return &CodeBuoyDesignatorType{
	}
}

func (CodeBuoyType) GongGetAssociationName() any {
	return &CodeBuoyType{
	}
}

func (CodeColourType) GongGetAssociationName() any {
	return &CodeColourType{
	}
}

func (CodeCommunicationModeType) GongGetAssociationName() any {
	return &CodeCommunicationModeType{
	}
}

func (CodeContaminationType) GongGetAssociationName() any {
	return &CodeContaminationType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeDeclaredDistanceType) GongGetAssociationName() any {
	return &CodeDeclaredDistanceType{
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

func (CodeFrictionDeviceType) GongGetAssociationName() any {
	return &CodeFrictionDeviceType{
	}
}

func (CodeFrictionEstimateType) GongGetAssociationName() any {
	return &CodeFrictionEstimateType{
	}
}

func (CodeGradeSeparationType) GongGetAssociationName() any {
	return &CodeGradeSeparationType{
	}
}

func (CodeGuidanceLineType) GongGetAssociationName() any {
	return &CodeGuidanceLineType{
	}
}

func (CodeHelicopterPerformanceType) GongGetAssociationName() any {
	return &CodeHelicopterPerformanceType{
	}
}

func (CodeHoldingCategoryType) GongGetAssociationName() any {
	return &CodeHoldingCategoryType{
	}
}

func (CodeIATAType) GongGetAssociationName() any {
	return &CodeIATAType{
	}
}

func (CodeICAOType) GongGetAssociationName() any {
	return &CodeICAOType{
	}
}

func (CodeLightHoldingPositionType) GongGetAssociationName() any {
	return &CodeLightHoldingPositionType{
	}
}

func (CodeLightIntensityType) GongGetAssociationName() any {
	return &CodeLightIntensityType{
	}
}

func (CodeLightSourceType) GongGetAssociationName() any {
	return &CodeLightSourceType{
	}
}

func (CodeLightingJARType) GongGetAssociationName() any {
	return &CodeLightingJARType{
	}
}

func (CodeLoadingBridgeType) GongGetAssociationName() any {
	return &CodeLoadingBridgeType{
	}
}

func (CodeLogicalOperatorType) GongGetAssociationName() any {
	return &CodeLogicalOperatorType{
	}
}

func (CodeMarkingConditionType) GongGetAssociationName() any {
	return &CodeMarkingConditionType{
	}
}

func (CodeMarkingStyleType) GongGetAssociationName() any {
	return &CodeMarkingStyleType{
	}
}

func (CodeMeteoConditionsType) GongGetAssociationName() any {
	return &CodeMeteoConditionsType{
	}
}

func (CodeMilitaryOperationsType) GongGetAssociationName() any {
	return &CodeMilitaryOperationsType{
	}
}

func (CodeMilitaryStatusType) GongGetAssociationName() any {
	return &CodeMilitaryStatusType{
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

func (CodeOperationAirportHeliportType) GongGetAssociationName() any {
	return &CodeOperationAirportHeliportType{
	}
}

func (CodeOperationManoeuvringAreaType) GongGetAssociationName() any {
	return &CodeOperationManoeuvringAreaType{
	}
}

func (CodePCNMethodType) GongGetAssociationName() any {
	return &CodePCNMethodType{
	}
}

func (CodePCNPavementType) GongGetAssociationName() any {
	return &CodePCNPavementType{
	}
}

func (CodePCNSubgradeType) GongGetAssociationName() any {
	return &CodePCNSubgradeType{
	}
}

func (CodePCNTyrePressureType) GongGetAssociationName() any {
	return &CodePCNTyrePressureType{
	}
}

func (CodeProtectAreaSectionType) GongGetAssociationName() any {
	return &CodeProtectAreaSectionType{
	}
}

func (CodeRVRReadingType) GongGetAssociationName() any {
	return &CodeRVRReadingType{
	}
}

func (CodeRVSMType) GongGetAssociationName() any {
	return &CodeRVSMType{
	}
}

func (CodeRoadType) GongGetAssociationName() any {
	return &CodeRoadType{
	}
}

func (CodeRunwayElementType) GongGetAssociationName() any {
	return &CodeRunwayElementType{
	}
}

func (CodeRunwayMarkingType) GongGetAssociationName() any {
	return &CodeRunwayMarkingType{
	}
}

func (CodeRunwayPointRoleType) GongGetAssociationName() any {
	return &CodeRunwayPointRoleType{
	}
}

func (CodeRunwayProtectionAreaType) GongGetAssociationName() any {
	return &CodeRunwayProtectionAreaType{
	}
}

func (CodeRunwaySectionType) GongGetAssociationName() any {
	return &CodeRunwaySectionType{
	}
}

func (CodeRunwayType) GongGetAssociationName() any {
	return &CodeRunwayType{
	}
}

func (CodeSideType) GongGetAssociationName() any {
	return &CodeSideType{
	}
}

func (CodeStatusAirportType) GongGetAssociationName() any {
	return &CodeStatusAirportType{
	}
}

func (CodeStatusOperationsType) GongGetAssociationName() any {
	return &CodeStatusOperationsType{
	}
}

func (CodeSurfaceCompositionType) GongGetAssociationName() any {
	return &CodeSurfaceCompositionType{
	}
}

func (CodeSurfaceConditionType) GongGetAssociationName() any {
	return &CodeSurfaceConditionType{
	}
}

func (CodeSurfacePreparationType) GongGetAssociationName() any {
	return &CodeSurfacePreparationType{
	}
}

func (CodeTLOFSectionType) GongGetAssociationName() any {
	return &CodeTLOFSectionType{
	}
}

func (CodeTaxiwayElementType) GongGetAssociationName() any {
	return &CodeTaxiwayElementType{
	}
}

func (CodeTaxiwaySectionType) GongGetAssociationName() any {
	return &CodeTaxiwaySectionType{
	}
}

func (CodeTaxiwayType) GongGetAssociationName() any {
	return &CodeTaxiwayType{
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

func (CodeUsageLimitationType) GongGetAssociationName() any {
	return &CodeUsageLimitationType{
	}
}

func (CodeVASISType) GongGetAssociationName() any {
	return &CodeVASISType{
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

func (CodeVisualDockingGuidanceType) GongGetAssociationName() any {
	return &CodeVisualDockingGuidanceType{
	}
}

func (CodeWakeTurbulenceType) GongGetAssociationName() any {
	return &CodeWakeTurbulenceType{
	}
}

func (CodeWorkAreaType) GongGetAssociationName() any {
	return &CodeWorkAreaType{
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

func (DateTimeType) GongGetAssociationName() any {
	return &DateTimeType{
	}
}

func (DateType) GongGetAssociationName() any {
	return &DateType{
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

func (ValDepthType) GongGetAssociationName() any {
	return &ValDepthType{
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

func (ValFLType) GongGetAssociationName() any {
	return &ValFLType{
	}
}

func (ValFrictionType) GongGetAssociationName() any {
	return &ValFrictionType{
	}
}

func (ValLCNType) GongGetAssociationName() any {
	return &ValLCNType{
	}
}

func (ValLightIntensityType) GongGetAssociationName() any {
	return &ValLightIntensityType{
	}
}

func (ValMagneticVariationChangeType) GongGetAssociationName() any {
	return &ValMagneticVariationChangeType{
	}
}

func (ValMagneticVariationType) GongGetAssociationName() any {
	return &ValMagneticVariationType{
	}
}

func (ValPCNType) GongGetAssociationName() any {
	return &ValPCNType{
	}
}

func (ValPercentType) GongGetAssociationName() any {
	return &ValPercentType{
	}
}

func (ValPressureType) GongGetAssociationName() any {
	return &ValPressureType{
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
	// reverse maps of direct associations of CodeAircraftStandType
	case CodeAircraftStandType:
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
	// reverse maps of direct associations of CodeAirportHeliportCollocationType
	case CodeAirportHeliportCollocationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportHeliportDesignatorType
	case CodeAirportHeliportDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportHeliportType
	case CodeAirportHeliportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportWarningType
	case CodeAirportWarningType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachGuidanceType
	case CodeApproachGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachLightingICAOType
	case CodeApproachLightingICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachLightingType
	case CodeApproachLightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApronElementType
	case CodeApronElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApronSectionType
	case CodeApronSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArrestingGearEnergyAbsorbType
	case CodeArrestingGearEnergyAbsorbType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArrestingGearEngageDeviceType
	case CodeArrestingGearEngageDeviceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityRoleType
	case CodeAuthorityRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBuoyDesignatorType
	case CodeBuoyDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBuoyType
	case CodeBuoyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeColourType
	case CodeColourType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeContaminationType
	case CodeContaminationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDeclaredDistanceType
	case CodeDeclaredDistanceType:
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
	// reverse maps of direct associations of CodeFrictionDeviceType
	case CodeFrictionDeviceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFrictionEstimateType
	case CodeFrictionEstimateType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGradeSeparationType
	case CodeGradeSeparationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGuidanceLineType
	case CodeGuidanceLineType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHelicopterPerformanceType
	case CodeHelicopterPerformanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingCategoryType
	case CodeHoldingCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIATAType
	case CodeIATAType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeICAOType
	case CodeICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightHoldingPositionType
	case CodeLightHoldingPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightIntensityType
	case CodeLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightSourceType
	case CodeLightSourceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightingJARType
	case CodeLightingJARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLoadingBridgeType
	case CodeLoadingBridgeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLogicalOperatorType
	case CodeLogicalOperatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkingConditionType
	case CodeMarkingConditionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkingStyleType
	case CodeMarkingStyleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMeteoConditionsType
	case CodeMeteoConditionsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryStatusType
	case CodeMilitaryStatusType:
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
	// reverse maps of direct associations of CodeOperationAirportHeliportType
	case CodeOperationAirportHeliportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOperationManoeuvringAreaType
	case CodeOperationManoeuvringAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNMethodType
	case CodePCNMethodType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNPavementType
	case CodePCNPavementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNSubgradeType
	case CodePCNSubgradeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNTyrePressureType
	case CodePCNTyrePressureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProtectAreaSectionType
	case CodeProtectAreaSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVRReadingType
	case CodeRVRReadingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRoadType
	case CodeRoadType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayElementType
	case CodeRunwayElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayMarkingType
	case CodeRunwayMarkingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayPointRoleType
	case CodeRunwayPointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayProtectionAreaType
	case CodeRunwayProtectionAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwaySectionType
	case CodeRunwaySectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayType
	case CodeRunwayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSideType
	case CodeSideType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusAirportType
	case CodeStatusAirportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfaceCompositionType
	case CodeSurfaceCompositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfaceConditionType
	case CodeSurfaceConditionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfacePreparationType
	case CodeSurfacePreparationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTLOFSectionType
	case CodeTLOFSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwayElementType
	case CodeTaxiwayElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwaySectionType
	case CodeTaxiwaySectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwayType
	case CodeTaxiwayType:
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
	// reverse maps of direct associations of CodeUsageLimitationType
	case CodeUsageLimitationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVASISType
	case CodeVASISType:
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
	// reverse maps of direct associations of CodeVisualDockingGuidanceType
	case CodeVisualDockingGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWakeTurbulenceType
	case CodeWakeTurbulenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWorkAreaType
	case CodeWorkAreaType:
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
	// reverse maps of direct associations of DateTimeType
	case DateTimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DateType
	case DateType:
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
	// reverse maps of direct associations of ValDepthType
	case ValDepthType:
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
	// reverse maps of direct associations of ValFLType
	case ValFLType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValFrictionType
	case ValFrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValLCNType
	case ValLCNType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValLightIntensityType
	case ValLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationChangeType
	case ValMagneticVariationChangeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationType
	case ValMagneticVariationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPCNType
	case ValPCNType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPercentType
	case ValPercentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPressureType
	case ValPressureType:
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
	// reverse maps of direct associations of CodeAircraftStandType
	case CodeAircraftStandType:
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
	// reverse maps of direct associations of CodeAirportHeliportCollocationType
	case CodeAirportHeliportCollocationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportHeliportDesignatorType
	case CodeAirportHeliportDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportHeliportType
	case CodeAirportHeliportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirportWarningType
	case CodeAirportWarningType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachGuidanceType
	case CodeApproachGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachLightingICAOType
	case CodeApproachLightingICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApproachLightingType
	case CodeApproachLightingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApronElementType
	case CodeApronElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeApronSectionType
	case CodeApronSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArrestingGearEnergyAbsorbType
	case CodeArrestingGearEnergyAbsorbType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeArrestingGearEngageDeviceType
	case CodeArrestingGearEngageDeviceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAuthorityRoleType
	case CodeAuthorityRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBuoyDesignatorType
	case CodeBuoyDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeBuoyType
	case CodeBuoyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeColourType
	case CodeColourType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeCommunicationModeType
	case CodeCommunicationModeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeContaminationType
	case CodeContaminationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDeclaredDistanceType
	case CodeDeclaredDistanceType:
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
	// reverse maps of direct associations of CodeFrictionDeviceType
	case CodeFrictionDeviceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeFrictionEstimateType
	case CodeFrictionEstimateType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGradeSeparationType
	case CodeGradeSeparationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGuidanceLineType
	case CodeGuidanceLineType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHelicopterPerformanceType
	case CodeHelicopterPerformanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeHoldingCategoryType
	case CodeHoldingCategoryType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeIATAType
	case CodeIATAType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeICAOType
	case CodeICAOType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightHoldingPositionType
	case CodeLightHoldingPositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightIntensityType
	case CodeLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightSourceType
	case CodeLightSourceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLightingJARType
	case CodeLightingJARType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLoadingBridgeType
	case CodeLoadingBridgeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeLogicalOperatorType
	case CodeLogicalOperatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkingConditionType
	case CodeMarkingConditionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMarkingStyleType
	case CodeMarkingStyleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMeteoConditionsType
	case CodeMeteoConditionsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryStatusType
	case CodeMilitaryStatusType:
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
	// reverse maps of direct associations of CodeOperationAirportHeliportType
	case CodeOperationAirportHeliportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeOperationManoeuvringAreaType
	case CodeOperationManoeuvringAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNMethodType
	case CodePCNMethodType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNPavementType
	case CodePCNPavementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNSubgradeType
	case CodePCNSubgradeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodePCNTyrePressureType
	case CodePCNTyrePressureType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeProtectAreaSectionType
	case CodeProtectAreaSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVRReadingType
	case CodeRVRReadingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRoadType
	case CodeRoadType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayElementType
	case CodeRunwayElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayMarkingType
	case CodeRunwayMarkingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayPointRoleType
	case CodeRunwayPointRoleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayProtectionAreaType
	case CodeRunwayProtectionAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwaySectionType
	case CodeRunwaySectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeRunwayType
	case CodeRunwayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSideType
	case CodeSideType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusAirportType
	case CodeStatusAirportType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfaceCompositionType
	case CodeSurfaceCompositionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfaceConditionType
	case CodeSurfaceConditionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeSurfacePreparationType
	case CodeSurfacePreparationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTLOFSectionType
	case CodeTLOFSectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwayElementType
	case CodeTaxiwayElementType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwaySectionType
	case CodeTaxiwaySectionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeTaxiwayType
	case CodeTaxiwayType:
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
	// reverse maps of direct associations of CodeUsageLimitationType
	case CodeUsageLimitationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVASISType
	case CodeVASISType:
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
	// reverse maps of direct associations of CodeVisualDockingGuidanceType
	case CodeVisualDockingGuidanceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWakeTurbulenceType
	case CodeWakeTurbulenceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeWorkAreaType
	case CodeWorkAreaType:
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
	// reverse maps of direct associations of DateTimeType
	case DateTimeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DateType
	case DateType:
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
	// reverse maps of direct associations of ValDepthType
	case ValDepthType:
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
	// reverse maps of direct associations of ValFLType
	case ValFLType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValFrictionType
	case ValFrictionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValLCNType
	case ValLCNType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValLightIntensityType
	case ValLightIntensityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationChangeType
	case ValMagneticVariationChangeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValMagneticVariationType
	case ValMagneticVariationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPCNType
	case ValPCNType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPercentType
	case ValPercentType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ValPressureType
	case ValPressureType:
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

func (*CodeAircraftStandType) GongGetReverseFields() []GongReverseField {
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

func (*CodeAirportHeliportCollocationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirportHeliportDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirportHeliportType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirportWarningType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachGuidanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachLightingICAOType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApproachLightingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApronElementType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeApronSectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeArrestingGearEnergyAbsorbType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeArrestingGearEngageDeviceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAuthorityRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeBuoyDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeBuoyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeColourType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeCommunicationModeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeContaminationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDeclaredDistanceType) GongGetReverseFields() []GongReverseField {
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

func (*CodeFrictionDeviceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeFrictionEstimateType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeGradeSeparationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeGuidanceLineType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeHelicopterPerformanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeHoldingCategoryType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeIATAType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeICAOType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLightHoldingPositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLightIntensityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLightSourceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLightingJARType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLoadingBridgeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeLogicalOperatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMarkingConditionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMarkingStyleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMeteoConditionsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryOperationsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryStatusType) GongGetReverseFields() []GongReverseField {
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

func (*CodeOperationAirportHeliportType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeOperationManoeuvringAreaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePCNMethodType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePCNPavementType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePCNSubgradeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodePCNTyrePressureType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeProtectAreaSectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRVRReadingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRVSMType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRoadType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwayElementType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwayMarkingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwayPointRoleType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwayProtectionAreaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwaySectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeRunwayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSideType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusAirportType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusOperationsType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSurfaceCompositionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSurfaceConditionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeSurfacePreparationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTLOFSectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTaxiwayElementType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTaxiwaySectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeTaxiwayType) GongGetReverseFields() []GongReverseField {
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

func (*CodeUsageLimitationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVASISType) GongGetReverseFields() []GongReverseField {
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

func (*CodeVisualDockingGuidanceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeWakeTurbulenceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeWorkAreaType) GongGetReverseFields() []GongReverseField {
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

func (*DateTimeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DateType) GongGetReverseFields() []GongReverseField {
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

func (*ValDepthType) GongGetReverseFields() []GongReverseField {
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

func (*ValFLType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValFrictionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValLCNType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValLightIntensityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValMagneticVariationChangeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValMagneticVariationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValPCNType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValPercentType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*ValPressureType) GongGetReverseFields() []GongReverseField {
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

func (codeaircraftstandtype *CodeAircraftStandType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeapronelementtype *CodeApronElementType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeapronsectiontype *CodeApronSectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codebuoytype *CodeBuoyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codecolourtype *CodeColourType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codecontaminationtype *CodeContaminationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeiatatype *CodeIATAType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeicaotype *CodeICAOType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codelightsourcetype *CodeLightSourceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codelightingjartype *CodeLightingJARType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codepcnmethodtype *CodePCNMethodType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codepcnpavementtype *CodePCNPavementType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codervrreadingtype *CodeRVRReadingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (coderoadtype *CodeRoadType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (coderunwaytype *CodeRunwayType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codestatusairporttype *CodeStatusAirportType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codevasistype *CodeVASISType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeworkareatype *CodeWorkAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (datetimetype *DateTimeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (datetype *DateType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (valdepthtype *ValDepthType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valfltype *ValFLType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (valfrictiontype *ValFrictionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (vallcntype *ValLCNType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (vallightintensitytype *ValLightIntensityType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (valpcntype *ValPCNType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valpercenttype *ValPercentType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (valpressuretype *ValPressureType) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (codeaircraftstandtype *CodeAircraftStandType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeaircraftstandtype.Name
	case "NilReason":
		res.valueString = codeaircraftstandtype.NilReason
	case "EnclosedText":
		res.valueString = codeaircraftstandtype.EnclosedText
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

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairportheliportcollocationtype.Name
	case "NilReason":
		res.valueString = codeairportheliportcollocationtype.NilReason
	case "EnclosedText":
		res.valueString = codeairportheliportcollocationtype.EnclosedText
	}
	return
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairportheliportdesignatortype.Name
	case "NilReason":
		res.valueString = codeairportheliportdesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codeairportheliportdesignatortype.EnclosedText
	}
	return
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairportheliporttype.Name
	case "NilReason":
		res.valueString = codeairportheliporttype.NilReason
	case "EnclosedText":
		res.valueString = codeairportheliporttype.EnclosedText
	}
	return
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairportwarningtype.Name
	case "NilReason":
		res.valueString = codeairportwarningtype.NilReason
	case "EnclosedText":
		res.valueString = codeairportwarningtype.EnclosedText
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

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachlightingicaotype.Name
	case "NilReason":
		res.valueString = codeapproachlightingicaotype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachlightingicaotype.EnclosedText
	}
	return
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapproachlightingtype.Name
	case "NilReason":
		res.valueString = codeapproachlightingtype.NilReason
	case "EnclosedText":
		res.valueString = codeapproachlightingtype.EnclosedText
	}
	return
}

func (codeapronelementtype *CodeApronElementType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapronelementtype.Name
	case "NilReason":
		res.valueString = codeapronelementtype.NilReason
	case "EnclosedText":
		res.valueString = codeapronelementtype.EnclosedText
	}
	return
}

func (codeapronsectiontype *CodeApronSectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeapronsectiontype.Name
	case "NilReason":
		res.valueString = codeapronsectiontype.NilReason
	case "EnclosedText":
		res.valueString = codeapronsectiontype.EnclosedText
	}
	return
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codearrestinggearenergyabsorbtype.Name
	case "NilReason":
		res.valueString = codearrestinggearenergyabsorbtype.NilReason
	case "EnclosedText":
		res.valueString = codearrestinggearenergyabsorbtype.EnclosedText
	}
	return
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codearrestinggearengagedevicetype.Name
	case "NilReason":
		res.valueString = codearrestinggearengagedevicetype.NilReason
	case "EnclosedText":
		res.valueString = codearrestinggearengagedevicetype.EnclosedText
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

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codebuoydesignatortype.Name
	case "NilReason":
		res.valueString = codebuoydesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codebuoydesignatortype.EnclosedText
	}
	return
}

func (codebuoytype *CodeBuoyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codebuoytype.Name
	case "NilReason":
		res.valueString = codebuoytype.NilReason
	case "EnclosedText":
		res.valueString = codebuoytype.EnclosedText
	}
	return
}

func (codecolourtype *CodeColourType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecolourtype.Name
	case "NilReason":
		res.valueString = codecolourtype.NilReason
	case "EnclosedText":
		res.valueString = codecolourtype.EnclosedText
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

func (codecontaminationtype *CodeContaminationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codecontaminationtype.Name
	case "NilReason":
		res.valueString = codecontaminationtype.NilReason
	case "EnclosedText":
		res.valueString = codecontaminationtype.EnclosedText
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

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codedeclareddistancetype.Name
	case "NilReason":
		res.valueString = codedeclareddistancetype.NilReason
	case "EnclosedText":
		res.valueString = codedeclareddistancetype.EnclosedText
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefrictiondevicetype.Name
	case "NilReason":
		res.valueString = codefrictiondevicetype.NilReason
	case "EnclosedText":
		res.valueString = codefrictiondevicetype.EnclosedText
	}
	return
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codefrictionestimatetype.Name
	case "NilReason":
		res.valueString = codefrictionestimatetype.NilReason
	case "EnclosedText":
		res.valueString = codefrictionestimatetype.EnclosedText
	}
	return
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codegradeseparationtype.Name
	case "NilReason":
		res.valueString = codegradeseparationtype.NilReason
	case "EnclosedText":
		res.valueString = codegradeseparationtype.EnclosedText
	}
	return
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeguidancelinetype.Name
	case "NilReason":
		res.valueString = codeguidancelinetype.NilReason
	case "EnclosedText":
		res.valueString = codeguidancelinetype.EnclosedText
	}
	return
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codehelicopterperformancetype.Name
	case "NilReason":
		res.valueString = codehelicopterperformancetype.NilReason
	case "EnclosedText":
		res.valueString = codehelicopterperformancetype.EnclosedText
	}
	return
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeholdingcategorytype.Name
	case "NilReason":
		res.valueString = codeholdingcategorytype.NilReason
	case "EnclosedText":
		res.valueString = codeholdingcategorytype.EnclosedText
	}
	return
}

func (codeiatatype *CodeIATAType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeiatatype.Name
	case "NilReason":
		res.valueString = codeiatatype.NilReason
	case "EnclosedText":
		res.valueString = codeiatatype.EnclosedText
	}
	return
}

func (codeicaotype *CodeICAOType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeicaotype.Name
	case "NilReason":
		res.valueString = codeicaotype.NilReason
	case "EnclosedText":
		res.valueString = codeicaotype.EnclosedText
	}
	return
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelightholdingpositiontype.Name
	case "NilReason":
		res.valueString = codelightholdingpositiontype.NilReason
	case "EnclosedText":
		res.valueString = codelightholdingpositiontype.EnclosedText
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

func (codelightsourcetype *CodeLightSourceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelightsourcetype.Name
	case "NilReason":
		res.valueString = codelightsourcetype.NilReason
	case "EnclosedText":
		res.valueString = codelightsourcetype.EnclosedText
	}
	return
}

func (codelightingjartype *CodeLightingJARType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelightingjartype.Name
	case "NilReason":
		res.valueString = codelightingjartype.NilReason
	case "EnclosedText":
		res.valueString = codelightingjartype.EnclosedText
	}
	return
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeloadingbridgetype.Name
	case "NilReason":
		res.valueString = codeloadingbridgetype.NilReason
	case "EnclosedText":
		res.valueString = codeloadingbridgetype.EnclosedText
	}
	return
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codelogicaloperatortype.Name
	case "NilReason":
		res.valueString = codelogicaloperatortype.NilReason
	case "EnclosedText":
		res.valueString = codelogicaloperatortype.EnclosedText
	}
	return
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemarkingconditiontype.Name
	case "NilReason":
		res.valueString = codemarkingconditiontype.NilReason
	case "EnclosedText":
		res.valueString = codemarkingconditiontype.EnclosedText
	}
	return
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codemarkingstyletype.Name
	case "NilReason":
		res.valueString = codemarkingstyletype.NilReason
	case "EnclosedText":
		res.valueString = codemarkingstyletype.EnclosedText
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

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeoperationairportheliporttype.Name
	case "NilReason":
		res.valueString = codeoperationairportheliporttype.NilReason
	case "EnclosedText":
		res.valueString = codeoperationairportheliporttype.EnclosedText
	}
	return
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeoperationmanoeuvringareatype.Name
	case "NilReason":
		res.valueString = codeoperationmanoeuvringareatype.NilReason
	case "EnclosedText":
		res.valueString = codeoperationmanoeuvringareatype.EnclosedText
	}
	return
}

func (codepcnmethodtype *CodePCNMethodType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepcnmethodtype.Name
	case "NilReason":
		res.valueString = codepcnmethodtype.NilReason
	case "EnclosedText":
		res.valueString = codepcnmethodtype.EnclosedText
	}
	return
}

func (codepcnpavementtype *CodePCNPavementType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepcnpavementtype.Name
	case "NilReason":
		res.valueString = codepcnpavementtype.NilReason
	case "EnclosedText":
		res.valueString = codepcnpavementtype.EnclosedText
	}
	return
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepcnsubgradetype.Name
	case "NilReason":
		res.valueString = codepcnsubgradetype.NilReason
	case "EnclosedText":
		res.valueString = codepcnsubgradetype.EnclosedText
	}
	return
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codepcntyrepressuretype.Name
	case "NilReason":
		res.valueString = codepcntyrepressuretype.NilReason
	case "EnclosedText":
		res.valueString = codepcntyrepressuretype.EnclosedText
	}
	return
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeprotectareasectiontype.Name
	case "NilReason":
		res.valueString = codeprotectareasectiontype.NilReason
	case "EnclosedText":
		res.valueString = codeprotectareasectiontype.EnclosedText
	}
	return
}

func (codervrreadingtype *CodeRVRReadingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codervrreadingtype.Name
	case "NilReason":
		res.valueString = codervrreadingtype.NilReason
	case "EnclosedText":
		res.valueString = codervrreadingtype.EnclosedText
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

func (coderoadtype *CodeRoadType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderoadtype.Name
	case "NilReason":
		res.valueString = coderoadtype.NilReason
	case "EnclosedText":
		res.valueString = coderoadtype.EnclosedText
	}
	return
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwayelementtype.Name
	case "NilReason":
		res.valueString = coderunwayelementtype.NilReason
	case "EnclosedText":
		res.valueString = coderunwayelementtype.EnclosedText
	}
	return
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwaymarkingtype.Name
	case "NilReason":
		res.valueString = coderunwaymarkingtype.NilReason
	case "EnclosedText":
		res.valueString = coderunwaymarkingtype.EnclosedText
	}
	return
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwaypointroletype.Name
	case "NilReason":
		res.valueString = coderunwaypointroletype.NilReason
	case "EnclosedText":
		res.valueString = coderunwaypointroletype.EnclosedText
	}
	return
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwayprotectionareatype.Name
	case "NilReason":
		res.valueString = coderunwayprotectionareatype.NilReason
	case "EnclosedText":
		res.valueString = coderunwayprotectionareatype.EnclosedText
	}
	return
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwaysectiontype.Name
	case "NilReason":
		res.valueString = coderunwaysectiontype.NilReason
	case "EnclosedText":
		res.valueString = coderunwaysectiontype.EnclosedText
	}
	return
}

func (coderunwaytype *CodeRunwayType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coderunwaytype.Name
	case "NilReason":
		res.valueString = coderunwaytype.NilReason
	case "EnclosedText":
		res.valueString = coderunwaytype.EnclosedText
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

func (codestatusairporttype *CodeStatusAirportType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusairporttype.Name
	case "NilReason":
		res.valueString = codestatusairporttype.NilReason
	case "EnclosedText":
		res.valueString = codestatusairporttype.EnclosedText
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

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesurfacecompositiontype.Name
	case "NilReason":
		res.valueString = codesurfacecompositiontype.NilReason
	case "EnclosedText":
		res.valueString = codesurfacecompositiontype.EnclosedText
	}
	return
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesurfaceconditiontype.Name
	case "NilReason":
		res.valueString = codesurfaceconditiontype.NilReason
	case "EnclosedText":
		res.valueString = codesurfaceconditiontype.EnclosedText
	}
	return
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codesurfacepreparationtype.Name
	case "NilReason":
		res.valueString = codesurfacepreparationtype.NilReason
	case "EnclosedText":
		res.valueString = codesurfacepreparationtype.EnclosedText
	}
	return
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetlofsectiontype.Name
	case "NilReason":
		res.valueString = codetlofsectiontype.NilReason
	case "EnclosedText":
		res.valueString = codetlofsectiontype.EnclosedText
	}
	return
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetaxiwayelementtype.Name
	case "NilReason":
		res.valueString = codetaxiwayelementtype.NilReason
	case "EnclosedText":
		res.valueString = codetaxiwayelementtype.EnclosedText
	}
	return
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetaxiwaysectiontype.Name
	case "NilReason":
		res.valueString = codetaxiwaysectiontype.NilReason
	case "EnclosedText":
		res.valueString = codetaxiwaysectiontype.EnclosedText
	}
	return
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codetaxiwaytype.Name
	case "NilReason":
		res.valueString = codetaxiwaytype.NilReason
	case "EnclosedText":
		res.valueString = codetaxiwaytype.EnclosedText
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

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeusagelimitationtype.Name
	case "NilReason":
		res.valueString = codeusagelimitationtype.NilReason
	case "EnclosedText":
		res.valueString = codeusagelimitationtype.EnclosedText
	}
	return
}

func (codevasistype *CodeVASISType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codevasistype.Name
	case "NilReason":
		res.valueString = codevasistype.NilReason
	case "EnclosedText":
		res.valueString = codevasistype.EnclosedText
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

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codevisualdockingguidancetype.Name
	case "NilReason":
		res.valueString = codevisualdockingguidancetype.NilReason
	case "EnclosedText":
		res.valueString = codevisualdockingguidancetype.EnclosedText
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

func (codeworkareatype *CodeWorkAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeworkareatype.Name
	case "NilReason":
		res.valueString = codeworkareatype.NilReason
	case "EnclosedText":
		res.valueString = codeworkareatype.EnclosedText
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

func (datetimetype *DateTimeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = datetimetype.Name
	case "NilReason":
		res.valueString = datetimetype.NilReason
	case "EnclosedText":
		res.valueString = datetimetype.EnclosedText
	}
	return
}

func (datetype *DateType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = datetype.Name
	case "NilReason":
		res.valueString = datetype.NilReason
	case "EnclosedText":
		res.valueString = datetype.EnclosedText
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

func (valdepthtype *ValDepthType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valdepthtype.Name
	case "Uom":
		res.valueString = valdepthtype.Uom
	case "NilReason":
		res.valueString = valdepthtype.NilReason
	case "EnclosedText":
		res.valueString = valdepthtype.EnclosedText
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

func (valfltype *ValFLType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valfltype.Name
	case "Uom":
		res.valueString = valfltype.Uom
	case "NilReason":
		res.valueString = valfltype.NilReason
	case "EnclosedText":
		res.valueString = fmt.Sprintf("%d", valfltype.EnclosedText)
		res.valueInt = valfltype.EnclosedText
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (valfrictiontype *ValFrictionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valfrictiontype.Name
	case "NilReason":
		res.valueString = valfrictiontype.NilReason
	case "EnclosedText":
		res.valueString = valfrictiontype.EnclosedText
	}
	return
}

func (vallcntype *ValLCNType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = vallcntype.Name
	case "NilReason":
		res.valueString = vallcntype.NilReason
	case "EnclosedText":
		res.valueString = vallcntype.EnclosedText
	}
	return
}

func (vallightintensitytype *ValLightIntensityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = vallightintensitytype.Name
	case "Uom":
		res.valueString = vallightintensitytype.Uom
	case "NilReason":
		res.valueString = vallightintensitytype.NilReason
	case "EnclosedText":
		res.valueString = vallightintensitytype.EnclosedText
	}
	return
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valmagneticvariationchangetype.Name
	case "NilReason":
		res.valueString = valmagneticvariationchangetype.NilReason
	case "EnclosedText":
		res.valueString = valmagneticvariationchangetype.EnclosedText
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

func (valpcntype *ValPCNType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valpcntype.Name
	case "NilReason":
		res.valueString = valpcntype.NilReason
	case "EnclosedText":
		res.valueString = valpcntype.EnclosedText
	}
	return
}

func (valpercenttype *ValPercentType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valpercenttype.Name
	case "NilReason":
		res.valueString = valpercenttype.NilReason
	case "EnclosedText":
		res.valueString = valpercenttype.EnclosedText
	}
	return
}

func (valpressuretype *ValPressureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = valpressuretype.Name
	case "Uom":
		res.valueString = valpressuretype.Uom
	case "NilReason":
		res.valueString = valpressuretype.NilReason
	case "EnclosedText":
		res.valueString = valpressuretype.EnclosedText
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

func (codeaircraftstandtype *CodeAircraftStandType) GongGetGongstructName() string {
	return "CodeAircraftStandType"
}

func (codeaircrafttype *CodeAircraftType) GongGetGongstructName() string {
	return "CodeAircraftType"
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetGongstructName() string {
	return "CodeAircraftWingspanClassType"
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetGongstructName() string {
	return "CodeAirportHeliportCollocationType"
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetGongstructName() string {
	return "CodeAirportHeliportDesignatorType"
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetGongstructName() string {
	return "CodeAirportHeliportType"
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetGongstructName() string {
	return "CodeAirportWarningType"
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetGongstructName() string {
	return "CodeApproachGuidanceType"
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetGongstructName() string {
	return "CodeApproachLightingICAOType"
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetGongstructName() string {
	return "CodeApproachLightingType"
}

func (codeapronelementtype *CodeApronElementType) GongGetGongstructName() string {
	return "CodeApronElementType"
}

func (codeapronsectiontype *CodeApronSectionType) GongGetGongstructName() string {
	return "CodeApronSectionType"
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetGongstructName() string {
	return "CodeArrestingGearEnergyAbsorbType"
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetGongstructName() string {
	return "CodeArrestingGearEngageDeviceType"
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetGongstructName() string {
	return "CodeAuthorityRoleType"
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetGongstructName() string {
	return "CodeBuoyDesignatorType"
}

func (codebuoytype *CodeBuoyType) GongGetGongstructName() string {
	return "CodeBuoyType"
}

func (codecolourtype *CodeColourType) GongGetGongstructName() string {
	return "CodeColourType"
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetGongstructName() string {
	return "CodeCommunicationModeType"
}

func (codecontaminationtype *CodeContaminationType) GongGetGongstructName() string {
	return "CodeContaminationType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetGongstructName() string {
	return "CodeDeclaredDistanceType"
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

func (codeflightruletype *CodeFlightRuleType) GongGetGongstructName() string {
	return "CodeFlightRuleType"
}

func (codeflightstatustype *CodeFlightStatusType) GongGetGongstructName() string {
	return "CodeFlightStatusType"
}

func (codeflighttype *CodeFlightType) GongGetGongstructName() string {
	return "CodeFlightType"
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetGongstructName() string {
	return "CodeFrictionDeviceType"
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetGongstructName() string {
	return "CodeFrictionEstimateType"
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetGongstructName() string {
	return "CodeGradeSeparationType"
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetGongstructName() string {
	return "CodeGuidanceLineType"
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetGongstructName() string {
	return "CodeHelicopterPerformanceType"
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetGongstructName() string {
	return "CodeHoldingCategoryType"
}

func (codeiatatype *CodeIATAType) GongGetGongstructName() string {
	return "CodeIATAType"
}

func (codeicaotype *CodeICAOType) GongGetGongstructName() string {
	return "CodeICAOType"
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetGongstructName() string {
	return "CodeLightHoldingPositionType"
}

func (codelightintensitytype *CodeLightIntensityType) GongGetGongstructName() string {
	return "CodeLightIntensityType"
}

func (codelightsourcetype *CodeLightSourceType) GongGetGongstructName() string {
	return "CodeLightSourceType"
}

func (codelightingjartype *CodeLightingJARType) GongGetGongstructName() string {
	return "CodeLightingJARType"
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetGongstructName() string {
	return "CodeLoadingBridgeType"
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetGongstructName() string {
	return "CodeLogicalOperatorType"
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetGongstructName() string {
	return "CodeMarkingConditionType"
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetGongstructName() string {
	return "CodeMarkingStyleType"
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetGongstructName() string {
	return "CodeMeteoConditionsType"
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetGongstructName() string {
	return "CodeMilitaryOperationsType"
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetGongstructName() string {
	return "CodeMilitaryStatusType"
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

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetGongstructName() string {
	return "CodeOperationAirportHeliportType"
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetGongstructName() string {
	return "CodeOperationManoeuvringAreaType"
}

func (codepcnmethodtype *CodePCNMethodType) GongGetGongstructName() string {
	return "CodePCNMethodType"
}

func (codepcnpavementtype *CodePCNPavementType) GongGetGongstructName() string {
	return "CodePCNPavementType"
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetGongstructName() string {
	return "CodePCNSubgradeType"
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetGongstructName() string {
	return "CodePCNTyrePressureType"
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetGongstructName() string {
	return "CodeProtectAreaSectionType"
}

func (codervrreadingtype *CodeRVRReadingType) GongGetGongstructName() string {
	return "CodeRVRReadingType"
}

func (codervsmtype *CodeRVSMType) GongGetGongstructName() string {
	return "CodeRVSMType"
}

func (coderoadtype *CodeRoadType) GongGetGongstructName() string {
	return "CodeRoadType"
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetGongstructName() string {
	return "CodeRunwayElementType"
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetGongstructName() string {
	return "CodeRunwayMarkingType"
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetGongstructName() string {
	return "CodeRunwayPointRoleType"
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetGongstructName() string {
	return "CodeRunwayProtectionAreaType"
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetGongstructName() string {
	return "CodeRunwaySectionType"
}

func (coderunwaytype *CodeRunwayType) GongGetGongstructName() string {
	return "CodeRunwayType"
}

func (codesidetype *CodeSideType) GongGetGongstructName() string {
	return "CodeSideType"
}

func (codestatusairporttype *CodeStatusAirportType) GongGetGongstructName() string {
	return "CodeStatusAirportType"
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetGongstructName() string {
	return "CodeStatusOperationsType"
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetGongstructName() string {
	return "CodeSurfaceCompositionType"
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetGongstructName() string {
	return "CodeSurfaceConditionType"
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetGongstructName() string {
	return "CodeSurfacePreparationType"
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetGongstructName() string {
	return "CodeTLOFSectionType"
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetGongstructName() string {
	return "CodeTaxiwayElementType"
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetGongstructName() string {
	return "CodeTaxiwaySectionType"
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetGongstructName() string {
	return "CodeTaxiwayType"
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

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetGongstructName() string {
	return "CodeUsageLimitationType"
}

func (codevasistype *CodeVASISType) GongGetGongstructName() string {
	return "CodeVASISType"
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetGongstructName() string {
	return "CodeValueInterpretationType"
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetGongstructName() string {
	return "CodeVerticalDatumType"
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetGongstructName() string {
	return "CodeVisualDockingGuidanceType"
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetGongstructName() string {
	return "CodeWakeTurbulenceType"
}

func (codeworkareatype *CodeWorkAreaType) GongGetGongstructName() string {
	return "CodeWorkAreaType"
}

func (codeyesnotype *CodeYesNoType) GongGetGongstructName() string {
	return "CodeYesNoType"
}

func (datemonthdaytype *DateMonthDayType) GongGetGongstructName() string {
	return "DateMonthDayType"
}

func (datetimetype *DateTimeType) GongGetGongstructName() string {
	return "DateTimeType"
}

func (datetype *DateType) GongGetGongstructName() string {
	return "DateType"
}

func (dateyeartype *DateYearType) GongGetGongstructName() string {
	return "DateYearType"
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

func (valdepthtype *ValDepthType) GongGetGongstructName() string {
	return "ValDepthType"
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

func (valfltype *ValFLType) GongGetGongstructName() string {
	return "ValFLType"
}

func (valfrictiontype *ValFrictionType) GongGetGongstructName() string {
	return "ValFrictionType"
}

func (vallcntype *ValLCNType) GongGetGongstructName() string {
	return "ValLCNType"
}

func (vallightintensitytype *ValLightIntensityType) GongGetGongstructName() string {
	return "ValLightIntensityType"
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetGongstructName() string {
	return "ValMagneticVariationChangeType"
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetGongstructName() string {
	return "ValMagneticVariationType"
}

func (valpcntype *ValPCNType) GongGetGongstructName() string {
	return "ValPCNType"
}

func (valpercenttype *ValPercentType) GongGetGongstructName() string {
	return "ValPercentType"
}

func (valpressuretype *ValPressureType) GongGetGongstructName() string {
	return "ValPressureType"
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
	__gong__rebuildMapString(stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftStandTypes, &stage.CodeAircraftStandTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirportHeliportCollocationTypes, &stage.CodeAirportHeliportCollocationTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirportHeliportDesignatorTypes, &stage.CodeAirportHeliportDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirportHeliportTypes, &stage.CodeAirportHeliportTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirportWarningTypes, &stage.CodeAirportWarningTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachLightingICAOTypes, &stage.CodeApproachLightingICAOTypes_mapString)

	__gong__rebuildMapString(stage.CodeApproachLightingTypes, &stage.CodeApproachLightingTypes_mapString)

	__gong__rebuildMapString(stage.CodeApronElementTypes, &stage.CodeApronElementTypes_mapString)

	__gong__rebuildMapString(stage.CodeApronSectionTypes, &stage.CodeApronSectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeArrestingGearEnergyAbsorbTypes, &stage.CodeArrestingGearEnergyAbsorbTypes_mapString)

	__gong__rebuildMapString(stage.CodeArrestingGearEngageDeviceTypes, &stage.CodeArrestingGearEngageDeviceTypes_mapString)

	__gong__rebuildMapString(stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeBuoyDesignatorTypes, &stage.CodeBuoyDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeBuoyTypes, &stage.CodeBuoyTypes_mapString)

	__gong__rebuildMapString(stage.CodeColourTypes, &stage.CodeColourTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString)

	__gong__rebuildMapString(stage.CodeContaminationTypes, &stage.CodeContaminationTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeDeclaredDistanceTypes, &stage.CodeDeclaredDistanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_mapString)

	__gong__rebuildMapString(stage.CodeDirectionTypes, &stage.CodeDirectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_mapString)

	__gong__rebuildMapString(stage.CodeFlightTypes, &stage.CodeFlightTypes_mapString)

	__gong__rebuildMapString(stage.CodeFrictionDeviceTypes, &stage.CodeFrictionDeviceTypes_mapString)

	__gong__rebuildMapString(stage.CodeFrictionEstimateTypes, &stage.CodeFrictionEstimateTypes_mapString)

	__gong__rebuildMapString(stage.CodeGradeSeparationTypes, &stage.CodeGradeSeparationTypes_mapString)

	__gong__rebuildMapString(stage.CodeGuidanceLineTypes, &stage.CodeGuidanceLineTypes_mapString)

	__gong__rebuildMapString(stage.CodeHelicopterPerformanceTypes, &stage.CodeHelicopterPerformanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeHoldingCategoryTypes, &stage.CodeHoldingCategoryTypes_mapString)

	__gong__rebuildMapString(stage.CodeIATATypes, &stage.CodeIATATypes_mapString)

	__gong__rebuildMapString(stage.CodeICAOTypes, &stage.CodeICAOTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightHoldingPositionTypes, &stage.CodeLightHoldingPositionTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightingJARTypes, &stage.CodeLightingJARTypes_mapString)

	__gong__rebuildMapString(stage.CodeLoadingBridgeTypes, &stage.CodeLoadingBridgeTypes_mapString)

	__gong__rebuildMapString(stage.CodeLogicalOperatorTypes, &stage.CodeLogicalOperatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeMarkingConditionTypes, &stage.CodeMarkingConditionTypes_mapString)

	__gong__rebuildMapString(stage.CodeMarkingStyleTypes, &stage.CodeMarkingStyleTypes_mapString)

	__gong__rebuildMapString(stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeOperationAirportHeliportTypes, &stage.CodeOperationAirportHeliportTypes_mapString)

	__gong__rebuildMapString(stage.CodeOperationManoeuvringAreaTypes, &stage.CodeOperationManoeuvringAreaTypes_mapString)

	__gong__rebuildMapString(stage.CodePCNMethodTypes, &stage.CodePCNMethodTypes_mapString)

	__gong__rebuildMapString(stage.CodePCNPavementTypes, &stage.CodePCNPavementTypes_mapString)

	__gong__rebuildMapString(stage.CodePCNSubgradeTypes, &stage.CodePCNSubgradeTypes_mapString)

	__gong__rebuildMapString(stage.CodePCNTyrePressureTypes, &stage.CodePCNTyrePressureTypes_mapString)

	__gong__rebuildMapString(stage.CodeProtectAreaSectionTypes, &stage.CodeProtectAreaSectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVRReadingTypes, &stage.CodeRVRReadingTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString)

	__gong__rebuildMapString(stage.CodeRoadTypes, &stage.CodeRoadTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwayElementTypes, &stage.CodeRunwayElementTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwayMarkingTypes, &stage.CodeRunwayMarkingTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwayPointRoleTypes, &stage.CodeRunwayPointRoleTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwayProtectionAreaTypes, &stage.CodeRunwayProtectionAreaTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwaySectionTypes, &stage.CodeRunwaySectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeRunwayTypes, &stage.CodeRunwayTypes_mapString)

	__gong__rebuildMapString(stage.CodeSideTypes, &stage.CodeSideTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusAirportTypes, &stage.CodeStatusAirportTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeSurfaceCompositionTypes, &stage.CodeSurfaceCompositionTypes_mapString)

	__gong__rebuildMapString(stage.CodeSurfaceConditionTypes, &stage.CodeSurfaceConditionTypes_mapString)

	__gong__rebuildMapString(stage.CodeSurfacePreparationTypes, &stage.CodeSurfacePreparationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTLOFSectionTypes, &stage.CodeTLOFSectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeTaxiwayElementTypes, &stage.CodeTaxiwayElementTypes_mapString)

	__gong__rebuildMapString(stage.CodeTaxiwaySectionTypes, &stage.CodeTaxiwaySectionTypes_mapString)

	__gong__rebuildMapString(stage.CodeTaxiwayTypes, &stage.CodeTaxiwayTypes_mapString)

	__gong__rebuildMapString(stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString)

	__gong__rebuildMapString(stage.CodeUsageLimitationTypes, &stage.CodeUsageLimitationTypes_mapString)

	__gong__rebuildMapString(stage.CodeVASISTypes, &stage.CodeVASISTypes_mapString)

	__gong__rebuildMapString(stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeVisualDockingGuidanceTypes, &stage.CodeVisualDockingGuidanceTypes_mapString)

	__gong__rebuildMapString(stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeWorkAreaTypes, &stage.CodeWorkAreaTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.DateTimeTypes, &stage.DateTimeTypes_mapString)

	__gong__rebuildMapString(stage.DateTypes, &stage.DateTypes_mapString)

	__gong__rebuildMapString(stage.DateYearTypes, &stage.DateYearTypes_mapString)

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

	__gong__rebuildMapString(stage.ValDepthTypes, &stage.ValDepthTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValFLTypes, &stage.ValFLTypes_mapString)

	__gong__rebuildMapString(stage.ValFrictionTypes, &stage.ValFrictionTypes_mapString)

	__gong__rebuildMapString(stage.ValLCNTypes, &stage.ValLCNTypes_mapString)

	__gong__rebuildMapString(stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_mapString)

	__gong__rebuildMapString(stage.ValMagneticVariationChangeTypes, &stage.ValMagneticVariationChangeTypes_mapString)

	__gong__rebuildMapString(stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_mapString)

	__gong__rebuildMapString(stage.ValPCNTypes, &stage.ValPCNTypes_mapString)

	__gong__rebuildMapString(stage.ValPercentTypes, &stage.ValPercentTypes_mapString)

	__gong__rebuildMapString(stage.ValPressureTypes, &stage.ValPressureTypes_mapString)

	__gong__rebuildMapString(stage.ValSlopeTypes, &stage.ValSlopeTypes_mapString)

	__gong__rebuildMapString(stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString)

	__gong__rebuildMapString(stage.ValTemperatureTypes, &stage.ValTemperatureTypes_mapString)

	__gong__rebuildMapString(stage.ValWeightTypes, &stage.ValWeightTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
