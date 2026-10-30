// generated code - do not edit
package spatialrepresentation

// insertion point
func (inst *AbstractMD_SpatialRepresentation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_CellGeometryCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_DimensionNameTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Dimension_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Georectified_Type":
		switch reverseField.Fieldname {
		case "AxisDimensionProperties":
			if _md_georectified_type, ok := stage.MD_Georectified_Type_AxisDimensionProperties_reverseMap[inst]; ok {
				res = _md_georectified_type.Name
			}
		}
	case "MD_Georeferenceable_Type":
		switch reverseField.Fieldname {
		case "AxisDimensionProperties":
			if _md_georeferenceable_type, ok := stage.MD_Georeferenceable_Type_AxisDimensionProperties_reverseMap[inst]; ok {
				res = _md_georeferenceable_type.Name
			}
		}
	case "MD_GridSpatialRepresentation_Type":
		switch reverseField.Fieldname {
		case "AxisDimensionProperties":
			if _md_gridspatialrepresentation_type, ok := stage.MD_GridSpatialRepresentation_Type_AxisDimensionProperties_reverseMap[inst]; ok {
				res = _md_gridspatialrepresentation_type.Name
			}
		}
	}
	return
}

func (inst *MD_Dimension_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_GeometricObjectTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_GeometricObjects_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_VectorSpatialRepresentation_Type":
		switch reverseField.Fieldname {
		case "GeometricObjects":
			if _md_vectorspatialrepresentation_type, ok := stage.MD_VectorSpatialRepresentation_Type_GeometricObjects_reverseMap[inst]; ok {
				res = _md_vectorspatialrepresentation_type.Name
			}
		}
	}
	return
}

func (inst *MD_GeometricObjects_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Georectified_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Georectified_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Georeferenceable_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Georeferenceable_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_GridSpatialRepresentation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_GridSpatialRepresentation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_PixelOrientationCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_SpatialRepresentation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_TopologyLevelCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_VectorSpatialRepresentation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_VectorSpatialRepresentation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
