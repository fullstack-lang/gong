// generated code - do not edit
package gmd

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDS_Aggregate_TypeCreateCallback != nil {
		stage.OnAfterAbstractDS_Aggregate_TypeCreateCallback.OnAfterCreate(stage, abstractds_aggregate_type)
	}
}

func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDS_Aggregate_TypeUpdateCallback != nil {
		var frontAbstractDS_Aggregate_Type *AbstractDS_Aggregate_Type
		if front != nil {
			frontAbstractDS_Aggregate_Type, _ = front.(*AbstractDS_Aggregate_Type)
		}
		stage.OnAfterAbstractDS_Aggregate_TypeUpdateCallback.OnAfterUpdate(stage, abstractds_aggregate_type, frontAbstractDS_Aggregate_Type)
	}
}

func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDS_Aggregate_TypeDeleteCallback != nil {
		var frontAbstractDS_Aggregate_Type *AbstractDS_Aggregate_Type
		if front != nil {
			frontAbstractDS_Aggregate_Type, _ = front.(*AbstractDS_Aggregate_Type)
		}
		stage.OnAfterAbstractDS_Aggregate_TypeDeleteCallback.OnAfterDelete(stage, abstractds_aggregate_type, frontAbstractDS_Aggregate_Type)
	}
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Aggregate_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Aggregate_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_aggregate_propertytype)
	}
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Aggregate_PropertyTypeUpdateCallback != nil {
		var frontDS_Aggregate_PropertyType *DS_Aggregate_PropertyType
		if front != nil {
			frontDS_Aggregate_PropertyType, _ = front.(*DS_Aggregate_PropertyType)
		}
		stage.OnAfterDS_Aggregate_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_aggregate_propertytype, frontDS_Aggregate_PropertyType)
	}
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Aggregate_PropertyTypeDeleteCallback != nil {
		var frontDS_Aggregate_PropertyType *DS_Aggregate_PropertyType
		if front != nil {
			frontDS_Aggregate_PropertyType, _ = front.(*DS_Aggregate_PropertyType)
		}
		stage.OnAfterDS_Aggregate_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_aggregate_propertytype, frontDS_Aggregate_PropertyType)
	}
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_DataSet_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_DataSet_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_dataset_propertytype)
	}
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_DataSet_PropertyTypeUpdateCallback != nil {
		var frontDS_DataSet_PropertyType *DS_DataSet_PropertyType
		if front != nil {
			frontDS_DataSet_PropertyType, _ = front.(*DS_DataSet_PropertyType)
		}
		stage.OnAfterDS_DataSet_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_dataset_propertytype, frontDS_DataSet_PropertyType)
	}
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_DataSet_PropertyTypeDeleteCallback != nil {
		var frontDS_DataSet_PropertyType *DS_DataSet_PropertyType
		if front != nil {
			frontDS_DataSet_PropertyType, _ = front.(*DS_DataSet_PropertyType)
		}
		stage.OnAfterDS_DataSet_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_dataset_propertytype, frontDS_DataSet_PropertyType)
	}
}

func (ds_dataset_type *DS_DataSet_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_DataSet_TypeCreateCallback != nil {
		stage.OnAfterDS_DataSet_TypeCreateCallback.OnAfterCreate(stage, ds_dataset_type)
	}
}

func (ds_dataset_type *DS_DataSet_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_DataSet_TypeUpdateCallback != nil {
		var frontDS_DataSet_Type *DS_DataSet_Type
		if front != nil {
			frontDS_DataSet_Type, _ = front.(*DS_DataSet_Type)
		}
		stage.OnAfterDS_DataSet_TypeUpdateCallback.OnAfterUpdate(stage, ds_dataset_type, frontDS_DataSet_Type)
	}
}

