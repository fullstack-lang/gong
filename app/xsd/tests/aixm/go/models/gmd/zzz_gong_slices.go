// generated code - do not edit
package gmd

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AbstractDS_Aggregate_Type
	// insertion point per field
	stage.AbstractDS_Aggregate_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*AbstractDS_Aggregate_Type)
	for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
		_ = abstractds_aggregate_type
		for _, _ds_dataset_propertytype := range abstractds_aggregate_type.ComposedOf {
			stage.AbstractDS_Aggregate_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = abstractds_aggregate_type
		}
	}
	stage.AbstractDS_Aggregate_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*AbstractDS_Aggregate_Type)
	for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
		_ = abstractds_aggregate_type
		for _, _md_metadata_propertytype := range abstractds_aggregate_type.SeriesMetadata {
			stage.AbstractDS_Aggregate_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = abstractds_aggregate_type
		}
	}
	stage.AbstractDS_Aggregate_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*AbstractDS_Aggregate_Type)
	for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
		_ = abstractds_aggregate_type
		for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Subset {
			stage.AbstractDS_Aggregate_Type_Subset_reverseMap[_ds_aggregate_propertytype] = abstractds_aggregate_type
		}
	}
	stage.AbstractDS_Aggregate_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*AbstractDS_Aggregate_Type)
	for abstractds_aggregate_type := range stage.AbstractDS_Aggregate_Types {
		_ = abstractds_aggregate_type
		for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Superset {
			stage.AbstractDS_Aggregate_Type_Superset_reverseMap[_ds_aggregate_propertytype] = abstractds_aggregate_type
		}
	}

	// Compute reverse map for named struct DS_DataSet_Type
	// insertion point per field
	stage.DS_DataSet_Type_Has_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_DataSet_Type)
	for ds_dataset_type := range stage.DS_DataSet_Types {
		_ = ds_dataset_type
		for _, _md_metadata_propertytype := range ds_dataset_type.Has {
			stage.DS_DataSet_Type_Has_reverseMap[_md_metadata_propertytype] = ds_dataset_type
		}
	}
	stage.DS_DataSet_Type_PartOf_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_DataSet_Type)
	for ds_dataset_type := range stage.DS_DataSet_Types {
		_ = ds_dataset_type
		for _, _ds_aggregate_propertytype := range ds_dataset_type.PartOf {
			stage.DS_DataSet_Type_PartOf_reverseMap[_ds_aggregate_propertytype] = ds_dataset_type
		}
	}

	// Compute reverse map for named struct DS_Initiative_Type
	// insertion point per field
	stage.DS_Initiative_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_Initiative_Type)
	for ds_initiative_type := range stage.DS_Initiative_Types {
		_ = ds_initiative_type
		for _, _ds_dataset_propertytype := range ds_initiative_type.ComposedOf {
			stage.DS_Initiative_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_initiative_type
		}
	}
	stage.DS_Initiative_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_Initiative_Type)
	for ds_initiative_type := range stage.DS_Initiative_Types {
		_ = ds_initiative_type
		for _, _md_metadata_propertytype := range ds_initiative_type.SeriesMetadata {
			stage.DS_Initiative_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_initiative_type
		}
	}
	stage.DS_Initiative_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Initiative_Type)
	for ds_initiative_type := range stage.DS_Initiative_Types {
		_ = ds_initiative_type
		for _, _ds_aggregate_propertytype := range ds_initiative_type.Subset {
			stage.DS_Initiative_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_initiative_type
		}
	}
	stage.DS_Initiative_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Initiative_Type)
	for ds_initiative_type := range stage.DS_Initiative_Types {
		_ = ds_initiative_type
		for _, _ds_aggregate_propertytype := range ds_initiative_type.Superset {
			stage.DS_Initiative_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_initiative_type
		}
	}

	// Compute reverse map for named struct DS_OtherAggregate_Type
	// insertion point per field
	stage.DS_OtherAggregate_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_OtherAggregate_Type)
	for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
		_ = ds_otheraggregate_type
		for _, _ds_dataset_propertytype := range ds_otheraggregate_type.ComposedOf {
			stage.DS_OtherAggregate_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_otheraggregate_type
		}
	}
	stage.DS_OtherAggregate_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_OtherAggregate_Type)
	for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
		_ = ds_otheraggregate_type
		for _, _md_metadata_propertytype := range ds_otheraggregate_type.SeriesMetadata {
			stage.DS_OtherAggregate_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_otheraggregate_type
		}
	}
	stage.DS_OtherAggregate_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_OtherAggregate_Type)
	for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
		_ = ds_otheraggregate_type
		for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Subset {
			stage.DS_OtherAggregate_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_otheraggregate_type
		}
	}
	stage.DS_OtherAggregate_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_OtherAggregate_Type)
	for ds_otheraggregate_type := range stage.DS_OtherAggregate_Types {
		_ = ds_otheraggregate_type
		for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Superset {
			stage.DS_OtherAggregate_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_otheraggregate_type
		}
	}

	// Compute reverse map for named struct DS_Platform_Type
	// insertion point per field
	stage.DS_Platform_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_Platform_Type)
	for ds_platform_type := range stage.DS_Platform_Types {
		_ = ds_platform_type
		for _, _ds_dataset_propertytype := range ds_platform_type.ComposedOf {
			stage.DS_Platform_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_platform_type
		}
	}
	stage.DS_Platform_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_Platform_Type)
	for ds_platform_type := range stage.DS_Platform_Types {
		_ = ds_platform_type
		for _, _md_metadata_propertytype := range ds_platform_type.SeriesMetadata {
			stage.DS_Platform_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_platform_type
		}
	}
	stage.DS_Platform_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Platform_Type)
	for ds_platform_type := range stage.DS_Platform_Types {
		_ = ds_platform_type
		for _, _ds_aggregate_propertytype := range ds_platform_type.Subset {
			stage.DS_Platform_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_platform_type
		}
	}
	stage.DS_Platform_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Platform_Type)
	for ds_platform_type := range stage.DS_Platform_Types {
		_ = ds_platform_type
		for _, _ds_aggregate_propertytype := range ds_platform_type.Superset {
			stage.DS_Platform_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_platform_type
		}
	}

	// Compute reverse map for named struct DS_ProductionSeries_Type
	// insertion point per field
	stage.DS_ProductionSeries_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_ProductionSeries_Type)
	for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
		_ = ds_productionseries_type
		for _, _ds_dataset_propertytype := range ds_productionseries_type.ComposedOf {
			stage.DS_ProductionSeries_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_productionseries_type
		}
	}
	stage.DS_ProductionSeries_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_ProductionSeries_Type)
	for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
		_ = ds_productionseries_type
		for _, _md_metadata_propertytype := range ds_productionseries_type.SeriesMetadata {
			stage.DS_ProductionSeries_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_productionseries_type
		}
	}
	stage.DS_ProductionSeries_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_ProductionSeries_Type)
	for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
		_ = ds_productionseries_type
		for _, _ds_aggregate_propertytype := range ds_productionseries_type.Subset {
			stage.DS_ProductionSeries_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_productionseries_type
		}
	}
	stage.DS_ProductionSeries_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_ProductionSeries_Type)
	for ds_productionseries_type := range stage.DS_ProductionSeries_Types {
		_ = ds_productionseries_type
		for _, _ds_aggregate_propertytype := range ds_productionseries_type.Superset {
			stage.DS_ProductionSeries_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_productionseries_type
		}
	}

	// Compute reverse map for named struct DS_Sensor_Type
	// insertion point per field
	stage.DS_Sensor_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_Sensor_Type)
	for ds_sensor_type := range stage.DS_Sensor_Types {
		_ = ds_sensor_type
		for _, _ds_dataset_propertytype := range ds_sensor_type.ComposedOf {
			stage.DS_Sensor_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_sensor_type
		}
	}
	stage.DS_Sensor_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_Sensor_Type)
	for ds_sensor_type := range stage.DS_Sensor_Types {
		_ = ds_sensor_type
		for _, _md_metadata_propertytype := range ds_sensor_type.SeriesMetadata {
			stage.DS_Sensor_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_sensor_type
		}
	}
	stage.DS_Sensor_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Sensor_Type)
	for ds_sensor_type := range stage.DS_Sensor_Types {
		_ = ds_sensor_type
		for _, _ds_aggregate_propertytype := range ds_sensor_type.Subset {
			stage.DS_Sensor_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_sensor_type
		}
	}
	stage.DS_Sensor_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Sensor_Type)
	for ds_sensor_type := range stage.DS_Sensor_Types {
		_ = ds_sensor_type
		for _, _ds_aggregate_propertytype := range ds_sensor_type.Superset {
			stage.DS_Sensor_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_sensor_type
		}
	}

	// Compute reverse map for named struct DS_Series_Type
	// insertion point per field
	stage.DS_Series_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_Series_Type)
	for ds_series_type := range stage.DS_Series_Types {
		_ = ds_series_type
		for _, _ds_dataset_propertytype := range ds_series_type.ComposedOf {
			stage.DS_Series_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_series_type
		}
	}
	stage.DS_Series_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_Series_Type)
	for ds_series_type := range stage.DS_Series_Types {
		_ = ds_series_type
		for _, _md_metadata_propertytype := range ds_series_type.SeriesMetadata {
			stage.DS_Series_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_series_type
		}
	}
	stage.DS_Series_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Series_Type)
	for ds_series_type := range stage.DS_Series_Types {
		_ = ds_series_type
		for _, _ds_aggregate_propertytype := range ds_series_type.Subset {
			stage.DS_Series_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_series_type
		}
	}
	stage.DS_Series_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_Series_Type)
	for ds_series_type := range stage.DS_Series_Types {
		_ = ds_series_type
		for _, _ds_aggregate_propertytype := range ds_series_type.Superset {
			stage.DS_Series_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_series_type
		}
	}

	// Compute reverse map for named struct DS_StereoMate_Type
	// insertion point per field
	stage.DS_StereoMate_Type_ComposedOf_reverseMap = make(map[*DS_DataSet_PropertyType]*DS_StereoMate_Type)
	for ds_stereomate_type := range stage.DS_StereoMate_Types {
		_ = ds_stereomate_type
		for _, _ds_dataset_propertytype := range ds_stereomate_type.ComposedOf {
			stage.DS_StereoMate_Type_ComposedOf_reverseMap[_ds_dataset_propertytype] = ds_stereomate_type
		}
	}
	stage.DS_StereoMate_Type_SeriesMetadata_reverseMap = make(map[*MD_Metadata_PropertyType]*DS_StereoMate_Type)
	for ds_stereomate_type := range stage.DS_StereoMate_Types {
		_ = ds_stereomate_type
		for _, _md_metadata_propertytype := range ds_stereomate_type.SeriesMetadata {
			stage.DS_StereoMate_Type_SeriesMetadata_reverseMap[_md_metadata_propertytype] = ds_stereomate_type
		}
	}
	stage.DS_StereoMate_Type_Subset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_StereoMate_Type)
	for ds_stereomate_type := range stage.DS_StereoMate_Types {
		_ = ds_stereomate_type
		for _, _ds_aggregate_propertytype := range ds_stereomate_type.Subset {
			stage.DS_StereoMate_Type_Subset_reverseMap[_ds_aggregate_propertytype] = ds_stereomate_type
		}
	}
	stage.DS_StereoMate_Type_Superset_reverseMap = make(map[*DS_Aggregate_PropertyType]*DS_StereoMate_Type)
	for ds_stereomate_type := range stage.DS_StereoMate_Types {
		_ = ds_stereomate_type
		for _, _ds_aggregate_propertytype := range ds_stereomate_type.Superset {
			stage.DS_StereoMate_Type_Superset_reverseMap[_ds_aggregate_propertytype] = ds_stereomate_type
		}
	}

	// Compute reverse map for named struct MD_Metadata_Type
	// insertion point per field
	stage.MD_Metadata_Type_Series_reverseMap = make(map[*DS_Aggregate_PropertyType]*MD_Metadata_Type)
	for md_metadata_type := range stage.MD_Metadata_Types {
		_ = md_metadata_type
		for _, _ds_aggregate_propertytype := range md_metadata_type.Series {
			stage.MD_Metadata_Type_Series_reverseMap[_ds_aggregate_propertytype] = md_metadata_type
		}
	}
	stage.MD_Metadata_Type_Describes_reverseMap = make(map[*DS_DataSet_PropertyType]*MD_Metadata_Type)
	for md_metadata_type := range stage.MD_Metadata_Types {
		_ = md_metadata_type
		for _, _ds_dataset_propertytype := range md_metadata_type.Describes {
			stage.MD_Metadata_Type_Describes_reverseMap[_ds_dataset_propertytype] = md_metadata_type
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AbstractDS_Aggregate_Types)

	res = __gong__appendInstances(res, stage.DS_Aggregate_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_DataSet_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_DataSet_Types)

	res = __gong__appendInstances(res, stage.DS_Initiative_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_Initiative_Types)

	res = __gong__appendInstances(res, stage.DS_OtherAggregate_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_OtherAggregate_Types)

	res = __gong__appendInstances(res, stage.DS_Platform_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_Platform_Types)

	res = __gong__appendInstances(res, stage.DS_ProductionSeries_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_ProductionSeries_Types)

	res = __gong__appendInstances(res, stage.DS_Sensor_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_Sensor_Types)

	res = __gong__appendInstances(res, stage.DS_Series_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_Series_Types)

	res = __gong__appendInstances(res, stage.DS_StereoMate_PropertyTypes)

	res = __gong__appendInstances(res, stage.DS_StereoMate_Types)

	res = __gong__appendInstances(res, stage.MD_Metadata_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Metadata_Types)

	return
}

