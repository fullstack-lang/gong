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

	CodeAirspaceActivityTypes                map[*CodeAirspaceActivityType]struct{}
	CodeAirspaceActivityTypes_instance       map[*CodeAirspaceActivityType]*CodeAirspaceActivityType
	CodeAirspaceActivityTypes_mapString      map[string]*CodeAirspaceActivityType
	CodeAirspaceActivityTypeOrder            uint
	CodeAirspaceActivityType_stagedOrder     map[*CodeAirspaceActivityType]uint
	CodeAirspaceActivityType_orderStaged     map[uint]*CodeAirspaceActivityType
	CodeAirspaceActivityTypes_reference      map[*CodeAirspaceActivityType]*CodeAirspaceActivityType
	CodeAirspaceActivityTypes_referenceOrder map[*CodeAirspaceActivityType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceActivityTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceActivityType]
	OnAfterCodeAirspaceActivityTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceActivityType]
	OnAfterCodeAirspaceActivityTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceActivityType]

	CodeAirspaceAggregationTypes                map[*CodeAirspaceAggregationType]struct{}
	CodeAirspaceAggregationTypes_instance       map[*CodeAirspaceAggregationType]*CodeAirspaceAggregationType
	CodeAirspaceAggregationTypes_mapString      map[string]*CodeAirspaceAggregationType
	CodeAirspaceAggregationTypeOrder            uint
	CodeAirspaceAggregationType_stagedOrder     map[*CodeAirspaceAggregationType]uint
	CodeAirspaceAggregationType_orderStaged     map[uint]*CodeAirspaceAggregationType
	CodeAirspaceAggregationTypes_reference      map[*CodeAirspaceAggregationType]*CodeAirspaceAggregationType
	CodeAirspaceAggregationTypes_referenceOrder map[*CodeAirspaceAggregationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceAggregationTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceAggregationType]
	OnAfterCodeAirspaceAggregationTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceAggregationType]
	OnAfterCodeAirspaceAggregationTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceAggregationType]

	CodeAirspaceClassificationTypes                map[*CodeAirspaceClassificationType]struct{}
	CodeAirspaceClassificationTypes_instance       map[*CodeAirspaceClassificationType]*CodeAirspaceClassificationType
	CodeAirspaceClassificationTypes_mapString      map[string]*CodeAirspaceClassificationType
	CodeAirspaceClassificationTypeOrder            uint
	CodeAirspaceClassificationType_stagedOrder     map[*CodeAirspaceClassificationType]uint
	CodeAirspaceClassificationType_orderStaged     map[uint]*CodeAirspaceClassificationType
	CodeAirspaceClassificationTypes_reference      map[*CodeAirspaceClassificationType]*CodeAirspaceClassificationType
	CodeAirspaceClassificationTypes_referenceOrder map[*CodeAirspaceClassificationType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceClassificationTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceClassificationType]
	OnAfterCodeAirspaceClassificationTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceClassificationType]
	OnAfterCodeAirspaceClassificationTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceClassificationType]

	CodeAirspaceDependencyTypes                map[*CodeAirspaceDependencyType]struct{}
	CodeAirspaceDependencyTypes_instance       map[*CodeAirspaceDependencyType]*CodeAirspaceDependencyType
	CodeAirspaceDependencyTypes_mapString      map[string]*CodeAirspaceDependencyType
	CodeAirspaceDependencyTypeOrder            uint
	CodeAirspaceDependencyType_stagedOrder     map[*CodeAirspaceDependencyType]uint
	CodeAirspaceDependencyType_orderStaged     map[uint]*CodeAirspaceDependencyType
	CodeAirspaceDependencyTypes_reference      map[*CodeAirspaceDependencyType]*CodeAirspaceDependencyType
	CodeAirspaceDependencyTypes_referenceOrder map[*CodeAirspaceDependencyType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceDependencyTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceDependencyType]
	OnAfterCodeAirspaceDependencyTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceDependencyType]
	OnAfterCodeAirspaceDependencyTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceDependencyType]

	CodeAirspaceDesignatorTypes                map[*CodeAirspaceDesignatorType]struct{}
	CodeAirspaceDesignatorTypes_instance       map[*CodeAirspaceDesignatorType]*CodeAirspaceDesignatorType
	CodeAirspaceDesignatorTypes_mapString      map[string]*CodeAirspaceDesignatorType
	CodeAirspaceDesignatorTypeOrder            uint
	CodeAirspaceDesignatorType_stagedOrder     map[*CodeAirspaceDesignatorType]uint
	CodeAirspaceDesignatorType_orderStaged     map[uint]*CodeAirspaceDesignatorType
	CodeAirspaceDesignatorTypes_reference      map[*CodeAirspaceDesignatorType]*CodeAirspaceDesignatorType
	CodeAirspaceDesignatorTypes_referenceOrder map[*CodeAirspaceDesignatorType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceDesignatorTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceDesignatorType]
	OnAfterCodeAirspaceDesignatorTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceDesignatorType]
	OnAfterCodeAirspaceDesignatorTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceDesignatorType]

	CodeAirspaceTypes                map[*CodeAirspaceType]struct{}
	CodeAirspaceTypes_instance       map[*CodeAirspaceType]*CodeAirspaceType
	CodeAirspaceTypes_mapString      map[string]*CodeAirspaceType
	CodeAirspaceTypeOrder            uint
	CodeAirspaceType_stagedOrder     map[*CodeAirspaceType]uint
	CodeAirspaceType_orderStaged     map[uint]*CodeAirspaceType
	CodeAirspaceTypes_reference      map[*CodeAirspaceType]*CodeAirspaceType
	CodeAirspaceTypes_referenceOrder map[*CodeAirspaceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeAirspaceTypeCreateCallback GongOnAfterCreateInterface[CodeAirspaceType]
	OnAfterCodeAirspaceTypeUpdateCallback GongOnAfterUpdateInterface[CodeAirspaceType]
	OnAfterCodeAirspaceTypeDeleteCallback GongOnAfterDeleteInterface[CodeAirspaceType]

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

	CodeGeoBorderTypes                map[*CodeGeoBorderType]struct{}
	CodeGeoBorderTypes_instance       map[*CodeGeoBorderType]*CodeGeoBorderType
	CodeGeoBorderTypes_mapString      map[string]*CodeGeoBorderType
	CodeGeoBorderTypeOrder            uint
	CodeGeoBorderType_stagedOrder     map[*CodeGeoBorderType]uint
	CodeGeoBorderType_orderStaged     map[uint]*CodeGeoBorderType
	CodeGeoBorderTypes_reference      map[*CodeGeoBorderType]*CodeGeoBorderType
	CodeGeoBorderTypes_referenceOrder map[*CodeGeoBorderType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeGeoBorderTypeCreateCallback GongOnAfterCreateInterface[CodeGeoBorderType]
	OnAfterCodeGeoBorderTypeUpdateCallback GongOnAfterUpdateInterface[CodeGeoBorderType]
	OnAfterCodeGeoBorderTypeDeleteCallback GongOnAfterDeleteInterface[CodeGeoBorderType]

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

	CodeStatusAirspaceTypes                map[*CodeStatusAirspaceType]struct{}
	CodeStatusAirspaceTypes_instance       map[*CodeStatusAirspaceType]*CodeStatusAirspaceType
	CodeStatusAirspaceTypes_mapString      map[string]*CodeStatusAirspaceType
	CodeStatusAirspaceTypeOrder            uint
	CodeStatusAirspaceType_stagedOrder     map[*CodeStatusAirspaceType]uint
	CodeStatusAirspaceType_orderStaged     map[uint]*CodeStatusAirspaceType
	CodeStatusAirspaceTypes_reference      map[*CodeStatusAirspaceType]*CodeStatusAirspaceType
	CodeStatusAirspaceTypes_referenceOrder map[*CodeStatusAirspaceType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusAirspaceTypeCreateCallback GongOnAfterCreateInterface[CodeStatusAirspaceType]
	OnAfterCodeStatusAirspaceTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusAirspaceType]
	OnAfterCodeStatusAirspaceTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusAirspaceType]

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

	__gong__clearReferences(&stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_instance, &stage.CodeAircraftTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_instance, &stage.CodeAircraftWingspanClassTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceActivityTypes_reference, &stage.CodeAirspaceActivityTypes_instance, &stage.CodeAirspaceActivityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceAggregationTypes_reference, &stage.CodeAirspaceAggregationTypes_instance, &stage.CodeAirspaceAggregationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceClassificationTypes_reference, &stage.CodeAirspaceClassificationTypes_instance, &stage.CodeAirspaceClassificationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceDependencyTypes_reference, &stage.CodeAirspaceDependencyTypes_instance, &stage.CodeAirspaceDependencyTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceDesignatorTypes_reference, &stage.CodeAirspaceDesignatorTypes_instance, &stage.CodeAirspaceDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAirspaceTypes_reference, &stage.CodeAirspaceTypes_instance, &stage.CodeAirspaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_instance, &stage.CodeAltitudeUseTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeAuthorityTypes_reference, &stage.CodeAuthorityTypes_instance, &stage.CodeAuthorityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_instance, &stage.CodeCommunicationModeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_instance, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeGeoBorderTypes_reference, &stage.CodeGeoBorderTypes_instance, &stage.CodeGeoBorderTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_instance, &stage.CodeMilitaryOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_instance, &stage.CodeNavigationEquipmentTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_instance, &stage.CodeNavigationSpecificationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_instance, &stage.CodeRVSMTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusAirspaceTypes_reference, &stage.CodeStatusAirspaceTypes_instance, &stage.CodeStatusAirspaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_instance, &stage.CodeTransponderTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_instance, &stage.CodeValueInterpretationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_instance, &stage.CodeVerticalReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_instance, &stage.CodeWakeTurbulenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.NoNumberTypes_reference, &stage.NoNumberTypes_instance, &stage.NoNumberTypes_referenceOrder)

	__gong__clearReferences(&stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_instance, &stage.NoSequenceTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNameTypes_reference, &stage.TextNameTypes_instance, &stage.TextNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNoteTypes_reference, &stage.TextNoteTypes_instance, &stage.TextNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_instance, &stage.TextPropertyNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValFLTypes_reference, &stage.ValFLTypes_instance, &stage.ValFLTypes_referenceOrder)

	__gong__clearReferences(&stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_instance, &stage.ValSpeedTypes_referenceOrder)

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

	stage.CodeAircraftTypeOrder = __gong__recomputeOrder(stage.CodeAircraftType_stagedOrder)

	stage.CodeAircraftWingspanClassTypeOrder = __gong__recomputeOrder(stage.CodeAircraftWingspanClassType_stagedOrder)

	stage.CodeAirspaceActivityTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceActivityType_stagedOrder)

	stage.CodeAirspaceAggregationTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceAggregationType_stagedOrder)

	stage.CodeAirspaceClassificationTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceClassificationType_stagedOrder)

	stage.CodeAirspaceDependencyTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceDependencyType_stagedOrder)

	stage.CodeAirspaceDesignatorTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceDesignatorType_stagedOrder)

	stage.CodeAirspaceTypeOrder = __gong__recomputeOrder(stage.CodeAirspaceType_stagedOrder)

	stage.CodeAltitudeUseTypeOrder = __gong__recomputeOrder(stage.CodeAltitudeUseType_stagedOrder)

	stage.CodeAuthorityTypeOrder = __gong__recomputeOrder(stage.CodeAuthorityType_stagedOrder)

	stage.CodeCommunicationModeTypeOrder = __gong__recomputeOrder(stage.CodeCommunicationModeType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeEquipmentAntiCollisionTypeOrder = __gong__recomputeOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder)

	stage.CodeGeoBorderTypeOrder = __gong__recomputeOrder(stage.CodeGeoBorderType_stagedOrder)

	stage.CodeMilitaryOperationsTypeOrder = __gong__recomputeOrder(stage.CodeMilitaryOperationsType_stagedOrder)

	stage.CodeNavigationEquipmentTypeOrder = __gong__recomputeOrder(stage.CodeNavigationEquipmentType_stagedOrder)

	stage.CodeNavigationSpecificationTypeOrder = __gong__recomputeOrder(stage.CodeNavigationSpecificationType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeRVSMTypeOrder = __gong__recomputeOrder(stage.CodeRVSMType_stagedOrder)

	stage.CodeStatusAirspaceTypeOrder = __gong__recomputeOrder(stage.CodeStatusAirspaceType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeTransponderTypeOrder = __gong__recomputeOrder(stage.CodeTransponderType_stagedOrder)

	stage.CodeValueInterpretationTypeOrder = __gong__recomputeOrder(stage.CodeValueInterpretationType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeVerticalReferenceTypeOrder = __gong__recomputeOrder(stage.CodeVerticalReferenceType_stagedOrder)

	stage.CodeWakeTurbulenceTypeOrder = __gong__recomputeOrder(stage.CodeWakeTurbulenceType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.NoNumberTypeOrder = __gong__recomputeOrder(stage.NoNumberType_stagedOrder)

	stage.NoSequenceTypeOrder = __gong__recomputeOrder(stage.NoSequenceType_stagedOrder)

	stage.TextNameTypeOrder = __gong__recomputeOrder(stage.TextNameType_stagedOrder)

	stage.TextNoteTypeOrder = __gong__recomputeOrder(stage.TextNoteType_stagedOrder)

	stage.TextPropertyNameTypeOrder = __gong__recomputeOrder(stage.TextPropertyNameType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValFLTypeOrder = __gong__recomputeOrder(stage.ValFLType_stagedOrder)

	stage.ValSpeedTypeOrder = __gong__recomputeOrder(stage.ValSpeedType_stagedOrder)

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

func (*CodeAirspaceActivityType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityType_stagedOrder)
}

func (*CodeAirspaceActivityType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceActivityType_orderStaged[order]
}

func (*CodeAirspaceActivityType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceActivityTypes_mapString
}

func (*CodeAirspaceActivityType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceActivityTypes
}

func (*CodeAirspaceActivityType) GongNewInstance() any {
	return new(CodeAirspaceActivityType)
}

func (*CodeAirspaceAggregationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationType_stagedOrder)
}

func (*CodeAirspaceAggregationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceAggregationType_orderStaged[order]
}

func (*CodeAirspaceAggregationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceAggregationTypes_mapString
}

func (*CodeAirspaceAggregationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceAggregationTypes
}

func (*CodeAirspaceAggregationType) GongNewInstance() any {
	return new(CodeAirspaceAggregationType)
}

func (*CodeAirspaceClassificationType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationType_stagedOrder)
}

func (*CodeAirspaceClassificationType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceClassificationType_orderStaged[order]
}

func (*CodeAirspaceClassificationType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceClassificationTypes_mapString
}

func (*CodeAirspaceClassificationType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceClassificationTypes
}

func (*CodeAirspaceClassificationType) GongNewInstance() any {
	return new(CodeAirspaceClassificationType)
}

func (*CodeAirspaceDependencyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyType_stagedOrder)
}

func (*CodeAirspaceDependencyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceDependencyType_orderStaged[order]
}

func (*CodeAirspaceDependencyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceDependencyTypes_mapString
}

func (*CodeAirspaceDependencyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceDependencyTypes
}

func (*CodeAirspaceDependencyType) GongNewInstance() any {
	return new(CodeAirspaceDependencyType)
}

func (*CodeAirspaceDesignatorType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorType_stagedOrder)
}

