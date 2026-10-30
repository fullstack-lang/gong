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
			for instance := range stage.A_extension_4s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_4", instances, (*A_extension_4)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractAirportGroundServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractAirportGroundServiceType", instances, (*AbstractAirportGroundServiceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractServiceType", instances, (*AbstractServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTrafficSeparationServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTrafficSeparationServiceType", instances, (*AbstractTrafficSeparationServiceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AirTrafficControlServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceTimeSlicePropertyType", instances, (*AirTrafficControlServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceTimeSliceType", instances, (*AirTrafficControlServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceType", instances, (*AirTrafficControlServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceTimeSlicePropertyType", instances, (*AirTrafficManagementServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceTimeSliceType", instances, (*AirTrafficManagementServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceType", instances, (*AirTrafficManagementServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceTimeSlicePropertyType", instances, (*AircraftGroundServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceTimeSliceType", instances, (*AircraftGroundServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceType", instances, (*AircraftGroundServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceTimeSlicePropertyType", instances, (*AirportClearanceServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceTimeSliceType", instances, (*AirportClearanceServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceType", instances, (*AirportClearanceServiceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AirportSuppliesServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceTimeSlicePropertyType", instances, (*AirportSuppliesServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceTimeSliceType", instances, (*AirportSuppliesServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceType", instances, (*AirportSuppliesServiceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CallsignDetailPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CallsignDetailPropertyType", instances, (*CallsignDetailPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CallsignDetailTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CallsignDetailType", instances, (*CallsignDetailType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.DesignatedPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointPropertyType", instances, (*DesignatedPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderPropertyType", instances, (*DirectionFinderPropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.FireFightingServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceTimeSlicePropertyType", instances, (*FireFightingServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceTimeSliceType", instances, (*FireFightingServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceType", instances, (*FireFightingServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FuelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FuelPropertyType", instances, (*FuelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FuelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FuelType", instances, (*FuelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightSystemPropertyType", instances, (*GroundLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceTimeSlicePropertyType", instances, (*GroundTrafficControlServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceTimeSliceType", instances, (*GroundTrafficControlServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceType", instances, (*GroundTrafficControlServiceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.InformationServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceTimeSlicePropertyType", instances, (*InformationServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceTimeSliceType", instances, (*InformationServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceType", instances, (*InformationServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightActivationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightActivationPropertyType", instances, (*LightActivationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightActivationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightActivationType", instances, (*LightActivationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.NavaidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidPropertyType", instances, (*NavaidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NitrogenPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NitrogenPropertyType", instances, (*NitrogenPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NitrogenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NitrogenType", instances, (*NitrogenType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.OilPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OilPropertyType", instances, (*OilPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OilTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OilType", instances, (*OilType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.OrganisationAuthorityAssociationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityAssociationPropertyType", instances, (*OrganisationAuthorityAssociationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityAssociationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityAssociationType", instances, (*OrganisationAuthorityAssociationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.OrganisationAuthorityTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityTimeSlicePropertyType", instances, (*OrganisationAuthorityTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityTimeSliceType", instances, (*OrganisationAuthorityTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityType", instances, (*OrganisationAuthorityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OxygenPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OxygenPropertyType", instances, (*OxygenPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OxygenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OxygenType", instances, (*OxygenType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceTimeSlicePropertyType", instances, (*PassengerServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceTimeSliceType", instances, (*PassengerServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceType", instances, (*PassengerServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingTimeSlicePropertyType", instances, (*PilotControlledLightingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingTimeSliceType", instances, (*PilotControlledLightingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingType", instances, (*PilotControlledLightingType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedurePropertyType", instances, (*ProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelPropertyType", instances, (*RadioCommunicationChannelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelTimeSlicePropertyType", instances, (*RadioCommunicationChannelTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelTimeSliceType", instances, (*RadioCommunicationChannelTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelType", instances, (*RadioCommunicationChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationOperationalStatusPropertyType", instances, (*RadioCommunicationOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationOperationalStatusType", instances, (*RadioCommunicationOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayCentrelinePointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointPropertyType", instances, (*RunwayCentrelinePointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceTimeSlicePropertyType", instances, (*SearchRescueServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceTimeSliceType", instances, (*SearchRescueServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceType", instances, (*SearchRescueServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ServiceOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ServiceOperationalStatusPropertyType", instances, (*ServiceOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ServiceOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ServiceOperationalStatusType", instances, (*ServiceOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateTimeSlicePropertyType", instances, (*SpecialDateTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateTimeSliceType", instances, (*SpecialDateTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateType", instances, (*SpecialDateType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.StandardLevelColumnTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnTimeSlicePropertyType", instances, (*StandardLevelColumnTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnTimeSliceType", instances, (*StandardLevelColumnTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnType", instances, (*StandardLevelColumnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelPropertyType", instances, (*StandardLevelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorTimeSlicePropertyType", instances, (*StandardLevelSectorTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorTimeSliceType", instances, (*StandardLevelSectorTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorType", instances, (*StandardLevelSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTablePropertyType", instances, (*StandardLevelTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableTimeSlicePropertyType", instances, (*StandardLevelTableTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableTimeSliceType", instances, (*StandardLevelTableTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableType", instances, (*StandardLevelTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelType", instances, (*StandardLevelType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.UnitAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitAvailabilityPropertyType", instances, (*UnitAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitAvailabilityType", instances, (*UnitAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitDependencyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitDependencyPropertyType", instances, (*UnitDependencyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitDependencyType", instances, (*UnitDependencyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.UnitTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitTimeSlicePropertyType", instances, (*UnitTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitTimeSliceType", instances, (*UnitTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitType", instances, (*UnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORPropertyType", instances, (*VORPropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
