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
	// Compute reverse map for named struct AeronauticalGroundLightTimeSliceType
	// insertion point per field
	stage.AeronauticalGroundLightTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AeronauticalGroundLightTimeSliceType)
	for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
		_ = aeronauticalgroundlighttimeslicetype
		for _, _notepropertytype := range aeronauticalgroundlighttimeslicetype.Annotation {
			stage.AeronauticalGroundLightTimeSliceType_Annotation_reverseMap[_notepropertytype] = aeronauticalgroundlighttimeslicetype
		}
	}
	stage.AeronauticalGroundLightTimeSliceType_Extension_reverseMap = make(map[*A_extension_3]*AeronauticalGroundLightTimeSliceType)
	for aeronauticalgroundlighttimeslicetype := range stage.AeronauticalGroundLightTimeSliceTypes {
		_ = aeronauticalgroundlighttimeslicetype
		for _, _a_extension_3 := range aeronauticalgroundlighttimeslicetype.Extension {
			stage.AeronauticalGroundLightTimeSliceType_Extension_reverseMap[_a_extension_3] = aeronauticalgroundlighttimeslicetype
		}
	}

	// Compute reverse map for named struct AeronauticalGroundLightType
	// insertion point per field
	stage.AeronauticalGroundLightType_TimeSlice_reverseMap = make(map[*AeronauticalGroundLightTimeSlicePropertyType]*AeronauticalGroundLightType)
	for aeronauticalgroundlighttype := range stage.AeronauticalGroundLightTypes {
		_ = aeronauticalgroundlighttype
		for _, _aeronauticalgroundlighttimeslicepropertytype := range aeronauticalgroundlighttype.TimeSlice {
			stage.AeronauticalGroundLightType_TimeSlice_reverseMap[_aeronauticalgroundlighttimeslicepropertytype] = aeronauticalgroundlighttype
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

	// Compute reverse map for named struct ElevatedCurveType
	// insertion point per field
	stage.ElevatedCurveType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _notepropertytype := range elevatedcurvetype.Annotation {
			stage.ElevatedCurveType_Annotation_reverseMap[_notepropertytype] = elevatedcurvetype
		}
	}
	stage.ElevatedCurveType_Extension_reverseMap = make(map[*A_extension]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _a_extension := range elevatedcurvetype.Extension {
			stage.ElevatedCurveType_Extension_reverseMap[_a_extension] = elevatedcurvetype
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
	stage.ElevatedPointType_Extension_reverseMap = make(map[*A_extension_1]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _a_extension_1 := range elevatedpointtype.Extension {
			stage.ElevatedPointType_Extension_reverseMap[_a_extension_1] = elevatedpointtype
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
	stage.ElevatedSurfaceType_Extension_reverseMap = make(map[*A_extension_2]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _a_extension_2 := range elevatedsurfacetype.Extension {
			stage.ElevatedSurfaceType_Extension_reverseMap[_a_extension_2] = elevatedsurfacetype
		}
	}

	// Compute reverse map for named struct LightElementStatusType
	// insertion point per field
	stage.LightElementStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*LightElementStatusType)
	for lightelementstatustype := range stage.LightElementStatusTypes {
		_ = lightelementstatustype
		for _, _timesheetpropertytype := range lightelementstatustype.TimeInterval {
			stage.LightElementStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = lightelementstatustype
		}
	}
	stage.LightElementStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*LightElementStatusType)
	for lightelementstatustype := range stage.LightElementStatusTypes {
		_ = lightelementstatustype
		for _, _notepropertytype := range lightelementstatustype.Annotation {
			stage.LightElementStatusType_Annotation_reverseMap[_notepropertytype] = lightelementstatustype
		}
	}
	stage.LightElementStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*LightElementStatusType)
	for lightelementstatustype := range stage.LightElementStatusTypes {
		_ = lightelementstatustype
		for _, _organisationauthoritypropertytype := range lightelementstatustype.SpecialDateAuthority {
			stage.LightElementStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = lightelementstatustype
		}
	}
	stage.LightElementStatusType_Extension_reverseMap = make(map[*A_extension_11]*LightElementStatusType)
	for lightelementstatustype := range stage.LightElementStatusTypes {
		_ = lightelementstatustype
		for _, _a_extension_11 := range lightelementstatustype.Extension {
			stage.LightElementStatusType_Extension_reverseMap[_a_extension_11] = lightelementstatustype
		}
	}

	// Compute reverse map for named struct LightElementType
	// insertion point per field
	stage.LightElementType_Annotation_reverseMap = make(map[*NotePropertyType]*LightElementType)
	for lightelementtype := range stage.LightElementTypes {
		_ = lightelementtype
		for _, _notepropertytype := range lightelementtype.Annotation {
			stage.LightElementType_Annotation_reverseMap[_notepropertytype] = lightelementtype
		}
	}
	stage.LightElementType_Availability_reverseMap = make(map[*LightElementStatusPropertyType]*LightElementType)
	for lightelementtype := range stage.LightElementTypes {
		_ = lightelementtype
		for _, _lightelementstatuspropertytype := range lightelementtype.Availability {
			stage.LightElementType_Availability_reverseMap[_lightelementstatuspropertytype] = lightelementtype
		}
	}
	stage.LightElementType_Extension_reverseMap = make(map[*A_extension_10]*LightElementType)
	for lightelementtype := range stage.LightElementTypes {
		_ = lightelementtype
		for _, _a_extension_10 := range lightelementtype.Extension {
			stage.LightElementType_Extension_reverseMap[_a_extension_10] = lightelementtype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_4]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_4 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_4] = linguisticnotetype
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
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_5]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_5 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_5] = notetype
		}
	}

	// Compute reverse map for named struct ObstacleAreaTimeSliceType
	// insertion point per field
	stage.ObstacleAreaTimeSliceType_Obstacle_reverseMap = make(map[*VerticalStructurePropertyType]*ObstacleAreaTimeSliceType)
	for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
		_ = obstacleareatimeslicetype
		for _, _verticalstructurepropertytype := range obstacleareatimeslicetype.Obstacle {
			stage.ObstacleAreaTimeSliceType_Obstacle_reverseMap[_verticalstructurepropertytype] = obstacleareatimeslicetype
		}
	}
	stage.ObstacleAreaTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstacleAreaTimeSliceType)
	for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
		_ = obstacleareatimeslicetype
		for _, _notepropertytype := range obstacleareatimeslicetype.Annotation {
			stage.ObstacleAreaTimeSliceType_Annotation_reverseMap[_notepropertytype] = obstacleareatimeslicetype
		}
	}
	stage.ObstacleAreaTimeSliceType_Extension_reverseMap = make(map[*A_extension_6]*ObstacleAreaTimeSliceType)
	for obstacleareatimeslicetype := range stage.ObstacleAreaTimeSliceTypes {
		_ = obstacleareatimeslicetype
		for _, _a_extension_6 := range obstacleareatimeslicetype.Extension {
			stage.ObstacleAreaTimeSliceType_Extension_reverseMap[_a_extension_6] = obstacleareatimeslicetype
		}
	}

	// Compute reverse map for named struct ObstacleAreaType
	// insertion point per field
	stage.ObstacleAreaType_TimeSlice_reverseMap = make(map[*ObstacleAreaTimeSlicePropertyType]*ObstacleAreaType)
	for obstacleareatype := range stage.ObstacleAreaTypes {
		_ = obstacleareatype
		for _, _obstacleareatimeslicepropertytype := range obstacleareatype.TimeSlice {
			stage.ObstacleAreaType_TimeSlice_reverseMap[_obstacleareatimeslicepropertytype] = obstacleareatype
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

	// Compute reverse map for named struct SurfaceType
	// insertion point per field
	stage.SurfaceType_Annotation_reverseMap = make(map[*NotePropertyType]*SurfaceType)
	for surfacetype := range stage.SurfaceTypes {
		_ = surfacetype
		for _, _notepropertytype := range surfacetype.Annotation {
			stage.SurfaceType_Annotation_reverseMap[_notepropertytype] = surfacetype
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
	stage.TimesheetType_Extension_reverseMap = make(map[*A_extension_12]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _a_extension_12 := range timesheettype.Extension {
			stage.TimesheetType_Extension_reverseMap[_a_extension_12] = timesheettype
		}
	}

	// Compute reverse map for named struct VerticalStructureLightingStatusType
	// insertion point per field
	stage.VerticalStructureLightingStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*VerticalStructureLightingStatusType)
	for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
		_ = verticalstructurelightingstatustype
		for _, _timesheetpropertytype := range verticalstructurelightingstatustype.TimeInterval {
			stage.VerticalStructureLightingStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = verticalstructurelightingstatustype
		}
	}
	stage.VerticalStructureLightingStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*VerticalStructureLightingStatusType)
	for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
		_ = verticalstructurelightingstatustype
		for _, _notepropertytype := range verticalstructurelightingstatustype.Annotation {
			stage.VerticalStructureLightingStatusType_Annotation_reverseMap[_notepropertytype] = verticalstructurelightingstatustype
		}
	}
	stage.VerticalStructureLightingStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*VerticalStructureLightingStatusType)
	for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
		_ = verticalstructurelightingstatustype
		for _, _organisationauthoritypropertytype := range verticalstructurelightingstatustype.SpecialDateAuthority {
			stage.VerticalStructureLightingStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = verticalstructurelightingstatustype
		}
	}
	stage.VerticalStructureLightingStatusType_Extension_reverseMap = make(map[*A_extension_8]*VerticalStructureLightingStatusType)
	for verticalstructurelightingstatustype := range stage.VerticalStructureLightingStatusTypes {
		_ = verticalstructurelightingstatustype
		for _, _a_extension_8 := range verticalstructurelightingstatustype.Extension {
			stage.VerticalStructureLightingStatusType_Extension_reverseMap[_a_extension_8] = verticalstructurelightingstatustype
		}
	}

	// Compute reverse map for named struct VerticalStructurePartType
	// insertion point per field
	stage.VerticalStructurePartType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*VerticalStructurePartType)
	for verticalstructureparttype := range stage.VerticalStructurePartTypes {
		_ = verticalstructureparttype
		for _, _timesheetpropertytype := range verticalstructureparttype.TimeInterval {
			stage.VerticalStructurePartType_TimeInterval_reverseMap[_timesheetpropertytype] = verticalstructureparttype
		}
	}
	stage.VerticalStructurePartType_Annotation_reverseMap = make(map[*NotePropertyType]*VerticalStructurePartType)
	for verticalstructureparttype := range stage.VerticalStructurePartTypes {
		_ = verticalstructureparttype
		for _, _notepropertytype := range verticalstructureparttype.Annotation {
			stage.VerticalStructurePartType_Annotation_reverseMap[_notepropertytype] = verticalstructureparttype
		}
	}
	stage.VerticalStructurePartType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*VerticalStructurePartType)
	for verticalstructureparttype := range stage.VerticalStructurePartTypes {
		_ = verticalstructureparttype
		for _, _organisationauthoritypropertytype := range verticalstructureparttype.SpecialDateAuthority {
			stage.VerticalStructurePartType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = verticalstructureparttype
		}
	}
	stage.VerticalStructurePartType_Lighting_reverseMap = make(map[*LightElementPropertyType]*VerticalStructurePartType)
	for verticalstructureparttype := range stage.VerticalStructurePartTypes {
		_ = verticalstructureparttype
		for _, _lightelementpropertytype := range verticalstructureparttype.Lighting {
			stage.VerticalStructurePartType_Lighting_reverseMap[_lightelementpropertytype] = verticalstructureparttype
		}
	}
	stage.VerticalStructurePartType_Extension_reverseMap = make(map[*A_extension_9]*VerticalStructurePartType)
	for verticalstructureparttype := range stage.VerticalStructurePartTypes {
		_ = verticalstructureparttype
		for _, _a_extension_9 := range verticalstructureparttype.Extension {
			stage.VerticalStructurePartType_Extension_reverseMap[_a_extension_9] = verticalstructureparttype
		}
	}

	// Compute reverse map for named struct VerticalStructureTimeSliceType
	// insertion point per field
	stage.VerticalStructureTimeSliceType_Part_reverseMap = make(map[*VerticalStructurePartPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _verticalstructurepartpropertytype := range verticalstructuretimeslicetype.Part {
			stage.VerticalStructureTimeSliceType_Part_reverseMap[_verticalstructurepartpropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_HostedPassengerService_reverseMap = make(map[*PassengerServicePropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _passengerservicepropertytype := range verticalstructuretimeslicetype.HostedPassengerService {
			stage.VerticalStructureTimeSliceType_HostedPassengerService_reverseMap[_passengerservicepropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_SupportedGroundLight_reverseMap = make(map[*GroundLightSystemPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _groundlightsystempropertytype := range verticalstructuretimeslicetype.SupportedGroundLight {
			stage.VerticalStructureTimeSliceType_SupportedGroundLight_reverseMap[_groundlightsystempropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_HostedNavaidEquipment_reverseMap = make(map[*NavaidEquipmentPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _navaidequipmentpropertytype := range verticalstructuretimeslicetype.HostedNavaidEquipment {
			stage.VerticalStructureTimeSliceType_HostedNavaidEquipment_reverseMap[_navaidequipmentpropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_HostedSpecialNavStation_reverseMap = make(map[*SpecialNavigationStationPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _specialnavigationstationpropertytype := range verticalstructuretimeslicetype.HostedSpecialNavStation {
			stage.VerticalStructureTimeSliceType_HostedSpecialNavStation_reverseMap[_specialnavigationstationpropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_HostedUnit_reverseMap = make(map[*UnitPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _unitpropertytype := range verticalstructuretimeslicetype.HostedUnit {
			stage.VerticalStructureTimeSliceType_HostedUnit_reverseMap[_unitpropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_HostedOrganisation_reverseMap = make(map[*OrganisationAuthorityPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _organisationauthoritypropertytype := range verticalstructuretimeslicetype.HostedOrganisation {
			stage.VerticalStructureTimeSliceType_HostedOrganisation_reverseMap[_organisationauthoritypropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_SupportedService_reverseMap = make(map[*ServicePropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _servicepropertytype := range verticalstructuretimeslicetype.SupportedService {
			stage.VerticalStructureTimeSliceType_SupportedService_reverseMap[_servicepropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _notepropertytype := range verticalstructuretimeslicetype.Annotation {
			stage.VerticalStructureTimeSliceType_Annotation_reverseMap[_notepropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_LightingAvailability_reverseMap = make(map[*VerticalStructureLightingStatusPropertyType]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _verticalstructurelightingstatuspropertytype := range verticalstructuretimeslicetype.LightingAvailability {
			stage.VerticalStructureTimeSliceType_LightingAvailability_reverseMap[_verticalstructurelightingstatuspropertytype] = verticalstructuretimeslicetype
		}
	}
	stage.VerticalStructureTimeSliceType_Extension_reverseMap = make(map[*A_extension_7]*VerticalStructureTimeSliceType)
	for verticalstructuretimeslicetype := range stage.VerticalStructureTimeSliceTypes {
		_ = verticalstructuretimeslicetype
		for _, _a_extension_7 := range verticalstructuretimeslicetype.Extension {
			stage.VerticalStructureTimeSliceType_Extension_reverseMap[_a_extension_7] = verticalstructuretimeslicetype
		}
	}

	// Compute reverse map for named struct VerticalStructureType
	// insertion point per field
	stage.VerticalStructureType_TimeSlice_reverseMap = make(map[*VerticalStructureTimeSlicePropertyType]*VerticalStructureType)
	for verticalstructuretype := range stage.VerticalStructureTypes {
		_ = verticalstructuretype
		for _, _verticalstructuretimeslicepropertytype := range verticalstructuretype.TimeSlice {
			stage.VerticalStructureType_TimeSlice_reverseMap[_verticalstructuretimeslicepropertytype] = verticalstructuretype
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

	res = __gong__appendInstances(res, stage.A_extension_2s)

	res = __gong__appendInstances(res, stage.A_extension_3s)

	res = __gong__appendInstances(res, stage.A_extension_4s)

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AeronauticalGroundLightTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AeronauticalGroundLightTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AeronauticalGroundLightTypes)

	res = __gong__appendInstances(res, stage.AirportHeliportPropertyTypes)

	res = __gong__appendInstances(res, stage.CurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedCurvePropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedCurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointTypes)

	res = __gong__appendInstances(res, stage.ElevatedSurfacePropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedSurfaceTypes)

	res = __gong__appendInstances(res, stage.GroundLightSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.LightElementPropertyTypes)

	res = __gong__appendInstances(res, stage.LightElementStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.LightElementStatusTypes)

	res = __gong__appendInstances(res, stage.LightElementTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.MarkerBeaconPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidEquipmentPropertyTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.ObstacleAreaTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ObstacleAreaTimeSliceTypes)

	res = __gong__appendInstances(res, stage.ObstacleAreaTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.PassengerServicePropertyTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.RunwayDirectionPropertyTypes)

	res = __gong__appendInstances(res, stage.ServicePropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationPropertyTypes)

	res = __gong__appendInstances(res, stage.SurfacePropertyTypes)

	res = __gong__appendInstances(res, stage.SurfaceTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	res = __gong__appendInstances(res, stage.UnitPropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructureLightingStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructureLightingStatusTypes)

	res = __gong__appendInstances(res, stage.VerticalStructurePartPropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructurePartTypes)

	res = __gong__appendInstances(res, stage.VerticalStructurePropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructureTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructureTimeSliceTypes)

	res = __gong__appendInstances(res, stage.VerticalStructureTypes)

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

func (a_extension_2 *A_extension_2) GongCopy() GongstructIF {
	newInstance := new(A_extension_2)
	a_extension_2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_3 *A_extension_3) GongCopy() GongstructIF {
	newInstance := new(A_extension_3)
	a_extension_3.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_4 *A_extension_4) GongCopy() GongstructIF {
	newInstance := new(A_extension_4)
	a_extension_4.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_5 *A_extension_5) GongCopy() GongstructIF {
	newInstance := new(A_extension_5)
	a_extension_5.GongCopyBasicFields(newInstance)
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AeronauticalGroundLightTimeSlicePropertyType)
	aeronauticalgroundlighttimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AeronauticalGroundLightTimeSliceType)
	aeronauticalgroundlighttimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongCopy() GongstructIF {
	newInstance := new(AeronauticalGroundLightType)
	aeronauticalgroundlighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportHeliportPropertyType)
	airportheliportpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curvetype *CurveType) GongCopy() GongstructIF {
	newInstance := new(CurveType)
	curvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongCopy() GongstructIF {
	newInstance := new(ElevatedCurvePropertyType)
	elevatedcurvepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedcurvetype *ElevatedCurveType) GongCopy() GongstructIF {
	newInstance := new(ElevatedCurveType)
	elevatedcurvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointPropertyType)
	elevatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointtype *ElevatedPointType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointType)
	elevatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongCopy() GongstructIF {
	newInstance := new(ElevatedSurfacePropertyType)
	elevatedsurfacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongCopy() GongstructIF {
	newInstance := new(ElevatedSurfaceType)
	elevatedsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(GroundLightSystemPropertyType)
	groundlightsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightelementpropertytype *LightElementPropertyType) GongCopy() GongstructIF {
	newInstance := new(LightElementPropertyType)
	lightelementpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(LightElementStatusPropertyType)
	lightelementstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightelementstatustype *LightElementStatusType) GongCopy() GongstructIF {
	newInstance := new(LightElementStatusType)
	lightelementstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightelementtype *LightElementType) GongCopy() GongstructIF {
	newInstance := new(LightElementType)
	lightelementtype.GongCopyBasicFields(newInstance)
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

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongCopy() GongstructIF {
	newInstance := new(MarkerBeaconPropertyType)
	markerbeaconpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidEquipmentPropertyType)
	navaidequipmentpropertytype.GongCopyBasicFields(newInstance)
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

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAreaTimeSlicePropertyType)
	obstacleareatimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAreaTimeSliceType)
	obstacleareatimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleareatype *ObstacleAreaType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAreaType)
	obstacleareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityPropertyType)
	organisationauthoritypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongCopy() GongstructIF {
	newInstance := new(PassengerServicePropertyType)
	passengerservicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointtype *PointType) GongCopy() GongstructIF {
	newInstance := new(PointType)
	pointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayDirectionPropertyType)
	runwaydirectionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (servicepropertytype *ServicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ServicePropertyType)
	servicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationPropertyType)
	specialnavigationstationpropertytype.GongCopyBasicFields(newInstance)
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

func (unitpropertytype *UnitPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitPropertyType)
	unitpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructureLightingStatusPropertyType)
	verticalstructurelightingstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructureLightingStatusType)
	verticalstructurelightingstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructurePartPropertyType)
	verticalstructurepartpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructureparttype *VerticalStructurePartType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructurePartType)
	verticalstructureparttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructurePropertyType)
	verticalstructurepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructureTimeSlicePropertyType)
	verticalstructuretimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructureTimeSliceType)
	verticalstructuretimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructuretype *VerticalStructureType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructureType)
	verticalstructuretype.GongCopyBasicFields(newInstance)
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

func (a_extension_2 *A_extension_2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_2)
}

func (a_extension_3 *A_extension_3) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_3)
}

func (a_extension_4 *A_extension_4) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_4)
}

func (a_extension_5 *A_extension_5) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_5)
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aeronauticalgroundlighttimeslicepropertytype)
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aeronauticalgroundlighttimeslicetype)
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aeronauticalgroundlighttype)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportheliportpropertytype)
}

func (curvetype *CurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvetype)
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedcurvepropertytype)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedcurvetype)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointpropertytype)
}

func (elevatedpointtype *ElevatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointtype)
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedsurfacepropertytype)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedsurfacetype)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groundlightsystempropertytype)
}

func (lightelementpropertytype *LightElementPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightelementpropertytype)
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightelementstatuspropertytype)
}

func (lightelementstatustype *LightElementStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightelementstatustype)
}

func (lightelementtype *LightElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightelementtype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, markerbeaconpropertytype)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidequipmentpropertytype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleareatimeslicepropertytype)
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleareatimeslicetype)
}