// insertion point per named struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongCopy() GongstructIF {
	newInstance := new(AbstractDS_Aggregate_Type)
	abstractds_aggregate_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_Aggregate_PropertyType)
	ds_aggregate_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_DataSet_PropertyType)
	ds_dataset_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_dataset_type *DS_DataSet_Type) GongCopy() GongstructIF {
	newInstance := new(DS_DataSet_Type)
	ds_dataset_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_Initiative_PropertyType)
	ds_initiative_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_initiative_type *DS_Initiative_Type) GongCopy() GongstructIF {
	newInstance := new(DS_Initiative_Type)
	ds_initiative_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_OtherAggregate_PropertyType)
	ds_otheraggregate_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongCopy() GongstructIF {
	newInstance := new(DS_OtherAggregate_Type)
	ds_otheraggregate_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_Platform_PropertyType)
	ds_platform_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_platform_type *DS_Platform_Type) GongCopy() GongstructIF {
	newInstance := new(DS_Platform_Type)
	ds_platform_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_ProductionSeries_PropertyType)
	ds_productionseries_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongCopy() GongstructIF {
	newInstance := new(DS_ProductionSeries_Type)
	ds_productionseries_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_Sensor_PropertyType)
	ds_sensor_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_sensor_type *DS_Sensor_Type) GongCopy() GongstructIF {
	newInstance := new(DS_Sensor_Type)
	ds_sensor_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_series_propertytype *DS_Series_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_Series_PropertyType)
	ds_series_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_series_type *DS_Series_Type) GongCopy() GongstructIF {
	newInstance := new(DS_Series_Type)
	ds_series_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongCopy() GongstructIF {
	newInstance := new(DS_StereoMate_PropertyType)
	ds_stereomate_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ds_stereomate_type *DS_StereoMate_Type) GongCopy() GongstructIF {
	newInstance := new(DS_StereoMate_Type)
	ds_stereomate_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_Metadata_PropertyType)
	md_metadata_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_metadata_type *MD_Metadata_Type) GongCopy() GongstructIF {
	newInstance := new(MD_Metadata_Type)
	md_metadata_type.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractds_aggregate_type)
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_aggregate_propertytype)
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_dataset_propertytype)
}

