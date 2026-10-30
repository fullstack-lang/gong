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