func (*CodeAirspaceDesignatorType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceDesignatorType_orderStaged[order]
}

func (*CodeAirspaceDesignatorType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceDesignatorTypes_mapString
}

func (*CodeAirspaceDesignatorType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceDesignatorTypes
}

func (*CodeAirspaceDesignatorType) GongNewInstance() any {
	return new(CodeAirspaceDesignatorType)
}

func (*CodeAirspaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeAirspaceTypes, stage.CodeAirspaceType_stagedOrder)
}

func (*CodeAirspaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeAirspaceType_orderStaged[order]
}

func (*CodeAirspaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeAirspaceTypes_mapString
}

func (*CodeAirspaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeAirspaceTypes
}

func (*CodeAirspaceType) GongNewInstance() any {
	return new(CodeAirspaceType)
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

func (*CodeGeoBorderType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeGeoBorderTypes, stage.CodeGeoBorderType_stagedOrder)
}

func (*CodeGeoBorderType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeGeoBorderType_orderStaged[order]
}

func (*CodeGeoBorderType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeGeoBorderTypes_mapString
}

func (*CodeGeoBorderType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeGeoBorderTypes
}

func (*CodeGeoBorderType) GongNewInstance() any {
	return new(CodeGeoBorderType)
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

func (*CodeStatusAirspaceType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceType_stagedOrder)
}

