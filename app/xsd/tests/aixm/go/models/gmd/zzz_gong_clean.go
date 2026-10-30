// generated code - do not edit
package gmd

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
// Clean garbage collect unstaged instances that are referenced by AbstractDS_Aggregate_Type
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractds_aggregate_type.ComposedOf) || modified
	modified = stage.CleanSlice(&abstractds_aggregate_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&abstractds_aggregate_type.Subset) || modified
	modified = stage.CleanSlice(&abstractds_aggregate_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Aggregate_PropertyType
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_OtherAggregate) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_StereoMate) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_Series) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_Platform) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_Sensor) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_ProductionSeries) || modified
	modified = stage.CleanPointer(&ds_aggregate_propertytype.DS_Initiative) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_DataSet_PropertyType
func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_dataset_propertytype.DS_DataSet) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_DataSet_Type
func (ds_dataset_type *DS_DataSet_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_dataset_type.Has) || modified
	modified = stage.CleanSlice(&ds_dataset_type.PartOf) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Initiative_PropertyType
func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_initiative_propertytype.DS_Initiative) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Initiative_Type
func (ds_initiative_type *DS_Initiative_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_initiative_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_initiative_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_initiative_type.Subset) || modified
	modified = stage.CleanSlice(&ds_initiative_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_OtherAggregate_PropertyType
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_otheraggregate_propertytype.DS_OtherAggregate) || modified
	modified = stage.CleanPointer(&ds_otheraggregate_propertytype.DS_StereoMate) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_OtherAggregate_Type
func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_otheraggregate_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_otheraggregate_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_otheraggregate_type.Subset) || modified
	modified = stage.CleanSlice(&ds_otheraggregate_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Platform_PropertyType
func (ds_platform_propertytype *DS_Platform_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_platform_propertytype.DS_Platform) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Platform_Type
func (ds_platform_type *DS_Platform_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_platform_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_platform_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_platform_type.Subset) || modified
	modified = stage.CleanSlice(&ds_platform_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_ProductionSeries_PropertyType
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_productionseries_propertytype.DS_ProductionSeries) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_ProductionSeries_Type
func (ds_productionseries_type *DS_ProductionSeries_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_productionseries_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_productionseries_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_productionseries_type.Subset) || modified
	modified = stage.CleanSlice(&ds_productionseries_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Sensor_PropertyType
func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_sensor_propertytype.DS_Sensor) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Sensor_Type
func (ds_sensor_type *DS_Sensor_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_sensor_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_sensor_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_sensor_type.Subset) || modified
	modified = stage.CleanSlice(&ds_sensor_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Series_PropertyType
func (ds_series_propertytype *DS_Series_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_series_propertytype.DS_Series) || modified
	modified = stage.CleanPointer(&ds_series_propertytype.DS_Platform) || modified
	modified = stage.CleanPointer(&ds_series_propertytype.DS_Sensor) || modified
	modified = stage.CleanPointer(&ds_series_propertytype.DS_ProductionSeries) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Series_Type
func (ds_series_type *DS_Series_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_series_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_series_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_series_type.Subset) || modified
	modified = stage.CleanSlice(&ds_series_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_StereoMate_PropertyType
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_stereomate_propertytype.DS_StereoMate) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_StereoMate_Type
func (ds_stereomate_type *DS_StereoMate_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ds_stereomate_type.ComposedOf) || modified
	modified = stage.CleanSlice(&ds_stereomate_type.SeriesMetadata) || modified
	modified = stage.CleanSlice(&ds_stereomate_type.Subset) || modified
	modified = stage.CleanSlice(&ds_stereomate_type.Superset) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Metadata_PropertyType
func (md_metadata_propertytype *MD_Metadata_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_metadata_propertytype.MD_Metadata) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Metadata_Type
func (md_metadata_type *MD_Metadata_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_metadata_type.Series) || modified
	modified = stage.CleanSlice(&md_metadata_type.Describes) || modified
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