func (obstacleareatype *ObstacleAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleareatype)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritypropertytype)
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, passengerservicepropertytype)
}

func (pointtype *PointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointtype)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaydirectionpropertytype)
}

func (servicepropertytype *ServicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, servicepropertytype)
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationpropertytype)
}

func (surfacepropertytype *SurfacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacepropertytype)
}

func (surfacetype *SurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacetype)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheetpropertytype)
}

func (timesheettype *TimesheetType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheettype)
}

func (unitpropertytype *UnitPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitpropertytype)
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurelightingstatuspropertytype)
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurelightingstatustype)
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurepartpropertytype)
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructureparttype)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurepropertytype)
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructuretimeslicepropertytype)
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructuretimeslicetype)
}

func (verticalstructuretype *VerticalStructureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructuretype)
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
		stage.AeronauticalGroundLightTimeSlicePropertyTypes,
		stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder,
		stage.AeronauticalGroundLightTimeSlicePropertyTypes_reference,
		&stage.AeronauticalGroundLightTimeSlicePropertyTypes_referenceOrder,
		stage.AeronauticalGroundLightTimeSlicePropertyTypes_instance,
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
		stage.AeronauticalGroundLightTimeSliceTypes,
		stage.AeronauticalGroundLightTimeSliceType_stagedOrder,
		stage.AeronauticalGroundLightTimeSliceTypes_reference,
		&stage.AeronauticalGroundLightTimeSliceTypes_referenceOrder,
		stage.AeronauticalGroundLightTimeSliceTypes_instance,
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
		stage.AeronauticalGroundLightTypes,
		stage.AeronauticalGroundLightType_stagedOrder,
		stage.AeronauticalGroundLightTypes_reference,
		&stage.AeronauticalGroundLightTypes_referenceOrder,
		stage.AeronauticalGroundLightTypes_instance,
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
		stage.ElevatedCurvePropertyTypes,
		stage.ElevatedCurvePropertyType_stagedOrder,
		stage.ElevatedCurvePropertyTypes_reference,
		&stage.ElevatedCurvePropertyTypes_referenceOrder,
		stage.ElevatedCurvePropertyTypes_instance,
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
		stage.ElevatedPointPropertyTypes,
		stage.ElevatedPointPropertyType_stagedOrder,
		stage.ElevatedPointPropertyTypes_reference,
		&stage.ElevatedPointPropertyTypes_referenceOrder,
		stage.ElevatedPointPropertyTypes_instance,
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
		stage.ElevatedSurfacePropertyTypes,
		stage.ElevatedSurfacePropertyType_stagedOrder,
		stage.ElevatedSurfacePropertyTypes_reference,
		&stage.ElevatedSurfacePropertyTypes_referenceOrder,
		stage.ElevatedSurfacePropertyTypes_instance,
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
		stage.GroundLightSystemPropertyTypes,
		stage.GroundLightSystemPropertyType_stagedOrder,
		stage.GroundLightSystemPropertyTypes_reference,
		&stage.GroundLightSystemPropertyTypes_referenceOrder,
		stage.GroundLightSystemPropertyTypes_instance,
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
		stage.LightElementPropertyTypes,
		stage.LightElementPropertyType_stagedOrder,
		stage.LightElementPropertyTypes_reference,
		&stage.LightElementPropertyTypes_referenceOrder,
		stage.LightElementPropertyTypes_instance,
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
		stage.LightElementStatusPropertyTypes,
		stage.LightElementStatusPropertyType_stagedOrder,
		stage.LightElementStatusPropertyTypes_reference,
		&stage.LightElementStatusPropertyTypes_referenceOrder,
		stage.LightElementStatusPropertyTypes_instance,
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
		stage.LightElementStatusTypes,
		stage.LightElementStatusType_stagedOrder,
		stage.LightElementStatusTypes_reference,
		&stage.LightElementStatusTypes_referenceOrder,
		stage.LightElementStatusTypes_instance,
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
		stage.LightElementTypes,
		stage.LightElementType_stagedOrder,
		stage.LightElementTypes_reference,
		&stage.LightElementTypes_referenceOrder,
		stage.LightElementTypes_instance,
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
		stage.MarkerBeaconPropertyTypes,
		stage.MarkerBeaconPropertyType_stagedOrder,
		stage.MarkerBeaconPropertyTypes_reference,
		&stage.MarkerBeaconPropertyTypes_referenceOrder,
		stage.MarkerBeaconPropertyTypes_instance,
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
		stage.NavaidEquipmentPropertyTypes,
		stage.NavaidEquipmentPropertyType_stagedOrder,
		stage.NavaidEquipmentPropertyTypes_reference,
		&stage.NavaidEquipmentPropertyTypes_referenceOrder,
		stage.NavaidEquipmentPropertyTypes_instance,
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
		stage.ObstacleAreaTimeSlicePropertyTypes,
		stage.ObstacleAreaTimeSlicePropertyType_stagedOrder,
		stage.ObstacleAreaTimeSlicePropertyTypes_reference,
		&stage.ObstacleAreaTimeSlicePropertyTypes_referenceOrder,
		stage.ObstacleAreaTimeSlicePropertyTypes_instance,
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
		stage.ObstacleAreaTimeSliceTypes,
		stage.ObstacleAreaTimeSliceType_stagedOrder,
		stage.ObstacleAreaTimeSliceTypes_reference,
		&stage.ObstacleAreaTimeSliceTypes_referenceOrder,
		stage.ObstacleAreaTimeSliceTypes_instance,
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
		stage.ObstacleAreaTypes,
		stage.ObstacleAreaType_stagedOrder,
		stage.ObstacleAreaTypes_reference,
		&stage.ObstacleAreaTypes_referenceOrder,
		stage.ObstacleAreaTypes_instance,
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
		stage.PassengerServicePropertyTypes,
		stage.PassengerServicePropertyType_stagedOrder,
		stage.PassengerServicePropertyTypes_reference,
		&stage.PassengerServicePropertyTypes_referenceOrder,
		stage.PassengerServicePropertyTypes_instance,
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
		stage.ServicePropertyTypes,
		stage.ServicePropertyType_stagedOrder,
		stage.ServicePropertyTypes_reference,
		&stage.ServicePropertyTypes_referenceOrder,
		stage.ServicePropertyTypes_instance,
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
		stage.SpecialNavigationStationPropertyTypes,
		stage.SpecialNavigationStationPropertyType_stagedOrder,
		stage.SpecialNavigationStationPropertyTypes_reference,
		&stage.SpecialNavigationStationPropertyTypes_referenceOrder,
		stage.SpecialNavigationStationPropertyTypes_instance,
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
		stage.UnitPropertyTypes,
		stage.UnitPropertyType_stagedOrder,
		stage.UnitPropertyTypes_reference,
		&stage.UnitPropertyTypes_referenceOrder,
		stage.UnitPropertyTypes_instance,
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
		stage.VerticalStructureLightingStatusPropertyTypes,
		stage.VerticalStructureLightingStatusPropertyType_stagedOrder,
		stage.VerticalStructureLightingStatusPropertyTypes_reference,
		&stage.VerticalStructureLightingStatusPropertyTypes_referenceOrder,
		stage.VerticalStructureLightingStatusPropertyTypes_instance,
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
		stage.VerticalStructureLightingStatusTypes,
		stage.VerticalStructureLightingStatusType_stagedOrder,
		stage.VerticalStructureLightingStatusTypes_reference,
		&stage.VerticalStructureLightingStatusTypes_referenceOrder,
		stage.VerticalStructureLightingStatusTypes_instance,
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
		stage.VerticalStructurePartPropertyTypes,
		stage.VerticalStructurePartPropertyType_stagedOrder,
		stage.VerticalStructurePartPropertyTypes_reference,
		&stage.VerticalStructurePartPropertyTypes_referenceOrder,
		stage.VerticalStructurePartPropertyTypes_instance,
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
		stage.VerticalStructurePartTypes,
		stage.VerticalStructurePartType_stagedOrder,
		stage.VerticalStructurePartTypes_reference,
		&stage.VerticalStructurePartTypes_referenceOrder,
		stage.VerticalStructurePartTypes_instance,
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
	computeCommitsForType(
		stage,
		stage.VerticalStructureTimeSlicePropertyTypes,
		stage.VerticalStructureTimeSlicePropertyType_stagedOrder,
		stage.VerticalStructureTimeSlicePropertyTypes_reference,
		&stage.VerticalStructureTimeSlicePropertyTypes_referenceOrder,
		stage.VerticalStructureTimeSlicePropertyTypes_instance,
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
		stage.VerticalStructureTimeSliceTypes,
		stage.VerticalStructureTimeSliceType_stagedOrder,
		stage.VerticalStructureTimeSliceTypes_reference,
		&stage.VerticalStructureTimeSliceTypes_referenceOrder,
		stage.VerticalStructureTimeSliceTypes_instance,
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
		stage.VerticalStructureTypes,
		stage.VerticalStructureType_stagedOrder,
		stage.VerticalStructureTypes_reference,
		&stage.VerticalStructureTypes_referenceOrder,
		stage.VerticalStructureTypes_instance,
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

	__gong__computeReferencePass1(stage, stage.A_extension_2s, &stage.A_extension_2s_reference, &stage.A_extension_2s_referenceOrder, &stage.A_extension_2s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_3s, &stage.A_extension_3s_reference, &stage.A_extension_3s_referenceOrder, &stage.A_extension_3s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_4s, &stage.A_extension_4s_reference, &stage.A_extension_4s_referenceOrder, &stage.A_extension_4s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AeronauticalGroundLightTimeSlicePropertyTypes, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_reference, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_referenceOrder, &stage.AeronauticalGroundLightTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AeronauticalGroundLightTimeSliceTypes, &stage.AeronauticalGroundLightTimeSliceTypes_reference, &stage.AeronauticalGroundLightTimeSliceTypes_referenceOrder, &stage.AeronauticalGroundLightTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AeronauticalGroundLightTypes, &stage.AeronauticalGroundLightTypes_reference, &stage.AeronauticalGroundLightTypes_referenceOrder, &stage.AeronauticalGroundLightTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_referenceOrder, &stage.AirportHeliportPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurveTypes, &stage.CurveTypes_reference, &stage.CurveTypes_referenceOrder, &stage.CurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedCurvePropertyTypes, &stage.ElevatedCurvePropertyTypes_reference, &stage.ElevatedCurvePropertyTypes_referenceOrder, &stage.ElevatedCurvePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_referenceOrder, &stage.ElevatedCurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointPropertyTypes, &stage.ElevatedPointPropertyTypes_reference, &stage.ElevatedPointPropertyTypes_referenceOrder, &stage.ElevatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointTypes, &stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_referenceOrder, &stage.ElevatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedSurfacePropertyTypes, &stage.ElevatedSurfacePropertyTypes_reference, &stage.ElevatedSurfacePropertyTypes_referenceOrder, &stage.ElevatedSurfacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_referenceOrder, &stage.ElevatedSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.GroundLightSystemPropertyTypes, &stage.GroundLightSystemPropertyTypes_reference, &stage.GroundLightSystemPropertyTypes_referenceOrder, &stage.GroundLightSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightElementPropertyTypes, &stage.LightElementPropertyTypes_reference, &stage.LightElementPropertyTypes_referenceOrder, &stage.LightElementPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightElementStatusPropertyTypes, &stage.LightElementStatusPropertyTypes_reference, &stage.LightElementStatusPropertyTypes_referenceOrder, &stage.LightElementStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightElementStatusTypes, &stage.LightElementStatusTypes_reference, &stage.LightElementStatusTypes_referenceOrder, &stage.LightElementStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightElementTypes, &stage.LightElementTypes_reference, &stage.LightElementTypes_referenceOrder, &stage.LightElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.MarkerBeaconPropertyTypes, &stage.MarkerBeaconPropertyTypes_reference, &stage.MarkerBeaconPropertyTypes_referenceOrder, &stage.MarkerBeaconPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidEquipmentPropertyTypes, &stage.NavaidEquipmentPropertyTypes_reference, &stage.NavaidEquipmentPropertyTypes_referenceOrder, &stage.NavaidEquipmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAreaTimeSlicePropertyTypes, &stage.ObstacleAreaTimeSlicePropertyTypes_reference, &stage.ObstacleAreaTimeSlicePropertyTypes_referenceOrder, &stage.ObstacleAreaTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAreaTimeSliceTypes, &stage.ObstacleAreaTimeSliceTypes_reference, &stage.ObstacleAreaTimeSliceTypes_referenceOrder, &stage.ObstacleAreaTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAreaTypes, &stage.ObstacleAreaTypes_reference, &stage.ObstacleAreaTypes_referenceOrder, &stage.ObstacleAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PassengerServicePropertyTypes, &stage.PassengerServicePropertyTypes_reference, &stage.PassengerServicePropertyTypes_referenceOrder, &stage.PassengerServicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayDirectionPropertyTypes, &stage.RunwayDirectionPropertyTypes_reference, &stage.RunwayDirectionPropertyTypes_referenceOrder, &stage.RunwayDirectionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ServicePropertyTypes, &stage.ServicePropertyTypes_reference, &stage.ServicePropertyTypes_referenceOrder, &stage.ServicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationPropertyTypes, &stage.SpecialNavigationStationPropertyTypes_reference, &stage.SpecialNavigationStationPropertyTypes_referenceOrder, &stage.SpecialNavigationStationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_referenceOrder, &stage.SurfacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfaceTypes, &stage.SurfaceTypes_reference, &stage.SurfaceTypes_referenceOrder, &stage.SurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitPropertyTypes, &stage.UnitPropertyTypes_reference, &stage.UnitPropertyTypes_referenceOrder, &stage.UnitPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructureLightingStatusPropertyTypes, &stage.VerticalStructureLightingStatusPropertyTypes_reference, &stage.VerticalStructureLightingStatusPropertyTypes_referenceOrder, &stage.VerticalStructureLightingStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructureLightingStatusTypes, &stage.VerticalStructureLightingStatusTypes_reference, &stage.VerticalStructureLightingStatusTypes_referenceOrder, &stage.VerticalStructureLightingStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructurePartPropertyTypes, &stage.VerticalStructurePartPropertyTypes_reference, &stage.VerticalStructurePartPropertyTypes_referenceOrder, &stage.VerticalStructurePartPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructurePartTypes, &stage.VerticalStructurePartTypes_reference, &stage.VerticalStructurePartTypes_referenceOrder, &stage.VerticalStructurePartTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructurePropertyTypes, &stage.VerticalStructurePropertyTypes_reference, &stage.VerticalStructurePropertyTypes_referenceOrder, &stage.VerticalStructurePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructureTimeSlicePropertyTypes, &stage.VerticalStructureTimeSlicePropertyTypes_reference, &stage.VerticalStructureTimeSlicePropertyTypes_referenceOrder, &stage.VerticalStructureTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructureTimeSliceTypes, &stage.VerticalStructureTimeSliceTypes_reference, &stage.VerticalStructureTimeSliceTypes_referenceOrder, &stage.VerticalStructureTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructureTypes, &stage.VerticalStructureTypes_reference, &stage.VerticalStructureTypes_referenceOrder, &stage.VerticalStructureTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.A_extensions, stage.A_extensions_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_1s, stage.A_extension_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_10s, stage.A_extension_10s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_11s, stage.A_extension_11s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_12s, stage.A_extension_12s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_2s, stage.A_extension_2s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_3s, stage.A_extension_3s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_4s, stage.A_extension_4s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurveTypes, stage.CurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedCurveTypes, stage.ElevatedCurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointTypes, stage.ElevatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightElementPropertyTypes, stage.LightElementPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightElementStatusTypes, stage.LightElementStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightElementTypes, stage.LightElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAreaTypes, stage.ObstacleAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ServicePropertyTypes, stage.ServicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfacePropertyTypes, stage.SurfacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfaceTypes, stage.SurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitPropertyTypes, stage.UnitPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructurePartTypes, stage.VerticalStructurePartTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructureTypes, stage.VerticalStructureTypes_reference, stage)

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

func (a_extension_2 *A_extension_2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_2_stagedOrder, stage.A_extension_2s_referenceOrder, a_extension_2, "A_extension_2")
}