func (ds_dataset_type *DS_DataSet_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_dataset_type)
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_initiative_propertytype)
}

func (ds_initiative_type *DS_Initiative_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_initiative_type)
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_otheraggregate_propertytype)
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_otheraggregate_type)
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_platform_propertytype)
}

func (ds_platform_type *DS_Platform_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_platform_type)
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_productionseries_propertytype)
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_productionseries_type)
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_sensor_propertytype)
}

func (ds_sensor_type *DS_Sensor_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_sensor_type)
}

func (ds_series_propertytype *DS_Series_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_series_propertytype)
}

func (ds_series_type *DS_Series_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_series_type)
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_stereomate_propertytype)
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ds_stereomate_type)
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_metadata_propertytype)
}

func (md_metadata_type *MD_Metadata_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_metadata_type)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.AbstractDS_Aggregate_Types,
		stage.AbstractDS_Aggregate_Type_stagedOrder,
		stage.AbstractDS_Aggregate_Types_reference,
		&stage.AbstractDS_Aggregate_Types_referenceOrder,
		stage.AbstractDS_Aggregate_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Aggregate_PropertyTypes,
		stage.DS_Aggregate_PropertyType_stagedOrder,
		stage.DS_Aggregate_PropertyTypes_reference,
		&stage.DS_Aggregate_PropertyTypes_referenceOrder,
		stage.DS_Aggregate_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_DataSet_PropertyTypes,
		stage.DS_DataSet_PropertyType_stagedOrder,
		stage.DS_DataSet_PropertyTypes_reference,
		&stage.DS_DataSet_PropertyTypes_referenceOrder,
		stage.DS_DataSet_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_DataSet_Types,
		stage.DS_DataSet_Type_stagedOrder,
		stage.DS_DataSet_Types_reference,
		&stage.DS_DataSet_Types_referenceOrder,
		stage.DS_DataSet_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Initiative_PropertyTypes,
		stage.DS_Initiative_PropertyType_stagedOrder,
		stage.DS_Initiative_PropertyTypes_reference,
		&stage.DS_Initiative_PropertyTypes_referenceOrder,
		stage.DS_Initiative_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Initiative_Types,
		stage.DS_Initiative_Type_stagedOrder,
		stage.DS_Initiative_Types_reference,
		&stage.DS_Initiative_Types_referenceOrder,
		stage.DS_Initiative_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_OtherAggregate_PropertyTypes,
		stage.DS_OtherAggregate_PropertyType_stagedOrder,
		stage.DS_OtherAggregate_PropertyTypes_reference,
		&stage.DS_OtherAggregate_PropertyTypes_referenceOrder,
		stage.DS_OtherAggregate_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_OtherAggregate_Types,
		stage.DS_OtherAggregate_Type_stagedOrder,
		stage.DS_OtherAggregate_Types_reference,
		&stage.DS_OtherAggregate_Types_referenceOrder,
		stage.DS_OtherAggregate_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Platform_PropertyTypes,
		stage.DS_Platform_PropertyType_stagedOrder,
		stage.DS_Platform_PropertyTypes_reference,
		&stage.DS_Platform_PropertyTypes_referenceOrder,
		stage.DS_Platform_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Platform_Types,
		stage.DS_Platform_Type_stagedOrder,
		stage.DS_Platform_Types_reference,
		&stage.DS_Platform_Types_referenceOrder,
		stage.DS_Platform_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_ProductionSeries_PropertyTypes,
		stage.DS_ProductionSeries_PropertyType_stagedOrder,
		stage.DS_ProductionSeries_PropertyTypes_reference,
		&stage.DS_ProductionSeries_PropertyTypes_referenceOrder,
		stage.DS_ProductionSeries_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_ProductionSeries_Types,
		stage.DS_ProductionSeries_Type_stagedOrder,
		stage.DS_ProductionSeries_Types_reference,
		&stage.DS_ProductionSeries_Types_referenceOrder,
		stage.DS_ProductionSeries_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Sensor_PropertyTypes,
		stage.DS_Sensor_PropertyType_stagedOrder,
		stage.DS_Sensor_PropertyTypes_reference,
		&stage.DS_Sensor_PropertyTypes_referenceOrder,
		stage.DS_Sensor_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Sensor_Types,
		stage.DS_Sensor_Type_stagedOrder,
		stage.DS_Sensor_Types_reference,
		&stage.DS_Sensor_Types_referenceOrder,
		stage.DS_Sensor_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Series_PropertyTypes,
		stage.DS_Series_PropertyType_stagedOrder,
		stage.DS_Series_PropertyTypes_reference,
		&stage.DS_Series_PropertyTypes_referenceOrder,
		stage.DS_Series_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_Series_Types,
		stage.DS_Series_Type_stagedOrder,
		stage.DS_Series_Types_reference,
		&stage.DS_Series_Types_referenceOrder,
		stage.DS_Series_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_StereoMate_PropertyTypes,
		stage.DS_StereoMate_PropertyType_stagedOrder,
		stage.DS_StereoMate_PropertyTypes_reference,
		&stage.DS_StereoMate_PropertyTypes_referenceOrder,
		stage.DS_StereoMate_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DS_StereoMate_Types,
		stage.DS_StereoMate_Type_stagedOrder,
		stage.DS_StereoMate_Types_reference,
		&stage.DS_StereoMate_Types_referenceOrder,
		stage.DS_StereoMate_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_Metadata_PropertyTypes,
		stage.MD_Metadata_PropertyType_stagedOrder,
		stage.MD_Metadata_PropertyTypes_reference,
		&stage.MD_Metadata_PropertyTypes_referenceOrder,
		stage.MD_Metadata_PropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MD_Metadata_Types,
		stage.MD_Metadata_Type_stagedOrder,
		stage.MD_Metadata_Types_reference,
		&stage.MD_Metadata_Types_referenceOrder,
		stage.MD_Metadata_Types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.AbstractDS_Aggregate_Types, &stage.AbstractDS_Aggregate_Types_reference, &stage.AbstractDS_Aggregate_Types_referenceOrder, &stage.AbstractDS_Aggregate_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_Aggregate_PropertyTypes, &stage.DS_Aggregate_PropertyTypes_reference, &stage.DS_Aggregate_PropertyTypes_referenceOrder, &stage.DS_Aggregate_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_DataSet_PropertyTypes, &stage.DS_DataSet_PropertyTypes_reference, &stage.DS_DataSet_PropertyTypes_referenceOrder, &stage.DS_DataSet_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_DataSet_Types, &stage.DS_DataSet_Types_reference, &stage.DS_DataSet_Types_referenceOrder, &stage.DS_DataSet_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_Initiative_PropertyTypes, &stage.DS_Initiative_PropertyTypes_reference, &stage.DS_Initiative_PropertyTypes_referenceOrder, &stage.DS_Initiative_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_Initiative_Types, &stage.DS_Initiative_Types_reference, &stage.DS_Initiative_Types_referenceOrder, &stage.DS_Initiative_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_OtherAggregate_PropertyTypes, &stage.DS_OtherAggregate_PropertyTypes_reference, &stage.DS_OtherAggregate_PropertyTypes_referenceOrder, &stage.DS_OtherAggregate_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_OtherAggregate_Types, &stage.DS_OtherAggregate_Types_reference, &stage.DS_OtherAggregate_Types_referenceOrder, &stage.DS_OtherAggregate_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_Platform_PropertyTypes, &stage.DS_Platform_PropertyTypes_reference, &stage.DS_Platform_PropertyTypes_referenceOrder, &stage.DS_Platform_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_Platform_Types, &stage.DS_Platform_Types_reference, &stage.DS_Platform_Types_referenceOrder, &stage.DS_Platform_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_ProductionSeries_PropertyTypes, &stage.DS_ProductionSeries_PropertyTypes_reference, &stage.DS_ProductionSeries_PropertyTypes_referenceOrder, &stage.DS_ProductionSeries_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_ProductionSeries_Types, &stage.DS_ProductionSeries_Types_reference, &stage.DS_ProductionSeries_Types_referenceOrder, &stage.DS_ProductionSeries_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_Sensor_PropertyTypes, &stage.DS_Sensor_PropertyTypes_reference, &stage.DS_Sensor_PropertyTypes_referenceOrder, &stage.DS_Sensor_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_Sensor_Types, &stage.DS_Sensor_Types_reference, &stage.DS_Sensor_Types_referenceOrder, &stage.DS_Sensor_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_Series_PropertyTypes, &stage.DS_Series_PropertyTypes_reference, &stage.DS_Series_PropertyTypes_referenceOrder, &stage.DS_Series_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_Series_Types, &stage.DS_Series_Types_reference, &stage.DS_Series_Types_referenceOrder, &stage.DS_Series_Types_instance)

	__gong__computeReferencePass1(stage, stage.DS_StereoMate_PropertyTypes, &stage.DS_StereoMate_PropertyTypes_reference, &stage.DS_StereoMate_PropertyTypes_referenceOrder, &stage.DS_StereoMate_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DS_StereoMate_Types, &stage.DS_StereoMate_Types_reference, &stage.DS_StereoMate_Types_referenceOrder, &stage.DS_StereoMate_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_Metadata_PropertyTypes, &stage.MD_Metadata_PropertyTypes_reference, &stage.MD_Metadata_PropertyTypes_referenceOrder, &stage.MD_Metadata_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Metadata_Types, &stage.MD_Metadata_Types_reference, &stage.MD_Metadata_Types_referenceOrder, &stage.MD_Metadata_Types_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_DataSet_Types, stage.DS_DataSet_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_Initiative_Types, stage.DS_Initiative_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_Platform_Types, stage.DS_Platform_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_Sensor_Types, stage.DS_Sensor_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_Series_Types, stage.DS_Series_Types_reference, stage)

	__gong__computeReferencePass2(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DS_StereoMate_Types, stage.DS_StereoMate_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Metadata_Types, stage.MD_Metadata_Types_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractDS_Aggregate_Type_stagedOrder, stage.AbstractDS_Aggregate_Types_referenceOrder, abstractds_aggregate_type, "AbstractDS_Aggregate_Type")
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Aggregate_PropertyType_stagedOrder, stage.DS_Aggregate_PropertyTypes_referenceOrder, ds_aggregate_propertytype, "DS_Aggregate_PropertyType")
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_DataSet_PropertyType_stagedOrder, stage.DS_DataSet_PropertyTypes_referenceOrder, ds_dataset_propertytype, "DS_DataSet_PropertyType")
}

