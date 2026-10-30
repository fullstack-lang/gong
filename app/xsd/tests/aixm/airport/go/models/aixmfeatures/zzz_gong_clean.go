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

// Clean garbage collect unstaged instances that are referenced by AircraftStandContaminationPropertyType
func (aircraftstandcontaminationpropertytype *AircraftStandContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aircraftstandcontaminationpropertytype.AircraftStandContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftStandContaminationType
func (aircraftstandcontaminationtype *AircraftStandContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftstandcontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&aircraftstandcontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&aircraftstandcontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&aircraftstandcontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftStandTimeSlicePropertyType
func (aircraftstandtimeslicepropertytype *AircraftStandTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aircraftstandtimeslicepropertytype.AircraftStandTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftStandTimeSliceType
func (aircraftstandtimeslicetype *AircraftStandTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftstandtimeslicetype.Contaminant) || modified
	modified = stage.CleanSlice(&aircraftstandtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&aircraftstandtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&aircraftstandtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aircraftstandtimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&aircraftstandtimeslicetype.Location) || modified
	modified = stage.CleanPointer(&aircraftstandtimeslicetype.ApronLocation) || modified
	modified = stage.CleanPointer(&aircraftstandtimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftStandType
func (aircraftstandtype *AircraftStandType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftstandtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportAvailabilityPropertyType
func (airportheliportavailabilitypropertytype *AirportHeliportAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportavailabilitypropertytype.AirportHeliportAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportAvailabilityType
func (airportheliportavailabilitytype *AirportHeliportAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&airportheliportavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliportavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&airportheliportavailabilitytype.Usage) || modified
	modified = stage.CleanSlice(&airportheliportavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportCollocationTimeSlicePropertyType
func (airportheliportcollocationtimeslicepropertytype *AirportHeliportCollocationTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportcollocationtimeslicepropertytype.AirportHeliportCollocationTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportCollocationTimeSliceType
func (airportheliportcollocationtimeslicetype *AirportHeliportCollocationTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportcollocationtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliportcollocationtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportcollocationtimeslicetype.HostAirport) || modified
	modified = stage.CleanPointer(&airportheliportcollocationtimeslicetype.DependentAirport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportCollocationType
func (airportheliportcollocationtype *AirportHeliportCollocationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportcollocationtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportContaminationPropertyType
func (airportheliportcontaminationpropertytype *AirportHeliportContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportcontaminationpropertytype.AirportHeliportContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportContaminationType
func (airportheliportcontaminationtype *AirportHeliportContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportcontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&airportheliportcontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&airportheliportcontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliportcontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportResponsibilityOrganisationPropertyType
func (airportheliportresponsibilityorganisationpropertytype *AirportHeliportResponsibilityOrganisationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportresponsibilityorganisationpropertytype.AirportHeliportResponsibilityOrganisation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportResponsibilityOrganisationType
func (airportheliportresponsibilityorganisationtype *AirportHeliportResponsibilityOrganisationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportresponsibilityorganisationtype.TimeInterval) || modified
	modified = stage.CleanSlice(&airportheliportresponsibilityorganisationtype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliportresponsibilityorganisationtype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&airportheliportresponsibilityorganisationtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportresponsibilityorganisationtype.TheOrganisationAuthority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportTimeSlicePropertyType
func (airportheliporttimeslicepropertytype *AirportHeliportTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliporttimeslicepropertytype.AirportHeliportTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportTimeSliceType
func (airportheliporttimeslicetype *AirportHeliportTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliporttimeslicetype.Contaminant) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.ServedCity) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.AltimeterSource) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliporttimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportheliporttimeslicetype.ResponsibleOrganisation) || modified
	modified = stage.CleanPointer(&airportheliporttimeslicetype.ARP) || modified
	modified = stage.CleanPointer(&airportheliporttimeslicetype.AviationBoundary) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportType
func (airportheliporttype *AirportHeliportType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliporttype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportUsagePropertyType
func (airportheliportusagepropertytype *AirportHeliportUsagePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportusagepropertytype.AirportHeliportUsage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHeliportUsageType
func (airportheliportusagetype *AirportHeliportUsageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportheliportusagetype.Contact) || modified
	modified = stage.CleanSlice(&airportheliportusagetype.Annotation) || modified
	modified = stage.CleanSlice(&airportheliportusagetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportheliportusagetype.Selection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHotSpotTimeSlicePropertyType
func (airporthotspottimeslicepropertytype *AirportHotSpotTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airporthotspottimeslicepropertytype.AirportHotSpotTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHotSpotTimeSliceType
func (airporthotspottimeslicetype *AirportHotSpotTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airporthotspottimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airporthotspottimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airporthotspottimeslicetype.Area) || modified
	modified = stage.CleanPointer(&airporthotspottimeslicetype.AffectedAirport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportHotSpotType
func (airporthotspottype *AirportHotSpotType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airporthotspottype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportProtectionAreaMarkingTimeSlicePropertyType
func (airportprotectionareamarkingtimeslicepropertytype *AirportProtectionAreaMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportprotectionareamarkingtimeslicepropertytype.AirportProtectionAreaMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportProtectionAreaMarkingTimeSliceType
func (airportprotectionareamarkingtimeslicetype *AirportProtectionAreaMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportprotectionareamarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&airportprotectionareamarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airportprotectionareamarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportprotectionareamarkingtimeslicetype.MarkedProtectionArea) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportProtectionAreaMarkingType
func (airportprotectionareamarkingtype *AirportProtectionAreaMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportprotectionareamarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AltimeterSourceStatusPropertyType
func (altimetersourcestatuspropertytype *AltimeterSourceStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&altimetersourcestatuspropertytype.AltimeterSourceStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AltimeterSourceStatusType
func (altimetersourcestatustype *AltimeterSourceStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&altimetersourcestatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&altimetersourcestatustype.Annotation) || modified
	modified = stage.CleanSlice(&altimetersourcestatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&altimetersourcestatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AltimeterSourceTimeSlicePropertyType
func (altimetersourcetimeslicepropertytype *AltimeterSourceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&altimetersourcetimeslicepropertytype.AltimeterSourceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AltimeterSourceTimeSliceType
func (altimetersourcetimeslicetype *AltimeterSourceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&altimetersourcetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&altimetersourcetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&altimetersourcetimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AltimeterSourceType
func (altimetersourcetype *AltimeterSourceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&altimetersourcetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachLightingSystemTimeSlicePropertyType
func (approachlightingsystemtimeslicepropertytype *ApproachLightingSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&approachlightingsystemtimeslicepropertytype.ApproachLightingSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachLightingSystemTimeSliceType
func (approachlightingsystemtimeslicetype *ApproachLightingSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachlightingsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&approachlightingsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&approachlightingsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&approachlightingsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&approachlightingsystemtimeslicetype.ServedRunwayDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachLightingSystemType
func (approachlightingsystemtype *ApproachLightingSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachlightingsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronAreaAvailabilityPropertyType
func (apronareaavailabilitypropertytype *ApronAreaAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&apronareaavailabilitypropertytype.ApronAreaAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronAreaAvailabilityType
func (apronareaavailabilitytype *ApronAreaAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronareaavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&apronareaavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&apronareaavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&apronareaavailabilitytype.Usage) || modified
	modified = stage.CleanSlice(&apronareaavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronAreaUsagePropertyType
func (apronareausagepropertytype *ApronAreaUsagePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&apronareausagepropertytype.ApronAreaUsage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronAreaUsageType
func (apronareausagetype *ApronAreaUsageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronareausagetype.Contact) || modified
	modified = stage.CleanSlice(&apronareausagetype.Annotation) || modified
	modified = stage.CleanSlice(&apronareausagetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&apronareausagetype.Selection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronContaminationPropertyType
func (aproncontaminationpropertytype *ApronContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aproncontaminationpropertytype.ApronContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronContaminationType
func (aproncontaminationtype *ApronContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aproncontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&aproncontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&aproncontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&aproncontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronElementTimeSlicePropertyType
func (apronelementtimeslicepropertytype *ApronElementTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&apronelementtimeslicepropertytype.ApronElementTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronElementTimeSliceType
func (apronelementtimeslicetype *ApronElementTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronelementtimeslicetype.SupplyService) || modified
	modified = stage.CleanSlice(&apronelementtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&apronelementtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&apronelementtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&apronelementtimeslicetype.AssociatedApron) || modified
	modified = stage.CleanPointer(&apronelementtimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&apronelementtimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronElementType
func (apronelementtype *ApronElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronelementtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronLightSystemTimeSlicePropertyType
func (apronlightsystemtimeslicepropertytype *ApronLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&apronlightsystemtimeslicepropertytype.ApronLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronLightSystemTimeSliceType
func (apronlightsystemtimeslicetype *ApronLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronlightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&apronlightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&apronlightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&apronlightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&apronlightsystemtimeslicetype.LightedApron) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronLightSystemType
func (apronlightsystemtype *ApronLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronlightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronMarkingTimeSlicePropertyType
func (apronmarkingtimeslicepropertytype *ApronMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&apronmarkingtimeslicepropertytype.ApronMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronMarkingTimeSliceType
func (apronmarkingtimeslicetype *ApronMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronmarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&apronmarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&apronmarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&apronmarkingtimeslicetype.MarkedApron) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronMarkingType
func (apronmarkingtype *ApronMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&apronmarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronTimeSlicePropertyType
func (aprontimeslicepropertytype *ApronTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aprontimeslicepropertytype.ApronTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronTimeSliceType
func (aprontimeslicetype *ApronTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aprontimeslicetype.Contaminant) || modified
	modified = stage.CleanSlice(&aprontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&aprontimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&aprontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aprontimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&aprontimeslicetype.AssociatedAirportHeliport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApronType
func (aprontype *ApronType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aprontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrestingGearTimeSlicePropertyType
func (arrestinggeartimeslicepropertytype *ArrestingGearTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arrestinggeartimeslicepropertytype.ArrestingGearTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrestingGearTimeSliceType
func (arrestinggeartimeslicetype *ArrestingGearTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrestinggeartimeslicetype.RunwayDirection) || modified
	modified = stage.CleanSlice(&arrestinggeartimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&arrestinggeartimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arrestinggeartimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&arrestinggeartimeslicetype.Extent_surfaceExtent) || modified
	modified = stage.CleanPointer(&arrestinggeartimeslicetype.Extent_curveExtent) || modified
	modified = stage.CleanPointer(&arrestinggeartimeslicetype.Extent_pointExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrestingGearType
func (arrestinggeartype *ArrestingGearType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrestinggeartype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CityPropertyType
func (citypropertytype *CityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&citypropertytype.City) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CityType
func (citytype *CityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&citytype.Annotation) || modified
	modified = stage.CleanSlice(&citytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ConditionCombinationPropertyType
func (conditioncombinationpropertytype *ConditionCombinationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&conditioncombinationpropertytype.ConditionCombination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConditionCombinationType
func (conditioncombinationtype *ConditionCombinationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&conditioncombinationtype.TimeInterval) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.Annotation) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.Weather) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.Aircraft) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.Flight) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.SubCondition) || modified
	modified = stage.CleanSlice(&conditioncombinationtype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by CurveType
func (curvetype *CurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvetype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaMarkingTimeSlicePropertyType
func (deicingareamarkingtimeslicepropertytype *DeicingAreaMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&deicingareamarkingtimeslicepropertytype.DeicingAreaMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaMarkingTimeSliceType
func (deicingareamarkingtimeslicetype *DeicingAreaMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deicingareamarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&deicingareamarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&deicingareamarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&deicingareamarkingtimeslicetype.MarkedDeicingArea) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaMarkingType
func (deicingareamarkingtype *DeicingAreaMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deicingareamarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaTimeSlicePropertyType
func (deicingareatimeslicepropertytype *DeicingAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&deicingareatimeslicepropertytype.DeicingAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaTimeSliceType
func (deicingareatimeslicetype *DeicingAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deicingareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&deicingareatimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&deicingareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&deicingareatimeslicetype.AssociatedApron) || modified
	modified = stage.CleanPointer(&deicingareatimeslicetype.TaxiwayLocation) || modified
	modified = stage.CleanPointer(&deicingareatimeslicetype.StandLocation) || modified
	modified = stage.CleanPointer(&deicingareatimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&deicingareatimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DeicingAreaType
func (deicingareatype *DeicingAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&deicingareatype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by FlightCharacteristicPropertyType
func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightcharacteristicpropertytype.FlightCharacteristic) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightCharacteristicType
func (flightcharacteristictype *FlightCharacteristicType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightcharacteristictype.Annotation) || modified
	modified = stage.CleanSlice(&flightcharacteristictype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FloatingDockSiteTimeSlicePropertyType
func (floatingdocksitetimeslicepropertytype *FloatingDockSiteTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&floatingdocksitetimeslicepropertytype.FloatingDockSiteTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FloatingDockSiteTimeSliceType
func (floatingdocksitetimeslicetype *FloatingDockSiteTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&floatingdocksitetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&floatingdocksitetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&floatingdocksitetimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FloatingDockSiteType
func (floatingdocksitetype *FloatingDockSiteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&floatingdocksitetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GroundLightingAvailabilityPropertyType
func (groundlightingavailabilitypropertytype *GroundLightingAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&groundlightingavailabilitypropertytype.GroundLightingAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GroundLightingAvailabilityType
func (groundlightingavailabilitytype *GroundLightingAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&groundlightingavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&groundlightingavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&groundlightingavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&groundlightingavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineLightSystemTimeSlicePropertyType
func (guidancelinelightsystemtimeslicepropertytype *GuidanceLineLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinelightsystemtimeslicepropertytype.GuidanceLineLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineLightSystemTimeSliceType
func (guidancelinelightsystemtimeslicetype *GuidanceLineLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinelightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&guidancelinelightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&guidancelinelightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&guidancelinelightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinelightsystemtimeslicetype.LightedGuidanceLine) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineLightSystemType
func (guidancelinelightsystemtype *GuidanceLineLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinelightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineMarkingTimeSlicePropertyType
func (guidancelinemarkingtimeslicepropertytype *GuidanceLineMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinemarkingtimeslicepropertytype.GuidanceLineMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineMarkingTimeSliceType
func (guidancelinemarkingtimeslicetype *GuidanceLineMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinemarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&guidancelinemarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&guidancelinemarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinemarkingtimeslicetype.MarkedGuidanceLine) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineMarkingType
func (guidancelinemarkingtype *GuidanceLineMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinemarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineTimeSlicePropertyType
func (guidancelinetimeslicepropertytype *GuidanceLineTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinetimeslicepropertytype.GuidanceLineTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineTimeSliceType
func (guidancelinetimeslicetype *GuidanceLineTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinetimeslicetype.ConnectedTouchDownLiftOff) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.ConnectedRunwayCentrelinePoint) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.ConnectedApron) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.ConnectedStand) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.ConnectedTaxiway) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&guidancelinetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&guidancelinetimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GuidanceLineType
func (guidancelinetype *GuidanceLineType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&guidancelinetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by ManoeuvringAreaAvailabilityPropertyType
func (manoeuvringareaavailabilitypropertytype *ManoeuvringAreaAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&manoeuvringareaavailabilitypropertytype.ManoeuvringAreaAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ManoeuvringAreaAvailabilityType
func (manoeuvringareaavailabilitytype *ManoeuvringAreaAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&manoeuvringareaavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&manoeuvringareaavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&manoeuvringareaavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&manoeuvringareaavailabilitytype.Usage) || modified
	modified = stage.CleanSlice(&manoeuvringareaavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ManoeuvringAreaUsagePropertyType
func (manoeuvringareausagepropertytype *ManoeuvringAreaUsagePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&manoeuvringareausagepropertytype.ManoeuvringAreaUsage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ManoeuvringAreaUsageType
func (manoeuvringareausagetype *ManoeuvringAreaUsageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&manoeuvringareausagetype.Contact) || modified
	modified = stage.CleanSlice(&manoeuvringareausagetype.Annotation) || modified
	modified = stage.CleanSlice(&manoeuvringareausagetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&manoeuvringareausagetype.Selection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkingBuoyTimeSlicePropertyType
func (markingbuoytimeslicepropertytype *MarkingBuoyTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&markingbuoytimeslicepropertytype.MarkingBuoyTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkingBuoyTimeSliceType
func (markingbuoytimeslicetype *MarkingBuoyTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&markingbuoytimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&markingbuoytimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&markingbuoytimeslicetype.TheSeaplaneLandingArea) || modified
	modified = stage.CleanPointer(&markingbuoytimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkingBuoyType
func (markingbuoytype *MarkingBuoyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&markingbuoytype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkingElementPropertyType
func (markingelementpropertytype *MarkingElementPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&markingelementpropertytype.MarkingElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MarkingElementType
func (markingelementtype *MarkingElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&markingelementtype.Annotation) || modified
	modified = stage.CleanSlice(&markingelementtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&markingelementtype.Extent_curveExtent) || modified
	modified = stage.CleanPointer(&markingelementtype.Extent_surfaceExtent) || modified
	modified = stage.CleanPointer(&markingelementtype.Extent_location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MeteorologyPropertyType
func (meteorologypropertytype *MeteorologyPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&meteorologypropertytype.Meteorology) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MeteorologyType
func (meteorologytype *MeteorologyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&meteorologytype.Annotation) || modified
	modified = stage.CleanSlice(&meteorologytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidEquipmentDistancePropertyType
func (navaidequipmentdistancepropertytype *NavaidEquipmentDistancePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navaidequipmentdistancepropertytype.NavaidEquipmentDistance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavaidEquipmentDistanceType
func (navaidequipmentdistancetype *NavaidEquipmentDistanceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navaidequipmentdistancetype.Annotation) || modified
	modified = stage.CleanSlice(&navaidequipmentdistancetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navaidequipmentdistancetype.TheNavaidEquipment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NonMovementAreaTimeSlicePropertyType
func (nonmovementareatimeslicepropertytype *NonMovementAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&nonmovementareatimeslicepropertytype.NonMovementAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NonMovementAreaTimeSliceType
func (nonmovementareatimeslicetype *NonMovementAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&nonmovementareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&nonmovementareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&nonmovementareatimeslicetype.AssociatedAirportHeliport) || modified
	modified = stage.CleanPointer(&nonmovementareatimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NonMovementAreaType
func (nonmovementareatype *NonMovementAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&nonmovementareatype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by PassengerLoadingBridgeTimeSlicePropertyType
func (passengerloadingbridgetimeslicepropertytype *PassengerLoadingBridgeTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&passengerloadingbridgetimeslicepropertytype.PassengerLoadingBridgeTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassengerLoadingBridgeTimeSliceType
func (passengerloadingbridgetimeslicetype *PassengerLoadingBridgeTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&passengerloadingbridgetimeslicetype.AssociatedStand) || modified
	modified = stage.CleanSlice(&passengerloadingbridgetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&passengerloadingbridgetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&passengerloadingbridgetimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassengerLoadingBridgeType
func (passengerloadingbridgetype *PassengerLoadingBridgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&passengerloadingbridgetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by RidgePropertyType
func (ridgepropertytype *RidgePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ridgepropertytype.Ridge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RidgeType
func (ridgetype *RidgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ridgetype.Annotation) || modified
	modified = stage.CleanSlice(&ridgetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RoadTimeSlicePropertyType
func (roadtimeslicepropertytype *RoadTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&roadtimeslicepropertytype.RoadTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RoadTimeSliceType
func (roadtimeslicetype *RoadTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&roadtimeslicetype.AccessibleStand) || modified
	modified = stage.CleanSlice(&roadtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&roadtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&roadtimeslicetype.AssociatedAirport) || modified
	modified = stage.CleanPointer(&roadtimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&roadtimeslicetype.SurfaceExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RoadType
func (roadtype *RoadType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&roadtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayBlastPadTimeSlicePropertyType
func (runwayblastpadtimeslicepropertytype *RunwayBlastPadTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwayblastpadtimeslicepropertytype.RunwayBlastPadTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayBlastPadTimeSliceType
func (runwayblastpadtimeslicetype *RunwayBlastPadTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayblastpadtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwayblastpadtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwayblastpadtimeslicetype.UsedRunwayDirection) || modified
	modified = stage.CleanPointer(&runwayblastpadtimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&runwayblastpadtimeslicetype.SurfaceProperties) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayBlastPadType
func (runwayblastpadtype *RunwayBlastPadType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayblastpadtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayCentrelinePointTimeSlicePropertyType
func (runwaycentrelinepointtimeslicepropertytype *RunwayCentrelinePointTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaycentrelinepointtimeslicepropertytype.RunwayCentrelinePointTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayCentrelinePointTimeSliceType
func (runwaycentrelinepointtimeslicetype *RunwayCentrelinePointTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaycentrelinepointtimeslicetype.AssociatedDeclaredDistance) || modified
	modified = stage.CleanSlice(&runwaycentrelinepointtimeslicetype.NavaidEquipment) || modified
	modified = stage.CleanSlice(&runwaycentrelinepointtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaycentrelinepointtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwaycentrelinepointtimeslicetype.Location) || modified
	modified = stage.CleanPointer(&runwaycentrelinepointtimeslicetype.OnRunway) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayCentrelinePointType
func (runwaycentrelinepointtype *RunwayCentrelinePointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaycentrelinepointtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayContaminationPropertyType
func (runwaycontaminationpropertytype *RunwayContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaycontaminationpropertytype.RunwayContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayContaminationType
func (runwaycontaminationtype *RunwayContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaycontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&runwaycontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&runwaycontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&runwaycontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDeclaredDistancePropertyType
func (runwaydeclareddistancepropertytype *RunwayDeclaredDistancePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaydeclareddistancepropertytype.RunwayDeclaredDistance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDeclaredDistanceType
func (runwaydeclareddistancetype *RunwayDeclaredDistanceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydeclareddistancetype.DeclaredValue) || modified
	modified = stage.CleanSlice(&runwaydeclareddistancetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaydeclareddistancetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDeclaredDistanceValuePropertyType
func (runwaydeclareddistancevaluepropertytype *RunwayDeclaredDistanceValuePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaydeclareddistancevaluepropertytype.RunwayDeclaredDistanceValue) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDeclaredDistanceValueType
func (runwaydeclareddistancevaluetype *RunwayDeclaredDistanceValueType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydeclareddistancevaluetype.TimeInterval) || modified
	modified = stage.CleanSlice(&runwaydeclareddistancevaluetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaydeclareddistancevaluetype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&runwaydeclareddistancevaluetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionLightSystemTimeSlicePropertyType
func (runwaydirectionlightsystemtimeslicepropertytype *RunwayDirectionLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaydirectionlightsystemtimeslicepropertytype.RunwayDirectionLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionLightSystemTimeSliceType
func (runwaydirectionlightsystemtimeslicetype *RunwayDirectionLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydirectionlightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&runwaydirectionlightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&runwaydirectionlightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaydirectionlightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwaydirectionlightsystemtimeslicetype.AssociatedRunwayDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionLightSystemType
func (runwaydirectionlightsystemtype *RunwayDirectionLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydirectionlightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionTimeSlicePropertyType
func (runwaydirectiontimeslicepropertytype *RunwayDirectionTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaydirectiontimeslicepropertytype.RunwayDirectionTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionTimeSliceType
func (runwaydirectiontimeslicetype *RunwayDirectionTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydirectiontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaydirectiontimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&runwaydirectiontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwaydirectiontimeslicetype.UsedRunway) || modified
	modified = stage.CleanPointer(&runwaydirectiontimeslicetype.StartingElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayDirectionType
func (runwaydirectiontype *RunwayDirectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaydirectiontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayElementTimeSlicePropertyType
func (runwayelementtimeslicepropertytype *RunwayElementTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwayelementtimeslicepropertytype.RunwayElementTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayElementTimeSliceType
func (runwayelementtimeslicetype *RunwayElementTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayelementtimeslicetype.AssociatedRunway) || modified
	modified = stage.CleanSlice(&runwayelementtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwayelementtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&runwayelementtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwayelementtimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&runwayelementtimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayElementType
func (runwayelementtype *RunwayElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayelementtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayMarkingTimeSlicePropertyType
func (runwaymarkingtimeslicepropertytype *RunwayMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaymarkingtimeslicepropertytype.RunwayMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayMarkingTimeSliceType
func (runwaymarkingtimeslicetype *RunwayMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaymarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&runwaymarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaymarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwaymarkingtimeslicetype.MarkedRunway) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayMarkingType
func (runwaymarkingtype *RunwayMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaymarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaLightSystemTimeSlicePropertyType
func (runwayprotectarealightsystemtimeslicepropertytype *RunwayProtectAreaLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwayprotectarealightsystemtimeslicepropertytype.RunwayProtectAreaLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaLightSystemTimeSliceType
func (runwayprotectarealightsystemtimeslicetype *RunwayProtectAreaLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayprotectarealightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&runwayprotectarealightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&runwayprotectarealightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwayprotectarealightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwayprotectarealightsystemtimeslicetype.LightedArea) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaLightSystemType
func (runwayprotectarealightsystemtype *RunwayProtectAreaLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayprotectarealightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaTimeSlicePropertyType
func (runwayprotectareatimeslicepropertytype *RunwayProtectAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwayprotectareatimeslicepropertytype.RunwayProtectAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaTimeSliceType
func (runwayprotectareatimeslicetype *RunwayProtectAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayprotectareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwayprotectareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwayprotectareatimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&runwayprotectareatimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&runwayprotectareatimeslicetype.ProtectedRunwayDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayProtectAreaType
func (runwayprotectareatype *RunwayProtectAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayprotectareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwaySectionContaminationPropertyType
func (runwaysectioncontaminationpropertytype *RunwaySectionContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaysectioncontaminationpropertytype.RunwaySectionContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwaySectionContaminationType
func (runwaysectioncontaminationtype *RunwaySectionContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaysectioncontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&runwaysectioncontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&runwaysectioncontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&runwaysectioncontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayTimeSlicePropertyType
func (runwaytimeslicepropertytype *RunwayTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwaytimeslicepropertytype.RunwayTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayTimeSliceType
func (runwaytimeslicetype *RunwayTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaytimeslicetype.OverallContaminant) || modified
	modified = stage.CleanSlice(&runwaytimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwaytimeslicetype.AreaContaminant) || modified
	modified = stage.CleanSlice(&runwaytimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwaytimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&runwaytimeslicetype.AssociatedAirportHeliport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayType
func (runwaytype *RunwayType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwaytype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayVisualRangeTimeSlicePropertyType
func (runwayvisualrangetimeslicepropertytype *RunwayVisualRangeTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runwayvisualrangetimeslicepropertytype.RunwayVisualRangeTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayVisualRangeTimeSliceType
func (runwayvisualrangetimeslicetype *RunwayVisualRangeTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayvisualrangetimeslicetype.AssociatedRunwayDirection) || modified
	modified = stage.CleanSlice(&runwayvisualrangetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&runwayvisualrangetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&runwayvisualrangetimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RunwayVisualRangeType
func (runwayvisualrangetype *RunwayVisualRangeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&runwayvisualrangetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneLandingAreaTimeSlicePropertyType
func (seaplanelandingareatimeslicepropertytype *SeaplaneLandingAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&seaplanelandingareatimeslicepropertytype.SeaplaneLandingAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneLandingAreaTimeSliceType
func (seaplanelandingareatimeslicetype *SeaplaneLandingAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&seaplanelandingareatimeslicetype.RampSite) || modified
	modified = stage.CleanSlice(&seaplanelandingareatimeslicetype.DockSite) || modified
	modified = stage.CleanSlice(&seaplanelandingareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&seaplanelandingareatimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&seaplanelandingareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&seaplanelandingareatimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneLandingAreaType
func (seaplanelandingareatype *SeaplaneLandingAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&seaplanelandingareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneRampSiteTimeSlicePropertyType
func (seaplanerampsitetimeslicepropertytype *SeaplaneRampSiteTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&seaplanerampsitetimeslicepropertytype.SeaplaneRampSiteTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneRampSiteTimeSliceType
func (seaplanerampsitetimeslicetype *SeaplaneRampSiteTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&seaplanerampsitetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&seaplanerampsitetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&seaplanerampsitetimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&seaplanerampsitetimeslicetype.Centreline) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SeaplaneRampSiteType
func (seaplanerampsitetype *SeaplaneRampSiteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&seaplanerampsitetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandMarkingTimeSlicePropertyType
func (standmarkingtimeslicepropertytype *StandMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standmarkingtimeslicepropertytype.StandMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandMarkingTimeSliceType
func (standmarkingtimeslicetype *StandMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standmarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&standmarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standmarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&standmarkingtimeslicetype.MarkedStand) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandMarkingType
func (standmarkingtype *StandMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standmarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceCharacteristicsPropertyType
func (surfacecharacteristicspropertytype *SurfaceCharacteristicsPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surfacecharacteristicspropertytype.SurfaceCharacteristics) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceCharacteristicsType
func (surfacecharacteristicstype *SurfaceCharacteristicsType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacecharacteristicstype.Annotation) || modified
	modified = stage.CleanSlice(&surfacecharacteristicstype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceContaminationLayerPropertyType
func (surfacecontaminationlayerpropertytype *SurfaceContaminationLayerPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surfacecontaminationlayerpropertytype.SurfaceContaminationLayer) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceContaminationLayerType
func (surfacecontaminationlayertype *SurfaceContaminationLayerType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacecontaminationlayertype.Extent) || modified
	modified = stage.CleanSlice(&surfacecontaminationlayertype.Annotation) || modified
	modified = stage.CleanSlice(&surfacecontaminationlayertype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceType
func (surfacetype *SurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacetype.Annotation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurveyControlPointTimeSlicePropertyType
func (surveycontrolpointtimeslicepropertytype *SurveyControlPointTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surveycontrolpointtimeslicepropertytype.SurveyControlPointTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurveyControlPointTimeSliceType
func (surveycontrolpointtimeslicetype *SurveyControlPointTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surveycontrolpointtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&surveycontrolpointtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&surveycontrolpointtimeslicetype.AssociatedAirportHeliport) || modified
	modified = stage.CleanPointer(&surveycontrolpointtimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurveyControlPointType
func (surveycontrolpointtype *SurveyControlPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surveycontrolpointtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionLightSystemTimeSlicePropertyType
func (taxiholdingpositionlightsystemtimeslicepropertytype *TaxiHoldingPositionLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositionlightsystemtimeslicepropertytype.TaxiHoldingPositionLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionLightSystemTimeSliceType
func (taxiholdingpositionlightsystemtimeslicetype *TaxiHoldingPositionLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositionlightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&taxiholdingpositionlightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&taxiholdingpositionlightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiholdingpositionlightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositionlightsystemtimeslicetype.TaxiHolding) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionLightSystemType
func (taxiholdingpositionlightsystemtype *TaxiHoldingPositionLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositionlightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionMarkingTimeSlicePropertyType
func (taxiholdingpositionmarkingtimeslicepropertytype *TaxiHoldingPositionMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositionmarkingtimeslicepropertytype.TaxiHoldingPositionMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionMarkingTimeSliceType
func (taxiholdingpositionmarkingtimeslicetype *TaxiHoldingPositionMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositionmarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&taxiholdingpositionmarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiholdingpositionmarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositionmarkingtimeslicetype.MarkedTaxiHold) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionMarkingType
func (taxiholdingpositionmarkingtype *TaxiHoldingPositionMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositionmarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionTimeSlicePropertyType
func (taxiholdingpositiontimeslicepropertytype *TaxiHoldingPositionTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositiontimeslicepropertytype.TaxiHoldingPositionTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionTimeSliceType
func (taxiholdingpositiontimeslicetype *TaxiHoldingPositionTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositiontimeslicetype.ProtectedRunway) || modified
	modified = stage.CleanSlice(&taxiholdingpositiontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiholdingpositiontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiholdingpositiontimeslicetype.AssociatedGuidanceLine) || modified
	modified = stage.CleanPointer(&taxiholdingpositiontimeslicetype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiHoldingPositionType
func (taxiholdingpositiontype *TaxiHoldingPositionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiholdingpositiontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayContaminationPropertyType
func (taxiwaycontaminationpropertytype *TaxiwayContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaycontaminationpropertytype.TaxiwayContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayContaminationType
func (taxiwaycontaminationtype *TaxiwayContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaycontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&taxiwaycontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&taxiwaycontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&taxiwaycontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayElementTimeSlicePropertyType
func (taxiwayelementtimeslicepropertytype *TaxiwayElementTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiwayelementtimeslicepropertytype.TaxiwayElementTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayElementTimeSliceType
func (taxiwayelementtimeslicetype *TaxiwayElementTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwayelementtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiwayelementtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&taxiwayelementtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiwayelementtimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&taxiwayelementtimeslicetype.AssociatedTaxiway) || modified
	modified = stage.CleanPointer(&taxiwayelementtimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayElementType
func (taxiwayelementtype *TaxiwayElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwayelementtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayLightSystemTimeSlicePropertyType
func (taxiwaylightsystemtimeslicepropertytype *TaxiwayLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaylightsystemtimeslicepropertytype.TaxiwayLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayLightSystemTimeSliceType
func (taxiwaylightsystemtimeslicetype *TaxiwayLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaylightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&taxiwaylightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&taxiwaylightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiwaylightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaylightsystemtimeslicetype.LightedTaxiway) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayLightSystemType
func (taxiwaylightsystemtype *TaxiwayLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaylightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayMarkingTimeSlicePropertyType
func (taxiwaymarkingtimeslicepropertytype *TaxiwayMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaymarkingtimeslicepropertytype.TaxiwayMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayMarkingTimeSliceType
func (taxiwaymarkingtimeslicetype *TaxiwayMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaymarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&taxiwaymarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiwaymarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaymarkingtimeslicetype.MarkedTaxiway) || modified
	modified = stage.CleanPointer(&taxiwaymarkingtimeslicetype.MarkedElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayMarkingType
func (taxiwaymarkingtype *TaxiwayMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaymarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayTimeSlicePropertyType
func (taxiwaytimeslicepropertytype *TaxiwayTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaytimeslicepropertytype.TaxiwayTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayTimeSliceType
func (taxiwaytimeslicetype *TaxiwayTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaytimeslicetype.Contaminant) || modified
	modified = stage.CleanSlice(&taxiwaytimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&taxiwaytimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&taxiwaytimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&taxiwaytimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&taxiwaytimeslicetype.AssociatedAirportHeliport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaxiwayType
func (taxiwaytype *TaxiwayType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taxiwaytype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffContaminationPropertyType
func (touchdownliftoffcontaminationpropertytype *TouchDownLiftOffContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftoffcontaminationpropertytype.TouchDownLiftOffContamination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffContaminationType
func (touchdownliftoffcontaminationtype *TouchDownLiftOffContaminationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftoffcontaminationtype.CriticalRidge) || modified
	modified = stage.CleanSlice(&touchdownliftoffcontaminationtype.Layer) || modified
	modified = stage.CleanSlice(&touchdownliftoffcontaminationtype.Annotation) || modified
	modified = stage.CleanSlice(&touchdownliftoffcontaminationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffLightSystemTimeSlicePropertyType
func (touchdownliftofflightsystemtimeslicepropertytype *TouchDownLiftOffLightSystemTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftofflightsystemtimeslicepropertytype.TouchDownLiftOffLightSystemTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffLightSystemTimeSliceType
func (touchdownliftofflightsystemtimeslicetype *TouchDownLiftOffLightSystemTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftofflightsystemtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&touchdownliftofflightsystemtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&touchdownliftofflightsystemtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&touchdownliftofflightsystemtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftofflightsystemtimeslicetype.LightedTouchDownLiftOff) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffLightSystemType
func (touchdownliftofflightsystemtype *TouchDownLiftOffLightSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftofflightsystemtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffMarkingTimeSlicePropertyType
func (touchdownliftoffmarkingtimeslicepropertytype *TouchDownLiftOffMarkingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftoffmarkingtimeslicepropertytype.TouchDownLiftOffMarkingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffMarkingTimeSliceType
func (touchdownliftoffmarkingtimeslicetype *TouchDownLiftOffMarkingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftoffmarkingtimeslicetype.Element) || modified
	modified = stage.CleanSlice(&touchdownliftoffmarkingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&touchdownliftoffmarkingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftoffmarkingtimeslicetype.MarkedTouchDownLiftOff) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffMarkingType
func (touchdownliftoffmarkingtype *TouchDownLiftOffMarkingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftoffmarkingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffSafeAreaTimeSlicePropertyType
func (touchdownliftoffsafeareatimeslicepropertytype *TouchDownLiftOffSafeAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftoffsafeareatimeslicepropertytype.TouchDownLiftOffSafeAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffSafeAreaTimeSliceType
func (touchdownliftoffsafeareatimeslicetype *TouchDownLiftOffSafeAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftoffsafeareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&touchdownliftoffsafeareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftoffsafeareatimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&touchdownliftoffsafeareatimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&touchdownliftoffsafeareatimeslicetype.ProtectedTouchDownLiftOff) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffSafeAreaType
func (touchdownliftoffsafeareatype *TouchDownLiftOffSafeAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftoffsafeareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffTimeSlicePropertyType
func (touchdownliftofftimeslicepropertytype *TouchDownLiftOffTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftofftimeslicepropertytype.TouchDownLiftOffTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffTimeSliceType
func (touchdownliftofftimeslicetype *TouchDownLiftOffTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftofftimeslicetype.Contaminant) || modified
	modified = stage.CleanSlice(&touchdownliftofftimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&touchdownliftofftimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&touchdownliftofftimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&touchdownliftofftimeslicetype.AimingPoint) || modified
	modified = stage.CleanPointer(&touchdownliftofftimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&touchdownliftofftimeslicetype.SurfaceProperties) || modified
	modified = stage.CleanPointer(&touchdownliftofftimeslicetype.AssociatedAirportHeliport) || modified
	modified = stage.CleanPointer(&touchdownliftofftimeslicetype.ApproachTakeOffArea) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TouchDownLiftOffType
func (touchdownliftofftype *TouchDownLiftOffType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&touchdownliftofftype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by VisualGlideSlopeIndicatorTimeSlicePropertyType
func (visualglideslopeindicatortimeslicepropertytype *VisualGlideSlopeIndicatorTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&visualglideslopeindicatortimeslicepropertytype.VisualGlideSlopeIndicatorTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VisualGlideSlopeIndicatorTimeSliceType
func (visualglideslopeindicatortimeslicetype *VisualGlideSlopeIndicatorTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&visualglideslopeindicatortimeslicetype.Element) || modified
	modified = stage.CleanSlice(&visualglideslopeindicatortimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&visualglideslopeindicatortimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&visualglideslopeindicatortimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&visualglideslopeindicatortimeslicetype.RunwayDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VisualGlideSlopeIndicatorType
func (visualglideslopeindicatortype *VisualGlideSlopeIndicatorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&visualglideslopeindicatortype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by WorkAreaTimeSlicePropertyType
func (workareatimeslicepropertytype *WorkAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&workareatimeslicepropertytype.WorkAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by WorkAreaTimeSliceType
func (workareatimeslicetype *WorkAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&workareatimeslicetype.Activation) || modified
	modified = stage.CleanSlice(&workareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&workareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&workareatimeslicetype.AssociatedAirportHeliport) || modified
	modified = stage.CleanPointer(&workareatimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by WorkAreaType
func (workareatype *WorkAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&workareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by WorkareaActivityPropertyType
func (workareaactivitypropertytype *WorkareaActivityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&workareaactivitypropertytype.WorkareaActivity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by WorkareaActivityType
func (workareaactivitytype *WorkareaActivityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&workareaactivitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&workareaactivitytype.Annotation) || modified
	modified = stage.CleanSlice(&workareaactivitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&workareaactivitytype.Extension) || modified
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
