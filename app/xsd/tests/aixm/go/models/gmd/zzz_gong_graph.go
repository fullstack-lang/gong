// generated code - do not edit
package gmd

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractDS_Aggregate_Types[abstractds_aggregate_type]
	return ok
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Aggregate_PropertyTypes[ds_aggregate_propertytype]
	return ok
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_DataSet_PropertyTypes[ds_dataset_propertytype]
	return ok
}

func (ds_dataset_type *DS_DataSet_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_DataSet_Types[ds_dataset_type]
	return ok
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Initiative_PropertyTypes[ds_initiative_propertytype]
	return ok
}

func (ds_initiative_type *DS_Initiative_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Initiative_Types[ds_initiative_type]
	return ok
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_OtherAggregate_PropertyTypes[ds_otheraggregate_propertytype]
	return ok
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_OtherAggregate_Types[ds_otheraggregate_type]
	return ok
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Platform_PropertyTypes[ds_platform_propertytype]
	return ok
}

func (ds_platform_type *DS_Platform_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Platform_Types[ds_platform_type]
	return ok
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_ProductionSeries_PropertyTypes[ds_productionseries_propertytype]
	return ok
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_ProductionSeries_Types[ds_productionseries_type]
	return ok
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Sensor_PropertyTypes[ds_sensor_propertytype]
	return ok
}

func (ds_sensor_type *DS_Sensor_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Sensor_Types[ds_sensor_type]
	return ok
}

func (ds_series_propertytype *DS_Series_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Series_PropertyTypes[ds_series_propertytype]
	return ok
}

func (ds_series_type *DS_Series_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_Series_Types[ds_series_type]
	return ok
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_StereoMate_PropertyTypes[ds_stereomate_propertytype]
	return ok
}

func (ds_stereomate_type *DS_StereoMate_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DS_StereoMate_Types[ds_stereomate_type]
	return ok
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Metadata_PropertyTypes[md_metadata_propertytype]
	return ok
}