func (ds_dataset_type *DS_DataSet_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_DataSet_Type_stagedOrder, stage.DS_DataSet_Types_referenceOrder, ds_dataset_type, "DS_DataSet_Type")
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Initiative_PropertyType_stagedOrder, stage.DS_Initiative_PropertyTypes_referenceOrder, ds_initiative_propertytype, "DS_Initiative_PropertyType")
}

func (ds_initiative_type *DS_Initiative_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Initiative_Type_stagedOrder, stage.DS_Initiative_Types_referenceOrder, ds_initiative_type, "DS_Initiative_Type")
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_OtherAggregate_PropertyType_stagedOrder, stage.DS_OtherAggregate_PropertyTypes_referenceOrder, ds_otheraggregate_propertytype, "DS_OtherAggregate_PropertyType")
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_OtherAggregate_Type_stagedOrder, stage.DS_OtherAggregate_Types_referenceOrder, ds_otheraggregate_type, "DS_OtherAggregate_Type")
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Platform_PropertyType_stagedOrder, stage.DS_Platform_PropertyTypes_referenceOrder, ds_platform_propertytype, "DS_Platform_PropertyType")
}

func (ds_platform_type *DS_Platform_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Platform_Type_stagedOrder, stage.DS_Platform_Types_referenceOrder, ds_platform_type, "DS_Platform_Type")
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_ProductionSeries_PropertyType_stagedOrder, stage.DS_ProductionSeries_PropertyTypes_referenceOrder, ds_productionseries_propertytype, "DS_ProductionSeries_PropertyType")
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_ProductionSeries_Type_stagedOrder, stage.DS_ProductionSeries_Types_referenceOrder, ds_productionseries_type, "DS_ProductionSeries_Type")
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Sensor_PropertyType_stagedOrder, stage.DS_Sensor_PropertyTypes_referenceOrder, ds_sensor_propertytype, "DS_Sensor_PropertyType")
}

