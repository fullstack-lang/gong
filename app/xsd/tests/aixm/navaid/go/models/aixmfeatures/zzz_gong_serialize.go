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
			for instance := range stage.AbstractNavaidEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractNavaidEquipmentType", instances, (*AbstractNavaidEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractNavigationSystemCheckpointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractNavigationSystemCheckpointType", instances, (*AbstractNavigationSystemCheckpointType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractRadarEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractRadarEquipmentType", instances, (*AbstractRadarEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurveillanceRadarType", instances, (*AbstractSurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspacePropertyType", instances, (*AirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationTimeSlicePropertyType", instances, (*AngleIndicationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationTimeSliceType", instances, (*AngleIndicationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationType", instances, (*AngleIndicationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForNavaidEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForNavaidEquipmentPropertyType", instances, (*AuthorityForNavaidEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForNavaidEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForNavaidEquipmentType", instances, (*AuthorityForNavaidEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationStationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationStationPropertyType", instances, (*AuthorityForSpecialNavigationStationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationStationType", instances, (*AuthorityForSpecialNavigationStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationSystemPropertyType", instances, (*AuthorityForSpecialNavigationSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationSystemType", instances, (*AuthorityForSpecialNavigationSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthTimeSlicePropertyType", instances, (*AzimuthTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthTimeSliceType", instances, (*AzimuthTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthType", instances, (*AzimuthType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSTimeSlicePropertyType", instances, (*CheckpointINSTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSTimeSliceType", instances, (*CheckpointINSTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSType", instances, (*CheckpointINSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORTimeSlicePropertyType", instances, (*CheckpointVORTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORTimeSliceType", instances, (*CheckpointVORTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORType", instances, (*CheckpointVORType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.DMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMETimeSlicePropertyType", instances, (*DMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMETimeSliceType", instances, (*DMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMEType", instances, (*DMEType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.DesignatedPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointTimeSlicePropertyType", instances, (*DesignatedPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointTimeSliceType", instances, (*DesignatedPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointType", instances, (*DesignatedPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderTimeSlicePropertyType", instances, (*DirectionFinderTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderTimeSliceType", instances, (*DirectionFinderTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderType", instances, (*DirectionFinderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationTimeSlicePropertyType", instances, (*DistanceIndicationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationTimeSliceType", instances, (*DistanceIndicationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationType", instances, (*DistanceIndicationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedPointPropertyType", instances, (*ElevatedPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ElevationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationTimeSlicePropertyType", instances, (*ElevationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationTimeSliceType", instances, (*ElevationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationType", instances, (*ElevationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathTimeSlicePropertyType", instances, (*GlidepathTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathTimeSliceType", instances, (*GlidepathTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathType", instances, (*GlidepathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServicePropertyType", instances, (*InformationServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.LocalizerTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerTimeSlicePropertyType", instances, (*LocalizerTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerTimeSliceType", instances, (*LocalizerTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerType", instances, (*LocalizerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconTimeSlicePropertyType", instances, (*MarkerBeaconTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconTimeSliceType", instances, (*MarkerBeaconTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconType", instances, (*MarkerBeaconType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBTimeSlicePropertyType", instances, (*NDBTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBTimeSliceType", instances, (*NDBTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBType", instances, (*NDBType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidComponentPropertyType", instances, (*NavaidComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidComponentType", instances, (*NavaidComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentMonitoringPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentMonitoringPropertyType", instances, (*NavaidEquipmentMonitoringPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentMonitoringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentMonitoringType", instances, (*NavaidEquipmentMonitoringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentPropertyType", instances, (*NavaidEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidOperationalStatusPropertyType", instances, (*NavaidOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidOperationalStatusType", instances, (*NavaidOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.NavaidTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidTimeSlicePropertyType", instances, (*NavaidTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidTimeSliceType", instances, (*NavaidTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidType", instances, (*NavaidType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.PrecisionApproachRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarTimeSlicePropertyType", instances, (*PrecisionApproachRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarTimeSliceType", instances, (*PrecisionApproachRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarType", instances, (*PrecisionApproachRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarTimeSlicePropertyType", instances, (*PrimarySurveillanceRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarTimeSliceType", instances, (*PrimarySurveillanceRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarType", instances, (*PrimarySurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarComponentPropertyType", instances, (*RadarComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarComponentType", instances, (*RadarComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarEquipmentPropertyType", instances, (*RadarEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemTimeSlicePropertyType", instances, (*RadarSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemTimeSliceType", instances, (*RadarSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemType", instances, (*RadarSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ReflectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ReflectorPropertyType", instances, (*ReflectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ReflectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ReflectorType", instances, (*ReflectorType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayPropertyType", instances, (*RunwayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFTimeSlicePropertyType", instances, (*SDFTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFTimeSliceType", instances, (*SDFTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFType", instances, (*SDFType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarTimeSlicePropertyType", instances, (*SecondarySurveillanceRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarTimeSliceType", instances, (*SecondarySurveillanceRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarType", instances, (*SecondarySurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceTimeSlicePropertyType", instances, (*SignificantPointInAirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceTimeSliceType", instances, (*SignificantPointInAirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceType", instances, (*SignificantPointInAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationStatusPropertyType", instances, (*SpecialNavigationStationStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationStatusType", instances, (*SpecialNavigationStationStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationTimeSlicePropertyType", instances, (*SpecialNavigationStationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationTimeSliceType", instances, (*SpecialNavigationStationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationType", instances, (*SpecialNavigationStationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.SpecialNavigationSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemTimeSlicePropertyType", instances, (*SpecialNavigationSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemTimeSliceType", instances, (*SpecialNavigationSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemType", instances, (*SpecialNavigationSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveillanceGroundStationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveillanceGroundStationPropertyType", instances, (*SurveillanceGroundStationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveillanceGroundStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveillanceGroundStationType", instances, (*SurveillanceGroundStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANTimeSlicePropertyType", instances, (*TACANTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANTimeSliceType", instances, (*TACANTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANType", instances, (*TACANType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.UnitPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitPropertyType", instances, (*UnitPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORPropertyType", instances, (*VORPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORTimeSlicePropertyType", instances, (*VORTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORTimeSliceType", instances, (*VORTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORType", instances, (*VORType)(nil).GongGetFieldHeaders(), addIDs)
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
