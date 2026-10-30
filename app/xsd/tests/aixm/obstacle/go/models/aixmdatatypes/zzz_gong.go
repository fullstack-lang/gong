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

	CodeGroundLightingTypes                map[*CodeGroundLightingType]struct{}
	CodeGroundLightingTypes_instance       map[*CodeGroundLightingType]*CodeGroundLightingType
	CodeGroundLightingTypes_mapString      map[string]*CodeGroundLightingType
	CodeGroundLightingTypeOrder            uint
	CodeGroundLightingType_stagedOrder     map[*CodeGroundLightingType]uint
	CodeGroundLightingType_orderStaged     map[uint]*CodeGroundLightingType
	CodeGroundLightingTypes_reference      map[*CodeGroundLightingType]*CodeGroundLightingType
	CodeGroundLightingTypes_referenceOrder map[*CodeGroundLightingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeGroundLightingTypeCreateCallback GongOnAfterCreateInterface[CodeGroundLightingType]
	OnAfterCodeGroundLightingTypeUpdateCallback GongOnAfterUpdateInterface[CodeGroundLightingType]
	OnAfterCodeGroundLightingTypeDeleteCallback GongOnAfterDeleteInterface[CodeGroundLightingType]

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

	CodeObstacleAreaTypes                map[*CodeObstacleAreaType]struct{}
	CodeObstacleAreaTypes_instance       map[*CodeObstacleAreaType]*CodeObstacleAreaType
	CodeObstacleAreaTypes_mapString      map[string]*CodeObstacleAreaType
	CodeObstacleAreaTypeOrder            uint
	CodeObstacleAreaType_stagedOrder     map[*CodeObstacleAreaType]uint
	CodeObstacleAreaType_orderStaged     map[uint]*CodeObstacleAreaType
	CodeObstacleAreaTypes_reference      map[*CodeObstacleAreaType]*CodeObstacleAreaType
	CodeObstacleAreaTypes_referenceOrder map[*CodeObstacleAreaType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeObstacleAreaTypeCreateCallback GongOnAfterCreateInterface[CodeObstacleAreaType]
	OnAfterCodeObstacleAreaTypeUpdateCallback GongOnAfterUpdateInterface[CodeObstacleAreaType]
	OnAfterCodeObstacleAreaTypeDeleteCallback GongOnAfterDeleteInterface[CodeObstacleAreaType]

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

	CodeStatusConstructionTypes                map[*CodeStatusConstructionType]struct{}
	CodeStatusConstructionTypes_instance       map[*CodeStatusConstructionType]*CodeStatusConstructionType
	CodeStatusConstructionTypes_mapString      map[string]*CodeStatusConstructionType
	CodeStatusConstructionTypeOrder            uint
	CodeStatusConstructionType_stagedOrder     map[*CodeStatusConstructionType]uint
	CodeStatusConstructionType_orderStaged     map[uint]*CodeStatusConstructionType
	CodeStatusConstructionTypes_reference      map[*CodeStatusConstructionType]*CodeStatusConstructionType
	CodeStatusConstructionTypes_referenceOrder map[*CodeStatusConstructionType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeStatusConstructionTypeCreateCallback GongOnAfterCreateInterface[CodeStatusConstructionType]
	OnAfterCodeStatusConstructionTypeUpdateCallback GongOnAfterUpdateInterface[CodeStatusConstructionType]
	OnAfterCodeStatusConstructionTypeDeleteCallback GongOnAfterDeleteInterface[CodeStatusConstructionType]

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

	CodeVerticalStructureMarkingTypes                map[*CodeVerticalStructureMarkingType]struct{}
	CodeVerticalStructureMarkingTypes_instance       map[*CodeVerticalStructureMarkingType]*CodeVerticalStructureMarkingType
	CodeVerticalStructureMarkingTypes_mapString      map[string]*CodeVerticalStructureMarkingType
	CodeVerticalStructureMarkingTypeOrder            uint
	CodeVerticalStructureMarkingType_stagedOrder     map[*CodeVerticalStructureMarkingType]uint
	CodeVerticalStructureMarkingType_orderStaged     map[uint]*CodeVerticalStructureMarkingType
	CodeVerticalStructureMarkingTypes_reference      map[*CodeVerticalStructureMarkingType]*CodeVerticalStructureMarkingType
	CodeVerticalStructureMarkingTypes_referenceOrder map[*CodeVerticalStructureMarkingType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVerticalStructureMarkingTypeCreateCallback GongOnAfterCreateInterface[CodeVerticalStructureMarkingType]
	OnAfterCodeVerticalStructureMarkingTypeUpdateCallback GongOnAfterUpdateInterface[CodeVerticalStructureMarkingType]
	OnAfterCodeVerticalStructureMarkingTypeDeleteCallback GongOnAfterDeleteInterface[CodeVerticalStructureMarkingType]

	CodeVerticalStructureMaterialTypes                map[*CodeVerticalStructureMaterialType]struct{}
	CodeVerticalStructureMaterialTypes_instance       map[*CodeVerticalStructureMaterialType]*CodeVerticalStructureMaterialType
	CodeVerticalStructureMaterialTypes_mapString      map[string]*CodeVerticalStructureMaterialType
	CodeVerticalStructureMaterialTypeOrder            uint
	CodeVerticalStructureMaterialType_stagedOrder     map[*CodeVerticalStructureMaterialType]uint
	CodeVerticalStructureMaterialType_orderStaged     map[uint]*CodeVerticalStructureMaterialType
	CodeVerticalStructureMaterialTypes_reference      map[*CodeVerticalStructureMaterialType]*CodeVerticalStructureMaterialType
	CodeVerticalStructureMaterialTypes_referenceOrder map[*CodeVerticalStructureMaterialType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVerticalStructureMaterialTypeCreateCallback GongOnAfterCreateInterface[CodeVerticalStructureMaterialType]
	OnAfterCodeVerticalStructureMaterialTypeUpdateCallback GongOnAfterUpdateInterface[CodeVerticalStructureMaterialType]
	OnAfterCodeVerticalStructureMaterialTypeDeleteCallback GongOnAfterDeleteInterface[CodeVerticalStructureMaterialType]

	CodeVerticalStructureTypes                map[*CodeVerticalStructureType]struct{}
	CodeVerticalStructureTypes_instance       map[*CodeVerticalStructureType]*CodeVerticalStructureType
	CodeVerticalStructureTypes_mapString      map[string]*CodeVerticalStructureType
	CodeVerticalStructureTypeOrder            uint
	CodeVerticalStructureType_stagedOrder     map[*CodeVerticalStructureType]uint
	CodeVerticalStructureType_orderStaged     map[uint]*CodeVerticalStructureType
	CodeVerticalStructureTypes_reference      map[*CodeVerticalStructureType]*CodeVerticalStructureType
	CodeVerticalStructureTypes_referenceOrder map[*CodeVerticalStructureType]uint

	// insertion point for slice of pointers maps
	OnAfterCodeVerticalStructureTypeCreateCallback GongOnAfterCreateInterface[CodeVerticalStructureType]
	OnAfterCodeVerticalStructureTypeUpdateCallback GongOnAfterUpdateInterface[CodeVerticalStructureType]
	OnAfterCodeVerticalStructureTypeDeleteCallback GongOnAfterDeleteInterface[CodeVerticalStructureType]

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
	__gong__clearReferences(&stage.CodeColourTypes_reference, &stage.CodeColourTypes_instance, &stage.CodeColourTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeDayTypes_reference, &stage.CodeDayTypes_instance, &stage.CodeDayTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeGroundLightingTypes_reference, &stage.CodeGroundLightingTypes_instance, &stage.CodeGroundLightingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_instance, &stage.CodeLightIntensityTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeLightSourceTypes_reference, &stage.CodeLightSourceTypes_instance, &stage.CodeLightSourceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_instance, &stage.CodeNotePurposeTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstacleAreaTypes_reference, &stage.CodeObstacleAreaTypes_instance, &stage.CodeObstacleAreaTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_instance, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusConstructionTypes_reference, &stage.CodeStatusConstructionTypes_instance, &stage.CodeStatusConstructionTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_instance, &stage.CodeStatusOperationsTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_instance, &stage.CodeTimeEventCombinationTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_instance, &stage.CodeTimeEventTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_instance, &stage.CodeTimeReferenceTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_instance, &stage.CodeVerticalDatumTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalStructureMarkingTypes_reference, &stage.CodeVerticalStructureMarkingTypes_instance, &stage.CodeVerticalStructureMarkingTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalStructureMaterialTypes_reference, &stage.CodeVerticalStructureMaterialTypes_instance, &stage.CodeVerticalStructureMaterialTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeVerticalStructureTypes_reference, &stage.CodeVerticalStructureTypes_instance, &stage.CodeVerticalStructureTypes_referenceOrder)

	__gong__clearReferences(&stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_instance, &stage.CodeYesNoTypes_referenceOrder)

	__gong__clearReferences(&stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_instance, &stage.DateMonthDayTypes_referenceOrder)

	__gong__clearReferences(&stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_instance, &stage.TextDesignatorTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNameTypes_reference, &stage.TextNameTypes_instance, &stage.TextNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TextNoteTypes_reference, &stage.TextNoteTypes_instance, &stage.TextNoteTypes_referenceOrder)

	__gong__clearReferences(&stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_instance, &stage.TextPropertyNameTypes_referenceOrder)

	__gong__clearReferences(&stage.TimeTypes_reference, &stage.TimeTypes_instance, &stage.TimeTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_instance, &stage.ValDistanceSignedTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_instance, &stage.ValDistanceTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_instance, &stage.ValDistanceVerticalTypes_referenceOrder)

	__gong__clearReferences(&stage.ValDurationTypes_reference, &stage.ValDurationTypes_instance, &stage.ValDurationTypes_referenceOrder)

	__gong__clearReferences(&stage.ValLightIntensityTypes_reference, &stage.ValLightIntensityTypes_instance, &stage.ValLightIntensityTypes_referenceOrder)

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
	stage.CodeColourTypeOrder = __gong__recomputeOrder(stage.CodeColourType_stagedOrder)

	stage.CodeDayTypeOrder = __gong__recomputeOrder(stage.CodeDayType_stagedOrder)

	stage.CodeGroundLightingTypeOrder = __gong__recomputeOrder(stage.CodeGroundLightingType_stagedOrder)

	stage.CodeLightIntensityTypeOrder = __gong__recomputeOrder(stage.CodeLightIntensityType_stagedOrder)

	stage.CodeLightSourceTypeOrder = __gong__recomputeOrder(stage.CodeLightSourceType_stagedOrder)

	stage.CodeNotePurposeTypeOrder = __gong__recomputeOrder(stage.CodeNotePurposeType_stagedOrder)

	stage.CodeObstacleAreaTypeOrder = __gong__recomputeOrder(stage.CodeObstacleAreaType_stagedOrder)

	stage.CodeObstacleAssessmentSurfaceTypeOrder = __gong__recomputeOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder)

	stage.CodeStatusConstructionTypeOrder = __gong__recomputeOrder(stage.CodeStatusConstructionType_stagedOrder)

	stage.CodeStatusOperationsTypeOrder = __gong__recomputeOrder(stage.CodeStatusOperationsType_stagedOrder)

	stage.CodeTimeEventCombinationTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventCombinationType_stagedOrder)

	stage.CodeTimeEventTypeOrder = __gong__recomputeOrder(stage.CodeTimeEventType_stagedOrder)

	stage.CodeTimeReferenceTypeOrder = __gong__recomputeOrder(stage.CodeTimeReferenceType_stagedOrder)

	stage.CodeVerticalDatumTypeOrder = __gong__recomputeOrder(stage.CodeVerticalDatumType_stagedOrder)

	stage.CodeVerticalStructureMarkingTypeOrder = __gong__recomputeOrder(stage.CodeVerticalStructureMarkingType_stagedOrder)

	stage.CodeVerticalStructureMaterialTypeOrder = __gong__recomputeOrder(stage.CodeVerticalStructureMaterialType_stagedOrder)

	stage.CodeVerticalStructureTypeOrder = __gong__recomputeOrder(stage.CodeVerticalStructureType_stagedOrder)

	stage.CodeYesNoTypeOrder = __gong__recomputeOrder(stage.CodeYesNoType_stagedOrder)

	stage.DateMonthDayTypeOrder = __gong__recomputeOrder(stage.DateMonthDayType_stagedOrder)

	stage.TextDesignatorTypeOrder = __gong__recomputeOrder(stage.TextDesignatorType_stagedOrder)

	stage.TextNameTypeOrder = __gong__recomputeOrder(stage.TextNameType_stagedOrder)

	stage.TextNoteTypeOrder = __gong__recomputeOrder(stage.TextNoteType_stagedOrder)

	stage.TextPropertyNameTypeOrder = __gong__recomputeOrder(stage.TextPropertyNameType_stagedOrder)

	stage.TimeTypeOrder = __gong__recomputeOrder(stage.TimeType_stagedOrder)

	stage.ValDistanceSignedTypeOrder = __gong__recomputeOrder(stage.ValDistanceSignedType_stagedOrder)

	stage.ValDistanceTypeOrder = __gong__recomputeOrder(stage.ValDistanceType_stagedOrder)

	stage.ValDistanceVerticalTypeOrder = __gong__recomputeOrder(stage.ValDistanceVerticalType_stagedOrder)

	stage.ValDurationTypeOrder = __gong__recomputeOrder(stage.ValDurationType_stagedOrder)

	stage.ValLightIntensityTypeOrder = __gong__recomputeOrder(stage.ValLightIntensityType_stagedOrder)

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

func (*CodeGroundLightingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeGroundLightingTypes, stage.CodeGroundLightingType_stagedOrder)
}

func (*CodeGroundLightingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeGroundLightingType_orderStaged[order]
}

func (*CodeGroundLightingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeGroundLightingTypes_mapString
}

func (*CodeGroundLightingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeGroundLightingTypes
}

func (*CodeGroundLightingType) GongNewInstance() any {
	return new(CodeGroundLightingType)
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

func (*CodeObstacleAreaType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaType_stagedOrder)
}

func (*CodeObstacleAreaType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeObstacleAreaType_orderStaged[order]
}

func (*CodeObstacleAreaType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeObstacleAreaTypes_mapString
}

func (*CodeObstacleAreaType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeObstacleAreaTypes
}

func (*CodeObstacleAreaType) GongNewInstance() any {
	return new(CodeObstacleAreaType)
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

func (*CodeStatusConstructionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionType_stagedOrder)
}

func (*CodeStatusConstructionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeStatusConstructionType_orderStaged[order]
}

func (*CodeStatusConstructionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeStatusConstructionTypes_mapString
}

func (*CodeStatusConstructionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeStatusConstructionTypes
}

func (*CodeStatusConstructionType) GongNewInstance() any {
	return new(CodeStatusConstructionType)
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

func (*CodeVerticalStructureMarkingType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingType_stagedOrder)
}

func (*CodeVerticalStructureMarkingType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVerticalStructureMarkingType_orderStaged[order]
}

func (*CodeVerticalStructureMarkingType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVerticalStructureMarkingTypes_mapString
}

func (*CodeVerticalStructureMarkingType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVerticalStructureMarkingTypes
}

func (*CodeVerticalStructureMarkingType) GongNewInstance() any {
	return new(CodeVerticalStructureMarkingType)
}

func (*CodeVerticalStructureMaterialType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialType_stagedOrder)
}

func (*CodeVerticalStructureMaterialType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVerticalStructureMaterialType_orderStaged[order]
}

func (*CodeVerticalStructureMaterialType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVerticalStructureMaterialTypes_mapString
}

func (*CodeVerticalStructureMaterialType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVerticalStructureMaterialTypes
}

func (*CodeVerticalStructureMaterialType) GongNewInstance() any {
	return new(CodeVerticalStructureMaterialType)
}

func (*CodeVerticalStructureType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureType_stagedOrder)
}