func (*CodeStatusAirspaceType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusAirspaceType_orderStaged[order]
}

func (*CodeStatusAirspaceType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusAirspaceTypes_mapString
}

func (*CodeStatusAirspaceType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusAirspaceTypes
}

func (*CodeStatusAirspaceType) GongNewInstance() any {
	return new(CodeStatusAirspaceType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
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

		CodeAircraftTypes:           make(map[*CodeAircraftType]struct{}),
		CodeAircraftTypes_mapString: make(map[string]*CodeAircraftType),

		CodeAircraftWingspanClassTypes:           make(map[*CodeAircraftWingspanClassType]struct{}),
		CodeAircraftWingspanClassTypes_mapString: make(map[string]*CodeAircraftWingspanClassType),

		CodeAirspaceActivityTypes:           make(map[*CodeAirspaceActivityType]struct{}),
		CodeAirspaceActivityTypes_mapString: make(map[string]*CodeAirspaceActivityType),

		CodeAirspaceAggregationTypes:           make(map[*CodeAirspaceAggregationType]struct{}),
		CodeAirspaceAggregationTypes_mapString: make(map[string]*CodeAirspaceAggregationType),

		CodeAirspaceClassificationTypes:           make(map[*CodeAirspaceClassificationType]struct{}),
		CodeAirspaceClassificationTypes_mapString: make(map[string]*CodeAirspaceClassificationType),

		CodeAirspaceDependencyTypes:           make(map[*CodeAirspaceDependencyType]struct{}),
		CodeAirspaceDependencyTypes_mapString: make(map[string]*CodeAirspaceDependencyType),

		CodeAirspaceDesignatorTypes:           make(map[*CodeAirspaceDesignatorType]struct{}),
		CodeAirspaceDesignatorTypes_mapString: make(map[string]*CodeAirspaceDesignatorType),

		CodeAirspaceTypes:           make(map[*CodeAirspaceType]struct{}),
		CodeAirspaceTypes_mapString: make(map[string]*CodeAirspaceType),

		CodeAltitudeUseTypes:           make(map[*CodeAltitudeUseType]struct{}),
		CodeAltitudeUseTypes_mapString: make(map[string]*CodeAltitudeUseType),

		CodeAuthorityTypes:           make(map[*CodeAuthorityType]struct{}),
		CodeAuthorityTypes_mapString: make(map[string]*CodeAuthorityType),

		CodeCommunicationModeTypes:           make(map[*CodeCommunicationModeType]struct{}),
		CodeCommunicationModeTypes_mapString: make(map[string]*CodeCommunicationModeType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeEquipmentAntiCollisionTypes:           make(map[*CodeEquipmentAntiCollisionType]struct{}),
		CodeEquipmentAntiCollisionTypes_mapString: make(map[string]*CodeEquipmentAntiCollisionType),

		CodeGeoBorderTypes:           make(map[*CodeGeoBorderType]struct{}),
		CodeGeoBorderTypes_mapString: make(map[string]*CodeGeoBorderType),

		CodeMilitaryOperationsTypes:           make(map[*CodeMilitaryOperationsType]struct{}),
		CodeMilitaryOperationsTypes_mapString: make(map[string]*CodeMilitaryOperationsType),

		CodeNavigationEquipmentTypes:           make(map[*CodeNavigationEquipmentType]struct{}),
		CodeNavigationEquipmentTypes_mapString: make(map[string]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationTypes:           make(map[*CodeNavigationSpecificationType]struct{}),
		CodeNavigationSpecificationTypes_mapString: make(map[string]*CodeNavigationSpecificationType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodeRVSMTypes:           make(map[*CodeRVSMType]struct{}),
		CodeRVSMTypes_mapString: make(map[string]*CodeRVSMType),

		CodeStatusAirspaceTypes:           make(map[*CodeStatusAirspaceType]struct{}),
		CodeStatusAirspaceTypes_mapString: make(map[string]*CodeStatusAirspaceType),

		CodeTimeEventCombinationTypes:           make(map[*CodeTimeEventCombinationType]struct{}),
		CodeTimeEventCombinationTypes_mapString: make(map[string]*CodeTimeEventCombinationType),

		CodeTimeEventTypes:           make(map[*CodeTimeEventType]struct{}),
		CodeTimeEventTypes_mapString: make(map[string]*CodeTimeEventType),

		CodeTimeReferenceTypes:           make(map[*CodeTimeReferenceType]struct{}),
		CodeTimeReferenceTypes_mapString: make(map[string]*CodeTimeReferenceType),

		CodeTransponderTypes:           make(map[*CodeTransponderType]struct{}),
		CodeTransponderTypes_mapString: make(map[string]*CodeTransponderType),

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

		TextNameTypes:           make(map[*TextNameType]struct{}),
		TextNameTypes_mapString: make(map[string]*TextNameType),

		TextNoteTypes:           make(map[*TextNoteType]struct{}),
		TextNoteTypes_mapString: make(map[string]*TextNoteType),

		TextPropertyNameTypes:           make(map[*TextPropertyNameType]struct{}),
		TextPropertyNameTypes_mapString: make(map[string]*TextPropertyNameType),

		TimeTypes:           make(map[*TimeType]struct{}),
		TimeTypes_mapString: make(map[string]*TimeType),

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

		ValSpeedTypes:           make(map[*ValSpeedType]struct{}),
		ValSpeedTypes_mapString: make(map[string]*ValSpeedType),

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

		CodeAircraftType_stagedOrder: make(map[*CodeAircraftType]uint),
		CodeAircraftType_orderStaged: make(map[uint]*CodeAircraftType),
		CodeAircraftTypes_reference:  make(map[*CodeAircraftType]*CodeAircraftType),

		CodeAircraftWingspanClassType_stagedOrder: make(map[*CodeAircraftWingspanClassType]uint),
		CodeAircraftWingspanClassType_orderStaged: make(map[uint]*CodeAircraftWingspanClassType),
		CodeAircraftWingspanClassTypes_reference:  make(map[*CodeAircraftWingspanClassType]*CodeAircraftWingspanClassType),

		CodeAirspaceActivityType_stagedOrder: make(map[*CodeAirspaceActivityType]uint),
		CodeAirspaceActivityType_orderStaged: make(map[uint]*CodeAirspaceActivityType),
		CodeAirspaceActivityTypes_reference:  make(map[*CodeAirspaceActivityType]*CodeAirspaceActivityType),

		CodeAirspaceAggregationType_stagedOrder: make(map[*CodeAirspaceAggregationType]uint),
		CodeAirspaceAggregationType_orderStaged: make(map[uint]*CodeAirspaceAggregationType),
		CodeAirspaceAggregationTypes_reference:  make(map[*CodeAirspaceAggregationType]*CodeAirspaceAggregationType),

		CodeAirspaceClassificationType_stagedOrder: make(map[*CodeAirspaceClassificationType]uint),
		CodeAirspaceClassificationType_orderStaged: make(map[uint]*CodeAirspaceClassificationType),
		CodeAirspaceClassificationTypes_reference:  make(map[*CodeAirspaceClassificationType]*CodeAirspaceClassificationType),

		CodeAirspaceDependencyType_stagedOrder: make(map[*CodeAirspaceDependencyType]uint),
		CodeAirspaceDependencyType_orderStaged: make(map[uint]*CodeAirspaceDependencyType),
		CodeAirspaceDependencyTypes_reference:  make(map[*CodeAirspaceDependencyType]*CodeAirspaceDependencyType),

		CodeAirspaceDesignatorType_stagedOrder: make(map[*CodeAirspaceDesignatorType]uint),
		CodeAirspaceDesignatorType_orderStaged: make(map[uint]*CodeAirspaceDesignatorType),
		CodeAirspaceDesignatorTypes_reference:  make(map[*CodeAirspaceDesignatorType]*CodeAirspaceDesignatorType),

		CodeAirspaceType_stagedOrder: make(map[*CodeAirspaceType]uint),
		CodeAirspaceType_orderStaged: make(map[uint]*CodeAirspaceType),
		CodeAirspaceTypes_reference:  make(map[*CodeAirspaceType]*CodeAirspaceType),

		CodeAltitudeUseType_stagedOrder: make(map[*CodeAltitudeUseType]uint),
		CodeAltitudeUseType_orderStaged: make(map[uint]*CodeAltitudeUseType),
		CodeAltitudeUseTypes_reference:  make(map[*CodeAltitudeUseType]*CodeAltitudeUseType),

		CodeAuthorityType_stagedOrder: make(map[*CodeAuthorityType]uint),
		CodeAuthorityType_orderStaged: make(map[uint]*CodeAuthorityType),
		CodeAuthorityTypes_reference:  make(map[*CodeAuthorityType]*CodeAuthorityType),

		CodeCommunicationModeType_stagedOrder: make(map[*CodeCommunicationModeType]uint),
		CodeCommunicationModeType_orderStaged: make(map[uint]*CodeCommunicationModeType),
		CodeCommunicationModeTypes_reference:  make(map[*CodeCommunicationModeType]*CodeCommunicationModeType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeEquipmentAntiCollisionType_stagedOrder: make(map[*CodeEquipmentAntiCollisionType]uint),
		CodeEquipmentAntiCollisionType_orderStaged: make(map[uint]*CodeEquipmentAntiCollisionType),
		CodeEquipmentAntiCollisionTypes_reference:  make(map[*CodeEquipmentAntiCollisionType]*CodeEquipmentAntiCollisionType),

		CodeGeoBorderType_stagedOrder: make(map[*CodeGeoBorderType]uint),
		CodeGeoBorderType_orderStaged: make(map[uint]*CodeGeoBorderType),
		CodeGeoBorderTypes_reference:  make(map[*CodeGeoBorderType]*CodeGeoBorderType),

		CodeMilitaryOperationsType_stagedOrder: make(map[*CodeMilitaryOperationsType]uint),
		CodeMilitaryOperationsType_orderStaged: make(map[uint]*CodeMilitaryOperationsType),
		CodeMilitaryOperationsTypes_reference:  make(map[*CodeMilitaryOperationsType]*CodeMilitaryOperationsType),

		CodeNavigationEquipmentType_stagedOrder: make(map[*CodeNavigationEquipmentType]uint),
		CodeNavigationEquipmentType_orderStaged: make(map[uint]*CodeNavigationEquipmentType),
		CodeNavigationEquipmentTypes_reference:  make(map[*CodeNavigationEquipmentType]*CodeNavigationEquipmentType),

		CodeNavigationSpecificationType_stagedOrder: make(map[*CodeNavigationSpecificationType]uint),
		CodeNavigationSpecificationType_orderStaged: make(map[uint]*CodeNavigationSpecificationType),
		CodeNavigationSpecificationTypes_reference:  make(map[*CodeNavigationSpecificationType]*CodeNavigationSpecificationType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodeRVSMType_stagedOrder: make(map[*CodeRVSMType]uint),
		CodeRVSMType_orderStaged: make(map[uint]*CodeRVSMType),
		CodeRVSMTypes_reference:  make(map[*CodeRVSMType]*CodeRVSMType),

		CodeStatusAirspaceType_stagedOrder: make(map[*CodeStatusAirspaceType]uint),
		CodeStatusAirspaceType_orderStaged: make(map[uint]*CodeStatusAirspaceType),
		CodeStatusAirspaceTypes_reference:  make(map[*CodeStatusAirspaceType]*CodeStatusAirspaceType),

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

		TextNameType_stagedOrder: make(map[*TextNameType]uint),
		TextNameType_orderStaged: make(map[uint]*TextNameType),
		TextNameTypes_reference:  make(map[*TextNameType]*TextNameType),

		TextNoteType_stagedOrder: make(map[*TextNoteType]uint),
		TextNoteType_orderStaged: make(map[uint]*TextNoteType),
		TextNoteTypes_reference:  make(map[*TextNoteType]*TextNoteType),

		TextPropertyNameType_stagedOrder: make(map[*TextPropertyNameType]uint),
		TextPropertyNameType_orderStaged: make(map[uint]*TextPropertyNameType),
		TextPropertyNameTypes_reference:  make(map[*TextPropertyNameType]*TextPropertyNameType),

		TimeType_stagedOrder: make(map[*TimeType]uint),
		TimeType_orderStaged: make(map[uint]*TimeType),
		TimeTypes_reference:  make(map[*TimeType]*TimeType),

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

		ValSpeedType_stagedOrder: make(map[*ValSpeedType]uint),
		ValSpeedType_orderStaged: make(map[uint]*ValSpeedType),
		ValSpeedTypes_reference:  make(map[*ValSpeedType]*ValSpeedType),

		ValWeightType_stagedOrder: make(map[*ValWeightType]uint),
		ValWeightType_orderStaged: make(map[uint]*ValWeightType),
		ValWeightTypes_reference:  make(map[*ValWeightType]*ValWeightType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeAircraftCategoryType": &CodeAircraftCategoryTypeUnmarshaller{},

			"CodeAircraftEngineNumberType": &CodeAircraftEngineNumberTypeUnmarshaller{},

			"CodeAircraftEngineType": &CodeAircraftEngineTypeUnmarshaller{},

			"CodeAircraftICAOType": &CodeAircraftICAOTypeUnmarshaller{},

			"CodeAircraftType": &CodeAircraftTypeUnmarshaller{},

			"CodeAircraftWingspanClassType": &CodeAircraftWingspanClassTypeUnmarshaller{},

			"CodeAirspaceActivityType": &CodeAirspaceActivityTypeUnmarshaller{},

			"CodeAirspaceAggregationType": &CodeAirspaceAggregationTypeUnmarshaller{},

			"CodeAirspaceClassificationType": &CodeAirspaceClassificationTypeUnmarshaller{},

			"CodeAirspaceDependencyType": &CodeAirspaceDependencyTypeUnmarshaller{},

			"CodeAirspaceDesignatorType": &CodeAirspaceDesignatorTypeUnmarshaller{},

			"CodeAirspaceType": &CodeAirspaceTypeUnmarshaller{},

			"CodeAltitudeUseType": &CodeAltitudeUseTypeUnmarshaller{},

			"CodeAuthorityType": &CodeAuthorityTypeUnmarshaller{},

			"CodeCommunicationModeType": &CodeCommunicationModeTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeEquipmentAntiCollisionType": &CodeEquipmentAntiCollisionTypeUnmarshaller{},

			"CodeGeoBorderType": &CodeGeoBorderTypeUnmarshaller{},

			"CodeMilitaryOperationsType": &CodeMilitaryOperationsTypeUnmarshaller{},

			"CodeNavigationEquipmentType": &CodeNavigationEquipmentTypeUnmarshaller{},

			"CodeNavigationSpecificationType": &CodeNavigationSpecificationTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeRVSMType": &CodeRVSMTypeUnmarshaller{},

			"CodeStatusAirspaceType": &CodeStatusAirspaceTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeTransponderType": &CodeTransponderTypeUnmarshaller{},

			"CodeValueInterpretationType": &CodeValueInterpretationTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeVerticalReferenceType": &CodeVerticalReferenceTypeUnmarshaller{},

			"CodeWakeTurbulenceType": &CodeWakeTurbulenceTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"NoNumberType": &NoNumberTypeUnmarshaller{},

			"NoSequenceType": &NoSequenceTypeUnmarshaller{},

			"TextNameType": &TextNameTypeUnmarshaller{},

			"TextNoteType": &TextNoteTypeUnmarshaller{},

			"TextPropertyNameType": &TextPropertyNameTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValFLType": &ValFLTypeUnmarshaller{},

			"ValSpeedType": &ValSpeedTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeAircraftType"] = len(stage.CodeAircraftTypes)
	stage.Map_GongStructName_InstancesNb["CodeAircraftWingspanClassType"] = len(stage.CodeAircraftWingspanClassTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceActivityType"] = len(stage.CodeAirspaceActivityTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceAggregationType"] = len(stage.CodeAirspaceAggregationTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceClassificationType"] = len(stage.CodeAirspaceClassificationTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceDependencyType"] = len(stage.CodeAirspaceDependencyTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceDesignatorType"] = len(stage.CodeAirspaceDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["CodeAirspaceType"] = len(stage.CodeAirspaceTypes)
	stage.Map_GongStructName_InstancesNb["CodeAltitudeUseType"] = len(stage.CodeAltitudeUseTypes)
	stage.Map_GongStructName_InstancesNb["CodeAuthorityType"] = len(stage.CodeAuthorityTypes)
	stage.Map_GongStructName_InstancesNb["CodeCommunicationModeType"] = len(stage.CodeCommunicationModeTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeEquipmentAntiCollisionType"] = len(stage.CodeEquipmentAntiCollisionTypes)
	stage.Map_GongStructName_InstancesNb["CodeGeoBorderType"] = len(stage.CodeGeoBorderTypes)
	stage.Map_GongStructName_InstancesNb["CodeMilitaryOperationsType"] = len(stage.CodeMilitaryOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationEquipmentType"] = len(stage.CodeNavigationEquipmentTypes)
	stage.Map_GongStructName_InstancesNb["CodeNavigationSpecificationType"] = len(stage.CodeNavigationSpecificationTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeRVSMType"] = len(stage.CodeRVSMTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusAirspaceType"] = len(stage.CodeStatusAirspaceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeTransponderType"] = len(stage.CodeTransponderTypes)
	stage.Map_GongStructName_InstancesNb["CodeValueInterpretationType"] = len(stage.CodeValueInterpretationTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalReferenceType"] = len(stage.CodeVerticalReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeWakeTurbulenceType"] = len(stage.CodeWakeTurbulenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["NoNumberType"] = len(stage.NoNumberTypes)
	stage.Map_GongStructName_InstancesNb["NoSequenceType"] = len(stage.NoSequenceTypes)
	stage.Map_GongStructName_InstancesNb["TextNameType"] = len(stage.TextNameTypes)
	stage.Map_GongStructName_InstancesNb["TextNoteType"] = len(stage.TextNoteTypes)
	stage.Map_GongStructName_InstancesNb["TextPropertyNameType"] = len(stage.TextPropertyNameTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValFLType"] = len(stage.ValFLTypes)
	stage.Map_GongStructName_InstancesNb["ValSpeedType"] = len(stage.ValSpeedTypes)
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

// Stage puts codeairspaceactivitytype to the model stage
func (codeairspaceactivitytype *CodeAirspaceActivityType) Stage(stage *Stage) *CodeAirspaceActivityType {
	__gong__stage(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityType_stagedOrder, stage.CodeAirspaceActivityType_orderStaged, &stage.CodeAirspaceActivityTypeOrder, stage.CodeAirspaceActivityTypes_mapString, codeairspaceactivitytype, codeairspaceactivitytype.Name)
	return codeairspaceactivitytype
}

// StagePreserveOrder puts codeairspaceactivitytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceActivityTypeOrder
// - update stage.CodeAirspaceActivityTypeOrder accordingly
func (codeairspaceactivitytype *CodeAirspaceActivityType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityType_stagedOrder, stage.CodeAirspaceActivityType_orderStaged, &stage.CodeAirspaceActivityTypeOrder, stage.CodeAirspaceActivityTypes_mapString, codeairspaceactivitytype, order, codeairspaceactivitytype.Name)
}

// Unstage removes codeairspaceactivitytype off the model stage
func (codeairspaceactivitytype *CodeAirspaceActivityType) Unstage(stage *Stage) *CodeAirspaceActivityType {
	__gong__unstage(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityTypes_mapString, codeairspaceactivitytype, codeairspaceactivitytype.Name)
	return codeairspaceactivitytype
}

// UnstageVoid removes codeairspaceactivitytype off the model stage
func (codeairspaceactivitytype *CodeAirspaceActivityType) UnstageVoid(stage *Stage) {
	codeairspaceactivitytype.Unstage(stage)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) StageVoid(stage *Stage) {
	codeairspaceactivitytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspaceactivitytype *CodeAirspaceActivityType) GetName() (res string) {
	return codeairspaceactivitytype.Name
}

// for satisfaction of GongStruct interface
func (codeairspaceactivitytype *CodeAirspaceActivityType) SetName(name string) {
	codeairspaceactivitytype.Name = name
}

// Stage puts codeairspaceaggregationtype to the model stage
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) Stage(stage *Stage) *CodeAirspaceAggregationType {
	__gong__stage(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationType_stagedOrder, stage.CodeAirspaceAggregationType_orderStaged, &stage.CodeAirspaceAggregationTypeOrder, stage.CodeAirspaceAggregationTypes_mapString, codeairspaceaggregationtype, codeairspaceaggregationtype.Name)
	return codeairspaceaggregationtype
}

// StagePreserveOrder puts codeairspaceaggregationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceAggregationTypeOrder
// - update stage.CodeAirspaceAggregationTypeOrder accordingly
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationType_stagedOrder, stage.CodeAirspaceAggregationType_orderStaged, &stage.CodeAirspaceAggregationTypeOrder, stage.CodeAirspaceAggregationTypes_mapString, codeairspaceaggregationtype, order, codeairspaceaggregationtype.Name)
}

// Unstage removes codeairspaceaggregationtype off the model stage
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) Unstage(stage *Stage) *CodeAirspaceAggregationType {
	__gong__unstage(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationTypes_mapString, codeairspaceaggregationtype, codeairspaceaggregationtype.Name)
	return codeairspaceaggregationtype
}

// UnstageVoid removes codeairspaceaggregationtype off the model stage
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) UnstageVoid(stage *Stage) {
	codeairspaceaggregationtype.Unstage(stage)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) StageVoid(stage *Stage) {
	codeairspaceaggregationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GetName() (res string) {
	return codeairspaceaggregationtype.Name
}

// for satisfaction of GongStruct interface
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) SetName(name string) {
	codeairspaceaggregationtype.Name = name
}

// Stage puts codeairspaceclassificationtype to the model stage
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) Stage(stage *Stage) *CodeAirspaceClassificationType {
	__gong__stage(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationType_stagedOrder, stage.CodeAirspaceClassificationType_orderStaged, &stage.CodeAirspaceClassificationTypeOrder, stage.CodeAirspaceClassificationTypes_mapString, codeairspaceclassificationtype, codeairspaceclassificationtype.Name)
	return codeairspaceclassificationtype
}

// StagePreserveOrder puts codeairspaceclassificationtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceClassificationTypeOrder
// - update stage.CodeAirspaceClassificationTypeOrder accordingly
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationType_stagedOrder, stage.CodeAirspaceClassificationType_orderStaged, &stage.CodeAirspaceClassificationTypeOrder, stage.CodeAirspaceClassificationTypes_mapString, codeairspaceclassificationtype, order, codeairspaceclassificationtype.Name)
}

// Unstage removes codeairspaceclassificationtype off the model stage
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) Unstage(stage *Stage) *CodeAirspaceClassificationType {
	__gong__unstage(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationTypes_mapString, codeairspaceclassificationtype, codeairspaceclassificationtype.Name)
	return codeairspaceclassificationtype
}

// UnstageVoid removes codeairspaceclassificationtype off the model stage
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) UnstageVoid(stage *Stage) {
	codeairspaceclassificationtype.Unstage(stage)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) StageVoid(stage *Stage) {
	codeairspaceclassificationtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GetName() (res string) {
	return codeairspaceclassificationtype.Name
}

// for satisfaction of GongStruct interface
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) SetName(name string) {
	codeairspaceclassificationtype.Name = name
}

// Stage puts codeairspacedependencytype to the model stage
func (codeairspacedependencytype *CodeAirspaceDependencyType) Stage(stage *Stage) *CodeAirspaceDependencyType {
	__gong__stage(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyType_stagedOrder, stage.CodeAirspaceDependencyType_orderStaged, &stage.CodeAirspaceDependencyTypeOrder, stage.CodeAirspaceDependencyTypes_mapString, codeairspacedependencytype, codeairspacedependencytype.Name)
	return codeairspacedependencytype
}

// StagePreserveOrder puts codeairspacedependencytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceDependencyTypeOrder
// - update stage.CodeAirspaceDependencyTypeOrder accordingly
func (codeairspacedependencytype *CodeAirspaceDependencyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyType_stagedOrder, stage.CodeAirspaceDependencyType_orderStaged, &stage.CodeAirspaceDependencyTypeOrder, stage.CodeAirspaceDependencyTypes_mapString, codeairspacedependencytype, order, codeairspacedependencytype.Name)
}

// Unstage removes codeairspacedependencytype off the model stage
func (codeairspacedependencytype *CodeAirspaceDependencyType) Unstage(stage *Stage) *CodeAirspaceDependencyType {
	__gong__unstage(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyTypes_mapString, codeairspacedependencytype, codeairspacedependencytype.Name)
	return codeairspacedependencytype
}

// UnstageVoid removes codeairspacedependencytype off the model stage
func (codeairspacedependencytype *CodeAirspaceDependencyType) UnstageVoid(stage *Stage) {
	codeairspacedependencytype.Unstage(stage)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) StageVoid(stage *Stage) {
	codeairspacedependencytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspacedependencytype *CodeAirspaceDependencyType) GetName() (res string) {
	return codeairspacedependencytype.Name
}

// for satisfaction of GongStruct interface
func (codeairspacedependencytype *CodeAirspaceDependencyType) SetName(name string) {
	codeairspacedependencytype.Name = name
}

// Stage puts codeairspacedesignatortype to the model stage
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) Stage(stage *Stage) *CodeAirspaceDesignatorType {
	__gong__stage(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorType_stagedOrder, stage.CodeAirspaceDesignatorType_orderStaged, &stage.CodeAirspaceDesignatorTypeOrder, stage.CodeAirspaceDesignatorTypes_mapString, codeairspacedesignatortype, codeairspacedesignatortype.Name)
	return codeairspacedesignatortype
}

// StagePreserveOrder puts codeairspacedesignatortype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceDesignatorTypeOrder
// - update stage.CodeAirspaceDesignatorTypeOrder accordingly
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorType_stagedOrder, stage.CodeAirspaceDesignatorType_orderStaged, &stage.CodeAirspaceDesignatorTypeOrder, stage.CodeAirspaceDesignatorTypes_mapString, codeairspacedesignatortype, order, codeairspacedesignatortype.Name)
}