func (ds_dataset_type *DS_DataSet_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_DataSet_TypeDeleteCallback != nil {
		var frontDS_DataSet_Type *DS_DataSet_Type
		if front != nil {
			frontDS_DataSet_Type, _ = front.(*DS_DataSet_Type)
		}
		stage.OnAfterDS_DataSet_TypeDeleteCallback.OnAfterDelete(stage, ds_dataset_type, frontDS_DataSet_Type)
	}
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Initiative_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Initiative_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_initiative_propertytype)
	}
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Initiative_PropertyTypeUpdateCallback != nil {
		var frontDS_Initiative_PropertyType *DS_Initiative_PropertyType
		if front != nil {
			frontDS_Initiative_PropertyType, _ = front.(*DS_Initiative_PropertyType)
		}
		stage.OnAfterDS_Initiative_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_initiative_propertytype, frontDS_Initiative_PropertyType)
	}
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Initiative_PropertyTypeDeleteCallback != nil {
		var frontDS_Initiative_PropertyType *DS_Initiative_PropertyType
		if front != nil {
			frontDS_Initiative_PropertyType, _ = front.(*DS_Initiative_PropertyType)
		}
		stage.OnAfterDS_Initiative_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_initiative_propertytype, frontDS_Initiative_PropertyType)
	}
}

func (ds_initiative_type *DS_Initiative_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Initiative_TypeCreateCallback != nil {
		stage.OnAfterDS_Initiative_TypeCreateCallback.OnAfterCreate(stage, ds_initiative_type)
	}
}

func (ds_initiative_type *DS_Initiative_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Initiative_TypeUpdateCallback != nil {
		var frontDS_Initiative_Type *DS_Initiative_Type
		if front != nil {
			frontDS_Initiative_Type, _ = front.(*DS_Initiative_Type)
		}
		stage.OnAfterDS_Initiative_TypeUpdateCallback.OnAfterUpdate(stage, ds_initiative_type, frontDS_Initiative_Type)
	}
}

func (ds_initiative_type *DS_Initiative_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Initiative_TypeDeleteCallback != nil {
		var frontDS_Initiative_Type *DS_Initiative_Type
		if front != nil {
			frontDS_Initiative_Type, _ = front.(*DS_Initiative_Type)
		}
		stage.OnAfterDS_Initiative_TypeDeleteCallback.OnAfterDelete(stage, ds_initiative_type, frontDS_Initiative_Type)
	}
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_OtherAggregate_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_OtherAggregate_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_otheraggregate_propertytype)
	}
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_OtherAggregate_PropertyTypeUpdateCallback != nil {
		var frontDS_OtherAggregate_PropertyType *DS_OtherAggregate_PropertyType
		if front != nil {
			frontDS_OtherAggregate_PropertyType, _ = front.(*DS_OtherAggregate_PropertyType)
		}
		stage.OnAfterDS_OtherAggregate_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_otheraggregate_propertytype, frontDS_OtherAggregate_PropertyType)
	}
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_OtherAggregate_PropertyTypeDeleteCallback != nil {
		var frontDS_OtherAggregate_PropertyType *DS_OtherAggregate_PropertyType
		if front != nil {
			frontDS_OtherAggregate_PropertyType, _ = front.(*DS_OtherAggregate_PropertyType)
		}
		stage.OnAfterDS_OtherAggregate_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_otheraggregate_propertytype, frontDS_OtherAggregate_PropertyType)
	}
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_OtherAggregate_TypeCreateCallback != nil {
		stage.OnAfterDS_OtherAggregate_TypeCreateCallback.OnAfterCreate(stage, ds_otheraggregate_type)
	}
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_OtherAggregate_TypeUpdateCallback != nil {
		var frontDS_OtherAggregate_Type *DS_OtherAggregate_Type
		if front != nil {
			frontDS_OtherAggregate_Type, _ = front.(*DS_OtherAggregate_Type)
		}
		stage.OnAfterDS_OtherAggregate_TypeUpdateCallback.OnAfterUpdate(stage, ds_otheraggregate_type, frontDS_OtherAggregate_Type)
	}
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_OtherAggregate_TypeDeleteCallback != nil {
		var frontDS_OtherAggregate_Type *DS_OtherAggregate_Type
		if front != nil {
			frontDS_OtherAggregate_Type, _ = front.(*DS_OtherAggregate_Type)
		}
		stage.OnAfterDS_OtherAggregate_TypeDeleteCallback.OnAfterDelete(stage, ds_otheraggregate_type, frontDS_OtherAggregate_Type)
	}
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Platform_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Platform_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_platform_propertytype)
	}
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Platform_PropertyTypeUpdateCallback != nil {
		var frontDS_Platform_PropertyType *DS_Platform_PropertyType
		if front != nil {
			frontDS_Platform_PropertyType, _ = front.(*DS_Platform_PropertyType)
		}
		stage.OnAfterDS_Platform_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_platform_propertytype, frontDS_Platform_PropertyType)
	}
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Platform_PropertyTypeDeleteCallback != nil {
		var frontDS_Platform_PropertyType *DS_Platform_PropertyType
		if front != nil {
			frontDS_Platform_PropertyType, _ = front.(*DS_Platform_PropertyType)
		}
		stage.OnAfterDS_Platform_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_platform_propertytype, frontDS_Platform_PropertyType)
	}
}

