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
// Clean garbage collect unstaged instances that are referenced by AngleIndicationTimeSlicePropertyType
func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&angleindicationtimeslicepropertytype.AngleIndicationTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AngleIndicationTimeSliceType
func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&angleindicationtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&angleindicationtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&angleindicationtimeslicetype.Fix) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_navaidSystem) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_position) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_runwayPoint) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_aimingPoint) || modified
	modified = stage.CleanPointer(&angleindicationtimeslicetype.PointChoice_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AngleIndicationType
func (angleindicationtype *AngleIndicationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&angleindicationtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForNavaidEquipmentPropertyType
func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForNavaidEquipmentType
func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityfornavaidequipmenttype.Annotation) || modified
	modified = stage.CleanSlice(&authorityfornavaidequipmenttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&authorityfornavaidequipmenttype.TheOrganisationAuthority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForSpecialNavigationStationPropertyType
func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForSpecialNavigationStationType
func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityforspecialnavigationstationtype.Annotation) || modified
	modified = stage.CleanSlice(&authorityforspecialnavigationstationtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&authorityforspecialnavigationstationtype.TheOrganisationAuthority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForSpecialNavigationSystemPropertyType
func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForSpecialNavigationSystemType
func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityforspecialnavigationsystemtype.Annotation) || modified
	modified = stage.CleanSlice(&authorityforspecialnavigationsystemtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&authorityforspecialnavigationsystemtype.TheOrganisationAuthority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AzimuthTimeSlicePropertyType
func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&azimuthtimeslicepropertytype.AzimuthTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AzimuthTimeSliceType
func (azimuthtimeslicetype *AzimuthTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&azimuthtimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&azimuthtimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&azimuthtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&azimuthtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&azimuthtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&azimuthtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AzimuthType
func (azimuthtype *AzimuthType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&azimuthtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointINSTimeSlicePropertyType
func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&checkpointinstimeslicepropertytype.CheckpointINSTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointINSTimeSliceType
func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&checkpointinstimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&checkpointinstimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&checkpointinstimeslicetype.Position) || modified
	modified = stage.CleanPointer(&checkpointinstimeslicetype.AirportHeliport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointINSType
func (checkpointinstype *CheckpointINSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&checkpointinstype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointVORTimeSlicePropertyType
func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&checkpointvortimeslicepropertytype.CheckpointVORTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointVORTimeSliceType
func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&checkpointvortimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&checkpointvortimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&checkpointvortimeslicetype.Position) || modified
	modified = stage.CleanPointer(&checkpointvortimeslicetype.AirportHeliport) || modified
	modified = stage.CleanPointer(&checkpointvortimeslicetype.CheckPointFacility) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckpointVORType
func (checkpointvortype *CheckpointVORType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&checkpointvortype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ContactInformationPropertyType
func (contactinformationpropertytype *ContactInformationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&contactinformationpropertytype.ContactInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ContactInformationType
func (contactinformationtype *ContactInformationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&contactinformationtype.Address) || modified
	modified = stage.CleanSlice(&contactinformationtype.NetworkNode) || modified
	modified = stage.CleanSlice(&contactinformationtype.PhoneFax) || modified
	modified = stage.CleanSlice(&contactinformationtype.Annotation) || modified
	modified = stage.CleanSlice(&contactinformationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DMETimeSlicePropertyType
func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dmetimeslicepropertytype.DMETimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DMETimeSliceType
func (dmetimeslicetype *DMETimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dmetimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&dmetimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&dmetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&dmetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&dmetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dmetimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DMEType
func (dmetype *DMEType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dmetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DesignatedPointTimeSlicePropertyType
func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&designatedpointtimeslicepropertytype.DesignatedPointTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DesignatedPointTimeSliceType
func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&designatedpointtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&designatedpointtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&designatedpointtimeslicetype.Location) || modified
	modified = stage.CleanPointer(&designatedpointtimeslicetype.AimingPoint) || modified
	modified = stage.CleanPointer(&designatedpointtimeslicetype.AirportHeliport) || modified
	modified = stage.CleanPointer(&designatedpointtimeslicetype.RunwayPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DesignatedPointType
func (designatedpointtype *DesignatedPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&designatedpointtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionFinderTimeSlicePropertyType
func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directionfindertimeslicepropertytype.DirectionFinderTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionFinderTimeSliceType
func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directionfindertimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&directionfindertimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&directionfindertimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&directionfindertimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&directionfindertimeslicetype.InformationProvision) || modified
	modified = stage.CleanSlice(&directionfindertimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&directionfindertimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionFinderType
func (directionfindertype *DirectionFinderType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directionfindertype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DistanceIndicationTimeSlicePropertyType
func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DistanceIndicationTimeSliceType
func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&distanceindicationtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&distanceindicationtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.Fix) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_navaidSystem) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_position) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_runwayPoint) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_aimingPoint) || modified
	modified = stage.CleanPointer(&distanceindicationtimeslicetype.PointChoice_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DistanceIndicationType
func (distanceindicationtype *DistanceIndicationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&distanceindicationtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by ElevationTimeSlicePropertyType
func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&elevationtimeslicepropertytype.ElevationTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevationTimeSliceType
func (elevationtimeslicetype *ElevationTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevationtimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&elevationtimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&elevationtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&elevationtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&elevationtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&elevationtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ElevationType
func (elevationtype *ElevationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevationtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GlidepathTimeSlicePropertyType
func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&glidepathtimeslicepropertytype.GlidepathTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GlidepathTimeSliceType
func (glidepathtimeslicetype *GlidepathTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&glidepathtimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&glidepathtimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&glidepathtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&glidepathtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&glidepathtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&glidepathtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GlidepathType
func (glidepathtype *GlidepathType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&glidepathtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by LocalizerTimeSlicePropertyType
func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&localizertimeslicepropertytype.LocalizerTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LocalizerTimeSliceType
func (localizertimeslicetype *LocalizerTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&localizertimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&localizertimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&localizertimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&localizertimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&localizertimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&localizertimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LocalizerType
func (localizertype *LocalizerType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&localizertype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkerBeaconTimeSlicePropertyType
func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkerBeaconTimeSliceType
func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&markerbeacontimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&markerbeacontimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&markerbeacontimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&markerbeacontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&markerbeacontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&markerbeacontimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkerBeaconType
func (markerbeacontype *MarkerBeaconType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&markerbeacontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NDBTimeSlicePropertyType
func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ndbtimeslicepropertytype.NDBTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NDBTimeSliceType
func (ndbtimeslicetype *NDBTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ndbtimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&ndbtimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&ndbtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&ndbtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&ndbtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ndbtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NDBType
func (ndbtype *NDBType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ndbtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidComponentPropertyType
func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navaidcomponentpropertytype.NavaidComponent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidComponentType
func (navaidcomponenttype *NavaidComponentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidcomponenttype.Annotation) || modified
	modified = stage.CleanSlice(&navaidcomponenttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navaidcomponenttype.TheNavaidEquipment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidEquipmentMonitoringPropertyType
func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidEquipmentMonitoringType
func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidequipmentmonitoringtype.TimeInterval) || modified
	modified = stage.CleanSlice(&navaidequipmentmonitoringtype.Annotation) || modified
	modified = stage.CleanSlice(&navaidequipmentmonitoringtype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&navaidequipmentmonitoringtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidOperationalStatusPropertyType
func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navaidoperationalstatuspropertytype.NavaidOperationalStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidOperationalStatusType
func (navaidoperationalstatustype *NavaidOperationalStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidoperationalstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&navaidoperationalstatustype.Annotation) || modified
	modified = stage.CleanSlice(&navaidoperationalstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&navaidoperationalstatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidTimeSlicePropertyType
func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navaidtimeslicepropertytype.NavaidTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidTimeSliceType
func (navaidtimeslicetype *NavaidTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidtimeslicetype.TouchDownLiftOff) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.NavaidEquipment) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.RunwayDirection) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.ServedAirport) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&navaidtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navaidtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidType
func (navaidtype *NavaidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by OnlineContactPropertyType
func (onlinecontactpropertytype *OnlineContactPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&onlinecontactpropertytype.OnlineContact) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OnlineContactType
func (onlinecontacttype *OnlineContactType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&onlinecontacttype.TimeInterval) || modified
	modified = stage.CleanSlice(&onlinecontacttype.Annotation) || modified
	modified = stage.CleanSlice(&onlinecontacttype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&onlinecontacttype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PointPropertyType
func (pointpropertytype *PointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pointpropertytype.Point) || modified
	modified = stage.CleanPointer(&pointpropertytype.ElevatedPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PointType
func (pointtype *PointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointtype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PostalAddressPropertyType
func (postaladdresspropertytype *PostalAddressPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&postaladdresspropertytype.PostalAddress) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PostalAddressType
func (postaladdresstype *PostalAddressType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&postaladdresstype.TimeInterval) || modified
	modified = stage.CleanSlice(&postaladdresstype.Annotation) || modified
	modified = stage.CleanSlice(&postaladdresstype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&postaladdresstype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PrecisionApproachRadarTimeSlicePropertyType
func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrecisionApproachRadarTimeSliceType
func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&precisionapproachradartimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&precisionapproachradartimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&precisionapproachradartimeslicetype.Reflector) || modified
	modified = stage.CleanSlice(&precisionapproachradartimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&precisionapproachradartimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrecisionApproachRadarType
func (precisionapproachradartype *PrecisionApproachRadarType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&precisionapproachradartype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PrimarySurveillanceRadarTimeSlicePropertyType
func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrimarySurveillanceRadarTimeSliceType
func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&primarysurveillanceradartimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&primarysurveillanceradartimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&primarysurveillanceradartimeslicetype.GroundStation) || modified
	modified = stage.CleanSlice(&primarysurveillanceradartimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&primarysurveillanceradartimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrimarySurveillanceRadarType
func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&primarysurveillanceradartype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RadarComponentPropertyType
func (radarcomponentpropertytype *RadarComponentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&radarcomponentpropertytype.RadarComponent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadarComponentType
func (radarcomponenttype *RadarComponentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radarcomponenttype.Annotation) || modified
	modified = stage.CleanSlice(&radarcomponenttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&radarcomponenttype.TheRadarEquipment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadarSystemTimeSlicePropertyType
func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&radarsystemtimeslicepropertytype.RadarSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadarSystemTimeSliceType
func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radarsystemtimeslicetype.RadarEquipment) || modified
	modified = stage.CleanSlice(&radarsystemtimeslicetype.Office) || modified
	modified = stage.CleanSlice(&radarsystemtimeslicetype.PARRunway) || modified
	modified = stage.CleanSlice(&radarsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&radarsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&radarsystemtimeslicetype.AirportHeliport) || modified
	modified = stage.CleanPointer(&radarsystemtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadarSystemType
func (radarsystemtype *RadarSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radarsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ReflectorPropertyType
func (reflectorpropertytype *ReflectorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&reflectorpropertytype.Reflector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ReflectorType
func (reflectortype *ReflectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&reflectortype.Annotation) || modified
	modified = stage.CleanSlice(&reflectortype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&reflectortype.TouchdownReflector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SDFTimeSlicePropertyType
func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&sdftimeslicepropertytype.SDFTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SDFTimeSliceType
func (sdftimeslicetype *SDFTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sdftimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&sdftimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&sdftimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&sdftimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&sdftimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&sdftimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SDFType
func (sdftype *SDFType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sdftype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SecondarySurveillanceRadarTimeSlicePropertyType
func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SecondarySurveillanceRadarTimeSliceType
func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&secondarysurveillanceradartimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&secondarysurveillanceradartimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&secondarysurveillanceradartimeslicetype.GroundStation) || modified
	modified = stage.CleanSlice(&secondarysurveillanceradartimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&secondarysurveillanceradartimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SecondarySurveillanceRadarType
func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&secondarysurveillanceradartype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SignificantPointInAirspaceTimeSlicePropertyType
func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SignificantPointInAirspaceTimeSliceType
func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&significantpointinairspacetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&significantpointinairspacetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.ContainingAirspace) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_navaidSystem) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_position) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_runwayPoint) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_aimingPoint) || modified
	modified = stage.CleanPointer(&significantpointinairspacetimeslicetype.Location_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SignificantPointInAirspaceType
func (significantpointinairspacetype *SignificantPointInAirspaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&significantpointinairspacetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationStationStatusPropertyType
func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationStationStatusType
func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialnavigationstationstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&specialnavigationstationstatustype.Annotation) || modified
	modified = stage.CleanSlice(&specialnavigationstationstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&specialnavigationstationstatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationStationTimeSlicePropertyType
func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationStationTimeSliceType
func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialnavigationstationtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&specialnavigationstationtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&specialnavigationstationtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&specialnavigationstationtimeslicetype.SystemChain) || modified
	modified = stage.CleanPointer(&specialnavigationstationtimeslicetype.ResponsibleOrganisation) || modified
	modified = stage.CleanPointer(&specialnavigationstationtimeslicetype.Position) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationStationType
func (specialnavigationstationtype *SpecialNavigationStationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialnavigationstationtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationSystemTimeSlicePropertyType
func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationSystemTimeSliceType
func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialnavigationsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&specialnavigationsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&specialnavigationsystemtimeslicetype.ResponsibleOrganisation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialNavigationSystemType
func (specialnavigationsystemtype *SpecialNavigationSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialnavigationsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurveillanceGroundStationPropertyType
func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surveillancegroundstationpropertytype.SurveillanceGroundStation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurveillanceGroundStationType
func (surveillancegroundstationtype *SurveillanceGroundStationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surveillancegroundstationtype.Annotation) || modified
	modified = stage.CleanSlice(&surveillancegroundstationtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&surveillancegroundstationtype.TheUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TACANTimeSlicePropertyType
func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tacantimeslicepropertytype.TACANTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TACANTimeSliceType
func (tacantimeslicetype *TACANTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tacantimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&tacantimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&tacantimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&tacantimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&tacantimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&tacantimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TACANType
func (tacantype *TACANType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tacantype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TelephoneContactPropertyType
func (telephonecontactpropertytype *TelephoneContactPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&telephonecontactpropertytype.TelephoneContact) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TelephoneContactType
func (telephonecontacttype *TelephoneContactType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&telephonecontacttype.TimeInterval) || modified
	modified = stage.CleanSlice(&telephonecontacttype.Annotation) || modified
	modified = stage.CleanSlice(&telephonecontacttype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&telephonecontacttype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by VORTimeSlicePropertyType
func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&vortimeslicepropertytype.VORTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VORTimeSliceType
func (vortimeslicetype *VORTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&vortimeslicetype.Authority) || modified
	modified = stage.CleanSlice(&vortimeslicetype.Monitoring) || modified
	modified = stage.CleanSlice(&vortimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&vortimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&vortimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&vortimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VORType
func (vortype *VORType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&vortype.TimeSlice) || modified
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