func (ds_sensor_type *DS_Sensor_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Sensor_Type_stagedOrder, stage.DS_Sensor_Types_referenceOrder, ds_sensor_type, "DS_Sensor_Type")
}

func (ds_series_propertytype *DS_Series_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Series_PropertyType_stagedOrder, stage.DS_Series_PropertyTypes_referenceOrder, ds_series_propertytype, "DS_Series_PropertyType")
}

func (ds_series_type *DS_Series_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_Series_Type_stagedOrder, stage.DS_Series_Types_referenceOrder, ds_series_type, "DS_Series_Type")
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_StereoMate_PropertyType_stagedOrder, stage.DS_StereoMate_PropertyTypes_referenceOrder, ds_stereomate_propertytype, "DS_StereoMate_PropertyType")
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DS_StereoMate_Type_stagedOrder, stage.DS_StereoMate_Types_referenceOrder, ds_stereomate_type, "DS_StereoMate_Type")
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Metadata_PropertyType_stagedOrder, stage.MD_Metadata_PropertyTypes_referenceOrder, md_metadata_propertytype, "MD_Metadata_PropertyType")
}

func (md_metadata_type *MD_Metadata_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Metadata_Type_stagedOrder, stage.MD_Metadata_Types_referenceOrder, md_metadata_type, "MD_Metadata_Type")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractds_aggregate_type, abstractds_aggregate_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractds_aggregate_type.GongGetIdentifier(stage)
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_aggregate_propertytype, ds_aggregate_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_aggregate_propertytype.GongGetIdentifier(stage)
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_dataset_propertytype, ds_dataset_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_dataset_propertytype.GongGetIdentifier(stage)
}

