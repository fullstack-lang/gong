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

// Clean garbage collect unstaged instances that are referenced by ElevatedCurveType
func (elevatedcurvetype *ElevatedCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevatedcurvetype.Annotation) || modified
	modified = stage.CleanSlice(&elevatedcurvetype.Extension) || modified
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ElevatedSurfaceType
func (elevatedsurfacetype *ElevatedSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&elevatedsurfacetype.Annotation) || modified
	modified = stage.CleanSlice(&elevatedsurfacetype.Extension) || modified
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