func (*CodeVerticalStructureType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.CodeVerticalStructureType_orderStaged[order]
}

func (*CodeVerticalStructureType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.CodeVerticalStructureTypes_mapString
}

func (*CodeVerticalStructureType) GongGetInstancesSet(stage *Stage) any {
	return &stage.CodeVerticalStructureTypes
}

func (*CodeVerticalStructureType) GongNewInstance() any {
	return new(CodeVerticalStructureType)
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
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
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
		CodeColourTypes:           make(map[*CodeColourType]struct{}),
		CodeColourTypes_mapString: make(map[string]*CodeColourType),

		CodeDayTypes:           make(map[*CodeDayType]struct{}),
		CodeDayTypes_mapString: make(map[string]*CodeDayType),

		CodeGroundLightingTypes:           make(map[*CodeGroundLightingType]struct{}),
		CodeGroundLightingTypes_mapString: make(map[string]*CodeGroundLightingType),

		CodeLightIntensityTypes:           make(map[*CodeLightIntensityType]struct{}),
		CodeLightIntensityTypes_mapString: make(map[string]*CodeLightIntensityType),

		CodeLightSourceTypes:           make(map[*CodeLightSourceType]struct{}),
		CodeLightSourceTypes_mapString: make(map[string]*CodeLightSourceType),

		CodeNotePurposeTypes:           make(map[*CodeNotePurposeType]struct{}),
		CodeNotePurposeTypes_mapString: make(map[string]*CodeNotePurposeType),

		CodeObstacleAreaTypes:           make(map[*CodeObstacleAreaType]struct{}),
		CodeObstacleAreaTypes_mapString: make(map[string]*CodeObstacleAreaType),

		CodeObstacleAssessmentSurfaceTypes:           make(map[*CodeObstacleAssessmentSurfaceType]struct{}),
		CodeObstacleAssessmentSurfaceTypes_mapString: make(map[string]*CodeObstacleAssessmentSurfaceType),

		CodeStatusConstructionTypes:           make(map[*CodeStatusConstructionType]struct{}),
		CodeStatusConstructionTypes_mapString: make(map[string]*CodeStatusConstructionType),

		CodeStatusOperationsTypes:           make(map[*CodeStatusOperationsType]struct{}),
		CodeStatusOperationsTypes_mapString: make(map[string]*CodeStatusOperationsType),

		CodeTimeEventCombinationTypes:           make(map[*CodeTimeEventCombinationType]struct{}),
		CodeTimeEventCombinationTypes_mapString: make(map[string]*CodeTimeEventCombinationType),

		CodeTimeEventTypes:           make(map[*CodeTimeEventType]struct{}),
		CodeTimeEventTypes_mapString: make(map[string]*CodeTimeEventType),

		CodeTimeReferenceTypes:           make(map[*CodeTimeReferenceType]struct{}),
		CodeTimeReferenceTypes_mapString: make(map[string]*CodeTimeReferenceType),

		CodeVerticalDatumTypes:           make(map[*CodeVerticalDatumType]struct{}),
		CodeVerticalDatumTypes_mapString: make(map[string]*CodeVerticalDatumType),

		CodeVerticalStructureMarkingTypes:           make(map[*CodeVerticalStructureMarkingType]struct{}),
		CodeVerticalStructureMarkingTypes_mapString: make(map[string]*CodeVerticalStructureMarkingType),

		CodeVerticalStructureMaterialTypes:           make(map[*CodeVerticalStructureMaterialType]struct{}),
		CodeVerticalStructureMaterialTypes_mapString: make(map[string]*CodeVerticalStructureMaterialType),

		CodeVerticalStructureTypes:           make(map[*CodeVerticalStructureType]struct{}),
		CodeVerticalStructureTypes_mapString: make(map[string]*CodeVerticalStructureType),

		CodeYesNoTypes:           make(map[*CodeYesNoType]struct{}),
		CodeYesNoTypes_mapString: make(map[string]*CodeYesNoType),

		DateMonthDayTypes:           make(map[*DateMonthDayType]struct{}),
		DateMonthDayTypes_mapString: make(map[string]*DateMonthDayType),

		TextDesignatorTypes:           make(map[*TextDesignatorType]struct{}),
		TextDesignatorTypes_mapString: make(map[string]*TextDesignatorType),

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

		ValLightIntensityTypes:           make(map[*ValLightIntensityType]struct{}),
		ValLightIntensityTypes_mapString: make(map[string]*ValLightIntensityType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CodeColourType_stagedOrder: make(map[*CodeColourType]uint),
		CodeColourType_orderStaged: make(map[uint]*CodeColourType),
		CodeColourTypes_reference:  make(map[*CodeColourType]*CodeColourType),

		CodeDayType_stagedOrder: make(map[*CodeDayType]uint),
		CodeDayType_orderStaged: make(map[uint]*CodeDayType),
		CodeDayTypes_reference:  make(map[*CodeDayType]*CodeDayType),

		CodeGroundLightingType_stagedOrder: make(map[*CodeGroundLightingType]uint),
		CodeGroundLightingType_orderStaged: make(map[uint]*CodeGroundLightingType),
		CodeGroundLightingTypes_reference:  make(map[*CodeGroundLightingType]*CodeGroundLightingType),

		CodeLightIntensityType_stagedOrder: make(map[*CodeLightIntensityType]uint),
		CodeLightIntensityType_orderStaged: make(map[uint]*CodeLightIntensityType),
		CodeLightIntensityTypes_reference:  make(map[*CodeLightIntensityType]*CodeLightIntensityType),

		CodeLightSourceType_stagedOrder: make(map[*CodeLightSourceType]uint),
		CodeLightSourceType_orderStaged: make(map[uint]*CodeLightSourceType),
		CodeLightSourceTypes_reference:  make(map[*CodeLightSourceType]*CodeLightSourceType),

		CodeNotePurposeType_stagedOrder: make(map[*CodeNotePurposeType]uint),
		CodeNotePurposeType_orderStaged: make(map[uint]*CodeNotePurposeType),
		CodeNotePurposeTypes_reference:  make(map[*CodeNotePurposeType]*CodeNotePurposeType),

		CodeObstacleAreaType_stagedOrder: make(map[*CodeObstacleAreaType]uint),
		CodeObstacleAreaType_orderStaged: make(map[uint]*CodeObstacleAreaType),
		CodeObstacleAreaTypes_reference:  make(map[*CodeObstacleAreaType]*CodeObstacleAreaType),

		CodeObstacleAssessmentSurfaceType_stagedOrder: make(map[*CodeObstacleAssessmentSurfaceType]uint),
		CodeObstacleAssessmentSurfaceType_orderStaged: make(map[uint]*CodeObstacleAssessmentSurfaceType),
		CodeObstacleAssessmentSurfaceTypes_reference:  make(map[*CodeObstacleAssessmentSurfaceType]*CodeObstacleAssessmentSurfaceType),

		CodeStatusConstructionType_stagedOrder: make(map[*CodeStatusConstructionType]uint),
		CodeStatusConstructionType_orderStaged: make(map[uint]*CodeStatusConstructionType),
		CodeStatusConstructionTypes_reference:  make(map[*CodeStatusConstructionType]*CodeStatusConstructionType),

		CodeStatusOperationsType_stagedOrder: make(map[*CodeStatusOperationsType]uint),
		CodeStatusOperationsType_orderStaged: make(map[uint]*CodeStatusOperationsType),
		CodeStatusOperationsTypes_reference:  make(map[*CodeStatusOperationsType]*CodeStatusOperationsType),

		CodeTimeEventCombinationType_stagedOrder: make(map[*CodeTimeEventCombinationType]uint),
		CodeTimeEventCombinationType_orderStaged: make(map[uint]*CodeTimeEventCombinationType),
		CodeTimeEventCombinationTypes_reference:  make(map[*CodeTimeEventCombinationType]*CodeTimeEventCombinationType),

		CodeTimeEventType_stagedOrder: make(map[*CodeTimeEventType]uint),
		CodeTimeEventType_orderStaged: make(map[uint]*CodeTimeEventType),
		CodeTimeEventTypes_reference:  make(map[*CodeTimeEventType]*CodeTimeEventType),

		CodeTimeReferenceType_stagedOrder: make(map[*CodeTimeReferenceType]uint),
		CodeTimeReferenceType_orderStaged: make(map[uint]*CodeTimeReferenceType),
		CodeTimeReferenceTypes_reference:  make(map[*CodeTimeReferenceType]*CodeTimeReferenceType),

		CodeVerticalDatumType_stagedOrder: make(map[*CodeVerticalDatumType]uint),
		CodeVerticalDatumType_orderStaged: make(map[uint]*CodeVerticalDatumType),
		CodeVerticalDatumTypes_reference:  make(map[*CodeVerticalDatumType]*CodeVerticalDatumType),

		CodeVerticalStructureMarkingType_stagedOrder: make(map[*CodeVerticalStructureMarkingType]uint),
		CodeVerticalStructureMarkingType_orderStaged: make(map[uint]*CodeVerticalStructureMarkingType),
		CodeVerticalStructureMarkingTypes_reference:  make(map[*CodeVerticalStructureMarkingType]*CodeVerticalStructureMarkingType),

		CodeVerticalStructureMaterialType_stagedOrder: make(map[*CodeVerticalStructureMaterialType]uint),
		CodeVerticalStructureMaterialType_orderStaged: make(map[uint]*CodeVerticalStructureMaterialType),
		CodeVerticalStructureMaterialTypes_reference:  make(map[*CodeVerticalStructureMaterialType]*CodeVerticalStructureMaterialType),

		CodeVerticalStructureType_stagedOrder: make(map[*CodeVerticalStructureType]uint),
		CodeVerticalStructureType_orderStaged: make(map[uint]*CodeVerticalStructureType),
		CodeVerticalStructureTypes_reference:  make(map[*CodeVerticalStructureType]*CodeVerticalStructureType),

		CodeYesNoType_stagedOrder: make(map[*CodeYesNoType]uint),
		CodeYesNoType_orderStaged: make(map[uint]*CodeYesNoType),
		CodeYesNoTypes_reference:  make(map[*CodeYesNoType]*CodeYesNoType),

		DateMonthDayType_stagedOrder: make(map[*DateMonthDayType]uint),
		DateMonthDayType_orderStaged: make(map[uint]*DateMonthDayType),
		DateMonthDayTypes_reference:  make(map[*DateMonthDayType]*DateMonthDayType),

		TextDesignatorType_stagedOrder: make(map[*TextDesignatorType]uint),
		TextDesignatorType_orderStaged: make(map[uint]*TextDesignatorType),
		TextDesignatorTypes_reference:  make(map[*TextDesignatorType]*TextDesignatorType),

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

		ValLightIntensityType_stagedOrder: make(map[*ValLightIntensityType]uint),
		ValLightIntensityType_orderStaged: make(map[uint]*ValLightIntensityType),
		ValLightIntensityTypes_reference:  make(map[*ValLightIntensityType]*ValLightIntensityType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"CodeColourType": &CodeColourTypeUnmarshaller{},

			"CodeDayType": &CodeDayTypeUnmarshaller{},

			"CodeGroundLightingType": &CodeGroundLightingTypeUnmarshaller{},

			"CodeLightIntensityType": &CodeLightIntensityTypeUnmarshaller{},

			"CodeLightSourceType": &CodeLightSourceTypeUnmarshaller{},

			"CodeNotePurposeType": &CodeNotePurposeTypeUnmarshaller{},

			"CodeObstacleAreaType": &CodeObstacleAreaTypeUnmarshaller{},

			"CodeObstacleAssessmentSurfaceType": &CodeObstacleAssessmentSurfaceTypeUnmarshaller{},

			"CodeStatusConstructionType": &CodeStatusConstructionTypeUnmarshaller{},

			"CodeStatusOperationsType": &CodeStatusOperationsTypeUnmarshaller{},

			"CodeTimeEventCombinationType": &CodeTimeEventCombinationTypeUnmarshaller{},

			"CodeTimeEventType": &CodeTimeEventTypeUnmarshaller{},

			"CodeTimeReferenceType": &CodeTimeReferenceTypeUnmarshaller{},

			"CodeVerticalDatumType": &CodeVerticalDatumTypeUnmarshaller{},

			"CodeVerticalStructureMarkingType": &CodeVerticalStructureMarkingTypeUnmarshaller{},

			"CodeVerticalStructureMaterialType": &CodeVerticalStructureMaterialTypeUnmarshaller{},

			"CodeVerticalStructureType": &CodeVerticalStructureTypeUnmarshaller{},

			"CodeYesNoType": &CodeYesNoTypeUnmarshaller{},

			"DateMonthDayType": &DateMonthDayTypeUnmarshaller{},

			"TextDesignatorType": &TextDesignatorTypeUnmarshaller{},

			"TextNameType": &TextNameTypeUnmarshaller{},

			"TextNoteType": &TextNoteTypeUnmarshaller{},

			"TextPropertyNameType": &TextPropertyNameTypeUnmarshaller{},

			"TimeType": &TimeTypeUnmarshaller{},

			"ValDistanceSignedType": &ValDistanceSignedTypeUnmarshaller{},

			"ValDistanceType": &ValDistanceTypeUnmarshaller{},

			"ValDistanceVerticalType": &ValDistanceVerticalTypeUnmarshaller{},

			"ValDurationType": &ValDurationTypeUnmarshaller{},

			"ValLightIntensityType": &ValLightIntensityTypeUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["CodeColourType"] = len(stage.CodeColourTypes)
	stage.Map_GongStructName_InstancesNb["CodeDayType"] = len(stage.CodeDayTypes)
	stage.Map_GongStructName_InstancesNb["CodeGroundLightingType"] = len(stage.CodeGroundLightingTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightIntensityType"] = len(stage.CodeLightIntensityTypes)
	stage.Map_GongStructName_InstancesNb["CodeLightSourceType"] = len(stage.CodeLightSourceTypes)
	stage.Map_GongStructName_InstancesNb["CodeNotePurposeType"] = len(stage.CodeNotePurposeTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstacleAreaType"] = len(stage.CodeObstacleAreaTypes)
	stage.Map_GongStructName_InstancesNb["CodeObstacleAssessmentSurfaceType"] = len(stage.CodeObstacleAssessmentSurfaceTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusConstructionType"] = len(stage.CodeStatusConstructionTypes)
	stage.Map_GongStructName_InstancesNb["CodeStatusOperationsType"] = len(stage.CodeStatusOperationsTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventCombinationType"] = len(stage.CodeTimeEventCombinationTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeEventType"] = len(stage.CodeTimeEventTypes)
	stage.Map_GongStructName_InstancesNb["CodeTimeReferenceType"] = len(stage.CodeTimeReferenceTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalDatumType"] = len(stage.CodeVerticalDatumTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalStructureMarkingType"] = len(stage.CodeVerticalStructureMarkingTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalStructureMaterialType"] = len(stage.CodeVerticalStructureMaterialTypes)
	stage.Map_GongStructName_InstancesNb["CodeVerticalStructureType"] = len(stage.CodeVerticalStructureTypes)
	stage.Map_GongStructName_InstancesNb["CodeYesNoType"] = len(stage.CodeYesNoTypes)
	stage.Map_GongStructName_InstancesNb["DateMonthDayType"] = len(stage.DateMonthDayTypes)
	stage.Map_GongStructName_InstancesNb["TextDesignatorType"] = len(stage.TextDesignatorTypes)
	stage.Map_GongStructName_InstancesNb["TextNameType"] = len(stage.TextNameTypes)
	stage.Map_GongStructName_InstancesNb["TextNoteType"] = len(stage.TextNoteTypes)
	stage.Map_GongStructName_InstancesNb["TextPropertyNameType"] = len(stage.TextPropertyNameTypes)
	stage.Map_GongStructName_InstancesNb["TimeType"] = len(stage.TimeTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceSignedType"] = len(stage.ValDistanceSignedTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceType"] = len(stage.ValDistanceTypes)
	stage.Map_GongStructName_InstancesNb["ValDistanceVerticalType"] = len(stage.ValDistanceVerticalTypes)
	stage.Map_GongStructName_InstancesNb["ValDurationType"] = len(stage.ValDurationTypes)
	stage.Map_GongStructName_InstancesNb["ValLightIntensityType"] = len(stage.ValLightIntensityTypes)
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

// Stage puts codegroundlightingtype to the model stage
func (codegroundlightingtype *CodeGroundLightingType) Stage(stage *Stage) *CodeGroundLightingType {
	__gong__stage(stage.CodeGroundLightingTypes, stage.CodeGroundLightingType_stagedOrder, stage.CodeGroundLightingType_orderStaged, &stage.CodeGroundLightingTypeOrder, stage.CodeGroundLightingTypes_mapString, codegroundlightingtype, codegroundlightingtype.Name)
	return codegroundlightingtype
}

// StagePreserveOrder puts codegroundlightingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeGroundLightingTypeOrder
// - update stage.CodeGroundLightingTypeOrder accordingly
func (codegroundlightingtype *CodeGroundLightingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeGroundLightingTypes, stage.CodeGroundLightingType_stagedOrder, stage.CodeGroundLightingType_orderStaged, &stage.CodeGroundLightingTypeOrder, stage.CodeGroundLightingTypes_mapString, codegroundlightingtype, order, codegroundlightingtype.Name)
}

// Unstage removes codegroundlightingtype off the model stage
func (codegroundlightingtype *CodeGroundLightingType) Unstage(stage *Stage) *CodeGroundLightingType {
	__gong__unstage(stage.CodeGroundLightingTypes, stage.CodeGroundLightingTypes_mapString, codegroundlightingtype, codegroundlightingtype.Name)
	return codegroundlightingtype
}

// UnstageVoid removes codegroundlightingtype off the model stage
func (codegroundlightingtype *CodeGroundLightingType) UnstageVoid(stage *Stage) {
	codegroundlightingtype.Unstage(stage)
}

func (codegroundlightingtype *CodeGroundLightingType) StageVoid(stage *Stage) {
	codegroundlightingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codegroundlightingtype *CodeGroundLightingType) GetName() (res string) {
	return codegroundlightingtype.Name
}

// for satisfaction of GongStruct interface
func (codegroundlightingtype *CodeGroundLightingType) SetName(name string) {
	codegroundlightingtype.Name = name
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

// Stage puts codeobstacleareatype to the model stage
func (codeobstacleareatype *CodeObstacleAreaType) Stage(stage *Stage) *CodeObstacleAreaType {
	__gong__stage(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaType_stagedOrder, stage.CodeObstacleAreaType_orderStaged, &stage.CodeObstacleAreaTypeOrder, stage.CodeObstacleAreaTypes_mapString, codeobstacleareatype, codeobstacleareatype.Name)
	return codeobstacleareatype
}

// StagePreserveOrder puts codeobstacleareatype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeObstacleAreaTypeOrder
// - update stage.CodeObstacleAreaTypeOrder accordingly
func (codeobstacleareatype *CodeObstacleAreaType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaType_stagedOrder, stage.CodeObstacleAreaType_orderStaged, &stage.CodeObstacleAreaTypeOrder, stage.CodeObstacleAreaTypes_mapString, codeobstacleareatype, order, codeobstacleareatype.Name)
}

// Unstage removes codeobstacleareatype off the model stage
func (codeobstacleareatype *CodeObstacleAreaType) Unstage(stage *Stage) *CodeObstacleAreaType {
	__gong__unstage(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaTypes_mapString, codeobstacleareatype, codeobstacleareatype.Name)
	return codeobstacleareatype
}

// UnstageVoid removes codeobstacleareatype off the model stage
func (codeobstacleareatype *CodeObstacleAreaType) UnstageVoid(stage *Stage) {
	codeobstacleareatype.Unstage(stage)
}

func (codeobstacleareatype *CodeObstacleAreaType) StageVoid(stage *Stage) {
	codeobstacleareatype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeobstacleareatype *CodeObstacleAreaType) GetName() (res string) {
	return codeobstacleareatype.Name
}

// for satisfaction of GongStruct interface
func (codeobstacleareatype *CodeObstacleAreaType) SetName(name string) {
	codeobstacleareatype.Name = name
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

// Stage puts codestatusconstructiontype to the model stage
func (codestatusconstructiontype *CodeStatusConstructionType) Stage(stage *Stage) *CodeStatusConstructionType {
	__gong__stage(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionType_stagedOrder, stage.CodeStatusConstructionType_orderStaged, &stage.CodeStatusConstructionTypeOrder, stage.CodeStatusConstructionTypes_mapString, codestatusconstructiontype, codestatusconstructiontype.Name)
	return codestatusconstructiontype
}

// StagePreserveOrder puts codestatusconstructiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeStatusConstructionTypeOrder
// - update stage.CodeStatusConstructionTypeOrder accordingly
func (codestatusconstructiontype *CodeStatusConstructionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionType_stagedOrder, stage.CodeStatusConstructionType_orderStaged, &stage.CodeStatusConstructionTypeOrder, stage.CodeStatusConstructionTypes_mapString, codestatusconstructiontype, order, codestatusconstructiontype.Name)
}

// Unstage removes codestatusconstructiontype off the model stage
func (codestatusconstructiontype *CodeStatusConstructionType) Unstage(stage *Stage) *CodeStatusConstructionType {
	__gong__unstage(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionTypes_mapString, codestatusconstructiontype, codestatusconstructiontype.Name)
	return codestatusconstructiontype
}

// UnstageVoid removes codestatusconstructiontype off the model stage
func (codestatusconstructiontype *CodeStatusConstructionType) UnstageVoid(stage *Stage) {
	codestatusconstructiontype.Unstage(stage)
}

func (codestatusconstructiontype *CodeStatusConstructionType) StageVoid(stage *Stage) {
	codestatusconstructiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codestatusconstructiontype *CodeStatusConstructionType) GetName() (res string) {
	return codestatusconstructiontype.Name
}

// for satisfaction of GongStruct interface
func (codestatusconstructiontype *CodeStatusConstructionType) SetName(name string) {
	codestatusconstructiontype.Name = name
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

// Stage puts codeverticalstructuremarkingtype to the model stage
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) Stage(stage *Stage) *CodeVerticalStructureMarkingType {
	__gong__stage(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingType_stagedOrder, stage.CodeVerticalStructureMarkingType_orderStaged, &stage.CodeVerticalStructureMarkingTypeOrder, stage.CodeVerticalStructureMarkingTypes_mapString, codeverticalstructuremarkingtype, codeverticalstructuremarkingtype.Name)
	return codeverticalstructuremarkingtype
}

// StagePreserveOrder puts codeverticalstructuremarkingtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVerticalStructureMarkingTypeOrder
// - update stage.CodeVerticalStructureMarkingTypeOrder accordingly
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingType_stagedOrder, stage.CodeVerticalStructureMarkingType_orderStaged, &stage.CodeVerticalStructureMarkingTypeOrder, stage.CodeVerticalStructureMarkingTypes_mapString, codeverticalstructuremarkingtype, order, codeverticalstructuremarkingtype.Name)
}

// Unstage removes codeverticalstructuremarkingtype off the model stage
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) Unstage(stage *Stage) *CodeVerticalStructureMarkingType {
	__gong__unstage(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingTypes_mapString, codeverticalstructuremarkingtype, codeverticalstructuremarkingtype.Name)
	return codeverticalstructuremarkingtype
}

// UnstageVoid removes codeverticalstructuremarkingtype off the model stage
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) UnstageVoid(stage *Stage) {
	codeverticalstructuremarkingtype.Unstage(stage)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) StageVoid(stage *Stage) {
	codeverticalstructuremarkingtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GetName() (res string) {
	return codeverticalstructuremarkingtype.Name
}

// for satisfaction of GongStruct interface
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) SetName(name string) {
	codeverticalstructuremarkingtype.Name = name
}

// Stage puts codeverticalstructurematerialtype to the model stage
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) Stage(stage *Stage) *CodeVerticalStructureMaterialType {
	__gong__stage(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialType_stagedOrder, stage.CodeVerticalStructureMaterialType_orderStaged, &stage.CodeVerticalStructureMaterialTypeOrder, stage.CodeVerticalStructureMaterialTypes_mapString, codeverticalstructurematerialtype, codeverticalstructurematerialtype.Name)
	return codeverticalstructurematerialtype
}

// StagePreserveOrder puts codeverticalstructurematerialtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVerticalStructureMaterialTypeOrder
// - update stage.CodeVerticalStructureMaterialTypeOrder accordingly
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialType_stagedOrder, stage.CodeVerticalStructureMaterialType_orderStaged, &stage.CodeVerticalStructureMaterialTypeOrder, stage.CodeVerticalStructureMaterialTypes_mapString, codeverticalstructurematerialtype, order, codeverticalstructurematerialtype.Name)
}