// Unstage removes codeairspacedesignatortype off the model stage
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) Unstage(stage *Stage) *CodeAirspaceDesignatorType {
	__gong__unstage(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorTypes_mapString, codeairspacedesignatortype, codeairspacedesignatortype.Name)
	return codeairspacedesignatortype
}

// UnstageVoid removes codeairspacedesignatortype off the model stage
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) UnstageVoid(stage *Stage) {
	codeairspacedesignatortype.Unstage(stage)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) StageVoid(stage *Stage) {
	codeairspacedesignatortype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GetName() (res string) {
	return codeairspacedesignatortype.Name
}

// for satisfaction of GongStruct interface
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) SetName(name string) {
	codeairspacedesignatortype.Name = name
}

// Stage puts codeairspacetype to the model stage
func (codeairspacetype *CodeAirspaceType) Stage(stage *Stage) *CodeAirspaceType {
	__gong__stage(stage.CodeAirspaceTypes, stage.CodeAirspaceType_stagedOrder, stage.CodeAirspaceType_orderStaged, &stage.CodeAirspaceTypeOrder, stage.CodeAirspaceTypes_mapString, codeairspacetype, codeairspacetype.Name)
	return codeairspacetype
}

// StagePreserveOrder puts codeairspacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeAirspaceTypeOrder
// - update stage.CodeAirspaceTypeOrder accordingly
func (codeairspacetype *CodeAirspaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeAirspaceTypes, stage.CodeAirspaceType_stagedOrder, stage.CodeAirspaceType_orderStaged, &stage.CodeAirspaceTypeOrder, stage.CodeAirspaceTypes_mapString, codeairspacetype, order, codeairspacetype.Name)
}