func (a_extension_3 *A_extension_3) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_3_stagedOrder, stage.A_extension_3s_referenceOrder, a_extension_3, "A_extension_3")
}

func (a_extension_4 *A_extension_4) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_4_stagedOrder, stage.A_extension_4s_referenceOrder, a_extension_4, "A_extension_4")
}

func (a_extension_5 *A_extension_5) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_5_stagedOrder, stage.A_extension_5s_referenceOrder, a_extension_5, "A_extension_5")
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder, stage.AeronauticalGroundLightTimeSlicePropertyTypes_referenceOrder, aeronauticalgroundlighttimeslicepropertytype, "AeronauticalGroundLightTimeSlicePropertyType")
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AeronauticalGroundLightTimeSliceType_stagedOrder, stage.AeronauticalGroundLightTimeSliceTypes_referenceOrder, aeronauticalgroundlighttimeslicetype, "AeronauticalGroundLightTimeSliceType")
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AeronauticalGroundLightType_stagedOrder, stage.AeronauticalGroundLightTypes_referenceOrder, aeronauticalgroundlighttype, "AeronauticalGroundLightType")
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyTypes_referenceOrder, airportheliportpropertytype, "AirportHeliportPropertyType")
}

func (curvetype *CurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurveType_stagedOrder, stage.CurveTypes_referenceOrder, curvetype, "CurveType")
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedCurvePropertyType_stagedOrder, stage.ElevatedCurvePropertyTypes_referenceOrder, elevatedcurvepropertytype, "ElevatedCurvePropertyType")
}