// Unstage removes codeverticalstructurematerialtype off the model stage
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) Unstage(stage *Stage) *CodeVerticalStructureMaterialType {
	__gong__unstage(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialTypes_mapString, codeverticalstructurematerialtype, codeverticalstructurematerialtype.Name)
	return codeverticalstructurematerialtype
}

// UnstageVoid removes codeverticalstructurematerialtype off the model stage
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) UnstageVoid(stage *Stage) {
	codeverticalstructurematerialtype.Unstage(stage)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) StageVoid(stage *Stage) {
	codeverticalstructurematerialtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GetName() (res string) {
	return codeverticalstructurematerialtype.Name
}

// for satisfaction of GongStruct interface
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) SetName(name string) {
	codeverticalstructurematerialtype.Name = name
}

// Stage puts codeverticalstructuretype to the model stage
func (codeverticalstructuretype *CodeVerticalStructureType) Stage(stage *Stage) *CodeVerticalStructureType {
	__gong__stage(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureType_stagedOrder, stage.CodeVerticalStructureType_orderStaged, &stage.CodeVerticalStructureTypeOrder, stage.CodeVerticalStructureTypes_mapString, codeverticalstructuretype, codeverticalstructuretype.Name)
	return codeverticalstructuretype
}

// StagePreserveOrder puts codeverticalstructuretype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodeVerticalStructureTypeOrder
// - update stage.CodeVerticalStructureTypeOrder accordingly
func (codeverticalstructuretype *CodeVerticalStructureType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureType_stagedOrder, stage.CodeVerticalStructureType_orderStaged, &stage.CodeVerticalStructureTypeOrder, stage.CodeVerticalStructureTypes_mapString, codeverticalstructuretype, order, codeverticalstructuretype.Name)
}