// Unstage removes codeairspacetype off the model stage
func (codeairspacetype *CodeAirspaceType) Unstage(stage *Stage) *CodeAirspaceType {
	__gong__unstage(stage.CodeAirspaceTypes, stage.CodeAirspaceTypes_mapString, codeairspacetype, codeairspacetype.Name)
	return codeairspacetype
}

// UnstageVoid removes codeairspacetype off the model stage
func (codeairspacetype *CodeAirspaceType) UnstageVoid(stage *Stage) {
	codeairspacetype.Unstage(stage)
}

func (codeairspacetype *CodeAirspaceType) StageVoid(stage *Stage) {
	codeairspacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeairspacetype *CodeAirspaceType) GetName() (res string) {
	return codeairspacetype.Name
}

// for satisfaction of GongStruct interface
func (codeairspacetype *CodeAirspaceType) SetName(name string) {
	codeairspacetype.Name = name
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

// Stage puts codegeobordertype to the model stage
func (codegeobordertype *CodeGeoBorderType) Stage(stage *Stage) *CodeGeoBorderType {
	__gong__stage(stage.CodeGeoBorderTypes, stage.CodeGeoBorderType_stagedOrder, stage.CodeGeoBorderType_orderStaged, &stage.CodeGeoBorderTypeOrder, stage.CodeGeoBorderTypes_mapString, codegeobordertype, codegeobordertype.Name)
	return codegeobordertype
}

// StagePreserveOrder puts codegeobordertype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeGeoBorderTypeOrder
// - update stage.CodeGeoBorderTypeOrder accordingly
func (codegeobordertype *CodeGeoBorderType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeGeoBorderTypes, stage.CodeGeoBorderType_stagedOrder, stage.CodeGeoBorderType_orderStaged, &stage.CodeGeoBorderTypeOrder, stage.CodeGeoBorderTypes_mapString, codegeobordertype, order, codegeobordertype.Name)
}