func (elevatedcurvetype *ElevatedCurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveTypes_referenceOrder, elevatedcurvetype, "ElevatedCurveType")
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointPropertyType_stagedOrder, stage.ElevatedPointPropertyTypes_referenceOrder, elevatedpointpropertytype, "ElevatedPointPropertyType")
}

func (elevatedpointtype *ElevatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointType_stagedOrder, stage.ElevatedPointTypes_referenceOrder, elevatedpointtype, "ElevatedPointType")
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedSurfacePropertyType_stagedOrder, stage.ElevatedSurfacePropertyTypes_referenceOrder, elevatedsurfacepropertytype, "ElevatedSurfacePropertyType")
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceTypes_referenceOrder, elevatedsurfacetype, "ElevatedSurfaceType")
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroundLightSystemPropertyType_stagedOrder, stage.GroundLightSystemPropertyTypes_referenceOrder, groundlightsystempropertytype, "GroundLightSystemPropertyType")
}

func (lightelementpropertytype *LightElementPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightElementPropertyType_stagedOrder, stage.LightElementPropertyTypes_referenceOrder, lightelementpropertytype, "LightElementPropertyType")
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightElementStatusPropertyType_stagedOrder, stage.LightElementStatusPropertyTypes_referenceOrder, lightelementstatuspropertytype, "LightElementStatusPropertyType")
}