// Unstage removes codeverticalstructuretype off the model stage
func (codeverticalstructuretype *CodeVerticalStructureType) Unstage(stage *Stage) *CodeVerticalStructureType {
	__gong__unstage(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureTypes_mapString, codeverticalstructuretype, codeverticalstructuretype.Name)
	return codeverticalstructuretype
}

// UnstageVoid removes codeverticalstructuretype off the model stage
func (codeverticalstructuretype *CodeVerticalStructureType) UnstageVoid(stage *Stage) {
	codeverticalstructuretype.Unstage(stage)
}

func (codeverticalstructuretype *CodeVerticalStructureType) StageVoid(stage *Stage) {
	codeverticalstructuretype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (codeverticalstructuretype *CodeVerticalStructureType) GetName() (res string) {
	return codeverticalstructuretype.Name
}

// for satisfaction of GongStruct interface
func (codeverticalstructuretype *CodeVerticalStructureType) SetName(name string) {
	codeverticalstructuretype.Name = name
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

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.CodeColourTypes, &stage.CodeColourTypes_mapString, &stage.CodeColourType_stagedOrder, &stage.CodeColourTypeOrder)

	__gong__resetStageType(&stage.CodeDayTypes, &stage.CodeDayTypes_mapString, &stage.CodeDayType_stagedOrder, &stage.CodeDayTypeOrder)

	__gong__resetStageType(&stage.CodeGroundLightingTypes, &stage.CodeGroundLightingTypes_mapString, &stage.CodeGroundLightingType_stagedOrder, &stage.CodeGroundLightingTypeOrder)

	__gong__resetStageType(&stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString, &stage.CodeLightIntensityType_stagedOrder, &stage.CodeLightIntensityTypeOrder)

	__gong__resetStageType(&stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_mapString, &stage.CodeLightSourceType_stagedOrder, &stage.CodeLightSourceTypeOrder)

	__gong__resetStageType(&stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString, &stage.CodeNotePurposeType_stagedOrder, &stage.CodeNotePurposeTypeOrder)

	__gong__resetStageType(&stage.CodeObstacleAreaTypes, &stage.CodeObstacleAreaTypes_mapString, &stage.CodeObstacleAreaType_stagedOrder, &stage.CodeObstacleAreaTypeOrder)

	__gong__resetStageType(&stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString, &stage.CodeObstacleAssessmentSurfaceType_stagedOrder, &stage.CodeObstacleAssessmentSurfaceTypeOrder)

	__gong__resetStageType(&stage.CodeStatusConstructionTypes, &stage.CodeStatusConstructionTypes_mapString, &stage.CodeStatusConstructionType_stagedOrder, &stage.CodeStatusConstructionTypeOrder)

	__gong__resetStageType(&stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString, &stage.CodeStatusOperationsType_stagedOrder, &stage.CodeStatusOperationsTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString, &stage.CodeTimeEventCombinationType_stagedOrder, &stage.CodeTimeEventCombinationTypeOrder)

	__gong__resetStageType(&stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString, &stage.CodeTimeEventType_stagedOrder, &stage.CodeTimeEventTypeOrder)

	__gong__resetStageType(&stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString, &stage.CodeTimeReferenceType_stagedOrder, &stage.CodeTimeReferenceTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString, &stage.CodeVerticalDatumType_stagedOrder, &stage.CodeVerticalDatumTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalStructureMarkingTypes, &stage.CodeVerticalStructureMarkingTypes_mapString, &stage.CodeVerticalStructureMarkingType_stagedOrder, &stage.CodeVerticalStructureMarkingTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalStructureMaterialTypes, &stage.CodeVerticalStructureMaterialTypes_mapString, &stage.CodeVerticalStructureMaterialType_stagedOrder, &stage.CodeVerticalStructureMaterialTypeOrder)

	__gong__resetStageType(&stage.CodeVerticalStructureTypes, &stage.CodeVerticalStructureTypes_mapString, &stage.CodeVerticalStructureType_stagedOrder, &stage.CodeVerticalStructureTypeOrder)

	__gong__resetStageType(&stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString, &stage.CodeYesNoType_stagedOrder, &stage.CodeYesNoTypeOrder)

	__gong__resetStageType(&stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString, &stage.DateMonthDayType_stagedOrder, &stage.DateMonthDayTypeOrder)

	__gong__resetStageType(&stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString, &stage.TextDesignatorType_stagedOrder, &stage.TextDesignatorTypeOrder)

	__gong__resetStageType(&stage.TextNameTypes, &stage.TextNameTypes_mapString, &stage.TextNameType_stagedOrder, &stage.TextNameTypeOrder)

	__gong__resetStageType(&stage.TextNoteTypes, &stage.TextNoteTypes_mapString, &stage.TextNoteType_stagedOrder, &stage.TextNoteTypeOrder)

	__gong__resetStageType(&stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString, &stage.TextPropertyNameType_stagedOrder, &stage.TextPropertyNameTypeOrder)

	__gong__resetStageType(&stage.TimeTypes, &stage.TimeTypes_mapString, &stage.TimeType_stagedOrder, &stage.TimeTypeOrder)

	__gong__resetStageType(&stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString, &stage.ValDistanceSignedType_stagedOrder, &stage.ValDistanceSignedTypeOrder)

	__gong__resetStageType(&stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString, &stage.ValDistanceType_stagedOrder, &stage.ValDistanceTypeOrder)

	__gong__resetStageType(&stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString, &stage.ValDistanceVerticalType_stagedOrder, &stage.ValDistanceVerticalTypeOrder)

	__gong__resetStageType(&stage.ValDurationTypes, &stage.ValDurationTypes_mapString, &stage.ValDurationType_stagedOrder, &stage.ValDurationTypeOrder)

	__gong__resetStageType(&stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_mapString, &stage.ValLightIntensityType_stagedOrder, &stage.ValLightIntensityTypeOrder)

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
func (CodeColourType) GongGetAssociationName() any {
	return &CodeColourType{
	}
}

func (CodeDayType) GongGetAssociationName() any {
	return &CodeDayType{
	}
}

func (CodeGroundLightingType) GongGetAssociationName() any {
	return &CodeGroundLightingType{
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

func (CodeNotePurposeType) GongGetAssociationName() any {
	return &CodeNotePurposeType{
	}
}

func (CodeObstacleAreaType) GongGetAssociationName() any {
	return &CodeObstacleAreaType{
	}
}

func (CodeObstacleAssessmentSurfaceType) GongGetAssociationName() any {
	return &CodeObstacleAssessmentSurfaceType{
	}
}

func (CodeStatusConstructionType) GongGetAssociationName() any {
	return &CodeStatusConstructionType{
	}
}

func (CodeStatusOperationsType) GongGetAssociationName() any {
	return &CodeStatusOperationsType{
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

func (CodeVerticalDatumType) GongGetAssociationName() any {
	return &CodeVerticalDatumType{
	}
}

func (CodeVerticalStructureMarkingType) GongGetAssociationName() any {
	return &CodeVerticalStructureMarkingType{
	}
}

func (CodeVerticalStructureMaterialType) GongGetAssociationName() any {
	return &CodeVerticalStructureMaterialType{
	}
}

func (CodeVerticalStructureType) GongGetAssociationName() any {
	return &CodeVerticalStructureType{
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

func (ValLightIntensityType) GongGetAssociationName() any {
	return &ValLightIntensityType{
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
	// reverse maps of direct associations of CodeColourType
	case CodeColourType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGroundLightingType
	case CodeGroundLightingType:
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
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAreaType
	case CodeObstacleAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAssessmentSurfaceType
	case CodeObstacleAssessmentSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusConstructionType
	case CodeStatusConstructionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
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
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureMarkingType
	case CodeVerticalStructureMarkingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureMaterialType
	case CodeVerticalStructureMaterialType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureType
	case CodeVerticalStructureType:
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
	// reverse maps of direct associations of ValLightIntensityType
	case ValLightIntensityType:
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
	// reverse maps of direct associations of CodeColourType
	case CodeColourType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeDayType
	case CodeDayType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeGroundLightingType
	case CodeGroundLightingType:
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
	// reverse maps of direct associations of CodeNotePurposeType
	case CodeNotePurposeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAreaType
	case CodeObstacleAreaType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeObstacleAssessmentSurfaceType
	case CodeObstacleAssessmentSurfaceType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusConstructionType
	case CodeStatusConstructionType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeStatusOperationsType
	case CodeStatusOperationsType:
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
	// reverse maps of direct associations of CodeVerticalDatumType
	case CodeVerticalDatumType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureMarkingType
	case CodeVerticalStructureMarkingType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureMaterialType
	case CodeVerticalStructureMaterialType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CodeVerticalStructureType
	case CodeVerticalStructureType:
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
	// reverse maps of direct associations of ValLightIntensityType
	case ValLightIntensityType:
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
func (*CodeColourType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeDayType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeGroundLightingType) GongGetReverseFields() []GongReverseField {
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

func (*CodeNotePurposeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeObstacleAreaType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeObstacleAssessmentSurfaceType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusConstructionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeStatusOperationsType) GongGetReverseFields() []GongReverseField {
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

func (*CodeVerticalDatumType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalStructureMarkingType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalStructureMaterialType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*CodeVerticalStructureType) GongGetReverseFields() []GongReverseField {
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

func (*ValLightIntensityType) GongGetReverseFields() []GongReverseField {
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

func (codegroundlightingtype *CodeGroundLightingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeobstacleareatype *CodeObstacleAreaType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
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

func (codegroundlightingtype *CodeGroundLightingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codegroundlightingtype.Name
	case "NilReason":
		res.valueString = codegroundlightingtype.NilReason
	case "EnclosedText":
		res.valueString = codegroundlightingtype.EnclosedText
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

func (codeobstacleareatype *CodeObstacleAreaType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeobstacleareatype.Name
	case "NilReason":
		res.valueString = codeobstacleareatype.NilReason
	case "EnclosedText":
		res.valueString = codeobstacleareatype.EnclosedText
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

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codestatusconstructiontype.Name
	case "NilReason":
		res.valueString = codestatusconstructiontype.NilReason
	case "EnclosedText":
		res.valueString = codestatusconstructiontype.EnclosedText
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

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeverticalstructuremarkingtype.Name
	case "NilReason":
		res.valueString = codeverticalstructuremarkingtype.NilReason
	case "EnclosedText":
		res.valueString = codeverticalstructuremarkingtype.EnclosedText
	}
	return
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeverticalstructurematerialtype.Name
	case "NilReason":
		res.valueString = codeverticalstructurematerialtype.NilReason
	case "EnclosedText":
		res.valueString = codeverticalstructurematerialtype.EnclosedText
	}
	return
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = codeverticalstructuretype.Name
	case "NilReason":
		res.valueString = codeverticalstructuretype.NilReason
	case "EnclosedText":
		res.valueString = codeverticalstructuretype.EnclosedText
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (codecolourtype *CodeColourType) GongGetGongstructName() string {
	return "CodeColourType"
}

func (codedaytype *CodeDayType) GongGetGongstructName() string {
	return "CodeDayType"
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetGongstructName() string {
	return "CodeGroundLightingType"
}

func (codelightintensitytype *CodeLightIntensityType) GongGetGongstructName() string {
	return "CodeLightIntensityType"
}

func (codelightsourcetype *CodeLightSourceType) GongGetGongstructName() string {
	return "CodeLightSourceType"
}

func (codenotepurposetype *CodeNotePurposeType) GongGetGongstructName() string {
	return "CodeNotePurposeType"
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetGongstructName() string {
	return "CodeObstacleAreaType"
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetGongstructName() string {
	return "CodeObstacleAssessmentSurfaceType"
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetGongstructName() string {
	return "CodeStatusConstructionType"
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetGongstructName() string {
	return "CodeStatusOperationsType"
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

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetGongstructName() string {
	return "CodeVerticalDatumType"
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetGongstructName() string {
	return "CodeVerticalStructureMarkingType"
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetGongstructName() string {
	return "CodeVerticalStructureMaterialType"
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetGongstructName() string {
	return "CodeVerticalStructureType"
}

func (codeyesnotype *CodeYesNoType) GongGetGongstructName() string {
	return "CodeYesNoType"
}

func (datemonthdaytype *DateMonthDayType) GongGetGongstructName() string {
	return "DateMonthDayType"
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

func (vallightintensitytype *ValLightIntensityType) GongGetGongstructName() string {
	return "ValLightIntensityType"
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
	__gong__rebuildMapString(stage.CodeColourTypes, &stage.CodeColourTypes_mapString)

	__gong__rebuildMapString(stage.CodeDayTypes, &stage.CodeDayTypes_mapString)

	__gong__rebuildMapString(stage.CodeGroundLightingTypes, &stage.CodeGroundLightingTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_mapString)

	__gong__rebuildMapString(stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_mapString)

	__gong__rebuildMapString(stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstacleAreaTypes, &stage.CodeObstacleAreaTypes_mapString)

	__gong__rebuildMapString(stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusConstructionTypes, &stage.CodeStatusConstructionTypes_mapString)

	__gong__rebuildMapString(stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_mapString)

	__gong__rebuildMapString(stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalStructureMarkingTypes, &stage.CodeVerticalStructureMarkingTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalStructureMaterialTypes, &stage.CodeVerticalStructureMaterialTypes_mapString)

	__gong__rebuildMapString(stage.CodeVerticalStructureTypes, &stage.CodeVerticalStructureTypes_mapString)

	__gong__rebuildMapString(stage.CodeYesNoTypes, &stage.CodeYesNoTypes_mapString)

	__gong__rebuildMapString(stage.DateMonthDayTypes, &stage.DateMonthDayTypes_mapString)

	__gong__rebuildMapString(stage.TextDesignatorTypes, &stage.TextDesignatorTypes_mapString)

	__gong__rebuildMapString(stage.TextNameTypes, &stage.TextNameTypes_mapString)

	__gong__rebuildMapString(stage.TextNoteTypes, &stage.TextNoteTypes_mapString)

	__gong__rebuildMapString(stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_mapString)

	__gong__rebuildMapString(stage.TimeTypes, &stage.TimeTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceTypes, &stage.ValDistanceTypes_mapString)

	__gong__rebuildMapString(stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_mapString)

	__gong__rebuildMapString(stage.ValDurationTypes, &stage.ValDurationTypes_mapString)

	__gong__rebuildMapString(stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