// Unstage removes codegeobordertype off the model stage
func (codegeobordertype *CodeGeoBorderType) Unstage(stage *Stage) *CodeGeoBorderType {
	__gong__unstage(stage.CodeGeoBorderTypes, stage.CodeGeoBorderTypes_mapString, codegeobordertype, codegeobordertype.Name)
	return codegeobordertype
}

// UnstageVoid removes codegeobordertype off the model stage
func (codegeobordertype *CodeGeoBorderType) UnstageVoid(stage *Stage) {
	codegeobordertype.Unstage(stage)
}

func (codegeobordertype *CodeGeoBorderType) StageVoid(stage *Stage) {
	codegeobordertype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codegeobordertype *CodeGeoBorderType) GetName() (res string) {
	return codegeobordertype.Name
}

// for satisfaction of GongStruct interface
func (codegeobordertype *CodeGeoBorderType) SetName(name string) {
	codegeobordertype.Name = name
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

// Stage puts codestatusairspacetype to the model stage
func (codestatusairspacetype *CodeStatusAirspaceType) Stage(stage *Stage) *CodeStatusAirspaceType {
	__gong__stage(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceType_stagedOrder, stage.CodeStatusAirspaceType_orderStaged, &stage.CodeStatusAirspaceTypeOrder, stage.CodeStatusAirspaceTypes_mapString, codestatusairspacetype, codestatusairspacetype.Name)
	return codestatusairspacetype
}

// StagePreserveOrder puts codestatusairspacetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusAirspaceTypeOrder
// - update stage.CodeStatusAirspaceTypeOrder accordingly
func (codestatusairspacetype *CodeStatusAirspaceType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceType_stagedOrder, stage.CodeStatusAirspaceType_orderStaged, &stage.CodeStatusAirspaceTypeOrder, stage.CodeStatusAirspaceTypes_mapString, codestatusairspacetype, order, codestatusairspacetype.Name)
}

