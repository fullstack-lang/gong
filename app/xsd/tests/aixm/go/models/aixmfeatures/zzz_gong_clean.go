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
// Clean garbage collect unstaged instances that are referenced by AerialRefuellingAnchorPropertyType
func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellinganchorpropertytype.AerialRefuellingAnchor) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingAnchorType
func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aerialrefuellinganchortype.VerticalExtent) || modified
	modified = stage.CleanSlice(&aerialrefuellinganchortype.Point) || modified
	modified = stage.CleanSlice(&aerialrefuellinganchortype.Annotation) || modified
	modified = stage.CleanSlice(&aerialrefuellinganchortype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellinganchortype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingPointPropertyType
func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingpointpropertytype.AerialRefuellingPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingPointType
func (aerialrefuellingpointtype *AerialRefuellingPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aerialrefuellingpointtype.FacilityMakeup) || modified
	modified = stage.CleanSlice(&aerialrefuellingpointtype.Annotation) || modified
	modified = stage.CleanSlice(&aerialrefuellingpointtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_navaidSystem) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_position) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_runwayPoint) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_aimingPoint) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.PointChoice_airportReferencePoint) || modified
	modified = stage.CleanPointer(&aerialrefuellingpointtype.ExtendedServiceVolume) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingTimeSlicePropertyType
func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingtimeslicepropertytype.AerialRefuellingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingTimeSliceType
func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.ProtectingAirspace) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.Track) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.Anchor) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.ManagingOrganisation) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&aerialrefuellingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingtimeslicetype.OppositeTrack) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingTrackPropertyType
func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingtrackpropertytype.AerialRefuellingTrack) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingTrackType
func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aerialrefuellingtracktype.Point) || modified
	modified = stage.CleanSlice(&aerialrefuellingtracktype.VerticalExtent) || modified
	modified = stage.CleanSlice(&aerialrefuellingtracktype.Annotation) || modified
	modified = stage.CleanSlice(&aerialrefuellingtracktype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aerialrefuellingtracktype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AerialRefuellingType
func (aerialrefuellingtype *AerialRefuellingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aerialrefuellingtype.TimeSlice) || modified
	// insertion point per field
	return
}

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

// Clean garbage collect unstaged instances that are referenced by AirTrafficControlServiceTimeSlicePropertyType
func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airtrafficcontrolservicetimeslicepropertytype.AirTrafficControlServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirTrafficControlServiceTimeSliceType
func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientAirport) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientAirspace) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientRoute) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientProcedure) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientHolding) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.ClientAerialRefuelling) || modified
	modified = stage.CleanSlice(&airtrafficcontrolservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airtrafficcontrolservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&airtrafficcontrolservicetimeslicetype.ServiceProvider) || modified
	modified = stage.CleanPointer(&airtrafficcontrolservicetimeslicetype.AircraftLocator) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirTrafficControlServiceType
func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airtrafficcontrolservicetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AirTrafficManagementServiceTimeSlicePropertyType
func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airtrafficmanagementservicetimeslicepropertytype.AirTrafficManagementServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirTrafficManagementServiceTimeSliceType
func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.ClientAirspace) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.ClientAerialRefuelling) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.ClientRoute) || modified
	modified = stage.CleanSlice(&airtrafficmanagementservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airtrafficmanagementservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&airtrafficmanagementservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirTrafficManagementServiceType
func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airtrafficmanagementservicetype.TimeSlice) || modified
	// insertion point per field
	return
}

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

// Clean garbage collect unstaged instances that are referenced by AircraftGroundServiceTimeSlicePropertyType
func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&aircraftgroundservicetimeslicepropertytype.AircraftGroundServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftGroundServiceTimeSliceType
func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&aircraftgroundservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&aircraftgroundservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&aircraftgroundservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AircraftGroundServiceType
func (aircraftgroundservicetype *AircraftGroundServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&aircraftgroundservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by AirportClearanceServiceTimeSlicePropertyType
func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportclearanceservicetimeslicepropertytype.AirportClearanceServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportClearanceServiceTimeSliceType
func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&airportclearanceservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportclearanceservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&airportclearanceservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportClearanceServiceType
func (airportclearanceservicetype *AirportClearanceServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportclearanceservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by AirportSuppliesServiceTimeSlicePropertyType
func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airportsuppliesservicetimeslicepropertytype.AirportSuppliesServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportSuppliesServiceTimeSliceType
func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.FuelSupply) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.OilSupply) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.NitrogenSupply) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.OxygenSupply) || modified
	modified = stage.CleanSlice(&airportsuppliesservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airportsuppliesservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&airportsuppliesservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirportSuppliesServiceType
func (airportsuppliesservicetype *AirportSuppliesServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airportsuppliesservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by AirspaceBorderCrossingTimeSlicePropertyType
func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&airspacebordercrossingtimeslicepropertytype.AirspaceBorderCrossingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceBorderCrossingTimeSliceType
func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacebordercrossingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&airspacebordercrossingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&airspacebordercrossingtimeslicetype.ExitedAirspace) || modified
	modified = stage.CleanPointer(&airspacebordercrossingtimeslicetype.EnteredAirspace) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AirspaceBorderCrossingType
func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&airspacebordercrossingtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by AltitudeAdjustmentPropertyType
func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&altitudeadjustmentpropertytype.AltitudeAdjustment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AltitudeAdjustmentType
func (altitudeadjustmenttype *AltitudeAdjustmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&altitudeadjustmenttype.Annotation) || modified
	modified = stage.CleanSlice(&altitudeadjustmenttype.Extension) || modified
	// insertion point per field
	return
}

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

// Clean garbage collect unstaged instances that are referenced by AngleUsePropertyType
func (angleusepropertytype *AngleUsePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&angleusepropertytype.AngleUse) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AngleUseType
func (angleusetype *AngleUseType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&angleusetype.Annotation) || modified
	modified = stage.CleanSlice(&angleusetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&angleusetype.TheAngleIndication) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachAltitudeTablePropertyType
func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&approachaltitudetablepropertytype.ApproachAltitudeTable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachAltitudeTableType
func (approachaltitudetabletype *ApproachAltitudeTableType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachaltitudetabletype.Annotation) || modified
	modified = stage.CleanSlice(&approachaltitudetabletype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachConditionPropertyType
func (approachconditionpropertytype *ApproachConditionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&approachconditionpropertytype.ApproachCondition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachConditionType
func (approachconditiontype *ApproachConditionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachconditiontype.CirclingRestriction) || modified
	modified = stage.CleanSlice(&approachconditiontype.AircraftCategory) || modified
	modified = stage.CleanSlice(&approachconditiontype.LandingArea) || modified
	modified = stage.CleanSlice(&approachconditiontype.DesignSurface) || modified
	modified = stage.CleanSlice(&approachconditiontype.Annotation) || modified
	modified = stage.CleanSlice(&approachconditiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&approachconditiontype.MinimumSet) || modified
	modified = stage.CleanPointer(&approachconditiontype.Altimeter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachDistanceTablePropertyType
func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&approachdistancetablepropertytype.ApproachDistanceTable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachDistanceTableType
func (approachdistancetabletype *ApproachDistanceTableType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachdistancetabletype.Annotation) || modified
	modified = stage.CleanSlice(&approachdistancetabletype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by ApproachTimingTablePropertyType
func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&approachtimingtablepropertytype.ApproachTimingTable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ApproachTimingTableType
func (approachtimingtabletype *ApproachTimingTableType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&approachtimingtabletype.Annotation) || modified
	modified = stage.CleanSlice(&approachtimingtabletype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by ArrivalFeederLegTimeSlicePropertyType
func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicepropertytype.ArrivalFeederLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrivalFeederLegTimeSliceType
func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrivalfeederlegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&arrivalfeederlegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&arrivalfeederlegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&arrivalfeederlegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&arrivalfeederlegtimeslicetype.Approach) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrivalFeederLegType
func (arrivalfeederlegtype *ArrivalFeederLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrivalfeederlegtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrivalLegTimeSlicePropertyType
func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arrivallegtimeslicepropertytype.ArrivalLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrivalLegTimeSliceType
func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrivallegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&arrivallegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&arrivallegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&arrivallegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arrivallegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&arrivallegtimeslicetype.Arrival) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrivalLegType
func (arrivallegtype *ArrivalLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrivallegtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForAerialRefuellingPropertyType
func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&authorityforaerialrefuellingpropertytype.AuthorityForAerialRefuelling) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AuthorityForAerialRefuellingType
func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&authorityforaerialrefuellingtype.Annotation) || modified
	modified = stage.CleanSlice(&authorityforaerialrefuellingtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&authorityforaerialrefuellingtype.TheOrganisationAuthority) || modified
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

// Clean garbage collect unstaged instances that are referenced by CallsignDetailPropertyType
func (callsigndetailpropertytype *CallsignDetailPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&callsigndetailpropertytype.CallsignDetail) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CallsignDetailType
func (callsigndetailtype *CallsignDetailType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&callsigndetailtype.Annotation) || modified
	modified = stage.CleanSlice(&callsigndetailtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ChangeOverPointTimeSlicePropertyType
func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&changeoverpointtimeslicepropertytype.ChangeOverPointTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ChangeOverPointTimeSliceType
func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&changeoverpointtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&changeoverpointtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_navaidSystem) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_position) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_runwayPoint) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_aimingPoint) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.Location_airportReferencePoint) || modified
	modified = stage.CleanPointer(&changeoverpointtimeslicetype.ApplicableRoutePortion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ChangeOverPointType
func (changeoverpointtype *ChangeOverPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&changeoverpointtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by CircleSectorPropertyType
func (circlesectorpropertytype *CircleSectorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&circlesectorpropertytype.CircleSector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleSectorType
func (circlesectortype *CircleSectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circlesectortype.Annotation) || modified
	modified = stage.CleanSlice(&circlesectortype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CirclingAreaTimeSlicePropertyType
func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&circlingareatimeslicepropertytype.CirclingAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CirclingAreaTimeSliceType
func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circlingareatimeslicetype.Condition) || modified
	modified = stage.CleanSlice(&circlingareatimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&circlingareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&circlingareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&circlingareatimeslicetype.Extent) || modified
	modified = stage.CleanPointer(&circlingareatimeslicetype.Approach) || modified
	modified = stage.CleanPointer(&circlingareatimeslicetype.AircraftCategory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CirclingAreaType
func (circlingareatype *CirclingAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circlingareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CirclingRestrictionPropertyType
func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&circlingrestrictionpropertytype.CirclingRestriction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CirclingRestrictionType
func (circlingrestrictiontype *CirclingRestrictionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circlingrestrictiontype.TimeInterval) || modified
	modified = stage.CleanSlice(&circlingrestrictiontype.Annotation) || modified
	modified = stage.CleanSlice(&circlingrestrictiontype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&circlingrestrictiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&circlingrestrictiontype.SectorDescription) || modified
	modified = stage.CleanPointer(&circlingrestrictiontype.RestrictionArea) || modified
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

// Clean garbage collect unstaged instances that are referenced by DepartureArrivalConditionPropertyType
func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&departurearrivalconditionpropertytype.DepartureArrivalCondition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DepartureArrivalConditionType
func (departurearrivalconditiontype *DepartureArrivalConditionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&departurearrivalconditiontype.Annotation) || modified
	modified = stage.CleanSlice(&departurearrivalconditiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&departurearrivalconditiontype.EngineType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DepartureLegTimeSlicePropertyType
func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&departurelegtimeslicepropertytype.DepartureLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DepartureLegTimeSliceType
func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&departurelegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&departurelegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&departurelegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&departurelegtimeslicetype.Condition) || modified
	modified = stage.CleanSlice(&departurelegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&departurelegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&departurelegtimeslicetype.Departure) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DepartureLegType
func (departurelegtype *DepartureLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&departurelegtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by DirectFlightClassPropertyType
func (directflightclasspropertytype *DirectFlightClassPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directflightclasspropertytype.DirectFlightClass) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectFlightClassType
func (directflightclasstype *DirectFlightClassType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directflightclasstype.Annotation) || modified
	modified = stage.CleanSlice(&directflightclasstype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectFlightPropertyType
func (directflightpropertytype *DirectFlightPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directflightpropertytype.DirectFlightClass) || modified
	modified = stage.CleanPointer(&directflightpropertytype.DirectFlightSegment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectFlightSegmentPropertyType
func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directflightsegmentpropertytype.DirectFlightSegment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectFlightSegmentType
func (directflightsegmenttype *DirectFlightSegmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directflightsegmenttype.Annotation) || modified
	modified = stage.CleanSlice(&directflightsegmenttype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&directflightsegmenttype.End_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.End_navaidSystem) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.End_position) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.End_runwayPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.End_aimingPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.End_airportReferencePoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_navaidSystem) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_position) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_runwayPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_aimingPoint) || modified
	modified = stage.CleanPointer(&directflightsegmenttype.Start_airportReferencePoint) || modified
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

// Clean garbage collect unstaged instances that are referenced by EnRouteSegmentPointPropertyType
func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&enroutesegmentpointpropertytype.EnRouteSegmentPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EnRouteSegmentPointType
func (enroutesegmentpointtype *EnRouteSegmentPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&enroutesegmentpointtype.FacilityMakeup) || modified
	modified = stage.CleanSlice(&enroutesegmentpointtype.Annotation) || modified
	modified = stage.CleanSlice(&enroutesegmentpointtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_navaidSystem) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_position) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_runwayPoint) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_aimingPoint) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.PointChoice_airportReferencePoint) || modified
	modified = stage.CleanPointer(&enroutesegmentpointtype.ExtendedServiceVolume) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EquipmentUnavailableAdjustmentColumnPropertyType
