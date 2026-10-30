// generated code - do not edit
package aixmfeatures

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (a_extension *A_extension) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extensions[a_extension]
	return ok
}

func (a_extension_1 *A_extension_1) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_1s[a_extension_1]
	return ok
}

func (a_extension_10 *A_extension_10) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_10s[a_extension_10]
	return ok
}

func (a_extension_11 *A_extension_11) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_11s[a_extension_11]
	return ok
}

func (a_extension_12 *A_extension_12) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_12s[a_extension_12]
	return ok
}

func (a_extension_2 *A_extension_2) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_2s[a_extension_2]
	return ok
}

func (a_extension_3 *A_extension_3) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_3s[a_extension_3]
	return ok
}

func (a_extension_4 *A_extension_4) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_4s[a_extension_4]
	return ok
}

func (a_extension_5 *A_extension_5) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_5s[a_extension_5]
	return ok
}

func (a_extension_6 *A_extension_6) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_6s[a_extension_6]
	return ok
}

func (a_extension_7 *A_extension_7) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_7s[a_extension_7]
	return ok
}

func (a_extension_8 *A_extension_8) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_8s[a_extension_8]
	return ok
}

func (a_extension_9 *A_extension_9) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_9s[a_extension_9]
	return ok
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractPropertiesWithScheduleTypes[abstractpropertieswithscheduletype]
	return ok
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AeronauticalGroundLightTimeSlicePropertyTypes[aeronauticalgroundlighttimeslicepropertytype]
	return ok
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AeronauticalGroundLightTimeSliceTypes[aeronauticalgroundlighttimeslicetype]
	return ok
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AeronauticalGroundLightTypes[aeronauticalgroundlighttype]
	return ok
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirportHeliportPropertyTypes[airportheliportpropertytype]
	return ok
}

func (curvetype *CurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurveTypes[curvetype]
	return ok
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedCurvePropertyTypes[elevatedcurvepropertytype]
	return ok
}

func (elevatedcurvetype *ElevatedCurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedCurveTypes[elevatedcurvetype]
	return ok
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedPointPropertyTypes[elevatedpointpropertytype]
	return ok
}

func (elevatedpointtype *ElevatedPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedPointTypes[elevatedpointtype]
	return ok
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedSurfacePropertyTypes[elevatedsurfacepropertytype]
	return ok
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedSurfaceTypes[elevatedsurfacetype]
	return ok
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroundLightSystemPropertyTypes[groundlightsystempropertytype]
	return ok
}

func (lightelementpropertytype *LightElementPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LightElementPropertyTypes[lightelementpropertytype]
	return ok
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LightElementStatusPropertyTypes[lightelementstatuspropertytype]
	return ok
}

func (lightelementstatustype *LightElementStatusType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LightElementStatusTypes[lightelementstatustype]
	return ok
}

func (lightelementtype *LightElementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LightElementTypes[lightelementtype]
	return ok
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LinguisticNotePropertyTypes[linguisticnotepropertytype]
	return ok
}

func (linguisticnotetype *LinguisticNoteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LinguisticNoteTypes[linguisticnotetype]
	return ok
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MarkerBeaconPropertyTypes[markerbeaconpropertytype]
	return ok
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NavaidEquipmentPropertyTypes[navaidequipmentpropertytype]
	return ok
}

func (notepropertytype *NotePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NotePropertyTypes[notepropertytype]
	return ok
}

func (notetype *NoteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTypes[notetype]
	return ok
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ObstacleAreaTimeSlicePropertyTypes[obstacleareatimeslicepropertytype]
	return ok
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ObstacleAreaTimeSliceTypes[obstacleareatimeslicetype]
	return ok
}

func (obstacleareatype *ObstacleAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ObstacleAreaTypes[obstacleareatype]
	return ok
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OrganisationAuthorityPropertyTypes[organisationauthoritypropertytype]
	return ok
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PassengerServicePropertyTypes[passengerservicepropertytype]
	return ok
}

func (pointtype *PointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointTypes[pointtype]
	return ok
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RunwayDirectionPropertyTypes[runwaydirectionpropertytype]
	return ok
}

func (servicepropertytype *ServicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ServicePropertyTypes[servicepropertytype]
	return ok
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SpecialNavigationStationPropertyTypes[specialnavigationstationpropertytype]
	return ok
}

func (surfacepropertytype *SurfacePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SurfacePropertyTypes[surfacepropertytype]
	return ok
}

func (surfacetype *SurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SurfaceTypes[surfacetype]
	return ok
}

func (timesheetpropertytype *TimesheetPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimesheetPropertyTypes[timesheetpropertytype]
	return ok
}

func (timesheettype *TimesheetType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimesheetTypes[timesheettype]
	return ok
}

func (unitpropertytype *UnitPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.UnitPropertyTypes[unitpropertytype]
	return ok
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructureLightingStatusPropertyTypes[verticalstructurelightingstatuspropertytype]
	return ok
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructureLightingStatusTypes[verticalstructurelightingstatustype]
	return ok
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructurePartPropertyTypes[verticalstructurepartpropertytype]
	return ok
}

func (verticalstructureparttype *VerticalStructurePartType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructurePartTypes[verticalstructureparttype]
	return ok
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructurePropertyTypes[verticalstructurepropertytype]
	return ok
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructureTimeSlicePropertyTypes[verticalstructuretimeslicepropertytype]
	return ok
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructureTimeSliceTypes[verticalstructuretimeslicetype]
	return ok
}