func (lightelementstatustype *LightElementStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightElementStatusType_stagedOrder, stage.LightElementStatusTypes_referenceOrder, lightelementstatustype, "LightElementStatusType")
}

func (lightelementtype *LightElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightElementType_stagedOrder, stage.LightElementTypes_referenceOrder, lightelementtype, "LightElementType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MarkerBeaconPropertyType_stagedOrder, stage.MarkerBeaconPropertyTypes_referenceOrder, markerbeaconpropertytype, "MarkerBeaconPropertyType")
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidEquipmentPropertyType_stagedOrder, stage.NavaidEquipmentPropertyTypes_referenceOrder, navaidequipmentpropertytype, "NavaidEquipmentPropertyType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAreaTimeSlicePropertyType_stagedOrder, stage.ObstacleAreaTimeSlicePropertyTypes_referenceOrder, obstacleareatimeslicepropertytype, "ObstacleAreaTimeSlicePropertyType")
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAreaTimeSliceType_stagedOrder, stage.ObstacleAreaTimeSliceTypes_referenceOrder, obstacleareatimeslicetype, "ObstacleAreaTimeSliceType")
}

func (obstacleareatype *ObstacleAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAreaType_stagedOrder, stage.ObstacleAreaTypes_referenceOrder, obstacleareatype, "ObstacleAreaType")
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyTypes_referenceOrder, organisationauthoritypropertytype, "OrganisationAuthorityPropertyType")
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PassengerServicePropertyType_stagedOrder, stage.PassengerServicePropertyTypes_referenceOrder, passengerservicepropertytype, "PassengerServicePropertyType")
}

