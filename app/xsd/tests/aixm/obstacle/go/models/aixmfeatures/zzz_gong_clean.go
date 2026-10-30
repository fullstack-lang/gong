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
// Clean garbage collect unstaged instances that are referenced by AeronauticalGroundLightTimeSlicePropertyType
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AeronauticalGroundLightTimeSliceType
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aeronauticalgroundlighttimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&aeronauticalgroundlighttimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aeronauticalgroundlighttimeslicetype.StructureBeacon) || modified
	modified = stage.CleanPointer(&aeronauticalgroundlighttimeslicetype.AerodromeBeacon) || modified
	modified = stage.CleanPointer(&aeronauticalgroundlighttimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AeronauticalGroundLightType
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aeronauticalgroundlighttype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CurveType
func (curvetype *CurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvetype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevatedCurvePropertyType
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&elevatedcurvepropertytype.ElevatedCurve) || modified
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

// Clean garbage collect unstaged instances that are referenced by ElevatedPointPropertyType
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&elevatedpointpropertytype.ElevatedPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevatedPointType
func (elevatedpointtype *ElevatedPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevatedpointtype.Annotation) || modified
	modified = stage.CleanSlice(&elevatedpointtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevatedSurfacePropertyType
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&elevatedsurfacepropertytype.ElevatedSurface) || modified
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

// Clean garbage collect unstaged instances that are referenced by LightElementPropertyType
func (lightelementpropertytype *LightElementPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&lightelementpropertytype.LightElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LightElementStatusPropertyType
func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&lightelementstatuspropertytype.LightElementStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LightElementStatusType
func (lightelementstatustype *LightElementStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&lightelementstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&lightelementstatustype.Annotation) || modified
	modified = stage.CleanSlice(&lightelementstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&lightelementstatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LightElementType
func (lightelementtype *LightElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&lightelementtype.Annotation) || modified
	modified = stage.CleanSlice(&lightelementtype.Availability) || modified
	modified = stage.CleanSlice(&lightelementtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&lightelementtype.Location) || modified
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

// Clean garbage collect unstaged instances that are referenced by ObstacleAreaTimeSlicePropertyType
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstacleAreaTimeSliceType
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obstacleareatimeslicetype.Obstacle) || modified
	modified = stage.CleanSlice(&obstacleareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&obstacleareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&obstacleareatimeslicetype.Reference_ownerOrganisation) || modified
	modified = stage.CleanPointer(&obstacleareatimeslicetype.Reference_ownerRunway) || modified
	modified = stage.CleanPointer(&obstacleareatimeslicetype.Reference_ownerAirport) || modified
	modified = stage.CleanPointer(&obstacleareatimeslicetype.SurfaceExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstacleAreaType
func (obstacleareatype *ObstacleAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obstacleareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PointType
func (pointtype *PointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointtype.Annotation) || modified
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

// Clean garbage collect unstaged instances that are referenced by VerticalStructureLightingStatusPropertyType
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructureLightingStatusType
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalstructurelightingstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&verticalstructurelightingstatustype.Annotation) || modified
	modified = stage.CleanSlice(&verticalstructurelightingstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&verticalstructurelightingstatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructurePartPropertyType
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticalstructurepartpropertytype.VerticalStructurePart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructurePartType
func (verticalstructureparttype *VerticalStructurePartType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalstructureparttype.TimeInterval) || modified
	modified = stage.CleanSlice(&verticalstructureparttype.Annotation) || modified
	modified = stage.CleanSlice(&verticalstructureparttype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&verticalstructureparttype.Lighting) || modified
	modified = stage.CleanSlice(&verticalstructureparttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&verticalstructureparttype.HorizontalProjection_surfaceExtent) || modified
	modified = stage.CleanPointer(&verticalstructureparttype.HorizontalProjection_linearExtent) || modified
	modified = stage.CleanPointer(&verticalstructureparttype.HorizontalProjection_location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructureTimeSlicePropertyType
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructureTimeSliceType
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.Part) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.HostedPassengerService) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.SupportedGroundLight) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.HostedNavaidEquipment) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.HostedSpecialNavStation) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.HostedUnit) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.HostedOrganisation) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.SupportedService) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.LightingAvailability) || modified
	modified = stage.CleanSlice(&verticalstructuretimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&verticalstructuretimeslicetype.Marker) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalStructureType
func (verticalstructuretype *VerticalStructureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalstructuretype.TimeSlice) || modified
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