func (ds_dataset_type *DS_DataSet_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_dataset_type, ds_dataset_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_dataset_type *DS_DataSet_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_dataset_type.GongGetIdentifier(stage)
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_initiative_propertytype, ds_initiative_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_initiative_propertytype.GongGetIdentifier(stage)
}

func (ds_initiative_type *DS_Initiative_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_initiative_type, ds_initiative_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_initiative_type *DS_Initiative_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_initiative_type.GongGetIdentifier(stage)
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_otheraggregate_propertytype, ds_otheraggregate_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_otheraggregate_propertytype.GongGetIdentifier(stage)
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_otheraggregate_type, ds_otheraggregate_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_otheraggregate_type.GongGetIdentifier(stage)
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_platform_propertytype, ds_platform_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_platform_propertytype *DS_Platform_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_platform_propertytype.GongGetIdentifier(stage)
}

func (ds_platform_type *DS_Platform_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_platform_type, ds_platform_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_platform_type *DS_Platform_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_platform_type.GongGetIdentifier(stage)
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_productionseries_propertytype, ds_productionseries_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_productionseries_propertytype.GongGetIdentifier(stage)
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_productionseries_type, ds_productionseries_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_productionseries_type *DS_ProductionSeries_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_productionseries_type.GongGetIdentifier(stage)
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_sensor_propertytype, ds_sensor_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_sensor_propertytype.GongGetIdentifier(stage)
}