func (ds_platform_type *DS_Platform_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Platform_TypeCreateCallback != nil {
		stage.OnAfterDS_Platform_TypeCreateCallback.OnAfterCreate(stage, ds_platform_type)
	}
}

func (ds_platform_type *DS_Platform_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Platform_TypeUpdateCallback != nil {
		var frontDS_Platform_Type *DS_Platform_Type
		if front != nil {
			frontDS_Platform_Type, _ = front.(*DS_Platform_Type)
		}
		stage.OnAfterDS_Platform_TypeUpdateCallback.OnAfterUpdate(stage, ds_platform_type, frontDS_Platform_Type)
	}
}

func (ds_platform_type *DS_Platform_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Platform_TypeDeleteCallback != nil {
		var frontDS_Platform_Type *DS_Platform_Type
		if front != nil {
			frontDS_Platform_Type, _ = front.(*DS_Platform_Type)
		}
		stage.OnAfterDS_Platform_TypeDeleteCallback.OnAfterDelete(stage, ds_platform_type, frontDS_Platform_Type)
	}
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_ProductionSeries_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_ProductionSeries_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_productionseries_propertytype)
	}
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_ProductionSeries_PropertyTypeUpdateCallback != nil {
		var frontDS_ProductionSeries_PropertyType *DS_ProductionSeries_PropertyType
		if front != nil {
			frontDS_ProductionSeries_PropertyType, _ = front.(*DS_ProductionSeries_PropertyType)
		}
		stage.OnAfterDS_ProductionSeries_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_productionseries_propertytype, frontDS_ProductionSeries_PropertyType)
	}
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_ProductionSeries_PropertyTypeDeleteCallback != nil {
		var frontDS_ProductionSeries_PropertyType *DS_ProductionSeries_PropertyType
		if front != nil {
			frontDS_ProductionSeries_PropertyType, _ = front.(*DS_ProductionSeries_PropertyType)
		}
		stage.OnAfterDS_ProductionSeries_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_productionseries_propertytype, frontDS_ProductionSeries_PropertyType)
	}
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_ProductionSeries_TypeCreateCallback != nil {
		stage.OnAfterDS_ProductionSeries_TypeCreateCallback.OnAfterCreate(stage, ds_productionseries_type)
	}
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_ProductionSeries_TypeUpdateCallback != nil {
		var frontDS_ProductionSeries_Type *DS_ProductionSeries_Type
		if front != nil {
			frontDS_ProductionSeries_Type, _ = front.(*DS_ProductionSeries_Type)
		}
		stage.OnAfterDS_ProductionSeries_TypeUpdateCallback.OnAfterUpdate(stage, ds_productionseries_type, frontDS_ProductionSeries_Type)
	}
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_ProductionSeries_TypeDeleteCallback != nil {
		var frontDS_ProductionSeries_Type *DS_ProductionSeries_Type
		if front != nil {
			frontDS_ProductionSeries_Type, _ = front.(*DS_ProductionSeries_Type)
		}
		stage.OnAfterDS_ProductionSeries_TypeDeleteCallback.OnAfterDelete(stage, ds_productionseries_type, frontDS_ProductionSeries_Type)
	}
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Sensor_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Sensor_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_sensor_propertytype)
	}
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Sensor_PropertyTypeUpdateCallback != nil {
		var frontDS_Sensor_PropertyType *DS_Sensor_PropertyType
		if front != nil {
			frontDS_Sensor_PropertyType, _ = front.(*DS_Sensor_PropertyType)
		}
		stage.OnAfterDS_Sensor_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_sensor_propertytype, frontDS_Sensor_PropertyType)
	}
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Sensor_PropertyTypeDeleteCallback != nil {
		var frontDS_Sensor_PropertyType *DS_Sensor_PropertyType
		if front != nil {
			frontDS_Sensor_PropertyType, _ = front.(*DS_Sensor_PropertyType)
		}
		stage.OnAfterDS_Sensor_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_sensor_propertytype, frontDS_Sensor_PropertyType)
	}
}

