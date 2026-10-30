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
			for instance := range stage.A_extension_43s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_43", instances, (*A_extension_43)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_44s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_44", instances, (*A_extension_44)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_45s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_45", instances, (*A_extension_45)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_46s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_46", instances, (*A_extension_46)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_47s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_47", instances, (*A_extension_47)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_48s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_48", instances, (*A_extension_48)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_49s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_49", instances, (*A_extension_49)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.A_extension_50s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_50", instances, (*A_extension_50)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_51s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_51", instances, (*A_extension_51)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_52s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_52", instances, (*A_extension_52)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_53s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_53", instances, (*A_extension_53)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_54s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_54", instances, (*A_extension_54)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_55s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_55", instances, (*A_extension_55)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_56s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_56", instances, (*A_extension_56)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_57s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_57", instances, (*A_extension_57)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_58s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_58", instances, (*A_extension_58)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractApproachLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractApproachLegType", instances, (*AbstractApproachLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractProcedureType", instances, (*AbstractProcedureType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractSegmentLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSegmentLegType", instances, (*AbstractSegmentLegType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AerialRefuellingPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPointType", instances, (*AerialRefuellingPointType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AltimeterSourcePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourcePropertyType", instances, (*AltimeterSourcePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ApproachAltitudeTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachAltitudeTablePropertyType", instances, (*ApproachAltitudeTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachAltitudeTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachAltitudeTableType", instances, (*ApproachAltitudeTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachConditionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachConditionPropertyType", instances, (*ApproachConditionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachConditionType", instances, (*ApproachConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachDistanceTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachDistanceTablePropertyType", instances, (*ApproachDistanceTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachDistanceTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachDistanceTableType", instances, (*ApproachDistanceTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachTimingTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachTimingTablePropertyType", instances, (*ApproachTimingTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachTimingTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachTimingTableType", instances, (*ApproachTimingTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegTimeSlicePropertyType", instances, (*ArrivalFeederLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegTimeSliceType", instances, (*ArrivalFeederLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegType", instances, (*ArrivalFeederLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegTimeSlicePropertyType", instances, (*ArrivalLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegTimeSliceType", instances, (*ArrivalLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegType", instances, (*ArrivalLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleSectorPropertyType", instances, (*CircleSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleSectorType", instances, (*CircleSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaTimeSlicePropertyType", instances, (*CirclingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaTimeSliceType", instances, (*CirclingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaType", instances, (*CirclingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingRestrictionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingRestrictionPropertyType", instances, (*CirclingRestrictionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingRestrictionType", instances, (*CirclingRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.DepartureArrivalConditionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureArrivalConditionPropertyType", instances, (*DepartureArrivalConditionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureArrivalConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureArrivalConditionType", instances, (*DepartureArrivalConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegTimeSlicePropertyType", instances, (*DepartureLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegTimeSliceType", instances, (*DepartureLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegType", instances, (*DepartureLegType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.EnRouteSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnRouteSegmentPointType", instances, (*EnRouteSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentColumnPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentColumnPropertyType", instances, (*EquipmentUnavailableAdjustmentColumnPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentColumnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentColumnType", instances, (*EquipmentUnavailableAdjustmentColumnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentPropertyType", instances, (*EquipmentUnavailableAdjustmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentType", instances, (*EquipmentUnavailableAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FASDataBlockPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FASDataBlockPropertyType", instances, (*FASDataBlockPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FASDataBlockTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FASDataBlockType", instances, (*FASDataBlockType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegTimeSlicePropertyType", instances, (*FinalLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegTimeSliceType", instances, (*FinalLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegType", instances, (*FinalLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalProfilePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalProfilePropertyType", instances, (*FinalProfilePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalProfileTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalProfileType", instances, (*FinalProfileType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentTimeSlicePropertyType", instances, (*HoldingAssessmentTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentTimeSliceType", instances, (*HoldingAssessmentTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentType", instances, (*HoldingAssessmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDistancePropertyType", instances, (*HoldingPatternDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDistanceType", instances, (*HoldingPatternDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDurationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDurationPropertyType", instances, (*HoldingPatternDurationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDurationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDurationType", instances, (*HoldingPatternDurationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternPropertyType", instances, (*HoldingPatternPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternTimeSlicePropertyType", instances, (*HoldingPatternTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternTimeSliceType", instances, (*HoldingPatternTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternType", instances, (*HoldingPatternType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingUsePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingUsePropertyType", instances, (*HoldingUsePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingUseType", instances, (*HoldingUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegTimeSlicePropertyType", instances, (*InitialLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegTimeSliceType", instances, (*InitialLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegType", instances, (*InitialLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedurePropertyType", instances, (*InstrumentApproachProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureTimeSlicePropertyType", instances, (*InstrumentApproachProcedureTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureTimeSliceType", instances, (*InstrumentApproachProcedureTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureType", instances, (*InstrumentApproachProcedureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegTimeSlicePropertyType", instances, (*IntermediateLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegTimeSliceType", instances, (*IntermediateLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegType", instances, (*IntermediateLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LandingTakeoffAreaCollectionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LandingTakeoffAreaCollectionPropertyType", instances, (*LandingTakeoffAreaCollectionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LandingTakeoffAreaCollectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LandingTakeoffAreaCollectionType", instances, (*LandingTakeoffAreaCollectionType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.MinimaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinimaPropertyType", instances, (*MinimaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinimaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinimaType", instances, (*MinimaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachGroupPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachGroupPropertyType", instances, (*MissedApproachGroupPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachGroupTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachGroupType", instances, (*MissedApproachGroupType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegTimeSlicePropertyType", instances, (*MissedApproachLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegTimeSliceType", instances, (*MissedApproachLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegType", instances, (*MissedApproachLegType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.NavigationAreaRestrictionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionTimeSlicePropertyType", instances, (*NavigationAreaRestrictionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionTimeSliceType", instances, (*NavigationAreaRestrictionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionType", instances, (*NavigationAreaRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaSectorPropertyType", instances, (*NavigationAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaSectorType", instances, (*NavigationAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaTimeSlicePropertyType", instances, (*NavigationAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaTimeSliceType", instances, (*NavigationAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaType", instances, (*NavigationAreaType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ProcedureAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureAvailabilityPropertyType", instances, (*ProcedureAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureAvailabilityType", instances, (*ProcedureAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMETimeSlicePropertyType", instances, (*ProcedureDMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMETimeSliceType", instances, (*ProcedureDMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMEType", instances, (*ProcedureDMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedurePropertyType", instances, (*ProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionLegPropertyType", instances, (*ProcedureTransitionLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionLegType", instances, (*ProcedureTransitionLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionPropertyType", instances, (*ProcedureTransitionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionType", instances, (*ProcedureTransitionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemPropertyType", instances, (*RadarSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayCentrelinePointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointPropertyType", instances, (*RunwayCentrelinePointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionPropertyType", instances, (*RunwayDirectionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaPropertyType", instances, (*SafeAltitudeAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaSectorPropertyType", instances, (*SafeAltitudeAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaSectorType", instances, (*SafeAltitudeAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaTimeSlicePropertyType", instances, (*SafeAltitudeAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaTimeSliceType", instances, (*SafeAltitudeAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaType", instances, (*SafeAltitudeAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SectorDesignPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SectorDesignPropertyType", instances, (*SectorDesignPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SectorDesignTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SectorDesignType", instances, (*SectorDesignType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SegmentLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SegmentLegPropertyType", instances, (*SegmentLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SegmentPointPropertyType", instances, (*SegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemPropertyType", instances, (*SpecialNavigationSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.StandardInstrumentArrivalTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalTimeSlicePropertyType", instances, (*StandardInstrumentArrivalTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalTimeSliceType", instances, (*StandardInstrumentArrivalTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalType", instances, (*StandardInstrumentArrivalType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.StandardInstrumentDepartureTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureTimeSlicePropertyType", instances, (*StandardInstrumentDepartureTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDepartureTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureTimeSliceType", instances, (*StandardInstrumentDepartureTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDepartureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureType", instances, (*StandardInstrumentDepartureType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.TerminalArrivalAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaSectorPropertyType", instances, (*TerminalArrivalAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaSectorType", instances, (*TerminalArrivalAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaTimeSlicePropertyType", instances, (*TerminalArrivalAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaTimeSliceType", instances, (*TerminalArrivalAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaType", instances, (*TerminalArrivalAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalSegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalSegmentPointPropertyType", instances, (*TerminalSegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalSegmentPointType", instances, (*TerminalSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.UnplannedHoldingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingPropertyType", instances, (*UnplannedHoldingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingTimeSlicePropertyType", instances, (*UnplannedHoldingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingTimeSliceType", instances, (*UnplannedHoldingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingType", instances, (*UnplannedHoldingType)(nil).GongGetFieldHeaders(), addIDs)
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
