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

func (a_extension_13 *A_extension_13) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_13s[a_extension_13]
	return ok
}

func (a_extension_14 *A_extension_14) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_extension_14s[a_extension_14]
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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AircraftCharacteristicPropertyTypes[aircraftcharacteristicpropertytype]
	return ok
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AircraftCharacteristicTypes[aircraftcharacteristictype]
	return ok
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceActivationPropertyTypes[airspaceactivationpropertytype]
	return ok
}

func (airspaceactivationtype *AirspaceActivationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceActivationTypes[airspaceactivationtype]
	return ok
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceGeometryComponentPropertyTypes[airspacegeometrycomponentpropertytype]
	return ok
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceGeometryComponentTypes[airspacegeometrycomponenttype]
	return ok
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceLayerClassPropertyTypes[airspacelayerclasspropertytype]
	return ok
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceLayerClassTypes[airspacelayerclasstype]
	return ok
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceLayerPropertyTypes[airspacelayerpropertytype]
	return ok
}

func (airspacelayertype *AirspaceLayerType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceLayerTypes[airspacelayertype]
	return ok
}

func (airspacepropertytype *AirspacePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspacePropertyTypes[airspacepropertytype]
	return ok
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceTimeSlicePropertyTypes[airspacetimeslicepropertytype]
	return ok
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceTimeSliceTypes[airspacetimeslicetype]
	return ok
}

func (airspacetype *AirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceTypes[airspacetype]
	return ok
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceVolumeDependencyPropertyTypes[airspacevolumedependencypropertytype]
	return ok
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceVolumeDependencyTypes[airspacevolumedependencytype]
	return ok
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceVolumePropertyTypes[airspacevolumepropertytype]
	return ok
}

func (airspacevolumetype *AirspaceVolumeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AirspaceVolumeTypes[airspacevolumetype]
	return ok
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AuthorityForAirspaceTimeSlicePropertyTypes[authorityforairspacetimeslicepropertytype]
	return ok
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AuthorityForAirspaceTimeSliceTypes[authorityforairspacetimeslicetype]
	return ok
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AuthorityForAirspaceTypes[authorityforairspacetype]
	return ok
}

func (curvepropertytype *CurvePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurvePropertyTypes[curvepropertytype]
	return ok
}

func (curvetype *CurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CurveTypes[curvetype]
	return ok
}

func (elevatedcurvetype *ElevatedCurveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedCurveTypes[elevatedcurvetype]
	return ok
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ElevatedSurfaceTypes[elevatedsurfacetype]
	return ok
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeoBorderTimeSlicePropertyTypes[geobordertimeslicepropertytype]
	return ok
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeoBorderTimeSliceTypes[geobordertimeslicetype]
	return ok
}

