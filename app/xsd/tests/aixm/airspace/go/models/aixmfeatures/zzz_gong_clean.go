// generated code - do not edit
package aixmfeatures

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AircraftCharacteristicPropertyType
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aircraftcharacteristicpropertytype.AircraftCharacteristic) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftCharacteristicType
func (aircraftcharacteristictype *AircraftCharacteristicType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftcharacteristictype.Annotation) || modified
	modified = stage.CleanSlice(&aircraftcharacteristictype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceActivationPropertyType
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspaceactivationpropertytype.AirspaceActivation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceActivationType
func (airspaceactivationtype *AirspaceActivationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspaceactivationtype.TimeInterval) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.Annotation) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.Levels) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.User) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.Aircraft) || modified
	modified = stage.CleanSlice(&airspaceactivationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceGeometryComponentPropertyType
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacegeometrycomponentpropertytype.AirspaceGeometryComponent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceGeometryComponentType
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacegeometrycomponenttype.Annotation) || modified
	modified = stage.CleanSlice(&airspacegeometrycomponenttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacegeometrycomponenttype.TheAirspaceVolume) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceLayerClassPropertyType
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacelayerclasspropertytype.AirspaceLayerClass) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceLayerClassType
func (airspacelayerclasstype *AirspaceLayerClassType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacelayerclasstype.TimeInterval) || modified
	modified = stage.CleanSlice(&airspacelayerclasstype.Annotation) || modified
	modified = stage.CleanSlice(&airspacelayerclasstype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&airspacelayerclasstype.AssociatedLevels) || modified
	modified = stage.CleanSlice(&airspacelayerclasstype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceLayerPropertyType
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacelayerpropertytype.AirspaceLayer) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceLayerType
func (airspacelayertype *AirspaceLayerType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacelayertype.Annotation) || modified
	modified = stage.CleanSlice(&airspacelayertype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacelayertype.DiscreteLevelSeries) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceTimeSlicePropertyType
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacetimeslicepropertytype.AirspaceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceTimeSliceType
func (airspacetimeslicetype *AirspaceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacetimeslicetype.Class) || modified
	modified = stage.CleanSlice(&airspacetimeslicetype.GeometryComponent) || modified
	modified = stage.CleanSlice(&airspacetimeslicetype.Activation) || modified
	modified = stage.CleanSlice(&airspacetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airspacetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacetimeslicetype.ProtectedRoute) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceType
func (airspacetype *AirspaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceVolumeDependencyPropertyType
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacevolumedependencypropertytype.AirspaceVolumeDependency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceVolumeDependencyType
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacevolumedependencytype.Annotation) || modified
	modified = stage.CleanSlice(&airspacevolumedependencytype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacevolumedependencytype.TheAirspace) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceVolumePropertyType
func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacevolumepropertytype.AirspaceVolume) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceVolumeType
func (airspacevolumetype *AirspaceVolumeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacevolumetype.Annotation) || modified
	modified = stage.CleanSlice(&airspacevolumetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacevolumetype.HorizontalProjection) || modified
	modified = stage.CleanPointer(&airspacevolumetype.Centreline) || modified
	modified = stage.CleanPointer(&airspacevolumetype.ContributorAirspace) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForAirspaceTimeSlicePropertyType
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForAirspaceTimeSliceType
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityforairspacetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&authorityforairspacetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&authorityforairspacetimeslicetype.ResponsibleOrganisation) || modified
	modified = stage.CleanPointer(&authorityforairspacetimeslicetype.AssignedAirspace) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForAirspaceType
func (authorityforairspacetype *AuthorityForAirspaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityforairspacetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CurvePropertyType
func (curvepropertytype *CurvePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&curvepropertytype.Curve) || modified
	modified = stage.CleanPointer(&curvepropertytype.ElevatedCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CurveType
func (curvetype *CurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvetype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevatedCurveType
func (elevatedcurvetype *ElevatedCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevatedcurvetype.Annotation) || modified
	modified = stage.CleanSlice(&elevatedcurvetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevatedSurfaceType
func (elevatedsurfacetype *ElevatedSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevatedsurfacetype.Annotation) || modified
	modified = stage.CleanSlice(&elevatedsurfacetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GeoBorderTimeSlicePropertyType
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geobordertimeslicepropertytype.GeoBorderTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeoBorderTimeSliceType
func (geobordertimeslicetype *GeoBorderTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geobordertimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&geobordertimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geobordertimeslicetype.Border) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeoBorderType
func (geobordertype *GeoBorderType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geobordertype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LinguisticNotePropertyType
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&linguisticnotepropertytype.LinguisticNote) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinguisticNoteType
func (linguisticnotetype *LinguisticNoteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linguisticnotetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePropertyType
func (notepropertytype *NotePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notepropertytype.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteType
func (notetype *NoteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&notetype.TranslatedNote) || modified
	modified = stage.CleanSlice(&notetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfacePropertyType
func (surfacepropertytype *SurfacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surfacepropertytype.Surface) || modified
	modified = stage.CleanPointer(&surfacepropertytype.ElevatedSurface) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceType
func (surfacetype *SurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacetype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TimesheetPropertyType
func (timesheetpropertytype *TimesheetPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timesheetpropertytype.Timesheet) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimesheetType
func (timesheettype *TimesheetType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timesheettype.Annotation) || modified
	modified = stage.CleanSlice(&timesheettype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