func (ds_sensor_type *DS_Sensor_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_sensor_type, ds_sensor_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_sensor_type *DS_Sensor_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_sensor_type.GongGetIdentifier(stage)
}

func (ds_series_propertytype *DS_Series_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_series_propertytype, ds_series_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_series_propertytype *DS_Series_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_series_propertytype.GongGetIdentifier(stage)
}

func (ds_series_type *DS_Series_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_series_type, ds_series_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_series_type *DS_Series_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_series_type.GongGetIdentifier(stage)
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_stereomate_propertytype, ds_stereomate_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_stereomate_propertytype.GongGetIdentifier(stage)
}

func (ds_stereomate_type *DS_StereoMate_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ds_stereomate_type, ds_stereomate_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ds_stereomate_type *DS_StereoMate_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return ds_stereomate_type.GongGetIdentifier(stage)
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_metadata_propertytype, md_metadata_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_metadata_propertytype *MD_Metadata_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_metadata_propertytype.GongGetIdentifier(stage)
}

func (md_metadata_type *MD_Metadata_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_metadata_type, md_metadata_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_metadata_type *MD_Metadata_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_metadata_type.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractds_aggregate_type.GongGetIdentifier(stage), "AbstractDS_Aggregate_Type", abstractds_aggregate_type.Name)
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_aggregate_propertytype.GongGetIdentifier(stage), "DS_Aggregate_PropertyType", ds_aggregate_propertytype.Name)
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_dataset_propertytype.GongGetIdentifier(stage), "DS_DataSet_PropertyType", ds_dataset_propertytype.Name)
}