func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&equipmentunavailableadjustmentcolumnpropertytype.EquipmentUnavailableAdjustmentColumn) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EquipmentUnavailableAdjustmentColumnType
func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&equipmentunavailableadjustmentcolumntype.Annotation) || modified
	modified = stage.CleanSlice(&equipmentunavailableadjustmentcolumntype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EquipmentUnavailableAdjustmentPropertyType
func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&equipmentunavailableadjustmentpropertytype.EquipmentUnavailableAdjustment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EquipmentUnavailableAdjustmentType
func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&equipmentunavailableadjustmenttype.AdjustmentINOPCol) || modified
	modified = stage.CleanSlice(&equipmentunavailableadjustmenttype.Annotation) || modified
	modified = stage.CleanSlice(&equipmentunavailableadjustmenttype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FASDataBlockPropertyType
func (fasdatablockpropertytype *FASDataBlockPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&fasdatablockpropertytype.FASDataBlock) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FASDataBlockType
func (fasdatablocktype *FASDataBlockType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&fasdatablocktype.Annotation) || modified
	modified = stage.CleanSlice(&fasdatablocktype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FinalLegTimeSlicePropertyType
func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&finallegtimeslicepropertytype.FinalLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FinalLegTimeSliceType
func (finallegtimeslicetype *FinalLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&finallegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&finallegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&finallegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&finallegtimeslicetype.Condition) || modified
	modified = stage.CleanSlice(&finallegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&finallegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.Approach) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_navaidSystem) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_position) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_runwayPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_aimingPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FinalPathAlignmentPoint_airportReferencePoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.VisualDescentPoint) || modified
	modified = stage.CleanPointer(&finallegtimeslicetype.FASData) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FinalLegType
func (finallegtype *FinalLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&finallegtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FinalProfilePropertyType
func (finalprofilepropertytype *FinalProfilePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&finalprofilepropertytype.FinalProfile) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FinalProfileType
func (finalprofiletype *FinalProfileType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&finalprofiletype.Altitude) || modified
	modified = stage.CleanSlice(&finalprofiletype.Distance) || modified
	modified = stage.CleanSlice(&finalprofiletype.Timing) || modified
	modified = stage.CleanSlice(&finalprofiletype.Annotation) || modified
	modified = stage.CleanSlice(&finalprofiletype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FireFightingServiceTimeSlicePropertyType
func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&firefightingservicetimeslicepropertytype.FireFightingServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FireFightingServiceTimeSliceType
func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&firefightingservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&firefightingservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&firefightingservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FireFightingServiceType
func (firefightingservicetype *FireFightingServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&firefightingservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by FlightConditionCircumstancePropertyType
func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightconditioncircumstancepropertytype.FlightConditionCircumstance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightConditionCircumstanceType
func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightconditioncircumstancetype.Annotation) || modified
	modified = stage.CleanSlice(&flightconditioncircumstancetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightConditionCombinationPropertyType
func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightconditioncombinationpropertytype.FlightConditionCombination) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightConditionCombinationType
func (flightconditioncombinationtype *FlightConditionCombinationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightconditioncombinationtype.TimeInterval) || modified
	modified = stage.CleanSlice(&flightconditioncombinationtype.Annotation) || modified
	modified = stage.CleanSlice(&flightconditioncombinationtype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&flightconditioncombinationtype.Element) || modified
	modified = stage.CleanSlice(&flightconditioncombinationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightConditionElementPropertyType
func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightconditionelementpropertytype.FlightConditionElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightConditionElementType
func (flightconditionelementtype *FlightConditionElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightconditionelementtype.FlightLevel) || modified
	modified = stage.CleanSlice(&flightconditionelementtype.Annotation) || modified
	modified = stage.CleanSlice(&flightconditionelementtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_weather) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_aircraft) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_flight) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_directFlightCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_borderCrossingCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_operand) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_routePortionCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_standardInstrumentDepartureCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_standardInstrumentArrivalCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_organisationCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_navaidSystem) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_position) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_runwayPoint) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_aimingPoint) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.SignificantPointCondition_airportReferencePoint) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_airspaceCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_airportHeliportCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.FlightCondition_aerialRefuellingCondition) || modified
	modified = stage.CleanPointer(&flightconditionelementtype.OperationalCondition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionLevelPropertyType
func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightrestrictionlevelpropertytype.FlightRestrictionLevel) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionLevelType
func (flightrestrictionleveltype *FlightRestrictionLevelType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightrestrictionleveltype.Annotation) || modified
	modified = stage.CleanSlice(&flightrestrictionleveltype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionRoutePropertyType
func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightrestrictionroutepropertytype.FlightRestrictionRoute) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionRouteType
func (flightrestrictionroutetype *FlightRestrictionRouteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightrestrictionroutetype.RouteElement) || modified
	modified = stage.CleanSlice(&flightrestrictionroutetype.Contact) || modified
	modified = stage.CleanSlice(&flightrestrictionroutetype.Annotation) || modified
	modified = stage.CleanSlice(&flightrestrictionroutetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionTimeSlicePropertyType
func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightrestrictiontimeslicepropertytype.FlightRestrictionTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionTimeSliceType
func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightrestrictiontimeslicetype.RegulatedRoute) || modified
	modified = stage.CleanSlice(&flightrestrictiontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&flightrestrictiontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&flightrestrictiontimeslicetype.Flight) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRestrictionType
func (flightrestrictiontype *FlightRestrictionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightrestrictiontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRoutingElementPropertyType
func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&flightroutingelementpropertytype.FlightRoutingElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FlightRoutingElementType
func (flightroutingelementtype *FlightRoutingElementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&flightroutingelementtype.FlightLevel) || modified
	modified = stage.CleanSlice(&flightroutingelementtype.Annotation) || modified
	modified = stage.CleanSlice(&flightroutingelementtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&flightroutingelementtype.Element_directFlightElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_routePortionElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_standardInstrumentDepartureElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_standardInstrumentArrivalElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_navaidSystem) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_position) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_runwayPoint) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_aimingPoint) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.PointElement_airportReferencePoint) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_airspaceElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_airportHeliportElement) || modified
	modified = stage.CleanPointer(&flightroutingelementtype.Element_aerialRefuellingElement) || modified
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

