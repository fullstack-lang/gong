// generated code - do not edit
package aixmfeatures

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.A_extensions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension", instances, (*A_extension)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_1", instances, (*A_extension_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_10s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_10", instances, (*A_extension_10)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_11s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_11", instances, (*A_extension_11)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_12s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_12", instances, (*A_extension_12)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_13s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_13", instances, (*A_extension_13)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_14s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_14", instances, (*A_extension_14)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_15s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_15", instances, (*A_extension_15)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_16s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_16", instances, (*A_extension_16)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_17s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_17", instances, (*A_extension_17)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_18s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_18", instances, (*A_extension_18)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_19s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_19", instances, (*A_extension_19)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_2s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_2", instances, (*A_extension_2)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_20s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_20", instances, (*A_extension_20)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_21s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_21", instances, (*A_extension_21)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_22s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_22", instances, (*A_extension_22)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_23s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_23", instances, (*A_extension_23)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_24s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_24", instances, (*A_extension_24)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_25s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_25", instances, (*A_extension_25)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_26s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_26", instances, (*A_extension_26)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_27s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_27", instances, (*A_extension_27)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_28s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_28", instances, (*A_extension_28)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_29s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_29", instances, (*A_extension_29)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_3s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_3", instances, (*A_extension_3)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_30s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_30", instances, (*A_extension_30)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_31s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_31", instances, (*A_extension_31)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_32s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_32", instances, (*A_extension_32)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_33s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_33", instances, (*A_extension_33)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_34s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_34", instances, (*A_extension_34)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_35s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_35", instances, (*A_extension_35)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_36s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_36", instances, (*A_extension_36)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_37s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_37", instances, (*A_extension_37)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_38s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_38", instances, (*A_extension_38)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_39s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_39", instances, (*A_extension_39)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_4s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_4", instances, (*A_extension_4)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_40s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_40", instances, (*A_extension_40)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_41s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_41", instances, (*A_extension_41)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_42s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_42", instances, (*A_extension_42)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_5s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_5", instances, (*A_extension_5)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_6s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_6", instances, (*A_extension_6)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_7s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_7", instances, (*A_extension_7)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_8s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_8", instances, (*A_extension_8)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_9s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_9", instances, (*A_extension_9)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDirectFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDirectFlightType", instances, (*AbstractDirectFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractPropertiesWithScheduleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractPropertiesWithScheduleType", instances, (*AbstractPropertiesWithScheduleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSegmentPointType", instances, (*AbstractSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingAnchorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingAnchorPropertyType", instances, (*AerialRefuellingAnchorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingAnchorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingAnchorType", instances, (*AerialRefuellingAnchorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPointPropertyType", instances, (*AerialRefuellingPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPointType", instances, (*AerialRefuellingPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPropertyType", instances, (*AerialRefuellingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTimeSlicePropertyType", instances, (*AerialRefuellingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTimeSliceType", instances, (*AerialRefuellingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTrackPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTrackPropertyType", instances, (*AerialRefuellingTrackPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTrackTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTrackType", instances, (*AerialRefuellingTrackType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingType", instances, (*AerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicPropertyType", instances, (*AircraftCharacteristicPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicType", instances, (*AircraftCharacteristicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportPropertyType", instances, (*AirportHeliportPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingPropertyType", instances, (*AirspaceBorderCrossingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingTimeSlicePropertyType", instances, (*AirspaceBorderCrossingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingTimeSliceType", instances, (*AirspaceBorderCrossingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingType", instances, (*AirspaceBorderCrossingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerPropertyType", instances, (*AirspaceLayerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerType", instances, (*AirspaceLayerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspacePropertyType", instances, (*AirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltitudeAdjustmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltitudeAdjustmentPropertyType", instances, (*AltitudeAdjustmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltitudeAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltitudeAdjustmentType", instances, (*AltitudeAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationPropertyType", instances, (*AngleIndicationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleUsePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleUsePropertyType", instances, (*AngleUsePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleUseType", instances, (*AngleUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAerialRefuellingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAerialRefuellingPropertyType", instances, (*AuthorityForAerialRefuellingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAerialRefuellingType", instances, (*AuthorityForAerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointTimeSlicePropertyType", instances, (*ChangeOverPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointTimeSliceType", instances, (*ChangeOverPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointType", instances, (*ChangeOverPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ContactInformationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ContactInformationPropertyType", instances, (*ContactInformationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ContactInformationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ContactInformationType", instances, (*ContactInformationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurvePropertyType", instances, (*CurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveType", instances, (*CurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMEPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMEPropertyType", instances, (*DMEPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointPropertyType", instances, (*DesignatedPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightClassType", instances, (*DirectFlightClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightPropertyType", instances, (*DirectFlightPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightSegmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightSegmentPropertyType", instances, (*DirectFlightSegmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightSegmentType", instances, (*DirectFlightSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationPropertyType", instances, (*DistanceIndicationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedCurveType", instances, (*ElevatedCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedPointType", instances, (*ElevatedPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedSurfaceType", instances, (*ElevatedSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnRouteSegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnRouteSegmentPointPropertyType", instances, (*EnRouteSegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnRouteSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnRouteSegmentPointType", instances, (*EnRouteSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightCharacteristicPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightCharacteristicPropertyType", instances, (*FlightCharacteristicPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightCharacteristicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightCharacteristicType", instances, (*FlightCharacteristicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCircumstancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCircumstancePropertyType", instances, (*FlightConditionCircumstancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCircumstanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCircumstanceType", instances, (*FlightConditionCircumstanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCombinationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCombinationPropertyType", instances, (*FlightConditionCombinationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCombinationType", instances, (*FlightConditionCombinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionElementPropertyType", instances, (*FlightConditionElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionElementType", instances, (*FlightConditionElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionLevelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionLevelPropertyType", instances, (*FlightRestrictionLevelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionLevelType", instances, (*FlightRestrictionLevelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionRoutePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionRoutePropertyType", instances, (*FlightRestrictionRoutePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionRouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionRouteType", instances, (*FlightRestrictionRouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionTimeSlicePropertyType", instances, (*FlightRestrictionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionTimeSliceType", instances, (*FlightRestrictionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionType", instances, (*FlightRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRoutingElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRoutingElementPropertyType", instances, (*FlightRoutingElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRoutingElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRoutingElementType", instances, (*FlightRoutingElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNotePropertyType", instances, (*LinguisticNotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNoteType", instances, (*LinguisticNoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeteorologyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeteorologyPropertyType", instances, (*MeteorologyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeteorologyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeteorologyType", instances, (*MeteorologyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidPropertyType", instances, (*NavaidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NotePropertyType", instances, (*NotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoteType", instances, (*NoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAssessmentAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAssessmentAreaPropertyType", instances, (*ObstacleAssessmentAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAssessmentAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAssessmentAreaType", instances, (*ObstacleAssessmentAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstaclePlacementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstaclePlacementPropertyType", instances, (*ObstaclePlacementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstaclePlacementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstaclePlacementType", instances, (*ObstaclePlacementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstructionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstructionPropertyType", instances, (*ObstructionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstructionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstructionType", instances, (*ObstructionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OnlineContactPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OnlineContactPropertyType", instances, (*OnlineContactPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OnlineContactTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OnlineContactType", instances, (*OnlineContactType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityPropertyType", instances, (*OrganisationAuthorityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointPropertyType", instances, (*PointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointReferencePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointReferencePropertyType", instances, (*PointReferencePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointReferenceType", instances, (*PointReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointType", instances, (*PointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PostalAddressPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PostalAddressPropertyType", instances, (*PostalAddressPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PostalAddressTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PostalAddressType", instances, (*PostalAddressType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioFrequencyAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioFrequencyAreaPropertyType", instances, (*RadioFrequencyAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteAvailabilityPropertyType", instances, (*RouteAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteAvailabilityType", instances, (*RouteAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMETimeSlicePropertyType", instances, (*RouteDMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMETimeSliceType", instances, (*RouteDMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMEType", instances, (*RouteDMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePortionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePortionPropertyType", instances, (*RoutePortionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePortionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePortionType", instances, (*RoutePortionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePropertyType", instances, (*RoutePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentTimeSlicePropertyType", instances, (*RouteSegmentTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentTimeSliceType", instances, (*RouteSegmentTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentType", instances, (*RouteSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteTimeSlicePropertyType", instances, (*RouteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteTimeSliceType", instances, (*RouteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteType", instances, (*RouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresTimeSlicePropertyType", instances, (*RulesProceduresTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresTimeSliceType", instances, (*RulesProceduresTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresType", instances, (*RulesProceduresType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointPropertyType", instances, (*RunwayCentrelinePointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalPropertyType", instances, (*StandardInstrumentArrivalPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDeparturePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDeparturePropertyType", instances, (*StandardInstrumentDeparturePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnPropertyType", instances, (*StandardLevelColumnPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfacePropertyType", instances, (*SurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceType", instances, (*SurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TelephoneContactPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TelephoneContactPropertyType", instances, (*TelephoneContactPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TelephoneContactTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TelephoneContactType", instances, (*TelephoneContactType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetPropertyType", instances, (*TimesheetPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetType", instances, (*TimesheetType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffPropertyType", instances, (*TouchDownLiftOffPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructurePropertyType", instances, (*VerticalStructurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab GongExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// GongTabulator is an interface for writing to a table strings
type GongTabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type Tabulator = GongTabulator

type GongExcelizeTabulator struct {
	f *excelize.File
}

type ExcelizeTabulator = GongExcelizeTabulator

func (tab *GongExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *GongExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *GongExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *GongExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type GongstructPtr](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type GongstructPtr](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