func (pointtype *PointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointType_stagedOrder, stage.PointTypes_referenceOrder, pointtype, "PointType")
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayDirectionPropertyType_stagedOrder, stage.RunwayDirectionPropertyTypes_referenceOrder, runwaydirectionpropertytype, "RunwayDirectionPropertyType")
}

func (servicepropertytype *ServicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ServicePropertyType_stagedOrder, stage.ServicePropertyTypes_referenceOrder, servicepropertytype, "ServicePropertyType")
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationPropertyType_stagedOrder, stage.SpecialNavigationStationPropertyTypes_referenceOrder, specialnavigationstationpropertytype, "SpecialNavigationStationPropertyType")
}

func (surfacepropertytype *SurfacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfacePropertyType_stagedOrder, stage.SurfacePropertyTypes_referenceOrder, surfacepropertytype, "SurfacePropertyType")
}

func (surfacetype *SurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfaceType_stagedOrder, stage.SurfaceTypes_referenceOrder, surfacetype, "SurfaceType")
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyTypes_referenceOrder, timesheetpropertytype, "TimesheetPropertyType")
}

func (timesheettype *TimesheetType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetType_stagedOrder, stage.TimesheetTypes_referenceOrder, timesheettype, "TimesheetType")
}

func (unitpropertytype *UnitPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitPropertyType_stagedOrder, stage.UnitPropertyTypes_referenceOrder, unitpropertytype, "UnitPropertyType")
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructureLightingStatusPropertyType_stagedOrder, stage.VerticalStructureLightingStatusPropertyTypes_referenceOrder, verticalstructurelightingstatuspropertytype, "VerticalStructureLightingStatusPropertyType")
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructureLightingStatusType_stagedOrder, stage.VerticalStructureLightingStatusTypes_referenceOrder, verticalstructurelightingstatustype, "VerticalStructureLightingStatusType")
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructurePartPropertyType_stagedOrder, stage.VerticalStructurePartPropertyTypes_referenceOrder, verticalstructurepartpropertytype, "VerticalStructurePartPropertyType")
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructurePartType_stagedOrder, stage.VerticalStructurePartTypes_referenceOrder, verticalstructureparttype, "VerticalStructurePartType")
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructurePropertyType_stagedOrder, stage.VerticalStructurePropertyTypes_referenceOrder, verticalstructurepropertytype, "VerticalStructurePropertyType")
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructureTimeSlicePropertyType_stagedOrder, stage.VerticalStructureTimeSlicePropertyTypes_referenceOrder, verticalstructuretimeslicepropertytype, "VerticalStructureTimeSlicePropertyType")
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructureTimeSliceType_stagedOrder, stage.VerticalStructureTimeSliceTypes_referenceOrder, verticalstructuretimeslicetype, "VerticalStructureTimeSliceType")
}