// Clean garbage collect unstaged instances that are referenced by FuelPropertyType
func (fuelpropertytype *FuelPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&fuelpropertytype.Fuel) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FuelType
func (fueltype *FuelType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&fueltype.Annotation) || modified
	modified = stage.CleanSlice(&fueltype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by GroundTrafficControlServiceTimeSlicePropertyType
func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&groundtrafficcontrolservicetimeslicepropertytype.GroundTrafficControlServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GroundTrafficControlServiceTimeSliceType
func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&groundtrafficcontrolservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&groundtrafficcontrolservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&groundtrafficcontrolservicetimeslicetype.ServiceProvider) || modified
	modified = stage.CleanPointer(&groundtrafficcontrolservicetimeslicetype.ClientAirport) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GroundTrafficControlServiceType
func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&groundtrafficcontrolservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by HoldingAssessmentTimeSlicePropertyType
func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&holdingassessmenttimeslicepropertytype.HoldingAssessmentTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingAssessmentTimeSliceType
func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingassessmenttimeslicetype.ObstacleAssessment) || modified
	modified = stage.CleanSlice(&holdingassessmenttimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&holdingassessmenttimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&holdingassessmenttimeslicetype.HoldingPoint) || modified
	modified = stage.CleanPointer(&holdingassessmenttimeslicetype.UnplannedHolding) || modified
	modified = stage.CleanPointer(&holdingassessmenttimeslicetype.AssessedHoldingPattern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingAssessmentType
func (holdingassessmenttype *HoldingAssessmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingassessmenttype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternDistancePropertyType
func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&holdingpatterndistancepropertytype.HoldingPatternDistance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternDistanceType
func (holdingpatterndistancetype *HoldingPatternDistanceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingpatterndistancetype.Annotation) || modified
	modified = stage.CleanSlice(&holdingpatterndistancetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternDurationPropertyType
func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&holdingpatterndurationpropertytype.HoldingPatternDuration) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternDurationType
func (holdingpatterndurationtype *HoldingPatternDurationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingpatterndurationtype.Annotation) || modified
	modified = stage.CleanSlice(&holdingpatterndurationtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternTimeSlicePropertyType
func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&holdingpatterntimeslicepropertytype.HoldingPatternTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternTimeSliceType
func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingpatterntimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&holdingpatterntimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&holdingpatterntimeslicetype.OutboundLegSpan_endPoint) || modified
	modified = stage.CleanPointer(&holdingpatterntimeslicetype.OutboundLegSpan_endDistance) || modified
	modified = stage.CleanPointer(&holdingpatterntimeslicetype.OutboundLegSpan_endTime) || modified
	modified = stage.CleanPointer(&holdingpatterntimeslicetype.HoldingPoint) || modified
	modified = stage.CleanPointer(&holdingpatterntimeslicetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingPatternType
func (holdingpatterntype *HoldingPatternType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingpatterntype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingUsePropertyType
func (holdingusepropertytype *HoldingUsePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&holdingusepropertytype.HoldingUse) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HoldingUseType
func (holdingusetype *HoldingUseType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&holdingusetype.Annotation) || modified
	modified = stage.CleanSlice(&holdingusetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&holdingusetype.TheHoldingPattern) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InformationServiceTimeSlicePropertyType
func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&informationservicetimeslicepropertytype.InformationServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InformationServiceTimeSliceType
func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&informationservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.NavaidBroadcast) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientAirspace) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientAirport) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientRoute) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientProcedure) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientHolding) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.ClientAerialRefuelling) || modified
	modified = stage.CleanSlice(&informationservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&informationservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&informationservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InformationServiceType
func (informationservicetype *InformationServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&informationservicetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialLegTimeSlicePropertyType
func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&initiallegtimeslicepropertytype.InitialLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialLegTimeSliceType
func (initiallegtimeslicetype *InitialLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&initiallegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&initiallegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&initiallegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&initiallegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&initiallegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&initiallegtimeslicetype.Approach) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialLegType
func (initiallegtype *InitialLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&initiallegtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InstrumentApproachProcedureTimeSlicePropertyType
func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicepropertytype.InstrumentApproachProcedureTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InstrumentApproachProcedureTimeSliceType
func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.AircraftCharacteristic) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.FlightTransition) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.MissedInstruction) || modified
	modified = stage.CleanSlice(&instrumentapproachproceduretimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.GuidanceFacility_specialNavigationSystem) || modified
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.GuidanceFacility_radar) || modified
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.GuidanceFacility_navaid) || modified
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.SafeAltitude) || modified
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.Landing) || modified
	modified = stage.CleanPointer(&instrumentapproachproceduretimeslicetype.FinalProfile) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InstrumentApproachProcedureType
func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&instrumentapproachproceduretype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by IntermediateLegTimeSlicePropertyType
func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&intermediatelegtimeslicepropertytype.IntermediateLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by IntermediateLegTimeSliceType
func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&intermediatelegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&intermediatelegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&intermediatelegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&intermediatelegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&intermediatelegtimeslicetype.Approach) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by IntermediateLegType
func (intermediatelegtype *IntermediateLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&intermediatelegtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LandingTakeoffAreaCollectionPropertyType
func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&landingtakeoffareacollectionpropertytype.LandingTakeoffAreaCollection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LandingTakeoffAreaCollectionType
func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&landingtakeoffareacollectiontype.Runway) || modified
	modified = stage.CleanSlice(&landingtakeoffareacollectiontype.TLOF) || modified
	modified = stage.CleanSlice(&landingtakeoffareacollectiontype.Annotation) || modified
	modified = stage.CleanSlice(&landingtakeoffareacollectiontype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LightActivationPropertyType
func (lightactivationpropertytype *LightActivationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&lightactivationpropertytype.LightActivation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LightActivationType
func (lightactivationtype *LightActivationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&lightactivationtype.Annotation) || modified
	modified = stage.CleanSlice(&lightactivationtype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by MinimaPropertyType
func (minimapropertytype *MinimaPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&minimapropertytype.Minima) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MinimaType
func (minimatype *MinimaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&minimatype.AdjustmentINOP) || modified
	modified = stage.CleanSlice(&minimatype.Annotation) || modified
	modified = stage.CleanSlice(&minimatype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MissedApproachGroupPropertyType
func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&missedapproachgrouppropertytype.MissedApproachGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MissedApproachGroupType
func (missedapproachgrouptype *MissedApproachGroupType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&missedapproachgrouptype.Altimeter) || modified
	modified = stage.CleanSlice(&missedapproachgrouptype.Annotation) || modified
	modified = stage.CleanSlice(&missedapproachgrouptype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MissedApproachLegTimeSlicePropertyType
func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&missedapproachlegtimeslicepropertytype.MissedApproachLegTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MissedApproachLegTimeSliceType
func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&missedapproachlegtimeslicetype.AircraftCategory) || modified
	modified = stage.CleanSlice(&missedapproachlegtimeslicetype.DesignSurface) || modified
	modified = stage.CleanSlice(&missedapproachlegtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&missedapproachlegtimeslicetype.Condition) || modified
	modified = stage.CleanSlice(&missedapproachlegtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.StartPoint) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.EndPoint) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.Trajectory) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.ArcCentre) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.Angle) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.Distance) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.Holding) || modified
	modified = stage.CleanPointer(&missedapproachlegtimeslicetype.Approach) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MissedApproachLegType