func (verticalstructuretype *VerticalStructureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalStructureTypes[verticalstructuretype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (a_extension *A_extension) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension) {
		return
	}

	a_extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_1 *A_extension_1) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_1) {
		return
	}

	a_extension_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_10 *A_extension_10) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_10) {
		return
	}

	a_extension_10.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_11 *A_extension_11) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_11) {
		return
	}

	a_extension_11.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_12 *A_extension_12) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_12) {
		return
	}

	a_extension_12.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_2 *A_extension_2) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_2) {
		return
	}

	a_extension_2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_3 *A_extension_3) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_3) {
		return
	}

	a_extension_3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_4 *A_extension_4) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_4) {
		return
	}

	a_extension_4.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_5 *A_extension_5) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_5) {
		return
	}

	a_extension_5.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_6 *A_extension_6) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_6) {
		return
	}

	a_extension_6.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_7 *A_extension_7) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_7) {
		return
	}

	a_extension_7.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_8 *A_extension_8) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_8) {
		return
	}

	a_extension_8.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_9 *A_extension_9) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_9) {
		return
	}

	a_extension_9.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractpropertieswithscheduletype) {
		return
	}

	abstractpropertieswithscheduletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aeronauticalgroundlighttimeslicepropertytype) {
		return
	}

	aeronauticalgroundlighttimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
		stage.StageBranch(aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aeronauticalgroundlighttimeslicetype) {
		return
	}

	aeronauticalgroundlighttimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
		stage.StageBranch(aeronauticalgroundlighttimeslicetype.StructureBeacon)
	}
	if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
		stage.StageBranch(aeronauticalgroundlighttimeslicetype.AerodromeBeacon)
	}
	if aeronauticalgroundlighttimeslicetype.Location != nil {
		stage.StageBranch(aeronauticalgroundlighttimeslicetype.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aeronauticalgroundlighttimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_3 := range aeronauticalgroundlighttimeslicetype.Extension {
		stage.StageBranch(_a_extension_3)
	}

}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aeronauticalgroundlighttype) {
		return
	}

	aeronauticalgroundlighttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _aeronauticalgroundlighttimeslicepropertytype := range aeronauticalgroundlighttype.TimeSlice {
		stage.StageBranch(_aeronauticalgroundlighttimeslicepropertytype)
	}

}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airportheliportpropertytype) {
		return
	}

	airportheliportpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (curvetype *CurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curvetype) {
		return
	}

	curvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range curvetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}

}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedcurvepropertytype) {
		return
	}

	elevatedcurvepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedcurvepropertytype.ElevatedCurve != nil {
		stage.StageBranch(elevatedcurvepropertytype.ElevatedCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedcurvetype *ElevatedCurveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedcurvetype) {
		return
	}

	elevatedcurvetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedcurvetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension := range elevatedcurvetype.Extension {
		stage.StageBranch(_a_extension)
	}

}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedpointpropertytype) {
		return
	}

	elevatedpointpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedpointpropertytype.ElevatedPoint != nil {
		stage.StageBranch(elevatedpointpropertytype.ElevatedPoint)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedpointtype *ElevatedPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedpointtype) {
		return
	}

	elevatedpointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedpointtype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_1 := range elevatedpointtype.Extension {
		stage.StageBranch(_a_extension_1)
	}

}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedsurfacepropertytype) {
		return
	}

	elevatedsurfacepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedsurfacepropertytype.ElevatedSurface != nil {
		stage.StageBranch(elevatedsurfacepropertytype.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedsurfacetype *ElevatedSurfaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elevatedsurfacetype) {
		return
	}

	elevatedsurfacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedsurfacetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_2 := range elevatedsurfacetype.Extension {
		stage.StageBranch(_a_extension_2)
	}

}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(groundlightsystempropertytype) {
		return
	}

	groundlightsystempropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementpropertytype *LightElementPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lightelementpropertytype) {
		return
	}

	lightelementpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementpropertytype.LightElement != nil {
		stage.StageBranch(lightelementpropertytype.LightElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lightelementstatuspropertytype) {
		return
	}

	lightelementstatuspropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementstatuspropertytype.LightElementStatus != nil {
		stage.StageBranch(lightelementstatuspropertytype.LightElementStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementstatustype *LightElementStatusType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lightelementstatustype) {
		return
	}

	lightelementstatustype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range lightelementstatustype.TimeInterval {
		stage.StageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range lightelementstatustype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range lightelementstatustype.SpecialDateAuthority {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _a_extension_11 := range lightelementstatustype.Extension {
		stage.StageBranch(_a_extension_11)
	}

}

func (lightelementtype *LightElementType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lightelementtype) {
		return
	}

	lightelementtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementtype.Location != nil {
		stage.StageBranch(lightelementtype.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range lightelementtype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _lightelementstatuspropertytype := range lightelementtype.Availability {
		stage.StageBranch(_lightelementstatuspropertytype)
	}
	for _, _a_extension_10 := range lightelementtype.Extension {
		stage.StageBranch(_a_extension_10)
	}

}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linguisticnotepropertytype) {
		return
	}

	linguisticnotepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linguisticnotepropertytype.LinguisticNote != nil {
		stage.StageBranch(linguisticnotepropertytype.LinguisticNote)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linguisticnotetype *LinguisticNoteType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linguisticnotetype) {
		return
	}

	linguisticnotetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_extension_4 := range linguisticnotetype.Extension {
		stage.StageBranch(_a_extension_4)
	}

}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(markerbeaconpropertytype) {
		return
	}

	markerbeaconpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(navaidequipmentpropertytype) {
		return
	}

	navaidequipmentpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notepropertytype *NotePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notepropertytype) {
		return
	}

	notepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notepropertytype.Note != nil {
		stage.StageBranch(notepropertytype.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetype *NoteType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notetype) {
		return
	}

	notetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
		stage.StageBranch(_linguisticnotepropertytype)
	}
	for _, _a_extension_5 := range notetype.Extension {
		stage.StageBranch(_a_extension_5)
	}

}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(obstacleareatimeslicepropertytype) {
		return
	}

	obstacleareatimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
		stage.StageBranch(obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(obstacleareatimeslicetype) {
		return
	}

	obstacleareatimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
		stage.StageBranch(obstacleareatimeslicetype.Reference_ownerOrganisation)
	}
	if obstacleareatimeslicetype.Reference_ownerRunway != nil {
		stage.StageBranch(obstacleareatimeslicetype.Reference_ownerRunway)
	}
	if obstacleareatimeslicetype.Reference_ownerAirport != nil {
		stage.StageBranch(obstacleareatimeslicetype.Reference_ownerAirport)
	}
	if obstacleareatimeslicetype.SurfaceExtent != nil {
		stage.StageBranch(obstacleareatimeslicetype.SurfaceExtent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepropertytype := range obstacleareatimeslicetype.Obstacle {
		stage.StageBranch(_verticalstructurepropertytype)
	}
	for _, _notepropertytype := range obstacleareatimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_6 := range obstacleareatimeslicetype.Extension {
		stage.StageBranch(_a_extension_6)
	}

}

func (obstacleareatype *ObstacleAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(obstacleareatype) {
		return
	}

	obstacleareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _obstacleareatimeslicepropertytype := range obstacleareatype.TimeSlice {
		stage.StageBranch(_obstacleareatimeslicepropertytype)
	}

}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(organisationauthoritypropertytype) {
		return
	}

	organisationauthoritypropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (passengerservicepropertytype *PassengerServicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(passengerservicepropertytype) {
		return
	}

	passengerservicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pointtype) {
		return
	}

	pointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range pointtype.Annotation {
		stage.StageBranch(_notepropertytype)
	}

}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(runwaydirectionpropertytype) {
		return
	}

	runwaydirectionpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (servicepropertytype *ServicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(servicepropertytype) {
		return
	}

	servicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(specialnavigationstationpropertytype) {
		return
	}

	specialnavigationstationpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (surfacepropertytype *SurfacePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(surfacepropertytype) {
		return
	}

	surfacepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if surfacepropertytype.Surface != nil {
		stage.StageBranch(surfacepropertytype.Surface)
	}
	if surfacepropertytype.ElevatedSurface != nil {
		stage.StageBranch(surfacepropertytype.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (surfacetype *SurfaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(surfacetype) {
		return
	}

	surfacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range surfacetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}

}

func (timesheetpropertytype *TimesheetPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timesheetpropertytype) {
		return
	}

	timesheetpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timesheetpropertytype.Timesheet != nil {
		stage.StageBranch(timesheetpropertytype.Timesheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timesheettype *TimesheetType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timesheettype) {
		return
	}

	timesheettype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range timesheettype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_12 := range timesheettype.Extension {
		stage.StageBranch(_a_extension_12)
	}

}

func (unitpropertytype *UnitPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(unitpropertytype) {
		return
	}

	unitpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructurelightingstatuspropertytype) {
		return
	}

	verticalstructurelightingstatuspropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
		stage.StageBranch(verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructurelightingstatustype) {
		return
	}

	verticalstructurelightingstatustype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructurelightingstatustype.TimeInterval {
		stage.StageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range verticalstructurelightingstatustype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructurelightingstatustype.SpecialDateAuthority {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _a_extension_8 := range verticalstructurelightingstatustype.Extension {
		stage.StageBranch(_a_extension_8)
	}

}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructurepartpropertytype) {
		return
	}

	verticalstructurepartpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurepartpropertytype.VerticalStructurePart != nil {
		stage.StageBranch(verticalstructurepartpropertytype.VerticalStructurePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructureparttype *VerticalStructurePartType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructureparttype) {
		return
	}

	verticalstructureparttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
		stage.StageBranch(verticalstructureparttype.HorizontalProjection_surfaceExtent)
	}
	if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
		stage.StageBranch(verticalstructureparttype.HorizontalProjection_linearExtent)
	}
	if verticalstructureparttype.HorizontalProjection_location != nil {
		stage.StageBranch(verticalstructureparttype.HorizontalProjection_location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructureparttype.TimeInterval {
		stage.StageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range verticalstructureparttype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructureparttype.SpecialDateAuthority {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _lightelementpropertytype := range verticalstructureparttype.Lighting {
		stage.StageBranch(_lightelementpropertytype)
	}
	for _, _a_extension_9 := range verticalstructureparttype.Extension {
		stage.StageBranch(_a_extension_9)
	}

}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructurepropertytype) {
		return
	}

	verticalstructurepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructuretimeslicepropertytype) {
		return
	}

	verticalstructuretimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
		stage.StageBranch(verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructuretimeslicetype) {
		return
	}

	verticalstructuretimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicetype.Marker != nil {
		stage.StageBranch(verticalstructuretimeslicetype.Marker)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepartpropertytype := range verticalstructuretimeslicetype.Part {
		stage.StageBranch(_verticalstructurepartpropertytype)
	}
	for _, _passengerservicepropertytype := range verticalstructuretimeslicetype.HostedPassengerService {
		stage.StageBranch(_passengerservicepropertytype)
	}
	for _, _groundlightsystempropertytype := range verticalstructuretimeslicetype.SupportedGroundLight {
		stage.StageBranch(_groundlightsystempropertytype)
	}
	for _, _navaidequipmentpropertytype := range verticalstructuretimeslicetype.HostedNavaidEquipment {
		stage.StageBranch(_navaidequipmentpropertytype)
	}
	for _, _specialnavigationstationpropertytype := range verticalstructuretimeslicetype.HostedSpecialNavStation {
		stage.StageBranch(_specialnavigationstationpropertytype)
	}
	for _, _unitpropertytype := range verticalstructuretimeslicetype.HostedUnit {
		stage.StageBranch(_unitpropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructuretimeslicetype.HostedOrganisation {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _servicepropertytype := range verticalstructuretimeslicetype.SupportedService {
		stage.StageBranch(_servicepropertytype)
	}
	for _, _notepropertytype := range verticalstructuretimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _verticalstructurelightingstatuspropertytype := range verticalstructuretimeslicetype.LightingAvailability {
		stage.StageBranch(_verticalstructurelightingstatuspropertytype)
	}
	for _, _a_extension_7 := range verticalstructuretimeslicetype.Extension {
		stage.StageBranch(_a_extension_7)
	}

}

func (verticalstructuretype *VerticalStructureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticalstructuretype) {
		return
	}

	verticalstructuretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructuretimeslicepropertytype := range verticalstructuretype.TimeSlice {
		stage.StageBranch(_verticalstructuretimeslicepropertytype)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *A_extension:
		toT := GongCopyBranchA_extension(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_1:
		toT := GongCopyBranchA_extension_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_10:
		toT := GongCopyBranchA_extension_10(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_11:
		toT := GongCopyBranchA_extension_11(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_12:
		toT := GongCopyBranchA_extension_12(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_2:
		toT := GongCopyBranchA_extension_2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_3:
		toT := GongCopyBranchA_extension_3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_4:
		toT := GongCopyBranchA_extension_4(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_5:
		toT := GongCopyBranchA_extension_5(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_6:
		toT := GongCopyBranchA_extension_6(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_7:
		toT := GongCopyBranchA_extension_7(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_8:
		toT := GongCopyBranchA_extension_8(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_9:
		toT := GongCopyBranchA_extension_9(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractPropertiesWithScheduleType:
		toT := GongCopyBranchAbstractPropertiesWithScheduleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AeronauticalGroundLightTimeSlicePropertyType:
		toT := GongCopyBranchAeronauticalGroundLightTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AeronauticalGroundLightTimeSliceType:
		toT := GongCopyBranchAeronauticalGroundLightTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AeronauticalGroundLightType:
		toT := GongCopyBranchAeronauticalGroundLightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirportHeliportPropertyType:
		toT := GongCopyBranchAirportHeliportPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurveType:
		toT := GongCopyBranchCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedCurvePropertyType:
		toT := GongCopyBranchElevatedCurvePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedCurveType:
		toT := GongCopyBranchElevatedCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedPointPropertyType:
		toT := GongCopyBranchElevatedPointPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedPointType:
		toT := GongCopyBranchElevatedPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedSurfacePropertyType:
		toT := GongCopyBranchElevatedSurfacePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedSurfaceType:
		toT := GongCopyBranchElevatedSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroundLightSystemPropertyType:
		toT := GongCopyBranchGroundLightSystemPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LightElementPropertyType:
		toT := GongCopyBranchLightElementPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LightElementStatusPropertyType:
		toT := GongCopyBranchLightElementStatusPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LightElementStatusType:
		toT := GongCopyBranchLightElementStatusType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LightElementType:
		toT := GongCopyBranchLightElementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinguisticNotePropertyType:
		toT := GongCopyBranchLinguisticNotePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinguisticNoteType:
		toT := GongCopyBranchLinguisticNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MarkerBeaconPropertyType:
		toT := GongCopyBranchMarkerBeaconPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NavaidEquipmentPropertyType:
		toT := GongCopyBranchNavaidEquipmentPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePropertyType:
		toT := GongCopyBranchNotePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteType:
		toT := GongCopyBranchNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ObstacleAreaTimeSlicePropertyType:
		toT := GongCopyBranchObstacleAreaTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ObstacleAreaTimeSliceType:
		toT := GongCopyBranchObstacleAreaTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ObstacleAreaType:
		toT := GongCopyBranchObstacleAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OrganisationAuthorityPropertyType:
		toT := GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PassengerServicePropertyType:
		toT := GongCopyBranchPassengerServicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointType:
		toT := GongCopyBranchPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RunwayDirectionPropertyType:
		toT := GongCopyBranchRunwayDirectionPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ServicePropertyType:
		toT := GongCopyBranchServicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpecialNavigationStationPropertyType:
		toT := GongCopyBranchSpecialNavigationStationPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SurfacePropertyType:
		toT := GongCopyBranchSurfacePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SurfaceType:
		toT := GongCopyBranchSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimesheetPropertyType:
		toT := GongCopyBranchTimesheetPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimesheetType:
		toT := GongCopyBranchTimesheetType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UnitPropertyType:
		toT := GongCopyBranchUnitPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructureLightingStatusPropertyType:
		toT := GongCopyBranchVerticalStructureLightingStatusPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructureLightingStatusType:
		toT := GongCopyBranchVerticalStructureLightingStatusType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructurePartPropertyType:
		toT := GongCopyBranchVerticalStructurePartPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructurePartType:
		toT := GongCopyBranchVerticalStructurePartType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructurePropertyType:
		toT := GongCopyBranchVerticalStructurePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructureTimeSlicePropertyType:
		toT := GongCopyBranchVerticalStructureTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructureTimeSliceType:
		toT := GongCopyBranchVerticalStructureTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalStructureType:
		toT := GongCopyBranchVerticalStructureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchA_extension(mapOrigCopy map[any]any, a_extensionFrom *A_extension) (a_extensionTo *A_extension) {
	var alreadyCopied bool
	a_extensionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extensionFrom)
	if alreadyCopied {
		return
	}
	a_extensionFrom.GongCopyBasicFields(a_extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_1(mapOrigCopy map[any]any, a_extension_1From *A_extension_1) (a_extension_1To *A_extension_1) {
	var alreadyCopied bool
	a_extension_1To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_1From)
	if alreadyCopied {
		return
	}
	a_extension_1From.GongCopyBasicFields(a_extension_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_10(mapOrigCopy map[any]any, a_extension_10From *A_extension_10) (a_extension_10To *A_extension_10) {
	var alreadyCopied bool
	a_extension_10To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_10From)
	if alreadyCopied {
		return
	}
	a_extension_10From.GongCopyBasicFields(a_extension_10To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_11(mapOrigCopy map[any]any, a_extension_11From *A_extension_11) (a_extension_11To *A_extension_11) {
	var alreadyCopied bool
	a_extension_11To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_11From)
	if alreadyCopied {
		return
	}
	a_extension_11From.GongCopyBasicFields(a_extension_11To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_12(mapOrigCopy map[any]any, a_extension_12From *A_extension_12) (a_extension_12To *A_extension_12) {
	var alreadyCopied bool
	a_extension_12To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_12From)
	if alreadyCopied {
		return
	}
	a_extension_12From.GongCopyBasicFields(a_extension_12To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_2(mapOrigCopy map[any]any, a_extension_2From *A_extension_2) (a_extension_2To *A_extension_2) {
	var alreadyCopied bool
	a_extension_2To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_2From)
	if alreadyCopied {
		return
	}
	a_extension_2From.GongCopyBasicFields(a_extension_2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_3(mapOrigCopy map[any]any, a_extension_3From *A_extension_3) (a_extension_3To *A_extension_3) {
	var alreadyCopied bool
	a_extension_3To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_3From)
	if alreadyCopied {
		return
	}
	a_extension_3From.GongCopyBasicFields(a_extension_3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_4(mapOrigCopy map[any]any, a_extension_4From *A_extension_4) (a_extension_4To *A_extension_4) {
	var alreadyCopied bool
	a_extension_4To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_4From)
	if alreadyCopied {
		return
	}
	a_extension_4From.GongCopyBasicFields(a_extension_4To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_5(mapOrigCopy map[any]any, a_extension_5From *A_extension_5) (a_extension_5To *A_extension_5) {
	var alreadyCopied bool
	a_extension_5To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_5From)
	if alreadyCopied {
		return
	}
	a_extension_5From.GongCopyBasicFields(a_extension_5To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_6(mapOrigCopy map[any]any, a_extension_6From *A_extension_6) (a_extension_6To *A_extension_6) {
	var alreadyCopied bool
	a_extension_6To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_6From)
	if alreadyCopied {
		return
	}
	a_extension_6From.GongCopyBasicFields(a_extension_6To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_7(mapOrigCopy map[any]any, a_extension_7From *A_extension_7) (a_extension_7To *A_extension_7) {
	var alreadyCopied bool
	a_extension_7To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_7From)
	if alreadyCopied {
		return
	}
	a_extension_7From.GongCopyBasicFields(a_extension_7To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_8(mapOrigCopy map[any]any, a_extension_8From *A_extension_8) (a_extension_8To *A_extension_8) {
	var alreadyCopied bool
	a_extension_8To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_8From)
	if alreadyCopied {
		return
	}
	a_extension_8From.GongCopyBasicFields(a_extension_8To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_9(mapOrigCopy map[any]any, a_extension_9From *A_extension_9) (a_extension_9To *A_extension_9) {
	var alreadyCopied bool
	a_extension_9To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_9From)
	if alreadyCopied {
		return
	}
	a_extension_9From.GongCopyBasicFields(a_extension_9To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractPropertiesWithScheduleType(mapOrigCopy map[any]any, abstractpropertieswithscheduletypeFrom *AbstractPropertiesWithScheduleType) (abstractpropertieswithscheduletypeTo *AbstractPropertiesWithScheduleType) {
	var alreadyCopied bool
	abstractpropertieswithscheduletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractpropertieswithscheduletypeFrom)
	if alreadyCopied {
		return
	}
	abstractpropertieswithscheduletypeFrom.GongCopyBasicFields(abstractpropertieswithscheduletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAeronauticalGroundLightTimeSlicePropertyType(mapOrigCopy map[any]any, aeronauticalgroundlighttimeslicepropertytypeFrom *AeronauticalGroundLightTimeSlicePropertyType) (aeronauticalgroundlighttimeslicepropertytypeTo *AeronauticalGroundLightTimeSlicePropertyType) {
	var alreadyCopied bool
	aeronauticalgroundlighttimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aeronauticalgroundlighttimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	aeronauticalgroundlighttimeslicepropertytypeFrom.GongCopyBasicFields(aeronauticalgroundlighttimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicepropertytypeFrom.AeronauticalGroundLightTimeSlice != nil {
		aeronauticalgroundlighttimeslicepropertytypeTo.AeronauticalGroundLightTimeSlice = GongCopyBranchAeronauticalGroundLightTimeSliceType(mapOrigCopy, aeronauticalgroundlighttimeslicepropertytypeFrom.AeronauticalGroundLightTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAeronauticalGroundLightTimeSliceType(mapOrigCopy map[any]any, aeronauticalgroundlighttimeslicetypeFrom *AeronauticalGroundLightTimeSliceType) (aeronauticalgroundlighttimeslicetypeTo *AeronauticalGroundLightTimeSliceType) {
	var alreadyCopied bool
	aeronauticalgroundlighttimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aeronauticalgroundlighttimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	aeronauticalgroundlighttimeslicetypeFrom.GongCopyBasicFields(aeronauticalgroundlighttimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicetypeFrom.StructureBeacon != nil {
		aeronauticalgroundlighttimeslicetypeTo.StructureBeacon = GongCopyBranchVerticalStructurePropertyType(mapOrigCopy, aeronauticalgroundlighttimeslicetypeFrom.StructureBeacon)
	}
	if aeronauticalgroundlighttimeslicetypeFrom.AerodromeBeacon != nil {
		aeronauticalgroundlighttimeslicetypeTo.AerodromeBeacon = GongCopyBranchAirportHeliportPropertyType(mapOrigCopy, aeronauticalgroundlighttimeslicetypeFrom.AerodromeBeacon)
	}
	if aeronauticalgroundlighttimeslicetypeFrom.Location != nil {
		aeronauticalgroundlighttimeslicetypeTo.Location = GongCopyBranchElevatedPointPropertyType(mapOrigCopy, aeronauticalgroundlighttimeslicetypeFrom.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aeronauticalgroundlighttimeslicetypeFrom.Annotation {
		aeronauticalgroundlighttimeslicetypeTo.Annotation = append(aeronauticalgroundlighttimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_3 := range aeronauticalgroundlighttimeslicetypeFrom.Extension {
		aeronauticalgroundlighttimeslicetypeTo.Extension = append(aeronauticalgroundlighttimeslicetypeTo.Extension, GongCopyBranchA_extension_3(mapOrigCopy, _a_extension_3))
	}

	return
}

func GongCopyBranchAeronauticalGroundLightType(mapOrigCopy map[any]any, aeronauticalgroundlighttypeFrom *AeronauticalGroundLightType) (aeronauticalgroundlighttypeTo *AeronauticalGroundLightType) {
	var alreadyCopied bool
	aeronauticalgroundlighttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aeronauticalgroundlighttypeFrom)
	if alreadyCopied {
		return
	}
	aeronauticalgroundlighttypeFrom.GongCopyBasicFields(aeronauticalgroundlighttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _aeronauticalgroundlighttimeslicepropertytype := range aeronauticalgroundlighttypeFrom.TimeSlice {
		aeronauticalgroundlighttypeTo.TimeSlice = append(aeronauticalgroundlighttypeTo.TimeSlice, GongCopyBranchAeronauticalGroundLightTimeSlicePropertyType(mapOrigCopy, _aeronauticalgroundlighttimeslicepropertytype))
	}

	return
}

func GongCopyBranchAirportHeliportPropertyType(mapOrigCopy map[any]any, airportheliportpropertytypeFrom *AirportHeliportPropertyType) (airportheliportpropertytypeTo *AirportHeliportPropertyType) {
	var alreadyCopied bool
	airportheliportpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airportheliportpropertytypeFrom)
	if alreadyCopied {
		return
	}
	airportheliportpropertytypeFrom.GongCopyBasicFields(airportheliportpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCurveType(mapOrigCopy map[any]any, curvetypeFrom *CurveType) (curvetypeTo *CurveType) {
	var alreadyCopied bool
	curvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curvetypeFrom)
	if alreadyCopied {
		return
	}
	curvetypeFrom.GongCopyBasicFields(curvetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range curvetypeFrom.Annotation {
		curvetypeTo.Annotation = append(curvetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}

	return
}

func GongCopyBranchElevatedCurvePropertyType(mapOrigCopy map[any]any, elevatedcurvepropertytypeFrom *ElevatedCurvePropertyType) (elevatedcurvepropertytypeTo *ElevatedCurvePropertyType) {
	var alreadyCopied bool
	elevatedcurvepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedcurvepropertytypeFrom)
	if alreadyCopied {
		return
	}
	elevatedcurvepropertytypeFrom.GongCopyBasicFields(elevatedcurvepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if elevatedcurvepropertytypeFrom.ElevatedCurve != nil {
		elevatedcurvepropertytypeTo.ElevatedCurve = GongCopyBranchElevatedCurveType(mapOrigCopy, elevatedcurvepropertytypeFrom.ElevatedCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElevatedCurveType(mapOrigCopy map[any]any, elevatedcurvetypeFrom *ElevatedCurveType) (elevatedcurvetypeTo *ElevatedCurveType) {
	var alreadyCopied bool
	elevatedcurvetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedcurvetypeFrom)
	if alreadyCopied {
		return
	}
	elevatedcurvetypeFrom.GongCopyBasicFields(elevatedcurvetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedcurvetypeFrom.Annotation {
		elevatedcurvetypeTo.Annotation = append(elevatedcurvetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension := range elevatedcurvetypeFrom.Extension {
		elevatedcurvetypeTo.Extension = append(elevatedcurvetypeTo.Extension, GongCopyBranchA_extension(mapOrigCopy, _a_extension))
	}

	return
}

func GongCopyBranchElevatedPointPropertyType(mapOrigCopy map[any]any, elevatedpointpropertytypeFrom *ElevatedPointPropertyType) (elevatedpointpropertytypeTo *ElevatedPointPropertyType) {
	var alreadyCopied bool
	elevatedpointpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedpointpropertytypeFrom)
	if alreadyCopied {
		return
	}
	elevatedpointpropertytypeFrom.GongCopyBasicFields(elevatedpointpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if elevatedpointpropertytypeFrom.ElevatedPoint != nil {
		elevatedpointpropertytypeTo.ElevatedPoint = GongCopyBranchElevatedPointType(mapOrigCopy, elevatedpointpropertytypeFrom.ElevatedPoint)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElevatedPointType(mapOrigCopy map[any]any, elevatedpointtypeFrom *ElevatedPointType) (elevatedpointtypeTo *ElevatedPointType) {
	var alreadyCopied bool
	elevatedpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedpointtypeFrom)
	if alreadyCopied {
		return
	}
	elevatedpointtypeFrom.GongCopyBasicFields(elevatedpointtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedpointtypeFrom.Annotation {
		elevatedpointtypeTo.Annotation = append(elevatedpointtypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_1 := range elevatedpointtypeFrom.Extension {
		elevatedpointtypeTo.Extension = append(elevatedpointtypeTo.Extension, GongCopyBranchA_extension_1(mapOrigCopy, _a_extension_1))
	}

	return
}

func GongCopyBranchElevatedSurfacePropertyType(mapOrigCopy map[any]any, elevatedsurfacepropertytypeFrom *ElevatedSurfacePropertyType) (elevatedsurfacepropertytypeTo *ElevatedSurfacePropertyType) {
	var alreadyCopied bool
	elevatedsurfacepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedsurfacepropertytypeFrom)
	if alreadyCopied {
		return
	}
	elevatedsurfacepropertytypeFrom.GongCopyBasicFields(elevatedsurfacepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if elevatedsurfacepropertytypeFrom.ElevatedSurface != nil {
		elevatedsurfacepropertytypeTo.ElevatedSurface = GongCopyBranchElevatedSurfaceType(mapOrigCopy, elevatedsurfacepropertytypeFrom.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElevatedSurfaceType(mapOrigCopy map[any]any, elevatedsurfacetypeFrom *ElevatedSurfaceType) (elevatedsurfacetypeTo *ElevatedSurfaceType) {
	var alreadyCopied bool
	elevatedsurfacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elevatedsurfacetypeFrom)
	if alreadyCopied {
		return
	}
	elevatedsurfacetypeFrom.GongCopyBasicFields(elevatedsurfacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedsurfacetypeFrom.Annotation {
		elevatedsurfacetypeTo.Annotation = append(elevatedsurfacetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_2 := range elevatedsurfacetypeFrom.Extension {
		elevatedsurfacetypeTo.Extension = append(elevatedsurfacetypeTo.Extension, GongCopyBranchA_extension_2(mapOrigCopy, _a_extension_2))
	}

	return
}

func GongCopyBranchGroundLightSystemPropertyType(mapOrigCopy map[any]any, groundlightsystempropertytypeFrom *GroundLightSystemPropertyType) (groundlightsystempropertytypeTo *GroundLightSystemPropertyType) {
	var alreadyCopied bool
	groundlightsystempropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groundlightsystempropertytypeFrom)
	if alreadyCopied {
		return
	}
	groundlightsystempropertytypeFrom.GongCopyBasicFields(groundlightsystempropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLightElementPropertyType(mapOrigCopy map[any]any, lightelementpropertytypeFrom *LightElementPropertyType) (lightelementpropertytypeTo *LightElementPropertyType) {
	var alreadyCopied bool
	lightelementpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lightelementpropertytypeFrom)
	if alreadyCopied {
		return
	}
	lightelementpropertytypeFrom.GongCopyBasicFields(lightelementpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if lightelementpropertytypeFrom.LightElement != nil {
		lightelementpropertytypeTo.LightElement = GongCopyBranchLightElementType(mapOrigCopy, lightelementpropertytypeFrom.LightElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLightElementStatusPropertyType(mapOrigCopy map[any]any, lightelementstatuspropertytypeFrom *LightElementStatusPropertyType) (lightelementstatuspropertytypeTo *LightElementStatusPropertyType) {
	var alreadyCopied bool
	lightelementstatuspropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lightelementstatuspropertytypeFrom)
	if alreadyCopied {
		return
	}
	lightelementstatuspropertytypeFrom.GongCopyBasicFields(lightelementstatuspropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if lightelementstatuspropertytypeFrom.LightElementStatus != nil {
		lightelementstatuspropertytypeTo.LightElementStatus = GongCopyBranchLightElementStatusType(mapOrigCopy, lightelementstatuspropertytypeFrom.LightElementStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLightElementStatusType(mapOrigCopy map[any]any, lightelementstatustypeFrom *LightElementStatusType) (lightelementstatustypeTo *LightElementStatusType) {
	var alreadyCopied bool
	lightelementstatustypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lightelementstatustypeFrom)
	if alreadyCopied {
		return
	}
	lightelementstatustypeFrom.GongCopyBasicFields(lightelementstatustypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range lightelementstatustypeFrom.TimeInterval {
		lightelementstatustypeTo.TimeInterval = append(lightelementstatustypeTo.TimeInterval, GongCopyBranchTimesheetPropertyType(mapOrigCopy, _timesheetpropertytype))
	}
	for _, _notepropertytype := range lightelementstatustypeFrom.Annotation {
		lightelementstatustypeTo.Annotation = append(lightelementstatustypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _organisationauthoritypropertytype := range lightelementstatustypeFrom.SpecialDateAuthority {
		lightelementstatustypeTo.SpecialDateAuthority = append(lightelementstatustypeTo.SpecialDateAuthority, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _a_extension_11 := range lightelementstatustypeFrom.Extension {
		lightelementstatustypeTo.Extension = append(lightelementstatustypeTo.Extension, GongCopyBranchA_extension_11(mapOrigCopy, _a_extension_11))
	}

	return
}

func GongCopyBranchLightElementType(mapOrigCopy map[any]any, lightelementtypeFrom *LightElementType) (lightelementtypeTo *LightElementType) {
	var alreadyCopied bool
	lightelementtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lightelementtypeFrom)
	if alreadyCopied {
		return
	}
	lightelementtypeFrom.GongCopyBasicFields(lightelementtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if lightelementtypeFrom.Location != nil {
		lightelementtypeTo.Location = GongCopyBranchElevatedPointPropertyType(mapOrigCopy, lightelementtypeFrom.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range lightelementtypeFrom.Annotation {
		lightelementtypeTo.Annotation = append(lightelementtypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _lightelementstatuspropertytype := range lightelementtypeFrom.Availability {
		lightelementtypeTo.Availability = append(lightelementtypeTo.Availability, GongCopyBranchLightElementStatusPropertyType(mapOrigCopy, _lightelementstatuspropertytype))
	}
	for _, _a_extension_10 := range lightelementtypeFrom.Extension {
		lightelementtypeTo.Extension = append(lightelementtypeTo.Extension, GongCopyBranchA_extension_10(mapOrigCopy, _a_extension_10))
	}

	return
}

func GongCopyBranchLinguisticNotePropertyType(mapOrigCopy map[any]any, linguisticnotepropertytypeFrom *LinguisticNotePropertyType) (linguisticnotepropertytypeTo *LinguisticNotePropertyType) {
	var alreadyCopied bool
	linguisticnotepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linguisticnotepropertytypeFrom)
	if alreadyCopied {
		return
	}
	linguisticnotepropertytypeFrom.GongCopyBasicFields(linguisticnotepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if linguisticnotepropertytypeFrom.LinguisticNote != nil {
		linguisticnotepropertytypeTo.LinguisticNote = GongCopyBranchLinguisticNoteType(mapOrigCopy, linguisticnotepropertytypeFrom.LinguisticNote)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLinguisticNoteType(mapOrigCopy map[any]any, linguisticnotetypeFrom *LinguisticNoteType) (linguisticnotetypeTo *LinguisticNoteType) {
	var alreadyCopied bool
	linguisticnotetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linguisticnotetypeFrom)
	if alreadyCopied {
		return
	}
	linguisticnotetypeFrom.GongCopyBasicFields(linguisticnotetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_extension_4 := range linguisticnotetypeFrom.Extension {
		linguisticnotetypeTo.Extension = append(linguisticnotetypeTo.Extension, GongCopyBranchA_extension_4(mapOrigCopy, _a_extension_4))
	}

	return
}

func GongCopyBranchMarkerBeaconPropertyType(mapOrigCopy map[any]any, markerbeaconpropertytypeFrom *MarkerBeaconPropertyType) (markerbeaconpropertytypeTo *MarkerBeaconPropertyType) {
	var alreadyCopied bool
	markerbeaconpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, markerbeaconpropertytypeFrom)
	if alreadyCopied {
		return
	}
	markerbeaconpropertytypeFrom.GongCopyBasicFields(markerbeaconpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNavaidEquipmentPropertyType(mapOrigCopy map[any]any, navaidequipmentpropertytypeFrom *NavaidEquipmentPropertyType) (navaidequipmentpropertytypeTo *NavaidEquipmentPropertyType) {
	var alreadyCopied bool
	navaidequipmentpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, navaidequipmentpropertytypeFrom)
	if alreadyCopied {
		return
	}
	navaidequipmentpropertytypeFrom.GongCopyBasicFields(navaidequipmentpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotePropertyType(mapOrigCopy map[any]any, notepropertytypeFrom *NotePropertyType) (notepropertytypeTo *NotePropertyType) {
	var alreadyCopied bool
	notepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notepropertytypeFrom)
	if alreadyCopied {
		return
	}
	notepropertytypeFrom.GongCopyBasicFields(notepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if notepropertytypeFrom.Note != nil {
		notepropertytypeTo.Note = GongCopyBranchNoteType(mapOrigCopy, notepropertytypeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteType(mapOrigCopy map[any]any, notetypeFrom *NoteType) (notetypeTo *NoteType) {
	var alreadyCopied bool
	notetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notetypeFrom)
	if alreadyCopied {
		return
	}
	notetypeFrom.GongCopyBasicFields(notetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linguisticnotepropertytype := range notetypeFrom.TranslatedNote {
		notetypeTo.TranslatedNote = append(notetypeTo.TranslatedNote, GongCopyBranchLinguisticNotePropertyType(mapOrigCopy, _linguisticnotepropertytype))
	}
	for _, _a_extension_5 := range notetypeFrom.Extension {
		notetypeTo.Extension = append(notetypeTo.Extension, GongCopyBranchA_extension_5(mapOrigCopy, _a_extension_5))
	}

	return
}

func GongCopyBranchObstacleAreaTimeSlicePropertyType(mapOrigCopy map[any]any, obstacleareatimeslicepropertytypeFrom *ObstacleAreaTimeSlicePropertyType) (obstacleareatimeslicepropertytypeTo *ObstacleAreaTimeSlicePropertyType) {
	var alreadyCopied bool
	obstacleareatimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, obstacleareatimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	obstacleareatimeslicepropertytypeFrom.GongCopyBasicFields(obstacleareatimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicepropertytypeFrom.ObstacleAreaTimeSlice != nil {
		obstacleareatimeslicepropertytypeTo.ObstacleAreaTimeSlice = GongCopyBranchObstacleAreaTimeSliceType(mapOrigCopy, obstacleareatimeslicepropertytypeFrom.ObstacleAreaTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchObstacleAreaTimeSliceType(mapOrigCopy map[any]any, obstacleareatimeslicetypeFrom *ObstacleAreaTimeSliceType) (obstacleareatimeslicetypeTo *ObstacleAreaTimeSliceType) {
	var alreadyCopied bool
	obstacleareatimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, obstacleareatimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	obstacleareatimeslicetypeFrom.GongCopyBasicFields(obstacleareatimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicetypeFrom.Reference_ownerOrganisation != nil {
		obstacleareatimeslicetypeTo.Reference_ownerOrganisation = GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, obstacleareatimeslicetypeFrom.Reference_ownerOrganisation)
	}
	if obstacleareatimeslicetypeFrom.Reference_ownerRunway != nil {
		obstacleareatimeslicetypeTo.Reference_ownerRunway = GongCopyBranchRunwayDirectionPropertyType(mapOrigCopy, obstacleareatimeslicetypeFrom.Reference_ownerRunway)
	}
	if obstacleareatimeslicetypeFrom.Reference_ownerAirport != nil {
		obstacleareatimeslicetypeTo.Reference_ownerAirport = GongCopyBranchAirportHeliportPropertyType(mapOrigCopy, obstacleareatimeslicetypeFrom.Reference_ownerAirport)
	}
	if obstacleareatimeslicetypeFrom.SurfaceExtent != nil {
		obstacleareatimeslicetypeTo.SurfaceExtent = GongCopyBranchSurfacePropertyType(mapOrigCopy, obstacleareatimeslicetypeFrom.SurfaceExtent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepropertytype := range obstacleareatimeslicetypeFrom.Obstacle {
		obstacleareatimeslicetypeTo.Obstacle = append(obstacleareatimeslicetypeTo.Obstacle, GongCopyBranchVerticalStructurePropertyType(mapOrigCopy, _verticalstructurepropertytype))
	}
	for _, _notepropertytype := range obstacleareatimeslicetypeFrom.Annotation {
		obstacleareatimeslicetypeTo.Annotation = append(obstacleareatimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_6 := range obstacleareatimeslicetypeFrom.Extension {
		obstacleareatimeslicetypeTo.Extension = append(obstacleareatimeslicetypeTo.Extension, GongCopyBranchA_extension_6(mapOrigCopy, _a_extension_6))
	}

	return
}

func GongCopyBranchObstacleAreaType(mapOrigCopy map[any]any, obstacleareatypeFrom *ObstacleAreaType) (obstacleareatypeTo *ObstacleAreaType) {
	var alreadyCopied bool
	obstacleareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, obstacleareatypeFrom)
	if alreadyCopied {
		return
	}
	obstacleareatypeFrom.GongCopyBasicFields(obstacleareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _obstacleareatimeslicepropertytype := range obstacleareatypeFrom.TimeSlice {
		obstacleareatypeTo.TimeSlice = append(obstacleareatypeTo.TimeSlice, GongCopyBranchObstacleAreaTimeSlicePropertyType(mapOrigCopy, _obstacleareatimeslicepropertytype))
	}

	return
}

func GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy map[any]any, organisationauthoritypropertytypeFrom *OrganisationAuthorityPropertyType) (organisationauthoritypropertytypeTo *OrganisationAuthorityPropertyType) {
	var alreadyCopied bool
	organisationauthoritypropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, organisationauthoritypropertytypeFrom)
	if alreadyCopied {
		return
	}
	organisationauthoritypropertytypeFrom.GongCopyBasicFields(organisationauthoritypropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPassengerServicePropertyType(mapOrigCopy map[any]any, passengerservicepropertytypeFrom *PassengerServicePropertyType) (passengerservicepropertytypeTo *PassengerServicePropertyType) {
	var alreadyCopied bool
	passengerservicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, passengerservicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	passengerservicepropertytypeFrom.GongCopyBasicFields(passengerservicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointType(mapOrigCopy map[any]any, pointtypeFrom *PointType) (pointtypeTo *PointType) {
	var alreadyCopied bool
	pointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointtypeFrom)
	if alreadyCopied {
		return
	}
	pointtypeFrom.GongCopyBasicFields(pointtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range pointtypeFrom.Annotation {
		pointtypeTo.Annotation = append(pointtypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}

	return
}

func GongCopyBranchRunwayDirectionPropertyType(mapOrigCopy map[any]any, runwaydirectionpropertytypeFrom *RunwayDirectionPropertyType) (runwaydirectionpropertytypeTo *RunwayDirectionPropertyType) {
	var alreadyCopied bool
	runwaydirectionpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, runwaydirectionpropertytypeFrom)
	if alreadyCopied {
		return
	}
	runwaydirectionpropertytypeFrom.GongCopyBasicFields(runwaydirectionpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchServicePropertyType(mapOrigCopy map[any]any, servicepropertytypeFrom *ServicePropertyType) (servicepropertytypeTo *ServicePropertyType) {
	var alreadyCopied bool
	servicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, servicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	servicepropertytypeFrom.GongCopyBasicFields(servicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSpecialNavigationStationPropertyType(mapOrigCopy map[any]any, specialnavigationstationpropertytypeFrom *SpecialNavigationStationPropertyType) (specialnavigationstationpropertytypeTo *SpecialNavigationStationPropertyType) {
	var alreadyCopied bool
	specialnavigationstationpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, specialnavigationstationpropertytypeFrom)
	if alreadyCopied {
		return
	}
	specialnavigationstationpropertytypeFrom.GongCopyBasicFields(specialnavigationstationpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSurfacePropertyType(mapOrigCopy map[any]any, surfacepropertytypeFrom *SurfacePropertyType) (surfacepropertytypeTo *SurfacePropertyType) {
	var alreadyCopied bool
	surfacepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, surfacepropertytypeFrom)
	if alreadyCopied {
		return
	}
	surfacepropertytypeFrom.GongCopyBasicFields(surfacepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if surfacepropertytypeFrom.Surface != nil {
		surfacepropertytypeTo.Surface = GongCopyBranchSurfaceType(mapOrigCopy, surfacepropertytypeFrom.Surface)
	}
	if surfacepropertytypeFrom.ElevatedSurface != nil {
		surfacepropertytypeTo.ElevatedSurface = GongCopyBranchElevatedSurfaceType(mapOrigCopy, surfacepropertytypeFrom.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSurfaceType(mapOrigCopy map[any]any, surfacetypeFrom *SurfaceType) (surfacetypeTo *SurfaceType) {
	var alreadyCopied bool
	surfacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, surfacetypeFrom)
	if alreadyCopied {
		return
	}
	surfacetypeFrom.GongCopyBasicFields(surfacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range surfacetypeFrom.Annotation {
		surfacetypeTo.Annotation = append(surfacetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}

	return
}

func GongCopyBranchTimesheetPropertyType(mapOrigCopy map[any]any, timesheetpropertytypeFrom *TimesheetPropertyType) (timesheetpropertytypeTo *TimesheetPropertyType) {
	var alreadyCopied bool
	timesheetpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timesheetpropertytypeFrom)
	if alreadyCopied {
		return
	}
	timesheetpropertytypeFrom.GongCopyBasicFields(timesheetpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timesheetpropertytypeFrom.Timesheet != nil {
		timesheetpropertytypeTo.Timesheet = GongCopyBranchTimesheetType(mapOrigCopy, timesheetpropertytypeFrom.Timesheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimesheetType(mapOrigCopy map[any]any, timesheettypeFrom *TimesheetType) (timesheettypeTo *TimesheetType) {
	var alreadyCopied bool
	timesheettypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timesheettypeFrom)
	if alreadyCopied {
		return
	}
	timesheettypeFrom.GongCopyBasicFields(timesheettypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range timesheettypeFrom.Annotation {
		timesheettypeTo.Annotation = append(timesheettypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_12 := range timesheettypeFrom.Extension {
		timesheettypeTo.Extension = append(timesheettypeTo.Extension, GongCopyBranchA_extension_12(mapOrigCopy, _a_extension_12))
	}

	return
}

func GongCopyBranchUnitPropertyType(mapOrigCopy map[any]any, unitpropertytypeFrom *UnitPropertyType) (unitpropertytypeTo *UnitPropertyType) {
	var alreadyCopied bool
	unitpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, unitpropertytypeFrom)
	if alreadyCopied {
		return
	}
	unitpropertytypeFrom.GongCopyBasicFields(unitpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalStructureLightingStatusPropertyType(mapOrigCopy map[any]any, verticalstructurelightingstatuspropertytypeFrom *VerticalStructureLightingStatusPropertyType) (verticalstructurelightingstatuspropertytypeTo *VerticalStructureLightingStatusPropertyType) {
	var alreadyCopied bool
	verticalstructurelightingstatuspropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructurelightingstatuspropertytypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructurelightingstatuspropertytypeFrom.GongCopyBasicFields(verticalstructurelightingstatuspropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurelightingstatuspropertytypeFrom.VerticalStructureLightingStatus != nil {
		verticalstructurelightingstatuspropertytypeTo.VerticalStructureLightingStatus = GongCopyBranchVerticalStructureLightingStatusType(mapOrigCopy, verticalstructurelightingstatuspropertytypeFrom.VerticalStructureLightingStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalStructureLightingStatusType(mapOrigCopy map[any]any, verticalstructurelightingstatustypeFrom *VerticalStructureLightingStatusType) (verticalstructurelightingstatustypeTo *VerticalStructureLightingStatusType) {
	var alreadyCopied bool
	verticalstructurelightingstatustypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructurelightingstatustypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructurelightingstatustypeFrom.GongCopyBasicFields(verticalstructurelightingstatustypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructurelightingstatustypeFrom.TimeInterval {
		verticalstructurelightingstatustypeTo.TimeInterval = append(verticalstructurelightingstatustypeTo.TimeInterval, GongCopyBranchTimesheetPropertyType(mapOrigCopy, _timesheetpropertytype))
	}
	for _, _notepropertytype := range verticalstructurelightingstatustypeFrom.Annotation {
		verticalstructurelightingstatustypeTo.Annotation = append(verticalstructurelightingstatustypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _organisationauthoritypropertytype := range verticalstructurelightingstatustypeFrom.SpecialDateAuthority {
		verticalstructurelightingstatustypeTo.SpecialDateAuthority = append(verticalstructurelightingstatustypeTo.SpecialDateAuthority, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _a_extension_8 := range verticalstructurelightingstatustypeFrom.Extension {
		verticalstructurelightingstatustypeTo.Extension = append(verticalstructurelightingstatustypeTo.Extension, GongCopyBranchA_extension_8(mapOrigCopy, _a_extension_8))
	}

	return
}

func GongCopyBranchVerticalStructurePartPropertyType(mapOrigCopy map[any]any, verticalstructurepartpropertytypeFrom *VerticalStructurePartPropertyType) (verticalstructurepartpropertytypeTo *VerticalStructurePartPropertyType) {
	var alreadyCopied bool
	verticalstructurepartpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructurepartpropertytypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructurepartpropertytypeFrom.GongCopyBasicFields(verticalstructurepartpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurepartpropertytypeFrom.VerticalStructurePart != nil {
		verticalstructurepartpropertytypeTo.VerticalStructurePart = GongCopyBranchVerticalStructurePartType(mapOrigCopy, verticalstructurepartpropertytypeFrom.VerticalStructurePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalStructurePartType(mapOrigCopy map[any]any, verticalstructureparttypeFrom *VerticalStructurePartType) (verticalstructureparttypeTo *VerticalStructurePartType) {
	var alreadyCopied bool
	verticalstructureparttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructureparttypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructureparttypeFrom.GongCopyBasicFields(verticalstructureparttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructureparttypeFrom.HorizontalProjection_surfaceExtent != nil {
		verticalstructureparttypeTo.HorizontalProjection_surfaceExtent = GongCopyBranchElevatedSurfacePropertyType(mapOrigCopy, verticalstructureparttypeFrom.HorizontalProjection_surfaceExtent)
	}
	if verticalstructureparttypeFrom.HorizontalProjection_linearExtent != nil {
		verticalstructureparttypeTo.HorizontalProjection_linearExtent = GongCopyBranchElevatedCurvePropertyType(mapOrigCopy, verticalstructureparttypeFrom.HorizontalProjection_linearExtent)
	}
	if verticalstructureparttypeFrom.HorizontalProjection_location != nil {
		verticalstructureparttypeTo.HorizontalProjection_location = GongCopyBranchElevatedPointPropertyType(mapOrigCopy, verticalstructureparttypeFrom.HorizontalProjection_location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructureparttypeFrom.TimeInterval {
		verticalstructureparttypeTo.TimeInterval = append(verticalstructureparttypeTo.TimeInterval, GongCopyBranchTimesheetPropertyType(mapOrigCopy, _timesheetpropertytype))
	}
	for _, _notepropertytype := range verticalstructureparttypeFrom.Annotation {
		verticalstructureparttypeTo.Annotation = append(verticalstructureparttypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _organisationauthoritypropertytype := range verticalstructureparttypeFrom.SpecialDateAuthority {
		verticalstructureparttypeTo.SpecialDateAuthority = append(verticalstructureparttypeTo.SpecialDateAuthority, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _lightelementpropertytype := range verticalstructureparttypeFrom.Lighting {
		verticalstructureparttypeTo.Lighting = append(verticalstructureparttypeTo.Lighting, GongCopyBranchLightElementPropertyType(mapOrigCopy, _lightelementpropertytype))
	}
	for _, _a_extension_9 := range verticalstructureparttypeFrom.Extension {
		verticalstructureparttypeTo.Extension = append(verticalstructureparttypeTo.Extension, GongCopyBranchA_extension_9(mapOrigCopy, _a_extension_9))
	}

	return
}

func GongCopyBranchVerticalStructurePropertyType(mapOrigCopy map[any]any, verticalstructurepropertytypeFrom *VerticalStructurePropertyType) (verticalstructurepropertytypeTo *VerticalStructurePropertyType) {
	var alreadyCopied bool
	verticalstructurepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructurepropertytypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructurepropertytypeFrom.GongCopyBasicFields(verticalstructurepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalStructureTimeSlicePropertyType(mapOrigCopy map[any]any, verticalstructuretimeslicepropertytypeFrom *VerticalStructureTimeSlicePropertyType) (verticalstructuretimeslicepropertytypeTo *VerticalStructureTimeSlicePropertyType) {
	var alreadyCopied bool
	verticalstructuretimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructuretimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructuretimeslicepropertytypeFrom.GongCopyBasicFields(verticalstructuretimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicepropertytypeFrom.VerticalStructureTimeSlice != nil {
		verticalstructuretimeslicepropertytypeTo.VerticalStructureTimeSlice = GongCopyBranchVerticalStructureTimeSliceType(mapOrigCopy, verticalstructuretimeslicepropertytypeFrom.VerticalStructureTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalStructureTimeSliceType(mapOrigCopy map[any]any, verticalstructuretimeslicetypeFrom *VerticalStructureTimeSliceType) (verticalstructuretimeslicetypeTo *VerticalStructureTimeSliceType) {
	var alreadyCopied bool
	verticalstructuretimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructuretimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructuretimeslicetypeFrom.GongCopyBasicFields(verticalstructuretimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicetypeFrom.Marker != nil {
		verticalstructuretimeslicetypeTo.Marker = GongCopyBranchMarkerBeaconPropertyType(mapOrigCopy, verticalstructuretimeslicetypeFrom.Marker)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepartpropertytype := range verticalstructuretimeslicetypeFrom.Part {
		verticalstructuretimeslicetypeTo.Part = append(verticalstructuretimeslicetypeTo.Part, GongCopyBranchVerticalStructurePartPropertyType(mapOrigCopy, _verticalstructurepartpropertytype))
	}
	for _, _passengerservicepropertytype := range verticalstructuretimeslicetypeFrom.HostedPassengerService {
		verticalstructuretimeslicetypeTo.HostedPassengerService = append(verticalstructuretimeslicetypeTo.HostedPassengerService, GongCopyBranchPassengerServicePropertyType(mapOrigCopy, _passengerservicepropertytype))
	}
	for _, _groundlightsystempropertytype := range verticalstructuretimeslicetypeFrom.SupportedGroundLight {
		verticalstructuretimeslicetypeTo.SupportedGroundLight = append(verticalstructuretimeslicetypeTo.SupportedGroundLight, GongCopyBranchGroundLightSystemPropertyType(mapOrigCopy, _groundlightsystempropertytype))
	}
	for _, _navaidequipmentpropertytype := range verticalstructuretimeslicetypeFrom.HostedNavaidEquipment {
		verticalstructuretimeslicetypeTo.HostedNavaidEquipment = append(verticalstructuretimeslicetypeTo.HostedNavaidEquipment, GongCopyBranchNavaidEquipmentPropertyType(mapOrigCopy, _navaidequipmentpropertytype))
	}
	for _, _specialnavigationstationpropertytype := range verticalstructuretimeslicetypeFrom.HostedSpecialNavStation {
		verticalstructuretimeslicetypeTo.HostedSpecialNavStation = append(verticalstructuretimeslicetypeTo.HostedSpecialNavStation, GongCopyBranchSpecialNavigationStationPropertyType(mapOrigCopy, _specialnavigationstationpropertytype))
	}
	for _, _unitpropertytype := range verticalstructuretimeslicetypeFrom.HostedUnit {
		verticalstructuretimeslicetypeTo.HostedUnit = append(verticalstructuretimeslicetypeTo.HostedUnit, GongCopyBranchUnitPropertyType(mapOrigCopy, _unitpropertytype))
	}
	for _, _organisationauthoritypropertytype := range verticalstructuretimeslicetypeFrom.HostedOrganisation {
		verticalstructuretimeslicetypeTo.HostedOrganisation = append(verticalstructuretimeslicetypeTo.HostedOrganisation, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _servicepropertytype := range verticalstructuretimeslicetypeFrom.SupportedService {
		verticalstructuretimeslicetypeTo.SupportedService = append(verticalstructuretimeslicetypeTo.SupportedService, GongCopyBranchServicePropertyType(mapOrigCopy, _servicepropertytype))
	}
	for _, _notepropertytype := range verticalstructuretimeslicetypeFrom.Annotation {
		verticalstructuretimeslicetypeTo.Annotation = append(verticalstructuretimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _verticalstructurelightingstatuspropertytype := range verticalstructuretimeslicetypeFrom.LightingAvailability {
		verticalstructuretimeslicetypeTo.LightingAvailability = append(verticalstructuretimeslicetypeTo.LightingAvailability, GongCopyBranchVerticalStructureLightingStatusPropertyType(mapOrigCopy, _verticalstructurelightingstatuspropertytype))
	}
	for _, _a_extension_7 := range verticalstructuretimeslicetypeFrom.Extension {
		verticalstructuretimeslicetypeTo.Extension = append(verticalstructuretimeslicetypeTo.Extension, GongCopyBranchA_extension_7(mapOrigCopy, _a_extension_7))
	}

	return
}

func GongCopyBranchVerticalStructureType(mapOrigCopy map[any]any, verticalstructuretypeFrom *VerticalStructureType) (verticalstructuretypeTo *VerticalStructureType) {
	var alreadyCopied bool
	verticalstructuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticalstructuretypeFrom)
	if alreadyCopied {
		return
	}
	verticalstructuretypeFrom.GongCopyBasicFields(verticalstructuretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructuretimeslicepropertytype := range verticalstructuretypeFrom.TimeSlice {
		verticalstructuretypeTo.TimeSlice = append(verticalstructuretypeTo.TimeSlice, GongCopyBranchVerticalStructureTimeSlicePropertyType(mapOrigCopy, _verticalstructuretimeslicepropertytype))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (a_extension *A_extension) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension) {
		return
	}

	a_extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_1 *A_extension_1) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_1) {
		return
	}

	a_extension_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_10 *A_extension_10) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_10) {
		return
	}

	a_extension_10.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_11 *A_extension_11) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_11) {
		return
	}

	a_extension_11.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_12 *A_extension_12) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_12) {
		return
	}

	a_extension_12.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_2 *A_extension_2) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_2) {
		return
	}

	a_extension_2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_3 *A_extension_3) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_3) {
		return
	}

	a_extension_3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_4 *A_extension_4) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_4) {
		return
	}

	a_extension_4.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_5 *A_extension_5) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_5) {
		return
	}

	a_extension_5.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_6 *A_extension_6) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_6) {
		return
	}

	a_extension_6.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_7 *A_extension_7) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_7) {
		return
	}

	a_extension_7.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_8 *A_extension_8) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_8) {
		return
	}

	a_extension_8.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_9 *A_extension_9) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_9) {
		return
	}

	a_extension_9.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractpropertieswithscheduletype) {
		return
	}

	abstractpropertieswithscheduletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aeronauticalgroundlighttimeslicepropertytype) {
		return
	}

	aeronauticalgroundlighttimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
		stage.UnstageBranch(aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aeronauticalgroundlighttimeslicetype) {
		return
	}

	aeronauticalgroundlighttimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
		stage.UnstageBranch(aeronauticalgroundlighttimeslicetype.StructureBeacon)
	}
	if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
		stage.UnstageBranch(aeronauticalgroundlighttimeslicetype.AerodromeBeacon)
	}
	if aeronauticalgroundlighttimeslicetype.Location != nil {
		stage.UnstageBranch(aeronauticalgroundlighttimeslicetype.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aeronauticalgroundlighttimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_3 := range aeronauticalgroundlighttimeslicetype.Extension {
		stage.UnstageBranch(_a_extension_3)
	}

}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aeronauticalgroundlighttype) {
		return
	}

	aeronauticalgroundlighttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _aeronauticalgroundlighttimeslicepropertytype := range aeronauticalgroundlighttype.TimeSlice {
		stage.UnstageBranch(_aeronauticalgroundlighttimeslicepropertytype)
	}

}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airportheliportpropertytype) {
		return
	}

	airportheliportpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (curvetype *CurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curvetype) {
		return
	}

	curvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range curvetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}

}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedcurvepropertytype) {
		return
	}

	elevatedcurvepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedcurvepropertytype.ElevatedCurve != nil {
		stage.UnstageBranch(elevatedcurvepropertytype.ElevatedCurve)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedcurvetype *ElevatedCurveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedcurvetype) {
		return
	}

	elevatedcurvetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedcurvetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension := range elevatedcurvetype.Extension {
		stage.UnstageBranch(_a_extension)
	}

}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedpointpropertytype) {
		return
	}

	elevatedpointpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedpointpropertytype.ElevatedPoint != nil {
		stage.UnstageBranch(elevatedpointpropertytype.ElevatedPoint)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedpointtype *ElevatedPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedpointtype) {
		return
	}

	elevatedpointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedpointtype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_1 := range elevatedpointtype.Extension {
		stage.UnstageBranch(_a_extension_1)
	}

}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedsurfacepropertytype) {
		return
	}

	elevatedsurfacepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if elevatedsurfacepropertytype.ElevatedSurface != nil {
		stage.UnstageBranch(elevatedsurfacepropertytype.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elevatedsurfacetype *ElevatedSurfaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elevatedsurfacetype) {
		return
	}

	elevatedsurfacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range elevatedsurfacetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_2 := range elevatedsurfacetype.Extension {
		stage.UnstageBranch(_a_extension_2)
	}

}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(groundlightsystempropertytype) {
		return
	}

	groundlightsystempropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementpropertytype *LightElementPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lightelementpropertytype) {
		return
	}

	lightelementpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementpropertytype.LightElement != nil {
		stage.UnstageBranch(lightelementpropertytype.LightElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lightelementstatuspropertytype) {
		return
	}

	lightelementstatuspropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementstatuspropertytype.LightElementStatus != nil {
		stage.UnstageBranch(lightelementstatuspropertytype.LightElementStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lightelementstatustype *LightElementStatusType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lightelementstatustype) {
		return
	}

	lightelementstatustype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range lightelementstatustype.TimeInterval {
		stage.UnstageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range lightelementstatustype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range lightelementstatustype.SpecialDateAuthority {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _a_extension_11 := range lightelementstatustype.Extension {
		stage.UnstageBranch(_a_extension_11)
	}

}

func (lightelementtype *LightElementType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lightelementtype) {
		return
	}

	lightelementtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lightelementtype.Location != nil {
		stage.UnstageBranch(lightelementtype.Location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range lightelementtype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _lightelementstatuspropertytype := range lightelementtype.Availability {
		stage.UnstageBranch(_lightelementstatuspropertytype)
	}
	for _, _a_extension_10 := range lightelementtype.Extension {
		stage.UnstageBranch(_a_extension_10)
	}

}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linguisticnotepropertytype) {
		return
	}

	linguisticnotepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linguisticnotepropertytype.LinguisticNote != nil {
		stage.UnstageBranch(linguisticnotepropertytype.LinguisticNote)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linguisticnotetype *LinguisticNoteType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linguisticnotetype) {
		return
	}

	linguisticnotetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_extension_4 := range linguisticnotetype.Extension {
		stage.UnstageBranch(_a_extension_4)
	}

}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(markerbeaconpropertytype) {
		return
	}

	markerbeaconpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(navaidequipmentpropertytype) {
		return
	}

	navaidequipmentpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notepropertytype *NotePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notepropertytype) {
		return
	}

	notepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notepropertytype.Note != nil {
		stage.UnstageBranch(notepropertytype.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetype *NoteType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notetype) {
		return
	}

	notetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
		stage.UnstageBranch(_linguisticnotepropertytype)
	}
	for _, _a_extension_5 := range notetype.Extension {
		stage.UnstageBranch(_a_extension_5)
	}

}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(obstacleareatimeslicepropertytype) {
		return
	}

	obstacleareatimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
		stage.UnstageBranch(obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(obstacleareatimeslicetype) {
		return
	}

	obstacleareatimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
		stage.UnstageBranch(obstacleareatimeslicetype.Reference_ownerOrganisation)
	}
	if obstacleareatimeslicetype.Reference_ownerRunway != nil {
		stage.UnstageBranch(obstacleareatimeslicetype.Reference_ownerRunway)
	}
	if obstacleareatimeslicetype.Reference_ownerAirport != nil {
		stage.UnstageBranch(obstacleareatimeslicetype.Reference_ownerAirport)
	}
	if obstacleareatimeslicetype.SurfaceExtent != nil {
		stage.UnstageBranch(obstacleareatimeslicetype.SurfaceExtent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepropertytype := range obstacleareatimeslicetype.Obstacle {
		stage.UnstageBranch(_verticalstructurepropertytype)
	}
	for _, _notepropertytype := range obstacleareatimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_6 := range obstacleareatimeslicetype.Extension {
		stage.UnstageBranch(_a_extension_6)
	}

}

func (obstacleareatype *ObstacleAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(obstacleareatype) {
		return
	}

	obstacleareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _obstacleareatimeslicepropertytype := range obstacleareatype.TimeSlice {
		stage.UnstageBranch(_obstacleareatimeslicepropertytype)
	}

}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(organisationauthoritypropertytype) {
		return
	}

	organisationauthoritypropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (passengerservicepropertytype *PassengerServicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(passengerservicepropertytype) {
		return
	}

	passengerservicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pointtype) {
		return
	}

	pointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range pointtype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}

}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(runwaydirectionpropertytype) {
		return
	}

	runwaydirectionpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (servicepropertytype *ServicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(servicepropertytype) {
		return
	}

	servicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(specialnavigationstationpropertytype) {
		return
	}

	specialnavigationstationpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (surfacepropertytype *SurfacePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(surfacepropertytype) {
		return
	}

	surfacepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if surfacepropertytype.Surface != nil {
		stage.UnstageBranch(surfacepropertytype.Surface)
	}
	if surfacepropertytype.ElevatedSurface != nil {
		stage.UnstageBranch(surfacepropertytype.ElevatedSurface)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (surfacetype *SurfaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(surfacetype) {
		return
	}

	surfacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range surfacetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}

}

func (timesheetpropertytype *TimesheetPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timesheetpropertytype) {
		return
	}

	timesheetpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timesheetpropertytype.Timesheet != nil {
		stage.UnstageBranch(timesheetpropertytype.Timesheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timesheettype *TimesheetType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timesheettype) {
		return
	}

	timesheettype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range timesheettype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_12 := range timesheettype.Extension {
		stage.UnstageBranch(_a_extension_12)
	}

}

func (unitpropertytype *UnitPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(unitpropertytype) {
		return
	}

	unitpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructurelightingstatuspropertytype) {
		return
	}

	verticalstructurelightingstatuspropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
		stage.UnstageBranch(verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructurelightingstatustype) {
		return
	}

	verticalstructurelightingstatustype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructurelightingstatustype.TimeInterval {
		stage.UnstageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range verticalstructurelightingstatustype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructurelightingstatustype.SpecialDateAuthority {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _a_extension_8 := range verticalstructurelightingstatustype.Extension {
		stage.UnstageBranch(_a_extension_8)
	}

}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructurepartpropertytype) {
		return
	}

	verticalstructurepartpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructurepartpropertytype.VerticalStructurePart != nil {
		stage.UnstageBranch(verticalstructurepartpropertytype.VerticalStructurePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructureparttype *VerticalStructurePartType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructureparttype) {
		return
	}

	verticalstructureparttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
		stage.UnstageBranch(verticalstructureparttype.HorizontalProjection_surfaceExtent)
	}
	if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
		stage.UnstageBranch(verticalstructureparttype.HorizontalProjection_linearExtent)
	}
	if verticalstructureparttype.HorizontalProjection_location != nil {
		stage.UnstageBranch(verticalstructureparttype.HorizontalProjection_location)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range verticalstructureparttype.TimeInterval {
		stage.UnstageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range verticalstructureparttype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructureparttype.SpecialDateAuthority {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _lightelementpropertytype := range verticalstructureparttype.Lighting {
		stage.UnstageBranch(_lightelementpropertytype)
	}
	for _, _a_extension_9 := range verticalstructureparttype.Extension {
		stage.UnstageBranch(_a_extension_9)
	}

}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructurepropertytype) {
		return
	}

	verticalstructurepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructuretimeslicepropertytype) {
		return
	}

	verticalstructuretimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
		stage.UnstageBranch(verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructuretimeslicetype) {
		return
	}

	verticalstructuretimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if verticalstructuretimeslicetype.Marker != nil {
		stage.UnstageBranch(verticalstructuretimeslicetype.Marker)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructurepartpropertytype := range verticalstructuretimeslicetype.Part {
		stage.UnstageBranch(_verticalstructurepartpropertytype)
	}
	for _, _passengerservicepropertytype := range verticalstructuretimeslicetype.HostedPassengerService {
		stage.UnstageBranch(_passengerservicepropertytype)
	}
	for _, _groundlightsystempropertytype := range verticalstructuretimeslicetype.SupportedGroundLight {
		stage.UnstageBranch(_groundlightsystempropertytype)
	}
	for _, _navaidequipmentpropertytype := range verticalstructuretimeslicetype.HostedNavaidEquipment {
		stage.UnstageBranch(_navaidequipmentpropertytype)
	}
	for _, _specialnavigationstationpropertytype := range verticalstructuretimeslicetype.HostedSpecialNavStation {
		stage.UnstageBranch(_specialnavigationstationpropertytype)
	}
	for _, _unitpropertytype := range verticalstructuretimeslicetype.HostedUnit {
		stage.UnstageBranch(_unitpropertytype)
	}
	for _, _organisationauthoritypropertytype := range verticalstructuretimeslicetype.HostedOrganisation {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _servicepropertytype := range verticalstructuretimeslicetype.SupportedService {
		stage.UnstageBranch(_servicepropertytype)
	}
	for _, _notepropertytype := range verticalstructuretimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _verticalstructurelightingstatuspropertytype := range verticalstructuretimeslicetype.LightingAvailability {
		stage.UnstageBranch(_verticalstructurelightingstatuspropertytype)
	}
	for _, _a_extension_7 := range verticalstructuretimeslicetype.Extension {
		stage.UnstageBranch(_a_extension_7)
	}

}

func (verticalstructuretype *VerticalStructureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticalstructuretype) {
		return
	}

	verticalstructuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _verticalstructuretimeslicepropertytype := range verticalstructuretype.TimeSlice {
		stage.UnstageBranch(_verticalstructuretimeslicepropertytype)
	}

}

// insertion point for pointer reconstruction from references
func (reference *A_extension) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_10) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_10) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_11) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_11) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_12) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_12) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_2) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_3) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_3) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_4) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_4) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_5) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_5) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_6) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_6) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_7) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_7) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_8) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_8) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_9) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_9) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractPropertiesWithScheduleType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractPropertiesWithScheduleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AeronauticalGroundLightTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AeronauticalGroundLightTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AeronauticalGroundLightTimeSlice, stage.AeronauticalGroundLightTimeSliceTypes_reference, instance.AeronauticalGroundLightTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *AeronauticalGroundLightTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AeronauticalGroundLightTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StructureBeacon, stage.VerticalStructurePropertyTypes_reference, instance.StructureBeacon)
	__gong__reconstructPointer(&reference.AerodromeBeacon, stage.AirportHeliportPropertyTypes_reference, instance.AerodromeBeacon)
	__gong__reconstructPointer(&reference.Location, stage.ElevatedPointPropertyTypes_reference, instance.Location)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_3s_reference, instance.Extension)
}

func (reference *AeronauticalGroundLightType) GongReconstructPointersFromReferences(stage *Stage, instance *AeronauticalGroundLightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.AeronauticalGroundLightTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

func (reference *AirportHeliportPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirportHeliportPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CurveType) GongReconstructPointersFromReferences(stage *Stage, instance *CurveType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
}

func (reference *ElevatedCurvePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedCurvePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ElevatedCurve, stage.ElevatedCurveTypes_reference, instance.ElevatedCurve)
	// insertion point for slice of pointers field
}

func (reference *ElevatedCurveType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedCurveType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extensions_reference, instance.Extension)
}

func (reference *ElevatedPointPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedPointPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ElevatedPoint, stage.ElevatedPointTypes_reference, instance.ElevatedPoint)
	// insertion point for slice of pointers field
}

func (reference *ElevatedPointType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedPointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_1s_reference, instance.Extension)
}

func (reference *ElevatedSurfacePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedSurfacePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ElevatedSurface, stage.ElevatedSurfaceTypes_reference, instance.ElevatedSurface)
	// insertion point for slice of pointers field
}

func (reference *ElevatedSurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedSurfaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_2s_reference, instance.Extension)
}

func (reference *GroundLightSystemPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *GroundLightSystemPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LightElementPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *LightElementPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LightElement, stage.LightElementTypes_reference, instance.LightElement)
	// insertion point for slice of pointers field
}

func (reference *LightElementStatusPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *LightElementStatusPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LightElementStatus, stage.LightElementStatusTypes_reference, instance.LightElementStatus)
	// insertion point for slice of pointers field
}

func (reference *LightElementStatusType) GongReconstructPointersFromReferences(stage *Stage, instance *LightElementStatusType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeInterval, stage.TimesheetPropertyTypes_reference, instance.TimeInterval)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_reference, instance.SpecialDateAuthority)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_11s_reference, instance.Extension)
}

func (reference *LightElementType) GongReconstructPointersFromReferences(stage *Stage, instance *LightElementType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Location, stage.ElevatedPointPropertyTypes_reference, instance.Location)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Availability, stage.LightElementStatusPropertyTypes_reference, instance.Availability)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_10s_reference, instance.Extension)
}

func (reference *LinguisticNotePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *LinguisticNotePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LinguisticNote, stage.LinguisticNoteTypes_reference, instance.LinguisticNote)
	// insertion point for slice of pointers field
}

func (reference *LinguisticNoteType) GongReconstructPointersFromReferences(stage *Stage, instance *LinguisticNoteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_4s_reference, instance.Extension)
}

func (reference *MarkerBeaconPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MarkerBeaconPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NavaidEquipmentPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *NavaidEquipmentPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NotePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *NotePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.NoteTypes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *NoteType) GongReconstructPointersFromReferences(stage *Stage, instance *NoteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TranslatedNote, stage.LinguisticNotePropertyTypes_reference, instance.TranslatedNote)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_5s_reference, instance.Extension)
}

func (reference *ObstacleAreaTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ObstacleAreaTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ObstacleAreaTimeSlice, stage.ObstacleAreaTimeSliceTypes_reference, instance.ObstacleAreaTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *ObstacleAreaTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *ObstacleAreaTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Reference_ownerOrganisation, stage.OrganisationAuthorityPropertyTypes_reference, instance.Reference_ownerOrganisation)
	__gong__reconstructPointer(&reference.Reference_ownerRunway, stage.RunwayDirectionPropertyTypes_reference, instance.Reference_ownerRunway)
	__gong__reconstructPointer(&reference.Reference_ownerAirport, stage.AirportHeliportPropertyTypes_reference, instance.Reference_ownerAirport)
	__gong__reconstructPointer(&reference.SurfaceExtent, stage.SurfacePropertyTypes_reference, instance.SurfaceExtent)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Obstacle, stage.VerticalStructurePropertyTypes_reference, instance.Obstacle)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_6s_reference, instance.Extension)
}

func (reference *ObstacleAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *ObstacleAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.ObstacleAreaTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

func (reference *OrganisationAuthorityPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *OrganisationAuthorityPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PassengerServicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *PassengerServicePropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PointType) GongReconstructPointersFromReferences(stage *Stage, instance *PointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
}

func (reference *RunwayDirectionPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *RunwayDirectionPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ServicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ServicePropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SpecialNavigationStationPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *SpecialNavigationStationPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SurfacePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *SurfacePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Surface, stage.SurfaceTypes_reference, instance.Surface)
	__gong__reconstructPointer(&reference.ElevatedSurface, stage.ElevatedSurfaceTypes_reference, instance.ElevatedSurface)
	// insertion point for slice of pointers field
}

func (reference *SurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *SurfaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
}

func (reference *TimesheetPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *TimesheetPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Timesheet, stage.TimesheetTypes_reference, instance.Timesheet)
	// insertion point for slice of pointers field
}

func (reference *TimesheetType) GongReconstructPointersFromReferences(stage *Stage, instance *TimesheetType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_12s_reference, instance.Extension)
}

func (reference *UnitPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *UnitPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VerticalStructureLightingStatusPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructureLightingStatusPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.VerticalStructureLightingStatus, stage.VerticalStructureLightingStatusTypes_reference, instance.VerticalStructureLightingStatus)
	// insertion point for slice of pointers field
}

func (reference *VerticalStructureLightingStatusType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructureLightingStatusType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeInterval, stage.TimesheetPropertyTypes_reference, instance.TimeInterval)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_reference, instance.SpecialDateAuthority)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_8s_reference, instance.Extension)
}

func (reference *VerticalStructurePartPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructurePartPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.VerticalStructurePart, stage.VerticalStructurePartTypes_reference, instance.VerticalStructurePart)
	// insertion point for slice of pointers field
}

func (reference *VerticalStructurePartType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructurePartType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.HorizontalProjection_surfaceExtent, stage.ElevatedSurfacePropertyTypes_reference, instance.HorizontalProjection_surfaceExtent)
	__gong__reconstructPointer(&reference.HorizontalProjection_linearExtent, stage.ElevatedCurvePropertyTypes_reference, instance.HorizontalProjection_linearExtent)
	__gong__reconstructPointer(&reference.HorizontalProjection_location, stage.ElevatedPointPropertyTypes_reference, instance.HorizontalProjection_location)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeInterval, stage.TimesheetPropertyTypes_reference, instance.TimeInterval)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_reference, instance.SpecialDateAuthority)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lighting, stage.LightElementPropertyTypes_reference, instance.Lighting)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_9s_reference, instance.Extension)
}

func (reference *VerticalStructurePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructurePropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VerticalStructureTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructureTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.VerticalStructureTimeSlice, stage.VerticalStructureTimeSliceTypes_reference, instance.VerticalStructureTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *VerticalStructureTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructureTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Marker, stage.MarkerBeaconPropertyTypes_reference, instance.Marker)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part, stage.VerticalStructurePartPropertyTypes_reference, instance.Part)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HostedPassengerService, stage.PassengerServicePropertyTypes_reference, instance.HostedPassengerService)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SupportedGroundLight, stage.GroundLightSystemPropertyTypes_reference, instance.SupportedGroundLight)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HostedNavaidEquipment, stage.NavaidEquipmentPropertyTypes_reference, instance.HostedNavaidEquipment)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HostedSpecialNavStation, stage.SpecialNavigationStationPropertyTypes_reference, instance.HostedSpecialNavStation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HostedUnit, stage.UnitPropertyTypes_reference, instance.HostedUnit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HostedOrganisation, stage.OrganisationAuthorityPropertyTypes_reference, instance.HostedOrganisation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SupportedService, stage.ServicePropertyTypes_reference, instance.SupportedService)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.LightingAvailability, stage.VerticalStructureLightingStatusPropertyTypes_reference, instance.LightingAvailability)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_7s_reference, instance.Extension)
}

func (reference *VerticalStructureType) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalStructureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.VerticalStructureTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

// insertion point for pointer reconstruction from instances
func (reference *A_extension) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_10) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_11) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_12) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_3) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_4) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_5) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_6) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_7) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_8) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_9) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractPropertiesWithScheduleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AeronauticalGroundLightTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AeronauticalGroundLightTimeSlice, stage.AeronauticalGroundLightTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AeronauticalGroundLightTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StructureBeacon, stage.VerticalStructurePropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.AerodromeBeacon, stage.AirportHeliportPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Location, stage.ElevatedPointPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_3s_instance)
}

func (reference *AeronauticalGroundLightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.AeronauticalGroundLightTimeSlicePropertyTypes_instance)
}

func (reference *AirportHeliportPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
}

func (reference *ElevatedCurvePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ElevatedCurve, stage.ElevatedCurveTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ElevatedCurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extensions_instance)
}

func (reference *ElevatedPointPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ElevatedPoint, stage.ElevatedPointTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ElevatedPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_1s_instance)
}

func (reference *ElevatedSurfacePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ElevatedSurface, stage.ElevatedSurfaceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ElevatedSurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_2s_instance)
}

func (reference *GroundLightSystemPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LightElementPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LightElement, stage.LightElementTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *LightElementStatusPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LightElementStatus, stage.LightElementStatusTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *LightElementStatusType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeInterval, stage.TimesheetPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_11s_instance)
}

func (reference *LightElementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Location, stage.ElevatedPointPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Availability, stage.LightElementStatusPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_10s_instance)
}

func (reference *LinguisticNotePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LinguisticNote, stage.LinguisticNoteTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *LinguisticNoteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_4s_instance)
}

func (reference *MarkerBeaconPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NavaidEquipmentPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NotePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.NoteTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TranslatedNote, stage.LinguisticNotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_5s_instance)
}

func (reference *ObstacleAreaTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ObstacleAreaTimeSlice, stage.ObstacleAreaTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ObstacleAreaTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Reference_ownerOrganisation, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Reference_ownerRunway, stage.RunwayDirectionPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Reference_ownerAirport, stage.AirportHeliportPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.SurfaceExtent, stage.SurfacePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Obstacle, stage.VerticalStructurePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_6s_instance)
}

func (reference *ObstacleAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.ObstacleAreaTimeSlicePropertyTypes_instance)
}

func (reference *OrganisationAuthorityPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PassengerServicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
}

func (reference *RunwayDirectionPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ServicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SpecialNavigationStationPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SurfacePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Surface, stage.SurfaceTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ElevatedSurface, stage.ElevatedSurfaceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *SurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
}

func (reference *TimesheetPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Timesheet, stage.TimesheetTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *TimesheetType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_12s_instance)
}

func (reference *UnitPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VerticalStructureLightingStatusPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.VerticalStructureLightingStatus, stage.VerticalStructureLightingStatusTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *VerticalStructureLightingStatusType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeInterval, stage.TimesheetPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_8s_instance)
}

func (reference *VerticalStructurePartPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.VerticalStructurePart, stage.VerticalStructurePartTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *VerticalStructurePartType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.HorizontalProjection_surfaceExtent, stage.ElevatedSurfacePropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.HorizontalProjection_linearExtent, stage.ElevatedCurvePropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.HorizontalProjection_location, stage.ElevatedPointPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeInterval, stage.TimesheetPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lighting, stage.LightElementPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_9s_instance)
}

func (reference *VerticalStructurePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VerticalStructureTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.VerticalStructureTimeSlice, stage.VerticalStructureTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *VerticalStructureTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Marker, stage.MarkerBeaconPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part, stage.VerticalStructurePartPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HostedPassengerService, stage.PassengerServicePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SupportedGroundLight, stage.GroundLightSystemPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HostedNavaidEquipment, stage.NavaidEquipmentPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HostedSpecialNavStation, stage.SpecialNavigationStationPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HostedUnit, stage.UnitPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HostedOrganisation, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SupportedService, stage.ServicePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.LightingAvailability, stage.VerticalStructureLightingStatusPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_7s_instance)
}

func (reference *VerticalStructureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.VerticalStructureTimeSlicePropertyTypes_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension *A_extension) GongDiff(stage *Stage, a_extensionOther *A_extension) (diffs []string) {
	// insertion point for field diffs
	if a_extension.Name != a_extensionOther.Name {
		diffs = append(diffs, a_extension.GongMarshallField(stage, "Name"))
	}
	if a_extension.Owns != a_extensionOther.Owns {
		diffs = append(diffs, a_extension.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_1 *A_extension_1) GongDiff(stage *Stage, a_extension_1Other *A_extension_1) (diffs []string) {
	// insertion point for field diffs
	if a_extension_1.Name != a_extension_1Other.Name {
		diffs = append(diffs, a_extension_1.GongMarshallField(stage, "Name"))
	}
	if a_extension_1.Owns != a_extension_1Other.Owns {
		diffs = append(diffs, a_extension_1.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_10 *A_extension_10) GongDiff(stage *Stage, a_extension_10Other *A_extension_10) (diffs []string) {
	// insertion point for field diffs
	if a_extension_10.Name != a_extension_10Other.Name {
		diffs = append(diffs, a_extension_10.GongMarshallField(stage, "Name"))
	}
	if a_extension_10.Owns != a_extension_10Other.Owns {
		diffs = append(diffs, a_extension_10.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_11 *A_extension_11) GongDiff(stage *Stage, a_extension_11Other *A_extension_11) (diffs []string) {
	// insertion point for field diffs
	if a_extension_11.Name != a_extension_11Other.Name {
		diffs = append(diffs, a_extension_11.GongMarshallField(stage, "Name"))
	}
	if a_extension_11.Owns != a_extension_11Other.Owns {
		diffs = append(diffs, a_extension_11.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_12 *A_extension_12) GongDiff(stage *Stage, a_extension_12Other *A_extension_12) (diffs []string) {
	// insertion point for field diffs
	if a_extension_12.Name != a_extension_12Other.Name {
		diffs = append(diffs, a_extension_12.GongMarshallField(stage, "Name"))
	}
	if a_extension_12.Owns != a_extension_12Other.Owns {
		diffs = append(diffs, a_extension_12.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_2 *A_extension_2) GongDiff(stage *Stage, a_extension_2Other *A_extension_2) (diffs []string) {
	// insertion point for field diffs
	if a_extension_2.Name != a_extension_2Other.Name {
		diffs = append(diffs, a_extension_2.GongMarshallField(stage, "Name"))
	}
	if a_extension_2.Owns != a_extension_2Other.Owns {
		diffs = append(diffs, a_extension_2.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_3 *A_extension_3) GongDiff(stage *Stage, a_extension_3Other *A_extension_3) (diffs []string) {
	// insertion point for field diffs
	if a_extension_3.Name != a_extension_3Other.Name {
		diffs = append(diffs, a_extension_3.GongMarshallField(stage, "Name"))
	}
	if a_extension_3.Owns != a_extension_3Other.Owns {
		diffs = append(diffs, a_extension_3.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_4 *A_extension_4) GongDiff(stage *Stage, a_extension_4Other *A_extension_4) (diffs []string) {
	// insertion point for field diffs
	if a_extension_4.Name != a_extension_4Other.Name {
		diffs = append(diffs, a_extension_4.GongMarshallField(stage, "Name"))
	}
	if a_extension_4.Owns != a_extension_4Other.Owns {
		diffs = append(diffs, a_extension_4.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_5 *A_extension_5) GongDiff(stage *Stage, a_extension_5Other *A_extension_5) (diffs []string) {
	// insertion point for field diffs
	if a_extension_5.Name != a_extension_5Other.Name {
		diffs = append(diffs, a_extension_5.GongMarshallField(stage, "Name"))
	}
	if a_extension_5.Owns != a_extension_5Other.Owns {
		diffs = append(diffs, a_extension_5.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_6 *A_extension_6) GongDiff(stage *Stage, a_extension_6Other *A_extension_6) (diffs []string) {
	// insertion point for field diffs
	if a_extension_6.Name != a_extension_6Other.Name {
		diffs = append(diffs, a_extension_6.GongMarshallField(stage, "Name"))
	}
	if a_extension_6.Owns != a_extension_6Other.Owns {
		diffs = append(diffs, a_extension_6.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_7 *A_extension_7) GongDiff(stage *Stage, a_extension_7Other *A_extension_7) (diffs []string) {
	// insertion point for field diffs
	if a_extension_7.Name != a_extension_7Other.Name {
		diffs = append(diffs, a_extension_7.GongMarshallField(stage, "Name"))
	}
	if a_extension_7.Owns != a_extension_7Other.Owns {
		diffs = append(diffs, a_extension_7.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_8 *A_extension_8) GongDiff(stage *Stage, a_extension_8Other *A_extension_8) (diffs []string) {
	// insertion point for field diffs
	if a_extension_8.Name != a_extension_8Other.Name {
		diffs = append(diffs, a_extension_8.GongMarshallField(stage, "Name"))
	}
	if a_extension_8.Owns != a_extension_8Other.Owns {
		diffs = append(diffs, a_extension_8.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_9 *A_extension_9) GongDiff(stage *Stage, a_extension_9Other *A_extension_9) (diffs []string) {
	// insertion point for field diffs
	if a_extension_9.Name != a_extension_9Other.Name {
		diffs = append(diffs, a_extension_9.GongMarshallField(stage, "Name"))
	}
	if a_extension_9.Owns != a_extension_9Other.Owns {
		diffs = append(diffs, a_extension_9.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongDiff(stage *Stage, abstractpropertieswithscheduletypeOther *AbstractPropertiesWithScheduleType) (diffs []string) {
	// insertion point for field diffs
	if abstractpropertieswithscheduletype.Name != abstractpropertieswithscheduletypeOther.Name {
		diffs = append(diffs, abstractpropertieswithscheduletype.GongMarshallField(stage, "Name"))
	}
	if abstractpropertieswithscheduletype.Id != abstractpropertieswithscheduletypeOther.Id {
		diffs = append(diffs, abstractpropertieswithscheduletype.GongMarshallField(stage, "Id"))
	}
	if abstractpropertieswithscheduletype.Description != abstractpropertieswithscheduletypeOther.Description {
		diffs = append(diffs, abstractpropertieswithscheduletype.GongMarshallField(stage, "Description"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongDiff(stage *Stage, aeronauticalgroundlighttimeslicepropertytypeOther *AeronauticalGroundLightTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if aeronauticalgroundlighttimeslicepropertytype.Name != aeronauticalgroundlighttimeslicepropertytypeOther.Name {
		diffs = append(diffs, aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if aeronauticalgroundlighttimeslicepropertytype.Owns != aeronauticalgroundlighttimeslicepropertytypeOther.Owns {
		diffs = append(diffs, aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != aeronauticalgroundlighttimeslicepropertytypeOther.AeronauticalGroundLightTimeSlice {
		diffs = append(diffs, aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "AeronauticalGroundLightTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongDiff(stage *Stage, aeronauticalgroundlighttimeslicetypeOther *AeronauticalGroundLightTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if aeronauticalgroundlighttimeslicetype.Name != aeronauticalgroundlighttimeslicetypeOther.Name {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Name"))
	}
	if aeronauticalgroundlighttimeslicetype.Id != aeronauticalgroundlighttimeslicetypeOther.Id {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Id"))
	}
	if aeronauticalgroundlighttimeslicetype.Description != aeronauticalgroundlighttimeslicetypeOther.Description {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Description"))
	}
	if aeronauticalgroundlighttimeslicetype.Interpretation != aeronauticalgroundlighttimeslicetypeOther.Interpretation {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if aeronauticalgroundlighttimeslicetype.SequenceNumber != aeronauticalgroundlighttimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if aeronauticalgroundlighttimeslicetype.CorrectionNumber != aeronauticalgroundlighttimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if aeronauticalgroundlighttimeslicetype.StructureBeacon != aeronauticalgroundlighttimeslicetypeOther.StructureBeacon {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "StructureBeacon"))
	}
	if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != aeronauticalgroundlighttimeslicetypeOther.AerodromeBeacon {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "AerodromeBeacon"))
	}
	if aeronauticalgroundlighttimeslicetype.Location != aeronauticalgroundlighttimeslicetypeOther.Location {
		diffs = append(diffs, aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Location"))
	}
	if ops := __gong__diffSliceOfPointers(stage, aeronauticalgroundlighttimeslicetype, "Annotation", aeronauticalgroundlighttimeslicetypeOther.Annotation, aeronauticalgroundlighttimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, aeronauticalgroundlighttimeslicetype, "Extension", aeronauticalgroundlighttimeslicetypeOther.Extension, aeronauticalgroundlighttimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongDiff(stage *Stage, aeronauticalgroundlighttypeOther *AeronauticalGroundLightType) (diffs []string) {
	// insertion point for field diffs
	if aeronauticalgroundlighttype.Name != aeronauticalgroundlighttypeOther.Name {
		diffs = append(diffs, aeronauticalgroundlighttype.GongMarshallField(stage, "Name"))
	}
	if aeronauticalgroundlighttype.Id != aeronauticalgroundlighttypeOther.Id {
		diffs = append(diffs, aeronauticalgroundlighttype.GongMarshallField(stage, "Id"))
	}
	if aeronauticalgroundlighttype.Description != aeronauticalgroundlighttypeOther.Description {
		diffs = append(diffs, aeronauticalgroundlighttype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, aeronauticalgroundlighttype, "TimeSlice", aeronauticalgroundlighttypeOther.TimeSlice, aeronauticalgroundlighttype.TimeSlice); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airportheliportpropertytype *AirportHeliportPropertyType) GongDiff(stage *Stage, airportheliportpropertytypeOther *AirportHeliportPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airportheliportpropertytype.Name != airportheliportpropertytypeOther.Name {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Name"))
	}
	if airportheliportpropertytype.Owns != airportheliportpropertytypeOther.Owns {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Owns"))
	}
	if airportheliportpropertytype.NilReason != airportheliportpropertytypeOther.NilReason {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airportheliportpropertytype.Type_XSD != airportheliportpropertytypeOther.Type_XSD {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if airportheliportpropertytype.Href != airportheliportpropertytypeOther.Href {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Href"))
	}
	if airportheliportpropertytype.Role != airportheliportpropertytypeOther.Role {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Role"))
	}
	if airportheliportpropertytype.Arcrole != airportheliportpropertytypeOther.Arcrole {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if airportheliportpropertytype.Title != airportheliportpropertytypeOther.Title {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Title"))
	}
	if airportheliportpropertytype.Show != airportheliportpropertytypeOther.Show {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Show"))
	}
	if airportheliportpropertytype.Actuate != airportheliportpropertytypeOther.Actuate {
		diffs = append(diffs, airportheliportpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curvetype *CurveType) GongDiff(stage *Stage, curvetypeOther *CurveType) (diffs []string) {
	// insertion point for field diffs
	if curvetype.Name != curvetypeOther.Name {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Name"))
	}
	if curvetype.Id != curvetypeOther.Id {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Id"))
	}
	if curvetype.Description != curvetypeOther.Description {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "Description"))
	}
	if curvetype.SrsName != curvetypeOther.SrsName {
		diffs = append(diffs, curvetype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, curvetype, "Annotation", curvetypeOther.Annotation, curvetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongDiff(stage *Stage, elevatedcurvepropertytypeOther *ElevatedCurvePropertyType) (diffs []string) {
	// insertion point for field diffs
	if elevatedcurvepropertytype.Name != elevatedcurvepropertytypeOther.Name {
		diffs = append(diffs, elevatedcurvepropertytype.GongMarshallField(stage, "Name"))
	}
	if elevatedcurvepropertytype.NilReason != elevatedcurvepropertytypeOther.NilReason {
		diffs = append(diffs, elevatedcurvepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if elevatedcurvepropertytype.ElevatedCurve != elevatedcurvepropertytypeOther.ElevatedCurve {
		diffs = append(diffs, elevatedcurvepropertytype.GongMarshallField(stage, "ElevatedCurve"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedcurvetype *ElevatedCurveType) GongDiff(stage *Stage, elevatedcurvetypeOther *ElevatedCurveType) (diffs []string) {
	// insertion point for field diffs
	if elevatedcurvetype.Name != elevatedcurvetypeOther.Name {
		diffs = append(diffs, elevatedcurvetype.GongMarshallField(stage, "Name"))
	}
	if elevatedcurvetype.Id != elevatedcurvetypeOther.Id {
		diffs = append(diffs, elevatedcurvetype.GongMarshallField(stage, "Id"))
	}
	if elevatedcurvetype.Description != elevatedcurvetypeOther.Description {
		diffs = append(diffs, elevatedcurvetype.GongMarshallField(stage, "Description"))
	}
	if elevatedcurvetype.SrsName != elevatedcurvetypeOther.SrsName {
		diffs = append(diffs, elevatedcurvetype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedcurvetype, "Annotation", elevatedcurvetypeOther.Annotation, elevatedcurvetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedcurvetype, "Extension", elevatedcurvetypeOther.Extension, elevatedcurvetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongDiff(stage *Stage, elevatedpointpropertytypeOther *ElevatedPointPropertyType) (diffs []string) {
	// insertion point for field diffs
	if elevatedpointpropertytype.Name != elevatedpointpropertytypeOther.Name {
		diffs = append(diffs, elevatedpointpropertytype.GongMarshallField(stage, "Name"))
	}
	if elevatedpointpropertytype.NilReason != elevatedpointpropertytypeOther.NilReason {
		diffs = append(diffs, elevatedpointpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if elevatedpointpropertytype.ElevatedPoint != elevatedpointpropertytypeOther.ElevatedPoint {
		diffs = append(diffs, elevatedpointpropertytype.GongMarshallField(stage, "ElevatedPoint"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedpointtype *ElevatedPointType) GongDiff(stage *Stage, elevatedpointtypeOther *ElevatedPointType) (diffs []string) {
	// insertion point for field diffs
	if elevatedpointtype.Name != elevatedpointtypeOther.Name {
		diffs = append(diffs, elevatedpointtype.GongMarshallField(stage, "Name"))
	}
	if elevatedpointtype.Id != elevatedpointtypeOther.Id {
		diffs = append(diffs, elevatedpointtype.GongMarshallField(stage, "Id"))
	}
	if elevatedpointtype.Description != elevatedpointtypeOther.Description {
		diffs = append(diffs, elevatedpointtype.GongMarshallField(stage, "Description"))
	}
	if elevatedpointtype.SrsName != elevatedpointtypeOther.SrsName {
		diffs = append(diffs, elevatedpointtype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedpointtype, "Annotation", elevatedpointtypeOther.Annotation, elevatedpointtype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedpointtype, "Extension", elevatedpointtypeOther.Extension, elevatedpointtype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongDiff(stage *Stage, elevatedsurfacepropertytypeOther *ElevatedSurfacePropertyType) (diffs []string) {
	// insertion point for field diffs
	if elevatedsurfacepropertytype.Name != elevatedsurfacepropertytypeOther.Name {
		diffs = append(diffs, elevatedsurfacepropertytype.GongMarshallField(stage, "Name"))
	}
	if elevatedsurfacepropertytype.NilReason != elevatedsurfacepropertytypeOther.NilReason {
		diffs = append(diffs, elevatedsurfacepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if elevatedsurfacepropertytype.ElevatedSurface != elevatedsurfacepropertytypeOther.ElevatedSurface {
		diffs = append(diffs, elevatedsurfacepropertytype.GongMarshallField(stage, "ElevatedSurface"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elevatedsurfacetype *ElevatedSurfaceType) GongDiff(stage *Stage, elevatedsurfacetypeOther *ElevatedSurfaceType) (diffs []string) {
	// insertion point for field diffs
	if elevatedsurfacetype.Name != elevatedsurfacetypeOther.Name {
		diffs = append(diffs, elevatedsurfacetype.GongMarshallField(stage, "Name"))
	}
	if elevatedsurfacetype.Id != elevatedsurfacetypeOther.Id {
		diffs = append(diffs, elevatedsurfacetype.GongMarshallField(stage, "Id"))
	}
	if elevatedsurfacetype.Description != elevatedsurfacetypeOther.Description {
		diffs = append(diffs, elevatedsurfacetype.GongMarshallField(stage, "Description"))
	}
	if elevatedsurfacetype.SrsName != elevatedsurfacetypeOther.SrsName {
		diffs = append(diffs, elevatedsurfacetype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedsurfacetype, "Annotation", elevatedsurfacetypeOther.Annotation, elevatedsurfacetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, elevatedsurfacetype, "Extension", elevatedsurfacetypeOther.Extension, elevatedsurfacetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongDiff(stage *Stage, groundlightsystempropertytypeOther *GroundLightSystemPropertyType) (diffs []string) {
	// insertion point for field diffs
	if groundlightsystempropertytype.Name != groundlightsystempropertytypeOther.Name {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Name"))
	}
	if groundlightsystempropertytype.Owns != groundlightsystempropertytypeOther.Owns {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Owns"))
	}
	if groundlightsystempropertytype.NilReason != groundlightsystempropertytypeOther.NilReason {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "NilReason"))
	}
	if groundlightsystempropertytype.Type_XSD != groundlightsystempropertytypeOther.Type_XSD {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if groundlightsystempropertytype.Href != groundlightsystempropertytypeOther.Href {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Href"))
	}
	if groundlightsystempropertytype.Role != groundlightsystempropertytypeOther.Role {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Role"))
	}
	if groundlightsystempropertytype.Arcrole != groundlightsystempropertytypeOther.Arcrole {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if groundlightsystempropertytype.Title != groundlightsystempropertytypeOther.Title {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Title"))
	}
	if groundlightsystempropertytype.Show != groundlightsystempropertytypeOther.Show {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Show"))
	}
	if groundlightsystempropertytype.Actuate != groundlightsystempropertytypeOther.Actuate {
		diffs = append(diffs, groundlightsystempropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lightelementpropertytype *LightElementPropertyType) GongDiff(stage *Stage, lightelementpropertytypeOther *LightElementPropertyType) (diffs []string) {
	// insertion point for field diffs
	if lightelementpropertytype.Name != lightelementpropertytypeOther.Name {
		diffs = append(diffs, lightelementpropertytype.GongMarshallField(stage, "Name"))
	}
	if lightelementpropertytype.NilReason != lightelementpropertytypeOther.NilReason {
		diffs = append(diffs, lightelementpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if lightelementpropertytype.LightElement != lightelementpropertytypeOther.LightElement {
		diffs = append(diffs, lightelementpropertytype.GongMarshallField(stage, "LightElement"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongDiff(stage *Stage, lightelementstatuspropertytypeOther *LightElementStatusPropertyType) (diffs []string) {
	// insertion point for field diffs
	if lightelementstatuspropertytype.Name != lightelementstatuspropertytypeOther.Name {
		diffs = append(diffs, lightelementstatuspropertytype.GongMarshallField(stage, "Name"))
	}
	if lightelementstatuspropertytype.NilReason != lightelementstatuspropertytypeOther.NilReason {
		diffs = append(diffs, lightelementstatuspropertytype.GongMarshallField(stage, "NilReason"))
	}
	if lightelementstatuspropertytype.LightElementStatus != lightelementstatuspropertytypeOther.LightElementStatus {
		diffs = append(diffs, lightelementstatuspropertytype.GongMarshallField(stage, "LightElementStatus"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lightelementstatustype *LightElementStatusType) GongDiff(stage *Stage, lightelementstatustypeOther *LightElementStatusType) (diffs []string) {
	// insertion point for field diffs
	if lightelementstatustype.Name != lightelementstatustypeOther.Name {
		diffs = append(diffs, lightelementstatustype.GongMarshallField(stage, "Name"))
	}
	if lightelementstatustype.Id != lightelementstatustypeOther.Id {
		diffs = append(diffs, lightelementstatustype.GongMarshallField(stage, "Id"))
	}
	if lightelementstatustype.Description != lightelementstatustypeOther.Description {
		diffs = append(diffs, lightelementstatustype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementstatustype, "TimeInterval", lightelementstatustypeOther.TimeInterval, lightelementstatustype.TimeInterval); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementstatustype, "Annotation", lightelementstatustypeOther.Annotation, lightelementstatustype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementstatustype, "SpecialDateAuthority", lightelementstatustypeOther.SpecialDateAuthority, lightelementstatustype.SpecialDateAuthority); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementstatustype, "Extension", lightelementstatustypeOther.Extension, lightelementstatustype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lightelementtype *LightElementType) GongDiff(stage *Stage, lightelementtypeOther *LightElementType) (diffs []string) {
	// insertion point for field diffs
	if lightelementtype.Name != lightelementtypeOther.Name {
		diffs = append(diffs, lightelementtype.GongMarshallField(stage, "Name"))
	}
	if lightelementtype.Id != lightelementtypeOther.Id {
		diffs = append(diffs, lightelementtype.GongMarshallField(stage, "Id"))
	}
	if lightelementtype.Description != lightelementtypeOther.Description {
		diffs = append(diffs, lightelementtype.GongMarshallField(stage, "Description"))
	}
	if lightelementtype.Location != lightelementtypeOther.Location {
		diffs = append(diffs, lightelementtype.GongMarshallField(stage, "Location"))
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementtype, "Annotation", lightelementtypeOther.Annotation, lightelementtype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementtype, "Availability", lightelementtypeOther.Availability, lightelementtype.Availability); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, lightelementtype, "Extension", lightelementtypeOther.Extension, lightelementtype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongDiff(stage *Stage, linguisticnotepropertytypeOther *LinguisticNotePropertyType) (diffs []string) {
	// insertion point for field diffs
	if linguisticnotepropertytype.Name != linguisticnotepropertytypeOther.Name {
		diffs = append(diffs, linguisticnotepropertytype.GongMarshallField(stage, "Name"))
	}
	if linguisticnotepropertytype.NilReason != linguisticnotepropertytypeOther.NilReason {
		diffs = append(diffs, linguisticnotepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if linguisticnotepropertytype.LinguisticNote != linguisticnotepropertytypeOther.LinguisticNote {
		diffs = append(diffs, linguisticnotepropertytype.GongMarshallField(stage, "LinguisticNote"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linguisticnotetype *LinguisticNoteType) GongDiff(stage *Stage, linguisticnotetypeOther *LinguisticNoteType) (diffs []string) {
	// insertion point for field diffs
	if linguisticnotetype.Name != linguisticnotetypeOther.Name {
		diffs = append(diffs, linguisticnotetype.GongMarshallField(stage, "Name"))
	}
	if linguisticnotetype.Id != linguisticnotetypeOther.Id {
		diffs = append(diffs, linguisticnotetype.GongMarshallField(stage, "Id"))
	}
	if linguisticnotetype.Description != linguisticnotetypeOther.Description {
		diffs = append(diffs, linguisticnotetype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, linguisticnotetype, "Extension", linguisticnotetypeOther.Extension, linguisticnotetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongDiff(stage *Stage, markerbeaconpropertytypeOther *MarkerBeaconPropertyType) (diffs []string) {
	// insertion point for field diffs
	if markerbeaconpropertytype.Name != markerbeaconpropertytypeOther.Name {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Name"))
	}
	if markerbeaconpropertytype.Owns != markerbeaconpropertytypeOther.Owns {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Owns"))
	}
	if markerbeaconpropertytype.NilReason != markerbeaconpropertytypeOther.NilReason {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if markerbeaconpropertytype.Type_XSD != markerbeaconpropertytypeOther.Type_XSD {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if markerbeaconpropertytype.Href != markerbeaconpropertytypeOther.Href {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Href"))
	}
	if markerbeaconpropertytype.Role != markerbeaconpropertytypeOther.Role {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Role"))
	}
	if markerbeaconpropertytype.Arcrole != markerbeaconpropertytypeOther.Arcrole {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if markerbeaconpropertytype.Title != markerbeaconpropertytypeOther.Title {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Title"))
	}
	if markerbeaconpropertytype.Show != markerbeaconpropertytypeOther.Show {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Show"))
	}
	if markerbeaconpropertytype.Actuate != markerbeaconpropertytypeOther.Actuate {
		diffs = append(diffs, markerbeaconpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongDiff(stage *Stage, navaidequipmentpropertytypeOther *NavaidEquipmentPropertyType) (diffs []string) {
	// insertion point for field diffs
	if navaidequipmentpropertytype.Name != navaidequipmentpropertytypeOther.Name {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Name"))
	}
	if navaidequipmentpropertytype.Owns != navaidequipmentpropertytypeOther.Owns {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Owns"))
	}
	if navaidequipmentpropertytype.NilReason != navaidequipmentpropertytypeOther.NilReason {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if navaidequipmentpropertytype.Type_XSD != navaidequipmentpropertytypeOther.Type_XSD {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if navaidequipmentpropertytype.Href != navaidequipmentpropertytypeOther.Href {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Href"))
	}
	if navaidequipmentpropertytype.Role != navaidequipmentpropertytypeOther.Role {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Role"))
	}
	if navaidequipmentpropertytype.Arcrole != navaidequipmentpropertytypeOther.Arcrole {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if navaidequipmentpropertytype.Title != navaidequipmentpropertytypeOther.Title {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Title"))
	}
	if navaidequipmentpropertytype.Show != navaidequipmentpropertytypeOther.Show {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Show"))
	}
	if navaidequipmentpropertytype.Actuate != navaidequipmentpropertytypeOther.Actuate {
		diffs = append(diffs, navaidequipmentpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notepropertytype *NotePropertyType) GongDiff(stage *Stage, notepropertytypeOther *NotePropertyType) (diffs []string) {
	// insertion point for field diffs
	if notepropertytype.Name != notepropertytypeOther.Name {
		diffs = append(diffs, notepropertytype.GongMarshallField(stage, "Name"))
	}
	if notepropertytype.NilReason != notepropertytypeOther.NilReason {
		diffs = append(diffs, notepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if notepropertytype.Note != notepropertytypeOther.Note {
		diffs = append(diffs, notepropertytype.GongMarshallField(stage, "Note"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetype *NoteType) GongDiff(stage *Stage, notetypeOther *NoteType) (diffs []string) {
	// insertion point for field diffs
	if notetype.Name != notetypeOther.Name {
		diffs = append(diffs, notetype.GongMarshallField(stage, "Name"))
	}
	if notetype.Id != notetypeOther.Id {
		diffs = append(diffs, notetype.GongMarshallField(stage, "Id"))
	}
	if notetype.Description != notetypeOther.Description {
		diffs = append(diffs, notetype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notetype, "TranslatedNote", notetypeOther.TranslatedNote, notetype.TranslatedNote); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notetype, "Extension", notetypeOther.Extension, notetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongDiff(stage *Stage, obstacleareatimeslicepropertytypeOther *ObstacleAreaTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if obstacleareatimeslicepropertytype.Name != obstacleareatimeslicepropertytypeOther.Name {
		diffs = append(diffs, obstacleareatimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if obstacleareatimeslicepropertytype.Owns != obstacleareatimeslicepropertytypeOther.Owns {
		diffs = append(diffs, obstacleareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != obstacleareatimeslicepropertytypeOther.ObstacleAreaTimeSlice {
		diffs = append(diffs, obstacleareatimeslicepropertytype.GongMarshallField(stage, "ObstacleAreaTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongDiff(stage *Stage, obstacleareatimeslicetypeOther *ObstacleAreaTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if obstacleareatimeslicetype.Name != obstacleareatimeslicetypeOther.Name {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Name"))
	}
	if obstacleareatimeslicetype.Id != obstacleareatimeslicetypeOther.Id {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Id"))
	}
	if obstacleareatimeslicetype.Description != obstacleareatimeslicetypeOther.Description {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Description"))
	}
	if obstacleareatimeslicetype.Interpretation != obstacleareatimeslicetypeOther.Interpretation {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if obstacleareatimeslicetype.SequenceNumber != obstacleareatimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if obstacleareatimeslicetype.CorrectionNumber != obstacleareatimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if obstacleareatimeslicetype.Reference_ownerOrganisation != obstacleareatimeslicetypeOther.Reference_ownerOrganisation {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerOrganisation"))
	}
	if obstacleareatimeslicetype.Reference_ownerRunway != obstacleareatimeslicetypeOther.Reference_ownerRunway {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerRunway"))
	}
	if obstacleareatimeslicetype.Reference_ownerAirport != obstacleareatimeslicetypeOther.Reference_ownerAirport {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerAirport"))
	}
	if obstacleareatimeslicetype.SurfaceExtent != obstacleareatimeslicetypeOther.SurfaceExtent {
		diffs = append(diffs, obstacleareatimeslicetype.GongMarshallField(stage, "SurfaceExtent"))
	}
	if ops := __gong__diffSliceOfPointers(stage, obstacleareatimeslicetype, "Obstacle", obstacleareatimeslicetypeOther.Obstacle, obstacleareatimeslicetype.Obstacle); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, obstacleareatimeslicetype, "Annotation", obstacleareatimeslicetypeOther.Annotation, obstacleareatimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, obstacleareatimeslicetype, "Extension", obstacleareatimeslicetypeOther.Extension, obstacleareatimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (obstacleareatype *ObstacleAreaType) GongDiff(stage *Stage, obstacleareatypeOther *ObstacleAreaType) (diffs []string) {
	// insertion point for field diffs
	if obstacleareatype.Name != obstacleareatypeOther.Name {
		diffs = append(diffs, obstacleareatype.GongMarshallField(stage, "Name"))
	}
	if obstacleareatype.Id != obstacleareatypeOther.Id {
		diffs = append(diffs, obstacleareatype.GongMarshallField(stage, "Id"))
	}
	if obstacleareatype.Description != obstacleareatypeOther.Description {
		diffs = append(diffs, obstacleareatype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, obstacleareatype, "TimeSlice", obstacleareatypeOther.TimeSlice, obstacleareatype.TimeSlice); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongDiff(stage *Stage, organisationauthoritypropertytypeOther *OrganisationAuthorityPropertyType) (diffs []string) {
	// insertion point for field diffs
	if organisationauthoritypropertytype.Name != organisationauthoritypropertytypeOther.Name {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Name"))
	}
	if organisationauthoritypropertytype.Owns != organisationauthoritypropertytypeOther.Owns {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Owns"))
	}
	if organisationauthoritypropertytype.NilReason != organisationauthoritypropertytypeOther.NilReason {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "NilReason"))
	}
	if organisationauthoritypropertytype.Type_XSD != organisationauthoritypropertytypeOther.Type_XSD {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if organisationauthoritypropertytype.Href != organisationauthoritypropertytypeOther.Href {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Href"))
	}
	if organisationauthoritypropertytype.Role != organisationauthoritypropertytypeOther.Role {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Role"))
	}
	if organisationauthoritypropertytype.Arcrole != organisationauthoritypropertytypeOther.Arcrole {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if organisationauthoritypropertytype.Title != organisationauthoritypropertytypeOther.Title {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Title"))
	}
	if organisationauthoritypropertytype.Show != organisationauthoritypropertytypeOther.Show {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Show"))
	}
	if organisationauthoritypropertytype.Actuate != organisationauthoritypropertytypeOther.Actuate {
		diffs = append(diffs, organisationauthoritypropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (passengerservicepropertytype *PassengerServicePropertyType) GongDiff(stage *Stage, passengerservicepropertytypeOther *PassengerServicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if passengerservicepropertytype.Name != passengerservicepropertytypeOther.Name {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Name"))
	}
	if passengerservicepropertytype.Owns != passengerservicepropertytypeOther.Owns {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if passengerservicepropertytype.NilReason != passengerservicepropertytypeOther.NilReason {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if passengerservicepropertytype.Type_XSD != passengerservicepropertytypeOther.Type_XSD {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if passengerservicepropertytype.Href != passengerservicepropertytypeOther.Href {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Href"))
	}
	if passengerservicepropertytype.Role != passengerservicepropertytypeOther.Role {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Role"))
	}
	if passengerservicepropertytype.Arcrole != passengerservicepropertytypeOther.Arcrole {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if passengerservicepropertytype.Title != passengerservicepropertytypeOther.Title {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Title"))
	}
	if passengerservicepropertytype.Show != passengerservicepropertytypeOther.Show {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Show"))
	}
	if passengerservicepropertytype.Actuate != passengerservicepropertytypeOther.Actuate {
		diffs = append(diffs, passengerservicepropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointtype *PointType) GongDiff(stage *Stage, pointtypeOther *PointType) (diffs []string) {
	// insertion point for field diffs
	if pointtype.Name != pointtypeOther.Name {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Name"))
	}
	if pointtype.Id != pointtypeOther.Id {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Id"))
	}
	if pointtype.Description != pointtypeOther.Description {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Description"))
	}
	if pointtype.SrsName != pointtypeOther.SrsName {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, pointtype, "Annotation", pointtypeOther.Annotation, pointtype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongDiff(stage *Stage, runwaydirectionpropertytypeOther *RunwayDirectionPropertyType) (diffs []string) {
	// insertion point for field diffs
	if runwaydirectionpropertytype.Name != runwaydirectionpropertytypeOther.Name {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Name"))
	}
	if runwaydirectionpropertytype.Owns != runwaydirectionpropertytypeOther.Owns {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Owns"))
	}
	if runwaydirectionpropertytype.NilReason != runwaydirectionpropertytypeOther.NilReason {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if runwaydirectionpropertytype.Type_XSD != runwaydirectionpropertytypeOther.Type_XSD {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if runwaydirectionpropertytype.Href != runwaydirectionpropertytypeOther.Href {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Href"))
	}
	if runwaydirectionpropertytype.Role != runwaydirectionpropertytypeOther.Role {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Role"))
	}
	if runwaydirectionpropertytype.Arcrole != runwaydirectionpropertytypeOther.Arcrole {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if runwaydirectionpropertytype.Title != runwaydirectionpropertytypeOther.Title {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Title"))
	}
	if runwaydirectionpropertytype.Show != runwaydirectionpropertytypeOther.Show {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Show"))
	}
	if runwaydirectionpropertytype.Actuate != runwaydirectionpropertytypeOther.Actuate {
		diffs = append(diffs, runwaydirectionpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (servicepropertytype *ServicePropertyType) GongDiff(stage *Stage, servicepropertytypeOther *ServicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if servicepropertytype.Name != servicepropertytypeOther.Name {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Name"))
	}
	if servicepropertytype.Owns != servicepropertytypeOther.Owns {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if servicepropertytype.NilReason != servicepropertytypeOther.NilReason {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if servicepropertytype.Type_XSD != servicepropertytypeOther.Type_XSD {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if servicepropertytype.Href != servicepropertytypeOther.Href {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Href"))
	}
	if servicepropertytype.Role != servicepropertytypeOther.Role {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Role"))
	}
	if servicepropertytype.Arcrole != servicepropertytypeOther.Arcrole {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if servicepropertytype.Title != servicepropertytypeOther.Title {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Title"))
	}
	if servicepropertytype.Show != servicepropertytypeOther.Show {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Show"))
	}
	if servicepropertytype.Actuate != servicepropertytypeOther.Actuate {
		diffs = append(diffs, servicepropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongDiff(stage *Stage, specialnavigationstationpropertytypeOther *SpecialNavigationStationPropertyType) (diffs []string) {
	// insertion point for field diffs
	if specialnavigationstationpropertytype.Name != specialnavigationstationpropertytypeOther.Name {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Name"))
	}
	if specialnavigationstationpropertytype.Owns != specialnavigationstationpropertytypeOther.Owns {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Owns"))
	}
	if specialnavigationstationpropertytype.NilReason != specialnavigationstationpropertytypeOther.NilReason {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if specialnavigationstationpropertytype.Type_XSD != specialnavigationstationpropertytypeOther.Type_XSD {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if specialnavigationstationpropertytype.Href != specialnavigationstationpropertytypeOther.Href {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Href"))
	}
	if specialnavigationstationpropertytype.Role != specialnavigationstationpropertytypeOther.Role {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Role"))
	}
	if specialnavigationstationpropertytype.Arcrole != specialnavigationstationpropertytypeOther.Arcrole {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if specialnavigationstationpropertytype.Title != specialnavigationstationpropertytypeOther.Title {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Title"))
	}
	if specialnavigationstationpropertytype.Show != specialnavigationstationpropertytypeOther.Show {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Show"))
	}
	if specialnavigationstationpropertytype.Actuate != specialnavigationstationpropertytypeOther.Actuate {
		diffs = append(diffs, specialnavigationstationpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (surfacepropertytype *SurfacePropertyType) GongDiff(stage *Stage, surfacepropertytypeOther *SurfacePropertyType) (diffs []string) {
	// insertion point for field diffs
	if surfacepropertytype.Name != surfacepropertytypeOther.Name {
		diffs = append(diffs, surfacepropertytype.GongMarshallField(stage, "Name"))
	}
	if surfacepropertytype.NilReason != surfacepropertytypeOther.NilReason {
		diffs = append(diffs, surfacepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if surfacepropertytype.Surface != surfacepropertytypeOther.Surface {
		diffs = append(diffs, surfacepropertytype.GongMarshallField(stage, "Surface"))
	}
	if surfacepropertytype.ElevatedSurface != surfacepropertytypeOther.ElevatedSurface {
		diffs = append(diffs, surfacepropertytype.GongMarshallField(stage, "ElevatedSurface"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (surfacetype *SurfaceType) GongDiff(stage *Stage, surfacetypeOther *SurfaceType) (diffs []string) {
	// insertion point for field diffs
	if surfacetype.Name != surfacetypeOther.Name {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Name"))
	}
	if surfacetype.Id != surfacetypeOther.Id {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Id"))
	}
	if surfacetype.Description != surfacetypeOther.Description {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "Description"))
	}
	if surfacetype.SrsName != surfacetypeOther.SrsName {
		diffs = append(diffs, surfacetype.GongMarshallField(stage, "SrsName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, surfacetype, "Annotation", surfacetypeOther.Annotation, surfacetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timesheetpropertytype *TimesheetPropertyType) GongDiff(stage *Stage, timesheetpropertytypeOther *TimesheetPropertyType) (diffs []string) {
	// insertion point for field diffs
	if timesheetpropertytype.Name != timesheetpropertytypeOther.Name {
		diffs = append(diffs, timesheetpropertytype.GongMarshallField(stage, "Name"))
	}
	if timesheetpropertytype.NilReason != timesheetpropertytypeOther.NilReason {
		diffs = append(diffs, timesheetpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if timesheetpropertytype.Timesheet != timesheetpropertytypeOther.Timesheet {
		diffs = append(diffs, timesheetpropertytype.GongMarshallField(stage, "Timesheet"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timesheettype *TimesheetType) GongDiff(stage *Stage, timesheettypeOther *TimesheetType) (diffs []string) {
	// insertion point for field diffs
	if timesheettype.Name != timesheettypeOther.Name {
		diffs = append(diffs, timesheettype.GongMarshallField(stage, "Name"))
	}
	if timesheettype.Id != timesheettypeOther.Id {
		diffs = append(diffs, timesheettype.GongMarshallField(stage, "Id"))
	}
	if timesheettype.Description != timesheettypeOther.Description {
		diffs = append(diffs, timesheettype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, timesheettype, "Annotation", timesheettypeOther.Annotation, timesheettype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, timesheettype, "Extension", timesheettypeOther.Extension, timesheettype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (unitpropertytype *UnitPropertyType) GongDiff(stage *Stage, unitpropertytypeOther *UnitPropertyType) (diffs []string) {
	// insertion point for field diffs
	if unitpropertytype.Name != unitpropertytypeOther.Name {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Name"))
	}
	if unitpropertytype.Owns != unitpropertytypeOther.Owns {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Owns"))
	}
	if unitpropertytype.NilReason != unitpropertytypeOther.NilReason {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if unitpropertytype.Type_XSD != unitpropertytypeOther.Type_XSD {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if unitpropertytype.Href != unitpropertytypeOther.Href {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Href"))
	}
	if unitpropertytype.Role != unitpropertytypeOther.Role {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Role"))
	}
	if unitpropertytype.Arcrole != unitpropertytypeOther.Arcrole {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if unitpropertytype.Title != unitpropertytypeOther.Title {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Title"))
	}
	if unitpropertytype.Show != unitpropertytypeOther.Show {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Show"))
	}
	if unitpropertytype.Actuate != unitpropertytypeOther.Actuate {
		diffs = append(diffs, unitpropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongDiff(stage *Stage, verticalstructurelightingstatuspropertytypeOther *VerticalStructureLightingStatusPropertyType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructurelightingstatuspropertytype.Name != verticalstructurelightingstatuspropertytypeOther.Name {
		diffs = append(diffs, verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "Name"))
	}
	if verticalstructurelightingstatuspropertytype.NilReason != verticalstructurelightingstatuspropertytypeOther.NilReason {
		diffs = append(diffs, verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "NilReason"))
	}
	if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != verticalstructurelightingstatuspropertytypeOther.VerticalStructureLightingStatus {
		diffs = append(diffs, verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "VerticalStructureLightingStatus"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongDiff(stage *Stage, verticalstructurelightingstatustypeOther *VerticalStructureLightingStatusType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructurelightingstatustype.Name != verticalstructurelightingstatustypeOther.Name {
		diffs = append(diffs, verticalstructurelightingstatustype.GongMarshallField(stage, "Name"))
	}
	if verticalstructurelightingstatustype.Id != verticalstructurelightingstatustypeOther.Id {
		diffs = append(diffs, verticalstructurelightingstatustype.GongMarshallField(stage, "Id"))
	}
	if verticalstructurelightingstatustype.Description != verticalstructurelightingstatustypeOther.Description {
		diffs = append(diffs, verticalstructurelightingstatustype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructurelightingstatustype, "TimeInterval", verticalstructurelightingstatustypeOther.TimeInterval, verticalstructurelightingstatustype.TimeInterval); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructurelightingstatustype, "Annotation", verticalstructurelightingstatustypeOther.Annotation, verticalstructurelightingstatustype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructurelightingstatustype, "SpecialDateAuthority", verticalstructurelightingstatustypeOther.SpecialDateAuthority, verticalstructurelightingstatustype.SpecialDateAuthority); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructurelightingstatustype, "Extension", verticalstructurelightingstatustypeOther.Extension, verticalstructurelightingstatustype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongDiff(stage *Stage, verticalstructurepartpropertytypeOther *VerticalStructurePartPropertyType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructurepartpropertytype.Name != verticalstructurepartpropertytypeOther.Name {
		diffs = append(diffs, verticalstructurepartpropertytype.GongMarshallField(stage, "Name"))
	}
	if verticalstructurepartpropertytype.NilReason != verticalstructurepartpropertytypeOther.NilReason {
		diffs = append(diffs, verticalstructurepartpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if verticalstructurepartpropertytype.VerticalStructurePart != verticalstructurepartpropertytypeOther.VerticalStructurePart {
		diffs = append(diffs, verticalstructurepartpropertytype.GongMarshallField(stage, "VerticalStructurePart"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructureparttype *VerticalStructurePartType) GongDiff(stage *Stage, verticalstructureparttypeOther *VerticalStructurePartType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructureparttype.Name != verticalstructureparttypeOther.Name {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "Name"))
	}
	if verticalstructureparttype.Id != verticalstructureparttypeOther.Id {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "Id"))
	}
	if verticalstructureparttype.Description != verticalstructureparttypeOther.Description {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructureparttype, "TimeInterval", verticalstructureparttypeOther.TimeInterval, verticalstructureparttype.TimeInterval); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructureparttype, "Annotation", verticalstructureparttypeOther.Annotation, verticalstructureparttype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructureparttype, "SpecialDateAuthority", verticalstructureparttypeOther.SpecialDateAuthority, verticalstructureparttype.SpecialDateAuthority); ops != "" {
		diffs = append(diffs, ops)
	}
	if verticalstructureparttype.HorizontalProjection_surfaceExtent != verticalstructureparttypeOther.HorizontalProjection_surfaceExtent {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_surfaceExtent"))
	}
	if verticalstructureparttype.HorizontalProjection_linearExtent != verticalstructureparttypeOther.HorizontalProjection_linearExtent {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_linearExtent"))
	}
	if verticalstructureparttype.HorizontalProjection_location != verticalstructureparttypeOther.HorizontalProjection_location {
		diffs = append(diffs, verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_location"))
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructureparttype, "Lighting", verticalstructureparttypeOther.Lighting, verticalstructureparttype.Lighting); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructureparttype, "Extension", verticalstructureparttypeOther.Extension, verticalstructureparttype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongDiff(stage *Stage, verticalstructurepropertytypeOther *VerticalStructurePropertyType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructurepropertytype.Name != verticalstructurepropertytypeOther.Name {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Name"))
	}
	if verticalstructurepropertytype.Owns != verticalstructurepropertytypeOther.Owns {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Owns"))
	}
	if verticalstructurepropertytype.NilReason != verticalstructurepropertytypeOther.NilReason {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if verticalstructurepropertytype.Type_XSD != verticalstructurepropertytypeOther.Type_XSD {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if verticalstructurepropertytype.Href != verticalstructurepropertytypeOther.Href {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Href"))
	}
	if verticalstructurepropertytype.Role != verticalstructurepropertytypeOther.Role {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Role"))
	}
	if verticalstructurepropertytype.Arcrole != verticalstructurepropertytypeOther.Arcrole {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if verticalstructurepropertytype.Title != verticalstructurepropertytypeOther.Title {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Title"))
	}
	if verticalstructurepropertytype.Show != verticalstructurepropertytypeOther.Show {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Show"))
	}
	if verticalstructurepropertytype.Actuate != verticalstructurepropertytypeOther.Actuate {
		diffs = append(diffs, verticalstructurepropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongDiff(stage *Stage, verticalstructuretimeslicepropertytypeOther *VerticalStructureTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructuretimeslicepropertytype.Name != verticalstructuretimeslicepropertytypeOther.Name {
		diffs = append(diffs, verticalstructuretimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if verticalstructuretimeslicepropertytype.Owns != verticalstructuretimeslicepropertytypeOther.Owns {
		diffs = append(diffs, verticalstructuretimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != verticalstructuretimeslicepropertytypeOther.VerticalStructureTimeSlice {
		diffs = append(diffs, verticalstructuretimeslicepropertytype.GongMarshallField(stage, "VerticalStructureTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongDiff(stage *Stage, verticalstructuretimeslicetypeOther *VerticalStructureTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructuretimeslicetype.Name != verticalstructuretimeslicetypeOther.Name {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "Name"))
	}
	if verticalstructuretimeslicetype.Id != verticalstructuretimeslicetypeOther.Id {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "Id"))
	}
	if verticalstructuretimeslicetype.Description != verticalstructuretimeslicetypeOther.Description {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "Description"))
	}
	if verticalstructuretimeslicetype.Interpretation != verticalstructuretimeslicetypeOther.Interpretation {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if verticalstructuretimeslicetype.SequenceNumber != verticalstructuretimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if verticalstructuretimeslicetype.CorrectionNumber != verticalstructuretimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if verticalstructuretimeslicetype.Marker != verticalstructuretimeslicetypeOther.Marker {
		diffs = append(diffs, verticalstructuretimeslicetype.GongMarshallField(stage, "Marker"))
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "Part", verticalstructuretimeslicetypeOther.Part, verticalstructuretimeslicetype.Part); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "HostedPassengerService", verticalstructuretimeslicetypeOther.HostedPassengerService, verticalstructuretimeslicetype.HostedPassengerService); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "SupportedGroundLight", verticalstructuretimeslicetypeOther.SupportedGroundLight, verticalstructuretimeslicetype.SupportedGroundLight); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "HostedNavaidEquipment", verticalstructuretimeslicetypeOther.HostedNavaidEquipment, verticalstructuretimeslicetype.HostedNavaidEquipment); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "HostedSpecialNavStation", verticalstructuretimeslicetypeOther.HostedSpecialNavStation, verticalstructuretimeslicetype.HostedSpecialNavStation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "HostedUnit", verticalstructuretimeslicetypeOther.HostedUnit, verticalstructuretimeslicetype.HostedUnit); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "HostedOrganisation", verticalstructuretimeslicetypeOther.HostedOrganisation, verticalstructuretimeslicetype.HostedOrganisation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "SupportedService", verticalstructuretimeslicetypeOther.SupportedService, verticalstructuretimeslicetype.SupportedService); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "Annotation", verticalstructuretimeslicetypeOther.Annotation, verticalstructuretimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "LightingAvailability", verticalstructuretimeslicetypeOther.LightingAvailability, verticalstructuretimeslicetype.LightingAvailability); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretimeslicetype, "Extension", verticalstructuretimeslicetypeOther.Extension, verticalstructuretimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticalstructuretype *VerticalStructureType) GongDiff(stage *Stage, verticalstructuretypeOther *VerticalStructureType) (diffs []string) {
	// insertion point for field diffs
	if verticalstructuretype.Name != verticalstructuretypeOther.Name {
		diffs = append(diffs, verticalstructuretype.GongMarshallField(stage, "Name"))
	}
	if verticalstructuretype.Id != verticalstructuretypeOther.Id {
		diffs = append(diffs, verticalstructuretype.GongMarshallField(stage, "Id"))
	}
	if verticalstructuretype.Description != verticalstructuretypeOther.Description {
		diffs = append(diffs, verticalstructuretype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, verticalstructuretype, "TimeSlice", verticalstructuretypeOther.TimeSlice, verticalstructuretype.TimeSlice); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
