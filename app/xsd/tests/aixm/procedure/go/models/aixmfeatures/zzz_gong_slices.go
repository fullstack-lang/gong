// generated code - do not edit
package aixmfeatures

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AerialRefuellingPointType
	// insertion point per field
	stage.AerialRefuellingPointType_FacilityMakeup_reverseMap = make(map[*PointReferencePropertyType]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _pointreferencepropertytype := range aerialrefuellingpointtype.FacilityMakeup {
			stage.AerialRefuellingPointType_FacilityMakeup_reverseMap[_pointreferencepropertytype] = aerialrefuellingpointtype
		}
	}
	stage.AerialRefuellingPointType_Annotation_reverseMap = make(map[*NotePropertyType]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _notepropertytype := range aerialrefuellingpointtype.Annotation {
			stage.AerialRefuellingPointType_Annotation_reverseMap[_notepropertytype] = aerialrefuellingpointtype
		}
	}
	stage.AerialRefuellingPointType_Extension_reverseMap = make(map[*A_extension]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _a_extension := range aerialrefuellingpointtype.Extension {
			stage.AerialRefuellingPointType_Extension_reverseMap[_a_extension] = aerialrefuellingpointtype
		}
	}

	// Compute reverse map for named struct AircraftCharacteristicType
	// insertion point per field
	stage.AircraftCharacteristicType_Annotation_reverseMap = make(map[*NotePropertyType]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
			stage.AircraftCharacteristicType_Annotation_reverseMap[_notepropertytype] = aircraftcharacteristictype
		}
	}
	stage.AircraftCharacteristicType_Extension_reverseMap = make(map[*A_extension_51]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _a_extension_51 := range aircraftcharacteristictype.Extension {
			stage.AircraftCharacteristicType_Extension_reverseMap[_a_extension_51] = aircraftcharacteristictype
		}
	}

	// Compute reverse map for named struct AltitudeAdjustmentType
	// insertion point per field
	stage.AltitudeAdjustmentType_Annotation_reverseMap = make(map[*NotePropertyType]*AltitudeAdjustmentType)
	for altitudeadjustmenttype := range stage.AltitudeAdjustmentTypes {
		_ = altitudeadjustmenttype
		for _, _notepropertytype := range altitudeadjustmenttype.Annotation {
			stage.AltitudeAdjustmentType_Annotation_reverseMap[_notepropertytype] = altitudeadjustmenttype
		}
	}
	stage.AltitudeAdjustmentType_Extension_reverseMap = make(map[*A_extension_54]*AltitudeAdjustmentType)
	for altitudeadjustmenttype := range stage.AltitudeAdjustmentTypes {
		_ = altitudeadjustmenttype
		for _, _a_extension_54 := range altitudeadjustmenttype.Extension {
			stage.AltitudeAdjustmentType_Extension_reverseMap[_a_extension_54] = altitudeadjustmenttype
		}
	}

	// Compute reverse map for named struct AngleUseType
	// insertion point per field
	stage.AngleUseType_Annotation_reverseMap = make(map[*NotePropertyType]*AngleUseType)
	for angleusetype := range stage.AngleUseTypes {
		_ = angleusetype
		for _, _notepropertytype := range angleusetype.Annotation {
			stage.AngleUseType_Annotation_reverseMap[_notepropertytype] = angleusetype
		}
	}
	stage.AngleUseType_Extension_reverseMap = make(map[*A_extension_8]*AngleUseType)
	for angleusetype := range stage.AngleUseTypes {
		_ = angleusetype
		for _, _a_extension_8 := range angleusetype.Extension {
			stage.AngleUseType_Extension_reverseMap[_a_extension_8] = angleusetype
		}
	}

	// Compute reverse map for named struct ApproachAltitudeTableType
	// insertion point per field
	stage.ApproachAltitudeTableType_Annotation_reverseMap = make(map[*NotePropertyType]*ApproachAltitudeTableType)
	for approachaltitudetabletype := range stage.ApproachAltitudeTableTypes {
		_ = approachaltitudetabletype
		for _, _notepropertytype := range approachaltitudetabletype.Annotation {
			stage.ApproachAltitudeTableType_Annotation_reverseMap[_notepropertytype] = approachaltitudetabletype
		}
	}
	stage.ApproachAltitudeTableType_Extension_reverseMap = make(map[*A_extension_14]*ApproachAltitudeTableType)
	for approachaltitudetabletype := range stage.ApproachAltitudeTableTypes {
		_ = approachaltitudetabletype
		for _, _a_extension_14 := range approachaltitudetabletype.Extension {
			stage.ApproachAltitudeTableType_Extension_reverseMap[_a_extension_14] = approachaltitudetabletype
		}
	}

	// Compute reverse map for named struct ApproachConditionType
	// insertion point per field
	stage.ApproachConditionType_CirclingRestriction_reverseMap = make(map[*CirclingRestrictionPropertyType]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _circlingrestrictionpropertytype := range approachconditiontype.CirclingRestriction {
			stage.ApproachConditionType_CirclingRestriction_reverseMap[_circlingrestrictionpropertytype] = approachconditiontype
		}
	}
	stage.ApproachConditionType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _aircraftcharacteristicpropertytype := range approachconditiontype.AircraftCategory {
			stage.ApproachConditionType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = approachconditiontype
		}
	}
	stage.ApproachConditionType_LandingArea_reverseMap = make(map[*LandingTakeoffAreaCollectionPropertyType]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _landingtakeoffareacollectionpropertytype := range approachconditiontype.LandingArea {
			stage.ApproachConditionType_LandingArea_reverseMap[_landingtakeoffareacollectionpropertytype] = approachconditiontype
		}
	}
	stage.ApproachConditionType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _obstacleassessmentareapropertytype := range approachconditiontype.DesignSurface {
			stage.ApproachConditionType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = approachconditiontype
		}
	}
	stage.ApproachConditionType_Annotation_reverseMap = make(map[*NotePropertyType]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _notepropertytype := range approachconditiontype.Annotation {
			stage.ApproachConditionType_Annotation_reverseMap[_notepropertytype] = approachconditiontype
		}
	}
	stage.ApproachConditionType_Extension_reverseMap = make(map[*A_extension_15]*ApproachConditionType)
	for approachconditiontype := range stage.ApproachConditionTypes {
		_ = approachconditiontype
		for _, _a_extension_15 := range approachconditiontype.Extension {
			stage.ApproachConditionType_Extension_reverseMap[_a_extension_15] = approachconditiontype
		}
	}

	// Compute reverse map for named struct ApproachDistanceTableType
	// insertion point per field
	stage.ApproachDistanceTableType_Annotation_reverseMap = make(map[*NotePropertyType]*ApproachDistanceTableType)
	for approachdistancetabletype := range stage.ApproachDistanceTableTypes {
		_ = approachdistancetabletype
		for _, _notepropertytype := range approachdistancetabletype.Annotation {
			stage.ApproachDistanceTableType_Annotation_reverseMap[_notepropertytype] = approachdistancetabletype
		}
	}
	stage.ApproachDistanceTableType_Extension_reverseMap = make(map[*A_extension_16]*ApproachDistanceTableType)
	for approachdistancetabletype := range stage.ApproachDistanceTableTypes {
		_ = approachdistancetabletype
		for _, _a_extension_16 := range approachdistancetabletype.Extension {
			stage.ApproachDistanceTableType_Extension_reverseMap[_a_extension_16] = approachdistancetabletype
		}
	}

	// Compute reverse map for named struct ApproachTimingTableType
	// insertion point per field
	stage.ApproachTimingTableType_Annotation_reverseMap = make(map[*NotePropertyType]*ApproachTimingTableType)
	for approachtimingtabletype := range stage.ApproachTimingTableTypes {
		_ = approachtimingtabletype
		for _, _notepropertytype := range approachtimingtabletype.Annotation {
			stage.ApproachTimingTableType_Annotation_reverseMap[_notepropertytype] = approachtimingtabletype
		}
	}
	stage.ApproachTimingTableType_Extension_reverseMap = make(map[*A_extension_17]*ApproachTimingTableType)
	for approachtimingtabletype := range stage.ApproachTimingTableTypes {
		_ = approachtimingtabletype
		for _, _a_extension_17 := range approachtimingtabletype.Extension {
			stage.ApproachTimingTableType_Extension_reverseMap[_a_extension_17] = approachtimingtabletype
		}
	}

	// Compute reverse map for named struct ArrivalFeederLegTimeSliceType
	// insertion point per field
	stage.ArrivalFeederLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*ArrivalFeederLegTimeSliceType)
	for arrivalfeederlegtimeslicetype := range stage.ArrivalFeederLegTimeSliceTypes {
		_ = arrivalfeederlegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range arrivalfeederlegtimeslicetype.AircraftCategory {
			stage.ArrivalFeederLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = arrivalfeederlegtimeslicetype
		}
	}
	stage.ArrivalFeederLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*ArrivalFeederLegTimeSliceType)
	for arrivalfeederlegtimeslicetype := range stage.ArrivalFeederLegTimeSliceTypes {
		_ = arrivalfeederlegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range arrivalfeederlegtimeslicetype.DesignSurface {
			stage.ArrivalFeederLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = arrivalfeederlegtimeslicetype
		}
	}
	stage.ArrivalFeederLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ArrivalFeederLegTimeSliceType)
	for arrivalfeederlegtimeslicetype := range stage.ArrivalFeederLegTimeSliceTypes {
		_ = arrivalfeederlegtimeslicetype
		for _, _notepropertytype := range arrivalfeederlegtimeslicetype.Annotation {
			stage.ArrivalFeederLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = arrivalfeederlegtimeslicetype
		}
	}
	stage.ArrivalFeederLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_37]*ArrivalFeederLegTimeSliceType)
	for arrivalfeederlegtimeslicetype := range stage.ArrivalFeederLegTimeSliceTypes {
		_ = arrivalfeederlegtimeslicetype
		for _, _a_extension_37 := range arrivalfeederlegtimeslicetype.Extension {
			stage.ArrivalFeederLegTimeSliceType_Extension_reverseMap[_a_extension_37] = arrivalfeederlegtimeslicetype
		}
	}

	// Compute reverse map for named struct ArrivalFeederLegType
	// insertion point per field
	stage.ArrivalFeederLegType_TimeSlice_reverseMap = make(map[*ArrivalFeederLegTimeSlicePropertyType]*ArrivalFeederLegType)
	for arrivalfeederlegtype := range stage.ArrivalFeederLegTypes {
		_ = arrivalfeederlegtype
		for _, _arrivalfeederlegtimeslicepropertytype := range arrivalfeederlegtype.TimeSlice {
			stage.ArrivalFeederLegType_TimeSlice_reverseMap[_arrivalfeederlegtimeslicepropertytype] = arrivalfeederlegtype
		}
	}

	// Compute reverse map for named struct ArrivalLegTimeSliceType
	// insertion point per field
	stage.ArrivalLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*ArrivalLegTimeSliceType)
	for arrivallegtimeslicetype := range stage.ArrivalLegTimeSliceTypes {
		_ = arrivallegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range arrivallegtimeslicetype.AircraftCategory {
			stage.ArrivalLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = arrivallegtimeslicetype
		}
	}
	stage.ArrivalLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*ArrivalLegTimeSliceType)
	for arrivallegtimeslicetype := range stage.ArrivalLegTimeSliceTypes {
		_ = arrivallegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range arrivallegtimeslicetype.DesignSurface {
			stage.ArrivalLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = arrivallegtimeslicetype
		}
	}
	stage.ArrivalLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ArrivalLegTimeSliceType)
	for arrivallegtimeslicetype := range stage.ArrivalLegTimeSliceTypes {
		_ = arrivallegtimeslicetype
		for _, _notepropertytype := range arrivallegtimeslicetype.Annotation {
			stage.ArrivalLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = arrivallegtimeslicetype
		}
	}
	stage.ArrivalLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_38]*ArrivalLegTimeSliceType)
	for arrivallegtimeslicetype := range stage.ArrivalLegTimeSliceTypes {
		_ = arrivallegtimeslicetype
		for _, _a_extension_38 := range arrivallegtimeslicetype.Extension {
			stage.ArrivalLegTimeSliceType_Extension_reverseMap[_a_extension_38] = arrivallegtimeslicetype
		}
	}

	// Compute reverse map for named struct ArrivalLegType
	// insertion point per field
	stage.ArrivalLegType_TimeSlice_reverseMap = make(map[*ArrivalLegTimeSlicePropertyType]*ArrivalLegType)
	for arrivallegtype := range stage.ArrivalLegTypes {
		_ = arrivallegtype
		for _, _arrivallegtimeslicepropertytype := range arrivallegtype.TimeSlice {
			stage.ArrivalLegType_TimeSlice_reverseMap[_arrivallegtimeslicepropertytype] = arrivallegtype
		}
	}

	// Compute reverse map for named struct CircleSectorType
	// insertion point per field
	stage.CircleSectorType_Annotation_reverseMap = make(map[*NotePropertyType]*CircleSectorType)
	for circlesectortype := range stage.CircleSectorTypes {
		_ = circlesectortype
		for _, _notepropertytype := range circlesectortype.Annotation {
			stage.CircleSectorType_Annotation_reverseMap[_notepropertytype] = circlesectortype
		}
	}
	stage.CircleSectorType_Extension_reverseMap = make(map[*A_extension_52]*CircleSectorType)
	for circlesectortype := range stage.CircleSectorTypes {
		_ = circlesectortype
		for _, _a_extension_52 := range circlesectortype.Extension {
			stage.CircleSectorType_Extension_reverseMap[_a_extension_52] = circlesectortype
		}
	}

	// Compute reverse map for named struct CirclingAreaTimeSliceType
	// insertion point per field
	stage.CirclingAreaTimeSliceType_Condition_reverseMap = make(map[*ApproachConditionPropertyType]*CirclingAreaTimeSliceType)
	for circlingareatimeslicetype := range stage.CirclingAreaTimeSliceTypes {
		_ = circlingareatimeslicetype
		for _, _approachconditionpropertytype := range circlingareatimeslicetype.Condition {
			stage.CirclingAreaTimeSliceType_Condition_reverseMap[_approachconditionpropertytype] = circlingareatimeslicetype
		}
	}
	stage.CirclingAreaTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*CirclingAreaTimeSliceType)
	for circlingareatimeslicetype := range stage.CirclingAreaTimeSliceTypes {
		_ = circlingareatimeslicetype
		for _, _obstacleassessmentareapropertytype := range circlingareatimeslicetype.DesignSurface {
			stage.CirclingAreaTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = circlingareatimeslicetype
		}
	}
	stage.CirclingAreaTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*CirclingAreaTimeSliceType)
	for circlingareatimeslicetype := range stage.CirclingAreaTimeSliceTypes {
		_ = circlingareatimeslicetype
		for _, _notepropertytype := range circlingareatimeslicetype.Annotation {
			stage.CirclingAreaTimeSliceType_Annotation_reverseMap[_notepropertytype] = circlingareatimeslicetype
		}
	}
	stage.CirclingAreaTimeSliceType_Extension_reverseMap = make(map[*A_extension_24]*CirclingAreaTimeSliceType)
	for circlingareatimeslicetype := range stage.CirclingAreaTimeSliceTypes {
		_ = circlingareatimeslicetype
		for _, _a_extension_24 := range circlingareatimeslicetype.Extension {
			stage.CirclingAreaTimeSliceType_Extension_reverseMap[_a_extension_24] = circlingareatimeslicetype
		}
	}

	// Compute reverse map for named struct CirclingAreaType
	// insertion point per field
	stage.CirclingAreaType_TimeSlice_reverseMap = make(map[*CirclingAreaTimeSlicePropertyType]*CirclingAreaType)
	for circlingareatype := range stage.CirclingAreaTypes {
		_ = circlingareatype
		for _, _circlingareatimeslicepropertytype := range circlingareatype.TimeSlice {
			stage.CirclingAreaType_TimeSlice_reverseMap[_circlingareatimeslicepropertytype] = circlingareatype
		}
	}

	// Compute reverse map for named struct CirclingRestrictionType
	// insertion point per field
	stage.CirclingRestrictionType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*CirclingRestrictionType)
	for circlingrestrictiontype := range stage.CirclingRestrictionTypes {
		_ = circlingrestrictiontype
		for _, _timesheetpropertytype := range circlingrestrictiontype.TimeInterval {
			stage.CirclingRestrictionType_TimeInterval_reverseMap[_timesheetpropertytype] = circlingrestrictiontype
		}
	}
	stage.CirclingRestrictionType_Annotation_reverseMap = make(map[*NotePropertyType]*CirclingRestrictionType)
	for circlingrestrictiontype := range stage.CirclingRestrictionTypes {
		_ = circlingrestrictiontype
		for _, _notepropertytype := range circlingrestrictiontype.Annotation {
			stage.CirclingRestrictionType_Annotation_reverseMap[_notepropertytype] = circlingrestrictiontype
		}
	}
	stage.CirclingRestrictionType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*CirclingRestrictionType)
	for circlingrestrictiontype := range stage.CirclingRestrictionTypes {
		_ = circlingrestrictiontype
		for _, _organisationauthoritypropertytype := range circlingrestrictiontype.SpecialDateAuthority {
			stage.CirclingRestrictionType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = circlingrestrictiontype
		}
	}
	stage.CirclingRestrictionType_Extension_reverseMap = make(map[*A_extension_25]*CirclingRestrictionType)
	for circlingrestrictiontype := range stage.CirclingRestrictionTypes {
		_ = circlingrestrictiontype
		for _, _a_extension_25 := range circlingrestrictiontype.Extension {
			stage.CirclingRestrictionType_Extension_reverseMap[_a_extension_25] = circlingrestrictiontype
		}
	}

	// Compute reverse map for named struct CurveType
	// insertion point per field
	stage.CurveType_Annotation_reverseMap = make(map[*NotePropertyType]*CurveType)
	for curvetype := range stage.CurveTypes {
		_ = curvetype
		for _, _notepropertytype := range curvetype.Annotation {
			stage.CurveType_Annotation_reverseMap[_notepropertytype] = curvetype
		}
	}

	// Compute reverse map for named struct DepartureArrivalConditionType
	// insertion point per field
	stage.DepartureArrivalConditionType_Annotation_reverseMap = make(map[*NotePropertyType]*DepartureArrivalConditionType)
	for departurearrivalconditiontype := range stage.DepartureArrivalConditionTypes {
		_ = departurearrivalconditiontype
		for _, _notepropertytype := range departurearrivalconditiontype.Annotation {
			stage.DepartureArrivalConditionType_Annotation_reverseMap[_notepropertytype] = departurearrivalconditiontype
		}
	}
	stage.DepartureArrivalConditionType_Extension_reverseMap = make(map[*A_extension_30]*DepartureArrivalConditionType)
	for departurearrivalconditiontype := range stage.DepartureArrivalConditionTypes {
		_ = departurearrivalconditiontype
		for _, _a_extension_30 := range departurearrivalconditiontype.Extension {
			stage.DepartureArrivalConditionType_Extension_reverseMap[_a_extension_30] = departurearrivalconditiontype
		}
	}

	// Compute reverse map for named struct DepartureLegTimeSliceType
	// insertion point per field
	stage.DepartureLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*DepartureLegTimeSliceType)
	for departurelegtimeslicetype := range stage.DepartureLegTimeSliceTypes {
		_ = departurelegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range departurelegtimeslicetype.AircraftCategory {
			stage.DepartureLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = departurelegtimeslicetype
		}
	}
	stage.DepartureLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*DepartureLegTimeSliceType)
	for departurelegtimeslicetype := range stage.DepartureLegTimeSliceTypes {
		_ = departurelegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range departurelegtimeslicetype.DesignSurface {
			stage.DepartureLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = departurelegtimeslicetype
		}
	}
	stage.DepartureLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*DepartureLegTimeSliceType)
	for departurelegtimeslicetype := range stage.DepartureLegTimeSliceTypes {
		_ = departurelegtimeslicetype
		for _, _notepropertytype := range departurelegtimeslicetype.Annotation {
			stage.DepartureLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = departurelegtimeslicetype
		}
	}
	stage.DepartureLegTimeSliceType_Condition_reverseMap = make(map[*DepartureArrivalConditionPropertyType]*DepartureLegTimeSliceType)
	for departurelegtimeslicetype := range stage.DepartureLegTimeSliceTypes {
		_ = departurelegtimeslicetype
		for _, _departurearrivalconditionpropertytype := range departurelegtimeslicetype.Condition {
			stage.DepartureLegTimeSliceType_Condition_reverseMap[_departurearrivalconditionpropertytype] = departurelegtimeslicetype
		}
	}
	stage.DepartureLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_39]*DepartureLegTimeSliceType)
	for departurelegtimeslicetype := range stage.DepartureLegTimeSliceTypes {
		_ = departurelegtimeslicetype
		for _, _a_extension_39 := range departurelegtimeslicetype.Extension {
			stage.DepartureLegTimeSliceType_Extension_reverseMap[_a_extension_39] = departurelegtimeslicetype
		}
	}

	// Compute reverse map for named struct DepartureLegType
	// insertion point per field
	stage.DepartureLegType_TimeSlice_reverseMap = make(map[*DepartureLegTimeSlicePropertyType]*DepartureLegType)
	for departurelegtype := range stage.DepartureLegTypes {
		_ = departurelegtype
		for _, _departurelegtimeslicepropertytype := range departurelegtype.TimeSlice {
			stage.DepartureLegType_TimeSlice_reverseMap[_departurelegtimeslicepropertytype] = departurelegtype
		}
	}

	// Compute reverse map for named struct ElevatedCurveType
	// insertion point per field
	stage.ElevatedCurveType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _notepropertytype := range elevatedcurvetype.Annotation {
			stage.ElevatedCurveType_Annotation_reverseMap[_notepropertytype] = elevatedcurvetype
		}
	}
	stage.ElevatedCurveType_Extension_reverseMap = make(map[*A_extension_1]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _a_extension_1 := range elevatedcurvetype.Extension {
			stage.ElevatedCurveType_Extension_reverseMap[_a_extension_1] = elevatedcurvetype
		}
	}

	// Compute reverse map for named struct ElevatedPointType
	// insertion point per field
	stage.ElevatedPointType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _notepropertytype := range elevatedpointtype.Annotation {
			stage.ElevatedPointType_Annotation_reverseMap[_notepropertytype] = elevatedpointtype
		}
	}
	stage.ElevatedPointType_Extension_reverseMap = make(map[*A_extension_2]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _a_extension_2 := range elevatedpointtype.Extension {
			stage.ElevatedPointType_Extension_reverseMap[_a_extension_2] = elevatedpointtype
		}
	}

	// Compute reverse map for named struct ElevatedSurfaceType
	// insertion point per field
	stage.ElevatedSurfaceType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _notepropertytype := range elevatedsurfacetype.Annotation {
			stage.ElevatedSurfaceType_Annotation_reverseMap[_notepropertytype] = elevatedsurfacetype
		}
	}
	stage.ElevatedSurfaceType_Extension_reverseMap = make(map[*A_extension_3]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _a_extension_3 := range elevatedsurfacetype.Extension {
			stage.ElevatedSurfaceType_Extension_reverseMap[_a_extension_3] = elevatedsurfacetype
		}
	}

	// Compute reverse map for named struct EnRouteSegmentPointType
	// insertion point per field
	stage.EnRouteSegmentPointType_FacilityMakeup_reverseMap = make(map[*PointReferencePropertyType]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _pointreferencepropertytype := range enroutesegmentpointtype.FacilityMakeup {
			stage.EnRouteSegmentPointType_FacilityMakeup_reverseMap[_pointreferencepropertytype] = enroutesegmentpointtype
		}
	}
	stage.EnRouteSegmentPointType_Annotation_reverseMap = make(map[*NotePropertyType]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _notepropertytype := range enroutesegmentpointtype.Annotation {
			stage.EnRouteSegmentPointType_Annotation_reverseMap[_notepropertytype] = enroutesegmentpointtype
		}
	}
	stage.EnRouteSegmentPointType_Extension_reverseMap = make(map[*A_extension_9]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _a_extension_9 := range enroutesegmentpointtype.Extension {
			stage.EnRouteSegmentPointType_Extension_reverseMap[_a_extension_9] = enroutesegmentpointtype
		}
	}

	// Compute reverse map for named struct EquipmentUnavailableAdjustmentColumnType
	// insertion point per field
	stage.EquipmentUnavailableAdjustmentColumnType_Annotation_reverseMap = make(map[*NotePropertyType]*EquipmentUnavailableAdjustmentColumnType)
	for equipmentunavailableadjustmentcolumntype := range stage.EquipmentUnavailableAdjustmentColumnTypes {
		_ = equipmentunavailableadjustmentcolumntype
		for _, _notepropertytype := range equipmentunavailableadjustmentcolumntype.Annotation {
			stage.EquipmentUnavailableAdjustmentColumnType_Annotation_reverseMap[_notepropertytype] = equipmentunavailableadjustmentcolumntype
		}
	}
	stage.EquipmentUnavailableAdjustmentColumnType_Extension_reverseMap = make(map[*A_extension_27]*EquipmentUnavailableAdjustmentColumnType)
	for equipmentunavailableadjustmentcolumntype := range stage.EquipmentUnavailableAdjustmentColumnTypes {
		_ = equipmentunavailableadjustmentcolumntype
		for _, _a_extension_27 := range equipmentunavailableadjustmentcolumntype.Extension {
			stage.EquipmentUnavailableAdjustmentColumnType_Extension_reverseMap[_a_extension_27] = equipmentunavailableadjustmentcolumntype
		}
	}

	// Compute reverse map for named struct EquipmentUnavailableAdjustmentType
	// insertion point per field
	stage.EquipmentUnavailableAdjustmentType_AdjustmentINOPCol_reverseMap = make(map[*EquipmentUnavailableAdjustmentColumnPropertyType]*EquipmentUnavailableAdjustmentType)
	for equipmentunavailableadjustmenttype := range stage.EquipmentUnavailableAdjustmentTypes {
		_ = equipmentunavailableadjustmenttype
		for _, _equipmentunavailableadjustmentcolumnpropertytype := range equipmentunavailableadjustmenttype.AdjustmentINOPCol {
			stage.EquipmentUnavailableAdjustmentType_AdjustmentINOPCol_reverseMap[_equipmentunavailableadjustmentcolumnpropertytype] = equipmentunavailableadjustmenttype
		}
	}
	stage.EquipmentUnavailableAdjustmentType_Annotation_reverseMap = make(map[*NotePropertyType]*EquipmentUnavailableAdjustmentType)
	for equipmentunavailableadjustmenttype := range stage.EquipmentUnavailableAdjustmentTypes {
		_ = equipmentunavailableadjustmenttype
		for _, _notepropertytype := range equipmentunavailableadjustmenttype.Annotation {
			stage.EquipmentUnavailableAdjustmentType_Annotation_reverseMap[_notepropertytype] = equipmentunavailableadjustmenttype
		}
	}
	stage.EquipmentUnavailableAdjustmentType_Extension_reverseMap = make(map[*A_extension_26]*EquipmentUnavailableAdjustmentType)
	for equipmentunavailableadjustmenttype := range stage.EquipmentUnavailableAdjustmentTypes {
		_ = equipmentunavailableadjustmenttype
		for _, _a_extension_26 := range equipmentunavailableadjustmenttype.Extension {
			stage.EquipmentUnavailableAdjustmentType_Extension_reverseMap[_a_extension_26] = equipmentunavailableadjustmenttype
		}
	}

	// Compute reverse map for named struct FASDataBlockType
	// insertion point per field
	stage.FASDataBlockType_Annotation_reverseMap = make(map[*NotePropertyType]*FASDataBlockType)
	for fasdatablocktype := range stage.FASDataBlockTypes {
		_ = fasdatablocktype
		for _, _notepropertytype := range fasdatablocktype.Annotation {
			stage.FASDataBlockType_Annotation_reverseMap[_notepropertytype] = fasdatablocktype
		}
	}
	stage.FASDataBlockType_Extension_reverseMap = make(map[*A_extension_23]*FASDataBlockType)
	for fasdatablocktype := range stage.FASDataBlockTypes {
		_ = fasdatablocktype
		for _, _a_extension_23 := range fasdatablocktype.Extension {
			stage.FASDataBlockType_Extension_reverseMap[_a_extension_23] = fasdatablocktype
		}
	}

	// Compute reverse map for named struct FinalLegTimeSliceType
	// insertion point per field
	stage.FinalLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*FinalLegTimeSliceType)
	for finallegtimeslicetype := range stage.FinalLegTimeSliceTypes {
		_ = finallegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range finallegtimeslicetype.AircraftCategory {
			stage.FinalLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = finallegtimeslicetype
		}
	}
	stage.FinalLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*FinalLegTimeSliceType)
	for finallegtimeslicetype := range stage.FinalLegTimeSliceTypes {
		_ = finallegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range finallegtimeslicetype.DesignSurface {
			stage.FinalLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = finallegtimeslicetype
		}
	}
	stage.FinalLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*FinalLegTimeSliceType)
	for finallegtimeslicetype := range stage.FinalLegTimeSliceTypes {
		_ = finallegtimeslicetype
		for _, _notepropertytype := range finallegtimeslicetype.Annotation {
			stage.FinalLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = finallegtimeslicetype
		}
	}
	stage.FinalLegTimeSliceType_Condition_reverseMap = make(map[*ApproachConditionPropertyType]*FinalLegTimeSliceType)
	for finallegtimeslicetype := range stage.FinalLegTimeSliceTypes {
		_ = finallegtimeslicetype
		for _, _approachconditionpropertytype := range finallegtimeslicetype.Condition {
			stage.FinalLegTimeSliceType_Condition_reverseMap[_approachconditionpropertytype] = finallegtimeslicetype
		}
	}
	stage.FinalLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_40]*FinalLegTimeSliceType)
	for finallegtimeslicetype := range stage.FinalLegTimeSliceTypes {
		_ = finallegtimeslicetype
		for _, _a_extension_40 := range finallegtimeslicetype.Extension {
			stage.FinalLegTimeSliceType_Extension_reverseMap[_a_extension_40] = finallegtimeslicetype
		}
	}

	// Compute reverse map for named struct FinalLegType
	// insertion point per field
	stage.FinalLegType_TimeSlice_reverseMap = make(map[*FinalLegTimeSlicePropertyType]*FinalLegType)
	for finallegtype := range stage.FinalLegTypes {
		_ = finallegtype
		for _, _finallegtimeslicepropertytype := range finallegtype.TimeSlice {
			stage.FinalLegType_TimeSlice_reverseMap[_finallegtimeslicepropertytype] = finallegtype
		}
	}

	// Compute reverse map for named struct FinalProfileType
	// insertion point per field
	stage.FinalProfileType_Altitude_reverseMap = make(map[*ApproachAltitudeTablePropertyType]*FinalProfileType)
	for finalprofiletype := range stage.FinalProfileTypes {
		_ = finalprofiletype
		for _, _approachaltitudetablepropertytype := range finalprofiletype.Altitude {
			stage.FinalProfileType_Altitude_reverseMap[_approachaltitudetablepropertytype] = finalprofiletype
		}
	}
	stage.FinalProfileType_Distance_reverseMap = make(map[*ApproachDistanceTablePropertyType]*FinalProfileType)
	for finalprofiletype := range stage.FinalProfileTypes {
		_ = finalprofiletype
		for _, _approachdistancetablepropertytype := range finalprofiletype.Distance {
			stage.FinalProfileType_Distance_reverseMap[_approachdistancetablepropertytype] = finalprofiletype
		}
	}
	stage.FinalProfileType_Timing_reverseMap = make(map[*ApproachTimingTablePropertyType]*FinalProfileType)
	for finalprofiletype := range stage.FinalProfileTypes {
		_ = finalprofiletype
		for _, _approachtimingtablepropertytype := range finalprofiletype.Timing {
			stage.FinalProfileType_Timing_reverseMap[_approachtimingtablepropertytype] = finalprofiletype
		}
	}
	stage.FinalProfileType_Annotation_reverseMap = make(map[*NotePropertyType]*FinalProfileType)
	for finalprofiletype := range stage.FinalProfileTypes {
		_ = finalprofiletype
		for _, _notepropertytype := range finalprofiletype.Annotation {
			stage.FinalProfileType_Annotation_reverseMap[_notepropertytype] = finalprofiletype
		}
	}
	stage.FinalProfileType_Extension_reverseMap = make(map[*A_extension_18]*FinalProfileType)
	for finalprofiletype := range stage.FinalProfileTypes {
		_ = finalprofiletype
		for _, _a_extension_18 := range finalprofiletype.Extension {
			stage.FinalProfileType_Extension_reverseMap[_a_extension_18] = finalprofiletype
		}
	}

	// Compute reverse map for named struct HoldingAssessmentTimeSliceType
	// insertion point per field
	stage.HoldingAssessmentTimeSliceType_ObstacleAssessment_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*HoldingAssessmentTimeSliceType)
	for holdingassessmenttimeslicetype := range stage.HoldingAssessmentTimeSliceTypes {
		_ = holdingassessmenttimeslicetype
		for _, _obstacleassessmentareapropertytype := range holdingassessmenttimeslicetype.ObstacleAssessment {
			stage.HoldingAssessmentTimeSliceType_ObstacleAssessment_reverseMap[_obstacleassessmentareapropertytype] = holdingassessmenttimeslicetype
		}
	}
	stage.HoldingAssessmentTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*HoldingAssessmentTimeSliceType)
	for holdingassessmenttimeslicetype := range stage.HoldingAssessmentTimeSliceTypes {
		_ = holdingassessmenttimeslicetype
		for _, _notepropertytype := range holdingassessmenttimeslicetype.Annotation {
			stage.HoldingAssessmentTimeSliceType_Annotation_reverseMap[_notepropertytype] = holdingassessmenttimeslicetype
		}
	}
	stage.HoldingAssessmentTimeSliceType_Extension_reverseMap = make(map[*A_extension_55]*HoldingAssessmentTimeSliceType)
	for holdingassessmenttimeslicetype := range stage.HoldingAssessmentTimeSliceTypes {
		_ = holdingassessmenttimeslicetype
		for _, _a_extension_55 := range holdingassessmenttimeslicetype.Extension {
			stage.HoldingAssessmentTimeSliceType_Extension_reverseMap[_a_extension_55] = holdingassessmenttimeslicetype
		}
	}

	// Compute reverse map for named struct HoldingAssessmentType
	// insertion point per field
	stage.HoldingAssessmentType_TimeSlice_reverseMap = make(map[*HoldingAssessmentTimeSlicePropertyType]*HoldingAssessmentType)
	for holdingassessmenttype := range stage.HoldingAssessmentTypes {
		_ = holdingassessmenttype
		for _, _holdingassessmenttimeslicepropertytype := range holdingassessmenttype.TimeSlice {
			stage.HoldingAssessmentType_TimeSlice_reverseMap[_holdingassessmenttimeslicepropertytype] = holdingassessmenttype
		}
	}

	// Compute reverse map for named struct HoldingPatternDistanceType
	// insertion point per field
	stage.HoldingPatternDistanceType_Annotation_reverseMap = make(map[*NotePropertyType]*HoldingPatternDistanceType)
	for holdingpatterndistancetype := range stage.HoldingPatternDistanceTypes {
		_ = holdingpatterndistancetype
		for _, _notepropertytype := range holdingpatterndistancetype.Annotation {
			stage.HoldingPatternDistanceType_Annotation_reverseMap[_notepropertytype] = holdingpatterndistancetype
		}
	}
	stage.HoldingPatternDistanceType_Extension_reverseMap = make(map[*A_extension_5]*HoldingPatternDistanceType)
	for holdingpatterndistancetype := range stage.HoldingPatternDistanceTypes {
		_ = holdingpatterndistancetype
		for _, _a_extension_5 := range holdingpatterndistancetype.Extension {
			stage.HoldingPatternDistanceType_Extension_reverseMap[_a_extension_5] = holdingpatterndistancetype
		}
	}

	// Compute reverse map for named struct HoldingPatternDurationType
	// insertion point per field
	stage.HoldingPatternDurationType_Annotation_reverseMap = make(map[*NotePropertyType]*HoldingPatternDurationType)
	for holdingpatterndurationtype := range stage.HoldingPatternDurationTypes {
		_ = holdingpatterndurationtype
		for _, _notepropertytype := range holdingpatterndurationtype.Annotation {
			stage.HoldingPatternDurationType_Annotation_reverseMap[_notepropertytype] = holdingpatterndurationtype
		}
	}
	stage.HoldingPatternDurationType_Extension_reverseMap = make(map[*A_extension_6]*HoldingPatternDurationType)
	for holdingpatterndurationtype := range stage.HoldingPatternDurationTypes {
		_ = holdingpatterndurationtype
		for _, _a_extension_6 := range holdingpatterndurationtype.Extension {
			stage.HoldingPatternDurationType_Extension_reverseMap[_a_extension_6] = holdingpatterndurationtype
		}
	}

	// Compute reverse map for named struct HoldingPatternTimeSliceType
	// insertion point per field
	stage.HoldingPatternTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*HoldingPatternTimeSliceType)
	for holdingpatterntimeslicetype := range stage.HoldingPatternTimeSliceTypes {
		_ = holdingpatterntimeslicetype
		for _, _notepropertytype := range holdingpatterntimeslicetype.Annotation {
			stage.HoldingPatternTimeSliceType_Annotation_reverseMap[_notepropertytype] = holdingpatterntimeslicetype
		}
	}
	stage.HoldingPatternTimeSliceType_Extension_reverseMap = make(map[*A_extension_4]*HoldingPatternTimeSliceType)
	for holdingpatterntimeslicetype := range stage.HoldingPatternTimeSliceTypes {
		_ = holdingpatterntimeslicetype
		for _, _a_extension_4 := range holdingpatterntimeslicetype.Extension {
			stage.HoldingPatternTimeSliceType_Extension_reverseMap[_a_extension_4] = holdingpatterntimeslicetype
		}
	}

	// Compute reverse map for named struct HoldingPatternType
	// insertion point per field
	stage.HoldingPatternType_TimeSlice_reverseMap = make(map[*HoldingPatternTimeSlicePropertyType]*HoldingPatternType)
	for holdingpatterntype := range stage.HoldingPatternTypes {
		_ = holdingpatterntype
		for _, _holdingpatterntimeslicepropertytype := range holdingpatterntype.TimeSlice {
			stage.HoldingPatternType_TimeSlice_reverseMap[_holdingpatterntimeslicepropertytype] = holdingpatterntype
		}
	}

	// Compute reverse map for named struct HoldingUseType
	// insertion point per field
	stage.HoldingUseType_Annotation_reverseMap = make(map[*NotePropertyType]*HoldingUseType)
	for holdingusetype := range stage.HoldingUseTypes {
		_ = holdingusetype
		for _, _notepropertytype := range holdingusetype.Annotation {
			stage.HoldingUseType_Annotation_reverseMap[_notepropertytype] = holdingusetype
		}
	}
	stage.HoldingUseType_Extension_reverseMap = make(map[*A_extension_41]*HoldingUseType)
	for holdingusetype := range stage.HoldingUseTypes {
		_ = holdingusetype
		for _, _a_extension_41 := range holdingusetype.Extension {
			stage.HoldingUseType_Extension_reverseMap[_a_extension_41] = holdingusetype
		}
	}

	// Compute reverse map for named struct InitialLegTimeSliceType
	// insertion point per field
	stage.InitialLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*InitialLegTimeSliceType)
	for initiallegtimeslicetype := range stage.InitialLegTimeSliceTypes {
		_ = initiallegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range initiallegtimeslicetype.AircraftCategory {
			stage.InitialLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = initiallegtimeslicetype
		}
	}
	stage.InitialLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*InitialLegTimeSliceType)
	for initiallegtimeslicetype := range stage.InitialLegTimeSliceTypes {
		_ = initiallegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range initiallegtimeslicetype.DesignSurface {
			stage.InitialLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = initiallegtimeslicetype
		}
	}
	stage.InitialLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*InitialLegTimeSliceType)
	for initiallegtimeslicetype := range stage.InitialLegTimeSliceTypes {
		_ = initiallegtimeslicetype
		for _, _notepropertytype := range initiallegtimeslicetype.Annotation {
			stage.InitialLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = initiallegtimeslicetype
		}
	}
	stage.InitialLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_42]*InitialLegTimeSliceType)
	for initiallegtimeslicetype := range stage.InitialLegTimeSliceTypes {
		_ = initiallegtimeslicetype
		for _, _a_extension_42 := range initiallegtimeslicetype.Extension {
			stage.InitialLegTimeSliceType_Extension_reverseMap[_a_extension_42] = initiallegtimeslicetype
		}
	}

	// Compute reverse map for named struct InitialLegType
	// insertion point per field
	stage.InitialLegType_TimeSlice_reverseMap = make(map[*InitialLegTimeSlicePropertyType]*InitialLegType)
	for initiallegtype := range stage.InitialLegTypes {
		_ = initiallegtype
		for _, _initiallegtimeslicepropertytype := range initiallegtype.TimeSlice {
			stage.InitialLegType_TimeSlice_reverseMap[_initiallegtimeslicepropertytype] = initiallegtype
		}
	}

	// Compute reverse map for named struct InstrumentApproachProcedureTimeSliceType
	// insertion point per field
	stage.InstrumentApproachProcedureTimeSliceType_Availability_reverseMap = make(map[*ProcedureAvailabilityPropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _procedureavailabilitypropertytype := range instrumentapproachproceduretimeslicetype.Availability {
			stage.InstrumentApproachProcedureTimeSliceType_Availability_reverseMap[_procedureavailabilitypropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _airportheliportpropertytype := range instrumentapproachproceduretimeslicetype.AirportHeliport {
			stage.InstrumentApproachProcedureTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_AircraftCharacteristic_reverseMap = make(map[*AircraftCharacteristicPropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _aircraftcharacteristicpropertytype := range instrumentapproachproceduretimeslicetype.AircraftCharacteristic {
			stage.InstrumentApproachProcedureTimeSliceType_AircraftCharacteristic_reverseMap[_aircraftcharacteristicpropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_FlightTransition_reverseMap = make(map[*ProcedureTransitionPropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _proceduretransitionpropertytype := range instrumentapproachproceduretimeslicetype.FlightTransition {
			stage.InstrumentApproachProcedureTimeSliceType_FlightTransition_reverseMap[_proceduretransitionpropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _notepropertytype := range instrumentapproachproceduretimeslicetype.Annotation {
			stage.InstrumentApproachProcedureTimeSliceType_Annotation_reverseMap[_notepropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_MissedInstruction_reverseMap = make(map[*MissedApproachGroupPropertyType]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _missedapproachgrouppropertytype := range instrumentapproachproceduretimeslicetype.MissedInstruction {
			stage.InstrumentApproachProcedureTimeSliceType_MissedInstruction_reverseMap[_missedapproachgrouppropertytype] = instrumentapproachproceduretimeslicetype
		}
	}
	stage.InstrumentApproachProcedureTimeSliceType_Extension_reverseMap = make(map[*A_extension_19]*InstrumentApproachProcedureTimeSliceType)
	for instrumentapproachproceduretimeslicetype := range stage.InstrumentApproachProcedureTimeSliceTypes {
		_ = instrumentapproachproceduretimeslicetype
		for _, _a_extension_19 := range instrumentapproachproceduretimeslicetype.Extension {
			stage.InstrumentApproachProcedureTimeSliceType_Extension_reverseMap[_a_extension_19] = instrumentapproachproceduretimeslicetype
		}
	}

	// Compute reverse map for named struct InstrumentApproachProcedureType
	// insertion point per field
	stage.InstrumentApproachProcedureType_TimeSlice_reverseMap = make(map[*InstrumentApproachProcedureTimeSlicePropertyType]*InstrumentApproachProcedureType)
	for instrumentapproachproceduretype := range stage.InstrumentApproachProcedureTypes {
		_ = instrumentapproachproceduretype
		for _, _instrumentapproachproceduretimeslicepropertytype := range instrumentapproachproceduretype.TimeSlice {
			stage.InstrumentApproachProcedureType_TimeSlice_reverseMap[_instrumentapproachproceduretimeslicepropertytype] = instrumentapproachproceduretype
		}
	}

	// Compute reverse map for named struct IntermediateLegTimeSliceType
	// insertion point per field
	stage.IntermediateLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*IntermediateLegTimeSliceType)
	for intermediatelegtimeslicetype := range stage.IntermediateLegTimeSliceTypes {
		_ = intermediatelegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range intermediatelegtimeslicetype.AircraftCategory {
			stage.IntermediateLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = intermediatelegtimeslicetype
		}
	}
	stage.IntermediateLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*IntermediateLegTimeSliceType)
	for intermediatelegtimeslicetype := range stage.IntermediateLegTimeSliceTypes {
		_ = intermediatelegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range intermediatelegtimeslicetype.DesignSurface {
			stage.IntermediateLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = intermediatelegtimeslicetype
		}
	}
	stage.IntermediateLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*IntermediateLegTimeSliceType)
	for intermediatelegtimeslicetype := range stage.IntermediateLegTimeSliceTypes {
		_ = intermediatelegtimeslicetype
		for _, _notepropertytype := range intermediatelegtimeslicetype.Annotation {
			stage.IntermediateLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = intermediatelegtimeslicetype
		}
	}
	stage.IntermediateLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_43]*IntermediateLegTimeSliceType)
	for intermediatelegtimeslicetype := range stage.IntermediateLegTimeSliceTypes {
		_ = intermediatelegtimeslicetype
		for _, _a_extension_43 := range intermediatelegtimeslicetype.Extension {
			stage.IntermediateLegTimeSliceType_Extension_reverseMap[_a_extension_43] = intermediatelegtimeslicetype
		}
	}

	// Compute reverse map for named struct IntermediateLegType
	// insertion point per field
	stage.IntermediateLegType_TimeSlice_reverseMap = make(map[*IntermediateLegTimeSlicePropertyType]*IntermediateLegType)
	for intermediatelegtype := range stage.IntermediateLegTypes {
		_ = intermediatelegtype
		for _, _intermediatelegtimeslicepropertytype := range intermediatelegtype.TimeSlice {
			stage.IntermediateLegType_TimeSlice_reverseMap[_intermediatelegtimeslicepropertytype] = intermediatelegtype
		}
	}

	// Compute reverse map for named struct LandingTakeoffAreaCollectionType
	// insertion point per field
	stage.LandingTakeoffAreaCollectionType_Runway_reverseMap = make(map[*RunwayDirectionPropertyType]*LandingTakeoffAreaCollectionType)
	for landingtakeoffareacollectiontype := range stage.LandingTakeoffAreaCollectionTypes {
		_ = landingtakeoffareacollectiontype
		for _, _runwaydirectionpropertytype := range landingtakeoffareacollectiontype.Runway {
			stage.LandingTakeoffAreaCollectionType_Runway_reverseMap[_runwaydirectionpropertytype] = landingtakeoffareacollectiontype
		}
	}
	stage.LandingTakeoffAreaCollectionType_TLOF_reverseMap = make(map[*TouchDownLiftOffPropertyType]*LandingTakeoffAreaCollectionType)
	for landingtakeoffareacollectiontype := range stage.LandingTakeoffAreaCollectionTypes {
		_ = landingtakeoffareacollectiontype
		for _, _touchdownliftoffpropertytype := range landingtakeoffareacollectiontype.TLOF {
			stage.LandingTakeoffAreaCollectionType_TLOF_reverseMap[_touchdownliftoffpropertytype] = landingtakeoffareacollectiontype
		}
	}
	stage.LandingTakeoffAreaCollectionType_Annotation_reverseMap = make(map[*NotePropertyType]*LandingTakeoffAreaCollectionType)
	for landingtakeoffareacollectiontype := range stage.LandingTakeoffAreaCollectionTypes {
		_ = landingtakeoffareacollectiontype
		for _, _notepropertytype := range landingtakeoffareacollectiontype.Annotation {
			stage.LandingTakeoffAreaCollectionType_Annotation_reverseMap[_notepropertytype] = landingtakeoffareacollectiontype
		}
	}
	stage.LandingTakeoffAreaCollectionType_Extension_reverseMap = make(map[*A_extension_44]*LandingTakeoffAreaCollectionType)
	for landingtakeoffareacollectiontype := range stage.LandingTakeoffAreaCollectionTypes {
		_ = landingtakeoffareacollectiontype
		for _, _a_extension_44 := range landingtakeoffareacollectiontype.Extension {
			stage.LandingTakeoffAreaCollectionType_Extension_reverseMap[_a_extension_44] = landingtakeoffareacollectiontype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_12]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_12 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_12] = linguisticnotetype
		}
	}

	// Compute reverse map for named struct MinimaType
	// insertion point per field
	stage.MinimaType_AdjustmentINOP_reverseMap = make(map[*EquipmentUnavailableAdjustmentPropertyType]*MinimaType)
	for minimatype := range stage.MinimaTypes {
		_ = minimatype
		for _, _equipmentunavailableadjustmentpropertytype := range minimatype.AdjustmentINOP {
			stage.MinimaType_AdjustmentINOP_reverseMap[_equipmentunavailableadjustmentpropertytype] = minimatype
		}
	}
	stage.MinimaType_Annotation_reverseMap = make(map[*NotePropertyType]*MinimaType)
	for minimatype := range stage.MinimaTypes {
		_ = minimatype
		for _, _notepropertytype := range minimatype.Annotation {
			stage.MinimaType_Annotation_reverseMap[_notepropertytype] = minimatype
		}
	}
	stage.MinimaType_Extension_reverseMap = make(map[*A_extension_28]*MinimaType)
	for minimatype := range stage.MinimaTypes {
		_ = minimatype
		for _, _a_extension_28 := range minimatype.Extension {
			stage.MinimaType_Extension_reverseMap[_a_extension_28] = minimatype
		}
	}

	// Compute reverse map for named struct MissedApproachGroupType
	// insertion point per field
	stage.MissedApproachGroupType_Altimeter_reverseMap = make(map[*AltimeterSourcePropertyType]*MissedApproachGroupType)
	for missedapproachgrouptype := range stage.MissedApproachGroupTypes {
		_ = missedapproachgrouptype
		for _, _altimetersourcepropertytype := range missedapproachgrouptype.Altimeter {
			stage.MissedApproachGroupType_Altimeter_reverseMap[_altimetersourcepropertytype] = missedapproachgrouptype
		}
	}
	stage.MissedApproachGroupType_Annotation_reverseMap = make(map[*NotePropertyType]*MissedApproachGroupType)
	for missedapproachgrouptype := range stage.MissedApproachGroupTypes {
		_ = missedapproachgrouptype
		for _, _notepropertytype := range missedapproachgrouptype.Annotation {
			stage.MissedApproachGroupType_Annotation_reverseMap[_notepropertytype] = missedapproachgrouptype
		}
	}
	stage.MissedApproachGroupType_Extension_reverseMap = make(map[*A_extension_20]*MissedApproachGroupType)
	for missedapproachgrouptype := range stage.MissedApproachGroupTypes {
		_ = missedapproachgrouptype
		for _, _a_extension_20 := range missedapproachgrouptype.Extension {
			stage.MissedApproachGroupType_Extension_reverseMap[_a_extension_20] = missedapproachgrouptype
		}
	}

	// Compute reverse map for named struct MissedApproachLegTimeSliceType
	// insertion point per field
	stage.MissedApproachLegTimeSliceType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*MissedApproachLegTimeSliceType)
	for missedapproachlegtimeslicetype := range stage.MissedApproachLegTimeSliceTypes {
		_ = missedapproachlegtimeslicetype
		for _, _aircraftcharacteristicpropertytype := range missedapproachlegtimeslicetype.AircraftCategory {
			stage.MissedApproachLegTimeSliceType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = missedapproachlegtimeslicetype
		}
	}
	stage.MissedApproachLegTimeSliceType_DesignSurface_reverseMap = make(map[*ObstacleAssessmentAreaPropertyType]*MissedApproachLegTimeSliceType)
	for missedapproachlegtimeslicetype := range stage.MissedApproachLegTimeSliceTypes {
		_ = missedapproachlegtimeslicetype
		for _, _obstacleassessmentareapropertytype := range missedapproachlegtimeslicetype.DesignSurface {
			stage.MissedApproachLegTimeSliceType_DesignSurface_reverseMap[_obstacleassessmentareapropertytype] = missedapproachlegtimeslicetype
		}
	}
	stage.MissedApproachLegTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*MissedApproachLegTimeSliceType)
	for missedapproachlegtimeslicetype := range stage.MissedApproachLegTimeSliceTypes {
		_ = missedapproachlegtimeslicetype
		for _, _notepropertytype := range missedapproachlegtimeslicetype.Annotation {
			stage.MissedApproachLegTimeSliceType_Annotation_reverseMap[_notepropertytype] = missedapproachlegtimeslicetype
		}
	}
	stage.MissedApproachLegTimeSliceType_Condition_reverseMap = make(map[*ApproachConditionPropertyType]*MissedApproachLegTimeSliceType)
	for missedapproachlegtimeslicetype := range stage.MissedApproachLegTimeSliceTypes {
		_ = missedapproachlegtimeslicetype
		for _, _approachconditionpropertytype := range missedapproachlegtimeslicetype.Condition {
			stage.MissedApproachLegTimeSliceType_Condition_reverseMap[_approachconditionpropertytype] = missedapproachlegtimeslicetype
		}
	}
	stage.MissedApproachLegTimeSliceType_Extension_reverseMap = make(map[*A_extension_45]*MissedApproachLegTimeSliceType)
	for missedapproachlegtimeslicetype := range stage.MissedApproachLegTimeSliceTypes {
		_ = missedapproachlegtimeslicetype
		for _, _a_extension_45 := range missedapproachlegtimeslicetype.Extension {
			stage.MissedApproachLegTimeSliceType_Extension_reverseMap[_a_extension_45] = missedapproachlegtimeslicetype
		}
	}

	// Compute reverse map for named struct MissedApproachLegType
	// insertion point per field
	stage.MissedApproachLegType_TimeSlice_reverseMap = make(map[*MissedApproachLegTimeSlicePropertyType]*MissedApproachLegType)
	for missedapproachlegtype := range stage.MissedApproachLegTypes {
		_ = missedapproachlegtype
		for _, _missedapproachlegtimeslicepropertytype := range missedapproachlegtype.TimeSlice {
			stage.MissedApproachLegType_TimeSlice_reverseMap[_missedapproachlegtimeslicepropertytype] = missedapproachlegtype
		}
	}

	// Compute reverse map for named struct NavigationAreaRestrictionTimeSliceType
	// insertion point per field
	stage.NavigationAreaRestrictionTimeSliceType_Procedure_reverseMap = make(map[*ProcedurePropertyType]*NavigationAreaRestrictionTimeSliceType)
	for navigationarearestrictiontimeslicetype := range stage.NavigationAreaRestrictionTimeSliceTypes {
		_ = navigationarearestrictiontimeslicetype
		for _, _procedurepropertytype := range navigationarearestrictiontimeslicetype.Procedure {
			stage.NavigationAreaRestrictionTimeSliceType_Procedure_reverseMap[_procedurepropertytype] = navigationarearestrictiontimeslicetype
		}
	}
	stage.NavigationAreaRestrictionTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*NavigationAreaRestrictionTimeSliceType)
	for navigationarearestrictiontimeslicetype := range stage.NavigationAreaRestrictionTimeSliceTypes {
		_ = navigationarearestrictiontimeslicetype
		for _, _notepropertytype := range navigationarearestrictiontimeslicetype.Annotation {
			stage.NavigationAreaRestrictionTimeSliceType_Annotation_reverseMap[_notepropertytype] = navigationarearestrictiontimeslicetype
		}
	}
	stage.NavigationAreaRestrictionTimeSliceType_Extension_reverseMap = make(map[*A_extension_46]*NavigationAreaRestrictionTimeSliceType)
	for navigationarearestrictiontimeslicetype := range stage.NavigationAreaRestrictionTimeSliceTypes {
		_ = navigationarearestrictiontimeslicetype
		for _, _a_extension_46 := range navigationarearestrictiontimeslicetype.Extension {
			stage.NavigationAreaRestrictionTimeSliceType_Extension_reverseMap[_a_extension_46] = navigationarearestrictiontimeslicetype
		}
	}

	// Compute reverse map for named struct NavigationAreaRestrictionType
	// insertion point per field
	stage.NavigationAreaRestrictionType_TimeSlice_reverseMap = make(map[*NavigationAreaRestrictionTimeSlicePropertyType]*NavigationAreaRestrictionType)
	for navigationarearestrictiontype := range stage.NavigationAreaRestrictionTypes {
		_ = navigationarearestrictiontype
		for _, _navigationarearestrictiontimeslicepropertytype := range navigationarearestrictiontype.TimeSlice {
			stage.NavigationAreaRestrictionType_TimeSlice_reverseMap[_navigationarearestrictiontimeslicepropertytype] = navigationarearestrictiontype
		}
	}

	// Compute reverse map for named struct NavigationAreaSectorType
	// insertion point per field
	stage.NavigationAreaSectorType_SignificantObstacle_reverseMap = make(map[*ObstructionPropertyType]*NavigationAreaSectorType)
	for navigationareasectortype := range stage.NavigationAreaSectorTypes {
		_ = navigationareasectortype
		for _, _obstructionpropertytype := range navigationareasectortype.SignificantObstacle {
			stage.NavigationAreaSectorType_SignificantObstacle_reverseMap[_obstructionpropertytype] = navigationareasectortype
		}
	}
	stage.NavigationAreaSectorType_SectorCriteria_reverseMap = make(map[*SectorDesignPropertyType]*NavigationAreaSectorType)
	for navigationareasectortype := range stage.NavigationAreaSectorTypes {
		_ = navigationareasectortype
		for _, _sectordesignpropertytype := range navigationareasectortype.SectorCriteria {
			stage.NavigationAreaSectorType_SectorCriteria_reverseMap[_sectordesignpropertytype] = navigationareasectortype
		}
	}
	stage.NavigationAreaSectorType_Annotation_reverseMap = make(map[*NotePropertyType]*NavigationAreaSectorType)
	for navigationareasectortype := range stage.NavigationAreaSectorTypes {
		_ = navigationareasectortype
		for _, _notepropertytype := range navigationareasectortype.Annotation {
			stage.NavigationAreaSectorType_Annotation_reverseMap[_notepropertytype] = navigationareasectortype
		}
	}
	stage.NavigationAreaSectorType_Extension_reverseMap = make(map[*A_extension_32]*NavigationAreaSectorType)
	for navigationareasectortype := range stage.NavigationAreaSectorTypes {
		_ = navigationareasectortype
		for _, _a_extension_32 := range navigationareasectortype.Extension {
			stage.NavigationAreaSectorType_Extension_reverseMap[_a_extension_32] = navigationareasectortype
		}
	}

	// Compute reverse map for named struct NavigationAreaTimeSliceType
	// insertion point per field
	stage.NavigationAreaTimeSliceType_Sector_reverseMap = make(map[*NavigationAreaSectorPropertyType]*NavigationAreaTimeSliceType)
	for navigationareatimeslicetype := range stage.NavigationAreaTimeSliceTypes {
		_ = navigationareatimeslicetype
		for _, _navigationareasectorpropertytype := range navigationareatimeslicetype.Sector {
			stage.NavigationAreaTimeSliceType_Sector_reverseMap[_navigationareasectorpropertytype] = navigationareatimeslicetype
		}
	}
	stage.NavigationAreaTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*NavigationAreaTimeSliceType)
	for navigationareatimeslicetype := range stage.NavigationAreaTimeSliceTypes {
		_ = navigationareatimeslicetype
		for _, _notepropertytype := range navigationareatimeslicetype.Annotation {
			stage.NavigationAreaTimeSliceType_Annotation_reverseMap[_notepropertytype] = navigationareatimeslicetype
		}
	}
	stage.NavigationAreaTimeSliceType_Extension_reverseMap = make(map[*A_extension_31]*NavigationAreaTimeSliceType)
	for navigationareatimeslicetype := range stage.NavigationAreaTimeSliceTypes {
		_ = navigationareatimeslicetype
		for _, _a_extension_31 := range navigationareatimeslicetype.Extension {
			stage.NavigationAreaTimeSliceType_Extension_reverseMap[_a_extension_31] = navigationareatimeslicetype
		}
	}

	// Compute reverse map for named struct NavigationAreaType
	// insertion point per field
	stage.NavigationAreaType_TimeSlice_reverseMap = make(map[*NavigationAreaTimeSlicePropertyType]*NavigationAreaType)
	for navigationareatype := range stage.NavigationAreaTypes {
		_ = navigationareatype
		for _, _navigationareatimeslicepropertytype := range navigationareatype.TimeSlice {
			stage.NavigationAreaType_TimeSlice_reverseMap[_navigationareatimeslicepropertytype] = navigationareatype
		}
	}

	// Compute reverse map for named struct NoteType
	// insertion point per field
	stage.NoteType_TranslatedNote_reverseMap = make(map[*LinguisticNotePropertyType]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
			stage.NoteType_TranslatedNote_reverseMap[_linguisticnotepropertytype] = notetype
		}
	}
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_13]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_13 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_13] = notetype
		}
	}

	// Compute reverse map for named struct ObstacleAssessmentAreaType
	// insertion point per field
	stage.ObstacleAssessmentAreaType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _aircraftcharacteristicpropertytype := range obstacleassessmentareatype.AircraftCategory {
			stage.ObstacleAssessmentAreaType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_SignificantObstacle_reverseMap = make(map[*ObstructionPropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _obstructionpropertytype := range obstacleassessmentareatype.SignificantObstacle {
			stage.ObstacleAssessmentAreaType_SignificantObstacle_reverseMap[_obstructionpropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _notepropertytype := range obstacleassessmentareatype.Annotation {
			stage.ObstacleAssessmentAreaType_Annotation_reverseMap[_notepropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_Extension_reverseMap = make(map[*A_extension_56]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _a_extension_56 := range obstacleassessmentareatype.Extension {
			stage.ObstacleAssessmentAreaType_Extension_reverseMap[_a_extension_56] = obstacleassessmentareatype
		}
	}

	// Compute reverse map for named struct ObstaclePlacementType
	// insertion point per field
	stage.ObstaclePlacementType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstaclePlacementType)
	for obstacleplacementtype := range stage.ObstaclePlacementTypes {
		_ = obstacleplacementtype
		for _, _notepropertytype := range obstacleplacementtype.Annotation {
			stage.ObstaclePlacementType_Annotation_reverseMap[_notepropertytype] = obstacleplacementtype
		}
	}
	stage.ObstaclePlacementType_Extension_reverseMap = make(map[*A_extension_57]*ObstaclePlacementType)
	for obstacleplacementtype := range stage.ObstaclePlacementTypes {
		_ = obstacleplacementtype
		for _, _a_extension_57 := range obstacleplacementtype.Extension {
			stage.ObstaclePlacementType_Extension_reverseMap[_a_extension_57] = obstacleplacementtype
		}
	}

	// Compute reverse map for named struct ObstructionType
	// insertion point per field
	stage.ObstructionType_Adjustment_reverseMap = make(map[*AltitudeAdjustmentPropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _altitudeadjustmentpropertytype := range obstructiontype.Adjustment {
			stage.ObstructionType_Adjustment_reverseMap[_altitudeadjustmentpropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_ObstaclePlacement_reverseMap = make(map[*ObstaclePlacementPropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _obstacleplacementpropertytype := range obstructiontype.ObstaclePlacement {
			stage.ObstructionType_ObstaclePlacement_reverseMap[_obstacleplacementpropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _notepropertytype := range obstructiontype.Annotation {
			stage.ObstructionType_Annotation_reverseMap[_notepropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_Extension_reverseMap = make(map[*A_extension_58]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _a_extension_58 := range obstructiontype.Extension {
			stage.ObstructionType_Extension_reverseMap[_a_extension_58] = obstructiontype
		}
	}

	// Compute reverse map for named struct PointReferenceType
	// insertion point per field
	stage.PointReferenceType_FacilityAngle_reverseMap = make(map[*AngleUsePropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _angleusepropertytype := range pointreferencetype.FacilityAngle {
			stage.PointReferenceType_FacilityAngle_reverseMap[_angleusepropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_FacilityDistance_reverseMap = make(map[*DistanceIndicationPropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _distanceindicationpropertytype := range pointreferencetype.FacilityDistance {
			stage.PointReferenceType_FacilityDistance_reverseMap[_distanceindicationpropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_Annotation_reverseMap = make(map[*NotePropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _notepropertytype := range pointreferencetype.Annotation {
			stage.PointReferenceType_Annotation_reverseMap[_notepropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_Extension_reverseMap = make(map[*A_extension_10]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _a_extension_10 := range pointreferencetype.Extension {
			stage.PointReferenceType_Extension_reverseMap[_a_extension_10] = pointreferencetype
		}
	}

	// Compute reverse map for named struct PointType
	// insertion point per field
	stage.PointType_Annotation_reverseMap = make(map[*NotePropertyType]*PointType)
	for pointtype := range stage.PointTypes {
		_ = pointtype
		for _, _notepropertytype := range pointtype.Annotation {
			stage.PointType_Annotation_reverseMap[_notepropertytype] = pointtype
		}
	}

	// Compute reverse map for named struct ProcedureAvailabilityType
	// insertion point per field
	stage.ProcedureAvailabilityType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*ProcedureAvailabilityType)
	for procedureavailabilitytype := range stage.ProcedureAvailabilityTypes {
		_ = procedureavailabilitytype
		for _, _timesheetpropertytype := range procedureavailabilitytype.TimeInterval {
			stage.ProcedureAvailabilityType_TimeInterval_reverseMap[_timesheetpropertytype] = procedureavailabilitytype
		}
	}
	stage.ProcedureAvailabilityType_Annotation_reverseMap = make(map[*NotePropertyType]*ProcedureAvailabilityType)
	for procedureavailabilitytype := range stage.ProcedureAvailabilityTypes {
		_ = procedureavailabilitytype
		for _, _notepropertytype := range procedureavailabilitytype.Annotation {
			stage.ProcedureAvailabilityType_Annotation_reverseMap[_notepropertytype] = procedureavailabilitytype
		}
	}
	stage.ProcedureAvailabilityType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*ProcedureAvailabilityType)
	for procedureavailabilitytype := range stage.ProcedureAvailabilityTypes {
		_ = procedureavailabilitytype
		for _, _organisationauthoritypropertytype := range procedureavailabilitytype.SpecialDateAuthority {
			stage.ProcedureAvailabilityType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = procedureavailabilitytype
		}
	}
	stage.ProcedureAvailabilityType_Extension_reverseMap = make(map[*A_extension_50]*ProcedureAvailabilityType)
	for procedureavailabilitytype := range stage.ProcedureAvailabilityTypes {
		_ = procedureavailabilitytype
		for _, _a_extension_50 := range procedureavailabilitytype.Extension {
			stage.ProcedureAvailabilityType_Extension_reverseMap[_a_extension_50] = procedureavailabilitytype
		}
	}

	// Compute reverse map for named struct ProcedureDMETimeSliceType
	// insertion point per field
	stage.ProcedureDMETimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ProcedureDMETimeSliceType)
	for proceduredmetimeslicetype := range stage.ProcedureDMETimeSliceTypes {
		_ = proceduredmetimeslicetype
		for _, _notepropertytype := range proceduredmetimeslicetype.Annotation {
			stage.ProcedureDMETimeSliceType_Annotation_reverseMap[_notepropertytype] = proceduredmetimeslicetype
		}
	}
	stage.ProcedureDMETimeSliceType_Extension_reverseMap = make(map[*A_extension_47]*ProcedureDMETimeSliceType)
	for proceduredmetimeslicetype := range stage.ProcedureDMETimeSliceTypes {
		_ = proceduredmetimeslicetype
		for _, _a_extension_47 := range proceduredmetimeslicetype.Extension {
			stage.ProcedureDMETimeSliceType_Extension_reverseMap[_a_extension_47] = proceduredmetimeslicetype
		}
	}

	// Compute reverse map for named struct ProcedureDMEType
	// insertion point per field
	stage.ProcedureDMEType_TimeSlice_reverseMap = make(map[*ProcedureDMETimeSlicePropertyType]*ProcedureDMEType)
	for proceduredmetype := range stage.ProcedureDMETypes {
		_ = proceduredmetype
		for _, _proceduredmetimeslicepropertytype := range proceduredmetype.TimeSlice {
			stage.ProcedureDMEType_TimeSlice_reverseMap[_proceduredmetimeslicepropertytype] = proceduredmetype
		}
	}

	// Compute reverse map for named struct ProcedureTransitionLegType
	// insertion point per field
	stage.ProcedureTransitionLegType_Annotation_reverseMap = make(map[*NotePropertyType]*ProcedureTransitionLegType)
	for proceduretransitionlegtype := range stage.ProcedureTransitionLegTypes {
		_ = proceduretransitionlegtype
		for _, _notepropertytype := range proceduretransitionlegtype.Annotation {
			stage.ProcedureTransitionLegType_Annotation_reverseMap[_notepropertytype] = proceduretransitionlegtype
		}
	}
	stage.ProcedureTransitionLegType_Extension_reverseMap = make(map[*A_extension_49]*ProcedureTransitionLegType)
	for proceduretransitionlegtype := range stage.ProcedureTransitionLegTypes {
		_ = proceduretransitionlegtype
		for _, _a_extension_49 := range proceduretransitionlegtype.Extension {
			stage.ProcedureTransitionLegType_Extension_reverseMap[_a_extension_49] = proceduretransitionlegtype
		}
	}

	// Compute reverse map for named struct ProcedureTransitionType
	// insertion point per field
	stage.ProcedureTransitionType_TransitionLeg_reverseMap = make(map[*ProcedureTransitionLegPropertyType]*ProcedureTransitionType)
	for proceduretransitiontype := range stage.ProcedureTransitionTypes {
		_ = proceduretransitiontype
		for _, _proceduretransitionlegpropertytype := range proceduretransitiontype.TransitionLeg {
			stage.ProcedureTransitionType_TransitionLeg_reverseMap[_proceduretransitionlegpropertytype] = proceduretransitiontype
		}
	}
	stage.ProcedureTransitionType_Annotation_reverseMap = make(map[*NotePropertyType]*ProcedureTransitionType)
	for proceduretransitiontype := range stage.ProcedureTransitionTypes {
		_ = proceduretransitiontype
		for _, _notepropertytype := range proceduretransitiontype.Annotation {
			stage.ProcedureTransitionType_Annotation_reverseMap[_notepropertytype] = proceduretransitiontype
		}
	}
	stage.ProcedureTransitionType_Extension_reverseMap = make(map[*A_extension_48]*ProcedureTransitionType)
	for proceduretransitiontype := range stage.ProcedureTransitionTypes {
		_ = proceduretransitiontype
		for _, _a_extension_48 := range proceduretransitiontype.Extension {
			stage.ProcedureTransitionType_Extension_reverseMap[_a_extension_48] = proceduretransitiontype
		}
	}

	// Compute reverse map for named struct SafeAltitudeAreaSectorType
	// insertion point per field
	stage.SafeAltitudeAreaSectorType_SignificantObstacle_reverseMap = make(map[*ObstructionPropertyType]*SafeAltitudeAreaSectorType)
	for safealtitudeareasectortype := range stage.SafeAltitudeAreaSectorTypes {
		_ = safealtitudeareasectortype
		for _, _obstructionpropertytype := range safealtitudeareasectortype.SignificantObstacle {
			stage.SafeAltitudeAreaSectorType_SignificantObstacle_reverseMap[_obstructionpropertytype] = safealtitudeareasectortype
		}
	}
	stage.SafeAltitudeAreaSectorType_Annotation_reverseMap = make(map[*NotePropertyType]*SafeAltitudeAreaSectorType)
	for safealtitudeareasectortype := range stage.SafeAltitudeAreaSectorTypes {
		_ = safealtitudeareasectortype
		for _, _notepropertytype := range safealtitudeareasectortype.Annotation {
			stage.SafeAltitudeAreaSectorType_Annotation_reverseMap[_notepropertytype] = safealtitudeareasectortype
		}
	}
	stage.SafeAltitudeAreaSectorType_Extension_reverseMap = make(map[*A_extension_36]*SafeAltitudeAreaSectorType)
	for safealtitudeareasectortype := range stage.SafeAltitudeAreaSectorTypes {
		_ = safealtitudeareasectortype
		for _, _a_extension_36 := range safealtitudeareasectortype.Extension {
			stage.SafeAltitudeAreaSectorType_Extension_reverseMap[_a_extension_36] = safealtitudeareasectortype
		}
	}

	// Compute reverse map for named struct SafeAltitudeAreaTimeSliceType
	// insertion point per field
	stage.SafeAltitudeAreaTimeSliceType_Sector_reverseMap = make(map[*SafeAltitudeAreaSectorPropertyType]*SafeAltitudeAreaTimeSliceType)
	for safealtitudeareatimeslicetype := range stage.SafeAltitudeAreaTimeSliceTypes {
		_ = safealtitudeareatimeslicetype
		for _, _safealtitudeareasectorpropertytype := range safealtitudeareatimeslicetype.Sector {
			stage.SafeAltitudeAreaTimeSliceType_Sector_reverseMap[_safealtitudeareasectorpropertytype] = safealtitudeareatimeslicetype
		}
	}
	stage.SafeAltitudeAreaTimeSliceType_Location_reverseMap = make(map[*AirportHeliportPropertyType]*SafeAltitudeAreaTimeSliceType)
	for safealtitudeareatimeslicetype := range stage.SafeAltitudeAreaTimeSliceTypes {
		_ = safealtitudeareatimeslicetype
		for _, _airportheliportpropertytype := range safealtitudeareatimeslicetype.Location {
			stage.SafeAltitudeAreaTimeSliceType_Location_reverseMap[_airportheliportpropertytype] = safealtitudeareatimeslicetype
		}
	}
	stage.SafeAltitudeAreaTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SafeAltitudeAreaTimeSliceType)
	for safealtitudeareatimeslicetype := range stage.SafeAltitudeAreaTimeSliceTypes {
		_ = safealtitudeareatimeslicetype
		for _, _notepropertytype := range safealtitudeareatimeslicetype.Annotation {
			stage.SafeAltitudeAreaTimeSliceType_Annotation_reverseMap[_notepropertytype] = safealtitudeareatimeslicetype
		}
	}
	stage.SafeAltitudeAreaTimeSliceType_Extension_reverseMap = make(map[*A_extension_35]*SafeAltitudeAreaTimeSliceType)
	for safealtitudeareatimeslicetype := range stage.SafeAltitudeAreaTimeSliceTypes {
		_ = safealtitudeareatimeslicetype
		for _, _a_extension_35 := range safealtitudeareatimeslicetype.Extension {
			stage.SafeAltitudeAreaTimeSliceType_Extension_reverseMap[_a_extension_35] = safealtitudeareatimeslicetype
		}
	}

	// Compute reverse map for named struct SafeAltitudeAreaType
	// insertion point per field
	stage.SafeAltitudeAreaType_TimeSlice_reverseMap = make(map[*SafeAltitudeAreaTimeSlicePropertyType]*SafeAltitudeAreaType)
	for safealtitudeareatype := range stage.SafeAltitudeAreaTypes {
		_ = safealtitudeareatype
		for _, _safealtitudeareatimeslicepropertytype := range safealtitudeareatype.TimeSlice {
			stage.SafeAltitudeAreaType_TimeSlice_reverseMap[_safealtitudeareatimeslicepropertytype] = safealtitudeareatype
		}
	}

	// Compute reverse map for named struct SectorDesignType
	// insertion point per field
	stage.SectorDesignType_Annotation_reverseMap = make(map[*NotePropertyType]*SectorDesignType)
	for sectordesigntype := range stage.SectorDesignTypes {
		_ = sectordesigntype
		for _, _notepropertytype := range sectordesigntype.Annotation {
			stage.SectorDesignType_Annotation_reverseMap[_notepropertytype] = sectordesigntype
		}
	}
	stage.SectorDesignType_Extension_reverseMap = make(map[*A_extension_33]*SectorDesignType)
	for sectordesigntype := range stage.SectorDesignTypes {
		_ = sectordesigntype
		for _, _a_extension_33 := range sectordesigntype.Extension {
			stage.SectorDesignType_Extension_reverseMap[_a_extension_33] = sectordesigntype
		}
	}

	// Compute reverse map for named struct StandardInstrumentArrivalTimeSliceType
	// insertion point per field
	stage.StandardInstrumentArrivalTimeSliceType_Availability_reverseMap = make(map[*ProcedureAvailabilityPropertyType]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _procedureavailabilitypropertytype := range standardinstrumentarrivaltimeslicetype.Availability {
			stage.StandardInstrumentArrivalTimeSliceType_Availability_reverseMap[_procedureavailabilitypropertytype] = standardinstrumentarrivaltimeslicetype
		}
	}
	stage.StandardInstrumentArrivalTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _airportheliportpropertytype := range standardinstrumentarrivaltimeslicetype.AirportHeliport {
			stage.StandardInstrumentArrivalTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = standardinstrumentarrivaltimeslicetype
		}
	}
	stage.StandardInstrumentArrivalTimeSliceType_AircraftCharacteristic_reverseMap = make(map[*AircraftCharacteristicPropertyType]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _aircraftcharacteristicpropertytype := range standardinstrumentarrivaltimeslicetype.AircraftCharacteristic {
			stage.StandardInstrumentArrivalTimeSliceType_AircraftCharacteristic_reverseMap[_aircraftcharacteristicpropertytype] = standardinstrumentarrivaltimeslicetype
		}
	}
	stage.StandardInstrumentArrivalTimeSliceType_FlightTransition_reverseMap = make(map[*ProcedureTransitionPropertyType]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _proceduretransitionpropertytype := range standardinstrumentarrivaltimeslicetype.FlightTransition {
			stage.StandardInstrumentArrivalTimeSliceType_FlightTransition_reverseMap[_proceduretransitionpropertytype] = standardinstrumentarrivaltimeslicetype
		}
	}
	stage.StandardInstrumentArrivalTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _notepropertytype := range standardinstrumentarrivaltimeslicetype.Annotation {
			stage.StandardInstrumentArrivalTimeSliceType_Annotation_reverseMap[_notepropertytype] = standardinstrumentarrivaltimeslicetype
		}
	}
	stage.StandardInstrumentArrivalTimeSliceType_Extension_reverseMap = make(map[*A_extension_29]*StandardInstrumentArrivalTimeSliceType)
	for standardinstrumentarrivaltimeslicetype := range stage.StandardInstrumentArrivalTimeSliceTypes {
		_ = standardinstrumentarrivaltimeslicetype
		for _, _a_extension_29 := range standardinstrumentarrivaltimeslicetype.Extension {
			stage.StandardInstrumentArrivalTimeSliceType_Extension_reverseMap[_a_extension_29] = standardinstrumentarrivaltimeslicetype
		}
	}

	// Compute reverse map for named struct StandardInstrumentArrivalType
	// insertion point per field
	stage.StandardInstrumentArrivalType_TimeSlice_reverseMap = make(map[*StandardInstrumentArrivalTimeSlicePropertyType]*StandardInstrumentArrivalType)
	for standardinstrumentarrivaltype := range stage.StandardInstrumentArrivalTypes {
		_ = standardinstrumentarrivaltype
		for _, _standardinstrumentarrivaltimeslicepropertytype := range standardinstrumentarrivaltype.TimeSlice {
			stage.StandardInstrumentArrivalType_TimeSlice_reverseMap[_standardinstrumentarrivaltimeslicepropertytype] = standardinstrumentarrivaltype
		}
	}

	// Compute reverse map for named struct StandardInstrumentDepartureTimeSliceType
	// insertion point per field
	stage.StandardInstrumentDepartureTimeSliceType_Availability_reverseMap = make(map[*ProcedureAvailabilityPropertyType]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _procedureavailabilitypropertytype := range standardinstrumentdeparturetimeslicetype.Availability {
			stage.StandardInstrumentDepartureTimeSliceType_Availability_reverseMap[_procedureavailabilitypropertytype] = standardinstrumentdeparturetimeslicetype
		}
	}
	stage.StandardInstrumentDepartureTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _airportheliportpropertytype := range standardinstrumentdeparturetimeslicetype.AirportHeliport {
			stage.StandardInstrumentDepartureTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = standardinstrumentdeparturetimeslicetype
		}
	}
	stage.StandardInstrumentDepartureTimeSliceType_AircraftCharacteristic_reverseMap = make(map[*AircraftCharacteristicPropertyType]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _aircraftcharacteristicpropertytype := range standardinstrumentdeparturetimeslicetype.AircraftCharacteristic {
			stage.StandardInstrumentDepartureTimeSliceType_AircraftCharacteristic_reverseMap[_aircraftcharacteristicpropertytype] = standardinstrumentdeparturetimeslicetype
		}
	}
	stage.StandardInstrumentDepartureTimeSliceType_FlightTransition_reverseMap = make(map[*ProcedureTransitionPropertyType]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _proceduretransitionpropertytype := range standardinstrumentdeparturetimeslicetype.FlightTransition {
			stage.StandardInstrumentDepartureTimeSliceType_FlightTransition_reverseMap[_proceduretransitionpropertytype] = standardinstrumentdeparturetimeslicetype
		}
	}
	stage.StandardInstrumentDepartureTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _notepropertytype := range standardinstrumentdeparturetimeslicetype.Annotation {
			stage.StandardInstrumentDepartureTimeSliceType_Annotation_reverseMap[_notepropertytype] = standardinstrumentdeparturetimeslicetype
		}
	}
	stage.StandardInstrumentDepartureTimeSliceType_Extension_reverseMap = make(map[*A_extension_34]*StandardInstrumentDepartureTimeSliceType)
	for standardinstrumentdeparturetimeslicetype := range stage.StandardInstrumentDepartureTimeSliceTypes {
		_ = standardinstrumentdeparturetimeslicetype
		for _, _a_extension_34 := range standardinstrumentdeparturetimeslicetype.Extension {
			stage.StandardInstrumentDepartureTimeSliceType_Extension_reverseMap[_a_extension_34] = standardinstrumentdeparturetimeslicetype
		}
	}

	// Compute reverse map for named struct StandardInstrumentDepartureType
	// insertion point per field
	stage.StandardInstrumentDepartureType_TimeSlice_reverseMap = make(map[*StandardInstrumentDepartureTimeSlicePropertyType]*StandardInstrumentDepartureType)
	for standardinstrumentdeparturetype := range stage.StandardInstrumentDepartureTypes {
		_ = standardinstrumentdeparturetype
		for _, _standardinstrumentdeparturetimeslicepropertytype := range standardinstrumentdeparturetype.TimeSlice {
			stage.StandardInstrumentDepartureType_TimeSlice_reverseMap[_standardinstrumentdeparturetimeslicepropertytype] = standardinstrumentdeparturetype
		}
	}

	// Compute reverse map for named struct SurfaceType
	// insertion point per field
	stage.SurfaceType_Annotation_reverseMap = make(map[*NotePropertyType]*SurfaceType)
	for surfacetype := range stage.SurfaceTypes {
		_ = surfacetype
		for _, _notepropertytype := range surfacetype.Annotation {
			stage.SurfaceType_Annotation_reverseMap[_notepropertytype] = surfacetype
		}
	}

	// Compute reverse map for named struct TerminalArrivalAreaSectorType
	// insertion point per field
	stage.TerminalArrivalAreaSectorType_SignificantObstacle_reverseMap = make(map[*ObstructionPropertyType]*TerminalArrivalAreaSectorType)
	for terminalarrivalareasectortype := range stage.TerminalArrivalAreaSectorTypes {
		_ = terminalarrivalareasectortype
		for _, _obstructionpropertytype := range terminalarrivalareasectortype.SignificantObstacle {
			stage.TerminalArrivalAreaSectorType_SignificantObstacle_reverseMap[_obstructionpropertytype] = terminalarrivalareasectortype
		}
	}
	stage.TerminalArrivalAreaSectorType_Annotation_reverseMap = make(map[*NotePropertyType]*TerminalArrivalAreaSectorType)
	for terminalarrivalareasectortype := range stage.TerminalArrivalAreaSectorTypes {
		_ = terminalarrivalareasectortype
		for _, _notepropertytype := range terminalarrivalareasectortype.Annotation {
			stage.TerminalArrivalAreaSectorType_Annotation_reverseMap[_notepropertytype] = terminalarrivalareasectortype
		}
	}
	stage.TerminalArrivalAreaSectorType_Extension_reverseMap = make(map[*A_extension_22]*TerminalArrivalAreaSectorType)
	for terminalarrivalareasectortype := range stage.TerminalArrivalAreaSectorTypes {
		_ = terminalarrivalareasectortype
		for _, _a_extension_22 := range terminalarrivalareasectortype.Extension {
			stage.TerminalArrivalAreaSectorType_Extension_reverseMap[_a_extension_22] = terminalarrivalareasectortype
		}
	}

	// Compute reverse map for named struct TerminalArrivalAreaTimeSliceType
	// insertion point per field
	stage.TerminalArrivalAreaTimeSliceType_Sector_reverseMap = make(map[*TerminalArrivalAreaSectorPropertyType]*TerminalArrivalAreaTimeSliceType)
	for terminalarrivalareatimeslicetype := range stage.TerminalArrivalAreaTimeSliceTypes {
		_ = terminalarrivalareatimeslicetype
		for _, _terminalarrivalareasectorpropertytype := range terminalarrivalareatimeslicetype.Sector {
			stage.TerminalArrivalAreaTimeSliceType_Sector_reverseMap[_terminalarrivalareasectorpropertytype] = terminalarrivalareatimeslicetype
		}
	}
	stage.TerminalArrivalAreaTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*TerminalArrivalAreaTimeSliceType)
	for terminalarrivalareatimeslicetype := range stage.TerminalArrivalAreaTimeSliceTypes {
		_ = terminalarrivalareatimeslicetype
		for _, _notepropertytype := range terminalarrivalareatimeslicetype.Annotation {
			stage.TerminalArrivalAreaTimeSliceType_Annotation_reverseMap[_notepropertytype] = terminalarrivalareatimeslicetype
		}
	}
	stage.TerminalArrivalAreaTimeSliceType_Extension_reverseMap = make(map[*A_extension_21]*TerminalArrivalAreaTimeSliceType)
	for terminalarrivalareatimeslicetype := range stage.TerminalArrivalAreaTimeSliceTypes {
		_ = terminalarrivalareatimeslicetype
		for _, _a_extension_21 := range terminalarrivalareatimeslicetype.Extension {
			stage.TerminalArrivalAreaTimeSliceType_Extension_reverseMap[_a_extension_21] = terminalarrivalareatimeslicetype
		}
	}

	// Compute reverse map for named struct TerminalArrivalAreaType
	// insertion point per field
	stage.TerminalArrivalAreaType_TimeSlice_reverseMap = make(map[*TerminalArrivalAreaTimeSlicePropertyType]*TerminalArrivalAreaType)
	for terminalarrivalareatype := range stage.TerminalArrivalAreaTypes {
		_ = terminalarrivalareatype
		for _, _terminalarrivalareatimeslicepropertytype := range terminalarrivalareatype.TimeSlice {
			stage.TerminalArrivalAreaType_TimeSlice_reverseMap[_terminalarrivalareatimeslicepropertytype] = terminalarrivalareatype
		}
	}

	// Compute reverse map for named struct TerminalSegmentPointType
	// insertion point per field
	stage.TerminalSegmentPointType_FacilityMakeup_reverseMap = make(map[*PointReferencePropertyType]*TerminalSegmentPointType)
	for terminalsegmentpointtype := range stage.TerminalSegmentPointTypes {
		_ = terminalsegmentpointtype
		for _, _pointreferencepropertytype := range terminalsegmentpointtype.FacilityMakeup {
			stage.TerminalSegmentPointType_FacilityMakeup_reverseMap[_pointreferencepropertytype] = terminalsegmentpointtype
		}
	}
	stage.TerminalSegmentPointType_Annotation_reverseMap = make(map[*NotePropertyType]*TerminalSegmentPointType)
	for terminalsegmentpointtype := range stage.TerminalSegmentPointTypes {
		_ = terminalsegmentpointtype
		for _, _notepropertytype := range terminalsegmentpointtype.Annotation {
			stage.TerminalSegmentPointType_Annotation_reverseMap[_notepropertytype] = terminalsegmentpointtype
		}
	}
	stage.TerminalSegmentPointType_Extension_reverseMap = make(map[*A_extension_11]*TerminalSegmentPointType)
	for terminalsegmentpointtype := range stage.TerminalSegmentPointTypes {
		_ = terminalsegmentpointtype
		for _, _a_extension_11 := range terminalsegmentpointtype.Extension {
			stage.TerminalSegmentPointType_Extension_reverseMap[_a_extension_11] = terminalsegmentpointtype
		}
	}

	// Compute reverse map for named struct TimesheetType
	// insertion point per field
	stage.TimesheetType_Annotation_reverseMap = make(map[*NotePropertyType]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _notepropertytype := range timesheettype.Annotation {
			stage.TimesheetType_Annotation_reverseMap[_notepropertytype] = timesheettype
		}
	}
	stage.TimesheetType_Extension_reverseMap = make(map[*A_extension_53]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _a_extension_53 := range timesheettype.Extension {
			stage.TimesheetType_Extension_reverseMap[_a_extension_53] = timesheettype
		}
	}

	// Compute reverse map for named struct UnplannedHoldingTimeSliceType
	// insertion point per field
	stage.UnplannedHoldingTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*UnplannedHoldingTimeSliceType)
	for unplannedholdingtimeslicetype := range stage.UnplannedHoldingTimeSliceTypes {
		_ = unplannedholdingtimeslicetype
		for _, _notepropertytype := range unplannedholdingtimeslicetype.Annotation {
			stage.UnplannedHoldingTimeSliceType_Annotation_reverseMap[_notepropertytype] = unplannedholdingtimeslicetype
		}
	}
	stage.UnplannedHoldingTimeSliceType_Extension_reverseMap = make(map[*A_extension_7]*UnplannedHoldingTimeSliceType)
	for unplannedholdingtimeslicetype := range stage.UnplannedHoldingTimeSliceTypes {
		_ = unplannedholdingtimeslicetype
		for _, _a_extension_7 := range unplannedholdingtimeslicetype.Extension {
			stage.UnplannedHoldingTimeSliceType_Extension_reverseMap[_a_extension_7] = unplannedholdingtimeslicetype
		}
	}

	// Compute reverse map for named struct UnplannedHoldingType
	// insertion point per field
	stage.UnplannedHoldingType_TimeSlice_reverseMap = make(map[*UnplannedHoldingTimeSlicePropertyType]*UnplannedHoldingType)
	for unplannedholdingtype := range stage.UnplannedHoldingTypes {
		_ = unplannedholdingtype
		for _, _unplannedholdingtimeslicepropertytype := range unplannedholdingtype.TimeSlice {
			stage.UnplannedHoldingType_TimeSlice_reverseMap[_unplannedholdingtimeslicepropertytype] = unplannedholdingtype
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.A_extensions)

	res = __gong__appendInstances(res, stage.A_extension_1s)

	res = __gong__appendInstances(res, stage.A_extension_10s)

	res = __gong__appendInstances(res, stage.A_extension_11s)

	res = __gong__appendInstances(res, stage.A_extension_12s)

	res = __gong__appendInstances(res, stage.A_extension_13s)

	res = __gong__appendInstances(res, stage.A_extension_14s)

	res = __gong__appendInstances(res, stage.A_extension_15s)

	res = __gong__appendInstances(res, stage.A_extension_16s)

	res = __gong__appendInstances(res, stage.A_extension_17s)

	res = __gong__appendInstances(res, stage.A_extension_18s)

	res = __gong__appendInstances(res, stage.A_extension_19s)

	res = __gong__appendInstances(res, stage.A_extension_2s)

	res = __gong__appendInstances(res, stage.A_extension_20s)

	res = __gong__appendInstances(res, stage.A_extension_21s)

	res = __gong__appendInstances(res, stage.A_extension_22s)

	res = __gong__appendInstances(res, stage.A_extension_23s)

	res = __gong__appendInstances(res, stage.A_extension_24s)

	res = __gong__appendInstances(res, stage.A_extension_25s)

	res = __gong__appendInstances(res, stage.A_extension_26s)

	res = __gong__appendInstances(res, stage.A_extension_27s)

	res = __gong__appendInstances(res, stage.A_extension_28s)

	res = __gong__appendInstances(res, stage.A_extension_29s)

	res = __gong__appendInstances(res, stage.A_extension_3s)

	res = __gong__appendInstances(res, stage.A_extension_30s)

	res = __gong__appendInstances(res, stage.A_extension_31s)

	res = __gong__appendInstances(res, stage.A_extension_32s)

	res = __gong__appendInstances(res, stage.A_extension_33s)

	res = __gong__appendInstances(res, stage.A_extension_34s)

	res = __gong__appendInstances(res, stage.A_extension_35s)

	res = __gong__appendInstances(res, stage.A_extension_36s)

	res = __gong__appendInstances(res, stage.A_extension_37s)

	res = __gong__appendInstances(res, stage.A_extension_38s)

	res = __gong__appendInstances(res, stage.A_extension_39s)

	res = __gong__appendInstances(res, stage.A_extension_4s)

	res = __gong__appendInstances(res, stage.A_extension_40s)

	res = __gong__appendInstances(res, stage.A_extension_41s)

	res = __gong__appendInstances(res, stage.A_extension_42s)

	res = __gong__appendInstances(res, stage.A_extension_43s)

	res = __gong__appendInstances(res, stage.A_extension_44s)

	res = __gong__appendInstances(res, stage.A_extension_45s)

	res = __gong__appendInstances(res, stage.A_extension_46s)

	res = __gong__appendInstances(res, stage.A_extension_47s)

	res = __gong__appendInstances(res, stage.A_extension_48s)

	res = __gong__appendInstances(res, stage.A_extension_49s)

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_50s)

	res = __gong__appendInstances(res, stage.A_extension_51s)

	res = __gong__appendInstances(res, stage.A_extension_52s)

	res = __gong__appendInstances(res, stage.A_extension_53s)

	res = __gong__appendInstances(res, stage.A_extension_54s)

	res = __gong__appendInstances(res, stage.A_extension_55s)

	res = __gong__appendInstances(res, stage.A_extension_56s)

	res = __gong__appendInstances(res, stage.A_extension_57s)

	res = __gong__appendInstances(res, stage.A_extension_58s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractApproachLegTypes)

	res = __gong__appendInstances(res, stage.AbstractProcedureTypes)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AbstractSegmentLegTypes)

	res = __gong__appendInstances(res, stage.AbstractSegmentPointTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingPointTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicPropertyTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicTypes)

	res = __gong__appendInstances(res, stage.AirportHeliportPropertyTypes)

	res = __gong__appendInstances(res, stage.AltimeterSourcePropertyTypes)

	res = __gong__appendInstances(res, stage.AltitudeAdjustmentPropertyTypes)

	res = __gong__appendInstances(res, stage.AltitudeAdjustmentTypes)

	res = __gong__appendInstances(res, stage.AngleIndicationPropertyTypes)

	res = __gong__appendInstances(res, stage.AngleUsePropertyTypes)

	res = __gong__appendInstances(res, stage.AngleUseTypes)

	res = __gong__appendInstances(res, stage.ApproachAltitudeTablePropertyTypes)

	res = __gong__appendInstances(res, stage.ApproachAltitudeTableTypes)

	res = __gong__appendInstances(res, stage.ApproachConditionPropertyTypes)

	res = __gong__appendInstances(res, stage.ApproachConditionTypes)

	res = __gong__appendInstances(res, stage.ApproachDistanceTablePropertyTypes)

	res = __gong__appendInstances(res, stage.ApproachDistanceTableTypes)

	res = __gong__appendInstances(res, stage.ApproachTimingTablePropertyTypes)

	res = __gong__appendInstances(res, stage.ApproachTimingTableTypes)

	res = __gong__appendInstances(res, stage.ArrivalFeederLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ArrivalFeederLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.ArrivalFeederLegTypes)

	res = __gong__appendInstances(res, stage.ArrivalLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ArrivalLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.ArrivalLegTypes)

	res = __gong__appendInstances(res, stage.CircleSectorPropertyTypes)

	res = __gong__appendInstances(res, stage.CircleSectorTypes)

	res = __gong__appendInstances(res, stage.CirclingAreaTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.CirclingAreaTimeSliceTypes)

	res = __gong__appendInstances(res, stage.CirclingAreaTypes)

	res = __gong__appendInstances(res, stage.CirclingRestrictionPropertyTypes)

	res = __gong__appendInstances(res, stage.CirclingRestrictionTypes)

	res = __gong__appendInstances(res, stage.CurvePropertyTypes)

	res = __gong__appendInstances(res, stage.CurveTypes)

	res = __gong__appendInstances(res, stage.DMEPropertyTypes)

	res = __gong__appendInstances(res, stage.DepartureArrivalConditionPropertyTypes)

	res = __gong__appendInstances(res, stage.DepartureArrivalConditionTypes)

	res = __gong__appendInstances(res, stage.DepartureLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.DepartureLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.DepartureLegTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.DistanceIndicationPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedCurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointTypes)

	res = __gong__appendInstances(res, stage.ElevatedSurfaceTypes)

	res = __gong__appendInstances(res, stage.EnRouteSegmentPointTypes)

	res = __gong__appendInstances(res, stage.EquipmentUnavailableAdjustmentColumnPropertyTypes)

	res = __gong__appendInstances(res, stage.EquipmentUnavailableAdjustmentColumnTypes)

	res = __gong__appendInstances(res, stage.EquipmentUnavailableAdjustmentPropertyTypes)

	res = __gong__appendInstances(res, stage.EquipmentUnavailableAdjustmentTypes)

	res = __gong__appendInstances(res, stage.FASDataBlockPropertyTypes)

	res = __gong__appendInstances(res, stage.FASDataBlockTypes)

	res = __gong__appendInstances(res, stage.FinalLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.FinalLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.FinalLegTypes)

	res = __gong__appendInstances(res, stage.FinalProfilePropertyTypes)

	res = __gong__appendInstances(res, stage.FinalProfileTypes)

	res = __gong__appendInstances(res, stage.HoldingAssessmentTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingAssessmentTimeSliceTypes)

	res = __gong__appendInstances(res, stage.HoldingAssessmentTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternDistancePropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternDistanceTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternDurationPropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternDurationTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternPropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternTimeSliceTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternTypes)

	res = __gong__appendInstances(res, stage.HoldingUsePropertyTypes)

	res = __gong__appendInstances(res, stage.HoldingUseTypes)

	res = __gong__appendInstances(res, stage.InitialLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.InitialLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.InitialLegTypes)

	res = __gong__appendInstances(res, stage.InstrumentApproachProcedurePropertyTypes)

	res = __gong__appendInstances(res, stage.InstrumentApproachProcedureTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.InstrumentApproachProcedureTimeSliceTypes)

	res = __gong__appendInstances(res, stage.InstrumentApproachProcedureTypes)

	res = __gong__appendInstances(res, stage.IntermediateLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.IntermediateLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.IntermediateLegTypes)

	res = __gong__appendInstances(res, stage.LandingTakeoffAreaCollectionPropertyTypes)

	res = __gong__appendInstances(res, stage.LandingTakeoffAreaCollectionTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.MinimaPropertyTypes)

	res = __gong__appendInstances(res, stage.MinimaTypes)

	res = __gong__appendInstances(res, stage.MissedApproachGroupPropertyTypes)

	res = __gong__appendInstances(res, stage.MissedApproachGroupTypes)

	res = __gong__appendInstances(res, stage.MissedApproachLegTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.MissedApproachLegTimeSliceTypes)

	res = __gong__appendInstances(res, stage.MissedApproachLegTypes)

	res = __gong__appendInstances(res, stage.NavaidPropertyTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaRestrictionTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaRestrictionTimeSliceTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaRestrictionTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaSectorPropertyTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaSectorTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaTimeSliceTypes)

	res = __gong__appendInstances(res, stage.NavigationAreaTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.ObstacleAssessmentAreaPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstacleAssessmentAreaTypes)

	res = __gong__appendInstances(res, stage.ObstaclePlacementPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstaclePlacementTypes)

	res = __gong__appendInstances(res, stage.ObstructionPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstructionTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.PointPropertyTypes)

	res = __gong__appendInstances(res, stage.PointReferencePropertyTypes)

	res = __gong__appendInstances(res, stage.PointReferenceTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.ProcedureAvailabilityPropertyTypes)

	res = __gong__appendInstances(res, stage.ProcedureAvailabilityTypes)

	res = __gong__appendInstances(res, stage.ProcedureDMETimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ProcedureDMETimeSliceTypes)

	res = __gong__appendInstances(res, stage.ProcedureDMETypes)

	res = __gong__appendInstances(res, stage.ProcedurePropertyTypes)

	res = __gong__appendInstances(res, stage.ProcedureTransitionLegPropertyTypes)

	res = __gong__appendInstances(res, stage.ProcedureTransitionLegTypes)

	res = __gong__appendInstances(res, stage.ProcedureTransitionPropertyTypes)

	res = __gong__appendInstances(res, stage.ProcedureTransitionTypes)

	res = __gong__appendInstances(res, stage.RadarSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.RadioFrequencyAreaPropertyTypes)

	res = __gong__appendInstances(res, stage.RunwayCentrelinePointPropertyTypes)

	res = __gong__appendInstances(res, stage.RunwayDirectionPropertyTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaPropertyTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaSectorPropertyTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaSectorTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SafeAltitudeAreaTypes)

	res = __gong__appendInstances(res, stage.SectorDesignPropertyTypes)

	res = __gong__appendInstances(res, stage.SectorDesignTypes)

	res = __gong__appendInstances(res, stage.SegmentLegPropertyTypes)

	res = __gong__appendInstances(res, stage.SegmentPointPropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentArrivalPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentArrivalTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentArrivalTimeSliceTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentArrivalTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentDeparturePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentDepartureTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentDepartureTimeSliceTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentDepartureTypes)

	res = __gong__appendInstances(res, stage.SurfacePropertyTypes)

	res = __gong__appendInstances(res, stage.SurfaceTypes)

	res = __gong__appendInstances(res, stage.TerminalArrivalAreaSectorPropertyTypes)

	res = __gong__appendInstances(res, stage.TerminalArrivalAreaSectorTypes)

	res = __gong__appendInstances(res, stage.TerminalArrivalAreaTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.TerminalArrivalAreaTimeSliceTypes)

	res = __gong__appendInstances(res, stage.TerminalArrivalAreaTypes)

	res = __gong__appendInstances(res, stage.TerminalSegmentPointPropertyTypes)

	res = __gong__appendInstances(res, stage.TerminalSegmentPointTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	res = __gong__appendInstances(res, stage.TouchDownLiftOffPropertyTypes)

	res = __gong__appendInstances(res, stage.UnplannedHoldingPropertyTypes)

	res = __gong__appendInstances(res, stage.UnplannedHoldingTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.UnplannedHoldingTimeSliceTypes)

	res = __gong__appendInstances(res, stage.UnplannedHoldingTypes)

	res = __gong__appendInstances(res, stage.VerticalStructurePropertyTypes)

	return
}

// insertion point per named struct
func (a_extension *A_extension) GongCopy() GongstructIF {
	newInstance := new(A_extension)
	a_extension.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_1 *A_extension_1) GongCopy() GongstructIF {
	newInstance := new(A_extension_1)
	a_extension_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_10 *A_extension_10) GongCopy() GongstructIF {
	newInstance := new(A_extension_10)
	a_extension_10.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_11 *A_extension_11) GongCopy() GongstructIF {
	newInstance := new(A_extension_11)
	a_extension_11.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_12 *A_extension_12) GongCopy() GongstructIF {
	newInstance := new(A_extension_12)
	a_extension_12.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_13 *A_extension_13) GongCopy() GongstructIF {
	newInstance := new(A_extension_13)
	a_extension_13.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_14 *A_extension_14) GongCopy() GongstructIF {
	newInstance := new(A_extension_14)
	a_extension_14.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_15 *A_extension_15) GongCopy() GongstructIF {
	newInstance := new(A_extension_15)
	a_extension_15.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_16 *A_extension_16) GongCopy() GongstructIF {
	newInstance := new(A_extension_16)
	a_extension_16.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_17 *A_extension_17) GongCopy() GongstructIF {
	newInstance := new(A_extension_17)
	a_extension_17.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_18 *A_extension_18) GongCopy() GongstructIF {
	newInstance := new(A_extension_18)
	a_extension_18.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_19 *A_extension_19) GongCopy() GongstructIF {
	newInstance := new(A_extension_19)
	a_extension_19.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_2 *A_extension_2) GongCopy() GongstructIF {
	newInstance := new(A_extension_2)
	a_extension_2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_20 *A_extension_20) GongCopy() GongstructIF {
	newInstance := new(A_extension_20)
	a_extension_20.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_21 *A_extension_21) GongCopy() GongstructIF {
	newInstance := new(A_extension_21)
	a_extension_21.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_22 *A_extension_22) GongCopy() GongstructIF {
	newInstance := new(A_extension_22)
	a_extension_22.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_23 *A_extension_23) GongCopy() GongstructIF {
	newInstance := new(A_extension_23)
	a_extension_23.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_24 *A_extension_24) GongCopy() GongstructIF {
	newInstance := new(A_extension_24)
	a_extension_24.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_25 *A_extension_25) GongCopy() GongstructIF {
	newInstance := new(A_extension_25)
	a_extension_25.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_26 *A_extension_26) GongCopy() GongstructIF {
	newInstance := new(A_extension_26)
	a_extension_26.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_27 *A_extension_27) GongCopy() GongstructIF {
	newInstance := new(A_extension_27)
	a_extension_27.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_28 *A_extension_28) GongCopy() GongstructIF {
	newInstance := new(A_extension_28)
	a_extension_28.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_29 *A_extension_29) GongCopy() GongstructIF {
	newInstance := new(A_extension_29)
	a_extension_29.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_3 *A_extension_3) GongCopy() GongstructIF {
	newInstance := new(A_extension_3)
	a_extension_3.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_30 *A_extension_30) GongCopy() GongstructIF {
	newInstance := new(A_extension_30)
	a_extension_30.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_31 *A_extension_31) GongCopy() GongstructIF {
	newInstance := new(A_extension_31)
	a_extension_31.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_32 *A_extension_32) GongCopy() GongstructIF {
	newInstance := new(A_extension_32)
	a_extension_32.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_33 *A_extension_33) GongCopy() GongstructIF {
	newInstance := new(A_extension_33)
	a_extension_33.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_34 *A_extension_34) GongCopy() GongstructIF {
	newInstance := new(A_extension_34)
	a_extension_34.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_35 *A_extension_35) GongCopy() GongstructIF {
	newInstance := new(A_extension_35)
	a_extension_35.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_36 *A_extension_36) GongCopy() GongstructIF {
	newInstance := new(A_extension_36)
	a_extension_36.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_37 *A_extension_37) GongCopy() GongstructIF {
	newInstance := new(A_extension_37)
	a_extension_37.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_38 *A_extension_38) GongCopy() GongstructIF {
	newInstance := new(A_extension_38)
	a_extension_38.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_39 *A_extension_39) GongCopy() GongstructIF {
	newInstance := new(A_extension_39)
	a_extension_39.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_4 *A_extension_4) GongCopy() GongstructIF {
	newInstance := new(A_extension_4)
	a_extension_4.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_40 *A_extension_40) GongCopy() GongstructIF {
	newInstance := new(A_extension_40)
	a_extension_40.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_41 *A_extension_41) GongCopy() GongstructIF {
	newInstance := new(A_extension_41)
	a_extension_41.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_42 *A_extension_42) GongCopy() GongstructIF {
	newInstance := new(A_extension_42)
	a_extension_42.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_43 *A_extension_43) GongCopy() GongstructIF {
	newInstance := new(A_extension_43)
	a_extension_43.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_44 *A_extension_44) GongCopy() GongstructIF {
	newInstance := new(A_extension_44)
	a_extension_44.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_45 *A_extension_45) GongCopy() GongstructIF {
	newInstance := new(A_extension_45)
	a_extension_45.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_46 *A_extension_46) GongCopy() GongstructIF {
	newInstance := new(A_extension_46)
	a_extension_46.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_47 *A_extension_47) GongCopy() GongstructIF {
	newInstance := new(A_extension_47)
	a_extension_47.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_48 *A_extension_48) GongCopy() GongstructIF {
	newInstance := new(A_extension_48)
	a_extension_48.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_49 *A_extension_49) GongCopy() GongstructIF {
	newInstance := new(A_extension_49)
	a_extension_49.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_5 *A_extension_5) GongCopy() GongstructIF {
	newInstance := new(A_extension_5)
	a_extension_5.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_50 *A_extension_50) GongCopy() GongstructIF {
	newInstance := new(A_extension_50)
	a_extension_50.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_51 *A_extension_51) GongCopy() GongstructIF {
	newInstance := new(A_extension_51)
	a_extension_51.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_52 *A_extension_52) GongCopy() GongstructIF {
	newInstance := new(A_extension_52)
	a_extension_52.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_53 *A_extension_53) GongCopy() GongstructIF {
	newInstance := new(A_extension_53)
	a_extension_53.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_54 *A_extension_54) GongCopy() GongstructIF {
	newInstance := new(A_extension_54)
	a_extension_54.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_55 *A_extension_55) GongCopy() GongstructIF {
	newInstance := new(A_extension_55)
	a_extension_55.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_56 *A_extension_56) GongCopy() GongstructIF {
	newInstance := new(A_extension_56)
	a_extension_56.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_57 *A_extension_57) GongCopy() GongstructIF {
	newInstance := new(A_extension_57)
	a_extension_57.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_58 *A_extension_58) GongCopy() GongstructIF {
	newInstance := new(A_extension_58)
	a_extension_58.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_6 *A_extension_6) GongCopy() GongstructIF {
	newInstance := new(A_extension_6)
	a_extension_6.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_7 *A_extension_7) GongCopy() GongstructIF {
	newInstance := new(A_extension_7)
	a_extension_7.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_8 *A_extension_8) GongCopy() GongstructIF {
	newInstance := new(A_extension_8)
	a_extension_8.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_9 *A_extension_9) GongCopy() GongstructIF {
	newInstance := new(A_extension_9)
	a_extension_9.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractapproachlegtype *AbstractApproachLegType) GongCopy() GongstructIF {
	newInstance := new(AbstractApproachLegType)
	abstractapproachlegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractproceduretype *AbstractProcedureType) GongCopy() GongstructIF {
	newInstance := new(AbstractProcedureType)
	abstractproceduretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongCopy() GongstructIF {
	newInstance := new(AbstractSegmentLegType)
	abstractsegmentlegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongCopy() GongstructIF {
	newInstance := new(AbstractSegmentPointType)
	abstractsegmentpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingPointType)
	aerialrefuellingpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongCopy() GongstructIF {
	newInstance := new(AircraftCharacteristicPropertyType)
	aircraftcharacteristicpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongCopy() GongstructIF {
	newInstance := new(AircraftCharacteristicType)
	aircraftcharacteristictype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportHeliportPropertyType)
	airportheliportpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongCopy() GongstructIF {
	newInstance := new(AltimeterSourcePropertyType)
	altimetersourcepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(AltitudeAdjustmentPropertyType)
	altitudeadjustmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongCopy() GongstructIF {
	newInstance := new(AltitudeAdjustmentType)
	altitudeadjustmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongCopy() GongstructIF {
	newInstance := new(AngleIndicationPropertyType)
	angleindicationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleusepropertytype *AngleUsePropertyType) GongCopy() GongstructIF {
	newInstance := new(AngleUsePropertyType)
	angleusepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleusetype *AngleUseType) GongCopy() GongstructIF {
	newInstance := new(AngleUseType)
	angleusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongCopy() GongstructIF {
	newInstance := new(ApproachAltitudeTablePropertyType)
	approachaltitudetablepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongCopy() GongstructIF {
	newInstance := new(ApproachAltitudeTableType)
	approachaltitudetabletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongCopy() GongstructIF {
	newInstance := new(ApproachConditionPropertyType)
	approachconditionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachconditiontype *ApproachConditionType) GongCopy() GongstructIF {
	newInstance := new(ApproachConditionType)
	approachconditiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongCopy() GongstructIF {
	newInstance := new(ApproachDistanceTablePropertyType)
	approachdistancetablepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachdistancetabletype *ApproachDistanceTableType) GongCopy() GongstructIF {
	newInstance := new(ApproachDistanceTableType)
	approachdistancetabletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongCopy() GongstructIF {
	newInstance := new(ApproachTimingTablePropertyType)
	approachtimingtablepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (approachtimingtabletype *ApproachTimingTableType) GongCopy() GongstructIF {
	newInstance := new(ApproachTimingTableType)
	approachtimingtabletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ArrivalFeederLegTimeSlicePropertyType)
	arrivalfeederlegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ArrivalFeederLegTimeSliceType)
	arrivalfeederlegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongCopy() GongstructIF {
	newInstance := new(ArrivalFeederLegType)
	arrivalfeederlegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ArrivalLegTimeSlicePropertyType)
	arrivallegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ArrivalLegTimeSliceType)
	arrivallegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrivallegtype *ArrivalLegType) GongCopy() GongstructIF {
	newInstance := new(ArrivalLegType)
	arrivallegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongCopy() GongstructIF {
	newInstance := new(CircleSectorPropertyType)
	circlesectorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlesectortype *CircleSectorType) GongCopy() GongstructIF {
	newInstance := new(CircleSectorType)
	circlesectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(CirclingAreaTimeSlicePropertyType)
	circlingareatimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(CirclingAreaTimeSliceType)
	circlingareatimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlingareatype *CirclingAreaType) GongCopy() GongstructIF {
	newInstance := new(CirclingAreaType)
	circlingareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongCopy() GongstructIF {
	newInstance := new(CirclingRestrictionPropertyType)
	circlingrestrictionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlingrestrictiontype *CirclingRestrictionType) GongCopy() GongstructIF {
	newInstance := new(CirclingRestrictionType)
	circlingrestrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curvepropertytype *CurvePropertyType) GongCopy() GongstructIF {
	newInstance := new(CurvePropertyType)
	curvepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curvetype *CurveType) GongCopy() GongstructIF {
	newInstance := new(CurveType)
	curvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dmepropertytype *DMEPropertyType) GongCopy() GongstructIF {
	newInstance := new(DMEPropertyType)
	dmepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongCopy() GongstructIF {
	newInstance := new(DepartureArrivalConditionPropertyType)
	departurearrivalconditionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongCopy() GongstructIF {
	newInstance := new(DepartureArrivalConditionType)
	departurearrivalconditiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(DepartureLegTimeSlicePropertyType)
	departurelegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(DepartureLegTimeSliceType)
	departurelegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (departurelegtype *DepartureLegType) GongCopy() GongstructIF {
	newInstance := new(DepartureLegType)
	departurelegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointPropertyType)
	designatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongCopy() GongstructIF {
	newInstance := new(DistanceIndicationPropertyType)
	distanceindicationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedcurvetype *ElevatedCurveType) GongCopy() GongstructIF {
	newInstance := new(ElevatedCurveType)
	elevatedcurvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointtype *ElevatedPointType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointType)
	elevatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongCopy() GongstructIF {
	newInstance := new(ElevatedSurfaceType)
	elevatedsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongCopy() GongstructIF {
	newInstance := new(EnRouteSegmentPointType)
	enroutesegmentpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongCopy() GongstructIF {
	newInstance := new(EquipmentUnavailableAdjustmentColumnPropertyType)
	equipmentunavailableadjustmentcolumnpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongCopy() GongstructIF {
	newInstance := new(EquipmentUnavailableAdjustmentColumnType)
	equipmentunavailableadjustmentcolumntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(EquipmentUnavailableAdjustmentPropertyType)
	equipmentunavailableadjustmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongCopy() GongstructIF {
	newInstance := new(EquipmentUnavailableAdjustmentType)
	equipmentunavailableadjustmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongCopy() GongstructIF {
	newInstance := new(FASDataBlockPropertyType)
	fasdatablockpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fasdatablocktype *FASDataBlockType) GongCopy() GongstructIF {
	newInstance := new(FASDataBlockType)
	fasdatablocktype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(FinalLegTimeSlicePropertyType)
	finallegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(FinalLegTimeSliceType)
	finallegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (finallegtype *FinalLegType) GongCopy() GongstructIF {
	newInstance := new(FinalLegType)
	finallegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongCopy() GongstructIF {
	newInstance := new(FinalProfilePropertyType)
	finalprofilepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (finalprofiletype *FinalProfileType) GongCopy() GongstructIF {
	newInstance := new(FinalProfileType)
	finalprofiletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingAssessmentTimeSlicePropertyType)
	holdingassessmenttimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(HoldingAssessmentTimeSliceType)
	holdingassessmenttimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingassessmenttype *HoldingAssessmentType) GongCopy() GongstructIF {
	newInstance := new(HoldingAssessmentType)
	holdingassessmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternDistancePropertyType)
	holdingpatterndistancepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternDistanceType)
	holdingpatterndistancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternDurationPropertyType)
	holdingpatterndurationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternDurationType)
	holdingpatterndurationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternPropertyType)
	holdingpatternpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternTimeSlicePropertyType)
	holdingpatterntimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternTimeSliceType)
	holdingpatterntimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatterntype *HoldingPatternType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternType)
	holdingpatterntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingusepropertytype *HoldingUsePropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingUsePropertyType)
	holdingusepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingusetype *HoldingUseType) GongCopy() GongstructIF {
	newInstance := new(HoldingUseType)
	holdingusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(InitialLegTimeSlicePropertyType)
	initiallegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(InitialLegTimeSliceType)
	initiallegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (initiallegtype *InitialLegType) GongCopy() GongstructIF {
	newInstance := new(InitialLegType)
	initiallegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongCopy() GongstructIF {
	newInstance := new(InstrumentApproachProcedurePropertyType)
	instrumentapproachprocedurepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(InstrumentApproachProcedureTimeSlicePropertyType)
	instrumentapproachproceduretimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(InstrumentApproachProcedureTimeSliceType)
	instrumentapproachproceduretimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongCopy() GongstructIF {
	newInstance := new(InstrumentApproachProcedureType)
	instrumentapproachproceduretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(IntermediateLegTimeSlicePropertyType)
	intermediatelegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(IntermediateLegTimeSliceType)
	intermediatelegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (intermediatelegtype *IntermediateLegType) GongCopy() GongstructIF {
	newInstance := new(IntermediateLegType)
	intermediatelegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongCopy() GongstructIF {
	newInstance := new(LandingTakeoffAreaCollectionPropertyType)
	landingtakeoffareacollectionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongCopy() GongstructIF {
	newInstance := new(LandingTakeoffAreaCollectionType)
	landingtakeoffareacollectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongCopy() GongstructIF {
	newInstance := new(LinguisticNotePropertyType)
	linguisticnotepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linguisticnotetype *LinguisticNoteType) GongCopy() GongstructIF {
	newInstance := new(LinguisticNoteType)
	linguisticnotetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (minimapropertytype *MinimaPropertyType) GongCopy() GongstructIF {
	newInstance := new(MinimaPropertyType)
	minimapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (minimatype *MinimaType) GongCopy() GongstructIF {
	newInstance := new(MinimaType)
	minimatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongCopy() GongstructIF {
	newInstance := new(MissedApproachGroupPropertyType)
	missedapproachgrouppropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (missedapproachgrouptype *MissedApproachGroupType) GongCopy() GongstructIF {
	newInstance := new(MissedApproachGroupType)
	missedapproachgrouptype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(MissedApproachLegTimeSlicePropertyType)
	missedapproachlegtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(MissedApproachLegTimeSliceType)
	missedapproachlegtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (missedapproachlegtype *MissedApproachLegType) GongCopy() GongstructIF {
	newInstance := new(MissedApproachLegType)
	missedapproachlegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidpropertytype *NavaidPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidPropertyType)
	navaidpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaRestrictionTimeSlicePropertyType)
	navigationarearestrictiontimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaRestrictionTimeSliceType)
	navigationarearestrictiontimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaRestrictionType)
	navigationarearestrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaSectorPropertyType)
	navigationareasectorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationareasectortype *NavigationAreaSectorType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaSectorType)
	navigationareasectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaTimeSlicePropertyType)
	navigationareatimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaTimeSliceType)
	navigationareatimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navigationareatype *NavigationAreaType) GongCopy() GongstructIF {
	newInstance := new(NavigationAreaType)
	navigationareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notepropertytype *NotePropertyType) GongCopy() GongstructIF {
	newInstance := new(NotePropertyType)
	notepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetype *NoteType) GongCopy() GongstructIF {
	newInstance := new(NoteType)
	notetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAssessmentAreaPropertyType)
	obstacleassessmentareapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAssessmentAreaType)
	obstacleassessmentareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstaclePlacementPropertyType)
	obstacleplacementpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleplacementtype *ObstaclePlacementType) GongCopy() GongstructIF {
	newInstance := new(ObstaclePlacementType)
	obstacleplacementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstructionpropertytype *ObstructionPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstructionPropertyType)
	obstructionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstructiontype *ObstructionType) GongCopy() GongstructIF {
	newInstance := new(ObstructionType)
	obstructiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityPropertyType)
	organisationauthoritypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointpropertytype *PointPropertyType) GongCopy() GongstructIF {
	newInstance := new(PointPropertyType)
	pointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointreferencepropertytype *PointReferencePropertyType) GongCopy() GongstructIF {
	newInstance := new(PointReferencePropertyType)
	pointreferencepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointreferencetype *PointReferenceType) GongCopy() GongstructIF {
	newInstance := new(PointReferenceType)
	pointreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointtype *PointType) GongCopy() GongstructIF {
	newInstance := new(PointType)
	pointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedureAvailabilityPropertyType)
	procedureavailabilitypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongCopy() GongstructIF {
	newInstance := new(ProcedureAvailabilityType)
	procedureavailabilitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedureDMETimeSlicePropertyType)
	proceduredmetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ProcedureDMETimeSliceType)
	proceduredmetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduredmetype *ProcedureDMEType) GongCopy() GongstructIF {
	newInstance := new(ProcedureDMEType)
	proceduredmetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (procedurepropertytype *ProcedurePropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedurePropertyType)
	procedurepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedureTransitionLegPropertyType)
	proceduretransitionlegpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongCopy() GongstructIF {
	newInstance := new(ProcedureTransitionLegType)
	proceduretransitionlegtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedureTransitionPropertyType)
	proceduretransitionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (proceduretransitiontype *ProcedureTransitionType) GongCopy() GongstructIF {
	newInstance := new(ProcedureTransitionType)
	proceduretransitiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarsystempropertytype *RadarSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadarSystemPropertyType)
	radarsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadioFrequencyAreaPropertyType)
	radiofrequencyareapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayCentrelinePointPropertyType)
	runwaycentrelinepointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayDirectionPropertyType)
	runwaydirectionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaPropertyType)
	safealtitudeareapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaSectorPropertyType)
	safealtitudeareasectorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaSectorType)
	safealtitudeareasectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaTimeSlicePropertyType)
	safealtitudeareatimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaTimeSliceType)
	safealtitudeareatimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongCopy() GongstructIF {
	newInstance := new(SafeAltitudeAreaType)
	safealtitudeareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongCopy() GongstructIF {
	newInstance := new(SectorDesignPropertyType)
	sectordesignpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sectordesigntype *SectorDesignType) GongCopy() GongstructIF {
	newInstance := new(SectorDesignType)
	sectordesigntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongCopy() GongstructIF {
	newInstance := new(SegmentLegPropertyType)
	segmentlegpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(SegmentPointPropertyType)
	segmentpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationSystemPropertyType)
	specialnavigationsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentArrivalPropertyType)
	standardinstrumentarrivalpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentArrivalTimeSlicePropertyType)
	standardinstrumentarrivaltimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentArrivalTimeSliceType)
	standardinstrumentarrivaltimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentArrivalType)
	standardinstrumentarrivaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentDeparturePropertyType)
	standardinstrumentdeparturepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentDepartureTimeSlicePropertyType)
	standardinstrumentdeparturetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentDepartureTimeSliceType)
	standardinstrumentdeparturetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentDepartureType)
	standardinstrumentdeparturetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surfacepropertytype *SurfacePropertyType) GongCopy() GongstructIF {
	newInstance := new(SurfacePropertyType)
	surfacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surfacetype *SurfaceType) GongCopy() GongstructIF {
	newInstance := new(SurfaceType)
	surfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongCopy() GongstructIF {
	newInstance := new(TerminalArrivalAreaSectorPropertyType)
	terminalarrivalareasectorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongCopy() GongstructIF {
	newInstance := new(TerminalArrivalAreaSectorType)
	terminalarrivalareasectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(TerminalArrivalAreaTimeSlicePropertyType)
	terminalarrivalareatimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(TerminalArrivalAreaTimeSliceType)
	terminalarrivalareatimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongCopy() GongstructIF {
	newInstance := new(TerminalArrivalAreaType)
	terminalarrivalareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(TerminalSegmentPointPropertyType)
	terminalsegmentpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongCopy() GongstructIF {
	newInstance := new(TerminalSegmentPointType)
	terminalsegmentpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timesheetpropertytype *TimesheetPropertyType) GongCopy() GongstructIF {
	newInstance := new(TimesheetPropertyType)
	timesheetpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timesheettype *TimesheetType) GongCopy() GongstructIF {
	newInstance := new(TimesheetType)
	timesheettype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongCopy() GongstructIF {
	newInstance := new(TouchDownLiftOffPropertyType)
	touchdownliftoffpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnplannedHoldingPropertyType)
	unplannedholdingpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(UnplannedHoldingTimeSlicePropertyType)
	unplannedholdingtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(UnplannedHoldingTimeSliceType)
	unplannedholdingtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unplannedholdingtype *UnplannedHoldingType) GongCopy() GongstructIF {
	newInstance := new(UnplannedHoldingType)
	unplannedholdingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructurePropertyType)
	verticalstructurepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (a_extension *A_extension) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension)
}

func (a_extension_1 *A_extension_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_1)
}

func (a_extension_10 *A_extension_10) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_10)
}

func (a_extension_11 *A_extension_11) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_11)
}

func (a_extension_12 *A_extension_12) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_12)
}

func (a_extension_13 *A_extension_13) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_13)
}

func (a_extension_14 *A_extension_14) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_14)
}

func (a_extension_15 *A_extension_15) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_15)
}

func (a_extension_16 *A_extension_16) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_16)
}

func (a_extension_17 *A_extension_17) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_17)
}

func (a_extension_18 *A_extension_18) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_18)
}

func (a_extension_19 *A_extension_19) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_19)
}

func (a_extension_2 *A_extension_2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_2)
}

func (a_extension_20 *A_extension_20) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_20)
}

func (a_extension_21 *A_extension_21) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_21)
}

func (a_extension_22 *A_extension_22) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_22)
}

func (a_extension_23 *A_extension_23) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_23)
}

func (a_extension_24 *A_extension_24) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_24)
}

func (a_extension_25 *A_extension_25) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_25)
}

func (a_extension_26 *A_extension_26) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_26)
}

func (a_extension_27 *A_extension_27) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_27)
}

func (a_extension_28 *A_extension_28) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_28)
}

func (a_extension_29 *A_extension_29) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_29)
}

func (a_extension_3 *A_extension_3) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_3)
}

func (a_extension_30 *A_extension_30) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_30)
}

func (a_extension_31 *A_extension_31) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_31)
}

func (a_extension_32 *A_extension_32) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_32)
}

func (a_extension_33 *A_extension_33) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_33)
}

func (a_extension_34 *A_extension_34) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_34)
}

func (a_extension_35 *A_extension_35) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_35)
}

func (a_extension_36 *A_extension_36) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_36)
}

func (a_extension_37 *A_extension_37) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_37)
}

func (a_extension_38 *A_extension_38) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_38)
}

func (a_extension_39 *A_extension_39) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_39)
}

func (a_extension_4 *A_extension_4) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_4)
}

func (a_extension_40 *A_extension_40) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_40)
}

func (a_extension_41 *A_extension_41) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_41)
}

func (a_extension_42 *A_extension_42) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_42)
}

func (a_extension_43 *A_extension_43) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_43)
}

func (a_extension_44 *A_extension_44) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_44)
}

func (a_extension_45 *A_extension_45) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_45)
}

func (a_extension_46 *A_extension_46) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_46)
}

func (a_extension_47 *A_extension_47) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_47)
}

func (a_extension_48 *A_extension_48) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_48)
}

func (a_extension_49 *A_extension_49) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_49)
}

func (a_extension_5 *A_extension_5) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_5)
}

func (a_extension_50 *A_extension_50) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_50)
}

func (a_extension_51 *A_extension_51) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_51)
}

func (a_extension_52 *A_extension_52) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_52)
}

func (a_extension_53 *A_extension_53) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_53)
}

func (a_extension_54 *A_extension_54) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_54)
}

func (a_extension_55 *A_extension_55) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_55)
}

func (a_extension_56 *A_extension_56) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_56)
}

func (a_extension_57 *A_extension_57) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_57)
}

func (a_extension_58 *A_extension_58) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_58)
}

func (a_extension_6 *A_extension_6) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_6)
}

func (a_extension_7 *A_extension_7) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_7)
}

func (a_extension_8 *A_extension_8) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_8)
}

func (a_extension_9 *A_extension_9) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_9)
}

func (abstractapproachlegtype *AbstractApproachLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractapproachlegtype)
}

func (abstractproceduretype *AbstractProcedureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractproceduretype)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractsegmentlegtype)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractsegmentpointtype)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingpointtype)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristicpropertytype)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristictype)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportheliportpropertytype)
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, altimetersourcepropertytype)
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, altitudeadjustmentpropertytype)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, altitudeadjustmenttype)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleindicationpropertytype)
}

func (angleusepropertytype *AngleUsePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleusepropertytype)
}

func (angleusetype *AngleUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleusetype)
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachaltitudetablepropertytype)
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachaltitudetabletype)
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachconditionpropertytype)
}

func (approachconditiontype *ApproachConditionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachconditiontype)
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachdistancetablepropertytype)
}

func (approachdistancetabletype *ApproachDistanceTableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachdistancetabletype)
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachtimingtablepropertytype)
}

func (approachtimingtabletype *ApproachTimingTableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, approachtimingtabletype)
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivalfeederlegtimeslicepropertytype)
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivalfeederlegtimeslicetype)
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivalfeederlegtype)
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivallegtimeslicepropertytype)
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivallegtimeslicetype)
}

func (arrivallegtype *ArrivalLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrivallegtype)
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlesectorpropertytype)
}

func (circlesectortype *CircleSectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlesectortype)
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlingareatimeslicepropertytype)
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlingareatimeslicetype)
}

func (circlingareatype *CirclingAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlingareatype)
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlingrestrictionpropertytype)
}

func (circlingrestrictiontype *CirclingRestrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlingrestrictiontype)
}

func (curvepropertytype *CurvePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvepropertytype)
}

func (curvetype *CurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvetype)
}

func (dmepropertytype *DMEPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dmepropertytype)
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, departurearrivalconditionpropertytype)
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, departurearrivalconditiontype)
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, departurelegtimeslicepropertytype)
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, departurelegtimeslicetype)
}

func (departurelegtype *DepartureLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, departurelegtype)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointpropertytype)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distanceindicationpropertytype)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedcurvetype)
}

func (elevatedpointtype *ElevatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointtype)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedsurfacetype)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, enroutesegmentpointtype)
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, equipmentunavailableadjustmentcolumnpropertytype)
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, equipmentunavailableadjustmentcolumntype)
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, equipmentunavailableadjustmentpropertytype)
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, equipmentunavailableadjustmenttype)
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fasdatablockpropertytype)
}

func (fasdatablocktype *FASDataBlockType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fasdatablocktype)
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, finallegtimeslicepropertytype)
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, finallegtimeslicetype)
}

func (finallegtype *FinalLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, finallegtype)
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, finalprofilepropertytype)
}

func (finalprofiletype *FinalProfileType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, finalprofiletype)
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingassessmenttimeslicepropertytype)
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingassessmenttimeslicetype)
}

func (holdingassessmenttype *HoldingAssessmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingassessmenttype)
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterndistancepropertytype)
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterndistancetype)
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterndurationpropertytype)
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterndurationtype)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatternpropertytype)
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterntimeslicepropertytype)
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterntimeslicetype)
}

func (holdingpatterntype *HoldingPatternType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatterntype)
}

func (holdingusepropertytype *HoldingUsePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingusepropertytype)
}

func (holdingusetype *HoldingUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingusetype)
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, initiallegtimeslicepropertytype)
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, initiallegtimeslicetype)
}

func (initiallegtype *InitialLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, initiallegtype)
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentapproachprocedurepropertytype)
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentapproachproceduretimeslicepropertytype)
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentapproachproceduretimeslicetype)
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentapproachproceduretype)
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, intermediatelegtimeslicepropertytype)
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, intermediatelegtimeslicetype)
}

func (intermediatelegtype *IntermediateLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, intermediatelegtype)
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, landingtakeoffareacollectionpropertytype)
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, landingtakeoffareacollectiontype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (minimapropertytype *MinimaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, minimapropertytype)
}

func (minimatype *MinimaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, minimatype)
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, missedapproachgrouppropertytype)
}

func (missedapproachgrouptype *MissedApproachGroupType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, missedapproachgrouptype)
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, missedapproachlegtimeslicepropertytype)
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, missedapproachlegtimeslicetype)
}

func (missedapproachlegtype *MissedApproachLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, missedapproachlegtype)
}

func (navaidpropertytype *NavaidPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidpropertytype)
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationarearestrictiontimeslicepropertytype)
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationarearestrictiontimeslicetype)
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationarearestrictiontype)
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationareasectorpropertytype)
}

func (navigationareasectortype *NavigationAreaSectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationareasectortype)
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationareatimeslicepropertytype)
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationareatimeslicetype)
}

func (navigationareatype *NavigationAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navigationareatype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleassessmentareapropertytype)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleassessmentareatype)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleplacementpropertytype)
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleplacementtype)
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstructionpropertytype)
}

func (obstructiontype *ObstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstructiontype)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritypropertytype)
}

func (pointpropertytype *PointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointpropertytype)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointreferencepropertytype)
}

func (pointreferencetype *PointReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointreferencetype)
}

func (pointtype *PointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointtype)
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, procedureavailabilitypropertytype)
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, procedureavailabilitytype)
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduredmetimeslicepropertytype)
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduredmetimeslicetype)
}

func (proceduredmetype *ProcedureDMEType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduredmetype)
}

func (procedurepropertytype *ProcedurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, procedurepropertytype)
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduretransitionlegpropertytype)
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduretransitionlegtype)
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduretransitionpropertytype)
}

func (proceduretransitiontype *ProcedureTransitionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, proceduretransitiontype)
}

func (radarsystempropertytype *RadarSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarsystempropertytype)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiofrequencyareapropertytype)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaycentrelinepointpropertytype)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaydirectionpropertytype)
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareapropertytype)
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareasectorpropertytype)
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareasectortype)
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareatimeslicepropertytype)
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareatimeslicetype)
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, safealtitudeareatype)
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sectordesignpropertytype)
}

func (sectordesigntype *SectorDesignType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sectordesigntype)
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, segmentlegpropertytype)
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, segmentpointpropertytype)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationsystempropertytype)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentarrivalpropertytype)
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentarrivaltimeslicepropertytype)
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentarrivaltimeslicetype)
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentarrivaltype)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentdeparturepropertytype)
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentdeparturetimeslicepropertytype)
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentdeparturetimeslicetype)
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentdeparturetype)
}

func (surfacepropertytype *SurfacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacepropertytype)
}

func (surfacetype *SurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacetype)
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalarrivalareasectorpropertytype)
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalarrivalareasectortype)
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalarrivalareatimeslicepropertytype)
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalarrivalareatimeslicetype)
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalarrivalareatype)
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalsegmentpointpropertytype)
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, terminalsegmentpointtype)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheetpropertytype)
}

func (timesheettype *TimesheetType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheettype)
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, touchdownliftoffpropertytype)
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unplannedholdingpropertytype)
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unplannedholdingtimeslicepropertytype)
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unplannedholdingtimeslicetype)
}

func (unplannedholdingtype *UnplannedHoldingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unplannedholdingtype)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurepropertytype)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.A_extensions,
		stage.A_extension_stagedOrder,
		stage.A_extensions_reference,
		&stage.A_extensions_referenceOrder,
		stage.A_extensions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_1s,
		stage.A_extension_1_stagedOrder,
		stage.A_extension_1s_reference,
		&stage.A_extension_1s_referenceOrder,
		stage.A_extension_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_10s,
		stage.A_extension_10_stagedOrder,
		stage.A_extension_10s_reference,
		&stage.A_extension_10s_referenceOrder,
		stage.A_extension_10s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_11s,
		stage.A_extension_11_stagedOrder,
		stage.A_extension_11s_reference,
		&stage.A_extension_11s_referenceOrder,
		stage.A_extension_11s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_12s,
		stage.A_extension_12_stagedOrder,
		stage.A_extension_12s_reference,
		&stage.A_extension_12s_referenceOrder,
		stage.A_extension_12s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_13s,
		stage.A_extension_13_stagedOrder,
		stage.A_extension_13s_reference,
		&stage.A_extension_13s_referenceOrder,
		stage.A_extension_13s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_14s,
		stage.A_extension_14_stagedOrder,
		stage.A_extension_14s_reference,
		&stage.A_extension_14s_referenceOrder,
		stage.A_extension_14s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_15s,
		stage.A_extension_15_stagedOrder,
		stage.A_extension_15s_reference,
		&stage.A_extension_15s_referenceOrder,
		stage.A_extension_15s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_16s,
		stage.A_extension_16_stagedOrder,
		stage.A_extension_16s_reference,
		&stage.A_extension_16s_referenceOrder,
		stage.A_extension_16s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_17s,
		stage.A_extension_17_stagedOrder,
		stage.A_extension_17s_reference,
		&stage.A_extension_17s_referenceOrder,
		stage.A_extension_17s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_18s,
		stage.A_extension_18_stagedOrder,
		stage.A_extension_18s_reference,
		&stage.A_extension_18s_referenceOrder,
		stage.A_extension_18s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_19s,
		stage.A_extension_19_stagedOrder,
		stage.A_extension_19s_reference,
		&stage.A_extension_19s_referenceOrder,
		stage.A_extension_19s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_2s,
		stage.A_extension_2_stagedOrder,
		stage.A_extension_2s_reference,
		&stage.A_extension_2s_referenceOrder,
		stage.A_extension_2s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_20s,
		stage.A_extension_20_stagedOrder,
		stage.A_extension_20s_reference,
		&stage.A_extension_20s_referenceOrder,
		stage.A_extension_20s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_21s,
		stage.A_extension_21_stagedOrder,
		stage.A_extension_21s_reference,
		&stage.A_extension_21s_referenceOrder,
		stage.A_extension_21s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_22s,
		stage.A_extension_22_stagedOrder,
		stage.A_extension_22s_reference,
		&stage.A_extension_22s_referenceOrder,
		stage.A_extension_22s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_23s,
		stage.A_extension_23_stagedOrder,
		stage.A_extension_23s_reference,
		&stage.A_extension_23s_referenceOrder,
		stage.A_extension_23s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_24s,
		stage.A_extension_24_stagedOrder,
		stage.A_extension_24s_reference,
		&stage.A_extension_24s_referenceOrder,
		stage.A_extension_24s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_25s,
		stage.A_extension_25_stagedOrder,
		stage.A_extension_25s_reference,
		&stage.A_extension_25s_referenceOrder,
		stage.A_extension_25s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_26s,
		stage.A_extension_26_stagedOrder,
		stage.A_extension_26s_reference,
		&stage.A_extension_26s_referenceOrder,
		stage.A_extension_26s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_27s,
		stage.A_extension_27_stagedOrder,
		stage.A_extension_27s_reference,
		&stage.A_extension_27s_referenceOrder,
		stage.A_extension_27s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_28s,
		stage.A_extension_28_stagedOrder,
		stage.A_extension_28s_reference,
		&stage.A_extension_28s_referenceOrder,
		stage.A_extension_28s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_29s,
		stage.A_extension_29_stagedOrder,
		stage.A_extension_29s_reference,
		&stage.A_extension_29s_referenceOrder,
		stage.A_extension_29s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_3s,
		stage.A_extension_3_stagedOrder,
		stage.A_extension_3s_reference,
		&stage.A_extension_3s_referenceOrder,
		stage.A_extension_3s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_30s,
		stage.A_extension_30_stagedOrder,
		stage.A_extension_30s_reference,
		&stage.A_extension_30s_referenceOrder,
		stage.A_extension_30s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_31s,
		stage.A_extension_31_stagedOrder,
		stage.A_extension_31s_reference,
		&stage.A_extension_31s_referenceOrder,
		stage.A_extension_31s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_32s,
		stage.A_extension_32_stagedOrder,
		stage.A_extension_32s_reference,
		&stage.A_extension_32s_referenceOrder,
		stage.A_extension_32s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_33s,
		stage.A_extension_33_stagedOrder,
		stage.A_extension_33s_reference,
		&stage.A_extension_33s_referenceOrder,
		stage.A_extension_33s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_34s,
		stage.A_extension_34_stagedOrder,
		stage.A_extension_34s_reference,
		&stage.A_extension_34s_referenceOrder,
		stage.A_extension_34s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_35s,
		stage.A_extension_35_stagedOrder,
		stage.A_extension_35s_reference,
		&stage.A_extension_35s_referenceOrder,
		stage.A_extension_35s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_36s,
		stage.A_extension_36_stagedOrder,
		stage.A_extension_36s_reference,
		&stage.A_extension_36s_referenceOrder,
		stage.A_extension_36s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_37s,
		stage.A_extension_37_stagedOrder,
		stage.A_extension_37s_reference,
		&stage.A_extension_37s_referenceOrder,
		stage.A_extension_37s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_38s,
		stage.A_extension_38_stagedOrder,
		stage.A_extension_38s_reference,
		&stage.A_extension_38s_referenceOrder,
		stage.A_extension_38s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_39s,
		stage.A_extension_39_stagedOrder,
		stage.A_extension_39s_reference,
		&stage.A_extension_39s_referenceOrder,
		stage.A_extension_39s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_4s,
		stage.A_extension_4_stagedOrder,
		stage.A_extension_4s_reference,
		&stage.A_extension_4s_referenceOrder,
		stage.A_extension_4s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_40s,
		stage.A_extension_40_stagedOrder,
		stage.A_extension_40s_reference,
		&stage.A_extension_40s_referenceOrder,
		stage.A_extension_40s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_41s,
		stage.A_extension_41_stagedOrder,
		stage.A_extension_41s_reference,
		&stage.A_extension_41s_referenceOrder,
		stage.A_extension_41s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_42s,
		stage.A_extension_42_stagedOrder,
		stage.A_extension_42s_reference,
		&stage.A_extension_42s_referenceOrder,
		stage.A_extension_42s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_43s,
		stage.A_extension_43_stagedOrder,
		stage.A_extension_43s_reference,
		&stage.A_extension_43s_referenceOrder,
		stage.A_extension_43s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_44s,
		stage.A_extension_44_stagedOrder,
		stage.A_extension_44s_reference,
		&stage.A_extension_44s_referenceOrder,
		stage.A_extension_44s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_45s,
		stage.A_extension_45_stagedOrder,
		stage.A_extension_45s_reference,
		&stage.A_extension_45s_referenceOrder,
		stage.A_extension_45s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_46s,
		stage.A_extension_46_stagedOrder,
		stage.A_extension_46s_reference,
		&stage.A_extension_46s_referenceOrder,
		stage.A_extension_46s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_47s,
		stage.A_extension_47_stagedOrder,
		stage.A_extension_47s_reference,
		&stage.A_extension_47s_referenceOrder,
		stage.A_extension_47s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_48s,
		stage.A_extension_48_stagedOrder,
		stage.A_extension_48s_reference,
		&stage.A_extension_48s_referenceOrder,
		stage.A_extension_48s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_49s,
		stage.A_extension_49_stagedOrder,
		stage.A_extension_49s_reference,
		&stage.A_extension_49s_referenceOrder,
		stage.A_extension_49s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_5s,
		stage.A_extension_5_stagedOrder,
		stage.A_extension_5s_reference,
		&stage.A_extension_5s_referenceOrder,
		stage.A_extension_5s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_50s,
		stage.A_extension_50_stagedOrder,
		stage.A_extension_50s_reference,
		&stage.A_extension_50s_referenceOrder,
		stage.A_extension_50s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_51s,
		stage.A_extension_51_stagedOrder,
		stage.A_extension_51s_reference,
		&stage.A_extension_51s_referenceOrder,
		stage.A_extension_51s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_52s,
		stage.A_extension_52_stagedOrder,
		stage.A_extension_52s_reference,
		&stage.A_extension_52s_referenceOrder,
		stage.A_extension_52s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_53s,
		stage.A_extension_53_stagedOrder,
		stage.A_extension_53s_reference,
		&stage.A_extension_53s_referenceOrder,
		stage.A_extension_53s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_54s,
		stage.A_extension_54_stagedOrder,
		stage.A_extension_54s_reference,
		&stage.A_extension_54s_referenceOrder,
		stage.A_extension_54s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_55s,
		stage.A_extension_55_stagedOrder,
		stage.A_extension_55s_reference,
		&stage.A_extension_55s_referenceOrder,
		stage.A_extension_55s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_56s,
		stage.A_extension_56_stagedOrder,
		stage.A_extension_56s_reference,
		&stage.A_extension_56s_referenceOrder,
		stage.A_extension_56s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_57s,
		stage.A_extension_57_stagedOrder,
		stage.A_extension_57s_reference,
		&stage.A_extension_57s_referenceOrder,
		stage.A_extension_57s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_58s,
		stage.A_extension_58_stagedOrder,
		stage.A_extension_58s_reference,
		&stage.A_extension_58s_referenceOrder,
		stage.A_extension_58s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_6s,
		stage.A_extension_6_stagedOrder,
		stage.A_extension_6s_reference,
		&stage.A_extension_6s_referenceOrder,
		stage.A_extension_6s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_7s,
		stage.A_extension_7_stagedOrder,
		stage.A_extension_7s_reference,
		&stage.A_extension_7s_referenceOrder,
		stage.A_extension_7s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_8s,
		stage.A_extension_8_stagedOrder,
		stage.A_extension_8s_reference,
		&stage.A_extension_8s_referenceOrder,
		stage.A_extension_8s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_9s,
		stage.A_extension_9_stagedOrder,
		stage.A_extension_9s_reference,
		&stage.A_extension_9s_referenceOrder,
		stage.A_extension_9s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractApproachLegTypes,
		stage.AbstractApproachLegType_stagedOrder,
		stage.AbstractApproachLegTypes_reference,
		&stage.AbstractApproachLegTypes_referenceOrder,
		stage.AbstractApproachLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractProcedureTypes,
		stage.AbstractProcedureType_stagedOrder,
		stage.AbstractProcedureTypes_reference,
		&stage.AbstractProcedureTypes_referenceOrder,
		stage.AbstractProcedureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractPropertiesWithScheduleTypes,
		stage.AbstractPropertiesWithScheduleType_stagedOrder,
		stage.AbstractPropertiesWithScheduleTypes_reference,
		&stage.AbstractPropertiesWithScheduleTypes_referenceOrder,
		stage.AbstractPropertiesWithScheduleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractSegmentLegTypes,
		stage.AbstractSegmentLegType_stagedOrder,
		stage.AbstractSegmentLegTypes_reference,
		&stage.AbstractSegmentLegTypes_referenceOrder,
		stage.AbstractSegmentLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractSegmentPointTypes,
		stage.AbstractSegmentPointType_stagedOrder,
		stage.AbstractSegmentPointTypes_reference,
		&stage.AbstractSegmentPointTypes_referenceOrder,
		stage.AbstractSegmentPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AerialRefuellingPointTypes,
		stage.AerialRefuellingPointType_stagedOrder,
		stage.AerialRefuellingPointTypes_reference,
		&stage.AerialRefuellingPointTypes_referenceOrder,
		stage.AerialRefuellingPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftCharacteristicPropertyTypes,
		stage.AircraftCharacteristicPropertyType_stagedOrder,
		stage.AircraftCharacteristicPropertyTypes_reference,
		&stage.AircraftCharacteristicPropertyTypes_referenceOrder,
		stage.AircraftCharacteristicPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftCharacteristicTypes,
		stage.AircraftCharacteristicType_stagedOrder,
		stage.AircraftCharacteristicTypes_reference,
		&stage.AircraftCharacteristicTypes_referenceOrder,
		stage.AircraftCharacteristicTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportHeliportPropertyTypes,
		stage.AirportHeliportPropertyType_stagedOrder,
		stage.AirportHeliportPropertyTypes_reference,
		&stage.AirportHeliportPropertyTypes_referenceOrder,
		stage.AirportHeliportPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AltimeterSourcePropertyTypes,
		stage.AltimeterSourcePropertyType_stagedOrder,
		stage.AltimeterSourcePropertyTypes_reference,
		&stage.AltimeterSourcePropertyTypes_referenceOrder,
		stage.AltimeterSourcePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AltitudeAdjustmentPropertyTypes,
		stage.AltitudeAdjustmentPropertyType_stagedOrder,
		stage.AltitudeAdjustmentPropertyTypes_reference,
		&stage.AltitudeAdjustmentPropertyTypes_referenceOrder,
		stage.AltitudeAdjustmentPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AltitudeAdjustmentTypes,
		stage.AltitudeAdjustmentType_stagedOrder,
		stage.AltitudeAdjustmentTypes_reference,
		&stage.AltitudeAdjustmentTypes_referenceOrder,
		stage.AltitudeAdjustmentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AngleIndicationPropertyTypes,
		stage.AngleIndicationPropertyType_stagedOrder,
		stage.AngleIndicationPropertyTypes_reference,
		&stage.AngleIndicationPropertyTypes_referenceOrder,
		stage.AngleIndicationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AngleUsePropertyTypes,
		stage.AngleUsePropertyType_stagedOrder,
		stage.AngleUsePropertyTypes_reference,
		&stage.AngleUsePropertyTypes_referenceOrder,
		stage.AngleUsePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AngleUseTypes,
		stage.AngleUseType_stagedOrder,
		stage.AngleUseTypes_reference,
		&stage.AngleUseTypes_referenceOrder,
		stage.AngleUseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachAltitudeTablePropertyTypes,
		stage.ApproachAltitudeTablePropertyType_stagedOrder,
		stage.ApproachAltitudeTablePropertyTypes_reference,
		&stage.ApproachAltitudeTablePropertyTypes_referenceOrder,
		stage.ApproachAltitudeTablePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachAltitudeTableTypes,
		stage.ApproachAltitudeTableType_stagedOrder,
		stage.ApproachAltitudeTableTypes_reference,
		&stage.ApproachAltitudeTableTypes_referenceOrder,
		stage.ApproachAltitudeTableTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachConditionPropertyTypes,
		stage.ApproachConditionPropertyType_stagedOrder,
		stage.ApproachConditionPropertyTypes_reference,
		&stage.ApproachConditionPropertyTypes_referenceOrder,
		stage.ApproachConditionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachConditionTypes,
		stage.ApproachConditionType_stagedOrder,
		stage.ApproachConditionTypes_reference,
		&stage.ApproachConditionTypes_referenceOrder,
		stage.ApproachConditionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachDistanceTablePropertyTypes,
		stage.ApproachDistanceTablePropertyType_stagedOrder,
		stage.ApproachDistanceTablePropertyTypes_reference,
		&stage.ApproachDistanceTablePropertyTypes_referenceOrder,
		stage.ApproachDistanceTablePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachDistanceTableTypes,
		stage.ApproachDistanceTableType_stagedOrder,
		stage.ApproachDistanceTableTypes_reference,
		&stage.ApproachDistanceTableTypes_referenceOrder,
		stage.ApproachDistanceTableTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachTimingTablePropertyTypes,
		stage.ApproachTimingTablePropertyType_stagedOrder,
		stage.ApproachTimingTablePropertyTypes_reference,
		&stage.ApproachTimingTablePropertyTypes_referenceOrder,
		stage.ApproachTimingTablePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ApproachTimingTableTypes,
		stage.ApproachTimingTableType_stagedOrder,
		stage.ApproachTimingTableTypes_reference,
		&stage.ApproachTimingTableTypes_referenceOrder,
		stage.ApproachTimingTableTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalFeederLegTimeSlicePropertyTypes,
		stage.ArrivalFeederLegTimeSlicePropertyType_stagedOrder,
		stage.ArrivalFeederLegTimeSlicePropertyTypes_reference,
		&stage.ArrivalFeederLegTimeSlicePropertyTypes_referenceOrder,
		stage.ArrivalFeederLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalFeederLegTimeSliceTypes,
		stage.ArrivalFeederLegTimeSliceType_stagedOrder,
		stage.ArrivalFeederLegTimeSliceTypes_reference,
		&stage.ArrivalFeederLegTimeSliceTypes_referenceOrder,
		stage.ArrivalFeederLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalFeederLegTypes,
		stage.ArrivalFeederLegType_stagedOrder,
		stage.ArrivalFeederLegTypes_reference,
		&stage.ArrivalFeederLegTypes_referenceOrder,
		stage.ArrivalFeederLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalLegTimeSlicePropertyTypes,
		stage.ArrivalLegTimeSlicePropertyType_stagedOrder,
		stage.ArrivalLegTimeSlicePropertyTypes_reference,
		&stage.ArrivalLegTimeSlicePropertyTypes_referenceOrder,
		stage.ArrivalLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalLegTimeSliceTypes,
		stage.ArrivalLegTimeSliceType_stagedOrder,
		stage.ArrivalLegTimeSliceTypes_reference,
		&stage.ArrivalLegTimeSliceTypes_referenceOrder,
		stage.ArrivalLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ArrivalLegTypes,
		stage.ArrivalLegType_stagedOrder,
		stage.ArrivalLegTypes_reference,
		&stage.ArrivalLegTypes_referenceOrder,
		stage.ArrivalLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CircleSectorPropertyTypes,
		stage.CircleSectorPropertyType_stagedOrder,
		stage.CircleSectorPropertyTypes_reference,
		&stage.CircleSectorPropertyTypes_referenceOrder,
		stage.CircleSectorPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CircleSectorTypes,
		stage.CircleSectorType_stagedOrder,
		stage.CircleSectorTypes_reference,
		&stage.CircleSectorTypes_referenceOrder,
		stage.CircleSectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CirclingAreaTimeSlicePropertyTypes,
		stage.CirclingAreaTimeSlicePropertyType_stagedOrder,
		stage.CirclingAreaTimeSlicePropertyTypes_reference,
		&stage.CirclingAreaTimeSlicePropertyTypes_referenceOrder,
		stage.CirclingAreaTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CirclingAreaTimeSliceTypes,
		stage.CirclingAreaTimeSliceType_stagedOrder,
		stage.CirclingAreaTimeSliceTypes_reference,
		&stage.CirclingAreaTimeSliceTypes_referenceOrder,
		stage.CirclingAreaTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CirclingAreaTypes,
		stage.CirclingAreaType_stagedOrder,
		stage.CirclingAreaTypes_reference,
		&stage.CirclingAreaTypes_referenceOrder,
		stage.CirclingAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CirclingRestrictionPropertyTypes,
		stage.CirclingRestrictionPropertyType_stagedOrder,
		stage.CirclingRestrictionPropertyTypes_reference,
		&stage.CirclingRestrictionPropertyTypes_referenceOrder,
		stage.CirclingRestrictionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CirclingRestrictionTypes,
		stage.CirclingRestrictionType_stagedOrder,
		stage.CirclingRestrictionTypes_reference,
		&stage.CirclingRestrictionTypes_referenceOrder,
		stage.CirclingRestrictionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CurvePropertyTypes,
		stage.CurvePropertyType_stagedOrder,
		stage.CurvePropertyTypes_reference,
		&stage.CurvePropertyTypes_referenceOrder,
		stage.CurvePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CurveTypes,
		stage.CurveType_stagedOrder,
		stage.CurveTypes_reference,
		&stage.CurveTypes_referenceOrder,
		stage.CurveTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DMEPropertyTypes,
		stage.DMEPropertyType_stagedOrder,
		stage.DMEPropertyTypes_reference,
		&stage.DMEPropertyTypes_referenceOrder,
		stage.DMEPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DepartureArrivalConditionPropertyTypes,
		stage.DepartureArrivalConditionPropertyType_stagedOrder,
		stage.DepartureArrivalConditionPropertyTypes_reference,
		&stage.DepartureArrivalConditionPropertyTypes_referenceOrder,
		stage.DepartureArrivalConditionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DepartureArrivalConditionTypes,
		stage.DepartureArrivalConditionType_stagedOrder,
		stage.DepartureArrivalConditionTypes_reference,
		&stage.DepartureArrivalConditionTypes_referenceOrder,
		stage.DepartureArrivalConditionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DepartureLegTimeSlicePropertyTypes,
		stage.DepartureLegTimeSlicePropertyType_stagedOrder,
		stage.DepartureLegTimeSlicePropertyTypes_reference,
		&stage.DepartureLegTimeSlicePropertyTypes_referenceOrder,
		stage.DepartureLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DepartureLegTimeSliceTypes,
		stage.DepartureLegTimeSliceType_stagedOrder,
		stage.DepartureLegTimeSliceTypes_reference,
		&stage.DepartureLegTimeSliceTypes_referenceOrder,
		stage.DepartureLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DepartureLegTypes,
		stage.DepartureLegType_stagedOrder,
		stage.DepartureLegTypes_reference,
		&stage.DepartureLegTypes_referenceOrder,
		stage.DepartureLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DesignatedPointPropertyTypes,
		stage.DesignatedPointPropertyType_stagedOrder,
		stage.DesignatedPointPropertyTypes_reference,
		&stage.DesignatedPointPropertyTypes_referenceOrder,
		stage.DesignatedPointPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DistanceIndicationPropertyTypes,
		stage.DistanceIndicationPropertyType_stagedOrder,
		stage.DistanceIndicationPropertyTypes_reference,
		&stage.DistanceIndicationPropertyTypes_referenceOrder,
		stage.DistanceIndicationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ElevatedCurveTypes,
		stage.ElevatedCurveType_stagedOrder,
		stage.ElevatedCurveTypes_reference,
		&stage.ElevatedCurveTypes_referenceOrder,
		stage.ElevatedCurveTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ElevatedPointTypes,
		stage.ElevatedPointType_stagedOrder,
		stage.ElevatedPointTypes_reference,
		&stage.ElevatedPointTypes_referenceOrder,
		stage.ElevatedPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ElevatedSurfaceTypes,
		stage.ElevatedSurfaceType_stagedOrder,
		stage.ElevatedSurfaceTypes_reference,
		&stage.ElevatedSurfaceTypes_referenceOrder,
		stage.ElevatedSurfaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EnRouteSegmentPointTypes,
		stage.EnRouteSegmentPointType_stagedOrder,
		stage.EnRouteSegmentPointTypes_reference,
		&stage.EnRouteSegmentPointTypes_referenceOrder,
		stage.EnRouteSegmentPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EquipmentUnavailableAdjustmentColumnPropertyTypes,
		stage.EquipmentUnavailableAdjustmentColumnPropertyType_stagedOrder,
		stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_reference,
		&stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_referenceOrder,
		stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EquipmentUnavailableAdjustmentColumnTypes,
		stage.EquipmentUnavailableAdjustmentColumnType_stagedOrder,
		stage.EquipmentUnavailableAdjustmentColumnTypes_reference,
		&stage.EquipmentUnavailableAdjustmentColumnTypes_referenceOrder,
		stage.EquipmentUnavailableAdjustmentColumnTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EquipmentUnavailableAdjustmentPropertyTypes,
		stage.EquipmentUnavailableAdjustmentPropertyType_stagedOrder,
		stage.EquipmentUnavailableAdjustmentPropertyTypes_reference,
		&stage.EquipmentUnavailableAdjustmentPropertyTypes_referenceOrder,
		stage.EquipmentUnavailableAdjustmentPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EquipmentUnavailableAdjustmentTypes,
		stage.EquipmentUnavailableAdjustmentType_stagedOrder,
		stage.EquipmentUnavailableAdjustmentTypes_reference,
		&stage.EquipmentUnavailableAdjustmentTypes_referenceOrder,
		stage.EquipmentUnavailableAdjustmentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FASDataBlockPropertyTypes,
		stage.FASDataBlockPropertyType_stagedOrder,
		stage.FASDataBlockPropertyTypes_reference,
		&stage.FASDataBlockPropertyTypes_referenceOrder,
		stage.FASDataBlockPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FASDataBlockTypes,
		stage.FASDataBlockType_stagedOrder,
		stage.FASDataBlockTypes_reference,
		&stage.FASDataBlockTypes_referenceOrder,
		stage.FASDataBlockTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FinalLegTimeSlicePropertyTypes,
		stage.FinalLegTimeSlicePropertyType_stagedOrder,
		stage.FinalLegTimeSlicePropertyTypes_reference,
		&stage.FinalLegTimeSlicePropertyTypes_referenceOrder,
		stage.FinalLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FinalLegTimeSliceTypes,
		stage.FinalLegTimeSliceType_stagedOrder,
		stage.FinalLegTimeSliceTypes_reference,
		&stage.FinalLegTimeSliceTypes_referenceOrder,
		stage.FinalLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FinalLegTypes,
		stage.FinalLegType_stagedOrder,
		stage.FinalLegTypes_reference,
		&stage.FinalLegTypes_referenceOrder,
		stage.FinalLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FinalProfilePropertyTypes,
		stage.FinalProfilePropertyType_stagedOrder,
		stage.FinalProfilePropertyTypes_reference,
		&stage.FinalProfilePropertyTypes_referenceOrder,
		stage.FinalProfilePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FinalProfileTypes,
		stage.FinalProfileType_stagedOrder,
		stage.FinalProfileTypes_reference,
		&stage.FinalProfileTypes_referenceOrder,
		stage.FinalProfileTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingAssessmentTimeSlicePropertyTypes,
		stage.HoldingAssessmentTimeSlicePropertyType_stagedOrder,
		stage.HoldingAssessmentTimeSlicePropertyTypes_reference,
		&stage.HoldingAssessmentTimeSlicePropertyTypes_referenceOrder,
		stage.HoldingAssessmentTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingAssessmentTimeSliceTypes,
		stage.HoldingAssessmentTimeSliceType_stagedOrder,
		stage.HoldingAssessmentTimeSliceTypes_reference,
		&stage.HoldingAssessmentTimeSliceTypes_referenceOrder,
		stage.HoldingAssessmentTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingAssessmentTypes,
		stage.HoldingAssessmentType_stagedOrder,
		stage.HoldingAssessmentTypes_reference,
		&stage.HoldingAssessmentTypes_referenceOrder,
		stage.HoldingAssessmentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternDistancePropertyTypes,
		stage.HoldingPatternDistancePropertyType_stagedOrder,
		stage.HoldingPatternDistancePropertyTypes_reference,
		&stage.HoldingPatternDistancePropertyTypes_referenceOrder,
		stage.HoldingPatternDistancePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternDistanceTypes,
		stage.HoldingPatternDistanceType_stagedOrder,
		stage.HoldingPatternDistanceTypes_reference,
		&stage.HoldingPatternDistanceTypes_referenceOrder,
		stage.HoldingPatternDistanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternDurationPropertyTypes,
		stage.HoldingPatternDurationPropertyType_stagedOrder,
		stage.HoldingPatternDurationPropertyTypes_reference,
		&stage.HoldingPatternDurationPropertyTypes_referenceOrder,
		stage.HoldingPatternDurationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternDurationTypes,
		stage.HoldingPatternDurationType_stagedOrder,
		stage.HoldingPatternDurationTypes_reference,
		&stage.HoldingPatternDurationTypes_referenceOrder,
		stage.HoldingPatternDurationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternPropertyTypes,
		stage.HoldingPatternPropertyType_stagedOrder,
		stage.HoldingPatternPropertyTypes_reference,
		&stage.HoldingPatternPropertyTypes_referenceOrder,
		stage.HoldingPatternPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternTimeSlicePropertyTypes,
		stage.HoldingPatternTimeSlicePropertyType_stagedOrder,
		stage.HoldingPatternTimeSlicePropertyTypes_reference,
		&stage.HoldingPatternTimeSlicePropertyTypes_referenceOrder,
		stage.HoldingPatternTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternTimeSliceTypes,
		stage.HoldingPatternTimeSliceType_stagedOrder,
		stage.HoldingPatternTimeSliceTypes_reference,
		&stage.HoldingPatternTimeSliceTypes_referenceOrder,
		stage.HoldingPatternTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternTypes,
		stage.HoldingPatternType_stagedOrder,
		stage.HoldingPatternTypes_reference,
		&stage.HoldingPatternTypes_referenceOrder,
		stage.HoldingPatternTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingUsePropertyTypes,
		stage.HoldingUsePropertyType_stagedOrder,
		stage.HoldingUsePropertyTypes_reference,
		&stage.HoldingUsePropertyTypes_referenceOrder,
		stage.HoldingUsePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingUseTypes,
		stage.HoldingUseType_stagedOrder,
		stage.HoldingUseTypes_reference,
		&stage.HoldingUseTypes_referenceOrder,
		stage.HoldingUseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InitialLegTimeSlicePropertyTypes,
		stage.InitialLegTimeSlicePropertyType_stagedOrder,
		stage.InitialLegTimeSlicePropertyTypes_reference,
		&stage.InitialLegTimeSlicePropertyTypes_referenceOrder,
		stage.InitialLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InitialLegTimeSliceTypes,
		stage.InitialLegTimeSliceType_stagedOrder,
		stage.InitialLegTimeSliceTypes_reference,
		&stage.InitialLegTimeSliceTypes_referenceOrder,
		stage.InitialLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InitialLegTypes,
		stage.InitialLegType_stagedOrder,
		stage.InitialLegTypes_reference,
		&stage.InitialLegTypes_referenceOrder,
		stage.InitialLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentApproachProcedurePropertyTypes,
		stage.InstrumentApproachProcedurePropertyType_stagedOrder,
		stage.InstrumentApproachProcedurePropertyTypes_reference,
		&stage.InstrumentApproachProcedurePropertyTypes_referenceOrder,
		stage.InstrumentApproachProcedurePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentApproachProcedureTimeSlicePropertyTypes,
		stage.InstrumentApproachProcedureTimeSlicePropertyType_stagedOrder,
		stage.InstrumentApproachProcedureTimeSlicePropertyTypes_reference,
		&stage.InstrumentApproachProcedureTimeSlicePropertyTypes_referenceOrder,
		stage.InstrumentApproachProcedureTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentApproachProcedureTimeSliceTypes,
		stage.InstrumentApproachProcedureTimeSliceType_stagedOrder,
		stage.InstrumentApproachProcedureTimeSliceTypes_reference,
		&stage.InstrumentApproachProcedureTimeSliceTypes_referenceOrder,
		stage.InstrumentApproachProcedureTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentApproachProcedureTypes,
		stage.InstrumentApproachProcedureType_stagedOrder,
		stage.InstrumentApproachProcedureTypes_reference,
		&stage.InstrumentApproachProcedureTypes_referenceOrder,
		stage.InstrumentApproachProcedureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.IntermediateLegTimeSlicePropertyTypes,
		stage.IntermediateLegTimeSlicePropertyType_stagedOrder,
		stage.IntermediateLegTimeSlicePropertyTypes_reference,
		&stage.IntermediateLegTimeSlicePropertyTypes_referenceOrder,
		stage.IntermediateLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.IntermediateLegTimeSliceTypes,
		stage.IntermediateLegTimeSliceType_stagedOrder,
		stage.IntermediateLegTimeSliceTypes_reference,
		&stage.IntermediateLegTimeSliceTypes_referenceOrder,
		stage.IntermediateLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.IntermediateLegTypes,
		stage.IntermediateLegType_stagedOrder,
		stage.IntermediateLegTypes_reference,
		&stage.IntermediateLegTypes_referenceOrder,
		stage.IntermediateLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LandingTakeoffAreaCollectionPropertyTypes,
		stage.LandingTakeoffAreaCollectionPropertyType_stagedOrder,
		stage.LandingTakeoffAreaCollectionPropertyTypes_reference,
		&stage.LandingTakeoffAreaCollectionPropertyTypes_referenceOrder,
		stage.LandingTakeoffAreaCollectionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LandingTakeoffAreaCollectionTypes,
		stage.LandingTakeoffAreaCollectionType_stagedOrder,
		stage.LandingTakeoffAreaCollectionTypes_reference,
		&stage.LandingTakeoffAreaCollectionTypes_referenceOrder,
		stage.LandingTakeoffAreaCollectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinguisticNotePropertyTypes,
		stage.LinguisticNotePropertyType_stagedOrder,
		stage.LinguisticNotePropertyTypes_reference,
		&stage.LinguisticNotePropertyTypes_referenceOrder,
		stage.LinguisticNotePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinguisticNoteTypes,
		stage.LinguisticNoteType_stagedOrder,
		stage.LinguisticNoteTypes_reference,
		&stage.LinguisticNoteTypes_referenceOrder,
		stage.LinguisticNoteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MinimaPropertyTypes,
		stage.MinimaPropertyType_stagedOrder,
		stage.MinimaPropertyTypes_reference,
		&stage.MinimaPropertyTypes_referenceOrder,
		stage.MinimaPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MinimaTypes,
		stage.MinimaType_stagedOrder,
		stage.MinimaTypes_reference,
		&stage.MinimaTypes_referenceOrder,
		stage.MinimaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MissedApproachGroupPropertyTypes,
		stage.MissedApproachGroupPropertyType_stagedOrder,
		stage.MissedApproachGroupPropertyTypes_reference,
		&stage.MissedApproachGroupPropertyTypes_referenceOrder,
		stage.MissedApproachGroupPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MissedApproachGroupTypes,
		stage.MissedApproachGroupType_stagedOrder,
		stage.MissedApproachGroupTypes_reference,
		&stage.MissedApproachGroupTypes_referenceOrder,
		stage.MissedApproachGroupTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MissedApproachLegTimeSlicePropertyTypes,
		stage.MissedApproachLegTimeSlicePropertyType_stagedOrder,
		stage.MissedApproachLegTimeSlicePropertyTypes_reference,
		&stage.MissedApproachLegTimeSlicePropertyTypes_referenceOrder,
		stage.MissedApproachLegTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MissedApproachLegTimeSliceTypes,
		stage.MissedApproachLegTimeSliceType_stagedOrder,
		stage.MissedApproachLegTimeSliceTypes_reference,
		&stage.MissedApproachLegTimeSliceTypes_referenceOrder,
		stage.MissedApproachLegTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MissedApproachLegTypes,
		stage.MissedApproachLegType_stagedOrder,
		stage.MissedApproachLegTypes_reference,
		&stage.MissedApproachLegTypes_referenceOrder,
		stage.MissedApproachLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavaidPropertyTypes,
		stage.NavaidPropertyType_stagedOrder,
		stage.NavaidPropertyTypes_reference,
		&stage.NavaidPropertyTypes_referenceOrder,
		stage.NavaidPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaRestrictionTimeSlicePropertyTypes,
		stage.NavigationAreaRestrictionTimeSlicePropertyType_stagedOrder,
		stage.NavigationAreaRestrictionTimeSlicePropertyTypes_reference,
		&stage.NavigationAreaRestrictionTimeSlicePropertyTypes_referenceOrder,
		stage.NavigationAreaRestrictionTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaRestrictionTimeSliceTypes,
		stage.NavigationAreaRestrictionTimeSliceType_stagedOrder,
		stage.NavigationAreaRestrictionTimeSliceTypes_reference,
		&stage.NavigationAreaRestrictionTimeSliceTypes_referenceOrder,
		stage.NavigationAreaRestrictionTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaRestrictionTypes,
		stage.NavigationAreaRestrictionType_stagedOrder,
		stage.NavigationAreaRestrictionTypes_reference,
		&stage.NavigationAreaRestrictionTypes_referenceOrder,
		stage.NavigationAreaRestrictionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaSectorPropertyTypes,
		stage.NavigationAreaSectorPropertyType_stagedOrder,
		stage.NavigationAreaSectorPropertyTypes_reference,
		&stage.NavigationAreaSectorPropertyTypes_referenceOrder,
		stage.NavigationAreaSectorPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaSectorTypes,
		stage.NavigationAreaSectorType_stagedOrder,
		stage.NavigationAreaSectorTypes_reference,
		&stage.NavigationAreaSectorTypes_referenceOrder,
		stage.NavigationAreaSectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaTimeSlicePropertyTypes,
		stage.NavigationAreaTimeSlicePropertyType_stagedOrder,
		stage.NavigationAreaTimeSlicePropertyTypes_reference,
		&stage.NavigationAreaTimeSlicePropertyTypes_referenceOrder,
		stage.NavigationAreaTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaTimeSliceTypes,
		stage.NavigationAreaTimeSliceType_stagedOrder,
		stage.NavigationAreaTimeSliceTypes_reference,
		&stage.NavigationAreaTimeSliceTypes_referenceOrder,
		stage.NavigationAreaTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NavigationAreaTypes,
		stage.NavigationAreaType_stagedOrder,
		stage.NavigationAreaTypes_reference,
		&stage.NavigationAreaTypes_referenceOrder,
		stage.NavigationAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NotePropertyTypes,
		stage.NotePropertyType_stagedOrder,
		stage.NotePropertyTypes_reference,
		&stage.NotePropertyTypes_referenceOrder,
		stage.NotePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteTypes,
		stage.NoteType_stagedOrder,
		stage.NoteTypes_reference,
		&stage.NoteTypes_referenceOrder,
		stage.NoteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstacleAssessmentAreaPropertyTypes,
		stage.ObstacleAssessmentAreaPropertyType_stagedOrder,
		stage.ObstacleAssessmentAreaPropertyTypes_reference,
		&stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder,
		stage.ObstacleAssessmentAreaPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstacleAssessmentAreaTypes,
		stage.ObstacleAssessmentAreaType_stagedOrder,
		stage.ObstacleAssessmentAreaTypes_reference,
		&stage.ObstacleAssessmentAreaTypes_referenceOrder,
		stage.ObstacleAssessmentAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstaclePlacementPropertyTypes,
		stage.ObstaclePlacementPropertyType_stagedOrder,
		stage.ObstaclePlacementPropertyTypes_reference,
		&stage.ObstaclePlacementPropertyTypes_referenceOrder,
		stage.ObstaclePlacementPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstaclePlacementTypes,
		stage.ObstaclePlacementType_stagedOrder,
		stage.ObstaclePlacementTypes_reference,
		&stage.ObstaclePlacementTypes_referenceOrder,
		stage.ObstaclePlacementTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstructionPropertyTypes,
		stage.ObstructionPropertyType_stagedOrder,
		stage.ObstructionPropertyTypes_reference,
		&stage.ObstructionPropertyTypes_referenceOrder,
		stage.ObstructionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ObstructionTypes,
		stage.ObstructionType_stagedOrder,
		stage.ObstructionTypes_reference,
		&stage.ObstructionTypes_referenceOrder,
		stage.ObstructionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OrganisationAuthorityPropertyTypes,
		stage.OrganisationAuthorityPropertyType_stagedOrder,
		stage.OrganisationAuthorityPropertyTypes_reference,
		&stage.OrganisationAuthorityPropertyTypes_referenceOrder,
		stage.OrganisationAuthorityPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PointPropertyTypes,
		stage.PointPropertyType_stagedOrder,
		stage.PointPropertyTypes_reference,
		&stage.PointPropertyTypes_referenceOrder,
		stage.PointPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PointReferencePropertyTypes,
		stage.PointReferencePropertyType_stagedOrder,
		stage.PointReferencePropertyTypes_reference,
		&stage.PointReferencePropertyTypes_referenceOrder,
		stage.PointReferencePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PointReferenceTypes,
		stage.PointReferenceType_stagedOrder,
		stage.PointReferenceTypes_reference,
		&stage.PointReferenceTypes_referenceOrder,
		stage.PointReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PointTypes,
		stage.PointType_stagedOrder,
		stage.PointTypes_reference,
		&stage.PointTypes_referenceOrder,
		stage.PointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureAvailabilityPropertyTypes,
		stage.ProcedureAvailabilityPropertyType_stagedOrder,
		stage.ProcedureAvailabilityPropertyTypes_reference,
		&stage.ProcedureAvailabilityPropertyTypes_referenceOrder,
		stage.ProcedureAvailabilityPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureAvailabilityTypes,
		stage.ProcedureAvailabilityType_stagedOrder,
		stage.ProcedureAvailabilityTypes_reference,
		&stage.ProcedureAvailabilityTypes_referenceOrder,
		stage.ProcedureAvailabilityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureDMETimeSlicePropertyTypes,
		stage.ProcedureDMETimeSlicePropertyType_stagedOrder,
		stage.ProcedureDMETimeSlicePropertyTypes_reference,
		&stage.ProcedureDMETimeSlicePropertyTypes_referenceOrder,
		stage.ProcedureDMETimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureDMETimeSliceTypes,
		stage.ProcedureDMETimeSliceType_stagedOrder,
		stage.ProcedureDMETimeSliceTypes_reference,
		&stage.ProcedureDMETimeSliceTypes_referenceOrder,
		stage.ProcedureDMETimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureDMETypes,
		stage.ProcedureDMEType_stagedOrder,
		stage.ProcedureDMETypes_reference,
		&stage.ProcedureDMETypes_referenceOrder,
		stage.ProcedureDMETypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedurePropertyTypes,
		stage.ProcedurePropertyType_stagedOrder,
		stage.ProcedurePropertyTypes_reference,
		&stage.ProcedurePropertyTypes_referenceOrder,
		stage.ProcedurePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureTransitionLegPropertyTypes,
		stage.ProcedureTransitionLegPropertyType_stagedOrder,
		stage.ProcedureTransitionLegPropertyTypes_reference,
		&stage.ProcedureTransitionLegPropertyTypes_referenceOrder,
		stage.ProcedureTransitionLegPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureTransitionLegTypes,
		stage.ProcedureTransitionLegType_stagedOrder,
		stage.ProcedureTransitionLegTypes_reference,
		&stage.ProcedureTransitionLegTypes_referenceOrder,
		stage.ProcedureTransitionLegTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureTransitionPropertyTypes,
		stage.ProcedureTransitionPropertyType_stagedOrder,
		stage.ProcedureTransitionPropertyTypes_reference,
		&stage.ProcedureTransitionPropertyTypes_referenceOrder,
		stage.ProcedureTransitionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcedureTransitionTypes,
		stage.ProcedureTransitionType_stagedOrder,
		stage.ProcedureTransitionTypes_reference,
		&stage.ProcedureTransitionTypes_referenceOrder,
		stage.ProcedureTransitionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadarSystemPropertyTypes,
		stage.RadarSystemPropertyType_stagedOrder,
		stage.RadarSystemPropertyTypes_reference,
		&stage.RadarSystemPropertyTypes_referenceOrder,
		stage.RadarSystemPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioFrequencyAreaPropertyTypes,
		stage.RadioFrequencyAreaPropertyType_stagedOrder,
		stage.RadioFrequencyAreaPropertyTypes_reference,
		&stage.RadioFrequencyAreaPropertyTypes_referenceOrder,
		stage.RadioFrequencyAreaPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RunwayCentrelinePointPropertyTypes,
		stage.RunwayCentrelinePointPropertyType_stagedOrder,
		stage.RunwayCentrelinePointPropertyTypes_reference,
		&stage.RunwayCentrelinePointPropertyTypes_referenceOrder,
		stage.RunwayCentrelinePointPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RunwayDirectionPropertyTypes,
		stage.RunwayDirectionPropertyType_stagedOrder,
		stage.RunwayDirectionPropertyTypes_reference,
		&stage.RunwayDirectionPropertyTypes_referenceOrder,
		stage.RunwayDirectionPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaPropertyTypes,
		stage.SafeAltitudeAreaPropertyType_stagedOrder,
		stage.SafeAltitudeAreaPropertyTypes_reference,
		&stage.SafeAltitudeAreaPropertyTypes_referenceOrder,
		stage.SafeAltitudeAreaPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaSectorPropertyTypes,
		stage.SafeAltitudeAreaSectorPropertyType_stagedOrder,
		stage.SafeAltitudeAreaSectorPropertyTypes_reference,
		&stage.SafeAltitudeAreaSectorPropertyTypes_referenceOrder,
		stage.SafeAltitudeAreaSectorPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaSectorTypes,
		stage.SafeAltitudeAreaSectorType_stagedOrder,
		stage.SafeAltitudeAreaSectorTypes_reference,
		&stage.SafeAltitudeAreaSectorTypes_referenceOrder,
		stage.SafeAltitudeAreaSectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaTimeSlicePropertyTypes,
		stage.SafeAltitudeAreaTimeSlicePropertyType_stagedOrder,
		stage.SafeAltitudeAreaTimeSlicePropertyTypes_reference,
		&stage.SafeAltitudeAreaTimeSlicePropertyTypes_referenceOrder,
		stage.SafeAltitudeAreaTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaTimeSliceTypes,
		stage.SafeAltitudeAreaTimeSliceType_stagedOrder,
		stage.SafeAltitudeAreaTimeSliceTypes_reference,
		&stage.SafeAltitudeAreaTimeSliceTypes_referenceOrder,
		stage.SafeAltitudeAreaTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SafeAltitudeAreaTypes,
		stage.SafeAltitudeAreaType_stagedOrder,
		stage.SafeAltitudeAreaTypes_reference,
		&stage.SafeAltitudeAreaTypes_referenceOrder,
		stage.SafeAltitudeAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SectorDesignPropertyTypes,
		stage.SectorDesignPropertyType_stagedOrder,
		stage.SectorDesignPropertyTypes_reference,
		&stage.SectorDesignPropertyTypes_referenceOrder,
		stage.SectorDesignPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SectorDesignTypes,
		stage.SectorDesignType_stagedOrder,
		stage.SectorDesignTypes_reference,
		&stage.SectorDesignTypes_referenceOrder,
		stage.SectorDesignTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SegmentLegPropertyTypes,
		stage.SegmentLegPropertyType_stagedOrder,
		stage.SegmentLegPropertyTypes_reference,
		&stage.SegmentLegPropertyTypes_referenceOrder,
		stage.SegmentLegPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SegmentPointPropertyTypes,
		stage.SegmentPointPropertyType_stagedOrder,
		stage.SegmentPointPropertyTypes_reference,
		&stage.SegmentPointPropertyTypes_referenceOrder,
		stage.SegmentPointPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SpecialNavigationSystemPropertyTypes,
		stage.SpecialNavigationSystemPropertyType_stagedOrder,
		stage.SpecialNavigationSystemPropertyTypes_reference,
		&stage.SpecialNavigationSystemPropertyTypes_referenceOrder,
		stage.SpecialNavigationSystemPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentArrivalPropertyTypes,
		stage.StandardInstrumentArrivalPropertyType_stagedOrder,
		stage.StandardInstrumentArrivalPropertyTypes_reference,
		&stage.StandardInstrumentArrivalPropertyTypes_referenceOrder,
		stage.StandardInstrumentArrivalPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentArrivalTimeSlicePropertyTypes,
		stage.StandardInstrumentArrivalTimeSlicePropertyType_stagedOrder,
		stage.StandardInstrumentArrivalTimeSlicePropertyTypes_reference,
		&stage.StandardInstrumentArrivalTimeSlicePropertyTypes_referenceOrder,
		stage.StandardInstrumentArrivalTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentArrivalTimeSliceTypes,
		stage.StandardInstrumentArrivalTimeSliceType_stagedOrder,
		stage.StandardInstrumentArrivalTimeSliceTypes_reference,
		&stage.StandardInstrumentArrivalTimeSliceTypes_referenceOrder,
		stage.StandardInstrumentArrivalTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentArrivalTypes,
		stage.StandardInstrumentArrivalType_stagedOrder,
		stage.StandardInstrumentArrivalTypes_reference,
		&stage.StandardInstrumentArrivalTypes_referenceOrder,
		stage.StandardInstrumentArrivalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentDeparturePropertyTypes,
		stage.StandardInstrumentDeparturePropertyType_stagedOrder,
		stage.StandardInstrumentDeparturePropertyTypes_reference,
		&stage.StandardInstrumentDeparturePropertyTypes_referenceOrder,
		stage.StandardInstrumentDeparturePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentDepartureTimeSlicePropertyTypes,
		stage.StandardInstrumentDepartureTimeSlicePropertyType_stagedOrder,
		stage.StandardInstrumentDepartureTimeSlicePropertyTypes_reference,
		&stage.StandardInstrumentDepartureTimeSlicePropertyTypes_referenceOrder,
		stage.StandardInstrumentDepartureTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentDepartureTimeSliceTypes,
		stage.StandardInstrumentDepartureTimeSliceType_stagedOrder,
		stage.StandardInstrumentDepartureTimeSliceTypes_reference,
		&stage.StandardInstrumentDepartureTimeSliceTypes_referenceOrder,
		stage.StandardInstrumentDepartureTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardInstrumentDepartureTypes,
		stage.StandardInstrumentDepartureType_stagedOrder,
		stage.StandardInstrumentDepartureTypes_reference,
		&stage.StandardInstrumentDepartureTypes_referenceOrder,
		stage.StandardInstrumentDepartureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SurfacePropertyTypes,
		stage.SurfacePropertyType_stagedOrder,
		stage.SurfacePropertyTypes_reference,
		&stage.SurfacePropertyTypes_referenceOrder,
		stage.SurfacePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SurfaceTypes,
		stage.SurfaceType_stagedOrder,
		stage.SurfaceTypes_reference,
		&stage.SurfaceTypes_referenceOrder,
		stage.SurfaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalArrivalAreaSectorPropertyTypes,
		stage.TerminalArrivalAreaSectorPropertyType_stagedOrder,
		stage.TerminalArrivalAreaSectorPropertyTypes_reference,
		&stage.TerminalArrivalAreaSectorPropertyTypes_referenceOrder,
		stage.TerminalArrivalAreaSectorPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalArrivalAreaSectorTypes,
		stage.TerminalArrivalAreaSectorType_stagedOrder,
		stage.TerminalArrivalAreaSectorTypes_reference,
		&stage.TerminalArrivalAreaSectorTypes_referenceOrder,
		stage.TerminalArrivalAreaSectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalArrivalAreaTimeSlicePropertyTypes,
		stage.TerminalArrivalAreaTimeSlicePropertyType_stagedOrder,
		stage.TerminalArrivalAreaTimeSlicePropertyTypes_reference,
		&stage.TerminalArrivalAreaTimeSlicePropertyTypes_referenceOrder,
		stage.TerminalArrivalAreaTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalArrivalAreaTimeSliceTypes,
		stage.TerminalArrivalAreaTimeSliceType_stagedOrder,
		stage.TerminalArrivalAreaTimeSliceTypes_reference,
		&stage.TerminalArrivalAreaTimeSliceTypes_referenceOrder,
		stage.TerminalArrivalAreaTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalArrivalAreaTypes,
		stage.TerminalArrivalAreaType_stagedOrder,
		stage.TerminalArrivalAreaTypes_reference,
		&stage.TerminalArrivalAreaTypes_referenceOrder,
		stage.TerminalArrivalAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalSegmentPointPropertyTypes,
		stage.TerminalSegmentPointPropertyType_stagedOrder,
		stage.TerminalSegmentPointPropertyTypes_reference,
		&stage.TerminalSegmentPointPropertyTypes_referenceOrder,
		stage.TerminalSegmentPointPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TerminalSegmentPointTypes,
		stage.TerminalSegmentPointType_stagedOrder,
		stage.TerminalSegmentPointTypes_reference,
		&stage.TerminalSegmentPointTypes_referenceOrder,
		stage.TerminalSegmentPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimesheetPropertyTypes,
		stage.TimesheetPropertyType_stagedOrder,
		stage.TimesheetPropertyTypes_reference,
		&stage.TimesheetPropertyTypes_referenceOrder,
		stage.TimesheetPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimesheetTypes,
		stage.TimesheetType_stagedOrder,
		stage.TimesheetTypes_reference,
		&stage.TimesheetTypes_referenceOrder,
		stage.TimesheetTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TouchDownLiftOffPropertyTypes,
		stage.TouchDownLiftOffPropertyType_stagedOrder,
		stage.TouchDownLiftOffPropertyTypes_reference,
		&stage.TouchDownLiftOffPropertyTypes_referenceOrder,
		stage.TouchDownLiftOffPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnplannedHoldingPropertyTypes,
		stage.UnplannedHoldingPropertyType_stagedOrder,
		stage.UnplannedHoldingPropertyTypes_reference,
		&stage.UnplannedHoldingPropertyTypes_referenceOrder,
		stage.UnplannedHoldingPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnplannedHoldingTimeSlicePropertyTypes,
		stage.UnplannedHoldingTimeSlicePropertyType_stagedOrder,
		stage.UnplannedHoldingTimeSlicePropertyTypes_reference,
		&stage.UnplannedHoldingTimeSlicePropertyTypes_referenceOrder,
		stage.UnplannedHoldingTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnplannedHoldingTimeSliceTypes,
		stage.UnplannedHoldingTimeSliceType_stagedOrder,
		stage.UnplannedHoldingTimeSliceTypes_reference,
		&stage.UnplannedHoldingTimeSliceTypes_referenceOrder,
		stage.UnplannedHoldingTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnplannedHoldingTypes,
		stage.UnplannedHoldingType_stagedOrder,
		stage.UnplannedHoldingTypes_reference,
		&stage.UnplannedHoldingTypes_referenceOrder,
		stage.UnplannedHoldingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.VerticalStructurePropertyTypes,
		stage.VerticalStructurePropertyType_stagedOrder,
		stage.VerticalStructurePropertyTypes_reference,
		&stage.VerticalStructurePropertyTypes_referenceOrder,
		stage.VerticalStructurePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.A_extensions, &stage.A_extensions_reference, &stage.A_extensions_referenceOrder, &stage.A_extensions_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_1s, &stage.A_extension_1s_reference, &stage.A_extension_1s_referenceOrder, &stage.A_extension_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_10s, &stage.A_extension_10s_reference, &stage.A_extension_10s_referenceOrder, &stage.A_extension_10s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_11s, &stage.A_extension_11s_reference, &stage.A_extension_11s_referenceOrder, &stage.A_extension_11s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_12s, &stage.A_extension_12s_reference, &stage.A_extension_12s_referenceOrder, &stage.A_extension_12s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_13s, &stage.A_extension_13s_reference, &stage.A_extension_13s_referenceOrder, &stage.A_extension_13s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_14s, &stage.A_extension_14s_reference, &stage.A_extension_14s_referenceOrder, &stage.A_extension_14s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_15s, &stage.A_extension_15s_reference, &stage.A_extension_15s_referenceOrder, &stage.A_extension_15s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_16s, &stage.A_extension_16s_reference, &stage.A_extension_16s_referenceOrder, &stage.A_extension_16s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_17s, &stage.A_extension_17s_reference, &stage.A_extension_17s_referenceOrder, &stage.A_extension_17s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_18s, &stage.A_extension_18s_reference, &stage.A_extension_18s_referenceOrder, &stage.A_extension_18s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_19s, &stage.A_extension_19s_reference, &stage.A_extension_19s_referenceOrder, &stage.A_extension_19s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_2s, &stage.A_extension_2s_reference, &stage.A_extension_2s_referenceOrder, &stage.A_extension_2s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_20s, &stage.A_extension_20s_reference, &stage.A_extension_20s_referenceOrder, &stage.A_extension_20s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_21s, &stage.A_extension_21s_reference, &stage.A_extension_21s_referenceOrder, &stage.A_extension_21s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_22s, &stage.A_extension_22s_reference, &stage.A_extension_22s_referenceOrder, &stage.A_extension_22s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_23s, &stage.A_extension_23s_reference, &stage.A_extension_23s_referenceOrder, &stage.A_extension_23s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_24s, &stage.A_extension_24s_reference, &stage.A_extension_24s_referenceOrder, &stage.A_extension_24s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_25s, &stage.A_extension_25s_reference, &stage.A_extension_25s_referenceOrder, &stage.A_extension_25s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_26s, &stage.A_extension_26s_reference, &stage.A_extension_26s_referenceOrder, &stage.A_extension_26s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_27s, &stage.A_extension_27s_reference, &stage.A_extension_27s_referenceOrder, &stage.A_extension_27s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_28s, &stage.A_extension_28s_reference, &stage.A_extension_28s_referenceOrder, &stage.A_extension_28s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_29s, &stage.A_extension_29s_reference, &stage.A_extension_29s_referenceOrder, &stage.A_extension_29s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_3s, &stage.A_extension_3s_reference, &stage.A_extension_3s_referenceOrder, &stage.A_extension_3s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_30s, &stage.A_extension_30s_reference, &stage.A_extension_30s_referenceOrder, &stage.A_extension_30s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_31s, &stage.A_extension_31s_reference, &stage.A_extension_31s_referenceOrder, &stage.A_extension_31s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_32s, &stage.A_extension_32s_reference, &stage.A_extension_32s_referenceOrder, &stage.A_extension_32s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_33s, &stage.A_extension_33s_reference, &stage.A_extension_33s_referenceOrder, &stage.A_extension_33s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_34s, &stage.A_extension_34s_reference, &stage.A_extension_34s_referenceOrder, &stage.A_extension_34s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_35s, &stage.A_extension_35s_reference, &stage.A_extension_35s_referenceOrder, &stage.A_extension_35s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_36s, &stage.A_extension_36s_reference, &stage.A_extension_36s_referenceOrder, &stage.A_extension_36s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_37s, &stage.A_extension_37s_reference, &stage.A_extension_37s_referenceOrder, &stage.A_extension_37s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_38s, &stage.A_extension_38s_reference, &stage.A_extension_38s_referenceOrder, &stage.A_extension_38s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_39s, &stage.A_extension_39s_reference, &stage.A_extension_39s_referenceOrder, &stage.A_extension_39s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_4s, &stage.A_extension_4s_reference, &stage.A_extension_4s_referenceOrder, &stage.A_extension_4s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_40s, &stage.A_extension_40s_reference, &stage.A_extension_40s_referenceOrder, &stage.A_extension_40s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_41s, &stage.A_extension_41s_reference, &stage.A_extension_41s_referenceOrder, &stage.A_extension_41s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_42s, &stage.A_extension_42s_reference, &stage.A_extension_42s_referenceOrder, &stage.A_extension_42s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_43s, &stage.A_extension_43s_reference, &stage.A_extension_43s_referenceOrder, &stage.A_extension_43s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_44s, &stage.A_extension_44s_reference, &stage.A_extension_44s_referenceOrder, &stage.A_extension_44s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_45s, &stage.A_extension_45s_reference, &stage.A_extension_45s_referenceOrder, &stage.A_extension_45s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_46s, &stage.A_extension_46s_reference, &stage.A_extension_46s_referenceOrder, &stage.A_extension_46s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_47s, &stage.A_extension_47s_reference, &stage.A_extension_47s_referenceOrder, &stage.A_extension_47s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_48s, &stage.A_extension_48s_reference, &stage.A_extension_48s_referenceOrder, &stage.A_extension_48s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_49s, &stage.A_extension_49s_reference, &stage.A_extension_49s_referenceOrder, &stage.A_extension_49s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_50s, &stage.A_extension_50s_reference, &stage.A_extension_50s_referenceOrder, &stage.A_extension_50s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_51s, &stage.A_extension_51s_reference, &stage.A_extension_51s_referenceOrder, &stage.A_extension_51s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_52s, &stage.A_extension_52s_reference, &stage.A_extension_52s_referenceOrder, &stage.A_extension_52s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_53s, &stage.A_extension_53s_reference, &stage.A_extension_53s_referenceOrder, &stage.A_extension_53s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_54s, &stage.A_extension_54s_reference, &stage.A_extension_54s_referenceOrder, &stage.A_extension_54s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_55s, &stage.A_extension_55s_reference, &stage.A_extension_55s_referenceOrder, &stage.A_extension_55s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_56s, &stage.A_extension_56s_reference, &stage.A_extension_56s_referenceOrder, &stage.A_extension_56s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_57s, &stage.A_extension_57s_reference, &stage.A_extension_57s_referenceOrder, &stage.A_extension_57s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_58s, &stage.A_extension_58s_reference, &stage.A_extension_58s_referenceOrder, &stage.A_extension_58s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractApproachLegTypes, &stage.AbstractApproachLegTypes_reference, &stage.AbstractApproachLegTypes_referenceOrder, &stage.AbstractApproachLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractProcedureTypes, &stage.AbstractProcedureTypes_reference, &stage.AbstractProcedureTypes_referenceOrder, &stage.AbstractProcedureTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractSegmentLegTypes, &stage.AbstractSegmentLegTypes_reference, &stage.AbstractSegmentLegTypes_referenceOrder, &stage.AbstractSegmentLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractSegmentPointTypes, &stage.AbstractSegmentPointTypes_reference, &stage.AbstractSegmentPointTypes_referenceOrder, &stage.AbstractSegmentPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingPointTypes, &stage.AerialRefuellingPointTypes_reference, &stage.AerialRefuellingPointTypes_referenceOrder, &stage.AerialRefuellingPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicPropertyTypes, &stage.AircraftCharacteristicPropertyTypes_reference, &stage.AircraftCharacteristicPropertyTypes_referenceOrder, &stage.AircraftCharacteristicPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicTypes, &stage.AircraftCharacteristicTypes_reference, &stage.AircraftCharacteristicTypes_referenceOrder, &stage.AircraftCharacteristicTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_referenceOrder, &stage.AirportHeliportPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AltimeterSourcePropertyTypes, &stage.AltimeterSourcePropertyTypes_reference, &stage.AltimeterSourcePropertyTypes_referenceOrder, &stage.AltimeterSourcePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AltitudeAdjustmentPropertyTypes, &stage.AltitudeAdjustmentPropertyTypes_reference, &stage.AltitudeAdjustmentPropertyTypes_referenceOrder, &stage.AltitudeAdjustmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AltitudeAdjustmentTypes, &stage.AltitudeAdjustmentTypes_reference, &stage.AltitudeAdjustmentTypes_referenceOrder, &stage.AltitudeAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleIndicationPropertyTypes, &stage.AngleIndicationPropertyTypes_reference, &stage.AngleIndicationPropertyTypes_referenceOrder, &stage.AngleIndicationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleUsePropertyTypes, &stage.AngleUsePropertyTypes_reference, &stage.AngleUsePropertyTypes_referenceOrder, &stage.AngleUsePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleUseTypes, &stage.AngleUseTypes_reference, &stage.AngleUseTypes_referenceOrder, &stage.AngleUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachAltitudeTablePropertyTypes, &stage.ApproachAltitudeTablePropertyTypes_reference, &stage.ApproachAltitudeTablePropertyTypes_referenceOrder, &stage.ApproachAltitudeTablePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachAltitudeTableTypes, &stage.ApproachAltitudeTableTypes_reference, &stage.ApproachAltitudeTableTypes_referenceOrder, &stage.ApproachAltitudeTableTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachConditionPropertyTypes, &stage.ApproachConditionPropertyTypes_reference, &stage.ApproachConditionPropertyTypes_referenceOrder, &stage.ApproachConditionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachConditionTypes, &stage.ApproachConditionTypes_reference, &stage.ApproachConditionTypes_referenceOrder, &stage.ApproachConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachDistanceTablePropertyTypes, &stage.ApproachDistanceTablePropertyTypes_reference, &stage.ApproachDistanceTablePropertyTypes_referenceOrder, &stage.ApproachDistanceTablePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachDistanceTableTypes, &stage.ApproachDistanceTableTypes_reference, &stage.ApproachDistanceTableTypes_referenceOrder, &stage.ApproachDistanceTableTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachTimingTablePropertyTypes, &stage.ApproachTimingTablePropertyTypes_reference, &stage.ApproachTimingTablePropertyTypes_referenceOrder, &stage.ApproachTimingTablePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ApproachTimingTableTypes, &stage.ApproachTimingTableTypes_reference, &stage.ApproachTimingTableTypes_referenceOrder, &stage.ApproachTimingTableTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalFeederLegTimeSlicePropertyTypes, &stage.ArrivalFeederLegTimeSlicePropertyTypes_reference, &stage.ArrivalFeederLegTimeSlicePropertyTypes_referenceOrder, &stage.ArrivalFeederLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalFeederLegTimeSliceTypes, &stage.ArrivalFeederLegTimeSliceTypes_reference, &stage.ArrivalFeederLegTimeSliceTypes_referenceOrder, &stage.ArrivalFeederLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalFeederLegTypes, &stage.ArrivalFeederLegTypes_reference, &stage.ArrivalFeederLegTypes_referenceOrder, &stage.ArrivalFeederLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalLegTimeSlicePropertyTypes, &stage.ArrivalLegTimeSlicePropertyTypes_reference, &stage.ArrivalLegTimeSlicePropertyTypes_referenceOrder, &stage.ArrivalLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalLegTimeSliceTypes, &stage.ArrivalLegTimeSliceTypes_reference, &stage.ArrivalLegTimeSliceTypes_referenceOrder, &stage.ArrivalLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ArrivalLegTypes, &stage.ArrivalLegTypes_reference, &stage.ArrivalLegTypes_referenceOrder, &stage.ArrivalLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.CircleSectorPropertyTypes, &stage.CircleSectorPropertyTypes_reference, &stage.CircleSectorPropertyTypes_referenceOrder, &stage.CircleSectorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CircleSectorTypes, &stage.CircleSectorTypes_reference, &stage.CircleSectorTypes_referenceOrder, &stage.CircleSectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CirclingAreaTimeSlicePropertyTypes, &stage.CirclingAreaTimeSlicePropertyTypes_reference, &stage.CirclingAreaTimeSlicePropertyTypes_referenceOrder, &stage.CirclingAreaTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CirclingAreaTimeSliceTypes, &stage.CirclingAreaTimeSliceTypes_reference, &stage.CirclingAreaTimeSliceTypes_referenceOrder, &stage.CirclingAreaTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CirclingAreaTypes, &stage.CirclingAreaTypes_reference, &stage.CirclingAreaTypes_referenceOrder, &stage.CirclingAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CirclingRestrictionPropertyTypes, &stage.CirclingRestrictionPropertyTypes_reference, &stage.CirclingRestrictionPropertyTypes_referenceOrder, &stage.CirclingRestrictionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CirclingRestrictionTypes, &stage.CirclingRestrictionTypes_reference, &stage.CirclingRestrictionTypes_referenceOrder, &stage.CirclingRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurvePropertyTypes, &stage.CurvePropertyTypes_reference, &stage.CurvePropertyTypes_referenceOrder, &stage.CurvePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurveTypes, &stage.CurveTypes_reference, &stage.CurveTypes_referenceOrder, &stage.CurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.DMEPropertyTypes, &stage.DMEPropertyTypes_reference, &stage.DMEPropertyTypes_referenceOrder, &stage.DMEPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DepartureArrivalConditionPropertyTypes, &stage.DepartureArrivalConditionPropertyTypes_reference, &stage.DepartureArrivalConditionPropertyTypes_referenceOrder, &stage.DepartureArrivalConditionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DepartureArrivalConditionTypes, &stage.DepartureArrivalConditionTypes_reference, &stage.DepartureArrivalConditionTypes_referenceOrder, &stage.DepartureArrivalConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.DepartureLegTimeSlicePropertyTypes, &stage.DepartureLegTimeSlicePropertyTypes_reference, &stage.DepartureLegTimeSlicePropertyTypes_referenceOrder, &stage.DepartureLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DepartureLegTimeSliceTypes, &stage.DepartureLegTimeSliceTypes_reference, &stage.DepartureLegTimeSliceTypes_referenceOrder, &stage.DepartureLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.DepartureLegTypes, &stage.DepartureLegTypes_reference, &stage.DepartureLegTypes_referenceOrder, &stage.DepartureLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointPropertyTypes, &stage.DesignatedPointPropertyTypes_reference, &stage.DesignatedPointPropertyTypes_referenceOrder, &stage.DesignatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DistanceIndicationPropertyTypes, &stage.DistanceIndicationPropertyTypes_reference, &stage.DistanceIndicationPropertyTypes_referenceOrder, &stage.DistanceIndicationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_referenceOrder, &stage.ElevatedCurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointTypes, &stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_referenceOrder, &stage.ElevatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_referenceOrder, &stage.ElevatedSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.EnRouteSegmentPointTypes, &stage.EnRouteSegmentPointTypes_reference, &stage.EnRouteSegmentPointTypes_referenceOrder, &stage.EnRouteSegmentPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.EquipmentUnavailableAdjustmentColumnPropertyTypes, &stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_reference, &stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_referenceOrder, &stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.EquipmentUnavailableAdjustmentColumnTypes, &stage.EquipmentUnavailableAdjustmentColumnTypes_reference, &stage.EquipmentUnavailableAdjustmentColumnTypes_referenceOrder, &stage.EquipmentUnavailableAdjustmentColumnTypes_instance)

	__gong__computeReferencePass1(stage, stage.EquipmentUnavailableAdjustmentPropertyTypes, &stage.EquipmentUnavailableAdjustmentPropertyTypes_reference, &stage.EquipmentUnavailableAdjustmentPropertyTypes_referenceOrder, &stage.EquipmentUnavailableAdjustmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.EquipmentUnavailableAdjustmentTypes, &stage.EquipmentUnavailableAdjustmentTypes_reference, &stage.EquipmentUnavailableAdjustmentTypes_referenceOrder, &stage.EquipmentUnavailableAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.FASDataBlockPropertyTypes, &stage.FASDataBlockPropertyTypes_reference, &stage.FASDataBlockPropertyTypes_referenceOrder, &stage.FASDataBlockPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FASDataBlockTypes, &stage.FASDataBlockTypes_reference, &stage.FASDataBlockTypes_referenceOrder, &stage.FASDataBlockTypes_instance)

	__gong__computeReferencePass1(stage, stage.FinalLegTimeSlicePropertyTypes, &stage.FinalLegTimeSlicePropertyTypes_reference, &stage.FinalLegTimeSlicePropertyTypes_referenceOrder, &stage.FinalLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FinalLegTimeSliceTypes, &stage.FinalLegTimeSliceTypes_reference, &stage.FinalLegTimeSliceTypes_referenceOrder, &stage.FinalLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.FinalLegTypes, &stage.FinalLegTypes_reference, &stage.FinalLegTypes_referenceOrder, &stage.FinalLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.FinalProfilePropertyTypes, &stage.FinalProfilePropertyTypes_reference, &stage.FinalProfilePropertyTypes_referenceOrder, &stage.FinalProfilePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FinalProfileTypes, &stage.FinalProfileTypes_reference, &stage.FinalProfileTypes_referenceOrder, &stage.FinalProfileTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingAssessmentTimeSlicePropertyTypes, &stage.HoldingAssessmentTimeSlicePropertyTypes_reference, &stage.HoldingAssessmentTimeSlicePropertyTypes_referenceOrder, &stage.HoldingAssessmentTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingAssessmentTimeSliceTypes, &stage.HoldingAssessmentTimeSliceTypes_reference, &stage.HoldingAssessmentTimeSliceTypes_referenceOrder, &stage.HoldingAssessmentTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingAssessmentTypes, &stage.HoldingAssessmentTypes_reference, &stage.HoldingAssessmentTypes_referenceOrder, &stage.HoldingAssessmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternDistancePropertyTypes, &stage.HoldingPatternDistancePropertyTypes_reference, &stage.HoldingPatternDistancePropertyTypes_referenceOrder, &stage.HoldingPatternDistancePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternDistanceTypes, &stage.HoldingPatternDistanceTypes_reference, &stage.HoldingPatternDistanceTypes_referenceOrder, &stage.HoldingPatternDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternDurationPropertyTypes, &stage.HoldingPatternDurationPropertyTypes_reference, &stage.HoldingPatternDurationPropertyTypes_referenceOrder, &stage.HoldingPatternDurationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternDurationTypes, &stage.HoldingPatternDurationTypes_reference, &stage.HoldingPatternDurationTypes_referenceOrder, &stage.HoldingPatternDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternPropertyTypes, &stage.HoldingPatternPropertyTypes_reference, &stage.HoldingPatternPropertyTypes_referenceOrder, &stage.HoldingPatternPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternTimeSlicePropertyTypes, &stage.HoldingPatternTimeSlicePropertyTypes_reference, &stage.HoldingPatternTimeSlicePropertyTypes_referenceOrder, &stage.HoldingPatternTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternTimeSliceTypes, &stage.HoldingPatternTimeSliceTypes_reference, &stage.HoldingPatternTimeSliceTypes_referenceOrder, &stage.HoldingPatternTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternTypes, &stage.HoldingPatternTypes_reference, &stage.HoldingPatternTypes_referenceOrder, &stage.HoldingPatternTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingUsePropertyTypes, &stage.HoldingUsePropertyTypes_reference, &stage.HoldingUsePropertyTypes_referenceOrder, &stage.HoldingUsePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingUseTypes, &stage.HoldingUseTypes_reference, &stage.HoldingUseTypes_referenceOrder, &stage.HoldingUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.InitialLegTimeSlicePropertyTypes, &stage.InitialLegTimeSlicePropertyTypes_reference, &stage.InitialLegTimeSlicePropertyTypes_referenceOrder, &stage.InitialLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.InitialLegTimeSliceTypes, &stage.InitialLegTimeSliceTypes_reference, &stage.InitialLegTimeSliceTypes_referenceOrder, &stage.InitialLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.InitialLegTypes, &stage.InitialLegTypes_reference, &stage.InitialLegTypes_referenceOrder, &stage.InitialLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentApproachProcedurePropertyTypes, &stage.InstrumentApproachProcedurePropertyTypes_reference, &stage.InstrumentApproachProcedurePropertyTypes_referenceOrder, &stage.InstrumentApproachProcedurePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentApproachProcedureTimeSlicePropertyTypes, &stage.InstrumentApproachProcedureTimeSlicePropertyTypes_reference, &stage.InstrumentApproachProcedureTimeSlicePropertyTypes_referenceOrder, &stage.InstrumentApproachProcedureTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentApproachProcedureTimeSliceTypes, &stage.InstrumentApproachProcedureTimeSliceTypes_reference, &stage.InstrumentApproachProcedureTimeSliceTypes_referenceOrder, &stage.InstrumentApproachProcedureTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentApproachProcedureTypes, &stage.InstrumentApproachProcedureTypes_reference, &stage.InstrumentApproachProcedureTypes_referenceOrder, &stage.InstrumentApproachProcedureTypes_instance)

	__gong__computeReferencePass1(stage, stage.IntermediateLegTimeSlicePropertyTypes, &stage.IntermediateLegTimeSlicePropertyTypes_reference, &stage.IntermediateLegTimeSlicePropertyTypes_referenceOrder, &stage.IntermediateLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.IntermediateLegTimeSliceTypes, &stage.IntermediateLegTimeSliceTypes_reference, &stage.IntermediateLegTimeSliceTypes_referenceOrder, &stage.IntermediateLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.IntermediateLegTypes, &stage.IntermediateLegTypes_reference, &stage.IntermediateLegTypes_referenceOrder, &stage.IntermediateLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.LandingTakeoffAreaCollectionPropertyTypes, &stage.LandingTakeoffAreaCollectionPropertyTypes_reference, &stage.LandingTakeoffAreaCollectionPropertyTypes_referenceOrder, &stage.LandingTakeoffAreaCollectionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LandingTakeoffAreaCollectionTypes, &stage.LandingTakeoffAreaCollectionTypes_reference, &stage.LandingTakeoffAreaCollectionTypes_referenceOrder, &stage.LandingTakeoffAreaCollectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.MinimaPropertyTypes, &stage.MinimaPropertyTypes_reference, &stage.MinimaPropertyTypes_referenceOrder, &stage.MinimaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MinimaTypes, &stage.MinimaTypes_reference, &stage.MinimaTypes_referenceOrder, &stage.MinimaTypes_instance)

	__gong__computeReferencePass1(stage, stage.MissedApproachGroupPropertyTypes, &stage.MissedApproachGroupPropertyTypes_reference, &stage.MissedApproachGroupPropertyTypes_referenceOrder, &stage.MissedApproachGroupPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MissedApproachGroupTypes, &stage.MissedApproachGroupTypes_reference, &stage.MissedApproachGroupTypes_referenceOrder, &stage.MissedApproachGroupTypes_instance)

	__gong__computeReferencePass1(stage, stage.MissedApproachLegTimeSlicePropertyTypes, &stage.MissedApproachLegTimeSlicePropertyTypes_reference, &stage.MissedApproachLegTimeSlicePropertyTypes_referenceOrder, &stage.MissedApproachLegTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MissedApproachLegTimeSliceTypes, &stage.MissedApproachLegTimeSliceTypes_reference, &stage.MissedApproachLegTimeSliceTypes_referenceOrder, &stage.MissedApproachLegTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.MissedApproachLegTypes, &stage.MissedApproachLegTypes_reference, &stage.MissedApproachLegTypes_referenceOrder, &stage.MissedApproachLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidPropertyTypes, &stage.NavaidPropertyTypes_reference, &stage.NavaidPropertyTypes_referenceOrder, &stage.NavaidPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaRestrictionTimeSlicePropertyTypes, &stage.NavigationAreaRestrictionTimeSlicePropertyTypes_reference, &stage.NavigationAreaRestrictionTimeSlicePropertyTypes_referenceOrder, &stage.NavigationAreaRestrictionTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaRestrictionTimeSliceTypes, &stage.NavigationAreaRestrictionTimeSliceTypes_reference, &stage.NavigationAreaRestrictionTimeSliceTypes_referenceOrder, &stage.NavigationAreaRestrictionTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaRestrictionTypes, &stage.NavigationAreaRestrictionTypes_reference, &stage.NavigationAreaRestrictionTypes_referenceOrder, &stage.NavigationAreaRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaSectorPropertyTypes, &stage.NavigationAreaSectorPropertyTypes_reference, &stage.NavigationAreaSectorPropertyTypes_referenceOrder, &stage.NavigationAreaSectorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaSectorTypes, &stage.NavigationAreaSectorTypes_reference, &stage.NavigationAreaSectorTypes_referenceOrder, &stage.NavigationAreaSectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaTimeSlicePropertyTypes, &stage.NavigationAreaTimeSlicePropertyTypes_reference, &stage.NavigationAreaTimeSlicePropertyTypes_referenceOrder, &stage.NavigationAreaTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaTimeSliceTypes, &stage.NavigationAreaTimeSliceTypes_reference, &stage.NavigationAreaTimeSliceTypes_referenceOrder, &stage.NavigationAreaTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavigationAreaTypes, &stage.NavigationAreaTypes_reference, &stage.NavigationAreaTypes_referenceOrder, &stage.NavigationAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAssessmentAreaPropertyTypes, &stage.ObstacleAssessmentAreaPropertyTypes_reference, &stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder, &stage.ObstacleAssessmentAreaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAssessmentAreaTypes, &stage.ObstacleAssessmentAreaTypes_reference, &stage.ObstacleAssessmentAreaTypes_referenceOrder, &stage.ObstacleAssessmentAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstaclePlacementPropertyTypes, &stage.ObstaclePlacementPropertyTypes_reference, &stage.ObstaclePlacementPropertyTypes_referenceOrder, &stage.ObstaclePlacementPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstaclePlacementTypes, &stage.ObstaclePlacementTypes_reference, &stage.ObstaclePlacementTypes_referenceOrder, &stage.ObstaclePlacementTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstructionPropertyTypes, &stage.ObstructionPropertyTypes_reference, &stage.ObstructionPropertyTypes_referenceOrder, &stage.ObstructionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstructionTypes, &stage.ObstructionTypes_reference, &stage.ObstructionTypes_referenceOrder, &stage.ObstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointPropertyTypes, &stage.PointPropertyTypes_reference, &stage.PointPropertyTypes_referenceOrder, &stage.PointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointReferencePropertyTypes, &stage.PointReferencePropertyTypes_reference, &stage.PointReferencePropertyTypes_referenceOrder, &stage.PointReferencePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointReferenceTypes, &stage.PointReferenceTypes_reference, &stage.PointReferenceTypes_referenceOrder, &stage.PointReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureAvailabilityPropertyTypes, &stage.ProcedureAvailabilityPropertyTypes_reference, &stage.ProcedureAvailabilityPropertyTypes_referenceOrder, &stage.ProcedureAvailabilityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureAvailabilityTypes, &stage.ProcedureAvailabilityTypes_reference, &stage.ProcedureAvailabilityTypes_referenceOrder, &stage.ProcedureAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureDMETimeSlicePropertyTypes, &stage.ProcedureDMETimeSlicePropertyTypes_reference, &stage.ProcedureDMETimeSlicePropertyTypes_referenceOrder, &stage.ProcedureDMETimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureDMETimeSliceTypes, &stage.ProcedureDMETimeSliceTypes_reference, &stage.ProcedureDMETimeSliceTypes_referenceOrder, &stage.ProcedureDMETimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureDMETypes, &stage.ProcedureDMETypes_reference, &stage.ProcedureDMETypes_referenceOrder, &stage.ProcedureDMETypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedurePropertyTypes, &stage.ProcedurePropertyTypes_reference, &stage.ProcedurePropertyTypes_referenceOrder, &stage.ProcedurePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureTransitionLegPropertyTypes, &stage.ProcedureTransitionLegPropertyTypes_reference, &stage.ProcedureTransitionLegPropertyTypes_referenceOrder, &stage.ProcedureTransitionLegPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureTransitionLegTypes, &stage.ProcedureTransitionLegTypes_reference, &stage.ProcedureTransitionLegTypes_referenceOrder, &stage.ProcedureTransitionLegTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureTransitionPropertyTypes, &stage.ProcedureTransitionPropertyTypes_reference, &stage.ProcedureTransitionPropertyTypes_referenceOrder, &stage.ProcedureTransitionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedureTransitionTypes, &stage.ProcedureTransitionTypes_reference, &stage.ProcedureTransitionTypes_referenceOrder, &stage.ProcedureTransitionTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarSystemPropertyTypes, &stage.RadarSystemPropertyTypes_reference, &stage.RadarSystemPropertyTypes_referenceOrder, &stage.RadarSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioFrequencyAreaPropertyTypes, &stage.RadioFrequencyAreaPropertyTypes_reference, &stage.RadioFrequencyAreaPropertyTypes_referenceOrder, &stage.RadioFrequencyAreaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayCentrelinePointPropertyTypes, &stage.RunwayCentrelinePointPropertyTypes_reference, &stage.RunwayCentrelinePointPropertyTypes_referenceOrder, &stage.RunwayCentrelinePointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayDirectionPropertyTypes, &stage.RunwayDirectionPropertyTypes_reference, &stage.RunwayDirectionPropertyTypes_referenceOrder, &stage.RunwayDirectionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaPropertyTypes, &stage.SafeAltitudeAreaPropertyTypes_reference, &stage.SafeAltitudeAreaPropertyTypes_referenceOrder, &stage.SafeAltitudeAreaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaSectorPropertyTypes, &stage.SafeAltitudeAreaSectorPropertyTypes_reference, &stage.SafeAltitudeAreaSectorPropertyTypes_referenceOrder, &stage.SafeAltitudeAreaSectorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaSectorTypes, &stage.SafeAltitudeAreaSectorTypes_reference, &stage.SafeAltitudeAreaSectorTypes_referenceOrder, &stage.SafeAltitudeAreaSectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaTimeSlicePropertyTypes, &stage.SafeAltitudeAreaTimeSlicePropertyTypes_reference, &stage.SafeAltitudeAreaTimeSlicePropertyTypes_referenceOrder, &stage.SafeAltitudeAreaTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaTimeSliceTypes, &stage.SafeAltitudeAreaTimeSliceTypes_reference, &stage.SafeAltitudeAreaTimeSliceTypes_referenceOrder, &stage.SafeAltitudeAreaTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SafeAltitudeAreaTypes, &stage.SafeAltitudeAreaTypes_reference, &stage.SafeAltitudeAreaTypes_referenceOrder, &stage.SafeAltitudeAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.SectorDesignPropertyTypes, &stage.SectorDesignPropertyTypes_reference, &stage.SectorDesignPropertyTypes_referenceOrder, &stage.SectorDesignPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SectorDesignTypes, &stage.SectorDesignTypes_reference, &stage.SectorDesignTypes_referenceOrder, &stage.SectorDesignTypes_instance)

	__gong__computeReferencePass1(stage, stage.SegmentLegPropertyTypes, &stage.SegmentLegPropertyTypes_reference, &stage.SegmentLegPropertyTypes_referenceOrder, &stage.SegmentLegPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SegmentPointPropertyTypes, &stage.SegmentPointPropertyTypes_reference, &stage.SegmentPointPropertyTypes_referenceOrder, &stage.SegmentPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationSystemPropertyTypes, &stage.SpecialNavigationSystemPropertyTypes_reference, &stage.SpecialNavigationSystemPropertyTypes_referenceOrder, &stage.SpecialNavigationSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentArrivalPropertyTypes, &stage.StandardInstrumentArrivalPropertyTypes_reference, &stage.StandardInstrumentArrivalPropertyTypes_referenceOrder, &stage.StandardInstrumentArrivalPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentArrivalTimeSlicePropertyTypes, &stage.StandardInstrumentArrivalTimeSlicePropertyTypes_reference, &stage.StandardInstrumentArrivalTimeSlicePropertyTypes_referenceOrder, &stage.StandardInstrumentArrivalTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentArrivalTimeSliceTypes, &stage.StandardInstrumentArrivalTimeSliceTypes_reference, &stage.StandardInstrumentArrivalTimeSliceTypes_referenceOrder, &stage.StandardInstrumentArrivalTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentArrivalTypes, &stage.StandardInstrumentArrivalTypes_reference, &stage.StandardInstrumentArrivalTypes_referenceOrder, &stage.StandardInstrumentArrivalTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentDeparturePropertyTypes, &stage.StandardInstrumentDeparturePropertyTypes_reference, &stage.StandardInstrumentDeparturePropertyTypes_referenceOrder, &stage.StandardInstrumentDeparturePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentDepartureTimeSlicePropertyTypes, &stage.StandardInstrumentDepartureTimeSlicePropertyTypes_reference, &stage.StandardInstrumentDepartureTimeSlicePropertyTypes_referenceOrder, &stage.StandardInstrumentDepartureTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentDepartureTimeSliceTypes, &stage.StandardInstrumentDepartureTimeSliceTypes_reference, &stage.StandardInstrumentDepartureTimeSliceTypes_referenceOrder, &stage.StandardInstrumentDepartureTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentDepartureTypes, &stage.StandardInstrumentDepartureTypes_reference, &stage.StandardInstrumentDepartureTypes_referenceOrder, &stage.StandardInstrumentDepartureTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_referenceOrder, &stage.SurfacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfaceTypes, &stage.SurfaceTypes_reference, &stage.SurfaceTypes_referenceOrder, &stage.SurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalArrivalAreaSectorPropertyTypes, &stage.TerminalArrivalAreaSectorPropertyTypes_reference, &stage.TerminalArrivalAreaSectorPropertyTypes_referenceOrder, &stage.TerminalArrivalAreaSectorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalArrivalAreaSectorTypes, &stage.TerminalArrivalAreaSectorTypes_reference, &stage.TerminalArrivalAreaSectorTypes_referenceOrder, &stage.TerminalArrivalAreaSectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalArrivalAreaTimeSlicePropertyTypes, &stage.TerminalArrivalAreaTimeSlicePropertyTypes_reference, &stage.TerminalArrivalAreaTimeSlicePropertyTypes_referenceOrder, &stage.TerminalArrivalAreaTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalArrivalAreaTimeSliceTypes, &stage.TerminalArrivalAreaTimeSliceTypes_reference, &stage.TerminalArrivalAreaTimeSliceTypes_referenceOrder, &stage.TerminalArrivalAreaTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalArrivalAreaTypes, &stage.TerminalArrivalAreaTypes_reference, &stage.TerminalArrivalAreaTypes_referenceOrder, &stage.TerminalArrivalAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalSegmentPointPropertyTypes, &stage.TerminalSegmentPointPropertyTypes_reference, &stage.TerminalSegmentPointPropertyTypes_referenceOrder, &stage.TerminalSegmentPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TerminalSegmentPointTypes, &stage.TerminalSegmentPointTypes_reference, &stage.TerminalSegmentPointTypes_referenceOrder, &stage.TerminalSegmentPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	__gong__computeReferencePass1(stage, stage.TouchDownLiftOffPropertyTypes, &stage.TouchDownLiftOffPropertyTypes_reference, &stage.TouchDownLiftOffPropertyTypes_referenceOrder, &stage.TouchDownLiftOffPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnplannedHoldingPropertyTypes, &stage.UnplannedHoldingPropertyTypes_reference, &stage.UnplannedHoldingPropertyTypes_referenceOrder, &stage.UnplannedHoldingPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnplannedHoldingTimeSlicePropertyTypes, &stage.UnplannedHoldingTimeSlicePropertyTypes_reference, &stage.UnplannedHoldingTimeSlicePropertyTypes_referenceOrder, &stage.UnplannedHoldingTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnplannedHoldingTimeSliceTypes, &stage.UnplannedHoldingTimeSliceTypes_reference, &stage.UnplannedHoldingTimeSliceTypes_referenceOrder, &stage.UnplannedHoldingTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnplannedHoldingTypes, &stage.UnplannedHoldingTypes_reference, &stage.UnplannedHoldingTypes_referenceOrder, &stage.UnplannedHoldingTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructurePropertyTypes, &stage.VerticalStructurePropertyTypes_reference, &stage.VerticalStructurePropertyTypes_referenceOrder, &stage.VerticalStructurePropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.A_extensions, stage.A_extensions_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_1s, stage.A_extension_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_10s, stage.A_extension_10s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_11s, stage.A_extension_11s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_12s, stage.A_extension_12s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_13s, stage.A_extension_13s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_14s, stage.A_extension_14s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_15s, stage.A_extension_15s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_16s, stage.A_extension_16s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_17s, stage.A_extension_17s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_18s, stage.A_extension_18s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_19s, stage.A_extension_19s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_2s, stage.A_extension_2s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_20s, stage.A_extension_20s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_21s, stage.A_extension_21s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_22s, stage.A_extension_22s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_23s, stage.A_extension_23s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_24s, stage.A_extension_24s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_25s, stage.A_extension_25s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_26s, stage.A_extension_26s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_27s, stage.A_extension_27s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_28s, stage.A_extension_28s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_29s, stage.A_extension_29s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_3s, stage.A_extension_3s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_30s, stage.A_extension_30s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_31s, stage.A_extension_31s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_32s, stage.A_extension_32s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_33s, stage.A_extension_33s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_34s, stage.A_extension_34s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_35s, stage.A_extension_35s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_36s, stage.A_extension_36s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_37s, stage.A_extension_37s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_38s, stage.A_extension_38s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_39s, stage.A_extension_39s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_4s, stage.A_extension_4s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_40s, stage.A_extension_40s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_41s, stage.A_extension_41s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_42s, stage.A_extension_42s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_43s, stage.A_extension_43s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_44s, stage.A_extension_44s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_45s, stage.A_extension_45s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_46s, stage.A_extension_46s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_47s, stage.A_extension_47s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_48s, stage.A_extension_48s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_49s, stage.A_extension_49s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_50s, stage.A_extension_50s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_51s, stage.A_extension_51s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_52s, stage.A_extension_52s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_53s, stage.A_extension_53s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_54s, stage.A_extension_54s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_55s, stage.A_extension_55s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_56s, stage.A_extension_56s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_57s, stage.A_extension_57s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_58s, stage.A_extension_58s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractApproachLegTypes, stage.AbstractApproachLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractProcedureTypes, stage.AbstractProcedureTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractSegmentLegTypes, stage.AbstractSegmentLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractSegmentPointTypes, stage.AbstractSegmentPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingPointTypes, stage.AerialRefuellingPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AltimeterSourcePropertyTypes, stage.AltimeterSourcePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AltitudeAdjustmentPropertyTypes, stage.AltitudeAdjustmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AltitudeAdjustmentTypes, stage.AltitudeAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleIndicationPropertyTypes, stage.AngleIndicationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleUsePropertyTypes, stage.AngleUsePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleUseTypes, stage.AngleUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachAltitudeTablePropertyTypes, stage.ApproachAltitudeTablePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachAltitudeTableTypes, stage.ApproachAltitudeTableTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachConditionPropertyTypes, stage.ApproachConditionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachConditionTypes, stage.ApproachConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachDistanceTablePropertyTypes, stage.ApproachDistanceTablePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachDistanceTableTypes, stage.ApproachDistanceTableTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachTimingTablePropertyTypes, stage.ApproachTimingTablePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ApproachTimingTableTypes, stage.ApproachTimingTableTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalFeederLegTimeSlicePropertyTypes, stage.ArrivalFeederLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalFeederLegTimeSliceTypes, stage.ArrivalFeederLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalFeederLegTypes, stage.ArrivalFeederLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalLegTimeSlicePropertyTypes, stage.ArrivalLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalLegTimeSliceTypes, stage.ArrivalLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ArrivalLegTypes, stage.ArrivalLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.CircleSectorPropertyTypes, stage.CircleSectorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CircleSectorTypes, stage.CircleSectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CirclingAreaTimeSlicePropertyTypes, stage.CirclingAreaTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CirclingAreaTimeSliceTypes, stage.CirclingAreaTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CirclingAreaTypes, stage.CirclingAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CirclingRestrictionPropertyTypes, stage.CirclingRestrictionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CirclingRestrictionTypes, stage.CirclingRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurvePropertyTypes, stage.CurvePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurveTypes, stage.CurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.DMEPropertyTypes, stage.DMEPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DepartureArrivalConditionPropertyTypes, stage.DepartureArrivalConditionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DepartureArrivalConditionTypes, stage.DepartureArrivalConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.DepartureLegTimeSlicePropertyTypes, stage.DepartureLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DepartureLegTimeSliceTypes, stage.DepartureLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.DepartureLegTypes, stage.DepartureLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointPropertyTypes, stage.DesignatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DistanceIndicationPropertyTypes, stage.DistanceIndicationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedCurveTypes, stage.ElevatedCurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointTypes, stage.ElevatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.EnRouteSegmentPointTypes, stage.EnRouteSegmentPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.EquipmentUnavailableAdjustmentColumnPropertyTypes, stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.EquipmentUnavailableAdjustmentColumnTypes, stage.EquipmentUnavailableAdjustmentColumnTypes_reference, stage)

	__gong__computeReferencePass2(stage.EquipmentUnavailableAdjustmentPropertyTypes, stage.EquipmentUnavailableAdjustmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.EquipmentUnavailableAdjustmentTypes, stage.EquipmentUnavailableAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.FASDataBlockPropertyTypes, stage.FASDataBlockPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FASDataBlockTypes, stage.FASDataBlockTypes_reference, stage)

	__gong__computeReferencePass2(stage.FinalLegTimeSlicePropertyTypes, stage.FinalLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FinalLegTimeSliceTypes, stage.FinalLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.FinalLegTypes, stage.FinalLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.FinalProfilePropertyTypes, stage.FinalProfilePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FinalProfileTypes, stage.FinalProfileTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingAssessmentTimeSlicePropertyTypes, stage.HoldingAssessmentTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingAssessmentTimeSliceTypes, stage.HoldingAssessmentTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingAssessmentTypes, stage.HoldingAssessmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternDistancePropertyTypes, stage.HoldingPatternDistancePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternDistanceTypes, stage.HoldingPatternDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternDurationPropertyTypes, stage.HoldingPatternDurationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternDurationTypes, stage.HoldingPatternDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternPropertyTypes, stage.HoldingPatternPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternTimeSlicePropertyTypes, stage.HoldingPatternTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternTimeSliceTypes, stage.HoldingPatternTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternTypes, stage.HoldingPatternTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingUsePropertyTypes, stage.HoldingUsePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingUseTypes, stage.HoldingUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.InitialLegTimeSlicePropertyTypes, stage.InitialLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.InitialLegTimeSliceTypes, stage.InitialLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.InitialLegTypes, stage.InitialLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentApproachProcedurePropertyTypes, stage.InstrumentApproachProcedurePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentApproachProcedureTimeSlicePropertyTypes, stage.InstrumentApproachProcedureTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentApproachProcedureTimeSliceTypes, stage.InstrumentApproachProcedureTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentApproachProcedureTypes, stage.InstrumentApproachProcedureTypes_reference, stage)

	__gong__computeReferencePass2(stage.IntermediateLegTimeSlicePropertyTypes, stage.IntermediateLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.IntermediateLegTimeSliceTypes, stage.IntermediateLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.IntermediateLegTypes, stage.IntermediateLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.LandingTakeoffAreaCollectionPropertyTypes, stage.LandingTakeoffAreaCollectionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LandingTakeoffAreaCollectionTypes, stage.LandingTakeoffAreaCollectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.MinimaPropertyTypes, stage.MinimaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MinimaTypes, stage.MinimaTypes_reference, stage)

	__gong__computeReferencePass2(stage.MissedApproachGroupPropertyTypes, stage.MissedApproachGroupPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MissedApproachGroupTypes, stage.MissedApproachGroupTypes_reference, stage)

	__gong__computeReferencePass2(stage.MissedApproachLegTimeSlicePropertyTypes, stage.MissedApproachLegTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MissedApproachLegTimeSliceTypes, stage.MissedApproachLegTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.MissedApproachLegTypes, stage.MissedApproachLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidPropertyTypes, stage.NavaidPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaRestrictionTimeSlicePropertyTypes, stage.NavigationAreaRestrictionTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaRestrictionTimeSliceTypes, stage.NavigationAreaRestrictionTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaRestrictionTypes, stage.NavigationAreaRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaSectorPropertyTypes, stage.NavigationAreaSectorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaSectorTypes, stage.NavigationAreaSectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaTimeSlicePropertyTypes, stage.NavigationAreaTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaTimeSliceTypes, stage.NavigationAreaTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavigationAreaTypes, stage.NavigationAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAssessmentAreaPropertyTypes, stage.ObstacleAssessmentAreaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAssessmentAreaTypes, stage.ObstacleAssessmentAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstaclePlacementPropertyTypes, stage.ObstaclePlacementPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstaclePlacementTypes, stage.ObstaclePlacementTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstructionPropertyTypes, stage.ObstructionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstructionTypes, stage.ObstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointPropertyTypes, stage.PointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointReferencePropertyTypes, stage.PointReferencePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointReferenceTypes, stage.PointReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureAvailabilityPropertyTypes, stage.ProcedureAvailabilityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureAvailabilityTypes, stage.ProcedureAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureDMETimeSlicePropertyTypes, stage.ProcedureDMETimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureDMETimeSliceTypes, stage.ProcedureDMETimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureDMETypes, stage.ProcedureDMETypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedurePropertyTypes, stage.ProcedurePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureTransitionLegPropertyTypes, stage.ProcedureTransitionLegPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureTransitionLegTypes, stage.ProcedureTransitionLegTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureTransitionPropertyTypes, stage.ProcedureTransitionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedureTransitionTypes, stage.ProcedureTransitionTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarSystemPropertyTypes, stage.RadarSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioFrequencyAreaPropertyTypes, stage.RadioFrequencyAreaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayCentrelinePointPropertyTypes, stage.RunwayCentrelinePointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaPropertyTypes, stage.SafeAltitudeAreaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaSectorPropertyTypes, stage.SafeAltitudeAreaSectorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaSectorTypes, stage.SafeAltitudeAreaSectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaTimeSlicePropertyTypes, stage.SafeAltitudeAreaTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaTimeSliceTypes, stage.SafeAltitudeAreaTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SafeAltitudeAreaTypes, stage.SafeAltitudeAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.SectorDesignPropertyTypes, stage.SectorDesignPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SectorDesignTypes, stage.SectorDesignTypes_reference, stage)

	__gong__computeReferencePass2(stage.SegmentLegPropertyTypes, stage.SegmentLegPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SegmentPointPropertyTypes, stage.SegmentPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationSystemPropertyTypes, stage.SpecialNavigationSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentArrivalPropertyTypes, stage.StandardInstrumentArrivalPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentArrivalTimeSlicePropertyTypes, stage.StandardInstrumentArrivalTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentArrivalTimeSliceTypes, stage.StandardInstrumentArrivalTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentArrivalTypes, stage.StandardInstrumentArrivalTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentDeparturePropertyTypes, stage.StandardInstrumentDeparturePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentDepartureTimeSlicePropertyTypes, stage.StandardInstrumentDepartureTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentDepartureTimeSliceTypes, stage.StandardInstrumentDepartureTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentDepartureTypes, stage.StandardInstrumentDepartureTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfacePropertyTypes, stage.SurfacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfaceTypes, stage.SurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalArrivalAreaSectorPropertyTypes, stage.TerminalArrivalAreaSectorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalArrivalAreaSectorTypes, stage.TerminalArrivalAreaSectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalArrivalAreaTimeSlicePropertyTypes, stage.TerminalArrivalAreaTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalArrivalAreaTimeSliceTypes, stage.TerminalArrivalAreaTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalArrivalAreaTypes, stage.TerminalArrivalAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalSegmentPointPropertyTypes, stage.TerminalSegmentPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TerminalSegmentPointTypes, stage.TerminalSegmentPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	__gong__computeReferencePass2(stage.TouchDownLiftOffPropertyTypes, stage.TouchDownLiftOffPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnplannedHoldingPropertyTypes, stage.UnplannedHoldingPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnplannedHoldingTimeSlicePropertyTypes, stage.UnplannedHoldingTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnplannedHoldingTimeSliceTypes, stage.UnplannedHoldingTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnplannedHoldingTypes, stage.UnplannedHoldingTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (a_extension *A_extension) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_stagedOrder, stage.A_extensions_referenceOrder, a_extension, "A_extension")
}

func (a_extension_1 *A_extension_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_1_stagedOrder, stage.A_extension_1s_referenceOrder, a_extension_1, "A_extension_1")
}

func (a_extension_10 *A_extension_10) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_10_stagedOrder, stage.A_extension_10s_referenceOrder, a_extension_10, "A_extension_10")
}

func (a_extension_11 *A_extension_11) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_11_stagedOrder, stage.A_extension_11s_referenceOrder, a_extension_11, "A_extension_11")
}

func (a_extension_12 *A_extension_12) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_12_stagedOrder, stage.A_extension_12s_referenceOrder, a_extension_12, "A_extension_12")
}

func (a_extension_13 *A_extension_13) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_13_stagedOrder, stage.A_extension_13s_referenceOrder, a_extension_13, "A_extension_13")
}

func (a_extension_14 *A_extension_14) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_14_stagedOrder, stage.A_extension_14s_referenceOrder, a_extension_14, "A_extension_14")
}

func (a_extension_15 *A_extension_15) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_15_stagedOrder, stage.A_extension_15s_referenceOrder, a_extension_15, "A_extension_15")
}

func (a_extension_16 *A_extension_16) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_16_stagedOrder, stage.A_extension_16s_referenceOrder, a_extension_16, "A_extension_16")
}

func (a_extension_17 *A_extension_17) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_17_stagedOrder, stage.A_extension_17s_referenceOrder, a_extension_17, "A_extension_17")
}

func (a_extension_18 *A_extension_18) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_18_stagedOrder, stage.A_extension_18s_referenceOrder, a_extension_18, "A_extension_18")
}

func (a_extension_19 *A_extension_19) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_19_stagedOrder, stage.A_extension_19s_referenceOrder, a_extension_19, "A_extension_19")
}

func (a_extension_2 *A_extension_2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_2_stagedOrder, stage.A_extension_2s_referenceOrder, a_extension_2, "A_extension_2")
}

func (a_extension_20 *A_extension_20) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_20_stagedOrder, stage.A_extension_20s_referenceOrder, a_extension_20, "A_extension_20")
}

func (a_extension_21 *A_extension_21) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_21_stagedOrder, stage.A_extension_21s_referenceOrder, a_extension_21, "A_extension_21")
}

func (a_extension_22 *A_extension_22) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_22_stagedOrder, stage.A_extension_22s_referenceOrder, a_extension_22, "A_extension_22")
}

func (a_extension_23 *A_extension_23) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_23_stagedOrder, stage.A_extension_23s_referenceOrder, a_extension_23, "A_extension_23")
}

func (a_extension_24 *A_extension_24) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_24_stagedOrder, stage.A_extension_24s_referenceOrder, a_extension_24, "A_extension_24")
}

func (a_extension_25 *A_extension_25) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_25_stagedOrder, stage.A_extension_25s_referenceOrder, a_extension_25, "A_extension_25")
}

func (a_extension_26 *A_extension_26) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_26_stagedOrder, stage.A_extension_26s_referenceOrder, a_extension_26, "A_extension_26")
}

func (a_extension_27 *A_extension_27) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_27_stagedOrder, stage.A_extension_27s_referenceOrder, a_extension_27, "A_extension_27")
}

func (a_extension_28 *A_extension_28) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_28_stagedOrder, stage.A_extension_28s_referenceOrder, a_extension_28, "A_extension_28")
}

func (a_extension_29 *A_extension_29) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_29_stagedOrder, stage.A_extension_29s_referenceOrder, a_extension_29, "A_extension_29")
}

func (a_extension_3 *A_extension_3) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_3_stagedOrder, stage.A_extension_3s_referenceOrder, a_extension_3, "A_extension_3")
}

func (a_extension_30 *A_extension_30) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_30_stagedOrder, stage.A_extension_30s_referenceOrder, a_extension_30, "A_extension_30")
}

func (a_extension_31 *A_extension_31) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_31_stagedOrder, stage.A_extension_31s_referenceOrder, a_extension_31, "A_extension_31")
}

func (a_extension_32 *A_extension_32) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_32_stagedOrder, stage.A_extension_32s_referenceOrder, a_extension_32, "A_extension_32")
}

func (a_extension_33 *A_extension_33) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_33_stagedOrder, stage.A_extension_33s_referenceOrder, a_extension_33, "A_extension_33")
}

func (a_extension_34 *A_extension_34) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_34_stagedOrder, stage.A_extension_34s_referenceOrder, a_extension_34, "A_extension_34")
}

func (a_extension_35 *A_extension_35) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_35_stagedOrder, stage.A_extension_35s_referenceOrder, a_extension_35, "A_extension_35")
}

func (a_extension_36 *A_extension_36) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_36_stagedOrder, stage.A_extension_36s_referenceOrder, a_extension_36, "A_extension_36")
}

func (a_extension_37 *A_extension_37) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_37_stagedOrder, stage.A_extension_37s_referenceOrder, a_extension_37, "A_extension_37")
}

func (a_extension_38 *A_extension_38) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_38_stagedOrder, stage.A_extension_38s_referenceOrder, a_extension_38, "A_extension_38")
}

func (a_extension_39 *A_extension_39) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_39_stagedOrder, stage.A_extension_39s_referenceOrder, a_extension_39, "A_extension_39")
}

func (a_extension_4 *A_extension_4) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_4_stagedOrder, stage.A_extension_4s_referenceOrder, a_extension_4, "A_extension_4")
}

func (a_extension_40 *A_extension_40) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_40_stagedOrder, stage.A_extension_40s_referenceOrder, a_extension_40, "A_extension_40")
}

func (a_extension_41 *A_extension_41) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_41_stagedOrder, stage.A_extension_41s_referenceOrder, a_extension_41, "A_extension_41")
}

func (a_extension_42 *A_extension_42) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_42_stagedOrder, stage.A_extension_42s_referenceOrder, a_extension_42, "A_extension_42")
}

func (a_extension_43 *A_extension_43) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_43_stagedOrder, stage.A_extension_43s_referenceOrder, a_extension_43, "A_extension_43")
}

func (a_extension_44 *A_extension_44) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_44_stagedOrder, stage.A_extension_44s_referenceOrder, a_extension_44, "A_extension_44")
}

func (a_extension_45 *A_extension_45) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_45_stagedOrder, stage.A_extension_45s_referenceOrder, a_extension_45, "A_extension_45")
}

func (a_extension_46 *A_extension_46) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_46_stagedOrder, stage.A_extension_46s_referenceOrder, a_extension_46, "A_extension_46")
}

func (a_extension_47 *A_extension_47) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_47_stagedOrder, stage.A_extension_47s_referenceOrder, a_extension_47, "A_extension_47")
}

func (a_extension_48 *A_extension_48) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_48_stagedOrder, stage.A_extension_48s_referenceOrder, a_extension_48, "A_extension_48")
}

func (a_extension_49 *A_extension_49) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_49_stagedOrder, stage.A_extension_49s_referenceOrder, a_extension_49, "A_extension_49")
}

func (a_extension_5 *A_extension_5) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_5_stagedOrder, stage.A_extension_5s_referenceOrder, a_extension_5, "A_extension_5")
}

func (a_extension_50 *A_extension_50) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_50_stagedOrder, stage.A_extension_50s_referenceOrder, a_extension_50, "A_extension_50")
}

func (a_extension_51 *A_extension_51) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_51_stagedOrder, stage.A_extension_51s_referenceOrder, a_extension_51, "A_extension_51")
}

func (a_extension_52 *A_extension_52) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_52_stagedOrder, stage.A_extension_52s_referenceOrder, a_extension_52, "A_extension_52")
}

func (a_extension_53 *A_extension_53) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_53_stagedOrder, stage.A_extension_53s_referenceOrder, a_extension_53, "A_extension_53")
}

func (a_extension_54 *A_extension_54) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_54_stagedOrder, stage.A_extension_54s_referenceOrder, a_extension_54, "A_extension_54")
}

func (a_extension_55 *A_extension_55) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_55_stagedOrder, stage.A_extension_55s_referenceOrder, a_extension_55, "A_extension_55")
}

func (a_extension_56 *A_extension_56) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_56_stagedOrder, stage.A_extension_56s_referenceOrder, a_extension_56, "A_extension_56")
}

func (a_extension_57 *A_extension_57) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_57_stagedOrder, stage.A_extension_57s_referenceOrder, a_extension_57, "A_extension_57")
}

func (a_extension_58 *A_extension_58) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_58_stagedOrder, stage.A_extension_58s_referenceOrder, a_extension_58, "A_extension_58")
}

func (a_extension_6 *A_extension_6) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_6_stagedOrder, stage.A_extension_6s_referenceOrder, a_extension_6, "A_extension_6")
}

func (a_extension_7 *A_extension_7) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_7_stagedOrder, stage.A_extension_7s_referenceOrder, a_extension_7, "A_extension_7")
}

func (a_extension_8 *A_extension_8) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_8_stagedOrder, stage.A_extension_8s_referenceOrder, a_extension_8, "A_extension_8")
}

func (a_extension_9 *A_extension_9) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_9_stagedOrder, stage.A_extension_9s_referenceOrder, a_extension_9, "A_extension_9")
}

func (abstractapproachlegtype *AbstractApproachLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractApproachLegType_stagedOrder, stage.AbstractApproachLegTypes_referenceOrder, abstractapproachlegtype, "AbstractApproachLegType")
}

func (abstractproceduretype *AbstractProcedureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractProcedureType_stagedOrder, stage.AbstractProcedureTypes_referenceOrder, abstractproceduretype, "AbstractProcedureType")
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractSegmentLegType_stagedOrder, stage.AbstractSegmentLegTypes_referenceOrder, abstractsegmentlegtype, "AbstractSegmentLegType")
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractSegmentPointType_stagedOrder, stage.AbstractSegmentPointTypes_referenceOrder, abstractsegmentpointtype, "AbstractSegmentPointType")
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingPointType_stagedOrder, stage.AerialRefuellingPointTypes_referenceOrder, aerialrefuellingpointtype, "AerialRefuellingPointType")
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicPropertyType_stagedOrder, stage.AircraftCharacteristicPropertyTypes_referenceOrder, aircraftcharacteristicpropertytype, "AircraftCharacteristicPropertyType")
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicType_stagedOrder, stage.AircraftCharacteristicTypes_referenceOrder, aircraftcharacteristictype, "AircraftCharacteristicType")
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyTypes_referenceOrder, airportheliportpropertytype, "AirportHeliportPropertyType")
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AltimeterSourcePropertyType_stagedOrder, stage.AltimeterSourcePropertyTypes_referenceOrder, altimetersourcepropertytype, "AltimeterSourcePropertyType")
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AltitudeAdjustmentPropertyType_stagedOrder, stage.AltitudeAdjustmentPropertyTypes_referenceOrder, altitudeadjustmentpropertytype, "AltitudeAdjustmentPropertyType")
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AltitudeAdjustmentType_stagedOrder, stage.AltitudeAdjustmentTypes_referenceOrder, altitudeadjustmenttype, "AltitudeAdjustmentType")
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleIndicationPropertyType_stagedOrder, stage.AngleIndicationPropertyTypes_referenceOrder, angleindicationpropertytype, "AngleIndicationPropertyType")
}

func (angleusepropertytype *AngleUsePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleUsePropertyType_stagedOrder, stage.AngleUsePropertyTypes_referenceOrder, angleusepropertytype, "AngleUsePropertyType")
}

func (angleusetype *AngleUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleUseType_stagedOrder, stage.AngleUseTypes_referenceOrder, angleusetype, "AngleUseType")
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachAltitudeTablePropertyType_stagedOrder, stage.ApproachAltitudeTablePropertyTypes_referenceOrder, approachaltitudetablepropertytype, "ApproachAltitudeTablePropertyType")
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachAltitudeTableType_stagedOrder, stage.ApproachAltitudeTableTypes_referenceOrder, approachaltitudetabletype, "ApproachAltitudeTableType")
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachConditionPropertyType_stagedOrder, stage.ApproachConditionPropertyTypes_referenceOrder, approachconditionpropertytype, "ApproachConditionPropertyType")
}

func (approachconditiontype *ApproachConditionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachConditionType_stagedOrder, stage.ApproachConditionTypes_referenceOrder, approachconditiontype, "ApproachConditionType")
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachDistanceTablePropertyType_stagedOrder, stage.ApproachDistanceTablePropertyTypes_referenceOrder, approachdistancetablepropertytype, "ApproachDistanceTablePropertyType")
}

func (approachdistancetabletype *ApproachDistanceTableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachDistanceTableType_stagedOrder, stage.ApproachDistanceTableTypes_referenceOrder, approachdistancetabletype, "ApproachDistanceTableType")
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachTimingTablePropertyType_stagedOrder, stage.ApproachTimingTablePropertyTypes_referenceOrder, approachtimingtablepropertytype, "ApproachTimingTablePropertyType")
}

func (approachtimingtabletype *ApproachTimingTableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ApproachTimingTableType_stagedOrder, stage.ApproachTimingTableTypes_referenceOrder, approachtimingtabletype, "ApproachTimingTableType")
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalFeederLegTimeSlicePropertyType_stagedOrder, stage.ArrivalFeederLegTimeSlicePropertyTypes_referenceOrder, arrivalfeederlegtimeslicepropertytype, "ArrivalFeederLegTimeSlicePropertyType")
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalFeederLegTimeSliceType_stagedOrder, stage.ArrivalFeederLegTimeSliceTypes_referenceOrder, arrivalfeederlegtimeslicetype, "ArrivalFeederLegTimeSliceType")
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalFeederLegType_stagedOrder, stage.ArrivalFeederLegTypes_referenceOrder, arrivalfeederlegtype, "ArrivalFeederLegType")
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalLegTimeSlicePropertyType_stagedOrder, stage.ArrivalLegTimeSlicePropertyTypes_referenceOrder, arrivallegtimeslicepropertytype, "ArrivalLegTimeSlicePropertyType")
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalLegTimeSliceType_stagedOrder, stage.ArrivalLegTimeSliceTypes_referenceOrder, arrivallegtimeslicetype, "ArrivalLegTimeSliceType")
}

func (arrivallegtype *ArrivalLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArrivalLegType_stagedOrder, stage.ArrivalLegTypes_referenceOrder, arrivallegtype, "ArrivalLegType")
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CircleSectorPropertyType_stagedOrder, stage.CircleSectorPropertyTypes_referenceOrder, circlesectorpropertytype, "CircleSectorPropertyType")
}

func (circlesectortype *CircleSectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CircleSectorType_stagedOrder, stage.CircleSectorTypes_referenceOrder, circlesectortype, "CircleSectorType")
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CirclingAreaTimeSlicePropertyType_stagedOrder, stage.CirclingAreaTimeSlicePropertyTypes_referenceOrder, circlingareatimeslicepropertytype, "CirclingAreaTimeSlicePropertyType")
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CirclingAreaTimeSliceType_stagedOrder, stage.CirclingAreaTimeSliceTypes_referenceOrder, circlingareatimeslicetype, "CirclingAreaTimeSliceType")
}

func (circlingareatype *CirclingAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CirclingAreaType_stagedOrder, stage.CirclingAreaTypes_referenceOrder, circlingareatype, "CirclingAreaType")
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CirclingRestrictionPropertyType_stagedOrder, stage.CirclingRestrictionPropertyTypes_referenceOrder, circlingrestrictionpropertytype, "CirclingRestrictionPropertyType")
}

func (circlingrestrictiontype *CirclingRestrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CirclingRestrictionType_stagedOrder, stage.CirclingRestrictionTypes_referenceOrder, circlingrestrictiontype, "CirclingRestrictionType")
}

func (curvepropertytype *CurvePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurvePropertyType_stagedOrder, stage.CurvePropertyTypes_referenceOrder, curvepropertytype, "CurvePropertyType")
}

func (curvetype *CurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurveType_stagedOrder, stage.CurveTypes_referenceOrder, curvetype, "CurveType")
}

func (dmepropertytype *DMEPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DMEPropertyType_stagedOrder, stage.DMEPropertyTypes_referenceOrder, dmepropertytype, "DMEPropertyType")
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DepartureArrivalConditionPropertyType_stagedOrder, stage.DepartureArrivalConditionPropertyTypes_referenceOrder, departurearrivalconditionpropertytype, "DepartureArrivalConditionPropertyType")
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DepartureArrivalConditionType_stagedOrder, stage.DepartureArrivalConditionTypes_referenceOrder, departurearrivalconditiontype, "DepartureArrivalConditionType")
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DepartureLegTimeSlicePropertyType_stagedOrder, stage.DepartureLegTimeSlicePropertyTypes_referenceOrder, departurelegtimeslicepropertytype, "DepartureLegTimeSlicePropertyType")
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DepartureLegTimeSliceType_stagedOrder, stage.DepartureLegTimeSliceTypes_referenceOrder, departurelegtimeslicetype, "DepartureLegTimeSliceType")
}

func (departurelegtype *DepartureLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DepartureLegType_stagedOrder, stage.DepartureLegTypes_referenceOrder, departurelegtype, "DepartureLegType")
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointPropertyType_stagedOrder, stage.DesignatedPointPropertyTypes_referenceOrder, designatedpointpropertytype, "DesignatedPointPropertyType")
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DistanceIndicationPropertyType_stagedOrder, stage.DistanceIndicationPropertyTypes_referenceOrder, distanceindicationpropertytype, "DistanceIndicationPropertyType")
}

func (elevatedcurvetype *ElevatedCurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveTypes_referenceOrder, elevatedcurvetype, "ElevatedCurveType")
}

func (elevatedpointtype *ElevatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointType_stagedOrder, stage.ElevatedPointTypes_referenceOrder, elevatedpointtype, "ElevatedPointType")
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceTypes_referenceOrder, elevatedsurfacetype, "ElevatedSurfaceType")
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EnRouteSegmentPointType_stagedOrder, stage.EnRouteSegmentPointTypes_referenceOrder, enroutesegmentpointtype, "EnRouteSegmentPointType")
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EquipmentUnavailableAdjustmentColumnPropertyType_stagedOrder, stage.EquipmentUnavailableAdjustmentColumnPropertyTypes_referenceOrder, equipmentunavailableadjustmentcolumnpropertytype, "EquipmentUnavailableAdjustmentColumnPropertyType")
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EquipmentUnavailableAdjustmentColumnType_stagedOrder, stage.EquipmentUnavailableAdjustmentColumnTypes_referenceOrder, equipmentunavailableadjustmentcolumntype, "EquipmentUnavailableAdjustmentColumnType")
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EquipmentUnavailableAdjustmentPropertyType_stagedOrder, stage.EquipmentUnavailableAdjustmentPropertyTypes_referenceOrder, equipmentunavailableadjustmentpropertytype, "EquipmentUnavailableAdjustmentPropertyType")
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EquipmentUnavailableAdjustmentType_stagedOrder, stage.EquipmentUnavailableAdjustmentTypes_referenceOrder, equipmentunavailableadjustmenttype, "EquipmentUnavailableAdjustmentType")
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FASDataBlockPropertyType_stagedOrder, stage.FASDataBlockPropertyTypes_referenceOrder, fasdatablockpropertytype, "FASDataBlockPropertyType")
}

func (fasdatablocktype *FASDataBlockType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FASDataBlockType_stagedOrder, stage.FASDataBlockTypes_referenceOrder, fasdatablocktype, "FASDataBlockType")
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FinalLegTimeSlicePropertyType_stagedOrder, stage.FinalLegTimeSlicePropertyTypes_referenceOrder, finallegtimeslicepropertytype, "FinalLegTimeSlicePropertyType")
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FinalLegTimeSliceType_stagedOrder, stage.FinalLegTimeSliceTypes_referenceOrder, finallegtimeslicetype, "FinalLegTimeSliceType")
}

func (finallegtype *FinalLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FinalLegType_stagedOrder, stage.FinalLegTypes_referenceOrder, finallegtype, "FinalLegType")
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FinalProfilePropertyType_stagedOrder, stage.FinalProfilePropertyTypes_referenceOrder, finalprofilepropertytype, "FinalProfilePropertyType")
}

func (finalprofiletype *FinalProfileType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FinalProfileType_stagedOrder, stage.FinalProfileTypes_referenceOrder, finalprofiletype, "FinalProfileType")
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingAssessmentTimeSlicePropertyType_stagedOrder, stage.HoldingAssessmentTimeSlicePropertyTypes_referenceOrder, holdingassessmenttimeslicepropertytype, "HoldingAssessmentTimeSlicePropertyType")
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingAssessmentTimeSliceType_stagedOrder, stage.HoldingAssessmentTimeSliceTypes_referenceOrder, holdingassessmenttimeslicetype, "HoldingAssessmentTimeSliceType")
}

func (holdingassessmenttype *HoldingAssessmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingAssessmentType_stagedOrder, stage.HoldingAssessmentTypes_referenceOrder, holdingassessmenttype, "HoldingAssessmentType")
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternDistancePropertyType_stagedOrder, stage.HoldingPatternDistancePropertyTypes_referenceOrder, holdingpatterndistancepropertytype, "HoldingPatternDistancePropertyType")
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternDistanceType_stagedOrder, stage.HoldingPatternDistanceTypes_referenceOrder, holdingpatterndistancetype, "HoldingPatternDistanceType")
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternDurationPropertyType_stagedOrder, stage.HoldingPatternDurationPropertyTypes_referenceOrder, holdingpatterndurationpropertytype, "HoldingPatternDurationPropertyType")
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternDurationType_stagedOrder, stage.HoldingPatternDurationTypes_referenceOrder, holdingpatterndurationtype, "HoldingPatternDurationType")
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternPropertyType_stagedOrder, stage.HoldingPatternPropertyTypes_referenceOrder, holdingpatternpropertytype, "HoldingPatternPropertyType")
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternTimeSlicePropertyType_stagedOrder, stage.HoldingPatternTimeSlicePropertyTypes_referenceOrder, holdingpatterntimeslicepropertytype, "HoldingPatternTimeSlicePropertyType")
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternTimeSliceType_stagedOrder, stage.HoldingPatternTimeSliceTypes_referenceOrder, holdingpatterntimeslicetype, "HoldingPatternTimeSliceType")
}

func (holdingpatterntype *HoldingPatternType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternType_stagedOrder, stage.HoldingPatternTypes_referenceOrder, holdingpatterntype, "HoldingPatternType")
}

func (holdingusepropertytype *HoldingUsePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingUsePropertyType_stagedOrder, stage.HoldingUsePropertyTypes_referenceOrder, holdingusepropertytype, "HoldingUsePropertyType")
}

func (holdingusetype *HoldingUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingUseType_stagedOrder, stage.HoldingUseTypes_referenceOrder, holdingusetype, "HoldingUseType")
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InitialLegTimeSlicePropertyType_stagedOrder, stage.InitialLegTimeSlicePropertyTypes_referenceOrder, initiallegtimeslicepropertytype, "InitialLegTimeSlicePropertyType")
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InitialLegTimeSliceType_stagedOrder, stage.InitialLegTimeSliceTypes_referenceOrder, initiallegtimeslicetype, "InitialLegTimeSliceType")
}

func (initiallegtype *InitialLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InitialLegType_stagedOrder, stage.InitialLegTypes_referenceOrder, initiallegtype, "InitialLegType")
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentApproachProcedurePropertyType_stagedOrder, stage.InstrumentApproachProcedurePropertyTypes_referenceOrder, instrumentapproachprocedurepropertytype, "InstrumentApproachProcedurePropertyType")
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentApproachProcedureTimeSlicePropertyType_stagedOrder, stage.InstrumentApproachProcedureTimeSlicePropertyTypes_referenceOrder, instrumentapproachproceduretimeslicepropertytype, "InstrumentApproachProcedureTimeSlicePropertyType")
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentApproachProcedureTimeSliceType_stagedOrder, stage.InstrumentApproachProcedureTimeSliceTypes_referenceOrder, instrumentapproachproceduretimeslicetype, "InstrumentApproachProcedureTimeSliceType")
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentApproachProcedureType_stagedOrder, stage.InstrumentApproachProcedureTypes_referenceOrder, instrumentapproachproceduretype, "InstrumentApproachProcedureType")
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.IntermediateLegTimeSlicePropertyType_stagedOrder, stage.IntermediateLegTimeSlicePropertyTypes_referenceOrder, intermediatelegtimeslicepropertytype, "IntermediateLegTimeSlicePropertyType")
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.IntermediateLegTimeSliceType_stagedOrder, stage.IntermediateLegTimeSliceTypes_referenceOrder, intermediatelegtimeslicetype, "IntermediateLegTimeSliceType")
}

func (intermediatelegtype *IntermediateLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.IntermediateLegType_stagedOrder, stage.IntermediateLegTypes_referenceOrder, intermediatelegtype, "IntermediateLegType")
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LandingTakeoffAreaCollectionPropertyType_stagedOrder, stage.LandingTakeoffAreaCollectionPropertyTypes_referenceOrder, landingtakeoffareacollectionpropertytype, "LandingTakeoffAreaCollectionPropertyType")
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LandingTakeoffAreaCollectionType_stagedOrder, stage.LandingTakeoffAreaCollectionTypes_referenceOrder, landingtakeoffareacollectiontype, "LandingTakeoffAreaCollectionType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (minimapropertytype *MinimaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MinimaPropertyType_stagedOrder, stage.MinimaPropertyTypes_referenceOrder, minimapropertytype, "MinimaPropertyType")
}

func (minimatype *MinimaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MinimaType_stagedOrder, stage.MinimaTypes_referenceOrder, minimatype, "MinimaType")
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MissedApproachGroupPropertyType_stagedOrder, stage.MissedApproachGroupPropertyTypes_referenceOrder, missedapproachgrouppropertytype, "MissedApproachGroupPropertyType")
}

func (missedapproachgrouptype *MissedApproachGroupType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MissedApproachGroupType_stagedOrder, stage.MissedApproachGroupTypes_referenceOrder, missedapproachgrouptype, "MissedApproachGroupType")
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MissedApproachLegTimeSlicePropertyType_stagedOrder, stage.MissedApproachLegTimeSlicePropertyTypes_referenceOrder, missedapproachlegtimeslicepropertytype, "MissedApproachLegTimeSlicePropertyType")
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MissedApproachLegTimeSliceType_stagedOrder, stage.MissedApproachLegTimeSliceTypes_referenceOrder, missedapproachlegtimeslicetype, "MissedApproachLegTimeSliceType")
}

func (missedapproachlegtype *MissedApproachLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MissedApproachLegType_stagedOrder, stage.MissedApproachLegTypes_referenceOrder, missedapproachlegtype, "MissedApproachLegType")
}

func (navaidpropertytype *NavaidPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidPropertyType_stagedOrder, stage.NavaidPropertyTypes_referenceOrder, navaidpropertytype, "NavaidPropertyType")
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaRestrictionTimeSlicePropertyType_stagedOrder, stage.NavigationAreaRestrictionTimeSlicePropertyTypes_referenceOrder, navigationarearestrictiontimeslicepropertytype, "NavigationAreaRestrictionTimeSlicePropertyType")
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaRestrictionTimeSliceType_stagedOrder, stage.NavigationAreaRestrictionTimeSliceTypes_referenceOrder, navigationarearestrictiontimeslicetype, "NavigationAreaRestrictionTimeSliceType")
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaRestrictionType_stagedOrder, stage.NavigationAreaRestrictionTypes_referenceOrder, navigationarearestrictiontype, "NavigationAreaRestrictionType")
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaSectorPropertyType_stagedOrder, stage.NavigationAreaSectorPropertyTypes_referenceOrder, navigationareasectorpropertytype, "NavigationAreaSectorPropertyType")
}

func (navigationareasectortype *NavigationAreaSectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaSectorType_stagedOrder, stage.NavigationAreaSectorTypes_referenceOrder, navigationareasectortype, "NavigationAreaSectorType")
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaTimeSlicePropertyType_stagedOrder, stage.NavigationAreaTimeSlicePropertyTypes_referenceOrder, navigationareatimeslicepropertytype, "NavigationAreaTimeSlicePropertyType")
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaTimeSliceType_stagedOrder, stage.NavigationAreaTimeSliceTypes_referenceOrder, navigationareatimeslicetype, "NavigationAreaTimeSliceType")
}

func (navigationareatype *NavigationAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavigationAreaType_stagedOrder, stage.NavigationAreaTypes_referenceOrder, navigationareatype, "NavigationAreaType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAssessmentAreaPropertyType_stagedOrder, stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder, obstacleassessmentareapropertytype, "ObstacleAssessmentAreaPropertyType")
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAssessmentAreaType_stagedOrder, stage.ObstacleAssessmentAreaTypes_referenceOrder, obstacleassessmentareatype, "ObstacleAssessmentAreaType")
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstaclePlacementPropertyType_stagedOrder, stage.ObstaclePlacementPropertyTypes_referenceOrder, obstacleplacementpropertytype, "ObstaclePlacementPropertyType")
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstaclePlacementType_stagedOrder, stage.ObstaclePlacementTypes_referenceOrder, obstacleplacementtype, "ObstaclePlacementType")
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstructionPropertyType_stagedOrder, stage.ObstructionPropertyTypes_referenceOrder, obstructionpropertytype, "ObstructionPropertyType")
}

func (obstructiontype *ObstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstructionType_stagedOrder, stage.ObstructionTypes_referenceOrder, obstructiontype, "ObstructionType")
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyTypes_referenceOrder, organisationauthoritypropertytype, "OrganisationAuthorityPropertyType")
}

func (pointpropertytype *PointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointPropertyType_stagedOrder, stage.PointPropertyTypes_referenceOrder, pointpropertytype, "PointPropertyType")
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointReferencePropertyType_stagedOrder, stage.PointReferencePropertyTypes_referenceOrder, pointreferencepropertytype, "PointReferencePropertyType")
}

func (pointreferencetype *PointReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointReferenceType_stagedOrder, stage.PointReferenceTypes_referenceOrder, pointreferencetype, "PointReferenceType")
}

func (pointtype *PointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointType_stagedOrder, stage.PointTypes_referenceOrder, pointtype, "PointType")
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureAvailabilityPropertyType_stagedOrder, stage.ProcedureAvailabilityPropertyTypes_referenceOrder, procedureavailabilitypropertytype, "ProcedureAvailabilityPropertyType")
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureAvailabilityType_stagedOrder, stage.ProcedureAvailabilityTypes_referenceOrder, procedureavailabilitytype, "ProcedureAvailabilityType")
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureDMETimeSlicePropertyType_stagedOrder, stage.ProcedureDMETimeSlicePropertyTypes_referenceOrder, proceduredmetimeslicepropertytype, "ProcedureDMETimeSlicePropertyType")
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureDMETimeSliceType_stagedOrder, stage.ProcedureDMETimeSliceTypes_referenceOrder, proceduredmetimeslicetype, "ProcedureDMETimeSliceType")
}

func (proceduredmetype *ProcedureDMEType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureDMEType_stagedOrder, stage.ProcedureDMETypes_referenceOrder, proceduredmetype, "ProcedureDMEType")
}

func (procedurepropertytype *ProcedurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedurePropertyType_stagedOrder, stage.ProcedurePropertyTypes_referenceOrder, procedurepropertytype, "ProcedurePropertyType")
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureTransitionLegPropertyType_stagedOrder, stage.ProcedureTransitionLegPropertyTypes_referenceOrder, proceduretransitionlegpropertytype, "ProcedureTransitionLegPropertyType")
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureTransitionLegType_stagedOrder, stage.ProcedureTransitionLegTypes_referenceOrder, proceduretransitionlegtype, "ProcedureTransitionLegType")
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureTransitionPropertyType_stagedOrder, stage.ProcedureTransitionPropertyTypes_referenceOrder, proceduretransitionpropertytype, "ProcedureTransitionPropertyType")
}

func (proceduretransitiontype *ProcedureTransitionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedureTransitionType_stagedOrder, stage.ProcedureTransitionTypes_referenceOrder, proceduretransitiontype, "ProcedureTransitionType")
}

func (radarsystempropertytype *RadarSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarSystemPropertyType_stagedOrder, stage.RadarSystemPropertyTypes_referenceOrder, radarsystempropertytype, "RadarSystemPropertyType")
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioFrequencyAreaPropertyType_stagedOrder, stage.RadioFrequencyAreaPropertyTypes_referenceOrder, radiofrequencyareapropertytype, "RadioFrequencyAreaPropertyType")
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayCentrelinePointPropertyType_stagedOrder, stage.RunwayCentrelinePointPropertyTypes_referenceOrder, runwaycentrelinepointpropertytype, "RunwayCentrelinePointPropertyType")
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayDirectionPropertyType_stagedOrder, stage.RunwayDirectionPropertyTypes_referenceOrder, runwaydirectionpropertytype, "RunwayDirectionPropertyType")
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaPropertyType_stagedOrder, stage.SafeAltitudeAreaPropertyTypes_referenceOrder, safealtitudeareapropertytype, "SafeAltitudeAreaPropertyType")
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaSectorPropertyType_stagedOrder, stage.SafeAltitudeAreaSectorPropertyTypes_referenceOrder, safealtitudeareasectorpropertytype, "SafeAltitudeAreaSectorPropertyType")
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaSectorType_stagedOrder, stage.SafeAltitudeAreaSectorTypes_referenceOrder, safealtitudeareasectortype, "SafeAltitudeAreaSectorType")
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaTimeSlicePropertyType_stagedOrder, stage.SafeAltitudeAreaTimeSlicePropertyTypes_referenceOrder, safealtitudeareatimeslicepropertytype, "SafeAltitudeAreaTimeSlicePropertyType")
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaTimeSliceType_stagedOrder, stage.SafeAltitudeAreaTimeSliceTypes_referenceOrder, safealtitudeareatimeslicetype, "SafeAltitudeAreaTimeSliceType")
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SafeAltitudeAreaType_stagedOrder, stage.SafeAltitudeAreaTypes_referenceOrder, safealtitudeareatype, "SafeAltitudeAreaType")
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SectorDesignPropertyType_stagedOrder, stage.SectorDesignPropertyTypes_referenceOrder, sectordesignpropertytype, "SectorDesignPropertyType")
}

func (sectordesigntype *SectorDesignType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SectorDesignType_stagedOrder, stage.SectorDesignTypes_referenceOrder, sectordesigntype, "SectorDesignType")
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SegmentLegPropertyType_stagedOrder, stage.SegmentLegPropertyTypes_referenceOrder, segmentlegpropertytype, "SegmentLegPropertyType")
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SegmentPointPropertyType_stagedOrder, stage.SegmentPointPropertyTypes_referenceOrder, segmentpointpropertytype, "SegmentPointPropertyType")
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationSystemPropertyType_stagedOrder, stage.SpecialNavigationSystemPropertyTypes_referenceOrder, specialnavigationsystempropertytype, "SpecialNavigationSystemPropertyType")
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentArrivalPropertyType_stagedOrder, stage.StandardInstrumentArrivalPropertyTypes_referenceOrder, standardinstrumentarrivalpropertytype, "StandardInstrumentArrivalPropertyType")
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentArrivalTimeSlicePropertyType_stagedOrder, stage.StandardInstrumentArrivalTimeSlicePropertyTypes_referenceOrder, standardinstrumentarrivaltimeslicepropertytype, "StandardInstrumentArrivalTimeSlicePropertyType")
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentArrivalTimeSliceType_stagedOrder, stage.StandardInstrumentArrivalTimeSliceTypes_referenceOrder, standardinstrumentarrivaltimeslicetype, "StandardInstrumentArrivalTimeSliceType")
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentArrivalType_stagedOrder, stage.StandardInstrumentArrivalTypes_referenceOrder, standardinstrumentarrivaltype, "StandardInstrumentArrivalType")
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentDeparturePropertyType_stagedOrder, stage.StandardInstrumentDeparturePropertyTypes_referenceOrder, standardinstrumentdeparturepropertytype, "StandardInstrumentDeparturePropertyType")
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentDepartureTimeSlicePropertyType_stagedOrder, stage.StandardInstrumentDepartureTimeSlicePropertyTypes_referenceOrder, standardinstrumentdeparturetimeslicepropertytype, "StandardInstrumentDepartureTimeSlicePropertyType")
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentDepartureTimeSliceType_stagedOrder, stage.StandardInstrumentDepartureTimeSliceTypes_referenceOrder, standardinstrumentdeparturetimeslicetype, "StandardInstrumentDepartureTimeSliceType")
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentDepartureType_stagedOrder, stage.StandardInstrumentDepartureTypes_referenceOrder, standardinstrumentdeparturetype, "StandardInstrumentDepartureType")
}

func (surfacepropertytype *SurfacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfacePropertyType_stagedOrder, stage.SurfacePropertyTypes_referenceOrder, surfacepropertytype, "SurfacePropertyType")
}

func (surfacetype *SurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfaceType_stagedOrder, stage.SurfaceTypes_referenceOrder, surfacetype, "SurfaceType")
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalArrivalAreaSectorPropertyType_stagedOrder, stage.TerminalArrivalAreaSectorPropertyTypes_referenceOrder, terminalarrivalareasectorpropertytype, "TerminalArrivalAreaSectorPropertyType")
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalArrivalAreaSectorType_stagedOrder, stage.TerminalArrivalAreaSectorTypes_referenceOrder, terminalarrivalareasectortype, "TerminalArrivalAreaSectorType")
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalArrivalAreaTimeSlicePropertyType_stagedOrder, stage.TerminalArrivalAreaTimeSlicePropertyTypes_referenceOrder, terminalarrivalareatimeslicepropertytype, "TerminalArrivalAreaTimeSlicePropertyType")
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalArrivalAreaTimeSliceType_stagedOrder, stage.TerminalArrivalAreaTimeSliceTypes_referenceOrder, terminalarrivalareatimeslicetype, "TerminalArrivalAreaTimeSliceType")
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalArrivalAreaType_stagedOrder, stage.TerminalArrivalAreaTypes_referenceOrder, terminalarrivalareatype, "TerminalArrivalAreaType")
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalSegmentPointPropertyType_stagedOrder, stage.TerminalSegmentPointPropertyTypes_referenceOrder, terminalsegmentpointpropertytype, "TerminalSegmentPointPropertyType")
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TerminalSegmentPointType_stagedOrder, stage.TerminalSegmentPointTypes_referenceOrder, terminalsegmentpointtype, "TerminalSegmentPointType")
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyTypes_referenceOrder, timesheetpropertytype, "TimesheetPropertyType")
}

func (timesheettype *TimesheetType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetType_stagedOrder, stage.TimesheetTypes_referenceOrder, timesheettype, "TimesheetType")
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TouchDownLiftOffPropertyType_stagedOrder, stage.TouchDownLiftOffPropertyTypes_referenceOrder, touchdownliftoffpropertytype, "TouchDownLiftOffPropertyType")
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnplannedHoldingPropertyType_stagedOrder, stage.UnplannedHoldingPropertyTypes_referenceOrder, unplannedholdingpropertytype, "UnplannedHoldingPropertyType")
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnplannedHoldingTimeSlicePropertyType_stagedOrder, stage.UnplannedHoldingTimeSlicePropertyTypes_referenceOrder, unplannedholdingtimeslicepropertytype, "UnplannedHoldingTimeSlicePropertyType")
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnplannedHoldingTimeSliceType_stagedOrder, stage.UnplannedHoldingTimeSliceTypes_referenceOrder, unplannedholdingtimeslicetype, "UnplannedHoldingTimeSliceType")
}

func (unplannedholdingtype *UnplannedHoldingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnplannedHoldingType_stagedOrder, stage.UnplannedHoldingTypes_referenceOrder, unplannedholdingtype, "UnplannedHoldingType")
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructurePropertyType_stagedOrder, stage.VerticalStructurePropertyTypes_referenceOrder, verticalstructurepropertytype, "VerticalStructurePropertyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a_extension *A_extension) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension, a_extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension *A_extension) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension.GongGetIdentifier(stage)
}

func (a_extension_1 *A_extension_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_1, a_extension_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_1 *A_extension_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_1.GongGetIdentifier(stage)
}

func (a_extension_10 *A_extension_10) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_10, a_extension_10.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_10 *A_extension_10) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_10.GongGetIdentifier(stage)
}

func (a_extension_11 *A_extension_11) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_11, a_extension_11.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_11 *A_extension_11) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_11.GongGetIdentifier(stage)
}

func (a_extension_12 *A_extension_12) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_12, a_extension_12.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_12 *A_extension_12) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_12.GongGetIdentifier(stage)
}

func (a_extension_13 *A_extension_13) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_13, a_extension_13.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_13 *A_extension_13) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_13.GongGetIdentifier(stage)
}

func (a_extension_14 *A_extension_14) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_14, a_extension_14.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_14 *A_extension_14) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_14.GongGetIdentifier(stage)
}

func (a_extension_15 *A_extension_15) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_15, a_extension_15.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_15 *A_extension_15) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_15.GongGetIdentifier(stage)
}

func (a_extension_16 *A_extension_16) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_16, a_extension_16.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_16 *A_extension_16) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_16.GongGetIdentifier(stage)
}

func (a_extension_17 *A_extension_17) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_17, a_extension_17.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_17 *A_extension_17) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_17.GongGetIdentifier(stage)
}

func (a_extension_18 *A_extension_18) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_18, a_extension_18.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_18 *A_extension_18) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_18.GongGetIdentifier(stage)
}

func (a_extension_19 *A_extension_19) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_19, a_extension_19.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_19 *A_extension_19) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_19.GongGetIdentifier(stage)
}

func (a_extension_2 *A_extension_2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_2, a_extension_2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_2 *A_extension_2) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_2.GongGetIdentifier(stage)
}

func (a_extension_20 *A_extension_20) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_20, a_extension_20.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_20 *A_extension_20) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_20.GongGetIdentifier(stage)
}

func (a_extension_21 *A_extension_21) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_21, a_extension_21.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_21 *A_extension_21) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_21.GongGetIdentifier(stage)
}

func (a_extension_22 *A_extension_22) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_22, a_extension_22.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_22 *A_extension_22) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_22.GongGetIdentifier(stage)
}

func (a_extension_23 *A_extension_23) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_23, a_extension_23.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_23 *A_extension_23) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_23.GongGetIdentifier(stage)
}

func (a_extension_24 *A_extension_24) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_24, a_extension_24.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_24 *A_extension_24) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_24.GongGetIdentifier(stage)
}

func (a_extension_25 *A_extension_25) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_25, a_extension_25.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_25 *A_extension_25) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_25.GongGetIdentifier(stage)
}

func (a_extension_26 *A_extension_26) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_26, a_extension_26.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_26 *A_extension_26) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_26.GongGetIdentifier(stage)
}

func (a_extension_27 *A_extension_27) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_27, a_extension_27.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_27 *A_extension_27) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_27.GongGetIdentifier(stage)
}

func (a_extension_28 *A_extension_28) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_28, a_extension_28.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_28 *A_extension_28) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_28.GongGetIdentifier(stage)
}

func (a_extension_29 *A_extension_29) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_29, a_extension_29.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_29 *A_extension_29) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_29.GongGetIdentifier(stage)
}

func (a_extension_3 *A_extension_3) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_3, a_extension_3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_3 *A_extension_3) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_3.GongGetIdentifier(stage)
}

func (a_extension_30 *A_extension_30) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_30, a_extension_30.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_30 *A_extension_30) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_30.GongGetIdentifier(stage)
}

func (a_extension_31 *A_extension_31) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_31, a_extension_31.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_31 *A_extension_31) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_31.GongGetIdentifier(stage)
}

func (a_extension_32 *A_extension_32) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_32, a_extension_32.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_32 *A_extension_32) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_32.GongGetIdentifier(stage)
}

func (a_extension_33 *A_extension_33) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_33, a_extension_33.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_33 *A_extension_33) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_33.GongGetIdentifier(stage)
}

func (a_extension_34 *A_extension_34) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_34, a_extension_34.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_34 *A_extension_34) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_34.GongGetIdentifier(stage)
}

func (a_extension_35 *A_extension_35) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_35, a_extension_35.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_35 *A_extension_35) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_35.GongGetIdentifier(stage)
}

func (a_extension_36 *A_extension_36) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_36, a_extension_36.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_36 *A_extension_36) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_36.GongGetIdentifier(stage)
}

func (a_extension_37 *A_extension_37) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_37, a_extension_37.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_37 *A_extension_37) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_37.GongGetIdentifier(stage)
}

func (a_extension_38 *A_extension_38) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_38, a_extension_38.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_38 *A_extension_38) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_38.GongGetIdentifier(stage)
}

func (a_extension_39 *A_extension_39) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_39, a_extension_39.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_39 *A_extension_39) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_39.GongGetIdentifier(stage)
}

func (a_extension_4 *A_extension_4) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_4, a_extension_4.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_4 *A_extension_4) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_4.GongGetIdentifier(stage)
}

func (a_extension_40 *A_extension_40) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_40, a_extension_40.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_40 *A_extension_40) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_40.GongGetIdentifier(stage)
}

func (a_extension_41 *A_extension_41) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_41, a_extension_41.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_41 *A_extension_41) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_41.GongGetIdentifier(stage)
}

func (a_extension_42 *A_extension_42) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_42, a_extension_42.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_42 *A_extension_42) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_42.GongGetIdentifier(stage)
}

func (a_extension_43 *A_extension_43) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_43, a_extension_43.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_43 *A_extension_43) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_43.GongGetIdentifier(stage)
}

func (a_extension_44 *A_extension_44) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_44, a_extension_44.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_44 *A_extension_44) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_44.GongGetIdentifier(stage)
}

func (a_extension_45 *A_extension_45) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_45, a_extension_45.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_45 *A_extension_45) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_45.GongGetIdentifier(stage)
}

func (a_extension_46 *A_extension_46) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_46, a_extension_46.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_46 *A_extension_46) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_46.GongGetIdentifier(stage)
}

func (a_extension_47 *A_extension_47) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_47, a_extension_47.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_47 *A_extension_47) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_47.GongGetIdentifier(stage)
}

func (a_extension_48 *A_extension_48) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_48, a_extension_48.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_48 *A_extension_48) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_48.GongGetIdentifier(stage)
}

func (a_extension_49 *A_extension_49) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_49, a_extension_49.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_49 *A_extension_49) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_49.GongGetIdentifier(stage)
}

func (a_extension_5 *A_extension_5) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_5, a_extension_5.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_5 *A_extension_5) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_5.GongGetIdentifier(stage)
}

func (a_extension_50 *A_extension_50) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_50, a_extension_50.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_50 *A_extension_50) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_50.GongGetIdentifier(stage)
}

func (a_extension_51 *A_extension_51) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_51, a_extension_51.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_51 *A_extension_51) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_51.GongGetIdentifier(stage)
}

func (a_extension_52 *A_extension_52) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_52, a_extension_52.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_52 *A_extension_52) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_52.GongGetIdentifier(stage)
}

func (a_extension_53 *A_extension_53) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_53, a_extension_53.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_53 *A_extension_53) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_53.GongGetIdentifier(stage)
}

func (a_extension_54 *A_extension_54) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_54, a_extension_54.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_54 *A_extension_54) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_54.GongGetIdentifier(stage)
}

func (a_extension_55 *A_extension_55) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_55, a_extension_55.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_55 *A_extension_55) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_55.GongGetIdentifier(stage)
}

func (a_extension_56 *A_extension_56) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_56, a_extension_56.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_56 *A_extension_56) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_56.GongGetIdentifier(stage)
}

func (a_extension_57 *A_extension_57) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_57, a_extension_57.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_57 *A_extension_57) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_57.GongGetIdentifier(stage)
}

func (a_extension_58 *A_extension_58) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_58, a_extension_58.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_58 *A_extension_58) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_58.GongGetIdentifier(stage)
}

func (a_extension_6 *A_extension_6) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_6, a_extension_6.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_6 *A_extension_6) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_6.GongGetIdentifier(stage)
}

func (a_extension_7 *A_extension_7) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_7, a_extension_7.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_7 *A_extension_7) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_7.GongGetIdentifier(stage)
}

func (a_extension_8 *A_extension_8) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_8, a_extension_8.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_8 *A_extension_8) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_8.GongGetIdentifier(stage)
}

func (a_extension_9 *A_extension_9) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_9, a_extension_9.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_9 *A_extension_9) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_9.GongGetIdentifier(stage)
}

func (abstractapproachlegtype *AbstractApproachLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractapproachlegtype, abstractapproachlegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractapproachlegtype *AbstractApproachLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractapproachlegtype.GongGetIdentifier(stage)
}

func (abstractproceduretype *AbstractProcedureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractproceduretype, abstractproceduretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractproceduretype *AbstractProcedureType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractproceduretype.GongGetIdentifier(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractsegmentlegtype, abstractsegmentlegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractsegmentlegtype *AbstractSegmentLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractsegmentlegtype.GongGetIdentifier(stage)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractsegmentpointtype, abstractsegmentpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractsegmentpointtype.GongGetIdentifier(stage)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingpointtype, aerialrefuellingpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingpointtype.GongGetIdentifier(stage)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftcharacteristicpropertytype, aircraftcharacteristicpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftcharacteristictype, aircraftcharacteristictype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftcharacteristictype.GongGetIdentifier(stage)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportheliportpropertytype, airportheliportpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportheliportpropertytype.GongGetIdentifier(stage)
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(altimetersourcepropertytype, altimetersourcepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return altimetersourcepropertytype.GongGetIdentifier(stage)
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(altitudeadjustmentpropertytype, altitudeadjustmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return altitudeadjustmentpropertytype.GongGetIdentifier(stage)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(altitudeadjustmenttype, altitudeadjustmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return altitudeadjustmenttype.GongGetIdentifier(stage)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleindicationpropertytype, angleindicationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleindicationpropertytype.GongGetIdentifier(stage)
}

func (angleusepropertytype *AngleUsePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleusepropertytype, angleusepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleusepropertytype *AngleUsePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleusepropertytype.GongGetIdentifier(stage)
}

func (angleusetype *AngleUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleusetype, angleusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleusetype *AngleUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleusetype.GongGetIdentifier(stage)
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachaltitudetablepropertytype, approachaltitudetablepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachaltitudetablepropertytype.GongGetIdentifier(stage)
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachaltitudetabletype, approachaltitudetabletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachaltitudetabletype *ApproachAltitudeTableType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachaltitudetabletype.GongGetIdentifier(stage)
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachconditionpropertytype, approachconditionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachconditionpropertytype *ApproachConditionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachconditionpropertytype.GongGetIdentifier(stage)
}

func (approachconditiontype *ApproachConditionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachconditiontype, approachconditiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachconditiontype *ApproachConditionType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachconditiontype.GongGetIdentifier(stage)
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachdistancetablepropertytype, approachdistancetablepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachdistancetablepropertytype.GongGetIdentifier(stage)
}

func (approachdistancetabletype *ApproachDistanceTableType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachdistancetabletype, approachdistancetabletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachdistancetabletype *ApproachDistanceTableType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachdistancetabletype.GongGetIdentifier(stage)
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachtimingtablepropertytype, approachtimingtablepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachtimingtablepropertytype.GongGetIdentifier(stage)
}

func (approachtimingtabletype *ApproachTimingTableType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(approachtimingtabletype, approachtimingtabletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (approachtimingtabletype *ApproachTimingTableType) GongGetReferenceIdentifier(stage *Stage) string {
	return approachtimingtabletype.GongGetIdentifier(stage)
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivalfeederlegtimeslicepropertytype, arrivalfeederlegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivalfeederlegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivalfeederlegtimeslicetype, arrivalfeederlegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivalfeederlegtimeslicetype.GongGetIdentifier(stage)
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivalfeederlegtype, arrivalfeederlegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivalfeederlegtype *ArrivalFeederLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivalfeederlegtype.GongGetIdentifier(stage)
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivallegtimeslicepropertytype, arrivallegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivallegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivallegtimeslicetype, arrivallegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivallegtimeslicetype.GongGetIdentifier(stage)
}

func (arrivallegtype *ArrivalLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrivallegtype, arrivallegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrivallegtype *ArrivalLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return arrivallegtype.GongGetIdentifier(stage)
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlesectorpropertytype, circlesectorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlesectorpropertytype *CircleSectorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlesectorpropertytype.GongGetIdentifier(stage)
}

func (circlesectortype *CircleSectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlesectortype, circlesectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlesectortype *CircleSectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlesectortype.GongGetIdentifier(stage)
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlingareatimeslicepropertytype, circlingareatimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlingareatimeslicepropertytype.GongGetIdentifier(stage)
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlingareatimeslicetype, circlingareatimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlingareatimeslicetype.GongGetIdentifier(stage)
}

func (circlingareatype *CirclingAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlingareatype, circlingareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlingareatype *CirclingAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlingareatype.GongGetIdentifier(stage)
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlingrestrictionpropertytype, circlingrestrictionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlingrestrictionpropertytype.GongGetIdentifier(stage)
}

func (circlingrestrictiontype *CirclingRestrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlingrestrictiontype, circlingrestrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlingrestrictiontype *CirclingRestrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return circlingrestrictiontype.GongGetIdentifier(stage)
}

func (curvepropertytype *CurvePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curvepropertytype, curvepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curvepropertytype *CurvePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return curvepropertytype.GongGetIdentifier(stage)
}

func (curvetype *CurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curvetype, curvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curvetype *CurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return curvetype.GongGetIdentifier(stage)
}

func (dmepropertytype *DMEPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dmepropertytype, dmepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dmepropertytype *DMEPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return dmepropertytype.GongGetIdentifier(stage)
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(departurearrivalconditionpropertytype, departurearrivalconditionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return departurearrivalconditionpropertytype.GongGetIdentifier(stage)
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(departurearrivalconditiontype, departurearrivalconditiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (departurearrivalconditiontype *DepartureArrivalConditionType) GongGetReferenceIdentifier(stage *Stage) string {
	return departurearrivalconditiontype.GongGetIdentifier(stage)
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(departurelegtimeslicepropertytype, departurelegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return departurelegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(departurelegtimeslicetype, departurelegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return departurelegtimeslicetype.GongGetIdentifier(stage)
}

func (departurelegtype *DepartureLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(departurelegtype, departurelegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (departurelegtype *DepartureLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return departurelegtype.GongGetIdentifier(stage)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointpropertytype, designatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointpropertytype.GongGetIdentifier(stage)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distanceindicationpropertytype, distanceindicationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return distanceindicationpropertytype.GongGetIdentifier(stage)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedcurvetype, elevatedcurvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedcurvetype *ElevatedCurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedcurvetype.GongGetIdentifier(stage)
}

func (elevatedpointtype *ElevatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointtype, elevatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointtype *ElevatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointtype.GongGetIdentifier(stage)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedsurfacetype, elevatedsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedsurfacetype *ElevatedSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedsurfacetype.GongGetIdentifier(stage)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(enroutesegmentpointtype, enroutesegmentpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return enroutesegmentpointtype.GongGetIdentifier(stage)
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(equipmentunavailableadjustmentcolumnpropertytype, equipmentunavailableadjustmentcolumnpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return equipmentunavailableadjustmentcolumnpropertytype.GongGetIdentifier(stage)
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(equipmentunavailableadjustmentcolumntype, equipmentunavailableadjustmentcolumntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongGetReferenceIdentifier(stage *Stage) string {
	return equipmentunavailableadjustmentcolumntype.GongGetIdentifier(stage)
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(equipmentunavailableadjustmentpropertytype, equipmentunavailableadjustmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return equipmentunavailableadjustmentpropertytype.GongGetIdentifier(stage)
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(equipmentunavailableadjustmenttype, equipmentunavailableadjustmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return equipmentunavailableadjustmenttype.GongGetIdentifier(stage)
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fasdatablockpropertytype, fasdatablockpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fasdatablockpropertytype *FASDataBlockPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return fasdatablockpropertytype.GongGetIdentifier(stage)
}

func (fasdatablocktype *FASDataBlockType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fasdatablocktype, fasdatablocktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fasdatablocktype *FASDataBlockType) GongGetReferenceIdentifier(stage *Stage) string {
	return fasdatablocktype.GongGetIdentifier(stage)
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(finallegtimeslicepropertytype, finallegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return finallegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(finallegtimeslicetype, finallegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (finallegtimeslicetype *FinalLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return finallegtimeslicetype.GongGetIdentifier(stage)
}

func (finallegtype *FinalLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(finallegtype, finallegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (finallegtype *FinalLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return finallegtype.GongGetIdentifier(stage)
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(finalprofilepropertytype, finalprofilepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (finalprofilepropertytype *FinalProfilePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return finalprofilepropertytype.GongGetIdentifier(stage)
}

func (finalprofiletype *FinalProfileType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(finalprofiletype, finalprofiletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (finalprofiletype *FinalProfileType) GongGetReferenceIdentifier(stage *Stage) string {
	return finalprofiletype.GongGetIdentifier(stage)
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingassessmenttimeslicepropertytype, holdingassessmenttimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingassessmenttimeslicepropertytype.GongGetIdentifier(stage)
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingassessmenttimeslicetype, holdingassessmenttimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingassessmenttimeslicetype.GongGetIdentifier(stage)
}

func (holdingassessmenttype *HoldingAssessmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingassessmenttype, holdingassessmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingassessmenttype *HoldingAssessmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingassessmenttype.GongGetIdentifier(stage)
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterndistancepropertytype, holdingpatterndistancepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterndistancepropertytype.GongGetIdentifier(stage)
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterndistancetype, holdingpatterndistancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterndistancetype *HoldingPatternDistanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterndistancetype.GongGetIdentifier(stage)
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterndurationpropertytype, holdingpatterndurationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterndurationpropertytype.GongGetIdentifier(stage)
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterndurationtype, holdingpatterndurationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterndurationtype *HoldingPatternDurationType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterndurationtype.GongGetIdentifier(stage)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatternpropertytype, holdingpatternpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatternpropertytype.GongGetIdentifier(stage)
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterntimeslicepropertytype, holdingpatterntimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterntimeslicepropertytype.GongGetIdentifier(stage)
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterntimeslicetype, holdingpatterntimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterntimeslicetype.GongGetIdentifier(stage)
}

func (holdingpatterntype *HoldingPatternType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatterntype, holdingpatterntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatterntype *HoldingPatternType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatterntype.GongGetIdentifier(stage)
}

func (holdingusepropertytype *HoldingUsePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingusepropertytype, holdingusepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingusepropertytype *HoldingUsePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingusepropertytype.GongGetIdentifier(stage)
}

func (holdingusetype *HoldingUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingusetype, holdingusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingusetype *HoldingUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingusetype.GongGetIdentifier(stage)
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(initiallegtimeslicepropertytype, initiallegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return initiallegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(initiallegtimeslicetype, initiallegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initiallegtimeslicetype *InitialLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return initiallegtimeslicetype.GongGetIdentifier(stage)
}

func (initiallegtype *InitialLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(initiallegtype, initiallegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initiallegtype *InitialLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return initiallegtype.GongGetIdentifier(stage)
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentapproachprocedurepropertytype, instrumentapproachprocedurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentapproachprocedurepropertytype.GongGetIdentifier(stage)
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentapproachproceduretimeslicepropertytype, instrumentapproachproceduretimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentapproachproceduretimeslicepropertytype.GongGetIdentifier(stage)
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentapproachproceduretimeslicetype, instrumentapproachproceduretimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentapproachproceduretimeslicetype.GongGetIdentifier(stage)
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentapproachproceduretype, instrumentapproachproceduretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentapproachproceduretype.GongGetIdentifier(stage)
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(intermediatelegtimeslicepropertytype, intermediatelegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return intermediatelegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(intermediatelegtimeslicetype, intermediatelegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return intermediatelegtimeslicetype.GongGetIdentifier(stage)
}

func (intermediatelegtype *IntermediateLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(intermediatelegtype, intermediatelegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (intermediatelegtype *IntermediateLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return intermediatelegtype.GongGetIdentifier(stage)
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(landingtakeoffareacollectionpropertytype, landingtakeoffareacollectionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return landingtakeoffareacollectionpropertytype.GongGetIdentifier(stage)
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(landingtakeoffareacollectiontype, landingtakeoffareacollectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return landingtakeoffareacollectiontype.GongGetIdentifier(stage)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linguisticnotepropertytype, linguisticnotepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return linguisticnotepropertytype.GongGetIdentifier(stage)
}

func (linguisticnotetype *LinguisticNoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linguisticnotetype, linguisticnotetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linguisticnotetype *LinguisticNoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return linguisticnotetype.GongGetIdentifier(stage)
}

func (minimapropertytype *MinimaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(minimapropertytype, minimapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (minimapropertytype *MinimaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return minimapropertytype.GongGetIdentifier(stage)
}

func (minimatype *MinimaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(minimatype, minimatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (minimatype *MinimaType) GongGetReferenceIdentifier(stage *Stage) string {
	return minimatype.GongGetIdentifier(stage)
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(missedapproachgrouppropertytype, missedapproachgrouppropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return missedapproachgrouppropertytype.GongGetIdentifier(stage)
}

func (missedapproachgrouptype *MissedApproachGroupType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(missedapproachgrouptype, missedapproachgrouptype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (missedapproachgrouptype *MissedApproachGroupType) GongGetReferenceIdentifier(stage *Stage) string {
	return missedapproachgrouptype.GongGetIdentifier(stage)
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(missedapproachlegtimeslicepropertytype, missedapproachlegtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return missedapproachlegtimeslicepropertytype.GongGetIdentifier(stage)
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(missedapproachlegtimeslicetype, missedapproachlegtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return missedapproachlegtimeslicetype.GongGetIdentifier(stage)
}

func (missedapproachlegtype *MissedApproachLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(missedapproachlegtype, missedapproachlegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (missedapproachlegtype *MissedApproachLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return missedapproachlegtype.GongGetIdentifier(stage)
}

func (navaidpropertytype *NavaidPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidpropertytype, navaidpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidpropertytype *NavaidPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidpropertytype.GongGetIdentifier(stage)
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationarearestrictiontimeslicepropertytype, navigationarearestrictiontimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationarearestrictiontimeslicepropertytype.GongGetIdentifier(stage)
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationarearestrictiontimeslicetype, navigationarearestrictiontimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationarearestrictiontimeslicetype.GongGetIdentifier(stage)
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationarearestrictiontype, navigationarearestrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationarearestrictiontype.GongGetIdentifier(stage)
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationareasectorpropertytype, navigationareasectorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationareasectorpropertytype.GongGetIdentifier(stage)
}

func (navigationareasectortype *NavigationAreaSectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationareasectortype, navigationareasectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationareasectortype *NavigationAreaSectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationareasectortype.GongGetIdentifier(stage)
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationareatimeslicepropertytype, navigationareatimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationareatimeslicepropertytype.GongGetIdentifier(stage)
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationareatimeslicetype, navigationareatimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationareatimeslicetype.GongGetIdentifier(stage)
}

func (navigationareatype *NavigationAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navigationareatype, navigationareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navigationareatype *NavigationAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return navigationareatype.GongGetIdentifier(stage)
}

func (notepropertytype *NotePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notepropertytype, notepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notepropertytype *NotePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return notepropertytype.GongGetIdentifier(stage)
}

func (notetype *NoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetype, notetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetype *NoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return notetype.GongGetIdentifier(stage)
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleassessmentareapropertytype, obstacleassessmentareapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleassessmentareapropertytype.GongGetIdentifier(stage)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleassessmentareatype, obstacleassessmentareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleassessmentareatype.GongGetIdentifier(stage)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleplacementpropertytype, obstacleplacementpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleplacementpropertytype.GongGetIdentifier(stage)
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleplacementtype, obstacleplacementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleplacementtype *ObstaclePlacementType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleplacementtype.GongGetIdentifier(stage)
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstructionpropertytype, obstructionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstructionpropertytype *ObstructionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstructionpropertytype.GongGetIdentifier(stage)
}

func (obstructiontype *ObstructionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstructiontype, obstructiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstructiontype *ObstructionType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstructiontype.GongGetIdentifier(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritypropertytype, organisationauthoritypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritypropertytype.GongGetIdentifier(stage)
}

func (pointpropertytype *PointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointpropertytype, pointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointpropertytype *PointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointpropertytype.GongGetIdentifier(stage)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointreferencepropertytype, pointreferencepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointreferencepropertytype *PointReferencePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointreferencepropertytype.GongGetIdentifier(stage)
}

func (pointreferencetype *PointReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointreferencetype, pointreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointreferencetype *PointReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointreferencetype.GongGetIdentifier(stage)
}

func (pointtype *PointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointtype, pointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointtype *PointType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointtype.GongGetIdentifier(stage)
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(procedureavailabilitypropertytype, procedureavailabilitypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return procedureavailabilitypropertytype.GongGetIdentifier(stage)
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(procedureavailabilitytype, procedureavailabilitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (procedureavailabilitytype *ProcedureAvailabilityType) GongGetReferenceIdentifier(stage *Stage) string {
	return procedureavailabilitytype.GongGetIdentifier(stage)
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduredmetimeslicepropertytype, proceduredmetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduredmetimeslicepropertytype.GongGetIdentifier(stage)
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduredmetimeslicetype, proceduredmetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduredmetimeslicetype.GongGetIdentifier(stage)
}

func (proceduredmetype *ProcedureDMEType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduredmetype, proceduredmetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduredmetype *ProcedureDMEType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduredmetype.GongGetIdentifier(stage)
}

func (procedurepropertytype *ProcedurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(procedurepropertytype, procedurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (procedurepropertytype *ProcedurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return procedurepropertytype.GongGetIdentifier(stage)
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduretransitionlegpropertytype, proceduretransitionlegpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduretransitionlegpropertytype.GongGetIdentifier(stage)
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduretransitionlegtype, proceduretransitionlegtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduretransitionlegtype *ProcedureTransitionLegType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduretransitionlegtype.GongGetIdentifier(stage)
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduretransitionpropertytype, proceduretransitionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduretransitionpropertytype.GongGetIdentifier(stage)
}

func (proceduretransitiontype *ProcedureTransitionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(proceduretransitiontype, proceduretransitiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (proceduretransitiontype *ProcedureTransitionType) GongGetReferenceIdentifier(stage *Stage) string {
	return proceduretransitiontype.GongGetIdentifier(stage)
}

func (radarsystempropertytype *RadarSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarsystempropertytype, radarsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarsystempropertytype *RadarSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarsystempropertytype.GongGetIdentifier(stage)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiofrequencyareapropertytype, radiofrequencyareapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiofrequencyareapropertytype.GongGetIdentifier(stage)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaycentrelinepointpropertytype, runwaycentrelinepointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaydirectionpropertytype, runwaydirectionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaydirectionpropertytype.GongGetIdentifier(stage)
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareapropertytype, safealtitudeareapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareapropertytype.GongGetIdentifier(stage)
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareasectorpropertytype, safealtitudeareasectorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareasectorpropertytype.GongGetIdentifier(stage)
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareasectortype, safealtitudeareasectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareasectortype.GongGetIdentifier(stage)
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareatimeslicepropertytype, safealtitudeareatimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareatimeslicepropertytype.GongGetIdentifier(stage)
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareatimeslicetype, safealtitudeareatimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareatimeslicetype.GongGetIdentifier(stage)
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(safealtitudeareatype, safealtitudeareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (safealtitudeareatype *SafeAltitudeAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return safealtitudeareatype.GongGetIdentifier(stage)
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sectordesignpropertytype, sectordesignpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sectordesignpropertytype *SectorDesignPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return sectordesignpropertytype.GongGetIdentifier(stage)
}

func (sectordesigntype *SectorDesignType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sectordesigntype, sectordesigntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sectordesigntype *SectorDesignType) GongGetReferenceIdentifier(stage *Stage) string {
	return sectordesigntype.GongGetIdentifier(stage)
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(segmentlegpropertytype, segmentlegpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (segmentlegpropertytype *SegmentLegPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return segmentlegpropertytype.GongGetIdentifier(stage)
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(segmentpointpropertytype, segmentpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (segmentpointpropertytype *SegmentPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return segmentpointpropertytype.GongGetIdentifier(stage)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationsystempropertytype, specialnavigationsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationsystempropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentarrivalpropertytype, standardinstrumentarrivalpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentarrivalpropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentarrivaltimeslicepropertytype, standardinstrumentarrivaltimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentarrivaltimeslicepropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentarrivaltimeslicetype, standardinstrumentarrivaltimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentarrivaltimeslicetype.GongGetIdentifier(stage)
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentarrivaltype, standardinstrumentarrivaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentarrivaltype.GongGetIdentifier(stage)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentdeparturepropertytype, standardinstrumentdeparturepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentdeparturepropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentdeparturetimeslicepropertytype, standardinstrumentdeparturetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentdeparturetimeslicepropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentdeparturetimeslicetype, standardinstrumentdeparturetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentdeparturetimeslicetype.GongGetIdentifier(stage)
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentdeparturetype, standardinstrumentdeparturetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentdeparturetype.GongGetIdentifier(stage)
}

func (surfacepropertytype *SurfacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surfacepropertytype, surfacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surfacepropertytype *SurfacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return surfacepropertytype.GongGetIdentifier(stage)
}

func (surfacetype *SurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surfacetype, surfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surfacetype *SurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return surfacetype.GongGetIdentifier(stage)
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalarrivalareasectorpropertytype, terminalarrivalareasectorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalarrivalareasectorpropertytype.GongGetIdentifier(stage)
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalarrivalareasectortype, terminalarrivalareasectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalarrivalareasectortype.GongGetIdentifier(stage)
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalarrivalareatimeslicepropertytype, terminalarrivalareatimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalarrivalareatimeslicepropertytype.GongGetIdentifier(stage)
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalarrivalareatimeslicetype, terminalarrivalareatimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalarrivalareatimeslicetype.GongGetIdentifier(stage)
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalarrivalareatype, terminalarrivalareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalarrivalareatype *TerminalArrivalAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalarrivalareatype.GongGetIdentifier(stage)
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalsegmentpointpropertytype, terminalsegmentpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalsegmentpointpropertytype.GongGetIdentifier(stage)
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(terminalsegmentpointtype, terminalsegmentpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (terminalsegmentpointtype *TerminalSegmentPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return terminalsegmentpointtype.GongGetIdentifier(stage)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timesheetpropertytype, timesheetpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timesheetpropertytype *TimesheetPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return timesheetpropertytype.GongGetIdentifier(stage)
}

func (timesheettype *TimesheetType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timesheettype, timesheettype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timesheettype *TimesheetType) GongGetReferenceIdentifier(stage *Stage) string {
	return timesheettype.GongGetIdentifier(stage)
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(touchdownliftoffpropertytype, touchdownliftoffpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return touchdownliftoffpropertytype.GongGetIdentifier(stage)
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unplannedholdingpropertytype, unplannedholdingpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unplannedholdingpropertytype.GongGetIdentifier(stage)
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unplannedholdingtimeslicepropertytype, unplannedholdingtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unplannedholdingtimeslicepropertytype.GongGetIdentifier(stage)
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unplannedholdingtimeslicetype, unplannedholdingtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return unplannedholdingtimeslicetype.GongGetIdentifier(stage)
}

func (unplannedholdingtype *UnplannedHoldingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unplannedholdingtype, unplannedholdingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unplannedholdingtype *UnplannedHoldingType) GongGetReferenceIdentifier(stage *Stage) string {
	return unplannedholdingtype.GongGetIdentifier(stage)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurepropertytype, verticalstructurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurepropertytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a_extension *A_extension) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension.GongGetIdentifier(stage), "A_extension", a_extension.Name)
}

func (a_extension_1 *A_extension_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_1.GongGetIdentifier(stage), "A_extension_1", a_extension_1.Name)
}

func (a_extension_10 *A_extension_10) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_10.GongGetIdentifier(stage), "A_extension_10", a_extension_10.Name)
}

func (a_extension_11 *A_extension_11) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_11.GongGetIdentifier(stage), "A_extension_11", a_extension_11.Name)
}

func (a_extension_12 *A_extension_12) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_12.GongGetIdentifier(stage), "A_extension_12", a_extension_12.Name)
}

func (a_extension_13 *A_extension_13) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_13.GongGetIdentifier(stage), "A_extension_13", a_extension_13.Name)
}

func (a_extension_14 *A_extension_14) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_14.GongGetIdentifier(stage), "A_extension_14", a_extension_14.Name)
}

func (a_extension_15 *A_extension_15) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_15.GongGetIdentifier(stage), "A_extension_15", a_extension_15.Name)
}

func (a_extension_16 *A_extension_16) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_16.GongGetIdentifier(stage), "A_extension_16", a_extension_16.Name)
}

func (a_extension_17 *A_extension_17) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_17.GongGetIdentifier(stage), "A_extension_17", a_extension_17.Name)
}

func (a_extension_18 *A_extension_18) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_18.GongGetIdentifier(stage), "A_extension_18", a_extension_18.Name)
}

func (a_extension_19 *A_extension_19) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_19.GongGetIdentifier(stage), "A_extension_19", a_extension_19.Name)
}

func (a_extension_2 *A_extension_2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_2.GongGetIdentifier(stage), "A_extension_2", a_extension_2.Name)
}

func (a_extension_20 *A_extension_20) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_20.GongGetIdentifier(stage), "A_extension_20", a_extension_20.Name)
}

func (a_extension_21 *A_extension_21) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_21.GongGetIdentifier(stage), "A_extension_21", a_extension_21.Name)
}

func (a_extension_22 *A_extension_22) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_22.GongGetIdentifier(stage), "A_extension_22", a_extension_22.Name)
}

func (a_extension_23 *A_extension_23) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_23.GongGetIdentifier(stage), "A_extension_23", a_extension_23.Name)
}

func (a_extension_24 *A_extension_24) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_24.GongGetIdentifier(stage), "A_extension_24", a_extension_24.Name)
}

func (a_extension_25 *A_extension_25) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_25.GongGetIdentifier(stage), "A_extension_25", a_extension_25.Name)
}

func (a_extension_26 *A_extension_26) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_26.GongGetIdentifier(stage), "A_extension_26", a_extension_26.Name)
}

func (a_extension_27 *A_extension_27) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_27.GongGetIdentifier(stage), "A_extension_27", a_extension_27.Name)
}

func (a_extension_28 *A_extension_28) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_28.GongGetIdentifier(stage), "A_extension_28", a_extension_28.Name)
}

func (a_extension_29 *A_extension_29) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_29.GongGetIdentifier(stage), "A_extension_29", a_extension_29.Name)
}

func (a_extension_3 *A_extension_3) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_3.GongGetIdentifier(stage), "A_extension_3", a_extension_3.Name)
}

func (a_extension_30 *A_extension_30) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_30.GongGetIdentifier(stage), "A_extension_30", a_extension_30.Name)
}

func (a_extension_31 *A_extension_31) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_31.GongGetIdentifier(stage), "A_extension_31", a_extension_31.Name)
}

func (a_extension_32 *A_extension_32) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_32.GongGetIdentifier(stage), "A_extension_32", a_extension_32.Name)
}

func (a_extension_33 *A_extension_33) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_33.GongGetIdentifier(stage), "A_extension_33", a_extension_33.Name)
}

func (a_extension_34 *A_extension_34) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_34.GongGetIdentifier(stage), "A_extension_34", a_extension_34.Name)
}

func (a_extension_35 *A_extension_35) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_35.GongGetIdentifier(stage), "A_extension_35", a_extension_35.Name)
}

func (a_extension_36 *A_extension_36) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_36.GongGetIdentifier(stage), "A_extension_36", a_extension_36.Name)
}

func (a_extension_37 *A_extension_37) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_37.GongGetIdentifier(stage), "A_extension_37", a_extension_37.Name)
}

func (a_extension_38 *A_extension_38) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_38.GongGetIdentifier(stage), "A_extension_38", a_extension_38.Name)
}

func (a_extension_39 *A_extension_39) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_39.GongGetIdentifier(stage), "A_extension_39", a_extension_39.Name)
}

func (a_extension_4 *A_extension_4) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_4.GongGetIdentifier(stage), "A_extension_4", a_extension_4.Name)
}

func (a_extension_40 *A_extension_40) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_40.GongGetIdentifier(stage), "A_extension_40", a_extension_40.Name)
}

func (a_extension_41 *A_extension_41) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_41.GongGetIdentifier(stage), "A_extension_41", a_extension_41.Name)
}

func (a_extension_42 *A_extension_42) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_42.GongGetIdentifier(stage), "A_extension_42", a_extension_42.Name)
}

func (a_extension_43 *A_extension_43) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_43.GongGetIdentifier(stage), "A_extension_43", a_extension_43.Name)
}

func (a_extension_44 *A_extension_44) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_44.GongGetIdentifier(stage), "A_extension_44", a_extension_44.Name)
}

func (a_extension_45 *A_extension_45) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_45.GongGetIdentifier(stage), "A_extension_45", a_extension_45.Name)
}

func (a_extension_46 *A_extension_46) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_46.GongGetIdentifier(stage), "A_extension_46", a_extension_46.Name)
}

func (a_extension_47 *A_extension_47) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_47.GongGetIdentifier(stage), "A_extension_47", a_extension_47.Name)
}

func (a_extension_48 *A_extension_48) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_48.GongGetIdentifier(stage), "A_extension_48", a_extension_48.Name)
}

func (a_extension_49 *A_extension_49) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_49.GongGetIdentifier(stage), "A_extension_49", a_extension_49.Name)
}

func (a_extension_5 *A_extension_5) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_5.GongGetIdentifier(stage), "A_extension_5", a_extension_5.Name)
}

func (a_extension_50 *A_extension_50) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_50.GongGetIdentifier(stage), "A_extension_50", a_extension_50.Name)
}

func (a_extension_51 *A_extension_51) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_51.GongGetIdentifier(stage), "A_extension_51", a_extension_51.Name)
}

func (a_extension_52 *A_extension_52) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_52.GongGetIdentifier(stage), "A_extension_52", a_extension_52.Name)
}

func (a_extension_53 *A_extension_53) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_53.GongGetIdentifier(stage), "A_extension_53", a_extension_53.Name)
}

func (a_extension_54 *A_extension_54) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_54.GongGetIdentifier(stage), "A_extension_54", a_extension_54.Name)
}

func (a_extension_55 *A_extension_55) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_55.GongGetIdentifier(stage), "A_extension_55", a_extension_55.Name)
}

func (a_extension_56 *A_extension_56) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_56.GongGetIdentifier(stage), "A_extension_56", a_extension_56.Name)
}

func (a_extension_57 *A_extension_57) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_57.GongGetIdentifier(stage), "A_extension_57", a_extension_57.Name)
}

func (a_extension_58 *A_extension_58) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_58.GongGetIdentifier(stage), "A_extension_58", a_extension_58.Name)
}

func (a_extension_6 *A_extension_6) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_6.GongGetIdentifier(stage), "A_extension_6", a_extension_6.Name)
}

func (a_extension_7 *A_extension_7) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_7.GongGetIdentifier(stage), "A_extension_7", a_extension_7.Name)
}

func (a_extension_8 *A_extension_8) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_8.GongGetIdentifier(stage), "A_extension_8", a_extension_8.Name)
}

func (a_extension_9 *A_extension_9) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_9.GongGetIdentifier(stage), "A_extension_9", a_extension_9.Name)
}

func (abstractapproachlegtype *AbstractApproachLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractapproachlegtype.GongGetIdentifier(stage), "AbstractApproachLegType", abstractapproachlegtype.Name)
}

func (abstractproceduretype *AbstractProcedureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractproceduretype.GongGetIdentifier(stage), "AbstractProcedureType", abstractproceduretype.Name)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractsegmentlegtype.GongGetIdentifier(stage), "AbstractSegmentLegType", abstractsegmentlegtype.Name)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractsegmentpointtype.GongGetIdentifier(stage), "AbstractSegmentPointType", abstractsegmentpointtype.Name)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingpointtype.GongGetIdentifier(stage), "AerialRefuellingPointType", aerialrefuellingpointtype.Name)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristicpropertytype.GongGetIdentifier(stage), "AircraftCharacteristicPropertyType", aircraftcharacteristicpropertytype.Name)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristictype.GongGetIdentifier(stage), "AircraftCharacteristicType", aircraftcharacteristictype.Name)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportheliportpropertytype.GongGetIdentifier(stage), "AirportHeliportPropertyType", airportheliportpropertytype.Name)
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(altimetersourcepropertytype.GongGetIdentifier(stage), "AltimeterSourcePropertyType", altimetersourcepropertytype.Name)
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(altitudeadjustmentpropertytype.GongGetIdentifier(stage), "AltitudeAdjustmentPropertyType", altitudeadjustmentpropertytype.Name)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(altitudeadjustmenttype.GongGetIdentifier(stage), "AltitudeAdjustmentType", altitudeadjustmenttype.Name)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleindicationpropertytype.GongGetIdentifier(stage), "AngleIndicationPropertyType", angleindicationpropertytype.Name)
}

func (angleusepropertytype *AngleUsePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleusepropertytype.GongGetIdentifier(stage), "AngleUsePropertyType", angleusepropertytype.Name)
}

func (angleusetype *AngleUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleusetype.GongGetIdentifier(stage), "AngleUseType", angleusetype.Name)
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachaltitudetablepropertytype.GongGetIdentifier(stage), "ApproachAltitudeTablePropertyType", approachaltitudetablepropertytype.Name)
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachaltitudetabletype.GongGetIdentifier(stage), "ApproachAltitudeTableType", approachaltitudetabletype.Name)
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachconditionpropertytype.GongGetIdentifier(stage), "ApproachConditionPropertyType", approachconditionpropertytype.Name)
}

func (approachconditiontype *ApproachConditionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachconditiontype.GongGetIdentifier(stage), "ApproachConditionType", approachconditiontype.Name)
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachdistancetablepropertytype.GongGetIdentifier(stage), "ApproachDistanceTablePropertyType", approachdistancetablepropertytype.Name)
}

func (approachdistancetabletype *ApproachDistanceTableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachdistancetabletype.GongGetIdentifier(stage), "ApproachDistanceTableType", approachdistancetabletype.Name)
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachtimingtablepropertytype.GongGetIdentifier(stage), "ApproachTimingTablePropertyType", approachtimingtablepropertytype.Name)
}

func (approachtimingtabletype *ApproachTimingTableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(approachtimingtabletype.GongGetIdentifier(stage), "ApproachTimingTableType", approachtimingtabletype.Name)
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivalfeederlegtimeslicepropertytype.GongGetIdentifier(stage), "ArrivalFeederLegTimeSlicePropertyType", arrivalfeederlegtimeslicepropertytype.Name)
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivalfeederlegtimeslicetype.GongGetIdentifier(stage), "ArrivalFeederLegTimeSliceType", arrivalfeederlegtimeslicetype.Name)
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivalfeederlegtype.GongGetIdentifier(stage), "ArrivalFeederLegType", arrivalfeederlegtype.Name)
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivallegtimeslicepropertytype.GongGetIdentifier(stage), "ArrivalLegTimeSlicePropertyType", arrivallegtimeslicepropertytype.Name)
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivallegtimeslicetype.GongGetIdentifier(stage), "ArrivalLegTimeSliceType", arrivallegtimeslicetype.Name)
}

func (arrivallegtype *ArrivalLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrivallegtype.GongGetIdentifier(stage), "ArrivalLegType", arrivallegtype.Name)
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlesectorpropertytype.GongGetIdentifier(stage), "CircleSectorPropertyType", circlesectorpropertytype.Name)
}

func (circlesectortype *CircleSectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlesectortype.GongGetIdentifier(stage), "CircleSectorType", circlesectortype.Name)
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlingareatimeslicepropertytype.GongGetIdentifier(stage), "CirclingAreaTimeSlicePropertyType", circlingareatimeslicepropertytype.Name)
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlingareatimeslicetype.GongGetIdentifier(stage), "CirclingAreaTimeSliceType", circlingareatimeslicetype.Name)
}

func (circlingareatype *CirclingAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlingareatype.GongGetIdentifier(stage), "CirclingAreaType", circlingareatype.Name)
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlingrestrictionpropertytype.GongGetIdentifier(stage), "CirclingRestrictionPropertyType", circlingrestrictionpropertytype.Name)
}

func (circlingrestrictiontype *CirclingRestrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlingrestrictiontype.GongGetIdentifier(stage), "CirclingRestrictionType", circlingrestrictiontype.Name)
}

func (curvepropertytype *CurvePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvepropertytype.GongGetIdentifier(stage), "CurvePropertyType", curvepropertytype.Name)
}

func (curvetype *CurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvetype.GongGetIdentifier(stage), "CurveType", curvetype.Name)
}

func (dmepropertytype *DMEPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dmepropertytype.GongGetIdentifier(stage), "DMEPropertyType", dmepropertytype.Name)
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(departurearrivalconditionpropertytype.GongGetIdentifier(stage), "DepartureArrivalConditionPropertyType", departurearrivalconditionpropertytype.Name)
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(departurearrivalconditiontype.GongGetIdentifier(stage), "DepartureArrivalConditionType", departurearrivalconditiontype.Name)
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(departurelegtimeslicepropertytype.GongGetIdentifier(stage), "DepartureLegTimeSlicePropertyType", departurelegtimeslicepropertytype.Name)
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(departurelegtimeslicetype.GongGetIdentifier(stage), "DepartureLegTimeSliceType", departurelegtimeslicetype.Name)
}

func (departurelegtype *DepartureLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(departurelegtype.GongGetIdentifier(stage), "DepartureLegType", departurelegtype.Name)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointpropertytype.GongGetIdentifier(stage), "DesignatedPointPropertyType", designatedpointpropertytype.Name)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distanceindicationpropertytype.GongGetIdentifier(stage), "DistanceIndicationPropertyType", distanceindicationpropertytype.Name)
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedcurvetype.GongGetIdentifier(stage), "ElevatedCurveType", elevatedcurvetype.Name)
}

func (elevatedpointtype *ElevatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointtype.GongGetIdentifier(stage), "ElevatedPointType", elevatedpointtype.Name)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedsurfacetype.GongGetIdentifier(stage), "ElevatedSurfaceType", elevatedsurfacetype.Name)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(enroutesegmentpointtype.GongGetIdentifier(stage), "EnRouteSegmentPointType", enroutesegmentpointtype.Name)
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(equipmentunavailableadjustmentcolumnpropertytype.GongGetIdentifier(stage), "EquipmentUnavailableAdjustmentColumnPropertyType", equipmentunavailableadjustmentcolumnpropertytype.Name)
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(equipmentunavailableadjustmentcolumntype.GongGetIdentifier(stage), "EquipmentUnavailableAdjustmentColumnType", equipmentunavailableadjustmentcolumntype.Name)
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(equipmentunavailableadjustmentpropertytype.GongGetIdentifier(stage), "EquipmentUnavailableAdjustmentPropertyType", equipmentunavailableadjustmentpropertytype.Name)
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(equipmentunavailableadjustmenttype.GongGetIdentifier(stage), "EquipmentUnavailableAdjustmentType", equipmentunavailableadjustmenttype.Name)
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fasdatablockpropertytype.GongGetIdentifier(stage), "FASDataBlockPropertyType", fasdatablockpropertytype.Name)
}

func (fasdatablocktype *FASDataBlockType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fasdatablocktype.GongGetIdentifier(stage), "FASDataBlockType", fasdatablocktype.Name)
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(finallegtimeslicepropertytype.GongGetIdentifier(stage), "FinalLegTimeSlicePropertyType", finallegtimeslicepropertytype.Name)
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(finallegtimeslicetype.GongGetIdentifier(stage), "FinalLegTimeSliceType", finallegtimeslicetype.Name)
}

func (finallegtype *FinalLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(finallegtype.GongGetIdentifier(stage), "FinalLegType", finallegtype.Name)
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(finalprofilepropertytype.GongGetIdentifier(stage), "FinalProfilePropertyType", finalprofilepropertytype.Name)
}

func (finalprofiletype *FinalProfileType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(finalprofiletype.GongGetIdentifier(stage), "FinalProfileType", finalprofiletype.Name)
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingassessmenttimeslicepropertytype.GongGetIdentifier(stage), "HoldingAssessmentTimeSlicePropertyType", holdingassessmenttimeslicepropertytype.Name)
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingassessmenttimeslicetype.GongGetIdentifier(stage), "HoldingAssessmentTimeSliceType", holdingassessmenttimeslicetype.Name)
}

func (holdingassessmenttype *HoldingAssessmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingassessmenttype.GongGetIdentifier(stage), "HoldingAssessmentType", holdingassessmenttype.Name)
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterndistancepropertytype.GongGetIdentifier(stage), "HoldingPatternDistancePropertyType", holdingpatterndistancepropertytype.Name)
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterndistancetype.GongGetIdentifier(stage), "HoldingPatternDistanceType", holdingpatterndistancetype.Name)
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterndurationpropertytype.GongGetIdentifier(stage), "HoldingPatternDurationPropertyType", holdingpatterndurationpropertytype.Name)
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterndurationtype.GongGetIdentifier(stage), "HoldingPatternDurationType", holdingpatterndurationtype.Name)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatternpropertytype.GongGetIdentifier(stage), "HoldingPatternPropertyType", holdingpatternpropertytype.Name)
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterntimeslicepropertytype.GongGetIdentifier(stage), "HoldingPatternTimeSlicePropertyType", holdingpatterntimeslicepropertytype.Name)
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterntimeslicetype.GongGetIdentifier(stage), "HoldingPatternTimeSliceType", holdingpatterntimeslicetype.Name)
}

func (holdingpatterntype *HoldingPatternType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatterntype.GongGetIdentifier(stage), "HoldingPatternType", holdingpatterntype.Name)
}

func (holdingusepropertytype *HoldingUsePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingusepropertytype.GongGetIdentifier(stage), "HoldingUsePropertyType", holdingusepropertytype.Name)
}

func (holdingusetype *HoldingUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingusetype.GongGetIdentifier(stage), "HoldingUseType", holdingusetype.Name)
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(initiallegtimeslicepropertytype.GongGetIdentifier(stage), "InitialLegTimeSlicePropertyType", initiallegtimeslicepropertytype.Name)
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(initiallegtimeslicetype.GongGetIdentifier(stage), "InitialLegTimeSliceType", initiallegtimeslicetype.Name)
}

func (initiallegtype *InitialLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(initiallegtype.GongGetIdentifier(stage), "InitialLegType", initiallegtype.Name)
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentapproachprocedurepropertytype.GongGetIdentifier(stage), "InstrumentApproachProcedurePropertyType", instrumentapproachprocedurepropertytype.Name)
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentapproachproceduretimeslicepropertytype.GongGetIdentifier(stage), "InstrumentApproachProcedureTimeSlicePropertyType", instrumentapproachproceduretimeslicepropertytype.Name)
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentapproachproceduretimeslicetype.GongGetIdentifier(stage), "InstrumentApproachProcedureTimeSliceType", instrumentapproachproceduretimeslicetype.Name)
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentapproachproceduretype.GongGetIdentifier(stage), "InstrumentApproachProcedureType", instrumentapproachproceduretype.Name)
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(intermediatelegtimeslicepropertytype.GongGetIdentifier(stage), "IntermediateLegTimeSlicePropertyType", intermediatelegtimeslicepropertytype.Name)
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(intermediatelegtimeslicetype.GongGetIdentifier(stage), "IntermediateLegTimeSliceType", intermediatelegtimeslicetype.Name)
}

func (intermediatelegtype *IntermediateLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(intermediatelegtype.GongGetIdentifier(stage), "IntermediateLegType", intermediatelegtype.Name)
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(landingtakeoffareacollectionpropertytype.GongGetIdentifier(stage), "LandingTakeoffAreaCollectionPropertyType", landingtakeoffareacollectionpropertytype.Name)
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(landingtakeoffareacollectiontype.GongGetIdentifier(stage), "LandingTakeoffAreaCollectionType", landingtakeoffareacollectiontype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (minimapropertytype *MinimaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(minimapropertytype.GongGetIdentifier(stage), "MinimaPropertyType", minimapropertytype.Name)
}

func (minimatype *MinimaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(minimatype.GongGetIdentifier(stage), "MinimaType", minimatype.Name)
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(missedapproachgrouppropertytype.GongGetIdentifier(stage), "MissedApproachGroupPropertyType", missedapproachgrouppropertytype.Name)
}

func (missedapproachgrouptype *MissedApproachGroupType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(missedapproachgrouptype.GongGetIdentifier(stage), "MissedApproachGroupType", missedapproachgrouptype.Name)
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(missedapproachlegtimeslicepropertytype.GongGetIdentifier(stage), "MissedApproachLegTimeSlicePropertyType", missedapproachlegtimeslicepropertytype.Name)
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(missedapproachlegtimeslicetype.GongGetIdentifier(stage), "MissedApproachLegTimeSliceType", missedapproachlegtimeslicetype.Name)
}

func (missedapproachlegtype *MissedApproachLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(missedapproachlegtype.GongGetIdentifier(stage), "MissedApproachLegType", missedapproachlegtype.Name)
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidpropertytype.GongGetIdentifier(stage), "NavaidPropertyType", navaidpropertytype.Name)
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationarearestrictiontimeslicepropertytype.GongGetIdentifier(stage), "NavigationAreaRestrictionTimeSlicePropertyType", navigationarearestrictiontimeslicepropertytype.Name)
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationarearestrictiontimeslicetype.GongGetIdentifier(stage), "NavigationAreaRestrictionTimeSliceType", navigationarearestrictiontimeslicetype.Name)
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationarearestrictiontype.GongGetIdentifier(stage), "NavigationAreaRestrictionType", navigationarearestrictiontype.Name)
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationareasectorpropertytype.GongGetIdentifier(stage), "NavigationAreaSectorPropertyType", navigationareasectorpropertytype.Name)
}

func (navigationareasectortype *NavigationAreaSectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationareasectortype.GongGetIdentifier(stage), "NavigationAreaSectorType", navigationareasectortype.Name)
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationareatimeslicepropertytype.GongGetIdentifier(stage), "NavigationAreaTimeSlicePropertyType", navigationareatimeslicepropertytype.Name)
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationareatimeslicetype.GongGetIdentifier(stage), "NavigationAreaTimeSliceType", navigationareatimeslicetype.Name)
}

func (navigationareatype *NavigationAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navigationareatype.GongGetIdentifier(stage), "NavigationAreaType", navigationareatype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleassessmentareapropertytype.GongGetIdentifier(stage), "ObstacleAssessmentAreaPropertyType", obstacleassessmentareapropertytype.Name)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleassessmentareatype.GongGetIdentifier(stage), "ObstacleAssessmentAreaType", obstacleassessmentareatype.Name)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleplacementpropertytype.GongGetIdentifier(stage), "ObstaclePlacementPropertyType", obstacleplacementpropertytype.Name)
}

func (obstacleplacementtype *ObstaclePlacementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleplacementtype.GongGetIdentifier(stage), "ObstaclePlacementType", obstacleplacementtype.Name)
}

func (obstructionpropertytype *ObstructionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstructionpropertytype.GongGetIdentifier(stage), "ObstructionPropertyType", obstructionpropertytype.Name)
}

func (obstructiontype *ObstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstructiontype.GongGetIdentifier(stage), "ObstructionType", obstructiontype.Name)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritypropertytype.GongGetIdentifier(stage), "OrganisationAuthorityPropertyType", organisationauthoritypropertytype.Name)
}

func (pointpropertytype *PointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointpropertytype.GongGetIdentifier(stage), "PointPropertyType", pointpropertytype.Name)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointreferencepropertytype.GongGetIdentifier(stage), "PointReferencePropertyType", pointreferencepropertytype.Name)
}

func (pointreferencetype *PointReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointreferencetype.GongGetIdentifier(stage), "PointReferenceType", pointreferencetype.Name)
}

func (pointtype *PointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointtype.GongGetIdentifier(stage), "PointType", pointtype.Name)
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(procedureavailabilitypropertytype.GongGetIdentifier(stage), "ProcedureAvailabilityPropertyType", procedureavailabilitypropertytype.Name)
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(procedureavailabilitytype.GongGetIdentifier(stage), "ProcedureAvailabilityType", procedureavailabilitytype.Name)
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduredmetimeslicepropertytype.GongGetIdentifier(stage), "ProcedureDMETimeSlicePropertyType", proceduredmetimeslicepropertytype.Name)
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduredmetimeslicetype.GongGetIdentifier(stage), "ProcedureDMETimeSliceType", proceduredmetimeslicetype.Name)
}

func (proceduredmetype *ProcedureDMEType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduredmetype.GongGetIdentifier(stage), "ProcedureDMEType", proceduredmetype.Name)
}

func (procedurepropertytype *ProcedurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(procedurepropertytype.GongGetIdentifier(stage), "ProcedurePropertyType", procedurepropertytype.Name)
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduretransitionlegpropertytype.GongGetIdentifier(stage), "ProcedureTransitionLegPropertyType", proceduretransitionlegpropertytype.Name)
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduretransitionlegtype.GongGetIdentifier(stage), "ProcedureTransitionLegType", proceduretransitionlegtype.Name)
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduretransitionpropertytype.GongGetIdentifier(stage), "ProcedureTransitionPropertyType", proceduretransitionpropertytype.Name)
}

func (proceduretransitiontype *ProcedureTransitionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(proceduretransitiontype.GongGetIdentifier(stage), "ProcedureTransitionType", proceduretransitiontype.Name)
}

func (radarsystempropertytype *RadarSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarsystempropertytype.GongGetIdentifier(stage), "RadarSystemPropertyType", radarsystempropertytype.Name)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiofrequencyareapropertytype.GongGetIdentifier(stage), "RadioFrequencyAreaPropertyType", radiofrequencyareapropertytype.Name)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaycentrelinepointpropertytype.GongGetIdentifier(stage), "RunwayCentrelinePointPropertyType", runwaycentrelinepointpropertytype.Name)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaydirectionpropertytype.GongGetIdentifier(stage), "RunwayDirectionPropertyType", runwaydirectionpropertytype.Name)
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareapropertytype.GongGetIdentifier(stage), "SafeAltitudeAreaPropertyType", safealtitudeareapropertytype.Name)
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareasectorpropertytype.GongGetIdentifier(stage), "SafeAltitudeAreaSectorPropertyType", safealtitudeareasectorpropertytype.Name)
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareasectortype.GongGetIdentifier(stage), "SafeAltitudeAreaSectorType", safealtitudeareasectortype.Name)
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareatimeslicepropertytype.GongGetIdentifier(stage), "SafeAltitudeAreaTimeSlicePropertyType", safealtitudeareatimeslicepropertytype.Name)
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareatimeslicetype.GongGetIdentifier(stage), "SafeAltitudeAreaTimeSliceType", safealtitudeareatimeslicetype.Name)
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(safealtitudeareatype.GongGetIdentifier(stage), "SafeAltitudeAreaType", safealtitudeareatype.Name)
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sectordesignpropertytype.GongGetIdentifier(stage), "SectorDesignPropertyType", sectordesignpropertytype.Name)
}

func (sectordesigntype *SectorDesignType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sectordesigntype.GongGetIdentifier(stage), "SectorDesignType", sectordesigntype.Name)
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(segmentlegpropertytype.GongGetIdentifier(stage), "SegmentLegPropertyType", segmentlegpropertytype.Name)
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(segmentpointpropertytype.GongGetIdentifier(stage), "SegmentPointPropertyType", segmentpointpropertytype.Name)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationsystempropertytype.GongGetIdentifier(stage), "SpecialNavigationSystemPropertyType", specialnavigationsystempropertytype.Name)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentarrivalpropertytype.GongGetIdentifier(stage), "StandardInstrumentArrivalPropertyType", standardinstrumentarrivalpropertytype.Name)
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentarrivaltimeslicepropertytype.GongGetIdentifier(stage), "StandardInstrumentArrivalTimeSlicePropertyType", standardinstrumentarrivaltimeslicepropertytype.Name)
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentarrivaltimeslicetype.GongGetIdentifier(stage), "StandardInstrumentArrivalTimeSliceType", standardinstrumentarrivaltimeslicetype.Name)
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentarrivaltype.GongGetIdentifier(stage), "StandardInstrumentArrivalType", standardinstrumentarrivaltype.Name)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentdeparturepropertytype.GongGetIdentifier(stage), "StandardInstrumentDeparturePropertyType", standardinstrumentdeparturepropertytype.Name)
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentdeparturetimeslicepropertytype.GongGetIdentifier(stage), "StandardInstrumentDepartureTimeSlicePropertyType", standardinstrumentdeparturetimeslicepropertytype.Name)
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentdeparturetimeslicetype.GongGetIdentifier(stage), "StandardInstrumentDepartureTimeSliceType", standardinstrumentdeparturetimeslicetype.Name)
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentdeparturetype.GongGetIdentifier(stage), "StandardInstrumentDepartureType", standardinstrumentdeparturetype.Name)
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacepropertytype.GongGetIdentifier(stage), "SurfacePropertyType", surfacepropertytype.Name)
}

func (surfacetype *SurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacetype.GongGetIdentifier(stage), "SurfaceType", surfacetype.Name)
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalarrivalareasectorpropertytype.GongGetIdentifier(stage), "TerminalArrivalAreaSectorPropertyType", terminalarrivalareasectorpropertytype.Name)
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalarrivalareasectortype.GongGetIdentifier(stage), "TerminalArrivalAreaSectorType", terminalarrivalareasectortype.Name)
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalarrivalareatimeslicepropertytype.GongGetIdentifier(stage), "TerminalArrivalAreaTimeSlicePropertyType", terminalarrivalareatimeslicepropertytype.Name)
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalarrivalareatimeslicetype.GongGetIdentifier(stage), "TerminalArrivalAreaTimeSliceType", terminalarrivalareatimeslicetype.Name)
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalarrivalareatype.GongGetIdentifier(stage), "TerminalArrivalAreaType", terminalarrivalareatype.Name)
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalsegmentpointpropertytype.GongGetIdentifier(stage), "TerminalSegmentPointPropertyType", terminalsegmentpointpropertytype.Name)
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(terminalsegmentpointtype.GongGetIdentifier(stage), "TerminalSegmentPointType", terminalsegmentpointtype.Name)
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheetpropertytype.GongGetIdentifier(stage), "TimesheetPropertyType", timesheetpropertytype.Name)
}

func (timesheettype *TimesheetType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheettype.GongGetIdentifier(stage), "TimesheetType", timesheettype.Name)
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(touchdownliftoffpropertytype.GongGetIdentifier(stage), "TouchDownLiftOffPropertyType", touchdownliftoffpropertytype.Name)
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unplannedholdingpropertytype.GongGetIdentifier(stage), "UnplannedHoldingPropertyType", unplannedholdingpropertytype.Name)
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unplannedholdingtimeslicepropertytype.GongGetIdentifier(stage), "UnplannedHoldingTimeSlicePropertyType", unplannedholdingtimeslicepropertytype.Name)
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unplannedholdingtimeslicetype.GongGetIdentifier(stage), "UnplannedHoldingTimeSliceType", unplannedholdingtimeslicetype.Name)
}

func (unplannedholdingtype *UnplannedHoldingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unplannedholdingtype.GongGetIdentifier(stage), "UnplannedHoldingType", unplannedholdingtype.Name)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurepropertytype.GongGetIdentifier(stage), "VerticalStructurePropertyType", verticalstructurepropertytype.Name)
}

// insertion point for unstaging
func (a_extension *A_extension) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension.GongGetReferenceIdentifier(stage))
}

func (a_extension_1 *A_extension_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_1.GongGetReferenceIdentifier(stage))
}

func (a_extension_10 *A_extension_10) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_10.GongGetReferenceIdentifier(stage))
}

func (a_extension_11 *A_extension_11) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_11.GongGetReferenceIdentifier(stage))
}

func (a_extension_12 *A_extension_12) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_12.GongGetReferenceIdentifier(stage))
}

func (a_extension_13 *A_extension_13) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_13.GongGetReferenceIdentifier(stage))
}

func (a_extension_14 *A_extension_14) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_14.GongGetReferenceIdentifier(stage))
}

func (a_extension_15 *A_extension_15) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_15.GongGetReferenceIdentifier(stage))
}

func (a_extension_16 *A_extension_16) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_16.GongGetReferenceIdentifier(stage))
}

func (a_extension_17 *A_extension_17) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_17.GongGetReferenceIdentifier(stage))
}

func (a_extension_18 *A_extension_18) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_18.GongGetReferenceIdentifier(stage))
}

func (a_extension_19 *A_extension_19) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_19.GongGetReferenceIdentifier(stage))
}

func (a_extension_2 *A_extension_2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_2.GongGetReferenceIdentifier(stage))
}

func (a_extension_20 *A_extension_20) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_20.GongGetReferenceIdentifier(stage))
}

func (a_extension_21 *A_extension_21) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_21.GongGetReferenceIdentifier(stage))
}

func (a_extension_22 *A_extension_22) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_22.GongGetReferenceIdentifier(stage))
}

func (a_extension_23 *A_extension_23) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_23.GongGetReferenceIdentifier(stage))
}

func (a_extension_24 *A_extension_24) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_24.GongGetReferenceIdentifier(stage))
}

func (a_extension_25 *A_extension_25) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_25.GongGetReferenceIdentifier(stage))
}

func (a_extension_26 *A_extension_26) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_26.GongGetReferenceIdentifier(stage))
}

func (a_extension_27 *A_extension_27) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_27.GongGetReferenceIdentifier(stage))
}

func (a_extension_28 *A_extension_28) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_28.GongGetReferenceIdentifier(stage))
}

func (a_extension_29 *A_extension_29) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_29.GongGetReferenceIdentifier(stage))
}

func (a_extension_3 *A_extension_3) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_3.GongGetReferenceIdentifier(stage))
}

func (a_extension_30 *A_extension_30) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_30.GongGetReferenceIdentifier(stage))
}

func (a_extension_31 *A_extension_31) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_31.GongGetReferenceIdentifier(stage))
}

func (a_extension_32 *A_extension_32) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_32.GongGetReferenceIdentifier(stage))
}

func (a_extension_33 *A_extension_33) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_33.GongGetReferenceIdentifier(stage))
}

func (a_extension_34 *A_extension_34) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_34.GongGetReferenceIdentifier(stage))
}

func (a_extension_35 *A_extension_35) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_35.GongGetReferenceIdentifier(stage))
}

func (a_extension_36 *A_extension_36) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_36.GongGetReferenceIdentifier(stage))
}

func (a_extension_37 *A_extension_37) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_37.GongGetReferenceIdentifier(stage))
}

func (a_extension_38 *A_extension_38) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_38.GongGetReferenceIdentifier(stage))
}

func (a_extension_39 *A_extension_39) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_39.GongGetReferenceIdentifier(stage))
}

func (a_extension_4 *A_extension_4) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_4.GongGetReferenceIdentifier(stage))
}

func (a_extension_40 *A_extension_40) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_40.GongGetReferenceIdentifier(stage))
}

func (a_extension_41 *A_extension_41) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_41.GongGetReferenceIdentifier(stage))
}

func (a_extension_42 *A_extension_42) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_42.GongGetReferenceIdentifier(stage))
}

func (a_extension_43 *A_extension_43) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_43.GongGetReferenceIdentifier(stage))
}

func (a_extension_44 *A_extension_44) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_44.GongGetReferenceIdentifier(stage))
}

func (a_extension_45 *A_extension_45) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_45.GongGetReferenceIdentifier(stage))
}

func (a_extension_46 *A_extension_46) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_46.GongGetReferenceIdentifier(stage))
}

func (a_extension_47 *A_extension_47) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_47.GongGetReferenceIdentifier(stage))
}

func (a_extension_48 *A_extension_48) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_48.GongGetReferenceIdentifier(stage))
}

func (a_extension_49 *A_extension_49) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_49.GongGetReferenceIdentifier(stage))
}

func (a_extension_5 *A_extension_5) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_5.GongGetReferenceIdentifier(stage))
}

func (a_extension_50 *A_extension_50) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_50.GongGetReferenceIdentifier(stage))
}

func (a_extension_51 *A_extension_51) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_51.GongGetReferenceIdentifier(stage))
}

func (a_extension_52 *A_extension_52) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_52.GongGetReferenceIdentifier(stage))
}

func (a_extension_53 *A_extension_53) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_53.GongGetReferenceIdentifier(stage))
}

func (a_extension_54 *A_extension_54) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_54.GongGetReferenceIdentifier(stage))
}

func (a_extension_55 *A_extension_55) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_55.GongGetReferenceIdentifier(stage))
}

func (a_extension_56 *A_extension_56) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_56.GongGetReferenceIdentifier(stage))
}

func (a_extension_57 *A_extension_57) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_57.GongGetReferenceIdentifier(stage))
}

func (a_extension_58 *A_extension_58) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_58.GongGetReferenceIdentifier(stage))
}

func (a_extension_6 *A_extension_6) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_6.GongGetReferenceIdentifier(stage))
}

func (a_extension_7 *A_extension_7) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_7.GongGetReferenceIdentifier(stage))
}

func (a_extension_8 *A_extension_8) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_8.GongGetReferenceIdentifier(stage))
}

func (a_extension_9 *A_extension_9) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_9.GongGetReferenceIdentifier(stage))
}

func (abstractapproachlegtype *AbstractApproachLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractapproachlegtype.GongGetReferenceIdentifier(stage))
}

func (abstractproceduretype *AbstractProcedureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractproceduretype.GongGetReferenceIdentifier(stage))
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractsegmentlegtype.GongGetReferenceIdentifier(stage))
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractsegmentpointtype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingpointtype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristicpropertytype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristictype.GongGetReferenceIdentifier(stage))
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportheliportpropertytype.GongGetReferenceIdentifier(stage))
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(altimetersourcepropertytype.GongGetReferenceIdentifier(stage))
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(altitudeadjustmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(altitudeadjustmenttype.GongGetReferenceIdentifier(stage))
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleindicationpropertytype.GongGetReferenceIdentifier(stage))
}

func (angleusepropertytype *AngleUsePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleusepropertytype.GongGetReferenceIdentifier(stage))
}

func (angleusetype *AngleUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleusetype.GongGetReferenceIdentifier(stage))
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachaltitudetablepropertytype.GongGetReferenceIdentifier(stage))
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachaltitudetabletype.GongGetReferenceIdentifier(stage))
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachconditionpropertytype.GongGetReferenceIdentifier(stage))
}

func (approachconditiontype *ApproachConditionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachconditiontype.GongGetReferenceIdentifier(stage))
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachdistancetablepropertytype.GongGetReferenceIdentifier(stage))
}

func (approachdistancetabletype *ApproachDistanceTableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachdistancetabletype.GongGetReferenceIdentifier(stage))
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachtimingtablepropertytype.GongGetReferenceIdentifier(stage))
}

func (approachtimingtabletype *ApproachTimingTableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(approachtimingtabletype.GongGetReferenceIdentifier(stage))
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivalfeederlegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivalfeederlegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivalfeederlegtype.GongGetReferenceIdentifier(stage))
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivallegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivallegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (arrivallegtype *ArrivalLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrivallegtype.GongGetReferenceIdentifier(stage))
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlesectorpropertytype.GongGetReferenceIdentifier(stage))
}

func (circlesectortype *CircleSectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlesectortype.GongGetReferenceIdentifier(stage))
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlingareatimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlingareatimeslicetype.GongGetReferenceIdentifier(stage))
}

func (circlingareatype *CirclingAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlingareatype.GongGetReferenceIdentifier(stage))
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlingrestrictionpropertytype.GongGetReferenceIdentifier(stage))
}

func (circlingrestrictiontype *CirclingRestrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlingrestrictiontype.GongGetReferenceIdentifier(stage))
}

func (curvepropertytype *CurvePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvepropertytype.GongGetReferenceIdentifier(stage))
}

func (curvetype *CurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvetype.GongGetReferenceIdentifier(stage))
}

func (dmepropertytype *DMEPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dmepropertytype.GongGetReferenceIdentifier(stage))
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(departurearrivalconditionpropertytype.GongGetReferenceIdentifier(stage))
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(departurearrivalconditiontype.GongGetReferenceIdentifier(stage))
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(departurelegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(departurelegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (departurelegtype *DepartureLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(departurelegtype.GongGetReferenceIdentifier(stage))
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distanceindicationpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedcurvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointtype *ElevatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointtype.GongGetReferenceIdentifier(stage))
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedsurfacetype.GongGetReferenceIdentifier(stage))
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(enroutesegmentpointtype.GongGetReferenceIdentifier(stage))
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(equipmentunavailableadjustmentcolumnpropertytype.GongGetReferenceIdentifier(stage))
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(equipmentunavailableadjustmentcolumntype.GongGetReferenceIdentifier(stage))
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(equipmentunavailableadjustmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(equipmentunavailableadjustmenttype.GongGetReferenceIdentifier(stage))
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fasdatablockpropertytype.GongGetReferenceIdentifier(stage))
}

func (fasdatablocktype *FASDataBlockType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fasdatablocktype.GongGetReferenceIdentifier(stage))
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(finallegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(finallegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (finallegtype *FinalLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(finallegtype.GongGetReferenceIdentifier(stage))
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(finalprofilepropertytype.GongGetReferenceIdentifier(stage))
}

func (finalprofiletype *FinalProfileType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(finalprofiletype.GongGetReferenceIdentifier(stage))
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingassessmenttimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingassessmenttimeslicetype.GongGetReferenceIdentifier(stage))
}

func (holdingassessmenttype *HoldingAssessmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingassessmenttype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterndistancepropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterndistancetype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterndurationpropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterndurationtype.GongGetReferenceIdentifier(stage))
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatternpropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterntimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterntimeslicetype.GongGetReferenceIdentifier(stage))
}

func (holdingpatterntype *HoldingPatternType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatterntype.GongGetReferenceIdentifier(stage))
}

func (holdingusepropertytype *HoldingUsePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingusepropertytype.GongGetReferenceIdentifier(stage))
}

func (holdingusetype *HoldingUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingusetype.GongGetReferenceIdentifier(stage))
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(initiallegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(initiallegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (initiallegtype *InitialLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(initiallegtype.GongGetReferenceIdentifier(stage))
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentapproachprocedurepropertytype.GongGetReferenceIdentifier(stage))
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentapproachproceduretimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentapproachproceduretimeslicetype.GongGetReferenceIdentifier(stage))
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentapproachproceduretype.GongGetReferenceIdentifier(stage))
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(intermediatelegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(intermediatelegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (intermediatelegtype *IntermediateLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(intermediatelegtype.GongGetReferenceIdentifier(stage))
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(landingtakeoffareacollectionpropertytype.GongGetReferenceIdentifier(stage))
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(landingtakeoffareacollectiontype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (minimapropertytype *MinimaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(minimapropertytype.GongGetReferenceIdentifier(stage))
}

func (minimatype *MinimaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(minimatype.GongGetReferenceIdentifier(stage))
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(missedapproachgrouppropertytype.GongGetReferenceIdentifier(stage))
}

func (missedapproachgrouptype *MissedApproachGroupType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(missedapproachgrouptype.GongGetReferenceIdentifier(stage))
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(missedapproachlegtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(missedapproachlegtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (missedapproachlegtype *MissedApproachLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(missedapproachlegtype.GongGetReferenceIdentifier(stage))
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidpropertytype.GongGetReferenceIdentifier(stage))
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationarearestrictiontimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationarearestrictiontimeslicetype.GongGetReferenceIdentifier(stage))
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationarearestrictiontype.GongGetReferenceIdentifier(stage))
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationareasectorpropertytype.GongGetReferenceIdentifier(stage))
}

func (navigationareasectortype *NavigationAreaSectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationareasectortype.GongGetReferenceIdentifier(stage))
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationareatimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationareatimeslicetype.GongGetReferenceIdentifier(stage))
}

func (navigationareatype *NavigationAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navigationareatype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleassessmentareapropertytype.GongGetReferenceIdentifier(stage))
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleassessmentareatype.GongGetReferenceIdentifier(stage))
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleplacementpropertytype.GongGetReferenceIdentifier(stage))
}

func (obstacleplacementtype *ObstaclePlacementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleplacementtype.GongGetReferenceIdentifier(stage))
}

func (obstructionpropertytype *ObstructionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstructionpropertytype.GongGetReferenceIdentifier(stage))
}

func (obstructiontype *ObstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstructiontype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritypropertytype.GongGetReferenceIdentifier(stage))
}

func (pointpropertytype *PointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointpropertytype.GongGetReferenceIdentifier(stage))
}

func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointreferencepropertytype.GongGetReferenceIdentifier(stage))
}

func (pointreferencetype *PointReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointreferencetype.GongGetReferenceIdentifier(stage))
}

func (pointtype *PointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointtype.GongGetReferenceIdentifier(stage))
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(procedureavailabilitypropertytype.GongGetReferenceIdentifier(stage))
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(procedureavailabilitytype.GongGetReferenceIdentifier(stage))
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduredmetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduredmetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (proceduredmetype *ProcedureDMEType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduredmetype.GongGetReferenceIdentifier(stage))
}

func (procedurepropertytype *ProcedurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(procedurepropertytype.GongGetReferenceIdentifier(stage))
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduretransitionlegpropertytype.GongGetReferenceIdentifier(stage))
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduretransitionlegtype.GongGetReferenceIdentifier(stage))
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduretransitionpropertytype.GongGetReferenceIdentifier(stage))
}

func (proceduretransitiontype *ProcedureTransitionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(proceduretransitiontype.GongGetReferenceIdentifier(stage))
}

func (radarsystempropertytype *RadarSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiofrequencyareapropertytype.GongGetReferenceIdentifier(stage))
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaycentrelinepointpropertytype.GongGetReferenceIdentifier(stage))
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaydirectionpropertytype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareapropertytype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareasectorpropertytype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareasectortype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareatimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareatimeslicetype.GongGetReferenceIdentifier(stage))
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(safealtitudeareatype.GongGetReferenceIdentifier(stage))
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sectordesignpropertytype.GongGetReferenceIdentifier(stage))
}

func (sectordesigntype *SectorDesignType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sectordesigntype.GongGetReferenceIdentifier(stage))
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(segmentlegpropertytype.GongGetReferenceIdentifier(stage))
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(segmentpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentarrivalpropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentarrivaltimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentarrivaltimeslicetype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentarrivaltype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentdeparturepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentdeparturetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentdeparturetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentdeparturetype.GongGetReferenceIdentifier(stage))
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacepropertytype.GongGetReferenceIdentifier(stage))
}

func (surfacetype *SurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacetype.GongGetReferenceIdentifier(stage))
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalarrivalareasectorpropertytype.GongGetReferenceIdentifier(stage))
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalarrivalareasectortype.GongGetReferenceIdentifier(stage))
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalarrivalareatimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalarrivalareatimeslicetype.GongGetReferenceIdentifier(stage))
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalarrivalareatype.GongGetReferenceIdentifier(stage))
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalsegmentpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(terminalsegmentpointtype.GongGetReferenceIdentifier(stage))
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheetpropertytype.GongGetReferenceIdentifier(stage))
}

func (timesheettype *TimesheetType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheettype.GongGetReferenceIdentifier(stage))
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(touchdownliftoffpropertytype.GongGetReferenceIdentifier(stage))
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unplannedholdingpropertytype.GongGetReferenceIdentifier(stage))
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unplannedholdingtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unplannedholdingtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (unplannedholdingtype *UnplannedHoldingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unplannedholdingtype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurepropertytype.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__aixmfeatures__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