// Unstage removes codestatusairspacetype off the model stage
func (codestatusairspacetype *CodeStatusAirspaceType) Unstage(stage *Stage) *CodeStatusAirspaceType {
	__gong__unstage(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceTypes_mapString, codestatusairspacetype, codestatusairspacetype.Name)
	return codestatusairspacetype
}

// UnstageVoid removes codestatusairspacetype off the model stage
func (codestatusairspacetype *CodeStatusAirspaceType) UnstageVoid(stage *Stage) {
	codestatusairspacetype.Unstage(stage)
}

func (codestatusairspacetype *CodeStatusAirspaceType) StageVoid(stage *Stage) {
	codestatusairspacetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusairspacetype *CodeStatusAirspaceType) GetName() (res string) {
	return codestatusairspacetype.Name
}

// for satisfaction of GongStruct interface
func (codestatusairspacetype *CodeStatusAirspaceType) SetName(name string) {
	codestatusairspacetype.Name = name
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

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_mapString, &stage.CodeAircraftCategoryType_stagedOrder, &stage.CodeAircraftCategoryTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_mapString, &stage.CodeAircraftEngineNumberType_stagedOrder, &stage.CodeAircraftEngineNumberTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_mapString, &stage.CodeAircraftEngineType_stagedOrder, &stage.CodeAircraftEngineTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_mapString, &stage.CodeAircraftICAOType_stagedOrder, &stage.CodeAircraftICAOTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString, &stage.CodeAircraftType_stagedOrder, &stage.CodeAircraftTypeOrder)

	__gong__resetStageType(&stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString, &stage.CodeAircraftWingspanClassType_stagedOrder, &stage.CodeAircraftWingspanClassTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceActivityTypes, &stage.CodeAirspaceActivityTypes_mapString, &stage.CodeAirspaceActivityType_stagedOrder, &stage.CodeAirspaceActivityTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceAggregationTypes, &stage.CodeAirspaceAggregationTypes_mapString, &stage.CodeAirspaceAggregationType_stagedOrder, &stage.CodeAirspaceAggregationTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceClassificationTypes, &stage.CodeAirspaceClassificationTypes_mapString, &stage.CodeAirspaceClassificationType_stagedOrder, &stage.CodeAirspaceClassificationTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceDependencyTypes, &stage.CodeAirspaceDependencyTypes_mapString, &stage.CodeAirspaceDependencyType_stagedOrder, &stage.CodeAirspaceDependencyTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceDesignatorTypes, &stage.CodeAirspaceDesignatorTypes_mapString, &stage.CodeAirspaceDesignatorType_stagedOrder, &stage.CodeAirspaceDesignatorTypeOrder)

	__gong__resetStageType(&stage.CodeAirspaceTypes, &stage.CodeAirspaceTypes_mapString, &stage.CodeAirspaceType_stagedOrder, &stage.CodeAirspaceTypeOrder)

	__gong__resetStageType(&stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString, &stage.CodeAltitudeUseType_stagedOrder, &stage.CodeAltitudeUseTypeOrder)

	__gong__resetStageType(&stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_mapString, &stage.CodeAuthorityType_stagedOrder, &stage.CodeAuthorityTypeOrder)

	__gong__resetStageType(&stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString, &stage.CodeCommunicationModeType_stagedOrder, &stage.CodeCommunicationModeTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString, &stage.CodeEquipmentAntiCollisionType_stagedOrder, &stage.CodeEquipmentAntiCollisionTypeOrder)

	__gong__resetStageType(&stage.CodeGeoBorderTypes, &stage.CodeGeoBorderTypes_mapString, &stage.CodeGeoBorderType_stagedOrder, &stage.CodeGeoBorderTypeOrder)

	__gong__resetStageType(&stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString, &stage.CodeMilitaryOperationsType_stagedOrder, &stage.CodeMilitaryOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString, &stage.CodeNavigationEquipmentType_stagedOrder, &stage.CodeNavigationEquipmentTypeOrder)

	__gong__resetStageType(&stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString, &stage.CodeNavigationSpecificationType_stagedOrder, &stage.CodeNavigationSpecificationTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString, &stage.CodeRVSMType_stagedOrder, &stage.CodeRVSMTypeOrder)

	__gong__resetStageType(&stage.CodeStatusAirspaceTypes, &stage.CodeStatusAirspaceTypes_mapString, &stage.CodeStatusAirspaceType_stagedOrder, &stage.CodeStatusAirspaceTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString, &stage.CodeTransponderType_stagedOrder, &stage.CodeTransponderTypeOrder)

	__gong__resetStageType(&stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString, &stage.CodeValueInterpretationType_stagedOrder, &stage.CodeValueInterpretationTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString, &stage.CodeVerticalReferenceType_stagedOrder, &stage.CodeVerticalReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString, &stage.CodeWakeTurbulenceType_stagedOrder, &stage.CodeWakeTurbulenceTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.NoNumberTypes, &stage.NoNumberTypes_mapString, &stage.NoNumberType_stagedOrder, &stage.NoNumberTypeOrder)

	__gong__resetStageType(&stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString, &stage.NoSequenceType_stagedOrder, &stage.NoSequenceTypeOrder)

	__gong__resetStageType(&stage.TextNameTypes, &stage.TextNameTypes_mapString, &stage.TextNameType_stagedOrder, &stage.TextNameTypeOrder)

	__gong__resetStageType(&stage.TextNoteTypes, &stage.TextNoteTypes_mapString, &stage.TextNoteType_stagedOrder, &stage.TextNoteTypeOrder)

	__gong__resetStageType(&stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString, &stage.TextPropertyNameType_stagedOrder, &stage.TextPropertyNameTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValFLTypes, &stage.ValFLTypes_mapString, &stage.ValFLType_stagedOrder, &stage.ValFLTypeOrder)

	__gong__resetStageType(&stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString, &stage.ValSpeedType_stagedOrder, &stage.ValSpeedTypeOrder)

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

func (CodeAircraftType) GongGetAssociationName() any {
	return &CodeAircraftType{
	}
}

func (CodeAircraftWingspanClassType) GongGetAssociationName() any {
	return &CodeAircraftWingspanClassType{
	}
}

func (CodeAirspaceActivityType) GongGetAssociationName() any {
	return &CodeAirspaceActivityType{
	}
}

func (CodeAirspaceAggregationType) GongGetAssociationName() any {
	return &CodeAirspaceAggregationType{
	}
}

func (CodeAirspaceClassificationType) GongGetAssociationName() any {
	return &CodeAirspaceClassificationType{
	}
}

func (CodeAirspaceDependencyType) GongGetAssociationName() any {
	return &CodeAirspaceDependencyType{
	}
}

func (CodeAirspaceDesignatorType) GongGetAssociationName() any {
	return &CodeAirspaceDesignatorType{
	}
}