func (md_metadata_type *MD_Metadata_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Metadata_Types[md_metadata_type]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractds_aggregate_type) {
		return
	}

	abstractds_aggregate_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range abstractds_aggregate_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range abstractds_aggregate_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_aggregate_propertytype) {
		return
	}

	ds_aggregate_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_aggregate_propertytype.DS_OtherAggregate != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_OtherAggregate)
	}
	if ds_aggregate_propertytype.DS_StereoMate != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_StereoMate)
	}
	if ds_aggregate_propertytype.DS_Series != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_Series)
	}
	if ds_aggregate_propertytype.DS_Platform != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_Platform)
	}
	if ds_aggregate_propertytype.DS_Sensor != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_Sensor)
	}
	if ds_aggregate_propertytype.DS_ProductionSeries != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_ProductionSeries)
	}
	if ds_aggregate_propertytype.DS_Initiative != nil {
		stage.StageBranch(ds_aggregate_propertytype.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_dataset_propertytype) {
		return
	}

	ds_dataset_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_dataset_propertytype.DS_DataSet != nil {
		stage.StageBranch(ds_dataset_propertytype.DS_DataSet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_dataset_type *DS_DataSet_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_dataset_type) {
		return
	}

	ds_dataset_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_metadata_propertytype := range ds_dataset_type.Has {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_dataset_type.PartOf {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_initiative_propertytype) {
		return
	}

	ds_initiative_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_initiative_propertytype.DS_Initiative != nil {
		stage.StageBranch(ds_initiative_propertytype.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_initiative_type *DS_Initiative_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_initiative_type) {
		return
	}

	ds_initiative_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_initiative_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_initiative_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_otheraggregate_propertytype) {
		return
	}

	ds_otheraggregate_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
		stage.StageBranch(ds_otheraggregate_propertytype.DS_OtherAggregate)
	}
	if ds_otheraggregate_propertytype.DS_StereoMate != nil {
		stage.StageBranch(ds_otheraggregate_propertytype.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_otheraggregate_type) {
		return
	}

	ds_otheraggregate_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_otheraggregate_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_otheraggregate_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_platform_propertytype) {
		return
	}

	ds_platform_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_platform_propertytype.DS_Platform != nil {
		stage.StageBranch(ds_platform_propertytype.DS_Platform)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_platform_type *DS_Platform_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_platform_type) {
		return
	}

	ds_platform_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_platform_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_platform_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_platform_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_platform_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_productionseries_propertytype) {
		return
	}

	ds_productionseries_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_productionseries_propertytype.DS_ProductionSeries != nil {
		stage.StageBranch(ds_productionseries_propertytype.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_productionseries_type) {
		return
	}

	ds_productionseries_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_productionseries_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_productionseries_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_sensor_propertytype) {
		return
	}

	ds_sensor_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_sensor_propertytype.DS_Sensor != nil {
		stage.StageBranch(ds_sensor_propertytype.DS_Sensor)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_sensor_type *DS_Sensor_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_sensor_type) {
		return
	}

	ds_sensor_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_sensor_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_sensor_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_series_propertytype *DS_Series_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_series_propertytype) {
		return
	}

	ds_series_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_series_propertytype.DS_Series != nil {
		stage.StageBranch(ds_series_propertytype.DS_Series)
	}
	if ds_series_propertytype.DS_Platform != nil {
		stage.StageBranch(ds_series_propertytype.DS_Platform)
	}
	if ds_series_propertytype.DS_Sensor != nil {
		stage.StageBranch(ds_series_propertytype.DS_Sensor)
	}
	if ds_series_propertytype.DS_ProductionSeries != nil {
		stage.StageBranch(ds_series_propertytype.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_series_type *DS_Series_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_series_type) {
		return
	}

	ds_series_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_series_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_series_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_series_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_series_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_stereomate_propertytype) {
		return
	}

	ds_stereomate_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_stereomate_propertytype.DS_StereoMate != nil {
		stage.StageBranch(ds_stereomate_propertytype.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_stereomate_type *DS_StereoMate_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ds_stereomate_type) {
		return
	}

	ds_stereomate_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_stereomate_type.ComposedOf {
		stage.StageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_stereomate_type.SeriesMetadata {
		stage.StageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_type.Subset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_type.Superset {
		stage.StageBranch(_ds_aggregate_propertytype)
	}

}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_metadata_propertytype) {
		return
	}

	md_metadata_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_metadata_propertytype.MD_Metadata != nil {
		stage.StageBranch(md_metadata_propertytype.MD_Metadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadata_type *MD_Metadata_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_metadata_type) {
		return
	}

	md_metadata_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_aggregate_propertytype := range md_metadata_type.Series {
		stage.StageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_dataset_propertytype := range md_metadata_type.Describes {
		stage.StageBranch(_ds_dataset_propertytype)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AbstractDS_Aggregate_Type:
		toT := GongCopyBranchAbstractDS_Aggregate_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Aggregate_PropertyType:
		toT := GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_DataSet_PropertyType:
		toT := GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_DataSet_Type:
		toT := GongCopyBranchDS_DataSet_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Initiative_PropertyType:
		toT := GongCopyBranchDS_Initiative_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Initiative_Type:
		toT := GongCopyBranchDS_Initiative_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_OtherAggregate_PropertyType:
		toT := GongCopyBranchDS_OtherAggregate_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_OtherAggregate_Type:
		toT := GongCopyBranchDS_OtherAggregate_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Platform_PropertyType:
		toT := GongCopyBranchDS_Platform_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Platform_Type:
		toT := GongCopyBranchDS_Platform_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_ProductionSeries_PropertyType:
		toT := GongCopyBranchDS_ProductionSeries_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_ProductionSeries_Type:
		toT := GongCopyBranchDS_ProductionSeries_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Sensor_PropertyType:
		toT := GongCopyBranchDS_Sensor_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Sensor_Type:
		toT := GongCopyBranchDS_Sensor_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Series_PropertyType:
		toT := GongCopyBranchDS_Series_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_Series_Type:
		toT := GongCopyBranchDS_Series_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_StereoMate_PropertyType:
		toT := GongCopyBranchDS_StereoMate_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DS_StereoMate_Type:
		toT := GongCopyBranchDS_StereoMate_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Metadata_PropertyType:
		toT := GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Metadata_Type:
		toT := GongCopyBranchMD_Metadata_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractDS_Aggregate_Type(mapOrigCopy map[any]any, abstractds_aggregate_typeFrom *AbstractDS_Aggregate_Type) (abstractds_aggregate_typeTo *AbstractDS_Aggregate_Type) {
	var alreadyCopied bool
	abstractds_aggregate_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractds_aggregate_typeFrom)
	if alreadyCopied {
		return
	}
	abstractds_aggregate_typeFrom.GongCopyBasicFields(abstractds_aggregate_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range abstractds_aggregate_typeFrom.ComposedOf {
		abstractds_aggregate_typeTo.ComposedOf = append(abstractds_aggregate_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range abstractds_aggregate_typeFrom.SeriesMetadata {
		abstractds_aggregate_typeTo.SeriesMetadata = append(abstractds_aggregate_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_typeFrom.Subset {
		abstractds_aggregate_typeTo.Subset = append(abstractds_aggregate_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_typeFrom.Superset {
		abstractds_aggregate_typeTo.Superset = append(abstractds_aggregate_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy map[any]any, ds_aggregate_propertytypeFrom *DS_Aggregate_PropertyType) (ds_aggregate_propertytypeTo *DS_Aggregate_PropertyType) {
	var alreadyCopied bool
	ds_aggregate_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_aggregate_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_aggregate_propertytypeFrom.GongCopyBasicFields(ds_aggregate_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_aggregate_propertytypeFrom.DS_OtherAggregate != nil {
		ds_aggregate_propertytypeTo.DS_OtherAggregate = GongCopyBranchDS_OtherAggregate_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_OtherAggregate)
	}
	if ds_aggregate_propertytypeFrom.DS_StereoMate != nil {
		ds_aggregate_propertytypeTo.DS_StereoMate = GongCopyBranchDS_StereoMate_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_StereoMate)
	}
	if ds_aggregate_propertytypeFrom.DS_Series != nil {
		ds_aggregate_propertytypeTo.DS_Series = GongCopyBranchDS_Series_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_Series)
	}
	if ds_aggregate_propertytypeFrom.DS_Platform != nil {
		ds_aggregate_propertytypeTo.DS_Platform = GongCopyBranchDS_Platform_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_Platform)
	}
	if ds_aggregate_propertytypeFrom.DS_Sensor != nil {
		ds_aggregate_propertytypeTo.DS_Sensor = GongCopyBranchDS_Sensor_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_Sensor)
	}
	if ds_aggregate_propertytypeFrom.DS_ProductionSeries != nil {
		ds_aggregate_propertytypeTo.DS_ProductionSeries = GongCopyBranchDS_ProductionSeries_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_ProductionSeries)
	}
	if ds_aggregate_propertytypeFrom.DS_Initiative != nil {
		ds_aggregate_propertytypeTo.DS_Initiative = GongCopyBranchDS_Initiative_Type(mapOrigCopy, ds_aggregate_propertytypeFrom.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy map[any]any, ds_dataset_propertytypeFrom *DS_DataSet_PropertyType) (ds_dataset_propertytypeTo *DS_DataSet_PropertyType) {
	var alreadyCopied bool
	ds_dataset_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_dataset_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_dataset_propertytypeFrom.GongCopyBasicFields(ds_dataset_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_dataset_propertytypeFrom.DS_DataSet != nil {
		ds_dataset_propertytypeTo.DS_DataSet = GongCopyBranchDS_DataSet_Type(mapOrigCopy, ds_dataset_propertytypeFrom.DS_DataSet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_DataSet_Type(mapOrigCopy map[any]any, ds_dataset_typeFrom *DS_DataSet_Type) (ds_dataset_typeTo *DS_DataSet_Type) {
	var alreadyCopied bool
	ds_dataset_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_dataset_typeFrom)
	if alreadyCopied {
		return
	}
	ds_dataset_typeFrom.GongCopyBasicFields(ds_dataset_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_metadata_propertytype := range ds_dataset_typeFrom.Has {
		ds_dataset_typeTo.Has = append(ds_dataset_typeTo.Has, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_dataset_typeFrom.PartOf {
		ds_dataset_typeTo.PartOf = append(ds_dataset_typeTo.PartOf, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_Initiative_PropertyType(mapOrigCopy map[any]any, ds_initiative_propertytypeFrom *DS_Initiative_PropertyType) (ds_initiative_propertytypeTo *DS_Initiative_PropertyType) {
	var alreadyCopied bool
	ds_initiative_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_initiative_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_initiative_propertytypeFrom.GongCopyBasicFields(ds_initiative_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_initiative_propertytypeFrom.DS_Initiative != nil {
		ds_initiative_propertytypeTo.DS_Initiative = GongCopyBranchDS_Initiative_Type(mapOrigCopy, ds_initiative_propertytypeFrom.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_Initiative_Type(mapOrigCopy map[any]any, ds_initiative_typeFrom *DS_Initiative_Type) (ds_initiative_typeTo *DS_Initiative_Type) {
	var alreadyCopied bool
	ds_initiative_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_initiative_typeFrom)
	if alreadyCopied {
		return
	}
	ds_initiative_typeFrom.GongCopyBasicFields(ds_initiative_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_initiative_typeFrom.ComposedOf {
		ds_initiative_typeTo.ComposedOf = append(ds_initiative_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_initiative_typeFrom.SeriesMetadata {
		ds_initiative_typeTo.SeriesMetadata = append(ds_initiative_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_typeFrom.Subset {
		ds_initiative_typeTo.Subset = append(ds_initiative_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_typeFrom.Superset {
		ds_initiative_typeTo.Superset = append(ds_initiative_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_OtherAggregate_PropertyType(mapOrigCopy map[any]any, ds_otheraggregate_propertytypeFrom *DS_OtherAggregate_PropertyType) (ds_otheraggregate_propertytypeTo *DS_OtherAggregate_PropertyType) {
	var alreadyCopied bool
	ds_otheraggregate_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_otheraggregate_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_otheraggregate_propertytypeFrom.GongCopyBasicFields(ds_otheraggregate_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_otheraggregate_propertytypeFrom.DS_OtherAggregate != nil {
		ds_otheraggregate_propertytypeTo.DS_OtherAggregate = GongCopyBranchDS_OtherAggregate_Type(mapOrigCopy, ds_otheraggregate_propertytypeFrom.DS_OtherAggregate)
	}
	if ds_otheraggregate_propertytypeFrom.DS_StereoMate != nil {
		ds_otheraggregate_propertytypeTo.DS_StereoMate = GongCopyBranchDS_StereoMate_Type(mapOrigCopy, ds_otheraggregate_propertytypeFrom.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_OtherAggregate_Type(mapOrigCopy map[any]any, ds_otheraggregate_typeFrom *DS_OtherAggregate_Type) (ds_otheraggregate_typeTo *DS_OtherAggregate_Type) {
	var alreadyCopied bool
	ds_otheraggregate_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_otheraggregate_typeFrom)
	if alreadyCopied {
		return
	}
	ds_otheraggregate_typeFrom.GongCopyBasicFields(ds_otheraggregate_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_otheraggregate_typeFrom.ComposedOf {
		ds_otheraggregate_typeTo.ComposedOf = append(ds_otheraggregate_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_otheraggregate_typeFrom.SeriesMetadata {
		ds_otheraggregate_typeTo.SeriesMetadata = append(ds_otheraggregate_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_typeFrom.Subset {
		ds_otheraggregate_typeTo.Subset = append(ds_otheraggregate_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_typeFrom.Superset {
		ds_otheraggregate_typeTo.Superset = append(ds_otheraggregate_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_Platform_PropertyType(mapOrigCopy map[any]any, ds_platform_propertytypeFrom *DS_Platform_PropertyType) (ds_platform_propertytypeTo *DS_Platform_PropertyType) {
	var alreadyCopied bool
	ds_platform_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_platform_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_platform_propertytypeFrom.GongCopyBasicFields(ds_platform_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_platform_propertytypeFrom.DS_Platform != nil {
		ds_platform_propertytypeTo.DS_Platform = GongCopyBranchDS_Platform_Type(mapOrigCopy, ds_platform_propertytypeFrom.DS_Platform)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_Platform_Type(mapOrigCopy map[any]any, ds_platform_typeFrom *DS_Platform_Type) (ds_platform_typeTo *DS_Platform_Type) {
	var alreadyCopied bool
	ds_platform_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_platform_typeFrom)
	if alreadyCopied {
		return
	}
	ds_platform_typeFrom.GongCopyBasicFields(ds_platform_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_platform_typeFrom.ComposedOf {
		ds_platform_typeTo.ComposedOf = append(ds_platform_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_platform_typeFrom.SeriesMetadata {
		ds_platform_typeTo.SeriesMetadata = append(ds_platform_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_platform_typeFrom.Subset {
		ds_platform_typeTo.Subset = append(ds_platform_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_platform_typeFrom.Superset {
		ds_platform_typeTo.Superset = append(ds_platform_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_ProductionSeries_PropertyType(mapOrigCopy map[any]any, ds_productionseries_propertytypeFrom *DS_ProductionSeries_PropertyType) (ds_productionseries_propertytypeTo *DS_ProductionSeries_PropertyType) {
	var alreadyCopied bool
	ds_productionseries_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_productionseries_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_productionseries_propertytypeFrom.GongCopyBasicFields(ds_productionseries_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_productionseries_propertytypeFrom.DS_ProductionSeries != nil {
		ds_productionseries_propertytypeTo.DS_ProductionSeries = GongCopyBranchDS_ProductionSeries_Type(mapOrigCopy, ds_productionseries_propertytypeFrom.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_ProductionSeries_Type(mapOrigCopy map[any]any, ds_productionseries_typeFrom *DS_ProductionSeries_Type) (ds_productionseries_typeTo *DS_ProductionSeries_Type) {
	var alreadyCopied bool
	ds_productionseries_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_productionseries_typeFrom)
	if alreadyCopied {
		return
	}
	ds_productionseries_typeFrom.GongCopyBasicFields(ds_productionseries_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_productionseries_typeFrom.ComposedOf {
		ds_productionseries_typeTo.ComposedOf = append(ds_productionseries_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_productionseries_typeFrom.SeriesMetadata {
		ds_productionseries_typeTo.SeriesMetadata = append(ds_productionseries_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_typeFrom.Subset {
		ds_productionseries_typeTo.Subset = append(ds_productionseries_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_typeFrom.Superset {
		ds_productionseries_typeTo.Superset = append(ds_productionseries_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_Sensor_PropertyType(mapOrigCopy map[any]any, ds_sensor_propertytypeFrom *DS_Sensor_PropertyType) (ds_sensor_propertytypeTo *DS_Sensor_PropertyType) {
	var alreadyCopied bool
	ds_sensor_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_sensor_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_sensor_propertytypeFrom.GongCopyBasicFields(ds_sensor_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_sensor_propertytypeFrom.DS_Sensor != nil {
		ds_sensor_propertytypeTo.DS_Sensor = GongCopyBranchDS_Sensor_Type(mapOrigCopy, ds_sensor_propertytypeFrom.DS_Sensor)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_Sensor_Type(mapOrigCopy map[any]any, ds_sensor_typeFrom *DS_Sensor_Type) (ds_sensor_typeTo *DS_Sensor_Type) {
	var alreadyCopied bool
	ds_sensor_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_sensor_typeFrom)
	if alreadyCopied {
		return
	}
	ds_sensor_typeFrom.GongCopyBasicFields(ds_sensor_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_sensor_typeFrom.ComposedOf {
		ds_sensor_typeTo.ComposedOf = append(ds_sensor_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_sensor_typeFrom.SeriesMetadata {
		ds_sensor_typeTo.SeriesMetadata = append(ds_sensor_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_typeFrom.Subset {
		ds_sensor_typeTo.Subset = append(ds_sensor_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_typeFrom.Superset {
		ds_sensor_typeTo.Superset = append(ds_sensor_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_Series_PropertyType(mapOrigCopy map[any]any, ds_series_propertytypeFrom *DS_Series_PropertyType) (ds_series_propertytypeTo *DS_Series_PropertyType) {
	var alreadyCopied bool
	ds_series_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_series_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_series_propertytypeFrom.GongCopyBasicFields(ds_series_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_series_propertytypeFrom.DS_Series != nil {
		ds_series_propertytypeTo.DS_Series = GongCopyBranchDS_Series_Type(mapOrigCopy, ds_series_propertytypeFrom.DS_Series)
	}
	if ds_series_propertytypeFrom.DS_Platform != nil {
		ds_series_propertytypeTo.DS_Platform = GongCopyBranchDS_Platform_Type(mapOrigCopy, ds_series_propertytypeFrom.DS_Platform)
	}
	if ds_series_propertytypeFrom.DS_Sensor != nil {
		ds_series_propertytypeTo.DS_Sensor = GongCopyBranchDS_Sensor_Type(mapOrigCopy, ds_series_propertytypeFrom.DS_Sensor)
	}
	if ds_series_propertytypeFrom.DS_ProductionSeries != nil {
		ds_series_propertytypeTo.DS_ProductionSeries = GongCopyBranchDS_ProductionSeries_Type(mapOrigCopy, ds_series_propertytypeFrom.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_Series_Type(mapOrigCopy map[any]any, ds_series_typeFrom *DS_Series_Type) (ds_series_typeTo *DS_Series_Type) {
	var alreadyCopied bool
	ds_series_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_series_typeFrom)
	if alreadyCopied {
		return
	}
	ds_series_typeFrom.GongCopyBasicFields(ds_series_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_series_typeFrom.ComposedOf {
		ds_series_typeTo.ComposedOf = append(ds_series_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_series_typeFrom.SeriesMetadata {
		ds_series_typeTo.SeriesMetadata = append(ds_series_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_series_typeFrom.Subset {
		ds_series_typeTo.Subset = append(ds_series_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_series_typeFrom.Superset {
		ds_series_typeTo.Superset = append(ds_series_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchDS_StereoMate_PropertyType(mapOrigCopy map[any]any, ds_stereomate_propertytypeFrom *DS_StereoMate_PropertyType) (ds_stereomate_propertytypeTo *DS_StereoMate_PropertyType) {
	var alreadyCopied bool
	ds_stereomate_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_stereomate_propertytypeFrom)
	if alreadyCopied {
		return
	}
	ds_stereomate_propertytypeFrom.GongCopyBasicFields(ds_stereomate_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if ds_stereomate_propertytypeFrom.DS_StereoMate != nil {
		ds_stereomate_propertytypeTo.DS_StereoMate = GongCopyBranchDS_StereoMate_Type(mapOrigCopy, ds_stereomate_propertytypeFrom.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDS_StereoMate_Type(mapOrigCopy map[any]any, ds_stereomate_typeFrom *DS_StereoMate_Type) (ds_stereomate_typeTo *DS_StereoMate_Type) {
	var alreadyCopied bool
	ds_stereomate_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ds_stereomate_typeFrom)
	if alreadyCopied {
		return
	}
	ds_stereomate_typeFrom.GongCopyBasicFields(ds_stereomate_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_stereomate_typeFrom.ComposedOf {
		ds_stereomate_typeTo.ComposedOf = append(ds_stereomate_typeTo.ComposedOf, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}
	for _, _md_metadata_propertytype := range ds_stereomate_typeFrom.SeriesMetadata {
		ds_stereomate_typeTo.SeriesMetadata = append(ds_stereomate_typeTo.SeriesMetadata, GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy, _md_metadata_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_typeFrom.Subset {
		ds_stereomate_typeTo.Subset = append(ds_stereomate_typeTo.Subset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_typeFrom.Superset {
		ds_stereomate_typeTo.Superset = append(ds_stereomate_typeTo.Superset, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}

	return
}

func GongCopyBranchMD_Metadata_PropertyType(mapOrigCopy map[any]any, md_metadata_propertytypeFrom *MD_Metadata_PropertyType) (md_metadata_propertytypeTo *MD_Metadata_PropertyType) {
	var alreadyCopied bool
	md_metadata_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_metadata_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_metadata_propertytypeFrom.GongCopyBasicFields(md_metadata_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_metadata_propertytypeFrom.MD_Metadata != nil {
		md_metadata_propertytypeTo.MD_Metadata = GongCopyBranchMD_Metadata_Type(mapOrigCopy, md_metadata_propertytypeFrom.MD_Metadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Metadata_Type(mapOrigCopy map[any]any, md_metadata_typeFrom *MD_Metadata_Type) (md_metadata_typeTo *MD_Metadata_Type) {
	var alreadyCopied bool
	md_metadata_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_metadata_typeFrom)
	if alreadyCopied {
		return
	}
	md_metadata_typeFrom.GongCopyBasicFields(md_metadata_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_aggregate_propertytype := range md_metadata_typeFrom.Series {
		md_metadata_typeTo.Series = append(md_metadata_typeTo.Series, GongCopyBranchDS_Aggregate_PropertyType(mapOrigCopy, _ds_aggregate_propertytype))
	}
	for _, _ds_dataset_propertytype := range md_metadata_typeFrom.Describes {
		md_metadata_typeTo.Describes = append(md_metadata_typeTo.Describes, GongCopyBranchDS_DataSet_PropertyType(mapOrigCopy, _ds_dataset_propertytype))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractds_aggregate_type) {
		return
	}

	abstractds_aggregate_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range abstractds_aggregate_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range abstractds_aggregate_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_aggregate_propertytype) {
		return
	}

	ds_aggregate_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_aggregate_propertytype.DS_OtherAggregate != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_OtherAggregate)
	}
	if ds_aggregate_propertytype.DS_StereoMate != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_StereoMate)
	}
	if ds_aggregate_propertytype.DS_Series != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_Series)
	}
	if ds_aggregate_propertytype.DS_Platform != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_Platform)
	}
	if ds_aggregate_propertytype.DS_Sensor != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_Sensor)
	}
	if ds_aggregate_propertytype.DS_ProductionSeries != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_ProductionSeries)
	}
	if ds_aggregate_propertytype.DS_Initiative != nil {
		stage.UnstageBranch(ds_aggregate_propertytype.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_dataset_propertytype) {
		return
	}

	ds_dataset_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_dataset_propertytype.DS_DataSet != nil {
		stage.UnstageBranch(ds_dataset_propertytype.DS_DataSet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_dataset_type *DS_DataSet_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_dataset_type) {
		return
	}

	ds_dataset_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_metadata_propertytype := range ds_dataset_type.Has {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_dataset_type.PartOf {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_initiative_propertytype) {
		return
	}

	ds_initiative_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_initiative_propertytype.DS_Initiative != nil {
		stage.UnstageBranch(ds_initiative_propertytype.DS_Initiative)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_initiative_type *DS_Initiative_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_initiative_type) {
		return
	}

	ds_initiative_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_initiative_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_initiative_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_initiative_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_otheraggregate_propertytype) {
		return
	}

	ds_otheraggregate_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
		stage.UnstageBranch(ds_otheraggregate_propertytype.DS_OtherAggregate)
	}
	if ds_otheraggregate_propertytype.DS_StereoMate != nil {
		stage.UnstageBranch(ds_otheraggregate_propertytype.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_otheraggregate_type) {
		return
	}

	ds_otheraggregate_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_otheraggregate_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_otheraggregate_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_platform_propertytype) {
		return
	}

	ds_platform_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_platform_propertytype.DS_Platform != nil {
		stage.UnstageBranch(ds_platform_propertytype.DS_Platform)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_platform_type *DS_Platform_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_platform_type) {
		return
	}

	ds_platform_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_platform_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_platform_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_platform_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_platform_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_productionseries_propertytype) {
		return
	}

	ds_productionseries_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_productionseries_propertytype.DS_ProductionSeries != nil {
		stage.UnstageBranch(ds_productionseries_propertytype.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_productionseries_type) {
		return
	}

	ds_productionseries_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_productionseries_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_productionseries_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_productionseries_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_sensor_propertytype) {
		return
	}

	ds_sensor_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_sensor_propertytype.DS_Sensor != nil {
		stage.UnstageBranch(ds_sensor_propertytype.DS_Sensor)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_sensor_type *DS_Sensor_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_sensor_type) {
		return
	}

	ds_sensor_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_sensor_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_sensor_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_sensor_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_series_propertytype *DS_Series_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_series_propertytype) {
		return
	}

	ds_series_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_series_propertytype.DS_Series != nil {
		stage.UnstageBranch(ds_series_propertytype.DS_Series)
	}
	if ds_series_propertytype.DS_Platform != nil {
		stage.UnstageBranch(ds_series_propertytype.DS_Platform)
	}
	if ds_series_propertytype.DS_Sensor != nil {
		stage.UnstageBranch(ds_series_propertytype.DS_Sensor)
	}
	if ds_series_propertytype.DS_ProductionSeries != nil {
		stage.UnstageBranch(ds_series_propertytype.DS_ProductionSeries)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_series_type *DS_Series_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_series_type) {
		return
	}

	ds_series_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_series_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_series_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_series_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_series_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_stereomate_propertytype) {
		return
	}

	ds_stereomate_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ds_stereomate_propertytype.DS_StereoMate != nil {
		stage.UnstageBranch(ds_stereomate_propertytype.DS_StereoMate)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ds_stereomate_type *DS_StereoMate_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ds_stereomate_type) {
		return
	}

	ds_stereomate_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_dataset_propertytype := range ds_stereomate_type.ComposedOf {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}
	for _, _md_metadata_propertytype := range ds_stereomate_type.SeriesMetadata {
		stage.UnstageBranch(_md_metadata_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_type.Subset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_aggregate_propertytype := range ds_stereomate_type.Superset {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}

}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_metadata_propertytype) {
		return
	}

	md_metadata_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_metadata_propertytype.MD_Metadata != nil {
		stage.UnstageBranch(md_metadata_propertytype.MD_Metadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadata_type *MD_Metadata_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_metadata_type) {
		return
	}

	md_metadata_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ds_aggregate_propertytype := range md_metadata_type.Series {
		stage.UnstageBranch(_ds_aggregate_propertytype)
	}
	for _, _ds_dataset_propertytype := range md_metadata_type.Describes {
		stage.UnstageBranch(_ds_dataset_propertytype)
	}

}

// insertion point for pointer reconstruction from references
func (reference *AbstractDS_Aggregate_Type) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractDS_Aggregate_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_Aggregate_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Aggregate_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_OtherAggregate, stage.DS_OtherAggregate_Types_reference, instance.DS_OtherAggregate)
	__gong__reconstructPointer(&reference.DS_StereoMate, stage.DS_StereoMate_Types_reference, instance.DS_StereoMate)
	__gong__reconstructPointer(&reference.DS_Series, stage.DS_Series_Types_reference, instance.DS_Series)
	__gong__reconstructPointer(&reference.DS_Platform, stage.DS_Platform_Types_reference, instance.DS_Platform)
	__gong__reconstructPointer(&reference.DS_Sensor, stage.DS_Sensor_Types_reference, instance.DS_Sensor)
	__gong__reconstructPointer(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_reference, instance.DS_ProductionSeries)
	__gong__reconstructPointer(&reference.DS_Initiative, stage.DS_Initiative_Types_reference, instance.DS_Initiative)
	// insertion point for slice of pointers field
}

func (reference *DS_DataSet_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_DataSet_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_DataSet, stage.DS_DataSet_Types_reference, instance.DS_DataSet)
	// insertion point for slice of pointers field
}

func (reference *DS_DataSet_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_DataSet_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Has, stage.MD_Metadata_PropertyTypes_reference, instance.Has)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartOf, stage.DS_Aggregate_PropertyTypes_reference, instance.PartOf)
}

func (reference *DS_Initiative_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Initiative_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_Initiative, stage.DS_Initiative_Types_reference, instance.DS_Initiative)
	// insertion point for slice of pointers field
}

func (reference *DS_Initiative_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Initiative_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_OtherAggregate_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_OtherAggregate_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_OtherAggregate, stage.DS_OtherAggregate_Types_reference, instance.DS_OtherAggregate)
	__gong__reconstructPointer(&reference.DS_StereoMate, stage.DS_StereoMate_Types_reference, instance.DS_StereoMate)
	// insertion point for slice of pointers field
}

func (reference *DS_OtherAggregate_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_OtherAggregate_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_Platform_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Platform_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_Platform, stage.DS_Platform_Types_reference, instance.DS_Platform)
	// insertion point for slice of pointers field
}

func (reference *DS_Platform_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Platform_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_ProductionSeries_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_ProductionSeries_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_reference, instance.DS_ProductionSeries)
	// insertion point for slice of pointers field
}

func (reference *DS_ProductionSeries_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_ProductionSeries_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_Sensor_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Sensor_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_Sensor, stage.DS_Sensor_Types_reference, instance.DS_Sensor)
	// insertion point for slice of pointers field
}

func (reference *DS_Sensor_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Sensor_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_Series_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Series_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_Series, stage.DS_Series_Types_reference, instance.DS_Series)
	__gong__reconstructPointer(&reference.DS_Platform, stage.DS_Platform_Types_reference, instance.DS_Platform)
	__gong__reconstructPointer(&reference.DS_Sensor, stage.DS_Sensor_Types_reference, instance.DS_Sensor)
	__gong__reconstructPointer(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_reference, instance.DS_ProductionSeries)
	// insertion point for slice of pointers field
}

func (reference *DS_Series_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_Series_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *DS_StereoMate_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *DS_StereoMate_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DS_StereoMate, stage.DS_StereoMate_Types_reference, instance.DS_StereoMate)
	// insertion point for slice of pointers field
}

func (reference *DS_StereoMate_Type) GongReconstructPointersFromReferences(stage *Stage, instance *DS_StereoMate_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_reference, instance.ComposedOf)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_reference, instance.SeriesMetadata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Subset, stage.DS_Aggregate_PropertyTypes_reference, instance.Subset)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Superset, stage.DS_Aggregate_PropertyTypes_reference, instance.Superset)
}

func (reference *MD_Metadata_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Metadata_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_Metadata, stage.MD_Metadata_Types_reference, instance.MD_Metadata)
	// insertion point for slice of pointers field
}

func (reference *MD_Metadata_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Metadata_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Series, stage.DS_Aggregate_PropertyTypes_reference, instance.Series)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Describes, stage.DS_DataSet_PropertyTypes_reference, instance.Describes)
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractDS_Aggregate_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_Aggregate_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_OtherAggregate, stage.DS_OtherAggregate_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_StereoMate, stage.DS_StereoMate_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Series, stage.DS_Series_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Platform, stage.DS_Platform_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Sensor, stage.DS_Sensor_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Initiative, stage.DS_Initiative_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_DataSet_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_DataSet, stage.DS_DataSet_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_DataSet_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Has, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartOf, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_Initiative_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_Initiative, stage.DS_Initiative_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_Initiative_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_OtherAggregate_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_OtherAggregate, stage.DS_OtherAggregate_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_StereoMate, stage.DS_StereoMate_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_OtherAggregate_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_Platform_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_Platform, stage.DS_Platform_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_Platform_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_ProductionSeries_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_ProductionSeries_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_Sensor_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_Sensor, stage.DS_Sensor_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_Sensor_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_Series_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_Series, stage.DS_Series_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Platform, stage.DS_Platform_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_Sensor, stage.DS_Sensor_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.DS_ProductionSeries, stage.DS_ProductionSeries_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_Series_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *DS_StereoMate_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DS_StereoMate, stage.DS_StereoMate_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *DS_StereoMate_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComposedOf, stage.DS_DataSet_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SeriesMetadata, stage.MD_Metadata_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Subset, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Superset, stage.DS_Aggregate_PropertyTypes_instance)
}

func (reference *MD_Metadata_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_Metadata, stage.MD_Metadata_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Metadata_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Series, stage.DS_Aggregate_PropertyTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Describes, stage.DS_DataSet_PropertyTypes_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongDiff(stage *Stage, abstractds_aggregate_typeOther *AbstractDS_Aggregate_Type) (diffs []string) {
	// insertion point for field diffs
	if abstractds_aggregate_type.Name != abstractds_aggregate_typeOther.Name {
		diffs = append(diffs, abstractds_aggregate_type.GongMarshallField(stage, "Name"))
	}
	if abstractds_aggregate_type.Id != abstractds_aggregate_typeOther.Id {
		diffs = append(diffs, abstractds_aggregate_type.GongMarshallField(stage, "Id"))
	}
	if abstractds_aggregate_type.Uuid != abstractds_aggregate_typeOther.Uuid {
		diffs = append(diffs, abstractds_aggregate_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractds_aggregate_type, "ComposedOf", abstractds_aggregate_typeOther.ComposedOf, abstractds_aggregate_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractds_aggregate_type, "SeriesMetadata", abstractds_aggregate_typeOther.SeriesMetadata, abstractds_aggregate_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractds_aggregate_type, "Subset", abstractds_aggregate_typeOther.Subset, abstractds_aggregate_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractds_aggregate_type, "Superset", abstractds_aggregate_typeOther.Superset, abstractds_aggregate_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongDiff(stage *Stage, ds_aggregate_propertytypeOther *DS_Aggregate_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_aggregate_propertytype.Name != ds_aggregate_propertytypeOther.Name {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_aggregate_propertytype.NilReason != ds_aggregate_propertytypeOther.NilReason {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_aggregate_propertytype.Uuidref != ds_aggregate_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_aggregate_propertytype.Type_XSD != ds_aggregate_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_aggregate_propertytype.Href != ds_aggregate_propertytypeOther.Href {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_aggregate_propertytype.Role != ds_aggregate_propertytypeOther.Role {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_aggregate_propertytype.Arcrole != ds_aggregate_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_aggregate_propertytype.Title != ds_aggregate_propertytypeOther.Title {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_aggregate_propertytype.Show != ds_aggregate_propertytypeOther.Show {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_aggregate_propertytype.Actuate != ds_aggregate_propertytypeOther.Actuate {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_aggregate_propertytype.DS_OtherAggregate != ds_aggregate_propertytypeOther.DS_OtherAggregate {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_OtherAggregate"))
	}
	if ds_aggregate_propertytype.DS_StereoMate != ds_aggregate_propertytypeOther.DS_StereoMate {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
	}
	if ds_aggregate_propertytype.DS_Series != ds_aggregate_propertytypeOther.DS_Series {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_Series"))
	}
	if ds_aggregate_propertytype.DS_Platform != ds_aggregate_propertytypeOther.DS_Platform {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_Platform"))
	}
	if ds_aggregate_propertytype.DS_Sensor != ds_aggregate_propertytypeOther.DS_Sensor {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_Sensor"))
	}
	if ds_aggregate_propertytype.DS_ProductionSeries != ds_aggregate_propertytypeOther.DS_ProductionSeries {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
	}
	if ds_aggregate_propertytype.DS_Initiative != ds_aggregate_propertytypeOther.DS_Initiative {
		diffs = append(diffs, ds_aggregate_propertytype.GongMarshallField(stage, "DS_Initiative"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongDiff(stage *Stage, ds_dataset_propertytypeOther *DS_DataSet_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_dataset_propertytype.Name != ds_dataset_propertytypeOther.Name {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_dataset_propertytype.NilReason != ds_dataset_propertytypeOther.NilReason {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_dataset_propertytype.Uuidref != ds_dataset_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_dataset_propertytype.Type_XSD != ds_dataset_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_dataset_propertytype.Href != ds_dataset_propertytypeOther.Href {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_dataset_propertytype.Role != ds_dataset_propertytypeOther.Role {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_dataset_propertytype.Arcrole != ds_dataset_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_dataset_propertytype.Title != ds_dataset_propertytypeOther.Title {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_dataset_propertytype.Show != ds_dataset_propertytypeOther.Show {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_dataset_propertytype.Actuate != ds_dataset_propertytypeOther.Actuate {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_dataset_propertytype.DS_DataSet != ds_dataset_propertytypeOther.DS_DataSet {
		diffs = append(diffs, ds_dataset_propertytype.GongMarshallField(stage, "DS_DataSet"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_dataset_type *DS_DataSet_Type) GongDiff(stage *Stage, ds_dataset_typeOther *DS_DataSet_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_dataset_type.Name != ds_dataset_typeOther.Name {
		diffs = append(diffs, ds_dataset_type.GongMarshallField(stage, "Name"))
	}
	if ds_dataset_type.Id != ds_dataset_typeOther.Id {
		diffs = append(diffs, ds_dataset_type.GongMarshallField(stage, "Id"))
	}
	if ds_dataset_type.Uuid != ds_dataset_typeOther.Uuid {
		diffs = append(diffs, ds_dataset_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_dataset_type, "Has", ds_dataset_typeOther.Has, ds_dataset_type.Has); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_dataset_type, "PartOf", ds_dataset_typeOther.PartOf, ds_dataset_type.PartOf); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongDiff(stage *Stage, ds_initiative_propertytypeOther *DS_Initiative_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_initiative_propertytype.Name != ds_initiative_propertytypeOther.Name {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_initiative_propertytype.NilReason != ds_initiative_propertytypeOther.NilReason {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_initiative_propertytype.Uuidref != ds_initiative_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_initiative_propertytype.Type_XSD != ds_initiative_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_initiative_propertytype.Href != ds_initiative_propertytypeOther.Href {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_initiative_propertytype.Role != ds_initiative_propertytypeOther.Role {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_initiative_propertytype.Arcrole != ds_initiative_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_initiative_propertytype.Title != ds_initiative_propertytypeOther.Title {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_initiative_propertytype.Show != ds_initiative_propertytypeOther.Show {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_initiative_propertytype.Actuate != ds_initiative_propertytypeOther.Actuate {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_initiative_propertytype.DS_Initiative != ds_initiative_propertytypeOther.DS_Initiative {
		diffs = append(diffs, ds_initiative_propertytype.GongMarshallField(stage, "DS_Initiative"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_initiative_type *DS_Initiative_Type) GongDiff(stage *Stage, ds_initiative_typeOther *DS_Initiative_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_initiative_type.Name != ds_initiative_typeOther.Name {
		diffs = append(diffs, ds_initiative_type.GongMarshallField(stage, "Name"))
	}
	if ds_initiative_type.Id != ds_initiative_typeOther.Id {
		diffs = append(diffs, ds_initiative_type.GongMarshallField(stage, "Id"))
	}
	if ds_initiative_type.Uuid != ds_initiative_typeOther.Uuid {
		diffs = append(diffs, ds_initiative_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_initiative_type, "ComposedOf", ds_initiative_typeOther.ComposedOf, ds_initiative_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_initiative_type, "SeriesMetadata", ds_initiative_typeOther.SeriesMetadata, ds_initiative_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_initiative_type, "Subset", ds_initiative_typeOther.Subset, ds_initiative_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_initiative_type, "Superset", ds_initiative_typeOther.Superset, ds_initiative_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongDiff(stage *Stage, ds_otheraggregate_propertytypeOther *DS_OtherAggregate_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_otheraggregate_propertytype.Name != ds_otheraggregate_propertytypeOther.Name {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_otheraggregate_propertytype.NilReason != ds_otheraggregate_propertytypeOther.NilReason {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_otheraggregate_propertytype.Uuidref != ds_otheraggregate_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_otheraggregate_propertytype.Type_XSD != ds_otheraggregate_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_otheraggregate_propertytype.Href != ds_otheraggregate_propertytypeOther.Href {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_otheraggregate_propertytype.Role != ds_otheraggregate_propertytypeOther.Role {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_otheraggregate_propertytype.Arcrole != ds_otheraggregate_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_otheraggregate_propertytype.Title != ds_otheraggregate_propertytypeOther.Title {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_otheraggregate_propertytype.Show != ds_otheraggregate_propertytypeOther.Show {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_otheraggregate_propertytype.Actuate != ds_otheraggregate_propertytypeOther.Actuate {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_otheraggregate_propertytype.DS_OtherAggregate != ds_otheraggregate_propertytypeOther.DS_OtherAggregate {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "DS_OtherAggregate"))
	}
	if ds_otheraggregate_propertytype.DS_StereoMate != ds_otheraggregate_propertytypeOther.DS_StereoMate {
		diffs = append(diffs, ds_otheraggregate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongDiff(stage *Stage, ds_otheraggregate_typeOther *DS_OtherAggregate_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_otheraggregate_type.Name != ds_otheraggregate_typeOther.Name {
		diffs = append(diffs, ds_otheraggregate_type.GongMarshallField(stage, "Name"))
	}
	if ds_otheraggregate_type.Id != ds_otheraggregate_typeOther.Id {
		diffs = append(diffs, ds_otheraggregate_type.GongMarshallField(stage, "Id"))
	}
	if ds_otheraggregate_type.Uuid != ds_otheraggregate_typeOther.Uuid {
		diffs = append(diffs, ds_otheraggregate_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_otheraggregate_type, "ComposedOf", ds_otheraggregate_typeOther.ComposedOf, ds_otheraggregate_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_otheraggregate_type, "SeriesMetadata", ds_otheraggregate_typeOther.SeriesMetadata, ds_otheraggregate_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_otheraggregate_type, "Subset", ds_otheraggregate_typeOther.Subset, ds_otheraggregate_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_otheraggregate_type, "Superset", ds_otheraggregate_typeOther.Superset, ds_otheraggregate_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_platform_propertytype *DS_Platform_PropertyType) GongDiff(stage *Stage, ds_platform_propertytypeOther *DS_Platform_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_platform_propertytype.Name != ds_platform_propertytypeOther.Name {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_platform_propertytype.NilReason != ds_platform_propertytypeOther.NilReason {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_platform_propertytype.Uuidref != ds_platform_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_platform_propertytype.Type_XSD != ds_platform_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_platform_propertytype.Href != ds_platform_propertytypeOther.Href {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_platform_propertytype.Role != ds_platform_propertytypeOther.Role {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_platform_propertytype.Arcrole != ds_platform_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_platform_propertytype.Title != ds_platform_propertytypeOther.Title {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_platform_propertytype.Show != ds_platform_propertytypeOther.Show {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_platform_propertytype.Actuate != ds_platform_propertytypeOther.Actuate {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_platform_propertytype.DS_Platform != ds_platform_propertytypeOther.DS_Platform {
		diffs = append(diffs, ds_platform_propertytype.GongMarshallField(stage, "DS_Platform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_platform_type *DS_Platform_Type) GongDiff(stage *Stage, ds_platform_typeOther *DS_Platform_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_platform_type.Name != ds_platform_typeOther.Name {
		diffs = append(diffs, ds_platform_type.GongMarshallField(stage, "Name"))
	}
	if ds_platform_type.Id != ds_platform_typeOther.Id {
		diffs = append(diffs, ds_platform_type.GongMarshallField(stage, "Id"))
	}
	if ds_platform_type.Uuid != ds_platform_typeOther.Uuid {
		diffs = append(diffs, ds_platform_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_platform_type, "ComposedOf", ds_platform_typeOther.ComposedOf, ds_platform_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_platform_type, "SeriesMetadata", ds_platform_typeOther.SeriesMetadata, ds_platform_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_platform_type, "Subset", ds_platform_typeOther.Subset, ds_platform_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_platform_type, "Superset", ds_platform_typeOther.Superset, ds_platform_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongDiff(stage *Stage, ds_productionseries_propertytypeOther *DS_ProductionSeries_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_productionseries_propertytype.Name != ds_productionseries_propertytypeOther.Name {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_productionseries_propertytype.NilReason != ds_productionseries_propertytypeOther.NilReason {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_productionseries_propertytype.Uuidref != ds_productionseries_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_productionseries_propertytype.Type_XSD != ds_productionseries_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_productionseries_propertytype.Href != ds_productionseries_propertytypeOther.Href {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_productionseries_propertytype.Role != ds_productionseries_propertytypeOther.Role {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_productionseries_propertytype.Arcrole != ds_productionseries_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_productionseries_propertytype.Title != ds_productionseries_propertytypeOther.Title {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_productionseries_propertytype.Show != ds_productionseries_propertytypeOther.Show {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_productionseries_propertytype.Actuate != ds_productionseries_propertytypeOther.Actuate {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_productionseries_propertytype.DS_ProductionSeries != ds_productionseries_propertytypeOther.DS_ProductionSeries {
		diffs = append(diffs, ds_productionseries_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_productionseries_type *DS_ProductionSeries_Type) GongDiff(stage *Stage, ds_productionseries_typeOther *DS_ProductionSeries_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_productionseries_type.Name != ds_productionseries_typeOther.Name {
		diffs = append(diffs, ds_productionseries_type.GongMarshallField(stage, "Name"))
	}
	if ds_productionseries_type.Id != ds_productionseries_typeOther.Id {
		diffs = append(diffs, ds_productionseries_type.GongMarshallField(stage, "Id"))
	}
	if ds_productionseries_type.Uuid != ds_productionseries_typeOther.Uuid {
		diffs = append(diffs, ds_productionseries_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_productionseries_type, "ComposedOf", ds_productionseries_typeOther.ComposedOf, ds_productionseries_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_productionseries_type, "SeriesMetadata", ds_productionseries_typeOther.SeriesMetadata, ds_productionseries_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_productionseries_type, "Subset", ds_productionseries_typeOther.Subset, ds_productionseries_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_productionseries_type, "Superset", ds_productionseries_typeOther.Superset, ds_productionseries_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongDiff(stage *Stage, ds_sensor_propertytypeOther *DS_Sensor_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_sensor_propertytype.Name != ds_sensor_propertytypeOther.Name {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_sensor_propertytype.NilReason != ds_sensor_propertytypeOther.NilReason {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_sensor_propertytype.Uuidref != ds_sensor_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_sensor_propertytype.Type_XSD != ds_sensor_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_sensor_propertytype.Href != ds_sensor_propertytypeOther.Href {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_sensor_propertytype.Role != ds_sensor_propertytypeOther.Role {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_sensor_propertytype.Arcrole != ds_sensor_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_sensor_propertytype.Title != ds_sensor_propertytypeOther.Title {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_sensor_propertytype.Show != ds_sensor_propertytypeOther.Show {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_sensor_propertytype.Actuate != ds_sensor_propertytypeOther.Actuate {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_sensor_propertytype.DS_Sensor != ds_sensor_propertytypeOther.DS_Sensor {
		diffs = append(diffs, ds_sensor_propertytype.GongMarshallField(stage, "DS_Sensor"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_sensor_type *DS_Sensor_Type) GongDiff(stage *Stage, ds_sensor_typeOther *DS_Sensor_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_sensor_type.Name != ds_sensor_typeOther.Name {
		diffs = append(diffs, ds_sensor_type.GongMarshallField(stage, "Name"))
	}
	if ds_sensor_type.Id != ds_sensor_typeOther.Id {
		diffs = append(diffs, ds_sensor_type.GongMarshallField(stage, "Id"))
	}
	if ds_sensor_type.Uuid != ds_sensor_typeOther.Uuid {
		diffs = append(diffs, ds_sensor_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_sensor_type, "ComposedOf", ds_sensor_typeOther.ComposedOf, ds_sensor_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_sensor_type, "SeriesMetadata", ds_sensor_typeOther.SeriesMetadata, ds_sensor_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_sensor_type, "Subset", ds_sensor_typeOther.Subset, ds_sensor_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_sensor_type, "Superset", ds_sensor_typeOther.Superset, ds_sensor_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_series_propertytype *DS_Series_PropertyType) GongDiff(stage *Stage, ds_series_propertytypeOther *DS_Series_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_series_propertytype.Name != ds_series_propertytypeOther.Name {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_series_propertytype.NilReason != ds_series_propertytypeOther.NilReason {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_series_propertytype.Uuidref != ds_series_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_series_propertytype.Type_XSD != ds_series_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_series_propertytype.Href != ds_series_propertytypeOther.Href {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_series_propertytype.Role != ds_series_propertytypeOther.Role {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_series_propertytype.Arcrole != ds_series_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_series_propertytype.Title != ds_series_propertytypeOther.Title {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_series_propertytype.Show != ds_series_propertytypeOther.Show {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_series_propertytype.Actuate != ds_series_propertytypeOther.Actuate {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_series_propertytype.DS_Series != ds_series_propertytypeOther.DS_Series {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "DS_Series"))
	}
	if ds_series_propertytype.DS_Platform != ds_series_propertytypeOther.DS_Platform {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "DS_Platform"))
	}
	if ds_series_propertytype.DS_Sensor != ds_series_propertytypeOther.DS_Sensor {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "DS_Sensor"))
	}
	if ds_series_propertytype.DS_ProductionSeries != ds_series_propertytypeOther.DS_ProductionSeries {
		diffs = append(diffs, ds_series_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_series_type *DS_Series_Type) GongDiff(stage *Stage, ds_series_typeOther *DS_Series_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_series_type.Name != ds_series_typeOther.Name {
		diffs = append(diffs, ds_series_type.GongMarshallField(stage, "Name"))
	}
	if ds_series_type.Id != ds_series_typeOther.Id {
		diffs = append(diffs, ds_series_type.GongMarshallField(stage, "Id"))
	}
	if ds_series_type.Uuid != ds_series_typeOther.Uuid {
		diffs = append(diffs, ds_series_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_series_type, "ComposedOf", ds_series_typeOther.ComposedOf, ds_series_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_series_type, "SeriesMetadata", ds_series_typeOther.SeriesMetadata, ds_series_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_series_type, "Subset", ds_series_typeOther.Subset, ds_series_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_series_type, "Superset", ds_series_typeOther.Superset, ds_series_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongDiff(stage *Stage, ds_stereomate_propertytypeOther *DS_StereoMate_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if ds_stereomate_propertytype.Name != ds_stereomate_propertytypeOther.Name {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Name"))
	}
	if ds_stereomate_propertytype.NilReason != ds_stereomate_propertytypeOther.NilReason {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if ds_stereomate_propertytype.Uuidref != ds_stereomate_propertytypeOther.Uuidref {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if ds_stereomate_propertytype.Type_XSD != ds_stereomate_propertytypeOther.Type_XSD {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if ds_stereomate_propertytype.Href != ds_stereomate_propertytypeOther.Href {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Href"))
	}
	if ds_stereomate_propertytype.Role != ds_stereomate_propertytypeOther.Role {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Role"))
	}
	if ds_stereomate_propertytype.Arcrole != ds_stereomate_propertytypeOther.Arcrole {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if ds_stereomate_propertytype.Title != ds_stereomate_propertytypeOther.Title {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Title"))
	}
	if ds_stereomate_propertytype.Show != ds_stereomate_propertytypeOther.Show {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Show"))
	}
	if ds_stereomate_propertytype.Actuate != ds_stereomate_propertytypeOther.Actuate {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if ds_stereomate_propertytype.DS_StereoMate != ds_stereomate_propertytypeOther.DS_StereoMate {
		diffs = append(diffs, ds_stereomate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ds_stereomate_type *DS_StereoMate_Type) GongDiff(stage *Stage, ds_stereomate_typeOther *DS_StereoMate_Type) (diffs []string) {
	// insertion point for field diffs
	if ds_stereomate_type.Name != ds_stereomate_typeOther.Name {
		diffs = append(diffs, ds_stereomate_type.GongMarshallField(stage, "Name"))
	}
	if ds_stereomate_type.Id != ds_stereomate_typeOther.Id {
		diffs = append(diffs, ds_stereomate_type.GongMarshallField(stage, "Id"))
	}
	if ds_stereomate_type.Uuid != ds_stereomate_typeOther.Uuid {
		diffs = append(diffs, ds_stereomate_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_stereomate_type, "ComposedOf", ds_stereomate_typeOther.ComposedOf, ds_stereomate_type.ComposedOf); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_stereomate_type, "SeriesMetadata", ds_stereomate_typeOther.SeriesMetadata, ds_stereomate_type.SeriesMetadata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_stereomate_type, "Subset", ds_stereomate_typeOther.Subset, ds_stereomate_type.Subset); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ds_stereomate_type, "Superset", ds_stereomate_typeOther.Superset, ds_stereomate_type.Superset); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_metadata_propertytype *MD_Metadata_PropertyType) GongDiff(stage *Stage, md_metadata_propertytypeOther *MD_Metadata_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_metadata_propertytype.Name != md_metadata_propertytypeOther.Name {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_metadata_propertytype.NilReason != md_metadata_propertytypeOther.NilReason {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_metadata_propertytype.Uuidref != md_metadata_propertytypeOther.Uuidref {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_metadata_propertytype.Type_XSD != md_metadata_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_metadata_propertytype.Href != md_metadata_propertytypeOther.Href {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_metadata_propertytype.Role != md_metadata_propertytypeOther.Role {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_metadata_propertytype.Arcrole != md_metadata_propertytypeOther.Arcrole {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_metadata_propertytype.Title != md_metadata_propertytypeOther.Title {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_metadata_propertytype.Show != md_metadata_propertytypeOther.Show {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_metadata_propertytype.Actuate != md_metadata_propertytypeOther.Actuate {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_metadata_propertytype.MD_Metadata != md_metadata_propertytypeOther.MD_Metadata {
		diffs = append(diffs, md_metadata_propertytype.GongMarshallField(stage, "MD_Metadata"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_metadata_type *MD_Metadata_Type) GongDiff(stage *Stage, md_metadata_typeOther *MD_Metadata_Type) (diffs []string) {
	// insertion point for field diffs
	if md_metadata_type.Name != md_metadata_typeOther.Name {
		diffs = append(diffs, md_metadata_type.GongMarshallField(stage, "Name"))
	}
	if md_metadata_type.Id != md_metadata_typeOther.Id {
		diffs = append(diffs, md_metadata_type.GongMarshallField(stage, "Id"))
	}
	if md_metadata_type.Uuid != md_metadata_typeOther.Uuid {
		diffs = append(diffs, md_metadata_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_metadata_type, "Series", md_metadata_typeOther.Series, md_metadata_type.Series); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, md_metadata_type, "Describes", md_metadata_typeOther.Describes, md_metadata_type.Describes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
