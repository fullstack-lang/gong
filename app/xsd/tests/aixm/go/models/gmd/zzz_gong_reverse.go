// generated code - do not edit
package gmd

// insertion point
func (inst *AbstractDS_Aggregate_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Aggregate_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractDS_Aggregate_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _abstractds_aggregate_type, ok := stage.AbstractDS_Aggregate_Type_Subset_reverseMap[inst]; ok {
				res = _abstractds_aggregate_type.Name
			}
		case "Superset":
			if _abstractds_aggregate_type, ok := stage.AbstractDS_Aggregate_Type_Superset_reverseMap[inst]; ok {
				res = _abstractds_aggregate_type.Name
			}
		}
	case "DS_DataSet_Type":
		switch reverseField.Fieldname {
		case "PartOf":
			if _ds_dataset_type, ok := stage.DS_DataSet_Type_PartOf_reverseMap[inst]; ok {
				res = _ds_dataset_type.Name
			}
		}
	case "DS_Initiative_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_initiative_type, ok := stage.DS_Initiative_Type_Subset_reverseMap[inst]; ok {
				res = _ds_initiative_type.Name
			}
		case "Superset":
			if _ds_initiative_type, ok := stage.DS_Initiative_Type_Superset_reverseMap[inst]; ok {
				res = _ds_initiative_type.Name
			}
		}
	case "DS_OtherAggregate_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_otheraggregate_type, ok := stage.DS_OtherAggregate_Type_Subset_reverseMap[inst]; ok {
				res = _ds_otheraggregate_type.Name
			}
		case "Superset":
			if _ds_otheraggregate_type, ok := stage.DS_OtherAggregate_Type_Superset_reverseMap[inst]; ok {
				res = _ds_otheraggregate_type.Name
			}
		}
	case "DS_Platform_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_platform_type, ok := stage.DS_Platform_Type_Subset_reverseMap[inst]; ok {
				res = _ds_platform_type.Name
			}
		case "Superset":
			if _ds_platform_type, ok := stage.DS_Platform_Type_Superset_reverseMap[inst]; ok {
				res = _ds_platform_type.Name
			}
		}
	case "DS_ProductionSeries_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_productionseries_type, ok := stage.DS_ProductionSeries_Type_Subset_reverseMap[inst]; ok {
				res = _ds_productionseries_type.Name
			}
		case "Superset":
			if _ds_productionseries_type, ok := stage.DS_ProductionSeries_Type_Superset_reverseMap[inst]; ok {
				res = _ds_productionseries_type.Name
			}
		}
	case "DS_Sensor_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_sensor_type, ok := stage.DS_Sensor_Type_Subset_reverseMap[inst]; ok {
				res = _ds_sensor_type.Name
			}
		case "Superset":
			if _ds_sensor_type, ok := stage.DS_Sensor_Type_Superset_reverseMap[inst]; ok {
				res = _ds_sensor_type.Name
			}
		}
	case "DS_Series_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_series_type, ok := stage.DS_Series_Type_Subset_reverseMap[inst]; ok {
				res = _ds_series_type.Name
			}
		case "Superset":
			if _ds_series_type, ok := stage.DS_Series_Type_Superset_reverseMap[inst]; ok {
				res = _ds_series_type.Name
			}
		}
	case "DS_StereoMate_Type":
		switch reverseField.Fieldname {
		case "Subset":
			if _ds_stereomate_type, ok := stage.DS_StereoMate_Type_Subset_reverseMap[inst]; ok {
				res = _ds_stereomate_type.Name
			}
		case "Superset":
			if _ds_stereomate_type, ok := stage.DS_StereoMate_Type_Superset_reverseMap[inst]; ok {
				res = _ds_stereomate_type.Name
			}
		}
	case "MD_Metadata_Type":
		switch reverseField.Fieldname {
		case "Series":
			if _md_metadata_type, ok := stage.MD_Metadata_Type_Series_reverseMap[inst]; ok {
				res = _md_metadata_type.Name
			}
		}
	}
	return
}

func (inst *DS_DataSet_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractDS_Aggregate_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _abstractds_aggregate_type, ok := stage.AbstractDS_Aggregate_Type_ComposedOf_reverseMap[inst]; ok {
				res = _abstractds_aggregate_type.Name
			}
		}
	case "DS_Initiative_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_initiative_type, ok := stage.DS_Initiative_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_initiative_type.Name
			}
		}
	case "DS_OtherAggregate_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_otheraggregate_type, ok := stage.DS_OtherAggregate_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_otheraggregate_type.Name
			}
		}
	case "DS_Platform_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_platform_type, ok := stage.DS_Platform_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_platform_type.Name
			}
		}
	case "DS_ProductionSeries_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_productionseries_type, ok := stage.DS_ProductionSeries_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_productionseries_type.Name
			}
		}
	case "DS_Sensor_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_sensor_type, ok := stage.DS_Sensor_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_sensor_type.Name
			}
		}
	case "DS_Series_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_series_type, ok := stage.DS_Series_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_series_type.Name
			}
		}
	case "DS_StereoMate_Type":
		switch reverseField.Fieldname {
		case "ComposedOf":
			if _ds_stereomate_type, ok := stage.DS_StereoMate_Type_ComposedOf_reverseMap[inst]; ok {
				res = _ds_stereomate_type.Name
			}
		}
	case "MD_Metadata_Type":
		switch reverseField.Fieldname {
		case "Describes":
			if _md_metadata_type, ok := stage.MD_Metadata_Type_Describes_reverseMap[inst]; ok {
				res = _md_metadata_type.Name
			}
		}
	}
	return
}

func (inst *DS_DataSet_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Initiative_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Initiative_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_OtherAggregate_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_OtherAggregate_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Platform_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Platform_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_ProductionSeries_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_ProductionSeries_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Sensor_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Sensor_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Series_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Series_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_StereoMate_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_StereoMate_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Metadata_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractDS_Aggregate_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _abstractds_aggregate_type, ok := stage.AbstractDS_Aggregate_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _abstractds_aggregate_type.Name
			}
		}
	case "DS_DataSet_Type":
		switch reverseField.Fieldname {
		case "Has":
			if _ds_dataset_type, ok := stage.DS_DataSet_Type_Has_reverseMap[inst]; ok {
				res = _ds_dataset_type.Name
			}
		}
	case "DS_Initiative_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_initiative_type, ok := stage.DS_Initiative_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_initiative_type.Name
			}
		}
	case "DS_OtherAggregate_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_otheraggregate_type, ok := stage.DS_OtherAggregate_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_otheraggregate_type.Name
			}
		}
	case "DS_Platform_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_platform_type, ok := stage.DS_Platform_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_platform_type.Name
			}
		}
	case "DS_ProductionSeries_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_productionseries_type, ok := stage.DS_ProductionSeries_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_productionseries_type.Name
			}
		}
	case "DS_Sensor_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_sensor_type, ok := stage.DS_Sensor_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_sensor_type.Name
			}
		}
	case "DS_Series_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_series_type, ok := stage.DS_Series_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_series_type.Name
			}
		}
	case "DS_StereoMate_Type":
		switch reverseField.Fieldname {
		case "SeriesMetadata":
			if _ds_stereomate_type, ok := stage.DS_StereoMate_Type_SeriesMetadata_reverseMap[inst]; ok {
				res = _ds_stereomate_type.Name
			}
		}
	}
	return
}

func (inst *MD_Metadata_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