func (ds_dataset_type *DS_DataSet_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_dataset_type.GongGetIdentifier(stage), "DS_DataSet_Type", ds_dataset_type.Name)
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_initiative_propertytype.GongGetIdentifier(stage), "DS_Initiative_PropertyType", ds_initiative_propertytype.Name)
}

func (ds_initiative_type *DS_Initiative_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_initiative_type.GongGetIdentifier(stage), "DS_Initiative_Type", ds_initiative_type.Name)
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_otheraggregate_propertytype.GongGetIdentifier(stage), "DS_OtherAggregate_PropertyType", ds_otheraggregate_propertytype.Name)
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_otheraggregate_type.GongGetIdentifier(stage), "DS_OtherAggregate_Type", ds_otheraggregate_type.Name)
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_platform_propertytype.GongGetIdentifier(stage), "DS_Platform_PropertyType", ds_platform_propertytype.Name)
}

func (ds_platform_type *DS_Platform_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_platform_type.GongGetIdentifier(stage), "DS_Platform_Type", ds_platform_type.Name)
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_productionseries_propertytype.GongGetIdentifier(stage), "DS_ProductionSeries_PropertyType", ds_productionseries_propertytype.Name)
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_productionseries_type.GongGetIdentifier(stage), "DS_ProductionSeries_Type", ds_productionseries_type.Name)
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_sensor_propertytype.GongGetIdentifier(stage), "DS_Sensor_PropertyType", ds_sensor_propertytype.Name)
}

func (ds_sensor_type *DS_Sensor_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_sensor_type.GongGetIdentifier(stage), "DS_Sensor_Type", ds_sensor_type.Name)
}

func (ds_series_propertytype *DS_Series_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_series_propertytype.GongGetIdentifier(stage), "DS_Series_PropertyType", ds_series_propertytype.Name)
}

func (ds_series_type *DS_Series_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_series_type.GongGetIdentifier(stage), "DS_Series_Type", ds_series_type.Name)
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_stereomate_propertytype.GongGetIdentifier(stage), "DS_StereoMate_PropertyType", ds_stereomate_propertytype.Name)
}

func (ds_stereomate_type *DS_StereoMate_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ds_stereomate_type.GongGetIdentifier(stage), "DS_StereoMate_Type", ds_stereomate_type.Name)
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_metadata_propertytype.GongGetIdentifier(stage), "MD_Metadata_PropertyType", md_metadata_propertytype.Name)
}

func (md_metadata_type *MD_Metadata_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_metadata_type.GongGetIdentifier(stage), "MD_Metadata_Type", md_metadata_type.Name)
}

// insertion point for unstaging
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractds_aggregate_type.GongGetReferenceIdentifier(stage))
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_aggregate_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_dataset_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_dataset_type *DS_DataSet_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_dataset_type.GongGetReferenceIdentifier(stage))
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_initiative_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_initiative_type *DS_Initiative_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_initiative_type.GongGetReferenceIdentifier(stage))
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_otheraggregate_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_otheraggregate_type.GongGetReferenceIdentifier(stage))
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_platform_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_platform_type *DS_Platform_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_platform_type.GongGetReferenceIdentifier(stage))
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_productionseries_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_productionseries_type.GongGetReferenceIdentifier(stage))
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_sensor_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_sensor_type *DS_Sensor_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_sensor_type.GongGetReferenceIdentifier(stage))
}

func (ds_series_propertytype *DS_Series_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_series_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_series_type *DS_Series_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_series_type.GongGetReferenceIdentifier(stage))
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_stereomate_propertytype.GongGetReferenceIdentifier(stage))
}

func (ds_stereomate_type *DS_StereoMate_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ds_stereomate_type.GongGetReferenceIdentifier(stage))
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_metadata_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_metadata_type *MD_Metadata_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_metadata_type.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__gmd__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