func (missedapproachlegtype *MissedApproachLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&missedapproachlegtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by NavigationAreaRestrictionTimeSlicePropertyType
func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navigationarearestrictiontimeslicepropertytype.NavigationAreaRestrictionTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaRestrictionTimeSliceType
func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navigationarearestrictiontimeslicetype.Procedure) || modified
	modified = stage.CleanSlice(&navigationarearestrictiontimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&navigationarearestrictiontimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navigationarearestrictiontimeslicetype.DesignSurface) || modified
	modified = stage.CleanPointer(&navigationarearestrictiontimeslicetype.SectorDefinition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaRestrictionType
func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navigationarearestrictiontype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaSectorPropertyType
func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navigationareasectorpropertytype.NavigationAreaSector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaSectorType
func (navigationareasectortype *NavigationAreaSectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navigationareasectortype.SignificantObstacle) || modified
	modified = stage.CleanSlice(&navigationareasectortype.SectorCriteria) || modified
	modified = stage.CleanSlice(&navigationareasectortype.Annotation) || modified
	modified = stage.CleanSlice(&navigationareasectortype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navigationareasectortype.SectorDefinition) || modified
	modified = stage.CleanPointer(&navigationareasectortype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaTimeSlicePropertyType
func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&navigationareatimeslicepropertytype.NavigationAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaTimeSliceType
func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navigationareatimeslicetype.Sector) || modified
	modified = stage.CleanSlice(&navigationareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&navigationareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&navigationareatimeslicetype.Departure) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_navaidSystem) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_position) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_runwayPoint) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_aimingPoint) || modified
	modified = stage.CleanPointer(&navigationareatimeslicetype.CentrePoint_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NavigationAreaType
func (navigationareatype *NavigationAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&navigationareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NitrogenPropertyType
func (nitrogenpropertytype *NitrogenPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&nitrogenpropertytype.Nitrogen) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NitrogenType
func (nitrogentype *NitrogenType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&nitrogentype.Annotation) || modified
	modified = stage.CleanSlice(&nitrogentype.Extension) || modified
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ObstacleAssessmentAreaPropertyType
func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&obstacleassessmentareapropertytype.ObstacleAssessmentArea) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstacleAssessmentAreaType
func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obstacleassessmentareatype.AircraftCategory) || modified
	modified = stage.CleanSlice(&obstacleassessmentareatype.SignificantObstacle) || modified
	modified = stage.CleanSlice(&obstacleassessmentareatype.Annotation) || modified
	modified = stage.CleanSlice(&obstacleassessmentareatype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&obstacleassessmentareatype.Surface) || modified
	modified = stage.CleanPointer(&obstacleassessmentareatype.StartingCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstaclePlacementPropertyType
func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&obstacleplacementpropertytype.ObstaclePlacement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstaclePlacementType
func (obstacleplacementtype *ObstaclePlacementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obstacleplacementtype.Annotation) || modified
	modified = stage.CleanSlice(&obstacleplacementtype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstructionPropertyType
func (obstructionpropertytype *ObstructionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&obstructionpropertytype.Obstruction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObstructionType
func (obstructiontype *ObstructionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obstructiontype.Adjustment) || modified
	modified = stage.CleanSlice(&obstructiontype.ObstaclePlacement) || modified
	modified = stage.CleanSlice(&obstructiontype.Annotation) || modified
	modified = stage.CleanSlice(&obstructiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&obstructiontype.TheVerticalStructure) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OilPropertyType
func (oilpropertytype *OilPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&oilpropertytype.Oil) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OilType
func (oiltype *OilType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&oiltype.Annotation) || modified
	modified = stage.CleanSlice(&oiltype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by OrganisationAuthorityAssociationPropertyType
func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&organisationauthorityassociationpropertytype.OrganisationAuthorityAssociation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OrganisationAuthorityAssociationType
func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&organisationauthorityassociationtype.Annotation) || modified
	modified = stage.CleanSlice(&organisationauthorityassociationtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&organisationauthorityassociationtype.TheOrganisationAuthority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OrganisationAuthorityTimeSlicePropertyType
func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&organisationauthoritytimeslicepropertytype.OrganisationAuthorityTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OrganisationAuthorityTimeSliceType
func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&organisationauthoritytimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&organisationauthoritytimeslicetype.RelatedOrganisationAuthority) || modified
	modified = stage.CleanSlice(&organisationauthoritytimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&organisationauthoritytimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by OrganisationAuthorityType
func (organisationauthoritytype *OrganisationAuthorityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&organisationauthoritytype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by OxygenPropertyType
func (oxygenpropertytype *OxygenPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&oxygenpropertytype.Oxygen) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OxygenType
func (oxygentype *OxygenType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&oxygentype.Annotation) || modified
	modified = stage.CleanSlice(&oxygentype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by PassengerServiceTimeSlicePropertyType
func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&passengerservicetimeslicepropertytype.PassengerServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassengerServiceTimeSliceType
func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&passengerservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&passengerservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&passengerservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&passengerservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassengerServiceType
func (passengerservicetype *PassengerServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&passengerservicetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PilotControlledLightingTimeSlicePropertyType
func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pilotcontrolledlightingtimeslicepropertytype.PilotControlledLightingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PilotControlledLightingTimeSliceType
func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pilotcontrolledlightingtimeslicetype.ControlledLightIntensity) || modified
	modified = stage.CleanSlice(&pilotcontrolledlightingtimeslicetype.ActivatedGroundLighting) || modified
	modified = stage.CleanSlice(&pilotcontrolledlightingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&pilotcontrolledlightingtimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PilotControlledLightingType
func (pilotcontrolledlightingtype *PilotControlledLightingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pilotcontrolledlightingtype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by PointReferencePropertyType
func (pointreferencepropertytype *PointReferencePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pointreferencepropertytype.PointReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PointReferenceType
func (pointreferencetype *PointReferenceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointreferencetype.FacilityAngle) || modified
	modified = stage.CleanSlice(&pointreferencetype.FacilityDistance) || modified
	modified = stage.CleanSlice(&pointreferencetype.Annotation) || modified
	modified = stage.CleanSlice(&pointreferencetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&pointreferencetype.Point) || modified
	modified = stage.CleanPointer(&pointreferencetype.FixToleranceArea) || modified
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

// Clean garbage collect unstaged instances that are referenced by ProcedureAvailabilityPropertyType
func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&procedureavailabilitypropertytype.ProcedureAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureAvailabilityType
func (procedureavailabilitytype *ProcedureAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&procedureavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&procedureavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&procedureavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&procedureavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureDMETimeSlicePropertyType
func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&proceduredmetimeslicepropertytype.ProcedureDMETimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureDMETimeSliceType
func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&proceduredmetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&proceduredmetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&proceduredmetimeslicetype.DME) || modified
	modified = stage.CleanPointer(&proceduredmetimeslicetype.SegmentLeg) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureDMEType
func (proceduredmetype *ProcedureDMEType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&proceduredmetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureTransitionLegPropertyType
func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&proceduretransitionlegpropertytype.ProcedureTransitionLeg) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureTransitionLegType
func (proceduretransitionlegtype *ProcedureTransitionLegType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&proceduretransitionlegtype.Annotation) || modified
	modified = stage.CleanSlice(&proceduretransitionlegtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&proceduretransitionlegtype.TheSegmentLeg) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureTransitionPropertyType
func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&proceduretransitionpropertytype.ProcedureTransition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedureTransitionType
func (proceduretransitiontype *ProcedureTransitionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&proceduretransitiontype.TransitionLeg) || modified
	modified = stage.CleanSlice(&proceduretransitiontype.Annotation) || modified
	modified = stage.CleanSlice(&proceduretransitiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&proceduretransitiontype.DepartureRunwayTransition) || modified
	modified = stage.CleanPointer(&proceduretransitiontype.Trajectory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PropertiesWithSchedulePropertyType
func (propertieswithschedulepropertytype *PropertiesWithSchedulePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.AirportHeliportAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.AirportHeliportResponsibilityOrganisation) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.AltimeterSourceStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.ConditionCombination) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.WorkareaActivity) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.ApronAreaAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.GroundLightingAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.ManoeuvringAreaAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.RunwayDeclaredDistanceValue) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.AirspaceActivation) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.AirspaceLayerClass) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.NavaidEquipmentMonitoring) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.NavaidOperationalStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.SpecialNavigationStationStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.VerticalStructureLightingStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.VerticalStructurePart) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.UnitAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.CirclingRestriction) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.ProcedureAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.RouteAvailability) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.FlightConditionCombination) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.RadioCommunicationOperationalStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.ServiceOperationalStatus) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.OnlineContact) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.PostalAddress) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.TelephoneContact) || modified
	modified = stage.CleanPointer(&propertieswithschedulepropertytype.LightElementStatus) || modified
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