func (ds_sensor_type *DS_Sensor_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Sensor_TypeCreateCallback != nil {
		stage.OnAfterDS_Sensor_TypeCreateCallback.OnAfterCreate(stage, ds_sensor_type)
	}
}

func (ds_sensor_type *DS_Sensor_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Sensor_TypeUpdateCallback != nil {
		var frontDS_Sensor_Type *DS_Sensor_Type
		if front != nil {
			frontDS_Sensor_Type, _ = front.(*DS_Sensor_Type)
		}
		stage.OnAfterDS_Sensor_TypeUpdateCallback.OnAfterUpdate(stage, ds_sensor_type, frontDS_Sensor_Type)
	}
}

func (ds_sensor_type *DS_Sensor_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Sensor_TypeDeleteCallback != nil {
		var frontDS_Sensor_Type *DS_Sensor_Type
		if front != nil {
			frontDS_Sensor_Type, _ = front.(*DS_Sensor_Type)
		}
		stage.OnAfterDS_Sensor_TypeDeleteCallback.OnAfterDelete(stage, ds_sensor_type, frontDS_Sensor_Type)
	}
}

func (ds_series_propertytype *DS_Series_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Series_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Series_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_series_propertytype)
	}
}

func (ds_series_propertytype *DS_Series_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Series_PropertyTypeUpdateCallback != nil {
		var frontDS_Series_PropertyType *DS_Series_PropertyType
		if front != nil {
			frontDS_Series_PropertyType, _ = front.(*DS_Series_PropertyType)
		}
		stage.OnAfterDS_Series_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_series_propertytype, frontDS_Series_PropertyType)
	}
}

func (ds_series_propertytype *DS_Series_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Series_PropertyTypeDeleteCallback != nil {
		var frontDS_Series_PropertyType *DS_Series_PropertyType
		if front != nil {
			frontDS_Series_PropertyType, _ = front.(*DS_Series_PropertyType)
		}
		stage.OnAfterDS_Series_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_series_propertytype, frontDS_Series_PropertyType)
	}
}

func (ds_series_type *DS_Series_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Series_TypeCreateCallback != nil {
		stage.OnAfterDS_Series_TypeCreateCallback.OnAfterCreate(stage, ds_series_type)
	}
}

func (ds_series_type *DS_Series_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Series_TypeUpdateCallback != nil {
		var frontDS_Series_Type *DS_Series_Type
		if front != nil {
			frontDS_Series_Type, _ = front.(*DS_Series_Type)
		}
		stage.OnAfterDS_Series_TypeUpdateCallback.OnAfterUpdate(stage, ds_series_type, frontDS_Series_Type)
	}
}

