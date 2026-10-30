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
			for instance := range stage.A_extension_59s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_59", instances, (*A_extension_59)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.A_extension_60s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_60", instances, (*A_extension_60)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_61s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_61", instances, (*A_extension_61)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_62s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_62", instances, (*A_extension_62)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_63s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_63", instances, (*A_extension_63)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_64s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_64", instances, (*A_extension_64)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_65s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_65", instances, (*A_extension_65)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_66s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_66", instances, (*A_extension_66)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_67s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_67", instances, (*A_extension_67)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_68s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_68", instances, (*A_extension_68)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_69s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_69", instances, (*A_extension_69)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.A_extension_70s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_70", instances, (*A_extension_70)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_71s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_71", instances, (*A_extension_71)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_72s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_72", instances, (*A_extension_72)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_73s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_73", instances, (*A_extension_73)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_74s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_74", instances, (*A_extension_74)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_75s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_75", instances, (*A_extension_75)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_76s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_76", instances, (*A_extension_76)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_77s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_77", instances, (*A_extension_77)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_78s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_78", instances, (*A_extension_78)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_79s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_79", instances, (*A_extension_79)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.A_extension_80s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_80", instances, (*A_extension_80)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_81s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_81", instances, (*A_extension_81)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_82s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_82", instances, (*A_extension_82)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_83s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_83", instances, (*A_extension_83)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_84s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_84", instances, (*A_extension_84)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_85s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_85", instances, (*A_extension_85)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_86s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_86", instances, (*A_extension_86)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_87s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_87", instances, (*A_extension_87)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_88s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_88", instances, (*A_extension_88)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_89s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_89", instances, (*A_extension_89)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractAirportHeliportProtectionAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractAirportHeliportProtectionAreaType", instances, (*AbstractAirportHeliportProtectionAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGroundLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGroundLightSystemType", instances, (*AbstractGroundLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMarkingType", instances, (*AbstractMarkingType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AbstractSurfaceContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurfaceContaminationType", instances, (*AbstractSurfaceContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractUsageConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractUsageConditionType", instances, (*AbstractUsageConditionType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AircraftStandContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandContaminationPropertyType", instances, (*AircraftStandContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandContaminationType", instances, (*AircraftStandContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandPropertyType", instances, (*AircraftStandPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandTimeSlicePropertyType", instances, (*AircraftStandTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandTimeSliceType", instances, (*AircraftStandTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandType", instances, (*AircraftStandType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportAvailabilityPropertyType", instances, (*AirportHeliportAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportAvailabilityType", instances, (*AirportHeliportAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationTimeSlicePropertyType", instances, (*AirportHeliportCollocationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationTimeSliceType", instances, (*AirportHeliportCollocationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationType", instances, (*AirportHeliportCollocationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportContaminationPropertyType", instances, (*AirportHeliportContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportContaminationType", instances, (*AirportHeliportContaminationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AirportHeliportProtectionAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportProtectionAreaPropertyType", instances, (*AirportHeliportProtectionAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportResponsibilityOrganisationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportResponsibilityOrganisationPropertyType", instances, (*AirportHeliportResponsibilityOrganisationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportResponsibilityOrganisationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportResponsibilityOrganisationType", instances, (*AirportHeliportResponsibilityOrganisationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportTimeSlicePropertyType", instances, (*AirportHeliportTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportTimeSliceType", instances, (*AirportHeliportTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportType", instances, (*AirportHeliportType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportUsagePropertyType", instances, (*AirportHeliportUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportUsageType", instances, (*AirportHeliportUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotTimeSlicePropertyType", instances, (*AirportHotSpotTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotTimeSliceType", instances, (*AirportHotSpotTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotType", instances, (*AirportHotSpotType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingTimeSlicePropertyType", instances, (*AirportProtectionAreaMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingTimeSliceType", instances, (*AirportProtectionAreaMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingType", instances, (*AirportProtectionAreaMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServicePropertyType", instances, (*AirportSuppliesServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.AltimeterSourceStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceStatusPropertyType", instances, (*AltimeterSourceStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceStatusType", instances, (*AltimeterSourceStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceTimeSlicePropertyType", instances, (*AltimeterSourceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceTimeSliceType", instances, (*AltimeterSourceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceType", instances, (*AltimeterSourceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemTimeSlicePropertyType", instances, (*ApproachLightingSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemTimeSliceType", instances, (*ApproachLightingSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemType", instances, (*ApproachLightingSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaAvailabilityPropertyType", instances, (*ApronAreaAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaAvailabilityType", instances, (*ApronAreaAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaUsagePropertyType", instances, (*ApronAreaUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaUsageType", instances, (*ApronAreaUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronContaminationPropertyType", instances, (*ApronContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronContaminationType", instances, (*ApronContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementPropertyType", instances, (*ApronElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementTimeSlicePropertyType", instances, (*ApronElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementTimeSliceType", instances, (*ApronElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementType", instances, (*ApronElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemTimeSlicePropertyType", instances, (*ApronLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemTimeSliceType", instances, (*ApronLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemType", instances, (*ApronLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingTimeSlicePropertyType", instances, (*ApronMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingTimeSliceType", instances, (*ApronMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingType", instances, (*ApronMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronPropertyType", instances, (*ApronPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronTimeSlicePropertyType", instances, (*ApronTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronTimeSliceType", instances, (*ApronTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronType", instances, (*ApronType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearTimeSlicePropertyType", instances, (*ArrestingGearTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearTimeSliceType", instances, (*ArrestingGearTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearType", instances, (*ArrestingGearType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CityPropertyType", instances, (*CityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CityType", instances, (*CityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConditionCombinationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConditionCombinationPropertyType", instances, (*ConditionCombinationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConditionCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConditionCombinationType", instances, (*ConditionCombinationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveType", instances, (*CurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingTimeSlicePropertyType", instances, (*DeicingAreaMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingTimeSliceType", instances, (*DeicingAreaMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingType", instances, (*DeicingAreaMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaPropertyType", instances, (*DeicingAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaTimeSlicePropertyType", instances, (*DeicingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaTimeSliceType", instances, (*DeicingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaType", instances, (*DeicingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedCurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedCurvePropertyType", instances, (*ElevatedCurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ElevatedSurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedSurfacePropertyType", instances, (*ElevatedSurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.FloatingDockSitePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSitePropertyType", instances, (*FloatingDockSitePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteTimeSlicePropertyType", instances, (*FloatingDockSiteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteTimeSliceType", instances, (*FloatingDockSiteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteType", instances, (*FloatingDockSiteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightingAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightingAvailabilityPropertyType", instances, (*GroundLightingAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightingAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightingAvailabilityType", instances, (*GroundLightingAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemTimeSlicePropertyType", instances, (*GuidanceLineLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemTimeSliceType", instances, (*GuidanceLineLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemType", instances, (*GuidanceLineLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingTimeSlicePropertyType", instances, (*GuidanceLineMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingTimeSliceType", instances, (*GuidanceLineMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingType", instances, (*GuidanceLineMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLinePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLinePropertyType", instances, (*GuidanceLinePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineTimeSlicePropertyType", instances, (*GuidanceLineTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineTimeSliceType", instances, (*GuidanceLineTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineType", instances, (*GuidanceLineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementPropertyType", instances, (*LightElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementStatusPropertyType", instances, (*LightElementStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementStatusType", instances, (*LightElementStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementType", instances, (*LightElementType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ManoeuvringAreaAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaAvailabilityPropertyType", instances, (*ManoeuvringAreaAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaAvailabilityType", instances, (*ManoeuvringAreaAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaUsagePropertyType", instances, (*ManoeuvringAreaUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaUsageType", instances, (*ManoeuvringAreaUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyTimeSlicePropertyType", instances, (*MarkingBuoyTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyTimeSliceType", instances, (*MarkingBuoyTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyType", instances, (*MarkingBuoyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingElementPropertyType", instances, (*MarkingElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingElementType", instances, (*MarkingElementType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.NavaidEquipmentDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentDistancePropertyType", instances, (*NavaidEquipmentDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentDistanceType", instances, (*NavaidEquipmentDistanceType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.NonMovementAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaTimeSlicePropertyType", instances, (*NonMovementAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaTimeSliceType", instances, (*NonMovementAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaType", instances, (*NonMovementAreaType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.PassengerLoadingBridgeTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeTimeSlicePropertyType", instances, (*PassengerLoadingBridgeTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgeTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeTimeSliceType", instances, (*PassengerLoadingBridgeTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeType", instances, (*PassengerLoadingBridgeType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RidgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RidgePropertyType", instances, (*RidgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RidgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RidgeType", instances, (*RidgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadTimeSlicePropertyType", instances, (*RoadTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadTimeSliceType", instances, (*RoadTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadType", instances, (*RoadType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadTimeSlicePropertyType", instances, (*RunwayBlastPadTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadTimeSliceType", instances, (*RunwayBlastPadTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadType", instances, (*RunwayBlastPadType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayCentrelinePointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointTimeSlicePropertyType", instances, (*RunwayCentrelinePointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointTimeSliceType", instances, (*RunwayCentrelinePointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointType", instances, (*RunwayCentrelinePointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayContaminationPropertyType", instances, (*RunwayContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayContaminationType", instances, (*RunwayContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistancePropertyType", instances, (*RunwayDeclaredDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceType", instances, (*RunwayDeclaredDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceValuePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceValuePropertyType", instances, (*RunwayDeclaredDistanceValuePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceValueTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceValueType", instances, (*RunwayDeclaredDistanceValueType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemTimeSlicePropertyType", instances, (*RunwayDirectionLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemTimeSliceType", instances, (*RunwayDirectionLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemType", instances, (*RunwayDirectionLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayDirectionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionTimeSlicePropertyType", instances, (*RunwayDirectionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionTimeSliceType", instances, (*RunwayDirectionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionType", instances, (*RunwayDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementPropertyType", instances, (*RunwayElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementTimeSlicePropertyType", instances, (*RunwayElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementTimeSliceType", instances, (*RunwayElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementType", instances, (*RunwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingTimeSlicePropertyType", instances, (*RunwayMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingTimeSliceType", instances, (*RunwayMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingType", instances, (*RunwayMarkingType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.RunwayProtectAreaLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemTimeSlicePropertyType", instances, (*RunwayProtectAreaLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemTimeSliceType", instances, (*RunwayProtectAreaLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemType", instances, (*RunwayProtectAreaLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaPropertyType", instances, (*RunwayProtectAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaTimeSlicePropertyType", instances, (*RunwayProtectAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaTimeSliceType", instances, (*RunwayProtectAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaType", instances, (*RunwayProtectAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwaySectionContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwaySectionContaminationPropertyType", instances, (*RunwaySectionContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwaySectionContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwaySectionContaminationType", instances, (*RunwaySectionContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayTimeSlicePropertyType", instances, (*RunwayTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayTimeSliceType", instances, (*RunwayTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayType", instances, (*RunwayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeTimeSlicePropertyType", instances, (*RunwayVisualRangeTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeTimeSliceType", instances, (*RunwayVisualRangeTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeType", instances, (*RunwayVisualRangeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaPropertyType", instances, (*SeaplaneLandingAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaTimeSlicePropertyType", instances, (*SeaplaneLandingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaTimeSliceType", instances, (*SeaplaneLandingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaType", instances, (*SeaplaneLandingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSitePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSitePropertyType", instances, (*SeaplaneRampSitePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteTimeSlicePropertyType", instances, (*SeaplaneRampSiteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteTimeSliceType", instances, (*SeaplaneRampSiteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteType", instances, (*SeaplaneRampSiteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingTimeSlicePropertyType", instances, (*StandMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingTimeSliceType", instances, (*StandMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingType", instances, (*StandMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceCharacteristicsPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceCharacteristicsPropertyType", instances, (*SurfaceCharacteristicsPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceCharacteristicsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceCharacteristicsType", instances, (*SurfaceCharacteristicsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceContaminationLayerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceContaminationLayerPropertyType", instances, (*SurfaceContaminationLayerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceContaminationLayerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceContaminationLayerType", instances, (*SurfaceContaminationLayerType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.SurveyControlPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointTimeSlicePropertyType", instances, (*SurveyControlPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointTimeSliceType", instances, (*SurveyControlPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointType", instances, (*SurveyControlPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemTimeSlicePropertyType", instances, (*TaxiHoldingPositionLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemTimeSliceType", instances, (*TaxiHoldingPositionLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemType", instances, (*TaxiHoldingPositionLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingTimeSlicePropertyType", instances, (*TaxiHoldingPositionMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingTimeSliceType", instances, (*TaxiHoldingPositionMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingType", instances, (*TaxiHoldingPositionMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionPropertyType", instances, (*TaxiHoldingPositionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionTimeSlicePropertyType", instances, (*TaxiHoldingPositionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionTimeSliceType", instances, (*TaxiHoldingPositionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionType", instances, (*TaxiHoldingPositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayContaminationPropertyType", instances, (*TaxiwayContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayContaminationType", instances, (*TaxiwayContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementPropertyType", instances, (*TaxiwayElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementTimeSlicePropertyType", instances, (*TaxiwayElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementTimeSliceType", instances, (*TaxiwayElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementType", instances, (*TaxiwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemTimeSlicePropertyType", instances, (*TaxiwayLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemTimeSliceType", instances, (*TaxiwayLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemType", instances, (*TaxiwayLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingTimeSlicePropertyType", instances, (*TaxiwayMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingTimeSliceType", instances, (*TaxiwayMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingType", instances, (*TaxiwayMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayPropertyType", instances, (*TaxiwayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayTimeSlicePropertyType", instances, (*TaxiwayTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayTimeSliceType", instances, (*TaxiwayTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayType", instances, (*TaxiwayType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.TouchDownLiftOffContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffContaminationPropertyType", instances, (*TouchDownLiftOffContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffContaminationType", instances, (*TouchDownLiftOffContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemTimeSlicePropertyType", instances, (*TouchDownLiftOffLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemTimeSliceType", instances, (*TouchDownLiftOffLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemType", instances, (*TouchDownLiftOffLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingTimeSlicePropertyType", instances, (*TouchDownLiftOffMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingTimeSliceType", instances, (*TouchDownLiftOffMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingType", instances, (*TouchDownLiftOffMarkingType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.TouchDownLiftOffSafeAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaTimeSlicePropertyType", instances, (*TouchDownLiftOffSafeAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaTimeSliceType", instances, (*TouchDownLiftOffSafeAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaType", instances, (*TouchDownLiftOffSafeAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffTimeSlicePropertyType", instances, (*TouchDownLiftOffTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffTimeSliceType", instances, (*TouchDownLiftOffTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffType", instances, (*TouchDownLiftOffType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorTimeSlicePropertyType", instances, (*VisualGlideSlopeIndicatorTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorTimeSliceType", instances, (*VisualGlideSlopeIndicatorTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorType", instances, (*VisualGlideSlopeIndicatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaTimeSlicePropertyType", instances, (*WorkAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaTimeSliceType", instances, (*WorkAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaType", instances, (*WorkAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkareaActivityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkareaActivityPropertyType", instances, (*WorkareaActivityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkareaActivityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkareaActivityType", instances, (*WorkareaActivityType)(nil).GongGetFieldHeaders(), addIDs)
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