func (verticalstructuretype *VerticalStructureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructureType_stagedOrder, stage.VerticalStructureTypes_referenceOrder, verticalstructuretype, "VerticalStructureType")
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

func (a_extension_2 *A_extension_2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_2, a_extension_2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_2 *A_extension_2) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_2.GongGetIdentifier(stage)
}

func (a_extension_3 *A_extension_3) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_3, a_extension_3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_3 *A_extension_3) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_3.GongGetIdentifier(stage)
}

func (a_extension_4 *A_extension_4) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_4, a_extension_4.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_4 *A_extension_4) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_4.GongGetIdentifier(stage)
}

func (a_extension_5 *A_extension_5) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_5, a_extension_5.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_5 *A_extension_5) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_5.GongGetIdentifier(stage)
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aeronauticalgroundlighttimeslicepropertytype, aeronauticalgroundlighttimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage)
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aeronauticalgroundlighttimeslicetype, aeronauticalgroundlighttimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return aeronauticalgroundlighttimeslicetype.GongGetIdentifier(stage)
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aeronauticalgroundlighttype, aeronauticalgroundlighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongGetReferenceIdentifier(stage *Stage) string {
	return aeronauticalgroundlighttype.GongGetIdentifier(stage)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportheliportpropertytype, airportheliportpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportheliportpropertytype.GongGetIdentifier(stage)
}

func (curvetype *CurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curvetype, curvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curvetype *CurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return curvetype.GongGetIdentifier(stage)
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedcurvepropertytype, elevatedcurvepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedcurvepropertytype.GongGetIdentifier(stage)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedcurvetype, elevatedcurvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedcurvetype *ElevatedCurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedcurvetype.GongGetIdentifier(stage)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointpropertytype, elevatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointpropertytype.GongGetIdentifier(stage)
}

func (elevatedpointtype *ElevatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointtype, elevatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointtype *ElevatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointtype.GongGetIdentifier(stage)
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedsurfacepropertytype, elevatedsurfacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedsurfacepropertytype.GongGetIdentifier(stage)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedsurfacetype, elevatedsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedsurfacetype *ElevatedSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedsurfacetype.GongGetIdentifier(stage)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groundlightsystempropertytype, groundlightsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return groundlightsystempropertytype.GongGetIdentifier(stage)
}

func (lightelementpropertytype *LightElementPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightelementpropertytype, lightelementpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightelementpropertytype *LightElementPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightelementpropertytype.GongGetIdentifier(stage)
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightelementstatuspropertytype, lightelementstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightelementstatuspropertytype.GongGetIdentifier(stage)
}

func (lightelementstatustype *LightElementStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightelementstatustype, lightelementstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightelementstatustype *LightElementStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightelementstatustype.GongGetIdentifier(stage)
}

func (lightelementtype *LightElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightelementtype, lightelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightelementtype *LightElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightelementtype.GongGetIdentifier(stage)
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

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(markerbeaconpropertytype, markerbeaconpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return markerbeaconpropertytype.GongGetIdentifier(stage)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidequipmentpropertytype, navaidequipmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidequipmentpropertytype.GongGetIdentifier(stage)
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

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleareatimeslicepropertytype, obstacleareatimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleareatimeslicepropertytype.GongGetIdentifier(stage)
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleareatimeslicetype, obstacleareatimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleareatimeslicetype.GongGetIdentifier(stage)
}

func (obstacleareatype *ObstacleAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleareatype, obstacleareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleareatype *ObstacleAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleareatype.GongGetIdentifier(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritypropertytype, organisationauthoritypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritypropertytype.GongGetIdentifier(stage)
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(passengerservicepropertytype, passengerservicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (passengerservicepropertytype *PassengerServicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return passengerservicepropertytype.GongGetIdentifier(stage)
}

func (pointtype *PointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointtype, pointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointtype *PointType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointtype.GongGetIdentifier(stage)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaydirectionpropertytype, runwaydirectionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaydirectionpropertytype.GongGetIdentifier(stage)
}

func (servicepropertytype *ServicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(servicepropertytype, servicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (servicepropertytype *ServicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return servicepropertytype.GongGetIdentifier(stage)
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationpropertytype, specialnavigationstationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationpropertytype.GongGetIdentifier(stage)
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

func (unitpropertytype *UnitPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitpropertytype, unitpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitpropertytype *UnitPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitpropertytype.GongGetIdentifier(stage)
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurelightingstatuspropertytype, verticalstructurelightingstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage)
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurelightingstatustype, verticalstructurelightingstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurelightingstatustype.GongGetIdentifier(stage)
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurepartpropertytype, verticalstructurepartpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurepartpropertytype.GongGetIdentifier(stage)
}

func (verticalstructureparttype *VerticalStructurePartType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructureparttype, verticalstructureparttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructureparttype *VerticalStructurePartType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructureparttype.GongGetIdentifier(stage)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurepropertytype, verticalstructurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurepropertytype.GongGetIdentifier(stage)
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructuretimeslicepropertytype, verticalstructuretimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructuretimeslicepropertytype.GongGetIdentifier(stage)
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructuretimeslicetype, verticalstructuretimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructuretimeslicetype.GongGetIdentifier(stage)
}

func (verticalstructuretype *VerticalStructureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructuretype, verticalstructuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructuretype *VerticalStructureType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructuretype.GongGetIdentifier(stage)
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

func (a_extension_2 *A_extension_2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_2.GongGetIdentifier(stage), "A_extension_2", a_extension_2.Name)
}

func (a_extension_3 *A_extension_3) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_3.GongGetIdentifier(stage), "A_extension_3", a_extension_3.Name)
}

func (a_extension_4 *A_extension_4) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_4.GongGetIdentifier(stage), "A_extension_4", a_extension_4.Name)
}

func (a_extension_5 *A_extension_5) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_5.GongGetIdentifier(stage), "A_extension_5", a_extension_5.Name)
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage), "AeronauticalGroundLightTimeSlicePropertyType", aeronauticalgroundlighttimeslicepropertytype.Name)
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aeronauticalgroundlighttimeslicetype.GongGetIdentifier(stage), "AeronauticalGroundLightTimeSliceType", aeronauticalgroundlighttimeslicetype.Name)
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aeronauticalgroundlighttype.GongGetIdentifier(stage), "AeronauticalGroundLightType", aeronauticalgroundlighttype.Name)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportheliportpropertytype.GongGetIdentifier(stage), "AirportHeliportPropertyType", airportheliportpropertytype.Name)
}