func (ds_series_type *DS_Series_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Series_TypeDeleteCallback != nil {
		var frontDS_Series_Type *DS_Series_Type
		if front != nil {
			frontDS_Series_Type, _ = front.(*DS_Series_Type)
		}
		stage.OnAfterDS_Series_TypeDeleteCallback.OnAfterDelete(stage, ds_series_type, frontDS_Series_Type)
	}
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_StereoMate_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_StereoMate_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_stereomate_propertytype)
	}
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_StereoMate_PropertyTypeUpdateCallback != nil {
		var frontDS_StereoMate_PropertyType *DS_StereoMate_PropertyType
		if front != nil {
			frontDS_StereoMate_PropertyType, _ = front.(*DS_StereoMate_PropertyType)
		}
		stage.OnAfterDS_StereoMate_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_stereomate_propertytype, frontDS_StereoMate_PropertyType)
	}
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_StereoMate_PropertyTypeDeleteCallback != nil {
		var frontDS_StereoMate_PropertyType *DS_StereoMate_PropertyType
		if front != nil {
			frontDS_StereoMate_PropertyType, _ = front.(*DS_StereoMate_PropertyType)
		}
		stage.OnAfterDS_StereoMate_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_stereomate_propertytype, frontDS_StereoMate_PropertyType)
	}
}

func (ds_stereomate_type *DS_StereoMate_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_StereoMate_TypeCreateCallback != nil {
		stage.OnAfterDS_StereoMate_TypeCreateCallback.OnAfterCreate(stage, ds_stereomate_type)
	}
}

func (ds_stereomate_type *DS_StereoMate_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_StereoMate_TypeUpdateCallback != nil {
		var frontDS_StereoMate_Type *DS_StereoMate_Type
		if front != nil {
			frontDS_StereoMate_Type, _ = front.(*DS_StereoMate_Type)
		}
		stage.OnAfterDS_StereoMate_TypeUpdateCallback.OnAfterUpdate(stage, ds_stereomate_type, frontDS_StereoMate_Type)
	}
}

func (ds_stereomate_type *DS_StereoMate_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_StereoMate_TypeDeleteCallback != nil {
		var frontDS_StereoMate_Type *DS_StereoMate_Type
		if front != nil {
			frontDS_StereoMate_Type, _ = front.(*DS_StereoMate_Type)
		}
		stage.OnAfterDS_StereoMate_TypeDeleteCallback.OnAfterDelete(stage, ds_stereomate_type, frontDS_StereoMate_Type)
	}
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Metadata_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Metadata_PropertyTypeCreateCallback.OnAfterCreate(stage, md_metadata_propertytype)
	}
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Metadata_PropertyTypeUpdateCallback != nil {
		var frontMD_Metadata_PropertyType *MD_Metadata_PropertyType
		if front != nil {
			frontMD_Metadata_PropertyType, _ = front.(*MD_Metadata_PropertyType)
		}
		stage.OnAfterMD_Metadata_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_metadata_propertytype, frontMD_Metadata_PropertyType)
	}
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Metadata_PropertyTypeDeleteCallback != nil {
		var frontMD_Metadata_PropertyType *MD_Metadata_PropertyType
		if front != nil {
			frontMD_Metadata_PropertyType, _ = front.(*MD_Metadata_PropertyType)
		}
		stage.OnAfterMD_Metadata_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_metadata_propertytype, frontMD_Metadata_PropertyType)
	}
}

func (md_metadata_type *MD_Metadata_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Metadata_TypeCreateCallback != nil {
		stage.OnAfterMD_Metadata_TypeCreateCallback.OnAfterCreate(stage, md_metadata_type)
	}
}

func (md_metadata_type *MD_Metadata_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Metadata_TypeUpdateCallback != nil {
		var frontMD_Metadata_Type *MD_Metadata_Type
		if front != nil {
			frontMD_Metadata_Type, _ = front.(*MD_Metadata_Type)
		}
		stage.OnAfterMD_Metadata_TypeUpdateCallback.OnAfterUpdate(stage, md_metadata_type, frontMD_Metadata_Type)
	}
}

func (md_metadata_type *MD_Metadata_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Metadata_TypeDeleteCallback != nil {
		var frontMD_Metadata_Type *MD_Metadata_Type
		if front != nil {
			frontMD_Metadata_Type, _ = front.(*MD_Metadata_Type)
		}
		stage.OnAfterMD_Metadata_TypeDeleteCallback.OnAfterDelete(stage, md_metadata_type, frontMD_Metadata_Type)
	}
}