// Clean garbage collect unstaged instances that are referenced by RadioCommunicationChannelTimeSlicePropertyType
func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&radiocommunicationchanneltimeslicepropertytype.RadioCommunicationChannelTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioCommunicationChannelTimeSliceType
func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radiocommunicationchanneltimeslicetype.Location) || modified
	modified = stage.CleanSlice(&radiocommunicationchanneltimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&radiocommunicationchanneltimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&radiocommunicationchanneltimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioCommunicationChannelType
func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radiocommunicationchanneltype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioCommunicationOperationalStatusPropertyType
func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&radiocommunicationoperationalstatuspropertytype.RadioCommunicationOperationalStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioCommunicationOperationalStatusType
func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radiocommunicationoperationalstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&radiocommunicationoperationalstatustype.Annotation) || modified
	modified = stage.CleanSlice(&radiocommunicationoperationalstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&radiocommunicationoperationalstatustype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioFrequencyAreaTimeSlicePropertyType
func (radiofrequencyareatimeslicepropertytype *RadioFrequencyAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&radiofrequencyareatimeslicepropertytype.RadioFrequencyAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioFrequencyAreaTimeSliceType
func (radiofrequencyareatimeslicetype *RadioFrequencyAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radiofrequencyareatimeslicetype.Sector) || modified
	modified = stage.CleanSlice(&radiofrequencyareatimeslicetype.Extent) || modified
	modified = stage.CleanSlice(&radiofrequencyareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&radiofrequencyareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&radiofrequencyareatimeslicetype.Equipment_radar) || modified
	modified = stage.CleanPointer(&radiofrequencyareatimeslicetype.Equipment_precisionApproachRadar) || modified
	modified = stage.CleanPointer(&radiofrequencyareatimeslicetype.Equipment_frequency) || modified
	modified = stage.CleanPointer(&radiofrequencyareatimeslicetype.Equipment_specialNavigationStation) || modified
	modified = stage.CleanPointer(&radiofrequencyareatimeslicetype.Equipment_navaidEquipment) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RadioFrequencyAreaType
func (radiofrequencyareatype *RadioFrequencyAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&radiofrequencyareatype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by RouteAvailabilityPropertyType
func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&routeavailabilitypropertytype.RouteAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteAvailabilityType
func (routeavailabilitytype *RouteAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routeavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&routeavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&routeavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&routeavailabilitytype.Levels) || modified
	modified = stage.CleanSlice(&routeavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteDMETimeSlicePropertyType
func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&routedmetimeslicepropertytype.RouteDMETimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteDMETimeSliceType
func (routedmetimeslicetype *RouteDMETimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routedmetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&routedmetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&routedmetimeslicetype.ReferencedDME) || modified
	modified = stage.CleanPointer(&routedmetimeslicetype.ApplicableRoutePortion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteDMEType
func (routedmetype *RouteDMEType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routedmetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RoutePortionPropertyType
func (routeportionpropertytype *RoutePortionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&routeportionpropertytype.RoutePortion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RoutePortionType
func (routeportiontype *RoutePortionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routeportiontype.Annotation) || modified
	modified = stage.CleanSlice(&routeportiontype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&routeportiontype.Start_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.Start_navaidSystem) || modified
	modified = stage.CleanPointer(&routeportiontype.Start_position) || modified
	modified = stage.CleanPointer(&routeportiontype.Start_runwayPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.Start_aimingPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.Start_airportReferencePoint) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_navaidSystem) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_position) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_runwayPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_aimingPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.IntermediatePoint_airportReferencePoint) || modified
	modified = stage.CleanPointer(&routeportiontype.ReferencedRoute) || modified
	modified = stage.CleanPointer(&routeportiontype.End_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.End_navaidSystem) || modified
	modified = stage.CleanPointer(&routeportiontype.End_position) || modified
	modified = stage.CleanPointer(&routeportiontype.End_runwayPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.End_aimingPoint) || modified
	modified = stage.CleanPointer(&routeportiontype.End_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteSegmentTimeSlicePropertyType
func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&routesegmenttimeslicepropertytype.RouteSegmentTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteSegmentTimeSliceType
func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routesegmenttimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&routesegmenttimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&routesegmenttimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&routesegmenttimeslicetype.Start) || modified
	modified = stage.CleanPointer(&routesegmenttimeslicetype.RouteFormed) || modified
	modified = stage.CleanPointer(&routesegmenttimeslicetype.EvaluationArea) || modified
	modified = stage.CleanPointer(&routesegmenttimeslicetype.CurveExtent) || modified
	modified = stage.CleanPointer(&routesegmenttimeslicetype.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteSegmentType
func (routesegmenttype *RouteSegmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routesegmenttype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteTimeSlicePropertyType
func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&routetimeslicepropertytype.RouteTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteTimeSliceType
func (routetimeslicetype *RouteTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&routetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&routetimeslicetype.UserOrganisation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RouteType
func (routetype *RouteType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&routetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RulesProceduresTimeSlicePropertyType
func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rulesprocedurestimeslicepropertytype.RulesProceduresTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RulesProceduresTimeSliceType
func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rulesprocedurestimeslicetype.AffectedLocation) || modified
	modified = stage.CleanSlice(&rulesprocedurestimeslicetype.AffectedArea) || modified
	modified = stage.CleanSlice(&rulesprocedurestimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&rulesprocedurestimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RulesProceduresType
func (rulesprocedurestype *RulesProceduresType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rulesprocedurestype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by SafeAltitudeAreaSectorPropertyType
func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&safealtitudeareasectorpropertytype.SafeAltitudeAreaSector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SafeAltitudeAreaSectorType
func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&safealtitudeareasectortype.SignificantObstacle) || modified
	modified = stage.CleanSlice(&safealtitudeareasectortype.Annotation) || modified
	modified = stage.CleanSlice(&safealtitudeareasectortype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&safealtitudeareasectortype.Extent) || modified
	modified = stage.CleanPointer(&safealtitudeareasectortype.SectorDefinition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SafeAltitudeAreaTimeSlicePropertyType
func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&safealtitudeareatimeslicepropertytype.SafeAltitudeAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SafeAltitudeAreaTimeSliceType
func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&safealtitudeareatimeslicetype.Sector) || modified
	modified = stage.CleanSlice(&safealtitudeareatimeslicetype.Location) || modified
	modified = stage.CleanSlice(&safealtitudeareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&safealtitudeareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_navaidSystem) || modified
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_position) || modified
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_runwayPoint) || modified
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_aimingPoint) || modified
	modified = stage.CleanPointer(&safealtitudeareatimeslicetype.CentrePoint_airportReferencePoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SafeAltitudeAreaType
func (safealtitudeareatype *SafeAltitudeAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&safealtitudeareatype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by SearchRescueServiceTimeSlicePropertyType
func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&searchrescueservicetimeslicepropertytype.SearchRescueServiceTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SearchRescueServiceTimeSliceType
func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.Call_sign) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.RadioCommunication) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.GroundCommunication) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.ClientAirspace) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.ClientRoute) || modified
	modified = stage.CleanSlice(&searchrescueservicetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&searchrescueservicetimeslicetype.Location) || modified
	modified = stage.CleanPointer(&searchrescueservicetimeslicetype.ServiceProvider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SearchRescueServiceType
func (searchrescueservicetype *SearchRescueServiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&searchrescueservicetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by SectorDesignPropertyType
func (sectordesignpropertytype *SectorDesignPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&sectordesignpropertytype.SectorDesign) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SectorDesignType
func (sectordesigntype *SectorDesignType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sectordesigntype.Annotation) || modified
	modified = stage.CleanSlice(&sectordesigntype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SegmentPointPropertyType
func (segmentpointpropertytype *SegmentPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&segmentpointpropertytype.AerialRefuellingPoint) || modified
	modified = stage.CleanPointer(&segmentpointpropertytype.EnRouteSegmentPoint) || modified
	modified = stage.CleanPointer(&segmentpointpropertytype.TerminalSegmentPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ServiceOperationalStatusPropertyType
func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&serviceoperationalstatuspropertytype.ServiceOperationalStatus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ServiceOperationalStatusType
func (serviceoperationalstatustype *ServiceOperationalStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&serviceoperationalstatustype.TimeInterval) || modified
	modified = stage.CleanSlice(&serviceoperationalstatustype.Annotation) || modified
	modified = stage.CleanSlice(&serviceoperationalstatustype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&serviceoperationalstatustype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by SpecialDateTimeSlicePropertyType
func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&specialdatetimeslicepropertytype.SpecialDateTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialDateTimeSliceType
func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialdatetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&specialdatetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&specialdatetimeslicetype.Authority) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SpecialDateType
func (specialdatetype *SpecialDateType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&specialdatetype.TimeSlice) || modified
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

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentArrivalTimeSlicePropertyType
func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicepropertytype.StandardInstrumentArrivalTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentArrivalTimeSliceType
func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.AircraftCharacteristic) || modified
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.FlightTransition) || modified
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standardinstrumentarrivaltimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicetype.GuidanceFacility_specialNavigationSystem) || modified
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicetype.GuidanceFacility_radar) || modified
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicetype.GuidanceFacility_navaid) || modified
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicetype.SafeAltitude) || modified
	modified = stage.CleanPointer(&standardinstrumentarrivaltimeslicetype.Arrival) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentArrivalType
func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardinstrumentarrivaltype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentDepartureTimeSlicePropertyType
func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicepropertytype.StandardInstrumentDepartureTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentDepartureTimeSliceType
func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.AirportHeliport) || modified
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.AircraftCharacteristic) || modified
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.FlightTransition) || modified
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standardinstrumentdeparturetimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicetype.GuidanceFacility_specialNavigationSystem) || modified
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicetype.GuidanceFacility_radar) || modified
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicetype.GuidanceFacility_navaid) || modified
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicetype.SafeAltitude) || modified
	modified = stage.CleanPointer(&standardinstrumentdeparturetimeslicetype.Takeoff) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardInstrumentDepartureType
func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardinstrumentdeparturetype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelColumnTimeSlicePropertyType
func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardlevelcolumntimeslicepropertytype.StandardLevelColumnTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelColumnTimeSliceType
func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardlevelcolumntimeslicetype.Level) || modified
	modified = stage.CleanSlice(&standardlevelcolumntimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standardlevelcolumntimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&standardlevelcolumntimeslicetype.LevelTable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelColumnType
func (standardlevelcolumntype *StandardLevelColumnType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardlevelcolumntype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelPropertyType
func (standardlevelpropertytype *StandardLevelPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardlevelpropertytype.StandardLevel) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelSectorTimeSlicePropertyType
func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardlevelsectortimeslicepropertytype.StandardLevelSectorTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelSectorTimeSliceType
func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardlevelsectortimeslicetype.ApplicableAirspace) || modified
	modified = stage.CleanSlice(&standardlevelsectortimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standardlevelsectortimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&standardlevelsectortimeslicetype.ApplicableLevelColumn) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelSectorType
func (standardlevelsectortype *StandardLevelSectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardlevelsectortype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelTableTimeSlicePropertyType
func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&standardleveltabletimeslicepropertytype.StandardLevelTableTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelTableTimeSliceType
func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardleveltabletimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&standardleveltabletimeslicetype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelTableType
func (standardleveltabletype *StandardLevelTableType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardleveltabletype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StandardLevelType
func (standardleveltype *StandardLevelType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&standardleveltype.Annotation) || modified
	modified = stage.CleanSlice(&standardleveltype.Extension) || modified
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

// Clean garbage collect unstaged instances that are referenced by SurfaceContaminationPropertyType
func (surfacecontaminationpropertytype *SurfaceContaminationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.AircraftStandContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.AirportHeliportContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.ApronContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.RunwayContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.RunwaySectionContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.TaxiwayContamination) || modified
	modified = stage.CleanPointer(&surfacecontaminationpropertytype.TouchDownLiftOffContamination) || modified
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

// Clean garbage collect unstaged instances that are referenced by TerminalArrivalAreaSectorPropertyType
func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&terminalarrivalareasectorpropertytype.TerminalArrivalAreaSector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalArrivalAreaSectorType
func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&terminalarrivalareasectortype.SignificantObstacle) || modified
	modified = stage.CleanSlice(&terminalarrivalareasectortype.Annotation) || modified
	modified = stage.CleanSlice(&terminalarrivalareasectortype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&terminalarrivalareasectortype.SectorDefinition) || modified
	modified = stage.CleanPointer(&terminalarrivalareasectortype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalArrivalAreaTimeSlicePropertyType
func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&terminalarrivalareatimeslicepropertytype.TerminalArrivalAreaTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalArrivalAreaTimeSliceType
func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&terminalarrivalareatimeslicetype.Sector) || modified
	modified = stage.CleanSlice(&terminalarrivalareatimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&terminalarrivalareatimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_navaidSystem) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_position) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_runwayPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_aimingPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IF_airportReferencePoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_navaidSystem) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_position) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_runwayPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_aimingPoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.IAF_airportReferencePoint) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.Buffer) || modified
	modified = stage.CleanPointer(&terminalarrivalareatimeslicetype.ApproachRNAV) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalArrivalAreaType
