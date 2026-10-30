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