func (CodeAirspaceType) GongGetAssociationName() any {
	return &CodeAirspaceType{
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

func (CodeCommunicationModeType) GongGetAssociationName() any {
	return &CodeCommunicationModeType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeEquipmentAntiCollisionType) GongGetAssociationName() any {
	return &CodeEquipmentAntiCollisionType{
	}
}

func (CodeGeoBorderType) GongGetAssociationName() any {
	return &CodeGeoBorderType{
	}
}

func (CodeMilitaryOperationsType) GongGetAssociationName() any {
	return &CodeMilitaryOperationsType{
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

func (CodeRVSMType) GongGetAssociationName() any {
	return &CodeRVSMType{
	}
}

func (CodeStatusAirspaceType) GongGetAssociationName() any {
	return &CodeStatusAirspaceType{
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

func (TimeType) GongGetAssociationName() any {
	return &TimeType{
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

func (ValSpeedType) GongGetAssociationName() any {
	return &ValSpeedType{
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
	// reverse maps of direct associations of CodeAirspaceActivityType
	case CodeAirspaceActivityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceAggregationType
	case CodeAirspaceAggregationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceClassificationType
	case CodeAirspaceClassificationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceDependencyType
	case CodeAirspaceDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceDesignatorType
	case CodeAirspaceDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceType
	case CodeAirspaceType:
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
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGeoBorderType
	case CodeGeoBorderType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
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
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusAirspaceType
	case CodeStatusAirspaceType:
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
	// reverse maps of direct associations of TimeType
	case TimeType:
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
	// reverse maps of direct associations of CodeAirspaceActivityType
	case CodeAirspaceActivityType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceAggregationType
	case CodeAirspaceAggregationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceClassificationType
	case CodeAirspaceClassificationType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceDependencyType
	case CodeAirspaceDependencyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceDesignatorType
	case CodeAirspaceDesignatorType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeAirspaceType
	case CodeAirspaceType:
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
	// reverse maps of direct associations of CodeEquipmentAntiCollisionType
	case CodeEquipmentAntiCollisionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGeoBorderType
	case CodeGeoBorderType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeMilitaryOperationsType
	case CodeMilitaryOperationsType:
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
	// reverse maps of direct associations of CodeRVSMType
	case CodeRVSMType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusAirspaceType
	case CodeStatusAirspaceType:
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
	// reverse maps of direct associations of TimeType
	case TimeType:
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

func (*CodeAircraftType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAircraftWingspanClassType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceActivityType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceAggregationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceClassificationType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceDependencyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceDesignatorType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeAirspaceType) GongGetReverseFields() []GongReverseField {
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

func (*CodeCommunicationModeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeEquipmentAntiCollisionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeGeoBorderType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeMilitaryOperationsType) GongGetReverseFields() []GongReverseField {
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

func (*CodeRVSMType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusAirspaceType) GongGetReverseFields() []GongReverseField {
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

func (*TimeType) GongGetReverseFields() []GongReverseField {
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

func (*ValSpeedType) GongGetReverseFields() []GongReverseField {
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

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeairspacetype *CodeAirspaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codegeobordertype *CodeGeoBorderType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspaceactivitytype.Name
	case "NilReason":
		res.valueString = codeairspaceactivitytype.NilReason
	case "EnclosedText":
		res.valueString = codeairspaceactivitytype.EnclosedText
	}
	return
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspaceaggregationtype.Name
	case "NilReason":
		res.valueString = codeairspaceaggregationtype.NilReason
	case "EnclosedText":
		res.valueString = codeairspaceaggregationtype.EnclosedText
	}
	return
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspaceclassificationtype.Name
	case "NilReason":
		res.valueString = codeairspaceclassificationtype.NilReason
	case "EnclosedText":
		res.valueString = codeairspaceclassificationtype.EnclosedText
	}
	return
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspacedependencytype.Name
	case "NilReason":
		res.valueString = codeairspacedependencytype.NilReason
	case "EnclosedText":
		res.valueString = codeairspacedependencytype.EnclosedText
	}
	return
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspacedesignatortype.Name
	case "NilReason":
		res.valueString = codeairspacedesignatortype.NilReason
	case "EnclosedText":
		res.valueString = codeairspacedesignatortype.EnclosedText
	}
	return
}

func (codeairspacetype *CodeAirspaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeairspacetype.Name
	case "NilReason":
		res.valueString = codeairspacetype.NilReason
	case "EnclosedText":
		res.valueString = codeairspacetype.EnclosedText
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

func (codegeobordertype *CodeGeoBorderType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codegeobordertype.Name
	case "NilReason":
		res.valueString = codegeobordertype.NilReason
	case "EnclosedText":
		res.valueString = codegeobordertype.EnclosedText
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

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusairspacetype.Name
	case "NilReason":
		res.valueString = codestatusairspacetype.NilReason
	case "EnclosedText":
		res.valueString = codestatusairspacetype.EnclosedText
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

func (codeaircrafttype *CodeAircraftType) GongGetGongstructName() string {
	return "CodeAircraftType"
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetGongstructName() string {
	return "CodeAircraftWingspanClassType"
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetGongstructName() string {
	return "CodeAirspaceActivityType"
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetGongstructName() string {
	return "CodeAirspaceAggregationType"
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetGongstructName() string {
	return "CodeAirspaceClassificationType"
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetGongstructName() string {
	return "CodeAirspaceDependencyType"
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetGongstructName() string {
	return "CodeAirspaceDesignatorType"
}

func (codeairspacetype *CodeAirspaceType) GongGetGongstructName() string {
	return "CodeAirspaceType"
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetGongstructName() string {
	return "CodeAltitudeUseType"
}

func (codeauthoritytype *CodeAuthorityType) GongGetGongstructName() string {
	return "CodeAuthorityType"
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetGongstructName() string {
	return "CodeCommunicationModeType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetGongstructName() string {
	return "CodeEquipmentAntiCollisionType"
}

func (codegeobordertype *CodeGeoBorderType) GongGetGongstructName() string {
	return "CodeGeoBorderType"
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetGongstructName() string {
	return "CodeMilitaryOperationsType"
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

func (codervsmtype *CodeRVSMType) GongGetGongstructName() string {
	return "CodeRVSMType"
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetGongstructName() string {
	return "CodeStatusAirspaceType"
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

func (textnametype *TextNameType) GongGetGongstructName() string {
	return "TextNameType"
}

func (textnotetype *TextNoteType) GongGetGongstructName() string {
	return "TextNoteType"
}

func (textpropertynametype *TextPropertyNameType) GongGetGongstructName() string {
	return "TextPropertyNameType"
}

func (timetype *TimeType) GongGetGongstructName() string {
	return "TimeType"
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

func (valspeedtype *ValSpeedType) GongGetGongstructName() string {
	return "ValSpeedType"
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

	__gong__rebuildMapString(stage.CodeAircraftTypes, &stage.CodeAircraftTypes_mapString)

	__gong__rebuildMapString(stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceActivityTypes, &stage.CodeAirspaceActivityTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceAggregationTypes, &stage.CodeAirspaceAggregationTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceClassificationTypes, &stage.CodeAirspaceClassificationTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceDependencyTypes, &stage.CodeAirspaceDependencyTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceDesignatorTypes, &stage.CodeAirspaceDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.CodeAirspaceTypes, &stage.CodeAirspaceTypes_mapString)

	__gong__rebuildMapString(stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_mapString)

	__gong__rebuildMapString(stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_mapString)

	__gong__rebuildMapString(stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_mapString)

	__gong__rebuildMapString(stage.CodeGeoBorderTypes, &stage.CodeGeoBorderTypes_mapString)

	__gong__rebuildMapString(stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_mapString)

	__gong__rebuildMapString(stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeRVSMTypes, &stage.CodeRVSMTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusAirspaceTypes, &stage.CodeStatusAirspaceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeTransponderTypes, &stage.CodeTransponderTypes_mapString)

	__gong__rebuildMapString(stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.NoNumberTypes, &stage.NoNumberTypes_mapString)

	__gong__rebuildMapString(stage.NoSequenceTypes, &stage.NoSequenceTypes_mapString)

	__gong__rebuildMapString(stage.TextNameTypes, &stage.TextNameTypes_mapString)

	__gong__rebuildMapString(stage.TextNoteTypes, &stage.TextNoteTypes_mapString)

	__gong__rebuildMapString(stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValFLTypes, &stage.ValFLTypes_mapString)

	__gong__rebuildMapString(stage.ValSpeedTypes, &stage.ValSpeedTypes_mapString)

	__gong__rebuildMapString(stage.ValWeightTypes, &stage.ValWeightTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