func (terminalarrivalareatype *TerminalArrivalAreaType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&terminalarrivalareatype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalSegmentPointPropertyType
func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&terminalsegmentpointpropertytype.TerminalSegmentPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TerminalSegmentPointType
func (terminalsegmentpointtype *TerminalSegmentPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&terminalsegmentpointtype.FacilityMakeup) || modified
	modified = stage.CleanSlice(&terminalsegmentpointtype.Annotation) || modified
	modified = stage.CleanSlice(&terminalsegmentpointtype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_fixDesignatedPoint) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_navaidSystem) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_position) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_runwayPoint) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_aimingPoint) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.PointChoice_airportReferencePoint) || modified
	modified = stage.CleanPointer(&terminalsegmentpointtype.ExtendedServiceVolume) || modified
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

// Clean garbage collect unstaged instances that are referenced by UnitAvailabilityPropertyType
func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unitavailabilitypropertytype.UnitAvailability) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitAvailabilityType
func (unitavailabilitytype *UnitAvailabilityType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unitavailabilitytype.TimeInterval) || modified
	modified = stage.CleanSlice(&unitavailabilitytype.Annotation) || modified
	modified = stage.CleanSlice(&unitavailabilitytype.SpecialDateAuthority) || modified
	modified = stage.CleanSlice(&unitavailabilitytype.Extension) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitDependencyPropertyType
func (unitdependencypropertytype *UnitDependencyPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unitdependencypropertytype.UnitDependency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitDependencyType
func (unitdependencytype *UnitDependencyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unitdependencytype.Annotation) || modified
	modified = stage.CleanSlice(&unitdependencytype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&unitdependencytype.TheUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitTimeSlicePropertyType
func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unittimeslicepropertytype.UnitTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitTimeSliceType
func (unittimeslicetype *UnitTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unittimeslicetype.Contact) || modified
	modified = stage.CleanSlice(&unittimeslicetype.RelatedUnit) || modified
	modified = stage.CleanSlice(&unittimeslicetype.Availability) || modified
	modified = stage.CleanSlice(&unittimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&unittimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&unittimeslicetype.Position) || modified
	modified = stage.CleanPointer(&unittimeslicetype.AirportLocation) || modified
	modified = stage.CleanPointer(&unittimeslicetype.OwnerOrganisation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitType
func (unittype *UnitType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unittype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by UnplannedHoldingTimeSlicePropertyType
func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unplannedholdingtimeslicepropertytype.UnplannedHoldingTimeSlice) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnplannedHoldingTimeSliceType
func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unplannedholdingtimeslicetype.Annotation) || modified
	modified = stage.CleanSlice(&unplannedholdingtimeslicetype.Extension) || modified
	// insertion point per field
	modified = stage.CleanPointer(&unplannedholdingtimeslicetype.HoldingPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnplannedHoldingType
func (unplannedholdingtype *UnplannedHoldingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unplannedholdingtype.TimeSlice) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by UsageConditionPropertyType
func (usageconditionpropertytype *UsageConditionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&usageconditionpropertytype.AirportHeliportUsage) || modified
	modified = stage.CleanPointer(&usageconditionpropertytype.ApronAreaUsage) || modified
	modified = stage.CleanPointer(&usageconditionpropertytype.ManoeuvringAreaUsage) || modified
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
