// generated code - do not edit
package spatialrepresentation

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
// Clean garbage collect unstaged instances that are referenced by MD_Dimension_PropertyType
func (md_dimension_propertytype *MD_Dimension_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_dimension_propertytype.MD_Dimension) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Dimension_Type
func (md_dimension_type *MD_Dimension_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_dimension_type.DimensionName) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_GeometricObjects_PropertyType
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_geometricobjects_propertytype.MD_GeometricObjects) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_GeometricObjects_Type
func (md_geometricobjects_type *MD_GeometricObjects_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_geometricobjects_type.GeometricObjectType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Georectified_PropertyType
func (md_georectified_propertytype *MD_Georectified_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_georectified_propertytype.MD_Georectified) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Georectified_Type
func (md_georectified_type *MD_Georectified_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_georectified_type.AxisDimensionProperties) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_georectified_type.CellGeometry) || modified
	modified = stage.CleanPointer(&md_georectified_type.PointInPixel) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Georeferenceable_PropertyType
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_georeferenceable_propertytype.MD_Georeferenceable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Georeferenceable_Type
func (md_georeferenceable_type *MD_Georeferenceable_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_georeferenceable_type.AxisDimensionProperties) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_georeferenceable_type.CellGeometry) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_GridSpatialRepresentation_PropertyType
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation) || modified
	modified = stage.CleanPointer(&md_gridspatialrepresentation_propertytype.MD_Georeferenceable) || modified
	modified = stage.CleanPointer(&md_gridspatialrepresentation_propertytype.MD_Georectified) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_GridSpatialRepresentation_Type
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_gridspatialrepresentation_type.AxisDimensionProperties) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_gridspatialrepresentation_type.CellGeometry) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_SpatialRepresentation_PropertyType
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation) || modified
	modified = stage.CleanPointer(&md_spatialrepresentation_propertytype.MD_Georeferenceable) || modified
	modified = stage.CleanPointer(&md_spatialrepresentation_propertytype.MD_Georectified) || modified
	modified = stage.CleanPointer(&md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_VectorSpatialRepresentation_PropertyType
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_VectorSpatialRepresentation_Type
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_vectorspatialrepresentation_type.GeometricObjects) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_vectorspatialrepresentation_type.TopologyLevel) || modified
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