func (curvetype *CurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvetype.GongGetIdentifier(stage), "CurveType", curvetype.Name)
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedcurvepropertytype.GongGetIdentifier(stage), "ElevatedCurvePropertyType", elevatedcurvepropertytype.Name)
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedcurvetype.GongGetIdentifier(stage), "ElevatedCurveType", elevatedcurvetype.Name)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointpropertytype.GongGetIdentifier(stage), "ElevatedPointPropertyType", elevatedpointpropertytype.Name)
}

func (elevatedpointtype *ElevatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointtype.GongGetIdentifier(stage), "ElevatedPointType", elevatedpointtype.Name)
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedsurfacepropertytype.GongGetIdentifier(stage), "ElevatedSurfacePropertyType", elevatedsurfacepropertytype.Name)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedsurfacetype.GongGetIdentifier(stage), "ElevatedSurfaceType", elevatedsurfacetype.Name)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groundlightsystempropertytype.GongGetIdentifier(stage), "GroundLightSystemPropertyType", groundlightsystempropertytype.Name)
}

func (lightelementpropertytype *LightElementPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightelementpropertytype.GongGetIdentifier(stage), "LightElementPropertyType", lightelementpropertytype.Name)
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightelementstatuspropertytype.GongGetIdentifier(stage), "LightElementStatusPropertyType", lightelementstatuspropertytype.Name)
}

func (lightelementstatustype *LightElementStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightelementstatustype.GongGetIdentifier(stage), "LightElementStatusType", lightelementstatustype.Name)
}

func (lightelementtype *LightElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightelementtype.GongGetIdentifier(stage), "LightElementType", lightelementtype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(markerbeaconpropertytype.GongGetIdentifier(stage), "MarkerBeaconPropertyType", markerbeaconpropertytype.Name)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidequipmentpropertytype.GongGetIdentifier(stage), "NavaidEquipmentPropertyType", navaidequipmentpropertytype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleareatimeslicepropertytype.GongGetIdentifier(stage), "ObstacleAreaTimeSlicePropertyType", obstacleareatimeslicepropertytype.Name)
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleareatimeslicetype.GongGetIdentifier(stage), "ObstacleAreaTimeSliceType", obstacleareatimeslicetype.Name)
}

func (obstacleareatype *ObstacleAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleareatype.GongGetIdentifier(stage), "ObstacleAreaType", obstacleareatype.Name)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritypropertytype.GongGetIdentifier(stage), "OrganisationAuthorityPropertyType", organisationauthoritypropertytype.Name)
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(passengerservicepropertytype.GongGetIdentifier(stage), "PassengerServicePropertyType", passengerservicepropertytype.Name)
}

func (pointtype *PointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointtype.GongGetIdentifier(stage), "PointType", pointtype.Name)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaydirectionpropertytype.GongGetIdentifier(stage), "RunwayDirectionPropertyType", runwaydirectionpropertytype.Name)
}

func (servicepropertytype *ServicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(servicepropertytype.GongGetIdentifier(stage), "ServicePropertyType", servicepropertytype.Name)
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationpropertytype.GongGetIdentifier(stage), "SpecialNavigationStationPropertyType", specialnavigationstationpropertytype.Name)
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacepropertytype.GongGetIdentifier(stage), "SurfacePropertyType", surfacepropertytype.Name)
}

func (surfacetype *SurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacetype.GongGetIdentifier(stage), "SurfaceType", surfacetype.Name)
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheetpropertytype.GongGetIdentifier(stage), "TimesheetPropertyType", timesheetpropertytype.Name)
}

func (timesheettype *TimesheetType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheettype.GongGetIdentifier(stage), "TimesheetType", timesheettype.Name)
}

func (unitpropertytype *UnitPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitpropertytype.GongGetIdentifier(stage), "UnitPropertyType", unitpropertytype.Name)
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage), "VerticalStructureLightingStatusPropertyType", verticalstructurelightingstatuspropertytype.Name)
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurelightingstatustype.GongGetIdentifier(stage), "VerticalStructureLightingStatusType", verticalstructurelightingstatustype.Name)
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurepartpropertytype.GongGetIdentifier(stage), "VerticalStructurePartPropertyType", verticalstructurepartpropertytype.Name)
}

func (verticalstructureparttype *VerticalStructurePartType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructureparttype.GongGetIdentifier(stage), "VerticalStructurePartType", verticalstructureparttype.Name)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurepropertytype.GongGetIdentifier(stage), "VerticalStructurePropertyType", verticalstructurepropertytype.Name)
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructuretimeslicepropertytype.GongGetIdentifier(stage), "VerticalStructureTimeSlicePropertyType", verticalstructuretimeslicepropertytype.Name)
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructuretimeslicetype.GongGetIdentifier(stage), "VerticalStructureTimeSliceType", verticalstructuretimeslicetype.Name)
}

func (verticalstructuretype *VerticalStructureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructuretype.GongGetIdentifier(stage), "VerticalStructureType", verticalstructuretype.Name)
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

func (a_extension_2 *A_extension_2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_2.GongGetReferenceIdentifier(stage))
}

func (a_extension_3 *A_extension_3) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_3.GongGetReferenceIdentifier(stage))
}

func (a_extension_4 *A_extension_4) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_4.GongGetReferenceIdentifier(stage))
}

func (a_extension_5 *A_extension_5) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_5.GongGetReferenceIdentifier(stage))
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

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aeronauticalgroundlighttimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aeronauticalgroundlighttimeslicetype.GongGetReferenceIdentifier(stage))
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aeronauticalgroundlighttype.GongGetReferenceIdentifier(stage))
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportheliportpropertytype.GongGetReferenceIdentifier(stage))
}

func (curvetype *CurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedcurvepropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedcurvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointtype *ElevatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointtype.GongGetReferenceIdentifier(stage))
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedsurfacepropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedsurfacetype.GongGetReferenceIdentifier(stage))
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groundlightsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (lightelementpropertytype *LightElementPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightelementpropertytype.GongGetReferenceIdentifier(stage))
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightelementstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (lightelementstatustype *LightElementStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightelementstatustype.GongGetReferenceIdentifier(stage))
}

func (lightelementtype *LightElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightelementtype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(markerbeaconpropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidequipmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleareatimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleareatimeslicetype.GongGetReferenceIdentifier(stage))
}

func (obstacleareatype *ObstacleAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleareatype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritypropertytype.GongGetReferenceIdentifier(stage))
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(passengerservicepropertytype.GongGetReferenceIdentifier(stage))
}

func (pointtype *PointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointtype.GongGetReferenceIdentifier(stage))
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaydirectionpropertytype.GongGetReferenceIdentifier(stage))
}

func (servicepropertytype *ServicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(servicepropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationpropertytype.GongGetReferenceIdentifier(stage))
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacepropertytype.GongGetReferenceIdentifier(stage))
}

func (surfacetype *SurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacetype.GongGetReferenceIdentifier(stage))
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheetpropertytype.GongGetReferenceIdentifier(stage))
}

func (timesheettype *TimesheetType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheettype.GongGetReferenceIdentifier(stage))
}

func (unitpropertytype *UnitPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitpropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurelightingstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurelightingstatustype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurepartpropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructureparttype *VerticalStructurePartType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructureparttype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurepropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructuretimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructuretimeslicetype.GongGetReferenceIdentifier(stage))
}

func (verticalstructuretype *VerticalStructureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructuretype.GongGetReferenceIdentifier(stage))
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