func (geobordertype *GeoBorderType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GeoBorderTypes[geobordertype]
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

func (notepropertytype *NotePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NotePropertyTypes[notepropertytype]
	return ok
}

func (notetype *NoteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTypes[notetype]
	return ok
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OrganisationAuthorityPropertyTypes[organisationauthoritypropertytype]
	return ok
}

func (routepropertytype *RoutePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RoutePropertyTypes[routepropertytype]
	return ok
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StandardLevelColumnPropertyTypes[standardlevelcolumnpropertytype]
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

func (a_extension_13 *A_extension_13) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_13) {
		return
	}

	a_extension_13.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_14 *A_extension_14) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_extension_14) {
		return
	}

	a_extension_14.Stage(stage)

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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aircraftcharacteristicpropertytype) {
		return
	}

	aircraftcharacteristicpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
		stage.StageBranch(aircraftcharacteristicpropertytype.AircraftCharacteristic)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aircraftcharacteristictype) {
		return
	}

	aircraftcharacteristictype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_12 := range aircraftcharacteristictype.Extension {
		stage.StageBranch(_a_extension_12)
	}

}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspaceactivationpropertytype) {
		return
	}

	airspaceactivationpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspaceactivationpropertytype.AirspaceActivation != nil {
		stage.StageBranch(airspaceactivationpropertytype.AirspaceActivation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspaceactivationtype *AirspaceActivationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspaceactivationtype) {
		return
	}

	airspaceactivationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspaceactivationtype.TimeInterval {
		stage.StageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range airspaceactivationtype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtype.SpecialDateAuthority {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _airspacelayerpropertytype := range airspaceactivationtype.Levels {
		stage.StageBranch(_airspacelayerpropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtype.User {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _aircraftcharacteristicpropertytype := range airspaceactivationtype.Aircraft {
		stage.StageBranch(_aircraftcharacteristicpropertytype)
	}
	for _, _a_extension_1 := range airspaceactivationtype.Extension {
		stage.StageBranch(_a_extension_1)
	}

}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacegeometrycomponentpropertytype) {
		return
	}

	airspacegeometrycomponentpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
		stage.StageBranch(airspacegeometrycomponentpropertytype.AirspaceGeometryComponent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacegeometrycomponenttype) {
		return
	}

	airspacegeometrycomponenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
		stage.StageBranch(airspacegeometrycomponenttype.TheAirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacegeometrycomponenttype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_2 := range airspacegeometrycomponenttype.Extension {
		stage.StageBranch(_a_extension_2)
	}

}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacelayerclasspropertytype) {
		return
	}

	airspacelayerclasspropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
		stage.StageBranch(airspacelayerclasspropertytype.AirspaceLayerClass)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacelayerclasstype *AirspaceLayerClassType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacelayerclasstype) {
		return
	}

	airspacelayerclasstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspacelayerclasstype.TimeInterval {
		stage.StageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range airspacelayerclasstype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspacelayerclasstype.SpecialDateAuthority {
		stage.StageBranch(_organisationauthoritypropertytype)
	}
	for _, _airspacelayerpropertytype := range airspacelayerclasstype.AssociatedLevels {
		stage.StageBranch(_airspacelayerpropertytype)
	}
	for _, _a_extension_3 := range airspacelayerclasstype.Extension {
		stage.StageBranch(_a_extension_3)
	}

}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacelayerpropertytype) {
		return
	}

	airspacelayerpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerpropertytype.AirspaceLayer != nil {
		stage.StageBranch(airspacelayerpropertytype.AirspaceLayer)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacelayertype *AirspaceLayerType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacelayertype) {
		return
	}

	airspacelayertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayertype.DiscreteLevelSeries != nil {
		stage.StageBranch(airspacelayertype.DiscreteLevelSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacelayertype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_13 := range airspacelayertype.Extension {
		stage.StageBranch(_a_extension_13)
	}

}

func (airspacepropertytype *AirspacePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacepropertytype) {
		return
	}

	airspacepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacetimeslicepropertytype) {
		return
	}

	airspacetimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
		stage.StageBranch(airspacetimeslicepropertytype.AirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacetimeslicetype) {
		return
	}

	airspacetimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicetype.ProtectedRoute != nil {
		stage.StageBranch(airspacetimeslicetype.ProtectedRoute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacelayerclasspropertytype := range airspacetimeslicetype.Class {
		stage.StageBranch(_airspacelayerclasspropertytype)
	}
	for _, _airspacegeometrycomponentpropertytype := range airspacetimeslicetype.GeometryComponent {
		stage.StageBranch(_airspacegeometrycomponentpropertytype)
	}
	for _, _airspaceactivationpropertytype := range airspacetimeslicetype.Activation {
		stage.StageBranch(_airspaceactivationpropertytype)
	}
	for _, _notepropertytype := range airspacetimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension := range airspacetimeslicetype.Extension {
		stage.StageBranch(_a_extension)
	}

}

func (airspacetype *AirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacetype) {
		return
	}

	airspacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacetimeslicepropertytype := range airspacetype.TimeSlice {
		stage.StageBranch(_airspacetimeslicepropertytype)
	}

}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacevolumedependencypropertytype) {
		return
	}

	airspacevolumedependencypropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
		stage.StageBranch(airspacevolumedependencypropertytype.AirspaceVolumeDependency)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacevolumedependencytype) {
		return
	}

	airspacevolumedependencytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencytype.TheAirspace != nil {
		stage.StageBranch(airspacevolumedependencytype.TheAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumedependencytype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_5 := range airspacevolumedependencytype.Extension {
		stage.StageBranch(_a_extension_5)
	}

}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacevolumepropertytype) {
		return
	}

	airspacevolumepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumepropertytype.AirspaceVolume != nil {
		stage.StageBranch(airspacevolumepropertytype.AirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacevolumetype *AirspaceVolumeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(airspacevolumetype) {
		return
	}

	airspacevolumetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumetype.HorizontalProjection != nil {
		stage.StageBranch(airspacevolumetype.HorizontalProjection)
	}
	if airspacevolumetype.Centreline != nil {
		stage.StageBranch(airspacevolumetype.Centreline)
	}
	if airspacevolumetype.ContributorAirspace != nil {
		stage.StageBranch(airspacevolumetype.ContributorAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_4 := range airspacevolumetype.Extension {
		stage.StageBranch(_a_extension_4)
	}

}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(authorityforairspacetimeslicepropertytype) {
		return
	}

	authorityforairspacetimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
		stage.StageBranch(authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(authorityforairspacetimeslicetype) {
		return
	}

	authorityforairspacetimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
		stage.StageBranch(authorityforairspacetimeslicetype.ResponsibleOrganisation)
	}
	if authorityforairspacetimeslicetype.AssignedAirspace != nil {
		stage.StageBranch(authorityforairspacetimeslicetype.AssignedAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range authorityforairspacetimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_6 := range authorityforairspacetimeslicetype.Extension {
		stage.StageBranch(_a_extension_6)
	}

}

func (authorityforairspacetype *AuthorityForAirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(authorityforairspacetype) {
		return
	}

	authorityforairspacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _authorityforairspacetimeslicepropertytype := range authorityforairspacetype.TimeSlice {
		stage.StageBranch(_authorityforairspacetimeslicepropertytype)
	}

}

func (curvepropertytype *CurvePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curvepropertytype) {
		return
	}

	curvepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytype.Curve != nil {
		stage.StageBranch(curvepropertytype.Curve)
	}
	if curvepropertytype.ElevatedCurve != nil {
		stage.StageBranch(curvepropertytype.ElevatedCurve)
	}

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
	for _, _a_extension_8 := range elevatedcurvetype.Extension {
		stage.StageBranch(_a_extension_8)
	}

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
	for _, _a_extension_9 := range elevatedsurfacetype.Extension {
		stage.StageBranch(_a_extension_9)
	}

}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geobordertimeslicepropertytype) {
		return
	}

	geobordertimeslicepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
		stage.StageBranch(geobordertimeslicepropertytype.GeoBorderTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geobordertimeslicetype) {
		return
	}

	geobordertimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicetype.Border != nil {
		stage.StageBranch(geobordertimeslicetype.Border)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range geobordertimeslicetype.Annotation {
		stage.StageBranch(_notepropertytype)
	}
	for _, _a_extension_7 := range geobordertimeslicetype.Extension {
		stage.StageBranch(_a_extension_7)
	}

}

func (geobordertype *GeoBorderType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(geobordertype) {
		return
	}

	geobordertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geobordertimeslicepropertytype := range geobordertype.TimeSlice {
		stage.StageBranch(_geobordertimeslicepropertytype)
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
	for _, _a_extension_10 := range linguisticnotetype.Extension {
		stage.StageBranch(_a_extension_10)
	}

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
	for _, _a_extension_11 := range notetype.Extension {
		stage.StageBranch(_a_extension_11)
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

func (routepropertytype *RoutePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(routepropertytype) {
		return
	}

	routepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(standardlevelcolumnpropertytype) {
		return
	}

	standardlevelcolumnpropertytype.Stage(stage)

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
	for _, _a_extension_14 := range timesheettype.Extension {
		stage.StageBranch(_a_extension_14)
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

	case *A_extension_13:
		toT := GongCopyBranchA_extension_13(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_extension_14:
		toT := GongCopyBranchA_extension_14(mapOrigCopy, fromT)
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

	case *AircraftCharacteristicPropertyType:
		toT := GongCopyBranchAircraftCharacteristicPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AircraftCharacteristicType:
		toT := GongCopyBranchAircraftCharacteristicType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceActivationPropertyType:
		toT := GongCopyBranchAirspaceActivationPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceActivationType:
		toT := GongCopyBranchAirspaceActivationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceGeometryComponentPropertyType:
		toT := GongCopyBranchAirspaceGeometryComponentPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceGeometryComponentType:
		toT := GongCopyBranchAirspaceGeometryComponentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceLayerClassPropertyType:
		toT := GongCopyBranchAirspaceLayerClassPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceLayerClassType:
		toT := GongCopyBranchAirspaceLayerClassType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceLayerPropertyType:
		toT := GongCopyBranchAirspaceLayerPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceLayerType:
		toT := GongCopyBranchAirspaceLayerType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspacePropertyType:
		toT := GongCopyBranchAirspacePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceTimeSlicePropertyType:
		toT := GongCopyBranchAirspaceTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceTimeSliceType:
		toT := GongCopyBranchAirspaceTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceType:
		toT := GongCopyBranchAirspaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceVolumeDependencyPropertyType:
		toT := GongCopyBranchAirspaceVolumeDependencyPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceVolumeDependencyType:
		toT := GongCopyBranchAirspaceVolumeDependencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceVolumePropertyType:
		toT := GongCopyBranchAirspaceVolumePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AirspaceVolumeType:
		toT := GongCopyBranchAirspaceVolumeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AuthorityForAirspaceTimeSlicePropertyType:
		toT := GongCopyBranchAuthorityForAirspaceTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AuthorityForAirspaceTimeSliceType:
		toT := GongCopyBranchAuthorityForAirspaceTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AuthorityForAirspaceType:
		toT := GongCopyBranchAuthorityForAirspaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurvePropertyType:
		toT := GongCopyBranchCurvePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CurveType:
		toT := GongCopyBranchCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedCurveType:
		toT := GongCopyBranchElevatedCurveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ElevatedSurfaceType:
		toT := GongCopyBranchElevatedSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoBorderTimeSlicePropertyType:
		toT := GongCopyBranchGeoBorderTimeSlicePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoBorderTimeSliceType:
		toT := GongCopyBranchGeoBorderTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GeoBorderType:
		toT := GongCopyBranchGeoBorderType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinguisticNotePropertyType:
		toT := GongCopyBranchLinguisticNotePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinguisticNoteType:
		toT := GongCopyBranchLinguisticNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePropertyType:
		toT := GongCopyBranchNotePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteType:
		toT := GongCopyBranchNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OrganisationAuthorityPropertyType:
		toT := GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RoutePropertyType:
		toT := GongCopyBranchRoutePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StandardLevelColumnPropertyType:
		toT := GongCopyBranchStandardLevelColumnPropertyType(mapOrigCopy, fromT)
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

func GongCopyBranchA_extension_13(mapOrigCopy map[any]any, a_extension_13From *A_extension_13) (a_extension_13To *A_extension_13) {
	var alreadyCopied bool
	a_extension_13To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_13From)
	if alreadyCopied {
		return
	}
	a_extension_13From.GongCopyBasicFields(a_extension_13To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_extension_14(mapOrigCopy map[any]any, a_extension_14From *A_extension_14) (a_extension_14To *A_extension_14) {
	var alreadyCopied bool
	a_extension_14To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_extension_14From)
	if alreadyCopied {
		return
	}
	a_extension_14From.GongCopyBasicFields(a_extension_14To)

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

func GongCopyBranchAircraftCharacteristicPropertyType(mapOrigCopy map[any]any, aircraftcharacteristicpropertytypeFrom *AircraftCharacteristicPropertyType) (aircraftcharacteristicpropertytypeTo *AircraftCharacteristicPropertyType) {
	var alreadyCopied bool
	aircraftcharacteristicpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aircraftcharacteristicpropertytypeFrom)
	if alreadyCopied {
		return
	}
	aircraftcharacteristicpropertytypeFrom.GongCopyBasicFields(aircraftcharacteristicpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if aircraftcharacteristicpropertytypeFrom.AircraftCharacteristic != nil {
		aircraftcharacteristicpropertytypeTo.AircraftCharacteristic = GongCopyBranchAircraftCharacteristicType(mapOrigCopy, aircraftcharacteristicpropertytypeFrom.AircraftCharacteristic)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAircraftCharacteristicType(mapOrigCopy map[any]any, aircraftcharacteristictypeFrom *AircraftCharacteristicType) (aircraftcharacteristictypeTo *AircraftCharacteristicType) {
	var alreadyCopied bool
	aircraftcharacteristictypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aircraftcharacteristictypeFrom)
	if alreadyCopied {
		return
	}
	aircraftcharacteristictypeFrom.GongCopyBasicFields(aircraftcharacteristictypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aircraftcharacteristictypeFrom.Annotation {
		aircraftcharacteristictypeTo.Annotation = append(aircraftcharacteristictypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_12 := range aircraftcharacteristictypeFrom.Extension {
		aircraftcharacteristictypeTo.Extension = append(aircraftcharacteristictypeTo.Extension, GongCopyBranchA_extension_12(mapOrigCopy, _a_extension_12))
	}

	return
}

func GongCopyBranchAirspaceActivationPropertyType(mapOrigCopy map[any]any, airspaceactivationpropertytypeFrom *AirspaceActivationPropertyType) (airspaceactivationpropertytypeTo *AirspaceActivationPropertyType) {
	var alreadyCopied bool
	airspaceactivationpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspaceactivationpropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspaceactivationpropertytypeFrom.GongCopyBasicFields(airspaceactivationpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspaceactivationpropertytypeFrom.AirspaceActivation != nil {
		airspaceactivationpropertytypeTo.AirspaceActivation = GongCopyBranchAirspaceActivationType(mapOrigCopy, airspaceactivationpropertytypeFrom.AirspaceActivation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceActivationType(mapOrigCopy map[any]any, airspaceactivationtypeFrom *AirspaceActivationType) (airspaceactivationtypeTo *AirspaceActivationType) {
	var alreadyCopied bool
	airspaceactivationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspaceactivationtypeFrom)
	if alreadyCopied {
		return
	}
	airspaceactivationtypeFrom.GongCopyBasicFields(airspaceactivationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspaceactivationtypeFrom.TimeInterval {
		airspaceactivationtypeTo.TimeInterval = append(airspaceactivationtypeTo.TimeInterval, GongCopyBranchTimesheetPropertyType(mapOrigCopy, _timesheetpropertytype))
	}
	for _, _notepropertytype := range airspaceactivationtypeFrom.Annotation {
		airspaceactivationtypeTo.Annotation = append(airspaceactivationtypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtypeFrom.SpecialDateAuthority {
		airspaceactivationtypeTo.SpecialDateAuthority = append(airspaceactivationtypeTo.SpecialDateAuthority, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _airspacelayerpropertytype := range airspaceactivationtypeFrom.Levels {
		airspaceactivationtypeTo.Levels = append(airspaceactivationtypeTo.Levels, GongCopyBranchAirspaceLayerPropertyType(mapOrigCopy, _airspacelayerpropertytype))
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtypeFrom.User {
		airspaceactivationtypeTo.User = append(airspaceactivationtypeTo.User, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _aircraftcharacteristicpropertytype := range airspaceactivationtypeFrom.Aircraft {
		airspaceactivationtypeTo.Aircraft = append(airspaceactivationtypeTo.Aircraft, GongCopyBranchAircraftCharacteristicPropertyType(mapOrigCopy, _aircraftcharacteristicpropertytype))
	}
	for _, _a_extension_1 := range airspaceactivationtypeFrom.Extension {
		airspaceactivationtypeTo.Extension = append(airspaceactivationtypeTo.Extension, GongCopyBranchA_extension_1(mapOrigCopy, _a_extension_1))
	}

	return
}

func GongCopyBranchAirspaceGeometryComponentPropertyType(mapOrigCopy map[any]any, airspacegeometrycomponentpropertytypeFrom *AirspaceGeometryComponentPropertyType) (airspacegeometrycomponentpropertytypeTo *AirspaceGeometryComponentPropertyType) {
	var alreadyCopied bool
	airspacegeometrycomponentpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacegeometrycomponentpropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacegeometrycomponentpropertytypeFrom.GongCopyBasicFields(airspacegeometrycomponentpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponentpropertytypeFrom.AirspaceGeometryComponent != nil {
		airspacegeometrycomponentpropertytypeTo.AirspaceGeometryComponent = GongCopyBranchAirspaceGeometryComponentType(mapOrigCopy, airspacegeometrycomponentpropertytypeFrom.AirspaceGeometryComponent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceGeometryComponentType(mapOrigCopy map[any]any, airspacegeometrycomponenttypeFrom *AirspaceGeometryComponentType) (airspacegeometrycomponenttypeTo *AirspaceGeometryComponentType) {
	var alreadyCopied bool
	airspacegeometrycomponenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacegeometrycomponenttypeFrom)
	if alreadyCopied {
		return
	}
	airspacegeometrycomponenttypeFrom.GongCopyBasicFields(airspacegeometrycomponenttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponenttypeFrom.TheAirspaceVolume != nil {
		airspacegeometrycomponenttypeTo.TheAirspaceVolume = GongCopyBranchAirspaceVolumePropertyType(mapOrigCopy, airspacegeometrycomponenttypeFrom.TheAirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacegeometrycomponenttypeFrom.Annotation {
		airspacegeometrycomponenttypeTo.Annotation = append(airspacegeometrycomponenttypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_2 := range airspacegeometrycomponenttypeFrom.Extension {
		airspacegeometrycomponenttypeTo.Extension = append(airspacegeometrycomponenttypeTo.Extension, GongCopyBranchA_extension_2(mapOrigCopy, _a_extension_2))
	}

	return
}

func GongCopyBranchAirspaceLayerClassPropertyType(mapOrigCopy map[any]any, airspacelayerclasspropertytypeFrom *AirspaceLayerClassPropertyType) (airspacelayerclasspropertytypeTo *AirspaceLayerClassPropertyType) {
	var alreadyCopied bool
	airspacelayerclasspropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacelayerclasspropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacelayerclasspropertytypeFrom.GongCopyBasicFields(airspacelayerclasspropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerclasspropertytypeFrom.AirspaceLayerClass != nil {
		airspacelayerclasspropertytypeTo.AirspaceLayerClass = GongCopyBranchAirspaceLayerClassType(mapOrigCopy, airspacelayerclasspropertytypeFrom.AirspaceLayerClass)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceLayerClassType(mapOrigCopy map[any]any, airspacelayerclasstypeFrom *AirspaceLayerClassType) (airspacelayerclasstypeTo *AirspaceLayerClassType) {
	var alreadyCopied bool
	airspacelayerclasstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacelayerclasstypeFrom)
	if alreadyCopied {
		return
	}
	airspacelayerclasstypeFrom.GongCopyBasicFields(airspacelayerclasstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspacelayerclasstypeFrom.TimeInterval {
		airspacelayerclasstypeTo.TimeInterval = append(airspacelayerclasstypeTo.TimeInterval, GongCopyBranchTimesheetPropertyType(mapOrigCopy, _timesheetpropertytype))
	}
	for _, _notepropertytype := range airspacelayerclasstypeFrom.Annotation {
		airspacelayerclasstypeTo.Annotation = append(airspacelayerclasstypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _organisationauthoritypropertytype := range airspacelayerclasstypeFrom.SpecialDateAuthority {
		airspacelayerclasstypeTo.SpecialDateAuthority = append(airspacelayerclasstypeTo.SpecialDateAuthority, GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, _organisationauthoritypropertytype))
	}
	for _, _airspacelayerpropertytype := range airspacelayerclasstypeFrom.AssociatedLevels {
		airspacelayerclasstypeTo.AssociatedLevels = append(airspacelayerclasstypeTo.AssociatedLevels, GongCopyBranchAirspaceLayerPropertyType(mapOrigCopy, _airspacelayerpropertytype))
	}
	for _, _a_extension_3 := range airspacelayerclasstypeFrom.Extension {
		airspacelayerclasstypeTo.Extension = append(airspacelayerclasstypeTo.Extension, GongCopyBranchA_extension_3(mapOrigCopy, _a_extension_3))
	}

	return
}

func GongCopyBranchAirspaceLayerPropertyType(mapOrigCopy map[any]any, airspacelayerpropertytypeFrom *AirspaceLayerPropertyType) (airspacelayerpropertytypeTo *AirspaceLayerPropertyType) {
	var alreadyCopied bool
	airspacelayerpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacelayerpropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacelayerpropertytypeFrom.GongCopyBasicFields(airspacelayerpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerpropertytypeFrom.AirspaceLayer != nil {
		airspacelayerpropertytypeTo.AirspaceLayer = GongCopyBranchAirspaceLayerType(mapOrigCopy, airspacelayerpropertytypeFrom.AirspaceLayer)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceLayerType(mapOrigCopy map[any]any, airspacelayertypeFrom *AirspaceLayerType) (airspacelayertypeTo *AirspaceLayerType) {
	var alreadyCopied bool
	airspacelayertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacelayertypeFrom)
	if alreadyCopied {
		return
	}
	airspacelayertypeFrom.GongCopyBasicFields(airspacelayertypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayertypeFrom.DiscreteLevelSeries != nil {
		airspacelayertypeTo.DiscreteLevelSeries = GongCopyBranchStandardLevelColumnPropertyType(mapOrigCopy, airspacelayertypeFrom.DiscreteLevelSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacelayertypeFrom.Annotation {
		airspacelayertypeTo.Annotation = append(airspacelayertypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_13 := range airspacelayertypeFrom.Extension {
		airspacelayertypeTo.Extension = append(airspacelayertypeTo.Extension, GongCopyBranchA_extension_13(mapOrigCopy, _a_extension_13))
	}

	return
}

func GongCopyBranchAirspacePropertyType(mapOrigCopy map[any]any, airspacepropertytypeFrom *AirspacePropertyType) (airspacepropertytypeTo *AirspacePropertyType) {
	var alreadyCopied bool
	airspacepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacepropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacepropertytypeFrom.GongCopyBasicFields(airspacepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceTimeSlicePropertyType(mapOrigCopy map[any]any, airspacetimeslicepropertytypeFrom *AirspaceTimeSlicePropertyType) (airspacetimeslicepropertytypeTo *AirspaceTimeSlicePropertyType) {
	var alreadyCopied bool
	airspacetimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacetimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacetimeslicepropertytypeFrom.GongCopyBasicFields(airspacetimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicepropertytypeFrom.AirspaceTimeSlice != nil {
		airspacetimeslicepropertytypeTo.AirspaceTimeSlice = GongCopyBranchAirspaceTimeSliceType(mapOrigCopy, airspacetimeslicepropertytypeFrom.AirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceTimeSliceType(mapOrigCopy map[any]any, airspacetimeslicetypeFrom *AirspaceTimeSliceType) (airspacetimeslicetypeTo *AirspaceTimeSliceType) {
	var alreadyCopied bool
	airspacetimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacetimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	airspacetimeslicetypeFrom.GongCopyBasicFields(airspacetimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicetypeFrom.ProtectedRoute != nil {
		airspacetimeslicetypeTo.ProtectedRoute = GongCopyBranchRoutePropertyType(mapOrigCopy, airspacetimeslicetypeFrom.ProtectedRoute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacelayerclasspropertytype := range airspacetimeslicetypeFrom.Class {
		airspacetimeslicetypeTo.Class = append(airspacetimeslicetypeTo.Class, GongCopyBranchAirspaceLayerClassPropertyType(mapOrigCopy, _airspacelayerclasspropertytype))
	}
	for _, _airspacegeometrycomponentpropertytype := range airspacetimeslicetypeFrom.GeometryComponent {
		airspacetimeslicetypeTo.GeometryComponent = append(airspacetimeslicetypeTo.GeometryComponent, GongCopyBranchAirspaceGeometryComponentPropertyType(mapOrigCopy, _airspacegeometrycomponentpropertytype))
	}
	for _, _airspaceactivationpropertytype := range airspacetimeslicetypeFrom.Activation {
		airspacetimeslicetypeTo.Activation = append(airspacetimeslicetypeTo.Activation, GongCopyBranchAirspaceActivationPropertyType(mapOrigCopy, _airspaceactivationpropertytype))
	}
	for _, _notepropertytype := range airspacetimeslicetypeFrom.Annotation {
		airspacetimeslicetypeTo.Annotation = append(airspacetimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension := range airspacetimeslicetypeFrom.Extension {
		airspacetimeslicetypeTo.Extension = append(airspacetimeslicetypeTo.Extension, GongCopyBranchA_extension(mapOrigCopy, _a_extension))
	}

	return
}

func GongCopyBranchAirspaceType(mapOrigCopy map[any]any, airspacetypeFrom *AirspaceType) (airspacetypeTo *AirspaceType) {
	var alreadyCopied bool
	airspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacetypeFrom)
	if alreadyCopied {
		return
	}
	airspacetypeFrom.GongCopyBasicFields(airspacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacetimeslicepropertytype := range airspacetypeFrom.TimeSlice {
		airspacetypeTo.TimeSlice = append(airspacetypeTo.TimeSlice, GongCopyBranchAirspaceTimeSlicePropertyType(mapOrigCopy, _airspacetimeslicepropertytype))
	}

	return
}

func GongCopyBranchAirspaceVolumeDependencyPropertyType(mapOrigCopy map[any]any, airspacevolumedependencypropertytypeFrom *AirspaceVolumeDependencyPropertyType) (airspacevolumedependencypropertytypeTo *AirspaceVolumeDependencyPropertyType) {
	var alreadyCopied bool
	airspacevolumedependencypropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacevolumedependencypropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacevolumedependencypropertytypeFrom.GongCopyBasicFields(airspacevolumedependencypropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencypropertytypeFrom.AirspaceVolumeDependency != nil {
		airspacevolumedependencypropertytypeTo.AirspaceVolumeDependency = GongCopyBranchAirspaceVolumeDependencyType(mapOrigCopy, airspacevolumedependencypropertytypeFrom.AirspaceVolumeDependency)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceVolumeDependencyType(mapOrigCopy map[any]any, airspacevolumedependencytypeFrom *AirspaceVolumeDependencyType) (airspacevolumedependencytypeTo *AirspaceVolumeDependencyType) {
	var alreadyCopied bool
	airspacevolumedependencytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacevolumedependencytypeFrom)
	if alreadyCopied {
		return
	}
	airspacevolumedependencytypeFrom.GongCopyBasicFields(airspacevolumedependencytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencytypeFrom.TheAirspace != nil {
		airspacevolumedependencytypeTo.TheAirspace = GongCopyBranchAirspacePropertyType(mapOrigCopy, airspacevolumedependencytypeFrom.TheAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumedependencytypeFrom.Annotation {
		airspacevolumedependencytypeTo.Annotation = append(airspacevolumedependencytypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_5 := range airspacevolumedependencytypeFrom.Extension {
		airspacevolumedependencytypeTo.Extension = append(airspacevolumedependencytypeTo.Extension, GongCopyBranchA_extension_5(mapOrigCopy, _a_extension_5))
	}

	return
}

func GongCopyBranchAirspaceVolumePropertyType(mapOrigCopy map[any]any, airspacevolumepropertytypeFrom *AirspaceVolumePropertyType) (airspacevolumepropertytypeTo *AirspaceVolumePropertyType) {
	var alreadyCopied bool
	airspacevolumepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacevolumepropertytypeFrom)
	if alreadyCopied {
		return
	}
	airspacevolumepropertytypeFrom.GongCopyBasicFields(airspacevolumepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumepropertytypeFrom.AirspaceVolume != nil {
		airspacevolumepropertytypeTo.AirspaceVolume = GongCopyBranchAirspaceVolumeType(mapOrigCopy, airspacevolumepropertytypeFrom.AirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAirspaceVolumeType(mapOrigCopy map[any]any, airspacevolumetypeFrom *AirspaceVolumeType) (airspacevolumetypeTo *AirspaceVolumeType) {
	var alreadyCopied bool
	airspacevolumetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, airspacevolumetypeFrom)
	if alreadyCopied {
		return
	}
	airspacevolumetypeFrom.GongCopyBasicFields(airspacevolumetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumetypeFrom.HorizontalProjection != nil {
		airspacevolumetypeTo.HorizontalProjection = GongCopyBranchSurfacePropertyType(mapOrigCopy, airspacevolumetypeFrom.HorizontalProjection)
	}
	if airspacevolumetypeFrom.Centreline != nil {
		airspacevolumetypeTo.Centreline = GongCopyBranchCurvePropertyType(mapOrigCopy, airspacevolumetypeFrom.Centreline)
	}
	if airspacevolumetypeFrom.ContributorAirspace != nil {
		airspacevolumetypeTo.ContributorAirspace = GongCopyBranchAirspaceVolumeDependencyPropertyType(mapOrigCopy, airspacevolumetypeFrom.ContributorAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumetypeFrom.Annotation {
		airspacevolumetypeTo.Annotation = append(airspacevolumetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_4 := range airspacevolumetypeFrom.Extension {
		airspacevolumetypeTo.Extension = append(airspacevolumetypeTo.Extension, GongCopyBranchA_extension_4(mapOrigCopy, _a_extension_4))
	}

	return
}

func GongCopyBranchAuthorityForAirspaceTimeSlicePropertyType(mapOrigCopy map[any]any, authorityforairspacetimeslicepropertytypeFrom *AuthorityForAirspaceTimeSlicePropertyType) (authorityforairspacetimeslicepropertytypeTo *AuthorityForAirspaceTimeSlicePropertyType) {
	var alreadyCopied bool
	authorityforairspacetimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, authorityforairspacetimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	authorityforairspacetimeslicepropertytypeFrom.GongCopyBasicFields(authorityforairspacetimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicepropertytypeFrom.AuthorityForAirspaceTimeSlice != nil {
		authorityforairspacetimeslicepropertytypeTo.AuthorityForAirspaceTimeSlice = GongCopyBranchAuthorityForAirspaceTimeSliceType(mapOrigCopy, authorityforairspacetimeslicepropertytypeFrom.AuthorityForAirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAuthorityForAirspaceTimeSliceType(mapOrigCopy map[any]any, authorityforairspacetimeslicetypeFrom *AuthorityForAirspaceTimeSliceType) (authorityforairspacetimeslicetypeTo *AuthorityForAirspaceTimeSliceType) {
	var alreadyCopied bool
	authorityforairspacetimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, authorityforairspacetimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	authorityforairspacetimeslicetypeFrom.GongCopyBasicFields(authorityforairspacetimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicetypeFrom.ResponsibleOrganisation != nil {
		authorityforairspacetimeslicetypeTo.ResponsibleOrganisation = GongCopyBranchOrganisationAuthorityPropertyType(mapOrigCopy, authorityforairspacetimeslicetypeFrom.ResponsibleOrganisation)
	}
	if authorityforairspacetimeslicetypeFrom.AssignedAirspace != nil {
		authorityforairspacetimeslicetypeTo.AssignedAirspace = GongCopyBranchAirspacePropertyType(mapOrigCopy, authorityforairspacetimeslicetypeFrom.AssignedAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range authorityforairspacetimeslicetypeFrom.Annotation {
		authorityforairspacetimeslicetypeTo.Annotation = append(authorityforairspacetimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_6 := range authorityforairspacetimeslicetypeFrom.Extension {
		authorityforairspacetimeslicetypeTo.Extension = append(authorityforairspacetimeslicetypeTo.Extension, GongCopyBranchA_extension_6(mapOrigCopy, _a_extension_6))
	}

	return
}

func GongCopyBranchAuthorityForAirspaceType(mapOrigCopy map[any]any, authorityforairspacetypeFrom *AuthorityForAirspaceType) (authorityforairspacetypeTo *AuthorityForAirspaceType) {
	var alreadyCopied bool
	authorityforairspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, authorityforairspacetypeFrom)
	if alreadyCopied {
		return
	}
	authorityforairspacetypeFrom.GongCopyBasicFields(authorityforairspacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _authorityforairspacetimeslicepropertytype := range authorityforairspacetypeFrom.TimeSlice {
		authorityforairspacetypeTo.TimeSlice = append(authorityforairspacetypeTo.TimeSlice, GongCopyBranchAuthorityForAirspaceTimeSlicePropertyType(mapOrigCopy, _authorityforairspacetimeslicepropertytype))
	}

	return
}

func GongCopyBranchCurvePropertyType(mapOrigCopy map[any]any, curvepropertytypeFrom *CurvePropertyType) (curvepropertytypeTo *CurvePropertyType) {
	var alreadyCopied bool
	curvepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curvepropertytypeFrom)
	if alreadyCopied {
		return
	}
	curvepropertytypeFrom.GongCopyBasicFields(curvepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytypeFrom.Curve != nil {
		curvepropertytypeTo.Curve = GongCopyBranchCurveType(mapOrigCopy, curvepropertytypeFrom.Curve)
	}
	if curvepropertytypeFrom.ElevatedCurve != nil {
		curvepropertytypeTo.ElevatedCurve = GongCopyBranchElevatedCurveType(mapOrigCopy, curvepropertytypeFrom.ElevatedCurve)
	}

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
	for _, _a_extension_8 := range elevatedcurvetypeFrom.Extension {
		elevatedcurvetypeTo.Extension = append(elevatedcurvetypeTo.Extension, GongCopyBranchA_extension_8(mapOrigCopy, _a_extension_8))
	}

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
	for _, _a_extension_9 := range elevatedsurfacetypeFrom.Extension {
		elevatedsurfacetypeTo.Extension = append(elevatedsurfacetypeTo.Extension, GongCopyBranchA_extension_9(mapOrigCopy, _a_extension_9))
	}

	return
}

func GongCopyBranchGeoBorderTimeSlicePropertyType(mapOrigCopy map[any]any, geobordertimeslicepropertytypeFrom *GeoBorderTimeSlicePropertyType) (geobordertimeslicepropertytypeTo *GeoBorderTimeSlicePropertyType) {
	var alreadyCopied bool
	geobordertimeslicepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geobordertimeslicepropertytypeFrom)
	if alreadyCopied {
		return
	}
	geobordertimeslicepropertytypeFrom.GongCopyBasicFields(geobordertimeslicepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicepropertytypeFrom.GeoBorderTimeSlice != nil {
		geobordertimeslicepropertytypeTo.GeoBorderTimeSlice = GongCopyBranchGeoBorderTimeSliceType(mapOrigCopy, geobordertimeslicepropertytypeFrom.GeoBorderTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGeoBorderTimeSliceType(mapOrigCopy map[any]any, geobordertimeslicetypeFrom *GeoBorderTimeSliceType) (geobordertimeslicetypeTo *GeoBorderTimeSliceType) {
	var alreadyCopied bool
	geobordertimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geobordertimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	geobordertimeslicetypeFrom.GongCopyBasicFields(geobordertimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicetypeFrom.Border != nil {
		geobordertimeslicetypeTo.Border = GongCopyBranchCurvePropertyType(mapOrigCopy, geobordertimeslicetypeFrom.Border)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range geobordertimeslicetypeFrom.Annotation {
		geobordertimeslicetypeTo.Annotation = append(geobordertimeslicetypeTo.Annotation, GongCopyBranchNotePropertyType(mapOrigCopy, _notepropertytype))
	}
	for _, _a_extension_7 := range geobordertimeslicetypeFrom.Extension {
		geobordertimeslicetypeTo.Extension = append(geobordertimeslicetypeTo.Extension, GongCopyBranchA_extension_7(mapOrigCopy, _a_extension_7))
	}

	return
}

func GongCopyBranchGeoBorderType(mapOrigCopy map[any]any, geobordertypeFrom *GeoBorderType) (geobordertypeTo *GeoBorderType) {
	var alreadyCopied bool
	geobordertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, geobordertypeFrom)
	if alreadyCopied {
		return
	}
	geobordertypeFrom.GongCopyBasicFields(geobordertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geobordertimeslicepropertytype := range geobordertypeFrom.TimeSlice {
		geobordertypeTo.TimeSlice = append(geobordertypeTo.TimeSlice, GongCopyBranchGeoBorderTimeSlicePropertyType(mapOrigCopy, _geobordertimeslicepropertytype))
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
	for _, _a_extension_10 := range linguisticnotetypeFrom.Extension {
		linguisticnotetypeTo.Extension = append(linguisticnotetypeTo.Extension, GongCopyBranchA_extension_10(mapOrigCopy, _a_extension_10))
	}

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
	for _, _a_extension_11 := range notetypeFrom.Extension {
		notetypeTo.Extension = append(notetypeTo.Extension, GongCopyBranchA_extension_11(mapOrigCopy, _a_extension_11))
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

func GongCopyBranchRoutePropertyType(mapOrigCopy map[any]any, routepropertytypeFrom *RoutePropertyType) (routepropertytypeTo *RoutePropertyType) {
	var alreadyCopied bool
	routepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, routepropertytypeFrom)
	if alreadyCopied {
		return
	}
	routepropertytypeFrom.GongCopyBasicFields(routepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStandardLevelColumnPropertyType(mapOrigCopy map[any]any, standardlevelcolumnpropertytypeFrom *StandardLevelColumnPropertyType) (standardlevelcolumnpropertytypeTo *StandardLevelColumnPropertyType) {
	var alreadyCopied bool
	standardlevelcolumnpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, standardlevelcolumnpropertytypeFrom)
	if alreadyCopied {
		return
	}
	standardlevelcolumnpropertytypeFrom.GongCopyBasicFields(standardlevelcolumnpropertytypeTo)

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
	for _, _a_extension_14 := range timesheettypeFrom.Extension {
		timesheettypeTo.Extension = append(timesheettypeTo.Extension, GongCopyBranchA_extension_14(mapOrigCopy, _a_extension_14))
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

func (a_extension_13 *A_extension_13) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_13) {
		return
	}

	a_extension_13.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_extension_14 *A_extension_14) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_extension_14) {
		return
	}

	a_extension_14.Unstage(stage)

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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aircraftcharacteristicpropertytype) {
		return
	}

	aircraftcharacteristicpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
		stage.UnstageBranch(aircraftcharacteristicpropertytype.AircraftCharacteristic)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aircraftcharacteristictype) {
		return
	}

	aircraftcharacteristictype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_12 := range aircraftcharacteristictype.Extension {
		stage.UnstageBranch(_a_extension_12)
	}

}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspaceactivationpropertytype) {
		return
	}

	airspaceactivationpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspaceactivationpropertytype.AirspaceActivation != nil {
		stage.UnstageBranch(airspaceactivationpropertytype.AirspaceActivation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspaceactivationtype *AirspaceActivationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspaceactivationtype) {
		return
	}

	airspaceactivationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspaceactivationtype.TimeInterval {
		stage.UnstageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range airspaceactivationtype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtype.SpecialDateAuthority {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _airspacelayerpropertytype := range airspaceactivationtype.Levels {
		stage.UnstageBranch(_airspacelayerpropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspaceactivationtype.User {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _aircraftcharacteristicpropertytype := range airspaceactivationtype.Aircraft {
		stage.UnstageBranch(_aircraftcharacteristicpropertytype)
	}
	for _, _a_extension_1 := range airspaceactivationtype.Extension {
		stage.UnstageBranch(_a_extension_1)
	}

}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacegeometrycomponentpropertytype) {
		return
	}

	airspacegeometrycomponentpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
		stage.UnstageBranch(airspacegeometrycomponentpropertytype.AirspaceGeometryComponent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacegeometrycomponenttype) {
		return
	}

	airspacegeometrycomponenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
		stage.UnstageBranch(airspacegeometrycomponenttype.TheAirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacegeometrycomponenttype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_2 := range airspacegeometrycomponenttype.Extension {
		stage.UnstageBranch(_a_extension_2)
	}

}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacelayerclasspropertytype) {
		return
	}

	airspacelayerclasspropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
		stage.UnstageBranch(airspacelayerclasspropertytype.AirspaceLayerClass)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacelayerclasstype *AirspaceLayerClassType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacelayerclasstype) {
		return
	}

	airspacelayerclasstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _timesheetpropertytype := range airspacelayerclasstype.TimeInterval {
		stage.UnstageBranch(_timesheetpropertytype)
	}
	for _, _notepropertytype := range airspacelayerclasstype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _organisationauthoritypropertytype := range airspacelayerclasstype.SpecialDateAuthority {
		stage.UnstageBranch(_organisationauthoritypropertytype)
	}
	for _, _airspacelayerpropertytype := range airspacelayerclasstype.AssociatedLevels {
		stage.UnstageBranch(_airspacelayerpropertytype)
	}
	for _, _a_extension_3 := range airspacelayerclasstype.Extension {
		stage.UnstageBranch(_a_extension_3)
	}

}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacelayerpropertytype) {
		return
	}

	airspacelayerpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayerpropertytype.AirspaceLayer != nil {
		stage.UnstageBranch(airspacelayerpropertytype.AirspaceLayer)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacelayertype *AirspaceLayerType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacelayertype) {
		return
	}

	airspacelayertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacelayertype.DiscreteLevelSeries != nil {
		stage.UnstageBranch(airspacelayertype.DiscreteLevelSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacelayertype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_13 := range airspacelayertype.Extension {
		stage.UnstageBranch(_a_extension_13)
	}

}

func (airspacepropertytype *AirspacePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacepropertytype) {
		return
	}

	airspacepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacetimeslicepropertytype) {
		return
	}

	airspacetimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
		stage.UnstageBranch(airspacetimeslicepropertytype.AirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacetimeslicetype) {
		return
	}

	airspacetimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacetimeslicetype.ProtectedRoute != nil {
		stage.UnstageBranch(airspacetimeslicetype.ProtectedRoute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacelayerclasspropertytype := range airspacetimeslicetype.Class {
		stage.UnstageBranch(_airspacelayerclasspropertytype)
	}
	for _, _airspacegeometrycomponentpropertytype := range airspacetimeslicetype.GeometryComponent {
		stage.UnstageBranch(_airspacegeometrycomponentpropertytype)
	}
	for _, _airspaceactivationpropertytype := range airspacetimeslicetype.Activation {
		stage.UnstageBranch(_airspaceactivationpropertytype)
	}
	for _, _notepropertytype := range airspacetimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension := range airspacetimeslicetype.Extension {
		stage.UnstageBranch(_a_extension)
	}

}

func (airspacetype *AirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacetype) {
		return
	}

	airspacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _airspacetimeslicepropertytype := range airspacetype.TimeSlice {
		stage.UnstageBranch(_airspacetimeslicepropertytype)
	}

}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacevolumedependencypropertytype) {
		return
	}

	airspacevolumedependencypropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
		stage.UnstageBranch(airspacevolumedependencypropertytype.AirspaceVolumeDependency)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacevolumedependencytype) {
		return
	}

	airspacevolumedependencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumedependencytype.TheAirspace != nil {
		stage.UnstageBranch(airspacevolumedependencytype.TheAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumedependencytype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_5 := range airspacevolumedependencytype.Extension {
		stage.UnstageBranch(_a_extension_5)
	}

}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacevolumepropertytype) {
		return
	}

	airspacevolumepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumepropertytype.AirspaceVolume != nil {
		stage.UnstageBranch(airspacevolumepropertytype.AirspaceVolume)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (airspacevolumetype *AirspaceVolumeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(airspacevolumetype) {
		return
	}

	airspacevolumetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if airspacevolumetype.HorizontalProjection != nil {
		stage.UnstageBranch(airspacevolumetype.HorizontalProjection)
	}
	if airspacevolumetype.Centreline != nil {
		stage.UnstageBranch(airspacevolumetype.Centreline)
	}
	if airspacevolumetype.ContributorAirspace != nil {
		stage.UnstageBranch(airspacevolumetype.ContributorAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range airspacevolumetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_4 := range airspacevolumetype.Extension {
		stage.UnstageBranch(_a_extension_4)
	}

}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(authorityforairspacetimeslicepropertytype) {
		return
	}

	authorityforairspacetimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
		stage.UnstageBranch(authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(authorityforairspacetimeslicetype) {
		return
	}

	authorityforairspacetimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
		stage.UnstageBranch(authorityforairspacetimeslicetype.ResponsibleOrganisation)
	}
	if authorityforairspacetimeslicetype.AssignedAirspace != nil {
		stage.UnstageBranch(authorityforairspacetimeslicetype.AssignedAirspace)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range authorityforairspacetimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_6 := range authorityforairspacetimeslicetype.Extension {
		stage.UnstageBranch(_a_extension_6)
	}

}

func (authorityforairspacetype *AuthorityForAirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(authorityforairspacetype) {
		return
	}

	authorityforairspacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _authorityforairspacetimeslicepropertytype := range authorityforairspacetype.TimeSlice {
		stage.UnstageBranch(_authorityforairspacetimeslicepropertytype)
	}

}

func (curvepropertytype *CurvePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curvepropertytype) {
		return
	}

	curvepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if curvepropertytype.Curve != nil {
		stage.UnstageBranch(curvepropertytype.Curve)
	}
	if curvepropertytype.ElevatedCurve != nil {
		stage.UnstageBranch(curvepropertytype.ElevatedCurve)
	}

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
	for _, _a_extension_8 := range elevatedcurvetype.Extension {
		stage.UnstageBranch(_a_extension_8)
	}

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
	for _, _a_extension_9 := range elevatedsurfacetype.Extension {
		stage.UnstageBranch(_a_extension_9)
	}

}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geobordertimeslicepropertytype) {
		return
	}

	geobordertimeslicepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
		stage.UnstageBranch(geobordertimeslicepropertytype.GeoBorderTimeSlice)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geobordertimeslicetype) {
		return
	}

	geobordertimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if geobordertimeslicetype.Border != nil {
		stage.UnstageBranch(geobordertimeslicetype.Border)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _notepropertytype := range geobordertimeslicetype.Annotation {
		stage.UnstageBranch(_notepropertytype)
	}
	for _, _a_extension_7 := range geobordertimeslicetype.Extension {
		stage.UnstageBranch(_a_extension_7)
	}

}

func (geobordertype *GeoBorderType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(geobordertype) {
		return
	}

	geobordertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _geobordertimeslicepropertytype := range geobordertype.TimeSlice {
		stage.UnstageBranch(_geobordertimeslicepropertytype)
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
	for _, _a_extension_10 := range linguisticnotetype.Extension {
		stage.UnstageBranch(_a_extension_10)
	}

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
	for _, _a_extension_11 := range notetype.Extension {
		stage.UnstageBranch(_a_extension_11)
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

func (routepropertytype *RoutePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(routepropertytype) {
		return
	}

	routepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(standardlevelcolumnpropertytype) {
		return
	}

	standardlevelcolumnpropertytype.Unstage(stage)

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
	for _, _a_extension_14 := range timesheettype.Extension {
		stage.UnstageBranch(_a_extension_14)
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

func (reference *A_extension_13) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_13) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_extension_14) GongReconstructPointersFromReferences(stage *Stage, instance *A_extension_14) {
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

func (reference *AircraftCharacteristicPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AircraftCharacteristicPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AircraftCharacteristic, stage.AircraftCharacteristicTypes_reference, instance.AircraftCharacteristic)
	// insertion point for slice of pointers field
}

func (reference *AircraftCharacteristicType) GongReconstructPointersFromReferences(stage *Stage, instance *AircraftCharacteristicType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_12s_reference, instance.Extension)
}

func (reference *AirspaceActivationPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceActivationPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceActivation, stage.AirspaceActivationTypes_reference, instance.AirspaceActivation)
	// insertion point for slice of pointers field
}

func (reference *AirspaceActivationType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceActivationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeInterval, stage.TimesheetPropertyTypes_reference, instance.TimeInterval)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_reference, instance.SpecialDateAuthority)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Levels, stage.AirspaceLayerPropertyTypes_reference, instance.Levels)
	__gong__reconstructSliceOfPointersFromReferences(&reference.User, stage.OrganisationAuthorityPropertyTypes_reference, instance.User)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Aircraft, stage.AircraftCharacteristicPropertyTypes_reference, instance.Aircraft)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_1s_reference, instance.Extension)
}

func (reference *AirspaceGeometryComponentPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceGeometryComponentPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceGeometryComponent, stage.AirspaceGeometryComponentTypes_reference, instance.AirspaceGeometryComponent)
	// insertion point for slice of pointers field
}

func (reference *AirspaceGeometryComponentType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceGeometryComponentType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TheAirspaceVolume, stage.AirspaceVolumePropertyTypes_reference, instance.TheAirspaceVolume)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_2s_reference, instance.Extension)
}

func (reference *AirspaceLayerClassPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceLayerClassPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceLayerClass, stage.AirspaceLayerClassTypes_reference, instance.AirspaceLayerClass)
	// insertion point for slice of pointers field
}

func (reference *AirspaceLayerClassType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceLayerClassType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeInterval, stage.TimesheetPropertyTypes_reference, instance.TimeInterval)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_reference, instance.SpecialDateAuthority)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AssociatedLevels, stage.AirspaceLayerPropertyTypes_reference, instance.AssociatedLevels)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_3s_reference, instance.Extension)
}

func (reference *AirspaceLayerPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceLayerPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceLayer, stage.AirspaceLayerTypes_reference, instance.AirspaceLayer)
	// insertion point for slice of pointers field
}

func (reference *AirspaceLayerType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceLayerType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DiscreteLevelSeries, stage.StandardLevelColumnPropertyTypes_reference, instance.DiscreteLevelSeries)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_13s_reference, instance.Extension)
}

func (reference *AirspacePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspacePropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AirspaceTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceTimeSlice, stage.AirspaceTimeSliceTypes_reference, instance.AirspaceTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *AirspaceTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ProtectedRoute, stage.RoutePropertyTypes_reference, instance.ProtectedRoute)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Class, stage.AirspaceLayerClassPropertyTypes_reference, instance.Class)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeometryComponent, stage.AirspaceGeometryComponentPropertyTypes_reference, instance.GeometryComponent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Activation, stage.AirspaceActivationPropertyTypes_reference, instance.Activation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extensions_reference, instance.Extension)
}

func (reference *AirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.AirspaceTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

func (reference *AirspaceVolumeDependencyPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceVolumeDependencyPropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceVolumeDependency, stage.AirspaceVolumeDependencyTypes_reference, instance.AirspaceVolumeDependency)
	// insertion point for slice of pointers field
}

func (reference *AirspaceVolumeDependencyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceVolumeDependencyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TheAirspace, stage.AirspacePropertyTypes_reference, instance.TheAirspace)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_5s_reference, instance.Extension)
}

func (reference *AirspaceVolumePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceVolumePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AirspaceVolume, stage.AirspaceVolumeTypes_reference, instance.AirspaceVolume)
	// insertion point for slice of pointers field
}

func (reference *AirspaceVolumeType) GongReconstructPointersFromReferences(stage *Stage, instance *AirspaceVolumeType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.HorizontalProjection, stage.SurfacePropertyTypes_reference, instance.HorizontalProjection)
	__gong__reconstructPointer(&reference.Centreline, stage.CurvePropertyTypes_reference, instance.Centreline)
	__gong__reconstructPointer(&reference.ContributorAirspace, stage.AirspaceVolumeDependencyPropertyTypes_reference, instance.ContributorAirspace)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_4s_reference, instance.Extension)
}

func (reference *AuthorityForAirspaceTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AuthorityForAirspaceTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AuthorityForAirspaceTimeSlice, stage.AuthorityForAirspaceTimeSliceTypes_reference, instance.AuthorityForAirspaceTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *AuthorityForAirspaceTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AuthorityForAirspaceTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ResponsibleOrganisation, stage.OrganisationAuthorityPropertyTypes_reference, instance.ResponsibleOrganisation)
	__gong__reconstructPointer(&reference.AssignedAirspace, stage.AirspacePropertyTypes_reference, instance.AssignedAirspace)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_6s_reference, instance.Extension)
}

func (reference *AuthorityForAirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *AuthorityForAirspaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.AuthorityForAirspaceTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

func (reference *CurvePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *CurvePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Curve, stage.CurveTypes_reference, instance.Curve)
	__gong__reconstructPointer(&reference.ElevatedCurve, stage.ElevatedCurveTypes_reference, instance.ElevatedCurve)
	// insertion point for slice of pointers field
}

func (reference *CurveType) GongReconstructPointersFromReferences(stage *Stage, instance *CurveType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
}

func (reference *ElevatedCurveType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedCurveType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_8s_reference, instance.Extension)
}

func (reference *ElevatedSurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *ElevatedSurfaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_9s_reference, instance.Extension)
}

func (reference *GeoBorderTimeSlicePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *GeoBorderTimeSlicePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.GeoBorderTimeSlice, stage.GeoBorderTimeSliceTypes_reference, instance.GeoBorderTimeSlice)
	// insertion point for slice of pointers field
}

func (reference *GeoBorderTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *GeoBorderTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Border, stage.CurvePropertyTypes_reference, instance.Border)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Annotation, stage.NotePropertyTypes_reference, instance.Annotation)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_7s_reference, instance.Extension)
}

func (reference *GeoBorderType) GongReconstructPointersFromReferences(stage *Stage, instance *GeoBorderType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TimeSlice, stage.GeoBorderTimeSlicePropertyTypes_reference, instance.TimeSlice)
}

func (reference *LinguisticNotePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *LinguisticNotePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LinguisticNote, stage.LinguisticNoteTypes_reference, instance.LinguisticNote)
	// insertion point for slice of pointers field
}

func (reference *LinguisticNoteType) GongReconstructPointersFromReferences(stage *Stage, instance *LinguisticNoteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_10s_reference, instance.Extension)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_11s_reference, instance.Extension)
}

func (reference *OrganisationAuthorityPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *OrganisationAuthorityPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RoutePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *RoutePropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StandardLevelColumnPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *StandardLevelColumnPropertyType) {
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Extension, stage.A_extension_14s_reference, instance.Extension)
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

func (reference *A_extension_13) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_extension_14) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *AircraftCharacteristicPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AircraftCharacteristic, stage.AircraftCharacteristicTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AircraftCharacteristicType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_12s_instance)
}

func (reference *AirspaceActivationPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceActivation, stage.AirspaceActivationTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceActivationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeInterval, stage.TimesheetPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Levels, stage.AirspaceLayerPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.User, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Aircraft, stage.AircraftCharacteristicPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_1s_instance)
}

func (reference *AirspaceGeometryComponentPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceGeometryComponent, stage.AirspaceGeometryComponentTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceGeometryComponentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TheAirspaceVolume, stage.AirspaceVolumePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_2s_instance)
}

func (reference *AirspaceLayerClassPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceLayerClass, stage.AirspaceLayerClassTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceLayerClassType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeInterval, stage.TimesheetPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SpecialDateAuthority, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AssociatedLevels, stage.AirspaceLayerPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_3s_instance)
}

func (reference *AirspaceLayerPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceLayer, stage.AirspaceLayerTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceLayerType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DiscreteLevelSeries, stage.StandardLevelColumnPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_13s_instance)
}

func (reference *AirspacePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AirspaceTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceTimeSlice, stage.AirspaceTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ProtectedRoute, stage.RoutePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Class, stage.AirspaceLayerClassPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeometryComponent, stage.AirspaceGeometryComponentPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Activation, stage.AirspaceActivationPropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extensions_instance)
}

func (reference *AirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.AirspaceTimeSlicePropertyTypes_instance)
}

func (reference *AirspaceVolumeDependencyPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceVolumeDependency, stage.AirspaceVolumeDependencyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceVolumeDependencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TheAirspace, stage.AirspacePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_5s_instance)
}

func (reference *AirspaceVolumePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AirspaceVolume, stage.AirspaceVolumeTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AirspaceVolumeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.HorizontalProjection, stage.SurfacePropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Centreline, stage.CurvePropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ContributorAirspace, stage.AirspaceVolumeDependencyPropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_4s_instance)
}

func (reference *AuthorityForAirspaceTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AuthorityForAirspaceTimeSlice, stage.AuthorityForAirspaceTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AuthorityForAirspaceTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ResponsibleOrganisation, stage.OrganisationAuthorityPropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.AssignedAirspace, stage.AirspacePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_6s_instance)
}

func (reference *AuthorityForAirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.AuthorityForAirspaceTimeSlicePropertyTypes_instance)
}

func (reference *CurvePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Curve, stage.CurveTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ElevatedCurve, stage.ElevatedCurveTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *CurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
}

func (reference *ElevatedCurveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_8s_instance)
}

func (reference *ElevatedSurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_9s_instance)
}

func (reference *GeoBorderTimeSlicePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.GeoBorderTimeSlice, stage.GeoBorderTimeSliceTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *GeoBorderTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Border, stage.CurvePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Annotation, stage.NotePropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_7s_instance)
}

func (reference *GeoBorderType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TimeSlice, stage.GeoBorderTimeSlicePropertyTypes_instance)
}

func (reference *LinguisticNotePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LinguisticNote, stage.LinguisticNoteTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *LinguisticNoteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_10s_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_11s_instance)
}

func (reference *OrganisationAuthorityPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RoutePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StandardLevelColumnPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Extension, stage.A_extension_14s_instance)
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
func (a_extension_13 *A_extension_13) GongDiff(stage *Stage, a_extension_13Other *A_extension_13) (diffs []string) {
	// insertion point for field diffs
	if a_extension_13.Name != a_extension_13Other.Name {
		diffs = append(diffs, a_extension_13.GongMarshallField(stage, "Name"))
	}
	if a_extension_13.Owns != a_extension_13Other.Owns {
		diffs = append(diffs, a_extension_13.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_extension_14 *A_extension_14) GongDiff(stage *Stage, a_extension_14Other *A_extension_14) (diffs []string) {
	// insertion point for field diffs
	if a_extension_14.Name != a_extension_14Other.Name {
		diffs = append(diffs, a_extension_14.GongMarshallField(stage, "Name"))
	}
	if a_extension_14.Owns != a_extension_14Other.Owns {
		diffs = append(diffs, a_extension_14.GongMarshallField(stage, "Owns"))
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
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongDiff(stage *Stage, aircraftcharacteristicpropertytypeOther *AircraftCharacteristicPropertyType) (diffs []string) {
	// insertion point for field diffs
	if aircraftcharacteristicpropertytype.Name != aircraftcharacteristicpropertytypeOther.Name {
		diffs = append(diffs, aircraftcharacteristicpropertytype.GongMarshallField(stage, "Name"))
	}
	if aircraftcharacteristicpropertytype.NilReason != aircraftcharacteristicpropertytypeOther.NilReason {
		diffs = append(diffs, aircraftcharacteristicpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if aircraftcharacteristicpropertytype.AircraftCharacteristic != aircraftcharacteristicpropertytypeOther.AircraftCharacteristic {
		diffs = append(diffs, aircraftcharacteristicpropertytype.GongMarshallField(stage, "AircraftCharacteristic"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aircraftcharacteristictype *AircraftCharacteristicType) GongDiff(stage *Stage, aircraftcharacteristictypeOther *AircraftCharacteristicType) (diffs []string) {
	// insertion point for field diffs
	if aircraftcharacteristictype.Name != aircraftcharacteristictypeOther.Name {
		diffs = append(diffs, aircraftcharacteristictype.GongMarshallField(stage, "Name"))
	}
	if aircraftcharacteristictype.Id != aircraftcharacteristictypeOther.Id {
		diffs = append(diffs, aircraftcharacteristictype.GongMarshallField(stage, "Id"))
	}
	if aircraftcharacteristictype.Description != aircraftcharacteristictypeOther.Description {
		diffs = append(diffs, aircraftcharacteristictype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, aircraftcharacteristictype, "Annotation", aircraftcharacteristictypeOther.Annotation, aircraftcharacteristictype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, aircraftcharacteristictype, "Extension", aircraftcharacteristictypeOther.Extension, aircraftcharacteristictype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongDiff(stage *Stage, airspaceactivationpropertytypeOther *AirspaceActivationPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspaceactivationpropertytype.Name != airspaceactivationpropertytypeOther.Name {
		diffs = append(diffs, airspaceactivationpropertytype.GongMarshallField(stage, "Name"))
	}
	if airspaceactivationpropertytype.NilReason != airspaceactivationpropertytypeOther.NilReason {
		diffs = append(diffs, airspaceactivationpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspaceactivationpropertytype.AirspaceActivation != airspaceactivationpropertytypeOther.AirspaceActivation {
		diffs = append(diffs, airspaceactivationpropertytype.GongMarshallField(stage, "AirspaceActivation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspaceactivationtype *AirspaceActivationType) GongDiff(stage *Stage, airspaceactivationtypeOther *AirspaceActivationType) (diffs []string) {
	// insertion point for field diffs
	if airspaceactivationtype.Name != airspaceactivationtypeOther.Name {
		diffs = append(diffs, airspaceactivationtype.GongMarshallField(stage, "Name"))
	}
	if airspaceactivationtype.Id != airspaceactivationtypeOther.Id {
		diffs = append(diffs, airspaceactivationtype.GongMarshallField(stage, "Id"))
	}
	if airspaceactivationtype.Description != airspaceactivationtypeOther.Description {
		diffs = append(diffs, airspaceactivationtype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "TimeInterval", airspaceactivationtypeOther.TimeInterval, airspaceactivationtype.TimeInterval); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "Annotation", airspaceactivationtypeOther.Annotation, airspaceactivationtype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "SpecialDateAuthority", airspaceactivationtypeOther.SpecialDateAuthority, airspaceactivationtype.SpecialDateAuthority); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "Levels", airspaceactivationtypeOther.Levels, airspaceactivationtype.Levels); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "User", airspaceactivationtypeOther.User, airspaceactivationtype.User); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "Aircraft", airspaceactivationtypeOther.Aircraft, airspaceactivationtype.Aircraft); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspaceactivationtype, "Extension", airspaceactivationtypeOther.Extension, airspaceactivationtype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongDiff(stage *Stage, airspacegeometrycomponentpropertytypeOther *AirspaceGeometryComponentPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacegeometrycomponentpropertytype.Name != airspacegeometrycomponentpropertytypeOther.Name {
		diffs = append(diffs, airspacegeometrycomponentpropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacegeometrycomponentpropertytype.NilReason != airspacegeometrycomponentpropertytypeOther.NilReason {
		diffs = append(diffs, airspacegeometrycomponentpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != airspacegeometrycomponentpropertytypeOther.AirspaceGeometryComponent {
		diffs = append(diffs, airspacegeometrycomponentpropertytype.GongMarshallField(stage, "AirspaceGeometryComponent"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongDiff(stage *Stage, airspacegeometrycomponenttypeOther *AirspaceGeometryComponentType) (diffs []string) {
	// insertion point for field diffs
	if airspacegeometrycomponenttype.Name != airspacegeometrycomponenttypeOther.Name {
		diffs = append(diffs, airspacegeometrycomponenttype.GongMarshallField(stage, "Name"))
	}
	if airspacegeometrycomponenttype.Id != airspacegeometrycomponenttypeOther.Id {
		diffs = append(diffs, airspacegeometrycomponenttype.GongMarshallField(stage, "Id"))
	}
	if airspacegeometrycomponenttype.Description != airspacegeometrycomponenttypeOther.Description {
		diffs = append(diffs, airspacegeometrycomponenttype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacegeometrycomponenttype, "Annotation", airspacegeometrycomponenttypeOther.Annotation, airspacegeometrycomponenttype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if airspacegeometrycomponenttype.TheAirspaceVolume != airspacegeometrycomponenttypeOther.TheAirspaceVolume {
		diffs = append(diffs, airspacegeometrycomponenttype.GongMarshallField(stage, "TheAirspaceVolume"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacegeometrycomponenttype, "Extension", airspacegeometrycomponenttypeOther.Extension, airspacegeometrycomponenttype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongDiff(stage *Stage, airspacelayerclasspropertytypeOther *AirspaceLayerClassPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacelayerclasspropertytype.Name != airspacelayerclasspropertytypeOther.Name {
		diffs = append(diffs, airspacelayerclasspropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacelayerclasspropertytype.NilReason != airspacelayerclasspropertytypeOther.NilReason {
		diffs = append(diffs, airspacelayerclasspropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacelayerclasspropertytype.AirspaceLayerClass != airspacelayerclasspropertytypeOther.AirspaceLayerClass {
		diffs = append(diffs, airspacelayerclasspropertytype.GongMarshallField(stage, "AirspaceLayerClass"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacelayerclasstype *AirspaceLayerClassType) GongDiff(stage *Stage, airspacelayerclasstypeOther *AirspaceLayerClassType) (diffs []string) {
	// insertion point for field diffs
	if airspacelayerclasstype.Name != airspacelayerclasstypeOther.Name {
		diffs = append(diffs, airspacelayerclasstype.GongMarshallField(stage, "Name"))
	}
	if airspacelayerclasstype.Id != airspacelayerclasstypeOther.Id {
		diffs = append(diffs, airspacelayerclasstype.GongMarshallField(stage, "Id"))
	}
	if airspacelayerclasstype.Description != airspacelayerclasstypeOther.Description {
		diffs = append(diffs, airspacelayerclasstype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayerclasstype, "TimeInterval", airspacelayerclasstypeOther.TimeInterval, airspacelayerclasstype.TimeInterval); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayerclasstype, "Annotation", airspacelayerclasstypeOther.Annotation, airspacelayerclasstype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayerclasstype, "SpecialDateAuthority", airspacelayerclasstypeOther.SpecialDateAuthority, airspacelayerclasstype.SpecialDateAuthority); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayerclasstype, "AssociatedLevels", airspacelayerclasstypeOther.AssociatedLevels, airspacelayerclasstype.AssociatedLevels); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayerclasstype, "Extension", airspacelayerclasstypeOther.Extension, airspacelayerclasstype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongDiff(stage *Stage, airspacelayerpropertytypeOther *AirspaceLayerPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacelayerpropertytype.Name != airspacelayerpropertytypeOther.Name {
		diffs = append(diffs, airspacelayerpropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacelayerpropertytype.NilReason != airspacelayerpropertytypeOther.NilReason {
		diffs = append(diffs, airspacelayerpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacelayerpropertytype.AirspaceLayer != airspacelayerpropertytypeOther.AirspaceLayer {
		diffs = append(diffs, airspacelayerpropertytype.GongMarshallField(stage, "AirspaceLayer"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacelayertype *AirspaceLayerType) GongDiff(stage *Stage, airspacelayertypeOther *AirspaceLayerType) (diffs []string) {
	// insertion point for field diffs
	if airspacelayertype.Name != airspacelayertypeOther.Name {
		diffs = append(diffs, airspacelayertype.GongMarshallField(stage, "Name"))
	}
	if airspacelayertype.Id != airspacelayertypeOther.Id {
		diffs = append(diffs, airspacelayertype.GongMarshallField(stage, "Id"))
	}
	if airspacelayertype.Description != airspacelayertypeOther.Description {
		diffs = append(diffs, airspacelayertype.GongMarshallField(stage, "Description"))
	}
	if airspacelayertype.DiscreteLevelSeries != airspacelayertypeOther.DiscreteLevelSeries {
		diffs = append(diffs, airspacelayertype.GongMarshallField(stage, "DiscreteLevelSeries"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayertype, "Annotation", airspacelayertypeOther.Annotation, airspacelayertype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacelayertype, "Extension", airspacelayertypeOther.Extension, airspacelayertype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacepropertytype *AirspacePropertyType) GongDiff(stage *Stage, airspacepropertytypeOther *AirspacePropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacepropertytype.Name != airspacepropertytypeOther.Name {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacepropertytype.Owns != airspacepropertytypeOther.Owns {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Owns"))
	}
	if airspacepropertytype.NilReason != airspacepropertytypeOther.NilReason {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacepropertytype.Type_XSD != airspacepropertytypeOther.Type_XSD {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if airspacepropertytype.Href != airspacepropertytypeOther.Href {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Href"))
	}
	if airspacepropertytype.Role != airspacepropertytypeOther.Role {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Role"))
	}
	if airspacepropertytype.Arcrole != airspacepropertytypeOther.Arcrole {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if airspacepropertytype.Title != airspacepropertytypeOther.Title {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Title"))
	}
	if airspacepropertytype.Show != airspacepropertytypeOther.Show {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Show"))
	}
	if airspacepropertytype.Actuate != airspacepropertytypeOther.Actuate {
		diffs = append(diffs, airspacepropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongDiff(stage *Stage, airspacetimeslicepropertytypeOther *AirspaceTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacetimeslicepropertytype.Name != airspacetimeslicepropertytypeOther.Name {
		diffs = append(diffs, airspacetimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacetimeslicepropertytype.Owns != airspacetimeslicepropertytypeOther.Owns {
		diffs = append(diffs, airspacetimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if airspacetimeslicepropertytype.AirspaceTimeSlice != airspacetimeslicepropertytypeOther.AirspaceTimeSlice {
		diffs = append(diffs, airspacetimeslicepropertytype.GongMarshallField(stage, "AirspaceTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacetimeslicetype *AirspaceTimeSliceType) GongDiff(stage *Stage, airspacetimeslicetypeOther *AirspaceTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if airspacetimeslicetype.Name != airspacetimeslicetypeOther.Name {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "Name"))
	}
	if airspacetimeslicetype.Id != airspacetimeslicetypeOther.Id {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "Id"))
	}
	if airspacetimeslicetype.Description != airspacetimeslicetypeOther.Description {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "Description"))
	}
	if airspacetimeslicetype.Interpretation != airspacetimeslicetypeOther.Interpretation {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if airspacetimeslicetype.SequenceNumber != airspacetimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if airspacetimeslicetype.CorrectionNumber != airspacetimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetimeslicetype, "Class", airspacetimeslicetypeOther.Class, airspacetimeslicetype.Class); ops != "" {
		diffs = append(diffs, ops)
	}
	if airspacetimeslicetype.ProtectedRoute != airspacetimeslicetypeOther.ProtectedRoute {
		diffs = append(diffs, airspacetimeslicetype.GongMarshallField(stage, "ProtectedRoute"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetimeslicetype, "GeometryComponent", airspacetimeslicetypeOther.GeometryComponent, airspacetimeslicetype.GeometryComponent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetimeslicetype, "Activation", airspacetimeslicetypeOther.Activation, airspacetimeslicetype.Activation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetimeslicetype, "Annotation", airspacetimeslicetypeOther.Annotation, airspacetimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetimeslicetype, "Extension", airspacetimeslicetypeOther.Extension, airspacetimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacetype *AirspaceType) GongDiff(stage *Stage, airspacetypeOther *AirspaceType) (diffs []string) {
	// insertion point for field diffs
	if airspacetype.Name != airspacetypeOther.Name {
		diffs = append(diffs, airspacetype.GongMarshallField(stage, "Name"))
	}
	if airspacetype.Id != airspacetypeOther.Id {
		diffs = append(diffs, airspacetype.GongMarshallField(stage, "Id"))
	}
	if airspacetype.Description != airspacetypeOther.Description {
		diffs = append(diffs, airspacetype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacetype, "TimeSlice", airspacetypeOther.TimeSlice, airspacetype.TimeSlice); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongDiff(stage *Stage, airspacevolumedependencypropertytypeOther *AirspaceVolumeDependencyPropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacevolumedependencypropertytype.Name != airspacevolumedependencypropertytypeOther.Name {
		diffs = append(diffs, airspacevolumedependencypropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacevolumedependencypropertytype.NilReason != airspacevolumedependencypropertytypeOther.NilReason {
		diffs = append(diffs, airspacevolumedependencypropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacevolumedependencypropertytype.AirspaceVolumeDependency != airspacevolumedependencypropertytypeOther.AirspaceVolumeDependency {
		diffs = append(diffs, airspacevolumedependencypropertytype.GongMarshallField(stage, "AirspaceVolumeDependency"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongDiff(stage *Stage, airspacevolumedependencytypeOther *AirspaceVolumeDependencyType) (diffs []string) {
	// insertion point for field diffs
	if airspacevolumedependencytype.Name != airspacevolumedependencytypeOther.Name {
		diffs = append(diffs, airspacevolumedependencytype.GongMarshallField(stage, "Name"))
	}
	if airspacevolumedependencytype.Id != airspacevolumedependencytypeOther.Id {
		diffs = append(diffs, airspacevolumedependencytype.GongMarshallField(stage, "Id"))
	}
	if airspacevolumedependencytype.Description != airspacevolumedependencytypeOther.Description {
		diffs = append(diffs, airspacevolumedependencytype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacevolumedependencytype, "Annotation", airspacevolumedependencytypeOther.Annotation, airspacevolumedependencytype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if airspacevolumedependencytype.TheAirspace != airspacevolumedependencytypeOther.TheAirspace {
		diffs = append(diffs, airspacevolumedependencytype.GongMarshallField(stage, "TheAirspace"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacevolumedependencytype, "Extension", airspacevolumedependencytypeOther.Extension, airspacevolumedependencytype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongDiff(stage *Stage, airspacevolumepropertytypeOther *AirspaceVolumePropertyType) (diffs []string) {
	// insertion point for field diffs
	if airspacevolumepropertytype.Name != airspacevolumepropertytypeOther.Name {
		diffs = append(diffs, airspacevolumepropertytype.GongMarshallField(stage, "Name"))
	}
	if airspacevolumepropertytype.NilReason != airspacevolumepropertytypeOther.NilReason {
		diffs = append(diffs, airspacevolumepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if airspacevolumepropertytype.AirspaceVolume != airspacevolumepropertytypeOther.AirspaceVolume {
		diffs = append(diffs, airspacevolumepropertytype.GongMarshallField(stage, "AirspaceVolume"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (airspacevolumetype *AirspaceVolumeType) GongDiff(stage *Stage, airspacevolumetypeOther *AirspaceVolumeType) (diffs []string) {
	// insertion point for field diffs
	if airspacevolumetype.Name != airspacevolumetypeOther.Name {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "Name"))
	}
	if airspacevolumetype.Id != airspacevolumetypeOther.Id {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "Id"))
	}
	if airspacevolumetype.Description != airspacevolumetypeOther.Description {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "Description"))
	}
	if airspacevolumetype.HorizontalProjection != airspacevolumetypeOther.HorizontalProjection {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "HorizontalProjection"))
	}
	if airspacevolumetype.Centreline != airspacevolumetypeOther.Centreline {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "Centreline"))
	}
	if airspacevolumetype.ContributorAirspace != airspacevolumetypeOther.ContributorAirspace {
		diffs = append(diffs, airspacevolumetype.GongMarshallField(stage, "ContributorAirspace"))
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacevolumetype, "Annotation", airspacevolumetypeOther.Annotation, airspacevolumetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, airspacevolumetype, "Extension", airspacevolumetypeOther.Extension, airspacevolumetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongDiff(stage *Stage, authorityforairspacetimeslicepropertytypeOther *AuthorityForAirspaceTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if authorityforairspacetimeslicepropertytype.Name != authorityforairspacetimeslicepropertytypeOther.Name {
		diffs = append(diffs, authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if authorityforairspacetimeslicepropertytype.Owns != authorityforairspacetimeslicepropertytypeOther.Owns {
		diffs = append(diffs, authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != authorityforairspacetimeslicepropertytypeOther.AuthorityForAirspaceTimeSlice {
		diffs = append(diffs, authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "AuthorityForAirspaceTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongDiff(stage *Stage, authorityforairspacetimeslicetypeOther *AuthorityForAirspaceTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if authorityforairspacetimeslicetype.Name != authorityforairspacetimeslicetypeOther.Name {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "Name"))
	}
	if authorityforairspacetimeslicetype.Id != authorityforairspacetimeslicetypeOther.Id {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "Id"))
	}
	if authorityforairspacetimeslicetype.Description != authorityforairspacetimeslicetypeOther.Description {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "Description"))
	}
	if authorityforairspacetimeslicetype.Interpretation != authorityforairspacetimeslicetypeOther.Interpretation {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if authorityforairspacetimeslicetype.SequenceNumber != authorityforairspacetimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if authorityforairspacetimeslicetype.CorrectionNumber != authorityforairspacetimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if authorityforairspacetimeslicetype.ResponsibleOrganisation != authorityforairspacetimeslicetypeOther.ResponsibleOrganisation {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "ResponsibleOrganisation"))
	}
	if authorityforairspacetimeslicetype.AssignedAirspace != authorityforairspacetimeslicetypeOther.AssignedAirspace {
		diffs = append(diffs, authorityforairspacetimeslicetype.GongMarshallField(stage, "AssignedAirspace"))
	}
	if ops := __gong__diffSliceOfPointers(stage, authorityforairspacetimeslicetype, "Annotation", authorityforairspacetimeslicetypeOther.Annotation, authorityforairspacetimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, authorityforairspacetimeslicetype, "Extension", authorityforairspacetimeslicetypeOther.Extension, authorityforairspacetimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (authorityforairspacetype *AuthorityForAirspaceType) GongDiff(stage *Stage, authorityforairspacetypeOther *AuthorityForAirspaceType) (diffs []string) {
	// insertion point for field diffs
	if authorityforairspacetype.Name != authorityforairspacetypeOther.Name {
		diffs = append(diffs, authorityforairspacetype.GongMarshallField(stage, "Name"))
	}
	if authorityforairspacetype.Id != authorityforairspacetypeOther.Id {
		diffs = append(diffs, authorityforairspacetype.GongMarshallField(stage, "Id"))
	}
	if authorityforairspacetype.Description != authorityforairspacetypeOther.Description {
		diffs = append(diffs, authorityforairspacetype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, authorityforairspacetype, "TimeSlice", authorityforairspacetypeOther.TimeSlice, authorityforairspacetype.TimeSlice); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curvepropertytype *CurvePropertyType) GongDiff(stage *Stage, curvepropertytypeOther *CurvePropertyType) (diffs []string) {
	// insertion point for field diffs
	if curvepropertytype.Name != curvepropertytypeOther.Name {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Name"))
	}
	if curvepropertytype.NilReason != curvepropertytypeOther.NilReason {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if curvepropertytype.Curve != curvepropertytypeOther.Curve {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "Curve"))
	}
	if curvepropertytype.ElevatedCurve != curvepropertytypeOther.ElevatedCurve {
		diffs = append(diffs, curvepropertytype.GongMarshallField(stage, "ElevatedCurve"))
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
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongDiff(stage *Stage, geobordertimeslicepropertytypeOther *GeoBorderTimeSlicePropertyType) (diffs []string) {
	// insertion point for field diffs
	if geobordertimeslicepropertytype.Name != geobordertimeslicepropertytypeOther.Name {
		diffs = append(diffs, geobordertimeslicepropertytype.GongMarshallField(stage, "Name"))
	}
	if geobordertimeslicepropertytype.Owns != geobordertimeslicepropertytypeOther.Owns {
		diffs = append(diffs, geobordertimeslicepropertytype.GongMarshallField(stage, "Owns"))
	}
	if geobordertimeslicepropertytype.GeoBorderTimeSlice != geobordertimeslicepropertytypeOther.GeoBorderTimeSlice {
		diffs = append(diffs, geobordertimeslicepropertytype.GongMarshallField(stage, "GeoBorderTimeSlice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geobordertimeslicetype *GeoBorderTimeSliceType) GongDiff(stage *Stage, geobordertimeslicetypeOther *GeoBorderTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if geobordertimeslicetype.Name != geobordertimeslicetypeOther.Name {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "Name"))
	}
	if geobordertimeslicetype.Id != geobordertimeslicetypeOther.Id {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "Id"))
	}
	if geobordertimeslicetype.Description != geobordertimeslicetypeOther.Description {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "Description"))
	}
	if geobordertimeslicetype.Interpretation != geobordertimeslicetypeOther.Interpretation {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if geobordertimeslicetype.SequenceNumber != geobordertimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if geobordertimeslicetype.CorrectionNumber != geobordertimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if geobordertimeslicetype.Border != geobordertimeslicetypeOther.Border {
		diffs = append(diffs, geobordertimeslicetype.GongMarshallField(stage, "Border"))
	}
	if ops := __gong__diffSliceOfPointers(stage, geobordertimeslicetype, "Annotation", geobordertimeslicetypeOther.Annotation, geobordertimeslicetype.Annotation); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, geobordertimeslicetype, "Extension", geobordertimeslicetypeOther.Extension, geobordertimeslicetype.Extension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (geobordertype *GeoBorderType) GongDiff(stage *Stage, geobordertypeOther *GeoBorderType) (diffs []string) {
	// insertion point for field diffs
	if geobordertype.Name != geobordertypeOther.Name {
		diffs = append(diffs, geobordertype.GongMarshallField(stage, "Name"))
	}
	if geobordertype.Id != geobordertypeOther.Id {
		diffs = append(diffs, geobordertype.GongMarshallField(stage, "Id"))
	}
	if geobordertype.Description != geobordertypeOther.Description {
		diffs = append(diffs, geobordertype.GongMarshallField(stage, "Description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, geobordertype, "TimeSlice", geobordertypeOther.TimeSlice, geobordertype.TimeSlice); ops != "" {
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
func (routepropertytype *RoutePropertyType) GongDiff(stage *Stage, routepropertytypeOther *RoutePropertyType) (diffs []string) {
	// insertion point for field diffs
	if routepropertytype.Name != routepropertytypeOther.Name {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Name"))
	}
	if routepropertytype.Owns != routepropertytypeOther.Owns {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Owns"))
	}
	if routepropertytype.NilReason != routepropertytypeOther.NilReason {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if routepropertytype.Type_XSD != routepropertytypeOther.Type_XSD {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if routepropertytype.Href != routepropertytypeOther.Href {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Href"))
	}
	if routepropertytype.Role != routepropertytypeOther.Role {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Role"))
	}
	if routepropertytype.Arcrole != routepropertytypeOther.Arcrole {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if routepropertytype.Title != routepropertytypeOther.Title {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Title"))
	}
	if routepropertytype.Show != routepropertytypeOther.Show {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Show"))
	}
	if routepropertytype.Actuate != routepropertytypeOther.Actuate {
		diffs = append(diffs, routepropertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongDiff(stage *Stage, standardlevelcolumnpropertytypeOther *StandardLevelColumnPropertyType) (diffs []string) {
	// insertion point for field diffs
	if standardlevelcolumnpropertytype.Name != standardlevelcolumnpropertytypeOther.Name {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Name"))
	}
	if standardlevelcolumnpropertytype.Owns != standardlevelcolumnpropertytypeOther.Owns {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Owns"))
	}
	if standardlevelcolumnpropertytype.NilReason != standardlevelcolumnpropertytypeOther.NilReason {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if standardlevelcolumnpropertytype.Type_XSD != standardlevelcolumnpropertytypeOther.Type_XSD {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if standardlevelcolumnpropertytype.Href != standardlevelcolumnpropertytypeOther.Href {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Href"))
	}
	if standardlevelcolumnpropertytype.Role != standardlevelcolumnpropertytypeOther.Role {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Role"))
	}
	if standardlevelcolumnpropertytype.Arcrole != standardlevelcolumnpropertytypeOther.Arcrole {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if standardlevelcolumnpropertytype.Title != standardlevelcolumnpropertytypeOther.Title {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Title"))
	}
	if standardlevelcolumnpropertytype.Show != standardlevelcolumnpropertytypeOther.Show {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Show"))
	}
	if standardlevelcolumnpropertytype.Actuate != standardlevelcolumnpropertytypeOther.Actuate {
		diffs = append(diffs, standardlevelcolumnpropertytype.GongMarshallField(stage, "Actuate"))
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
